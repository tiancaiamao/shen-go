package main

import . "github.com/tiancaiamao/shen-go/kl"

var DeclarationsMain = MakeNative(func(__e *ControlFlow) {
tmp7573 := PrimSet(symshen_4_dhistory_d, Nil)

_ = tmp7573

tmp7574 := PrimSet(symshen_4_dtc_d, False)

_ = tmp7574

tmp7575 := Call(__e, PrimFunc(symvector), MakeNumber(20000))


tmp7576 := PrimSet(sym_dproperty_1vector_d, tmp7575)

_ = tmp7576

tmp7577 := MakeNative(func(__e *ControlFlow) {
X := __e.Get(1)
_ = X
__e.TailApply(PrimFunc(symshen_4macros), X)
return
}, 1)

tmp7578 := PrimCons(symshen_4macros, tmp7577)

tmp7579 := PrimCons(tmp7578, Nil)

tmp7580 := PrimSet(sym_dmacros_d, tmp7579)

_ = tmp7580

tmp7581 := PrimSet(symshen_4_dgensym_d, MakeNumber(0))

_ = tmp7581

tmp7582 := PrimSet(symshen_4_dtracking_d, Nil)

_ = tmp7582

tmp7583 := PrimSet(symshen_4_dprofiled_d, Nil)

_ = tmp7583

tmp7584 := PrimSet(sym_dhome_1directory_d, MakeString(""))

_ = tmp7584

tmp7585 := PrimCons(symtype, Nil)

tmp7586 := PrimCons(syminput_7, tmp7585)

tmp7587 := PrimCons(symopen, tmp7586)

tmp7588 := PrimCons(symset, tmp7587)

tmp7589 := PrimCons(symwhere, tmp7588)

tmp7590 := PrimCons(symlet, tmp7589)

tmp7591 := PrimCons(symlambda, tmp7590)

tmp7592 := PrimCons(symcons, tmp7591)

tmp7593 := PrimCons(sym_8v, tmp7592)

tmp7594 := PrimCons(sym_8s, tmp7593)

tmp7595 := PrimCons(sym_8p, tmp7594)

tmp7596 := PrimSet(symshen_4_dspecial_d, tmp7595)

_ = tmp7596

tmp7597 := PrimSet(symshen_4_dextraspecial_d, Nil)

_ = tmp7597

tmp7598 := PrimSet(symshen_4_dspy_d, False)

_ = tmp7598

tmp7599 := PrimSet(symshen_4_ddatatypes_d, Nil)

_ = tmp7599

tmp7600 := PrimSet(symshen_4_dalldatatypes_d, Nil)

_ = tmp7600

tmp7601 := PrimSet(symshen_4_dshen_1type_1theory_1enabled_2_d, True)

_ = tmp7601

tmp7602 := PrimSet(symshen_4_dpackage_d, symnull)

_ = tmp7602

tmp7603 := PrimSet(symshen_4_dsynonyms_d, Nil)

_ = tmp7603

tmp7604 := PrimSet(symshen_4_dsystem_d, Nil)

_ = tmp7604

tmp7605 := PrimSet(symshen_4_dsigf_d, Nil)

_ = tmp7605

tmp7606 := PrimSet(symshen_4_doccurs_d, True)

_ = tmp7606

tmp7607 := PrimSet(symshen_4_dfactorise_2_d, False)

_ = tmp7607

tmp7608 := PrimSet(symshen_4_dmaxinferences_d, MakeNumber(1000000))

_ = tmp7608

tmp7609 := PrimSet(sym_dmaximum_1print_1sequence_1size_d, MakeNumber(20))

_ = tmp7609

tmp7610 := PrimSet(symshen_4_dcall_d, MakeNumber(0))

_ = tmp7610

tmp7611 := PrimSet(symshen_4_dinfs_d, MakeNumber(0))

_ = tmp7611

tmp7612 := PrimSet(sym_dhush_d, False)

_ = tmp7612

tmp7613 := PrimSet(symshen_4_doptimise_d, False)

_ = tmp7613

tmp7614 := PrimSet(sym_dversion_d, MakeString("39.2"))

_ = tmp7614

tmp7615 := PrimSet(symshen_4_dnames_d, Nil)

_ = tmp7615

tmp7616 := PrimSet(symshen_4_dstep_d, False)

_ = tmp7616

tmp7617 := PrimSet(symshen_4_dit_d, MakeString(""))

_ = tmp7617

tmp7618 := PrimSet(symshen_4_dresidue_d, Nil)

_ = tmp7618

tmp7619 := PrimSet(symshen_4_dprolog_1memory_d, MakeNumber(1000))

_ = tmp7619

tmp7620 := PrimSet(symshen_4_dloading_2_d, False)

_ = tmp7620

tmp7621 := PrimSet(symshen_4_duserdefs_d, Nil)

_ = tmp7621

tmp7622 := MakeNative(func(__e *ControlFlow) {
tmp7623 := MakeNative(func(__e *ControlFlow) {
Z5776 := __e.Get(1)
_ = Z5776
__e.TailApply(PrimFunc(symshen_4typename), Z5776)
return
}, 1)

tmp7624 := PrimValue(symshen_4_dalldatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp7623, tmp7624)
return


}, 0)

tmp7625 := Call(__e, ns2_1set, symdatatypes, tmp7622)


_ = tmp7625

tmp7626 := MakeNative(func(__e *ControlFlow) {
tmp7627 := MakeNative(func(__e *ControlFlow) {
Z5777 := __e.Get(1)
_ = Z5777
__e.TailApply(PrimFunc(symshen_4typename), Z5777)
return
}, 1)

tmp7628 := PrimValue(symshen_4_ddatatypes_d)

__e.TailApply(PrimFunc(symmap), tmp7627, tmp7628)
return


}, 0)

tmp7629 := Call(__e, ns2_1set, symincluded, tmp7626)


_ = tmp7629

tmp7630 := MakeNative(func(__e *ControlFlow) {
V5780 := __e.Get(1)
_ = V5780
tmp7635 := PrimIsPair(V5780)

if True == tmp7635 {
tmp7631 := PrimHead(V5780)

tmp7632 := PrimStr(tmp7631)

tmp7633 := Call(__e, PrimFunc(symshen_4typename_1h), tmp7632)


__e.Return(PrimIntern(tmp7633))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.typename")))
return
}


}, 1)

tmp7636 := Call(__e, ns2_1set, symshen_4typename, tmp7630)


_ = tmp7636

tmp7637 := MakeNative(func(__e *ControlFlow) {
V5781 := __e.Get(1)
_ = V5781
tmp7644 := PrimEqual(MakeString("#type"), V5781)

if True == tmp7644 {
__e.Return(MakeString(""))
return
} else {
tmp7642 := Call(__e, PrimFunc(symshen_4_7string_2), V5781)


if True == tmp7642 {
tmp7638 := Call(__e, PrimFunc(symhdstr), V5781)


tmp7639 := PrimTailString(V5781)

tmp7640 := Call(__e, PrimFunc(symshen_4typename_1h), tmp7639)


__e.Return(PrimStringConcat(tmp7638, tmp7640))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.typename-h")))
return
}


}


}, 1)

tmp7645 := Call(__e, ns2_1set, symshen_4typename_1h, tmp7637)


_ = tmp7645

tmp7646 := MakeNative(func(__e *ControlFlow) {
V5782 := __e.Get(1)
_ = V5782
tmp7650 := PrimLessThan(V5782, MakeNumber(0))

if True == tmp7650 {
__e.Return(PrimValue(symshen_4_dprolog_1memory_d))
return
} else {
tmp7648 := PrimIsInteger(V5782)

if True == tmp7648 {
__e.Return(PrimSet(symshen_4_dprolog_1memory_d, V5782))
return
} else {
__e.Return(PrimSimpleError(MakeString("prolog memory expects an integer value\n")))
return
}


}


}, 1)

tmp7651 := Call(__e, ns2_1set, symprolog_1memory, tmp7646)


_ = tmp7651

tmp7652 := MakeNative(func(__e *ControlFlow) {
V5785 := __e.Get(1)
_ = V5785
tmp7668 := PrimEqual(Nil, V5785)

if True == tmp7668 {
__e.Return(Nil)
return
} else {
tmp7666 := PrimIsPair(V5785)

var ifres7662 Obj

if True == tmp7666 {
tmp7664 := PrimTail(V5785)

tmp7665 := PrimIsPair(tmp7664)

var ifres7663 Obj

if True == tmp7665 {
ifres7663 = True


} else {
ifres7663 = False


}

ifres7662 = ifres7663


} else {
ifres7662 = False


}

if True == ifres7662 {
tmp7653 := MakeNative(func(__e *ControlFlow) {
W5786 := __e.Get(1)
_ = W5786
tmp7654 := PrimTail(V5785)

tmp7655 := PrimTail(tmp7654)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7655)
return


}, 1)

tmp7656 := PrimHead(V5785)

tmp7657 := PrimTail(V5785)

tmp7658 := PrimHead(tmp7657)

tmp7659 := PrimValue(sym_dproperty_1vector_d)

tmp7660 := Call(__e, PrimFunc(symput), tmp7656, symarity, tmp7658, tmp7659)


__e.TailApply(tmp7653, tmp7660)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp7669 := Call(__e, ns2_1set, symshen_4initialise_1lambda_1tables, tmp7652)


_ = tmp7669

tmp7670 := MakeNative(func(__e *ControlFlow) {
V5787 := __e.Get(1)
_ = V5787
tmp7671 := MakeNative(func(__e *ControlFlow) {
tmp7672 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V5787, symarity, tmp7672)
return


}, 0)

tmp7673 := MakeNative(func(__e *ControlFlow) {
Z5788 := __e.Get(1)
_ = Z5788
__e.Return(MakeNumber(-1))
return
}, 1)

__e.TailApply(try_1catch, tmp7671, tmp7673)
return


}, 1)

tmp7674 := Call(__e, ns2_1set, symarity, tmp7670)


_ = tmp7674

tmp7675 := MakeNative(func(__e *ControlFlow) {
V5791 := __e.Get(1)
_ = V5791
tmp7691 := PrimEqual(Nil, V5791)

if True == tmp7691 {
__e.Return(Nil)
return
} else {
tmp7689 := PrimIsPair(V5791)

var ifres7685 Obj

if True == tmp7689 {
tmp7687 := PrimTail(V5791)

tmp7688 := PrimIsPair(tmp7687)

var ifres7686 Obj

if True == tmp7688 {
ifres7686 = True


} else {
ifres7686 = False


}

ifres7685 = ifres7686


} else {
ifres7685 = False


}

if True == ifres7685 {
tmp7676 := MakeNative(func(__e *ControlFlow) {
W5792 := __e.Get(1)
_ = W5792
tmp7677 := PrimTail(V5791)

tmp7678 := PrimTail(tmp7677)

__e.TailApply(PrimFunc(symshen_4initialise_1arity_1table), tmp7678)
return


}, 1)

tmp7679 := PrimHead(V5791)

tmp7680 := PrimTail(V5791)

tmp7681 := PrimHead(tmp7680)

tmp7682 := PrimValue(sym_dproperty_1vector_d)

tmp7683 := Call(__e, PrimFunc(symput), tmp7679, symarity, tmp7681, tmp7682)


__e.TailApply(tmp7676, tmp7683)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.initialise-arity-table")))
return
}


}


}, 1)

tmp7692 := Call(__e, ns2_1set, symshen_4initialise_1arity_1table, tmp7675)


_ = tmp7692

tmp7693 := PrimCons(MakeNumber(2), Nil)

tmp7694 := PrimCons(sym_8s, tmp7693)

tmp7695 := PrimCons(MakeNumber(2), tmp7694)

tmp7696 := PrimCons(sym_8v, tmp7695)

tmp7697 := PrimCons(MakeNumber(2), tmp7696)

tmp7698 := PrimCons(sym_8p, tmp7697)

tmp7699 := PrimCons(MakeNumber(1), tmp7698)

tmp7700 := PrimCons(sym_5_b_6, tmp7699)

tmp7701 := PrimCons(MakeNumber(1), tmp7700)

tmp7702 := PrimCons(sym_5end_6, tmp7701)

tmp7703 := PrimCons(MakeNumber(1), tmp7702)

tmp7704 := PrimCons(sym_5e_6, tmp7703)

tmp7705 := PrimCons(MakeNumber(2), tmp7704)

tmp7706 := PrimCons(sym_a_a, tmp7705)

tmp7707 := PrimCons(MakeNumber(2), tmp7706)

tmp7708 := PrimCons(sym_1, tmp7707)

tmp7709 := PrimCons(MakeNumber(2), tmp7708)

tmp7710 := PrimCons(sym_c, tmp7709)

tmp7711 := PrimCons(MakeNumber(2), tmp7710)

tmp7712 := PrimCons(sym_d, tmp7711)

tmp7713 := PrimCons(MakeNumber(2), tmp7712)

tmp7714 := PrimCons(sym_7, tmp7713)

tmp7715 := PrimCons(MakeNumber(1), tmp7714)

tmp7716 := PrimCons(symy_1or_1n_2, tmp7715)

tmp7717 := PrimCons(MakeNumber(2), tmp7716)

tmp7718 := PrimCons(symwrite_1to_1file, tmp7717)

tmp7719 := PrimCons(MakeNumber(2), tmp7718)

tmp7720 := PrimCons(symwrite_1byte, tmp7719)

tmp7721 := PrimCons(MakeNumber(5), tmp7720)

tmp7722 := PrimCons(symwhen, tmp7721)

tmp7723 := PrimCons(MakeNumber(0), tmp7722)

tmp7724 := PrimCons(symversion, tmp7723)

tmp7725 := PrimCons(MakeNumber(5), tmp7724)

tmp7726 := PrimCons(symvar_2, tmp7725)

tmp7727 := PrimCons(MakeNumber(1), tmp7726)

tmp7728 := PrimCons(symvariable_2, tmp7727)

tmp7729 := PrimCons(MakeNumber(1), tmp7728)

tmp7730 := PrimCons(symvalue, tmp7729)

tmp7731 := PrimCons(MakeNumber(3), tmp7730)

tmp7732 := PrimCons(symvector_1_6, tmp7731)

tmp7733 := PrimCons(MakeNumber(1), tmp7732)

tmp7734 := PrimCons(symvector_2, tmp7733)

tmp7735 := PrimCons(MakeNumber(1), tmp7734)

tmp7736 := PrimCons(symvector, tmp7735)

tmp7737 := PrimCons(MakeNumber(0), tmp7736)

tmp7738 := PrimCons(symuserdefs, tmp7737)

tmp7739 := PrimCons(MakeNumber(2), tmp7738)

tmp7740 := PrimCons(symupdate_1lambda_1table, tmp7739)

tmp7741 := PrimCons(MakeNumber(1), tmp7740)

tmp7742 := PrimCons(symundefmacro, tmp7741)

tmp7743 := PrimCons(MakeNumber(1), tmp7742)

tmp7744 := PrimCons(symuntrack, tmp7743)

tmp7745 := PrimCons(MakeNumber(2), tmp7744)

tmp7746 := PrimCons(symunion, tmp7745)

tmp7747 := PrimCons(MakeNumber(1), tmp7746)

tmp7748 := PrimCons(symunprofile, tmp7747)

tmp7749 := PrimCons(MakeNumber(3), tmp7748)

tmp7750 := PrimCons(symunput, tmp7749)

tmp7751 := PrimCons(MakeNumber(1), tmp7750)

tmp7752 := PrimCons(symundefmacro, tmp7751)

tmp7753 := PrimCons(MakeNumber(5), tmp7752)

tmp7754 := PrimCons(symreturn, tmp7753)

tmp7755 := PrimCons(MakeNumber(2), tmp7754)

tmp7756 := PrimCons(symtype, tmp7755)

tmp7757 := PrimCons(MakeNumber(1), tmp7756)

tmp7758 := PrimCons(symtuple_2, tmp7757)

tmp7759 := PrimCons(MakeNumber(2), tmp7758)

tmp7760 := PrimCons(symtrap_1error, tmp7759)

tmp7761 := PrimCons(MakeNumber(0), tmp7760)

tmp7762 := PrimCons(symtracked, tmp7761)

tmp7763 := PrimCons(MakeNumber(1), tmp7762)

tmp7764 := PrimCons(symtrack, tmp7763)

tmp7765 := PrimCons(MakeNumber(1), tmp7764)

tmp7766 := PrimCons(symtlstr, tmp7765)

tmp7767 := PrimCons(MakeNumber(1), tmp7766)

tmp7768 := PrimCons(symthaw, tmp7767)

tmp7769 := PrimCons(MakeNumber(0), tmp7768)

tmp7770 := PrimCons(symtc_2, tmp7769)

tmp7771 := PrimCons(MakeNumber(1), tmp7770)

tmp7772 := PrimCons(symtc, tmp7771)

tmp7773 := PrimCons(MakeNumber(1), tmp7772)

tmp7774 := PrimCons(symtl, tmp7773)

tmp7775 := PrimCons(MakeNumber(1), tmp7774)

tmp7776 := PrimCons(symtail, tmp7775)

tmp7777 := PrimCons(MakeNumber(1), tmp7776)

tmp7778 := PrimCons(symsystemf, tmp7777)

tmp7779 := PrimCons(MakeNumber(1), tmp7778)

tmp7780 := PrimCons(symsymbol_2, tmp7779)

tmp7781 := PrimCons(MakeNumber(1), tmp7780)

tmp7782 := PrimCons(symsum, tmp7781)

tmp7783 := PrimCons(MakeNumber(3), tmp7782)

tmp7784 := PrimCons(symsubst, tmp7783)

tmp7785 := PrimCons(MakeNumber(1), tmp7784)

tmp7786 := PrimCons(symstring_2, tmp7785)

tmp7787 := PrimCons(MakeNumber(1), tmp7786)

tmp7788 := PrimCons(symstring_1_6symbol, tmp7787)

tmp7789 := PrimCons(MakeNumber(1), tmp7788)

tmp7790 := PrimCons(symstring_1_6n, tmp7789)

tmp7791 := PrimCons(MakeNumber(1), tmp7790)

tmp7792 := PrimCons(symstr, tmp7791)

tmp7793 := PrimCons(MakeNumber(0), tmp7792)

tmp7794 := PrimCons(symstoutput, tmp7793)

tmp7795 := PrimCons(MakeNumber(0), tmp7794)

tmp7796 := PrimCons(symstinput, tmp7795)

tmp7797 := PrimCons(MakeNumber(0), tmp7796)

tmp7798 := PrimCons(symstep_2, tmp7797)

tmp7799 := PrimCons(MakeNumber(1), tmp7798)

tmp7800 := PrimCons(symstep, tmp7799)

tmp7801 := PrimCons(MakeNumber(0), tmp7800)

tmp7802 := PrimCons(symspy_2, tmp7801)

tmp7803 := PrimCons(MakeNumber(1), tmp7802)

tmp7804 := PrimCons(symspy, tmp7803)

tmp7805 := PrimCons(MakeNumber(2), tmp7804)

tmp7806 := PrimCons(symspecialise, tmp7805)

tmp7807 := PrimCons(MakeNumber(1), tmp7806)

tmp7808 := PrimCons(symsnd, tmp7807)

tmp7809 := PrimCons(MakeNumber(1), tmp7808)

tmp7810 := PrimCons(symsimple_1error, tmp7809)

tmp7811 := PrimCons(MakeNumber(2), tmp7810)

tmp7812 := PrimCons(symset, tmp7811)

tmp7813 := PrimCons(MakeNumber(1), tmp7812)

tmp7814 := PrimCons(symreverse, tmp7813)

tmp7815 := PrimCons(MakeNumber(2), tmp7814)

tmp7816 := PrimCons(symremove, tmp7815)

tmp7817 := PrimCons(MakeNumber(0), tmp7816)

tmp7818 := PrimCons(symrelease, tmp7817)

tmp7819 := PrimCons(MakeNumber(1), tmp7818)

tmp7820 := PrimCons(symreceive, tmp7819)

tmp7821 := PrimCons(MakeNumber(1), tmp7820)

tmp7822 := PrimCons(symshen_4read_1unit_1string, tmp7821)

tmp7823 := PrimCons(MakeNumber(1), tmp7822)

tmp7824 := PrimCons(symread_1from_1string_1unprocessed, tmp7823)

tmp7825 := PrimCons(MakeNumber(1), tmp7824)

tmp7826 := PrimCons(symread_1from_1string, tmp7825)

tmp7827 := PrimCons(MakeNumber(1), tmp7826)

tmp7828 := PrimCons(symread_1byte, tmp7827)

tmp7829 := PrimCons(MakeNumber(1), tmp7828)

tmp7830 := PrimCons(symread, tmp7829)

tmp7831 := PrimCons(MakeNumber(1), tmp7830)

tmp7832 := PrimCons(symread_1file, tmp7831)

tmp7833 := PrimCons(MakeNumber(1), tmp7832)

tmp7834 := PrimCons(symread_1file_1as_1bytelist, tmp7833)

tmp7835 := PrimCons(MakeNumber(1), tmp7834)

tmp7836 := PrimCons(symread_1file_1as_1string, tmp7835)

tmp7837 := PrimCons(MakeNumber(4), tmp7836)

tmp7838 := PrimCons(symput, tmp7837)

tmp7839 := PrimCons(MakeNumber(1), tmp7838)

tmp7840 := PrimCons(symprotect, tmp7839)

tmp7841 := PrimCons(MakeNumber(1), tmp7840)

tmp7842 := PrimCons(sympreclude_1all_1but, tmp7841)

tmp7843 := PrimCons(MakeNumber(1), tmp7842)

tmp7844 := PrimCons(sympreclude, tmp7843)

tmp7845 := PrimCons(MakeNumber(1), tmp7844)

tmp7846 := PrimCons(symps, tmp7845)

tmp7847 := PrimCons(MakeNumber(2), tmp7846)

tmp7848 := PrimCons(sympr, tmp7847)

tmp7849 := PrimCons(MakeNumber(1), tmp7848)

tmp7850 := PrimCons(symprofile_1results, tmp7849)

tmp7851 := PrimCons(MakeNumber(1), tmp7850)

tmp7852 := PrimCons(symprolog_1memory, tmp7851)

tmp7853 := PrimCons(MakeNumber(1), tmp7852)

tmp7854 := PrimCons(symshen_4printF, tmp7853)

tmp7855 := PrimCons(MakeNumber(1), tmp7854)

tmp7856 := PrimCons(symshen_4print_1freshterm, tmp7855)

tmp7857 := PrimCons(MakeNumber(1), tmp7856)

tmp7858 := PrimCons(symshen_4print_1prolog_1vector, tmp7857)

tmp7859 := PrimCons(MakeNumber(1), tmp7858)

tmp7860 := PrimCons(symprofile, tmp7859)

tmp7861 := PrimCons(MakeNumber(1), tmp7860)

tmp7862 := PrimCons(symprint, tmp7861)

tmp7863 := PrimCons(MakeNumber(1), tmp7862)

tmp7864 := PrimCons(sympreclude_1all_1but, tmp7863)

tmp7865 := PrimCons(MakeNumber(2), tmp7864)

tmp7866 := PrimCons(sympos, tmp7865)

tmp7867 := PrimCons(MakeNumber(0), tmp7866)

tmp7868 := PrimCons(symporters, tmp7867)

tmp7869 := PrimCons(MakeNumber(0), tmp7868)

tmp7870 := PrimCons(symport, tmp7869)

tmp7871 := PrimCons(MakeNumber(1), tmp7870)

tmp7872 := PrimCons(sympackage_2, tmp7871)

tmp7873 := PrimCons(MakeNumber(3), tmp7872)

tmp7874 := PrimCons(sympackage, tmp7873)

tmp7875 := PrimCons(MakeNumber(0), tmp7874)

tmp7876 := PrimCons(symos, tmp7875)

tmp7877 := PrimCons(MakeNumber(2), tmp7876)

tmp7878 := PrimCons(symor, tmp7877)

tmp7879 := PrimCons(MakeNumber(0), tmp7878)

tmp7880 := PrimCons(symoptimise_2, tmp7879)

tmp7881 := PrimCons(MakeNumber(1), tmp7880)

tmp7882 := PrimCons(symoptimise, tmp7881)

tmp7883 := PrimCons(MakeNumber(2), tmp7882)

tmp7884 := PrimCons(symopen, tmp7883)

tmp7885 := PrimCons(MakeNumber(1), tmp7884)

tmp7886 := PrimCons(symoccurs_1check, tmp7885)

tmp7887 := PrimCons(MakeNumber(0), tmp7886)

tmp7888 := PrimCons(symoccurs_2, tmp7887)

tmp7889 := PrimCons(MakeNumber(2), tmp7888)

tmp7890 := PrimCons(symoccurrences, tmp7889)

tmp7891 := PrimCons(MakeNumber(1), tmp7890)

tmp7892 := PrimCons(symoccurs_1check, tmp7891)

tmp7893 := PrimCons(MakeNumber(1), tmp7892)

tmp7894 := PrimCons(symnumber_2, tmp7893)

tmp7895 := PrimCons(MakeNumber(1), tmp7894)

tmp7896 := PrimCons(symn_1_6string, tmp7895)

tmp7897 := PrimCons(MakeNumber(2), tmp7896)

tmp7898 := PrimCons(symnth, tmp7897)

tmp7899 := PrimCons(MakeNumber(1), tmp7898)

tmp7900 := PrimCons(symnot, tmp7899)

tmp7901 := PrimCons(MakeNumber(1), tmp7900)

tmp7902 := PrimCons(symnl, tmp7901)

tmp7903 := PrimCons(MakeNumber(1), tmp7902)

tmp7904 := PrimCons(symmaxinferences, tmp7903)

tmp7905 := PrimCons(MakeNumber(2), tmp7904)

tmp7906 := PrimCons(symmapcan, tmp7905)

tmp7907 := PrimCons(MakeNumber(2), tmp7906)

tmp7908 := PrimCons(symmap, tmp7907)

tmp7909 := PrimCons(MakeNumber(1), tmp7908)

tmp7910 := PrimCons(symmacroexpand, tmp7909)

tmp7911 := PrimCons(MakeNumber(1), tmp7910)

tmp7912 := PrimCons(symvector, tmp7911)

tmp7913 := PrimCons(MakeNumber(2), tmp7912)

tmp7914 := PrimCons(sym_5_a, tmp7913)

tmp7915 := PrimCons(MakeNumber(2), tmp7914)

tmp7916 := PrimCons(sym_5, tmp7915)

tmp7917 := PrimCons(MakeNumber(1), tmp7916)

tmp7918 := PrimCons(symload, tmp7917)

tmp7919 := PrimCons(MakeNumber(1), tmp7918)

tmp7920 := PrimCons(symlist, tmp7919)

tmp7921 := PrimCons(MakeNumber(1), tmp7920)

tmp7922 := PrimCons(symlineread, tmp7921)

tmp7923 := PrimCons(MakeNumber(1), tmp7922)

tmp7924 := PrimCons(symlimit, tmp7923)

tmp7925 := PrimCons(MakeNumber(1), tmp7924)

tmp7926 := PrimCons(symlength, tmp7925)

tmp7927 := PrimCons(MakeNumber(0), tmp7926)

tmp7928 := PrimCons(symlanguage, tmp7927)

tmp7929 := PrimCons(MakeNumber(6), tmp7928)

tmp7930 := PrimCons(symis_b, tmp7929)

tmp7931 := PrimCons(MakeNumber(6), tmp7930)

tmp7932 := PrimCons(symis, tmp7931)

tmp7933 := PrimCons(MakeNumber(0), tmp7932)

tmp7934 := PrimCons(symit, tmp7933)

tmp7935 := PrimCons(MakeNumber(1), tmp7934)

tmp7936 := PrimCons(syminternal, tmp7935)

tmp7937 := PrimCons(MakeNumber(2), tmp7936)

tmp7938 := PrimCons(symintersection, tmp7937)

tmp7939 := PrimCons(MakeNumber(1), tmp7938)

tmp7940 := PrimCons(syminclude_1all_1but, tmp7939)

tmp7941 := PrimCons(MakeNumber(0), tmp7940)

tmp7942 := PrimCons(symimplementation, tmp7941)

tmp7943 := PrimCons(MakeNumber(2), tmp7942)

tmp7944 := PrimCons(syminput_7, tmp7943)

tmp7945 := PrimCons(MakeNumber(1), tmp7944)

tmp7946 := PrimCons(syminput, tmp7945)

tmp7947 := PrimCons(MakeNumber(0), tmp7946)

tmp7948 := PrimCons(syminferences, tmp7947)

tmp7949 := PrimCons(MakeNumber(1), tmp7948)

tmp7950 := PrimCons(symintern, tmp7949)

tmp7951 := PrimCons(MakeNumber(1), tmp7950)

tmp7952 := PrimCons(syminternal, tmp7951)

tmp7953 := PrimCons(MakeNumber(1), tmp7952)

tmp7954 := PrimCons(syminteger_2, tmp7953)

tmp7955 := PrimCons(MakeNumber(1), tmp7954)

tmp7956 := PrimCons(symin_1package, tmp7955)

tmp7957 := PrimCons(MakeNumber(0), tmp7956)

tmp7958 := PrimCons(symincluded, tmp7957)

tmp7959 := PrimCons(MakeNumber(1), tmp7958)

tmp7960 := PrimCons(syminclude, tmp7959)

tmp7961 := PrimCons(MakeNumber(3), tmp7960)

tmp7962 := PrimCons(symif, tmp7961)

tmp7963 := PrimCons(MakeNumber(1), tmp7962)

tmp7964 := PrimCons(symhush, tmp7963)

tmp7965 := PrimCons(MakeNumber(0), tmp7964)

tmp7966 := PrimCons(symhush_2, tmp7965)

tmp7967 := PrimCons(MakeNumber(1), tmp7966)

tmp7968 := PrimCons(symhead, tmp7967)

tmp7969 := PrimCons(MakeNumber(1), tmp7968)

tmp7970 := PrimCons(symhdstr, tmp7969)

tmp7971 := PrimCons(MakeNumber(1), tmp7970)

tmp7972 := PrimCons(symhdv, tmp7971)

tmp7973 := PrimCons(MakeNumber(1), tmp7972)

tmp7974 := PrimCons(symhd, tmp7973)

tmp7975 := PrimCons(MakeNumber(2), tmp7974)

tmp7976 := PrimCons(symhash, tmp7975)

tmp7977 := PrimCons(MakeNumber(2), tmp7976)

tmp7978 := PrimCons(sym_a, tmp7977)

tmp7979 := PrimCons(MakeNumber(2), tmp7978)

tmp7980 := PrimCons(sym_6_a, tmp7979)

tmp7981 := PrimCons(MakeNumber(2), tmp7980)

tmp7982 := PrimCons(sym_6, tmp7981)

tmp7983 := PrimCons(MakeNumber(2), tmp7982)

tmp7984 := PrimCons(sym_5_1vector, tmp7983)

tmp7985 := PrimCons(MakeNumber(2), tmp7984)

tmp7986 := PrimCons(sym_5_1address, tmp7985)

tmp7987 := PrimCons(MakeNumber(3), tmp7986)

tmp7988 := PrimCons(symaddress_1_6, tmp7987)

tmp7989 := PrimCons(MakeNumber(1), tmp7988)

tmp7990 := PrimCons(symget_1time, tmp7989)

tmp7991 := PrimCons(MakeNumber(3), tmp7990)

tmp7992 := PrimCons(symget, tmp7991)

tmp7993 := PrimCons(MakeNumber(1), tmp7992)

tmp7994 := PrimCons(symgensym, tmp7993)

tmp7995 := PrimCons(MakeNumber(1), tmp7994)

tmp7996 := PrimCons(symfunction, tmp7995)

tmp7997 := PrimCons(MakeNumber(1), tmp7996)

tmp7998 := PrimCons(symfn, tmp7997)

tmp7999 := PrimCons(MakeNumber(1), tmp7998)

tmp8000 := PrimCons(symfst, tmp7999)

tmp8001 := PrimCons(MakeNumber(0), tmp8000)

tmp8002 := PrimCons(symfresh, tmp8001)

tmp8003 := PrimCons(MakeNumber(1), tmp8002)

tmp8004 := PrimCons(symfreeze, tmp8003)

tmp8005 := PrimCons(MakeNumber(5), tmp8004)

tmp8006 := PrimCons(symfork, tmp8005)

tmp8007 := PrimCons(MakeNumber(1), tmp8006)

tmp8008 := PrimCons(symforeign, tmp8007)

tmp8009 := PrimCons(MakeNumber(7), tmp8008)

tmp8010 := PrimCons(symfindall, tmp8009)

tmp8011 := PrimCons(MakeNumber(2), tmp8010)

tmp8012 := PrimCons(symfix, tmp8011)

tmp8013 := PrimCons(MakeNumber(0), tmp8012)

tmp8014 := PrimCons(symfail, tmp8013)

tmp8015 := PrimCons(MakeNumber(2), tmp8014)

tmp8016 := PrimCons(symfail_1if, tmp8015)

tmp8017 := PrimCons(MakeNumber(0), tmp8016)

tmp8018 := PrimCons(symfactorise_2, tmp8017)

tmp8019 := PrimCons(MakeNumber(1), tmp8018)

tmp8020 := PrimCons(symfactorise, tmp8019)

tmp8021 := PrimCons(MakeNumber(1), tmp8020)

tmp8022 := PrimCons(symexternal, tmp8021)

tmp8023 := PrimCons(MakeNumber(1), tmp8022)

tmp8024 := PrimCons(symexplode, tmp8023)

tmp8025 := PrimCons(MakeNumber(1), tmp8024)

tmp8026 := PrimCons(symeval_1kl, tmp8025)

tmp8027 := PrimCons(MakeNumber(1), tmp8026)

tmp8028 := PrimCons(symeval, tmp8027)

tmp8029 := PrimCons(MakeNumber(1), tmp8028)

tmp8030 := PrimCons(symerror_1to_1string, tmp8029)

tmp8031 := PrimCons(MakeNumber(1), tmp8030)

tmp8032 := PrimCons(symexternal, tmp8031)

tmp8033 := PrimCons(MakeNumber(1), tmp8032)

tmp8034 := PrimCons(symenable_1type_1theory, tmp8033)

tmp8035 := PrimCons(MakeNumber(1), tmp8034)

tmp8036 := PrimCons(symempty_2, tmp8035)

tmp8037 := PrimCons(MakeNumber(2), tmp8036)

tmp8038 := PrimCons(symelement_2, tmp8037)

tmp8039 := PrimCons(MakeNumber(2), tmp8038)

tmp8040 := PrimCons(symdo, tmp8039)

tmp8041 := PrimCons(MakeNumber(2), tmp8040)

tmp8042 := PrimCons(symdifference, tmp8041)

tmp8043 := PrimCons(MakeNumber(1), tmp8042)

tmp8044 := PrimCons(symdestroy, tmp8043)

tmp8045 := PrimCons(MakeNumber(2), tmp8044)

tmp8046 := PrimCons(symdeclare, tmp8045)

tmp8047 := PrimCons(MakeNumber(0), tmp8046)

tmp8048 := PrimCons(symdatatypes, tmp8047)

tmp8049 := PrimCons(MakeNumber(1), tmp8048)

tmp8050 := PrimCons(symclose, tmp8049)

tmp8051 := PrimCons(MakeNumber(2), tmp8050)

tmp8052 := PrimCons(symcn, tmp8051)

tmp8053 := PrimCons(MakeNumber(1), tmp8052)

tmp8054 := PrimCons(symcons_2, tmp8053)

tmp8055 := PrimCons(MakeNumber(2), tmp8054)

tmp8056 := PrimCons(symcons, tmp8055)

tmp8057 := PrimCons(MakeNumber(2), tmp8056)

tmp8058 := PrimCons(symconcat, tmp8057)

tmp8059 := PrimCons(MakeNumber(2), tmp8058)

tmp8060 := PrimCons(symcompile, tmp8059)

tmp8061 := PrimCons(MakeNumber(1), tmp8060)

tmp8062 := PrimCons(symcd, tmp8061)

tmp8063 := PrimCons(MakeNumber(5), tmp8062)

tmp8064 := PrimCons(symcall, tmp8063)

tmp8065 := PrimCons(MakeNumber(6), tmp8064)

tmp8066 := PrimCons(symbind, tmp8065)

tmp8067 := PrimCons(MakeNumber(1), tmp8066)

tmp8068 := PrimCons(symbound_2, tmp8067)

tmp8069 := PrimCons(MakeNumber(1), tmp8068)

tmp8070 := PrimCons(symbootstrap, tmp8069)

tmp8071 := PrimCons(MakeNumber(1), tmp8070)

tmp8072 := PrimCons(symboolean_2, tmp8071)

tmp8073 := PrimCons(MakeNumber(1), tmp8072)

tmp8074 := PrimCons(symatom_2, tmp8073)

tmp8075 := PrimCons(MakeNumber(2), tmp8074)

tmp8076 := PrimCons(symassoc, tmp8075)

tmp8077 := PrimCons(MakeNumber(1), tmp8076)

tmp8078 := PrimCons(symarity, tmp8077)

tmp8079 := PrimCons(MakeNumber(2), tmp8078)

tmp8080 := PrimCons(symappend, tmp8079)

tmp8081 := PrimCons(MakeNumber(2), tmp8080)

tmp8082 := PrimCons(symand, tmp8081)

tmp8083 := PrimCons(MakeNumber(2), tmp8082)

tmp8084 := PrimCons(symadjoin, tmp8083)

tmp8085 := PrimCons(MakeNumber(3), tmp8084)

tmp8086 := PrimCons(symaddress_1_6, tmp8085)

tmp8087 := PrimCons(MakeNumber(1), tmp8086)

tmp8088 := PrimCons(symabsvector, tmp8087)

tmp8089 := PrimCons(MakeNumber(1), tmp8088)

tmp8090 := PrimCons(symabsvector_2, tmp8089)

tmp8091 := PrimCons(MakeNumber(0), tmp8090)

tmp8092 := PrimCons(symabort, tmp8091)

tmp8093 := Call(__e, PrimFunc(symshen_4initialise_1arity_1table), tmp8092)


_ = tmp8093

tmp8094 := MakeNative(func(__e *ControlFlow) {
V5793 := __e.Get(1)
_ = V5793
tmp8095 := MakeNative(func(__e *ControlFlow) {
W5794 := __e.Get(1)
_ = W5794
tmp8096 := MakeNative(func(__e *ControlFlow) {
W5795 := __e.Get(1)
_ = W5795
__e.Return(V5793)
return
}, 1)

tmp8097 := Call(__e, PrimFunc(symadjoin), V5793, W5794)


tmp8098 := PrimValue(sym_dproperty_1vector_d)

tmp8099 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8097, tmp8098)


__e.TailApply(tmp8096, tmp8099)
return


}, 1)

tmp8100 := PrimValue(sym_dproperty_1vector_d)

tmp8101 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp8100)


__e.TailApply(tmp8095, tmp8101)
return


}, 1)

tmp8102 := Call(__e, ns2_1set, symsystemf, tmp8094)


_ = tmp8102

tmp8103 := MakeNative(func(__e *ControlFlow) {
V5796 := __e.Get(1)
_ = V5796
V5797 := __e.Get(2)
_ = V5797
tmp8105 := Call(__e, PrimFunc(symelement_2), V5796, V5797)


if True == tmp8105 {
__e.Return(V5797)
return
} else {
__e.Return(PrimCons(V5796, V5797))
return
}


}, 2)

tmp8106 := Call(__e, ns2_1set, symadjoin, tmp8103)


_ = tmp8106

tmp8107 := PrimIntern(MakeString(":"))

tmp8108 := PrimIntern(MakeString(";"))

tmp8109 := PrimIntern(MakeString(":="))

tmp8110 := PrimIntern(MakeString(","))

tmp8111 := Call(__e, PrimFunc(symvector), MakeNumber(0))


tmp8112 := PrimIntern(MakeString("bar!"))

tmp8113 := PrimCons(symabort, Nil)

tmp8114 := PrimCons(symabsvector, tmp8113)

tmp8115 := PrimCons(symabsvector_2, tmp8114)

tmp8116 := PrimCons(symaddress_1_6, tmp8115)

tmp8117 := PrimCons(sym_5_1address, tmp8116)

tmp8118 := PrimCons(symadjoin, tmp8117)

tmp8119 := PrimCons(symand, tmp8118)

tmp8120 := PrimCons(symappend, tmp8119)

tmp8121 := PrimCons(symarity, tmp8120)

tmp8122 := PrimCons(symassoc, tmp8121)

tmp8123 := PrimCons(symassertz, tmp8122)

tmp8124 := PrimCons(symasserta, tmp8123)

tmp8125 := PrimCons(symatom_2, tmp8124)

tmp8126 := PrimCons(symstep_2, tmp8125)

tmp8127 := PrimCons(symspy_2, tmp8126)

tmp8128 := PrimCons(tmp8112, tmp8127)

tmp8129 := PrimCons(symbootstrap, tmp8128)

tmp8130 := PrimCons(symboolean, tmp8129)

tmp8131 := PrimCons(symboolean_2, tmp8130)

tmp8132 := PrimCons(symbound_2, tmp8131)

tmp8133 := PrimCons(symbind, tmp8132)

tmp8134 := PrimCons(symclose, tmp8133)

tmp8135 := PrimCons(symcall, tmp8134)

tmp8136 := PrimCons(symcases, tmp8135)

tmp8137 := PrimCons(symcd, tmp8136)

tmp8138 := PrimCons(symcompile, tmp8137)

tmp8139 := PrimCons(symconcat, tmp8138)

tmp8140 := PrimCons(symcond, tmp8139)

tmp8141 := PrimCons(symcons, tmp8140)

tmp8142 := PrimCons(symcons_2, tmp8141)

tmp8143 := PrimCons(symcn, tmp8142)

tmp8144 := PrimCons(symdatatypes, tmp8143)

tmp8145 := PrimCons(symdatatype, tmp8144)

tmp8146 := PrimCons(symdeclare, tmp8145)

tmp8147 := PrimCons(symdefprolog, tmp8146)

tmp8148 := PrimCons(symdefcc, tmp8147)

tmp8149 := PrimCons(symdefmacro, tmp8148)

tmp8150 := PrimCons(symdefine, tmp8149)

tmp8151 := PrimCons(symdefun, tmp8150)

tmp8152 := PrimCons(symdestroy, tmp8151)

tmp8153 := PrimCons(symdifference, tmp8152)

tmp8154 := PrimCons(symdo, tmp8153)

tmp8155 := PrimCons(symelement_2, tmp8154)

tmp8156 := PrimCons(symempty_2, tmp8155)

tmp8157 := PrimCons(symerror, tmp8156)

tmp8158 := PrimCons(symerror_1to_1string, tmp8157)

tmp8159 := PrimCons(symeval, tmp8158)

tmp8160 := PrimCons(symeval_1kl, tmp8159)

tmp8161 := PrimCons(symexception, tmp8160)

tmp8162 := PrimCons(symexternal, tmp8161)

tmp8163 := PrimCons(symexplode, tmp8162)

tmp8164 := PrimCons(symenable_1type_1theory, tmp8163)

tmp8165 := PrimCons(False, tmp8164)

tmp8166 := PrimCons(symfindall, tmp8165)

tmp8167 := PrimCons(symfactorise_2, tmp8166)

tmp8168 := PrimCons(symfactorise, tmp8167)

tmp8169 := PrimCons(symfail_1if, tmp8168)

tmp8170 := PrimCons(symfail, tmp8169)

tmp8171 := PrimCons(symfile, tmp8170)

tmp8172 := PrimCons(symfix, tmp8171)

tmp8173 := PrimCons(symforeign, tmp8172)

tmp8174 := PrimCons(symfork, tmp8173)

tmp8175 := PrimCons(symfresh, tmp8174)

tmp8176 := PrimCons(symfreeze, tmp8175)

tmp8177 := PrimCons(symfst, tmp8176)

tmp8178 := PrimCons(symfunction, tmp8177)

tmp8179 := PrimCons(symfn, tmp8178)

tmp8180 := PrimCons(symgensym, tmp8179)

tmp8181 := PrimCons(symget_1time, tmp8180)

tmp8182 := PrimCons(symget, tmp8181)

tmp8183 := PrimCons(symhash, tmp8182)

tmp8184 := PrimCons(symhdstr, tmp8183)

tmp8185 := PrimCons(symhdv, tmp8184)

tmp8186 := PrimCons(symhd, tmp8185)

tmp8187 := PrimCons(symhead, tmp8186)

tmp8188 := PrimCons(symhush_2, tmp8187)

tmp8189 := PrimCons(symhush_2, tmp8188)

tmp8190 := PrimCons(symif, tmp8189)

tmp8191 := PrimCons(symimplementation, tmp8190)

tmp8192 := PrimCons(syminternal, tmp8191)

tmp8193 := PrimCons(symin_1package, tmp8192)

tmp8194 := PrimCons(symin, tmp8193)

tmp8195 := PrimCons(symis_b, tmp8194)

tmp8196 := PrimCons(symis, tmp8195)

tmp8197 := PrimCons(symit, tmp8196)

tmp8198 := PrimCons(syminclude_1all_1but, tmp8197)

tmp8199 := PrimCons(syminclude, tmp8198)

tmp8200 := PrimCons(symincluded, tmp8199)

tmp8201 := PrimCons(syminput_7, tmp8200)

tmp8202 := PrimCons(syminput, tmp8201)

tmp8203 := PrimCons(syminteger_2, tmp8202)

tmp8204 := PrimCons(symintern, tmp8203)

tmp8205 := PrimCons(syminferences, tmp8204)

tmp8206 := PrimCons(symintersection, tmp8205)

tmp8207 := PrimCons(symis, tmp8206)

tmp8208 := PrimCons(symlanguage, tmp8207)

tmp8209 := PrimCons(symlambda, tmp8208)

tmp8210 := PrimCons(symlazy, tmp8209)

tmp8211 := PrimCons(symlet, tmp8210)

tmp8212 := PrimCons(symlength, tmp8211)

tmp8213 := PrimCons(symlimit, tmp8212)

tmp8214 := PrimCons(symlineread, tmp8213)

tmp8215 := PrimCons(symlist, tmp8214)

tmp8216 := PrimCons(symloaded, tmp8215)

tmp8217 := PrimCons(symload, tmp8216)

tmp8218 := PrimCons(symmake_1string, tmp8217)

tmp8219 := PrimCons(symmap, tmp8218)

tmp8220 := PrimCons(symmapcan, tmp8219)

tmp8221 := PrimCons(symmaxinferences, tmp8220)

tmp8222 := PrimCons(symmacroexpand, tmp8221)

tmp8223 := PrimCons(symmode, tmp8222)

tmp8224 := PrimCons(symnl, tmp8223)

tmp8225 := PrimCons(symnot, tmp8224)

tmp8226 := PrimCons(symnth, tmp8225)

tmp8227 := PrimCons(symnull, tmp8226)

tmp8228 := PrimCons(symnumber, tmp8227)

tmp8229 := PrimCons(symnumber_2, tmp8228)

tmp8230 := PrimCons(symn_1_6string, tmp8229)

tmp8231 := PrimCons(symoccurs_2, tmp8230)

tmp8232 := PrimCons(symoccurs_1check, tmp8231)

tmp8233 := PrimCons(symoccurrences, tmp8232)

tmp8234 := PrimCons(symopen, tmp8233)

tmp8235 := PrimCons(symoptimise_2, tmp8234)

tmp8236 := PrimCons(symoptimise, tmp8235)

tmp8237 := PrimCons(symor, tmp8236)

tmp8238 := PrimCons(symos, tmp8237)

tmp8239 := PrimCons(symout, tmp8238)

tmp8240 := PrimCons(symoutput, tmp8239)

tmp8241 := PrimCons(sympackage, tmp8240)

tmp8242 := PrimCons(symport, tmp8241)

tmp8243 := PrimCons(symporters, tmp8242)

tmp8244 := PrimCons(sympos, tmp8243)

tmp8245 := PrimCons(sympr, tmp8244)

tmp8246 := PrimCons(symprint, tmp8245)

tmp8247 := PrimCons(symprolog_1memory, tmp8246)

tmp8248 := PrimCons(symprofile, tmp8247)

tmp8249 := PrimCons(symprofile_1results, tmp8248)

tmp8250 := PrimCons(symprotect, tmp8249)

tmp8251 := PrimCons(symprolog_2, tmp8250)

tmp8252 := PrimCons(symps, tmp8251)

tmp8253 := PrimCons(sympreclude_1all_1but, tmp8252)

tmp8254 := PrimCons(sympreclude, tmp8253)

tmp8255 := PrimCons(symput, tmp8254)

tmp8256 := PrimCons(sympackage_2, tmp8255)

tmp8257 := PrimCons(symread_1from_1string_1unprocessed, tmp8256)

tmp8258 := PrimCons(symread_1from_1string, tmp8257)

tmp8259 := PrimCons(symread_1byte, tmp8258)

tmp8260 := PrimCons(symread_1file_1as_1string, tmp8259)

tmp8261 := PrimCons(symread_1file_1as_1bytelist, tmp8260)

tmp8262 := PrimCons(symread_1file, tmp8261)

tmp8263 := PrimCons(symreceive, tmp8262)

tmp8264 := PrimCons(symread, tmp8263)

tmp8265 := PrimCons(symrelease, tmp8264)

tmp8266 := PrimCons(symremove, tmp8265)

tmp8267 := PrimCons(symretract, tmp8266)

tmp8268 := PrimCons(symreverse, tmp8267)

tmp8269 := PrimCons(symrun, tmp8268)

tmp8270 := PrimCons(symstr, tmp8269)

tmp8271 := PrimCons(symsave, tmp8270)

tmp8272 := PrimCons(symset, tmp8271)

tmp8273 := PrimCons(symsimple_1error, tmp8272)

tmp8274 := PrimCons(symsnd, tmp8273)

tmp8275 := PrimCons(symspecialise, tmp8274)

tmp8276 := PrimCons(symspy, tmp8275)

tmp8277 := PrimCons(symstep, tmp8276)

tmp8278 := PrimCons(symstoutput, tmp8277)

tmp8279 := PrimCons(symstinput, tmp8278)

tmp8280 := PrimCons(symstring, tmp8279)

tmp8281 := PrimCons(symstream, tmp8280)

tmp8282 := PrimCons(symstring_1_6n, tmp8281)

tmp8283 := PrimCons(symstring_2, tmp8282)

tmp8284 := PrimCons(symsubst, tmp8283)

tmp8285 := PrimCons(symsum, tmp8284)

tmp8286 := PrimCons(symstring_1_6symbol, tmp8285)

tmp8287 := PrimCons(symsymbol_2, tmp8286)

tmp8288 := PrimCons(symsymbol, tmp8287)

tmp8289 := PrimCons(symsystem_1S_2, tmp8288)

tmp8290 := PrimCons(symsynonyms, tmp8289)

tmp8291 := PrimCons(symsystemf, tmp8290)

tmp8292 := PrimCons(symtail, tmp8291)

tmp8293 := PrimCons(symtlv, tmp8292)

tmp8294 := PrimCons(symtlstr, tmp8293)

tmp8295 := PrimCons(symtl, tmp8294)

tmp8296 := PrimCons(symtc, tmp8295)

tmp8297 := PrimCons(symtc_2, tmp8296)

tmp8298 := PrimCons(symthaw, tmp8297)

tmp8299 := PrimCons(symtime, tmp8298)

tmp8300 := PrimCons(symtrack, tmp8299)

tmp8301 := PrimCons(symtracked, tmp8300)

tmp8302 := PrimCons(symtrap_1error, tmp8301)

tmp8303 := PrimCons(True, tmp8302)

tmp8304 := PrimCons(symtuple_2, tmp8303)

tmp8305 := PrimCons(symtype, tmp8304)

tmp8306 := PrimCons(symreturn, tmp8305)

tmp8307 := PrimCons(symundefmacro, tmp8306)

tmp8308 := PrimCons(symunprofile, tmp8307)

tmp8309 := PrimCons(symunput, tmp8308)

tmp8310 := PrimCons(symunion, tmp8309)

tmp8311 := PrimCons(symunix, tmp8310)

tmp8312 := PrimCons(symunit, tmp8311)

tmp8313 := PrimCons(symuntrack, tmp8312)

tmp8314 := PrimCons(symunspecialise, tmp8313)

tmp8315 := PrimCons(symupdate_1lambda_1table, tmp8314)

tmp8316 := PrimCons(symu_b, tmp8315)

tmp8317 := PrimCons(symuserdefs, tmp8316)

tmp8318 := PrimCons(symvector_2, tmp8317)

tmp8319 := PrimCons(symvector, tmp8318)

tmp8320 := PrimCons(sym_5_1vector, tmp8319)

tmp8321 := PrimCons(symvector_1_6, tmp8320)

tmp8322 := PrimCons(symvalue, tmp8321)

tmp8323 := PrimCons(symvar_2, tmp8322)

tmp8324 := PrimCons(symvariable_2, tmp8323)

tmp8325 := PrimCons(symverified, tmp8324)

tmp8326 := PrimCons(symversion, tmp8325)

tmp8327 := PrimCons(symwhen, tmp8326)

tmp8328 := PrimCons(symwhere, tmp8327)

tmp8329 := PrimCons(symwrite_1byte, tmp8328)

tmp8330 := PrimCons(symwrite_1to_1file, tmp8329)

tmp8331 := PrimCons(symy_1or_1n_2, tmp8330)

tmp8332 := PrimCons(tmp8111, tmp8331)

tmp8333 := PrimCons(sym_6_6, tmp8332)

tmp8334 := PrimCons(sym_5, tmp8333)

tmp8335 := PrimCons(sym_5_a, tmp8334)

tmp8336 := PrimCons(sym_7, tmp8335)

tmp8337 := PrimCons(sym_d, tmp8336)

tmp8338 := PrimCons(sym_c, tmp8337)

tmp8339 := PrimCons(sym_1, tmp8338)

tmp8340 := PrimCons(sym_3, tmp8339)

tmp8341 := PrimCons(sym_5end_6, tmp8340)

tmp8342 := PrimCons(sym_5_b_6, tmp8341)

tmp8343 := PrimCons(sym_c_4, tmp8342)

tmp8344 := PrimCons(sym_a_a_6, tmp8343)

tmp8345 := PrimCons(sym_6, tmp8344)

tmp8346 := PrimCons(sym_6_a, tmp8345)

tmp8347 := PrimCons(sym_a, tmp8346)

tmp8348 := PrimCons(sym_a_a, tmp8347)

tmp8349 := PrimCons(sym_5e_6, tmp8348)

tmp8350 := PrimCons(sym_1_6, tmp8349)

tmp8351 := PrimCons(sym_5_1, tmp8350)

tmp8352 := PrimCons(sym_dhush_d, tmp8351)

tmp8353 := PrimCons(sym_dporters_d, tmp8352)

tmp8354 := PrimCons(sym_dport_d, tmp8353)

tmp8355 := PrimCons(sym_8s, tmp8354)

tmp8356 := PrimCons(sym_8p, tmp8355)

tmp8357 := PrimCons(sym_8v, tmp8356)

tmp8358 := PrimCons(sym_dproperty_1vector_d, tmp8357)

tmp8359 := PrimCons(sym_drelease_d, tmp8358)

tmp8360 := PrimCons(sym_dos_d, tmp8359)

tmp8361 := PrimCons(sym_dmacros_d, tmp8360)

tmp8362 := PrimCons(sym_dmaximum_1print_1sequence_1size_d, tmp8361)

tmp8363 := PrimCons(sym_dversion_d, tmp8362)

tmp8364 := PrimCons(sym_dhome_1directory_d, tmp8363)

tmp8365 := PrimCons(sym_dstoutput_d, tmp8364)

tmp8366 := PrimCons(sym_dstinput_d, tmp8365)

tmp8367 := PrimCons(sym_dimplementation_d, tmp8366)

tmp8368 := PrimCons(sym_dlanguage_d, tmp8367)

tmp8369 := PrimCons(sym__, tmp8368)

tmp8370 := PrimCons(tmp8110, tmp8369)

tmp8371 := PrimCons(tmp8109, tmp8370)

tmp8372 := PrimCons(tmp8108, tmp8371)

tmp8373 := PrimCons(tmp8107, tmp8372)

tmp8374 := PrimCons(sym_e_e, tmp8373)

tmp8375 := PrimCons(sym_5_1_1, tmp8374)

tmp8376 := PrimCons(sym_1_1_6, tmp8375)

tmp8377 := PrimCons(sym_i, tmp8376)

tmp8378 := PrimCons(sym_j, tmp8377)

tmp8379 := PrimCons(sym_b, tmp8378)

tmp8380 := PrimValue(sym_dproperty_1vector_d)

tmp8381 := Call(__e, PrimFunc(symput), symshen, symshen_4external_1symbols, tmp8379, tmp8380)


_ = tmp8381

tmp8382 := MakeNative(func(__e *ControlFlow) {
V5798 := __e.Get(1)
_ = V5798
tmp8383 := MakeNative(func(__e *ControlFlow) {
W5799 := __e.Get(1)
_ = W5799
tmp8391 := PrimEqual(W5799, MakeNumber(-1))

var ifres8388 Obj

if True == tmp8391 {
ifres8388 = True


} else {
tmp8390 := PrimEqual(W5799, MakeNumber(0))

var ifres8389 Obj

if True == tmp8390 {
ifres8389 = True


} else {
ifres8389 = False


}

ifres8388 = ifres8389


}

if True == ifres8388 {
__e.Return(Nil)
return
} else {
tmp8384 := PrimCons(V5798, Nil)

tmp8385 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp8384, W5799)


tmp8386 := Call(__e, PrimFunc(symeval_1kl), tmp8385)


__e.Return(PrimCons(V5798, tmp8386))
return


}


}, 1)

tmp8392 := Call(__e, PrimFunc(symarity), V5798)


__e.TailApply(tmp8383, tmp8392)
return


}, 1)

tmp8393 := Call(__e, ns2_1set, symshen_4lambda_1entry, tmp8382)


_ = tmp8393

tmp8394 := MakeNative(func(__e *ControlFlow) {
V5800 := __e.Get(1)
_ = V5800
tmp8395 := MakeNative(func(__e *ControlFlow) {
W5801 := __e.Get(1)
_ = W5801
tmp8396 := MakeNative(func(__e *ControlFlow) {
Z5803 := __e.Get(1)
_ = Z5803
__e.TailApply(PrimFunc(symshen_4tuple), Z5803)
return
}, 1)

tmp8397 := PrimCons(symshen_4tuple, tmp8396)

tmp8398 := MakeNative(func(__e *ControlFlow) {
Z5804 := __e.Get(1)
_ = Z5804
__e.TailApply(PrimFunc(symshen_4pvar), Z5804)
return
}, 1)

tmp8399 := PrimCons(symshen_4pvar, tmp8398)

tmp8400 := MakeNative(func(__e *ControlFlow) {
Z5805 := __e.Get(1)
_ = Z5805
__e.TailApply(PrimFunc(symshen_4print_1prolog_1vector), Z5805)
return
}, 1)

tmp8401 := PrimCons(symshen_4print_1prolog_1vector, tmp8400)

tmp8402 := MakeNative(func(__e *ControlFlow) {
Z5806 := __e.Get(1)
_ = Z5806
__e.TailApply(PrimFunc(symshen_4print_1freshterm), Z5806)
return
}, 1)

tmp8403 := PrimCons(symshen_4print_1freshterm, tmp8402)

tmp8404 := MakeNative(func(__e *ControlFlow) {
Z5807 := __e.Get(1)
_ = Z5807
__e.TailApply(PrimFunc(symshen_4printF), Z5807)
return
}, 1)

tmp8405 := PrimCons(symshen_4printF, tmp8404)

tmp8406 := PrimCons(tmp8405, W5801)

tmp8407 := PrimCons(tmp8403, tmp8406)

tmp8408 := PrimCons(tmp8401, tmp8407)

tmp8409 := PrimCons(tmp8399, tmp8408)

tmp8410 := PrimCons(tmp8397, tmp8409)

__e.Return(PrimSet(symshen_4_dlambdatable_d, tmp8410))
return


}, 1)

tmp8411 := MakeNative(func(__e *ControlFlow) {
Z5802 := __e.Get(1)
_ = Z5802
__e.TailApply(PrimFunc(symshen_4lambda_1entry), Z5802)
return
}, 1)

tmp8412 := Call(__e, PrimFunc(symmap), tmp8411, V5800)


__e.TailApply(tmp8395, tmp8412)
return


}, 1)

tmp8413 := Call(__e, ns2_1set, symshen_4build_1lambda_1table, tmp8394)


_ = tmp8413

tmp8414 := Call(__e, PrimFunc(symexternal), symshen)


__e.TailApply(PrimFunc(symshen_4build_1lambda_1table), tmp8414)
return




}, 0)

