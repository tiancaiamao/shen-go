package main

import . "github.com/tiancaiamao/shen-go/kl"

var CoreMain = MakeNative(func(__e *ControlFlow) {
tmp2585 := MakeNative(func(__e *ControlFlow) {
V502 := __e.Get(1)
_ = V502
tmp2586 := MakeNative(func(__e *ControlFlow) {
W503 := __e.Get(1)
_ = W503
__e.TailApply(PrimFunc(symshen_4record_1and_1evaluate), W503)
return
}, 1)

tmp2587 := Call(__e, PrimFunc(symshen_4shen_1_6kl_1h), V502)


__e.TailApply(tmp2586, tmp2587)
return


}, 1)

tmp2588 := Call(__e, ns2_1set, symshen_4shen_1_6kl, tmp2585)


_ = tmp2588

tmp2589 := MakeNative(func(__e *ControlFlow) {
V504 := __e.Get(1)
_ = V504
tmp2642 := PrimIsPair(V504)

var ifres2616 Obj

if True == tmp2642 {
tmp2640 := PrimHead(V504)

tmp2641 := PrimEqual(symdefun, tmp2640)

var ifres2618 Obj

if True == tmp2641 {
tmp2638 := PrimTail(V504)

tmp2639 := PrimIsPair(tmp2638)

var ifres2620 Obj

if True == tmp2639 {
tmp2635 := PrimTail(V504)

tmp2636 := PrimTail(tmp2635)

tmp2637 := PrimIsPair(tmp2636)

var ifres2622 Obj

if True == tmp2637 {
tmp2631 := PrimTail(V504)

tmp2632 := PrimTail(tmp2631)

tmp2633 := PrimTail(tmp2632)

tmp2634 := PrimIsPair(tmp2633)

var ifres2624 Obj

if True == tmp2634 {
tmp2626 := PrimTail(V504)

tmp2627 := PrimTail(tmp2626)

tmp2628 := PrimTail(tmp2627)

tmp2629 := PrimTail(tmp2628)

tmp2630 := PrimEqual(Nil, tmp2629)

var ifres2625 Obj

if True == tmp2630 {
ifres2625 = True


} else {
ifres2625 = False


}

ifres2624 = ifres2625


} else {
ifres2624 = False


}

var ifres2623 Obj

if True == ifres2624 {
ifres2623 = True


} else {
ifres2623 = False


}

ifres2622 = ifres2623


} else {
ifres2622 = False


}

var ifres2621 Obj

if True == ifres2622 {
ifres2621 = True


} else {
ifres2621 = False


}

ifres2620 = ifres2621


} else {
ifres2620 = False


}

var ifres2619 Obj

if True == ifres2620 {
ifres2619 = True


} else {
ifres2619 = False


}

ifres2618 = ifres2619


} else {
ifres2618 = False


}

var ifres2617 Obj

if True == ifres2618 {
ifres2617 = True


} else {
ifres2617 = False


}

ifres2616 = ifres2617


} else {
ifres2616 = False


}

if True == ifres2616 {
tmp2590 := MakeNative(func(__e *ControlFlow) {
W505 := __e.Get(1)
_ = W505
tmp2591 := MakeNative(func(__e *ControlFlow) {
W506 := __e.Get(1)
_ = W506
tmp2592 := MakeNative(func(__e *ControlFlow) {
W507 := __e.Get(1)
_ = W507
tmp2593 := MakeNative(func(__e *ControlFlow) {
W508 := __e.Get(1)
_ = W508
tmp2594 := PrimTail(V504)

tmp2595 := PrimHead(tmp2594)

__e.TailApply(PrimFunc(symshen_4fn_1print), tmp2595)
return


}, 1)

tmp2596 := Call(__e, PrimFunc(symeval_1kl), V504)


__e.TailApply(tmp2593, tmp2596)
return


}, 1)

tmp2597 := PrimTail(V504)

tmp2598 := PrimHead(tmp2597)

tmp2599 := Call(__e, PrimFunc(symshen_4record_1kl), tmp2598, V504)


__e.TailApply(tmp2592, tmp2599)
return


}, 1)

tmp2600 := PrimTail(V504)

tmp2601 := PrimHead(tmp2600)

tmp2602 := PrimTail(V504)

tmp2603 := PrimTail(tmp2602)

tmp2604 := PrimHead(tmp2603)

tmp2605 := Call(__e, PrimFunc(symlength), tmp2604)


tmp2606 := Call(__e, PrimFunc(symshen_4store_1arity), tmp2601, tmp2605)


__e.TailApply(tmp2591, tmp2606)
return


}, 1)

tmp2612 := PrimTail(V504)

tmp2613 := PrimHead(tmp2612)

tmp2614 := Call(__e, PrimFunc(symshen_4sysfunc_2), tmp2613)


var ifres2607 Obj

if True == tmp2614 {
tmp2608 := PrimTail(V504)

tmp2609 := PrimHead(tmp2608)

tmp2610 := Call(__e, PrimFunc(symshen_4app), tmp2609, MakeString(" is not a legitimate function name\n"), symshen_4a)


tmp2611 := PrimSimpleError(tmp2610)

ifres2607 = tmp2611


} else {
ifres2607 = symshen_4skip


}

__e.TailApply(tmp2590, ifres2607)
return


} else {
__e.Return(V504)
return
}


}, 1)

tmp2643 := Call(__e, ns2_1set, symshen_4record_1and_1evaluate, tmp2589)


_ = tmp2643

tmp2644 := MakeNative(func(__e *ControlFlow) {
V509 := __e.Get(1)
_ = V509
tmp2745 := PrimIsPair(V509)

var ifres2737 Obj

if True == tmp2745 {
tmp2743 := PrimHead(V509)

tmp2744 := PrimEqual(symdefine, tmp2743)

var ifres2739 Obj

if True == tmp2744 {
tmp2741 := PrimTail(V509)

tmp2742 := PrimIsPair(tmp2741)

var ifres2740 Obj

if True == tmp2742 {
ifres2740 = True


} else {
ifres2740 = False


}

ifres2739 = ifres2740


} else {
ifres2739 = False


}

var ifres2738 Obj

if True == ifres2739 {
ifres2738 = True


} else {
ifres2738 = False


}

ifres2737 = ifres2738


} else {
ifres2737 = False


}

if True == ifres2737 {
tmp2645 := PrimTail(V509)

tmp2646 := PrimHead(tmp2645)

tmp2647 := PrimTail(V509)

tmp2648 := PrimTail(tmp2647)

__e.TailApply(PrimFunc(symshen_4shendef_1_6kldef), tmp2646, tmp2648)
return


} else {
tmp2735 := PrimIsPair(V509)

var ifres2709 Obj

if True == tmp2735 {
tmp2733 := PrimHead(V509)

tmp2734 := PrimEqual(symdefun, tmp2733)

var ifres2711 Obj

if True == tmp2734 {
tmp2731 := PrimTail(V509)

tmp2732 := PrimIsPair(tmp2731)

var ifres2713 Obj

if True == tmp2732 {
tmp2728 := PrimTail(V509)

tmp2729 := PrimTail(tmp2728)

tmp2730 := PrimIsPair(tmp2729)

var ifres2715 Obj

if True == tmp2730 {
tmp2724 := PrimTail(V509)

tmp2725 := PrimTail(tmp2724)

tmp2726 := PrimTail(tmp2725)

tmp2727 := PrimIsPair(tmp2726)

var ifres2717 Obj

if True == tmp2727 {
tmp2719 := PrimTail(V509)

tmp2720 := PrimTail(tmp2719)

tmp2721 := PrimTail(tmp2720)

tmp2722 := PrimTail(tmp2721)

tmp2723 := PrimEqual(Nil, tmp2722)

var ifres2718 Obj

if True == tmp2723 {
ifres2718 = True


} else {
ifres2718 = False


}

ifres2717 = ifres2718


} else {
ifres2717 = False


}

var ifres2716 Obj

if True == ifres2717 {
ifres2716 = True


} else {
ifres2716 = False


}

ifres2715 = ifres2716


} else {
ifres2715 = False


}

var ifres2714 Obj

if True == ifres2715 {
ifres2714 = True


} else {
ifres2714 = False


}

ifres2713 = ifres2714


} else {
ifres2713 = False


}

var ifres2712 Obj

if True == ifres2713 {
ifres2712 = True


} else {
ifres2712 = False


}

ifres2711 = ifres2712


} else {
ifres2711 = False


}

var ifres2710 Obj

if True == ifres2711 {
ifres2710 = True


} else {
ifres2710 = False


}

ifres2709 = ifres2710


} else {
ifres2709 = False


}

if True == ifres2709 {
__e.Return(V509)
return
} else {
tmp2707 := PrimIsPair(V509)

var ifres2688 Obj

if True == tmp2707 {
tmp2705 := PrimHead(V509)

tmp2706 := PrimEqual(symtype, tmp2705)

var ifres2690 Obj

if True == tmp2706 {
tmp2703 := PrimTail(V509)

tmp2704 := PrimIsPair(tmp2703)

var ifres2692 Obj

if True == tmp2704 {
tmp2700 := PrimTail(V509)

tmp2701 := PrimTail(tmp2700)

tmp2702 := PrimIsPair(tmp2701)

var ifres2694 Obj

if True == tmp2702 {
tmp2696 := PrimTail(V509)

tmp2697 := PrimTail(tmp2696)

tmp2698 := PrimTail(tmp2697)

tmp2699 := PrimEqual(Nil, tmp2698)

var ifres2695 Obj

if True == tmp2699 {
ifres2695 = True


} else {
ifres2695 = False


}

ifres2694 = ifres2695


} else {
ifres2694 = False


}

var ifres2693 Obj

if True == ifres2694 {
ifres2693 = True


} else {
ifres2693 = False


}

ifres2692 = ifres2693


} else {
ifres2692 = False


}

var ifres2691 Obj

if True == ifres2692 {
ifres2691 = True


} else {
ifres2691 = False


}

ifres2690 = ifres2691


} else {
ifres2690 = False


}

var ifres2689 Obj

if True == ifres2690 {
ifres2689 = True


} else {
ifres2689 = False


}

ifres2688 = ifres2689


} else {
ifres2688 = False


}

if True == ifres2688 {
tmp2649 := PrimTail(V509)

tmp2650 := PrimHead(tmp2649)

tmp2651 := PrimTail(V509)

tmp2652 := PrimTail(tmp2651)

tmp2653 := PrimHead(tmp2652)

tmp2654 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2653)


tmp2655 := PrimCons(tmp2654, Nil)

tmp2656 := PrimCons(tmp2650, tmp2655)

__e.Return(PrimCons(symtype, tmp2656))
return


} else {
tmp2686 := PrimIsPair(V509)

var ifres2667 Obj

if True == tmp2686 {
tmp2684 := PrimHead(V509)

tmp2685 := PrimEqual(syminput_7, tmp2684)

var ifres2669 Obj

if True == tmp2685 {
tmp2682 := PrimTail(V509)

tmp2683 := PrimIsPair(tmp2682)

var ifres2671 Obj

if True == tmp2683 {
tmp2679 := PrimTail(V509)

tmp2680 := PrimTail(tmp2679)

tmp2681 := PrimIsPair(tmp2680)

var ifres2673 Obj

if True == tmp2681 {
tmp2675 := PrimTail(V509)

tmp2676 := PrimTail(tmp2675)

tmp2677 := PrimTail(tmp2676)

tmp2678 := PrimEqual(Nil, tmp2677)

var ifres2674 Obj

if True == tmp2678 {
ifres2674 = True


} else {
ifres2674 = False


}

ifres2673 = ifres2674


} else {
ifres2673 = False


}

var ifres2672 Obj

if True == ifres2673 {
ifres2672 = True


} else {
ifres2672 = False


}

ifres2671 = ifres2672


} else {
ifres2671 = False


}

var ifres2670 Obj

if True == ifres2671 {
ifres2670 = True


} else {
ifres2670 = False


}

ifres2669 = ifres2670


} else {
ifres2669 = False


}

var ifres2668 Obj

if True == ifres2669 {
ifres2668 = True


} else {
ifres2668 = False


}

ifres2667 = ifres2668


} else {
ifres2667 = False


}

if True == ifres2667 {
tmp2657 := PrimTail(V509)

tmp2658 := PrimHead(tmp2657)

tmp2659 := Call(__e, PrimFunc(symshen_4rcons__form), tmp2658)


tmp2660 := PrimTail(V509)

tmp2661 := PrimTail(tmp2660)

tmp2662 := PrimCons(tmp2659, tmp2661)

__e.Return(PrimCons(syminput_7, tmp2662))
return


} else {
tmp2665 := PrimIsPair(V509)

if True == tmp2665 {
tmp2663 := MakeNative(func(__e *ControlFlow) {
Z510 := __e.Get(1)
_ = Z510
__e.TailApply(PrimFunc(symshen_4shen_1_6kl_1h), Z510)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2663, V509)
return


} else {
__e.Return(V509)
return
}


}


}


}


}


}, 1)

tmp2746 := Call(__e, ns2_1set, symshen_4shen_1_6kl_1h, tmp2644)


_ = tmp2746

tmp2747 := MakeNative(func(__e *ControlFlow) {
V511 := __e.Get(1)
_ = V511
V512 := __e.Get(2)
_ = V512
tmp2748 := MakeNative(func(__e *ControlFlow) {
Z513 := __e.Get(1)
_ = Z513
__e.TailApply(PrimFunc(symshen_4_5define_6), Z513)
return
}, 1)

tmp2749 := PrimCons(V511, V512)

__e.TailApply(PrimFunc(symcompile), tmp2748, tmp2749)
return


}, 2)

tmp2750 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef, tmp2747)


_ = tmp2750

tmp2751 := MakeNative(func(__e *ControlFlow) {
V514 := __e.Get(1)
_ = V514
tmp2752 := MakeNative(func(__e *ControlFlow) {
W515 := __e.Get(1)
_ = W515
tmp2775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W515)


if True == tmp2775 {
tmp2753 := MakeNative(func(__e *ControlFlow) {
W526 := __e.Get(1)
_ = W526
tmp2755 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W526)


if True == tmp2755 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W526)
return
}


}, 1)

tmp2756 := MakeNative(func(__e *ControlFlow) {
W527 := __e.Get(1)
_ = W527
tmp2771 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W527)


if True == tmp2771 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2757 := MakeNative(func(__e *ControlFlow) {
W528 := __e.Get(1)
_ = W528
tmp2758 := MakeNative(func(__e *ControlFlow) {
W529 := __e.Get(1)
_ = W529
tmp2759 := MakeNative(func(__e *ControlFlow) {
W530 := __e.Get(1)
_ = W530
tmp2766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W530)


if True == tmp2766 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2760 := MakeNative(func(__e *ControlFlow) {
W531 := __e.Get(1)
_ = W531
tmp2761 := MakeNative(func(__e *ControlFlow) {
W532 := __e.Get(1)
_ = W532
tmp2762 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W528, W531)


__e.TailApply(PrimFunc(symshen_4comb), W532, tmp2762)
return


}, 1)

tmp2763 := Call(__e, PrimFunc(symshen_4in_1_6), W530)


__e.TailApply(tmp2761, tmp2763)
return


}, 1)

tmp2764 := Call(__e, PrimFunc(symshen_4_5_1out), W530)


__e.TailApply(tmp2760, tmp2764)
return


}


}, 1)

tmp2767 := Call(__e, PrimFunc(symshen_4_5rules_6), W529)


__e.TailApply(tmp2759, tmp2767)
return


}, 1)

tmp2768 := Call(__e, PrimFunc(symshen_4in_1_6), W527)


__e.TailApply(tmp2758, tmp2768)
return


}, 1)

tmp2769 := Call(__e, PrimFunc(symshen_4_5_1out), W527)


__e.TailApply(tmp2757, tmp2769)
return


}


}, 1)

tmp2772 := Call(__e, PrimFunc(symshen_4_5name_6), V514)


tmp2773 := Call(__e, tmp2756, tmp2772)


__e.TailApply(tmp2753, tmp2773)
return


} else {
__e.Return(W515)
return
}


}, 1)

tmp2776 := MakeNative(func(__e *ControlFlow) {
W516 := __e.Get(1)
_ = W516
tmp2805 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W516)


if True == tmp2805 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2777 := MakeNative(func(__e *ControlFlow) {
W517 := __e.Get(1)
_ = W517
tmp2778 := MakeNative(func(__e *ControlFlow) {
W518 := __e.Get(1)
_ = W518
tmp2801 := Call(__e, PrimFunc(symshen_4hds_a_2), W518, sym_i)


if True == tmp2801 {
tmp2779 := MakeNative(func(__e *ControlFlow) {
W519 := __e.Get(1)
_ = W519
tmp2780 := MakeNative(func(__e *ControlFlow) {
W520 := __e.Get(1)
_ = W520
tmp2797 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W520)


if True == tmp2797 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2781 := MakeNative(func(__e *ControlFlow) {
W521 := __e.Get(1)
_ = W521
tmp2794 := Call(__e, PrimFunc(symshen_4hds_a_2), W521, sym_j)


if True == tmp2794 {
tmp2782 := MakeNative(func(__e *ControlFlow) {
W522 := __e.Get(1)
_ = W522
tmp2783 := MakeNative(func(__e *ControlFlow) {
W523 := __e.Get(1)
_ = W523
tmp2790 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W523)


if True == tmp2790 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2784 := MakeNative(func(__e *ControlFlow) {
W524 := __e.Get(1)
_ = W524
tmp2785 := MakeNative(func(__e *ControlFlow) {
W525 := __e.Get(1)
_ = W525
tmp2786 := Call(__e, PrimFunc(symshen_4shendef_1_6kldef_1h), W517, W524)


__e.TailApply(PrimFunc(symshen_4comb), W525, tmp2786)
return


}, 1)

tmp2787 := Call(__e, PrimFunc(symshen_4in_1_6), W523)


__e.TailApply(tmp2785, tmp2787)
return


}, 1)

tmp2788 := Call(__e, PrimFunc(symshen_4_5_1out), W523)


__e.TailApply(tmp2784, tmp2788)
return


}


}, 1)

tmp2791 := Call(__e, PrimFunc(symshen_4_5rules_6), W522)


__e.TailApply(tmp2783, tmp2791)
return


}, 1)

tmp2792 := Call(__e, PrimFunc(symtail), W521)


__e.TailApply(tmp2782, tmp2792)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2795 := Call(__e, PrimFunc(symshen_4in_1_6), W520)


__e.TailApply(tmp2781, tmp2795)
return


}


}, 1)

tmp2798 := Call(__e, PrimFunc(symshen_4_5signature_6), W519)


__e.TailApply(tmp2780, tmp2798)
return


}, 1)

tmp2799 := Call(__e, PrimFunc(symtail), W518)


__e.TailApply(tmp2779, tmp2799)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2802 := Call(__e, PrimFunc(symshen_4in_1_6), W516)


__e.TailApply(tmp2778, tmp2802)
return


}, 1)

tmp2803 := Call(__e, PrimFunc(symshen_4_5_1out), W516)


__e.TailApply(tmp2777, tmp2803)
return


}


}, 1)

tmp2806 := Call(__e, PrimFunc(symshen_4_5name_6), V514)


tmp2807 := Call(__e, tmp2776, tmp2806)


__e.TailApply(tmp2752, tmp2807)
return


}, 1)

tmp2808 := Call(__e, ns2_1set, symshen_4_5define_6, tmp2751)


_ = tmp2808

tmp2809 := MakeNative(func(__e *ControlFlow) {
V533 := __e.Get(1)
_ = V533
V534 := __e.Get(2)
_ = V534
tmp2810 := MakeNative(func(__e *ControlFlow) {
W535 := __e.Get(1)
_ = W535
tmp2811 := MakeNative(func(__e *ControlFlow) {
W537 := __e.Get(1)
_ = W537
tmp2812 := MakeNative(func(__e *ControlFlow) {
W538 := __e.Get(1)
_ = W538
tmp2813 := MakeNative(func(__e *ControlFlow) {
W540 := __e.Get(1)
_ = W540
tmp2814 := MakeNative(func(__e *ControlFlow) {
W541 := __e.Get(1)
_ = W541
__e.Return(W541)
return
}, 1)

tmp2815 := Call(__e, PrimFunc(symshen_4compile_1to_1kl), V533, W540, W537)


tmp2816 := Call(__e, PrimFunc(symshen_4factorise_1code), tmp2815)


__e.TailApply(tmp2814, tmp2816)
return


}, 1)

tmp2817 := Call(__e, PrimFunc(symshen_4unprotect), V534)


__e.TailApply(tmp2813, tmp2817)
return


}, 1)

tmp2818 := MakeNative(func(__e *ControlFlow) {
Z539 := __e.Get(1)
_ = Z539
__e.TailApply(PrimFunc(symshen_4free_1var_1chk), V533, Z539)
return
}, 1)

tmp2819 := Call(__e, PrimFunc(symmap), tmp2818, V534)


__e.TailApply(tmp2812, tmp2819)
return


}, 1)

tmp2820 := Call(__e, PrimFunc(symshen_4arity_1chk), V533, W535)


__e.TailApply(tmp2811, tmp2820)
return


}, 1)

tmp2821 := MakeNative(func(__e *ControlFlow) {
Z536 := __e.Get(1)
_ = Z536
__e.TailApply(PrimFunc(symfst), Z536)
return
}, 1)

tmp2822 := Call(__e, PrimFunc(symmap), tmp2821, V534)


__e.TailApply(tmp2810, tmp2822)
return


}, 2)

tmp2823 := Call(__e, ns2_1set, symshen_4shendef_1_6kldef_1h, tmp2809)


_ = tmp2823

tmp2824 := MakeNative(func(__e *ControlFlow) {
V542 := __e.Get(1)
_ = V542
tmp2850 := Call(__e, PrimFunc(symtuple_2), V542)


if True == tmp2850 {
tmp2825 := Call(__e, PrimFunc(symfst), V542)


tmp2826 := Call(__e, PrimFunc(symshen_4unprotect), tmp2825)


tmp2827 := Call(__e, PrimFunc(symsnd), V542)


tmp2828 := Call(__e, PrimFunc(symshen_4unprotect), tmp2827)


__e.TailApply(PrimFunc(sym_8p), tmp2826, tmp2828)
return


} else {
tmp2848 := PrimIsPair(V542)

var ifres2835 Obj

if True == tmp2848 {
tmp2846 := PrimHead(V542)

tmp2847 := PrimEqual(symprotect, tmp2846)

var ifres2837 Obj

if True == tmp2847 {
tmp2844 := PrimTail(V542)

tmp2845 := PrimIsPair(tmp2844)

var ifres2839 Obj

if True == tmp2845 {
tmp2841 := PrimTail(V542)

tmp2842 := PrimTail(tmp2841)

tmp2843 := PrimEqual(Nil, tmp2842)

var ifres2840 Obj

if True == tmp2843 {
ifres2840 = True


} else {
ifres2840 = False


}

ifres2839 = ifres2840


} else {
ifres2839 = False


}

var ifres2838 Obj

if True == ifres2839 {
ifres2838 = True


} else {
ifres2838 = False


}

ifres2837 = ifres2838


} else {
ifres2837 = False


}

var ifres2836 Obj

if True == ifres2837 {
ifres2836 = True


} else {
ifres2836 = False


}

ifres2835 = ifres2836


} else {
ifres2835 = False


}

if True == ifres2835 {
tmp2829 := PrimTail(V542)

tmp2830 := PrimHead(tmp2829)

__e.TailApply(PrimFunc(symshen_4unprotect), tmp2830)
return


} else {
tmp2833 := PrimIsPair(V542)

if True == tmp2833 {
tmp2831 := MakeNative(func(__e *ControlFlow) {
Z543 := __e.Get(1)
_ = Z543
__e.TailApply(PrimFunc(symshen_4unprotect), Z543)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp2831, V542)
return


} else {
__e.Return(V542)
return
}


}


}


}, 1)

tmp2851 := Call(__e, ns2_1set, symshen_4unprotect, tmp2824)


_ = tmp2851

tmp2852 := MakeNative(func(__e *ControlFlow) {
V544 := __e.Get(1)
_ = V544
tmp2853 := MakeNative(func(__e *ControlFlow) {
W545 := __e.Get(1)
_ = W545
tmp2855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W545)


if True == tmp2855 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W545)
return
}


}, 1)

tmp2871 := PrimIsPair(V544)

var ifres2856 Obj

if True == tmp2871 {
tmp2857 := MakeNative(func(__e *ControlFlow) {
W546 := __e.Get(1)
_ = W546
tmp2858 := MakeNative(func(__e *ControlFlow) {
W547 := __e.Get(1)
_ = W547
tmp2866 := PrimIsSymbol(W546)

var ifres2862 Obj

if True == tmp2866 {
tmp2864 := PrimIsVariable(W546)

tmp2865 := PrimNot(tmp2864)

var ifres2863 Obj

if True == tmp2865 {
ifres2863 = True


} else {
ifres2863 = False


}

ifres2862 = ifres2863


} else {
ifres2862 = False


}

var ifres2859 Obj

if True == ifres2862 {
ifres2859 = W546


} else {
tmp2860 := Call(__e, PrimFunc(symshen_4app), W546, MakeString(" is not a legitimate function name.\n"), symshen_4a)


tmp2861 := PrimSimpleError(tmp2860)

ifres2859 = tmp2861


}

__e.TailApply(PrimFunc(symshen_4comb), W547, ifres2859)
return


}, 1)

tmp2867 := Call(__e, PrimFunc(symtail), V544)


__e.TailApply(tmp2858, tmp2867)
return


}, 1)

tmp2868 := Call(__e, PrimFunc(symhead), V544)


tmp2869 := Call(__e, tmp2857, tmp2868)


ifres2856 = tmp2869


} else {
tmp2870 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2856 = tmp2870


}

__e.TailApply(tmp2853, ifres2856)
return


}, 1)

tmp2872 := Call(__e, ns2_1set, symshen_4_5name_6, tmp2852)


_ = tmp2872

tmp2873 := MakeNative(func(__e *ControlFlow) {
V548 := __e.Get(1)
_ = V548
tmp2874 := MakeNative(func(__e *ControlFlow) {
W549 := __e.Get(1)
_ = W549
tmp2886 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W549)


if True == tmp2886 {
tmp2875 := MakeNative(func(__e *ControlFlow) {
W555 := __e.Get(1)
_ = W555
tmp2877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W555)


if True == tmp2877 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W555)
return
}


}, 1)

tmp2878 := MakeNative(func(__e *ControlFlow) {
W556 := __e.Get(1)
_ = W556
tmp2882 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W556)


if True == tmp2882 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2879 := MakeNative(func(__e *ControlFlow) {
W557 := __e.Get(1)
_ = W557
__e.TailApply(PrimFunc(symshen_4comb), W557, Nil)
return
}, 1)

tmp2880 := Call(__e, PrimFunc(symshen_4in_1_6), W556)


__e.TailApply(tmp2879, tmp2880)
return


}


}, 1)

tmp2883 := Call(__e, PrimFunc(sym_5e_6), V548)


tmp2884 := Call(__e, tmp2878, tmp2883)


__e.TailApply(tmp2875, tmp2884)
return


} else {
__e.Return(W549)
return
}


}, 1)

tmp2908 := PrimIsPair(V548)

var ifres2887 Obj

if True == tmp2908 {
tmp2888 := MakeNative(func(__e *ControlFlow) {
W550 := __e.Get(1)
_ = W550
tmp2889 := MakeNative(func(__e *ControlFlow) {
W551 := __e.Get(1)
_ = W551
tmp2890 := MakeNative(func(__e *ControlFlow) {
W552 := __e.Get(1)
_ = W552
tmp2902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W552)


if True == tmp2902 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2891 := MakeNative(func(__e *ControlFlow) {
W553 := __e.Get(1)
_ = W553
tmp2892 := MakeNative(func(__e *ControlFlow) {
W554 := __e.Get(1)
_ = W554
tmp2895 := PrimCons(sym_j, Nil)

tmp2896 := PrimCons(sym_i, tmp2895)

tmp2897 := Call(__e, PrimFunc(symelement_2), W550, tmp2896)


tmp2898 := PrimNot(tmp2897)

if True == tmp2898 {
tmp2893 := PrimCons(W550, W553)

__e.TailApply(PrimFunc(symshen_4comb), W554, tmp2893)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp2899 := Call(__e, PrimFunc(symshen_4in_1_6), W552)


__e.TailApply(tmp2892, tmp2899)
return


}, 1)

tmp2900 := Call(__e, PrimFunc(symshen_4_5_1out), W552)


__e.TailApply(tmp2891, tmp2900)
return


}


}, 1)

tmp2903 := Call(__e, PrimFunc(symshen_4_5signature_6), W551)


__e.TailApply(tmp2890, tmp2903)
return


}, 1)

tmp2904 := Call(__e, PrimFunc(symtail), V548)


__e.TailApply(tmp2889, tmp2904)
return


}, 1)

tmp2905 := Call(__e, PrimFunc(symhead), V548)


tmp2906 := Call(__e, tmp2888, tmp2905)


ifres2887 = tmp2906


} else {
tmp2907 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres2887 = tmp2907


}

__e.TailApply(tmp2874, ifres2887)
return


}, 1)

tmp2909 := Call(__e, ns2_1set, symshen_4_5signature_6, tmp2873)


_ = tmp2909

tmp2910 := MakeNative(func(__e *ControlFlow) {
V558 := __e.Get(1)
_ = V558
tmp2911 := MakeNative(func(__e *ControlFlow) {
W559 := __e.Get(1)
_ = W559
tmp2930 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W559)


if True == tmp2930 {
tmp2912 := MakeNative(func(__e *ControlFlow) {
W566 := __e.Get(1)
_ = W566
tmp2914 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W566)


if True == tmp2914 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W566)
return
}


}, 1)

tmp2915 := MakeNative(func(__e *ControlFlow) {
W567 := __e.Get(1)
_ = W567
tmp2926 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W567)


if True == tmp2926 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2916 := MakeNative(func(__e *ControlFlow) {
W568 := __e.Get(1)
_ = W568
tmp2917 := MakeNative(func(__e *ControlFlow) {
W569 := __e.Get(1)
_ = W569
tmp2922 := Call(__e, PrimFunc(symempty_2), W568)


var ifres2918 Obj

if True == tmp2922 {
ifres2918 = Nil


} else {
tmp2919 := Call(__e, PrimFunc(symshen_4app), W568, MakeString("\n ..."), symshen_4r)


tmp2920 := PrimStringConcat(MakeString("Shen syntax error here:\n "), tmp2919)

tmp2921 := PrimSimpleError(tmp2920)

ifres2918 = tmp2921


}

__e.TailApply(PrimFunc(symshen_4comb), W569, ifres2918)
return


}, 1)

tmp2923 := Call(__e, PrimFunc(symshen_4in_1_6), W567)


__e.TailApply(tmp2917, tmp2923)
return


}, 1)

tmp2924 := Call(__e, PrimFunc(symshen_4_5_1out), W567)


__e.TailApply(tmp2916, tmp2924)
return


}


}, 1)

tmp2927 := Call(__e, PrimFunc(sym_5_b_6), V558)


tmp2928 := Call(__e, tmp2915, tmp2927)


__e.TailApply(tmp2912, tmp2928)
return


} else {
__e.Return(W559)
return
}


}, 1)

tmp2931 := MakeNative(func(__e *ControlFlow) {
W560 := __e.Get(1)
_ = W560
tmp2947 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W560)


if True == tmp2947 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2932 := MakeNative(func(__e *ControlFlow) {
W561 := __e.Get(1)
_ = W561
tmp2933 := MakeNative(func(__e *ControlFlow) {
W562 := __e.Get(1)
_ = W562
tmp2934 := MakeNative(func(__e *ControlFlow) {
W563 := __e.Get(1)
_ = W563
tmp2942 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W563)


if True == tmp2942 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp2935 := MakeNative(func(__e *ControlFlow) {
W564 := __e.Get(1)
_ = W564
tmp2936 := MakeNative(func(__e *ControlFlow) {
W565 := __e.Get(1)
_ = W565
tmp2937 := Call(__e, PrimFunc(symshen_4linearise), W561)


tmp2938 := PrimCons(tmp2937, W564)

__e.TailApply(PrimFunc(symshen_4comb), W565, tmp2938)
return


}, 1)

tmp2939 := Call(__e, PrimFunc(symshen_4in_1_6), W563)


__e.TailApply(tmp2936, tmp2939)
return


}, 1)

tmp2940 := Call(__e, PrimFunc(symshen_4_5_1out), W563)


__e.TailApply(tmp2935, tmp2940)
return


}


}, 1)

tmp2943 := Call(__e, PrimFunc(symshen_4_5rules_6), W562)


__e.TailApply(tmp2934, tmp2943)
return


}, 1)

tmp2944 := Call(__e, PrimFunc(symshen_4in_1_6), W560)


__e.TailApply(tmp2933, tmp2944)
return


}, 1)

tmp2945 := Call(__e, PrimFunc(symshen_4_5_1out), W560)


__e.TailApply(tmp2932, tmp2945)
return


}


}, 1)

tmp2948 := Call(__e, PrimFunc(symshen_4_5rule_6), V558)


tmp2949 := Call(__e, tmp2931, tmp2948)


__e.TailApply(tmp2911, tmp2949)
return


}, 1)

tmp2950 := Call(__e, ns2_1set, symshen_4_5rules_6, tmp2910)


_ = tmp2950

tmp2951 := MakeNative(func(__e *ControlFlow) {
V572 := __e.Get(1)
_ = V572
tmp2956 := Call(__e, PrimFunc(symtuple_2), V572)


if True == tmp2956 {
tmp2952 := Call(__e, PrimFunc(symfst), V572)


tmp2953 := Call(__e, PrimFunc(symfst), V572)


tmp2954 := Call(__e, PrimFunc(symsnd), V572)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2952, tmp2953, Nil, tmp2954)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise")))
return
}


}, 1)

tmp2957 := Call(__e, ns2_1set, symshen_4linearise, tmp2951)


_ = tmp2957

tmp2958 := MakeNative(func(__e *ControlFlow) {
V585 := __e.Get(1)
_ = V585
V586 := __e.Get(2)
_ = V586
V587 := __e.Get(3)
_ = V587
V588 := __e.Get(4)
_ = V588
tmp2996 := PrimEqual(Nil, V585)

if True == tmp2996 {
__e.TailApply(PrimFunc(sym_8p), V586, V588)
return
} else {
tmp2994 := PrimIsPair(V585)

var ifres2990 Obj

if True == tmp2994 {
tmp2992 := PrimHead(V585)

tmp2993 := PrimIsPair(tmp2992)

var ifres2991 Obj

if True == tmp2993 {
ifres2991 = True


} else {
ifres2991 = False


}

ifres2990 = ifres2991


} else {
ifres2990 = False


}

if True == ifres2990 {
tmp2959 := PrimHead(V585)

tmp2960 := PrimTail(V585)

tmp2961 := Call(__e, PrimFunc(symappend), tmp2959, tmp2960)


__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2961, V586, V587, V588)
return


} else {
tmp2988 := PrimIsPair(V585)

var ifres2984 Obj

if True == tmp2988 {
tmp2986 := PrimHead(V585)

tmp2987 := PrimIsVariable(tmp2986)

var ifres2985 Obj

if True == tmp2987 {
ifres2985 = True


} else {
ifres2985 = False


}

ifres2984 = ifres2985


} else {
ifres2984 = False


}

if True == ifres2984 {
tmp2978 := PrimHead(V585)

tmp2979 := Call(__e, PrimFunc(symelement_2), tmp2978, V587)


if True == tmp2979 {
tmp2962 := MakeNative(func(__e *ControlFlow) {
W589 := __e.Get(1)
_ = W589
tmp2963 := PrimTail(V585)

tmp2964 := PrimHead(V585)

tmp2965 := Call(__e, PrimFunc(symshen_4rep_1X), tmp2964, W589, V586)


tmp2966 := PrimHead(V585)

tmp2967 := PrimCons(tmp2966, Nil)

tmp2968 := PrimCons(W589, tmp2967)

tmp2969 := PrimCons(sym_a, tmp2968)

tmp2970 := PrimCons(V588, Nil)

tmp2971 := PrimCons(tmp2969, tmp2970)

tmp2972 := PrimCons(symwhere, tmp2971)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2963, tmp2965, V587, tmp2972)
return


}, 1)

tmp2973 := Call(__e, PrimFunc(symgensym), symV)


__e.TailApply(tmp2962, tmp2973)
return


} else {
tmp2974 := PrimTail(V585)

tmp2975 := PrimHead(V585)

tmp2976 := PrimCons(tmp2975, V587)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2974, V586, tmp2976, V588)
return


}


} else {
tmp2982 := PrimIsPair(V585)

if True == tmp2982 {
tmp2980 := PrimTail(V585)

__e.TailApply(PrimFunc(symshen_4linearise_1h), tmp2980, V586, V587, V588)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.linearise-h")))
return
}


}


}


}


}, 4)

tmp2997 := Call(__e, ns2_1set, symshen_4linearise_1h, tmp2958)


_ = tmp2997

tmp2998 := MakeNative(func(__e *ControlFlow) {
V590 := __e.Get(1)
_ = V590
tmp2999 := MakeNative(func(__e *ControlFlow) {
W591 := __e.Get(1)
_ = W591
tmp3087 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W591)


if True == tmp3087 {
tmp3000 := MakeNative(func(__e *ControlFlow) {
W601 := __e.Get(1)
_ = W601
tmp3065 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W601)


if True == tmp3065 {
tmp3001 := MakeNative(func(__e *ControlFlow) {
W608 := __e.Get(1)
_ = W608
tmp3028 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W608)


if True == tmp3028 {
tmp3002 := MakeNative(func(__e *ControlFlow) {
W618 := __e.Get(1)
_ = W618
tmp3004 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W618)


if True == tmp3004 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W618)
return
}


}, 1)

tmp3005 := MakeNative(func(__e *ControlFlow) {
W619 := __e.Get(1)
_ = W619
tmp3024 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W619)


if True == tmp3024 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3006 := MakeNative(func(__e *ControlFlow) {
W620 := __e.Get(1)
_ = W620
tmp3007 := MakeNative(func(__e *ControlFlow) {
W621 := __e.Get(1)
_ = W621
tmp3020 := Call(__e, PrimFunc(symshen_4hds_a_2), W621, sym_5_1)


if True == tmp3020 {
tmp3008 := MakeNative(func(__e *ControlFlow) {
W622 := __e.Get(1)
_ = W622
tmp3017 := PrimIsPair(W622)

if True == tmp3017 {
tmp3009 := MakeNative(func(__e *ControlFlow) {
W623 := __e.Get(1)
_ = W623
tmp3010 := MakeNative(func(__e *ControlFlow) {
W624 := __e.Get(1)
_ = W624
tmp3011 := PrimCons(W623, Nil)

tmp3012 := PrimCons(symshen_4choicepoint_b, tmp3011)

tmp3013 := Call(__e, PrimFunc(sym_8p), W620, tmp3012)


__e.TailApply(PrimFunc(symshen_4comb), W624, tmp3013)
return


}, 1)

tmp3014 := Call(__e, PrimFunc(symtail), W622)


__e.TailApply(tmp3010, tmp3014)
return


}, 1)

tmp3015 := Call(__e, PrimFunc(symhead), W622)


__e.TailApply(tmp3009, tmp3015)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3018 := Call(__e, PrimFunc(symtail), W621)


__e.TailApply(tmp3008, tmp3018)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3021 := Call(__e, PrimFunc(symshen_4in_1_6), W619)


__e.TailApply(tmp3007, tmp3021)
return


}, 1)

tmp3022 := Call(__e, PrimFunc(symshen_4_5_1out), W619)


__e.TailApply(tmp3006, tmp3022)
return


}


}, 1)

tmp3025 := Call(__e, PrimFunc(symshen_4_5patterns_6), V590)


tmp3026 := Call(__e, tmp3005, tmp3025)


__e.TailApply(tmp3002, tmp3026)
return


} else {
__e.Return(W608)
return
}


}, 1)

tmp3029 := MakeNative(func(__e *ControlFlow) {
W609 := __e.Get(1)
_ = W609
tmp3061 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W609)


if True == tmp3061 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3030 := MakeNative(func(__e *ControlFlow) {
W610 := __e.Get(1)
_ = W610
tmp3031 := MakeNative(func(__e *ControlFlow) {
W611 := __e.Get(1)
_ = W611
tmp3057 := Call(__e, PrimFunc(symshen_4hds_a_2), W611, sym_5_1)


if True == tmp3057 {
tmp3032 := MakeNative(func(__e *ControlFlow) {
W612 := __e.Get(1)
_ = W612
tmp3054 := PrimIsPair(W612)

if True == tmp3054 {
tmp3033 := MakeNative(func(__e *ControlFlow) {
W613 := __e.Get(1)
_ = W613
tmp3034 := MakeNative(func(__e *ControlFlow) {
W614 := __e.Get(1)
_ = W614
tmp3050 := Call(__e, PrimFunc(symshen_4hds_a_2), W614, symwhere)


if True == tmp3050 {
tmp3035 := MakeNative(func(__e *ControlFlow) {
W615 := __e.Get(1)
_ = W615
tmp3047 := PrimIsPair(W615)

if True == tmp3047 {
tmp3036 := MakeNative(func(__e *ControlFlow) {
W616 := __e.Get(1)
_ = W616
tmp3037 := MakeNative(func(__e *ControlFlow) {
W617 := __e.Get(1)
_ = W617
tmp3038 := PrimCons(W613, Nil)

tmp3039 := PrimCons(symshen_4choicepoint_b, tmp3038)

tmp3040 := PrimCons(tmp3039, Nil)

tmp3041 := PrimCons(W616, tmp3040)

tmp3042 := PrimCons(symwhere, tmp3041)

tmp3043 := Call(__e, PrimFunc(sym_8p), W610, tmp3042)


__e.TailApply(PrimFunc(symshen_4comb), W617, tmp3043)
return


}, 1)

tmp3044 := Call(__e, PrimFunc(symtail), W615)


__e.TailApply(tmp3037, tmp3044)
return


}, 1)

tmp3045 := Call(__e, PrimFunc(symhead), W615)


__e.TailApply(tmp3036, tmp3045)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3048 := Call(__e, PrimFunc(symtail), W614)


__e.TailApply(tmp3035, tmp3048)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3051 := Call(__e, PrimFunc(symtail), W612)


__e.TailApply(tmp3034, tmp3051)
return


}, 1)

tmp3052 := Call(__e, PrimFunc(symhead), W612)


__e.TailApply(tmp3033, tmp3052)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3055 := Call(__e, PrimFunc(symtail), W611)


__e.TailApply(tmp3032, tmp3055)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3058 := Call(__e, PrimFunc(symshen_4in_1_6), W609)


__e.TailApply(tmp3031, tmp3058)
return


}, 1)

tmp3059 := Call(__e, PrimFunc(symshen_4_5_1out), W609)


__e.TailApply(tmp3030, tmp3059)
return


}


}, 1)

tmp3062 := Call(__e, PrimFunc(symshen_4_5patterns_6), V590)


tmp3063 := Call(__e, tmp3029, tmp3062)


__e.TailApply(tmp3001, tmp3063)
return


} else {
__e.Return(W601)
return
}


}, 1)

tmp3066 := MakeNative(func(__e *ControlFlow) {
W602 := __e.Get(1)
_ = W602
tmp3083 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W602)


if True == tmp3083 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3067 := MakeNative(func(__e *ControlFlow) {
W603 := __e.Get(1)
_ = W603
tmp3068 := MakeNative(func(__e *ControlFlow) {
W604 := __e.Get(1)
_ = W604
tmp3079 := Call(__e, PrimFunc(symshen_4hds_a_2), W604, sym_1_6)


if True == tmp3079 {
tmp3069 := MakeNative(func(__e *ControlFlow) {
W605 := __e.Get(1)
_ = W605
tmp3076 := PrimIsPair(W605)

if True == tmp3076 {
tmp3070 := MakeNative(func(__e *ControlFlow) {
W606 := __e.Get(1)
_ = W606
tmp3071 := MakeNative(func(__e *ControlFlow) {
W607 := __e.Get(1)
_ = W607
tmp3072 := Call(__e, PrimFunc(sym_8p), W603, W606)


__e.TailApply(PrimFunc(symshen_4comb), W607, tmp3072)
return


}, 1)

tmp3073 := Call(__e, PrimFunc(symtail), W605)


__e.TailApply(tmp3071, tmp3073)
return


}, 1)

tmp3074 := Call(__e, PrimFunc(symhead), W605)


__e.TailApply(tmp3070, tmp3074)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3077 := Call(__e, PrimFunc(symtail), W604)


__e.TailApply(tmp3069, tmp3077)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3080 := Call(__e, PrimFunc(symshen_4in_1_6), W602)


__e.TailApply(tmp3068, tmp3080)
return


}, 1)

tmp3081 := Call(__e, PrimFunc(symshen_4_5_1out), W602)


__e.TailApply(tmp3067, tmp3081)
return


}


}, 1)

tmp3084 := Call(__e, PrimFunc(symshen_4_5patterns_6), V590)


tmp3085 := Call(__e, tmp3066, tmp3084)


__e.TailApply(tmp3000, tmp3085)
return


} else {
__e.Return(W591)
return
}


}, 1)

tmp3088 := MakeNative(func(__e *ControlFlow) {
W592 := __e.Get(1)
_ = W592
tmp3118 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W592)


if True == tmp3118 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3089 := MakeNative(func(__e *ControlFlow) {
W593 := __e.Get(1)
_ = W593
tmp3090 := MakeNative(func(__e *ControlFlow) {
W594 := __e.Get(1)
_ = W594
tmp3114 := Call(__e, PrimFunc(symshen_4hds_a_2), W594, sym_1_6)


if True == tmp3114 {
tmp3091 := MakeNative(func(__e *ControlFlow) {
W595 := __e.Get(1)
_ = W595
tmp3111 := PrimIsPair(W595)

if True == tmp3111 {
tmp3092 := MakeNative(func(__e *ControlFlow) {
W596 := __e.Get(1)
_ = W596
tmp3093 := MakeNative(func(__e *ControlFlow) {
W597 := __e.Get(1)
_ = W597
tmp3107 := Call(__e, PrimFunc(symshen_4hds_a_2), W597, symwhere)


if True == tmp3107 {
tmp3094 := MakeNative(func(__e *ControlFlow) {
W598 := __e.Get(1)
_ = W598
tmp3104 := PrimIsPair(W598)

if True == tmp3104 {
tmp3095 := MakeNative(func(__e *ControlFlow) {
W599 := __e.Get(1)
_ = W599
tmp3096 := MakeNative(func(__e *ControlFlow) {
W600 := __e.Get(1)
_ = W600
tmp3097 := PrimCons(W596, Nil)

tmp3098 := PrimCons(W599, tmp3097)

tmp3099 := PrimCons(symwhere, tmp3098)

tmp3100 := Call(__e, PrimFunc(sym_8p), W593, tmp3099)


__e.TailApply(PrimFunc(symshen_4comb), W600, tmp3100)
return


}, 1)

tmp3101 := Call(__e, PrimFunc(symtail), W598)


__e.TailApply(tmp3096, tmp3101)
return


}, 1)

tmp3102 := Call(__e, PrimFunc(symhead), W598)


__e.TailApply(tmp3095, tmp3102)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3105 := Call(__e, PrimFunc(symtail), W597)


__e.TailApply(tmp3094, tmp3105)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3108 := Call(__e, PrimFunc(symtail), W595)


__e.TailApply(tmp3093, tmp3108)
return


}, 1)

tmp3109 := Call(__e, PrimFunc(symhead), W595)


__e.TailApply(tmp3092, tmp3109)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3112 := Call(__e, PrimFunc(symtail), W594)


__e.TailApply(tmp3091, tmp3112)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3115 := Call(__e, PrimFunc(symshen_4in_1_6), W592)


__e.TailApply(tmp3090, tmp3115)
return


}, 1)

tmp3116 := Call(__e, PrimFunc(symshen_4_5_1out), W592)


__e.TailApply(tmp3089, tmp3116)
return


}


}, 1)

tmp3119 := Call(__e, PrimFunc(symshen_4_5patterns_6), V590)


tmp3120 := Call(__e, tmp3088, tmp3119)


__e.TailApply(tmp2999, tmp3120)
return


}, 1)

tmp3121 := Call(__e, ns2_1set, symshen_4_5rule_6, tmp2998)


_ = tmp3121

tmp3122 := MakeNative(func(__e *ControlFlow) {
V625 := __e.Get(1)
_ = V625
tmp3123 := MakeNative(func(__e *ControlFlow) {
W626 := __e.Get(1)
_ = W626
tmp3135 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W626)


if True == tmp3135 {
tmp3124 := MakeNative(func(__e *ControlFlow) {
W633 := __e.Get(1)
_ = W633
tmp3126 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W633)


if True == tmp3126 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W633)
return
}


}, 1)

tmp3127 := MakeNative(func(__e *ControlFlow) {
W634 := __e.Get(1)
_ = W634
tmp3131 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W634)


if True == tmp3131 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3128 := MakeNative(func(__e *ControlFlow) {
W635 := __e.Get(1)
_ = W635
__e.TailApply(PrimFunc(symshen_4comb), W635, Nil)
return
}, 1)

tmp3129 := Call(__e, PrimFunc(symshen_4in_1_6), W634)


__e.TailApply(tmp3128, tmp3129)
return


}


}, 1)

tmp3132 := Call(__e, PrimFunc(sym_5e_6), V625)


tmp3133 := Call(__e, tmp3127, tmp3132)


__e.TailApply(tmp3124, tmp3133)
return


} else {
__e.Return(W626)
return
}


}, 1)

tmp3136 := MakeNative(func(__e *ControlFlow) {
W627 := __e.Get(1)
_ = W627
tmp3151 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W627)


if True == tmp3151 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3137 := MakeNative(func(__e *ControlFlow) {
W628 := __e.Get(1)
_ = W628
tmp3138 := MakeNative(func(__e *ControlFlow) {
W629 := __e.Get(1)
_ = W629
tmp3139 := MakeNative(func(__e *ControlFlow) {
W630 := __e.Get(1)
_ = W630
tmp3146 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W630)


if True == tmp3146 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3140 := MakeNative(func(__e *ControlFlow) {
W631 := __e.Get(1)
_ = W631
tmp3141 := MakeNative(func(__e *ControlFlow) {
W632 := __e.Get(1)
_ = W632
tmp3142 := PrimCons(W628, W631)

__e.TailApply(PrimFunc(symshen_4comb), W632, tmp3142)
return


}, 1)

tmp3143 := Call(__e, PrimFunc(symshen_4in_1_6), W630)


__e.TailApply(tmp3141, tmp3143)
return


}, 1)

tmp3144 := Call(__e, PrimFunc(symshen_4_5_1out), W630)


__e.TailApply(tmp3140, tmp3144)
return


}


}, 1)

tmp3147 := Call(__e, PrimFunc(symshen_4_5patterns_6), W629)


__e.TailApply(tmp3139, tmp3147)
return


}, 1)

tmp3148 := Call(__e, PrimFunc(symshen_4in_1_6), W627)


__e.TailApply(tmp3138, tmp3148)
return


}, 1)

tmp3149 := Call(__e, PrimFunc(symshen_4_5_1out), W627)


__e.TailApply(tmp3137, tmp3149)
return


}


}, 1)

tmp3152 := Call(__e, PrimFunc(symshen_4_5pattern_6), V625)


tmp3153 := Call(__e, tmp3136, tmp3152)


__e.TailApply(tmp3123, tmp3153)
return


}, 1)

tmp3154 := Call(__e, ns2_1set, symshen_4_5patterns_6, tmp3122)


_ = tmp3154

tmp3155 := MakeNative(func(__e *ControlFlow) {
V636 := __e.Get(1)
_ = V636
tmp3156 := MakeNative(func(__e *ControlFlow) {
W637 := __e.Get(1)
_ = W637
tmp3211 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W637)


if True == tmp3211 {
tmp3157 := MakeNative(func(__e *ControlFlow) {
W651 := __e.Get(1)
_ = W651
tmp3185 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W651)


if True == tmp3185 {
tmp3158 := MakeNative(func(__e *ControlFlow) {
W658 := __e.Get(1)
_ = W658
tmp3172 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W658)


if True == tmp3172 {
tmp3159 := MakeNative(func(__e *ControlFlow) {
W661 := __e.Get(1)
_ = W661
tmp3161 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W661)


if True == tmp3161 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W661)
return
}


}, 1)

tmp3162 := MakeNative(func(__e *ControlFlow) {
W662 := __e.Get(1)
_ = W662
tmp3168 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W662)


if True == tmp3168 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3163 := MakeNative(func(__e *ControlFlow) {
W663 := __e.Get(1)
_ = W663
tmp3164 := MakeNative(func(__e *ControlFlow) {
W664 := __e.Get(1)
_ = W664
__e.TailApply(PrimFunc(symshen_4comb), W664, W663)
return
}, 1)

tmp3165 := Call(__e, PrimFunc(symshen_4in_1_6), W662)


__e.TailApply(tmp3164, tmp3165)
return


}, 1)

tmp3166 := Call(__e, PrimFunc(symshen_4_5_1out), W662)


__e.TailApply(tmp3163, tmp3166)
return


}


}, 1)

tmp3169 := Call(__e, PrimFunc(symshen_4_5simple_1pattern_6), V636)


tmp3170 := Call(__e, tmp3162, tmp3169)


__e.TailApply(tmp3159, tmp3170)
return


} else {
__e.Return(W658)
return
}


}, 1)

tmp3183 := PrimIsPair(V636)

var ifres3173 Obj

if True == tmp3183 {
tmp3174 := MakeNative(func(__e *ControlFlow) {
W659 := __e.Get(1)
_ = W659
tmp3175 := MakeNative(func(__e *ControlFlow) {
W660 := __e.Get(1)
_ = W660
tmp3178 := PrimIsPair(W659)

if True == tmp3178 {
tmp3176 := Call(__e, PrimFunc(symshen_4constructor_1error), W659)


__e.TailApply(PrimFunc(symshen_4comb), W660, tmp3176)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3179 := Call(__e, PrimFunc(symtail), V636)


__e.TailApply(tmp3175, tmp3179)
return


}, 1)

tmp3180 := Call(__e, PrimFunc(symhead), V636)


tmp3181 := Call(__e, tmp3174, tmp3180)


ifres3173 = tmp3181


} else {
tmp3182 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3173 = tmp3182


}

__e.TailApply(tmp3158, ifres3173)
return


} else {
__e.Return(W651)
return
}


}, 1)

tmp3209 := Call(__e, PrimFunc(symshen_4ccons_2), V636)


var ifres3186 Obj

if True == tmp3209 {
tmp3187 := MakeNative(func(__e *ControlFlow) {
W652 := __e.Get(1)
_ = W652
tmp3188 := MakeNative(func(__e *ControlFlow) {
W653 := __e.Get(1)
_ = W653
tmp3204 := Call(__e, PrimFunc(symshen_4hds_a_2), W652, symvector)


if True == tmp3204 {
tmp3189 := MakeNative(func(__e *ControlFlow) {
W654 := __e.Get(1)
_ = W654
tmp3201 := Call(__e, PrimFunc(symshen_4hds_a_2), W654, MakeNumber(0))


if True == tmp3201 {
tmp3190 := MakeNative(func(__e *ControlFlow) {
W655 := __e.Get(1)
_ = W655
tmp3191 := MakeNative(func(__e *ControlFlow) {
W656 := __e.Get(1)
_ = W656
tmp3197 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W656)


if True == tmp3197 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3192 := MakeNative(func(__e *ControlFlow) {
W657 := __e.Get(1)
_ = W657
tmp3193 := PrimCons(MakeNumber(0), Nil)

tmp3194 := PrimCons(symvector, tmp3193)

__e.TailApply(PrimFunc(symshen_4comb), W653, tmp3194)
return


}, 1)

tmp3195 := Call(__e, PrimFunc(symshen_4in_1_6), W656)


__e.TailApply(tmp3192, tmp3195)
return


}


}, 1)

tmp3198 := Call(__e, PrimFunc(sym_5end_6), W655)


__e.TailApply(tmp3191, tmp3198)
return


}, 1)

tmp3199 := Call(__e, PrimFunc(symtail), W654)


__e.TailApply(tmp3190, tmp3199)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3202 := Call(__e, PrimFunc(symtail), W652)


__e.TailApply(tmp3189, tmp3202)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3205 := Call(__e, PrimFunc(symtail), V636)


__e.TailApply(tmp3188, tmp3205)
return


}, 1)

tmp3206 := Call(__e, PrimFunc(symhead), V636)


tmp3207 := Call(__e, tmp3187, tmp3206)


ifres3186 = tmp3207


} else {
tmp3208 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3186 = tmp3208


}

__e.TailApply(tmp3157, ifres3186)
return


} else {
__e.Return(W637)
return
}


}, 1)

tmp3252 := Call(__e, PrimFunc(symshen_4ccons_2), V636)


var ifres3212 Obj

if True == tmp3252 {
tmp3213 := MakeNative(func(__e *ControlFlow) {
W638 := __e.Get(1)
_ = W638
tmp3214 := MakeNative(func(__e *ControlFlow) {
W639 := __e.Get(1)
_ = W639
tmp3215 := MakeNative(func(__e *ControlFlow) {
W640 := __e.Get(1)
_ = W640
tmp3246 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W640)


if True == tmp3246 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3216 := MakeNative(func(__e *ControlFlow) {
W641 := __e.Get(1)
_ = W641
tmp3217 := MakeNative(func(__e *ControlFlow) {
W642 := __e.Get(1)
_ = W642
tmp3218 := MakeNative(func(__e *ControlFlow) {
W643 := __e.Get(1)
_ = W643
tmp3241 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W643)


if True == tmp3241 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3219 := MakeNative(func(__e *ControlFlow) {
W644 := __e.Get(1)
_ = W644
tmp3220 := MakeNative(func(__e *ControlFlow) {
W645 := __e.Get(1)
_ = W645
tmp3221 := MakeNative(func(__e *ControlFlow) {
W646 := __e.Get(1)
_ = W646
tmp3236 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W646)


if True == tmp3236 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3222 := MakeNative(func(__e *ControlFlow) {
W647 := __e.Get(1)
_ = W647
tmp3223 := MakeNative(func(__e *ControlFlow) {
W648 := __e.Get(1)
_ = W648
tmp3224 := MakeNative(func(__e *ControlFlow) {
W649 := __e.Get(1)
_ = W649
tmp3231 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W649)


if True == tmp3231 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3225 := MakeNative(func(__e *ControlFlow) {
W650 := __e.Get(1)
_ = W650
tmp3226 := PrimCons(W647, Nil)

tmp3227 := PrimCons(W644, tmp3226)

tmp3228 := PrimCons(W641, tmp3227)

__e.TailApply(PrimFunc(symshen_4comb), W639, tmp3228)
return


}, 1)

tmp3229 := Call(__e, PrimFunc(symshen_4in_1_6), W649)


__e.TailApply(tmp3225, tmp3229)
return


}


}, 1)

tmp3232 := Call(__e, PrimFunc(sym_5end_6), W648)


__e.TailApply(tmp3224, tmp3232)
return


}, 1)

tmp3233 := Call(__e, PrimFunc(symshen_4in_1_6), W646)


__e.TailApply(tmp3223, tmp3233)
return


}, 1)

tmp3234 := Call(__e, PrimFunc(symshen_4_5_1out), W646)


__e.TailApply(tmp3222, tmp3234)
return


}


}, 1)

tmp3237 := Call(__e, PrimFunc(symshen_4_5pattern2_6), W645)


__e.TailApply(tmp3221, tmp3237)
return


}, 1)

tmp3238 := Call(__e, PrimFunc(symshen_4in_1_6), W643)


__e.TailApply(tmp3220, tmp3238)
return


}, 1)

tmp3239 := Call(__e, PrimFunc(symshen_4_5_1out), W643)


__e.TailApply(tmp3219, tmp3239)
return


}


}, 1)

tmp3242 := Call(__e, PrimFunc(symshen_4_5pattern1_6), W642)


__e.TailApply(tmp3218, tmp3242)
return


}, 1)

tmp3243 := Call(__e, PrimFunc(symshen_4in_1_6), W640)


__e.TailApply(tmp3217, tmp3243)
return


}, 1)

tmp3244 := Call(__e, PrimFunc(symshen_4_5_1out), W640)


__e.TailApply(tmp3216, tmp3244)
return


}


}, 1)

tmp3247 := Call(__e, PrimFunc(symshen_4_5constructor_6), W638)


__e.TailApply(tmp3215, tmp3247)
return


}, 1)

tmp3248 := Call(__e, PrimFunc(symtail), V636)


__e.TailApply(tmp3214, tmp3248)
return


}, 1)

tmp3249 := Call(__e, PrimFunc(symhead), V636)


tmp3250 := Call(__e, tmp3213, tmp3249)


ifres3212 = tmp3250


} else {
tmp3251 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3212 = tmp3251


}

__e.TailApply(tmp3156, ifres3212)
return


}, 1)

tmp3253 := Call(__e, ns2_1set, symshen_4_5pattern_6, tmp3155)


_ = tmp3253

tmp3254 := MakeNative(func(__e *ControlFlow) {
V665 := __e.Get(1)
_ = V665
tmp3255 := MakeNative(func(__e *ControlFlow) {
W666 := __e.Get(1)
_ = W666
tmp3257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W666)


if True == tmp3257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W666)
return
}


}, 1)

tmp3267 := PrimIsPair(V665)

var ifres3258 Obj

if True == tmp3267 {
tmp3259 := MakeNative(func(__e *ControlFlow) {
W667 := __e.Get(1)
_ = W667
tmp3260 := MakeNative(func(__e *ControlFlow) {
W668 := __e.Get(1)
_ = W668
tmp3262 := Call(__e, PrimFunc(symshen_4constructor_2), W667)


if True == tmp3262 {
__e.TailApply(PrimFunc(symshen_4comb), W668, W667)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3263 := Call(__e, PrimFunc(symtail), V665)


__e.TailApply(tmp3260, tmp3263)
return


}, 1)

tmp3264 := Call(__e, PrimFunc(symhead), V665)


tmp3265 := Call(__e, tmp3259, tmp3264)


ifres3258 = tmp3265


} else {
tmp3266 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3258 = tmp3266


}

__e.TailApply(tmp3255, ifres3258)
return


}, 1)

tmp3268 := Call(__e, ns2_1set, symshen_4_5constructor_6, tmp3254)


_ = tmp3268

tmp3269 := MakeNative(func(__e *ControlFlow) {
V669 := __e.Get(1)
_ = V669
tmp3270 := PrimCons(sym_8v, Nil)

tmp3271 := PrimCons(sym_8s, tmp3270)

tmp3272 := PrimCons(sym_8p, tmp3271)

tmp3273 := PrimCons(symcons, tmp3272)

__e.TailApply(PrimFunc(symelement_2), V669, tmp3273)
return


}, 1)

tmp3274 := Call(__e, ns2_1set, symshen_4constructor_2, tmp3269)


_ = tmp3274

tmp3275 := MakeNative(func(__e *ControlFlow) {
V670 := __e.Get(1)
_ = V670
tmp3276 := Call(__e, PrimFunc(symshen_4app), V670, MakeString(" is not a legitimate constructor\n"), symshen_4r)


__e.Return(PrimSimpleError(tmp3276))
return


}, 1)

tmp3277 := Call(__e, ns2_1set, symshen_4constructor_1error, tmp3275)


_ = tmp3277

tmp3278 := MakeNative(func(__e *ControlFlow) {
V671 := __e.Get(1)
_ = V671
tmp3279 := MakeNative(func(__e *ControlFlow) {
W672 := __e.Get(1)
_ = W672
tmp3297 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W672)


if True == tmp3297 {
tmp3280 := MakeNative(func(__e *ControlFlow) {
W675 := __e.Get(1)
_ = W675
tmp3282 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W675)


if True == tmp3282 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W675)
return
}


}, 1)

tmp3295 := PrimIsPair(V671)

var ifres3283 Obj

if True == tmp3295 {
tmp3284 := MakeNative(func(__e *ControlFlow) {
W676 := __e.Get(1)
_ = W676
tmp3285 := MakeNative(func(__e *ControlFlow) {
W677 := __e.Get(1)
_ = W677
tmp3287 := PrimCons(sym_5_1, Nil)

tmp3288 := PrimCons(sym_1_6, tmp3287)

tmp3289 := Call(__e, PrimFunc(symelement_2), W676, tmp3288)


tmp3290 := PrimNot(tmp3289)

if True == tmp3290 {
__e.TailApply(PrimFunc(symshen_4comb), W677, W676)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3291 := Call(__e, PrimFunc(symtail), V671)


__e.TailApply(tmp3285, tmp3291)
return


}, 1)

tmp3292 := Call(__e, PrimFunc(symhead), V671)


tmp3293 := Call(__e, tmp3284, tmp3292)


ifres3283 = tmp3293


} else {
tmp3294 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3283 = tmp3294


}

__e.TailApply(tmp3280, ifres3283)
return


} else {
__e.Return(W672)
return
}


}, 1)

tmp3308 := PrimIsPair(V671)

var ifres3298 Obj

if True == tmp3308 {
tmp3299 := MakeNative(func(__e *ControlFlow) {
W673 := __e.Get(1)
_ = W673
tmp3300 := MakeNative(func(__e *ControlFlow) {
W674 := __e.Get(1)
_ = W674
tmp3303 := PrimEqual(W673, sym__)

if True == tmp3303 {
tmp3301 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(PrimFunc(symshen_4comb), W674, tmp3301)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp3304 := Call(__e, PrimFunc(symtail), V671)


__e.TailApply(tmp3300, tmp3304)
return


}, 1)

tmp3305 := Call(__e, PrimFunc(symhead), V671)


tmp3306 := Call(__e, tmp3299, tmp3305)


ifres3298 = tmp3306


} else {
tmp3307 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres3298 = tmp3307


}

__e.TailApply(tmp3279, ifres3298)
return


}, 1)

tmp3309 := Call(__e, ns2_1set, symshen_4_5simple_1pattern_6, tmp3278)


_ = tmp3309

tmp3310 := MakeNative(func(__e *ControlFlow) {
V678 := __e.Get(1)
_ = V678
tmp3311 := MakeNative(func(__e *ControlFlow) {
W679 := __e.Get(1)
_ = W679
tmp3313 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W679)


if True == tmp3313 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W679)
return
}


}, 1)

tmp3314 := MakeNative(func(__e *ControlFlow) {
W680 := __e.Get(1)
_ = W680
tmp3320 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W680)


if True == tmp3320 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3315 := MakeNative(func(__e *ControlFlow) {
W681 := __e.Get(1)
_ = W681
tmp3316 := MakeNative(func(__e *ControlFlow) {
W682 := __e.Get(1)
_ = W682
__e.TailApply(PrimFunc(symshen_4comb), W682, W681)
return
}, 1)

tmp3317 := Call(__e, PrimFunc(symshen_4in_1_6), W680)


__e.TailApply(tmp3316, tmp3317)
return


}, 1)

tmp3318 := Call(__e, PrimFunc(symshen_4_5_1out), W680)


__e.TailApply(tmp3315, tmp3318)
return


}


}, 1)

tmp3321 := Call(__e, PrimFunc(symshen_4_5pattern_6), V678)


tmp3322 := Call(__e, tmp3314, tmp3321)


__e.TailApply(tmp3311, tmp3322)
return


}, 1)

tmp3323 := Call(__e, ns2_1set, symshen_4_5pattern1_6, tmp3310)


_ = tmp3323

tmp3324 := MakeNative(func(__e *ControlFlow) {
V683 := __e.Get(1)
_ = V683
tmp3325 := MakeNative(func(__e *ControlFlow) {
W684 := __e.Get(1)
_ = W684
tmp3327 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W684)


if True == tmp3327 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W684)
return
}


}, 1)

tmp3328 := MakeNative(func(__e *ControlFlow) {
W685 := __e.Get(1)
_ = W685
tmp3334 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W685)


if True == tmp3334 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp3329 := MakeNative(func(__e *ControlFlow) {
W686 := __e.Get(1)
_ = W686
tmp3330 := MakeNative(func(__e *ControlFlow) {
W687 := __e.Get(1)
_ = W687
__e.TailApply(PrimFunc(symshen_4comb), W687, W686)
return
}, 1)

tmp3331 := Call(__e, PrimFunc(symshen_4in_1_6), W685)


__e.TailApply(tmp3330, tmp3331)
return


}, 1)

tmp3332 := Call(__e, PrimFunc(symshen_4_5_1out), W685)


__e.TailApply(tmp3329, tmp3332)
return


}


}, 1)

tmp3335 := Call(__e, PrimFunc(symshen_4_5pattern_6), V683)


tmp3336 := Call(__e, tmp3328, tmp3335)


__e.TailApply(tmp3325, tmp3336)
return


}, 1)

tmp3337 := Call(__e, ns2_1set, symshen_4_5pattern2_6, tmp3324)


_ = tmp3337

tmp3338 := MakeNative(func(__e *ControlFlow) {
V688 := __e.Get(1)
_ = V688
tmp3339 := MakeNative(func(__e *ControlFlow) {
W689 := __e.Get(1)
_ = W689
tmp3340 := MakeNative(func(__e *ControlFlow) {
W690 := __e.Get(1)
_ = W690
tmp3341 := MakeNative(func(__e *ControlFlow) {
W691 := __e.Get(1)
_ = W691
__e.Return(W691)
return
}, 1)

tmp3342 := PrimStr(V688)

tmp3343 := Call(__e, PrimFunc(sym_8s), tmp3342, MakeString(")"))


tmp3344 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp3343)


tmp3345 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp3344)


tmp3346 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp3345)


tmp3347 := Call(__e, PrimFunc(sym_8s), MakeString("("), tmp3346)


tmp3348 := PrimVectorSet(W690, MakeNumber(1), tmp3347)

__e.TailApply(tmp3341, tmp3348)
return


}, 1)

tmp3349 := PrimVectorSet(W689, MakeNumber(0), symshen_4printF)

__e.TailApply(tmp3340, tmp3349)
return


}, 1)

tmp3350 := PrimAbsvector(MakeNumber(2))

__e.TailApply(tmp3339, tmp3350)
return


}, 1)

tmp3351 := Call(__e, ns2_1set, symshen_4fn_1print, tmp3338)


_ = tmp3351

tmp3352 := MakeNative(func(__e *ControlFlow) {
V692 := __e.Get(1)
_ = V692
__e.Return(PrimVectorGet(V692, MakeNumber(1)))
return
}, 1)

tmp3353 := Call(__e, ns2_1set, symshen_4printF, tmp3352)


_ = tmp3353

tmp3354 := MakeNative(func(__e *ControlFlow) {
V697 := __e.Get(1)
_ = V697
V698 := __e.Get(2)
_ = V698
tmp3378 := PrimIsPair(V698)

var ifres3374 Obj

if True == tmp3378 {
tmp3376 := PrimTail(V698)

tmp3377 := PrimEqual(Nil, tmp3376)

var ifres3375 Obj

if True == tmp3377 {
ifres3375 = True


} else {
ifres3375 = False


}

ifres3374 = ifres3375


} else {
ifres3374 = False


}

if True == ifres3374 {
tmp3355 := PrimHead(V698)

__e.TailApply(PrimFunc(symlength), tmp3355)
return


} else {
tmp3372 := PrimIsPair(V698)

var ifres3360 Obj

if True == tmp3372 {
tmp3370 := PrimTail(V698)

tmp3371 := PrimIsPair(tmp3370)

var ifres3362 Obj

if True == tmp3371 {
tmp3364 := PrimHead(V698)

tmp3365 := Call(__e, PrimFunc(symlength), tmp3364)


tmp3366 := PrimTail(V698)

tmp3367 := PrimHead(tmp3366)

tmp3368 := Call(__e, PrimFunc(symlength), tmp3367)


tmp3369 := PrimEqual(tmp3365, tmp3368)

var ifres3363 Obj

if True == tmp3369 {
ifres3363 = True


} else {
ifres3363 = False


}

ifres3362 = ifres3363


} else {
ifres3362 = False


}

var ifres3361 Obj

if True == ifres3362 {
ifres3361 = True


} else {
ifres3361 = False


}

ifres3360 = ifres3361


} else {
ifres3360 = False


}

if True == ifres3360 {
tmp3356 := PrimTail(V698)

__e.TailApply(PrimFunc(symshen_4arity_1chk), V697, tmp3356)
return


} else {
tmp3357 := Call(__e, PrimFunc(symshen_4app), V697, MakeString("\n"), symshen_4a)


tmp3358 := PrimStringConcat(MakeString("arity error in "), tmp3357)

__e.Return(PrimSimpleError(tmp3358))
return


}


}


}, 2)

tmp3379 := Call(__e, ns2_1set, symshen_4arity_1chk, tmp3354)


_ = tmp3379

tmp3380 := MakeNative(func(__e *ControlFlow) {
V699 := __e.Get(1)
_ = V699
V700 := __e.Get(2)
_ = V700
tmp3386 := Call(__e, PrimFunc(symtuple_2), V700)


if True == tmp3386 {
tmp3381 := Call(__e, PrimFunc(symfst), V700)


tmp3382 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3381)


tmp3383 := Call(__e, PrimFunc(symsnd), V700)


tmp3384 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3382, tmp3383)


__e.TailApply(PrimFunc(symshen_4free_1variable_1error_1message), V699, tmp3384)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.free-var-chk")))
return
}


}, 2)

tmp3387 := Call(__e, ns2_1set, symshen_4free_1var_1chk, tmp3380)


_ = tmp3387

tmp3388 := MakeNative(func(__e *ControlFlow) {
V701 := __e.Get(1)
_ = V701
V702 := __e.Get(2)
_ = V702
tmp3400 := Call(__e, PrimFunc(symempty_2), V702)


if True == tmp3400 {
__e.Return(symshen_4skip)
return
} else {
tmp3389 := Call(__e, PrimFunc(symshen_4app), V701, MakeString(":"), symshen_4a)


tmp3390 := PrimStringConcat(MakeString("free variables in "), tmp3389)

tmp3391 := Call(__e, PrimFunc(symstoutput))


tmp3392 := Call(__e, PrimFunc(sympr), tmp3390, tmp3391)


_ = tmp3392

tmp3393 := MakeNative(func(__e *ControlFlow) {
Z703 := __e.Get(1)
_ = Z703
tmp3394 := Call(__e, PrimFunc(symshen_4app), Z703, MakeString(""), symshen_4a)


tmp3395 := PrimStringConcat(MakeString(" "), tmp3394)

tmp3396 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp3395, tmp3396)
return


}, 1)

tmp3397 := Call(__e, PrimFunc(symmap), tmp3393, V702)


_ = tmp3397

tmp3398 := Call(__e, PrimFunc(symnl), MakeNumber(1))


_ = tmp3398

__e.TailApply(PrimFunc(symabort))
return


}


}, 2)

tmp3401 := Call(__e, ns2_1set, symshen_4free_1variable_1error_1message, tmp3388)


_ = tmp3401

tmp3402 := MakeNative(func(__e *ControlFlow) {
V706 := __e.Get(1)
_ = V706
tmp3410 := PrimIsVariable(V706)

if True == tmp3410 {
__e.Return(PrimCons(V706, Nil))
return
} else {
tmp3408 := PrimIsPair(V706)

if True == tmp3408 {
tmp3403 := PrimHead(V706)

tmp3404 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3403)


tmp3405 := PrimTail(V706)

tmp3406 := Call(__e, PrimFunc(symshen_4extract_1vars), tmp3405)


__e.TailApply(PrimFunc(symunion), tmp3404, tmp3406)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp3411 := Call(__e, ns2_1set, symshen_4extract_1vars, tmp3402)


_ = tmp3411

tmp3412 := MakeNative(func(__e *ControlFlow) {
V711 := __e.Get(1)
_ = V711
V712 := __e.Get(2)
_ = V712
tmp3502 := PrimIsPair(V712)

var ifres3489 Obj

if True == tmp3502 {
tmp3500 := PrimHead(V712)

tmp3501 := PrimEqual(symprotect, tmp3500)

var ifres3491 Obj

if True == tmp3501 {
tmp3498 := PrimTail(V712)

tmp3499 := PrimIsPair(tmp3498)

var ifres3493 Obj

if True == tmp3499 {
tmp3495 := PrimTail(V712)

tmp3496 := PrimTail(tmp3495)

tmp3497 := PrimEqual(Nil, tmp3496)

var ifres3494 Obj

if True == tmp3497 {
ifres3494 = True


} else {
ifres3494 = False


}

ifres3493 = ifres3494


} else {
ifres3493 = False


}

var ifres3492 Obj

if True == ifres3493 {
ifres3492 = True


} else {
ifres3492 = False


}

ifres3491 = ifres3492


} else {
ifres3491 = False


}

var ifres3490 Obj

if True == ifres3491 {
ifres3490 = True


} else {
ifres3490 = False


}

ifres3489 = ifres3490


} else {
ifres3489 = False


}

if True == ifres3489 {
__e.Return(Nil)
return
} else {
tmp3487 := PrimIsPair(V712)

var ifres3461 Obj

if True == tmp3487 {
tmp3485 := PrimHead(V712)

tmp3486 := PrimEqual(symlet, tmp3485)

var ifres3463 Obj

if True == tmp3486 {
tmp3483 := PrimTail(V712)

tmp3484 := PrimIsPair(tmp3483)

var ifres3465 Obj

if True == tmp3484 {
tmp3480 := PrimTail(V712)

tmp3481 := PrimTail(tmp3480)

tmp3482 := PrimIsPair(tmp3481)

var ifres3467 Obj

if True == tmp3482 {
tmp3476 := PrimTail(V712)

tmp3477 := PrimTail(tmp3476)

tmp3478 := PrimTail(tmp3477)

tmp3479 := PrimIsPair(tmp3478)

var ifres3469 Obj

if True == tmp3479 {
tmp3471 := PrimTail(V712)

tmp3472 := PrimTail(tmp3471)

tmp3473 := PrimTail(tmp3472)

tmp3474 := PrimTail(tmp3473)

tmp3475 := PrimEqual(Nil, tmp3474)

var ifres3470 Obj

if True == tmp3475 {
ifres3470 = True


} else {
ifres3470 = False


}

ifres3469 = ifres3470


} else {
ifres3469 = False


}

var ifres3468 Obj

if True == ifres3469 {
ifres3468 = True


} else {
ifres3468 = False


}

ifres3467 = ifres3468


} else {
ifres3467 = False


}

var ifres3466 Obj

if True == ifres3467 {
ifres3466 = True


} else {
ifres3466 = False


}

ifres3465 = ifres3466


} else {
ifres3465 = False


}

var ifres3464 Obj

if True == ifres3465 {
ifres3464 = True


} else {
ifres3464 = False


}

ifres3463 = ifres3464


} else {
ifres3463 = False


}

var ifres3462 Obj

if True == ifres3463 {
ifres3462 = True


} else {
ifres3462 = False


}

ifres3461 = ifres3462


} else {
ifres3461 = False


}

if True == ifres3461 {
tmp3413 := PrimTail(V712)

tmp3414 := PrimTail(tmp3413)

tmp3415 := PrimHead(tmp3414)

tmp3416 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V711, tmp3415)


tmp3417 := PrimTail(V712)

tmp3418 := PrimHead(tmp3417)

tmp3419 := PrimCons(tmp3418, V711)

tmp3420 := PrimTail(V712)

tmp3421 := PrimTail(tmp3420)

tmp3422 := PrimTail(tmp3421)

tmp3423 := PrimHead(tmp3422)

tmp3424 := Call(__e, PrimFunc(symshen_4find_1free_1vars), tmp3419, tmp3423)


__e.TailApply(PrimFunc(symunion), tmp3416, tmp3424)
return


} else {
tmp3459 := PrimIsPair(V712)

var ifres3440 Obj

if True == tmp3459 {
tmp3457 := PrimHead(V712)

tmp3458 := PrimEqual(symlambda, tmp3457)

var ifres3442 Obj

if True == tmp3458 {
tmp3455 := PrimTail(V712)

tmp3456 := PrimIsPair(tmp3455)

var ifres3444 Obj

if True == tmp3456 {
tmp3452 := PrimTail(V712)

tmp3453 := PrimTail(tmp3452)

tmp3454 := PrimIsPair(tmp3453)

var ifres3446 Obj

if True == tmp3454 {
tmp3448 := PrimTail(V712)

tmp3449 := PrimTail(tmp3448)

tmp3450 := PrimTail(tmp3449)

tmp3451 := PrimEqual(Nil, tmp3450)

var ifres3447 Obj

if True == tmp3451 {
ifres3447 = True


} else {
ifres3447 = False


}

ifres3446 = ifres3447


} else {
ifres3446 = False


}

var ifres3445 Obj

if True == ifres3446 {
ifres3445 = True


} else {
ifres3445 = False


}

ifres3444 = ifres3445


} else {
ifres3444 = False


}

var ifres3443 Obj

if True == ifres3444 {
ifres3443 = True


} else {
ifres3443 = False


}

ifres3442 = ifres3443


} else {
ifres3442 = False


}

var ifres3441 Obj

if True == ifres3442 {
ifres3441 = True


} else {
ifres3441 = False


}

ifres3440 = ifres3441


} else {
ifres3440 = False


}

if True == ifres3440 {
tmp3425 := PrimTail(V712)

tmp3426 := PrimHead(tmp3425)

tmp3427 := PrimCons(tmp3426, V711)

tmp3428 := PrimTail(V712)

tmp3429 := PrimTail(tmp3428)

tmp3430 := PrimHead(tmp3429)

__e.TailApply(PrimFunc(symshen_4find_1free_1vars), tmp3427, tmp3430)
return


} else {
tmp3438 := PrimIsPair(V712)

if True == tmp3438 {
tmp3431 := PrimHead(V712)

tmp3432 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V711, tmp3431)


tmp3433 := PrimTail(V712)

tmp3434 := Call(__e, PrimFunc(symshen_4find_1free_1vars), V711, tmp3433)


__e.TailApply(PrimFunc(symunion), tmp3432, tmp3434)
return


} else {
tmp3436 := Call(__e, PrimFunc(symshen_4free_1variable_2), V712, V711)


if True == tmp3436 {
__e.Return(PrimCons(V712, Nil))
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

tmp3503 := Call(__e, ns2_1set, symshen_4find_1free_1vars, tmp3412)


_ = tmp3503

tmp3504 := MakeNative(func(__e *ControlFlow) {
V713 := __e.Get(1)
_ = V713
V714 := __e.Get(2)
_ = V714
tmp3509 := PrimIsVariable(V713)

if True == tmp3509 {
tmp3506 := Call(__e, PrimFunc(symelement_2), V713, V714)


tmp3507 := PrimNot(tmp3506)

if True == tmp3507 {
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

tmp3510 := Call(__e, ns2_1set, symshen_4free_1variable_2, tmp3504)


_ = tmp3510

tmp3511 := MakeNative(func(__e *ControlFlow) {
V715 := __e.Get(1)
_ = V715
V716 := __e.Get(2)
_ = V716
tmp3512 := PrimValue(symshen_4_duserdefs_d)

tmp3513 := Call(__e, PrimFunc(symadjoin), V715, tmp3512)


tmp3514 := PrimSet(symshen_4_duserdefs_d, tmp3513)

_ = tmp3514

tmp3515 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V715, symshen_4source, V716, tmp3515)
return


}, 2)

tmp3516 := Call(__e, ns2_1set, symshen_4record_1kl, tmp3511)


_ = tmp3516

tmp3517 := MakeNative(func(__e *ControlFlow) {
V717 := __e.Get(1)
_ = V717
V718 := __e.Get(2)
_ = V718
V719 := __e.Get(3)
_ = V719
tmp3518 := MakeNative(func(__e *ControlFlow) {
W720 := __e.Get(1)
_ = W720
tmp3519 := MakeNative(func(__e *ControlFlow) {
W721 := __e.Get(1)
_ = W721
tmp3520 := MakeNative(func(__e *ControlFlow) {
W722 := __e.Get(1)
_ = W722
__e.Return(W722)
return
}, 1)

tmp3521 := Call(__e, PrimFunc(symshen_4cond_1form), W721)


tmp3522 := PrimCons(tmp3521, Nil)

tmp3523 := PrimCons(W720, tmp3522)

tmp3524 := PrimCons(V717, tmp3523)

tmp3525 := PrimCons(symdefun, tmp3524)

__e.TailApply(tmp3520, tmp3525)
return


}, 1)

tmp3526 := Call(__e, PrimFunc(symshen_4kl_1body), V718, W720)


tmp3527 := Call(__e, PrimFunc(symshen_4scan_1body), V717, tmp3526)


__e.TailApply(tmp3519, tmp3527)
return


}, 1)

tmp3528 := Call(__e, PrimFunc(symshen_4parameters), V719)


__e.TailApply(tmp3518, tmp3528)
return


}, 3)

tmp3529 := Call(__e, ns2_1set, symshen_4compile_1to_1kl, tmp3517)


_ = tmp3529

tmp3530 := MakeNative(func(__e *ControlFlow) {
V723 := __e.Get(1)
_ = V723
tmp3535 := PrimEqual(MakeNumber(0), V723)

if True == tmp3535 {
__e.Return(Nil)
return
} else {
tmp3531 := Call(__e, PrimFunc(symgensym), symV)


tmp3532 := PrimNumberSubtract(V723, MakeNumber(1))

tmp3533 := Call(__e, PrimFunc(symshen_4parameters), tmp3532)


__e.Return(PrimCons(tmp3531, tmp3533))
return


}


}, 1)

tmp3536 := Call(__e, ns2_1set, symshen_4parameters, tmp3530)


_ = tmp3536

tmp3537 := MakeNative(func(__e *ControlFlow) {
V726 := __e.Get(1)
_ = V726
tmp3561 := PrimIsPair(V726)

var ifres3541 Obj

if True == tmp3561 {
tmp3559 := PrimHead(V726)

tmp3560 := PrimIsPair(tmp3559)

var ifres3543 Obj

if True == tmp3560 {
tmp3556 := PrimHead(V726)

tmp3557 := PrimHead(tmp3556)

tmp3558 := PrimEqual(True, tmp3557)

var ifres3545 Obj

if True == tmp3558 {
tmp3553 := PrimHead(V726)

tmp3554 := PrimTail(tmp3553)

tmp3555 := PrimIsPair(tmp3554)

var ifres3547 Obj

if True == tmp3555 {
tmp3549 := PrimHead(V726)

tmp3550 := PrimTail(tmp3549)

tmp3551 := PrimTail(tmp3550)

tmp3552 := PrimEqual(Nil, tmp3551)

var ifres3548 Obj

if True == tmp3552 {
ifres3548 = True


} else {
ifres3548 = False


}

ifres3547 = ifres3548


} else {
ifres3547 = False


}

var ifres3546 Obj

if True == ifres3547 {
ifres3546 = True


} else {
ifres3546 = False


}

ifres3545 = ifres3546


} else {
ifres3545 = False


}

var ifres3544 Obj

if True == ifres3545 {
ifres3544 = True


} else {
ifres3544 = False


}

ifres3543 = ifres3544


} else {
ifres3543 = False


}

var ifres3542 Obj

if True == ifres3543 {
ifres3542 = True


} else {
ifres3542 = False


}

ifres3541 = ifres3542


} else {
ifres3541 = False


}

if True == ifres3541 {
tmp3538 := PrimHead(V726)

tmp3539 := PrimTail(tmp3538)

__e.Return(PrimHead(tmp3539))
return


} else {
__e.Return(PrimCons(symcond, V726))
return
}


}, 1)

tmp3562 := Call(__e, ns2_1set, symshen_4cond_1form, tmp3537)


_ = tmp3562

tmp3563 := MakeNative(func(__e *ControlFlow) {
V735 := __e.Get(1)
_ = V735
V736 := __e.Get(2)
_ = V736
tmp3607 := PrimEqual(Nil, V736)

if True == tmp3607 {
tmp3564 := PrimCons(V735, Nil)

tmp3565 := PrimCons(symshen_4f_1error, tmp3564)

tmp3566 := PrimCons(tmp3565, Nil)

tmp3567 := PrimCons(True, tmp3566)

__e.Return(PrimCons(tmp3567, Nil))
return


} else {
tmp3605 := PrimIsPair(V736)

var ifres3601 Obj

if True == tmp3605 {
tmp3603 := PrimHead(V736)

tmp3604 := Call(__e, PrimFunc(symshen_4choicepoint_2), tmp3603)


var ifres3602 Obj

if True == tmp3604 {
ifres3602 = True


} else {
ifres3602 = False


}

ifres3601 = ifres3602


} else {
ifres3601 = False


}

if True == ifres3601 {
tmp3568 := Call(__e, PrimFunc(symgensym), symFreeze)


tmp3569 := Call(__e, PrimFunc(symgensym), symResult)


tmp3570 := PrimHead(V736)

tmp3571 := PrimTail(V736)

__e.TailApply(PrimFunc(symshen_4choicepoint), V735, tmp3568, tmp3569, tmp3570, tmp3571)
return


} else {
tmp3599 := PrimIsPair(V736)

var ifres3579 Obj

if True == tmp3599 {
tmp3597 := PrimHead(V736)

tmp3598 := PrimIsPair(tmp3597)

var ifres3581 Obj

if True == tmp3598 {
tmp3594 := PrimHead(V736)

tmp3595 := PrimHead(tmp3594)

tmp3596 := PrimEqual(True, tmp3595)

var ifres3583 Obj

if True == tmp3596 {
tmp3591 := PrimHead(V736)

tmp3592 := PrimTail(tmp3591)

tmp3593 := PrimIsPair(tmp3592)

var ifres3585 Obj

if True == tmp3593 {
tmp3587 := PrimHead(V736)

tmp3588 := PrimTail(tmp3587)

tmp3589 := PrimTail(tmp3588)

tmp3590 := PrimEqual(Nil, tmp3589)

var ifres3586 Obj

if True == tmp3590 {
ifres3586 = True


} else {
ifres3586 = False


}

ifres3585 = ifres3586


} else {
ifres3585 = False


}

var ifres3584 Obj

if True == ifres3585 {
ifres3584 = True


} else {
ifres3584 = False


}

ifres3583 = ifres3584


} else {
ifres3583 = False


}

var ifres3582 Obj

if True == ifres3583 {
ifres3582 = True


} else {
ifres3582 = False


}

ifres3581 = ifres3582


} else {
ifres3581 = False


}

var ifres3580 Obj

if True == ifres3581 {
ifres3580 = True


} else {
ifres3580 = False


}

ifres3579 = ifres3580


} else {
ifres3579 = False


}

if True == ifres3579 {
tmp3572 := PrimHead(V736)

__e.Return(PrimCons(tmp3572, Nil))
return


} else {
tmp3577 := PrimIsPair(V736)

if True == tmp3577 {
tmp3573 := PrimHead(V736)

tmp3574 := PrimTail(V736)

tmp3575 := Call(__e, PrimFunc(symshen_4scan_1body), V735, tmp3574)


__e.Return(PrimCons(tmp3573, tmp3575))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.scan-body")))
return
}


}


}


}


}, 2)

tmp3608 := Call(__e, ns2_1set, symshen_4scan_1body, tmp3563)


_ = tmp3608

tmp3609 := MakeNative(func(__e *ControlFlow) {
V743 := __e.Get(1)
_ = V743
tmp3644 := PrimIsPair(V743)

var ifres3611 Obj

if True == tmp3644 {
tmp3642 := PrimTail(V743)

tmp3643 := PrimIsPair(tmp3642)

var ifres3613 Obj

if True == tmp3643 {
tmp3639 := PrimTail(V743)

tmp3640 := PrimHead(tmp3639)

tmp3641 := PrimIsPair(tmp3640)

var ifres3615 Obj

if True == tmp3641 {
tmp3635 := PrimTail(V743)

tmp3636 := PrimHead(tmp3635)

tmp3637 := PrimHead(tmp3636)

tmp3638 := PrimEqual(symshen_4choicepoint_b, tmp3637)

var ifres3617 Obj

if True == tmp3638 {
tmp3631 := PrimTail(V743)

tmp3632 := PrimHead(tmp3631)

tmp3633 := PrimTail(tmp3632)

tmp3634 := PrimIsPair(tmp3633)

var ifres3619 Obj

if True == tmp3634 {
tmp3626 := PrimTail(V743)

tmp3627 := PrimHead(tmp3626)

tmp3628 := PrimTail(tmp3627)

tmp3629 := PrimTail(tmp3628)

tmp3630 := PrimEqual(Nil, tmp3629)

var ifres3621 Obj

if True == tmp3630 {
tmp3623 := PrimTail(V743)

tmp3624 := PrimTail(tmp3623)

tmp3625 := PrimEqual(Nil, tmp3624)

var ifres3622 Obj

if True == tmp3625 {
ifres3622 = True


} else {
ifres3622 = False


}

ifres3621 = ifres3622


} else {
ifres3621 = False


}

var ifres3620 Obj

if True == ifres3621 {
ifres3620 = True


} else {
ifres3620 = False


}

ifres3619 = ifres3620


} else {
ifres3619 = False


}

var ifres3618 Obj

if True == ifres3619 {
ifres3618 = True


} else {
ifres3618 = False


}

ifres3617 = ifres3618


} else {
ifres3617 = False


}

var ifres3616 Obj

if True == ifres3617 {
ifres3616 = True


} else {
ifres3616 = False


}

ifres3615 = ifres3616


} else {
ifres3615 = False


}

var ifres3614 Obj

if True == ifres3615 {
ifres3614 = True


} else {
ifres3614 = False


}

ifres3613 = ifres3614


} else {
ifres3613 = False


}

var ifres3612 Obj

if True == ifres3613 {
ifres3612 = True


} else {
ifres3612 = False


}

ifres3611 = ifres3612


} else {
ifres3611 = False


}

if True == ifres3611 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp3645 := Call(__e, ns2_1set, symshen_4choicepoint_2, tmp3609)


_ = tmp3645

tmp3646 := MakeNative(func(__e *ControlFlow) {
V759 := __e.Get(1)
_ = V759
V760 := __e.Get(2)
_ = V760
V761 := __e.Get(3)
_ = V761
V762 := __e.Get(4)
_ = V762
V763 := __e.Get(5)
_ = V763
tmp3838 := PrimIsPair(V762)

var ifres3760 Obj

if True == tmp3838 {
tmp3836 := PrimTail(V762)

tmp3837 := PrimIsPair(tmp3836)

var ifres3762 Obj

if True == tmp3837 {
tmp3833 := PrimTail(V762)

tmp3834 := PrimHead(tmp3833)

tmp3835 := PrimIsPair(tmp3834)

var ifres3764 Obj

if True == tmp3835 {
tmp3829 := PrimTail(V762)

tmp3830 := PrimHead(tmp3829)

tmp3831 := PrimTail(tmp3830)

tmp3832 := PrimIsPair(tmp3831)

var ifres3766 Obj

if True == tmp3832 {
tmp3824 := PrimTail(V762)

tmp3825 := PrimHead(tmp3824)

tmp3826 := PrimTail(tmp3825)

tmp3827 := PrimHead(tmp3826)

tmp3828 := PrimIsPair(tmp3827)

var ifres3768 Obj

if True == tmp3828 {
tmp3818 := PrimTail(V762)

tmp3819 := PrimHead(tmp3818)

tmp3820 := PrimTail(tmp3819)

tmp3821 := PrimHead(tmp3820)

tmp3822 := PrimHead(tmp3821)

tmp3823 := PrimEqual(symfail_1if, tmp3822)

var ifres3770 Obj

if True == tmp3823 {
tmp3812 := PrimTail(V762)

tmp3813 := PrimHead(tmp3812)

tmp3814 := PrimTail(tmp3813)

tmp3815 := PrimHead(tmp3814)

tmp3816 := PrimTail(tmp3815)

tmp3817 := PrimIsPair(tmp3816)

var ifres3772 Obj

if True == tmp3817 {
tmp3805 := PrimTail(V762)

tmp3806 := PrimHead(tmp3805)

tmp3807 := PrimTail(tmp3806)

tmp3808 := PrimHead(tmp3807)

tmp3809 := PrimTail(tmp3808)

tmp3810 := PrimTail(tmp3809)

tmp3811 := PrimIsPair(tmp3810)

var ifres3774 Obj

if True == tmp3811 {
tmp3797 := PrimTail(V762)

tmp3798 := PrimHead(tmp3797)

tmp3799 := PrimTail(tmp3798)

tmp3800 := PrimHead(tmp3799)

tmp3801 := PrimTail(tmp3800)

tmp3802 := PrimTail(tmp3801)

tmp3803 := PrimTail(tmp3802)

tmp3804 := PrimEqual(Nil, tmp3803)

var ifres3776 Obj

if True == tmp3804 {
tmp3792 := PrimTail(V762)

tmp3793 := PrimHead(tmp3792)

tmp3794 := PrimTail(tmp3793)

tmp3795 := PrimTail(tmp3794)

tmp3796 := PrimEqual(Nil, tmp3795)

var ifres3778 Obj

if True == tmp3796 {
tmp3789 := PrimTail(V762)

tmp3790 := PrimTail(tmp3789)

tmp3791 := PrimEqual(Nil, tmp3790)

var ifres3780 Obj

if True == tmp3791 {
tmp3782 := PrimTail(V762)

tmp3783 := PrimHead(tmp3782)

tmp3784 := PrimTail(tmp3783)

tmp3785 := PrimHead(tmp3784)

tmp3786 := PrimTail(tmp3785)

tmp3787 := PrimHead(tmp3786)

tmp3788 := PrimEqual(V759, tmp3787)

var ifres3781 Obj

if True == tmp3788 {
ifres3781 = True


} else {
ifres3781 = False


}

ifres3780 = ifres3781


} else {
ifres3780 = False


}

var ifres3779 Obj

if True == ifres3780 {
ifres3779 = True


} else {
ifres3779 = False


}

ifres3778 = ifres3779


} else {
ifres3778 = False


}

var ifres3777 Obj

if True == ifres3778 {
ifres3777 = True


} else {
ifres3777 = False


}

ifres3776 = ifres3777


} else {
ifres3776 = False


}

var ifres3775 Obj

if True == ifres3776 {
ifres3775 = True


} else {
ifres3775 = False


}

ifres3774 = ifres3775


} else {
ifres3774 = False


}

var ifres3773 Obj

if True == ifres3774 {
ifres3773 = True


} else {
ifres3773 = False


}

ifres3772 = ifres3773


} else {
ifres3772 = False


}

var ifres3771 Obj

if True == ifres3772 {
ifres3771 = True


} else {
ifres3771 = False


}

ifres3770 = ifres3771


} else {
ifres3770 = False


}

var ifres3769 Obj

if True == ifres3770 {
ifres3769 = True


} else {
ifres3769 = False


}

ifres3768 = ifres3769


} else {
ifres3768 = False


}

var ifres3767 Obj

if True == ifres3768 {
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

if True == ifres3760 {
tmp3647 := PrimTail(V762)

tmp3648 := PrimHead(tmp3647)

tmp3649 := PrimTail(tmp3648)

tmp3650 := PrimHead(tmp3649)

tmp3651 := PrimTail(tmp3650)

tmp3652 := PrimHead(tmp3651)

tmp3653 := Call(__e, PrimFunc(symshen_4scan_1body), tmp3652, V763)


tmp3654 := PrimCons(symcond, tmp3653)

tmp3655 := PrimCons(tmp3654, Nil)

tmp3656 := PrimCons(symfreeze, tmp3655)

tmp3657 := PrimHead(V762)

tmp3658 := PrimTail(V762)

tmp3659 := PrimHead(tmp3658)

tmp3660 := PrimTail(tmp3659)

tmp3661 := PrimHead(tmp3660)

tmp3662 := PrimTail(tmp3661)

tmp3663 := PrimTail(tmp3662)

tmp3664 := PrimHead(tmp3663)

tmp3665 := PrimTail(V762)

tmp3666 := PrimHead(tmp3665)

tmp3667 := PrimTail(tmp3666)

tmp3668 := PrimHead(tmp3667)

tmp3669 := PrimTail(tmp3668)

tmp3670 := PrimHead(tmp3669)

tmp3671 := PrimCons(V761, Nil)

tmp3672 := PrimCons(tmp3670, tmp3671)

tmp3673 := PrimCons(V760, Nil)

tmp3674 := PrimCons(symthaw, tmp3673)

tmp3675 := PrimCons(V761, Nil)

tmp3676 := PrimCons(tmp3674, tmp3675)

tmp3677 := PrimCons(tmp3672, tmp3676)

tmp3678 := PrimCons(symif, tmp3677)

tmp3679 := PrimCons(tmp3678, Nil)

tmp3680 := PrimCons(tmp3664, tmp3679)

tmp3681 := PrimCons(V761, tmp3680)

tmp3682 := PrimCons(symlet, tmp3681)

tmp3683 := PrimCons(V760, Nil)

tmp3684 := PrimCons(symthaw, tmp3683)

tmp3685 := PrimCons(tmp3684, Nil)

tmp3686 := PrimCons(tmp3682, tmp3685)

tmp3687 := PrimCons(tmp3657, tmp3686)

tmp3688 := PrimCons(symif, tmp3687)

tmp3689 := PrimCons(tmp3688, Nil)

tmp3690 := PrimCons(tmp3656, tmp3689)

tmp3691 := PrimCons(V760, tmp3690)

tmp3692 := PrimCons(symlet, tmp3691)

tmp3693 := PrimCons(tmp3692, Nil)

tmp3694 := PrimCons(True, tmp3693)

__e.Return(PrimCons(tmp3694, Nil))
return


} else {
tmp3758 := PrimIsPair(V762)

var ifres3731 Obj

if True == tmp3758 {
tmp3756 := PrimTail(V762)

tmp3757 := PrimIsPair(tmp3756)

var ifres3733 Obj

if True == tmp3757 {
tmp3753 := PrimTail(V762)

tmp3754 := PrimHead(tmp3753)

tmp3755 := PrimIsPair(tmp3754)

var ifres3735 Obj

if True == tmp3755 {
tmp3749 := PrimTail(V762)

tmp3750 := PrimHead(tmp3749)

tmp3751 := PrimTail(tmp3750)

tmp3752 := PrimIsPair(tmp3751)

var ifres3737 Obj

if True == tmp3752 {
tmp3744 := PrimTail(V762)

tmp3745 := PrimHead(tmp3744)

tmp3746 := PrimTail(tmp3745)

tmp3747 := PrimTail(tmp3746)

tmp3748 := PrimEqual(Nil, tmp3747)

var ifres3739 Obj

if True == tmp3748 {
tmp3741 := PrimTail(V762)

tmp3742 := PrimTail(tmp3741)

tmp3743 := PrimEqual(Nil, tmp3742)

var ifres3740 Obj

if True == tmp3743 {
ifres3740 = True


} else {
ifres3740 = False


}

ifres3739 = ifres3740


} else {
ifres3739 = False


}

var ifres3738 Obj

if True == ifres3739 {
ifres3738 = True


} else {
ifres3738 = False


}

ifres3737 = ifres3738


} else {
ifres3737 = False


}

var ifres3736 Obj

if True == ifres3737 {
ifres3736 = True


} else {
ifres3736 = False


}

ifres3735 = ifres3736


} else {
ifres3735 = False


}

var ifres3734 Obj

if True == ifres3735 {
ifres3734 = True


} else {
ifres3734 = False


}

ifres3733 = ifres3734


} else {
ifres3733 = False


}

var ifres3732 Obj

if True == ifres3733 {
ifres3732 = True


} else {
ifres3732 = False


}

ifres3731 = ifres3732


} else {
ifres3731 = False


}

if True == ifres3731 {
tmp3695 := Call(__e, PrimFunc(symshen_4scan_1body), V759, V763)


tmp3696 := PrimCons(symcond, tmp3695)

tmp3697 := PrimCons(tmp3696, Nil)

tmp3698 := PrimCons(symfreeze, tmp3697)

tmp3699 := PrimHead(V762)

tmp3700 := PrimTail(V762)

tmp3701 := PrimHead(tmp3700)

tmp3702 := PrimTail(tmp3701)

tmp3703 := PrimHead(tmp3702)

tmp3704 := PrimCons(symfail, Nil)

tmp3705 := PrimCons(tmp3704, Nil)

tmp3706 := PrimCons(V761, tmp3705)

tmp3707 := PrimCons(sym_a, tmp3706)

tmp3708 := PrimCons(V760, Nil)

tmp3709 := PrimCons(symthaw, tmp3708)

tmp3710 := PrimCons(V761, Nil)

tmp3711 := PrimCons(tmp3709, tmp3710)

tmp3712 := PrimCons(tmp3707, tmp3711)

tmp3713 := PrimCons(symif, tmp3712)

tmp3714 := PrimCons(tmp3713, Nil)

tmp3715 := PrimCons(tmp3703, tmp3714)

tmp3716 := PrimCons(V761, tmp3715)

tmp3717 := PrimCons(symlet, tmp3716)

tmp3718 := PrimCons(V760, Nil)

tmp3719 := PrimCons(symthaw, tmp3718)

tmp3720 := PrimCons(tmp3719, Nil)

tmp3721 := PrimCons(tmp3717, tmp3720)

tmp3722 := PrimCons(tmp3699, tmp3721)

tmp3723 := PrimCons(symif, tmp3722)

tmp3724 := PrimCons(tmp3723, Nil)

tmp3725 := PrimCons(tmp3698, tmp3724)

tmp3726 := PrimCons(V760, tmp3725)

tmp3727 := PrimCons(symlet, tmp3726)

tmp3728 := PrimCons(tmp3727, Nil)

tmp3729 := PrimCons(True, tmp3728)

__e.Return(PrimCons(tmp3729, Nil))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.choicepoint")))
return
}


}


}, 5)

tmp3839 := Call(__e, ns2_1set, symshen_4choicepoint, tmp3646)


_ = tmp3839

tmp3840 := MakeNative(func(__e *ControlFlow) {
V765 := __e.Get(1)
_ = V765
V766 := __e.Get(2)
_ = V766
V767 := __e.Get(3)
_ = V767
tmp3854 := PrimEqual(V765, V767)

if True == tmp3854 {
__e.Return(V766)
return
} else {
tmp3852 := PrimIsPair(V767)

if True == tmp3852 {
tmp3841 := MakeNative(func(__e *ControlFlow) {
W768 := __e.Get(1)
_ = W768
tmp3847 := PrimHead(V767)

tmp3848 := PrimEqual(W768, tmp3847)

if True == tmp3848 {
tmp3842 := PrimHead(V767)

tmp3843 := PrimTail(V767)

tmp3844 := Call(__e, PrimFunc(symshen_4rep_1X), V765, V766, tmp3843)


__e.Return(PrimCons(tmp3842, tmp3844))
return


} else {
tmp3845 := PrimTail(V767)

__e.Return(PrimCons(W768, tmp3845))
return


}


}, 1)

tmp3849 := PrimHead(V767)

tmp3850 := Call(__e, PrimFunc(symshen_4rep_1X), V765, V766, tmp3849)


__e.TailApply(tmp3841, tmp3850)
return


} else {
__e.Return(V767)
return
}


}


}, 3)

tmp3855 := Call(__e, ns2_1set, symshen_4rep_1X, tmp3840)


_ = tmp3855

tmp3856 := MakeNative(func(__e *ControlFlow) {
V769 := __e.Get(1)
_ = V769
V770 := __e.Get(2)
_ = V770
tmp3857 := MakeNative(func(__e *ControlFlow) {
Z771 := __e.Get(1)
_ = Z771
tmp3858 := Call(__e, PrimFunc(symfst), Z771)


tmp3859 := Call(__e, PrimFunc(symsnd), Z771)


tmp3860 := Call(__e, PrimFunc(symshen_4alpha_1convert), tmp3859)


__e.TailApply(PrimFunc(symshen_4triple_1stack), Nil, tmp3858, V770, tmp3860)
return


}, 1)

__e.TailApply(PrimFunc(symmap), tmp3857, V769)
return


}, 2)

tmp3861 := Call(__e, ns2_1set, symshen_4kl_1body, tmp3856)


_ = tmp3861

tmp3862 := MakeNative(func(__e *ControlFlow) {
V772 := __e.Get(1)
_ = V772
tmp3945 := PrimIsPair(V772)

var ifres3926 Obj

if True == tmp3945 {
tmp3943 := PrimHead(V772)

tmp3944 := PrimEqual(symlambda, tmp3943)

var ifres3928 Obj

if True == tmp3944 {
tmp3941 := PrimTail(V772)

tmp3942 := PrimIsPair(tmp3941)

var ifres3930 Obj

if True == tmp3942 {
tmp3938 := PrimTail(V772)

tmp3939 := PrimTail(tmp3938)

tmp3940 := PrimIsPair(tmp3939)

var ifres3932 Obj

if True == tmp3940 {
tmp3934 := PrimTail(V772)

tmp3935 := PrimTail(tmp3934)

tmp3936 := PrimTail(tmp3935)

tmp3937 := PrimEqual(Nil, tmp3936)

var ifres3933 Obj

if True == tmp3937 {
ifres3933 = True


} else {
ifres3933 = False


}

ifres3932 = ifres3933


} else {
ifres3932 = False


}

var ifres3931 Obj

if True == ifres3932 {
ifres3931 = True


} else {
ifres3931 = False


}

ifres3930 = ifres3931


} else {
ifres3930 = False


}

var ifres3929 Obj

if True == ifres3930 {
ifres3929 = True


} else {
ifres3929 = False


}

ifres3928 = ifres3929


} else {
ifres3928 = False


}

var ifres3927 Obj

if True == ifres3928 {
ifres3927 = True


} else {
ifres3927 = False


}

ifres3926 = ifres3927


} else {
ifres3926 = False


}

if True == ifres3926 {
tmp3863 := MakeNative(func(__e *ControlFlow) {
W773 := __e.Get(1)
_ = W773
tmp3864 := MakeNative(func(__e *ControlFlow) {
W774 := __e.Get(1)
_ = W774
tmp3865 := MakeNative(func(__e *ControlFlow) {
Z775 := __e.Get(1)
_ = Z775
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z775)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3865, W774)
return


}, 1)

tmp3866 := PrimTail(V772)

tmp3867 := PrimHead(tmp3866)

tmp3868 := PrimTail(V772)

tmp3869 := PrimTail(tmp3868)

tmp3870 := PrimHead(tmp3869)

tmp3871 := Call(__e, PrimFunc(symshen_4beta), tmp3867, W773, tmp3870)


tmp3872 := PrimCons(tmp3871, Nil)

tmp3873 := PrimCons(W773, tmp3872)

tmp3874 := PrimCons(symlambda, tmp3873)

__e.TailApply(tmp3864, tmp3874)
return


}, 1)

tmp3875 := Call(__e, PrimFunc(symgensym), symZ)


__e.TailApply(tmp3863, tmp3875)
return


} else {
tmp3924 := PrimIsPair(V772)

var ifres3898 Obj

if True == tmp3924 {
tmp3922 := PrimHead(V772)

tmp3923 := PrimEqual(symlet, tmp3922)

var ifres3900 Obj

if True == tmp3923 {
tmp3920 := PrimTail(V772)

tmp3921 := PrimIsPair(tmp3920)

var ifres3902 Obj

if True == tmp3921 {
tmp3917 := PrimTail(V772)

tmp3918 := PrimTail(tmp3917)

tmp3919 := PrimIsPair(tmp3918)

var ifres3904 Obj

if True == tmp3919 {
tmp3913 := PrimTail(V772)

tmp3914 := PrimTail(tmp3913)

tmp3915 := PrimTail(tmp3914)

tmp3916 := PrimIsPair(tmp3915)

var ifres3906 Obj

if True == tmp3916 {
tmp3908 := PrimTail(V772)

tmp3909 := PrimTail(tmp3908)

tmp3910 := PrimTail(tmp3909)

tmp3911 := PrimTail(tmp3910)

tmp3912 := PrimEqual(Nil, tmp3911)

var ifres3907 Obj

if True == tmp3912 {
ifres3907 = True


} else {
ifres3907 = False


}

ifres3906 = ifres3907


} else {
ifres3906 = False


}

var ifres3905 Obj

if True == ifres3906 {
ifres3905 = True


} else {
ifres3905 = False


}

ifres3904 = ifres3905


} else {
ifres3904 = False


}

var ifres3903 Obj

if True == ifres3904 {
ifres3903 = True


} else {
ifres3903 = False


}

ifres3902 = ifres3903


} else {
ifres3902 = False


}

var ifres3901 Obj

if True == ifres3902 {
ifres3901 = True


} else {
ifres3901 = False


}

ifres3900 = ifres3901


} else {
ifres3900 = False


}

var ifres3899 Obj

if True == ifres3900 {
ifres3899 = True


} else {
ifres3899 = False


}

ifres3898 = ifres3899


} else {
ifres3898 = False


}

if True == ifres3898 {
tmp3876 := MakeNative(func(__e *ControlFlow) {
W776 := __e.Get(1)
_ = W776
tmp3877 := MakeNative(func(__e *ControlFlow) {
W777 := __e.Get(1)
_ = W777
tmp3878 := MakeNative(func(__e *ControlFlow) {
Z778 := __e.Get(1)
_ = Z778
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z778)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3878, W777)
return


}, 1)

tmp3879 := PrimTail(V772)

tmp3880 := PrimTail(tmp3879)

tmp3881 := PrimHead(tmp3880)

tmp3882 := PrimTail(V772)

tmp3883 := PrimHead(tmp3882)

tmp3884 := PrimTail(V772)

tmp3885 := PrimTail(tmp3884)

tmp3886 := PrimTail(tmp3885)

tmp3887 := PrimHead(tmp3886)

tmp3888 := Call(__e, PrimFunc(symshen_4beta), tmp3883, W776, tmp3887)


tmp3889 := PrimCons(tmp3888, Nil)

tmp3890 := PrimCons(tmp3881, tmp3889)

tmp3891 := PrimCons(W776, tmp3890)

tmp3892 := PrimCons(symlet, tmp3891)

__e.TailApply(tmp3877, tmp3892)
return


}, 1)

tmp3893 := Call(__e, PrimFunc(symgensym), symW)


__e.TailApply(tmp3876, tmp3893)
return


} else {
tmp3896 := PrimIsPair(V772)

if True == tmp3896 {
tmp3894 := MakeNative(func(__e *ControlFlow) {
Z779 := __e.Get(1)
_ = Z779
__e.TailApply(PrimFunc(symshen_4alpha_1convert), Z779)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp3894, V772)
return


} else {
__e.Return(V772)
return
}


}


}


}, 1)

tmp3946 := Call(__e, ns2_1set, symshen_4alpha_1convert, tmp3862)


_ = tmp3946

tmp3947 := MakeNative(func(__e *ControlFlow) {
V788 := __e.Get(1)
_ = V788
V789 := __e.Get(2)
_ = V789
V790 := __e.Get(3)
_ = V790
V791 := __e.Get(4)
_ = V791
tmp4077 := PrimEqual(Nil, V789)

var ifres4052 Obj

if True == tmp4077 {
tmp4076 := PrimEqual(Nil, V790)

var ifres4054 Obj

if True == tmp4076 {
tmp4075 := PrimIsPair(V791)

var ifres4056 Obj

if True == tmp4075 {
tmp4073 := PrimHead(V791)

tmp4074 := PrimEqual(symwhere, tmp4073)

var ifres4058 Obj

if True == tmp4074 {
tmp4071 := PrimTail(V791)

tmp4072 := PrimIsPair(tmp4071)

var ifres4060 Obj

if True == tmp4072 {
tmp4068 := PrimTail(V791)

tmp4069 := PrimTail(tmp4068)

tmp4070 := PrimIsPair(tmp4069)

var ifres4062 Obj

if True == tmp4070 {
tmp4064 := PrimTail(V791)

tmp4065 := PrimTail(tmp4064)

tmp4066 := PrimTail(tmp4065)

tmp4067 := PrimEqual(Nil, tmp4066)

var ifres4063 Obj

if True == tmp4067 {
ifres4063 = True


} else {
ifres4063 = False


}

ifres4062 = ifres4063


} else {
ifres4062 = False


}

var ifres4061 Obj

if True == ifres4062 {
ifres4061 = True


} else {
ifres4061 = False


}

ifres4060 = ifres4061


} else {
ifres4060 = False


}

var ifres4059 Obj

if True == ifres4060 {
ifres4059 = True


} else {
ifres4059 = False


}

ifres4058 = ifres4059


} else {
ifres4058 = False


}

var ifres4057 Obj

if True == ifres4058 {
ifres4057 = True


} else {
ifres4057 = False


}

ifres4056 = ifres4057


} else {
ifres4056 = False


}

var ifres4055 Obj

if True == ifres4056 {
ifres4055 = True


} else {
ifres4055 = False


}

ifres4054 = ifres4055


} else {
ifres4054 = False


}

var ifres4053 Obj

if True == ifres4054 {
ifres4053 = True


} else {
ifres4053 = False


}

ifres4052 = ifres4053


} else {
ifres4052 = False


}

if True == ifres4052 {
tmp3948 := PrimTail(V791)

tmp3949 := PrimHead(tmp3948)

tmp3950 := PrimCons(tmp3949, V788)

tmp3951 := PrimTail(V791)

tmp3952 := PrimTail(tmp3951)

tmp3953 := PrimHead(tmp3952)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3950, Nil, Nil, tmp3953)
return


} else {
tmp4050 := PrimEqual(Nil, V789)

var ifres4047 Obj

if True == tmp4050 {
tmp4049 := PrimEqual(Nil, V790)

var ifres4048 Obj

if True == tmp4049 {
ifres4048 = True


} else {
ifres4048 = False


}

ifres4047 = ifres4048


} else {
ifres4047 = False


}

if True == ifres4047 {
tmp3954 := Call(__e, PrimFunc(symreverse), V788)


tmp3955 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp3954)


tmp3956 := PrimCons(V791, Nil)

__e.Return(PrimCons(tmp3955, tmp3956))
return


} else {
tmp4045 := PrimIsPair(V789)

var ifres4038 Obj

if True == tmp4045 {
tmp4044 := PrimIsPair(V790)

var ifres4040 Obj

if True == tmp4044 {
tmp4042 := PrimHead(V789)

tmp4043 := PrimIsVariable(tmp4042)

var ifres4041 Obj

if True == tmp4043 {
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
tmp3957 := PrimTail(V789)

tmp3958 := PrimTail(V790)

tmp3959 := PrimHead(V789)

tmp3960 := PrimHead(V790)

tmp3961 := Call(__e, PrimFunc(symshen_4beta), tmp3959, tmp3960, V791)


__e.TailApply(PrimFunc(symshen_4triple_1stack), V788, tmp3957, tmp3958, tmp3961)
return


} else {
tmp4036 := PrimIsPair(V789)

var ifres4011 Obj

if True == tmp4036 {
tmp4034 := PrimHead(V789)

tmp4035 := PrimIsPair(tmp4034)

var ifres4013 Obj

if True == tmp4035 {
tmp4031 := PrimHead(V789)

tmp4032 := PrimTail(tmp4031)

tmp4033 := PrimIsPair(tmp4032)

var ifres4015 Obj

if True == tmp4033 {
tmp4027 := PrimHead(V789)

tmp4028 := PrimTail(tmp4027)

tmp4029 := PrimTail(tmp4028)

tmp4030 := PrimIsPair(tmp4029)

var ifres4017 Obj

if True == tmp4030 {
tmp4022 := PrimHead(V789)

tmp4023 := PrimTail(tmp4022)

tmp4024 := PrimTail(tmp4023)

tmp4025 := PrimTail(tmp4024)

tmp4026 := PrimEqual(Nil, tmp4025)

var ifres4019 Obj

if True == tmp4026 {
tmp4021 := PrimIsPair(V790)

var ifres4020 Obj

if True == tmp4021 {
ifres4020 = True


} else {
ifres4020 = False


}

ifres4019 = ifres4020


} else {
ifres4019 = False


}

var ifres4018 Obj

if True == ifres4019 {
ifres4018 = True


} else {
ifres4018 = False


}

ifres4017 = ifres4018


} else {
ifres4017 = False


}

var ifres4016 Obj

if True == ifres4017 {
ifres4016 = True


} else {
ifres4016 = False


}

ifres4015 = ifres4016


} else {
ifres4015 = False


}

var ifres4014 Obj

if True == ifres4015 {
ifres4014 = True


} else {
ifres4014 = False


}

ifres4013 = ifres4014


} else {
ifres4013 = False


}

var ifres4012 Obj

if True == ifres4013 {
ifres4012 = True


} else {
ifres4012 = False


}

ifres4011 = ifres4012


} else {
ifres4011 = False


}

if True == ifres4011 {
tmp3962 := PrimHead(V789)

tmp3963 := PrimHead(tmp3962)

tmp3964 := Call(__e, PrimFunc(symshen_4op_1test), tmp3963)


tmp3965 := PrimHead(V790)

tmp3966 := PrimCons(tmp3965, Nil)

tmp3967 := PrimCons(tmp3964, tmp3966)

tmp3968 := PrimCons(tmp3967, V788)

tmp3969 := PrimHead(V789)

tmp3970 := PrimTail(tmp3969)

tmp3971 := PrimHead(tmp3970)

tmp3972 := PrimHead(V789)

tmp3973 := PrimTail(tmp3972)

tmp3974 := PrimTail(tmp3973)

tmp3975 := PrimHead(tmp3974)

tmp3976 := PrimTail(V789)

tmp3977 := PrimCons(tmp3975, tmp3976)

tmp3978 := PrimCons(tmp3971, tmp3977)

tmp3979 := PrimHead(V789)

tmp3980 := PrimHead(tmp3979)

tmp3981 := Call(__e, PrimFunc(symshen_4op1), tmp3980)


tmp3982 := PrimHead(V790)

tmp3983 := PrimCons(tmp3982, Nil)

tmp3984 := PrimCons(tmp3981, tmp3983)

tmp3985 := PrimHead(V789)

tmp3986 := PrimHead(tmp3985)

tmp3987 := Call(__e, PrimFunc(symshen_4op2), tmp3986)


tmp3988 := PrimHead(V790)

tmp3989 := PrimCons(tmp3988, Nil)

tmp3990 := PrimCons(tmp3987, tmp3989)

tmp3991 := PrimTail(V790)

tmp3992 := PrimCons(tmp3990, tmp3991)

tmp3993 := PrimCons(tmp3984, tmp3992)

tmp3994 := PrimHead(V789)

tmp3995 := PrimHead(V790)

tmp3996 := Call(__e, PrimFunc(symshen_4beta), tmp3994, tmp3995, V791)


__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp3968, tmp3978, tmp3993, tmp3996)
return


} else {
tmp4009 := PrimIsPair(V789)

var ifres4006 Obj

if True == tmp4009 {
tmp4008 := PrimIsPair(V790)

var ifres4007 Obj

if True == tmp4008 {
ifres4007 = True


} else {
ifres4007 = False


}

ifres4006 = ifres4007


} else {
ifres4006 = False


}

if True == ifres4006 {
tmp3997 := PrimHead(V789)

tmp3998 := PrimHead(V790)

tmp3999 := PrimCons(tmp3998, Nil)

tmp4000 := PrimCons(tmp3997, tmp3999)

tmp4001 := PrimCons(sym_a, tmp4000)

tmp4002 := PrimCons(tmp4001, V788)

tmp4003 := PrimTail(V789)

tmp4004 := PrimTail(V790)

__e.TailApply(PrimFunc(symshen_4triple_1stack), tmp4002, tmp4003, tmp4004, V791)
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

tmp4078 := Call(__e, ns2_1set, symshen_4triple_1stack, tmp3947)


_ = tmp4078

tmp4079 := MakeNative(func(__e *ControlFlow) {
V794 := __e.Get(1)
_ = V794
tmp4098 := PrimEqual(Nil, V794)

if True == tmp4098 {
__e.Return(True)
return
} else {
tmp4096 := PrimIsPair(V794)

var ifres4092 Obj

if True == tmp4096 {
tmp4094 := PrimTail(V794)

tmp4095 := PrimEqual(Nil, tmp4094)

var ifres4093 Obj

if True == tmp4095 {
ifres4093 = True


} else {
ifres4093 = False


}

ifres4092 = ifres4093


} else {
ifres4092 = False


}

if True == ifres4092 {
__e.Return(PrimHead(V794))
return
} else {
tmp4090 := PrimIsPair(V794)

var ifres4086 Obj

if True == tmp4090 {
tmp4088 := PrimTail(V794)

tmp4089 := PrimIsPair(tmp4088)

var ifres4087 Obj

if True == tmp4089 {
ifres4087 = True


} else {
ifres4087 = False


}

ifres4086 = ifres4087


} else {
ifres4086 = False


}

if True == ifres4086 {
tmp4080 := PrimHead(V794)

tmp4081 := PrimTail(V794)

tmp4082 := Call(__e, PrimFunc(symshen_4rectify_1test), tmp4081)


tmp4083 := PrimCons(tmp4082, Nil)

tmp4084 := PrimCons(tmp4080, tmp4083)

__e.Return(PrimCons(symand, tmp4084))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.rectify-test")))
return
}


}


}


}, 1)

tmp4099 := Call(__e, ns2_1set, symshen_4rectify_1test, tmp4079)


_ = tmp4099

tmp4100 := MakeNative(func(__e *ControlFlow) {
V804 := __e.Get(1)
_ = V804
V805 := __e.Get(2)
_ = V805
V806 := __e.Get(3)
_ = V806
tmp4177 := PrimEqual(V804, V806)

if True == tmp4177 {
__e.Return(V805)
return
} else {
tmp4175 := PrimIsPair(V806)

var ifres4151 Obj

if True == tmp4175 {
tmp4173 := PrimHead(V806)

tmp4174 := PrimEqual(symlambda, tmp4173)

var ifres4153 Obj

if True == tmp4174 {
tmp4171 := PrimTail(V806)

tmp4172 := PrimIsPair(tmp4171)

var ifres4155 Obj

if True == tmp4172 {
tmp4168 := PrimTail(V806)

tmp4169 := PrimTail(tmp4168)

tmp4170 := PrimIsPair(tmp4169)

var ifres4157 Obj

if True == tmp4170 {
tmp4164 := PrimTail(V806)

tmp4165 := PrimTail(tmp4164)

tmp4166 := PrimTail(tmp4165)

tmp4167 := PrimEqual(Nil, tmp4166)

var ifres4159 Obj

if True == tmp4167 {
tmp4161 := PrimTail(V806)

tmp4162 := PrimHead(tmp4161)

tmp4163 := PrimEqual(V804, tmp4162)

var ifres4160 Obj

if True == tmp4163 {
ifres4160 = True


} else {
ifres4160 = False


}

ifres4159 = ifres4160


} else {
ifres4159 = False


}

var ifres4158 Obj

if True == ifres4159 {
ifres4158 = True


} else {
ifres4158 = False


}

ifres4157 = ifres4158


} else {
ifres4157 = False


}

var ifres4156 Obj

if True == ifres4157 {
ifres4156 = True


} else {
ifres4156 = False


}

ifres4155 = ifres4156


} else {
ifres4155 = False


}

var ifres4154 Obj

if True == ifres4155 {
ifres4154 = True


} else {
ifres4154 = False


}

ifres4153 = ifres4154


} else {
ifres4153 = False


}

var ifres4152 Obj

if True == ifres4153 {
ifres4152 = True


} else {
ifres4152 = False


}

ifres4151 = ifres4152


} else {
ifres4151 = False


}

if True == ifres4151 {
__e.Return(V806)
return
} else {
tmp4149 := PrimIsPair(V806)

var ifres4118 Obj

if True == tmp4149 {
tmp4147 := PrimHead(V806)

tmp4148 := PrimEqual(symlet, tmp4147)

var ifres4120 Obj

if True == tmp4148 {
tmp4145 := PrimTail(V806)

tmp4146 := PrimIsPair(tmp4145)

var ifres4122 Obj

if True == tmp4146 {
tmp4142 := PrimTail(V806)

tmp4143 := PrimTail(tmp4142)

tmp4144 := PrimIsPair(tmp4143)

var ifres4124 Obj

if True == tmp4144 {
tmp4138 := PrimTail(V806)

tmp4139 := PrimTail(tmp4138)

tmp4140 := PrimTail(tmp4139)

tmp4141 := PrimIsPair(tmp4140)

var ifres4126 Obj

if True == tmp4141 {
tmp4133 := PrimTail(V806)

tmp4134 := PrimTail(tmp4133)

tmp4135 := PrimTail(tmp4134)

tmp4136 := PrimTail(tmp4135)

tmp4137 := PrimEqual(Nil, tmp4136)

var ifres4128 Obj

if True == tmp4137 {
tmp4130 := PrimTail(V806)

tmp4131 := PrimHead(tmp4130)

tmp4132 := PrimEqual(V804, tmp4131)

var ifres4129 Obj

if True == tmp4132 {
ifres4129 = True


} else {
ifres4129 = False


}

ifres4128 = ifres4129


} else {
ifres4128 = False


}

var ifres4127 Obj

if True == ifres4128 {
ifres4127 = True


} else {
ifres4127 = False


}

ifres4126 = ifres4127


} else {
ifres4126 = False


}

var ifres4125 Obj

if True == ifres4126 {
ifres4125 = True


} else {
ifres4125 = False


}

ifres4124 = ifres4125


} else {
ifres4124 = False


}

var ifres4123 Obj

if True == ifres4124 {
ifres4123 = True


} else {
ifres4123 = False


}

ifres4122 = ifres4123


} else {
ifres4122 = False


}

var ifres4121 Obj

if True == ifres4122 {
ifres4121 = True


} else {
ifres4121 = False


}

ifres4120 = ifres4121


} else {
ifres4120 = False


}

var ifres4119 Obj

if True == ifres4120 {
ifres4119 = True


} else {
ifres4119 = False


}

ifres4118 = ifres4119


} else {
ifres4118 = False


}

if True == ifres4118 {
tmp4101 := PrimTail(V806)

tmp4102 := PrimHead(tmp4101)

tmp4103 := PrimTail(V806)

tmp4104 := PrimHead(tmp4103)

tmp4105 := PrimTail(V806)

tmp4106 := PrimTail(tmp4105)

tmp4107 := PrimHead(tmp4106)

tmp4108 := Call(__e, PrimFunc(symshen_4beta), tmp4104, V805, tmp4107)


tmp4109 := PrimTail(V806)

tmp4110 := PrimTail(tmp4109)

tmp4111 := PrimTail(tmp4110)

tmp4112 := PrimCons(tmp4108, tmp4111)

tmp4113 := PrimCons(tmp4102, tmp4112)

__e.Return(PrimCons(symlet, tmp4113))
return


} else {
tmp4116 := PrimIsPair(V806)

if True == tmp4116 {
tmp4114 := MakeNative(func(__e *ControlFlow) {
Z807 := __e.Get(1)
_ = Z807
__e.TailApply(PrimFunc(symshen_4beta), V804, V805, Z807)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4114, V806)
return


} else {
__e.Return(V806)
return
}


}


}


}


}, 3)

tmp4178 := Call(__e, ns2_1set, symshen_4beta, tmp4100)


_ = tmp4178

tmp4179 := MakeNative(func(__e *ControlFlow) {
V810 := __e.Get(1)
_ = V810
tmp4187 := PrimEqual(symcons, V810)

if True == tmp4187 {
__e.Return(symhd)
return
} else {
tmp4185 := PrimEqual(sym_8s, V810)

if True == tmp4185 {
__e.Return(symhdstr)
return
} else {
tmp4183 := PrimEqual(sym_8p, V810)

if True == tmp4183 {
__e.Return(symfst)
return
} else {
tmp4181 := PrimEqual(sym_8v, V810)

if True == tmp4181 {
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

tmp4188 := Call(__e, ns2_1set, symshen_4op1, tmp4179)


_ = tmp4188

tmp4189 := MakeNative(func(__e *ControlFlow) {
V813 := __e.Get(1)
_ = V813
tmp4197 := PrimEqual(symcons, V813)

if True == tmp4197 {
__e.Return(symtl)
return
} else {
tmp4195 := PrimEqual(sym_8s, V813)

if True == tmp4195 {
__e.Return(symtlstr)
return
} else {
tmp4193 := PrimEqual(sym_8p, V813)

if True == tmp4193 {
__e.Return(symsnd)
return
} else {
tmp4191 := PrimEqual(sym_8v, V813)

if True == tmp4191 {
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

tmp4198 := Call(__e, ns2_1set, symshen_4op2, tmp4189)


_ = tmp4198

tmp4199 := MakeNative(func(__e *ControlFlow) {
V816 := __e.Get(1)
_ = V816
tmp4207 := PrimEqual(symcons, V816)

if True == tmp4207 {
__e.Return(symcons_2)
return
} else {
tmp4205 := PrimEqual(sym_8s, V816)

if True == tmp4205 {
__e.Return(symshen_4_7string_2)
return
} else {
tmp4203 := PrimEqual(sym_8p, V816)

if True == tmp4203 {
__e.Return(symtuple_2)
return
} else {
tmp4201 := PrimEqual(sym_8v, V816)

if True == tmp4201 {
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

tmp4208 := Call(__e, ns2_1set, symshen_4op_1test, tmp4199)


_ = tmp4208

tmp4209 := MakeNative(func(__e *ControlFlow) {
V817 := __e.Get(1)
_ = V817
tmp4211 := PrimEqual(MakeString(""), V817)

if True == tmp4211 {
__e.Return(False)
return
} else {
__e.Return(PrimIsString(V817))
return
}


}, 1)

tmp4212 := Call(__e, ns2_1set, symshen_4_7string_2, tmp4209)


_ = tmp4212

tmp4213 := MakeNative(func(__e *ControlFlow) {
V818 := __e.Get(1)
_ = V818
tmp4215 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp4216 := PrimEqual(V818, tmp4215)

if True == tmp4216 {
__e.Return(False)
return
} else {
__e.TailApply(PrimFunc(symvector_2), V818)
return
}


}, 1)

tmp4217 := Call(__e, ns2_1set, symshen_4_7vector_2, tmp4213)


_ = tmp4217

tmp4218 := MakeNative(func(__e *ControlFlow) {
V821 := __e.Get(1)
_ = V821
tmp4222 := PrimEqual(sym_7, V821)

if True == tmp4222 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, True))
return
} else {
tmp4220 := PrimEqual(sym_1, V821)

if True == tmp4220 {
__e.Return(PrimSet(symshen_4_dfactorise_2_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("factorise expects a + or a -\n")))
return
}


}


}, 1)

tmp4223 := Call(__e, ns2_1set, symfactorise, tmp4218)


_ = tmp4223

tmp4224 := MakeNative(func(__e *ControlFlow) {
V822 := __e.Get(1)
_ = V822
tmp4226 := PrimValue(symshen_4_dfactorise_2_d)

if True == tmp4226 {
__e.TailApply(PrimFunc(symshen_4factor), V822)
return
} else {
__e.Return(V822)
return
}


}, 1)

tmp4227 := Call(__e, ns2_1set, symshen_4factorise_1code, tmp4224)


_ = tmp4227

tmp4228 := MakeNative(func(__e *ControlFlow) {
V823 := __e.Get(1)
_ = V823
tmp4285 := PrimIsPair(V823)

var ifres4244 Obj

if True == tmp4285 {
tmp4283 := PrimHead(V823)

tmp4284 := PrimEqual(symdefun, tmp4283)

var ifres4246 Obj

if True == tmp4284 {
tmp4281 := PrimTail(V823)

tmp4282 := PrimIsPair(tmp4281)

var ifres4248 Obj

if True == tmp4282 {
tmp4278 := PrimTail(V823)

tmp4279 := PrimTail(tmp4278)

tmp4280 := PrimIsPair(tmp4279)

var ifres4250 Obj

if True == tmp4280 {
tmp4274 := PrimTail(V823)

tmp4275 := PrimTail(tmp4274)

tmp4276 := PrimTail(tmp4275)

tmp4277 := PrimIsPair(tmp4276)

var ifres4252 Obj

if True == tmp4277 {
tmp4269 := PrimTail(V823)

tmp4270 := PrimTail(tmp4269)

tmp4271 := PrimTail(tmp4270)

tmp4272 := PrimHead(tmp4271)

tmp4273 := PrimIsPair(tmp4272)

var ifres4254 Obj

if True == tmp4273 {
tmp4263 := PrimTail(V823)

tmp4264 := PrimTail(tmp4263)

tmp4265 := PrimTail(tmp4264)

tmp4266 := PrimHead(tmp4265)

tmp4267 := PrimHead(tmp4266)

tmp4268 := PrimEqual(symcond, tmp4267)

var ifres4256 Obj

if True == tmp4268 {
tmp4258 := PrimTail(V823)

tmp4259 := PrimTail(tmp4258)

tmp4260 := PrimTail(tmp4259)

tmp4261 := PrimTail(tmp4260)

tmp4262 := PrimEqual(Nil, tmp4261)

var ifres4257 Obj

if True == tmp4262 {
ifres4257 = True


} else {
ifres4257 = False


}

ifres4256 = ifres4257


} else {
ifres4256 = False


}

var ifres4255 Obj

if True == ifres4256 {
ifres4255 = True


} else {
ifres4255 = False


}

ifres4254 = ifres4255


} else {
ifres4254 = False


}

var ifres4253 Obj

if True == ifres4254 {
ifres4253 = True


} else {
ifres4253 = False


}

ifres4252 = ifres4253


} else {
ifres4252 = False


}

var ifres4251 Obj

if True == ifres4252 {
ifres4251 = True


} else {
ifres4251 = False


}

ifres4250 = ifres4251


} else {
ifres4250 = False


}

var ifres4249 Obj

if True == ifres4250 {
ifres4249 = True


} else {
ifres4249 = False


}

ifres4248 = ifres4249


} else {
ifres4248 = False


}

var ifres4247 Obj

if True == ifres4248 {
ifres4247 = True


} else {
ifres4247 = False


}

ifres4246 = ifres4247


} else {
ifres4246 = False


}

var ifres4245 Obj

if True == ifres4246 {
ifres4245 = True


} else {
ifres4245 = False


}

ifres4244 = ifres4245


} else {
ifres4244 = False


}

if True == ifres4244 {
tmp4229 := PrimTail(V823)

tmp4230 := PrimHead(tmp4229)

tmp4231 := PrimTail(V823)

tmp4232 := PrimTail(tmp4231)

tmp4233 := PrimHead(tmp4232)

tmp4234 := PrimTail(V823)

tmp4235 := PrimTail(tmp4234)

tmp4236 := PrimTail(tmp4235)

tmp4237 := PrimHead(tmp4236)

tmp4238 := PrimTail(tmp4237)

tmp4239 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4238)


tmp4240 := PrimCons(tmp4239, Nil)

tmp4241 := PrimCons(tmp4233, tmp4240)

tmp4242 := PrimCons(tmp4230, tmp4241)

__e.Return(PrimCons(symdefun, tmp4242))
return


} else {
__e.Return(V823)
return
}


}, 1)

tmp4286 := Call(__e, ns2_1set, symshen_4factor, tmp4228)


_ = tmp4286

tmp4287 := MakeNative(func(__e *ControlFlow) {
V826 := __e.Get(1)
_ = V826
tmp4443 := PrimIsPair(V826)

var ifres4423 Obj

if True == tmp4443 {
tmp4441 := PrimHead(V826)

tmp4442 := PrimIsPair(tmp4441)

var ifres4425 Obj

if True == tmp4442 {
tmp4438 := PrimHead(V826)

tmp4439 := PrimHead(tmp4438)

tmp4440 := PrimEqual(True, tmp4439)

var ifres4427 Obj

if True == tmp4440 {
tmp4435 := PrimHead(V826)

tmp4436 := PrimTail(tmp4435)

tmp4437 := PrimIsPair(tmp4436)

var ifres4429 Obj

if True == tmp4437 {
tmp4431 := PrimHead(V826)

tmp4432 := PrimTail(tmp4431)

tmp4433 := PrimTail(tmp4432)

tmp4434 := PrimEqual(Nil, tmp4433)

var ifres4430 Obj

if True == tmp4434 {
ifres4430 = True


} else {
ifres4430 = False


}

ifres4429 = ifres4430


} else {
ifres4429 = False


}

var ifres4428 Obj

if True == ifres4429 {
ifres4428 = True


} else {
ifres4428 = False


}

ifres4427 = ifres4428


} else {
ifres4427 = False


}

var ifres4426 Obj

if True == ifres4427 {
ifres4426 = True


} else {
ifres4426 = False


}

ifres4425 = ifres4426


} else {
ifres4425 = False


}

var ifres4424 Obj

if True == ifres4425 {
ifres4424 = True


} else {
ifres4424 = False


}

ifres4423 = ifres4424


} else {
ifres4423 = False


}

if True == ifres4423 {
tmp4288 := PrimHead(V826)

tmp4289 := PrimTail(tmp4288)

__e.Return(PrimHead(tmp4289))
return


} else {
tmp4421 := PrimIsPair(V826)

var ifres4374 Obj

if True == tmp4421 {
tmp4419 := PrimHead(V826)

tmp4420 := PrimIsPair(tmp4419)

var ifres4376 Obj

if True == tmp4420 {
tmp4416 := PrimHead(V826)

tmp4417 := PrimHead(tmp4416)

tmp4418 := PrimIsPair(tmp4417)

var ifres4378 Obj

if True == tmp4418 {
tmp4412 := PrimHead(V826)

tmp4413 := PrimHead(tmp4412)

tmp4414 := PrimHead(tmp4413)

tmp4415 := PrimEqual(symand, tmp4414)

var ifres4380 Obj

if True == tmp4415 {
tmp4408 := PrimHead(V826)

tmp4409 := PrimHead(tmp4408)

tmp4410 := PrimTail(tmp4409)

tmp4411 := PrimIsPair(tmp4410)

var ifres4382 Obj

if True == tmp4411 {
tmp4403 := PrimHead(V826)

tmp4404 := PrimHead(tmp4403)

tmp4405 := PrimTail(tmp4404)

tmp4406 := PrimTail(tmp4405)

tmp4407 := PrimIsPair(tmp4406)

var ifres4384 Obj

if True == tmp4407 {
tmp4397 := PrimHead(V826)

tmp4398 := PrimHead(tmp4397)

tmp4399 := PrimTail(tmp4398)

tmp4400 := PrimTail(tmp4399)

tmp4401 := PrimTail(tmp4400)

tmp4402 := PrimEqual(Nil, tmp4401)

var ifres4386 Obj

if True == tmp4402 {
tmp4394 := PrimHead(V826)

tmp4395 := PrimTail(tmp4394)

tmp4396 := PrimIsPair(tmp4395)

var ifres4388 Obj

if True == tmp4396 {
tmp4390 := PrimHead(V826)

tmp4391 := PrimTail(tmp4390)

tmp4392 := PrimTail(tmp4391)

tmp4393 := PrimEqual(Nil, tmp4392)

var ifres4389 Obj

if True == tmp4393 {
ifres4389 = True


} else {
ifres4389 = False


}

ifres4388 = ifres4389


} else {
ifres4388 = False


}

var ifres4387 Obj

if True == ifres4388 {
ifres4387 = True


} else {
ifres4387 = False


}

ifres4386 = ifres4387


} else {
ifres4386 = False


}

var ifres4385 Obj

if True == ifres4386 {
ifres4385 = True


} else {
ifres4385 = False


}

ifres4384 = ifres4385


} else {
ifres4384 = False


}

var ifres4383 Obj

if True == ifres4384 {
ifres4383 = True


} else {
ifres4383 = False


}

ifres4382 = ifres4383


} else {
ifres4382 = False


}

var ifres4381 Obj

if True == ifres4382 {
ifres4381 = True


} else {
ifres4381 = False


}

ifres4380 = ifres4381


} else {
ifres4380 = False


}

var ifres4379 Obj

if True == ifres4380 {
ifres4379 = True


} else {
ifres4379 = False


}

ifres4378 = ifres4379


} else {
ifres4378 = False


}

var ifres4377 Obj

if True == ifres4378 {
ifres4377 = True


} else {
ifres4377 = False


}

ifres4376 = ifres4377


} else {
ifres4376 = False


}

var ifres4375 Obj

if True == ifres4376 {
ifres4375 = True


} else {
ifres4375 = False


}

ifres4374 = ifres4375


} else {
ifres4374 = False


}

if True == ifres4374 {
tmp4290 := MakeNative(func(__e *ControlFlow) {
W827 := __e.Get(1)
_ = W827
tmp4291 := MakeNative(func(__e *ControlFlow) {
W828 := __e.Get(1)
_ = W828
tmp4339 := Call(__e, PrimFunc(symshen_4bad_1pivot_2), W828)


if True == tmp4339 {
tmp4292 := PrimHead(V826)

tmp4293 := PrimHead(tmp4292)

tmp4294 := PrimHead(V826)

tmp4295 := PrimTail(tmp4294)

tmp4296 := PrimHead(tmp4295)

tmp4297 := PrimTail(V826)

tmp4298 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4297)


tmp4299 := PrimCons(tmp4298, Nil)

tmp4300 := PrimCons(tmp4296, tmp4299)

tmp4301 := PrimCons(tmp4293, tmp4300)

__e.Return(PrimCons(symif, tmp4301))
return


} else {
tmp4302 := MakeNative(func(__e *ControlFlow) {
W829 := __e.Get(1)
_ = W829
tmp4303 := MakeNative(func(__e *ControlFlow) {
W830 := __e.Get(1)
_ = W830
tmp4304 := MakeNative(func(__e *ControlFlow) {
W831 := __e.Get(1)
_ = W831
tmp4305 := MakeNative(func(__e *ControlFlow) {
W832 := __e.Get(1)
_ = W832
tmp4306 := MakeNative(func(__e *ControlFlow) {
W833 := __e.Get(1)
_ = W833
__e.TailApply(PrimFunc(symshen_4remove_1indirection), W833)
return
}, 1)

tmp4307 := PrimCons(W830, Nil)

tmp4308 := PrimCons(symfreeze, tmp4307)

tmp4309 := PrimHead(V826)

tmp4310 := PrimHead(tmp4309)

tmp4311 := PrimTail(tmp4310)

tmp4312 := PrimHead(tmp4311)

tmp4313 := PrimHead(V826)

tmp4314 := PrimHead(tmp4313)

tmp4315 := PrimTail(tmp4314)

tmp4316 := PrimHead(tmp4315)

tmp4317 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W832)


tmp4318 := Call(__e, PrimFunc(symshen_4factor_1selectors), tmp4316, tmp4317)


tmp4319 := PrimCons(W831, Nil)

tmp4320 := PrimCons(symthaw, tmp4319)

tmp4321 := PrimCons(tmp4320, Nil)

tmp4322 := PrimCons(tmp4318, tmp4321)

tmp4323 := PrimCons(tmp4312, tmp4322)

tmp4324 := PrimCons(symif, tmp4323)

tmp4325 := PrimCons(tmp4324, Nil)

tmp4326 := PrimCons(tmp4308, tmp4325)

tmp4327 := PrimCons(W831, tmp4326)

tmp4328 := PrimCons(symlet, tmp4327)

__e.TailApply(tmp4306, tmp4328)
return


}, 1)

tmp4329 := PrimCons(W831, Nil)

tmp4330 := PrimCons(symthaw, tmp4329)

tmp4331 := PrimCons(tmp4330, Nil)

tmp4332 := PrimCons(True, tmp4331)

tmp4333 := PrimCons(tmp4332, W828)

tmp4334 := Call(__e, PrimFunc(symreverse), tmp4333)


__e.TailApply(tmp4305, tmp4334)
return


}, 1)

tmp4335 := Call(__e, PrimFunc(symgensym), symGoTo)


__e.TailApply(tmp4304, tmp4335)
return


}, 1)

tmp4336 := Call(__e, PrimFunc(symshen_4factor_1recognisors), W829)


__e.TailApply(tmp4303, tmp4336)
return


}, 1)

tmp4337 := Call(__e, PrimFunc(symsnd), W827)


__e.TailApply(tmp4302, tmp4337)
return


}


}, 1)

tmp4340 := Call(__e, PrimFunc(symfst), W827)


__e.TailApply(tmp4291, tmp4340)
return


}, 1)

tmp4341 := PrimHead(V826)

tmp4342 := PrimHead(tmp4341)

tmp4343 := PrimTail(tmp4342)

tmp4344 := PrimHead(tmp4343)

tmp4345 := Call(__e, PrimFunc(symshen_4pivot_1on), tmp4344, V826, Nil)


__e.TailApply(tmp4290, tmp4345)
return


} else {
tmp4372 := PrimIsPair(V826)

var ifres4357 Obj

if True == tmp4372 {
tmp4370 := PrimHead(V826)

tmp4371 := PrimIsPair(tmp4370)

var ifres4359 Obj

if True == tmp4371 {
tmp4367 := PrimHead(V826)

tmp4368 := PrimTail(tmp4367)

tmp4369 := PrimIsPair(tmp4368)

var ifres4361 Obj

if True == tmp4369 {
tmp4363 := PrimHead(V826)

tmp4364 := PrimTail(tmp4363)

tmp4365 := PrimTail(tmp4364)

tmp4366 := PrimEqual(Nil, tmp4365)

var ifres4362 Obj

if True == tmp4366 {
ifres4362 = True


} else {
ifres4362 = False


}

ifres4361 = ifres4362


} else {
ifres4361 = False


}

var ifres4360 Obj

if True == ifres4361 {
ifres4360 = True


} else {
ifres4360 = False


}

ifres4359 = ifres4360


} else {
ifres4359 = False


}

var ifres4358 Obj

if True == ifres4359 {
ifres4358 = True


} else {
ifres4358 = False


}

ifres4357 = ifres4358


} else {
ifres4357 = False


}

if True == ifres4357 {
tmp4346 := PrimHead(V826)

tmp4347 := PrimHead(tmp4346)

tmp4348 := PrimHead(V826)

tmp4349 := PrimTail(tmp4348)

tmp4350 := PrimHead(tmp4349)

tmp4351 := PrimTail(V826)

tmp4352 := Call(__e, PrimFunc(symshen_4factor_1recognisors), tmp4351)


tmp4353 := PrimCons(tmp4352, Nil)

tmp4354 := PrimCons(tmp4350, tmp4353)

tmp4355 := PrimCons(tmp4347, tmp4354)

__e.Return(PrimCons(symif, tmp4355))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.factor-recognisors")))
return
}


}


}


}, 1)

tmp4444 := Call(__e, ns2_1set, symshen_4factor_1recognisors, tmp4287)


_ = tmp4444

tmp4445 := MakeNative(func(__e *ControlFlow) {
V838 := __e.Get(1)
_ = V838
tmp4451 := PrimIsPair(V838)

var ifres4447 Obj

if True == tmp4451 {
tmp4449 := PrimTail(V838)

tmp4450 := PrimEqual(Nil, tmp4449)

var ifres4448 Obj

if True == tmp4450 {
ifres4448 = True


} else {
ifres4448 = False


}

ifres4447 = ifres4448


} else {
ifres4447 = False


}

if True == ifres4447 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp4452 := Call(__e, ns2_1set, symshen_4bad_1pivot_2, tmp4445)


_ = tmp4452

tmp4453 := MakeNative(func(__e *ControlFlow) {
V839 := __e.Get(1)
_ = V839
tmp4568 := PrimIsPair(V839)

var ifres4468 Obj

if True == tmp4568 {
tmp4566 := PrimHead(V839)

tmp4567 := PrimEqual(symlet, tmp4566)

var ifres4470 Obj

if True == tmp4567 {
tmp4564 := PrimTail(V839)

tmp4565 := PrimIsPair(tmp4564)

var ifres4472 Obj

if True == tmp4565 {
tmp4561 := PrimTail(V839)

tmp4562 := PrimTail(tmp4561)

tmp4563 := PrimIsPair(tmp4562)

var ifres4474 Obj

if True == tmp4563 {
tmp4557 := PrimTail(V839)

tmp4558 := PrimTail(tmp4557)

tmp4559 := PrimHead(tmp4558)

tmp4560 := PrimIsPair(tmp4559)

var ifres4476 Obj

if True == tmp4560 {
tmp4552 := PrimTail(V839)

tmp4553 := PrimTail(tmp4552)

tmp4554 := PrimHead(tmp4553)

tmp4555 := PrimHead(tmp4554)

tmp4556 := PrimEqual(symfreeze, tmp4555)

var ifres4478 Obj

if True == tmp4556 {
tmp4547 := PrimTail(V839)

tmp4548 := PrimTail(tmp4547)

tmp4549 := PrimHead(tmp4548)

tmp4550 := PrimTail(tmp4549)

tmp4551 := PrimIsPair(tmp4550)

var ifres4480 Obj

if True == tmp4551 {
tmp4541 := PrimTail(V839)

tmp4542 := PrimTail(tmp4541)

tmp4543 := PrimHead(tmp4542)

tmp4544 := PrimTail(tmp4543)

tmp4545 := PrimHead(tmp4544)

tmp4546 := PrimIsPair(tmp4545)

var ifres4482 Obj

if True == tmp4546 {
tmp4534 := PrimTail(V839)

tmp4535 := PrimTail(tmp4534)

tmp4536 := PrimHead(tmp4535)

tmp4537 := PrimTail(tmp4536)

tmp4538 := PrimHead(tmp4537)

tmp4539 := PrimHead(tmp4538)

tmp4540 := PrimEqual(symthaw, tmp4539)

var ifres4484 Obj

if True == tmp4540 {
tmp4527 := PrimTail(V839)

tmp4528 := PrimTail(tmp4527)

tmp4529 := PrimHead(tmp4528)

tmp4530 := PrimTail(tmp4529)

tmp4531 := PrimHead(tmp4530)

tmp4532 := PrimTail(tmp4531)

tmp4533 := PrimIsPair(tmp4532)

var ifres4486 Obj

if True == tmp4533 {
tmp4519 := PrimTail(V839)

tmp4520 := PrimTail(tmp4519)

tmp4521 := PrimHead(tmp4520)

tmp4522 := PrimTail(tmp4521)

tmp4523 := PrimHead(tmp4522)

tmp4524 := PrimTail(tmp4523)

tmp4525 := PrimTail(tmp4524)

tmp4526 := PrimEqual(Nil, tmp4525)

var ifres4488 Obj

if True == tmp4526 {
tmp4513 := PrimTail(V839)

tmp4514 := PrimTail(tmp4513)

tmp4515 := PrimHead(tmp4514)

tmp4516 := PrimTail(tmp4515)

tmp4517 := PrimTail(tmp4516)

tmp4518 := PrimEqual(Nil, tmp4517)

var ifres4490 Obj

if True == tmp4518 {
tmp4509 := PrimTail(V839)

tmp4510 := PrimTail(tmp4509)

tmp4511 := PrimTail(tmp4510)

tmp4512 := PrimIsPair(tmp4511)

var ifres4492 Obj

if True == tmp4512 {
tmp4504 := PrimTail(V839)

tmp4505 := PrimTail(tmp4504)

tmp4506 := PrimTail(tmp4505)

tmp4507 := PrimTail(tmp4506)

tmp4508 := PrimEqual(Nil, tmp4507)

var ifres4494 Obj

if True == tmp4508 {
tmp4496 := PrimTail(V839)

tmp4497 := PrimTail(tmp4496)

tmp4498 := PrimHead(tmp4497)

tmp4499 := PrimTail(tmp4498)

tmp4500 := PrimHead(tmp4499)

tmp4501 := PrimTail(tmp4500)

tmp4502 := PrimHead(tmp4501)

tmp4503 := PrimIsSymbol(tmp4502)

var ifres4495 Obj

if True == tmp4503 {
ifres4495 = True


} else {
ifres4495 = False


}

ifres4494 = ifres4495


} else {
ifres4494 = False


}

var ifres4493 Obj

if True == ifres4494 {
ifres4493 = True


} else {
ifres4493 = False


}

ifres4492 = ifres4493


} else {
ifres4492 = False


}

var ifres4491 Obj

if True == ifres4492 {
ifres4491 = True


} else {
ifres4491 = False


}

ifres4490 = ifres4491


} else {
ifres4490 = False


}

var ifres4489 Obj

if True == ifres4490 {
ifres4489 = True


} else {
ifres4489 = False


}

ifres4488 = ifres4489


} else {
ifres4488 = False


}

var ifres4487 Obj

if True == ifres4488 {
ifres4487 = True


} else {
ifres4487 = False


}

ifres4486 = ifres4487


} else {
ifres4486 = False


}

var ifres4485 Obj

if True == ifres4486 {
ifres4485 = True


} else {
ifres4485 = False


}

ifres4484 = ifres4485


} else {
ifres4484 = False


}

var ifres4483 Obj

if True == ifres4484 {
ifres4483 = True


} else {
ifres4483 = False


}

ifres4482 = ifres4483


} else {
ifres4482 = False


}

var ifres4481 Obj

if True == ifres4482 {
ifres4481 = True


} else {
ifres4481 = False


}

ifres4480 = ifres4481


} else {
ifres4480 = False


}

var ifres4479 Obj

if True == ifres4480 {
ifres4479 = True


} else {
ifres4479 = False


}

ifres4478 = ifres4479


} else {
ifres4478 = False


}

var ifres4477 Obj

if True == ifres4478 {
ifres4477 = True


} else {
ifres4477 = False


}

ifres4476 = ifres4477


} else {
ifres4476 = False


}

var ifres4475 Obj

if True == ifres4476 {
ifres4475 = True


} else {
ifres4475 = False


}

ifres4474 = ifres4475


} else {
ifres4474 = False


}

var ifres4473 Obj

if True == ifres4474 {
ifres4473 = True


} else {
ifres4473 = False


}

ifres4472 = ifres4473


} else {
ifres4472 = False


}

var ifres4471 Obj

if True == ifres4472 {
ifres4471 = True


} else {
ifres4471 = False


}

ifres4470 = ifres4471


} else {
ifres4470 = False


}

var ifres4469 Obj

if True == ifres4470 {
ifres4469 = True


} else {
ifres4469 = False


}

ifres4468 = ifres4469


} else {
ifres4468 = False


}

if True == ifres4468 {
tmp4454 := PrimTail(V839)

tmp4455 := PrimTail(tmp4454)

tmp4456 := PrimHead(tmp4455)

tmp4457 := PrimTail(tmp4456)

tmp4458 := PrimHead(tmp4457)

tmp4459 := PrimTail(tmp4458)

tmp4460 := PrimHead(tmp4459)

tmp4461 := PrimTail(V839)

tmp4462 := PrimHead(tmp4461)

tmp4463 := PrimTail(V839)

tmp4464 := PrimTail(tmp4463)

tmp4465 := PrimTail(tmp4464)

tmp4466 := PrimHead(tmp4465)

__e.TailApply(PrimFunc(symsubst), tmp4460, tmp4462, tmp4466)
return


} else {
__e.Return(V839)
return
}


}, 1)

tmp4569 := Call(__e, ns2_1set, symshen_4remove_1indirection, tmp4453)


_ = tmp4569

tmp4570 := MakeNative(func(__e *ControlFlow) {
V842 := __e.Get(1)
_ = V842
V843 := __e.Get(2)
_ = V843
V844 := __e.Get(3)
_ = V844
tmp4669 := PrimIsPair(V843)

var ifres4615 Obj

if True == tmp4669 {
tmp4667 := PrimHead(V843)

tmp4668 := PrimIsPair(tmp4667)

var ifres4617 Obj

if True == tmp4668 {
tmp4664 := PrimHead(V843)

tmp4665 := PrimHead(tmp4664)

tmp4666 := PrimIsPair(tmp4665)

var ifres4619 Obj

if True == tmp4666 {
tmp4660 := PrimHead(V843)

tmp4661 := PrimHead(tmp4660)

tmp4662 := PrimHead(tmp4661)

tmp4663 := PrimEqual(symand, tmp4662)

var ifres4621 Obj

if True == tmp4663 {
tmp4656 := PrimHead(V843)

tmp4657 := PrimHead(tmp4656)

tmp4658 := PrimTail(tmp4657)

tmp4659 := PrimIsPair(tmp4658)

var ifres4623 Obj

if True == tmp4659 {
tmp4651 := PrimHead(V843)

tmp4652 := PrimHead(tmp4651)

tmp4653 := PrimTail(tmp4652)

tmp4654 := PrimTail(tmp4653)

tmp4655 := PrimIsPair(tmp4654)

var ifres4625 Obj

if True == tmp4655 {
tmp4645 := PrimHead(V843)

tmp4646 := PrimHead(tmp4645)

tmp4647 := PrimTail(tmp4646)

tmp4648 := PrimTail(tmp4647)

tmp4649 := PrimTail(tmp4648)

tmp4650 := PrimEqual(Nil, tmp4649)

var ifres4627 Obj

if True == tmp4650 {
tmp4642 := PrimHead(V843)

tmp4643 := PrimTail(tmp4642)

tmp4644 := PrimIsPair(tmp4643)

var ifres4629 Obj

if True == tmp4644 {
tmp4638 := PrimHead(V843)

tmp4639 := PrimTail(tmp4638)

tmp4640 := PrimTail(tmp4639)

tmp4641 := PrimEqual(Nil, tmp4640)

var ifres4631 Obj

if True == tmp4641 {
tmp4633 := PrimHead(V843)

tmp4634 := PrimHead(tmp4633)

tmp4635 := PrimTail(tmp4634)

tmp4636 := PrimHead(tmp4635)

tmp4637 := PrimEqual(V842, tmp4636)

var ifres4632 Obj

if True == tmp4637 {
ifres4632 = True


} else {
ifres4632 = False


}

ifres4631 = ifres4632


} else {
ifres4631 = False


}

var ifres4630 Obj

if True == ifres4631 {
ifres4630 = True


} else {
ifres4630 = False


}

ifres4629 = ifres4630


} else {
ifres4629 = False


}

var ifres4628 Obj

if True == ifres4629 {
ifres4628 = True


} else {
ifres4628 = False


}

ifres4627 = ifres4628


} else {
ifres4627 = False


}

var ifres4626 Obj

if True == ifres4627 {
ifres4626 = True


} else {
ifres4626 = False


}

ifres4625 = ifres4626


} else {
ifres4625 = False


}

var ifres4624 Obj

if True == ifres4625 {
ifres4624 = True


} else {
ifres4624 = False


}

ifres4623 = ifres4624


} else {
ifres4623 = False


}

var ifres4622 Obj

if True == ifres4623 {
ifres4622 = True


} else {
ifres4622 = False


}

ifres4621 = ifres4622


} else {
ifres4621 = False


}

var ifres4620 Obj

if True == ifres4621 {
ifres4620 = True


} else {
ifres4620 = False


}

ifres4619 = ifres4620


} else {
ifres4619 = False


}

var ifres4618 Obj

if True == ifres4619 {
ifres4618 = True


} else {
ifres4618 = False


}

ifres4617 = ifres4618


} else {
ifres4617 = False


}

var ifres4616 Obj

if True == ifres4617 {
ifres4616 = True


} else {
ifres4616 = False


}

ifres4615 = ifres4616


} else {
ifres4615 = False


}

if True == ifres4615 {
tmp4571 := PrimHead(V843)

tmp4572 := PrimHead(tmp4571)

tmp4573 := PrimTail(tmp4572)

tmp4574 := PrimHead(tmp4573)

tmp4575 := PrimTail(V843)

tmp4576 := PrimHead(V843)

tmp4577 := PrimHead(tmp4576)

tmp4578 := PrimTail(tmp4577)

tmp4579 := PrimTail(tmp4578)

tmp4580 := PrimHead(tmp4579)

tmp4581 := PrimHead(V843)

tmp4582 := PrimTail(tmp4581)

tmp4583 := PrimCons(tmp4580, tmp4582)

tmp4584 := PrimCons(tmp4583, V844)

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4574, tmp4575, tmp4584)
return


} else {
tmp4613 := PrimIsPair(V843)

var ifres4593 Obj

if True == tmp4613 {
tmp4611 := PrimHead(V843)

tmp4612 := PrimIsPair(tmp4611)

var ifres4595 Obj

if True == tmp4612 {
tmp4608 := PrimHead(V843)

tmp4609 := PrimTail(tmp4608)

tmp4610 := PrimIsPair(tmp4609)

var ifres4597 Obj

if True == tmp4610 {
tmp4604 := PrimHead(V843)

tmp4605 := PrimTail(tmp4604)

tmp4606 := PrimTail(tmp4605)

tmp4607 := PrimEqual(Nil, tmp4606)

var ifres4599 Obj

if True == tmp4607 {
tmp4601 := PrimHead(V843)

tmp4602 := PrimHead(tmp4601)

tmp4603 := PrimEqual(V842, tmp4602)

var ifres4600 Obj

if True == tmp4603 {
ifres4600 = True


} else {
ifres4600 = False


}

ifres4599 = ifres4600


} else {
ifres4599 = False


}

var ifres4598 Obj

if True == ifres4599 {
ifres4598 = True


} else {
ifres4598 = False


}

ifres4597 = ifres4598


} else {
ifres4597 = False


}

var ifres4596 Obj

if True == ifres4597 {
ifres4596 = True


} else {
ifres4596 = False


}

ifres4595 = ifres4596


} else {
ifres4595 = False


}

var ifres4594 Obj

if True == ifres4595 {
ifres4594 = True


} else {
ifres4594 = False


}

ifres4593 = ifres4594


} else {
ifres4593 = False


}

if True == ifres4593 {
tmp4585 := PrimHead(V843)

tmp4586 := PrimHead(tmp4585)

tmp4587 := PrimTail(V843)

tmp4588 := PrimHead(V843)

tmp4589 := PrimTail(tmp4588)

tmp4590 := PrimCons(True, tmp4589)

tmp4591 := PrimCons(tmp4590, V844)

__e.TailApply(PrimFunc(symshen_4pivot_1on), tmp4586, tmp4587, tmp4591)
return


} else {
__e.TailApply(PrimFunc(sym_8p), V844, V843)
return
}


}


}, 3)

tmp4670 := Call(__e, ns2_1set, symshen_4pivot_1on, tmp4570)


_ = tmp4670

tmp4671 := MakeNative(func(__e *ControlFlow) {
V847 := __e.Get(1)
_ = V847
V848 := __e.Get(2)
_ = V848
tmp4695 := PrimIsPair(V847)

var ifres4686 Obj

if True == tmp4695 {
tmp4693 := PrimTail(V847)

tmp4694 := PrimIsPair(tmp4693)

var ifres4688 Obj

if True == tmp4694 {
tmp4690 := PrimTail(V847)

tmp4691 := PrimTail(tmp4690)

tmp4692 := PrimEqual(Nil, tmp4691)

var ifres4689 Obj

if True == tmp4692 {
ifres4689 = True


} else {
ifres4689 = False


}

ifres4688 = ifres4689


} else {
ifres4688 = False


}

var ifres4687 Obj

if True == ifres4688 {
ifres4687 = True


} else {
ifres4687 = False


}

ifres4686 = ifres4687


} else {
ifres4686 = False


}

if True == ifres4686 {
tmp4672 := MakeNative(func(__e *ControlFlow) {
W849 := __e.Get(1)
_ = W849
tmp4682 := PrimEqual(symshen_4skip, W849)

if True == tmp4682 {
__e.Return(V848)
return
} else {
tmp4673 := Call(__e, PrimFunc(symshen_4op1), W849)


tmp4674 := PrimTail(V847)

tmp4675 := PrimCons(tmp4673, tmp4674)

tmp4676 := Call(__e, PrimFunc(symshen_4op2), W849)


tmp4677 := PrimTail(V847)

tmp4678 := PrimCons(tmp4676, tmp4677)

tmp4679 := PrimCons(tmp4678, Nil)

tmp4680 := PrimCons(tmp4675, tmp4679)

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4680, V848)
return


}


}, 1)

tmp4683 := PrimHead(V847)

tmp4684 := Call(__e, PrimFunc(symshen_4op), tmp4683)


__e.TailApply(tmp4672, tmp4684)
return


} else {
__e.Return(V848)
return
}


}, 2)

tmp4696 := Call(__e, ns2_1set, symshen_4factor_1selectors, tmp4671)


_ = tmp4696

tmp4697 := MakeNative(func(__e *ControlFlow) {
V852 := __e.Get(1)
_ = V852
tmp4705 := PrimEqual(symcons_2, V852)

if True == tmp4705 {
__e.Return(symcons)
return
} else {
tmp4703 := PrimEqual(symshen_4_7string_2, V852)

if True == tmp4703 {
__e.Return(sym_8s)
return
} else {
tmp4701 := PrimEqual(symshen_4_7vector_2, V852)

if True == tmp4701 {
__e.Return(sym_8v)
return
} else {
tmp4699 := PrimEqual(symtuple_2, V852)

if True == tmp4699 {
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

tmp4706 := Call(__e, ns2_1set, symshen_4op, tmp4697)


_ = tmp4706

tmp4707 := MakeNative(func(__e *ControlFlow) {
V853 := __e.Get(1)
_ = V853
V854 := __e.Get(2)
_ = V854
tmp4726 := PrimEqual(Nil, V853)

if True == tmp4726 {
__e.Return(V854)
return
} else {
tmp4724 := PrimIsPair(V853)

if True == tmp4724 {
tmp4720 := PrimHead(V853)

tmp4721 := Call(__e, PrimFunc(symoccurrences), tmp4720, V854)


tmp4722 := PrimGreatThan(tmp4721, MakeNumber(1))

if True == tmp4722 {
tmp4708 := MakeNative(func(__e *ControlFlow) {
W855 := __e.Get(1)
_ = W855
tmp4709 := PrimHead(V853)

tmp4710 := PrimTail(V853)

tmp4711 := PrimHead(V853)

tmp4712 := Call(__e, PrimFunc(symsubst), W855, tmp4711, V854)


tmp4713 := Call(__e, PrimFunc(symshen_4factor_1selectors_1h), tmp4710, tmp4712)


tmp4714 := PrimCons(tmp4713, Nil)

tmp4715 := PrimCons(tmp4709, tmp4714)

tmp4716 := PrimCons(W855, tmp4715)

__e.Return(PrimCons(symlet, tmp4716))
return


}, 1)

tmp4717 := Call(__e, PrimFunc(symgensym), symSelect)


__e.TailApply(tmp4708, tmp4717)
return


} else {
tmp4718 := PrimTail(V853)

__e.TailApply(PrimFunc(symshen_4factor_1selectors_1h), tmp4718, V854)
return


}


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.factor-selectors-h")))
return
}


}


}, 2)

__e.TailApply(ns2_1set, symshen_4factor_1selectors_1h, tmp4707)
return




}, 0)

