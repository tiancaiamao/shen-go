package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp2970 := MakeNative(func(__e *ControlFlow) {
		V451 := __e.Get(1)
		_ = V451
		tmp2986 := PrimEqual(Nil, V451)

		if True == tmp2986 {
			__e.Return(Nil)
			return
		} else {
			tmp2985 := PrimIsPair(V451)

			var ifres2981 Obj

			if True == tmp2985 {
				tmp2983 := PrimTail(V451)

				tmp2984 := PrimIsPair(tmp2983)

				var ifres2982 Obj

				if True == tmp2984 {
					ifres2982 = True

				} else {
					ifres2982 = False

				}

				ifres2981 = ifres2982

			} else {
				ifres2981 = False

			}

			if True == ifres2981 {
				tmp2973 := MakeNative(func(__e *ControlFlow) {
					DecArity := __e.Get(1)
					_ = DecArity
					tmp2974 := PrimTail(V451)

					tmp2975 := PrimTail(tmp2974)

					__e.TailApply(PrimNS2Value(symshen_4initialise__arity__table), tmp2975)
					return

				}, 1)

				tmp2976 := PrimHead(V451)

				tmp2977 := PrimTail(V451)

				tmp2978 := PrimHead(tmp2977)

				tmp2979 := PrimNS3Value(sym_dproperty_1vector_d)

				tmp2980 := Call(__e, PrimNS2Value(symput), tmp2976, symarity, tmp2978, tmp2979)

				__e.TailApply(tmp2973, tmp2980)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4initialise__arity__table)
				return
			}

		}

	}, 1)

	tmp2987 := Call(__e, PrimNS1Value(symns2_1set), symshen_4initialise__arity__table, tmp2970)

	_ = tmp2987

	tmp2988 := MakeNative(func(__e *ControlFlow) {
		V453 := __e.Get(1)
		_ = V453
		tmp2989 := MakeNative(func(__e *ControlFlow) {
			tmp2990 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V453, symarity, tmp2990)
			return

		}, 0)

		tmp2991 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(MakeNumber(-1))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp2989, tmp2991)
		return

	}, 1)

	tmp2992 := Call(__e, PrimNS1Value(symns2_1set), symarity, tmp2988)

	_ = tmp2992

	tmp2993 := MakeNative(func(__e *ControlFlow) {
		V455 := __e.Get(1)
		_ = V455
		tmp2994 := MakeNative(func(__e *ControlFlow) {
			Shen := __e.Get(1)
			_ = Shen
			tmp2995 := MakeNative(func(__e *ControlFlow) {
				External := __e.Get(1)
				_ = External
				tmp2996 := MakeNative(func(__e *ControlFlow) {
					Place := __e.Get(1)
					_ = Place
					__e.Return(V455)
					return
				}, 1)

				tmp2997 := Call(__e, PrimNS2Value(symadjoin), V455, External)

				tmp2998 := PrimNS3Value(sym_dproperty_1vector_d)

				tmp2999 := Call(__e, PrimNS2Value(symput), Shen, symshen_4external_1symbols, tmp2997, tmp2998)

				__e.TailApply(tmp2996, tmp2999)
				return

			}, 1)

			tmp3000 := PrimNS3Value(sym_dproperty_1vector_d)

			tmp3001 := Call(__e, PrimNS2Value(symget), Shen, symshen_4external_1symbols, tmp3000)

			__e.TailApply(tmp2995, tmp3001)
			return

		}, 1)

		tmp3002 := PrimIntern(MakeString("shen"))

		__e.TailApply(tmp2994, tmp3002)
		return

	}, 1)

	tmp3003 := Call(__e, PrimNS1Value(symns2_1set), symsystemf, tmp2993)

	_ = tmp3003

	tmp3004 := MakeNative(func(__e *ControlFlow) {
		V458 := __e.Get(1)
		_ = V458
		V459 := __e.Get(2)
		_ = V459
		tmp3006 := Call(__e, PrimNS2Value(symelement_2), V458, V459)

		if True == tmp3006 {
			__e.Return(V459)
			return
		} else {
			__e.Return(PrimCons(V458, V459))
			return
		}

	}, 2)

	tmp3007 := Call(__e, PrimNS1Value(symns2_1set), symadjoin, tmp3004)

	_ = tmp3007

	tmp3008 := MakeNative(func(__e *ControlFlow) {
		V461 := __e.Get(1)
		_ = V461
		tmp3021 := PrimEqual(sympackage, V461)

		if True == tmp3021 {
			__e.Return(Nil)
			return
		} else {
			tmp3020 := PrimEqual(symreceive, V461)

			if True == tmp3020 {
				__e.Return(Nil)
				return
			} else {
				tmp3011 := MakeNative(func(__e *ControlFlow) {
					ArityF := __e.Get(1)
					_ = ArityF
					tmp3018 := PrimEqual(ArityF, MakeNumber(-1))

					if True == tmp3018 {
						__e.Return(Nil)
						return
					} else {
						tmp3017 := PrimEqual(ArityF, MakeNumber(0))

						if True == tmp3017 {
							__e.Return(Nil)
							return
						} else {
							tmp3014 := Call(__e, PrimNS2Value(symshen_4lambda_1form), V461, ArityF)

							tmp3015 := Call(__e, PrimNS2Value(symeval_1kl), tmp3014)

							tmp3016 := PrimCons(V461, tmp3015)

							__e.Return(PrimCons(tmp3016, Nil))
							return

						}

					}

				}, 1)

				tmp3019 := Call(__e, PrimNS2Value(symarity), V461)

				__e.TailApply(tmp3011, tmp3019)
				return

			}

		}

	}, 1)

	tmp3022 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1form_1entry, tmp3008)

	_ = tmp3022

	tmp3023 := MakeNative(func(__e *ControlFlow) {
		V464 := __e.Get(1)
		_ = V464
		V465 := __e.Get(2)
		_ = V465
		tmp3032 := PrimEqual(MakeNumber(0), V465)

		if True == tmp3032 {
			__e.Return(V464)
			return
		} else {
			tmp3025 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp3026 := Call(__e, PrimNS2Value(symshen_4add_1end), V464, X)

				tmp3027 := PrimNumberSubtract(V465, MakeNumber(1))

				tmp3028 := Call(__e, PrimNS2Value(symshen_4lambda_1form), tmp3026, tmp3027)

				tmp3029 := PrimCons(tmp3028, Nil)

				tmp3030 := PrimCons(X, tmp3029)

				__e.Return(PrimCons(symlambda, tmp3030))
				return

			}, 1)

			tmp3031 := Call(__e, PrimNS2Value(symgensym), symV)

			__e.TailApply(tmp3025, tmp3031)
			return

		}

	}, 2)

	tmp3033 := Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1form, tmp3023)

	_ = tmp3033

	tmp3034 := MakeNative(func(__e *ControlFlow) {
		V468 := __e.Get(1)
		_ = V468
		V469 := __e.Get(2)
		_ = V469
		tmp3038 := PrimIsPair(V468)

		if True == tmp3038 {
			tmp3037 := PrimCons(V469, Nil)

			__e.TailApply(PrimNS2Value(symappend), V468, tmp3037)
			return

		} else {
			tmp3036 := PrimCons(V469, Nil)

			__e.Return(PrimCons(V468, tmp3036))
			return

		}

	}, 2)

	tmp3039 := Call(__e, PrimNS1Value(symns2_1set), symshen_4add_1end, tmp3034)

	_ = tmp3039

	tmp3040 := MakeNative(func(__e *ControlFlow) {
		V471 := __e.Get(1)
		_ = V471
		tmp3045 := PrimIsPair(V471)

		if True == tmp3045 {
			tmp3042 := PrimHead(V471)

			tmp3043 := PrimTail(V471)

			tmp3044 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symput), tmp3042, symshen_4lambda_1form, tmp3043, tmp3044)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4set_1lambda_1form_1entry)
			return
		}

	}, 1)

	tmp3046 := Call(__e, PrimNS1Value(symns2_1set), symshen_4set_1lambda_1form_1entry, tmp3040)

	_ = tmp3046

	tmp3047 := MakeNative(func(__e *ControlFlow) {
		V473 := __e.Get(1)
		_ = V473
		tmp3048 := PrimNS3Value(symshen_4_dspecial_d)

		tmp3049 := PrimCons(V473, tmp3048)

		tmp3050 := PrimNS3Set(symshen_4_dspecial_d, tmp3049)

		_ = tmp3050

		__e.Return(V473)
		return

	}, 1)

	tmp3051 := Call(__e, PrimNS1Value(symns2_1set), symspecialise, tmp3047)

	_ = tmp3051

	tmp3052 := MakeNative(func(__e *ControlFlow) {
		V475 := __e.Get(1)
		_ = V475
		tmp3053 := PrimNS3Value(symshen_4_dspecial_d)

		tmp3054 := Call(__e, PrimNS2Value(symremove), V475, tmp3053)

		tmp3055 := PrimNS3Set(symshen_4_dspecial_d, tmp3054)

		_ = tmp3055

		__e.Return(V475)
		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symunspecialise, tmp3052)
	return

}, 0)
