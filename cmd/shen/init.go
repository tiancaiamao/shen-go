package main

import . "github.com/tiancaiamao/shen-go/kl"

var InitMain = MakeNative(func(__e *ControlFlow) {
	tmp75 := MakeNative(func(__e *ControlFlow) {
		tmp76 := PrimSet(symshen_4_dhistory_d, Nil)

		_ = tmp76

		tmp77 := PrimSet(symshen_4_dtc_d, False)

		_ = tmp77

		tmp78 := Call(__e, PrimFunc(symshen_4dict), MakeNumber(20000))

		tmp79 := PrimSet(sym_dproperty_1vector_d, tmp78)

		_ = tmp79

		tmp80 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimFunc(symshen_4macros), X)
			return
		}, 1)

		tmp81 := PrimCons(symshen_4macros, tmp80)

		tmp82 := PrimCons(tmp81, Nil)

		tmp83 := PrimSet(sym_dmacros_d, tmp82)

		_ = tmp83

		tmp84 := PrimSet(symshen_4_dgensym_d, MakeNumber(0))

		_ = tmp84

		tmp85 := PrimSet(symshen_4_dtracking_d, Nil)

		_ = tmp85

		tmp86 := PrimSet(symshen_4_dprofiled_d, Nil)

		_ = tmp86

		tmp87 := PrimCons(symtype, Nil)

		tmp88 := PrimCons(syminput_7, tmp87)

		tmp89 := PrimCons(symopen, tmp88)

		tmp90 := PrimCons(symset, tmp89)

		tmp91 := PrimCons(symwhere, tmp90)

		tmp92 := PrimCons(symlet, tmp91)

		tmp93 := PrimCons(symlambda, tmp92)

		tmp94 := PrimCons(symcons, tmp93)

		tmp95 := PrimCons(sym_8v, tmp94)

		tmp96 := PrimCons(sym_8s, tmp95)

		tmp97 := PrimCons(sym_8p, tmp96)

		tmp98 := PrimSet(symshen_4_dspecial_d, tmp97)

		_ = tmp98

		tmp99 := PrimSet(symshen_4_dextraspecial_d, Nil)

		_ = tmp99

		tmp100 := PrimSet(symshen_4_dspy_d, False)

		_ = tmp100

		tmp101 := PrimSet(symshen_4_ddatatypes_d, Nil)

		_ = tmp101

		tmp102 := PrimSet(symshen_4_dalldatatypes_d, Nil)

		_ = tmp102

		tmp103 := PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

		_ = tmp103

		tmp104 := PrimSet(symshen_4_dpackage_d, symnull)

		_ = tmp104

		tmp105 := PrimSet(symshen_4_dsynonyms_d, Nil)

		_ = tmp105

		tmp106 := PrimSet(symshen_4_dsystem_d, Nil)

		_ = tmp106

		tmp107 := PrimSet(symshen_4_doccurs_d, True)

		_ = tmp107

		tmp108 := PrimSet(symshen_4_dfactorise_2_d, False)

		_ = tmp108

		tmp109 := PrimSet(symshen_4_dmaxinferences_d, MakeNumber(1000000))

		_ = tmp109

		tmp110 := PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

		_ = tmp110

		tmp111 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

		_ = tmp111

		tmp112 := PrimSet(symshen_4_dinfs_d, MakeNumber(0))

		_ = tmp112

		tmp113 := PrimSet(sym_dhush_d, False)

		_ = tmp113

		tmp114 := PrimSet(symshen_4_doptimise_d, False)

		_ = tmp114

		tmp115 := PrimSet(sym_dversion_d, MakeString("41.1"))

		_ = tmp115

		tmp116 := PrimSet(symshen_4_dnames_d, Nil)

		_ = tmp116

		tmp117 := PrimSet(symshen_4_dstep_d, False)

		_ = tmp117

		tmp118 := PrimSet(symshen_4_dit_d, MakeString(""))

		_ = tmp118

		tmp119 := PrimSet(symshen_4_dresidue_d, Nil)

		_ = tmp119

		tmp120 := PrimSet(sym_dabsolute_d, Nil)

		_ = tmp120

		tmp121 := PrimSet(symshen_4_dprolog_1memory_d, MakeNumber(1000))

		_ = tmp121

		tmp122 := PrimSet(symshen_4_dloading_2_d, False)

		_ = tmp122

		tmp123 := PrimSet(symshen_4_duserdefs_d, Nil)

		_ = tmp123

		tmp124 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.Return(X)
			return
		}, 1)

		tmp125 := PrimSet(symshen_4_ddemodulation_1function_d, tmp124)

		_ = tmp125

		tmp128 := Call(__e, PrimFunc(symbound_2), sym_dhome_1directory_d)

		tmp129 := PrimNot(tmp128)

		var ifres126 Obj

		if True == tmp129 {
			tmp127 := PrimSet(sym_dhome_1directory_d, MakeString(""))

			ifres126 = tmp127

		} else {
			ifres126 = symshen_4skip

		}

		_ = ifres126

		tmp133 := Call(__e, PrimFunc(symbound_2), sym_dsterror_d)

		tmp134 := PrimNot(tmp133)

		var ifres130 Obj

		if True == tmp134 {
			tmp131 := PrimValue(sym_dstoutput_d)

			tmp132 := PrimSet(sym_dsterror_d, tmp131)

			ifres130 = tmp132

		} else {
			ifres130 = symshen_4skip

		}

		_ = ifres130

		tmp135 := Call(__e, PrimFunc(symprolog_1memory), MakeNumber(10000))

		_ = tmp135

		tmp136 := PrimSet(symshen_4_dloading_2_d, False)

		_ = tmp136

		tmp137 := PrimCons(MakeNumber(2), Nil)

		tmp138 := PrimCons(sym_8s, tmp137)

		tmp139 := PrimCons(MakeNumber(2), tmp138)

		tmp140 := PrimCons(sym_8v, tmp139)

		tmp141 := PrimCons(MakeNumber(2), tmp140)

		tmp142 := PrimCons(sym_8p, tmp141)

		tmp143 := PrimCons(MakeNumber(1), tmp142)

		tmp144 := PrimCons(sym_5_b_6, tmp143)

		tmp145 := PrimCons(MakeNumber(1), tmp144)

		tmp146 := PrimCons(sym_5end_6, tmp145)

		tmp147 := PrimCons(MakeNumber(1), tmp146)

		tmp148 := PrimCons(sym_5e_6, tmp147)

		tmp149 := PrimCons(MakeNumber(2), tmp148)

		tmp150 := PrimCons(sym_a_a, tmp149)

		tmp151 := PrimCons(MakeNumber(2), tmp150)

		tmp152 := PrimCons(sym_1, tmp151)

		tmp153 := PrimCons(MakeNumber(2), tmp152)

		tmp154 := PrimCons(sym_c, tmp153)

		tmp155 := PrimCons(MakeNumber(2), tmp154)

		tmp156 := PrimCons(sym_d, tmp155)

		tmp157 := PrimCons(MakeNumber(2), tmp156)

		tmp158 := PrimCons(sym_7, tmp157)

		tmp159 := PrimCons(MakeNumber(1), tmp158)

		tmp160 := PrimCons(symy_1or_1n_2, tmp159)

		tmp161 := PrimCons(MakeNumber(2), tmp160)

		tmp162 := PrimCons(symwrite_1to_1file, tmp161)

		tmp163 := PrimCons(MakeNumber(2), tmp162)

		tmp164 := PrimCons(symwrite_1byte, tmp163)

		tmp165 := PrimCons(MakeNumber(5), tmp164)

		tmp166 := PrimCons(symwhen, tmp165)

		tmp167 := PrimCons(MakeNumber(0), tmp166)

		tmp168 := PrimCons(symversion, tmp167)

		tmp169 := PrimCons(MakeNumber(5), tmp168)

		tmp170 := PrimCons(symvar_2, tmp169)

		tmp171 := PrimCons(MakeNumber(1), tmp170)

		tmp172 := PrimCons(symvariable_2, tmp171)

		tmp173 := PrimCons(MakeNumber(1), tmp172)

		tmp174 := PrimCons(symvalue, tmp173)

		tmp175 := PrimCons(MakeNumber(3), tmp174)

		tmp176 := PrimCons(symvector_1_6, tmp175)

		tmp177 := PrimCons(MakeNumber(1), tmp176)

		tmp178 := PrimCons(symvector_2, tmp177)

		tmp179 := PrimCons(MakeNumber(1), tmp178)

		tmp180 := PrimCons(symvector, tmp179)

		tmp181 := PrimCons(MakeNumber(0), tmp180)

		tmp182 := PrimCons(symuserdefs, tmp181)

		tmp183 := PrimCons(MakeNumber(2), tmp182)

		tmp184 := PrimCons(symupdate_1lambda_1table, tmp183)

		tmp185 := PrimCons(MakeNumber(1), tmp184)

		tmp186 := PrimCons(symundefmacro, tmp185)

		tmp187 := PrimCons(MakeNumber(1), tmp186)

		tmp188 := PrimCons(symuntrack, tmp187)

		tmp189 := PrimCons(MakeNumber(2), tmp188)

		tmp190 := PrimCons(symunion, tmp189)

		tmp191 := PrimCons(MakeNumber(1), tmp190)

		tmp192 := PrimCons(symunprofile, tmp191)

		tmp193 := PrimCons(MakeNumber(3), tmp192)

		tmp194 := PrimCons(symunput, tmp193)

		tmp195 := PrimCons(MakeNumber(1), tmp194)

		tmp196 := PrimCons(symundefmacro, tmp195)

		tmp197 := PrimCons(MakeNumber(1), tmp196)

		tmp198 := PrimCons(symunabsolute, tmp197)

		tmp199 := PrimCons(MakeNumber(5), tmp198)

		tmp200 := PrimCons(symreturn, tmp199)

		tmp201 := PrimCons(MakeNumber(2), tmp200)

		tmp202 := PrimCons(symtype, tmp201)

		tmp203 := PrimCons(MakeNumber(1), tmp202)

		tmp204 := PrimCons(symtuple_2, tmp203)

		tmp205 := PrimCons(MakeNumber(2), tmp204)

		tmp206 := PrimCons(symtrap_1error, tmp205)

		tmp207 := PrimCons(MakeNumber(0), tmp206)

		tmp208 := PrimCons(symtracked, tmp207)

		tmp209 := PrimCons(MakeNumber(1), tmp208)

		tmp210 := PrimCons(symtrack, tmp209)

		tmp211 := PrimCons(MakeNumber(1), tmp210)

		tmp212 := PrimCons(symtlstr, tmp211)

		tmp213 := PrimCons(MakeNumber(1), tmp212)

		tmp214 := PrimCons(symthaw, tmp213)

		tmp215 := PrimCons(MakeNumber(0), tmp214)

		tmp216 := PrimCons(symtc_2, tmp215)

		tmp217 := PrimCons(MakeNumber(1), tmp216)

		tmp218 := PrimCons(symtc, tmp217)

		tmp219 := PrimCons(MakeNumber(1), tmp218)

		tmp220 := PrimCons(symtl, tmp219)

		tmp221 := PrimCons(MakeNumber(1), tmp220)

		tmp222 := PrimCons(symtail, tmp221)

		tmp223 := PrimCons(MakeNumber(1), tmp222)

		tmp224 := PrimCons(symsystemf, tmp223)

		tmp225 := PrimCons(MakeNumber(1), tmp224)

		tmp226 := PrimCons(symsymbol_2, tmp225)

		tmp227 := PrimCons(MakeNumber(1), tmp226)

		tmp228 := PrimCons(symsum, tmp227)

		tmp229 := PrimCons(MakeNumber(3), tmp228)

		tmp230 := PrimCons(symsubst, tmp229)

		tmp231 := PrimCons(MakeNumber(1), tmp230)

		tmp232 := PrimCons(symstring_2, tmp231)

		tmp233 := PrimCons(MakeNumber(1), tmp232)

		tmp234 := PrimCons(symstring_1_6symbol, tmp233)

		tmp235 := PrimCons(MakeNumber(1), tmp234)

		tmp236 := PrimCons(symstring_1_6n, tmp235)

		tmp237 := PrimCons(MakeNumber(1), tmp236)

		tmp238 := PrimCons(symstr, tmp237)

		tmp239 := PrimCons(MakeNumber(0), tmp238)

		tmp240 := PrimCons(symstoutput, tmp239)

		tmp241 := PrimCons(MakeNumber(0), tmp240)

		tmp242 := PrimCons(symstinput, tmp241)

		tmp243 := PrimCons(MakeNumber(0), tmp242)

		tmp244 := PrimCons(symshen_4step_2, tmp243)

		tmp245 := PrimCons(MakeNumber(1), tmp244)

		tmp246 := PrimCons(symstep, tmp245)

		tmp247 := PrimCons(MakeNumber(0), tmp246)

		tmp248 := PrimCons(symshen_4spy_2, tmp247)

		tmp249 := PrimCons(MakeNumber(1), tmp248)

		tmp250 := PrimCons(symspy, tmp249)

		tmp251 := PrimCons(MakeNumber(2), tmp250)

		tmp252 := PrimCons(symspecialise, tmp251)

		tmp253 := PrimCons(MakeNumber(1), tmp252)

		tmp254 := PrimCons(symsnd, tmp253)

		tmp255 := PrimCons(MakeNumber(1), tmp254)

		tmp256 := PrimCons(symsimple_1error, tmp255)

		tmp257 := PrimCons(MakeNumber(2), tmp256)

		tmp258 := PrimCons(symset, tmp257)

		tmp259 := PrimCons(MakeNumber(1), tmp258)

		tmp260 := PrimCons(symreverse, tmp259)

		tmp261 := PrimCons(MakeNumber(2), tmp260)

		tmp262 := PrimCons(symremove, tmp261)

		tmp263 := PrimCons(MakeNumber(0), tmp262)

		tmp264 := PrimCons(symrelease, tmp263)

		tmp265 := PrimCons(MakeNumber(1), tmp264)

		tmp266 := PrimCons(symreceive, tmp265)

		tmp267 := PrimCons(MakeNumber(1), tmp266)

		tmp268 := PrimCons(symshen_4read_1unit_1string, tmp267)

		tmp269 := PrimCons(MakeNumber(1), tmp268)

		tmp270 := PrimCons(symread_1from_1string_1unprocessed, tmp269)

		tmp271 := PrimCons(MakeNumber(1), tmp270)

		tmp272 := PrimCons(symread_1from_1string, tmp271)

		tmp273 := PrimCons(MakeNumber(1), tmp272)

		tmp274 := PrimCons(symread_1byte, tmp273)

		tmp275 := PrimCons(MakeNumber(1), tmp274)

		tmp276 := PrimCons(symread, tmp275)

		tmp277 := PrimCons(MakeNumber(1), tmp276)

		tmp278 := PrimCons(symread_1file, tmp277)

		tmp279 := PrimCons(MakeNumber(1), tmp278)

		tmp280 := PrimCons(symread_1file_1as_1bytelist, tmp279)

		tmp281 := PrimCons(MakeNumber(1), tmp280)

		tmp282 := PrimCons(symread_1file_1as_1string, tmp281)

		tmp283 := PrimCons(MakeNumber(4), tmp282)

		tmp284 := PrimCons(symput, tmp283)

		tmp285 := PrimCons(MakeNumber(1), tmp284)

		tmp286 := PrimCons(symprotect, tmp285)

		tmp287 := PrimCons(MakeNumber(1), tmp286)

		tmp288 := PrimCons(sympreclude_1all_1but, tmp287)

		tmp289 := PrimCons(MakeNumber(1), tmp288)

		tmp290 := PrimCons(sympreclude, tmp289)

		tmp291 := PrimCons(MakeNumber(1), tmp290)

		tmp292 := PrimCons(symps, tmp291)

		tmp293 := PrimCons(MakeNumber(2), tmp292)

		tmp294 := PrimCons(sympr, tmp293)

		tmp295 := PrimCons(MakeNumber(1), tmp294)

		tmp296 := PrimCons(symprofile_1results, tmp295)

		tmp297 := PrimCons(MakeNumber(1), tmp296)

		tmp298 := PrimCons(symprolog_1memory, tmp297)

		tmp299 := PrimCons(MakeNumber(1), tmp298)

		tmp300 := PrimCons(symshen_4printF, tmp299)

		tmp301 := PrimCons(MakeNumber(1), tmp300)

		tmp302 := PrimCons(symshen_4print_1freshterm, tmp301)

		tmp303 := PrimCons(MakeNumber(1), tmp302)

		tmp304 := PrimCons(symshen_4print_1prolog_1vector, tmp303)

		tmp305 := PrimCons(MakeNumber(1), tmp304)

		tmp306 := PrimCons(symprofile, tmp305)

		tmp307 := PrimCons(MakeNumber(1), tmp306)

		tmp308 := PrimCons(symprint, tmp307)

		tmp309 := PrimCons(MakeNumber(1), tmp308)

		tmp310 := PrimCons(sympreclude_1all_1but, tmp309)

		tmp311 := PrimCons(MakeNumber(2), tmp310)

		tmp312 := PrimCons(sympos, tmp311)

		tmp313 := PrimCons(MakeNumber(0), tmp312)

		tmp314 := PrimCons(symporters, tmp313)

		tmp315 := PrimCons(MakeNumber(0), tmp314)

		tmp316 := PrimCons(symport, tmp315)

		tmp317 := PrimCons(MakeNumber(1), tmp316)

		tmp318 := PrimCons(sympackage_2, tmp317)

		tmp319 := PrimCons(MakeNumber(3), tmp318)

		tmp320 := PrimCons(sympackage, tmp319)

		tmp321 := PrimCons(MakeNumber(0), tmp320)

		tmp322 := PrimCons(symos, tmp321)

		tmp323 := PrimCons(MakeNumber(2), tmp322)

		tmp324 := PrimCons(symor, tmp323)

		tmp325 := PrimCons(MakeNumber(0), tmp324)

		tmp326 := PrimCons(symoptimise_2, tmp325)

		tmp327 := PrimCons(MakeNumber(1), tmp326)

		tmp328 := PrimCons(symoptimise, tmp327)

		tmp329 := PrimCons(MakeNumber(2), tmp328)

		tmp330 := PrimCons(symopen, tmp329)

		tmp331 := PrimCons(MakeNumber(1), tmp330)

		tmp332 := PrimCons(symoccurs_1check, tmp331)

		tmp333 := PrimCons(MakeNumber(0), tmp332)

		tmp334 := PrimCons(symoccurs_2, tmp333)

		tmp335 := PrimCons(MakeNumber(2), tmp334)

		tmp336 := PrimCons(symoccurrences, tmp335)

		tmp337 := PrimCons(MakeNumber(1), tmp336)

		tmp338 := PrimCons(symoccurs_1check, tmp337)

		tmp339 := PrimCons(MakeNumber(1), tmp338)

		tmp340 := PrimCons(symnumber_2, tmp339)

		tmp341 := PrimCons(MakeNumber(1), tmp340)

		tmp342 := PrimCons(symn_1_6string, tmp341)

		tmp343 := PrimCons(MakeNumber(2), tmp342)

		tmp344 := PrimCons(symnth, tmp343)

		tmp345 := PrimCons(MakeNumber(1), tmp344)

		tmp346 := PrimCons(symnot, tmp345)

		tmp347 := PrimCons(MakeNumber(1), tmp346)

		tmp348 := PrimCons(symnl, tmp347)

		tmp349 := PrimCons(MakeNumber(1), tmp348)

		tmp350 := PrimCons(symmaxinferences, tmp349)

		tmp351 := PrimCons(MakeNumber(2), tmp350)

		tmp352 := PrimCons(symmapcan, tmp351)

		tmp353 := PrimCons(MakeNumber(2), tmp352)

		tmp354 := PrimCons(symmap, tmp353)

		tmp355 := PrimCons(MakeNumber(1), tmp354)

		tmp356 := PrimCons(symmacroexpand, tmp355)

		tmp357 := PrimCons(MakeNumber(1), tmp356)

		tmp358 := PrimCons(symvector, tmp357)

		tmp359 := PrimCons(MakeNumber(2), tmp358)

		tmp360 := PrimCons(sym_5_a, tmp359)

		tmp361 := PrimCons(MakeNumber(2), tmp360)

		tmp362 := PrimCons(sym_5, tmp361)

		tmp363 := PrimCons(MakeNumber(1), tmp362)

		tmp364 := PrimCons(symload, tmp363)

		tmp365 := PrimCons(MakeNumber(1), tmp364)

		tmp366 := PrimCons(symlist, tmp365)

		tmp367 := PrimCons(MakeNumber(1), tmp366)

		tmp368 := PrimCons(symlineread, tmp367)

		tmp369 := PrimCons(MakeNumber(1), tmp368)

		tmp370 := PrimCons(symlimit, tmp369)

		tmp371 := PrimCons(MakeNumber(1), tmp370)

		tmp372 := PrimCons(symlength, tmp371)

		tmp373 := PrimCons(MakeNumber(0), tmp372)

		tmp374 := PrimCons(symlanguage, tmp373)

		tmp375 := PrimCons(MakeNumber(6), tmp374)

		tmp376 := PrimCons(symis_b, tmp375)

		tmp377 := PrimCons(MakeNumber(6), tmp376)

		tmp378 := PrimCons(symis, tmp377)

		tmp379 := PrimCons(MakeNumber(0), tmp378)

		tmp380 := PrimCons(symit, tmp379)

		tmp381 := PrimCons(MakeNumber(1), tmp380)

		tmp382 := PrimCons(syminternal, tmp381)

		tmp383 := PrimCons(MakeNumber(2), tmp382)

		tmp384 := PrimCons(symintersection, tmp383)

		tmp385 := PrimCons(MakeNumber(1), tmp384)

		tmp386 := PrimCons(syminclude_1all_1but, tmp385)

		tmp387 := PrimCons(MakeNumber(0), tmp386)

		tmp388 := PrimCons(symimplementation, tmp387)

		tmp389 := PrimCons(MakeNumber(2), tmp388)

		tmp390 := PrimCons(syminput_7, tmp389)

		tmp391 := PrimCons(MakeNumber(1), tmp390)

		tmp392 := PrimCons(syminput, tmp391)

		tmp393 := PrimCons(MakeNumber(0), tmp392)

		tmp394 := PrimCons(syminferences, tmp393)

		tmp395 := PrimCons(MakeNumber(1), tmp394)

		tmp396 := PrimCons(symintern, tmp395)

		tmp397 := PrimCons(MakeNumber(1), tmp396)

		tmp398 := PrimCons(syminternal, tmp397)

		tmp399 := PrimCons(MakeNumber(1), tmp398)

		tmp400 := PrimCons(syminteger_2, tmp399)

		tmp401 := PrimCons(MakeNumber(1), tmp400)

		tmp402 := PrimCons(symin_1package, tmp401)

		tmp403 := PrimCons(MakeNumber(0), tmp402)

		tmp404 := PrimCons(symshen_4included, tmp403)

		tmp405 := PrimCons(MakeNumber(1), tmp404)

		tmp406 := PrimCons(syminclude, tmp405)

		tmp407 := PrimCons(MakeNumber(3), tmp406)

		tmp408 := PrimCons(symif, tmp407)

		tmp409 := PrimCons(MakeNumber(1), tmp408)

		tmp410 := PrimCons(symhush, tmp409)

		tmp411 := PrimCons(MakeNumber(0), tmp410)

		tmp412 := PrimCons(symhush_2, tmp411)

		tmp413 := PrimCons(MakeNumber(1), tmp412)

		tmp414 := PrimCons(symhead, tmp413)

		tmp415 := PrimCons(MakeNumber(1), tmp414)

		tmp416 := PrimCons(symhdstr, tmp415)

		tmp417 := PrimCons(MakeNumber(1), tmp416)

		tmp418 := PrimCons(symhdv, tmp417)

		tmp419 := PrimCons(MakeNumber(1), tmp418)

		tmp420 := PrimCons(symhd, tmp419)

		tmp421 := PrimCons(MakeNumber(2), tmp420)

		tmp422 := PrimCons(symhash, tmp421)

		tmp423 := PrimCons(MakeNumber(2), tmp422)

		tmp424 := PrimCons(sym_a, tmp423)

		tmp425 := PrimCons(MakeNumber(2), tmp424)

		tmp426 := PrimCons(sym_6_a, tmp425)

		tmp427 := PrimCons(MakeNumber(2), tmp426)

		tmp428 := PrimCons(sym_6, tmp427)

		tmp429 := PrimCons(MakeNumber(2), tmp428)

		tmp430 := PrimCons(sym_5_1vector, tmp429)

		tmp431 := PrimCons(MakeNumber(2), tmp430)

		tmp432 := PrimCons(sym_5_1address, tmp431)

		tmp433 := PrimCons(MakeNumber(3), tmp432)

		tmp434 := PrimCons(symaddress_1_6, tmp433)

		tmp435 := PrimCons(MakeNumber(1), tmp434)

		tmp436 := PrimCons(symget_1time, tmp435)

		tmp437 := PrimCons(MakeNumber(3), tmp436)

		tmp438 := PrimCons(symget, tmp437)

		tmp439 := PrimCons(MakeNumber(1), tmp438)

		tmp440 := PrimCons(symgensym, tmp439)

		tmp441 := PrimCons(MakeNumber(1), tmp440)

		tmp442 := PrimCons(symfunction, tmp441)

		tmp443 := PrimCons(MakeNumber(1), tmp442)

		tmp444 := PrimCons(symfn, tmp443)

		tmp445 := PrimCons(MakeNumber(1), tmp444)

		tmp446 := PrimCons(symfst, tmp445)

		tmp447 := PrimCons(MakeNumber(0), tmp446)

		tmp448 := PrimCons(symfresh, tmp447)

		tmp449 := PrimCons(MakeNumber(1), tmp448)

		tmp450 := PrimCons(symfreeze, tmp449)

		tmp451 := PrimCons(MakeNumber(5), tmp450)

		tmp452 := PrimCons(symfork, tmp451)

		tmp453 := PrimCons(MakeNumber(1), tmp452)

		tmp454 := PrimCons(symforeign, tmp453)

		tmp455 := PrimCons(MakeNumber(7), tmp454)

		tmp456 := PrimCons(symfindall, tmp455)

		tmp457 := PrimCons(MakeNumber(2), tmp456)

		tmp458 := PrimCons(symfix, tmp457)

		tmp459 := PrimCons(MakeNumber(0), tmp458)

		tmp460 := PrimCons(symfail, tmp459)

		tmp461 := PrimCons(MakeNumber(2), tmp460)

		tmp462 := PrimCons(symfail_1if, tmp461)

		tmp463 := PrimCons(MakeNumber(0), tmp462)

		tmp464 := PrimCons(symfactorise_2, tmp463)

		tmp465 := PrimCons(MakeNumber(1), tmp464)

		tmp466 := PrimCons(symfactorise, tmp465)

		tmp467 := PrimCons(MakeNumber(1), tmp466)

		tmp468 := PrimCons(symexternal, tmp467)

		tmp469 := PrimCons(MakeNumber(1), tmp468)

		tmp470 := PrimCons(symexplode, tmp469)

		tmp471 := PrimCons(MakeNumber(1), tmp470)

		tmp472 := PrimCons(symeval_1kl, tmp471)

		tmp473 := PrimCons(MakeNumber(1), tmp472)

		tmp474 := PrimCons(symeval, tmp473)

		tmp475 := PrimCons(MakeNumber(1), tmp474)

		tmp476 := PrimCons(symerror_1to_1string, tmp475)

		tmp477 := PrimCons(MakeNumber(1), tmp476)

		tmp478 := PrimCons(symexternal, tmp477)

		tmp479 := PrimCons(MakeNumber(1), tmp478)

		tmp480 := PrimCons(symenable_1type_1theory, tmp479)

		tmp481 := PrimCons(MakeNumber(1), tmp480)

		tmp482 := PrimCons(symempty_2, tmp481)

		tmp483 := PrimCons(MakeNumber(2), tmp482)

		tmp484 := PrimCons(symelement_2, tmp483)

		tmp485 := PrimCons(MakeNumber(2), tmp484)

		tmp486 := PrimCons(symdo, tmp485)

		tmp487 := PrimCons(MakeNumber(2), tmp486)

		tmp488 := PrimCons(symdifference, tmp487)

		tmp489 := PrimCons(MakeNumber(1), tmp488)

		tmp490 := PrimCons(symdestroy, tmp489)

		tmp491 := PrimCons(MakeNumber(2), tmp490)

		tmp492 := PrimCons(symdeclare, tmp491)

		tmp493 := PrimCons(MakeNumber(0), tmp492)

		tmp494 := PrimCons(symdatatypes, tmp493)

		tmp495 := PrimCons(MakeNumber(1), tmp494)

		tmp496 := PrimCons(symclose, tmp495)

		tmp497 := PrimCons(MakeNumber(2), tmp496)

		tmp498 := PrimCons(symcn, tmp497)

		tmp499 := PrimCons(MakeNumber(1), tmp498)

		tmp500 := PrimCons(symcons_2, tmp499)

		tmp501 := PrimCons(MakeNumber(2), tmp500)

		tmp502 := PrimCons(symcons, tmp501)

		tmp503 := PrimCons(MakeNumber(2), tmp502)

		tmp504 := PrimCons(symconcat, tmp503)

		tmp505 := PrimCons(MakeNumber(2), tmp504)

		tmp506 := PrimCons(symcompile, tmp505)

		tmp507 := PrimCons(MakeNumber(1), tmp506)

		tmp508 := PrimCons(symcd, tmp507)

		tmp509 := PrimCons(MakeNumber(5), tmp508)

		tmp510 := PrimCons(symcall, tmp509)

		tmp511 := PrimCons(MakeNumber(6), tmp510)

		tmp512 := PrimCons(symbind, tmp511)

		tmp513 := PrimCons(MakeNumber(1), tmp512)

		tmp514 := PrimCons(symbound_2, tmp513)

		tmp515 := PrimCons(MakeNumber(1), tmp514)

		tmp516 := PrimCons(symbootstrap, tmp515)

		tmp517 := PrimCons(MakeNumber(1), tmp516)

		tmp518 := PrimCons(symboolean_2, tmp517)

		tmp519 := PrimCons(MakeNumber(1), tmp518)

		tmp520 := PrimCons(symatom_2, tmp519)

		tmp521 := PrimCons(MakeNumber(2), tmp520)

		tmp522 := PrimCons(symassoc, tmp521)

		tmp523 := PrimCons(MakeNumber(1), tmp522)

		tmp524 := PrimCons(symarity, tmp523)

		tmp525 := PrimCons(MakeNumber(2), tmp524)

		tmp526 := PrimCons(symappend, tmp525)

		tmp527 := PrimCons(MakeNumber(2), tmp526)

		tmp528 := PrimCons(symand, tmp527)

		tmp529 := PrimCons(MakeNumber(2), tmp528)

		tmp530 := PrimCons(symadjoin, tmp529)

		tmp531 := PrimCons(MakeNumber(3), tmp530)

		tmp532 := PrimCons(symaddress_1_6, tmp531)

		tmp533 := PrimCons(MakeNumber(1), tmp532)

		tmp534 := PrimCons(symabsvector, tmp533)

		tmp535 := PrimCons(MakeNumber(1), tmp534)

		tmp536 := PrimCons(symabsvector_2, tmp535)

		tmp537 := PrimCons(MakeNumber(1), tmp536)

		tmp538 := PrimCons(symabsolute, tmp537)

		tmp539 := PrimCons(MakeNumber(0), tmp538)

		tmp540 := PrimCons(symabort, tmp539)

		tmp541 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp540)

		_ = tmp541

		tmp542 := PrimIntern(MakeString(":"))

		tmp543 := PrimIntern(MakeString(";"))

		tmp544 := PrimIntern(MakeString(":="))

		tmp545 := PrimIntern(MakeString(","))

		tmp546 := Call(__e, PrimFunc(symvector), MakeNumber(0))

		tmp547 := PrimIntern(MakeString("bar!"))

		tmp548 := PrimCons(symabort, Nil)

		tmp549 := PrimCons(symabsolute, tmp548)

		tmp550 := PrimCons(symabsvector, tmp549)

		tmp551 := PrimCons(symabsvector_2, tmp550)

		tmp552 := PrimCons(symaddress_1_6, tmp551)

		tmp553 := PrimCons(sym_5_1address, tmp552)

		tmp554 := PrimCons(symadjoin, tmp553)

		tmp555 := PrimCons(symand, tmp554)

		tmp556 := PrimCons(symappend, tmp555)

		tmp557 := PrimCons(symarity, tmp556)

		tmp558 := PrimCons(symassoc, tmp557)

		tmp559 := PrimCons(symassertz, tmp558)

		tmp560 := PrimCons(symasserta, tmp559)

		tmp561 := PrimCons(symatom_2, tmp560)

		tmp562 := PrimCons(tmp547, tmp561)

		tmp563 := PrimCons(symbootstrap, tmp562)

		tmp564 := PrimCons(symboolean, tmp563)

		tmp565 := PrimCons(symboolean_2, tmp564)

		tmp566 := PrimCons(symbound_2, tmp565)

		tmp567 := PrimCons(symbind, tmp566)

		tmp568 := PrimCons(symclose, tmp567)

		tmp569 := PrimCons(symcall, tmp568)

		tmp570 := PrimCons(symcases, tmp569)

		tmp571 := PrimCons(symcd, tmp570)

		tmp572 := PrimCons(symcompile, tmp571)

		tmp573 := PrimCons(symconcat, tmp572)

		tmp574 := PrimCons(symcond, tmp573)

		tmp575 := PrimCons(symcons, tmp574)

		tmp576 := PrimCons(symcons_2, tmp575)

		tmp577 := PrimCons(symcn, tmp576)

		tmp578 := PrimCons(symshen_4ctxt, tmp577)

		tmp579 := PrimCons(symdatatypes, tmp578)

		tmp580 := PrimCons(symdatatype, tmp579)

		tmp581 := PrimCons(symdeclare, tmp580)

		tmp582 := PrimCons(symdefprolog, tmp581)

		tmp583 := PrimCons(symdefcc, tmp582)

		tmp584 := PrimCons(symdefmacro, tmp583)

		tmp585 := PrimCons(symdefine, tmp584)

		tmp586 := PrimCons(symdefun, tmp585)

		tmp587 := PrimCons(symdestroy, tmp586)

		tmp588 := PrimCons(symdifference, tmp587)

		tmp589 := PrimCons(symdo, tmp588)

		tmp590 := PrimCons(symelement_2, tmp589)

		tmp591 := PrimCons(symempty_2, tmp590)

		tmp592 := PrimCons(symerror, tmp591)

		tmp593 := PrimCons(symerror_1to_1string, tmp592)

		tmp594 := PrimCons(symeval, tmp593)

		tmp595 := PrimCons(symeval_1kl, tmp594)

		tmp596 := PrimCons(symexception, tmp595)

		tmp597 := PrimCons(symexternal, tmp596)

		tmp598 := PrimCons(symexplode, tmp597)

		tmp599 := PrimCons(symenable_1type_1theory, tmp598)

		tmp600 := PrimCons(False, tmp599)

		tmp601 := PrimCons(symfindall, tmp600)

		tmp602 := PrimCons(symfactorise, tmp601)

		tmp603 := PrimCons(symfail_1if, tmp602)

		tmp604 := PrimCons(symfail, tmp603)

		tmp605 := PrimCons(symfile, tmp604)

		tmp606 := PrimCons(symfix, tmp605)

		tmp607 := PrimCons(symforeign, tmp606)

		tmp608 := PrimCons(symfork, tmp607)

		tmp609 := PrimCons(symfresh, tmp608)

		tmp610 := PrimCons(symfreeze, tmp609)

		tmp611 := PrimCons(symfst, tmp610)

		tmp612 := PrimCons(symfunction, tmp611)

		tmp613 := PrimCons(symfn, tmp612)

		tmp614 := PrimCons(symgensym, tmp613)

		tmp615 := PrimCons(symget_1time, tmp614)

		tmp616 := PrimCons(symget, tmp615)

		tmp617 := PrimCons(symhash, tmp616)

		tmp618 := PrimCons(symhdstr, tmp617)

		tmp619 := PrimCons(symhdv, tmp618)

		tmp620 := PrimCons(symhd, tmp619)

		tmp621 := PrimCons(symhead, tmp620)

		tmp622 := PrimCons(symif, tmp621)

		tmp623 := PrimCons(symimplementation, tmp622)

		tmp624 := PrimCons(syminternal, tmp623)

		tmp625 := PrimCons(symin_1package, tmp624)

		tmp626 := PrimCons(symin, tmp625)

		tmp627 := PrimCons(symis_b, tmp626)

		tmp628 := PrimCons(symis, tmp627)

		tmp629 := PrimCons(symit, tmp628)

		tmp630 := PrimCons(syminclude_1all_1but, tmp629)

		tmp631 := PrimCons(syminclude, tmp630)

		tmp632 := PrimCons(syminline, tmp631)

		tmp633 := PrimCons(syminput_7, tmp632)

		tmp634 := PrimCons(syminput, tmp633)

		tmp635 := PrimCons(syminteger_2, tmp634)

		tmp636 := PrimCons(symintern, tmp635)

		tmp637 := PrimCons(syminferences, tmp636)

		tmp638 := PrimCons(symintersection, tmp637)

		tmp639 := PrimCons(symis, tmp638)

		tmp640 := PrimCons(symlanguage, tmp639)

		tmp641 := PrimCons(symlambda, tmp640)

		tmp642 := PrimCons(symlazy, tmp641)

		tmp643 := PrimCons(symlet, tmp642)

		tmp644 := PrimCons(symlength, tmp643)

		tmp645 := PrimCons(symlimit, tmp644)

		tmp646 := PrimCons(symlineread, tmp645)

		tmp647 := PrimCons(symlist, tmp646)

		tmp648 := PrimCons(symloaded, tmp647)

		tmp649 := PrimCons(symload, tmp648)

		tmp650 := PrimCons(symmake_1string, tmp649)

		tmp651 := PrimCons(symmap, tmp650)

		tmp652 := PrimCons(symmapcan, tmp651)

		tmp653 := PrimCons(symmaxinferences, tmp652)

		tmp654 := PrimCons(symmacroexpand, tmp653)

		tmp655 := PrimCons(symmode, tmp654)

		tmp656 := PrimCons(symnl, tmp655)

		tmp657 := PrimCons(symnot, tmp656)

		tmp658 := PrimCons(symnth, tmp657)

		tmp659 := PrimCons(symnull, tmp658)

		tmp660 := PrimCons(symnumber, tmp659)

		tmp661 := PrimCons(symnumber_2, tmp660)

		tmp662 := PrimCons(symn_1_6string, tmp661)

		tmp663 := PrimCons(symoccurs_1check, tmp662)

		tmp664 := PrimCons(symoccurrences, tmp663)

		tmp665 := PrimCons(symopen, tmp664)

		tmp666 := PrimCons(symoptimise, tmp665)

		tmp667 := PrimCons(symor, tmp666)

		tmp668 := PrimCons(symos, tmp667)

		tmp669 := PrimCons(symout, tmp668)

		tmp670 := PrimCons(symoutput, tmp669)

		tmp671 := PrimCons(sympackage, tmp670)

		tmp672 := PrimCons(symport, tmp671)

		tmp673 := PrimCons(symporters, tmp672)

		tmp674 := PrimCons(sympos, tmp673)

		tmp675 := PrimCons(sympr, tmp674)

		tmp676 := PrimCons(symprint, tmp675)

		tmp677 := PrimCons(symprolog_1memory, tmp676)

		tmp678 := PrimCons(symprofile, tmp677)

		tmp679 := PrimCons(symprofile_1results, tmp678)

		tmp680 := PrimCons(symprotect, tmp679)

		tmp681 := PrimCons(symprolog_2, tmp680)

		tmp682 := PrimCons(symps, tmp681)

		tmp683 := PrimCons(sympreclude_1all_1but, tmp682)

		tmp684 := PrimCons(sympreclude, tmp683)

		tmp685 := PrimCons(symput, tmp684)

		tmp686 := PrimCons(sympackage_2, tmp685)

		tmp687 := PrimCons(symread_1from_1string_1unprocessed, tmp686)

		tmp688 := PrimCons(symread_1from_1string, tmp687)

		tmp689 := PrimCons(symread_1byte, tmp688)

		tmp690 := PrimCons(symread_1file_1as_1string, tmp689)

		tmp691 := PrimCons(symread_1file_1as_1bytelist, tmp690)

		tmp692 := PrimCons(symread_1file, tmp691)

		tmp693 := PrimCons(symreceive, tmp692)

		tmp694 := PrimCons(symread, tmp693)

		tmp695 := PrimCons(symrelease, tmp694)

		tmp696 := PrimCons(symremove, tmp695)

		tmp697 := PrimCons(symretract, tmp696)

		tmp698 := PrimCons(symreverse, tmp697)

		tmp699 := PrimCons(symrun, tmp698)

		tmp700 := PrimCons(symstr, tmp699)

		tmp701 := PrimCons(symsave, tmp700)

		tmp702 := PrimCons(symset, tmp701)

		tmp703 := PrimCons(symsimple_1error, tmp702)

		tmp704 := PrimCons(symsnd, tmp703)

		tmp705 := PrimCons(symspecialise, tmp704)

		tmp706 := PrimCons(symspy, tmp705)

		tmp707 := PrimCons(symstep, tmp706)

		tmp708 := PrimCons(symstoutput, tmp707)

		tmp709 := PrimCons(symsterror, tmp708)

		tmp710 := PrimCons(symstinput, tmp709)

		tmp711 := PrimCons(symstring, tmp710)

		tmp712 := PrimCons(symstream, tmp711)

		tmp713 := PrimCons(symstring_1_6n, tmp712)

		tmp714 := PrimCons(symstring_2, tmp713)

		tmp715 := PrimCons(symsubst, tmp714)

		tmp716 := PrimCons(symsum, tmp715)

		tmp717 := PrimCons(symstring_1_6symbol, tmp716)

		tmp718 := PrimCons(symsymbol_2, tmp717)

		tmp719 := PrimCons(symsymbol, tmp718)

		tmp720 := PrimCons(symsynonyms, tmp719)

		tmp721 := PrimCons(symsystemf, tmp720)

		tmp722 := PrimCons(symtail, tmp721)

		tmp723 := PrimCons(symtlv, tmp722)

		tmp724 := PrimCons(symtlstr, tmp723)

		tmp725 := PrimCons(symtl, tmp724)

		tmp726 := PrimCons(symtc, tmp725)

		tmp727 := PrimCons(symtc_2, tmp726)

		tmp728 := PrimCons(symthaw, tmp727)

		tmp729 := PrimCons(symtime, tmp728)

		tmp730 := PrimCons(symtrack, tmp729)

		tmp731 := PrimCons(symtrap_1error, tmp730)

		tmp732 := PrimCons(True, tmp731)

		tmp733 := PrimCons(symtuple_2, tmp732)

		tmp734 := PrimCons(symtype, tmp733)

		tmp735 := PrimCons(symreturn, tmp734)

		tmp736 := PrimCons(symunabsolute, tmp735)

		tmp737 := PrimCons(symundefmacro, tmp736)

		tmp738 := PrimCons(symunprofile, tmp737)

		tmp739 := PrimCons(symunput, tmp738)

		tmp740 := PrimCons(symunion, tmp739)

		tmp741 := PrimCons(symshen_4unix, tmp740)

		tmp742 := PrimCons(symunit, tmp741)

		tmp743 := PrimCons(symuntrack, tmp742)

		tmp744 := PrimCons(symunspecialise, tmp743)

		tmp745 := PrimCons(symupdate_1lambda_1table, tmp744)

		tmp746 := PrimCons(symu_b, tmp745)

		tmp747 := PrimCons(symvector_2, tmp746)

		tmp748 := PrimCons(symvector, tmp747)

		tmp749 := PrimCons(sym_5_1vector, tmp748)

		tmp750 := PrimCons(symvector_1_6, tmp749)

		tmp751 := PrimCons(symvalue, tmp750)

		tmp752 := PrimCons(symvar_2, tmp751)

		tmp753 := PrimCons(symvariable_2, tmp752)

		tmp754 := PrimCons(symverified, tmp753)

		tmp755 := PrimCons(symversion, tmp754)

		tmp756 := PrimCons(symwarn, tmp755)

		tmp757 := PrimCons(symwhen, tmp756)

		tmp758 := PrimCons(symwhere, tmp757)

		tmp759 := PrimCons(symwrite_1byte, tmp758)

		tmp760 := PrimCons(symwrite_1to_1file, tmp759)

		tmp761 := PrimCons(symy_1or_1n_2, tmp760)

		tmp762 := PrimCons(tmp546, tmp761)

		tmp763 := PrimCons(sym_6_6, tmp762)

		tmp764 := PrimCons(sym_5, tmp763)

		tmp765 := PrimCons(sym_5_a, tmp764)

		tmp766 := PrimCons(sym_7, tmp765)

		tmp767 := PrimCons(sym_d, tmp766)

		tmp768 := PrimCons(sym_c, tmp767)

		tmp769 := PrimCons(sym_1, tmp768)

		tmp770 := PrimCons(sym_3, tmp769)

		tmp771 := PrimCons(sym_5end_6, tmp770)

		tmp772 := PrimCons(sym_5_b_6, tmp771)

		tmp773 := PrimCons(sym_c_4, tmp772)

		tmp774 := PrimCons(sym_a_a_6, tmp773)

		tmp775 := PrimCons(sym_6, tmp774)

		tmp776 := PrimCons(sym_6_a, tmp775)

		tmp777 := PrimCons(sym_a, tmp776)

		tmp778 := PrimCons(sym_a_a, tmp777)

		tmp779 := PrimCons(sym_5e_6, tmp778)

		tmp780 := PrimCons(sym_1_6, tmp779)

		tmp781 := PrimCons(sym_5_1, tmp780)

		tmp782 := PrimCons(sym_dhush_d, tmp781)

		tmp783 := PrimCons(sym_dporters_d, tmp782)

		tmp784 := PrimCons(sym_dport_d, tmp783)

		tmp785 := PrimCons(sym_8s, tmp784)

		tmp786 := PrimCons(sym_8p, tmp785)

		tmp787 := PrimCons(sym_8v, tmp786)

		tmp788 := PrimCons(sym_dproperty_1vector_d, tmp787)

		tmp789 := PrimCons(sym_drelease_d, tmp788)

		tmp790 := PrimCons(sym_dos_d, tmp789)

		tmp791 := PrimCons(sym_dmacros_d, tmp790)

		tmp792 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp791)

		tmp793 := PrimCons(sym_dversion_d, tmp792)

		tmp794 := PrimCons(sym_dhome_1directory_d, tmp793)

		tmp795 := PrimCons(sym_dstoutput_d, tmp794)

		tmp796 := PrimCons(sym_dsterror_d, tmp795)

		tmp797 := PrimCons(sym_dstinput_d, tmp796)

		tmp798 := PrimCons(sym_dimplementation_d, tmp797)

		tmp799 := PrimCons(sym_dlanguage_d, tmp798)

		tmp800 := PrimCons(sym__, tmp799)

		tmp801 := PrimCons(tmp545, tmp800)

		tmp802 := PrimCons(tmp544, tmp801)

		tmp803 := PrimCons(tmp543, tmp802)

		tmp804 := PrimCons(tmp542, tmp803)

		tmp805 := PrimCons(sym_e_e, tmp804)

		tmp806 := PrimCons(sym_5_1_1, tmp805)

		tmp807 := PrimCons(sym_1_1_6, tmp806)

		tmp808 := PrimCons(sym_i, tmp807)

		tmp809 := PrimCons(sym_j, tmp808)

		tmp810 := PrimCons(sym_b, tmp809)

		tmp811 := PrimValue(sym_dproperty_1vector_d)

		tmp812 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp810, tmp811)

		_ = tmp812

		tmp813 := PrimAbsvector(MakeNumber(0))

		__e.Return(PrimSet(symshen_4_dempty_1absvector_d, tmp813))
		return

	}, 0)

	tmp814 := Call(__e, ns2_1set, symshen_4initialise_1environment, tmp75)

	_ = tmp814

	tmp815 := MakeNative(func(__e *ControlFlow) {
		tmp816 := PrimSet(symshen_4_dsigf_d, Nil)

		_ = tmp816

		tmp817 := MakeNative(func(__e *ControlFlow) {
			V5927 := __e.Get(1)
			_ = V5927
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5923 := __e.Get(1)
				_ = B5923
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5924 := __e.Get(1)
					_ = L5924
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5925 := __e.Get(1)
						_ = Key5925
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5926 := __e.Get(1)
							_ = C5926
							tmp818 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp819 := PrimCons(A, Nil)

								tmp820 := PrimCons(sym_1_1_6, tmp819)

								tmp821 := Call(__e, PrimFunc(symis_b), V5927, tmp820, B5923, L5924, Key5925, C5926)

								__e.TailApply(PrimFunc(symshen_4gc), B5923, tmp821)
								return

							}, 1)

							tmp822 := Call(__e, PrimFunc(symshen_4newpv), B5923)

							__e.TailApply(tmp818, tmp822)
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

		tmp823 := PrimValue(symshen_4_dsigf_d)

		tmp824 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabort, tmp817, tmp823)

		tmp825 := PrimSet(symshen_4_dsigf_d, tmp824)

		_ = tmp825

		tmp826 := MakeNative(func(__e *ControlFlow) {
			V5932 := __e.Get(1)
			_ = V5932
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5928 := __e.Get(1)
				_ = B5928
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5929 := __e.Get(1)
					_ = L5929
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5930 := __e.Get(1)
						_ = Key5930
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5931 := __e.Get(1)
							_ = C5931
							tmp827 := PrimCons(symstring, Nil)

							tmp828 := PrimCons(symlist, tmp827)

							tmp829 := PrimCons(tmp828, Nil)

							tmp830 := PrimCons(sym_1_1_6, tmp829)

							tmp831 := PrimCons(symstring, tmp830)

							__e.TailApply(PrimFunc(symis_b), V5932, tmp831, B5928, L5929, Key5930, C5931)
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

		tmp832 := PrimValue(symshen_4_dsigf_d)

		tmp833 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsolute, tmp826, tmp832)

		tmp834 := PrimSet(symshen_4_dsigf_d, tmp833)

		_ = tmp834

		tmp835 := MakeNative(func(__e *ControlFlow) {
			V5937 := __e.Get(1)
			_ = V5937
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5933 := __e.Get(1)
				_ = B5933
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5934 := __e.Get(1)
					_ = L5934
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5935 := __e.Get(1)
						_ = Key5935
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5936 := __e.Get(1)
							_ = C5936
							tmp836 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp837 := PrimCons(symboolean, Nil)

								tmp838 := PrimCons(sym_1_1_6, tmp837)

								tmp839 := PrimCons(A, tmp838)

								tmp840 := Call(__e, PrimFunc(symis_b), V5937, tmp839, B5933, L5934, Key5935, C5936)

								__e.TailApply(PrimFunc(symshen_4gc), B5933, tmp840)
								return

							}, 1)

							tmp841 := Call(__e, PrimFunc(symshen_4newpv), B5933)

							__e.TailApply(tmp836, tmp841)
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

		tmp842 := PrimValue(symshen_4_dsigf_d)

		tmp843 := Call(__e, PrimFunc(symshen_4assoc_1_6), symabsvector_2, tmp835, tmp842)

		tmp844 := PrimSet(symshen_4_dsigf_d, tmp843)

		_ = tmp844

		tmp845 := MakeNative(func(__e *ControlFlow) {
			V5942 := __e.Get(1)
			_ = V5942
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5938 := __e.Get(1)
				_ = B5938
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5939 := __e.Get(1)
					_ = L5939
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5940 := __e.Get(1)
						_ = Key5940
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5941 := __e.Get(1)
							_ = C5941
							tmp846 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp847 := PrimCons(A, Nil)

								tmp848 := PrimCons(symlist, tmp847)

								tmp849 := PrimCons(A, Nil)

								tmp850 := PrimCons(symlist, tmp849)

								tmp851 := PrimCons(tmp850, Nil)

								tmp852 := PrimCons(sym_1_1_6, tmp851)

								tmp853 := PrimCons(tmp848, tmp852)

								tmp854 := PrimCons(tmp853, Nil)

								tmp855 := PrimCons(sym_1_1_6, tmp854)

								tmp856 := PrimCons(A, tmp855)

								tmp857 := Call(__e, PrimFunc(symis_b), V5942, tmp856, B5938, L5939, Key5940, C5941)

								__e.TailApply(PrimFunc(symshen_4gc), B5938, tmp857)
								return

							}, 1)

							tmp858 := Call(__e, PrimFunc(symshen_4newpv), B5938)

							__e.TailApply(tmp846, tmp858)
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

		tmp859 := PrimValue(symshen_4_dsigf_d)

		tmp860 := Call(__e, PrimFunc(symshen_4assoc_1_6), symadjoin, tmp845, tmp859)

		tmp861 := PrimSet(symshen_4_dsigf_d, tmp860)

		_ = tmp861

		tmp862 := MakeNative(func(__e *ControlFlow) {
			V5947 := __e.Get(1)
			_ = V5947
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5943 := __e.Get(1)
				_ = B5943
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5944 := __e.Get(1)
					_ = L5944
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5945 := __e.Get(1)
						_ = Key5945
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5946 := __e.Get(1)
							_ = C5946
							tmp863 := PrimCons(symboolean, Nil)

							tmp864 := PrimCons(sym_1_1_6, tmp863)

							tmp865 := PrimCons(symboolean, tmp864)

							tmp866 := PrimCons(tmp865, Nil)

							tmp867 := PrimCons(sym_1_1_6, tmp866)

							tmp868 := PrimCons(symboolean, tmp867)

							__e.TailApply(PrimFunc(symis_b), V5947, tmp868, B5943, L5944, Key5945, C5946)
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

		tmp869 := PrimValue(symshen_4_dsigf_d)

		tmp870 := Call(__e, PrimFunc(symshen_4assoc_1_6), symand, tmp862, tmp869)

		tmp871 := PrimSet(symshen_4_dsigf_d, tmp870)

		_ = tmp871

		tmp872 := MakeNative(func(__e *ControlFlow) {
			V5952 := __e.Get(1)
			_ = V5952
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5948 := __e.Get(1)
				_ = B5948
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5949 := __e.Get(1)
					_ = L5949
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5950 := __e.Get(1)
						_ = Key5950
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5951 := __e.Get(1)
							_ = C5951
							tmp873 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp874 := PrimCons(symstring, Nil)

								tmp875 := PrimCons(sym_1_1_6, tmp874)

								tmp876 := PrimCons(symsymbol, tmp875)

								tmp877 := PrimCons(tmp876, Nil)

								tmp878 := PrimCons(sym_1_1_6, tmp877)

								tmp879 := PrimCons(symstring, tmp878)

								tmp880 := PrimCons(tmp879, Nil)

								tmp881 := PrimCons(sym_1_1_6, tmp880)

								tmp882 := PrimCons(A, tmp881)

								tmp883 := Call(__e, PrimFunc(symis_b), V5952, tmp882, B5948, L5949, Key5950, C5951)

								__e.TailApply(PrimFunc(symshen_4gc), B5948, tmp883)
								return

							}, 1)

							tmp884 := Call(__e, PrimFunc(symshen_4newpv), B5948)

							__e.TailApply(tmp873, tmp884)
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

		tmp885 := PrimValue(symshen_4_dsigf_d)

		tmp886 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4app, tmp872, tmp885)

		tmp887 := PrimSet(symshen_4_dsigf_d, tmp886)

		_ = tmp887

		tmp888 := MakeNative(func(__e *ControlFlow) {
			V5957 := __e.Get(1)
			_ = V5957
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5953 := __e.Get(1)
				_ = B5953
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5954 := __e.Get(1)
					_ = L5954
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5955 := __e.Get(1)
						_ = Key5955
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5956 := __e.Get(1)
							_ = C5956
							tmp889 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp890 := PrimCons(A, Nil)

								tmp891 := PrimCons(symlist, tmp890)

								tmp892 := PrimCons(A, Nil)

								tmp893 := PrimCons(symlist, tmp892)

								tmp894 := PrimCons(A, Nil)

								tmp895 := PrimCons(symlist, tmp894)

								tmp896 := PrimCons(tmp895, Nil)

								tmp897 := PrimCons(sym_1_1_6, tmp896)

								tmp898 := PrimCons(tmp893, tmp897)

								tmp899 := PrimCons(tmp898, Nil)

								tmp900 := PrimCons(sym_1_1_6, tmp899)

								tmp901 := PrimCons(tmp891, tmp900)

								tmp902 := Call(__e, PrimFunc(symis_b), V5957, tmp901, B5953, L5954, Key5955, C5956)

								__e.TailApply(PrimFunc(symshen_4gc), B5953, tmp902)
								return

							}, 1)

							tmp903 := Call(__e, PrimFunc(symshen_4newpv), B5953)

							__e.TailApply(tmp889, tmp903)
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

		tmp904 := PrimValue(symshen_4_dsigf_d)

		tmp905 := Call(__e, PrimFunc(symshen_4assoc_1_6), symappend, tmp888, tmp904)

		tmp906 := PrimSet(symshen_4_dsigf_d, tmp905)

		_ = tmp906

		tmp907 := MakeNative(func(__e *ControlFlow) {
			V5962 := __e.Get(1)
			_ = V5962
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5958 := __e.Get(1)
				_ = B5958
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5959 := __e.Get(1)
					_ = L5959
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5960 := __e.Get(1)
						_ = Key5960
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5961 := __e.Get(1)
							_ = C5961
							tmp908 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp909 := PrimCons(symnumber, Nil)

								tmp910 := PrimCons(sym_1_1_6, tmp909)

								tmp911 := PrimCons(A, tmp910)

								tmp912 := Call(__e, PrimFunc(symis_b), V5962, tmp911, B5958, L5959, Key5960, C5961)

								__e.TailApply(PrimFunc(symshen_4gc), B5958, tmp912)
								return

							}, 1)

							tmp913 := Call(__e, PrimFunc(symshen_4newpv), B5958)

							__e.TailApply(tmp908, tmp913)
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

		tmp914 := PrimValue(symshen_4_dsigf_d)

		tmp915 := Call(__e, PrimFunc(symshen_4assoc_1_6), symarity, tmp907, tmp914)

		tmp916 := PrimSet(symshen_4_dsigf_d, tmp915)

		_ = tmp916

		tmp917 := MakeNative(func(__e *ControlFlow) {
			V5967 := __e.Get(1)
			_ = V5967
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5963 := __e.Get(1)
				_ = B5963
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5964 := __e.Get(1)
					_ = L5964
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5965 := __e.Get(1)
						_ = Key5965
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5966 := __e.Get(1)
							_ = C5966
							tmp918 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp919 := PrimCons(A, Nil)

								tmp920 := PrimCons(symlist, tmp919)

								tmp921 := PrimCons(tmp920, Nil)

								tmp922 := PrimCons(symlist, tmp921)

								tmp923 := PrimCons(A, Nil)

								tmp924 := PrimCons(symlist, tmp923)

								tmp925 := PrimCons(tmp924, Nil)

								tmp926 := PrimCons(sym_1_1_6, tmp925)

								tmp927 := PrimCons(tmp922, tmp926)

								tmp928 := PrimCons(tmp927, Nil)

								tmp929 := PrimCons(sym_1_1_6, tmp928)

								tmp930 := PrimCons(A, tmp929)

								tmp931 := Call(__e, PrimFunc(symis_b), V5967, tmp930, B5963, L5964, Key5965, C5966)

								__e.TailApply(PrimFunc(symshen_4gc), B5963, tmp931)
								return

							}, 1)

							tmp932 := Call(__e, PrimFunc(symshen_4newpv), B5963)

							__e.TailApply(tmp918, tmp932)
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

		tmp933 := PrimValue(symshen_4_dsigf_d)

		tmp934 := Call(__e, PrimFunc(symshen_4assoc_1_6), symassoc, tmp917, tmp933)

		tmp935 := PrimSet(symshen_4_dsigf_d, tmp934)

		_ = tmp935

		tmp936 := MakeNative(func(__e *ControlFlow) {
			V5972 := __e.Get(1)
			_ = V5972
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5968 := __e.Get(1)
				_ = B5968
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5969 := __e.Get(1)
					_ = L5969
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5970 := __e.Get(1)
						_ = Key5970
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5971 := __e.Get(1)
							_ = C5971
							tmp937 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp938 := PrimCons(symboolean, Nil)

								tmp939 := PrimCons(sym_1_1_6, tmp938)

								tmp940 := PrimCons(A, tmp939)

								tmp941 := Call(__e, PrimFunc(symis_b), V5972, tmp940, B5968, L5969, Key5970, C5971)

								__e.TailApply(PrimFunc(symshen_4gc), B5968, tmp941)
								return

							}, 1)

							tmp942 := Call(__e, PrimFunc(symshen_4newpv), B5968)

							__e.TailApply(tmp937, tmp942)
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

		tmp943 := PrimValue(symshen_4_dsigf_d)

		tmp944 := Call(__e, PrimFunc(symshen_4assoc_1_6), symatom_2, tmp936, tmp943)

		tmp945 := PrimSet(symshen_4_dsigf_d, tmp944)

		_ = tmp945

		tmp946 := MakeNative(func(__e *ControlFlow) {
			V5977 := __e.Get(1)
			_ = V5977
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5973 := __e.Get(1)
				_ = B5973
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5974 := __e.Get(1)
					_ = L5974
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5975 := __e.Get(1)
						_ = Key5975
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5976 := __e.Get(1)
							_ = C5976
							tmp947 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp948 := PrimCons(symboolean, Nil)

								tmp949 := PrimCons(sym_1_1_6, tmp948)

								tmp950 := PrimCons(A, tmp949)

								tmp951 := Call(__e, PrimFunc(symis_b), V5977, tmp950, B5973, L5974, Key5975, C5976)

								__e.TailApply(PrimFunc(symshen_4gc), B5973, tmp951)
								return

							}, 1)

							tmp952 := Call(__e, PrimFunc(symshen_4newpv), B5973)

							__e.TailApply(tmp947, tmp952)
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

		tmp953 := PrimValue(symshen_4_dsigf_d)

		tmp954 := Call(__e, PrimFunc(symshen_4assoc_1_6), symboolean_2, tmp946, tmp953)

		tmp955 := PrimSet(symshen_4_dsigf_d, tmp954)

		_ = tmp955

		tmp956 := MakeNative(func(__e *ControlFlow) {
			V5982 := __e.Get(1)
			_ = V5982
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5978 := __e.Get(1)
				_ = B5978
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5979 := __e.Get(1)
					_ = L5979
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5980 := __e.Get(1)
						_ = Key5980
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5981 := __e.Get(1)
							_ = C5981
							tmp957 := PrimCons(symstring, Nil)

							tmp958 := PrimCons(sym_1_1_6, tmp957)

							tmp959 := PrimCons(symstring, tmp958)

							__e.TailApply(PrimFunc(symis_b), V5982, tmp959, B5978, L5979, Key5980, C5981)
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

		tmp960 := PrimValue(symshen_4_dsigf_d)

		tmp961 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbootstrap, tmp956, tmp960)

		tmp962 := PrimSet(symshen_4_dsigf_d, tmp961)

		_ = tmp962

		tmp963 := MakeNative(func(__e *ControlFlow) {
			V5987 := __e.Get(1)
			_ = V5987
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5983 := __e.Get(1)
				_ = B5983
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5984 := __e.Get(1)
					_ = L5984
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5985 := __e.Get(1)
						_ = Key5985
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5986 := __e.Get(1)
							_ = C5986
							tmp964 := PrimCons(symboolean, Nil)

							tmp965 := PrimCons(sym_1_1_6, tmp964)

							tmp966 := PrimCons(symsymbol, tmp965)

							__e.TailApply(PrimFunc(symis_b), V5987, tmp966, B5983, L5984, Key5985, C5986)
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

		tmp967 := PrimValue(symshen_4_dsigf_d)

		tmp968 := Call(__e, PrimFunc(symshen_4assoc_1_6), symbound_2, tmp963, tmp967)

		tmp969 := PrimSet(symshen_4_dsigf_d, tmp968)

		_ = tmp969

		tmp970 := MakeNative(func(__e *ControlFlow) {
			V5992 := __e.Get(1)
			_ = V5992
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5988 := __e.Get(1)
				_ = B5988
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5989 := __e.Get(1)
					_ = L5989
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5990 := __e.Get(1)
						_ = Key5990
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5991 := __e.Get(1)
							_ = C5991
							tmp971 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp972 := PrimCons(A, Nil)

								tmp973 := PrimCons(symlist, tmp972)

								tmp974 := PrimCons(symboolean, Nil)

								tmp975 := PrimCons(sym_1_1_6, tmp974)

								tmp976 := PrimCons(tmp973, tmp975)

								tmp977 := Call(__e, PrimFunc(symis_b), V5992, tmp976, B5988, L5989, Key5990, C5991)

								__e.TailApply(PrimFunc(symshen_4gc), B5988, tmp977)
								return

							}, 1)

							tmp978 := Call(__e, PrimFunc(symshen_4newpv), B5988)

							__e.TailApply(tmp971, tmp978)
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

		tmp979 := PrimValue(symshen_4_dsigf_d)

		tmp980 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4ccons_2, tmp970, tmp979)

		tmp981 := PrimSet(symshen_4_dsigf_d, tmp980)

		_ = tmp981

		tmp982 := MakeNative(func(__e *ControlFlow) {
			V5997 := __e.Get(1)
			_ = V5997
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5993 := __e.Get(1)
				_ = B5993
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5994 := __e.Get(1)
					_ = L5994
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key5995 := __e.Get(1)
						_ = Key5995
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C5996 := __e.Get(1)
							_ = C5996
							tmp983 := PrimCons(symstring, Nil)

							tmp984 := PrimCons(sym_1_1_6, tmp983)

							tmp985 := PrimCons(symstring, tmp984)

							__e.TailApply(PrimFunc(symis_b), V5997, tmp985, B5993, L5994, Key5995, C5996)
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

		tmp986 := PrimValue(symshen_4_dsigf_d)

		tmp987 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcd, tmp982, tmp986)

		tmp988 := PrimSet(symshen_4_dsigf_d, tmp987)

		_ = tmp988

		tmp989 := MakeNative(func(__e *ControlFlow) {
			V6002 := __e.Get(1)
			_ = V6002
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B5998 := __e.Get(1)
				_ = B5998
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L5999 := __e.Get(1)
					_ = L5999
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6000 := __e.Get(1)
						_ = Key6000
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6001 := __e.Get(1)
							_ = C6001
							tmp990 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp991 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp992 := PrimCons(A, Nil)

									tmp993 := PrimCons(symstream, tmp992)

									tmp994 := PrimCons(B, Nil)

									tmp995 := PrimCons(symlist, tmp994)

									tmp996 := PrimCons(tmp995, Nil)

									tmp997 := PrimCons(sym_1_1_6, tmp996)

									tmp998 := PrimCons(tmp993, tmp997)

									tmp999 := Call(__e, PrimFunc(symis_b), V6002, tmp998, B5998, L5999, Key6000, C6001)

									__e.TailApply(PrimFunc(symshen_4gc), B5998, tmp999)
									return

								}, 1)

								tmp1000 := Call(__e, PrimFunc(symshen_4newpv), B5998)

								tmp1001 := Call(__e, tmp991, tmp1000)

								__e.TailApply(PrimFunc(symshen_4gc), B5998, tmp1001)
								return

							}, 1)

							tmp1002 := Call(__e, PrimFunc(symshen_4newpv), B5998)

							__e.TailApply(tmp990, tmp1002)
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

		tmp1003 := PrimValue(symshen_4_dsigf_d)

		tmp1004 := Call(__e, PrimFunc(symshen_4assoc_1_6), symclose, tmp989, tmp1003)

		tmp1005 := PrimSet(symshen_4_dsigf_d, tmp1004)

		_ = tmp1005

		tmp1006 := MakeNative(func(__e *ControlFlow) {
			V6007 := __e.Get(1)
			_ = V6007
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6003 := __e.Get(1)
				_ = B6003
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6004 := __e.Get(1)
					_ = L6004
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6005 := __e.Get(1)
						_ = Key6005
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6006 := __e.Get(1)
							_ = C6006
							tmp1007 := PrimCons(symstring, Nil)

							tmp1008 := PrimCons(sym_1_1_6, tmp1007)

							tmp1009 := PrimCons(symstring, tmp1008)

							tmp1010 := PrimCons(tmp1009, Nil)

							tmp1011 := PrimCons(sym_1_1_6, tmp1010)

							tmp1012 := PrimCons(symstring, tmp1011)

							__e.TailApply(PrimFunc(symis_b), V6007, tmp1012, B6003, L6004, Key6005, C6006)
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

		tmp1013 := PrimValue(symshen_4_dsigf_d)

		tmp1014 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcn, tmp1006, tmp1013)

		tmp1015 := PrimSet(symshen_4_dsigf_d, tmp1014)

		_ = tmp1015

		tmp1016 := MakeNative(func(__e *ControlFlow) {
			V6012 := __e.Get(1)
			_ = V6012
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6008 := __e.Get(1)
				_ = B6008
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6009 := __e.Get(1)
					_ = L6009
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6010 := __e.Get(1)
						_ = Key6010
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6011 := __e.Get(1)
							_ = C6011
							tmp1017 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1018 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1019 := PrimCons(A, Nil)

									tmp1020 := PrimCons(symlist, tmp1019)

									tmp1021 := PrimCons(A, Nil)

									tmp1022 := PrimCons(symlist, tmp1021)

									tmp1023 := PrimCons(B, Nil)

									tmp1024 := PrimCons(tmp1022, tmp1023)

									tmp1025 := PrimCons(symstr, tmp1024)

									tmp1026 := PrimCons(tmp1025, Nil)

									tmp1027 := PrimCons(sym_1_1_6, tmp1026)

									tmp1028 := PrimCons(tmp1020, tmp1027)

									tmp1029 := PrimCons(A, Nil)

									tmp1030 := PrimCons(symlist, tmp1029)

									tmp1031 := PrimCons(B, Nil)

									tmp1032 := PrimCons(sym_1_1_6, tmp1031)

									tmp1033 := PrimCons(tmp1030, tmp1032)

									tmp1034 := PrimCons(tmp1033, Nil)

									tmp1035 := PrimCons(sym_1_1_6, tmp1034)

									tmp1036 := PrimCons(tmp1028, tmp1035)

									tmp1037 := Call(__e, PrimFunc(symis_b), V6012, tmp1036, B6008, L6009, Key6010, C6011)

									__e.TailApply(PrimFunc(symshen_4gc), B6008, tmp1037)
									return

								}, 1)

								tmp1038 := Call(__e, PrimFunc(symshen_4newpv), B6008)

								tmp1039 := Call(__e, tmp1018, tmp1038)

								__e.TailApply(PrimFunc(symshen_4gc), B6008, tmp1039)
								return

							}, 1)

							tmp1040 := Call(__e, PrimFunc(symshen_4newpv), B6008)

							__e.TailApply(tmp1017, tmp1040)
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

		tmp1041 := PrimValue(symshen_4_dsigf_d)

		tmp1042 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcompile, tmp1016, tmp1041)

		tmp1043 := PrimSet(symshen_4_dsigf_d, tmp1042)

		_ = tmp1043

		tmp1044 := MakeNative(func(__e *ControlFlow) {
			V6017 := __e.Get(1)
			_ = V6017
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6013 := __e.Get(1)
				_ = B6013
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6014 := __e.Get(1)
					_ = L6014
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6015 := __e.Get(1)
						_ = Key6015
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6016 := __e.Get(1)
							_ = C6016
							tmp1045 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1046 := PrimCons(symboolean, Nil)

								tmp1047 := PrimCons(sym_1_1_6, tmp1046)

								tmp1048 := PrimCons(A, tmp1047)

								tmp1049 := Call(__e, PrimFunc(symis_b), V6017, tmp1048, B6013, L6014, Key6015, C6016)

								__e.TailApply(PrimFunc(symshen_4gc), B6013, tmp1049)
								return

							}, 1)

							tmp1050 := Call(__e, PrimFunc(symshen_4newpv), B6013)

							__e.TailApply(tmp1045, tmp1050)
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

		tmp1051 := PrimValue(symshen_4_dsigf_d)

		tmp1052 := Call(__e, PrimFunc(symshen_4assoc_1_6), symcons_2, tmp1044, tmp1051)

		tmp1053 := PrimSet(symshen_4_dsigf_d, tmp1052)

		_ = tmp1053

		tmp1054 := MakeNative(func(__e *ControlFlow) {
			V6022 := __e.Get(1)
			_ = V6022
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6018 := __e.Get(1)
				_ = B6018
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6019 := __e.Get(1)
					_ = L6019
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6020 := __e.Get(1)
						_ = Key6020
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6021 := __e.Get(1)
							_ = C6021
							tmp1055 := PrimCons(symsymbol, Nil)

							tmp1056 := PrimCons(symlist, tmp1055)

							tmp1057 := PrimCons(tmp1056, Nil)

							tmp1058 := PrimCons(sym_1_1_6, tmp1057)

							__e.TailApply(PrimFunc(symis_b), V6022, tmp1058, B6018, L6019, Key6020, C6021)
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

		tmp1059 := PrimValue(symshen_4_dsigf_d)

		tmp1060 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdatatypes, tmp1054, tmp1059)

		tmp1061 := PrimSet(symshen_4_dsigf_d, tmp1060)

		_ = tmp1061

		tmp1062 := MakeNative(func(__e *ControlFlow) {
			V6027 := __e.Get(1)
			_ = V6027
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6023 := __e.Get(1)
				_ = B6023
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6024 := __e.Get(1)
					_ = L6024
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6025 := __e.Get(1)
						_ = Key6025
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6026 := __e.Get(1)
							_ = C6026
							tmp1063 := PrimCons(symsymbol, Nil)

							tmp1064 := PrimCons(sym_1_1_6, tmp1063)

							tmp1065 := PrimCons(symsymbol, tmp1064)

							__e.TailApply(PrimFunc(symis_b), V6027, tmp1065, B6023, L6024, Key6025, C6026)
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

		tmp1066 := PrimValue(symshen_4_dsigf_d)

		tmp1067 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdestroy, tmp1062, tmp1066)

		tmp1068 := PrimSet(symshen_4_dsigf_d, tmp1067)

		_ = tmp1068

		tmp1069 := MakeNative(func(__e *ControlFlow) {
			V6032 := __e.Get(1)
			_ = V6032
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6028 := __e.Get(1)
				_ = B6028
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6029 := __e.Get(1)
					_ = L6029
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6030 := __e.Get(1)
						_ = Key6030
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6031 := __e.Get(1)
							_ = C6031
							tmp1070 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1071 := PrimCons(A, Nil)

								tmp1072 := PrimCons(symlist, tmp1071)

								tmp1073 := PrimCons(A, Nil)

								tmp1074 := PrimCons(symlist, tmp1073)

								tmp1075 := PrimCons(A, Nil)

								tmp1076 := PrimCons(symlist, tmp1075)

								tmp1077 := PrimCons(tmp1076, Nil)

								tmp1078 := PrimCons(sym_1_1_6, tmp1077)

								tmp1079 := PrimCons(tmp1074, tmp1078)

								tmp1080 := PrimCons(tmp1079, Nil)

								tmp1081 := PrimCons(sym_1_1_6, tmp1080)

								tmp1082 := PrimCons(tmp1072, tmp1081)

								tmp1083 := Call(__e, PrimFunc(symis_b), V6032, tmp1082, B6028, L6029, Key6030, C6031)

								__e.TailApply(PrimFunc(symshen_4gc), B6028, tmp1083)
								return

							}, 1)

							tmp1084 := Call(__e, PrimFunc(symshen_4newpv), B6028)

							__e.TailApply(tmp1070, tmp1084)
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

		tmp1085 := PrimValue(symshen_4_dsigf_d)

		tmp1086 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdifference, tmp1069, tmp1085)

		tmp1087 := PrimSet(symshen_4_dsigf_d, tmp1086)

		_ = tmp1087

		tmp1088 := MakeNative(func(__e *ControlFlow) {
			V6037 := __e.Get(1)
			_ = V6037
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6033 := __e.Get(1)
				_ = B6033
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6034 := __e.Get(1)
					_ = L6034
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6035 := __e.Get(1)
						_ = Key6035
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6036 := __e.Get(1)
							_ = C6036
							tmp1089 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1090 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1091 := PrimCons(B, Nil)

									tmp1092 := PrimCons(sym_1_1_6, tmp1091)

									tmp1093 := PrimCons(B, tmp1092)

									tmp1094 := PrimCons(tmp1093, Nil)

									tmp1095 := PrimCons(sym_1_1_6, tmp1094)

									tmp1096 := PrimCons(A, tmp1095)

									tmp1097 := Call(__e, PrimFunc(symis_b), V6037, tmp1096, B6033, L6034, Key6035, C6036)

									__e.TailApply(PrimFunc(symshen_4gc), B6033, tmp1097)
									return

								}, 1)

								tmp1098 := Call(__e, PrimFunc(symshen_4newpv), B6033)

								tmp1099 := Call(__e, tmp1090, tmp1098)

								__e.TailApply(PrimFunc(symshen_4gc), B6033, tmp1099)
								return

							}, 1)

							tmp1100 := Call(__e, PrimFunc(symshen_4newpv), B6033)

							__e.TailApply(tmp1089, tmp1100)
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

		tmp1101 := PrimValue(symshen_4_dsigf_d)

		tmp1102 := Call(__e, PrimFunc(symshen_4assoc_1_6), symdo, tmp1088, tmp1101)

		tmp1103 := PrimSet(symshen_4_dsigf_d, tmp1102)

		_ = tmp1103

		tmp1104 := MakeNative(func(__e *ControlFlow) {
			V6042 := __e.Get(1)
			_ = V6042
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6038 := __e.Get(1)
				_ = B6038
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6039 := __e.Get(1)
					_ = L6039
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6040 := __e.Get(1)
						_ = Key6040
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6041 := __e.Get(1)
							_ = C6041
							tmp1105 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1106 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1107 := PrimCons(A, Nil)

									tmp1108 := PrimCons(symlist, tmp1107)

									tmp1109 := PrimCons(A, Nil)

									tmp1110 := PrimCons(symlist, tmp1109)

									tmp1111 := PrimCons(B, Nil)

									tmp1112 := PrimCons(symlist, tmp1111)

									tmp1113 := PrimCons(tmp1112, Nil)

									tmp1114 := PrimCons(tmp1110, tmp1113)

									tmp1115 := PrimCons(symstr, tmp1114)

									tmp1116 := PrimCons(tmp1115, Nil)

									tmp1117 := PrimCons(sym_1_1_6, tmp1116)

									tmp1118 := PrimCons(tmp1108, tmp1117)

									tmp1119 := Call(__e, PrimFunc(symis_b), V6042, tmp1118, B6038, L6039, Key6040, C6041)

									__e.TailApply(PrimFunc(symshen_4gc), B6038, tmp1119)
									return

								}, 1)

								tmp1120 := Call(__e, PrimFunc(symshen_4newpv), B6038)

								tmp1121 := Call(__e, tmp1106, tmp1120)

								__e.TailApply(PrimFunc(symshen_4gc), B6038, tmp1121)
								return

							}, 1)

							tmp1122 := Call(__e, PrimFunc(symshen_4newpv), B6038)

							__e.TailApply(tmp1105, tmp1122)
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

		tmp1123 := PrimValue(symshen_4_dsigf_d)

		tmp1124 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5e_6, tmp1104, tmp1123)

		tmp1125 := PrimSet(symshen_4_dsigf_d, tmp1124)

		_ = tmp1125

		tmp1126 := MakeNative(func(__e *ControlFlow) {
			V6047 := __e.Get(1)
			_ = V6047
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6043 := __e.Get(1)
				_ = B6043
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6044 := __e.Get(1)
					_ = L6044
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6045 := __e.Get(1)
						_ = Key6045
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6046 := __e.Get(1)
							_ = C6046
							tmp1127 := MakeNative(func(__e *ControlFlow) {
								B := __e.Get(1)
								_ = B
								tmp1128 := MakeNative(func(__e *ControlFlow) {
									A := __e.Get(1)
									_ = A
									tmp1129 := PrimCons(A, Nil)

									tmp1130 := PrimCons(symlist, tmp1129)

									tmp1131 := PrimCons(B, Nil)

									tmp1132 := PrimCons(symlist, tmp1131)

									tmp1133 := PrimCons(A, Nil)

									tmp1134 := PrimCons(symlist, tmp1133)

									tmp1135 := PrimCons(tmp1134, Nil)

									tmp1136 := PrimCons(tmp1132, tmp1135)

									tmp1137 := PrimCons(symstr, tmp1136)

									tmp1138 := PrimCons(tmp1137, Nil)

									tmp1139 := PrimCons(sym_1_1_6, tmp1138)

									tmp1140 := PrimCons(tmp1130, tmp1139)

									tmp1141 := Call(__e, PrimFunc(symis_b), V6047, tmp1140, B6043, L6044, Key6045, C6046)

									__e.TailApply(PrimFunc(symshen_4gc), B6043, tmp1141)
									return

								}, 1)

								tmp1142 := Call(__e, PrimFunc(symshen_4newpv), B6043)

								tmp1143 := Call(__e, tmp1128, tmp1142)

								__e.TailApply(PrimFunc(symshen_4gc), B6043, tmp1143)
								return

							}, 1)

							tmp1144 := Call(__e, PrimFunc(symshen_4newpv), B6043)

							__e.TailApply(tmp1127, tmp1144)
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

		tmp1145 := PrimValue(symshen_4_dsigf_d)

		tmp1146 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_b_6, tmp1126, tmp1145)

		tmp1147 := PrimSet(symshen_4_dsigf_d, tmp1146)

		_ = tmp1147

		tmp1148 := MakeNative(func(__e *ControlFlow) {
			V6052 := __e.Get(1)
			_ = V6052
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6048 := __e.Get(1)
				_ = B6048
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6049 := __e.Get(1)
					_ = L6049
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6050 := __e.Get(1)
						_ = Key6050
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6051 := __e.Get(1)
							_ = C6051
							tmp1149 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1150 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1151 := PrimCons(A, Nil)

									tmp1152 := PrimCons(symlist, tmp1151)

									tmp1153 := PrimCons(A, Nil)

									tmp1154 := PrimCons(symlist, tmp1153)

									tmp1155 := PrimCons(B, Nil)

									tmp1156 := PrimCons(symlist, tmp1155)

									tmp1157 := PrimCons(tmp1156, Nil)

									tmp1158 := PrimCons(tmp1154, tmp1157)

									tmp1159 := PrimCons(symstr, tmp1158)

									tmp1160 := PrimCons(tmp1159, Nil)

									tmp1161 := PrimCons(sym_1_1_6, tmp1160)

									tmp1162 := PrimCons(tmp1152, tmp1161)

									tmp1163 := Call(__e, PrimFunc(symis_b), V6052, tmp1162, B6048, L6049, Key6050, C6051)

									__e.TailApply(PrimFunc(symshen_4gc), B6048, tmp1163)
									return

								}, 1)

								tmp1164 := Call(__e, PrimFunc(symshen_4newpv), B6048)

								tmp1165 := Call(__e, tmp1150, tmp1164)

								__e.TailApply(PrimFunc(symshen_4gc), B6048, tmp1165)
								return

							}, 1)

							tmp1166 := Call(__e, PrimFunc(symshen_4newpv), B6048)

							__e.TailApply(tmp1149, tmp1166)
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

		tmp1167 := PrimValue(symshen_4_dsigf_d)

		tmp1168 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5end_6, tmp1148, tmp1167)

		tmp1169 := PrimSet(symshen_4_dsigf_d, tmp1168)

		_ = tmp1169

		tmp1170 := MakeNative(func(__e *ControlFlow) {
			V6057 := __e.Get(1)
			_ = V6057
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6053 := __e.Get(1)
				_ = B6053
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6054 := __e.Get(1)
					_ = L6054
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6055 := __e.Get(1)
						_ = Key6055
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6056 := __e.Get(1)
							_ = C6056
							tmp1171 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1172 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1173 := PrimCons(A, Nil)

									tmp1174 := PrimCons(symlist, tmp1173)

									tmp1175 := PrimCons(B, Nil)

									tmp1176 := PrimCons(tmp1174, tmp1175)

									tmp1177 := PrimCons(symstr, tmp1176)

									tmp1178 := PrimCons(symboolean, Nil)

									tmp1179 := PrimCons(sym_1_1_6, tmp1178)

									tmp1180 := PrimCons(tmp1177, tmp1179)

									tmp1181 := Call(__e, PrimFunc(symis_b), V6057, tmp1180, B6053, L6054, Key6055, C6056)

									__e.TailApply(PrimFunc(symshen_4gc), B6053, tmp1181)
									return

								}, 1)

								tmp1182 := Call(__e, PrimFunc(symshen_4newpv), B6053)

								tmp1183 := Call(__e, tmp1172, tmp1182)

								__e.TailApply(PrimFunc(symshen_4gc), B6053, tmp1183)
								return

							}, 1)

							tmp1184 := Call(__e, PrimFunc(symshen_4newpv), B6053)

							__e.TailApply(tmp1171, tmp1184)
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

		tmp1185 := PrimValue(symshen_4_dsigf_d)

		tmp1186 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure_2, tmp1170, tmp1185)

		tmp1187 := PrimSet(symshen_4_dsigf_d, tmp1186)

		_ = tmp1187

		tmp1188 := MakeNative(func(__e *ControlFlow) {
			V6062 := __e.Get(1)
			_ = V6062
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6058 := __e.Get(1)
				_ = B6058
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6059 := __e.Get(1)
					_ = L6059
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6060 := __e.Get(1)
						_ = Key6060
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6061 := __e.Get(1)
							_ = C6061
							tmp1189 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1190 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1191 := PrimCons(A, Nil)

									tmp1192 := PrimCons(symlist, tmp1191)

									tmp1193 := PrimCons(B, Nil)

									tmp1194 := PrimCons(tmp1192, tmp1193)

									tmp1195 := PrimCons(symstr, tmp1194)

									tmp1196 := PrimCons(tmp1195, Nil)

									tmp1197 := PrimCons(sym_1_1_6, tmp1196)

									tmp1198 := Call(__e, PrimFunc(symis_b), V6062, tmp1197, B6058, L6059, Key6060, C6061)

									__e.TailApply(PrimFunc(symshen_4gc), B6058, tmp1198)
									return

								}, 1)

								tmp1199 := Call(__e, PrimFunc(symshen_4newpv), B6058)

								tmp1200 := Call(__e, tmp1190, tmp1199)

								__e.TailApply(PrimFunc(symshen_4gc), B6058, tmp1200)
								return

							}, 1)

							tmp1201 := Call(__e, PrimFunc(symshen_4newpv), B6058)

							__e.TailApply(tmp1189, tmp1201)
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

		tmp1202 := PrimValue(symshen_4_dsigf_d)

		tmp1203 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4parse_1failure, tmp1188, tmp1202)

		tmp1204 := PrimSet(symshen_4_dsigf_d, tmp1203)

		_ = tmp1204

		tmp1205 := MakeNative(func(__e *ControlFlow) {
			V6067 := __e.Get(1)
			_ = V6067
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6063 := __e.Get(1)
				_ = B6063
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6064 := __e.Get(1)
					_ = L6064
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6065 := __e.Get(1)
						_ = Key6065
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6066 := __e.Get(1)
							_ = C6066
							tmp1206 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1207 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1208 := PrimCons(A, Nil)

									tmp1209 := PrimCons(symlist, tmp1208)

									tmp1210 := PrimCons(B, Nil)

									tmp1211 := PrimCons(tmp1209, tmp1210)

									tmp1212 := PrimCons(symstr, tmp1211)

									tmp1213 := PrimCons(B, Nil)

									tmp1214 := PrimCons(sym_1_1_6, tmp1213)

									tmp1215 := PrimCons(tmp1212, tmp1214)

									tmp1216 := Call(__e, PrimFunc(symis_b), V6067, tmp1215, B6063, L6064, Key6065, C6066)

									__e.TailApply(PrimFunc(symshen_4gc), B6063, tmp1216)
									return

								}, 1)

								tmp1217 := Call(__e, PrimFunc(symshen_4newpv), B6063)

								tmp1218 := Call(__e, tmp1207, tmp1217)

								__e.TailApply(PrimFunc(symshen_4gc), B6063, tmp1218)
								return

							}, 1)

							tmp1219 := Call(__e, PrimFunc(symshen_4newpv), B6063)

							__e.TailApply(tmp1206, tmp1219)
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

		tmp1220 := PrimValue(symshen_4_dsigf_d)

		tmp1221 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4_5_1out, tmp1205, tmp1220)

		tmp1222 := PrimSet(symshen_4_dsigf_d, tmp1221)

		_ = tmp1222

		tmp1223 := MakeNative(func(__e *ControlFlow) {
			V6072 := __e.Get(1)
			_ = V6072
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6068 := __e.Get(1)
				_ = B6068
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6069 := __e.Get(1)
					_ = L6069
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6070 := __e.Get(1)
						_ = Key6070
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6071 := __e.Get(1)
							_ = C6071
							tmp1224 := MakeNative(func(__e *ControlFlow) {
								B := __e.Get(1)
								_ = B
								tmp1225 := MakeNative(func(__e *ControlFlow) {
									A := __e.Get(1)
									_ = A
									tmp1226 := PrimCons(A, Nil)

									tmp1227 := PrimCons(symlist, tmp1226)

									tmp1228 := PrimCons(B, Nil)

									tmp1229 := PrimCons(tmp1227, tmp1228)

									tmp1230 := PrimCons(symstr, tmp1229)

									tmp1231 := PrimCons(A, Nil)

									tmp1232 := PrimCons(symlist, tmp1231)

									tmp1233 := PrimCons(tmp1232, Nil)

									tmp1234 := PrimCons(sym_1_1_6, tmp1233)

									tmp1235 := PrimCons(tmp1230, tmp1234)

									tmp1236 := Call(__e, PrimFunc(symis_b), V6072, tmp1235, B6068, L6069, Key6070, C6071)

									__e.TailApply(PrimFunc(symshen_4gc), B6068, tmp1236)
									return

								}, 1)

								tmp1237 := Call(__e, PrimFunc(symshen_4newpv), B6068)

								tmp1238 := Call(__e, tmp1225, tmp1237)

								__e.TailApply(PrimFunc(symshen_4gc), B6068, tmp1238)
								return

							}, 1)

							tmp1239 := Call(__e, PrimFunc(symshen_4newpv), B6068)

							__e.TailApply(tmp1224, tmp1239)
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

		tmp1240 := PrimValue(symshen_4_dsigf_d)

		tmp1241 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4in_1_6, tmp1223, tmp1240)

		tmp1242 := PrimSet(symshen_4_dsigf_d, tmp1241)

		_ = tmp1242

		tmp1243 := MakeNative(func(__e *ControlFlow) {
			V6077 := __e.Get(1)
			_ = V6077
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6073 := __e.Get(1)
				_ = B6073
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6074 := __e.Get(1)
					_ = L6074
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6075 := __e.Get(1)
						_ = Key6075
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6076 := __e.Get(1)
							_ = C6076
							tmp1244 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1245 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1246 := PrimCons(A, Nil)

									tmp1247 := PrimCons(symlist, tmp1246)

									tmp1248 := PrimCons(A, Nil)

									tmp1249 := PrimCons(symlist, tmp1248)

									tmp1250 := PrimCons(B, Nil)

									tmp1251 := PrimCons(tmp1249, tmp1250)

									tmp1252 := PrimCons(symstr, tmp1251)

									tmp1253 := PrimCons(tmp1252, Nil)

									tmp1254 := PrimCons(sym_1_1_6, tmp1253)

									tmp1255 := PrimCons(B, tmp1254)

									tmp1256 := PrimCons(tmp1255, Nil)

									tmp1257 := PrimCons(sym_1_1_6, tmp1256)

									tmp1258 := PrimCons(tmp1247, tmp1257)

									tmp1259 := Call(__e, PrimFunc(symis_b), V6077, tmp1258, B6073, L6074, Key6075, C6076)

									__e.TailApply(PrimFunc(symshen_4gc), B6073, tmp1259)
									return

								}, 1)

								tmp1260 := Call(__e, PrimFunc(symshen_4newpv), B6073)

								tmp1261 := Call(__e, tmp1245, tmp1260)

								__e.TailApply(PrimFunc(symshen_4gc), B6073, tmp1261)
								return

							}, 1)

							tmp1262 := Call(__e, PrimFunc(symshen_4newpv), B6073)

							__e.TailApply(tmp1244, tmp1262)
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

		tmp1263 := PrimValue(symshen_4_dsigf_d)

		tmp1264 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4comb, tmp1243, tmp1263)

		tmp1265 := PrimSet(symshen_4_dsigf_d, tmp1264)

		_ = tmp1265

		tmp1266 := MakeNative(func(__e *ControlFlow) {
			V6082 := __e.Get(1)
			_ = V6082
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6078 := __e.Get(1)
				_ = B6078
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6079 := __e.Get(1)
					_ = L6079
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6080 := __e.Get(1)
						_ = Key6080
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6081 := __e.Get(1)
							_ = C6081
							tmp1267 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1268 := PrimCons(A, Nil)

								tmp1269 := PrimCons(symlist, tmp1268)

								tmp1270 := PrimCons(symboolean, Nil)

								tmp1271 := PrimCons(sym_1_1_6, tmp1270)

								tmp1272 := PrimCons(tmp1269, tmp1271)

								tmp1273 := PrimCons(tmp1272, Nil)

								tmp1274 := PrimCons(sym_1_1_6, tmp1273)

								tmp1275 := PrimCons(A, tmp1274)

								tmp1276 := Call(__e, PrimFunc(symis_b), V6082, tmp1275, B6078, L6079, Key6080, C6081)

								__e.TailApply(PrimFunc(symshen_4gc), B6078, tmp1276)
								return

							}, 1)

							tmp1277 := Call(__e, PrimFunc(symshen_4newpv), B6078)

							__e.TailApply(tmp1267, tmp1277)
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

		tmp1278 := PrimValue(symshen_4_dsigf_d)

		tmp1279 := Call(__e, PrimFunc(symshen_4assoc_1_6), symelement_2, tmp1266, tmp1278)

		tmp1280 := PrimSet(symshen_4_dsigf_d, tmp1279)

		_ = tmp1280

		tmp1281 := MakeNative(func(__e *ControlFlow) {
			V6087 := __e.Get(1)
			_ = V6087
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6083 := __e.Get(1)
				_ = B6083
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6084 := __e.Get(1)
					_ = L6084
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6085 := __e.Get(1)
						_ = Key6085
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6086 := __e.Get(1)
							_ = C6086
							tmp1282 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1283 := PrimCons(symboolean, Nil)

								tmp1284 := PrimCons(sym_1_1_6, tmp1283)

								tmp1285 := PrimCons(A, tmp1284)

								tmp1286 := Call(__e, PrimFunc(symis_b), V6087, tmp1285, B6083, L6084, Key6085, C6086)

								__e.TailApply(PrimFunc(symshen_4gc), B6083, tmp1286)
								return

							}, 1)

							tmp1287 := Call(__e, PrimFunc(symshen_4newpv), B6083)

							__e.TailApply(tmp1282, tmp1287)
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

		tmp1288 := PrimValue(symshen_4_dsigf_d)

		tmp1289 := Call(__e, PrimFunc(symshen_4assoc_1_6), symempty_2, tmp1281, tmp1288)

		tmp1290 := PrimSet(symshen_4_dsigf_d, tmp1289)

		_ = tmp1290

		tmp1291 := MakeNative(func(__e *ControlFlow) {
			V6092 := __e.Get(1)
			_ = V6092
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6088 := __e.Get(1)
				_ = B6088
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6089 := __e.Get(1)
					_ = L6089
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6090 := __e.Get(1)
						_ = Key6090
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6091 := __e.Get(1)
							_ = C6091
							tmp1292 := PrimCons(symboolean, Nil)

							tmp1293 := PrimCons(sym_1_1_6, tmp1292)

							tmp1294 := PrimCons(symsymbol, tmp1293)

							__e.TailApply(PrimFunc(symis_b), V6092, tmp1294, B6088, L6089, Key6090, C6091)
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

		tmp1295 := PrimValue(symshen_4_dsigf_d)

		tmp1296 := Call(__e, PrimFunc(symshen_4assoc_1_6), symenable_1type_1theory, tmp1291, tmp1295)

		tmp1297 := PrimSet(symshen_4_dsigf_d, tmp1296)

		_ = tmp1297

		tmp1298 := MakeNative(func(__e *ControlFlow) {
			V6097 := __e.Get(1)
			_ = V6097
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6093 := __e.Get(1)
				_ = B6093
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6094 := __e.Get(1)
					_ = L6094
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6095 := __e.Get(1)
						_ = Key6095
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6096 := __e.Get(1)
							_ = C6096
							tmp1299 := PrimCons(symsymbol, Nil)

							tmp1300 := PrimCons(symlist, tmp1299)

							tmp1301 := PrimCons(tmp1300, Nil)

							tmp1302 := PrimCons(sym_1_1_6, tmp1301)

							tmp1303 := PrimCons(symsymbol, tmp1302)

							__e.TailApply(PrimFunc(symis_b), V6097, tmp1303, B6093, L6094, Key6095, C6096)
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

		tmp1304 := PrimValue(symshen_4_dsigf_d)

		tmp1305 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexternal, tmp1298, tmp1304)

		tmp1306 := PrimSet(symshen_4_dsigf_d, tmp1305)

		_ = tmp1306

		tmp1307 := MakeNative(func(__e *ControlFlow) {
			V6102 := __e.Get(1)
			_ = V6102
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6098 := __e.Get(1)
				_ = B6098
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6099 := __e.Get(1)
					_ = L6099
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6100 := __e.Get(1)
						_ = Key6100
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6101 := __e.Get(1)
							_ = C6101
							tmp1308 := PrimCons(symstring, Nil)

							tmp1309 := PrimCons(sym_1_1_6, tmp1308)

							tmp1310 := PrimCons(symexception, tmp1309)

							__e.TailApply(PrimFunc(symis_b), V6102, tmp1310, B6098, L6099, Key6100, C6101)
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

		tmp1311 := PrimValue(symshen_4_dsigf_d)

		tmp1312 := Call(__e, PrimFunc(symshen_4assoc_1_6), symerror_1to_1string, tmp1307, tmp1311)

		tmp1313 := PrimSet(symshen_4_dsigf_d, tmp1312)

		_ = tmp1313

		tmp1314 := MakeNative(func(__e *ControlFlow) {
			V6107 := __e.Get(1)
			_ = V6107
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6103 := __e.Get(1)
				_ = B6103
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6104 := __e.Get(1)
					_ = L6104
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6105 := __e.Get(1)
						_ = Key6105
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6106 := __e.Get(1)
							_ = C6106
							tmp1315 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1316 := PrimCons(symstring, Nil)

								tmp1317 := PrimCons(symlist, tmp1316)

								tmp1318 := PrimCons(tmp1317, Nil)

								tmp1319 := PrimCons(sym_1_1_6, tmp1318)

								tmp1320 := PrimCons(A, tmp1319)

								tmp1321 := Call(__e, PrimFunc(symis_b), V6107, tmp1320, B6103, L6104, Key6105, C6106)

								__e.TailApply(PrimFunc(symshen_4gc), B6103, tmp1321)
								return

							}, 1)

							tmp1322 := Call(__e, PrimFunc(symshen_4newpv), B6103)

							__e.TailApply(tmp1315, tmp1322)
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

		tmp1323 := PrimValue(symshen_4_dsigf_d)

		tmp1324 := Call(__e, PrimFunc(symshen_4assoc_1_6), symexplode, tmp1314, tmp1323)

		tmp1325 := PrimSet(symshen_4_dsigf_d, tmp1324)

		_ = tmp1325

		tmp1326 := MakeNative(func(__e *ControlFlow) {
			V6112 := __e.Get(1)
			_ = V6112
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6108 := __e.Get(1)
				_ = B6108
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6109 := __e.Get(1)
					_ = L6109
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6110 := __e.Get(1)
						_ = Key6110
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6111 := __e.Get(1)
							_ = C6111
							tmp1327 := PrimCons(symsymbol, Nil)

							tmp1328 := PrimCons(sym_1_1_6, tmp1327)

							tmp1329 := PrimCons(symsymbol, tmp1328)

							__e.TailApply(PrimFunc(symis_b), V6112, tmp1329, B6108, L6109, Key6110, C6111)
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

		tmp1330 := PrimValue(symshen_4_dsigf_d)

		tmp1331 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise, tmp1326, tmp1330)

		tmp1332 := PrimSet(symshen_4_dsigf_d, tmp1331)

		_ = tmp1332

		tmp1333 := MakeNative(func(__e *ControlFlow) {
			V6117 := __e.Get(1)
			_ = V6117
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6113 := __e.Get(1)
				_ = B6113
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6114 := __e.Get(1)
					_ = L6114
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6115 := __e.Get(1)
						_ = Key6115
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6116 := __e.Get(1)
							_ = C6116
							tmp1334 := PrimCons(symboolean, Nil)

							tmp1335 := PrimCons(sym_1_1_6, tmp1334)

							__e.TailApply(PrimFunc(symis_b), V6117, tmp1335, B6113, L6114, Key6115, C6116)
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

		tmp1336 := PrimValue(symshen_4_dsigf_d)

		tmp1337 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfactorise_2, tmp1333, tmp1336)

		tmp1338 := PrimSet(symshen_4_dsigf_d, tmp1337)

		_ = tmp1338

		tmp1339 := MakeNative(func(__e *ControlFlow) {
			V6122 := __e.Get(1)
			_ = V6122
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6118 := __e.Get(1)
				_ = B6118
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6119 := __e.Get(1)
					_ = L6119
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6120 := __e.Get(1)
						_ = Key6120
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6121 := __e.Get(1)
							_ = C6121
							tmp1340 := PrimCons(symsymbol, Nil)

							tmp1341 := PrimCons(sym_1_1_6, tmp1340)

							__e.TailApply(PrimFunc(symis_b), V6122, tmp1341, B6118, L6119, Key6120, C6121)
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

		tmp1342 := PrimValue(symshen_4_dsigf_d)

		tmp1343 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfail, tmp1339, tmp1342)

		tmp1344 := PrimSet(symshen_4_dsigf_d, tmp1343)

		_ = tmp1344

		tmp1345 := MakeNative(func(__e *ControlFlow) {
			V6127 := __e.Get(1)
			_ = V6127
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6123 := __e.Get(1)
				_ = B6123
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6124 := __e.Get(1)
					_ = L6124
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6125 := __e.Get(1)
						_ = Key6125
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6126 := __e.Get(1)
							_ = C6126
							tmp1346 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1347 := PrimCons(A, Nil)

								tmp1348 := PrimCons(sym_1_1_6, tmp1347)

								tmp1349 := PrimCons(A, tmp1348)

								tmp1350 := PrimCons(A, Nil)

								tmp1351 := PrimCons(sym_1_1_6, tmp1350)

								tmp1352 := PrimCons(A, tmp1351)

								tmp1353 := PrimCons(tmp1352, Nil)

								tmp1354 := PrimCons(sym_1_1_6, tmp1353)

								tmp1355 := PrimCons(tmp1349, tmp1354)

								tmp1356 := Call(__e, PrimFunc(symis_b), V6127, tmp1355, B6123, L6124, Key6125, C6126)

								__e.TailApply(PrimFunc(symshen_4gc), B6123, tmp1356)
								return

							}, 1)

							tmp1357 := Call(__e, PrimFunc(symshen_4newpv), B6123)

							__e.TailApply(tmp1346, tmp1357)
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

		tmp1358 := PrimValue(symshen_4_dsigf_d)

		tmp1359 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfix, tmp1345, tmp1358)

		tmp1360 := PrimSet(symshen_4_dsigf_d, tmp1359)

		_ = tmp1360

		tmp1361 := MakeNative(func(__e *ControlFlow) {
			V6132 := __e.Get(1)
			_ = V6132
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6128 := __e.Get(1)
				_ = B6128
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6129 := __e.Get(1)
					_ = L6129
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6130 := __e.Get(1)
						_ = Key6130
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6131 := __e.Get(1)
							_ = C6131
							tmp1362 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1363 := PrimCons(A, Nil)

								tmp1364 := PrimCons(symlazy, tmp1363)

								tmp1365 := PrimCons(tmp1364, Nil)

								tmp1366 := PrimCons(sym_1_1_6, tmp1365)

								tmp1367 := PrimCons(A, tmp1366)

								tmp1368 := Call(__e, PrimFunc(symis_b), V6132, tmp1367, B6128, L6129, Key6130, C6131)

								__e.TailApply(PrimFunc(symshen_4gc), B6128, tmp1368)
								return

							}, 1)

							tmp1369 := Call(__e, PrimFunc(symshen_4newpv), B6128)

							__e.TailApply(tmp1362, tmp1369)
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

		tmp1370 := PrimValue(symshen_4_dsigf_d)

		tmp1371 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfreeze, tmp1361, tmp1370)

		tmp1372 := PrimSet(symshen_4_dsigf_d, tmp1371)

		_ = tmp1372

		tmp1373 := MakeNative(func(__e *ControlFlow) {
			V6137 := __e.Get(1)
			_ = V6137
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6133 := __e.Get(1)
				_ = B6133
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6134 := __e.Get(1)
					_ = L6134
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6135 := __e.Get(1)
						_ = Key6135
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6136 := __e.Get(1)
							_ = C6136
							tmp1374 := MakeNative(func(__e *ControlFlow) {
								B := __e.Get(1)
								_ = B
								tmp1375 := MakeNative(func(__e *ControlFlow) {
									A := __e.Get(1)
									_ = A
									tmp1376 := PrimCons(B, Nil)

									tmp1377 := PrimCons(sym_d, tmp1376)

									tmp1378 := PrimCons(A, tmp1377)

									tmp1379 := PrimCons(A, Nil)

									tmp1380 := PrimCons(sym_1_1_6, tmp1379)

									tmp1381 := PrimCons(tmp1378, tmp1380)

									tmp1382 := Call(__e, PrimFunc(symis_b), V6137, tmp1381, B6133, L6134, Key6135, C6136)

									__e.TailApply(PrimFunc(symshen_4gc), B6133, tmp1382)
									return

								}, 1)

								tmp1383 := Call(__e, PrimFunc(symshen_4newpv), B6133)

								tmp1384 := Call(__e, tmp1375, tmp1383)

								__e.TailApply(PrimFunc(symshen_4gc), B6133, tmp1384)
								return

							}, 1)

							tmp1385 := Call(__e, PrimFunc(symshen_4newpv), B6133)

							__e.TailApply(tmp1374, tmp1385)
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

		tmp1386 := PrimValue(symshen_4_dsigf_d)

		tmp1387 := Call(__e, PrimFunc(symshen_4assoc_1_6), symfst, tmp1373, tmp1386)

		tmp1388 := PrimSet(symshen_4_dsigf_d, tmp1387)

		_ = tmp1388

		tmp1389 := MakeNative(func(__e *ControlFlow) {
			V6142 := __e.Get(1)
			_ = V6142
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6138 := __e.Get(1)
				_ = B6138
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6139 := __e.Get(1)
					_ = L6139
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6140 := __e.Get(1)
						_ = Key6140
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6141 := __e.Get(1)
							_ = C6141
							tmp1390 := PrimCons(symsymbol, Nil)

							tmp1391 := PrimCons(sym_1_1_6, tmp1390)

							tmp1392 := PrimCons(symsymbol, tmp1391)

							__e.TailApply(PrimFunc(symis_b), V6142, tmp1392, B6138, L6139, Key6140, C6141)
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

		tmp1393 := PrimValue(symshen_4_dsigf_d)

		tmp1394 := Call(__e, PrimFunc(symshen_4assoc_1_6), symgensym, tmp1389, tmp1393)

		tmp1395 := PrimSet(symshen_4_dsigf_d, tmp1394)

		_ = tmp1395

		tmp1396 := MakeNative(func(__e *ControlFlow) {
			V6147 := __e.Get(1)
			_ = V6147
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6143 := __e.Get(1)
				_ = B6143
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6144 := __e.Get(1)
					_ = L6144
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6145 := __e.Get(1)
						_ = Key6145
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6146 := __e.Get(1)
							_ = C6146
							tmp1397 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1398 := PrimCons(A, Nil)

								tmp1399 := PrimCons(symlist, tmp1398)

								tmp1400 := PrimCons(symboolean, Nil)

								tmp1401 := PrimCons(sym_1_1_6, tmp1400)

								tmp1402 := PrimCons(A, tmp1401)

								tmp1403 := PrimCons(tmp1402, Nil)

								tmp1404 := PrimCons(sym_1_1_6, tmp1403)

								tmp1405 := PrimCons(tmp1399, tmp1404)

								tmp1406 := Call(__e, PrimFunc(symis_b), V6147, tmp1405, B6143, L6144, Key6145, C6146)

								__e.TailApply(PrimFunc(symshen_4gc), B6143, tmp1406)
								return

							}, 1)

							tmp1407 := Call(__e, PrimFunc(symshen_4newpv), B6143)

							__e.TailApply(tmp1397, tmp1407)
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

		tmp1408 := PrimValue(symshen_4_dsigf_d)

		tmp1409 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4hds_a_2, tmp1396, tmp1408)

		tmp1410 := PrimSet(symshen_4_dsigf_d, tmp1409)

		_ = tmp1410

		tmp1411 := MakeNative(func(__e *ControlFlow) {
			V6152 := __e.Get(1)
			_ = V6152
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6148 := __e.Get(1)
				_ = B6148
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6149 := __e.Get(1)
					_ = L6149
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6150 := __e.Get(1)
						_ = Key6150
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6151 := __e.Get(1)
							_ = C6151
							tmp1412 := PrimCons(symboolean, Nil)

							tmp1413 := PrimCons(sym_1_1_6, tmp1412)

							tmp1414 := PrimCons(symsymbol, tmp1413)

							__e.TailApply(PrimFunc(symis_b), V6152, tmp1414, B6148, L6149, Key6150, C6151)
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

		tmp1415 := PrimValue(symshen_4_dsigf_d)

		tmp1416 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush, tmp1411, tmp1415)

		tmp1417 := PrimSet(symshen_4_dsigf_d, tmp1416)

		_ = tmp1417

		tmp1418 := MakeNative(func(__e *ControlFlow) {
			V6157 := __e.Get(1)
			_ = V6157
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6153 := __e.Get(1)
				_ = B6153
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6154 := __e.Get(1)
					_ = L6154
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6155 := __e.Get(1)
						_ = Key6155
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6156 := __e.Get(1)
							_ = C6156
							tmp1419 := PrimCons(symboolean, Nil)

							tmp1420 := PrimCons(sym_1_1_6, tmp1419)

							__e.TailApply(PrimFunc(symis_b), V6157, tmp1420, B6153, L6154, Key6155, C6156)
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

		tmp1421 := PrimValue(symshen_4_dsigf_d)

		tmp1422 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhush_2, tmp1418, tmp1421)

		tmp1423 := PrimSet(symshen_4_dsigf_d, tmp1422)

		_ = tmp1423

		tmp1424 := MakeNative(func(__e *ControlFlow) {
			V6162 := __e.Get(1)
			_ = V6162
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6158 := __e.Get(1)
				_ = B6158
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6159 := __e.Get(1)
					_ = L6159
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6160 := __e.Get(1)
						_ = Key6160
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6161 := __e.Get(1)
							_ = C6161
							tmp1425 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1426 := PrimCons(A, Nil)

								tmp1427 := PrimCons(symvector, tmp1426)

								tmp1428 := PrimCons(A, Nil)

								tmp1429 := PrimCons(sym_1_1_6, tmp1428)

								tmp1430 := PrimCons(symnumber, tmp1429)

								tmp1431 := PrimCons(tmp1430, Nil)

								tmp1432 := PrimCons(sym_1_1_6, tmp1431)

								tmp1433 := PrimCons(tmp1427, tmp1432)

								tmp1434 := Call(__e, PrimFunc(symis_b), V6162, tmp1433, B6158, L6159, Key6160, C6161)

								__e.TailApply(PrimFunc(symshen_4gc), B6158, tmp1434)
								return

							}, 1)

							tmp1435 := Call(__e, PrimFunc(symshen_4newpv), B6158)

							__e.TailApply(tmp1425, tmp1435)
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

		tmp1436 := PrimValue(symshen_4_dsigf_d)

		tmp1437 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_1vector, tmp1424, tmp1436)

		tmp1438 := PrimSet(symshen_4_dsigf_d, tmp1437)

		_ = tmp1438

		tmp1439 := MakeNative(func(__e *ControlFlow) {
			V6167 := __e.Get(1)
			_ = V6167
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6163 := __e.Get(1)
				_ = B6163
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6164 := __e.Get(1)
					_ = L6164
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6165 := __e.Get(1)
						_ = Key6165
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6166 := __e.Get(1)
							_ = C6166
							tmp1440 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1441 := PrimCons(A, Nil)

								tmp1442 := PrimCons(symvector, tmp1441)

								tmp1443 := PrimCons(A, Nil)

								tmp1444 := PrimCons(symvector, tmp1443)

								tmp1445 := PrimCons(tmp1444, Nil)

								tmp1446 := PrimCons(sym_1_1_6, tmp1445)

								tmp1447 := PrimCons(A, tmp1446)

								tmp1448 := PrimCons(tmp1447, Nil)

								tmp1449 := PrimCons(sym_1_1_6, tmp1448)

								tmp1450 := PrimCons(symnumber, tmp1449)

								tmp1451 := PrimCons(tmp1450, Nil)

								tmp1452 := PrimCons(sym_1_1_6, tmp1451)

								tmp1453 := PrimCons(tmp1442, tmp1452)

								tmp1454 := Call(__e, PrimFunc(symis_b), V6167, tmp1453, B6163, L6164, Key6165, C6166)

								__e.TailApply(PrimFunc(symshen_4gc), B6163, tmp1454)
								return

							}, 1)

							tmp1455 := Call(__e, PrimFunc(symshen_4newpv), B6163)

							__e.TailApply(tmp1440, tmp1455)
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

		tmp1456 := PrimValue(symshen_4_dsigf_d)

		tmp1457 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_1_6, tmp1439, tmp1456)

		tmp1458 := PrimSet(symshen_4_dsigf_d, tmp1457)

		_ = tmp1458

		tmp1459 := MakeNative(func(__e *ControlFlow) {
			V6172 := __e.Get(1)
			_ = V6172
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6168 := __e.Get(1)
				_ = B6168
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6169 := __e.Get(1)
					_ = L6169
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6170 := __e.Get(1)
						_ = Key6170
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6171 := __e.Get(1)
							_ = C6171
							tmp1460 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1461 := PrimCons(A, Nil)

								tmp1462 := PrimCons(symvector, tmp1461)

								tmp1463 := PrimCons(tmp1462, Nil)

								tmp1464 := PrimCons(sym_1_1_6, tmp1463)

								tmp1465 := PrimCons(symnumber, tmp1464)

								tmp1466 := Call(__e, PrimFunc(symis_b), V6172, tmp1465, B6168, L6169, Key6170, C6171)

								__e.TailApply(PrimFunc(symshen_4gc), B6168, tmp1466)
								return

							}, 1)

							tmp1467 := Call(__e, PrimFunc(symshen_4newpv), B6168)

							__e.TailApply(tmp1460, tmp1467)
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

		tmp1468 := PrimValue(symshen_4_dsigf_d)

		tmp1469 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector, tmp1459, tmp1468)

		tmp1470 := PrimSet(symshen_4_dsigf_d, tmp1469)

		_ = tmp1470

		tmp1471 := MakeNative(func(__e *ControlFlow) {
			V6177 := __e.Get(1)
			_ = V6177
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6173 := __e.Get(1)
				_ = B6173
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6174 := __e.Get(1)
					_ = L6174
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6175 := __e.Get(1)
						_ = Key6175
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6176 := __e.Get(1)
							_ = C6176
							tmp1472 := PrimCons(symnumber, Nil)

							tmp1473 := PrimCons(sym_1_1_6, tmp1472)

							tmp1474 := PrimCons(symsymbol, tmp1473)

							__e.TailApply(PrimFunc(symis_b), V6177, tmp1474, B6173, L6174, Key6175, C6176)
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

		tmp1475 := PrimValue(symshen_4_dsigf_d)

		tmp1476 := Call(__e, PrimFunc(symshen_4assoc_1_6), symget_1time, tmp1471, tmp1475)

		tmp1477 := PrimSet(symshen_4_dsigf_d, tmp1476)

		_ = tmp1477

		tmp1478 := MakeNative(func(__e *ControlFlow) {
			V6182 := __e.Get(1)
			_ = V6182
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6178 := __e.Get(1)
				_ = B6178
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6179 := __e.Get(1)
					_ = L6179
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6180 := __e.Get(1)
						_ = Key6180
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6181 := __e.Get(1)
							_ = C6181
							tmp1479 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1480 := PrimCons(symnumber, Nil)

								tmp1481 := PrimCons(sym_1_1_6, tmp1480)

								tmp1482 := PrimCons(symnumber, tmp1481)

								tmp1483 := PrimCons(tmp1482, Nil)

								tmp1484 := PrimCons(sym_1_1_6, tmp1483)

								tmp1485 := PrimCons(A, tmp1484)

								tmp1486 := Call(__e, PrimFunc(symis_b), V6182, tmp1485, B6178, L6179, Key6180, C6181)

								__e.TailApply(PrimFunc(symshen_4gc), B6178, tmp1486)
								return

							}, 1)

							tmp1487 := Call(__e, PrimFunc(symshen_4newpv), B6178)

							__e.TailApply(tmp1479, tmp1487)
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

		tmp1488 := PrimValue(symshen_4_dsigf_d)

		tmp1489 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhash, tmp1478, tmp1488)

		tmp1490 := PrimSet(symshen_4_dsigf_d, tmp1489)

		_ = tmp1490

		tmp1491 := MakeNative(func(__e *ControlFlow) {
			V6187 := __e.Get(1)
			_ = V6187
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6183 := __e.Get(1)
				_ = B6183
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6184 := __e.Get(1)
					_ = L6184
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6185 := __e.Get(1)
						_ = Key6185
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6186 := __e.Get(1)
							_ = C6186
							tmp1492 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1493 := PrimCons(A, Nil)

								tmp1494 := PrimCons(symlist, tmp1493)

								tmp1495 := PrimCons(A, Nil)

								tmp1496 := PrimCons(sym_1_1_6, tmp1495)

								tmp1497 := PrimCons(tmp1494, tmp1496)

								tmp1498 := Call(__e, PrimFunc(symis_b), V6187, tmp1497, B6183, L6184, Key6185, C6186)

								__e.TailApply(PrimFunc(symshen_4gc), B6183, tmp1498)
								return

							}, 1)

							tmp1499 := Call(__e, PrimFunc(symshen_4newpv), B6183)

							__e.TailApply(tmp1492, tmp1499)
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

		tmp1500 := PrimValue(symshen_4_dsigf_d)

		tmp1501 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhead, tmp1491, tmp1500)

		tmp1502 := PrimSet(symshen_4_dsigf_d, tmp1501)

		_ = tmp1502

		tmp1503 := MakeNative(func(__e *ControlFlow) {
			V6192 := __e.Get(1)
			_ = V6192
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6188 := __e.Get(1)
				_ = B6188
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6189 := __e.Get(1)
					_ = L6189
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6190 := __e.Get(1)
						_ = Key6190
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6191 := __e.Get(1)
							_ = C6191
							tmp1504 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1505 := PrimCons(A, Nil)

								tmp1506 := PrimCons(symvector, tmp1505)

								tmp1507 := PrimCons(A, Nil)

								tmp1508 := PrimCons(sym_1_1_6, tmp1507)

								tmp1509 := PrimCons(tmp1506, tmp1508)

								tmp1510 := Call(__e, PrimFunc(symis_b), V6192, tmp1509, B6188, L6189, Key6190, C6191)

								__e.TailApply(PrimFunc(symshen_4gc), B6188, tmp1510)
								return

							}, 1)

							tmp1511 := Call(__e, PrimFunc(symshen_4newpv), B6188)

							__e.TailApply(tmp1504, tmp1511)
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

		tmp1512 := PrimValue(symshen_4_dsigf_d)

		tmp1513 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdv, tmp1503, tmp1512)

		tmp1514 := PrimSet(symshen_4_dsigf_d, tmp1513)

		_ = tmp1514

		tmp1515 := MakeNative(func(__e *ControlFlow) {
			V6197 := __e.Get(1)
			_ = V6197
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6193 := __e.Get(1)
				_ = B6193
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6194 := __e.Get(1)
					_ = L6194
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6195 := __e.Get(1)
						_ = Key6195
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6196 := __e.Get(1)
							_ = C6196
							tmp1516 := PrimCons(symstring, Nil)

							tmp1517 := PrimCons(sym_1_1_6, tmp1516)

							tmp1518 := PrimCons(symstring, tmp1517)

							__e.TailApply(PrimFunc(symis_b), V6197, tmp1518, B6193, L6194, Key6195, C6196)
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

		tmp1519 := PrimValue(symshen_4_dsigf_d)

		tmp1520 := Call(__e, PrimFunc(symshen_4assoc_1_6), symhdstr, tmp1515, tmp1519)

		tmp1521 := PrimSet(symshen_4_dsigf_d, tmp1520)

		_ = tmp1521

		tmp1522 := MakeNative(func(__e *ControlFlow) {
			V6202 := __e.Get(1)
			_ = V6202
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6198 := __e.Get(1)
				_ = B6198
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6199 := __e.Get(1)
					_ = L6199
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6200 := __e.Get(1)
						_ = Key6200
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6201 := __e.Get(1)
							_ = C6201
							tmp1523 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1524 := PrimCons(A, Nil)

								tmp1525 := PrimCons(sym_1_1_6, tmp1524)

								tmp1526 := PrimCons(A, tmp1525)

								tmp1527 := PrimCons(tmp1526, Nil)

								tmp1528 := PrimCons(sym_1_1_6, tmp1527)

								tmp1529 := PrimCons(A, tmp1528)

								tmp1530 := PrimCons(tmp1529, Nil)

								tmp1531 := PrimCons(sym_1_1_6, tmp1530)

								tmp1532 := PrimCons(symboolean, tmp1531)

								tmp1533 := Call(__e, PrimFunc(symis_b), V6202, tmp1532, B6198, L6199, Key6200, C6201)

								__e.TailApply(PrimFunc(symshen_4gc), B6198, tmp1533)
								return

							}, 1)

							tmp1534 := Call(__e, PrimFunc(symshen_4newpv), B6198)

							__e.TailApply(tmp1523, tmp1534)
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

		tmp1535 := PrimValue(symshen_4_dsigf_d)

		tmp1536 := Call(__e, PrimFunc(symshen_4assoc_1_6), symif, tmp1522, tmp1535)

		tmp1537 := PrimSet(symshen_4_dsigf_d, tmp1536)

		_ = tmp1537

		tmp1538 := MakeNative(func(__e *ControlFlow) {
			V6207 := __e.Get(1)
			_ = V6207
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6203 := __e.Get(1)
				_ = B6203
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6204 := __e.Get(1)
					_ = L6204
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6205 := __e.Get(1)
						_ = Key6205
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6206 := __e.Get(1)
							_ = C6206
							tmp1539 := PrimCons(symsymbol, Nil)

							tmp1540 := PrimCons(sym_1_1_6, tmp1539)

							tmp1541 := PrimCons(symsymbol, tmp1540)

							__e.TailApply(PrimFunc(symis_b), V6207, tmp1541, B6203, L6204, Key6205, C6206)
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

		tmp1542 := PrimValue(symshen_4_dsigf_d)

		tmp1543 := Call(__e, PrimFunc(symshen_4assoc_1_6), symin_1package, tmp1538, tmp1542)

		tmp1544 := PrimSet(symshen_4_dsigf_d, tmp1543)

		_ = tmp1544

		tmp1545 := MakeNative(func(__e *ControlFlow) {
			V6212 := __e.Get(1)
			_ = V6212
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6208 := __e.Get(1)
				_ = B6208
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6209 := __e.Get(1)
					_ = L6209
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6210 := __e.Get(1)
						_ = Key6210
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6211 := __e.Get(1)
							_ = C6211
							tmp1546 := PrimCons(symstring, Nil)

							tmp1547 := PrimCons(sym_1_1_6, tmp1546)

							__e.TailApply(PrimFunc(symis_b), V6212, tmp1547, B6208, L6209, Key6210, C6211)
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

		tmp1548 := PrimValue(symshen_4_dsigf_d)

		tmp1549 := Call(__e, PrimFunc(symshen_4assoc_1_6), symit, tmp1545, tmp1548)

		tmp1550 := PrimSet(symshen_4_dsigf_d, tmp1549)

		_ = tmp1550

		tmp1551 := MakeNative(func(__e *ControlFlow) {
			V6217 := __e.Get(1)
			_ = V6217
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6213 := __e.Get(1)
				_ = B6213
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6214 := __e.Get(1)
					_ = L6214
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6215 := __e.Get(1)
						_ = Key6215
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6216 := __e.Get(1)
							_ = C6216
							tmp1552 := PrimCons(symstring, Nil)

							tmp1553 := PrimCons(sym_1_1_6, tmp1552)

							__e.TailApply(PrimFunc(symis_b), V6217, tmp1553, B6213, L6214, Key6215, C6216)
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

		tmp1554 := PrimValue(symshen_4_dsigf_d)

		tmp1555 := Call(__e, PrimFunc(symshen_4assoc_1_6), symimplementation, tmp1551, tmp1554)

		tmp1556 := PrimSet(symshen_4_dsigf_d, tmp1555)

		_ = tmp1556

		tmp1557 := MakeNative(func(__e *ControlFlow) {
			V6222 := __e.Get(1)
			_ = V6222
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6218 := __e.Get(1)
				_ = B6218
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6219 := __e.Get(1)
					_ = L6219
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6220 := __e.Get(1)
						_ = Key6220
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6221 := __e.Get(1)
							_ = C6221
							tmp1558 := PrimCons(symsymbol, Nil)

							tmp1559 := PrimCons(symlist, tmp1558)

							tmp1560 := PrimCons(symsymbol, Nil)

							tmp1561 := PrimCons(symlist, tmp1560)

							tmp1562 := PrimCons(tmp1561, Nil)

							tmp1563 := PrimCons(sym_1_1_6, tmp1562)

							tmp1564 := PrimCons(tmp1559, tmp1563)

							__e.TailApply(PrimFunc(symis_b), V6222, tmp1564, B6218, L6219, Key6220, C6221)
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

		tmp1565 := PrimValue(symshen_4_dsigf_d)

		tmp1566 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude, tmp1557, tmp1565)

		tmp1567 := PrimSet(symshen_4_dsigf_d, tmp1566)

		_ = tmp1567

		tmp1568 := MakeNative(func(__e *ControlFlow) {
			V6227 := __e.Get(1)
			_ = V6227
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6223 := __e.Get(1)
				_ = B6223
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6224 := __e.Get(1)
					_ = L6224
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6225 := __e.Get(1)
						_ = Key6225
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6226 := __e.Get(1)
							_ = C6226
							tmp1569 := PrimCons(symsymbol, Nil)

							tmp1570 := PrimCons(symlist, tmp1569)

							tmp1571 := PrimCons(symsymbol, Nil)

							tmp1572 := PrimCons(symlist, tmp1571)

							tmp1573 := PrimCons(tmp1572, Nil)

							tmp1574 := PrimCons(sym_1_1_6, tmp1573)

							tmp1575 := PrimCons(tmp1570, tmp1574)

							__e.TailApply(PrimFunc(symis_b), V6227, tmp1575, B6223, L6224, Key6225, C6226)
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

		tmp1576 := PrimValue(symshen_4_dsigf_d)

		tmp1577 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminclude_1all_1but, tmp1568, tmp1576)

		tmp1578 := PrimSet(symshen_4_dsigf_d, tmp1577)

		_ = tmp1578

		tmp1579 := MakeNative(func(__e *ControlFlow) {
			V6232 := __e.Get(1)
			_ = V6232
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6228 := __e.Get(1)
				_ = B6228
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6229 := __e.Get(1)
					_ = L6229
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6230 := __e.Get(1)
						_ = Key6230
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6231 := __e.Get(1)
							_ = C6231
							tmp1580 := PrimCons(symsymbol, Nil)

							tmp1581 := PrimCons(symlist, tmp1580)

							tmp1582 := PrimCons(tmp1581, Nil)

							tmp1583 := PrimCons(sym_1_1_6, tmp1582)

							__e.TailApply(PrimFunc(symis_b), V6232, tmp1583, B6228, L6229, Key6230, C6231)
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

		tmp1584 := PrimValue(symshen_4_dsigf_d)

		tmp1585 := Call(__e, PrimFunc(symshen_4assoc_1_6), symincluded, tmp1579, tmp1584)

		tmp1586 := PrimSet(symshen_4_dsigf_d, tmp1585)

		_ = tmp1586

		tmp1587 := MakeNative(func(__e *ControlFlow) {
			V6237 := __e.Get(1)
			_ = V6237
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6233 := __e.Get(1)
				_ = B6233
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6234 := __e.Get(1)
					_ = L6234
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6235 := __e.Get(1)
						_ = Key6235
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6236 := __e.Get(1)
							_ = C6236
							tmp1588 := PrimCons(symnumber, Nil)

							tmp1589 := PrimCons(sym_1_1_6, tmp1588)

							__e.TailApply(PrimFunc(symis_b), V6237, tmp1589, B6233, L6234, Key6235, C6236)
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

		tmp1590 := PrimValue(symshen_4_dsigf_d)

		tmp1591 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminferences, tmp1587, tmp1590)

		tmp1592 := PrimSet(symshen_4_dsigf_d, tmp1591)

		_ = tmp1592

		tmp1593 := MakeNative(func(__e *ControlFlow) {
			V6242 := __e.Get(1)
			_ = V6242
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6238 := __e.Get(1)
				_ = B6238
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6239 := __e.Get(1)
					_ = L6239
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6240 := __e.Get(1)
						_ = Key6240
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6241 := __e.Get(1)
							_ = C6241
							tmp1594 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1595 := PrimCons(symstring, Nil)

								tmp1596 := PrimCons(sym_1_1_6, tmp1595)

								tmp1597 := PrimCons(symstring, tmp1596)

								tmp1598 := PrimCons(tmp1597, Nil)

								tmp1599 := PrimCons(sym_1_1_6, tmp1598)

								tmp1600 := PrimCons(A, tmp1599)

								tmp1601 := Call(__e, PrimFunc(symis_b), V6242, tmp1600, B6238, L6239, Key6240, C6241)

								__e.TailApply(PrimFunc(symshen_4gc), B6238, tmp1601)
								return

							}, 1)

							tmp1602 := Call(__e, PrimFunc(symshen_4newpv), B6238)

							__e.TailApply(tmp1594, tmp1602)
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

		tmp1603 := PrimValue(symshen_4_dsigf_d)

		tmp1604 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4insert, tmp1593, tmp1603)

		tmp1605 := PrimSet(symshen_4_dsigf_d, tmp1604)

		_ = tmp1605

		tmp1606 := MakeNative(func(__e *ControlFlow) {
			V6247 := __e.Get(1)
			_ = V6247
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6243 := __e.Get(1)
				_ = B6243
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6244 := __e.Get(1)
					_ = L6244
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6245 := __e.Get(1)
						_ = Key6245
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6246 := __e.Get(1)
							_ = C6246
							tmp1607 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1608 := PrimCons(symboolean, Nil)

								tmp1609 := PrimCons(sym_1_1_6, tmp1608)

								tmp1610 := PrimCons(A, tmp1609)

								tmp1611 := Call(__e, PrimFunc(symis_b), V6247, tmp1610, B6243, L6244, Key6245, C6246)

								__e.TailApply(PrimFunc(symshen_4gc), B6243, tmp1611)
								return

							}, 1)

							tmp1612 := Call(__e, PrimFunc(symshen_4newpv), B6243)

							__e.TailApply(tmp1607, tmp1612)
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

		tmp1613 := PrimValue(symshen_4_dsigf_d)

		tmp1614 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminteger_2, tmp1606, tmp1613)

		tmp1615 := PrimSet(symshen_4_dsigf_d, tmp1614)

		_ = tmp1615

		tmp1616 := MakeNative(func(__e *ControlFlow) {
			V6252 := __e.Get(1)
			_ = V6252
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6248 := __e.Get(1)
				_ = B6248
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6249 := __e.Get(1)
					_ = L6249
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6250 := __e.Get(1)
						_ = Key6250
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6251 := __e.Get(1)
							_ = C6251
							tmp1617 := PrimCons(symsymbol, Nil)

							tmp1618 := PrimCons(symlist, tmp1617)

							tmp1619 := PrimCons(tmp1618, Nil)

							tmp1620 := PrimCons(sym_1_1_6, tmp1619)

							tmp1621 := PrimCons(symsymbol, tmp1620)

							__e.TailApply(PrimFunc(symis_b), V6252, tmp1621, B6248, L6249, Key6250, C6251)
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

		tmp1622 := PrimValue(symshen_4_dsigf_d)

		tmp1623 := Call(__e, PrimFunc(symshen_4assoc_1_6), syminternal, tmp1616, tmp1622)

		tmp1624 := PrimSet(symshen_4_dsigf_d, tmp1623)

		_ = tmp1624

		tmp1625 := MakeNative(func(__e *ControlFlow) {
			V6257 := __e.Get(1)
			_ = V6257
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6253 := __e.Get(1)
				_ = B6253
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6254 := __e.Get(1)
					_ = L6254
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6255 := __e.Get(1)
						_ = Key6255
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6256 := __e.Get(1)
							_ = C6256
							tmp1626 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1627 := PrimCons(A, Nil)

								tmp1628 := PrimCons(symlist, tmp1627)

								tmp1629 := PrimCons(A, Nil)

								tmp1630 := PrimCons(symlist, tmp1629)

								tmp1631 := PrimCons(A, Nil)

								tmp1632 := PrimCons(symlist, tmp1631)

								tmp1633 := PrimCons(tmp1632, Nil)

								tmp1634 := PrimCons(sym_1_1_6, tmp1633)

								tmp1635 := PrimCons(tmp1630, tmp1634)

								tmp1636 := PrimCons(tmp1635, Nil)

								tmp1637 := PrimCons(sym_1_1_6, tmp1636)

								tmp1638 := PrimCons(tmp1628, tmp1637)

								tmp1639 := Call(__e, PrimFunc(symis_b), V6257, tmp1638, B6253, L6254, Key6255, C6256)

								__e.TailApply(PrimFunc(symshen_4gc), B6253, tmp1639)
								return

							}, 1)

							tmp1640 := Call(__e, PrimFunc(symshen_4newpv), B6253)

							__e.TailApply(tmp1626, tmp1640)
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

		tmp1641 := PrimValue(symshen_4_dsigf_d)

		tmp1642 := Call(__e, PrimFunc(symshen_4assoc_1_6), symintersection, tmp1625, tmp1641)

		tmp1643 := PrimSet(symshen_4_dsigf_d, tmp1642)

		_ = tmp1643

		tmp1644 := MakeNative(func(__e *ControlFlow) {
			V6262 := __e.Get(1)
			_ = V6262
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6258 := __e.Get(1)
				_ = B6258
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6259 := __e.Get(1)
					_ = L6259
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6260 := __e.Get(1)
						_ = Key6260
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6261 := __e.Get(1)
							_ = C6261
							tmp1645 := PrimCons(symstring, Nil)

							tmp1646 := PrimCons(sym_1_1_6, tmp1645)

							__e.TailApply(PrimFunc(symis_b), V6262, tmp1646, B6258, L6259, Key6260, C6261)
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

		tmp1647 := PrimValue(symshen_4_dsigf_d)

		tmp1648 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlanguage, tmp1644, tmp1647)

		tmp1649 := PrimSet(symshen_4_dsigf_d, tmp1648)

		_ = tmp1649

		tmp1650 := MakeNative(func(__e *ControlFlow) {
			V6267 := __e.Get(1)
			_ = V6267
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6263 := __e.Get(1)
				_ = B6263
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6264 := __e.Get(1)
					_ = L6264
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6265 := __e.Get(1)
						_ = Key6265
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6266 := __e.Get(1)
							_ = C6266
							tmp1651 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1652 := PrimCons(A, Nil)

								tmp1653 := PrimCons(symlist, tmp1652)

								tmp1654 := PrimCons(symnumber, Nil)

								tmp1655 := PrimCons(sym_1_1_6, tmp1654)

								tmp1656 := PrimCons(tmp1653, tmp1655)

								tmp1657 := Call(__e, PrimFunc(symis_b), V6267, tmp1656, B6263, L6264, Key6265, C6266)

								__e.TailApply(PrimFunc(symshen_4gc), B6263, tmp1657)
								return

							}, 1)

							tmp1658 := Call(__e, PrimFunc(symshen_4newpv), B6263)

							__e.TailApply(tmp1651, tmp1658)
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

		tmp1659 := PrimValue(symshen_4_dsigf_d)

		tmp1660 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlength, tmp1650, tmp1659)

		tmp1661 := PrimSet(symshen_4_dsigf_d, tmp1660)

		_ = tmp1661

		tmp1662 := MakeNative(func(__e *ControlFlow) {
			V6272 := __e.Get(1)
			_ = V6272
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6268 := __e.Get(1)
				_ = B6268
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6269 := __e.Get(1)
					_ = L6269
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6270 := __e.Get(1)
						_ = Key6270
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6271 := __e.Get(1)
							_ = C6271
							tmp1663 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1664 := PrimCons(A, Nil)

								tmp1665 := PrimCons(symvector, tmp1664)

								tmp1666 := PrimCons(symnumber, Nil)

								tmp1667 := PrimCons(sym_1_1_6, tmp1666)

								tmp1668 := PrimCons(tmp1665, tmp1667)

								tmp1669 := Call(__e, PrimFunc(symis_b), V6272, tmp1668, B6268, L6269, Key6270, C6271)

								__e.TailApply(PrimFunc(symshen_4gc), B6268, tmp1669)
								return

							}, 1)

							tmp1670 := Call(__e, PrimFunc(symshen_4newpv), B6268)

							__e.TailApply(tmp1663, tmp1670)
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

		tmp1671 := PrimValue(symshen_4_dsigf_d)

		tmp1672 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlimit, tmp1662, tmp1671)

		tmp1673 := PrimSet(symshen_4_dsigf_d, tmp1672)

		_ = tmp1673

		tmp1674 := MakeNative(func(__e *ControlFlow) {
			V6277 := __e.Get(1)
			_ = V6277
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6273 := __e.Get(1)
				_ = B6273
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6274 := __e.Get(1)
					_ = L6274
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6275 := __e.Get(1)
						_ = Key6275
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6276 := __e.Get(1)
							_ = C6276
							tmp1675 := PrimCons(symin, Nil)

							tmp1676 := PrimCons(symstream, tmp1675)

							tmp1677 := PrimCons(symunit, Nil)

							tmp1678 := PrimCons(symlist, tmp1677)

							tmp1679 := PrimCons(tmp1678, Nil)

							tmp1680 := PrimCons(sym_1_1_6, tmp1679)

							tmp1681 := PrimCons(tmp1676, tmp1680)

							__e.TailApply(PrimFunc(symis_b), V6277, tmp1681, B6273, L6274, Key6275, C6276)
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

		tmp1682 := PrimValue(symshen_4_dsigf_d)

		tmp1683 := Call(__e, PrimFunc(symshen_4assoc_1_6), symlineread, tmp1674, tmp1682)

		tmp1684 := PrimSet(symshen_4_dsigf_d, tmp1683)

		_ = tmp1684

		tmp1685 := MakeNative(func(__e *ControlFlow) {
			V6282 := __e.Get(1)
			_ = V6282
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6278 := __e.Get(1)
				_ = B6278
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6279 := __e.Get(1)
					_ = L6279
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6280 := __e.Get(1)
						_ = Key6280
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6281 := __e.Get(1)
							_ = C6281
							tmp1686 := PrimCons(symsymbol, Nil)

							tmp1687 := PrimCons(sym_1_1_6, tmp1686)

							tmp1688 := PrimCons(symstring, tmp1687)

							__e.TailApply(PrimFunc(symis_b), V6282, tmp1688, B6278, L6279, Key6280, C6281)
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

		tmp1689 := PrimValue(symshen_4_dsigf_d)

		tmp1690 := Call(__e, PrimFunc(symshen_4assoc_1_6), symload, tmp1685, tmp1689)

		tmp1691 := PrimSet(symshen_4_dsigf_d, tmp1690)

		_ = tmp1691

		tmp1692 := MakeNative(func(__e *ControlFlow) {
			V6287 := __e.Get(1)
			_ = V6287
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6283 := __e.Get(1)
				_ = B6283
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6284 := __e.Get(1)
					_ = L6284
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6285 := __e.Get(1)
						_ = Key6285
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6286 := __e.Get(1)
							_ = C6286
							tmp1693 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1694 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1695 := PrimCons(B, Nil)

									tmp1696 := PrimCons(sym_1_1_6, tmp1695)

									tmp1697 := PrimCons(A, tmp1696)

									tmp1698 := PrimCons(A, Nil)

									tmp1699 := PrimCons(symlist, tmp1698)

									tmp1700 := PrimCons(B, Nil)

									tmp1701 := PrimCons(symlist, tmp1700)

									tmp1702 := PrimCons(tmp1701, Nil)

									tmp1703 := PrimCons(sym_1_1_6, tmp1702)

									tmp1704 := PrimCons(tmp1699, tmp1703)

									tmp1705 := PrimCons(tmp1704, Nil)

									tmp1706 := PrimCons(sym_1_1_6, tmp1705)

									tmp1707 := PrimCons(tmp1697, tmp1706)

									tmp1708 := Call(__e, PrimFunc(symis_b), V6287, tmp1707, B6283, L6284, Key6285, C6286)

									__e.TailApply(PrimFunc(symshen_4gc), B6283, tmp1708)
									return

								}, 1)

								tmp1709 := Call(__e, PrimFunc(symshen_4newpv), B6283)

								tmp1710 := Call(__e, tmp1694, tmp1709)

								__e.TailApply(PrimFunc(symshen_4gc), B6283, tmp1710)
								return

							}, 1)

							tmp1711 := Call(__e, PrimFunc(symshen_4newpv), B6283)

							__e.TailApply(tmp1693, tmp1711)
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

		tmp1712 := PrimValue(symshen_4_dsigf_d)

		tmp1713 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmap, tmp1692, tmp1712)

		tmp1714 := PrimSet(symshen_4_dsigf_d, tmp1713)

		_ = tmp1714

		tmp1715 := MakeNative(func(__e *ControlFlow) {
			V6292 := __e.Get(1)
			_ = V6292
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6288 := __e.Get(1)
				_ = B6288
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6289 := __e.Get(1)
					_ = L6289
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6290 := __e.Get(1)
						_ = Key6290
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6291 := __e.Get(1)
							_ = C6291
							tmp1716 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1717 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1718 := PrimCons(B, Nil)

									tmp1719 := PrimCons(symlist, tmp1718)

									tmp1720 := PrimCons(tmp1719, Nil)

									tmp1721 := PrimCons(sym_1_1_6, tmp1720)

									tmp1722 := PrimCons(A, tmp1721)

									tmp1723 := PrimCons(A, Nil)

									tmp1724 := PrimCons(symlist, tmp1723)

									tmp1725 := PrimCons(B, Nil)

									tmp1726 := PrimCons(symlist, tmp1725)

									tmp1727 := PrimCons(tmp1726, Nil)

									tmp1728 := PrimCons(sym_1_1_6, tmp1727)

									tmp1729 := PrimCons(tmp1724, tmp1728)

									tmp1730 := PrimCons(tmp1729, Nil)

									tmp1731 := PrimCons(sym_1_1_6, tmp1730)

									tmp1732 := PrimCons(tmp1722, tmp1731)

									tmp1733 := Call(__e, PrimFunc(symis_b), V6292, tmp1732, B6288, L6289, Key6290, C6291)

									__e.TailApply(PrimFunc(symshen_4gc), B6288, tmp1733)
									return

								}, 1)

								tmp1734 := Call(__e, PrimFunc(symshen_4newpv), B6288)

								tmp1735 := Call(__e, tmp1717, tmp1734)

								__e.TailApply(PrimFunc(symshen_4gc), B6288, tmp1735)
								return

							}, 1)

							tmp1736 := Call(__e, PrimFunc(symshen_4newpv), B6288)

							__e.TailApply(tmp1716, tmp1736)
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

		tmp1737 := PrimValue(symshen_4_dsigf_d)

		tmp1738 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmapcan, tmp1715, tmp1737)

		tmp1739 := PrimSet(symshen_4_dsigf_d, tmp1738)

		_ = tmp1739

		tmp1740 := MakeNative(func(__e *ControlFlow) {
			V6297 := __e.Get(1)
			_ = V6297
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6293 := __e.Get(1)
				_ = B6293
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6294 := __e.Get(1)
					_ = L6294
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6295 := __e.Get(1)
						_ = Key6295
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6296 := __e.Get(1)
							_ = C6296
							tmp1741 := PrimCons(symnumber, Nil)

							tmp1742 := PrimCons(sym_1_1_6, tmp1741)

							tmp1743 := PrimCons(symnumber, tmp1742)

							__e.TailApply(PrimFunc(symis_b), V6297, tmp1743, B6293, L6294, Key6295, C6296)
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

		tmp1744 := PrimValue(symshen_4_dsigf_d)

		tmp1745 := Call(__e, PrimFunc(symshen_4assoc_1_6), symmaxinferences, tmp1740, tmp1744)

		tmp1746 := PrimSet(symshen_4_dsigf_d, tmp1745)

		_ = tmp1746

		tmp1747 := MakeNative(func(__e *ControlFlow) {
			V6302 := __e.Get(1)
			_ = V6302
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6298 := __e.Get(1)
				_ = B6298
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6299 := __e.Get(1)
					_ = L6299
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6300 := __e.Get(1)
						_ = Key6300
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6301 := __e.Get(1)
							_ = C6301
							tmp1748 := PrimCons(symstring, Nil)

							tmp1749 := PrimCons(sym_1_1_6, tmp1748)

							tmp1750 := PrimCons(symnumber, tmp1749)

							__e.TailApply(PrimFunc(symis_b), V6302, tmp1750, B6298, L6299, Key6300, C6301)
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

		tmp1751 := PrimValue(symshen_4_dsigf_d)

		tmp1752 := Call(__e, PrimFunc(symshen_4assoc_1_6), symn_1_6string, tmp1747, tmp1751)

		tmp1753 := PrimSet(symshen_4_dsigf_d, tmp1752)

		_ = tmp1753

		tmp1754 := MakeNative(func(__e *ControlFlow) {
			V6307 := __e.Get(1)
			_ = V6307
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6303 := __e.Get(1)
				_ = B6303
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6304 := __e.Get(1)
					_ = L6304
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6305 := __e.Get(1)
						_ = Key6305
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6306 := __e.Get(1)
							_ = C6306
							tmp1755 := PrimCons(symnumber, Nil)

							tmp1756 := PrimCons(sym_1_1_6, tmp1755)

							tmp1757 := PrimCons(symnumber, tmp1756)

							__e.TailApply(PrimFunc(symis_b), V6307, tmp1757, B6303, L6304, Key6305, C6306)
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

		tmp1758 := PrimValue(symshen_4_dsigf_d)

		tmp1759 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnl, tmp1754, tmp1758)

		tmp1760 := PrimSet(symshen_4_dsigf_d, tmp1759)

		_ = tmp1760

		tmp1761 := MakeNative(func(__e *ControlFlow) {
			V6312 := __e.Get(1)
			_ = V6312
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6308 := __e.Get(1)
				_ = B6308
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6309 := __e.Get(1)
					_ = L6309
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6310 := __e.Get(1)
						_ = Key6310
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6311 := __e.Get(1)
							_ = C6311
							tmp1762 := PrimCons(symboolean, Nil)

							tmp1763 := PrimCons(sym_1_1_6, tmp1762)

							tmp1764 := PrimCons(symboolean, tmp1763)

							__e.TailApply(PrimFunc(symis_b), V6312, tmp1764, B6308, L6309, Key6310, C6311)
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

		tmp1765 := PrimValue(symshen_4_dsigf_d)

		tmp1766 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnot, tmp1761, tmp1765)

		tmp1767 := PrimSet(symshen_4_dsigf_d, tmp1766)

		_ = tmp1767

		tmp1768 := MakeNative(func(__e *ControlFlow) {
			V6317 := __e.Get(1)
			_ = V6317
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6313 := __e.Get(1)
				_ = B6313
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6314 := __e.Get(1)
					_ = L6314
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6315 := __e.Get(1)
						_ = Key6315
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6316 := __e.Get(1)
							_ = C6316
							tmp1769 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1770 := PrimCons(A, Nil)

								tmp1771 := PrimCons(symlist, tmp1770)

								tmp1772 := PrimCons(A, Nil)

								tmp1773 := PrimCons(sym_1_1_6, tmp1772)

								tmp1774 := PrimCons(tmp1771, tmp1773)

								tmp1775 := PrimCons(tmp1774, Nil)

								tmp1776 := PrimCons(sym_1_1_6, tmp1775)

								tmp1777 := PrimCons(symnumber, tmp1776)

								tmp1778 := Call(__e, PrimFunc(symis_b), V6317, tmp1777, B6313, L6314, Key6315, C6316)

								__e.TailApply(PrimFunc(symshen_4gc), B6313, tmp1778)
								return

							}, 1)

							tmp1779 := Call(__e, PrimFunc(symshen_4newpv), B6313)

							__e.TailApply(tmp1769, tmp1779)
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

		tmp1780 := PrimValue(symshen_4_dsigf_d)

		tmp1781 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnth, tmp1768, tmp1780)

		tmp1782 := PrimSet(symshen_4_dsigf_d, tmp1781)

		_ = tmp1782

		tmp1783 := MakeNative(func(__e *ControlFlow) {
			V6322 := __e.Get(1)
			_ = V6322
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6318 := __e.Get(1)
				_ = B6318
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6319 := __e.Get(1)
					_ = L6319
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6320 := __e.Get(1)
						_ = Key6320
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6321 := __e.Get(1)
							_ = C6321
							tmp1784 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1785 := PrimCons(symboolean, Nil)

								tmp1786 := PrimCons(sym_1_1_6, tmp1785)

								tmp1787 := PrimCons(A, tmp1786)

								tmp1788 := Call(__e, PrimFunc(symis_b), V6322, tmp1787, B6318, L6319, Key6320, C6321)

								__e.TailApply(PrimFunc(symshen_4gc), B6318, tmp1788)
								return

							}, 1)

							tmp1789 := Call(__e, PrimFunc(symshen_4newpv), B6318)

							__e.TailApply(tmp1784, tmp1789)
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

		tmp1790 := PrimValue(symshen_4_dsigf_d)

		tmp1791 := Call(__e, PrimFunc(symshen_4assoc_1_6), symnumber_2, tmp1783, tmp1790)

		tmp1792 := PrimSet(symshen_4_dsigf_d, tmp1791)

		_ = tmp1792

		tmp1793 := MakeNative(func(__e *ControlFlow) {
			V6327 := __e.Get(1)
			_ = V6327
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6323 := __e.Get(1)
				_ = B6323
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6324 := __e.Get(1)
					_ = L6324
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6325 := __e.Get(1)
						_ = Key6325
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6326 := __e.Get(1)
							_ = C6326
							tmp1794 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1795 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp1796 := PrimCons(symnumber, Nil)

									tmp1797 := PrimCons(sym_1_1_6, tmp1796)

									tmp1798 := PrimCons(B, tmp1797)

									tmp1799 := PrimCons(tmp1798, Nil)

									tmp1800 := PrimCons(sym_1_1_6, tmp1799)

									tmp1801 := PrimCons(A, tmp1800)

									tmp1802 := Call(__e, PrimFunc(symis_b), V6327, tmp1801, B6323, L6324, Key6325, C6326)

									__e.TailApply(PrimFunc(symshen_4gc), B6323, tmp1802)
									return

								}, 1)

								tmp1803 := Call(__e, PrimFunc(symshen_4newpv), B6323)

								tmp1804 := Call(__e, tmp1795, tmp1803)

								__e.TailApply(PrimFunc(symshen_4gc), B6323, tmp1804)
								return

							}, 1)

							tmp1805 := Call(__e, PrimFunc(symshen_4newpv), B6323)

							__e.TailApply(tmp1794, tmp1805)
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

		tmp1806 := PrimValue(symshen_4_dsigf_d)

		tmp1807 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurrences, tmp1793, tmp1806)

		tmp1808 := PrimSet(symshen_4_dsigf_d, tmp1807)

		_ = tmp1808

		tmp1809 := MakeNative(func(__e *ControlFlow) {
			V6332 := __e.Get(1)
			_ = V6332
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6328 := __e.Get(1)
				_ = B6328
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6329 := __e.Get(1)
					_ = L6329
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6330 := __e.Get(1)
						_ = Key6330
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6331 := __e.Get(1)
							_ = C6331
							tmp1810 := PrimCons(symboolean, Nil)

							tmp1811 := PrimCons(sym_1_1_6, tmp1810)

							tmp1812 := PrimCons(symsymbol, tmp1811)

							__e.TailApply(PrimFunc(symis_b), V6332, tmp1812, B6328, L6329, Key6330, C6331)
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

		tmp1813 := PrimValue(symshen_4_dsigf_d)

		tmp1814 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_1check, tmp1809, tmp1813)

		tmp1815 := PrimSet(symshen_4_dsigf_d, tmp1814)

		_ = tmp1815

		tmp1816 := MakeNative(func(__e *ControlFlow) {
			V6337 := __e.Get(1)
			_ = V6337
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6333 := __e.Get(1)
				_ = B6333
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6334 := __e.Get(1)
					_ = L6334
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6335 := __e.Get(1)
						_ = Key6335
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6336 := __e.Get(1)
							_ = C6336
							tmp1817 := PrimCons(symboolean, Nil)

							tmp1818 := PrimCons(sym_1_1_6, tmp1817)

							__e.TailApply(PrimFunc(symis_b), V6337, tmp1818, B6333, L6334, Key6335, C6336)
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

		tmp1819 := PrimValue(symshen_4_dsigf_d)

		tmp1820 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoccurs_2, tmp1816, tmp1819)

		tmp1821 := PrimSet(symshen_4_dsigf_d, tmp1820)

		_ = tmp1821

		tmp1822 := MakeNative(func(__e *ControlFlow) {
			V6342 := __e.Get(1)
			_ = V6342
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6338 := __e.Get(1)
				_ = B6338
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6339 := __e.Get(1)
					_ = L6339
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6340 := __e.Get(1)
						_ = Key6340
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6341 := __e.Get(1)
							_ = C6341
							tmp1823 := PrimCons(symboolean, Nil)

							tmp1824 := PrimCons(sym_1_1_6, tmp1823)

							tmp1825 := PrimCons(symsymbol, tmp1824)

							__e.TailApply(PrimFunc(symis_b), V6342, tmp1825, B6338, L6339, Key6340, C6341)
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

		tmp1826 := PrimValue(symshen_4_dsigf_d)

		tmp1827 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise, tmp1822, tmp1826)

		tmp1828 := PrimSet(symshen_4_dsigf_d, tmp1827)

		_ = tmp1828

		tmp1829 := MakeNative(func(__e *ControlFlow) {
			V6347 := __e.Get(1)
			_ = V6347
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6343 := __e.Get(1)
				_ = B6343
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6344 := __e.Get(1)
					_ = L6344
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6345 := __e.Get(1)
						_ = Key6345
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6346 := __e.Get(1)
							_ = C6346
							tmp1830 := PrimCons(symboolean, Nil)

							tmp1831 := PrimCons(sym_1_1_6, tmp1830)

							__e.TailApply(PrimFunc(symis_b), V6347, tmp1831, B6343, L6344, Key6345, C6346)
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

		tmp1832 := PrimValue(symshen_4_dsigf_d)

		tmp1833 := Call(__e, PrimFunc(symshen_4assoc_1_6), symoptimise_2, tmp1829, tmp1832)

		tmp1834 := PrimSet(symshen_4_dsigf_d, tmp1833)

		_ = tmp1834

		tmp1835 := MakeNative(func(__e *ControlFlow) {
			V6352 := __e.Get(1)
			_ = V6352
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6348 := __e.Get(1)
				_ = B6348
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6349 := __e.Get(1)
					_ = L6349
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6350 := __e.Get(1)
						_ = Key6350
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6351 := __e.Get(1)
							_ = C6351
							tmp1836 := PrimCons(symboolean, Nil)

							tmp1837 := PrimCons(sym_1_1_6, tmp1836)

							tmp1838 := PrimCons(symboolean, tmp1837)

							tmp1839 := PrimCons(tmp1838, Nil)

							tmp1840 := PrimCons(sym_1_1_6, tmp1839)

							tmp1841 := PrimCons(symboolean, tmp1840)

							__e.TailApply(PrimFunc(symis_b), V6352, tmp1841, B6348, L6349, Key6350, C6351)
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

		tmp1842 := PrimValue(symshen_4_dsigf_d)

		tmp1843 := Call(__e, PrimFunc(symshen_4assoc_1_6), symor, tmp1835, tmp1842)

		tmp1844 := PrimSet(symshen_4_dsigf_d, tmp1843)

		_ = tmp1844

		tmp1845 := MakeNative(func(__e *ControlFlow) {
			V6357 := __e.Get(1)
			_ = V6357
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6353 := __e.Get(1)
				_ = B6353
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6354 := __e.Get(1)
					_ = L6354
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6355 := __e.Get(1)
						_ = Key6355
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6356 := __e.Get(1)
							_ = C6356
							tmp1846 := PrimCons(symstring, Nil)

							tmp1847 := PrimCons(sym_1_1_6, tmp1846)

							__e.TailApply(PrimFunc(symis_b), V6357, tmp1847, B6353, L6354, Key6355, C6356)
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

		tmp1848 := PrimValue(symshen_4_dsigf_d)

		tmp1849 := Call(__e, PrimFunc(symshen_4assoc_1_6), symos, tmp1845, tmp1848)

		tmp1850 := PrimSet(symshen_4_dsigf_d, tmp1849)

		_ = tmp1850

		tmp1851 := MakeNative(func(__e *ControlFlow) {
			V6362 := __e.Get(1)
			_ = V6362
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6358 := __e.Get(1)
				_ = B6358
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6359 := __e.Get(1)
					_ = L6359
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6360 := __e.Get(1)
						_ = Key6360
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6361 := __e.Get(1)
							_ = C6361
							tmp1852 := PrimCons(symboolean, Nil)

							tmp1853 := PrimCons(sym_1_1_6, tmp1852)

							tmp1854 := PrimCons(symsymbol, tmp1853)

							__e.TailApply(PrimFunc(symis_b), V6362, tmp1854, B6358, L6359, Key6360, C6361)
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

		tmp1855 := PrimValue(symshen_4_dsigf_d)

		tmp1856 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympackage_2, tmp1851, tmp1855)

		tmp1857 := PrimSet(symshen_4_dsigf_d, tmp1856)

		_ = tmp1857

		tmp1858 := MakeNative(func(__e *ControlFlow) {
			V6367 := __e.Get(1)
			_ = V6367
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6363 := __e.Get(1)
				_ = B6363
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6364 := __e.Get(1)
					_ = L6364
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6365 := __e.Get(1)
						_ = Key6365
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6366 := __e.Get(1)
							_ = C6366
							tmp1859 := PrimCons(symstring, Nil)

							tmp1860 := PrimCons(sym_1_1_6, tmp1859)

							__e.TailApply(PrimFunc(symis_b), V6367, tmp1860, B6363, L6364, Key6365, C6366)
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

		tmp1861 := PrimValue(symshen_4_dsigf_d)

		tmp1862 := Call(__e, PrimFunc(symshen_4assoc_1_6), symport, tmp1858, tmp1861)

		tmp1863 := PrimSet(symshen_4_dsigf_d, tmp1862)

		_ = tmp1863

		tmp1864 := MakeNative(func(__e *ControlFlow) {
			V6372 := __e.Get(1)
			_ = V6372
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6368 := __e.Get(1)
				_ = B6368
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6369 := __e.Get(1)
					_ = L6369
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6370 := __e.Get(1)
						_ = Key6370
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6371 := __e.Get(1)
							_ = C6371
							tmp1865 := PrimCons(symstring, Nil)

							tmp1866 := PrimCons(sym_1_1_6, tmp1865)

							__e.TailApply(PrimFunc(symis_b), V6372, tmp1866, B6368, L6369, Key6370, C6371)
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

		tmp1867 := PrimValue(symshen_4_dsigf_d)

		tmp1868 := Call(__e, PrimFunc(symshen_4assoc_1_6), symporters, tmp1864, tmp1867)

		tmp1869 := PrimSet(symshen_4_dsigf_d, tmp1868)

		_ = tmp1869

		tmp1870 := MakeNative(func(__e *ControlFlow) {
			V6377 := __e.Get(1)
			_ = V6377
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6373 := __e.Get(1)
				_ = B6373
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6374 := __e.Get(1)
					_ = L6374
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6375 := __e.Get(1)
						_ = Key6375
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6376 := __e.Get(1)
							_ = C6376
							tmp1871 := PrimCons(symstring, Nil)

							tmp1872 := PrimCons(sym_1_1_6, tmp1871)

							tmp1873 := PrimCons(symnumber, tmp1872)

							tmp1874 := PrimCons(tmp1873, Nil)

							tmp1875 := PrimCons(sym_1_1_6, tmp1874)

							tmp1876 := PrimCons(symstring, tmp1875)

							__e.TailApply(PrimFunc(symis_b), V6377, tmp1876, B6373, L6374, Key6375, C6376)
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

		tmp1877 := PrimValue(symshen_4_dsigf_d)

		tmp1878 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympos, tmp1870, tmp1877)

		tmp1879 := PrimSet(symshen_4_dsigf_d, tmp1878)

		_ = tmp1879

		tmp1880 := MakeNative(func(__e *ControlFlow) {
			V6382 := __e.Get(1)
			_ = V6382
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6378 := __e.Get(1)
				_ = B6378
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6379 := __e.Get(1)
					_ = L6379
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6380 := __e.Get(1)
						_ = Key6380
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6381 := __e.Get(1)
							_ = C6381
							tmp1881 := PrimCons(symout, Nil)

							tmp1882 := PrimCons(symstream, tmp1881)

							tmp1883 := PrimCons(symstring, Nil)

							tmp1884 := PrimCons(sym_1_1_6, tmp1883)

							tmp1885 := PrimCons(tmp1882, tmp1884)

							tmp1886 := PrimCons(tmp1885, Nil)

							tmp1887 := PrimCons(sym_1_1_6, tmp1886)

							tmp1888 := PrimCons(symstring, tmp1887)

							__e.TailApply(PrimFunc(symis_b), V6382, tmp1888, B6378, L6379, Key6380, C6381)
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

		tmp1889 := PrimValue(symshen_4_dsigf_d)

		tmp1890 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympr, tmp1880, tmp1889)

		tmp1891 := PrimSet(symshen_4_dsigf_d, tmp1890)

		_ = tmp1891

		tmp1892 := MakeNative(func(__e *ControlFlow) {
			V6387 := __e.Get(1)
			_ = V6387
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6383 := __e.Get(1)
				_ = B6383
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6384 := __e.Get(1)
					_ = L6384
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6385 := __e.Get(1)
						_ = Key6385
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6386 := __e.Get(1)
							_ = C6386
							tmp1893 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1894 := PrimCons(A, Nil)

								tmp1895 := PrimCons(sym_1_1_6, tmp1894)

								tmp1896 := PrimCons(A, tmp1895)

								tmp1897 := Call(__e, PrimFunc(symis_b), V6387, tmp1896, B6383, L6384, Key6385, C6386)

								__e.TailApply(PrimFunc(symshen_4gc), B6383, tmp1897)
								return

							}, 1)

							tmp1898 := Call(__e, PrimFunc(symshen_4newpv), B6383)

							__e.TailApply(tmp1893, tmp1898)
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

		tmp1899 := PrimValue(symshen_4_dsigf_d)

		tmp1900 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprint, tmp1892, tmp1899)

		tmp1901 := PrimSet(symshen_4_dsigf_d, tmp1900)

		_ = tmp1901

		tmp1902 := MakeNative(func(__e *ControlFlow) {
			V6392 := __e.Get(1)
			_ = V6392
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6388 := __e.Get(1)
				_ = B6388
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6389 := __e.Get(1)
					_ = L6389
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6390 := __e.Get(1)
						_ = Key6390
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6391 := __e.Get(1)
							_ = C6391
							tmp1903 := PrimCons(symsymbol, Nil)

							tmp1904 := PrimCons(sym_1_1_6, tmp1903)

							tmp1905 := PrimCons(symsymbol, tmp1904)

							__e.TailApply(PrimFunc(symis_b), V6392, tmp1905, B6388, L6389, Key6390, C6391)
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

		tmp1906 := PrimValue(symshen_4_dsigf_d)

		tmp1907 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile, tmp1902, tmp1906)

		tmp1908 := PrimSet(symshen_4_dsigf_d, tmp1907)

		_ = tmp1908

		tmp1909 := MakeNative(func(__e *ControlFlow) {
			V6397 := __e.Get(1)
			_ = V6397
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6393 := __e.Get(1)
				_ = B6393
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6394 := __e.Get(1)
					_ = L6394
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6395 := __e.Get(1)
						_ = Key6395
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6396 := __e.Get(1)
							_ = C6396
							tmp1910 := PrimCons(symsymbol, Nil)

							tmp1911 := PrimCons(symlist, tmp1910)

							tmp1912 := PrimCons(symsymbol, Nil)

							tmp1913 := PrimCons(symlist, tmp1912)

							tmp1914 := PrimCons(tmp1913, Nil)

							tmp1915 := PrimCons(sym_1_1_6, tmp1914)

							tmp1916 := PrimCons(tmp1911, tmp1915)

							__e.TailApply(PrimFunc(symis_b), V6397, tmp1916, B6393, L6394, Key6395, C6396)
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

		tmp1917 := PrimValue(symshen_4_dsigf_d)

		tmp1918 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude, tmp1909, tmp1917)

		tmp1919 := PrimSet(symshen_4_dsigf_d, tmp1918)

		_ = tmp1919

		tmp1920 := MakeNative(func(__e *ControlFlow) {
			V6402 := __e.Get(1)
			_ = V6402
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6398 := __e.Get(1)
				_ = B6398
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6399 := __e.Get(1)
					_ = L6399
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6400 := __e.Get(1)
						_ = Key6400
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6401 := __e.Get(1)
							_ = C6401
							tmp1921 := PrimCons(symstring, Nil)

							tmp1922 := PrimCons(sym_1_1_6, tmp1921)

							tmp1923 := PrimCons(symstring, tmp1922)

							__e.TailApply(PrimFunc(symis_b), V6402, tmp1923, B6398, L6399, Key6400, C6401)
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

		tmp1924 := PrimValue(symshen_4_dsigf_d)

		tmp1925 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4proc_1nl, tmp1920, tmp1924)

		tmp1926 := PrimSet(symshen_4_dsigf_d, tmp1925)

		_ = tmp1926

		tmp1927 := MakeNative(func(__e *ControlFlow) {
			V6407 := __e.Get(1)
			_ = V6407
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6403 := __e.Get(1)
				_ = B6403
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6404 := __e.Get(1)
					_ = L6404
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6405 := __e.Get(1)
						_ = Key6405
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6406 := __e.Get(1)
							_ = C6406
							tmp1928 := PrimCons(symnumber, Nil)

							tmp1929 := PrimCons(sym_d, tmp1928)

							tmp1930 := PrimCons(symsymbol, tmp1929)

							tmp1931 := PrimCons(tmp1930, Nil)

							tmp1932 := PrimCons(sym_1_1_6, tmp1931)

							tmp1933 := PrimCons(symsymbol, tmp1932)

							__e.TailApply(PrimFunc(symis_b), V6407, tmp1933, B6403, L6404, Key6405, C6406)
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

		tmp1934 := PrimValue(symshen_4_dsigf_d)

		tmp1935 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprofile_1results, tmp1927, tmp1934)

		tmp1936 := PrimSet(symshen_4_dsigf_d, tmp1935)

		_ = tmp1936

		tmp1937 := MakeNative(func(__e *ControlFlow) {
			V6412 := __e.Get(1)
			_ = V6412
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6408 := __e.Get(1)
				_ = B6408
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6409 := __e.Get(1)
					_ = L6409
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6410 := __e.Get(1)
						_ = Key6410
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6411 := __e.Get(1)
							_ = C6411
							tmp1938 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp1939 := PrimCons(A, Nil)

								tmp1940 := PrimCons(sym_1_1_6, tmp1939)

								tmp1941 := PrimCons(A, tmp1940)

								tmp1942 := Call(__e, PrimFunc(symis_b), V6412, tmp1941, B6408, L6409, Key6410, C6411)

								__e.TailApply(PrimFunc(symshen_4gc), B6408, tmp1942)
								return

							}, 1)

							tmp1943 := Call(__e, PrimFunc(symshen_4newpv), B6408)

							__e.TailApply(tmp1938, tmp1943)
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

		tmp1944 := PrimValue(symshen_4_dsigf_d)

		tmp1945 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprotect, tmp1937, tmp1944)

		tmp1946 := PrimSet(symshen_4_dsigf_d, tmp1945)

		_ = tmp1946

		tmp1947 := MakeNative(func(__e *ControlFlow) {
			V6417 := __e.Get(1)
			_ = V6417
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6413 := __e.Get(1)
				_ = B6413
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6414 := __e.Get(1)
					_ = L6414
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6415 := __e.Get(1)
						_ = Key6415
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6416 := __e.Get(1)
							_ = C6416
							tmp1948 := PrimCons(symsymbol, Nil)

							tmp1949 := PrimCons(symlist, tmp1948)

							tmp1950 := PrimCons(symsymbol, Nil)

							tmp1951 := PrimCons(symlist, tmp1950)

							tmp1952 := PrimCons(tmp1951, Nil)

							tmp1953 := PrimCons(sym_1_1_6, tmp1952)

							tmp1954 := PrimCons(tmp1949, tmp1953)

							__e.TailApply(PrimFunc(symis_b), V6417, tmp1954, B6413, L6414, Key6415, C6416)
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

		tmp1955 := PrimValue(symshen_4_dsigf_d)

		tmp1956 := Call(__e, PrimFunc(symshen_4assoc_1_6), sympreclude_1all_1but, tmp1947, tmp1955)

		tmp1957 := PrimSet(symshen_4_dsigf_d, tmp1956)

		_ = tmp1957

		tmp1958 := MakeNative(func(__e *ControlFlow) {
			V6422 := __e.Get(1)
			_ = V6422
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6418 := __e.Get(1)
				_ = B6418
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6419 := __e.Get(1)
					_ = L6419
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6420 := __e.Get(1)
						_ = Key6420
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6421 := __e.Get(1)
							_ = C6421
							tmp1959 := PrimCons(symout, Nil)

							tmp1960 := PrimCons(symstream, tmp1959)

							tmp1961 := PrimCons(symstring, Nil)

							tmp1962 := PrimCons(sym_1_1_6, tmp1961)

							tmp1963 := PrimCons(tmp1960, tmp1962)

							tmp1964 := PrimCons(tmp1963, Nil)

							tmp1965 := PrimCons(sym_1_1_6, tmp1964)

							tmp1966 := PrimCons(symstring, tmp1965)

							__e.TailApply(PrimFunc(symis_b), V6422, tmp1966, B6418, L6419, Key6420, C6421)
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

		tmp1967 := PrimValue(symshen_4_dsigf_d)

		tmp1968 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4prhush, tmp1958, tmp1967)

		tmp1969 := PrimSet(symshen_4_dsigf_d, tmp1968)

		_ = tmp1969

		tmp1970 := MakeNative(func(__e *ControlFlow) {
			V6427 := __e.Get(1)
			_ = V6427
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6423 := __e.Get(1)
				_ = B6423
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6424 := __e.Get(1)
					_ = L6424
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6425 := __e.Get(1)
						_ = Key6425
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6426 := __e.Get(1)
							_ = C6426
							tmp1971 := PrimCons(symnumber, Nil)

							tmp1972 := PrimCons(sym_1_1_6, tmp1971)

							tmp1973 := PrimCons(symnumber, tmp1972)

							__e.TailApply(PrimFunc(symis_b), V6427, tmp1973, B6423, L6424, Key6425, C6426)
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

		tmp1974 := PrimValue(symshen_4_dsigf_d)

		tmp1975 := Call(__e, PrimFunc(symshen_4assoc_1_6), symprolog_1memory, tmp1970, tmp1974)

		tmp1976 := PrimSet(symshen_4_dsigf_d, tmp1975)

		_ = tmp1976

		tmp1977 := MakeNative(func(__e *ControlFlow) {
			V6432 := __e.Get(1)
			_ = V6432
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6428 := __e.Get(1)
				_ = B6428
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6429 := __e.Get(1)
					_ = L6429
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6430 := __e.Get(1)
						_ = Key6430
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6431 := __e.Get(1)
							_ = C6431
							tmp1978 := PrimCons(symunit, Nil)

							tmp1979 := PrimCons(symlist, tmp1978)

							tmp1980 := PrimCons(tmp1979, Nil)

							tmp1981 := PrimCons(sym_1_1_6, tmp1980)

							tmp1982 := PrimCons(symsymbol, tmp1981)

							__e.TailApply(PrimFunc(symis_b), V6432, tmp1982, B6428, L6429, Key6430, C6431)
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

		tmp1983 := PrimValue(symshen_4_dsigf_d)

		tmp1984 := Call(__e, PrimFunc(symshen_4assoc_1_6), symps, tmp1977, tmp1983)

		tmp1985 := PrimSet(symshen_4_dsigf_d, tmp1984)

		_ = tmp1985

		tmp1986 := MakeNative(func(__e *ControlFlow) {
			V6437 := __e.Get(1)
			_ = V6437
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6433 := __e.Get(1)
				_ = B6433
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6434 := __e.Get(1)
					_ = L6434
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6435 := __e.Get(1)
						_ = Key6435
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6436 := __e.Get(1)
							_ = C6436
							tmp1987 := PrimCons(symin, Nil)

							tmp1988 := PrimCons(symstream, tmp1987)

							tmp1989 := PrimCons(symunit, Nil)

							tmp1990 := PrimCons(sym_1_1_6, tmp1989)

							tmp1991 := PrimCons(tmp1988, tmp1990)

							__e.TailApply(PrimFunc(symis_b), V6437, tmp1991, B6433, L6434, Key6435, C6436)
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

		tmp1992 := PrimValue(symshen_4_dsigf_d)

		tmp1993 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread, tmp1986, tmp1992)

		tmp1994 := PrimSet(symshen_4_dsigf_d, tmp1993)

		_ = tmp1994

		tmp1995 := MakeNative(func(__e *ControlFlow) {
			V6442 := __e.Get(1)
			_ = V6442
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6438 := __e.Get(1)
				_ = B6438
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6439 := __e.Get(1)
					_ = L6439
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6440 := __e.Get(1)
						_ = Key6440
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6441 := __e.Get(1)
							_ = C6441
							tmp1996 := PrimCons(symin, Nil)

							tmp1997 := PrimCons(symstream, tmp1996)

							tmp1998 := PrimCons(symnumber, Nil)

							tmp1999 := PrimCons(sym_1_1_6, tmp1998)

							tmp2000 := PrimCons(tmp1997, tmp1999)

							__e.TailApply(PrimFunc(symis_b), V6442, tmp2000, B6438, L6439, Key6440, C6441)
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

		tmp2001 := PrimValue(symshen_4_dsigf_d)

		tmp2002 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1byte, tmp1995, tmp2001)

		tmp2003 := PrimSet(symshen_4_dsigf_d, tmp2002)

		_ = tmp2003

		tmp2004 := MakeNative(func(__e *ControlFlow) {
			V6447 := __e.Get(1)
			_ = V6447
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6443 := __e.Get(1)
				_ = B6443
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6444 := __e.Get(1)
					_ = L6444
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6445 := __e.Get(1)
						_ = Key6445
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6446 := __e.Get(1)
							_ = C6446
							tmp2005 := PrimCons(symnumber, Nil)

							tmp2006 := PrimCons(symlist, tmp2005)

							tmp2007 := PrimCons(tmp2006, Nil)

							tmp2008 := PrimCons(sym_1_1_6, tmp2007)

							tmp2009 := PrimCons(symstring, tmp2008)

							__e.TailApply(PrimFunc(symis_b), V6447, tmp2009, B6443, L6444, Key6445, C6446)
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

		tmp2010 := PrimValue(symshen_4_dsigf_d)

		tmp2011 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1bytelist, tmp2004, tmp2010)

		tmp2012 := PrimSet(symshen_4_dsigf_d, tmp2011)

		_ = tmp2012

		tmp2013 := MakeNative(func(__e *ControlFlow) {
			V6452 := __e.Get(1)
			_ = V6452
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6448 := __e.Get(1)
				_ = B6448
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6449 := __e.Get(1)
					_ = L6449
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6450 := __e.Get(1)
						_ = Key6450
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6451 := __e.Get(1)
							_ = C6451
							tmp2014 := PrimCons(symstring, Nil)

							tmp2015 := PrimCons(sym_1_1_6, tmp2014)

							tmp2016 := PrimCons(symstring, tmp2015)

							__e.TailApply(PrimFunc(symis_b), V6452, tmp2016, B6448, L6449, Key6450, C6451)
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

		tmp2017 := PrimValue(symshen_4_dsigf_d)

		tmp2018 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file_1as_1string, tmp2013, tmp2017)

		tmp2019 := PrimSet(symshen_4_dsigf_d, tmp2018)

		_ = tmp2019

		tmp2020 := MakeNative(func(__e *ControlFlow) {
			V6457 := __e.Get(1)
			_ = V6457
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6453 := __e.Get(1)
				_ = B6453
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6454 := __e.Get(1)
					_ = L6454
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6455 := __e.Get(1)
						_ = Key6455
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6456 := __e.Get(1)
							_ = C6456
							tmp2021 := PrimCons(symunit, Nil)

							tmp2022 := PrimCons(symlist, tmp2021)

							tmp2023 := PrimCons(tmp2022, Nil)

							tmp2024 := PrimCons(sym_1_1_6, tmp2023)

							tmp2025 := PrimCons(symstring, tmp2024)

							__e.TailApply(PrimFunc(symis_b), V6457, tmp2025, B6453, L6454, Key6455, C6456)
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

		tmp2026 := PrimValue(symshen_4_dsigf_d)

		tmp2027 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1file, tmp2020, tmp2026)

		tmp2028 := PrimSet(symshen_4_dsigf_d, tmp2027)

		_ = tmp2028

		tmp2029 := MakeNative(func(__e *ControlFlow) {
			V6462 := __e.Get(1)
			_ = V6462
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6458 := __e.Get(1)
				_ = B6458
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6459 := __e.Get(1)
					_ = L6459
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6460 := __e.Get(1)
						_ = Key6460
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6461 := __e.Get(1)
							_ = C6461
							tmp2030 := PrimCons(symunit, Nil)

							tmp2031 := PrimCons(symlist, tmp2030)

							tmp2032 := PrimCons(tmp2031, Nil)

							tmp2033 := PrimCons(sym_1_1_6, tmp2032)

							tmp2034 := PrimCons(symstring, tmp2033)

							__e.TailApply(PrimFunc(symis_b), V6462, tmp2034, B6458, L6459, Key6460, C6461)
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

		tmp2035 := PrimValue(symshen_4_dsigf_d)

		tmp2036 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string, tmp2029, tmp2035)

		tmp2037 := PrimSet(symshen_4_dsigf_d, tmp2036)

		_ = tmp2037

		tmp2038 := MakeNative(func(__e *ControlFlow) {
			V6467 := __e.Get(1)
			_ = V6467
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6463 := __e.Get(1)
				_ = B6463
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6464 := __e.Get(1)
					_ = L6464
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6465 := __e.Get(1)
						_ = Key6465
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6466 := __e.Get(1)
							_ = C6466
							tmp2039 := PrimCons(symunit, Nil)

							tmp2040 := PrimCons(symlist, tmp2039)

							tmp2041 := PrimCons(tmp2040, Nil)

							tmp2042 := PrimCons(sym_1_1_6, tmp2041)

							tmp2043 := PrimCons(symstring, tmp2042)

							__e.TailApply(PrimFunc(symis_b), V6467, tmp2043, B6463, L6464, Key6465, C6466)
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

		tmp2044 := PrimValue(symshen_4_dsigf_d)

		tmp2045 := Call(__e, PrimFunc(symshen_4assoc_1_6), symread_1from_1string_1unprocessed, tmp2038, tmp2044)

		tmp2046 := PrimSet(symshen_4_dsigf_d, tmp2045)

		_ = tmp2046

		tmp2047 := MakeNative(func(__e *ControlFlow) {
			V6472 := __e.Get(1)
			_ = V6472
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6468 := __e.Get(1)
				_ = B6468
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6469 := __e.Get(1)
					_ = L6469
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6470 := __e.Get(1)
						_ = Key6470
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6471 := __e.Get(1)
							_ = C6471
							tmp2048 := PrimCons(symstring, Nil)

							tmp2049 := PrimCons(sym_1_1_6, tmp2048)

							__e.TailApply(PrimFunc(symis_b), V6472, tmp2049, B6468, L6469, Key6470, C6471)
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

		tmp2050 := PrimValue(symshen_4_dsigf_d)

		tmp2051 := Call(__e, PrimFunc(symshen_4assoc_1_6), symrelease, tmp2047, tmp2050)

		tmp2052 := PrimSet(symshen_4_dsigf_d, tmp2051)

		_ = tmp2052

		tmp2053 := MakeNative(func(__e *ControlFlow) {
			V6477 := __e.Get(1)
			_ = V6477
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6473 := __e.Get(1)
				_ = B6473
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6474 := __e.Get(1)
					_ = L6474
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6475 := __e.Get(1)
						_ = Key6475
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6476 := __e.Get(1)
							_ = C6476
							tmp2054 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2055 := PrimCons(A, Nil)

								tmp2056 := PrimCons(symlist, tmp2055)

								tmp2057 := PrimCons(A, Nil)

								tmp2058 := PrimCons(symlist, tmp2057)

								tmp2059 := PrimCons(tmp2058, Nil)

								tmp2060 := PrimCons(sym_1_1_6, tmp2059)

								tmp2061 := PrimCons(tmp2056, tmp2060)

								tmp2062 := PrimCons(tmp2061, Nil)

								tmp2063 := PrimCons(sym_1_1_6, tmp2062)

								tmp2064 := PrimCons(A, tmp2063)

								tmp2065 := Call(__e, PrimFunc(symis_b), V6477, tmp2064, B6473, L6474, Key6475, C6476)

								__e.TailApply(PrimFunc(symshen_4gc), B6473, tmp2065)
								return

							}, 1)

							tmp2066 := Call(__e, PrimFunc(symshen_4newpv), B6473)

							__e.TailApply(tmp2054, tmp2066)
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

		tmp2067 := PrimValue(symshen_4_dsigf_d)

		tmp2068 := Call(__e, PrimFunc(symshen_4assoc_1_6), symremove, tmp2053, tmp2067)

		tmp2069 := PrimSet(symshen_4_dsigf_d, tmp2068)

		_ = tmp2069

		tmp2070 := MakeNative(func(__e *ControlFlow) {
			V6482 := __e.Get(1)
			_ = V6482
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6478 := __e.Get(1)
				_ = B6478
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6479 := __e.Get(1)
					_ = L6479
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6480 := __e.Get(1)
						_ = Key6480
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6481 := __e.Get(1)
							_ = C6481
							tmp2071 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2072 := PrimCons(A, Nil)

								tmp2073 := PrimCons(symlist, tmp2072)

								tmp2074 := PrimCons(A, Nil)

								tmp2075 := PrimCons(symlist, tmp2074)

								tmp2076 := PrimCons(tmp2075, Nil)

								tmp2077 := PrimCons(sym_1_1_6, tmp2076)

								tmp2078 := PrimCons(tmp2073, tmp2077)

								tmp2079 := Call(__e, PrimFunc(symis_b), V6482, tmp2078, B6478, L6479, Key6480, C6481)

								__e.TailApply(PrimFunc(symshen_4gc), B6478, tmp2079)
								return

							}, 1)

							tmp2080 := Call(__e, PrimFunc(symshen_4newpv), B6478)

							__e.TailApply(tmp2071, tmp2080)
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

		tmp2081 := PrimValue(symshen_4_dsigf_d)

		tmp2082 := Call(__e, PrimFunc(symshen_4assoc_1_6), symreverse, tmp2070, tmp2081)

		tmp2083 := PrimSet(symshen_4_dsigf_d, tmp2082)

		_ = tmp2083

		tmp2084 := MakeNative(func(__e *ControlFlow) {
			V6487 := __e.Get(1)
			_ = V6487
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6483 := __e.Get(1)
				_ = B6483
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6484 := __e.Get(1)
					_ = L6484
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6485 := __e.Get(1)
						_ = Key6485
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6486 := __e.Get(1)
							_ = C6486
							tmp2085 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2086 := PrimCons(A, Nil)

								tmp2087 := PrimCons(sym_1_1_6, tmp2086)

								tmp2088 := PrimCons(symstring, tmp2087)

								tmp2089 := Call(__e, PrimFunc(symis_b), V6487, tmp2088, B6483, L6484, Key6485, C6486)

								__e.TailApply(PrimFunc(symshen_4gc), B6483, tmp2089)
								return

							}, 1)

							tmp2090 := Call(__e, PrimFunc(symshen_4newpv), B6483)

							__e.TailApply(tmp2085, tmp2090)
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

		tmp2091 := PrimValue(symshen_4_dsigf_d)

		tmp2092 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsimple_1error, tmp2084, tmp2091)

		tmp2093 := PrimSet(symshen_4_dsigf_d, tmp2092)

		_ = tmp2093

		tmp2094 := MakeNative(func(__e *ControlFlow) {
			V6492 := __e.Get(1)
			_ = V6492
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6488 := __e.Get(1)
				_ = B6488
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6489 := __e.Get(1)
					_ = L6489
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6490 := __e.Get(1)
						_ = Key6490
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6491 := __e.Get(1)
							_ = C6491
							tmp2095 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2096 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp2097 := PrimCons(B, Nil)

									tmp2098 := PrimCons(sym_d, tmp2097)

									tmp2099 := PrimCons(A, tmp2098)

									tmp2100 := PrimCons(B, Nil)

									tmp2101 := PrimCons(sym_1_1_6, tmp2100)

									tmp2102 := PrimCons(tmp2099, tmp2101)

									tmp2103 := Call(__e, PrimFunc(symis_b), V6492, tmp2102, B6488, L6489, Key6490, C6491)

									__e.TailApply(PrimFunc(symshen_4gc), B6488, tmp2103)
									return

								}, 1)

								tmp2104 := Call(__e, PrimFunc(symshen_4newpv), B6488)

								tmp2105 := Call(__e, tmp2096, tmp2104)

								__e.TailApply(PrimFunc(symshen_4gc), B6488, tmp2105)
								return

							}, 1)

							tmp2106 := Call(__e, PrimFunc(symshen_4newpv), B6488)

							__e.TailApply(tmp2095, tmp2106)
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

		tmp2107 := PrimValue(symshen_4_dsigf_d)

		tmp2108 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsnd, tmp2094, tmp2107)

		tmp2109 := PrimSet(symshen_4_dsigf_d, tmp2108)

		_ = tmp2109

		tmp2110 := MakeNative(func(__e *ControlFlow) {
			V6497 := __e.Get(1)
			_ = V6497
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6493 := __e.Get(1)
				_ = B6493
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6494 := __e.Get(1)
					_ = L6494
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6495 := __e.Get(1)
						_ = Key6495
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6496 := __e.Get(1)
							_ = C6496
							tmp2111 := PrimCons(symsymbol, Nil)

							tmp2112 := PrimCons(sym_1_1_6, tmp2111)

							tmp2113 := PrimCons(symnumber, tmp2112)

							tmp2114 := PrimCons(tmp2113, Nil)

							tmp2115 := PrimCons(sym_1_1_6, tmp2114)

							tmp2116 := PrimCons(symsymbol, tmp2115)

							__e.TailApply(PrimFunc(symis_b), V6497, tmp2116, B6493, L6494, Key6495, C6496)
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

		tmp2117 := PrimValue(symshen_4_dsigf_d)

		tmp2118 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspecialise, tmp2110, tmp2117)

		tmp2119 := PrimSet(symshen_4_dsigf_d, tmp2118)

		_ = tmp2119

		tmp2120 := MakeNative(func(__e *ControlFlow) {
			V6502 := __e.Get(1)
			_ = V6502
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6498 := __e.Get(1)
				_ = B6498
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6499 := __e.Get(1)
					_ = L6499
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6500 := __e.Get(1)
						_ = Key6500
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6501 := __e.Get(1)
							_ = C6501
							tmp2121 := PrimCons(symboolean, Nil)

							tmp2122 := PrimCons(sym_1_1_6, tmp2121)

							tmp2123 := PrimCons(symsymbol, tmp2122)

							__e.TailApply(PrimFunc(symis_b), V6502, tmp2123, B6498, L6499, Key6500, C6501)
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

		tmp2124 := PrimValue(symshen_4_dsigf_d)

		tmp2125 := Call(__e, PrimFunc(symshen_4assoc_1_6), symspy, tmp2120, tmp2124)

		tmp2126 := PrimSet(symshen_4_dsigf_d, tmp2125)

		_ = tmp2126

		tmp2127 := MakeNative(func(__e *ControlFlow) {
			V6507 := __e.Get(1)
			_ = V6507
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6503 := __e.Get(1)
				_ = B6503
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6504 := __e.Get(1)
					_ = L6504
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6505 := __e.Get(1)
						_ = Key6505
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6506 := __e.Get(1)
							_ = C6506
							tmp2128 := PrimCons(symboolean, Nil)

							tmp2129 := PrimCons(sym_1_1_6, tmp2128)

							__e.TailApply(PrimFunc(symis_b), V6507, tmp2129, B6503, L6504, Key6505, C6506)
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

		tmp2130 := PrimValue(symshen_4_dsigf_d)

		tmp2131 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4spy_2, tmp2127, tmp2130)

		tmp2132 := PrimSet(symshen_4_dsigf_d, tmp2131)

		_ = tmp2132

		tmp2133 := MakeNative(func(__e *ControlFlow) {
			V6512 := __e.Get(1)
			_ = V6512
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6508 := __e.Get(1)
				_ = B6508
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6509 := __e.Get(1)
					_ = L6509
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6510 := __e.Get(1)
						_ = Key6510
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6511 := __e.Get(1)
							_ = C6511
							tmp2134 := PrimCons(symboolean, Nil)

							tmp2135 := PrimCons(sym_1_1_6, tmp2134)

							tmp2136 := PrimCons(symsymbol, tmp2135)

							__e.TailApply(PrimFunc(symis_b), V6512, tmp2136, B6508, L6509, Key6510, C6511)
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

		tmp2137 := PrimValue(symshen_4_dsigf_d)

		tmp2138 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstep, tmp2133, tmp2137)

		tmp2139 := PrimSet(symshen_4_dsigf_d, tmp2138)

		_ = tmp2139

		tmp2140 := MakeNative(func(__e *ControlFlow) {
			V6517 := __e.Get(1)
			_ = V6517
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6513 := __e.Get(1)
				_ = B6513
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6514 := __e.Get(1)
					_ = L6514
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6515 := __e.Get(1)
						_ = Key6515
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6516 := __e.Get(1)
							_ = C6516
							tmp2141 := PrimCons(symboolean, Nil)

							tmp2142 := PrimCons(sym_1_1_6, tmp2141)

							__e.TailApply(PrimFunc(symis_b), V6517, tmp2142, B6513, L6514, Key6515, C6516)
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

		tmp2143 := PrimValue(symshen_4_dsigf_d)

		tmp2144 := Call(__e, PrimFunc(symshen_4assoc_1_6), symshen_4step_2, tmp2140, tmp2143)

		tmp2145 := PrimSet(symshen_4_dsigf_d, tmp2144)

		_ = tmp2145

		tmp2146 := MakeNative(func(__e *ControlFlow) {
			V6522 := __e.Get(1)
			_ = V6522
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6518 := __e.Get(1)
				_ = B6518
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6519 := __e.Get(1)
					_ = L6519
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6520 := __e.Get(1)
						_ = Key6520
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6521 := __e.Get(1)
							_ = C6521
							tmp2147 := PrimCons(symin, Nil)

							tmp2148 := PrimCons(symstream, tmp2147)

							tmp2149 := PrimCons(tmp2148, Nil)

							tmp2150 := PrimCons(sym_1_1_6, tmp2149)

							__e.TailApply(PrimFunc(symis_b), V6522, tmp2150, B6518, L6519, Key6520, C6521)
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

		tmp2151 := PrimValue(symshen_4_dsigf_d)

		tmp2152 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstinput, tmp2146, tmp2151)

		tmp2153 := PrimSet(symshen_4_dsigf_d, tmp2152)

		_ = tmp2153

		tmp2154 := MakeNative(func(__e *ControlFlow) {
			V6527 := __e.Get(1)
			_ = V6527
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6523 := __e.Get(1)
				_ = B6523
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6524 := __e.Get(1)
					_ = L6524
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6525 := __e.Get(1)
						_ = Key6525
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6526 := __e.Get(1)
							_ = C6526
							tmp2155 := PrimCons(symout, Nil)

							tmp2156 := PrimCons(symstream, tmp2155)

							tmp2157 := PrimCons(tmp2156, Nil)

							tmp2158 := PrimCons(sym_1_1_6, tmp2157)

							__e.TailApply(PrimFunc(symis_b), V6527, tmp2158, B6523, L6524, Key6525, C6526)
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

		tmp2159 := PrimValue(symshen_4_dsigf_d)

		tmp2160 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsterror, tmp2154, tmp2159)

		tmp2161 := PrimSet(symshen_4_dsigf_d, tmp2160)

		_ = tmp2161

		tmp2162 := MakeNative(func(__e *ControlFlow) {
			V6532 := __e.Get(1)
			_ = V6532
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6528 := __e.Get(1)
				_ = B6528
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6529 := __e.Get(1)
					_ = L6529
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6530 := __e.Get(1)
						_ = Key6530
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6531 := __e.Get(1)
							_ = C6531
							tmp2163 := PrimCons(symout, Nil)

							tmp2164 := PrimCons(symstream, tmp2163)

							tmp2165 := PrimCons(tmp2164, Nil)

							tmp2166 := PrimCons(sym_1_1_6, tmp2165)

							__e.TailApply(PrimFunc(symis_b), V6532, tmp2166, B6528, L6529, Key6530, C6531)
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

		tmp2167 := PrimValue(symshen_4_dsigf_d)

		tmp2168 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstoutput, tmp2162, tmp2167)

		tmp2169 := PrimSet(symshen_4_dsigf_d, tmp2168)

		_ = tmp2169

		tmp2170 := MakeNative(func(__e *ControlFlow) {
			V6537 := __e.Get(1)
			_ = V6537
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6533 := __e.Get(1)
				_ = B6533
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6534 := __e.Get(1)
					_ = L6534
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6535 := __e.Get(1)
						_ = Key6535
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6536 := __e.Get(1)
							_ = C6536
							tmp2171 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2172 := PrimCons(symboolean, Nil)

								tmp2173 := PrimCons(sym_1_1_6, tmp2172)

								tmp2174 := PrimCons(A, tmp2173)

								tmp2175 := Call(__e, PrimFunc(symis_b), V6537, tmp2174, B6533, L6534, Key6535, C6536)

								__e.TailApply(PrimFunc(symshen_4gc), B6533, tmp2175)
								return

							}, 1)

							tmp2176 := Call(__e, PrimFunc(symshen_4newpv), B6533)

							__e.TailApply(tmp2171, tmp2176)
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

		tmp2177 := PrimValue(symshen_4_dsigf_d)

		tmp2178 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_2, tmp2170, tmp2177)

		tmp2179 := PrimSet(symshen_4_dsigf_d, tmp2178)

		_ = tmp2179

		tmp2180 := MakeNative(func(__e *ControlFlow) {
			V6542 := __e.Get(1)
			_ = V6542
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6538 := __e.Get(1)
				_ = B6538
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6539 := __e.Get(1)
					_ = L6539
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6540 := __e.Get(1)
						_ = Key6540
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6541 := __e.Get(1)
							_ = C6541
							tmp2181 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2182 := PrimCons(symstring, Nil)

								tmp2183 := PrimCons(sym_1_1_6, tmp2182)

								tmp2184 := PrimCons(A, tmp2183)

								tmp2185 := Call(__e, PrimFunc(symis_b), V6542, tmp2184, B6538, L6539, Key6540, C6541)

								__e.TailApply(PrimFunc(symshen_4gc), B6538, tmp2185)
								return

							}, 1)

							tmp2186 := Call(__e, PrimFunc(symshen_4newpv), B6538)

							__e.TailApply(tmp2181, tmp2186)
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

		tmp2187 := PrimValue(symshen_4_dsigf_d)

		tmp2188 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstr, tmp2180, tmp2187)

		tmp2189 := PrimSet(symshen_4_dsigf_d, tmp2188)

		_ = tmp2189

		tmp2190 := MakeNative(func(__e *ControlFlow) {
			V6547 := __e.Get(1)
			_ = V6547
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6543 := __e.Get(1)
				_ = B6543
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6544 := __e.Get(1)
					_ = L6544
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6545 := __e.Get(1)
						_ = Key6545
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6546 := __e.Get(1)
							_ = C6546
							tmp2191 := PrimCons(symnumber, Nil)

							tmp2192 := PrimCons(sym_1_1_6, tmp2191)

							tmp2193 := PrimCons(symstring, tmp2192)

							__e.TailApply(PrimFunc(symis_b), V6547, tmp2193, B6543, L6544, Key6545, C6546)
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

		tmp2194 := PrimValue(symshen_4_dsigf_d)

		tmp2195 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6n, tmp2190, tmp2194)

		tmp2196 := PrimSet(symshen_4_dsigf_d, tmp2195)

		_ = tmp2196

		tmp2197 := MakeNative(func(__e *ControlFlow) {
			V6552 := __e.Get(1)
			_ = V6552
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6548 := __e.Get(1)
				_ = B6548
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6549 := __e.Get(1)
					_ = L6549
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6550 := __e.Get(1)
						_ = Key6550
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6551 := __e.Get(1)
							_ = C6551
							tmp2198 := PrimCons(symsymbol, Nil)

							tmp2199 := PrimCons(sym_1_1_6, tmp2198)

							tmp2200 := PrimCons(symstring, tmp2199)

							__e.TailApply(PrimFunc(symis_b), V6552, tmp2200, B6548, L6549, Key6550, C6551)
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

		tmp2201 := PrimValue(symshen_4_dsigf_d)

		tmp2202 := Call(__e, PrimFunc(symshen_4assoc_1_6), symstring_1_6symbol, tmp2197, tmp2201)

		tmp2203 := PrimSet(symshen_4_dsigf_d, tmp2202)

		_ = tmp2203

		tmp2204 := MakeNative(func(__e *ControlFlow) {
			V6557 := __e.Get(1)
			_ = V6557
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6553 := __e.Get(1)
				_ = B6553
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6554 := __e.Get(1)
					_ = L6554
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6555 := __e.Get(1)
						_ = Key6555
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6556 := __e.Get(1)
							_ = C6556
							tmp2205 := PrimCons(symnumber, Nil)

							tmp2206 := PrimCons(symlist, tmp2205)

							tmp2207 := PrimCons(symnumber, Nil)

							tmp2208 := PrimCons(sym_1_1_6, tmp2207)

							tmp2209 := PrimCons(tmp2206, tmp2208)

							__e.TailApply(PrimFunc(symis_b), V6557, tmp2209, B6553, L6554, Key6555, C6556)
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

		tmp2210 := PrimValue(symshen_4_dsigf_d)

		tmp2211 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsum, tmp2204, tmp2210)

		tmp2212 := PrimSet(symshen_4_dsigf_d, tmp2211)

		_ = tmp2212

		tmp2213 := MakeNative(func(__e *ControlFlow) {
			V6562 := __e.Get(1)
			_ = V6562
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6558 := __e.Get(1)
				_ = B6558
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6559 := __e.Get(1)
					_ = L6559
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6560 := __e.Get(1)
						_ = Key6560
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6561 := __e.Get(1)
							_ = C6561
							tmp2214 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2215 := PrimCons(symboolean, Nil)

								tmp2216 := PrimCons(sym_1_1_6, tmp2215)

								tmp2217 := PrimCons(A, tmp2216)

								tmp2218 := Call(__e, PrimFunc(symis_b), V6562, tmp2217, B6558, L6559, Key6560, C6561)

								__e.TailApply(PrimFunc(symshen_4gc), B6558, tmp2218)
								return

							}, 1)

							tmp2219 := Call(__e, PrimFunc(symshen_4newpv), B6558)

							__e.TailApply(tmp2214, tmp2219)
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

		tmp2220 := PrimValue(symshen_4_dsigf_d)

		tmp2221 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsymbol_2, tmp2213, tmp2220)

		tmp2222 := PrimSet(symshen_4_dsigf_d, tmp2221)

		_ = tmp2222

		tmp2223 := MakeNative(func(__e *ControlFlow) {
			V6567 := __e.Get(1)
			_ = V6567
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6563 := __e.Get(1)
				_ = B6563
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6564 := __e.Get(1)
					_ = L6564
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6565 := __e.Get(1)
						_ = Key6565
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6566 := __e.Get(1)
							_ = C6566
							tmp2224 := PrimCons(symsymbol, Nil)

							tmp2225 := PrimCons(sym_1_1_6, tmp2224)

							tmp2226 := PrimCons(symsymbol, tmp2225)

							__e.TailApply(PrimFunc(symis_b), V6567, tmp2226, B6563, L6564, Key6565, C6566)
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

		tmp2227 := PrimValue(symshen_4_dsigf_d)

		tmp2228 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystemf, tmp2223, tmp2227)

		tmp2229 := PrimSet(symshen_4_dsigf_d, tmp2228)

		_ = tmp2229

		tmp2230 := MakeNative(func(__e *ControlFlow) {
			V6572 := __e.Get(1)
			_ = V6572
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6568 := __e.Get(1)
				_ = B6568
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6569 := __e.Get(1)
					_ = L6569
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6570 := __e.Get(1)
						_ = Key6570
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6571 := __e.Get(1)
							_ = C6571
							tmp2231 := PrimCons(symboolean, Nil)

							tmp2232 := PrimCons(sym_1_1_6, tmp2231)

							__e.TailApply(PrimFunc(symis_b), V6572, tmp2232, B6568, L6569, Key6570, C6571)
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

		tmp2233 := PrimValue(symshen_4_dsigf_d)

		tmp2234 := Call(__e, PrimFunc(symshen_4assoc_1_6), symsystem_1S_2, tmp2230, tmp2233)

		tmp2235 := PrimSet(symshen_4_dsigf_d, tmp2234)

		_ = tmp2235

		tmp2236 := MakeNative(func(__e *ControlFlow) {
			V6577 := __e.Get(1)
			_ = V6577
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6573 := __e.Get(1)
				_ = B6573
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6574 := __e.Get(1)
					_ = L6574
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6575 := __e.Get(1)
						_ = Key6575
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6576 := __e.Get(1)
							_ = C6576
							tmp2237 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2238 := PrimCons(A, Nil)

								tmp2239 := PrimCons(symlist, tmp2238)

								tmp2240 := PrimCons(A, Nil)

								tmp2241 := PrimCons(symlist, tmp2240)

								tmp2242 := PrimCons(tmp2241, Nil)

								tmp2243 := PrimCons(sym_1_1_6, tmp2242)

								tmp2244 := PrimCons(tmp2239, tmp2243)

								tmp2245 := Call(__e, PrimFunc(symis_b), V6577, tmp2244, B6573, L6574, Key6575, C6576)

								__e.TailApply(PrimFunc(symshen_4gc), B6573, tmp2245)
								return

							}, 1)

							tmp2246 := Call(__e, PrimFunc(symshen_4newpv), B6573)

							__e.TailApply(tmp2237, tmp2246)
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

		tmp2247 := PrimValue(symshen_4_dsigf_d)

		tmp2248 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtail, tmp2236, tmp2247)

		tmp2249 := PrimSet(symshen_4_dsigf_d, tmp2248)

		_ = tmp2249

		tmp2250 := MakeNative(func(__e *ControlFlow) {
			V6582 := __e.Get(1)
			_ = V6582
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6578 := __e.Get(1)
				_ = B6578
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6579 := __e.Get(1)
					_ = L6579
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6580 := __e.Get(1)
						_ = Key6580
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6581 := __e.Get(1)
							_ = C6581
							tmp2251 := PrimCons(symstring, Nil)

							tmp2252 := PrimCons(sym_1_1_6, tmp2251)

							tmp2253 := PrimCons(symstring, tmp2252)

							__e.TailApply(PrimFunc(symis_b), V6582, tmp2253, B6578, L6579, Key6580, C6581)
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

		tmp2254 := PrimValue(symshen_4_dsigf_d)

		tmp2255 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlstr, tmp2250, tmp2254)

		tmp2256 := PrimSet(symshen_4_dsigf_d, tmp2255)

		_ = tmp2256

		tmp2257 := MakeNative(func(__e *ControlFlow) {
			V6587 := __e.Get(1)
			_ = V6587
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6583 := __e.Get(1)
				_ = B6583
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6584 := __e.Get(1)
					_ = L6584
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6585 := __e.Get(1)
						_ = Key6585
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6586 := __e.Get(1)
							_ = C6586
							tmp2258 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2259 := PrimCons(A, Nil)

								tmp2260 := PrimCons(symvector, tmp2259)

								tmp2261 := PrimCons(A, Nil)

								tmp2262 := PrimCons(symvector, tmp2261)

								tmp2263 := PrimCons(tmp2262, Nil)

								tmp2264 := PrimCons(sym_1_1_6, tmp2263)

								tmp2265 := PrimCons(tmp2260, tmp2264)

								tmp2266 := Call(__e, PrimFunc(symis_b), V6587, tmp2265, B6583, L6584, Key6585, C6586)

								__e.TailApply(PrimFunc(symshen_4gc), B6583, tmp2266)
								return

							}, 1)

							tmp2267 := Call(__e, PrimFunc(symshen_4newpv), B6583)

							__e.TailApply(tmp2258, tmp2267)
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

		tmp2268 := PrimValue(symshen_4_dsigf_d)

		tmp2269 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtlv, tmp2257, tmp2268)

		tmp2270 := PrimSet(symshen_4_dsigf_d, tmp2269)

		_ = tmp2270

		tmp2271 := MakeNative(func(__e *ControlFlow) {
			V6592 := __e.Get(1)
			_ = V6592
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6588 := __e.Get(1)
				_ = B6588
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6589 := __e.Get(1)
					_ = L6589
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6590 := __e.Get(1)
						_ = Key6590
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6591 := __e.Get(1)
							_ = C6591
							tmp2272 := PrimCons(symboolean, Nil)

							tmp2273 := PrimCons(sym_1_1_6, tmp2272)

							tmp2274 := PrimCons(symsymbol, tmp2273)

							__e.TailApply(PrimFunc(symis_b), V6592, tmp2274, B6588, L6589, Key6590, C6591)
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

		tmp2275 := PrimValue(symshen_4_dsigf_d)

		tmp2276 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc, tmp2271, tmp2275)

		tmp2277 := PrimSet(symshen_4_dsigf_d, tmp2276)

		_ = tmp2277

		tmp2278 := MakeNative(func(__e *ControlFlow) {
			V6597 := __e.Get(1)
			_ = V6597
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6593 := __e.Get(1)
				_ = B6593
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6594 := __e.Get(1)
					_ = L6594
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6595 := __e.Get(1)
						_ = Key6595
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6596 := __e.Get(1)
							_ = C6596
							tmp2279 := PrimCons(symboolean, Nil)

							tmp2280 := PrimCons(sym_1_1_6, tmp2279)

							__e.TailApply(PrimFunc(symis_b), V6597, tmp2280, B6593, L6594, Key6595, C6596)
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

		tmp2281 := PrimValue(symshen_4_dsigf_d)

		tmp2282 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtc_2, tmp2278, tmp2281)

		tmp2283 := PrimSet(symshen_4_dsigf_d, tmp2282)

		_ = tmp2283

		tmp2284 := MakeNative(func(__e *ControlFlow) {
			V6602 := __e.Get(1)
			_ = V6602
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6598 := __e.Get(1)
				_ = B6598
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6599 := __e.Get(1)
					_ = L6599
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6600 := __e.Get(1)
						_ = Key6600
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6601 := __e.Get(1)
							_ = C6601
							tmp2285 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2286 := PrimCons(A, Nil)

								tmp2287 := PrimCons(symlazy, tmp2286)

								tmp2288 := PrimCons(A, Nil)

								tmp2289 := PrimCons(sym_1_1_6, tmp2288)

								tmp2290 := PrimCons(tmp2287, tmp2289)

								tmp2291 := Call(__e, PrimFunc(symis_b), V6602, tmp2290, B6598, L6599, Key6600, C6601)

								__e.TailApply(PrimFunc(symshen_4gc), B6598, tmp2291)
								return

							}, 1)

							tmp2292 := Call(__e, PrimFunc(symshen_4newpv), B6598)

							__e.TailApply(tmp2285, tmp2292)
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

		tmp2293 := PrimValue(symshen_4_dsigf_d)

		tmp2294 := Call(__e, PrimFunc(symshen_4assoc_1_6), symthaw, tmp2284, tmp2293)

		tmp2295 := PrimSet(symshen_4_dsigf_d, tmp2294)

		_ = tmp2295

		tmp2296 := MakeNative(func(__e *ControlFlow) {
			V6607 := __e.Get(1)
			_ = V6607
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6603 := __e.Get(1)
				_ = B6603
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6604 := __e.Get(1)
					_ = L6604
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6605 := __e.Get(1)
						_ = Key6605
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6606 := __e.Get(1)
							_ = C6606
							tmp2297 := PrimCons(symsymbol, Nil)

							tmp2298 := PrimCons(sym_1_1_6, tmp2297)

							tmp2299 := PrimCons(symsymbol, tmp2298)

							__e.TailApply(PrimFunc(symis_b), V6607, tmp2299, B6603, L6604, Key6605, C6606)
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

		tmp2300 := PrimValue(symshen_4_dsigf_d)

		tmp2301 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrack, tmp2296, tmp2300)

		tmp2302 := PrimSet(symshen_4_dsigf_d, tmp2301)

		_ = tmp2302

		tmp2303 := MakeNative(func(__e *ControlFlow) {
			V6612 := __e.Get(1)
			_ = V6612
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6608 := __e.Get(1)
				_ = B6608
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6609 := __e.Get(1)
					_ = L6609
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6610 := __e.Get(1)
						_ = Key6610
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6611 := __e.Get(1)
							_ = C6611
							tmp2304 := PrimCons(symsymbol, Nil)

							tmp2305 := PrimCons(symlist, tmp2304)

							tmp2306 := PrimCons(tmp2305, Nil)

							tmp2307 := PrimCons(sym_1_1_6, tmp2306)

							__e.TailApply(PrimFunc(symis_b), V6612, tmp2307, B6608, L6609, Key6610, C6611)
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

		tmp2308 := PrimValue(symshen_4_dsigf_d)

		tmp2309 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtracked, tmp2303, tmp2308)

		tmp2310 := PrimSet(symshen_4_dsigf_d, tmp2309)

		_ = tmp2310

		tmp2311 := MakeNative(func(__e *ControlFlow) {
			V6617 := __e.Get(1)
			_ = V6617
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6613 := __e.Get(1)
				_ = B6613
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6614 := __e.Get(1)
					_ = L6614
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6615 := __e.Get(1)
						_ = Key6615
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6616 := __e.Get(1)
							_ = C6616
							tmp2312 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2313 := PrimCons(A, Nil)

								tmp2314 := PrimCons(sym_1_1_6, tmp2313)

								tmp2315 := PrimCons(symexception, tmp2314)

								tmp2316 := PrimCons(A, Nil)

								tmp2317 := PrimCons(sym_1_1_6, tmp2316)

								tmp2318 := PrimCons(tmp2315, tmp2317)

								tmp2319 := PrimCons(tmp2318, Nil)

								tmp2320 := PrimCons(sym_1_1_6, tmp2319)

								tmp2321 := PrimCons(A, tmp2320)

								tmp2322 := Call(__e, PrimFunc(symis_b), V6617, tmp2321, B6613, L6614, Key6615, C6616)

								__e.TailApply(PrimFunc(symshen_4gc), B6613, tmp2322)
								return

							}, 1)

							tmp2323 := Call(__e, PrimFunc(symshen_4newpv), B6613)

							__e.TailApply(tmp2312, tmp2323)
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

		tmp2324 := PrimValue(symshen_4_dsigf_d)

		tmp2325 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtrap_1error, tmp2311, tmp2324)

		tmp2326 := PrimSet(symshen_4_dsigf_d, tmp2325)

		_ = tmp2326

		tmp2327 := MakeNative(func(__e *ControlFlow) {
			V6622 := __e.Get(1)
			_ = V6622
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6618 := __e.Get(1)
				_ = B6618
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6619 := __e.Get(1)
					_ = L6619
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6620 := __e.Get(1)
						_ = Key6620
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6621 := __e.Get(1)
							_ = C6621
							tmp2328 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2329 := PrimCons(symboolean, Nil)

								tmp2330 := PrimCons(sym_1_1_6, tmp2329)

								tmp2331 := PrimCons(A, tmp2330)

								tmp2332 := Call(__e, PrimFunc(symis_b), V6622, tmp2331, B6618, L6619, Key6620, C6621)

								__e.TailApply(PrimFunc(symshen_4gc), B6618, tmp2332)
								return

							}, 1)

							tmp2333 := Call(__e, PrimFunc(symshen_4newpv), B6618)

							__e.TailApply(tmp2328, tmp2333)
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

		tmp2334 := PrimValue(symshen_4_dsigf_d)

		tmp2335 := Call(__e, PrimFunc(symshen_4assoc_1_6), symtuple_2, tmp2327, tmp2334)

		tmp2336 := PrimSet(symshen_4_dsigf_d, tmp2335)

		_ = tmp2336

		tmp2337 := MakeNative(func(__e *ControlFlow) {
			V6627 := __e.Get(1)
			_ = V6627
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6623 := __e.Get(1)
				_ = B6623
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6624 := __e.Get(1)
					_ = L6624
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6625 := __e.Get(1)
						_ = Key6625
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6626 := __e.Get(1)
							_ = C6626
							tmp2338 := PrimCons(symstring, Nil)

							tmp2339 := PrimCons(symlist, tmp2338)

							tmp2340 := PrimCons(tmp2339, Nil)

							tmp2341 := PrimCons(sym_1_1_6, tmp2340)

							tmp2342 := PrimCons(symstring, tmp2341)

							__e.TailApply(PrimFunc(symis_b), V6627, tmp2342, B6623, L6624, Key6625, C6626)
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

		tmp2343 := PrimValue(symshen_4_dsigf_d)

		tmp2344 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunabsolute, tmp2337, tmp2343)

		tmp2345 := PrimSet(symshen_4_dsigf_d, tmp2344)

		_ = tmp2345

		tmp2346 := MakeNative(func(__e *ControlFlow) {
			V6632 := __e.Get(1)
			_ = V6632
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6628 := __e.Get(1)
				_ = B6628
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6629 := __e.Get(1)
					_ = L6629
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6630 := __e.Get(1)
						_ = Key6630
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6631 := __e.Get(1)
							_ = C6631
							tmp2347 := PrimCons(symsymbol, Nil)

							tmp2348 := PrimCons(sym_1_1_6, tmp2347)

							tmp2349 := PrimCons(symsymbol, tmp2348)

							__e.TailApply(PrimFunc(symis_b), V6632, tmp2349, B6628, L6629, Key6630, C6631)
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

		tmp2350 := PrimValue(symshen_4_dsigf_d)

		tmp2351 := Call(__e, PrimFunc(symshen_4assoc_1_6), symundefmacro, tmp2346, tmp2350)

		tmp2352 := PrimSet(symshen_4_dsigf_d, tmp2351)

		_ = tmp2352

		tmp2353 := MakeNative(func(__e *ControlFlow) {
			V6637 := __e.Get(1)
			_ = V6637
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6633 := __e.Get(1)
				_ = B6633
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6634 := __e.Get(1)
					_ = L6634
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6635 := __e.Get(1)
						_ = Key6635
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6636 := __e.Get(1)
							_ = C6636
							tmp2354 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2355 := PrimCons(A, Nil)

								tmp2356 := PrimCons(symlist, tmp2355)

								tmp2357 := PrimCons(A, Nil)

								tmp2358 := PrimCons(symlist, tmp2357)

								tmp2359 := PrimCons(A, Nil)

								tmp2360 := PrimCons(symlist, tmp2359)

								tmp2361 := PrimCons(tmp2360, Nil)

								tmp2362 := PrimCons(sym_1_1_6, tmp2361)

								tmp2363 := PrimCons(tmp2358, tmp2362)

								tmp2364 := PrimCons(tmp2363, Nil)

								tmp2365 := PrimCons(sym_1_1_6, tmp2364)

								tmp2366 := PrimCons(tmp2356, tmp2365)

								tmp2367 := Call(__e, PrimFunc(symis_b), V6637, tmp2366, B6633, L6634, Key6635, C6636)

								__e.TailApply(PrimFunc(symshen_4gc), B6633, tmp2367)
								return

							}, 1)

							tmp2368 := Call(__e, PrimFunc(symshen_4newpv), B6633)

							__e.TailApply(tmp2354, tmp2368)
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

		tmp2369 := PrimValue(symshen_4_dsigf_d)

		tmp2370 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunion, tmp2353, tmp2369)

		tmp2371 := PrimSet(symshen_4_dsigf_d, tmp2370)

		_ = tmp2371

		tmp2372 := MakeNative(func(__e *ControlFlow) {
			V6642 := __e.Get(1)
			_ = V6642
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6638 := __e.Get(1)
				_ = B6638
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6639 := __e.Get(1)
					_ = L6639
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6640 := __e.Get(1)
						_ = Key6640
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6641 := __e.Get(1)
							_ = C6641
							tmp2373 := PrimCons(symsymbol, Nil)

							tmp2374 := PrimCons(sym_1_1_6, tmp2373)

							tmp2375 := PrimCons(symsymbol, tmp2374)

							__e.TailApply(PrimFunc(symis_b), V6642, tmp2375, B6638, L6639, Key6640, C6641)
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

		tmp2376 := PrimValue(symshen_4_dsigf_d)

		tmp2377 := Call(__e, PrimFunc(symshen_4assoc_1_6), symunprofile, tmp2372, tmp2376)

		tmp2378 := PrimSet(symshen_4_dsigf_d, tmp2377)

		_ = tmp2378

		tmp2379 := MakeNative(func(__e *ControlFlow) {
			V6647 := __e.Get(1)
			_ = V6647
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6643 := __e.Get(1)
				_ = B6643
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6644 := __e.Get(1)
					_ = L6644
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6645 := __e.Get(1)
						_ = Key6645
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6646 := __e.Get(1)
							_ = C6646
							tmp2380 := PrimCons(symsymbol, Nil)

							tmp2381 := PrimCons(sym_1_1_6, tmp2380)

							tmp2382 := PrimCons(symsymbol, tmp2381)

							__e.TailApply(PrimFunc(symis_b), V6647, tmp2382, B6643, L6644, Key6645, C6646)
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

		tmp2383 := PrimValue(symshen_4_dsigf_d)

		tmp2384 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuntrack, tmp2379, tmp2383)

		tmp2385 := PrimSet(symshen_4_dsigf_d, tmp2384)

		_ = tmp2385

		tmp2386 := MakeNative(func(__e *ControlFlow) {
			V6652 := __e.Get(1)
			_ = V6652
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6648 := __e.Get(1)
				_ = B6648
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6649 := __e.Get(1)
					_ = L6649
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6650 := __e.Get(1)
						_ = Key6650
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6651 := __e.Get(1)
							_ = C6651
							tmp2387 := PrimCons(symsymbol, Nil)

							tmp2388 := PrimCons(symlist, tmp2387)

							tmp2389 := PrimCons(tmp2388, Nil)

							tmp2390 := PrimCons(sym_1_1_6, tmp2389)

							__e.TailApply(PrimFunc(symis_b), V6652, tmp2390, B6648, L6649, Key6650, C6651)
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

		tmp2391 := PrimValue(symshen_4_dsigf_d)

		tmp2392 := Call(__e, PrimFunc(symshen_4assoc_1_6), symuserdefs, tmp2386, tmp2391)

		tmp2393 := PrimSet(symshen_4_dsigf_d, tmp2392)

		_ = tmp2393

		tmp2394 := MakeNative(func(__e *ControlFlow) {
			V6657 := __e.Get(1)
			_ = V6657
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6653 := __e.Get(1)
				_ = B6653
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6654 := __e.Get(1)
					_ = L6654
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6655 := __e.Get(1)
						_ = Key6655
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6656 := __e.Get(1)
							_ = C6656
							tmp2395 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2396 := PrimCons(symboolean, Nil)

								tmp2397 := PrimCons(sym_1_1_6, tmp2396)

								tmp2398 := PrimCons(A, tmp2397)

								tmp2399 := Call(__e, PrimFunc(symis_b), V6657, tmp2398, B6653, L6654, Key6655, C6656)

								__e.TailApply(PrimFunc(symshen_4gc), B6653, tmp2399)
								return

							}, 1)

							tmp2400 := Call(__e, PrimFunc(symshen_4newpv), B6653)

							__e.TailApply(tmp2395, tmp2400)
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

		tmp2401 := PrimValue(symshen_4_dsigf_d)

		tmp2402 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvariable_2, tmp2394, tmp2401)

		tmp2403 := PrimSet(symshen_4_dsigf_d, tmp2402)

		_ = tmp2403

		tmp2404 := MakeNative(func(__e *ControlFlow) {
			V6662 := __e.Get(1)
			_ = V6662
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6658 := __e.Get(1)
				_ = B6658
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6659 := __e.Get(1)
					_ = L6659
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6660 := __e.Get(1)
						_ = Key6660
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6661 := __e.Get(1)
							_ = C6661
							tmp2405 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2406 := PrimCons(symboolean, Nil)

								tmp2407 := PrimCons(sym_1_1_6, tmp2406)

								tmp2408 := PrimCons(A, tmp2407)

								tmp2409 := Call(__e, PrimFunc(symis_b), V6662, tmp2408, B6658, L6659, Key6660, C6661)

								__e.TailApply(PrimFunc(symshen_4gc), B6658, tmp2409)
								return

							}, 1)

							tmp2410 := Call(__e, PrimFunc(symshen_4newpv), B6658)

							__e.TailApply(tmp2405, tmp2410)
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

		tmp2411 := PrimValue(symshen_4_dsigf_d)

		tmp2412 := Call(__e, PrimFunc(symshen_4assoc_1_6), symvector_2, tmp2404, tmp2411)

		tmp2413 := PrimSet(symshen_4_dsigf_d, tmp2412)

		_ = tmp2413

		tmp2414 := MakeNative(func(__e *ControlFlow) {
			V6667 := __e.Get(1)
			_ = V6667
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6663 := __e.Get(1)
				_ = B6663
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6664 := __e.Get(1)
					_ = L6664
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6665 := __e.Get(1)
						_ = Key6665
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6666 := __e.Get(1)
							_ = C6666
							tmp2415 := PrimCons(symstring, Nil)

							tmp2416 := PrimCons(sym_1_1_6, tmp2415)

							__e.TailApply(PrimFunc(symis_b), V6667, tmp2416, B6663, L6664, Key6665, C6666)
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

		tmp2417 := PrimValue(symshen_4_dsigf_d)

		tmp2418 := Call(__e, PrimFunc(symshen_4assoc_1_6), symversion, tmp2414, tmp2417)

		tmp2419 := PrimSet(symshen_4_dsigf_d, tmp2418)

		_ = tmp2419

		tmp2420 := MakeNative(func(__e *ControlFlow) {
			V6672 := __e.Get(1)
			_ = V6672
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6668 := __e.Get(1)
				_ = B6668
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6669 := __e.Get(1)
					_ = L6669
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6670 := __e.Get(1)
						_ = Key6670
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6671 := __e.Get(1)
							_ = C6671
							tmp2421 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2422 := PrimCons(A, Nil)

								tmp2423 := PrimCons(sym_1_1_6, tmp2422)

								tmp2424 := PrimCons(A, tmp2423)

								tmp2425 := PrimCons(tmp2424, Nil)

								tmp2426 := PrimCons(sym_1_1_6, tmp2425)

								tmp2427 := PrimCons(symstring, tmp2426)

								tmp2428 := Call(__e, PrimFunc(symis_b), V6672, tmp2427, B6668, L6669, Key6670, C6671)

								__e.TailApply(PrimFunc(symshen_4gc), B6668, tmp2428)
								return

							}, 1)

							tmp2429 := Call(__e, PrimFunc(symshen_4newpv), B6668)

							__e.TailApply(tmp2421, tmp2429)
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

		tmp2430 := PrimValue(symshen_4_dsigf_d)

		tmp2431 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1to_1file, tmp2420, tmp2430)

		tmp2432 := PrimSet(symshen_4_dsigf_d, tmp2431)

		_ = tmp2432

		tmp2433 := MakeNative(func(__e *ControlFlow) {
			V6677 := __e.Get(1)
			_ = V6677
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6673 := __e.Get(1)
				_ = B6673
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6674 := __e.Get(1)
					_ = L6674
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6675 := __e.Get(1)
						_ = Key6675
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6676 := __e.Get(1)
							_ = C6676
							tmp2434 := PrimCons(symout, Nil)

							tmp2435 := PrimCons(symstream, tmp2434)

							tmp2436 := PrimCons(symnumber, Nil)

							tmp2437 := PrimCons(sym_1_1_6, tmp2436)

							tmp2438 := PrimCons(tmp2435, tmp2437)

							tmp2439 := PrimCons(tmp2438, Nil)

							tmp2440 := PrimCons(sym_1_1_6, tmp2439)

							tmp2441 := PrimCons(symnumber, tmp2440)

							__e.TailApply(PrimFunc(symis_b), V6677, tmp2441, B6673, L6674, Key6675, C6676)
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

		tmp2442 := PrimValue(symshen_4_dsigf_d)

		tmp2443 := Call(__e, PrimFunc(symshen_4assoc_1_6), symwrite_1byte, tmp2433, tmp2442)

		tmp2444 := PrimSet(symshen_4_dsigf_d, tmp2443)

		_ = tmp2444

		tmp2445 := MakeNative(func(__e *ControlFlow) {
			V6682 := __e.Get(1)
			_ = V6682
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6678 := __e.Get(1)
				_ = B6678
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6679 := __e.Get(1)
					_ = L6679
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6680 := __e.Get(1)
						_ = Key6680
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6681 := __e.Get(1)
							_ = C6681
							tmp2446 := PrimCons(symboolean, Nil)

							tmp2447 := PrimCons(sym_1_1_6, tmp2446)

							tmp2448 := PrimCons(symstring, tmp2447)

							__e.TailApply(PrimFunc(symis_b), V6682, tmp2448, B6678, L6679, Key6680, C6681)
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

		tmp2449 := PrimValue(symshen_4_dsigf_d)

		tmp2450 := Call(__e, PrimFunc(symshen_4assoc_1_6), symy_1or_1n_2, tmp2445, tmp2449)

		tmp2451 := PrimSet(symshen_4_dsigf_d, tmp2450)

		_ = tmp2451

		tmp2452 := MakeNative(func(__e *ControlFlow) {
			V6687 := __e.Get(1)
			_ = V6687
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6683 := __e.Get(1)
				_ = B6683
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6684 := __e.Get(1)
					_ = L6684
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6685 := __e.Get(1)
						_ = Key6685
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6686 := __e.Get(1)
							_ = C6686
							tmp2453 := PrimCons(symboolean, Nil)

							tmp2454 := PrimCons(sym_1_1_6, tmp2453)

							tmp2455 := PrimCons(symnumber, tmp2454)

							tmp2456 := PrimCons(tmp2455, Nil)

							tmp2457 := PrimCons(sym_1_1_6, tmp2456)

							tmp2458 := PrimCons(symnumber, tmp2457)

							__e.TailApply(PrimFunc(symis_b), V6687, tmp2458, B6683, L6684, Key6685, C6686)
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

		tmp2459 := PrimValue(symshen_4_dsigf_d)

		tmp2460 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6, tmp2452, tmp2459)

		tmp2461 := PrimSet(symshen_4_dsigf_d, tmp2460)

		_ = tmp2461

		tmp2462 := MakeNative(func(__e *ControlFlow) {
			V6692 := __e.Get(1)
			_ = V6692
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6688 := __e.Get(1)
				_ = B6688
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6689 := __e.Get(1)
					_ = L6689
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6690 := __e.Get(1)
						_ = Key6690
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6691 := __e.Get(1)
							_ = C6691
							tmp2463 := PrimCons(symboolean, Nil)

							tmp2464 := PrimCons(sym_1_1_6, tmp2463)

							tmp2465 := PrimCons(symnumber, tmp2464)

							tmp2466 := PrimCons(tmp2465, Nil)

							tmp2467 := PrimCons(sym_1_1_6, tmp2466)

							tmp2468 := PrimCons(symnumber, tmp2467)

							__e.TailApply(PrimFunc(symis_b), V6692, tmp2468, B6688, L6689, Key6690, C6691)
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

		tmp2469 := PrimValue(symshen_4_dsigf_d)

		tmp2470 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5, tmp2462, tmp2469)

		tmp2471 := PrimSet(symshen_4_dsigf_d, tmp2470)

		_ = tmp2471

		tmp2472 := MakeNative(func(__e *ControlFlow) {
			V6697 := __e.Get(1)
			_ = V6697
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6693 := __e.Get(1)
				_ = B6693
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6694 := __e.Get(1)
					_ = L6694
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6695 := __e.Get(1)
						_ = Key6695
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6696 := __e.Get(1)
							_ = C6696
							tmp2473 := PrimCons(symboolean, Nil)

							tmp2474 := PrimCons(sym_1_1_6, tmp2473)

							tmp2475 := PrimCons(symnumber, tmp2474)

							tmp2476 := PrimCons(tmp2475, Nil)

							tmp2477 := PrimCons(sym_1_1_6, tmp2476)

							tmp2478 := PrimCons(symnumber, tmp2477)

							__e.TailApply(PrimFunc(symis_b), V6697, tmp2478, B6693, L6694, Key6695, C6696)
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

		tmp2479 := PrimValue(symshen_4_dsigf_d)

		tmp2480 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_6_a, tmp2472, tmp2479)

		tmp2481 := PrimSet(symshen_4_dsigf_d, tmp2480)

		_ = tmp2481

		tmp2482 := MakeNative(func(__e *ControlFlow) {
			V6702 := __e.Get(1)
			_ = V6702
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6698 := __e.Get(1)
				_ = B6698
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6699 := __e.Get(1)
					_ = L6699
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6700 := __e.Get(1)
						_ = Key6700
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6701 := __e.Get(1)
							_ = C6701
							tmp2483 := PrimCons(symboolean, Nil)

							tmp2484 := PrimCons(sym_1_1_6, tmp2483)

							tmp2485 := PrimCons(symnumber, tmp2484)

							tmp2486 := PrimCons(tmp2485, Nil)

							tmp2487 := PrimCons(sym_1_1_6, tmp2486)

							tmp2488 := PrimCons(symnumber, tmp2487)

							__e.TailApply(PrimFunc(symis_b), V6702, tmp2488, B6698, L6699, Key6700, C6701)
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

		tmp2489 := PrimValue(symshen_4_dsigf_d)

		tmp2490 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_5_a, tmp2482, tmp2489)

		tmp2491 := PrimSet(symshen_4_dsigf_d, tmp2490)

		_ = tmp2491

		tmp2492 := MakeNative(func(__e *ControlFlow) {
			V6707 := __e.Get(1)
			_ = V6707
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6703 := __e.Get(1)
				_ = B6703
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6704 := __e.Get(1)
					_ = L6704
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6705 := __e.Get(1)
						_ = Key6705
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6706 := __e.Get(1)
							_ = C6706
							tmp2493 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2494 := PrimCons(symboolean, Nil)

								tmp2495 := PrimCons(sym_1_1_6, tmp2494)

								tmp2496 := PrimCons(A, tmp2495)

								tmp2497 := PrimCons(tmp2496, Nil)

								tmp2498 := PrimCons(sym_1_1_6, tmp2497)

								tmp2499 := PrimCons(A, tmp2498)

								tmp2500 := Call(__e, PrimFunc(symis_b), V6707, tmp2499, B6703, L6704, Key6705, C6706)

								__e.TailApply(PrimFunc(symshen_4gc), B6703, tmp2500)
								return

							}, 1)

							tmp2501 := Call(__e, PrimFunc(symshen_4newpv), B6703)

							__e.TailApply(tmp2493, tmp2501)
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

		tmp2502 := PrimValue(symshen_4_dsigf_d)

		tmp2503 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a, tmp2492, tmp2502)

		tmp2504 := PrimSet(symshen_4_dsigf_d, tmp2503)

		_ = tmp2504

		tmp2505 := MakeNative(func(__e *ControlFlow) {
			V6712 := __e.Get(1)
			_ = V6712
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6708 := __e.Get(1)
				_ = B6708
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6709 := __e.Get(1)
					_ = L6709
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6710 := __e.Get(1)
						_ = Key6710
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6711 := __e.Get(1)
							_ = C6711
							tmp2506 := PrimCons(symnumber, Nil)

							tmp2507 := PrimCons(sym_1_1_6, tmp2506)

							tmp2508 := PrimCons(symnumber, tmp2507)

							tmp2509 := PrimCons(tmp2508, Nil)

							tmp2510 := PrimCons(sym_1_1_6, tmp2509)

							tmp2511 := PrimCons(symnumber, tmp2510)

							__e.TailApply(PrimFunc(symis_b), V6712, tmp2511, B6708, L6709, Key6710, C6711)
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

		tmp2512 := PrimValue(symshen_4_dsigf_d)

		tmp2513 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_7, tmp2505, tmp2512)

		tmp2514 := PrimSet(symshen_4_dsigf_d, tmp2513)

		_ = tmp2514

		tmp2515 := MakeNative(func(__e *ControlFlow) {
			V6717 := __e.Get(1)
			_ = V6717
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6713 := __e.Get(1)
				_ = B6713
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6714 := __e.Get(1)
					_ = L6714
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6715 := __e.Get(1)
						_ = Key6715
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6716 := __e.Get(1)
							_ = C6716
							tmp2516 := PrimCons(symnumber, Nil)

							tmp2517 := PrimCons(sym_1_1_6, tmp2516)

							tmp2518 := PrimCons(symnumber, tmp2517)

							tmp2519 := PrimCons(tmp2518, Nil)

							tmp2520 := PrimCons(sym_1_1_6, tmp2519)

							tmp2521 := PrimCons(symnumber, tmp2520)

							__e.TailApply(PrimFunc(symis_b), V6717, tmp2521, B6713, L6714, Key6715, C6716)
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

		tmp2522 := PrimValue(symshen_4_dsigf_d)

		tmp2523 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_c, tmp2515, tmp2522)

		tmp2524 := PrimSet(symshen_4_dsigf_d, tmp2523)

		_ = tmp2524

		tmp2525 := MakeNative(func(__e *ControlFlow) {
			V6722 := __e.Get(1)
			_ = V6722
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6718 := __e.Get(1)
				_ = B6718
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6719 := __e.Get(1)
					_ = L6719
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6720 := __e.Get(1)
						_ = Key6720
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6721 := __e.Get(1)
							_ = C6721
							tmp2526 := PrimCons(symnumber, Nil)

							tmp2527 := PrimCons(sym_1_1_6, tmp2526)

							tmp2528 := PrimCons(symnumber, tmp2527)

							tmp2529 := PrimCons(tmp2528, Nil)

							tmp2530 := PrimCons(sym_1_1_6, tmp2529)

							tmp2531 := PrimCons(symnumber, tmp2530)

							__e.TailApply(PrimFunc(symis_b), V6722, tmp2531, B6718, L6719, Key6720, C6721)
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

		tmp2532 := PrimValue(symshen_4_dsigf_d)

		tmp2533 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_1, tmp2525, tmp2532)

		tmp2534 := PrimSet(symshen_4_dsigf_d, tmp2533)

		_ = tmp2534

		tmp2535 := MakeNative(func(__e *ControlFlow) {
			V6727 := __e.Get(1)
			_ = V6727
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6723 := __e.Get(1)
				_ = B6723
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6724 := __e.Get(1)
					_ = L6724
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6725 := __e.Get(1)
						_ = Key6725
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6726 := __e.Get(1)
							_ = C6726
							tmp2536 := PrimCons(symnumber, Nil)

							tmp2537 := PrimCons(sym_1_1_6, tmp2536)

							tmp2538 := PrimCons(symnumber, tmp2537)

							tmp2539 := PrimCons(tmp2538, Nil)

							tmp2540 := PrimCons(sym_1_1_6, tmp2539)

							tmp2541 := PrimCons(symnumber, tmp2540)

							__e.TailApply(PrimFunc(symis_b), V6727, tmp2541, B6723, L6724, Key6725, C6726)
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

		tmp2542 := PrimValue(symshen_4_dsigf_d)

		tmp2543 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_d, tmp2535, tmp2542)

		tmp2544 := PrimSet(symshen_4_dsigf_d, tmp2543)

		_ = tmp2544

		tmp2545 := MakeNative(func(__e *ControlFlow) {
			V6732 := __e.Get(1)
			_ = V6732
			__e.Return(MakeNative(func(__e *ControlFlow) {
				B6728 := __e.Get(1)
				_ = B6728
				__e.Return(MakeNative(func(__e *ControlFlow) {
					L6729 := __e.Get(1)
					_ = L6729
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Key6730 := __e.Get(1)
						_ = Key6730
						__e.Return(MakeNative(func(__e *ControlFlow) {
							C6731 := __e.Get(1)
							_ = C6731
							tmp2546 := MakeNative(func(__e *ControlFlow) {
								A := __e.Get(1)
								_ = A
								tmp2547 := MakeNative(func(__e *ControlFlow) {
									B := __e.Get(1)
									_ = B
									tmp2548 := PrimCons(symboolean, Nil)

									tmp2549 := PrimCons(sym_1_1_6, tmp2548)

									tmp2550 := PrimCons(B, tmp2549)

									tmp2551 := PrimCons(tmp2550, Nil)

									tmp2552 := PrimCons(sym_1_1_6, tmp2551)

									tmp2553 := PrimCons(A, tmp2552)

									tmp2554 := Call(__e, PrimFunc(symis_b), V6732, tmp2553, B6728, L6729, Key6730, C6731)

									__e.TailApply(PrimFunc(symshen_4gc), B6728, tmp2554)
									return

								}, 1)

								tmp2555 := Call(__e, PrimFunc(symshen_4newpv), B6728)

								tmp2556 := Call(__e, tmp2547, tmp2555)

								__e.TailApply(PrimFunc(symshen_4gc), B6728, tmp2556)
								return

							}, 1)

							tmp2557 := Call(__e, PrimFunc(symshen_4newpv), B6728)

							__e.TailApply(tmp2546, tmp2557)
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

		tmp2558 := PrimValue(symshen_4_dsigf_d)

		tmp2559 := Call(__e, PrimFunc(symshen_4assoc_1_6), sym_a_a, tmp2545, tmp2558)

		__e.Return(PrimSet(symshen_4_dsigf_d, tmp2559))
		return

	}, 0)

	tmp2560 := Call(__e, ns2_1set, symshen_4initialise_1signedfuncs, tmp815)

	_ = tmp2560

	tmp2561 := MakeNative(func(__e *ControlFlow) {
		tmp2562 := MakeNative(func(__e *ControlFlow) {
			Y1196 := __e.Get(1)
			_ = Y1196
			__e.TailApply(PrimFunc(symshen_4tuple), Y1196)
			return
		}, 1)

		tmp2563 := PrimCons(symshen_4tuple, tmp2562)

		tmp2564 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2563)

		_ = tmp2564

		tmp2565 := MakeNative(func(__e *ControlFlow) {
			Y1195 := __e.Get(1)
			_ = Y1195
			__e.TailApply(PrimFunc(symshen_4pvar), Y1195)
			return
		}, 1)

		tmp2566 := PrimCons(symshen_4pvar, tmp2565)

		tmp2567 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2566)

		_ = tmp2567

		tmp2568 := MakeNative(func(__e *ControlFlow) {
			Y1194 := __e.Get(1)
			_ = Y1194
			__e.TailApply(PrimFunc(symshen_4dictionary), Y1194)
			return
		}, 1)

		tmp2569 := PrimCons(symshen_4dictionary, tmp2568)

		tmp2570 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2569)

		_ = tmp2570

		tmp2571 := MakeNative(func(__e *ControlFlow) {
			Y1193 := __e.Get(1)
			_ = Y1193
			__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1193)
			return
		}, 1)

		tmp2572 := PrimCons(symshen_4print_1prolog_1vector, tmp2571)

		tmp2573 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2572)

		_ = tmp2573

		tmp2574 := MakeNative(func(__e *ControlFlow) {
			Y1192 := __e.Get(1)
			_ = Y1192
			__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1192)
			return
		}, 1)

		tmp2575 := PrimCons(symshen_4print_1freshterm, tmp2574)

		tmp2576 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2575)

		_ = tmp2576

		tmp2577 := MakeNative(func(__e *ControlFlow) {
			Y1191 := __e.Get(1)
			_ = Y1191
			__e.TailApply(PrimFunc(symshen_4printF), Y1191)
			return
		}, 1)

		tmp2578 := PrimCons(symshen_4printF, tmp2577)

		tmp2579 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2578)

		_ = tmp2579

		tmp2580 := MakeNative(func(__e *ControlFlow) {
			Y1190 := __e.Get(1)
			_ = Y1190
			__e.TailApply(PrimFunc(symabsolute), Y1190)
			return
		}, 1)

		tmp2581 := PrimCons(symabsolute, tmp2580)

		tmp2582 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2581)

		_ = tmp2582

		tmp2583 := MakeNative(func(__e *ControlFlow) {
			Y1189 := __e.Get(1)
			_ = Y1189
			__e.Return(PrimIsVector(Y1189))
			return
		}, 1)

		tmp2584 := PrimCons(symabsvector_2, tmp2583)

		tmp2585 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2584)

		_ = tmp2585

		tmp2586 := MakeNative(func(__e *ControlFlow) {
			Y1188 := __e.Get(1)
			_ = Y1188
			__e.Return(PrimAbsvector(Y1188))
			return
		}, 1)

		tmp2587 := PrimCons(symabsvector, tmp2586)

		tmp2588 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2587)

		_ = tmp2588

		tmp2589 := MakeNative(func(__e *ControlFlow) {
			Y1185 := __e.Get(1)
			_ = Y1185
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1186 := __e.Get(1)
				_ = Y1186
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1187 := __e.Get(1)
					_ = Y1187
					__e.Return(PrimVectorSet(Y1185, Y1186, Y1187))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp2590 := PrimCons(symaddress_1_6, tmp2589)

		tmp2591 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2590)

		_ = tmp2591

		tmp2592 := MakeNative(func(__e *ControlFlow) {
			Y1183 := __e.Get(1)
			_ = Y1183
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1184 := __e.Get(1)
				_ = Y1184
				__e.TailApply(PrimFunc(symadjoin), Y1183, Y1184)
				return
			}, 1))
			return
		}, 1)

		tmp2593 := PrimCons(symadjoin, tmp2592)

		tmp2594 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2593)

		_ = tmp2594

		tmp2595 := MakeNative(func(__e *ControlFlow) {
			Y1181 := __e.Get(1)
			_ = Y1181
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1182 := __e.Get(1)
				_ = Y1182
				if True == Y1181 {
					if True == Y1182 {
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

		tmp2598 := PrimCons(symand, tmp2595)

		tmp2599 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2598)

		_ = tmp2599

		tmp2600 := MakeNative(func(__e *ControlFlow) {
			Y1179 := __e.Get(1)
			_ = Y1179
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1180 := __e.Get(1)
				_ = Y1180
				__e.TailApply(PrimFunc(symappend), Y1179, Y1180)
				return
			}, 1))
			return
		}, 1)

		tmp2601 := PrimCons(symappend, tmp2600)

		tmp2602 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2601)

		_ = tmp2602

		tmp2603 := MakeNative(func(__e *ControlFlow) {
			Y1178 := __e.Get(1)
			_ = Y1178
			__e.TailApply(PrimFunc(symarity), Y1178)
			return
		}, 1)

		tmp2604 := PrimCons(symarity, tmp2603)

		tmp2605 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2604)

		_ = tmp2605

		tmp2606 := MakeNative(func(__e *ControlFlow) {
			Y1176 := __e.Get(1)
			_ = Y1176
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1177 := __e.Get(1)
				_ = Y1177
				__e.TailApply(PrimFunc(symassoc), Y1176, Y1177)
				return
			}, 1))
			return
		}, 1)

		tmp2607 := PrimCons(symassoc, tmp2606)

		tmp2608 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2607)

		_ = tmp2608

		tmp2609 := MakeNative(func(__e *ControlFlow) {
			Y1175 := __e.Get(1)
			_ = Y1175
			__e.TailApply(PrimFunc(symatom_2), Y1175)
			return
		}, 1)

		tmp2610 := PrimCons(symatom_2, tmp2609)

		tmp2611 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2610)

		_ = tmp2611

		tmp2612 := MakeNative(func(__e *ControlFlow) {
			Y1174 := __e.Get(1)
			_ = Y1174
			__e.TailApply(PrimFunc(symboolean_2), Y1174)
			return
		}, 1)

		tmp2613 := PrimCons(symboolean_2, tmp2612)

		tmp2614 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2613)

		_ = tmp2614

		tmp2615 := MakeNative(func(__e *ControlFlow) {
			Y1173 := __e.Get(1)
			_ = Y1173
			__e.TailApply(PrimFunc(symbootstrap), Y1173)
			return
		}, 1)

		tmp2616 := PrimCons(symbootstrap, tmp2615)

		tmp2617 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2616)

		_ = tmp2617

		tmp2618 := MakeNative(func(__e *ControlFlow) {
			Y1172 := __e.Get(1)
			_ = Y1172
			__e.TailApply(PrimFunc(symbound_2), Y1172)
			return
		}, 1)

		tmp2619 := PrimCons(symbound_2, tmp2618)

		tmp2620 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2619)

		_ = tmp2620

		tmp2621 := MakeNative(func(__e *ControlFlow) {
			Y1166 := __e.Get(1)
			_ = Y1166
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1167 := __e.Get(1)
				_ = Y1167
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1168 := __e.Get(1)
					_ = Y1168
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y1169 := __e.Get(1)
						_ = Y1169
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y1170 := __e.Get(1)
							_ = Y1170
							__e.Return(MakeNative(func(__e *ControlFlow) {
								Y1171 := __e.Get(1)
								_ = Y1171
								__e.TailApply(PrimFunc(symbind), Y1166, Y1167, Y1168, Y1169, Y1170, Y1171)
								return
							}, 1))
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

		tmp2622 := PrimCons(symbind, tmp2621)

		tmp2623 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2622)

		_ = tmp2623

		tmp2624 := MakeNative(func(__e *ControlFlow) {
			Y1161 := __e.Get(1)
			_ = Y1161
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1162 := __e.Get(1)
				_ = Y1162
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1163 := __e.Get(1)
					_ = Y1163
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y1164 := __e.Get(1)
						_ = Y1164
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y1165 := __e.Get(1)
							_ = Y1165
							__e.TailApply(PrimFunc(symcall), Y1161, Y1162, Y1163, Y1164, Y1165)
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

		tmp2625 := PrimCons(symcall, tmp2624)

		tmp2626 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2625)

		_ = tmp2626

		tmp2627 := MakeNative(func(__e *ControlFlow) {
			Y1160 := __e.Get(1)
			_ = Y1160
			__e.TailApply(PrimFunc(symcd), Y1160)
			return
		}, 1)

		tmp2628 := PrimCons(symcd, tmp2627)

		tmp2629 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2628)

		_ = tmp2629

		tmp2630 := MakeNative(func(__e *ControlFlow) {
			Y1158 := __e.Get(1)
			_ = Y1158
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1159 := __e.Get(1)
				_ = Y1159
				__e.TailApply(PrimFunc(symcompile), Y1158, Y1159)
				return
			}, 1))
			return
		}, 1)

		tmp2631 := PrimCons(symcompile, tmp2630)

		tmp2632 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2631)

		_ = tmp2632

		tmp2633 := MakeNative(func(__e *ControlFlow) {
			Y1156 := __e.Get(1)
			_ = Y1156
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1157 := __e.Get(1)
				_ = Y1157
				__e.TailApply(PrimFunc(symconcat), Y1156, Y1157)
				return
			}, 1))
			return
		}, 1)

		tmp2634 := PrimCons(symconcat, tmp2633)

		tmp2635 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2634)

		_ = tmp2635

		tmp2636 := MakeNative(func(__e *ControlFlow) {
			Y1154 := __e.Get(1)
			_ = Y1154
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1155 := __e.Get(1)
				_ = Y1155
				__e.Return(PrimCons(Y1154, Y1155))
				return
			}, 1))
			return
		}, 1)

		tmp2637 := PrimCons(symcons, tmp2636)

		tmp2638 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2637)

		_ = tmp2638

		tmp2639 := MakeNative(func(__e *ControlFlow) {
			Y1153 := __e.Get(1)
			_ = Y1153
			__e.Return(PrimIsPair(Y1153))
			return
		}, 1)

		tmp2640 := PrimCons(symcons_2, tmp2639)

		tmp2641 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2640)

		_ = tmp2641

		tmp2642 := MakeNative(func(__e *ControlFlow) {
			Y1151 := __e.Get(1)
			_ = Y1151
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1152 := __e.Get(1)
				_ = Y1152
				__e.Return(PrimStringConcat(Y1151, Y1152))
				return
			}, 1))
			return
		}, 1)

		tmp2643 := PrimCons(symcn, tmp2642)

		tmp2644 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2643)

		_ = tmp2644

		tmp2645 := MakeNative(func(__e *ControlFlow) {
			Y1150 := __e.Get(1)
			_ = Y1150
			__e.Return(PrimCloseStream(Y1150))
			return
		}, 1)

		tmp2646 := PrimCons(symclose, tmp2645)

		tmp2647 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2646)

		_ = tmp2647

		tmp2648 := MakeNative(func(__e *ControlFlow) {
			Y1148 := __e.Get(1)
			_ = Y1148
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1149 := __e.Get(1)
				_ = Y1149
				__e.TailApply(PrimFunc(symdeclare), Y1148, Y1149)
				return
			}, 1))
			return
		}, 1)

		tmp2649 := PrimCons(symdeclare, tmp2648)

		tmp2650 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2649)

		_ = tmp2650

		tmp2651 := MakeNative(func(__e *ControlFlow) {
			Y1147 := __e.Get(1)
			_ = Y1147
			__e.TailApply(PrimFunc(symdestroy), Y1147)
			return
		}, 1)

		tmp2652 := PrimCons(symdestroy, tmp2651)

		tmp2653 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2652)

		_ = tmp2653

		tmp2654 := MakeNative(func(__e *ControlFlow) {
			Y1145 := __e.Get(1)
			_ = Y1145
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1146 := __e.Get(1)
				_ = Y1146
				__e.TailApply(PrimFunc(symdifference), Y1145, Y1146)
				return
			}, 1))
			return
		}, 1)

		tmp2655 := PrimCons(symdifference, tmp2654)

		tmp2656 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2655)

		_ = tmp2656

		tmp2657 := MakeNative(func(__e *ControlFlow) {
			Y1143 := __e.Get(1)
			_ = Y1143
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1144 := __e.Get(1)
				_ = Y1144
				_ = Y1143

				__e.Return(Y1144)
				return

			}, 1))
			return
		}, 1)

		tmp2658 := PrimCons(symdo, tmp2657)

		tmp2659 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2658)

		_ = tmp2659

		tmp2660 := MakeNative(func(__e *ControlFlow) {
			Y1141 := __e.Get(1)
			_ = Y1141
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1142 := __e.Get(1)
				_ = Y1142
				__e.TailApply(PrimFunc(symelement_2), Y1141, Y1142)
				return
			}, 1))
			return
		}, 1)

		tmp2661 := PrimCons(symelement_2, tmp2660)

		tmp2662 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2661)

		_ = tmp2662

		tmp2663 := MakeNative(func(__e *ControlFlow) {
			Y1140 := __e.Get(1)
			_ = Y1140
			__e.TailApply(PrimFunc(symempty_2), Y1140)
			return
		}, 1)

		tmp2664 := PrimCons(symempty_2, tmp2663)

		tmp2665 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2664)

		_ = tmp2665

		tmp2666 := MakeNative(func(__e *ControlFlow) {
			Y1139 := __e.Get(1)
			_ = Y1139
			__e.TailApply(PrimFunc(symenable_1type_1theory), Y1139)
			return
		}, 1)

		tmp2667 := PrimCons(symenable_1type_1theory, tmp2666)

		tmp2668 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2667)

		_ = tmp2668

		tmp2669 := MakeNative(func(__e *ControlFlow) {
			Y1138 := __e.Get(1)
			_ = Y1138
			__e.TailApply(PrimFunc(symexternal), Y1138)
			return
		}, 1)

		tmp2670 := PrimCons(symexternal, tmp2669)

		tmp2671 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2670)

		_ = tmp2671

		tmp2672 := MakeNative(func(__e *ControlFlow) {
			Y1137 := __e.Get(1)
			_ = Y1137
			__e.Return(PrimErrorToString(Y1137))
			return
		}, 1)

		tmp2673 := PrimCons(symerror_1to_1string, tmp2672)

		tmp2674 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2673)

		_ = tmp2674

		tmp2675 := MakeNative(func(__e *ControlFlow) {
			Y1136 := __e.Get(1)
			_ = Y1136
			__e.TailApply(PrimFunc(symeval), Y1136)
			return
		}, 1)

		tmp2676 := PrimCons(symeval, tmp2675)

		tmp2677 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2676)

		_ = tmp2677

		tmp2678 := MakeNative(func(__e *ControlFlow) {
			Y1135 := __e.Get(1)
			_ = Y1135
			__e.TailApply(PrimFunc(symeval_1kl), Y1135)
			return
		}, 1)

		tmp2679 := PrimCons(symeval_1kl, tmp2678)

		tmp2680 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2679)

		_ = tmp2680

		tmp2681 := MakeNative(func(__e *ControlFlow) {
			Y1134 := __e.Get(1)
			_ = Y1134
			__e.TailApply(PrimFunc(symexplode), Y1134)
			return
		}, 1)

		tmp2682 := PrimCons(symexplode, tmp2681)

		tmp2683 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2682)

		_ = tmp2683

		tmp2684 := MakeNative(func(__e *ControlFlow) {
			Y1133 := __e.Get(1)
			_ = Y1133
			__e.TailApply(PrimFunc(symexternal), Y1133)
			return
		}, 1)

		tmp2685 := PrimCons(symexternal, tmp2684)

		tmp2686 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2685)

		_ = tmp2686

		tmp2687 := MakeNative(func(__e *ControlFlow) {
			Y1132 := __e.Get(1)
			_ = Y1132
			__e.TailApply(PrimFunc(symfactorise), Y1132)
			return
		}, 1)

		tmp2688 := PrimCons(symfactorise, tmp2687)

		tmp2689 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2688)

		_ = tmp2689

		tmp2690 := MakeNative(func(__e *ControlFlow) {
			Y1130 := __e.Get(1)
			_ = Y1130
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1131 := __e.Get(1)
				_ = Y1131
				__e.TailApply(PrimFunc(symfail_1if), Y1130, Y1131)
				return
			}, 1))
			return
		}, 1)

		tmp2691 := PrimCons(symfail_1if, tmp2690)

		tmp2692 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2691)

		_ = tmp2692

		tmp2693 := MakeNative(func(__e *ControlFlow) {
			Y1128 := __e.Get(1)
			_ = Y1128
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1129 := __e.Get(1)
				_ = Y1129
				__e.TailApply(PrimFunc(symfix), Y1128, Y1129)
				return
			}, 1))
			return
		}, 1)

		tmp2694 := PrimCons(symfix, tmp2693)

		tmp2695 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2694)

		_ = tmp2695

		tmp2696 := MakeNative(func(__e *ControlFlow) {
			Y1121 := __e.Get(1)
			_ = Y1121
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1122 := __e.Get(1)
				_ = Y1122
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1123 := __e.Get(1)
					_ = Y1123
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y1124 := __e.Get(1)
						_ = Y1124
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y1125 := __e.Get(1)
							_ = Y1125
							__e.Return(MakeNative(func(__e *ControlFlow) {
								Y1126 := __e.Get(1)
								_ = Y1126
								__e.Return(MakeNative(func(__e *ControlFlow) {
									Y1127 := __e.Get(1)
									_ = Y1127
									__e.TailApply(PrimFunc(symfindall), Y1121, Y1122, Y1123, Y1124, Y1125, Y1126, Y1127)
									return
								}, 1))
								return
							}, 1))
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

		tmp2697 := PrimCons(symfindall, tmp2696)

		tmp2698 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2697)

		_ = tmp2698

		tmp2699 := MakeNative(func(__e *ControlFlow) {
			Y1116 := __e.Get(1)
			_ = Y1116
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1117 := __e.Get(1)
				_ = Y1117
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1118 := __e.Get(1)
					_ = Y1118
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y1119 := __e.Get(1)
						_ = Y1119
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y1120 := __e.Get(1)
							_ = Y1120
							__e.TailApply(PrimFunc(symfork), Y1116, Y1117, Y1118, Y1119, Y1120)
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

		tmp2700 := PrimCons(symfork, tmp2699)

		tmp2701 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2700)

		_ = tmp2701

		tmp2702 := MakeNative(func(__e *ControlFlow) {
			Y1115 := __e.Get(1)
			_ = Y1115
			__e.Return(MakeNative(func(__e *ControlFlow) {
				__e.Return(Y1115)
				return
			}, 0))
			return
		}, 1)

		tmp2703 := PrimCons(symfreeze, tmp2702)

		tmp2704 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2703)

		_ = tmp2704

		tmp2705 := MakeNative(func(__e *ControlFlow) {
			Y1114 := __e.Get(1)
			_ = Y1114
			__e.TailApply(PrimFunc(symfst), Y1114)
			return
		}, 1)

		tmp2706 := PrimCons(symfst, tmp2705)

		tmp2707 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2706)

		_ = tmp2707

		tmp2708 := MakeNative(func(__e *ControlFlow) {
			Y1113 := __e.Get(1)
			_ = Y1113
			__e.TailApply(PrimFunc(symfn), Y1113)
			return
		}, 1)

		tmp2709 := PrimCons(symfn, tmp2708)

		tmp2710 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2709)

		_ = tmp2710

		tmp2711 := MakeNative(func(__e *ControlFlow) {
			Y1112 := __e.Get(1)
			_ = Y1112
			__e.TailApply(PrimFunc(symfunction), Y1112)
			return
		}, 1)

		tmp2712 := PrimCons(symfunction, tmp2711)

		tmp2713 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2712)

		_ = tmp2713

		tmp2714 := MakeNative(func(__e *ControlFlow) {
			Y1111 := __e.Get(1)
			_ = Y1111
			__e.TailApply(PrimFunc(symgensym), Y1111)
			return
		}, 1)

		tmp2715 := PrimCons(symgensym, tmp2714)

		tmp2716 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2715)

		_ = tmp2716

		tmp2717 := MakeNative(func(__e *ControlFlow) {
			Y1108 := __e.Get(1)
			_ = Y1108
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1109 := __e.Get(1)
				_ = Y1109
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1110 := __e.Get(1)
					_ = Y1110
					__e.TailApply(PrimFunc(symget), Y1108, Y1109, Y1110)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp2718 := PrimCons(symget, tmp2717)

		tmp2719 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2718)

		_ = tmp2719

		tmp2720 := MakeNative(func(__e *ControlFlow) {
			Y1107 := __e.Get(1)
			_ = Y1107
			__e.Return(PrimGetTime(Y1107))
			return
		}, 1)

		tmp2721 := PrimCons(symget_1time, tmp2720)

		tmp2722 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2721)

		_ = tmp2722

		tmp2723 := MakeNative(func(__e *ControlFlow) {
			Y1104 := __e.Get(1)
			_ = Y1104
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1105 := __e.Get(1)
				_ = Y1105
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1106 := __e.Get(1)
					_ = Y1106
					__e.Return(PrimVectorSet(Y1104, Y1105, Y1106))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp2724 := PrimCons(symaddress_1_6, tmp2723)

		tmp2725 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2724)

		_ = tmp2725

		tmp2726 := MakeNative(func(__e *ControlFlow) {
			Y1102 := __e.Get(1)
			_ = Y1102
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1103 := __e.Get(1)
				_ = Y1103
				__e.Return(PrimVectorGet(Y1102, Y1103))
				return
			}, 1))
			return
		}, 1)

		tmp2727 := PrimCons(sym_5_1address, tmp2726)

		tmp2728 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2727)

		_ = tmp2728

		tmp2729 := MakeNative(func(__e *ControlFlow) {
			Y1100 := __e.Get(1)
			_ = Y1100
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1101 := __e.Get(1)
				_ = Y1101
				__e.TailApply(PrimFunc(sym_5_1vector), Y1100, Y1101)
				return
			}, 1))
			return
		}, 1)

		tmp2730 := PrimCons(sym_5_1vector, tmp2729)

		tmp2731 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2730)

		_ = tmp2731

		tmp2732 := MakeNative(func(__e *ControlFlow) {
			Y1098 := __e.Get(1)
			_ = Y1098
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1099 := __e.Get(1)
				_ = Y1099
				__e.Return(PrimGreatThan(Y1098, Y1099))
				return
			}, 1))
			return
		}, 1)

		tmp2733 := PrimCons(sym_6, tmp2732)

		tmp2734 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2733)

		_ = tmp2734

		tmp2735 := MakeNative(func(__e *ControlFlow) {
			Y1096 := __e.Get(1)
			_ = Y1096
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1097 := __e.Get(1)
				_ = Y1097
				__e.Return(PrimGreatEqual(Y1096, Y1097))
				return
			}, 1))
			return
		}, 1)

		tmp2736 := PrimCons(sym_6_a, tmp2735)

		tmp2737 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2736)

		_ = tmp2737

		tmp2738 := MakeNative(func(__e *ControlFlow) {
			Y1094 := __e.Get(1)
			_ = Y1094
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1095 := __e.Get(1)
				_ = Y1095
				__e.Return(PrimEqual(Y1094, Y1095))
				return
			}, 1))
			return
		}, 1)

		tmp2739 := PrimCons(sym_a, tmp2738)

		tmp2740 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2739)

		_ = tmp2740

		tmp2741 := MakeNative(func(__e *ControlFlow) {
			Y1092 := __e.Get(1)
			_ = Y1092
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1093 := __e.Get(1)
				_ = Y1093
				__e.TailApply(PrimFunc(symhash), Y1092, Y1093)
				return
			}, 1))
			return
		}, 1)

		tmp2742 := PrimCons(symhash, tmp2741)

		tmp2743 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2742)

		_ = tmp2743

		tmp2744 := MakeNative(func(__e *ControlFlow) {
			Y1091 := __e.Get(1)
			_ = Y1091
			__e.Return(PrimHead(Y1091))
			return
		}, 1)

		tmp2745 := PrimCons(symhd, tmp2744)

		tmp2746 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2745)

		_ = tmp2746

		tmp2747 := MakeNative(func(__e *ControlFlow) {
			Y1090 := __e.Get(1)
			_ = Y1090
			__e.TailApply(PrimFunc(symhdv), Y1090)
			return
		}, 1)

		tmp2748 := PrimCons(symhdv, tmp2747)

		tmp2749 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2748)

		_ = tmp2749

		tmp2750 := MakeNative(func(__e *ControlFlow) {
			Y1089 := __e.Get(1)
			_ = Y1089
			__e.TailApply(PrimFunc(symhdstr), Y1089)
			return
		}, 1)

		tmp2751 := PrimCons(symhdstr, tmp2750)

		tmp2752 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2751)

		_ = tmp2752

		tmp2753 := MakeNative(func(__e *ControlFlow) {
			Y1088 := __e.Get(1)
			_ = Y1088
			__e.TailApply(PrimFunc(symhead), Y1088)
			return
		}, 1)

		tmp2754 := PrimCons(symhead, tmp2753)

		tmp2755 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2754)

		_ = tmp2755

		tmp2756 := MakeNative(func(__e *ControlFlow) {
			Y1087 := __e.Get(1)
			_ = Y1087
			__e.TailApply(PrimFunc(symhush), Y1087)
			return
		}, 1)

		tmp2757 := PrimCons(symhush, tmp2756)

		tmp2758 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2757)

		_ = tmp2758

		tmp2759 := MakeNative(func(__e *ControlFlow) {
			Y1084 := __e.Get(1)
			_ = Y1084
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1085 := __e.Get(1)
				_ = Y1085
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1086 := __e.Get(1)
					_ = Y1086
					if True == Y1084 {
						__e.Return(Y1085)
						return
					} else {
						__e.Return(Y1086)
						return
					}
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp2761 := PrimCons(symif, tmp2759)

		tmp2762 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2761)

		_ = tmp2762

		tmp2763 := MakeNative(func(__e *ControlFlow) {
			Y1083 := __e.Get(1)
			_ = Y1083
			__e.TailApply(PrimFunc(syminclude), Y1083)
			return
		}, 1)

		tmp2764 := PrimCons(syminclude, tmp2763)

		tmp2765 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2764)

		_ = tmp2765

		tmp2766 := MakeNative(func(__e *ControlFlow) {
			Y1082 := __e.Get(1)
			_ = Y1082
			__e.TailApply(PrimFunc(symin_1package), Y1082)
			return
		}, 1)

		tmp2767 := PrimCons(symin_1package, tmp2766)

		tmp2768 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2767)

		_ = tmp2768

		tmp2769 := MakeNative(func(__e *ControlFlow) {
			Y1081 := __e.Get(1)
			_ = Y1081
			__e.Return(PrimIsInteger(Y1081))
			return
		}, 1)

		tmp2770 := PrimCons(syminteger_2, tmp2769)

		tmp2771 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2770)

		_ = tmp2771

		tmp2772 := MakeNative(func(__e *ControlFlow) {
			Y1080 := __e.Get(1)
			_ = Y1080
			__e.TailApply(PrimFunc(syminternal), Y1080)
			return
		}, 1)

		tmp2773 := PrimCons(syminternal, tmp2772)

		tmp2774 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2773)

		_ = tmp2774

		tmp2775 := MakeNative(func(__e *ControlFlow) {
			Y1079 := __e.Get(1)
			_ = Y1079
			__e.Return(PrimIntern(Y1079))
			return
		}, 1)

		tmp2776 := PrimCons(symintern, tmp2775)

		tmp2777 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2776)

		_ = tmp2777

		tmp2778 := MakeNative(func(__e *ControlFlow) {
			Y1078 := __e.Get(1)
			_ = Y1078
			__e.TailApply(PrimFunc(syminput), Y1078)
			return
		}, 1)

		tmp2779 := PrimCons(syminput, tmp2778)

		tmp2780 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2779)

		_ = tmp2780

		tmp2781 := MakeNative(func(__e *ControlFlow) {
			Y1076 := __e.Get(1)
			_ = Y1076
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1077 := __e.Get(1)
				_ = Y1077
				__e.TailApply(PrimFunc(syminput_7), Y1076, Y1077)
				return
			}, 1))
			return
		}, 1)

		tmp2782 := PrimCons(syminput_7, tmp2781)

		tmp2783 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2782)

		_ = tmp2783

		tmp2784 := MakeNative(func(__e *ControlFlow) {
			Y1075 := __e.Get(1)
			_ = Y1075
			__e.TailApply(PrimFunc(syminclude_1all_1but), Y1075)
			return
		}, 1)

		tmp2785 := PrimCons(syminclude_1all_1but, tmp2784)

		tmp2786 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2785)

		_ = tmp2786

		tmp2787 := MakeNative(func(__e *ControlFlow) {
			Y1073 := __e.Get(1)
			_ = Y1073
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1074 := __e.Get(1)
				_ = Y1074
				__e.TailApply(PrimFunc(symintersection), Y1073, Y1074)
				return
			}, 1))
			return
		}, 1)

		tmp2788 := PrimCons(symintersection, tmp2787)

		tmp2789 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2788)

		_ = tmp2789

		tmp2790 := MakeNative(func(__e *ControlFlow) {
			Y1072 := __e.Get(1)
			_ = Y1072
			__e.TailApply(PrimFunc(syminternal), Y1072)
			return
		}, 1)

		tmp2791 := PrimCons(syminternal, tmp2790)

		tmp2792 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2791)

		_ = tmp2792

		tmp2793 := MakeNative(func(__e *ControlFlow) {
			Y1066 := __e.Get(1)
			_ = Y1066
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1067 := __e.Get(1)
				_ = Y1067
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1068 := __e.Get(1)
					_ = Y1068
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y1069 := __e.Get(1)
						_ = Y1069
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y1070 := __e.Get(1)
							_ = Y1070
							__e.Return(MakeNative(func(__e *ControlFlow) {
								Y1071 := __e.Get(1)
								_ = Y1071
								__e.TailApply(PrimFunc(symis), Y1066, Y1067, Y1068, Y1069, Y1070, Y1071)
								return
							}, 1))
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

		tmp2794 := PrimCons(symis, tmp2793)

		tmp2795 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2794)

		_ = tmp2795

		tmp2796 := MakeNative(func(__e *ControlFlow) {
			Y1060 := __e.Get(1)
			_ = Y1060
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1061 := __e.Get(1)
				_ = Y1061
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1062 := __e.Get(1)
					_ = Y1062
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y1063 := __e.Get(1)
						_ = Y1063
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y1064 := __e.Get(1)
							_ = Y1064
							__e.Return(MakeNative(func(__e *ControlFlow) {
								Y1065 := __e.Get(1)
								_ = Y1065
								__e.TailApply(PrimFunc(symis_b), Y1060, Y1061, Y1062, Y1063, Y1064, Y1065)
								return
							}, 1))
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

		tmp2797 := PrimCons(symis_b, tmp2796)

		tmp2798 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2797)

		_ = tmp2798

		tmp2799 := MakeNative(func(__e *ControlFlow) {
			Y1059 := __e.Get(1)
			_ = Y1059
			__e.TailApply(PrimFunc(symlength), Y1059)
			return
		}, 1)

		tmp2800 := PrimCons(symlength, tmp2799)

		tmp2801 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2800)

		_ = tmp2801

		tmp2802 := MakeNative(func(__e *ControlFlow) {
			Y1058 := __e.Get(1)
			_ = Y1058
			__e.TailApply(PrimFunc(symlimit), Y1058)
			return
		}, 1)

		tmp2803 := PrimCons(symlimit, tmp2802)

		tmp2804 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2803)

		_ = tmp2804

		tmp2805 := MakeNative(func(__e *ControlFlow) {
			Y1057 := __e.Get(1)
			_ = Y1057
			__e.TailApply(PrimFunc(symlineread), Y1057)
			return
		}, 1)

		tmp2806 := PrimCons(symlineread, tmp2805)

		tmp2807 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2806)

		_ = tmp2807

		tmp2808 := MakeNative(func(__e *ControlFlow) {
			Y1056 := __e.Get(1)
			_ = Y1056
			__e.TailApply(PrimFunc(symload), Y1056)
			return
		}, 1)

		tmp2809 := PrimCons(symload, tmp2808)

		tmp2810 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2809)

		_ = tmp2810

		tmp2811 := MakeNative(func(__e *ControlFlow) {
			Y1054 := __e.Get(1)
			_ = Y1054
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1055 := __e.Get(1)
				_ = Y1055
				__e.Return(PrimLessThan(Y1054, Y1055))
				return
			}, 1))
			return
		}, 1)

		tmp2812 := PrimCons(sym_5, tmp2811)

		tmp2813 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2812)

		_ = tmp2813

		tmp2814 := MakeNative(func(__e *ControlFlow) {
			Y1052 := __e.Get(1)
			_ = Y1052
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1053 := __e.Get(1)
				_ = Y1053
				__e.Return(PrimLessEqual(Y1052, Y1053))
				return
			}, 1))
			return
		}, 1)

		tmp2815 := PrimCons(sym_5_a, tmp2814)

		tmp2816 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2815)

		_ = tmp2816

		tmp2817 := MakeNative(func(__e *ControlFlow) {
			Y1051 := __e.Get(1)
			_ = Y1051
			__e.TailApply(PrimFunc(symvector), Y1051)
			return
		}, 1)

		tmp2818 := PrimCons(symvector, tmp2817)

		tmp2819 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2818)

		_ = tmp2819

		tmp2820 := MakeNative(func(__e *ControlFlow) {
			Y1050 := __e.Get(1)
			_ = Y1050
			__e.TailApply(PrimFunc(symmacroexpand), Y1050)
			return
		}, 1)

		tmp2821 := PrimCons(symmacroexpand, tmp2820)

		tmp2822 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2821)

		_ = tmp2822

		tmp2823 := MakeNative(func(__e *ControlFlow) {
			Y1048 := __e.Get(1)
			_ = Y1048
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1049 := __e.Get(1)
				_ = Y1049
				__e.TailApply(PrimFunc(symmap), Y1048, Y1049)
				return
			}, 1))
			return
		}, 1)

		tmp2824 := PrimCons(symmap, tmp2823)

		tmp2825 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2824)

		_ = tmp2825

		tmp2826 := MakeNative(func(__e *ControlFlow) {
			Y1046 := __e.Get(1)
			_ = Y1046
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1047 := __e.Get(1)
				_ = Y1047
				__e.TailApply(PrimFunc(symmapcan), Y1046, Y1047)
				return
			}, 1))
			return
		}, 1)

		tmp2827 := PrimCons(symmapcan, tmp2826)

		tmp2828 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2827)

		_ = tmp2828

		tmp2829 := MakeNative(func(__e *ControlFlow) {
			Y1045 := __e.Get(1)
			_ = Y1045
			__e.TailApply(PrimFunc(symmaxinferences), Y1045)
			return
		}, 1)

		tmp2830 := PrimCons(symmaxinferences, tmp2829)

		tmp2831 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2830)

		_ = tmp2831

		tmp2832 := MakeNative(func(__e *ControlFlow) {
			Y1044 := __e.Get(1)
			_ = Y1044
			__e.TailApply(PrimFunc(symnl), Y1044)
			return
		}, 1)

		tmp2833 := PrimCons(symnl, tmp2832)

		tmp2834 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2833)

		_ = tmp2834

		tmp2835 := MakeNative(func(__e *ControlFlow) {
			Y1043 := __e.Get(1)
			_ = Y1043
			__e.Return(PrimNot(Y1043))
			return
		}, 1)

		tmp2836 := PrimCons(symnot, tmp2835)

		tmp2837 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2836)

		_ = tmp2837

		tmp2838 := MakeNative(func(__e *ControlFlow) {
			Y1041 := __e.Get(1)
			_ = Y1041
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1042 := __e.Get(1)
				_ = Y1042
				__e.TailApply(PrimFunc(symnth), Y1041, Y1042)
				return
			}, 1))
			return
		}, 1)

		tmp2839 := PrimCons(symnth, tmp2838)

		tmp2840 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2839)

		_ = tmp2840

		tmp2841 := MakeNative(func(__e *ControlFlow) {
			Y1040 := __e.Get(1)
			_ = Y1040
			__e.Return(PrimNumberToString(Y1040))
			return
		}, 1)

		tmp2842 := PrimCons(symn_1_6string, tmp2841)

		tmp2843 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2842)

		_ = tmp2843

		tmp2844 := MakeNative(func(__e *ControlFlow) {
			Y1039 := __e.Get(1)
			_ = Y1039
			__e.Return(PrimIsNumber(Y1039))
			return
		}, 1)

		tmp2845 := PrimCons(symnumber_2, tmp2844)

		tmp2846 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2845)

		_ = tmp2846

		tmp2847 := MakeNative(func(__e *ControlFlow) {
			Y1038 := __e.Get(1)
			_ = Y1038
			__e.TailApply(PrimFunc(symoccurs_1check), Y1038)
			return
		}, 1)

		tmp2848 := PrimCons(symoccurs_1check, tmp2847)

		tmp2849 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2848)

		_ = tmp2849

		tmp2850 := MakeNative(func(__e *ControlFlow) {
			Y1036 := __e.Get(1)
			_ = Y1036
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1037 := __e.Get(1)
				_ = Y1037
				__e.TailApply(PrimFunc(symoccurrences), Y1036, Y1037)
				return
			}, 1))
			return
		}, 1)

		tmp2851 := PrimCons(symoccurrences, tmp2850)

		tmp2852 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2851)

		_ = tmp2852

		tmp2853 := MakeNative(func(__e *ControlFlow) {
			Y1035 := __e.Get(1)
			_ = Y1035
			__e.TailApply(PrimFunc(symoccurs_1check), Y1035)
			return
		}, 1)

		tmp2854 := PrimCons(symoccurs_1check, tmp2853)

		tmp2855 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2854)

		_ = tmp2855

		tmp2856 := MakeNative(func(__e *ControlFlow) {
			Y1033 := __e.Get(1)
			_ = Y1033
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1034 := __e.Get(1)
				_ = Y1034
				__e.Return(PrimOpenStream(Y1033, Y1034))
				return
			}, 1))
			return
		}, 1)

		tmp2857 := PrimCons(symopen, tmp2856)

		tmp2858 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2857)

		_ = tmp2858

		tmp2859 := MakeNative(func(__e *ControlFlow) {
			Y1032 := __e.Get(1)
			_ = Y1032
			__e.TailApply(PrimFunc(symoptimise), Y1032)
			return
		}, 1)

		tmp2860 := PrimCons(symoptimise, tmp2859)

		tmp2861 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2860)

		_ = tmp2861

		tmp2862 := MakeNative(func(__e *ControlFlow) {
			Y1030 := __e.Get(1)
			_ = Y1030
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1031 := __e.Get(1)
				_ = Y1031
				if True == Y1030 {
					__e.Return(True)
					return
				} else {
					if True == Y1031 {
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

		tmp2865 := PrimCons(symor, tmp2862)

		tmp2866 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2865)

		_ = tmp2866

		tmp2867 := MakeNative(func(__e *ControlFlow) {
			Y1029 := __e.Get(1)
			_ = Y1029
			__e.TailApply(PrimFunc(sympackage_2), Y1029)
			return
		}, 1)

		tmp2868 := PrimCons(sympackage_2, tmp2867)

		tmp2869 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2868)

		_ = tmp2869

		tmp2870 := MakeNative(func(__e *ControlFlow) {
			Y1027 := __e.Get(1)
			_ = Y1027
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1028 := __e.Get(1)
				_ = Y1028
				__e.Return(PrimPos(Y1027, Y1028))
				return
			}, 1))
			return
		}, 1)

		tmp2871 := PrimCons(sympos, tmp2870)

		tmp2872 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2871)

		_ = tmp2872

		tmp2873 := MakeNative(func(__e *ControlFlow) {
			Y1026 := __e.Get(1)
			_ = Y1026
			__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1026)
			return
		}, 1)

		tmp2874 := PrimCons(sympreclude_1all_1but, tmp2873)

		tmp2875 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2874)

		_ = tmp2875

		tmp2876 := MakeNative(func(__e *ControlFlow) {
			Y1025 := __e.Get(1)
			_ = Y1025
			__e.TailApply(PrimFunc(symprint), Y1025)
			return
		}, 1)

		tmp2877 := PrimCons(symprint, tmp2876)

		tmp2878 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2877)

		_ = tmp2878

		tmp2879 := MakeNative(func(__e *ControlFlow) {
			Y1024 := __e.Get(1)
			_ = Y1024
			__e.TailApply(PrimFunc(symprofile), Y1024)
			return
		}, 1)

		tmp2880 := PrimCons(symprofile, tmp2879)

		tmp2881 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2880)

		_ = tmp2881

		tmp2882 := MakeNative(func(__e *ControlFlow) {
			Y1023 := __e.Get(1)
			_ = Y1023
			__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Y1023)
			return
		}, 1)

		tmp2883 := PrimCons(symshen_4print_1prolog_1vector, tmp2882)

		tmp2884 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2883)

		_ = tmp2884

		tmp2885 := MakeNative(func(__e *ControlFlow) {
			Y1022 := __e.Get(1)
			_ = Y1022
			__e.TailApply(PrimFunc(symshen_4print_1freshterm), Y1022)
			return
		}, 1)

		tmp2886 := PrimCons(symshen_4print_1freshterm, tmp2885)

		tmp2887 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2886)

		_ = tmp2887

		tmp2888 := MakeNative(func(__e *ControlFlow) {
			Y1021 := __e.Get(1)
			_ = Y1021
			__e.TailApply(PrimFunc(symshen_4printF), Y1021)
			return
		}, 1)

		tmp2889 := PrimCons(symshen_4printF, tmp2888)

		tmp2890 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2889)

		_ = tmp2890

		tmp2891 := MakeNative(func(__e *ControlFlow) {
			Y1020 := __e.Get(1)
			_ = Y1020
			__e.TailApply(PrimFunc(symprolog_1memory), Y1020)
			return
		}, 1)

		tmp2892 := PrimCons(symprolog_1memory, tmp2891)

		tmp2893 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2892)

		_ = tmp2893

		tmp2894 := MakeNative(func(__e *ControlFlow) {
			Y1019 := __e.Get(1)
			_ = Y1019
			__e.TailApply(PrimFunc(symprofile_1results), Y1019)
			return
		}, 1)

		tmp2895 := PrimCons(symprofile_1results, tmp2894)

		tmp2896 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2895)

		_ = tmp2896

		tmp2897 := MakeNative(func(__e *ControlFlow) {
			Y1017 := __e.Get(1)
			_ = Y1017
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1018 := __e.Get(1)
				_ = Y1018
				__e.TailApply(PrimFunc(sympr), Y1017, Y1018)
				return
			}, 1))
			return
		}, 1)

		tmp2898 := PrimCons(sympr, tmp2897)

		tmp2899 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2898)

		_ = tmp2899

		tmp2900 := MakeNative(func(__e *ControlFlow) {
			Y1016 := __e.Get(1)
			_ = Y1016
			__e.TailApply(PrimFunc(symps), Y1016)
			return
		}, 1)

		tmp2901 := PrimCons(symps, tmp2900)

		tmp2902 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2901)

		_ = tmp2902

		tmp2903 := MakeNative(func(__e *ControlFlow) {
			Y1015 := __e.Get(1)
			_ = Y1015
			__e.TailApply(PrimFunc(sympreclude), Y1015)
			return
		}, 1)

		tmp2904 := PrimCons(sympreclude, tmp2903)

		tmp2905 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2904)

		_ = tmp2905

		tmp2906 := MakeNative(func(__e *ControlFlow) {
			Y1014 := __e.Get(1)
			_ = Y1014
			__e.TailApply(PrimFunc(sympreclude_1all_1but), Y1014)
			return
		}, 1)

		tmp2907 := PrimCons(sympreclude_1all_1but, tmp2906)

		tmp2908 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2907)

		_ = tmp2908

		tmp2909 := MakeNative(func(__e *ControlFlow) {
			Y1013 := __e.Get(1)
			_ = Y1013
			__e.TailApply(PrimFunc(symprotect), Y1013)
			return
		}, 1)

		tmp2910 := PrimCons(symprotect, tmp2909)

		tmp2911 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2910)

		_ = tmp2911

		tmp2912 := MakeNative(func(__e *ControlFlow) {
			Y1009 := __e.Get(1)
			_ = Y1009
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1010 := __e.Get(1)
				_ = Y1010
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y1011 := __e.Get(1)
					_ = Y1011
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y1012 := __e.Get(1)
						_ = Y1012
						__e.TailApply(PrimFunc(symput), Y1009, Y1010, Y1011, Y1012)
						return
					}, 1))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp2913 := PrimCons(symput, tmp2912)

		tmp2914 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2913)

		_ = tmp2914

		tmp2915 := MakeNative(func(__e *ControlFlow) {
			Y1008 := __e.Get(1)
			_ = Y1008
			__e.Return(PrimReadFileAsString(Y1008))
			return
		}, 1)

		tmp2916 := PrimCons(symread_1file_1as_1string, tmp2915)

		tmp2917 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2916)

		_ = tmp2917

		tmp2918 := MakeNative(func(__e *ControlFlow) {
			Y1007 := __e.Get(1)
			_ = Y1007
			__e.Return(PrimReadFileAsByteList(Y1007))
			return
		}, 1)

		tmp2919 := PrimCons(symread_1file_1as_1bytelist, tmp2918)

		tmp2920 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2919)

		_ = tmp2920

		tmp2921 := MakeNative(func(__e *ControlFlow) {
			Y1006 := __e.Get(1)
			_ = Y1006
			__e.TailApply(PrimFunc(symread_1file), Y1006)
			return
		}, 1)

		tmp2922 := PrimCons(symread_1file, tmp2921)

		tmp2923 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2922)

		_ = tmp2923

		tmp2924 := MakeNative(func(__e *ControlFlow) {
			Y1005 := __e.Get(1)
			_ = Y1005
			__e.TailApply(PrimFunc(symread), Y1005)
			return
		}, 1)

		tmp2925 := PrimCons(symread, tmp2924)

		tmp2926 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2925)

		_ = tmp2926

		tmp2927 := MakeNative(func(__e *ControlFlow) {
			Y1004 := __e.Get(1)
			_ = Y1004
			__e.Return(PrimReadByte(Y1004))
			return
		}, 1)

		tmp2928 := PrimCons(symread_1byte, tmp2927)

		tmp2929 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2928)

		_ = tmp2929

		tmp2930 := MakeNative(func(__e *ControlFlow) {
			Y1003 := __e.Get(1)
			_ = Y1003
			__e.TailApply(PrimFunc(symread_1from_1string), Y1003)
			return
		}, 1)

		tmp2931 := PrimCons(symread_1from_1string, tmp2930)

		tmp2932 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2931)

		_ = tmp2932

		tmp2933 := MakeNative(func(__e *ControlFlow) {
			Y1002 := __e.Get(1)
			_ = Y1002
			__e.TailApply(PrimFunc(symread_1from_1string_1unprocessed), Y1002)
			return
		}, 1)

		tmp2934 := PrimCons(symread_1from_1string_1unprocessed, tmp2933)

		tmp2935 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2934)

		_ = tmp2935

		tmp2936 := MakeNative(func(__e *ControlFlow) {
			Y1001 := __e.Get(1)
			_ = Y1001
			__e.TailApply(PrimFunc(symshen_4read_1unit_1string), Y1001)
			return
		}, 1)

		tmp2937 := PrimCons(symshen_4read_1unit_1string, tmp2936)

		tmp2938 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2937)

		_ = tmp2938

		tmp2939 := MakeNative(func(__e *ControlFlow) {
			Y999 := __e.Get(1)
			_ = Y999
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y1000 := __e.Get(1)
				_ = Y1000
				__e.TailApply(PrimFunc(symremove), Y999, Y1000)
				return
			}, 1))
			return
		}, 1)

		tmp2940 := PrimCons(symremove, tmp2939)

		tmp2941 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2940)

		_ = tmp2941

		tmp2942 := MakeNative(func(__e *ControlFlow) {
			Y998 := __e.Get(1)
			_ = Y998
			__e.TailApply(PrimFunc(symreverse), Y998)
			return
		}, 1)

		tmp2943 := PrimCons(symreverse, tmp2942)

		tmp2944 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2943)

		_ = tmp2944

		tmp2945 := MakeNative(func(__e *ControlFlow) {
			Y996 := __e.Get(1)
			_ = Y996
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y997 := __e.Get(1)
				_ = Y997
				__e.Return(PrimSet(Y996, Y997))
				return
			}, 1))
			return
		}, 1)

		tmp2946 := PrimCons(symset, tmp2945)

		tmp2947 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2946)

		_ = tmp2947

		tmp2948 := MakeNative(func(__e *ControlFlow) {
			Y995 := __e.Get(1)
			_ = Y995
			__e.Return(PrimSimpleError(Y995))
			return
		}, 1)

		tmp2949 := PrimCons(symsimple_1error, tmp2948)

		tmp2950 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2949)

		_ = tmp2950

		tmp2951 := MakeNative(func(__e *ControlFlow) {
			Y994 := __e.Get(1)
			_ = Y994
			__e.TailApply(PrimFunc(symsnd), Y994)
			return
		}, 1)

		tmp2952 := PrimCons(symsnd, tmp2951)

		tmp2953 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2952)

		_ = tmp2953

		tmp2954 := MakeNative(func(__e *ControlFlow) {
			Y992 := __e.Get(1)
			_ = Y992
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y993 := __e.Get(1)
				_ = Y993
				__e.TailApply(PrimFunc(symspecialise), Y992, Y993)
				return
			}, 1))
			return
		}, 1)

		tmp2955 := PrimCons(symspecialise, tmp2954)

		tmp2956 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2955)

		_ = tmp2956

		tmp2957 := MakeNative(func(__e *ControlFlow) {
			Y991 := __e.Get(1)
			_ = Y991
			__e.TailApply(PrimFunc(symspy), Y991)
			return
		}, 1)

		tmp2958 := PrimCons(symspy, tmp2957)

		tmp2959 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2958)

		_ = tmp2959

		tmp2960 := MakeNative(func(__e *ControlFlow) {
			Y990 := __e.Get(1)
			_ = Y990
			__e.TailApply(PrimFunc(symstep), Y990)
			return
		}, 1)

		tmp2961 := PrimCons(symstep, tmp2960)

		tmp2962 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2961)

		_ = tmp2962

		tmp2963 := MakeNative(func(__e *ControlFlow) {
			Y989 := __e.Get(1)
			_ = Y989
			__e.Return(PrimStr(Y989))
			return
		}, 1)

		tmp2964 := PrimCons(symstr, tmp2963)

		tmp2965 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2964)

		_ = tmp2965

		tmp2966 := MakeNative(func(__e *ControlFlow) {
			Y988 := __e.Get(1)
			_ = Y988
			__e.Return(PrimStringToNumber(Y988))
			return
		}, 1)

		tmp2967 := PrimCons(symstring_1_6n, tmp2966)

		tmp2968 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2967)

		_ = tmp2968

		tmp2969 := MakeNative(func(__e *ControlFlow) {
			Y987 := __e.Get(1)
			_ = Y987
			__e.TailApply(PrimFunc(symstring_1_6symbol), Y987)
			return
		}, 1)

		tmp2970 := PrimCons(symstring_1_6symbol, tmp2969)

		tmp2971 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2970)

		_ = tmp2971

		tmp2972 := MakeNative(func(__e *ControlFlow) {
			Y986 := __e.Get(1)
			_ = Y986
			__e.Return(PrimIsString(Y986))
			return
		}, 1)

		tmp2973 := PrimCons(symstring_2, tmp2972)

		tmp2974 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2973)

		_ = tmp2974

		tmp2975 := MakeNative(func(__e *ControlFlow) {
			Y983 := __e.Get(1)
			_ = Y983
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y984 := __e.Get(1)
				_ = Y984
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y985 := __e.Get(1)
					_ = Y985
					__e.TailApply(PrimFunc(symsubst), Y983, Y984, Y985)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp2976 := PrimCons(symsubst, tmp2975)

		tmp2977 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2976)

		_ = tmp2977

		tmp2978 := MakeNative(func(__e *ControlFlow) {
			Y982 := __e.Get(1)
			_ = Y982
			__e.TailApply(PrimFunc(symsum), Y982)
			return
		}, 1)

		tmp2979 := PrimCons(symsum, tmp2978)

		tmp2980 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2979)

		_ = tmp2980

		tmp2981 := MakeNative(func(__e *ControlFlow) {
			Y981 := __e.Get(1)
			_ = Y981
			__e.Return(PrimIsSymbol(Y981))
			return
		}, 1)

		tmp2982 := PrimCons(symsymbol_2, tmp2981)

		tmp2983 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2982)

		_ = tmp2983

		tmp2984 := MakeNative(func(__e *ControlFlow) {
			Y980 := __e.Get(1)
			_ = Y980
			__e.TailApply(PrimFunc(symsystemf), Y980)
			return
		}, 1)

		tmp2985 := PrimCons(symsystemf, tmp2984)

		tmp2986 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2985)

		_ = tmp2986

		tmp2987 := MakeNative(func(__e *ControlFlow) {
			Y979 := __e.Get(1)
			_ = Y979
			__e.TailApply(PrimFunc(symtail), Y979)
			return
		}, 1)

		tmp2988 := PrimCons(symtail, tmp2987)

		tmp2989 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2988)

		_ = tmp2989

		tmp2990 := MakeNative(func(__e *ControlFlow) {
			Y978 := __e.Get(1)
			_ = Y978
			__e.Return(PrimTail(Y978))
			return
		}, 1)

		tmp2991 := PrimCons(symtl, tmp2990)

		tmp2992 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2991)

		_ = tmp2992

		tmp2993 := MakeNative(func(__e *ControlFlow) {
			Y977 := __e.Get(1)
			_ = Y977
			__e.TailApply(PrimFunc(symtc), Y977)
			return
		}, 1)

		tmp2994 := PrimCons(symtc, tmp2993)

		tmp2995 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2994)

		_ = tmp2995

		tmp2996 := MakeNative(func(__e *ControlFlow) {
			Y976 := __e.Get(1)
			_ = Y976
			__e.TailApply(PrimFunc(symthaw), Y976)
			return
		}, 1)

		tmp2997 := PrimCons(symthaw, tmp2996)

		tmp2998 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp2997)

		_ = tmp2998

		tmp2999 := MakeNative(func(__e *ControlFlow) {
			Y975 := __e.Get(1)
			_ = Y975
			__e.Return(PrimTailString(Y975))
			return
		}, 1)

		tmp3000 := PrimCons(symtlstr, tmp2999)

		tmp3001 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3000)

		_ = tmp3001

		tmp3002 := MakeNative(func(__e *ControlFlow) {
			Y974 := __e.Get(1)
			_ = Y974
			__e.TailApply(PrimFunc(symtrack), Y974)
			return
		}, 1)

		tmp3003 := PrimCons(symtrack, tmp3002)

		tmp3004 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3003)

		_ = tmp3004

		tmp3005 := MakeNative(func(__e *ControlFlow) {
			Y972 := __e.Get(1)
			_ = Y972
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y973 := __e.Get(1)
				_ = Y973
				tmp3006 := MakeNative(func(__e *ControlFlow) {
					__e.Return(Y972)
					return
				}, 0)

				__e.TailApply(try_1catch, tmp3006, Y973)
				return

			}, 1))
			return
		}, 1)

		tmp3007 := PrimCons(symtrap_1error, tmp3005)

		tmp3008 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3007)

		_ = tmp3008

		tmp3009 := MakeNative(func(__e *ControlFlow) {
			Y971 := __e.Get(1)
			_ = Y971
			__e.TailApply(PrimFunc(symtuple_2), Y971)
			return
		}, 1)

		tmp3010 := PrimCons(symtuple_2, tmp3009)

		tmp3011 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3010)

		_ = tmp3011

		tmp3012 := MakeNative(func(__e *ControlFlow) {
			Y969 := __e.Get(1)
			_ = Y969
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y970 := __e.Get(1)
				_ = Y970
				__e.Return(Y969)
				return
			}, 1))
			return
		}, 1)

		tmp3013 := PrimCons(symtype, tmp3012)

		tmp3014 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3013)

		_ = tmp3014

		tmp3015 := MakeNative(func(__e *ControlFlow) {
			Y964 := __e.Get(1)
			_ = Y964
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y965 := __e.Get(1)
				_ = Y965
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y966 := __e.Get(1)
					_ = Y966
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y967 := __e.Get(1)
						_ = Y967
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y968 := __e.Get(1)
							_ = Y968
							__e.TailApply(PrimFunc(symreturn), Y964, Y965, Y966, Y967, Y968)
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

		tmp3016 := PrimCons(symreturn, tmp3015)

		tmp3017 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3016)

		_ = tmp3017

		tmp3018 := MakeNative(func(__e *ControlFlow) {
			Y963 := __e.Get(1)
			_ = Y963
			__e.TailApply(PrimFunc(symunabsolute), Y963)
			return
		}, 1)

		tmp3019 := PrimCons(symunabsolute, tmp3018)

		tmp3020 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3019)

		_ = tmp3020

		tmp3021 := MakeNative(func(__e *ControlFlow) {
			Y962 := __e.Get(1)
			_ = Y962
			__e.TailApply(PrimFunc(symundefmacro), Y962)
			return
		}, 1)

		tmp3022 := PrimCons(symundefmacro, tmp3021)

		tmp3023 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3022)

		_ = tmp3023

		tmp3024 := MakeNative(func(__e *ControlFlow) {
			Y959 := __e.Get(1)
			_ = Y959
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y960 := __e.Get(1)
				_ = Y960
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y961 := __e.Get(1)
					_ = Y961
					__e.TailApply(PrimFunc(symunput), Y959, Y960, Y961)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp3025 := PrimCons(symunput, tmp3024)

		tmp3026 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3025)

		_ = tmp3026

		tmp3027 := MakeNative(func(__e *ControlFlow) {
			Y958 := __e.Get(1)
			_ = Y958
			__e.TailApply(PrimFunc(symunprofile), Y958)
			return
		}, 1)

		tmp3028 := PrimCons(symunprofile, tmp3027)

		tmp3029 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3028)

		_ = tmp3029

		tmp3030 := MakeNative(func(__e *ControlFlow) {
			Y956 := __e.Get(1)
			_ = Y956
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y957 := __e.Get(1)
				_ = Y957
				__e.TailApply(PrimFunc(symunion), Y956, Y957)
				return
			}, 1))
			return
		}, 1)

		tmp3031 := PrimCons(symunion, tmp3030)

		tmp3032 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3031)

		_ = tmp3032

		tmp3033 := MakeNative(func(__e *ControlFlow) {
			Y955 := __e.Get(1)
			_ = Y955
			__e.TailApply(PrimFunc(symuntrack), Y955)
			return
		}, 1)

		tmp3034 := PrimCons(symuntrack, tmp3033)

		tmp3035 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3034)

		_ = tmp3035

		tmp3036 := MakeNative(func(__e *ControlFlow) {
			Y954 := __e.Get(1)
			_ = Y954
			__e.TailApply(PrimFunc(symundefmacro), Y954)
			return
		}, 1)

		tmp3037 := PrimCons(symundefmacro, tmp3036)

		tmp3038 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3037)

		_ = tmp3038

		tmp3039 := MakeNative(func(__e *ControlFlow) {
			Y952 := __e.Get(1)
			_ = Y952
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y953 := __e.Get(1)
				_ = Y953
				__e.TailApply(PrimFunc(symupdate_1lambda_1table), Y952, Y953)
				return
			}, 1))
			return
		}, 1)

		tmp3040 := PrimCons(symupdate_1lambda_1table, tmp3039)

		tmp3041 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3040)

		_ = tmp3041

		tmp3042 := MakeNative(func(__e *ControlFlow) {
			Y951 := __e.Get(1)
			_ = Y951
			__e.TailApply(PrimFunc(symvector), Y951)
			return
		}, 1)

		tmp3043 := PrimCons(symvector, tmp3042)

		tmp3044 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3043)

		_ = tmp3044

		tmp3045 := MakeNative(func(__e *ControlFlow) {
			Y950 := __e.Get(1)
			_ = Y950
			__e.TailApply(PrimFunc(symvector_2), Y950)
			return
		}, 1)

		tmp3046 := PrimCons(symvector_2, tmp3045)

		tmp3047 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3046)

		_ = tmp3047

		tmp3048 := MakeNative(func(__e *ControlFlow) {
			Y947 := __e.Get(1)
			_ = Y947
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y948 := __e.Get(1)
				_ = Y948
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y949 := __e.Get(1)
					_ = Y949
					__e.TailApply(PrimFunc(symvector_1_6), Y947, Y948, Y949)
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		tmp3049 := PrimCons(symvector_1_6, tmp3048)

		tmp3050 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3049)

		_ = tmp3050

		tmp3051 := MakeNative(func(__e *ControlFlow) {
			Y946 := __e.Get(1)
			_ = Y946
			__e.Return(PrimValue(Y946))
			return
		}, 1)

		tmp3052 := PrimCons(symvalue, tmp3051)

		tmp3053 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3052)

		_ = tmp3053

		tmp3054 := MakeNative(func(__e *ControlFlow) {
			Y945 := __e.Get(1)
			_ = Y945
			__e.Return(PrimIsVariable(Y945))
			return
		}, 1)

		tmp3055 := PrimCons(symvariable_2, tmp3054)

		tmp3056 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3055)

		_ = tmp3056

		tmp3057 := MakeNative(func(__e *ControlFlow) {
			Y940 := __e.Get(1)
			_ = Y940
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y941 := __e.Get(1)
				_ = Y941
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y942 := __e.Get(1)
					_ = Y942
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y943 := __e.Get(1)
						_ = Y943
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y944 := __e.Get(1)
							_ = Y944
							__e.TailApply(PrimFunc(symvar_2), Y940, Y941, Y942, Y943, Y944)
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

		tmp3058 := PrimCons(symvar_2, tmp3057)

		tmp3059 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3058)

		_ = tmp3059

		tmp3060 := MakeNative(func(__e *ControlFlow) {
			Y935 := __e.Get(1)
			_ = Y935
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y936 := __e.Get(1)
				_ = Y936
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Y937 := __e.Get(1)
					_ = Y937
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Y938 := __e.Get(1)
						_ = Y938
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Y939 := __e.Get(1)
							_ = Y939
							__e.TailApply(PrimFunc(symwhen), Y935, Y936, Y937, Y938, Y939)
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

		tmp3061 := PrimCons(symwhen, tmp3060)

		tmp3062 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3061)

		_ = tmp3062

		tmp3063 := MakeNative(func(__e *ControlFlow) {
			Y933 := __e.Get(1)
			_ = Y933
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y934 := __e.Get(1)
				_ = Y934
				__e.Return(PrimWriteByte(Y933, Y934))
				return
			}, 1))
			return
		}, 1)

		tmp3064 := PrimCons(symwrite_1byte, tmp3063)

		tmp3065 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3064)

		_ = tmp3065

		tmp3066 := MakeNative(func(__e *ControlFlow) {
			Y931 := __e.Get(1)
			_ = Y931
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y932 := __e.Get(1)
				_ = Y932
				__e.TailApply(PrimFunc(symwrite_1to_1file), Y931, Y932)
				return
			}, 1))
			return
		}, 1)

		tmp3067 := PrimCons(symwrite_1to_1file, tmp3066)

		tmp3068 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3067)

		_ = tmp3068

		tmp3069 := MakeNative(func(__e *ControlFlow) {
			Y930 := __e.Get(1)
			_ = Y930
			__e.TailApply(PrimFunc(symy_1or_1n_2), Y930)
			return
		}, 1)

		tmp3070 := PrimCons(symy_1or_1n_2, tmp3069)

		tmp3071 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3070)

		_ = tmp3071

		tmp3072 := MakeNative(func(__e *ControlFlow) {
			Y928 := __e.Get(1)
			_ = Y928
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y929 := __e.Get(1)
				_ = Y929
				__e.Return(PrimNumberAdd(Y928, Y929))
				return
			}, 1))
			return
		}, 1)

		tmp3073 := PrimCons(sym_7, tmp3072)

		tmp3074 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3073)

		_ = tmp3074

		tmp3075 := MakeNative(func(__e *ControlFlow) {
			Y926 := __e.Get(1)
			_ = Y926
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y927 := __e.Get(1)
				_ = Y927
				__e.Return(PrimNumberMultiply(Y926, Y927))
				return
			}, 1))
			return
		}, 1)

		tmp3076 := PrimCons(sym_d, tmp3075)

		tmp3077 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3076)

		_ = tmp3077

		tmp3078 := MakeNative(func(__e *ControlFlow) {
			Y924 := __e.Get(1)
			_ = Y924
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y925 := __e.Get(1)
				_ = Y925
				__e.Return(PrimNumberDivide(Y924, Y925))
				return
			}, 1))
			return
		}, 1)

		tmp3079 := PrimCons(sym_c, tmp3078)

		tmp3080 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3079)

		_ = tmp3080

		tmp3081 := MakeNative(func(__e *ControlFlow) {
			Y922 := __e.Get(1)
			_ = Y922
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y923 := __e.Get(1)
				_ = Y923
				__e.Return(PrimNumberSubtract(Y922, Y923))
				return
			}, 1))
			return
		}, 1)

		tmp3082 := PrimCons(sym_1, tmp3081)

		tmp3083 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3082)

		_ = tmp3083

		tmp3084 := MakeNative(func(__e *ControlFlow) {
			Y920 := __e.Get(1)
			_ = Y920
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y921 := __e.Get(1)
				_ = Y921
				__e.TailApply(PrimFunc(sym_a_a), Y920, Y921)
				return
			}, 1))
			return
		}, 1)

		tmp3085 := PrimCons(sym_a_a, tmp3084)

		tmp3086 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3085)

		_ = tmp3086

		tmp3087 := MakeNative(func(__e *ControlFlow) {
			Y919 := __e.Get(1)
			_ = Y919
			__e.TailApply(PrimFunc(sym_5e_6), Y919)
			return
		}, 1)

		tmp3088 := PrimCons(sym_5e_6, tmp3087)

		tmp3089 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3088)

		_ = tmp3089

		tmp3090 := MakeNative(func(__e *ControlFlow) {
			Y918 := __e.Get(1)
			_ = Y918
			__e.TailApply(PrimFunc(sym_5end_6), Y918)
			return
		}, 1)

		tmp3091 := PrimCons(sym_5end_6, tmp3090)

		tmp3092 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3091)

		_ = tmp3092

		tmp3093 := MakeNative(func(__e *ControlFlow) {
			Y917 := __e.Get(1)
			_ = Y917
			__e.TailApply(PrimFunc(sym_5_b_6), Y917)
			return
		}, 1)

		tmp3094 := PrimCons(sym_5_b_6, tmp3093)

		tmp3095 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3094)

		_ = tmp3095

		tmp3096 := MakeNative(func(__e *ControlFlow) {
			Y915 := __e.Get(1)
			_ = Y915
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y916 := __e.Get(1)
				_ = Y916
				__e.TailApply(PrimFunc(sym_8p), Y915, Y916)
				return
			}, 1))
			return
		}, 1)

		tmp3097 := PrimCons(sym_8p, tmp3096)

		tmp3098 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3097)

		_ = tmp3098

		tmp3099 := MakeNative(func(__e *ControlFlow) {
			Y913 := __e.Get(1)
			_ = Y913
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y914 := __e.Get(1)
				_ = Y914
				__e.TailApply(PrimFunc(sym_8v), Y913, Y914)
				return
			}, 1))
			return
		}, 1)

		tmp3100 := PrimCons(sym_8v, tmp3099)

		tmp3101 := Call(__e, PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3100)

		_ = tmp3101

		tmp3102 := MakeNative(func(__e *ControlFlow) {
			Y911 := __e.Get(1)
			_ = Y911
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Y912 := __e.Get(1)
				_ = Y912
				__e.TailApply(PrimFunc(sym_8s), Y911, Y912)
				return
			}, 1))
			return
		}, 1)

		tmp3103 := PrimCons(sym_8s, tmp3102)

		__e.TailApply(PrimFunc(symshen_4set_1lambda_1form_1entry), tmp3103)
		return

	}, 0)

	tmp3104 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1forms, tmp2561)

	_ = tmp3104

	tmp3105 := MakeNative(func(__e *ControlFlow) {
		tmp3106 := Call(__e, PrimFunc(symshen_4initialise_1environment))

		_ = tmp3106

		tmp3107 := Call(__e, PrimFunc(symshen_4initialise_1lambda_1forms))

		_ = tmp3107

		__e.TailApply(PrimFunc(symshen_4initialise_1signedfuncs))
		return

	}, 0)

	__e.TailApply(ns2_1set, symshen_4initialise, tmp3105)
	return

}, 0)

var sym_5_1address = MakeSymbol("<-address")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symdestroy = MakeSymbol("destroy")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symK = MakeSymbol("K")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4continue = MakeSymbol("shen.continue")
var symhush = MakeSymbol("hush")
var sym_dabsolute_d = MakeSymbol("*absolute*")
var symcons = MakeSymbol("cons")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symlazy = MakeSymbol("lazy")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symtc_2 = MakeSymbol("tc?")
var symdefcc = MakeSymbol("defcc")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4unix = MakeSymbol("shen.unix")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symoccurrences = MakeSymbol("occurrences")
var symParse = MakeSymbol("Parse")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4rule_1_6body = MakeSymbol("shen.rule->body")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symfactorise = MakeSymbol("factorise")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symstinput = MakeSymbol("stinput")
var symdefun = MakeSymbol("defun")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symempty_2 = MakeSymbol("empty?")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symeval = MakeSymbol("eval")
var sym_8s = MakeSymbol("@s")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symtail = MakeSymbol("tail")
var symin = MakeSymbol("in")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symoutput = MakeSymbol("output")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symspy = MakeSymbol("spy")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var sym_5 = MakeSymbol("<")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symnull = MakeSymbol("null")
var symadjoin = MakeSymbol("adjoin")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symhdstr = MakeSymbol("hdstr")
var symnl = MakeSymbol("nl")
var symporters = MakeSymbol("porters")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4raise_1syntax_1error = MakeSymbol("shen.raise-syntax-error")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symeval_1kl = MakeSymbol("eval-kl")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symdo = MakeSymbol("do")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var sym_5_b_6 = MakeSymbol("<!>")
var symcall = MakeSymbol("call")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symmake_1string = MakeSymbol("make-string")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symshen_4received = MakeSymbol("shen.received")
var sym_a_a_6 = MakeSymbol("==>")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symRecord = MakeSymbol("Record")
var symStart = MakeSymbol("Start")
var sym_8v = MakeSymbol("@v")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symassertz = MakeSymbol("assertz")
var symstring = MakeSymbol("string")
var sym_dlanguage_d = MakeSymbol("*language*")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symbound_2 = MakeSymbol("bound?")
var symcons_2 = MakeSymbol("cons?")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symprolog_1memory = MakeSymbol("prolog-memory")
var sym_e = MakeSymbol("&")
var symS = MakeSymbol("S")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symintern = MakeSymbol("intern")
var symA = MakeSymbol("A")
var symshen_4initialise = MakeSymbol("shen.initialise")
var symshen_4not_1tuple = MakeSymbol("shen.not-tuple")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symtlstr = MakeSymbol("tlstr")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symloaded = MakeSymbol("loaded")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var sym_6 = MakeSymbol(">")
var symand = MakeSymbol("and")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symunabsolute = MakeSymbol("unabsolute")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symget = MakeSymbol("get")
var sym_dhush_d = MakeSymbol("*hush*")
var symremove = MakeSymbol("remove")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var sympackage_2 = MakeSymbol("package?")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symport = MakeSymbol("port")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symprofile = MakeSymbol("profile")
var symshen_4not_1dictionary = MakeSymbol("shen.not-dictionary")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symfail = MakeSymbol("fail")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4correct = MakeSymbol("shen.correct")
var symshen_4source = MakeSymbol("shen.source")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var sym_5_1 = MakeSymbol("<-")
var symFinish = MakeSymbol("Finish")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symwarn = MakeSymbol("warn")
var symAssumptions = MakeSymbol("Assumptions")
var symfile = MakeSymbol("file")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var sym_5_a = MakeSymbol("<=")
var symhush_2 = MakeSymbol("hush?")
var symshen_4eos = MakeSymbol("shen.eos")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var sym_6_a = MakeSymbol(">=")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symor = MakeSymbol("or")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var sym_3 = MakeSymbol("$")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symAssumption = MakeSymbol("Assumption")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symverified = MakeSymbol("verified")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symsave = MakeSymbol("save")
var symprolog_2 = MakeSymbol("prolog?")
var symunput = MakeSymbol("unput")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symmode = MakeSymbol("mode")
var symshen_4premises_1_6goals = MakeSymbol("shen.premises->goals")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symnth = MakeSymbol("nth")
var symatom_2 = MakeSymbol("atom?")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var symshen_4repl = MakeSymbol("shen.repl")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symTl = MakeSymbol("Tl")
var symW = MakeSymbol("W")
var symshen_4op2 = MakeSymbol("shen.op2")
var symout = MakeSymbol("out")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symZ = MakeSymbol("Z")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symreceive = MakeSymbol("receive")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var syminput = MakeSymbol("input")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symmapcan = MakeSymbol("mapcan")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symfactorise_2 = MakeSymbol("factorise?")
var sym_dos_d = MakeSymbol("*os*")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4reader_1error = MakeSymbol("shen.reader-error")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symundefmacro = MakeSymbol("undefmacro")
var symshen_4remove_1bystanders = MakeSymbol("shen.remove-bystanders")
var sym__ = MakeSymbol("_")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symerror = MakeSymbol("error")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symretract = MakeSymbol("retract")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symRemainder = MakeSymbol("Remainder")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symstream = MakeSymbol("stream")
var symHd = MakeSymbol("Hd")
var symhdv = MakeSymbol("hdv")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4syntax_1error_1message = MakeSymbol("shen.syntax-error-message")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symctxt = MakeSymbol("ctxt")
var symshen_4decons = MakeSymbol("shen.decons")
var syminternal = MakeSymbol("internal")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4macros = MakeSymbol("shen.macros")
var sym_6_6 = MakeSymbol(">>")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symP = MakeSymbol("P")
var symshen_4macro_1_8c = MakeSymbol("shen.macro-@c")
var symshen_4specialise_1consume = MakeSymbol("shen.specialise-consume")
var sympreclude = MakeSymbol("preclude")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symoptimise_2 = MakeSymbol("optimise?")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4_8ch = MakeSymbol("shen.@ch")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symlimit = MakeSymbol("limit")
var symos = MakeSymbol("os")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symFreeze = MakeSymbol("Freeze")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symTm = MakeSymbol("Tm")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var syminferences = MakeSymbol("inferences")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symshen_4process_1cond_1clauses = MakeSymbol("shen.process-cond-clauses")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symload = MakeSymbol("load")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symtrack = MakeSymbol("track")
var symnumber = MakeSymbol("number")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var sympos = MakeSymbol("pos")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symdatatype = MakeSymbol("datatype")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4top = MakeSymbol("shen.top")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symtype = MakeSymbol("type")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var sympackage = MakeSymbol("package")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symrun = MakeSymbol("run")
var symHypotheses = MakeSymbol("Hypotheses")
var symassoc = MakeSymbol("assoc")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symis = MakeSymbol("is")
var sym_i = MakeSymbol("{")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symlist = MakeSymbol("list")
var symB = MakeSymbol("B")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symwhere = MakeSymbol("where")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4macro_1_8ch = MakeSymbol("shen.macro-@ch")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4magless = MakeSymbol("shen.magless")
var symstr = MakeSymbol("str")
var symshen_4partial = MakeSymbol("shen.partial")
var symopen = MakeSymbol("open")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symasserta = MakeSymbol("asserta")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var sym_j = MakeSymbol("}")
var symversion = MakeSymbol("version")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symprint = MakeSymbol("print")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symhash = MakeSymbol("hash")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4g = MakeSymbol("shen.g")
var symprotect = MakeSymbol("protect")
var symfn = MakeSymbol("fn")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symfunction = MakeSymbol("function")
var symappend = MakeSymbol("append")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var sym_dsterror_d = MakeSymbol("*sterror*")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4mod = MakeSymbol("shen.mod")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4line = MakeSymbol("shen.line")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symshen_4_1m = MakeSymbol("shen.-m")
var symreturn = MakeSymbol("return")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symunion = MakeSymbol("union")
var symget_1time = MakeSymbol("get-time")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4beta = MakeSymbol("shen.beta")
var symfix = MakeSymbol("fix")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symcompile = MakeSymbol("compile")
var symNewV = MakeSymbol("NewV")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symdifference = MakeSymbol("difference")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var sym_c_4 = MakeSymbol("/.")
var symtime = MakeSymbol("time")
var symC = MakeSymbol("C")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4comb = MakeSymbol("shen.comb")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symis_b = MakeSymbol("is!")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symcd = MakeSymbol("cd")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symforeign = MakeSymbol("foreign")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symKey = MakeSymbol("Key")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4nvars = MakeSymbol("shen.nvars")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4special_2 = MakeSymbol("shen.special?")
var sym_1 = MakeSymbol("-")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symlet = MakeSymbol("let")
var symlambda = MakeSymbol("lambda")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symboolean = MakeSymbol("boolean")
var symsum = MakeSymbol("sum")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4not_1pvar = MakeSymbol("shen.not-pvar")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var sym_7 = MakeSymbol("+")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4t = MakeSymbol("shen.t")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symY = MakeSymbol("Y")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symmacroexpand = MakeSymbol("macroexpand")
var symlength = MakeSymbol("length")
var symdeclare = MakeSymbol("declare")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symcases = MakeSymbol("cases")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var sym_5_1vector = MakeSymbol("<-vector")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symnumber_2 = MakeSymbol("number?")
var symX = MakeSymbol("X")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var sym_8p = MakeSymbol("@p")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symGoTo = MakeSymbol("GoTo")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symfindall = MakeSymbol("findall")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symmap = MakeSymbol("map")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symspecialise = MakeSymbol("specialise")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symvalue = MakeSymbol("value")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symexplode = MakeSymbol("explode")
var sym_dporters_d = MakeSymbol("*porters*")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symunit = MakeSymbol("unit")
var symuserdefs = MakeSymbol("userdefs")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symshen_4specialise_1member = MakeSymbol("shen.specialise-member")
var symexception = MakeSymbol("exception")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var sym_e_e = MakeSymbol("&&")
var symsymbol_2 = MakeSymbol("symbol?")
var symsubst = MakeSymbol("subst")
var symsterror = MakeSymbol("sterror")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4step_2 = MakeSymbol("shen.step?")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4tame = MakeSymbol("shen.tame")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symwhen = MakeSymbol("when")
var symhead = MakeSymbol("head")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symput = MakeSymbol("put")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symoccurs_2 = MakeSymbol("occurs?")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symvector = MakeSymbol("vector")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var sym_dport_d = MakeSymbol("*port*")
var symL = MakeSymbol("L")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symAction = MakeSymbol("Action")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symtuple_2 = MakeSymbol("tuple?")
var sym_a_a = MakeSymbol("==")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4member_1clause = MakeSymbol("shen.member-clause")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symabort = MakeSymbol("abort")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4consume = MakeSymbol("shen.consume")
var sympr = MakeSymbol("pr")
var symimplementation = MakeSymbol("implementation")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var syminline = MakeSymbol("inline")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4ctxt = MakeSymbol("shen.ctxt")
var symfst = MakeSymbol("fst")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4_dempty_1absvector_d = MakeSymbol("shen.*empty-absvector*")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symclose = MakeSymbol("close")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4dict_1values = MakeSymbol("shen.dict-values")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symResult = MakeSymbol("Result")
var symit = MakeSymbol("it")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4included = MakeSymbol("shen.included")
var symshen_4_8c = MakeSymbol("shen.@c")
var symu_b = MakeSymbol("u!")
var symshen_4member = MakeSymbol("shen.member")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var symSelect = MakeSymbol("Select")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symMessage = MakeSymbol("Message")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symsymbol = MakeSymbol("symbol")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symtc = MakeSymbol("tc")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symshen_4demode = MakeSymbol("shen.demode")
var symvector_2 = MakeSymbol("vector?")
var symconcat = MakeSymbol("concat")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4passive_1bind = MakeSymbol("shen.passive-bind")
var symreverse = MakeSymbol("reverse")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symrelease = MakeSymbol("release")
var sym_4_4_4 = MakeSymbol("...")
var symabsolute = MakeSymbol("absolute")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symhd = MakeSymbol("hd")
var symshen_4reader_1error_1message = MakeSymbol("shen.reader-error-message")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4cons_1form_1respect_1modes = MakeSymbol("shen.cons-form-respect-modes")
var sym_d = MakeSymbol("*")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var sym_5e_6 = MakeSymbol("<e>")
var symcond = MakeSymbol("cond")
var symprofile_1results = MakeSymbol("profile-results")
var symsnd = MakeSymbol("snd")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symfork = MakeSymbol("fork")
var symshen_4passive_1variables = MakeSymbol("shen.passive-variables")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symnot = MakeSymbol("not")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symvar_2 = MakeSymbol("var?")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4processed = MakeSymbol("shen.processed")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4dict_2 = MakeSymbol("shen.dict?")
var symabsvector = MakeSymbol("absvector")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4rule_1_6head = MakeSymbol("shen.rule->head")
var symshen_4abs = MakeSymbol("shen.abs")
var symtl = MakeSymbol("tl")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symlineread = MakeSymbol("lineread")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symsystemf = MakeSymbol("systemf")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symcn = MakeSymbol("cn")
var symfreeze = MakeSymbol("freeze")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symbind = MakeSymbol("bind")
var symshen_4dict_1rm = MakeSymbol("shen.dict-rm")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4app = MakeSymbol("shen.app")
var symvariable_2 = MakeSymbol("variable?")
var symfresh = MakeSymbol("fresh")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symdatatypes = MakeSymbol("datatypes")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symtracked = MakeSymbol("tracked")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symshen_4factor = MakeSymbol("shen.factor")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4goto = MakeSymbol("shen.goto")
var symuntrack = MakeSymbol("untrack")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symshen = MakeSymbol("shen")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symthaw = MakeSymbol("thaw")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4partial_1parse_1failure_2 = MakeSymbol("shen.partial-parse-failure?")
var symshen_4predicate = MakeSymbol("shen.predicate")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symincluded = MakeSymbol("included")
var symgensym = MakeSymbol("gensym")
var symintersection = MakeSymbol("intersection")
var symshen_4f = MakeSymbol("shen.f")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var sym_a = MakeSymbol("=")
var sym_5_1_1 = MakeSymbol("<--")
var syminclude = MakeSymbol("include")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symexternal = MakeSymbol("external")
var symin_1package = MakeSymbol("in-package")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symshen_4side_1conditions_1_6goals = MakeSymbol("shen.side-conditions->goals")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var syminput_7 = MakeSymbol("input+")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symps = MakeSymbol("ps")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symNewAssumptions = MakeSymbol("NewAssumptions")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symread_1file = MakeSymbol("read-file")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symwrite_1byte = MakeSymbol("write-byte")
var symtlv = MakeSymbol("tlv")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var sym_c = MakeSymbol("/")
var symlanguage = MakeSymbol("language")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symdefmacro = MakeSymbol("defmacro")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symstep = MakeSymbol("step")
var symstring_2 = MakeSymbol("string?")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symbar_b = MakeSymbol("bar!")
var symset = MakeSymbol("set")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symV = MakeSymbol("V")
var sym_b = MakeSymbol("!")
var symshen_4dict_1keys = MakeSymbol("shen.dict-keys")
var symshen_4op1 = MakeSymbol("shen.op1")
var symTime = MakeSymbol("Time")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symdefine = MakeSymbol("define")
var symif = MakeSymbol("if")
var sym_5end_6 = MakeSymbol("<end>")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symshen_4typename = MakeSymbol("shen.typename")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symaddress_1_6 = MakeSymbol("address->")
var symelement_2 = MakeSymbol("element?")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4_ddemodulation_1function_d = MakeSymbol("shen.*demodulation-function*")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symoptimise = MakeSymbol("optimise")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symunprofile = MakeSymbol("unprofile")
var symread = MakeSymbol("read")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symstoutput = MakeSymbol("stoutput")
var sym_1_6 = MakeSymbol("->")
var symshen_4consume_1clause = MakeSymbol("shen.consume-clause")
var symshen_4spy_2 = MakeSymbol("shen.spy?")
var symarity = MakeSymbol("arity")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4loop = MakeSymbol("shen.loop")
