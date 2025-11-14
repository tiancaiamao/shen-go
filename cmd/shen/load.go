package main

import . "github.com/tiancaiamao/shen-go/kl"

var LoadMain = MakeNative(func(__e *ControlFlow) {
tmp9500 := MakeNative(func(__e *ControlFlow) {
V874 := __e.Get(1)
_ = V874
tmp9501 := MakeNative(func(__e *ControlFlow) {
W875 := __e.Get(1)
_ = W875
tmp9502 := MakeNative(func(__e *ControlFlow) {
W876 := __e.Get(1)
_ = W876
tmp9503 := MakeNative(func(__e *ControlFlow) {
W882 := __e.Get(1)
_ = W882
__e.Return(symloaded)
return
}, 1)

var ifres9504 Obj

if True == W875 {
tmp9505 := Call(__e, PrimFunc(syminferences))


tmp9506 := Call(__e, PrimFunc(symshen_4app), tmp9505, MakeString(" inferences\n"), symshen_4a)


tmp9507 := PrimStringConcat(MakeString("\ntypechecked in "), tmp9506)

tmp9508 := Call(__e, PrimFunc(symstoutput))


tmp9509 := Call(__e, PrimFunc(sympr), tmp9507, tmp9508)


ifres9504 = tmp9509


} else {
ifres9504 = symshen_4skip


}

__e.TailApply(tmp9503, ifres9504)
return


}, 1)

tmp9510 := MakeNative(func(__e *ControlFlow) {
W877 := __e.Get(1)
_ = W877
tmp9511 := MakeNative(func(__e *ControlFlow) {
W878 := __e.Get(1)
_ = W878
tmp9512 := MakeNative(func(__e *ControlFlow) {
W879 := __e.Get(1)
_ = W879
tmp9513 := MakeNative(func(__e *ControlFlow) {
W880 := __e.Get(1)
_ = W880
tmp9514 := MakeNative(func(__e *ControlFlow) {
W881 := __e.Get(1)
_ = W881
__e.Return(W878)
return
}, 1)

tmp9515 := PrimStr(W880)

tmp9516 := PrimStringConcat(tmp9515, MakeString(" secs\n"))

tmp9517 := PrimStringConcat(MakeString("\nrun time: "), tmp9516)

tmp9518 := Call(__e, PrimFunc(symstoutput))


tmp9519 := Call(__e, PrimFunc(sympr), tmp9517, tmp9518)


__e.TailApply(tmp9514, tmp9519)
return


}, 1)

tmp9520 := PrimNumberSubtract(W879, W877)

__e.TailApply(tmp9513, tmp9520)
return


}, 1)

tmp9521 := PrimGetTime(symrun)

__e.TailApply(tmp9512, tmp9521)
return


}, 1)

tmp9522 := Call(__e, PrimFunc(symread_1file), V874)


tmp9523 := Call(__e, PrimFunc(symshen_4load_1help), W875, tmp9522)


__e.TailApply(tmp9511, tmp9523)
return


}, 1)

tmp9524 := PrimGetTime(symrun)

tmp9525 := Call(__e, tmp9510, tmp9524)


__e.TailApply(tmp9502, tmp9525)
return


}, 1)

tmp9526 := PrimValue(symshen_4_dtc_d)

__e.TailApply(tmp9501, tmp9526)
return


}, 1)

tmp9527 := Call(__e, ns2_1set, symload, tmp9500)


_ = tmp9527

tmp9528 := MakeNative(func(__e *ControlFlow) {
V885 := __e.Get(1)
_ = V885
V886 := __e.Get(2)
_ = V886
tmp9530 := PrimEqual(False, V885)

if True == tmp9530 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V886)
return
} else {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V886)
return
}


}, 2)

tmp9531 := Call(__e, ns2_1set, symshen_4load_1help, tmp9528)


_ = tmp9531

tmp9532 := MakeNative(func(__e *ControlFlow) {
V887 := __e.Get(1)
_ = V887
tmp9533 := MakeNative(func(__e *ControlFlow) {
Z888 := __e.Get(1)
_ = Z888
tmp9534 := Call(__e, PrimFunc(symshen_4shen_1_6kl), Z888)


tmp9535 := Call(__e, PrimFunc(symeval_1kl), tmp9534)


tmp9536 := Call(__e, PrimFunc(symshen_4app), tmp9535, MakeString("\n"), symshen_4s)


tmp9537 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp9536, tmp9537)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp9533, V887)
return


}, 1)

tmp9538 := Call(__e, ns2_1set, symshen_4eval_1and_1print, tmp9532)


_ = tmp9538

tmp9539 := MakeNative(func(__e *ControlFlow) {
V889 := __e.Get(1)
_ = V889
tmp9540 := MakeNative(func(__e *ControlFlow) {
W890 := __e.Get(1)
_ = W890
tmp9541 := MakeNative(func(__e *ControlFlow) {
W892 := __e.Get(1)
_ = W892
tmp9542 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4work_1through), V889)
return
}, 0)

tmp9543 := MakeNative(func(__e *ControlFlow) {
Z894 := __e.Get(1)
_ = Z894
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z894, W890)
return
}, 1)

__e.TailApply(try_1catch, tmp9542, tmp9543)
return


}, 1)

tmp9544 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4assumetypes), W890)
return
}, 0)

tmp9545 := MakeNative(func(__e *ControlFlow) {
Z893 := __e.Get(1)
_ = Z893
__e.TailApply(PrimFunc(symshen_4unwind_1types), Z893, W890)
return
}, 1)

tmp9546 := Call(__e, try_1catch, tmp9544, tmp9545)


__e.TailApply(tmp9541, tmp9546)
return


}, 1)

tmp9547 := MakeNative(func(__e *ControlFlow) {
Z891 := __e.Get(1)
_ = Z891
__e.TailApply(PrimFunc(symshen_4typetable), Z891)
return
}, 1)

tmp9548 := Call(__e, PrimFunc(symmapcan), tmp9547, V889)


__e.TailApply(tmp9540, tmp9548)
return


}, 1)

tmp9549 := Call(__e, ns2_1set, symshen_4check_1eval_1and_1print, tmp9539)


_ = tmp9549

tmp9550 := MakeNative(func(__e *ControlFlow) {
V899 := __e.Get(1)
_ = V899
tmp9595 := PrimIsPair(V899)

var ifres9576 Obj

if True == tmp9595 {
tmp9593 := PrimHead(V899)

tmp9594 := PrimEqual(symdefine, tmp9593)

var ifres9578 Obj

if True == tmp9594 {
tmp9591 := PrimTail(V899)

tmp9592 := PrimIsPair(tmp9591)

var ifres9580 Obj

if True == tmp9592 {
tmp9588 := PrimTail(V899)

tmp9589 := PrimTail(tmp9588)

tmp9590 := PrimIsPair(tmp9589)

var ifres9582 Obj

if True == tmp9590 {
tmp9584 := PrimTail(V899)

tmp9585 := PrimTail(tmp9584)

tmp9586 := PrimHead(tmp9585)

tmp9587 := PrimEqual(sym_i, tmp9586)

var ifres9583 Obj

if True == tmp9587 {
ifres9583 = True


} else {
ifres9583 = False


}

ifres9582 = ifres9583


} else {
ifres9582 = False


}

var ifres9581 Obj

if True == ifres9582 {
ifres9581 = True


} else {
ifres9581 = False


}

ifres9580 = ifres9581


} else {
ifres9580 = False


}

var ifres9579 Obj

if True == ifres9580 {
ifres9579 = True


} else {
ifres9579 = False


}

ifres9578 = ifres9579


} else {
ifres9578 = False


}

var ifres9577 Obj

if True == ifres9578 {
ifres9577 = True


} else {
ifres9577 = False


}

ifres9576 = ifres9577


} else {
ifres9576 = False


}

if True == ifres9576 {
tmp9551 := PrimTail(V899)

tmp9552 := PrimHead(tmp9551)

tmp9553 := PrimTail(V899)

tmp9554 := PrimHead(tmp9553)

tmp9555 := PrimTail(V899)

tmp9556 := PrimTail(tmp9555)

tmp9557 := PrimTail(tmp9556)

tmp9558 := Call(__e, PrimFunc(symshen_4type_1F), tmp9554, tmp9557)


tmp9559 := Call(__e, PrimFunc(symshen_4rectify_1type), tmp9558)


tmp9560 := PrimCons(tmp9559, Nil)

__e.Return(PrimCons(tmp9552, tmp9560))
return


} else {
tmp9574 := PrimIsPair(V899)

var ifres9566 Obj

if True == tmp9574 {
tmp9572 := PrimHead(V899)

tmp9573 := PrimEqual(symdefine, tmp9572)

var ifres9568 Obj

if True == tmp9573 {
tmp9570 := PrimTail(V899)

tmp9571 := PrimIsPair(tmp9570)

var ifres9569 Obj

if True == tmp9571 {
ifres9569 = True


} else {
ifres9569 = False


}

ifres9568 = ifres9569


} else {
ifres9568 = False


}

var ifres9567 Obj

if True == ifres9568 {
ifres9567 = True


} else {
ifres9567 = False


}

ifres9566 = ifres9567


} else {
ifres9566 = False


}

if True == ifres9566 {
tmp9561 := PrimTail(V899)

tmp9562 := PrimHead(tmp9561)

tmp9563 := Call(__e, PrimFunc(symshen_4app), tmp9562, MakeString("\n"), symshen_4a)


tmp9564 := PrimStringConcat(MakeString("missing { in "), tmp9563)

__e.Return(PrimSimpleError(tmp9564))
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp9596 := Call(__e, ns2_1set, symshen_4typetable, tmp9550)


_ = tmp9596

tmp9597 := MakeNative(func(__e *ControlFlow) {
V906 := __e.Get(1)
_ = V906
V907 := __e.Get(2)
_ = V907
tmp9610 := PrimIsPair(V907)

var ifres9606 Obj

if True == tmp9610 {
tmp9608 := PrimHead(V907)

tmp9609 := PrimEqual(sym_j, tmp9608)

var ifres9607 Obj

if True == tmp9609 {
ifres9607 = True


} else {
ifres9607 = False


}

ifres9606 = ifres9607


} else {
ifres9606 = False


}

if True == ifres9606 {
__e.Return(Nil)
return
} else {
tmp9604 := PrimIsPair(V907)

if True == tmp9604 {
tmp9598 := PrimHead(V907)

tmp9599 := PrimTail(V907)

tmp9600 := Call(__e, PrimFunc(symshen_4type_1F), V906, tmp9599)


__e.Return(PrimCons(tmp9598, tmp9600))
return


} else {
tmp9601 := Call(__e, PrimFunc(symshen_4app), V906, MakeString("\n"), symshen_4a)


tmp9602 := PrimStringConcat(MakeString("missing } in "), tmp9601)

__e.Return(PrimSimpleError(tmp9602))
return


}


}


}, 2)

tmp9611 := Call(__e, ns2_1set, symshen_4type_1F, tmp9597)


_ = tmp9611

tmp9612 := MakeNative(func(__e *ControlFlow) {
V910 := __e.Get(1)
_ = V910
tmp9626 := PrimEqual(Nil, V910)

if True == tmp9626 {
__e.Return(Nil)
return
} else {
tmp9624 := PrimIsPair(V910)

var ifres9620 Obj

if True == tmp9624 {
tmp9622 := PrimTail(V910)

tmp9623 := PrimIsPair(tmp9622)

var ifres9621 Obj

if True == tmp9623 {
ifres9621 = True


} else {
ifres9621 = False


}

ifres9620 = ifres9621


} else {
ifres9620 = False


}

if True == ifres9620 {
tmp9613 := PrimHead(V910)

tmp9614 := PrimTail(V910)

tmp9615 := PrimHead(tmp9614)

tmp9616 := Call(__e, PrimFunc(symdeclare), tmp9613, tmp9615)


_ = tmp9616

tmp9617 := PrimTail(V910)

tmp9618 := PrimTail(tmp9617)

__e.TailApply(PrimFunc(symshen_4assumetypes), tmp9618)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assumetype")))
return
}


}


}, 1)

tmp9627 := Call(__e, ns2_1set, symshen_4assumetypes, tmp9612)


_ = tmp9627

tmp9628 := MakeNative(func(__e *ControlFlow) {
V915 := __e.Get(1)
_ = V915
V916 := __e.Get(2)
_ = V916
tmp9639 := PrimIsPair(V916)

var ifres9635 Obj

if True == tmp9639 {
tmp9637 := PrimTail(V916)

tmp9638 := PrimIsPair(tmp9637)

var ifres9636 Obj

if True == tmp9638 {
ifres9636 = True


} else {
ifres9636 = False


}

ifres9635 = ifres9636


} else {
ifres9635 = False


}

if True == ifres9635 {
tmp9629 := PrimHead(V916)

tmp9630 := Call(__e, PrimFunc(symdestroy), tmp9629)


_ = tmp9630

tmp9631 := PrimTail(V916)

tmp9632 := PrimTail(tmp9631)

__e.TailApply(PrimFunc(symshen_4unwind_1types), V915, tmp9632)
return


} else {
tmp9633 := PrimErrorToString(V915)

__e.Return(PrimSimpleError(tmp9633))
return


}


}, 2)

tmp9640 := Call(__e, ns2_1set, symshen_4unwind_1types, tmp9628)


_ = tmp9640

tmp9641 := MakeNative(func(__e *ControlFlow) {
V919 := __e.Get(1)
_ = V919
tmp9690 := PrimEqual(Nil, V919)

if True == tmp9690 {
__e.Return(Nil)
return
} else {
tmp9688 := PrimIsPair(V919)

var ifres9673 Obj

if True == tmp9688 {
tmp9686 := PrimTail(V919)

tmp9687 := PrimIsPair(tmp9686)

var ifres9675 Obj

if True == tmp9687 {
tmp9683 := PrimTail(V919)

tmp9684 := PrimTail(tmp9683)

tmp9685 := PrimIsPair(tmp9684)

var ifres9677 Obj

if True == tmp9685 {
tmp9679 := PrimTail(V919)

tmp9680 := PrimHead(tmp9679)

tmp9681 := PrimIntern(MakeString(":"))

tmp9682 := PrimEqual(tmp9680, tmp9681)

var ifres9678 Obj

if True == tmp9682 {
ifres9678 = True


} else {
ifres9678 = False


}

ifres9677 = ifres9678


} else {
ifres9677 = False


}

var ifres9676 Obj

if True == ifres9677 {
ifres9676 = True


} else {
ifres9676 = False


}

ifres9675 = ifres9676


} else {
ifres9675 = False


}

var ifres9674 Obj

if True == ifres9675 {
ifres9674 = True


} else {
ifres9674 = False


}

ifres9673 = ifres9674


} else {
ifres9673 = False


}

if True == ifres9673 {
tmp9642 := MakeNative(func(__e *ControlFlow) {
W920 := __e.Get(1)
_ = W920
tmp9658 := PrimEqual(W920, False)

if True == tmp9658 {
__e.TailApply(PrimFunc(symshen_4type_1error))
return
} else {
tmp9643 := MakeNative(func(__e *ControlFlow) {
W921 := __e.Get(1)
_ = W921
tmp9644 := MakeNative(func(__e *ControlFlow) {
W922 := __e.Get(1)
_ = W922
tmp9645 := PrimTail(V919)

tmp9646 := PrimTail(tmp9645)

tmp9647 := PrimTail(tmp9646)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9647)
return


}, 1)

tmp9648 := Call(__e, PrimFunc(symshen_4pretty_1type), W920)


tmp9649 := Call(__e, PrimFunc(symshen_4app), tmp9648, MakeString("\n"), symshen_4r)


tmp9650 := PrimStringConcat(MakeString(" : "), tmp9649)

tmp9651 := Call(__e, PrimFunc(symshen_4app), W921, tmp9650, symshen_4s)


tmp9652 := Call(__e, PrimFunc(symstoutput))


tmp9653 := Call(__e, PrimFunc(sympr), tmp9651, tmp9652)


__e.TailApply(tmp9644, tmp9653)
return


}, 1)

tmp9654 := PrimHead(V919)

tmp9655 := Call(__e, PrimFunc(symshen_4shen_1_6kl), tmp9654)


tmp9656 := Call(__e, PrimFunc(symeval_1kl), tmp9655)


__e.TailApply(tmp9643, tmp9656)
return


}


}, 1)

tmp9659 := PrimHead(V919)

tmp9660 := PrimTail(V919)

tmp9661 := PrimTail(tmp9660)

tmp9662 := PrimHead(tmp9661)

tmp9663 := Call(__e, PrimFunc(symshen_4typecheck), tmp9659, tmp9662)


__e.TailApply(tmp9642, tmp9663)
return


} else {
tmp9671 := PrimIsPair(V919)

if True == tmp9671 {
tmp9664 := PrimHead(V919)

tmp9665 := PrimIntern(MakeString(":"))

tmp9666 := PrimTail(V919)

tmp9667 := PrimCons(symA, tmp9666)

tmp9668 := PrimCons(tmp9665, tmp9667)

tmp9669 := PrimCons(tmp9664, tmp9668)

__e.TailApply(PrimFunc(symshen_4work_1through), tmp9669)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.work-through")))
return
}


}


}


}, 1)

tmp9691 := Call(__e, ns2_1set, symshen_4work_1through, tmp9641)


_ = tmp9691

tmp9692 := MakeNative(func(__e *ControlFlow) {
V924 := __e.Get(1)
_ = V924
tmp9834 := PrimIsPair(V924)

var ifres9708 Obj

if True == tmp9834 {
tmp9832 := PrimHead(V924)

tmp9833 := PrimIsPair(tmp9832)

var ifres9710 Obj

if True == tmp9833 {
tmp9829 := PrimHead(V924)

tmp9830 := PrimHead(tmp9829)

tmp9831 := PrimEqual(symlist, tmp9830)

var ifres9712 Obj

if True == tmp9831 {
tmp9826 := PrimHead(V924)

tmp9827 := PrimTail(tmp9826)

tmp9828 := PrimIsPair(tmp9827)

var ifres9714 Obj

if True == tmp9828 {
tmp9822 := PrimHead(V924)

tmp9823 := PrimTail(tmp9822)

tmp9824 := PrimTail(tmp9823)

tmp9825 := PrimEqual(Nil, tmp9824)

var ifres9716 Obj

if True == tmp9825 {
tmp9820 := PrimTail(V924)

tmp9821 := PrimIsPair(tmp9820)

var ifres9718 Obj

if True == tmp9821 {
tmp9817 := PrimTail(V924)

tmp9818 := PrimHead(tmp9817)

tmp9819 := PrimEqual(sym_1_1_6, tmp9818)

var ifres9720 Obj

if True == tmp9819 {
tmp9814 := PrimTail(V924)

tmp9815 := PrimTail(tmp9814)

tmp9816 := PrimIsPair(tmp9815)

var ifres9722 Obj

if True == tmp9816 {
tmp9810 := PrimTail(V924)

tmp9811 := PrimTail(tmp9810)

tmp9812 := PrimHead(tmp9811)

tmp9813 := PrimIsPair(tmp9812)

var ifres9724 Obj

if True == tmp9813 {
tmp9805 := PrimTail(V924)

tmp9806 := PrimTail(tmp9805)

tmp9807 := PrimHead(tmp9806)

tmp9808 := PrimHead(tmp9807)

tmp9809 := PrimEqual(symstr, tmp9808)

var ifres9726 Obj

if True == tmp9809 {
tmp9800 := PrimTail(V924)

tmp9801 := PrimTail(tmp9800)

tmp9802 := PrimHead(tmp9801)

tmp9803 := PrimTail(tmp9802)

tmp9804 := PrimIsPair(tmp9803)

var ifres9728 Obj

if True == tmp9804 {
tmp9794 := PrimTail(V924)

tmp9795 := PrimTail(tmp9794)

tmp9796 := PrimHead(tmp9795)

tmp9797 := PrimTail(tmp9796)

tmp9798 := PrimHead(tmp9797)

tmp9799 := PrimIsPair(tmp9798)

var ifres9730 Obj

if True == tmp9799 {
tmp9787 := PrimTail(V924)

tmp9788 := PrimTail(tmp9787)

tmp9789 := PrimHead(tmp9788)

tmp9790 := PrimTail(tmp9789)

tmp9791 := PrimHead(tmp9790)

tmp9792 := PrimHead(tmp9791)

tmp9793 := PrimEqual(symlist, tmp9792)

var ifres9732 Obj

if True == tmp9793 {
tmp9780 := PrimTail(V924)

tmp9781 := PrimTail(tmp9780)

tmp9782 := PrimHead(tmp9781)

tmp9783 := PrimTail(tmp9782)

tmp9784 := PrimHead(tmp9783)

tmp9785 := PrimTail(tmp9784)

tmp9786 := PrimIsPair(tmp9785)

var ifres9734 Obj

if True == tmp9786 {
tmp9772 := PrimTail(V924)

tmp9773 := PrimTail(tmp9772)

tmp9774 := PrimHead(tmp9773)

tmp9775 := PrimTail(tmp9774)

tmp9776 := PrimHead(tmp9775)

tmp9777 := PrimTail(tmp9776)

tmp9778 := PrimTail(tmp9777)

tmp9779 := PrimEqual(Nil, tmp9778)

var ifres9736 Obj

if True == tmp9779 {
tmp9766 := PrimTail(V924)

tmp9767 := PrimTail(tmp9766)

tmp9768 := PrimHead(tmp9767)

tmp9769 := PrimTail(tmp9768)

tmp9770 := PrimTail(tmp9769)

tmp9771 := PrimIsPair(tmp9770)

var ifres9738 Obj

if True == tmp9771 {
tmp9759 := PrimTail(V924)

tmp9760 := PrimTail(tmp9759)

tmp9761 := PrimHead(tmp9760)

tmp9762 := PrimTail(tmp9761)

tmp9763 := PrimTail(tmp9762)

tmp9764 := PrimTail(tmp9763)

tmp9765 := PrimEqual(Nil, tmp9764)

var ifres9740 Obj

if True == tmp9765 {
tmp9755 := PrimTail(V924)

tmp9756 := PrimTail(tmp9755)

tmp9757 := PrimTail(tmp9756)

tmp9758 := PrimEqual(Nil, tmp9757)

var ifres9742 Obj

if True == tmp9758 {
tmp9744 := PrimHead(V924)

tmp9745 := PrimTail(tmp9744)

tmp9746 := PrimHead(tmp9745)

tmp9747 := PrimTail(V924)

tmp9748 := PrimTail(tmp9747)

tmp9749 := PrimHead(tmp9748)

tmp9750 := PrimTail(tmp9749)

tmp9751 := PrimHead(tmp9750)

tmp9752 := PrimTail(tmp9751)

tmp9753 := PrimHead(tmp9752)

tmp9754 := PrimEqual(tmp9746, tmp9753)

var ifres9743 Obj

if True == tmp9754 {
ifres9743 = True


} else {
ifres9743 = False


}

ifres9742 = ifres9743


} else {
ifres9742 = False


}

var ifres9741 Obj

if True == ifres9742 {
ifres9741 = True


} else {
ifres9741 = False


}

ifres9740 = ifres9741


} else {
ifres9740 = False


}

var ifres9739 Obj

if True == ifres9740 {
ifres9739 = True


} else {
ifres9739 = False


}

ifres9738 = ifres9739


} else {
ifres9738 = False


}

var ifres9737 Obj

if True == ifres9738 {
ifres9737 = True


} else {
ifres9737 = False


}

ifres9736 = ifres9737


} else {
ifres9736 = False


}

var ifres9735 Obj

if True == ifres9736 {
ifres9735 = True


} else {
ifres9735 = False


}

ifres9734 = ifres9735


} else {
ifres9734 = False


}

var ifres9733 Obj

if True == ifres9734 {
ifres9733 = True


} else {
ifres9733 = False


}

ifres9732 = ifres9733


} else {
ifres9732 = False


}

var ifres9731 Obj

if True == ifres9732 {
ifres9731 = True


} else {
ifres9731 = False


}

ifres9730 = ifres9731


} else {
ifres9730 = False


}

var ifres9729 Obj

if True == ifres9730 {
ifres9729 = True


} else {
ifres9729 = False


}

ifres9728 = ifres9729


} else {
ifres9728 = False


}

var ifres9727 Obj

if True == ifres9728 {
ifres9727 = True


} else {
ifres9727 = False


}

ifres9726 = ifres9727


} else {
ifres9726 = False


}

var ifres9725 Obj

if True == ifres9726 {
ifres9725 = True


} else {
ifres9725 = False


}

ifres9724 = ifres9725


} else {
ifres9724 = False


}

var ifres9723 Obj

if True == ifres9724 {
ifres9723 = True


} else {
ifres9723 = False


}

ifres9722 = ifres9723


} else {
ifres9722 = False


}

var ifres9721 Obj

if True == ifres9722 {
ifres9721 = True


} else {
ifres9721 = False


}

ifres9720 = ifres9721


} else {
ifres9720 = False


}

var ifres9719 Obj

if True == ifres9720 {
ifres9719 = True


} else {
ifres9719 = False


}

ifres9718 = ifres9719


} else {
ifres9718 = False


}

var ifres9717 Obj

if True == ifres9718 {
ifres9717 = True


} else {
ifres9717 = False


}

ifres9716 = ifres9717


} else {
ifres9716 = False


}

var ifres9715 Obj

if True == ifres9716 {
ifres9715 = True


} else {
ifres9715 = False


}

ifres9714 = ifres9715


} else {
ifres9714 = False


}

var ifres9713 Obj

if True == ifres9714 {
ifres9713 = True


} else {
ifres9713 = False


}

ifres9712 = ifres9713


} else {
ifres9712 = False


}

var ifres9711 Obj

if True == ifres9712 {
ifres9711 = True


} else {
ifres9711 = False


}

ifres9710 = ifres9711


} else {
ifres9710 = False


}

var ifres9709 Obj

if True == ifres9710 {
ifres9709 = True


} else {
ifres9709 = False


}

ifres9708 = ifres9709


} else {
ifres9708 = False


}

if True == ifres9708 {
tmp9693 := PrimTail(V924)

tmp9694 := PrimTail(tmp9693)

tmp9695 := PrimHead(tmp9694)

tmp9696 := PrimTail(tmp9695)

tmp9697 := PrimHead(tmp9696)

tmp9698 := PrimTail(V924)

tmp9699 := PrimTail(tmp9698)

tmp9700 := PrimHead(tmp9699)

tmp9701 := PrimTail(tmp9700)

tmp9702 := PrimTail(tmp9701)

tmp9703 := PrimCons(sym_a_a_6, tmp9702)

__e.Return(PrimCons(tmp9697, tmp9703))
return


} else {
tmp9706 := PrimIsPair(V924)

if True == tmp9706 {
tmp9704 := MakeNative(func(__e *ControlFlow) {
Z925 := __e.Get(1)
_ = Z925
__e.TailApply(PrimFunc(symshen_4pretty_1type), Z925)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9704, V924)
return


} else {
__e.Return(V924)
return
}


}


}, 1)

tmp9835 := Call(__e, ns2_1set, symshen_4pretty_1type, tmp9692)


_ = tmp9835

tmp9836 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimSimpleError(MakeString("type error\n")))
return
}, 0)

tmp9837 := Call(__e, ns2_1set, symshen_4type_1error, tmp9836)


_ = tmp9837

tmp9838 := MakeNative(func(__e *ControlFlow) {
V926 := __e.Get(1)
_ = V926
tmp9839 := MakeNative(func(__e *ControlFlow) {
W927 := __e.Get(1)
_ = W927
tmp9840 := MakeNative(func(__e *ControlFlow) {
W928 := __e.Get(1)
_ = W928
tmp9841 := MakeNative(func(__e *ControlFlow) {
W929 := __e.Get(1)
_ = W929
tmp9842 := MakeNative(func(__e *ControlFlow) {
W930 := __e.Get(1)
_ = W930
tmp9843 := MakeNative(func(__e *ControlFlow) {
W932 := __e.Get(1)
_ = W932
__e.Return(W927)
return
}, 1)

tmp9844 := Call(__e, PrimFunc(symshen_4write_1kl), W930, W929)


__e.TailApply(tmp9843, tmp9844)
return


}, 1)

tmp9845 := MakeNative(func(__e *ControlFlow) {
Z931 := __e.Get(1)
_ = Z931
tmp9846 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), Z931)


__e.TailApply(PrimFunc(symshen_4partial), tmp9846)
return


}, 1)

tmp9847 := Call(__e, PrimFunc(symmap), tmp9845, W928)


__e.TailApply(tmp9842, tmp9847)
return


}, 1)

tmp9848 := PrimOpenStream(W927, symout)

__e.TailApply(tmp9841, tmp9848)
return


}, 1)

tmp9849 := Call(__e, PrimFunc(symread_1file), V926)


__e.TailApply(tmp9840, tmp9849)
return


}, 1)

tmp9850 := Call(__e, PrimFunc(symshen_4klfile), V926)


__e.TailApply(tmp9839, tmp9850)
return


}, 1)

tmp9851 := Call(__e, ns2_1set, symbootstrap, tmp9838)


_ = tmp9851

tmp9852 := MakeNative(func(__e *ControlFlow) {
V933 := __e.Get(1)
_ = V933
tmp9875 := PrimIsPair(V933)

var ifres9862 Obj

if True == tmp9875 {
tmp9873 := PrimHead(V933)

tmp9874 := PrimEqual(symshen_4f_1error, tmp9873)

var ifres9864 Obj

if True == tmp9874 {
tmp9871 := PrimTail(V933)

tmp9872 := PrimIsPair(tmp9871)

var ifres9866 Obj

if True == tmp9872 {
tmp9868 := PrimTail(V933)

tmp9869 := PrimTail(tmp9868)

tmp9870 := PrimEqual(Nil, tmp9869)

var ifres9867 Obj

if True == tmp9870 {
ifres9867 = True


} else {
ifres9867 = False


}

ifres9866 = ifres9867


} else {
ifres9866 = False


}

var ifres9865 Obj

if True == ifres9866 {
ifres9865 = True


} else {
ifres9865 = False


}

ifres9864 = ifres9865


} else {
ifres9864 = False


}

var ifres9863 Obj

if True == ifres9864 {
ifres9863 = True


} else {
ifres9863 = False


}

ifres9862 = ifres9863


} else {
ifres9862 = False


}

if True == ifres9862 {
tmp9853 := PrimTail(V933)

tmp9854 := PrimHead(tmp9853)

tmp9855 := PrimStr(tmp9854)

tmp9856 := PrimStringConcat(MakeString("partial function "), tmp9855)

tmp9857 := PrimCons(tmp9856, Nil)

__e.Return(PrimCons(symsimple_1error, tmp9857))
return


} else {
tmp9860 := PrimIsPair(V933)

if True == tmp9860 {
tmp9858 := MakeNative(func(__e *ControlFlow) {
Z934 := __e.Get(1)
_ = Z934
__e.TailApply(PrimFunc(symshen_4partial), Z934)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp9858, V933)
return


} else {
__e.Return(V933)
return
}


}


}, 1)

tmp9876 := Call(__e, ns2_1set, symshen_4partial, tmp9852)


_ = tmp9876

tmp9877 := MakeNative(func(__e *ControlFlow) {
V937 := __e.Get(1)
_ = V937
V938 := __e.Get(2)
_ = V938
tmp9891 := PrimEqual(Nil, V937)

if True == tmp9891 {
__e.Return(PrimCloseStream(V938))
return
} else {
tmp9889 := PrimIsPair(V937)

var ifres9885 Obj

if True == tmp9889 {
tmp9887 := PrimHead(V937)

tmp9888 := PrimIsPair(tmp9887)

var ifres9886 Obj

if True == tmp9888 {
ifres9886 = True


} else {
ifres9886 = False


}

ifres9885 = ifres9886


} else {
ifres9885 = False


}

if True == ifres9885 {
tmp9878 := PrimTail(V937)

tmp9879 := PrimHead(V937)

tmp9880 := Call(__e, PrimFunc(symshen_4write_1kl_1h), tmp9879, V938)


_ = tmp9880

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9878, V938)
return


} else {
tmp9883 := PrimIsPair(V937)

if True == tmp9883 {
tmp9881 := PrimTail(V937)

__e.TailApply(PrimFunc(symshen_4write_1kl), tmp9881, V938)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.write-kl")))
return
}


}


}


}, 2)

tmp9892 := Call(__e, ns2_1set, symshen_4write_1kl, tmp9877)


_ = tmp9892

tmp9893 := MakeNative(func(__e *ControlFlow) {
V941 := __e.Get(1)
_ = V941
V942 := __e.Get(2)
_ = V942
tmp9933 := PrimIsPair(V941)

var ifres9896 Obj

if True == tmp9933 {
tmp9931 := PrimHead(V941)

tmp9932 := PrimEqual(symdefun, tmp9931)

var ifres9898 Obj

if True == tmp9932 {
tmp9929 := PrimTail(V941)

tmp9930 := PrimIsPair(tmp9929)

var ifres9900 Obj

if True == tmp9930 {
tmp9926 := PrimTail(V941)

tmp9927 := PrimHead(tmp9926)

tmp9928 := PrimEqual(symfail, tmp9927)

var ifres9902 Obj

if True == tmp9928 {
tmp9923 := PrimTail(V941)

tmp9924 := PrimTail(tmp9923)

tmp9925 := PrimIsPair(tmp9924)

var ifres9904 Obj

if True == tmp9925 {
tmp9919 := PrimTail(V941)

tmp9920 := PrimTail(tmp9919)

tmp9921 := PrimHead(tmp9920)

tmp9922 := PrimEqual(Nil, tmp9921)

var ifres9906 Obj

if True == tmp9922 {
tmp9915 := PrimTail(V941)

tmp9916 := PrimTail(tmp9915)

tmp9917 := PrimTail(tmp9916)

tmp9918 := PrimIsPair(tmp9917)

var ifres9908 Obj

if True == tmp9918 {
tmp9910 := PrimTail(V941)

tmp9911 := PrimTail(tmp9910)

tmp9912 := PrimTail(tmp9911)

tmp9913 := PrimTail(tmp9912)

tmp9914 := PrimEqual(Nil, tmp9913)

var ifres9909 Obj

if True == tmp9914 {
ifres9909 = True


} else {
ifres9909 = False


}

ifres9908 = ifres9909


} else {
ifres9908 = False


}

var ifres9907 Obj

if True == ifres9908 {
ifres9907 = True


} else {
ifres9907 = False


}

ifres9906 = ifres9907


} else {
ifres9906 = False


}

var ifres9905 Obj

if True == ifres9906 {
ifres9905 = True


} else {
ifres9905 = False


}

ifres9904 = ifres9905


} else {
ifres9904 = False


}

var ifres9903 Obj

if True == ifres9904 {
ifres9903 = True


} else {
ifres9903 = False


}

ifres9902 = ifres9903


} else {
ifres9902 = False


}

var ifres9901 Obj

if True == ifres9902 {
ifres9901 = True


} else {
ifres9901 = False


}

ifres9900 = ifres9901


} else {
ifres9900 = False


}

var ifres9899 Obj

if True == ifres9900 {
ifres9899 = True


} else {
ifres9899 = False


}

ifres9898 = ifres9899


} else {
ifres9898 = False


}

var ifres9897 Obj

if True == ifres9898 {
ifres9897 = True


} else {
ifres9897 = False


}

ifres9896 = ifres9897


} else {
ifres9896 = False


}

if True == ifres9896 {
__e.TailApply(PrimFunc(sympr), MakeString("(defun fail () shen.fail!)"), V942)
return
} else {
tmp9894 := Call(__e, PrimFunc(symshen_4app), V941, MakeString("\n\n"), symshen_4r)


__e.TailApply(PrimFunc(sympr), tmp9894, V942)
return


}


}, 2)

tmp9934 := Call(__e, ns2_1set, symshen_4write_1kl_1h, tmp9893)


_ = tmp9934

tmp9935 := MakeNative(func(__e *ControlFlow) {
V943 := __e.Get(1)
_ = V943
tmp9944 := PrimEqual(MakeString(""), V943)

if True == tmp9944 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9942 := PrimEqual(MakeString(".shen"), V943)

if True == tmp9942 {
__e.Return(MakeString(".kl"))
return
} else {
tmp9940 := Call(__e, PrimFunc(symshen_4_7string_2), V943)


if True == tmp9940 {
tmp9936 := Call(__e, PrimFunc(symhdstr), V943)


tmp9937 := PrimTailString(V943)

tmp9938 := Call(__e, PrimFunc(symshen_4klfile), tmp9937)


__e.TailApply(PrimFunc(sym_8s), tmp9936, tmp9938)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.klfile")))
return
}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4klfile, tmp9935)
return




}, 0)

