package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4f__error Obj               // shen.f_error
var __defun__shen_4tracked_2 Obj              // shen.tracked?
var __defun__track Obj                        // track
var __defun__shen_4track_1function Obj        // shen.track-function
var __defun__shen_4insert_1tracking_1code Obj // shen.insert-tracking-code
var __defun__step Obj                         // step
var __defun__spy Obj                          // spy
var __defun__shen_4terpri_1or_1read_1char Obj // shen.terpri-or-read-char
var __defun__shen_4check_1byte Obj            // shen.check-byte
var __defun__shen_4input_1track Obj           // shen.input-track
var __defun__shen_4recursively_1print Obj     // shen.recursively-print
var __defun__shen_4spaces Obj                 // shen.spaces
var __defun__shen_4output_1track Obj          // shen.output-track
var __defun__untrack Obj                      // untrack
var __defun__profile Obj                      // profile
var __defun__shen_4profile_1help Obj          // shen.profile-help
var __defun__unprofile Obj                    // unprofile
var __defun__shen_4profile_1func Obj          // shen.profile-func
var __defun__profile_1results Obj             // profile-results
var __defun__shen_4get_1profile Obj           // shen.get-profile
var __defun__shen_4put_1profile Obj           // shen.put-profile

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg15974 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__ctx.Return(reg15974)
		return
	}, 0))
	__defun__shen_4f__error = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3105 := __args[0]
		_ = V3105
		reg15975 := MakeString("partial function ")
		reg15976 := MakeString(";\n")
		reg15977 := MakeSymbol("shen.a")
		reg15978 := __e.Call(__defun__shen_4app, V3105, reg15976, reg15977)
		reg15979 := PrimStringConcat(reg15975, reg15978)
		reg15980 := __e.Call(__defun__stoutput)
		reg15981 := __e.Call(__defun__shen_4prhush, reg15979, reg15980)
		_ = reg15981
		reg15982 := __e.Call(__defun__shen_4tracked_2, V3105)
		reg15983 := PrimNot(reg15982)
		var reg15994 Obj
		if reg15983 == True {
			reg15984 := MakeString("track ")
			reg15985 := MakeString("? ")
			reg15986 := MakeSymbol("shen.a")
			reg15987 := __e.Call(__defun__shen_4app, V3105, reg15985, reg15986)
			reg15988 := PrimStringConcat(reg15984, reg15987)
			reg15989 := __e.Call(__defun__y_1or_1n_2, reg15988)
			var reg15992 Obj
			if reg15989 == True {
				reg15990 := True
				reg15992 = reg15990
			} else {
				reg15991 := False
				reg15992 = reg15991
			}
			reg15994 = reg15992
		} else {
			reg15993 := False
			reg15994 = reg15993
		}
		var reg15998 Obj
		if reg15994 == True {
			reg15995 := __e.Call(__defun__ps, V3105)
			reg15996 := __e.Call(__defun__shen_4track_1function, reg15995)
			reg15998 = reg15996
		} else {
			reg15997 := MakeSymbol("shen.ok")
			reg15998 = reg15997
		}
		_ = reg15998
		reg15999 := MakeString("aborted")
		reg16000 := PrimSimpleError(reg15999)
		__ctx.Return(reg16000)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.f_error", value: __defun__shen_4f__error})

	__defun__shen_4tracked_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3107 := __args[0]
		_ = V3107
		reg16001 := MakeSymbol("shen.*tracking*")
		reg16002 := PrimValue(reg16001)
		__ctx.TailApply(__defun__element_2, V3107, reg16002)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.tracked?", value: __defun__shen_4tracked_2})

	__defun__track = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3109 := __args[0]
		_ = V3109
		reg16004 := __e.Call(__defun__ps, V3109)
		Source := reg16004
		_ = Source
		__ctx.TailApply(__defun__shen_4track_1function, Source)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "track", value: __defun__track})

	__defun__shen_4track_1function = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3111 := __args[0]
		_ = V3111
		reg16006 := PrimIsPair(V3111)
		var reg16049 Obj
		if reg16006 == True {
			reg16007 := MakeSymbol("defun")
			reg16008 := PrimHead(V3111)
			reg16009 := PrimEqual(reg16007, reg16008)
			var reg16044 Obj
			if reg16009 == True {
				reg16010 := PrimTail(V3111)
				reg16011 := PrimIsPair(reg16010)
				var reg16039 Obj
				if reg16011 == True {
					reg16012 := PrimTail(V3111)
					reg16013 := PrimTail(reg16012)
					reg16014 := PrimIsPair(reg16013)
					var reg16034 Obj
					if reg16014 == True {
						reg16015 := PrimTail(V3111)
						reg16016 := PrimTail(reg16015)
						reg16017 := PrimTail(reg16016)
						reg16018 := PrimIsPair(reg16017)
						var reg16029 Obj
						if reg16018 == True {
							reg16019 := Nil
							reg16020 := PrimTail(V3111)
							reg16021 := PrimTail(reg16020)
							reg16022 := PrimTail(reg16021)
							reg16023 := PrimTail(reg16022)
							reg16024 := PrimEqual(reg16019, reg16023)
							var reg16027 Obj
							if reg16024 == True {
								reg16025 := True
								reg16027 = reg16025
							} else {
								reg16026 := False
								reg16027 = reg16026
							}
							reg16029 = reg16027
						} else {
							reg16028 := False
							reg16029 = reg16028
						}
						var reg16032 Obj
						if reg16029 == True {
							reg16030 := True
							reg16032 = reg16030
						} else {
							reg16031 := False
							reg16032 = reg16031
						}
						reg16034 = reg16032
					} else {
						reg16033 := False
						reg16034 = reg16033
					}
					var reg16037 Obj
					if reg16034 == True {
						reg16035 := True
						reg16037 = reg16035
					} else {
						reg16036 := False
						reg16037 = reg16036
					}
					reg16039 = reg16037
				} else {
					reg16038 := False
					reg16039 = reg16038
				}
				var reg16042 Obj
				if reg16039 == True {
					reg16040 := True
					reg16042 = reg16040
				} else {
					reg16041 := False
					reg16042 = reg16041
				}
				reg16044 = reg16042
			} else {
				reg16043 := False
				reg16044 = reg16043
			}
			var reg16047 Obj
			if reg16044 == True {
				reg16045 := True
				reg16047 = reg16045
			} else {
				reg16046 := False
				reg16047 = reg16046
			}
			reg16049 = reg16047
		} else {
			reg16048 := False
			reg16049 = reg16048
		}
		if reg16049 == True {
			reg16050 := MakeSymbol("defun")
			reg16051 := PrimTail(V3111)
			reg16052 := PrimHead(reg16051)
			reg16053 := PrimTail(V3111)
			reg16054 := PrimTail(reg16053)
			reg16055 := PrimHead(reg16054)
			reg16056 := PrimTail(V3111)
			reg16057 := PrimHead(reg16056)
			reg16058 := PrimTail(V3111)
			reg16059 := PrimTail(reg16058)
			reg16060 := PrimHead(reg16059)
			reg16061 := PrimTail(V3111)
			reg16062 := PrimTail(reg16061)
			reg16063 := PrimTail(reg16062)
			reg16064 := PrimHead(reg16063)
			reg16065 := __e.Call(__defun__shen_4insert_1tracking_1code, reg16057, reg16060, reg16064)
			reg16066 := Nil
			reg16067 := PrimCons(reg16065, reg16066)
			reg16068 := PrimCons(reg16055, reg16067)
			reg16069 := PrimCons(reg16052, reg16068)
			reg16070 := PrimCons(reg16050, reg16069)
			KL := reg16070
			_ = KL
			reg16071 := PrimEvalKL(__e, KL)
			Ob := reg16071
			_ = Ob
			reg16072 := MakeSymbol("shen.*tracking*")
			reg16073 := MakeSymbol("shen.*tracking*")
			reg16074 := PrimValue(reg16073)
			reg16075 := PrimCons(Ob, reg16074)
			reg16076 := PrimSet(reg16072, reg16075)
			Tr := reg16076
			_ = Tr
			__ctx.Return(Ob)
			return
		} else {
			reg16077 := MakeSymbol("shen.track-function")
			__ctx.TailApply(__defun__shen_4f__error, reg16077)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.track-function", value: __defun__shen_4track_1function})

	__defun__shen_4insert_1tracking_1code = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3115 := __args[0]
		_ = V3115
		V3116 := __args[1]
		_ = V3116
		V3117 := __args[2]
		_ = V3117
		reg16079 := MakeSymbol("do")
		reg16080 := MakeSymbol("set")
		reg16081 := MakeSymbol("shen.*call*")
		reg16082 := MakeSymbol("+")
		reg16083 := MakeSymbol("value")
		reg16084 := MakeSymbol("shen.*call*")
		reg16085 := Nil
		reg16086 := PrimCons(reg16084, reg16085)
		reg16087 := PrimCons(reg16083, reg16086)
		reg16088 := MakeNumber(1)
		reg16089 := Nil
		reg16090 := PrimCons(reg16088, reg16089)
		reg16091 := PrimCons(reg16087, reg16090)
		reg16092 := PrimCons(reg16082, reg16091)
		reg16093 := Nil
		reg16094 := PrimCons(reg16092, reg16093)
		reg16095 := PrimCons(reg16081, reg16094)
		reg16096 := PrimCons(reg16080, reg16095)
		reg16097 := MakeSymbol("do")
		reg16098 := MakeSymbol("shen.input-track")
		reg16099 := MakeSymbol("value")
		reg16100 := MakeSymbol("shen.*call*")
		reg16101 := Nil
		reg16102 := PrimCons(reg16100, reg16101)
		reg16103 := PrimCons(reg16099, reg16102)
		reg16104 := __e.Call(__defun__shen_4cons__form, V3116)
		reg16105 := Nil
		reg16106 := PrimCons(reg16104, reg16105)
		reg16107 := PrimCons(V3115, reg16106)
		reg16108 := PrimCons(reg16103, reg16107)
		reg16109 := PrimCons(reg16098, reg16108)
		reg16110 := MakeSymbol("do")
		reg16111 := MakeSymbol("shen.terpri-or-read-char")
		reg16112 := Nil
		reg16113 := PrimCons(reg16111, reg16112)
		reg16114 := MakeSymbol("let")
		reg16115 := MakeSymbol("Result")
		reg16116 := MakeSymbol("do")
		reg16117 := MakeSymbol("shen.output-track")
		reg16118 := MakeSymbol("value")
		reg16119 := MakeSymbol("shen.*call*")
		reg16120 := Nil
		reg16121 := PrimCons(reg16119, reg16120)
		reg16122 := PrimCons(reg16118, reg16121)
		reg16123 := MakeSymbol("Result")
		reg16124 := Nil
		reg16125 := PrimCons(reg16123, reg16124)
		reg16126 := PrimCons(V3115, reg16125)
		reg16127 := PrimCons(reg16122, reg16126)
		reg16128 := PrimCons(reg16117, reg16127)
		reg16129 := MakeSymbol("do")
		reg16130 := MakeSymbol("set")
		reg16131 := MakeSymbol("shen.*call*")
		reg16132 := MakeSymbol("-")
		reg16133 := MakeSymbol("value")
		reg16134 := MakeSymbol("shen.*call*")
		reg16135 := Nil
		reg16136 := PrimCons(reg16134, reg16135)
		reg16137 := PrimCons(reg16133, reg16136)
		reg16138 := MakeNumber(1)
		reg16139 := Nil
		reg16140 := PrimCons(reg16138, reg16139)
		reg16141 := PrimCons(reg16137, reg16140)
		reg16142 := PrimCons(reg16132, reg16141)
		reg16143 := Nil
		reg16144 := PrimCons(reg16142, reg16143)
		reg16145 := PrimCons(reg16131, reg16144)
		reg16146 := PrimCons(reg16130, reg16145)
		reg16147 := MakeSymbol("do")
		reg16148 := MakeSymbol("shen.terpri-or-read-char")
		reg16149 := Nil
		reg16150 := PrimCons(reg16148, reg16149)
		reg16151 := MakeSymbol("Result")
		reg16152 := Nil
		reg16153 := PrimCons(reg16151, reg16152)
		reg16154 := PrimCons(reg16150, reg16153)
		reg16155 := PrimCons(reg16147, reg16154)
		reg16156 := Nil
		reg16157 := PrimCons(reg16155, reg16156)
		reg16158 := PrimCons(reg16146, reg16157)
		reg16159 := PrimCons(reg16129, reg16158)
		reg16160 := Nil
		reg16161 := PrimCons(reg16159, reg16160)
		reg16162 := PrimCons(reg16128, reg16161)
		reg16163 := PrimCons(reg16116, reg16162)
		reg16164 := Nil
		reg16165 := PrimCons(reg16163, reg16164)
		reg16166 := PrimCons(V3117, reg16165)
		reg16167 := PrimCons(reg16115, reg16166)
		reg16168 := PrimCons(reg16114, reg16167)
		reg16169 := Nil
		reg16170 := PrimCons(reg16168, reg16169)
		reg16171 := PrimCons(reg16113, reg16170)
		reg16172 := PrimCons(reg16110, reg16171)
		reg16173 := Nil
		reg16174 := PrimCons(reg16172, reg16173)
		reg16175 := PrimCons(reg16109, reg16174)
		reg16176 := PrimCons(reg16097, reg16175)
		reg16177 := Nil
		reg16178 := PrimCons(reg16176, reg16177)
		reg16179 := PrimCons(reg16096, reg16178)
		reg16180 := PrimCons(reg16079, reg16179)
		__ctx.Return(reg16180)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.insert-tracking-code", value: __defun__shen_4insert_1tracking_1code})

	__defun__step = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3123 := __args[0]
		_ = V3123
		reg16181 := MakeSymbol("+")
		reg16182 := PrimEqual(reg16181, V3123)
		if reg16182 == True {
			reg16183 := MakeSymbol("shen.*step*")
			reg16184 := True
			reg16185 := PrimSet(reg16183, reg16184)
			__ctx.Return(reg16185)
			return
		} else {
			reg16186 := MakeSymbol("-")
			reg16187 := PrimEqual(reg16186, V3123)
			if reg16187 == True {
				reg16188 := MakeSymbol("shen.*step*")
				reg16189 := False
				reg16190 := PrimSet(reg16188, reg16189)
				__ctx.Return(reg16190)
				return
			} else {
				reg16191 := MakeString("step expects a + or a -.\n")
				reg16192 := PrimSimpleError(reg16191)
				__ctx.Return(reg16192)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "step", value: __defun__step})

	__defun__spy = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3129 := __args[0]
		_ = V3129
		reg16193 := MakeSymbol("+")
		reg16194 := PrimEqual(reg16193, V3129)
		if reg16194 == True {
			reg16195 := MakeSymbol("shen.*spy*")
			reg16196 := True
			reg16197 := PrimSet(reg16195, reg16196)
			__ctx.Return(reg16197)
			return
		} else {
			reg16198 := MakeSymbol("-")
			reg16199 := PrimEqual(reg16198, V3129)
			if reg16199 == True {
				reg16200 := MakeSymbol("shen.*spy*")
				reg16201 := False
				reg16202 := PrimSet(reg16200, reg16201)
				__ctx.Return(reg16202)
				return
			} else {
				reg16203 := MakeString("spy expects a + or a -.\n")
				reg16204 := PrimSimpleError(reg16203)
				__ctx.Return(reg16204)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "spy", value: __defun__spy})

	__defun__shen_4terpri_1or_1read_1char = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg16205 := MakeSymbol("shen.*step*")
		reg16206 := PrimValue(reg16205)
		if reg16206 == True {
			reg16207 := MakeSymbol("*stinput*")
			reg16208 := PrimValue(reg16207)
			reg16209 := PrimReadByte(reg16208)
			__ctx.TailApply(__defun__shen_4check_1byte, reg16209)
			return
		} else {
			reg16211 := MakeNumber(1)
			__ctx.TailApply(__defun__nl, reg16211)
			return
		}
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.terpri-or-read-char", value: __defun__shen_4terpri_1or_1read_1char})

	__defun__shen_4check_1byte = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3135 := __args[0]
		_ = V3135
		reg16213 := __e.Call(__defun__shen_4hat)
		reg16214 := PrimEqual(V3135, reg16213)
		if reg16214 == True {
			reg16215 := MakeString("aborted")
			reg16216 := PrimSimpleError(reg16215)
			__ctx.Return(reg16216)
			return
		} else {
			reg16217 := True
			__ctx.Return(reg16217)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.check-byte", value: __defun__shen_4check_1byte})

	__defun__shen_4input_1track = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3139 := __args[0]
		_ = V3139
		V3140 := __args[1]
		_ = V3140
		V3141 := __args[2]
		_ = V3141
		reg16218 := MakeString("\n")
		reg16219 := __e.Call(__defun__shen_4spaces, V3139)
		reg16220 := MakeString("<")
		reg16221 := MakeString("> Inputs to ")
		reg16222 := MakeString(" \n")
		reg16223 := __e.Call(__defun__shen_4spaces, V3139)
		reg16224 := MakeString("")
		reg16225 := MakeSymbol("shen.a")
		reg16226 := __e.Call(__defun__shen_4app, reg16223, reg16224, reg16225)
		reg16227 := PrimStringConcat(reg16222, reg16226)
		reg16228 := MakeSymbol("shen.a")
		reg16229 := __e.Call(__defun__shen_4app, V3140, reg16227, reg16228)
		reg16230 := PrimStringConcat(reg16221, reg16229)
		reg16231 := MakeSymbol("shen.a")
		reg16232 := __e.Call(__defun__shen_4app, V3139, reg16230, reg16231)
		reg16233 := PrimStringConcat(reg16220, reg16232)
		reg16234 := MakeSymbol("shen.a")
		reg16235 := __e.Call(__defun__shen_4app, reg16219, reg16233, reg16234)
		reg16236 := PrimStringConcat(reg16218, reg16235)
		reg16237 := __e.Call(__defun__stoutput)
		reg16238 := __e.Call(__defun__shen_4prhush, reg16236, reg16237)
		_ = reg16238
		__ctx.TailApply(__defun__shen_4recursively_1print, V3141)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.input-track", value: __defun__shen_4input_1track})

	__defun__shen_4recursively_1print = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3143 := __args[0]
		_ = V3143
		reg16240 := Nil
		reg16241 := PrimEqual(reg16240, V3143)
		if reg16241 == True {
			reg16242 := MakeString(" ==>")
			reg16243 := __e.Call(__defun__stoutput)
			__ctx.TailApply(__defun__shen_4prhush, reg16242, reg16243)
			return
		} else {
			reg16245 := PrimIsPair(V3143)
			if reg16245 == True {
				reg16246 := PrimHead(V3143)
				reg16247 := __e.Call(__defun__print, reg16246)
				_ = reg16247
				reg16248 := MakeString(", ")
				reg16249 := __e.Call(__defun__stoutput)
				reg16250 := __e.Call(__defun__shen_4prhush, reg16248, reg16249)
				_ = reg16250
				reg16251 := PrimTail(V3143)
				__ctx.TailApply(__defun__shen_4recursively_1print, reg16251)
				return
			} else {
				reg16253 := MakeSymbol("shen.recursively-print")
				__ctx.TailApply(__defun__shen_4f__error, reg16253)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.recursively-print", value: __defun__shen_4recursively_1print})

	__defun__shen_4spaces = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3145 := __args[0]
		_ = V3145
		reg16255 := MakeNumber(0)
		reg16256 := PrimEqual(reg16255, V3145)
		if reg16256 == True {
			reg16257 := MakeString("")
			__ctx.Return(reg16257)
			return
		} else {
			reg16258 := MakeString(" ")
			reg16259 := MakeNumber(1)
			reg16260 := PrimNumberSubtract(V3145, reg16259)
			reg16261 := __e.Call(__defun__shen_4spaces, reg16260)
			reg16262 := PrimStringConcat(reg16258, reg16261)
			__ctx.Return(reg16262)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.spaces", value: __defun__shen_4spaces})

	__defun__shen_4output_1track = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3149 := __args[0]
		_ = V3149
		V3150 := __args[1]
		_ = V3150
		V3151 := __args[2]
		_ = V3151
		reg16263 := MakeString("\n")
		reg16264 := __e.Call(__defun__shen_4spaces, V3149)
		reg16265 := MakeString("<")
		reg16266 := MakeString("> Output from ")
		reg16267 := MakeString(" \n")
		reg16268 := __e.Call(__defun__shen_4spaces, V3149)
		reg16269 := MakeString("==> ")
		reg16270 := MakeString("")
		reg16271 := MakeSymbol("shen.s")
		reg16272 := __e.Call(__defun__shen_4app, V3151, reg16270, reg16271)
		reg16273 := PrimStringConcat(reg16269, reg16272)
		reg16274 := MakeSymbol("shen.a")
		reg16275 := __e.Call(__defun__shen_4app, reg16268, reg16273, reg16274)
		reg16276 := PrimStringConcat(reg16267, reg16275)
		reg16277 := MakeSymbol("shen.a")
		reg16278 := __e.Call(__defun__shen_4app, V3150, reg16276, reg16277)
		reg16279 := PrimStringConcat(reg16266, reg16278)
		reg16280 := MakeSymbol("shen.a")
		reg16281 := __e.Call(__defun__shen_4app, V3149, reg16279, reg16280)
		reg16282 := PrimStringConcat(reg16265, reg16281)
		reg16283 := MakeSymbol("shen.a")
		reg16284 := __e.Call(__defun__shen_4app, reg16264, reg16282, reg16283)
		reg16285 := PrimStringConcat(reg16263, reg16284)
		reg16286 := __e.Call(__defun__stoutput)
		__ctx.TailApply(__defun__shen_4prhush, reg16285, reg16286)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.output-track", value: __defun__shen_4output_1track})

	__defun__untrack = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3153 := __args[0]
		_ = V3153
		reg16288 := MakeSymbol("shen.*tracking*")
		reg16289 := PrimValue(reg16288)
		Tracking := reg16289
		_ = Tracking
		reg16290 := MakeSymbol("shen.*tracking*")
		reg16291 := __e.Call(__defun__remove, V3153, Tracking)
		reg16292 := PrimSet(reg16290, reg16291)
		Tracking = reg16292
		_ = Tracking
		reg16293 := __e.Call(__defun__ps, V3153)
		__ctx.TailApply(__defun__eval, reg16293)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "untrack", value: __defun__untrack})

	__defun__profile = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3155 := __args[0]
		_ = V3155
		reg16295 := __e.Call(__defun__ps, V3155)
		__ctx.TailApply(__defun__shen_4profile_1help, reg16295)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "profile", value: __defun__profile})

	__defun__shen_4profile_1help = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3161 := __args[0]
		_ = V3161
		reg16297 := PrimIsPair(V3161)
		var reg16340 Obj
		if reg16297 == True {
			reg16298 := MakeSymbol("defun")
			reg16299 := PrimHead(V3161)
			reg16300 := PrimEqual(reg16298, reg16299)
			var reg16335 Obj
			if reg16300 == True {
				reg16301 := PrimTail(V3161)
				reg16302 := PrimIsPair(reg16301)
				var reg16330 Obj
				if reg16302 == True {
					reg16303 := PrimTail(V3161)
					reg16304 := PrimTail(reg16303)
					reg16305 := PrimIsPair(reg16304)
					var reg16325 Obj
					if reg16305 == True {
						reg16306 := PrimTail(V3161)
						reg16307 := PrimTail(reg16306)
						reg16308 := PrimTail(reg16307)
						reg16309 := PrimIsPair(reg16308)
						var reg16320 Obj
						if reg16309 == True {
							reg16310 := Nil
							reg16311 := PrimTail(V3161)
							reg16312 := PrimTail(reg16311)
							reg16313 := PrimTail(reg16312)
							reg16314 := PrimTail(reg16313)
							reg16315 := PrimEqual(reg16310, reg16314)
							var reg16318 Obj
							if reg16315 == True {
								reg16316 := True
								reg16318 = reg16316
							} else {
								reg16317 := False
								reg16318 = reg16317
							}
							reg16320 = reg16318
						} else {
							reg16319 := False
							reg16320 = reg16319
						}
						var reg16323 Obj
						if reg16320 == True {
							reg16321 := True
							reg16323 = reg16321
						} else {
							reg16322 := False
							reg16323 = reg16322
						}
						reg16325 = reg16323
					} else {
						reg16324 := False
						reg16325 = reg16324
					}
					var reg16328 Obj
					if reg16325 == True {
						reg16326 := True
						reg16328 = reg16326
					} else {
						reg16327 := False
						reg16328 = reg16327
					}
					reg16330 = reg16328
				} else {
					reg16329 := False
					reg16330 = reg16329
				}
				var reg16333 Obj
				if reg16330 == True {
					reg16331 := True
					reg16333 = reg16331
				} else {
					reg16332 := False
					reg16333 = reg16332
				}
				reg16335 = reg16333
			} else {
				reg16334 := False
				reg16335 = reg16334
			}
			var reg16338 Obj
			if reg16335 == True {
				reg16336 := True
				reg16338 = reg16336
			} else {
				reg16337 := False
				reg16338 = reg16337
			}
			reg16340 = reg16338
		} else {
			reg16339 := False
			reg16340 = reg16339
		}
		if reg16340 == True {
			reg16341 := MakeSymbol("shen.f")
			reg16342 := __e.Call(__defun__gensym, reg16341)
			G := reg16342
			_ = G
			reg16343 := MakeSymbol("defun")
			reg16344 := PrimTail(V3161)
			reg16345 := PrimHead(reg16344)
			reg16346 := PrimTail(V3161)
			reg16347 := PrimTail(reg16346)
			reg16348 := PrimHead(reg16347)
			reg16349 := PrimTail(V3161)
			reg16350 := PrimHead(reg16349)
			reg16351 := PrimTail(V3161)
			reg16352 := PrimTail(reg16351)
			reg16353 := PrimHead(reg16352)
			reg16354 := PrimTail(V3161)
			reg16355 := PrimTail(reg16354)
			reg16356 := PrimHead(reg16355)
			reg16357 := PrimCons(G, reg16356)
			reg16358 := __e.Call(__defun__shen_4profile_1func, reg16350, reg16353, reg16357)
			reg16359 := Nil
			reg16360 := PrimCons(reg16358, reg16359)
			reg16361 := PrimCons(reg16348, reg16360)
			reg16362 := PrimCons(reg16345, reg16361)
			reg16363 := PrimCons(reg16343, reg16362)
			Profile := reg16363
			_ = Profile
			reg16364 := MakeSymbol("defun")
			reg16365 := PrimTail(V3161)
			reg16366 := PrimTail(reg16365)
			reg16367 := PrimHead(reg16366)
			reg16368 := PrimTail(V3161)
			reg16369 := PrimHead(reg16368)
			reg16370 := PrimTail(V3161)
			reg16371 := PrimTail(reg16370)
			reg16372 := PrimTail(reg16371)
			reg16373 := PrimHead(reg16372)
			reg16374 := __e.Call(__defun__subst, G, reg16369, reg16373)
			reg16375 := Nil
			reg16376 := PrimCons(reg16374, reg16375)
			reg16377 := PrimCons(reg16367, reg16376)
			reg16378 := PrimCons(G, reg16377)
			reg16379 := PrimCons(reg16364, reg16378)
			Def := reg16379
			_ = Def
			reg16380 := __e.Call(__defun__shen_4eval_1without_1macros, Profile)
			CompileProfile := reg16380
			_ = CompileProfile
			reg16381 := __e.Call(__defun__shen_4eval_1without_1macros, Def)
			CompileG := reg16381
			_ = CompileG
			reg16382 := PrimTail(V3161)
			reg16383 := PrimHead(reg16382)
			__ctx.Return(reg16383)
			return
		} else {
			reg16384 := MakeString("Cannot profile.\n")
			reg16385 := PrimSimpleError(reg16384)
			__ctx.Return(reg16385)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.profile-help", value: __defun__shen_4profile_1help})

	__defun__unprofile = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3163 := __args[0]
		_ = V3163
		__ctx.TailApply(__defun__untrack, V3163)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "unprofile", value: __defun__unprofile})

	__defun__shen_4profile_1func = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3167 := __args[0]
		_ = V3167
		V3168 := __args[1]
		_ = V3168
		V3169 := __args[2]
		_ = V3169
		reg16387 := MakeSymbol("let")
		reg16388 := MakeSymbol("Start")
		reg16389 := MakeSymbol("get-time")
		reg16390 := MakeSymbol("run")
		reg16391 := Nil
		reg16392 := PrimCons(reg16390, reg16391)
		reg16393 := PrimCons(reg16389, reg16392)
		reg16394 := MakeSymbol("let")
		reg16395 := MakeSymbol("Result")
		reg16396 := MakeSymbol("let")
		reg16397 := MakeSymbol("Finish")
		reg16398 := MakeSymbol("-")
		reg16399 := MakeSymbol("get-time")
		reg16400 := MakeSymbol("run")
		reg16401 := Nil
		reg16402 := PrimCons(reg16400, reg16401)
		reg16403 := PrimCons(reg16399, reg16402)
		reg16404 := MakeSymbol("Start")
		reg16405 := Nil
		reg16406 := PrimCons(reg16404, reg16405)
		reg16407 := PrimCons(reg16403, reg16406)
		reg16408 := PrimCons(reg16398, reg16407)
		reg16409 := MakeSymbol("let")
		reg16410 := MakeSymbol("Record")
		reg16411 := MakeSymbol("shen.put-profile")
		reg16412 := MakeSymbol("+")
		reg16413 := MakeSymbol("shen.get-profile")
		reg16414 := Nil
		reg16415 := PrimCons(V3167, reg16414)
		reg16416 := PrimCons(reg16413, reg16415)
		reg16417 := MakeSymbol("Finish")
		reg16418 := Nil
		reg16419 := PrimCons(reg16417, reg16418)
		reg16420 := PrimCons(reg16416, reg16419)
		reg16421 := PrimCons(reg16412, reg16420)
		reg16422 := Nil
		reg16423 := PrimCons(reg16421, reg16422)
		reg16424 := PrimCons(V3167, reg16423)
		reg16425 := PrimCons(reg16411, reg16424)
		reg16426 := MakeSymbol("Result")
		reg16427 := Nil
		reg16428 := PrimCons(reg16426, reg16427)
		reg16429 := PrimCons(reg16425, reg16428)
		reg16430 := PrimCons(reg16410, reg16429)
		reg16431 := PrimCons(reg16409, reg16430)
		reg16432 := Nil
		reg16433 := PrimCons(reg16431, reg16432)
		reg16434 := PrimCons(reg16408, reg16433)
		reg16435 := PrimCons(reg16397, reg16434)
		reg16436 := PrimCons(reg16396, reg16435)
		reg16437 := Nil
		reg16438 := PrimCons(reg16436, reg16437)
		reg16439 := PrimCons(V3169, reg16438)
		reg16440 := PrimCons(reg16395, reg16439)
		reg16441 := PrimCons(reg16394, reg16440)
		reg16442 := Nil
		reg16443 := PrimCons(reg16441, reg16442)
		reg16444 := PrimCons(reg16393, reg16443)
		reg16445 := PrimCons(reg16388, reg16444)
		reg16446 := PrimCons(reg16387, reg16445)
		__ctx.Return(reg16446)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.profile-func", value: __defun__shen_4profile_1func})

	__defun__profile_1results = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3171 := __args[0]
		_ = V3171
		reg16447 := __e.Call(__defun__shen_4get_1profile, V3171)
		Results := reg16447
		_ = Results
		reg16448 := MakeNumber(0)
		reg16449 := __e.Call(__defun__shen_4put_1profile, V3171, reg16448)
		Initialise := reg16449
		_ = Initialise
		__ctx.TailApply(__defun___8p, V3171, Results)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "profile-results", value: __defun__profile_1results})

	__defun__shen_4get_1profile = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3173 := __args[0]
		_ = V3173
		reg16451 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg16452 := MakeSymbol("profile")
			reg16453 := MakeSymbol("*property-vector*")
			reg16454 := PrimValue(reg16453)
			__ctx.TailApply(__defun__get, V3173, reg16452, reg16454)
			return
		}, 0)
		reg16456 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg16457 := MakeNumber(0)
			__ctx.Return(reg16457)
			return
		}, 1)
		reg16458 := __e.Try(reg16451).Catch(reg16456)
		__ctx.Return(reg16458)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.get-profile", value: __defun__shen_4get_1profile})

	__defun__shen_4put_1profile = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V3176 := __args[0]
		_ = V3176
		V3177 := __args[1]
		_ = V3177
		reg16459 := MakeSymbol("profile")
		reg16460 := MakeSymbol("*property-vector*")
		reg16461 := PrimValue(reg16460)
		__ctx.TailApply(__defun__put, V3176, reg16459, V3177, reg16461)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.put-profile", value: __defun__shen_4put_1profile})

}
