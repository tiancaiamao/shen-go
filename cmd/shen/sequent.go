package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4datatype_1error Obj // shen.datatype-error
var __defun__shen_4_5datatype_1rules_6 Obj // shen.<datatype-rules>
var __defun__shen_4_5datatype_1rule_6 Obj // shen.<datatype-rule>
var __defun__shen_4_5side_1conditions_6 Obj // shen.<side-conditions>
var __defun__shen_4_5side_1condition_6 Obj // shen.<side-condition>
var __defun__shen_4_5variable_2_6 Obj // shen.<variable?>
var __defun__shen_4_5expr_6 Obj // shen.<expr>
var __defun__shen_4remove_1bar Obj // shen.remove-bar
var __defun__shen_4_5premises_6 Obj // shen.<premises>
var __defun__shen_4_5semicolon_1symbol_6 Obj // shen.<semicolon-symbol>
var __defun__shen_4_5premise_6 Obj // shen.<premise>
var __defun__shen_4_5conclusion_6 Obj // shen.<conclusion>
var __defun__shen_4sequent Obj // shen.sequent
var __defun__shen_4_5formulae_6 Obj // shen.<formulae>
var __defun__shen_4_5comma_1symbol_6 Obj // shen.<comma-symbol>
var __defun__shen_4_5formula_6 Obj // shen.<formula>
var __defun__shen_4_5type_6 Obj // shen.<type>
var __defun__shen_4_5doubleunderline_6 Obj // shen.<doubleunderline>
var __defun__shen_4_5singleunderline_6 Obj // shen.<singleunderline>
var __defun__shen_4singleunderline_2 Obj // shen.singleunderline?
var __defun__shen_4sh_2 Obj // shen.sh?
var __defun__shen_4doubleunderline_2 Obj // shen.doubleunderline?
var __defun__shen_4dh_2 Obj // shen.dh?
var __defun__shen_4process_1datatype Obj // shen.process-datatype
var __defun__shen_4remember_1datatype Obj // shen.remember-datatype
var __defun__shen_4rules_1_6horn_1clauses Obj // shen.rules->horn-clauses
var __defun__shen_4double_1_6singles Obj // shen.double->singles
var __defun__shen_4right_1rule Obj // shen.right-rule
var __defun__shen_4left_1rule Obj // shen.left-rule
var __defun__shen_4right_1_6left Obj // shen.right->left
var __defun__shen_4rule_1_6horn_1clause Obj // shen.rule->horn-clause
var __defun__shen_4rule_1_6horn_1clause_1head Obj // shen.rule->horn-clause-head
var __defun__shen_4mode_1ify Obj // shen.mode-ify
var __defun__shen_4rule_1_6horn_1clause_1body Obj // shen.rule->horn-clause-body
var __defun__shen_4construct_1search_1literals Obj // shen.construct-search-literals
var __defun__shen_4csl_1help Obj // shen.csl-help
var __defun__shen_4construct_1search_1clauses Obj // shen.construct-search-clauses
var __defun__shen_4construct_1search_1clause Obj // shen.construct-search-clause
var __defun__shen_4construct_1base_1search_1clause Obj // shen.construct-base-search-clause
var __defun__shen_4construct_1recursive_1search_1clause Obj // shen.construct-recursive-search-clause
var __defun__shen_4construct_1side_1literals Obj // shen.construct-side-literals
var __defun__shen_4construct_1premiss_1literal Obj // shen.construct-premiss-literal
var __defun__shen_4construct_1context Obj // shen.construct-context
var __defun__shen_4recursive__cons__form Obj // shen.recursive_cons_form
var __defun__preclude Obj // preclude
var __defun__shen_4preclude_1h Obj // shen.preclude-h
var __defun__include Obj // include
var __defun__shen_4include_1h Obj // shen.include-h
var __defun__preclude_1all_1but Obj // preclude-all-but
var __defun__include_1all_1but Obj // include-all-but
var __defun__shen_4synonyms_1help Obj // shen.synonyms-help
var __defun__shen_4pushnew Obj // shen.pushnew
var __defun__shen_4demod_1rule Obj // shen.demod-rule
var __defun__shen_4lambda_1of_1defun Obj // shen.lambda-of-defun
var __defun__shen_4update_1demodulation_1function Obj // shen.update-demodulation-function
var __defun__shen_4default_1rule Obj // shen.default-rule

func init() {
__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg297444 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
__ctx.Return(reg297444)
return
}, 0))
__defun__shen_4datatype_1error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2635 := __args[0]
_ = V2635
reg297445 := PrimIsPair(V2635)
var reg297461 Obj
if reg297445 == True {
reg297446 := PrimTail(V2635)
reg297447 := PrimIsPair(reg297446)
var reg297456 Obj
if reg297447 == True {
reg297448 := Nil;
reg297449 := PrimTail(V2635)
reg297450 := PrimTail(reg297449)
reg297451 := PrimEqual(reg297448, reg297450)
var reg297454 Obj
if reg297451 == True {
reg297452 := True;
reg297454 = reg297452
} else {
reg297453 := False;
reg297454 = reg297453
}
reg297456 = reg297454
} else {
reg297455 := False;
reg297456 = reg297455
}
var reg297459 Obj
if reg297456 == True {
reg297457 := True;
reg297459 = reg297457
} else {
reg297458 := False;
reg297459 = reg297458
}
reg297461 = reg297459
} else {
reg297460 := False;
reg297461 = reg297460
}
if reg297461 == True {
reg297462 := MakeString("datatype syntax error here:\n\n ")
reg297463 := MakeNumber(50)
reg297464 := PrimHead(V2635)
reg297465 := __e.Call(__defun__shen_4next_150, reg297463, reg297464)
reg297466 := MakeString("\n")
reg297467 := MakeSymbol("shen.a")
reg297468 := __e.Call(__defun__shen_4app, reg297465, reg297466, reg297467)
reg297469 := PrimStringConcat(reg297462, reg297468)
reg297470 := PrimSimpleError(reg297469)
__ctx.Return(reg297470)
return
} else {
reg297471 := MakeSymbol("shen.datatype-error")
__ctx.TailApply(__defun__shen_4f__error, reg297471)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.datatype-error", value: __defun__shen_4datatype_1error})

__defun__shen_4_5datatype_1rules_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2637 := __args[0]
_ = V2637
reg297473 := __e.Call(__defun__shen_4_5datatype_1rule_6, V2637)
Parse__shen_4_5datatype_1rule_6 := reg297473
_ = Parse__shen_4_5datatype_1rule_6
reg297474 := __e.Call(__defun__fail)
reg297475 := PrimEqual(reg297474, Parse__shen_4_5datatype_1rule_6)
reg297476 := PrimNot(reg297475)
var reg297489 Obj
if reg297476 == True {
reg297477 := __e.Call(__defun__shen_4_5datatype_1rules_6, Parse__shen_4_5datatype_1rule_6)
Parse__shen_4_5datatype_1rules_6 := reg297477
_ = Parse__shen_4_5datatype_1rules_6
reg297478 := __e.Call(__defun__fail)
reg297479 := PrimEqual(reg297478, Parse__shen_4_5datatype_1rules_6)
reg297480 := PrimNot(reg297479)
var reg297487 Obj
if reg297480 == True {
reg297481 := PrimHead(Parse__shen_4_5datatype_1rules_6)
reg297482 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5datatype_1rule_6)
reg297483 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5datatype_1rules_6)
reg297484 := PrimCons(reg297482, reg297483)
reg297485 := __e.Call(__defun__shen_4pair, reg297481, reg297484)
reg297487 = reg297485
} else {
reg297486 := __e.Call(__defun__fail)
reg297487 = reg297486
}
reg297489 = reg297487
} else {
reg297488 := __e.Call(__defun__fail)
reg297489 = reg297488
}
YaccParse := reg297489
_ = YaccParse
reg297490 := __e.Call(__defun__fail)
reg297491 := PrimEqual(YaccParse, reg297490)
if reg297491 == True {
reg297492 := __e.Call(__defun___5e_6, V2637)
Parse___5e_6 := reg297492
_ = Parse___5e_6
reg297493 := __e.Call(__defun__fail)
reg297494 := PrimEqual(reg297493, Parse___5e_6)
reg297495 := PrimNot(reg297494)
if reg297495 == True {
reg297496 := PrimHead(Parse___5e_6)
reg297497 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg297496, reg297497)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<datatype-rules>", value: __defun__shen_4_5datatype_1rules_6})

__defun__shen_4_5datatype_1rule_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2639 := __args[0]
_ = V2639
reg297500 := __e.Call(__defun__shen_4_5side_1conditions_6, V2639)
Parse__shen_4_5side_1conditions_6 := reg297500
_ = Parse__shen_4_5side_1conditions_6
reg297501 := __e.Call(__defun__fail)
reg297502 := PrimEqual(reg297501, Parse__shen_4_5side_1conditions_6)
reg297503 := PrimNot(reg297502)
var reg297534 Obj
if reg297503 == True {
reg297504 := __e.Call(__defun__shen_4_5premises_6, Parse__shen_4_5side_1conditions_6)
Parse__shen_4_5premises_6 := reg297504
_ = Parse__shen_4_5premises_6
reg297505 := __e.Call(__defun__fail)
reg297506 := PrimEqual(reg297505, Parse__shen_4_5premises_6)
reg297507 := PrimNot(reg297506)
var reg297532 Obj
if reg297507 == True {
reg297508 := __e.Call(__defun__shen_4_5singleunderline_6, Parse__shen_4_5premises_6)
Parse__shen_4_5singleunderline_6 := reg297508
_ = Parse__shen_4_5singleunderline_6
reg297509 := __e.Call(__defun__fail)
reg297510 := PrimEqual(reg297509, Parse__shen_4_5singleunderline_6)
reg297511 := PrimNot(reg297510)
var reg297530 Obj
if reg297511 == True {
reg297512 := __e.Call(__defun__shen_4_5conclusion_6, Parse__shen_4_5singleunderline_6)
Parse__shen_4_5conclusion_6 := reg297512
_ = Parse__shen_4_5conclusion_6
reg297513 := __e.Call(__defun__fail)
reg297514 := PrimEqual(reg297513, Parse__shen_4_5conclusion_6)
reg297515 := PrimNot(reg297514)
var reg297528 Obj
if reg297515 == True {
reg297516 := PrimHead(Parse__shen_4_5conclusion_6)
reg297517 := MakeSymbol("shen.single")
reg297518 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5side_1conditions_6)
reg297519 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5premises_6)
reg297520 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5conclusion_6)
reg297521 := Nil;
reg297522 := PrimCons(reg297520, reg297521)
reg297523 := PrimCons(reg297519, reg297522)
reg297524 := PrimCons(reg297518, reg297523)
reg297525 := __e.Call(__defun__shen_4sequent, reg297517, reg297524)
reg297526 := __e.Call(__defun__shen_4pair, reg297516, reg297525)
reg297528 = reg297526
} else {
reg297527 := __e.Call(__defun__fail)
reg297528 = reg297527
}
reg297530 = reg297528
} else {
reg297529 := __e.Call(__defun__fail)
reg297530 = reg297529
}
reg297532 = reg297530
} else {
reg297531 := __e.Call(__defun__fail)
reg297532 = reg297531
}
reg297534 = reg297532
} else {
reg297533 := __e.Call(__defun__fail)
reg297534 = reg297533
}
YaccParse := reg297534
_ = YaccParse
reg297535 := __e.Call(__defun__fail)
reg297536 := PrimEqual(YaccParse, reg297535)
if reg297536 == True {
reg297537 := __e.Call(__defun__shen_4_5side_1conditions_6, V2639)
Parse__shen_4_5side_1conditions_6 := reg297537
_ = Parse__shen_4_5side_1conditions_6
reg297538 := __e.Call(__defun__fail)
reg297539 := PrimEqual(reg297538, Parse__shen_4_5side_1conditions_6)
reg297540 := PrimNot(reg297539)
if reg297540 == True {
reg297541 := __e.Call(__defun__shen_4_5premises_6, Parse__shen_4_5side_1conditions_6)
Parse__shen_4_5premises_6 := reg297541
_ = Parse__shen_4_5premises_6
reg297542 := __e.Call(__defun__fail)
reg297543 := PrimEqual(reg297542, Parse__shen_4_5premises_6)
reg297544 := PrimNot(reg297543)
if reg297544 == True {
reg297545 := __e.Call(__defun__shen_4_5doubleunderline_6, Parse__shen_4_5premises_6)
Parse__shen_4_5doubleunderline_6 := reg297545
_ = Parse__shen_4_5doubleunderline_6
reg297546 := __e.Call(__defun__fail)
reg297547 := PrimEqual(reg297546, Parse__shen_4_5doubleunderline_6)
reg297548 := PrimNot(reg297547)
if reg297548 == True {
reg297549 := __e.Call(__defun__shen_4_5conclusion_6, Parse__shen_4_5doubleunderline_6)
Parse__shen_4_5conclusion_6 := reg297549
_ = Parse__shen_4_5conclusion_6
reg297550 := __e.Call(__defun__fail)
reg297551 := PrimEqual(reg297550, Parse__shen_4_5conclusion_6)
reg297552 := PrimNot(reg297551)
if reg297552 == True {
reg297553 := PrimHead(Parse__shen_4_5conclusion_6)
reg297554 := MakeSymbol("shen.double")
reg297555 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5side_1conditions_6)
reg297556 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5premises_6)
reg297557 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5conclusion_6)
reg297558 := Nil;
reg297559 := PrimCons(reg297557, reg297558)
reg297560 := PrimCons(reg297556, reg297559)
reg297561 := PrimCons(reg297555, reg297560)
reg297562 := __e.Call(__defun__shen_4sequent, reg297554, reg297561)
__ctx.TailApply(__defun__shen_4pair, reg297553, reg297562)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<datatype-rule>", value: __defun__shen_4_5datatype_1rule_6})

__defun__shen_4_5side_1conditions_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2641 := __args[0]
_ = V2641
reg297568 := __e.Call(__defun__shen_4_5side_1condition_6, V2641)
Parse__shen_4_5side_1condition_6 := reg297568
_ = Parse__shen_4_5side_1condition_6
reg297569 := __e.Call(__defun__fail)
reg297570 := PrimEqual(reg297569, Parse__shen_4_5side_1condition_6)
reg297571 := PrimNot(reg297570)
var reg297584 Obj
if reg297571 == True {
reg297572 := __e.Call(__defun__shen_4_5side_1conditions_6, Parse__shen_4_5side_1condition_6)
Parse__shen_4_5side_1conditions_6 := reg297572
_ = Parse__shen_4_5side_1conditions_6
reg297573 := __e.Call(__defun__fail)
reg297574 := PrimEqual(reg297573, Parse__shen_4_5side_1conditions_6)
reg297575 := PrimNot(reg297574)
var reg297582 Obj
if reg297575 == True {
reg297576 := PrimHead(Parse__shen_4_5side_1conditions_6)
reg297577 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5side_1condition_6)
reg297578 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5side_1conditions_6)
reg297579 := PrimCons(reg297577, reg297578)
reg297580 := __e.Call(__defun__shen_4pair, reg297576, reg297579)
reg297582 = reg297580
} else {
reg297581 := __e.Call(__defun__fail)
reg297582 = reg297581
}
reg297584 = reg297582
} else {
reg297583 := __e.Call(__defun__fail)
reg297584 = reg297583
}
YaccParse := reg297584
_ = YaccParse
reg297585 := __e.Call(__defun__fail)
reg297586 := PrimEqual(YaccParse, reg297585)
if reg297586 == True {
reg297587 := __e.Call(__defun___5e_6, V2641)
Parse___5e_6 := reg297587
_ = Parse___5e_6
reg297588 := __e.Call(__defun__fail)
reg297589 := PrimEqual(reg297588, Parse___5e_6)
reg297590 := PrimNot(reg297589)
if reg297590 == True {
reg297591 := PrimHead(Parse___5e_6)
reg297592 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg297591, reg297592)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<side-conditions>", value: __defun__shen_4_5side_1conditions_6})

__defun__shen_4_5side_1condition_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2643 := __args[0]
_ = V2643
reg297595 := PrimHead(V2643)
reg297596 := PrimIsPair(reg297595)
var reg297605 Obj
if reg297596 == True {
reg297597 := MakeSymbol("if")
reg297598 := PrimHead(V2643)
reg297599 := PrimHead(reg297598)
reg297600 := PrimEqual(reg297597, reg297599)
var reg297603 Obj
if reg297600 == True {
reg297601 := True;
reg297603 = reg297601
} else {
reg297602 := False;
reg297603 = reg297602
}
reg297605 = reg297603
} else {
reg297604 := False;
reg297605 = reg297604
}
var reg297624 Obj
if reg297605 == True {
reg297606 := PrimHead(V2643)
reg297607 := PrimTail(reg297606)
reg297608 := __e.Call(__defun__shen_4hdtl, V2643)
reg297609 := __e.Call(__defun__shen_4pair, reg297607, reg297608)
reg297610 := __e.Call(__defun__shen_4_5expr_6, reg297609)
Parse__shen_4_5expr_6 := reg297610
_ = Parse__shen_4_5expr_6
reg297611 := __e.Call(__defun__fail)
reg297612 := PrimEqual(reg297611, Parse__shen_4_5expr_6)
reg297613 := PrimNot(reg297612)
var reg297622 Obj
if reg297613 == True {
reg297614 := PrimHead(Parse__shen_4_5expr_6)
reg297615 := MakeSymbol("if")
reg297616 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5expr_6)
reg297617 := Nil;
reg297618 := PrimCons(reg297616, reg297617)
reg297619 := PrimCons(reg297615, reg297618)
reg297620 := __e.Call(__defun__shen_4pair, reg297614, reg297619)
reg297622 = reg297620
} else {
reg297621 := __e.Call(__defun__fail)
reg297622 = reg297621
}
reg297624 = reg297622
} else {
reg297623 := __e.Call(__defun__fail)
reg297624 = reg297623
}
YaccParse := reg297624
_ = YaccParse
reg297625 := __e.Call(__defun__fail)
reg297626 := PrimEqual(YaccParse, reg297625)
if reg297626 == True {
reg297627 := PrimHead(V2643)
reg297628 := PrimIsPair(reg297627)
var reg297637 Obj
if reg297628 == True {
reg297629 := MakeSymbol("let")
reg297630 := PrimHead(V2643)
reg297631 := PrimHead(reg297630)
reg297632 := PrimEqual(reg297629, reg297631)
var reg297635 Obj
if reg297632 == True {
reg297633 := True;
reg297635 = reg297633
} else {
reg297634 := False;
reg297635 = reg297634
}
reg297637 = reg297635
} else {
reg297636 := False;
reg297637 = reg297636
}
if reg297637 == True {
reg297638 := PrimHead(V2643)
reg297639 := PrimTail(reg297638)
reg297640 := __e.Call(__defun__shen_4hdtl, V2643)
reg297641 := __e.Call(__defun__shen_4pair, reg297639, reg297640)
reg297642 := __e.Call(__defun__shen_4_5variable_2_6, reg297641)
Parse__shen_4_5variable_2_6 := reg297642
_ = Parse__shen_4_5variable_2_6
reg297643 := __e.Call(__defun__fail)
reg297644 := PrimEqual(reg297643, Parse__shen_4_5variable_2_6)
reg297645 := PrimNot(reg297644)
if reg297645 == True {
reg297646 := __e.Call(__defun__shen_4_5expr_6, Parse__shen_4_5variable_2_6)
Parse__shen_4_5expr_6 := reg297646
_ = Parse__shen_4_5expr_6
reg297647 := __e.Call(__defun__fail)
reg297648 := PrimEqual(reg297647, Parse__shen_4_5expr_6)
reg297649 := PrimNot(reg297648)
if reg297649 == True {
reg297650 := PrimHead(Parse__shen_4_5expr_6)
reg297651 := MakeSymbol("let")
reg297652 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5variable_2_6)
reg297653 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5expr_6)
reg297654 := Nil;
reg297655 := PrimCons(reg297653, reg297654)
reg297656 := PrimCons(reg297652, reg297655)
reg297657 := PrimCons(reg297651, reg297656)
__ctx.TailApply(__defun__shen_4pair, reg297650, reg297657)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<side-condition>", value: __defun__shen_4_5side_1condition_6})

__defun__shen_4_5variable_2_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2645 := __args[0]
_ = V2645
reg297662 := PrimHead(V2645)
reg297663 := PrimIsPair(reg297662)
if reg297663 == True {
reg297664 := PrimHead(V2645)
reg297665 := PrimHead(reg297664)
Parse__X := reg297665
_ = Parse__X
reg297666 := PrimIsVariable(Parse__X)
if reg297666 == True {
reg297667 := PrimHead(V2645)
reg297668 := PrimTail(reg297667)
reg297669 := __e.Call(__defun__shen_4hdtl, V2645)
reg297670 := __e.Call(__defun__shen_4pair, reg297668, reg297669)
reg297671 := PrimHead(reg297670)
__ctx.TailApply(__defun__shen_4pair, reg297671, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<variable?>", value: __defun__shen_4_5variable_2_6})

__defun__shen_4_5expr_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2647 := __args[0]
_ = V2647
reg297675 := PrimHead(V2647)
reg297676 := PrimIsPair(reg297675)
if reg297676 == True {
reg297677 := PrimHead(V2647)
reg297678 := PrimHead(reg297677)
Parse__X := reg297678
_ = Parse__X
reg297679 := MakeSymbol(">>")
reg297680 := MakeSymbol(";")
reg297681 := Nil;
reg297682 := PrimCons(reg297680, reg297681)
reg297683 := PrimCons(reg297679, reg297682)
reg297684 := __e.Call(__defun__element_2, Parse__X, reg297683)
var reg297696 Obj
if reg297684 == True {
reg297685 := True;
reg297696 = reg297685
} else {
reg297686 := __e.Call(__defun__shen_4singleunderline_2, Parse__X)
var reg297692 Obj
if reg297686 == True {
reg297687 := True;
reg297692 = reg297687
} else {
reg297688 := __e.Call(__defun__shen_4doubleunderline_2, Parse__X)
var reg297691 Obj
if reg297688 == True {
reg297689 := True;
reg297691 = reg297689
} else {
reg297690 := False;
reg297691 = reg297690
}
reg297692 = reg297691
}
var reg297695 Obj
if reg297692 == True {
reg297693 := True;
reg297695 = reg297693
} else {
reg297694 := False;
reg297695 = reg297694
}
reg297696 = reg297695
}
reg297697 := PrimNot(reg297696)
if reg297697 == True {
reg297698 := PrimHead(V2647)
reg297699 := PrimTail(reg297698)
reg297700 := __e.Call(__defun__shen_4hdtl, V2647)
reg297701 := __e.Call(__defun__shen_4pair, reg297699, reg297700)
reg297702 := PrimHead(reg297701)
reg297703 := __e.Call(__defun__shen_4remove_1bar, Parse__X)
__ctx.TailApply(__defun__shen_4pair, reg297702, reg297703)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<expr>", value: __defun__shen_4_5expr_6})

__defun__shen_4remove_1bar = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2649 := __args[0]
_ = V2649
reg297707 := PrimIsPair(V2649)
var reg297741 Obj
if reg297707 == True {
reg297708 := PrimTail(V2649)
reg297709 := PrimIsPair(reg297708)
var reg297736 Obj
if reg297709 == True {
reg297710 := PrimTail(V2649)
reg297711 := PrimTail(reg297710)
reg297712 := PrimIsPair(reg297711)
var reg297731 Obj
if reg297712 == True {
reg297713 := Nil;
reg297714 := PrimTail(V2649)
reg297715 := PrimTail(reg297714)
reg297716 := PrimTail(reg297715)
reg297717 := PrimEqual(reg297713, reg297716)
var reg297726 Obj
if reg297717 == True {
reg297718 := PrimTail(V2649)
reg297719 := PrimHead(reg297718)
reg297720 := MakeSymbol("bar!")
reg297721 := PrimEqual(reg297719, reg297720)
var reg297724 Obj
if reg297721 == True {
reg297722 := True;
reg297724 = reg297722
} else {
reg297723 := False;
reg297724 = reg297723
}
reg297726 = reg297724
} else {
reg297725 := False;
reg297726 = reg297725
}
var reg297729 Obj
if reg297726 == True {
reg297727 := True;
reg297729 = reg297727
} else {
reg297728 := False;
reg297729 = reg297728
}
reg297731 = reg297729
} else {
reg297730 := False;
reg297731 = reg297730
}
var reg297734 Obj
if reg297731 == True {
reg297732 := True;
reg297734 = reg297732
} else {
reg297733 := False;
reg297734 = reg297733
}
reg297736 = reg297734
} else {
reg297735 := False;
reg297736 = reg297735
}
var reg297739 Obj
if reg297736 == True {
reg297737 := True;
reg297739 = reg297737
} else {
reg297738 := False;
reg297739 = reg297738
}
reg297741 = reg297739
} else {
reg297740 := False;
reg297741 = reg297740
}
if reg297741 == True {
reg297742 := PrimHead(V2649)
reg297743 := PrimTail(V2649)
reg297744 := PrimTail(reg297743)
reg297745 := PrimHead(reg297744)
reg297746 := PrimCons(reg297742, reg297745)
__ctx.Return(reg297746)
return
} else {
reg297747 := PrimIsPair(V2649)
if reg297747 == True {
reg297748 := PrimHead(V2649)
reg297749 := __e.Call(__defun__shen_4remove_1bar, reg297748)
reg297750 := PrimTail(V2649)
reg297751 := __e.Call(__defun__shen_4remove_1bar, reg297750)
reg297752 := PrimCons(reg297749, reg297751)
__ctx.Return(reg297752)
return
} else {
__ctx.Return(V2649)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.remove-bar", value: __defun__shen_4remove_1bar})

__defun__shen_4_5premises_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2651 := __args[0]
_ = V2651
reg297753 := __e.Call(__defun__shen_4_5premise_6, V2651)
Parse__shen_4_5premise_6 := reg297753
_ = Parse__shen_4_5premise_6
reg297754 := __e.Call(__defun__fail)
reg297755 := PrimEqual(reg297754, Parse__shen_4_5premise_6)
reg297756 := PrimNot(reg297755)
var reg297775 Obj
if reg297756 == True {
reg297757 := __e.Call(__defun__shen_4_5semicolon_1symbol_6, Parse__shen_4_5premise_6)
Parse__shen_4_5semicolon_1symbol_6 := reg297757
_ = Parse__shen_4_5semicolon_1symbol_6
reg297758 := __e.Call(__defun__fail)
reg297759 := PrimEqual(reg297758, Parse__shen_4_5semicolon_1symbol_6)
reg297760 := PrimNot(reg297759)
var reg297773 Obj
if reg297760 == True {
reg297761 := __e.Call(__defun__shen_4_5premises_6, Parse__shen_4_5semicolon_1symbol_6)
Parse__shen_4_5premises_6 := reg297761
_ = Parse__shen_4_5premises_6
reg297762 := __e.Call(__defun__fail)
reg297763 := PrimEqual(reg297762, Parse__shen_4_5premises_6)
reg297764 := PrimNot(reg297763)
var reg297771 Obj
if reg297764 == True {
reg297765 := PrimHead(Parse__shen_4_5premises_6)
reg297766 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5premise_6)
reg297767 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5premises_6)
reg297768 := PrimCons(reg297766, reg297767)
reg297769 := __e.Call(__defun__shen_4pair, reg297765, reg297768)
reg297771 = reg297769
} else {
reg297770 := __e.Call(__defun__fail)
reg297771 = reg297770
}
reg297773 = reg297771
} else {
reg297772 := __e.Call(__defun__fail)
reg297773 = reg297772
}
reg297775 = reg297773
} else {
reg297774 := __e.Call(__defun__fail)
reg297775 = reg297774
}
YaccParse := reg297775
_ = YaccParse
reg297776 := __e.Call(__defun__fail)
reg297777 := PrimEqual(YaccParse, reg297776)
if reg297777 == True {
reg297778 := __e.Call(__defun___5e_6, V2651)
Parse___5e_6 := reg297778
_ = Parse___5e_6
reg297779 := __e.Call(__defun__fail)
reg297780 := PrimEqual(reg297779, Parse___5e_6)
reg297781 := PrimNot(reg297780)
if reg297781 == True {
reg297782 := PrimHead(Parse___5e_6)
reg297783 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg297782, reg297783)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<premises>", value: __defun__shen_4_5premises_6})

__defun__shen_4_5semicolon_1symbol_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2653 := __args[0]
_ = V2653
reg297786 := PrimHead(V2653)
reg297787 := PrimIsPair(reg297786)
if reg297787 == True {
reg297788 := PrimHead(V2653)
reg297789 := PrimHead(reg297788)
Parse__X := reg297789
_ = Parse__X
reg297790 := MakeSymbol(";")
reg297791 := PrimEqual(Parse__X, reg297790)
if reg297791 == True {
reg297792 := PrimHead(V2653)
reg297793 := PrimTail(reg297792)
reg297794 := __e.Call(__defun__shen_4hdtl, V2653)
reg297795 := __e.Call(__defun__shen_4pair, reg297793, reg297794)
reg297796 := PrimHead(reg297795)
reg297797 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg297796, reg297797)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<semicolon-symbol>", value: __defun__shen_4_5semicolon_1symbol_6})

__defun__shen_4_5premise_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2655 := __args[0]
_ = V2655
reg297801 := PrimHead(V2655)
reg297802 := PrimIsPair(reg297801)
var reg297811 Obj
if reg297802 == True {
reg297803 := MakeSymbol("!")
reg297804 := PrimHead(V2655)
reg297805 := PrimHead(reg297804)
reg297806 := PrimEqual(reg297803, reg297805)
var reg297809 Obj
if reg297806 == True {
reg297807 := True;
reg297809 = reg297807
} else {
reg297808 := False;
reg297809 = reg297808
}
reg297811 = reg297809
} else {
reg297810 := False;
reg297811 = reg297810
}
var reg297820 Obj
if reg297811 == True {
reg297812 := PrimHead(V2655)
reg297813 := PrimTail(reg297812)
reg297814 := __e.Call(__defun__shen_4hdtl, V2655)
reg297815 := __e.Call(__defun__shen_4pair, reg297813, reg297814)
reg297816 := PrimHead(reg297815)
reg297817 := MakeSymbol("!")
reg297818 := __e.Call(__defun__shen_4pair, reg297816, reg297817)
reg297820 = reg297818
} else {
reg297819 := __e.Call(__defun__fail)
reg297820 = reg297819
}
YaccParse := reg297820
_ = YaccParse
reg297821 := __e.Call(__defun__fail)
reg297822 := PrimEqual(YaccParse, reg297821)
if reg297822 == True {
reg297823 := __e.Call(__defun__shen_4_5formulae_6, V2655)
Parse__shen_4_5formulae_6 := reg297823
_ = Parse__shen_4_5formulae_6
reg297824 := __e.Call(__defun__fail)
reg297825 := PrimEqual(reg297824, Parse__shen_4_5formulae_6)
reg297826 := PrimNot(reg297825)
var reg297856 Obj
if reg297826 == True {
reg297827 := PrimHead(Parse__shen_4_5formulae_6)
reg297828 := PrimIsPair(reg297827)
var reg297837 Obj
if reg297828 == True {
reg297829 := MakeSymbol(">>")
reg297830 := PrimHead(Parse__shen_4_5formulae_6)
reg297831 := PrimHead(reg297830)
reg297832 := PrimEqual(reg297829, reg297831)
var reg297835 Obj
if reg297832 == True {
reg297833 := True;
reg297835 = reg297833
} else {
reg297834 := False;
reg297835 = reg297834
}
reg297837 = reg297835
} else {
reg297836 := False;
reg297837 = reg297836
}
var reg297854 Obj
if reg297837 == True {
reg297838 := PrimHead(Parse__shen_4_5formulae_6)
reg297839 := PrimTail(reg297838)
reg297840 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formulae_6)
reg297841 := __e.Call(__defun__shen_4pair, reg297839, reg297840)
reg297842 := __e.Call(__defun__shen_4_5formula_6, reg297841)
Parse__shen_4_5formula_6 := reg297842
_ = Parse__shen_4_5formula_6
reg297843 := __e.Call(__defun__fail)
reg297844 := PrimEqual(reg297843, Parse__shen_4_5formula_6)
reg297845 := PrimNot(reg297844)
var reg297852 Obj
if reg297845 == True {
reg297846 := PrimHead(Parse__shen_4_5formula_6)
reg297847 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formulae_6)
reg297848 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formula_6)
reg297849 := __e.Call(__defun__shen_4sequent, reg297847, reg297848)
reg297850 := __e.Call(__defun__shen_4pair, reg297846, reg297849)
reg297852 = reg297850
} else {
reg297851 := __e.Call(__defun__fail)
reg297852 = reg297851
}
reg297854 = reg297852
} else {
reg297853 := __e.Call(__defun__fail)
reg297854 = reg297853
}
reg297856 = reg297854
} else {
reg297855 := __e.Call(__defun__fail)
reg297856 = reg297855
}
YaccParse := reg297856
_ = YaccParse
reg297857 := __e.Call(__defun__fail)
reg297858 := PrimEqual(YaccParse, reg297857)
if reg297858 == True {
reg297859 := __e.Call(__defun__shen_4_5formula_6, V2655)
Parse__shen_4_5formula_6 := reg297859
_ = Parse__shen_4_5formula_6
reg297860 := __e.Call(__defun__fail)
reg297861 := PrimEqual(reg297860, Parse__shen_4_5formula_6)
reg297862 := PrimNot(reg297861)
if reg297862 == True {
reg297863 := PrimHead(Parse__shen_4_5formula_6)
reg297864 := Nil;
reg297865 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formula_6)
reg297866 := __e.Call(__defun__shen_4sequent, reg297864, reg297865)
__ctx.TailApply(__defun__shen_4pair, reg297863, reg297866)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<premise>", value: __defun__shen_4_5premise_6})

__defun__shen_4_5conclusion_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2657 := __args[0]
_ = V2657
reg297869 := __e.Call(__defun__shen_4_5formulae_6, V2657)
Parse__shen_4_5formulae_6 := reg297869
_ = Parse__shen_4_5formulae_6
reg297870 := __e.Call(__defun__fail)
reg297871 := PrimEqual(reg297870, Parse__shen_4_5formulae_6)
reg297872 := PrimNot(reg297871)
var reg297908 Obj
if reg297872 == True {
reg297873 := PrimHead(Parse__shen_4_5formulae_6)
reg297874 := PrimIsPair(reg297873)
var reg297883 Obj
if reg297874 == True {
reg297875 := MakeSymbol(">>")
reg297876 := PrimHead(Parse__shen_4_5formulae_6)
reg297877 := PrimHead(reg297876)
reg297878 := PrimEqual(reg297875, reg297877)
var reg297881 Obj
if reg297878 == True {
reg297879 := True;
reg297881 = reg297879
} else {
reg297880 := False;
reg297881 = reg297880
}
reg297883 = reg297881
} else {
reg297882 := False;
reg297883 = reg297882
}
var reg297906 Obj
if reg297883 == True {
reg297884 := PrimHead(Parse__shen_4_5formulae_6)
reg297885 := PrimTail(reg297884)
reg297886 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formulae_6)
reg297887 := __e.Call(__defun__shen_4pair, reg297885, reg297886)
reg297888 := __e.Call(__defun__shen_4_5formula_6, reg297887)
Parse__shen_4_5formula_6 := reg297888
_ = Parse__shen_4_5formula_6
reg297889 := __e.Call(__defun__fail)
reg297890 := PrimEqual(reg297889, Parse__shen_4_5formula_6)
reg297891 := PrimNot(reg297890)
var reg297904 Obj
if reg297891 == True {
reg297892 := __e.Call(__defun__shen_4_5semicolon_1symbol_6, Parse__shen_4_5formula_6)
Parse__shen_4_5semicolon_1symbol_6 := reg297892
_ = Parse__shen_4_5semicolon_1symbol_6
reg297893 := __e.Call(__defun__fail)
reg297894 := PrimEqual(reg297893, Parse__shen_4_5semicolon_1symbol_6)
reg297895 := PrimNot(reg297894)
var reg297902 Obj
if reg297895 == True {
reg297896 := PrimHead(Parse__shen_4_5semicolon_1symbol_6)
reg297897 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formulae_6)
reg297898 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formula_6)
reg297899 := __e.Call(__defun__shen_4sequent, reg297897, reg297898)
reg297900 := __e.Call(__defun__shen_4pair, reg297896, reg297899)
reg297902 = reg297900
} else {
reg297901 := __e.Call(__defun__fail)
reg297902 = reg297901
}
reg297904 = reg297902
} else {
reg297903 := __e.Call(__defun__fail)
reg297904 = reg297903
}
reg297906 = reg297904
} else {
reg297905 := __e.Call(__defun__fail)
reg297906 = reg297905
}
reg297908 = reg297906
} else {
reg297907 := __e.Call(__defun__fail)
reg297908 = reg297907
}
YaccParse := reg297908
_ = YaccParse
reg297909 := __e.Call(__defun__fail)
reg297910 := PrimEqual(YaccParse, reg297909)
if reg297910 == True {
reg297911 := __e.Call(__defun__shen_4_5formula_6, V2657)
Parse__shen_4_5formula_6 := reg297911
_ = Parse__shen_4_5formula_6
reg297912 := __e.Call(__defun__fail)
reg297913 := PrimEqual(reg297912, Parse__shen_4_5formula_6)
reg297914 := PrimNot(reg297913)
if reg297914 == True {
reg297915 := __e.Call(__defun__shen_4_5semicolon_1symbol_6, Parse__shen_4_5formula_6)
Parse__shen_4_5semicolon_1symbol_6 := reg297915
_ = Parse__shen_4_5semicolon_1symbol_6
reg297916 := __e.Call(__defun__fail)
reg297917 := PrimEqual(reg297916, Parse__shen_4_5semicolon_1symbol_6)
reg297918 := PrimNot(reg297917)
if reg297918 == True {
reg297919 := PrimHead(Parse__shen_4_5semicolon_1symbol_6)
reg297920 := Nil;
reg297921 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formula_6)
reg297922 := __e.Call(__defun__shen_4sequent, reg297920, reg297921)
__ctx.TailApply(__defun__shen_4pair, reg297919, reg297922)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<conclusion>", value: __defun__shen_4_5conclusion_6})

__defun__shen_4sequent = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2660 := __args[0]
_ = V2660
V2661 := __args[1]
_ = V2661
__ctx.TailApply(__defun___8p, V2660, V2661)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.sequent", value: __defun__shen_4sequent})

__defun__shen_4_5formulae_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2663 := __args[0]
_ = V2663
reg297927 := __e.Call(__defun__shen_4_5formula_6, V2663)
Parse__shen_4_5formula_6 := reg297927
_ = Parse__shen_4_5formula_6
reg297928 := __e.Call(__defun__fail)
reg297929 := PrimEqual(reg297928, Parse__shen_4_5formula_6)
reg297930 := PrimNot(reg297929)
var reg297949 Obj
if reg297930 == True {
reg297931 := __e.Call(__defun__shen_4_5comma_1symbol_6, Parse__shen_4_5formula_6)
Parse__shen_4_5comma_1symbol_6 := reg297931
_ = Parse__shen_4_5comma_1symbol_6
reg297932 := __e.Call(__defun__fail)
reg297933 := PrimEqual(reg297932, Parse__shen_4_5comma_1symbol_6)
reg297934 := PrimNot(reg297933)
var reg297947 Obj
if reg297934 == True {
reg297935 := __e.Call(__defun__shen_4_5formulae_6, Parse__shen_4_5comma_1symbol_6)
Parse__shen_4_5formulae_6 := reg297935
_ = Parse__shen_4_5formulae_6
reg297936 := __e.Call(__defun__fail)
reg297937 := PrimEqual(reg297936, Parse__shen_4_5formulae_6)
reg297938 := PrimNot(reg297937)
var reg297945 Obj
if reg297938 == True {
reg297939 := PrimHead(Parse__shen_4_5formulae_6)
reg297940 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formula_6)
reg297941 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formulae_6)
reg297942 := PrimCons(reg297940, reg297941)
reg297943 := __e.Call(__defun__shen_4pair, reg297939, reg297942)
reg297945 = reg297943
} else {
reg297944 := __e.Call(__defun__fail)
reg297945 = reg297944
}
reg297947 = reg297945
} else {
reg297946 := __e.Call(__defun__fail)
reg297947 = reg297946
}
reg297949 = reg297947
} else {
reg297948 := __e.Call(__defun__fail)
reg297949 = reg297948
}
YaccParse := reg297949
_ = YaccParse
reg297950 := __e.Call(__defun__fail)
reg297951 := PrimEqual(YaccParse, reg297950)
if reg297951 == True {
reg297952 := __e.Call(__defun__shen_4_5formula_6, V2663)
Parse__shen_4_5formula_6 := reg297952
_ = Parse__shen_4_5formula_6
reg297953 := __e.Call(__defun__fail)
reg297954 := PrimEqual(reg297953, Parse__shen_4_5formula_6)
reg297955 := PrimNot(reg297954)
var reg297962 Obj
if reg297955 == True {
reg297956 := PrimHead(Parse__shen_4_5formula_6)
reg297957 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5formula_6)
reg297958 := Nil;
reg297959 := PrimCons(reg297957, reg297958)
reg297960 := __e.Call(__defun__shen_4pair, reg297956, reg297959)
reg297962 = reg297960
} else {
reg297961 := __e.Call(__defun__fail)
reg297962 = reg297961
}
YaccParse := reg297962
_ = YaccParse
reg297963 := __e.Call(__defun__fail)
reg297964 := PrimEqual(YaccParse, reg297963)
if reg297964 == True {
reg297965 := __e.Call(__defun___5e_6, V2663)
Parse___5e_6 := reg297965
_ = Parse___5e_6
reg297966 := __e.Call(__defun__fail)
reg297967 := PrimEqual(reg297966, Parse___5e_6)
reg297968 := PrimNot(reg297967)
if reg297968 == True {
reg297969 := PrimHead(Parse___5e_6)
reg297970 := Nil;
__ctx.TailApply(__defun__shen_4pair, reg297969, reg297970)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<formulae>", value: __defun__shen_4_5formulae_6})

__defun__shen_4_5comma_1symbol_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2665 := __args[0]
_ = V2665
reg297973 := PrimHead(V2665)
reg297974 := PrimIsPair(reg297973)
if reg297974 == True {
reg297975 := PrimHead(V2665)
reg297976 := PrimHead(reg297975)
Parse__X := reg297976
_ = Parse__X
reg297977 := MakeString(",")
reg297978 := PrimIntern(reg297977)
reg297979 := PrimEqual(Parse__X, reg297978)
if reg297979 == True {
reg297980 := PrimHead(V2665)
reg297981 := PrimTail(reg297980)
reg297982 := __e.Call(__defun__shen_4hdtl, V2665)
reg297983 := __e.Call(__defun__shen_4pair, reg297981, reg297982)
reg297984 := PrimHead(reg297983)
reg297985 := MakeSymbol("shen.skip")
__ctx.TailApply(__defun__shen_4pair, reg297984, reg297985)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<comma-symbol>", value: __defun__shen_4_5comma_1symbol_6})

__defun__shen_4_5formula_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2667 := __args[0]
_ = V2667
reg297989 := __e.Call(__defun__shen_4_5expr_6, V2667)
Parse__shen_4_5expr_6 := reg297989
_ = Parse__shen_4_5expr_6
reg297990 := __e.Call(__defun__fail)
reg297991 := PrimEqual(reg297990, Parse__shen_4_5expr_6)
reg297992 := PrimNot(reg297991)
var reg298028 Obj
if reg297992 == True {
reg297993 := PrimHead(Parse__shen_4_5expr_6)
reg297994 := PrimIsPair(reg297993)
var reg298003 Obj
if reg297994 == True {
reg297995 := MakeSymbol(":")
reg297996 := PrimHead(Parse__shen_4_5expr_6)
reg297997 := PrimHead(reg297996)
reg297998 := PrimEqual(reg297995, reg297997)
var reg298001 Obj
if reg297998 == True {
reg297999 := True;
reg298001 = reg297999
} else {
reg298000 := False;
reg298001 = reg298000
}
reg298003 = reg298001
} else {
reg298002 := False;
reg298003 = reg298002
}
var reg298026 Obj
if reg298003 == True {
reg298004 := PrimHead(Parse__shen_4_5expr_6)
reg298005 := PrimTail(reg298004)
reg298006 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5expr_6)
reg298007 := __e.Call(__defun__shen_4pair, reg298005, reg298006)
reg298008 := __e.Call(__defun__shen_4_5type_6, reg298007)
Parse__shen_4_5type_6 := reg298008
_ = Parse__shen_4_5type_6
reg298009 := __e.Call(__defun__fail)
reg298010 := PrimEqual(reg298009, Parse__shen_4_5type_6)
reg298011 := PrimNot(reg298010)
var reg298024 Obj
if reg298011 == True {
reg298012 := PrimHead(Parse__shen_4_5type_6)
reg298013 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5expr_6)
reg298014 := __e.Call(__defun__shen_4curry, reg298013)
reg298015 := MakeSymbol(":")
reg298016 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5type_6)
reg298017 := __e.Call(__defun__shen_4demodulate, reg298016)
reg298018 := Nil;
reg298019 := PrimCons(reg298017, reg298018)
reg298020 := PrimCons(reg298015, reg298019)
reg298021 := PrimCons(reg298014, reg298020)
reg298022 := __e.Call(__defun__shen_4pair, reg298012, reg298021)
reg298024 = reg298022
} else {
reg298023 := __e.Call(__defun__fail)
reg298024 = reg298023
}
reg298026 = reg298024
} else {
reg298025 := __e.Call(__defun__fail)
reg298026 = reg298025
}
reg298028 = reg298026
} else {
reg298027 := __e.Call(__defun__fail)
reg298028 = reg298027
}
YaccParse := reg298028
_ = YaccParse
reg298029 := __e.Call(__defun__fail)
reg298030 := PrimEqual(YaccParse, reg298029)
if reg298030 == True {
reg298031 := __e.Call(__defun__shen_4_5expr_6, V2667)
Parse__shen_4_5expr_6 := reg298031
_ = Parse__shen_4_5expr_6
reg298032 := __e.Call(__defun__fail)
reg298033 := PrimEqual(reg298032, Parse__shen_4_5expr_6)
reg298034 := PrimNot(reg298033)
if reg298034 == True {
reg298035 := PrimHead(Parse__shen_4_5expr_6)
reg298036 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5expr_6)
__ctx.TailApply(__defun__shen_4pair, reg298035, reg298036)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.Return(YaccParse)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<formula>", value: __defun__shen_4_5formula_6})

__defun__shen_4_5type_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2669 := __args[0]
_ = V2669
reg298039 := __e.Call(__defun__shen_4_5expr_6, V2669)
Parse__shen_4_5expr_6 := reg298039
_ = Parse__shen_4_5expr_6
reg298040 := __e.Call(__defun__fail)
reg298041 := PrimEqual(reg298040, Parse__shen_4_5expr_6)
reg298042 := PrimNot(reg298041)
if reg298042 == True {
reg298043 := PrimHead(Parse__shen_4_5expr_6)
reg298044 := __e.Call(__defun__shen_4hdtl, Parse__shen_4_5expr_6)
reg298045 := __e.Call(__defun__shen_4curry_1type, reg298044)
__ctx.TailApply(__defun__shen_4pair, reg298043, reg298045)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<type>", value: __defun__shen_4_5type_6})

__defun__shen_4_5doubleunderline_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2671 := __args[0]
_ = V2671
reg298048 := PrimHead(V2671)
reg298049 := PrimIsPair(reg298048)
if reg298049 == True {
reg298050 := PrimHead(V2671)
reg298051 := PrimHead(reg298050)
Parse__X := reg298051
_ = Parse__X
reg298052 := __e.Call(__defun__shen_4doubleunderline_2, Parse__X)
if reg298052 == True {
reg298053 := PrimHead(V2671)
reg298054 := PrimTail(reg298053)
reg298055 := __e.Call(__defun__shen_4hdtl, V2671)
reg298056 := __e.Call(__defun__shen_4pair, reg298054, reg298055)
reg298057 := PrimHead(reg298056)
__ctx.TailApply(__defun__shen_4pair, reg298057, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<doubleunderline>", value: __defun__shen_4_5doubleunderline_6})

__defun__shen_4_5singleunderline_6 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2673 := __args[0]
_ = V2673
reg298061 := PrimHead(V2673)
reg298062 := PrimIsPair(reg298061)
if reg298062 == True {
reg298063 := PrimHead(V2673)
reg298064 := PrimHead(reg298063)
Parse__X := reg298064
_ = Parse__X
reg298065 := __e.Call(__defun__shen_4singleunderline_2, Parse__X)
if reg298065 == True {
reg298066 := PrimHead(V2673)
reg298067 := PrimTail(reg298066)
reg298068 := __e.Call(__defun__shen_4hdtl, V2673)
reg298069 := __e.Call(__defun__shen_4pair, reg298067, reg298068)
reg298070 := PrimHead(reg298069)
__ctx.TailApply(__defun__shen_4pair, reg298070, Parse__X)
return
} else {
__ctx.TailApply(__defun__fail)
return
}
} else {
__ctx.TailApply(__defun__fail)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.<singleunderline>", value: __defun__shen_4_5singleunderline_6})

__defun__shen_4singleunderline_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2675 := __args[0]
_ = V2675
reg298074 := PrimIsSymbol(V2675)
if reg298074 == True {
reg298075 := PrimStr(V2675)
reg298076 := __e.Call(__defun__shen_4sh_2, reg298075)
if reg298076 == True {
reg298077 := True;
__ctx.Return(reg298077)
return
} else {
reg298078 := False;
__ctx.Return(reg298078)
return
}
} else {
reg298079 := False;
__ctx.Return(reg298079)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.singleunderline?", value: __defun__shen_4singleunderline_2})

__defun__shen_4sh_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2677 := __args[0]
_ = V2677
reg298080 := MakeString("_")
reg298081 := PrimEqual(reg298080, V2677)
if reg298081 == True {
reg298082 := True;
__ctx.Return(reg298082)
return
} else {
reg298083 := MakeNumber(0)
reg298084 := PrimPos(V2677, reg298083)
reg298085 := MakeString("_")
reg298086 := PrimEqual(reg298084, reg298085)
if reg298086 == True {
reg298087 := PrimTailString(V2677)
reg298088 := __e.Call(__defun__shen_4sh_2, reg298087)
if reg298088 == True {
reg298089 := True;
__ctx.Return(reg298089)
return
} else {
reg298090 := False;
__ctx.Return(reg298090)
return
}
} else {
reg298091 := False;
__ctx.Return(reg298091)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.sh?", value: __defun__shen_4sh_2})

__defun__shen_4doubleunderline_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2679 := __args[0]
_ = V2679
reg298092 := PrimIsSymbol(V2679)
if reg298092 == True {
reg298093 := PrimStr(V2679)
reg298094 := __e.Call(__defun__shen_4dh_2, reg298093)
if reg298094 == True {
reg298095 := True;
__ctx.Return(reg298095)
return
} else {
reg298096 := False;
__ctx.Return(reg298096)
return
}
} else {
reg298097 := False;
__ctx.Return(reg298097)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.doubleunderline?", value: __defun__shen_4doubleunderline_2})

__defun__shen_4dh_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2681 := __args[0]
_ = V2681
reg298098 := MakeString("=")
reg298099 := PrimEqual(reg298098, V2681)
if reg298099 == True {
reg298100 := True;
__ctx.Return(reg298100)
return
} else {
reg298101 := MakeNumber(0)
reg298102 := PrimPos(V2681, reg298101)
reg298103 := MakeString("=")
reg298104 := PrimEqual(reg298102, reg298103)
if reg298104 == True {
reg298105 := PrimTailString(V2681)
reg298106 := __e.Call(__defun__shen_4dh_2, reg298105)
if reg298106 == True {
reg298107 := True;
__ctx.Return(reg298107)
return
} else {
reg298108 := False;
__ctx.Return(reg298108)
return
}
} else {
reg298109 := False;
__ctx.Return(reg298109)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.dh?", value: __defun__shen_4dh_2})

__defun__shen_4process_1datatype = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2684 := __args[0]
_ = V2684
V2685 := __args[1]
_ = V2685
reg298110 := __e.Call(__defun__shen_4rules_1_6horn_1clauses, V2684, V2685)
reg298111 := __e.Call(__defun__shen_4s_1prolog, reg298110)
__ctx.TailApply(__defun__shen_4remember_1datatype, reg298111)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.process-datatype", value: __defun__shen_4process_1datatype})

__defun__shen_4remember_1datatype = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2691 := __args[0]
_ = V2691
reg298113 := PrimIsPair(V2691)
if reg298113 == True {
reg298114 := MakeSymbol("shen.*datatypes*")
reg298115 := PrimHead(V2691)
reg298116 := MakeSymbol("shen.*datatypes*")
reg298117 := PrimValue(reg298116)
reg298118 := __e.Call(__defun__adjoin, reg298115, reg298117)
reg298119 := PrimSet(reg298114, reg298118)
_ = reg298119
reg298120 := MakeSymbol("shen.*alldatatypes*")
reg298121 := PrimHead(V2691)
reg298122 := MakeSymbol("shen.*alldatatypes*")
reg298123 := PrimValue(reg298122)
reg298124 := __e.Call(__defun__adjoin, reg298121, reg298123)
reg298125 := PrimSet(reg298120, reg298124)
_ = reg298125
reg298126 := PrimHead(V2691)
__ctx.Return(reg298126)
return
} else {
reg298127 := MakeSymbol("shen.remember-datatype")
__ctx.TailApply(__defun__shen_4f__error, reg298127)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.remember-datatype", value: __defun__shen_4remember_1datatype})

__defun__shen_4rules_1_6horn_1clauses = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2696 := __args[0]
_ = V2696
V2697 := __args[1]
_ = V2697
reg298129 := Nil;
reg298130 := PrimEqual(reg298129, V2697)
if reg298130 == True {
reg298131 := Nil;
__ctx.Return(reg298131)
return
} else {
reg298132 := PrimIsPair(V2697)
var reg298148 Obj
if reg298132 == True {
reg298133 := PrimHead(V2697)
reg298134 := __e.Call(__defun__tuple_2, reg298133)
var reg298143 Obj
if reg298134 == True {
reg298135 := MakeSymbol("shen.single")
reg298136 := PrimHead(V2697)
reg298137 := __e.Call(__defun__fst, reg298136)
reg298138 := PrimEqual(reg298135, reg298137)
var reg298141 Obj
if reg298138 == True {
reg298139 := True;
reg298141 = reg298139
} else {
reg298140 := False;
reg298141 = reg298140
}
reg298143 = reg298141
} else {
reg298142 := False;
reg298143 = reg298142
}
var reg298146 Obj
if reg298143 == True {
reg298144 := True;
reg298146 = reg298144
} else {
reg298145 := False;
reg298146 = reg298145
}
reg298148 = reg298146
} else {
reg298147 := False;
reg298148 = reg298147
}
if reg298148 == True {
reg298149 := PrimHead(V2697)
reg298150 := __e.Call(__defun__snd, reg298149)
reg298151 := __e.Call(__defun__shen_4rule_1_6horn_1clause, V2696, reg298150)
reg298152 := PrimTail(V2697)
reg298153 := __e.Call(__defun__shen_4rules_1_6horn_1clauses, V2696, reg298152)
reg298154 := PrimCons(reg298151, reg298153)
__ctx.Return(reg298154)
return
} else {
reg298155 := PrimIsPair(V2697)
var reg298171 Obj
if reg298155 == True {
reg298156 := PrimHead(V2697)
reg298157 := __e.Call(__defun__tuple_2, reg298156)
var reg298166 Obj
if reg298157 == True {
reg298158 := MakeSymbol("shen.double")
reg298159 := PrimHead(V2697)
reg298160 := __e.Call(__defun__fst, reg298159)
reg298161 := PrimEqual(reg298158, reg298160)
var reg298164 Obj
if reg298161 == True {
reg298162 := True;
reg298164 = reg298162
} else {
reg298163 := False;
reg298164 = reg298163
}
reg298166 = reg298164
} else {
reg298165 := False;
reg298166 = reg298165
}
var reg298169 Obj
if reg298166 == True {
reg298167 := True;
reg298169 = reg298167
} else {
reg298168 := False;
reg298169 = reg298168
}
reg298171 = reg298169
} else {
reg298170 := False;
reg298171 = reg298170
}
if reg298171 == True {
reg298172 := PrimHead(V2697)
reg298173 := __e.Call(__defun__snd, reg298172)
reg298174 := __e.Call(__defun__shen_4double_1_6singles, reg298173)
reg298175 := PrimTail(V2697)
reg298176 := __e.Call(__defun__append, reg298174, reg298175)
__ctx.TailApply(__defun__shen_4rules_1_6horn_1clauses, V2696, reg298176)
return
} else {
reg298178 := MakeSymbol("shen.rules->horn-clauses")
__ctx.TailApply(__defun__shen_4f__error, reg298178)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.rules->horn-clauses", value: __defun__shen_4rules_1_6horn_1clauses})

__defun__shen_4double_1_6singles = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2699 := __args[0]
_ = V2699
reg298180 := __e.Call(__defun__shen_4right_1rule, V2699)
reg298181 := __e.Call(__defun__shen_4left_1rule, V2699)
reg298182 := Nil;
reg298183 := PrimCons(reg298181, reg298182)
reg298184 := PrimCons(reg298180, reg298183)
__ctx.Return(reg298184)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.double->singles", value: __defun__shen_4double_1_6singles})

__defun__shen_4right_1rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2701 := __args[0]
_ = V2701
reg298185 := MakeSymbol("shen.single")
__ctx.TailApply(__defun___8p, reg298185, V2701)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.right-rule", value: __defun__shen_4right_1rule})

__defun__shen_4left_1rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2703 := __args[0]
_ = V2703
reg298187 := PrimIsPair(V2703)
var reg298232 Obj
if reg298187 == True {
reg298188 := PrimTail(V2703)
reg298189 := PrimIsPair(reg298188)
var reg298227 Obj
if reg298189 == True {
reg298190 := PrimTail(V2703)
reg298191 := PrimTail(reg298190)
reg298192 := PrimIsPair(reg298191)
var reg298222 Obj
if reg298192 == True {
reg298193 := PrimTail(V2703)
reg298194 := PrimTail(reg298193)
reg298195 := PrimHead(reg298194)
reg298196 := __e.Call(__defun__tuple_2, reg298195)
var reg298217 Obj
if reg298196 == True {
reg298197 := Nil;
reg298198 := PrimTail(V2703)
reg298199 := PrimTail(reg298198)
reg298200 := PrimHead(reg298199)
reg298201 := __e.Call(__defun__fst, reg298200)
reg298202 := PrimEqual(reg298197, reg298201)
var reg298212 Obj
if reg298202 == True {
reg298203 := Nil;
reg298204 := PrimTail(V2703)
reg298205 := PrimTail(reg298204)
reg298206 := PrimTail(reg298205)
reg298207 := PrimEqual(reg298203, reg298206)
var reg298210 Obj
if reg298207 == True {
reg298208 := True;
reg298210 = reg298208
} else {
reg298209 := False;
reg298210 = reg298209
}
reg298212 = reg298210
} else {
reg298211 := False;
reg298212 = reg298211
}
var reg298215 Obj
if reg298212 == True {
reg298213 := True;
reg298215 = reg298213
} else {
reg298214 := False;
reg298215 = reg298214
}
reg298217 = reg298215
} else {
reg298216 := False;
reg298217 = reg298216
}
var reg298220 Obj
if reg298217 == True {
reg298218 := True;
reg298220 = reg298218
} else {
reg298219 := False;
reg298220 = reg298219
}
reg298222 = reg298220
} else {
reg298221 := False;
reg298222 = reg298221
}
var reg298225 Obj
if reg298222 == True {
reg298223 := True;
reg298225 = reg298223
} else {
reg298224 := False;
reg298225 = reg298224
}
reg298227 = reg298225
} else {
reg298226 := False;
reg298227 = reg298226
}
var reg298230 Obj
if reg298227 == True {
reg298228 := True;
reg298230 = reg298228
} else {
reg298229 := False;
reg298230 = reg298229
}
reg298232 = reg298230
} else {
reg298231 := False;
reg298232 = reg298231
}
if reg298232 == True {
reg298233 := MakeSymbol("Qv")
reg298234 := __e.Call(__defun__gensym, reg298233)
Q := reg298234
_ = Q
reg298235 := PrimTail(V2703)
reg298236 := PrimTail(reg298235)
reg298237 := PrimHead(reg298236)
reg298238 := __e.Call(__defun__snd, reg298237)
reg298239 := Nil;
reg298240 := PrimCons(reg298238, reg298239)
reg298241 := __e.Call(__defun___8p, reg298240, Q)
NewConclusion := reg298241
_ = NewConclusion
reg298242 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4right_1_6left, X)
return
}, 1)
reg298244 := PrimTail(V2703)
reg298245 := PrimHead(reg298244)
reg298246 := __e.Call(__defun__map, reg298242, reg298245)
reg298247 := __e.Call(__defun___8p, reg298246, Q)
reg298248 := Nil;
reg298249 := PrimCons(reg298247, reg298248)
NewPremises := reg298249
_ = NewPremises
reg298250 := MakeSymbol("shen.single")
reg298251 := PrimHead(V2703)
reg298252 := Nil;
reg298253 := PrimCons(NewConclusion, reg298252)
reg298254 := PrimCons(NewPremises, reg298253)
reg298255 := PrimCons(reg298251, reg298254)
__ctx.TailApply(__defun___8p, reg298250, reg298255)
return
} else {
reg298257 := MakeSymbol("shen.left-rule")
__ctx.TailApply(__defun__shen_4f__error, reg298257)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.left-rule", value: __defun__shen_4left_1rule})

__defun__shen_4right_1_6left = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2709 := __args[0]
_ = V2709
reg298259 := __e.Call(__defun__tuple_2, V2709)
var reg298267 Obj
if reg298259 == True {
reg298260 := Nil;
reg298261 := __e.Call(__defun__fst, V2709)
reg298262 := PrimEqual(reg298260, reg298261)
var reg298265 Obj
if reg298262 == True {
reg298263 := True;
reg298265 = reg298263
} else {
reg298264 := False;
reg298265 = reg298264
}
reg298267 = reg298265
} else {
reg298266 := False;
reg298267 = reg298266
}
if reg298267 == True {
__ctx.TailApply(__defun__snd, V2709)
return
} else {
reg298269 := MakeString("syntax error with ==========\n")
reg298270 := PrimSimpleError(reg298269)
__ctx.Return(reg298270)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.right->left", value: __defun__shen_4right_1_6left})

__defun__shen_4rule_1_6horn_1clause = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2712 := __args[0]
_ = V2712
V2713 := __args[1]
_ = V2713
reg298271 := PrimIsPair(V2713)
var reg298305 Obj
if reg298271 == True {
reg298272 := PrimTail(V2713)
reg298273 := PrimIsPair(reg298272)
var reg298300 Obj
if reg298273 == True {
reg298274 := PrimTail(V2713)
reg298275 := PrimTail(reg298274)
reg298276 := PrimIsPair(reg298275)
var reg298295 Obj
if reg298276 == True {
reg298277 := PrimTail(V2713)
reg298278 := PrimTail(reg298277)
reg298279 := PrimHead(reg298278)
reg298280 := __e.Call(__defun__tuple_2, reg298279)
var reg298290 Obj
if reg298280 == True {
reg298281 := Nil;
reg298282 := PrimTail(V2713)
reg298283 := PrimTail(reg298282)
reg298284 := PrimTail(reg298283)
reg298285 := PrimEqual(reg298281, reg298284)
var reg298288 Obj
if reg298285 == True {
reg298286 := True;
reg298288 = reg298286
} else {
reg298287 := False;
reg298288 = reg298287
}
reg298290 = reg298288
} else {
reg298289 := False;
reg298290 = reg298289
}
var reg298293 Obj
if reg298290 == True {
reg298291 := True;
reg298293 = reg298291
} else {
reg298292 := False;
reg298293 = reg298292
}
reg298295 = reg298293
} else {
reg298294 := False;
reg298295 = reg298294
}
var reg298298 Obj
if reg298295 == True {
reg298296 := True;
reg298298 = reg298296
} else {
reg298297 := False;
reg298298 = reg298297
}
reg298300 = reg298298
} else {
reg298299 := False;
reg298300 = reg298299
}
var reg298303 Obj
if reg298300 == True {
reg298301 := True;
reg298303 = reg298301
} else {
reg298302 := False;
reg298303 = reg298302
}
reg298305 = reg298303
} else {
reg298304 := False;
reg298305 = reg298304
}
if reg298305 == True {
reg298306 := PrimTail(V2713)
reg298307 := PrimTail(reg298306)
reg298308 := PrimHead(reg298307)
reg298309 := __e.Call(__defun__snd, reg298308)
reg298310 := __e.Call(__defun__shen_4rule_1_6horn_1clause_1head, V2712, reg298309)
reg298311 := MakeSymbol(":-")
reg298312 := PrimHead(V2713)
reg298313 := PrimTail(V2713)
reg298314 := PrimHead(reg298313)
reg298315 := PrimTail(V2713)
reg298316 := PrimTail(reg298315)
reg298317 := PrimHead(reg298316)
reg298318 := __e.Call(__defun__fst, reg298317)
reg298319 := __e.Call(__defun__shen_4rule_1_6horn_1clause_1body, reg298312, reg298314, reg298318)
reg298320 := Nil;
reg298321 := PrimCons(reg298319, reg298320)
reg298322 := PrimCons(reg298311, reg298321)
reg298323 := PrimCons(reg298310, reg298322)
__ctx.Return(reg298323)
return
} else {
reg298324 := MakeSymbol("shen.rule->horn-clause")
__ctx.TailApply(__defun__shen_4f__error, reg298324)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.rule->horn-clause", value: __defun__shen_4rule_1_6horn_1clause})

__defun__shen_4rule_1_6horn_1clause_1head = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2716 := __args[0]
_ = V2716
V2717 := __args[1]
_ = V2717
reg298326 := __e.Call(__defun__shen_4mode_1ify, V2717)
reg298327 := MakeSymbol("Context_1957")
reg298328 := Nil;
reg298329 := PrimCons(reg298327, reg298328)
reg298330 := PrimCons(reg298326, reg298329)
reg298331 := PrimCons(V2716, reg298330)
__ctx.Return(reg298331)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.rule->horn-clause-head", value: __defun__shen_4rule_1_6horn_1clause_1head})

__defun__shen_4mode_1ify = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2719 := __args[0]
_ = V2719
reg298332 := PrimIsPair(V2719)
var reg298366 Obj
if reg298332 == True {
reg298333 := PrimTail(V2719)
reg298334 := PrimIsPair(reg298333)
var reg298361 Obj
if reg298334 == True {
reg298335 := MakeSymbol(":")
reg298336 := PrimTail(V2719)
reg298337 := PrimHead(reg298336)
reg298338 := PrimEqual(reg298335, reg298337)
var reg298356 Obj
if reg298338 == True {
reg298339 := PrimTail(V2719)
reg298340 := PrimTail(reg298339)
reg298341 := PrimIsPair(reg298340)
var reg298351 Obj
if reg298341 == True {
reg298342 := Nil;
reg298343 := PrimTail(V2719)
reg298344 := PrimTail(reg298343)
reg298345 := PrimTail(reg298344)
reg298346 := PrimEqual(reg298342, reg298345)
var reg298349 Obj
if reg298346 == True {
reg298347 := True;
reg298349 = reg298347
} else {
reg298348 := False;
reg298349 = reg298348
}
reg298351 = reg298349
} else {
reg298350 := False;
reg298351 = reg298350
}
var reg298354 Obj
if reg298351 == True {
reg298352 := True;
reg298354 = reg298352
} else {
reg298353 := False;
reg298354 = reg298353
}
reg298356 = reg298354
} else {
reg298355 := False;
reg298356 = reg298355
}
var reg298359 Obj
if reg298356 == True {
reg298357 := True;
reg298359 = reg298357
} else {
reg298358 := False;
reg298359 = reg298358
}
reg298361 = reg298359
} else {
reg298360 := False;
reg298361 = reg298360
}
var reg298364 Obj
if reg298361 == True {
reg298362 := True;
reg298364 = reg298362
} else {
reg298363 := False;
reg298364 = reg298363
}
reg298366 = reg298364
} else {
reg298365 := False;
reg298366 = reg298365
}
if reg298366 == True {
reg298367 := MakeSymbol("mode")
reg298368 := PrimHead(V2719)
reg298369 := MakeSymbol(":")
reg298370 := MakeSymbol("mode")
reg298371 := PrimTail(V2719)
reg298372 := PrimTail(reg298371)
reg298373 := PrimHead(reg298372)
reg298374 := MakeSymbol("+")
reg298375 := Nil;
reg298376 := PrimCons(reg298374, reg298375)
reg298377 := PrimCons(reg298373, reg298376)
reg298378 := PrimCons(reg298370, reg298377)
reg298379 := Nil;
reg298380 := PrimCons(reg298378, reg298379)
reg298381 := PrimCons(reg298369, reg298380)
reg298382 := PrimCons(reg298368, reg298381)
reg298383 := MakeSymbol("-")
reg298384 := Nil;
reg298385 := PrimCons(reg298383, reg298384)
reg298386 := PrimCons(reg298382, reg298385)
reg298387 := PrimCons(reg298367, reg298386)
__ctx.Return(reg298387)
return
} else {
__ctx.Return(V2719)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.mode-ify", value: __defun__shen_4mode_1ify})

__defun__shen_4rule_1_6horn_1clause_1body = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2723 := __args[0]
_ = V2723
V2724 := __args[1]
_ = V2724
V2725 := __args[2]
_ = V2725
reg298388 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4extract__vars, X)
return
}, 1)
reg298390 := __e.Call(__defun__map, reg298388, V2725)
Variables := reg298390
_ = Variables
reg298391 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg298392 := MakeSymbol("shen.cl")
__ctx.TailApply(__defun__gensym, reg298392)
return
}, 1)
reg298394 := __e.Call(__defun__map, reg298391, V2725)
Predicates := reg298394
_ = Predicates
reg298395 := MakeSymbol("Context_1957")
reg298396 := MakeSymbol("Context1_1957")
reg298397 := __e.Call(__defun__shen_4construct_1search_1literals, Predicates, Variables, reg298395, reg298396)
SearchLiterals := reg298397
_ = SearchLiterals
reg298398 := __e.Call(__defun__shen_4construct_1search_1clauses, Predicates, V2725, Variables)
SearchClauses := reg298398
_ = SearchClauses
reg298399 := __e.Call(__defun__shen_4construct_1side_1literals, V2723)
SideLiterals := reg298399
_ = SideLiterals
reg298400 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
reg298401 := __e.Call(__defun__empty_2, V2725)
__ctx.TailApply(__defun__shen_4construct_1premiss_1literal, X, reg298401)
return
}, 1)
reg298403 := __e.Call(__defun__map, reg298400, V2724)
PremissLiterals := reg298403
_ = PremissLiterals
reg298404 := __e.Call(__defun__append, SideLiterals, PremissLiterals)
__ctx.TailApply(__defun__append, SearchLiterals, reg298404)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.rule->horn-clause-body", value: __defun__shen_4rule_1_6horn_1clause_1body})

__defun__shen_4construct_1search_1literals = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2734 := __args[0]
_ = V2734
V2735 := __args[1]
_ = V2735
V2736 := __args[2]
_ = V2736
V2737 := __args[3]
_ = V2737
reg298406 := Nil;
reg298407 := PrimEqual(reg298406, V2734)
var reg298414 Obj
if reg298407 == True {
reg298408 := Nil;
reg298409 := PrimEqual(reg298408, V2735)
var reg298412 Obj
if reg298409 == True {
reg298410 := True;
reg298412 = reg298410
} else {
reg298411 := False;
reg298412 = reg298411
}
reg298414 = reg298412
} else {
reg298413 := False;
reg298414 = reg298413
}
if reg298414 == True {
reg298415 := Nil;
__ctx.Return(reg298415)
return
} else {
__ctx.TailApply(__defun__shen_4csl_1help, V2734, V2735, V2736, V2737)
return
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.construct-search-literals", value: __defun__shen_4construct_1search_1literals})

__defun__shen_4csl_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2744 := __args[0]
_ = V2744
V2745 := __args[1]
_ = V2745
V2746 := __args[2]
_ = V2746
V2747 := __args[3]
_ = V2747
reg298417 := Nil;
reg298418 := PrimEqual(reg298417, V2744)
var reg298425 Obj
if reg298418 == True {
reg298419 := Nil;
reg298420 := PrimEqual(reg298419, V2745)
var reg298423 Obj
if reg298420 == True {
reg298421 := True;
reg298423 = reg298421
} else {
reg298422 := False;
reg298423 = reg298422
}
reg298425 = reg298423
} else {
reg298424 := False;
reg298425 = reg298424
}
if reg298425 == True {
reg298426 := MakeSymbol("bind")
reg298427 := MakeSymbol("ContextOut_1957")
reg298428 := Nil;
reg298429 := PrimCons(V2746, reg298428)
reg298430 := PrimCons(reg298427, reg298429)
reg298431 := PrimCons(reg298426, reg298430)
reg298432 := Nil;
reg298433 := PrimCons(reg298431, reg298432)
__ctx.Return(reg298433)
return
} else {
reg298434 := PrimIsPair(V2744)
var reg298440 Obj
if reg298434 == True {
reg298435 := PrimIsPair(V2745)
var reg298438 Obj
if reg298435 == True {
reg298436 := True;
reg298438 = reg298436
} else {
reg298437 := False;
reg298438 = reg298437
}
reg298440 = reg298438
} else {
reg298439 := False;
reg298440 = reg298439
}
if reg298440 == True {
reg298441 := PrimHead(V2744)
reg298442 := PrimHead(V2745)
reg298443 := PrimCons(V2747, reg298442)
reg298444 := PrimCons(V2746, reg298443)
reg298445 := PrimCons(reg298441, reg298444)
reg298446 := PrimTail(V2744)
reg298447 := PrimTail(V2745)
reg298448 := MakeSymbol("Context")
reg298449 := __e.Call(__defun__gensym, reg298448)
reg298450 := __e.Call(__defun__shen_4csl_1help, reg298446, reg298447, V2747, reg298449)
reg298451 := PrimCons(reg298445, reg298450)
__ctx.Return(reg298451)
return
} else {
reg298452 := MakeSymbol("shen.csl-help")
__ctx.TailApply(__defun__shen_4f__error, reg298452)
return
}
}
}, 4)
__initDefs = append(__initDefs, defType{name: "shen.csl-help", value: __defun__shen_4csl_1help})

__defun__shen_4construct_1search_1clauses = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2751 := __args[0]
_ = V2751
V2752 := __args[1]
_ = V2752
V2753 := __args[2]
_ = V2753
reg298454 := Nil;
reg298455 := PrimEqual(reg298454, V2751)
var reg298469 Obj
if reg298455 == True {
reg298456 := Nil;
reg298457 := PrimEqual(reg298456, V2752)
var reg298464 Obj
if reg298457 == True {
reg298458 := Nil;
reg298459 := PrimEqual(reg298458, V2753)
var reg298462 Obj
if reg298459 == True {
reg298460 := True;
reg298462 = reg298460
} else {
reg298461 := False;
reg298462 = reg298461
}
reg298464 = reg298462
} else {
reg298463 := False;
reg298464 = reg298463
}
var reg298467 Obj
if reg298464 == True {
reg298465 := True;
reg298467 = reg298465
} else {
reg298466 := False;
reg298467 = reg298466
}
reg298469 = reg298467
} else {
reg298468 := False;
reg298469 = reg298468
}
if reg298469 == True {
reg298470 := MakeSymbol("shen.skip")
__ctx.Return(reg298470)
return
} else {
reg298471 := PrimIsPair(V2751)
var reg298483 Obj
if reg298471 == True {
reg298472 := PrimIsPair(V2752)
var reg298478 Obj
if reg298472 == True {
reg298473 := PrimIsPair(V2753)
var reg298476 Obj
if reg298473 == True {
reg298474 := True;
reg298476 = reg298474
} else {
reg298475 := False;
reg298476 = reg298475
}
reg298478 = reg298476
} else {
reg298477 := False;
reg298478 = reg298477
}
var reg298481 Obj
if reg298478 == True {
reg298479 := True;
reg298481 = reg298479
} else {
reg298480 := False;
reg298481 = reg298480
}
reg298483 = reg298481
} else {
reg298482 := False;
reg298483 = reg298482
}
if reg298483 == True {
reg298484 := PrimHead(V2751)
reg298485 := PrimHead(V2752)
reg298486 := PrimHead(V2753)
reg298487 := __e.Call(__defun__shen_4construct_1search_1clause, reg298484, reg298485, reg298486)
_ = reg298487
reg298488 := PrimTail(V2751)
reg298489 := PrimTail(V2752)
reg298490 := PrimTail(V2753)
__ctx.TailApply(__defun__shen_4construct_1search_1clauses, reg298488, reg298489, reg298490)
return
} else {
reg298492 := MakeSymbol("shen.construct-search-clauses")
__ctx.TailApply(__defun__shen_4f__error, reg298492)
return
}
}
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.construct-search-clauses", value: __defun__shen_4construct_1search_1clauses})

__defun__shen_4construct_1search_1clause = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2757 := __args[0]
_ = V2757
V2758 := __args[1]
_ = V2758
V2759 := __args[2]
_ = V2759
reg298494 := __e.Call(__defun__shen_4construct_1base_1search_1clause, V2757, V2758, V2759)
reg298495 := __e.Call(__defun__shen_4construct_1recursive_1search_1clause, V2757, V2758, V2759)
reg298496 := Nil;
reg298497 := PrimCons(reg298495, reg298496)
reg298498 := PrimCons(reg298494, reg298497)
__ctx.TailApply(__defun__shen_4s_1prolog, reg298498)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.construct-search-clause", value: __defun__shen_4construct_1search_1clause})

__defun__shen_4construct_1base_1search_1clause = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2763 := __args[0]
_ = V2763
V2764 := __args[1]
_ = V2764
V2765 := __args[2]
_ = V2765
reg298500 := __e.Call(__defun__shen_4mode_1ify, V2764)
reg298501 := MakeSymbol("In_1957")
reg298502 := PrimCons(reg298500, reg298501)
reg298503 := MakeSymbol("In_1957")
reg298504 := PrimCons(reg298503, V2765)
reg298505 := PrimCons(reg298502, reg298504)
reg298506 := PrimCons(V2763, reg298505)
reg298507 := MakeSymbol(":-")
reg298508 := Nil;
reg298509 := Nil;
reg298510 := PrimCons(reg298508, reg298509)
reg298511 := PrimCons(reg298507, reg298510)
reg298512 := PrimCons(reg298506, reg298511)
__ctx.Return(reg298512)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.construct-base-search-clause", value: __defun__shen_4construct_1base_1search_1clause})

__defun__shen_4construct_1recursive_1search_1clause = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2769 := __args[0]
_ = V2769
V2770 := __args[1]
_ = V2770
V2771 := __args[2]
_ = V2771
reg298513 := MakeSymbol("Assumption_1957")
reg298514 := MakeSymbol("Assumptions_1957")
reg298515 := PrimCons(reg298513, reg298514)
reg298516 := MakeSymbol("Assumption_1957")
reg298517 := MakeSymbol("Out_1957")
reg298518 := PrimCons(reg298516, reg298517)
reg298519 := PrimCons(reg298518, V2771)
reg298520 := PrimCons(reg298515, reg298519)
reg298521 := PrimCons(V2769, reg298520)
reg298522 := MakeSymbol(":-")
reg298523 := MakeSymbol("Assumptions_1957")
reg298524 := MakeSymbol("Out_1957")
reg298525 := PrimCons(reg298524, V2771)
reg298526 := PrimCons(reg298523, reg298525)
reg298527 := PrimCons(V2769, reg298526)
reg298528 := Nil;
reg298529 := PrimCons(reg298527, reg298528)
reg298530 := Nil;
reg298531 := PrimCons(reg298529, reg298530)
reg298532 := PrimCons(reg298522, reg298531)
reg298533 := PrimCons(reg298521, reg298532)
__ctx.Return(reg298533)
return
}, 3)
__initDefs = append(__initDefs, defType{name: "shen.construct-recursive-search-clause", value: __defun__shen_4construct_1recursive_1search_1clause})

__defun__shen_4construct_1side_1literals = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2777 := __args[0]
_ = V2777
reg298534 := Nil;
reg298535 := PrimEqual(reg298534, V2777)
if reg298535 == True {
reg298536 := Nil;
__ctx.Return(reg298536)
return
} else {
reg298537 := PrimIsPair(V2777)
var reg298571 Obj
if reg298537 == True {
reg298538 := PrimHead(V2777)
reg298539 := PrimIsPair(reg298538)
var reg298566 Obj
if reg298539 == True {
reg298540 := MakeSymbol("if")
reg298541 := PrimHead(V2777)
reg298542 := PrimHead(reg298541)
reg298543 := PrimEqual(reg298540, reg298542)
var reg298561 Obj
if reg298543 == True {
reg298544 := PrimHead(V2777)
reg298545 := PrimTail(reg298544)
reg298546 := PrimIsPair(reg298545)
var reg298556 Obj
if reg298546 == True {
reg298547 := Nil;
reg298548 := PrimHead(V2777)
reg298549 := PrimTail(reg298548)
reg298550 := PrimTail(reg298549)
reg298551 := PrimEqual(reg298547, reg298550)
var reg298554 Obj
if reg298551 == True {
reg298552 := True;
reg298554 = reg298552
} else {
reg298553 := False;
reg298554 = reg298553
}
reg298556 = reg298554
} else {
reg298555 := False;
reg298556 = reg298555
}
var reg298559 Obj
if reg298556 == True {
reg298557 := True;
reg298559 = reg298557
} else {
reg298558 := False;
reg298559 = reg298558
}
reg298561 = reg298559
} else {
reg298560 := False;
reg298561 = reg298560
}
var reg298564 Obj
if reg298561 == True {
reg298562 := True;
reg298564 = reg298562
} else {
reg298563 := False;
reg298564 = reg298563
}
reg298566 = reg298564
} else {
reg298565 := False;
reg298566 = reg298565
}
var reg298569 Obj
if reg298566 == True {
reg298567 := True;
reg298569 = reg298567
} else {
reg298568 := False;
reg298569 = reg298568
}
reg298571 = reg298569
} else {
reg298570 := False;
reg298571 = reg298570
}
if reg298571 == True {
reg298572 := MakeSymbol("when")
reg298573 := PrimHead(V2777)
reg298574 := PrimTail(reg298573)
reg298575 := PrimCons(reg298572, reg298574)
reg298576 := PrimTail(V2777)
reg298577 := __e.Call(__defun__shen_4construct_1side_1literals, reg298576)
reg298578 := PrimCons(reg298575, reg298577)
__ctx.Return(reg298578)
return
} else {
reg298579 := PrimIsPair(V2777)
var reg298623 Obj
if reg298579 == True {
reg298580 := PrimHead(V2777)
reg298581 := PrimIsPair(reg298580)
var reg298618 Obj
if reg298581 == True {
reg298582 := MakeSymbol("let")
reg298583 := PrimHead(V2777)
reg298584 := PrimHead(reg298583)
reg298585 := PrimEqual(reg298582, reg298584)
var reg298613 Obj
if reg298585 == True {
reg298586 := PrimHead(V2777)
reg298587 := PrimTail(reg298586)
reg298588 := PrimIsPair(reg298587)
var reg298608 Obj
if reg298588 == True {
reg298589 := PrimHead(V2777)
reg298590 := PrimTail(reg298589)
reg298591 := PrimTail(reg298590)
reg298592 := PrimIsPair(reg298591)
var reg298603 Obj
if reg298592 == True {
reg298593 := Nil;
reg298594 := PrimHead(V2777)
reg298595 := PrimTail(reg298594)
reg298596 := PrimTail(reg298595)
reg298597 := PrimTail(reg298596)
reg298598 := PrimEqual(reg298593, reg298597)
var reg298601 Obj
if reg298598 == True {
reg298599 := True;
reg298601 = reg298599
} else {
reg298600 := False;
reg298601 = reg298600
}
reg298603 = reg298601
} else {
reg298602 := False;
reg298603 = reg298602
}
var reg298606 Obj
if reg298603 == True {
reg298604 := True;
reg298606 = reg298604
} else {
reg298605 := False;
reg298606 = reg298605
}
reg298608 = reg298606
} else {
reg298607 := False;
reg298608 = reg298607
}
var reg298611 Obj
if reg298608 == True {
reg298609 := True;
reg298611 = reg298609
} else {
reg298610 := False;
reg298611 = reg298610
}
reg298613 = reg298611
} else {
reg298612 := False;
reg298613 = reg298612
}
var reg298616 Obj
if reg298613 == True {
reg298614 := True;
reg298616 = reg298614
} else {
reg298615 := False;
reg298616 = reg298615
}
reg298618 = reg298616
} else {
reg298617 := False;
reg298618 = reg298617
}
var reg298621 Obj
if reg298618 == True {
reg298619 := True;
reg298621 = reg298619
} else {
reg298620 := False;
reg298621 = reg298620
}
reg298623 = reg298621
} else {
reg298622 := False;
reg298623 = reg298622
}
if reg298623 == True {
reg298624 := MakeSymbol("is")
reg298625 := PrimHead(V2777)
reg298626 := PrimTail(reg298625)
reg298627 := PrimCons(reg298624, reg298626)
reg298628 := PrimTail(V2777)
reg298629 := __e.Call(__defun__shen_4construct_1side_1literals, reg298628)
reg298630 := PrimCons(reg298627, reg298629)
__ctx.Return(reg298630)
return
} else {
reg298631 := PrimIsPair(V2777)
if reg298631 == True {
reg298632 := PrimTail(V2777)
__ctx.TailApply(__defun__shen_4construct_1side_1literals, reg298632)
return
} else {
reg298634 := MakeSymbol("shen.construct-side-literals")
__ctx.TailApply(__defun__shen_4f__error, reg298634)
return
}
}
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.construct-side-literals", value: __defun__shen_4construct_1side_1literals})

__defun__shen_4construct_1premiss_1literal = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2784 := __args[0]
_ = V2784
V2785 := __args[1]
_ = V2785
reg298636 := __e.Call(__defun__tuple_2, V2784)
if reg298636 == True {
reg298637 := MakeSymbol("shen.t*")
reg298638 := __e.Call(__defun__snd, V2784)
reg298639 := __e.Call(__defun__shen_4recursive__cons__form, reg298638)
reg298640 := __e.Call(__defun__fst, V2784)
reg298641 := __e.Call(__defun__shen_4construct_1context, V2785, reg298640)
reg298642 := Nil;
reg298643 := PrimCons(reg298641, reg298642)
reg298644 := PrimCons(reg298639, reg298643)
reg298645 := PrimCons(reg298637, reg298644)
__ctx.Return(reg298645)
return
} else {
reg298646 := MakeSymbol("!")
reg298647 := PrimEqual(reg298646, V2784)
if reg298647 == True {
reg298648 := MakeSymbol("cut")
reg298649 := MakeSymbol("Throwcontrol")
reg298650 := Nil;
reg298651 := PrimCons(reg298649, reg298650)
reg298652 := PrimCons(reg298648, reg298651)
__ctx.Return(reg298652)
return
} else {
reg298653 := MakeSymbol("shen.construct-premiss-literal")
__ctx.TailApply(__defun__shen_4f__error, reg298653)
return
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.construct-premiss-literal", value: __defun__shen_4construct_1premiss_1literal})

__defun__shen_4construct_1context = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2788 := __args[0]
_ = V2788
V2789 := __args[1]
_ = V2789
reg298655 := True;
reg298656 := PrimEqual(reg298655, V2788)
var reg298663 Obj
if reg298656 == True {
reg298657 := Nil;
reg298658 := PrimEqual(reg298657, V2789)
var reg298661 Obj
if reg298658 == True {
reg298659 := True;
reg298661 = reg298659
} else {
reg298660 := False;
reg298661 = reg298660
}
reg298663 = reg298661
} else {
reg298662 := False;
reg298663 = reg298662
}
if reg298663 == True {
reg298664 := MakeSymbol("Context_1957")
__ctx.Return(reg298664)
return
} else {
reg298665 := False;
reg298666 := PrimEqual(reg298665, V2788)
var reg298673 Obj
if reg298666 == True {
reg298667 := Nil;
reg298668 := PrimEqual(reg298667, V2789)
var reg298671 Obj
if reg298668 == True {
reg298669 := True;
reg298671 = reg298669
} else {
reg298670 := False;
reg298671 = reg298670
}
reg298673 = reg298671
} else {
reg298672 := False;
reg298673 = reg298672
}
if reg298673 == True {
reg298674 := MakeSymbol("ContextOut_1957")
__ctx.Return(reg298674)
return
} else {
reg298675 := PrimIsPair(V2789)
if reg298675 == True {
reg298676 := MakeSymbol("cons")
reg298677 := PrimHead(V2789)
reg298678 := __e.Call(__defun__shen_4recursive__cons__form, reg298677)
reg298679 := PrimTail(V2789)
reg298680 := __e.Call(__defun__shen_4construct_1context, V2788, reg298679)
reg298681 := Nil;
reg298682 := PrimCons(reg298680, reg298681)
reg298683 := PrimCons(reg298678, reg298682)
reg298684 := PrimCons(reg298676, reg298683)
__ctx.Return(reg298684)
return
} else {
reg298685 := MakeSymbol("shen.construct-context")
__ctx.TailApply(__defun__shen_4f__error, reg298685)
return
}
}
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.construct-context", value: __defun__shen_4construct_1context})

__defun__shen_4recursive__cons__form = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2791 := __args[0]
_ = V2791
reg298687 := PrimIsPair(V2791)
if reg298687 == True {
reg298688 := MakeSymbol("cons")
reg298689 := PrimHead(V2791)
reg298690 := __e.Call(__defun__shen_4recursive__cons__form, reg298689)
reg298691 := PrimTail(V2791)
reg298692 := __e.Call(__defun__shen_4recursive__cons__form, reg298691)
reg298693 := Nil;
reg298694 := PrimCons(reg298692, reg298693)
reg298695 := PrimCons(reg298690, reg298694)
reg298696 := PrimCons(reg298688, reg298695)
__ctx.Return(reg298696)
return
} else {
__ctx.Return(V2791)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.recursive_cons_form", value: __defun__shen_4recursive__cons__form})

__defun__preclude = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2793 := __args[0]
_ = V2793
reg298697 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4intern_1type, X)
return
}, 1)
reg298699 := __e.Call(__defun__map, reg298697, V2793)
__ctx.TailApply(__defun__shen_4preclude_1h, reg298699)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "preclude", value: __defun__preclude})

__defun__shen_4preclude_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2795 := __args[0]
_ = V2795
reg298701 := MakeSymbol("shen.*datatypes*")
reg298702 := MakeSymbol("shen.*datatypes*")
reg298703 := PrimValue(reg298702)
reg298704 := __e.Call(__defun__difference, reg298703, V2795)
reg298705 := PrimSet(reg298701, reg298704)
FilterDatatypes := reg298705
_ = FilterDatatypes
reg298706 := MakeSymbol("shen.*datatypes*")
reg298707 := PrimValue(reg298706)
__ctx.Return(reg298707)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.preclude-h", value: __defun__shen_4preclude_1h})

__defun__include = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2797 := __args[0]
_ = V2797
reg298708 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4intern_1type, X)
return
}, 1)
reg298710 := __e.Call(__defun__map, reg298708, V2797)
__ctx.TailApply(__defun__shen_4include_1h, reg298710)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "include", value: __defun__include})

__defun__shen_4include_1h = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2799 := __args[0]
_ = V2799
reg298712 := MakeSymbol("shen.*alldatatypes*")
reg298713 := PrimValue(reg298712)
reg298714 := __e.Call(__defun__intersection, V2799, reg298713)
ValidTypes := reg298714
_ = ValidTypes
reg298715 := MakeSymbol("shen.*datatypes*")
reg298716 := MakeSymbol("shen.*datatypes*")
reg298717 := PrimValue(reg298716)
reg298718 := __e.Call(__defun__union, ValidTypes, reg298717)
reg298719 := PrimSet(reg298715, reg298718)
NewDatatypes := reg298719
_ = NewDatatypes
reg298720 := MakeSymbol("shen.*datatypes*")
reg298721 := PrimValue(reg298720)
__ctx.Return(reg298721)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.include-h", value: __defun__shen_4include_1h})

__defun__preclude_1all_1but = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2801 := __args[0]
_ = V2801
reg298722 := MakeSymbol("shen.*alldatatypes*")
reg298723 := PrimValue(reg298722)
reg298724 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4intern_1type, X)
return
}, 1)
reg298726 := __e.Call(__defun__map, reg298724, V2801)
reg298727 := __e.Call(__defun__difference, reg298723, reg298726)
__ctx.TailApply(__defun__shen_4preclude_1h, reg298727)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "preclude-all-but", value: __defun__preclude_1all_1but})

__defun__include_1all_1but = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2803 := __args[0]
_ = V2803
reg298729 := MakeSymbol("shen.*alldatatypes*")
reg298730 := PrimValue(reg298729)
reg298731 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4intern_1type, X)
return
}, 1)
reg298733 := __e.Call(__defun__map, reg298731, V2803)
reg298734 := __e.Call(__defun__difference, reg298730, reg298733)
__ctx.TailApply(__defun__shen_4include_1h, reg298734)
return
}, 1)
__initDefs = append(__initDefs, defType{name: "include-all-but", value: __defun__include_1all_1but})

__defun__shen_4synonyms_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2809 := __args[0]
_ = V2809
reg298736 := Nil;
reg298737 := PrimEqual(reg298736, V2809)
if reg298737 == True {
reg298738 := MakeSymbol("shen.*tc*")
reg298739 := PrimValue(reg298738)
reg298740 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
X := __args[0]
_ = X
__ctx.TailApply(__defun__shen_4demod_1rule, X)
return
}, 1)
reg298742 := MakeSymbol("shen.*synonyms*")
reg298743 := PrimValue(reg298742)
reg298744 := __e.Call(__defun__mapcan, reg298740, reg298743)
__ctx.TailApply(__defun__shen_4update_1demodulation_1function, reg298739, reg298744)
return
} else {
reg298746 := PrimIsPair(V2809)
var reg298753 Obj
if reg298746 == True {
reg298747 := PrimTail(V2809)
reg298748 := PrimIsPair(reg298747)
var reg298751 Obj
if reg298748 == True {
reg298749 := True;
reg298751 = reg298749
} else {
reg298750 := False;
reg298751 = reg298750
}
reg298753 = reg298751
} else {
reg298752 := False;
reg298753 = reg298752
}
if reg298753 == True {
reg298754 := PrimTail(V2809)
reg298755 := PrimHead(reg298754)
reg298756 := __e.Call(__defun__shen_4extract__vars, reg298755)
reg298757 := PrimHead(V2809)
reg298758 := __e.Call(__defun__shen_4extract__vars, reg298757)
reg298759 := __e.Call(__defun__difference, reg298756, reg298758)
Vs := reg298759
_ = Vs
reg298760 := __e.Call(__defun__empty_2, Vs)
if reg298760 == True {
reg298761 := PrimHead(V2809)
reg298762 := PrimTail(V2809)
reg298763 := PrimHead(reg298762)
reg298764 := Nil;
reg298765 := PrimCons(reg298763, reg298764)
reg298766 := PrimCons(reg298761, reg298765)
reg298767 := MakeSymbol("shen.*synonyms*")
reg298768 := __e.Call(__defun__shen_4pushnew, reg298766, reg298767)
_ = reg298768
reg298769 := PrimTail(V2809)
reg298770 := PrimTail(reg298769)
__ctx.TailApply(__defun__shen_4synonyms_1help, reg298770)
return
} else {
reg298772 := PrimTail(V2809)
reg298773 := PrimHead(reg298772)
__ctx.TailApply(__defun__shen_4free__variable__warnings, reg298773, Vs)
return
}
} else {
reg298775 := MakeString("odd number of synonyms\n")
reg298776 := PrimSimpleError(reg298775)
__ctx.Return(reg298776)
return
}
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.synonyms-help", value: __defun__shen_4synonyms_1help})

__defun__shen_4pushnew = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2812 := __args[0]
_ = V2812
V2813 := __args[1]
_ = V2813
reg298777 := PrimValue(V2813)
reg298778 := __e.Call(__defun__element_2, V2812, reg298777)
if reg298778 == True {
reg298779 := PrimValue(V2813)
__ctx.Return(reg298779)
return
} else {
reg298780 := PrimValue(V2813)
reg298781 := PrimCons(V2812, reg298780)
reg298782 := PrimSet(V2813, reg298781)
__ctx.Return(reg298782)
return
}
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.pushnew", value: __defun__shen_4pushnew})

__defun__shen_4demod_1rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2815 := __args[0]
_ = V2815
reg298783 := PrimIsPair(V2815)
var reg298799 Obj
if reg298783 == True {
reg298784 := PrimTail(V2815)
reg298785 := PrimIsPair(reg298784)
var reg298794 Obj
if reg298785 == True {
reg298786 := Nil;
reg298787 := PrimTail(V2815)
reg298788 := PrimTail(reg298787)
reg298789 := PrimEqual(reg298786, reg298788)
var reg298792 Obj
if reg298789 == True {
reg298790 := True;
reg298792 = reg298790
} else {
reg298791 := False;
reg298792 = reg298791
}
reg298794 = reg298792
} else {
reg298793 := False;
reg298794 = reg298793
}
var reg298797 Obj
if reg298794 == True {
reg298795 := True;
reg298797 = reg298795
} else {
reg298796 := False;
reg298797 = reg298796
}
reg298799 = reg298797
} else {
reg298798 := False;
reg298799 = reg298798
}
if reg298799 == True {
reg298800 := PrimHead(V2815)
reg298801 := __e.Call(__defun__shen_4rcons__form, reg298800)
reg298802 := MakeSymbol("->")
reg298803 := PrimTail(V2815)
reg298804 := PrimHead(reg298803)
reg298805 := __e.Call(__defun__shen_4rcons__form, reg298804)
reg298806 := Nil;
reg298807 := PrimCons(reg298805, reg298806)
reg298808 := PrimCons(reg298802, reg298807)
reg298809 := PrimCons(reg298801, reg298808)
__ctx.Return(reg298809)
return
} else {
reg298810 := MakeSymbol("shen.demod-rule")
__ctx.TailApply(__defun__shen_4f__error, reg298810)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.demod-rule", value: __defun__shen_4demod_1rule})

__defun__shen_4lambda_1of_1defun = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2821 := __args[0]
_ = V2821
reg298812 := PrimIsPair(V2821)
var reg298875 Obj
if reg298812 == True {
reg298813 := MakeSymbol("defun")
reg298814 := PrimHead(V2821)
reg298815 := PrimEqual(reg298813, reg298814)
var reg298870 Obj
if reg298815 == True {
reg298816 := PrimTail(V2821)
reg298817 := PrimIsPair(reg298816)
var reg298865 Obj
if reg298817 == True {
reg298818 := PrimTail(V2821)
reg298819 := PrimTail(reg298818)
reg298820 := PrimIsPair(reg298819)
var reg298860 Obj
if reg298820 == True {
reg298821 := PrimTail(V2821)
reg298822 := PrimTail(reg298821)
reg298823 := PrimHead(reg298822)
reg298824 := PrimIsPair(reg298823)
var reg298855 Obj
if reg298824 == True {
reg298825 := Nil;
reg298826 := PrimTail(V2821)
reg298827 := PrimTail(reg298826)
reg298828 := PrimHead(reg298827)
reg298829 := PrimTail(reg298828)
reg298830 := PrimEqual(reg298825, reg298829)
var reg298850 Obj
if reg298830 == True {
reg298831 := PrimTail(V2821)
reg298832 := PrimTail(reg298831)
reg298833 := PrimTail(reg298832)
reg298834 := PrimIsPair(reg298833)
var reg298845 Obj
if reg298834 == True {
reg298835 := Nil;
reg298836 := PrimTail(V2821)
reg298837 := PrimTail(reg298836)
reg298838 := PrimTail(reg298837)
reg298839 := PrimTail(reg298838)
reg298840 := PrimEqual(reg298835, reg298839)
var reg298843 Obj
if reg298840 == True {
reg298841 := True;
reg298843 = reg298841
} else {
reg298842 := False;
reg298843 = reg298842
}
reg298845 = reg298843
} else {
reg298844 := False;
reg298845 = reg298844
}
var reg298848 Obj
if reg298845 == True {
reg298846 := True;
reg298848 = reg298846
} else {
reg298847 := False;
reg298848 = reg298847
}
reg298850 = reg298848
} else {
reg298849 := False;
reg298850 = reg298849
}
var reg298853 Obj
if reg298850 == True {
reg298851 := True;
reg298853 = reg298851
} else {
reg298852 := False;
reg298853 = reg298852
}
reg298855 = reg298853
} else {
reg298854 := False;
reg298855 = reg298854
}
var reg298858 Obj
if reg298855 == True {
reg298856 := True;
reg298858 = reg298856
} else {
reg298857 := False;
reg298858 = reg298857
}
reg298860 = reg298858
} else {
reg298859 := False;
reg298860 = reg298859
}
var reg298863 Obj
if reg298860 == True {
reg298861 := True;
reg298863 = reg298861
} else {
reg298862 := False;
reg298863 = reg298862
}
reg298865 = reg298863
} else {
reg298864 := False;
reg298865 = reg298864
}
var reg298868 Obj
if reg298865 == True {
reg298866 := True;
reg298868 = reg298866
} else {
reg298867 := False;
reg298868 = reg298867
}
reg298870 = reg298868
} else {
reg298869 := False;
reg298870 = reg298869
}
var reg298873 Obj
if reg298870 == True {
reg298871 := True;
reg298873 = reg298871
} else {
reg298872 := False;
reg298873 = reg298872
}
reg298875 = reg298873
} else {
reg298874 := False;
reg298875 = reg298874
}
if reg298875 == True {
reg298876 := MakeSymbol("/.")
reg298877 := PrimTail(V2821)
reg298878 := PrimTail(reg298877)
reg298879 := PrimHead(reg298878)
reg298880 := PrimHead(reg298879)
reg298881 := PrimTail(V2821)
reg298882 := PrimTail(reg298881)
reg298883 := PrimTail(reg298882)
reg298884 := PrimCons(reg298880, reg298883)
reg298885 := PrimCons(reg298876, reg298884)
__ctx.TailApply(__defun__eval, reg298885)
return
} else {
reg298887 := MakeSymbol("shen.lambda-of-defun")
__ctx.TailApply(__defun__shen_4f__error, reg298887)
return
}
}, 1)
__initDefs = append(__initDefs, defType{name: "shen.lambda-of-defun", value: __defun__shen_4lambda_1of_1defun})

__defun__shen_4update_1demodulation_1function = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
V2824 := __args[0]
_ = V2824
V2825 := __args[1]
_ = V2825
reg298889 := MakeSymbol("-")
reg298890 := __e.Call(__defun__tc, reg298889)
_ = reg298890
reg298891 := MakeSymbol("shen.*demodulation-function*")
reg298892 := MakeSymbol("define")
reg298893 := MakeSymbol("shen.demod")
reg298894 := __e.Call(__defun__shen_4default_1rule)
reg298895 := __e.Call(__defun__append, V2825, reg298894)
reg298896 := PrimCons(reg298893, reg298895)
reg298897 := PrimCons(reg298892, reg298896)
reg298898 := __e.Call(__defun__shen_4elim_1def, reg298897)
reg298899 := __e.Call(__defun__shen_4lambda_1of_1defun, reg298898)
reg298900 := PrimSet(reg298891, reg298899)
_ = reg298900
var reg298904 Obj
if V2824 == True {
reg298901 := MakeSymbol("+")
reg298902 := __e.Call(__defun__tc, reg298901)
reg298904 = reg298902
} else {
reg298903 := MakeSymbol("shen.skip")
reg298904 = reg298903
}
_ = reg298904
reg298905 := MakeSymbol("synonyms")
__ctx.Return(reg298905)
return
}, 2)
__initDefs = append(__initDefs, defType{name: "shen.update-demodulation-function", value: __defun__shen_4update_1demodulation_1function})

__defun__shen_4default_1rule = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
reg298906 := MakeSymbol("X")
reg298907 := MakeSymbol("->")
reg298908 := MakeSymbol("X")
reg298909 := Nil;
reg298910 := PrimCons(reg298908, reg298909)
reg298911 := PrimCons(reg298907, reg298910)
reg298912 := PrimCons(reg298906, reg298911)
__ctx.Return(reg298912)
return
}, 0)
__initDefs = append(__initDefs, defType{name: "shen.default-rule", value: __defun__shen_4default_1rule})

}
