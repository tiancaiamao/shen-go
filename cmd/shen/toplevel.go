package main

import . "github.com/tiancaiamao/shen-go/kl"

var TopLevelMain = MakeNative(func(__e *ControlFlow) {
tmp8415 := MakeNative(func(__e *ControlFlow) {
tmp8416 := Call(__e, PrimFunc(symshen_4credits))


_ = tmp8416

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8417 := Call(__e, ns2_1set, symshen_4shen, tmp8415)


_ = tmp8417

tmp8418 := MakeNative(func(__e *ControlFlow) {
tmp8419 := Call(__e, PrimFunc(symshen_4initialise__environment))


_ = tmp8419

tmp8420 := Call(__e, PrimFunc(symshen_4prompt))


_ = tmp8420

tmp8421 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4read_1evaluate_1print))
return
}, 0)

tmp8422 := MakeNative(func(__e *ControlFlow) {
Z5297 := __e.Get(1)
_ = Z5297
tmp8423 := PrimErrorToString(Z5297)

tmp8424 := Call(__e, PrimFunc(symstoutput))


tmp8425 := Call(__e, PrimFunc(sympr), tmp8423, tmp8424)


_ = tmp8425

__e.TailApply(PrimFunc(symnl), MakeNumber(0))
return


}, 1)

tmp8426 := Call(__e, try_1catch, tmp8421, tmp8422)


_ = tmp8426

__e.TailApply(PrimFunc(symshen_4loop))
return


}, 0)

tmp8427 := Call(__e, ns2_1set, symshen_4loop, tmp8418)


_ = tmp8427

tmp8428 := MakeNative(func(__e *ControlFlow) {
tmp8429 := Call(__e, PrimFunc(symstoutput))


tmp8430 := Call(__e, PrimFunc(sympr), MakeString("\nShen, www.shenlanguage.org, copyright (C) 2010-2024, Mark Tarver\n"), tmp8429)


_ = tmp8430

tmp8431 := PrimValue(sym_dversion_d)

tmp8432 := PrimValue(sym_dlanguage_d)

tmp8433 := PrimValue(sym_dimplementation_d)

tmp8434 := PrimValue(sym_drelease_d)

tmp8435 := Call(__e, PrimFunc(symshen_4app), tmp8434, MakeString("\n"), symshen_4a)


tmp8436 := PrimStringConcat(MakeString(" "), tmp8435)

tmp8437 := Call(__e, PrimFunc(symshen_4app), tmp8433, tmp8436, symshen_4a)


tmp8438 := PrimStringConcat(MakeString(", platform: "), tmp8437)

tmp8439 := Call(__e, PrimFunc(symshen_4app), tmp8432, tmp8438, symshen_4a)


tmp8440 := PrimStringConcat(MakeString(", language: "), tmp8439)

tmp8441 := Call(__e, PrimFunc(symshen_4app), tmp8431, tmp8440, symshen_4a)


tmp8442 := PrimStringConcat(MakeString("version: S"), tmp8441)

tmp8443 := Call(__e, PrimFunc(symstoutput))


tmp8444 := Call(__e, PrimFunc(sympr), tmp8442, tmp8443)


_ = tmp8444

tmp8445 := PrimValue(sym_dport_d)

tmp8446 := PrimValue(sym_dporters_d)

tmp8447 := Call(__e, PrimFunc(symshen_4app), tmp8446, MakeString("\n\n"), symshen_4a)


tmp8448 := PrimStringConcat(MakeString(", ported by "), tmp8447)

tmp8449 := Call(__e, PrimFunc(symshen_4app), tmp8445, tmp8448, symshen_4a)


tmp8450 := PrimStringConcat(MakeString("port "), tmp8449)

tmp8451 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8450, tmp8451)
return


}, 0)

tmp8452 := Call(__e, ns2_1set, symshen_4credits, tmp8428)


_ = tmp8452

tmp8453 := MakeNative(func(__e *ControlFlow) {
tmp8454 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp8454

__e.Return(PrimSet(symshen_4_dinfs_d, MakeNumber(0)))
return


}, 0)

tmp8455 := Call(__e, ns2_1set, symshen_4initialise__environment, tmp8453)


_ = tmp8455

tmp8456 := MakeNative(func(__e *ControlFlow) {
tmp8468 := PrimValue(symshen_4_dtc_d)

if True == tmp8468 {
tmp8457 := PrimValue(symshen_4_dhistory_d)

tmp8458 := Call(__e, PrimFunc(symlength), tmp8457)


tmp8459 := Call(__e, PrimFunc(symshen_4app), tmp8458, MakeString("+) "), symshen_4a)


tmp8460 := PrimStringConcat(MakeString("\n("), tmp8459)

tmp8461 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8460, tmp8461)
return


} else {
tmp8462 := PrimValue(symshen_4_dhistory_d)

tmp8463 := Call(__e, PrimFunc(symlength), tmp8462)


tmp8464 := Call(__e, PrimFunc(symshen_4app), tmp8463, MakeString("-) "), symshen_4a)


tmp8465 := PrimStringConcat(MakeString("\n("), tmp8464)

tmp8466 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8465, tmp8466)
return


}


}, 0)

tmp8469 := Call(__e, ns2_1set, symshen_4prompt, tmp8456)


_ = tmp8469

tmp8470 := MakeNative(func(__e *ControlFlow) {
tmp8471 := MakeNative(func(__e *ControlFlow) {
W5298 := __e.Get(1)
_ = W5298
tmp8472 := MakeNative(func(__e *ControlFlow) {
W5299 := __e.Get(1)
_ = W5299
tmp8473 := MakeNative(func(__e *ControlFlow) {
W5300 := __e.Get(1)
_ = W5300
tmp8474 := PrimValue(symshen_4_dtc_d)

__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5299, W5300, tmp8474)
return


}, 1)

tmp8475 := Call(__e, PrimFunc(symshen_4update_1history))


__e.TailApply(tmp8473, tmp8475)
return


}, 1)

tmp8476 := Call(__e, PrimFunc(symstinput))


tmp8477 := Call(__e, PrimFunc(symlineread), tmp8476)


tmp8478 := Call(__e, PrimFunc(symshen_4package_1user_1input), W5298, tmp8477)


__e.TailApply(tmp8472, tmp8478)
return


}, 1)

tmp8479 := PrimValue(symshen_4_dpackage_d)

__e.TailApply(tmp8471, tmp8479)
return


}, 0)

tmp8480 := Call(__e, ns2_1set, symshen_4read_1evaluate_1print, tmp8470)


_ = tmp8480

tmp8481 := MakeNative(func(__e *ControlFlow) {
V5301 := __e.Get(1)
_ = V5301
V5302 := __e.Get(2)
_ = V5302
tmp8488 := PrimEqual(symnull, V5301)

if True == tmp8488 {
__e.Return(V5302)
return
} else {
tmp8482 := MakeNative(func(__e *ControlFlow) {
W5303 := __e.Get(1)
_ = W5303
tmp8483 := MakeNative(func(__e *ControlFlow) {
W5304 := __e.Get(1)
_ = W5304
tmp8484 := MakeNative(func(__e *ControlFlow) {
Z5305 := __e.Get(1)
_ = Z5305
__e.TailApply(PrimFunc(symshen_4pui_1h), W5303, W5304, Z5305)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8484, V5302)
return


}, 1)

tmp8485 := Call(__e, PrimFunc(symexternal), V5301)


__e.TailApply(tmp8483, tmp8485)
return


}, 1)

tmp8486 := PrimStr(V5301)

__e.TailApply(tmp8482, tmp8486)
return


}


}, 2)

tmp8489 := Call(__e, ns2_1set, symshen_4package_1user_1input, tmp8481)


_ = tmp8489

tmp8490 := MakeNative(func(__e *ControlFlow) {
V5310 := __e.Get(1)
_ = V5310
V5311 := __e.Get(2)
_ = V5311
V5312 := __e.Get(3)
_ = V5312
tmp8531 := PrimIsPair(V5312)

var ifres8518 Obj

if True == tmp8531 {
tmp8529 := PrimHead(V5312)

tmp8530 := PrimEqual(symfn, tmp8529)

var ifres8520 Obj

if True == tmp8530 {
tmp8527 := PrimTail(V5312)

tmp8528 := PrimIsPair(tmp8527)

var ifres8522 Obj

if True == tmp8528 {
tmp8524 := PrimTail(V5312)

tmp8525 := PrimTail(tmp8524)

tmp8526 := PrimEqual(Nil, tmp8525)

var ifres8523 Obj

if True == tmp8526 {
ifres8523 = True


} else {
ifres8523 = False


}

ifres8522 = ifres8523


} else {
ifres8522 = False


}

var ifres8521 Obj

if True == ifres8522 {
ifres8521 = True


} else {
ifres8521 = False


}

ifres8520 = ifres8521


} else {
ifres8520 = False


}

var ifres8519 Obj

if True == ifres8520 {
ifres8519 = True


} else {
ifres8519 = False


}

ifres8518 = ifres8519


} else {
ifres8518 = False


}

if True == ifres8518 {
tmp8496 := PrimTail(V5312)

tmp8497 := PrimHead(tmp8496)

tmp8498 := Call(__e, PrimFunc(symshen_4internal_2), tmp8497, V5310, V5311)


if True == tmp8498 {
tmp8491 := PrimTail(V5312)

tmp8492 := PrimHead(tmp8491)

tmp8493 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5310, tmp8492)


tmp8494 := PrimCons(tmp8493, Nil)

__e.Return(PrimCons(symfn, tmp8494))
return


} else {
__e.Return(V5312)
return
}


} else {
tmp8516 := PrimIsPair(V5312)

if True == tmp8516 {
tmp8513 := PrimHead(V5312)

tmp8514 := Call(__e, PrimFunc(symshen_4internal_2), tmp8513, V5310, V5311)


if True == tmp8514 {
tmp8499 := PrimHead(V5312)

tmp8500 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V5310, tmp8499)


tmp8501 := MakeNative(func(__e *ControlFlow) {
Z5313 := __e.Get(1)
_ = Z5313
__e.TailApply(PrimFunc(symshen_4pui_1h), V5310, V5311, Z5313)
return
}, 1)

tmp8502 := PrimTail(V5312)

tmp8503 := Call(__e, PrimFunc(symmap), tmp8501, tmp8502)


__e.Return(PrimCons(tmp8500, tmp8503))
return


} else {
tmp8510 := PrimHead(V5312)

tmp8511 := PrimIsPair(tmp8510)

if True == tmp8511 {
tmp8504 := MakeNative(func(__e *ControlFlow) {
Z5314 := __e.Get(1)
_ = Z5314
__e.TailApply(PrimFunc(symshen_4pui_1h), V5310, V5311, Z5314)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp8504, V5312)
return


} else {
tmp8505 := PrimHead(V5312)

tmp8506 := MakeNative(func(__e *ControlFlow) {
Z5315 := __e.Get(1)
_ = Z5315
__e.TailApply(PrimFunc(symshen_4pui_1h), V5310, V5311, Z5315)
return
}, 1)

tmp8507 := PrimTail(V5312)

tmp8508 := Call(__e, PrimFunc(symmap), tmp8506, tmp8507)


__e.Return(PrimCons(tmp8505, tmp8508))
return


}


}


} else {
__e.Return(V5312)
return
}


}


}, 3)

tmp8532 := Call(__e, ns2_1set, symshen_4pui_1h, tmp8490)


_ = tmp8532

tmp8533 := MakeNative(func(__e *ControlFlow) {
tmp8534 := Call(__e, PrimFunc(symit))


tmp8535 := Call(__e, PrimFunc(symshen_4trim_1it), tmp8534)


tmp8536 := PrimValue(symshen_4_dhistory_d)

tmp8537 := PrimCons(tmp8535, tmp8536)

__e.Return(PrimSet(symshen_4_dhistory_d, tmp8537))
return


}, 0)

tmp8538 := Call(__e, ns2_1set, symshen_4update_1history, tmp8533)


_ = tmp8538

tmp8539 := MakeNative(func(__e *ControlFlow) {
V5316 := __e.Get(1)
_ = V5316
tmp8547 := Call(__e, PrimFunc(symshen_4_7string_2), V5316)


var ifres8542 Obj

if True == tmp8547 {
tmp8544 := Call(__e, PrimFunc(symhdstr), V5316)


tmp8545 := PrimStringToNumber(tmp8544)

tmp8546 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8545)


var ifres8543 Obj

if True == tmp8546 {
ifres8543 = True


} else {
ifres8543 = False


}

ifres8542 = ifres8543


} else {
ifres8542 = False


}

if True == ifres8542 {
tmp8540 := PrimTailString(V5316)

__e.TailApply(PrimFunc(symshen_4trim_1it), tmp8540)
return


} else {
__e.Return(V5316)
return
}


}, 1)

tmp8548 := Call(__e, ns2_1set, symshen_4trim_1it, tmp8539)


_ = tmp8548

tmp8549 := MakeNative(func(__e *ControlFlow) {
V5335 := __e.Get(1)
_ = V5335
V5336 := __e.Get(2)
_ = V5336
V5337 := __e.Get(3)
_ = V5337
tmp8679 := PrimIsPair(V5335)

var ifres8648 Obj

if True == tmp8679 {
tmp8677 := PrimTail(V5335)

tmp8678 := PrimEqual(Nil, tmp8677)

var ifres8650 Obj

if True == tmp8678 {
tmp8676 := PrimIsPair(V5336)

var ifres8652 Obj

if True == tmp8676 {
tmp8674 := PrimHead(V5336)

tmp8675 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8674)


var ifres8654 Obj

if True == tmp8675 {
tmp8671 := PrimHead(V5336)

tmp8672 := Call(__e, PrimFunc(symhdstr), tmp8671)


tmp8673 := PrimEqual(MakeString("!"), tmp8672)

var ifres8656 Obj

if True == tmp8673 {
tmp8668 := PrimHead(V5336)

tmp8669 := PrimTailString(tmp8668)

tmp8670 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8669)


var ifres8658 Obj

if True == tmp8670 {
tmp8664 := PrimHead(V5336)

tmp8665 := PrimTailString(tmp8664)

tmp8666 := Call(__e, PrimFunc(symhdstr), tmp8665)


tmp8667 := PrimEqual(MakeString("!"), tmp8666)

var ifres8660 Obj

if True == tmp8667 {
tmp8662 := PrimTail(V5336)

tmp8663 := PrimIsPair(tmp8662)

var ifres8661 Obj

if True == tmp8663 {
ifres8661 = True


} else {
ifres8661 = False


}

ifres8660 = ifres8661


} else {
ifres8660 = False


}

var ifres8659 Obj

if True == ifres8660 {
ifres8659 = True


} else {
ifres8659 = False


}

ifres8658 = ifres8659


} else {
ifres8658 = False


}

var ifres8657 Obj

if True == ifres8658 {
ifres8657 = True


} else {
ifres8657 = False


}

ifres8656 = ifres8657


} else {
ifres8656 = False


}

var ifres8655 Obj

if True == ifres8656 {
ifres8655 = True


} else {
ifres8655 = False


}

ifres8654 = ifres8655


} else {
ifres8654 = False


}

var ifres8653 Obj

if True == ifres8654 {
ifres8653 = True


} else {
ifres8653 = False


}

ifres8652 = ifres8653


} else {
ifres8652 = False


}

var ifres8651 Obj

if True == ifres8652 {
ifres8651 = True


} else {
ifres8651 = False


}

ifres8650 = ifres8651


} else {
ifres8650 = False


}

var ifres8649 Obj

if True == ifres8650 {
ifres8649 = True


} else {
ifres8649 = False


}

ifres8648 = ifres8649


} else {
ifres8648 = False


}

if True == ifres8648 {
tmp8550 := MakeNative(func(__e *ControlFlow) {
W5338 := __e.Get(1)
_ = W5338
tmp8551 := MakeNative(func(__e *ControlFlow) {
W5339 := __e.Get(1)
_ = W5339
tmp8552 := MakeNative(func(__e *ControlFlow) {
W5340 := __e.Get(1)
_ = W5340
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5338, W5339, V5337)
return
}, 1)

tmp8553 := PrimTail(V5336)

tmp8554 := PrimHead(tmp8553)

tmp8555 := Call(__e, PrimFunc(symshen_4app), tmp8554, MakeString("\n"), symshen_4a)


tmp8556 := Call(__e, PrimFunc(symstoutput))


tmp8557 := Call(__e, PrimFunc(sympr), tmp8555, tmp8556)


__e.TailApply(tmp8552, tmp8557)
return


}, 1)

tmp8558 := PrimTail(V5336)

tmp8559 := PrimHead(tmp8558)

tmp8560 := PrimTail(V5336)

tmp8561 := PrimCons(tmp8559, tmp8560)

tmp8562 := PrimSet(symshen_4_dhistory_d, tmp8561)

__e.TailApply(tmp8551, tmp8562)
return


}, 1)

tmp8563 := PrimTail(V5336)

tmp8564 := PrimHead(tmp8563)

tmp8565 := Call(__e, PrimFunc(symread_1from_1string), tmp8564)


__e.TailApply(tmp8550, tmp8565)
return


} else {
tmp8646 := PrimIsPair(V5335)

var ifres8630 Obj

if True == tmp8646 {
tmp8644 := PrimTail(V5335)

tmp8645 := PrimEqual(Nil, tmp8644)

var ifres8632 Obj

if True == tmp8645 {
tmp8643 := PrimIsPair(V5336)

var ifres8634 Obj

if True == tmp8643 {
tmp8641 := PrimHead(V5336)

tmp8642 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8641)


var ifres8636 Obj

if True == tmp8642 {
tmp8638 := PrimHead(V5336)

tmp8639 := Call(__e, PrimFunc(symhdstr), tmp8638)


tmp8640 := PrimEqual(MakeString("!"), tmp8639)

var ifres8637 Obj

if True == tmp8640 {
ifres8637 = True


} else {
ifres8637 = False


}

ifres8636 = ifres8637


} else {
ifres8636 = False


}

var ifres8635 Obj

if True == ifres8636 {
ifres8635 = True


} else {
ifres8635 = False


}

ifres8634 = ifres8635


} else {
ifres8634 = False


}

var ifres8633 Obj

if True == ifres8634 {
ifres8633 = True


} else {
ifres8633 = False


}

ifres8632 = ifres8633


} else {
ifres8632 = False


}

var ifres8631 Obj

if True == ifres8632 {
ifres8631 = True


} else {
ifres8631 = False


}

ifres8630 = ifres8631


} else {
ifres8630 = False


}

if True == ifres8630 {
tmp8566 := MakeNative(func(__e *ControlFlow) {
W5341 := __e.Get(1)
_ = W5341
tmp8567 := MakeNative(func(__e *ControlFlow) {
W5342 := __e.Get(1)
_ = W5342
tmp8568 := MakeNative(func(__e *ControlFlow) {
W5343 := __e.Get(1)
_ = W5343
tmp8569 := MakeNative(func(__e *ControlFlow) {
W5344 := __e.Get(1)
_ = W5344
tmp8570 := MakeNative(func(__e *ControlFlow) {
W5345 := __e.Get(1)
_ = W5345
__e.TailApply(PrimFunc(symshen_4evaluate_1lineread), W5344, W5345, V5337)
return
}, 1)

tmp8571 := PrimTail(V5336)

tmp8572 := PrimCons(W5342, tmp8571)

tmp8573 := PrimSet(symshen_4_dhistory_d, tmp8572)

__e.TailApply(tmp8570, tmp8573)
return


}, 1)

tmp8574 := Call(__e, PrimFunc(symread_1from_1string), W5342)


__e.TailApply(tmp8569, tmp8574)
return


}, 1)

tmp8575 := Call(__e, PrimFunc(symshen_4app), W5342, MakeString("\n"), symshen_4a)


tmp8576 := Call(__e, PrimFunc(symstoutput))


tmp8577 := Call(__e, PrimFunc(sympr), tmp8575, tmp8576)


__e.TailApply(tmp8568, tmp8577)
return


}, 1)

tmp8578 := PrimHead(V5336)

tmp8579 := PrimTailString(tmp8578)

tmp8580 := PrimTail(V5336)

tmp8581 := Call(__e, PrimFunc(symshen_4use_1history), W5341, tmp8579, tmp8580)


__e.TailApply(tmp8567, tmp8581)
return


}, 1)

tmp8587 := PrimHead(V5336)

tmp8588 := PrimTailString(tmp8587)

tmp8589 := PrimEqual(tmp8588, MakeString(""))

var ifres8582 Obj

if True == tmp8589 {
ifres8582 = Nil


} else {
tmp8583 := PrimHead(V5336)

tmp8584 := PrimTailString(tmp8583)

tmp8585 := Call(__e, PrimFunc(symread_1from_1string), tmp8584)


tmp8586 := PrimHead(tmp8585)

ifres8582 = tmp8586


}

__e.TailApply(tmp8566, ifres8582)
return


} else {
tmp8628 := PrimIsPair(V5335)

var ifres8612 Obj

if True == tmp8628 {
tmp8626 := PrimTail(V5335)

tmp8627 := PrimEqual(Nil, tmp8626)

var ifres8614 Obj

if True == tmp8627 {
tmp8625 := PrimIsPair(V5336)

var ifres8616 Obj

if True == tmp8625 {
tmp8623 := PrimHead(V5336)

tmp8624 := Call(__e, PrimFunc(symshen_4_7string_2), tmp8623)


var ifres8618 Obj

if True == tmp8624 {
tmp8620 := PrimHead(V5336)

tmp8621 := Call(__e, PrimFunc(symhdstr), tmp8620)


tmp8622 := PrimEqual(MakeString("%"), tmp8621)

var ifres8619 Obj

if True == tmp8622 {
ifres8619 = True


} else {
ifres8619 = False


}

ifres8618 = ifres8619


} else {
ifres8618 = False


}

var ifres8617 Obj

if True == ifres8618 {
ifres8617 = True


} else {
ifres8617 = False


}

ifres8616 = ifres8617


} else {
ifres8616 = False


}

var ifres8615 Obj

if True == ifres8616 {
ifres8615 = True


} else {
ifres8615 = False


}

ifres8614 = ifres8615


} else {
ifres8614 = False


}

var ifres8613 Obj

if True == ifres8614 {
ifres8613 = True


} else {
ifres8613 = False


}

ifres8612 = ifres8613


} else {
ifres8612 = False


}

if True == ifres8612 {
tmp8590 := MakeNative(func(__e *ControlFlow) {
W5346 := __e.Get(1)
_ = W5346
tmp8591 := MakeNative(func(__e *ControlFlow) {
W5347 := __e.Get(1)
_ = W5347
tmp8592 := MakeNative(func(__e *ControlFlow) {
W5348 := __e.Get(1)
_ = W5348
__e.TailApply(PrimFunc(symabort))
return
}, 1)

tmp8593 := PrimTail(V5336)

tmp8594 := PrimSet(symshen_4_dhistory_d, tmp8593)

__e.TailApply(tmp8592, tmp8594)
return


}, 1)

tmp8595 := PrimHead(V5336)

tmp8596 := PrimTailString(tmp8595)

tmp8597 := PrimTail(V5336)

tmp8598 := Call(__e, PrimFunc(symshen_4peek_1history), W5346, tmp8596, tmp8597)


__e.TailApply(tmp8591, tmp8598)
return


}, 1)

tmp8604 := PrimHead(V5336)

tmp8605 := PrimTailString(tmp8604)

tmp8606 := PrimEqual(tmp8605, MakeString(""))

var ifres8599 Obj

if True == tmp8606 {
ifres8599 = Nil


} else {
tmp8600 := PrimHead(V5336)

tmp8601 := PrimTailString(tmp8600)

tmp8602 := Call(__e, PrimFunc(symread_1from_1string), tmp8601)


tmp8603 := PrimHead(tmp8602)

ifres8599 = tmp8603


}

__e.TailApply(tmp8590, ifres8599)
return


} else {
tmp8610 := PrimEqual(True, V5337)

if True == tmp8610 {
__e.TailApply(PrimFunc(symshen_4check_1eval_1and_1print), V5335)
return
} else {
tmp8608 := PrimEqual(False, V5337)

if True == tmp8608 {
__e.TailApply(PrimFunc(symshen_4eval_1and_1print), V5335)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.evaluate-lineread")))
return
}


}


}


}


}


}, 3)

tmp8680 := Call(__e, ns2_1set, symshen_4evaluate_1lineread, tmp8549)


_ = tmp8680

tmp8681 := MakeNative(func(__e *ControlFlow) {
V5349 := __e.Get(1)
_ = V5349
V5350 := __e.Get(2)
_ = V5350
V5351 := __e.Get(3)
_ = V5351
tmp8687 := PrimIsInteger(V5349)

if True == tmp8687 {
tmp8682 := PrimNumberAdd(MakeNumber(1), V5349)

tmp8683 := Call(__e, PrimFunc(symreverse), V5351)


__e.TailApply(PrimFunc(symnth), tmp8682, tmp8683)
return


} else {
tmp8685 := PrimIsSymbol(V5349)

if True == tmp8685 {
__e.TailApply(PrimFunc(symshen_4string_1match), V5350, V5351)
return
} else {
__e.Return(PrimSimpleError(MakeString("! expects a number or a symbol\n")))
return
}


}


}, 3)

tmp8688 := Call(__e, ns2_1set, symshen_4use_1history, tmp8681)


_ = tmp8688

tmp8689 := MakeNative(func(__e *ControlFlow) {
V5352 := __e.Get(1)
_ = V5352
V5353 := __e.Get(2)
_ = V5353
V5354 := __e.Get(3)
_ = V5354
tmp8703 := PrimIsInteger(V5352)

if True == tmp8703 {
tmp8690 := PrimNumberAdd(MakeNumber(1), V5352)

tmp8691 := Call(__e, PrimFunc(symreverse), V5354)


tmp8692 := Call(__e, PrimFunc(symnth), tmp8690, tmp8691)


tmp8693 := Call(__e, PrimFunc(symshen_4app), tmp8692, MakeString(""), symshen_4a)


tmp8694 := PrimStringConcat(MakeString("\n"), tmp8693)

tmp8695 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp8694, tmp8695)
return


} else {
tmp8701 := PrimEqual(V5353, MakeString(""))

var ifres8698 Obj

if True == tmp8701 {
ifres8698 = True


} else {
tmp8700 := PrimIsSymbol(V5352)

var ifres8699 Obj

if True == tmp8700 {
ifres8699 = True


} else {
ifres8699 = False


}

ifres8698 = ifres8699


}

if True == ifres8698 {
tmp8696 := Call(__e, PrimFunc(symreverse), V5354)


__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), MakeNumber(0), V5353, tmp8696)
return


} else {
__e.Return(PrimSimpleError(MakeString("% expects a number or a symbol\n")))
return
}


}


}, 3)

tmp8704 := Call(__e, ns2_1set, symshen_4peek_1history, tmp8689)


_ = tmp8704

tmp8705 := MakeNative(func(__e *ControlFlow) {
V5364 := __e.Get(1)
_ = V5364
V5365 := __e.Get(2)
_ = V5365
tmp8716 := PrimEqual(Nil, V5365)

if True == tmp8716 {
__e.Return(PrimSimpleError(MakeString("\ninput not found")))
return
} else {
tmp8714 := PrimIsPair(V5365)

var ifres8710 Obj

if True == tmp8714 {
tmp8712 := PrimHead(V5365)

tmp8713 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5364, tmp8712)


var ifres8711 Obj

if True == tmp8713 {
ifres8711 = True


} else {
ifres8711 = False


}

ifres8710 = ifres8711


} else {
ifres8710 = False


}

if True == ifres8710 {
__e.Return(PrimHead(V5365))
return
} else {
tmp8708 := PrimIsPair(V5365)

if True == tmp8708 {
tmp8706 := PrimTail(V5365)

__e.TailApply(PrimFunc(symshen_4string_1match), V5364, tmp8706)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.string-match")))
return
}


}


}


}, 2)

tmp8717 := Call(__e, ns2_1set, symshen_4string_1match, tmp8705)


_ = tmp8717

tmp8718 := MakeNative(func(__e *ControlFlow) {
V5373 := __e.Get(1)
_ = V5373
V5374 := __e.Get(2)
_ = V5374
tmp8755 := PrimEqual(MakeString(""), V5373)

if True == tmp8755 {
__e.Return(True)
return
} else {
tmp8753 := Call(__e, PrimFunc(symshen_4_7string_2), V5373)


var ifres8748 Obj

if True == tmp8753 {
tmp8750 := Call(__e, PrimFunc(symhdstr), V5373)


tmp8751 := PrimStringToNumber(tmp8750)

tmp8752 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8751)


var ifres8749 Obj

if True == tmp8752 {
ifres8749 = True


} else {
ifres8749 = False


}

ifres8748 = ifres8749


} else {
ifres8748 = False


}

if True == ifres8748 {
tmp8719 := PrimTailString(V5373)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8719, V5374)
return


} else {
tmp8746 := Call(__e, PrimFunc(symshen_4_7string_2), V5374)


var ifres8741 Obj

if True == tmp8746 {
tmp8743 := Call(__e, PrimFunc(symhdstr), V5374)


tmp8744 := PrimStringToNumber(tmp8743)

tmp8745 := Call(__e, PrimFunc(symshen_4whitespace_2), tmp8744)


var ifres8742 Obj

if True == tmp8745 {
ifres8742 = True


} else {
ifres8742 = False


}

ifres8741 = ifres8742


} else {
ifres8741 = False


}

if True == ifres8741 {
tmp8720 := PrimTailString(V5374)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5373, tmp8720)
return


} else {
tmp8739 := Call(__e, PrimFunc(symshen_4_7string_2), V5374)


var ifres8735 Obj

if True == tmp8739 {
tmp8737 := Call(__e, PrimFunc(symhdstr), V5374)


tmp8738 := PrimEqual(MakeString("("), tmp8737)

var ifres8736 Obj

if True == tmp8738 {
ifres8736 = True


} else {
ifres8736 = False


}

ifres8735 = ifres8736


} else {
ifres8735 = False


}

if True == ifres8735 {
tmp8721 := PrimTailString(V5374)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), V5373, tmp8721)
return


} else {
tmp8733 := Call(__e, PrimFunc(symshen_4_7string_2), V5373)


var ifres8725 Obj

if True == tmp8733 {
tmp8732 := Call(__e, PrimFunc(symshen_4_7string_2), V5374)


var ifres8727 Obj

if True == tmp8732 {
tmp8729 := Call(__e, PrimFunc(symhdstr), V5373)


tmp8730 := Call(__e, PrimFunc(symhdstr), V5374)


tmp8731 := PrimEqual(tmp8729, tmp8730)

var ifres8728 Obj

if True == tmp8731 {
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

if True == ifres8725 {
tmp8722 := PrimTailString(V5373)

tmp8723 := PrimTailString(V5374)

__e.TailApply(PrimFunc(symshen_4string_1prefix_2), tmp8722, tmp8723)
return


} else {
__e.Return(False)
return
}


}


}


}


}


}, 2)

tmp8756 := Call(__e, ns2_1set, symshen_4string_1prefix_2, tmp8718)


_ = tmp8756

tmp8757 := MakeNative(func(__e *ControlFlow) {
V5385 := __e.Get(1)
_ = V5385
V5386 := __e.Get(2)
_ = V5386
V5387 := __e.Get(3)
_ = V5387
tmp8772 := PrimEqual(Nil, V5387)

if True == tmp8772 {
__e.Return(symshen_4skip)
return
} else {
tmp8770 := PrimIsPair(V5387)

if True == tmp8770 {
tmp8765 := PrimHead(V5387)

tmp8766 := Call(__e, PrimFunc(symshen_4string_1prefix_2), V5386, tmp8765)


var ifres8758 Obj

if True == tmp8766 {
tmp8759 := PrimHead(V5387)

tmp8760 := Call(__e, PrimFunc(symshen_4app), tmp8759, MakeString("\n"), symshen_4a)


tmp8761 := PrimStringConcat(MakeString(". "), tmp8760)

tmp8762 := Call(__e, PrimFunc(symshen_4app), V5385, tmp8761, symshen_4a)


tmp8763 := Call(__e, PrimFunc(symstoutput))


tmp8764 := Call(__e, PrimFunc(sympr), tmp8762, tmp8763)


ifres8758 = tmp8764


} else {
ifres8758 = symshen_4skip


}

_ = ifres8758

tmp8767 := PrimNumberAdd(V5385, MakeNumber(1))

tmp8768 := PrimTail(V5387)

__e.TailApply(PrimFunc(symshen_4recursive_1string_1match), tmp8767, V5386, tmp8768)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursive-string-match")))
return
}


}


}, 3)

__e.TailApply(ns2_1set, symshen_4recursive_1string_1match, tmp8757)
return




}, 0)

