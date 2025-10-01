package main

import . "github.com/tiancaiamao/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp17505 := MakeNative(func(__e *ControlFlow) {
V4662 := __e.Get(1)
_ = V4662
V4663 := __e.Get(2)
_ = V4663
tmp17506 := MakeNative(func(__e *ControlFlow) {
Rectify := __e.Get(1)
_ = Rectify
tmp17507 := MakeNative(func(__e *ControlFlow) {
Variant_2 := __e.Get(1)
_ = Variant_2
tmp17508 := MakeNative(func(__e *ControlFlow) {
Abstraction := __e.Get(1)
_ = Abstraction
tmp17509 := MakeNative(func(__e *ControlFlow) {
UpDate := __e.Get(1)
_ = UpDate
__e.Return(V4662)
return
}, 1)

tmp17510 := PrimValue(symshen_4_dsigf_d)

tmp17511 := Call(__e, PrimFunc(symshen_4assoc_1_6), V4662, Abstraction, tmp17510)


tmp17512 := PrimSet(symshen_4_dsigf_d, tmp17511)

__e.TailApply(tmp17509, tmp17512)
return


}, 1)

tmp17513 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V4663)


tmp17514 := Call(__e, PrimFunc(symeval_1kl), tmp17513)


__e.TailApply(tmp17508, tmp17514)
return


}, 1)

tmp17515 := MakeNative(func(__e *ControlFlow) {
V4624 := __e.Get(1)
_ = V4624
__e.Return(MakeNative(func(__e *ControlFlow) {
L4625 := __e.Get(1)
_ = L4625
__e.Return(MakeNative(func(__e *ControlFlow) {
K4626 := __e.Get(1)
_ = K4626
__e.Return(MakeNative(func(__e *ControlFlow) {
C4627 := __e.Get(1)
_ = C4627
tmp17516 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17516

tmp17517 := Call(__e, PrimFunc(symshen_4deref), V4662, V4624)


tmp17518 := Call(__e, PrimFunc(symreceive), tmp17517)


tmp17519 := Call(__e, PrimFunc(symshen_4deref), Rectify, V4624)


tmp17520 := Call(__e, PrimFunc(symreceive), tmp17519)


__e.TailApply(PrimFunc(symshen_4variancy), tmp17518, tmp17520, V4624, L4625, K4626, C4627)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp17521 := Call(__e, PrimFunc(symshen_4reset_1prolog_1vector))


tmp17522 := Call(__e, tmp17515, tmp17521)


tmp17523 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp17524 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp17523)


tmp17525 := Call(__e, PrimFunc(sym_8v), True, tmp17524)


tmp17526 := Call(__e, tmp17522, tmp17525)


tmp17527 := Call(__e, tmp17526, MakeNumber(0))


tmp17528 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp17529 := Call(__e, tmp17527, tmp17528)


__e.TailApply(tmp17507, tmp17529)
return


}, 1)

tmp17530 := Call(__e, PrimFunc(symshen_4rectify_1type), V4663)


__e.TailApply(tmp17506, tmp17530)
return


}, 2)

tmp17531 := Call(__e, ns2_1set, symdeclare, tmp17505)


_ = tmp17531

tmp17532 := MakeNative(func(__e *ControlFlow) {
V4664 := __e.Get(1)
_ = V4664
V4665 := __e.Get(2)
_ = V4665
V4666 := __e.Get(3)
_ = V4666
V4667 := __e.Get(4)
_ = V4667
V4668 := __e.Get(5)
_ = V4668
V4669 := __e.Get(6)
_ = V4669
tmp17539 := Call(__e, PrimFunc(symshen_4unlocked_2), V4667)


if True == tmp17539 {
tmp17533 := MakeNative(func(__e *ControlFlow) {
A := __e.Get(1)
_ = A
tmp17534 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17534

tmp17535 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V4664, A, V4665, V4666, V4667, V4668, V4669)
return
}, 0)

tmp17536 := Call(__e, PrimFunc(symshen_4system_1S_1h), V4664, A, Nil, V4666, V4667, V4668, tmp17535)


__e.TailApply(PrimFunc(symshen_4gc), V4666, tmp17536)
return


}, 1)

tmp17537 := Call(__e, PrimFunc(symshen_4newpv), V4666)


__e.TailApply(tmp17533, tmp17537)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp17540 := Call(__e, ns2_1set, symshen_4variancy, tmp17532)


_ = tmp17540

tmp17541 := MakeNative(func(__e *ControlFlow) {
V4670 := __e.Get(1)
_ = V4670
V4671 := __e.Get(2)
_ = V4671
V4672 := __e.Get(3)
_ = V4672
V4673 := __e.Get(4)
_ = V4673
V4674 := __e.Get(5)
_ = V4674
V4675 := __e.Get(6)
_ = V4675
V4676 := __e.Get(7)
_ = V4676
tmp17542 := MakeNative(func(__e *ControlFlow) {
K4637 := __e.Get(1)
_ = K4637
tmp17543 := MakeNative(func(__e *ControlFlow) {
C4642 := __e.Get(1)
_ = C4642
tmp17567 := PrimEqual(C4642, False)

if True == tmp17567 {
tmp17544 := MakeNative(func(__e *ControlFlow) {
C4645 := __e.Get(1)
_ = C4645
tmp17561 := PrimEqual(C4645, False)

if True == tmp17561 {
tmp17545 := MakeNative(func(__e *ControlFlow) {
C4646 := __e.Get(1)
_ = C4646
tmp17547 := PrimEqual(C4646, False)

if True == tmp17547 {
__e.TailApply(PrimFunc(symshen_4unlock), V4674, K4637)
return
} else {
__e.Return(C4646)
return
}


}, 1)

tmp17559 := Call(__e, PrimFunc(symshen_4unlocked_2), V4674)


var ifres17548 Obj

if True == tmp17559 {
tmp17549 := MakeNative(func(__e *ControlFlow) {
Warning := __e.Get(1)
_ = Warning
tmp17550 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17550

tmp17551 := Call(__e, PrimFunc(symshen_4deref), V4670, V4673)


tmp17552 := Call(__e, PrimFunc(symshen_4app), tmp17551, MakeString(" may create errors\n"), symshen_4a)


tmp17553 := PrimStringConcat(MakeString("warning: changing the type of "), tmp17552)

tmp17554 := Call(__e, PrimFunc(symstoutput))


tmp17555 := Call(__e, PrimFunc(sympr), tmp17553, tmp17554)


tmp17556 := Call(__e, PrimFunc(symis), Warning, tmp17555, V4673, V4674, K4637, V4676)


__e.TailApply(PrimFunc(symshen_4gc), V4673, tmp17556)
return


}, 1)

tmp17557 := Call(__e, PrimFunc(symshen_4newpv), V4673)


tmp17558 := Call(__e, tmp17549, tmp17557)


ifres17548 = tmp17558


} else {
ifres17548 = False


}

__e.TailApply(tmp17545, ifres17548)
return


} else {
__e.Return(C4645)
return
}


}, 1)

tmp17565 := Call(__e, PrimFunc(symshen_4unlocked_2), V4674)


var ifres17562 Obj

if True == tmp17565 {
tmp17563 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17563

tmp17564 := Call(__e, PrimFunc(symis_b), V4671, V4672, V4673, V4674, K4637, V4676)


ifres17562 = tmp17564


} else {
ifres17562 = False


}

__e.TailApply(tmp17544, ifres17562)
return


} else {
__e.Return(C4642)
return
}


}, 1)

tmp17579 := Call(__e, PrimFunc(symshen_4unlocked_2), V4674)


var ifres17568 Obj

if True == tmp17579 {
tmp17569 := MakeNative(func(__e *ControlFlow) {
Tm4643 := __e.Get(1)
_ = Tm4643
tmp17570 := MakeNative(func(__e *ControlFlow) {
GoTo4644 := __e.Get(1)
_ = GoTo4644
tmp17574 := PrimEqual(Tm4643, symsymbol)

if True == tmp17574 {
__e.TailApply(PrimFunc(symthaw), GoTo4644)
return
} else {
tmp17572 := Call(__e, PrimFunc(symshen_4pvar_2), Tm4643)


if True == tmp17572 {
__e.TailApply(PrimFunc(symshen_4bind_b), Tm4643, symsymbol, V4673, GoTo4644)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp17575 := MakeNative(func(__e *ControlFlow) {
tmp17576 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp17576

__e.TailApply(PrimFunc(symshen_4cut), V4673, V4674, K4637, V4676)
return


}, 0)

__e.TailApply(tmp17570, tmp17575)
return


}, 1)

tmp17577 := Call(__e, PrimFunc(symshen_4lazyderef), V4671, V4673)


tmp17578 := Call(__e, tmp17569, tmp17577)


ifres17568 = tmp17578


} else {
ifres17568 = False


}

__e.TailApply(tmp17543, ifres17568)
return


}, 1)

tmp17580 := PrimNumberAdd(V4675, MakeNumber(1))

__e.TailApply(tmp17542, tmp17580)
return


}, 7)

tmp17581 := Call(__e, ns2_1set, symshen_4variants_2, tmp17541)


_ = tmp17581

tmp17582 := MakeNative(func(__e *ControlFlow) {
V4677 := __e.Get(1)
_ = V4677
tmp17583 := MakeNative(func(__e *ControlFlow) {
Bindings := __e.Get(1)
_ = Bindings
tmp17584 := MakeNative(func(__e *ControlFlow) {
Lock := __e.Get(1)
_ = Lock
tmp17585 := MakeNative(func(__e *ControlFlow) {
Key := __e.Get(1)
_ = Key
tmp17586 := MakeNative(func(__e *ControlFlow) {
Continuation := __e.Get(1)
_ = Continuation
tmp17587 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
tmp17588 := MakeNative(func(__e *ControlFlow) {
Vs := __e.Get(1)
_ = Vs
tmp17589 := Call(__e, PrimFunc(symshen_4rcons__form), V4677)


tmp17590 := PrimCons(Continuation, Nil)

tmp17591 := PrimCons(Key, tmp17590)

tmp17592 := PrimCons(Lock, tmp17591)

tmp17593 := PrimCons(Bindings, tmp17592)

tmp17594 := PrimCons(tmp17589, tmp17593)

tmp17595 := PrimCons(V, tmp17594)

tmp17596 := PrimCons(symis_b, tmp17595)

tmp17597 := Call(__e, PrimFunc(symshen_4stpart), Vs, tmp17596, Bindings)


tmp17598 := PrimCons(tmp17597, Nil)

tmp17599 := PrimCons(Continuation, tmp17598)

tmp17600 := PrimCons(symlambda, tmp17599)

tmp17601 := PrimCons(tmp17600, Nil)

tmp17602 := PrimCons(Key, tmp17601)

tmp17603 := PrimCons(symlambda, tmp17602)

tmp17604 := PrimCons(tmp17603, Nil)

tmp17605 := PrimCons(Lock, tmp17604)

tmp17606 := PrimCons(symlambda, tmp17605)

tmp17607 := PrimCons(tmp17606, Nil)

tmp17608 := PrimCons(Bindings, tmp17607)

tmp17609 := PrimCons(symlambda, tmp17608)

tmp17610 := PrimCons(tmp17609, Nil)

tmp17611 := PrimCons(V, tmp17610)

__e.Return(PrimCons(symlambda, tmp17611))
return


}, 1)

tmp17612 := Call(__e, PrimFunc(symshen_4extract_1vars), V4677)


__e.TailApply(tmp17588, tmp17612)
return


}, 1)

tmp17613 := Call(__e, PrimFunc(symprotect), symV)


tmp17614 := Call(__e, PrimFunc(symgensym), tmp17613)


__e.TailApply(tmp17587, tmp17614)
return


}, 1)

tmp17615 := Call(__e, PrimFunc(symprotect), symC)


tmp17616 := Call(__e, PrimFunc(symgensym), tmp17615)


__e.TailApply(tmp17586, tmp17616)
return


}, 1)

tmp17617 := Call(__e, PrimFunc(symprotect), symKey)


tmp17618 := Call(__e, PrimFunc(symgensym), tmp17617)


__e.TailApply(tmp17585, tmp17618)
return


}, 1)

tmp17619 := Call(__e, PrimFunc(symprotect), symL)


tmp17620 := Call(__e, PrimFunc(symgensym), tmp17619)


__e.TailApply(tmp17584, tmp17620)
return


}, 1)

tmp17621 := Call(__e, PrimFunc(symprotect), symB)


tmp17622 := Call(__e, PrimFunc(symgensym), tmp17621)


__e.TailApply(tmp17583, tmp17622)
return


}, 1)

tmp17623 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp17582)


_ = tmp17623

tmp17624 := MakeNative(func(__e *ControlFlow) {
V4678 := __e.Get(1)
_ = V4678
__e.Return(V4678)
return
}, 1)

tmp17625 := Call(__e, ns2_1set, symshen_4demod, tmp17624)


_ = tmp17625

tmp17626 := PrimCons(symA, Nil)

tmp17627 := PrimCons(sym_1_1_6, tmp17626)

tmp17628 := Call(__e, PrimFunc(symdeclare), symabort, tmp17627)


_ = tmp17628

tmp17629 := PrimCons(symboolean, Nil)

tmp17630 := PrimCons(sym_1_1_6, tmp17629)

tmp17631 := PrimCons(symA, tmp17630)

tmp17632 := Call(__e, PrimFunc(symdeclare), symabsvector_2, tmp17631)


_ = tmp17632

tmp17633 := PrimCons(symA, Nil)

tmp17634 := PrimCons(symlist, tmp17633)

tmp17635 := PrimCons(symA, Nil)

tmp17636 := PrimCons(symlist, tmp17635)

tmp17637 := PrimCons(tmp17636, Nil)

tmp17638 := PrimCons(sym_1_1_6, tmp17637)

tmp17639 := PrimCons(tmp17634, tmp17638)

tmp17640 := PrimCons(tmp17639, Nil)

tmp17641 := PrimCons(sym_1_1_6, tmp17640)

tmp17642 := PrimCons(symA, tmp17641)

tmp17643 := Call(__e, PrimFunc(symdeclare), symadjoin, tmp17642)


_ = tmp17643

tmp17644 := PrimCons(symboolean, Nil)

tmp17645 := PrimCons(sym_1_1_6, tmp17644)

tmp17646 := PrimCons(symboolean, tmp17645)

tmp17647 := PrimCons(tmp17646, Nil)

tmp17648 := PrimCons(sym_1_1_6, tmp17647)

tmp17649 := PrimCons(symboolean, tmp17648)

tmp17650 := Call(__e, PrimFunc(symdeclare), symand, tmp17649)


_ = tmp17650

tmp17651 := PrimCons(symstring, Nil)

tmp17652 := PrimCons(sym_1_1_6, tmp17651)

tmp17653 := PrimCons(symsymbol, tmp17652)

tmp17654 := PrimCons(tmp17653, Nil)

tmp17655 := PrimCons(sym_1_1_6, tmp17654)

tmp17656 := PrimCons(symstring, tmp17655)

tmp17657 := PrimCons(tmp17656, Nil)

tmp17658 := PrimCons(sym_1_1_6, tmp17657)

tmp17659 := PrimCons(symA, tmp17658)

tmp17660 := Call(__e, PrimFunc(symdeclare), symshen_4app, tmp17659)


_ = tmp17660

tmp17661 := PrimCons(symA, Nil)

tmp17662 := PrimCons(symlist, tmp17661)

tmp17663 := PrimCons(symA, Nil)

tmp17664 := PrimCons(symlist, tmp17663)

tmp17665 := PrimCons(symA, Nil)

tmp17666 := PrimCons(symlist, tmp17665)

tmp17667 := PrimCons(tmp17666, Nil)

tmp17668 := PrimCons(sym_1_1_6, tmp17667)

tmp17669 := PrimCons(tmp17664, tmp17668)

tmp17670 := PrimCons(tmp17669, Nil)

tmp17671 := PrimCons(sym_1_1_6, tmp17670)

tmp17672 := PrimCons(tmp17662, tmp17671)

tmp17673 := Call(__e, PrimFunc(symdeclare), symappend, tmp17672)


_ = tmp17673

tmp17674 := PrimCons(symnumber, Nil)

tmp17675 := PrimCons(sym_1_1_6, tmp17674)

tmp17676 := PrimCons(symA, tmp17675)

tmp17677 := Call(__e, PrimFunc(symdeclare), symarity, tmp17676)


_ = tmp17677

tmp17678 := PrimCons(symA, Nil)

tmp17679 := PrimCons(symlist, tmp17678)

tmp17680 := PrimCons(tmp17679, Nil)

tmp17681 := PrimCons(symlist, tmp17680)

tmp17682 := PrimCons(symA, Nil)

tmp17683 := PrimCons(symlist, tmp17682)

tmp17684 := PrimCons(tmp17683, Nil)

tmp17685 := PrimCons(sym_1_1_6, tmp17684)

tmp17686 := PrimCons(tmp17681, tmp17685)

tmp17687 := PrimCons(tmp17686, Nil)

tmp17688 := PrimCons(sym_1_1_6, tmp17687)

tmp17689 := PrimCons(symA, tmp17688)

tmp17690 := Call(__e, PrimFunc(symdeclare), symassoc, tmp17689)


_ = tmp17690

tmp17691 := PrimCons(symboolean, Nil)

tmp17692 := PrimCons(sym_1_1_6, tmp17691)

tmp17693 := PrimCons(symA, tmp17692)

tmp17694 := Call(__e, PrimFunc(symdeclare), symatom_2, tmp17693)


_ = tmp17694

tmp17695 := PrimCons(symboolean, Nil)

tmp17696 := PrimCons(sym_1_1_6, tmp17695)

tmp17697 := PrimCons(symA, tmp17696)

tmp17698 := Call(__e, PrimFunc(symdeclare), symboolean_2, tmp17697)


_ = tmp17698

tmp17699 := PrimCons(symstring, Nil)

tmp17700 := PrimCons(sym_1_1_6, tmp17699)

tmp17701 := PrimCons(symstring, tmp17700)

tmp17702 := Call(__e, PrimFunc(symdeclare), symbootstrap, tmp17701)


_ = tmp17702

tmp17703 := PrimCons(symboolean, Nil)

tmp17704 := PrimCons(sym_1_1_6, tmp17703)

tmp17705 := PrimCons(symsymbol, tmp17704)

tmp17706 := Call(__e, PrimFunc(symdeclare), symbound_2, tmp17705)


_ = tmp17706

tmp17707 := PrimCons(symstring, Nil)

tmp17708 := PrimCons(sym_1_1_6, tmp17707)

tmp17709 := PrimCons(symstring, tmp17708)

tmp17710 := Call(__e, PrimFunc(symdeclare), symcd, tmp17709)


_ = tmp17710

tmp17711 := PrimCons(symA, Nil)

tmp17712 := PrimCons(symstream, tmp17711)

tmp17713 := PrimCons(symB, Nil)

tmp17714 := PrimCons(symlist, tmp17713)

tmp17715 := PrimCons(tmp17714, Nil)

tmp17716 := PrimCons(sym_1_1_6, tmp17715)

tmp17717 := PrimCons(tmp17712, tmp17716)

tmp17718 := Call(__e, PrimFunc(symdeclare), symclose, tmp17717)


_ = tmp17718

tmp17719 := PrimCons(symstring, Nil)

tmp17720 := PrimCons(sym_1_1_6, tmp17719)

tmp17721 := PrimCons(symstring, tmp17720)

tmp17722 := PrimCons(tmp17721, Nil)

tmp17723 := PrimCons(sym_1_1_6, tmp17722)

tmp17724 := PrimCons(symstring, tmp17723)

tmp17725 := Call(__e, PrimFunc(symdeclare), symcn, tmp17724)


_ = tmp17725

tmp17726 := PrimCons(symA, Nil)

tmp17727 := PrimCons(symlist, tmp17726)

tmp17728 := PrimCons(symB, Nil)

tmp17729 := PrimCons(tmp17727, tmp17728)

tmp17730 := PrimCons(symstr, tmp17729)

tmp17731 := PrimCons(symA, Nil)

tmp17732 := PrimCons(symlist, tmp17731)

tmp17733 := PrimCons(symC, Nil)

tmp17734 := PrimCons(tmp17732, tmp17733)

tmp17735 := PrimCons(symstr, tmp17734)

tmp17736 := PrimCons(tmp17735, Nil)

tmp17737 := PrimCons(sym_1_1_6, tmp17736)

tmp17738 := PrimCons(tmp17730, tmp17737)

tmp17739 := PrimCons(symA, Nil)

tmp17740 := PrimCons(symlist, tmp17739)

tmp17741 := PrimCons(symC, Nil)

tmp17742 := PrimCons(sym_1_1_6, tmp17741)

tmp17743 := PrimCons(tmp17740, tmp17742)

tmp17744 := PrimCons(tmp17743, Nil)

tmp17745 := PrimCons(sym_1_1_6, tmp17744)

tmp17746 := PrimCons(tmp17738, tmp17745)

tmp17747 := Call(__e, PrimFunc(symdeclare), symcompile, tmp17746)


_ = tmp17747

tmp17748 := PrimCons(symboolean, Nil)

tmp17749 := PrimCons(sym_1_1_6, tmp17748)

tmp17750 := PrimCons(symA, tmp17749)

tmp17751 := Call(__e, PrimFunc(symdeclare), symcons_2, tmp17750)


_ = tmp17751

tmp17752 := PrimCons(symB, Nil)

tmp17753 := PrimCons(sym_1_1_6, tmp17752)

tmp17754 := PrimCons(symA, tmp17753)

tmp17755 := PrimCons(symsymbol, Nil)

tmp17756 := PrimCons(sym_1_1_6, tmp17755)

tmp17757 := PrimCons(tmp17754, tmp17756)

tmp17758 := Call(__e, PrimFunc(symdeclare), symdestroy, tmp17757)


_ = tmp17758

tmp17759 := PrimCons(symA, Nil)

tmp17760 := PrimCons(symlist, tmp17759)

tmp17761 := PrimCons(symA, Nil)

tmp17762 := PrimCons(symlist, tmp17761)

tmp17763 := PrimCons(symA, Nil)

tmp17764 := PrimCons(symlist, tmp17763)

tmp17765 := PrimCons(tmp17764, Nil)

tmp17766 := PrimCons(sym_1_1_6, tmp17765)

tmp17767 := PrimCons(tmp17762, tmp17766)

tmp17768 := PrimCons(tmp17767, Nil)

tmp17769 := PrimCons(sym_1_1_6, tmp17768)

tmp17770 := PrimCons(tmp17760, tmp17769)

tmp17771 := Call(__e, PrimFunc(symdeclare), symdifference, tmp17770)


_ = tmp17771

tmp17772 := PrimCons(symB, Nil)

tmp17773 := PrimCons(sym_1_1_6, tmp17772)

tmp17774 := PrimCons(symB, tmp17773)

tmp17775 := PrimCons(tmp17774, Nil)

tmp17776 := PrimCons(sym_1_1_6, tmp17775)

tmp17777 := PrimCons(symA, tmp17776)

tmp17778 := Call(__e, PrimFunc(symdeclare), symdo, tmp17777)


_ = tmp17778

tmp17779 := PrimCons(symA, Nil)

tmp17780 := PrimCons(symlist, tmp17779)

tmp17781 := PrimCons(symB, Nil)

tmp17782 := PrimCons(tmp17780, tmp17781)

tmp17783 := PrimCons(symstr, tmp17782)

tmp17784 := PrimCons(symA, Nil)

tmp17785 := PrimCons(symlist, tmp17784)

tmp17786 := PrimCons(symC, Nil)

tmp17787 := PrimCons(symlist, tmp17786)

tmp17788 := PrimCons(tmp17787, Nil)

tmp17789 := PrimCons(tmp17785, tmp17788)

tmp17790 := PrimCons(symstr, tmp17789)

tmp17791 := PrimCons(tmp17790, Nil)

tmp17792 := PrimCons(sym_1_1_6, tmp17791)

tmp17793 := PrimCons(tmp17783, tmp17792)

tmp17794 := Call(__e, PrimFunc(symdeclare), sym_5e_6, tmp17793)


_ = tmp17794

tmp17795 := PrimCons(symA, Nil)

tmp17796 := PrimCons(symlist, tmp17795)

tmp17797 := PrimCons(symB, Nil)

tmp17798 := PrimCons(tmp17796, tmp17797)

tmp17799 := PrimCons(symstr, tmp17798)

tmp17800 := PrimCons(symA, Nil)

tmp17801 := PrimCons(symlist, tmp17800)

tmp17802 := PrimCons(symA, Nil)

tmp17803 := PrimCons(symlist, tmp17802)

tmp17804 := PrimCons(tmp17803, Nil)

tmp17805 := PrimCons(tmp17801, tmp17804)

tmp17806 := PrimCons(symstr, tmp17805)

tmp17807 := PrimCons(tmp17806, Nil)

tmp17808 := PrimCons(sym_1_1_6, tmp17807)

tmp17809 := PrimCons(tmp17799, tmp17808)

tmp17810 := Call(__e, PrimFunc(symdeclare), sym_5_b_6, tmp17809)


_ = tmp17810

tmp17811 := PrimCons(symA, Nil)

tmp17812 := PrimCons(symlist, tmp17811)

tmp17813 := PrimCons(symB, Nil)

tmp17814 := PrimCons(tmp17812, tmp17813)

tmp17815 := PrimCons(symstr, tmp17814)

tmp17816 := PrimCons(symA, Nil)

tmp17817 := PrimCons(symlist, tmp17816)

tmp17818 := PrimCons(symB, Nil)

tmp17819 := PrimCons(tmp17817, tmp17818)

tmp17820 := PrimCons(symstr, tmp17819)

tmp17821 := PrimCons(tmp17820, Nil)

tmp17822 := PrimCons(sym_1_1_6, tmp17821)

tmp17823 := PrimCons(tmp17815, tmp17822)

tmp17824 := Call(__e, PrimFunc(symdeclare), symshen_4_5end_6, tmp17823)


_ = tmp17824

tmp17825 := PrimCons(symA, Nil)

tmp17826 := PrimCons(symlist, tmp17825)

tmp17827 := PrimCons(symB, Nil)

tmp17828 := PrimCons(tmp17826, tmp17827)

tmp17829 := PrimCons(symstr, tmp17828)

tmp17830 := PrimCons(symboolean, Nil)

tmp17831 := PrimCons(sym_1_1_6, tmp17830)

tmp17832 := PrimCons(symA, tmp17831)

tmp17833 := PrimCons(tmp17832, Nil)

tmp17834 := PrimCons(sym_1_1_6, tmp17833)

tmp17835 := PrimCons(tmp17829, tmp17834)

tmp17836 := Call(__e, PrimFunc(symdeclare), symshen_4_ahd_2, tmp17835)


_ = tmp17836

tmp17837 := PrimCons(symA, Nil)

tmp17838 := PrimCons(symlist, tmp17837)

tmp17839 := PrimCons(symB, Nil)

tmp17840 := PrimCons(tmp17838, tmp17839)

tmp17841 := PrimCons(symstr, tmp17840)

tmp17842 := PrimCons(symA, Nil)

tmp17843 := PrimCons(sym_1_1_6, tmp17842)

tmp17844 := PrimCons(tmp17841, tmp17843)

tmp17845 := Call(__e, PrimFunc(symdeclare), symshen_4hds, tmp17844)


_ = tmp17845

tmp17846 := PrimCons(symA, Nil)

tmp17847 := PrimCons(symlist, tmp17846)

tmp17848 := PrimCons(symB, Nil)

tmp17849 := PrimCons(tmp17847, tmp17848)

tmp17850 := PrimCons(symstr, tmp17849)

tmp17851 := PrimCons(symA, Nil)

tmp17852 := PrimCons(symlist, tmp17851)

tmp17853 := PrimCons(symB, Nil)

tmp17854 := PrimCons(tmp17852, tmp17853)

tmp17855 := PrimCons(symstr, tmp17854)

tmp17856 := PrimCons(tmp17855, Nil)

tmp17857 := PrimCons(sym_1_1_6, tmp17856)

tmp17858 := PrimCons(tmp17850, tmp17857)

tmp17859 := Call(__e, PrimFunc(symdeclare), symshen_4tls, tmp17858)


_ = tmp17859

tmp17860 := PrimCons(symA, Nil)

tmp17861 := PrimCons(symlist, tmp17860)

tmp17862 := PrimCons(symB, Nil)

tmp17863 := PrimCons(tmp17861, tmp17862)

tmp17864 := PrimCons(symstr, tmp17863)

tmp17865 := PrimCons(symboolean, Nil)

tmp17866 := PrimCons(sym_1_1_6, tmp17865)

tmp17867 := PrimCons(tmp17864, tmp17866)

tmp17868 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure_2, tmp17867)


_ = tmp17868

tmp17869 := PrimCons(symA, Nil)

tmp17870 := PrimCons(symlist, tmp17869)

tmp17871 := PrimCons(symB, Nil)

tmp17872 := PrimCons(tmp17870, tmp17871)

tmp17873 := PrimCons(symstr, tmp17872)

tmp17874 := PrimCons(tmp17873, Nil)

tmp17875 := PrimCons(sym_1_1_6, tmp17874)

tmp17876 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure, tmp17875)


_ = tmp17876

tmp17877 := PrimCons(symA, Nil)

tmp17878 := PrimCons(symlist, tmp17877)

tmp17879 := PrimCons(symB, Nil)

tmp17880 := PrimCons(tmp17878, tmp17879)

tmp17881 := PrimCons(symstr, tmp17880)

tmp17882 := PrimCons(symB, Nil)

tmp17883 := PrimCons(sym_1_1_6, tmp17882)

tmp17884 := PrimCons(tmp17881, tmp17883)

tmp17885 := Call(__e, PrimFunc(symdeclare), symshen_4_5_1out, tmp17884)


_ = tmp17885

tmp17886 := PrimCons(symA, Nil)

tmp17887 := PrimCons(symlist, tmp17886)

tmp17888 := PrimCons(symB, Nil)

tmp17889 := PrimCons(tmp17887, tmp17888)

tmp17890 := PrimCons(symstr, tmp17889)

tmp17891 := PrimCons(symA, Nil)

tmp17892 := PrimCons(symlist, tmp17891)

tmp17893 := PrimCons(tmp17892, Nil)

tmp17894 := PrimCons(sym_1_1_6, tmp17893)

tmp17895 := PrimCons(tmp17890, tmp17894)

tmp17896 := Call(__e, PrimFunc(symdeclare), symshen_4in_1_6, tmp17895)


_ = tmp17896

tmp17897 := PrimCons(symA, Nil)

tmp17898 := PrimCons(symlist, tmp17897)

tmp17899 := PrimCons(symB, Nil)

tmp17900 := PrimCons(tmp17898, tmp17899)

tmp17901 := PrimCons(symstr, tmp17900)

tmp17902 := PrimCons(symboolean, Nil)

tmp17903 := PrimCons(sym_1_1_6, tmp17902)

tmp17904 := PrimCons(tmp17901, tmp17903)

tmp17905 := Call(__e, PrimFunc(symdeclare), symshen_4non_1empty_1stream_2, tmp17904)


_ = tmp17905

tmp17906 := PrimCons(symA, Nil)

tmp17907 := PrimCons(symlist, tmp17906)

tmp17908 := PrimCons(symA, Nil)

tmp17909 := PrimCons(symlist, tmp17908)

tmp17910 := PrimCons(symB, Nil)

tmp17911 := PrimCons(tmp17909, tmp17910)

tmp17912 := PrimCons(symstr, tmp17911)

tmp17913 := PrimCons(tmp17912, Nil)

tmp17914 := PrimCons(sym_1_1_6, tmp17913)

tmp17915 := PrimCons(symB, tmp17914)

tmp17916 := PrimCons(tmp17915, Nil)

tmp17917 := PrimCons(sym_1_1_6, tmp17916)

tmp17918 := PrimCons(tmp17907, tmp17917)

tmp17919 := Call(__e, PrimFunc(symdeclare), symshen_4comb, tmp17918)


_ = tmp17919

tmp17920 := PrimCons(symB, Nil)

tmp17921 := PrimCons(symA, tmp17920)

tmp17922 := PrimCons(symstr, tmp17921)

tmp17923 := PrimCons(symD, Nil)

tmp17924 := PrimCons(symC, tmp17923)

tmp17925 := PrimCons(symstr, tmp17924)

tmp17926 := PrimCons(symD, Nil)

tmp17927 := PrimCons(symC, tmp17926)

tmp17928 := PrimCons(symstr, tmp17927)

tmp17929 := PrimCons(tmp17928, Nil)

tmp17930 := PrimCons(symA, tmp17929)

tmp17931 := PrimCons(symstr, tmp17930)

tmp17932 := PrimCons(tmp17931, Nil)

tmp17933 := PrimCons(sym_1_1_6, tmp17932)

tmp17934 := PrimCons(tmp17925, tmp17933)

tmp17935 := PrimCons(tmp17934, Nil)

tmp17936 := PrimCons(sym_1_1_6, tmp17935)

tmp17937 := PrimCons(tmp17922, tmp17936)

tmp17938 := Call(__e, PrimFunc(symdeclare), symshen_4headstream, tmp17937)


_ = tmp17938

tmp17939 := PrimCons(symA, Nil)

tmp17940 := PrimCons(symlist, tmp17939)

tmp17941 := PrimCons(symB, Nil)

tmp17942 := PrimCons(tmp17940, tmp17941)

tmp17943 := PrimCons(symstr, tmp17942)

tmp17944 := PrimCons(symA, Nil)

tmp17945 := PrimCons(symlist, tmp17944)

tmp17946 := PrimCons(symB, Nil)

tmp17947 := PrimCons(tmp17945, tmp17946)

tmp17948 := PrimCons(symstr, tmp17947)

tmp17949 := PrimCons(tmp17948, Nil)

tmp17950 := PrimCons(sym_1_1_6, tmp17949)

tmp17951 := PrimCons(tmp17943, tmp17950)

tmp17952 := Call(__e, PrimFunc(symdeclare), symshen_4tlstream, tmp17951)


_ = tmp17952

tmp17953 := PrimCons(symA, Nil)

tmp17954 := PrimCons(symlist, tmp17953)

tmp17955 := PrimCons(symboolean, Nil)

tmp17956 := PrimCons(sym_1_1_6, tmp17955)

tmp17957 := PrimCons(tmp17954, tmp17956)

tmp17958 := PrimCons(tmp17957, Nil)

tmp17959 := PrimCons(sym_1_1_6, tmp17958)

tmp17960 := PrimCons(symA, tmp17959)

tmp17961 := Call(__e, PrimFunc(symdeclare), symelement_2, tmp17960)


_ = tmp17961

tmp17962 := PrimCons(symboolean, Nil)

tmp17963 := PrimCons(sym_1_1_6, tmp17962)

tmp17964 := PrimCons(symA, tmp17963)

tmp17965 := Call(__e, PrimFunc(symdeclare), symempty_2, tmp17964)


_ = tmp17965

tmp17966 := PrimCons(symboolean, Nil)

tmp17967 := PrimCons(sym_1_1_6, tmp17966)

tmp17968 := PrimCons(symsymbol, tmp17967)

tmp17969 := Call(__e, PrimFunc(symdeclare), symenable_1type_1theory, tmp17968)


_ = tmp17969

tmp17970 := PrimCons(symsymbol, Nil)

tmp17971 := PrimCons(symlist, tmp17970)

tmp17972 := PrimCons(tmp17971, Nil)

tmp17973 := PrimCons(sym_1_1_6, tmp17972)

tmp17974 := PrimCons(symsymbol, tmp17973)

tmp17975 := Call(__e, PrimFunc(symdeclare), symexternal, tmp17974)


_ = tmp17975

tmp17976 := PrimCons(symstring, Nil)

tmp17977 := PrimCons(sym_1_1_6, tmp17976)

tmp17978 := PrimCons(symexception, tmp17977)

tmp17979 := Call(__e, PrimFunc(symdeclare), symerror_1to_1string, tmp17978)


_ = tmp17979

tmp17980 := PrimCons(symstring, Nil)

tmp17981 := PrimCons(symlist, tmp17980)

tmp17982 := PrimCons(tmp17981, Nil)

tmp17983 := PrimCons(sym_1_1_6, tmp17982)

tmp17984 := PrimCons(symA, tmp17983)

tmp17985 := Call(__e, PrimFunc(symdeclare), symexplode, tmp17984)


_ = tmp17985

tmp17986 := PrimCons(symsymbol, Nil)

tmp17987 := PrimCons(sym_1_1_6, tmp17986)

tmp17988 := PrimCons(symsymbol, tmp17987)

tmp17989 := Call(__e, PrimFunc(symdeclare), symfactorise, tmp17988)


_ = tmp17989

tmp17990 := PrimCons(symsymbol, Nil)

tmp17991 := PrimCons(sym_1_1_6, tmp17990)

tmp17992 := Call(__e, PrimFunc(symdeclare), symfail, tmp17991)


_ = tmp17992

tmp17993 := PrimCons(symA, Nil)

tmp17994 := PrimCons(sym_1_1_6, tmp17993)

tmp17995 := PrimCons(symA, tmp17994)

tmp17996 := PrimCons(symA, Nil)

tmp17997 := PrimCons(sym_1_1_6, tmp17996)

tmp17998 := PrimCons(symA, tmp17997)

tmp17999 := PrimCons(tmp17998, Nil)

tmp18000 := PrimCons(sym_1_1_6, tmp17999)

tmp18001 := PrimCons(tmp17995, tmp18000)

tmp18002 := Call(__e, PrimFunc(symdeclare), symfix, tmp18001)


_ = tmp18002

tmp18003 := PrimCons(symA, Nil)

tmp18004 := PrimCons(symlazy, tmp18003)

tmp18005 := PrimCons(tmp18004, Nil)

tmp18006 := PrimCons(sym_1_1_6, tmp18005)

tmp18007 := PrimCons(symA, tmp18006)

tmp18008 := Call(__e, PrimFunc(symdeclare), symfreeze, tmp18007)


_ = tmp18008

tmp18009 := PrimCons(symB, Nil)

tmp18010 := PrimCons(sym_d, tmp18009)

tmp18011 := PrimCons(symA, tmp18010)

tmp18012 := PrimCons(symA, Nil)

tmp18013 := PrimCons(sym_1_1_6, tmp18012)

tmp18014 := PrimCons(tmp18011, tmp18013)

tmp18015 := Call(__e, PrimFunc(symdeclare), symfst, tmp18014)


_ = tmp18015

tmp18016 := PrimCons(symsymbol, Nil)

tmp18017 := PrimCons(sym_1_1_6, tmp18016)

tmp18018 := PrimCons(symsymbol, tmp18017)

tmp18019 := Call(__e, PrimFunc(symdeclare), symgensym, tmp18018)


_ = tmp18019

tmp18020 := PrimCons(symA, Nil)

tmp18021 := PrimCons(symvector, tmp18020)

tmp18022 := PrimCons(symA, Nil)

tmp18023 := PrimCons(sym_1_1_6, tmp18022)

tmp18024 := PrimCons(symnumber, tmp18023)

tmp18025 := PrimCons(tmp18024, Nil)

tmp18026 := PrimCons(sym_1_1_6, tmp18025)

tmp18027 := PrimCons(tmp18021, tmp18026)

tmp18028 := Call(__e, PrimFunc(symdeclare), sym_5_1vector, tmp18027)


_ = tmp18028

tmp18029 := PrimCons(symA, Nil)

tmp18030 := PrimCons(symvector, tmp18029)

tmp18031 := PrimCons(symA, Nil)

tmp18032 := PrimCons(symvector, tmp18031)

tmp18033 := PrimCons(tmp18032, Nil)

tmp18034 := PrimCons(sym_1_1_6, tmp18033)

tmp18035 := PrimCons(symA, tmp18034)

tmp18036 := PrimCons(tmp18035, Nil)

tmp18037 := PrimCons(sym_1_1_6, tmp18036)

tmp18038 := PrimCons(symnumber, tmp18037)

tmp18039 := PrimCons(tmp18038, Nil)

tmp18040 := PrimCons(sym_1_1_6, tmp18039)

tmp18041 := PrimCons(tmp18030, tmp18040)

tmp18042 := Call(__e, PrimFunc(symdeclare), symvector_1_6, tmp18041)


_ = tmp18042

tmp18043 := PrimCons(symA, Nil)

tmp18044 := PrimCons(symvector, tmp18043)

tmp18045 := PrimCons(tmp18044, Nil)

tmp18046 := PrimCons(sym_1_1_6, tmp18045)

tmp18047 := PrimCons(symnumber, tmp18046)

tmp18048 := Call(__e, PrimFunc(symdeclare), symvector, tmp18047)


_ = tmp18048

tmp18049 := PrimCons(symnumber, Nil)

tmp18050 := PrimCons(sym_1_1_6, tmp18049)

tmp18051 := PrimCons(symsymbol, tmp18050)

tmp18052 := Call(__e, PrimFunc(symdeclare), symget_1time, tmp18051)


_ = tmp18052

tmp18053 := PrimCons(symnumber, Nil)

tmp18054 := PrimCons(sym_1_1_6, tmp18053)

tmp18055 := PrimCons(symnumber, tmp18054)

tmp18056 := PrimCons(tmp18055, Nil)

tmp18057 := PrimCons(sym_1_1_6, tmp18056)

tmp18058 := PrimCons(symA, tmp18057)

tmp18059 := Call(__e, PrimFunc(symdeclare), symhash, tmp18058)


_ = tmp18059

tmp18060 := PrimCons(symA, Nil)

tmp18061 := PrimCons(symlist, tmp18060)

tmp18062 := PrimCons(symA, Nil)

tmp18063 := PrimCons(sym_1_1_6, tmp18062)

tmp18064 := PrimCons(tmp18061, tmp18063)

tmp18065 := Call(__e, PrimFunc(symdeclare), symhead, tmp18064)


_ = tmp18065

tmp18066 := PrimCons(symA, Nil)

tmp18067 := PrimCons(symvector, tmp18066)

tmp18068 := PrimCons(symA, Nil)

tmp18069 := PrimCons(sym_1_1_6, tmp18068)

tmp18070 := PrimCons(tmp18067, tmp18069)

tmp18071 := Call(__e, PrimFunc(symdeclare), symhdv, tmp18070)


_ = tmp18071

tmp18072 := PrimCons(symstring, Nil)

tmp18073 := PrimCons(sym_1_1_6, tmp18072)

tmp18074 := PrimCons(symstring, tmp18073)

tmp18075 := Call(__e, PrimFunc(symdeclare), symhdstr, tmp18074)


_ = tmp18075

tmp18076 := PrimCons(symA, Nil)

tmp18077 := PrimCons(sym_1_1_6, tmp18076)

tmp18078 := PrimCons(symA, tmp18077)

tmp18079 := PrimCons(tmp18078, Nil)

tmp18080 := PrimCons(sym_1_1_6, tmp18079)

tmp18081 := PrimCons(symA, tmp18080)

tmp18082 := PrimCons(tmp18081, Nil)

tmp18083 := PrimCons(sym_1_1_6, tmp18082)

tmp18084 := PrimCons(symboolean, tmp18083)

tmp18085 := Call(__e, PrimFunc(symdeclare), symif, tmp18084)


_ = tmp18085

tmp18086 := PrimCons(symsymbol, Nil)

tmp18087 := PrimCons(sym_1_1_6, tmp18086)

tmp18088 := PrimCons(symsymbol, tmp18087)

tmp18089 := Call(__e, PrimFunc(symdeclare), symin_1package, tmp18088)


_ = tmp18089

tmp18090 := PrimCons(symstring, Nil)

tmp18091 := PrimCons(sym_1_1_6, tmp18090)

tmp18092 := Call(__e, PrimFunc(symdeclare), symit, tmp18091)


_ = tmp18092

tmp18093 := PrimCons(symstring, Nil)

tmp18094 := PrimCons(sym_1_1_6, tmp18093)

tmp18095 := Call(__e, PrimFunc(symdeclare), symimplementation, tmp18094)


_ = tmp18095

tmp18096 := PrimCons(symsymbol, Nil)

tmp18097 := PrimCons(symlist, tmp18096)

tmp18098 := PrimCons(symsymbol, Nil)

tmp18099 := PrimCons(symlist, tmp18098)

tmp18100 := PrimCons(tmp18099, Nil)

tmp18101 := PrimCons(sym_1_1_6, tmp18100)

tmp18102 := PrimCons(tmp18097, tmp18101)

tmp18103 := Call(__e, PrimFunc(symdeclare), syminclude, tmp18102)


_ = tmp18103

tmp18104 := PrimCons(symsymbol, Nil)

tmp18105 := PrimCons(symlist, tmp18104)

tmp18106 := PrimCons(symsymbol, Nil)

tmp18107 := PrimCons(symlist, tmp18106)

tmp18108 := PrimCons(tmp18107, Nil)

tmp18109 := PrimCons(sym_1_1_6, tmp18108)

tmp18110 := PrimCons(tmp18105, tmp18109)

tmp18111 := Call(__e, PrimFunc(symdeclare), syminclude_1all_1but, tmp18110)


_ = tmp18111

tmp18112 := PrimCons(symnumber, Nil)

tmp18113 := PrimCons(sym_1_1_6, tmp18112)

tmp18114 := Call(__e, PrimFunc(symdeclare), syminferences, tmp18113)


_ = tmp18114

tmp18115 := PrimCons(symstring, Nil)

tmp18116 := PrimCons(sym_1_1_6, tmp18115)

tmp18117 := PrimCons(symstring, tmp18116)

tmp18118 := PrimCons(tmp18117, Nil)

tmp18119 := PrimCons(sym_1_1_6, tmp18118)

tmp18120 := PrimCons(symA, tmp18119)

tmp18121 := Call(__e, PrimFunc(symdeclare), symshen_4insert, tmp18120)


_ = tmp18121

tmp18122 := PrimCons(symboolean, Nil)

tmp18123 := PrimCons(sym_1_1_6, tmp18122)

tmp18124 := PrimCons(symA, tmp18123)

tmp18125 := Call(__e, PrimFunc(symdeclare), syminteger_2, tmp18124)


_ = tmp18125

tmp18126 := PrimCons(symsymbol, Nil)

tmp18127 := PrimCons(symlist, tmp18126)

tmp18128 := PrimCons(tmp18127, Nil)

tmp18129 := PrimCons(sym_1_1_6, tmp18128)

tmp18130 := PrimCons(symsymbol, tmp18129)

tmp18131 := Call(__e, PrimFunc(symdeclare), syminternal, tmp18130)


_ = tmp18131

tmp18132 := PrimCons(symA, Nil)

tmp18133 := PrimCons(symlist, tmp18132)

tmp18134 := PrimCons(symA, Nil)

tmp18135 := PrimCons(symlist, tmp18134)

tmp18136 := PrimCons(symA, Nil)

tmp18137 := PrimCons(symlist, tmp18136)

tmp18138 := PrimCons(tmp18137, Nil)

tmp18139 := PrimCons(sym_1_1_6, tmp18138)

tmp18140 := PrimCons(tmp18135, tmp18139)

tmp18141 := PrimCons(tmp18140, Nil)

tmp18142 := PrimCons(sym_1_1_6, tmp18141)

tmp18143 := PrimCons(tmp18133, tmp18142)

tmp18144 := Call(__e, PrimFunc(symdeclare), symintersection, tmp18143)


_ = tmp18144

tmp18145 := PrimCons(symstring, Nil)

tmp18146 := PrimCons(sym_1_1_6, tmp18145)

tmp18147 := Call(__e, PrimFunc(symdeclare), symlanguage, tmp18146)


_ = tmp18147

tmp18148 := PrimCons(symA, Nil)

tmp18149 := PrimCons(symlist, tmp18148)

tmp18150 := PrimCons(symnumber, Nil)

tmp18151 := PrimCons(sym_1_1_6, tmp18150)

tmp18152 := PrimCons(tmp18149, tmp18151)

tmp18153 := Call(__e, PrimFunc(symdeclare), symlength, tmp18152)


_ = tmp18153

tmp18154 := PrimCons(symA, Nil)

tmp18155 := PrimCons(symvector, tmp18154)

tmp18156 := PrimCons(symnumber, Nil)

tmp18157 := PrimCons(sym_1_1_6, tmp18156)

tmp18158 := PrimCons(tmp18155, tmp18157)

tmp18159 := Call(__e, PrimFunc(symdeclare), symlimit, tmp18158)


_ = tmp18159

tmp18160 := PrimCons(symin, Nil)

tmp18161 := PrimCons(symstream, tmp18160)

tmp18162 := PrimCons(symunit, Nil)

tmp18163 := PrimCons(symlist, tmp18162)

tmp18164 := PrimCons(tmp18163, Nil)

tmp18165 := PrimCons(sym_1_1_6, tmp18164)

tmp18166 := PrimCons(tmp18161, tmp18165)

tmp18167 := Call(__e, PrimFunc(symdeclare), symlineread, tmp18166)


_ = tmp18167

tmp18168 := PrimCons(symsymbol, Nil)

tmp18169 := PrimCons(sym_1_1_6, tmp18168)

tmp18170 := PrimCons(symstring, tmp18169)

tmp18171 := Call(__e, PrimFunc(symdeclare), symload, tmp18170)


_ = tmp18171

tmp18172 := PrimCons(symB, Nil)

tmp18173 := PrimCons(sym_1_1_6, tmp18172)

tmp18174 := PrimCons(symA, tmp18173)

tmp18175 := PrimCons(symA, Nil)

tmp18176 := PrimCons(symlist, tmp18175)

tmp18177 := PrimCons(symB, Nil)

tmp18178 := PrimCons(symlist, tmp18177)

tmp18179 := PrimCons(tmp18178, Nil)

tmp18180 := PrimCons(sym_1_1_6, tmp18179)

tmp18181 := PrimCons(tmp18176, tmp18180)

tmp18182 := PrimCons(tmp18181, Nil)

tmp18183 := PrimCons(sym_1_1_6, tmp18182)

tmp18184 := PrimCons(tmp18174, tmp18183)

tmp18185 := Call(__e, PrimFunc(symdeclare), symmap, tmp18184)


_ = tmp18185

tmp18186 := PrimCons(symB, Nil)

tmp18187 := PrimCons(symlist, tmp18186)

tmp18188 := PrimCons(tmp18187, Nil)

tmp18189 := PrimCons(sym_1_1_6, tmp18188)

tmp18190 := PrimCons(symA, tmp18189)

tmp18191 := PrimCons(symA, Nil)

tmp18192 := PrimCons(symlist, tmp18191)

tmp18193 := PrimCons(symB, Nil)

tmp18194 := PrimCons(symlist, tmp18193)

tmp18195 := PrimCons(tmp18194, Nil)

tmp18196 := PrimCons(sym_1_1_6, tmp18195)

tmp18197 := PrimCons(tmp18192, tmp18196)

tmp18198 := PrimCons(tmp18197, Nil)

tmp18199 := PrimCons(sym_1_1_6, tmp18198)

tmp18200 := PrimCons(tmp18190, tmp18199)

tmp18201 := Call(__e, PrimFunc(symdeclare), symmapcan, tmp18200)


_ = tmp18201

tmp18202 := PrimCons(symnumber, Nil)

tmp18203 := PrimCons(sym_1_1_6, tmp18202)

tmp18204 := PrimCons(symnumber, tmp18203)

tmp18205 := Call(__e, PrimFunc(symdeclare), symmaxinferences, tmp18204)


_ = tmp18205

tmp18206 := PrimCons(symstring, Nil)

tmp18207 := PrimCons(sym_1_1_6, tmp18206)

tmp18208 := PrimCons(symnumber, tmp18207)

tmp18209 := Call(__e, PrimFunc(symdeclare), symn_1_6string, tmp18208)


_ = tmp18209

tmp18210 := PrimCons(symnumber, Nil)

tmp18211 := PrimCons(sym_1_1_6, tmp18210)

tmp18212 := PrimCons(symnumber, tmp18211)

tmp18213 := Call(__e, PrimFunc(symdeclare), symnl, tmp18212)


_ = tmp18213

tmp18214 := PrimCons(symboolean, Nil)

tmp18215 := PrimCons(sym_1_1_6, tmp18214)

tmp18216 := PrimCons(symboolean, tmp18215)

tmp18217 := Call(__e, PrimFunc(symdeclare), symnot, tmp18216)


_ = tmp18217

tmp18218 := PrimCons(symA, Nil)

tmp18219 := PrimCons(symlist, tmp18218)

tmp18220 := PrimCons(symA, Nil)

tmp18221 := PrimCons(sym_1_1_6, tmp18220)

tmp18222 := PrimCons(tmp18219, tmp18221)

tmp18223 := PrimCons(tmp18222, Nil)

tmp18224 := PrimCons(sym_1_1_6, tmp18223)

tmp18225 := PrimCons(symnumber, tmp18224)

tmp18226 := Call(__e, PrimFunc(symdeclare), symnth, tmp18225)


_ = tmp18226

tmp18227 := PrimCons(symboolean, Nil)

tmp18228 := PrimCons(sym_1_1_6, tmp18227)

tmp18229 := PrimCons(symA, tmp18228)

tmp18230 := Call(__e, PrimFunc(symdeclare), symnumber_2, tmp18229)


_ = tmp18230

tmp18231 := PrimCons(symnumber, Nil)

tmp18232 := PrimCons(sym_1_1_6, tmp18231)

tmp18233 := PrimCons(symB, tmp18232)

tmp18234 := PrimCons(tmp18233, Nil)

tmp18235 := PrimCons(sym_1_1_6, tmp18234)

tmp18236 := PrimCons(symA, tmp18235)

tmp18237 := Call(__e, PrimFunc(symdeclare), symoccurrences, tmp18236)


_ = tmp18237

tmp18238 := PrimCons(symboolean, Nil)

tmp18239 := PrimCons(sym_1_1_6, tmp18238)

tmp18240 := PrimCons(symsymbol, tmp18239)

tmp18241 := Call(__e, PrimFunc(symdeclare), symoccurs_1check, tmp18240)


_ = tmp18241

tmp18242 := PrimCons(symboolean, Nil)

tmp18243 := PrimCons(sym_1_1_6, tmp18242)

tmp18244 := PrimCons(symsymbol, tmp18243)

tmp18245 := Call(__e, PrimFunc(symdeclare), symoptimise, tmp18244)


_ = tmp18245

tmp18246 := PrimCons(symboolean, Nil)

tmp18247 := PrimCons(sym_1_1_6, tmp18246)

tmp18248 := PrimCons(symboolean, tmp18247)

tmp18249 := PrimCons(tmp18248, Nil)

tmp18250 := PrimCons(sym_1_1_6, tmp18249)

tmp18251 := PrimCons(symboolean, tmp18250)

tmp18252 := Call(__e, PrimFunc(symdeclare), symor, tmp18251)


_ = tmp18252

tmp18253 := PrimCons(symstring, Nil)

tmp18254 := PrimCons(sym_1_1_6, tmp18253)

tmp18255 := Call(__e, PrimFunc(symdeclare), symos, tmp18254)


_ = tmp18255

tmp18256 := PrimCons(symboolean, Nil)

tmp18257 := PrimCons(sym_1_1_6, tmp18256)

tmp18258 := PrimCons(symsymbol, tmp18257)

tmp18259 := Call(__e, PrimFunc(symdeclare), sympackage_2, tmp18258)


_ = tmp18259

tmp18260 := PrimCons(symstring, Nil)

tmp18261 := PrimCons(sym_1_1_6, tmp18260)

tmp18262 := Call(__e, PrimFunc(symdeclare), symport, tmp18261)


_ = tmp18262

tmp18263 := PrimCons(symstring, Nil)

tmp18264 := PrimCons(sym_1_1_6, tmp18263)

tmp18265 := Call(__e, PrimFunc(symdeclare), symporters, tmp18264)


_ = tmp18265

tmp18266 := PrimCons(symstring, Nil)

tmp18267 := PrimCons(sym_1_1_6, tmp18266)

tmp18268 := PrimCons(symnumber, tmp18267)

tmp18269 := PrimCons(tmp18268, Nil)

tmp18270 := PrimCons(sym_1_1_6, tmp18269)

tmp18271 := PrimCons(symstring, tmp18270)

tmp18272 := Call(__e, PrimFunc(symdeclare), sympos, tmp18271)


_ = tmp18272

tmp18273 := PrimCons(symout, Nil)

tmp18274 := PrimCons(symstream, tmp18273)

tmp18275 := PrimCons(symstring, Nil)

tmp18276 := PrimCons(sym_1_1_6, tmp18275)

tmp18277 := PrimCons(tmp18274, tmp18276)

tmp18278 := PrimCons(tmp18277, Nil)

tmp18279 := PrimCons(sym_1_1_6, tmp18278)

tmp18280 := PrimCons(symstring, tmp18279)

tmp18281 := Call(__e, PrimFunc(symdeclare), sympr, tmp18280)


_ = tmp18281

tmp18282 := PrimCons(symA, Nil)

tmp18283 := PrimCons(sym_1_1_6, tmp18282)

tmp18284 := PrimCons(symA, tmp18283)

tmp18285 := Call(__e, PrimFunc(symdeclare), symprint, tmp18284)


_ = tmp18285

tmp18286 := PrimCons(symsymbol, Nil)

tmp18287 := PrimCons(sym_1_1_6, tmp18286)

tmp18288 := PrimCons(symsymbol, tmp18287)

tmp18289 := Call(__e, PrimFunc(symdeclare), symprofile, tmp18288)


_ = tmp18289

tmp18290 := PrimCons(symsymbol, Nil)

tmp18291 := PrimCons(symlist, tmp18290)

tmp18292 := PrimCons(symsymbol, Nil)

tmp18293 := PrimCons(symlist, tmp18292)

tmp18294 := PrimCons(tmp18293, Nil)

tmp18295 := PrimCons(sym_1_1_6, tmp18294)

tmp18296 := PrimCons(tmp18291, tmp18295)

tmp18297 := Call(__e, PrimFunc(symdeclare), sympreclude, tmp18296)


_ = tmp18297

tmp18298 := PrimCons(symstring, Nil)

tmp18299 := PrimCons(sym_1_1_6, tmp18298)

tmp18300 := PrimCons(symstring, tmp18299)

tmp18301 := Call(__e, PrimFunc(symdeclare), symshen_4proc_1nl, tmp18300)


_ = tmp18301

tmp18302 := PrimCons(symnumber, Nil)

tmp18303 := PrimCons(sym_d, tmp18302)

tmp18304 := PrimCons(symsymbol, tmp18303)

tmp18305 := PrimCons(tmp18304, Nil)

tmp18306 := PrimCons(sym_1_1_6, tmp18305)

tmp18307 := PrimCons(symsymbol, tmp18306)

tmp18308 := Call(__e, PrimFunc(symdeclare), symprofile_1results, tmp18307)


_ = tmp18308

tmp18309 := PrimCons(symA, Nil)

tmp18310 := PrimCons(sym_1_1_6, tmp18309)

tmp18311 := PrimCons(symA, tmp18310)

tmp18312 := Call(__e, PrimFunc(symdeclare), symprotect, tmp18311)


_ = tmp18312

tmp18313 := PrimCons(symsymbol, Nil)

tmp18314 := PrimCons(symlist, tmp18313)

tmp18315 := PrimCons(symsymbol, Nil)

tmp18316 := PrimCons(symlist, tmp18315)

tmp18317 := PrimCons(tmp18316, Nil)

tmp18318 := PrimCons(sym_1_1_6, tmp18317)

tmp18319 := PrimCons(tmp18314, tmp18318)

tmp18320 := Call(__e, PrimFunc(symdeclare), sympreclude_1all_1but, tmp18319)


_ = tmp18320

tmp18321 := PrimCons(symout, Nil)

tmp18322 := PrimCons(symstream, tmp18321)

tmp18323 := PrimCons(symstring, Nil)

tmp18324 := PrimCons(sym_1_1_6, tmp18323)

tmp18325 := PrimCons(tmp18322, tmp18324)

tmp18326 := PrimCons(tmp18325, Nil)

tmp18327 := PrimCons(sym_1_1_6, tmp18326)

tmp18328 := PrimCons(symstring, tmp18327)

tmp18329 := Call(__e, PrimFunc(symdeclare), symshen_4prhush, tmp18328)


_ = tmp18329

tmp18330 := PrimCons(symnumber, Nil)

tmp18331 := PrimCons(sym_1_1_6, tmp18330)

tmp18332 := PrimCons(symnumber, tmp18331)

tmp18333 := Call(__e, PrimFunc(symdeclare), symprolog_1memory, tmp18332)


_ = tmp18333

tmp18334 := PrimCons(symunit, Nil)

tmp18335 := PrimCons(symlist, tmp18334)

tmp18336 := PrimCons(tmp18335, Nil)

tmp18337 := PrimCons(sym_1_1_6, tmp18336)

tmp18338 := PrimCons(symsymbol, tmp18337)

tmp18339 := Call(__e, PrimFunc(symdeclare), symps, tmp18338)


_ = tmp18339

tmp18340 := PrimCons(symin, Nil)

tmp18341 := PrimCons(symstream, tmp18340)

tmp18342 := PrimCons(symunit, Nil)

tmp18343 := PrimCons(sym_1_1_6, tmp18342)

tmp18344 := PrimCons(tmp18341, tmp18343)

tmp18345 := Call(__e, PrimFunc(symdeclare), symread, tmp18344)


_ = tmp18345

tmp18346 := PrimCons(symin, Nil)

tmp18347 := PrimCons(symstream, tmp18346)

tmp18348 := PrimCons(symnumber, Nil)

tmp18349 := PrimCons(sym_1_1_6, tmp18348)

tmp18350 := PrimCons(tmp18347, tmp18349)

tmp18351 := Call(__e, PrimFunc(symdeclare), symread_1byte, tmp18350)


_ = tmp18351

tmp18352 := PrimCons(symnumber, Nil)

tmp18353 := PrimCons(symlist, tmp18352)

tmp18354 := PrimCons(tmp18353, Nil)

tmp18355 := PrimCons(sym_1_1_6, tmp18354)

tmp18356 := PrimCons(symstring, tmp18355)

tmp18357 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1bytelist, tmp18356)


_ = tmp18357

tmp18358 := PrimCons(symstring, Nil)

tmp18359 := PrimCons(sym_1_1_6, tmp18358)

tmp18360 := PrimCons(symstring, tmp18359)

tmp18361 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1string, tmp18360)


_ = tmp18361

tmp18362 := PrimCons(symunit, Nil)

tmp18363 := PrimCons(symlist, tmp18362)

tmp18364 := PrimCons(tmp18363, Nil)

tmp18365 := PrimCons(sym_1_1_6, tmp18364)

tmp18366 := PrimCons(symstring, tmp18365)

tmp18367 := Call(__e, PrimFunc(symdeclare), symread_1file, tmp18366)


_ = tmp18367

tmp18368 := PrimCons(symunit, Nil)

tmp18369 := PrimCons(symlist, tmp18368)

tmp18370 := PrimCons(tmp18369, Nil)

tmp18371 := PrimCons(sym_1_1_6, tmp18370)

tmp18372 := PrimCons(symstring, tmp18371)

tmp18373 := Call(__e, PrimFunc(symdeclare), symread_1from_1string, tmp18372)


_ = tmp18373

tmp18374 := PrimCons(symunit, Nil)

tmp18375 := PrimCons(symlist, tmp18374)

tmp18376 := PrimCons(tmp18375, Nil)

tmp18377 := PrimCons(sym_1_1_6, tmp18376)

tmp18378 := PrimCons(symstring, tmp18377)

tmp18379 := Call(__e, PrimFunc(symdeclare), symread_1from_1string_1unprocessed, tmp18378)


_ = tmp18379

tmp18380 := PrimCons(symstring, Nil)

tmp18381 := PrimCons(sym_1_1_6, tmp18380)

tmp18382 := Call(__e, PrimFunc(symdeclare), symrelease, tmp18381)


_ = tmp18382

tmp18383 := PrimCons(symA, Nil)

tmp18384 := PrimCons(symlist, tmp18383)

tmp18385 := PrimCons(symA, Nil)

tmp18386 := PrimCons(symlist, tmp18385)

tmp18387 := PrimCons(tmp18386, Nil)

tmp18388 := PrimCons(sym_1_1_6, tmp18387)

tmp18389 := PrimCons(tmp18384, tmp18388)

tmp18390 := PrimCons(tmp18389, Nil)

tmp18391 := PrimCons(sym_1_1_6, tmp18390)

tmp18392 := PrimCons(symA, tmp18391)

tmp18393 := Call(__e, PrimFunc(symdeclare), symremove, tmp18392)


_ = tmp18393

tmp18394 := PrimCons(symA, Nil)

tmp18395 := PrimCons(symlist, tmp18394)

tmp18396 := PrimCons(symA, Nil)

tmp18397 := PrimCons(symlist, tmp18396)

tmp18398 := PrimCons(tmp18397, Nil)

tmp18399 := PrimCons(sym_1_1_6, tmp18398)

tmp18400 := PrimCons(tmp18395, tmp18399)

tmp18401 := Call(__e, PrimFunc(symdeclare), symreverse, tmp18400)


_ = tmp18401

tmp18402 := PrimCons(symA, Nil)

tmp18403 := PrimCons(sym_1_1_6, tmp18402)

tmp18404 := PrimCons(symstring, tmp18403)

tmp18405 := Call(__e, PrimFunc(symdeclare), symsimple_1error, tmp18404)


_ = tmp18405

tmp18406 := PrimCons(symB, Nil)

tmp18407 := PrimCons(sym_d, tmp18406)

tmp18408 := PrimCons(symA, tmp18407)

tmp18409 := PrimCons(symB, Nil)

tmp18410 := PrimCons(sym_1_1_6, tmp18409)

tmp18411 := PrimCons(tmp18408, tmp18410)

tmp18412 := Call(__e, PrimFunc(symdeclare), symsnd, tmp18411)


_ = tmp18412

tmp18413 := PrimCons(symsymbol, Nil)

tmp18414 := PrimCons(sym_1_1_6, tmp18413)

tmp18415 := PrimCons(symnumber, tmp18414)

tmp18416 := PrimCons(tmp18415, Nil)

tmp18417 := PrimCons(sym_1_1_6, tmp18416)

tmp18418 := PrimCons(symsymbol, tmp18417)

tmp18419 := Call(__e, PrimFunc(symdeclare), symspecialise, tmp18418)


_ = tmp18419

tmp18420 := PrimCons(symboolean, Nil)

tmp18421 := PrimCons(sym_1_1_6, tmp18420)

tmp18422 := PrimCons(symsymbol, tmp18421)

tmp18423 := Call(__e, PrimFunc(symdeclare), symspy, tmp18422)


_ = tmp18423

tmp18424 := PrimCons(symboolean, Nil)

tmp18425 := PrimCons(sym_1_1_6, tmp18424)

tmp18426 := PrimCons(symsymbol, tmp18425)

tmp18427 := Call(__e, PrimFunc(symdeclare), symstep, tmp18426)


_ = tmp18427

tmp18428 := PrimCons(symin, Nil)

tmp18429 := PrimCons(symstream, tmp18428)

tmp18430 := PrimCons(tmp18429, Nil)

tmp18431 := PrimCons(sym_1_1_6, tmp18430)

tmp18432 := Call(__e, PrimFunc(symdeclare), symstinput, tmp18431)


_ = tmp18432

tmp18433 := PrimCons(symout, Nil)

tmp18434 := PrimCons(symstream, tmp18433)

tmp18435 := PrimCons(tmp18434, Nil)

tmp18436 := PrimCons(sym_1_1_6, tmp18435)

tmp18437 := Call(__e, PrimFunc(symdeclare), symstoutput, tmp18436)


_ = tmp18437

tmp18438 := PrimCons(symboolean, Nil)

tmp18439 := PrimCons(sym_1_1_6, tmp18438)

tmp18440 := PrimCons(symA, tmp18439)

tmp18441 := Call(__e, PrimFunc(symdeclare), symstring_2, tmp18440)


_ = tmp18441

tmp18442 := PrimCons(symstring, Nil)

tmp18443 := PrimCons(sym_1_1_6, tmp18442)

tmp18444 := PrimCons(symA, tmp18443)

tmp18445 := Call(__e, PrimFunc(symdeclare), symstr, tmp18444)


_ = tmp18445

tmp18446 := PrimCons(symnumber, Nil)

tmp18447 := PrimCons(sym_1_1_6, tmp18446)

tmp18448 := PrimCons(symstring, tmp18447)

tmp18449 := Call(__e, PrimFunc(symdeclare), symstring_1_6n, tmp18448)


_ = tmp18449

tmp18450 := PrimCons(symsymbol, Nil)

tmp18451 := PrimCons(sym_1_1_6, tmp18450)

tmp18452 := PrimCons(symstring, tmp18451)

tmp18453 := Call(__e, PrimFunc(symdeclare), symstring_1_6symbol, tmp18452)


_ = tmp18453

tmp18454 := PrimCons(symnumber, Nil)

tmp18455 := PrimCons(symlist, tmp18454)

tmp18456 := PrimCons(symnumber, Nil)

tmp18457 := PrimCons(sym_1_1_6, tmp18456)

tmp18458 := PrimCons(tmp18455, tmp18457)

tmp18459 := Call(__e, PrimFunc(symdeclare), symsum, tmp18458)


_ = tmp18459

tmp18460 := PrimCons(symboolean, Nil)

tmp18461 := PrimCons(sym_1_1_6, tmp18460)

tmp18462 := PrimCons(symA, tmp18461)

tmp18463 := Call(__e, PrimFunc(symdeclare), symsymbol_2, tmp18462)


_ = tmp18463

tmp18464 := PrimCons(symsymbol, Nil)

tmp18465 := PrimCons(sym_1_1_6, tmp18464)

tmp18466 := PrimCons(symsymbol, tmp18465)

tmp18467 := Call(__e, PrimFunc(symdeclare), symsystemf, tmp18466)


_ = tmp18467

tmp18468 := PrimCons(symA, Nil)

tmp18469 := PrimCons(symlist, tmp18468)

tmp18470 := PrimCons(symA, Nil)

tmp18471 := PrimCons(symlist, tmp18470)

tmp18472 := PrimCons(tmp18471, Nil)

tmp18473 := PrimCons(sym_1_1_6, tmp18472)

tmp18474 := PrimCons(tmp18469, tmp18473)

tmp18475 := Call(__e, PrimFunc(symdeclare), symtail, tmp18474)


_ = tmp18475

tmp18476 := PrimCons(symstring, Nil)

tmp18477 := PrimCons(sym_1_1_6, tmp18476)

tmp18478 := PrimCons(symstring, tmp18477)

tmp18479 := Call(__e, PrimFunc(symdeclare), symtlstr, tmp18478)


_ = tmp18479

tmp18480 := PrimCons(symA, Nil)

tmp18481 := PrimCons(symvector, tmp18480)

tmp18482 := PrimCons(symA, Nil)

tmp18483 := PrimCons(symvector, tmp18482)

tmp18484 := PrimCons(tmp18483, Nil)

tmp18485 := PrimCons(sym_1_1_6, tmp18484)

tmp18486 := PrimCons(tmp18481, tmp18485)

tmp18487 := Call(__e, PrimFunc(symdeclare), symtlv, tmp18486)


_ = tmp18487

tmp18488 := PrimCons(symboolean, Nil)

tmp18489 := PrimCons(sym_1_1_6, tmp18488)

tmp18490 := PrimCons(symsymbol, tmp18489)

tmp18491 := Call(__e, PrimFunc(symdeclare), symtc, tmp18490)


_ = tmp18491

tmp18492 := PrimCons(symboolean, Nil)

tmp18493 := PrimCons(sym_1_1_6, tmp18492)

tmp18494 := Call(__e, PrimFunc(symdeclare), symtc_2, tmp18493)


_ = tmp18494

tmp18495 := PrimCons(symA, Nil)

tmp18496 := PrimCons(symlazy, tmp18495)

tmp18497 := PrimCons(symA, Nil)

tmp18498 := PrimCons(sym_1_1_6, tmp18497)

tmp18499 := PrimCons(tmp18496, tmp18498)

tmp18500 := Call(__e, PrimFunc(symdeclare), symthaw, tmp18499)


_ = tmp18500

tmp18501 := PrimCons(symsymbol, Nil)

tmp18502 := PrimCons(sym_1_1_6, tmp18501)

tmp18503 := PrimCons(symsymbol, tmp18502)

tmp18504 := Call(__e, PrimFunc(symdeclare), symtrack, tmp18503)


_ = tmp18504

tmp18505 := PrimCons(symA, Nil)

tmp18506 := PrimCons(sym_1_1_6, tmp18505)

tmp18507 := PrimCons(symexception, tmp18506)

tmp18508 := PrimCons(symA, Nil)

tmp18509 := PrimCons(sym_1_1_6, tmp18508)

tmp18510 := PrimCons(tmp18507, tmp18509)

tmp18511 := PrimCons(tmp18510, Nil)

tmp18512 := PrimCons(sym_1_1_6, tmp18511)

tmp18513 := PrimCons(symA, tmp18512)

tmp18514 := Call(__e, PrimFunc(symdeclare), symtrap_1error, tmp18513)


_ = tmp18514

tmp18515 := PrimCons(symboolean, Nil)

tmp18516 := PrimCons(sym_1_1_6, tmp18515)

tmp18517 := PrimCons(symA, tmp18516)

tmp18518 := Call(__e, PrimFunc(symdeclare), symtuple_2, tmp18517)


_ = tmp18518

tmp18519 := PrimCons(symsymbol, Nil)

tmp18520 := PrimCons(sym_1_1_6, tmp18519)

tmp18521 := PrimCons(symsymbol, tmp18520)

tmp18522 := Call(__e, PrimFunc(symdeclare), symundefmacro, tmp18521)


_ = tmp18522

tmp18523 := PrimCons(symA, Nil)

tmp18524 := PrimCons(symlist, tmp18523)

tmp18525 := PrimCons(symA, Nil)

tmp18526 := PrimCons(symlist, tmp18525)

tmp18527 := PrimCons(symA, Nil)

tmp18528 := PrimCons(symlist, tmp18527)

tmp18529 := PrimCons(tmp18528, Nil)

tmp18530 := PrimCons(sym_1_1_6, tmp18529)

tmp18531 := PrimCons(tmp18526, tmp18530)

tmp18532 := PrimCons(tmp18531, Nil)

tmp18533 := PrimCons(sym_1_1_6, tmp18532)

tmp18534 := PrimCons(tmp18524, tmp18533)

tmp18535 := Call(__e, PrimFunc(symdeclare), symunion, tmp18534)


_ = tmp18535

tmp18536 := PrimCons(symsymbol, Nil)

tmp18537 := PrimCons(sym_1_1_6, tmp18536)

tmp18538 := PrimCons(symsymbol, tmp18537)

tmp18539 := Call(__e, PrimFunc(symdeclare), symunprofile, tmp18538)


_ = tmp18539

tmp18540 := PrimCons(symsymbol, Nil)

tmp18541 := PrimCons(sym_1_1_6, tmp18540)

tmp18542 := PrimCons(symsymbol, tmp18541)

tmp18543 := Call(__e, PrimFunc(symdeclare), symuntrack, tmp18542)


_ = tmp18543

tmp18544 := PrimCons(symboolean, Nil)

tmp18545 := PrimCons(sym_1_1_6, tmp18544)

tmp18546 := PrimCons(symA, tmp18545)

tmp18547 := Call(__e, PrimFunc(symdeclare), symvariable_2, tmp18546)


_ = tmp18547

tmp18548 := PrimCons(symboolean, Nil)

tmp18549 := PrimCons(sym_1_1_6, tmp18548)

tmp18550 := PrimCons(symA, tmp18549)

tmp18551 := Call(__e, PrimFunc(symdeclare), symvector_2, tmp18550)


_ = tmp18551

tmp18552 := PrimCons(symstring, Nil)

tmp18553 := PrimCons(sym_1_1_6, tmp18552)

tmp18554 := Call(__e, PrimFunc(symdeclare), symversion, tmp18553)


_ = tmp18554

tmp18555 := PrimCons(symA, Nil)

tmp18556 := PrimCons(sym_1_1_6, tmp18555)

tmp18557 := PrimCons(symA, tmp18556)

tmp18558 := PrimCons(tmp18557, Nil)

tmp18559 := PrimCons(sym_1_1_6, tmp18558)

tmp18560 := PrimCons(symstring, tmp18559)

tmp18561 := Call(__e, PrimFunc(symdeclare), symwrite_1to_1file, tmp18560)


_ = tmp18561

tmp18562 := PrimCons(symout, Nil)

tmp18563 := PrimCons(symstream, tmp18562)

tmp18564 := PrimCons(symnumber, Nil)

tmp18565 := PrimCons(sym_1_1_6, tmp18564)

tmp18566 := PrimCons(tmp18563, tmp18565)

tmp18567 := PrimCons(tmp18566, Nil)

tmp18568 := PrimCons(sym_1_1_6, tmp18567)

tmp18569 := PrimCons(symnumber, tmp18568)

tmp18570 := Call(__e, PrimFunc(symdeclare), symwrite_1byte, tmp18569)


_ = tmp18570

tmp18571 := PrimCons(symboolean, Nil)

tmp18572 := PrimCons(sym_1_1_6, tmp18571)

tmp18573 := PrimCons(symstring, tmp18572)

tmp18574 := Call(__e, PrimFunc(symdeclare), symy_1or_1n_2, tmp18573)


_ = tmp18574

tmp18575 := PrimCons(symboolean, Nil)

tmp18576 := PrimCons(sym_1_1_6, tmp18575)

tmp18577 := PrimCons(symnumber, tmp18576)

tmp18578 := PrimCons(tmp18577, Nil)

tmp18579 := PrimCons(sym_1_1_6, tmp18578)

tmp18580 := PrimCons(symnumber, tmp18579)

tmp18581 := Call(__e, PrimFunc(symdeclare), sym_6, tmp18580)


_ = tmp18581

tmp18582 := PrimCons(symboolean, Nil)

tmp18583 := PrimCons(sym_1_1_6, tmp18582)

tmp18584 := PrimCons(symnumber, tmp18583)

tmp18585 := PrimCons(tmp18584, Nil)

tmp18586 := PrimCons(sym_1_1_6, tmp18585)

tmp18587 := PrimCons(symnumber, tmp18586)

tmp18588 := Call(__e, PrimFunc(symdeclare), sym_5, tmp18587)


_ = tmp18588

tmp18589 := PrimCons(symboolean, Nil)

tmp18590 := PrimCons(sym_1_1_6, tmp18589)

tmp18591 := PrimCons(symnumber, tmp18590)

tmp18592 := PrimCons(tmp18591, Nil)

tmp18593 := PrimCons(sym_1_1_6, tmp18592)

tmp18594 := PrimCons(symnumber, tmp18593)

tmp18595 := Call(__e, PrimFunc(symdeclare), sym_6_a, tmp18594)


_ = tmp18595

tmp18596 := PrimCons(symboolean, Nil)

tmp18597 := PrimCons(sym_1_1_6, tmp18596)

tmp18598 := PrimCons(symnumber, tmp18597)

tmp18599 := PrimCons(tmp18598, Nil)

tmp18600 := PrimCons(sym_1_1_6, tmp18599)

tmp18601 := PrimCons(symnumber, tmp18600)

tmp18602 := Call(__e, PrimFunc(symdeclare), sym_5_a, tmp18601)


_ = tmp18602

tmp18603 := PrimCons(symboolean, Nil)

tmp18604 := PrimCons(sym_1_1_6, tmp18603)

tmp18605 := PrimCons(symA, tmp18604)

tmp18606 := PrimCons(tmp18605, Nil)

tmp18607 := PrimCons(sym_1_1_6, tmp18606)

tmp18608 := PrimCons(symA, tmp18607)

tmp18609 := Call(__e, PrimFunc(symdeclare), sym_a, tmp18608)


_ = tmp18609

tmp18610 := PrimCons(symnumber, Nil)

tmp18611 := PrimCons(sym_1_1_6, tmp18610)

tmp18612 := PrimCons(symnumber, tmp18611)

tmp18613 := PrimCons(tmp18612, Nil)

tmp18614 := PrimCons(sym_1_1_6, tmp18613)

tmp18615 := PrimCons(symnumber, tmp18614)

tmp18616 := Call(__e, PrimFunc(symdeclare), sym_7, tmp18615)


_ = tmp18616

tmp18617 := PrimCons(symnumber, Nil)

tmp18618 := PrimCons(sym_1_1_6, tmp18617)

tmp18619 := PrimCons(symnumber, tmp18618)

tmp18620 := PrimCons(tmp18619, Nil)

tmp18621 := PrimCons(sym_1_1_6, tmp18620)

tmp18622 := PrimCons(symnumber, tmp18621)

tmp18623 := Call(__e, PrimFunc(symdeclare), sym_c, tmp18622)


_ = tmp18623

tmp18624 := PrimCons(symnumber, Nil)

tmp18625 := PrimCons(sym_1_1_6, tmp18624)

tmp18626 := PrimCons(symnumber, tmp18625)

tmp18627 := PrimCons(tmp18626, Nil)

tmp18628 := PrimCons(sym_1_1_6, tmp18627)

tmp18629 := PrimCons(symnumber, tmp18628)

tmp18630 := Call(__e, PrimFunc(symdeclare), sym_1, tmp18629)


_ = tmp18630

tmp18631 := PrimCons(symnumber, Nil)

tmp18632 := PrimCons(sym_1_1_6, tmp18631)

tmp18633 := PrimCons(symnumber, tmp18632)

tmp18634 := PrimCons(tmp18633, Nil)

tmp18635 := PrimCons(sym_1_1_6, tmp18634)

tmp18636 := PrimCons(symnumber, tmp18635)

tmp18637 := Call(__e, PrimFunc(symdeclare), sym_d, tmp18636)


_ = tmp18637

tmp18638 := PrimCons(symboolean, Nil)

tmp18639 := PrimCons(sym_1_1_6, tmp18638)

tmp18640 := PrimCons(symB, tmp18639)

tmp18641 := PrimCons(tmp18640, Nil)

tmp18642 := PrimCons(sym_1_1_6, tmp18641)

tmp18643 := PrimCons(symA, tmp18642)

__e.TailApply(PrimFunc(symdeclare), sym_a_a, tmp18643)
return




}, 0)

var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
var symshen_4eos = MakeSymbol("shen.eos")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symprotect = MakeSymbol("protect")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4non_1empty_1stream_2 = MakeSymbol("shen.non-empty-stream?")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var symshen_4let_1macro = MakeSymbol("shen.let-macro")
var symdefprolog = MakeSymbol("defprolog")
var symL = MakeSymbol("L")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4compile_1premises = MakeSymbol("shen.compile-premises")
var symconcat = MakeSymbol("concat")
var symdifference = MakeSymbol("difference")
var symlanguage = MakeSymbol("language")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symuntrack = MakeSymbol("untrack")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var sym_a = MakeSymbol("=")
var symshen_4defprolog_1macro = MakeSymbol("shen.defprolog-macro")
var symunprofile = MakeSymbol("unprofile")
var symshen_4pprint = MakeSymbol("shen.pprint")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symS = MakeSymbol("S")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symclose = MakeSymbol("close")
var sym_6_6 = MakeSymbol(">>")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symeval_1kl = MakeSymbol("eval-kl")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symverified = MakeSymbol("verified")
var symshen_4goals = MakeSymbol("shen.goals")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symabort = MakeSymbol("abort")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symtrack = MakeSymbol("track")
var symfindall = MakeSymbol("findall")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symhdv = MakeSymbol("hdv")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4split_1cases = MakeSymbol("shen.split-cases")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4assoc_1macro = MakeSymbol("shen.assoc-macro")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var symput = MakeSymbol("put")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symMessage = MakeSymbol("Message")
var symGoTo = MakeSymbol("GoTo")
var symshen_4goto = MakeSymbol("shen.goto")
var symshen_4return_2 = MakeSymbol("shen.return?")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symundefmacro = MakeSymbol("undefmacro")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4show = MakeSymbol("shen.show")
var symKey = MakeSymbol("Key")
var symrelease = MakeSymbol("release")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symread_1file = MakeSymbol("read-file")
var symshen_4shen = MakeSymbol("shen.shen")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var syminternal = MakeSymbol("internal")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4let_b = MakeSymbol("shen.let!")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4input_1macro = MakeSymbol("shen.input-macro")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symtlv = MakeSymbol("tlv")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var sym_dhush_d = MakeSymbol("*hush*")
var symshen_4error_1macro = MakeSymbol("shen.error-macro")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symappend = MakeSymbol("append")
var symshen_4mod = MakeSymbol("shen.mod")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symfix = MakeSymbol("fix")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symmapcan = MakeSymbol("mapcan")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symexception = MakeSymbol("exception")
var symFinish = MakeSymbol("Finish")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4optimise_1typing = MakeSymbol("shen.optimise-typing")
var symshen_4join = MakeSymbol("shen.join")
var syminput_7 = MakeSymbol("input+")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symshen_4print_1residue = MakeSymbol("shen.print-residue")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symoutput = MakeSymbol("output")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symshen_4compile_1assumption = MakeSymbol("shen.compile-assumption")
var symshen_4tabulate_1passive = MakeSymbol("shen.tabulate-passive")
var symshen_4source = MakeSymbol("shen.source")
var symassoc = MakeSymbol("assoc")
var sym_a_a = MakeSymbol("==")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symwhen = MakeSymbol("when")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symshen_4conscode = MakeSymbol("shen.conscode")
var sym_dport_d = MakeSymbol("*port*")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4vertical = MakeSymbol("shen.vertical")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symabsvector = MakeSymbol("absvector")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symshen_4_dprolog_1vector_d = MakeSymbol("shen.*prolog-vector*")
var symshen_4keep_1looking = MakeSymbol("shen.keep-looking")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symsnd = MakeSymbol("snd")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var symshen_4defmacro_1macro = MakeSymbol("shen.defmacro-macro")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symmacroexpand = MakeSymbol("macroexpand")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symV = MakeSymbol("V")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symshen_4parsed_2 = MakeSymbol("shen.parsed?")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var sym_dversion_d = MakeSymbol("*version*")
var symand = MakeSymbol("and")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symfork = MakeSymbol("fork")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var sym_5_1vector = MakeSymbol("<-vector")
var symshen_4datatype_1macro = MakeSymbol("shen.datatype-macro")
var sym_d = MakeSymbol("*")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symstoutput = MakeSymbol("stoutput")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4op = MakeSymbol("shen.op")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symnumber_2 = MakeSymbol("number?")
var symbind = MakeSymbol("bind")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symvector_2 = MakeSymbol("vector?")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symstinput = MakeSymbol("stinput")
var symport = MakeSymbol("port")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symK = MakeSymbol("K")
var symTime = MakeSymbol("Time")
var symshen_4passive_2 = MakeSymbol("shen.passive?")
var sym_4_4_4 = MakeSymbol("...")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var sym_i = MakeSymbol("{")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4a = MakeSymbol("shen.a")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4else = MakeSymbol("shen.else")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symshen_4curry = MakeSymbol("shen.curry")
var symstep = MakeSymbol("step")
var symget_1time = MakeSymbol("get-time")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symParse = MakeSymbol("Parse")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symvalue = MakeSymbol("value")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var sym_b = MakeSymbol("!")
var symshen_4foundit_b = MakeSymbol("shen.foundit!")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var sympr = MakeSymbol("pr")
var symcases = MakeSymbol("cases")
var symsymbol = MakeSymbol("symbol")
var symshen_4compile_1premise = MakeSymbol("shen.compile-premise")
var symshen_4app = MakeSymbol("shen.app")
var sym_dstinput_d = MakeSymbol("*stinput*")
var syminferences = MakeSymbol("inferences")
var symos = MakeSymbol("os")
var symadjoin = MakeSymbol("adjoin")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symdefun = MakeSymbol("defun")
var symshen_4tlstream = MakeSymbol("shen.tlstream")
var symshen_4t = MakeSymbol("shen.t")
var symtl = MakeSymbol("tl")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symit = MakeSymbol("it")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symstr = MakeSymbol("str")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symexternal = MakeSymbol("external")
var symshen_4_dlambdatable_d = MakeSymbol("shen.*lambdatable*")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var symshen_4selectors = MakeSymbol("shen.selectors")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symps = MakeSymbol("ps")
var sympackage_2 = MakeSymbol("package?")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symshen_4occurs_2 = MakeSymbol("shen.occurs?")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4compile_1premise_1h = MakeSymbol("shen.compile-premise-h")
var symtuple_2 = MakeSymbol("tuple?")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symResult = MakeSymbol("Result")
var symshen_4op2 = MakeSymbol("shen.op2")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4abs_1macro = MakeSymbol("shen.abs-macro")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symspy = MakeSymbol("spy")
var symfunction = MakeSymbol("function")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var symY = MakeSymbol("Y")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symloaded = MakeSymbol("loaded")
var symTm = MakeSymbol("Tm")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4_ahd_2 = MakeSymbol("shen.=hd?")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symsymbol_2 = MakeSymbol("symbol?")
var symstring_2 = MakeSymbol("string?")
var symread_1byte = MakeSymbol("read-byte")
var symunit = MakeSymbol("unit")
var symshen_4head_1foundit_b = MakeSymbol("shen.head-foundit!")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4optimise_1passive_1h = MakeSymbol("shen.optimise-passive-h")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var symmap = MakeSymbol("map")
var symoptimise = MakeSymbol("optimise")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symout = MakeSymbol("out")
var symlambda = MakeSymbol("lambda")
var symFreeze = MakeSymbol("Freeze")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symstring_1_6n = MakeSymbol("string->n")
var symgensym = MakeSymbol("gensym")
var sym_dlanguage_d = MakeSymbol("*language*")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4inline_2 = MakeSymbol("shen.inline?")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symlazy = MakeSymbol("lazy")
var symlimit = MakeSymbol("limit")
var sym_8v = MakeSymbol("@v")
var symdefcc = MakeSymbol("defcc")
var symB = MakeSymbol("B")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symshen_4hdstream = MakeSymbol("shen.hdstream")
var symtc = MakeSymbol("tc")
var symshen_4comb = MakeSymbol("shen.comb")
var symtlstr = MakeSymbol("tlstr")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symRecord = MakeSymbol("Record")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symcons_2 = MakeSymbol("cons?")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symor = MakeSymbol("or")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symread = MakeSymbol("read")
var syminput = MakeSymbol("input")
var symwarn = MakeSymbol("warn")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4correct = MakeSymbol("shen.correct")
var sym_8p = MakeSymbol("@p")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symsum = MakeSymbol("sum")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var sym_6_a = MakeSymbol(">=")
var symrun = MakeSymbol("run")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4freshen = MakeSymbol("shen.freshen")
var symfst = MakeSymbol("fst")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4cases_1macro = MakeSymbol("shen.cases-macro")
var symmake_1string = MakeSymbol("make-string")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symcall = MakeSymbol("call")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symfn = MakeSymbol("fn")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symn_1_6string = MakeSymbol("n->string")
var symtime = MakeSymbol("time")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symatom_2 = MakeSymbol("atom?")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var sympackage = MakeSymbol("package")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4change_1pointer_1value = MakeSymbol("shen.change-pointer-value")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symshen_4rpted_2 = MakeSymbol("shen.rpted?")
var symin = MakeSymbol("in")
var symboolean = MakeSymbol("boolean")
var symshen_4compile_1side_1condition = MakeSymbol("shen.compile-side-condition")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symsynonyms = MakeSymbol("synonyms")
var symshen_4_8s_1macro = MakeSymbol("shen.@s-macro")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4skip = MakeSymbol("shen.skip")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symhd = MakeSymbol("hd")
var symshen_4branch = MakeSymbol("shen.branch")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4f = MakeSymbol("shen.f")
var symunput = MakeSymbol("unput")
var symexplode = MakeSymbol("explode")
var symshen_4_5end_6 = MakeSymbol("shen.<end>")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4unix = MakeSymbol("shen.unix")
var sym_a_a_6 = MakeSymbol("==>")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symhash = MakeSymbol("hash")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symcn = MakeSymbol("cn")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var symdo = MakeSymbol("do")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var sym_c_4 = MakeSymbol("/.")
var symshen_4headstream = MakeSymbol("shen.headstream")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symwhere = MakeSymbol("where")
var symdeclare = MakeSymbol("declare")
var symshen_4_1m = MakeSymbol("shen.-m")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symshen_4nvars = MakeSymbol("shen.nvars")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symfreeze = MakeSymbol("freeze")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symshen_4initialise_1lambda_1tables = MakeSymbol("shen.initialise-lambda-tables")
var sympreclude = MakeSymbol("preclude")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4tame = MakeSymbol("shen.tame")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symP = MakeSymbol("P")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4compile_1consequent = MakeSymbol("shen.compile-consequent")
var symshen_4optimise_1passive = MakeSymbol("shen.optimise-passive")
var symget = MakeSymbol("get")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symshen_4prolog_1macro = MakeSymbol("shen.prolog-macro")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symSynCons = MakeSymbol("SynCons")
var symeval = MakeSymbol("eval")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4output_1macro = MakeSymbol("shen.output-macro")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symA = MakeSymbol("A")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4r = MakeSymbol("shen.r")
var sym_5_1 = MakeSymbol("<-")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symshen_4interror = MakeSymbol("shen.interror")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symcd = MakeSymbol("cd")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symelement_2 = MakeSymbol("element?")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symopen = MakeSymbol("open")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symlineread = MakeSymbol("lineread")
var sym_3 = MakeSymbol("$")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4cons_1form_1with_1modes = MakeSymbol("shen.cons-form-with-modes")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symis = MakeSymbol("is")
var symintern = MakeSymbol("intern")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4compile_1assumptions = MakeSymbol("shen.compile-assumptions")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symshen_4make_1string_1macro = MakeSymbol("shen.make-string-macro")
var sym_5_a = MakeSymbol("<=")
var symstream = MakeSymbol("stream")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen = MakeSymbol("shen")
var symdatatype = MakeSymbol("datatype")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var sym_1 = MakeSymbol("-")
var symversion = MakeSymbol("version")
var sym_dporters_d = MakeSymbol("*porters*")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symarity = MakeSymbol("arity")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symset = MakeSymbol("set")
var symfile = MakeSymbol("file")
var syminline = MakeSymbol("inline")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symhead = MakeSymbol("head")
var symnot = MakeSymbol("not")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symshen_4search = MakeSymbol("shen.search")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symbound_2 = MakeSymbol("bound?")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symD = MakeSymbol("D")
var symvariable_2 = MakeSymbol("variable?")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symdefine = MakeSymbol("define")
var sym_6 = MakeSymbol(">")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symNews = MakeSymbol("News")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symimplementation = MakeSymbol("implementation")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4tls = MakeSymbol("shen.tls")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symC = MakeSymbol("C")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symprolog_2 = MakeSymbol("prolog?")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4cons_1form_1no_1modes = MakeSymbol("shen.cons-form-no-modes")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symtype = MakeSymbol("type")
var symif = MakeSymbol("if")
var sym_5 = MakeSymbol("<")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4remove_1if_1unused = MakeSymbol("shen.remove-if-unused")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symshen_4nl_1macro = MakeSymbol("shen.nl-macro")
var symprofile = MakeSymbol("profile")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4decons = MakeSymbol("shen.decons")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var sym_5_b_6 = MakeSymbol("<!>")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var sym_5_1address = MakeSymbol("<-address")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symshen_4modh = MakeSymbol("shen.modh")
var symtail = MakeSymbol("tail")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symload = MakeSymbol("load")
var symshen_4passive = MakeSymbol("shen.passive")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symintersection = MakeSymbol("intersection")
var symcons = MakeSymbol("cons")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4synonyms_1macro = MakeSymbol("shen.synonyms-macro")
var symshen_4lock = MakeSymbol("shen.lock")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symnull = MakeSymbol("null")
var sym_7 = MakeSymbol("+")
var symprint = MakeSymbol("print")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symremove = MakeSymbol("remove")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4put_cget_1macro = MakeSymbol("shen.put/get-macro")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symreverse = MakeSymbol("reverse")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symreceive = MakeSymbol("receive")
var sym_e_e = MakeSymbol("&&")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symtc_2 = MakeSymbol("tc?")
var symporters = MakeSymbol("porters")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var symshen_4defcc_1macro = MakeSymbol("shen.defcc-macro")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symnth = MakeSymbol("nth")
var sym_c = MakeSymbol("/")
var symdefmacro = MakeSymbol("defmacro")
var symshen_4reset_1prolog_1vector = MakeSymbol("shen.reset-prolog-vector")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symerror = MakeSymbol("error")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symis_b = MakeSymbol("is!")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symunion = MakeSymbol("union")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var syminclude = MakeSymbol("include")
var symu_b = MakeSymbol("u!")
var symPrevious = MakeSymbol("Previous")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4remove_1pointer = MakeSymbol("shen.remove-pointer")
var sym_j = MakeSymbol("}")
var symshen_4hds = MakeSymbol("shen.hds")
var symlet = MakeSymbol("let")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var sym_8s = MakeSymbol("@s")
var symempty_2 = MakeSymbol("empty?")
var symcompile = MakeSymbol("compile")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var symlength = MakeSymbol("length")
var symshen_4u_b_1macro = MakeSymbol("shen.u!-macro")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symhdstr = MakeSymbol("hdstr")
var sym_1_6 = MakeSymbol("->")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symsystemf = MakeSymbol("systemf")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symfresh = MakeSymbol("fresh")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symstring = MakeSymbol("string")
var symStart = MakeSymbol("Start")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symnl = MakeSymbol("nl")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symsubst = MakeSymbol("subst")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symprofile_1results = MakeSymbol("profile-results")
var symlist = MakeSymbol("list")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4horizontal = MakeSymbol("shen.horizontal")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symshen_4s = MakeSymbol("shen.s")
var symfactorise = MakeSymbol("factorise")
var symshen_4nchars = MakeSymbol("shen.nchars")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4body_1foundit_b = MakeSymbol("shen.body-foundit!")
var symfail = MakeSymbol("fail")
var symdestroy = MakeSymbol("destroy")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symnumber = MakeSymbol("number")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symspecialise = MakeSymbol("specialise")
var symshen_4procedure_1call = MakeSymbol("shen.procedure-call")
var symvar_2 = MakeSymbol("var?")
var sympos = MakeSymbol("pos")
var symsave = MakeSymbol("save")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var sym_e = MakeSymbol("&")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var sym__ = MakeSymbol("_")
var symshen_4beta = MakeSymbol("shen.beta")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symreturn = MakeSymbol("return")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symvector = MakeSymbol("vector")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symcond = MakeSymbol("cond")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4compile_1side_1conditions = MakeSymbol("shen.compile-side-conditions")
var symshen_4compile_1search_1procedure = MakeSymbol("shen.compile-search-procedure")
var symthaw = MakeSymbol("thaw")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symbar_b = MakeSymbol("bar!")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4then = MakeSymbol("shen.then")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symmode = MakeSymbol("mode")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var sym_dos_d = MakeSymbol("*os*")
var symin_1package = MakeSymbol("in-package")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symshen_4pushsemantics = MakeSymbol("shen.pushsemantics")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4timer_1macro = MakeSymbol("shen.timer-macro")
var symX = MakeSymbol("X")
var symshen_4no_1action = MakeSymbol("shen.no-action")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
