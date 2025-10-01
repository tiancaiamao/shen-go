package main

import . "github.com/tiancaiamao/shen-go/kl"

var MacrosMain = MakeNative(func(__e *ControlFlow) {
tmp8149 := MakeNative(func(__e *ControlFlow) {
V1786 := __e.Get(1)
_ = V1786
tmp8150 := MakeNative(func(__e *ControlFlow) {
Fs := __e.Get(1)
_ = Fs
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V1786, Fs, Fs)
return
}, 1)

tmp8151 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(PrimTail(X))
return
}, 1)

tmp8152 := PrimValue(sym_dmacros_d)

tmp8153 := Call(__e, PrimFunc(symmap), tmp8151, tmp8152)


__e.TailApply(tmp8150, tmp8153)
return


}, 1)

tmp8154 := Call(__e, ns2_1set, symmacroexpand, tmp8149)


_ = tmp8154

tmp8155 := MakeNative(func(__e *ControlFlow) {
V1795 := __e.Get(1)
_ = V1795
V1796 := __e.Get(2)
_ = V1796
V1797 := __e.Get(3)
_ = V1797
tmp8165 := PrimEqual(Nil, V1796)

if True == tmp8165 {
__e.Return(V1795)
return
} else {
tmp8163 := PrimIsPair(V1796)

if True == tmp8163 {
tmp8156 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp8159 := PrimEqual(V1795, Y)

if True == tmp8159 {
tmp8157 := PrimTail(V1796)

__e.TailApply(PrimFunc(symshen_4macroexpand_1h), V1795, tmp8157, V1797)
return


} else {
__e.TailApply(PrimFunc(symshen_4macroexpand_1h), Y, V1797, V1797)
return
}


}, 1)

tmp8160 := PrimHead(V1796)

tmp8161 := Call(__e, PrimFunc(symshen_4walk), tmp8160, V1795)


__e.TailApply(tmp8156, tmp8161)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.macroexpand-h")))
return
}


}


}, 3)

tmp8166 := Call(__e, ns2_1set, symshen_4macroexpand_1h, tmp8155)


_ = tmp8166

tmp8167 := MakeNative(func(__e *ControlFlow) {
V1798 := __e.Get(1)
_ = V1798
V1799 := __e.Get(2)
_ = V1799
tmp8171 := PrimIsPair(V1799)

if True == tmp8171 {
tmp8168 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
__e.TailApply(PrimFunc(symshen_4walk), V1798, Z)
return
}, 1)

tmp8169 := Call(__e, PrimFunc(symmap), tmp8168, V1799)


__e.TailApply(V1798, tmp8169)
return


} else {
__e.TailApply(V1798, V1799)
return
}


}, 2)

tmp8172 := Call(__e, ns2_1set, symshen_4walk, tmp8167)


_ = tmp8172

tmp8173 := MakeNative(func(__e *ControlFlow) {
V1800 := __e.Get(1)
_ = V1800
tmp8206 := PrimIsPair(V1800)

var ifres8198 Obj

if True == tmp8206 {
tmp8204 := PrimHead(V1800)

tmp8205 := PrimEqual(symdefmacro, tmp8204)

var ifres8200 Obj

if True == tmp8205 {
tmp8202 := PrimTail(V1800)

tmp8203 := PrimIsPair(tmp8202)

var ifres8201 Obj

if True == tmp8203 {
ifres8201 = True


} else {
ifres8201 = False


}

ifres8200 = ifres8201


} else {
ifres8200 = False


}

var ifres8199 Obj

if True == ifres8200 {
ifres8199 = True


} else {
ifres8199 = False


}

ifres8198 = ifres8199


} else {
ifres8198 = False


}

if True == ifres8198 {
tmp8174 := MakeNative(func(__e *ControlFlow) {
Default := __e.Get(1)
_ = Default
tmp8175 := MakeNative(func(__e *ControlFlow) {
Def := __e.Get(1)
_ = Def
tmp8176 := MakeNative(func(__e *ControlFlow) {
Record := __e.Get(1)
_ = Record
tmp8177 := PrimTail(V1800)

__e.Return(PrimHead(tmp8177))
return


}, 1)

tmp8178 := PrimTail(V1800)

tmp8179 := PrimHead(tmp8178)

tmp8180 := PrimTail(V1800)

tmp8181 := PrimHead(tmp8180)

tmp8182 := Call(__e, PrimFunc(symfn), tmp8181)


tmp8183 := Call(__e, PrimFunc(symshen_4record_1macro), tmp8179, tmp8182)


__e.TailApply(tmp8176, tmp8183)
return


}, 1)

tmp8184 := PrimTail(V1800)

tmp8185 := PrimHead(tmp8184)

tmp8186 := PrimTail(V1800)

tmp8187 := PrimTail(tmp8186)

tmp8188 := Call(__e, PrimFunc(symappend), tmp8187, Default)


tmp8189 := PrimCons(tmp8185, tmp8188)

tmp8190 := PrimCons(symdefine, tmp8189)

tmp8191 := Call(__e, PrimFunc(symeval), tmp8190)


__e.TailApply(tmp8175, tmp8191)
return


}, 1)

tmp8192 := Call(__e, PrimFunc(symprotect), symX)


tmp8193 := Call(__e, PrimFunc(symprotect), symX)


tmp8194 := PrimCons(tmp8193, Nil)

tmp8195 := PrimCons(sym_1_6, tmp8194)

tmp8196 := PrimCons(tmp8192, tmp8195)

__e.TailApply(tmp8174, tmp8196)
return


} else {
__e.Return(V1800)
return
}


}, 1)

tmp8207 := Call(__e, ns2_1set, symshen_4defmacro_1macro, tmp8173)


_ = tmp8207

tmp8208 := MakeNative(func(__e *ControlFlow) {
V1801 := __e.Get(1)
_ = V1801
tmp8227 := PrimIsPair(V1801)

var ifres8214 Obj

if True == tmp8227 {
tmp8225 := PrimHead(V1801)

tmp8226 := PrimEqual(symu_b, tmp8225)

var ifres8216 Obj

if True == tmp8226 {
tmp8223 := PrimTail(V1801)

tmp8224 := PrimIsPair(tmp8223)

var ifres8218 Obj

if True == tmp8224 {
tmp8220 := PrimTail(V1801)

tmp8221 := PrimTail(tmp8220)

tmp8222 := PrimEqual(Nil, tmp8221)

var ifres8219 Obj

if True == tmp8222 {
ifres8219 = True


} else {
ifres8219 = False


}

ifres8218 = ifres8219


} else {
ifres8218 = False


}

var ifres8217 Obj

if True == ifres8218 {
ifres8217 = True


} else {
ifres8217 = False


}

ifres8216 = ifres8217


} else {
ifres8216 = False


}

var ifres8215 Obj

if True == ifres8216 {
ifres8215 = True


} else {
ifres8215 = False


}

ifres8214 = ifres8215


} else {
ifres8214 = False


}

if True == ifres8214 {
tmp8209 := PrimTail(V1801)

tmp8210 := PrimHead(tmp8209)

tmp8211 := Call(__e, PrimFunc(symshen_4make_1uppercase), tmp8210)


tmp8212 := PrimCons(tmp8211, Nil)

__e.Return(PrimCons(symprotect, tmp8212))
return


} else {
__e.Return(V1801)
return
}


}, 1)

tmp8228 := Call(__e, ns2_1set, symshen_4u_b_1macro, tmp8208)


_ = tmp8228

tmp8229 := MakeNative(func(__e *ControlFlow) {
V1802 := __e.Get(1)
_ = V1802
tmp8230 := PrimStr(V1802)

tmp8231 := Call(__e, PrimFunc(symshen_4mu_1h), tmp8230)


__e.Return(PrimIntern(tmp8231))
return


}, 1)

tmp8232 := Call(__e, ns2_1set, symshen_4make_1uppercase, tmp8229)


_ = tmp8232

tmp8233 := MakeNative(func(__e *ControlFlow) {
V1803 := __e.Get(1)
_ = V1803
tmp8252 := PrimEqual(MakeString(""), V1803)

if True == tmp8252 {
__e.Return(MakeString(""))
return
} else {
tmp8250 := Call(__e, PrimFunc(symshen_4_7string_2), V1803)


if True == tmp8250 {
tmp8234 := MakeNative(func(__e *ControlFlow) {
ASCII := __e.Get(1)
_ = ASCII
tmp8235 := MakeNative(func(__e *ControlFlow) {
ASCII_132 := __e.Get(1)
_ = ASCII_132
tmp8236 := MakeNative(func(__e *ControlFlow) {
Upper := __e.Get(1)
_ = Upper
tmp8237 := PrimTailString(V1803)

tmp8238 := Call(__e, PrimFunc(symshen_4mu_1h), tmp8237)


__e.TailApply(PrimFunc(sym_8s), Upper, tmp8238)
return


}, 1)

tmp8245 := PrimGreatEqual(ASCII, MakeNumber(97))

var ifres8242 Obj

if True == tmp8245 {
tmp8244 := PrimLessEqual(ASCII, MakeNumber(122))

var ifres8243 Obj

if True == tmp8244 {
ifres8243 = True


} else {
ifres8243 = False


}

ifres8242 = ifres8243


} else {
ifres8242 = False


}

var ifres8239 Obj

if True == ifres8242 {
tmp8240 := PrimNumberToString(ASCII_132)

ifres8239 = tmp8240


} else {
tmp8241 := Call(__e, PrimFunc(symhdstr), V1803)


ifres8239 = tmp8241


}

__e.TailApply(tmp8236, ifres8239)
return


}, 1)

tmp8246 := PrimNumberSubtract(ASCII, MakeNumber(32))

__e.TailApply(tmp8235, tmp8246)
return


}, 1)

tmp8247 := Call(__e, PrimFunc(symhdstr), V1803)


tmp8248 := PrimStringToNumber(tmp8247)

__e.TailApply(tmp8234, tmp8248)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4mu_1h)
return
}


}


}, 1)

tmp8253 := Call(__e, ns2_1set, symshen_4mu_1h, tmp8233)


_ = tmp8253

tmp8254 := MakeNative(func(__e *ControlFlow) {
V1804 := __e.Get(1)
_ = V1804
V1805 := __e.Get(2)
_ = V1805
tmp8255 := PrimValue(sym_dmacros_d)

tmp8256 := Call(__e, PrimFunc(symshen_4update_1assoc), V1804, V1805, tmp8255)


__e.Return(PrimSet(sym_dmacros_d, tmp8256))
return


}, 2)

tmp8257 := Call(__e, ns2_1set, symshen_4record_1macro, tmp8254)


_ = tmp8257

tmp8258 := MakeNative(func(__e *ControlFlow) {
V1815 := __e.Get(1)
_ = V1815
V1816 := __e.Get(2)
_ = V1816
V1817 := __e.Get(3)
_ = V1817
tmp8281 := PrimEqual(Nil, V1817)

if True == tmp8281 {
tmp8259 := PrimCons(V1815, V1816)

__e.Return(PrimCons(tmp8259, Nil))
return


} else {
tmp8279 := PrimIsPair(V1817)

var ifres8270 Obj

if True == tmp8279 {
tmp8277 := PrimHead(V1817)

tmp8278 := PrimIsPair(tmp8277)

var ifres8272 Obj

if True == tmp8278 {
tmp8274 := PrimHead(V1817)

tmp8275 := PrimHead(tmp8274)

tmp8276 := PrimEqual(V1815, tmp8275)

var ifres8273 Obj

if True == tmp8276 {
ifres8273 = True


} else {
ifres8273 = False


}

ifres8272 = ifres8273


} else {
ifres8272 = False


}

var ifres8271 Obj

if True == ifres8272 {
ifres8271 = True


} else {
ifres8271 = False


}

ifres8270 = ifres8271


} else {
ifres8270 = False


}

if True == ifres8270 {
tmp8260 := PrimHead(V1817)

tmp8261 := PrimHead(tmp8260)

tmp8262 := PrimCons(tmp8261, V1816)

tmp8263 := PrimTail(V1817)

__e.Return(PrimCons(tmp8262, tmp8263))
return


} else {
tmp8268 := PrimIsPair(V1817)

if True == tmp8268 {
tmp8264 := PrimHead(V1817)

tmp8265 := PrimTail(V1817)

tmp8266 := Call(__e, PrimFunc(symshen_4update_1assoc), V1815, V1816, tmp8265)


__e.Return(PrimCons(tmp8264, tmp8266))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.update-assoc")))
return
}


}


}


}, 3)

tmp8282 := Call(__e, ns2_1set, symshen_4update_1assoc, tmp8258)


_ = tmp8282

tmp8283 := MakeNative(func(__e *ControlFlow) {
V1818 := __e.Get(1)
_ = V1818
tmp8299 := PrimIsPair(V1818)

var ifres8291 Obj

if True == tmp8299 {
tmp8297 := PrimHead(V1818)

tmp8298 := PrimEqual(symerror, tmp8297)

var ifres8293 Obj

if True == tmp8298 {
tmp8295 := PrimTail(V1818)

tmp8296 := PrimIsPair(tmp8295)

var ifres8294 Obj

if True == tmp8296 {
ifres8294 = True


} else {
ifres8294 = False


}

ifres8293 = ifres8294


} else {
ifres8293 = False


}

var ifres8292 Obj

if True == ifres8293 {
ifres8292 = True


} else {
ifres8292 = False


}

ifres8291 = ifres8292


} else {
ifres8291 = False


}

if True == ifres8291 {
tmp8284 := PrimTail(V1818)

tmp8285 := PrimHead(tmp8284)

tmp8286 := PrimTail(V1818)

tmp8287 := PrimTail(tmp8286)

tmp8288 := Call(__e, PrimFunc(symshen_4mkstr), tmp8285, tmp8287)


tmp8289 := PrimCons(tmp8288, Nil)

__e.Return(PrimCons(symsimple_1error, tmp8289))
return


} else {
__e.Return(V1818)
return
}


}, 1)

tmp8300 := Call(__e, ns2_1set, symshen_4error_1macro, tmp8283)


_ = tmp8300

tmp8301 := MakeNative(func(__e *ControlFlow) {
V1819 := __e.Get(1)
_ = V1819
tmp8339 := PrimIsPair(V1819)

var ifres8331 Obj

if True == tmp8339 {
tmp8337 := PrimHead(V1819)

tmp8338 := PrimEqual(symoutput, tmp8337)

var ifres8333 Obj

if True == tmp8338 {
tmp8335 := PrimTail(V1819)

tmp8336 := PrimIsPair(tmp8335)

var ifres8334 Obj

if True == tmp8336 {
ifres8334 = True


} else {
ifres8334 = False


}

ifres8333 = ifres8334


} else {
ifres8333 = False


}

var ifres8332 Obj

if True == ifres8333 {
ifres8332 = True


} else {
ifres8332 = False


}

ifres8331 = ifres8332


} else {
ifres8331 = False


}

if True == ifres8331 {
tmp8302 := PrimTail(V1819)

tmp8303 := PrimHead(tmp8302)

tmp8304 := PrimTail(V1819)

tmp8305 := PrimTail(tmp8304)

tmp8306 := Call(__e, PrimFunc(symshen_4mkstr), tmp8303, tmp8305)


tmp8307 := PrimCons(symstoutput, Nil)

tmp8308 := PrimCons(tmp8307, Nil)

tmp8309 := PrimCons(tmp8306, tmp8308)

__e.Return(PrimCons(sympr, tmp8309))
return


} else {
tmp8329 := PrimIsPair(V1819)

var ifres8316 Obj

if True == tmp8329 {
tmp8327 := PrimHead(V1819)

tmp8328 := PrimEqual(sympr, tmp8327)

var ifres8318 Obj

if True == tmp8328 {
tmp8325 := PrimTail(V1819)

tmp8326 := PrimIsPair(tmp8325)

var ifres8320 Obj

if True == tmp8326 {
tmp8322 := PrimTail(V1819)

tmp8323 := PrimTail(tmp8322)

tmp8324 := PrimEqual(Nil, tmp8323)

var ifres8321 Obj

if True == tmp8324 {
ifres8321 = True


} else {
ifres8321 = False


}

ifres8320 = ifres8321


} else {
ifres8320 = False


}

var ifres8319 Obj

if True == ifres8320 {
ifres8319 = True


} else {
ifres8319 = False


}

ifres8318 = ifres8319


} else {
ifres8318 = False


}

var ifres8317 Obj

if True == ifres8318 {
ifres8317 = True


} else {
ifres8317 = False


}

ifres8316 = ifres8317


} else {
ifres8316 = False


}

if True == ifres8316 {
tmp8310 := PrimTail(V1819)

tmp8311 := PrimHead(tmp8310)

tmp8312 := PrimCons(symstoutput, Nil)

tmp8313 := PrimCons(tmp8312, Nil)

tmp8314 := PrimCons(tmp8311, tmp8313)

__e.Return(PrimCons(sympr, tmp8314))
return


} else {
__e.Return(V1819)
return
}


}


}, 1)

tmp8340 := Call(__e, ns2_1set, symshen_4output_1macro, tmp8301)


_ = tmp8340

tmp8341 := MakeNative(func(__e *ControlFlow) {
V1820 := __e.Get(1)
_ = V1820
tmp8355 := PrimIsPair(V1820)

var ifres8347 Obj

if True == tmp8355 {
tmp8353 := PrimHead(V1820)

tmp8354 := PrimEqual(symmake_1string, tmp8353)

var ifres8349 Obj

if True == tmp8354 {
tmp8351 := PrimTail(V1820)

tmp8352 := PrimIsPair(tmp8351)

var ifres8350 Obj

if True == tmp8352 {
ifres8350 = True


} else {
ifres8350 = False


}

ifres8349 = ifres8350


} else {
ifres8349 = False


}

var ifres8348 Obj

if True == ifres8349 {
ifres8348 = True


} else {
ifres8348 = False


}

ifres8347 = ifres8348


} else {
ifres8347 = False


}

if True == ifres8347 {
tmp8342 := PrimTail(V1820)

tmp8343 := PrimHead(tmp8342)

tmp8344 := PrimTail(V1820)

tmp8345 := PrimTail(tmp8344)

__e.TailApply(PrimFunc(symshen_4mkstr), tmp8343, tmp8345)
return


} else {
__e.Return(V1820)
return
}


}, 1)

tmp8356 := Call(__e, ns2_1set, symshen_4make_1string_1macro, tmp8341)


_ = tmp8356

tmp8357 := MakeNative(func(__e *ControlFlow) {
V1821 := __e.Get(1)
_ = V1821
tmp8432 := PrimIsPair(V1821)

var ifres8424 Obj

if True == tmp8432 {
tmp8430 := PrimHead(V1821)

tmp8431 := PrimEqual(symlineread, tmp8430)

var ifres8426 Obj

if True == tmp8431 {
tmp8428 := PrimTail(V1821)

tmp8429 := PrimEqual(Nil, tmp8428)

var ifres8427 Obj

if True == tmp8429 {
ifres8427 = True


} else {
ifres8427 = False


}

ifres8426 = ifres8427


} else {
ifres8426 = False


}

var ifres8425 Obj

if True == ifres8426 {
ifres8425 = True


} else {
ifres8425 = False


}

ifres8424 = ifres8425


} else {
ifres8424 = False


}

if True == ifres8424 {
tmp8358 := PrimCons(symstinput, Nil)

tmp8359 := PrimCons(tmp8358, Nil)

__e.Return(PrimCons(symlineread, tmp8359))
return


} else {
tmp8422 := PrimIsPair(V1821)

var ifres8414 Obj

if True == tmp8422 {
tmp8420 := PrimHead(V1821)

tmp8421 := PrimEqual(syminput, tmp8420)

var ifres8416 Obj

if True == tmp8421 {
tmp8418 := PrimTail(V1821)

tmp8419 := PrimEqual(Nil, tmp8418)

var ifres8417 Obj

if True == tmp8419 {
ifres8417 = True


} else {
ifres8417 = False


}

ifres8416 = ifres8417


} else {
ifres8416 = False


}

var ifres8415 Obj

if True == ifres8416 {
ifres8415 = True


} else {
ifres8415 = False


}

ifres8414 = ifres8415


} else {
ifres8414 = False


}

if True == ifres8414 {
tmp8360 := PrimCons(symstinput, Nil)

tmp8361 := PrimCons(tmp8360, Nil)

__e.Return(PrimCons(syminput, tmp8361))
return


} else {
tmp8412 := PrimIsPair(V1821)

var ifres8404 Obj

if True == tmp8412 {
tmp8410 := PrimHead(V1821)

tmp8411 := PrimEqual(symread, tmp8410)

var ifres8406 Obj

if True == tmp8411 {
tmp8408 := PrimTail(V1821)

tmp8409 := PrimEqual(Nil, tmp8408)

var ifres8407 Obj

if True == tmp8409 {
ifres8407 = True


} else {
ifres8407 = False


}

ifres8406 = ifres8407


} else {
ifres8406 = False


}

var ifres8405 Obj

if True == ifres8406 {
ifres8405 = True


} else {
ifres8405 = False


}

ifres8404 = ifres8405


} else {
ifres8404 = False


}

if True == ifres8404 {
tmp8362 := PrimCons(symstinput, Nil)

tmp8363 := PrimCons(tmp8362, Nil)

__e.Return(PrimCons(symread, tmp8363))
return


} else {
tmp8402 := PrimIsPair(V1821)

var ifres8389 Obj

if True == tmp8402 {
tmp8400 := PrimHead(V1821)

tmp8401 := PrimEqual(syminput_7, tmp8400)

var ifres8391 Obj

if True == tmp8401 {
tmp8398 := PrimTail(V1821)

tmp8399 := PrimIsPair(tmp8398)

var ifres8393 Obj

if True == tmp8399 {
tmp8395 := PrimTail(V1821)

tmp8396 := PrimTail(tmp8395)

tmp8397 := PrimEqual(Nil, tmp8396)

var ifres8394 Obj

if True == tmp8397 {
ifres8394 = True


} else {
ifres8394 = False


}

ifres8393 = ifres8394


} else {
ifres8393 = False


}

var ifres8392 Obj

if True == ifres8393 {
ifres8392 = True


} else {
ifres8392 = False


}

ifres8391 = ifres8392


} else {
ifres8391 = False


}

var ifres8390 Obj

if True == ifres8391 {
ifres8390 = True


} else {
ifres8390 = False


}

ifres8389 = ifres8390


} else {
ifres8389 = False


}

if True == ifres8389 {
tmp8364 := PrimTail(V1821)

tmp8365 := PrimHead(tmp8364)

tmp8366 := PrimCons(symstinput, Nil)

tmp8367 := PrimCons(tmp8366, Nil)

tmp8368 := PrimCons(tmp8365, tmp8367)

__e.Return(PrimCons(syminput_7, tmp8368))
return


} else {
tmp8387 := PrimIsPair(V1821)

var ifres8379 Obj

if True == tmp8387 {
tmp8385 := PrimHead(V1821)

tmp8386 := PrimEqual(symread_1byte, tmp8385)

var ifres8381 Obj

if True == tmp8386 {
tmp8383 := PrimTail(V1821)

tmp8384 := PrimEqual(Nil, tmp8383)

var ifres8382 Obj

if True == tmp8384 {
ifres8382 = True


} else {
ifres8382 = False


}

ifres8381 = ifres8382


} else {
ifres8381 = False


}

var ifres8380 Obj

if True == ifres8381 {
ifres8380 = True


} else {
ifres8380 = False


}

ifres8379 = ifres8380


} else {
ifres8379 = False


}

if True == ifres8379 {
tmp8376 := Call(__e, PrimFunc(symstinput))


tmp8377 := Call(__e, PrimFunc(symshen_4char_1stinput_2), tmp8376)


if True == tmp8377 {
tmp8369 := PrimCons(symstinput, Nil)

tmp8370 := PrimCons(tmp8369, Nil)

tmp8371 := PrimCons(symshen_4read_1unit_1string, tmp8370)

tmp8372 := PrimCons(tmp8371, Nil)

__e.Return(PrimCons(symstring_1_6n, tmp8372))
return


} else {
tmp8373 := PrimCons(symstinput, Nil)

tmp8374 := PrimCons(tmp8373, Nil)

__e.Return(PrimCons(symread_1byte, tmp8374))
return


}


} else {
__e.Return(V1821)
return
}


}


}


}


}


}, 1)

tmp8433 := Call(__e, ns2_1set, symshen_4input_1macro, tmp8357)


_ = tmp8433

tmp8434 := MakeNative(func(__e *ControlFlow) {
V1822 := __e.Get(1)
_ = V1822
tmp8441 := PrimIsPair(V1822)

var ifres8437 Obj

if True == tmp8441 {
tmp8439 := PrimHead(V1822)

tmp8440 := PrimEqual(symdefcc, tmp8439)

var ifres8438 Obj

if True == tmp8440 {
ifres8438 = True


} else {
ifres8438 = False


}

ifres8437 = ifres8438


} else {
ifres8437 = False


}

if True == ifres8437 {
tmp8435 := PrimTail(V1822)

__e.TailApply(PrimFunc(symshen_4yacc_1_6shen), tmp8435)
return


} else {
__e.Return(V1822)
return
}


}, 1)

tmp8442 := Call(__e, ns2_1set, symshen_4defcc_1macro, tmp8434)


_ = tmp8442

tmp8443 := MakeNative(func(__e *ControlFlow) {
V1823 := __e.Get(1)
_ = V1823
tmp8450 := PrimIsPair(V1823)

var ifres8446 Obj

if True == tmp8450 {
tmp8448 := PrimHead(V1823)

tmp8449 := PrimEqual(symprolog_2, tmp8448)

var ifres8447 Obj

if True == tmp8449 {
ifres8447 = True


} else {
ifres8447 = False


}

ifres8446 = ifres8447


} else {
ifres8446 = False


}

if True == ifres8446 {
tmp8444 := PrimTail(V1823)

__e.TailApply(PrimFunc(symshen_4call_1prolog), tmp8444)
return


} else {
__e.Return(V1823)
return
}


}, 1)

tmp8451 := Call(__e, ns2_1set, symshen_4prolog_1macro, tmp8443)


_ = tmp8451

tmp8452 := MakeNative(func(__e *ControlFlow) {
V1824 := __e.Get(1)
_ = V1824
tmp8453 := MakeNative(func(__e *ControlFlow) {
Bindings := __e.Get(1)
_ = Bindings
tmp8454 := MakeNative(func(__e *ControlFlow) {
Lock := __e.Get(1)
_ = Lock
tmp8455 := MakeNative(func(__e *ControlFlow) {
Key := __e.Get(1)
_ = Key
tmp8456 := MakeNative(func(__e *ControlFlow) {
Continuation := __e.Get(1)
_ = Continuation
tmp8457 := MakeNative(func(__e *ControlFlow) {
CLiterals := __e.Get(1)
_ = CLiterals
tmp8458 := MakeNative(func(__e *ControlFlow) {
Received := __e.Get(1)
_ = Received
tmp8459 := MakeNative(func(__e *ControlFlow) {
B := __e.Get(1)
_ = B
tmp8460 := MakeNative(func(__e *ControlFlow) {
L := __e.Get(1)
_ = L
tmp8461 := MakeNative(func(__e *ControlFlow) {
K := __e.Get(1)
_ = K
tmp8462 := MakeNative(func(__e *ControlFlow) {
C := __e.Get(1)
_ = C
tmp8463 := MakeNative(func(__e *ControlFlow) {
Lambda := __e.Get(1)
_ = Lambda
tmp8464 := PrimCons(Continuation, Nil)

tmp8465 := PrimCons(Key, tmp8464)

tmp8466 := PrimCons(Lock, tmp8465)

tmp8467 := PrimCons(Bindings, tmp8466)

__e.Return(PrimCons(Lambda, tmp8467))
return


}, 1)

tmp8468 := Call(__e, PrimFunc(symshen_4continue), Received, CLiterals, B, L, K, C)


tmp8469 := PrimCons(tmp8468, Nil)

tmp8470 := PrimCons(C, tmp8469)

tmp8471 := PrimCons(symlambda, tmp8470)

tmp8472 := PrimCons(tmp8471, Nil)

tmp8473 := PrimCons(K, tmp8472)

tmp8474 := PrimCons(symlambda, tmp8473)

tmp8475 := PrimCons(tmp8474, Nil)

tmp8476 := PrimCons(L, tmp8475)

tmp8477 := PrimCons(symlambda, tmp8476)

tmp8478 := PrimCons(tmp8477, Nil)

tmp8479 := PrimCons(B, tmp8478)

tmp8480 := PrimCons(symlambda, tmp8479)

__e.TailApply(tmp8463, tmp8480)
return


}, 1)

tmp8481 := Call(__e, PrimFunc(symprotect), symC)


tmp8482 := Call(__e, PrimFunc(symgensym), tmp8481)


__e.TailApply(tmp8462, tmp8482)
return


}, 1)

tmp8483 := Call(__e, PrimFunc(symprotect), symK)


tmp8484 := Call(__e, PrimFunc(symgensym), tmp8483)


__e.TailApply(tmp8461, tmp8484)
return


}, 1)

tmp8485 := Call(__e, PrimFunc(symprotect), symL)


tmp8486 := Call(__e, PrimFunc(symgensym), tmp8485)


__e.TailApply(tmp8460, tmp8486)
return


}, 1)

tmp8487 := Call(__e, PrimFunc(symprotect), symV)


tmp8488 := Call(__e, PrimFunc(symgensym), tmp8487)


__e.TailApply(tmp8459, tmp8488)
return


}, 1)

tmp8489 := Call(__e, PrimFunc(symshen_4received), V1824)


__e.TailApply(tmp8458, tmp8489)
return


}, 1)

tmp8490 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4_5body_6), X)
return
}, 1)

tmp8491 := Call(__e, PrimFunc(symcompile), tmp8490, V1824)


__e.TailApply(tmp8457, tmp8491)
return


}, 1)

tmp8492 := PrimCons(True, Nil)

tmp8493 := PrimCons(symfreeze, tmp8492)

__e.TailApply(tmp8456, tmp8493)
return


}, 1)

__e.TailApply(tmp8455, MakeNumber(0))
return


}, 1)

tmp8494 := PrimCons(MakeNumber(0), Nil)

tmp8495 := PrimCons(symvector, tmp8494)

tmp8496 := PrimCons(tmp8495, Nil)

tmp8497 := PrimCons(MakeNumber(0), tmp8496)

tmp8498 := PrimCons(True, tmp8497)

tmp8499 := PrimCons(sym_8v, tmp8498)

__e.TailApply(tmp8454, tmp8499)
return


}, 1)

tmp8500 := PrimCons(symshen_4reset_1prolog_1vector, Nil)

__e.TailApply(tmp8453, tmp8500)
return


}, 1)

tmp8501 := Call(__e, ns2_1set, symshen_4call_1prolog, tmp8452)


_ = tmp8501

tmp8502 := MakeNative(func(__e *ControlFlow) {
V1827 := __e.Get(1)
_ = V1827
tmp8523 := PrimIsPair(V1827)

var ifres8510 Obj

if True == tmp8523 {
tmp8521 := PrimHead(V1827)

tmp8522 := PrimEqual(symreceive, tmp8521)

var ifres8512 Obj

if True == tmp8522 {
tmp8519 := PrimTail(V1827)

tmp8520 := PrimIsPair(tmp8519)

var ifres8514 Obj

if True == tmp8520 {
tmp8516 := PrimTail(V1827)

tmp8517 := PrimTail(tmp8516)

tmp8518 := PrimEqual(Nil, tmp8517)

var ifres8515 Obj

if True == tmp8518 {
ifres8515 = True


} else {
ifres8515 = False


}

ifres8514 = ifres8515


} else {
ifres8514 = False


}

var ifres8513 Obj

if True == ifres8514 {
ifres8513 = True


} else {
ifres8513 = False


}

ifres8512 = ifres8513


} else {
ifres8512 = False


}

var ifres8511 Obj

if True == ifres8512 {
ifres8511 = True


} else {
ifres8511 = False


}

ifres8510 = ifres8511


} else {
ifres8510 = False


}

if True == ifres8510 {
__e.Return(PrimTail(V1827))
return
} else {
tmp8508 := PrimIsPair(V1827)

if True == tmp8508 {
tmp8503 := PrimHead(V1827)

tmp8504 := Call(__e, PrimFunc(symshen_4received), tmp8503)


tmp8505 := PrimTail(V1827)

tmp8506 := Call(__e, PrimFunc(symshen_4received), tmp8505)


__e.TailApply(PrimFunc(symunion), tmp8504, tmp8506)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp8524 := Call(__e, ns2_1set, symshen_4received, tmp8502)


_ = tmp8524

tmp8525 := MakeNative(func(__e *ControlFlow) {
tmp8526 := PrimValue(symshen_4_dprolog_1vector_d)

__e.Return(PrimVectorSet(tmp8526, MakeNumber(1), MakeNumber(2)))
return


}, 0)

tmp8527 := Call(__e, ns2_1set, symshen_4reset_1prolog_1vector, tmp8525)


_ = tmp8527

tmp8528 := MakeNative(func(__e *ControlFlow) {
V1828 := __e.Get(1)
_ = V1828
__e.Return(V1828)
return
}, 1)

tmp8529 := Call(__e, ns2_1set, symreceive, tmp8528)


_ = tmp8529

tmp8530 := MakeNative(func(__e *ControlFlow) {
V1829 := __e.Get(1)
_ = V1829
tmp8544 := PrimIsPair(V1829)

var ifres8536 Obj

if True == tmp8544 {
tmp8542 := PrimHead(V1829)

tmp8543 := PrimEqual(symdefprolog, tmp8542)

var ifres8538 Obj

if True == tmp8543 {
tmp8540 := PrimTail(V1829)

tmp8541 := PrimIsPair(tmp8540)

var ifres8539 Obj

if True == tmp8541 {
ifres8539 = True


} else {
ifres8539 = False


}

ifres8538 = ifres8539


} else {
ifres8538 = False


}

var ifres8537 Obj

if True == ifres8538 {
ifres8537 = True


} else {
ifres8537 = False


}

ifres8536 = ifres8537


} else {
ifres8536 = False


}

if True == ifres8536 {
tmp8531 := PrimTail(V1829)

tmp8532 := PrimHead(tmp8531)

tmp8533 := PrimTail(V1829)

tmp8534 := PrimTail(tmp8533)

__e.TailApply(PrimFunc(symshen_4compile_1prolog), tmp8532, tmp8534)
return


} else {
__e.Return(V1829)
return
}


}, 1)

tmp8545 := Call(__e, ns2_1set, symshen_4defprolog_1macro, tmp8530)


_ = tmp8545

tmp8546 := MakeNative(func(__e *ControlFlow) {
V1830 := __e.Get(1)
_ = V1830
tmp8566 := PrimIsPair(V1830)

var ifres8558 Obj

if True == tmp8566 {
tmp8564 := PrimHead(V1830)

tmp8565 := PrimEqual(symdatatype, tmp8564)

var ifres8560 Obj

if True == tmp8565 {
tmp8562 := PrimTail(V1830)

tmp8563 := PrimIsPair(tmp8562)

var ifres8561 Obj

if True == tmp8563 {
ifres8561 = True


} else {
ifres8561 = False


}

ifres8560 = ifres8561


} else {
ifres8560 = False


}

var ifres8559 Obj

if True == ifres8560 {
ifres8559 = True


} else {
ifres8559 = False


}

ifres8558 = ifres8559


} else {
ifres8558 = False


}

if True == ifres8558 {
tmp8547 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp8548 := MakeNative(func(__e *ControlFlow) {
Compile := __e.Get(1)
_ = Compile
__e.Return(D)
return
}, 1)

tmp8549 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4_5datatype_6), X)
return
}, 1)

tmp8550 := PrimTail(V1830)

tmp8551 := PrimTail(tmp8550)

tmp8552 := PrimCons(D, tmp8551)

tmp8553 := Call(__e, PrimFunc(symcompile), tmp8549, tmp8552)


__e.TailApply(tmp8548, tmp8553)
return


}, 1)

tmp8554 := PrimTail(V1830)

tmp8555 := PrimHead(tmp8554)

tmp8556 := Call(__e, PrimFunc(symshen_4intern_1type), tmp8555)


__e.TailApply(tmp8547, tmp8556)
return


} else {
__e.Return(V1830)
return
}


}, 1)

tmp8567 := Call(__e, ns2_1set, symshen_4datatype_1macro, tmp8546)


_ = tmp8567

tmp8568 := MakeNative(func(__e *ControlFlow) {
V1831 := __e.Get(1)
_ = V1831
tmp8576 := PrimIsPair(V1831)

if True == tmp8576 {
tmp8569 := PrimHead(V1831)

tmp8570 := Call(__e, PrimFunc(symshen_4rcons__form), tmp8569)


tmp8571 := PrimTail(V1831)

tmp8572 := Call(__e, PrimFunc(symshen_4rcons__form), tmp8571)


tmp8573 := PrimCons(tmp8572, Nil)

tmp8574 := PrimCons(tmp8570, tmp8573)

__e.Return(PrimCons(symcons, tmp8574))
return


} else {
__e.Return(V1831)
return
}


}, 1)

tmp8577 := Call(__e, ns2_1set, symshen_4rcons__form, tmp8568)


_ = tmp8577

tmp8578 := MakeNative(func(__e *ControlFlow) {
V1832 := __e.Get(1)
_ = V1832
tmp8579 := PrimStr(V1832)

tmp8580 := PrimStringConcat(tmp8579, MakeString("#type"))

__e.Return(PrimIntern(tmp8580))
return


}, 1)

tmp8581 := Call(__e, ns2_1set, symshen_4intern_1type, tmp8578)


_ = tmp8581

tmp8582 := MakeNative(func(__e *ControlFlow) {
V1833 := __e.Get(1)
_ = V1833
tmp8648 := PrimIsPair(V1833)

var ifres8629 Obj

if True == tmp8648 {
tmp8646 := PrimHead(V1833)

tmp8647 := PrimEqual(sym_8s, tmp8646)

var ifres8631 Obj

if True == tmp8647 {
tmp8644 := PrimTail(V1833)

tmp8645 := PrimIsPair(tmp8644)

var ifres8633 Obj

if True == tmp8645 {
tmp8641 := PrimTail(V1833)

tmp8642 := PrimTail(tmp8641)

tmp8643 := PrimIsPair(tmp8642)

var ifres8635 Obj

if True == tmp8643 {
tmp8637 := PrimTail(V1833)

tmp8638 := PrimTail(tmp8637)

tmp8639 := PrimTail(tmp8638)

tmp8640 := PrimIsPair(tmp8639)

var ifres8636 Obj

if True == tmp8640 {
ifres8636 = True


} else {
ifres8636 = False


}

ifres8635 = ifres8636


} else {
ifres8635 = False


}

var ifres8634 Obj

if True == ifres8635 {
ifres8634 = True


} else {
ifres8634 = False


}

ifres8633 = ifres8634


} else {
ifres8633 = False


}

var ifres8632 Obj

if True == ifres8633 {
ifres8632 = True


} else {
ifres8632 = False


}

ifres8631 = ifres8632


} else {
ifres8631 = False


}

var ifres8630 Obj

if True == ifres8631 {
ifres8630 = True


} else {
ifres8630 = False


}

ifres8629 = ifres8630


} else {
ifres8629 = False


}

if True == ifres8629 {
tmp8583 := PrimTail(V1833)

tmp8584 := PrimHead(tmp8583)

tmp8585 := PrimTail(V1833)

tmp8586 := PrimTail(tmp8585)

tmp8587 := PrimCons(sym_8s, tmp8586)

tmp8588 := Call(__e, PrimFunc(symshen_4_8s_1macro), tmp8587)


tmp8589 := PrimCons(tmp8588, Nil)

tmp8590 := PrimCons(tmp8584, tmp8589)

__e.Return(PrimCons(sym_8s, tmp8590))
return


} else {
tmp8627 := PrimIsPair(V1833)

var ifres8603 Obj

if True == tmp8627 {
tmp8625 := PrimHead(V1833)

tmp8626 := PrimEqual(sym_8s, tmp8625)

var ifres8605 Obj

if True == tmp8626 {
tmp8623 := PrimTail(V1833)

tmp8624 := PrimIsPair(tmp8623)

var ifres8607 Obj

if True == tmp8624 {
tmp8620 := PrimTail(V1833)

tmp8621 := PrimTail(tmp8620)

tmp8622 := PrimIsPair(tmp8621)

var ifres8609 Obj

if True == tmp8622 {
tmp8616 := PrimTail(V1833)

tmp8617 := PrimTail(tmp8616)

tmp8618 := PrimTail(tmp8617)

tmp8619 := PrimEqual(Nil, tmp8618)

var ifres8611 Obj

if True == tmp8619 {
tmp8613 := PrimTail(V1833)

tmp8614 := PrimHead(tmp8613)

tmp8615 := PrimIsString(tmp8614)

var ifres8612 Obj

if True == tmp8615 {
ifres8612 = True


} else {
ifres8612 = False


}

ifres8611 = ifres8612


} else {
ifres8611 = False


}

var ifres8610 Obj

if True == ifres8611 {
ifres8610 = True


} else {
ifres8610 = False


}

ifres8609 = ifres8610


} else {
ifres8609 = False


}

var ifres8608 Obj

if True == ifres8609 {
ifres8608 = True


} else {
ifres8608 = False


}

ifres8607 = ifres8608


} else {
ifres8607 = False


}

var ifres8606 Obj

if True == ifres8607 {
ifres8606 = True


} else {
ifres8606 = False


}

ifres8605 = ifres8606


} else {
ifres8605 = False


}

var ifres8604 Obj

if True == ifres8605 {
ifres8604 = True


} else {
ifres8604 = False


}

ifres8603 = ifres8604


} else {
ifres8603 = False


}

if True == ifres8603 {
tmp8591 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
tmp8597 := Call(__e, PrimFunc(symlength), E)


tmp8598 := PrimGreatThan(tmp8597, MakeNumber(1))

if True == tmp8598 {
tmp8592 := PrimTail(V1833)

tmp8593 := PrimTail(tmp8592)

tmp8594 := Call(__e, PrimFunc(symappend), E, tmp8593)


tmp8595 := PrimCons(sym_8s, tmp8594)

__e.TailApply(PrimFunc(symshen_4_8s_1macro), tmp8595)
return


} else {
__e.Return(V1833)
return
}


}, 1)

tmp8599 := PrimTail(V1833)

tmp8600 := PrimHead(tmp8599)

tmp8601 := Call(__e, PrimFunc(symexplode), tmp8600)


__e.TailApply(tmp8591, tmp8601)
return


} else {
__e.Return(V1833)
return
}


}


}, 1)

tmp8649 := Call(__e, ns2_1set, symshen_4_8s_1macro, tmp8582)


_ = tmp8649

tmp8650 := MakeNative(func(__e *ControlFlow) {
V1834 := __e.Get(1)
_ = V1834
tmp8659 := PrimIsPair(V1834)

var ifres8655 Obj

if True == tmp8659 {
tmp8657 := PrimHead(V1834)

tmp8658 := PrimEqual(symsynonyms, tmp8657)

var ifres8656 Obj

if True == tmp8658 {
ifres8656 = True


} else {
ifres8656 = False


}

ifres8655 = ifres8656


} else {
ifres8655 = False


}

if True == ifres8655 {
tmp8651 := PrimTail(V1834)

tmp8652 := PrimValue(symshen_4_dsynonyms_d)

tmp8653 := Call(__e, PrimFunc(symappend), tmp8651, tmp8652)


__e.TailApply(PrimFunc(symshen_4synonyms_1h), tmp8653)
return


} else {
__e.Return(V1834)
return
}


}, 1)

tmp8660 := Call(__e, ns2_1set, symshen_4synonyms_1macro, tmp8650)


_ = tmp8660

tmp8661 := MakeNative(func(__e *ControlFlow) {
V1835 := __e.Get(1)
_ = V1835
tmp8662 := MakeNative(func(__e *ControlFlow) {
CurryTypes := __e.Get(1)
_ = CurryTypes
tmp8663 := MakeNative(func(__e *ControlFlow) {
Eval := __e.Get(1)
_ = Eval
__e.Return(symsynonyms)
return
}, 1)

tmp8664 := Call(__e, PrimFunc(symshen_4compile_1synonyms), CurryTypes)


tmp8665 := PrimCons(symshen_4demod, tmp8664)

tmp8666 := PrimCons(symdefine, tmp8665)

tmp8667 := Call(__e, PrimFunc(symeval), tmp8666)


__e.TailApply(tmp8663, tmp8667)
return


}, 1)

tmp8668 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4curry_1type), X)
return
}, 1)

tmp8669 := Call(__e, PrimFunc(symmap), tmp8668, V1835)


__e.TailApply(tmp8662, tmp8669)
return


}, 1)

tmp8670 := Call(__e, ns2_1set, symshen_4synonyms_1h, tmp8661)


_ = tmp8670

tmp8671 := MakeNative(func(__e *ControlFlow) {
V1838 := __e.Get(1)
_ = V1838
tmp8694 := PrimEqual(Nil, V1838)

if True == tmp8694 {
tmp8672 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp8673 := PrimCons(X, Nil)

tmp8674 := PrimCons(sym_1_6, tmp8673)

__e.Return(PrimCons(X, tmp8674))
return


}, 1)

tmp8675 := Call(__e, PrimFunc(symprotect), symX)


tmp8676 := Call(__e, PrimFunc(symgensym), tmp8675)


__e.TailApply(tmp8672, tmp8676)
return


} else {
tmp8692 := PrimIsPair(V1838)

var ifres8688 Obj

if True == tmp8692 {
tmp8690 := PrimTail(V1838)

tmp8691 := PrimIsPair(tmp8690)

var ifres8689 Obj

if True == tmp8691 {
ifres8689 = True


} else {
ifres8689 = False


}

ifres8688 = ifres8689


} else {
ifres8688 = False


}

if True == ifres8688 {
tmp8677 := PrimHead(V1838)

tmp8678 := Call(__e, PrimFunc(symshen_4rcons__form), tmp8677)


tmp8679 := PrimTail(V1838)

tmp8680 := PrimHead(tmp8679)

tmp8681 := Call(__e, PrimFunc(symshen_4rcons__form), tmp8680)


tmp8682 := PrimTail(V1838)

tmp8683 := PrimTail(tmp8682)

tmp8684 := Call(__e, PrimFunc(symshen_4compile_1synonyms), tmp8683)


tmp8685 := PrimCons(tmp8681, tmp8684)

tmp8686 := PrimCons(sym_1_6, tmp8685)

__e.Return(PrimCons(tmp8678, tmp8686))
return


} else {
__e.Return(PrimSimpleError(MakeString("synonyms requires an even number of arguments\n")))
return
}


}


}, 1)

tmp8695 := Call(__e, ns2_1set, symshen_4compile_1synonyms, tmp8671)


_ = tmp8695

tmp8696 := MakeNative(func(__e *ControlFlow) {
V1839 := __e.Get(1)
_ = V1839
tmp8707 := PrimIsPair(V1839)

var ifres8699 Obj

if True == tmp8707 {
tmp8705 := PrimHead(V1839)

tmp8706 := PrimEqual(symnl, tmp8705)

var ifres8701 Obj

if True == tmp8706 {
tmp8703 := PrimTail(V1839)

tmp8704 := PrimEqual(Nil, tmp8703)

var ifres8702 Obj

if True == tmp8704 {
ifres8702 = True


} else {
ifres8702 = False


}

ifres8701 = ifres8702


} else {
ifres8701 = False


}

var ifres8700 Obj

if True == ifres8701 {
ifres8700 = True


} else {
ifres8700 = False


}

ifres8699 = ifres8700


} else {
ifres8699 = False


}

if True == ifres8699 {
tmp8697 := PrimCons(MakeNumber(1), Nil)

__e.Return(PrimCons(symnl, tmp8697))
return


} else {
__e.Return(V1839)
return
}


}, 1)

tmp8708 := Call(__e, ns2_1set, symshen_4nl_1macro, tmp8696)


_ = tmp8708

tmp8709 := MakeNative(func(__e *ControlFlow) {
V1840 := __e.Get(1)
_ = V1840
tmp8748 := PrimIsPair(V1840)

var ifres8721 Obj

if True == tmp8748 {
tmp8746 := PrimTail(V1840)

tmp8747 := PrimIsPair(tmp8746)

var ifres8723 Obj

if True == tmp8747 {
tmp8743 := PrimTail(V1840)

tmp8744 := PrimTail(tmp8743)

tmp8745 := PrimIsPair(tmp8744)

var ifres8725 Obj

if True == tmp8745 {
tmp8739 := PrimTail(V1840)

tmp8740 := PrimTail(tmp8739)

tmp8741 := PrimTail(tmp8740)

tmp8742 := PrimIsPair(tmp8741)

var ifres8727 Obj

if True == tmp8742 {
tmp8729 := PrimHead(V1840)

tmp8730 := PrimCons(symdo, Nil)

tmp8731 := PrimCons(sym_d, tmp8730)

tmp8732 := PrimCons(sym_7, tmp8731)

tmp8733 := PrimCons(symor, tmp8732)

tmp8734 := PrimCons(symand, tmp8733)

tmp8735 := PrimCons(symappend, tmp8734)

tmp8736 := PrimCons(sym_8v, tmp8735)

tmp8737 := PrimCons(sym_8p, tmp8736)

tmp8738 := Call(__e, PrimFunc(symelement_2), tmp8729, tmp8737)


var ifres8728 Obj

if True == tmp8738 {
ifres8728 = True


} else {
ifres8728 = False


}

ifres8727 = ifres8728


} else {
ifres8727 = False


}

var ifres8726 Obj

if True == ifres8727 {
ifres8726 = True


} else {
ifres8726 = False


}

ifres8725 = ifres8726


} else {
ifres8725 = False


}

var ifres8724 Obj

if True == ifres8725 {
ifres8724 = True


} else {
ifres8724 = False


}

ifres8723 = ifres8724


} else {
ifres8723 = False


}

var ifres8722 Obj

if True == ifres8723 {
ifres8722 = True


} else {
ifres8722 = False


}

ifres8721 = ifres8722


} else {
ifres8721 = False


}

if True == ifres8721 {
tmp8710 := PrimHead(V1840)

tmp8711 := PrimTail(V1840)

tmp8712 := PrimHead(tmp8711)

tmp8713 := PrimHead(V1840)

tmp8714 := PrimTail(V1840)

tmp8715 := PrimTail(tmp8714)

tmp8716 := PrimCons(tmp8713, tmp8715)

tmp8717 := Call(__e, PrimFunc(symshen_4assoc_1macro), tmp8716)


tmp8718 := PrimCons(tmp8717, Nil)

tmp8719 := PrimCons(tmp8712, tmp8718)

__e.Return(PrimCons(tmp8710, tmp8719))
return


} else {
__e.Return(V1840)
return
}


}, 1)

tmp8749 := Call(__e, ns2_1set, symshen_4assoc_1macro, tmp8709)


_ = tmp8749

tmp8750 := MakeNative(func(__e *ControlFlow) {
V1841 := __e.Get(1)
_ = V1841
tmp8826 := PrimIsPair(V1841)

var ifres8800 Obj

if True == tmp8826 {
tmp8824 := PrimHead(V1841)

tmp8825 := PrimEqual(symlet, tmp8824)

var ifres8802 Obj

if True == tmp8825 {
tmp8822 := PrimTail(V1841)

tmp8823 := PrimIsPair(tmp8822)

var ifres8804 Obj

if True == tmp8823 {
tmp8819 := PrimTail(V1841)

tmp8820 := PrimTail(tmp8819)

tmp8821 := PrimIsPair(tmp8820)

var ifres8806 Obj

if True == tmp8821 {
tmp8815 := PrimTail(V1841)

tmp8816 := PrimTail(tmp8815)

tmp8817 := PrimTail(tmp8816)

tmp8818 := PrimIsPair(tmp8817)

var ifres8808 Obj

if True == tmp8818 {
tmp8810 := PrimTail(V1841)

tmp8811 := PrimTail(tmp8810)

tmp8812 := PrimTail(tmp8811)

tmp8813 := PrimTail(tmp8812)

tmp8814 := PrimIsPair(tmp8813)

var ifres8809 Obj

if True == tmp8814 {
ifres8809 = True


} else {
ifres8809 = False


}

ifres8808 = ifres8809


} else {
ifres8808 = False


}

var ifres8807 Obj

if True == ifres8808 {
ifres8807 = True


} else {
ifres8807 = False


}

ifres8806 = ifres8807


} else {
ifres8806 = False


}

var ifres8805 Obj

if True == ifres8806 {
ifres8805 = True


} else {
ifres8805 = False


}

ifres8804 = ifres8805


} else {
ifres8804 = False


}

var ifres8803 Obj

if True == ifres8804 {
ifres8803 = True


} else {
ifres8803 = False


}

ifres8802 = ifres8803


} else {
ifres8802 = False


}

var ifres8801 Obj

if True == ifres8802 {
ifres8801 = True


} else {
ifres8801 = False


}

ifres8800 = ifres8801


} else {
ifres8800 = False


}

if True == ifres8800 {
tmp8751 := PrimTail(V1841)

tmp8752 := PrimHead(tmp8751)

tmp8753 := PrimTail(V1841)

tmp8754 := PrimTail(tmp8753)

tmp8755 := PrimHead(tmp8754)

tmp8756 := PrimTail(V1841)

tmp8757 := PrimTail(tmp8756)

tmp8758 := PrimTail(tmp8757)

tmp8759 := PrimCons(symlet, tmp8758)

tmp8760 := Call(__e, PrimFunc(symshen_4let_1macro), tmp8759)


tmp8761 := PrimCons(tmp8760, Nil)

tmp8762 := PrimCons(tmp8755, tmp8761)

tmp8763 := PrimCons(tmp8752, tmp8762)

__e.Return(PrimCons(symlet, tmp8763))
return


} else {
tmp8798 := PrimIsPair(V1841)

var ifres8772 Obj

if True == tmp8798 {
tmp8796 := PrimHead(V1841)

tmp8797 := PrimEqual(symlet, tmp8796)

var ifres8774 Obj

if True == tmp8797 {
tmp8794 := PrimTail(V1841)

tmp8795 := PrimIsPair(tmp8794)

var ifres8776 Obj

if True == tmp8795 {
tmp8791 := PrimTail(V1841)

tmp8792 := PrimTail(tmp8791)

tmp8793 := PrimIsPair(tmp8792)

var ifres8778 Obj

if True == tmp8793 {
tmp8787 := PrimTail(V1841)

tmp8788 := PrimTail(tmp8787)

tmp8789 := PrimTail(tmp8788)

tmp8790 := PrimIsPair(tmp8789)

var ifres8780 Obj

if True == tmp8790 {
tmp8782 := PrimTail(V1841)

tmp8783 := PrimTail(tmp8782)

tmp8784 := PrimTail(tmp8783)

tmp8785 := PrimTail(tmp8784)

tmp8786 := PrimEqual(Nil, tmp8785)

var ifres8781 Obj

if True == tmp8786 {
ifres8781 = True


} else {
ifres8781 = False


}

ifres8780 = ifres8781


} else {
ifres8780 = False


}

var ifres8779 Obj

if True == ifres8780 {
ifres8779 = True


} else {
ifres8779 = False


}

ifres8778 = ifres8779


} else {
ifres8778 = False


}

var ifres8777 Obj

if True == ifres8778 {
ifres8777 = True


} else {
ifres8777 = False


}

ifres8776 = ifres8777


} else {
ifres8776 = False


}

var ifres8775 Obj

if True == ifres8776 {
ifres8775 = True


} else {
ifres8775 = False


}

ifres8774 = ifres8775


} else {
ifres8774 = False


}

var ifres8773 Obj

if True == ifres8774 {
ifres8773 = True


} else {
ifres8773 = False


}

ifres8772 = ifres8773


} else {
ifres8772 = False


}

if True == ifres8772 {
tmp8768 := PrimTail(V1841)

tmp8769 := PrimHead(tmp8768)

tmp8770 := PrimIsVariable(tmp8769)

if True == tmp8770 {
__e.Return(V1841)
return
} else {
tmp8764 := PrimTail(V1841)

tmp8765 := PrimHead(tmp8764)

tmp8766 := Call(__e, PrimFunc(symshen_4app), tmp8765, MakeString(" is not a variable\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp8766))
return


}


} else {
__e.Return(V1841)
return
}


}


}, 1)

tmp8827 := Call(__e, ns2_1set, symshen_4let_1macro, tmp8750)


_ = tmp8827

tmp8828 := MakeNative(func(__e *ControlFlow) {
V1842 := __e.Get(1)
_ = V1842
tmp8886 := PrimIsPair(V1842)

var ifres8867 Obj

if True == tmp8886 {
tmp8884 := PrimHead(V1842)

tmp8885 := PrimEqual(sym_c_4, tmp8884)

var ifres8869 Obj

if True == tmp8885 {
tmp8882 := PrimTail(V1842)

tmp8883 := PrimIsPair(tmp8882)

var ifres8871 Obj

if True == tmp8883 {
tmp8879 := PrimTail(V1842)

tmp8880 := PrimTail(tmp8879)

tmp8881 := PrimIsPair(tmp8880)

var ifres8873 Obj

if True == tmp8881 {
tmp8875 := PrimTail(V1842)

tmp8876 := PrimTail(tmp8875)

tmp8877 := PrimTail(tmp8876)

tmp8878 := PrimIsPair(tmp8877)

var ifres8874 Obj

if True == tmp8878 {
ifres8874 = True


} else {
ifres8874 = False


}

ifres8873 = ifres8874


} else {
ifres8873 = False


}

var ifres8872 Obj

if True == ifres8873 {
ifres8872 = True


} else {
ifres8872 = False


}

ifres8871 = ifres8872


} else {
ifres8871 = False


}

var ifres8870 Obj

if True == ifres8871 {
ifres8870 = True


} else {
ifres8870 = False


}

ifres8869 = ifres8870


} else {
ifres8869 = False


}

var ifres8868 Obj

if True == ifres8869 {
ifres8868 = True


} else {
ifres8868 = False


}

ifres8867 = ifres8868


} else {
ifres8867 = False


}

if True == ifres8867 {
tmp8829 := PrimTail(V1842)

tmp8830 := PrimHead(tmp8829)

tmp8831 := PrimTail(V1842)

tmp8832 := PrimTail(tmp8831)

tmp8833 := PrimCons(sym_c_4, tmp8832)

tmp8834 := Call(__e, PrimFunc(symshen_4abs_1macro), tmp8833)


tmp8835 := PrimCons(tmp8834, Nil)

tmp8836 := PrimCons(tmp8830, tmp8835)

__e.Return(PrimCons(symlambda, tmp8836))
return


} else {
tmp8865 := PrimIsPair(V1842)

var ifres8846 Obj

if True == tmp8865 {
tmp8863 := PrimHead(V1842)

tmp8864 := PrimEqual(sym_c_4, tmp8863)

var ifres8848 Obj

if True == tmp8864 {
tmp8861 := PrimTail(V1842)

tmp8862 := PrimIsPair(tmp8861)

var ifres8850 Obj

if True == tmp8862 {
tmp8858 := PrimTail(V1842)

tmp8859 := PrimTail(tmp8858)

tmp8860 := PrimIsPair(tmp8859)

var ifres8852 Obj

if True == tmp8860 {
tmp8854 := PrimTail(V1842)

tmp8855 := PrimTail(tmp8854)

tmp8856 := PrimTail(tmp8855)

tmp8857 := PrimEqual(Nil, tmp8856)

var ifres8853 Obj

if True == tmp8857 {
ifres8853 = True


} else {
ifres8853 = False


}

ifres8852 = ifres8853


} else {
ifres8852 = False


}

var ifres8851 Obj

if True == ifres8852 {
ifres8851 = True


} else {
ifres8851 = False


}

ifres8850 = ifres8851


} else {
ifres8850 = False


}

var ifres8849 Obj

if True == ifres8850 {
ifres8849 = True


} else {
ifres8849 = False


}

ifres8848 = ifres8849


} else {
ifres8848 = False


}

var ifres8847 Obj

if True == ifres8848 {
ifres8847 = True


} else {
ifres8847 = False


}

ifres8846 = ifres8847


} else {
ifres8846 = False


}

if True == ifres8846 {
tmp8842 := PrimTail(V1842)

tmp8843 := PrimHead(tmp8842)

tmp8844 := PrimIsVariable(tmp8843)

if True == tmp8844 {
tmp8837 := PrimTail(V1842)

__e.Return(PrimCons(symlambda, tmp8837))
return


} else {
tmp8838 := PrimTail(V1842)

tmp8839 := PrimHead(tmp8838)

tmp8840 := Call(__e, PrimFunc(symshen_4app), tmp8839, MakeString(" is not a variable\n"), symshen_4s)


__e.Return(PrimSimpleError(tmp8840))
return


}


} else {
__e.Return(V1842)
return
}


}


}, 1)

tmp8887 := Call(__e, ns2_1set, symshen_4abs_1macro, tmp8828)


_ = tmp8887

tmp8888 := MakeNative(func(__e *ControlFlow) {
V1845 := __e.Get(1)
_ = V1845
tmp8984 := PrimIsPair(V1845)

var ifres8966 Obj

if True == tmp8984 {
tmp8982 := PrimHead(V1845)

tmp8983 := PrimEqual(symcases, tmp8982)

var ifres8968 Obj

if True == tmp8983 {
tmp8980 := PrimTail(V1845)

tmp8981 := PrimIsPair(tmp8980)

var ifres8970 Obj

if True == tmp8981 {
tmp8977 := PrimTail(V1845)

tmp8978 := PrimHead(tmp8977)

tmp8979 := PrimEqual(True, tmp8978)

var ifres8972 Obj

if True == tmp8979 {
tmp8974 := PrimTail(V1845)

tmp8975 := PrimTail(tmp8974)

tmp8976 := PrimIsPair(tmp8975)

var ifres8973 Obj

if True == tmp8976 {
ifres8973 = True


} else {
ifres8973 = False


}

ifres8972 = ifres8973


} else {
ifres8972 = False


}

var ifres8971 Obj

if True == ifres8972 {
ifres8971 = True


} else {
ifres8971 = False


}

ifres8970 = ifres8971


} else {
ifres8970 = False


}

var ifres8969 Obj

if True == ifres8970 {
ifres8969 = True


} else {
ifres8969 = False


}

ifres8968 = ifres8969


} else {
ifres8968 = False


}

var ifres8967 Obj

if True == ifres8968 {
ifres8967 = True


} else {
ifres8967 = False


}

ifres8966 = ifres8967


} else {
ifres8966 = False


}

if True == ifres8966 {
tmp8889 := PrimTail(V1845)

tmp8890 := PrimTail(tmp8889)

__e.Return(PrimHead(tmp8890))
return


} else {
tmp8964 := PrimIsPair(V1845)

var ifres8945 Obj

if True == tmp8964 {
tmp8962 := PrimHead(V1845)

tmp8963 := PrimEqual(symcases, tmp8962)

var ifres8947 Obj

if True == tmp8963 {
tmp8960 := PrimTail(V1845)

tmp8961 := PrimIsPair(tmp8960)

var ifres8949 Obj

if True == tmp8961 {
tmp8957 := PrimTail(V1845)

tmp8958 := PrimTail(tmp8957)

tmp8959 := PrimIsPair(tmp8958)

var ifres8951 Obj

if True == tmp8959 {
tmp8953 := PrimTail(V1845)

tmp8954 := PrimTail(tmp8953)

tmp8955 := PrimTail(tmp8954)

tmp8956 := PrimEqual(Nil, tmp8955)

var ifres8952 Obj

if True == tmp8956 {
ifres8952 = True


} else {
ifres8952 = False


}

ifres8951 = ifres8952


} else {
ifres8951 = False


}

var ifres8950 Obj

if True == ifres8951 {
ifres8950 = True


} else {
ifres8950 = False


}

ifres8949 = ifres8950


} else {
ifres8949 = False


}

var ifres8948 Obj

if True == ifres8949 {
ifres8948 = True


} else {
ifres8948 = False


}

ifres8947 = ifres8948


} else {
ifres8947 = False


}

var ifres8946 Obj

if True == ifres8947 {
ifres8946 = True


} else {
ifres8946 = False


}

ifres8945 = ifres8946


} else {
ifres8945 = False


}

if True == ifres8945 {
tmp8891 := PrimTail(V1845)

tmp8892 := PrimHead(tmp8891)

tmp8893 := PrimTail(V1845)

tmp8894 := PrimTail(tmp8893)

tmp8895 := PrimHead(tmp8894)

tmp8896 := PrimCons(MakeString("error: cases exhausted"), Nil)

tmp8897 := PrimCons(symsimple_1error, tmp8896)

tmp8898 := PrimCons(tmp8897, Nil)

tmp8899 := PrimCons(tmp8895, tmp8898)

tmp8900 := PrimCons(tmp8892, tmp8899)

__e.Return(PrimCons(symif, tmp8900))
return


} else {
tmp8943 := PrimIsPair(V1845)

var ifres8930 Obj

if True == tmp8943 {
tmp8941 := PrimHead(V1845)

tmp8942 := PrimEqual(symcases, tmp8941)

var ifres8932 Obj

if True == tmp8942 {
tmp8939 := PrimTail(V1845)

tmp8940 := PrimIsPair(tmp8939)

var ifres8934 Obj

if True == tmp8940 {
tmp8936 := PrimTail(V1845)

tmp8937 := PrimTail(tmp8936)

tmp8938 := PrimIsPair(tmp8937)

var ifres8935 Obj

if True == tmp8938 {
ifres8935 = True


} else {
ifres8935 = False


}

ifres8934 = ifres8935


} else {
ifres8934 = False


}

var ifres8933 Obj

if True == ifres8934 {
ifres8933 = True


} else {
ifres8933 = False


}

ifres8932 = ifres8933


} else {
ifres8932 = False


}

var ifres8931 Obj

if True == ifres8932 {
ifres8931 = True


} else {
ifres8931 = False


}

ifres8930 = ifres8931


} else {
ifres8930 = False


}

if True == ifres8930 {
tmp8901 := PrimTail(V1845)

tmp8902 := PrimHead(tmp8901)

tmp8903 := PrimTail(V1845)

tmp8904 := PrimTail(tmp8903)

tmp8905 := PrimHead(tmp8904)

tmp8906 := PrimTail(V1845)

tmp8907 := PrimTail(tmp8906)

tmp8908 := PrimTail(tmp8907)

tmp8909 := PrimCons(symcases, tmp8908)

tmp8910 := Call(__e, PrimFunc(symshen_4cases_1macro), tmp8909)


tmp8911 := PrimCons(tmp8910, Nil)

tmp8912 := PrimCons(tmp8905, tmp8911)

tmp8913 := PrimCons(tmp8902, tmp8912)

__e.Return(PrimCons(symif, tmp8913))
return


} else {
tmp8928 := PrimIsPair(V1845)

var ifres8915 Obj

if True == tmp8928 {
tmp8926 := PrimHead(V1845)

tmp8927 := PrimEqual(symcases, tmp8926)

var ifres8917 Obj

if True == tmp8927 {
tmp8924 := PrimTail(V1845)

tmp8925 := PrimIsPair(tmp8924)

var ifres8919 Obj

if True == tmp8925 {
tmp8921 := PrimTail(V1845)

tmp8922 := PrimTail(tmp8921)

tmp8923 := PrimEqual(Nil, tmp8922)

var ifres8920 Obj

if True == tmp8923 {
ifres8920 = True


} else {
ifres8920 = False


}

ifres8919 = ifres8920


} else {
ifres8919 = False


}

var ifres8918 Obj

if True == ifres8919 {
ifres8918 = True


} else {
ifres8918 = False


}

ifres8917 = ifres8918


} else {
ifres8917 = False


}

var ifres8916 Obj

if True == ifres8917 {
ifres8916 = True


} else {
ifres8916 = False


}

ifres8915 = ifres8916


} else {
ifres8915 = False


}

if True == ifres8915 {
__e.Return(PrimSimpleError(MakeString("error: odd number of case elements\n")))
return
} else {
__e.Return(V1845)
return
}


}


}


}


}, 1)

tmp8985 := Call(__e, ns2_1set, symshen_4cases_1macro, tmp8888)


_ = tmp8985

tmp8986 := MakeNative(func(__e *ControlFlow) {
V1846 := __e.Get(1)
_ = V1846
tmp9043 := PrimIsPair(V1846)

var ifres9030 Obj

if True == tmp9043 {
tmp9041 := PrimHead(V1846)

tmp9042 := PrimEqual(symtime, tmp9041)

var ifres9032 Obj

if True == tmp9042 {
tmp9039 := PrimTail(V1846)

tmp9040 := PrimIsPair(tmp9039)

var ifres9034 Obj

if True == tmp9040 {
tmp9036 := PrimTail(V1846)

tmp9037 := PrimTail(tmp9036)

tmp9038 := PrimEqual(Nil, tmp9037)

var ifres9035 Obj

if True == tmp9038 {
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

var ifres9031 Obj

if True == ifres9032 {
ifres9031 = True


} else {
ifres9031 = False


}

ifres9030 = ifres9031


} else {
ifres9030 = False


}

if True == ifres9030 {
tmp8987 := Call(__e, PrimFunc(symprotect), symStart)


tmp8988 := PrimCons(symrun, Nil)

tmp8989 := PrimCons(symget_1time, tmp8988)

tmp8990 := Call(__e, PrimFunc(symprotect), symResult)


tmp8991 := PrimTail(V1846)

tmp8992 := PrimHead(tmp8991)

tmp8993 := Call(__e, PrimFunc(symprotect), symFinish)


tmp8994 := PrimCons(symrun, Nil)

tmp8995 := PrimCons(symget_1time, tmp8994)

tmp8996 := Call(__e, PrimFunc(symprotect), symTime)


tmp8997 := Call(__e, PrimFunc(symprotect), symFinish)


tmp8998 := Call(__e, PrimFunc(symprotect), symStart)


tmp8999 := PrimCons(tmp8998, Nil)

tmp9000 := PrimCons(tmp8997, tmp8999)

tmp9001 := PrimCons(sym_1, tmp9000)

tmp9002 := Call(__e, PrimFunc(symprotect), symMessage)


tmp9003 := Call(__e, PrimFunc(symprotect), symTime)


tmp9004 := PrimCons(tmp9003, Nil)

tmp9005 := PrimCons(symstr, tmp9004)

tmp9006 := PrimCons(MakeString(" secs\n"), Nil)

tmp9007 := PrimCons(tmp9005, tmp9006)

tmp9008 := PrimCons(symcn, tmp9007)

tmp9009 := PrimCons(tmp9008, Nil)

tmp9010 := PrimCons(MakeString("\nrun time: "), tmp9009)

tmp9011 := PrimCons(symcn, tmp9010)

tmp9012 := PrimCons(symstoutput, Nil)

tmp9013 := PrimCons(tmp9012, Nil)

tmp9014 := PrimCons(tmp9011, tmp9013)

tmp9015 := PrimCons(sympr, tmp9014)

tmp9016 := Call(__e, PrimFunc(symprotect), symResult)


tmp9017 := PrimCons(tmp9016, Nil)

tmp9018 := PrimCons(tmp9015, tmp9017)

tmp9019 := PrimCons(tmp9002, tmp9018)

tmp9020 := PrimCons(tmp9001, tmp9019)

tmp9021 := PrimCons(tmp8996, tmp9020)

tmp9022 := PrimCons(tmp8995, tmp9021)

tmp9023 := PrimCons(tmp8993, tmp9022)

tmp9024 := PrimCons(tmp8992, tmp9023)

tmp9025 := PrimCons(tmp8990, tmp9024)

tmp9026 := PrimCons(tmp8989, tmp9025)

tmp9027 := PrimCons(tmp8987, tmp9026)

tmp9028 := PrimCons(symlet, tmp9027)

__e.TailApply(PrimFunc(symshen_4let_1macro), tmp9028)
return


} else {
__e.Return(V1846)
return
}


}, 1)

tmp9044 := Call(__e, ns2_1set, symshen_4timer_1macro, tmp8986)


_ = tmp9044

tmp9045 := MakeNative(func(__e *ControlFlow) {
V1847 := __e.Get(1)
_ = V1847
tmp9052 := PrimIsPair(V1847)

if True == tmp9052 {
tmp9046 := PrimHead(V1847)

tmp9047 := PrimTail(V1847)

tmp9048 := Call(__e, PrimFunc(symshen_4tuple_1up), tmp9047)


tmp9049 := PrimCons(tmp9048, Nil)

tmp9050 := PrimCons(tmp9046, tmp9049)

__e.Return(PrimCons(sym_8p, tmp9050))
return


} else {
__e.Return(V1847)
return
}


}, 1)

tmp9053 := Call(__e, ns2_1set, symshen_4tuple_1up, tmp9045)


_ = tmp9053

tmp9054 := MakeNative(func(__e *ControlFlow) {
V1848 := __e.Get(1)
_ = V1848
tmp9159 := PrimIsPair(V1848)

var ifres9133 Obj

if True == tmp9159 {
tmp9157 := PrimHead(V1848)

tmp9158 := PrimEqual(symput, tmp9157)

var ifres9135 Obj

if True == tmp9158 {
tmp9155 := PrimTail(V1848)

tmp9156 := PrimIsPair(tmp9155)

var ifres9137 Obj

if True == tmp9156 {
tmp9152 := PrimTail(V1848)

tmp9153 := PrimTail(tmp9152)

tmp9154 := PrimIsPair(tmp9153)

var ifres9139 Obj

if True == tmp9154 {
tmp9148 := PrimTail(V1848)

tmp9149 := PrimTail(tmp9148)

tmp9150 := PrimTail(tmp9149)

tmp9151 := PrimIsPair(tmp9150)

var ifres9141 Obj

if True == tmp9151 {
tmp9143 := PrimTail(V1848)

tmp9144 := PrimTail(tmp9143)

tmp9145 := PrimTail(tmp9144)

tmp9146 := PrimTail(tmp9145)

tmp9147 := PrimEqual(Nil, tmp9146)

var ifres9142 Obj

if True == tmp9147 {
ifres9142 = True


} else {
ifres9142 = False


}

ifres9141 = ifres9142


} else {
ifres9141 = False


}

var ifres9140 Obj

if True == ifres9141 {
ifres9140 = True


} else {
ifres9140 = False


}

ifres9139 = ifres9140


} else {
ifres9139 = False


}

var ifres9138 Obj

if True == ifres9139 {
ifres9138 = True


} else {
ifres9138 = False


}

ifres9137 = ifres9138


} else {
ifres9137 = False


}

var ifres9136 Obj

if True == ifres9137 {
ifres9136 = True


} else {
ifres9136 = False


}

ifres9135 = ifres9136


} else {
ifres9135 = False


}

var ifres9134 Obj

if True == ifres9135 {
ifres9134 = True


} else {
ifres9134 = False


}

ifres9133 = ifres9134


} else {
ifres9133 = False


}

if True == ifres9133 {
tmp9055 := PrimTail(V1848)

tmp9056 := PrimHead(tmp9055)

tmp9057 := PrimTail(V1848)

tmp9058 := PrimTail(tmp9057)

tmp9059 := PrimHead(tmp9058)

tmp9060 := PrimTail(V1848)

tmp9061 := PrimTail(tmp9060)

tmp9062 := PrimTail(tmp9061)

tmp9063 := PrimHead(tmp9062)

tmp9064 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp9065 := PrimCons(symvalue, tmp9064)

tmp9066 := PrimCons(tmp9065, Nil)

tmp9067 := PrimCons(tmp9063, tmp9066)

tmp9068 := PrimCons(tmp9059, tmp9067)

tmp9069 := PrimCons(tmp9056, tmp9068)

__e.Return(PrimCons(symput, tmp9069))
return


} else {
tmp9131 := PrimIsPair(V1848)

var ifres9112 Obj

if True == tmp9131 {
tmp9129 := PrimHead(V1848)

tmp9130 := PrimEqual(symget, tmp9129)

var ifres9114 Obj

if True == tmp9130 {
tmp9127 := PrimTail(V1848)

tmp9128 := PrimIsPair(tmp9127)

var ifres9116 Obj

if True == tmp9128 {
tmp9124 := PrimTail(V1848)

tmp9125 := PrimTail(tmp9124)

tmp9126 := PrimIsPair(tmp9125)

var ifres9118 Obj

if True == tmp9126 {
tmp9120 := PrimTail(V1848)

tmp9121 := PrimTail(tmp9120)

tmp9122 := PrimTail(tmp9121)

tmp9123 := PrimEqual(Nil, tmp9122)

var ifres9119 Obj

if True == tmp9123 {
ifres9119 = True


} else {
ifres9119 = False


}

ifres9118 = ifres9119


} else {
ifres9118 = False


}

var ifres9117 Obj

if True == ifres9118 {
ifres9117 = True


} else {
ifres9117 = False


}

ifres9116 = ifres9117


} else {
ifres9116 = False


}

var ifres9115 Obj

if True == ifres9116 {
ifres9115 = True


} else {
ifres9115 = False


}

ifres9114 = ifres9115


} else {
ifres9114 = False


}

var ifres9113 Obj

if True == ifres9114 {
ifres9113 = True


} else {
ifres9113 = False


}

ifres9112 = ifres9113


} else {
ifres9112 = False


}

if True == ifres9112 {
tmp9070 := PrimTail(V1848)

tmp9071 := PrimHead(tmp9070)

tmp9072 := PrimTail(V1848)

tmp9073 := PrimTail(tmp9072)

tmp9074 := PrimHead(tmp9073)

tmp9075 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp9076 := PrimCons(symvalue, tmp9075)

tmp9077 := PrimCons(tmp9076, Nil)

tmp9078 := PrimCons(tmp9074, tmp9077)

tmp9079 := PrimCons(tmp9071, tmp9078)

__e.Return(PrimCons(symget, tmp9079))
return


} else {
tmp9110 := PrimIsPair(V1848)

var ifres9091 Obj

if True == tmp9110 {
tmp9108 := PrimHead(V1848)

tmp9109 := PrimEqual(symunput, tmp9108)

var ifres9093 Obj

if True == tmp9109 {
tmp9106 := PrimTail(V1848)

tmp9107 := PrimIsPair(tmp9106)

var ifres9095 Obj

if True == tmp9107 {
tmp9103 := PrimTail(V1848)

tmp9104 := PrimTail(tmp9103)

tmp9105 := PrimIsPair(tmp9104)

var ifres9097 Obj

if True == tmp9105 {
tmp9099 := PrimTail(V1848)

tmp9100 := PrimTail(tmp9099)

tmp9101 := PrimTail(tmp9100)

tmp9102 := PrimEqual(Nil, tmp9101)

var ifres9098 Obj

if True == tmp9102 {
ifres9098 = True


} else {
ifres9098 = False


}

ifres9097 = ifres9098


} else {
ifres9097 = False


}

var ifres9096 Obj

if True == ifres9097 {
ifres9096 = True


} else {
ifres9096 = False


}

ifres9095 = ifres9096


} else {
ifres9095 = False


}

var ifres9094 Obj

if True == ifres9095 {
ifres9094 = True


} else {
ifres9094 = False


}

ifres9093 = ifres9094


} else {
ifres9093 = False


}

var ifres9092 Obj

if True == ifres9093 {
ifres9092 = True


} else {
ifres9092 = False


}

ifres9091 = ifres9092


} else {
ifres9091 = False


}

if True == ifres9091 {
tmp9080 := PrimTail(V1848)

tmp9081 := PrimHead(tmp9080)

tmp9082 := PrimTail(V1848)

tmp9083 := PrimTail(tmp9082)

tmp9084 := PrimHead(tmp9083)

tmp9085 := PrimCons(sym_dproperty_1vector_d, Nil)

tmp9086 := PrimCons(symvalue, tmp9085)

tmp9087 := PrimCons(tmp9086, Nil)

tmp9088 := PrimCons(tmp9084, tmp9087)

tmp9089 := PrimCons(tmp9081, tmp9088)

__e.Return(PrimCons(symunput, tmp9089))
return


} else {
__e.Return(V1848)
return
}


}


}


}, 1)

tmp9160 := Call(__e, ns2_1set, symshen_4put_cget_1macro, tmp9054)


_ = tmp9160

tmp9161 := MakeNative(func(__e *ControlFlow) {
V1849 := __e.Get(1)
_ = V1849
tmp9162 := PrimValue(sym_dmacros_d)

tmp9163 := Call(__e, PrimFunc(symassoc), V1849, tmp9162)


tmp9164 := PrimValue(sym_dmacros_d)

tmp9165 := Call(__e, PrimFunc(symremove), tmp9163, tmp9164)


tmp9166 := PrimSet(sym_dmacros_d, tmp9165)

_ = tmp9166

__e.Return(V1849)
return


}, 1)

__e.TailApply(ns2_1set, symundefmacro, tmp9161)
return




}, 0)

