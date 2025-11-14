package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp13473 := MakeNative(func(__e *ControlFlow) {
V5419 := __e.Get(1)
_ = V5419
tmp13474 := Call(__e, PrimFunc(symshen_4app), V5419, MakeString(";\n"), symshen_4a)


tmp13475 := PrimStringConcat(MakeString("partial function "), tmp13474)

tmp13476 := Call(__e, PrimFunc(symstoutput))


tmp13477 := Call(__e, PrimFunc(sympr), tmp13475, tmp13476)


_ = tmp13477

tmp13486 := Call(__e, PrimFunc(symshen_4tracked_2), V5419)


tmp13487 := PrimNot(tmp13486)

var ifres13481 Obj

if True == tmp13487 {
tmp13483 := Call(__e, PrimFunc(symshen_4app), V5419, MakeString("? "), symshen_4a)


tmp13484 := PrimStringConcat(MakeString("track "), tmp13483)

tmp13485 := Call(__e, PrimFunc(symy_1or_1n_2), tmp13484)


var ifres13482 Obj

if True == tmp13485 {
ifres13482 = True


} else {
ifres13482 = False


}

ifres13481 = ifres13482


} else {
ifres13481 = False


}

var ifres13478 Obj

if True == ifres13481 {
tmp13479 := Call(__e, PrimFunc(symps), V5419)


tmp13480 := Call(__e, PrimFunc(symshen_4track_1function), tmp13479)


ifres13478 = tmp13480


} else {
ifres13478 = symshen_4ok


}

_ = ifres13478

__e.Return(PrimSimpleError(MakeString("aborted")))
return


}, 1)

tmp13488 := Call(__e, ns2_1set, symshen_4f_1error, tmp13473)


_ = tmp13488

tmp13489 := MakeNative(func(__e *ControlFlow) {
V5420 := __e.Get(1)
_ = V5420
tmp13490 := PrimValue(symshen_4_dtracking_d)

__e.TailApply(PrimFunc(symelement_2), V5420, tmp13490)
return


}, 1)

tmp13491 := Call(__e, ns2_1set, symshen_4tracked_2, tmp13489)


_ = tmp13491

tmp13492 := MakeNative(func(__e *ControlFlow) {
V5421 := __e.Get(1)
_ = V5421
tmp13493 := MakeNative(func(__e *ControlFlow) {
W5422 := __e.Get(1)
_ = W5422
__e.TailApply(PrimFunc(symshen_4track_1function), W5422)
return
}, 1)

tmp13494 := Call(__e, PrimFunc(symps), V5421)


__e.TailApply(tmp13493, tmp13494)
return


}, 1)

tmp13495 := Call(__e, ns2_1set, symtrack, tmp13492)


_ = tmp13495

tmp13496 := MakeNative(func(__e *ControlFlow) {
V5425 := __e.Get(1)
_ = V5425
tmp13553 := PrimIsPair(V5425)

var ifres13527 Obj

if True == tmp13553 {
tmp13551 := PrimHead(V5425)

tmp13552 := PrimEqual(symdefun, tmp13551)

var ifres13529 Obj

if True == tmp13552 {
tmp13549 := PrimTail(V5425)

tmp13550 := PrimIsPair(tmp13549)

var ifres13531 Obj

if True == tmp13550 {
tmp13546 := PrimTail(V5425)

tmp13547 := PrimTail(tmp13546)

tmp13548 := PrimIsPair(tmp13547)

var ifres13533 Obj

if True == tmp13548 {
tmp13542 := PrimTail(V5425)

tmp13543 := PrimTail(tmp13542)

tmp13544 := PrimTail(tmp13543)

tmp13545 := PrimIsPair(tmp13544)

var ifres13535 Obj

if True == tmp13545 {
tmp13537 := PrimTail(V5425)

tmp13538 := PrimTail(tmp13537)

tmp13539 := PrimTail(tmp13538)

tmp13540 := PrimTail(tmp13539)

tmp13541 := PrimEqual(Nil, tmp13540)

var ifres13536 Obj

if True == tmp13541 {
ifres13536 = True


} else {
ifres13536 = False


}

ifres13535 = ifres13536


} else {
ifres13535 = False


}

var ifres13534 Obj

if True == ifres13535 {
ifres13534 = True


} else {
ifres13534 = False


}

ifres13533 = ifres13534


} else {
ifres13533 = False


}

var ifres13532 Obj

if True == ifres13533 {
ifres13532 = True


} else {
ifres13532 = False


}

ifres13531 = ifres13532


} else {
ifres13531 = False


}

var ifres13530 Obj

if True == ifres13531 {
ifres13530 = True


} else {
ifres13530 = False


}

ifres13529 = ifres13530


} else {
ifres13529 = False


}

var ifres13528 Obj

if True == ifres13529 {
ifres13528 = True


} else {
ifres13528 = False


}

ifres13527 = ifres13528


} else {
ifres13527 = False


}

if True == ifres13527 {
tmp13497 := MakeNative(func(__e *ControlFlow) {
W5426 := __e.Get(1)
_ = W5426
tmp13498 := MakeNative(func(__e *ControlFlow) {
W5427 := __e.Get(1)
_ = W5427
tmp13499 := MakeNative(func(__e *ControlFlow) {
W5428 := __e.Get(1)
_ = W5428
tmp13500 := PrimTail(V5425)

__e.Return(PrimHead(tmp13500))
return


}, 1)

tmp13501 := PrimTail(V5425)

tmp13502 := PrimHead(tmp13501)

tmp13503 := PrimValue(symshen_4_dtracking_d)

tmp13504 := Call(__e, PrimFunc(symadjoin), tmp13502, tmp13503)


tmp13505 := PrimSet(symshen_4_dtracking_d, tmp13504)

__e.TailApply(tmp13499, tmp13505)
return


}, 1)

tmp13506 := Call(__e, PrimFunc(symeval_1kl), W5426)


__e.TailApply(tmp13498, tmp13506)
return


}, 1)

tmp13507 := PrimTail(V5425)

tmp13508 := PrimHead(tmp13507)

tmp13509 := PrimTail(V5425)

tmp13510 := PrimTail(tmp13509)

tmp13511 := PrimHead(tmp13510)

tmp13512 := PrimTail(V5425)

tmp13513 := PrimHead(tmp13512)

tmp13514 := PrimTail(V5425)

tmp13515 := PrimTail(tmp13514)

tmp13516 := PrimHead(tmp13515)

tmp13517 := PrimTail(V5425)

tmp13518 := PrimTail(tmp13517)

tmp13519 := PrimTail(tmp13518)

tmp13520 := PrimHead(tmp13519)

tmp13521 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp13513, tmp13516, tmp13520)


tmp13522 := PrimCons(tmp13521, Nil)

tmp13523 := PrimCons(tmp13511, tmp13522)

tmp13524 := PrimCons(tmp13508, tmp13523)

tmp13525 := PrimCons(symdefun, tmp13524)

__e.TailApply(tmp13497, tmp13525)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.track-function")))
return
}


}, 1)

tmp13554 := Call(__e, ns2_1set, symshen_4track_1function, tmp13496)


_ = tmp13554

tmp13555 := MakeNative(func(__e *ControlFlow) {
V5429 := __e.Get(1)
_ = V5429
V5430 := __e.Get(2)
_ = V5430
V5431 := __e.Get(3)
_ = V5431
tmp13556 := PrimCons(symshen_4_dcall_d, Nil)

tmp13557 := PrimCons(symvalue, tmp13556)

tmp13558 := PrimCons(MakeNumber(1), Nil)

tmp13559 := PrimCons(tmp13557, tmp13558)

tmp13560 := PrimCons(sym_7, tmp13559)

tmp13561 := PrimCons(tmp13560, Nil)

tmp13562 := PrimCons(symshen_4_dcall_d, tmp13561)

tmp13563 := PrimCons(symset, tmp13562)

tmp13564 := PrimCons(symshen_4_dcall_d, Nil)

tmp13565 := PrimCons(symvalue, tmp13564)

tmp13566 := Call(__e, PrimFunc(symshen_4prolog_1track), V5431, V5430)


tmp13567 := Call(__e, PrimFunc(symshen_4cons_1form), tmp13566)


tmp13568 := PrimCons(tmp13567, Nil)

tmp13569 := PrimCons(V5429, tmp13568)

tmp13570 := PrimCons(tmp13565, tmp13569)

tmp13571 := PrimCons(symshen_4input_1track, tmp13570)

tmp13572 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp13573 := PrimCons(symshen_4_dcall_d, Nil)

tmp13574 := PrimCons(symvalue, tmp13573)

tmp13575 := PrimCons(symResult, Nil)

tmp13576 := PrimCons(V5429, tmp13575)

tmp13577 := PrimCons(tmp13574, tmp13576)

tmp13578 := PrimCons(symshen_4output_1track, tmp13577)

tmp13579 := PrimCons(symshen_4_dcall_d, Nil)

tmp13580 := PrimCons(symvalue, tmp13579)

tmp13581 := PrimCons(MakeNumber(1), Nil)

tmp13582 := PrimCons(tmp13580, tmp13581)

tmp13583 := PrimCons(sym_1, tmp13582)

tmp13584 := PrimCons(tmp13583, Nil)

tmp13585 := PrimCons(symshen_4_dcall_d, tmp13584)

tmp13586 := PrimCons(symset, tmp13585)

tmp13587 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp13588 := PrimCons(symResult, Nil)

tmp13589 := PrimCons(tmp13587, tmp13588)

tmp13590 := PrimCons(symdo, tmp13589)

tmp13591 := PrimCons(tmp13590, Nil)

tmp13592 := PrimCons(tmp13586, tmp13591)

tmp13593 := PrimCons(symdo, tmp13592)

tmp13594 := PrimCons(tmp13593, Nil)

tmp13595 := PrimCons(tmp13578, tmp13594)

tmp13596 := PrimCons(symdo, tmp13595)

tmp13597 := PrimCons(tmp13596, Nil)

tmp13598 := PrimCons(V5431, tmp13597)

tmp13599 := PrimCons(symResult, tmp13598)

tmp13600 := PrimCons(symlet, tmp13599)

tmp13601 := PrimCons(tmp13600, Nil)

tmp13602 := PrimCons(tmp13572, tmp13601)

tmp13603 := PrimCons(symdo, tmp13602)

tmp13604 := PrimCons(tmp13603, Nil)

tmp13605 := PrimCons(tmp13571, tmp13604)

tmp13606 := PrimCons(symdo, tmp13605)

tmp13607 := PrimCons(tmp13606, Nil)

tmp13608 := PrimCons(tmp13563, tmp13607)

__e.Return(PrimCons(symdo, tmp13608))
return


}, 3)

tmp13609 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp13555)


_ = tmp13609

tmp13610 := MakeNative(func(__e *ControlFlow) {
V5432 := __e.Get(1)
_ = V5432
V5433 := __e.Get(2)
_ = V5433
tmp13613 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5432)


tmp13614 := PrimEqual(tmp13613, MakeNumber(0))

if True == tmp13614 {
__e.Return(V5433)
return
} else {
tmp13611 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5433)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5433, tmp13611)
return


}


}, 2)

tmp13615 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp13610)


_ = tmp13615

tmp13616 := MakeNative(func(__e *ControlFlow) {
V5436 := __e.Get(1)
_ = V5436
tmp13645 := PrimEqual(Nil, V5436)

if True == tmp13645 {
__e.Return(Nil)
return
} else {
tmp13643 := PrimIsPair(V5436)

var ifres13621 Obj

if True == tmp13643 {
tmp13641 := PrimTail(V5436)

tmp13642 := PrimIsPair(tmp13641)

var ifres13623 Obj

if True == tmp13642 {
tmp13638 := PrimTail(V5436)

tmp13639 := PrimTail(tmp13638)

tmp13640 := PrimIsPair(tmp13639)

var ifres13625 Obj

if True == tmp13640 {
tmp13634 := PrimTail(V5436)

tmp13635 := PrimTail(tmp13634)

tmp13636 := PrimTail(tmp13635)

tmp13637 := PrimIsPair(tmp13636)

var ifres13627 Obj

if True == tmp13637 {
tmp13629 := PrimTail(V5436)

tmp13630 := PrimTail(tmp13629)

tmp13631 := PrimTail(tmp13630)

tmp13632 := PrimTail(tmp13631)

tmp13633 := PrimEqual(Nil, tmp13632)

var ifres13628 Obj

if True == tmp13633 {
ifres13628 = True


} else {
ifres13628 = False


}

ifres13627 = ifres13628


} else {
ifres13627 = False


}

var ifres13626 Obj

if True == ifres13627 {
ifres13626 = True


} else {
ifres13626 = False


}

ifres13625 = ifres13626


} else {
ifres13625 = False


}

var ifres13624 Obj

if True == ifres13625 {
ifres13624 = True


} else {
ifres13624 = False


}

ifres13623 = ifres13624


} else {
ifres13623 = False


}

var ifres13622 Obj

if True == ifres13623 {
ifres13622 = True


} else {
ifres13622 = False


}

ifres13621 = ifres13622


} else {
ifres13621 = False


}

if True == ifres13621 {
__e.Return(PrimHead(V5436))
return
} else {
tmp13619 := PrimIsPair(V5436)

if True == tmp13619 {
tmp13617 := PrimTail(V5436)

__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp13617)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.vector-parameter")))
return
}


}


}


}, 1)

tmp13646 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp13616)


_ = tmp13646

tmp13647 := MakeNative(func(__e *ControlFlow) {
V5439 := __e.Get(1)
_ = V5439
V5440 := __e.Get(2)
_ = V5440
tmp13681 := PrimEqual(Nil, V5440)

if True == tmp13681 {
__e.Return(V5439)
return
} else {
tmp13679 := PrimIsPair(V5439)

var ifres13657 Obj

if True == tmp13679 {
tmp13677 := PrimTail(V5439)

tmp13678 := PrimIsPair(tmp13677)

var ifres13659 Obj

if True == tmp13678 {
tmp13674 := PrimTail(V5439)

tmp13675 := PrimTail(tmp13674)

tmp13676 := PrimIsPair(tmp13675)

var ifres13661 Obj

if True == tmp13676 {
tmp13670 := PrimTail(V5439)

tmp13671 := PrimTail(tmp13670)

tmp13672 := PrimTail(tmp13671)

tmp13673 := PrimIsPair(tmp13672)

var ifres13663 Obj

if True == tmp13673 {
tmp13665 := PrimTail(V5439)

tmp13666 := PrimTail(tmp13665)

tmp13667 := PrimTail(tmp13666)

tmp13668 := PrimTail(tmp13667)

tmp13669 := PrimEqual(Nil, tmp13668)

var ifres13664 Obj

if True == tmp13669 {
ifres13664 = True


} else {
ifres13664 = False


}

ifres13663 = ifres13664


} else {
ifres13663 = False


}

var ifres13662 Obj

if True == ifres13663 {
ifres13662 = True


} else {
ifres13662 = False


}

ifres13661 = ifres13662


} else {
ifres13661 = False


}

var ifres13660 Obj

if True == ifres13661 {
ifres13660 = True


} else {
ifres13660 = False


}

ifres13659 = ifres13660


} else {
ifres13659 = False


}

var ifres13658 Obj

if True == ifres13659 {
ifres13658 = True


} else {
ifres13658 = False


}

ifres13657 = ifres13658


} else {
ifres13657 = False


}

if True == ifres13657 {
__e.Return(V5439)
return
} else {
tmp13655 := PrimIsPair(V5439)

if True == tmp13655 {
tmp13648 := PrimHead(V5439)

tmp13649 := PrimCons(V5440, Nil)

tmp13650 := PrimCons(tmp13648, tmp13649)

tmp13651 := PrimCons(symshen_4deref, tmp13650)

tmp13652 := PrimTail(V5439)

tmp13653 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp13652, V5440)


__e.Return(PrimCons(tmp13651, tmp13653))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.vector-dereference")))
return
}


}


}


}, 2)

tmp13682 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp13647)


_ = tmp13682

tmp13683 := MakeNative(func(__e *ControlFlow) {
V5443 := __e.Get(1)
_ = V5443
tmp13687 := PrimEqual(sym_7, V5443)

if True == tmp13687 {
__e.Return(PrimSet(symshen_4_dstep_d, True))
return
} else {
tmp13685 := PrimEqual(sym_1, V5443)

if True == tmp13685 {
__e.Return(PrimSet(symshen_4_dstep_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
return
}


}


}, 1)

tmp13688 := Call(__e, ns2_1set, symstep, tmp13683)


_ = tmp13688

tmp13689 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dstep_d))
return
}, 0)

tmp13690 := Call(__e, ns2_1set, symstep_2, tmp13689)


_ = tmp13690

tmp13691 := MakeNative(func(__e *ControlFlow) {
V5446 := __e.Get(1)
_ = V5446
tmp13695 := PrimEqual(sym_7, V5446)

if True == tmp13695 {
__e.Return(PrimSet(symshen_4_dspy_d, True))
return
} else {
tmp13693 := PrimEqual(sym_1, V5446)

if True == tmp13693 {
__e.Return(PrimSet(symshen_4_dspy_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
return
}


}


}, 1)

tmp13696 := Call(__e, ns2_1set, symspy, tmp13691)


_ = tmp13696

tmp13697 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dspy_d))
return
}, 0)

tmp13698 := Call(__e, ns2_1set, symspy_2, tmp13697)


_ = tmp13698

tmp13699 := MakeNative(func(__e *ControlFlow) {
tmp13703 := PrimValue(symshen_4_dstep_d)

if True == tmp13703 {
tmp13700 := PrimValue(sym_dstinput_d)

tmp13701 := PrimReadByte(tmp13700)

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp13701)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp13704 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp13699)


_ = tmp13704

tmp13705 := MakeNative(func(__e *ControlFlow) {
V5449 := __e.Get(1)
_ = V5449
tmp13707 := PrimEqual(MakeNumber(94), V5449)

if True == tmp13707 {
__e.Return(PrimSimpleError(MakeString("aborted")))
return
} else {
__e.Return(True)
return
}


}, 1)

tmp13708 := Call(__e, ns2_1set, symshen_4check_1byte, tmp13705)


_ = tmp13708

tmp13709 := MakeNative(func(__e *ControlFlow) {
V5450 := __e.Get(1)
_ = V5450
V5451 := __e.Get(2)
_ = V5451
V5452 := __e.Get(3)
_ = V5452
tmp13710 := Call(__e, PrimFunc(symshen_4spaces), V5450)


tmp13711 := Call(__e, PrimFunc(symshen_4spaces), V5450)


tmp13712 := Call(__e, PrimFunc(symshen_4app), tmp13711, MakeString(""), symshen_4a)


tmp13713 := PrimStringConcat(MakeString(" \n"), tmp13712)

tmp13714 := Call(__e, PrimFunc(symshen_4app), V5451, tmp13713, symshen_4a)


tmp13715 := PrimStringConcat(MakeString("> Inputs to "), tmp13714)

tmp13716 := Call(__e, PrimFunc(symshen_4app), V5450, tmp13715, symshen_4a)


tmp13717 := PrimStringConcat(MakeString("<"), tmp13716)

tmp13718 := Call(__e, PrimFunc(symshen_4app), tmp13710, tmp13717, symshen_4a)


tmp13719 := PrimStringConcat(MakeString("\n"), tmp13718)

tmp13720 := Call(__e, PrimFunc(symstoutput))


tmp13721 := Call(__e, PrimFunc(sympr), tmp13719, tmp13720)


_ = tmp13721

__e.TailApply(PrimFunc(symshen_4recursively_1print), V5452)
return


}, 3)

tmp13722 := Call(__e, ns2_1set, symshen_4input_1track, tmp13709)


_ = tmp13722

tmp13723 := MakeNative(func(__e *ControlFlow) {
V5455 := __e.Get(1)
_ = V5455
tmp13733 := PrimEqual(Nil, V5455)

if True == tmp13733 {
tmp13724 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp13724)
return


} else {
tmp13731 := PrimIsPair(V5455)

if True == tmp13731 {
tmp13725 := PrimHead(V5455)

tmp13726 := Call(__e, PrimFunc(symprint), tmp13725)


_ = tmp13726

tmp13727 := Call(__e, PrimFunc(symstoutput))


tmp13728 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp13727)


_ = tmp13728

tmp13729 := PrimTail(V5455)

__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13729)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursively-print")))
return
}


}


}, 1)

tmp13734 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp13723)


_ = tmp13734

tmp13735 := MakeNative(func(__e *ControlFlow) {
V5456 := __e.Get(1)
_ = V5456
tmp13739 := PrimEqual(MakeNumber(0), V5456)

if True == tmp13739 {
__e.Return(MakeString(""))
return
} else {
tmp13736 := PrimNumberSubtract(V5456, MakeNumber(1))

tmp13737 := Call(__e, PrimFunc(symshen_4spaces), tmp13736)


__e.Return(PrimStringConcat(MakeString(" "), tmp13737))
return


}


}, 1)

tmp13740 := Call(__e, ns2_1set, symshen_4spaces, tmp13735)


_ = tmp13740

tmp13741 := MakeNative(func(__e *ControlFlow) {
V5457 := __e.Get(1)
_ = V5457
V5458 := __e.Get(2)
_ = V5458
V5459 := __e.Get(3)
_ = V5459
tmp13742 := Call(__e, PrimFunc(symshen_4spaces), V5457)


tmp13743 := Call(__e, PrimFunc(symshen_4spaces), V5457)


tmp13744 := Call(__e, PrimFunc(symshen_4app), V5459, MakeString(""), symshen_4s)


tmp13745 := PrimStringConcat(MakeString("==> "), tmp13744)

tmp13746 := Call(__e, PrimFunc(symshen_4app), tmp13743, tmp13745, symshen_4a)


tmp13747 := PrimStringConcat(MakeString(" \n"), tmp13746)

tmp13748 := Call(__e, PrimFunc(symshen_4app), V5458, tmp13747, symshen_4a)


tmp13749 := PrimStringConcat(MakeString("> Output from "), tmp13748)

tmp13750 := Call(__e, PrimFunc(symshen_4app), V5457, tmp13749, symshen_4a)


tmp13751 := PrimStringConcat(MakeString("<"), tmp13750)

tmp13752 := Call(__e, PrimFunc(symshen_4app), tmp13742, tmp13751, symshen_4a)


tmp13753 := PrimStringConcat(MakeString("\n"), tmp13752)

tmp13754 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13753, tmp13754)
return


}, 3)

tmp13755 := Call(__e, ns2_1set, symshen_4output_1track, tmp13741)


_ = tmp13755

tmp13756 := MakeNative(func(__e *ControlFlow) {
V5460 := __e.Get(1)
_ = V5460
tmp13757 := PrimValue(symshen_4_dtracking_d)

tmp13758 := Call(__e, PrimFunc(symremove), V5460, tmp13757)


tmp13759 := PrimSet(symshen_4_dtracking_d, tmp13758)

_ = tmp13759

tmp13760 := MakeNative(func(__e *ControlFlow) {
tmp13761 := Call(__e, PrimFunc(symps), V5460)


__e.TailApply(PrimFunc(symeval), tmp13761)
return


}, 0)

tmp13762 := MakeNative(func(__e *ControlFlow) {
Z5461 := __e.Get(1)
_ = Z5461
__e.Return(V5460)
return
}, 1)

tmp13763 := Call(__e, try_1catch, tmp13760, tmp13762)


_ = tmp13763

__e.Return(V5460)
return


}, 1)

tmp13764 := Call(__e, ns2_1set, symuntrack, tmp13756)


_ = tmp13764

tmp13765 := MakeNative(func(__e *ControlFlow) {
V5462 := __e.Get(1)
_ = V5462
V5463 := __e.Get(2)
_ = V5463
__e.TailApply(PrimFunc(symshen_4remove_1h), V5462, V5463, Nil)
return
}, 2)

tmp13766 := Call(__e, ns2_1set, symremove, tmp13765)


_ = tmp13766

tmp13767 := MakeNative(func(__e *ControlFlow) {
V5473 := __e.Get(1)
_ = V5473
V5474 := __e.Get(2)
_ = V5474
V5475 := __e.Get(3)
_ = V5475
tmp13782 := PrimEqual(Nil, V5474)

if True == tmp13782 {
__e.TailApply(PrimFunc(symreverse), V5475)
return
} else {
tmp13780 := PrimIsPair(V5474)

var ifres13776 Obj

if True == tmp13780 {
tmp13778 := PrimHead(V5474)

tmp13779 := PrimEqual(V5473, tmp13778)

var ifres13777 Obj

if True == tmp13779 {
ifres13777 = True


} else {
ifres13777 = False


}

ifres13776 = ifres13777


} else {
ifres13776 = False


}

if True == ifres13776 {
tmp13768 := PrimHead(V5474)

tmp13769 := PrimTail(V5474)

__e.TailApply(PrimFunc(symshen_4remove_1h), tmp13768, tmp13769, V5475)
return


} else {
tmp13774 := PrimIsPair(V5474)

if True == tmp13774 {
tmp13770 := PrimTail(V5474)

tmp13771 := PrimHead(V5474)

tmp13772 := PrimCons(tmp13771, V5475)

__e.TailApply(PrimFunc(symshen_4remove_1h), V5473, tmp13770, tmp13772)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-h")))
return
}


}


}


}, 3)

tmp13783 := Call(__e, ns2_1set, symshen_4remove_1h, tmp13767)


_ = tmp13783

tmp13784 := MakeNative(func(__e *ControlFlow) {
V5476 := __e.Get(1)
_ = V5476
tmp13785 := PrimValue(symshen_4_dprofiled_d)

tmp13786 := PrimCons(V5476, tmp13785)

tmp13787 := PrimSet(symshen_4_dprofiled_d, tmp13786)

_ = tmp13787

tmp13788 := Call(__e, PrimFunc(symps), V5476)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp13788)
return


}, 1)

tmp13789 := Call(__e, ns2_1set, symprofile, tmp13784)


_ = tmp13789

tmp13790 := MakeNative(func(__e *ControlFlow) {
V5479 := __e.Get(1)
_ = V5479
tmp13860 := PrimIsPair(V5479)

var ifres13834 Obj

if True == tmp13860 {
tmp13858 := PrimHead(V5479)

tmp13859 := PrimEqual(symdefun, tmp13858)

var ifres13836 Obj

if True == tmp13859 {
tmp13856 := PrimTail(V5479)

tmp13857 := PrimIsPair(tmp13856)

var ifres13838 Obj

if True == tmp13857 {
tmp13853 := PrimTail(V5479)

tmp13854 := PrimTail(tmp13853)

tmp13855 := PrimIsPair(tmp13854)

var ifres13840 Obj

if True == tmp13855 {
tmp13849 := PrimTail(V5479)

tmp13850 := PrimTail(tmp13849)

tmp13851 := PrimTail(tmp13850)

tmp13852 := PrimIsPair(tmp13851)

var ifres13842 Obj

if True == tmp13852 {
tmp13844 := PrimTail(V5479)

tmp13845 := PrimTail(tmp13844)

tmp13846 := PrimTail(tmp13845)

tmp13847 := PrimTail(tmp13846)

tmp13848 := PrimEqual(Nil, tmp13847)

var ifres13843 Obj

if True == tmp13848 {
ifres13843 = True


} else {
ifres13843 = False


}

ifres13842 = ifres13843


} else {
ifres13842 = False


}

var ifres13841 Obj

if True == ifres13842 {
ifres13841 = True


} else {
ifres13841 = False


}

ifres13840 = ifres13841


} else {
ifres13840 = False


}

var ifres13839 Obj

if True == ifres13840 {
ifres13839 = True


} else {
ifres13839 = False


}

ifres13838 = ifres13839


} else {
ifres13838 = False


}

var ifres13837 Obj

if True == ifres13838 {
ifres13837 = True


} else {
ifres13837 = False


}

ifres13836 = ifres13837


} else {
ifres13836 = False


}

var ifres13835 Obj

if True == ifres13836 {
ifres13835 = True


} else {
ifres13835 = False


}

ifres13834 = ifres13835


} else {
ifres13834 = False


}

if True == ifres13834 {
tmp13791 := MakeNative(func(__e *ControlFlow) {
W5480 := __e.Get(1)
_ = W5480
tmp13792 := MakeNative(func(__e *ControlFlow) {
W5481 := __e.Get(1)
_ = W5481
tmp13793 := MakeNative(func(__e *ControlFlow) {
W5482 := __e.Get(1)
_ = W5482
tmp13794 := MakeNative(func(__e *ControlFlow) {
W5483 := __e.Get(1)
_ = W5483
tmp13795 := MakeNative(func(__e *ControlFlow) {
W5484 := __e.Get(1)
_ = W5484
tmp13796 := PrimTail(V5479)

__e.Return(PrimHead(tmp13796))
return


}, 1)

tmp13797 := Call(__e, PrimFunc(symeval_1kl), W5482)


__e.TailApply(tmp13795, tmp13797)
return


}, 1)

tmp13798 := Call(__e, PrimFunc(symeval_1kl), W5481)


__e.TailApply(tmp13794, tmp13798)
return


}, 1)

tmp13799 := PrimTail(V5479)

tmp13800 := PrimTail(tmp13799)

tmp13801 := PrimHead(tmp13800)

tmp13802 := PrimTail(V5479)

tmp13803 := PrimHead(tmp13802)

tmp13804 := PrimTail(V5479)

tmp13805 := PrimTail(tmp13804)

tmp13806 := PrimTail(tmp13805)

tmp13807 := PrimHead(tmp13806)

tmp13808 := Call(__e, PrimFunc(symsubst), W5480, tmp13803, tmp13807)


tmp13809 := PrimCons(tmp13808, Nil)

tmp13810 := PrimCons(tmp13801, tmp13809)

tmp13811 := PrimCons(W5480, tmp13810)

tmp13812 := PrimCons(symdefun, tmp13811)

__e.TailApply(tmp13793, tmp13812)
return


}, 1)

tmp13813 := PrimTail(V5479)

tmp13814 := PrimHead(tmp13813)

tmp13815 := PrimTail(V5479)

tmp13816 := PrimTail(tmp13815)

tmp13817 := PrimHead(tmp13816)

tmp13818 := PrimTail(V5479)

tmp13819 := PrimHead(tmp13818)

tmp13820 := PrimTail(V5479)

tmp13821 := PrimTail(tmp13820)

tmp13822 := PrimHead(tmp13821)

tmp13823 := PrimTail(V5479)

tmp13824 := PrimTail(tmp13823)

tmp13825 := PrimHead(tmp13824)

tmp13826 := PrimCons(W5480, tmp13825)

tmp13827 := Call(__e, PrimFunc(symshen_4profile_1func), tmp13819, tmp13822, tmp13826)


tmp13828 := PrimCons(tmp13827, Nil)

tmp13829 := PrimCons(tmp13817, tmp13828)

tmp13830 := PrimCons(tmp13814, tmp13829)

tmp13831 := PrimCons(symdefun, tmp13830)

__e.TailApply(tmp13792, tmp13831)
return


}, 1)

tmp13832 := Call(__e, PrimFunc(symgensym), symshen_4f)


__e.TailApply(tmp13791, tmp13832)
return


} else {
__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
return
}


}, 1)

tmp13861 := Call(__e, ns2_1set, symshen_4profile_1help, tmp13790)


_ = tmp13861

tmp13862 := MakeNative(func(__e *ControlFlow) {
V5485 := __e.Get(1)
_ = V5485
tmp13863 := PrimValue(symshen_4_dprofiled_d)

tmp13864 := Call(__e, PrimFunc(symremove), V5485, tmp13863)


tmp13865 := PrimSet(symshen_4_dprofiled_d, tmp13864)

_ = tmp13865

tmp13866 := MakeNative(func(__e *ControlFlow) {
tmp13867 := Call(__e, PrimFunc(symps), V5485)


__e.TailApply(PrimFunc(symeval), tmp13867)
return


}, 0)

tmp13868 := MakeNative(func(__e *ControlFlow) {
Z5486 := __e.Get(1)
_ = Z5486
__e.Return(V5485)
return
}, 1)

__e.TailApply(try_1catch, tmp13866, tmp13868)
return


}, 1)

tmp13869 := Call(__e, ns2_1set, symunprofile, tmp13862)


_ = tmp13869

tmp13870 := MakeNative(func(__e *ControlFlow) {
V5487 := __e.Get(1)
_ = V5487
tmp13871 := PrimValue(symshen_4_dprofiled_d)

__e.TailApply(PrimFunc(symelement_2), V5487, tmp13871)
return


}, 1)

tmp13872 := Call(__e, ns2_1set, symshen_4profiled_2, tmp13870)


_ = tmp13872

tmp13873 := MakeNative(func(__e *ControlFlow) {
V5488 := __e.Get(1)
_ = V5488
V5489 := __e.Get(2)
_ = V5489
V5490 := __e.Get(3)
_ = V5490
tmp13874 := PrimCons(symrun, Nil)

tmp13875 := PrimCons(symget_1time, tmp13874)

tmp13876 := PrimCons(symrun, Nil)

tmp13877 := PrimCons(symget_1time, tmp13876)

tmp13878 := PrimCons(symStart, Nil)

tmp13879 := PrimCons(tmp13877, tmp13878)

tmp13880 := PrimCons(sym_1, tmp13879)

tmp13881 := PrimCons(V5488, Nil)

tmp13882 := PrimCons(symshen_4get_1profile, tmp13881)

tmp13883 := PrimCons(symFinish, Nil)

tmp13884 := PrimCons(tmp13882, tmp13883)

tmp13885 := PrimCons(sym_7, tmp13884)

tmp13886 := PrimCons(tmp13885, Nil)

tmp13887 := PrimCons(V5488, tmp13886)

tmp13888 := PrimCons(symshen_4put_1profile, tmp13887)

tmp13889 := PrimCons(symResult, Nil)

tmp13890 := PrimCons(tmp13888, tmp13889)

tmp13891 := PrimCons(symRecord, tmp13890)

tmp13892 := PrimCons(symlet, tmp13891)

tmp13893 := PrimCons(tmp13892, Nil)

tmp13894 := PrimCons(tmp13880, tmp13893)

tmp13895 := PrimCons(symFinish, tmp13894)

tmp13896 := PrimCons(symlet, tmp13895)

tmp13897 := PrimCons(tmp13896, Nil)

tmp13898 := PrimCons(V5490, tmp13897)

tmp13899 := PrimCons(symResult, tmp13898)

tmp13900 := PrimCons(symlet, tmp13899)

tmp13901 := PrimCons(tmp13900, Nil)

tmp13902 := PrimCons(tmp13875, tmp13901)

tmp13903 := PrimCons(symStart, tmp13902)

__e.Return(PrimCons(symlet, tmp13903))
return


}, 3)

tmp13904 := Call(__e, ns2_1set, symshen_4profile_1func, tmp13873)


_ = tmp13904

tmp13905 := MakeNative(func(__e *ControlFlow) {
V5491 := __e.Get(1)
_ = V5491
tmp13906 := MakeNative(func(__e *ControlFlow) {
W5492 := __e.Get(1)
_ = W5492
tmp13907 := MakeNative(func(__e *ControlFlow) {
W5493 := __e.Get(1)
_ = W5493
__e.TailApply(PrimFunc(sym_8p), V5491, W5492)
return
}, 1)

tmp13908 := Call(__e, PrimFunc(symshen_4put_1profile), V5491, MakeNumber(0))


__e.TailApply(tmp13907, tmp13908)
return


}, 1)

tmp13909 := Call(__e, PrimFunc(symshen_4get_1profile), V5491)


__e.TailApply(tmp13906, tmp13909)
return


}, 1)

tmp13910 := Call(__e, ns2_1set, symprofile_1results, tmp13905)


_ = tmp13910

tmp13911 := MakeNative(func(__e *ControlFlow) {
V5494 := __e.Get(1)
_ = V5494
tmp13912 := MakeNative(func(__e *ControlFlow) {
tmp13913 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V5494, symprofile, tmp13913)
return


}, 0)

tmp13914 := MakeNative(func(__e *ControlFlow) {
Z5495 := __e.Get(1)
_ = Z5495
__e.Return(MakeNumber(0))
return
}, 1)

__e.TailApply(try_1catch, tmp13912, tmp13914)
return


}, 1)

tmp13915 := Call(__e, ns2_1set, symshen_4get_1profile, tmp13911)


_ = tmp13915

tmp13916 := MakeNative(func(__e *ControlFlow) {
V5496 := __e.Get(1)
_ = V5496
V5497 := __e.Get(2)
_ = V5497
tmp13917 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V5496, symprofile, V5497, tmp13917)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp13916)
return




}, 0)

