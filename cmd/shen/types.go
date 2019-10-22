package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__declare Obj // declare
var __defun__shen_4demodulate Obj // shen.demodulate
var __defun__shen_4variancy_1test Obj // shen.variancy-test
var __defun__shen_4variant_2 Obj // shen.variant?

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316482 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg316482)
return
}, 0))
__defun__declare = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4123 := __args[0]
_ = V4123
V4124 := __args[1]
_ = V4124
reg316483 := MakeSymbol("shen.*signedfuncs*")
reg316484 := PrimCons(V4123, V4124)
reg316485 := MakeSymbol("shen.*signedfuncs*")
reg316486 := PrimValue(reg316485)
reg316487 := PrimCons(reg316484, reg316486)
reg316488 := PrimSet(reg316483, reg316487)
Record := reg316488
_ = Record
reg316489 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4variancy_1test, V4123, V4124)
return
}, 0)
reg316491 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg316492 := MakeSymbol("shen.skip")
__ctx.Return(reg316492)
return
}, 1)
reg316493 := __e.Try(reg316489).Catch(reg316491)
Variancy := reg316493
_ = Variancy
reg316494 := __e.Call(__defun__shen_4demodulate, V4124)
reg316495 := __e.Call(__defun__shen_4rcons__form, reg316494)
Type := reg316495
_ = Type
reg316496 := MakeSymbol("shen.type-signature-of-")
reg316497 := __e.Call(__defun__concat, reg316496, V4123)
F_d := reg316497
_ = F_d
reg316498 := MakeNumber(1)
reg316499 := __e.Call(__defun__shen_4parameters, reg316498)
Parameters := reg316499
_ = Parameters
reg316500 := MakeSymbol("X")
reg316501 := Nil;
reg316502 := PrimCons(reg316500, reg316501)
reg316503 := PrimCons(F_d, reg316502)
reg316504 := MakeSymbol(":-")
reg316505 := MakeSymbol("unify!")
reg316506 := MakeSymbol("X")
reg316507 := Nil;
reg316508 := PrimCons(Type, reg316507)
reg316509 := PrimCons(reg316506, reg316508)
reg316510 := PrimCons(reg316505, reg316509)
reg316511 := Nil;
reg316512 := PrimCons(reg316510, reg316511)
reg316513 := Nil;
reg316514 := PrimCons(reg316512, reg316513)
reg316515 := PrimCons(reg316504, reg316514)
reg316516 := PrimCons(reg316503, reg316515)
Clause := reg316516
_ = Clause
reg316517 := __e.Call(__defun__shen_4aum, Clause, Parameters)
AUM__instruction := reg316517
_ = AUM__instruction
reg316518 := __e.Call(__defun__shen_4aum__to__shen, AUM__instruction)
Code := reg316518
_ = Code
reg316519 := MakeSymbol("define")
reg316520 := MakeSymbol("ProcessN")
reg316521 := MakeSymbol("Continuation")
reg316522 := Nil;
reg316523 := PrimCons(reg316521, reg316522)
reg316524 := PrimCons(reg316520, reg316523)
reg316525 := MakeSymbol("->")
reg316526 := Nil;
reg316527 := PrimCons(Code, reg316526)
reg316528 := PrimCons(reg316525, reg316527)
reg316529 := __e.Call(__defun__append, reg316524, reg316528)
reg316530 := __e.Call(__defun__append, Parameters, reg316529)
reg316531 := PrimCons(F_d, reg316530)
reg316532 := PrimCons(reg316519, reg316531)
ShenDef := reg316532
_ = ShenDef
reg316533 := __e.Call(__defun__shen_4eval_1without_1macros, ShenDef)
Eval := reg316533
_ = Eval
__ctx.Return(V4123)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "declare", value: __defun__declare})

__defun__shen_4demodulate = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4126 := __args[0]
_ = V4126
reg316534 := MakeSymbol("shen.*demodulation-function*")
reg316535 := PrimValue(reg316534)
reg316536 := __e.Call(__defun__shen_4walk, reg316535, V4126)
Demod := reg316536
_ = Demod
reg316537 := PrimEqual(Demod, V4126)
if reg316537 == True {
__ctx.Return(V4126)
return
} else {
__ctx.TailApply(__defun__shen_4demodulate, Demod)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.demodulate", value: __defun__shen_4demodulate})

__defun__shen_4variancy_1test = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4129 := __args[0]
_ = V4129
V4130 := __args[1]
_ = V4130
reg316539 := MakeSymbol("B")
reg316540 := __e.Call(__defun__shen_4typecheck, V4129, reg316539)
TypeF := reg316540
_ = TypeF
reg316541 := MakeSymbol("symbol")
reg316542 := PrimEqual(reg316541, TypeF)
var reg316554 Obj
if reg316542 == True {
reg316543 := MakeSymbol("shen.skip")
reg316554 = reg316543
} else {
reg316544 := __e.Call(__defun__shen_4variant_2, TypeF, V4130)
var reg316553 Obj
if reg316544 == True {
reg316545 := MakeSymbol("shen.skip")
reg316553 = reg316545
} else {
reg316546 := MakeString("warning: changing the type of ")
reg316547 := MakeString(" may create errors\n")
reg316548 := MakeSymbol("shen.a")
reg316549 := __e.Call(__defun__shen_4app, V4129, reg316547, reg316548)
reg316550 := PrimStringConcat(reg316546, reg316549)
reg316551 := __e.Call(__defun__stoutput)
reg316552 := __e.Call(__defun__shen_4prhush, reg316550, reg316551)
reg316553 = reg316552
}
reg316554 = reg316553
}
Check := reg316554
_ = Check
reg316555 := MakeSymbol("shen.skip")
__ctx.Return(reg316555)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.variancy-test", value: __defun__shen_4variancy_1test})

__defun__shen_4variant_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4143 := __args[0]
_ = V4143
V4144 := __args[1]
_ = V4144
reg316556 := PrimEqual(V4144, V4143)
if reg316556 == True {
reg316557 := True;
__ctx.Return(reg316557)
return
} else {
reg316558 := PrimIsPair(V4143)
var reg316572 Obj
if reg316558 == True {
reg316559 := PrimIsPair(V4144)
var reg316567 Obj
if reg316559 == True {
reg316560 := PrimHead(V4144)
reg316561 := PrimHead(V4143)
reg316562 := PrimEqual(reg316560, reg316561)
var reg316565 Obj
if reg316562 == True {
reg316563 := True;
reg316565 = reg316563
} else {
reg316564 := False;
reg316565 = reg316564
}
reg316567 = reg316565
} else {
reg316566 := False;
reg316567 = reg316566
}
var reg316570 Obj
if reg316567 == True {
reg316568 := True;
reg316570 = reg316568
} else {
reg316569 := False;
reg316570 = reg316569
}
reg316572 = reg316570
} else {
reg316571 := False;
reg316572 = reg316571
}
if reg316572 == True {
reg316573 := PrimTail(V4143)
reg316574 := PrimTail(V4144)
__ctx.TailApply(__defun__shen_4variant_2, reg316573, reg316574)
return
} else {
reg316576 := PrimIsPair(V4143)
var reg316596 Obj
if reg316576 == True {
reg316577 := PrimIsPair(V4144)
var reg316591 Obj
if reg316577 == True {
reg316578 := PrimHead(V4143)
reg316579 := __e.Call(__defun__shen_4pvar_2, reg316578)
var reg316586 Obj
if reg316579 == True {
reg316580 := PrimHead(V4144)
reg316581 := PrimIsVariable(reg316580)
var reg316584 Obj
if reg316581 == True {
reg316582 := True;
reg316584 = reg316582
} else {
reg316583 := False;
reg316584 = reg316583
}
reg316586 = reg316584
} else {
reg316585 := False;
reg316586 = reg316585
}
var reg316589 Obj
if reg316586 == True {
reg316587 := True;
reg316589 = reg316587
} else {
reg316588 := False;
reg316589 = reg316588
}
reg316591 = reg316589
} else {
reg316590 := False;
reg316591 = reg316590
}
var reg316594 Obj
if reg316591 == True {
reg316592 := True;
reg316594 = reg316592
} else {
reg316593 := False;
reg316594 = reg316593
}
reg316596 = reg316594
} else {
reg316595 := False;
reg316596 = reg316595
}
if reg316596 == True {
reg316597 := MakeSymbol("shen.a")
reg316598 := PrimHead(V4143)
reg316599 := PrimTail(V4143)
reg316600 := __e.Call(__defun__subst, reg316597, reg316598, reg316599)
reg316601 := MakeSymbol("shen.a")
reg316602 := PrimHead(V4144)
reg316603 := PrimTail(V4144)
reg316604 := __e.Call(__defun__subst, reg316601, reg316602, reg316603)
__ctx.TailApply(__defun__shen_4variant_2, reg316600, reg316604)
return
} else {
reg316606 := PrimIsPair(V4143)
var reg316626 Obj
if reg316606 == True {
reg316607 := PrimHead(V4143)
reg316608 := PrimIsPair(reg316607)
var reg316621 Obj
if reg316608 == True {
reg316609 := PrimIsPair(V4144)
var reg316616 Obj
if reg316609 == True {
reg316610 := PrimHead(V4144)
reg316611 := PrimIsPair(reg316610)
var reg316614 Obj
if reg316611 == True {
reg316612 := True;
reg316614 = reg316612
} else {
reg316613 := False;
reg316614 = reg316613
}
reg316616 = reg316614
} else {
reg316615 := False;
reg316616 = reg316615
}
var reg316619 Obj
if reg316616 == True {
reg316617 := True;
reg316619 = reg316617
} else {
reg316618 := False;
reg316619 = reg316618
}
reg316621 = reg316619
} else {
reg316620 := False;
reg316621 = reg316620
}
var reg316624 Obj
if reg316621 == True {
reg316622 := True;
reg316624 = reg316622
} else {
reg316623 := False;
reg316624 = reg316623
}
reg316626 = reg316624
} else {
reg316625 := False;
reg316626 = reg316625
}
if reg316626 == True {
reg316627 := PrimHead(V4143)
reg316628 := PrimTail(V4143)
reg316629 := __e.Call(__defun__append, reg316627, reg316628)
reg316630 := PrimHead(V4144)
reg316631 := PrimTail(V4144)
reg316632 := __e.Call(__defun__append, reg316630, reg316631)
__ctx.TailApply(__defun__shen_4variant_2, reg316629, reg316632)
return
} else {
reg316634 := False;
__ctx.Return(reg316634)
return
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.variant?", value: __defun__shen_4variant_2})

__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316635 := MakeSymbol("absvector?")
reg316636 := MakeSymbol("A")
reg316637 := MakeSymbol("-->")
reg316638 := MakeSymbol("boolean")
reg316639 := Nil;
reg316640 := PrimCons(reg316638, reg316639)
reg316641 := PrimCons(reg316637, reg316640)
reg316642 := PrimCons(reg316636, reg316641)
__ctx.TailApply(__defun__declare, reg316635, reg316642)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316644 := MakeSymbol("adjoin")
reg316645 := MakeSymbol("A")
reg316646 := MakeSymbol("-->")
reg316647 := MakeSymbol("list")
reg316648 := MakeSymbol("A")
reg316649 := Nil;
reg316650 := PrimCons(reg316648, reg316649)
reg316651 := PrimCons(reg316647, reg316650)
reg316652 := MakeSymbol("-->")
reg316653 := MakeSymbol("list")
reg316654 := MakeSymbol("A")
reg316655 := Nil;
reg316656 := PrimCons(reg316654, reg316655)
reg316657 := PrimCons(reg316653, reg316656)
reg316658 := Nil;
reg316659 := PrimCons(reg316657, reg316658)
reg316660 := PrimCons(reg316652, reg316659)
reg316661 := PrimCons(reg316651, reg316660)
reg316662 := Nil;
reg316663 := PrimCons(reg316661, reg316662)
reg316664 := PrimCons(reg316646, reg316663)
reg316665 := PrimCons(reg316645, reg316664)
__ctx.TailApply(__defun__declare, reg316644, reg316665)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316667 := MakeSymbol("and")
reg316668 := MakeSymbol("boolean")
reg316669 := MakeSymbol("-->")
reg316670 := MakeSymbol("boolean")
reg316671 := MakeSymbol("-->")
reg316672 := MakeSymbol("boolean")
reg316673 := Nil;
reg316674 := PrimCons(reg316672, reg316673)
reg316675 := PrimCons(reg316671, reg316674)
reg316676 := PrimCons(reg316670, reg316675)
reg316677 := Nil;
reg316678 := PrimCons(reg316676, reg316677)
reg316679 := PrimCons(reg316669, reg316678)
reg316680 := PrimCons(reg316668, reg316679)
__ctx.TailApply(__defun__declare, reg316667, reg316680)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316682 := MakeSymbol("shen.app")
reg316683 := MakeSymbol("A")
reg316684 := MakeSymbol("-->")
reg316685 := MakeSymbol("string")
reg316686 := MakeSymbol("-->")
reg316687 := MakeSymbol("symbol")
reg316688 := MakeSymbol("-->")
reg316689 := MakeSymbol("string")
reg316690 := Nil;
reg316691 := PrimCons(reg316689, reg316690)
reg316692 := PrimCons(reg316688, reg316691)
reg316693 := PrimCons(reg316687, reg316692)
reg316694 := Nil;
reg316695 := PrimCons(reg316693, reg316694)
reg316696 := PrimCons(reg316686, reg316695)
reg316697 := PrimCons(reg316685, reg316696)
reg316698 := Nil;
reg316699 := PrimCons(reg316697, reg316698)
reg316700 := PrimCons(reg316684, reg316699)
reg316701 := PrimCons(reg316683, reg316700)
__ctx.TailApply(__defun__declare, reg316682, reg316701)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316703 := MakeSymbol("append")
reg316704 := MakeSymbol("list")
reg316705 := MakeSymbol("A")
reg316706 := Nil;
reg316707 := PrimCons(reg316705, reg316706)
reg316708 := PrimCons(reg316704, reg316707)
reg316709 := MakeSymbol("-->")
reg316710 := MakeSymbol("list")
reg316711 := MakeSymbol("A")
reg316712 := Nil;
reg316713 := PrimCons(reg316711, reg316712)
reg316714 := PrimCons(reg316710, reg316713)
reg316715 := MakeSymbol("-->")
reg316716 := MakeSymbol("list")
reg316717 := MakeSymbol("A")
reg316718 := Nil;
reg316719 := PrimCons(reg316717, reg316718)
reg316720 := PrimCons(reg316716, reg316719)
reg316721 := Nil;
reg316722 := PrimCons(reg316720, reg316721)
reg316723 := PrimCons(reg316715, reg316722)
reg316724 := PrimCons(reg316714, reg316723)
reg316725 := Nil;
reg316726 := PrimCons(reg316724, reg316725)
reg316727 := PrimCons(reg316709, reg316726)
reg316728 := PrimCons(reg316708, reg316727)
__ctx.TailApply(__defun__declare, reg316703, reg316728)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316730 := MakeSymbol("arity")
reg316731 := MakeSymbol("A")
reg316732 := MakeSymbol("-->")
reg316733 := MakeSymbol("number")
reg316734 := Nil;
reg316735 := PrimCons(reg316733, reg316734)
reg316736 := PrimCons(reg316732, reg316735)
reg316737 := PrimCons(reg316731, reg316736)
__ctx.TailApply(__defun__declare, reg316730, reg316737)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316739 := MakeSymbol("assoc")
reg316740 := MakeSymbol("A")
reg316741 := MakeSymbol("-->")
reg316742 := MakeSymbol("list")
reg316743 := MakeSymbol("list")
reg316744 := MakeSymbol("A")
reg316745 := Nil;
reg316746 := PrimCons(reg316744, reg316745)
reg316747 := PrimCons(reg316743, reg316746)
reg316748 := Nil;
reg316749 := PrimCons(reg316747, reg316748)
reg316750 := PrimCons(reg316742, reg316749)
reg316751 := MakeSymbol("-->")
reg316752 := MakeSymbol("list")
reg316753 := MakeSymbol("A")
reg316754 := Nil;
reg316755 := PrimCons(reg316753, reg316754)
reg316756 := PrimCons(reg316752, reg316755)
reg316757 := Nil;
reg316758 := PrimCons(reg316756, reg316757)
reg316759 := PrimCons(reg316751, reg316758)
reg316760 := PrimCons(reg316750, reg316759)
reg316761 := Nil;
reg316762 := PrimCons(reg316760, reg316761)
reg316763 := PrimCons(reg316741, reg316762)
reg316764 := PrimCons(reg316740, reg316763)
__ctx.TailApply(__defun__declare, reg316739, reg316764)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316766 := MakeSymbol("boolean?")
reg316767 := MakeSymbol("A")
reg316768 := MakeSymbol("-->")
reg316769 := MakeSymbol("boolean")
reg316770 := Nil;
reg316771 := PrimCons(reg316769, reg316770)
reg316772 := PrimCons(reg316768, reg316771)
reg316773 := PrimCons(reg316767, reg316772)
__ctx.TailApply(__defun__declare, reg316766, reg316773)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316775 := MakeSymbol("bound?")
reg316776 := MakeSymbol("symbol")
reg316777 := MakeSymbol("-->")
reg316778 := MakeSymbol("boolean")
reg316779 := Nil;
reg316780 := PrimCons(reg316778, reg316779)
reg316781 := PrimCons(reg316777, reg316780)
reg316782 := PrimCons(reg316776, reg316781)
__ctx.TailApply(__defun__declare, reg316775, reg316782)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316784 := MakeSymbol("cd")
reg316785 := MakeSymbol("string")
reg316786 := MakeSymbol("-->")
reg316787 := MakeSymbol("string")
reg316788 := Nil;
reg316789 := PrimCons(reg316787, reg316788)
reg316790 := PrimCons(reg316786, reg316789)
reg316791 := PrimCons(reg316785, reg316790)
__ctx.TailApply(__defun__declare, reg316784, reg316791)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316793 := MakeSymbol("close")
reg316794 := MakeSymbol("stream")
reg316795 := MakeSymbol("A")
reg316796 := Nil;
reg316797 := PrimCons(reg316795, reg316796)
reg316798 := PrimCons(reg316794, reg316797)
reg316799 := MakeSymbol("-->")
reg316800 := MakeSymbol("list")
reg316801 := MakeSymbol("B")
reg316802 := Nil;
reg316803 := PrimCons(reg316801, reg316802)
reg316804 := PrimCons(reg316800, reg316803)
reg316805 := Nil;
reg316806 := PrimCons(reg316804, reg316805)
reg316807 := PrimCons(reg316799, reg316806)
reg316808 := PrimCons(reg316798, reg316807)
__ctx.TailApply(__defun__declare, reg316793, reg316808)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316810 := MakeSymbol("cn")
reg316811 := MakeSymbol("string")
reg316812 := MakeSymbol("-->")
reg316813 := MakeSymbol("string")
reg316814 := MakeSymbol("-->")
reg316815 := MakeSymbol("string")
reg316816 := Nil;
reg316817 := PrimCons(reg316815, reg316816)
reg316818 := PrimCons(reg316814, reg316817)
reg316819 := PrimCons(reg316813, reg316818)
reg316820 := Nil;
reg316821 := PrimCons(reg316819, reg316820)
reg316822 := PrimCons(reg316812, reg316821)
reg316823 := PrimCons(reg316811, reg316822)
__ctx.TailApply(__defun__declare, reg316810, reg316823)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316825 := MakeSymbol("compile")
reg316826 := MakeSymbol("A")
reg316827 := MakeSymbol("shen.==>")
reg316828 := MakeSymbol("B")
reg316829 := Nil;
reg316830 := PrimCons(reg316828, reg316829)
reg316831 := PrimCons(reg316827, reg316830)
reg316832 := PrimCons(reg316826, reg316831)
reg316833 := MakeSymbol("-->")
reg316834 := MakeSymbol("A")
reg316835 := MakeSymbol("-->")
reg316836 := MakeSymbol("A")
reg316837 := MakeSymbol("-->")
reg316838 := MakeSymbol("B")
reg316839 := Nil;
reg316840 := PrimCons(reg316838, reg316839)
reg316841 := PrimCons(reg316837, reg316840)
reg316842 := PrimCons(reg316836, reg316841)
reg316843 := MakeSymbol("-->")
reg316844 := MakeSymbol("B")
reg316845 := Nil;
reg316846 := PrimCons(reg316844, reg316845)
reg316847 := PrimCons(reg316843, reg316846)
reg316848 := PrimCons(reg316842, reg316847)
reg316849 := Nil;
reg316850 := PrimCons(reg316848, reg316849)
reg316851 := PrimCons(reg316835, reg316850)
reg316852 := PrimCons(reg316834, reg316851)
reg316853 := Nil;
reg316854 := PrimCons(reg316852, reg316853)
reg316855 := PrimCons(reg316833, reg316854)
reg316856 := PrimCons(reg316832, reg316855)
__ctx.TailApply(__defun__declare, reg316825, reg316856)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316858 := MakeSymbol("cons?")
reg316859 := MakeSymbol("A")
reg316860 := MakeSymbol("-->")
reg316861 := MakeSymbol("boolean")
reg316862 := Nil;
reg316863 := PrimCons(reg316861, reg316862)
reg316864 := PrimCons(reg316860, reg316863)
reg316865 := PrimCons(reg316859, reg316864)
__ctx.TailApply(__defun__declare, reg316858, reg316865)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316867 := MakeSymbol("destroy")
reg316868 := MakeSymbol("A")
reg316869 := MakeSymbol("-->")
reg316870 := MakeSymbol("B")
reg316871 := Nil;
reg316872 := PrimCons(reg316870, reg316871)
reg316873 := PrimCons(reg316869, reg316872)
reg316874 := PrimCons(reg316868, reg316873)
reg316875 := MakeSymbol("-->")
reg316876 := MakeSymbol("symbol")
reg316877 := Nil;
reg316878 := PrimCons(reg316876, reg316877)
reg316879 := PrimCons(reg316875, reg316878)
reg316880 := PrimCons(reg316874, reg316879)
__ctx.TailApply(__defun__declare, reg316867, reg316880)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316882 := MakeSymbol("difference")
reg316883 := MakeSymbol("list")
reg316884 := MakeSymbol("A")
reg316885 := Nil;
reg316886 := PrimCons(reg316884, reg316885)
reg316887 := PrimCons(reg316883, reg316886)
reg316888 := MakeSymbol("-->")
reg316889 := MakeSymbol("list")
reg316890 := MakeSymbol("A")
reg316891 := Nil;
reg316892 := PrimCons(reg316890, reg316891)
reg316893 := PrimCons(reg316889, reg316892)
reg316894 := MakeSymbol("-->")
reg316895 := MakeSymbol("list")
reg316896 := MakeSymbol("A")
reg316897 := Nil;
reg316898 := PrimCons(reg316896, reg316897)
reg316899 := PrimCons(reg316895, reg316898)
reg316900 := Nil;
reg316901 := PrimCons(reg316899, reg316900)
reg316902 := PrimCons(reg316894, reg316901)
reg316903 := PrimCons(reg316893, reg316902)
reg316904 := Nil;
reg316905 := PrimCons(reg316903, reg316904)
reg316906 := PrimCons(reg316888, reg316905)
reg316907 := PrimCons(reg316887, reg316906)
__ctx.TailApply(__defun__declare, reg316882, reg316907)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316909 := MakeSymbol("do")
reg316910 := MakeSymbol("A")
reg316911 := MakeSymbol("-->")
reg316912 := MakeSymbol("B")
reg316913 := MakeSymbol("-->")
reg316914 := MakeSymbol("B")
reg316915 := Nil;
reg316916 := PrimCons(reg316914, reg316915)
reg316917 := PrimCons(reg316913, reg316916)
reg316918 := PrimCons(reg316912, reg316917)
reg316919 := Nil;
reg316920 := PrimCons(reg316918, reg316919)
reg316921 := PrimCons(reg316911, reg316920)
reg316922 := PrimCons(reg316910, reg316921)
__ctx.TailApply(__defun__declare, reg316909, reg316922)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316924 := MakeSymbol("<e>")
reg316925 := MakeSymbol("list")
reg316926 := MakeSymbol("A")
reg316927 := Nil;
reg316928 := PrimCons(reg316926, reg316927)
reg316929 := PrimCons(reg316925, reg316928)
reg316930 := MakeSymbol("shen.==>")
reg316931 := MakeSymbol("list")
reg316932 := MakeSymbol("B")
reg316933 := Nil;
reg316934 := PrimCons(reg316932, reg316933)
reg316935 := PrimCons(reg316931, reg316934)
reg316936 := Nil;
reg316937 := PrimCons(reg316935, reg316936)
reg316938 := PrimCons(reg316930, reg316937)
reg316939 := PrimCons(reg316929, reg316938)
__ctx.TailApply(__defun__declare, reg316924, reg316939)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316941 := MakeSymbol("<!>")
reg316942 := MakeSymbol("list")
reg316943 := MakeSymbol("A")
reg316944 := Nil;
reg316945 := PrimCons(reg316943, reg316944)
reg316946 := PrimCons(reg316942, reg316945)
reg316947 := MakeSymbol("shen.==>")
reg316948 := MakeSymbol("list")
reg316949 := MakeSymbol("A")
reg316950 := Nil;
reg316951 := PrimCons(reg316949, reg316950)
reg316952 := PrimCons(reg316948, reg316951)
reg316953 := Nil;
reg316954 := PrimCons(reg316952, reg316953)
reg316955 := PrimCons(reg316947, reg316954)
reg316956 := PrimCons(reg316946, reg316955)
__ctx.TailApply(__defun__declare, reg316941, reg316956)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316958 := MakeSymbol("element?")
reg316959 := MakeSymbol("A")
reg316960 := MakeSymbol("-->")
reg316961 := MakeSymbol("list")
reg316962 := MakeSymbol("A")
reg316963 := Nil;
reg316964 := PrimCons(reg316962, reg316963)
reg316965 := PrimCons(reg316961, reg316964)
reg316966 := MakeSymbol("-->")
reg316967 := MakeSymbol("boolean")
reg316968 := Nil;
reg316969 := PrimCons(reg316967, reg316968)
reg316970 := PrimCons(reg316966, reg316969)
reg316971 := PrimCons(reg316965, reg316970)
reg316972 := Nil;
reg316973 := PrimCons(reg316971, reg316972)
reg316974 := PrimCons(reg316960, reg316973)
reg316975 := PrimCons(reg316959, reg316974)
__ctx.TailApply(__defun__declare, reg316958, reg316975)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316977 := MakeSymbol("empty?")
reg316978 := MakeSymbol("A")
reg316979 := MakeSymbol("-->")
reg316980 := MakeSymbol("boolean")
reg316981 := Nil;
reg316982 := PrimCons(reg316980, reg316981)
reg316983 := PrimCons(reg316979, reg316982)
reg316984 := PrimCons(reg316978, reg316983)
__ctx.TailApply(__defun__declare, reg316977, reg316984)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316986 := MakeSymbol("enable-type-theory")
reg316987 := MakeSymbol("symbol")
reg316988 := MakeSymbol("-->")
reg316989 := MakeSymbol("boolean")
reg316990 := Nil;
reg316991 := PrimCons(reg316989, reg316990)
reg316992 := PrimCons(reg316988, reg316991)
reg316993 := PrimCons(reg316987, reg316992)
__ctx.TailApply(__defun__declare, reg316986, reg316993)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg316995 := MakeSymbol("external")
reg316996 := MakeSymbol("symbol")
reg316997 := MakeSymbol("-->")
reg316998 := MakeSymbol("list")
reg316999 := MakeSymbol("symbol")
reg317000 := Nil;
reg317001 := PrimCons(reg316999, reg317000)
reg317002 := PrimCons(reg316998, reg317001)
reg317003 := Nil;
reg317004 := PrimCons(reg317002, reg317003)
reg317005 := PrimCons(reg316997, reg317004)
reg317006 := PrimCons(reg316996, reg317005)
__ctx.TailApply(__defun__declare, reg316995, reg317006)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317008 := MakeSymbol("error-to-string")
reg317009 := MakeSymbol("exception")
reg317010 := MakeSymbol("-->")
reg317011 := MakeSymbol("string")
reg317012 := Nil;
reg317013 := PrimCons(reg317011, reg317012)
reg317014 := PrimCons(reg317010, reg317013)
reg317015 := PrimCons(reg317009, reg317014)
__ctx.TailApply(__defun__declare, reg317008, reg317015)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317017 := MakeSymbol("explode")
reg317018 := MakeSymbol("A")
reg317019 := MakeSymbol("-->")
reg317020 := MakeSymbol("list")
reg317021 := MakeSymbol("string")
reg317022 := Nil;
reg317023 := PrimCons(reg317021, reg317022)
reg317024 := PrimCons(reg317020, reg317023)
reg317025 := Nil;
reg317026 := PrimCons(reg317024, reg317025)
reg317027 := PrimCons(reg317019, reg317026)
reg317028 := PrimCons(reg317018, reg317027)
__ctx.TailApply(__defun__declare, reg317017, reg317028)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317030 := MakeSymbol("fail")
reg317031 := MakeSymbol("-->")
reg317032 := MakeSymbol("symbol")
reg317033 := Nil;
reg317034 := PrimCons(reg317032, reg317033)
reg317035 := PrimCons(reg317031, reg317034)
__ctx.TailApply(__defun__declare, reg317030, reg317035)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317037 := MakeSymbol("fail-if")
reg317038 := MakeSymbol("symbol")
reg317039 := MakeSymbol("-->")
reg317040 := MakeSymbol("boolean")
reg317041 := Nil;
reg317042 := PrimCons(reg317040, reg317041)
reg317043 := PrimCons(reg317039, reg317042)
reg317044 := PrimCons(reg317038, reg317043)
reg317045 := MakeSymbol("-->")
reg317046 := MakeSymbol("symbol")
reg317047 := MakeSymbol("-->")
reg317048 := MakeSymbol("symbol")
reg317049 := Nil;
reg317050 := PrimCons(reg317048, reg317049)
reg317051 := PrimCons(reg317047, reg317050)
reg317052 := PrimCons(reg317046, reg317051)
reg317053 := Nil;
reg317054 := PrimCons(reg317052, reg317053)
reg317055 := PrimCons(reg317045, reg317054)
reg317056 := PrimCons(reg317044, reg317055)
__ctx.TailApply(__defun__declare, reg317037, reg317056)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317058 := MakeSymbol("fix")
reg317059 := MakeSymbol("A")
reg317060 := MakeSymbol("-->")
reg317061 := MakeSymbol("A")
reg317062 := Nil;
reg317063 := PrimCons(reg317061, reg317062)
reg317064 := PrimCons(reg317060, reg317063)
reg317065 := PrimCons(reg317059, reg317064)
reg317066 := MakeSymbol("-->")
reg317067 := MakeSymbol("A")
reg317068 := MakeSymbol("-->")
reg317069 := MakeSymbol("A")
reg317070 := Nil;
reg317071 := PrimCons(reg317069, reg317070)
reg317072 := PrimCons(reg317068, reg317071)
reg317073 := PrimCons(reg317067, reg317072)
reg317074 := Nil;
reg317075 := PrimCons(reg317073, reg317074)
reg317076 := PrimCons(reg317066, reg317075)
reg317077 := PrimCons(reg317065, reg317076)
__ctx.TailApply(__defun__declare, reg317058, reg317077)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317079 := MakeSymbol("freeze")
reg317080 := MakeSymbol("A")
reg317081 := MakeSymbol("-->")
reg317082 := MakeSymbol("lazy")
reg317083 := MakeSymbol("A")
reg317084 := Nil;
reg317085 := PrimCons(reg317083, reg317084)
reg317086 := PrimCons(reg317082, reg317085)
reg317087 := Nil;
reg317088 := PrimCons(reg317086, reg317087)
reg317089 := PrimCons(reg317081, reg317088)
reg317090 := PrimCons(reg317080, reg317089)
__ctx.TailApply(__defun__declare, reg317079, reg317090)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317092 := MakeSymbol("fst")
reg317093 := MakeSymbol("A")
reg317094 := MakeSymbol("*")
reg317095 := MakeSymbol("B")
reg317096 := Nil;
reg317097 := PrimCons(reg317095, reg317096)
reg317098 := PrimCons(reg317094, reg317097)
reg317099 := PrimCons(reg317093, reg317098)
reg317100 := MakeSymbol("-->")
reg317101 := MakeSymbol("A")
reg317102 := Nil;
reg317103 := PrimCons(reg317101, reg317102)
reg317104 := PrimCons(reg317100, reg317103)
reg317105 := PrimCons(reg317099, reg317104)
__ctx.TailApply(__defun__declare, reg317092, reg317105)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317107 := MakeSymbol("function")
reg317108 := MakeSymbol("A")
reg317109 := MakeSymbol("-->")
reg317110 := MakeSymbol("B")
reg317111 := Nil;
reg317112 := PrimCons(reg317110, reg317111)
reg317113 := PrimCons(reg317109, reg317112)
reg317114 := PrimCons(reg317108, reg317113)
reg317115 := MakeSymbol("-->")
reg317116 := MakeSymbol("A")
reg317117 := MakeSymbol("-->")
reg317118 := MakeSymbol("B")
reg317119 := Nil;
reg317120 := PrimCons(reg317118, reg317119)
reg317121 := PrimCons(reg317117, reg317120)
reg317122 := PrimCons(reg317116, reg317121)
reg317123 := Nil;
reg317124 := PrimCons(reg317122, reg317123)
reg317125 := PrimCons(reg317115, reg317124)
reg317126 := PrimCons(reg317114, reg317125)
__ctx.TailApply(__defun__declare, reg317107, reg317126)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317128 := MakeSymbol("gensym")
reg317129 := MakeSymbol("symbol")
reg317130 := MakeSymbol("-->")
reg317131 := MakeSymbol("symbol")
reg317132 := Nil;
reg317133 := PrimCons(reg317131, reg317132)
reg317134 := PrimCons(reg317130, reg317133)
reg317135 := PrimCons(reg317129, reg317134)
__ctx.TailApply(__defun__declare, reg317128, reg317135)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317137 := MakeSymbol("<-vector")
reg317138 := MakeSymbol("vector")
reg317139 := MakeSymbol("A")
reg317140 := Nil;
reg317141 := PrimCons(reg317139, reg317140)
reg317142 := PrimCons(reg317138, reg317141)
reg317143 := MakeSymbol("-->")
reg317144 := MakeSymbol("number")
reg317145 := MakeSymbol("-->")
reg317146 := MakeSymbol("A")
reg317147 := Nil;
reg317148 := PrimCons(reg317146, reg317147)
reg317149 := PrimCons(reg317145, reg317148)
reg317150 := PrimCons(reg317144, reg317149)
reg317151 := Nil;
reg317152 := PrimCons(reg317150, reg317151)
reg317153 := PrimCons(reg317143, reg317152)
reg317154 := PrimCons(reg317142, reg317153)
__ctx.TailApply(__defun__declare, reg317137, reg317154)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317156 := MakeSymbol("vector->")
reg317157 := MakeSymbol("vector")
reg317158 := MakeSymbol("A")
reg317159 := Nil;
reg317160 := PrimCons(reg317158, reg317159)
reg317161 := PrimCons(reg317157, reg317160)
reg317162 := MakeSymbol("-->")
reg317163 := MakeSymbol("number")
reg317164 := MakeSymbol("-->")
reg317165 := MakeSymbol("A")
reg317166 := MakeSymbol("-->")
reg317167 := MakeSymbol("vector")
reg317168 := MakeSymbol("A")
reg317169 := Nil;
reg317170 := PrimCons(reg317168, reg317169)
reg317171 := PrimCons(reg317167, reg317170)
reg317172 := Nil;
reg317173 := PrimCons(reg317171, reg317172)
reg317174 := PrimCons(reg317166, reg317173)
reg317175 := PrimCons(reg317165, reg317174)
reg317176 := Nil;
reg317177 := PrimCons(reg317175, reg317176)
reg317178 := PrimCons(reg317164, reg317177)
reg317179 := PrimCons(reg317163, reg317178)
reg317180 := Nil;
reg317181 := PrimCons(reg317179, reg317180)
reg317182 := PrimCons(reg317162, reg317181)
reg317183 := PrimCons(reg317161, reg317182)
__ctx.TailApply(__defun__declare, reg317156, reg317183)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317185 := MakeSymbol("vector")
reg317186 := MakeSymbol("number")
reg317187 := MakeSymbol("-->")
reg317188 := MakeSymbol("vector")
reg317189 := MakeSymbol("A")
reg317190 := Nil;
reg317191 := PrimCons(reg317189, reg317190)
reg317192 := PrimCons(reg317188, reg317191)
reg317193 := Nil;
reg317194 := PrimCons(reg317192, reg317193)
reg317195 := PrimCons(reg317187, reg317194)
reg317196 := PrimCons(reg317186, reg317195)
__ctx.TailApply(__defun__declare, reg317185, reg317196)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317198 := MakeSymbol("get-time")
reg317199 := MakeSymbol("symbol")
reg317200 := MakeSymbol("-->")
reg317201 := MakeSymbol("number")
reg317202 := Nil;
reg317203 := PrimCons(reg317201, reg317202)
reg317204 := PrimCons(reg317200, reg317203)
reg317205 := PrimCons(reg317199, reg317204)
__ctx.TailApply(__defun__declare, reg317198, reg317205)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317207 := MakeSymbol("hash")
reg317208 := MakeSymbol("A")
reg317209 := MakeSymbol("-->")
reg317210 := MakeSymbol("number")
reg317211 := MakeSymbol("-->")
reg317212 := MakeSymbol("number")
reg317213 := Nil;
reg317214 := PrimCons(reg317212, reg317213)
reg317215 := PrimCons(reg317211, reg317214)
reg317216 := PrimCons(reg317210, reg317215)
reg317217 := Nil;
reg317218 := PrimCons(reg317216, reg317217)
reg317219 := PrimCons(reg317209, reg317218)
reg317220 := PrimCons(reg317208, reg317219)
__ctx.TailApply(__defun__declare, reg317207, reg317220)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317222 := MakeSymbol("head")
reg317223 := MakeSymbol("list")
reg317224 := MakeSymbol("A")
reg317225 := Nil;
reg317226 := PrimCons(reg317224, reg317225)
reg317227 := PrimCons(reg317223, reg317226)
reg317228 := MakeSymbol("-->")
reg317229 := MakeSymbol("A")
reg317230 := Nil;
reg317231 := PrimCons(reg317229, reg317230)
reg317232 := PrimCons(reg317228, reg317231)
reg317233 := PrimCons(reg317227, reg317232)
__ctx.TailApply(__defun__declare, reg317222, reg317233)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317235 := MakeSymbol("hdv")
reg317236 := MakeSymbol("vector")
reg317237 := MakeSymbol("A")
reg317238 := Nil;
reg317239 := PrimCons(reg317237, reg317238)
reg317240 := PrimCons(reg317236, reg317239)
reg317241 := MakeSymbol("-->")
reg317242 := MakeSymbol("A")
reg317243 := Nil;
reg317244 := PrimCons(reg317242, reg317243)
reg317245 := PrimCons(reg317241, reg317244)
reg317246 := PrimCons(reg317240, reg317245)
__ctx.TailApply(__defun__declare, reg317235, reg317246)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317248 := MakeSymbol("hdstr")
reg317249 := MakeSymbol("string")
reg317250 := MakeSymbol("-->")
reg317251 := MakeSymbol("string")
reg317252 := Nil;
reg317253 := PrimCons(reg317251, reg317252)
reg317254 := PrimCons(reg317250, reg317253)
reg317255 := PrimCons(reg317249, reg317254)
__ctx.TailApply(__defun__declare, reg317248, reg317255)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317257 := MakeSymbol("if")
reg317258 := MakeSymbol("boolean")
reg317259 := MakeSymbol("-->")
reg317260 := MakeSymbol("A")
reg317261 := MakeSymbol("-->")
reg317262 := MakeSymbol("A")
reg317263 := MakeSymbol("-->")
reg317264 := MakeSymbol("A")
reg317265 := Nil;
reg317266 := PrimCons(reg317264, reg317265)
reg317267 := PrimCons(reg317263, reg317266)
reg317268 := PrimCons(reg317262, reg317267)
reg317269 := Nil;
reg317270 := PrimCons(reg317268, reg317269)
reg317271 := PrimCons(reg317261, reg317270)
reg317272 := PrimCons(reg317260, reg317271)
reg317273 := Nil;
reg317274 := PrimCons(reg317272, reg317273)
reg317275 := PrimCons(reg317259, reg317274)
reg317276 := PrimCons(reg317258, reg317275)
__ctx.TailApply(__defun__declare, reg317257, reg317276)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317278 := MakeSymbol("it")
reg317279 := MakeSymbol("-->")
reg317280 := MakeSymbol("string")
reg317281 := Nil;
reg317282 := PrimCons(reg317280, reg317281)
reg317283 := PrimCons(reg317279, reg317282)
__ctx.TailApply(__defun__declare, reg317278, reg317283)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317285 := MakeSymbol("implementation")
reg317286 := MakeSymbol("-->")
reg317287 := MakeSymbol("string")
reg317288 := Nil;
reg317289 := PrimCons(reg317287, reg317288)
reg317290 := PrimCons(reg317286, reg317289)
__ctx.TailApply(__defun__declare, reg317285, reg317290)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317292 := MakeSymbol("include")
reg317293 := MakeSymbol("list")
reg317294 := MakeSymbol("symbol")
reg317295 := Nil;
reg317296 := PrimCons(reg317294, reg317295)
reg317297 := PrimCons(reg317293, reg317296)
reg317298 := MakeSymbol("-->")
reg317299 := MakeSymbol("list")
reg317300 := MakeSymbol("symbol")
reg317301 := Nil;
reg317302 := PrimCons(reg317300, reg317301)
reg317303 := PrimCons(reg317299, reg317302)
reg317304 := Nil;
reg317305 := PrimCons(reg317303, reg317304)
reg317306 := PrimCons(reg317298, reg317305)
reg317307 := PrimCons(reg317297, reg317306)
__ctx.TailApply(__defun__declare, reg317292, reg317307)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317309 := MakeSymbol("include-all-but")
reg317310 := MakeSymbol("list")
reg317311 := MakeSymbol("symbol")
reg317312 := Nil;
reg317313 := PrimCons(reg317311, reg317312)
reg317314 := PrimCons(reg317310, reg317313)
reg317315 := MakeSymbol("-->")
reg317316 := MakeSymbol("list")
reg317317 := MakeSymbol("symbol")
reg317318 := Nil;
reg317319 := PrimCons(reg317317, reg317318)
reg317320 := PrimCons(reg317316, reg317319)
reg317321 := Nil;
reg317322 := PrimCons(reg317320, reg317321)
reg317323 := PrimCons(reg317315, reg317322)
reg317324 := PrimCons(reg317314, reg317323)
__ctx.TailApply(__defun__declare, reg317309, reg317324)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317326 := MakeSymbol("inferences")
reg317327 := MakeSymbol("-->")
reg317328 := MakeSymbol("number")
reg317329 := Nil;
reg317330 := PrimCons(reg317328, reg317329)
reg317331 := PrimCons(reg317327, reg317330)
__ctx.TailApply(__defun__declare, reg317326, reg317331)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317333 := MakeSymbol("shen.insert")
reg317334 := MakeSymbol("A")
reg317335 := MakeSymbol("-->")
reg317336 := MakeSymbol("string")
reg317337 := MakeSymbol("-->")
reg317338 := MakeSymbol("string")
reg317339 := Nil;
reg317340 := PrimCons(reg317338, reg317339)
reg317341 := PrimCons(reg317337, reg317340)
reg317342 := PrimCons(reg317336, reg317341)
reg317343 := Nil;
reg317344 := PrimCons(reg317342, reg317343)
reg317345 := PrimCons(reg317335, reg317344)
reg317346 := PrimCons(reg317334, reg317345)
__ctx.TailApply(__defun__declare, reg317333, reg317346)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317348 := MakeSymbol("integer?")
reg317349 := MakeSymbol("A")
reg317350 := MakeSymbol("-->")
reg317351 := MakeSymbol("boolean")
reg317352 := Nil;
reg317353 := PrimCons(reg317351, reg317352)
reg317354 := PrimCons(reg317350, reg317353)
reg317355 := PrimCons(reg317349, reg317354)
__ctx.TailApply(__defun__declare, reg317348, reg317355)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317357 := MakeSymbol("internal")
reg317358 := MakeSymbol("symbol")
reg317359 := MakeSymbol("-->")
reg317360 := MakeSymbol("list")
reg317361 := MakeSymbol("symbol")
reg317362 := Nil;
reg317363 := PrimCons(reg317361, reg317362)
reg317364 := PrimCons(reg317360, reg317363)
reg317365 := Nil;
reg317366 := PrimCons(reg317364, reg317365)
reg317367 := PrimCons(reg317359, reg317366)
reg317368 := PrimCons(reg317358, reg317367)
__ctx.TailApply(__defun__declare, reg317357, reg317368)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317370 := MakeSymbol("intersection")
reg317371 := MakeSymbol("list")
reg317372 := MakeSymbol("A")
reg317373 := Nil;
reg317374 := PrimCons(reg317372, reg317373)
reg317375 := PrimCons(reg317371, reg317374)
reg317376 := MakeSymbol("-->")
reg317377 := MakeSymbol("list")
reg317378 := MakeSymbol("A")
reg317379 := Nil;
reg317380 := PrimCons(reg317378, reg317379)
reg317381 := PrimCons(reg317377, reg317380)
reg317382 := MakeSymbol("-->")
reg317383 := MakeSymbol("list")
reg317384 := MakeSymbol("A")
reg317385 := Nil;
reg317386 := PrimCons(reg317384, reg317385)
reg317387 := PrimCons(reg317383, reg317386)
reg317388 := Nil;
reg317389 := PrimCons(reg317387, reg317388)
reg317390 := PrimCons(reg317382, reg317389)
reg317391 := PrimCons(reg317381, reg317390)
reg317392 := Nil;
reg317393 := PrimCons(reg317391, reg317392)
reg317394 := PrimCons(reg317376, reg317393)
reg317395 := PrimCons(reg317375, reg317394)
__ctx.TailApply(__defun__declare, reg317370, reg317395)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317397 := MakeSymbol("kill")
reg317398 := MakeSymbol("-->")
reg317399 := MakeSymbol("A")
reg317400 := Nil;
reg317401 := PrimCons(reg317399, reg317400)
reg317402 := PrimCons(reg317398, reg317401)
__ctx.TailApply(__defun__declare, reg317397, reg317402)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317404 := MakeSymbol("language")
reg317405 := MakeSymbol("-->")
reg317406 := MakeSymbol("string")
reg317407 := Nil;
reg317408 := PrimCons(reg317406, reg317407)
reg317409 := PrimCons(reg317405, reg317408)
__ctx.TailApply(__defun__declare, reg317404, reg317409)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317411 := MakeSymbol("length")
reg317412 := MakeSymbol("list")
reg317413 := MakeSymbol("A")
reg317414 := Nil;
reg317415 := PrimCons(reg317413, reg317414)
reg317416 := PrimCons(reg317412, reg317415)
reg317417 := MakeSymbol("-->")
reg317418 := MakeSymbol("number")
reg317419 := Nil;
reg317420 := PrimCons(reg317418, reg317419)
reg317421 := PrimCons(reg317417, reg317420)
reg317422 := PrimCons(reg317416, reg317421)
__ctx.TailApply(__defun__declare, reg317411, reg317422)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317424 := MakeSymbol("limit")
reg317425 := MakeSymbol("vector")
reg317426 := MakeSymbol("A")
reg317427 := Nil;
reg317428 := PrimCons(reg317426, reg317427)
reg317429 := PrimCons(reg317425, reg317428)
reg317430 := MakeSymbol("-->")
reg317431 := MakeSymbol("number")
reg317432 := Nil;
reg317433 := PrimCons(reg317431, reg317432)
reg317434 := PrimCons(reg317430, reg317433)
reg317435 := PrimCons(reg317429, reg317434)
__ctx.TailApply(__defun__declare, reg317424, reg317435)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317437 := MakeSymbol("load")
reg317438 := MakeSymbol("string")
reg317439 := MakeSymbol("-->")
reg317440 := MakeSymbol("symbol")
reg317441 := Nil;
reg317442 := PrimCons(reg317440, reg317441)
reg317443 := PrimCons(reg317439, reg317442)
reg317444 := PrimCons(reg317438, reg317443)
__ctx.TailApply(__defun__declare, reg317437, reg317444)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317446 := MakeSymbol("map")
reg317447 := MakeSymbol("A")
reg317448 := MakeSymbol("-->")
reg317449 := MakeSymbol("B")
reg317450 := Nil;
reg317451 := PrimCons(reg317449, reg317450)
reg317452 := PrimCons(reg317448, reg317451)
reg317453 := PrimCons(reg317447, reg317452)
reg317454 := MakeSymbol("-->")
reg317455 := MakeSymbol("list")
reg317456 := MakeSymbol("A")
reg317457 := Nil;
reg317458 := PrimCons(reg317456, reg317457)
reg317459 := PrimCons(reg317455, reg317458)
reg317460 := MakeSymbol("-->")
reg317461 := MakeSymbol("list")
reg317462 := MakeSymbol("B")
reg317463 := Nil;
reg317464 := PrimCons(reg317462, reg317463)
reg317465 := PrimCons(reg317461, reg317464)
reg317466 := Nil;
reg317467 := PrimCons(reg317465, reg317466)
reg317468 := PrimCons(reg317460, reg317467)
reg317469 := PrimCons(reg317459, reg317468)
reg317470 := Nil;
reg317471 := PrimCons(reg317469, reg317470)
reg317472 := PrimCons(reg317454, reg317471)
reg317473 := PrimCons(reg317453, reg317472)
__ctx.TailApply(__defun__declare, reg317446, reg317473)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317475 := MakeSymbol("mapcan")
reg317476 := MakeSymbol("A")
reg317477 := MakeSymbol("-->")
reg317478 := MakeSymbol("list")
reg317479 := MakeSymbol("B")
reg317480 := Nil;
reg317481 := PrimCons(reg317479, reg317480)
reg317482 := PrimCons(reg317478, reg317481)
reg317483 := Nil;
reg317484 := PrimCons(reg317482, reg317483)
reg317485 := PrimCons(reg317477, reg317484)
reg317486 := PrimCons(reg317476, reg317485)
reg317487 := MakeSymbol("-->")
reg317488 := MakeSymbol("list")
reg317489 := MakeSymbol("A")
reg317490 := Nil;
reg317491 := PrimCons(reg317489, reg317490)
reg317492 := PrimCons(reg317488, reg317491)
reg317493 := MakeSymbol("-->")
reg317494 := MakeSymbol("list")
reg317495 := MakeSymbol("B")
reg317496 := Nil;
reg317497 := PrimCons(reg317495, reg317496)
reg317498 := PrimCons(reg317494, reg317497)
reg317499 := Nil;
reg317500 := PrimCons(reg317498, reg317499)
reg317501 := PrimCons(reg317493, reg317500)
reg317502 := PrimCons(reg317492, reg317501)
reg317503 := Nil;
reg317504 := PrimCons(reg317502, reg317503)
reg317505 := PrimCons(reg317487, reg317504)
reg317506 := PrimCons(reg317486, reg317505)
__ctx.TailApply(__defun__declare, reg317475, reg317506)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317508 := MakeSymbol("maxinferences")
reg317509 := MakeSymbol("number")
reg317510 := MakeSymbol("-->")
reg317511 := MakeSymbol("number")
reg317512 := Nil;
reg317513 := PrimCons(reg317511, reg317512)
reg317514 := PrimCons(reg317510, reg317513)
reg317515 := PrimCons(reg317509, reg317514)
__ctx.TailApply(__defun__declare, reg317508, reg317515)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317517 := MakeSymbol("n->string")
reg317518 := MakeSymbol("number")
reg317519 := MakeSymbol("-->")
reg317520 := MakeSymbol("string")
reg317521 := Nil;
reg317522 := PrimCons(reg317520, reg317521)
reg317523 := PrimCons(reg317519, reg317522)
reg317524 := PrimCons(reg317518, reg317523)
__ctx.TailApply(__defun__declare, reg317517, reg317524)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317526 := MakeSymbol("nl")
reg317527 := MakeSymbol("number")
reg317528 := MakeSymbol("-->")
reg317529 := MakeSymbol("number")
reg317530 := Nil;
reg317531 := PrimCons(reg317529, reg317530)
reg317532 := PrimCons(reg317528, reg317531)
reg317533 := PrimCons(reg317527, reg317532)
__ctx.TailApply(__defun__declare, reg317526, reg317533)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317535 := MakeSymbol("not")
reg317536 := MakeSymbol("boolean")
reg317537 := MakeSymbol("-->")
reg317538 := MakeSymbol("boolean")
reg317539 := Nil;
reg317540 := PrimCons(reg317538, reg317539)
reg317541 := PrimCons(reg317537, reg317540)
reg317542 := PrimCons(reg317536, reg317541)
__ctx.TailApply(__defun__declare, reg317535, reg317542)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317544 := MakeSymbol("nth")
reg317545 := MakeSymbol("number")
reg317546 := MakeSymbol("-->")
reg317547 := MakeSymbol("list")
reg317548 := MakeSymbol("A")
reg317549 := Nil;
reg317550 := PrimCons(reg317548, reg317549)
reg317551 := PrimCons(reg317547, reg317550)
reg317552 := MakeSymbol("-->")
reg317553 := MakeSymbol("A")
reg317554 := Nil;
reg317555 := PrimCons(reg317553, reg317554)
reg317556 := PrimCons(reg317552, reg317555)
reg317557 := PrimCons(reg317551, reg317556)
reg317558 := Nil;
reg317559 := PrimCons(reg317557, reg317558)
reg317560 := PrimCons(reg317546, reg317559)
reg317561 := PrimCons(reg317545, reg317560)
__ctx.TailApply(__defun__declare, reg317544, reg317561)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317563 := MakeSymbol("number?")
reg317564 := MakeSymbol("A")
reg317565 := MakeSymbol("-->")
reg317566 := MakeSymbol("boolean")
reg317567 := Nil;
reg317568 := PrimCons(reg317566, reg317567)
reg317569 := PrimCons(reg317565, reg317568)
reg317570 := PrimCons(reg317564, reg317569)
__ctx.TailApply(__defun__declare, reg317563, reg317570)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317572 := MakeSymbol("occurrences")
reg317573 := MakeSymbol("A")
reg317574 := MakeSymbol("-->")
reg317575 := MakeSymbol("B")
reg317576 := MakeSymbol("-->")
reg317577 := MakeSymbol("number")
reg317578 := Nil;
reg317579 := PrimCons(reg317577, reg317578)
reg317580 := PrimCons(reg317576, reg317579)
reg317581 := PrimCons(reg317575, reg317580)
reg317582 := Nil;
reg317583 := PrimCons(reg317581, reg317582)
reg317584 := PrimCons(reg317574, reg317583)
reg317585 := PrimCons(reg317573, reg317584)
__ctx.TailApply(__defun__declare, reg317572, reg317585)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317587 := MakeSymbol("occurs-check")
reg317588 := MakeSymbol("symbol")
reg317589 := MakeSymbol("-->")
reg317590 := MakeSymbol("boolean")
reg317591 := Nil;
reg317592 := PrimCons(reg317590, reg317591)
reg317593 := PrimCons(reg317589, reg317592)
reg317594 := PrimCons(reg317588, reg317593)
__ctx.TailApply(__defun__declare, reg317587, reg317594)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317596 := MakeSymbol("optimise")
reg317597 := MakeSymbol("symbol")
reg317598 := MakeSymbol("-->")
reg317599 := MakeSymbol("boolean")
reg317600 := Nil;
reg317601 := PrimCons(reg317599, reg317600)
reg317602 := PrimCons(reg317598, reg317601)
reg317603 := PrimCons(reg317597, reg317602)
__ctx.TailApply(__defun__declare, reg317596, reg317603)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317605 := MakeSymbol("or")
reg317606 := MakeSymbol("boolean")
reg317607 := MakeSymbol("-->")
reg317608 := MakeSymbol("boolean")
reg317609 := MakeSymbol("-->")
reg317610 := MakeSymbol("boolean")
reg317611 := Nil;
reg317612 := PrimCons(reg317610, reg317611)
reg317613 := PrimCons(reg317609, reg317612)
reg317614 := PrimCons(reg317608, reg317613)
reg317615 := Nil;
reg317616 := PrimCons(reg317614, reg317615)
reg317617 := PrimCons(reg317607, reg317616)
reg317618 := PrimCons(reg317606, reg317617)
__ctx.TailApply(__defun__declare, reg317605, reg317618)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317620 := MakeSymbol("os")
reg317621 := MakeSymbol("-->")
reg317622 := MakeSymbol("string")
reg317623 := Nil;
reg317624 := PrimCons(reg317622, reg317623)
reg317625 := PrimCons(reg317621, reg317624)
__ctx.TailApply(__defun__declare, reg317620, reg317625)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317627 := MakeSymbol("package?")
reg317628 := MakeSymbol("symbol")
reg317629 := MakeSymbol("-->")
reg317630 := MakeSymbol("boolean")
reg317631 := Nil;
reg317632 := PrimCons(reg317630, reg317631)
reg317633 := PrimCons(reg317629, reg317632)
reg317634 := PrimCons(reg317628, reg317633)
__ctx.TailApply(__defun__declare, reg317627, reg317634)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317636 := MakeSymbol("port")
reg317637 := MakeSymbol("-->")
reg317638 := MakeSymbol("string")
reg317639 := Nil;
reg317640 := PrimCons(reg317638, reg317639)
reg317641 := PrimCons(reg317637, reg317640)
__ctx.TailApply(__defun__declare, reg317636, reg317641)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317643 := MakeSymbol("porters")
reg317644 := MakeSymbol("-->")
reg317645 := MakeSymbol("string")
reg317646 := Nil;
reg317647 := PrimCons(reg317645, reg317646)
reg317648 := PrimCons(reg317644, reg317647)
__ctx.TailApply(__defun__declare, reg317643, reg317648)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317650 := MakeSymbol("pos")
reg317651 := MakeSymbol("string")
reg317652 := MakeSymbol("-->")
reg317653 := MakeSymbol("number")
reg317654 := MakeSymbol("-->")
reg317655 := MakeSymbol("string")
reg317656 := Nil;
reg317657 := PrimCons(reg317655, reg317656)
reg317658 := PrimCons(reg317654, reg317657)
reg317659 := PrimCons(reg317653, reg317658)
reg317660 := Nil;
reg317661 := PrimCons(reg317659, reg317660)
reg317662 := PrimCons(reg317652, reg317661)
reg317663 := PrimCons(reg317651, reg317662)
__ctx.TailApply(__defun__declare, reg317650, reg317663)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317665 := MakeSymbol("pr")
reg317666 := MakeSymbol("string")
reg317667 := MakeSymbol("-->")
reg317668 := MakeSymbol("stream")
reg317669 := MakeSymbol("out")
reg317670 := Nil;
reg317671 := PrimCons(reg317669, reg317670)
reg317672 := PrimCons(reg317668, reg317671)
reg317673 := MakeSymbol("-->")
reg317674 := MakeSymbol("string")
reg317675 := Nil;
reg317676 := PrimCons(reg317674, reg317675)
reg317677 := PrimCons(reg317673, reg317676)
reg317678 := PrimCons(reg317672, reg317677)
reg317679 := Nil;
reg317680 := PrimCons(reg317678, reg317679)
reg317681 := PrimCons(reg317667, reg317680)
reg317682 := PrimCons(reg317666, reg317681)
__ctx.TailApply(__defun__declare, reg317665, reg317682)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317684 := MakeSymbol("print")
reg317685 := MakeSymbol("A")
reg317686 := MakeSymbol("-->")
reg317687 := MakeSymbol("A")
reg317688 := Nil;
reg317689 := PrimCons(reg317687, reg317688)
reg317690 := PrimCons(reg317686, reg317689)
reg317691 := PrimCons(reg317685, reg317690)
__ctx.TailApply(__defun__declare, reg317684, reg317691)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317693 := MakeSymbol("profile")
reg317694 := MakeSymbol("A")
reg317695 := MakeSymbol("-->")
reg317696 := MakeSymbol("B")
reg317697 := Nil;
reg317698 := PrimCons(reg317696, reg317697)
reg317699 := PrimCons(reg317695, reg317698)
reg317700 := PrimCons(reg317694, reg317699)
reg317701 := MakeSymbol("-->")
reg317702 := MakeSymbol("A")
reg317703 := MakeSymbol("-->")
reg317704 := MakeSymbol("B")
reg317705 := Nil;
reg317706 := PrimCons(reg317704, reg317705)
reg317707 := PrimCons(reg317703, reg317706)
reg317708 := PrimCons(reg317702, reg317707)
reg317709 := Nil;
reg317710 := PrimCons(reg317708, reg317709)
reg317711 := PrimCons(reg317701, reg317710)
reg317712 := PrimCons(reg317700, reg317711)
__ctx.TailApply(__defun__declare, reg317693, reg317712)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317714 := MakeSymbol("preclude")
reg317715 := MakeSymbol("list")
reg317716 := MakeSymbol("symbol")
reg317717 := Nil;
reg317718 := PrimCons(reg317716, reg317717)
reg317719 := PrimCons(reg317715, reg317718)
reg317720 := MakeSymbol("-->")
reg317721 := MakeSymbol("list")
reg317722 := MakeSymbol("symbol")
reg317723 := Nil;
reg317724 := PrimCons(reg317722, reg317723)
reg317725 := PrimCons(reg317721, reg317724)
reg317726 := Nil;
reg317727 := PrimCons(reg317725, reg317726)
reg317728 := PrimCons(reg317720, reg317727)
reg317729 := PrimCons(reg317719, reg317728)
__ctx.TailApply(__defun__declare, reg317714, reg317729)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317731 := MakeSymbol("shen.proc-nl")
reg317732 := MakeSymbol("string")
reg317733 := MakeSymbol("-->")
reg317734 := MakeSymbol("string")
reg317735 := Nil;
reg317736 := PrimCons(reg317734, reg317735)
reg317737 := PrimCons(reg317733, reg317736)
reg317738 := PrimCons(reg317732, reg317737)
__ctx.TailApply(__defun__declare, reg317731, reg317738)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317740 := MakeSymbol("profile-results")
reg317741 := MakeSymbol("A")
reg317742 := MakeSymbol("-->")
reg317743 := MakeSymbol("B")
reg317744 := Nil;
reg317745 := PrimCons(reg317743, reg317744)
reg317746 := PrimCons(reg317742, reg317745)
reg317747 := PrimCons(reg317741, reg317746)
reg317748 := MakeSymbol("-->")
reg317749 := MakeSymbol("A")
reg317750 := MakeSymbol("-->")
reg317751 := MakeSymbol("B")
reg317752 := Nil;
reg317753 := PrimCons(reg317751, reg317752)
reg317754 := PrimCons(reg317750, reg317753)
reg317755 := PrimCons(reg317749, reg317754)
reg317756 := MakeSymbol("*")
reg317757 := MakeSymbol("number")
reg317758 := Nil;
reg317759 := PrimCons(reg317757, reg317758)
reg317760 := PrimCons(reg317756, reg317759)
reg317761 := PrimCons(reg317755, reg317760)
reg317762 := Nil;
reg317763 := PrimCons(reg317761, reg317762)
reg317764 := PrimCons(reg317748, reg317763)
reg317765 := PrimCons(reg317747, reg317764)
__ctx.TailApply(__defun__declare, reg317740, reg317765)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317767 := MakeSymbol("protect")
reg317768 := MakeSymbol("symbol")
reg317769 := MakeSymbol("-->")
reg317770 := MakeSymbol("symbol")
reg317771 := Nil;
reg317772 := PrimCons(reg317770, reg317771)
reg317773 := PrimCons(reg317769, reg317772)
reg317774 := PrimCons(reg317768, reg317773)
__ctx.TailApply(__defun__declare, reg317767, reg317774)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317776 := MakeSymbol("preclude-all-but")
reg317777 := MakeSymbol("list")
reg317778 := MakeSymbol("symbol")
reg317779 := Nil;
reg317780 := PrimCons(reg317778, reg317779)
reg317781 := PrimCons(reg317777, reg317780)
reg317782 := MakeSymbol("-->")
reg317783 := MakeSymbol("list")
reg317784 := MakeSymbol("symbol")
reg317785 := Nil;
reg317786 := PrimCons(reg317784, reg317785)
reg317787 := PrimCons(reg317783, reg317786)
reg317788 := Nil;
reg317789 := PrimCons(reg317787, reg317788)
reg317790 := PrimCons(reg317782, reg317789)
reg317791 := PrimCons(reg317781, reg317790)
__ctx.TailApply(__defun__declare, reg317776, reg317791)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317793 := MakeSymbol("shen.prhush")
reg317794 := MakeSymbol("string")
reg317795 := MakeSymbol("-->")
reg317796 := MakeSymbol("stream")
reg317797 := MakeSymbol("out")
reg317798 := Nil;
reg317799 := PrimCons(reg317797, reg317798)
reg317800 := PrimCons(reg317796, reg317799)
reg317801 := MakeSymbol("-->")
reg317802 := MakeSymbol("string")
reg317803 := Nil;
reg317804 := PrimCons(reg317802, reg317803)
reg317805 := PrimCons(reg317801, reg317804)
reg317806 := PrimCons(reg317800, reg317805)
reg317807 := Nil;
reg317808 := PrimCons(reg317806, reg317807)
reg317809 := PrimCons(reg317795, reg317808)
reg317810 := PrimCons(reg317794, reg317809)
__ctx.TailApply(__defun__declare, reg317793, reg317810)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317812 := MakeSymbol("ps")
reg317813 := MakeSymbol("symbol")
reg317814 := MakeSymbol("-->")
reg317815 := MakeSymbol("list")
reg317816 := MakeSymbol("unit")
reg317817 := Nil;
reg317818 := PrimCons(reg317816, reg317817)
reg317819 := PrimCons(reg317815, reg317818)
reg317820 := Nil;
reg317821 := PrimCons(reg317819, reg317820)
reg317822 := PrimCons(reg317814, reg317821)
reg317823 := PrimCons(reg317813, reg317822)
__ctx.TailApply(__defun__declare, reg317812, reg317823)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317825 := MakeSymbol("read")
reg317826 := MakeSymbol("stream")
reg317827 := MakeSymbol("in")
reg317828 := Nil;
reg317829 := PrimCons(reg317827, reg317828)
reg317830 := PrimCons(reg317826, reg317829)
reg317831 := MakeSymbol("-->")
reg317832 := MakeSymbol("unit")
reg317833 := Nil;
reg317834 := PrimCons(reg317832, reg317833)
reg317835 := PrimCons(reg317831, reg317834)
reg317836 := PrimCons(reg317830, reg317835)
__ctx.TailApply(__defun__declare, reg317825, reg317836)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317838 := MakeSymbol("read-byte")
reg317839 := MakeSymbol("stream")
reg317840 := MakeSymbol("in")
reg317841 := Nil;
reg317842 := PrimCons(reg317840, reg317841)
reg317843 := PrimCons(reg317839, reg317842)
reg317844 := MakeSymbol("-->")
reg317845 := MakeSymbol("number")
reg317846 := Nil;
reg317847 := PrimCons(reg317845, reg317846)
reg317848 := PrimCons(reg317844, reg317847)
reg317849 := PrimCons(reg317843, reg317848)
__ctx.TailApply(__defun__declare, reg317838, reg317849)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317851 := MakeSymbol("read-file-as-bytelist")
reg317852 := MakeSymbol("string")
reg317853 := MakeSymbol("-->")
reg317854 := MakeSymbol("list")
reg317855 := MakeSymbol("number")
reg317856 := Nil;
reg317857 := PrimCons(reg317855, reg317856)
reg317858 := PrimCons(reg317854, reg317857)
reg317859 := Nil;
reg317860 := PrimCons(reg317858, reg317859)
reg317861 := PrimCons(reg317853, reg317860)
reg317862 := PrimCons(reg317852, reg317861)
__ctx.TailApply(__defun__declare, reg317851, reg317862)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317864 := MakeSymbol("read-file-as-string")
reg317865 := MakeSymbol("string")
reg317866 := MakeSymbol("-->")
reg317867 := MakeSymbol("string")
reg317868 := Nil;
reg317869 := PrimCons(reg317867, reg317868)
reg317870 := PrimCons(reg317866, reg317869)
reg317871 := PrimCons(reg317865, reg317870)
__ctx.TailApply(__defun__declare, reg317864, reg317871)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317873 := MakeSymbol("read-file")
reg317874 := MakeSymbol("string")
reg317875 := MakeSymbol("-->")
reg317876 := MakeSymbol("list")
reg317877 := MakeSymbol("unit")
reg317878 := Nil;
reg317879 := PrimCons(reg317877, reg317878)
reg317880 := PrimCons(reg317876, reg317879)
reg317881 := Nil;
reg317882 := PrimCons(reg317880, reg317881)
reg317883 := PrimCons(reg317875, reg317882)
reg317884 := PrimCons(reg317874, reg317883)
__ctx.TailApply(__defun__declare, reg317873, reg317884)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317886 := MakeSymbol("read-from-string")
reg317887 := MakeSymbol("string")
reg317888 := MakeSymbol("-->")
reg317889 := MakeSymbol("list")
reg317890 := MakeSymbol("unit")
reg317891 := Nil;
reg317892 := PrimCons(reg317890, reg317891)
reg317893 := PrimCons(reg317889, reg317892)
reg317894 := Nil;
reg317895 := PrimCons(reg317893, reg317894)
reg317896 := PrimCons(reg317888, reg317895)
reg317897 := PrimCons(reg317887, reg317896)
__ctx.TailApply(__defun__declare, reg317886, reg317897)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317899 := MakeSymbol("release")
reg317900 := MakeSymbol("-->")
reg317901 := MakeSymbol("string")
reg317902 := Nil;
reg317903 := PrimCons(reg317901, reg317902)
reg317904 := PrimCons(reg317900, reg317903)
__ctx.TailApply(__defun__declare, reg317899, reg317904)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317906 := MakeSymbol("remove")
reg317907 := MakeSymbol("A")
reg317908 := MakeSymbol("-->")
reg317909 := MakeSymbol("list")
reg317910 := MakeSymbol("A")
reg317911 := Nil;
reg317912 := PrimCons(reg317910, reg317911)
reg317913 := PrimCons(reg317909, reg317912)
reg317914 := MakeSymbol("-->")
reg317915 := MakeSymbol("list")
reg317916 := MakeSymbol("A")
reg317917 := Nil;
reg317918 := PrimCons(reg317916, reg317917)
reg317919 := PrimCons(reg317915, reg317918)
reg317920 := Nil;
reg317921 := PrimCons(reg317919, reg317920)
reg317922 := PrimCons(reg317914, reg317921)
reg317923 := PrimCons(reg317913, reg317922)
reg317924 := Nil;
reg317925 := PrimCons(reg317923, reg317924)
reg317926 := PrimCons(reg317908, reg317925)
reg317927 := PrimCons(reg317907, reg317926)
__ctx.TailApply(__defun__declare, reg317906, reg317927)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317929 := MakeSymbol("reverse")
reg317930 := MakeSymbol("list")
reg317931 := MakeSymbol("A")
reg317932 := Nil;
reg317933 := PrimCons(reg317931, reg317932)
reg317934 := PrimCons(reg317930, reg317933)
reg317935 := MakeSymbol("-->")
reg317936 := MakeSymbol("list")
reg317937 := MakeSymbol("A")
reg317938 := Nil;
reg317939 := PrimCons(reg317937, reg317938)
reg317940 := PrimCons(reg317936, reg317939)
reg317941 := Nil;
reg317942 := PrimCons(reg317940, reg317941)
reg317943 := PrimCons(reg317935, reg317942)
reg317944 := PrimCons(reg317934, reg317943)
__ctx.TailApply(__defun__declare, reg317929, reg317944)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317946 := MakeSymbol("simple-error")
reg317947 := MakeSymbol("string")
reg317948 := MakeSymbol("-->")
reg317949 := MakeSymbol("A")
reg317950 := Nil;
reg317951 := PrimCons(reg317949, reg317950)
reg317952 := PrimCons(reg317948, reg317951)
reg317953 := PrimCons(reg317947, reg317952)
__ctx.TailApply(__defun__declare, reg317946, reg317953)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317955 := MakeSymbol("snd")
reg317956 := MakeSymbol("A")
reg317957 := MakeSymbol("*")
reg317958 := MakeSymbol("B")
reg317959 := Nil;
reg317960 := PrimCons(reg317958, reg317959)
reg317961 := PrimCons(reg317957, reg317960)
reg317962 := PrimCons(reg317956, reg317961)
reg317963 := MakeSymbol("-->")
reg317964 := MakeSymbol("B")
reg317965 := Nil;
reg317966 := PrimCons(reg317964, reg317965)
reg317967 := PrimCons(reg317963, reg317966)
reg317968 := PrimCons(reg317962, reg317967)
__ctx.TailApply(__defun__declare, reg317955, reg317968)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317970 := MakeSymbol("specialise")
reg317971 := MakeSymbol("symbol")
reg317972 := MakeSymbol("-->")
reg317973 := MakeSymbol("symbol")
reg317974 := Nil;
reg317975 := PrimCons(reg317973, reg317974)
reg317976 := PrimCons(reg317972, reg317975)
reg317977 := PrimCons(reg317971, reg317976)
__ctx.TailApply(__defun__declare, reg317970, reg317977)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317979 := MakeSymbol("spy")
reg317980 := MakeSymbol("symbol")
reg317981 := MakeSymbol("-->")
reg317982 := MakeSymbol("boolean")
reg317983 := Nil;
reg317984 := PrimCons(reg317982, reg317983)
reg317985 := PrimCons(reg317981, reg317984)
reg317986 := PrimCons(reg317980, reg317985)
__ctx.TailApply(__defun__declare, reg317979, reg317986)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317988 := MakeSymbol("step")
reg317989 := MakeSymbol("symbol")
reg317990 := MakeSymbol("-->")
reg317991 := MakeSymbol("boolean")
reg317992 := Nil;
reg317993 := PrimCons(reg317991, reg317992)
reg317994 := PrimCons(reg317990, reg317993)
reg317995 := PrimCons(reg317989, reg317994)
__ctx.TailApply(__defun__declare, reg317988, reg317995)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg317997 := MakeSymbol("stinput")
reg317998 := MakeSymbol("-->")
reg317999 := MakeSymbol("stream")
reg318000 := MakeSymbol("in")
reg318001 := Nil;
reg318002 := PrimCons(reg318000, reg318001)
reg318003 := PrimCons(reg317999, reg318002)
reg318004 := Nil;
reg318005 := PrimCons(reg318003, reg318004)
reg318006 := PrimCons(reg317998, reg318005)
__ctx.TailApply(__defun__declare, reg317997, reg318006)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318008 := MakeSymbol("sterror")
reg318009 := MakeSymbol("-->")
reg318010 := MakeSymbol("stream")
reg318011 := MakeSymbol("out")
reg318012 := Nil;
reg318013 := PrimCons(reg318011, reg318012)
reg318014 := PrimCons(reg318010, reg318013)
reg318015 := Nil;
reg318016 := PrimCons(reg318014, reg318015)
reg318017 := PrimCons(reg318009, reg318016)
__ctx.TailApply(__defun__declare, reg318008, reg318017)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318019 := MakeSymbol("stoutput")
reg318020 := MakeSymbol("-->")
reg318021 := MakeSymbol("stream")
reg318022 := MakeSymbol("out")
reg318023 := Nil;
reg318024 := PrimCons(reg318022, reg318023)
reg318025 := PrimCons(reg318021, reg318024)
reg318026 := Nil;
reg318027 := PrimCons(reg318025, reg318026)
reg318028 := PrimCons(reg318020, reg318027)
__ctx.TailApply(__defun__declare, reg318019, reg318028)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318030 := MakeSymbol("string?")
reg318031 := MakeSymbol("A")
reg318032 := MakeSymbol("-->")
reg318033 := MakeSymbol("boolean")
reg318034 := Nil;
reg318035 := PrimCons(reg318033, reg318034)
reg318036 := PrimCons(reg318032, reg318035)
reg318037 := PrimCons(reg318031, reg318036)
__ctx.TailApply(__defun__declare, reg318030, reg318037)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318039 := MakeSymbol("str")
reg318040 := MakeSymbol("A")
reg318041 := MakeSymbol("-->")
reg318042 := MakeSymbol("string")
reg318043 := Nil;
reg318044 := PrimCons(reg318042, reg318043)
reg318045 := PrimCons(reg318041, reg318044)
reg318046 := PrimCons(reg318040, reg318045)
__ctx.TailApply(__defun__declare, reg318039, reg318046)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318048 := MakeSymbol("string->n")
reg318049 := MakeSymbol("string")
reg318050 := MakeSymbol("-->")
reg318051 := MakeSymbol("number")
reg318052 := Nil;
reg318053 := PrimCons(reg318051, reg318052)
reg318054 := PrimCons(reg318050, reg318053)
reg318055 := PrimCons(reg318049, reg318054)
__ctx.TailApply(__defun__declare, reg318048, reg318055)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318057 := MakeSymbol("string->symbol")
reg318058 := MakeSymbol("string")
reg318059 := MakeSymbol("-->")
reg318060 := MakeSymbol("symbol")
reg318061 := Nil;
reg318062 := PrimCons(reg318060, reg318061)
reg318063 := PrimCons(reg318059, reg318062)
reg318064 := PrimCons(reg318058, reg318063)
__ctx.TailApply(__defun__declare, reg318057, reg318064)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318066 := MakeSymbol("sum")
reg318067 := MakeSymbol("list")
reg318068 := MakeSymbol("number")
reg318069 := Nil;
reg318070 := PrimCons(reg318068, reg318069)
reg318071 := PrimCons(reg318067, reg318070)
reg318072 := MakeSymbol("-->")
reg318073 := MakeSymbol("number")
reg318074 := Nil;
reg318075 := PrimCons(reg318073, reg318074)
reg318076 := PrimCons(reg318072, reg318075)
reg318077 := PrimCons(reg318071, reg318076)
__ctx.TailApply(__defun__declare, reg318066, reg318077)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318079 := MakeSymbol("symbol?")
reg318080 := MakeSymbol("A")
reg318081 := MakeSymbol("-->")
reg318082 := MakeSymbol("boolean")
reg318083 := Nil;
reg318084 := PrimCons(reg318082, reg318083)
reg318085 := PrimCons(reg318081, reg318084)
reg318086 := PrimCons(reg318080, reg318085)
__ctx.TailApply(__defun__declare, reg318079, reg318086)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318088 := MakeSymbol("systemf")
reg318089 := MakeSymbol("symbol")
reg318090 := MakeSymbol("-->")
reg318091 := MakeSymbol("symbol")
reg318092 := Nil;
reg318093 := PrimCons(reg318091, reg318092)
reg318094 := PrimCons(reg318090, reg318093)
reg318095 := PrimCons(reg318089, reg318094)
__ctx.TailApply(__defun__declare, reg318088, reg318095)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318097 := MakeSymbol("tail")
reg318098 := MakeSymbol("list")
reg318099 := MakeSymbol("A")
reg318100 := Nil;
reg318101 := PrimCons(reg318099, reg318100)
reg318102 := PrimCons(reg318098, reg318101)
reg318103 := MakeSymbol("-->")
reg318104 := MakeSymbol("list")
reg318105 := MakeSymbol("A")
reg318106 := Nil;
reg318107 := PrimCons(reg318105, reg318106)
reg318108 := PrimCons(reg318104, reg318107)
reg318109 := Nil;
reg318110 := PrimCons(reg318108, reg318109)
reg318111 := PrimCons(reg318103, reg318110)
reg318112 := PrimCons(reg318102, reg318111)
__ctx.TailApply(__defun__declare, reg318097, reg318112)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318114 := MakeSymbol("tlstr")
reg318115 := MakeSymbol("string")
reg318116 := MakeSymbol("-->")
reg318117 := MakeSymbol("string")
reg318118 := Nil;
reg318119 := PrimCons(reg318117, reg318118)
reg318120 := PrimCons(reg318116, reg318119)
reg318121 := PrimCons(reg318115, reg318120)
__ctx.TailApply(__defun__declare, reg318114, reg318121)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318123 := MakeSymbol("tlv")
reg318124 := MakeSymbol("vector")
reg318125 := MakeSymbol("A")
reg318126 := Nil;
reg318127 := PrimCons(reg318125, reg318126)
reg318128 := PrimCons(reg318124, reg318127)
reg318129 := MakeSymbol("-->")
reg318130 := MakeSymbol("vector")
reg318131 := MakeSymbol("A")
reg318132 := Nil;
reg318133 := PrimCons(reg318131, reg318132)
reg318134 := PrimCons(reg318130, reg318133)
reg318135 := Nil;
reg318136 := PrimCons(reg318134, reg318135)
reg318137 := PrimCons(reg318129, reg318136)
reg318138 := PrimCons(reg318128, reg318137)
__ctx.TailApply(__defun__declare, reg318123, reg318138)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318140 := MakeSymbol("tc")
reg318141 := MakeSymbol("symbol")
reg318142 := MakeSymbol("-->")
reg318143 := MakeSymbol("boolean")
reg318144 := Nil;
reg318145 := PrimCons(reg318143, reg318144)
reg318146 := PrimCons(reg318142, reg318145)
reg318147 := PrimCons(reg318141, reg318146)
__ctx.TailApply(__defun__declare, reg318140, reg318147)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318149 := MakeSymbol("tc?")
reg318150 := MakeSymbol("-->")
reg318151 := MakeSymbol("boolean")
reg318152 := Nil;
reg318153 := PrimCons(reg318151, reg318152)
reg318154 := PrimCons(reg318150, reg318153)
__ctx.TailApply(__defun__declare, reg318149, reg318154)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318156 := MakeSymbol("thaw")
reg318157 := MakeSymbol("lazy")
reg318158 := MakeSymbol("A")
reg318159 := Nil;
reg318160 := PrimCons(reg318158, reg318159)
reg318161 := PrimCons(reg318157, reg318160)
reg318162 := MakeSymbol("-->")
reg318163 := MakeSymbol("A")
reg318164 := Nil;
reg318165 := PrimCons(reg318163, reg318164)
reg318166 := PrimCons(reg318162, reg318165)
reg318167 := PrimCons(reg318161, reg318166)
__ctx.TailApply(__defun__declare, reg318156, reg318167)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318169 := MakeSymbol("track")
reg318170 := MakeSymbol("symbol")
reg318171 := MakeSymbol("-->")
reg318172 := MakeSymbol("symbol")
reg318173 := Nil;
reg318174 := PrimCons(reg318172, reg318173)
reg318175 := PrimCons(reg318171, reg318174)
reg318176 := PrimCons(reg318170, reg318175)
__ctx.TailApply(__defun__declare, reg318169, reg318176)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318178 := MakeSymbol("trap-error")
reg318179 := MakeSymbol("A")
reg318180 := MakeSymbol("-->")
reg318181 := MakeSymbol("exception")
reg318182 := MakeSymbol("-->")
reg318183 := MakeSymbol("A")
reg318184 := Nil;
reg318185 := PrimCons(reg318183, reg318184)
reg318186 := PrimCons(reg318182, reg318185)
reg318187 := PrimCons(reg318181, reg318186)
reg318188 := MakeSymbol("-->")
reg318189 := MakeSymbol("A")
reg318190 := Nil;
reg318191 := PrimCons(reg318189, reg318190)
reg318192 := PrimCons(reg318188, reg318191)
reg318193 := PrimCons(reg318187, reg318192)
reg318194 := Nil;
reg318195 := PrimCons(reg318193, reg318194)
reg318196 := PrimCons(reg318180, reg318195)
reg318197 := PrimCons(reg318179, reg318196)
__ctx.TailApply(__defun__declare, reg318178, reg318197)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318199 := MakeSymbol("tuple?")
reg318200 := MakeSymbol("A")
reg318201 := MakeSymbol("-->")
reg318202 := MakeSymbol("boolean")
reg318203 := Nil;
reg318204 := PrimCons(reg318202, reg318203)
reg318205 := PrimCons(reg318201, reg318204)
reg318206 := PrimCons(reg318200, reg318205)
__ctx.TailApply(__defun__declare, reg318199, reg318206)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318208 := MakeSymbol("undefmacro")
reg318209 := MakeSymbol("symbol")
reg318210 := MakeSymbol("-->")
reg318211 := MakeSymbol("symbol")
reg318212 := Nil;
reg318213 := PrimCons(reg318211, reg318212)
reg318214 := PrimCons(reg318210, reg318213)
reg318215 := PrimCons(reg318209, reg318214)
__ctx.TailApply(__defun__declare, reg318208, reg318215)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318217 := MakeSymbol("union")
reg318218 := MakeSymbol("list")
reg318219 := MakeSymbol("A")
reg318220 := Nil;
reg318221 := PrimCons(reg318219, reg318220)
reg318222 := PrimCons(reg318218, reg318221)
reg318223 := MakeSymbol("-->")
reg318224 := MakeSymbol("list")
reg318225 := MakeSymbol("A")
reg318226 := Nil;
reg318227 := PrimCons(reg318225, reg318226)
reg318228 := PrimCons(reg318224, reg318227)
reg318229 := MakeSymbol("-->")
reg318230 := MakeSymbol("list")
reg318231 := MakeSymbol("A")
reg318232 := Nil;
reg318233 := PrimCons(reg318231, reg318232)
reg318234 := PrimCons(reg318230, reg318233)
reg318235 := Nil;
reg318236 := PrimCons(reg318234, reg318235)
reg318237 := PrimCons(reg318229, reg318236)
reg318238 := PrimCons(reg318228, reg318237)
reg318239 := Nil;
reg318240 := PrimCons(reg318238, reg318239)
reg318241 := PrimCons(reg318223, reg318240)
reg318242 := PrimCons(reg318222, reg318241)
__ctx.TailApply(__defun__declare, reg318217, reg318242)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318244 := MakeSymbol("unprofile")
reg318245 := MakeSymbol("A")
reg318246 := MakeSymbol("-->")
reg318247 := MakeSymbol("B")
reg318248 := Nil;
reg318249 := PrimCons(reg318247, reg318248)
reg318250 := PrimCons(reg318246, reg318249)
reg318251 := PrimCons(reg318245, reg318250)
reg318252 := MakeSymbol("-->")
reg318253 := MakeSymbol("A")
reg318254 := MakeSymbol("-->")
reg318255 := MakeSymbol("B")
reg318256 := Nil;
reg318257 := PrimCons(reg318255, reg318256)
reg318258 := PrimCons(reg318254, reg318257)
reg318259 := PrimCons(reg318253, reg318258)
reg318260 := Nil;
reg318261 := PrimCons(reg318259, reg318260)
reg318262 := PrimCons(reg318252, reg318261)
reg318263 := PrimCons(reg318251, reg318262)
__ctx.TailApply(__defun__declare, reg318244, reg318263)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318265 := MakeSymbol("untrack")
reg318266 := MakeSymbol("symbol")
reg318267 := MakeSymbol("-->")
reg318268 := MakeSymbol("symbol")
reg318269 := Nil;
reg318270 := PrimCons(reg318268, reg318269)
reg318271 := PrimCons(reg318267, reg318270)
reg318272 := PrimCons(reg318266, reg318271)
__ctx.TailApply(__defun__declare, reg318265, reg318272)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318274 := MakeSymbol("unspecialise")
reg318275 := MakeSymbol("symbol")
reg318276 := MakeSymbol("-->")
reg318277 := MakeSymbol("symbol")
reg318278 := Nil;
reg318279 := PrimCons(reg318277, reg318278)
reg318280 := PrimCons(reg318276, reg318279)
reg318281 := PrimCons(reg318275, reg318280)
__ctx.TailApply(__defun__declare, reg318274, reg318281)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318283 := MakeSymbol("variable?")
reg318284 := MakeSymbol("A")
reg318285 := MakeSymbol("-->")
reg318286 := MakeSymbol("boolean")
reg318287 := Nil;
reg318288 := PrimCons(reg318286, reg318287)
reg318289 := PrimCons(reg318285, reg318288)
reg318290 := PrimCons(reg318284, reg318289)
__ctx.TailApply(__defun__declare, reg318283, reg318290)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318292 := MakeSymbol("vector?")
reg318293 := MakeSymbol("A")
reg318294 := MakeSymbol("-->")
reg318295 := MakeSymbol("boolean")
reg318296 := Nil;
reg318297 := PrimCons(reg318295, reg318296)
reg318298 := PrimCons(reg318294, reg318297)
reg318299 := PrimCons(reg318293, reg318298)
__ctx.TailApply(__defun__declare, reg318292, reg318299)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318301 := MakeSymbol("version")
reg318302 := MakeSymbol("-->")
reg318303 := MakeSymbol("string")
reg318304 := Nil;
reg318305 := PrimCons(reg318303, reg318304)
reg318306 := PrimCons(reg318302, reg318305)
__ctx.TailApply(__defun__declare, reg318301, reg318306)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318308 := MakeSymbol("write-to-file")
reg318309 := MakeSymbol("string")
reg318310 := MakeSymbol("-->")
reg318311 := MakeSymbol("A")
reg318312 := MakeSymbol("-->")
reg318313 := MakeSymbol("A")
reg318314 := Nil;
reg318315 := PrimCons(reg318313, reg318314)
reg318316 := PrimCons(reg318312, reg318315)
reg318317 := PrimCons(reg318311, reg318316)
reg318318 := Nil;
reg318319 := PrimCons(reg318317, reg318318)
reg318320 := PrimCons(reg318310, reg318319)
reg318321 := PrimCons(reg318309, reg318320)
__ctx.TailApply(__defun__declare, reg318308, reg318321)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318323 := MakeSymbol("write-byte")
reg318324 := MakeSymbol("number")
reg318325 := MakeSymbol("-->")
reg318326 := MakeSymbol("stream")
reg318327 := MakeSymbol("out")
reg318328 := Nil;
reg318329 := PrimCons(reg318327, reg318328)
reg318330 := PrimCons(reg318326, reg318329)
reg318331 := MakeSymbol("-->")
reg318332 := MakeSymbol("number")
reg318333 := Nil;
reg318334 := PrimCons(reg318332, reg318333)
reg318335 := PrimCons(reg318331, reg318334)
reg318336 := PrimCons(reg318330, reg318335)
reg318337 := Nil;
reg318338 := PrimCons(reg318336, reg318337)
reg318339 := PrimCons(reg318325, reg318338)
reg318340 := PrimCons(reg318324, reg318339)
__ctx.TailApply(__defun__declare, reg318323, reg318340)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318342 := MakeSymbol("y-or-n?")
reg318343 := MakeSymbol("string")
reg318344 := MakeSymbol("-->")
reg318345 := MakeSymbol("boolean")
reg318346 := Nil;
reg318347 := PrimCons(reg318345, reg318346)
reg318348 := PrimCons(reg318344, reg318347)
reg318349 := PrimCons(reg318343, reg318348)
__ctx.TailApply(__defun__declare, reg318342, reg318349)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318351 := MakeSymbol(">")
reg318352 := MakeSymbol("number")
reg318353 := MakeSymbol("-->")
reg318354 := MakeSymbol("number")
reg318355 := MakeSymbol("-->")
reg318356 := MakeSymbol("boolean")
reg318357 := Nil;
reg318358 := PrimCons(reg318356, reg318357)
reg318359 := PrimCons(reg318355, reg318358)
reg318360 := PrimCons(reg318354, reg318359)
reg318361 := Nil;
reg318362 := PrimCons(reg318360, reg318361)
reg318363 := PrimCons(reg318353, reg318362)
reg318364 := PrimCons(reg318352, reg318363)
__ctx.TailApply(__defun__declare, reg318351, reg318364)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318366 := MakeSymbol("<")
reg318367 := MakeSymbol("number")
reg318368 := MakeSymbol("-->")
reg318369 := MakeSymbol("number")
reg318370 := MakeSymbol("-->")
reg318371 := MakeSymbol("boolean")
reg318372 := Nil;
reg318373 := PrimCons(reg318371, reg318372)
reg318374 := PrimCons(reg318370, reg318373)
reg318375 := PrimCons(reg318369, reg318374)
reg318376 := Nil;
reg318377 := PrimCons(reg318375, reg318376)
reg318378 := PrimCons(reg318368, reg318377)
reg318379 := PrimCons(reg318367, reg318378)
__ctx.TailApply(__defun__declare, reg318366, reg318379)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318381 := MakeSymbol(">=")
reg318382 := MakeSymbol("number")
reg318383 := MakeSymbol("-->")
reg318384 := MakeSymbol("number")
reg318385 := MakeSymbol("-->")
reg318386 := MakeSymbol("boolean")
reg318387 := Nil;
reg318388 := PrimCons(reg318386, reg318387)
reg318389 := PrimCons(reg318385, reg318388)
reg318390 := PrimCons(reg318384, reg318389)
reg318391 := Nil;
reg318392 := PrimCons(reg318390, reg318391)
reg318393 := PrimCons(reg318383, reg318392)
reg318394 := PrimCons(reg318382, reg318393)
__ctx.TailApply(__defun__declare, reg318381, reg318394)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318396 := MakeSymbol("<=")
reg318397 := MakeSymbol("number")
reg318398 := MakeSymbol("-->")
reg318399 := MakeSymbol("number")
reg318400 := MakeSymbol("-->")
reg318401 := MakeSymbol("boolean")
reg318402 := Nil;
reg318403 := PrimCons(reg318401, reg318402)
reg318404 := PrimCons(reg318400, reg318403)
reg318405 := PrimCons(reg318399, reg318404)
reg318406 := Nil;
reg318407 := PrimCons(reg318405, reg318406)
reg318408 := PrimCons(reg318398, reg318407)
reg318409 := PrimCons(reg318397, reg318408)
__ctx.TailApply(__defun__declare, reg318396, reg318409)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318411 := MakeSymbol("=")
reg318412 := MakeSymbol("A")
reg318413 := MakeSymbol("-->")
reg318414 := MakeSymbol("A")
reg318415 := MakeSymbol("-->")
reg318416 := MakeSymbol("boolean")
reg318417 := Nil;
reg318418 := PrimCons(reg318416, reg318417)
reg318419 := PrimCons(reg318415, reg318418)
reg318420 := PrimCons(reg318414, reg318419)
reg318421 := Nil;
reg318422 := PrimCons(reg318420, reg318421)
reg318423 := PrimCons(reg318413, reg318422)
reg318424 := PrimCons(reg318412, reg318423)
__ctx.TailApply(__defun__declare, reg318411, reg318424)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318426 := MakeSymbol("+")
reg318427 := MakeSymbol("number")
reg318428 := MakeSymbol("-->")
reg318429 := MakeSymbol("number")
reg318430 := MakeSymbol("-->")
reg318431 := MakeSymbol("number")
reg318432 := Nil;
reg318433 := PrimCons(reg318431, reg318432)
reg318434 := PrimCons(reg318430, reg318433)
reg318435 := PrimCons(reg318429, reg318434)
reg318436 := Nil;
reg318437 := PrimCons(reg318435, reg318436)
reg318438 := PrimCons(reg318428, reg318437)
reg318439 := PrimCons(reg318427, reg318438)
__ctx.TailApply(__defun__declare, reg318426, reg318439)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318441 := MakeSymbol("/")
reg318442 := MakeSymbol("number")
reg318443 := MakeSymbol("-->")
reg318444 := MakeSymbol("number")
reg318445 := MakeSymbol("-->")
reg318446 := MakeSymbol("number")
reg318447 := Nil;
reg318448 := PrimCons(reg318446, reg318447)
reg318449 := PrimCons(reg318445, reg318448)
reg318450 := PrimCons(reg318444, reg318449)
reg318451 := Nil;
reg318452 := PrimCons(reg318450, reg318451)
reg318453 := PrimCons(reg318443, reg318452)
reg318454 := PrimCons(reg318442, reg318453)
__ctx.TailApply(__defun__declare, reg318441, reg318454)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318456 := MakeSymbol("-")
reg318457 := MakeSymbol("number")
reg318458 := MakeSymbol("-->")
reg318459 := MakeSymbol("number")
reg318460 := MakeSymbol("-->")
reg318461 := MakeSymbol("number")
reg318462 := Nil;
reg318463 := PrimCons(reg318461, reg318462)
reg318464 := PrimCons(reg318460, reg318463)
reg318465 := PrimCons(reg318459, reg318464)
reg318466 := Nil;
reg318467 := PrimCons(reg318465, reg318466)
reg318468 := PrimCons(reg318458, reg318467)
reg318469 := PrimCons(reg318457, reg318468)
__ctx.TailApply(__defun__declare, reg318456, reg318469)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318471 := MakeSymbol("*")
reg318472 := MakeSymbol("number")
reg318473 := MakeSymbol("-->")
reg318474 := MakeSymbol("number")
reg318475 := MakeSymbol("-->")
reg318476 := MakeSymbol("number")
reg318477 := Nil;
reg318478 := PrimCons(reg318476, reg318477)
reg318479 := PrimCons(reg318475, reg318478)
reg318480 := PrimCons(reg318474, reg318479)
reg318481 := Nil;
reg318482 := PrimCons(reg318480, reg318481)
reg318483 := PrimCons(reg318473, reg318482)
reg318484 := PrimCons(reg318472, reg318483)
__ctx.TailApply(__defun__declare, reg318471, reg318484)
return
}, 0))
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg318486 := MakeSymbol("==")
reg318487 := MakeSymbol("A")
reg318488 := MakeSymbol("-->")
reg318489 := MakeSymbol("B")
reg318490 := MakeSymbol("-->")
reg318491 := MakeSymbol("boolean")
reg318492 := Nil;
reg318493 := PrimCons(reg318491, reg318492)
reg318494 := PrimCons(reg318490, reg318493)
reg318495 := PrimCons(reg318489, reg318494)
reg318496 := Nil;
reg318497 := PrimCons(reg318495, reg318496)
reg318498 := PrimCons(reg318488, reg318497)
reg318499 := PrimCons(reg318487, reg318498)
__ctx.TailApply(__defun__declare, reg318486, reg318499)
return
}, 0))
}
