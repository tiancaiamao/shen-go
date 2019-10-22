package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4shen Obj // shen.shen
var __defun__shen_4loop Obj // shen.loop
var __defun__shen_4toplevel_1display_1exception Obj // shen.toplevel-display-exception
var __defun__shen_4credits Obj // shen.credits
var __defun__shen_4initialise__environment Obj // shen.initialise_environment
var __defun__shen_4multiple_1set Obj // shen.multiple-set
var __defun__destroy Obj // destroy
var __defun__shen_4read_1evaluate_1print Obj // shen.read-evaluate-print
var __defun__shen_4retrieve_1from_1history_1if_1needed Obj // shen.retrieve-from-history-if-needed
var __defun__shen_4percent Obj // shen.percent
var __defun__shen_4exclamation Obj // shen.exclamation
var __defun__shen_4prbytes Obj // shen.prbytes
var __defun__shen_4update__history Obj // shen.update_history
var __defun__shen_4toplineread Obj // shen.toplineread
var __defun__shen_4toplineread__loop Obj // shen.toplineread_loop
var __defun__shen_4hat Obj // shen.hat
var __defun__shen_4newline Obj // shen.newline
var __defun__shen_4carriage_1return Obj // shen.carriage-return
var __defun__tc Obj // tc
var __defun__shen_4prompt Obj // shen.prompt
var __defun__shen_4toplevel Obj // shen.toplevel
var __defun__shen_4find_1past_1inputs Obj // shen.find-past-inputs
var __defun__shen_4make_1key Obj // shen.make-key
var __defun__shen_4trim_1gubbins Obj // shen.trim-gubbins
var __defun__shen_4space Obj // shen.space
var __defun__shen_4tab Obj // shen.tab
var __defun__shen_4left_1round Obj // shen.left-round
var __defun__shen_4find Obj // shen.find
var __defun__shen_4prefix_2 Obj // shen.prefix?
var __defun__shen_4print_1past_1inputs Obj // shen.print-past-inputs
var __defun__shen_4toplevel__evaluate Obj // shen.toplevel_evaluate
var __defun__shen_4typecheck_1and_1evaluate Obj // shen.typecheck-and-evaluate
var __defun__shen_4pretty_1type Obj // shen.pretty-type
var __defun__shen_4extract_1pvars Obj // shen.extract-pvars
var __defun__shen_4mult__subst Obj // shen.mult_subst

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291459 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg291459)
return
}, 0))
__defun__shen_4shen = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291460 := __e.Call(__defun__shen_4credits)
_ = reg291460
__ctx.TailApply(__defun__shen_4loop)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.shen", value: __defun__shen_4shen})

__defun__shen_4loop = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291462 := __e.Call(__defun__shen_4initialise__environment)
_ = reg291462
reg291463 := __e.Call(__defun__shen_4prompt)
_ = reg291463
reg291464 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
__ctx.TailApply(__defun__shen_4read_1evaluate_1print)
return
}, 0)
reg291466 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
__ctx.TailApply(__defun__shen_4toplevel_1display_1exception, E)
return
}, 1)
reg291468 := __e.Try(reg291464).Catch(reg291466)
_ = reg291468
__ctx.TailApply(__defun__shen_4loop)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.loop", value: __defun__shen_4loop})

__defun__shen_4toplevel_1display_1exception = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3935 := __args[0]
_ = V3935
reg291470 := PrimErrorToString(V3935)
reg291471 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__pr, reg291470, reg291471)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.toplevel-display-exception", value: __defun__shen_4toplevel_1display_1exception})

__defun__shen_4credits = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291473 := MakeString("\nShen, copyright (C) 2010-2015 Mark Tarver\n")
reg291474 := __e.Call(__defun__stoutput)
reg291475 := __e.Call(__defun__shen_4prhush, reg291473, reg291474)
_ = reg291475
reg291476 := MakeString("www.shenlanguage.org, ")
reg291477 := MakeSymbol("*version*")
reg291478 := PrimValue(reg291477)
reg291479 := MakeString("\n")
reg291480 := MakeSymbol("shen.a")
reg291481 := __e.Call(__defun__shen_4app, reg291478, reg291479, reg291480)
reg291482 := PrimStringConcat(reg291476, reg291481)
reg291483 := __e.Call(__defun__stoutput)
reg291484 := __e.Call(__defun__shen_4prhush, reg291482, reg291483)
_ = reg291484
reg291485 := MakeString("running under ")
reg291486 := MakeSymbol("*language*")
reg291487 := PrimValue(reg291486)
reg291488 := MakeString(", implementation: ")
reg291489 := MakeSymbol("*implementation*")
reg291490 := PrimValue(reg291489)
reg291491 := MakeString("")
reg291492 := MakeSymbol("shen.a")
reg291493 := __e.Call(__defun__shen_4app, reg291490, reg291491, reg291492)
reg291494 := PrimStringConcat(reg291488, reg291493)
reg291495 := MakeSymbol("shen.a")
reg291496 := __e.Call(__defun__shen_4app, reg291487, reg291494, reg291495)
reg291497 := PrimStringConcat(reg291485, reg291496)
reg291498 := __e.Call(__defun__stoutput)
reg291499 := __e.Call(__defun__shen_4prhush, reg291497, reg291498)
_ = reg291499
reg291500 := MakeString("\nport ")
reg291501 := MakeSymbol("*port*")
reg291502 := PrimValue(reg291501)
reg291503 := MakeString(" ported by ")
reg291504 := MakeSymbol("*porters*")
reg291505 := PrimValue(reg291504)
reg291506 := MakeString("\n")
reg291507 := MakeSymbol("shen.a")
reg291508 := __e.Call(__defun__shen_4app, reg291505, reg291506, reg291507)
reg291509 := PrimStringConcat(reg291503, reg291508)
reg291510 := MakeSymbol("shen.a")
reg291511 := __e.Call(__defun__shen_4app, reg291502, reg291509, reg291510)
reg291512 := PrimStringConcat(reg291500, reg291511)
reg291513 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg291512, reg291513)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.credits", value: __defun__shen_4credits})

__defun__shen_4initialise__environment = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291515 := MakeSymbol("shen.*call*")
reg291516 := MakeNumber(0)
reg291517 := MakeSymbol("shen.*infs*")
reg291518 := MakeNumber(0)
reg291519 := MakeSymbol("shen.*process-counter*")
reg291520 := MakeNumber(0)
reg291521 := MakeSymbol("shen.*catch*")
reg291522 := MakeNumber(0)
reg291523 := Nil;
reg291524 := PrimCons(reg291522, reg291523)
reg291525 := PrimCons(reg291521, reg291524)
reg291526 := PrimCons(reg291520, reg291525)
reg291527 := PrimCons(reg291519, reg291526)
reg291528 := PrimCons(reg291518, reg291527)
reg291529 := PrimCons(reg291517, reg291528)
reg291530 := PrimCons(reg291516, reg291529)
reg291531 := PrimCons(reg291515, reg291530)
__ctx.TailApply(__defun__shen_4multiple_1set, reg291531)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.initialise_environment", value: __defun__shen_4initialise__environment})

__defun__shen_4multiple_1set = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3937 := __args[0]
_ = V3937
reg291533 := Nil;
reg291534 := PrimEqual(reg291533, V3937)
if reg291534 == True {
reg291535 := Nil;
__ctx.Return(reg291535)
return
} else {
reg291536 := PrimIsPair(V3937)
var reg291543 Obj
if reg291536 == True {
reg291537 := PrimTail(V3937)
reg291538 := PrimIsPair(reg291537)
var reg291541 Obj
if reg291538 == True {
reg291539 := True;
reg291541 = reg291539
} else {
reg291540 := False;
reg291541 = reg291540
}
reg291543 = reg291541
} else {
reg291542 := False;
reg291543 = reg291542
}
if reg291543 == True {
reg291544 := PrimHead(V3937)
reg291545 := PrimTail(V3937)
reg291546 := PrimHead(reg291545)
reg291547 := PrimSet(reg291544, reg291546)
_ = reg291547
reg291548 := PrimTail(V3937)
reg291549 := PrimTail(reg291548)
__ctx.TailApply(__defun__shen_4multiple_1set, reg291549)
return
} else {
reg291551 := MakeSymbol("shen.multiple-set")
__ctx.TailApply(__defun__shen_4f__error, reg291551)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.multiple-set", value: __defun__shen_4multiple_1set})

__defun__destroy = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3939 := __args[0]
_ = V3939
reg291553 := MakeSymbol("symbol")
__ctx.TailApply(__defun__declare, V3939, reg291553)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "destroy", value: __defun__destroy})

__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291555 := MakeSymbol("shen.*history*")
reg291556 := Nil;
reg291557 := PrimSet(reg291555, reg291556)
__ctx.Return(reg291557)
return
}, 0))
__defun__shen_4read_1evaluate_1print = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291558 := __e.Call(__defun__shen_4toplineread)
Lineread := reg291558
_ = Lineread
reg291559 := MakeSymbol("shen.*history*")
reg291560 := PrimValue(reg291559)
History := reg291560
_ = History
reg291561 := __e.Call(__defun__shen_4retrieve_1from_1history_1if_1needed, Lineread, History)
NewLineread := reg291561
_ = NewLineread
reg291562 := __e.Call(__defun__shen_4update__history, NewLineread, History)
NewHistory := reg291562
_ = NewHistory
reg291563 := __e.Call(__defun__fst, NewLineread)
Parsed := reg291563
_ = Parsed
__ctx.TailApply(__defun__shen_4toplevel, Parsed)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.read-evaluate-print", value: __defun__shen_4read_1evaluate_1print})

__defun__shen_4retrieve_1from_1history_1if_1needed = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3951 := __args[0]
_ = V3951
V3952 := __args[1]
_ = V3952
reg291565 := __e.Call(__defun__tuple_2, V3951)
var reg291585 Obj
if reg291565 == True {
reg291566 := __e.Call(__defun__snd, V3951)
reg291567 := PrimIsPair(reg291566)
var reg291580 Obj
if reg291567 == True {
reg291568 := __e.Call(__defun__snd, V3951)
reg291569 := PrimHead(reg291568)
reg291570 := __e.Call(__defun__shen_4space)
reg291571 := __e.Call(__defun__shen_4newline)
reg291572 := Nil;
reg291573 := PrimCons(reg291571, reg291572)
reg291574 := PrimCons(reg291570, reg291573)
reg291575 := __e.Call(__defun__element_2, reg291569, reg291574)
var reg291578 Obj
if reg291575 == True {
reg291576 := True;
reg291578 = reg291576
} else {
reg291577 := False;
reg291578 = reg291577
}
reg291580 = reg291578
} else {
reg291579 := False;
reg291580 = reg291579
}
var reg291583 Obj
if reg291580 == True {
reg291581 := True;
reg291583 = reg291581
} else {
reg291582 := False;
reg291583 = reg291582
}
reg291585 = reg291583
} else {
reg291584 := False;
reg291585 = reg291584
}
if reg291585 == True {
reg291586 := __e.Call(__defun__fst, V3951)
reg291587 := __e.Call(__defun__snd, V3951)
reg291588 := PrimTail(reg291587)
reg291589 := __e.Call(__defun___8p, reg291586, reg291588)
__ctx.TailApply(__defun__shen_4retrieve_1from_1history_1if_1needed, reg291589, V3952)
return
} else {
reg291591 := __e.Call(__defun__tuple_2, V3951)
var reg291641 Obj
if reg291591 == True {
reg291592 := __e.Call(__defun__snd, V3951)
reg291593 := PrimIsPair(reg291592)
var reg291636 Obj
if reg291593 == True {
reg291594 := __e.Call(__defun__snd, V3951)
reg291595 := PrimTail(reg291594)
reg291596 := PrimIsPair(reg291595)
var reg291631 Obj
if reg291596 == True {
reg291597 := Nil;
reg291598 := __e.Call(__defun__snd, V3951)
reg291599 := PrimTail(reg291598)
reg291600 := PrimTail(reg291599)
reg291601 := PrimEqual(reg291597, reg291600)
var reg291626 Obj
if reg291601 == True {
reg291602 := PrimIsPair(V3952)
var reg291621 Obj
if reg291602 == True {
reg291603 := __e.Call(__defun__snd, V3951)
reg291604 := PrimHead(reg291603)
reg291605 := __e.Call(__defun__shen_4exclamation)
reg291606 := PrimEqual(reg291604, reg291605)
var reg291616 Obj
if reg291606 == True {
reg291607 := __e.Call(__defun__snd, V3951)
reg291608 := PrimTail(reg291607)
reg291609 := PrimHead(reg291608)
reg291610 := __e.Call(__defun__shen_4exclamation)
reg291611 := PrimEqual(reg291609, reg291610)
var reg291614 Obj
if reg291611 == True {
reg291612 := True;
reg291614 = reg291612
} else {
reg291613 := False;
reg291614 = reg291613
}
reg291616 = reg291614
} else {
reg291615 := False;
reg291616 = reg291615
}
var reg291619 Obj
if reg291616 == True {
reg291617 := True;
reg291619 = reg291617
} else {
reg291618 := False;
reg291619 = reg291618
}
reg291621 = reg291619
} else {
reg291620 := False;
reg291621 = reg291620
}
var reg291624 Obj
if reg291621 == True {
reg291622 := True;
reg291624 = reg291622
} else {
reg291623 := False;
reg291624 = reg291623
}
reg291626 = reg291624
} else {
reg291625 := False;
reg291626 = reg291625
}
var reg291629 Obj
if reg291626 == True {
reg291627 := True;
reg291629 = reg291627
} else {
reg291628 := False;
reg291629 = reg291628
}
reg291631 = reg291629
} else {
reg291630 := False;
reg291631 = reg291630
}
var reg291634 Obj
if reg291631 == True {
reg291632 := True;
reg291634 = reg291632
} else {
reg291633 := False;
reg291634 = reg291633
}
reg291636 = reg291634
} else {
reg291635 := False;
reg291636 = reg291635
}
var reg291639 Obj
if reg291636 == True {
reg291637 := True;
reg291639 = reg291637
} else {
reg291638 := False;
reg291639 = reg291638
}
reg291641 = reg291639
} else {
reg291640 := False;
reg291641 = reg291640
}
if reg291641 == True {
reg291642 := PrimHead(V3952)
reg291643 := __e.Call(__defun__snd, reg291642)
reg291644 := __e.Call(__defun__shen_4prbytes, reg291643)
PastPrint := reg291644
_ = PastPrint
reg291645 := PrimHead(V3952)
__ctx.Return(reg291645)
return
} else {
reg291646 := __e.Call(__defun__tuple_2, V3951)
var reg291662 Obj
if reg291646 == True {
reg291647 := __e.Call(__defun__snd, V3951)
reg291648 := PrimIsPair(reg291647)
var reg291657 Obj
if reg291648 == True {
reg291649 := __e.Call(__defun__snd, V3951)
reg291650 := PrimHead(reg291649)
reg291651 := __e.Call(__defun__shen_4exclamation)
reg291652 := PrimEqual(reg291650, reg291651)
var reg291655 Obj
if reg291652 == True {
reg291653 := True;
reg291655 = reg291653
} else {
reg291654 := False;
reg291655 = reg291654
}
reg291657 = reg291655
} else {
reg291656 := False;
reg291657 = reg291656
}
var reg291660 Obj
if reg291657 == True {
reg291658 := True;
reg291660 = reg291658
} else {
reg291659 := False;
reg291660 = reg291659
}
reg291662 = reg291660
} else {
reg291661 := False;
reg291662 = reg291661
}
if reg291662 == True {
reg291663 := __e.Call(__defun__snd, V3951)
reg291664 := PrimTail(reg291663)
reg291665 := __e.Call(__defun__shen_4make_1key, reg291664, V3952)
Key_2 := reg291665
_ = Key_2
reg291666 := __e.Call(__defun__shen_4find_1past_1inputs, Key_2, V3952)
reg291667 := __e.Call(__defun__head, reg291666)
Find := reg291667
_ = Find
reg291668 := __e.Call(__defun__snd, Find)
reg291669 := __e.Call(__defun__shen_4prbytes, reg291668)
PastPrint := reg291669
_ = PastPrint
__ctx.Return(Find)
return
} else {
reg291670 := __e.Call(__defun__tuple_2, V3951)
var reg291695 Obj
if reg291670 == True {
reg291671 := __e.Call(__defun__snd, V3951)
reg291672 := PrimIsPair(reg291671)
var reg291690 Obj
if reg291672 == True {
reg291673 := Nil;
reg291674 := __e.Call(__defun__snd, V3951)
reg291675 := PrimTail(reg291674)
reg291676 := PrimEqual(reg291673, reg291675)
var reg291685 Obj
if reg291676 == True {
reg291677 := __e.Call(__defun__snd, V3951)
reg291678 := PrimHead(reg291677)
reg291679 := __e.Call(__defun__shen_4percent)
reg291680 := PrimEqual(reg291678, reg291679)
var reg291683 Obj
if reg291680 == True {
reg291681 := True;
reg291683 = reg291681
} else {
reg291682 := False;
reg291683 = reg291682
}
reg291685 = reg291683
} else {
reg291684 := False;
reg291685 = reg291684
}
var reg291688 Obj
if reg291685 == True {
reg291686 := True;
reg291688 = reg291686
} else {
reg291687 := False;
reg291688 = reg291687
}
reg291690 = reg291688
} else {
reg291689 := False;
reg291690 = reg291689
}
var reg291693 Obj
if reg291690 == True {
reg291691 := True;
reg291693 = reg291691
} else {
reg291692 := False;
reg291693 = reg291692
}
reg291695 = reg291693
} else {
reg291694 := False;
reg291695 = reg291694
}
if reg291695 == True {
reg291696 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg291697 := True;
__ctx.Return(reg291697)
return
}, 1)
reg291698 := __e.Call(__defun__reverse, V3952)
reg291699 := MakeNumber(0)
reg291700 := __e.Call(__defun__shen_4print_1past_1inputs, reg291696, reg291698, reg291699)
_ = reg291700
__ctx.TailApply(__defun__abort)
return
} else {
reg291702 := __e.Call(__defun__tuple_2, V3951)
var reg291718 Obj
if reg291702 == True {
reg291703 := __e.Call(__defun__snd, V3951)
reg291704 := PrimIsPair(reg291703)
var reg291713 Obj
if reg291704 == True {
reg291705 := __e.Call(__defun__snd, V3951)
reg291706 := PrimHead(reg291705)
reg291707 := __e.Call(__defun__shen_4percent)
reg291708 := PrimEqual(reg291706, reg291707)
var reg291711 Obj
if reg291708 == True {
reg291709 := True;
reg291711 = reg291709
} else {
reg291710 := False;
reg291711 = reg291710
}
reg291713 = reg291711
} else {
reg291712 := False;
reg291713 = reg291712
}
var reg291716 Obj
if reg291713 == True {
reg291714 := True;
reg291716 = reg291714
} else {
reg291715 := False;
reg291716 = reg291715
}
reg291718 = reg291716
} else {
reg291717 := False;
reg291718 = reg291717
}
if reg291718 == True {
reg291719 := __e.Call(__defun__snd, V3951)
reg291720 := PrimTail(reg291719)
reg291721 := __e.Call(__defun__shen_4make_1key, reg291720, V3952)
Key_2 := reg291721
_ = Key_2
reg291722 := __e.Call(__defun__reverse, V3952)
reg291723 := MakeNumber(0)
reg291724 := __e.Call(__defun__shen_4print_1past_1inputs, Key_2, reg291722, reg291723)
Pastprint := reg291724
_ = Pastprint
__ctx.TailApply(__defun__abort)
return
} else {
__ctx.Return(V3951)
return
}
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.retrieve-from-history-if-needed", value: __defun__shen_4retrieve_1from_1history_1if_1needed})

__defun__shen_4percent = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291726 := MakeNumber(37)
__ctx.Return(reg291726)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.percent", value: __defun__shen_4percent})

__defun__shen_4exclamation = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291727 := MakeNumber(33)
__ctx.Return(reg291727)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.exclamation", value: __defun__shen_4exclamation})

__defun__shen_4prbytes = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3954 := __args[0]
_ = V3954
reg291728 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
Byte := __args[0]
_ = Byte
reg291729 := PrimNumberToString(Byte)
reg291730 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__pr, reg291729, reg291730)
return
}, 1)
reg291732 := __e.Call(__defun__shen_4for_1each, reg291728, V3954)
_ = reg291732
reg291733 := MakeNumber(1)
__ctx.TailApply(__defun__nl, reg291733)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.prbytes", value: __defun__shen_4prbytes})

__defun__shen_4update__history = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3957 := __args[0]
_ = V3957
V3958 := __args[1]
_ = V3958
reg291735 := MakeSymbol("shen.*history*")
reg291736 := PrimCons(V3957, V3958)
reg291737 := PrimSet(reg291735, reg291736)
__ctx.Return(reg291737)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.update_history", value: __defun__shen_4update__history})

__defun__shen_4toplineread = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291738 := __e.Call(__defun__stinput)
reg291739 := __e.Call(__defun__shen_4read_1char_1code, reg291738)
reg291740 := Nil;
__ctx.TailApply(__defun__shen_4toplineread__loop, reg291739, reg291740)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.toplineread", value: __defun__shen_4toplineread})

__defun__shen_4toplineread__loop = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3962 := __args[0]
_ = V3962
V3963 := __args[1]
_ = V3963
reg291742 := __e.Call(__defun__shen_4hat)
reg291743 := PrimEqual(V3962, reg291742)
if reg291743 == True {
reg291744 := MakeString("line read aborted")
reg291745 := PrimSimpleError(reg291744)
__ctx.Return(reg291745)
return
} else {
reg291746 := __e.Call(__defun__shen_4newline)
reg291747 := __e.Call(__defun__shen_4carriage_1return)
reg291748 := Nil;
reg291749 := PrimCons(reg291747, reg291748)
reg291750 := PrimCons(reg291746, reg291749)
reg291751 := __e.Call(__defun__element_2, V3962, reg291750)
if reg291751 == True {
reg291752 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg291754 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg291755 := MakeSymbol("shen.nextline")
__ctx.Return(reg291755)
return
}, 1)
reg291756 := __e.Call(__defun__compile, reg291752, V3963, reg291754)
Line := reg291756
_ = Line
reg291757 := __e.Call(__defun__shen_4record_1it, V3963)
It := reg291757
_ = It
reg291758 := MakeSymbol("shen.nextline")
reg291759 := PrimEqual(Line, reg291758)
var reg291765 Obj
if reg291759 == True {
reg291760 := True;
reg291765 = reg291760
} else {
reg291761 := __e.Call(__defun__empty_2, Line)
var reg291764 Obj
if reg291761 == True {
reg291762 := True;
reg291764 = reg291762
} else {
reg291763 := False;
reg291764 = reg291763
}
reg291765 = reg291764
}
if reg291765 == True {
reg291766 := __e.Call(__defun__stinput)
reg291767 := __e.Call(__defun__shen_4read_1char_1code, reg291766)
reg291768 := Nil;
reg291769 := PrimCons(V3962, reg291768)
reg291770 := __e.Call(__defun__append, V3963, reg291769)
__ctx.TailApply(__defun__shen_4toplineread__loop, reg291767, reg291770)
return
} else {
__ctx.TailApply(__defun___8p, Line, V3963)
return
}
} else {
reg291773 := __e.Call(__defun__stinput)
reg291774 := __e.Call(__defun__shen_4read_1char_1code, reg291773)
reg291775 := MakeNumber(-1)
reg291776 := PrimEqual(V3962, reg291775)
var reg291780 Obj
if reg291776 == True {
reg291780 = V3963
} else {
reg291777 := Nil;
reg291778 := PrimCons(V3962, reg291777)
reg291779 := __e.Call(__defun__append, V3963, reg291778)
reg291780 = reg291779
}
__ctx.TailApply(__defun__shen_4toplineread__loop, reg291774, reg291780)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.toplineread_loop", value: __defun__shen_4toplineread__loop})

__defun__shen_4hat = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291782 := MakeNumber(94)
__ctx.Return(reg291782)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.hat", value: __defun__shen_4hat})

__defun__shen_4newline = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291783 := MakeNumber(10)
__ctx.Return(reg291783)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.newline", value: __defun__shen_4newline})

__defun__shen_4carriage_1return = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291784 := MakeNumber(13)
__ctx.Return(reg291784)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.carriage-return", value: __defun__shen_4carriage_1return})

__defun__tc = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3969 := __args[0]
_ = V3969
reg291785 := MakeSymbol("+")
reg291786 := PrimEqual(reg291785, V3969)
if reg291786 == True {
reg291787 := MakeSymbol("shen.*tc*")
reg291788 := True;
reg291789 := PrimSet(reg291787, reg291788)
__ctx.Return(reg291789)
return
} else {
reg291790 := MakeSymbol("-")
reg291791 := PrimEqual(reg291790, V3969)
if reg291791 == True {
reg291792 := MakeSymbol("shen.*tc*")
reg291793 := False;
reg291794 := PrimSet(reg291792, reg291793)
__ctx.Return(reg291794)
return
} else {
reg291795 := MakeString("tc expects a + or -")
reg291796 := PrimSimpleError(reg291795)
__ctx.Return(reg291796)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "tc", value: __defun__tc})

__defun__shen_4prompt = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291797 := MakeSymbol("shen.*tc*")
reg291798 := PrimValue(reg291797)
if reg291798 == True {
reg291799 := MakeString("\n\n(")
reg291800 := MakeSymbol("shen.*history*")
reg291801 := PrimValue(reg291800)
reg291802 := __e.Call(__defun__length, reg291801)
reg291803 := MakeString("+) ")
reg291804 := MakeSymbol("shen.a")
reg291805 := __e.Call(__defun__shen_4app, reg291802, reg291803, reg291804)
reg291806 := PrimStringConcat(reg291799, reg291805)
reg291807 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg291806, reg291807)
return
} else {
reg291809 := MakeString("\n\n(")
reg291810 := MakeSymbol("shen.*history*")
reg291811 := PrimValue(reg291810)
reg291812 := __e.Call(__defun__length, reg291811)
reg291813 := MakeString("-) ")
reg291814 := MakeSymbol("shen.a")
reg291815 := __e.Call(__defun__shen_4app, reg291812, reg291813, reg291814)
reg291816 := PrimStringConcat(reg291809, reg291815)
reg291817 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg291816, reg291817)
return
}
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.prompt", value: __defun__shen_4prompt})

__defun__shen_4toplevel = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3971 := __args[0]
_ = V3971
reg291819 := MakeSymbol("shen.*tc*")
reg291820 := PrimValue(reg291819)
__ctx.TailApply(__defun__shen_4toplevel__evaluate, V3971, reg291820)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.toplevel", value: __defun__shen_4toplevel})

__defun__shen_4find_1past_1inputs = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3974 := __args[0]
_ = V3974
V3975 := __args[1]
_ = V3975
reg291822 := __e.Call(__defun__shen_4find, V3974, V3975)
F := reg291822
_ = F
reg291823 := __e.Call(__defun__empty_2, F)
if reg291823 == True {
reg291824 := MakeString("input not found\n")
reg291825 := PrimSimpleError(reg291824)
__ctx.Return(reg291825)
return
} else {
__ctx.Return(F)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.find-past-inputs", value: __defun__shen_4find_1past_1inputs})

__defun__shen_4make_1key = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3978 := __args[0]
_ = V3978
V3979 := __args[1]
_ = V3979
reg291826 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4_5st__input_6, X)
return
}, 1)
reg291828 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
E := __args[0]
_ = E
reg291829 := PrimIsPair(E)
if reg291829 == True {
reg291830 := MakeString("parse error here: ")
reg291831 := MakeString("\n")
reg291832 := MakeSymbol("shen.s")
reg291833 := __e.Call(__defun__shen_4app, E, reg291831, reg291832)
reg291834 := PrimStringConcat(reg291830, reg291833)
reg291835 := PrimSimpleError(reg291834)
__ctx.Return(reg291835)
return
} else {
reg291836 := MakeString("parse error\n")
reg291837 := PrimSimpleError(reg291836)
__ctx.Return(reg291837)
return
}
}, 1)
reg291838 := __e.Call(__defun__compile, reg291826, V3978, reg291828)
reg291839 := PrimHead(reg291838)
Atom := reg291839
_ = Atom
reg291840 := PrimIsInteger(Atom)
if reg291840 == True {
reg291841 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg291842 := MakeNumber(1)
reg291843 := PrimNumberAdd(Atom, reg291842)
reg291844 := __e.Call(__defun__reverse, V3979)
reg291845 := __e.Call(__defun__nth, reg291843, reg291844)
reg291846 := PrimEqual(X, reg291845)
__ctx.Return(reg291846)
return
}, 1)
__ctx.Return(reg291841)
return
} else {
reg291847 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg291848 := __e.Call(__defun__snd, X)
reg291849 := __e.Call(__defun__shen_4trim_1gubbins, reg291848)
__ctx.TailApply(__defun__shen_4prefix_2, V3978, reg291849)
return
}, 1)
__ctx.Return(reg291847)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.make-key", value: __defun__shen_4make_1key})

__defun__shen_4trim_1gubbins = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3981 := __args[0]
_ = V3981
reg291851 := PrimIsPair(V3981)
var reg291859 Obj
if reg291851 == True {
reg291852 := PrimHead(V3981)
reg291853 := __e.Call(__defun__shen_4space)
reg291854 := PrimEqual(reg291852, reg291853)
var reg291857 Obj
if reg291854 == True {
reg291855 := True;
reg291857 = reg291855
} else {
reg291856 := False;
reg291857 = reg291856
}
reg291859 = reg291857
} else {
reg291858 := False;
reg291859 = reg291858
}
if reg291859 == True {
reg291860 := PrimTail(V3981)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg291860)
return
} else {
reg291862 := PrimIsPair(V3981)
var reg291870 Obj
if reg291862 == True {
reg291863 := PrimHead(V3981)
reg291864 := __e.Call(__defun__shen_4newline)
reg291865 := PrimEqual(reg291863, reg291864)
var reg291868 Obj
if reg291865 == True {
reg291866 := True;
reg291868 = reg291866
} else {
reg291867 := False;
reg291868 = reg291867
}
reg291870 = reg291868
} else {
reg291869 := False;
reg291870 = reg291869
}
if reg291870 == True {
reg291871 := PrimTail(V3981)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg291871)
return
} else {
reg291873 := PrimIsPair(V3981)
var reg291881 Obj
if reg291873 == True {
reg291874 := PrimHead(V3981)
reg291875 := __e.Call(__defun__shen_4carriage_1return)
reg291876 := PrimEqual(reg291874, reg291875)
var reg291879 Obj
if reg291876 == True {
reg291877 := True;
reg291879 = reg291877
} else {
reg291878 := False;
reg291879 = reg291878
}
reg291881 = reg291879
} else {
reg291880 := False;
reg291881 = reg291880
}
if reg291881 == True {
reg291882 := PrimTail(V3981)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg291882)
return
} else {
reg291884 := PrimIsPair(V3981)
var reg291892 Obj
if reg291884 == True {
reg291885 := PrimHead(V3981)
reg291886 := __e.Call(__defun__shen_4tab)
reg291887 := PrimEqual(reg291885, reg291886)
var reg291890 Obj
if reg291887 == True {
reg291888 := True;
reg291890 = reg291888
} else {
reg291889 := False;
reg291890 = reg291889
}
reg291892 = reg291890
} else {
reg291891 := False;
reg291892 = reg291891
}
if reg291892 == True {
reg291893 := PrimTail(V3981)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg291893)
return
} else {
reg291895 := PrimIsPair(V3981)
var reg291903 Obj
if reg291895 == True {
reg291896 := PrimHead(V3981)
reg291897 := __e.Call(__defun__shen_4left_1round)
reg291898 := PrimEqual(reg291896, reg291897)
var reg291901 Obj
if reg291898 == True {
reg291899 := True;
reg291901 = reg291899
} else {
reg291900 := False;
reg291901 = reg291900
}
reg291903 = reg291901
} else {
reg291902 := False;
reg291903 = reg291902
}
if reg291903 == True {
reg291904 := PrimTail(V3981)
__ctx.TailApply(__defun__shen_4trim_1gubbins, reg291904)
return
} else {
__ctx.Return(V3981)
return
}
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.trim-gubbins", value: __defun__shen_4trim_1gubbins})

__defun__shen_4space = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291906 := MakeNumber(32)
__ctx.Return(reg291906)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.space", value: __defun__shen_4space})

__defun__shen_4tab = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291907 := MakeNumber(9)
__ctx.Return(reg291907)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.tab", value: __defun__shen_4tab})

__defun__shen_4left_1round = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg291908 := MakeNumber(40)
__ctx.Return(reg291908)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.left-round", value: __defun__shen_4left_1round})

__defun__shen_4find = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V3990 := __args[0]
_ = V3990
V3991 := __args[1]
_ = V3991
reg291909 := Nil;
reg291910 := PrimEqual(reg291909, V3991)
if reg291910 == True {
reg291911 := Nil;
__ctx.Return(reg291911)
return
} else {
reg291912 := PrimIsPair(V3991)
var reg291919 Obj
if reg291912 == True {
reg291913 := PrimHead(V3991)
reg291914 := __e.Call(V3990, reg291913)
var reg291917 Obj
if reg291914 == True {
reg291915 := True;
reg291917 = reg291915
} else {
reg291916 := False;
reg291917 = reg291916
}
reg291919 = reg291917
} else {
reg291918 := False;
reg291919 = reg291918
}
if reg291919 == True {
reg291920 := PrimHead(V3991)
reg291921 := PrimTail(V3991)
reg291922 := __e.Call(__defun__shen_4find, V3990, reg291921)
reg291923 := PrimCons(reg291920, reg291922)
__ctx.Return(reg291923)
return
} else {
reg291924 := PrimIsPair(V3991)
if reg291924 == True {
reg291925 := PrimTail(V3991)
__ctx.TailApply(__defun__shen_4find, V3990, reg291925)
return
} else {
reg291927 := MakeSymbol("shen.find")
__ctx.TailApply(__defun__shen_4f__error, reg291927)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.find", value: __defun__shen_4find})

__defun__shen_4prefix_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4005 := __args[0]
_ = V4005
V4006 := __args[1]
_ = V4006
reg291929 := Nil;
reg291930 := PrimEqual(reg291929, V4005)
if reg291930 == True {
reg291931 := True;
__ctx.Return(reg291931)
return
} else {
reg291932 := PrimIsPair(V4005)
var reg291946 Obj
if reg291932 == True {
reg291933 := PrimIsPair(V4006)
var reg291941 Obj
if reg291933 == True {
reg291934 := PrimHead(V4006)
reg291935 := PrimHead(V4005)
reg291936 := PrimEqual(reg291934, reg291935)
var reg291939 Obj
if reg291936 == True {
reg291937 := True;
reg291939 = reg291937
} else {
reg291938 := False;
reg291939 = reg291938
}
reg291941 = reg291939
} else {
reg291940 := False;
reg291941 = reg291940
}
var reg291944 Obj
if reg291941 == True {
reg291942 := True;
reg291944 = reg291942
} else {
reg291943 := False;
reg291944 = reg291943
}
reg291946 = reg291944
} else {
reg291945 := False;
reg291946 = reg291945
}
if reg291946 == True {
reg291947 := PrimTail(V4005)
reg291948 := PrimTail(V4006)
__ctx.TailApply(__defun__shen_4prefix_2, reg291947, reg291948)
return
} else {
reg291950 := False;
__ctx.Return(reg291950)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.prefix?", value: __defun__shen_4prefix_2})

__defun__shen_4print_1past_1inputs = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4018 := __args[0]
_ = V4018
V4019 := __args[1]
_ = V4019
V4020 := __args[2]
_ = V4020
reg291951 := Nil;
reg291952 := PrimEqual(reg291951, V4019)
if reg291952 == True {
reg291953 := MakeSymbol("_")
__ctx.Return(reg291953)
return
} else {
reg291954 := PrimIsPair(V4019)
var reg291962 Obj
if reg291954 == True {
reg291955 := PrimHead(V4019)
reg291956 := __e.Call(V4018, reg291955)
reg291957 := PrimNot(reg291956)
var reg291960 Obj
if reg291957 == True {
reg291958 := True;
reg291960 = reg291958
} else {
reg291959 := False;
reg291960 = reg291959
}
reg291962 = reg291960
} else {
reg291961 := False;
reg291962 = reg291961
}
if reg291962 == True {
reg291963 := PrimTail(V4019)
reg291964 := MakeNumber(1)
reg291965 := PrimNumberAdd(V4020, reg291964)
__ctx.TailApply(__defun__shen_4print_1past_1inputs, V4018, reg291963, reg291965)
return
} else {
reg291967 := PrimIsPair(V4019)
var reg291974 Obj
if reg291967 == True {
reg291968 := PrimHead(V4019)
reg291969 := __e.Call(__defun__tuple_2, reg291968)
var reg291972 Obj
if reg291969 == True {
reg291970 := True;
reg291972 = reg291970
} else {
reg291971 := False;
reg291972 = reg291971
}
reg291974 = reg291972
} else {
reg291973 := False;
reg291974 = reg291973
}
if reg291974 == True {
reg291975 := MakeString(". ")
reg291976 := MakeSymbol("shen.a")
reg291977 := __e.Call(__defun__shen_4app, V4020, reg291975, reg291976)
reg291978 := __e.Call(__defun__stoutput)
reg291979 := __e.Call(__defun__shen_4prhush, reg291977, reg291978)
_ = reg291979
reg291980 := PrimHead(V4019)
reg291981 := __e.Call(__defun__snd, reg291980)
reg291982 := __e.Call(__defun__shen_4prbytes, reg291981)
_ = reg291982
reg291983 := PrimTail(V4019)
reg291984 := MakeNumber(1)
reg291985 := PrimNumberAdd(V4020, reg291984)
__ctx.TailApply(__defun__shen_4print_1past_1inputs, V4018, reg291983, reg291985)
return
} else {
reg291987 := MakeSymbol("shen.print-past-inputs")
__ctx.TailApply(__defun__shen_4f__error, reg291987)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.print-past-inputs", value: __defun__shen_4print_1past_1inputs})

__defun__shen_4toplevel__evaluate = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4023 := __args[0]
_ = V4023
V4024 := __args[1]
_ = V4024
reg291989 := PrimIsPair(V4023)
var reg292030 Obj
if reg291989 == True {
reg291990 := PrimTail(V4023)
reg291991 := PrimIsPair(reg291990)
var reg292025 Obj
if reg291991 == True {
reg291992 := MakeSymbol(":")
reg291993 := PrimTail(V4023)
reg291994 := PrimHead(reg291993)
reg291995 := PrimEqual(reg291992, reg291994)
var reg292020 Obj
if reg291995 == True {
reg291996 := PrimTail(V4023)
reg291997 := PrimTail(reg291996)
reg291998 := PrimIsPair(reg291997)
var reg292015 Obj
if reg291998 == True {
reg291999 := Nil;
reg292000 := PrimTail(V4023)
reg292001 := PrimTail(reg292000)
reg292002 := PrimTail(reg292001)
reg292003 := PrimEqual(reg291999, reg292002)
var reg292010 Obj
if reg292003 == True {
reg292004 := True;
reg292005 := PrimEqual(reg292004, V4024)
var reg292008 Obj
if reg292005 == True {
reg292006 := True;
reg292008 = reg292006
} else {
reg292007 := False;
reg292008 = reg292007
}
reg292010 = reg292008
} else {
reg292009 := False;
reg292010 = reg292009
}
var reg292013 Obj
if reg292010 == True {
reg292011 := True;
reg292013 = reg292011
} else {
reg292012 := False;
reg292013 = reg292012
}
reg292015 = reg292013
} else {
reg292014 := False;
reg292015 = reg292014
}
var reg292018 Obj
if reg292015 == True {
reg292016 := True;
reg292018 = reg292016
} else {
reg292017 := False;
reg292018 = reg292017
}
reg292020 = reg292018
} else {
reg292019 := False;
reg292020 = reg292019
}
var reg292023 Obj
if reg292020 == True {
reg292021 := True;
reg292023 = reg292021
} else {
reg292022 := False;
reg292023 = reg292022
}
reg292025 = reg292023
} else {
reg292024 := False;
reg292025 = reg292024
}
var reg292028 Obj
if reg292025 == True {
reg292026 := True;
reg292028 = reg292026
} else {
reg292027 := False;
reg292028 = reg292027
}
reg292030 = reg292028
} else {
reg292029 := False;
reg292030 = reg292029
}
if reg292030 == True {
reg292031 := PrimHead(V4023)
reg292032 := PrimTail(V4023)
reg292033 := PrimTail(reg292032)
reg292034 := PrimHead(reg292033)
__ctx.TailApply(__defun__shen_4typecheck_1and_1evaluate, reg292031, reg292034)
return
} else {
reg292036 := PrimIsPair(V4023)
var reg292043 Obj
if reg292036 == True {
reg292037 := PrimTail(V4023)
reg292038 := PrimIsPair(reg292037)
var reg292041 Obj
if reg292038 == True {
reg292039 := True;
reg292041 = reg292039
} else {
reg292040 := False;
reg292041 = reg292040
}
reg292043 = reg292041
} else {
reg292042 := False;
reg292043 = reg292042
}
if reg292043 == True {
reg292044 := PrimHead(V4023)
reg292045 := Nil;
reg292046 := PrimCons(reg292044, reg292045)
reg292047 := __e.Call(__defun__shen_4toplevel__evaluate, reg292046, V4024)
_ = reg292047
reg292048 := MakeNumber(1)
reg292049 := __e.Call(__defun__nl, reg292048)
_ = reg292049
reg292050 := PrimTail(V4023)
__ctx.TailApply(__defun__shen_4toplevel__evaluate, reg292050, V4024)
return
} else {
reg292052 := PrimIsPair(V4023)
var reg292067 Obj
if reg292052 == True {
reg292053 := Nil;
reg292054 := PrimTail(V4023)
reg292055 := PrimEqual(reg292053, reg292054)
var reg292062 Obj
if reg292055 == True {
reg292056 := True;
reg292057 := PrimEqual(reg292056, V4024)
var reg292060 Obj
if reg292057 == True {
reg292058 := True;
reg292060 = reg292058
} else {
reg292059 := False;
reg292060 = reg292059
}
reg292062 = reg292060
} else {
reg292061 := False;
reg292062 = reg292061
}
var reg292065 Obj
if reg292062 == True {
reg292063 := True;
reg292065 = reg292063
} else {
reg292064 := False;
reg292065 = reg292064
}
reg292067 = reg292065
} else {
reg292066 := False;
reg292067 = reg292066
}
if reg292067 == True {
reg292068 := PrimHead(V4023)
reg292069 := MakeSymbol("A")
reg292070 := __e.Call(__defun__gensym, reg292069)
__ctx.TailApply(__defun__shen_4typecheck_1and_1evaluate, reg292068, reg292070)
return
} else {
reg292072 := PrimIsPair(V4023)
var reg292087 Obj
if reg292072 == True {
reg292073 := Nil;
reg292074 := PrimTail(V4023)
reg292075 := PrimEqual(reg292073, reg292074)
var reg292082 Obj
if reg292075 == True {
reg292076 := False;
reg292077 := PrimEqual(reg292076, V4024)
var reg292080 Obj
if reg292077 == True {
reg292078 := True;
reg292080 = reg292078
} else {
reg292079 := False;
reg292080 = reg292079
}
reg292082 = reg292080
} else {
reg292081 := False;
reg292082 = reg292081
}
var reg292085 Obj
if reg292082 == True {
reg292083 := True;
reg292085 = reg292083
} else {
reg292084 := False;
reg292085 = reg292084
}
reg292087 = reg292085
} else {
reg292086 := False;
reg292087 = reg292086
}
if reg292087 == True {
reg292088 := PrimHead(V4023)
reg292089 := __e.Call(__defun__shen_4eval_1without_1macros, reg292088)
Eval := reg292089
_ = Eval
__ctx.TailApply(__defun__print, Eval)
return
} else {
reg292091 := MakeSymbol("shen.toplevel_evaluate")
__ctx.TailApply(__defun__shen_4f__error, reg292091)
return
}
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.toplevel_evaluate", value: __defun__shen_4toplevel__evaluate})

__defun__shen_4typecheck_1and_1evaluate = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4027 := __args[0]
_ = V4027
V4028 := __args[1]
_ = V4028
reg292093 := __e.Call(__defun__shen_4typecheck, V4027, V4028)
Typecheck := reg292093
_ = Typecheck
reg292094 := False;
reg292095 := PrimEqual(Typecheck, reg292094)
if reg292095 == True {
reg292096 := MakeString("type error\n")
reg292097 := PrimSimpleError(reg292096)
__ctx.Return(reg292097)
return
} else {
reg292098 := __e.Call(__defun__shen_4eval_1without_1macros, V4027)
Eval := reg292098
_ = Eval
reg292099 := __e.Call(__defun__shen_4pretty_1type, Typecheck)
Type := reg292099
_ = Type
reg292100 := MakeString(" : ")
reg292101 := MakeString("")
reg292102 := MakeSymbol("shen.r")
reg292103 := __e.Call(__defun__shen_4app, Type, reg292101, reg292102)
reg292104 := PrimStringConcat(reg292100, reg292103)
reg292105 := MakeSymbol("shen.s")
reg292106 := __e.Call(__defun__shen_4app, Eval, reg292104, reg292105)
reg292107 := __e.Call(__defun__stoutput)
__ctx.TailApply(__defun__shen_4prhush, reg292106, reg292107)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.typecheck-and-evaluate", value: __defun__shen_4typecheck_1and_1evaluate})

__defun__shen_4pretty_1type = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4030 := __args[0]
_ = V4030
reg292109 := MakeSymbol("shen.*alphabet*")
reg292110 := PrimValue(reg292109)
reg292111 := __e.Call(__defun__shen_4extract_1pvars, V4030)
__ctx.TailApply(__defun__shen_4mult__subst, reg292110, reg292111, V4030)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.pretty-type", value: __defun__shen_4pretty_1type})

__defun__shen_4extract_1pvars = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4036 := __args[0]
_ = V4036
reg292113 := __e.Call(__defun__shen_4pvar_2, V4036)
if reg292113 == True {
reg292114 := Nil;
reg292115 := PrimCons(V4036, reg292114)
__ctx.Return(reg292115)
return
} else {
reg292116 := PrimIsPair(V4036)
if reg292116 == True {
reg292117 := PrimHead(V4036)
reg292118 := __e.Call(__defun__shen_4extract_1pvars, reg292117)
reg292119 := PrimTail(V4036)
reg292120 := __e.Call(__defun__shen_4extract_1pvars, reg292119)
__ctx.TailApply(__defun__union, reg292118, reg292120)
return
} else {
reg292122 := Nil;
__ctx.Return(reg292122)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.extract-pvars", value: __defun__shen_4extract_1pvars})

__defun__shen_4mult__subst = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V4044 := __args[0]
_ = V4044
V4045 := __args[1]
_ = V4045
V4046 := __args[2]
_ = V4046
reg292123 := Nil;
reg292124 := PrimEqual(reg292123, V4044)
if reg292124 == True {
__ctx.Return(V4046)
return
} else {
reg292125 := Nil;
reg292126 := PrimEqual(reg292125, V4045)
if reg292126 == True {
__ctx.Return(V4046)
return
} else {
reg292127 := PrimIsPair(V4044)
var reg292133 Obj
if reg292127 == True {
reg292128 := PrimIsPair(V4045)
var reg292131 Obj
if reg292128 == True {
reg292129 := True;
reg292131 = reg292129
} else {
reg292130 := False;
reg292131 = reg292130
}
reg292133 = reg292131
} else {
reg292132 := False;
reg292133 = reg292132
}
if reg292133 == True {
reg292134 := PrimTail(V4044)
reg292135 := PrimTail(V4045)
reg292136 := PrimHead(V4044)
reg292137 := PrimHead(V4045)
reg292138 := __e.Call(__defun__subst, reg292136, reg292137, V4046)
__ctx.TailApply(__defun__shen_4mult__subst, reg292134, reg292135, reg292138)
return
} else {
reg292140 := MakeSymbol("shen.mult_subst")
__ctx.TailApply(__defun__shen_4f__error, reg292140)
return
}
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.mult_subst", value: __defun__shen_4mult__subst})

}
