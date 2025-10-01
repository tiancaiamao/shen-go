package main

import . "github.com/tiancaiamao/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp2448 := MakeNative(func(__e *ControlFlow) {
V1447 := __e.Get(1)
_ = V1447
tmp2449 := MakeNative(func(__e *ControlFlow) {
KL := __e.Get(1)
_ = KL
__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), KL)
return
}, 1)

tmp2450 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V1447)


__e.TailApply(tmp2449, tmp2450)
return


}, 1)

tmp2451 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp2448)


_ = tmp2451

tmp2452 := MakeNative(func(__e *ControlFlow) {
V1448 := __e.Get(1)
_ = V1448
tmp2505 := PrimIsPair(V1448)

var ifres2479 Obj

if True == tmp2505 {
tmp2503 := PrimHead(V1448)

tmp2504 := PrimEqual(symdefun, tmp2503)

var ifres2481 Obj

if True == tmp2504 {
tmp2501 := PrimTail(V1448)

tmp2502 := PrimIsPair(tmp2501)

var ifres2483 Obj

if True == tmp2502 {
tmp2498 := PrimTail(V1448)

tmp2499 := PrimTail(tmp2498)

tmp2500 := PrimIsPair(tmp2499)

var ifres2485 Obj

if True == tmp2500 {
tmp2494 := PrimTail(V1448)

tmp2495 := PrimTail(tmp2494)

tmp2496 := PrimTail(tmp2495)

tmp2497 := PrimIsPair(tmp2496)

var ifres2487 Obj

if True == tmp2497 {
tmp2489 := PrimTail(V1448)

tmp2490 := PrimTail(tmp2489)

tmp2491 := PrimTail(tmp2490)

tmp2492 := PrimTail(tmp2491)

tmp2493 := PrimEqual(Nil, tmp2492)

var ifres2488 Obj

if True == tmp2493 {
ifres2488 = True


} else {
ifres2488 = False


}

ifres2487 = ifres2488


} else {
ifres2487 = False


}

var ifres2486 Obj

if True == ifres2487 {
ifres2486 = True


} else {
ifres2486 = False


}

ifres2485 = ifres2486


} else {
ifres2485 = False


}

var ifres2484 Obj

if True == ifres2485 {
ifres2484 = True


} else {
ifres2484 = False


}

ifres2483 = ifres2484


} else {
ifres2483 = False


}

var ifres2482 Obj

if True == ifres2483 {
ifres2482 = True


} else {
ifres2482 = False


}

ifres2481 = ifres2482


} else {
ifres2481 = False


}

var ifres2480 Obj

if True == ifres2481 {
ifres2480 = True


} else {
ifres2480 = False


}

ifres2479 = ifres2480


} else {
ifres2479 = False


}

if True == ifres2479 {
tmp2453 := MakeNative(func(__e *ControlFlow) {
SysfuncChk := __e.Get(1)
_ = SysfuncChk
tmp2454 := MakeNative(func(__e *ControlFlow) {
Arity := __e.Get(1)
_ = Arity
tmp2455 := MakeNative(func(__e *ControlFlow) {
Record := __e.Get(1)
_ = Record
tmp2456 := MakeNative(func(__e *ControlFlow) {
Eval := __e.Get(1)
_ = Eval
tmp2457 := PrimTail(V1448)

tmp2458 := PrimHead(tmp2457)

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp2458)
return


}, 1)

tmp2459 := Call(__e, PrimFunc(symeval_1kl), V1448)


__e.TailApply(tmp2456, tmp2459)
return


}, 1)

tmp2460 := PrimTail(V1448)

tmp2461 := PrimHead(tmp2460)

tmp2462 := Call(__e, PrimFunc(symshen_4record_1kl), tmp2461, V1448)


__e.TailApply(tmp2455, tmp2462)
return


}, 1)

tmp2463 := PrimTail(V1448)

tmp2464 := PrimHead(tmp2463)

tmp2465 := PrimTail(V1448)

tmp2466 := PrimTail(tmp2465)

tmp2467 := PrimHead(tmp2466)

tmp2468 := Call(__e, PrimFunc(symlength), tmp2467)


tmp2469 := Call(__e, PrimFunc(symshen_4store_1arity), tmp2464, tmp2468)


__e.TailApply(tmp2454, tmp2469)
return


}, 1)

tmp2475 := PrimTail(V1448)

tmp2476 := PrimHead(tmp2475)

tmp2477 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp2476)


var ifres2470 Obj

if True == tmp2477 {
tmp2471 := PrimTail(V1448)

tmp2472 := PrimHead(tmp2471)

tmp2473 := Call(__e, PrimFunc(symshen_4app), tmp2472, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2474 := PrimSimpleError(tmp2473)

ifres2470 = tmp2474


} else {
ifres2470 = symshen_4skip


}

__e.TailApply(tmp2453, ifres2470)
return


} else {
__e.Return(V1448)
return
}


}, 1)

tmp2506 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2452)


_ = tmp2506

tmp2507 := MakeNative(func(__e *ControlFlow) {
V1449 := __e.Get(1)
_ = V1449
tmp2608 := PrimIsPair(V1449)

var ifres2600 Obj

if True == tmp2608 {
tmp2606 := PrimHead(V1449)

tmp2607 := PrimEqual(symdefine, tmp2606)

var ifres2602 Obj

if True == tmp2607 {
tmp2604 := PrimTail(V1449)

tmp2605 := PrimIsPair(tmp2604)

var ifres2603 Obj

if True == tmp2605 {
ifres2603 = True


} else {
ifres2603 = False


}

ifres2602 = ifres2603


} else {
ifres2602 = False


}

var ifres2601 Obj

if True == ifres2602 {
ifres2601 = True


} else {
ifres2601 = False


}

ifres2600 = ifres2601


} else {
ifres2600 = False


}

if True == ifres2600 {
tmp2508 := PrimTail(V1449)

tmp2509 := PrimHead(tmp2508)

tmp2510 := PrimTail(V1449)

tmp2511 := PrimTail(tmp2510)

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp2509, tmp2511)
return


} else {
tmp2598 := PrimIsPair(V1449)

var ifres2572 Obj

if True == tmp2598 {
tmp2596 := PrimHead(V1449)

tmp2597 := PrimEqual(symdefun, tmp2596)

var ifres2574 Obj

if True == tmp2597 {
tmp2594 := PrimTail(V1449)

tmp2595 := PrimIsPair(tmp2594)

var ifres2576 Obj

if True == tmp2595 {
tmp2591 := PrimTail(V1449)

tmp2592 := PrimTail(tmp2591)

tmp2593 := PrimIsPair(tmp2592)

var ifres2578 Obj

if True == tmp2593 {
tmp2587 := PrimTail(V1449)

tmp2588 := PrimTail(tmp2587)

tmp2589 := PrimTail(tmp2588)

tmp2590 := PrimIsPair(tmp2589)

var ifres2580 Obj

if True == tmp2590 {
tmp2582 := PrimTail(V1449)

tmp2583 := PrimTail(tmp2582)

tmp2584 := PrimTail(tmp2583)

tmp2585 := PrimTail(tmp2584)

tmp2586 := PrimEqual(Nil, tmp2585)

var ifres2581 Obj

if True == tmp2586 {
ifres2581 = True


} else {
ifres2581 = False


}

ifres2580 = ifres2581


} else {
ifres2580 = False


}

var ifres2579 Obj

if True == ifres2580 {
ifres2579 = True


} else {
ifres2579 = False


}

ifres2578 = ifres2579


} else {
ifres2578 = False


}

var ifres2577 Obj

if True == ifres2578 {
ifres2577 = True


} else {
ifres2577 = False


}

ifres2576 = ifres2577


} else {
ifres2576 = False


}

var ifres2575 Obj

if True == ifres2576 {
ifres2575 = True


} else {
ifres2575 = False


}

ifres2574 = ifres2575


} else {
ifres2574 = False


}

var ifres2573 Obj

if True == ifres2574 {
ifres2573 = True


} else {
ifres2573 = False


}

ifres2572 = ifres2573


} else {
ifres2572 = False


}

if True == ifres2572 {
__e.Return(V1449)
return
} else {
tmp2570 := PrimIsPair(V1449)

var ifres2551 Obj

if True == tmp2570 {
tmp2568 := PrimHead(V1449)

tmp2569 := PrimEqual(symtype, tmp2568)

var ifres2553 Obj

if True == tmp2569 {
tmp2566 := PrimTail(V1449)

tmp2567 := PrimIsPair(tmp2566)

var ifres2555 Obj

if True == tmp2567 {
tmp2563 := PrimTail(V1449)

tmp2564 := PrimTail(tmp2563)

tmp2565 := PrimIsPair(tmp2564)

var ifres2557 Obj

if True == tmp2565 {
tmp2559 := PrimTail(V1449)

tmp2560 := PrimTail(tmp2559)

tmp2561 := PrimTail(tmp2560)

tmp2562 := PrimEqual(Nil, tmp2561)

var ifres2558 Obj

if True == tmp2562 {
ifres2558 = True


} else {
ifres2558 = False


}

ifres2557 = ifres2558


} else {
ifres2557 = False


}

var ifres2556 Obj

if True == ifres2557 {
ifres2556 = True


} else {
ifres2556 = False


}

ifres2555 = ifres2556


} else {
ifres2555 = False


}

var ifres2554 Obj

if True == ifres2555 {
ifres2554 = True


} else {
ifres2554 = False


}

ifres2553 = ifres2554


} else {
ifres2553 = False


}

var ifres2552 Obj

if True == ifres2553 {
ifres2552 = True


} else {
ifres2552 = False


}

ifres2551 = ifres2552


} else {
ifres2551 = False


}

if True == ifres2551 {
tmp2512 := PrimTail(V1449)

tmp2513 := PrimHead(tmp2512)

tmp2514 := PrimTail(V1449)

tmp2515 := PrimTail(tmp2514)

tmp2516 := PrimHead(tmp2515)

tmp2517 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2516)


tmp2518 := PrimCons(tmp2517, Nil)

tmp2519 := PrimCons(tmp2513, tmp2518)

__e.Return(PrimCons(symtype, tmp2519))
return


} else {
tmp2549 := PrimIsPair(V1449)

var ifres2530 Obj

if True == tmp2549 {
tmp2547 := PrimHead(V1449)

tmp2548 := PrimEqual(syminput_7, tmp2547)

var ifres2532 Obj

if True == tmp2548 {
tmp2545 := PrimTail(V1449)

tmp2546 := PrimIsPair(tmp2545)

var ifres2534 Obj

if True == tmp2546 {
tmp2542 := PrimTail(V1449)

tmp2543 := PrimTail(tmp2542)

tmp2544 := PrimIsPair(tmp2543)

var ifres2536 Obj

if True == tmp2544 {
tmp2538 := PrimTail(V1449)

tmp2539 := PrimTail(tmp2538)

tmp2540 := PrimTail(tmp2539)

tmp2541 := PrimEqual(Nil, tmp2540)

var ifres2537 Obj

if True == tmp2541 {
ifres2537 = True


} else {
ifres2537 = False


}

ifres2536 = ifres2537


} else {
ifres2536 = False


}

var ifres2535 Obj

if True == ifres2536 {
ifres2535 = True


} else {
ifres2535 = False


}

ifres2534 = ifres2535


} else {
ifres2534 = False


}

var ifres2533 Obj

if True == ifres2534 {
ifres2533 = True


} else {
ifres2533 = False


}

ifres2532 = ifres2533


} else {
ifres2532 = False


}

var ifres2531 Obj

if True == ifres2532 {
ifres2531 = True


} else {
ifres2531 = False


}

ifres2530 = ifres2531


} else {
ifres2530 = False


}

if True == ifres2530 {
tmp2520 := PrimTail(V1449)

tmp2521 := PrimHead(tmp2520)

tmp2522 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2521)


tmp2523 := PrimTail(V1449)

tmp2524 := PrimTail(tmp2523)

tmp2525 := PrimCons(tmp2522, tmp2524)

__e.Return(PrimCons(syminput_7, tmp2525))
return


} else {
tmp2528 := PrimIsPair(V1449)

if True == tmp2528 {
tmp2526 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2526, V1449)
return


} else {
__e.Return(V1449)
return
}


}


}


}


}


}, 1)

tmp2609 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2507)


_ = tmp2609

tmp2610 := MakeNative(func(__e *ControlFlow) {
V1450 := __e.Get(1)
_ = V1450
V1451 := __e.Get(2)
_ = V1451
tmp2611 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4_5define_6), X)
return
}, 1)

tmp2612 := PrimCons(V1450, V1451)

__e.TailApply(PrimFunc(symcompile), tmp2611, tmp2612)
return


}, 2)

tmp2613 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2610)


_ = tmp2613

tmp2614 := MakeNative(func(__e *ControlFlow) {
V1452 := __e.Get(1)
_ = V1452
tmp2615 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2633 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2633 {
tmp2616 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2618 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp2619 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5name_6 := __e.Get(1)
_ = Parseshen_4_5name_6
tmp2629 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5name_6)


if True == tmp2629 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2620 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rules_6 := __e.Get(1)
_ = Parseshen_4_5rules_6
tmp2626 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rules_6)


if True == tmp2626 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2621 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5rules_6)


tmp2622 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5name_6)


tmp2623 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rules_6)


tmp2624 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), tmp2622, tmp2623)


__e.TailApply(PrimFunc(symshen_4comb), tmp2621, tmp2624)
return


}


}, 1)

tmp2627 := Call(__e, PrimFunc(symshen_4_5rules_6), Parseshen_4_5name_6)


__e.TailApply(tmp2620, tmp2627)
return


}


}, 1)

tmp2630 := Call(__e, PrimFunc(symshen_4_5name_6), V1452)


tmp2631 := Call(__e, tmp2619, tmp2630)


__e.TailApply(tmp2616, tmp2631)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2634 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5name_6 := __e.Get(1)
_ = Parseshen_4_5name_6
tmp2656 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5name_6)


if True == tmp2656 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2654 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5name_6, sym_i)


if True == tmp2654 {
tmp2635 := MakeNative(func(__e *ControlFlow) {
News1313 := __e.Get(1)
_ = News1313
tmp2636 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5signature_6 := __e.Get(1)
_ = Parseshen_4_5signature_6
tmp2650 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5signature_6)


if True == tmp2650 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2648 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5signature_6, sym_j)


if True == tmp2648 {
tmp2637 := MakeNative(func(__e *ControlFlow) {
News1314 := __e.Get(1)
_ = News1314
tmp2638 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rules_6 := __e.Get(1)
_ = Parseshen_4_5rules_6
tmp2644 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rules_6)


if True == tmp2644 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2639 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5rules_6)


tmp2640 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5name_6)


tmp2641 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rules_6)


tmp2642 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), tmp2640, tmp2641)


__e.TailApply(PrimFunc(symshen_4comb), tmp2639, tmp2642)
return


}


}, 1)

tmp2645 := Call(__e, PrimFunc(symshen_4_5rules_6), News1314)


__e.TailApply(tmp2638, tmp2645)
return


}, 1)

tmp2646 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5signature_6)


__e.TailApply(tmp2637, tmp2646)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp2651 := Call(__e, PrimFunc(symshen_4_5signature_6), News1313)


__e.TailApply(tmp2636, tmp2651)
return


}, 1)

tmp2652 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5name_6)


__e.TailApply(tmp2635, tmp2652)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp2657 := Call(__e, PrimFunc(symshen_4_5name_6), V1452)


tmp2658 := Call(__e, tmp2634, tmp2657)


__e.TailApply(tmp2615, tmp2658)
return


}, 1)

tmp2659 := Call(__e, ns2_1set, symshen_4_5define_6, tmp2614)


_ = tmp2659

tmp2660 := MakeNative(func(__e *ControlFlow) {
V1453 := __e.Get(1)
_ = V1453
V1454 := __e.Get(2)
_ = V1454
tmp2661 := MakeNative(func(__e *ControlFlow) {
Ps := __e.Get(1)
_ = Ps
tmp2662 := MakeNative(func(__e *ControlFlow) {
Arity := __e.Get(1)
_ = Arity
tmp2663 := MakeNative(func(__e *ControlFlow) {
FreeVarChk := __e.Get(1)
_ = FreeVarChk
tmp2664 := MakeNative(func(__e *ControlFlow) {
KL := __e.Get(1)
_ = KL
__e.Return(KL)
return
}, 1)

tmp2665 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V1453, V1454, Arity)


tmp2666 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp2665)


__e.TailApply(tmp2664, tmp2666)
return


}, 1)

tmp2667 := MakeNative(func(__e *ControlFlow) {
R := __e.Get(1)
_ = R
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V1453, R)
return
}, 1)

tmp2668 := Call(__e, PrimFunc(symmap), tmp2667, V1454)


__e.TailApply(tmp2663, tmp2668)
return


}, 1)

tmp2669 := Call(__e, PrimFunc(symshen_4arity_1chk), V1453, Ps)


__e.TailApply(tmp2662, tmp2669)
return


}, 1)

tmp2670 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symfst), X)
return
}, 1)

tmp2671 := Call(__e, PrimFunc(symmap), tmp2670, V1454)


__e.TailApply(tmp2661, tmp2671)
return


}, 2)

tmp2672 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2660)


_ = tmp2672

tmp2673 := MakeNative(func(__e *ControlFlow) {
V1455 := __e.Get(1)
_ = V1455
tmp2674 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2676 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2676 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp2693 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1455)


var ifres2677 Obj

if True == tmp2693 {
tmp2678 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp2679 := MakeNative(func(__e *ControlFlow) {
News1316 := __e.Get(1)
_ = News1316
tmp2680 := Call(__e, PrimFunc(symshen_4in_1_6), News1316)


tmp2688 := PrimIsSymbol(X)

var ifres2684 Obj

if True == tmp2688 {
tmp2686 := PrimIsVariable(X)

tmp2687 := PrimNot(tmp2686)

var ifres2685 Obj

if True == tmp2687 {
ifres2685 = True


} else {
ifres2685 = False


}

ifres2684 = ifres2685


} else {
ifres2684 = False


}

var ifres2681 Obj

if True == ifres2684 {
ifres2681 = X


} else {
tmp2682 := Call(__e, PrimFunc(symshen_4app), X, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2683 := PrimSimpleError(tmp2682)

ifres2681 = tmp2683


}

__e.TailApply(PrimFunc(symshen_4comb), tmp2680, ifres2681)
return


}, 1)

tmp2689 := Call(__e, PrimFunc(symshen_4tls), V1455)


__e.TailApply(tmp2679, tmp2689)
return


}, 1)

tmp2690 := Call(__e, PrimFunc(symshen_4hds), V1455)


tmp2691 := Call(__e, tmp2678, tmp2690)


ifres2677 = tmp2691


} else {
tmp2692 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2677 = tmp2692


}

__e.TailApply(tmp2674, ifres2677)
return


}, 1)

tmp2694 := Call(__e, ns2_1set, symshen_4_5name_6, tmp2673)


_ = tmp2694

tmp2695 := MakeNative(func(__e *ControlFlow) {
V1456 := __e.Get(1)
_ = V1456
tmp2696 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2707 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2707 {
tmp2697 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2699 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2699 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp2700 := MakeNative(func(__e *ControlFlow) {
Parse_5e_6 := __e.Get(1)
_ = Parse_5e_6
tmp2703 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)


if True == tmp2703 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2701 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp2701, Nil)
return


}


}, 1)

tmp2704 := Call(__e, PrimFunc(sym_5e_6), V1456)


tmp2705 := Call(__e, tmp2700, tmp2704)


__e.TailApply(tmp2697, tmp2705)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2727 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1456)


var ifres2708 Obj

if True == tmp2727 {
tmp2709 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp2710 := MakeNative(func(__e *ControlFlow) {
News1318 := __e.Get(1)
_ = News1318
tmp2711 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5signature_6 := __e.Get(1)
_ = Parseshen_4_5signature_6
tmp2721 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5signature_6)


if True == tmp2721 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2716 := PrimCons(sym_j, Nil)

tmp2717 := PrimCons(sym_i, tmp2716)

tmp2718 := Call(__e, PrimFunc(symelement_2), X, tmp2717)


tmp2719 := PrimNot(tmp2718)

if True == tmp2719 {
tmp2712 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5signature_6)


tmp2713 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5signature_6)


tmp2714 := PrimCons(X, tmp2713)

__e.TailApply(PrimFunc(symshen_4comb), tmp2712, tmp2714)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp2722 := Call(__e, PrimFunc(symshen_4_5signature_6), News1318)


__e.TailApply(tmp2711, tmp2722)
return


}, 1)

tmp2723 := Call(__e, PrimFunc(symshen_4tls), V1456)


__e.TailApply(tmp2710, tmp2723)
return


}, 1)

tmp2724 := Call(__e, PrimFunc(symshen_4hds), V1456)


tmp2725 := Call(__e, tmp2709, tmp2724)


ifres2708 = tmp2725


} else {
tmp2726 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2708 = tmp2726


}

__e.TailApply(tmp2696, ifres2708)
return


}, 1)

tmp2728 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp2695)


_ = tmp2728

tmp2729 := MakeNative(func(__e *ControlFlow) {
V1457 := __e.Get(1)
_ = V1457
tmp2730 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2748 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2748 {
tmp2731 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2733 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2733 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp2734 := MakeNative(func(__e *ControlFlow) {
Parse_5_b_6 := __e.Get(1)
_ = Parse_5_b_6
tmp2744 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5_b_6)


if True == tmp2744 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2735 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5_b_6)


tmp2741 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)


tmp2742 := Call(__e, PrimFunc(symempty_2), tmp2741)


var ifres2736 Obj

if True == tmp2742 {
ifres2736 = Nil


} else {
tmp2737 := Call(__e, PrimFunc(symshen_4_5_1out), Parse_5_b_6)


tmp2738 := Call(__e, PrimFunc(symshen_4app), tmp2737, MakeString("\n ..."), symshen_4r)


tmp2739 := PrimStringConcat(MakeString("Shen syntax error here:\n "), tmp2738)

tmp2740 := PrimSimpleError(tmp2739)

ifres2736 = tmp2740


}

__e.TailApply(PrimFunc(symshen_4comb), tmp2735, ifres2736)
return


}


}, 1)

tmp2745 := Call(__e, PrimFunc(sym_5_b_6), V1457)


tmp2746 := Call(__e, tmp2734, tmp2745)


__e.TailApply(tmp2731, tmp2746)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2749 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rule_6 := __e.Get(1)
_ = Parseshen_4_5rule_6
tmp2760 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rule_6)


if True == tmp2760 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2750 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5rules_6 := __e.Get(1)
_ = Parseshen_4_5rules_6
tmp2757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rules_6)


if True == tmp2757 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2751 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5rules_6)


tmp2752 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rule_6)


tmp2753 := Call(__e, PrimFunc(symshen_4linearise), tmp2752)


tmp2754 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5rules_6)


tmp2755 := PrimCons(tmp2753, tmp2754)

__e.TailApply(PrimFunc(symshen_4comb), tmp2751, tmp2755)
return


}


}, 1)

tmp2758 := Call(__e, PrimFunc(symshen_4_5rules_6), Parseshen_4_5rule_6)


__e.TailApply(tmp2750, tmp2758)
return


}


}, 1)

tmp2761 := Call(__e, PrimFunc(symshen_4_5rule_6), V1457)


tmp2762 := Call(__e, tmp2749, tmp2761)


__e.TailApply(tmp2730, tmp2762)
return


}, 1)

tmp2763 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp2729)


_ = tmp2763

tmp2764 := MakeNative(func(__e *ControlFlow) {
V1460 := __e.Get(1)
_ = V1460
tmp2769 := Call(__e, PrimFunc(symtuple_2), V1460)


if True == tmp2769 {
tmp2765 := Call(__e, PrimFunc(symfst), V1460)


tmp2766 := Call(__e, PrimFunc(symfst), V1460)


tmp2767 := Call(__e, PrimFunc(symsnd), V1460)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2765, tmp2766, Nil, tmp2767)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise")))
return
}


}, 1)

tmp2770 := Call(__e, ns2_1set, symshen_4linearise, tmp2764)


_ = tmp2770

tmp2771 := MakeNative(func(__e *ControlFlow) {
V1473 := __e.Get(1)
_ = V1473
V1474 := __e.Get(2)
_ = V1474
V1475 := __e.Get(3)
_ = V1475
V1476 := __e.Get(4)
_ = V1476
tmp2810 := PrimEqual(Nil, V1473)

if True == tmp2810 {
__e.TailApply(PrimFunc(sym_8p), V1474, V1476)
return
} else {
tmp2808 := PrimIsPair(V1473)

var ifres2804 Obj

if True == tmp2808 {
tmp2806 := PrimHead(V1473)

tmp2807 := PrimIsPair(tmp2806)

var ifres2805 Obj

if True == tmp2807 {
ifres2805 = True


} else {
ifres2805 = False


}

ifres2804 = ifres2805


} else {
ifres2804 = False


}

if True == ifres2804 {
tmp2772 := PrimHead(V1473)

tmp2773 := PrimTail(V1473)

tmp2774 := Call(__e, PrimFunc(symappend), tmp2772, tmp2773)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2774, V1474, V1475, V1476)
return


} else {
tmp2802 := PrimIsPair(V1473)

var ifres2798 Obj

if True == tmp2802 {
tmp2800 := PrimHead(V1473)

tmp2801 := PrimIsVariable(tmp2800)

var ifres2799 Obj

if True == tmp2801 {
ifres2799 = True


} else {
ifres2799 = False


}

ifres2798 = ifres2799


} else {
ifres2798 = False


}

if True == ifres2798 {
tmp2792 := PrimHead(V1473)

tmp2793 := Call(__e, PrimFunc(symelement_2), tmp2792, V1475)


if True == tmp2793 {
tmp2775 := MakeNative(func(__e *ControlFlow) {
Z := __e.Get(1)
_ = Z
tmp2776 := PrimTail(V1473)

tmp2777 := PrimHead(V1473)

tmp2778 := Call(__e, PrimFunc(symshen_4rep_1X), tmp2777, Z, V1474)


tmp2779 := PrimHead(V1473)

tmp2780 := PrimCons(tmp2779, Nil)

tmp2781 := PrimCons(Z, tmp2780)

tmp2782 := PrimCons(sym_a, tmp2781)

tmp2783 := PrimCons(V1476, Nil)

tmp2784 := PrimCons(tmp2782, tmp2783)

tmp2785 := PrimCons(symwhere, tmp2784)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2776, tmp2778, V1475, tmp2785)
return


}, 1)

tmp2786 := Call(__e, PrimFunc(symprotect), symV)


tmp2787 := Call(__e, PrimFunc(symgensym), tmp2786)


__e.TailApply(tmp2775, tmp2787)
return


} else {
tmp2788 := PrimTail(V1473)

tmp2789 := PrimHead(V1473)

tmp2790 := PrimCons(tmp2789, V1475)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2788, V1474, tmp2790, V1476)
return


}


} else {
tmp2796 := PrimIsPair(V1473)

if True == tmp2796 {
tmp2794 := PrimTail(V1473)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2794, V1474, V1475, V1476)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise-h")))
return
}


}


}


}


}, 4)

tmp2811 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp2771)


_ = tmp2811

tmp2812 := MakeNative(func(__e *ControlFlow) {
V1477 := __e.Get(1)
_ = V1477
tmp2813 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2895 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2895 {
tmp2814 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2875 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2875 {
tmp2815 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2840 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2840 {
tmp2816 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2818 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2818 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp2819 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp2836 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp2836 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2834 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)


if True == tmp2834 {
tmp2820 := MakeNative(func(__e *ControlFlow) {
News1331 := __e.Get(1)
_ = News1331
tmp2831 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1331)


if True == tmp2831 {
tmp2821 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp2822 := MakeNative(func(__e *ControlFlow) {
News1332 := __e.Get(1)
_ = News1332
tmp2823 := Call(__e, PrimFunc(symshen_4in_1_6), News1332)


tmp2824 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp2825 := PrimCons(Action, Nil)

tmp2826 := PrimCons(symshen_4choicepoint_b, tmp2825)

tmp2827 := Call(__e, PrimFunc(sym_8p), tmp2824, tmp2826)


__e.TailApply(PrimFunc(symshen_4comb), tmp2823, tmp2827)
return


}, 1)

tmp2828 := Call(__e, PrimFunc(symshen_4tls), News1331)


__e.TailApply(tmp2822, tmp2828)
return


}, 1)

tmp2829 := Call(__e, PrimFunc(symshen_4hds), News1331)


__e.TailApply(tmp2821, tmp2829)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2832 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp2820, tmp2832)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp2837 := Call(__e, PrimFunc(symshen_4_5patterns_6), V1477)


tmp2838 := Call(__e, tmp2819, tmp2837)


__e.TailApply(tmp2816, tmp2838)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2841 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp2871 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp2871 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2869 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_5_1)


if True == tmp2869 {
tmp2842 := MakeNative(func(__e *ControlFlow) {
News1327 := __e.Get(1)
_ = News1327
tmp2866 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1327)


if True == tmp2866 {
tmp2843 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp2844 := MakeNative(func(__e *ControlFlow) {
News1328 := __e.Get(1)
_ = News1328
tmp2862 := Call(__e, PrimFunc(symshen_4_ahd_2), News1328, symwhere)


if True == tmp2862 {
tmp2845 := MakeNative(func(__e *ControlFlow) {
News1329 := __e.Get(1)
_ = News1329
tmp2859 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1329)


if True == tmp2859 {
tmp2846 := MakeNative(func(__e *ControlFlow) {
Guard := __e.Get(1)
_ = Guard
tmp2847 := MakeNative(func(__e *ControlFlow) {
News1330 := __e.Get(1)
_ = News1330
tmp2848 := Call(__e, PrimFunc(symshen_4in_1_6), News1330)


tmp2849 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp2850 := PrimCons(Action, Nil)

tmp2851 := PrimCons(symshen_4choicepoint_b, tmp2850)

tmp2852 := PrimCons(tmp2851, Nil)

tmp2853 := PrimCons(Guard, tmp2852)

tmp2854 := PrimCons(symwhere, tmp2853)

tmp2855 := Call(__e, PrimFunc(sym_8p), tmp2849, tmp2854)


__e.TailApply(PrimFunc(symshen_4comb), tmp2848, tmp2855)
return


}, 1)

tmp2856 := Call(__e, PrimFunc(symshen_4tls), News1329)


__e.TailApply(tmp2847, tmp2856)
return


}, 1)

tmp2857 := Call(__e, PrimFunc(symshen_4hds), News1329)


__e.TailApply(tmp2846, tmp2857)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2860 := Call(__e, PrimFunc(symshen_4tls), News1328)


__e.TailApply(tmp2845, tmp2860)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2863 := Call(__e, PrimFunc(symshen_4tls), News1327)


__e.TailApply(tmp2844, tmp2863)
return


}, 1)

tmp2864 := Call(__e, PrimFunc(symshen_4hds), News1327)


__e.TailApply(tmp2843, tmp2864)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2867 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp2842, tmp2867)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp2872 := Call(__e, PrimFunc(symshen_4_5patterns_6), V1477)


tmp2873 := Call(__e, tmp2841, tmp2872)


__e.TailApply(tmp2815, tmp2873)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2876 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp2891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp2891 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2889 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)


if True == tmp2889 {
tmp2877 := MakeNative(func(__e *ControlFlow) {
News1325 := __e.Get(1)
_ = News1325
tmp2886 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1325)


if True == tmp2886 {
tmp2878 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp2879 := MakeNative(func(__e *ControlFlow) {
News1326 := __e.Get(1)
_ = News1326
tmp2880 := Call(__e, PrimFunc(symshen_4in_1_6), News1326)


tmp2881 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp2882 := Call(__e, PrimFunc(sym_8p), tmp2881, Action)


__e.TailApply(PrimFunc(symshen_4comb), tmp2880, tmp2882)
return


}, 1)

tmp2883 := Call(__e, PrimFunc(symshen_4tls), News1325)


__e.TailApply(tmp2879, tmp2883)
return


}, 1)

tmp2884 := Call(__e, PrimFunc(symshen_4hds), News1325)


__e.TailApply(tmp2878, tmp2884)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2887 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp2877, tmp2887)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp2892 := Call(__e, PrimFunc(symshen_4_5patterns_6), V1477)


tmp2893 := Call(__e, tmp2876, tmp2892)


__e.TailApply(tmp2814, tmp2893)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2896 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp2924 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp2924 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2922 := Call(__e, PrimFunc(symshen_4_ahd_2), Parseshen_4_5patterns_6, sym_1_6)


if True == tmp2922 {
tmp2897 := MakeNative(func(__e *ControlFlow) {
News1321 := __e.Get(1)
_ = News1321
tmp2919 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1321)


if True == tmp2919 {
tmp2898 := MakeNative(func(__e *ControlFlow) {
Action := __e.Get(1)
_ = Action
tmp2899 := MakeNative(func(__e *ControlFlow) {
News1322 := __e.Get(1)
_ = News1322
tmp2915 := Call(__e, PrimFunc(symshen_4_ahd_2), News1322, symwhere)


if True == tmp2915 {
tmp2900 := MakeNative(func(__e *ControlFlow) {
News1323 := __e.Get(1)
_ = News1323
tmp2912 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), News1323)


if True == tmp2912 {
tmp2901 := MakeNative(func(__e *ControlFlow) {
Guard := __e.Get(1)
_ = Guard
tmp2902 := MakeNative(func(__e *ControlFlow) {
News1324 := __e.Get(1)
_ = News1324
tmp2903 := Call(__e, PrimFunc(symshen_4in_1_6), News1324)


tmp2904 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp2905 := PrimCons(Action, Nil)

tmp2906 := PrimCons(Guard, tmp2905)

tmp2907 := PrimCons(symwhere, tmp2906)

tmp2908 := Call(__e, PrimFunc(sym_8p), tmp2904, tmp2907)


__e.TailApply(PrimFunc(symshen_4comb), tmp2903, tmp2908)
return


}, 1)

tmp2909 := Call(__e, PrimFunc(symshen_4tls), News1323)


__e.TailApply(tmp2902, tmp2909)
return


}, 1)

tmp2910 := Call(__e, PrimFunc(symshen_4hds), News1323)


__e.TailApply(tmp2901, tmp2910)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2913 := Call(__e, PrimFunc(symshen_4tls), News1322)


__e.TailApply(tmp2900, tmp2913)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2916 := Call(__e, PrimFunc(symshen_4tls), News1321)


__e.TailApply(tmp2899, tmp2916)
return


}, 1)

tmp2917 := Call(__e, PrimFunc(symshen_4hds), News1321)


__e.TailApply(tmp2898, tmp2917)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2920 := Call(__e, PrimFunc(symshen_4tls), Parseshen_4_5patterns_6)


__e.TailApply(tmp2897, tmp2920)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}


}, 1)

tmp2925 := Call(__e, PrimFunc(symshen_4_5patterns_6), V1477)


tmp2926 := Call(__e, tmp2896, tmp2925)


__e.TailApply(tmp2813, tmp2926)
return


}, 1)

tmp2927 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp2812)


_ = tmp2927

tmp2928 := MakeNative(func(__e *ControlFlow) {
V1478 := __e.Get(1)
_ = V1478
tmp2929 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2940 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2940 {
tmp2930 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2932 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp2933 := MakeNative(func(__e *ControlFlow) {
Parse_5e_6 := __e.Get(1)
_ = Parse_5e_6
tmp2936 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)


if True == tmp2936 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2934 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp2934, Nil)
return


}


}, 1)

tmp2937 := Call(__e, PrimFunc(sym_5e_6), V1478)


tmp2938 := Call(__e, tmp2933, tmp2937)


__e.TailApply(tmp2930, tmp2938)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2941 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5pattern_6 := __e.Get(1)
_ = Parseshen_4_5pattern_6
tmp2951 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5pattern_6)


if True == tmp2951 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2942 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5patterns_6 := __e.Get(1)
_ = Parseshen_4_5patterns_6
tmp2948 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5patterns_6)


if True == tmp2948 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2943 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5patterns_6)


tmp2944 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5pattern_6)


tmp2945 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5patterns_6)


tmp2946 := PrimCons(tmp2944, tmp2945)

__e.TailApply(PrimFunc(symshen_4comb), tmp2943, tmp2946)
return


}


}, 1)

tmp2949 := Call(__e, PrimFunc(symshen_4_5patterns_6), Parseshen_4_5pattern_6)


__e.TailApply(tmp2942, tmp2949)
return


}


}, 1)

tmp2952 := Call(__e, PrimFunc(symshen_4_5pattern_6), V1478)


tmp2953 := Call(__e, tmp2941, tmp2952)


__e.TailApply(tmp2929, tmp2953)
return


}, 1)

tmp2954 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp2928)


_ = tmp2954

tmp2955 := MakeNative(func(__e *ControlFlow) {
V1479 := __e.Get(1)
_ = V1479
tmp2956 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp3010 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp3010 {
tmp2957 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2984 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2984 {
tmp2958 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2970 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2970 {
tmp2959 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp2961 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp2961 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp2962 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5simple_1pattern_6 := __e.Get(1)
_ = Parseshen_4_5simple_1pattern_6
tmp2966 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5simple_1pattern_6)


if True == tmp2966 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2963 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5simple_1pattern_6)


tmp2964 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5simple_1pattern_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp2963, tmp2964)
return


}


}, 1)

tmp2967 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V1479)


tmp2968 := Call(__e, tmp2962, tmp2967)


__e.TailApply(tmp2959, tmp2968)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp2982 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1479)


var ifres2971 Obj

if True == tmp2982 {
tmp2972 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp2973 := MakeNative(func(__e *ControlFlow) {
News1337 := __e.Get(1)
_ = News1337
tmp2977 := PrimIsPair(X)

if True == tmp2977 {
tmp2974 := Call(__e, PrimFunc(symshen_4in_1_6), News1337)


tmp2975 := Call(__e, PrimFunc(symshen_4constructor_1error), X)


__e.TailApply(PrimFunc(symshen_4comb), tmp2974, tmp2975)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2978 := Call(__e, PrimFunc(symshen_4tls), V1479)


__e.TailApply(tmp2973, tmp2978)
return


}, 1)

tmp2979 := Call(__e, PrimFunc(symshen_4hds), V1479)


tmp2980 := Call(__e, tmp2972, tmp2979)


ifres2971 = tmp2980


} else {
tmp2981 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2971 = tmp2981


}

__e.TailApply(tmp2958, ifres2971)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp3008 := Call(__e, PrimFunc(symshen_4ccons_2), V1479)


var ifres2985 Obj

if True == tmp3008 {
tmp2986 := MakeNative(func(__e *ControlFlow) {
SynCons := __e.Get(1)
_ = SynCons
tmp3002 := Call(__e, PrimFunc(symshen_4_ahd_2), SynCons, symvector)


if True == tmp3002 {
tmp2987 := MakeNative(func(__e *ControlFlow) {
News1335 := __e.Get(1)
_ = News1335
tmp2999 := Call(__e, PrimFunc(symshen_4_ahd_2), News1335, MakeNumber(0))


if True == tmp2999 {
tmp2988 := MakeNative(func(__e *ControlFlow) {
News1336 := __e.Get(1)
_ = News1336
tmp2989 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5end_6 := __e.Get(1)
_ = Parseshen_4_5end_6
tmp2995 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)


if True == tmp2995 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2990 := Call(__e, PrimFunc(symshen_4tlstream), V1479)


tmp2991 := Call(__e, PrimFunc(symshen_4in_1_6), tmp2990)


tmp2992 := PrimCons(MakeNumber(0), Nil)

tmp2993 := PrimCons(symvector, tmp2992)

__e.TailApply(PrimFunc(symshen_4comb), tmp2991, tmp2993)
return


}


}, 1)

tmp2996 := Call(__e, PrimFunc(symshen_4_5end_6), News1336)


__e.TailApply(tmp2989, tmp2996)
return


}, 1)

tmp2997 := Call(__e, PrimFunc(symshen_4tls), News1335)


__e.TailApply(tmp2988, tmp2997)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3000 := Call(__e, PrimFunc(symshen_4tls), SynCons)


__e.TailApply(tmp2987, tmp3000)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3003 := Call(__e, PrimFunc(symshen_4hds), V1479)


tmp3004 := Call(__e, PrimFunc(symshen_4_5_1out), V1479)


tmp3005 := Call(__e, PrimFunc(symshen_4comb), tmp3003, tmp3004)


tmp3006 := Call(__e, tmp2986, tmp3005)


ifres2985 = tmp3006


} else {
tmp3007 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2985 = tmp3007


}

__e.TailApply(tmp2957, ifres2985)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp3042 := Call(__e, PrimFunc(symshen_4ccons_2), V1479)


var ifres3011 Obj

if True == tmp3042 {
tmp3012 := MakeNative(func(__e *ControlFlow) {
SynCons := __e.Get(1)
_ = SynCons
tmp3013 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5constructor_6 := __e.Get(1)
_ = Parseshen_4_5constructor_6
tmp3035 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5constructor_6)


if True == tmp3035 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3014 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5pattern1_6 := __e.Get(1)
_ = Parseshen_4_5pattern1_6
tmp3032 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5pattern1_6)


if True == tmp3032 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3015 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5pattern2_6 := __e.Get(1)
_ = Parseshen_4_5pattern2_6
tmp3029 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5pattern2_6)


if True == tmp3029 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3016 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5end_6 := __e.Get(1)
_ = Parseshen_4_5end_6
tmp3026 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5end_6)


if True == tmp3026 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3017 := Call(__e, PrimFunc(symshen_4tlstream), V1479)


tmp3018 := Call(__e, PrimFunc(symshen_4in_1_6), tmp3017)


tmp3019 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5constructor_6)


tmp3020 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5pattern1_6)


tmp3021 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5pattern2_6)


tmp3022 := PrimCons(tmp3021, Nil)

tmp3023 := PrimCons(tmp3020, tmp3022)

tmp3024 := PrimCons(tmp3019, tmp3023)

__e.TailApply(PrimFunc(symshen_4comb), tmp3018, tmp3024)
return


}


}, 1)

tmp3027 := Call(__e, PrimFunc(symshen_4_5end_6), Parseshen_4_5pattern2_6)


__e.TailApply(tmp3016, tmp3027)
return


}


}, 1)

tmp3030 := Call(__e, PrimFunc(symshen_4_5pattern2_6), Parseshen_4_5pattern1_6)


__e.TailApply(tmp3015, tmp3030)
return


}


}, 1)

tmp3033 := Call(__e, PrimFunc(symshen_4_5pattern1_6), Parseshen_4_5constructor_6)


__e.TailApply(tmp3014, tmp3033)
return


}


}, 1)

tmp3036 := Call(__e, PrimFunc(symshen_4_5constructor_6), SynCons)


__e.TailApply(tmp3013, tmp3036)
return


}, 1)

tmp3037 := Call(__e, PrimFunc(symshen_4hds), V1479)


tmp3038 := Call(__e, PrimFunc(symshen_4_5_1out), V1479)


tmp3039 := Call(__e, PrimFunc(symshen_4comb), tmp3037, tmp3038)


tmp3040 := Call(__e, tmp3012, tmp3039)


ifres3011 = tmp3040


} else {
tmp3041 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3011 = tmp3041


}

__e.TailApply(tmp2956, ifres3011)
return


}, 1)

tmp3043 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp2955)


_ = tmp3043

tmp3044 := MakeNative(func(__e *ControlFlow) {
V1480 := __e.Get(1)
_ = V1480
tmp3045 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp3047 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp3047 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp3058 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1480)


var ifres3048 Obj

if True == tmp3058 {
tmp3049 := MakeNative(func(__e *ControlFlow) {
C := __e.Get(1)
_ = C
tmp3050 := MakeNative(func(__e *ControlFlow) {
News1339 := __e.Get(1)
_ = News1339
tmp3053 := Call(__e, PrimFunc(symshen_4constructor_2), C)


if True == tmp3053 {
tmp3051 := Call(__e, PrimFunc(symshen_4in_1_6), News1339)


__e.TailApply(PrimFunc(symshen_4comb), tmp3051, C)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3054 := Call(__e, PrimFunc(symshen_4tls), V1480)


__e.TailApply(tmp3050, tmp3054)
return


}, 1)

tmp3055 := Call(__e, PrimFunc(symshen_4hds), V1480)


tmp3056 := Call(__e, tmp3049, tmp3055)


ifres3048 = tmp3056


} else {
tmp3057 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3048 = tmp3057


}

__e.TailApply(tmp3045, ifres3048)
return


}, 1)

tmp3059 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp3044)


_ = tmp3059

tmp3060 := MakeNative(func(__e *ControlFlow) {
V1481 := __e.Get(1)
_ = V1481
tmp3061 := PrimCons(sym_8v, Nil)

tmp3062 := PrimCons(sym_8s, tmp3061)

tmp3063 := PrimCons(sym_8p, tmp3062)

tmp3064 := PrimCons(symcons, tmp3063)

__e.TailApply(PrimFunc(symelement_2), V1481, tmp3064)
return


}, 1)

tmp3065 := Call(__e, ns2_1set, symshen_4constructor_2, tmp3060)


_ = tmp3065

tmp3066 := MakeNative(func(__e *ControlFlow) {
V1482 := __e.Get(1)
_ = V1482
tmp3067 := Call(__e, PrimFunc(symshen_4app), V1482, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return(PrimSimpleError(tmp3067))
return


}, 1)

tmp3068 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp3066)


_ = tmp3068

tmp3069 := MakeNative(func(__e *ControlFlow) {
V1483 := __e.Get(1)
_ = V1483
tmp3070 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp3089 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp3089 {
tmp3071 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp3073 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp3073 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp3087 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1483)


var ifres3074 Obj

if True == tmp3087 {
tmp3075 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp3076 := MakeNative(func(__e *ControlFlow) {
News1342 := __e.Get(1)
_ = News1342
tmp3079 := PrimCons(sym_5_1, Nil)

tmp3080 := PrimCons(sym_1_6, tmp3079)

tmp3081 := Call(__e, PrimFunc(symelement_2), X, tmp3080)


tmp3082 := PrimNot(tmp3081)

if True == tmp3082 {
tmp3077 := Call(__e, PrimFunc(symshen_4in_1_6), News1342)


__e.TailApply(PrimFunc(symshen_4comb), tmp3077, X)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3083 := Call(__e, PrimFunc(symshen_4tls), V1483)


__e.TailApply(tmp3076, tmp3083)
return


}, 1)

tmp3084 := Call(__e, PrimFunc(symshen_4hds), V1483)


tmp3085 := Call(__e, tmp3075, tmp3084)


ifres3074 = tmp3085


} else {
tmp3086 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3074 = tmp3086


}

__e.TailApply(tmp3071, ifres3074)
return


} else {
__e.Return(Result)
return
}


}, 1)

tmp3102 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V1483)


var ifres3090 Obj

if True == tmp3102 {
tmp3091 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp3092 := MakeNative(func(__e *ControlFlow) {
News1341 := __e.Get(1)
_ = News1341
tmp3097 := PrimEqual(X, sym__)

if True == tmp3097 {
tmp3093 := Call(__e, PrimFunc(symshen_4in_1_6), News1341)


tmp3094 := Call(__e, PrimFunc(symprotect), symY)


tmp3095 := Call(__e, PrimFunc(symgensym), tmp3094)


__e.TailApply(PrimFunc(symshen_4comb), tmp3093, tmp3095)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3098 := Call(__e, PrimFunc(symshen_4tls), V1483)


__e.TailApply(tmp3092, tmp3098)
return


}, 1)

tmp3099 := Call(__e, PrimFunc(symshen_4hds), V1483)


tmp3100 := Call(__e, tmp3091, tmp3099)


ifres3090 = tmp3100


} else {
tmp3101 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3090 = tmp3101


}

__e.TailApply(tmp3070, ifres3090)
return


}, 1)

tmp3103 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3069)


_ = tmp3103

tmp3104 := MakeNative(func(__e *ControlFlow) {
V1484 := __e.Get(1)
_ = V1484
tmp3105 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp3107 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp3107 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp3108 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5pattern_6 := __e.Get(1)
_ = Parseshen_4_5pattern_6
tmp3112 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5pattern_6)


if True == tmp3112 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3109 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5pattern_6)


tmp3110 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5pattern_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp3109, tmp3110)
return


}


}, 1)

tmp3113 := Call(__e, PrimFunc(symshen_4_5pattern_6), V1484)


tmp3114 := Call(__e, tmp3108, tmp3113)


__e.TailApply(tmp3105, tmp3114)
return


}, 1)

tmp3115 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp3104)


_ = tmp3115

tmp3116 := MakeNative(func(__e *ControlFlow) {
V1485 := __e.Get(1)
_ = V1485
tmp3117 := MakeNative(func(__e *ControlFlow) {
Result := __e.Get(1)
_ = Result
tmp3119 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)


if True == tmp3119 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(Result)
return
}


}, 1)

tmp3120 := MakeNative(func(__e *ControlFlow) {
Parseshen_4_5pattern_6 := __e.Get(1)
_ = Parseshen_4_5pattern_6
tmp3124 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5pattern_6)


if True == tmp3124 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3121 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5pattern_6)


tmp3122 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5pattern_6)


__e.TailApply(PrimFunc(symshen_4comb), tmp3121, tmp3122)
return


}


}, 1)

tmp3125 := Call(__e, PrimFunc(symshen_4_5pattern_6), V1485)


tmp3126 := Call(__e, tmp3120, tmp3125)


__e.TailApply(tmp3117, tmp3126)
return


}, 1)

tmp3127 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp3116)


_ = tmp3127

tmp3128 := MakeNative(func(__e *ControlFlow) {
V1486 := __e.Get(1)
_ = V1486
tmp3129 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
tmp3130 := MakeNative(func(__e *ControlFlow) {
Print := __e.Get(1)
_ = Print
tmp3131 := MakeNative(func(__e *ControlFlow) {
Named := __e.Get(1)
_ = Named
__e.Return(Named)
return
}, 1)

tmp3132 := PrimStr(V1486)

tmp3133 := Call(__e, PrimFunc(sym_8s), tmp3132, MakeString(")"))


tmp3134 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp3133)


tmp3135 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp3134)


tmp3136 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp3135)


tmp3137 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp3136)


tmp3138 := PrimVectorSet(Print, MakeNumber(1), tmp3137)

__e.TailApply(tmp3131, tmp3138)
return


}, 1)

tmp3139 := PrimVectorSet(V, MakeNumber(0), symshen_4printF)

__e.TailApply(tmp3130, tmp3139)
return


}, 1)

tmp3140 := PrimAbsvector(MakeNumber(2))

__e.TailApply(tmp3129, tmp3140)
return


}, 1)

tmp3141 := Call(__e, ns2_1set, symshen_4fn_1print, tmp3128)


_ = tmp3141

tmp3142 := MakeNative(func(__e *ControlFlow) {
V1487 := __e.Get(1)
_ = V1487
__e.Return(PrimVectorGet(V1487, MakeNumber(1)))
return
}, 1)

tmp3143 := Call(__e, ns2_1set, symshen_4printF, tmp3142)


_ = tmp3143

tmp3144 := MakeNative(func(__e *ControlFlow) {
V1492 := __e.Get(1)
_ = V1492
V1493 := __e.Get(2)
_ = V1493
tmp3168 := PrimIsPair(V1493)

var ifres3164 Obj

if True == tmp3168 {
tmp3166 := PrimTail(V1493)

tmp3167 := PrimEqual(Nil, tmp3166)

var ifres3165 Obj

if True == tmp3167 {
ifres3165 = True


} else {
ifres3165 = False


}

ifres3164 = ifres3165


} else {
ifres3164 = False


}

if True == ifres3164 {
tmp3145 := PrimHead(V1493)

__e.TailApply(PrimFunc(symlength), tmp3145)
return


} else {
tmp3162 := PrimIsPair(V1493)

var ifres3150 Obj

if True == tmp3162 {
tmp3160 := PrimTail(V1493)

tmp3161 := PrimIsPair(tmp3160)

var ifres3152 Obj

if True == tmp3161 {
tmp3154 := PrimHead(V1493)

tmp3155 := Call(__e, PrimFunc(symlength), tmp3154)


tmp3156 := PrimTail(V1493)

tmp3157 := PrimHead(tmp3156)

tmp3158 := Call(__e, PrimFunc(symlength), tmp3157)


tmp3159 := PrimEqual(tmp3155, tmp3158)

var ifres3153 Obj

if True == tmp3159 {
ifres3153 = True


} else {
ifres3153 = False


}

ifres3152 = ifres3153


} else {
ifres3152 = False


}

var ifres3151 Obj

if True == ifres3152 {
ifres3151 = True


} else {
ifres3151 = False


}

ifres3150 = ifres3151


} else {
ifres3150 = False


}

if True == ifres3150 {
tmp3146 := PrimTail(V1493)

__e.TailApply(PrimFunc(symshen_4arity_1chk), V1492, tmp3146)
return


} else {
tmp3147 := Call(__e, PrimFunc(symshen_4app), V1492, MakeString("\n"), symshen_4a)


tmp3148 := PrimStringConcat(MakeString("arity error in "), tmp3147)

__e.Return(PrimSimpleError(tmp3148))
return


}


}


}, 2)

tmp3169 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp3144)


_ = tmp3169

tmp3170 := MakeNative(func(__e *ControlFlow) {
V1494 := __e.Get(1)
_ = V1494
V1495 := __e.Get(2)
_ = V1495
tmp3176 := Call(__e, PrimFunc(symtuple_2), V1495)


if True == tmp3176 {
tmp3171 := Call(__e, PrimFunc(symfst), V1495)


tmp3172 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3171)


tmp3173 := Call(__e, PrimFunc(symsnd), V1495)


tmp3174 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3172, tmp3173)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V1494, tmp3174)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4free_1var_1chk)
return
}


}, 2)

tmp3177 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp3170)


_ = tmp3177

tmp3178 := MakeNative(func(__e *ControlFlow) {
V1496 := __e.Get(1)
_ = V1496
V1497 := __e.Get(2)
_ = V1497
tmp3190 := Call(__e, PrimFunc(symempty_2), V1497)


if True == tmp3190 {
__e.Return(symshen_4skip)
return
} else {
tmp3179 := Call(__e, PrimFunc(symshen_4app), V1496, MakeString(":"), symshen_4a)


tmp3180 := PrimStringConcat(MakeString("free variables in "), tmp3179)

tmp3181 := Call(__e, PrimFunc(symstoutput))


tmp3182 := Call(__e, PrimFunc(sympr), tmp3180, tmp3181)


_ = tmp3182

tmp3183 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
tmp3184 := Call(__e, PrimFunc(symshen_4app), X, MakeString(""), symshen_4a)


tmp3185 := PrimStringConcat(MakeString(" "), tmp3184)

tmp3186 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp3185, tmp3186)
return


}, 1)

tmp3187 := Call(__e, PrimFunc(symmap), tmp3183, V1497)


_ = tmp3187

tmp3188 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp3188

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp3191 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3178)


_ = tmp3191

tmp3192 := MakeNative(func(__e *ControlFlow) {
V1500 := __e.Get(1)
_ = V1500
tmp3200 := PrimIsVariable(V1500)

if True == tmp3200 {
__e.Return(PrimCons(V1500, Nil))
return
} else {
tmp3198 := PrimIsPair(V1500)

if True == tmp3198 {
tmp3193 := PrimHead(V1500)

tmp3194 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3193)


tmp3195 := PrimTail(V1500)

tmp3196 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3195)


__e.TailApply(PrimFunc(symunion), tmp3194, tmp3196)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3201 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp3192)


_ = tmp3201

tmp3202 := MakeNative(func(__e *ControlFlow) {
V1505 := __e.Get(1)
_ = V1505
V1506 := __e.Get(2)
_ = V1506
tmp3292 := PrimIsPair(V1506)

var ifres3279 Obj

if True == tmp3292 {
tmp3290 := PrimHead(V1506)

tmp3291 := PrimEqual(symprotect, tmp3290)

var ifres3281 Obj

if True == tmp3291 {
tmp3288 := PrimTail(V1506)

tmp3289 := PrimIsPair(tmp3288)

var ifres3283 Obj

if True == tmp3289 {
tmp3285 := PrimTail(V1506)

tmp3286 := PrimTail(tmp3285)

tmp3287 := PrimEqual(Nil, tmp3286)

var ifres3284 Obj

if True == tmp3287 {
ifres3284 = True


} else {
ifres3284 = False


}

ifres3283 = ifres3284


} else {
ifres3283 = False


}

var ifres3282 Obj

if True == ifres3283 {
ifres3282 = True


} else {
ifres3282 = False


}

ifres3281 = ifres3282


} else {
ifres3281 = False


}

var ifres3280 Obj

if True == ifres3281 {
ifres3280 = True


} else {
ifres3280 = False


}

ifres3279 = ifres3280


} else {
ifres3279 = False


}

if True == ifres3279 {
__e.Return(Nil)
return
} else {
tmp3277 := PrimIsPair(V1506)

var ifres3251 Obj

if True == tmp3277 {
tmp3275 := PrimHead(V1506)

tmp3276 := PrimEqual(symlet, tmp3275)

var ifres3253 Obj

if True == tmp3276 {
tmp3273 := PrimTail(V1506)

tmp3274 := PrimIsPair(tmp3273)

var ifres3255 Obj

if True == tmp3274 {
tmp3270 := PrimTail(V1506)

tmp3271 := PrimTail(tmp3270)

tmp3272 := PrimIsPair(tmp3271)

var ifres3257 Obj

if True == tmp3272 {
tmp3266 := PrimTail(V1506)

tmp3267 := PrimTail(tmp3266)

tmp3268 := PrimTail(tmp3267)

tmp3269 := PrimIsPair(tmp3268)

var ifres3259 Obj

if True == tmp3269 {
tmp3261 := PrimTail(V1506)

tmp3262 := PrimTail(tmp3261)

tmp3263 := PrimTail(tmp3262)

tmp3264 := PrimTail(tmp3263)

tmp3265 := PrimEqual(Nil, tmp3264)

var ifres3260 Obj

if True == tmp3265 {
ifres3260 = True


} else {
ifres3260 = False


}

ifres3259 = ifres3260


} else {
ifres3259 = False


}

var ifres3258 Obj

if True == ifres3259 {
ifres3258 = True


} else {
ifres3258 = False


}

ifres3257 = ifres3258


} else {
ifres3257 = False


}

var ifres3256 Obj

if True == ifres3257 {
ifres3256 = True


} else {
ifres3256 = False


}

ifres3255 = ifres3256


} else {
ifres3255 = False


}

var ifres3254 Obj

if True == ifres3255 {
ifres3254 = True


} else {
ifres3254 = False


}

ifres3253 = ifres3254


} else {
ifres3253 = False


}

var ifres3252 Obj

if True == ifres3253 {
ifres3252 = True


} else {
ifres3252 = False


}

ifres3251 = ifres3252


} else {
ifres3251 = False


}

if True == ifres3251 {
tmp3203 := PrimTail(V1506)

tmp3204 := PrimTail(tmp3203)

tmp3205 := PrimHead(tmp3204)

tmp3206 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V1505, tmp3205)


tmp3207 := PrimTail(V1506)

tmp3208 := PrimHead(tmp3207)

tmp3209 := PrimCons(tmp3208, V1505)

tmp3210 := PrimTail(V1506)

tmp3211 := PrimTail(tmp3210)

tmp3212 := PrimTail(tmp3211)

tmp3213 := PrimHead(tmp3212)

tmp3214 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3209, tmp3213)


__e.TailApply(PrimFunc(symunion), tmp3206, tmp3214)
return


} else {
tmp3249 := PrimIsPair(V1506)

var ifres3230 Obj

if True == tmp3249 {
tmp3247 := PrimHead(V1506)

tmp3248 := PrimEqual(symlambda, tmp3247)

var ifres3232 Obj

if True == tmp3248 {
tmp3245 := PrimTail(V1506)

tmp3246 := PrimIsPair(tmp3245)

var ifres3234 Obj

if True == tmp3246 {
tmp3242 := PrimTail(V1506)

tmp3243 := PrimTail(tmp3242)

tmp3244 := PrimIsPair(tmp3243)

var ifres3236 Obj

if True == tmp3244 {
tmp3238 := PrimTail(V1506)

tmp3239 := PrimTail(tmp3238)

tmp3240 := PrimTail(tmp3239)

tmp3241 := PrimEqual(Nil, tmp3240)

var ifres3237 Obj

if True == tmp3241 {
ifres3237 = True


} else {
ifres3237 = False


}

ifres3236 = ifres3237


} else {
ifres3236 = False


}

var ifres3235 Obj

if True == ifres3236 {
ifres3235 = True


} else {
ifres3235 = False


}

ifres3234 = ifres3235


} else {
ifres3234 = False


}

var ifres3233 Obj

if True == ifres3234 {
ifres3233 = True


} else {
ifres3233 = False


}

ifres3232 = ifres3233


} else {
ifres3232 = False


}

var ifres3231 Obj

if True == ifres3232 {
ifres3231 = True


} else {
ifres3231 = False


}

ifres3230 = ifres3231


} else {
ifres3230 = False


}

if True == ifres3230 {
tmp3215 := PrimTail(V1506)

tmp3216 := PrimHead(tmp3215)

tmp3217 := PrimCons(tmp3216, V1505)

tmp3218 := PrimTail(V1506)

tmp3219 := PrimTail(tmp3218)

tmp3220 := PrimHead(tmp3219)

__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp3217, tmp3220)
return


} else {
tmp3228 := PrimIsPair(V1506)

if True == tmp3228 {
tmp3221 := PrimHead(V1506)

tmp3222 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V1505, tmp3221)


tmp3223 := PrimTail(V1506)

tmp3224 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V1505, tmp3223)


__e.TailApply(PrimFunc(symunion), tmp3222, tmp3224)
return


} else {
tmp3226 := Call(__e, PrimFunc(symshen_4free_1variable_2), V1506, V1505)


if True == tmp3226 {
__e.Return(PrimCons(V1506, Nil))
return
} else {
__e.Return(Nil)
return
}


}


}


}


}


}, 2)

tmp3293 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp3202)


_ = tmp3293

tmp3294 := MakeNative(func(__e *ControlFlow) {
V1507 := __e.Get(1)
_ = V1507
V1508 := __e.Get(2)
_ = V1508
tmp3299 := PrimIsVariable(V1507)

if True == tmp3299 {
tmp3296 := Call(__e, PrimFunc(symelement_2), V1507, V1508)


tmp3297 := PrimNot(tmp3296)

if True == tmp3297 {
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

tmp3300 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp3294)


_ = tmp3300

tmp3301 := MakeNative(func(__e *ControlFlow) {
V1509 := __e.Get(1)
_ = V1509
V1510 := __e.Get(2)
_ = V1510
tmp3302 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V1509, symshen_4source, V1510, tmp3302)
return


}, 2)

tmp3303 := Call(__e, ns2_1set, symshen_4record_1kl, tmp3301)


_ = tmp3303

tmp3304 := MakeNative(func(__e *ControlFlow) {
V1511 := __e.Get(1)
_ = V1511
V1512 := __e.Get(2)
_ = V1512
V1513 := __e.Get(3)
_ = V1513
tmp3305 := MakeNative(func(__e *ControlFlow) {
Parameters := __e.Get(1)
_ = Parameters
tmp3306 := MakeNative(func(__e *ControlFlow) {
Body := __e.Get(1)
_ = Body
tmp3307 := MakeNative(func(__e *ControlFlow) {
Defun := __e.Get(1)
_ = Defun
__e.Return(Defun)
return
}, 1)

tmp3308 := Call(__e, PrimFunc(symshen_4cond_1form), Body)


tmp3309 := PrimCons(tmp3308, Nil)

tmp3310 := PrimCons(Parameters, tmp3309)

tmp3311 := PrimCons(V1511, tmp3310)

tmp3312 := PrimCons(symdefun, tmp3311)

__e.TailApply(tmp3307, tmp3312)
return


}, 1)

tmp3313 := Call(__e, PrimFunc(symshen_4kl_1body), V1512, Parameters)


tmp3314 := Call(__e, PrimFunc(symshen_4scan_1body), V1511, tmp3313)


__e.TailApply(tmp3306, tmp3314)
return


}, 1)

tmp3315 := Call(__e, PrimFunc(symshen_4parameters), V1513)


__e.TailApply(tmp3305, tmp3315)
return


}, 3)

tmp3316 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3304)


_ = tmp3316

tmp3317 := MakeNative(func(__e *ControlFlow) {
V1514 := __e.Get(1)
_ = V1514
tmp3323 := PrimEqual(MakeNumber(0), V1514)

if True == tmp3323 {
__e.Return(Nil)
return
} else {
tmp3318 := Call(__e, PrimFunc(symprotect), symV)


tmp3319 := Call(__e, PrimFunc(symgensym), tmp3318)


tmp3320 := PrimNumberSubtract(V1514, MakeNumber(1))

tmp3321 := Call(__e, PrimFunc(symshen_4parameters), tmp3320)


__e.Return(PrimCons(tmp3319, tmp3321))
return


}


}, 1)

tmp3324 := Call(__e, ns2_1set, symshen_4parameters, tmp3317)


_ = tmp3324

tmp3325 := MakeNative(func(__e *ControlFlow) {
V1517 := __e.Get(1)
_ = V1517
tmp3349 := PrimIsPair(V1517)

var ifres3329 Obj

if True == tmp3349 {
tmp3347 := PrimHead(V1517)

tmp3348 := PrimIsPair(tmp3347)

var ifres3331 Obj

if True == tmp3348 {
tmp3344 := PrimHead(V1517)

tmp3345 := PrimHead(tmp3344)

tmp3346 := PrimEqual(True, tmp3345)

var ifres3333 Obj

if True == tmp3346 {
tmp3341 := PrimHead(V1517)

tmp3342 := PrimTail(tmp3341)

tmp3343 := PrimIsPair(tmp3342)

var ifres3335 Obj

if True == tmp3343 {
tmp3337 := PrimHead(V1517)

tmp3338 := PrimTail(tmp3337)

tmp3339 := PrimTail(tmp3338)

tmp3340 := PrimEqual(Nil, tmp3339)

var ifres3336 Obj

if True == tmp3340 {
ifres3336 = True


} else {
ifres3336 = False


}

ifres3335 = ifres3336


} else {
ifres3335 = False


}

var ifres3334 Obj

if True == ifres3335 {
ifres3334 = True


} else {
ifres3334 = False


}

ifres3333 = ifres3334


} else {
ifres3333 = False


}

var ifres3332 Obj

if True == ifres3333 {
ifres3332 = True


} else {
ifres3332 = False


}

ifres3331 = ifres3332


} else {
ifres3331 = False


}

var ifres3330 Obj

if True == ifres3331 {
ifres3330 = True


} else {
ifres3330 = False


}

ifres3329 = ifres3330


} else {
ifres3329 = False


}

if True == ifres3329 {
tmp3326 := PrimHead(V1517)

tmp3327 := PrimTail(tmp3326)

__e.Return(PrimHead(tmp3327))
return


} else {
__e.Return(PrimCons(symcond, V1517))
return
}


}, 1)

tmp3350 := Call(__e, ns2_1set, symshen_4cond_1form, tmp3325)


_ = tmp3350

tmp3351 := MakeNative(func(__e *ControlFlow) {
V1526 := __e.Get(1)
_ = V1526
V1527 := __e.Get(2)
_ = V1527
tmp3397 := PrimEqual(Nil, V1527)

if True == tmp3397 {
tmp3352 := PrimCons(V1526, Nil)

tmp3353 := PrimCons(symshen_4f_1error, tmp3352)

tmp3354 := PrimCons(tmp3353, Nil)

tmp3355 := PrimCons(True, tmp3354)

__e.Return(PrimCons(tmp3355, Nil))
return


} else {
tmp3395 := PrimIsPair(V1527)

var ifres3391 Obj

if True == tmp3395 {
tmp3393 := PrimHead(V1527)

tmp3394 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp3393)


var ifres3392 Obj

if True == tmp3394 {
ifres3392 = True


} else {
ifres3392 = False


}

ifres3391 = ifres3392


} else {
ifres3391 = False


}

if True == ifres3391 {
tmp3356 := Call(__e, PrimFunc(symprotect), symFreeze)


tmp3357 := Call(__e, PrimFunc(symgensym), tmp3356)


tmp3358 := Call(__e, PrimFunc(symprotect), symResult)


tmp3359 := Call(__e, PrimFunc(symgensym), tmp3358)


tmp3360 := PrimHead(V1527)

tmp3361 := PrimTail(V1527)

__e.TailApply(PrimFunc(symshen_4choicepoint), V1526, tmp3357, tmp3359, tmp3360, tmp3361)
return


} else {
tmp3389 := PrimIsPair(V1527)

var ifres3369 Obj

if True == tmp3389 {
tmp3387 := PrimHead(V1527)

tmp3388 := PrimIsPair(tmp3387)

var ifres3371 Obj

if True == tmp3388 {
tmp3384 := PrimHead(V1527)

tmp3385 := PrimHead(tmp3384)

tmp3386 := PrimEqual(True, tmp3385)

var ifres3373 Obj

if True == tmp3386 {
tmp3381 := PrimHead(V1527)

tmp3382 := PrimTail(tmp3381)

tmp3383 := PrimIsPair(tmp3382)

var ifres3375 Obj

if True == tmp3383 {
tmp3377 := PrimHead(V1527)

tmp3378 := PrimTail(tmp3377)

tmp3379 := PrimTail(tmp3378)

tmp3380 := PrimEqual(Nil, tmp3379)

var ifres3376 Obj

if True == tmp3380 {
ifres3376 = True


} else {
ifres3376 = False


}

ifres3375 = ifres3376


} else {
ifres3375 = False


}

var ifres3374 Obj

if True == ifres3375 {
ifres3374 = True


} else {
ifres3374 = False


}

ifres3373 = ifres3374


} else {
ifres3373 = False


}

var ifres3372 Obj

if True == ifres3373 {
ifres3372 = True


} else {
ifres3372 = False


}

ifres3371 = ifres3372


} else {
ifres3371 = False


}

var ifres3370 Obj

if True == ifres3371 {
ifres3370 = True


} else {
ifres3370 = False


}

ifres3369 = ifres3370


} else {
ifres3369 = False


}

if True == ifres3369 {
tmp3362 := PrimHead(V1527)

__e.Return(PrimCons(tmp3362, Nil))
return


} else {
tmp3367 := PrimIsPair(V1527)

if True == tmp3367 {
tmp3363 := PrimHead(V1527)

tmp3364 := PrimTail(V1527)

tmp3365 := Call(__e, PrimFunc(symshen_4scan_1body), V1526, tmp3364)


__e.Return(PrimCons(tmp3363, tmp3365))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.scan-body")))
return
}


}


}


}


}, 2)

tmp3398 := Call(__e, ns2_1set, symshen_4scan_1body, tmp3351)


_ = tmp3398

tmp3399 := MakeNative(func(__e *ControlFlow) {
V1534 := __e.Get(1)
_ = V1534
tmp3434 := PrimIsPair(V1534)

var ifres3401 Obj

if True == tmp3434 {
tmp3432 := PrimTail(V1534)

tmp3433 := PrimIsPair(tmp3432)

var ifres3403 Obj

if True == tmp3433 {
tmp3429 := PrimTail(V1534)

tmp3430 := PrimHead(tmp3429)

tmp3431 := PrimIsPair(tmp3430)

var ifres3405 Obj

if True == tmp3431 {
tmp3425 := PrimTail(V1534)

tmp3426 := PrimHead(tmp3425)

tmp3427 := PrimHead(tmp3426)

tmp3428 := PrimEqual(symshen_4choicepoint_b, tmp3427)

var ifres3407 Obj

if True == tmp3428 {
tmp3421 := PrimTail(V1534)

tmp3422 := PrimHead(tmp3421)

tmp3423 := PrimTail(tmp3422)

tmp3424 := PrimIsPair(tmp3423)

var ifres3409 Obj

if True == tmp3424 {
tmp3416 := PrimTail(V1534)

tmp3417 := PrimHead(tmp3416)

tmp3418 := PrimTail(tmp3417)

tmp3419 := PrimTail(tmp3418)

tmp3420 := PrimEqual(Nil, tmp3419)

var ifres3411 Obj

if True == tmp3420 {
tmp3413 := PrimTail(V1534)

tmp3414 := PrimTail(tmp3413)

tmp3415 := PrimEqual(Nil, tmp3414)

var ifres3412 Obj

if True == tmp3415 {
ifres3412 = True


} else {
ifres3412 = False


}

ifres3411 = ifres3412


} else {
ifres3411 = False


}

var ifres3410 Obj

if True == ifres3411 {
ifres3410 = True


} else {
ifres3410 = False


}

ifres3409 = ifres3410


} else {
ifres3409 = False


}

var ifres3408 Obj

if True == ifres3409 {
ifres3408 = True


} else {
ifres3408 = False


}

ifres3407 = ifres3408


} else {
ifres3407 = False


}

var ifres3406 Obj

if True == ifres3407 {
ifres3406 = True


} else {
ifres3406 = False


}

ifres3405 = ifres3406


} else {
ifres3405 = False


}

var ifres3404 Obj

if True == ifres3405 {
ifres3404 = True


} else {
ifres3404 = False


}

ifres3403 = ifres3404


} else {
ifres3403 = False


}

var ifres3402 Obj

if True == ifres3403 {
ifres3402 = True


} else {
ifres3402 = False


}

ifres3401 = ifres3402


} else {
ifres3401 = False


}

if True == ifres3401 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3435 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp3399)


_ = tmp3435

tmp3436 := MakeNative(func(__e *ControlFlow) {
V1550 := __e.Get(1)
_ = V1550
V1551 := __e.Get(2)
_ = V1551
V1552 := __e.Get(3)
_ = V1552
V1553 := __e.Get(4)
_ = V1553
V1554 := __e.Get(5)
_ = V1554
tmp3628 := PrimIsPair(V1553)

var ifres3550 Obj

if True == tmp3628 {
tmp3626 := PrimTail(V1553)

tmp3627 := PrimIsPair(tmp3626)

var ifres3552 Obj

if True == tmp3627 {
tmp3623 := PrimTail(V1553)

tmp3624 := PrimHead(tmp3623)

tmp3625 := PrimIsPair(tmp3624)

var ifres3554 Obj

if True == tmp3625 {
tmp3619 := PrimTail(V1553)

tmp3620 := PrimHead(tmp3619)

tmp3621 := PrimTail(tmp3620)

tmp3622 := PrimIsPair(tmp3621)

var ifres3556 Obj

if True == tmp3622 {
tmp3614 := PrimTail(V1553)

tmp3615 := PrimHead(tmp3614)

tmp3616 := PrimTail(tmp3615)

tmp3617 := PrimHead(tmp3616)

tmp3618 := PrimIsPair(tmp3617)

var ifres3558 Obj

if True == tmp3618 {
tmp3608 := PrimTail(V1553)

tmp3609 := PrimHead(tmp3608)

tmp3610 := PrimTail(tmp3609)

tmp3611 := PrimHead(tmp3610)

tmp3612 := PrimHead(tmp3611)

tmp3613 := PrimEqual(symfail_1if, tmp3612)

var ifres3560 Obj

if True == tmp3613 {
tmp3602 := PrimTail(V1553)

tmp3603 := PrimHead(tmp3602)

tmp3604 := PrimTail(tmp3603)

tmp3605 := PrimHead(tmp3604)

tmp3606 := PrimTail(tmp3605)

tmp3607 := PrimIsPair(tmp3606)

var ifres3562 Obj

if True == tmp3607 {
tmp3595 := PrimTail(V1553)

tmp3596 := PrimHead(tmp3595)

tmp3597 := PrimTail(tmp3596)

tmp3598 := PrimHead(tmp3597)

tmp3599 := PrimTail(tmp3598)

tmp3600 := PrimTail(tmp3599)

tmp3601 := PrimIsPair(tmp3600)

var ifres3564 Obj

if True == tmp3601 {
tmp3587 := PrimTail(V1553)

tmp3588 := PrimHead(tmp3587)

tmp3589 := PrimTail(tmp3588)

tmp3590 := PrimHead(tmp3589)

tmp3591 := PrimTail(tmp3590)

tmp3592 := PrimTail(tmp3591)

tmp3593 := PrimTail(tmp3592)

tmp3594 := PrimEqual(Nil, tmp3593)

var ifres3566 Obj

if True == tmp3594 {
tmp3582 := PrimTail(V1553)

tmp3583 := PrimHead(tmp3582)

tmp3584 := PrimTail(tmp3583)

tmp3585 := PrimTail(tmp3584)

tmp3586 := PrimEqual(Nil, tmp3585)

var ifres3568 Obj

if True == tmp3586 {
tmp3579 := PrimTail(V1553)

tmp3580 := PrimTail(tmp3579)

tmp3581 := PrimEqual(Nil, tmp3580)

var ifres3570 Obj

if True == tmp3581 {
tmp3572 := PrimTail(V1553)

tmp3573 := PrimHead(tmp3572)

tmp3574 := PrimTail(tmp3573)

tmp3575 := PrimHead(tmp3574)

tmp3576 := PrimTail(tmp3575)

tmp3577 := PrimHead(tmp3576)

tmp3578 := PrimEqual(V1550, tmp3577)

var ifres3571 Obj

if True == tmp3578 {
ifres3571 = True


} else {
ifres3571 = False


}

ifres3570 = ifres3571


} else {
ifres3570 = False


}

var ifres3569 Obj

if True == ifres3570 {
ifres3569 = True


} else {
ifres3569 = False


}

ifres3568 = ifres3569


} else {
ifres3568 = False


}

var ifres3567 Obj

if True == ifres3568 {
ifres3567 = True


} else {
ifres3567 = False


}

ifres3566 = ifres3567


} else {
ifres3566 = False


}

var ifres3565 Obj

if True == ifres3566 {
ifres3565 = True


} else {
ifres3565 = False


}

ifres3564 = ifres3565


} else {
ifres3564 = False


}

var ifres3563 Obj

if True == ifres3564 {
ifres3563 = True


} else {
ifres3563 = False


}

ifres3562 = ifres3563


} else {
ifres3562 = False


}

var ifres3561 Obj

if True == ifres3562 {
ifres3561 = True


} else {
ifres3561 = False


}

ifres3560 = ifres3561


} else {
ifres3560 = False


}

var ifres3559 Obj

if True == ifres3560 {
ifres3559 = True


} else {
ifres3559 = False


}

ifres3558 = ifres3559


} else {
ifres3558 = False


}

var ifres3557 Obj

if True == ifres3558 {
ifres3557 = True


} else {
ifres3557 = False


}

ifres3556 = ifres3557


} else {
ifres3556 = False


}

var ifres3555 Obj

if True == ifres3556 {
ifres3555 = True


} else {
ifres3555 = False


}

ifres3554 = ifres3555


} else {
ifres3554 = False


}

var ifres3553 Obj

if True == ifres3554 {
ifres3553 = True


} else {
ifres3553 = False


}

ifres3552 = ifres3553


} else {
ifres3552 = False


}

var ifres3551 Obj

if True == ifres3552 {
ifres3551 = True


} else {
ifres3551 = False


}

ifres3550 = ifres3551


} else {
ifres3550 = False


}

if True == ifres3550 {
tmp3437 := PrimTail(V1553)

tmp3438 := PrimHead(tmp3437)

tmp3439 := PrimTail(tmp3438)

tmp3440 := PrimHead(tmp3439)

tmp3441 := PrimTail(tmp3440)

tmp3442 := PrimHead(tmp3441)

tmp3443 := Call(__e, PrimFunc(symshen_4scan_1body), tmp3442, V1554)


tmp3444 := PrimCons(symcond, tmp3443)

tmp3445 := PrimCons(tmp3444, Nil)

tmp3446 := PrimCons(symfreeze, tmp3445)

tmp3447 := PrimHead(V1553)

tmp3448 := PrimTail(V1553)

tmp3449 := PrimHead(tmp3448)

tmp3450 := PrimTail(tmp3449)

tmp3451 := PrimHead(tmp3450)

tmp3452 := PrimTail(tmp3451)

tmp3453 := PrimTail(tmp3452)

tmp3454 := PrimHead(tmp3453)

tmp3455 := PrimTail(V1553)

tmp3456 := PrimHead(tmp3455)

tmp3457 := PrimTail(tmp3456)

tmp3458 := PrimHead(tmp3457)

tmp3459 := PrimTail(tmp3458)

tmp3460 := PrimHead(tmp3459)

tmp3461 := PrimCons(V1552, Nil)

tmp3462 := PrimCons(tmp3460, tmp3461)

tmp3463 := PrimCons(V1551, Nil)

tmp3464 := PrimCons(symthaw, tmp3463)

tmp3465 := PrimCons(V1552, Nil)

tmp3466 := PrimCons(tmp3464, tmp3465)

tmp3467 := PrimCons(tmp3462, tmp3466)

tmp3468 := PrimCons(symif, tmp3467)

tmp3469 := PrimCons(tmp3468, Nil)

tmp3470 := PrimCons(tmp3454, tmp3469)

tmp3471 := PrimCons(V1552, tmp3470)

tmp3472 := PrimCons(symlet, tmp3471)

tmp3473 := PrimCons(V1551, Nil)

tmp3474 := PrimCons(symthaw, tmp3473)

tmp3475 := PrimCons(tmp3474, Nil)

tmp3476 := PrimCons(tmp3472, tmp3475)

tmp3477 := PrimCons(tmp3447, tmp3476)

tmp3478 := PrimCons(symif, tmp3477)

tmp3479 := PrimCons(tmp3478, Nil)

tmp3480 := PrimCons(tmp3446, tmp3479)

tmp3481 := PrimCons(V1551, tmp3480)

tmp3482 := PrimCons(symlet, tmp3481)

tmp3483 := PrimCons(tmp3482, Nil)

tmp3484 := PrimCons(True, tmp3483)

__e.Return(PrimCons(tmp3484, Nil))
return


} else {
tmp3548 := PrimIsPair(V1553)

var ifres3521 Obj

if True == tmp3548 {
tmp3546 := PrimTail(V1553)

tmp3547 := PrimIsPair(tmp3546)

var ifres3523 Obj

if True == tmp3547 {
tmp3543 := PrimTail(V1553)

tmp3544 := PrimHead(tmp3543)

tmp3545 := PrimIsPair(tmp3544)

var ifres3525 Obj

if True == tmp3545 {
tmp3539 := PrimTail(V1553)

tmp3540 := PrimHead(tmp3539)

tmp3541 := PrimTail(tmp3540)

tmp3542 := PrimIsPair(tmp3541)

var ifres3527 Obj

if True == tmp3542 {
tmp3534 := PrimTail(V1553)

tmp3535 := PrimHead(tmp3534)

tmp3536 := PrimTail(tmp3535)

tmp3537 := PrimTail(tmp3536)

tmp3538 := PrimEqual(Nil, tmp3537)

var ifres3529 Obj

if True == tmp3538 {
tmp3531 := PrimTail(V1553)

tmp3532 := PrimTail(tmp3531)

tmp3533 := PrimEqual(Nil, tmp3532)

var ifres3530 Obj

if True == tmp3533 {
ifres3530 = True


} else {
ifres3530 = False


}

ifres3529 = ifres3530


} else {
ifres3529 = False


}

var ifres3528 Obj

if True == ifres3529 {
ifres3528 = True


} else {
ifres3528 = False


}

ifres3527 = ifres3528


} else {
ifres3527 = False


}

var ifres3526 Obj

if True == ifres3527 {
ifres3526 = True


} else {
ifres3526 = False


}

ifres3525 = ifres3526


} else {
ifres3525 = False


}

var ifres3524 Obj

if True == ifres3525 {
ifres3524 = True


} else {
ifres3524 = False


}

ifres3523 = ifres3524


} else {
ifres3523 = False


}

var ifres3522 Obj

if True == ifres3523 {
ifres3522 = True


} else {
ifres3522 = False


}

ifres3521 = ifres3522


} else {
ifres3521 = False


}

if True == ifres3521 {
tmp3485 := Call(__e, PrimFunc(symshen_4scan_1body), V1550, V1554)


tmp3486 := PrimCons(symcond, tmp3485)

tmp3487 := PrimCons(tmp3486, Nil)

tmp3488 := PrimCons(symfreeze, tmp3487)

tmp3489 := PrimHead(V1553)

tmp3490 := PrimTail(V1553)

tmp3491 := PrimHead(tmp3490)

tmp3492 := PrimTail(tmp3491)

tmp3493 := PrimHead(tmp3492)

tmp3494 := PrimCons(symfail, Nil)

tmp3495 := PrimCons(tmp3494, Nil)

tmp3496 := PrimCons(V1552, tmp3495)

tmp3497 := PrimCons(sym_a, tmp3496)

tmp3498 := PrimCons(V1551, Nil)

tmp3499 := PrimCons(symthaw, tmp3498)

tmp3500 := PrimCons(V1552, Nil)

tmp3501 := PrimCons(tmp3499, tmp3500)

tmp3502 := PrimCons(tmp3497, tmp3501)

tmp3503 := PrimCons(symif, tmp3502)

tmp3504 := PrimCons(tmp3503, Nil)

tmp3505 := PrimCons(tmp3493, tmp3504)

tmp3506 := PrimCons(V1552, tmp3505)

tmp3507 := PrimCons(symlet, tmp3506)

tmp3508 := PrimCons(V1551, Nil)

tmp3509 := PrimCons(symthaw, tmp3508)

tmp3510 := PrimCons(tmp3509, Nil)

tmp3511 := PrimCons(tmp3507, tmp3510)

tmp3512 := PrimCons(tmp3489, tmp3511)

tmp3513 := PrimCons(symif, tmp3512)

tmp3514 := PrimCons(tmp3513, Nil)

tmp3515 := PrimCons(tmp3488, tmp3514)

tmp3516 := PrimCons(V1551, tmp3515)

tmp3517 := PrimCons(symlet, tmp3516)

tmp3518 := PrimCons(tmp3517, Nil)

tmp3519 := PrimCons(True, tmp3518)

__e.Return(PrimCons(tmp3519, Nil))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.choicepoint")))
return
}


}


}, 5)

tmp3629 := Call(__e, ns2_1set, symshen_4choicepoint, tmp3436)


_ = tmp3629

tmp3630 := MakeNative(func(__e *ControlFlow) {
V1556 := __e.Get(1)
_ = V1556
V1557 := __e.Get(2)
_ = V1557
V1558 := __e.Get(3)
_ = V1558
tmp3644 := PrimEqual(V1556, V1558)

if True == tmp3644 {
__e.Return(V1557)
return
} else {
tmp3642 := PrimIsPair(V1558)

if True == tmp3642 {
tmp3631 := MakeNative(func(__e *ControlFlow) {
Rep := __e.Get(1)
_ = Rep
tmp3637 := PrimHead(V1558)

tmp3638 := PrimEqual(Rep, tmp3637)

if True == tmp3638 {
tmp3632 := PrimHead(V1558)

tmp3633 := PrimTail(V1558)

tmp3634 := Call(__e, PrimFunc(symshen_4rep_1X), V1556, V1557, tmp3633)


__e.Return(PrimCons(tmp3632, tmp3634))
return


} else {
tmp3635 := PrimTail(V1558)

__e.Return(PrimCons(Rep, tmp3635))
return


}


}, 1)

tmp3639 := PrimHead(V1558)

tmp3640 := Call(__e, PrimFunc(symshen_4rep_1X), V1556, V1557, tmp3639)


__e.TailApply(tmp3631, tmp3640)
return


} else {
__e.Return(V1558)
return
}


}


}, 3)

tmp3645 := Call(__e, ns2_1set, symshen_4rep_1X, tmp3630)


_ = tmp3645

tmp3646 := MakeNative(func(__e *ControlFlow) {
V1559 := __e.Get(1)
_ = V1559
V1560 := __e.Get(2)
_ = V1560
tmp3647 := MakeNative(func(__e *ControlFlow) {
R := __e.Get(1)
_ = R
tmp3648 := Call(__e, PrimFunc(symfst), R)


tmp3649 := Call(__e, PrimFunc(symsnd), R)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp3648, V1560, tmp3649)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp3647, V1559)
return


}, 2)

tmp3650 := Call(__e, ns2_1set, symshen_4kl_1body, tmp3646)


_ = tmp3650

tmp3651 := MakeNative(func(__e *ControlFlow) {
V1569 := __e.Get(1)
_ = V1569
V1570 := __e.Get(2)
_ = V1570
V1571 := __e.Get(3)
_ = V1571
V1572 := __e.Get(4)
_ = V1572
tmp3781 := PrimEqual(Nil, V1570)

var ifres3756 Obj

if True == tmp3781 {
tmp3780 := PrimEqual(Nil, V1571)

var ifres3758 Obj

if True == tmp3780 {
tmp3779 := PrimIsPair(V1572)

var ifres3760 Obj

if True == tmp3779 {
tmp3777 := PrimHead(V1572)

tmp3778 := PrimEqual(symwhere, tmp3777)

var ifres3762 Obj

if True == tmp3778 {
tmp3775 := PrimTail(V1572)

tmp3776 := PrimIsPair(tmp3775)

var ifres3764 Obj

if True == tmp3776 {
tmp3772 := PrimTail(V1572)

tmp3773 := PrimTail(tmp3772)

tmp3774 := PrimIsPair(tmp3773)

var ifres3766 Obj

if True == tmp3774 {
tmp3768 := PrimTail(V1572)

tmp3769 := PrimTail(tmp3768)

tmp3770 := PrimTail(tmp3769)

tmp3771 := PrimEqual(Nil, tmp3770)

var ifres3767 Obj

if True == tmp3771 {
ifres3767 = True


} else {
ifres3767 = False


}

ifres3766 = ifres3767


} else {
ifres3766 = False


}

var ifres3765 Obj

if True == ifres3766 {
ifres3765 = True


} else {
ifres3765 = False


}

ifres3764 = ifres3765


} else {
ifres3764 = False


}

var ifres3763 Obj

if True == ifres3764 {
ifres3763 = True


} else {
ifres3763 = False


}

ifres3762 = ifres3763


} else {
ifres3762 = False


}

var ifres3761 Obj

if True == ifres3762 {
ifres3761 = True


} else {
ifres3761 = False


}

ifres3760 = ifres3761


} else {
ifres3760 = False


}

var ifres3759 Obj

if True == ifres3760 {
ifres3759 = True


} else {
ifres3759 = False


}

ifres3758 = ifres3759


} else {
ifres3758 = False


}

var ifres3757 Obj

if True == ifres3758 {
ifres3757 = True


} else {
ifres3757 = False


}

ifres3756 = ifres3757


} else {
ifres3756 = False


}

if True == ifres3756 {
tmp3652 := PrimTail(V1572)

tmp3653 := PrimHead(tmp3652)

tmp3654 := PrimCons(tmp3653, V1569)

tmp3655 := PrimTail(V1572)

tmp3656 := PrimTail(tmp3655)

tmp3657 := PrimHead(tmp3656)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3654, Nil, Nil, tmp3657)
return


} else {
tmp3754 := PrimEqual(Nil, V1570)

var ifres3751 Obj

if True == tmp3754 {
tmp3753 := PrimEqual(Nil, V1571)

var ifres3752 Obj

if True == tmp3753 {
ifres3752 = True


} else {
ifres3752 = False


}

ifres3751 = ifres3752


} else {
ifres3751 = False


}

if True == ifres3751 {
tmp3658 := Call(__e, PrimFunc(symreverse), V1569)


tmp3659 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp3658)


tmp3660 := PrimCons(V1572, Nil)

__e.Return(PrimCons(tmp3659, tmp3660))
return


} else {
tmp3749 := PrimIsPair(V1570)

var ifres3742 Obj

if True == tmp3749 {
tmp3748 := PrimIsPair(V1571)

var ifres3744 Obj

if True == tmp3748 {
tmp3746 := PrimHead(V1570)

tmp3747 := PrimIsVariable(tmp3746)

var ifres3745 Obj

if True == tmp3747 {
ifres3745 = True


} else {
ifres3745 = False


}

ifres3744 = ifres3745


} else {
ifres3744 = False


}

var ifres3743 Obj

if True == ifres3744 {
ifres3743 = True


} else {
ifres3743 = False


}

ifres3742 = ifres3743


} else {
ifres3742 = False


}

if True == ifres3742 {
tmp3661 := PrimTail(V1570)

tmp3662 := PrimTail(V1571)

tmp3663 := PrimHead(V1570)

tmp3664 := PrimHead(V1571)

tmp3665 := Call(__e, PrimFunc(symshen_4beta), tmp3663, tmp3664, V1572)


__e.TailApply(PrimFunc(symshen_4triple_1stack), V1569, tmp3661, tmp3662, tmp3665)
return


} else {
tmp3740 := PrimIsPair(V1570)

var ifres3715 Obj

if True == tmp3740 {
tmp3738 := PrimHead(V1570)

tmp3739 := PrimIsPair(tmp3738)

var ifres3717 Obj

if True == tmp3739 {
tmp3735 := PrimHead(V1570)

tmp3736 := PrimTail(tmp3735)

tmp3737 := PrimIsPair(tmp3736)

var ifres3719 Obj

if True == tmp3737 {
tmp3731 := PrimHead(V1570)

tmp3732 := PrimTail(tmp3731)

tmp3733 := PrimTail(tmp3732)

tmp3734 := PrimIsPair(tmp3733)

var ifres3721 Obj

if True == tmp3734 {
tmp3726 := PrimHead(V1570)

tmp3727 := PrimTail(tmp3726)

tmp3728 := PrimTail(tmp3727)

tmp3729 := PrimTail(tmp3728)

tmp3730 := PrimEqual(Nil, tmp3729)

var ifres3723 Obj

if True == tmp3730 {
tmp3725 := PrimIsPair(V1571)

var ifres3724 Obj

if True == tmp3725 {
ifres3724 = True


} else {
ifres3724 = False


}

ifres3723 = ifres3724


} else {
ifres3723 = False


}

var ifres3722 Obj

if True == ifres3723 {
ifres3722 = True


} else {
ifres3722 = False


}

ifres3721 = ifres3722


} else {
ifres3721 = False


}

var ifres3720 Obj

if True == ifres3721 {
ifres3720 = True


} else {
ifres3720 = False


}

ifres3719 = ifres3720


} else {
ifres3719 = False


}

var ifres3718 Obj

if True == ifres3719 {
ifres3718 = True


} else {
ifres3718 = False


}

ifres3717 = ifres3718


} else {
ifres3717 = False


}

var ifres3716 Obj

if True == ifres3717 {
ifres3716 = True


} else {
ifres3716 = False


}

ifres3715 = ifres3716


} else {
ifres3715 = False


}

if True == ifres3715 {
tmp3666 := PrimHead(V1570)

tmp3667 := PrimHead(tmp3666)

tmp3668 := Call(__e, PrimFunc(symshen_4op_1test), tmp3667)


tmp3669 := PrimHead(V1571)

tmp3670 := PrimCons(tmp3669, Nil)

tmp3671 := PrimCons(tmp3668, tmp3670)

tmp3672 := PrimCons(tmp3671, V1569)

tmp3673 := PrimHead(V1570)

tmp3674 := PrimTail(tmp3673)

tmp3675 := PrimHead(tmp3674)

tmp3676 := PrimHead(V1570)

tmp3677 := PrimTail(tmp3676)

tmp3678 := PrimTail(tmp3677)

tmp3679 := PrimHead(tmp3678)

tmp3680 := PrimTail(V1570)

tmp3681 := PrimCons(tmp3679, tmp3680)

tmp3682 := PrimCons(tmp3675, tmp3681)

tmp3683 := PrimHead(V1570)

tmp3684 := PrimHead(tmp3683)

tmp3685 := Call(__e, PrimFunc(symshen_4op1), tmp3684)


tmp3686 := PrimHead(V1571)

tmp3687 := PrimCons(tmp3686, Nil)

tmp3688 := PrimCons(tmp3685, tmp3687)

tmp3689 := PrimHead(V1570)

tmp3690 := PrimHead(tmp3689)

tmp3691 := Call(__e, PrimFunc(symshen_4op2), tmp3690)


tmp3692 := PrimHead(V1571)

tmp3693 := PrimCons(tmp3692, Nil)

tmp3694 := PrimCons(tmp3691, tmp3693)

tmp3695 := PrimTail(V1571)

tmp3696 := PrimCons(tmp3694, tmp3695)

tmp3697 := PrimCons(tmp3688, tmp3696)

tmp3698 := PrimHead(V1570)

tmp3699 := PrimHead(V1571)

tmp3700 := Call(__e, PrimFunc(symshen_4beta), tmp3698, tmp3699, V1572)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3672, tmp3682, tmp3697, tmp3700)
return


} else {
tmp3713 := PrimIsPair(V1570)

var ifres3710 Obj

if True == tmp3713 {
tmp3712 := PrimIsPair(V1571)

var ifres3711 Obj

if True == tmp3712 {
ifres3711 = True


} else {
ifres3711 = False


}

ifres3710 = ifres3711


} else {
ifres3710 = False


}

if True == ifres3710 {
tmp3701 := PrimHead(V1570)

tmp3702 := PrimHead(V1571)

tmp3703 := PrimCons(tmp3702, Nil)

tmp3704 := PrimCons(tmp3701, tmp3703)

tmp3705 := PrimCons(sym_a, tmp3704)

tmp3706 := PrimCons(tmp3705, V1569)

tmp3707 := PrimTail(V1570)

tmp3708 := PrimTail(V1571)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3706, tmp3707, tmp3708, V1572)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.triple-stack")))
return
}


}


}


}


}


}, 4)

tmp3782 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp3651)


_ = tmp3782

tmp3783 := MakeNative(func(__e *ControlFlow) {
V1575 := __e.Get(1)
_ = V1575
tmp3802 := PrimEqual(Nil, V1575)

if True == tmp3802 {
__e.Return(True)
return
} else {
tmp3800 := PrimIsPair(V1575)

var ifres3796 Obj

if True == tmp3800 {
tmp3798 := PrimTail(V1575)

tmp3799 := PrimEqual(Nil, tmp3798)

var ifres3797 Obj

if True == tmp3799 {
ifres3797 = True


} else {
ifres3797 = False


}

ifres3796 = ifres3797


} else {
ifres3796 = False


}

if True == ifres3796 {
__e.Return(PrimHead(V1575))
return
} else {
tmp3794 := PrimIsPair(V1575)

var ifres3790 Obj

if True == tmp3794 {
tmp3792 := PrimTail(V1575)

tmp3793 := PrimIsPair(tmp3792)

var ifres3791 Obj

if True == tmp3793 {
ifres3791 = True


} else {
ifres3791 = False


}

ifres3790 = ifres3791


} else {
ifres3790 = False


}

if True == ifres3790 {
tmp3784 := PrimHead(V1575)

tmp3785 := PrimTail(V1575)

tmp3786 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp3785)


tmp3787 := PrimCons(tmp3786, Nil)

tmp3788 := PrimCons(tmp3784, tmp3787)

__e.Return(PrimCons(symand, tmp3788))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.rectify-test")))
return
}


}


}


}, 1)

tmp3803 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp3783)


_ = tmp3803

tmp3804 := MakeNative(func(__e *ControlFlow) {
V1585 := __e.Get(1)
_ = V1585
V1586 := __e.Get(2)
_ = V1586
V1587 := __e.Get(3)
_ = V1587
tmp3881 := PrimEqual(V1585, V1587)

if True == tmp3881 {
__e.Return(V1586)
return
} else {
tmp3879 := PrimIsPair(V1587)

var ifres3855 Obj

if True == tmp3879 {
tmp3877 := PrimHead(V1587)

tmp3878 := PrimEqual(symlambda, tmp3877)

var ifres3857 Obj

if True == tmp3878 {
tmp3875 := PrimTail(V1587)

tmp3876 := PrimIsPair(tmp3875)

var ifres3859 Obj

if True == tmp3876 {
tmp3872 := PrimTail(V1587)

tmp3873 := PrimTail(tmp3872)

tmp3874 := PrimIsPair(tmp3873)

var ifres3861 Obj

if True == tmp3874 {
tmp3868 := PrimTail(V1587)

tmp3869 := PrimTail(tmp3868)

tmp3870 := PrimTail(tmp3869)

tmp3871 := PrimEqual(Nil, tmp3870)

var ifres3863 Obj

if True == tmp3871 {
tmp3865 := PrimTail(V1587)

tmp3866 := PrimHead(tmp3865)

tmp3867 := PrimEqual(V1585, tmp3866)

var ifres3864 Obj

if True == tmp3867 {
ifres3864 = True


} else {
ifres3864 = False


}

ifres3863 = ifres3864


} else {
ifres3863 = False


}

var ifres3862 Obj

if True == ifres3863 {
ifres3862 = True


} else {
ifres3862 = False


}

ifres3861 = ifres3862


} else {
ifres3861 = False


}

var ifres3860 Obj

if True == ifres3861 {
ifres3860 = True


} else {
ifres3860 = False


}

ifres3859 = ifres3860


} else {
ifres3859 = False


}

var ifres3858 Obj

if True == ifres3859 {
ifres3858 = True


} else {
ifres3858 = False


}

ifres3857 = ifres3858


} else {
ifres3857 = False


}

var ifres3856 Obj

if True == ifres3857 {
ifres3856 = True


} else {
ifres3856 = False


}

ifres3855 = ifres3856


} else {
ifres3855 = False


}

if True == ifres3855 {
__e.Return(V1587)
return
} else {
tmp3853 := PrimIsPair(V1587)

var ifres3822 Obj

if True == tmp3853 {
tmp3851 := PrimHead(V1587)

tmp3852 := PrimEqual(symlet, tmp3851)

var ifres3824 Obj

if True == tmp3852 {
tmp3849 := PrimTail(V1587)

tmp3850 := PrimIsPair(tmp3849)

var ifres3826 Obj

if True == tmp3850 {
tmp3846 := PrimTail(V1587)

tmp3847 := PrimTail(tmp3846)

tmp3848 := PrimIsPair(tmp3847)

var ifres3828 Obj

if True == tmp3848 {
tmp3842 := PrimTail(V1587)

tmp3843 := PrimTail(tmp3842)

tmp3844 := PrimTail(tmp3843)

tmp3845 := PrimIsPair(tmp3844)

var ifres3830 Obj

if True == tmp3845 {
tmp3837 := PrimTail(V1587)

tmp3838 := PrimTail(tmp3837)

tmp3839 := PrimTail(tmp3838)

tmp3840 := PrimTail(tmp3839)

tmp3841 := PrimEqual(Nil, tmp3840)

var ifres3832 Obj

if True == tmp3841 {
tmp3834 := PrimTail(V1587)

tmp3835 := PrimHead(tmp3834)

tmp3836 := PrimEqual(V1585, tmp3835)

var ifres3833 Obj

if True == tmp3836 {
ifres3833 = True


} else {
ifres3833 = False


}

ifres3832 = ifres3833


} else {
ifres3832 = False


}

var ifres3831 Obj

if True == ifres3832 {
ifres3831 = True


} else {
ifres3831 = False


}

ifres3830 = ifres3831


} else {
ifres3830 = False


}

var ifres3829 Obj

if True == ifres3830 {
ifres3829 = True


} else {
ifres3829 = False


}

ifres3828 = ifres3829


} else {
ifres3828 = False


}

var ifres3827 Obj

if True == ifres3828 {
ifres3827 = True


} else {
ifres3827 = False


}

ifres3826 = ifres3827


} else {
ifres3826 = False


}

var ifres3825 Obj

if True == ifres3826 {
ifres3825 = True


} else {
ifres3825 = False


}

ifres3824 = ifres3825


} else {
ifres3824 = False


}

var ifres3823 Obj

if True == ifres3824 {
ifres3823 = True


} else {
ifres3823 = False


}

ifres3822 = ifres3823


} else {
ifres3822 = False


}

if True == ifres3822 {
tmp3805 := PrimTail(V1587)

tmp3806 := PrimHead(tmp3805)

tmp3807 := PrimTail(V1587)

tmp3808 := PrimHead(tmp3807)

tmp3809 := PrimTail(V1587)

tmp3810 := PrimTail(tmp3809)

tmp3811 := PrimHead(tmp3810)

tmp3812 := Call(__e, PrimFunc(symshen_4beta), tmp3808, V1586, tmp3811)


tmp3813 := PrimTail(V1587)

tmp3814 := PrimTail(tmp3813)

tmp3815 := PrimTail(tmp3814)

tmp3816 := PrimCons(tmp3812, tmp3815)

tmp3817 := PrimCons(tmp3806, tmp3816)

__e.Return(PrimCons(symlet, tmp3817))
return


} else {
tmp3820 := PrimIsPair(V1587)

if True == tmp3820 {
tmp3818 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
__e.TailApply(PrimFunc(symshen_4beta), V1585, V1586, V)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3818, V1587)
return


} else {
__e.Return(V1587)
return
}


}


}


}


}, 3)

tmp3882 := Call(__e, ns2_1set, symshen_4beta, tmp3804)


_ = tmp3882

tmp3883 := MakeNative(func(__e *ControlFlow) {
V1590 := __e.Get(1)
_ = V1590
tmp3891 := PrimEqual(symcons, V1590)

if True == tmp3891 {
__e.Return(symhd)
return
} else {
tmp3889 := PrimEqual(sym_8s, V1590)

if True == tmp3889 {
__e.Return(symhdstr)
return
} else {
tmp3887 := PrimEqual(sym_8p, V1590)

if True == tmp3887 {
__e.Return(symfst)
return
} else {
tmp3885 := PrimEqual(sym_8v, V1590)

if True == tmp3885 {
__e.Return(symhdv)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.op1")))
return
}


}


}


}


}, 1)

tmp3892 := Call(__e, ns2_1set, symshen_4op1, tmp3883)


_ = tmp3892

tmp3893 := MakeNative(func(__e *ControlFlow) {
V1593 := __e.Get(1)
_ = V1593
tmp3901 := PrimEqual(symcons, V1593)

if True == tmp3901 {
__e.Return(symtl)
return
} else {
tmp3899 := PrimEqual(sym_8s, V1593)

if True == tmp3899 {
__e.Return(symtlstr)
return
} else {
tmp3897 := PrimEqual(sym_8p, V1593)

if True == tmp3897 {
__e.Return(symsnd)
return
} else {
tmp3895 := PrimEqual(sym_8v, V1593)

if True == tmp3895 {
__e.Return(symtlv)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.op2")))
return
}


}


}


}


}, 1)

tmp3902 := Call(__e, ns2_1set, symshen_4op2, tmp3893)


_ = tmp3902

tmp3903 := MakeNative(func(__e *ControlFlow) {
V1596 := __e.Get(1)
_ = V1596
tmp3911 := PrimEqual(symcons, V1596)

if True == tmp3911 {
__e.Return(symcons_2)
return
} else {
tmp3909 := PrimEqual(sym_8s, V1596)

if True == tmp3909 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp3907 := PrimEqual(sym_8p, V1596)

if True == tmp3907 {
__e.Return(symtuple_2)
return
} else {
tmp3905 := PrimEqual(sym_8v, V1596)

if True == tmp3905 {
__e.Return(symshen_4_7vector_2)
return
} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.op-test")))
return
}


}


}


}


}, 1)

tmp3912 := Call(__e, ns2_1set, symshen_4op_1test, tmp3903)


_ = tmp3912

tmp3913 := MakeNative(func(__e *ControlFlow) {
V1597 := __e.Get(1)
_ = V1597
tmp3915 := PrimEqual(MakeString(""), V1597)

if True == tmp3915 {
__e.Return(False)
return
} else {
__e.Return(PrimIsString(V1597))
return
}


}, 1)

tmp3916 := Call(__e, ns2_1set, symshen_4_7string_2, tmp3913)


_ = tmp3916

tmp3917 := MakeNative(func(__e *ControlFlow) {
V1598 := __e.Get(1)
_ = V1598
tmp3919 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp3920 := PrimEqual(V1598, tmp3919)

if True == tmp3920 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V1598)
return
}


}, 1)

tmp3921 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp3917)


_ = tmp3921

tmp3922 := MakeNative(func(__e *ControlFlow) {
V1601 := __e.Get(1)
_ = V1601
tmp3926 := PrimEqual(sym_7, V1601)

if True == tmp3926 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, True))
return
} else {
tmp3924 := PrimEqual(sym_1, V1601)

if True == tmp3924 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("factorise expects a + or a -\n")))
return
}


}


}, 1)

tmp3927 := Call(__e, ns2_1set, symfactorise, tmp3922)


_ = tmp3927

tmp3928 := MakeNative(func(__e *ControlFlow) {
V1602 := __e.Get(1)
_ = V1602
tmp3995 := PrimIsPair(V1602)

var ifres3951 Obj

if True == tmp3995 {
tmp3993 := PrimHead(V1602)

tmp3994 := PrimEqual(symdefun, tmp3993)

var ifres3953 Obj

if True == tmp3994 {
tmp3991 := PrimTail(V1602)

tmp3992 := PrimIsPair(tmp3991)

var ifres3955 Obj

if True == tmp3992 {
tmp3988 := PrimTail(V1602)

tmp3989 := PrimTail(tmp3988)

tmp3990 := PrimIsPair(tmp3989)

var ifres3957 Obj

if True == tmp3990 {
tmp3984 := PrimTail(V1602)

tmp3985 := PrimTail(tmp3984)

tmp3986 := PrimTail(tmp3985)

tmp3987 := PrimIsPair(tmp3986)

var ifres3959 Obj

if True == tmp3987 {
tmp3979 := PrimTail(V1602)

tmp3980 := PrimTail(tmp3979)

tmp3981 := PrimTail(tmp3980)

tmp3982 := PrimHead(tmp3981)

tmp3983 := PrimIsPair(tmp3982)

var ifres3961 Obj

if True == tmp3983 {
tmp3973 := PrimTail(V1602)

tmp3974 := PrimTail(tmp3973)

tmp3975 := PrimTail(tmp3974)

tmp3976 := PrimHead(tmp3975)

tmp3977 := PrimHead(tmp3976)

tmp3978 := PrimEqual(symcond, tmp3977)

var ifres3963 Obj

if True == tmp3978 {
tmp3968 := PrimTail(V1602)

tmp3969 := PrimTail(tmp3968)

tmp3970 := PrimTail(tmp3969)

tmp3971 := PrimTail(tmp3970)

tmp3972 := PrimEqual(Nil, tmp3971)

var ifres3965 Obj

if True == tmp3972 {
tmp3967 := PrimValue(symshen_4_dfactorise_2_d)

var ifres3966 Obj

if True == tmp3967 {
ifres3966 = True


} else {
ifres3966 = False


}

ifres3965 = ifres3966


} else {
ifres3965 = False


}

var ifres3964 Obj

if True == ifres3965 {
ifres3964 = True


} else {
ifres3964 = False


}

ifres3963 = ifres3964


} else {
ifres3963 = False


}

var ifres3962 Obj

if True == ifres3963 {
ifres3962 = True


} else {
ifres3962 = False


}

ifres3961 = ifres3962


} else {
ifres3961 = False


}

var ifres3960 Obj

if True == ifres3961 {
ifres3960 = True


} else {
ifres3960 = False


}

ifres3959 = ifres3960


} else {
ifres3959 = False


}

var ifres3958 Obj

if True == ifres3959 {
ifres3958 = True


} else {
ifres3958 = False


}

ifres3957 = ifres3958


} else {
ifres3957 = False


}

var ifres3956 Obj

if True == ifres3957 {
ifres3956 = True


} else {
ifres3956 = False


}

ifres3955 = ifres3956


} else {
ifres3955 = False


}

var ifres3954 Obj

if True == ifres3955 {
ifres3954 = True


} else {
ifres3954 = False


}

ifres3953 = ifres3954


} else {
ifres3953 = False


}

var ifres3952 Obj

if True == ifres3953 {
ifres3952 = True


} else {
ifres3952 = False


}

ifres3951 = ifres3952


} else {
ifres3951 = False


}

if True == ifres3951 {
tmp3929 := PrimTail(V1602)

tmp3930 := PrimHead(tmp3929)

tmp3931 := PrimTail(V1602)

tmp3932 := PrimTail(tmp3931)

tmp3933 := PrimHead(tmp3932)

tmp3934 := PrimTail(V1602)

tmp3935 := PrimTail(tmp3934)

tmp3936 := PrimHead(tmp3935)

tmp3937 := PrimTail(V1602)

tmp3938 := PrimTail(tmp3937)

tmp3939 := PrimTail(tmp3938)

tmp3940 := PrimHead(tmp3939)

tmp3941 := PrimTail(tmp3940)

tmp3942 := PrimTail(V1602)

tmp3943 := PrimHead(tmp3942)

tmp3944 := PrimCons(tmp3943, Nil)

tmp3945 := PrimCons(symshen_4f_1error, tmp3944)

tmp3946 := Call(__e, PrimFunc(symshen_4vertical), tmp3936, tmp3941, tmp3945)


tmp3947 := PrimCons(tmp3946, Nil)

tmp3948 := PrimCons(tmp3933, tmp3947)

tmp3949 := PrimCons(tmp3930, tmp3948)

__e.Return(PrimCons(symdefun, tmp3949))
return


} else {
__e.Return(V1602)
return
}


}, 1)

tmp3996 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp3928)


_ = tmp3996

tmp3997 := MakeNative(func(__e *ControlFlow) {
V1615 := __e.Get(1)
_ = V1615
V1616 := __e.Get(2)
_ = V1616
V1617 := __e.Get(3)
_ = V1617
tmp4109 := PrimIsPair(V1616)

var ifres4089 Obj

if True == tmp4109 {
tmp4107 := PrimHead(V1616)

tmp4108 := PrimIsPair(tmp4107)

var ifres4091 Obj

if True == tmp4108 {
tmp4104 := PrimHead(V1616)

tmp4105 := PrimHead(tmp4104)

tmp4106 := PrimEqual(True, tmp4105)

var ifres4093 Obj

if True == tmp4106 {
tmp4101 := PrimHead(V1616)

tmp4102 := PrimTail(tmp4101)

tmp4103 := PrimIsPair(tmp4102)

var ifres4095 Obj

if True == tmp4103 {
tmp4097 := PrimHead(V1616)

tmp4098 := PrimTail(tmp4097)

tmp4099 := PrimTail(tmp4098)

tmp4100 := PrimEqual(Nil, tmp4099)

var ifres4096 Obj

if True == tmp4100 {
ifres4096 = True


} else {
ifres4096 = False


}

ifres4095 = ifres4096


} else {
ifres4095 = False


}

var ifres4094 Obj

if True == ifres4095 {
ifres4094 = True


} else {
ifres4094 = False


}

ifres4093 = ifres4094


} else {
ifres4093 = False


}

var ifres4092 Obj

if True == ifres4093 {
ifres4092 = True


} else {
ifres4092 = False


}

ifres4091 = ifres4092


} else {
ifres4091 = False


}

var ifres4090 Obj

if True == ifres4091 {
ifres4090 = True


} else {
ifres4090 = False


}

ifres4089 = ifres4090


} else {
ifres4089 = False


}

if True == ifres4089 {
tmp3998 := PrimHead(V1616)

tmp3999 := PrimTail(tmp3998)

__e.Return(PrimHead(tmp3999))
return


} else {
tmp4087 := PrimEqual(Nil, V1616)

if True == tmp4087 {
__e.Return(V1617)
return
} else {
tmp4085 := PrimIsPair(V1616)

var ifres4038 Obj

if True == tmp4085 {
tmp4083 := PrimHead(V1616)

tmp4084 := PrimIsPair(tmp4083)

var ifres4040 Obj

if True == tmp4084 {
tmp4080 := PrimHead(V1616)

tmp4081 := PrimHead(tmp4080)

tmp4082 := PrimIsPair(tmp4081)

var ifres4042 Obj

if True == tmp4082 {
tmp4076 := PrimHead(V1616)

tmp4077 := PrimHead(tmp4076)

tmp4078 := PrimHead(tmp4077)

tmp4079 := PrimEqual(symand, tmp4078)

var ifres4044 Obj

if True == tmp4079 {
tmp4072 := PrimHead(V1616)

tmp4073 := PrimHead(tmp4072)

tmp4074 := PrimTail(tmp4073)

tmp4075 := PrimIsPair(tmp4074)

var ifres4046 Obj

if True == tmp4075 {
tmp4067 := PrimHead(V1616)

tmp4068 := PrimHead(tmp4067)

tmp4069 := PrimTail(tmp4068)

tmp4070 := PrimTail(tmp4069)

tmp4071 := PrimIsPair(tmp4070)

var ifres4048 Obj

if True == tmp4071 {
tmp4061 := PrimHead(V1616)

tmp4062 := PrimHead(tmp4061)

tmp4063 := PrimTail(tmp4062)

tmp4064 := PrimTail(tmp4063)

tmp4065 := PrimTail(tmp4064)

tmp4066 := PrimEqual(Nil, tmp4065)

var ifres4050 Obj

if True == tmp4066 {
tmp4058 := PrimHead(V1616)

tmp4059 := PrimTail(tmp4058)

tmp4060 := PrimIsPair(tmp4059)

var ifres4052 Obj

if True == tmp4060 {
tmp4054 := PrimHead(V1616)

tmp4055 := PrimTail(tmp4054)

tmp4056 := PrimTail(tmp4055)

tmp4057 := PrimEqual(Nil, tmp4056)

var ifres4053 Obj

if True == tmp4057 {
ifres4053 = True


} else {
ifres4053 = False


}

ifres4052 = ifres4053


} else {
ifres4052 = False


}

var ifres4051 Obj

if True == ifres4052 {
ifres4051 = True


} else {
ifres4051 = False


}

ifres4050 = ifres4051


} else {
ifres4050 = False


}

var ifres4049 Obj

if True == ifres4050 {
ifres4049 = True


} else {
ifres4049 = False


}

ifres4048 = ifres4049


} else {
ifres4048 = False


}

var ifres4047 Obj

if True == ifres4048 {
ifres4047 = True


} else {
ifres4047 = False


}

ifres4046 = ifres4047


} else {
ifres4046 = False


}

var ifres4045 Obj

if True == ifres4046 {
ifres4045 = True


} else {
ifres4045 = False


}

ifres4044 = ifres4045


} else {
ifres4044 = False


}

var ifres4043 Obj

if True == ifres4044 {
ifres4043 = True


} else {
ifres4043 = False


}

ifres4042 = ifres4043


} else {
ifres4042 = False


}

var ifres4041 Obj

if True == ifres4042 {
ifres4041 = True


} else {
ifres4041 = False


}

ifres4040 = ifres4041


} else {
ifres4040 = False


}

var ifres4039 Obj

if True == ifres4040 {
ifres4039 = True


} else {
ifres4039 = False


}

ifres4038 = ifres4039


} else {
ifres4038 = False


}

if True == ifres4038 {
tmp4000 := MakeNative(func(__e *ControlFlow) {
Before_7After := __e.Get(1)
_ = Before_7After
tmp4001 := PrimHead(V1616)

tmp4002 := PrimHead(tmp4001)

tmp4003 := PrimTail(tmp4002)

tmp4004 := PrimHead(tmp4003)

__e.TailApply(PrimFunc(symshen_4branch), tmp4004, V1615, Before_7After, V1617)
return


}, 1)

tmp4005 := PrimHead(V1616)

tmp4006 := PrimHead(tmp4005)

tmp4007 := PrimTail(tmp4006)

tmp4008 := PrimHead(tmp4007)

tmp4009 := Call(__e, PrimFunc(symshen_4split_1cases), tmp4008, V1616, Nil)


__e.TailApply(tmp4000, tmp4009)
return


} else {
tmp4036 := PrimIsPair(V1616)

var ifres4021 Obj

if True == tmp4036 {
tmp4034 := PrimHead(V1616)

tmp4035 := PrimIsPair(tmp4034)

var ifres4023 Obj

if True == tmp4035 {
tmp4031 := PrimHead(V1616)

tmp4032 := PrimTail(tmp4031)

tmp4033 := PrimIsPair(tmp4032)

var ifres4025 Obj

if True == tmp4033 {
tmp4027 := PrimHead(V1616)

tmp4028 := PrimTail(tmp4027)

tmp4029 := PrimTail(tmp4028)

tmp4030 := PrimEqual(Nil, tmp4029)

var ifres4026 Obj

if True == tmp4030 {
ifres4026 = True


} else {
ifres4026 = False


}

ifres4025 = ifres4026


} else {
ifres4025 = False


}

var ifres4024 Obj

if True == ifres4025 {
ifres4024 = True


} else {
ifres4024 = False


}

ifres4023 = ifres4024


} else {
ifres4023 = False


}

var ifres4022 Obj

if True == ifres4023 {
ifres4022 = True


} else {
ifres4022 = False


}

ifres4021 = ifres4022


} else {
ifres4021 = False


}

if True == ifres4021 {
tmp4010 := PrimHead(V1616)

tmp4011 := PrimHead(tmp4010)

tmp4012 := PrimHead(V1616)

tmp4013 := PrimTail(tmp4012)

tmp4014 := PrimHead(tmp4013)

tmp4015 := PrimTail(V1616)

tmp4016 := Call(__e, PrimFunc(symshen_4vertical), V1615, tmp4015, V1617)


tmp4017 := PrimCons(tmp4016, Nil)

tmp4018 := PrimCons(tmp4014, tmp4017)

tmp4019 := PrimCons(tmp4011, tmp4018)

__e.Return(PrimCons(symif, tmp4019))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.vertical")))
return
}


}


}


}


}, 3)

tmp4110 := Call(__e, ns2_1set, symshen_4vertical, tmp3997)


_ = tmp4110

tmp4111 := MakeNative(func(__e *ControlFlow) {
V1622 := __e.Get(1)
_ = V1622
V1623 := __e.Get(2)
_ = V1623
V1624 := __e.Get(3)
_ = V1624
tmp4212 := PrimIsPair(V1623)

var ifres4158 Obj

if True == tmp4212 {
tmp4210 := PrimHead(V1623)

tmp4211 := PrimIsPair(tmp4210)

var ifres4160 Obj

if True == tmp4211 {
tmp4207 := PrimHead(V1623)

tmp4208 := PrimHead(tmp4207)

tmp4209 := PrimIsPair(tmp4208)

var ifres4162 Obj

if True == tmp4209 {
tmp4203 := PrimHead(V1623)

tmp4204 := PrimHead(tmp4203)

tmp4205 := PrimHead(tmp4204)

tmp4206 := PrimEqual(symand, tmp4205)

var ifres4164 Obj

if True == tmp4206 {
tmp4199 := PrimHead(V1623)

tmp4200 := PrimHead(tmp4199)

tmp4201 := PrimTail(tmp4200)

tmp4202 := PrimIsPair(tmp4201)

var ifres4166 Obj

if True == tmp4202 {
tmp4194 := PrimHead(V1623)

tmp4195 := PrimHead(tmp4194)

tmp4196 := PrimTail(tmp4195)

tmp4197 := PrimTail(tmp4196)

tmp4198 := PrimIsPair(tmp4197)

var ifres4168 Obj

if True == tmp4198 {
tmp4188 := PrimHead(V1623)

tmp4189 := PrimHead(tmp4188)

tmp4190 := PrimTail(tmp4189)

tmp4191 := PrimTail(tmp4190)

tmp4192 := PrimTail(tmp4191)

tmp4193 := PrimEqual(Nil, tmp4192)

var ifres4170 Obj

if True == tmp4193 {
tmp4185 := PrimHead(V1623)

tmp4186 := PrimTail(tmp4185)

tmp4187 := PrimIsPair(tmp4186)

var ifres4172 Obj

if True == tmp4187 {
tmp4181 := PrimHead(V1623)

tmp4182 := PrimTail(tmp4181)

tmp4183 := PrimTail(tmp4182)

tmp4184 := PrimEqual(Nil, tmp4183)

var ifres4174 Obj

if True == tmp4184 {
tmp4176 := PrimHead(V1623)

tmp4177 := PrimHead(tmp4176)

tmp4178 := PrimTail(tmp4177)

tmp4179 := PrimHead(tmp4178)

tmp4180 := PrimEqual(V1622, tmp4179)

var ifres4175 Obj

if True == tmp4180 {
ifres4175 = True


} else {
ifres4175 = False


}

ifres4174 = ifres4175


} else {
ifres4174 = False


}

var ifres4173 Obj

if True == ifres4174 {
ifres4173 = True


} else {
ifres4173 = False


}

ifres4172 = ifres4173


} else {
ifres4172 = False


}

var ifres4171 Obj

if True == ifres4172 {
ifres4171 = True


} else {
ifres4171 = False


}

ifres4170 = ifres4171


} else {
ifres4170 = False


}

var ifres4169 Obj

if True == ifres4170 {
ifres4169 = True


} else {
ifres4169 = False


}

ifres4168 = ifres4169


} else {
ifres4168 = False


}

var ifres4167 Obj

if True == ifres4168 {
ifres4167 = True


} else {
ifres4167 = False


}

ifres4166 = ifres4167


} else {
ifres4166 = False


}

var ifres4165 Obj

if True == ifres4166 {
ifres4165 = True


} else {
ifres4165 = False


}

ifres4164 = ifres4165


} else {
ifres4164 = False


}

var ifres4163 Obj

if True == ifres4164 {
ifres4163 = True


} else {
ifres4163 = False


}

ifres4162 = ifres4163


} else {
ifres4162 = False


}

var ifres4161 Obj

if True == ifres4162 {
ifres4161 = True


} else {
ifres4161 = False


}

ifres4160 = ifres4161


} else {
ifres4160 = False


}

var ifres4159 Obj

if True == ifres4160 {
ifres4159 = True


} else {
ifres4159 = False


}

ifres4158 = ifres4159


} else {
ifres4158 = False


}

if True == ifres4158 {
tmp4112 := PrimHead(V1623)

tmp4113 := PrimHead(tmp4112)

tmp4114 := PrimTail(tmp4113)

tmp4115 := PrimHead(tmp4114)

tmp4116 := PrimTail(V1623)

tmp4117 := PrimHead(V1623)

tmp4118 := PrimHead(tmp4117)

tmp4119 := PrimTail(tmp4118)

tmp4120 := PrimTail(tmp4119)

tmp4121 := PrimHead(tmp4120)

tmp4122 := PrimHead(V1623)

tmp4123 := PrimTail(tmp4122)

tmp4124 := PrimCons(tmp4121, tmp4123)

tmp4125 := PrimCons(tmp4124, V1624)

__e.TailApply(PrimFunc(symshen_4split_1cases), tmp4115, tmp4116, tmp4125)
return


} else {
tmp4156 := PrimIsPair(V1623)

var ifres4136 Obj

if True == tmp4156 {
tmp4154 := PrimHead(V1623)

tmp4155 := PrimIsPair(tmp4154)

var ifres4138 Obj

if True == tmp4155 {
tmp4151 := PrimHead(V1623)

tmp4152 := PrimTail(tmp4151)

tmp4153 := PrimIsPair(tmp4152)

var ifres4140 Obj

if True == tmp4153 {
tmp4147 := PrimHead(V1623)

tmp4148 := PrimTail(tmp4147)

tmp4149 := PrimTail(tmp4148)

tmp4150 := PrimEqual(Nil, tmp4149)

var ifres4142 Obj

if True == tmp4150 {
tmp4144 := PrimHead(V1623)

tmp4145 := PrimHead(tmp4144)

tmp4146 := PrimEqual(V1622, tmp4145)

var ifres4143 Obj

if True == tmp4146 {
ifres4143 = True


} else {
ifres4143 = False


}

ifres4142 = ifres4143


} else {
ifres4142 = False


}

var ifres4141 Obj

if True == ifres4142 {
ifres4141 = True


} else {
ifres4141 = False


}

ifres4140 = ifres4141


} else {
ifres4140 = False


}

var ifres4139 Obj

if True == ifres4140 {
ifres4139 = True


} else {
ifres4139 = False


}

ifres4138 = ifres4139


} else {
ifres4138 = False


}

var ifres4137 Obj

if True == ifres4138 {
ifres4137 = True


} else {
ifres4137 = False


}

ifres4136 = ifres4137


} else {
ifres4136 = False


}

if True == ifres4136 {
tmp4126 := PrimHead(V1623)

tmp4127 := PrimHead(tmp4126)

tmp4128 := PrimTail(V1623)

tmp4129 := PrimHead(V1623)

tmp4130 := PrimTail(tmp4129)

tmp4131 := PrimCons(True, tmp4130)

tmp4132 := PrimCons(tmp4131, V1624)

__e.TailApply(PrimFunc(symshen_4split_1cases), tmp4127, tmp4128, tmp4132)
return


} else {
tmp4133 := Call(__e, PrimFunc(symreverse), V1624)


tmp4134 := PrimCons(V1623, Nil)

__e.Return(PrimCons(tmp4133, tmp4134))
return


}


}


}, 3)

tmp4213 := Call(__e, ns2_1set, symshen_4split_1cases, tmp4111)


_ = tmp4213

tmp4214 := MakeNative(func(__e *ControlFlow) {
V1625 := __e.Get(1)
_ = V1625
V1626 := __e.Get(2)
_ = V1626
V1627 := __e.Get(3)
_ = V1627
V1628 := __e.Get(4)
_ = V1628
tmp4235 := PrimIsPair(V1627)

var ifres4226 Obj

if True == tmp4235 {
tmp4233 := PrimTail(V1627)

tmp4234 := PrimIsPair(tmp4233)

var ifres4228 Obj

if True == tmp4234 {
tmp4230 := PrimTail(V1627)

tmp4231 := PrimTail(tmp4230)

tmp4232 := PrimEqual(Nil, tmp4231)

var ifres4229 Obj

if True == tmp4232 {
ifres4229 = True


} else {
ifres4229 = False


}

ifres4228 = ifres4229


} else {
ifres4228 = False


}

var ifres4227 Obj

if True == ifres4228 {
ifres4227 = True


} else {
ifres4227 = False


}

ifres4226 = ifres4227


} else {
ifres4226 = False


}

if True == ifres4226 {
tmp4215 := MakeNative(func(__e *ControlFlow) {
Else := __e.Get(1)
_ = Else
tmp4216 := MakeNative(func(__e *ControlFlow) {
Then := __e.Get(1)
_ = Then
tmp4217 := PrimCons(Else, Nil)

tmp4218 := PrimCons(Then, tmp4217)

tmp4219 := PrimCons(V1625, tmp4218)

__e.Return(PrimCons(symif, tmp4219))
return


}, 1)

tmp4220 := PrimHead(V1627)

tmp4221 := Call(__e, PrimFunc(symshen_4then), V1625, V1626, tmp4220, Else)


__e.TailApply(tmp4216, tmp4221)
return


}, 1)

tmp4222 := PrimTail(V1627)

tmp4223 := PrimHead(tmp4222)

tmp4224 := Call(__e, PrimFunc(symshen_4else), V1626, tmp4223, V1628)


__e.TailApply(tmp4215, tmp4224)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4branch)
return
}


}, 4)

tmp4236 := Call(__e, ns2_1set, symshen_4branch, tmp4214)


_ = tmp4236

tmp4237 := MakeNative(func(__e *ControlFlow) {
V1629 := __e.Get(1)
_ = V1629
V1630 := __e.Get(2)
_ = V1630
V1631 := __e.Get(3)
_ = V1631
tmp4238 := MakeNative(func(__e *ControlFlow) {
Else := __e.Get(1)
_ = Else
tmp4240 := Call(__e, PrimFunc(symshen_4inline_2), Else)


if True == tmp4240 {
__e.Return(Else)
return
} else {
__e.TailApply(PrimFunc(symshen_4procedure_1call), V1629, Else)
return
}


}, 1)

tmp4241 := Call(__e, PrimFunc(symshen_4vertical), V1629, V1630, V1631)


__e.TailApply(tmp4238, tmp4241)
return


}, 3)

tmp4242 := Call(__e, ns2_1set, symshen_4else, tmp4237)


_ = tmp4242

tmp4243 := MakeNative(func(__e *ControlFlow) {
V1632 := __e.Get(1)
_ = V1632
V1633 := __e.Get(2)
_ = V1633
tmp4244 := MakeNative(func(__e *ControlFlow) {
F := __e.Get(1)
_ = F
tmp4245 := MakeNative(func(__e *ControlFlow) {
Used := __e.Get(1)
_ = Used
tmp4246 := MakeNative(func(__e *ControlFlow) {
KL := __e.Get(1)
_ = KL
tmp4247 := MakeNative(func(__e *ControlFlow) {
EvalKL := __e.Get(1)
_ = EvalKL
tmp4248 := MakeNative(func(__e *ControlFlow) {
Record := __e.Get(1)
_ = Record
__e.Return(PrimCons(F, Used))
return
}, 1)

tmp4249 := Call(__e, PrimFunc(symshen_4record_1kl), F, KL)


__e.TailApply(tmp4248, tmp4249)
return


}, 1)

tmp4250 := Call(__e, PrimFunc(symeval_1kl), KL)


__e.TailApply(tmp4247, tmp4250)
return


}, 1)

tmp4251 := PrimCons(V1633, Nil)

tmp4252 := PrimCons(Used, tmp4251)

tmp4253 := PrimCons(F, tmp4252)

tmp4254 := PrimCons(symdefun, tmp4253)

__e.TailApply(tmp4246, tmp4254)
return


}, 1)

tmp4255 := Call(__e, PrimFunc(symshen_4remove_1if_1unused), V1632, V1633)


__e.TailApply(tmp4245, tmp4255)
return


}, 1)

tmp4256 := Call(__e, PrimFunc(symgensym), symshen_4else)


__e.TailApply(tmp4244, tmp4256)
return


}, 2)

tmp4257 := Call(__e, ns2_1set, symshen_4procedure_1call, tmp4243)


_ = tmp4257

tmp4258 := MakeNative(func(__e *ControlFlow) {
V1640 := __e.Get(1)
_ = V1640
V1641 := __e.Get(2)
_ = V1641
tmp4269 := PrimEqual(Nil, V1640)

if True == tmp4269 {
__e.Return(Nil)
return
} else {
tmp4267 := PrimIsPair(V1640)

if True == tmp4267 {
tmp4264 := PrimHead(V1640)

tmp4265 := Call(__e, PrimFunc(symshen_4occurs_2), tmp4264, V1641)


if True == tmp4265 {
tmp4259 := PrimHead(V1640)

tmp4260 := PrimTail(V1640)

tmp4261 := Call(__e, PrimFunc(symshen_4remove_1if_1unused), tmp4260, V1641)


__e.Return(PrimCons(tmp4259, tmp4261))
return


} else {
tmp4262 := PrimTail(V1640)

__e.TailApply(PrimFunc(symshen_4remove_1if_1unused), tmp4262, V1641)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-if-unused")))
return
}


}


}, 2)

tmp4270 := Call(__e, ns2_1set, symshen_4remove_1if_1unused, tmp4258)


_ = tmp4270

tmp4271 := MakeNative(func(__e *ControlFlow) {
V1642 := __e.Get(1)
_ = V1642
V1643 := __e.Get(2)
_ = V1643
V1644 := __e.Get(3)
_ = V1644
V1645 := __e.Get(4)
_ = V1645
tmp4272 := Call(__e, PrimFunc(symshen_4selectors), V1642, V1644)


__e.TailApply(PrimFunc(symshen_4horizontal), tmp4272, V1643, V1644, V1645)
return


}, 4)

tmp4273 := Call(__e, ns2_1set, symshen_4then, tmp4271)


_ = tmp4273

tmp4274 := MakeNative(func(__e *ControlFlow) {
V1654 := __e.Get(1)
_ = V1654
V1655 := __e.Get(2)
_ = V1655
V1656 := __e.Get(3)
_ = V1656
V1657 := __e.Get(4)
_ = V1657
tmp4290 := PrimEqual(Nil, V1654)

if True == tmp4290 {
__e.TailApply(PrimFunc(symshen_4vertical), V1655, V1656, V1657)
return
} else {
tmp4288 := PrimIsPair(V1654)

if True == tmp4288 {
tmp4275 := MakeNative(func(__e *ControlFlow) {
V := __e.Get(1)
_ = V
tmp4276 := PrimHead(V1654)

tmp4277 := PrimTail(V1654)

tmp4278 := PrimCons(V, V1655)

tmp4279 := PrimHead(V1654)

tmp4280 := Call(__e, PrimFunc(symsubst), V, tmp4279, V1656)


tmp4281 := Call(__e, PrimFunc(symshen_4horizontal), tmp4277, tmp4278, tmp4280, V1657)


tmp4282 := PrimCons(tmp4281, Nil)

tmp4283 := PrimCons(tmp4276, tmp4282)

tmp4284 := PrimCons(V, tmp4283)

__e.Return(PrimCons(symlet, tmp4284))
return


}, 1)

tmp4285 := Call(__e, PrimFunc(symprotect), symV)


tmp4286 := Call(__e, PrimFunc(symgensym), tmp4285)


__e.TailApply(tmp4275, tmp4286)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.horizontal")))
return
}


}


}, 4)

tmp4291 := Call(__e, ns2_1set, symshen_4horizontal, tmp4274)


_ = tmp4291

tmp4292 := MakeNative(func(__e *ControlFlow) {
V1662 := __e.Get(1)
_ = V1662
V1663 := __e.Get(2)
_ = V1663
tmp4329 := PrimIsPair(V1662)

var ifres4315 Obj

if True == tmp4329 {
tmp4327 := PrimTail(V1662)

tmp4328 := PrimIsPair(tmp4327)

var ifres4317 Obj

if True == tmp4328 {
tmp4324 := PrimTail(V1662)

tmp4325 := PrimTail(tmp4324)

tmp4326 := PrimEqual(Nil, tmp4325)

var ifres4319 Obj

if True == tmp4326 {
tmp4321 := PrimHead(V1662)

tmp4322 := Call(__e, PrimFunc(symshen_4op), tmp4321)


tmp4323 := Call(__e, PrimFunc(symshen_4constructor_2), tmp4322)


var ifres4320 Obj

if True == tmp4323 {
ifres4320 = True


} else {
ifres4320 = False


}

ifres4319 = ifres4320


} else {
ifres4319 = False


}

var ifres4318 Obj

if True == ifres4319 {
ifres4318 = True


} else {
ifres4318 = False


}

ifres4317 = ifres4318


} else {
ifres4317 = False


}

var ifres4316 Obj

if True == ifres4317 {
ifres4316 = True


} else {
ifres4316 = False


}

ifres4315 = ifres4316


} else {
ifres4315 = False


}

if True == ifres4315 {
tmp4293 := MakeNative(func(__e *ControlFlow) {
Op := __e.Get(1)
_ = Op
tmp4294 := MakeNative(func(__e *ControlFlow) {
Hd := __e.Get(1)
_ = Hd
tmp4295 := MakeNative(func(__e *ControlFlow) {
Tl := __e.Get(1)
_ = Tl
tmp4296 := MakeNative(func(__e *ControlFlow) {
RptedHd_2 := __e.Get(1)
_ = RptedHd_2
tmp4297 := MakeNative(func(__e *ControlFlow) {
RptedTl_2 := __e.Get(1)
_ = RptedTl_2
var ifres4302 Obj

if True == RptedHd_2 {
var ifres4303 Obj

if True == RptedTl_2 {
ifres4303 = True


} else {
ifres4303 = False


}

ifres4302 = ifres4303


} else {
ifres4302 = False


}

if True == ifres4302 {
tmp4298 := PrimCons(Tl, Nil)

__e.Return(PrimCons(Hd, tmp4298))
return


} else {
if True == RptedHd_2 {
__e.Return(PrimCons(Hd, Nil))
return
} else {
if True == RptedTl_2 {
__e.Return(PrimCons(Tl, Nil))
return
} else {
__e.Return(Nil)
return
}
}
}


}, 1)

tmp4304 := Call(__e, PrimFunc(symshen_4rpted_2), Tl, V1663)


__e.TailApply(tmp4297, tmp4304)
return


}, 1)

tmp4305 := Call(__e, PrimFunc(symshen_4rpted_2), Hd, V1663)


__e.TailApply(tmp4296, tmp4305)
return


}, 1)

tmp4306 := Call(__e, PrimFunc(symshen_4op2), Op)


tmp4307 := PrimTail(V1662)

tmp4308 := PrimCons(tmp4306, tmp4307)

__e.TailApply(tmp4295, tmp4308)
return


}, 1)

tmp4309 := Call(__e, PrimFunc(symshen_4op1), Op)


tmp4310 := PrimTail(V1662)

tmp4311 := PrimCons(tmp4309, tmp4310)

__e.TailApply(tmp4294, tmp4311)
return


}, 1)

tmp4312 := PrimHead(V1662)

tmp4313 := Call(__e, PrimFunc(symshen_4op), tmp4312)


__e.TailApply(tmp4293, tmp4313)
return


} else {
__e.Return(Nil)
return
}


}, 2)

tmp4330 := Call(__e, ns2_1set, symshen_4selectors, tmp4292)


_ = tmp4330

tmp4331 := MakeNative(func(__e *ControlFlow) {
V1664 := __e.Get(1)
_ = V1664
V1665 := __e.Get(2)
_ = V1665
tmp4332 := Call(__e, PrimFunc(symoccurrences), V1664, V1665)


__e.Return(PrimGreatThan(tmp4332, MakeNumber(1)))
return


}, 2)

tmp4333 := Call(__e, ns2_1set, symshen_4rpted_2, tmp4331)


_ = tmp4333

tmp4334 := MakeNative(func(__e *ControlFlow) {
V1666 := __e.Get(1)
_ = V1666
tmp4342 := PrimIsPair(V1666)

if True == tmp4342 {
tmp4339 := PrimHead(V1666)

tmp4340 := Call(__e, PrimFunc(symatom_2), tmp4339)


if True == tmp4340 {
tmp4336 := PrimTail(V1666)

tmp4337 := Call(__e, PrimFunc(symshen_4inline_2), tmp4336)


if True == tmp4337 {
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


} else {
__e.TailApply(PrimFunc(symatom_2), V1666)
return
}


}, 1)

tmp4343 := Call(__e, ns2_1set, symshen_4inline_2, tmp4334)


_ = tmp4343

tmp4344 := MakeNative(func(__e *ControlFlow) {
V1669 := __e.Get(1)
_ = V1669
tmp4352 := PrimEqual(symcons_2, V1669)

if True == tmp4352 {
__e.Return(symcons)
return
} else {
tmp4350 := PrimEqual(symshen_4_7string_2, V1669)

if True == tmp4350 {
__e.Return(sym_8s)
return
} else {
tmp4348 := PrimEqual(symshen_4_7vector_2, V1669)

if True == tmp4348 {
__e.Return(sym_8v)
return
} else {
tmp4346 := PrimEqual(symtuple_2, V1669)

if True == tmp4346 {
__e.Return(sym_8p)
return
} else {
__e.Return(symshen_4skip)
return
}


}


}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4op, tmp4344)
return




}, 0)

