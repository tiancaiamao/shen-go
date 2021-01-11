package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen7148 := MakeNative(func(__e *ControlFlow) {
		V1722 := __e.Get(1)
		_ = V1722
		gen7145 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7146 := Call(__e, gen7145, V1722)

		var gen7147 Obj

		if True == gen7146 {
			gen7140 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen7141 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen7142 := Call(__e, gen7141, V1722)

			gen7143 := Call(__e, gen7140, gen7142)

			var gen7144 Obj

			if True == gen7143 {
				gen7134 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen7135 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen7136 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen7137 := Call(__e, gen7136, V1722)

				gen7138 := Call(__e, gen7135, gen7137)

				gen7139 := Call(__e, gen7134, Nil, gen7138)

				if True == gen7139 {
					gen7144 = True
				} else {
					gen7144 = False
				}

			} else {
				gen7144 = False
			}

			if True == gen7144 {
				gen7147 = True
			} else {
				gen7147 = False
			}

		} else {
			gen7147 = False
		}

		if True == gen7147 {
			gen7125 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen7126 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen7127 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen7128 := Call(__e, PrimNS1Value(symns2_1value), symshen_4next_150)

			gen7129 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen7130 := Call(__e, gen7129, V1722)

			gen7131 := Call(__e, gen7128, MakeNumber(50), gen7130)

			gen7132 := Call(__e, gen7127, gen7131, MakeString("\n"), symshen_4a)

			gen7133 := Call(__e, gen7126, MakeString("datatype syntax error here:\n\n "), gen7132)

			__e.TailApply(gen7125, gen7133)

			return

		} else {
			if True == True {
				gen7124 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen7124, symshen_4datatype_1error)

				return

			} else {
				gen7123 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen7123, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4datatype_1error, gen7148)

	gen7197 := MakeNative(func(__e *ControlFlow) {
		V1724 := __e.Get(1)
		_ = V1724
		gen7166 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7162 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7163 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7164 := Call(__e, gen7163)

			gen7165 := Call(__e, gen7162, YaccParse, gen7164)

			if True == gen7165 {
				gen7159 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					gen7153 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7154 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7155 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7156 := Call(__e, gen7155)

					gen7157 := Call(__e, gen7154, gen7156, Parse___5e_6)

					gen7158 := Call(__e, gen7153, gen7157)

					if True == gen7158 {
						gen7150 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7151 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7152 := Call(__e, gen7151, Parse___5e_6)

						__e.TailApply(gen7150, gen7152, Nil)

						return

					} else {
						gen7149 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7149)

						return

					}

				}, 1)

				gen7160 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				gen7161 := Call(__e, gen7160, V1724)

				__e.TailApply(gen7159, gen7161)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7193 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5datatype_1rule_6 := __e.Get(1)
			_ = Parse__shen_4_5datatype_1rule_6
			gen7187 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen7188 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7189 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7190 := Call(__e, gen7189)

			gen7191 := Call(__e, gen7188, gen7190, Parse__shen_4_5datatype_1rule_6)

			gen7192 := Call(__e, gen7187, gen7191)

			if True == gen7192 {
				gen7184 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5datatype_1rules_6 := __e.Get(1)
					_ = Parse__shen_4_5datatype_1rules_6
					gen7178 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7179 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7180 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7181 := Call(__e, gen7180)

					gen7182 := Call(__e, gen7179, gen7181, Parse__shen_4_5datatype_1rules_6)

					gen7183 := Call(__e, gen7178, gen7182)

					if True == gen7183 {
						gen7169 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7170 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7171 := Call(__e, gen7170, Parse__shen_4_5datatype_1rules_6)

						gen7172 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen7173 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen7174 := Call(__e, gen7173, Parse__shen_4_5datatype_1rule_6)

						gen7175 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen7176 := Call(__e, gen7175, Parse__shen_4_5datatype_1rules_6)

						gen7177 := Call(__e, gen7172, gen7174, gen7176)

						__e.TailApply(gen7169, gen7171, gen7177)

						return

					} else {
						gen7168 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7168)

						return

					}

				}, 1)

				gen7185 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5datatype_1rules_6)

				gen7186 := Call(__e, gen7185, Parse__shen_4_5datatype_1rule_6)

				__e.TailApply(gen7184, gen7186)

				return

			} else {
				gen7167 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen7167)

				return

			}

		}, 1)

		gen7194 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5datatype_1rule_6)

		gen7195 := Call(__e, gen7194, V1724)

		gen7196 := Call(__e, gen7193, gen7195)

		__e.TailApply(gen7166, gen7196)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5datatype_1rules_6, gen7197)

	gen7318 := MakeNative(func(__e *ControlFlow) {
		V1726 := __e.Get(1)
		_ = V1726
		gen7259 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7255 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7256 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7257 := Call(__e, gen7256)

			gen7258 := Call(__e, gen7255, YaccParse, gen7257)

			if True == gen7258 {
				gen7252 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5side_1conditions_6 := __e.Get(1)
					_ = Parse__shen_4_5side_1conditions_6
					gen7246 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7247 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7248 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7249 := Call(__e, gen7248)

					gen7250 := Call(__e, gen7247, gen7249, Parse__shen_4_5side_1conditions_6)

					gen7251 := Call(__e, gen7246, gen7250)

					if True == gen7251 {
						gen7243 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5premises_6 := __e.Get(1)
							_ = Parse__shen_4_5premises_6
							gen7237 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7238 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7239 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7240 := Call(__e, gen7239)

							gen7241 := Call(__e, gen7238, gen7240, Parse__shen_4_5premises_6)

							gen7242 := Call(__e, gen7237, gen7241)

							if True == gen7242 {
								gen7234 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5doubleunderline_6 := __e.Get(1)
									_ = Parse__shen_4_5doubleunderline_6
									gen7228 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen7229 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen7230 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen7231 := Call(__e, gen7230)

									gen7232 := Call(__e, gen7229, gen7231, Parse__shen_4_5doubleunderline_6)

									gen7233 := Call(__e, gen7228, gen7232)

									if True == gen7233 {
										gen7225 := MakeNative(func(__e *ControlFlow) {
											Parse__shen_4_5conclusion_6 := __e.Get(1)
											_ = Parse__shen_4_5conclusion_6
											gen7219 := Call(__e, PrimNS1Value(symns2_1value), symnot)

											gen7220 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

											gen7221 := Call(__e, PrimNS1Value(symns2_1value), symfail)

											gen7222 := Call(__e, gen7221)

											gen7223 := Call(__e, gen7220, gen7222, Parse__shen_4_5conclusion_6)

											gen7224 := Call(__e, gen7219, gen7223)

											if True == gen7224 {
												gen7202 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

												gen7203 := Call(__e, PrimNS1Value(symns2_1value), symhd)

												gen7204 := Call(__e, gen7203, Parse__shen_4_5conclusion_6)

												gen7205 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

												gen7206 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												gen7207 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												gen7208 := Call(__e, gen7207, Parse__shen_4_5side_1conditions_6)

												gen7209 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												gen7210 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												gen7211 := Call(__e, gen7210, Parse__shen_4_5premises_6)

												gen7212 := Call(__e, PrimNS1Value(symns2_1value), symcons)

												gen7213 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

												gen7214 := Call(__e, gen7213, Parse__shen_4_5conclusion_6)

												gen7215 := Call(__e, gen7212, gen7214, Nil)

												gen7216 := Call(__e, gen7209, gen7211, gen7215)

												gen7217 := Call(__e, gen7206, gen7208, gen7216)

												gen7218 := Call(__e, gen7205, symshen_4double, gen7217)

												__e.TailApply(gen7202, gen7204, gen7218)

												return

											} else {
												gen7201 := Call(__e, PrimNS1Value(symns2_1value), symfail)

												__e.TailApply(gen7201)

												return

											}

										}, 1)

										gen7226 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5conclusion_6)

										gen7227 := Call(__e, gen7226, Parse__shen_4_5doubleunderline_6)

										__e.TailApply(gen7225, gen7227)

										return

									} else {
										gen7200 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen7200)

										return

									}

								}, 1)

								gen7235 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5doubleunderline_6)

								gen7236 := Call(__e, gen7235, Parse__shen_4_5premises_6)

								__e.TailApply(gen7234, gen7236)

								return

							} else {
								gen7199 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7199)

								return

							}

						}, 1)

						gen7244 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premises_6)

						gen7245 := Call(__e, gen7244, Parse__shen_4_5side_1conditions_6)

						__e.TailApply(gen7243, gen7245)

						return

					} else {
						gen7198 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7198)

						return

					}

				}, 1)

				gen7253 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1conditions_6)

				gen7254 := Call(__e, gen7253, V1726)

				__e.TailApply(gen7252, gen7254)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7314 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5side_1conditions_6 := __e.Get(1)
			_ = Parse__shen_4_5side_1conditions_6
			gen7308 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen7309 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7310 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7311 := Call(__e, gen7310)

			gen7312 := Call(__e, gen7309, gen7311, Parse__shen_4_5side_1conditions_6)

			gen7313 := Call(__e, gen7308, gen7312)

			if True == gen7313 {
				gen7305 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5premises_6 := __e.Get(1)
					_ = Parse__shen_4_5premises_6
					gen7299 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7300 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7301 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7302 := Call(__e, gen7301)

					gen7303 := Call(__e, gen7300, gen7302, Parse__shen_4_5premises_6)

					gen7304 := Call(__e, gen7299, gen7303)

					if True == gen7304 {
						gen7296 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5singleunderline_6 := __e.Get(1)
							_ = Parse__shen_4_5singleunderline_6
							gen7290 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7291 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7292 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7293 := Call(__e, gen7292)

							gen7294 := Call(__e, gen7291, gen7293, Parse__shen_4_5singleunderline_6)

							gen7295 := Call(__e, gen7290, gen7294)

							if True == gen7295 {
								gen7287 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5conclusion_6 := __e.Get(1)
									_ = Parse__shen_4_5conclusion_6
									gen7281 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen7282 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen7283 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen7284 := Call(__e, gen7283)

									gen7285 := Call(__e, gen7282, gen7284, Parse__shen_4_5conclusion_6)

									gen7286 := Call(__e, gen7281, gen7285)

									if True == gen7286 {
										gen7264 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen7265 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen7266 := Call(__e, gen7265, Parse__shen_4_5conclusion_6)

										gen7267 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

										gen7268 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen7269 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7270 := Call(__e, gen7269, Parse__shen_4_5side_1conditions_6)

										gen7271 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen7272 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7273 := Call(__e, gen7272, Parse__shen_4_5premises_6)

										gen7274 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen7275 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7276 := Call(__e, gen7275, Parse__shen_4_5conclusion_6)

										gen7277 := Call(__e, gen7274, gen7276, Nil)

										gen7278 := Call(__e, gen7271, gen7273, gen7277)

										gen7279 := Call(__e, gen7268, gen7270, gen7278)

										gen7280 := Call(__e, gen7267, symshen_4single, gen7279)

										__e.TailApply(gen7264, gen7266, gen7280)

										return

									} else {
										gen7263 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen7263)

										return

									}

								}, 1)

								gen7288 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5conclusion_6)

								gen7289 := Call(__e, gen7288, Parse__shen_4_5singleunderline_6)

								__e.TailApply(gen7287, gen7289)

								return

							} else {
								gen7262 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7262)

								return

							}

						}, 1)

						gen7297 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5singleunderline_6)

						gen7298 := Call(__e, gen7297, Parse__shen_4_5premises_6)

						__e.TailApply(gen7296, gen7298)

						return

					} else {
						gen7261 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7261)

						return

					}

				}, 1)

				gen7306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premises_6)

				gen7307 := Call(__e, gen7306, Parse__shen_4_5side_1conditions_6)

				__e.TailApply(gen7305, gen7307)

				return

			} else {
				gen7260 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen7260)

				return

			}

		}, 1)

		gen7315 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1conditions_6)

		gen7316 := Call(__e, gen7315, V1726)

		gen7317 := Call(__e, gen7314, gen7316)

		__e.TailApply(gen7259, gen7317)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5datatype_1rule_6, gen7318)

	gen7367 := MakeNative(func(__e *ControlFlow) {
		V1728 := __e.Get(1)
		_ = V1728
		gen7336 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7332 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7333 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7334 := Call(__e, gen7333)

			gen7335 := Call(__e, gen7332, YaccParse, gen7334)

			if True == gen7335 {
				gen7329 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					gen7323 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7324 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7325 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7326 := Call(__e, gen7325)

					gen7327 := Call(__e, gen7324, gen7326, Parse___5e_6)

					gen7328 := Call(__e, gen7323, gen7327)

					if True == gen7328 {
						gen7320 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7321 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7322 := Call(__e, gen7321, Parse___5e_6)

						__e.TailApply(gen7320, gen7322, Nil)

						return

					} else {
						gen7319 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7319)

						return

					}

				}, 1)

				gen7330 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				gen7331 := Call(__e, gen7330, V1728)

				__e.TailApply(gen7329, gen7331)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7363 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5side_1condition_6 := __e.Get(1)
			_ = Parse__shen_4_5side_1condition_6
			gen7357 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen7358 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7359 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7360 := Call(__e, gen7359)

			gen7361 := Call(__e, gen7358, gen7360, Parse__shen_4_5side_1condition_6)

			gen7362 := Call(__e, gen7357, gen7361)

			if True == gen7362 {
				gen7354 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5side_1conditions_6 := __e.Get(1)
					_ = Parse__shen_4_5side_1conditions_6
					gen7348 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7349 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7350 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7351 := Call(__e, gen7350)

					gen7352 := Call(__e, gen7349, gen7351, Parse__shen_4_5side_1conditions_6)

					gen7353 := Call(__e, gen7348, gen7352)

					if True == gen7353 {
						gen7339 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7340 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7341 := Call(__e, gen7340, Parse__shen_4_5side_1conditions_6)

						gen7342 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen7343 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen7344 := Call(__e, gen7343, Parse__shen_4_5side_1condition_6)

						gen7345 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen7346 := Call(__e, gen7345, Parse__shen_4_5side_1conditions_6)

						gen7347 := Call(__e, gen7342, gen7344, gen7346)

						__e.TailApply(gen7339, gen7341, gen7347)

						return

					} else {
						gen7338 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7338)

						return

					}

				}, 1)

				gen7355 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1conditions_6)

				gen7356 := Call(__e, gen7355, Parse__shen_4_5side_1condition_6)

				__e.TailApply(gen7354, gen7356)

				return

			} else {
				gen7337 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen7337)

				return

			}

		}, 1)

		gen7364 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5side_1condition_6)

		gen7365 := Call(__e, gen7364, V1728)

		gen7366 := Call(__e, gen7363, gen7365)

		__e.TailApply(gen7336, gen7366)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5side_1conditions_6, gen7367)

	gen7460 := MakeNative(func(__e *ControlFlow) {
		V1732 := __e.Get(1)
		_ = V1732
		gen7422 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7418 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7419 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7420 := Call(__e, gen7419)

			gen7421 := Call(__e, gen7418, YaccParse, gen7420)

			if True == gen7421 {
				gen7413 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen7414 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen7415 := Call(__e, gen7414, V1732)

				gen7416 := Call(__e, gen7413, gen7415)

				var gen7417 Obj

				if True == gen7416 {
					gen7409 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7410 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					gen7411 := Call(__e, gen7410, V1732)

					gen7412 := Call(__e, gen7409, symlet, gen7411)

					if True == gen7412 {
						gen7417 = True
					} else {
						gen7417 = False
					}

				} else {
					gen7417 = False
				}

				if True == gen7417 {
					gen7402 := MakeNative(func(__e *ControlFlow) {
						NewStream1730 := __e.Get(1)
						_ = NewStream1730
						gen7399 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5variable_2_6 := __e.Get(1)
							_ = Parse__shen_4_5variable_2_6
							gen7393 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7394 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7395 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7396 := Call(__e, gen7395)

							gen7397 := Call(__e, gen7394, gen7396, Parse__shen_4_5variable_2_6)

							gen7398 := Call(__e, gen7393, gen7397)

							if True == gen7398 {
								gen7390 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5expr_6 := __e.Get(1)
									_ = Parse__shen_4_5expr_6
									gen7384 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen7385 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen7386 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen7387 := Call(__e, gen7386)

									gen7388 := Call(__e, gen7385, gen7387, Parse__shen_4_5expr_6)

									gen7389 := Call(__e, gen7384, gen7388)

									if True == gen7389 {
										gen7371 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen7372 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen7373 := Call(__e, gen7372, Parse__shen_4_5expr_6)

										gen7374 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen7375 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen7376 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7377 := Call(__e, gen7376, Parse__shen_4_5variable_2_6)

										gen7378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

										gen7379 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7380 := Call(__e, gen7379, Parse__shen_4_5expr_6)

										gen7381 := Call(__e, gen7378, gen7380, Nil)

										gen7382 := Call(__e, gen7375, gen7377, gen7381)

										gen7383 := Call(__e, gen7374, symlet, gen7382)

										__e.TailApply(gen7371, gen7373, gen7383)

										return

									} else {
										gen7370 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen7370)

										return

									}

								}, 1)

								gen7391 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

								gen7392 := Call(__e, gen7391, Parse__shen_4_5variable_2_6)

								__e.TailApply(gen7390, gen7392)

								return

							} else {
								gen7369 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7369)

								return

							}

						}, 1)

						gen7400 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5variable_2_6)

						gen7401 := Call(__e, gen7400, NewStream1730)

						__e.TailApply(gen7399, gen7401)

						return

					}, 1)

					gen7403 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7404 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen7405 := Call(__e, gen7404, V1732)

					gen7406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen7407 := Call(__e, gen7406, V1732)

					gen7408 := Call(__e, gen7403, gen7405, gen7407)

					__e.TailApply(gen7402, gen7408)

					return

				} else {
					gen7368 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen7368)

					return

				}

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7454 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7455 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen7456 := Call(__e, gen7455, V1732)

		gen7457 := Call(__e, gen7454, gen7456)

		var gen7458 Obj

		if True == gen7457 {
			gen7450 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7451 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen7452 := Call(__e, gen7451, V1732)

			gen7453 := Call(__e, gen7450, symif, gen7452)

			if True == gen7453 {
				gen7458 = True
			} else {
				gen7458 = False
			}

		} else {
			gen7458 = False
		}

		var gen7459 Obj

		if True == gen7458 {
			gen7443 := MakeNative(func(__e *ControlFlow) {
				NewStream1729 := __e.Get(1)
				_ = NewStream1729
				gen7440 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5expr_6 := __e.Get(1)
					_ = Parse__shen_4_5expr_6
					gen7434 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7435 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7436 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7437 := Call(__e, gen7436)

					gen7438 := Call(__e, gen7435, gen7437, Parse__shen_4_5expr_6)

					gen7439 := Call(__e, gen7434, gen7438)

					if True == gen7439 {
						gen7425 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7426 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7427 := Call(__e, gen7426, Parse__shen_4_5expr_6)

						gen7428 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen7429 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen7430 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen7431 := Call(__e, gen7430, Parse__shen_4_5expr_6)

						gen7432 := Call(__e, gen7429, gen7431, Nil)

						gen7433 := Call(__e, gen7428, symif, gen7432)

						__e.TailApply(gen7425, gen7427, gen7433)

						return

					} else {
						gen7424 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7424)

						return

					}

				}, 1)

				gen7441 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

				gen7442 := Call(__e, gen7441, NewStream1729)

				__e.TailApply(gen7440, gen7442)

				return

			}, 1)

			gen7444 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen7445 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen7446 := Call(__e, gen7445, V1732)

			gen7447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen7448 := Call(__e, gen7447, V1732)

			gen7449 := Call(__e, gen7444, gen7446, gen7448)

			gen7459 = Call(__e, gen7443, gen7449)

		} else {
			gen7423 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7459 = Call(__e, gen7423)

		}

		__e.TailApply(gen7422, gen7459)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5side_1condition_6, gen7460)

	gen7481 := MakeNative(func(__e *ControlFlow) {
		V1734 := __e.Get(1)
		_ = V1734
		gen7477 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7478 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen7479 := Call(__e, gen7478, V1734)

		gen7480 := Call(__e, gen7477, gen7479)

		if True == gen7480 {
			gen7474 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen7472 := Call(__e, PrimNS1Value(symns2_1value), symvariable_2)

				gen7473 := Call(__e, gen7472, Parse__X)

				if True == gen7473 {
					gen7463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7464 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen7465 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7466 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen7467 := Call(__e, gen7466, V1734)

					gen7468 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen7469 := Call(__e, gen7468, V1734)

					gen7470 := Call(__e, gen7465, gen7467, gen7469)

					gen7471 := Call(__e, gen7464, gen7470)

					__e.TailApply(gen7463, gen7471, Parse__X)

					return

				} else {
					gen7462 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen7462)

					return

				}

			}, 1)

			gen7475 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen7476 := Call(__e, gen7475, V1734)

			__e.TailApply(gen7474, gen7476)

			return

		} else {
			gen7461 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen7461)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5variable_2_6, gen7481)

	gen7516 := MakeNative(func(__e *ControlFlow) {
		V1736 := __e.Get(1)
		_ = V1736
		gen7512 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7513 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen7514 := Call(__e, gen7513, V1736)

		gen7515 := Call(__e, gen7512, gen7514)

		if True == gen7515 {
			gen7509 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen7495 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				gen7501 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen7502 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen7503 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen7504 := Call(__e, gen7503, sym_k, Nil)

				gen7505 := Call(__e, gen7502, sym_6_6, gen7504)

				gen7506 := Call(__e, gen7501, Parse__X, gen7505)

				var gen7507 Obj

				if True == gen7506 {
					gen7507 = True
				} else {
					gen7498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4singleunderline_2)

					gen7499 := Call(__e, gen7498, Parse__X)

					var gen7500 Obj

					if True == gen7499 {
						gen7500 = True
					} else {
						gen7496 := Call(__e, PrimNS1Value(symns2_1value), symshen_4doubleunderline_2)

						gen7497 := Call(__e, gen7496, Parse__X)

						if True == gen7497 {
							gen7500 = True
						} else {
							gen7500 = False
						}

					}

					if True == gen7500 {
						gen7507 = True
					} else {
						gen7507 = False
					}

				}

				gen7508 := Call(__e, gen7495, gen7507)

				if True == gen7508 {
					gen7484 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7485 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen7486 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7487 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen7488 := Call(__e, gen7487, V1736)

					gen7489 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen7490 := Call(__e, gen7489, V1736)

					gen7491 := Call(__e, gen7486, gen7488, gen7490)

					gen7492 := Call(__e, gen7485, gen7491)

					gen7493 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1bar)

					gen7494 := Call(__e, gen7493, Parse__X)

					__e.TailApply(gen7484, gen7492, gen7494)

					return

				} else {
					gen7483 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen7483)

					return

				}

			}, 1)

			gen7510 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen7511 := Call(__e, gen7510, V1736)

			__e.TailApply(gen7509, gen7511)

			return

		} else {
			gen7482 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen7482)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5expr_6, gen7516)

	gen7568 := MakeNative(func(__e *ControlFlow) {
		V1738 := __e.Get(1)
		_ = V1738
		gen7565 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7566 := Call(__e, gen7565, V1738)

		var gen7567 Obj

		if True == gen7566 {
			gen7560 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen7561 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen7562 := Call(__e, gen7561, V1738)

			gen7563 := Call(__e, gen7560, gen7562)

			var gen7564 Obj

			if True == gen7563 {
				gen7553 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen7554 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen7555 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen7556 := Call(__e, gen7555, V1738)

				gen7557 := Call(__e, gen7554, gen7556)

				gen7558 := Call(__e, gen7553, gen7557)

				var gen7559 Obj

				if True == gen7558 {
					gen7544 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7545 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen7546 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen7547 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen7548 := Call(__e, gen7547, V1738)

					gen7549 := Call(__e, gen7546, gen7548)

					gen7550 := Call(__e, gen7545, gen7549)

					gen7551 := Call(__e, gen7544, Nil, gen7550)

					var gen7552 Obj

					if True == gen7551 {
						gen7538 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen7539 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7540 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen7541 := Call(__e, gen7540, V1738)

						gen7542 := Call(__e, gen7539, gen7541)

						gen7543 := Call(__e, gen7538, gen7542, symbar_b)

						if True == gen7543 {
							gen7552 = True
						} else {
							gen7552 = False
						}

					} else {
						gen7552 = False
					}

					if True == gen7552 {
						gen7559 = True
					} else {
						gen7559 = False
					}

				} else {
					gen7559 = False
				}

				if True == gen7559 {
					gen7564 = True
				} else {
					gen7564 = False
				}

			} else {
				gen7564 = False
			}

			if True == gen7564 {
				gen7567 = True
			} else {
				gen7567 = False
			}

		} else {
			gen7567 = False
		}

		if True == gen7567 {
			gen7529 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen7530 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen7531 := Call(__e, gen7530, V1738)

			gen7532 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen7533 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen7534 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen7535 := Call(__e, gen7534, V1738)

			gen7536 := Call(__e, gen7533, gen7535)

			gen7537 := Call(__e, gen7532, gen7536)

			__e.TailApply(gen7529, gen7531, gen7537)

			return

		} else {
			gen7527 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen7528 := Call(__e, gen7527, V1738)

			if True == gen7528 {
				gen7518 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen7519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1bar)

				gen7520 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen7521 := Call(__e, gen7520, V1738)

				gen7522 := Call(__e, gen7519, gen7521)

				gen7523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remove_1bar)

				gen7524 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen7525 := Call(__e, gen7524, V1738)

				gen7526 := Call(__e, gen7523, gen7525)

				__e.TailApply(gen7518, gen7522, gen7526)

				return

			} else {
				if True == True {
					__e.Return(V1738)
					return
				} else {
					gen7517 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen7517, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1bar, gen7568)

	gen7627 := MakeNative(func(__e *ControlFlow) {
		V1740 := __e.Get(1)
		_ = V1740
		gen7586 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7582 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7583 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7584 := Call(__e, gen7583)

			gen7585 := Call(__e, gen7582, YaccParse, gen7584)

			if True == gen7585 {
				gen7579 := MakeNative(func(__e *ControlFlow) {
					Parse___5e_6 := __e.Get(1)
					_ = Parse___5e_6
					gen7573 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7574 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7575 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7576 := Call(__e, gen7575)

					gen7577 := Call(__e, gen7574, gen7576, Parse___5e_6)

					gen7578 := Call(__e, gen7573, gen7577)

					if True == gen7578 {
						gen7570 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7571 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7572 := Call(__e, gen7571, Parse___5e_6)

						__e.TailApply(gen7570, gen7572, Nil)

						return

					} else {
						gen7569 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7569)

						return

					}

				}, 1)

				gen7580 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

				gen7581 := Call(__e, gen7580, V1740)

				__e.TailApply(gen7579, gen7581)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7623 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5premise_6 := __e.Get(1)
			_ = Parse__shen_4_5premise_6
			gen7617 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen7618 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7619 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7620 := Call(__e, gen7619)

			gen7621 := Call(__e, gen7618, gen7620, Parse__shen_4_5premise_6)

			gen7622 := Call(__e, gen7617, gen7621)

			if True == gen7622 {
				gen7614 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
					_ = Parse__shen_4_5semicolon_1symbol_6
					gen7608 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7609 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7610 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7611 := Call(__e, gen7610)

					gen7612 := Call(__e, gen7609, gen7611, Parse__shen_4_5semicolon_1symbol_6)

					gen7613 := Call(__e, gen7608, gen7612)

					if True == gen7613 {
						gen7605 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5premises_6 := __e.Get(1)
							_ = Parse__shen_4_5premises_6
							gen7599 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7600 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7601 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7602 := Call(__e, gen7601)

							gen7603 := Call(__e, gen7600, gen7602, Parse__shen_4_5premises_6)

							gen7604 := Call(__e, gen7599, gen7603)

							if True == gen7604 {
								gen7590 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen7591 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen7592 := Call(__e, gen7591, Parse__shen_4_5premises_6)

								gen7593 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen7594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7595 := Call(__e, gen7594, Parse__shen_4_5premise_6)

								gen7596 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7597 := Call(__e, gen7596, Parse__shen_4_5premises_6)

								gen7598 := Call(__e, gen7593, gen7595, gen7597)

								__e.TailApply(gen7590, gen7592, gen7598)

								return

							} else {
								gen7589 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7589)

								return

							}

						}, 1)

						gen7606 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premises_6)

						gen7607 := Call(__e, gen7606, Parse__shen_4_5semicolon_1symbol_6)

						__e.TailApply(gen7605, gen7607)

						return

					} else {
						gen7588 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7588)

						return

					}

				}, 1)

				gen7615 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_1symbol_6)

				gen7616 := Call(__e, gen7615, Parse__shen_4_5premise_6)

				__e.TailApply(gen7614, gen7616)

				return

			} else {
				gen7587 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen7587)

				return

			}

		}, 1)

		gen7624 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5premise_6)

		gen7625 := Call(__e, gen7624, V1740)

		gen7626 := Call(__e, gen7623, gen7625)

		__e.TailApply(gen7586, gen7626)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5premises_6, gen7627)

	gen7648 := MakeNative(func(__e *ControlFlow) {
		V1742 := __e.Get(1)
		_ = V1742
		gen7644 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7645 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen7646 := Call(__e, gen7645, V1742)

		gen7647 := Call(__e, gen7644, gen7646)

		if True == gen7647 {
			gen7641 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen7639 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen7640 := Call(__e, gen7639, Parse__X, sym_k)

				if True == gen7640 {
					gen7630 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7631 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen7632 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7633 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen7634 := Call(__e, gen7633, V1742)

					gen7635 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen7636 := Call(__e, gen7635, V1742)

					gen7637 := Call(__e, gen7632, gen7634, gen7636)

					gen7638 := Call(__e, gen7631, gen7637)

					__e.TailApply(gen7630, gen7638, symshen_4skip)

					return

				} else {
					gen7629 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen7629)

					return

				}

			}, 1)

			gen7642 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen7643 := Call(__e, gen7642, V1742)

			__e.TailApply(gen7641, gen7643)

			return

		} else {
			gen7628 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen7628)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5semicolon_1symbol_6, gen7648)

	gen7744 := MakeNative(func(__e *ControlFlow) {
		V1746 := __e.Get(1)
		_ = V1746
		gen7722 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7718 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7719 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7720 := Call(__e, gen7719)

			gen7721 := Call(__e, gen7718, YaccParse, gen7720)

			if True == gen7721 {
				gen7670 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					gen7666 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7667 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7668 := Call(__e, gen7667)

					gen7669 := Call(__e, gen7666, YaccParse, gen7668)

					if True == gen7669 {
						gen7663 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formula_6 := __e.Get(1)
							_ = Parse__shen_4_5formula_6
							gen7657 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7658 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7659 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7660 := Call(__e, gen7659)

							gen7661 := Call(__e, gen7658, gen7660, Parse__shen_4_5formula_6)

							gen7662 := Call(__e, gen7657, gen7661)

							if True == gen7662 {
								gen7650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen7651 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen7652 := Call(__e, gen7651, Parse__shen_4_5formula_6)

								gen7653 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

								gen7654 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7655 := Call(__e, gen7654, Parse__shen_4_5formula_6)

								gen7656 := Call(__e, gen7653, Nil, gen7655)

								__e.TailApply(gen7650, gen7652, gen7656)

								return

							} else {
								gen7649 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7649)

								return

							}

						}, 1)

						gen7664 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

						gen7665 := Call(__e, gen7664, V1746)

						__e.TailApply(gen7663, gen7665)

						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				gen7714 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formulae_6 := __e.Get(1)
					_ = Parse__shen_4_5formulae_6
					gen7708 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7709 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7710 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7711 := Call(__e, gen7710)

					gen7712 := Call(__e, gen7709, gen7711, Parse__shen_4_5formulae_6)

					gen7713 := Call(__e, gen7708, gen7712)

					if True == gen7713 {
						gen7703 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen7704 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7705 := Call(__e, gen7704, Parse__shen_4_5formulae_6)

						gen7706 := Call(__e, gen7703, gen7705)

						var gen7707 Obj

						if True == gen7706 {
							gen7699 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7700 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

							gen7701 := Call(__e, gen7700, Parse__shen_4_5formulae_6)

							gen7702 := Call(__e, gen7699, sym_6_6, gen7701)

							if True == gen7702 {
								gen7707 = True
							} else {
								gen7707 = False
							}

						} else {
							gen7707 = False
						}

						if True == gen7707 {
							gen7692 := MakeNative(func(__e *ControlFlow) {
								NewStream1744 := __e.Get(1)
								_ = NewStream1744
								gen7689 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5formula_6 := __e.Get(1)
									_ = Parse__shen_4_5formula_6
									gen7683 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen7684 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen7685 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen7686 := Call(__e, gen7685)

									gen7687 := Call(__e, gen7684, gen7686, Parse__shen_4_5formula_6)

									gen7688 := Call(__e, gen7683, gen7687)

									if True == gen7688 {
										gen7674 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen7675 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen7676 := Call(__e, gen7675, Parse__shen_4_5formula_6)

										gen7677 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

										gen7678 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7679 := Call(__e, gen7678, Parse__shen_4_5formulae_6)

										gen7680 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7681 := Call(__e, gen7680, Parse__shen_4_5formula_6)

										gen7682 := Call(__e, gen7677, gen7679, gen7681)

										__e.TailApply(gen7674, gen7676, gen7682)

										return

									} else {
										gen7673 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen7673)

										return

									}

								}, 1)

								gen7690 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

								gen7691 := Call(__e, gen7690, NewStream1744)

								__e.TailApply(gen7689, gen7691)

								return

							}, 1)

							gen7693 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

							gen7694 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

							gen7695 := Call(__e, gen7694, Parse__shen_4_5formulae_6)

							gen7696 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

							gen7697 := Call(__e, gen7696, Parse__shen_4_5formulae_6)

							gen7698 := Call(__e, gen7693, gen7695, gen7697)

							__e.TailApply(gen7692, gen7698)

							return

						} else {
							gen7672 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							__e.TailApply(gen7672)

							return

						}

					} else {
						gen7671 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7671)

						return

					}

				}, 1)

				gen7715 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formulae_6)

				gen7716 := Call(__e, gen7715, V1746)

				gen7717 := Call(__e, gen7714, gen7716)

				__e.TailApply(gen7670, gen7717)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7738 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7739 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen7740 := Call(__e, gen7739, V1746)

		gen7741 := Call(__e, gen7738, gen7740)

		var gen7742 Obj

		if True == gen7741 {
			gen7734 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7735 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen7736 := Call(__e, gen7735, V1746)

			gen7737 := Call(__e, gen7734, sym_b, gen7736)

			if True == gen7737 {
				gen7742 = True
			} else {
				gen7742 = False
			}

		} else {
			gen7742 = False
		}

		var gen7743 Obj

		if True == gen7742 {
			gen7727 := MakeNative(func(__e *ControlFlow) {
				NewStream1743 := __e.Get(1)
				_ = NewStream1743
				gen7724 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen7725 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen7726 := Call(__e, gen7725, NewStream1743)

				__e.TailApply(gen7724, gen7726, sym_b)

				return

			}, 1)

			gen7728 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

			gen7729 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

			gen7730 := Call(__e, gen7729, V1746)

			gen7731 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

			gen7732 := Call(__e, gen7731, V1746)

			gen7733 := Call(__e, gen7728, gen7730, gen7732)

			gen7743 = Call(__e, gen7727, gen7733)

		} else {
			gen7723 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7743 = Call(__e, gen7723)

		}

		__e.TailApply(gen7722, gen7743)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5premise_6, gen7744)

	gen7834 := MakeNative(func(__e *ControlFlow) {
		V1749 := __e.Get(1)
		_ = V1749
		gen7776 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7772 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7773 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7774 := Call(__e, gen7773)

			gen7775 := Call(__e, gen7772, YaccParse, gen7774)

			if True == gen7775 {
				gen7769 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formula_6 := __e.Get(1)
					_ = Parse__shen_4_5formula_6
					gen7763 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7764 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7765 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7766 := Call(__e, gen7765)

					gen7767 := Call(__e, gen7764, gen7766, Parse__shen_4_5formula_6)

					gen7768 := Call(__e, gen7763, gen7767)

					if True == gen7768 {
						gen7760 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
							_ = Parse__shen_4_5semicolon_1symbol_6
							gen7754 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7755 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7756 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7757 := Call(__e, gen7756)

							gen7758 := Call(__e, gen7755, gen7757, Parse__shen_4_5semicolon_1symbol_6)

							gen7759 := Call(__e, gen7754, gen7758)

							if True == gen7759 {
								gen7747 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen7748 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen7749 := Call(__e, gen7748, Parse__shen_4_5semicolon_1symbol_6)

								gen7750 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

								gen7751 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7752 := Call(__e, gen7751, Parse__shen_4_5formula_6)

								gen7753 := Call(__e, gen7750, Nil, gen7752)

								__e.TailApply(gen7747, gen7749, gen7753)

								return

							} else {
								gen7746 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7746)

								return

							}

						}, 1)

						gen7761 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_1symbol_6)

						gen7762 := Call(__e, gen7761, Parse__shen_4_5formula_6)

						__e.TailApply(gen7760, gen7762)

						return

					} else {
						gen7745 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7745)

						return

					}

				}, 1)

				gen7770 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

				gen7771 := Call(__e, gen7770, V1749)

				__e.TailApply(gen7769, gen7771)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7830 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5formulae_6 := __e.Get(1)
			_ = Parse__shen_4_5formulae_6
			gen7824 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen7825 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7826 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7827 := Call(__e, gen7826)

			gen7828 := Call(__e, gen7825, gen7827, Parse__shen_4_5formulae_6)

			gen7829 := Call(__e, gen7824, gen7828)

			if True == gen7829 {
				gen7819 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen7820 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen7821 := Call(__e, gen7820, Parse__shen_4_5formulae_6)

				gen7822 := Call(__e, gen7819, gen7821)

				var gen7823 Obj

				if True == gen7822 {
					gen7815 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7816 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					gen7817 := Call(__e, gen7816, Parse__shen_4_5formulae_6)

					gen7818 := Call(__e, gen7815, sym_6_6, gen7817)

					if True == gen7818 {
						gen7823 = True
					} else {
						gen7823 = False
					}

				} else {
					gen7823 = False
				}

				if True == gen7823 {
					gen7808 := MakeNative(func(__e *ControlFlow) {
						NewStream1747 := __e.Get(1)
						_ = NewStream1747
						gen7805 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formula_6 := __e.Get(1)
							_ = Parse__shen_4_5formula_6
							gen7799 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7800 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7801 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7802 := Call(__e, gen7801)

							gen7803 := Call(__e, gen7800, gen7802, Parse__shen_4_5formula_6)

							gen7804 := Call(__e, gen7799, gen7803)

							if True == gen7804 {
								gen7796 := MakeNative(func(__e *ControlFlow) {
									Parse__shen_4_5semicolon_1symbol_6 := __e.Get(1)
									_ = Parse__shen_4_5semicolon_1symbol_6
									gen7790 := Call(__e, PrimNS1Value(symns2_1value), symnot)

									gen7791 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen7792 := Call(__e, PrimNS1Value(symns2_1value), symfail)

									gen7793 := Call(__e, gen7792)

									gen7794 := Call(__e, gen7791, gen7793, Parse__shen_4_5semicolon_1symbol_6)

									gen7795 := Call(__e, gen7790, gen7794)

									if True == gen7795 {
										gen7781 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

										gen7782 := Call(__e, PrimNS1Value(symns2_1value), symhd)

										gen7783 := Call(__e, gen7782, Parse__shen_4_5semicolon_1symbol_6)

										gen7784 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sequent)

										gen7785 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7786 := Call(__e, gen7785, Parse__shen_4_5formulae_6)

										gen7787 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

										gen7788 := Call(__e, gen7787, Parse__shen_4_5formula_6)

										gen7789 := Call(__e, gen7784, gen7786, gen7788)

										__e.TailApply(gen7781, gen7783, gen7789)

										return

									} else {
										gen7780 := Call(__e, PrimNS1Value(symns2_1value), symfail)

										__e.TailApply(gen7780)

										return

									}

								}, 1)

								gen7797 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5semicolon_1symbol_6)

								gen7798 := Call(__e, gen7797, Parse__shen_4_5formula_6)

								__e.TailApply(gen7796, gen7798)

								return

							} else {
								gen7779 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7779)

								return

							}

						}, 1)

						gen7806 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

						gen7807 := Call(__e, gen7806, NewStream1747)

						__e.TailApply(gen7805, gen7807)

						return

					}, 1)

					gen7809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7810 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen7811 := Call(__e, gen7810, Parse__shen_4_5formulae_6)

					gen7812 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen7813 := Call(__e, gen7812, Parse__shen_4_5formulae_6)

					gen7814 := Call(__e, gen7809, gen7811, gen7813)

					__e.TailApply(gen7808, gen7814)

					return

				} else {
					gen7778 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen7778)

					return

				}

			} else {
				gen7777 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen7777)

				return

			}

		}, 1)

		gen7831 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formulae_6)

		gen7832 := Call(__e, gen7831, V1749)

		gen7833 := Call(__e, gen7830, gen7832)

		__e.TailApply(gen7776, gen7833)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5conclusion_6, gen7834)

	gen7836 := MakeNative(func(__e *ControlFlow) {
		V1752 := __e.Get(1)
		_ = V1752
		V1753 := __e.Get(2)
		_ = V1753
		gen7835 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

		__e.TailApply(gen7835, V1752, V1753)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4sequent, gen7836)

	gen7918 := MakeNative(func(__e *ControlFlow) {
		V1755 := __e.Get(1)
		_ = V1755
		gen7877 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7873 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7874 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7875 := Call(__e, gen7874)

			gen7876 := Call(__e, gen7873, YaccParse, gen7875)

			if True == gen7876 {
				gen7854 := MakeNative(func(__e *ControlFlow) {
					YaccParse := __e.Get(1)
					_ = YaccParse
					gen7850 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7851 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7852 := Call(__e, gen7851)

					gen7853 := Call(__e, gen7850, YaccParse, gen7852)

					if True == gen7853 {
						gen7847 := MakeNative(func(__e *ControlFlow) {
							Parse___5e_6 := __e.Get(1)
							_ = Parse___5e_6
							gen7841 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7842 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7843 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7844 := Call(__e, gen7843)

							gen7845 := Call(__e, gen7842, gen7844, Parse___5e_6)

							gen7846 := Call(__e, gen7841, gen7845)

							if True == gen7846 {
								gen7838 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen7839 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen7840 := Call(__e, gen7839, Parse___5e_6)

								__e.TailApply(gen7838, gen7840, Nil)

								return

							} else {
								gen7837 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7837)

								return

							}

						}, 1)

						gen7848 := Call(__e, PrimNS1Value(symns2_1value), sym_5e_6)

						gen7849 := Call(__e, gen7848, V1755)

						__e.TailApply(gen7847, gen7849)

						return

					} else {
						__e.Return(YaccParse)
						return
					}

				}, 1)

				gen7869 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5formula_6 := __e.Get(1)
					_ = Parse__shen_4_5formula_6
					gen7863 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7864 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7865 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7866 := Call(__e, gen7865)

					gen7867 := Call(__e, gen7864, gen7866, Parse__shen_4_5formula_6)

					gen7868 := Call(__e, gen7863, gen7867)

					if True == gen7868 {
						gen7856 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7857 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7858 := Call(__e, gen7857, Parse__shen_4_5formula_6)

						gen7859 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen7860 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen7861 := Call(__e, gen7860, Parse__shen_4_5formula_6)

						gen7862 := Call(__e, gen7859, gen7861, Nil)

						__e.TailApply(gen7856, gen7858, gen7862)

						return

					} else {
						gen7855 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7855)

						return

					}

				}, 1)

				gen7870 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

				gen7871 := Call(__e, gen7870, V1755)

				gen7872 := Call(__e, gen7869, gen7871)

				__e.TailApply(gen7854, gen7872)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen7914 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5formula_6 := __e.Get(1)
			_ = Parse__shen_4_5formula_6
			gen7908 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen7909 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7910 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7911 := Call(__e, gen7910)

			gen7912 := Call(__e, gen7909, gen7911, Parse__shen_4_5formula_6)

			gen7913 := Call(__e, gen7908, gen7912)

			if True == gen7913 {
				gen7905 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5comma_1symbol_6 := __e.Get(1)
					_ = Parse__shen_4_5comma_1symbol_6
					gen7899 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7900 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7901 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7902 := Call(__e, gen7901)

					gen7903 := Call(__e, gen7900, gen7902, Parse__shen_4_5comma_1symbol_6)

					gen7904 := Call(__e, gen7899, gen7903)

					if True == gen7904 {
						gen7896 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5formulae_6 := __e.Get(1)
							_ = Parse__shen_4_5formulae_6
							gen7890 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7891 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7892 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7893 := Call(__e, gen7892)

							gen7894 := Call(__e, gen7891, gen7893, Parse__shen_4_5formulae_6)

							gen7895 := Call(__e, gen7890, gen7894)

							if True == gen7895 {
								gen7881 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen7882 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen7883 := Call(__e, gen7882, Parse__shen_4_5formulae_6)

								gen7884 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen7885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7886 := Call(__e, gen7885, Parse__shen_4_5formula_6)

								gen7887 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7888 := Call(__e, gen7887, Parse__shen_4_5formulae_6)

								gen7889 := Call(__e, gen7884, gen7886, gen7888)

								__e.TailApply(gen7881, gen7883, gen7889)

								return

							} else {
								gen7880 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7880)

								return

							}

						}, 1)

						gen7897 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formulae_6)

						gen7898 := Call(__e, gen7897, Parse__shen_4_5comma_1symbol_6)

						__e.TailApply(gen7896, gen7898)

						return

					} else {
						gen7879 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7879)

						return

					}

				}, 1)

				gen7906 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5comma_1symbol_6)

				gen7907 := Call(__e, gen7906, Parse__shen_4_5formula_6)

				__e.TailApply(gen7905, gen7907)

				return

			} else {
				gen7878 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen7878)

				return

			}

		}, 1)

		gen7915 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5formula_6)

		gen7916 := Call(__e, gen7915, V1755)

		gen7917 := Call(__e, gen7914, gen7916)

		__e.TailApply(gen7877, gen7917)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5formulae_6, gen7918)

	gen7941 := MakeNative(func(__e *ControlFlow) {
		V1757 := __e.Get(1)
		_ = V1757
		gen7937 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen7938 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen7939 := Call(__e, gen7938, V1757)

		gen7940 := Call(__e, gen7937, gen7939)

		if True == gen7940 {
			gen7934 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen7930 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen7931 := Call(__e, PrimNS1Value(symns2_1value), symintern)

				gen7932 := Call(__e, gen7931, MakeString(","))

				gen7933 := Call(__e, gen7930, Parse__X, gen7932)

				if True == gen7933 {
					gen7921 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7922 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen7923 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7924 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen7925 := Call(__e, gen7924, V1757)

					gen7926 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen7927 := Call(__e, gen7926, V1757)

					gen7928 := Call(__e, gen7923, gen7925, gen7927)

					gen7929 := Call(__e, gen7922, gen7928)

					__e.TailApply(gen7921, gen7929, symshen_4skip)

					return

				} else {
					gen7920 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen7920)

					return

				}

			}, 1)

			gen7935 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen7936 := Call(__e, gen7935, V1757)

			__e.TailApply(gen7934, gen7936)

			return

		} else {
			gen7919 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen7919)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5comma_1symbol_6, gen7941)

	gen8017 := MakeNative(func(__e *ControlFlow) {
		V1760 := __e.Get(1)
		_ = V1760
		gen7961 := MakeNative(func(__e *ControlFlow) {
			YaccParse := __e.Get(1)
			_ = YaccParse
			gen7957 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7958 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen7959 := Call(__e, gen7958)

			gen7960 := Call(__e, gen7957, YaccParse, gen7959)

			if True == gen7960 {
				gen7954 := MakeNative(func(__e *ControlFlow) {
					Parse__shen_4_5expr_6 := __e.Get(1)
					_ = Parse__shen_4_5expr_6
					gen7948 := Call(__e, PrimNS1Value(symns2_1value), symnot)

					gen7949 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7950 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen7951 := Call(__e, gen7950)

					gen7952 := Call(__e, gen7949, gen7951, Parse__shen_4_5expr_6)

					gen7953 := Call(__e, gen7948, gen7952)

					if True == gen7953 {
						gen7943 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

						gen7944 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen7945 := Call(__e, gen7944, Parse__shen_4_5expr_6)

						gen7946 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

						gen7947 := Call(__e, gen7946, Parse__shen_4_5expr_6)

						__e.TailApply(gen7943, gen7945, gen7947)

						return

					} else {
						gen7942 := Call(__e, PrimNS1Value(symns2_1value), symfail)

						__e.TailApply(gen7942)

						return

					}

				}, 1)

				gen7955 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

				gen7956 := Call(__e, gen7955, V1760)

				__e.TailApply(gen7954, gen7956)

				return

			} else {
				__e.Return(YaccParse)
				return
			}

		}, 1)

		gen8013 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5expr_6 := __e.Get(1)
			_ = Parse__shen_4_5expr_6
			gen8007 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen8008 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8009 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen8010 := Call(__e, gen8009)

			gen8011 := Call(__e, gen8008, gen8010, Parse__shen_4_5expr_6)

			gen8012 := Call(__e, gen8007, gen8011)

			if True == gen8012 {
				gen8002 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8003 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8004 := Call(__e, gen8003, Parse__shen_4_5expr_6)

				gen8005 := Call(__e, gen8002, gen8004)

				var gen8006 Obj

				if True == gen8005 {
					gen7998 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen7999 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

					gen8000 := Call(__e, gen7999, Parse__shen_4_5expr_6)

					gen8001 := Call(__e, gen7998, sym_h, gen8000)

					if True == gen8001 {
						gen8006 = True
					} else {
						gen8006 = False
					}

				} else {
					gen8006 = False
				}

				if True == gen8006 {
					gen7991 := MakeNative(func(__e *ControlFlow) {
						NewStream1758 := __e.Get(1)
						_ = NewStream1758
						gen7988 := MakeNative(func(__e *ControlFlow) {
							Parse__shen_4_5type_6 := __e.Get(1)
							_ = Parse__shen_4_5type_6
							gen7982 := Call(__e, PrimNS1Value(symns2_1value), symnot)

							gen7983 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen7984 := Call(__e, PrimNS1Value(symns2_1value), symfail)

							gen7985 := Call(__e, gen7984)

							gen7986 := Call(__e, gen7983, gen7985, Parse__shen_4_5type_6)

							gen7987 := Call(__e, gen7982, gen7986)

							if True == gen7987 {
								gen7965 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

								gen7966 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen7967 := Call(__e, gen7966, Parse__shen_4_5type_6)

								gen7968 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen7969 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry)

								gen7970 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7971 := Call(__e, gen7970, Parse__shen_4_5expr_6)

								gen7972 := Call(__e, gen7969, gen7971)

								gen7973 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen7974 := Call(__e, PrimNS1Value(symns2_1value), symcons)

								gen7975 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demodulate)

								gen7976 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

								gen7977 := Call(__e, gen7976, Parse__shen_4_5type_6)

								gen7978 := Call(__e, gen7975, gen7977)

								gen7979 := Call(__e, gen7974, gen7978, Nil)

								gen7980 := Call(__e, gen7973, sym_h, gen7979)

								gen7981 := Call(__e, gen7968, gen7972, gen7980)

								__e.TailApply(gen7965, gen7967, gen7981)

								return

							} else {
								gen7964 := Call(__e, PrimNS1Value(symns2_1value), symfail)

								__e.TailApply(gen7964)

								return

							}

						}, 1)

						gen7989 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5type_6)

						gen7990 := Call(__e, gen7989, NewStream1758)

						__e.TailApply(gen7988, gen7990)

						return

					}, 1)

					gen7992 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen7993 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen7994 := Call(__e, gen7993, Parse__shen_4_5expr_6)

					gen7995 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen7996 := Call(__e, gen7995, Parse__shen_4_5expr_6)

					gen7997 := Call(__e, gen7992, gen7994, gen7996)

					__e.TailApply(gen7991, gen7997)

					return

				} else {
					gen7963 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen7963)

					return

				}

			} else {
				gen7962 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen7962)

				return

			}

		}, 1)

		gen8014 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

		gen8015 := Call(__e, gen8014, V1760)

		gen8016 := Call(__e, gen8013, gen8015)

		__e.TailApply(gen7961, gen8016)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5formula_6, gen8017)

	gen8035 := MakeNative(func(__e *ControlFlow) {
		V1762 := __e.Get(1)
		_ = V1762
		gen8032 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5expr_6 := __e.Get(1)
			_ = Parse__shen_4_5expr_6
			gen8026 := Call(__e, PrimNS1Value(symns2_1value), symnot)

			gen8027 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8028 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen8029 := Call(__e, gen8028)

			gen8030 := Call(__e, gen8027, gen8029, Parse__shen_4_5expr_6)

			gen8031 := Call(__e, gen8026, gen8030)

			if True == gen8031 {
				gen8019 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

				gen8020 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8021 := Call(__e, gen8020, Parse__shen_4_5expr_6)

				gen8022 := Call(__e, PrimNS1Value(symns2_1value), symshen_4curry_1type)

				gen8023 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				gen8024 := Call(__e, gen8023, Parse__shen_4_5expr_6)

				gen8025 := Call(__e, gen8022, gen8024)

				__e.TailApply(gen8019, gen8021, gen8025)

				return

			} else {
				gen8018 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				__e.TailApply(gen8018)

				return

			}

		}, 1)

		gen8033 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5expr_6)

		gen8034 := Call(__e, gen8033, V1762)

		__e.TailApply(gen8032, gen8034)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5type_6, gen8035)

	gen8056 := MakeNative(func(__e *ControlFlow) {
		V1764 := __e.Get(1)
		_ = V1764
		gen8052 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8053 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen8054 := Call(__e, gen8053, V1764)

		gen8055 := Call(__e, gen8052, gen8054)

		if True == gen8055 {
			gen8049 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen8047 := Call(__e, PrimNS1Value(symns2_1value), symshen_4doubleunderline_2)

				gen8048 := Call(__e, gen8047, Parse__X)

				if True == gen8048 {
					gen8038 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen8039 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8040 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen8041 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen8042 := Call(__e, gen8041, V1764)

					gen8043 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen8044 := Call(__e, gen8043, V1764)

					gen8045 := Call(__e, gen8040, gen8042, gen8044)

					gen8046 := Call(__e, gen8039, gen8045)

					__e.TailApply(gen8038, gen8046, Parse__X)

					return

				} else {
					gen8037 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen8037)

					return

				}

			}, 1)

			gen8050 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen8051 := Call(__e, gen8050, V1764)

			__e.TailApply(gen8049, gen8051)

			return

		} else {
			gen8036 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen8036)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5doubleunderline_6, gen8056)

	gen8077 := MakeNative(func(__e *ControlFlow) {
		V1766 := __e.Get(1)
		_ = V1766
		gen8073 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8074 := Call(__e, PrimNS1Value(symns2_1value), symhd)

		gen8075 := Call(__e, gen8074, V1766)

		gen8076 := Call(__e, gen8073, gen8075)

		if True == gen8076 {
			gen8070 := MakeNative(func(__e *ControlFlow) {
				Parse__X := __e.Get(1)
				_ = Parse__X
				gen8068 := Call(__e, PrimNS1Value(symns2_1value), symshen_4singleunderline_2)

				gen8069 := Call(__e, gen8068, Parse__X)

				if True == gen8069 {
					gen8059 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen8060 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8061 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pair)

					gen8062 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlhd)

					gen8063 := Call(__e, gen8062, V1766)

					gen8064 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

					gen8065 := Call(__e, gen8064, V1766)

					gen8066 := Call(__e, gen8061, gen8063, gen8065)

					gen8067 := Call(__e, gen8060, gen8066)

					__e.TailApply(gen8059, gen8067, Parse__X)

					return

				} else {
					gen8058 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					__e.TailApply(gen8058)

					return

				}

			}, 1)

			gen8071 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdhd)

			gen8072 := Call(__e, gen8071, V1766)

			__e.TailApply(gen8070, gen8072)

			return

		} else {
			gen8057 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen8057)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_5singleunderline_6, gen8077)

	gen8084 := MakeNative(func(__e *ControlFlow) {
		V1768 := __e.Get(1)
		_ = V1768
		gen8082 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		gen8083 := Call(__e, gen8082, V1768)

		if True == gen8083 {
			gen8078 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sh_2)

			gen8079 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			gen8080 := Call(__e, gen8079, V1768)

			gen8081 := Call(__e, gen8078, gen8080)

			if True == gen8081 {
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

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4singleunderline_2, gen8084)

	gen8096 := MakeNative(func(__e *ControlFlow) {
		V1770 := __e.Get(1)
		_ = V1770
		gen8094 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8095 := Call(__e, gen8094, MakeString("_"), V1770)

		if True == gen8095 {
			__e.Return(True)
			return
		} else {
			if True == True {
				gen8090 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen8091 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen8092 := Call(__e, gen8091, V1770, MakeNumber(0))

				gen8093 := Call(__e, gen8090, gen8092, MakeString("_"))

				if True == gen8093 {
					gen8086 := Call(__e, PrimNS1Value(symns2_1value), symshen_4sh_2)

					gen8087 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen8088 := Call(__e, gen8087, V1770)

					gen8089 := Call(__e, gen8086, gen8088)

					if True == gen8089 {
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

			} else {
				gen8085 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8085, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4sh_2, gen8096)

	gen8103 := MakeNative(func(__e *ControlFlow) {
		V1772 := __e.Get(1)
		_ = V1772
		gen8101 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		gen8102 := Call(__e, gen8101, V1772)

		if True == gen8102 {
			gen8097 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dh_2)

			gen8098 := Call(__e, PrimNS1Value(symns2_1value), symstr)

			gen8099 := Call(__e, gen8098, V1772)

			gen8100 := Call(__e, gen8097, gen8099)

			if True == gen8100 {
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

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4doubleunderline_2, gen8103)

	gen8115 := MakeNative(func(__e *ControlFlow) {
		V1774 := __e.Get(1)
		_ = V1774
		gen8113 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8114 := Call(__e, gen8113, MakeString("="), V1774)

		if True == gen8114 {
			__e.Return(True)
			return
		} else {
			if True == True {
				gen8109 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen8110 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen8111 := Call(__e, gen8110, V1774, MakeNumber(0))

				gen8112 := Call(__e, gen8109, gen8111, MakeString("="))

				if True == gen8112 {
					gen8105 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dh_2)

					gen8106 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen8107 := Call(__e, gen8106, V1774)

					gen8108 := Call(__e, gen8105, gen8107)

					if True == gen8108 {
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

			} else {
				gen8104 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8104, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4dh_2, gen8115)

	gen8121 := MakeNative(func(__e *ControlFlow) {
		V1777 := __e.Get(1)
		_ = V1777
		V1778 := __e.Get(2)
		_ = V1778
		gen8116 := Call(__e, PrimNS1Value(symns2_1value), symshen_4remember_1datatype)

		gen8117 := Call(__e, PrimNS1Value(symns2_1value), symshen_4s_1prolog)

		gen8118 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rules_1_6horn_1clauses)

		gen8119 := Call(__e, gen8118, V1777, V1778)

		gen8120 := Call(__e, gen8117, gen8119)

		__e.TailApply(gen8116, gen8120)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4process_1datatype, gen8121)

	gen8141 := MakeNative(func(__e *ControlFlow) {
		V1784 := __e.Get(1)
		_ = V1784
		gen8139 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8140 := Call(__e, gen8139, V1784)

		if True == gen8140 {
			gen8124 := Call(__e, PrimNS1Value(symns2_1value), symset)

			gen8125 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			gen8126 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8127 := Call(__e, gen8126, V1784)

			gen8128 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen8129 := Call(__e, gen8128, symshen_4_ddatatypes_d)

			gen8130 := Call(__e, gen8125, gen8127, gen8129)

			Call(__e, gen8124, symshen_4_ddatatypes_d, gen8130)

			gen8131 := Call(__e, PrimNS1Value(symns2_1value), symset)

			gen8132 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			gen8133 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8134 := Call(__e, gen8133, V1784)

			gen8135 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen8136 := Call(__e, gen8135, symshen_4_dalldatatypes_d)

			gen8137 := Call(__e, gen8132, gen8134, gen8136)

			Call(__e, gen8131, symshen_4_dalldatatypes_d, gen8137)

			gen8138 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(gen8138, V1784)

			return

		} else {
			if True == True {
				gen8123 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen8123, symshen_4remember_1datatype)

				return

			} else {
				gen8122 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8122, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4remember_1datatype, gen8141)

	gen8196 := MakeNative(func(__e *ControlFlow) {
		V1789 := __e.Get(1)
		_ = V1789
		V1790 := __e.Get(2)
		_ = V1790
		gen8194 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8195 := Call(__e, gen8194, Nil, V1790)

		if True == gen8195 {
			__e.Return(Nil)
			return
		} else {
			gen8191 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8192 := Call(__e, gen8191, V1790)

			var gen8193 Obj

			if True == gen8192 {
				gen8186 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

				gen8187 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8188 := Call(__e, gen8187, V1790)

				gen8189 := Call(__e, gen8186, gen8188)

				var gen8190 Obj

				if True == gen8189 {
					gen8180 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen8181 := Call(__e, PrimNS1Value(symns2_1value), symfst)

					gen8182 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8183 := Call(__e, gen8182, V1790)

					gen8184 := Call(__e, gen8181, gen8183)

					gen8185 := Call(__e, gen8180, symshen_4single, gen8184)

					if True == gen8185 {
						gen8190 = True
					} else {
						gen8190 = False
					}

				} else {
					gen8190 = False
				}

				if True == gen8190 {
					gen8193 = True
				} else {
					gen8193 = False
				}

			} else {
				gen8193 = False
			}

			if True == gen8193 {
				gen8169 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8170 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rule_1_6horn_1clause)

				gen8171 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				gen8172 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8173 := Call(__e, gen8172, V1790)

				gen8174 := Call(__e, gen8171, gen8173)

				gen8175 := Call(__e, gen8170, V1789, gen8174)

				gen8176 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rules_1_6horn_1clauses)

				gen8177 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8178 := Call(__e, gen8177, V1790)

				gen8179 := Call(__e, gen8176, V1789, gen8178)

				__e.TailApply(gen8169, gen8175, gen8179)

				return

			} else {
				gen8166 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8167 := Call(__e, gen8166, V1790)

				var gen8168 Obj

				if True == gen8167 {
					gen8161 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					gen8162 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8163 := Call(__e, gen8162, V1790)

					gen8164 := Call(__e, gen8161, gen8163)

					var gen8165 Obj

					if True == gen8164 {
						gen8155 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen8156 := Call(__e, PrimNS1Value(symns2_1value), symfst)

						gen8157 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8158 := Call(__e, gen8157, V1790)

						gen8159 := Call(__e, gen8156, gen8158)

						gen8160 := Call(__e, gen8155, symshen_4double, gen8159)

						if True == gen8160 {
							gen8165 = True
						} else {
							gen8165 = False
						}

					} else {
						gen8165 = False
					}

					if True == gen8165 {
						gen8168 = True
					} else {
						gen8168 = False
					}

				} else {
					gen8168 = False
				}

				if True == gen8168 {
					gen8144 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rules_1_6horn_1clauses)

					gen8145 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					gen8146 := Call(__e, PrimNS1Value(symns2_1value), symshen_4double_1_6singles)

					gen8147 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

					gen8148 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8149 := Call(__e, gen8148, V1790)

					gen8150 := Call(__e, gen8147, gen8149)

					gen8151 := Call(__e, gen8146, gen8150)

					gen8152 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8153 := Call(__e, gen8152, V1790)

					gen8154 := Call(__e, gen8145, gen8151, gen8153)

					__e.TailApply(gen8144, V1789, gen8154)

					return

				} else {
					if True == True {
						gen8143 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen8143, symshen_4rules_1_6horn_1clauses)

						return

					} else {
						gen8142 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen8142, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4rules_1_6horn_1clauses, gen8196)

	gen8204 := MakeNative(func(__e *ControlFlow) {
		V1792 := __e.Get(1)
		_ = V1792
		gen8197 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8198 := Call(__e, PrimNS1Value(symns2_1value), symshen_4right_1rule)

		gen8199 := Call(__e, gen8198, V1792)

		gen8200 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8201 := Call(__e, PrimNS1Value(symns2_1value), symshen_4left_1rule)

		gen8202 := Call(__e, gen8201, V1792)

		gen8203 := Call(__e, gen8200, gen8202, Nil)

		__e.TailApply(gen8197, gen8199, gen8203)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4double_1_6singles, gen8204)

	gen8206 := MakeNative(func(__e *ControlFlow) {
		V1794 := __e.Get(1)
		_ = V1794
		gen8205 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

		__e.TailApply(gen8205, symshen_4single, V1794)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4right_1rule, gen8206)

	gen8290 := MakeNative(func(__e *ControlFlow) {
		V1796 := __e.Get(1)
		_ = V1796
		gen8287 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8288 := Call(__e, gen8287, V1796)

		var gen8289 Obj

		if True == gen8288 {
			gen8282 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8283 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8284 := Call(__e, gen8283, V1796)

			gen8285 := Call(__e, gen8282, gen8284)

			var gen8286 Obj

			if True == gen8285 {
				gen8275 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8276 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8277 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8278 := Call(__e, gen8277, V1796)

				gen8279 := Call(__e, gen8276, gen8278)

				gen8280 := Call(__e, gen8275, gen8279)

				var gen8281 Obj

				if True == gen8280 {
					gen8266 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					gen8267 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8268 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8269 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8270 := Call(__e, gen8269, V1796)

					gen8271 := Call(__e, gen8268, gen8270)

					gen8272 := Call(__e, gen8267, gen8271)

					gen8273 := Call(__e, gen8266, gen8272)

					var gen8274 Obj

					if True == gen8273 {
						gen8255 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen8256 := Call(__e, PrimNS1Value(symns2_1value), symfst)

						gen8257 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8258 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8259 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8260 := Call(__e, gen8259, V1796)

						gen8261 := Call(__e, gen8258, gen8260)

						gen8262 := Call(__e, gen8257, gen8261)

						gen8263 := Call(__e, gen8256, gen8262)

						gen8264 := Call(__e, gen8255, Nil, gen8263)

						var gen8265 Obj

						if True == gen8264 {
							gen8247 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen8248 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8249 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8250 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8251 := Call(__e, gen8250, V1796)

							gen8252 := Call(__e, gen8249, gen8251)

							gen8253 := Call(__e, gen8248, gen8252)

							gen8254 := Call(__e, gen8247, Nil, gen8253)

							if True == gen8254 {
								gen8265 = True
							} else {
								gen8265 = False
							}

						} else {
							gen8265 = False
						}

						if True == gen8265 {
							gen8274 = True
						} else {
							gen8274 = False
						}

					} else {
						gen8274 = False
					}

					if True == gen8274 {
						gen8281 = True
					} else {
						gen8281 = False
					}

				} else {
					gen8281 = False
				}

				if True == gen8281 {
					gen8286 = True
				} else {
					gen8286 = False
				}

			} else {
				gen8286 = False
			}

			if True == gen8286 {
				gen8289 = True
			} else {
				gen8289 = False
			}

		} else {
			gen8289 = False
		}

		if True == gen8289 {
			gen8244 := MakeNative(func(__e *ControlFlow) {
				Q := __e.Get(1)
				_ = Q
				gen8231 := MakeNative(func(__e *ControlFlow) {
					NewConclusion := __e.Get(1)
					_ = NewConclusion
					gen8218 := MakeNative(func(__e *ControlFlow) {
						NewPremises := __e.Get(1)
						_ = NewPremises
						gen8209 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

						gen8210 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen8211 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8212 := Call(__e, gen8211, V1796)

						gen8213 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen8214 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen8215 := Call(__e, gen8214, NewConclusion, Nil)

						gen8216 := Call(__e, gen8213, NewPremises, gen8215)

						gen8217 := Call(__e, gen8210, gen8212, gen8216)

						__e.TailApply(gen8209, symshen_4single, gen8217)

						return

					}, 1)

					gen8219 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen8220 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

					gen8221 := Call(__e, PrimNS1Value(symns2_1value), symmap)

					gen8223 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						gen8222 := Call(__e, PrimNS1Value(symns2_1value), symshen_4right_1_6left)

						__e.TailApply(gen8222, X)

						return

					}, 1)

					gen8224 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8225 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8226 := Call(__e, gen8225, V1796)

					gen8227 := Call(__e, gen8224, gen8226)

					gen8228 := Call(__e, gen8221, gen8223, gen8227)

					gen8229 := Call(__e, gen8220, gen8228, Q)

					gen8230 := Call(__e, gen8219, gen8229, Nil)

					__e.TailApply(gen8218, gen8230)

					return

				}, 1)

				gen8232 := Call(__e, PrimNS1Value(symns2_1value), sym_8p)

				gen8233 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8234 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

				gen8235 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8236 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8237 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8238 := Call(__e, gen8237, V1796)

				gen8239 := Call(__e, gen8236, gen8238)

				gen8240 := Call(__e, gen8235, gen8239)

				gen8241 := Call(__e, gen8234, gen8240)

				gen8242 := Call(__e, gen8233, gen8241, Nil)

				gen8243 := Call(__e, gen8232, gen8242, Q)

				__e.TailApply(gen8231, gen8243)

				return

			}, 1)

			gen8245 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

			gen8246 := Call(__e, gen8245, symQv)

			__e.TailApply(gen8244, gen8246)

			return

		} else {
			if True == True {
				gen8208 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen8208, symshen_4left_1rule)

				return

			} else {
				gen8207 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8207, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4left_1rule, gen8290)

	gen8301 := MakeNative(func(__e *ControlFlow) {
		V1802 := __e.Get(1)
		_ = V1802
		gen8298 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

		gen8299 := Call(__e, gen8298, V1802)

		var gen8300 Obj

		if True == gen8299 {
			gen8294 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8295 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			gen8296 := Call(__e, gen8295, V1802)

			gen8297 := Call(__e, gen8294, Nil, gen8296)

			if True == gen8297 {
				gen8300 = True
			} else {
				gen8300 = False
			}

		} else {
			gen8300 = False
		}

		if True == gen8300 {
			gen8293 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			__e.TailApply(gen8293, V1802)

			return

		} else {
			if True == True {
				gen8292 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8292, MakeString("syntax error with ==========\n"))

				return

			} else {
				gen8291 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8291, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4right_1_6left, gen8301)

	gen8367 := MakeNative(func(__e *ControlFlow) {
		V1805 := __e.Get(1)
		_ = V1805
		V1806 := __e.Get(2)
		_ = V1806
		gen8364 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8365 := Call(__e, gen8364, V1806)

		var gen8366 Obj

		if True == gen8365 {
			gen8359 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8360 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8361 := Call(__e, gen8360, V1806)

			gen8362 := Call(__e, gen8359, gen8361)

			var gen8363 Obj

			if True == gen8362 {
				gen8352 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8353 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8354 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8355 := Call(__e, gen8354, V1806)

				gen8356 := Call(__e, gen8353, gen8355)

				gen8357 := Call(__e, gen8352, gen8356)

				var gen8358 Obj

				if True == gen8357 {
					gen8343 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

					gen8344 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8345 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8346 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8347 := Call(__e, gen8346, V1806)

					gen8348 := Call(__e, gen8345, gen8347)

					gen8349 := Call(__e, gen8344, gen8348)

					gen8350 := Call(__e, gen8343, gen8349)

					var gen8351 Obj

					if True == gen8350 {
						gen8335 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen8336 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8337 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8338 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8339 := Call(__e, gen8338, V1806)

						gen8340 := Call(__e, gen8337, gen8339)

						gen8341 := Call(__e, gen8336, gen8340)

						gen8342 := Call(__e, gen8335, Nil, gen8341)

						if True == gen8342 {
							gen8351 = True
						} else {
							gen8351 = False
						}

					} else {
						gen8351 = False
					}

					if True == gen8351 {
						gen8358 = True
					} else {
						gen8358 = False
					}

				} else {
					gen8358 = False
				}

				if True == gen8358 {
					gen8363 = True
				} else {
					gen8363 = False
				}

			} else {
				gen8363 = False
			}

			if True == gen8363 {
				gen8366 = True
			} else {
				gen8366 = False
			}

		} else {
			gen8366 = False
		}

		if True == gen8366 {
			gen8304 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8305 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rule_1_6horn_1clause_1head)

			gen8306 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			gen8307 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8308 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8309 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8310 := Call(__e, gen8309, V1806)

			gen8311 := Call(__e, gen8308, gen8310)

			gen8312 := Call(__e, gen8307, gen8311)

			gen8313 := Call(__e, gen8306, gen8312)

			gen8314 := Call(__e, gen8305, V1805, gen8313)

			gen8315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8317 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rule_1_6horn_1clause_1body)

			gen8318 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8319 := Call(__e, gen8318, V1806)

			gen8320 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8321 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8322 := Call(__e, gen8321, V1806)

			gen8323 := Call(__e, gen8320, gen8322)

			gen8324 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			gen8325 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8326 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8327 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8328 := Call(__e, gen8327, V1806)

			gen8329 := Call(__e, gen8326, gen8328)

			gen8330 := Call(__e, gen8325, gen8329)

			gen8331 := Call(__e, gen8324, gen8330)

			gen8332 := Call(__e, gen8317, gen8319, gen8323, gen8331)

			gen8333 := Call(__e, gen8316, gen8332, Nil)

			gen8334 := Call(__e, gen8315, sym_h_1, gen8333)

			__e.TailApply(gen8304, gen8314, gen8334)

			return

		} else {
			if True == True {
				gen8303 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen8303, symshen_4rule_1_6horn_1clause)

				return

			} else {
				gen8302 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8302, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause, gen8367)

	gen8375 := MakeNative(func(__e *ControlFlow) {
		V1809 := __e.Get(1)
		_ = V1809
		V1810 := __e.Get(2)
		_ = V1810
		gen8368 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8369 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8370 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mode_1ify)

		gen8371 := Call(__e, gen8370, V1810)

		gen8372 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8373 := Call(__e, gen8372, symContext__1957, Nil)

		gen8374 := Call(__e, gen8369, gen8371, gen8373)

		__e.TailApply(gen8368, V1809, gen8374)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause_1head, gen8375)

	gen8432 := MakeNative(func(__e *ControlFlow) {
		V1812 := __e.Get(1)
		_ = V1812
		gen8429 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8430 := Call(__e, gen8429, V1812)

		var gen8431 Obj

		if True == gen8430 {
			gen8424 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8425 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8426 := Call(__e, gen8425, V1812)

			gen8427 := Call(__e, gen8424, gen8426)

			var gen8428 Obj

			if True == gen8427 {
				gen8417 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen8418 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8419 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8420 := Call(__e, gen8419, V1812)

				gen8421 := Call(__e, gen8418, gen8420)

				gen8422 := Call(__e, gen8417, sym_h, gen8421)

				var gen8423 Obj

				if True == gen8422 {
					gen8410 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen8411 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8412 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8413 := Call(__e, gen8412, V1812)

					gen8414 := Call(__e, gen8411, gen8413)

					gen8415 := Call(__e, gen8410, gen8414)

					var gen8416 Obj

					if True == gen8415 {
						gen8402 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen8403 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8404 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8405 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8406 := Call(__e, gen8405, V1812)

						gen8407 := Call(__e, gen8404, gen8406)

						gen8408 := Call(__e, gen8403, gen8407)

						gen8409 := Call(__e, gen8402, Nil, gen8408)

						if True == gen8409 {
							gen8416 = True
						} else {
							gen8416 = False
						}

					} else {
						gen8416 = False
					}

					if True == gen8416 {
						gen8423 = True
					} else {
						gen8423 = False
					}

				} else {
					gen8423 = False
				}

				if True == gen8423 {
					gen8428 = True
				} else {
					gen8428 = False
				}

			} else {
				gen8428 = False
			}

			if True == gen8428 {
				gen8431 = True
			} else {
				gen8431 = False
			}

		} else {
			gen8431 = False
		}

		if True == gen8431 {
			gen8377 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8378 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8379 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8380 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8381 := Call(__e, gen8380, V1812)

			gen8382 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8383 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8384 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8385 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8386 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8387 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8388 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8389 := Call(__e, gen8388, V1812)

			gen8390 := Call(__e, gen8387, gen8389)

			gen8391 := Call(__e, gen8386, gen8390)

			gen8392 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8393 := Call(__e, gen8392, sym_7, Nil)

			gen8394 := Call(__e, gen8385, gen8391, gen8393)

			gen8395 := Call(__e, gen8384, symmode, gen8394)

			gen8396 := Call(__e, gen8383, gen8395, Nil)

			gen8397 := Call(__e, gen8382, sym_h, gen8396)

			gen8398 := Call(__e, gen8379, gen8381, gen8397)

			gen8399 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8400 := Call(__e, gen8399, sym_1, Nil)

			gen8401 := Call(__e, gen8378, gen8398, gen8400)

			__e.TailApply(gen8377, symmode, gen8401)

			return

		} else {
			if True == True {
				__e.Return(V1812)
				return
			} else {
				gen8376 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8376, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4mode_1ify, gen8432)

	gen8462 := MakeNative(func(__e *ControlFlow) {
		V1816 := __e.Get(1)
		_ = V1816
		V1817 := __e.Get(2)
		_ = V1817
		V1818 := __e.Get(3)
		_ = V1818
		gen8457 := MakeNative(func(__e *ControlFlow) {
			Variables := __e.Get(1)
			_ = Variables
			gen8452 := MakeNative(func(__e *ControlFlow) {
				Predicates := __e.Get(1)
				_ = Predicates
				gen8449 := MakeNative(func(__e *ControlFlow) {
					SearchLiterals := __e.Get(1)
					_ = SearchLiterals
					gen8446 := MakeNative(func(__e *ControlFlow) {
						SearchClauses := __e.Get(1)
						_ = SearchClauses
						gen8443 := MakeNative(func(__e *ControlFlow) {
							SideLiterals := __e.Get(1)
							_ = SideLiterals
							gen8436 := MakeNative(func(__e *ControlFlow) {
								PremissLiterals := __e.Get(1)
								_ = PremissLiterals
								gen8433 := Call(__e, PrimNS1Value(symns2_1value), symappend)

								gen8434 := Call(__e, PrimNS1Value(symns2_1value), symappend)

								gen8435 := Call(__e, gen8434, SideLiterals, PremissLiterals)

								__e.TailApply(gen8433, SearchLiterals, gen8435)

								return

							}, 1)

							gen8437 := Call(__e, PrimNS1Value(symns2_1value), symmap)

							gen8441 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								gen8438 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1premiss_1literal)

								gen8439 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

								gen8440 := Call(__e, gen8439, V1818)

								__e.TailApply(gen8438, X, gen8440)

								return

							}, 1)

							gen8442 := Call(__e, gen8437, gen8441, V1817)

							__e.TailApply(gen8436, gen8442)

							return

						}, 1)

						gen8444 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

						gen8445 := Call(__e, gen8444, V1816)

						__e.TailApply(gen8443, gen8445)

						return

					}, 1)

					gen8447 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1clauses)

					gen8448 := Call(__e, gen8447, Predicates, V1818, Variables)

					__e.TailApply(gen8446, gen8448)

					return

				}, 1)

				gen8450 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1literals)

				gen8451 := Call(__e, gen8450, Predicates, Variables, symContext__1957, symContext1__1957)

				__e.TailApply(gen8449, gen8451)

				return

			}, 1)

			gen8453 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			gen8455 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen8454 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				__e.TailApply(gen8454, symshen_4cl)

				return

			}, 1)

			gen8456 := Call(__e, gen8453, gen8455, V1818)

			__e.TailApply(gen8452, gen8456)

			return

		}, 1)

		gen8458 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen8460 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen8459 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

			__e.TailApply(gen8459, X)

			return

		}, 1)

		gen8461 := Call(__e, gen8458, gen8460, V1818)

		__e.TailApply(gen8457, gen8461)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4rule_1_6horn_1clause_1body, gen8462)

	gen8470 := MakeNative(func(__e *ControlFlow) {
		V1827 := __e.Get(1)
		_ = V1827
		V1828 := __e.Get(2)
		_ = V1828
		V1829 := __e.Get(3)
		_ = V1829
		V1830 := __e.Get(4)
		_ = V1830
		gen8467 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8468 := Call(__e, gen8467, Nil, V1827)

		var gen8469 Obj

		if True == gen8468 {
			gen8465 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8466 := Call(__e, gen8465, Nil, V1828)

			if True == gen8466 {
				gen8469 = True
			} else {
				gen8469 = False
			}

		} else {
			gen8469 = False
		}

		if True == gen8469 {
			__e.Return(Nil)
			return
		} else {
			if True == True {
				gen8464 := Call(__e, PrimNS1Value(symns2_1value), symshen_4csl_1help)

				__e.TailApply(gen8464, V1827, V1828, V1829, V1830)

				return

			} else {
				gen8463 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8463, MakeString("no cond match"))

				return

			}
		}

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1literals, gen8470)

	gen8509 := MakeNative(func(__e *ControlFlow) {
		V1837 := __e.Get(1)
		_ = V1837
		V1838 := __e.Get(2)
		_ = V1838
		V1839 := __e.Get(3)
		_ = V1839
		V1840 := __e.Get(4)
		_ = V1840
		gen8506 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8507 := Call(__e, gen8506, Nil, V1837)

		var gen8508 Obj

		if True == gen8507 {
			gen8504 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8505 := Call(__e, gen8504, Nil, V1838)

			if True == gen8505 {
				gen8508 = True
			} else {
				gen8508 = False
			}

		} else {
			gen8508 = False
		}

		if True == gen8508 {
			gen8497 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8498 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8499 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8500 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8501 := Call(__e, gen8500, V1839, Nil)

			gen8502 := Call(__e, gen8499, symContextOut__1957, gen8501)

			gen8503 := Call(__e, gen8498, symbind, gen8502)

			__e.TailApply(gen8497, gen8503, Nil)

			return

		} else {
			gen8494 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8495 := Call(__e, gen8494, V1837)

			var gen8496 Obj

			if True == gen8495 {
				gen8492 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8493 := Call(__e, gen8492, V1838)

				if True == gen8493 {
					gen8496 = True
				} else {
					gen8496 = False
				}

			} else {
				gen8496 = False
			}

			if True == gen8496 {
				gen8473 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8474 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8475 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8476 := Call(__e, gen8475, V1837)

				gen8477 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8478 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8479 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8480 := Call(__e, gen8479, V1838)

				gen8481 := Call(__e, gen8478, V1840, gen8480)

				gen8482 := Call(__e, gen8477, V1839, gen8481)

				gen8483 := Call(__e, gen8474, gen8476, gen8482)

				gen8484 := Call(__e, PrimNS1Value(symns2_1value), symshen_4csl_1help)

				gen8485 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8486 := Call(__e, gen8485, V1837)

				gen8487 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8488 := Call(__e, gen8487, V1838)

				gen8489 := Call(__e, PrimNS1Value(symns2_1value), symgensym)

				gen8490 := Call(__e, gen8489, symContext)

				gen8491 := Call(__e, gen8484, gen8486, gen8488, V1840, gen8490)

				__e.TailApply(gen8473, gen8483, gen8491)

				return

			} else {
				if True == True {
					gen8472 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen8472, symshen_4csl_1help)

					return

				} else {
					gen8471 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen8471, MakeString("no cond match"))

					return

				}
			}

		}

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4csl_1help, gen8509)

	gen8542 := MakeNative(func(__e *ControlFlow) {
		V1844 := __e.Get(1)
		_ = V1844
		V1845 := __e.Get(2)
		_ = V1845
		V1846 := __e.Get(3)
		_ = V1846
		gen8539 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8540 := Call(__e, gen8539, Nil, V1844)

		var gen8541 Obj

		if True == gen8540 {
			gen8536 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8537 := Call(__e, gen8536, Nil, V1845)

			var gen8538 Obj

			if True == gen8537 {
				gen8534 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen8535 := Call(__e, gen8534, Nil, V1846)

				if True == gen8535 {
					gen8538 = True
				} else {
					gen8538 = False
				}

			} else {
				gen8538 = False
			}

			if True == gen8538 {
				gen8541 = True
			} else {
				gen8541 = False
			}

		} else {
			gen8541 = False
		}

		if True == gen8541 {
			__e.Return(symshen_4skip)
			return
		} else {
			gen8531 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8532 := Call(__e, gen8531, V1844)

			var gen8533 Obj

			if True == gen8532 {
				gen8528 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8529 := Call(__e, gen8528, V1845)

				var gen8530 Obj

				if True == gen8529 {
					gen8526 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen8527 := Call(__e, gen8526, V1846)

					if True == gen8527 {
						gen8530 = True
					} else {
						gen8530 = False
					}

				} else {
					gen8530 = False
				}

				if True == gen8530 {
					gen8533 = True
				} else {
					gen8533 = False
				}

			} else {
				gen8533 = False
			}

			if True == gen8533 {
				gen8512 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1clause)

				gen8513 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8514 := Call(__e, gen8513, V1844)

				gen8515 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8516 := Call(__e, gen8515, V1845)

				gen8517 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8518 := Call(__e, gen8517, V1846)

				Call(__e, gen8512, gen8514, gen8516, gen8518)

				gen8519 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1search_1clauses)

				gen8520 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8521 := Call(__e, gen8520, V1844)

				gen8522 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8523 := Call(__e, gen8522, V1845)

				gen8524 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8525 := Call(__e, gen8524, V1846)

				__e.TailApply(gen8519, gen8521, gen8523, gen8525)

				return

			} else {
				if True == True {
					gen8511 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen8511, symshen_4construct_1search_1clauses)

					return

				} else {
					gen8510 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen8510, MakeString("no cond match"))

					return

				}
			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1clauses, gen8542)

	gen8552 := MakeNative(func(__e *ControlFlow) {
		V1850 := __e.Get(1)
		_ = V1850
		V1851 := __e.Get(2)
		_ = V1851
		V1852 := __e.Get(3)
		_ = V1852
		gen8543 := Call(__e, PrimNS1Value(symns2_1value), symshen_4s_1prolog)

		gen8544 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8545 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1base_1search_1clause)

		gen8546 := Call(__e, gen8545, V1850, V1851, V1852)

		gen8547 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8548 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1recursive_1search_1clause)

		gen8549 := Call(__e, gen8548, V1850, V1851, V1852)

		gen8550 := Call(__e, gen8547, gen8549, Nil)

		gen8551 := Call(__e, gen8544, gen8546, gen8550)

		__e.TailApply(gen8543, gen8551)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1search_1clause, gen8552)

	gen8568 := MakeNative(func(__e *ControlFlow) {
		V1856 := __e.Get(1)
		_ = V1856
		V1857 := __e.Get(2)
		_ = V1857
		V1858 := __e.Get(3)
		_ = V1858
		gen8553 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8554 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8555 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8556 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8557 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mode_1ify)

		gen8558 := Call(__e, gen8557, V1857)

		gen8559 := Call(__e, gen8556, gen8558, symIn__1957)

		gen8560 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8561 := Call(__e, gen8560, symIn__1957, V1858)

		gen8562 := Call(__e, gen8555, gen8559, gen8561)

		gen8563 := Call(__e, gen8554, V1856, gen8562)

		gen8564 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8565 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8566 := Call(__e, gen8565, Nil, Nil)

		gen8567 := Call(__e, gen8564, sym_h_1, gen8566)

		__e.TailApply(gen8553, gen8563, gen8567)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1base_1search_1clause, gen8568)

	gen8592 := MakeNative(func(__e *ControlFlow) {
		V1862 := __e.Get(1)
		_ = V1862
		V1863 := __e.Get(2)
		_ = V1863
		V1864 := __e.Get(3)
		_ = V1864
		gen8569 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8570 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8571 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8572 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8573 := Call(__e, gen8572, symAssumption__1957, symAssumptions__1957)

		gen8574 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8575 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8576 := Call(__e, gen8575, symAssumption__1957, symOut__1957)

		gen8577 := Call(__e, gen8574, gen8576, V1864)

		gen8578 := Call(__e, gen8571, gen8573, gen8577)

		gen8579 := Call(__e, gen8570, V1862, gen8578)

		gen8580 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8581 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8582 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8583 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8584 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8585 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen8586 := Call(__e, gen8585, symOut__1957, V1864)

		gen8587 := Call(__e, gen8584, symAssumptions__1957, gen8586)

		gen8588 := Call(__e, gen8583, V1862, gen8587)

		gen8589 := Call(__e, gen8582, gen8588, Nil)

		gen8590 := Call(__e, gen8581, gen8589, Nil)

		gen8591 := Call(__e, gen8580, sym_h_1, gen8590)

		__e.TailApply(gen8569, gen8579, gen8591)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1recursive_1search_1clause, gen8592)

	gen8695 := MakeNative(func(__e *ControlFlow) {
		V1870 := __e.Get(1)
		_ = V1870
		gen8693 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8694 := Call(__e, gen8693, Nil, V1870)

		if True == gen8694 {
			__e.Return(Nil)
			return
		} else {
			gen8690 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8691 := Call(__e, gen8690, V1870)

			var gen8692 Obj

			if True == gen8691 {
				gen8685 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8686 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8687 := Call(__e, gen8686, V1870)

				gen8688 := Call(__e, gen8685, gen8687)

				var gen8689 Obj

				if True == gen8688 {
					gen8678 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen8679 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8680 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8681 := Call(__e, gen8680, V1870)

					gen8682 := Call(__e, gen8679, gen8681)

					gen8683 := Call(__e, gen8678, symif, gen8682)

					var gen8684 Obj

					if True == gen8683 {
						gen8671 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen8672 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8673 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8674 := Call(__e, gen8673, V1870)

						gen8675 := Call(__e, gen8672, gen8674)

						gen8676 := Call(__e, gen8671, gen8675)

						var gen8677 Obj

						if True == gen8676 {
							gen8663 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen8664 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8665 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8666 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen8667 := Call(__e, gen8666, V1870)

							gen8668 := Call(__e, gen8665, gen8667)

							gen8669 := Call(__e, gen8664, gen8668)

							gen8670 := Call(__e, gen8663, Nil, gen8669)

							if True == gen8670 {
								gen8677 = True
							} else {
								gen8677 = False
							}

						} else {
							gen8677 = False
						}

						if True == gen8677 {
							gen8684 = True
						} else {
							gen8684 = False
						}

					} else {
						gen8684 = False
					}

					if True == gen8684 {
						gen8689 = True
					} else {
						gen8689 = False
					}

				} else {
					gen8689 = False
				}

				if True == gen8689 {
					gen8692 = True
				} else {
					gen8692 = False
				}

			} else {
				gen8692 = False
			}

			if True == gen8692 {
				gen8652 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8653 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8654 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8655 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8656 := Call(__e, gen8655, V1870)

				gen8657 := Call(__e, gen8654, gen8656)

				gen8658 := Call(__e, gen8653, symwhen, gen8657)

				gen8659 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

				gen8660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8661 := Call(__e, gen8660, V1870)

				gen8662 := Call(__e, gen8659, gen8661)

				__e.TailApply(gen8652, gen8658, gen8662)

				return

			} else {
				gen8649 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8650 := Call(__e, gen8649, V1870)

				var gen8651 Obj

				if True == gen8650 {
					gen8644 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen8645 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8646 := Call(__e, gen8645, V1870)

					gen8647 := Call(__e, gen8644, gen8646)

					var gen8648 Obj

					if True == gen8647 {
						gen8637 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen8638 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8639 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8640 := Call(__e, gen8639, V1870)

						gen8641 := Call(__e, gen8638, gen8640)

						gen8642 := Call(__e, gen8637, symlet, gen8641)

						var gen8643 Obj

						if True == gen8642 {
							gen8630 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

							gen8631 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8632 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen8633 := Call(__e, gen8632, V1870)

							gen8634 := Call(__e, gen8631, gen8633)

							gen8635 := Call(__e, gen8630, gen8634)

							var gen8636 Obj

							if True == gen8635 {
								gen8621 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen8622 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen8623 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen8624 := Call(__e, PrimNS1Value(symns2_1value), symhd)

								gen8625 := Call(__e, gen8624, V1870)

								gen8626 := Call(__e, gen8623, gen8625)

								gen8627 := Call(__e, gen8622, gen8626)

								gen8628 := Call(__e, gen8621, gen8627)

								var gen8629 Obj

								if True == gen8628 {
									gen8611 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen8612 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen8613 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen8614 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen8615 := Call(__e, PrimNS1Value(symns2_1value), symhd)

									gen8616 := Call(__e, gen8615, V1870)

									gen8617 := Call(__e, gen8614, gen8616)

									gen8618 := Call(__e, gen8613, gen8617)

									gen8619 := Call(__e, gen8612, gen8618)

									gen8620 := Call(__e, gen8611, Nil, gen8619)

									if True == gen8620 {
										gen8629 = True
									} else {
										gen8629 = False
									}

								} else {
									gen8629 = False
								}

								if True == gen8629 {
									gen8636 = True
								} else {
									gen8636 = False
								}

							} else {
								gen8636 = False
							}

							if True == gen8636 {
								gen8643 = True
							} else {
								gen8643 = False
							}

						} else {
							gen8643 = False
						}

						if True == gen8643 {
							gen8648 = True
						} else {
							gen8648 = False
						}

					} else {
						gen8648 = False
					}

					if True == gen8648 {
						gen8651 = True
					} else {
						gen8651 = False
					}

				} else {
					gen8651 = False
				}

				if True == gen8651 {
					gen8600 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen8601 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen8602 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8603 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8604 := Call(__e, gen8603, V1870)

					gen8605 := Call(__e, gen8602, gen8604)

					gen8606 := Call(__e, gen8601, symis, gen8605)

					gen8607 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

					gen8608 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8609 := Call(__e, gen8608, V1870)

					gen8610 := Call(__e, gen8607, gen8609)

					__e.TailApply(gen8600, gen8606, gen8610)

					return

				} else {
					gen8598 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen8599 := Call(__e, gen8598, V1870)

					if True == gen8599 {
						gen8595 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1side_1literals)

						gen8596 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8597 := Call(__e, gen8596, V1870)

						__e.TailApply(gen8595, gen8597)

						return

					} else {
						if True == True {
							gen8594 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen8594, symshen_4construct_1side_1literals)

							return

						} else {
							gen8593 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen8593, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1side_1literals, gen8695)

	gen8718 := MakeNative(func(__e *ControlFlow) {
		V1877 := __e.Get(1)
		_ = V1877
		V1878 := __e.Get(2)
		_ = V1878
		gen8716 := Call(__e, PrimNS1Value(symns2_1value), symtuple_2)

		gen8717 := Call(__e, gen8716, V1877)

		if True == gen8717 {
			gen8703 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8704 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8705 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

			gen8706 := Call(__e, PrimNS1Value(symns2_1value), symsnd)

			gen8707 := Call(__e, gen8706, V1877)

			gen8708 := Call(__e, gen8705, gen8707)

			gen8709 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8710 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1context)

			gen8711 := Call(__e, PrimNS1Value(symns2_1value), symfst)

			gen8712 := Call(__e, gen8711, V1877)

			gen8713 := Call(__e, gen8710, V1878, gen8712)

			gen8714 := Call(__e, gen8709, gen8713, Nil)

			gen8715 := Call(__e, gen8704, gen8708, gen8714)

			__e.TailApply(gen8703, symshen_4t_d, gen8715)

			return

		} else {
			gen8701 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8702 := Call(__e, gen8701, sym_b, V1877)

			if True == gen8702 {
				gen8698 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8699 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen8700 := Call(__e, gen8699, symThrowcontrol, Nil)

				__e.TailApply(gen8698, symcut, gen8700)

				return

			} else {
				if True == True {
					gen8697 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen8697, symshen_4construct_1premiss_1literal)

					return

				} else {
					gen8696 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen8696, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1premiss_1literal, gen8718)

	gen8746 := MakeNative(func(__e *ControlFlow) {
		V1881 := __e.Get(1)
		_ = V1881
		V1882 := __e.Get(2)
		_ = V1882
		gen8743 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8744 := Call(__e, gen8743, True, V1881)

		var gen8745 Obj

		if True == gen8744 {
			gen8741 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8742 := Call(__e, gen8741, Nil, V1882)

			if True == gen8742 {
				gen8745 = True
			} else {
				gen8745 = False
			}

		} else {
			gen8745 = False
		}

		if True == gen8745 {
			__e.Return(symContext__1957)
			return
		} else {
			gen8738 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8739 := Call(__e, gen8738, False, V1881)

			var gen8740 Obj

			if True == gen8739 {
				gen8736 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen8737 := Call(__e, gen8736, Nil, V1882)

				if True == gen8737 {
					gen8740 = True
				} else {
					gen8740 = False
				}

			} else {
				gen8740 = False
			}

			if True == gen8740 {
				__e.Return(symContextOut__1957)
				return
			} else {
				gen8734 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8735 := Call(__e, gen8734, V1882)

				if True == gen8735 {
					gen8721 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen8722 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen8723 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

					gen8724 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen8725 := Call(__e, gen8724, V1882)

					gen8726 := Call(__e, gen8723, gen8725)

					gen8727 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen8728 := Call(__e, PrimNS1Value(symns2_1value), symshen_4construct_1context)

					gen8729 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8730 := Call(__e, gen8729, V1882)

					gen8731 := Call(__e, gen8728, V1881, gen8730)

					gen8732 := Call(__e, gen8727, gen8731, Nil)

					gen8733 := Call(__e, gen8722, gen8726, gen8732)

					__e.TailApply(gen8721, symcons, gen8733)

					return

				} else {
					if True == True {
						gen8720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen8720, symshen_4construct_1context)

						return

					} else {
						gen8719 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen8719, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4construct_1context, gen8746)

	gen8763 := MakeNative(func(__e *ControlFlow) {
		V1884 := __e.Get(1)
		_ = V1884
		gen8761 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8762 := Call(__e, gen8761, V1884)

		if True == gen8762 {
			gen8748 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8749 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8750 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

			gen8751 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8752 := Call(__e, gen8751, V1884)

			gen8753 := Call(__e, gen8750, gen8752)

			gen8754 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8755 := Call(__e, PrimNS1Value(symns2_1value), symshen_4recursive__cons__form)

			gen8756 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8757 := Call(__e, gen8756, V1884)

			gen8758 := Call(__e, gen8755, gen8757)

			gen8759 := Call(__e, gen8754, gen8758, Nil)

			gen8760 := Call(__e, gen8749, gen8753, gen8759)

			__e.TailApply(gen8748, symcons, gen8760)

			return

		} else {
			if True == True {
				__e.Return(V1884)
				return
			} else {
				gen8747 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8747, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4recursive__cons__form, gen8763)

	gen8769 := MakeNative(func(__e *ControlFlow) {
		V1886 := __e.Get(1)
		_ = V1886
		gen8764 := Call(__e, PrimNS1Value(symns2_1value), symshen_4preclude_1h)

		gen8765 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen8767 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen8766 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(gen8766, X)

			return

		}, 1)

		gen8768 := Call(__e, gen8765, gen8767, V1886)

		__e.TailApply(gen8764, gen8768)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), sympreclude, gen8769)

	gen8778 := MakeNative(func(__e *ControlFlow) {
		V1888 := __e.Get(1)
		_ = V1888
		gen8771 := MakeNative(func(__e *ControlFlow) {
			FilterDatatypes := __e.Get(1)
			_ = FilterDatatypes
			gen8770 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			__e.TailApply(gen8770, symshen_4_ddatatypes_d)

			return

		}, 1)

		gen8772 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen8773 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

		gen8774 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen8775 := Call(__e, gen8774, symshen_4_ddatatypes_d)

		gen8776 := Call(__e, gen8773, gen8775, V1888)

		gen8777 := Call(__e, gen8772, symshen_4_ddatatypes_d, gen8776)

		__e.TailApply(gen8771, gen8777)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4preclude_1h, gen8778)

	gen8784 := MakeNative(func(__e *ControlFlow) {
		V1890 := __e.Get(1)
		_ = V1890
		gen8779 := Call(__e, PrimNS1Value(symns2_1value), symshen_4include_1h)

		gen8780 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen8782 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen8781 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(gen8781, X)

			return

		}, 1)

		gen8783 := Call(__e, gen8780, gen8782, V1890)

		__e.TailApply(gen8779, gen8783)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), syminclude, gen8784)

	gen8798 := MakeNative(func(__e *ControlFlow) {
		V1892 := __e.Get(1)
		_ = V1892
		gen8793 := MakeNative(func(__e *ControlFlow) {
			ValidTypes := __e.Get(1)
			_ = ValidTypes
			gen8786 := MakeNative(func(__e *ControlFlow) {
				NewDatatypes := __e.Get(1)
				_ = NewDatatypes
				gen8785 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				__e.TailApply(gen8785, symshen_4_ddatatypes_d)

				return

			}, 1)

			gen8787 := Call(__e, PrimNS1Value(symns2_1value), symset)

			gen8788 := Call(__e, PrimNS1Value(symns2_1value), symunion)

			gen8789 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen8790 := Call(__e, gen8789, symshen_4_ddatatypes_d)

			gen8791 := Call(__e, gen8788, ValidTypes, gen8790)

			gen8792 := Call(__e, gen8787, symshen_4_ddatatypes_d, gen8791)

			__e.TailApply(gen8786, gen8792)

			return

		}, 1)

		gen8794 := Call(__e, PrimNS1Value(symns2_1value), symintersection)

		gen8795 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen8796 := Call(__e, gen8795, symshen_4_dalldatatypes_d)

		gen8797 := Call(__e, gen8794, V1892, gen8796)

		__e.TailApply(gen8793, gen8797)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4include_1h, gen8798)

	gen8808 := MakeNative(func(__e *ControlFlow) {
		V1894 := __e.Get(1)
		_ = V1894
		gen8799 := Call(__e, PrimNS1Value(symns2_1value), symshen_4preclude_1h)

		gen8800 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

		gen8801 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen8802 := Call(__e, gen8801, symshen_4_dalldatatypes_d)

		gen8803 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen8805 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen8804 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(gen8804, X)

			return

		}, 1)

		gen8806 := Call(__e, gen8803, gen8805, V1894)

		gen8807 := Call(__e, gen8800, gen8802, gen8806)

		__e.TailApply(gen8799, gen8807)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), sympreclude_1all_1but, gen8808)

	gen8818 := MakeNative(func(__e *ControlFlow) {
		V1896 := __e.Get(1)
		_ = V1896
		gen8809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4include_1h)

		gen8810 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

		gen8811 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen8812 := Call(__e, gen8811, symshen_4_dalldatatypes_d)

		gen8813 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen8815 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen8814 := Call(__e, PrimNS1Value(symns2_1value), symshen_4intern_1type)

			__e.TailApply(gen8814, X)

			return

		}, 1)

		gen8816 := Call(__e, gen8813, gen8815, V1896)

		gen8817 := Call(__e, gen8810, gen8812, gen8816)

		__e.TailApply(gen8809, gen8817)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), syminclude_1all_1but, gen8818)

	gen8875 := MakeNative(func(__e *ControlFlow) {
		V1902 := __e.Get(1)
		_ = V1902
		gen8873 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen8874 := Call(__e, gen8873, Nil, V1902)

		if True == gen8874 {
			gen8864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4update_1demodulation_1function)

			gen8865 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen8866 := Call(__e, gen8865, symshen_4_dtc_d)

			gen8867 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

			gen8869 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen8868 := Call(__e, PrimNS1Value(symns2_1value), symshen_4demod_1rule)

				__e.TailApply(gen8868, X)

				return

			}, 1)

			gen8870 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen8871 := Call(__e, gen8870, symshen_4_dsynonyms_d)

			gen8872 := Call(__e, gen8867, gen8869, gen8871)

			__e.TailApply(gen8864, gen8866, gen8872)

			return

		} else {
			gen8861 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8862 := Call(__e, gen8861, V1902)

			var gen8863 Obj

			if True == gen8862 {
				gen8857 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8858 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8859 := Call(__e, gen8858, V1902)

				gen8860 := Call(__e, gen8857, gen8859)

				if True == gen8860 {
					gen8863 = True
				} else {
					gen8863 = False
				}

			} else {
				gen8863 = False
			}

			if True == gen8863 {
				gen8844 := MakeNative(func(__e *ControlFlow) {
					Vs := __e.Get(1)
					_ = Vs
					gen8842 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

					gen8843 := Call(__e, gen8842, Vs)

					if True == gen8843 {
						gen8826 := Call(__e, PrimNS1Value(symns2_1value), symshen_4pushnew)

						gen8827 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen8828 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8829 := Call(__e, gen8828, V1902)

						gen8830 := Call(__e, PrimNS1Value(symns2_1value), symcons)

						gen8831 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8832 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8833 := Call(__e, gen8832, V1902)

						gen8834 := Call(__e, gen8831, gen8833)

						gen8835 := Call(__e, gen8830, gen8834, Nil)

						gen8836 := Call(__e, gen8827, gen8829, gen8835)

						Call(__e, gen8826, gen8836, symshen_4_dsynonyms_d)

						gen8837 := Call(__e, PrimNS1Value(symns2_1value), symshen_4synonyms_1help)

						gen8838 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8839 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8840 := Call(__e, gen8839, V1902)

						gen8841 := Call(__e, gen8838, gen8840)

						__e.TailApply(gen8837, gen8841)

						return

					} else {
						gen8821 := Call(__e, PrimNS1Value(symns2_1value), symshen_4free__variable__warnings)

						gen8822 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8823 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8824 := Call(__e, gen8823, V1902)

						gen8825 := Call(__e, gen8822, gen8824)

						__e.TailApply(gen8821, gen8825, Vs)

						return

					}

				}, 1)

				gen8845 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

				gen8846 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				gen8847 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8848 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8849 := Call(__e, gen8848, V1902)

				gen8850 := Call(__e, gen8847, gen8849)

				gen8851 := Call(__e, gen8846, gen8850)

				gen8852 := Call(__e, PrimNS1Value(symns2_1value), symshen_4extract__vars)

				gen8853 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen8854 := Call(__e, gen8853, V1902)

				gen8855 := Call(__e, gen8852, gen8854)

				gen8856 := Call(__e, gen8845, gen8851, gen8855)

				__e.TailApply(gen8844, gen8856)

				return

			} else {
				if True == True {
					gen8820 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen8820, MakeString("odd number of synonyms\n"))

					return

				} else {
					gen8819 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen8819, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4synonyms_1help, gen8875)

	gen8886 := MakeNative(func(__e *ControlFlow) {
		V1905 := __e.Get(1)
		_ = V1905
		V1906 := __e.Get(2)
		_ = V1906
		gen8882 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen8883 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen8884 := Call(__e, gen8883, V1906)

		gen8885 := Call(__e, gen8882, V1905, gen8884)

		if True == gen8885 {
			gen8881 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			__e.TailApply(gen8881, V1906)

			return

		} else {
			gen8876 := Call(__e, PrimNS1Value(symns2_1value), symset)

			gen8877 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8878 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen8879 := Call(__e, gen8878, V1906)

			gen8880 := Call(__e, gen8877, V1905, gen8879)

			__e.TailApply(gen8876, V1906, gen8880)

			return

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4pushnew, gen8886)

	gen8918 := MakeNative(func(__e *ControlFlow) {
		V1908 := __e.Get(1)
		_ = V1908
		gen8915 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8916 := Call(__e, gen8915, V1908)

		var gen8917 Obj

		if True == gen8916 {
			gen8910 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen8911 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8912 := Call(__e, gen8911, V1908)

			gen8913 := Call(__e, gen8910, gen8912)

			var gen8914 Obj

			if True == gen8913 {
				gen8904 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen8905 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8906 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8907 := Call(__e, gen8906, V1908)

				gen8908 := Call(__e, gen8905, gen8907)

				gen8909 := Call(__e, gen8904, Nil, gen8908)

				if True == gen8909 {
					gen8914 = True
				} else {
					gen8914 = False
				}

			} else {
				gen8914 = False
			}

			if True == gen8914 {
				gen8917 = True
			} else {
				gen8917 = False
			}

		} else {
			gen8917 = False
		}

		if True == gen8917 {
			gen8889 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8890 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			gen8891 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8892 := Call(__e, gen8891, V1908)

			gen8893 := Call(__e, gen8890, gen8892)

			gen8894 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8895 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8896 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			gen8897 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8898 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8899 := Call(__e, gen8898, V1908)

			gen8900 := Call(__e, gen8897, gen8899)

			gen8901 := Call(__e, gen8896, gen8900)

			gen8902 := Call(__e, gen8895, gen8901, Nil)

			gen8903 := Call(__e, gen8894, sym_1_6, gen8902)

			__e.TailApply(gen8889, gen8893, gen8903)

			return

		} else {
			if True == True {
				gen8888 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen8888, symshen_4demod_1rule)

				return

			} else {
				gen8887 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8887, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4demod_1rule, gen8918)

	gen8999 := MakeNative(func(__e *ControlFlow) {
		V1914 := __e.Get(1)
		_ = V1914
		gen8996 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen8997 := Call(__e, gen8996, V1914)

		var gen8998 Obj

		if True == gen8997 {
			gen8991 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen8992 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8993 := Call(__e, gen8992, V1914)

			gen8994 := Call(__e, gen8991, symdefun, gen8993)

			var gen8995 Obj

			if True == gen8994 {
				gen8986 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen8987 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen8988 := Call(__e, gen8987, V1914)

				gen8989 := Call(__e, gen8986, gen8988)

				var gen8990 Obj

				if True == gen8989 {
					gen8979 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen8980 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8981 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen8982 := Call(__e, gen8981, V1914)

					gen8983 := Call(__e, gen8980, gen8982)

					gen8984 := Call(__e, gen8979, gen8983)

					var gen8985 Obj

					if True == gen8984 {
						gen8970 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen8971 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen8972 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8973 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen8974 := Call(__e, gen8973, V1914)

						gen8975 := Call(__e, gen8972, gen8974)

						gen8976 := Call(__e, gen8971, gen8975)

						gen8977 := Call(__e, gen8970, gen8976)

						var gen8978 Obj

						if True == gen8977 {
							gen8959 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen8960 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8961 := Call(__e, PrimNS1Value(symns2_1value), symhd)

							gen8962 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8963 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen8964 := Call(__e, gen8963, V1914)

							gen8965 := Call(__e, gen8962, gen8964)

							gen8966 := Call(__e, gen8961, gen8965)

							gen8967 := Call(__e, gen8960, gen8966)

							gen8968 := Call(__e, gen8959, Nil, gen8967)

							var gen8969 Obj

							if True == gen8968 {
								gen8950 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

								gen8951 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen8952 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen8953 := Call(__e, PrimNS1Value(symns2_1value), symtl)

								gen8954 := Call(__e, gen8953, V1914)

								gen8955 := Call(__e, gen8952, gen8954)

								gen8956 := Call(__e, gen8951, gen8955)

								gen8957 := Call(__e, gen8950, gen8956)

								var gen8958 Obj

								if True == gen8957 {
									gen8940 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

									gen8941 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen8942 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen8943 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen8944 := Call(__e, PrimNS1Value(symns2_1value), symtl)

									gen8945 := Call(__e, gen8944, V1914)

									gen8946 := Call(__e, gen8943, gen8945)

									gen8947 := Call(__e, gen8942, gen8946)

									gen8948 := Call(__e, gen8941, gen8947)

									gen8949 := Call(__e, gen8940, Nil, gen8948)

									if True == gen8949 {
										gen8958 = True
									} else {
										gen8958 = False
									}

								} else {
									gen8958 = False
								}

								if True == gen8958 {
									gen8969 = True
								} else {
									gen8969 = False
								}

							} else {
								gen8969 = False
							}

							if True == gen8969 {
								gen8978 = True
							} else {
								gen8978 = False
							}

						} else {
							gen8978 = False
						}

						if True == gen8978 {
							gen8985 = True
						} else {
							gen8985 = False
						}

					} else {
						gen8985 = False
					}

					if True == gen8985 {
						gen8990 = True
					} else {
						gen8990 = False
					}

				} else {
					gen8990 = False
				}

				if True == gen8990 {
					gen8995 = True
				} else {
					gen8995 = False
				}

			} else {
				gen8995 = False
			}

			if True == gen8995 {
				gen8998 = True
			} else {
				gen8998 = False
			}

		} else {
			gen8998 = False
		}

		if True == gen8998 {
			gen8921 := Call(__e, PrimNS1Value(symns2_1value), symeval)

			gen8922 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8923 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen8924 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8925 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen8926 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8927 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8928 := Call(__e, gen8927, V1914)

			gen8929 := Call(__e, gen8926, gen8928)

			gen8930 := Call(__e, gen8925, gen8929)

			gen8931 := Call(__e, gen8924, gen8930)

			gen8932 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8933 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8934 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen8935 := Call(__e, gen8934, V1914)

			gen8936 := Call(__e, gen8933, gen8935)

			gen8937 := Call(__e, gen8932, gen8936)

			gen8938 := Call(__e, gen8923, gen8931, gen8937)

			gen8939 := Call(__e, gen8922, sym_c_4, gen8938)

			__e.TailApply(gen8921, gen8939)

			return

		} else {
			if True == True {
				gen8920 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen8920, symshen_4lambda_1of_1defun)

				return

			} else {
				gen8919 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen8919, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4lambda_1of_1defun, gen8999)

	gen9015 := MakeNative(func(__e *ControlFlow) {
		V1917 := __e.Get(1)
		_ = V1917
		V1918 := __e.Get(2)
		_ = V1918
		gen9000 := Call(__e, PrimNS1Value(symns2_1value), symtc)

		Call(__e, gen9000, sym_1)

		gen9001 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen9002 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lambda_1of_1defun)

		gen9003 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

		gen9004 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9005 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9006 := Call(__e, PrimNS1Value(symns2_1value), symappend)

		gen9007 := Call(__e, PrimNS1Value(symns2_1value), symshen_4default_1rule)

		gen9008 := Call(__e, gen9007)

		gen9009 := Call(__e, gen9006, V1918, gen9008)

		gen9010 := Call(__e, gen9005, symshen_4demod, gen9009)

		gen9011 := Call(__e, gen9004, symdefine, gen9010)

		gen9012 := Call(__e, gen9003, gen9011)

		gen9013 := Call(__e, gen9002, gen9012)

		Call(__e, gen9001, symshen_4_ddemodulation_1function_d, gen9013)

		if True == V1917 {
			gen9014 := Call(__e, PrimNS1Value(symns2_1value), symtc)

			Call(__e, gen9014, sym_7)

		} else {
			_ = symshen_4skip
		}

		__e.Return(symsynonyms)
		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4update_1demodulation_1function, gen9015)

	gen9021 := MakeNative(func(__e *ControlFlow) {
		gen9016 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9017 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9018 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen9019 := Call(__e, gen9018, symX, Nil)

		gen9020 := Call(__e, gen9017, sym_1_6, gen9019)

		__e.TailApply(gen9016, symX, gen9020)

		return

	}, 0)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4default_1rule, gen9021)

	return

}, 0)
