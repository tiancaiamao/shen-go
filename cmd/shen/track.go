package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp94138 := MakeNative(func(__e *ControlFlow) {
		V3135 := __e.Get(1)
		_ = V3135
		tmp94139 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp94140 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94141 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94142 := Call(__e, tmp94141, V3135, MakeString(";\n"), symshen_4a)

		tmp94143 := Call(__e, tmp94140, MakeString("partial function "), tmp94142)

		tmp94144 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp94145 := Call(__e, tmp94144)

		tmp94146 := Call(__e, tmp94139, tmp94143, tmp94145)

		_ = tmp94146

		tmp94160 := Call(__e, PrimNS1Value(symns2_1value), symnot)

		tmp94161 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tracked_2)

		tmp94162 := Call(__e, tmp94161, V3135)

		tmp94163 := Call(__e, tmp94160, tmp94162)

		var ifres94152 Obj

		if True == tmp94163 {
			tmp94154 := Call(__e, PrimNS1Value(symns2_1value), symy_1or_1n_2)

			tmp94155 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp94156 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			tmp94157 := Call(__e, tmp94156, V3135, MakeString("? "), symshen_4a)

			tmp94158 := Call(__e, tmp94155, MakeString("track "), tmp94157)

			tmp94159 := Call(__e, tmp94154, tmp94158)

			var ifres94153 Obj

			if True == tmp94159 {
				ifres94153 = True

			} else {
				ifres94153 = False

			}

			ifres94152 = ifres94153

		} else {
			ifres94152 = False

		}

		var ifres94147 Obj

		if True == ifres94152 {
			tmp94148 := Call(__e, PrimNS1Value(symns2_1value), symshen_4track_1function)

			tmp94149 := Call(__e, PrimNS1Value(symns2_1value), symps)

			tmp94150 := Call(__e, tmp94149, V3135)

			tmp94151 := Call(__e, tmp94148, tmp94150)

			ifres94147 = tmp94151

		} else {
			ifres94147 = symshen_4ok

		}

		_ = ifres94147

		tmp94164 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		__e.TailApply(tmp94164, MakeString("aborted"))
		return

	}, 1)

	tmp94165 := Call(__e, PrimNS1Value(symns2_1set), symshen_4f__error, tmp94138)

	_ = tmp94165

	tmp94166 := MakeNative(func(__e *ControlFlow) {
		V3137 := __e.Get(1)
		_ = V3137
		tmp94167 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		tmp94168 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp94169 := Call(__e, tmp94168, symshen_4_dtracking_d)

		__e.TailApply(tmp94167, V3137, tmp94169)
		return

	}, 1)

	tmp94170 := Call(__e, PrimNS1Value(symns2_1set), symshen_4tracked_2, tmp94166)

	_ = tmp94170

	tmp94171 := MakeNative(func(__e *ControlFlow) {
		V3139 := __e.Get(1)
		_ = V3139
		tmp94172 := MakeNative(func(__e *ControlFlow) {
			Source := __e.Get(1)
			_ = Source
			tmp94173 := Call(__e, PrimNS1Value(symns2_1value), symshen_4track_1function)

			__e.TailApply(tmp94173, Source)
			return

		}, 1)

		tmp94174 := Call(__e, PrimNS1Value(symns2_1value), symps)

		tmp94175 := Call(__e, tmp94174, V3139)

		__e.TailApply(tmp94172, tmp94175)
		return

	}, 1)

	tmp94176 := Call(__e, PrimNS1Value(symns2_1set), symtrack, tmp94171)

	_ = tmp94176

	tmp94177 := MakeNative(func(__e *ControlFlow) {
		V3141 := __e.Get(1)
		_ = V3141
		tmp94271 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp94272 := Call(__e, tmp94271, V3141)

		var ifres94229 Obj

		if True == tmp94272 {
			tmp94267 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp94268 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94269 := Call(__e, tmp94268, V3141)

			tmp94270 := Call(__e, tmp94267, symdefun, tmp94269)

			var ifres94231 Obj

			if True == tmp94270 {
				tmp94263 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp94264 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94265 := Call(__e, tmp94264, V3141)

				tmp94266 := Call(__e, tmp94263, tmp94265)

				var ifres94233 Obj

				if True == tmp94266 {
					tmp94257 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp94258 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94259 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94260 := Call(__e, tmp94259, V3141)

					tmp94261 := Call(__e, tmp94258, tmp94260)

					tmp94262 := Call(__e, tmp94257, tmp94261)

					var ifres94235 Obj

					if True == tmp94262 {
						tmp94249 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp94250 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94251 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94252 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94253 := Call(__e, tmp94252, V3141)

						tmp94254 := Call(__e, tmp94251, tmp94253)

						tmp94255 := Call(__e, tmp94250, tmp94254)

						tmp94256 := Call(__e, tmp94249, tmp94255)

						var ifres94237 Obj

						if True == tmp94256 {
							tmp94239 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp94240 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94241 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94242 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94243 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94244 := Call(__e, tmp94243, V3141)

							tmp94245 := Call(__e, tmp94242, tmp94244)

							tmp94246 := Call(__e, tmp94241, tmp94245)

							tmp94247 := Call(__e, tmp94240, tmp94246)

							tmp94248 := Call(__e, tmp94239, Nil, tmp94247)

							var ifres94238 Obj

							if True == tmp94248 {
								ifres94238 = True

							} else {
								ifres94238 = False

							}

							ifres94237 = ifres94238

						} else {
							ifres94237 = False

						}

						var ifres94236 Obj

						if True == ifres94237 {
							ifres94236 = True

						} else {
							ifres94236 = False

						}

						ifres94235 = ifres94236

					} else {
						ifres94235 = False

					}

					var ifres94234 Obj

					if True == ifres94235 {
						ifres94234 = True

					} else {
						ifres94234 = False

					}

					ifres94233 = ifres94234

				} else {
					ifres94233 = False

				}

				var ifres94232 Obj

				if True == ifres94233 {
					ifres94232 = True

				} else {
					ifres94232 = False

				}

				ifres94231 = ifres94232

			} else {
				ifres94231 = False

			}

			var ifres94230 Obj

			if True == ifres94231 {
				ifres94230 = True

			} else {
				ifres94230 = False

			}

			ifres94229 = ifres94230

		} else {
			ifres94229 = False

		}

		if True == ifres94229 {
			tmp94180 := MakeNative(func(__e *ControlFlow) {
				KL := __e.Get(1)
				_ = KL
				tmp94181 := MakeNative(func(__e *ControlFlow) {
					Ob := __e.Get(1)
					_ = Ob
					tmp94182 := MakeNative(func(__e *ControlFlow) {
						Tr := __e.Get(1)
						_ = Tr
						__e.Return(Ob)
						return
					}, 1)

					tmp94183 := Call(__e, PrimNS1Value(symns2_1value), symset)

					tmp94184 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp94185 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

					tmp94186 := Call(__e, tmp94185, symshen_4_dtracking_d)

					tmp94187 := Call(__e, tmp94184, Ob, tmp94186)

					tmp94188 := Call(__e, tmp94183, symshen_4_dtracking_d, tmp94187)

					__e.TailApply(tmp94182, tmp94188)
					return

				}, 1)

				tmp94189 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

				tmp94190 := Call(__e, tmp94189, KL)

				__e.TailApply(tmp94181, tmp94190)
				return

			}, 1)

			tmp94191 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp94192 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp94193 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94194 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94195 := Call(__e, tmp94194, V3141)

			tmp94196 := Call(__e, tmp94193, tmp94195)

			tmp94197 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp94198 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94199 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94200 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94201 := Call(__e, tmp94200, V3141)

			tmp94202 := Call(__e, tmp94199, tmp94201)

			tmp94203 := Call(__e, tmp94198, tmp94202)

			tmp94204 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			tmp94205 := Call(__e, PrimNS1Value(symns2_1value), symshen_4insert_1tracking_1code)

			tmp94206 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94207 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94208 := Call(__e, tmp94207, V3141)

			tmp94209 := Call(__e, tmp94206, tmp94208)

			tmp94210 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94211 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94212 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94213 := Call(__e, tmp94212, V3141)

			tmp94214 := Call(__e, tmp94211, tmp94213)

			tmp94215 := Call(__e, tmp94210, tmp94214)

			tmp94216 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94217 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94218 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94219 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			tmp94220 := Call(__e, tmp94219, V3141)

			tmp94221 := Call(__e, tmp94218, tmp94220)

			tmp94222 := Call(__e, tmp94217, tmp94221)

			tmp94223 := Call(__e, tmp94216, tmp94222)

			tmp94224 := Call(__e, tmp94205, tmp94209, tmp94215, tmp94223)

			tmp94225 := Call(__e, tmp94204, tmp94224, Nil)

			tmp94226 := Call(__e, tmp94197, tmp94203, tmp94225)

			tmp94227 := Call(__e, tmp94192, tmp94196, tmp94226)

			tmp94228 := Call(__e, tmp94191, symdefun, tmp94227)

			__e.TailApply(tmp94180, tmp94228)
			return

		} else {
			tmp94179 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

			__e.TailApply(tmp94179, symshen_4track_1function)
			return

		}

	}, 1)

	tmp94273 := Call(__e, PrimNS1Value(symns2_1set), symshen_4track_1function, tmp94177)

	_ = tmp94273

	tmp94274 := MakeNative(func(__e *ControlFlow) {
		V3145 := __e.Get(1)
		_ = V3145
		V3146 := __e.Get(2)
		_ = V3146
		V3147 := __e.Get(3)
		_ = V3147
		tmp94275 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94276 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94277 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94278 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94279 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94280 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94281 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94282 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94283 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94284 := Call(__e, tmp94283, symshen_4_dcall_d, Nil)

		tmp94285 := Call(__e, tmp94282, symvalue, tmp94284)

		tmp94286 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94287 := Call(__e, tmp94286, MakeNumber(1), Nil)

		tmp94288 := Call(__e, tmp94281, tmp94285, tmp94287)

		tmp94289 := Call(__e, tmp94280, sym_7, tmp94288)

		tmp94290 := Call(__e, tmp94279, tmp94289, Nil)

		tmp94291 := Call(__e, tmp94278, symshen_4_dcall_d, tmp94290)

		tmp94292 := Call(__e, tmp94277, symset, tmp94291)

		tmp94293 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94294 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94295 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94296 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94297 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94298 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94299 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94300 := Call(__e, tmp94299, symshen_4_dcall_d, Nil)

		tmp94301 := Call(__e, tmp94298, symvalue, tmp94300)

		tmp94302 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94303 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94304 := Call(__e, PrimNS1Value(symns2_1value), symshen_4cons__form)

		tmp94305 := Call(__e, tmp94304, V3146)

		tmp94306 := Call(__e, tmp94303, tmp94305, Nil)

		tmp94307 := Call(__e, tmp94302, V3145, tmp94306)

		tmp94308 := Call(__e, tmp94297, tmp94301, tmp94307)

		tmp94309 := Call(__e, tmp94296, symshen_4input_1track, tmp94308)

		tmp94310 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94311 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94312 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94313 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94314 := Call(__e, tmp94313, symshen_4terpri_1or_1read_1char, Nil)

		tmp94315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94317 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94326 := Call(__e, tmp94325, symshen_4_dcall_d, Nil)

		tmp94327 := Call(__e, tmp94324, symvalue, tmp94326)

		tmp94328 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94329 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94330 := Call(__e, tmp94329, symResult, Nil)

		tmp94331 := Call(__e, tmp94328, V3145, tmp94330)

		tmp94332 := Call(__e, tmp94323, tmp94327, tmp94331)

		tmp94333 := Call(__e, tmp94322, symshen_4output_1track, tmp94332)

		tmp94334 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94335 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94336 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94337 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94338 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94339 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94340 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94341 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94342 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94343 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94344 := Call(__e, tmp94343, symshen_4_dcall_d, Nil)

		tmp94345 := Call(__e, tmp94342, symvalue, tmp94344)

		tmp94346 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94347 := Call(__e, tmp94346, MakeNumber(1), Nil)

		tmp94348 := Call(__e, tmp94341, tmp94345, tmp94347)

		tmp94349 := Call(__e, tmp94340, sym_1, tmp94348)

		tmp94350 := Call(__e, tmp94339, tmp94349, Nil)

		tmp94351 := Call(__e, tmp94338, symshen_4_dcall_d, tmp94350)

		tmp94352 := Call(__e, tmp94337, symset, tmp94351)

		tmp94353 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94354 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94355 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94356 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94357 := Call(__e, tmp94356, symshen_4terpri_1or_1read_1char, Nil)

		tmp94358 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94359 := Call(__e, tmp94358, symResult, Nil)

		tmp94360 := Call(__e, tmp94355, tmp94357, tmp94359)

		tmp94361 := Call(__e, tmp94354, symdo, tmp94360)

		tmp94362 := Call(__e, tmp94353, tmp94361, Nil)

		tmp94363 := Call(__e, tmp94336, tmp94352, tmp94362)

		tmp94364 := Call(__e, tmp94335, symdo, tmp94363)

		tmp94365 := Call(__e, tmp94334, tmp94364, Nil)

		tmp94366 := Call(__e, tmp94321, tmp94333, tmp94365)

		tmp94367 := Call(__e, tmp94320, symdo, tmp94366)

		tmp94368 := Call(__e, tmp94319, tmp94367, Nil)

		tmp94369 := Call(__e, tmp94318, V3147, tmp94368)

		tmp94370 := Call(__e, tmp94317, symResult, tmp94369)

		tmp94371 := Call(__e, tmp94316, symlet, tmp94370)

		tmp94372 := Call(__e, tmp94315, tmp94371, Nil)

		tmp94373 := Call(__e, tmp94312, tmp94314, tmp94372)

		tmp94374 := Call(__e, tmp94311, symdo, tmp94373)

		tmp94375 := Call(__e, tmp94310, tmp94374, Nil)

		tmp94376 := Call(__e, tmp94295, tmp94309, tmp94375)

		tmp94377 := Call(__e, tmp94294, symdo, tmp94376)

		tmp94378 := Call(__e, tmp94293, tmp94377, Nil)

		tmp94379 := Call(__e, tmp94276, tmp94292, tmp94378)

		__e.TailApply(tmp94275, symdo, tmp94379)
		return

	}, 3)

	tmp94380 := Call(__e, PrimNS1Value(symns2_1set), symshen_4insert_1tracking_1code, tmp94274)

	_ = tmp94380

	tmp94381 := MakeNative(func(__e *ControlFlow) {
		V3153 := __e.Get(1)
		_ = V3153
		tmp94389 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp94390 := Call(__e, tmp94389, sym_7, V3153)

		if True == tmp94390 {
			tmp94388 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(tmp94388, symshen_4_dstep_d, True)
			return

		} else {
			tmp94386 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp94387 := Call(__e, tmp94386, sym_1, V3153)

			if True == tmp94387 {
				tmp94385 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(tmp94385, symshen_4_dstep_d, False)
				return

			} else {
				tmp94384 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp94384, MakeString("step expects a + or a -.\n"))
				return

			}

		}

	}, 1)

	tmp94391 := Call(__e, PrimNS1Value(symns2_1set), symstep, tmp94381)

	_ = tmp94391

	tmp94392 := MakeNative(func(__e *ControlFlow) {
		V3159 := __e.Get(1)
		_ = V3159
		tmp94400 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp94401 := Call(__e, tmp94400, sym_7, V3159)

		if True == tmp94401 {
			tmp94399 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(tmp94399, symshen_4_dspy_d, True)
			return

		} else {
			tmp94397 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp94398 := Call(__e, tmp94397, sym_1, V3159)

			if True == tmp94398 {
				tmp94396 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(tmp94396, symshen_4_dspy_d, False)
				return

			} else {
				tmp94395 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(tmp94395, MakeString("spy expects a + or a -.\n"))
				return

			}

		}

	}, 1)

	tmp94402 := Call(__e, PrimNS1Value(symns2_1set), symspy, tmp94392)

	_ = tmp94402

	tmp94403 := MakeNative(func(__e *ControlFlow) {
		tmp94411 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp94412 := Call(__e, tmp94411, symshen_4_dstep_d)

		if True == tmp94412 {
			tmp94406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4check_1byte)

			tmp94407 := Call(__e, PrimNS1Value(symns2_1value), symread_1byte)

			tmp94408 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp94409 := Call(__e, tmp94408, sym_dstinput_d)

			tmp94410 := Call(__e, tmp94407, tmp94409)

			__e.TailApply(tmp94406, tmp94410)
			return

		} else {
			tmp94405 := Call(__e, PrimNS1Value(symns2_1value), symnl)

			__e.TailApply(tmp94405, MakeNumber(1))
			return

		}

	}, 0)

	tmp94413 := Call(__e, PrimNS1Value(symns2_1set), symshen_4terpri_1or_1read_1char, tmp94403)

	_ = tmp94413

	tmp94414 := MakeNative(func(__e *ControlFlow) {
		V3165 := __e.Get(1)
		_ = V3165
		tmp94417 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp94418 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hat)

		tmp94419 := Call(__e, tmp94418)

		tmp94420 := Call(__e, tmp94417, V3165, tmp94419)

		if True == tmp94420 {
			tmp94416 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp94416, MakeString("aborted"))
			return

		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp94421 := Call(__e, PrimNS1Value(symns2_1set), symshen_4check_1byte, tmp94414)

	_ = tmp94421

	tmp94422 := MakeNative(func(__e *ControlFlow) {
		V3169 := __e.Get(1)
		_ = V3169
		V3170 := __e.Get(2)
		_ = V3170
		V3171 := __e.Get(3)
		_ = V3171
		tmp94423 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp94424 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94425 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94426 := Call(__e, PrimNS1Value(symns2_1value), symshen_4spaces)

		tmp94427 := Call(__e, tmp94426, V3169)

		tmp94428 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94429 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94430 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94431 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94432 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94433 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94434 := Call(__e, PrimNS1Value(symns2_1value), symshen_4spaces)

		tmp94435 := Call(__e, tmp94434, V3169)

		tmp94436 := Call(__e, tmp94433, tmp94435, MakeString(""), symshen_4a)

		tmp94437 := Call(__e, tmp94432, MakeString(" \n"), tmp94436)

		tmp94438 := Call(__e, tmp94431, V3170, tmp94437, symshen_4a)

		tmp94439 := Call(__e, tmp94430, MakeString("> Inputs to "), tmp94438)

		tmp94440 := Call(__e, tmp94429, V3169, tmp94439, symshen_4a)

		tmp94441 := Call(__e, tmp94428, MakeString("<"), tmp94440)

		tmp94442 := Call(__e, tmp94425, tmp94427, tmp94441, symshen_4a)

		tmp94443 := Call(__e, tmp94424, MakeString("\n"), tmp94442)

		tmp94444 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp94445 := Call(__e, tmp94444)

		tmp94446 := Call(__e, tmp94423, tmp94443, tmp94445)

		_ = tmp94446

		tmp94447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursively_1print)

		__e.TailApply(tmp94447, V3171)
		return

	}, 3)

	tmp94448 := Call(__e, PrimNS1Value(symns2_1set), symshen_4input_1track, tmp94422)

	_ = tmp94448

	tmp94449 := MakeNative(func(__e *ControlFlow) {
		V3173 := __e.Get(1)
		_ = V3173
		tmp94469 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp94470 := Call(__e, tmp94469, Nil, V3173)

		if True == tmp94470 {
			tmp94466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			tmp94467 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			tmp94468 := Call(__e, tmp94467)

			__e.TailApply(tmp94466, MakeString(" ==>"), tmp94468)
			return

		} else {
			tmp94464 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			tmp94465 := Call(__e, tmp94464, V3173)

			if True == tmp94465 {
				tmp94453 := Call(__e, PrimNS1Value(symns2_1value), symprint)

				tmp94454 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94455 := Call(__e, tmp94454, V3173)

				tmp94456 := Call(__e, tmp94453, tmp94455)

				_ = tmp94456

				tmp94457 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				tmp94458 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				tmp94459 := Call(__e, tmp94458)

				tmp94460 := Call(__e, tmp94457, MakeString(", "), tmp94459)

				_ = tmp94460

				tmp94461 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursively_1print)

				tmp94462 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94463 := Call(__e, tmp94462, V3173)

				__e.TailApply(tmp94461, tmp94463)
				return

			} else {
				tmp94452 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(tmp94452, symshen_4recursively_1print)
				return

			}

		}

	}, 1)

	tmp94471 := Call(__e, PrimNS1Value(symns2_1set), symshen_4recursively_1print, tmp94449)

	_ = tmp94471

	tmp94472 := MakeNative(func(__e *ControlFlow) {
		V3175 := __e.Get(1)
		_ = V3175
		tmp94479 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		tmp94480 := Call(__e, tmp94479, MakeNumber(0), V3175)

		if True == tmp94480 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp94474 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			tmp94475 := Call(__e, PrimNS1Value(symns2_1value), symshen_4spaces)

			tmp94476 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			tmp94477 := Call(__e, tmp94476, V3175, MakeNumber(1))

			tmp94478 := Call(__e, tmp94475, tmp94477)

			__e.TailApply(tmp94474, MakeString(" "), tmp94478)
			return

		}

	}, 1)

	tmp94481 := Call(__e, PrimNS1Value(symns2_1set), symshen_4spaces, tmp94472)

	_ = tmp94481

	tmp94482 := MakeNative(func(__e *ControlFlow) {
		V3179 := __e.Get(1)
		_ = V3179
		V3180 := __e.Get(2)
		_ = V3180
		V3181 := __e.Get(3)
		_ = V3181
		tmp94483 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		tmp94484 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94485 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94486 := Call(__e, PrimNS1Value(symns2_1value), symshen_4spaces)

		tmp94487 := Call(__e, tmp94486, V3179)

		tmp94488 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94490 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94492 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94493 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94494 := Call(__e, PrimNS1Value(symns2_1value), symshen_4spaces)

		tmp94495 := Call(__e, tmp94494, V3179)

		tmp94496 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		tmp94497 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		tmp94498 := Call(__e, tmp94497, V3181, MakeString(""), symshen_4s)

		tmp94499 := Call(__e, tmp94496, MakeString("==> "), tmp94498)

		tmp94500 := Call(__e, tmp94493, tmp94495, tmp94499, symshen_4a)

		tmp94501 := Call(__e, tmp94492, MakeString(" \n"), tmp94500)

		tmp94502 := Call(__e, tmp94491, V3180, tmp94501, symshen_4a)

		tmp94503 := Call(__e, tmp94490, MakeString("> Output from "), tmp94502)

		tmp94504 := Call(__e, tmp94489, V3179, tmp94503, symshen_4a)

		tmp94505 := Call(__e, tmp94488, MakeString("<"), tmp94504)

		tmp94506 := Call(__e, tmp94485, tmp94487, tmp94505, symshen_4a)

		tmp94507 := Call(__e, tmp94484, MakeString("\n"), tmp94506)

		tmp94508 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		tmp94509 := Call(__e, tmp94508)

		__e.TailApply(tmp94483, tmp94507, tmp94509)
		return

	}, 3)

	tmp94510 := Call(__e, PrimNS1Value(symns2_1set), symshen_4output_1track, tmp94482)

	_ = tmp94510

	tmp94511 := MakeNative(func(__e *ControlFlow) {
		V3183 := __e.Get(1)
		_ = V3183
		tmp94512 := MakeNative(func(__e *ControlFlow) {
			Tracking := __e.Get(1)
			_ = Tracking
			tmp94513 := MakeNative(func(__e *ControlFlow) {
				Tracking := __e.Get(1)
				_ = Tracking
				tmp94514 := Call(__e, PrimNS1Value(symns2_1value), symeval)

				tmp94515 := Call(__e, PrimNS1Value(symns2_1value), symps)

				tmp94516 := Call(__e, tmp94515, V3183)

				__e.TailApply(tmp94514, tmp94516)
				return

			}, 1)

			tmp94517 := Call(__e, PrimNS1Value(symns2_1value), symset)

			tmp94518 := Call(__e, PrimNS1Value(symns2_1value), symremove)

			tmp94519 := Call(__e, tmp94518, V3183, Tracking)

			tmp94520 := Call(__e, tmp94517, symshen_4_dtracking_d, tmp94519)

			__e.TailApply(tmp94513, tmp94520)
			return

		}, 1)

		tmp94521 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp94522 := Call(__e, tmp94521, symshen_4_dtracking_d)

		__e.TailApply(tmp94512, tmp94522)
		return

	}, 1)

	tmp94523 := Call(__e, PrimNS1Value(symns2_1set), symuntrack, tmp94511)

	_ = tmp94523

	tmp94524 := MakeNative(func(__e *ControlFlow) {
		V3185 := __e.Get(1)
		_ = V3185
		tmp94525 := Call(__e, PrimNS1Value(symns2_1value), symshen_4profile_1help)

		tmp94526 := Call(__e, PrimNS1Value(symns2_1value), symps)

		tmp94527 := Call(__e, tmp94526, V3185)

		__e.TailApply(tmp94525, tmp94527)
		return

	}, 1)

	tmp94528 := Call(__e, PrimNS1Value(symns2_1set), symprofile, tmp94524)

	_ = tmp94528

	tmp94529 := MakeNative(func(__e *ControlFlow) {
		V3191 := __e.Get(1)
		_ = V3191
		tmp94654 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		tmp94655 := Call(__e, tmp94654, V3191)

		var ifres94612 Obj

		if True == tmp94655 {
			tmp94650 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			tmp94651 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			tmp94652 := Call(__e, tmp94651, V3191)

			tmp94653 := Call(__e, tmp94650, symdefun, tmp94652)

			var ifres94614 Obj

			if True == tmp94653 {
				tmp94646 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				tmp94647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94648 := Call(__e, tmp94647, V3191)

				tmp94649 := Call(__e, tmp94646, tmp94648)

				var ifres94616 Obj

				if True == tmp94649 {
					tmp94640 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					tmp94641 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94642 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94643 := Call(__e, tmp94642, V3191)

					tmp94644 := Call(__e, tmp94641, tmp94643)

					tmp94645 := Call(__e, tmp94640, tmp94644)

					var ifres94618 Obj

					if True == tmp94645 {
						tmp94632 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						tmp94633 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94634 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94635 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						tmp94636 := Call(__e, tmp94635, V3191)

						tmp94637 := Call(__e, tmp94634, tmp94636)

						tmp94638 := Call(__e, tmp94633, tmp94637)

						tmp94639 := Call(__e, tmp94632, tmp94638)

						var ifres94620 Obj

						if True == tmp94639 {
							tmp94622 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							tmp94623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94624 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94625 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94626 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							tmp94627 := Call(__e, tmp94626, V3191)

							tmp94628 := Call(__e, tmp94625, tmp94627)

							tmp94629 := Call(__e, tmp94624, tmp94628)

							tmp94630 := Call(__e, tmp94623, tmp94629)

							tmp94631 := Call(__e, tmp94622, Nil, tmp94630)

							var ifres94621 Obj

							if True == tmp94631 {
								ifres94621 = True

							} else {
								ifres94621 = False

							}

							ifres94620 = ifres94621

						} else {
							ifres94620 = False

						}

						var ifres94619 Obj

						if True == ifres94620 {
							ifres94619 = True

						} else {
							ifres94619 = False

						}

						ifres94618 = ifres94619

					} else {
						ifres94618 = False

					}

					var ifres94617 Obj

					if True == ifres94618 {
						ifres94617 = True

					} else {
						ifres94617 = False

					}

					ifres94616 = ifres94617

				} else {
					ifres94616 = False

				}

				var ifres94615 Obj

				if True == ifres94616 {
					ifres94615 = True

				} else {
					ifres94615 = False

				}

				ifres94614 = ifres94615

			} else {
				ifres94614 = False

			}

			var ifres94613 Obj

			if True == ifres94614 {
				ifres94613 = True

			} else {
				ifres94613 = False

			}

			ifres94612 = ifres94613

		} else {
			ifres94612 = False

		}

		if True == ifres94612 {
			tmp94532 := MakeNative(func(__e *ControlFlow) {
				G := __e.Get(1)
				_ = G
				tmp94533 := MakeNative(func(__e *ControlFlow) {
					Profile := __e.Get(1)
					_ = Profile
					tmp94534 := MakeNative(func(__e *ControlFlow) {
						Def := __e.Get(1)
						_ = Def
						tmp94535 := MakeNative(func(__e *ControlFlow) {
							CompileProfile := __e.Get(1)
							_ = CompileProfile
							tmp94536 := MakeNative(func(__e *ControlFlow) {
								CompileG := __e.Get(1)
								_ = CompileG
								tmp94537 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								tmp94538 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								tmp94539 := Call(__e, tmp94538, V3191)

								__e.TailApply(tmp94537, tmp94539)
								return

							}, 1)

							tmp94540 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

							tmp94541 := Call(__e, tmp94540, Def)

							__e.TailApply(tmp94536, tmp94541)
							return

						}, 1)

						tmp94542 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

						tmp94543 := Call(__e, tmp94542, Profile)

						__e.TailApply(tmp94535, tmp94543)
						return

					}, 1)

					tmp94544 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp94545 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp94546 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp94547 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94548 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94549 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94550 := Call(__e, tmp94549, V3191)

					tmp94551 := Call(__e, tmp94548, tmp94550)

					tmp94552 := Call(__e, tmp94547, tmp94551)

					tmp94553 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					tmp94554 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					tmp94555 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94556 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94557 := Call(__e, tmp94556, V3191)

					tmp94558 := Call(__e, tmp94555, tmp94557)

					tmp94559 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					tmp94560 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94561 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94562 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					tmp94563 := Call(__e, tmp94562, V3191)

					tmp94564 := Call(__e, tmp94561, tmp94563)

					tmp94565 := Call(__e, tmp94560, tmp94564)

					tmp94566 := Call(__e, tmp94559, tmp94565)

					tmp94567 := Call(__e, tmp94554, G, tmp94558, tmp94566)

					tmp94568 := Call(__e, tmp94553, tmp94567, Nil)

					tmp94569 := Call(__e, tmp94546, tmp94552, tmp94568)

					tmp94570 := Call(__e, tmp94545, G, tmp94569)

					tmp94571 := Call(__e, tmp94544, symdefun, tmp94570)

					__e.TailApply(tmp94534, tmp94571)
					return

				}, 1)

				tmp94572 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp94573 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp94574 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94575 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94576 := Call(__e, tmp94575, V3191)

				tmp94577 := Call(__e, tmp94574, tmp94576)

				tmp94578 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp94579 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94580 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94581 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94582 := Call(__e, tmp94581, V3191)

				tmp94583 := Call(__e, tmp94580, tmp94582)

				tmp94584 := Call(__e, tmp94579, tmp94583)

				tmp94585 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp94586 := Call(__e, PrimNS1Value(symns2_1value), symshen_4profile_1func)

				tmp94587 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94588 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94589 := Call(__e, tmp94588, V3191)

				tmp94590 := Call(__e, tmp94587, tmp94589)

				tmp94591 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94592 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94593 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94594 := Call(__e, tmp94593, V3191)

				tmp94595 := Call(__e, tmp94592, tmp94594)

				tmp94596 := Call(__e, tmp94591, tmp94595)

				tmp94597 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				tmp94598 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				tmp94599 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94600 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				tmp94601 := Call(__e, tmp94600, V3191)

				tmp94602 := Call(__e, tmp94599, tmp94601)

				tmp94603 := Call(__e, tmp94598, tmp94602)

				tmp94604 := Call(__e, tmp94597, G, tmp94603)

				tmp94605 := Call(__e, tmp94586, tmp94590, tmp94596, tmp94604)

				tmp94606 := Call(__e, tmp94585, tmp94605, Nil)

				tmp94607 := Call(__e, tmp94578, tmp94584, tmp94606)

				tmp94608 := Call(__e, tmp94573, tmp94577, tmp94607)

				tmp94609 := Call(__e, tmp94572, symdefun, tmp94608)

				__e.TailApply(tmp94533, tmp94609)
				return

			}, 1)

			tmp94610 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			tmp94611 := Call(__e, tmp94610, symshen_4f)

			__e.TailApply(tmp94532, tmp94611)
			return

		} else {
			tmp94531 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(tmp94531, MakeString("Cannot profile.\n"))
			return

		}

	}, 1)

	tmp94656 := Call(__e, PrimNS1Value(symns2_1set), symshen_4profile_1help, tmp94529)

	_ = tmp94656

	tmp94657 := MakeNative(func(__e *ControlFlow) {
		V3193 := __e.Get(1)
		_ = V3193
		tmp94658 := Call(__e, PrimNS1Value(symns2_1value), symuntrack)

		__e.TailApply(tmp94658, V3193)
		return

	}, 1)

	tmp94659 := Call(__e, PrimNS1Value(symns2_1set), symunprofile, tmp94657)

	_ = tmp94659

	tmp94660 := MakeNative(func(__e *ControlFlow) {
		V3197 := __e.Get(1)
		_ = V3197
		V3198 := __e.Get(2)
		_ = V3198
		V3199 := __e.Get(3)
		_ = V3199
		tmp94661 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94662 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94663 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94664 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94665 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94666 := Call(__e, tmp94665, symrun, Nil)

		tmp94667 := Call(__e, tmp94664, symget_1time, tmp94666)

		tmp94668 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94669 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94670 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94671 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94672 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94673 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94674 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94675 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94676 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94677 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94678 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94679 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94680 := Call(__e, tmp94679, symrun, Nil)

		tmp94681 := Call(__e, tmp94678, symget_1time, tmp94680)

		tmp94682 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94683 := Call(__e, tmp94682, symStart, Nil)

		tmp94684 := Call(__e, tmp94677, tmp94681, tmp94683)

		tmp94685 := Call(__e, tmp94676, sym_1, tmp94684)

		tmp94686 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94687 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94688 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94689 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94690 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94691 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94692 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94693 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94694 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94695 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94696 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94697 := Call(__e, tmp94696, V3197, Nil)

		tmp94698 := Call(__e, tmp94695, symshen_4get_1profile, tmp94697)

		tmp94699 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94700 := Call(__e, tmp94699, symFinish, Nil)

		tmp94701 := Call(__e, tmp94694, tmp94698, tmp94700)

		tmp94702 := Call(__e, tmp94693, sym_7, tmp94701)

		tmp94703 := Call(__e, tmp94692, tmp94702, Nil)

		tmp94704 := Call(__e, tmp94691, V3197, tmp94703)

		tmp94705 := Call(__e, tmp94690, symshen_4put_1profile, tmp94704)

		tmp94706 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		tmp94707 := Call(__e, tmp94706, symResult, Nil)

		tmp94708 := Call(__e, tmp94689, tmp94705, tmp94707)

		tmp94709 := Call(__e, tmp94688, symRecord, tmp94708)

		tmp94710 := Call(__e, tmp94687, symlet, tmp94709)

		tmp94711 := Call(__e, tmp94686, tmp94710, Nil)

		tmp94712 := Call(__e, tmp94675, tmp94685, tmp94711)

		tmp94713 := Call(__e, tmp94674, symFinish, tmp94712)

		tmp94714 := Call(__e, tmp94673, symlet, tmp94713)

		tmp94715 := Call(__e, tmp94672, tmp94714, Nil)

		tmp94716 := Call(__e, tmp94671, V3199, tmp94715)

		tmp94717 := Call(__e, tmp94670, symResult, tmp94716)

		tmp94718 := Call(__e, tmp94669, symlet, tmp94717)

		tmp94719 := Call(__e, tmp94668, tmp94718, Nil)

		tmp94720 := Call(__e, tmp94663, tmp94667, tmp94719)

		tmp94721 := Call(__e, tmp94662, symStart, tmp94720)

		__e.TailApply(tmp94661, symlet, tmp94721)
		return

	}, 3)

	tmp94722 := Call(__e, PrimNS1Value(symns2_1set), symshen_4profile_1func, tmp94660)

	_ = tmp94722

	tmp94723 := MakeNative(func(__e *ControlFlow) {
		V3201 := __e.Get(1)
		_ = V3201
		tmp94724 := MakeNative(func(__e *ControlFlow) {
			Results := __e.Get(1)
			_ = Results
			tmp94725 := MakeNative(func(__e *ControlFlow) {
				Initialise := __e.Get(1)
				_ = Initialise
				tmp94726 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

				__e.TailApply(tmp94726, V3201, Results)
				return

			}, 1)

			tmp94727 := Call(__e, PrimNS1Value(symns2_1value), symshen_4put_1profile)

			tmp94728 := Call(__e, tmp94727, V3201, MakeNumber(0))

			__e.TailApply(tmp94725, tmp94728)
			return

		}, 1)

		tmp94729 := Call(__e, PrimNS1Value(symns2_1value), symshen_4get_1profile)

		tmp94730 := Call(__e, tmp94729, V3201)

		__e.TailApply(tmp94724, tmp94730)
		return

	}, 1)

	tmp94731 := Call(__e, PrimNS1Value(symns2_1set), symprofile_1results, tmp94723)

	_ = tmp94731

	tmp94732 := MakeNative(func(__e *ControlFlow) {
		V3203 := __e.Get(1)
		_ = V3203
		tmp94733 := MakeNative(func(__e *ControlFlow) {
			tmp94734 := Call(__e, PrimNS1Value(symns2_1value), symget)

			tmp94735 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			tmp94736 := Call(__e, tmp94735, sym_dproperty_1vector_d)

			__e.TailApply(tmp94734, V3203, symprofile, tmp94736)
			return

		}, 0)

		tmp94737 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(MakeNumber(0))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp94733, tmp94737)
		return

	}, 1)

	tmp94738 := Call(__e, PrimNS1Value(symns2_1set), symshen_4get_1profile, tmp94732)

	_ = tmp94738

	tmp94739 := MakeNative(func(__e *ControlFlow) {
		V3206 := __e.Get(1)
		_ = V3206
		V3207 := __e.Get(2)
		_ = V3207
		tmp94740 := Call(__e, PrimNS1Value(symns2_1value), symput)

		tmp94741 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		tmp94742 := Call(__e, tmp94741, sym_dproperty_1vector_d)

		__e.TailApply(tmp94740, V3206, symprofile, V3207, tmp94742)
		return

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4put_1profile, tmp94739)
	return

}, 0)
