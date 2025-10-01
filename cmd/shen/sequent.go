package main

import . "github.com/tiancaiamao/shen-go/kl"

var SequentMain = MakeNative(func(__e *ControlFlow) {
tmp11442 := MakeNative(func(__e *ControlFlow) {
V3119 := __e.Get(1)
_ = V3119
tmp11443 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11445 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11445 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11463 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V3119)


var ifres11446 Obj

if True == tmp11463 {
tmp11447 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp11448 := MakeNative(func(__e *ControlFlow) {
News2980 := __e.Get(1)
_ = News2980
tmp11449 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5datatype_1rules_6 := __e.Get(1)
_ = Parseshen_4_5datatype_1rules_6
tmp11457 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5datatype_1rules_6)


if True == tmp11457 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11450 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5datatype_1rules_6)


tmp11451 := MakeNative(func(__e *ControlFlow) {
Prolog := __e.Get(1)
_ = Prolog
tmp11452 := Call(__e, PrimFunc(symfn), D)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), D, tmp11452)
return


}, 1)

tmp11453 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5datatype_1rules_6)


tmp11454 := Call(__e, PrimFunc(symshen_4rules_1_6prolog), D, tmp11453)


tmp11455 := Call(__e, tmp11451, tmp11454)


__e.TailApply(PrimFunc(symshen_4comb), tmp11450, tmp11455)
return


}


}, 1)

tmp11458 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), News2980)


__e.TailApply(tmp11449, tmp11458)
return


}, 1)

tmp11459 := Call(__e, PrimFunc(symshen_4tls), V3119)


__e.TailApply(tmp11448, tmp11459)
return


}, 1)

tmp11460 := Call(__e, PrimFunc(symshen_4hds), V3119)


tmp11461 := Call(__e, tmp11447, tmp11460)


ifres11446 = tmp11461


} else {
tmp11462 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11446 = tmp11462


}

__e.TailApply(tmp11443, ifres11446)
return


}, 1)

tmp11464 := Call(__e, ns2_1set, symshen_4_5datatype_6, tmp11442)


_ = tmp11464

tmp11465 := MakeNative(func(__e *ControlFlow) {
V3120 := __e.Get(1)
_ = V3120
tmp11466 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11484 {
tmp11467 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11469 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11469 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11470 := MakeNative(func(__e *ControlFlow) {
Parse_5_b_6 := __e.Get(1)
_ = Parse_5_b_6
tmp11480 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5_b_6)


if True == tmp11480 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11471 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5_b_6)


tmp11477 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)


tmp11478 := Call(__e, PrimFunc(symempty_2), tmp11477)


var ifres11472 Obj

if True == tmp11478 {
ifres11472 = Nil


} else {
tmp11473 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)


tmp11474 := Call(__e, PrimFunc(symshen_4app), tmp11473, MakeString("\n ..."), symshen_4r)


tmp11475 := PrimStringConcat(MakeString("datatype syntax error here:\n "), tmp11474)

tmp11476 := PrimSimpleError(tmp11475)

ifres11472 = tmp11476


}

__e.TailApply(PrimFunc(symshen_4comb), tmp11471, ifres11472)
return


}


}, 1)

tmp11481 := Call(__e, PrimFunc(sym_5_b_6), V3120)


tmp11482 := Call(__e, tmp11470, tmp11481)


__e.TailApply(tmp11467, tmp11482)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11485 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5datatype_1rule_6 := __e.Get(1)
_ = Parseshen_4_5datatype_1rule_6
tmp11495 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5datatype_1rule_6)


if True == tmp11495 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11486 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5datatype_1rules_6 := __e.Get(1)
_ = Parseshen_4_5datatype_1rules_6
tmp11492 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5datatype_1rules_6)


if True == tmp11492 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11487 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5datatype_1rules_6)


tmp11488 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5datatype_1rule_6)


tmp11489 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5datatype_1rules_6)


tmp11490 := Call(__e, PrimFunc(symappend), tmp11488, tmp11489)


__e.TailApply(PrimFunc(symshen_4comb), tmp11487, tmp11490)
return


}


}, 1)

tmp11493 := Call(__e, PrimFunc(symshen_4_5datatype_1rules_6), Parseshen_4_5datatype_1rule_6)


__e.TailApply(tmp11486, tmp11493)
return


}


}, 1)

tmp11496 := Call(__e, PrimFunc(symshen_4_5datatype_1rule_6), V3120)


tmp11497 := Call(__e, tmp11485, tmp11496)


__e.TailApply(tmp11466, tmp11497)
return


}, 1)

tmp11498 := Call(__e, ns2_1set, symshen_4_5datatype_1rules_6, tmp11465)


_ = tmp11498

tmp11499 := MakeNative(func(__e *ControlFlow) {
V3121 := __e.Get(1)
_ = V3121
tmp11500 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11512 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11512 {
tmp11501 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11503 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11503 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11504 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5double_6 := __e.Get(1)
_ = Parseshen_4_5double_6
tmp11508 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5double_6)


if True == tmp11508 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11505 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5double_6)


tmp11506 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5double_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp11505, tmp11506)
return


}


}, 1)

tmp11509 := Call(__e, PrimFunc(symshen_4_5double_6), V3121)


tmp11510 := Call(__e, tmp11504, tmp11509)


__e.TailApply(tmp11501, tmp11510)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11513 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5single_6 := __e.Get(1)
_ = Parseshen_4_5single_6
tmp11517 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5single_6)


if True == tmp11517 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11514 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5single_6)


tmp11515 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5single_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp11514, tmp11515)
return


}


}, 1)

tmp11518 := Call(__e, PrimFunc(symshen_4_5single_6), V3121)


tmp11519 := Call(__e, tmp11513, tmp11518)


__e.TailApply(tmp11500, tmp11519)
return


}, 1)

tmp11520 := Call(__e, ns2_1set, symshen_4_5datatype_1rule_6, tmp11499)


_ = tmp11520

tmp11521 := MakeNative(func(__e *ControlFlow) {
V3122 := __e.Get(1)
_ = V3122
tmp11522 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11524 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11524 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11525 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sides_6 := __e.Get(1)
_ = Parseshen_4_5sides_6
tmp11551 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sides_6)


if True == tmp11551 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11526 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5prems_6 := __e.Get(1)
_ = Parseshen_4_5prems_6
tmp11548 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5prems_6)


if True == tmp11548 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11527 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sng_6 := __e.Get(1)
_ = Parseshen_4_5sng_6
tmp11545 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sng_6)


if True == tmp11545 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11528 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5conc_6 := __e.Get(1)
_ = Parseshen_4_5conc_6
tmp11542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5conc_6)


if True == tmp11542 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11529 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sc_6 := __e.Get(1)
_ = Parseshen_4_5sc_6
tmp11539 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)


if True == tmp11539 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11530 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sc_6)


tmp11531 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5sides_6)


tmp11532 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5prems_6)


tmp11533 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5conc_6)


tmp11534 := PrimCons(tmp11533, Nil)

tmp11535 := PrimCons(tmp11532, tmp11534)

tmp11536 := PrimCons(tmp11531, tmp11535)

tmp11537 := PrimCons(tmp11536, Nil)

__e.TailApply(PrimFunc(symshen_4comb), tmp11530, tmp11537)
return


}


}, 1)

tmp11540 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5conc_6)


__e.TailApply(tmp11529, tmp11540)
return


}


}, 1)

tmp11543 := Call(__e, PrimFunc(symshen_4_5conc_6), Parseshen_4_5sng_6)


__e.TailApply(tmp11528, tmp11543)
return


}


}, 1)

tmp11546 := Call(__e, PrimFunc(symshen_4_5sng_6), Parseshen_4_5prems_6)


__e.TailApply(tmp11527, tmp11546)
return


}


}, 1)

tmp11549 := Call(__e, PrimFunc(symshen_4_5prems_6), Parseshen_4_5sides_6)


__e.TailApply(tmp11526, tmp11549)
return


}


}, 1)

tmp11552 := Call(__e, PrimFunc(symshen_4_5sides_6), V3122)


tmp11553 := Call(__e, tmp11525, tmp11552)


__e.TailApply(tmp11522, tmp11553)
return


}, 1)

tmp11554 := Call(__e, ns2_1set, symshen_4_5single_6, tmp11521)


_ = tmp11554

tmp11555 := MakeNative(func(__e *ControlFlow) {
V3123 := __e.Get(1)
_ = V3123
tmp11556 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11558 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11558 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11559 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sides_6 := __e.Get(1)
_ = Parseshen_4_5sides_6
tmp11584 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sides_6)


if True == tmp11584 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11560 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formulae_6 := __e.Get(1)
_ = Parseshen_4_5formulae_6
tmp11581 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formulae_6)


if True == tmp11581 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11561 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5dbl_6 := __e.Get(1)
_ = Parseshen_4_5dbl_6
tmp11578 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5dbl_6)


if True == tmp11578 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11562 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11575 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11563 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sc_6 := __e.Get(1)
_ = Parseshen_4_5sc_6
tmp11572 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)


if True == tmp11572 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11564 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sc_6)


tmp11565 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5sides_6)


tmp11566 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formulae_6)


tmp11567 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11568 := PrimCons(tmp11567, Nil)

tmp11569 := PrimCons(Nil, tmp11568)

tmp11570 := Call(__e, PrimFunc(symshen_4lr_1rule), tmp11565, tmp11566, tmp11569)


__e.TailApply(PrimFunc(symshen_4comb), tmp11564, tmp11570)
return


}


}, 1)

tmp11573 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5formula_6)


__e.TailApply(tmp11563, tmp11573)
return


}


}, 1)

tmp11576 := Call(__e, PrimFunc(symshen_4_5formula_6), Parseshen_4_5dbl_6)


__e.TailApply(tmp11562, tmp11576)
return


}


}, 1)

tmp11579 := Call(__e, PrimFunc(symshen_4_5dbl_6), Parseshen_4_5formulae_6)


__e.TailApply(tmp11561, tmp11579)
return


}


}, 1)

tmp11582 := Call(__e, PrimFunc(symshen_4_5formulae_6), Parseshen_4_5sides_6)


__e.TailApply(tmp11560, tmp11582)
return


}


}, 1)

tmp11585 := Call(__e, PrimFunc(symshen_4_5sides_6), V3123)


tmp11586 := Call(__e, tmp11559, tmp11585)


__e.TailApply(tmp11556, tmp11586)
return


}, 1)

tmp11587 := Call(__e, ns2_1set, symshen_4_5double_6, tmp11555)


_ = tmp11587

tmp11588 := MakeNative(func(__e *ControlFlow) {
V3124 := __e.Get(1)
_ = V3124
tmp11589 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11608 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11608 {
tmp11590 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11592 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11592 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11593 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11604 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11604 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11594 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sc_6 := __e.Get(1)
_ = Parseshen_4_5sc_6
tmp11601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)


if True == tmp11601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11595 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sc_6)


tmp11596 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11597 := PrimCons(tmp11596, Nil)

tmp11598 := PrimCons(Nil, tmp11597)

tmp11599 := PrimCons(tmp11598, Nil)

__e.TailApply(PrimFunc(symshen_4comb), tmp11595, tmp11599)
return


}


}, 1)

tmp11602 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5formula_6)


__e.TailApply(tmp11594, tmp11602)
return


}


}, 1)

tmp11605 := Call(__e, PrimFunc(symshen_4_5formula_6), V3124)


tmp11606 := Call(__e, tmp11593, tmp11605)


__e.TailApply(tmp11590, tmp11606)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11609 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11625 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11625 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11610 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sc_6 := __e.Get(1)
_ = Parseshen_4_5sc_6
tmp11622 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)


if True == tmp11622 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11611 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formulae_6 := __e.Get(1)
_ = Parseshen_4_5formulae_6
tmp11619 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formulae_6)


if True == tmp11619 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11612 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5formulae_6)


tmp11613 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11614 := PrimCons(tmp11613, Nil)

tmp11615 := PrimCons(Nil, tmp11614)

tmp11616 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formulae_6)


tmp11617 := PrimCons(tmp11615, tmp11616)

__e.TailApply(PrimFunc(symshen_4comb), tmp11612, tmp11617)
return


}


}, 1)

tmp11620 := Call(__e, PrimFunc(symshen_4_5formulae_6), Parseshen_4_5sc_6)


__e.TailApply(tmp11611, tmp11620)
return


}


}, 1)

tmp11623 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5formula_6)


__e.TailApply(tmp11610, tmp11623)
return


}


}, 1)

tmp11626 := Call(__e, PrimFunc(symshen_4_5formula_6), V3124)


tmp11627 := Call(__e, tmp11609, tmp11626)


__e.TailApply(tmp11589, tmp11627)
return


}, 1)

tmp11628 := Call(__e, ns2_1set, symshen_4_5formulae_6, tmp11588)


_ = tmp11628

tmp11629 := MakeNative(func(__e *ControlFlow) {
V3125 := __e.Get(1)
_ = V3125
tmp11630 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11644 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11644 {
tmp11631 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11633 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11633 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11634 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11640 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11640 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11635 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5formula_6)


tmp11636 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11637 := PrimCons(tmp11636, Nil)

tmp11638 := PrimCons(Nil, tmp11637)

__e.TailApply(PrimFunc(symshen_4comb), tmp11635, tmp11638)
return


}


}, 1)

tmp11641 := Call(__e, PrimFunc(symshen_4_5formula_6), V3125)


tmp11642 := Call(__e, tmp11634, tmp11641)


__e.TailApply(tmp11631, tmp11642)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11645 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5ass_6 := __e.Get(1)
_ = Parseshen_4_5ass_6
tmp11660 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5ass_6)


if True == tmp11660 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11658 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5ass_6, sym_6_6)


if True == tmp11658 {
tmp11646 := MakeNative(func(__e *ControlFlow) {
News2987 := __e.Get(1)
_ = News2987
tmp11647 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11654 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11654 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11648 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5formula_6)


tmp11649 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5ass_6)


tmp11650 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11651 := PrimCons(tmp11650, Nil)

tmp11652 := PrimCons(tmp11649, tmp11651)

__e.TailApply(PrimFunc(symshen_4comb), tmp11648, tmp11652)
return


}


}, 1)

tmp11655 := Call(__e, PrimFunc(symshen_4_5formula_6), News2987)


__e.TailApply(tmp11647, tmp11655)
return


}, 1)

tmp11656 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5ass_6)


__e.TailApply(tmp11646, tmp11656)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp11661 := Call(__e, PrimFunc(symshen_4_5ass_6), V3125)


tmp11662 := Call(__e, tmp11645, tmp11661)


__e.TailApply(tmp11630, tmp11662)
return


}, 1)

tmp11663 := Call(__e, ns2_1set, symshen_4_5conc_6, tmp11629)


_ = tmp11663

tmp11664 := MakeNative(func(__e *ControlFlow) {
V3126 := __e.Get(1)
_ = V3126
tmp11665 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11676 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11676 {
tmp11666 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11668 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11668 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11669 := MakeNative(func(__e *ControlFlow) {
Parse_5e_6 := __e.Get(1)
_ = Parse_5e_6
tmp11672 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)


if True == tmp11672 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11670 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp11670, Nil)
return


}


}, 1)

tmp11673 := Call(__e, PrimFunc(sym_5e_6), V3126)


tmp11674 := Call(__e, tmp11669, tmp11673)


__e.TailApply(tmp11666, tmp11674)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11677 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5prem_6 := __e.Get(1)
_ = Parseshen_4_5prem_6
tmp11691 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5prem_6)


if True == tmp11691 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11678 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sc_6 := __e.Get(1)
_ = Parseshen_4_5sc_6
tmp11688 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sc_6)


if True == tmp11688 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11679 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5prems_6 := __e.Get(1)
_ = Parseshen_4_5prems_6
tmp11685 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5prems_6)


if True == tmp11685 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11680 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5prems_6)


tmp11681 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5prem_6)


tmp11682 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5prems_6)


tmp11683 := PrimCons(tmp11681, tmp11682)

__e.TailApply(PrimFunc(symshen_4comb), tmp11680, tmp11683)
return


}


}, 1)

tmp11686 := Call(__e, PrimFunc(symshen_4_5prems_6), Parseshen_4_5sc_6)


__e.TailApply(tmp11679, tmp11686)
return


}


}, 1)

tmp11689 := Call(__e, PrimFunc(symshen_4_5sc_6), Parseshen_4_5prem_6)


__e.TailApply(tmp11678, tmp11689)
return


}


}, 1)

tmp11692 := Call(__e, PrimFunc(symshen_4_5prem_6), V3126)


tmp11693 := Call(__e, tmp11677, tmp11692)


__e.TailApply(tmp11665, tmp11693)
return


}, 1)

tmp11694 := Call(__e, ns2_1set, symshen_4_5prems_6, tmp11664)


_ = tmp11694

tmp11695 := MakeNative(func(__e *ControlFlow) {
V3127 := __e.Get(1)
_ = V3127
tmp11696 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11731 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11731 {
tmp11697 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11711 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11711 {
tmp11698 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11700 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11700 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11701 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11707 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11707 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11702 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5formula_6)


tmp11703 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11704 := PrimCons(tmp11703, Nil)

tmp11705 := PrimCons(Nil, tmp11704)

__e.TailApply(PrimFunc(symshen_4comb), tmp11702, tmp11705)
return


}


}, 1)

tmp11708 := Call(__e, PrimFunc(symshen_4_5formula_6), V3127)


tmp11709 := Call(__e, tmp11701, tmp11708)


__e.TailApply(tmp11698, tmp11709)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11712 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5ass_6 := __e.Get(1)
_ = Parseshen_4_5ass_6
tmp11727 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5ass_6)


if True == tmp11727 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11725 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5ass_6, sym_6_6)


if True == tmp11725 {
tmp11713 := MakeNative(func(__e *ControlFlow) {
News2991 := __e.Get(1)
_ = News2991
tmp11714 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11721 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11715 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5formula_6)


tmp11716 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5ass_6)


tmp11717 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11718 := PrimCons(tmp11717, Nil)

tmp11719 := PrimCons(tmp11716, tmp11718)

__e.TailApply(PrimFunc(symshen_4comb), tmp11715, tmp11719)
return


}


}, 1)

tmp11722 := Call(__e, PrimFunc(symshen_4_5formula_6), News2991)


__e.TailApply(tmp11714, tmp11722)
return


}, 1)

tmp11723 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5ass_6)


__e.TailApply(tmp11713, tmp11723)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp11728 := Call(__e, PrimFunc(symshen_4_5ass_6), V3127)


tmp11729 := Call(__e, tmp11712, tmp11728)


__e.TailApply(tmp11697, tmp11729)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11738 := Call(__e, PrimFunc(symshen_4_ahd_2), V3127, sym_b)


var ifres11732 Obj

if True == tmp11738 {
tmp11733 := MakeNative(func(__e *ControlFlow) {
News2990 := __e.Get(1)
_ = News2990
tmp11734 := Call(__e, PrimFunc(symshen_4in_1_6), News2990)


__e.TailApply(PrimFunc(symshen_4comb), tmp11734, sym_b)
return


}, 1)

tmp11735 := Call(__e, PrimFunc(symshen_4tls), V3127)


tmp11736 := Call(__e, tmp11733, tmp11735)


ifres11732 = tmp11736


} else {
tmp11737 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11732 = tmp11737


}

__e.TailApply(tmp11696, ifres11732)
return


}, 1)

tmp11739 := Call(__e, ns2_1set, symshen_4_5prem_6, tmp11695)


_ = tmp11739

tmp11740 := MakeNative(func(__e *ControlFlow) {
V3128 := __e.Get(1)
_ = V3128
tmp11741 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11763 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11763 {
tmp11742 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11753 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11753 {
tmp11743 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11745 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11745 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11746 := MakeNative(func(__e *ControlFlow) {
Parse_5e_6 := __e.Get(1)
_ = Parse_5e_6
tmp11749 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)


if True == tmp11749 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11747 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp11747, Nil)
return


}


}, 1)

tmp11750 := Call(__e, PrimFunc(sym_5e_6), V3128)


tmp11751 := Call(__e, tmp11746, tmp11750)


__e.TailApply(tmp11743, tmp11751)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11754 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11759 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11759 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11755 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5formula_6)


tmp11756 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11757 := PrimCons(tmp11756, Nil)

__e.TailApply(PrimFunc(symshen_4comb), tmp11755, tmp11757)
return


}


}, 1)

tmp11760 := Call(__e, PrimFunc(symshen_4_5formula_6), V3128)


tmp11761 := Call(__e, tmp11754, tmp11760)


__e.TailApply(tmp11742, tmp11761)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11764 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5formula_6 := __e.Get(1)
_ = Parseshen_4_5formula_6
tmp11778 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5formula_6)


if True == tmp11778 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11765 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5iscomma_6 := __e.Get(1)
_ = Parseshen_4_5iscomma_6
tmp11775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5iscomma_6)


if True == tmp11775 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11766 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5ass_6 := __e.Get(1)
_ = Parseshen_4_5ass_6
tmp11772 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5ass_6)


if True == tmp11772 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11767 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5ass_6)


tmp11768 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5formula_6)


tmp11769 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5ass_6)


tmp11770 := PrimCons(tmp11768, tmp11769)

__e.TailApply(PrimFunc(symshen_4comb), tmp11767, tmp11770)
return


}


}, 1)

tmp11773 := Call(__e, PrimFunc(symshen_4_5ass_6), Parseshen_4_5iscomma_6)


__e.TailApply(tmp11766, tmp11773)
return


}


}, 1)

tmp11776 := Call(__e, PrimFunc(symshen_4_5iscomma_6), Parseshen_4_5formula_6)


__e.TailApply(tmp11765, tmp11776)
return


}


}, 1)

tmp11779 := Call(__e, PrimFunc(symshen_4_5formula_6), V3128)


tmp11780 := Call(__e, tmp11764, tmp11779)


__e.TailApply(tmp11741, tmp11780)
return


}, 1)

tmp11781 := Call(__e, ns2_1set, symshen_4_5ass_6, tmp11740)


_ = tmp11781

tmp11782 := MakeNative(func(__e *ControlFlow) {
V3129 := __e.Get(1)
_ = V3129
tmp11783 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11785 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11785 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11797 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V3129)


var ifres11786 Obj

if True == tmp11797 {
tmp11787 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp11788 := MakeNative(func(__e *ControlFlow) {
News2994 := __e.Get(1)
_ = News2994
tmp11791 := PrimIntern(MakeString(","))

tmp11792 := PrimEqual(X, tmp11791)

if True == tmp11792 {
tmp11789 := Call(__e, PrimFunc(symshen_4in_1_6), News2994)


__e.TailApply(PrimFunc(symshen_4comb), tmp11789, symshen_4skip)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11793 := Call(__e, PrimFunc(symshen_4tls), V3129)


__e.TailApply(tmp11788, tmp11793)
return


}, 1)

tmp11794 := Call(__e, PrimFunc(symshen_4hds), V3129)


tmp11795 := Call(__e, tmp11787, tmp11794)


ifres11786 = tmp11795


} else {
tmp11796 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11786 = tmp11796


}

__e.TailApply(tmp11783, ifres11786)
return


}, 1)

tmp11798 := Call(__e, ns2_1set, symshen_4_5iscomma_6, tmp11782)


_ = tmp11798

tmp11799 := MakeNative(func(__e *ControlFlow) {
V3130 := __e.Get(1)
_ = V3130
tmp11800 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11812 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11812 {
tmp11801 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11803 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11803 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11804 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5expr_6 := __e.Get(1)
_ = Parseshen_4_5expr_6
tmp11808 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5expr_6)


if True == tmp11808 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11805 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5expr_6)


tmp11806 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5expr_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp11805, tmp11806)
return


}


}, 1)

tmp11809 := Call(__e, PrimFunc(symshen_4_5expr_6), V3130)


tmp11810 := Call(__e, tmp11804, tmp11809)


__e.TailApply(tmp11801, tmp11810)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11813 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5expr_6 := __e.Get(1)
_ = Parseshen_4_5expr_6
tmp11832 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5expr_6)


if True == tmp11832 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11814 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5iscolon_6 := __e.Get(1)
_ = Parseshen_4_5iscolon_6
tmp11829 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5iscolon_6)


if True == tmp11829 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11815 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5type_6 := __e.Get(1)
_ = Parseshen_4_5type_6
tmp11826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5type_6)


if True == tmp11826 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11816 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5type_6)


tmp11817 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5expr_6)


tmp11818 := Call(__e, PrimFunc(symshen_4curry), tmp11817)


tmp11819 := PrimIntern(MakeString(":"))

tmp11820 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5type_6)


tmp11821 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp11820)


tmp11822 := PrimCons(tmp11821, Nil)

tmp11823 := PrimCons(tmp11819, tmp11822)

tmp11824 := PrimCons(tmp11818, tmp11823)

__e.TailApply(PrimFunc(symshen_4comb), tmp11816, tmp11824)
return


}


}, 1)

tmp11827 := Call(__e, PrimFunc(symshen_4_5type_6), Parseshen_4_5iscolon_6)


__e.TailApply(tmp11815, tmp11827)
return


}


}, 1)

tmp11830 := Call(__e, PrimFunc(symshen_4_5iscolon_6), Parseshen_4_5expr_6)


__e.TailApply(tmp11814, tmp11830)
return


}


}, 1)

tmp11833 := Call(__e, PrimFunc(symshen_4_5expr_6), V3130)


tmp11834 := Call(__e, tmp11813, tmp11833)


__e.TailApply(tmp11800, tmp11834)
return


}, 1)

tmp11835 := Call(__e, ns2_1set, symshen_4_5formula_6, tmp11799)


_ = tmp11835

tmp11836 := MakeNative(func(__e *ControlFlow) {
V3131 := __e.Get(1)
_ = V3131
tmp11837 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11839 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11851 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V3131)


var ifres11840 Obj

if True == tmp11851 {
tmp11841 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp11842 := MakeNative(func(__e *ControlFlow) {
News2997 := __e.Get(1)
_ = News2997
tmp11845 := PrimIntern(MakeString(":"))

tmp11846 := PrimEqual(X, tmp11845)

if True == tmp11846 {
tmp11843 := Call(__e, PrimFunc(symshen_4in_1_6), News2997)


__e.TailApply(PrimFunc(symshen_4comb), tmp11843, symshen_4skip)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11847 := Call(__e, PrimFunc(symshen_4tls), V3131)


__e.TailApply(tmp11842, tmp11847)
return


}, 1)

tmp11848 := Call(__e, PrimFunc(symshen_4hds), V3131)


tmp11849 := Call(__e, tmp11841, tmp11848)


ifres11840 = tmp11849


} else {
tmp11850 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11840 = tmp11850


}

__e.TailApply(tmp11837, ifres11840)
return


}, 1)

tmp11852 := Call(__e, ns2_1set, symshen_4_5iscolon_6, tmp11836)


_ = tmp11852

tmp11853 := MakeNative(func(__e *ControlFlow) {
V3132 := __e.Get(1)
_ = V3132
tmp11854 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11865 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11865 {
tmp11855 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11857 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11857 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11858 := MakeNative(func(__e *ControlFlow) {
Parse_5e_6 := __e.Get(1)
_ = Parse_5e_6
tmp11861 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)


if True == tmp11861 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11859 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp11859, Nil)
return


}


}, 1)

tmp11862 := Call(__e, PrimFunc(sym_5e_6), V3132)


tmp11863 := Call(__e, tmp11858, tmp11862)


__e.TailApply(tmp11855, tmp11863)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11866 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5side_6 := __e.Get(1)
_ = Parseshen_4_5side_6
tmp11876 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5side_6)


if True == tmp11876 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11867 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5sides_6 := __e.Get(1)
_ = Parseshen_4_5sides_6
tmp11873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sides_6)


if True == tmp11873 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp11868 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sides_6)


tmp11869 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5side_6)


tmp11870 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5sides_6)


tmp11871 := PrimCons(tmp11869, tmp11870)

__e.TailApply(PrimFunc(symshen_4comb), tmp11868, tmp11871)
return


}


}, 1)

tmp11874 := Call(__e, PrimFunc(symshen_4_5sides_6), Parseshen_4_5side_6)


__e.TailApply(tmp11867, tmp11874)
return


}


}, 1)

tmp11877 := Call(__e, PrimFunc(symshen_4_5side_6), V3132)


tmp11878 := Call(__e, tmp11866, tmp11877)


__e.TailApply(tmp11854, tmp11878)
return


}, 1)

tmp11879 := Call(__e, ns2_1set, symshen_4_5sides_6, tmp11853)


_ = tmp11879

tmp11880 := MakeNative(func(__e *ControlFlow) {
V3133 := __e.Get(1)
_ = V3133
tmp11881 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11933 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11933 {
tmp11882 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11909 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11909 {
tmp11883 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp11885 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp11885 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp11907 := Call(__e, PrimFunc(symshen_4_ahd_2), V3133, symshen_4let_b)


var ifres11886 Obj

if True == tmp11907 {
tmp11887 := MakeNative(func(__e *ControlFlow) {
News3005 := __e.Get(1)
_ = News3005
tmp11903 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3005)


if True == tmp11903 {
tmp11888 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp11889 := MakeNative(func(__e *ControlFlow) {
News3006 := __e.Get(1)
_ = News3006
tmp11899 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3006)


if True == tmp11899 {
tmp11890 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp11891 := MakeNative(func(__e *ControlFlow) {
News3007 := __e.Get(1)
_ = News3007
tmp11892 := Call(__e, PrimFunc(symshen_4in_1_6), News3007)


tmp11893 := PrimCons(Y, Nil)

tmp11894 := PrimCons(X, tmp11893)

tmp11895 := PrimCons(symshen_4let_b, tmp11894)

__e.TailApply(PrimFunc(symshen_4comb), tmp11892, tmp11895)
return


}, 1)

tmp11896 := Call(__e, PrimFunc(symshen_4tls), News3006)


__e.TailApply(tmp11891, tmp11896)
return


}, 1)

tmp11897 := Call(__e, PrimFunc(symshen_4hds), News3006)


__e.TailApply(tmp11890, tmp11897)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11900 := Call(__e, PrimFunc(symshen_4tls), News3005)


__e.TailApply(tmp11889, tmp11900)
return


}, 1)

tmp11901 := Call(__e, PrimFunc(symshen_4hds), News3005)


__e.TailApply(tmp11888, tmp11901)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11904 := Call(__e, PrimFunc(symshen_4tls), V3133)


tmp11905 := Call(__e, tmp11887, tmp11904)


ifres11886 = tmp11905


} else {
tmp11906 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11886 = tmp11906


}

__e.TailApply(tmp11883, ifres11886)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11931 := Call(__e, PrimFunc(symshen_4_ahd_2), V3133, symlet)


var ifres11910 Obj

if True == tmp11931 {
tmp11911 := MakeNative(func(__e *ControlFlow) {
News3002 := __e.Get(1)
_ = News3002
tmp11927 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3002)


if True == tmp11927 {
tmp11912 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp11913 := MakeNative(func(__e *ControlFlow) {
News3003 := __e.Get(1)
_ = News3003
tmp11923 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3003)


if True == tmp11923 {
tmp11914 := MakeNative(func(__e *ControlFlow) {
Y := __e.Get(1)
_ = Y
tmp11915 := MakeNative(func(__e *ControlFlow) {
News3004 := __e.Get(1)
_ = News3004
tmp11916 := Call(__e, PrimFunc(symshen_4in_1_6), News3004)


tmp11917 := PrimCons(Y, Nil)

tmp11918 := PrimCons(X, tmp11917)

tmp11919 := PrimCons(symlet, tmp11918)

__e.TailApply(PrimFunc(symshen_4comb), tmp11916, tmp11919)
return


}, 1)

tmp11920 := Call(__e, PrimFunc(symshen_4tls), News3003)


__e.TailApply(tmp11915, tmp11920)
return


}, 1)

tmp11921 := Call(__e, PrimFunc(symshen_4hds), News3003)


__e.TailApply(tmp11914, tmp11921)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11924 := Call(__e, PrimFunc(symshen_4tls), News3002)


__e.TailApply(tmp11913, tmp11924)
return


}, 1)

tmp11925 := Call(__e, PrimFunc(symshen_4hds), News3002)


__e.TailApply(tmp11912, tmp11925)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11928 := Call(__e, PrimFunc(symshen_4tls), V3133)


tmp11929 := Call(__e, tmp11911, tmp11928)


ifres11910 = tmp11929


} else {
tmp11930 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11910 = tmp11930


}

__e.TailApply(tmp11882, ifres11910)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp11948 := Call(__e, PrimFunc(symshen_4_ahd_2), V3133, symif)


var ifres11934 Obj

if True == tmp11948 {
tmp11935 := MakeNative(func(__e *ControlFlow) {
News3000 := __e.Get(1)
_ = News3000
tmp11944 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News3000)


if True == tmp11944 {
tmp11936 := MakeNative(func(__e *ControlFlow) {
P := __e.Get(1)
_ = P
tmp11937 := MakeNative(func(__e *ControlFlow) {
News3001 := __e.Get(1)
_ = News3001
tmp11938 := Call(__e, PrimFunc(symshen_4in_1_6), News3001)


tmp11939 := PrimCons(P, Nil)

tmp11940 := PrimCons(symif, tmp11939)

__e.TailApply(PrimFunc(symshen_4comb), tmp11938, tmp11940)
return


}, 1)

tmp11941 := Call(__e, PrimFunc(symshen_4tls), News3000)


__e.TailApply(tmp11937, tmp11941)
return


}, 1)

tmp11942 := Call(__e, PrimFunc(symshen_4hds), News3000)


__e.TailApply(tmp11936, tmp11942)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp11945 := Call(__e, PrimFunc(symshen_4tls), V3133)


tmp11946 := Call(__e, tmp11935, tmp11945)


ifres11934 = tmp11946


} else {
tmp11947 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres11934 = tmp11947


}

__e.TailApply(tmp11881, ifres11934)
return


}, 1)

tmp11949 := Call(__e, ns2_1set, symshen_4_5side_6, tmp11880)


_ = tmp11949

tmp11950 := MakeNative(func(__e *ControlFlow) {
V3140 := __e.Get(1)
_ = V3140
V3141 := __e.Get(2)
_ = V3141
V3142 := __e.Get(3)
_ = V3142
tmp11986 := PrimIsPair(V3142)

var ifres11973 Obj

if True == tmp11986 {
tmp11984 := PrimHead(V3142)

tmp11985 := PrimEqual(Nil, tmp11984)

var ifres11975 Obj

if True == tmp11985 {
tmp11982 := PrimTail(V3142)

tmp11983 := PrimIsPair(tmp11982)

var ifres11977 Obj

if True == tmp11983 {
tmp11979 := PrimTail(V3142)

tmp11980 := PrimTail(tmp11979)

tmp11981 := PrimEqual(Nil, tmp11980)

var ifres11978 Obj

if True == tmp11981 {
ifres11978 = True


} else {
ifres11978 = False


}

ifres11977 = ifres11978


} else {
ifres11977 = False


}

var ifres11976 Obj

if True == ifres11977 {
ifres11976 = True


} else {
ifres11976 = False


}

ifres11975 = ifres11976


} else {
ifres11975 = False


}

var ifres11974 Obj

if True == ifres11975 {
ifres11974 = True


} else {
ifres11974 = False


}

ifres11973 = ifres11974


} else {
ifres11973 = False


}

if True == ifres11973 {
tmp11951 := MakeNative(func(__e *ControlFlow) {
P := __e.Get(1)
_ = P
tmp11952 := MakeNative(func(__e *ControlFlow) {
LConc := __e.Get(1)
_ = LConc
tmp11953 := MakeNative(func(__e *ControlFlow) {
LPrem := __e.Get(1)
_ = LPrem
tmp11954 := MakeNative(func(__e *ControlFlow) {
Left := __e.Get(1)
_ = Left
tmp11955 := MakeNative(func(__e *ControlFlow) {
Right := __e.Get(1)
_ = Right
tmp11956 := PrimCons(Left, Nil)

__e.Return(PrimCons(Right, tmp11956))
return


}, 1)

tmp11957 := PrimCons(V3142, Nil)

tmp11958 := PrimCons(V3141, tmp11957)

tmp11959 := PrimCons(V3140, tmp11958)

__e.TailApply(tmp11955, tmp11959)
return


}, 1)

tmp11960 := PrimCons(LPrem, Nil)

tmp11961 := PrimCons(LConc, Nil)

tmp11962 := PrimCons(tmp11960, tmp11961)

tmp11963 := PrimCons(V3140, tmp11962)

__e.TailApply(tmp11954, tmp11963)
return


}, 1)

tmp11964 := Call(__e, PrimFunc(symshen_4coll_1formulae), V3141)


tmp11965 := PrimCons(P, Nil)

tmp11966 := PrimCons(tmp11964, tmp11965)

__e.TailApply(tmp11953, tmp11966)
return


}, 1)

tmp11967 := PrimTail(V3142)

tmp11968 := PrimCons(P, Nil)

tmp11969 := PrimCons(tmp11967, tmp11968)

__e.TailApply(tmp11952, tmp11969)
return


}, 1)

tmp11970 := Call(__e, PrimFunc(symprotect), symP)


tmp11971 := Call(__e, PrimFunc(symgensym), tmp11970)


__e.TailApply(tmp11951, tmp11971)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.lr-rule")))
return
}


}, 3)

tmp11987 := Call(__e, ns2_1set, symshen_4lr_1rule, tmp11950)


_ = tmp11987

tmp11988 := MakeNative(func(__e *ControlFlow) {
V3145 := __e.Get(1)
_ = V3145
tmp12017 := PrimEqual(Nil, V3145)

if True == tmp12017 {
__e.Return(Nil)
return
} else {
tmp12015 := PrimIsPair(V3145)

var ifres11995 Obj

if True == tmp12015 {
tmp12013 := PrimHead(V3145)

tmp12014 := PrimIsPair(tmp12013)

var ifres11997 Obj

if True == tmp12014 {
tmp12010 := PrimHead(V3145)

tmp12011 := PrimHead(tmp12010)

tmp12012 := PrimEqual(Nil, tmp12011)

var ifres11999 Obj

if True == tmp12012 {
tmp12007 := PrimHead(V3145)

tmp12008 := PrimTail(tmp12007)

tmp12009 := PrimIsPair(tmp12008)

var ifres12001 Obj

if True == tmp12009 {
tmp12003 := PrimHead(V3145)

tmp12004 := PrimTail(tmp12003)

tmp12005 := PrimTail(tmp12004)

tmp12006 := PrimEqual(Nil, tmp12005)

var ifres12002 Obj

if True == tmp12006 {
ifres12002 = True


} else {
ifres12002 = False


}

ifres12001 = ifres12002


} else {
ifres12001 = False


}

var ifres12000 Obj

if True == ifres12001 {
ifres12000 = True


} else {
ifres12000 = False


}

ifres11999 = ifres12000


} else {
ifres11999 = False


}

var ifres11998 Obj

if True == ifres11999 {
ifres11998 = True


} else {
ifres11998 = False


}

ifres11997 = ifres11998


} else {
ifres11997 = False


}

var ifres11996 Obj

if True == ifres11997 {
ifres11996 = True


} else {
ifres11996 = False


}

ifres11995 = ifres11996


} else {
ifres11995 = False


}

if True == ifres11995 {
tmp11989 := PrimHead(V3145)

tmp11990 := PrimTail(tmp11989)

tmp11991 := PrimHead(tmp11990)

tmp11992 := PrimTail(V3145)

tmp11993 := Call(__e, PrimFunc(symshen_4coll_1formulae), tmp11992)


__e.Return(PrimCons(tmp11991, tmp11993))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.coll-formulae")))
return
}


}


}, 1)

tmp12018 := Call(__e, ns2_1set, symshen_4coll_1formulae, tmp11988)


_ = tmp12018

tmp12019 := MakeNative(func(__e *ControlFlow) {
V3146 := __e.Get(1)
_ = V3146
tmp12020 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp12022 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp12022 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp12035 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V3146)


var ifres12023 Obj

if True == tmp12035 {
tmp12024 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp12025 := MakeNative(func(__e *ControlFlow) {
News3009 := __e.Get(1)
_ = News3009
tmp12029 := Call(__e, PrimFunc(symshen_4key_1in_1sequent_1calculus_2), X)


tmp12030 := PrimNot(tmp12029)

if True == tmp12030 {
tmp12026 := Call(__e, PrimFunc(symshen_4in_1_6), News3009)


tmp12027 := Call(__e, PrimFunc(symmacroexpand), X)


__e.TailApply(PrimFunc(symshen_4comb), tmp12026, tmp12027)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12031 := Call(__e, PrimFunc(symshen_4tls), V3146)


__e.TailApply(tmp12025, tmp12031)
return


}, 1)

tmp12032 := Call(__e, PrimFunc(symshen_4hds), V3146)


tmp12033 := Call(__e, tmp12024, tmp12032)


ifres12023 = tmp12033


} else {
tmp12034 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12023 = tmp12034


}

__e.TailApply(tmp12020, ifres12023)
return


}, 1)

tmp12036 := Call(__e, ns2_1set, symshen_4_5expr_6, tmp12019)


_ = tmp12036

tmp12037 := MakeNative(func(__e *ControlFlow) {
V3147 := __e.Get(1)
_ = V3147
tmp12044 := PrimIntern(MakeString(";"))

tmp12045 := PrimIntern(MakeString(","))

tmp12046 := PrimIntern(MakeString(":"))

tmp12047 := PrimCons(sym_5_1_1, Nil)

tmp12048 := PrimCons(tmp12046, tmp12047)

tmp12049 := PrimCons(tmp12045, tmp12048)

tmp12050 := PrimCons(tmp12044, tmp12049)

tmp12051 := PrimCons(sym_6_6, tmp12050)

tmp12052 := Call(__e, PrimFunc(symelement_2), V3147, tmp12051)


if True == tmp12052 {
__e.Return(True)
return
} else {
tmp12042 := Call(__e, PrimFunc(symshen_4sng_2), V3147)


var ifres12039 Obj

if True == tmp12042 {
ifres12039 = True


} else {
tmp12041 := Call(__e, PrimFunc(symshen_4dbl_2), V3147)


var ifres12040 Obj

if True == tmp12041 {
ifres12040 = True


} else {
ifres12040 = False


}

ifres12039 = ifres12040


}

if True == ifres12039 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp12053 := Call(__e, ns2_1set, symshen_4key_1in_1sequent_1calculus_2, tmp12037)


_ = tmp12053

tmp12054 := MakeNative(func(__e *ControlFlow) {
V3148 := __e.Get(1)
_ = V3148
tmp12055 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp12057 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp12057 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp12058 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5expr_6 := __e.Get(1)
_ = Parseshen_4_5expr_6
tmp12062 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5expr_6)


if True == tmp12062 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp12059 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5expr_6)


tmp12060 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5expr_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp12059, tmp12060)
return


}


}, 1)

tmp12063 := Call(__e, PrimFunc(symshen_4_5expr_6), V3148)


tmp12064 := Call(__e, tmp12058, tmp12063)


__e.TailApply(tmp12055, tmp12064)
return


}, 1)

tmp12065 := Call(__e, ns2_1set, symshen_4_5type_6, tmp12054)


_ = tmp12065

tmp12066 := MakeNative(func(__e *ControlFlow) {
V3149 := __e.Get(1)
_ = V3149
tmp12067 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp12069 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp12069 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp12080 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V3149)


var ifres12070 Obj

if True == tmp12080 {
tmp12071 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp12072 := MakeNative(func(__e *ControlFlow) {
News3012 := __e.Get(1)
_ = News3012
tmp12075 := Call(__e, PrimFunc(symshen_4dbl_2), X)


if True == tmp12075 {
tmp12073 := Call(__e, PrimFunc(symshen_4in_1_6), News3012)


__e.TailApply(PrimFunc(symshen_4comb), tmp12073, X)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12076 := Call(__e, PrimFunc(symshen_4tls), V3149)


__e.TailApply(tmp12072, tmp12076)
return


}, 1)

tmp12077 := Call(__e, PrimFunc(symshen_4hds), V3149)


tmp12078 := Call(__e, tmp12071, tmp12077)


ifres12070 = tmp12078


} else {
tmp12079 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12070 = tmp12079


}

__e.TailApply(tmp12067, ifres12070)
return


}, 1)

tmp12081 := Call(__e, ns2_1set, symshen_4_5dbl_6, tmp12066)


_ = tmp12081

tmp12082 := MakeNative(func(__e *ControlFlow) {
V3150 := __e.Get(1)
_ = V3150
tmp12083 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp12085 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp12085 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp12096 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V3150)


var ifres12086 Obj

if True == tmp12096 {
tmp12087 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp12088 := MakeNative(func(__e *ControlFlow) {
News3014 := __e.Get(1)
_ = News3014
tmp12091 := Call(__e, PrimFunc(symshen_4sng_2), X)


if True == tmp12091 {
tmp12089 := Call(__e, PrimFunc(symshen_4in_1_6), News3014)


__e.TailApply(PrimFunc(symshen_4comb), tmp12089, X)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp12092 := Call(__e, PrimFunc(symshen_4tls), V3150)


__e.TailApply(tmp12088, tmp12092)
return


}, 1)

tmp12093 := Call(__e, PrimFunc(symshen_4hds), V3150)


tmp12094 := Call(__e, tmp12087, tmp12093)


ifres12086 = tmp12094


} else {
tmp12095 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres12086 = tmp12095


}

__e.TailApply(tmp12083, ifres12086)
return


}, 1)

tmp12097 := Call(__e, ns2_1set, symshen_4_5sng_6, tmp12082)


_ = tmp12097

tmp12098 := MakeNative(func(__e *ControlFlow) {
V3151 := __e.Get(1)
_ = V3151
tmp12103 := PrimIsSymbol(V3151)

if True == tmp12103 {
tmp12100 := PrimStr(V3151)

tmp12101 := Call(__e, PrimFunc(symshen_4sng_1h_2), tmp12100)


if True == tmp12101 {
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

tmp12104 := Call(__e, ns2_1set, symshen_4sng_2, tmp12098)


_ = tmp12104

tmp12105 := MakeNative(func(__e *ControlFlow) {
V3154 := __e.Get(1)
_ = V3154
tmp12114 := PrimEqual(MakeString("___"), V3154)

if True == tmp12114 {
__e.Return(True)
return
} else {
tmp12112 := Call(__e, PrimFunc(symshen_4_7string_2), V3154)


var ifres12108 Obj

if True == tmp12112 {
tmp12110 := Call(__e, PrimFunc(symhdstr), V3154)


tmp12111 := PrimEqual(MakeString("_"), tmp12110)

var ifres12109 Obj

if True == tmp12111 {
ifres12109 = True


} else {
ifres12109 = False


}

ifres12108 = ifres12109


} else {
ifres12108 = False


}

if True == ifres12108 {
tmp12106 := PrimTailString(V3154)

__e.TailApply(PrimFunc(symshen_4sng_1h_2), tmp12106)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12115 := Call(__e, ns2_1set, symshen_4sng_1h_2, tmp12105)


_ = tmp12115

tmp12116 := MakeNative(func(__e *ControlFlow) {
V3155 := __e.Get(1)
_ = V3155
tmp12121 := PrimIsSymbol(V3155)

if True == tmp12121 {
tmp12118 := PrimStr(V3155)

tmp12119 := Call(__e, PrimFunc(symshen_4dbl_1h_2), tmp12118)


if True == tmp12119 {
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

tmp12122 := Call(__e, ns2_1set, symshen_4dbl_2, tmp12116)


_ = tmp12122

tmp12123 := MakeNative(func(__e *ControlFlow) {
V3158 := __e.Get(1)
_ = V3158
tmp12132 := PrimEqual(MakeString("==="), V3158)

if True == tmp12132 {
__e.Return(True)
return
} else {
tmp12130 := Call(__e, PrimFunc(symshen_4_7string_2), V3158)


var ifres12126 Obj

if True == tmp12130 {
tmp12128 := Call(__e, PrimFunc(symhdstr), V3158)


tmp12129 := PrimEqual(MakeString("="), tmp12128)

var ifres12127 Obj

if True == tmp12129 {
ifres12127 = True


} else {
ifres12127 = False


}

ifres12126 = ifres12127


} else {
ifres12126 = False


}

if True == ifres12126 {
tmp12124 := PrimTailString(V3158)

__e.TailApply(PrimFunc(symshen_4dbl_1h_2), tmp12124)
return


} else {
__e.Return(False)
return
}


}


}, 1)

tmp12133 := Call(__e, ns2_1set, symshen_4dbl_1h_2, tmp12123)


_ = tmp12133

tmp12134 := MakeNative(func(__e *ControlFlow) {
V3159 := __e.Get(1)
_ = V3159
V3160 := __e.Get(2)
_ = V3160
tmp12135 := PrimValue(symshen_4_ddatatypes_d)

tmp12136 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3159, V3160, tmp12135)


tmp12137 := PrimSet(symshen_4_ddatatypes_d, tmp12136)

_ = tmp12137

tmp12138 := PrimValue(symshen_4_dalldatatypes_d)

tmp12139 := Call(__e, PrimFunc(symshen_4assoc_1_6), V3159, V3160, tmp12138)


tmp12140 := PrimSet(symshen_4_dalldatatypes_d, tmp12139)

_ = tmp12140

__e.Return(V3159)
return


}, 2)

tmp12141 := Call(__e, ns2_1set, symshen_4remember_1datatype, tmp12134)


_ = tmp12141

tmp12142 := MakeNative(func(__e *ControlFlow) {
V3161 := __e.Get(1)
_ = V3161
V3162 := __e.Get(2)
_ = V3162
tmp12143 := MakeNative(func(__e *ControlFlow) {
Clauses := __e.Get(1)
_ = Clauses
tmp12144 := PrimCons(V3161, Clauses)

tmp12145 := PrimCons(symdefprolog, tmp12144)

__e.TailApply(PrimFunc(symeval), tmp12145)
return


}, 1)

tmp12146 := MakeNative(func(__e *ControlFlow) {
Rule := __e.Get(1)
_ = Rule
__e.TailApply(PrimFunc(symshen_4rule_1_6clause), Rule)
return
}, 1)

tmp12147 := Call(__e, PrimFunc(symmapcan), tmp12146, V3162)


__e.TailApply(tmp12143, tmp12147)
return


}, 2)

tmp12148 := Call(__e, ns2_1set, symshen_4rules_1_6prolog, tmp12142)


_ = tmp12148

tmp12149 := MakeNative(func(__e *ControlFlow) {
V3165 := __e.Get(1)
_ = V3165
tmp12225 := PrimIsPair(V3165)

var ifres12189 Obj

if True == tmp12225 {
tmp12223 := PrimTail(V3165)

tmp12224 := PrimIsPair(tmp12223)

var ifres12191 Obj

if True == tmp12224 {
tmp12220 := PrimTail(V3165)

tmp12221 := PrimTail(tmp12220)

tmp12222 := PrimIsPair(tmp12221)

var ifres12193 Obj

if True == tmp12222 {
tmp12216 := PrimTail(V3165)

tmp12217 := PrimTail(tmp12216)

tmp12218 := PrimHead(tmp12217)

tmp12219 := PrimIsPair(tmp12218)

var ifres12195 Obj

if True == tmp12219 {
tmp12211 := PrimTail(V3165)

tmp12212 := PrimTail(tmp12211)

tmp12213 := PrimHead(tmp12212)

tmp12214 := PrimTail(tmp12213)

tmp12215 := PrimIsPair(tmp12214)

var ifres12197 Obj

if True == tmp12215 {
tmp12205 := PrimTail(V3165)

tmp12206 := PrimTail(tmp12205)

tmp12207 := PrimHead(tmp12206)

tmp12208 := PrimTail(tmp12207)

tmp12209 := PrimTail(tmp12208)

tmp12210 := PrimEqual(Nil, tmp12209)

var ifres12199 Obj

if True == tmp12210 {
tmp12201 := PrimTail(V3165)

tmp12202 := PrimTail(tmp12201)

tmp12203 := PrimTail(tmp12202)

tmp12204 := PrimEqual(Nil, tmp12203)

var ifres12200 Obj

if True == tmp12204 {
ifres12200 = True


} else {
ifres12200 = False


}

ifres12199 = ifres12200


} else {
ifres12199 = False


}

var ifres12198 Obj

if True == ifres12199 {
ifres12198 = True


} else {
ifres12198 = False


}

ifres12197 = ifres12198


} else {
ifres12197 = False


}

var ifres12196 Obj

if True == ifres12197 {
ifres12196 = True


} else {
ifres12196 = False


}

ifres12195 = ifres12196


} else {
ifres12195 = False


}

var ifres12194 Obj

if True == ifres12195 {
ifres12194 = True


} else {
ifres12194 = False


}

ifres12193 = ifres12194


} else {
ifres12193 = False


}

var ifres12192 Obj

if True == ifres12193 {
ifres12192 = True


} else {
ifres12192 = False


}

ifres12191 = ifres12192


} else {
ifres12191 = False


}

var ifres12190 Obj

if True == ifres12191 {
ifres12190 = True


} else {
ifres12190 = False


}

ifres12189 = ifres12190


} else {
ifres12189 = False


}

if True == ifres12189 {
tmp12150 := MakeNative(func(__e *ControlFlow) {
Constraints := __e.Get(1)
_ = Constraints
tmp12151 := MakeNative(func(__e *ControlFlow) {
HypVs := __e.Get(1)
_ = HypVs
tmp12152 := MakeNative(func(__e *ControlFlow) {
Active := __e.Get(1)
_ = Active
tmp12153 := MakeNative(func(__e *ControlFlow) {
Head := __e.Get(1)
_ = Head
tmp12154 := MakeNative(func(__e *ControlFlow) {
Goals := __e.Get(1)
_ = Goals
tmp12155 := PrimCons(sym_5_1_1, Nil)

tmp12156 := PrimIntern(MakeString(";"))

tmp12157 := PrimCons(tmp12156, Nil)

tmp12158 := Call(__e, PrimFunc(symappend), Goals, tmp12157)


tmp12159 := Call(__e, PrimFunc(symappend), tmp12155, tmp12158)


__e.TailApply(PrimFunc(symappend), Head, tmp12159)
return


}, 1)

tmp12160 := PrimTail(V3165)

tmp12161 := PrimTail(tmp12160)

tmp12162 := PrimHead(tmp12161)

tmp12163 := PrimHead(tmp12162)

tmp12164 := PrimHead(V3165)

tmp12165 := PrimTail(V3165)

tmp12166 := PrimHead(tmp12165)

tmp12167 := Call(__e, PrimFunc(symshen_4goals), Constraints, tmp12163, tmp12164, tmp12166, HypVs, Active)


__e.TailApply(tmp12154, tmp12167)
return


}, 1)

tmp12168 := PrimTail(V3165)

tmp12169 := PrimTail(tmp12168)

tmp12170 := PrimHead(tmp12169)

tmp12171 := PrimTail(tmp12170)

tmp12172 := PrimHead(tmp12171)

tmp12173 := Call(__e, PrimFunc(symshen_4compile_1consequent), tmp12172, HypVs)


__e.TailApply(tmp12153, tmp12173)
return


}, 1)

tmp12174 := PrimTail(V3165)

tmp12175 := PrimTail(tmp12174)

tmp12176 := PrimHead(tmp12175)

tmp12177 := PrimTail(tmp12176)

tmp12178 := PrimHead(tmp12177)

tmp12179 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp12178)


__e.TailApply(tmp12152, tmp12179)
return


}, 1)

tmp12180 := PrimTail(V3165)

tmp12181 := PrimTail(tmp12180)

tmp12182 := PrimHead(tmp12181)

tmp12183 := PrimHead(tmp12182)

tmp12184 := Call(__e, PrimFunc(symlength), tmp12183)


tmp12185 := PrimNumberAdd(MakeNumber(1), tmp12184)

tmp12186 := Call(__e, PrimFunc(symshen_4nvars), tmp12185)


__e.TailApply(tmp12151, tmp12186)
return


}, 1)

tmp12187 := Call(__e, PrimFunc(symshen_4extract_1vars), V3165)


__e.TailApply(tmp12150, tmp12187)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.rule->clause")))
return
}


}, 1)

tmp12226 := Call(__e, ns2_1set, symshen_4rule_1_6clause, tmp12149)


_ = tmp12226

tmp12227 := MakeNative(func(__e *ControlFlow) {
V3172 := __e.Get(1)
_ = V3172
V3173 := __e.Get(2)
_ = V3173
tmp12232 := PrimIsPair(V3173)

if True == tmp12232 {
tmp12228 := Call(__e, PrimFunc(symshen_4optimise_1typing), V3172)


tmp12229 := PrimHead(V3173)

tmp12230 := PrimCons(tmp12229, Nil)

__e.Return(PrimCons(tmp12228, tmp12230))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-consequent")))
return
}


}, 2)

tmp12233 := Call(__e, ns2_1set, symshen_4compile_1consequent, tmp12227)


_ = tmp12233

tmp12234 := MakeNative(func(__e *ControlFlow) {
V3174 := __e.Get(1)
_ = V3174
tmp12240 := PrimEqual(MakeNumber(0), V3174)

if True == tmp12240 {
__e.Return(Nil)
return
} else {
tmp12235 := Call(__e, PrimFunc(symprotect), symV)


tmp12236 := Call(__e, PrimFunc(symgensym), tmp12235)


tmp12237 := PrimNumberSubtract(V3174, MakeNumber(1))

tmp12238 := Call(__e, PrimFunc(symshen_4nvars), tmp12237)


__e.Return(PrimCons(tmp12236, tmp12238))
return


}


}, 1)

tmp12241 := Call(__e, ns2_1set, symshen_4nvars, tmp12234)


_ = tmp12241

tmp12242 := MakeNative(func(__e *ControlFlow) {
V3175 := __e.Get(1)
_ = V3175
tmp12278 := PrimIsPair(V3175)

var ifres12257 Obj

if True == tmp12278 {
tmp12276 := PrimTail(V3175)

tmp12277 := PrimIsPair(tmp12276)

var ifres12259 Obj

if True == tmp12277 {
tmp12273 := PrimTail(V3175)

tmp12274 := PrimTail(tmp12273)

tmp12275 := PrimIsPair(tmp12274)

var ifres12261 Obj

if True == tmp12275 {
tmp12269 := PrimTail(V3175)

tmp12270 := PrimTail(tmp12269)

tmp12271 := PrimTail(tmp12270)

tmp12272 := PrimEqual(Nil, tmp12271)

var ifres12263 Obj

if True == tmp12272 {
tmp12265 := PrimTail(V3175)

tmp12266 := PrimHead(tmp12265)

tmp12267 := PrimIntern(MakeString(":"))

tmp12268 := PrimEqual(tmp12266, tmp12267)

var ifres12264 Obj

if True == tmp12268 {
ifres12264 = True


} else {
ifres12264 = False


}

ifres12263 = ifres12264


} else {
ifres12263 = False


}

var ifres12262 Obj

if True == ifres12263 {
ifres12262 = True


} else {
ifres12262 = False


}

ifres12261 = ifres12262


} else {
ifres12261 = False


}

var ifres12260 Obj

if True == ifres12261 {
ifres12260 = True


} else {
ifres12260 = False


}

ifres12259 = ifres12260


} else {
ifres12259 = False


}

var ifres12258 Obj

if True == ifres12259 {
ifres12258 = True


} else {
ifres12258 = False


}

ifres12257 = ifres12258


} else {
ifres12257 = False


}

if True == ifres12257 {
tmp12243 := PrimHead(V3175)

tmp12244 := PrimTail(V3175)

tmp12245 := PrimHead(tmp12244)

tmp12246 := PrimTail(V3175)

tmp12247 := PrimTail(tmp12246)

tmp12248 := PrimCons(sym_7, tmp12247)

tmp12249 := PrimCons(tmp12248, Nil)

tmp12250 := PrimCons(tmp12245, tmp12249)

tmp12251 := PrimCons(tmp12243, tmp12250)

tmp12252 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp12251)


tmp12253 := PrimCons(tmp12252, Nil)

__e.Return(PrimCons(sym_1, tmp12253))
return


} else {
tmp12254 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), V3175)


tmp12255 := PrimCons(tmp12254, Nil)

__e.Return(PrimCons(sym_7, tmp12255))
return


}


}, 1)

tmp12279 := Call(__e, ns2_1set, symshen_4optimise_1typing, tmp12242)


_ = tmp12279

tmp12280 := MakeNative(func(__e *ControlFlow) {
V3176 := __e.Get(1)
_ = V3176
tmp12370 := PrimIsPair(V3176)

var ifres12357 Obj

if True == tmp12370 {
tmp12368 := PrimHead(V3176)

tmp12369 := PrimEqual(sym_1, tmp12368)

var ifres12359 Obj

if True == tmp12369 {
tmp12366 := PrimTail(V3176)

tmp12367 := PrimIsPair(tmp12366)

var ifres12361 Obj

if True == tmp12367 {
tmp12363 := PrimTail(V3176)

tmp12364 := PrimTail(tmp12363)

tmp12365 := PrimEqual(Nil, tmp12364)

var ifres12362 Obj

if True == tmp12365 {
ifres12362 = True


} else {
ifres12362 = False


}

ifres12361 = ifres12362


} else {
ifres12361 = False


}

var ifres12360 Obj

if True == ifres12361 {
ifres12360 = True


} else {
ifres12360 = False


}

ifres12359 = ifres12360


} else {
ifres12359 = False


}

var ifres12358 Obj

if True == ifres12359 {
ifres12358 = True


} else {
ifres12358 = False


}

ifres12357 = ifres12358


} else {
ifres12357 = False


}

if True == ifres12357 {
tmp12281 := PrimTail(V3176)

tmp12282 := PrimHead(tmp12281)

tmp12283 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp12282)


tmp12284 := PrimCons(tmp12283, Nil)

__e.Return(PrimCons(sym_1, tmp12284))
return


} else {
tmp12355 := PrimIsPair(V3176)

var ifres12342 Obj

if True == tmp12355 {
tmp12353 := PrimHead(V3176)

tmp12354 := PrimEqual(sym_7, tmp12353)

var ifres12344 Obj

if True == tmp12354 {
tmp12351 := PrimTail(V3176)

tmp12352 := PrimIsPair(tmp12351)

var ifres12346 Obj

if True == tmp12352 {
tmp12348 := PrimTail(V3176)

tmp12349 := PrimTail(tmp12348)

tmp12350 := PrimEqual(Nil, tmp12349)

var ifres12347 Obj

if True == tmp12350 {
ifres12347 = True


} else {
ifres12347 = False


}

ifres12346 = ifres12347


} else {
ifres12346 = False


}

var ifres12345 Obj

if True == ifres12346 {
ifres12345 = True


} else {
ifres12345 = False


}

ifres12344 = ifres12345


} else {
ifres12344 = False


}

var ifres12343 Obj

if True == ifres12344 {
ifres12343 = True


} else {
ifres12343 = False


}

ifres12342 = ifres12343


} else {
ifres12342 = False


}

if True == ifres12342 {
tmp12285 := PrimTail(V3176)

tmp12286 := PrimHead(tmp12285)

tmp12287 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp12286)


tmp12288 := PrimCons(tmp12287, Nil)

__e.Return(PrimCons(sym_7, tmp12288))
return


} else {
tmp12340 := PrimIsPair(V3176)

var ifres12321 Obj

if True == tmp12340 {
tmp12338 := PrimHead(V3176)

tmp12339 := PrimEqual(symmode, tmp12338)

var ifres12323 Obj

if True == tmp12339 {
tmp12336 := PrimTail(V3176)

tmp12337 := PrimIsPair(tmp12336)

var ifres12325 Obj

if True == tmp12337 {
tmp12333 := PrimTail(V3176)

tmp12334 := PrimTail(tmp12333)

tmp12335 := PrimIsPair(tmp12334)

var ifres12327 Obj

if True == tmp12335 {
tmp12329 := PrimTail(V3176)

tmp12330 := PrimTail(tmp12329)

tmp12331 := PrimTail(tmp12330)

tmp12332 := PrimEqual(Nil, tmp12331)

var ifres12328 Obj

if True == tmp12332 {
ifres12328 = True


} else {
ifres12328 = False


}

ifres12327 = ifres12328


} else {
ifres12327 = False


}

var ifres12326 Obj

if True == ifres12327 {
ifres12326 = True


} else {
ifres12326 = False


}

ifres12325 = ifres12326


} else {
ifres12325 = False


}

var ifres12324 Obj

if True == ifres12325 {
ifres12324 = True


} else {
ifres12324 = False


}

ifres12323 = ifres12324


} else {
ifres12323 = False


}

var ifres12322 Obj

if True == ifres12323 {
ifres12322 = True


} else {
ifres12322 = False


}

ifres12321 = ifres12322


} else {
ifres12321 = False


}

if True == ifres12321 {
tmp12289 := PrimTail(V3176)

tmp12290 := PrimTail(tmp12289)

tmp12291 := PrimHead(tmp12290)

tmp12292 := PrimTail(V3176)

tmp12293 := PrimHead(tmp12292)

tmp12294 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp12293)


tmp12295 := PrimCons(tmp12294, Nil)

__e.Return(PrimCons(tmp12291, tmp12295))
return


} else {
tmp12319 := PrimIsPair(V3176)

var ifres12306 Obj

if True == tmp12319 {
tmp12317 := PrimHead(V3176)

tmp12318 := PrimEqual(symbar_b, tmp12317)

var ifres12308 Obj

if True == tmp12318 {
tmp12315 := PrimTail(V3176)

tmp12316 := PrimIsPair(tmp12315)

var ifres12310 Obj

if True == tmp12316 {
tmp12312 := PrimTail(V3176)

tmp12313 := PrimTail(tmp12312)

tmp12314 := PrimEqual(Nil, tmp12313)

var ifres12311 Obj

if True == tmp12314 {
ifres12311 = True


} else {
ifres12311 = False


}

ifres12310 = ifres12311


} else {
ifres12310 = False


}

var ifres12309 Obj

if True == ifres12310 {
ifres12309 = True


} else {
ifres12309 = False


}

ifres12308 = ifres12309


} else {
ifres12308 = False


}

var ifres12307 Obj

if True == ifres12308 {
ifres12307 = True


} else {
ifres12307 = False


}

ifres12306 = ifres12307


} else {
ifres12306 = False


}

if True == ifres12306 {
tmp12296 := PrimTail(V3176)

__e.Return(PrimHead(tmp12296))
return


} else {
tmp12304 := PrimIsPair(V3176)

if True == tmp12304 {
tmp12297 := PrimHead(V3176)

tmp12298 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp12297)


tmp12299 := PrimTail(V3176)

tmp12300 := Call(__e, PrimFunc(symshen_4cons_1form_1with_1modes), tmp12299)


tmp12301 := PrimCons(tmp12300, Nil)

tmp12302 := PrimCons(tmp12298, tmp12301)

__e.Return(PrimCons(symcons, tmp12302))
return


} else {
__e.Return(V3176)
return
}


}


}


}


}


}, 1)

tmp12371 := Call(__e, ns2_1set, symshen_4cons_1form_1with_1modes, tmp12280)


_ = tmp12371

tmp12372 := MakeNative(func(__e *ControlFlow) {
V3177 := __e.Get(1)
_ = V3177
V3178 := __e.Get(2)
_ = V3178
V3179 := __e.Get(3)
_ = V3179
V3180 := __e.Get(4)
_ = V3180
V3181 := __e.Get(5)
_ = V3181
V3182 := __e.Get(6)
_ = V3182
tmp12373 := MakeNative(func(__e *ControlFlow) {
GoalsAs := __e.Get(1)
_ = GoalsAs
tmp12374 := MakeNative(func(__e *ControlFlow) {
GoalsS := __e.Get(1)
_ = GoalsS
tmp12375 := MakeNative(func(__e *ControlFlow) {
GoalsP := __e.Get(1)
_ = GoalsP
tmp12376 := Call(__e, PrimFunc(symappend), GoalsS, GoalsP)


__e.TailApply(PrimFunc(symappend), GoalsAs, tmp12376)
return


}, 1)

tmp12377 := Call(__e, PrimFunc(symshen_4compile_1premises), V3180, V3181)


__e.TailApply(tmp12375, tmp12377)
return


}, 1)

tmp12378 := Call(__e, PrimFunc(symshen_4compile_1side_1conditions), V3179)


__e.TailApply(tmp12374, tmp12378)
return


}, 1)

tmp12379 := Call(__e, PrimFunc(symshen_4compile_1assumptions), V3178, V3177, V3181, V3182)


__e.TailApply(tmp12373, tmp12379)
return


}, 6)

tmp12380 := Call(__e, ns2_1set, symshen_4goals, tmp12372)


_ = tmp12380

tmp12381 := MakeNative(func(__e *ControlFlow) {
V3197 := __e.Get(1)
_ = V3197
V3198 := __e.Get(2)
_ = V3198
V3199 := __e.Get(3)
_ = V3199
V3200 := __e.Get(4)
_ = V3200
tmp12404 := PrimEqual(Nil, V3197)

if True == tmp12404 {
__e.Return(Nil)
return
} else {
tmp12402 := PrimIsPair(V3197)

var ifres12395 Obj

if True == tmp12402 {
tmp12401 := PrimIsPair(V3199)

var ifres12397 Obj

if True == tmp12401 {
tmp12399 := PrimTail(V3199)

tmp12400 := PrimIsPair(tmp12399)

var ifres12398 Obj

if True == tmp12400 {
ifres12398 = True


} else {
ifres12398 = False


}

ifres12397 = ifres12398


} else {
ifres12397 = False


}

var ifres12396 Obj

if True == ifres12397 {
ifres12396 = True


} else {
ifres12396 = False


}

ifres12395 = ifres12396


} else {
ifres12395 = False


}

if True == ifres12395 {
tmp12382 := MakeNative(func(__e *ControlFlow) {
NewActive := __e.Get(1)
_ = NewActive
tmp12383 := PrimHead(V3197)

tmp12384 := PrimHead(V3199)

tmp12385 := PrimTail(V3199)

tmp12386 := PrimHead(tmp12385)

tmp12387 := Call(__e, PrimFunc(symshen_4compile_1assumption), tmp12383, tmp12384, tmp12386, V3198, V3200)


tmp12388 := PrimTail(V3197)

tmp12389 := PrimTail(V3199)

tmp12390 := Call(__e, PrimFunc(symshen_4compile_1assumptions), tmp12388, V3198, tmp12389, NewActive)


__e.Return(PrimCons(tmp12387, tmp12390))
return


}, 1)

tmp12391 := PrimHead(V3197)

tmp12392 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp12391)


tmp12393 := Call(__e, PrimFunc(symappend), tmp12392, V3200)


__e.TailApply(tmp12382, tmp12393)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-assumptions")))
return
}


}


}, 4)

tmp12405 := Call(__e, ns2_1set, symshen_4compile_1assumptions, tmp12381)


_ = tmp12405

tmp12406 := MakeNative(func(__e *ControlFlow) {
V3201 := __e.Get(1)
_ = V3201
V3202 := __e.Get(2)
_ = V3202
V3203 := __e.Get(3)
_ = V3203
V3204 := __e.Get(4)
_ = V3204
V3205 := __e.Get(5)
_ = V3205
tmp12407 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp12408 := MakeNative(func(__e *ControlFlow) {
Compile := __e.Get(1)
_ = Compile
tmp12409 := PrimCons(V3203, V3204)

tmp12410 := PrimCons(Nil, tmp12409)

tmp12411 := PrimCons(V3202, tmp12410)

__e.Return(PrimCons(F, tmp12411))
return


}, 1)

tmp12412 := Call(__e, PrimFunc(symshen_4compile_1search_1procedure), F, V3201, V3202, V3203, V3204, V3205)


__e.TailApply(tmp12408, tmp12412)
return


}, 1)

tmp12413 := Call(__e, PrimFunc(symgensym), symshen_4search)


__e.TailApply(tmp12407, tmp12413)
return


}, 5)

tmp12414 := Call(__e, ns2_1set, symshen_4compile_1assumption, tmp12406)


_ = tmp12414

tmp12415 := MakeNative(func(__e *ControlFlow) {
V3206 := __e.Get(1)
_ = V3206
V3207 := __e.Get(2)
_ = V3207
V3208 := __e.Get(3)
_ = V3208
V3209 := __e.Get(4)
_ = V3209
V3210 := __e.Get(5)
_ = V3210
V3211 := __e.Get(6)
_ = V3211
tmp12416 := MakeNative(func(__e *ControlFlow) {
Past := __e.Get(1)
_ = Past
tmp12417 := MakeNative(func(__e *ControlFlow) {
Base := __e.Get(1)
_ = Base
tmp12418 := MakeNative(func(__e *ControlFlow) {
Recursive := __e.Get(1)
_ = Recursive
tmp12419 := Call(__e, PrimFunc(symappend), Base, Recursive)


tmp12420 := PrimCons(V3206, tmp12419)

tmp12421 := PrimCons(symdefprolog, tmp12420)

__e.TailApply(PrimFunc(symeval), tmp12421)
return


}, 1)

tmp12422 := Call(__e, PrimFunc(symshen_4keep_1looking), V3206, V3208, Past, V3209, V3210)


__e.TailApply(tmp12418, tmp12422)
return


}, 1)

tmp12423 := Call(__e, PrimFunc(symshen_4foundit_b), V3207, V3208, Past, V3209, V3210, V3211)


__e.TailApply(tmp12417, tmp12423)
return


}, 1)

tmp12424 := Call(__e, PrimFunc(symprotect), symPrevious)


tmp12425 := Call(__e, PrimFunc(symgensym), tmp12424)


__e.TailApply(tmp12416, tmp12425)
return


}, 6)

tmp12426 := Call(__e, ns2_1set, symshen_4compile_1search_1procedure, tmp12415)


_ = tmp12426

tmp12427 := MakeNative(func(__e *ControlFlow) {
V3212 := __e.Get(1)
_ = V3212
V3213 := __e.Get(2)
_ = V3213
V3214 := __e.Get(3)
_ = V3214
V3215 := __e.Get(4)
_ = V3215
V3216 := __e.Get(5)
_ = V3216
V3217 := __e.Get(6)
_ = V3217
tmp12428 := MakeNative(func(__e *ControlFlow) {
Passive := __e.Get(1)
_ = Passive
tmp12429 := MakeNative(func(__e *ControlFlow) {
Table := __e.Get(1)
_ = Table
tmp12430 := MakeNative(func(__e *ControlFlow) {
Head := __e.Get(1)
_ = Head
tmp12431 := MakeNative(func(__e *ControlFlow) {
Body := __e.Get(1)
_ = Body
tmp12432 := PrimCons(sym_5_1_1, Nil)

tmp12433 := PrimIntern(MakeString(";"))

tmp12434 := PrimCons(tmp12433, Nil)

tmp12435 := Call(__e, PrimFunc(symappend), Body, tmp12434)


tmp12436 := Call(__e, PrimFunc(symappend), tmp12432, tmp12435)


__e.TailApply(PrimFunc(symappend), Head, tmp12436)
return


}, 1)

tmp12437 := Call(__e, PrimFunc(symshen_4body_1foundit_b), V3213, V3214, V3215, Table)


__e.TailApply(tmp12431, tmp12437)
return


}, 1)

tmp12438 := Call(__e, PrimFunc(symshen_4head_1foundit_b), V3212, V3213, V3214, V3215, V3216, Table)


__e.TailApply(tmp12430, tmp12438)
return


}, 1)

tmp12439 := Call(__e, PrimFunc(symshen_4tabulate_1passive), Passive)


__e.TailApply(tmp12429, tmp12439)
return


}, 1)

tmp12440 := Call(__e, PrimFunc(symshen_4passive), V3212, V3217)


__e.TailApply(tmp12428, tmp12440)
return


}, 6)

tmp12441 := Call(__e, ns2_1set, symshen_4foundit_b, tmp12427)


_ = tmp12441

tmp12442 := MakeNative(func(__e *ControlFlow) {
V3218 := __e.Get(1)
_ = V3218
V3219 := __e.Get(2)
_ = V3219
V3220 := __e.Get(3)
_ = V3220
V3221 := __e.Get(4)
_ = V3221
V3222 := __e.Get(5)
_ = V3222
tmp12443 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp12444 := MakeNative(func(__e *ControlFlow) {
Head := __e.Get(1)
_ = Head
tmp12445 := MakeNative(func(__e *ControlFlow) {
Body := __e.Get(1)
_ = Body
tmp12446 := PrimCons(sym_5_1_1, Nil)

tmp12447 := PrimIntern(MakeString(";"))

tmp12448 := PrimCons(tmp12447, Nil)

tmp12449 := Call(__e, PrimFunc(symappend), Body, tmp12448)


tmp12450 := Call(__e, PrimFunc(symappend), tmp12446, tmp12449)


__e.TailApply(PrimFunc(symappend), Head, tmp12450)
return


}, 1)

tmp12451 := PrimCons(V3220, Nil)

tmp12452 := PrimCons(X, tmp12451)

tmp12453 := PrimCons(symcons, tmp12452)

tmp12454 := PrimCons(V3221, V3222)

tmp12455 := PrimCons(tmp12453, tmp12454)

tmp12456 := PrimCons(V3219, tmp12455)

tmp12457 := PrimCons(V3218, tmp12456)

tmp12458 := PrimCons(tmp12457, Nil)

__e.TailApply(tmp12445, tmp12458)
return


}, 1)

tmp12459 := PrimCons(V3219, Nil)

tmp12460 := PrimCons(X, tmp12459)

tmp12461 := PrimCons(symcons, tmp12460)

tmp12462 := PrimCons(tmp12461, Nil)

tmp12463 := PrimCons(sym_1, tmp12462)

tmp12464 := PrimCons(V3221, V3222)

tmp12465 := PrimCons(V3220, tmp12464)

tmp12466 := PrimCons(tmp12463, tmp12465)

__e.TailApply(tmp12444, tmp12466)
return


}, 1)

tmp12467 := Call(__e, PrimFunc(symprotect), symV)


tmp12468 := Call(__e, PrimFunc(symgensym), tmp12467)


__e.TailApply(tmp12443, tmp12468)
return


}, 5)

tmp12469 := Call(__e, ns2_1set, symshen_4keep_1looking, tmp12442)


_ = tmp12469

tmp12470 := MakeNative(func(__e *ControlFlow) {
V3227 := __e.Get(1)
_ = V3227
V3228 := __e.Get(2)
_ = V3228
tmp12478 := PrimIsPair(V3227)

if True == tmp12478 {
tmp12471 := PrimHead(V3227)

tmp12472 := Call(__e, PrimFunc(symshen_4passive), tmp12471, V3228)


tmp12473 := PrimTail(V3227)

tmp12474 := Call(__e, PrimFunc(symshen_4passive), tmp12473, V3228)


__e.TailApply(PrimFunc(symunion), tmp12472, tmp12474)
return


} else {
tmp12476 := Call(__e, PrimFunc(symshen_4passive_2), V3227, V3228)


if True == tmp12476 {
__e.Return(PrimCons(V3227, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}, 2)

tmp12479 := Call(__e, ns2_1set, symshen_4passive, tmp12470)


_ = tmp12479

tmp12480 := MakeNative(func(__e *ControlFlow) {
V3229 := __e.Get(1)
_ = V3229
V3230 := __e.Get(2)
_ = V3230
tmp12484 := Call(__e, PrimFunc(symelement_2), V3229, V3230)


tmp12485 := PrimNot(tmp12484)

if True == tmp12485 {
tmp12482 := PrimIsVariable(V3229)

if True == tmp12482 {
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


}, 2)

tmp12486 := Call(__e, ns2_1set, symshen_4passive_2, tmp12480)


_ = tmp12486

tmp12487 := MakeNative(func(__e *ControlFlow) {
V3231 := __e.Get(1)
_ = V3231
tmp12488 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp12489 := Call(__e, PrimFunc(symprotect), symV)


tmp12490 := Call(__e, PrimFunc(symgensym), tmp12489)


__e.Return(PrimCons(X, tmp12490))
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp12488, V3231)
return


}, 1)

tmp12491 := Call(__e, ns2_1set, symshen_4tabulate_1passive, tmp12487)


_ = tmp12491

tmp12492 := MakeNative(func(__e *ControlFlow) {
V3232 := __e.Get(1)
_ = V3232
V3233 := __e.Get(2)
_ = V3233
V3234 := __e.Get(3)
_ = V3234
V3235 := __e.Get(4)
_ = V3235
V3236 := __e.Get(5)
_ = V3236
V3237 := __e.Get(6)
_ = V3237
tmp12493 := MakeNative(func(__e *ControlFlow) {
Optimise := __e.Get(1)
_ = Optimise
tmp12494 := Call(__e, PrimFunc(symshen_4optimise_1typing), V3232)


tmp12495 := PrimCons(V3233, Nil)

tmp12496 := PrimCons(tmp12494, tmp12495)

tmp12497 := PrimCons(symcons, tmp12496)

tmp12498 := PrimCons(tmp12497, Nil)

tmp12499 := PrimCons(sym_1, tmp12498)

tmp12500 := PrimCons(V3235, Optimise)

tmp12501 := PrimCons(V3234, tmp12500)

__e.Return(PrimCons(tmp12499, tmp12501))
return


}, 1)

tmp12502 := Call(__e, PrimFunc(symshen_4optimise_1passive), V3236, V3237)


__e.TailApply(tmp12493, tmp12502)
return


}, 6)

tmp12503 := Call(__e, ns2_1set, symshen_4head_1foundit_b, tmp12492)


_ = tmp12503

tmp12504 := MakeNative(func(__e *ControlFlow) {
V3238 := __e.Get(1)
_ = V3238
V3239 := __e.Get(2)
_ = V3239
tmp12505 := MakeNative(func(__e *ControlFlow) {
C := __e.Get(1)
_ = C
__e.TailApply(PrimFunc(symshen_4optimise_1passive_1h), C, V3239)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12505, V3238)
return


}, 2)

tmp12506 := Call(__e, ns2_1set, symshen_4optimise_1passive, tmp12504)


_ = tmp12506

tmp12507 := MakeNative(func(__e *ControlFlow) {
V3240 := __e.Get(1)
_ = V3240
V3241 := __e.Get(2)
_ = V3241
tmp12508 := MakeNative(func(__e *ControlFlow) {
Entry := __e.Get(1)
_ = Entry
tmp12510 := Call(__e, PrimFunc(symempty_2), Entry)


if True == tmp12510 {
__e.Return(V3240)
return
} else {
__e.Return(PrimTail(Entry))
return
}


}, 1)

tmp12511 := Call(__e, PrimFunc(symassoc), V3240, V3241)


__e.TailApply(tmp12508, tmp12511)
return


}, 2)

tmp12512 := Call(__e, ns2_1set, symshen_4optimise_1passive_1h, tmp12507)


_ = tmp12512

tmp12513 := MakeNative(func(__e *ControlFlow) {
V3250 := __e.Get(1)
_ = V3250
V3251 := __e.Get(2)
_ = V3251
V3252 := __e.Get(3)
_ = V3252
V3253 := __e.Get(4)
_ = V3253
tmp12540 := PrimEqual(Nil, V3253)

if True == tmp12540 {
tmp12514 := PrimCons(V3251, Nil)

tmp12515 := PrimCons(MakeNumber(1), tmp12514)

tmp12516 := PrimCons(V3250, Nil)

tmp12517 := PrimCons(MakeNumber(1), tmp12516)

tmp12518 := PrimCons(tmp12517, Nil)

tmp12519 := PrimCons(tmp12515, tmp12518)

tmp12520 := PrimCons(symappend, tmp12519)

tmp12521 := PrimCons(tmp12520, Nil)

tmp12522 := PrimCons(V3252, tmp12521)

tmp12523 := PrimCons(symbind, tmp12522)

__e.Return(PrimCons(tmp12523, Nil))
return


} else {
tmp12538 := PrimIsPair(V3253)

var ifres12534 Obj

if True == tmp12538 {
tmp12536 := PrimHead(V3253)

tmp12537 := PrimIsPair(tmp12536)

var ifres12535 Obj

if True == tmp12537 {
ifres12535 = True


} else {
ifres12535 = False


}

ifres12534 = ifres12535


} else {
ifres12534 = False


}

if True == ifres12534 {
tmp12524 := PrimHead(V3253)

tmp12525 := PrimTail(tmp12524)

tmp12526 := PrimHead(V3253)

tmp12527 := PrimHead(tmp12526)

tmp12528 := PrimCons(tmp12527, Nil)

tmp12529 := PrimCons(tmp12525, tmp12528)

tmp12530 := PrimCons(symbind, tmp12529)

tmp12531 := PrimTail(V3253)

tmp12532 := Call(__e, PrimFunc(symshen_4body_1foundit_b), V3250, V3251, V3252, tmp12531)


__e.Return(PrimCons(tmp12530, tmp12532))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.body-foundit!")))
return
}


}


}, 4)

tmp12541 := Call(__e, ns2_1set, symshen_4body_1foundit_b, tmp12513)


_ = tmp12541

tmp12542 := MakeNative(func(__e *ControlFlow) {
V3254 := __e.Get(1)
_ = V3254
tmp12543 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4compile_1side_1condition), X)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12543, V3254)
return


}, 1)

tmp12544 := Call(__e, ns2_1set, symshen_4compile_1side_1conditions, tmp12542)


_ = tmp12544

tmp12545 := MakeNative(func(__e *ControlFlow) {
V3257 := __e.Get(1)
_ = V3257
tmp12605 := PrimIsPair(V3257)

var ifres12586 Obj

if True == tmp12605 {
tmp12603 := PrimHead(V3257)

tmp12604 := PrimEqual(symlet, tmp12603)

var ifres12588 Obj

if True == tmp12604 {
tmp12601 := PrimTail(V3257)

tmp12602 := PrimIsPair(tmp12601)

var ifres12590 Obj

if True == tmp12602 {
tmp12598 := PrimTail(V3257)

tmp12599 := PrimTail(tmp12598)

tmp12600 := PrimIsPair(tmp12599)

var ifres12592 Obj

if True == tmp12600 {
tmp12594 := PrimTail(V3257)

tmp12595 := PrimTail(tmp12594)

tmp12596 := PrimTail(tmp12595)

tmp12597 := PrimEqual(Nil, tmp12596)

var ifres12593 Obj

if True == tmp12597 {
ifres12593 = True


} else {
ifres12593 = False


}

ifres12592 = ifres12593


} else {
ifres12592 = False


}

var ifres12591 Obj

if True == ifres12592 {
ifres12591 = True


} else {
ifres12591 = False


}

ifres12590 = ifres12591


} else {
ifres12590 = False


}

var ifres12589 Obj

if True == ifres12590 {
ifres12589 = True


} else {
ifres12589 = False


}

ifres12588 = ifres12589


} else {
ifres12588 = False


}

var ifres12587 Obj

if True == ifres12588 {
ifres12587 = True


} else {
ifres12587 = False


}

ifres12586 = ifres12587


} else {
ifres12586 = False


}

if True == ifres12586 {
tmp12546 := PrimTail(V3257)

__e.Return(PrimCons(symis, tmp12546))
return


} else {
tmp12584 := PrimIsPair(V3257)

var ifres12565 Obj

if True == tmp12584 {
tmp12582 := PrimHead(V3257)

tmp12583 := PrimEqual(symshen_4let_b, tmp12582)

var ifres12567 Obj

if True == tmp12583 {
tmp12580 := PrimTail(V3257)

tmp12581 := PrimIsPair(tmp12580)

var ifres12569 Obj

if True == tmp12581 {
tmp12577 := PrimTail(V3257)

tmp12578 := PrimTail(tmp12577)

tmp12579 := PrimIsPair(tmp12578)

var ifres12571 Obj

if True == tmp12579 {
tmp12573 := PrimTail(V3257)

tmp12574 := PrimTail(tmp12573)

tmp12575 := PrimTail(tmp12574)

tmp12576 := PrimEqual(Nil, tmp12575)

var ifres12572 Obj

if True == tmp12576 {
ifres12572 = True


} else {
ifres12572 = False


}

ifres12571 = ifres12572


} else {
ifres12571 = False


}

var ifres12570 Obj

if True == ifres12571 {
ifres12570 = True


} else {
ifres12570 = False


}

ifres12569 = ifres12570


} else {
ifres12569 = False


}

var ifres12568 Obj

if True == ifres12569 {
ifres12568 = True


} else {
ifres12568 = False


}

ifres12567 = ifres12568


} else {
ifres12567 = False


}

var ifres12566 Obj

if True == ifres12567 {
ifres12566 = True


} else {
ifres12566 = False


}

ifres12565 = ifres12566


} else {
ifres12565 = False


}

if True == ifres12565 {
tmp12547 := PrimTail(V3257)

__e.Return(PrimCons(symis_b, tmp12547))
return


} else {
tmp12563 := PrimIsPair(V3257)

var ifres12550 Obj

if True == tmp12563 {
tmp12561 := PrimHead(V3257)

tmp12562 := PrimEqual(symif, tmp12561)

var ifres12552 Obj

if True == tmp12562 {
tmp12559 := PrimTail(V3257)

tmp12560 := PrimIsPair(tmp12559)

var ifres12554 Obj

if True == tmp12560 {
tmp12556 := PrimTail(V3257)

tmp12557 := PrimTail(tmp12556)

tmp12558 := PrimEqual(Nil, tmp12557)

var ifres12555 Obj

if True == tmp12558 {
ifres12555 = True


} else {
ifres12555 = False


}

ifres12554 = ifres12555


} else {
ifres12554 = False


}

var ifres12553 Obj

if True == ifres12554 {
ifres12553 = True


} else {
ifres12553 = False


}

ifres12552 = ifres12553


} else {
ifres12552 = False


}

var ifres12551 Obj

if True == ifres12552 {
ifres12551 = True


} else {
ifres12551 = False


}

ifres12550 = ifres12551


} else {
ifres12550 = False


}

if True == ifres12550 {
tmp12548 := PrimTail(V3257)

__e.Return(PrimCons(symwhen, tmp12548))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-side-condition")))
return
}


}


}


}, 1)

tmp12606 := Call(__e, ns2_1set, symshen_4compile_1side_1condition, tmp12545)


_ = tmp12606

tmp12607 := MakeNative(func(__e *ControlFlow) {
V3258 := __e.Get(1)
_ = V3258
V3259 := __e.Get(2)
_ = V3259
tmp12608 := MakeNative(func(__e *ControlFlow) {
Hyp := __e.Get(1)
_ = Hyp
tmp12609 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4compile_1premise), X, Hyp)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12609, V3258)
return


}, 1)

tmp12610 := Call(__e, PrimFunc(symreverse), V3259)


tmp12611 := PrimHead(tmp12610)

__e.TailApply(tmp12608, tmp12611)
return


}, 2)

tmp12612 := Call(__e, ns2_1set, symshen_4compile_1premises, tmp12607)


_ = tmp12612

tmp12613 := MakeNative(func(__e *ControlFlow) {
V3266 := __e.Get(1)
_ = V3266
V3267 := __e.Get(2)
_ = V3267
tmp12630 := PrimEqual(sym_b, V3266)

if True == tmp12630 {
__e.Return(sym_b)
return
} else {
tmp12628 := PrimIsPair(V3266)

var ifres12619 Obj

if True == tmp12628 {
tmp12626 := PrimTail(V3266)

tmp12627 := PrimIsPair(tmp12626)

var ifres12621 Obj

if True == tmp12627 {
tmp12623 := PrimTail(V3266)

tmp12624 := PrimTail(tmp12623)

tmp12625 := PrimEqual(Nil, tmp12624)

var ifres12622 Obj

if True == tmp12625 {
ifres12622 = True


} else {
ifres12622 = False


}

ifres12621 = ifres12622


} else {
ifres12621 = False


}

var ifres12620 Obj

if True == ifres12621 {
ifres12620 = True


} else {
ifres12620 = False


}

ifres12619 = ifres12620


} else {
ifres12619 = False


}

if True == ifres12619 {
tmp12614 := PrimHead(V3266)

tmp12615 := Call(__e, PrimFunc(symreverse), tmp12614)


tmp12616 := PrimTail(V3266)

tmp12617 := PrimHead(tmp12616)

__e.TailApply(PrimFunc(symshen_4compile_1premise_1h), tmp12615, tmp12617, V3267)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.premise")))
return
}


}


}, 2)

tmp12631 := Call(__e, ns2_1set, symshen_4compile_1premise, tmp12613)


_ = tmp12631

tmp12632 := MakeNative(func(__e *ControlFlow) {
V3274 := __e.Get(1)
_ = V3274
V3275 := __e.Get(2)
_ = V3275
V3276 := __e.Get(3)
_ = V3276
tmp12645 := PrimEqual(Nil, V3274)

if True == tmp12645 {
tmp12633 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), V3275)


tmp12634 := PrimCons(V3276, Nil)

tmp12635 := PrimCons(tmp12633, tmp12634)

__e.Return(PrimCons(symshen_4system_1S, tmp12635))
return


} else {
tmp12643 := PrimIsPair(V3274)

if True == tmp12643 {
tmp12636 := PrimTail(V3274)

tmp12637 := PrimHead(V3274)

tmp12638 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp12637)


tmp12639 := PrimCons(V3276, Nil)

tmp12640 := PrimCons(tmp12638, tmp12639)

tmp12641 := PrimCons(symcons, tmp12640)

__e.TailApply(PrimFunc(symshen_4compile_1premise_1h), tmp12636, V3275, tmp12641)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.compile-premise-h")))
return
}


}


}, 3)

tmp12646 := Call(__e, ns2_1set, symshen_4compile_1premise_1h, tmp12632)


_ = tmp12646

tmp12647 := MakeNative(func(__e *ControlFlow) {
V3277 := __e.Get(1)
_ = V3277
tmp12671 := PrimIsPair(V3277)

var ifres12658 Obj

if True == tmp12671 {
tmp12669 := PrimHead(V3277)

tmp12670 := PrimEqual(symbar_b, tmp12669)

var ifres12660 Obj

if True == tmp12670 {
tmp12667 := PrimTail(V3277)

tmp12668 := PrimIsPair(tmp12667)

var ifres12662 Obj

if True == tmp12668 {
tmp12664 := PrimTail(V3277)

tmp12665 := PrimTail(tmp12664)

tmp12666 := PrimEqual(Nil, tmp12665)

var ifres12663 Obj

if True == tmp12666 {
ifres12663 = True


} else {
ifres12663 = False


}

ifres12662 = ifres12663


} else {
ifres12662 = False


}

var ifres12661 Obj

if True == ifres12662 {
ifres12661 = True


} else {
ifres12661 = False


}

ifres12660 = ifres12661


} else {
ifres12660 = False


}

var ifres12659 Obj

if True == ifres12660 {
ifres12659 = True


} else {
ifres12659 = False


}

ifres12658 = ifres12659


} else {
ifres12658 = False


}

if True == ifres12658 {
tmp12648 := PrimTail(V3277)

__e.Return(PrimHead(tmp12648))
return


} else {
tmp12656 := PrimIsPair(V3277)

if True == tmp12656 {
tmp12649 := PrimHead(V3277)

tmp12650 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp12649)


tmp12651 := PrimTail(V3277)

tmp12652 := Call(__e, PrimFunc(symshen_4cons_1form_1no_1modes), tmp12651)


tmp12653 := PrimCons(tmp12652, Nil)

tmp12654 := PrimCons(tmp12650, tmp12653)

__e.Return(PrimCons(symcons, tmp12654))
return


} else {
__e.Return(V3277)
return
}


}


}, 1)

tmp12672 := Call(__e, ns2_1set, symshen_4cons_1form_1no_1modes, tmp12647)


_ = tmp12672

tmp12673 := MakeNative(func(__e *ControlFlow) {
V3278 := __e.Get(1)
_ = V3278
tmp12674 := MakeNative(func(__e *ControlFlow) {
InternTypes := __e.Get(1)
_ = InternTypes
tmp12675 := MakeNative(func(__e *ControlFlow) {
Datatypes := __e.Get(1)
_ = Datatypes
tmp12676 := MakeNative(func(__e *ControlFlow) {
Remove := __e.Get(1)
_ = Remove
tmp12677 := MakeNative(func(__e *ControlFlow) {
NewDatatypes := __e.Get(1)
_ = NewDatatypes
__e.TailApply(PrimFunc(symshen_4show_1datatypes), NewDatatypes)
return
}, 1)

tmp12678 := PrimSet(symshen_4_ddatatypes_d, Remove)

__e.TailApply(tmp12677, tmp12678)
return


}, 1)

tmp12679 := Call(__e, PrimFunc(symshen_4remove_1datatypes), InternTypes, Datatypes)


__e.TailApply(tmp12676, tmp12679)
return


}, 1)

tmp12680 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp12675, tmp12680)
return


}, 1)

tmp12681 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4intern_1type), X)
return
}, 1)

tmp12682 := Call(__e, PrimFunc(symmap), tmp12681, V3278)


__e.TailApply(tmp12674, tmp12682)
return


}, 1)

tmp12683 := Call(__e, ns2_1set, sympreclude, tmp12673)


_ = tmp12683

tmp12684 := MakeNative(func(__e *ControlFlow) {
V3283 := __e.Get(1)
_ = V3283
V3284 := __e.Get(2)
_ = V3284
tmp12691 := PrimEqual(Nil, V3283)

if True == tmp12691 {
__e.Return(V3284)
return
} else {
tmp12689 := PrimIsPair(V3283)

if True == tmp12689 {
tmp12685 := PrimTail(V3283)

tmp12686 := PrimHead(V3283)

tmp12687 := Call(__e, PrimFunc(symshen_4unassoc), tmp12686, V3284)


__e.TailApply(PrimFunc(symshen_4remove_1datatypes), tmp12685, tmp12687)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-datatypes")))
return
}


}


}, 2)

tmp12692 := Call(__e, ns2_1set, symshen_4remove_1datatypes, tmp12684)


_ = tmp12692

tmp12693 := MakeNative(func(__e *ControlFlow) {
V3294 := __e.Get(1)
_ = V3294
V3295 := __e.Get(2)
_ = V3295
tmp12711 := PrimEqual(Nil, V3295)

if True == tmp12711 {
__e.Return(Nil)
return
} else {
tmp12709 := PrimIsPair(V3295)

var ifres12700 Obj

if True == tmp12709 {
tmp12707 := PrimHead(V3295)

tmp12708 := PrimIsPair(tmp12707)

var ifres12702 Obj

if True == tmp12708 {
tmp12704 := PrimHead(V3295)

tmp12705 := PrimHead(tmp12704)

tmp12706 := PrimEqual(V3294, tmp12705)

var ifres12703 Obj

if True == tmp12706 {
ifres12703 = True


} else {
ifres12703 = False


}

ifres12702 = ifres12703


} else {
ifres12702 = False


}

var ifres12701 Obj

if True == ifres12702 {
ifres12701 = True


} else {
ifres12701 = False


}

ifres12700 = ifres12701


} else {
ifres12700 = False


}

if True == ifres12700 {
__e.Return(PrimTail(V3295))
return
} else {
tmp12698 := PrimIsPair(V3295)

if True == tmp12698 {
tmp12694 := PrimHead(V3295)

tmp12695 := PrimTail(V3295)

tmp12696 := Call(__e, PrimFunc(symshen_4unassoc), V3294, tmp12695)


__e.Return(PrimCons(tmp12694, tmp12696))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.unassoc")))
return
}


}


}


}, 2)

tmp12712 := Call(__e, ns2_1set, symshen_4unassoc, tmp12693)


_ = tmp12712

tmp12713 := MakeNative(func(__e *ControlFlow) {
V3296 := __e.Get(1)
_ = V3296
tmp12714 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.Return(PrimHead(X))
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp12714, V3296)
return


}, 1)

tmp12715 := Call(__e, ns2_1set, symshen_4show_1datatypes, tmp12713)


_ = tmp12715

tmp12716 := MakeNative(func(__e *ControlFlow) {
V3297 := __e.Get(1)
_ = V3297
tmp12717 := MakeNative(func(__e *ControlFlow) {
InternTypes := __e.Get(1)
_ = InternTypes
tmp12718 := MakeNative(func(__e *ControlFlow) {
Remember := __e.Get(1)
_ = Remember
tmp12719 := MakeNative(func(__e *ControlFlow) {
Datatypes := __e.Get(1)
_ = Datatypes
__e.TailApply(PrimFunc(symshen_4show_1datatypes), Datatypes)
return
}, 1)

tmp12720 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(tmp12719, tmp12720)
return


}, 1)

tmp12721 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp12722 := Call(__e, PrimFunc(symfn), D)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), D, tmp12722)
return


}, 1)

tmp12723 := Call(__e, PrimFunc(symmap), tmp12721, InternTypes)


__e.TailApply(tmp12718, tmp12723)
return


}, 1)

tmp12724 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4intern_1type), X)
return
}, 1)

tmp12725 := Call(__e, PrimFunc(symmap), tmp12724, V3297)


__e.TailApply(tmp12717, tmp12725)
return


}, 1)

tmp12726 := Call(__e, ns2_1set, syminclude, tmp12716)


_ = tmp12726

tmp12727 := MakeNative(func(__e *ControlFlow) {
V3298 := __e.Get(1)
_ = V3298
tmp12728 := MakeNative(func(__e *ControlFlow) {
Initialise := __e.Get(1)
_ = Initialise
tmp12729 := MakeNative(func(__e *ControlFlow) {
InternTypes := __e.Get(1)
_ = InternTypes
tmp12730 := MakeNative(func(__e *ControlFlow) {
NewDatatypes := __e.Get(1)
_ = NewDatatypes
tmp12731 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symshen_4show_1datatypes), tmp12731)
return


}, 1)

tmp12732 := MakeNative(func(__e *ControlFlow) {
D := __e.Get(1)
_ = D
tmp12733 := Call(__e, PrimFunc(symfn), D)


__e.TailApply(PrimFunc(symshen_4remember_1datatype), D, tmp12733)
return


}, 1)

tmp12734 := Call(__e, PrimFunc(symmap), tmp12732, InternTypes)


__e.TailApply(tmp12730, tmp12734)
return


}, 1)

tmp12735 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4intern_1type), X)
return
}, 1)

tmp12736 := Call(__e, PrimFunc(symmap), tmp12735, V3298)


__e.TailApply(tmp12729, tmp12736)
return


}, 1)

tmp12737 := PrimSet(symshen_4_ddatatypes_d, Nil)

__e.TailApply(tmp12728, tmp12737)
return


}, 1)

tmp12738 := Call(__e, ns2_1set, sympreclude_1all_1but, tmp12727)


_ = tmp12738

tmp12739 := MakeNative(func(__e *ControlFlow) {
V3299 := __e.Get(1)
_ = V3299
tmp12740 := MakeNative(func(__e *ControlFlow) {
InternTypes := __e.Get(1)
_ = InternTypes
tmp12741 := MakeNative(func(__e *ControlFlow) {
AllDatatypes := __e.Get(1)
_ = AllDatatypes
tmp12742 := MakeNative(func(__e *ControlFlow) {
Datatypes := __e.Get(1)
_ = Datatypes
__e.TailApply(PrimFunc(symshen_4show_1datatypes), Datatypes)
return
}, 1)

tmp12743 := Call(__e, PrimFunc(symshen_4remove_1datatypes), InternTypes, AllDatatypes)


tmp12744 := PrimSet(symshen_4_ddatatypes_d, tmp12743)

__e.TailApply(tmp12742, tmp12744)
return


}, 1)

tmp12745 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(tmp12741, tmp12745)
return


}, 1)

tmp12746 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4intern_1type), X)
return
}, 1)

tmp12747 := Call(__e, PrimFunc(symmap), tmp12746, V3299)


__e.TailApply(tmp12740, tmp12747)
return


}, 1)

__e.TailApply(ns2_1set, syminclude_1all_1but, tmp12739)
return




}, 0)

