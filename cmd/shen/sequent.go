package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4datatype_1error Obj                      // shen.datatype-error
var __defun__shen_4_5datatype_1rules_6 Obj                  // shen.<datatype-rules>
var __defun__shen_4_5datatype_1rule_6 Obj                   // shen.<datatype-rule>
var __defun__shen_4_5side_1conditions_6 Obj                 // shen.<side-conditions>
var __defun__shen_4_5side_1condition_6 Obj                  // shen.<side-condition>
var __defun__shen_4_5variable_2_6 Obj                       // shen.<variable?>
var __defun__shen_4_5expr_6 Obj                             // shen.<expr>
var __defun__shen_4remove_1bar Obj                          // shen.remove-bar
var __defun__shen_4_5premises_6 Obj                         // shen.<premises>
var __defun__shen_4_5semicolon_1symbol_6 Obj                // shen.<semicolon-symbol>
var __defun__shen_4_5premise_6 Obj                          // shen.<premise>
var __defun__shen_4_5conclusion_6 Obj                       // shen.<conclusion>
var __defun__shen_4sequent Obj                              // shen.sequent
var __defun__shen_4_5formulae_6 Obj                         // shen.<formulae>
var __defun__shen_4_5comma_1symbol_6 Obj                    // shen.<comma-symbol>
var __defun__shen_4_5formula_6 Obj                          // shen.<formula>
var __defun__shen_4_5type_6 Obj                             // shen.<type>
var __defun__shen_4_5doubleunderline_6 Obj                  // shen.<doubleunderline>
var __defun__shen_4_5singleunderline_6 Obj                  // shen.<singleunderline>
var __defun__shen_4singleunderline_2 Obj                    // shen.singleunderline?
var __defun__shen_4sh_2 Obj                                 // shen.sh?
var __defun__shen_4doubleunderline_2 Obj                    // shen.doubleunderline?
var __defun__shen_4dh_2 Obj                                 // shen.dh?
var __defun__shen_4process_1datatype Obj                    // shen.process-datatype
var __defun__shen_4remember_1datatype Obj                   // shen.remember-datatype
var __defun__shen_4rules_1_6horn_1clauses Obj               // shen.rules->horn-clauses
var __defun__shen_4double_1_6singles Obj                    // shen.double->singles
var __defun__shen_4right_1rule Obj                          // shen.right-rule
var __defun__shen_4left_1rule Obj                           // shen.left-rule
var __defun__shen_4right_1_6left Obj                        // shen.right->left
var __defun__shen_4rule_1_6horn_1clause Obj                 // shen.rule->horn-clause
var __defun__shen_4rule_1_6horn_1clause_1head Obj           // shen.rule->horn-clause-head
var __defun__shen_4mode_1ify Obj                            // shen.mode-ify
var __defun__shen_4rule_1_6horn_1clause_1body Obj           // shen.rule->horn-clause-body
var __defun__shen_4construct_1search_1literals Obj          // shen.construct-search-literals
var __defun__shen_4csl_1help Obj                            // shen.csl-help
var __defun__shen_4construct_1search_1clauses Obj           // shen.construct-search-clauses
var __defun__shen_4construct_1search_1clause Obj            // shen.construct-search-clause
var __defun__shen_4construct_1base_1search_1clause Obj      // shen.construct-base-search-clause
var __defun__shen_4construct_1recursive_1search_1clause Obj // shen.construct-recursive-search-clause
var __defun__shen_4construct_1side_1literals Obj            // shen.construct-side-literals
var __defun__shen_4construct_1premiss_1literal Obj          // shen.construct-premiss-literal
var __defun__shen_4construct_1context Obj                   // shen.construct-context
var __defun__shen_4recursive__cons__form Obj                // shen.recursive_cons_form
var __defun__preclude Obj                                   // preclude
var __defun__shen_4preclude_1h Obj                          // shen.preclude-h
var __defun__include Obj                                    // include
var __defun__shen_4include_1h Obj                           // shen.include-h
var __defun__preclude_1all_1but Obj                         // preclude-all-but
var __defun__include_1all_1but Obj                          // include-all-but
var __defun__shen_4synonyms_1help Obj                       // shen.synonyms-help
var __defun__shen_4pushnew Obj                              // shen.pushnew
var __defun__shen_4demod_1rule Obj                          // shen.demod-rule
var __defun__shen_4lambda_1of_1defun Obj                    // shen.lambda-of-defun
var __defun__shen_4update_1demodulation_1function Obj       // shen.update-demodulation-function
var __defun__shen_4default_1rule Obj                        // shen.default-rule

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		reg5949 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__e.Return(reg5949)
		return
	}, 0))
	__defun__shen_4datatype_1error = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1692 := __args[0]
		_ = V1692
		reg5950 := PrimIsPair(V1692)
		var reg5966 Obj
		if reg5950 == True {
			reg5951 := PrimTail(V1692)
			reg5952 := PrimIsPair(reg5951)
			var reg5961 Obj
			if reg5952 == True {
				reg5953 := Nil
				reg5954 := PrimTail(V1692)
				reg5955 := PrimTail(reg5954)
				reg5956 := PrimEqual(reg5953, reg5955)
				var reg5959 Obj
				if reg5956 == True {
					reg5957 := True
					reg5959 = reg5957
				} else {
					reg5958 := False
					reg5959 = reg5958
				}
				reg5961 = reg5959
			} else {
				reg5960 := False
				reg5961 = reg5960
			}
			var reg5964 Obj
			if reg5961 == True {
				reg5962 := True
				reg5964 = reg5962
			} else {
				reg5963 := False
				reg5964 = reg5963
			}
			reg5966 = reg5964
		} else {
			reg5965 := False
			reg5966 = reg5965
		}
		if reg5966 == True {
			reg5967 := MakeString("datatype syntax error here:\n\n ")
			reg5968 := MakeNumber(50)
			reg5969 := PrimHead(V1692)
			reg5970 := Call(__e, __defun__shen_4next_150, reg5968, reg5969)
			reg5971 := MakeString("\n")
			reg5972 := MakeSymbol("shen.a")
			reg5973 := Call(__e, __defun__shen_4app, reg5970, reg5971, reg5972)
			reg5974 := PrimStringConcat(reg5967, reg5973)
			reg5975 := PrimSimpleError(reg5974)
			__e.Return(reg5975)
			return
		} else {
			reg5976 := MakeSymbol("shen.datatype-error")
			__e.TailApply(__defun__shen_4f__error, reg5976)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.datatype-error", value: __defun__shen_4datatype_1error})

	__defun__shen_4_5datatype_1rules_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1694 := __args[0]
		_ = V1694
		reg5978 := Call(__e, __defun__shen_4_5datatype_1rule_6, V1694)
		Parse__shen_4_5datatype_1rule_6 := reg5978
		_ = Parse__shen_4_5datatype_1rule_6
		reg5979 := Call(__e, __defun__fail)
		reg5980 := PrimEqual(reg5979, Parse__shen_4_5datatype_1rule_6)
		reg5981 := PrimNot(reg5980)
		var reg5994 Obj
		if reg5981 == True {
			reg5982 := Call(__e, __defun__shen_4_5datatype_1rules_6, Parse__shen_4_5datatype_1rule_6)
			Parse__shen_4_5datatype_1rules_6 := reg5982
			_ = Parse__shen_4_5datatype_1rules_6
			reg5983 := Call(__e, __defun__fail)
			reg5984 := PrimEqual(reg5983, Parse__shen_4_5datatype_1rules_6)
			reg5985 := PrimNot(reg5984)
			var reg5992 Obj
			if reg5985 == True {
				reg5986 := PrimHead(Parse__shen_4_5datatype_1rules_6)
				reg5987 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5datatype_1rule_6)
				reg5988 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5datatype_1rules_6)
				reg5989 := PrimCons(reg5987, reg5988)
				reg5990 := Call(__e, __defun__shen_4pair, reg5986, reg5989)
				reg5992 = reg5990
			} else {
				reg5991 := Call(__e, __defun__fail)
				reg5992 = reg5991
			}
			reg5994 = reg5992
		} else {
			reg5993 := Call(__e, __defun__fail)
			reg5994 = reg5993
		}
		YaccParse := reg5994
		_ = YaccParse
		reg5995 := Call(__e, __defun__fail)
		reg5996 := PrimEqual(YaccParse, reg5995)
		if reg5996 == True {
			reg5997 := Call(__e, __defun___5e_6, V1694)
			Parse___5e_6 := reg5997
			_ = Parse___5e_6
			reg5998 := Call(__e, __defun__fail)
			reg5999 := PrimEqual(reg5998, Parse___5e_6)
			reg6000 := PrimNot(reg5999)
			if reg6000 == True {
				reg6001 := PrimHead(Parse___5e_6)
				reg6002 := Nil
				__e.TailApply(__defun__shen_4pair, reg6001, reg6002)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<datatype-rules>", value: __defun__shen_4_5datatype_1rules_6})

	__defun__shen_4_5datatype_1rule_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1696 := __args[0]
		_ = V1696
		reg6005 := Call(__e, __defun__shen_4_5side_1conditions_6, V1696)
		Parse__shen_4_5side_1conditions_6 := reg6005
		_ = Parse__shen_4_5side_1conditions_6
		reg6006 := Call(__e, __defun__fail)
		reg6007 := PrimEqual(reg6006, Parse__shen_4_5side_1conditions_6)
		reg6008 := PrimNot(reg6007)
		var reg6039 Obj
		if reg6008 == True {
			reg6009 := Call(__e, __defun__shen_4_5premises_6, Parse__shen_4_5side_1conditions_6)
			Parse__shen_4_5premises_6 := reg6009
			_ = Parse__shen_4_5premises_6
			reg6010 := Call(__e, __defun__fail)
			reg6011 := PrimEqual(reg6010, Parse__shen_4_5premises_6)
			reg6012 := PrimNot(reg6011)
			var reg6037 Obj
			if reg6012 == True {
				reg6013 := Call(__e, __defun__shen_4_5singleunderline_6, Parse__shen_4_5premises_6)
				Parse__shen_4_5singleunderline_6 := reg6013
				_ = Parse__shen_4_5singleunderline_6
				reg6014 := Call(__e, __defun__fail)
				reg6015 := PrimEqual(reg6014, Parse__shen_4_5singleunderline_6)
				reg6016 := PrimNot(reg6015)
				var reg6035 Obj
				if reg6016 == True {
					reg6017 := Call(__e, __defun__shen_4_5conclusion_6, Parse__shen_4_5singleunderline_6)
					Parse__shen_4_5conclusion_6 := reg6017
					_ = Parse__shen_4_5conclusion_6
					reg6018 := Call(__e, __defun__fail)
					reg6019 := PrimEqual(reg6018, Parse__shen_4_5conclusion_6)
					reg6020 := PrimNot(reg6019)
					var reg6033 Obj
					if reg6020 == True {
						reg6021 := PrimHead(Parse__shen_4_5conclusion_6)
						reg6022 := MakeSymbol("shen.single")
						reg6023 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5side_1conditions_6)
						reg6024 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5premises_6)
						reg6025 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5conclusion_6)
						reg6026 := Nil
						reg6027 := PrimCons(reg6025, reg6026)
						reg6028 := PrimCons(reg6024, reg6027)
						reg6029 := PrimCons(reg6023, reg6028)
						reg6030 := Call(__e, __defun__shen_4sequent, reg6022, reg6029)
						reg6031 := Call(__e, __defun__shen_4pair, reg6021, reg6030)
						reg6033 = reg6031
					} else {
						reg6032 := Call(__e, __defun__fail)
						reg6033 = reg6032
					}
					reg6035 = reg6033
				} else {
					reg6034 := Call(__e, __defun__fail)
					reg6035 = reg6034
				}
				reg6037 = reg6035
			} else {
				reg6036 := Call(__e, __defun__fail)
				reg6037 = reg6036
			}
			reg6039 = reg6037
		} else {
			reg6038 := Call(__e, __defun__fail)
			reg6039 = reg6038
		}
		YaccParse := reg6039
		_ = YaccParse
		reg6040 := Call(__e, __defun__fail)
		reg6041 := PrimEqual(YaccParse, reg6040)
		if reg6041 == True {
			reg6042 := Call(__e, __defun__shen_4_5side_1conditions_6, V1696)
			Parse__shen_4_5side_1conditions_6 := reg6042
			_ = Parse__shen_4_5side_1conditions_6
			reg6043 := Call(__e, __defun__fail)
			reg6044 := PrimEqual(reg6043, Parse__shen_4_5side_1conditions_6)
			reg6045 := PrimNot(reg6044)
			if reg6045 == True {
				reg6046 := Call(__e, __defun__shen_4_5premises_6, Parse__shen_4_5side_1conditions_6)
				Parse__shen_4_5premises_6 := reg6046
				_ = Parse__shen_4_5premises_6
				reg6047 := Call(__e, __defun__fail)
				reg6048 := PrimEqual(reg6047, Parse__shen_4_5premises_6)
				reg6049 := PrimNot(reg6048)
				if reg6049 == True {
					reg6050 := Call(__e, __defun__shen_4_5doubleunderline_6, Parse__shen_4_5premises_6)
					Parse__shen_4_5doubleunderline_6 := reg6050
					_ = Parse__shen_4_5doubleunderline_6
					reg6051 := Call(__e, __defun__fail)
					reg6052 := PrimEqual(reg6051, Parse__shen_4_5doubleunderline_6)
					reg6053 := PrimNot(reg6052)
					if reg6053 == True {
						reg6054 := Call(__e, __defun__shen_4_5conclusion_6, Parse__shen_4_5doubleunderline_6)
						Parse__shen_4_5conclusion_6 := reg6054
						_ = Parse__shen_4_5conclusion_6
						reg6055 := Call(__e, __defun__fail)
						reg6056 := PrimEqual(reg6055, Parse__shen_4_5conclusion_6)
						reg6057 := PrimNot(reg6056)
						if reg6057 == True {
							reg6058 := PrimHead(Parse__shen_4_5conclusion_6)
							reg6059 := MakeSymbol("shen.double")
							reg6060 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5side_1conditions_6)
							reg6061 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5premises_6)
							reg6062 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5conclusion_6)
							reg6063 := Nil
							reg6064 := PrimCons(reg6062, reg6063)
							reg6065 := PrimCons(reg6061, reg6064)
							reg6066 := PrimCons(reg6060, reg6065)
							reg6067 := Call(__e, __defun__shen_4sequent, reg6059, reg6066)
							__e.TailApply(__defun__shen_4pair, reg6058, reg6067)
							return
						} else {
							__e.TailApply(__defun__fail)
							return
						}
					} else {
						__e.TailApply(__defun__fail)
						return
					}
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<datatype-rule>", value: __defun__shen_4_5datatype_1rule_6})

	__defun__shen_4_5side_1conditions_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1698 := __args[0]
		_ = V1698
		reg6073 := Call(__e, __defun__shen_4_5side_1condition_6, V1698)
		Parse__shen_4_5side_1condition_6 := reg6073
		_ = Parse__shen_4_5side_1condition_6
		reg6074 := Call(__e, __defun__fail)
		reg6075 := PrimEqual(reg6074, Parse__shen_4_5side_1condition_6)
		reg6076 := PrimNot(reg6075)
		var reg6089 Obj
		if reg6076 == True {
			reg6077 := Call(__e, __defun__shen_4_5side_1conditions_6, Parse__shen_4_5side_1condition_6)
			Parse__shen_4_5side_1conditions_6 := reg6077
			_ = Parse__shen_4_5side_1conditions_6
			reg6078 := Call(__e, __defun__fail)
			reg6079 := PrimEqual(reg6078, Parse__shen_4_5side_1conditions_6)
			reg6080 := PrimNot(reg6079)
			var reg6087 Obj
			if reg6080 == True {
				reg6081 := PrimHead(Parse__shen_4_5side_1conditions_6)
				reg6082 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5side_1condition_6)
				reg6083 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5side_1conditions_6)
				reg6084 := PrimCons(reg6082, reg6083)
				reg6085 := Call(__e, __defun__shen_4pair, reg6081, reg6084)
				reg6087 = reg6085
			} else {
				reg6086 := Call(__e, __defun__fail)
				reg6087 = reg6086
			}
			reg6089 = reg6087
		} else {
			reg6088 := Call(__e, __defun__fail)
			reg6089 = reg6088
		}
		YaccParse := reg6089
		_ = YaccParse
		reg6090 := Call(__e, __defun__fail)
		reg6091 := PrimEqual(YaccParse, reg6090)
		if reg6091 == True {
			reg6092 := Call(__e, __defun___5e_6, V1698)
			Parse___5e_6 := reg6092
			_ = Parse___5e_6
			reg6093 := Call(__e, __defun__fail)
			reg6094 := PrimEqual(reg6093, Parse___5e_6)
			reg6095 := PrimNot(reg6094)
			if reg6095 == True {
				reg6096 := PrimHead(Parse___5e_6)
				reg6097 := Nil
				__e.TailApply(__defun__shen_4pair, reg6096, reg6097)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<side-conditions>", value: __defun__shen_4_5side_1conditions_6})

	__defun__shen_4_5side_1condition_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1702 := __args[0]
		_ = V1702
		reg6100 := PrimHead(V1702)
		reg6101 := PrimIsPair(reg6100)
		var reg6109 Obj
		if reg6101 == True {
			reg6102 := MakeSymbol("if")
			reg6103 := Call(__e, __defun__shen_4hdhd, V1702)
			reg6104 := PrimEqual(reg6102, reg6103)
			var reg6107 Obj
			if reg6104 == True {
				reg6105 := True
				reg6107 = reg6105
			} else {
				reg6106 := False
				reg6107 = reg6106
			}
			reg6109 = reg6107
		} else {
			reg6108 := False
			reg6109 = reg6108
		}
		var reg6127 Obj
		if reg6109 == True {
			reg6110 := Call(__e, __defun__shen_4tlhd, V1702)
			reg6111 := Call(__e, __defun__shen_4hdtl, V1702)
			reg6112 := Call(__e, __defun__shen_4pair, reg6110, reg6111)
			NewStream1699 := reg6112
			_ = NewStream1699
			reg6113 := Call(__e, __defun__shen_4_5expr_6, NewStream1699)
			Parse__shen_4_5expr_6 := reg6113
			_ = Parse__shen_4_5expr_6
			reg6114 := Call(__e, __defun__fail)
			reg6115 := PrimEqual(reg6114, Parse__shen_4_5expr_6)
			reg6116 := PrimNot(reg6115)
			var reg6125 Obj
			if reg6116 == True {
				reg6117 := PrimHead(Parse__shen_4_5expr_6)
				reg6118 := MakeSymbol("if")
				reg6119 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5expr_6)
				reg6120 := Nil
				reg6121 := PrimCons(reg6119, reg6120)
				reg6122 := PrimCons(reg6118, reg6121)
				reg6123 := Call(__e, __defun__shen_4pair, reg6117, reg6122)
				reg6125 = reg6123
			} else {
				reg6124 := Call(__e, __defun__fail)
				reg6125 = reg6124
			}
			reg6127 = reg6125
		} else {
			reg6126 := Call(__e, __defun__fail)
			reg6127 = reg6126
		}
		YaccParse := reg6127
		_ = YaccParse
		reg6128 := Call(__e, __defun__fail)
		reg6129 := PrimEqual(YaccParse, reg6128)
		if reg6129 == True {
			reg6130 := PrimHead(V1702)
			reg6131 := PrimIsPair(reg6130)
			var reg6139 Obj
			if reg6131 == True {
				reg6132 := MakeSymbol("let")
				reg6133 := Call(__e, __defun__shen_4hdhd, V1702)
				reg6134 := PrimEqual(reg6132, reg6133)
				var reg6137 Obj
				if reg6134 == True {
					reg6135 := True
					reg6137 = reg6135
				} else {
					reg6136 := False
					reg6137 = reg6136
				}
				reg6139 = reg6137
			} else {
				reg6138 := False
				reg6139 = reg6138
			}
			if reg6139 == True {
				reg6140 := Call(__e, __defun__shen_4tlhd, V1702)
				reg6141 := Call(__e, __defun__shen_4hdtl, V1702)
				reg6142 := Call(__e, __defun__shen_4pair, reg6140, reg6141)
				NewStream1700 := reg6142
				_ = NewStream1700
				reg6143 := Call(__e, __defun__shen_4_5variable_2_6, NewStream1700)
				Parse__shen_4_5variable_2_6 := reg6143
				_ = Parse__shen_4_5variable_2_6
				reg6144 := Call(__e, __defun__fail)
				reg6145 := PrimEqual(reg6144, Parse__shen_4_5variable_2_6)
				reg6146 := PrimNot(reg6145)
				if reg6146 == True {
					reg6147 := Call(__e, __defun__shen_4_5expr_6, Parse__shen_4_5variable_2_6)
					Parse__shen_4_5expr_6 := reg6147
					_ = Parse__shen_4_5expr_6
					reg6148 := Call(__e, __defun__fail)
					reg6149 := PrimEqual(reg6148, Parse__shen_4_5expr_6)
					reg6150 := PrimNot(reg6149)
					if reg6150 == True {
						reg6151 := PrimHead(Parse__shen_4_5expr_6)
						reg6152 := MakeSymbol("let")
						reg6153 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5variable_2_6)
						reg6154 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5expr_6)
						reg6155 := Nil
						reg6156 := PrimCons(reg6154, reg6155)
						reg6157 := PrimCons(reg6153, reg6156)
						reg6158 := PrimCons(reg6152, reg6157)
						__e.TailApply(__defun__shen_4pair, reg6151, reg6158)
						return
					} else {
						__e.TailApply(__defun__fail)
						return
					}
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<side-condition>", value: __defun__shen_4_5side_1condition_6})

	__defun__shen_4_5variable_2_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1704 := __args[0]
		_ = V1704
		reg6163 := PrimHead(V1704)
		reg6164 := PrimIsPair(reg6163)
		if reg6164 == True {
			reg6165 := Call(__e, __defun__shen_4hdhd, V1704)
			Parse__X := reg6165
			_ = Parse__X
			reg6166 := PrimIsVariable(Parse__X)
			if reg6166 == True {
				reg6167 := Call(__e, __defun__shen_4tlhd, V1704)
				reg6168 := Call(__e, __defun__shen_4hdtl, V1704)
				reg6169 := Call(__e, __defun__shen_4pair, reg6167, reg6168)
				reg6170 := PrimHead(reg6169)
				__e.TailApply(__defun__shen_4pair, reg6170, Parse__X)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<variable?>", value: __defun__shen_4_5variable_2_6})

	__defun__shen_4_5expr_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1706 := __args[0]
		_ = V1706
		reg6174 := PrimHead(V1706)
		reg6175 := PrimIsPair(reg6174)
		if reg6175 == True {
			reg6176 := Call(__e, __defun__shen_4hdhd, V1706)
			Parse__X := reg6176
			_ = Parse__X
			reg6177 := MakeSymbol(">>")
			reg6178 := MakeSymbol(";")
			reg6179 := Nil
			reg6180 := PrimCons(reg6178, reg6179)
			reg6181 := PrimCons(reg6177, reg6180)
			reg6182 := Call(__e, __defun__element_2, Parse__X, reg6181)
			var reg6194 Obj
			if reg6182 == True {
				reg6183 := True
				reg6194 = reg6183
			} else {
				reg6184 := Call(__e, __defun__shen_4singleunderline_2, Parse__X)
				var reg6190 Obj
				if reg6184 == True {
					reg6185 := True
					reg6190 = reg6185
				} else {
					reg6186 := Call(__e, __defun__shen_4doubleunderline_2, Parse__X)
					var reg6189 Obj
					if reg6186 == True {
						reg6187 := True
						reg6189 = reg6187
					} else {
						reg6188 := False
						reg6189 = reg6188
					}
					reg6190 = reg6189
				}
				var reg6193 Obj
				if reg6190 == True {
					reg6191 := True
					reg6193 = reg6191
				} else {
					reg6192 := False
					reg6193 = reg6192
				}
				reg6194 = reg6193
			}
			reg6195 := PrimNot(reg6194)
			if reg6195 == True {
				reg6196 := Call(__e, __defun__shen_4tlhd, V1706)
				reg6197 := Call(__e, __defun__shen_4hdtl, V1706)
				reg6198 := Call(__e, __defun__shen_4pair, reg6196, reg6197)
				reg6199 := PrimHead(reg6198)
				reg6200 := Call(__e, __defun__shen_4remove_1bar, Parse__X)
				__e.TailApply(__defun__shen_4pair, reg6199, reg6200)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<expr>", value: __defun__shen_4_5expr_6})

	__defun__shen_4remove_1bar = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1708 := __args[0]
		_ = V1708
		reg6204 := PrimIsPair(V1708)
		var reg6238 Obj
		if reg6204 == True {
			reg6205 := PrimTail(V1708)
			reg6206 := PrimIsPair(reg6205)
			var reg6233 Obj
			if reg6206 == True {
				reg6207 := PrimTail(V1708)
				reg6208 := PrimTail(reg6207)
				reg6209 := PrimIsPair(reg6208)
				var reg6228 Obj
				if reg6209 == True {
					reg6210 := Nil
					reg6211 := PrimTail(V1708)
					reg6212 := PrimTail(reg6211)
					reg6213 := PrimTail(reg6212)
					reg6214 := PrimEqual(reg6210, reg6213)
					var reg6223 Obj
					if reg6214 == True {
						reg6215 := PrimTail(V1708)
						reg6216 := PrimHead(reg6215)
						reg6217 := MakeSymbol("bar!")
						reg6218 := PrimEqual(reg6216, reg6217)
						var reg6221 Obj
						if reg6218 == True {
							reg6219 := True
							reg6221 = reg6219
						} else {
							reg6220 := False
							reg6221 = reg6220
						}
						reg6223 = reg6221
					} else {
						reg6222 := False
						reg6223 = reg6222
					}
					var reg6226 Obj
					if reg6223 == True {
						reg6224 := True
						reg6226 = reg6224
					} else {
						reg6225 := False
						reg6226 = reg6225
					}
					reg6228 = reg6226
				} else {
					reg6227 := False
					reg6228 = reg6227
				}
				var reg6231 Obj
				if reg6228 == True {
					reg6229 := True
					reg6231 = reg6229
				} else {
					reg6230 := False
					reg6231 = reg6230
				}
				reg6233 = reg6231
			} else {
				reg6232 := False
				reg6233 = reg6232
			}
			var reg6236 Obj
			if reg6233 == True {
				reg6234 := True
				reg6236 = reg6234
			} else {
				reg6235 := False
				reg6236 = reg6235
			}
			reg6238 = reg6236
		} else {
			reg6237 := False
			reg6238 = reg6237
		}
		if reg6238 == True {
			reg6239 := PrimHead(V1708)
			reg6240 := PrimTail(V1708)
			reg6241 := PrimTail(reg6240)
			reg6242 := PrimHead(reg6241)
			reg6243 := PrimCons(reg6239, reg6242)
			__e.Return(reg6243)
			return
		} else {
			reg6244 := PrimIsPair(V1708)
			if reg6244 == True {
				reg6245 := PrimHead(V1708)
				reg6246 := Call(__e, __defun__shen_4remove_1bar, reg6245)
				reg6247 := PrimTail(V1708)
				reg6248 := Call(__e, __defun__shen_4remove_1bar, reg6247)
				reg6249 := PrimCons(reg6246, reg6248)
				__e.Return(reg6249)
				return
			} else {
				__e.Return(V1708)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.remove-bar", value: __defun__shen_4remove_1bar})

	__defun__shen_4_5premises_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1710 := __args[0]
		_ = V1710
		reg6250 := Call(__e, __defun__shen_4_5premise_6, V1710)
		Parse__shen_4_5premise_6 := reg6250
		_ = Parse__shen_4_5premise_6
		reg6251 := Call(__e, __defun__fail)
		reg6252 := PrimEqual(reg6251, Parse__shen_4_5premise_6)
		reg6253 := PrimNot(reg6252)
		var reg6272 Obj
		if reg6253 == True {
			reg6254 := Call(__e, __defun__shen_4_5semicolon_1symbol_6, Parse__shen_4_5premise_6)
			Parse__shen_4_5semicolon_1symbol_6 := reg6254
			_ = Parse__shen_4_5semicolon_1symbol_6
			reg6255 := Call(__e, __defun__fail)
			reg6256 := PrimEqual(reg6255, Parse__shen_4_5semicolon_1symbol_6)
			reg6257 := PrimNot(reg6256)
			var reg6270 Obj
			if reg6257 == True {
				reg6258 := Call(__e, __defun__shen_4_5premises_6, Parse__shen_4_5semicolon_1symbol_6)
				Parse__shen_4_5premises_6 := reg6258
				_ = Parse__shen_4_5premises_6
				reg6259 := Call(__e, __defun__fail)
				reg6260 := PrimEqual(reg6259, Parse__shen_4_5premises_6)
				reg6261 := PrimNot(reg6260)
				var reg6268 Obj
				if reg6261 == True {
					reg6262 := PrimHead(Parse__shen_4_5premises_6)
					reg6263 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5premise_6)
					reg6264 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5premises_6)
					reg6265 := PrimCons(reg6263, reg6264)
					reg6266 := Call(__e, __defun__shen_4pair, reg6262, reg6265)
					reg6268 = reg6266
				} else {
					reg6267 := Call(__e, __defun__fail)
					reg6268 = reg6267
				}
				reg6270 = reg6268
			} else {
				reg6269 := Call(__e, __defun__fail)
				reg6270 = reg6269
			}
			reg6272 = reg6270
		} else {
			reg6271 := Call(__e, __defun__fail)
			reg6272 = reg6271
		}
		YaccParse := reg6272
		_ = YaccParse
		reg6273 := Call(__e, __defun__fail)
		reg6274 := PrimEqual(YaccParse, reg6273)
		if reg6274 == True {
			reg6275 := Call(__e, __defun___5e_6, V1710)
			Parse___5e_6 := reg6275
			_ = Parse___5e_6
			reg6276 := Call(__e, __defun__fail)
			reg6277 := PrimEqual(reg6276, Parse___5e_6)
			reg6278 := PrimNot(reg6277)
			if reg6278 == True {
				reg6279 := PrimHead(Parse___5e_6)
				reg6280 := Nil
				__e.TailApply(__defun__shen_4pair, reg6279, reg6280)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<premises>", value: __defun__shen_4_5premises_6})

	__defun__shen_4_5semicolon_1symbol_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1712 := __args[0]
		_ = V1712
		reg6283 := PrimHead(V1712)
		reg6284 := PrimIsPair(reg6283)
		if reg6284 == True {
			reg6285 := Call(__e, __defun__shen_4hdhd, V1712)
			Parse__X := reg6285
			_ = Parse__X
			reg6286 := MakeSymbol(";")
			reg6287 := PrimEqual(Parse__X, reg6286)
			if reg6287 == True {
				reg6288 := Call(__e, __defun__shen_4tlhd, V1712)
				reg6289 := Call(__e, __defun__shen_4hdtl, V1712)
				reg6290 := Call(__e, __defun__shen_4pair, reg6288, reg6289)
				reg6291 := PrimHead(reg6290)
				reg6292 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg6291, reg6292)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<semicolon-symbol>", value: __defun__shen_4_5semicolon_1symbol_6})

	__defun__shen_4_5premise_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1716 := __args[0]
		_ = V1716
		reg6296 := PrimHead(V1716)
		reg6297 := PrimIsPair(reg6296)
		var reg6305 Obj
		if reg6297 == True {
			reg6298 := MakeSymbol("!")
			reg6299 := Call(__e, __defun__shen_4hdhd, V1716)
			reg6300 := PrimEqual(reg6298, reg6299)
			var reg6303 Obj
			if reg6300 == True {
				reg6301 := True
				reg6303 = reg6301
			} else {
				reg6302 := False
				reg6303 = reg6302
			}
			reg6305 = reg6303
		} else {
			reg6304 := False
			reg6305 = reg6304
		}
		var reg6313 Obj
		if reg6305 == True {
			reg6306 := Call(__e, __defun__shen_4tlhd, V1716)
			reg6307 := Call(__e, __defun__shen_4hdtl, V1716)
			reg6308 := Call(__e, __defun__shen_4pair, reg6306, reg6307)
			NewStream1713 := reg6308
			_ = NewStream1713
			reg6309 := PrimHead(NewStream1713)
			reg6310 := MakeSymbol("!")
			reg6311 := Call(__e, __defun__shen_4pair, reg6309, reg6310)
			reg6313 = reg6311
		} else {
			reg6312 := Call(__e, __defun__fail)
			reg6313 = reg6312
		}
		YaccParse := reg6313
		_ = YaccParse
		reg6314 := Call(__e, __defun__fail)
		reg6315 := PrimEqual(YaccParse, reg6314)
		if reg6315 == True {
			reg6316 := Call(__e, __defun__shen_4_5formulae_6, V1716)
			Parse__shen_4_5formulae_6 := reg6316
			_ = Parse__shen_4_5formulae_6
			reg6317 := Call(__e, __defun__fail)
			reg6318 := PrimEqual(reg6317, Parse__shen_4_5formulae_6)
			reg6319 := PrimNot(reg6318)
			var reg6347 Obj
			if reg6319 == True {
				reg6320 := PrimHead(Parse__shen_4_5formulae_6)
				reg6321 := PrimIsPair(reg6320)
				var reg6329 Obj
				if reg6321 == True {
					reg6322 := MakeSymbol(">>")
					reg6323 := Call(__e, __defun__shen_4hdhd, Parse__shen_4_5formulae_6)
					reg6324 := PrimEqual(reg6322, reg6323)
					var reg6327 Obj
					if reg6324 == True {
						reg6325 := True
						reg6327 = reg6325
					} else {
						reg6326 := False
						reg6327 = reg6326
					}
					reg6329 = reg6327
				} else {
					reg6328 := False
					reg6329 = reg6328
				}
				var reg6345 Obj
				if reg6329 == True {
					reg6330 := Call(__e, __defun__shen_4tlhd, Parse__shen_4_5formulae_6)
					reg6331 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formulae_6)
					reg6332 := Call(__e, __defun__shen_4pair, reg6330, reg6331)
					NewStream1714 := reg6332
					_ = NewStream1714
					reg6333 := Call(__e, __defun__shen_4_5formula_6, NewStream1714)
					Parse__shen_4_5formula_6 := reg6333
					_ = Parse__shen_4_5formula_6
					reg6334 := Call(__e, __defun__fail)
					reg6335 := PrimEqual(reg6334, Parse__shen_4_5formula_6)
					reg6336 := PrimNot(reg6335)
					var reg6343 Obj
					if reg6336 == True {
						reg6337 := PrimHead(Parse__shen_4_5formula_6)
						reg6338 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formulae_6)
						reg6339 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formula_6)
						reg6340 := Call(__e, __defun__shen_4sequent, reg6338, reg6339)
						reg6341 := Call(__e, __defun__shen_4pair, reg6337, reg6340)
						reg6343 = reg6341
					} else {
						reg6342 := Call(__e, __defun__fail)
						reg6343 = reg6342
					}
					reg6345 = reg6343
				} else {
					reg6344 := Call(__e, __defun__fail)
					reg6345 = reg6344
				}
				reg6347 = reg6345
			} else {
				reg6346 := Call(__e, __defun__fail)
				reg6347 = reg6346
			}
			YaccParse := reg6347
			_ = YaccParse
			reg6348 := Call(__e, __defun__fail)
			reg6349 := PrimEqual(YaccParse, reg6348)
			if reg6349 == True {
				reg6350 := Call(__e, __defun__shen_4_5formula_6, V1716)
				Parse__shen_4_5formula_6 := reg6350
				_ = Parse__shen_4_5formula_6
				reg6351 := Call(__e, __defun__fail)
				reg6352 := PrimEqual(reg6351, Parse__shen_4_5formula_6)
				reg6353 := PrimNot(reg6352)
				if reg6353 == True {
					reg6354 := PrimHead(Parse__shen_4_5formula_6)
					reg6355 := Nil
					reg6356 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formula_6)
					reg6357 := Call(__e, __defun__shen_4sequent, reg6355, reg6356)
					__e.TailApply(__defun__shen_4pair, reg6354, reg6357)
					return
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.Return(YaccParse)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<premise>", value: __defun__shen_4_5premise_6})

	__defun__shen_4_5conclusion_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1719 := __args[0]
		_ = V1719
		reg6360 := Call(__e, __defun__shen_4_5formulae_6, V1719)
		Parse__shen_4_5formulae_6 := reg6360
		_ = Parse__shen_4_5formulae_6
		reg6361 := Call(__e, __defun__fail)
		reg6362 := PrimEqual(reg6361, Parse__shen_4_5formulae_6)
		reg6363 := PrimNot(reg6362)
		var reg6397 Obj
		if reg6363 == True {
			reg6364 := PrimHead(Parse__shen_4_5formulae_6)
			reg6365 := PrimIsPair(reg6364)
			var reg6373 Obj
			if reg6365 == True {
				reg6366 := MakeSymbol(">>")
				reg6367 := Call(__e, __defun__shen_4hdhd, Parse__shen_4_5formulae_6)
				reg6368 := PrimEqual(reg6366, reg6367)
				var reg6371 Obj
				if reg6368 == True {
					reg6369 := True
					reg6371 = reg6369
				} else {
					reg6370 := False
					reg6371 = reg6370
				}
				reg6373 = reg6371
			} else {
				reg6372 := False
				reg6373 = reg6372
			}
			var reg6395 Obj
			if reg6373 == True {
				reg6374 := Call(__e, __defun__shen_4tlhd, Parse__shen_4_5formulae_6)
				reg6375 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formulae_6)
				reg6376 := Call(__e, __defun__shen_4pair, reg6374, reg6375)
				NewStream1717 := reg6376
				_ = NewStream1717
				reg6377 := Call(__e, __defun__shen_4_5formula_6, NewStream1717)
				Parse__shen_4_5formula_6 := reg6377
				_ = Parse__shen_4_5formula_6
				reg6378 := Call(__e, __defun__fail)
				reg6379 := PrimEqual(reg6378, Parse__shen_4_5formula_6)
				reg6380 := PrimNot(reg6379)
				var reg6393 Obj
				if reg6380 == True {
					reg6381 := Call(__e, __defun__shen_4_5semicolon_1symbol_6, Parse__shen_4_5formula_6)
					Parse__shen_4_5semicolon_1symbol_6 := reg6381
					_ = Parse__shen_4_5semicolon_1symbol_6
					reg6382 := Call(__e, __defun__fail)
					reg6383 := PrimEqual(reg6382, Parse__shen_4_5semicolon_1symbol_6)
					reg6384 := PrimNot(reg6383)
					var reg6391 Obj
					if reg6384 == True {
						reg6385 := PrimHead(Parse__shen_4_5semicolon_1symbol_6)
						reg6386 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formulae_6)
						reg6387 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formula_6)
						reg6388 := Call(__e, __defun__shen_4sequent, reg6386, reg6387)
						reg6389 := Call(__e, __defun__shen_4pair, reg6385, reg6388)
						reg6391 = reg6389
					} else {
						reg6390 := Call(__e, __defun__fail)
						reg6391 = reg6390
					}
					reg6393 = reg6391
				} else {
					reg6392 := Call(__e, __defun__fail)
					reg6393 = reg6392
				}
				reg6395 = reg6393
			} else {
				reg6394 := Call(__e, __defun__fail)
				reg6395 = reg6394
			}
			reg6397 = reg6395
		} else {
			reg6396 := Call(__e, __defun__fail)
			reg6397 = reg6396
		}
		YaccParse := reg6397
		_ = YaccParse
		reg6398 := Call(__e, __defun__fail)
		reg6399 := PrimEqual(YaccParse, reg6398)
		if reg6399 == True {
			reg6400 := Call(__e, __defun__shen_4_5formula_6, V1719)
			Parse__shen_4_5formula_6 := reg6400
			_ = Parse__shen_4_5formula_6
			reg6401 := Call(__e, __defun__fail)
			reg6402 := PrimEqual(reg6401, Parse__shen_4_5formula_6)
			reg6403 := PrimNot(reg6402)
			if reg6403 == True {
				reg6404 := Call(__e, __defun__shen_4_5semicolon_1symbol_6, Parse__shen_4_5formula_6)
				Parse__shen_4_5semicolon_1symbol_6 := reg6404
				_ = Parse__shen_4_5semicolon_1symbol_6
				reg6405 := Call(__e, __defun__fail)
				reg6406 := PrimEqual(reg6405, Parse__shen_4_5semicolon_1symbol_6)
				reg6407 := PrimNot(reg6406)
				if reg6407 == True {
					reg6408 := PrimHead(Parse__shen_4_5semicolon_1symbol_6)
					reg6409 := Nil
					reg6410 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formula_6)
					reg6411 := Call(__e, __defun__shen_4sequent, reg6409, reg6410)
					__e.TailApply(__defun__shen_4pair, reg6408, reg6411)
					return
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<conclusion>", value: __defun__shen_4_5conclusion_6})

	__defun__shen_4sequent = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1722 := __args[0]
		_ = V1722
		V1723 := __args[1]
		_ = V1723
		__e.TailApply(__defun___8p, V1722, V1723)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.sequent", value: __defun__shen_4sequent})

	__defun__shen_4_5formulae_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1725 := __args[0]
		_ = V1725
		reg6416 := Call(__e, __defun__shen_4_5formula_6, V1725)
		Parse__shen_4_5formula_6 := reg6416
		_ = Parse__shen_4_5formula_6
		reg6417 := Call(__e, __defun__fail)
		reg6418 := PrimEqual(reg6417, Parse__shen_4_5formula_6)
		reg6419 := PrimNot(reg6418)
		var reg6438 Obj
		if reg6419 == True {
			reg6420 := Call(__e, __defun__shen_4_5comma_1symbol_6, Parse__shen_4_5formula_6)
			Parse__shen_4_5comma_1symbol_6 := reg6420
			_ = Parse__shen_4_5comma_1symbol_6
			reg6421 := Call(__e, __defun__fail)
			reg6422 := PrimEqual(reg6421, Parse__shen_4_5comma_1symbol_6)
			reg6423 := PrimNot(reg6422)
			var reg6436 Obj
			if reg6423 == True {
				reg6424 := Call(__e, __defun__shen_4_5formulae_6, Parse__shen_4_5comma_1symbol_6)
				Parse__shen_4_5formulae_6 := reg6424
				_ = Parse__shen_4_5formulae_6
				reg6425 := Call(__e, __defun__fail)
				reg6426 := PrimEqual(reg6425, Parse__shen_4_5formulae_6)
				reg6427 := PrimNot(reg6426)
				var reg6434 Obj
				if reg6427 == True {
					reg6428 := PrimHead(Parse__shen_4_5formulae_6)
					reg6429 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formula_6)
					reg6430 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formulae_6)
					reg6431 := PrimCons(reg6429, reg6430)
					reg6432 := Call(__e, __defun__shen_4pair, reg6428, reg6431)
					reg6434 = reg6432
				} else {
					reg6433 := Call(__e, __defun__fail)
					reg6434 = reg6433
				}
				reg6436 = reg6434
			} else {
				reg6435 := Call(__e, __defun__fail)
				reg6436 = reg6435
			}
			reg6438 = reg6436
		} else {
			reg6437 := Call(__e, __defun__fail)
			reg6438 = reg6437
		}
		YaccParse := reg6438
		_ = YaccParse
		reg6439 := Call(__e, __defun__fail)
		reg6440 := PrimEqual(YaccParse, reg6439)
		if reg6440 == True {
			reg6441 := Call(__e, __defun__shen_4_5formula_6, V1725)
			Parse__shen_4_5formula_6 := reg6441
			_ = Parse__shen_4_5formula_6
			reg6442 := Call(__e, __defun__fail)
			reg6443 := PrimEqual(reg6442, Parse__shen_4_5formula_6)
			reg6444 := PrimNot(reg6443)
			var reg6451 Obj
			if reg6444 == True {
				reg6445 := PrimHead(Parse__shen_4_5formula_6)
				reg6446 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5formula_6)
				reg6447 := Nil
				reg6448 := PrimCons(reg6446, reg6447)
				reg6449 := Call(__e, __defun__shen_4pair, reg6445, reg6448)
				reg6451 = reg6449
			} else {
				reg6450 := Call(__e, __defun__fail)
				reg6451 = reg6450
			}
			YaccParse := reg6451
			_ = YaccParse
			reg6452 := Call(__e, __defun__fail)
			reg6453 := PrimEqual(YaccParse, reg6452)
			if reg6453 == True {
				reg6454 := Call(__e, __defun___5e_6, V1725)
				Parse___5e_6 := reg6454
				_ = Parse___5e_6
				reg6455 := Call(__e, __defun__fail)
				reg6456 := PrimEqual(reg6455, Parse___5e_6)
				reg6457 := PrimNot(reg6456)
				if reg6457 == True {
					reg6458 := PrimHead(Parse___5e_6)
					reg6459 := Nil
					__e.TailApply(__defun__shen_4pair, reg6458, reg6459)
					return
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.Return(YaccParse)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<formulae>", value: __defun__shen_4_5formulae_6})

	__defun__shen_4_5comma_1symbol_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1727 := __args[0]
		_ = V1727
		reg6462 := PrimHead(V1727)
		reg6463 := PrimIsPair(reg6462)
		if reg6463 == True {
			reg6464 := Call(__e, __defun__shen_4hdhd, V1727)
			Parse__X := reg6464
			_ = Parse__X
			reg6465 := MakeString(",")
			reg6466 := PrimIntern(reg6465)
			reg6467 := PrimEqual(Parse__X, reg6466)
			if reg6467 == True {
				reg6468 := Call(__e, __defun__shen_4tlhd, V1727)
				reg6469 := Call(__e, __defun__shen_4hdtl, V1727)
				reg6470 := Call(__e, __defun__shen_4pair, reg6468, reg6469)
				reg6471 := PrimHead(reg6470)
				reg6472 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg6471, reg6472)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<comma-symbol>", value: __defun__shen_4_5comma_1symbol_6})

	__defun__shen_4_5formula_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1730 := __args[0]
		_ = V1730
		reg6476 := Call(__e, __defun__shen_4_5expr_6, V1730)
		Parse__shen_4_5expr_6 := reg6476
		_ = Parse__shen_4_5expr_6
		reg6477 := Call(__e, __defun__fail)
		reg6478 := PrimEqual(reg6477, Parse__shen_4_5expr_6)
		reg6479 := PrimNot(reg6478)
		var reg6513 Obj
		if reg6479 == True {
			reg6480 := PrimHead(Parse__shen_4_5expr_6)
			reg6481 := PrimIsPair(reg6480)
			var reg6489 Obj
			if reg6481 == True {
				reg6482 := MakeSymbol(":")
				reg6483 := Call(__e, __defun__shen_4hdhd, Parse__shen_4_5expr_6)
				reg6484 := PrimEqual(reg6482, reg6483)
				var reg6487 Obj
				if reg6484 == True {
					reg6485 := True
					reg6487 = reg6485
				} else {
					reg6486 := False
					reg6487 = reg6486
				}
				reg6489 = reg6487
			} else {
				reg6488 := False
				reg6489 = reg6488
			}
			var reg6511 Obj
			if reg6489 == True {
				reg6490 := Call(__e, __defun__shen_4tlhd, Parse__shen_4_5expr_6)
				reg6491 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5expr_6)
				reg6492 := Call(__e, __defun__shen_4pair, reg6490, reg6491)
				NewStream1728 := reg6492
				_ = NewStream1728
				reg6493 := Call(__e, __defun__shen_4_5type_6, NewStream1728)
				Parse__shen_4_5type_6 := reg6493
				_ = Parse__shen_4_5type_6
				reg6494 := Call(__e, __defun__fail)
				reg6495 := PrimEqual(reg6494, Parse__shen_4_5type_6)
				reg6496 := PrimNot(reg6495)
				var reg6509 Obj
				if reg6496 == True {
					reg6497 := PrimHead(Parse__shen_4_5type_6)
					reg6498 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5expr_6)
					reg6499 := Call(__e, __defun__shen_4curry, reg6498)
					reg6500 := MakeSymbol(":")
					reg6501 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5type_6)
					reg6502 := Call(__e, __defun__shen_4demodulate, reg6501)
					reg6503 := Nil
					reg6504 := PrimCons(reg6502, reg6503)
					reg6505 := PrimCons(reg6500, reg6504)
					reg6506 := PrimCons(reg6499, reg6505)
					reg6507 := Call(__e, __defun__shen_4pair, reg6497, reg6506)
					reg6509 = reg6507
				} else {
					reg6508 := Call(__e, __defun__fail)
					reg6509 = reg6508
				}
				reg6511 = reg6509
			} else {
				reg6510 := Call(__e, __defun__fail)
				reg6511 = reg6510
			}
			reg6513 = reg6511
		} else {
			reg6512 := Call(__e, __defun__fail)
			reg6513 = reg6512
		}
		YaccParse := reg6513
		_ = YaccParse
		reg6514 := Call(__e, __defun__fail)
		reg6515 := PrimEqual(YaccParse, reg6514)
		if reg6515 == True {
			reg6516 := Call(__e, __defun__shen_4_5expr_6, V1730)
			Parse__shen_4_5expr_6 := reg6516
			_ = Parse__shen_4_5expr_6
			reg6517 := Call(__e, __defun__fail)
			reg6518 := PrimEqual(reg6517, Parse__shen_4_5expr_6)
			reg6519 := PrimNot(reg6518)
			if reg6519 == True {
				reg6520 := PrimHead(Parse__shen_4_5expr_6)
				reg6521 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5expr_6)
				__e.TailApply(__defun__shen_4pair, reg6520, reg6521)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<formula>", value: __defun__shen_4_5formula_6})

	__defun__shen_4_5type_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1732 := __args[0]
		_ = V1732
		reg6524 := Call(__e, __defun__shen_4_5expr_6, V1732)
		Parse__shen_4_5expr_6 := reg6524
		_ = Parse__shen_4_5expr_6
		reg6525 := Call(__e, __defun__fail)
		reg6526 := PrimEqual(reg6525, Parse__shen_4_5expr_6)
		reg6527 := PrimNot(reg6526)
		if reg6527 == True {
			reg6528 := PrimHead(Parse__shen_4_5expr_6)
			reg6529 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5expr_6)
			reg6530 := Call(__e, __defun__shen_4curry_1type, reg6529)
			__e.TailApply(__defun__shen_4pair, reg6528, reg6530)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<type>", value: __defun__shen_4_5type_6})

	__defun__shen_4_5doubleunderline_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1734 := __args[0]
		_ = V1734
		reg6533 := PrimHead(V1734)
		reg6534 := PrimIsPair(reg6533)
		if reg6534 == True {
			reg6535 := Call(__e, __defun__shen_4hdhd, V1734)
			Parse__X := reg6535
			_ = Parse__X
			reg6536 := Call(__e, __defun__shen_4doubleunderline_2, Parse__X)
			if reg6536 == True {
				reg6537 := Call(__e, __defun__shen_4tlhd, V1734)
				reg6538 := Call(__e, __defun__shen_4hdtl, V1734)
				reg6539 := Call(__e, __defun__shen_4pair, reg6537, reg6538)
				reg6540 := PrimHead(reg6539)
				__e.TailApply(__defun__shen_4pair, reg6540, Parse__X)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<doubleunderline>", value: __defun__shen_4_5doubleunderline_6})

	__defun__shen_4_5singleunderline_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1736 := __args[0]
		_ = V1736
		reg6544 := PrimHead(V1736)
		reg6545 := PrimIsPair(reg6544)
		if reg6545 == True {
			reg6546 := Call(__e, __defun__shen_4hdhd, V1736)
			Parse__X := reg6546
			_ = Parse__X
			reg6547 := Call(__e, __defun__shen_4singleunderline_2, Parse__X)
			if reg6547 == True {
				reg6548 := Call(__e, __defun__shen_4tlhd, V1736)
				reg6549 := Call(__e, __defun__shen_4hdtl, V1736)
				reg6550 := Call(__e, __defun__shen_4pair, reg6548, reg6549)
				reg6551 := PrimHead(reg6550)
				__e.TailApply(__defun__shen_4pair, reg6551, Parse__X)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<singleunderline>", value: __defun__shen_4_5singleunderline_6})

	__defun__shen_4singleunderline_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1738 := __args[0]
		_ = V1738
		reg6555 := PrimIsSymbol(V1738)
		if reg6555 == True {
			reg6556 := PrimStr(V1738)
			reg6557 := Call(__e, __defun__shen_4sh_2, reg6556)
			if reg6557 == True {
				reg6558 := True
				__e.Return(reg6558)
				return
			} else {
				reg6559 := False
				__e.Return(reg6559)
				return
			}
		} else {
			reg6560 := False
			__e.Return(reg6560)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.singleunderline?", value: __defun__shen_4singleunderline_2})

	__defun__shen_4sh_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1740 := __args[0]
		_ = V1740
		reg6561 := MakeString("_")
		reg6562 := PrimEqual(reg6561, V1740)
		if reg6562 == True {
			reg6563 := True
			__e.Return(reg6563)
			return
		} else {
			reg6564 := MakeNumber(0)
			reg6565 := PrimPos(V1740, reg6564)
			reg6566 := MakeString("_")
			reg6567 := PrimEqual(reg6565, reg6566)
			if reg6567 == True {
				reg6568 := PrimTailString(V1740)
				reg6569 := Call(__e, __defun__shen_4sh_2, reg6568)
				if reg6569 == True {
					reg6570 := True
					__e.Return(reg6570)
					return
				} else {
					reg6571 := False
					__e.Return(reg6571)
					return
				}
			} else {
				reg6572 := False
				__e.Return(reg6572)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.sh?", value: __defun__shen_4sh_2})

	__defun__shen_4doubleunderline_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1742 := __args[0]
		_ = V1742
		reg6573 := PrimIsSymbol(V1742)
		if reg6573 == True {
			reg6574 := PrimStr(V1742)
			reg6575 := Call(__e, __defun__shen_4dh_2, reg6574)
			if reg6575 == True {
				reg6576 := True
				__e.Return(reg6576)
				return
			} else {
				reg6577 := False
				__e.Return(reg6577)
				return
			}
		} else {
			reg6578 := False
			__e.Return(reg6578)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.doubleunderline?", value: __defun__shen_4doubleunderline_2})

	__defun__shen_4dh_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1744 := __args[0]
		_ = V1744
		reg6579 := MakeString("=")
		reg6580 := PrimEqual(reg6579, V1744)
		if reg6580 == True {
			reg6581 := True
			__e.Return(reg6581)
			return
		} else {
			reg6582 := MakeNumber(0)
			reg6583 := PrimPos(V1744, reg6582)
			reg6584 := MakeString("=")
			reg6585 := PrimEqual(reg6583, reg6584)
			if reg6585 == True {
				reg6586 := PrimTailString(V1744)
				reg6587 := Call(__e, __defun__shen_4dh_2, reg6586)
				if reg6587 == True {
					reg6588 := True
					__e.Return(reg6588)
					return
				} else {
					reg6589 := False
					__e.Return(reg6589)
					return
				}
			} else {
				reg6590 := False
				__e.Return(reg6590)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dh?", value: __defun__shen_4dh_2})

	__defun__shen_4process_1datatype = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1747 := __args[0]
		_ = V1747
		V1748 := __args[1]
		_ = V1748
		reg6591 := Call(__e, __defun__shen_4rules_1_6horn_1clauses, V1747, V1748)
		reg6592 := Call(__e, __defun__shen_4s_1prolog, reg6591)
		__e.TailApply(__defun__shen_4remember_1datatype, reg6592)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.process-datatype", value: __defun__shen_4process_1datatype})

	__defun__shen_4remember_1datatype = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1754 := __args[0]
		_ = V1754
		reg6594 := PrimIsPair(V1754)
		if reg6594 == True {
			reg6595 := MakeSymbol("shen.*datatypes*")
			reg6596 := PrimHead(V1754)
			reg6597 := MakeSymbol("shen.*datatypes*")
			reg6598 := PrimValue(reg6597)
			reg6599 := Call(__e, __defun__adjoin, reg6596, reg6598)
			reg6600 := PrimSet(reg6595, reg6599)
			_ = reg6600
			reg6601 := MakeSymbol("shen.*alldatatypes*")
			reg6602 := PrimHead(V1754)
			reg6603 := MakeSymbol("shen.*alldatatypes*")
			reg6604 := PrimValue(reg6603)
			reg6605 := Call(__e, __defun__adjoin, reg6602, reg6604)
			reg6606 := PrimSet(reg6601, reg6605)
			_ = reg6606
			reg6607 := PrimHead(V1754)
			__e.Return(reg6607)
			return
		} else {
			reg6608 := MakeSymbol("shen.remember-datatype")
			__e.TailApply(__defun__shen_4f__error, reg6608)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.remember-datatype", value: __defun__shen_4remember_1datatype})

	__defun__shen_4rules_1_6horn_1clauses = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1759 := __args[0]
		_ = V1759
		V1760 := __args[1]
		_ = V1760
		reg6610 := Nil
		reg6611 := PrimEqual(reg6610, V1760)
		if reg6611 == True {
			reg6612 := Nil
			__e.Return(reg6612)
			return
		} else {
			reg6613 := PrimIsPair(V1760)
			var reg6629 Obj
			if reg6613 == True {
				reg6614 := PrimHead(V1760)
				reg6615 := Call(__e, __defun__tuple_2, reg6614)
				var reg6624 Obj
				if reg6615 == True {
					reg6616 := MakeSymbol("shen.single")
					reg6617 := PrimHead(V1760)
					reg6618 := Call(__e, __defun__fst, reg6617)
					reg6619 := PrimEqual(reg6616, reg6618)
					var reg6622 Obj
					if reg6619 == True {
						reg6620 := True
						reg6622 = reg6620
					} else {
						reg6621 := False
						reg6622 = reg6621
					}
					reg6624 = reg6622
				} else {
					reg6623 := False
					reg6624 = reg6623
				}
				var reg6627 Obj
				if reg6624 == True {
					reg6625 := True
					reg6627 = reg6625
				} else {
					reg6626 := False
					reg6627 = reg6626
				}
				reg6629 = reg6627
			} else {
				reg6628 := False
				reg6629 = reg6628
			}
			if reg6629 == True {
				reg6630 := PrimHead(V1760)
				reg6631 := Call(__e, __defun__snd, reg6630)
				reg6632 := Call(__e, __defun__shen_4rule_1_6horn_1clause, V1759, reg6631)
				reg6633 := PrimTail(V1760)
				reg6634 := Call(__e, __defun__shen_4rules_1_6horn_1clauses, V1759, reg6633)
				reg6635 := PrimCons(reg6632, reg6634)
				__e.Return(reg6635)
				return
			} else {
				reg6636 := PrimIsPair(V1760)
				var reg6652 Obj
				if reg6636 == True {
					reg6637 := PrimHead(V1760)
					reg6638 := Call(__e, __defun__tuple_2, reg6637)
					var reg6647 Obj
					if reg6638 == True {
						reg6639 := MakeSymbol("shen.double")
						reg6640 := PrimHead(V1760)
						reg6641 := Call(__e, __defun__fst, reg6640)
						reg6642 := PrimEqual(reg6639, reg6641)
						var reg6645 Obj
						if reg6642 == True {
							reg6643 := True
							reg6645 = reg6643
						} else {
							reg6644 := False
							reg6645 = reg6644
						}
						reg6647 = reg6645
					} else {
						reg6646 := False
						reg6647 = reg6646
					}
					var reg6650 Obj
					if reg6647 == True {
						reg6648 := True
						reg6650 = reg6648
					} else {
						reg6649 := False
						reg6650 = reg6649
					}
					reg6652 = reg6650
				} else {
					reg6651 := False
					reg6652 = reg6651
				}
				if reg6652 == True {
					reg6653 := PrimHead(V1760)
					reg6654 := Call(__e, __defun__snd, reg6653)
					reg6655 := Call(__e, __defun__shen_4double_1_6singles, reg6654)
					reg6656 := PrimTail(V1760)
					reg6657 := Call(__e, __defun__append, reg6655, reg6656)
					__e.TailApply(__defun__shen_4rules_1_6horn_1clauses, V1759, reg6657)
					return
				} else {
					reg6659 := MakeSymbol("shen.rules->horn-clauses")
					__e.TailApply(__defun__shen_4f__error, reg6659)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.rules->horn-clauses", value: __defun__shen_4rules_1_6horn_1clauses})

	__defun__shen_4double_1_6singles = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1762 := __args[0]
		_ = V1762
		reg6661 := Call(__e, __defun__shen_4right_1rule, V1762)
		reg6662 := Call(__e, __defun__shen_4left_1rule, V1762)
		reg6663 := Nil
		reg6664 := PrimCons(reg6662, reg6663)
		reg6665 := PrimCons(reg6661, reg6664)
		__e.Return(reg6665)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.double->singles", value: __defun__shen_4double_1_6singles})

	__defun__shen_4right_1rule = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1764 := __args[0]
		_ = V1764
		reg6666 := MakeSymbol("shen.single")
		__e.TailApply(__defun___8p, reg6666, V1764)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.right-rule", value: __defun__shen_4right_1rule})

	__defun__shen_4left_1rule = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1766 := __args[0]
		_ = V1766
		reg6668 := PrimIsPair(V1766)
		var reg6713 Obj
		if reg6668 == True {
			reg6669 := PrimTail(V1766)
			reg6670 := PrimIsPair(reg6669)
			var reg6708 Obj
			if reg6670 == True {
				reg6671 := PrimTail(V1766)
				reg6672 := PrimTail(reg6671)
				reg6673 := PrimIsPair(reg6672)
				var reg6703 Obj
				if reg6673 == True {
					reg6674 := PrimTail(V1766)
					reg6675 := PrimTail(reg6674)
					reg6676 := PrimHead(reg6675)
					reg6677 := Call(__e, __defun__tuple_2, reg6676)
					var reg6698 Obj
					if reg6677 == True {
						reg6678 := Nil
						reg6679 := PrimTail(V1766)
						reg6680 := PrimTail(reg6679)
						reg6681 := PrimHead(reg6680)
						reg6682 := Call(__e, __defun__fst, reg6681)
						reg6683 := PrimEqual(reg6678, reg6682)
						var reg6693 Obj
						if reg6683 == True {
							reg6684 := Nil
							reg6685 := PrimTail(V1766)
							reg6686 := PrimTail(reg6685)
							reg6687 := PrimTail(reg6686)
							reg6688 := PrimEqual(reg6684, reg6687)
							var reg6691 Obj
							if reg6688 == True {
								reg6689 := True
								reg6691 = reg6689
							} else {
								reg6690 := False
								reg6691 = reg6690
							}
							reg6693 = reg6691
						} else {
							reg6692 := False
							reg6693 = reg6692
						}
						var reg6696 Obj
						if reg6693 == True {
							reg6694 := True
							reg6696 = reg6694
						} else {
							reg6695 := False
							reg6696 = reg6695
						}
						reg6698 = reg6696
					} else {
						reg6697 := False
						reg6698 = reg6697
					}
					var reg6701 Obj
					if reg6698 == True {
						reg6699 := True
						reg6701 = reg6699
					} else {
						reg6700 := False
						reg6701 = reg6700
					}
					reg6703 = reg6701
				} else {
					reg6702 := False
					reg6703 = reg6702
				}
				var reg6706 Obj
				if reg6703 == True {
					reg6704 := True
					reg6706 = reg6704
				} else {
					reg6705 := False
					reg6706 = reg6705
				}
				reg6708 = reg6706
			} else {
				reg6707 := False
				reg6708 = reg6707
			}
			var reg6711 Obj
			if reg6708 == True {
				reg6709 := True
				reg6711 = reg6709
			} else {
				reg6710 := False
				reg6711 = reg6710
			}
			reg6713 = reg6711
		} else {
			reg6712 := False
			reg6713 = reg6712
		}
		if reg6713 == True {
			reg6714 := MakeSymbol("Qv")
			reg6715 := Call(__e, __defun__gensym, reg6714)
			Q := reg6715
			_ = Q
			reg6716 := PrimTail(V1766)
			reg6717 := PrimTail(reg6716)
			reg6718 := PrimHead(reg6717)
			reg6719 := Call(__e, __defun__snd, reg6718)
			reg6720 := Nil
			reg6721 := PrimCons(reg6719, reg6720)
			reg6722 := Call(__e, __defun___8p, reg6721, Q)
			NewConclusion := reg6722
			_ = NewConclusion
			reg6723 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(__defun__shen_4right_1_6left, X)
				return
			}, 1)
			reg6725 := PrimTail(V1766)
			reg6726 := PrimHead(reg6725)
			reg6727 := Call(__e, __defun__map, reg6723, reg6726)
			reg6728 := Call(__e, __defun___8p, reg6727, Q)
			reg6729 := Nil
			reg6730 := PrimCons(reg6728, reg6729)
			NewPremises := reg6730
			_ = NewPremises
			reg6731 := MakeSymbol("shen.single")
			reg6732 := PrimHead(V1766)
			reg6733 := Nil
			reg6734 := PrimCons(NewConclusion, reg6733)
			reg6735 := PrimCons(NewPremises, reg6734)
			reg6736 := PrimCons(reg6732, reg6735)
			__e.TailApply(__defun___8p, reg6731, reg6736)
			return
		} else {
			reg6738 := MakeSymbol("shen.left-rule")
			__e.TailApply(__defun__shen_4f__error, reg6738)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.left-rule", value: __defun__shen_4left_1rule})

	__defun__shen_4right_1_6left = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1772 := __args[0]
		_ = V1772
		reg6740 := Call(__e, __defun__tuple_2, V1772)
		var reg6748 Obj
		if reg6740 == True {
			reg6741 := Nil
			reg6742 := Call(__e, __defun__fst, V1772)
			reg6743 := PrimEqual(reg6741, reg6742)
			var reg6746 Obj
			if reg6743 == True {
				reg6744 := True
				reg6746 = reg6744
			} else {
				reg6745 := False
				reg6746 = reg6745
			}
			reg6748 = reg6746
		} else {
			reg6747 := False
			reg6748 = reg6747
		}
		if reg6748 == True {
			__e.TailApply(__defun__snd, V1772)
			return
		} else {
			reg6750 := MakeString("syntax error with ==========\n")
			reg6751 := PrimSimpleError(reg6750)
			__e.Return(reg6751)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.right->left", value: __defun__shen_4right_1_6left})

	__defun__shen_4rule_1_6horn_1clause = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1775 := __args[0]
		_ = V1775
		V1776 := __args[1]
		_ = V1776
		reg6752 := PrimIsPair(V1776)
		var reg6786 Obj
		if reg6752 == True {
			reg6753 := PrimTail(V1776)
			reg6754 := PrimIsPair(reg6753)
			var reg6781 Obj
			if reg6754 == True {
				reg6755 := PrimTail(V1776)
				reg6756 := PrimTail(reg6755)
				reg6757 := PrimIsPair(reg6756)
				var reg6776 Obj
				if reg6757 == True {
					reg6758 := PrimTail(V1776)
					reg6759 := PrimTail(reg6758)
					reg6760 := PrimHead(reg6759)
					reg6761 := Call(__e, __defun__tuple_2, reg6760)
					var reg6771 Obj
					if reg6761 == True {
						reg6762 := Nil
						reg6763 := PrimTail(V1776)
						reg6764 := PrimTail(reg6763)
						reg6765 := PrimTail(reg6764)
						reg6766 := PrimEqual(reg6762, reg6765)
						var reg6769 Obj
						if reg6766 == True {
							reg6767 := True
							reg6769 = reg6767
						} else {
							reg6768 := False
							reg6769 = reg6768
						}
						reg6771 = reg6769
					} else {
						reg6770 := False
						reg6771 = reg6770
					}
					var reg6774 Obj
					if reg6771 == True {
						reg6772 := True
						reg6774 = reg6772
					} else {
						reg6773 := False
						reg6774 = reg6773
					}
					reg6776 = reg6774
				} else {
					reg6775 := False
					reg6776 = reg6775
				}
				var reg6779 Obj
				if reg6776 == True {
					reg6777 := True
					reg6779 = reg6777
				} else {
					reg6778 := False
					reg6779 = reg6778
				}
				reg6781 = reg6779
			} else {
				reg6780 := False
				reg6781 = reg6780
			}
			var reg6784 Obj
			if reg6781 == True {
				reg6782 := True
				reg6784 = reg6782
			} else {
				reg6783 := False
				reg6784 = reg6783
			}
			reg6786 = reg6784
		} else {
			reg6785 := False
			reg6786 = reg6785
		}
		if reg6786 == True {
			reg6787 := PrimTail(V1776)
			reg6788 := PrimTail(reg6787)
			reg6789 := PrimHead(reg6788)
			reg6790 := Call(__e, __defun__snd, reg6789)
			reg6791 := Call(__e, __defun__shen_4rule_1_6horn_1clause_1head, V1775, reg6790)
			reg6792 := MakeSymbol(":-")
			reg6793 := PrimHead(V1776)
			reg6794 := PrimTail(V1776)
			reg6795 := PrimHead(reg6794)
			reg6796 := PrimTail(V1776)
			reg6797 := PrimTail(reg6796)
			reg6798 := PrimHead(reg6797)
			reg6799 := Call(__e, __defun__fst, reg6798)
			reg6800 := Call(__e, __defun__shen_4rule_1_6horn_1clause_1body, reg6793, reg6795, reg6799)
			reg6801 := Nil
			reg6802 := PrimCons(reg6800, reg6801)
			reg6803 := PrimCons(reg6792, reg6802)
			reg6804 := PrimCons(reg6791, reg6803)
			__e.Return(reg6804)
			return
		} else {
			reg6805 := MakeSymbol("shen.rule->horn-clause")
			__e.TailApply(__defun__shen_4f__error, reg6805)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.rule->horn-clause", value: __defun__shen_4rule_1_6horn_1clause})

	__defun__shen_4rule_1_6horn_1clause_1head = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1779 := __args[0]
		_ = V1779
		V1780 := __args[1]
		_ = V1780
		reg6807 := Call(__e, __defun__shen_4mode_1ify, V1780)
		reg6808 := MakeSymbol("Context_1957")
		reg6809 := Nil
		reg6810 := PrimCons(reg6808, reg6809)
		reg6811 := PrimCons(reg6807, reg6810)
		reg6812 := PrimCons(V1779, reg6811)
		__e.Return(reg6812)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.rule->horn-clause-head", value: __defun__shen_4rule_1_6horn_1clause_1head})

	__defun__shen_4mode_1ify = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1782 := __args[0]
		_ = V1782
		reg6813 := PrimIsPair(V1782)
		var reg6847 Obj
		if reg6813 == True {
			reg6814 := PrimTail(V1782)
			reg6815 := PrimIsPair(reg6814)
			var reg6842 Obj
			if reg6815 == True {
				reg6816 := MakeSymbol(":")
				reg6817 := PrimTail(V1782)
				reg6818 := PrimHead(reg6817)
				reg6819 := PrimEqual(reg6816, reg6818)
				var reg6837 Obj
				if reg6819 == True {
					reg6820 := PrimTail(V1782)
					reg6821 := PrimTail(reg6820)
					reg6822 := PrimIsPair(reg6821)
					var reg6832 Obj
					if reg6822 == True {
						reg6823 := Nil
						reg6824 := PrimTail(V1782)
						reg6825 := PrimTail(reg6824)
						reg6826 := PrimTail(reg6825)
						reg6827 := PrimEqual(reg6823, reg6826)
						var reg6830 Obj
						if reg6827 == True {
							reg6828 := True
							reg6830 = reg6828
						} else {
							reg6829 := False
							reg6830 = reg6829
						}
						reg6832 = reg6830
					} else {
						reg6831 := False
						reg6832 = reg6831
					}
					var reg6835 Obj
					if reg6832 == True {
						reg6833 := True
						reg6835 = reg6833
					} else {
						reg6834 := False
						reg6835 = reg6834
					}
					reg6837 = reg6835
				} else {
					reg6836 := False
					reg6837 = reg6836
				}
				var reg6840 Obj
				if reg6837 == True {
					reg6838 := True
					reg6840 = reg6838
				} else {
					reg6839 := False
					reg6840 = reg6839
				}
				reg6842 = reg6840
			} else {
				reg6841 := False
				reg6842 = reg6841
			}
			var reg6845 Obj
			if reg6842 == True {
				reg6843 := True
				reg6845 = reg6843
			} else {
				reg6844 := False
				reg6845 = reg6844
			}
			reg6847 = reg6845
		} else {
			reg6846 := False
			reg6847 = reg6846
		}
		if reg6847 == True {
			reg6848 := MakeSymbol("mode")
			reg6849 := PrimHead(V1782)
			reg6850 := MakeSymbol(":")
			reg6851 := MakeSymbol("mode")
			reg6852 := PrimTail(V1782)
			reg6853 := PrimTail(reg6852)
			reg6854 := PrimHead(reg6853)
			reg6855 := MakeSymbol("+")
			reg6856 := Nil
			reg6857 := PrimCons(reg6855, reg6856)
			reg6858 := PrimCons(reg6854, reg6857)
			reg6859 := PrimCons(reg6851, reg6858)
			reg6860 := Nil
			reg6861 := PrimCons(reg6859, reg6860)
			reg6862 := PrimCons(reg6850, reg6861)
			reg6863 := PrimCons(reg6849, reg6862)
			reg6864 := MakeSymbol("-")
			reg6865 := Nil
			reg6866 := PrimCons(reg6864, reg6865)
			reg6867 := PrimCons(reg6863, reg6866)
			reg6868 := PrimCons(reg6848, reg6867)
			__e.Return(reg6868)
			return
		} else {
			__e.Return(V1782)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.mode-ify", value: __defun__shen_4mode_1ify})

	__defun__shen_4rule_1_6horn_1clause_1body = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1786 := __args[0]
		_ = V1786
		V1787 := __args[1]
		_ = V1787
		V1788 := __args[2]
		_ = V1788
		reg6869 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4extract__vars, X)
			return
		}, 1)
		reg6871 := Call(__e, __defun__map, reg6869, V1788)
		Variables := reg6871
		_ = Variables
		reg6872 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			reg6873 := MakeSymbol("shen.cl")
			__e.TailApply(__defun__gensym, reg6873)
			return
		}, 1)
		reg6875 := Call(__e, __defun__map, reg6872, V1788)
		Predicates := reg6875
		_ = Predicates
		reg6876 := MakeSymbol("Context_1957")
		reg6877 := MakeSymbol("Context1_1957")
		reg6878 := Call(__e, __defun__shen_4construct_1search_1literals, Predicates, Variables, reg6876, reg6877)
		SearchLiterals := reg6878
		_ = SearchLiterals
		reg6879 := Call(__e, __defun__shen_4construct_1search_1clauses, Predicates, V1788, Variables)
		SearchClauses := reg6879
		_ = SearchClauses
		reg6880 := Call(__e, __defun__shen_4construct_1side_1literals, V1786)
		SideLiterals := reg6880
		_ = SideLiterals
		reg6881 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			reg6882 := Call(__e, __defun__empty_2, V1788)
			__e.TailApply(__defun__shen_4construct_1premiss_1literal, X, reg6882)
			return
		}, 1)
		reg6884 := Call(__e, __defun__map, reg6881, V1787)
		PremissLiterals := reg6884
		_ = PremissLiterals
		reg6885 := Call(__e, __defun__append, SideLiterals, PremissLiterals)
		__e.TailApply(__defun__append, SearchLiterals, reg6885)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.rule->horn-clause-body", value: __defun__shen_4rule_1_6horn_1clause_1body})

	__defun__shen_4construct_1search_1literals = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1797 := __args[0]
		_ = V1797
		V1798 := __args[1]
		_ = V1798
		V1799 := __args[2]
		_ = V1799
		V1800 := __args[3]
		_ = V1800
		reg6887 := Nil
		reg6888 := PrimEqual(reg6887, V1797)
		var reg6895 Obj
		if reg6888 == True {
			reg6889 := Nil
			reg6890 := PrimEqual(reg6889, V1798)
			var reg6893 Obj
			if reg6890 == True {
				reg6891 := True
				reg6893 = reg6891
			} else {
				reg6892 := False
				reg6893 = reg6892
			}
			reg6895 = reg6893
		} else {
			reg6894 := False
			reg6895 = reg6894
		}
		if reg6895 == True {
			reg6896 := Nil
			__e.Return(reg6896)
			return
		} else {
			__e.TailApply(__defun__shen_4csl_1help, V1797, V1798, V1799, V1800)
			return
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.construct-search-literals", value: __defun__shen_4construct_1search_1literals})

	__defun__shen_4csl_1help = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1807 := __args[0]
		_ = V1807
		V1808 := __args[1]
		_ = V1808
		V1809 := __args[2]
		_ = V1809
		V1810 := __args[3]
		_ = V1810
		reg6898 := Nil
		reg6899 := PrimEqual(reg6898, V1807)
		var reg6906 Obj
		if reg6899 == True {
			reg6900 := Nil
			reg6901 := PrimEqual(reg6900, V1808)
			var reg6904 Obj
			if reg6901 == True {
				reg6902 := True
				reg6904 = reg6902
			} else {
				reg6903 := False
				reg6904 = reg6903
			}
			reg6906 = reg6904
		} else {
			reg6905 := False
			reg6906 = reg6905
		}
		if reg6906 == True {
			reg6907 := MakeSymbol("bind")
			reg6908 := MakeSymbol("ContextOut_1957")
			reg6909 := Nil
			reg6910 := PrimCons(V1809, reg6909)
			reg6911 := PrimCons(reg6908, reg6910)
			reg6912 := PrimCons(reg6907, reg6911)
			reg6913 := Nil
			reg6914 := PrimCons(reg6912, reg6913)
			__e.Return(reg6914)
			return
		} else {
			reg6915 := PrimIsPair(V1807)
			var reg6921 Obj
			if reg6915 == True {
				reg6916 := PrimIsPair(V1808)
				var reg6919 Obj
				if reg6916 == True {
					reg6917 := True
					reg6919 = reg6917
				} else {
					reg6918 := False
					reg6919 = reg6918
				}
				reg6921 = reg6919
			} else {
				reg6920 := False
				reg6921 = reg6920
			}
			if reg6921 == True {
				reg6922 := PrimHead(V1807)
				reg6923 := PrimHead(V1808)
				reg6924 := PrimCons(V1810, reg6923)
				reg6925 := PrimCons(V1809, reg6924)
				reg6926 := PrimCons(reg6922, reg6925)
				reg6927 := PrimTail(V1807)
				reg6928 := PrimTail(V1808)
				reg6929 := MakeSymbol("Context")
				reg6930 := Call(__e, __defun__gensym, reg6929)
				reg6931 := Call(__e, __defun__shen_4csl_1help, reg6927, reg6928, V1810, reg6930)
				reg6932 := PrimCons(reg6926, reg6931)
				__e.Return(reg6932)
				return
			} else {
				reg6933 := MakeSymbol("shen.csl-help")
				__e.TailApply(__defun__shen_4f__error, reg6933)
				return
			}
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.csl-help", value: __defun__shen_4csl_1help})

	__defun__shen_4construct_1search_1clauses = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1814 := __args[0]
		_ = V1814
		V1815 := __args[1]
		_ = V1815
		V1816 := __args[2]
		_ = V1816
		reg6935 := Nil
		reg6936 := PrimEqual(reg6935, V1814)
		var reg6950 Obj
		if reg6936 == True {
			reg6937 := Nil
			reg6938 := PrimEqual(reg6937, V1815)
			var reg6945 Obj
			if reg6938 == True {
				reg6939 := Nil
				reg6940 := PrimEqual(reg6939, V1816)
				var reg6943 Obj
				if reg6940 == True {
					reg6941 := True
					reg6943 = reg6941
				} else {
					reg6942 := False
					reg6943 = reg6942
				}
				reg6945 = reg6943
			} else {
				reg6944 := False
				reg6945 = reg6944
			}
			var reg6948 Obj
			if reg6945 == True {
				reg6946 := True
				reg6948 = reg6946
			} else {
				reg6947 := False
				reg6948 = reg6947
			}
			reg6950 = reg6948
		} else {
			reg6949 := False
			reg6950 = reg6949
		}
		if reg6950 == True {
			reg6951 := MakeSymbol("shen.skip")
			__e.Return(reg6951)
			return
		} else {
			reg6952 := PrimIsPair(V1814)
			var reg6964 Obj
			if reg6952 == True {
				reg6953 := PrimIsPair(V1815)
				var reg6959 Obj
				if reg6953 == True {
					reg6954 := PrimIsPair(V1816)
					var reg6957 Obj
					if reg6954 == True {
						reg6955 := True
						reg6957 = reg6955
					} else {
						reg6956 := False
						reg6957 = reg6956
					}
					reg6959 = reg6957
				} else {
					reg6958 := False
					reg6959 = reg6958
				}
				var reg6962 Obj
				if reg6959 == True {
					reg6960 := True
					reg6962 = reg6960
				} else {
					reg6961 := False
					reg6962 = reg6961
				}
				reg6964 = reg6962
			} else {
				reg6963 := False
				reg6964 = reg6963
			}
			if reg6964 == True {
				reg6965 := PrimHead(V1814)
				reg6966 := PrimHead(V1815)
				reg6967 := PrimHead(V1816)
				reg6968 := Call(__e, __defun__shen_4construct_1search_1clause, reg6965, reg6966, reg6967)
				_ = reg6968
				reg6969 := PrimTail(V1814)
				reg6970 := PrimTail(V1815)
				reg6971 := PrimTail(V1816)
				__e.TailApply(__defun__shen_4construct_1search_1clauses, reg6969, reg6970, reg6971)
				return
			} else {
				reg6973 := MakeSymbol("shen.construct-search-clauses")
				__e.TailApply(__defun__shen_4f__error, reg6973)
				return
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.construct-search-clauses", value: __defun__shen_4construct_1search_1clauses})

	__defun__shen_4construct_1search_1clause = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1820 := __args[0]
		_ = V1820
		V1821 := __args[1]
		_ = V1821
		V1822 := __args[2]
		_ = V1822
		reg6975 := Call(__e, __defun__shen_4construct_1base_1search_1clause, V1820, V1821, V1822)
		reg6976 := Call(__e, __defun__shen_4construct_1recursive_1search_1clause, V1820, V1821, V1822)
		reg6977 := Nil
		reg6978 := PrimCons(reg6976, reg6977)
		reg6979 := PrimCons(reg6975, reg6978)
		__e.TailApply(__defun__shen_4s_1prolog, reg6979)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.construct-search-clause", value: __defun__shen_4construct_1search_1clause})

	__defun__shen_4construct_1base_1search_1clause = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1826 := __args[0]
		_ = V1826
		V1827 := __args[1]
		_ = V1827
		V1828 := __args[2]
		_ = V1828
		reg6981 := Call(__e, __defun__shen_4mode_1ify, V1827)
		reg6982 := MakeSymbol("In_1957")
		reg6983 := PrimCons(reg6981, reg6982)
		reg6984 := MakeSymbol("In_1957")
		reg6985 := PrimCons(reg6984, V1828)
		reg6986 := PrimCons(reg6983, reg6985)
		reg6987 := PrimCons(V1826, reg6986)
		reg6988 := MakeSymbol(":-")
		reg6989 := Nil
		reg6990 := Nil
		reg6991 := PrimCons(reg6989, reg6990)
		reg6992 := PrimCons(reg6988, reg6991)
		reg6993 := PrimCons(reg6987, reg6992)
		__e.Return(reg6993)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.construct-base-search-clause", value: __defun__shen_4construct_1base_1search_1clause})

	__defun__shen_4construct_1recursive_1search_1clause = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1832 := __args[0]
		_ = V1832
		V1833 := __args[1]
		_ = V1833
		V1834 := __args[2]
		_ = V1834
		reg6994 := MakeSymbol("Assumption_1957")
		reg6995 := MakeSymbol("Assumptions_1957")
		reg6996 := PrimCons(reg6994, reg6995)
		reg6997 := MakeSymbol("Assumption_1957")
		reg6998 := MakeSymbol("Out_1957")
		reg6999 := PrimCons(reg6997, reg6998)
		reg7000 := PrimCons(reg6999, V1834)
		reg7001 := PrimCons(reg6996, reg7000)
		reg7002 := PrimCons(V1832, reg7001)
		reg7003 := MakeSymbol(":-")
		reg7004 := MakeSymbol("Assumptions_1957")
		reg7005 := MakeSymbol("Out_1957")
		reg7006 := PrimCons(reg7005, V1834)
		reg7007 := PrimCons(reg7004, reg7006)
		reg7008 := PrimCons(V1832, reg7007)
		reg7009 := Nil
		reg7010 := PrimCons(reg7008, reg7009)
		reg7011 := Nil
		reg7012 := PrimCons(reg7010, reg7011)
		reg7013 := PrimCons(reg7003, reg7012)
		reg7014 := PrimCons(reg7002, reg7013)
		__e.Return(reg7014)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.construct-recursive-search-clause", value: __defun__shen_4construct_1recursive_1search_1clause})

	__defun__shen_4construct_1side_1literals = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1840 := __args[0]
		_ = V1840
		reg7015 := Nil
		reg7016 := PrimEqual(reg7015, V1840)
		if reg7016 == True {
			reg7017 := Nil
			__e.Return(reg7017)
			return
		} else {
			reg7018 := PrimIsPair(V1840)
			var reg7052 Obj
			if reg7018 == True {
				reg7019 := PrimHead(V1840)
				reg7020 := PrimIsPair(reg7019)
				var reg7047 Obj
				if reg7020 == True {
					reg7021 := MakeSymbol("if")
					reg7022 := PrimHead(V1840)
					reg7023 := PrimHead(reg7022)
					reg7024 := PrimEqual(reg7021, reg7023)
					var reg7042 Obj
					if reg7024 == True {
						reg7025 := PrimHead(V1840)
						reg7026 := PrimTail(reg7025)
						reg7027 := PrimIsPair(reg7026)
						var reg7037 Obj
						if reg7027 == True {
							reg7028 := Nil
							reg7029 := PrimHead(V1840)
							reg7030 := PrimTail(reg7029)
							reg7031 := PrimTail(reg7030)
							reg7032 := PrimEqual(reg7028, reg7031)
							var reg7035 Obj
							if reg7032 == True {
								reg7033 := True
								reg7035 = reg7033
							} else {
								reg7034 := False
								reg7035 = reg7034
							}
							reg7037 = reg7035
						} else {
							reg7036 := False
							reg7037 = reg7036
						}
						var reg7040 Obj
						if reg7037 == True {
							reg7038 := True
							reg7040 = reg7038
						} else {
							reg7039 := False
							reg7040 = reg7039
						}
						reg7042 = reg7040
					} else {
						reg7041 := False
						reg7042 = reg7041
					}
					var reg7045 Obj
					if reg7042 == True {
						reg7043 := True
						reg7045 = reg7043
					} else {
						reg7044 := False
						reg7045 = reg7044
					}
					reg7047 = reg7045
				} else {
					reg7046 := False
					reg7047 = reg7046
				}
				var reg7050 Obj
				if reg7047 == True {
					reg7048 := True
					reg7050 = reg7048
				} else {
					reg7049 := False
					reg7050 = reg7049
				}
				reg7052 = reg7050
			} else {
				reg7051 := False
				reg7052 = reg7051
			}
			if reg7052 == True {
				reg7053 := MakeSymbol("when")
				reg7054 := PrimHead(V1840)
				reg7055 := PrimTail(reg7054)
				reg7056 := PrimCons(reg7053, reg7055)
				reg7057 := PrimTail(V1840)
				reg7058 := Call(__e, __defun__shen_4construct_1side_1literals, reg7057)
				reg7059 := PrimCons(reg7056, reg7058)
				__e.Return(reg7059)
				return
			} else {
				reg7060 := PrimIsPair(V1840)
				var reg7104 Obj
				if reg7060 == True {
					reg7061 := PrimHead(V1840)
					reg7062 := PrimIsPair(reg7061)
					var reg7099 Obj
					if reg7062 == True {
						reg7063 := MakeSymbol("let")
						reg7064 := PrimHead(V1840)
						reg7065 := PrimHead(reg7064)
						reg7066 := PrimEqual(reg7063, reg7065)
						var reg7094 Obj
						if reg7066 == True {
							reg7067 := PrimHead(V1840)
							reg7068 := PrimTail(reg7067)
							reg7069 := PrimIsPair(reg7068)
							var reg7089 Obj
							if reg7069 == True {
								reg7070 := PrimHead(V1840)
								reg7071 := PrimTail(reg7070)
								reg7072 := PrimTail(reg7071)
								reg7073 := PrimIsPair(reg7072)
								var reg7084 Obj
								if reg7073 == True {
									reg7074 := Nil
									reg7075 := PrimHead(V1840)
									reg7076 := PrimTail(reg7075)
									reg7077 := PrimTail(reg7076)
									reg7078 := PrimTail(reg7077)
									reg7079 := PrimEqual(reg7074, reg7078)
									var reg7082 Obj
									if reg7079 == True {
										reg7080 := True
										reg7082 = reg7080
									} else {
										reg7081 := False
										reg7082 = reg7081
									}
									reg7084 = reg7082
								} else {
									reg7083 := False
									reg7084 = reg7083
								}
								var reg7087 Obj
								if reg7084 == True {
									reg7085 := True
									reg7087 = reg7085
								} else {
									reg7086 := False
									reg7087 = reg7086
								}
								reg7089 = reg7087
							} else {
								reg7088 := False
								reg7089 = reg7088
							}
							var reg7092 Obj
							if reg7089 == True {
								reg7090 := True
								reg7092 = reg7090
							} else {
								reg7091 := False
								reg7092 = reg7091
							}
							reg7094 = reg7092
						} else {
							reg7093 := False
							reg7094 = reg7093
						}
						var reg7097 Obj
						if reg7094 == True {
							reg7095 := True
							reg7097 = reg7095
						} else {
							reg7096 := False
							reg7097 = reg7096
						}
						reg7099 = reg7097
					} else {
						reg7098 := False
						reg7099 = reg7098
					}
					var reg7102 Obj
					if reg7099 == True {
						reg7100 := True
						reg7102 = reg7100
					} else {
						reg7101 := False
						reg7102 = reg7101
					}
					reg7104 = reg7102
				} else {
					reg7103 := False
					reg7104 = reg7103
				}
				if reg7104 == True {
					reg7105 := MakeSymbol("is")
					reg7106 := PrimHead(V1840)
					reg7107 := PrimTail(reg7106)
					reg7108 := PrimCons(reg7105, reg7107)
					reg7109 := PrimTail(V1840)
					reg7110 := Call(__e, __defun__shen_4construct_1side_1literals, reg7109)
					reg7111 := PrimCons(reg7108, reg7110)
					__e.Return(reg7111)
					return
				} else {
					reg7112 := PrimIsPair(V1840)
					if reg7112 == True {
						reg7113 := PrimTail(V1840)
						__e.TailApply(__defun__shen_4construct_1side_1literals, reg7113)
						return
					} else {
						reg7115 := MakeSymbol("shen.construct-side-literals")
						__e.TailApply(__defun__shen_4f__error, reg7115)
						return
					}
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.construct-side-literals", value: __defun__shen_4construct_1side_1literals})

	__defun__shen_4construct_1premiss_1literal = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1847 := __args[0]
		_ = V1847
		V1848 := __args[1]
		_ = V1848
		reg7117 := Call(__e, __defun__tuple_2, V1847)
		if reg7117 == True {
			reg7118 := MakeSymbol("shen.t*")
			reg7119 := Call(__e, __defun__snd, V1847)
			reg7120 := Call(__e, __defun__shen_4recursive__cons__form, reg7119)
			reg7121 := Call(__e, __defun__fst, V1847)
			reg7122 := Call(__e, __defun__shen_4construct_1context, V1848, reg7121)
			reg7123 := Nil
			reg7124 := PrimCons(reg7122, reg7123)
			reg7125 := PrimCons(reg7120, reg7124)
			reg7126 := PrimCons(reg7118, reg7125)
			__e.Return(reg7126)
			return
		} else {
			reg7127 := MakeSymbol("!")
			reg7128 := PrimEqual(reg7127, V1847)
			if reg7128 == True {
				reg7129 := MakeSymbol("cut")
				reg7130 := MakeSymbol("Throwcontrol")
				reg7131 := Nil
				reg7132 := PrimCons(reg7130, reg7131)
				reg7133 := PrimCons(reg7129, reg7132)
				__e.Return(reg7133)
				return
			} else {
				reg7134 := MakeSymbol("shen.construct-premiss-literal")
				__e.TailApply(__defun__shen_4f__error, reg7134)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.construct-premiss-literal", value: __defun__shen_4construct_1premiss_1literal})

	__defun__shen_4construct_1context = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1851 := __args[0]
		_ = V1851
		V1852 := __args[1]
		_ = V1852
		reg7136 := True
		reg7137 := PrimEqual(reg7136, V1851)
		var reg7144 Obj
		if reg7137 == True {
			reg7138 := Nil
			reg7139 := PrimEqual(reg7138, V1852)
			var reg7142 Obj
			if reg7139 == True {
				reg7140 := True
				reg7142 = reg7140
			} else {
				reg7141 := False
				reg7142 = reg7141
			}
			reg7144 = reg7142
		} else {
			reg7143 := False
			reg7144 = reg7143
		}
		if reg7144 == True {
			reg7145 := MakeSymbol("Context_1957")
			__e.Return(reg7145)
			return
		} else {
			reg7146 := False
			reg7147 := PrimEqual(reg7146, V1851)
			var reg7154 Obj
			if reg7147 == True {
				reg7148 := Nil
				reg7149 := PrimEqual(reg7148, V1852)
				var reg7152 Obj
				if reg7149 == True {
					reg7150 := True
					reg7152 = reg7150
				} else {
					reg7151 := False
					reg7152 = reg7151
				}
				reg7154 = reg7152
			} else {
				reg7153 := False
				reg7154 = reg7153
			}
			if reg7154 == True {
				reg7155 := MakeSymbol("ContextOut_1957")
				__e.Return(reg7155)
				return
			} else {
				reg7156 := PrimIsPair(V1852)
				if reg7156 == True {
					reg7157 := MakeSymbol("cons")
					reg7158 := PrimHead(V1852)
					reg7159 := Call(__e, __defun__shen_4recursive__cons__form, reg7158)
					reg7160 := PrimTail(V1852)
					reg7161 := Call(__e, __defun__shen_4construct_1context, V1851, reg7160)
					reg7162 := Nil
					reg7163 := PrimCons(reg7161, reg7162)
					reg7164 := PrimCons(reg7159, reg7163)
					reg7165 := PrimCons(reg7157, reg7164)
					__e.Return(reg7165)
					return
				} else {
					reg7166 := MakeSymbol("shen.construct-context")
					__e.TailApply(__defun__shen_4f__error, reg7166)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.construct-context", value: __defun__shen_4construct_1context})

	__defun__shen_4recursive__cons__form = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1854 := __args[0]
		_ = V1854
		reg7168 := PrimIsPair(V1854)
		if reg7168 == True {
			reg7169 := MakeSymbol("cons")
			reg7170 := PrimHead(V1854)
			reg7171 := Call(__e, __defun__shen_4recursive__cons__form, reg7170)
			reg7172 := PrimTail(V1854)
			reg7173 := Call(__e, __defun__shen_4recursive__cons__form, reg7172)
			reg7174 := Nil
			reg7175 := PrimCons(reg7173, reg7174)
			reg7176 := PrimCons(reg7171, reg7175)
			reg7177 := PrimCons(reg7169, reg7176)
			__e.Return(reg7177)
			return
		} else {
			__e.Return(V1854)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.recursive_cons_form", value: __defun__shen_4recursive__cons__form})

	__defun__preclude = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1856 := __args[0]
		_ = V1856
		reg7178 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4intern_1type, X)
			return
		}, 1)
		reg7180 := Call(__e, __defun__map, reg7178, V1856)
		__e.TailApply(__defun__shen_4preclude_1h, reg7180)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "preclude", value: __defun__preclude})

	__defun__shen_4preclude_1h = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1858 := __args[0]
		_ = V1858
		reg7182 := MakeSymbol("shen.*datatypes*")
		reg7183 := MakeSymbol("shen.*datatypes*")
		reg7184 := PrimValue(reg7183)
		reg7185 := Call(__e, __defun__difference, reg7184, V1858)
		reg7186 := PrimSet(reg7182, reg7185)
		FilterDatatypes := reg7186
		_ = FilterDatatypes
		reg7187 := MakeSymbol("shen.*datatypes*")
		reg7188 := PrimValue(reg7187)
		__e.Return(reg7188)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.preclude-h", value: __defun__shen_4preclude_1h})

	__defun__include = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1860 := __args[0]
		_ = V1860
		reg7189 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4intern_1type, X)
			return
		}, 1)
		reg7191 := Call(__e, __defun__map, reg7189, V1860)
		__e.TailApply(__defun__shen_4include_1h, reg7191)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "include", value: __defun__include})

	__defun__shen_4include_1h = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1862 := __args[0]
		_ = V1862
		reg7193 := MakeSymbol("shen.*alldatatypes*")
		reg7194 := PrimValue(reg7193)
		reg7195 := Call(__e, __defun__intersection, V1862, reg7194)
		ValidTypes := reg7195
		_ = ValidTypes
		reg7196 := MakeSymbol("shen.*datatypes*")
		reg7197 := MakeSymbol("shen.*datatypes*")
		reg7198 := PrimValue(reg7197)
		reg7199 := Call(__e, __defun__union, ValidTypes, reg7198)
		reg7200 := PrimSet(reg7196, reg7199)
		NewDatatypes := reg7200
		_ = NewDatatypes
		reg7201 := MakeSymbol("shen.*datatypes*")
		reg7202 := PrimValue(reg7201)
		__e.Return(reg7202)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.include-h", value: __defun__shen_4include_1h})

	__defun__preclude_1all_1but = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1864 := __args[0]
		_ = V1864
		reg7203 := MakeSymbol("shen.*alldatatypes*")
		reg7204 := PrimValue(reg7203)
		reg7205 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4intern_1type, X)
			return
		}, 1)
		reg7207 := Call(__e, __defun__map, reg7205, V1864)
		reg7208 := Call(__e, __defun__difference, reg7204, reg7207)
		__e.TailApply(__defun__shen_4preclude_1h, reg7208)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "preclude-all-but", value: __defun__preclude_1all_1but})

	__defun__include_1all_1but = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1866 := __args[0]
		_ = V1866
		reg7210 := MakeSymbol("shen.*alldatatypes*")
		reg7211 := PrimValue(reg7210)
		reg7212 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4intern_1type, X)
			return
		}, 1)
		reg7214 := Call(__e, __defun__map, reg7212, V1866)
		reg7215 := Call(__e, __defun__difference, reg7211, reg7214)
		__e.TailApply(__defun__shen_4include_1h, reg7215)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "include-all-but", value: __defun__include_1all_1but})

	__defun__shen_4synonyms_1help = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1872 := __args[0]
		_ = V1872
		reg7217 := Nil
		reg7218 := PrimEqual(reg7217, V1872)
		if reg7218 == True {
			reg7219 := MakeSymbol("shen.*tc*")
			reg7220 := PrimValue(reg7219)
			reg7221 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(__defun__shen_4demod_1rule, X)
				return
			}, 1)
			reg7223 := MakeSymbol("shen.*synonyms*")
			reg7224 := PrimValue(reg7223)
			reg7225 := Call(__e, __defun__mapcan, reg7221, reg7224)
			__e.TailApply(__defun__shen_4update_1demodulation_1function, reg7220, reg7225)
			return
		} else {
			reg7227 := PrimIsPair(V1872)
			var reg7234 Obj
			if reg7227 == True {
				reg7228 := PrimTail(V1872)
				reg7229 := PrimIsPair(reg7228)
				var reg7232 Obj
				if reg7229 == True {
					reg7230 := True
					reg7232 = reg7230
				} else {
					reg7231 := False
					reg7232 = reg7231
				}
				reg7234 = reg7232
			} else {
				reg7233 := False
				reg7234 = reg7233
			}
			if reg7234 == True {
				reg7235 := PrimTail(V1872)
				reg7236 := PrimHead(reg7235)
				reg7237 := Call(__e, __defun__shen_4extract__vars, reg7236)
				reg7238 := PrimHead(V1872)
				reg7239 := Call(__e, __defun__shen_4extract__vars, reg7238)
				reg7240 := Call(__e, __defun__difference, reg7237, reg7239)
				Vs := reg7240
				_ = Vs
				reg7241 := Call(__e, __defun__empty_2, Vs)
				if reg7241 == True {
					reg7242 := PrimHead(V1872)
					reg7243 := PrimTail(V1872)
					reg7244 := PrimHead(reg7243)
					reg7245 := Nil
					reg7246 := PrimCons(reg7244, reg7245)
					reg7247 := PrimCons(reg7242, reg7246)
					reg7248 := MakeSymbol("shen.*synonyms*")
					reg7249 := Call(__e, __defun__shen_4pushnew, reg7247, reg7248)
					_ = reg7249
					reg7250 := PrimTail(V1872)
					reg7251 := PrimTail(reg7250)
					__e.TailApply(__defun__shen_4synonyms_1help, reg7251)
					return
				} else {
					reg7253 := PrimTail(V1872)
					reg7254 := PrimHead(reg7253)
					__e.TailApply(__defun__shen_4free__variable__warnings, reg7254, Vs)
					return
				}
			} else {
				reg7256 := MakeString("odd number of synonyms\n")
				reg7257 := PrimSimpleError(reg7256)
				__e.Return(reg7257)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.synonyms-help", value: __defun__shen_4synonyms_1help})

	__defun__shen_4pushnew = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1875 := __args[0]
		_ = V1875
		V1876 := __args[1]
		_ = V1876
		reg7258 := PrimValue(V1876)
		reg7259 := Call(__e, __defun__element_2, V1875, reg7258)
		if reg7259 == True {
			reg7260 := PrimValue(V1876)
			__e.Return(reg7260)
			return
		} else {
			reg7261 := PrimValue(V1876)
			reg7262 := PrimCons(V1875, reg7261)
			reg7263 := PrimSet(V1876, reg7262)
			__e.Return(reg7263)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.pushnew", value: __defun__shen_4pushnew})

	__defun__shen_4demod_1rule = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1878 := __args[0]
		_ = V1878
		reg7264 := PrimIsPair(V1878)
		var reg7280 Obj
		if reg7264 == True {
			reg7265 := PrimTail(V1878)
			reg7266 := PrimIsPair(reg7265)
			var reg7275 Obj
			if reg7266 == True {
				reg7267 := Nil
				reg7268 := PrimTail(V1878)
				reg7269 := PrimTail(reg7268)
				reg7270 := PrimEqual(reg7267, reg7269)
				var reg7273 Obj
				if reg7270 == True {
					reg7271 := True
					reg7273 = reg7271
				} else {
					reg7272 := False
					reg7273 = reg7272
				}
				reg7275 = reg7273
			} else {
				reg7274 := False
				reg7275 = reg7274
			}
			var reg7278 Obj
			if reg7275 == True {
				reg7276 := True
				reg7278 = reg7276
			} else {
				reg7277 := False
				reg7278 = reg7277
			}
			reg7280 = reg7278
		} else {
			reg7279 := False
			reg7280 = reg7279
		}
		if reg7280 == True {
			reg7281 := PrimHead(V1878)
			reg7282 := Call(__e, __defun__shen_4rcons__form, reg7281)
			reg7283 := MakeSymbol("->")
			reg7284 := PrimTail(V1878)
			reg7285 := PrimHead(reg7284)
			reg7286 := Call(__e, __defun__shen_4rcons__form, reg7285)
			reg7287 := Nil
			reg7288 := PrimCons(reg7286, reg7287)
			reg7289 := PrimCons(reg7283, reg7288)
			reg7290 := PrimCons(reg7282, reg7289)
			__e.Return(reg7290)
			return
		} else {
			reg7291 := MakeSymbol("shen.demod-rule")
			__e.TailApply(__defun__shen_4f__error, reg7291)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.demod-rule", value: __defun__shen_4demod_1rule})

	__defun__shen_4lambda_1of_1defun = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1884 := __args[0]
		_ = V1884
		reg7293 := PrimIsPair(V1884)
		var reg7356 Obj
		if reg7293 == True {
			reg7294 := MakeSymbol("defun")
			reg7295 := PrimHead(V1884)
			reg7296 := PrimEqual(reg7294, reg7295)
			var reg7351 Obj
			if reg7296 == True {
				reg7297 := PrimTail(V1884)
				reg7298 := PrimIsPair(reg7297)
				var reg7346 Obj
				if reg7298 == True {
					reg7299 := PrimTail(V1884)
					reg7300 := PrimTail(reg7299)
					reg7301 := PrimIsPair(reg7300)
					var reg7341 Obj
					if reg7301 == True {
						reg7302 := PrimTail(V1884)
						reg7303 := PrimTail(reg7302)
						reg7304 := PrimHead(reg7303)
						reg7305 := PrimIsPair(reg7304)
						var reg7336 Obj
						if reg7305 == True {
							reg7306 := Nil
							reg7307 := PrimTail(V1884)
							reg7308 := PrimTail(reg7307)
							reg7309 := PrimHead(reg7308)
							reg7310 := PrimTail(reg7309)
							reg7311 := PrimEqual(reg7306, reg7310)
							var reg7331 Obj
							if reg7311 == True {
								reg7312 := PrimTail(V1884)
								reg7313 := PrimTail(reg7312)
								reg7314 := PrimTail(reg7313)
								reg7315 := PrimIsPair(reg7314)
								var reg7326 Obj
								if reg7315 == True {
									reg7316 := Nil
									reg7317 := PrimTail(V1884)
									reg7318 := PrimTail(reg7317)
									reg7319 := PrimTail(reg7318)
									reg7320 := PrimTail(reg7319)
									reg7321 := PrimEqual(reg7316, reg7320)
									var reg7324 Obj
									if reg7321 == True {
										reg7322 := True
										reg7324 = reg7322
									} else {
										reg7323 := False
										reg7324 = reg7323
									}
									reg7326 = reg7324
								} else {
									reg7325 := False
									reg7326 = reg7325
								}
								var reg7329 Obj
								if reg7326 == True {
									reg7327 := True
									reg7329 = reg7327
								} else {
									reg7328 := False
									reg7329 = reg7328
								}
								reg7331 = reg7329
							} else {
								reg7330 := False
								reg7331 = reg7330
							}
							var reg7334 Obj
							if reg7331 == True {
								reg7332 := True
								reg7334 = reg7332
							} else {
								reg7333 := False
								reg7334 = reg7333
							}
							reg7336 = reg7334
						} else {
							reg7335 := False
							reg7336 = reg7335
						}
						var reg7339 Obj
						if reg7336 == True {
							reg7337 := True
							reg7339 = reg7337
						} else {
							reg7338 := False
							reg7339 = reg7338
						}
						reg7341 = reg7339
					} else {
						reg7340 := False
						reg7341 = reg7340
					}
					var reg7344 Obj
					if reg7341 == True {
						reg7342 := True
						reg7344 = reg7342
					} else {
						reg7343 := False
						reg7344 = reg7343
					}
					reg7346 = reg7344
				} else {
					reg7345 := False
					reg7346 = reg7345
				}
				var reg7349 Obj
				if reg7346 == True {
					reg7347 := True
					reg7349 = reg7347
				} else {
					reg7348 := False
					reg7349 = reg7348
				}
				reg7351 = reg7349
			} else {
				reg7350 := False
				reg7351 = reg7350
			}
			var reg7354 Obj
			if reg7351 == True {
				reg7352 := True
				reg7354 = reg7352
			} else {
				reg7353 := False
				reg7354 = reg7353
			}
			reg7356 = reg7354
		} else {
			reg7355 := False
			reg7356 = reg7355
		}
		if reg7356 == True {
			reg7357 := MakeSymbol("/.")
			reg7358 := PrimTail(V1884)
			reg7359 := PrimTail(reg7358)
			reg7360 := PrimHead(reg7359)
			reg7361 := PrimHead(reg7360)
			reg7362 := PrimTail(V1884)
			reg7363 := PrimTail(reg7362)
			reg7364 := PrimTail(reg7363)
			reg7365 := PrimCons(reg7361, reg7364)
			reg7366 := PrimCons(reg7357, reg7365)
			__e.TailApply(__defun__eval, reg7366)
			return
		} else {
			reg7368 := MakeSymbol("shen.lambda-of-defun")
			__e.TailApply(__defun__shen_4f__error, reg7368)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.lambda-of-defun", value: __defun__shen_4lambda_1of_1defun})

	__defun__shen_4update_1demodulation_1function = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1887 := __args[0]
		_ = V1887
		V1888 := __args[1]
		_ = V1888
		reg7370 := MakeSymbol("-")
		reg7371 := Call(__e, __defun__tc, reg7370)
		_ = reg7371
		reg7372 := MakeSymbol("shen.*demodulation-function*")
		reg7373 := MakeSymbol("define")
		reg7374 := MakeSymbol("shen.demod")
		reg7375 := Call(__e, __defun__shen_4default_1rule)
		reg7376 := Call(__e, __defun__append, V1888, reg7375)
		reg7377 := PrimCons(reg7374, reg7376)
		reg7378 := PrimCons(reg7373, reg7377)
		reg7379 := Call(__e, __defun__shen_4elim_1def, reg7378)
		reg7380 := Call(__e, __defun__shen_4lambda_1of_1defun, reg7379)
		reg7381 := PrimSet(reg7372, reg7380)
		_ = reg7381
		var reg7385 Obj
		if V1887 == True {
			reg7382 := MakeSymbol("+")
			reg7383 := Call(__e, __defun__tc, reg7382)
			reg7385 = reg7383
		} else {
			reg7384 := MakeSymbol("shen.skip")
			reg7385 = reg7384
		}
		_ = reg7385
		reg7386 := MakeSymbol("synonyms")
		__e.Return(reg7386)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.update-demodulation-function", value: __defun__shen_4update_1demodulation_1function})

	__defun__shen_4default_1rule = MakeNative(func(__e Evaluator, __args ...Obj) {
		reg7387 := MakeSymbol("X")
		reg7388 := MakeSymbol("->")
		reg7389 := MakeSymbol("X")
		reg7390 := Nil
		reg7391 := PrimCons(reg7389, reg7390)
		reg7392 := PrimCons(reg7388, reg7391)
		reg7393 := PrimCons(reg7387, reg7392)
		__e.Return(reg7393)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.default-rule", value: __defun__shen_4default_1rule})

}
