package main

import . "github.com/tiancaiamao/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
tmp18362 := MakeNative(func(__e *ControlFlow) {
V5536 := __e.Get(1)
_ = V5536
V5537 := __e.Get(2)
_ = V5537
tmp18363 := MakeNative(func(__e *ControlFlow) {
W5538 := __e.Get(1)
_ = W5538
tmp18364 := MakeNative(func(__e *ControlFlow) {
W5539 := __e.Get(1)
_ = W5539
tmp18365 := MakeNative(func(__e *ControlFlow) {
W5544 := __e.Get(1)
_ = W5544
tmp18366 := MakeNative(func(__e *ControlFlow) {
W5545 := __e.Get(1)
_ = W5545
__e.Return(V5536)
return
}, 1)

tmp18367 := PrimValue(symshen_4_dsigf_d)

tmp18368 := Call(__e, PrimFunc(symshen_4assoc_1_6), V5536, W5544, tmp18367)


tmp18369 := PrimSet(symshen_4_dsigf_d, tmp18368)

__e.TailApply(tmp18366, tmp18369)
return


}, 1)

tmp18370 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V5537)


tmp18371 := Call(__e, PrimFunc(symeval_1kl), tmp18370)


__e.TailApply(tmp18365, tmp18371)
return


}, 1)

tmp18372 := MakeNative(func(__e *ControlFlow) {
Z5540 := __e.Get(1)
_ = Z5540
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5541 := __e.Get(1)
_ = Z5541
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5542 := __e.Get(1)
_ = Z5542
__e.Return(MakeNative(func(__e *ControlFlow) {
Z5543 := __e.Get(1)
_ = Z5543
tmp18373 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18373

tmp18374 := Call(__e, PrimFunc(symshen_4deref), V5536, Z5540)


tmp18375 := Call(__e, PrimFunc(symreceive), tmp18374)


tmp18376 := Call(__e, PrimFunc(symshen_4deref), W5538, Z5540)


tmp18377 := Call(__e, PrimFunc(symreceive), tmp18376)


__e.TailApply(PrimFunc(symshen_4variancy), tmp18375, tmp18377, Z5540, Z5541, Z5542, Z5543)
return


}, 1))
return
}, 1))
return
}, 1))
return
}, 1)

tmp18378 := Call(__e, PrimFunc(symshen_4prolog_1vector))


tmp18379 := Call(__e, tmp18372, tmp18378)


tmp18380 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp18381 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp18380)


tmp18382 := Call(__e, PrimFunc(sym_8v), True, tmp18381)


tmp18383 := Call(__e, tmp18379, tmp18382)


tmp18384 := Call(__e, tmp18383, MakeNumber(0))


tmp18385 := MakeNative(func(__e *ControlFlow) {
__e.Return(True)
return
}, 0)

tmp18386 := Call(__e, tmp18384, tmp18385)


__e.TailApply(tmp18364, tmp18386)
return


}, 1)

tmp18387 := Call(__e, PrimFunc(symshen_4rectify_1type), V5537)


__e.TailApply(tmp18363, tmp18387)
return


}, 2)

tmp18388 := Call(__e, ns2_1set, symdeclare, tmp18362)


_ = tmp18388

tmp18389 := MakeNative(func(__e *ControlFlow) {
V5546 := __e.Get(1)
_ = V5546
V5547 := __e.Get(2)
_ = V5547
V5548 := __e.Get(3)
_ = V5548
V5549 := __e.Get(4)
_ = V5549
V5550 := __e.Get(5)
_ = V5550
V5551 := __e.Get(6)
_ = V5551
tmp18396 := Call(__e, PrimFunc(symshen_4unlocked_2), V5549)


if True == tmp18396 {
tmp18390 := MakeNative(func(__e *ControlFlow) {
W5552 := __e.Get(1)
_ = W5552
tmp18391 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18391

tmp18392 := MakeNative(func(__e *ControlFlow) {
__e.TailApply(PrimFunc(symshen_4variants_2), V5546, W5552, V5547, V5548, V5549, V5550, V5551)
return
}, 0)

tmp18393 := Call(__e, PrimFunc(symshen_4system_1S_1h), V5546, W5552, Nil, V5548, V5549, V5550, tmp18392)


__e.TailApply(PrimFunc(symshen_4gc), V5548, tmp18393)
return


}, 1)

tmp18394 := Call(__e, PrimFunc(symshen_4newpv), V5548)


__e.TailApply(tmp18390, tmp18394)
return


} else {
__e.Return(False)
return
}


}, 6)

tmp18397 := Call(__e, ns2_1set, symshen_4variancy, tmp18389)


_ = tmp18397

tmp18398 := MakeNative(func(__e *ControlFlow) {
V5553 := __e.Get(1)
_ = V5553
V5554 := __e.Get(2)
_ = V5554
V5555 := __e.Get(3)
_ = V5555
V5556 := __e.Get(4)
_ = V5556
V5557 := __e.Get(5)
_ = V5557
V5558 := __e.Get(6)
_ = V5558
V5559 := __e.Get(7)
_ = V5559
tmp18399 := MakeNative(func(__e *ControlFlow) {
W5560 := __e.Get(1)
_ = W5560
tmp18400 := MakeNative(func(__e *ControlFlow) {
W5561 := __e.Get(1)
_ = W5561
tmp18424 := PrimEqual(W5561, False)

if True == tmp18424 {
tmp18401 := MakeNative(func(__e *ControlFlow) {
W5564 := __e.Get(1)
_ = W5564
tmp18418 := PrimEqual(W5564, False)

if True == tmp18418 {
tmp18402 := MakeNative(func(__e *ControlFlow) {
W5565 := __e.Get(1)
_ = W5565
tmp18404 := PrimEqual(W5565, False)

if True == tmp18404 {
__e.TailApply(PrimFunc(symshen_4unlock), V5557, W5560)
return
} else {
__e.Return(W5565)
return
}


}, 1)

tmp18416 := Call(__e, PrimFunc(symshen_4unlocked_2), V5557)


var ifres18405 Obj

if True == tmp18416 {
tmp18406 := MakeNative(func(__e *ControlFlow) {
W5566 := __e.Get(1)
_ = W5566
tmp18407 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18407

tmp18408 := Call(__e, PrimFunc(symshen_4deref), V5553, V5556)


tmp18409 := Call(__e, PrimFunc(symshen_4app), tmp18408, MakeString(" may create errors\n"), symshen_4a)


tmp18410 := PrimStringConcat(MakeString("warning: changing the type of "), tmp18409)

tmp18411 := Call(__e, PrimFunc(symstoutput))


tmp18412 := Call(__e, PrimFunc(sympr), tmp18410, tmp18411)


tmp18413 := Call(__e, PrimFunc(symis), W5566, tmp18412, V5556, V5557, W5560, V5559)


__e.TailApply(PrimFunc(symshen_4gc), V5556, tmp18413)
return


}, 1)

tmp18414 := Call(__e, PrimFunc(symshen_4newpv), V5556)


tmp18415 := Call(__e, tmp18406, tmp18414)


ifres18405 = tmp18415


} else {
ifres18405 = False


}

__e.TailApply(tmp18402, ifres18405)
return


} else {
__e.Return(W5564)
return
}


}, 1)

tmp18422 := Call(__e, PrimFunc(symshen_4unlocked_2), V5557)


var ifres18419 Obj

if True == tmp18422 {
tmp18420 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18420

tmp18421 := Call(__e, PrimFunc(symis_b), V5554, V5555, V5556, V5557, W5560, V5559)


ifres18419 = tmp18421


} else {
ifres18419 = False


}

__e.TailApply(tmp18401, ifres18419)
return


} else {
__e.Return(W5561)
return
}


}, 1)

tmp18436 := Call(__e, PrimFunc(symshen_4unlocked_2), V5557)


var ifres18425 Obj

if True == tmp18436 {
tmp18426 := MakeNative(func(__e *ControlFlow) {
W5562 := __e.Get(1)
_ = W5562
tmp18427 := MakeNative(func(__e *ControlFlow) {
W5563 := __e.Get(1)
_ = W5563
tmp18431 := PrimEqual(W5562, symsymbol)

if True == tmp18431 {
__e.TailApply(PrimFunc(symthaw), W5563)
return
} else {
tmp18429 := Call(__e, PrimFunc(symshen_4pvar_2), W5562)


if True == tmp18429 {
__e.TailApply(PrimFunc(symshen_4bind_b), W5562, symsymbol, V5556, W5563)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp18432 := MakeNative(func(__e *ControlFlow) {
tmp18433 := Call(__e, PrimFunc(symshen_4incinfs))


_ = tmp18433

__e.TailApply(PrimFunc(symshen_4cut), V5556, V5557, W5560, V5559)
return


}, 0)

__e.TailApply(tmp18427, tmp18432)
return


}, 1)

tmp18434 := Call(__e, PrimFunc(symshen_4lazyderef), V5554, V5556)


tmp18435 := Call(__e, tmp18426, tmp18434)


ifres18425 = tmp18435


} else {
ifres18425 = False


}

__e.TailApply(tmp18400, ifres18425)
return


}, 1)

tmp18437 := PrimNumberAdd(V5558, MakeNumber(1))

__e.TailApply(tmp18399, tmp18437)
return


}, 7)

tmp18438 := Call(__e, ns2_1set, symshen_4variants_2, tmp18398)


_ = tmp18438

tmp18439 := MakeNative(func(__e *ControlFlow) {
V5567 := __e.Get(1)
_ = V5567
tmp18440 := MakeNative(func(__e *ControlFlow) {
W5568 := __e.Get(1)
_ = W5568
tmp18441 := MakeNative(func(__e *ControlFlow) {
W5569 := __e.Get(1)
_ = W5569
tmp18442 := MakeNative(func(__e *ControlFlow) {
W5570 := __e.Get(1)
_ = W5570
tmp18443 := MakeNative(func(__e *ControlFlow) {
W5571 := __e.Get(1)
_ = W5571
tmp18444 := MakeNative(func(__e *ControlFlow) {
W5572 := __e.Get(1)
_ = W5572
tmp18445 := MakeNative(func(__e *ControlFlow) {
W5573 := __e.Get(1)
_ = W5573
tmp18446 := Call(__e, PrimFunc(symshen_4rcons__form), V5567)


tmp18447 := PrimCons(W5571, Nil)

tmp18448 := PrimCons(W5570, tmp18447)

tmp18449 := PrimCons(W5569, tmp18448)

tmp18450 := PrimCons(W5568, tmp18449)

tmp18451 := PrimCons(tmp18446, tmp18450)

tmp18452 := PrimCons(W5572, tmp18451)

tmp18453 := PrimCons(symis_b, tmp18452)

tmp18454 := Call(__e, PrimFunc(symshen_4stpart), W5573, tmp18453, W5568)


tmp18455 := PrimCons(tmp18454, Nil)

tmp18456 := PrimCons(W5571, tmp18455)

tmp18457 := PrimCons(symlambda, tmp18456)

tmp18458 := PrimCons(tmp18457, Nil)

tmp18459 := PrimCons(W5570, tmp18458)

tmp18460 := PrimCons(symlambda, tmp18459)

tmp18461 := PrimCons(tmp18460, Nil)

tmp18462 := PrimCons(W5569, tmp18461)

tmp18463 := PrimCons(symlambda, tmp18462)

tmp18464 := PrimCons(tmp18463, Nil)

tmp18465 := PrimCons(W5568, tmp18464)

tmp18466 := PrimCons(symlambda, tmp18465)

tmp18467 := PrimCons(tmp18466, Nil)

tmp18468 := PrimCons(W5572, tmp18467)

__e.Return(PrimCons(symlambda, tmp18468))
return


}, 1)

tmp18469 := Call(__e, PrimFunc(symshen_4extract_1vars), V5567)


__e.TailApply(tmp18445, tmp18469)
return


}, 1)

tmp18470 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp18444, tmp18470)
return


}, 1)

tmp18471 := Call(__e, PrimFunc(symgensym), symC)


__e.TailApply(tmp18443, tmp18471)
return


}, 1)

tmp18472 := Call(__e, PrimFunc(symgensym), symKey)


__e.TailApply(tmp18442, tmp18472)
return


}, 1)

tmp18473 := Call(__e, PrimFunc(symgensym), symL)


__e.TailApply(tmp18441, tmp18473)
return


}, 1)

tmp18474 := Call(__e, PrimFunc(symgensym), symB)


__e.TailApply(tmp18440, tmp18474)
return


}, 1)

tmp18475 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp18439)


_ = tmp18475

tmp18476 := MakeNative(func(__e *ControlFlow) {
V5574 := __e.Get(1)
_ = V5574
__e.Return(V5574)
return
}, 1)

tmp18477 := Call(__e, ns2_1set, symshen_4demod, tmp18476)


_ = tmp18477

tmp18478 := PrimCons(symA, Nil)

tmp18479 := PrimCons(sym_1_1_6, tmp18478)

tmp18480 := Call(__e, PrimFunc(symdeclare), symabort, tmp18479)


_ = tmp18480

tmp18481 := PrimCons(symboolean, Nil)

tmp18482 := PrimCons(sym_1_1_6, tmp18481)

tmp18483 := PrimCons(symA, tmp18482)

tmp18484 := Call(__e, PrimFunc(symdeclare), symabsvector_2, tmp18483)


_ = tmp18484

tmp18485 := PrimCons(symA, Nil)

tmp18486 := PrimCons(symlist, tmp18485)

tmp18487 := PrimCons(symA, Nil)

tmp18488 := PrimCons(symlist, tmp18487)

tmp18489 := PrimCons(tmp18488, Nil)

tmp18490 := PrimCons(sym_1_1_6, tmp18489)

tmp18491 := PrimCons(tmp18486, tmp18490)

tmp18492 := PrimCons(tmp18491, Nil)

tmp18493 := PrimCons(sym_1_1_6, tmp18492)

tmp18494 := PrimCons(symA, tmp18493)

tmp18495 := Call(__e, PrimFunc(symdeclare), symadjoin, tmp18494)


_ = tmp18495

tmp18496 := PrimCons(symboolean, Nil)

tmp18497 := PrimCons(sym_1_1_6, tmp18496)

tmp18498 := PrimCons(symboolean, tmp18497)

tmp18499 := PrimCons(tmp18498, Nil)

tmp18500 := PrimCons(sym_1_1_6, tmp18499)

tmp18501 := PrimCons(symboolean, tmp18500)

tmp18502 := Call(__e, PrimFunc(symdeclare), symand, tmp18501)


_ = tmp18502

tmp18503 := PrimCons(symstring, Nil)

tmp18504 := PrimCons(sym_1_1_6, tmp18503)

tmp18505 := PrimCons(symsymbol, tmp18504)

tmp18506 := PrimCons(tmp18505, Nil)

tmp18507 := PrimCons(sym_1_1_6, tmp18506)

tmp18508 := PrimCons(symstring, tmp18507)

tmp18509 := PrimCons(tmp18508, Nil)

tmp18510 := PrimCons(sym_1_1_6, tmp18509)

tmp18511 := PrimCons(symA, tmp18510)

tmp18512 := Call(__e, PrimFunc(symdeclare), symshen_4app, tmp18511)


_ = tmp18512

tmp18513 := PrimCons(symA, Nil)

tmp18514 := PrimCons(symlist, tmp18513)

tmp18515 := PrimCons(symA, Nil)

tmp18516 := PrimCons(symlist, tmp18515)

tmp18517 := PrimCons(symA, Nil)

tmp18518 := PrimCons(symlist, tmp18517)

tmp18519 := PrimCons(tmp18518, Nil)

tmp18520 := PrimCons(sym_1_1_6, tmp18519)

tmp18521 := PrimCons(tmp18516, tmp18520)

tmp18522 := PrimCons(tmp18521, Nil)

tmp18523 := PrimCons(sym_1_1_6, tmp18522)

tmp18524 := PrimCons(tmp18514, tmp18523)

tmp18525 := Call(__e, PrimFunc(symdeclare), symappend, tmp18524)


_ = tmp18525

tmp18526 := PrimCons(symnumber, Nil)

tmp18527 := PrimCons(sym_1_1_6, tmp18526)

tmp18528 := PrimCons(symA, tmp18527)

tmp18529 := Call(__e, PrimFunc(symdeclare), symarity, tmp18528)


_ = tmp18529

tmp18530 := PrimCons(symA, Nil)

tmp18531 := PrimCons(symlist, tmp18530)

tmp18532 := PrimCons(tmp18531, Nil)

tmp18533 := PrimCons(symlist, tmp18532)

tmp18534 := PrimCons(symA, Nil)

tmp18535 := PrimCons(symlist, tmp18534)

tmp18536 := PrimCons(tmp18535, Nil)

tmp18537 := PrimCons(sym_1_1_6, tmp18536)

tmp18538 := PrimCons(tmp18533, tmp18537)

tmp18539 := PrimCons(tmp18538, Nil)

tmp18540 := PrimCons(sym_1_1_6, tmp18539)

tmp18541 := PrimCons(symA, tmp18540)

tmp18542 := Call(__e, PrimFunc(symdeclare), symassoc, tmp18541)


_ = tmp18542

tmp18543 := PrimCons(symboolean, Nil)

tmp18544 := PrimCons(sym_1_1_6, tmp18543)

tmp18545 := PrimCons(symA, tmp18544)

tmp18546 := Call(__e, PrimFunc(symdeclare), symatom_2, tmp18545)


_ = tmp18546

tmp18547 := PrimCons(symboolean, Nil)

tmp18548 := PrimCons(sym_1_1_6, tmp18547)

tmp18549 := PrimCons(symA, tmp18548)

tmp18550 := Call(__e, PrimFunc(symdeclare), symboolean_2, tmp18549)


_ = tmp18550

tmp18551 := PrimCons(symstring, Nil)

tmp18552 := PrimCons(sym_1_1_6, tmp18551)

tmp18553 := PrimCons(symstring, tmp18552)

tmp18554 := Call(__e, PrimFunc(symdeclare), symbootstrap, tmp18553)


_ = tmp18554

tmp18555 := PrimCons(symboolean, Nil)

tmp18556 := PrimCons(sym_1_1_6, tmp18555)

tmp18557 := PrimCons(symsymbol, tmp18556)

tmp18558 := Call(__e, PrimFunc(symdeclare), symbound_2, tmp18557)


_ = tmp18558

tmp18559 := PrimCons(symA, Nil)

tmp18560 := PrimCons(symlist, tmp18559)

tmp18561 := PrimCons(symboolean, Nil)

tmp18562 := PrimCons(sym_1_1_6, tmp18561)

tmp18563 := PrimCons(tmp18560, tmp18562)

tmp18564 := Call(__e, PrimFunc(symdeclare), symshen_4ccons_2, tmp18563)


_ = tmp18564

tmp18565 := PrimCons(symstring, Nil)

tmp18566 := PrimCons(sym_1_1_6, tmp18565)

tmp18567 := PrimCons(symstring, tmp18566)

tmp18568 := Call(__e, PrimFunc(symdeclare), symcd, tmp18567)


_ = tmp18568

tmp18569 := PrimCons(symA, Nil)

tmp18570 := PrimCons(symstream, tmp18569)

tmp18571 := PrimCons(symB, Nil)

tmp18572 := PrimCons(symlist, tmp18571)

tmp18573 := PrimCons(tmp18572, Nil)

tmp18574 := PrimCons(sym_1_1_6, tmp18573)

tmp18575 := PrimCons(tmp18570, tmp18574)

tmp18576 := Call(__e, PrimFunc(symdeclare), symclose, tmp18575)


_ = tmp18576

tmp18577 := PrimCons(symstring, Nil)

tmp18578 := PrimCons(sym_1_1_6, tmp18577)

tmp18579 := PrimCons(symstring, tmp18578)

tmp18580 := PrimCons(tmp18579, Nil)

tmp18581 := PrimCons(sym_1_1_6, tmp18580)

tmp18582 := PrimCons(symstring, tmp18581)

tmp18583 := Call(__e, PrimFunc(symdeclare), symcn, tmp18582)


_ = tmp18583

tmp18584 := PrimCons(symA, Nil)

tmp18585 := PrimCons(symlist, tmp18584)

tmp18586 := PrimCons(symA, Nil)

tmp18587 := PrimCons(symlist, tmp18586)

tmp18588 := PrimCons(symB, Nil)

tmp18589 := PrimCons(tmp18587, tmp18588)

tmp18590 := PrimCons(symstr, tmp18589)

tmp18591 := PrimCons(tmp18590, Nil)

tmp18592 := PrimCons(sym_1_1_6, tmp18591)

tmp18593 := PrimCons(tmp18585, tmp18592)

tmp18594 := PrimCons(symA, Nil)

tmp18595 := PrimCons(symlist, tmp18594)

tmp18596 := PrimCons(symB, Nil)

tmp18597 := PrimCons(sym_1_1_6, tmp18596)

tmp18598 := PrimCons(tmp18595, tmp18597)

tmp18599 := PrimCons(tmp18598, Nil)

tmp18600 := PrimCons(sym_1_1_6, tmp18599)

tmp18601 := PrimCons(tmp18593, tmp18600)

tmp18602 := Call(__e, PrimFunc(symdeclare), symcompile, tmp18601)


_ = tmp18602

tmp18603 := PrimCons(symboolean, Nil)

tmp18604 := PrimCons(sym_1_1_6, tmp18603)

tmp18605 := PrimCons(symA, tmp18604)

tmp18606 := Call(__e, PrimFunc(symdeclare), symcons_2, tmp18605)


_ = tmp18606

tmp18607 := PrimCons(symsymbol, Nil)

tmp18608 := PrimCons(symlist, tmp18607)

tmp18609 := PrimCons(tmp18608, Nil)

tmp18610 := PrimCons(sym_1_1_6, tmp18609)

tmp18611 := Call(__e, PrimFunc(symdeclare), symdatatypes, tmp18610)


_ = tmp18611

tmp18612 := PrimCons(symsymbol, Nil)

tmp18613 := PrimCons(sym_1_1_6, tmp18612)

tmp18614 := PrimCons(symsymbol, tmp18613)

tmp18615 := Call(__e, PrimFunc(symdeclare), symdestroy, tmp18614)


_ = tmp18615

tmp18616 := PrimCons(symA, Nil)

tmp18617 := PrimCons(symlist, tmp18616)

tmp18618 := PrimCons(symA, Nil)

tmp18619 := PrimCons(symlist, tmp18618)

tmp18620 := PrimCons(symA, Nil)

tmp18621 := PrimCons(symlist, tmp18620)

tmp18622 := PrimCons(tmp18621, Nil)

tmp18623 := PrimCons(sym_1_1_6, tmp18622)

tmp18624 := PrimCons(tmp18619, tmp18623)

tmp18625 := PrimCons(tmp18624, Nil)

tmp18626 := PrimCons(sym_1_1_6, tmp18625)

tmp18627 := PrimCons(tmp18617, tmp18626)

tmp18628 := Call(__e, PrimFunc(symdeclare), symdifference, tmp18627)


_ = tmp18628

tmp18629 := PrimCons(symB, Nil)

tmp18630 := PrimCons(sym_1_1_6, tmp18629)

tmp18631 := PrimCons(symB, tmp18630)

tmp18632 := PrimCons(tmp18631, Nil)

tmp18633 := PrimCons(sym_1_1_6, tmp18632)

tmp18634 := PrimCons(symA, tmp18633)

tmp18635 := Call(__e, PrimFunc(symdeclare), symdo, tmp18634)


_ = tmp18635

tmp18636 := PrimCons(symA, Nil)

tmp18637 := PrimCons(symlist, tmp18636)

tmp18638 := PrimCons(symA, Nil)

tmp18639 := PrimCons(symlist, tmp18638)

tmp18640 := PrimCons(symB, Nil)

tmp18641 := PrimCons(symlist, tmp18640)

tmp18642 := PrimCons(tmp18641, Nil)

tmp18643 := PrimCons(tmp18639, tmp18642)

tmp18644 := PrimCons(symstr, tmp18643)

tmp18645 := PrimCons(tmp18644, Nil)

tmp18646 := PrimCons(sym_1_1_6, tmp18645)

tmp18647 := PrimCons(tmp18637, tmp18646)

tmp18648 := Call(__e, PrimFunc(symdeclare), sym_5e_6, tmp18647)


_ = tmp18648

tmp18649 := PrimCons(symA, Nil)

tmp18650 := PrimCons(symlist, tmp18649)

tmp18651 := PrimCons(symB, Nil)

tmp18652 := PrimCons(symlist, tmp18651)

tmp18653 := PrimCons(symA, Nil)

tmp18654 := PrimCons(symlist, tmp18653)

tmp18655 := PrimCons(tmp18654, Nil)

tmp18656 := PrimCons(tmp18652, tmp18655)

tmp18657 := PrimCons(symstr, tmp18656)

tmp18658 := PrimCons(tmp18657, Nil)

tmp18659 := PrimCons(sym_1_1_6, tmp18658)

tmp18660 := PrimCons(tmp18650, tmp18659)

tmp18661 := Call(__e, PrimFunc(symdeclare), sym_5_b_6, tmp18660)


_ = tmp18661

tmp18662 := PrimCons(symA, Nil)

tmp18663 := PrimCons(symlist, tmp18662)

tmp18664 := PrimCons(symA, Nil)

tmp18665 := PrimCons(symlist, tmp18664)

tmp18666 := PrimCons(symB, Nil)

tmp18667 := PrimCons(symlist, tmp18666)

tmp18668 := PrimCons(tmp18667, Nil)

tmp18669 := PrimCons(tmp18665, tmp18668)

tmp18670 := PrimCons(symstr, tmp18669)

tmp18671 := PrimCons(tmp18670, Nil)

tmp18672 := PrimCons(sym_1_1_6, tmp18671)

tmp18673 := PrimCons(tmp18663, tmp18672)

tmp18674 := Call(__e, PrimFunc(symdeclare), sym_5end_6, tmp18673)


_ = tmp18674

tmp18675 := PrimCons(symA, Nil)

tmp18676 := PrimCons(symlist, tmp18675)

tmp18677 := PrimCons(symB, Nil)

tmp18678 := PrimCons(tmp18676, tmp18677)

tmp18679 := PrimCons(symstr, tmp18678)

tmp18680 := PrimCons(symboolean, Nil)

tmp18681 := PrimCons(sym_1_1_6, tmp18680)

tmp18682 := PrimCons(tmp18679, tmp18681)

tmp18683 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure_2, tmp18682)


_ = tmp18683

tmp18684 := PrimCons(symA, Nil)

tmp18685 := PrimCons(symlist, tmp18684)

tmp18686 := PrimCons(symB, Nil)

tmp18687 := PrimCons(tmp18685, tmp18686)

tmp18688 := PrimCons(symstr, tmp18687)

tmp18689 := PrimCons(tmp18688, Nil)

tmp18690 := PrimCons(sym_1_1_6, tmp18689)

tmp18691 := Call(__e, PrimFunc(symdeclare), symshen_4parse_1failure, tmp18690)


_ = tmp18691

tmp18692 := PrimCons(symA, Nil)

tmp18693 := PrimCons(symlist, tmp18692)

tmp18694 := PrimCons(symB, Nil)

tmp18695 := PrimCons(tmp18693, tmp18694)

tmp18696 := PrimCons(symstr, tmp18695)

tmp18697 := PrimCons(symB, Nil)

tmp18698 := PrimCons(sym_1_1_6, tmp18697)

tmp18699 := PrimCons(tmp18696, tmp18698)

tmp18700 := Call(__e, PrimFunc(symdeclare), symshen_4_5_1out, tmp18699)


_ = tmp18700

tmp18701 := PrimCons(symA, Nil)

tmp18702 := PrimCons(symlist, tmp18701)

tmp18703 := PrimCons(symB, Nil)

tmp18704 := PrimCons(tmp18702, tmp18703)

tmp18705 := PrimCons(symstr, tmp18704)

tmp18706 := PrimCons(symA, Nil)

tmp18707 := PrimCons(symlist, tmp18706)

tmp18708 := PrimCons(tmp18707, Nil)

tmp18709 := PrimCons(sym_1_1_6, tmp18708)

tmp18710 := PrimCons(tmp18705, tmp18709)

tmp18711 := Call(__e, PrimFunc(symdeclare), symshen_4in_1_6, tmp18710)


_ = tmp18711

tmp18712 := PrimCons(symA, Nil)

tmp18713 := PrimCons(symlist, tmp18712)

tmp18714 := PrimCons(symA, Nil)

tmp18715 := PrimCons(symlist, tmp18714)

tmp18716 := PrimCons(symB, Nil)

tmp18717 := PrimCons(tmp18715, tmp18716)

tmp18718 := PrimCons(symstr, tmp18717)

tmp18719 := PrimCons(tmp18718, Nil)

tmp18720 := PrimCons(sym_1_1_6, tmp18719)

tmp18721 := PrimCons(symB, tmp18720)

tmp18722 := PrimCons(tmp18721, Nil)

tmp18723 := PrimCons(sym_1_1_6, tmp18722)

tmp18724 := PrimCons(tmp18713, tmp18723)

tmp18725 := Call(__e, PrimFunc(symdeclare), symshen_4comb, tmp18724)


_ = tmp18725

tmp18726 := PrimCons(symA, Nil)

tmp18727 := PrimCons(symlist, tmp18726)

tmp18728 := PrimCons(symboolean, Nil)

tmp18729 := PrimCons(sym_1_1_6, tmp18728)

tmp18730 := PrimCons(tmp18727, tmp18729)

tmp18731 := PrimCons(tmp18730, Nil)

tmp18732 := PrimCons(sym_1_1_6, tmp18731)

tmp18733 := PrimCons(symA, tmp18732)

tmp18734 := Call(__e, PrimFunc(symdeclare), symelement_2, tmp18733)


_ = tmp18734

tmp18735 := PrimCons(symboolean, Nil)

tmp18736 := PrimCons(sym_1_1_6, tmp18735)

tmp18737 := PrimCons(symA, tmp18736)

tmp18738 := Call(__e, PrimFunc(symdeclare), symempty_2, tmp18737)


_ = tmp18738

tmp18739 := PrimCons(symboolean, Nil)

tmp18740 := PrimCons(sym_1_1_6, tmp18739)

tmp18741 := PrimCons(symsymbol, tmp18740)

tmp18742 := Call(__e, PrimFunc(symdeclare), symenable_1type_1theory, tmp18741)


_ = tmp18742

tmp18743 := PrimCons(symsymbol, Nil)

tmp18744 := PrimCons(symlist, tmp18743)

tmp18745 := PrimCons(tmp18744, Nil)

tmp18746 := PrimCons(sym_1_1_6, tmp18745)

tmp18747 := PrimCons(symsymbol, tmp18746)

tmp18748 := Call(__e, PrimFunc(symdeclare), symexternal, tmp18747)


_ = tmp18748

tmp18749 := PrimCons(symstring, Nil)

tmp18750 := PrimCons(sym_1_1_6, tmp18749)

tmp18751 := PrimCons(symexception, tmp18750)

tmp18752 := Call(__e, PrimFunc(symdeclare), symerror_1to_1string, tmp18751)


_ = tmp18752

tmp18753 := PrimCons(symstring, Nil)

tmp18754 := PrimCons(symlist, tmp18753)

tmp18755 := PrimCons(tmp18754, Nil)

tmp18756 := PrimCons(sym_1_1_6, tmp18755)

tmp18757 := PrimCons(symA, tmp18756)

tmp18758 := Call(__e, PrimFunc(symdeclare), symexplode, tmp18757)


_ = tmp18758

tmp18759 := PrimCons(symsymbol, Nil)

tmp18760 := PrimCons(sym_1_1_6, tmp18759)

tmp18761 := PrimCons(symsymbol, tmp18760)

tmp18762 := Call(__e, PrimFunc(symdeclare), symfactorise, tmp18761)


_ = tmp18762

tmp18763 := PrimCons(symboolean, Nil)

tmp18764 := PrimCons(sym_1_1_6, tmp18763)

tmp18765 := Call(__e, PrimFunc(symdeclare), symfactorise_2, tmp18764)


_ = tmp18765

tmp18766 := PrimCons(symsymbol, Nil)

tmp18767 := PrimCons(sym_1_1_6, tmp18766)

tmp18768 := Call(__e, PrimFunc(symdeclare), symfail, tmp18767)


_ = tmp18768

tmp18769 := PrimCons(symA, Nil)

tmp18770 := PrimCons(sym_1_1_6, tmp18769)

tmp18771 := PrimCons(symA, tmp18770)

tmp18772 := PrimCons(symA, Nil)

tmp18773 := PrimCons(sym_1_1_6, tmp18772)

tmp18774 := PrimCons(symA, tmp18773)

tmp18775 := PrimCons(tmp18774, Nil)

tmp18776 := PrimCons(sym_1_1_6, tmp18775)

tmp18777 := PrimCons(tmp18771, tmp18776)

tmp18778 := Call(__e, PrimFunc(symdeclare), symfix, tmp18777)


_ = tmp18778

tmp18779 := PrimCons(symA, Nil)

tmp18780 := PrimCons(symlazy, tmp18779)

tmp18781 := PrimCons(tmp18780, Nil)

tmp18782 := PrimCons(sym_1_1_6, tmp18781)

tmp18783 := PrimCons(symA, tmp18782)

tmp18784 := Call(__e, PrimFunc(symdeclare), symfreeze, tmp18783)


_ = tmp18784

tmp18785 := PrimCons(symB, Nil)

tmp18786 := PrimCons(sym_d, tmp18785)

tmp18787 := PrimCons(symA, tmp18786)

tmp18788 := PrimCons(symA, Nil)

tmp18789 := PrimCons(sym_1_1_6, tmp18788)

tmp18790 := PrimCons(tmp18787, tmp18789)

tmp18791 := Call(__e, PrimFunc(symdeclare), symfst, tmp18790)


_ = tmp18791

tmp18792 := PrimCons(symsymbol, Nil)

tmp18793 := PrimCons(sym_1_1_6, tmp18792)

tmp18794 := PrimCons(symsymbol, tmp18793)

tmp18795 := Call(__e, PrimFunc(symdeclare), symgensym, tmp18794)


_ = tmp18795

tmp18796 := PrimCons(symA, Nil)

tmp18797 := PrimCons(symlist, tmp18796)

tmp18798 := PrimCons(symboolean, Nil)

tmp18799 := PrimCons(sym_1_1_6, tmp18798)

tmp18800 := PrimCons(symA, tmp18799)

tmp18801 := PrimCons(tmp18800, Nil)

tmp18802 := PrimCons(sym_1_1_6, tmp18801)

tmp18803 := PrimCons(tmp18797, tmp18802)

tmp18804 := Call(__e, PrimFunc(symdeclare), symshen_4hds_a_2, tmp18803)


_ = tmp18804

tmp18805 := PrimCons(symboolean, Nil)

tmp18806 := PrimCons(sym_1_1_6, tmp18805)

tmp18807 := PrimCons(symsymbol, tmp18806)

tmp18808 := Call(__e, PrimFunc(symdeclare), symhush, tmp18807)


_ = tmp18808

tmp18809 := PrimCons(symboolean, Nil)

tmp18810 := PrimCons(sym_1_1_6, tmp18809)

tmp18811 := Call(__e, PrimFunc(symdeclare), symhush_2, tmp18810)


_ = tmp18811

tmp18812 := PrimCons(symA, Nil)

tmp18813 := PrimCons(symvector, tmp18812)

tmp18814 := PrimCons(symA, Nil)

tmp18815 := PrimCons(sym_1_1_6, tmp18814)

tmp18816 := PrimCons(symnumber, tmp18815)

tmp18817 := PrimCons(tmp18816, Nil)

tmp18818 := PrimCons(sym_1_1_6, tmp18817)

tmp18819 := PrimCons(tmp18813, tmp18818)

tmp18820 := Call(__e, PrimFunc(symdeclare), sym_5_1vector, tmp18819)


_ = tmp18820

tmp18821 := PrimCons(symA, Nil)

tmp18822 := PrimCons(symvector, tmp18821)

tmp18823 := PrimCons(symA, Nil)

tmp18824 := PrimCons(symvector, tmp18823)

tmp18825 := PrimCons(tmp18824, Nil)

tmp18826 := PrimCons(sym_1_1_6, tmp18825)

tmp18827 := PrimCons(symA, tmp18826)

tmp18828 := PrimCons(tmp18827, Nil)

tmp18829 := PrimCons(sym_1_1_6, tmp18828)

tmp18830 := PrimCons(symnumber, tmp18829)

tmp18831 := PrimCons(tmp18830, Nil)

tmp18832 := PrimCons(sym_1_1_6, tmp18831)

tmp18833 := PrimCons(tmp18822, tmp18832)

tmp18834 := Call(__e, PrimFunc(symdeclare), symvector_1_6, tmp18833)


_ = tmp18834

tmp18835 := PrimCons(symA, Nil)

tmp18836 := PrimCons(symvector, tmp18835)

tmp18837 := PrimCons(tmp18836, Nil)

tmp18838 := PrimCons(sym_1_1_6, tmp18837)

tmp18839 := PrimCons(symnumber, tmp18838)

tmp18840 := Call(__e, PrimFunc(symdeclare), symvector, tmp18839)


_ = tmp18840

tmp18841 := PrimCons(symnumber, Nil)

tmp18842 := PrimCons(sym_1_1_6, tmp18841)

tmp18843 := PrimCons(symsymbol, tmp18842)

tmp18844 := Call(__e, PrimFunc(symdeclare), symget_1time, tmp18843)


_ = tmp18844

tmp18845 := PrimCons(symnumber, Nil)

tmp18846 := PrimCons(sym_1_1_6, tmp18845)

tmp18847 := PrimCons(symnumber, tmp18846)

tmp18848 := PrimCons(tmp18847, Nil)

tmp18849 := PrimCons(sym_1_1_6, tmp18848)

tmp18850 := PrimCons(symA, tmp18849)

tmp18851 := Call(__e, PrimFunc(symdeclare), symhash, tmp18850)


_ = tmp18851

tmp18852 := PrimCons(symA, Nil)

tmp18853 := PrimCons(symlist, tmp18852)

tmp18854 := PrimCons(symA, Nil)

tmp18855 := PrimCons(sym_1_1_6, tmp18854)

tmp18856 := PrimCons(tmp18853, tmp18855)

tmp18857 := Call(__e, PrimFunc(symdeclare), symhead, tmp18856)


_ = tmp18857

tmp18858 := PrimCons(symA, Nil)

tmp18859 := PrimCons(symvector, tmp18858)

tmp18860 := PrimCons(symA, Nil)

tmp18861 := PrimCons(sym_1_1_6, tmp18860)

tmp18862 := PrimCons(tmp18859, tmp18861)

tmp18863 := Call(__e, PrimFunc(symdeclare), symhdv, tmp18862)


_ = tmp18863

tmp18864 := PrimCons(symstring, Nil)

tmp18865 := PrimCons(sym_1_1_6, tmp18864)

tmp18866 := PrimCons(symstring, tmp18865)

tmp18867 := Call(__e, PrimFunc(symdeclare), symhdstr, tmp18866)


_ = tmp18867

tmp18868 := PrimCons(symA, Nil)

tmp18869 := PrimCons(sym_1_1_6, tmp18868)

tmp18870 := PrimCons(symA, tmp18869)

tmp18871 := PrimCons(tmp18870, Nil)

tmp18872 := PrimCons(sym_1_1_6, tmp18871)

tmp18873 := PrimCons(symA, tmp18872)

tmp18874 := PrimCons(tmp18873, Nil)

tmp18875 := PrimCons(sym_1_1_6, tmp18874)

tmp18876 := PrimCons(symboolean, tmp18875)

tmp18877 := Call(__e, PrimFunc(symdeclare), symif, tmp18876)


_ = tmp18877

tmp18878 := PrimCons(symsymbol, Nil)

tmp18879 := PrimCons(sym_1_1_6, tmp18878)

tmp18880 := PrimCons(symsymbol, tmp18879)

tmp18881 := Call(__e, PrimFunc(symdeclare), symin_1package, tmp18880)


_ = tmp18881

tmp18882 := PrimCons(symstring, Nil)

tmp18883 := PrimCons(sym_1_1_6, tmp18882)

tmp18884 := Call(__e, PrimFunc(symdeclare), symit, tmp18883)


_ = tmp18884

tmp18885 := PrimCons(symstring, Nil)

tmp18886 := PrimCons(sym_1_1_6, tmp18885)

tmp18887 := Call(__e, PrimFunc(symdeclare), symimplementation, tmp18886)


_ = tmp18887

tmp18888 := PrimCons(symsymbol, Nil)

tmp18889 := PrimCons(symlist, tmp18888)

tmp18890 := PrimCons(symsymbol, Nil)

tmp18891 := PrimCons(symlist, tmp18890)

tmp18892 := PrimCons(tmp18891, Nil)

tmp18893 := PrimCons(sym_1_1_6, tmp18892)

tmp18894 := PrimCons(tmp18889, tmp18893)

tmp18895 := Call(__e, PrimFunc(symdeclare), syminclude, tmp18894)


_ = tmp18895

tmp18896 := PrimCons(symsymbol, Nil)

tmp18897 := PrimCons(symlist, tmp18896)

tmp18898 := PrimCons(symsymbol, Nil)

tmp18899 := PrimCons(symlist, tmp18898)

tmp18900 := PrimCons(tmp18899, Nil)

tmp18901 := PrimCons(sym_1_1_6, tmp18900)

tmp18902 := PrimCons(tmp18897, tmp18901)

tmp18903 := Call(__e, PrimFunc(symdeclare), syminclude_1all_1but, tmp18902)


_ = tmp18903

tmp18904 := PrimCons(symsymbol, Nil)

tmp18905 := PrimCons(symlist, tmp18904)

tmp18906 := PrimCons(tmp18905, Nil)

tmp18907 := PrimCons(sym_1_1_6, tmp18906)

tmp18908 := Call(__e, PrimFunc(symdeclare), symincluded, tmp18907)


_ = tmp18908

tmp18909 := PrimCons(symnumber, Nil)

tmp18910 := PrimCons(sym_1_1_6, tmp18909)

tmp18911 := Call(__e, PrimFunc(symdeclare), syminferences, tmp18910)


_ = tmp18911

tmp18912 := PrimCons(symstring, Nil)

tmp18913 := PrimCons(sym_1_1_6, tmp18912)

tmp18914 := PrimCons(symstring, tmp18913)

tmp18915 := PrimCons(tmp18914, Nil)

tmp18916 := PrimCons(sym_1_1_6, tmp18915)

tmp18917 := PrimCons(symA, tmp18916)

tmp18918 := Call(__e, PrimFunc(symdeclare), symshen_4insert, tmp18917)


_ = tmp18918

tmp18919 := PrimCons(symboolean, Nil)

tmp18920 := PrimCons(sym_1_1_6, tmp18919)

tmp18921 := PrimCons(symA, tmp18920)

tmp18922 := Call(__e, PrimFunc(symdeclare), syminteger_2, tmp18921)


_ = tmp18922

tmp18923 := PrimCons(symsymbol, Nil)

tmp18924 := PrimCons(symlist, tmp18923)

tmp18925 := PrimCons(tmp18924, Nil)

tmp18926 := PrimCons(sym_1_1_6, tmp18925)

tmp18927 := PrimCons(symsymbol, tmp18926)

tmp18928 := Call(__e, PrimFunc(symdeclare), syminternal, tmp18927)


_ = tmp18928

tmp18929 := PrimCons(symA, Nil)

tmp18930 := PrimCons(symlist, tmp18929)

tmp18931 := PrimCons(symA, Nil)

tmp18932 := PrimCons(symlist, tmp18931)

tmp18933 := PrimCons(symA, Nil)

tmp18934 := PrimCons(symlist, tmp18933)

tmp18935 := PrimCons(tmp18934, Nil)

tmp18936 := PrimCons(sym_1_1_6, tmp18935)

tmp18937 := PrimCons(tmp18932, tmp18936)

tmp18938 := PrimCons(tmp18937, Nil)

tmp18939 := PrimCons(sym_1_1_6, tmp18938)

tmp18940 := PrimCons(tmp18930, tmp18939)

tmp18941 := Call(__e, PrimFunc(symdeclare), symintersection, tmp18940)


_ = tmp18941

tmp18942 := PrimCons(symstring, Nil)

tmp18943 := PrimCons(sym_1_1_6, tmp18942)

tmp18944 := Call(__e, PrimFunc(symdeclare), symlanguage, tmp18943)


_ = tmp18944

tmp18945 := PrimCons(symA, Nil)

tmp18946 := PrimCons(symlist, tmp18945)

tmp18947 := PrimCons(symnumber, Nil)

tmp18948 := PrimCons(sym_1_1_6, tmp18947)

tmp18949 := PrimCons(tmp18946, tmp18948)

tmp18950 := Call(__e, PrimFunc(symdeclare), symlength, tmp18949)


_ = tmp18950

tmp18951 := PrimCons(symA, Nil)

tmp18952 := PrimCons(symvector, tmp18951)

tmp18953 := PrimCons(symnumber, Nil)

tmp18954 := PrimCons(sym_1_1_6, tmp18953)

tmp18955 := PrimCons(tmp18952, tmp18954)

tmp18956 := Call(__e, PrimFunc(symdeclare), symlimit, tmp18955)


_ = tmp18956

tmp18957 := PrimCons(symin, Nil)

tmp18958 := PrimCons(symstream, tmp18957)

tmp18959 := PrimCons(symunit, Nil)

tmp18960 := PrimCons(symlist, tmp18959)

tmp18961 := PrimCons(tmp18960, Nil)

tmp18962 := PrimCons(sym_1_1_6, tmp18961)

tmp18963 := PrimCons(tmp18958, tmp18962)

tmp18964 := Call(__e, PrimFunc(symdeclare), symlineread, tmp18963)


_ = tmp18964

tmp18965 := PrimCons(symsymbol, Nil)

tmp18966 := PrimCons(sym_1_1_6, tmp18965)

tmp18967 := PrimCons(symstring, tmp18966)

tmp18968 := Call(__e, PrimFunc(symdeclare), symload, tmp18967)


_ = tmp18968

tmp18969 := PrimCons(symB, Nil)

tmp18970 := PrimCons(sym_1_1_6, tmp18969)

tmp18971 := PrimCons(symA, tmp18970)

tmp18972 := PrimCons(symA, Nil)

tmp18973 := PrimCons(symlist, tmp18972)

tmp18974 := PrimCons(symB, Nil)

tmp18975 := PrimCons(symlist, tmp18974)

tmp18976 := PrimCons(tmp18975, Nil)

tmp18977 := PrimCons(sym_1_1_6, tmp18976)

tmp18978 := PrimCons(tmp18973, tmp18977)

tmp18979 := PrimCons(tmp18978, Nil)

tmp18980 := PrimCons(sym_1_1_6, tmp18979)

tmp18981 := PrimCons(tmp18971, tmp18980)

tmp18982 := Call(__e, PrimFunc(symdeclare), symmap, tmp18981)


_ = tmp18982

tmp18983 := PrimCons(symB, Nil)

tmp18984 := PrimCons(symlist, tmp18983)

tmp18985 := PrimCons(tmp18984, Nil)

tmp18986 := PrimCons(sym_1_1_6, tmp18985)

tmp18987 := PrimCons(symA, tmp18986)

tmp18988 := PrimCons(symA, Nil)

tmp18989 := PrimCons(symlist, tmp18988)

tmp18990 := PrimCons(symB, Nil)

tmp18991 := PrimCons(symlist, tmp18990)

tmp18992 := PrimCons(tmp18991, Nil)

tmp18993 := PrimCons(sym_1_1_6, tmp18992)

tmp18994 := PrimCons(tmp18989, tmp18993)

tmp18995 := PrimCons(tmp18994, Nil)

tmp18996 := PrimCons(sym_1_1_6, tmp18995)

tmp18997 := PrimCons(tmp18987, tmp18996)

tmp18998 := Call(__e, PrimFunc(symdeclare), symmapcan, tmp18997)


_ = tmp18998

tmp18999 := PrimCons(symnumber, Nil)

tmp19000 := PrimCons(sym_1_1_6, tmp18999)

tmp19001 := PrimCons(symnumber, tmp19000)

tmp19002 := Call(__e, PrimFunc(symdeclare), symmaxinferences, tmp19001)


_ = tmp19002

tmp19003 := PrimCons(symstring, Nil)

tmp19004 := PrimCons(sym_1_1_6, tmp19003)

tmp19005 := PrimCons(symnumber, tmp19004)

tmp19006 := Call(__e, PrimFunc(symdeclare), symn_1_6string, tmp19005)


_ = tmp19006

tmp19007 := PrimCons(symnumber, Nil)

tmp19008 := PrimCons(sym_1_1_6, tmp19007)

tmp19009 := PrimCons(symnumber, tmp19008)

tmp19010 := Call(__e, PrimFunc(symdeclare), symnl, tmp19009)


_ = tmp19010

tmp19011 := PrimCons(symboolean, Nil)

tmp19012 := PrimCons(sym_1_1_6, tmp19011)

tmp19013 := PrimCons(symboolean, tmp19012)

tmp19014 := Call(__e, PrimFunc(symdeclare), symnot, tmp19013)


_ = tmp19014

tmp19015 := PrimCons(symA, Nil)

tmp19016 := PrimCons(symlist, tmp19015)

tmp19017 := PrimCons(symA, Nil)

tmp19018 := PrimCons(sym_1_1_6, tmp19017)

tmp19019 := PrimCons(tmp19016, tmp19018)

tmp19020 := PrimCons(tmp19019, Nil)

tmp19021 := PrimCons(sym_1_1_6, tmp19020)

tmp19022 := PrimCons(symnumber, tmp19021)

tmp19023 := Call(__e, PrimFunc(symdeclare), symnth, tmp19022)


_ = tmp19023

tmp19024 := PrimCons(symboolean, Nil)

tmp19025 := PrimCons(sym_1_1_6, tmp19024)

tmp19026 := PrimCons(symA, tmp19025)

tmp19027 := Call(__e, PrimFunc(symdeclare), symnumber_2, tmp19026)


_ = tmp19027

tmp19028 := PrimCons(symnumber, Nil)

tmp19029 := PrimCons(sym_1_1_6, tmp19028)

tmp19030 := PrimCons(symB, tmp19029)

tmp19031 := PrimCons(tmp19030, Nil)

tmp19032 := PrimCons(sym_1_1_6, tmp19031)

tmp19033 := PrimCons(symA, tmp19032)

tmp19034 := Call(__e, PrimFunc(symdeclare), symoccurrences, tmp19033)


_ = tmp19034

tmp19035 := PrimCons(symboolean, Nil)

tmp19036 := PrimCons(sym_1_1_6, tmp19035)

tmp19037 := PrimCons(symsymbol, tmp19036)

tmp19038 := Call(__e, PrimFunc(symdeclare), symoccurs_1check, tmp19037)


_ = tmp19038

tmp19039 := PrimCons(symboolean, Nil)

tmp19040 := PrimCons(sym_1_1_6, tmp19039)

tmp19041 := Call(__e, PrimFunc(symdeclare), symoccurs_2, tmp19040)


_ = tmp19041

tmp19042 := PrimCons(symboolean, Nil)

tmp19043 := PrimCons(sym_1_1_6, tmp19042)

tmp19044 := PrimCons(symsymbol, tmp19043)

tmp19045 := Call(__e, PrimFunc(symdeclare), symoptimise, tmp19044)


_ = tmp19045

tmp19046 := PrimCons(symboolean, Nil)

tmp19047 := PrimCons(sym_1_1_6, tmp19046)

tmp19048 := Call(__e, PrimFunc(symdeclare), symoptimise_2, tmp19047)


_ = tmp19048

tmp19049 := PrimCons(symboolean, Nil)

tmp19050 := PrimCons(sym_1_1_6, tmp19049)

tmp19051 := PrimCons(symboolean, tmp19050)

tmp19052 := PrimCons(tmp19051, Nil)

tmp19053 := PrimCons(sym_1_1_6, tmp19052)

tmp19054 := PrimCons(symboolean, tmp19053)

tmp19055 := Call(__e, PrimFunc(symdeclare), symor, tmp19054)


_ = tmp19055

tmp19056 := PrimCons(symstring, Nil)

tmp19057 := PrimCons(sym_1_1_6, tmp19056)

tmp19058 := Call(__e, PrimFunc(symdeclare), symos, tmp19057)


_ = tmp19058

tmp19059 := PrimCons(symboolean, Nil)

tmp19060 := PrimCons(sym_1_1_6, tmp19059)

tmp19061 := PrimCons(symsymbol, tmp19060)

tmp19062 := Call(__e, PrimFunc(symdeclare), sympackage_2, tmp19061)


_ = tmp19062

tmp19063 := PrimCons(symstring, Nil)

tmp19064 := PrimCons(sym_1_1_6, tmp19063)

tmp19065 := Call(__e, PrimFunc(symdeclare), symport, tmp19064)


_ = tmp19065

tmp19066 := PrimCons(symstring, Nil)

tmp19067 := PrimCons(sym_1_1_6, tmp19066)

tmp19068 := Call(__e, PrimFunc(symdeclare), symporters, tmp19067)


_ = tmp19068

tmp19069 := PrimCons(symstring, Nil)

tmp19070 := PrimCons(sym_1_1_6, tmp19069)

tmp19071 := PrimCons(symnumber, tmp19070)

tmp19072 := PrimCons(tmp19071, Nil)

tmp19073 := PrimCons(sym_1_1_6, tmp19072)

tmp19074 := PrimCons(symstring, tmp19073)

tmp19075 := Call(__e, PrimFunc(symdeclare), sympos, tmp19074)


_ = tmp19075

tmp19076 := PrimCons(symout, Nil)

tmp19077 := PrimCons(symstream, tmp19076)

tmp19078 := PrimCons(symstring, Nil)

tmp19079 := PrimCons(sym_1_1_6, tmp19078)

tmp19080 := PrimCons(tmp19077, tmp19079)

tmp19081 := PrimCons(tmp19080, Nil)

tmp19082 := PrimCons(sym_1_1_6, tmp19081)

tmp19083 := PrimCons(symstring, tmp19082)

tmp19084 := Call(__e, PrimFunc(symdeclare), sympr, tmp19083)


_ = tmp19084

tmp19085 := PrimCons(symA, Nil)

tmp19086 := PrimCons(sym_1_1_6, tmp19085)

tmp19087 := PrimCons(symA, tmp19086)

tmp19088 := Call(__e, PrimFunc(symdeclare), symprint, tmp19087)


_ = tmp19088

tmp19089 := PrimCons(symsymbol, Nil)

tmp19090 := PrimCons(sym_1_1_6, tmp19089)

tmp19091 := PrimCons(symsymbol, tmp19090)

tmp19092 := Call(__e, PrimFunc(symdeclare), symprofile, tmp19091)


_ = tmp19092

tmp19093 := PrimCons(symsymbol, Nil)

tmp19094 := PrimCons(symlist, tmp19093)

tmp19095 := PrimCons(symsymbol, Nil)

tmp19096 := PrimCons(symlist, tmp19095)

tmp19097 := PrimCons(tmp19096, Nil)

tmp19098 := PrimCons(sym_1_1_6, tmp19097)

tmp19099 := PrimCons(tmp19094, tmp19098)

tmp19100 := Call(__e, PrimFunc(symdeclare), sympreclude, tmp19099)


_ = tmp19100

tmp19101 := PrimCons(symstring, Nil)

tmp19102 := PrimCons(sym_1_1_6, tmp19101)

tmp19103 := PrimCons(symstring, tmp19102)

tmp19104 := Call(__e, PrimFunc(symdeclare), symshen_4proc_1nl, tmp19103)


_ = tmp19104

tmp19105 := PrimCons(symnumber, Nil)

tmp19106 := PrimCons(sym_d, tmp19105)

tmp19107 := PrimCons(symsymbol, tmp19106)

tmp19108 := PrimCons(tmp19107, Nil)

tmp19109 := PrimCons(sym_1_1_6, tmp19108)

tmp19110 := PrimCons(symsymbol, tmp19109)

tmp19111 := Call(__e, PrimFunc(symdeclare), symprofile_1results, tmp19110)


_ = tmp19111

tmp19112 := PrimCons(symA, Nil)

tmp19113 := PrimCons(sym_1_1_6, tmp19112)

tmp19114 := PrimCons(symA, tmp19113)

tmp19115 := Call(__e, PrimFunc(symdeclare), symprotect, tmp19114)


_ = tmp19115

tmp19116 := PrimCons(symsymbol, Nil)

tmp19117 := PrimCons(symlist, tmp19116)

tmp19118 := PrimCons(symsymbol, Nil)

tmp19119 := PrimCons(symlist, tmp19118)

tmp19120 := PrimCons(tmp19119, Nil)

tmp19121 := PrimCons(sym_1_1_6, tmp19120)

tmp19122 := PrimCons(tmp19117, tmp19121)

tmp19123 := Call(__e, PrimFunc(symdeclare), sympreclude_1all_1but, tmp19122)


_ = tmp19123

tmp19124 := PrimCons(symout, Nil)

tmp19125 := PrimCons(symstream, tmp19124)

tmp19126 := PrimCons(symstring, Nil)

tmp19127 := PrimCons(sym_1_1_6, tmp19126)

tmp19128 := PrimCons(tmp19125, tmp19127)

tmp19129 := PrimCons(tmp19128, Nil)

tmp19130 := PrimCons(sym_1_1_6, tmp19129)

tmp19131 := PrimCons(symstring, tmp19130)

tmp19132 := Call(__e, PrimFunc(symdeclare), symshen_4prhush, tmp19131)


_ = tmp19132

tmp19133 := PrimCons(symnumber, Nil)

tmp19134 := PrimCons(sym_1_1_6, tmp19133)

tmp19135 := PrimCons(symnumber, tmp19134)

tmp19136 := Call(__e, PrimFunc(symdeclare), symprolog_1memory, tmp19135)


_ = tmp19136

tmp19137 := PrimCons(symunit, Nil)

tmp19138 := PrimCons(symlist, tmp19137)

tmp19139 := PrimCons(tmp19138, Nil)

tmp19140 := PrimCons(sym_1_1_6, tmp19139)

tmp19141 := PrimCons(symsymbol, tmp19140)

tmp19142 := Call(__e, PrimFunc(symdeclare), symps, tmp19141)


_ = tmp19142

tmp19143 := PrimCons(symin, Nil)

tmp19144 := PrimCons(symstream, tmp19143)

tmp19145 := PrimCons(symunit, Nil)

tmp19146 := PrimCons(sym_1_1_6, tmp19145)

tmp19147 := PrimCons(tmp19144, tmp19146)

tmp19148 := Call(__e, PrimFunc(symdeclare), symread, tmp19147)


_ = tmp19148

tmp19149 := PrimCons(symin, Nil)

tmp19150 := PrimCons(symstream, tmp19149)

tmp19151 := PrimCons(symnumber, Nil)

tmp19152 := PrimCons(sym_1_1_6, tmp19151)

tmp19153 := PrimCons(tmp19150, tmp19152)

tmp19154 := Call(__e, PrimFunc(symdeclare), symread_1byte, tmp19153)


_ = tmp19154

tmp19155 := PrimCons(symnumber, Nil)

tmp19156 := PrimCons(symlist, tmp19155)

tmp19157 := PrimCons(tmp19156, Nil)

tmp19158 := PrimCons(sym_1_1_6, tmp19157)

tmp19159 := PrimCons(symstring, tmp19158)

tmp19160 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1bytelist, tmp19159)


_ = tmp19160

tmp19161 := PrimCons(symstring, Nil)

tmp19162 := PrimCons(sym_1_1_6, tmp19161)

tmp19163 := PrimCons(symstring, tmp19162)

tmp19164 := Call(__e, PrimFunc(symdeclare), symread_1file_1as_1string, tmp19163)


_ = tmp19164

tmp19165 := PrimCons(symunit, Nil)

tmp19166 := PrimCons(symlist, tmp19165)

tmp19167 := PrimCons(tmp19166, Nil)

tmp19168 := PrimCons(sym_1_1_6, tmp19167)

tmp19169 := PrimCons(symstring, tmp19168)

tmp19170 := Call(__e, PrimFunc(symdeclare), symread_1file, tmp19169)


_ = tmp19170

tmp19171 := PrimCons(symunit, Nil)

tmp19172 := PrimCons(symlist, tmp19171)

tmp19173 := PrimCons(tmp19172, Nil)

tmp19174 := PrimCons(sym_1_1_6, tmp19173)

tmp19175 := PrimCons(symstring, tmp19174)

tmp19176 := Call(__e, PrimFunc(symdeclare), symread_1from_1string, tmp19175)


_ = tmp19176

tmp19177 := PrimCons(symunit, Nil)

tmp19178 := PrimCons(symlist, tmp19177)

tmp19179 := PrimCons(tmp19178, Nil)

tmp19180 := PrimCons(sym_1_1_6, tmp19179)

tmp19181 := PrimCons(symstring, tmp19180)

tmp19182 := Call(__e, PrimFunc(symdeclare), symread_1from_1string_1unprocessed, tmp19181)


_ = tmp19182

tmp19183 := PrimCons(symstring, Nil)

tmp19184 := PrimCons(sym_1_1_6, tmp19183)

tmp19185 := Call(__e, PrimFunc(symdeclare), symrelease, tmp19184)


_ = tmp19185

tmp19186 := PrimCons(symA, Nil)

tmp19187 := PrimCons(symlist, tmp19186)

tmp19188 := PrimCons(symA, Nil)

tmp19189 := PrimCons(symlist, tmp19188)

tmp19190 := PrimCons(tmp19189, Nil)

tmp19191 := PrimCons(sym_1_1_6, tmp19190)

tmp19192 := PrimCons(tmp19187, tmp19191)

tmp19193 := PrimCons(tmp19192, Nil)

tmp19194 := PrimCons(sym_1_1_6, tmp19193)

tmp19195 := PrimCons(symA, tmp19194)

tmp19196 := Call(__e, PrimFunc(symdeclare), symremove, tmp19195)


_ = tmp19196

tmp19197 := PrimCons(symA, Nil)

tmp19198 := PrimCons(symlist, tmp19197)

tmp19199 := PrimCons(symA, Nil)

tmp19200 := PrimCons(symlist, tmp19199)

tmp19201 := PrimCons(tmp19200, Nil)

tmp19202 := PrimCons(sym_1_1_6, tmp19201)

tmp19203 := PrimCons(tmp19198, tmp19202)

tmp19204 := Call(__e, PrimFunc(symdeclare), symreverse, tmp19203)


_ = tmp19204

tmp19205 := PrimCons(symA, Nil)

tmp19206 := PrimCons(sym_1_1_6, tmp19205)

tmp19207 := PrimCons(symstring, tmp19206)

tmp19208 := Call(__e, PrimFunc(symdeclare), symsimple_1error, tmp19207)


_ = tmp19208

tmp19209 := PrimCons(symB, Nil)

tmp19210 := PrimCons(sym_d, tmp19209)

tmp19211 := PrimCons(symA, tmp19210)

tmp19212 := PrimCons(symB, Nil)

tmp19213 := PrimCons(sym_1_1_6, tmp19212)

tmp19214 := PrimCons(tmp19211, tmp19213)

tmp19215 := Call(__e, PrimFunc(symdeclare), symsnd, tmp19214)


_ = tmp19215

tmp19216 := PrimCons(symsymbol, Nil)

tmp19217 := PrimCons(sym_1_1_6, tmp19216)

tmp19218 := PrimCons(symnumber, tmp19217)

tmp19219 := PrimCons(tmp19218, Nil)

tmp19220 := PrimCons(sym_1_1_6, tmp19219)

tmp19221 := PrimCons(symsymbol, tmp19220)

tmp19222 := Call(__e, PrimFunc(symdeclare), symspecialise, tmp19221)


_ = tmp19222

tmp19223 := PrimCons(symboolean, Nil)

tmp19224 := PrimCons(sym_1_1_6, tmp19223)

tmp19225 := PrimCons(symsymbol, tmp19224)

tmp19226 := Call(__e, PrimFunc(symdeclare), symspy, tmp19225)


_ = tmp19226

tmp19227 := PrimCons(symboolean, Nil)

tmp19228 := PrimCons(sym_1_1_6, tmp19227)

tmp19229 := Call(__e, PrimFunc(symdeclare), symspy_2, tmp19228)


_ = tmp19229

tmp19230 := PrimCons(symboolean, Nil)

tmp19231 := PrimCons(sym_1_1_6, tmp19230)

tmp19232 := PrimCons(symsymbol, tmp19231)

tmp19233 := Call(__e, PrimFunc(symdeclare), symstep, tmp19232)


_ = tmp19233

tmp19234 := PrimCons(symboolean, Nil)

tmp19235 := PrimCons(sym_1_1_6, tmp19234)

tmp19236 := Call(__e, PrimFunc(symdeclare), symstep_2, tmp19235)


_ = tmp19236

tmp19237 := PrimCons(symin, Nil)

tmp19238 := PrimCons(symstream, tmp19237)

tmp19239 := PrimCons(tmp19238, Nil)

tmp19240 := PrimCons(sym_1_1_6, tmp19239)

tmp19241 := Call(__e, PrimFunc(symdeclare), symstinput, tmp19240)


_ = tmp19241

tmp19242 := PrimCons(symout, Nil)

tmp19243 := PrimCons(symstream, tmp19242)

tmp19244 := PrimCons(tmp19243, Nil)

tmp19245 := PrimCons(sym_1_1_6, tmp19244)

tmp19246 := Call(__e, PrimFunc(symdeclare), symstoutput, tmp19245)


_ = tmp19246

tmp19247 := PrimCons(symboolean, Nil)

tmp19248 := PrimCons(sym_1_1_6, tmp19247)

tmp19249 := PrimCons(symA, tmp19248)

tmp19250 := Call(__e, PrimFunc(symdeclare), symstring_2, tmp19249)


_ = tmp19250

tmp19251 := PrimCons(symstring, Nil)

tmp19252 := PrimCons(sym_1_1_6, tmp19251)

tmp19253 := PrimCons(symA, tmp19252)

tmp19254 := Call(__e, PrimFunc(symdeclare), symstr, tmp19253)


_ = tmp19254

tmp19255 := PrimCons(symnumber, Nil)

tmp19256 := PrimCons(sym_1_1_6, tmp19255)

tmp19257 := PrimCons(symstring, tmp19256)

tmp19258 := Call(__e, PrimFunc(symdeclare), symstring_1_6n, tmp19257)


_ = tmp19258

tmp19259 := PrimCons(symsymbol, Nil)

tmp19260 := PrimCons(sym_1_1_6, tmp19259)

tmp19261 := PrimCons(symstring, tmp19260)

tmp19262 := Call(__e, PrimFunc(symdeclare), symstring_1_6symbol, tmp19261)


_ = tmp19262

tmp19263 := PrimCons(symnumber, Nil)

tmp19264 := PrimCons(symlist, tmp19263)

tmp19265 := PrimCons(symnumber, Nil)

tmp19266 := PrimCons(sym_1_1_6, tmp19265)

tmp19267 := PrimCons(tmp19264, tmp19266)

tmp19268 := Call(__e, PrimFunc(symdeclare), symsum, tmp19267)


_ = tmp19268

tmp19269 := PrimCons(symboolean, Nil)

tmp19270 := PrimCons(sym_1_1_6, tmp19269)

tmp19271 := PrimCons(symA, tmp19270)

tmp19272 := Call(__e, PrimFunc(symdeclare), symsymbol_2, tmp19271)


_ = tmp19272

tmp19273 := PrimCons(symsymbol, Nil)

tmp19274 := PrimCons(sym_1_1_6, tmp19273)

tmp19275 := PrimCons(symsymbol, tmp19274)

tmp19276 := Call(__e, PrimFunc(symdeclare), symsystemf, tmp19275)


_ = tmp19276

tmp19277 := PrimCons(symboolean, Nil)

tmp19278 := PrimCons(sym_1_1_6, tmp19277)

tmp19279 := Call(__e, PrimFunc(symdeclare), symsystem_1S_2, tmp19278)


_ = tmp19279

tmp19280 := PrimCons(symA, Nil)

tmp19281 := PrimCons(symlist, tmp19280)

tmp19282 := PrimCons(symA, Nil)

tmp19283 := PrimCons(symlist, tmp19282)

tmp19284 := PrimCons(tmp19283, Nil)

tmp19285 := PrimCons(sym_1_1_6, tmp19284)

tmp19286 := PrimCons(tmp19281, tmp19285)

tmp19287 := Call(__e, PrimFunc(symdeclare), symtail, tmp19286)


_ = tmp19287

tmp19288 := PrimCons(symstring, Nil)

tmp19289 := PrimCons(sym_1_1_6, tmp19288)

tmp19290 := PrimCons(symstring, tmp19289)

tmp19291 := Call(__e, PrimFunc(symdeclare), symtlstr, tmp19290)


_ = tmp19291

tmp19292 := PrimCons(symA, Nil)

tmp19293 := PrimCons(symvector, tmp19292)

tmp19294 := PrimCons(symA, Nil)

tmp19295 := PrimCons(symvector, tmp19294)

tmp19296 := PrimCons(tmp19295, Nil)

tmp19297 := PrimCons(sym_1_1_6, tmp19296)

tmp19298 := PrimCons(tmp19293, tmp19297)

tmp19299 := Call(__e, PrimFunc(symdeclare), symtlv, tmp19298)


_ = tmp19299

tmp19300 := PrimCons(symboolean, Nil)

tmp19301 := PrimCons(sym_1_1_6, tmp19300)

tmp19302 := PrimCons(symsymbol, tmp19301)

tmp19303 := Call(__e, PrimFunc(symdeclare), symtc, tmp19302)


_ = tmp19303

tmp19304 := PrimCons(symboolean, Nil)

tmp19305 := PrimCons(sym_1_1_6, tmp19304)

tmp19306 := Call(__e, PrimFunc(symdeclare), symtc_2, tmp19305)


_ = tmp19306

tmp19307 := PrimCons(symA, Nil)

tmp19308 := PrimCons(symlazy, tmp19307)

tmp19309 := PrimCons(symA, Nil)

tmp19310 := PrimCons(sym_1_1_6, tmp19309)

tmp19311 := PrimCons(tmp19308, tmp19310)

tmp19312 := Call(__e, PrimFunc(symdeclare), symthaw, tmp19311)


_ = tmp19312

tmp19313 := PrimCons(symsymbol, Nil)

tmp19314 := PrimCons(sym_1_1_6, tmp19313)

tmp19315 := PrimCons(symsymbol, tmp19314)

tmp19316 := Call(__e, PrimFunc(symdeclare), symtrack, tmp19315)


_ = tmp19316

tmp19317 := PrimCons(symsymbol, Nil)

tmp19318 := PrimCons(symlist, tmp19317)

tmp19319 := PrimCons(tmp19318, Nil)

tmp19320 := PrimCons(sym_1_1_6, tmp19319)

tmp19321 := Call(__e, PrimFunc(symdeclare), symtracked, tmp19320)


_ = tmp19321

tmp19322 := PrimCons(symA, Nil)

tmp19323 := PrimCons(sym_1_1_6, tmp19322)

tmp19324 := PrimCons(symexception, tmp19323)

tmp19325 := PrimCons(symA, Nil)

tmp19326 := PrimCons(sym_1_1_6, tmp19325)

tmp19327 := PrimCons(tmp19324, tmp19326)

tmp19328 := PrimCons(tmp19327, Nil)

tmp19329 := PrimCons(sym_1_1_6, tmp19328)

tmp19330 := PrimCons(symA, tmp19329)

tmp19331 := Call(__e, PrimFunc(symdeclare), symtrap_1error, tmp19330)


_ = tmp19331

tmp19332 := PrimCons(symboolean, Nil)

tmp19333 := PrimCons(sym_1_1_6, tmp19332)

tmp19334 := PrimCons(symA, tmp19333)

tmp19335 := Call(__e, PrimFunc(symdeclare), symtuple_2, tmp19334)


_ = tmp19335

tmp19336 := PrimCons(symsymbol, Nil)

tmp19337 := PrimCons(sym_1_1_6, tmp19336)

tmp19338 := PrimCons(symsymbol, tmp19337)

tmp19339 := Call(__e, PrimFunc(symdeclare), symundefmacro, tmp19338)


_ = tmp19339

tmp19340 := PrimCons(symA, Nil)

tmp19341 := PrimCons(symlist, tmp19340)

tmp19342 := PrimCons(symA, Nil)

tmp19343 := PrimCons(symlist, tmp19342)

tmp19344 := PrimCons(symA, Nil)

tmp19345 := PrimCons(symlist, tmp19344)

tmp19346 := PrimCons(tmp19345, Nil)

tmp19347 := PrimCons(sym_1_1_6, tmp19346)

tmp19348 := PrimCons(tmp19343, tmp19347)

tmp19349 := PrimCons(tmp19348, Nil)

tmp19350 := PrimCons(sym_1_1_6, tmp19349)

tmp19351 := PrimCons(tmp19341, tmp19350)

tmp19352 := Call(__e, PrimFunc(symdeclare), symunion, tmp19351)


_ = tmp19352

tmp19353 := PrimCons(symsymbol, Nil)

tmp19354 := PrimCons(sym_1_1_6, tmp19353)

tmp19355 := PrimCons(symsymbol, tmp19354)

tmp19356 := Call(__e, PrimFunc(symdeclare), symunprofile, tmp19355)


_ = tmp19356

tmp19357 := PrimCons(symsymbol, Nil)

tmp19358 := PrimCons(sym_1_1_6, tmp19357)

tmp19359 := PrimCons(symsymbol, tmp19358)

tmp19360 := Call(__e, PrimFunc(symdeclare), symuntrack, tmp19359)


_ = tmp19360

tmp19361 := PrimCons(symsymbol, Nil)

tmp19362 := PrimCons(symlist, tmp19361)

tmp19363 := PrimCons(tmp19362, Nil)

tmp19364 := PrimCons(sym_1_1_6, tmp19363)

tmp19365 := Call(__e, PrimFunc(symdeclare), symuserdefs, tmp19364)


_ = tmp19365

tmp19366 := PrimCons(symboolean, Nil)

tmp19367 := PrimCons(sym_1_1_6, tmp19366)

tmp19368 := PrimCons(symA, tmp19367)

tmp19369 := Call(__e, PrimFunc(symdeclare), symvariable_2, tmp19368)


_ = tmp19369

tmp19370 := PrimCons(symboolean, Nil)

tmp19371 := PrimCons(sym_1_1_6, tmp19370)

tmp19372 := PrimCons(symA, tmp19371)

tmp19373 := Call(__e, PrimFunc(symdeclare), symvector_2, tmp19372)


_ = tmp19373

tmp19374 := PrimCons(symstring, Nil)

tmp19375 := PrimCons(sym_1_1_6, tmp19374)

tmp19376 := Call(__e, PrimFunc(symdeclare), symversion, tmp19375)


_ = tmp19376

tmp19377 := PrimCons(symA, Nil)

tmp19378 := PrimCons(sym_1_1_6, tmp19377)

tmp19379 := PrimCons(symA, tmp19378)

tmp19380 := PrimCons(tmp19379, Nil)

tmp19381 := PrimCons(sym_1_1_6, tmp19380)

tmp19382 := PrimCons(symstring, tmp19381)

tmp19383 := Call(__e, PrimFunc(symdeclare), symwrite_1to_1file, tmp19382)


_ = tmp19383

tmp19384 := PrimCons(symout, Nil)

tmp19385 := PrimCons(symstream, tmp19384)

tmp19386 := PrimCons(symnumber, Nil)

tmp19387 := PrimCons(sym_1_1_6, tmp19386)

tmp19388 := PrimCons(tmp19385, tmp19387)

tmp19389 := PrimCons(tmp19388, Nil)

tmp19390 := PrimCons(sym_1_1_6, tmp19389)

tmp19391 := PrimCons(symnumber, tmp19390)

tmp19392 := Call(__e, PrimFunc(symdeclare), symwrite_1byte, tmp19391)


_ = tmp19392

tmp19393 := PrimCons(symboolean, Nil)

tmp19394 := PrimCons(sym_1_1_6, tmp19393)

tmp19395 := PrimCons(symstring, tmp19394)

tmp19396 := Call(__e, PrimFunc(symdeclare), symy_1or_1n_2, tmp19395)


_ = tmp19396

tmp19397 := PrimCons(symboolean, Nil)

tmp19398 := PrimCons(sym_1_1_6, tmp19397)

tmp19399 := PrimCons(symnumber, tmp19398)

tmp19400 := PrimCons(tmp19399, Nil)

tmp19401 := PrimCons(sym_1_1_6, tmp19400)

tmp19402 := PrimCons(symnumber, tmp19401)

tmp19403 := Call(__e, PrimFunc(symdeclare), sym_6, tmp19402)


_ = tmp19403

tmp19404 := PrimCons(symboolean, Nil)

tmp19405 := PrimCons(sym_1_1_6, tmp19404)

tmp19406 := PrimCons(symnumber, tmp19405)

tmp19407 := PrimCons(tmp19406, Nil)

tmp19408 := PrimCons(sym_1_1_6, tmp19407)

tmp19409 := PrimCons(symnumber, tmp19408)

tmp19410 := Call(__e, PrimFunc(symdeclare), sym_5, tmp19409)


_ = tmp19410

tmp19411 := PrimCons(symboolean, Nil)

tmp19412 := PrimCons(sym_1_1_6, tmp19411)

tmp19413 := PrimCons(symnumber, tmp19412)

tmp19414 := PrimCons(tmp19413, Nil)

tmp19415 := PrimCons(sym_1_1_6, tmp19414)

tmp19416 := PrimCons(symnumber, tmp19415)

tmp19417 := Call(__e, PrimFunc(symdeclare), sym_6_a, tmp19416)


_ = tmp19417

tmp19418 := PrimCons(symboolean, Nil)

tmp19419 := PrimCons(sym_1_1_6, tmp19418)

tmp19420 := PrimCons(symnumber, tmp19419)

tmp19421 := PrimCons(tmp19420, Nil)

tmp19422 := PrimCons(sym_1_1_6, tmp19421)

tmp19423 := PrimCons(symnumber, tmp19422)

tmp19424 := Call(__e, PrimFunc(symdeclare), sym_5_a, tmp19423)


_ = tmp19424

tmp19425 := PrimCons(symboolean, Nil)

tmp19426 := PrimCons(sym_1_1_6, tmp19425)

tmp19427 := PrimCons(symA, tmp19426)

tmp19428 := PrimCons(tmp19427, Nil)

tmp19429 := PrimCons(sym_1_1_6, tmp19428)

tmp19430 := PrimCons(symA, tmp19429)

tmp19431 := Call(__e, PrimFunc(symdeclare), sym_a, tmp19430)


_ = tmp19431

tmp19432 := PrimCons(symnumber, Nil)

tmp19433 := PrimCons(sym_1_1_6, tmp19432)

tmp19434 := PrimCons(symnumber, tmp19433)

tmp19435 := PrimCons(tmp19434, Nil)

tmp19436 := PrimCons(sym_1_1_6, tmp19435)

tmp19437 := PrimCons(symnumber, tmp19436)

tmp19438 := Call(__e, PrimFunc(symdeclare), sym_7, tmp19437)


_ = tmp19438

tmp19439 := PrimCons(symnumber, Nil)

tmp19440 := PrimCons(sym_1_1_6, tmp19439)

tmp19441 := PrimCons(symnumber, tmp19440)

tmp19442 := PrimCons(tmp19441, Nil)

tmp19443 := PrimCons(sym_1_1_6, tmp19442)

tmp19444 := PrimCons(symnumber, tmp19443)

tmp19445 := Call(__e, PrimFunc(symdeclare), sym_c, tmp19444)


_ = tmp19445

tmp19446 := PrimCons(symnumber, Nil)

tmp19447 := PrimCons(sym_1_1_6, tmp19446)

tmp19448 := PrimCons(symnumber, tmp19447)

tmp19449 := PrimCons(tmp19448, Nil)

tmp19450 := PrimCons(sym_1_1_6, tmp19449)

tmp19451 := PrimCons(symnumber, tmp19450)

tmp19452 := Call(__e, PrimFunc(symdeclare), sym_1, tmp19451)


_ = tmp19452

tmp19453 := PrimCons(symnumber, Nil)

tmp19454 := PrimCons(sym_1_1_6, tmp19453)

tmp19455 := PrimCons(symnumber, tmp19454)

tmp19456 := PrimCons(tmp19455, Nil)

tmp19457 := PrimCons(sym_1_1_6, tmp19456)

tmp19458 := PrimCons(symnumber, tmp19457)

tmp19459 := Call(__e, PrimFunc(symdeclare), sym_d, tmp19458)


_ = tmp19459

tmp19460 := PrimCons(symboolean, Nil)

tmp19461 := PrimCons(sym_1_1_6, tmp19460)

tmp19462 := PrimCons(symB, tmp19461)

tmp19463 := PrimCons(tmp19462, Nil)

tmp19464 := PrimCons(sym_1_1_6, tmp19463)

tmp19465 := PrimCons(symA, tmp19464)

__e.TailApply(PrimFunc(symdeclare), sym_a_a, tmp19465)
return




}, 0)

var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symtlstr = MakeSymbol("tlstr")
var symshen_4_5returns_6 = MakeSymbol("shen.<returns>")
var symforeign = MakeSymbol("foreign")
var symverified = MakeSymbol("verified")
var symshen_4mod = MakeSymbol("shen.mod")
var sym_dhush_d = MakeSymbol("*hush*")
var symmake_1string = MakeSymbol("make-string")
var symshen_4trim_1it = MakeSymbol("shen.trim-it")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symshen_4optimise_1passive = MakeSymbol("shen.optimise-passive")
var syminferences = MakeSymbol("inferences")
var symshen_4_doptimise_d = MakeSymbol("shen.*optimise*")
var symshen_4shendef_1_6kldef = MakeSymbol("shen.shendef->kldef")
var symshen_4nchars = MakeSymbol("shen.nchars")
var symshen_4_5float_6 = MakeSymbol("shen.<float>")
var symstep = MakeSymbol("step")
var symshen_4peek_1history = MakeSymbol("shen.peek-history")
var symshen_4demod = MakeSymbol("shen.demod")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symbar_b = MakeSymbol("bar!")
var symnumber_2 = MakeSymbol("number?")
var symshen_4dynamic_1default = MakeSymbol("shen.dynamic-default")
var symshen_4_7m = MakeSymbol("shen.+m")
var symshen_4prolog_1fbody = MakeSymbol("shen.prolog-fbody")
var symshen_4gc = MakeSymbol("shen.gc")
var symshen_4expand_1mode_1forms = MakeSymbol("shen.expand-mode-forms")
var symshen_4prodbutzero = MakeSymbol("shen.prodbutzero")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4_5head_6 = MakeSymbol("shen.<head>")
var symshen_4ticket_1number = MakeSymbol("shen.ticket-number")
var symshen_4head_1foundit_b = MakeSymbol("shen.head-foundit!")
var sym_dstinput_d = MakeSymbol("*stinput*")
var symshen_4unlocked_2 = MakeSymbol("shen.unlocked?")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4_5rule_d_6 = MakeSymbol("shen.<rule*>")
var symshen_4_5semantics_6 = MakeSymbol("shen.<semantics>")
var symshen_4yacc_1semantics = MakeSymbol("shen.yacc-semantics")
var symshen_4conscode = MakeSymbol("shen.conscode")
var symshen_4posint_2 = MakeSymbol("shen.posint?")
var symboolean_2 = MakeSymbol("boolean?")
var symnth = MakeSymbol("nth")
var sym_dporters_d = MakeSymbol("*porters*")
var symvar_2 = MakeSymbol("var?")
var symshen_4record_1macro = MakeSymbol("shen.record-macro")
var symHd = MakeSymbol("Hd")
var symshen_4char_1stoutput_2 = MakeSymbol("shen.char-stoutput?")
var symFreeze = MakeSymbol("Freeze")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4_dsize_1prolog_1vector_d = MakeSymbol("shen.*size-prolog-vector*")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symhush_2 = MakeSymbol("hush?")
var symshen_4string_1_6byte = MakeSymbol("shen.string->byte")
var sym_dmaximum_1print_1sequence_1size_d = MakeSymbol("*maximum-print-sequence-size*")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4return_2 = MakeSymbol("shen.return?")
var sym_6_6 = MakeSymbol(">>")
var symshen_4_5ass_6 = MakeSymbol("shen.<ass>")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symclose = MakeSymbol("close")
var symshen_4use_1history = MakeSymbol("shen.use-history")
var symshen_4add_1sexpr = MakeSymbol("shen.add-sexpr")
var symshen_4compile_1prolog = MakeSymbol("shen.compile-prolog")
var symshen_4_dalldatatypes_d = MakeSymbol("shen.*alldatatypes*")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4compile_1side_1condition = MakeSymbol("shen.compile-side-condition")
var symshen_4system_1S = MakeSymbol("shen.system-S")
var symshen_4remove_1datatypes = MakeSymbol("shen.remove-datatypes")
var symtuple_2 = MakeSymbol("tuple?")
var symmap = MakeSymbol("map")
var symbound_2 = MakeSymbol("bound?")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symW = MakeSymbol("W")
var symcons_2 = MakeSymbol("cons?")
var symshen_4freshterms = MakeSymbol("shen.freshterms")
var symAction = MakeSymbol("Action")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symn_1_6string = MakeSymbol("n->string")
var symshen_4recursive_1string_1match = MakeSymbol("shen.recursive-string-match")
var symshen_4prolog_1vector = MakeSymbol("shen.prolog-vector")
var symshen_4passive = MakeSymbol("shen.passive")
var symshen_4freshen = MakeSymbol("shen.freshen")
var sympr = MakeSymbol("pr")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var symshen_4process_1sexprs = MakeSymbol("shen.process-sexprs")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var symshen_4compile_1synonyms = MakeSymbol("shen.compile-synonyms")
var symshen_4_5wildcard_6 = MakeSymbol("shen.<wildcard>")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symshen_4modh = MakeSymbol("shen.modh")
var symoptimise = MakeSymbol("optimise")
var symshen_4hush = MakeSymbol("shen.hush")
var symshen_4choicepoint_b = MakeSymbol("shen.choicepoint!")
var symshen_4unpackage_emacroexpand = MakeSymbol("shen.unpackage&macroexpand")
var symtrack = MakeSymbol("track")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symcd = MakeSymbol("cd")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symshen_4prolog_1parameters = MakeSymbol("shen.prolog-parameters")
var symshen_4_5sig_drules_6 = MakeSymbol("shen.<sig*rules>")
var symRemainder = MakeSymbol("Remainder")
var symshen_4abs = MakeSymbol("shen.abs")
var symlambda = MakeSymbol("lambda")
var sym_5_1_1 = MakeSymbol("<--")
var symshen_4evaluate_1lineread = MakeSymbol("shen.evaluate-lineread")
var symshen_4process_1assoc = MakeSymbol("shen.process-assoc")
var symshen_4compile_1head = MakeSymbol("shen.compile-head")
var symshen_4fits_2 = MakeSymbol("shen.fits?")
var symshen_4cons_1case_1plus = MakeSymbol("shen.cons-case-plus")
var symelement_2 = MakeSymbol("element?")
var symshen_4_dsigf_d = MakeSymbol("shen.*sigf*")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symlet = MakeSymbol("let")
var symif = MakeSymbol("if")
var symlineread = MakeSymbol("lineread")
var symshen_4_5lowC_6 = MakeSymbol("shen.<lowC>")
var symshen_4compute_1integer_1h = MakeSymbol("shen.compute-integer-h")
var symstinput = MakeSymbol("stinput")
var symshen_4_duserdefs_d = MakeSymbol("shen.*userdefs*")
var symarity = MakeSymbol("arity")
var symdefun = MakeSymbol("defun")
var symshen_4record_1kl = MakeSymbol("shen.record-kl")
var symshen_4comb = MakeSymbol("shen.comb")
var symY = MakeSymbol("Y")
var symshen_4_5s_1exprs1_6 = MakeSymbol("shen.<s-exprs1>")
var sym_8p = MakeSymbol("@p")
var symread = MakeSymbol("read")
var symshen_4process_1after_1type = MakeSymbol("shen.process-after-type")
var symlazy = MakeSymbol("lazy")
var symshen_4process_1synonyms = MakeSymbol("shen.process-synonyms")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4tabulate_1passive = MakeSymbol("shen.tabulate-passive")
var symshen_4by_1hypothesis = MakeSymbol("shen.by-hypothesis")
var symdifference = MakeSymbol("difference")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var sympreclude = MakeSymbol("preclude")
var symshen_4pac_1h = MakeSymbol("shen.pac-h")
var symshen_4lock = MakeSymbol("shen.lock")
var symRecord = MakeSymbol("Record")
var symshen_4process_1yacc_1semantics = MakeSymbol("shen.process-yacc-semantics")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var sym_dos_d = MakeSymbol("*os*")
var symshen_4free_1variable_2 = MakeSymbol("shen.free-variable?")
var symshen_4read_1unit_1string = MakeSymbol("shen.read-unit-string")
var symshen_4_dspy_d = MakeSymbol("shen.*spy*")
var symprolog_2 = MakeSymbol("prolog?")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4variablecode = MakeSymbol("shen.variablecode")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symstring = MakeSymbol("string")
var symshen_4create_1skeleton = MakeSymbol("shen.create-skeleton")
var symshen_4decrement_1ticket = MakeSymbol("shen.decrement-ticket")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4variants_2 = MakeSymbol("shen.variants?")
var symunit = MakeSymbol("unit")
var symshen_4_5single_6 = MakeSymbol("shen.<single>")
var symshen_4objectcode = MakeSymbol("shen.objectcode")
var symshen_4_5c_1rules_6 = MakeSymbol("shen.<c-rules>")
var sym_1_6 = MakeSymbol("->")
var symspy_2 = MakeSymbol("spy?")
var symshen_4goto_1h = MakeSymbol("shen.goto-h")
var symshen_4unwind = MakeSymbol("shen.unwind")
var symshen_4_5iscolon_6 = MakeSymbol("shen.<iscolon>")
var symshen_4search_1user_1datatypes = MakeSymbol("shen.search-user-datatypes")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4rep_1X = MakeSymbol("shen.rep-X")
var symshen_4constructor_2 = MakeSymbol("shen.constructor?")
var symshen_4read_1file_1as_1bytelist_1help = MakeSymbol("shen.read-file-as-bytelist-help")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4compute_1fraction = MakeSymbol("shen.compute-fraction")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4findall_1h = MakeSymbol("shen.findall-h")
var symoccurs_2 = MakeSymbol("occurs?")
var sym_i = MakeSymbol("{")
var symreceive = MakeSymbol("receive")
var symdefcc = MakeSymbol("defcc")
var symshen_4_5sc_6 = MakeSymbol("shen.<sc>")
var symshen_4bind_b = MakeSymbol("shen.bind!")
var symshen_4compile_1assumptions = MakeSymbol("shen.compile-assumptions")
var symshen_4_5rules_d_6 = MakeSymbol("shen.<rules*>")
var symhash = MakeSymbol("hash")
var symshen_4compute_1fraction_1h = MakeSymbol("shen.compute-fraction-h")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symcons = MakeSymbol("cons")
var symshen_4prolog_1keyword_2 = MakeSymbol("shen.prolog-keyword?")
var symB = MakeSymbol("B")
var symshen_4freeze_1literals = MakeSymbol("shen.freeze-literals")
var symshen_4_5non_1terminal_2_6 = MakeSymbol("shen.<non-terminal?>")
var symKey = MakeSymbol("Key")
var symshen_4bad_1pivot_2 = MakeSymbol("shen.bad-pivot?")
var symshen_4op = MakeSymbol("shen.op")
var sym_5 = MakeSymbol("<")
var symshen_4deref_1calls = MakeSymbol("shen.deref-calls")
var symshen_4monomorphic_2 = MakeSymbol("shen.monomorphic?")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4find_1free_1vars = MakeSymbol("shen.find-free-vars")
var symshen_4mu_1h = MakeSymbol("shen.mu-h")
var symshen_4horn_1clause_1procedure = MakeSymbol("shen.horn-clause-procedure")
var symTm = MakeSymbol("Tm")
var symabort = MakeSymbol("abort")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symget_1time = MakeSymbol("get-time")
var symcall = MakeSymbol("call")
var symmode = MakeSymbol("mode")
var symshen_4combine_1c_1code = MakeSymbol("shen.combine-c-code")
var symshen_4alpha_1convert = MakeSymbol("shen.alpha-convert")
var symshen_4process_1application = MakeSymbol("shen.process-application")
var symintern = MakeSymbol("intern")
var sym_5_1address = MakeSymbol("<-address")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4profiled_2 = MakeSymbol("shen.profiled?")
var symshen_4factor_1selectors = MakeSymbol("shen.factor-selectors")
var symshen_4nothing_1doing_2 = MakeSymbol("shen.nothing-doing?")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4semicolon_2 = MakeSymbol("shen.semicolon?")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4lookupsig = MakeSymbol("shen.lookupsig")
var symshen_4use_1type_1info = MakeSymbol("shen.use-type-info")
var sym_dhome_1directory_d = MakeSymbol("*home-directory*")
var symprint = MakeSymbol("print")
var symshen_4free_1var_1chk = MakeSymbol("shen.free-var-chk")
var symZ = MakeSymbol("Z")
var symshen_4_5e_1number_6 = MakeSymbol("shen.<e-number>")
var sym_1_1_6 = MakeSymbol("-->")
var symshen_4type_1F = MakeSymbol("shen.type-F")
var symshen_4nvars = MakeSymbol("shen.nvars")
var sym_8s = MakeSymbol("@s")
var symshen_4_dfactorise_2_d = MakeSymbol("shen.*factorise?*")
var symassoc = MakeSymbol("assoc")
var symshen_4overbind = MakeSymbol("shen.overbind")
var symshen_4prolog_1track = MakeSymbol("shen.prolog-track")
var symtlv = MakeSymbol("tlv")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4internal_1to_1P_2 = MakeSymbol("shen.internal-to-P?")
var symshen_4initialise_1lambda_1tables = MakeSymbol("shen.initialise-lambda-tables")
var symshen_4_5hterm_6 = MakeSymbol("shen.<hterm>")
var symshen_4occurs_1check_2 = MakeSymbol("shen.occurs-check?")
var symP = MakeSymbol("P")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symconcat = MakeSymbol("concat")
var symlength = MakeSymbol("length")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4undefined_1f_2 = MakeSymbol("shen.undefined-f?")
var symshen_4nextticket = MakeSymbol("shen.nextticket")
var symTl = MakeSymbol("Tl")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var sym_dlanguage_d = MakeSymbol("*language*")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symtl = MakeSymbol("tl")
var symshen_4_dit_d = MakeSymbol("shen.*it*")
var symshen_4macros = MakeSymbol("shen.macros")
var symstep_2 = MakeSymbol("step?")
var sym_b = MakeSymbol("!")
var symvector_1_6 = MakeSymbol("vector->")
var symshen_4triple_1stack = MakeSymbol("shen.triple-stack")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4string_1match = MakeSymbol("shen.string-match")
var symshen_4process_1let = MakeSymbol("shen.process-let")
var symshen_4partial = MakeSymbol("shen.partial")
var symshen_4_5iscomma_6 = MakeSymbol("shen.<iscomma>")
var symshen_4_5yacc_6 = MakeSymbol("shen.<yacc>")
var symshen_4insert = MakeSymbol("shen.insert")
var symdefine = MakeSymbol("define")
var symshen_4factorise_1code = MakeSymbol("shen.factorise-code")
var symshen_4printF = MakeSymbol("shen.printF")
var symshen_4factor_1selectors_1h = MakeSymbol("shen.factor-selectors-h")
var symshen_4process_1datatype = MakeSymbol("shen.process-datatype")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4system_1S_1h = MakeSymbol("shen.system-S-h")
var symempty_2 = MakeSymbol("empty?")
var symshen_4_dgensym_d = MakeSymbol("shen.*gensym*")
var symreverse = MakeSymbol("reverse")
var symlanguage = MakeSymbol("language")
var symtc = MakeSymbol("tc")
var symshen_4compute_1E = MakeSymbol("shen.compute-E")
var symshen_4record_1external = MakeSymbol("shen.record-external")
var symshen_4_dprolog_1memory_d = MakeSymbol("shen.*prolog-memory*")
var sym_4_4_4 = MakeSymbol("...")
var sym_6 = MakeSymbol(">")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symshen_4_1m = MakeSymbol("shen.-m")
var symhdstr = MakeSymbol("hdstr")
var symtail = MakeSymbol("tail")
var symprotect = MakeSymbol("protect")
var symshen_4write_1chars = MakeSymbol("shen.write-chars")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4body_1foundit_b = MakeSymbol("shen.body-foundit!")
var symshen_4l_1rules = MakeSymbol("shen.l-rules")
var symshen_4myassume = MakeSymbol("shen.myassume")
var symshen_4external_1symbols = MakeSymbol("shen.external-symbols")
var syminteger_2 = MakeSymbol("integer?")
var symadjoin = MakeSymbol("adjoin")
var symshen_4non_1application_2 = MakeSymbol("shen.non-application?")
var symtrap_1error = MakeSymbol("trap-error")
var symshen_4cons_1case_1minus = MakeSymbol("shen.cons-case-minus")
var symshen_4decons = MakeSymbol("shen.decons")
var symeval = MakeSymbol("eval")
var symvariable_2 = MakeSymbol("variable?")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4this_1symbol_1is_1unbound = MakeSymbol("shen.this-symbol-is-unbound")
var symoptimise_2 = MakeSymbol("optimise?")
var symshen_4shen_1_6kl_1h = MakeSymbol("shen.shen->kl-h")
var symshen_4rectify_1test = MakeSymbol("shen.rectify-test")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4free_1variable_1error_1message = MakeSymbol("shen.free-variable-error-message")
var symor = MakeSymbol("or")
var symdefprolog = MakeSymbol("defprolog")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4_5bterm_6 = MakeSymbol("shen.<bterm>")
var symshen_4compile_1body = MakeSymbol("shen.compile-body")
var symshen_4passive_2 = MakeSymbol("shen.passive?")
var symshen_4sigxrules = MakeSymbol("shen.sigxrules")
var sym_dmacros_d = MakeSymbol("*macros*")
var symshen_4initialise_1arity_1table = MakeSymbol("shen.initialise-arity-table")
var symshen_4eval_1and_1print = MakeSymbol("shen.eval-and-print")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4update_1assoc = MakeSymbol("shen.update-assoc")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symshen_4_5sng_6 = MakeSymbol("shen.<sng>")
var symshen_4compute_1integer = MakeSymbol("shen.compute-integer")
var symshen_4_ddatatypes_d = MakeSymbol("shen.*datatypes*")
var symshen_4_dsynonyms_d = MakeSymbol("shen.*synonyms*")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symTime = MakeSymbol("Time")
var symshen_4_5conc_6 = MakeSymbol("shen.<conc>")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var sym_dversion_d = MakeSymbol("*version*")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symfile = MakeSymbol("file")
var sym_c_4 = MakeSymbol("/.")
var symshen_4_5datatype_6 = MakeSymbol("shen.<datatype>")
var symshen_4_5hterm2_6 = MakeSymbol("shen.<hterm2>")
var symshen_4_5non_1terminal_1name_6 = MakeSymbol("shen.<non-terminal-name>")
var symshen_4_dtracking_d = MakeSymbol("shen.*tracking*")
var symshen_4factor_1recognisors = MakeSymbol("shen.factor-recognisors")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4_dshen_1type_1theory_1enabled_2_d = MakeSymbol("shen.*shen-type-theory-enabled?*")
var symshen_4in_1_6 = MakeSymbol("shen.in->")
var symhd = MakeSymbol("hd")
var symshen_4lambda_1function = MakeSymbol("shen.lambda-function")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4_5control_6 = MakeSymbol("shen.<control>")
var symshen_4_5fraction_6 = MakeSymbol("shen.<fraction>")
var symsynonyms = MakeSymbol("synonyms")
var symmacroexpand = MakeSymbol("macroexpand")
var sym_5e_6 = MakeSymbol("<e>")
var symGoTo = MakeSymbol("GoTo")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_dnames_d = MakeSymbol("shen.*names*")
var symfunction = MakeSymbol("function")
var symtracked = MakeSymbol("tracked")
var sym_8v = MakeSymbol("@v")
var symshen_4execute_1store_1arity = MakeSymbol("shen.execute-store-arity")
var symshen_4misc_2 = MakeSymbol("shen.misc?")
var symshen_4walk = MakeSymbol("shen.walk")
var symshen_4sng_1h_2 = MakeSymbol("shen.sng-h?")
var symshen_4foundit_b = MakeSymbol("shen.foundit!")
var symshen_4simple_1curry = MakeSymbol("shen.simple-curry")
var symshen_4typename_1h = MakeSymbol("shen.typename-h")
var symoutput = MakeSymbol("output")
var symshen_4terms = MakeSymbol("shen.terms")
var symshen_4demode = MakeSymbol("shen.demode")
var symshen_4compile_1premise = MakeSymbol("shen.compile-premise")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symfork = MakeSymbol("fork")
var symshen_4openlock = MakeSymbol("shen.openlock")
var symshen_4_5side_6 = MakeSymbol("shen.<side>")
var symshen_4source = MakeSymbol("shen.source")
var symnl = MakeSymbol("nl")
var sym_dimplementation_d = MakeSymbol("*implementation*")
var symshen_4hds_a_2 = MakeSymbol("shen.hds=?")
var symshen_4t_d_1correct = MakeSymbol("shen.t*-correct")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4intern_1in_1package = MakeSymbol("shen.intern-in-package")
var symincluded = MakeSymbol("included")
var symshen_4build_1lambda_1table = MakeSymbol("shen.build-lambda-table")
var symshen_4wildcard_2 = MakeSymbol("shen.wildcard?")
var symshen_4invoke = MakeSymbol("shen.invoke")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4unassoc = MakeSymbol("shen.unassoc")
var symSelect = MakeSymbol("Select")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4prhush = MakeSymbol("shen.prhush")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4lambda_1entry = MakeSymbol("shen.lambda-entry")
var symshen_4out_1of_1bounds = MakeSymbol("shen.out-of-bounds")
var symshen_4f_1error = MakeSymbol("shen.f-error")
var symcases = MakeSymbol("cases")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4hascut_2 = MakeSymbol("shen.hascut?")
var symshen_4signal_1def = MakeSymbol("shen.signal-def")
var symfactorise_2 = MakeSymbol("factorise?")
var symshen_4reverse_1help = MakeSymbol("shen.reverse-help")
var symshen_4_5shortnatters_6 = MakeSymbol("shen.<shortnatters>")
var symMessage = MakeSymbol("Message")
var symC = MakeSymbol("C")
var symshen_4_5c_1rule_6 = MakeSymbol("shen.<c-rule>")
var symshen_4terminalcode = MakeSymbol("shen.terminalcode")
var symshen_4pivot_1on = MakeSymbol("shen.pivot-on")
var symit = MakeSymbol("it")
var symshen_4lowercase_2 = MakeSymbol("shen.lowercase?")
var symshen_4shen_1call_2 = MakeSymbol("shen.shen-call?")
var sym_c = MakeSymbol("/")
var symload = MakeSymbol("load")
var symshen_4continue = MakeSymbol("shen.continue")
var symshen_4lchh = MakeSymbol("shen.lchh")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4_dinfs_d = MakeSymbol("shen.*infs*")
var symspecialise = MakeSymbol("specialise")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symopen = MakeSymbol("open")
var symdefmacro = MakeSymbol("defmacro")
var symnumber = MakeSymbol("number")
var symshen_4_5_1out = MakeSymbol("shen.<-out")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symbootstrap = MakeSymbol("bootstrap")
var symshen_4_5body_6 = MakeSymbol("shen.<body>")
var symshen_4lch = MakeSymbol("shen.lch")
var symshen_4deref_1terms = MakeSymbol("shen.deref-terms")
var symshen_4choicepoint_2 = MakeSymbol("shen.choicepoint?")
var symshen_4op1 = MakeSymbol("shen.op1")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4_5integer_6 = MakeSymbol("shen.<integer>")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4package_1symbols = MakeSymbol("shen.package-symbols")
var symshen_4_5hterm1_6 = MakeSymbol("shen.<hterm1>")
var symshen_4_5literal_6 = MakeSymbol("shen.<literal>")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symcn = MakeSymbol("cn")
var symshen_4process_1def = MakeSymbol("shen.process-def")
var symshen_4synonyms_1h = MakeSymbol("shen.synonyms-h")
var symshen_4_5sides_6 = MakeSymbol("shen.<sides>")
var symshen_4optimise_1typing = MakeSymbol("shen.optimise-typing")
var symshen_4prolog_1abstraction = MakeSymbol("shen.prolog-abstraction")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symvector = MakeSymbol("vector")
var symexplode = MakeSymbol("explode")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4top = MakeSymbol("shen.top")
var symshen_4optimise_1passive_1h = MakeSymbol("shen.optimise-passive-h")
var symshen_4variancy = MakeSymbol("shen.variancy")
var symversion = MakeSymbol("version")
var symshen_4shendef_1_6kldef_1h = MakeSymbol("shen.shendef->kldef-h")
var symloaded = MakeSymbol("loaded")
var symshen_4process_1lambda = MakeSymbol("shen.process-lambda")
var symshen_4_5clauses_6 = MakeSymbol("shen.<clauses>")
var symshen_4_5bterms_6 = MakeSymbol("shen.<bterms>")
var symshen_4variable_1case = MakeSymbol("shen.variable-case")
var symshen_4tame = MakeSymbol("shen.tame")
var symfail = MakeSymbol("fail")
var symshen_4_doccurs_d = MakeSymbol("shen.*occurs*")
var symout = MakeSymbol("out")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4zero_1place_2 = MakeSymbol("shen.zero-place?")
var symshen_4_dstep_d = MakeSymbol("shen.*step*")
var symfindall = MakeSymbol("findall")
var symshen_4write_1kl_1h = MakeSymbol("shen.write-kl-h")
var symshen_4_dmaxinferences_d = MakeSymbol("shen.*maxinferences*")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symwrite_1byte = MakeSymbol("write-byte")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4lr_1rule = MakeSymbol("shen.lr-rule")
var symshen_4compile_1premises = MakeSymbol("shen.compile-premises")
var symvector_2 = MakeSymbol("vector?")
var sym_5_1 = MakeSymbol("<-")
var symshen_4kl_1body = MakeSymbol("shen.kl-body")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4update_1lambdatable = MakeSymbol("shen.update-lambdatable")
var symset = MakeSymbol("set")
var symshen_4shen = MakeSymbol("shen.shen")
var symshen_4make_1prolog_1variable = MakeSymbol("shen.make-prolog-variable")
var symshen_4s = MakeSymbol("shen.s")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var symshen_4partial_1application_d_2 = MakeSymbol("shen.partial-application*?")
var sympos = MakeSymbol("pos")
var symshen_4string_1prefix_2 = MakeSymbol("shen.string-prefix?")
var symshen_4unlock = MakeSymbol("shen.unlock")
var symshen_4coll_1formulae = MakeSymbol("shen.coll-formulae")
var symshen_4a = MakeSymbol("shen.a")
var sym_dproperty_1vector_d = MakeSymbol("*property-vector*")
var symshen_4_dextraspecial_d = MakeSymbol("shen.*extraspecial*")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symshen_4freshen_1rule = MakeSymbol("shen.freshen-rule")
var symS = MakeSymbol("S")
var symshen_4_dcall_d = MakeSymbol("shen.*call*")
var symshen_4write_1kl = MakeSymbol("shen.write-kl")
var symshen_4insert_1info = MakeSymbol("shen.insert-info")
var symshen_4prolog_1vector_1size = MakeSymbol("shen.prolog-vector-size")
var symshen_4compile_1consequent = MakeSymbol("shen.compile-consequent")
var symshen_4compile_1side_1conditions = MakeSymbol("shen.compile-side-conditions")
var symshen_4t_d_1rule_1h = MakeSymbol("shen.t*-rule-h")
var symshen_4_5packagename_6 = MakeSymbol("shen.<packagename>")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4process_1applications = MakeSymbol("shen.process-applications")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4_dspecial_d = MakeSymbol("shen.*special*")
var symcompile = MakeSymbol("compile")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4_5s_1exprs2_6 = MakeSymbol("shen.<s-exprs2>")
var symshen_4_dsystem_d = MakeSymbol("shen.*system*")
var symfresh = MakeSymbol("fresh")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symsystemf = MakeSymbol("systemf")
var symshen_4print_1prolog_1vector = MakeSymbol("shen.print-prolog-vector")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4newname = MakeSymbol("shen.newname")
var symshen_4compile_1search_1procedure = MakeSymbol("shen.compile-search-procedure")
var symshen_4c_1rule_1_6shen = MakeSymbol("shen.c-rule->shen")
var symshen_4t = MakeSymbol("shen.t")
var symprofile = MakeSymbol("profile")
var symshen_4non_1terminal_2 = MakeSymbol("shen.non-terminal?")
var symshen_4syntax_1item_2 = MakeSymbol("shen.syntax-item?")
var symshen_4wildcardcode = MakeSymbol("shen.wildcardcode")
var sym_a_a = MakeSymbol("==")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4process_1read_1byte = MakeSymbol("shen.process-read-byte")
var symL = MakeSymbol("L")
var symshen_4extract_1free_1vars = MakeSymbol("shen.extract-free-vars")
var symshen_4non_1terminalcode = MakeSymbol("shen.non-terminalcode")
var symdo = MakeSymbol("do")
var symunion = MakeSymbol("union")
var symshen_4eos = MakeSymbol("shen.eos")
var symshen_4op2 = MakeSymbol("shen.op2")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4loading_2 = MakeSymbol("shen.loading?")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symwhen = MakeSymbol("when")
var symrun = MakeSymbol("run")
var symtime = MakeSymbol("time")
var symshen_4update_1history = MakeSymbol("shen.update-history")
var symshen_4assumetypes = MakeSymbol("shen.assumetypes")
var symshen_4_1null_1 = MakeSymbol("shen.-null-")
var symintersection = MakeSymbol("intersection")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symdeclare = MakeSymbol("declare")
var symboolean = MakeSymbol("boolean")
var symunix = MakeSymbol("unix")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4package_1user_1input = MakeSymbol("shen.package-user-input")
var symshen_4callrec = MakeSymbol("shen.callrec")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4pui_1h = MakeSymbol("shen.pui-h")
var symshen_4dbl_1h_2 = MakeSymbol("shen.dbl-h?")
var symshen_4input_1track = MakeSymbol("shen.input-track")
var symshen_4_dlambdatable_d = MakeSymbol("shen.*lambdatable*")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4rectify_1type = MakeSymbol("shen.rectify-type")
var symuntrack = MakeSymbol("untrack")
var symshen_4prolog_1arity_1check = MakeSymbol("shen.prolog-arity-check")
var symshen_4_5prem_6 = MakeSymbol("shen.<prem>")
var symshen_4toplevel_1forms = MakeSymbol("shen.toplevel-forms")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symundefmacro = MakeSymbol("undefmacro")
var symprofile_1results = MakeSymbol("profile-results")
var sym_5_a = MakeSymbol("<=")
var symis_b = MakeSymbol("is!")
var symdatatype = MakeSymbol("datatype")
var symshen_4compile_1assumption = MakeSymbol("shen.compile-assumption")
var symshen_4cons_1form_1no_1modes = MakeSymbol("shen.cons-form-no-modes")
var symshen_4linearise_1h = MakeSymbol("shen.linearise-h")
var symshen_4_5s_1exprs_6 = MakeSymbol("shen.<s-exprs>")
var symshen_4char_1stinput_2 = MakeSymbol("shen.char-stinput?")
var symshen_4find_1arities = MakeSymbol("shen.find-arities")
var symstr = MakeSymbol("str")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4atom_1case_1minus = MakeSymbol("shen.atom-case-minus")
var symshen_4compile_1premise_1h = MakeSymbol("shen.compile-premise-h")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var syminclude = MakeSymbol("include")
var symshen_4deref_1forked_1literals = MakeSymbol("shen.deref-forked-literals")
var symshen_4primitive = MakeSymbol("shen.primitive")
var symfix = MakeSymbol("fix")
var symsubst = MakeSymbol("subst")
var symfreeze = MakeSymbol("freeze")
var symshen_4unpackage = MakeSymbol("shen.unpackage")
var symshen = MakeSymbol("shen")
var symshen_4fn_1call = MakeSymbol("shen.fn-call")
var symshen_4retract_1clause = MakeSymbol("shen.retract-clause")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4_dtc_d = MakeSymbol("shen.*tc*")
var symoccurrences = MakeSymbol("occurrences")
var symin = MakeSymbol("in")
var symbind = MakeSymbol("bind")
var symassertz = MakeSymbol("assertz")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4ok = MakeSymbol("shen.ok")
var symshen_4freshen_1type = MakeSymbol("shen.freshen-type")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4application_2 = MakeSymbol("shen.application?")
var symprolog_1memory = MakeSymbol("prolog-memory")
var symexception = MakeSymbol("exception")
var symshen_4dynamic = MakeSymbol("shen.dynamic")
var symshen_4vector_1dereference = MakeSymbol("shen.vector-dereference")
var symshen_4c_1rules_1_6shen = MakeSymbol("shen.c-rules->shen")
var symporters = MakeSymbol("porters")
var symimplementation = MakeSymbol("implementation")
var symshen_4fn_1print = MakeSymbol("shen.fn-print")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4_5lowE_6 = MakeSymbol("shen.<lowE>")
var symreturn = MakeSymbol("return")
var symshen_4_5dbl_6 = MakeSymbol("shen.<dbl>")
var symshen_4keep_1looking = MakeSymbol("shen.keep-looking")
var syminternal = MakeSymbol("internal")
var symlimit = MakeSymbol("limit")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4my_1read_1byte = MakeSymbol("shen.my-read-byte")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4_5notdbq_6 = MakeSymbol("shen.<notdbq>")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4goals = MakeSymbol("shen.goals")
var symnull = MakeSymbol("null")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var sym_dstoutput_d = MakeSymbol("*stoutput*")
var symshen_4r = MakeSymbol("shen.r")
var symshen_4op_1test = MakeSymbol("shen.op-test")
var symshen_4_7vector_2 = MakeSymbol("shen.+vector?")
var symsimple_1error = MakeSymbol("simple-error")
var symlist = MakeSymbol("list")
var symshen_4bytes_1_6string = MakeSymbol("shen.bytes->string")
var sympackage = MakeSymbol("package")
var symshen_4foreign_2 = MakeSymbol("shen.foreign?")
var symshen_4lowercase_1symbol_2 = MakeSymbol("shen.lowercase-symbol?")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4rule_1_6clause = MakeSymbol("shen.rule->clause")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4hashkey = MakeSymbol("shen.hashkey")
var symfn = MakeSymbol("fn")
var symshen_4str_1_6bytes = MakeSymbol("shen.str->bytes")
var symstring_2 = MakeSymbol("string?")
var sym_a_a_6 = MakeSymbol("==>")
var symshen_4received = MakeSymbol("shen.received")
var symshen_4show_1datatypes = MakeSymbol("shen.show-datatypes")
var symshen_4f = MakeSymbol("shen.f")
var sym_1 = MakeSymbol("-")
var symos = MakeSymbol("os")
var symshen_4freshterm = MakeSymbol("shen.freshterm")
var symshen_4parse_1failure = MakeSymbol("shen.parse-failure")
var symand = MakeSymbol("and")
var syminput = MakeSymbol("input")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4_dloading_2_d = MakeSymbol("shen.*loading?*")
var symuserdefs = MakeSymbol("userdefs")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4special_1case = MakeSymbol("shen.special-case")
var symshen_4typename = MakeSymbol("shen.typename")
var symshen_4work_1through = MakeSymbol("shen.work-through")
var symshen_4g = MakeSymbol("shen.g")
var symDelta = MakeSymbol("Delta")
var symget = MakeSymbol("get")
var sym_5_1vector = MakeSymbol("<-vector")
var symrelease = MakeSymbol("release")
var sym_3 = MakeSymbol("$")
var symshen_4vector_1parameter = MakeSymbol("shen.vector-parameter")
var symshen_4show = MakeSymbol("shen.show")
var symshen_4prtl = MakeSymbol("shen.prtl")
var symshen_4freshen_1sig = MakeSymbol("shen.freshen-sig")
var symshen_4change_1pointer_1value = MakeSymbol("shen.change-pointer-value")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4fn_1call_2 = MakeSymbol("shen.fn-call?")
var symshen_4check_1eval_1and_1print = MakeSymbol("shen.check-eval-and-print")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4map_1h = MakeSymbol("shen.map-h")
var symeval_1kl = MakeSymbol("eval-kl")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var sym_7 = MakeSymbol("+")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var sym_a = MakeSymbol("=")
var symcond = MakeSymbol("cond")
var symshen_4call_1prolog = MakeSymbol("shen.call-prolog")
var symshen_4_5constructor_6 = MakeSymbol("shen.<constructor>")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symFinish = MakeSymbol("Finish")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4correct = MakeSymbol("shen.correct")
var symshen_4processed = MakeSymbol("shen.processed")
var symshen_4find_1types = MakeSymbol("shen.find-types")
var symshen_4write_1string = MakeSymbol("shen.write-string")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symis = MakeSymbol("is")
var symshen_4predicate = MakeSymbol("shen.predicate")
var sym_e = MakeSymbol("&")
var symshen_4_5packagechar_6 = MakeSymbol("shen.<packagechar>")
var symshen_4ccons_2 = MakeSymbol("shen.ccons?")
var symread_1file = MakeSymbol("read-file")
var symshen_4internal_1to_1shen_2 = MakeSymbol("shen.internal-to-shen?")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4_5colon_1equal_6 = MakeSymbol("shen.<colon-equal>")
var symport = MakeSymbol("port")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4overapplication_2 = MakeSymbol("shen.overapplication?")
var symunprofile = MakeSymbol("unprofile")
var symspy = MakeSymbol("spy")
var symasserta = MakeSymbol("asserta")
var sym_e_e = MakeSymbol("&&")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4freshterm_2 = MakeSymbol("shen.freshterm?")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symStart = MakeSymbol("Start")
var symshen_4stpart = MakeSymbol("shen.stpart")
var symshen_4autocomplete = MakeSymbol("shen.autocomplete")
var symnot = MakeSymbol("not")
var symshen_4scan_1body = MakeSymbol("shen.scan-body")
var symshen_4choicepoint = MakeSymbol("shen.choicepoint")
var symshen_4atom_1case_1plus = MakeSymbol("shen.atom-case-plus")
var symshen_4p_1hyps = MakeSymbol("shen.p-hyps")
var symshen_4_5syntax_6 = MakeSymbol("shen.<syntax>")
var symput = MakeSymbol("put")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symwhere = MakeSymbol("where")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4try_1parse = MakeSymbol("shen.try-parse")
var symshen_4goto = MakeSymbol("shen.goto")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4_dprofiled_d = MakeSymbol("shen.*profiled*")
var symshen_4klfile = MakeSymbol("shen.klfile")
var symshen_4cons_1form_1with_1modes = MakeSymbol("shen.cons-form-with-modes")
var symshen_4terpri_1or_1read_1char = MakeSymbol("shen.terpri-or-read-char")
var symshen_4output_1track = MakeSymbol("shen.output-track")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4whitespace_2 = MakeSymbol("shen.whitespace?")
var symshen_4internal_2 = MakeSymbol("shen.internal?")
var symshen_4key_1in_1sequent_1calculus_2 = MakeSymbol("shen.key-in-sequent-calculus?")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symshen_4_5syntax_1item_6 = MakeSymbol("shen.<syntax-item>")
var symtype = MakeSymbol("type")
var symshen_4print_1freshterm = MakeSymbol("shen.print-freshterm")
var symshen_4rules_1_6prolog = MakeSymbol("shen.rules->prolog")
var symfst = MakeSymbol("fst")
var symshen_4remove_1pointer = MakeSymbol("shen.remove-pointer")
var symsum = MakeSymbol("sum")
var symhead = MakeSymbol("head")
var symshen_4_dresidue_d = MakeSymbol("shen.*residue*")
var symshen_4_5hash_6 = MakeSymbol("shen.<hash>")
var sym_d = MakeSymbol("*")
var symshen_4colon_1equal_2 = MakeSymbol("shen.colon-equal?")
var symsnd = MakeSymbol("snd")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4print_1residue = MakeSymbol("shen.print-residue")
var symshen_4assoc_1_6 = MakeSymbol("shen.assoc->")
var symdatatypes = MakeSymbol("datatypes")
var symatom_2 = MakeSymbol("atom?")
var symabsvector = MakeSymbol("absvector")
var symshen_4process_1cases = MakeSymbol("shen.process-cases")
var symhdv = MakeSymbol("hdv")
var symsymbol = MakeSymbol("symbol")
var symu_b = MakeSymbol("u!")
var symshen_4_5prems_6 = MakeSymbol("shen.<prems>")
var sym_5end_6 = MakeSymbol("<end>")
var symshen_4sng_2 = MakeSymbol("shen.sng?")
var symhush = MakeSymbol("hush")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4call_1dynamic = MakeSymbol("shen.call-dynamic")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symshen_4search = MakeSymbol("shen.search")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symtc_2 = MakeSymbol("tc?")
var sym_j = MakeSymbol("}")
var sympackage_2 = MakeSymbol("package?")
var symshen_4_5shortnatter_6 = MakeSymbol("shen.<shortnatter>")
var symshen_4make_1uppercase = MakeSymbol("shen.make-uppercase")
var symshen_4yacc_1syntax = MakeSymbol("shen.yacc-syntax")
var symshen_4unprotect = MakeSymbol("shen.unprotect")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4_5numeral_6 = MakeSymbol("shen.<numeral>")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4process_1_8s = MakeSymbol("shen.process-@s")
var symshen_4locked_2 = MakeSymbol("shen.locked?")
var symshen_4prterm = MakeSymbol("shen.prterm")
var symshen_4extract_1vars = MakeSymbol("shen.extract-vars")
var symfactorise = MakeSymbol("factorise")
var sym_6_a = MakeSymbol(">=")
var symgensym = MakeSymbol("gensym")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symstoutput = MakeSymbol("stoutput")
var sym_dport_d = MakeSymbol("*port*")
var symsystem_1S_2 = MakeSymbol("system-S?")
var symupdate_1lambda_1table = MakeSymbol("update-lambda-table")
var symshen_4factor = MakeSymbol("shen.factor")
var symA = MakeSymbol("A")
var symremove = MakeSymbol("remove")
var symshen_4compile_1to_1kl = MakeSymbol("shen.compile-to-kl")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symread_1byte = MakeSymbol("read-byte")
var symsave = MakeSymbol("save")
var symshen_4process_1time = MakeSymbol("shen.process-time")
var symX = MakeSymbol("X")
var symshen_4assert_d = MakeSymbol("shen.assert*")
var symunput = MakeSymbol("unput")
var symdestroy = MakeSymbol("destroy")
var symin_1package = MakeSymbol("in-package")
var symshen_4parse_1failure_2 = MakeSymbol("shen.parse-failure?")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4i_1failed_b = MakeSymbol("shen.i-failed!")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4_dhistory_d = MakeSymbol("shen.*history*")
var symps = MakeSymbol("ps")
var symappend = MakeSymbol("append")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symV = MakeSymbol("V")
var symshen_4beta = MakeSymbol("shen.beta")
var symshen_4remove_1indirection = MakeSymbol("shen.remove-indirection")
var symshen_4expt = MakeSymbol("shen.expt")
var symK = MakeSymbol("K")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symshen_4skip = MakeSymbol("shen.skip")
var syminput_7 = MakeSymbol("input+")
var sym__ = MakeSymbol("_")
var symshen_4cons_1form = MakeSymbol("shen.cons-form")
var symstream = MakeSymbol("stream")
var symshen_4bottom = MakeSymbol("shen.bottom")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4arity_1chk = MakeSymbol("shen.arity-chk")
var symshen_4_5longnatter_6 = MakeSymbol("shen.<longnatter>")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symretract = MakeSymbol("retract")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4function_1calls = MakeSymbol("shen.function-calls")
var symexternal = MakeSymbol("external")
var symshen_4record_1and_1evaluate = MakeSymbol("shen.record-and-evaluate")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4dbl_2 = MakeSymbol("shen.dbl?")
var symshen_4type_1error = MakeSymbol("shen.type-error")
var symshen_4cut = MakeSymbol("shen.cut")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symshen_4_5double_6 = MakeSymbol("shen.<double>")
var sym_drelease_d = MakeSymbol("*release*")
var symshen_4_5simple_1pattern_6 = MakeSymbol("shen.<simple-pattern>")
var symvalue = MakeSymbol("value")
var symerror = MakeSymbol("error")
var symParse = MakeSymbol("Parse")
var symthaw = MakeSymbol("thaw")
var symmapcan = MakeSymbol("mapcan")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4unpack_1foreign = MakeSymbol("shen.unpack-foreign")
var symshen_4macroexpand_1h = MakeSymbol("shen.macroexpand-h")
var symshen_4let_b = MakeSymbol("shen.let!")
var symPrevious = MakeSymbol("Previous")
var symshen_4t_d = MakeSymbol("shen.t*")
var symResult = MakeSymbol("Result")
var symread_1from_1string_1unprocessed = MakeSymbol("read-from-string-unprocessed")
var symshen_4find_1arity = MakeSymbol("shen.find-arity")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4_5clause_6 = MakeSymbol("shen.<clause>")
var symshen_4t_d_1integrity = MakeSymbol("shen.t*-integrity")
var symshen_4_5yaccsig_6 = MakeSymbol("shen.<yaccsig>")
var symshen_4_5packagenames_6 = MakeSymbol("shen.<packagenames>")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4_dpackage_d = MakeSymbol("shen.*package*")
