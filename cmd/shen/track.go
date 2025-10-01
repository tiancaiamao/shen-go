package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
tmp12748 := MakeNative(func(__e *ControlFlow) {
V4559 := __e.Get(1)
_ = V4559
tmp12749 := Call(__e, PrimFunc(symshen_4app), V4559, MakeString(";\n"), symshen_4a)


tmp12750 := PrimStringConcat(MakeString("partial function "), tmp12749)

tmp12751 := Call(__e, PrimFunc(symstoutput))


tmp12752 := Call(__e, PrimFunc(sympr), tmp12750, tmp12751)


_ = tmp12752

tmp12761 := Call(__e, PrimFunc(symshen_4tracked_2), V4559)


tmp12762 := PrimNot(tmp12761)

var ifres12756 Obj

if True == tmp12762 {
tmp12758 := Call(__e, PrimFunc(symshen_4app), V4559, MakeString("? "), symshen_4a)


tmp12759 := PrimStringConcat(MakeString("track "), tmp12758)

tmp12760 := Call(__e, PrimFunc(symy_1or_1n_2), tmp12759)


var ifres12757 Obj

if True == tmp12760 {
ifres12757 = True


} else {
ifres12757 = False


}

ifres12756 = ifres12757


} else {
ifres12756 = False


}

var ifres12753 Obj

if True == ifres12756 {
tmp12754 := Call(__e, PrimFunc(symps), V4559)


tmp12755 := Call(__e, PrimFunc(symshen_4track_1function), tmp12754)


ifres12753 = tmp12755


} else {
ifres12753 = symshen_4ok


}

_ = ifres12753

__e.Return(PrimSimpleError(MakeString("aborted")))
return


}, 1)

tmp12763 := Call(__e, ns2_1set, symshen_4f_1error, tmp12748)


_ = tmp12763

tmp12764 := MakeNative(func(__e *ControlFlow) {
V4560 := __e.Get(1)
_ = V4560
tmp12765 := PrimValue(symshen_4_dtracking_d)

__e.TailApply(PrimFunc(symelement_2), V4560, tmp12765)
return


}, 1)

tmp12766 := Call(__e, ns2_1set, symshen_4tracked_2, tmp12764)


_ = tmp12766

tmp12767 := MakeNative(func(__e *ControlFlow) {
V4561 := __e.Get(1)
_ = V4561
tmp12768 := MakeNative(func(__e *ControlFlow) {
Source := __e.Get(1)
_ = Source
__e.TailApply(PrimFunc(symshen_4track_1function), Source)
return
}, 1)

tmp12769 := Call(__e, PrimFunc(symps), V4561)


__e.TailApply(tmp12768, tmp12769)
return


}, 1)

tmp12770 := Call(__e, ns2_1set, symtrack, tmp12767)


_ = tmp12770

tmp12771 := MakeNative(func(__e *ControlFlow) {
V4564 := __e.Get(1)
_ = V4564
tmp12825 := PrimIsPair(V4564)

var ifres12799 Obj

if True == tmp12825 {
tmp12823 := PrimHead(V4564)

tmp12824 := PrimEqual(symdefun, tmp12823)

var ifres12801 Obj

if True == tmp12824 {
tmp12821 := PrimTail(V4564)

tmp12822 := PrimIsPair(tmp12821)

var ifres12803 Obj

if True == tmp12822 {
tmp12818 := PrimTail(V4564)

tmp12819 := PrimTail(tmp12818)

tmp12820 := PrimIsPair(tmp12819)

var ifres12805 Obj

if True == tmp12820 {
tmp12814 := PrimTail(V4564)

tmp12815 := PrimTail(tmp12814)

tmp12816 := PrimTail(tmp12815)

tmp12817 := PrimIsPair(tmp12816)

var ifres12807 Obj

if True == tmp12817 {
tmp12809 := PrimTail(V4564)

tmp12810 := PrimTail(tmp12809)

tmp12811 := PrimTail(tmp12810)

tmp12812 := PrimTail(tmp12811)

tmp12813 := PrimEqual(Nil, tmp12812)

var ifres12808 Obj

if True == tmp12813 {
ifres12808 = True


} else {
ifres12808 = False


}

ifres12807 = ifres12808


} else {
ifres12807 = False


}

var ifres12806 Obj

if True == ifres12807 {
ifres12806 = True


} else {
ifres12806 = False


}

ifres12805 = ifres12806


} else {
ifres12805 = False


}

var ifres12804 Obj

if True == ifres12805 {
ifres12804 = True


} else {
ifres12804 = False


}

ifres12803 = ifres12804


} else {
ifres12803 = False


}

var ifres12802 Obj

if True == ifres12803 {
ifres12802 = True


} else {
ifres12802 = False


}

ifres12801 = ifres12802


} else {
ifres12801 = False


}

var ifres12800 Obj

if True == ifres12801 {
ifres12800 = True


} else {
ifres12800 = False


}

ifres12799 = ifres12800


} else {
ifres12799 = False


}

if True == ifres12799 {
tmp12772 := MakeNative(func(__e *ControlFlow) {
KL := __e.Get(1)
_ = KL
tmp12773 := MakeNative(func(__e *ControlFlow) {
Ob := __e.Get(1)
_ = Ob
tmp12774 := MakeNative(func(__e *ControlFlow) {
Tr := __e.Get(1)
_ = Tr
__e.Return(Ob)
return
}, 1)

tmp12775 := PrimValue(symshen_4_dtracking_d)

tmp12776 := PrimCons(Ob, tmp12775)

tmp12777 := PrimSet(symshen_4_dtracking_d, tmp12776)

__e.TailApply(tmp12774, tmp12777)
return


}, 1)

tmp12778 := Call(__e, PrimFunc(symeval_1kl), KL)


__e.TailApply(tmp12773, tmp12778)
return


}, 1)

tmp12779 := PrimTail(V4564)

tmp12780 := PrimHead(tmp12779)

tmp12781 := PrimTail(V4564)

tmp12782 := PrimTail(tmp12781)

tmp12783 := PrimHead(tmp12782)

tmp12784 := PrimTail(V4564)

tmp12785 := PrimHead(tmp12784)

tmp12786 := PrimTail(V4564)

tmp12787 := PrimTail(tmp12786)

tmp12788 := PrimHead(tmp12787)

tmp12789 := PrimTail(V4564)

tmp12790 := PrimTail(tmp12789)

tmp12791 := PrimTail(tmp12790)

tmp12792 := PrimHead(tmp12791)

tmp12793 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp12785, tmp12788, tmp12792)


tmp12794 := PrimCons(tmp12793, Nil)

tmp12795 := PrimCons(tmp12783, tmp12794)

tmp12796 := PrimCons(tmp12780, tmp12795)

tmp12797 := PrimCons(symdefun, tmp12796)

__e.TailApply(tmp12772, tmp12797)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.track-function")))
return
}


}, 1)

tmp12826 := Call(__e, ns2_1set, symshen_4track_1function, tmp12771)


_ = tmp12826

tmp12827 := MakeNative(func(__e *ControlFlow) {
V4565 := __e.Get(1)
_ = V4565
V4566 := __e.Get(2)
_ = V4566
V4567 := __e.Get(3)
_ = V4567
tmp12828 := PrimCons(symshen_4_dcall_d, Nil)

tmp12829 := PrimCons(symvalue, tmp12828)

tmp12830 := PrimCons(MakeNumber(1), Nil)

tmp12831 := PrimCons(tmp12829, tmp12830)

tmp12832 := PrimCons(sym_7, tmp12831)

tmp12833 := PrimCons(tmp12832, Nil)

tmp12834 := PrimCons(symshen_4_dcall_d, tmp12833)

tmp12835 := PrimCons(symset, tmp12834)

tmp12836 := PrimCons(symshen_4_dcall_d, Nil)

tmp12837 := PrimCons(symvalue, tmp12836)

tmp12838 := Call(__e, PrimFunc(symshen_4prolog_1track), V4567, V4566)


tmp12839 := Call(__e, PrimFunc(symshen_4cons_1form), tmp12838)


tmp12840 := PrimCons(tmp12839, Nil)

tmp12841 := PrimCons(V4565, tmp12840)

tmp12842 := PrimCons(tmp12837, tmp12841)

tmp12843 := PrimCons(symshen_4input_1track, tmp12842)

tmp12844 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp12845 := Call(__e, PrimFunc(symprotect), symResult)


tmp12846 := PrimCons(symshen_4_dcall_d, Nil)

tmp12847 := PrimCons(symvalue, tmp12846)

tmp12848 := Call(__e, PrimFunc(symprotect), symResult)


tmp12849 := PrimCons(tmp12848, Nil)

tmp12850 := PrimCons(V4565, tmp12849)

tmp12851 := PrimCons(tmp12847, tmp12850)

tmp12852 := PrimCons(symshen_4output_1track, tmp12851)

tmp12853 := PrimCons(symshen_4_dcall_d, Nil)

tmp12854 := PrimCons(symvalue, tmp12853)

tmp12855 := PrimCons(MakeNumber(1), Nil)

tmp12856 := PrimCons(tmp12854, tmp12855)

tmp12857 := PrimCons(sym_1, tmp12856)

tmp12858 := PrimCons(tmp12857, Nil)

tmp12859 := PrimCons(symshen_4_dcall_d, tmp12858)

tmp12860 := PrimCons(symset, tmp12859)

tmp12861 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

tmp12862 := Call(__e, PrimFunc(symprotect), symResult)


tmp12863 := PrimCons(tmp12862, Nil)

tmp12864 := PrimCons(tmp12861, tmp12863)

tmp12865 := PrimCons(symdo, tmp12864)

tmp12866 := PrimCons(tmp12865, Nil)

tmp12867 := PrimCons(tmp12860, tmp12866)

tmp12868 := PrimCons(symdo, tmp12867)

tmp12869 := PrimCons(tmp12868, Nil)

tmp12870 := PrimCons(tmp12852, tmp12869)

tmp12871 := PrimCons(symdo, tmp12870)

tmp12872 := PrimCons(tmp12871, Nil)

tmp12873 := PrimCons(V4567, tmp12872)

tmp12874 := PrimCons(tmp12845, tmp12873)

tmp12875 := PrimCons(symlet, tmp12874)

tmp12876 := PrimCons(tmp12875, Nil)

tmp12877 := PrimCons(tmp12844, tmp12876)

tmp12878 := PrimCons(symdo, tmp12877)

tmp12879 := PrimCons(tmp12878, Nil)

tmp12880 := PrimCons(tmp12843, tmp12879)

tmp12881 := PrimCons(symdo, tmp12880)

tmp12882 := PrimCons(tmp12881, Nil)

tmp12883 := PrimCons(tmp12835, tmp12882)

__e.Return(PrimCons(symdo, tmp12883))
return


}, 3)

tmp12884 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp12827)


_ = tmp12884

tmp12885 := MakeNative(func(__e *ControlFlow) {
V4568 := __e.Get(1)
_ = V4568
V4569 := __e.Get(2)
_ = V4569
tmp12888 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V4568)


tmp12889 := PrimEqual(tmp12888, MakeNumber(0))

if True == tmp12889 {
__e.Return(V4569)
return
} else {
tmp12886 := Call(__e, PrimFunc(symshen_4vector_1parameter), V4569)


__e.TailApply(PrimFunc(symshen_4vector_1dereference), V4569, tmp12886)
return


}


}, 2)

tmp12890 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp12885)


_ = tmp12890

tmp12891 := MakeNative(func(__e *ControlFlow) {
V4572 := __e.Get(1)
_ = V4572
tmp12920 := PrimEqual(Nil, V4572)

if True == tmp12920 {
__e.Return(Nil)
return
} else {
tmp12918 := PrimIsPair(V4572)

var ifres12896 Obj

if True == tmp12918 {
tmp12916 := PrimTail(V4572)

tmp12917 := PrimIsPair(tmp12916)

var ifres12898 Obj

if True == tmp12917 {
tmp12913 := PrimTail(V4572)

tmp12914 := PrimTail(tmp12913)

tmp12915 := PrimIsPair(tmp12914)

var ifres12900 Obj

if True == tmp12915 {
tmp12909 := PrimTail(V4572)

tmp12910 := PrimTail(tmp12909)

tmp12911 := PrimTail(tmp12910)

tmp12912 := PrimIsPair(tmp12911)

var ifres12902 Obj

if True == tmp12912 {
tmp12904 := PrimTail(V4572)

tmp12905 := PrimTail(tmp12904)

tmp12906 := PrimTail(tmp12905)

tmp12907 := PrimTail(tmp12906)

tmp12908 := PrimEqual(Nil, tmp12907)

var ifres12903 Obj

if True == tmp12908 {
ifres12903 = True


} else {
ifres12903 = False


}

ifres12902 = ifres12903


} else {
ifres12902 = False


}

var ifres12901 Obj

if True == ifres12902 {
ifres12901 = True


} else {
ifres12901 = False


}

ifres12900 = ifres12901


} else {
ifres12900 = False


}

var ifres12899 Obj

if True == ifres12900 {
ifres12899 = True


} else {
ifres12899 = False


}

ifres12898 = ifres12899


} else {
ifres12898 = False


}

var ifres12897 Obj

if True == ifres12898 {
ifres12897 = True


} else {
ifres12897 = False


}

ifres12896 = ifres12897


} else {
ifres12896 = False


}

if True == ifres12896 {
__e.Return(PrimHead(V4572))
return
} else {
tmp12894 := PrimIsPair(V4572)

if True == tmp12894 {
tmp12892 := PrimTail(V4572)

__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp12892)
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1parameter)
return
}


}


}


}, 1)

tmp12921 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp12891)


_ = tmp12921

tmp12922 := MakeNative(func(__e *ControlFlow) {
V4575 := __e.Get(1)
_ = V4575
V4576 := __e.Get(2)
_ = V4576
tmp12956 := PrimEqual(Nil, V4576)

if True == tmp12956 {
__e.Return(V4575)
return
} else {
tmp12954 := PrimIsPair(V4575)

var ifres12932 Obj

if True == tmp12954 {
tmp12952 := PrimTail(V4575)

tmp12953 := PrimIsPair(tmp12952)

var ifres12934 Obj

if True == tmp12953 {
tmp12949 := PrimTail(V4575)

tmp12950 := PrimTail(tmp12949)

tmp12951 := PrimIsPair(tmp12950)

var ifres12936 Obj

if True == tmp12951 {
tmp12945 := PrimTail(V4575)

tmp12946 := PrimTail(tmp12945)

tmp12947 := PrimTail(tmp12946)

tmp12948 := PrimIsPair(tmp12947)

var ifres12938 Obj

if True == tmp12948 {
tmp12940 := PrimTail(V4575)

tmp12941 := PrimTail(tmp12940)

tmp12942 := PrimTail(tmp12941)

tmp12943 := PrimTail(tmp12942)

tmp12944 := PrimEqual(Nil, tmp12943)

var ifres12939 Obj

if True == tmp12944 {
ifres12939 = True


} else {
ifres12939 = False


}

ifres12938 = ifres12939


} else {
ifres12938 = False


}

var ifres12937 Obj

if True == ifres12938 {
ifres12937 = True


} else {
ifres12937 = False


}

ifres12936 = ifres12937


} else {
ifres12936 = False


}

var ifres12935 Obj

if True == ifres12936 {
ifres12935 = True


} else {
ifres12935 = False


}

ifres12934 = ifres12935


} else {
ifres12934 = False


}

var ifres12933 Obj

if True == ifres12934 {
ifres12933 = True


} else {
ifres12933 = False


}

ifres12932 = ifres12933


} else {
ifres12932 = False


}

if True == ifres12932 {
__e.Return(V4575)
return
} else {
tmp12930 := PrimIsPair(V4575)

if True == tmp12930 {
tmp12923 := PrimHead(V4575)

tmp12924 := PrimCons(V4576, Nil)

tmp12925 := PrimCons(tmp12923, tmp12924)

tmp12926 := PrimCons(symshen_4deref, tmp12925)

tmp12927 := PrimTail(V4575)

tmp12928 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp12927, V4576)


__e.Return(PrimCons(tmp12926, tmp12928))
return


} else {
__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1dereference)
return
}


}


}


}, 2)

tmp12957 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp12922)


_ = tmp12957

tmp12958 := MakeNative(func(__e *ControlFlow) {
V4579 := __e.Get(1)
_ = V4579
tmp12962 := PrimEqual(sym_7, V4579)

if True == tmp12962 {
__e.Return(PrimSet(symshen_4_dstep_d, True))
return
} else {
tmp12960 := PrimEqual(sym_1, V4579)

if True == tmp12960 {
__e.Return(PrimSet(symshen_4_dstep_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
return
}


}


}, 1)

tmp12963 := Call(__e, ns2_1set, symstep, tmp12958)


_ = tmp12963

tmp12964 := MakeNative(func(__e *ControlFlow) {
V4582 := __e.Get(1)
_ = V4582
tmp12968 := PrimEqual(sym_7, V4582)

if True == tmp12968 {
__e.Return(PrimSet(symshen_4_dspy_d, True))
return
} else {
tmp12966 := PrimEqual(sym_1, V4582)

if True == tmp12966 {
__e.Return(PrimSet(symshen_4_dspy_d, False))
return
} else {
__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
return
}


}


}, 1)

tmp12969 := Call(__e, ns2_1set, symspy, tmp12964)


_ = tmp12969

tmp12970 := MakeNative(func(__e *ControlFlow) {
tmp12974 := PrimValue(symshen_4_dstep_d)

if True == tmp12974 {
tmp12971 := PrimValue(sym_dstinput_d)

tmp12972 := PrimReadByte(tmp12971)

__e.TailApply(PrimFunc(symshen_4check_1byte), tmp12972)
return


} else {
__e.TailApply(PrimFunc(symnl), MakeNumber(1))
return
}


}, 0)

tmp12975 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp12970)


_ = tmp12975

tmp12976 := MakeNative(func(__e *ControlFlow) {
V4585 := __e.Get(1)
_ = V4585
tmp12978 := PrimEqual(MakeNumber(94), V4585)

if True == tmp12978 {
__e.Return(PrimSimpleError(MakeString("aborted")))
return
} else {
__e.Return(True)
return
}


}, 1)

tmp12979 := Call(__e, ns2_1set, symshen_4check_1byte, tmp12976)


_ = tmp12979

tmp12980 := MakeNative(func(__e *ControlFlow) {
V4586 := __e.Get(1)
_ = V4586
V4587 := __e.Get(2)
_ = V4587
V4588 := __e.Get(3)
_ = V4588
tmp12981 := Call(__e, PrimFunc(symshen_4spaces), V4586)


tmp12982 := Call(__e, PrimFunc(symshen_4spaces), V4586)


tmp12983 := Call(__e, PrimFunc(symshen_4app), tmp12982, MakeString(""), symshen_4a)


tmp12984 := PrimStringConcat(MakeString(" \n"), tmp12983)

tmp12985 := Call(__e, PrimFunc(symshen_4app), V4587, tmp12984, symshen_4a)


tmp12986 := PrimStringConcat(MakeString("> Inputs to "), tmp12985)

tmp12987 := Call(__e, PrimFunc(symshen_4app), V4586, tmp12986, symshen_4a)


tmp12988 := PrimStringConcat(MakeString("<"), tmp12987)

tmp12989 := Call(__e, PrimFunc(symshen_4app), tmp12981, tmp12988, symshen_4a)


tmp12990 := PrimStringConcat(MakeString("\n"), tmp12989)

tmp12991 := Call(__e, PrimFunc(symstoutput))


tmp12992 := Call(__e, PrimFunc(sympr), tmp12990, tmp12991)


_ = tmp12992

__e.TailApply(PrimFunc(symshen_4recursively_1print), V4588)
return


}, 3)

tmp12993 := Call(__e, ns2_1set, symshen_4input_1track, tmp12980)


_ = tmp12993

tmp12994 := MakeNative(func(__e *ControlFlow) {
V4591 := __e.Get(1)
_ = V4591
tmp13004 := PrimEqual(Nil, V4591)

if True == tmp13004 {
tmp12995 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp12995)
return


} else {
tmp13002 := PrimIsPair(V4591)

if True == tmp13002 {
tmp12996 := PrimHead(V4591)

tmp12997 := Call(__e, PrimFunc(symprint), tmp12996)


_ = tmp12997

tmp12998 := Call(__e, PrimFunc(symstoutput))


tmp12999 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp12998)


_ = tmp12999

tmp13000 := PrimTail(V4591)

__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp13000)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursively-print")))
return
}


}


}, 1)

tmp13005 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp12994)


_ = tmp13005

tmp13006 := MakeNative(func(__e *ControlFlow) {
V4592 := __e.Get(1)
_ = V4592
tmp13010 := PrimEqual(MakeNumber(0), V4592)

if True == tmp13010 {
__e.Return(MakeString(""))
return
} else {
tmp13007 := PrimNumberSubtract(V4592, MakeNumber(1))

tmp13008 := Call(__e, PrimFunc(symshen_4spaces), tmp13007)


__e.Return(PrimStringConcat(MakeString(" "), tmp13008))
return


}


}, 1)

tmp13011 := Call(__e, ns2_1set, symshen_4spaces, tmp13006)


_ = tmp13011

tmp13012 := MakeNative(func(__e *ControlFlow) {
V4593 := __e.Get(1)
_ = V4593
V4594 := __e.Get(2)
_ = V4594
V4595 := __e.Get(3)
_ = V4595
tmp13013 := Call(__e, PrimFunc(symshen_4spaces), V4593)


tmp13014 := Call(__e, PrimFunc(symshen_4spaces), V4593)


tmp13015 := Call(__e, PrimFunc(symshen_4app), V4595, MakeString(""), symshen_4s)


tmp13016 := PrimStringConcat(MakeString("==> "), tmp13015)

tmp13017 := Call(__e, PrimFunc(symshen_4app), tmp13014, tmp13016, symshen_4a)


tmp13018 := PrimStringConcat(MakeString(" \n"), tmp13017)

tmp13019 := Call(__e, PrimFunc(symshen_4app), V4594, tmp13018, symshen_4a)


tmp13020 := PrimStringConcat(MakeString("> Output from "), tmp13019)

tmp13021 := Call(__e, PrimFunc(symshen_4app), V4593, tmp13020, symshen_4a)


tmp13022 := PrimStringConcat(MakeString("<"), tmp13021)

tmp13023 := Call(__e, PrimFunc(symshen_4app), tmp13013, tmp13022, symshen_4a)


tmp13024 := PrimStringConcat(MakeString("\n"), tmp13023)

tmp13025 := Call(__e, PrimFunc(symstoutput))


__e.TailApply(PrimFunc(sympr), tmp13024, tmp13025)
return


}, 3)

tmp13026 := Call(__e, ns2_1set, symshen_4output_1track, tmp13012)


_ = tmp13026

tmp13027 := MakeNative(func(__e *ControlFlow) {
V4596 := __e.Get(1)
_ = V4596
tmp13028 := PrimValue(symshen_4_dtracking_d)

tmp13029 := Call(__e, PrimFunc(symremove), V4596, tmp13028)


tmp13030 := PrimSet(symshen_4_dtracking_d, tmp13029)

_ = tmp13030

tmp13031 := MakeNative(func(__e *ControlFlow) {
tmp13032 := Call(__e, PrimFunc(symps), V4596)


__e.TailApply(PrimFunc(symeval), tmp13032)
return


}, 0)

tmp13033 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(V4596)
return
}, 1)

__e.TailApply(try_1catch, tmp13031, tmp13033)
return


}, 1)

tmp13034 := Call(__e, ns2_1set, symuntrack, tmp13027)


_ = tmp13034

tmp13035 := MakeNative(func(__e *ControlFlow) {
V4597 := __e.Get(1)
_ = V4597
V4598 := __e.Get(2)
_ = V4598
__e.TailApply(PrimFunc(symshen_4remove_1h), V4597, V4598, Nil)
return
}, 2)

tmp13036 := Call(__e, ns2_1set, symremove, tmp13035)


_ = tmp13036

tmp13037 := MakeNative(func(__e *ControlFlow) {
V4608 := __e.Get(1)
_ = V4608
V4609 := __e.Get(2)
_ = V4609
V4610 := __e.Get(3)
_ = V4610
tmp13052 := PrimEqual(Nil, V4609)

if True == tmp13052 {
__e.TailApply(PrimFunc(symreverse), V4610)
return
} else {
tmp13050 := PrimIsPair(V4609)

var ifres13046 Obj

if True == tmp13050 {
tmp13048 := PrimHead(V4609)

tmp13049 := PrimEqual(V4608, tmp13048)

var ifres13047 Obj

if True == tmp13049 {
ifres13047 = True


} else {
ifres13047 = False


}

ifres13046 = ifres13047


} else {
ifres13046 = False


}

if True == ifres13046 {
tmp13038 := PrimHead(V4609)

tmp13039 := PrimTail(V4609)

__e.TailApply(PrimFunc(symshen_4remove_1h), tmp13038, tmp13039, V4610)
return


} else {
tmp13044 := PrimIsPair(V4609)

if True == tmp13044 {
tmp13040 := PrimTail(V4609)

tmp13041 := PrimHead(V4609)

tmp13042 := PrimCons(tmp13041, V4610)

__e.TailApply(PrimFunc(symshen_4remove_1h), V4608, tmp13040, tmp13042)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-h")))
return
}


}


}


}, 3)

tmp13053 := Call(__e, ns2_1set, symshen_4remove_1h, tmp13037)


_ = tmp13053

tmp13054 := MakeNative(func(__e *ControlFlow) {
V4611 := __e.Get(1)
_ = V4611
tmp13055 := PrimValue(symshen_4_dprofiled_d)

tmp13056 := PrimCons(V4611, tmp13055)

tmp13057 := PrimSet(symshen_4_dprofiled_d, tmp13056)

_ = tmp13057

tmp13058 := Call(__e, PrimFunc(symps), V4611)


__e.TailApply(PrimFunc(symshen_4profile_1help), tmp13058)
return


}, 1)

tmp13059 := Call(__e, ns2_1set, symprofile, tmp13054)


_ = tmp13059

tmp13060 := MakeNative(func(__e *ControlFlow) {
V4614 := __e.Get(1)
_ = V4614
tmp13130 := PrimIsPair(V4614)

var ifres13104 Obj

if True == tmp13130 {
tmp13128 := PrimHead(V4614)

tmp13129 := PrimEqual(symdefun, tmp13128)

var ifres13106 Obj

if True == tmp13129 {
tmp13126 := PrimTail(V4614)

tmp13127 := PrimIsPair(tmp13126)

var ifres13108 Obj

if True == tmp13127 {
tmp13123 := PrimTail(V4614)

tmp13124 := PrimTail(tmp13123)

tmp13125 := PrimIsPair(tmp13124)

var ifres13110 Obj

if True == tmp13125 {
tmp13119 := PrimTail(V4614)

tmp13120 := PrimTail(tmp13119)

tmp13121 := PrimTail(tmp13120)

tmp13122 := PrimIsPair(tmp13121)

var ifres13112 Obj

if True == tmp13122 {
tmp13114 := PrimTail(V4614)

tmp13115 := PrimTail(tmp13114)

tmp13116 := PrimTail(tmp13115)

tmp13117 := PrimTail(tmp13116)

tmp13118 := PrimEqual(Nil, tmp13117)

var ifres13113 Obj

if True == tmp13118 {
ifres13113 = True


} else {
ifres13113 = False


}

ifres13112 = ifres13113


} else {
ifres13112 = False


}

var ifres13111 Obj

if True == ifres13112 {
ifres13111 = True


} else {
ifres13111 = False


}

ifres13110 = ifres13111


} else {
ifres13110 = False


}

var ifres13109 Obj

if True == ifres13110 {
ifres13109 = True


} else {
ifres13109 = False


}

ifres13108 = ifres13109


} else {
ifres13108 = False


}

var ifres13107 Obj

if True == ifres13108 {
ifres13107 = True


} else {
ifres13107 = False


}

ifres13106 = ifres13107


} else {
ifres13106 = False


}

var ifres13105 Obj

if True == ifres13106 {
ifres13105 = True


} else {
ifres13105 = False


}

ifres13104 = ifres13105


} else {
ifres13104 = False


}

if True == ifres13104 {
tmp13061 := MakeNative(func(__e *ControlFlow) {
G := __e.Get(1)
_ = G
tmp13062 := MakeNative(func(__e *ControlFlow) {
Profile := __e.Get(1)
_ = Profile
tmp13063 := MakeNative(func(__e *ControlFlow) {
Def := __e.Get(1)
_ = Def
tmp13064 := MakeNative(func(__e *ControlFlow) {
CompileProfile := __e.Get(1)
_ = CompileProfile
tmp13065 := MakeNative(func(__e *ControlFlow) {
CompileG := __e.Get(1)
_ = CompileG
tmp13066 := PrimTail(V4614)

__e.Return(PrimHead(tmp13066))
return


}, 1)

tmp13067 := Call(__e, PrimFunc(symeval_1kl), Def)


__e.TailApply(tmp13065, tmp13067)
return


}, 1)

tmp13068 := Call(__e, PrimFunc(symeval_1kl), Profile)


__e.TailApply(tmp13064, tmp13068)
return


}, 1)

tmp13069 := PrimTail(V4614)

tmp13070 := PrimTail(tmp13069)

tmp13071 := PrimHead(tmp13070)

tmp13072 := PrimTail(V4614)

tmp13073 := PrimHead(tmp13072)

tmp13074 := PrimTail(V4614)

tmp13075 := PrimTail(tmp13074)

tmp13076 := PrimTail(tmp13075)

tmp13077 := PrimHead(tmp13076)

tmp13078 := Call(__e, PrimFunc(symsubst), G, tmp13073, tmp13077)


tmp13079 := PrimCons(tmp13078, Nil)

tmp13080 := PrimCons(tmp13071, tmp13079)

tmp13081 := PrimCons(G, tmp13080)

tmp13082 := PrimCons(symdefun, tmp13081)

__e.TailApply(tmp13063, tmp13082)
return


}, 1)

tmp13083 := PrimTail(V4614)

tmp13084 := PrimHead(tmp13083)

tmp13085 := PrimTail(V4614)

tmp13086 := PrimTail(tmp13085)

tmp13087 := PrimHead(tmp13086)

tmp13088 := PrimTail(V4614)

tmp13089 := PrimHead(tmp13088)

tmp13090 := PrimTail(V4614)

tmp13091 := PrimTail(tmp13090)

tmp13092 := PrimHead(tmp13091)

tmp13093 := PrimTail(V4614)

tmp13094 := PrimTail(tmp13093)

tmp13095 := PrimHead(tmp13094)

tmp13096 := PrimCons(G, tmp13095)

tmp13097 := Call(__e, PrimFunc(symshen_4profile_1func), tmp13089, tmp13092, tmp13096)


tmp13098 := PrimCons(tmp13097, Nil)

tmp13099 := PrimCons(tmp13087, tmp13098)

tmp13100 := PrimCons(tmp13084, tmp13099)

tmp13101 := PrimCons(symdefun, tmp13100)

__e.TailApply(tmp13062, tmp13101)
return


}, 1)

tmp13102 := Call(__e, PrimFunc(symgensym), symshen_4f)


__e.TailApply(tmp13061, tmp13102)
return


} else {
__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
return
}


}, 1)

tmp13131 := Call(__e, ns2_1set, symshen_4profile_1help, tmp13060)


_ = tmp13131

tmp13132 := MakeNative(func(__e *ControlFlow) {
V4615 := __e.Get(1)
_ = V4615
tmp13133 := PrimValue(symshen_4_dprofiled_d)

tmp13134 := Call(__e, PrimFunc(symremove), V4615, tmp13133)


tmp13135 := PrimSet(symshen_4_dprofiled_d, tmp13134)

_ = tmp13135

tmp13136 := MakeNative(func(__e *ControlFlow) {
tmp13137 := Call(__e, PrimFunc(symps), V4615)


__e.TailApply(PrimFunc(symeval), tmp13137)
return


}, 0)

tmp13138 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(V4615)
return
}, 1)

__e.TailApply(try_1catch, tmp13136, tmp13138)
return


}, 1)

tmp13139 := Call(__e, ns2_1set, symunprofile, tmp13132)


_ = tmp13139

tmp13140 := MakeNative(func(__e *ControlFlow) {
V4616 := __e.Get(1)
_ = V4616
tmp13141 := PrimValue(symshen_4_dprofiled_d)

__e.TailApply(PrimFunc(symelement_2), V4616, tmp13141)
return


}, 1)

tmp13142 := Call(__e, ns2_1set, symshen_4profiled_2, tmp13140)


_ = tmp13142

tmp13143 := MakeNative(func(__e *ControlFlow) {
V4617 := __e.Get(1)
_ = V4617
V4618 := __e.Get(2)
_ = V4618
V4619 := __e.Get(3)
_ = V4619
tmp13144 := Call(__e, PrimFunc(symprotect), symStart)


tmp13145 := PrimCons(symrun, Nil)

tmp13146 := PrimCons(symget_1time, tmp13145)

tmp13147 := Call(__e, PrimFunc(symprotect), symResult)


tmp13148 := Call(__e, PrimFunc(symprotect), symFinish)


tmp13149 := PrimCons(symrun, Nil)

tmp13150 := PrimCons(symget_1time, tmp13149)

tmp13151 := Call(__e, PrimFunc(symprotect), symStart)


tmp13152 := PrimCons(tmp13151, Nil)

tmp13153 := PrimCons(tmp13150, tmp13152)

tmp13154 := PrimCons(sym_1, tmp13153)

tmp13155 := Call(__e, PrimFunc(symprotect), symRecord)


tmp13156 := PrimCons(V4617, Nil)

tmp13157 := PrimCons(symshen_4get_1profile, tmp13156)

tmp13158 := Call(__e, PrimFunc(symprotect), symFinish)


tmp13159 := PrimCons(tmp13158, Nil)

tmp13160 := PrimCons(tmp13157, tmp13159)

tmp13161 := PrimCons(sym_7, tmp13160)

tmp13162 := PrimCons(tmp13161, Nil)

tmp13163 := PrimCons(V4617, tmp13162)

tmp13164 := PrimCons(symshen_4put_1profile, tmp13163)

tmp13165 := Call(__e, PrimFunc(symprotect), symResult)


tmp13166 := PrimCons(tmp13165, Nil)

tmp13167 := PrimCons(tmp13164, tmp13166)

tmp13168 := PrimCons(tmp13155, tmp13167)

tmp13169 := PrimCons(symlet, tmp13168)

tmp13170 := PrimCons(tmp13169, Nil)

tmp13171 := PrimCons(tmp13154, tmp13170)

tmp13172 := PrimCons(tmp13148, tmp13171)

tmp13173 := PrimCons(symlet, tmp13172)

tmp13174 := PrimCons(tmp13173, Nil)

tmp13175 := PrimCons(V4619, tmp13174)

tmp13176 := PrimCons(tmp13147, tmp13175)

tmp13177 := PrimCons(symlet, tmp13176)

tmp13178 := PrimCons(tmp13177, Nil)

tmp13179 := PrimCons(tmp13146, tmp13178)

tmp13180 := PrimCons(tmp13144, tmp13179)

__e.Return(PrimCons(symlet, tmp13180))
return


}, 3)

tmp13181 := Call(__e, ns2_1set, symshen_4profile_1func, tmp13143)


_ = tmp13181

tmp13182 := MakeNative(func(__e *ControlFlow) {
V4620 := __e.Get(1)
_ = V4620
tmp13183 := MakeNative(func(__e *ControlFlow) {
Results := __e.Get(1)
_ = Results
tmp13184 := MakeNative(func(__e *ControlFlow) {
Initialise := __e.Get(1)
_ = Initialise
__e.TailApply(PrimFunc(sym_8p), V4620, Results)
return
}, 1)

tmp13185 := Call(__e, PrimFunc(symshen_4put_1profile), V4620, MakeNumber(0))


__e.TailApply(tmp13184, tmp13185)
return


}, 1)

tmp13186 := Call(__e, PrimFunc(symshen_4get_1profile), V4620)


__e.TailApply(tmp13183, tmp13186)
return


}, 1)

tmp13187 := Call(__e, ns2_1set, symprofile_1results, tmp13182)


_ = tmp13187

tmp13188 := MakeNative(func(__e *ControlFlow) {
V4621 := __e.Get(1)
_ = V4621
tmp13189 := MakeNative(func(__e *ControlFlow) {
tmp13190 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V4621, symprofile, tmp13190)
return


}, 0)

tmp13191 := MakeNative(func(__e *ControlFlow) {
E := __e.Get(1)
_ = E
__e.Return(MakeNumber(0))
return
}, 1)

__e.TailApply(try_1catch, tmp13189, tmp13191)
return


}, 1)

tmp13192 := Call(__e, ns2_1set, symshen_4get_1profile, tmp13188)


_ = tmp13192

tmp13193 := MakeNative(func(__e *ControlFlow) {
V4622 := __e.Get(1)
_ = V4622
V4623 := __e.Get(2)
_ = V4623
tmp13194 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V4622, symprofile, V4623, tmp13194)
return


}, 2)

__e.TailApply(ns2_1set, symshen_4put_1profile, tmp13193)
return




}, 0)

