package main

import . "github.com/tiancaiamao/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
tmp4727 := MakeNative(func(__e *ControlFlow) {
V2160 := __e.Get(1)
_ = V2160
tmp4728 := MakeNative(func(__e *ControlFlow) {
W2161 := __e.Get(1)
_ = W2161
tmp4729 := MakeNative(func(__e *ControlFlow) {
W2162 := __e.Get(1)
_ = W2162
tmp4730 := MakeNative(func(__e *ControlFlow) {
W2165 := __e.Get(1)
_ = W2165
__e.Return(W2165)
return
}, 1)

tmp4731 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2162)


__e.TailApply(tmp4730, tmp4731)
return


}, 1)

tmp4732 := MakeNative(func(__e *ControlFlow) {
tmp4733 := MakeNative(func(__e *ControlFlow) {
Z2163 := __e.Get(1)
_ = Z2163
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2163)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4733, W2161)
return


}, 0)

tmp4734 := MakeNative(func(__e *ControlFlow) {
Z2164 := __e.Get(1)
_ = Z2164
tmp4735 := PrimValue(symshen_4_dresidue_d)

__e.TailApply(PrimFunc(symshen_4print_1residue), tmp4735)
return


}, 1)

tmp4736 := Call(__e, try_1catch, tmp4732, tmp4734)


__e.TailApply(tmp4729, tmp4736)
return


}, 1)

tmp4737 := PrimReadFileAsByteList(V2160)

__e.TailApply(tmp4728, tmp4737)
return


}, 1)

tmp4738 := Call(__e, ns2_1set, symread_1file, tmp4727)


_ = tmp4738

tmp4739 := MakeNative(func(__e *ControlFlow) {
V2166 := __e.Get(1)
_ = V2166
tmp4740 := MakeNative(func(__e *ControlFlow) {
W2167 := __e.Get(1)
_ = W2167
__e.TailApply(PrimFunc(symshen_4nchars), MakeNumber(50), V2166)
return
}, 1)

tmp4741 := Call(__e, PrimFunc(symstoutput))


tmp4742 := Call(__e, PrimFunc(sympr), MakeString("syntax error here:\n"), tmp4741)


__e.TailApply(tmp4740, tmp4742)
return


}, 1)

tmp4743 := Call(__e, ns2_1set, symshen_4print_1residue, tmp4739)


_ = tmp4743

tmp4744 := MakeNative(func(__e *ControlFlow) {
V2172 := __e.Get(1)
_ = V2172
V2173 := __e.Get(2)
_ = V2173
tmp4760 := PrimEqual(MakeNumber(0), V2172)

if True == tmp4760 {
tmp4745 := Call(__e, PrimFunc(symstoutput))


tmp4746 := Call(__e, PrimFunc(sympr), MakeString(" ..."), tmp4745)


_ = tmp4746

__e.TailApply(PrimFunc(symabort))
return


} else {
tmp4758 := PrimEqual(Nil, V2173)

if True == tmp4758 {
tmp4747 := Call(__e, PrimFunc(symstoutput))


tmp4748 := Call(__e, PrimFunc(sympr), MakeString(" ..."), tmp4747)


_ = tmp4748

__e.TailApply(PrimFunc(symabort))
return


} else {
tmp4756 := PrimIsPair(V2173)

if True == tmp4756 {
tmp4749 := PrimHead(V2173)

tmp4750 := PrimNumberToString(tmp4749)

tmp4751 := Call(__e, PrimFunc(symstoutput))


tmp4752 := Call(__e, PrimFunc(sympr), tmp4750, tmp4751)


_ = tmp4752

tmp4753 := PrimNumberSubtract(V2172, MakeNumber(1))

tmp4754 := PrimTail(V2173)

__e.TailApply(PrimFunc(symshen_4nchars), tmp4753, tmp4754)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.nchars")))
return
}


}


}


}, 2)

tmp4761 := Call(__e, ns2_1set, symshen_4nchars, tmp4744)


_ = tmp4761

tmp4762 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dit_d))
return
}, 0)

tmp4763 := Call(__e, ns2_1set, symit, tmp4762)


_ = tmp4763

tmp4764 := MakeNative(func(__e *ControlFlow) {
V2174 := __e.Get(1)
_ = V2174
tmp4765 := MakeNative(func(__e *ControlFlow) {
W2175 := __e.Get(1)
_ = W2175
tmp4766 := MakeNative(func(__e *ControlFlow) {
W2176 := __e.Get(1)
_ = W2176
tmp4767 := MakeNative(func(__e *ControlFlow) {
W2177 := __e.Get(1)
_ = W2177
tmp4768 := MakeNative(func(__e *ControlFlow) {
W2178 := __e.Get(1)
_ = W2178
__e.TailApply(PrimFunc(symreverse), W2177)
return
}, 1)

tmp4769 := PrimCloseStream(W2175)

__e.TailApply(tmp4768, tmp4769)
return


}, 1)

tmp4770 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), W2175, W2176, Nil)


__e.TailApply(tmp4767, tmp4770)
return


}, 1)

tmp4771 := PrimReadByte(W2175)

__e.TailApply(tmp4766, tmp4771)
return


}, 1)

tmp4772 := PrimOpenStream(V2174, symin)

__e.TailApply(tmp4765, tmp4772)
return


}, 1)

tmp4773 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp4764)


_ = tmp4773

tmp4774 := MakeNative(func(__e *ControlFlow) {
V2179 := __e.Get(1)
_ = V2179
V2180 := __e.Get(2)
_ = V2180
V2181 := __e.Get(3)
_ = V2181
tmp4778 := PrimEqual(MakeNumber(-1), V2180)

if True == tmp4778 {
__e.Return(V2181)
return
} else {
tmp4775 := PrimReadByte(V2179)

tmp4776 := PrimCons(V2180, V2181)

__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2179, tmp4775, tmp4776)
return


}


}, 3)

tmp4779 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp4774)


_ = tmp4779

tmp4780 := MakeNative(func(__e *ControlFlow) {
V2182 := __e.Get(1)
_ = V2182
tmp4781 := MakeNative(func(__e *ControlFlow) {
W2183 := __e.Get(1)
_ = W2183
tmp4782 := PrimReadByte(W2183)

__e.TailApply(PrimFunc(symshen_4rfas_1h), W2183, tmp4782, MakeString(""))
return


}, 1)

tmp4783 := PrimOpenStream(V2182, symin)

__e.TailApply(tmp4781, tmp4783)
return


}, 1)

tmp4784 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp4780)


_ = tmp4784

tmp4785 := MakeNative(func(__e *ControlFlow) {
V2184 := __e.Get(1)
_ = V2184
V2185 := __e.Get(2)
_ = V2185
V2186 := __e.Get(3)
_ = V2186
tmp4791 := PrimEqual(MakeNumber(-1), V2185)

if True == tmp4791 {
tmp4786 := PrimCloseStream(V2184)

_ = tmp4786

__e.Return(V2186)
return


} else {
tmp4787 := PrimReadByte(V2184)

tmp4788 := PrimNumberToString(V2185)

tmp4789 := PrimStringConcat(V2186, tmp4788)

__e.TailApply(PrimFunc(symshen_4rfas_1h), V2184, tmp4787, tmp4789)
return


}


}, 3)

tmp4792 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp4785)


_ = tmp4792

tmp4793 := MakeNative(func(__e *ControlFlow) {
V2187 := __e.Get(1)
_ = V2187
tmp4794 := Call(__e, PrimFunc(symread), V2187)


__e.TailApply(PrimFunc(symeval_1kl), tmp4794)
return


}, 1)

tmp4795 := Call(__e, ns2_1set, syminput, tmp4793)


_ = tmp4795

tmp4796 := MakeNative(func(__e *ControlFlow) {
V2188 := __e.Get(1)
_ = V2188
V2189 := __e.Get(2)
_ = V2189
tmp4797 := MakeNative(func(__e *ControlFlow) {
W2190 := __e.Get(1)
_ = W2190
tmp4798 := MakeNative(func(__e *ControlFlow) {
W2191 := __e.Get(1)
_ = W2191
tmp4804 := Call(__e, PrimFunc(symshen_4rectify_1type), V2188)


tmp4805 := Call(__e, PrimFunc(symshen_4typecheck), W2191, tmp4804)


tmp4806 := PrimEqual(False, tmp4805)

if True == tmp4806 {
tmp4799 := Call(__e, PrimFunc(symshen_4app), V2188, MakeString("\n"), symshen_4r)


tmp4800 := PrimStringConcat(MakeString(" is not of type "), tmp4799)

tmp4801 := Call(__e, PrimFunc(symshen_4app), W2191, tmp4800, symshen_4r)


tmp4802 := PrimStringConcat(MakeString("type error: "), tmp4801)

__e.Return(PrimSimpleError(tmp4802))
return


} else {
__e.TailApply(PrimFunc(symeval_1kl), W2191)
return
}


}, 1)

tmp4807 := Call(__e, PrimFunc(symread), V2189)


__e.TailApply(tmp4798, tmp4807)
return


}, 1)

tmp4808 := Call(__e, PrimFunc(symshen_4monotype), V2188)


__e.TailApply(tmp4797, tmp4808)
return


}, 2)

tmp4809 := Call(__e, ns2_1set, syminput_7, tmp4796)


_ = tmp4809

tmp4810 := MakeNative(func(__e *ControlFlow) {
V2192 := __e.Get(1)
_ = V2192
tmp4817 := PrimIsPair(V2192)

if True == tmp4817 {
tmp4811 := MakeNative(func(__e *ControlFlow) {
Z2193 := __e.Get(1)
_ = Z2193
__e.TailApply(PrimFunc(symshen_4monotype), Z2193)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4811, V2192)
return


} else {
tmp4815 := PrimIsVariable(V2192)

if True == tmp4815 {
tmp4812 := Call(__e, PrimFunc(symshen_4app), V2192, MakeString("\n"), symshen_4a)


tmp4813 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp4812)

__e.Return(PrimSimpleError(tmp4813))
return


} else {
__e.Return(V2192)
return
}


}


}, 1)

tmp4818 := Call(__e, ns2_1set, symshen_4monotype, tmp4810)


_ = tmp4818

tmp4819 := MakeNative(func(__e *ControlFlow) {
V2194 := __e.Get(1)
_ = V2194
tmp4820 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2194)


tmp4821 := MakeNative(func(__e *ControlFlow) {
Z2195 := __e.Get(1)
_ = Z2195
__e.TailApply(PrimFunc(symshen_4return_2), Z2195)
return
}, 1)

__e.TailApply(PrimFunc(symshen_4read_1loop), V2194, tmp4820, Nil, tmp4821)
return


}, 1)

tmp4822 := Call(__e, ns2_1set, symlineread, tmp4819)


_ = tmp4822

tmp4823 := MakeNative(func(__e *ControlFlow) {
V2196 := __e.Get(1)
_ = V2196
tmp4824 := MakeNative(func(__e *ControlFlow) {
W2197 := __e.Get(1)
_ = W2197
tmp4825 := MakeNative(func(__e *ControlFlow) {
W2198 := __e.Get(1)
_ = W2198
tmp4826 := MakeNative(func(__e *ControlFlow) {
W2200 := __e.Get(1)
_ = W2200
__e.Return(W2200)
return
}, 1)

tmp4827 := Call(__e, PrimFunc(symshen_4process_1sexprs), W2198)


__e.TailApply(tmp4826, tmp4827)
return


}, 1)

tmp4828 := MakeNative(func(__e *ControlFlow) {
Z2199 := __e.Get(1)
_ = Z2199
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2199)
return
}, 1)

tmp4829 := Call(__e, PrimFunc(symcompile), tmp4828, W2197)


__e.TailApply(tmp4825, tmp4829)
return


}, 1)

tmp4830 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2196)


__e.TailApply(tmp4824, tmp4830)
return


}, 1)

tmp4831 := Call(__e, ns2_1set, symread_1from_1string, tmp4823)


_ = tmp4831

tmp4832 := MakeNative(func(__e *ControlFlow) {
V2201 := __e.Get(1)
_ = V2201
tmp4833 := MakeNative(func(__e *ControlFlow) {
W2202 := __e.Get(1)
_ = W2202
tmp4834 := MakeNative(func(__e *ControlFlow) {
W2203 := __e.Get(1)
_ = W2203
__e.Return(W2203)
return
}, 1)

tmp4835 := MakeNative(func(__e *ControlFlow) {
Z2204 := __e.Get(1)
_ = Z2204
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2204)
return
}, 1)

tmp4836 := Call(__e, PrimFunc(symcompile), tmp4835, W2202)


__e.TailApply(tmp4834, tmp4836)
return


}, 1)

tmp4837 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2201)


__e.TailApply(tmp4833, tmp4837)
return


}, 1)

tmp4838 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp4832)


_ = tmp4838

tmp4839 := MakeNative(func(__e *ControlFlow) {
V2205 := __e.Get(1)
_ = V2205
tmp4847 := PrimEqual(MakeString(""), V2205)

if True == tmp4847 {
__e.Return(Nil)
return
} else {
tmp4845 := Call(__e, PrimFunc(symshen_4_7string_2), V2205)


if True == tmp4845 {
tmp4840 := Call(__e, PrimFunc(symhdstr), V2205)


tmp4841 := PrimStringToNumber(tmp4840)

tmp4842 := PrimTailString(V2205)

tmp4843 := Call(__e, PrimFunc(symshen_4str_1_6bytes), tmp4842)


__e.Return(PrimCons(tmp4841, tmp4843))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.str->bytes")))
return
}


}


}, 1)

tmp4848 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp4839)


_ = tmp4848

tmp4849 := MakeNative(func(__e *ControlFlow) {
V2206 := __e.Get(1)
_ = V2206
tmp4850 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2206)


tmp4851 := MakeNative(func(__e *ControlFlow) {
Z2207 := __e.Get(1)
_ = Z2207
__e.TailApply(PrimFunc(symshen_4whitespace_2), Z2207)
return
}, 1)

tmp4852 := Call(__e, PrimFunc(symshen_4read_1loop), V2206, tmp4850, Nil, tmp4851)


__e.Return(PrimHead(tmp4852))
return


}, 1)

tmp4853 := Call(__e, ns2_1set, symread, tmp4849)


_ = tmp4853

tmp4854 := MakeNative(func(__e *ControlFlow) {
V2208 := __e.Get(1)
_ = V2208
tmp4857 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2208)


if True == tmp4857 {
tmp4855 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2208)


__e.Return(PrimStringToNumber(tmp4855))
return


} else {
__e.Return(PrimReadByte(V2208))
return
}


}, 1)

tmp4858 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp4854)


_ = tmp4858

tmp4859 := MakeNative(func(__e *ControlFlow) {
V2213 := __e.Get(1)
_ = V2213
V2214 := __e.Get(2)
_ = V2214
V2215 := __e.Get(3)
_ = V2215
V2216 := __e.Get(4)
_ = V2216
tmp4882 := PrimEqual(MakeNumber(94), V2214)

if True == tmp4882 {
__e.Return(PrimSimpleError(MakeString("read aborted")))
return
} else {
tmp4880 := PrimEqual(MakeNumber(-1), V2214)

if True == tmp4880 {
tmp4862 := Call(__e, PrimFunc(symempty_2), V2215)


if True == tmp4862 {
__e.Return(PrimSimpleError(MakeString("error: empty stream")))
return
} else {
tmp4860 := MakeNative(func(__e *ControlFlow) {
Z2217 := __e.Get(1)
_ = Z2217
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2217)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4860, V2215)
return


}


} else {
tmp4878 := PrimEqual(MakeNumber(0), V2214)

if True == tmp4878 {
tmp4863 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2213)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2213, tmp4863, V2215, V2216)
return


} else {
tmp4876 := Call(__e, V2216, V2214)


if True == tmp4876 {
tmp4864 := MakeNative(func(__e *ControlFlow) {
W2218 := __e.Get(1)
_ = W2218
tmp4870 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2218)


if True == tmp4870 {
tmp4865 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2213)


tmp4866 := PrimCons(V2214, Nil)

tmp4867 := Call(__e, PrimFunc(symappend), V2215, tmp4866)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2213, tmp4865, tmp4867, V2216)
return


} else {
tmp4868 := Call(__e, PrimFunc(symshen_4record_1it), V2215)


_ = tmp4868

__e.Return(W2218)
return


}


}, 1)

tmp4871 := Call(__e, PrimFunc(symshen_4try_1parse), V2215)


__e.TailApply(tmp4864, tmp4871)
return


} else {
tmp4872 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2213)


tmp4873 := PrimCons(V2214, Nil)

tmp4874 := Call(__e, PrimFunc(symappend), V2215, tmp4873)


__e.TailApply(PrimFunc(symshen_4read_1loop), V2213, tmp4872, tmp4874, V2216)
return


}


}


}


}


}, 4)

tmp4883 := Call(__e, ns2_1set, symshen_4read_1loop, tmp4859)


_ = tmp4883

tmp4884 := MakeNative(func(__e *ControlFlow) {
V2219 := __e.Get(1)
_ = V2219
tmp4885 := MakeNative(func(__e *ControlFlow) {
W2220 := __e.Get(1)
_ = W2220
tmp4887 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), W2220)


if True == tmp4887 {
__e.Return(symshen_4i_1failed_b)
return
} else {
__e.TailApply(PrimFunc(symshen_4process_1sexprs), W2220)
return
}


}, 1)

tmp4888 := MakeNative(func(__e *ControlFlow) {
tmp4889 := MakeNative(func(__e *ControlFlow) {
Z2221 := __e.Get(1)
_ = Z2221
__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), Z2221)
return
}, 1)

__e.TailApply(PrimFunc(symcompile), tmp4889, V2219)
return


}, 0)

tmp4890 := MakeNative(func(__e *ControlFlow) {
Z2222 := __e.Get(1)
_ = Z2222
__e.Return(symshen_4i_1failed_b)
return
}, 1)

tmp4891 := Call(__e, try_1catch, tmp4888, tmp4890)


__e.TailApply(tmp4885, tmp4891)
return


}, 1)

tmp4892 := Call(__e, ns2_1set, symshen_4try_1parse, tmp4884)


_ = tmp4892

tmp4893 := MakeNative(func(__e *ControlFlow) {
V2225 := __e.Get(1)
_ = V2225
tmp4897 := PrimEqual(symshen_4i_1failed_b, V2225)

if True == tmp4897 {
__e.Return(True)
return
} else {
tmp4895 := PrimEqual(Nil, V2225)

if True == tmp4895 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp4898 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp4893)


_ = tmp4898

tmp4899 := MakeNative(func(__e *ControlFlow) {
V2226 := __e.Get(1)
_ = V2226
tmp4900 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2226)


__e.Return(PrimSet(symshen_4_dit_d, tmp4900))
return


}, 1)

tmp4901 := Call(__e, ns2_1set, symshen_4record_1it, tmp4899)


_ = tmp4901

tmp4902 := MakeNative(func(__e *ControlFlow) {
V2227 := __e.Get(1)
_ = V2227
tmp4910 := PrimEqual(Nil, V2227)

if True == tmp4910 {
__e.Return(MakeString(""))
return
} else {
tmp4908 := PrimIsPair(V2227)

if True == tmp4908 {
tmp4903 := PrimHead(V2227)

tmp4904 := PrimNumberToString(tmp4903)

tmp4905 := PrimTail(V2227)

tmp4906 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp4905)


__e.Return(PrimStringConcat(tmp4904, tmp4906))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.bytes->string")))
return
}


}


}, 1)

tmp4911 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp4902)


_ = tmp4911

tmp4912 := MakeNative(func(__e *ControlFlow) {
V2228 := __e.Get(1)
_ = V2228
tmp4913 := MakeNative(func(__e *ControlFlow) {
W2229 := __e.Get(1)
_ = W2229
tmp4914 := MakeNative(func(__e *ControlFlow) {
W2230 := __e.Get(1)
_ = W2230
tmp4915 := MakeNative(func(__e *ControlFlow) {
W2231 := __e.Get(1)
_ = W2231
tmp4916 := MakeNative(func(__e *ControlFlow) {
Z2232 := __e.Get(1)
_ = Z2232
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2232, W2231)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4916, W2229)
return


}, 1)

tmp4917 := Call(__e, PrimFunc(symshen_4find_1types), W2229)


__e.TailApply(tmp4915, tmp4917)
return


}, 1)

tmp4918 := Call(__e, PrimFunc(symshen_4find_1arities), W2229)


__e.TailApply(tmp4914, tmp4918)
return


}, 1)

tmp4919 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2228)


__e.TailApply(tmp4913, tmp4919)
return


}, 1)

tmp4920 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp4912)


_ = tmp4920

tmp4921 := MakeNative(func(__e *ControlFlow) {
V2233 := __e.Get(1)
_ = V2233
tmp4943 := PrimIsPair(V2233)

var ifres4934 Obj

if True == tmp4943 {
tmp4941 := PrimTail(V2233)

tmp4942 := PrimIsPair(tmp4941)

var ifres4936 Obj

if True == tmp4942 {
tmp4938 := PrimHead(V2233)

tmp4939 := PrimIntern(MakeString(":"))

tmp4940 := PrimEqual(tmp4938, tmp4939)

var ifres4937 Obj

if True == tmp4940 {
ifres4937 = True


} else {
ifres4937 = False


}

ifres4936 = ifres4937


} else {
ifres4936 = False


}

var ifres4935 Obj

if True == ifres4936 {
ifres4935 = True


} else {
ifres4935 = False


}

ifres4934 = ifres4935


} else {
ifres4934 = False


}

if True == ifres4934 {
tmp4922 := PrimTail(V2233)

tmp4923 := PrimHead(tmp4922)

tmp4924 := PrimTail(V2233)

tmp4925 := PrimTail(tmp4924)

tmp4926 := Call(__e, PrimFunc(symshen_4find_1types), tmp4925)


__e.Return(PrimCons(tmp4923, tmp4926))
return


} else {
tmp4932 := PrimIsPair(V2233)

if True == tmp4932 {
tmp4927 := PrimHead(V2233)

tmp4928 := Call(__e, PrimFunc(symshen_4find_1types), tmp4927)


tmp4929 := PrimTail(V2233)

tmp4930 := Call(__e, PrimFunc(symshen_4find_1types), tmp4929)


__e.TailApply(PrimFunc(symappend), tmp4928, tmp4930)
return


} else {
__e.Return(Nil)
return
}


}


}, 1)

tmp4944 := Call(__e, ns2_1set, symshen_4find_1types, tmp4921)


_ = tmp4944

tmp4945 := MakeNative(func(__e *ControlFlow) {
V2236 := __e.Get(1)
_ = V2236
tmp4994 := PrimIsPair(V2236)

var ifres4975 Obj

if True == tmp4994 {
tmp4992 := PrimHead(V2236)

tmp4993 := PrimEqual(symdefine, tmp4992)

var ifres4977 Obj

if True == tmp4993 {
tmp4990 := PrimTail(V2236)

tmp4991 := PrimIsPair(tmp4990)

var ifres4979 Obj

if True == tmp4991 {
tmp4987 := PrimTail(V2236)

tmp4988 := PrimTail(tmp4987)

tmp4989 := PrimIsPair(tmp4988)

var ifres4981 Obj

if True == tmp4989 {
tmp4983 := PrimTail(V2236)

tmp4984 := PrimTail(tmp4983)

tmp4985 := PrimHead(tmp4984)

tmp4986 := PrimEqual(sym_i, tmp4985)

var ifres4982 Obj

if True == tmp4986 {
ifres4982 = True


} else {
ifres4982 = False


}

ifres4981 = ifres4982


} else {
ifres4981 = False


}

var ifres4980 Obj

if True == ifres4981 {
ifres4980 = True


} else {
ifres4980 = False


}

ifres4979 = ifres4980


} else {
ifres4979 = False


}

var ifres4978 Obj

if True == ifres4979 {
ifres4978 = True


} else {
ifres4978 = False


}

ifres4977 = ifres4978


} else {
ifres4977 = False


}

var ifres4976 Obj

if True == ifres4977 {
ifres4976 = True


} else {
ifres4976 = False


}

ifres4975 = ifres4976


} else {
ifres4975 = False


}

if True == ifres4975 {
tmp4946 := PrimTail(V2236)

tmp4947 := PrimHead(tmp4946)

tmp4948 := PrimTail(V2236)

tmp4949 := PrimHead(tmp4948)

tmp4950 := PrimTail(V2236)

tmp4951 := PrimTail(tmp4950)

tmp4952 := PrimTail(tmp4951)

tmp4953 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4949, MakeNumber(1), tmp4952)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4947, tmp4953)
return


} else {
tmp4973 := PrimIsPair(V2236)

var ifres4965 Obj

if True == tmp4973 {
tmp4971 := PrimHead(V2236)

tmp4972 := PrimEqual(symdefine, tmp4971)

var ifres4967 Obj

if True == tmp4972 {
tmp4969 := PrimTail(V2236)

tmp4970 := PrimIsPair(tmp4969)

var ifres4968 Obj

if True == tmp4970 {
ifres4968 = True


} else {
ifres4968 = False


}

ifres4967 = ifres4968


} else {
ifres4967 = False


}

var ifres4966 Obj

if True == ifres4967 {
ifres4966 = True


} else {
ifres4966 = False


}

ifres4965 = ifres4966


} else {
ifres4965 = False


}

if True == ifres4965 {
tmp4954 := PrimTail(V2236)

tmp4955 := PrimHead(tmp4954)

tmp4956 := PrimTail(V2236)

tmp4957 := PrimHead(tmp4956)

tmp4958 := PrimTail(V2236)

tmp4959 := PrimTail(tmp4958)

tmp4960 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4957, MakeNumber(0), tmp4959)


__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4955, tmp4960)
return


} else {
tmp4963 := PrimIsPair(V2236)

if True == tmp4963 {
tmp4961 := MakeNative(func(__e *ControlFlow) {
Z2237 := __e.Get(1)
_ = Z2237
__e.TailApply(PrimFunc(symshen_4find_1arities), Z2237)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp4961, V2236)
return


} else {
__e.Return(symshen_4skip)
return
}


}


}


}, 1)

tmp4995 := Call(__e, ns2_1set, symshen_4find_1arities, tmp4945)


_ = tmp4995

tmp4996 := MakeNative(func(__e *ControlFlow) {
V2238 := __e.Get(1)
_ = V2238
V2239 := __e.Get(2)
_ = V2239
tmp4997 := MakeNative(func(__e *ControlFlow) {
W2240 := __e.Get(1)
_ = W2240
tmp5008 := PrimEqual(W2240, MakeNumber(-1))

if True == tmp5008 {
__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2238, V2239)
return
} else {
tmp5006 := PrimEqual(W2240, V2239)

if True == tmp5006 {
__e.Return(symshen_4skip)
return
} else {
tmp5004 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2238)


if True == tmp5004 {
tmp4998 := Call(__e, PrimFunc(symshen_4app), V2238, MakeString(" is a system function\n"), symshen_4a)


__e.Return(PrimSimpleError(tmp4998))
return


} else {
tmp4999 := Call(__e, PrimFunc(symshen_4app), V2238, MakeString(" may cause errors\n"), symshen_4a)


tmp5000 := PrimStringConcat(MakeString("changing the arity of "), tmp4999)

tmp5001 := Call(__e, PrimFunc(symstoutput))


tmp5002 := Call(__e, PrimFunc(sympr), tmp5000, tmp5001)


_ = tmp5002

__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2238, V2239)
return


}


}


}


}, 1)

tmp5009 := Call(__e, PrimFunc(symarity), V2238)


__e.TailApply(tmp4997, tmp5009)
return


}, 2)

tmp5010 := Call(__e, ns2_1set, symshen_4store_1arity, tmp4996)


_ = tmp5010

tmp5011 := MakeNative(func(__e *ControlFlow) {
V2241 := __e.Get(1)
_ = V2241
V2242 := __e.Get(2)
_ = V2242
tmp5016 := PrimEqual(MakeNumber(0), V2242)

if True == tmp5016 {
tmp5012 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2241, symarity, MakeNumber(0), tmp5012)
return


} else {
tmp5013 := PrimValue(sym_dproperty_1vector_d)

tmp5014 := Call(__e, PrimFunc(symput), V2241, symarity, V2242, tmp5013)


_ = tmp5014

__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2241, V2242)
return


}


}, 2)

tmp5017 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp5011)


_ = tmp5017

tmp5018 := MakeNative(func(__e *ControlFlow) {
V2243 := __e.Get(1)
_ = V2243
V2244 := __e.Get(2)
_ = V2244
tmp5019 := MakeNative(func(__e *ControlFlow) {
W2245 := __e.Get(1)
_ = W2245
tmp5020 := MakeNative(func(__e *ControlFlow) {
W2246 := __e.Get(1)
_ = W2246
tmp5021 := MakeNative(func(__e *ControlFlow) {
W2247 := __e.Get(1)
_ = W2247
tmp5022 := MakeNative(func(__e *ControlFlow) {
W2248 := __e.Get(1)
_ = W2248
__e.Return(W2248)
return
}, 1)

tmp5023 := PrimSet(symshen_4_dlambdatable_d, W2247)

__e.TailApply(tmp5022, tmp5023)
return


}, 1)

tmp5024 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2243, W2246, W2245)


__e.TailApply(tmp5021, tmp5024)
return


}, 1)

tmp5025 := PrimCons(V2243, Nil)

tmp5026 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5025, V2244)


tmp5027 := Call(__e, PrimFunc(symeval_1kl), tmp5026)


__e.TailApply(tmp5020, tmp5027)
return


}, 1)

tmp5028 := PrimValue(symshen_4_dlambdatable_d)

__e.TailApply(tmp5019, tmp5028)
return


}, 2)

tmp5029 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp5018)


_ = tmp5029

tmp5030 := MakeNative(func(__e *ControlFlow) {
V2251 := __e.Get(1)
_ = V2251
V2252 := __e.Get(2)
_ = V2252
tmp5048 := PrimEqual(MakeNumber(0), V2252)

if True == tmp5048 {
__e.Return(symshen_4skip)
return
} else {
tmp5046 := PrimEqual(MakeNumber(1), V2252)

if True == tmp5046 {
tmp5031 := MakeNative(func(__e *ControlFlow) {
W2253 := __e.Get(1)
_ = W2253
tmp5032 := PrimCons(W2253, Nil)

tmp5033 := Call(__e, PrimFunc(symappend), V2251, tmp5032)


tmp5034 := PrimCons(tmp5033, Nil)

tmp5035 := PrimCons(W2253, tmp5034)

__e.Return(PrimCons(symlambda, tmp5035))
return


}, 1)

tmp5036 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5031, tmp5036)
return


} else {
tmp5037 := MakeNative(func(__e *ControlFlow) {
W2254 := __e.Get(1)
_ = W2254
tmp5038 := PrimCons(W2254, Nil)

tmp5039 := Call(__e, PrimFunc(symappend), V2251, tmp5038)


tmp5040 := PrimNumberSubtract(V2252, MakeNumber(1))

tmp5041 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp5039, tmp5040)


tmp5042 := PrimCons(tmp5041, Nil)

tmp5043 := PrimCons(W2254, tmp5042)

__e.Return(PrimCons(symlambda, tmp5043))
return


}, 1)

tmp5044 := Call(__e, PrimFunc(symgensym), symY)


__e.TailApply(tmp5037, tmp5044)
return


}


}


}, 2)

tmp5049 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp5030)


_ = tmp5049

tmp5050 := MakeNative(func(__e *ControlFlow) {
V2264 := __e.Get(1)
_ = V2264
V2265 := __e.Get(2)
_ = V2265
V2266 := __e.Get(3)
_ = V2266
tmp5073 := PrimEqual(Nil, V2266)

if True == tmp5073 {
tmp5051 := PrimCons(V2264, V2265)

__e.Return(PrimCons(tmp5051, Nil))
return


} else {
tmp5071 := PrimIsPair(V2266)

var ifres5062 Obj

if True == tmp5071 {
tmp5069 := PrimHead(V2266)

tmp5070 := PrimIsPair(tmp5069)

var ifres5064 Obj

if True == tmp5070 {
tmp5066 := PrimHead(V2266)

tmp5067 := PrimHead(tmp5066)

tmp5068 := PrimEqual(V2264, tmp5067)

var ifres5065 Obj

if True == tmp5068 {
ifres5065 = True


} else {
ifres5065 = False


}

ifres5064 = ifres5065


} else {
ifres5064 = False


}

var ifres5063 Obj

if True == ifres5064 {
ifres5063 = True


} else {
ifres5063 = False


}

ifres5062 = ifres5063


} else {
ifres5062 = False


}

if True == ifres5062 {
tmp5052 := PrimHead(V2266)

tmp5053 := PrimHead(tmp5052)

tmp5054 := PrimCons(tmp5053, V2265)

tmp5055 := PrimTail(V2266)

__e.Return(PrimCons(tmp5054, tmp5055))
return


} else {
tmp5060 := PrimIsPair(V2266)

if True == tmp5060 {
tmp5056 := PrimHead(V2266)

tmp5057 := PrimTail(V2266)

tmp5058 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2264, V2265, tmp5057)


__e.Return(PrimCons(tmp5056, tmp5058))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.assoc->")))
return
}


}


}


}, 3)

tmp5074 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp5050)


_ = tmp5074

tmp5075 := MakeNative(func(__e *ControlFlow) {
V2281 := __e.Get(1)
_ = V2281
V2282 := __e.Get(2)
_ = V2282
V2283 := __e.Get(3)
_ = V2283
tmp5122 := PrimEqual(MakeNumber(0), V2282)

var ifres5115 Obj

if True == tmp5122 {
tmp5121 := PrimIsPair(V2283)

var ifres5117 Obj

if True == tmp5121 {
tmp5119 := PrimHead(V2283)

tmp5120 := PrimEqual(tmp5119, sym_1_6)

var ifres5118 Obj

if True == tmp5120 {
ifres5118 = True


} else {
ifres5118 = False


}

ifres5117 = ifres5118


} else {
ifres5117 = False


}

var ifres5116 Obj

if True == ifres5117 {
ifres5116 = True


} else {
ifres5116 = False


}

ifres5115 = ifres5116


} else {
ifres5115 = False


}

if True == ifres5115 {
__e.Return(MakeNumber(0))
return
} else {
tmp5113 := PrimEqual(MakeNumber(0), V2282)

var ifres5106 Obj

if True == tmp5113 {
tmp5112 := PrimIsPair(V2283)

var ifres5108 Obj

if True == tmp5112 {
tmp5110 := PrimHead(V2283)

tmp5111 := PrimEqual(tmp5110, sym_5_1)

var ifres5109 Obj

if True == tmp5111 {
ifres5109 = True


} else {
ifres5109 = False


}

ifres5108 = ifres5109


} else {
ifres5108 = False


}

var ifres5107 Obj

if True == ifres5108 {
ifres5107 = True


} else {
ifres5107 = False


}

ifres5106 = ifres5107


} else {
ifres5106 = False


}

if True == ifres5106 {
__e.Return(MakeNumber(0))
return
} else {
tmp5104 := PrimEqual(MakeNumber(0), V2282)

var ifres5101 Obj

if True == tmp5104 {
tmp5103 := PrimIsPair(V2283)

var ifres5102 Obj

if True == tmp5103 {
ifres5102 = True


} else {
ifres5102 = False


}

ifres5101 = ifres5102


} else {
ifres5101 = False


}

if True == ifres5101 {
tmp5076 := PrimTail(V2283)

tmp5077 := Call(__e, PrimFunc(symshen_4find_1arity), V2281, MakeNumber(0), tmp5076)


__e.Return(PrimNumberAdd(MakeNumber(1), tmp5077))
return


} else {
tmp5099 := PrimEqual(MakeNumber(1), V2282)

var ifres5092 Obj

if True == tmp5099 {
tmp5098 := PrimIsPair(V2283)

var ifres5094 Obj

if True == tmp5098 {
tmp5096 := PrimHead(V2283)

tmp5097 := PrimEqual(sym_j, tmp5096)

var ifres5095 Obj

if True == tmp5097 {
ifres5095 = True


} else {
ifres5095 = False


}

ifres5094 = ifres5095


} else {
ifres5094 = False


}

var ifres5093 Obj

if True == ifres5094 {
ifres5093 = True


} else {
ifres5093 = False


}

ifres5092 = ifres5093


} else {
ifres5092 = False


}

if True == ifres5092 {
tmp5078 := PrimTail(V2283)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2281, MakeNumber(0), tmp5078)
return


} else {
tmp5090 := PrimEqual(MakeNumber(1), V2282)

var ifres5087 Obj

if True == tmp5090 {
tmp5089 := PrimIsPair(V2283)

var ifres5088 Obj

if True == tmp5089 {
ifres5088 = True


} else {
ifres5088 = False


}

ifres5087 = ifres5088


} else {
ifres5087 = False


}

if True == ifres5087 {
tmp5079 := PrimTail(V2283)

__e.TailApply(PrimFunc(symshen_4find_1arity), V2281, MakeNumber(1), tmp5079)
return


} else {
tmp5085 := PrimEqual(MakeNumber(1), V2282)

if True == tmp5085 {
tmp5080 := Call(__e, PrimFunc(symshen_4app), V2281, MakeString(" definition: missing }\n"), symshen_4a)


tmp5081 := PrimStringConcat(MakeString("syntax error in "), tmp5080)

__e.Return(PrimSimpleError(tmp5081))
return


} else {
tmp5082 := Call(__e, PrimFunc(symshen_4app), V2281, MakeString(" definition: missing -> or <-\n"), symshen_4a)


tmp5083 := PrimStringConcat(MakeString("syntax error in "), tmp5082)

__e.Return(PrimSimpleError(tmp5083))
return


}


}


}


}


}


}


}, 3)

tmp5123 := Call(__e, ns2_1set, symshen_4find_1arity, tmp5075)


_ = tmp5123

tmp5124 := MakeNative(func(__e *ControlFlow) {
V2284 := __e.Get(1)
_ = V2284
tmp5125 := MakeNative(func(__e *ControlFlow) {
W2285 := __e.Get(1)
_ = W2285
tmp5370 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2285)


if True == tmp5370 {
tmp5126 := MakeNative(func(__e *ControlFlow) {
W2296 := __e.Get(1)
_ = W2296
tmp5338 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2296)


if True == tmp5338 {
tmp5127 := MakeNative(func(__e *ControlFlow) {
W2307 := __e.Get(1)
_ = W2307
tmp5320 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2307)


if True == tmp5320 {
tmp5128 := MakeNative(func(__e *ControlFlow) {
W2313 := __e.Get(1)
_ = W2313
tmp5302 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2313)


if True == tmp5302 {
tmp5129 := MakeNative(func(__e *ControlFlow) {
W2319 := __e.Get(1)
_ = W2319
tmp5284 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2319)


if True == tmp5284 {
tmp5130 := MakeNative(func(__e *ControlFlow) {
W2325 := __e.Get(1)
_ = W2325
tmp5265 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2325)


if True == tmp5265 {
tmp5131 := MakeNative(func(__e *ControlFlow) {
W2331 := __e.Get(1)
_ = W2331
tmp5240 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2331)


if True == tmp5240 {
tmp5132 := MakeNative(func(__e *ControlFlow) {
W2339 := __e.Get(1)
_ = W2339
tmp5221 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2339)


if True == tmp5221 {
tmp5133 := MakeNative(func(__e *ControlFlow) {
W2345 := __e.Get(1)
_ = W2345
tmp5202 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2345)


if True == tmp5202 {
tmp5134 := MakeNative(func(__e *ControlFlow) {
W2351 := __e.Get(1)
_ = W2351
tmp5185 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2351)


if True == tmp5185 {
tmp5135 := MakeNative(func(__e *ControlFlow) {
W2357 := __e.Get(1)
_ = W2357
tmp5165 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2357)


if True == tmp5165 {
tmp5136 := MakeNative(func(__e *ControlFlow) {
W2364 := __e.Get(1)
_ = W2364
tmp5148 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2364)


if True == tmp5148 {
tmp5137 := MakeNative(func(__e *ControlFlow) {
W2370 := __e.Get(1)
_ = W2370
tmp5139 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2370)


if True == tmp5139 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2370)
return
}


}, 1)

tmp5140 := MakeNative(func(__e *ControlFlow) {
W2371 := __e.Get(1)
_ = W2371
tmp5144 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2371)


if True == tmp5144 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5141 := MakeNative(func(__e *ControlFlow) {
W2372 := __e.Get(1)
_ = W2372
__e.TailApply(PrimFunc(symshen_4comb), W2372, Nil)
return
}, 1)

tmp5142 := Call(__e, PrimFunc(symshen_4in_1_6), W2371)


__e.TailApply(tmp5141, tmp5142)
return


}


}, 1)

tmp5145 := Call(__e, PrimFunc(sym_5e_6), V2284)


tmp5146 := Call(__e, tmp5140, tmp5145)


__e.TailApply(tmp5137, tmp5146)
return


} else {
__e.Return(W2364)
return
}


}, 1)

tmp5149 := MakeNative(func(__e *ControlFlow) {
W2365 := __e.Get(1)
_ = W2365
tmp5161 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2365)


if True == tmp5161 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5150 := MakeNative(func(__e *ControlFlow) {
W2366 := __e.Get(1)
_ = W2366
tmp5151 := MakeNative(func(__e *ControlFlow) {
W2367 := __e.Get(1)
_ = W2367
tmp5157 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2367)


if True == tmp5157 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5152 := MakeNative(func(__e *ControlFlow) {
W2368 := __e.Get(1)
_ = W2368
tmp5153 := MakeNative(func(__e *ControlFlow) {
W2369 := __e.Get(1)
_ = W2369
__e.TailApply(PrimFunc(symshen_4comb), W2369, W2368)
return
}, 1)

tmp5154 := Call(__e, PrimFunc(symshen_4in_1_6), W2367)


__e.TailApply(tmp5153, tmp5154)
return


}, 1)

tmp5155 := Call(__e, PrimFunc(symshen_4_5_1out), W2367)


__e.TailApply(tmp5152, tmp5155)
return


}


}, 1)

tmp5158 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2366)


__e.TailApply(tmp5151, tmp5158)
return


}, 1)

tmp5159 := Call(__e, PrimFunc(symshen_4in_1_6), W2365)


__e.TailApply(tmp5150, tmp5159)
return


}


}, 1)

tmp5162 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2284)


tmp5163 := Call(__e, tmp5149, tmp5162)


__e.TailApply(tmp5136, tmp5163)
return


} else {
__e.Return(W2357)
return
}


}, 1)

tmp5166 := MakeNative(func(__e *ControlFlow) {
W2358 := __e.Get(1)
_ = W2358
tmp5181 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2358)


if True == tmp5181 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5167 := MakeNative(func(__e *ControlFlow) {
W2359 := __e.Get(1)
_ = W2359
tmp5168 := MakeNative(func(__e *ControlFlow) {
W2360 := __e.Get(1)
_ = W2360
tmp5169 := MakeNative(func(__e *ControlFlow) {
W2361 := __e.Get(1)
_ = W2361
tmp5176 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2361)


if True == tmp5176 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5170 := MakeNative(func(__e *ControlFlow) {
W2362 := __e.Get(1)
_ = W2362
tmp5171 := MakeNative(func(__e *ControlFlow) {
W2363 := __e.Get(1)
_ = W2363
tmp5172 := PrimCons(W2359, W2362)

__e.TailApply(PrimFunc(symshen_4comb), W2363, tmp5172)
return


}, 1)

tmp5173 := Call(__e, PrimFunc(symshen_4in_1_6), W2361)


__e.TailApply(tmp5171, tmp5173)
return


}, 1)

tmp5174 := Call(__e, PrimFunc(symshen_4_5_1out), W2361)


__e.TailApply(tmp5170, tmp5174)
return


}


}, 1)

tmp5177 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2360)


__e.TailApply(tmp5169, tmp5177)
return


}, 1)

tmp5178 := Call(__e, PrimFunc(symshen_4in_1_6), W2358)


__e.TailApply(tmp5168, tmp5178)
return


}, 1)

tmp5179 := Call(__e, PrimFunc(symshen_4_5_1out), W2358)


__e.TailApply(tmp5167, tmp5179)
return


}


}, 1)

tmp5182 := Call(__e, PrimFunc(symshen_4_5atom_6), V2284)


tmp5183 := Call(__e, tmp5166, tmp5182)


__e.TailApply(tmp5135, tmp5183)
return


} else {
__e.Return(W2351)
return
}


}, 1)

tmp5186 := MakeNative(func(__e *ControlFlow) {
W2352 := __e.Get(1)
_ = W2352
tmp5198 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2352)


if True == tmp5198 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5187 := MakeNative(func(__e *ControlFlow) {
W2353 := __e.Get(1)
_ = W2353
tmp5188 := MakeNative(func(__e *ControlFlow) {
W2354 := __e.Get(1)
_ = W2354
tmp5194 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2354)


if True == tmp5194 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5189 := MakeNative(func(__e *ControlFlow) {
W2355 := __e.Get(1)
_ = W2355
tmp5190 := MakeNative(func(__e *ControlFlow) {
W2356 := __e.Get(1)
_ = W2356
__e.TailApply(PrimFunc(symshen_4comb), W2356, W2355)
return
}, 1)

tmp5191 := Call(__e, PrimFunc(symshen_4in_1_6), W2354)


__e.TailApply(tmp5190, tmp5191)
return


}, 1)

tmp5192 := Call(__e, PrimFunc(symshen_4_5_1out), W2354)


__e.TailApply(tmp5189, tmp5192)
return


}


}, 1)

tmp5195 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2353)


__e.TailApply(tmp5188, tmp5195)
return


}, 1)

tmp5196 := Call(__e, PrimFunc(symshen_4in_1_6), W2352)


__e.TailApply(tmp5187, tmp5196)
return


}


}, 1)

tmp5199 := Call(__e, PrimFunc(symshen_4_5comment_6), V2284)


tmp5200 := Call(__e, tmp5186, tmp5199)


__e.TailApply(tmp5134, tmp5200)
return


} else {
__e.Return(W2345)
return
}


}, 1)

tmp5203 := MakeNative(func(__e *ControlFlow) {
W2346 := __e.Get(1)
_ = W2346
tmp5217 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2346)


if True == tmp5217 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5204 := MakeNative(func(__e *ControlFlow) {
W2347 := __e.Get(1)
_ = W2347
tmp5205 := MakeNative(func(__e *ControlFlow) {
W2348 := __e.Get(1)
_ = W2348
tmp5213 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2348)


if True == tmp5213 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5206 := MakeNative(func(__e *ControlFlow) {
W2349 := __e.Get(1)
_ = W2349
tmp5207 := MakeNative(func(__e *ControlFlow) {
W2350 := __e.Get(1)
_ = W2350
tmp5208 := PrimIntern(MakeString(","))

tmp5209 := PrimCons(tmp5208, W2349)

__e.TailApply(PrimFunc(symshen_4comb), W2350, tmp5209)
return


}, 1)

tmp5210 := Call(__e, PrimFunc(symshen_4in_1_6), W2348)


__e.TailApply(tmp5207, tmp5210)
return


}, 1)

tmp5211 := Call(__e, PrimFunc(symshen_4_5_1out), W2348)


__e.TailApply(tmp5206, tmp5211)
return


}


}, 1)

tmp5214 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2347)


__e.TailApply(tmp5205, tmp5214)
return


}, 1)

tmp5215 := Call(__e, PrimFunc(symshen_4in_1_6), W2346)


__e.TailApply(tmp5204, tmp5215)
return


}


}, 1)

tmp5218 := Call(__e, PrimFunc(symshen_4_5comma_6), V2284)


tmp5219 := Call(__e, tmp5203, tmp5218)


__e.TailApply(tmp5133, tmp5219)
return


} else {
__e.Return(W2339)
return
}


}, 1)

tmp5222 := MakeNative(func(__e *ControlFlow) {
W2340 := __e.Get(1)
_ = W2340
tmp5236 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2340)


if True == tmp5236 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5223 := MakeNative(func(__e *ControlFlow) {
W2341 := __e.Get(1)
_ = W2341
tmp5224 := MakeNative(func(__e *ControlFlow) {
W2342 := __e.Get(1)
_ = W2342
tmp5232 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2342)


if True == tmp5232 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5225 := MakeNative(func(__e *ControlFlow) {
W2343 := __e.Get(1)
_ = W2343
tmp5226 := MakeNative(func(__e *ControlFlow) {
W2344 := __e.Get(1)
_ = W2344
tmp5227 := PrimIntern(MakeString(":"))

tmp5228 := PrimCons(tmp5227, W2343)

__e.TailApply(PrimFunc(symshen_4comb), W2344, tmp5228)
return


}, 1)

tmp5229 := Call(__e, PrimFunc(symshen_4in_1_6), W2342)


__e.TailApply(tmp5226, tmp5229)
return


}, 1)

tmp5230 := Call(__e, PrimFunc(symshen_4_5_1out), W2342)


__e.TailApply(tmp5225, tmp5230)
return


}


}, 1)

tmp5233 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2341)


__e.TailApply(tmp5224, tmp5233)
return


}, 1)

tmp5234 := Call(__e, PrimFunc(symshen_4in_1_6), W2340)


__e.TailApply(tmp5223, tmp5234)
return


}


}, 1)

tmp5237 := Call(__e, PrimFunc(symshen_4_5colon_6), V2284)


tmp5238 := Call(__e, tmp5222, tmp5237)


__e.TailApply(tmp5132, tmp5238)
return


} else {
__e.Return(W2331)
return
}


}, 1)

tmp5241 := MakeNative(func(__e *ControlFlow) {
W2332 := __e.Get(1)
_ = W2332
tmp5261 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2332)


if True == tmp5261 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5242 := MakeNative(func(__e *ControlFlow) {
W2333 := __e.Get(1)
_ = W2333
tmp5243 := MakeNative(func(__e *ControlFlow) {
W2334 := __e.Get(1)
_ = W2334
tmp5257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2334)


if True == tmp5257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5244 := MakeNative(func(__e *ControlFlow) {
W2335 := __e.Get(1)
_ = W2335
tmp5245 := MakeNative(func(__e *ControlFlow) {
W2336 := __e.Get(1)
_ = W2336
tmp5253 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2336)


if True == tmp5253 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5246 := MakeNative(func(__e *ControlFlow) {
W2337 := __e.Get(1)
_ = W2337
tmp5247 := MakeNative(func(__e *ControlFlow) {
W2338 := __e.Get(1)
_ = W2338
tmp5248 := PrimIntern(MakeString(":="))

tmp5249 := PrimCons(tmp5248, W2337)

__e.TailApply(PrimFunc(symshen_4comb), W2338, tmp5249)
return


}, 1)

tmp5250 := Call(__e, PrimFunc(symshen_4in_1_6), W2336)


__e.TailApply(tmp5247, tmp5250)
return


}, 1)

tmp5251 := Call(__e, PrimFunc(symshen_4_5_1out), W2336)


__e.TailApply(tmp5246, tmp5251)
return


}


}, 1)

tmp5254 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2335)


__e.TailApply(tmp5245, tmp5254)
return


}, 1)

tmp5255 := Call(__e, PrimFunc(symshen_4in_1_6), W2334)


__e.TailApply(tmp5244, tmp5255)
return


}


}, 1)

tmp5258 := Call(__e, PrimFunc(symshen_4_5equal_6), W2333)


__e.TailApply(tmp5243, tmp5258)
return


}, 1)

tmp5259 := Call(__e, PrimFunc(symshen_4in_1_6), W2332)


__e.TailApply(tmp5242, tmp5259)
return


}


}, 1)

tmp5262 := Call(__e, PrimFunc(symshen_4_5colon_6), V2284)


tmp5263 := Call(__e, tmp5241, tmp5262)


__e.TailApply(tmp5131, tmp5263)
return


} else {
__e.Return(W2325)
return
}


}, 1)

tmp5266 := MakeNative(func(__e *ControlFlow) {
W2326 := __e.Get(1)
_ = W2326
tmp5280 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2326)


if True == tmp5280 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5267 := MakeNative(func(__e *ControlFlow) {
W2327 := __e.Get(1)
_ = W2327
tmp5268 := MakeNative(func(__e *ControlFlow) {
W2328 := __e.Get(1)
_ = W2328
tmp5276 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2328)


if True == tmp5276 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5269 := MakeNative(func(__e *ControlFlow) {
W2329 := __e.Get(1)
_ = W2329
tmp5270 := MakeNative(func(__e *ControlFlow) {
W2330 := __e.Get(1)
_ = W2330
tmp5271 := PrimIntern(MakeString(";"))

tmp5272 := PrimCons(tmp5271, W2329)

__e.TailApply(PrimFunc(symshen_4comb), W2330, tmp5272)
return


}, 1)

tmp5273 := Call(__e, PrimFunc(symshen_4in_1_6), W2328)


__e.TailApply(tmp5270, tmp5273)
return


}, 1)

tmp5274 := Call(__e, PrimFunc(symshen_4_5_1out), W2328)


__e.TailApply(tmp5269, tmp5274)
return


}


}, 1)

tmp5277 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2327)


__e.TailApply(tmp5268, tmp5277)
return


}, 1)

tmp5278 := Call(__e, PrimFunc(symshen_4in_1_6), W2326)


__e.TailApply(tmp5267, tmp5278)
return


}


}, 1)

tmp5281 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2284)


tmp5282 := Call(__e, tmp5266, tmp5281)


__e.TailApply(tmp5130, tmp5282)
return


} else {
__e.Return(W2319)
return
}


}, 1)

tmp5285 := MakeNative(func(__e *ControlFlow) {
W2320 := __e.Get(1)
_ = W2320
tmp5298 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2320)


if True == tmp5298 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5286 := MakeNative(func(__e *ControlFlow) {
W2321 := __e.Get(1)
_ = W2321
tmp5287 := MakeNative(func(__e *ControlFlow) {
W2322 := __e.Get(1)
_ = W2322
tmp5294 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2322)


if True == tmp5294 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5288 := MakeNative(func(__e *ControlFlow) {
W2323 := __e.Get(1)
_ = W2323
tmp5289 := MakeNative(func(__e *ControlFlow) {
W2324 := __e.Get(1)
_ = W2324
tmp5290 := PrimCons(symbar_b, W2323)

__e.TailApply(PrimFunc(symshen_4comb), W2324, tmp5290)
return


}, 1)

tmp5291 := Call(__e, PrimFunc(symshen_4in_1_6), W2322)


__e.TailApply(tmp5289, tmp5291)
return


}, 1)

tmp5292 := Call(__e, PrimFunc(symshen_4_5_1out), W2322)


__e.TailApply(tmp5288, tmp5292)
return


}


}, 1)

tmp5295 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2321)


__e.TailApply(tmp5287, tmp5295)
return


}, 1)

tmp5296 := Call(__e, PrimFunc(symshen_4in_1_6), W2320)


__e.TailApply(tmp5286, tmp5296)
return


}


}, 1)

tmp5299 := Call(__e, PrimFunc(symshen_4_5bar_6), V2284)


tmp5300 := Call(__e, tmp5285, tmp5299)


__e.TailApply(tmp5129, tmp5300)
return


} else {
__e.Return(W2313)
return
}


}, 1)

tmp5303 := MakeNative(func(__e *ControlFlow) {
W2314 := __e.Get(1)
_ = W2314
tmp5316 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2314)


if True == tmp5316 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5304 := MakeNative(func(__e *ControlFlow) {
W2315 := __e.Get(1)
_ = W2315
tmp5305 := MakeNative(func(__e *ControlFlow) {
W2316 := __e.Get(1)
_ = W2316
tmp5312 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2316)


if True == tmp5312 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5306 := MakeNative(func(__e *ControlFlow) {
W2317 := __e.Get(1)
_ = W2317
tmp5307 := MakeNative(func(__e *ControlFlow) {
W2318 := __e.Get(1)
_ = W2318
tmp5308 := PrimCons(sym_j, W2317)

__e.TailApply(PrimFunc(symshen_4comb), W2318, tmp5308)
return


}, 1)

tmp5309 := Call(__e, PrimFunc(symshen_4in_1_6), W2316)


__e.TailApply(tmp5307, tmp5309)
return


}, 1)

tmp5310 := Call(__e, PrimFunc(symshen_4_5_1out), W2316)


__e.TailApply(tmp5306, tmp5310)
return


}


}, 1)

tmp5313 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2315)


__e.TailApply(tmp5305, tmp5313)
return


}, 1)

tmp5314 := Call(__e, PrimFunc(symshen_4in_1_6), W2314)


__e.TailApply(tmp5304, tmp5314)
return


}


}, 1)

tmp5317 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2284)


tmp5318 := Call(__e, tmp5303, tmp5317)


__e.TailApply(tmp5128, tmp5318)
return


} else {
__e.Return(W2307)
return
}


}, 1)

tmp5321 := MakeNative(func(__e *ControlFlow) {
W2308 := __e.Get(1)
_ = W2308
tmp5334 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2308)


if True == tmp5334 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5322 := MakeNative(func(__e *ControlFlow) {
W2309 := __e.Get(1)
_ = W2309
tmp5323 := MakeNative(func(__e *ControlFlow) {
W2310 := __e.Get(1)
_ = W2310
tmp5330 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2310)


if True == tmp5330 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5324 := MakeNative(func(__e *ControlFlow) {
W2311 := __e.Get(1)
_ = W2311
tmp5325 := MakeNative(func(__e *ControlFlow) {
W2312 := __e.Get(1)
_ = W2312
tmp5326 := PrimCons(sym_i, W2311)

__e.TailApply(PrimFunc(symshen_4comb), W2312, tmp5326)
return


}, 1)

tmp5327 := Call(__e, PrimFunc(symshen_4in_1_6), W2310)


__e.TailApply(tmp5325, tmp5327)
return


}, 1)

tmp5328 := Call(__e, PrimFunc(symshen_4_5_1out), W2310)


__e.TailApply(tmp5324, tmp5328)
return


}


}, 1)

tmp5331 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), W2309)


__e.TailApply(tmp5323, tmp5331)
return


}, 1)

tmp5332 := Call(__e, PrimFunc(symshen_4in_1_6), W2308)


__e.TailApply(tmp5322, tmp5332)
return


}


}, 1)

tmp5335 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2284)


tmp5336 := Call(__e, tmp5321, tmp5335)


__e.TailApply(tmp5127, tmp5336)
return


} else {
__e.Return(W2296)
return
}


}, 1)

tmp5339 := MakeNative(func(__e *ControlFlow) {
W2297 := __e.Get(1)
_ = W2297
tmp5366 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2297)


if True == tmp5366 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5340 := MakeNative(func(__e *ControlFlow) {
W2298 := __e.Get(1)
_ = W2298
tmp5341 := MakeNative(func(__e *ControlFlow) {
W2299 := __e.Get(1)
_ = W2299
tmp5362 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2299)


if True == tmp5362 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5342 := MakeNative(func(__e *ControlFlow) {
W2300 := __e.Get(1)
_ = W2300
tmp5343 := MakeNative(func(__e *ControlFlow) {
W2301 := __e.Get(1)
_ = W2301
tmp5344 := MakeNative(func(__e *ControlFlow) {
W2302 := __e.Get(1)
_ = W2302
tmp5357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2302)


if True == tmp5357 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5345 := MakeNative(func(__e *ControlFlow) {
W2303 := __e.Get(1)
_ = W2303
tmp5346 := MakeNative(func(__e *ControlFlow) {
W2304 := __e.Get(1)
_ = W2304
tmp5353 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2304)


if True == tmp5353 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5347 := MakeNative(func(__e *ControlFlow) {
W2305 := __e.Get(1)
_ = W2305
tmp5348 := MakeNative(func(__e *ControlFlow) {
W2306 := __e.Get(1)
_ = W2306
tmp5349 := Call(__e, PrimFunc(symshen_4add_1sexpr), W2300, W2305)


__e.TailApply(PrimFunc(symshen_4comb), W2306, tmp5349)
return


}, 1)

tmp5350 := Call(__e, PrimFunc(symshen_4in_1_6), W2304)


__e.TailApply(tmp5348, tmp5350)
return


}, 1)

tmp5351 := Call(__e, PrimFunc(symshen_4_5_1out), W2304)


__e.TailApply(tmp5347, tmp5351)
return


}


}, 1)

tmp5354 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2303)


__e.TailApply(tmp5346, tmp5354)
return


}, 1)

tmp5355 := Call(__e, PrimFunc(symshen_4in_1_6), W2302)


__e.TailApply(tmp5345, tmp5355)
return


}


}, 1)

tmp5358 := Call(__e, PrimFunc(symshen_4_5rrb_6), W2301)


__e.TailApply(tmp5344, tmp5358)
return


}, 1)

tmp5359 := Call(__e, PrimFunc(symshen_4in_1_6), W2299)


__e.TailApply(tmp5343, tmp5359)
return


}, 1)

tmp5360 := Call(__e, PrimFunc(symshen_4_5_1out), W2299)


__e.TailApply(tmp5342, tmp5360)
return


}


}, 1)

tmp5363 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2298)


__e.TailApply(tmp5341, tmp5363)
return


}, 1)

tmp5364 := Call(__e, PrimFunc(symshen_4in_1_6), W2297)


__e.TailApply(tmp5340, tmp5364)
return


}


}, 1)

tmp5367 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2284)


tmp5368 := Call(__e, tmp5339, tmp5367)


__e.TailApply(tmp5126, tmp5368)
return


} else {
__e.Return(W2285)
return
}


}, 1)

tmp5371 := MakeNative(func(__e *ControlFlow) {
W2286 := __e.Get(1)
_ = W2286
tmp5399 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2286)


if True == tmp5399 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5372 := MakeNative(func(__e *ControlFlow) {
W2287 := __e.Get(1)
_ = W2287
tmp5373 := MakeNative(func(__e *ControlFlow) {
W2288 := __e.Get(1)
_ = W2288
tmp5395 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2288)


if True == tmp5395 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5374 := MakeNative(func(__e *ControlFlow) {
W2289 := __e.Get(1)
_ = W2289
tmp5375 := MakeNative(func(__e *ControlFlow) {
W2290 := __e.Get(1)
_ = W2290
tmp5376 := MakeNative(func(__e *ControlFlow) {
W2291 := __e.Get(1)
_ = W2291
tmp5390 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2291)


if True == tmp5390 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5377 := MakeNative(func(__e *ControlFlow) {
W2292 := __e.Get(1)
_ = W2292
tmp5378 := MakeNative(func(__e *ControlFlow) {
W2293 := __e.Get(1)
_ = W2293
tmp5386 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2293)


if True == tmp5386 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5379 := MakeNative(func(__e *ControlFlow) {
W2294 := __e.Get(1)
_ = W2294
tmp5380 := MakeNative(func(__e *ControlFlow) {
W2295 := __e.Get(1)
_ = W2295
tmp5381 := Call(__e, PrimFunc(symshen_4cons_1form), W2289)


tmp5382 := PrimCons(tmp5381, W2294)

__e.TailApply(PrimFunc(symshen_4comb), W2295, tmp5382)
return


}, 1)

tmp5383 := Call(__e, PrimFunc(symshen_4in_1_6), W2293)


__e.TailApply(tmp5380, tmp5383)
return


}, 1)

tmp5384 := Call(__e, PrimFunc(symshen_4_5_1out), W2293)


__e.TailApply(tmp5379, tmp5384)
return


}


}, 1)

tmp5387 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), W2292)


__e.TailApply(tmp5378, tmp5387)
return


}, 1)

tmp5388 := Call(__e, PrimFunc(symshen_4in_1_6), W2291)


__e.TailApply(tmp5377, tmp5388)
return


}


}, 1)

tmp5391 := Call(__e, PrimFunc(symshen_4_5rsb_6), W2290)


__e.TailApply(tmp5376, tmp5391)
return


}, 1)

tmp5392 := Call(__e, PrimFunc(symshen_4in_1_6), W2288)


__e.TailApply(tmp5375, tmp5392)
return


}, 1)

tmp5393 := Call(__e, PrimFunc(symshen_4_5_1out), W2288)


__e.TailApply(tmp5374, tmp5393)
return


}


}, 1)

tmp5396 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), W2287)


__e.TailApply(tmp5373, tmp5396)
return


}, 1)

tmp5397 := Call(__e, PrimFunc(symshen_4in_1_6), W2286)


__e.TailApply(tmp5372, tmp5397)
return


}


}, 1)

tmp5400 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2284)


tmp5401 := Call(__e, tmp5371, tmp5400)


__e.TailApply(tmp5125, tmp5401)
return


}, 1)

tmp5402 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp5124)


_ = tmp5402

tmp5403 := MakeNative(func(__e *ControlFlow) {
V2373 := __e.Get(1)
_ = V2373
V2374 := __e.Get(2)
_ = V2374
tmp5421 := PrimIsPair(V2373)

var ifres5408 Obj

if True == tmp5421 {
tmp5419 := PrimHead(V2373)

tmp5420 := PrimEqual(sym_3, tmp5419)

var ifres5410 Obj

if True == tmp5420 {
tmp5417 := PrimTail(V2373)

tmp5418 := PrimIsPair(tmp5417)

var ifres5412 Obj

if True == tmp5418 {
tmp5414 := PrimTail(V2373)

tmp5415 := PrimTail(tmp5414)

tmp5416 := PrimEqual(Nil, tmp5415)

var ifres5413 Obj

if True == tmp5416 {
ifres5413 = True


} else {
ifres5413 = False


}

ifres5412 = ifres5413


} else {
ifres5412 = False


}

var ifres5411 Obj

if True == ifres5412 {
ifres5411 = True


} else {
ifres5411 = False


}

ifres5410 = ifres5411


} else {
ifres5410 = False


}

var ifres5409 Obj

if True == ifres5410 {
ifres5409 = True


} else {
ifres5409 = False


}

ifres5408 = ifres5409


} else {
ifres5408 = False


}

if True == ifres5408 {
tmp5404 := PrimTail(V2373)

tmp5405 := PrimHead(tmp5404)

tmp5406 := Call(__e, PrimFunc(symexplode), tmp5405)


__e.TailApply(PrimFunc(symappend), tmp5406, V2374)
return


} else {
__e.Return(PrimCons(V2373, V2374))
return
}


}, 2)

tmp5422 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp5403)


_ = tmp5422

tmp5423 := MakeNative(func(__e *ControlFlow) {
V2375 := __e.Get(1)
_ = V2375
tmp5424 := MakeNative(func(__e *ControlFlow) {
W2376 := __e.Get(1)
_ = W2376
tmp5426 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2376)


if True == tmp5426 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2376)
return
}


}, 1)

tmp5432 := Call(__e, PrimFunc(symshen_4hds_a_2), V2375, MakeNumber(91))


var ifres5427 Obj

if True == tmp5432 {
tmp5428 := MakeNative(func(__e *ControlFlow) {
W2377 := __e.Get(1)
_ = W2377
__e.TailApply(PrimFunc(symshen_4comb), W2377, symshen_4skip)
return
}, 1)

tmp5429 := Call(__e, PrimFunc(symtail), V2375)


tmp5430 := Call(__e, tmp5428, tmp5429)


ifres5427 = tmp5430


} else {
tmp5431 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5427 = tmp5431


}

__e.TailApply(tmp5424, ifres5427)
return


}, 1)

tmp5433 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp5423)


_ = tmp5433

tmp5434 := MakeNative(func(__e *ControlFlow) {
V2378 := __e.Get(1)
_ = V2378
tmp5435 := MakeNative(func(__e *ControlFlow) {
W2379 := __e.Get(1)
_ = W2379
tmp5437 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2379)


if True == tmp5437 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2379)
return
}


}, 1)

tmp5443 := Call(__e, PrimFunc(symshen_4hds_a_2), V2378, MakeNumber(93))


var ifres5438 Obj

if True == tmp5443 {
tmp5439 := MakeNative(func(__e *ControlFlow) {
W2380 := __e.Get(1)
_ = W2380
__e.TailApply(PrimFunc(symshen_4comb), W2380, symshen_4skip)
return
}, 1)

tmp5440 := Call(__e, PrimFunc(symtail), V2378)


tmp5441 := Call(__e, tmp5439, tmp5440)


ifres5438 = tmp5441


} else {
tmp5442 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5438 = tmp5442


}

__e.TailApply(tmp5435, ifres5438)
return


}, 1)

tmp5444 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp5434)


_ = tmp5444

tmp5445 := MakeNative(func(__e *ControlFlow) {
V2381 := __e.Get(1)
_ = V2381
tmp5446 := MakeNative(func(__e *ControlFlow) {
W2382 := __e.Get(1)
_ = W2382
tmp5448 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2382)


if True == tmp5448 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2382)
return
}


}, 1)

tmp5449 := MakeNative(func(__e *ControlFlow) {
W2383 := __e.Get(1)
_ = W2383
tmp5455 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2383)


if True == tmp5455 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5450 := MakeNative(func(__e *ControlFlow) {
W2384 := __e.Get(1)
_ = W2384
tmp5451 := MakeNative(func(__e *ControlFlow) {
W2385 := __e.Get(1)
_ = W2385
__e.TailApply(PrimFunc(symshen_4comb), W2385, W2384)
return
}, 1)

tmp5452 := Call(__e, PrimFunc(symshen_4in_1_6), W2383)


__e.TailApply(tmp5451, tmp5452)
return


}, 1)

tmp5453 := Call(__e, PrimFunc(symshen_4_5_1out), W2383)


__e.TailApply(tmp5450, tmp5453)
return


}


}, 1)

tmp5456 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2381)


tmp5457 := Call(__e, tmp5449, tmp5456)


__e.TailApply(tmp5446, tmp5457)
return


}, 1)

tmp5458 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5445)


_ = tmp5458

tmp5459 := MakeNative(func(__e *ControlFlow) {
V2386 := __e.Get(1)
_ = V2386
tmp5460 := MakeNative(func(__e *ControlFlow) {
W2387 := __e.Get(1)
_ = W2387
tmp5462 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2387)


if True == tmp5462 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2387)
return
}


}, 1)

tmp5463 := MakeNative(func(__e *ControlFlow) {
W2388 := __e.Get(1)
_ = W2388
tmp5469 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2388)


if True == tmp5469 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5464 := MakeNative(func(__e *ControlFlow) {
W2389 := __e.Get(1)
_ = W2389
tmp5465 := MakeNative(func(__e *ControlFlow) {
W2390 := __e.Get(1)
_ = W2390
__e.TailApply(PrimFunc(symshen_4comb), W2390, W2389)
return
}, 1)

tmp5466 := Call(__e, PrimFunc(symshen_4in_1_6), W2388)


__e.TailApply(tmp5465, tmp5466)
return


}, 1)

tmp5467 := Call(__e, PrimFunc(symshen_4_5_1out), W2388)


__e.TailApply(tmp5464, tmp5467)
return


}


}, 1)

tmp5470 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2386)


tmp5471 := Call(__e, tmp5463, tmp5470)


__e.TailApply(tmp5460, tmp5471)
return


}, 1)

tmp5472 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5459)


_ = tmp5472

tmp5473 := MakeNative(func(__e *ControlFlow) {
V2392 := __e.Get(1)
_ = V2392
tmp5530 := PrimEqual(Nil, V2392)

if True == tmp5530 {
__e.Return(Nil)
return
} else {
tmp5528 := PrimIsPair(V2392)

var ifres5508 Obj

if True == tmp5528 {
tmp5526 := PrimTail(V2392)

tmp5527 := PrimIsPair(tmp5526)

var ifres5510 Obj

if True == tmp5527 {
tmp5523 := PrimTail(V2392)

tmp5524 := PrimTail(tmp5523)

tmp5525 := PrimIsPair(tmp5524)

var ifres5512 Obj

if True == tmp5525 {
tmp5519 := PrimTail(V2392)

tmp5520 := PrimTail(tmp5519)

tmp5521 := PrimTail(tmp5520)

tmp5522 := PrimEqual(Nil, tmp5521)

var ifres5514 Obj

if True == tmp5522 {
tmp5516 := PrimTail(V2392)

tmp5517 := PrimHead(tmp5516)

tmp5518 := PrimEqual(tmp5517, symbar_b)

var ifres5515 Obj

if True == tmp5518 {
ifres5515 = True


} else {
ifres5515 = False


}

ifres5514 = ifres5515


} else {
ifres5514 = False


}

var ifres5513 Obj

if True == ifres5514 {
ifres5513 = True


} else {
ifres5513 = False


}

ifres5512 = ifres5513


} else {
ifres5512 = False


}

var ifres5511 Obj

if True == ifres5512 {
ifres5511 = True


} else {
ifres5511 = False


}

ifres5510 = ifres5511


} else {
ifres5510 = False


}

var ifres5509 Obj

if True == ifres5510 {
ifres5509 = True


} else {
ifres5509 = False


}

ifres5508 = ifres5509


} else {
ifres5508 = False


}

if True == ifres5508 {
tmp5474 := PrimHead(V2392)

tmp5475 := PrimTail(V2392)

tmp5476 := PrimTail(tmp5475)

tmp5477 := PrimCons(tmp5474, tmp5476)

__e.Return(PrimCons(symcons, tmp5477))
return


} else {
tmp5506 := PrimIsPair(V2392)

var ifres5486 Obj

if True == tmp5506 {
tmp5504 := PrimTail(V2392)

tmp5505 := PrimIsPair(tmp5504)

var ifres5488 Obj

if True == tmp5505 {
tmp5501 := PrimTail(V2392)

tmp5502 := PrimTail(tmp5501)

tmp5503 := PrimIsPair(tmp5502)

var ifres5490 Obj

if True == tmp5503 {
tmp5497 := PrimTail(V2392)

tmp5498 := PrimTail(tmp5497)

tmp5499 := PrimTail(tmp5498)

tmp5500 := PrimIsPair(tmp5499)

var ifres5492 Obj

if True == tmp5500 {
tmp5494 := PrimTail(V2392)

tmp5495 := PrimHead(tmp5494)

tmp5496 := PrimEqual(tmp5495, symbar_b)

var ifres5493 Obj

if True == tmp5496 {
ifres5493 = True


} else {
ifres5493 = False


}

ifres5492 = ifres5493


} else {
ifres5492 = False


}

var ifres5491 Obj

if True == ifres5492 {
ifres5491 = True


} else {
ifres5491 = False


}

ifres5490 = ifres5491


} else {
ifres5490 = False


}

var ifres5489 Obj

if True == ifres5490 {
ifres5489 = True


} else {
ifres5489 = False


}

ifres5488 = ifres5489


} else {
ifres5488 = False


}

var ifres5487 Obj

if True == ifres5488 {
ifres5487 = True


} else {
ifres5487 = False


}

ifres5486 = ifres5487


} else {
ifres5486 = False


}

if True == ifres5486 {
__e.Return(PrimSimpleError(MakeString("misapplication of |\n")))
return
} else {
tmp5484 := PrimIsPair(V2392)

if True == tmp5484 {
tmp5478 := PrimHead(V2392)

tmp5479 := PrimTail(V2392)

tmp5480 := Call(__e, PrimFunc(symshen_4cons_1form), tmp5479)


tmp5481 := PrimCons(tmp5480, Nil)

tmp5482 := PrimCons(tmp5478, tmp5481)

__e.Return(PrimCons(symcons, tmp5482))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.cons-form")))
return
}


}


}


}


}, 1)

tmp5531 := Call(__e, ns2_1set, symshen_4cons_1form, tmp5473)


_ = tmp5531

tmp5532 := MakeNative(func(__e *ControlFlow) {
V2393 := __e.Get(1)
_ = V2393
tmp5533 := MakeNative(func(__e *ControlFlow) {
W2394 := __e.Get(1)
_ = W2394
tmp5535 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2394)


if True == tmp5535 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2394)
return
}


}, 1)

tmp5541 := Call(__e, PrimFunc(symshen_4hds_a_2), V2393, MakeNumber(40))


var ifres5536 Obj

if True == tmp5541 {
tmp5537 := MakeNative(func(__e *ControlFlow) {
W2395 := __e.Get(1)
_ = W2395
__e.TailApply(PrimFunc(symshen_4comb), W2395, symshen_4skip)
return
}, 1)

tmp5538 := Call(__e, PrimFunc(symtail), V2393)


tmp5539 := Call(__e, tmp5537, tmp5538)


ifres5536 = tmp5539


} else {
tmp5540 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5536 = tmp5540


}

__e.TailApply(tmp5533, ifres5536)
return


}, 1)

tmp5542 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp5532)


_ = tmp5542

tmp5543 := MakeNative(func(__e *ControlFlow) {
V2396 := __e.Get(1)
_ = V2396
tmp5544 := MakeNative(func(__e *ControlFlow) {
W2397 := __e.Get(1)
_ = W2397
tmp5546 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2397)


if True == tmp5546 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2397)
return
}


}, 1)

tmp5552 := Call(__e, PrimFunc(symshen_4hds_a_2), V2396, MakeNumber(41))


var ifres5547 Obj

if True == tmp5552 {
tmp5548 := MakeNative(func(__e *ControlFlow) {
W2398 := __e.Get(1)
_ = W2398
__e.TailApply(PrimFunc(symshen_4comb), W2398, symshen_4skip)
return
}, 1)

tmp5549 := Call(__e, PrimFunc(symtail), V2396)


tmp5550 := Call(__e, tmp5548, tmp5549)


ifres5547 = tmp5550


} else {
tmp5551 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5547 = tmp5551


}

__e.TailApply(tmp5544, ifres5547)
return


}, 1)

tmp5553 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp5543)


_ = tmp5553

tmp5554 := MakeNative(func(__e *ControlFlow) {
V2399 := __e.Get(1)
_ = V2399
tmp5555 := MakeNative(func(__e *ControlFlow) {
W2400 := __e.Get(1)
_ = W2400
tmp5557 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2400)


if True == tmp5557 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2400)
return
}


}, 1)

tmp5563 := Call(__e, PrimFunc(symshen_4hds_a_2), V2399, MakeNumber(123))


var ifres5558 Obj

if True == tmp5563 {
tmp5559 := MakeNative(func(__e *ControlFlow) {
W2401 := __e.Get(1)
_ = W2401
__e.TailApply(PrimFunc(symshen_4comb), W2401, symshen_4skip)
return
}, 1)

tmp5560 := Call(__e, PrimFunc(symtail), V2399)


tmp5561 := Call(__e, tmp5559, tmp5560)


ifres5558 = tmp5561


} else {
tmp5562 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5558 = tmp5562


}

__e.TailApply(tmp5555, ifres5558)
return


}, 1)

tmp5564 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp5554)


_ = tmp5564

tmp5565 := MakeNative(func(__e *ControlFlow) {
V2402 := __e.Get(1)
_ = V2402
tmp5566 := MakeNative(func(__e *ControlFlow) {
W2403 := __e.Get(1)
_ = W2403
tmp5568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2403)


if True == tmp5568 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2403)
return
}


}, 1)

tmp5574 := Call(__e, PrimFunc(symshen_4hds_a_2), V2402, MakeNumber(125))


var ifres5569 Obj

if True == tmp5574 {
tmp5570 := MakeNative(func(__e *ControlFlow) {
W2404 := __e.Get(1)
_ = W2404
__e.TailApply(PrimFunc(symshen_4comb), W2404, symshen_4skip)
return
}, 1)

tmp5571 := Call(__e, PrimFunc(symtail), V2402)


tmp5572 := Call(__e, tmp5570, tmp5571)


ifres5569 = tmp5572


} else {
tmp5573 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5569 = tmp5573


}

__e.TailApply(tmp5566, ifres5569)
return


}, 1)

tmp5575 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp5565)


_ = tmp5575

tmp5576 := MakeNative(func(__e *ControlFlow) {
V2405 := __e.Get(1)
_ = V2405
tmp5577 := MakeNative(func(__e *ControlFlow) {
W2406 := __e.Get(1)
_ = W2406
tmp5579 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2406)


if True == tmp5579 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2406)
return
}


}, 1)

tmp5585 := Call(__e, PrimFunc(symshen_4hds_a_2), V2405, MakeNumber(124))


var ifres5580 Obj

if True == tmp5585 {
tmp5581 := MakeNative(func(__e *ControlFlow) {
W2407 := __e.Get(1)
_ = W2407
__e.TailApply(PrimFunc(symshen_4comb), W2407, symshen_4skip)
return
}, 1)

tmp5582 := Call(__e, PrimFunc(symtail), V2405)


tmp5583 := Call(__e, tmp5581, tmp5582)


ifres5580 = tmp5583


} else {
tmp5584 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5580 = tmp5584


}

__e.TailApply(tmp5577, ifres5580)
return


}, 1)

tmp5586 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp5576)


_ = tmp5586

tmp5587 := MakeNative(func(__e *ControlFlow) {
V2408 := __e.Get(1)
_ = V2408
tmp5588 := MakeNative(func(__e *ControlFlow) {
W2409 := __e.Get(1)
_ = W2409
tmp5590 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2409)


if True == tmp5590 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2409)
return
}


}, 1)

tmp5596 := Call(__e, PrimFunc(symshen_4hds_a_2), V2408, MakeNumber(59))


var ifres5591 Obj

if True == tmp5596 {
tmp5592 := MakeNative(func(__e *ControlFlow) {
W2410 := __e.Get(1)
_ = W2410
__e.TailApply(PrimFunc(symshen_4comb), W2410, symshen_4skip)
return
}, 1)

tmp5593 := Call(__e, PrimFunc(symtail), V2408)


tmp5594 := Call(__e, tmp5592, tmp5593)


ifres5591 = tmp5594


} else {
tmp5595 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5591 = tmp5595


}

__e.TailApply(tmp5588, ifres5591)
return


}, 1)

tmp5597 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp5587)


_ = tmp5597

tmp5598 := MakeNative(func(__e *ControlFlow) {
V2411 := __e.Get(1)
_ = V2411
tmp5599 := MakeNative(func(__e *ControlFlow) {
W2412 := __e.Get(1)
_ = W2412
tmp5601 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2412)


if True == tmp5601 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2412)
return
}


}, 1)

tmp5607 := Call(__e, PrimFunc(symshen_4hds_a_2), V2411, MakeNumber(58))


var ifres5602 Obj

if True == tmp5607 {
tmp5603 := MakeNative(func(__e *ControlFlow) {
W2413 := __e.Get(1)
_ = W2413
__e.TailApply(PrimFunc(symshen_4comb), W2413, symshen_4skip)
return
}, 1)

tmp5604 := Call(__e, PrimFunc(symtail), V2411)


tmp5605 := Call(__e, tmp5603, tmp5604)


ifres5602 = tmp5605


} else {
tmp5606 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5602 = tmp5606


}

__e.TailApply(tmp5599, ifres5602)
return


}, 1)

tmp5608 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp5598)


_ = tmp5608

tmp5609 := MakeNative(func(__e *ControlFlow) {
V2414 := __e.Get(1)
_ = V2414
tmp5610 := MakeNative(func(__e *ControlFlow) {
W2415 := __e.Get(1)
_ = W2415
tmp5612 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2415)


if True == tmp5612 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2415)
return
}


}, 1)

tmp5618 := Call(__e, PrimFunc(symshen_4hds_a_2), V2414, MakeNumber(44))


var ifres5613 Obj

if True == tmp5618 {
tmp5614 := MakeNative(func(__e *ControlFlow) {
W2416 := __e.Get(1)
_ = W2416
__e.TailApply(PrimFunc(symshen_4comb), W2416, symshen_4skip)
return
}, 1)

tmp5615 := Call(__e, PrimFunc(symtail), V2414)


tmp5616 := Call(__e, tmp5614, tmp5615)


ifres5613 = tmp5616


} else {
tmp5617 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5613 = tmp5617


}

__e.TailApply(tmp5610, ifres5613)
return


}, 1)

tmp5619 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp5609)


_ = tmp5619

tmp5620 := MakeNative(func(__e *ControlFlow) {
V2417 := __e.Get(1)
_ = V2417
tmp5621 := MakeNative(func(__e *ControlFlow) {
W2418 := __e.Get(1)
_ = W2418
tmp5623 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2418)


if True == tmp5623 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2418)
return
}


}, 1)

tmp5629 := Call(__e, PrimFunc(symshen_4hds_a_2), V2417, MakeNumber(61))


var ifres5624 Obj

if True == tmp5629 {
tmp5625 := MakeNative(func(__e *ControlFlow) {
W2419 := __e.Get(1)
_ = W2419
__e.TailApply(PrimFunc(symshen_4comb), W2419, symshen_4skip)
return
}, 1)

tmp5626 := Call(__e, PrimFunc(symtail), V2417)


tmp5627 := Call(__e, tmp5625, tmp5626)


ifres5624 = tmp5627


} else {
tmp5628 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5624 = tmp5628


}

__e.TailApply(tmp5621, ifres5624)
return


}, 1)

tmp5630 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp5620)


_ = tmp5630

tmp5631 := MakeNative(func(__e *ControlFlow) {
V2420 := __e.Get(1)
_ = V2420
tmp5632 := MakeNative(func(__e *ControlFlow) {
W2421 := __e.Get(1)
_ = W2421
tmp5644 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2421)


if True == tmp5644 {
tmp5633 := MakeNative(func(__e *ControlFlow) {
W2424 := __e.Get(1)
_ = W2424
tmp5635 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2424)


if True == tmp5635 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2424)
return
}


}, 1)

tmp5636 := MakeNative(func(__e *ControlFlow) {
W2425 := __e.Get(1)
_ = W2425
tmp5640 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2425)


if True == tmp5640 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5637 := MakeNative(func(__e *ControlFlow) {
W2426 := __e.Get(1)
_ = W2426
__e.TailApply(PrimFunc(symshen_4comb), W2426, symshen_4skip)
return
}, 1)

tmp5638 := Call(__e, PrimFunc(symshen_4in_1_6), W2425)


__e.TailApply(tmp5637, tmp5638)
return


}


}, 1)

tmp5641 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2420)


tmp5642 := Call(__e, tmp5636, tmp5641)


__e.TailApply(tmp5633, tmp5642)
return


} else {
__e.Return(W2421)
return
}


}, 1)

tmp5645 := MakeNative(func(__e *ControlFlow) {
W2422 := __e.Get(1)
_ = W2422
tmp5649 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2422)


if True == tmp5649 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5646 := MakeNative(func(__e *ControlFlow) {
W2423 := __e.Get(1)
_ = W2423
__e.TailApply(PrimFunc(symshen_4comb), W2423, symshen_4skip)
return
}, 1)

tmp5647 := Call(__e, PrimFunc(symshen_4in_1_6), W2422)


__e.TailApply(tmp5646, tmp5647)
return


}


}, 1)

tmp5650 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2420)


tmp5651 := Call(__e, tmp5645, tmp5650)


__e.TailApply(tmp5632, tmp5651)
return


}, 1)

tmp5652 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp5631)


_ = tmp5652

tmp5653 := MakeNative(func(__e *ControlFlow) {
V2427 := __e.Get(1)
_ = V2427
tmp5654 := MakeNative(func(__e *ControlFlow) {
W2428 := __e.Get(1)
_ = W2428
tmp5656 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2428)


if True == tmp5656 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2428)
return
}


}, 1)

tmp5657 := MakeNative(func(__e *ControlFlow) {
W2429 := __e.Get(1)
_ = W2429
tmp5679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2429)


if True == tmp5679 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5658 := MakeNative(func(__e *ControlFlow) {
W2430 := __e.Get(1)
_ = W2430
tmp5659 := MakeNative(func(__e *ControlFlow) {
W2431 := __e.Get(1)
_ = W2431
tmp5675 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2431)


if True == tmp5675 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5660 := MakeNative(func(__e *ControlFlow) {
W2432 := __e.Get(1)
_ = W2432
tmp5661 := MakeNative(func(__e *ControlFlow) {
W2433 := __e.Get(1)
_ = W2433
tmp5671 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2433)


if True == tmp5671 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5662 := MakeNative(func(__e *ControlFlow) {
W2434 := __e.Get(1)
_ = W2434
tmp5663 := MakeNative(func(__e *ControlFlow) {
W2435 := __e.Get(1)
_ = W2435
tmp5667 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2435)


if True == tmp5667 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5664 := MakeNative(func(__e *ControlFlow) {
W2436 := __e.Get(1)
_ = W2436
__e.TailApply(PrimFunc(symshen_4comb), W2436, symshen_4skip)
return
}, 1)

tmp5665 := Call(__e, PrimFunc(symshen_4in_1_6), W2435)


__e.TailApply(tmp5664, tmp5665)
return


}


}, 1)

tmp5668 := Call(__e, PrimFunc(symshen_4_5returns_6), W2434)


__e.TailApply(tmp5663, tmp5668)
return


}, 1)

tmp5669 := Call(__e, PrimFunc(symshen_4in_1_6), W2433)


__e.TailApply(tmp5662, tmp5669)
return


}


}, 1)

tmp5672 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2432)


__e.TailApply(tmp5661, tmp5672)
return


}, 1)

tmp5673 := Call(__e, PrimFunc(symshen_4in_1_6), W2431)


__e.TailApply(tmp5660, tmp5673)
return


}


}, 1)

tmp5676 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2430)


__e.TailApply(tmp5659, tmp5676)
return


}, 1)

tmp5677 := Call(__e, PrimFunc(symshen_4in_1_6), W2429)


__e.TailApply(tmp5658, tmp5677)
return


}


}, 1)

tmp5680 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2427)


tmp5681 := Call(__e, tmp5657, tmp5680)


__e.TailApply(tmp5654, tmp5681)
return


}, 1)

tmp5682 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp5653)


_ = tmp5682

tmp5683 := MakeNative(func(__e *ControlFlow) {
V2437 := __e.Get(1)
_ = V2437
tmp5684 := MakeNative(func(__e *ControlFlow) {
W2438 := __e.Get(1)
_ = W2438
tmp5686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2438)


if True == tmp5686 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2438)
return
}


}, 1)

tmp5692 := Call(__e, PrimFunc(symshen_4hds_a_2), V2437, MakeNumber(92))


var ifres5687 Obj

if True == tmp5692 {
tmp5688 := MakeNative(func(__e *ControlFlow) {
W2439 := __e.Get(1)
_ = W2439
__e.TailApply(PrimFunc(symshen_4comb), W2439, symshen_4skip)
return
}, 1)

tmp5689 := Call(__e, PrimFunc(symtail), V2437)


tmp5690 := Call(__e, tmp5688, tmp5689)


ifres5687 = tmp5690


} else {
tmp5691 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5687 = tmp5691


}

__e.TailApply(tmp5684, ifres5687)
return


}, 1)

tmp5693 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp5683)


_ = tmp5693

tmp5694 := MakeNative(func(__e *ControlFlow) {
V2440 := __e.Get(1)
_ = V2440
tmp5695 := MakeNative(func(__e *ControlFlow) {
W2441 := __e.Get(1)
_ = W2441
tmp5707 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2441)


if True == tmp5707 {
tmp5696 := MakeNative(func(__e *ControlFlow) {
W2446 := __e.Get(1)
_ = W2446
tmp5698 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2446)


if True == tmp5698 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2446)
return
}


}, 1)

tmp5699 := MakeNative(func(__e *ControlFlow) {
W2447 := __e.Get(1)
_ = W2447
tmp5703 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2447)


if True == tmp5703 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5700 := MakeNative(func(__e *ControlFlow) {
W2448 := __e.Get(1)
_ = W2448
__e.TailApply(PrimFunc(symshen_4comb), W2448, symshen_4skip)
return
}, 1)

tmp5701 := Call(__e, PrimFunc(symshen_4in_1_6), W2447)


__e.TailApply(tmp5700, tmp5701)
return


}


}, 1)

tmp5704 := Call(__e, PrimFunc(sym_5e_6), V2440)


tmp5705 := Call(__e, tmp5699, tmp5704)


__e.TailApply(tmp5696, tmp5705)
return


} else {
__e.Return(W2441)
return
}


}, 1)

tmp5708 := MakeNative(func(__e *ControlFlow) {
W2442 := __e.Get(1)
_ = W2442
tmp5718 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2442)


if True == tmp5718 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5709 := MakeNative(func(__e *ControlFlow) {
W2443 := __e.Get(1)
_ = W2443
tmp5710 := MakeNative(func(__e *ControlFlow) {
W2444 := __e.Get(1)
_ = W2444
tmp5714 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2444)


if True == tmp5714 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5711 := MakeNative(func(__e *ControlFlow) {
W2445 := __e.Get(1)
_ = W2445
__e.TailApply(PrimFunc(symshen_4comb), W2445, symshen_4skip)
return
}, 1)

tmp5712 := Call(__e, PrimFunc(symshen_4in_1_6), W2444)


__e.TailApply(tmp5711, tmp5712)
return


}


}, 1)

tmp5715 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), W2443)


__e.TailApply(tmp5710, tmp5715)
return


}, 1)

tmp5716 := Call(__e, PrimFunc(symshen_4in_1_6), W2442)


__e.TailApply(tmp5709, tmp5716)
return


}


}, 1)

tmp5719 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2440)


tmp5720 := Call(__e, tmp5708, tmp5719)


__e.TailApply(tmp5695, tmp5720)
return


}, 1)

tmp5721 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp5694)


_ = tmp5721

tmp5722 := MakeNative(func(__e *ControlFlow) {
V2449 := __e.Get(1)
_ = V2449
tmp5723 := MakeNative(func(__e *ControlFlow) {
W2450 := __e.Get(1)
_ = W2450
tmp5725 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2450)


if True == tmp5725 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2450)
return
}


}, 1)

tmp5736 := PrimIsPair(V2449)

var ifres5726 Obj

if True == tmp5736 {
tmp5727 := MakeNative(func(__e *ControlFlow) {
W2451 := __e.Get(1)
_ = W2451
tmp5728 := MakeNative(func(__e *ControlFlow) {
W2452 := __e.Get(1)
_ = W2452
tmp5730 := Call(__e, PrimFunc(symshen_4return_2), W2451)


tmp5731 := PrimNot(tmp5730)

if True == tmp5731 {
__e.TailApply(PrimFunc(symshen_4comb), W2452, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5732 := Call(__e, PrimFunc(symtail), V2449)


__e.TailApply(tmp5728, tmp5732)
return


}, 1)

tmp5733 := Call(__e, PrimFunc(symhead), V2449)


tmp5734 := Call(__e, tmp5727, tmp5733)


ifres5726 = tmp5734


} else {
tmp5735 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5726 = tmp5735


}

__e.TailApply(tmp5723, ifres5726)
return


}, 1)

tmp5737 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp5722)


_ = tmp5737

tmp5738 := MakeNative(func(__e *ControlFlow) {
V2453 := __e.Get(1)
_ = V2453
tmp5739 := MakeNative(func(__e *ControlFlow) {
W2454 := __e.Get(1)
_ = W2454
tmp5751 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2454)


if True == tmp5751 {
tmp5740 := MakeNative(func(__e *ControlFlow) {
W2459 := __e.Get(1)
_ = W2459
tmp5742 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2459)


if True == tmp5742 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2459)
return
}


}, 1)

tmp5743 := MakeNative(func(__e *ControlFlow) {
W2460 := __e.Get(1)
_ = W2460
tmp5747 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2460)


if True == tmp5747 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5744 := MakeNative(func(__e *ControlFlow) {
W2461 := __e.Get(1)
_ = W2461
__e.TailApply(PrimFunc(symshen_4comb), W2461, symshen_4skip)
return
}, 1)

tmp5745 := Call(__e, PrimFunc(symshen_4in_1_6), W2460)


__e.TailApply(tmp5744, tmp5745)
return


}


}, 1)

tmp5748 := Call(__e, PrimFunc(symshen_4_5return_6), V2453)


tmp5749 := Call(__e, tmp5743, tmp5748)


__e.TailApply(tmp5740, tmp5749)
return


} else {
__e.Return(W2454)
return
}


}, 1)

tmp5752 := MakeNative(func(__e *ControlFlow) {
W2455 := __e.Get(1)
_ = W2455
tmp5762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2455)


if True == tmp5762 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5753 := MakeNative(func(__e *ControlFlow) {
W2456 := __e.Get(1)
_ = W2456
tmp5754 := MakeNative(func(__e *ControlFlow) {
W2457 := __e.Get(1)
_ = W2457
tmp5758 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2457)


if True == tmp5758 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5755 := MakeNative(func(__e *ControlFlow) {
W2458 := __e.Get(1)
_ = W2458
__e.TailApply(PrimFunc(symshen_4comb), W2458, symshen_4skip)
return
}, 1)

tmp5756 := Call(__e, PrimFunc(symshen_4in_1_6), W2457)


__e.TailApply(tmp5755, tmp5756)
return


}


}, 1)

tmp5759 := Call(__e, PrimFunc(symshen_4_5returns_6), W2456)


__e.TailApply(tmp5754, tmp5759)
return


}, 1)

tmp5760 := Call(__e, PrimFunc(symshen_4in_1_6), W2455)


__e.TailApply(tmp5753, tmp5760)
return


}


}, 1)

tmp5763 := Call(__e, PrimFunc(symshen_4_5return_6), V2453)


tmp5764 := Call(__e, tmp5752, tmp5763)


__e.TailApply(tmp5739, tmp5764)
return


}, 1)

tmp5765 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp5738)


_ = tmp5765

tmp5766 := MakeNative(func(__e *ControlFlow) {
V2462 := __e.Get(1)
_ = V2462
tmp5767 := MakeNative(func(__e *ControlFlow) {
W2463 := __e.Get(1)
_ = W2463
tmp5769 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2463)


if True == tmp5769 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2463)
return
}


}, 1)

tmp5779 := PrimIsPair(V2462)

var ifres5770 Obj

if True == tmp5779 {
tmp5771 := MakeNative(func(__e *ControlFlow) {
W2464 := __e.Get(1)
_ = W2464
tmp5772 := MakeNative(func(__e *ControlFlow) {
W2465 := __e.Get(1)
_ = W2465
tmp5774 := Call(__e, PrimFunc(symshen_4return_2), W2464)


if True == tmp5774 {
__e.TailApply(PrimFunc(symshen_4comb), W2465, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5775 := Call(__e, PrimFunc(symtail), V2462)


__e.TailApply(tmp5772, tmp5775)
return


}, 1)

tmp5776 := Call(__e, PrimFunc(symhead), V2462)


tmp5777 := Call(__e, tmp5771, tmp5776)


ifres5770 = tmp5777


} else {
tmp5778 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5770 = tmp5778


}

__e.TailApply(tmp5767, ifres5770)
return


}, 1)

tmp5780 := Call(__e, ns2_1set, symshen_4_5return_6, tmp5766)


_ = tmp5780

tmp5781 := MakeNative(func(__e *ControlFlow) {
V2466 := __e.Get(1)
_ = V2466
tmp5782 := PrimCons(MakeNumber(13), Nil)

tmp5783 := PrimCons(MakeNumber(10), tmp5782)

tmp5784 := PrimCons(MakeNumber(9), tmp5783)

__e.TailApply(PrimFunc(symelement_2), V2466, tmp5784)
return


}, 1)

tmp5785 := Call(__e, ns2_1set, symshen_4return_2, tmp5781)


_ = tmp5785

tmp5786 := MakeNative(func(__e *ControlFlow) {
V2467 := __e.Get(1)
_ = V2467
tmp5787 := MakeNative(func(__e *ControlFlow) {
W2468 := __e.Get(1)
_ = W2468
tmp5789 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2468)


if True == tmp5789 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2468)
return
}


}, 1)

tmp5790 := MakeNative(func(__e *ControlFlow) {
W2469 := __e.Get(1)
_ = W2469
tmp5806 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2469)


if True == tmp5806 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5791 := MakeNative(func(__e *ControlFlow) {
W2470 := __e.Get(1)
_ = W2470
tmp5792 := MakeNative(func(__e *ControlFlow) {
W2471 := __e.Get(1)
_ = W2471
tmp5802 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2471)


if True == tmp5802 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5793 := MakeNative(func(__e *ControlFlow) {
W2472 := __e.Get(1)
_ = W2472
tmp5794 := MakeNative(func(__e *ControlFlow) {
W2473 := __e.Get(1)
_ = W2473
tmp5798 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2473)


if True == tmp5798 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5795 := MakeNative(func(__e *ControlFlow) {
W2474 := __e.Get(1)
_ = W2474
__e.TailApply(PrimFunc(symshen_4comb), W2474, symshen_4skip)
return
}, 1)

tmp5796 := Call(__e, PrimFunc(symshen_4in_1_6), W2473)


__e.TailApply(tmp5795, tmp5796)
return


}


}, 1)

tmp5799 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2472)


__e.TailApply(tmp5794, tmp5799)
return


}, 1)

tmp5800 := Call(__e, PrimFunc(symshen_4in_1_6), W2471)


__e.TailApply(tmp5793, tmp5800)
return


}


}, 1)

tmp5803 := Call(__e, PrimFunc(symshen_4_5times_6), W2470)


__e.TailApply(tmp5792, tmp5803)
return


}, 1)

tmp5804 := Call(__e, PrimFunc(symshen_4in_1_6), W2469)


__e.TailApply(tmp5791, tmp5804)
return


}


}, 1)

tmp5807 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2467)


tmp5808 := Call(__e, tmp5790, tmp5807)


__e.TailApply(tmp5787, tmp5808)
return


}, 1)

tmp5809 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp5786)


_ = tmp5809

tmp5810 := MakeNative(func(__e *ControlFlow) {
V2475 := __e.Get(1)
_ = V2475
tmp5811 := MakeNative(func(__e *ControlFlow) {
W2476 := __e.Get(1)
_ = W2476
tmp5813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2476)


if True == tmp5813 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2476)
return
}


}, 1)

tmp5819 := Call(__e, PrimFunc(symshen_4hds_a_2), V2475, MakeNumber(42))


var ifres5814 Obj

if True == tmp5819 {
tmp5815 := MakeNative(func(__e *ControlFlow) {
W2477 := __e.Get(1)
_ = W2477
__e.TailApply(PrimFunc(symshen_4comb), W2477, symshen_4skip)
return
}, 1)

tmp5816 := Call(__e, PrimFunc(symtail), V2475)


tmp5817 := Call(__e, tmp5815, tmp5816)


ifres5814 = tmp5817


} else {
tmp5818 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5814 = tmp5818


}

__e.TailApply(tmp5811, ifres5814)
return


}, 1)

tmp5820 := Call(__e, ns2_1set, symshen_4_5times_6, tmp5810)


_ = tmp5820

tmp5821 := MakeNative(func(__e *ControlFlow) {
V2478 := __e.Get(1)
_ = V2478
tmp5822 := MakeNative(func(__e *ControlFlow) {
W2479 := __e.Get(1)
_ = W2479
tmp5855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2479)


if True == tmp5855 {
tmp5823 := MakeNative(func(__e *ControlFlow) {
W2484 := __e.Get(1)
_ = W2484
tmp5840 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2484)


if True == tmp5840 {
tmp5824 := MakeNative(func(__e *ControlFlow) {
W2489 := __e.Get(1)
_ = W2489
tmp5826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2489)


if True == tmp5826 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2489)
return
}


}, 1)

tmp5838 := PrimIsPair(V2478)

var ifres5827 Obj

if True == tmp5838 {
tmp5828 := MakeNative(func(__e *ControlFlow) {
W2490 := __e.Get(1)
_ = W2490
tmp5829 := MakeNative(func(__e *ControlFlow) {
W2491 := __e.Get(1)
_ = W2491
tmp5833 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2491)


if True == tmp5833 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5830 := MakeNative(func(__e *ControlFlow) {
W2492 := __e.Get(1)
_ = W2492
__e.TailApply(PrimFunc(symshen_4comb), W2492, symshen_4skip)
return
}, 1)

tmp5831 := Call(__e, PrimFunc(symshen_4in_1_6), W2491)


__e.TailApply(tmp5830, tmp5831)
return


}


}, 1)

tmp5834 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2490)


__e.TailApply(tmp5829, tmp5834)
return


}, 1)

tmp5835 := Call(__e, PrimFunc(symtail), V2478)


tmp5836 := Call(__e, tmp5828, tmp5835)


ifres5827 = tmp5836


} else {
tmp5837 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5827 = tmp5837


}

__e.TailApply(tmp5824, ifres5827)
return


} else {
__e.Return(W2484)
return
}


}, 1)

tmp5841 := MakeNative(func(__e *ControlFlow) {
W2485 := __e.Get(1)
_ = W2485
tmp5851 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2485)


if True == tmp5851 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5842 := MakeNative(func(__e *ControlFlow) {
W2486 := __e.Get(1)
_ = W2486
tmp5843 := MakeNative(func(__e *ControlFlow) {
W2487 := __e.Get(1)
_ = W2487
tmp5847 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2487)


if True == tmp5847 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5844 := MakeNative(func(__e *ControlFlow) {
W2488 := __e.Get(1)
_ = W2488
__e.TailApply(PrimFunc(symshen_4comb), W2488, symshen_4skip)
return
}, 1)

tmp5845 := Call(__e, PrimFunc(symshen_4in_1_6), W2487)


__e.TailApply(tmp5844, tmp5845)
return


}


}, 1)

tmp5848 := Call(__e, PrimFunc(symshen_4_5backslash_6), W2486)


__e.TailApply(tmp5843, tmp5848)
return


}, 1)

tmp5849 := Call(__e, PrimFunc(symshen_4in_1_6), W2485)


__e.TailApply(tmp5842, tmp5849)
return


}


}, 1)

tmp5852 := Call(__e, PrimFunc(symshen_4_5times_6), V2478)


tmp5853 := Call(__e, tmp5841, tmp5852)


__e.TailApply(tmp5823, tmp5853)
return


} else {
__e.Return(W2479)
return
}


}, 1)

tmp5856 := MakeNative(func(__e *ControlFlow) {
W2480 := __e.Get(1)
_ = W2480
tmp5866 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2480)


if True == tmp5866 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5857 := MakeNative(func(__e *ControlFlow) {
W2481 := __e.Get(1)
_ = W2481
tmp5858 := MakeNative(func(__e *ControlFlow) {
W2482 := __e.Get(1)
_ = W2482
tmp5862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2482)


if True == tmp5862 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5859 := MakeNative(func(__e *ControlFlow) {
W2483 := __e.Get(1)
_ = W2483
__e.TailApply(PrimFunc(symshen_4comb), W2483, symshen_4skip)
return
}, 1)

tmp5860 := Call(__e, PrimFunc(symshen_4in_1_6), W2482)


__e.TailApply(tmp5859, tmp5860)
return


}


}, 1)

tmp5863 := Call(__e, PrimFunc(symshen_4_5longnatter_6), W2481)


__e.TailApply(tmp5858, tmp5863)
return


}, 1)

tmp5864 := Call(__e, PrimFunc(symshen_4in_1_6), W2480)


__e.TailApply(tmp5857, tmp5864)
return


}


}, 1)

tmp5867 := Call(__e, PrimFunc(symshen_4_5comment_6), V2478)


tmp5868 := Call(__e, tmp5856, tmp5867)


__e.TailApply(tmp5822, tmp5868)
return


}, 1)

tmp5869 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp5821)


_ = tmp5869

tmp5870 := MakeNative(func(__e *ControlFlow) {
V2493 := __e.Get(1)
_ = V2493
tmp5871 := MakeNative(func(__e *ControlFlow) {
W2494 := __e.Get(1)
_ = W2494
tmp5902 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2494)


if True == tmp5902 {
tmp5872 := MakeNative(func(__e *ControlFlow) {
W2498 := __e.Get(1)
_ = W2498
tmp5891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2498)


if True == tmp5891 {
tmp5873 := MakeNative(func(__e *ControlFlow) {
W2502 := __e.Get(1)
_ = W2502
tmp5875 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2502)


if True == tmp5875 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2502)
return
}


}, 1)

tmp5876 := MakeNative(func(__e *ControlFlow) {
W2503 := __e.Get(1)
_ = W2503
tmp5887 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2503)


if True == tmp5887 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5877 := MakeNative(func(__e *ControlFlow) {
W2504 := __e.Get(1)
_ = W2504
tmp5878 := MakeNative(func(__e *ControlFlow) {
W2505 := __e.Get(1)
_ = W2505
tmp5883 := PrimEqual(W2504, MakeString("<>"))

var ifres5879 Obj

if True == tmp5883 {
tmp5880 := PrimCons(MakeNumber(0), Nil)

tmp5881 := PrimCons(symvector, tmp5880)

ifres5879 = tmp5881


} else {
tmp5882 := PrimIntern(W2504)

ifres5879 = tmp5882


}

__e.TailApply(PrimFunc(symshen_4comb), W2505, ifres5879)
return


}, 1)

tmp5884 := Call(__e, PrimFunc(symshen_4in_1_6), W2503)


__e.TailApply(tmp5878, tmp5884)
return


}, 1)

tmp5885 := Call(__e, PrimFunc(symshen_4_5_1out), W2503)


__e.TailApply(tmp5877, tmp5885)
return


}


}, 1)

tmp5888 := Call(__e, PrimFunc(symshen_4_5sym_6), V2493)


tmp5889 := Call(__e, tmp5876, tmp5888)


__e.TailApply(tmp5873, tmp5889)
return


} else {
__e.Return(W2498)
return
}


}, 1)

tmp5892 := MakeNative(func(__e *ControlFlow) {
W2499 := __e.Get(1)
_ = W2499
tmp5898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2499)


if True == tmp5898 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5893 := MakeNative(func(__e *ControlFlow) {
W2500 := __e.Get(1)
_ = W2500
tmp5894 := MakeNative(func(__e *ControlFlow) {
W2501 := __e.Get(1)
_ = W2501
__e.TailApply(PrimFunc(symshen_4comb), W2501, W2500)
return
}, 1)

tmp5895 := Call(__e, PrimFunc(symshen_4in_1_6), W2499)


__e.TailApply(tmp5894, tmp5895)
return


}, 1)

tmp5896 := Call(__e, PrimFunc(symshen_4_5_1out), W2499)


__e.TailApply(tmp5893, tmp5896)
return


}


}, 1)

tmp5899 := Call(__e, PrimFunc(symshen_4_5number_6), V2493)


tmp5900 := Call(__e, tmp5892, tmp5899)


__e.TailApply(tmp5872, tmp5900)
return


} else {
__e.Return(W2494)
return
}


}, 1)

tmp5903 := MakeNative(func(__e *ControlFlow) {
W2495 := __e.Get(1)
_ = W2495
tmp5909 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2495)


if True == tmp5909 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5904 := MakeNative(func(__e *ControlFlow) {
W2496 := __e.Get(1)
_ = W2496
tmp5905 := MakeNative(func(__e *ControlFlow) {
W2497 := __e.Get(1)
_ = W2497
__e.TailApply(PrimFunc(symshen_4comb), W2497, W2496)
return
}, 1)

tmp5906 := Call(__e, PrimFunc(symshen_4in_1_6), W2495)


__e.TailApply(tmp5905, tmp5906)
return


}, 1)

tmp5907 := Call(__e, PrimFunc(symshen_4_5_1out), W2495)


__e.TailApply(tmp5904, tmp5907)
return


}


}, 1)

tmp5910 := Call(__e, PrimFunc(symshen_4_5str_6), V2493)


tmp5911 := Call(__e, tmp5903, tmp5910)


__e.TailApply(tmp5871, tmp5911)
return


}, 1)

tmp5912 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp5870)


_ = tmp5912

tmp5913 := MakeNative(func(__e *ControlFlow) {
V2506 := __e.Get(1)
_ = V2506
tmp5914 := MakeNative(func(__e *ControlFlow) {
W2507 := __e.Get(1)
_ = W2507
tmp5916 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2507)


if True == tmp5916 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2507)
return
}


}, 1)

tmp5917 := MakeNative(func(__e *ControlFlow) {
W2508 := __e.Get(1)
_ = W2508
tmp5932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2508)


if True == tmp5932 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5918 := MakeNative(func(__e *ControlFlow) {
W2509 := __e.Get(1)
_ = W2509
tmp5919 := MakeNative(func(__e *ControlFlow) {
W2510 := __e.Get(1)
_ = W2510
tmp5920 := MakeNative(func(__e *ControlFlow) {
W2511 := __e.Get(1)
_ = W2511
tmp5927 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2511)


if True == tmp5927 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp5921 := MakeNative(func(__e *ControlFlow) {
W2512 := __e.Get(1)
_ = W2512
tmp5922 := MakeNative(func(__e *ControlFlow) {
W2513 := __e.Get(1)
_ = W2513
tmp5923 := PrimStringConcat(W2509, W2512)

__e.TailApply(PrimFunc(symshen_4comb), W2513, tmp5923)
return


}, 1)

tmp5924 := Call(__e, PrimFunc(symshen_4in_1_6), W2511)


__e.TailApply(tmp5922, tmp5924)
return


}, 1)

tmp5925 := Call(__e, PrimFunc(symshen_4_5_1out), W2511)


__e.TailApply(tmp5921, tmp5925)
return


}


}, 1)

tmp5928 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2510)


__e.TailApply(tmp5920, tmp5928)
return


}, 1)

tmp5929 := Call(__e, PrimFunc(symshen_4in_1_6), W2508)


__e.TailApply(tmp5919, tmp5929)
return


}, 1)

tmp5930 := Call(__e, PrimFunc(symshen_4_5_1out), W2508)


__e.TailApply(tmp5918, tmp5930)
return


}


}, 1)

tmp5933 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2506)


tmp5934 := Call(__e, tmp5917, tmp5933)


__e.TailApply(tmp5914, tmp5934)
return


}, 1)

tmp5935 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp5913)


_ = tmp5935

tmp5936 := MakeNative(func(__e *ControlFlow) {
V2514 := __e.Get(1)
_ = V2514
tmp5937 := MakeNative(func(__e *ControlFlow) {
W2515 := __e.Get(1)
_ = W2515
tmp5939 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2515)


if True == tmp5939 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2515)
return
}


}, 1)

tmp5950 := PrimIsPair(V2514)

var ifres5940 Obj

if True == tmp5950 {
tmp5941 := MakeNative(func(__e *ControlFlow) {
W2516 := __e.Get(1)
_ = W2516
tmp5942 := MakeNative(func(__e *ControlFlow) {
W2517 := __e.Get(1)
_ = W2517
tmp5945 := Call(__e, PrimFunc(symshen_4alpha_2), W2516)


if True == tmp5945 {
tmp5943 := PrimNumberToString(W2516)

__e.TailApply(PrimFunc(symshen_4comb), W2517, tmp5943)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp5946 := Call(__e, PrimFunc(symtail), V2514)


__e.TailApply(tmp5942, tmp5946)
return


}, 1)

tmp5947 := Call(__e, PrimFunc(symhead), V2514)


tmp5948 := Call(__e, tmp5941, tmp5947)


ifres5940 = tmp5948


} else {
tmp5949 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres5940 = tmp5949


}

__e.TailApply(tmp5937, ifres5940)
return


}, 1)

tmp5951 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp5936)


_ = tmp5951

tmp5952 := MakeNative(func(__e *ControlFlow) {
V2518 := __e.Get(1)
_ = V2518
tmp5959 := Call(__e, PrimFunc(symshen_4lowercase_2), V2518)


if True == tmp5959 {
__e.Return(True)
return
} else {
tmp5957 := Call(__e, PrimFunc(symshen_4uppercase_2), V2518)


var ifres5954 Obj

if True == tmp5957 {
ifres5954 = True


} else {
tmp5956 := Call(__e, PrimFunc(symshen_4misc_2), V2518)


var ifres5955 Obj

if True == tmp5956 {
ifres5955 = True


} else {
ifres5955 = False


}

ifres5954 = ifres5955


}

if True == ifres5954 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp5960 := Call(__e, ns2_1set, symshen_4alpha_2, tmp5952)


_ = tmp5960

tmp5961 := MakeNative(func(__e *ControlFlow) {
V2519 := __e.Get(1)
_ = V2519
tmp5965 := PrimGreatEqual(V2519, MakeNumber(97))

if True == tmp5965 {
tmp5963 := PrimLessEqual(V2519, MakeNumber(122))

if True == tmp5963 {
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


}, 1)

tmp5966 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp5961)


_ = tmp5966

tmp5967 := MakeNative(func(__e *ControlFlow) {
V2520 := __e.Get(1)
_ = V2520
tmp5971 := PrimGreatEqual(V2520, MakeNumber(65))

if True == tmp5971 {
tmp5969 := PrimLessEqual(V2520, MakeNumber(90))

if True == tmp5969 {
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


}, 1)

tmp5972 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp5967)


_ = tmp5972

tmp5973 := MakeNative(func(__e *ControlFlow) {
V2521 := __e.Get(1)
_ = V2521
tmp5974 := PrimCons(MakeNumber(96), Nil)

tmp5975 := PrimCons(MakeNumber(35), tmp5974)

tmp5976 := PrimCons(MakeNumber(39), tmp5975)

tmp5977 := PrimCons(MakeNumber(37), tmp5976)

tmp5978 := PrimCons(MakeNumber(38), tmp5977)

tmp5979 := PrimCons(MakeNumber(60), tmp5978)

tmp5980 := PrimCons(MakeNumber(62), tmp5979)

tmp5981 := PrimCons(MakeNumber(46), tmp5980)

tmp5982 := PrimCons(MakeNumber(126), tmp5981)

tmp5983 := PrimCons(MakeNumber(64), tmp5982)

tmp5984 := PrimCons(MakeNumber(33), tmp5983)

tmp5985 := PrimCons(MakeNumber(36), tmp5984)

tmp5986 := PrimCons(MakeNumber(63), tmp5985)

tmp5987 := PrimCons(MakeNumber(95), tmp5986)

tmp5988 := PrimCons(MakeNumber(43), tmp5987)

tmp5989 := PrimCons(MakeNumber(47), tmp5988)

tmp5990 := PrimCons(MakeNumber(42), tmp5989)

tmp5991 := PrimCons(MakeNumber(45), tmp5990)

tmp5992 := PrimCons(MakeNumber(61), tmp5991)

__e.TailApply(PrimFunc(symelement_2), V2521, tmp5992)
return


}, 1)

tmp5993 := Call(__e, ns2_1set, symshen_4misc_2, tmp5973)


_ = tmp5993

tmp5994 := MakeNative(func(__e *ControlFlow) {
V2522 := __e.Get(1)
_ = V2522
tmp5995 := MakeNative(func(__e *ControlFlow) {
W2523 := __e.Get(1)
_ = W2523
tmp6007 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2523)


if True == tmp6007 {
tmp5996 := MakeNative(func(__e *ControlFlow) {
W2530 := __e.Get(1)
_ = W2530
tmp5998 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2530)


if True == tmp5998 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2530)
return
}


}, 1)

tmp5999 := MakeNative(func(__e *ControlFlow) {
W2531 := __e.Get(1)
_ = W2531
tmp6003 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2531)


if True == tmp6003 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6000 := MakeNative(func(__e *ControlFlow) {
W2532 := __e.Get(1)
_ = W2532
__e.TailApply(PrimFunc(symshen_4comb), W2532, MakeString(""))
return
}, 1)

tmp6001 := Call(__e, PrimFunc(symshen_4in_1_6), W2531)


__e.TailApply(tmp6000, tmp6001)
return


}


}, 1)

tmp6004 := Call(__e, PrimFunc(sym_5e_6), V2522)


tmp6005 := Call(__e, tmp5999, tmp6004)


__e.TailApply(tmp5996, tmp6005)
return


} else {
__e.Return(W2523)
return
}


}, 1)

tmp6008 := MakeNative(func(__e *ControlFlow) {
W2524 := __e.Get(1)
_ = W2524
tmp6023 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2524)


if True == tmp6023 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6009 := MakeNative(func(__e *ControlFlow) {
W2525 := __e.Get(1)
_ = W2525
tmp6010 := MakeNative(func(__e *ControlFlow) {
W2526 := __e.Get(1)
_ = W2526
tmp6011 := MakeNative(func(__e *ControlFlow) {
W2527 := __e.Get(1)
_ = W2527
tmp6018 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2527)


if True == tmp6018 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6012 := MakeNative(func(__e *ControlFlow) {
W2528 := __e.Get(1)
_ = W2528
tmp6013 := MakeNative(func(__e *ControlFlow) {
W2529 := __e.Get(1)
_ = W2529
tmp6014 := PrimStringConcat(W2525, W2528)

__e.TailApply(PrimFunc(symshen_4comb), W2529, tmp6014)
return


}, 1)

tmp6015 := Call(__e, PrimFunc(symshen_4in_1_6), W2527)


__e.TailApply(tmp6013, tmp6015)
return


}, 1)

tmp6016 := Call(__e, PrimFunc(symshen_4_5_1out), W2527)


__e.TailApply(tmp6012, tmp6016)
return


}


}, 1)

tmp6019 := Call(__e, PrimFunc(symshen_4_5alphanums_6), W2526)


__e.TailApply(tmp6011, tmp6019)
return


}, 1)

tmp6020 := Call(__e, PrimFunc(symshen_4in_1_6), W2524)


__e.TailApply(tmp6010, tmp6020)
return


}, 1)

tmp6021 := Call(__e, PrimFunc(symshen_4_5_1out), W2524)


__e.TailApply(tmp6009, tmp6021)
return


}


}, 1)

tmp6024 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2522)


tmp6025 := Call(__e, tmp6008, tmp6024)


__e.TailApply(tmp5995, tmp6025)
return


}, 1)

tmp6026 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp5994)


_ = tmp6026

tmp6027 := MakeNative(func(__e *ControlFlow) {
V2533 := __e.Get(1)
_ = V2533
tmp6028 := MakeNative(func(__e *ControlFlow) {
W2534 := __e.Get(1)
_ = W2534
tmp6042 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2534)


if True == tmp6042 {
tmp6029 := MakeNative(func(__e *ControlFlow) {
W2538 := __e.Get(1)
_ = W2538
tmp6031 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2538)


if True == tmp6031 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2538)
return
}


}, 1)

tmp6032 := MakeNative(func(__e *ControlFlow) {
W2539 := __e.Get(1)
_ = W2539
tmp6038 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2539)


if True == tmp6038 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6033 := MakeNative(func(__e *ControlFlow) {
W2540 := __e.Get(1)
_ = W2540
tmp6034 := MakeNative(func(__e *ControlFlow) {
W2541 := __e.Get(1)
_ = W2541
__e.TailApply(PrimFunc(symshen_4comb), W2541, W2540)
return
}, 1)

tmp6035 := Call(__e, PrimFunc(symshen_4in_1_6), W2539)


__e.TailApply(tmp6034, tmp6035)
return


}, 1)

tmp6036 := Call(__e, PrimFunc(symshen_4_5_1out), W2539)


__e.TailApply(tmp6033, tmp6036)
return


}


}, 1)

tmp6039 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2533)


tmp6040 := Call(__e, tmp6032, tmp6039)


__e.TailApply(tmp6029, tmp6040)
return


} else {
__e.Return(W2534)
return
}


}, 1)

tmp6043 := MakeNative(func(__e *ControlFlow) {
W2535 := __e.Get(1)
_ = W2535
tmp6049 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2535)


if True == tmp6049 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6044 := MakeNative(func(__e *ControlFlow) {
W2536 := __e.Get(1)
_ = W2536
tmp6045 := MakeNative(func(__e *ControlFlow) {
W2537 := __e.Get(1)
_ = W2537
__e.TailApply(PrimFunc(symshen_4comb), W2537, W2536)
return
}, 1)

tmp6046 := Call(__e, PrimFunc(symshen_4in_1_6), W2535)


__e.TailApply(tmp6045, tmp6046)
return


}, 1)

tmp6047 := Call(__e, PrimFunc(symshen_4_5_1out), W2535)


__e.TailApply(tmp6044, tmp6047)
return


}


}, 1)

tmp6050 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2533)


tmp6051 := Call(__e, tmp6043, tmp6050)


__e.TailApply(tmp6028, tmp6051)
return


}, 1)

tmp6052 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp6027)


_ = tmp6052

tmp6053 := MakeNative(func(__e *ControlFlow) {
V2542 := __e.Get(1)
_ = V2542
tmp6054 := MakeNative(func(__e *ControlFlow) {
W2543 := __e.Get(1)
_ = W2543
tmp6056 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2543)


if True == tmp6056 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2543)
return
}


}, 1)

tmp6067 := PrimIsPair(V2542)

var ifres6057 Obj

if True == tmp6067 {
tmp6058 := MakeNative(func(__e *ControlFlow) {
W2544 := __e.Get(1)
_ = W2544
tmp6059 := MakeNative(func(__e *ControlFlow) {
W2545 := __e.Get(1)
_ = W2545
tmp6062 := Call(__e, PrimFunc(symshen_4digit_2), W2544)


if True == tmp6062 {
tmp6060 := PrimNumberToString(W2544)

__e.TailApply(PrimFunc(symshen_4comb), W2545, tmp6060)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6063 := Call(__e, PrimFunc(symtail), V2542)


__e.TailApply(tmp6059, tmp6063)
return


}, 1)

tmp6064 := Call(__e, PrimFunc(symhead), V2542)


tmp6065 := Call(__e, tmp6058, tmp6064)


ifres6057 = tmp6065


} else {
tmp6066 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6057 = tmp6066


}

__e.TailApply(tmp6054, ifres6057)
return


}, 1)

tmp6068 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp6053)


_ = tmp6068

tmp6069 := MakeNative(func(__e *ControlFlow) {
V2546 := __e.Get(1)
_ = V2546
tmp6073 := PrimGreatEqual(V2546, MakeNumber(48))

if True == tmp6073 {
tmp6071 := PrimLessEqual(V2546, MakeNumber(57))

if True == tmp6071 {
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


}, 1)

tmp6074 := Call(__e, ns2_1set, symshen_4digit_2, tmp6069)


_ = tmp6074

tmp6075 := MakeNative(func(__e *ControlFlow) {
V2547 := __e.Get(1)
_ = V2547
tmp6076 := MakeNative(func(__e *ControlFlow) {
W2548 := __e.Get(1)
_ = W2548
tmp6078 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2548)


if True == tmp6078 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2548)
return
}


}, 1)

tmp6079 := MakeNative(func(__e *ControlFlow) {
W2549 := __e.Get(1)
_ = W2549
tmp6097 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2549)


if True == tmp6097 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6080 := MakeNative(func(__e *ControlFlow) {
W2550 := __e.Get(1)
_ = W2550
tmp6081 := MakeNative(func(__e *ControlFlow) {
W2551 := __e.Get(1)
_ = W2551
tmp6093 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2551)


if True == tmp6093 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6082 := MakeNative(func(__e *ControlFlow) {
W2552 := __e.Get(1)
_ = W2552
tmp6083 := MakeNative(func(__e *ControlFlow) {
W2553 := __e.Get(1)
_ = W2553
tmp6084 := MakeNative(func(__e *ControlFlow) {
W2554 := __e.Get(1)
_ = W2554
tmp6088 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2554)


if True == tmp6088 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6085 := MakeNative(func(__e *ControlFlow) {
W2555 := __e.Get(1)
_ = W2555
__e.TailApply(PrimFunc(symshen_4comb), W2555, W2552)
return
}, 1)

tmp6086 := Call(__e, PrimFunc(symshen_4in_1_6), W2554)


__e.TailApply(tmp6085, tmp6086)
return


}


}, 1)

tmp6089 := Call(__e, PrimFunc(symshen_4_5dbq_6), W2553)


__e.TailApply(tmp6084, tmp6089)
return


}, 1)

tmp6090 := Call(__e, PrimFunc(symshen_4in_1_6), W2551)


__e.TailApply(tmp6083, tmp6090)
return


}, 1)

tmp6091 := Call(__e, PrimFunc(symshen_4_5_1out), W2551)


__e.TailApply(tmp6082, tmp6091)
return


}


}, 1)

tmp6094 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2550)


__e.TailApply(tmp6081, tmp6094)
return


}, 1)

tmp6095 := Call(__e, PrimFunc(symshen_4in_1_6), W2549)


__e.TailApply(tmp6080, tmp6095)
return


}


}, 1)

tmp6098 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2547)


tmp6099 := Call(__e, tmp6079, tmp6098)


__e.TailApply(tmp6076, tmp6099)
return


}, 1)

tmp6100 := Call(__e, ns2_1set, symshen_4_5str_6, tmp6075)


_ = tmp6100

tmp6101 := MakeNative(func(__e *ControlFlow) {
V2556 := __e.Get(1)
_ = V2556
tmp6102 := MakeNative(func(__e *ControlFlow) {
W2557 := __e.Get(1)
_ = W2557
tmp6104 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2557)


if True == tmp6104 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2557)
return
}


}, 1)

tmp6110 := Call(__e, PrimFunc(symshen_4hds_a_2), V2556, MakeNumber(34))


var ifres6105 Obj

if True == tmp6110 {
tmp6106 := MakeNative(func(__e *ControlFlow) {
W2558 := __e.Get(1)
_ = W2558
__e.TailApply(PrimFunc(symshen_4comb), W2558, symshen_4skip)
return
}, 1)

tmp6107 := Call(__e, PrimFunc(symtail), V2556)


tmp6108 := Call(__e, tmp6106, tmp6107)


ifres6105 = tmp6108


} else {
tmp6109 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6105 = tmp6109


}

__e.TailApply(tmp6102, ifres6105)
return


}, 1)

tmp6111 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp6101)


_ = tmp6111

tmp6112 := MakeNative(func(__e *ControlFlow) {
V2559 := __e.Get(1)
_ = V2559
tmp6113 := MakeNative(func(__e *ControlFlow) {
W2560 := __e.Get(1)
_ = W2560
tmp6125 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2560)


if True == tmp6125 {
tmp6114 := MakeNative(func(__e *ControlFlow) {
W2567 := __e.Get(1)
_ = W2567
tmp6116 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2567)


if True == tmp6116 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2567)
return
}


}, 1)

tmp6117 := MakeNative(func(__e *ControlFlow) {
W2568 := __e.Get(1)
_ = W2568
tmp6121 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2568)


if True == tmp6121 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6118 := MakeNative(func(__e *ControlFlow) {
W2569 := __e.Get(1)
_ = W2569
__e.TailApply(PrimFunc(symshen_4comb), W2569, MakeString(""))
return
}, 1)

tmp6119 := Call(__e, PrimFunc(symshen_4in_1_6), W2568)


__e.TailApply(tmp6118, tmp6119)
return


}


}, 1)

tmp6122 := Call(__e, PrimFunc(sym_5e_6), V2559)


tmp6123 := Call(__e, tmp6117, tmp6122)


__e.TailApply(tmp6114, tmp6123)
return


} else {
__e.Return(W2560)
return
}


}, 1)

tmp6126 := MakeNative(func(__e *ControlFlow) {
W2561 := __e.Get(1)
_ = W2561
tmp6141 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2561)


if True == tmp6141 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6127 := MakeNative(func(__e *ControlFlow) {
W2562 := __e.Get(1)
_ = W2562
tmp6128 := MakeNative(func(__e *ControlFlow) {
W2563 := __e.Get(1)
_ = W2563
tmp6129 := MakeNative(func(__e *ControlFlow) {
W2564 := __e.Get(1)
_ = W2564
tmp6136 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2564)


if True == tmp6136 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6130 := MakeNative(func(__e *ControlFlow) {
W2565 := __e.Get(1)
_ = W2565
tmp6131 := MakeNative(func(__e *ControlFlow) {
W2566 := __e.Get(1)
_ = W2566
tmp6132 := PrimStringConcat(W2562, W2565)

__e.TailApply(PrimFunc(symshen_4comb), W2566, tmp6132)
return


}, 1)

tmp6133 := Call(__e, PrimFunc(symshen_4in_1_6), W2564)


__e.TailApply(tmp6131, tmp6133)
return


}, 1)

tmp6134 := Call(__e, PrimFunc(symshen_4_5_1out), W2564)


__e.TailApply(tmp6130, tmp6134)
return


}


}, 1)

tmp6137 := Call(__e, PrimFunc(symshen_4_5strcontents_6), W2563)


__e.TailApply(tmp6129, tmp6137)
return


}, 1)

tmp6138 := Call(__e, PrimFunc(symshen_4in_1_6), W2561)


__e.TailApply(tmp6128, tmp6138)
return


}, 1)

tmp6139 := Call(__e, PrimFunc(symshen_4_5_1out), W2561)


__e.TailApply(tmp6127, tmp6139)
return


}


}, 1)

tmp6142 := Call(__e, PrimFunc(symshen_4_5strc_6), V2559)


tmp6143 := Call(__e, tmp6126, tmp6142)


__e.TailApply(tmp6113, tmp6143)
return


}, 1)

tmp6144 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp6112)


_ = tmp6144

tmp6145 := MakeNative(func(__e *ControlFlow) {
V2570 := __e.Get(1)
_ = V2570
tmp6146 := MakeNative(func(__e *ControlFlow) {
W2571 := __e.Get(1)
_ = W2571
tmp6160 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2571)


if True == tmp6160 {
tmp6147 := MakeNative(func(__e *ControlFlow) {
W2575 := __e.Get(1)
_ = W2575
tmp6149 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2575)


if True == tmp6149 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2575)
return
}


}, 1)

tmp6150 := MakeNative(func(__e *ControlFlow) {
W2576 := __e.Get(1)
_ = W2576
tmp6156 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2576)


if True == tmp6156 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6151 := MakeNative(func(__e *ControlFlow) {
W2577 := __e.Get(1)
_ = W2577
tmp6152 := MakeNative(func(__e *ControlFlow) {
W2578 := __e.Get(1)
_ = W2578
__e.TailApply(PrimFunc(symshen_4comb), W2578, W2577)
return
}, 1)

tmp6153 := Call(__e, PrimFunc(symshen_4in_1_6), W2576)


__e.TailApply(tmp6152, tmp6153)
return


}, 1)

tmp6154 := Call(__e, PrimFunc(symshen_4_5_1out), W2576)


__e.TailApply(tmp6151, tmp6154)
return


}


}, 1)

tmp6157 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2570)


tmp6158 := Call(__e, tmp6150, tmp6157)


__e.TailApply(tmp6147, tmp6158)
return


} else {
__e.Return(W2571)
return
}


}, 1)

tmp6161 := MakeNative(func(__e *ControlFlow) {
W2572 := __e.Get(1)
_ = W2572
tmp6167 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2572)


if True == tmp6167 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6162 := MakeNative(func(__e *ControlFlow) {
W2573 := __e.Get(1)
_ = W2573
tmp6163 := MakeNative(func(__e *ControlFlow) {
W2574 := __e.Get(1)
_ = W2574
__e.TailApply(PrimFunc(symshen_4comb), W2574, W2573)
return
}, 1)

tmp6164 := Call(__e, PrimFunc(symshen_4in_1_6), W2572)


__e.TailApply(tmp6163, tmp6164)
return


}, 1)

tmp6165 := Call(__e, PrimFunc(symshen_4_5_1out), W2572)


__e.TailApply(tmp6162, tmp6165)
return


}


}, 1)

tmp6168 := Call(__e, PrimFunc(symshen_4_5control_6), V2570)


tmp6169 := Call(__e, tmp6161, tmp6168)


__e.TailApply(tmp6146, tmp6169)
return


}, 1)

tmp6170 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp6145)


_ = tmp6170

tmp6171 := MakeNative(func(__e *ControlFlow) {
V2579 := __e.Get(1)
_ = V2579
tmp6172 := MakeNative(func(__e *ControlFlow) {
W2580 := __e.Get(1)
_ = W2580
tmp6174 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2580)


if True == tmp6174 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2580)
return
}


}, 1)

tmp6175 := MakeNative(func(__e *ControlFlow) {
W2581 := __e.Get(1)
_ = W2581
tmp6200 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2581)


if True == tmp6200 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6176 := MakeNative(func(__e *ControlFlow) {
W2582 := __e.Get(1)
_ = W2582
tmp6177 := MakeNative(func(__e *ControlFlow) {
W2583 := __e.Get(1)
_ = W2583
tmp6196 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2583)


if True == tmp6196 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6178 := MakeNative(func(__e *ControlFlow) {
W2584 := __e.Get(1)
_ = W2584
tmp6179 := MakeNative(func(__e *ControlFlow) {
W2585 := __e.Get(1)
_ = W2585
tmp6192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2585)


if True == tmp6192 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6180 := MakeNative(func(__e *ControlFlow) {
W2586 := __e.Get(1)
_ = W2586
tmp6181 := MakeNative(func(__e *ControlFlow) {
W2587 := __e.Get(1)
_ = W2587
tmp6182 := MakeNative(func(__e *ControlFlow) {
W2588 := __e.Get(1)
_ = W2588
tmp6187 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2588)


if True == tmp6187 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6183 := MakeNative(func(__e *ControlFlow) {
W2589 := __e.Get(1)
_ = W2589
tmp6184 := PrimNumberToString(W2586)

__e.TailApply(PrimFunc(symshen_4comb), W2589, tmp6184)
return


}, 1)

tmp6185 := Call(__e, PrimFunc(symshen_4in_1_6), W2588)


__e.TailApply(tmp6183, tmp6185)
return


}


}, 1)

tmp6188 := Call(__e, PrimFunc(symshen_4_5semicolon_6), W2587)


__e.TailApply(tmp6182, tmp6188)
return


}, 1)

tmp6189 := Call(__e, PrimFunc(symshen_4in_1_6), W2585)


__e.TailApply(tmp6181, tmp6189)
return


}, 1)

tmp6190 := Call(__e, PrimFunc(symshen_4_5_1out), W2585)


__e.TailApply(tmp6180, tmp6190)
return


}


}, 1)

tmp6193 := Call(__e, PrimFunc(symshen_4_5integer_6), W2584)


__e.TailApply(tmp6179, tmp6193)
return


}, 1)

tmp6194 := Call(__e, PrimFunc(symshen_4in_1_6), W2583)


__e.TailApply(tmp6178, tmp6194)
return


}


}, 1)

tmp6197 := Call(__e, PrimFunc(symshen_4_5hash_6), W2582)


__e.TailApply(tmp6177, tmp6197)
return


}, 1)

tmp6198 := Call(__e, PrimFunc(symshen_4in_1_6), W2581)


__e.TailApply(tmp6176, tmp6198)
return


}


}, 1)

tmp6201 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2579)


tmp6202 := Call(__e, tmp6175, tmp6201)


__e.TailApply(tmp6172, tmp6202)
return


}, 1)

tmp6203 := Call(__e, ns2_1set, symshen_4_5control_6, tmp6171)


_ = tmp6203

tmp6204 := MakeNative(func(__e *ControlFlow) {
V2590 := __e.Get(1)
_ = V2590
tmp6205 := MakeNative(func(__e *ControlFlow) {
W2591 := __e.Get(1)
_ = W2591
tmp6207 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2591)


if True == tmp6207 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2591)
return
}


}, 1)

tmp6219 := PrimIsPair(V2590)

var ifres6208 Obj

if True == tmp6219 {
tmp6209 := MakeNative(func(__e *ControlFlow) {
W2592 := __e.Get(1)
_ = W2592
tmp6210 := MakeNative(func(__e *ControlFlow) {
W2593 := __e.Get(1)
_ = W2593
tmp6213 := PrimEqual(W2592, MakeNumber(34))

tmp6214 := PrimNot(tmp6213)

if True == tmp6214 {
tmp6211 := PrimNumberToString(W2592)

__e.TailApply(PrimFunc(symshen_4comb), W2593, tmp6211)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6215 := Call(__e, PrimFunc(symtail), V2590)


__e.TailApply(tmp6210, tmp6215)
return


}, 1)

tmp6216 := Call(__e, PrimFunc(symhead), V2590)


tmp6217 := Call(__e, tmp6209, tmp6216)


ifres6208 = tmp6217


} else {
tmp6218 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6208 = tmp6218


}

__e.TailApply(tmp6205, ifres6208)
return


}, 1)

tmp6220 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp6204)


_ = tmp6220

tmp6221 := MakeNative(func(__e *ControlFlow) {
V2594 := __e.Get(1)
_ = V2594
tmp6222 := MakeNative(func(__e *ControlFlow) {
W2595 := __e.Get(1)
_ = W2595
tmp6224 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2595)


if True == tmp6224 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2595)
return
}


}, 1)

tmp6230 := Call(__e, PrimFunc(symshen_4hds_a_2), V2594, MakeNumber(99))


var ifres6225 Obj

if True == tmp6230 {
tmp6226 := MakeNative(func(__e *ControlFlow) {
W2596 := __e.Get(1)
_ = W2596
__e.TailApply(PrimFunc(symshen_4comb), W2596, symshen_4skip)
return
}, 1)

tmp6227 := Call(__e, PrimFunc(symtail), V2594)


tmp6228 := Call(__e, tmp6226, tmp6227)


ifres6225 = tmp6228


} else {
tmp6229 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6225 = tmp6229


}

__e.TailApply(tmp6222, ifres6225)
return


}, 1)

tmp6231 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp6221)


_ = tmp6231

tmp6232 := MakeNative(func(__e *ControlFlow) {
V2597 := __e.Get(1)
_ = V2597
tmp6233 := MakeNative(func(__e *ControlFlow) {
W2598 := __e.Get(1)
_ = W2598
tmp6235 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2598)


if True == tmp6235 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2598)
return
}


}, 1)

tmp6241 := Call(__e, PrimFunc(symshen_4hds_a_2), V2597, MakeNumber(35))


var ifres6236 Obj

if True == tmp6241 {
tmp6237 := MakeNative(func(__e *ControlFlow) {
W2599 := __e.Get(1)
_ = W2599
__e.TailApply(PrimFunc(symshen_4comb), W2599, symshen_4skip)
return
}, 1)

tmp6238 := Call(__e, PrimFunc(symtail), V2597)


tmp6239 := Call(__e, tmp6237, tmp6238)


ifres6236 = tmp6239


} else {
tmp6240 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6236 = tmp6240


}

__e.TailApply(tmp6233, ifres6236)
return


}, 1)

tmp6242 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp6232)


_ = tmp6242

tmp6243 := MakeNative(func(__e *ControlFlow) {
V2600 := __e.Get(1)
_ = V2600
tmp6244 := MakeNative(func(__e *ControlFlow) {
W2601 := __e.Get(1)
_ = W2601
tmp6300 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2601)


if True == tmp6300 {
tmp6245 := MakeNative(func(__e *ControlFlow) {
W2607 := __e.Get(1)
_ = W2607
tmp6283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2607)


if True == tmp6283 {
tmp6246 := MakeNative(func(__e *ControlFlow) {
W2613 := __e.Get(1)
_ = W2613
tmp6272 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2613)


if True == tmp6272 {
tmp6247 := MakeNative(func(__e *ControlFlow) {
W2617 := __e.Get(1)
_ = W2617
tmp6261 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2617)


if True == tmp6261 {
tmp6248 := MakeNative(func(__e *ControlFlow) {
W2621 := __e.Get(1)
_ = W2621
tmp6250 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2621)


if True == tmp6250 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2621)
return
}


}, 1)

tmp6251 := MakeNative(func(__e *ControlFlow) {
W2622 := __e.Get(1)
_ = W2622
tmp6257 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2622)


if True == tmp6257 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6252 := MakeNative(func(__e *ControlFlow) {
W2623 := __e.Get(1)
_ = W2623
tmp6253 := MakeNative(func(__e *ControlFlow) {
W2624 := __e.Get(1)
_ = W2624
__e.TailApply(PrimFunc(symshen_4comb), W2624, W2623)
return
}, 1)

tmp6254 := Call(__e, PrimFunc(symshen_4in_1_6), W2622)


__e.TailApply(tmp6253, tmp6254)
return


}, 1)

tmp6255 := Call(__e, PrimFunc(symshen_4_5_1out), W2622)


__e.TailApply(tmp6252, tmp6255)
return


}


}, 1)

tmp6258 := Call(__e, PrimFunc(symshen_4_5integer_6), V2600)


tmp6259 := Call(__e, tmp6251, tmp6258)


__e.TailApply(tmp6248, tmp6259)
return


} else {
__e.Return(W2617)
return
}


}, 1)

tmp6262 := MakeNative(func(__e *ControlFlow) {
W2618 := __e.Get(1)
_ = W2618
tmp6268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2618)


if True == tmp6268 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6263 := MakeNative(func(__e *ControlFlow) {
W2619 := __e.Get(1)
_ = W2619
tmp6264 := MakeNative(func(__e *ControlFlow) {
W2620 := __e.Get(1)
_ = W2620
__e.TailApply(PrimFunc(symshen_4comb), W2620, W2619)
return
}, 1)

tmp6265 := Call(__e, PrimFunc(symshen_4in_1_6), W2618)


__e.TailApply(tmp6264, tmp6265)
return


}, 1)

tmp6266 := Call(__e, PrimFunc(symshen_4_5_1out), W2618)


__e.TailApply(tmp6263, tmp6266)
return


}


}, 1)

tmp6269 := Call(__e, PrimFunc(symshen_4_5float_6), V2600)


tmp6270 := Call(__e, tmp6262, tmp6269)


__e.TailApply(tmp6247, tmp6270)
return


} else {
__e.Return(W2613)
return
}


}, 1)

tmp6273 := MakeNative(func(__e *ControlFlow) {
W2614 := __e.Get(1)
_ = W2614
tmp6279 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2614)


if True == tmp6279 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6274 := MakeNative(func(__e *ControlFlow) {
W2615 := __e.Get(1)
_ = W2615
tmp6275 := MakeNative(func(__e *ControlFlow) {
W2616 := __e.Get(1)
_ = W2616
__e.TailApply(PrimFunc(symshen_4comb), W2616, W2615)
return
}, 1)

tmp6276 := Call(__e, PrimFunc(symshen_4in_1_6), W2614)


__e.TailApply(tmp6275, tmp6276)
return


}, 1)

tmp6277 := Call(__e, PrimFunc(symshen_4_5_1out), W2614)


__e.TailApply(tmp6274, tmp6277)
return


}


}, 1)

tmp6280 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2600)


tmp6281 := Call(__e, tmp6273, tmp6280)


__e.TailApply(tmp6246, tmp6281)
return


} else {
__e.Return(W2607)
return
}


}, 1)

tmp6284 := MakeNative(func(__e *ControlFlow) {
W2608 := __e.Get(1)
_ = W2608
tmp6296 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2608)


if True == tmp6296 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6285 := MakeNative(func(__e *ControlFlow) {
W2609 := __e.Get(1)
_ = W2609
tmp6286 := MakeNative(func(__e *ControlFlow) {
W2610 := __e.Get(1)
_ = W2610
tmp6292 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2610)


if True == tmp6292 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6287 := MakeNative(func(__e *ControlFlow) {
W2611 := __e.Get(1)
_ = W2611
tmp6288 := MakeNative(func(__e *ControlFlow) {
W2612 := __e.Get(1)
_ = W2612
__e.TailApply(PrimFunc(symshen_4comb), W2612, W2611)
return
}, 1)

tmp6289 := Call(__e, PrimFunc(symshen_4in_1_6), W2610)


__e.TailApply(tmp6288, tmp6289)
return


}, 1)

tmp6290 := Call(__e, PrimFunc(symshen_4_5_1out), W2610)


__e.TailApply(tmp6287, tmp6290)
return


}


}, 1)

tmp6293 := Call(__e, PrimFunc(symshen_4_5number_6), W2609)


__e.TailApply(tmp6286, tmp6293)
return


}, 1)

tmp6294 := Call(__e, PrimFunc(symshen_4in_1_6), W2608)


__e.TailApply(tmp6285, tmp6294)
return


}


}, 1)

tmp6297 := Call(__e, PrimFunc(symshen_4_5plus_6), V2600)


tmp6298 := Call(__e, tmp6284, tmp6297)


__e.TailApply(tmp6245, tmp6298)
return


} else {
__e.Return(W2601)
return
}


}, 1)

tmp6301 := MakeNative(func(__e *ControlFlow) {
W2602 := __e.Get(1)
_ = W2602
tmp6314 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2602)


if True == tmp6314 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6302 := MakeNative(func(__e *ControlFlow) {
W2603 := __e.Get(1)
_ = W2603
tmp6303 := MakeNative(func(__e *ControlFlow) {
W2604 := __e.Get(1)
_ = W2604
tmp6310 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2604)


if True == tmp6310 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6304 := MakeNative(func(__e *ControlFlow) {
W2605 := __e.Get(1)
_ = W2605
tmp6305 := MakeNative(func(__e *ControlFlow) {
W2606 := __e.Get(1)
_ = W2606
tmp6306 := PrimNumberSubtract(MakeNumber(0), W2605)

__e.TailApply(PrimFunc(symshen_4comb), W2606, tmp6306)
return


}, 1)

tmp6307 := Call(__e, PrimFunc(symshen_4in_1_6), W2604)


__e.TailApply(tmp6305, tmp6307)
return


}, 1)

tmp6308 := Call(__e, PrimFunc(symshen_4_5_1out), W2604)


__e.TailApply(tmp6304, tmp6308)
return


}


}, 1)

tmp6311 := Call(__e, PrimFunc(symshen_4_5number_6), W2603)


__e.TailApply(tmp6303, tmp6311)
return


}, 1)

tmp6312 := Call(__e, PrimFunc(symshen_4in_1_6), W2602)


__e.TailApply(tmp6302, tmp6312)
return


}


}, 1)

tmp6315 := Call(__e, PrimFunc(symshen_4_5minus_6), V2600)


tmp6316 := Call(__e, tmp6301, tmp6315)


__e.TailApply(tmp6244, tmp6316)
return


}, 1)

tmp6317 := Call(__e, ns2_1set, symshen_4_5number_6, tmp6243)


_ = tmp6317

tmp6318 := MakeNative(func(__e *ControlFlow) {
V2625 := __e.Get(1)
_ = V2625
tmp6319 := MakeNative(func(__e *ControlFlow) {
W2626 := __e.Get(1)
_ = W2626
tmp6321 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2626)


if True == tmp6321 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2626)
return
}


}, 1)

tmp6327 := Call(__e, PrimFunc(symshen_4hds_a_2), V2625, MakeNumber(45))


var ifres6322 Obj

if True == tmp6327 {
tmp6323 := MakeNative(func(__e *ControlFlow) {
W2627 := __e.Get(1)
_ = W2627
__e.TailApply(PrimFunc(symshen_4comb), W2627, symshen_4skip)
return
}, 1)

tmp6324 := Call(__e, PrimFunc(symtail), V2625)


tmp6325 := Call(__e, tmp6323, tmp6324)


ifres6322 = tmp6325


} else {
tmp6326 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6322 = tmp6326


}

__e.TailApply(tmp6319, ifres6322)
return


}, 1)

tmp6328 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp6318)


_ = tmp6328

tmp6329 := MakeNative(func(__e *ControlFlow) {
V2628 := __e.Get(1)
_ = V2628
tmp6330 := MakeNative(func(__e *ControlFlow) {
W2629 := __e.Get(1)
_ = W2629
tmp6332 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2629)


if True == tmp6332 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2629)
return
}


}, 1)

tmp6338 := Call(__e, PrimFunc(symshen_4hds_a_2), V2628, MakeNumber(43))


var ifres6333 Obj

if True == tmp6338 {
tmp6334 := MakeNative(func(__e *ControlFlow) {
W2630 := __e.Get(1)
_ = W2630
__e.TailApply(PrimFunc(symshen_4comb), W2630, symshen_4skip)
return
}, 1)

tmp6335 := Call(__e, PrimFunc(symtail), V2628)


tmp6336 := Call(__e, tmp6334, tmp6335)


ifres6333 = tmp6336


} else {
tmp6337 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6333 = tmp6337


}

__e.TailApply(tmp6330, ifres6333)
return


}, 1)

tmp6339 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp6329)


_ = tmp6339

tmp6340 := MakeNative(func(__e *ControlFlow) {
V2631 := __e.Get(1)
_ = V2631
tmp6341 := MakeNative(func(__e *ControlFlow) {
W2632 := __e.Get(1)
_ = W2632
tmp6343 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2632)


if True == tmp6343 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2632)
return
}


}, 1)

tmp6344 := MakeNative(func(__e *ControlFlow) {
W2633 := __e.Get(1)
_ = W2633
tmp6351 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2633)


if True == tmp6351 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6345 := MakeNative(func(__e *ControlFlow) {
W2634 := __e.Get(1)
_ = W2634
tmp6346 := MakeNative(func(__e *ControlFlow) {
W2635 := __e.Get(1)
_ = W2635
tmp6347 := Call(__e, PrimFunc(symshen_4compute_1integer), W2634)


__e.TailApply(PrimFunc(symshen_4comb), W2635, tmp6347)
return


}, 1)

tmp6348 := Call(__e, PrimFunc(symshen_4in_1_6), W2633)


__e.TailApply(tmp6346, tmp6348)
return


}, 1)

tmp6349 := Call(__e, PrimFunc(symshen_4_5_1out), W2633)


__e.TailApply(tmp6345, tmp6349)
return


}


}, 1)

tmp6352 := Call(__e, PrimFunc(symshen_4_5digits_6), V2631)


tmp6353 := Call(__e, tmp6344, tmp6352)


__e.TailApply(tmp6341, tmp6353)
return


}, 1)

tmp6354 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp6340)


_ = tmp6354

tmp6355 := MakeNative(func(__e *ControlFlow) {
V2636 := __e.Get(1)
_ = V2636
tmp6356 := MakeNative(func(__e *ControlFlow) {
W2637 := __e.Get(1)
_ = W2637
tmp6371 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2637)


if True == tmp6371 {
tmp6357 := MakeNative(func(__e *ControlFlow) {
W2644 := __e.Get(1)
_ = W2644
tmp6359 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2644)


if True == tmp6359 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2644)
return
}


}, 1)

tmp6360 := MakeNative(func(__e *ControlFlow) {
W2645 := __e.Get(1)
_ = W2645
tmp6367 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2645)


if True == tmp6367 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6361 := MakeNative(func(__e *ControlFlow) {
W2646 := __e.Get(1)
_ = W2646
tmp6362 := MakeNative(func(__e *ControlFlow) {
W2647 := __e.Get(1)
_ = W2647
tmp6363 := PrimCons(W2646, Nil)

__e.TailApply(PrimFunc(symshen_4comb), W2647, tmp6363)
return


}, 1)

tmp6364 := Call(__e, PrimFunc(symshen_4in_1_6), W2645)


__e.TailApply(tmp6362, tmp6364)
return


}, 1)

tmp6365 := Call(__e, PrimFunc(symshen_4_5_1out), W2645)


__e.TailApply(tmp6361, tmp6365)
return


}


}, 1)

tmp6368 := Call(__e, PrimFunc(symshen_4_5digit_6), V2636)


tmp6369 := Call(__e, tmp6360, tmp6368)


__e.TailApply(tmp6357, tmp6369)
return


} else {
__e.Return(W2637)
return
}


}, 1)

tmp6372 := MakeNative(func(__e *ControlFlow) {
W2638 := __e.Get(1)
_ = W2638
tmp6387 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2638)


if True == tmp6387 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6373 := MakeNative(func(__e *ControlFlow) {
W2639 := __e.Get(1)
_ = W2639
tmp6374 := MakeNative(func(__e *ControlFlow) {
W2640 := __e.Get(1)
_ = W2640
tmp6375 := MakeNative(func(__e *ControlFlow) {
W2641 := __e.Get(1)
_ = W2641
tmp6382 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2641)


if True == tmp6382 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6376 := MakeNative(func(__e *ControlFlow) {
W2642 := __e.Get(1)
_ = W2642
tmp6377 := MakeNative(func(__e *ControlFlow) {
W2643 := __e.Get(1)
_ = W2643
tmp6378 := PrimCons(W2639, W2642)

__e.TailApply(PrimFunc(symshen_4comb), W2643, tmp6378)
return


}, 1)

tmp6379 := Call(__e, PrimFunc(symshen_4in_1_6), W2641)


__e.TailApply(tmp6377, tmp6379)
return


}, 1)

tmp6380 := Call(__e, PrimFunc(symshen_4_5_1out), W2641)


__e.TailApply(tmp6376, tmp6380)
return


}


}, 1)

tmp6383 := Call(__e, PrimFunc(symshen_4_5digits_6), W2640)


__e.TailApply(tmp6375, tmp6383)
return


}, 1)

tmp6384 := Call(__e, PrimFunc(symshen_4in_1_6), W2638)


__e.TailApply(tmp6374, tmp6384)
return


}, 1)

tmp6385 := Call(__e, PrimFunc(symshen_4_5_1out), W2638)


__e.TailApply(tmp6373, tmp6385)
return


}


}, 1)

tmp6388 := Call(__e, PrimFunc(symshen_4_5digit_6), V2636)


tmp6389 := Call(__e, tmp6372, tmp6388)


__e.TailApply(tmp6356, tmp6389)
return


}, 1)

tmp6390 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp6355)


_ = tmp6390

tmp6391 := MakeNative(func(__e *ControlFlow) {
V2648 := __e.Get(1)
_ = V2648
tmp6392 := MakeNative(func(__e *ControlFlow) {
W2649 := __e.Get(1)
_ = W2649
tmp6394 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2649)


if True == tmp6394 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2649)
return
}


}, 1)

tmp6405 := PrimIsPair(V2648)

var ifres6395 Obj

if True == tmp6405 {
tmp6396 := MakeNative(func(__e *ControlFlow) {
W2650 := __e.Get(1)
_ = W2650
tmp6397 := MakeNative(func(__e *ControlFlow) {
W2651 := __e.Get(1)
_ = W2651
tmp6400 := Call(__e, PrimFunc(symshen_4digit_2), W2650)


if True == tmp6400 {
tmp6398 := Call(__e, PrimFunc(symshen_4byte_1_6digit), W2650)


__e.TailApply(PrimFunc(symshen_4comb), W2651, tmp6398)
return


} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6401 := Call(__e, PrimFunc(symtail), V2648)


__e.TailApply(tmp6397, tmp6401)
return


}, 1)

tmp6402 := Call(__e, PrimFunc(symhead), V2648)


tmp6403 := Call(__e, tmp6396, tmp6402)


ifres6395 = tmp6403


} else {
tmp6404 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6395 = tmp6404


}

__e.TailApply(tmp6392, ifres6395)
return


}, 1)

tmp6406 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp6391)


_ = tmp6406

tmp6407 := MakeNative(func(__e *ControlFlow) {
V2652 := __e.Get(1)
_ = V2652
__e.Return(PrimNumberSubtract(V2652, MakeNumber(48)))
return
}, 1)

tmp6408 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp6407)


_ = tmp6408

tmp6409 := MakeNative(func(__e *ControlFlow) {
V2653 := __e.Get(1)
_ = V2653
tmp6410 := Call(__e, PrimFunc(symreverse), V2653)


__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp6410, MakeNumber(0))
return


}, 1)

tmp6411 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp6409)


_ = tmp6411

tmp6412 := MakeNative(func(__e *ControlFlow) {
V2656 := __e.Get(1)
_ = V2656
V2657 := __e.Get(2)
_ = V2657
tmp6422 := PrimEqual(Nil, V2656)

if True == tmp6422 {
__e.Return(MakeNumber(0))
return
} else {
tmp6420 := PrimIsPair(V2656)

if True == tmp6420 {
tmp6413 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2657)


tmp6414 := PrimHead(V2656)

tmp6415 := PrimNumberMultiply(tmp6413, tmp6414)

tmp6416 := PrimTail(V2656)

tmp6417 := PrimNumberAdd(V2657, MakeNumber(1))

tmp6418 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp6416, tmp6417)


__e.Return(PrimNumberAdd(tmp6415, tmp6418))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.compute-integer-h")))
return
}


}


}, 2)

tmp6423 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp6412)


_ = tmp6423

tmp6424 := MakeNative(func(__e *ControlFlow) {
V2660 := __e.Get(1)
_ = V2660
V2661 := __e.Get(2)
_ = V2661
tmp6432 := PrimEqual(MakeNumber(0), V2661)

if True == tmp6432 {
__e.Return(MakeNumber(1))
return
} else {
tmp6430 := PrimGreatThan(V2661, MakeNumber(0))

if True == tmp6430 {
tmp6425 := PrimNumberSubtract(V2661, MakeNumber(1))

tmp6426 := Call(__e, PrimFunc(symshen_4expt), V2660, tmp6425)


__e.Return(PrimNumberMultiply(V2660, tmp6426))
return


} else {
tmp6427 := PrimNumberAdd(V2661, MakeNumber(1))

tmp6428 := Call(__e, PrimFunc(symshen_4expt), V2660, tmp6427)


__e.Return(PrimNumberDivide(tmp6428, V2660))
return


}


}


}, 2)

tmp6433 := Call(__e, ns2_1set, symshen_4expt, tmp6424)


_ = tmp6433

tmp6434 := MakeNative(func(__e *ControlFlow) {
V2662 := __e.Get(1)
_ = V2662
tmp6435 := MakeNative(func(__e *ControlFlow) {
W2663 := __e.Get(1)
_ = W2663
tmp6455 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2663)


if True == tmp6455 {
tmp6436 := MakeNative(func(__e *ControlFlow) {
W2672 := __e.Get(1)
_ = W2672
tmp6438 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2672)


if True == tmp6438 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2672)
return
}


}, 1)

tmp6439 := MakeNative(func(__e *ControlFlow) {
W2673 := __e.Get(1)
_ = W2673
tmp6451 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2673)


if True == tmp6451 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6440 := MakeNative(func(__e *ControlFlow) {
W2674 := __e.Get(1)
_ = W2674
tmp6441 := MakeNative(func(__e *ControlFlow) {
W2675 := __e.Get(1)
_ = W2675
tmp6447 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2675)


if True == tmp6447 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6442 := MakeNative(func(__e *ControlFlow) {
W2676 := __e.Get(1)
_ = W2676
tmp6443 := MakeNative(func(__e *ControlFlow) {
W2677 := __e.Get(1)
_ = W2677
__e.TailApply(PrimFunc(symshen_4comb), W2677, W2676)
return
}, 1)

tmp6444 := Call(__e, PrimFunc(symshen_4in_1_6), W2675)


__e.TailApply(tmp6443, tmp6444)
return


}, 1)

tmp6445 := Call(__e, PrimFunc(symshen_4_5_1out), W2675)


__e.TailApply(tmp6442, tmp6445)
return


}


}, 1)

tmp6448 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2674)


__e.TailApply(tmp6441, tmp6448)
return


}, 1)

tmp6449 := Call(__e, PrimFunc(symshen_4in_1_6), W2673)


__e.TailApply(tmp6440, tmp6449)
return


}


}, 1)

tmp6452 := Call(__e, PrimFunc(symshen_4_5stop_6), V2662)


tmp6453 := Call(__e, tmp6439, tmp6452)


__e.TailApply(tmp6436, tmp6453)
return


} else {
__e.Return(W2663)
return
}


}, 1)

tmp6456 := MakeNative(func(__e *ControlFlow) {
W2664 := __e.Get(1)
_ = W2664
tmp6477 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2664)


if True == tmp6477 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6457 := MakeNative(func(__e *ControlFlow) {
W2665 := __e.Get(1)
_ = W2665
tmp6458 := MakeNative(func(__e *ControlFlow) {
W2666 := __e.Get(1)
_ = W2666
tmp6459 := MakeNative(func(__e *ControlFlow) {
W2667 := __e.Get(1)
_ = W2667
tmp6472 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2667)


if True == tmp6472 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6460 := MakeNative(func(__e *ControlFlow) {
W2668 := __e.Get(1)
_ = W2668
tmp6461 := MakeNative(func(__e *ControlFlow) {
W2669 := __e.Get(1)
_ = W2669
tmp6468 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2669)


if True == tmp6468 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6462 := MakeNative(func(__e *ControlFlow) {
W2670 := __e.Get(1)
_ = W2670
tmp6463 := MakeNative(func(__e *ControlFlow) {
W2671 := __e.Get(1)
_ = W2671
tmp6464 := PrimNumberAdd(W2665, W2670)

__e.TailApply(PrimFunc(symshen_4comb), W2671, tmp6464)
return


}, 1)

tmp6465 := Call(__e, PrimFunc(symshen_4in_1_6), W2669)


__e.TailApply(tmp6463, tmp6465)
return


}, 1)

tmp6466 := Call(__e, PrimFunc(symshen_4_5_1out), W2669)


__e.TailApply(tmp6462, tmp6466)
return


}


}, 1)

tmp6469 := Call(__e, PrimFunc(symshen_4_5fraction_6), W2668)


__e.TailApply(tmp6461, tmp6469)
return


}, 1)

tmp6470 := Call(__e, PrimFunc(symshen_4in_1_6), W2667)


__e.TailApply(tmp6460, tmp6470)
return


}


}, 1)

tmp6473 := Call(__e, PrimFunc(symshen_4_5stop_6), W2666)


__e.TailApply(tmp6459, tmp6473)
return


}, 1)

tmp6474 := Call(__e, PrimFunc(symshen_4in_1_6), W2664)


__e.TailApply(tmp6458, tmp6474)
return


}, 1)

tmp6475 := Call(__e, PrimFunc(symshen_4_5_1out), W2664)


__e.TailApply(tmp6457, tmp6475)
return


}


}, 1)

tmp6478 := Call(__e, PrimFunc(symshen_4_5integer_6), V2662)


tmp6479 := Call(__e, tmp6456, tmp6478)


__e.TailApply(tmp6435, tmp6479)
return


}, 1)

tmp6480 := Call(__e, ns2_1set, symshen_4_5float_6, tmp6434)


_ = tmp6480

tmp6481 := MakeNative(func(__e *ControlFlow) {
V2678 := __e.Get(1)
_ = V2678
tmp6482 := MakeNative(func(__e *ControlFlow) {
W2679 := __e.Get(1)
_ = W2679
tmp6484 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2679)


if True == tmp6484 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2679)
return
}


}, 1)

tmp6490 := Call(__e, PrimFunc(symshen_4hds_a_2), V2678, MakeNumber(46))


var ifres6485 Obj

if True == tmp6490 {
tmp6486 := MakeNative(func(__e *ControlFlow) {
W2680 := __e.Get(1)
_ = W2680
__e.TailApply(PrimFunc(symshen_4comb), W2680, symshen_4skip)
return
}, 1)

tmp6487 := Call(__e, PrimFunc(symtail), V2678)


tmp6488 := Call(__e, tmp6486, tmp6487)


ifres6485 = tmp6488


} else {
tmp6489 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6485 = tmp6489


}

__e.TailApply(tmp6482, ifres6485)
return


}, 1)

tmp6491 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp6481)


_ = tmp6491

tmp6492 := MakeNative(func(__e *ControlFlow) {
V2681 := __e.Get(1)
_ = V2681
tmp6493 := MakeNative(func(__e *ControlFlow) {
W2682 := __e.Get(1)
_ = W2682
tmp6495 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2682)


if True == tmp6495 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2682)
return
}


}, 1)

tmp6496 := MakeNative(func(__e *ControlFlow) {
W2683 := __e.Get(1)
_ = W2683
tmp6503 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2683)


if True == tmp6503 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6497 := MakeNative(func(__e *ControlFlow) {
W2684 := __e.Get(1)
_ = W2684
tmp6498 := MakeNative(func(__e *ControlFlow) {
W2685 := __e.Get(1)
_ = W2685
tmp6499 := Call(__e, PrimFunc(symshen_4compute_1fraction), W2684)


__e.TailApply(PrimFunc(symshen_4comb), W2685, tmp6499)
return


}, 1)

tmp6500 := Call(__e, PrimFunc(symshen_4in_1_6), W2683)


__e.TailApply(tmp6498, tmp6500)
return


}, 1)

tmp6501 := Call(__e, PrimFunc(symshen_4_5_1out), W2683)


__e.TailApply(tmp6497, tmp6501)
return


}


}, 1)

tmp6504 := Call(__e, PrimFunc(symshen_4_5digits_6), V2681)


tmp6505 := Call(__e, tmp6496, tmp6504)


__e.TailApply(tmp6493, tmp6505)
return


}, 1)

tmp6506 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp6492)


_ = tmp6506

tmp6507 := MakeNative(func(__e *ControlFlow) {
V2686 := __e.Get(1)
_ = V2686
__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V2686, MakeNumber(-1))
return
}, 1)

tmp6508 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp6507)


_ = tmp6508

tmp6509 := MakeNative(func(__e *ControlFlow) {
V2689 := __e.Get(1)
_ = V2689
V2690 := __e.Get(2)
_ = V2690
tmp6519 := PrimEqual(Nil, V2689)

if True == tmp6519 {
__e.Return(MakeNumber(0))
return
} else {
tmp6517 := PrimIsPair(V2689)

if True == tmp6517 {
tmp6510 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2690)


tmp6511 := PrimHead(V2689)

tmp6512 := PrimNumberMultiply(tmp6510, tmp6511)

tmp6513 := PrimTail(V2689)

tmp6514 := PrimNumberSubtract(V2690, MakeNumber(1))

tmp6515 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp6513, tmp6514)


__e.Return(PrimNumberAdd(tmp6512, tmp6515))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.compute-fraction-h")))
return
}


}


}, 2)

tmp6520 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp6509)


_ = tmp6520

tmp6521 := MakeNative(func(__e *ControlFlow) {
V2691 := __e.Get(1)
_ = V2691
tmp6522 := MakeNative(func(__e *ControlFlow) {
W2692 := __e.Get(1)
_ = W2692
tmp6551 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2692)


if True == tmp6551 {
tmp6523 := MakeNative(func(__e *ControlFlow) {
W2701 := __e.Get(1)
_ = W2701
tmp6525 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2701)


if True == tmp6525 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2701)
return
}


}, 1)

tmp6526 := MakeNative(func(__e *ControlFlow) {
W2702 := __e.Get(1)
_ = W2702
tmp6547 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2702)


if True == tmp6547 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6527 := MakeNative(func(__e *ControlFlow) {
W2703 := __e.Get(1)
_ = W2703
tmp6528 := MakeNative(func(__e *ControlFlow) {
W2704 := __e.Get(1)
_ = W2704
tmp6529 := MakeNative(func(__e *ControlFlow) {
W2705 := __e.Get(1)
_ = W2705
tmp6542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2705)


if True == tmp6542 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6530 := MakeNative(func(__e *ControlFlow) {
W2706 := __e.Get(1)
_ = W2706
tmp6531 := MakeNative(func(__e *ControlFlow) {
W2707 := __e.Get(1)
_ = W2707
tmp6538 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2707)


if True == tmp6538 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6532 := MakeNative(func(__e *ControlFlow) {
W2708 := __e.Get(1)
_ = W2708
tmp6533 := MakeNative(func(__e *ControlFlow) {
W2709 := __e.Get(1)
_ = W2709
tmp6534 := Call(__e, PrimFunc(symshen_4compute_1E), W2703, W2708)


__e.TailApply(PrimFunc(symshen_4comb), W2709, tmp6534)
return


}, 1)

tmp6535 := Call(__e, PrimFunc(symshen_4in_1_6), W2707)


__e.TailApply(tmp6533, tmp6535)
return


}, 1)

tmp6536 := Call(__e, PrimFunc(symshen_4_5_1out), W2707)


__e.TailApply(tmp6532, tmp6536)
return


}


}, 1)

tmp6539 := Call(__e, PrimFunc(symshen_4_5log10_6), W2706)


__e.TailApply(tmp6531, tmp6539)
return


}, 1)

tmp6540 := Call(__e, PrimFunc(symshen_4in_1_6), W2705)


__e.TailApply(tmp6530, tmp6540)
return


}


}, 1)

tmp6543 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2704)


__e.TailApply(tmp6529, tmp6543)
return


}, 1)

tmp6544 := Call(__e, PrimFunc(symshen_4in_1_6), W2702)


__e.TailApply(tmp6528, tmp6544)
return


}, 1)

tmp6545 := Call(__e, PrimFunc(symshen_4_5_1out), W2702)


__e.TailApply(tmp6527, tmp6545)
return


}


}, 1)

tmp6548 := Call(__e, PrimFunc(symshen_4_5integer_6), V2691)


tmp6549 := Call(__e, tmp6526, tmp6548)


__e.TailApply(tmp6523, tmp6549)
return


} else {
__e.Return(W2692)
return
}


}, 1)

tmp6552 := MakeNative(func(__e *ControlFlow) {
W2693 := __e.Get(1)
_ = W2693
tmp6573 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2693)


if True == tmp6573 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6553 := MakeNative(func(__e *ControlFlow) {
W2694 := __e.Get(1)
_ = W2694
tmp6554 := MakeNative(func(__e *ControlFlow) {
W2695 := __e.Get(1)
_ = W2695
tmp6555 := MakeNative(func(__e *ControlFlow) {
W2696 := __e.Get(1)
_ = W2696
tmp6568 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2696)


if True == tmp6568 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6556 := MakeNative(func(__e *ControlFlow) {
W2697 := __e.Get(1)
_ = W2697
tmp6557 := MakeNative(func(__e *ControlFlow) {
W2698 := __e.Get(1)
_ = W2698
tmp6564 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2698)


if True == tmp6564 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6558 := MakeNative(func(__e *ControlFlow) {
W2699 := __e.Get(1)
_ = W2699
tmp6559 := MakeNative(func(__e *ControlFlow) {
W2700 := __e.Get(1)
_ = W2700
tmp6560 := Call(__e, PrimFunc(symshen_4compute_1E), W2694, W2699)


__e.TailApply(PrimFunc(symshen_4comb), W2700, tmp6560)
return


}, 1)

tmp6561 := Call(__e, PrimFunc(symshen_4in_1_6), W2698)


__e.TailApply(tmp6559, tmp6561)
return


}, 1)

tmp6562 := Call(__e, PrimFunc(symshen_4_5_1out), W2698)


__e.TailApply(tmp6558, tmp6562)
return


}


}, 1)

tmp6565 := Call(__e, PrimFunc(symshen_4_5log10_6), W2697)


__e.TailApply(tmp6557, tmp6565)
return


}, 1)

tmp6566 := Call(__e, PrimFunc(symshen_4in_1_6), W2696)


__e.TailApply(tmp6556, tmp6566)
return


}


}, 1)

tmp6569 := Call(__e, PrimFunc(symshen_4_5lowE_6), W2695)


__e.TailApply(tmp6555, tmp6569)
return


}, 1)

tmp6570 := Call(__e, PrimFunc(symshen_4in_1_6), W2693)


__e.TailApply(tmp6554, tmp6570)
return


}, 1)

tmp6571 := Call(__e, PrimFunc(symshen_4_5_1out), W2693)


__e.TailApply(tmp6553, tmp6571)
return


}


}, 1)

tmp6574 := Call(__e, PrimFunc(symshen_4_5float_6), V2691)


tmp6575 := Call(__e, tmp6552, tmp6574)


__e.TailApply(tmp6522, tmp6575)
return


}, 1)

tmp6576 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp6521)


_ = tmp6576

tmp6577 := MakeNative(func(__e *ControlFlow) {
V2710 := __e.Get(1)
_ = V2710
tmp6578 := MakeNative(func(__e *ControlFlow) {
W2711 := __e.Get(1)
_ = W2711
tmp6611 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2711)


if True == tmp6611 {
tmp6579 := MakeNative(func(__e *ControlFlow) {
W2717 := __e.Get(1)
_ = W2717
tmp6593 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2717)


if True == tmp6593 {
tmp6580 := MakeNative(func(__e *ControlFlow) {
W2723 := __e.Get(1)
_ = W2723
tmp6582 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2723)


if True == tmp6582 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2723)
return
}


}, 1)

tmp6583 := MakeNative(func(__e *ControlFlow) {
W2724 := __e.Get(1)
_ = W2724
tmp6589 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2724)


if True == tmp6589 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6584 := MakeNative(func(__e *ControlFlow) {
W2725 := __e.Get(1)
_ = W2725
tmp6585 := MakeNative(func(__e *ControlFlow) {
W2726 := __e.Get(1)
_ = W2726
__e.TailApply(PrimFunc(symshen_4comb), W2726, W2725)
return
}, 1)

tmp6586 := Call(__e, PrimFunc(symshen_4in_1_6), W2724)


__e.TailApply(tmp6585, tmp6586)
return


}, 1)

tmp6587 := Call(__e, PrimFunc(symshen_4_5_1out), W2724)


__e.TailApply(tmp6584, tmp6587)
return


}


}, 1)

tmp6590 := Call(__e, PrimFunc(symshen_4_5integer_6), V2710)


tmp6591 := Call(__e, tmp6583, tmp6590)


__e.TailApply(tmp6580, tmp6591)
return


} else {
__e.Return(W2717)
return
}


}, 1)

tmp6594 := MakeNative(func(__e *ControlFlow) {
W2718 := __e.Get(1)
_ = W2718
tmp6607 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2718)


if True == tmp6607 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6595 := MakeNative(func(__e *ControlFlow) {
W2719 := __e.Get(1)
_ = W2719
tmp6596 := MakeNative(func(__e *ControlFlow) {
W2720 := __e.Get(1)
_ = W2720
tmp6603 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2720)


if True == tmp6603 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6597 := MakeNative(func(__e *ControlFlow) {
W2721 := __e.Get(1)
_ = W2721
tmp6598 := MakeNative(func(__e *ControlFlow) {
W2722 := __e.Get(1)
_ = W2722
tmp6599 := PrimNumberSubtract(MakeNumber(0), W2721)

__e.TailApply(PrimFunc(symshen_4comb), W2722, tmp6599)
return


}, 1)

tmp6600 := Call(__e, PrimFunc(symshen_4in_1_6), W2720)


__e.TailApply(tmp6598, tmp6600)
return


}, 1)

tmp6601 := Call(__e, PrimFunc(symshen_4_5_1out), W2720)


__e.TailApply(tmp6597, tmp6601)
return


}


}, 1)

tmp6604 := Call(__e, PrimFunc(symshen_4_5log10_6), W2719)


__e.TailApply(tmp6596, tmp6604)
return


}, 1)

tmp6605 := Call(__e, PrimFunc(symshen_4in_1_6), W2718)


__e.TailApply(tmp6595, tmp6605)
return


}


}, 1)

tmp6608 := Call(__e, PrimFunc(symshen_4_5minus_6), V2710)


tmp6609 := Call(__e, tmp6594, tmp6608)


__e.TailApply(tmp6579, tmp6609)
return


} else {
__e.Return(W2711)
return
}


}, 1)

tmp6612 := MakeNative(func(__e *ControlFlow) {
W2712 := __e.Get(1)
_ = W2712
tmp6624 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2712)


if True == tmp6624 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6613 := MakeNative(func(__e *ControlFlow) {
W2713 := __e.Get(1)
_ = W2713
tmp6614 := MakeNative(func(__e *ControlFlow) {
W2714 := __e.Get(1)
_ = W2714
tmp6620 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2714)


if True == tmp6620 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6615 := MakeNative(func(__e *ControlFlow) {
W2715 := __e.Get(1)
_ = W2715
tmp6616 := MakeNative(func(__e *ControlFlow) {
W2716 := __e.Get(1)
_ = W2716
__e.TailApply(PrimFunc(symshen_4comb), W2716, W2715)
return
}, 1)

tmp6617 := Call(__e, PrimFunc(symshen_4in_1_6), W2714)


__e.TailApply(tmp6616, tmp6617)
return


}, 1)

tmp6618 := Call(__e, PrimFunc(symshen_4_5_1out), W2714)


__e.TailApply(tmp6615, tmp6618)
return


}


}, 1)

tmp6621 := Call(__e, PrimFunc(symshen_4_5log10_6), W2713)


__e.TailApply(tmp6614, tmp6621)
return


}, 1)

tmp6622 := Call(__e, PrimFunc(symshen_4in_1_6), W2712)


__e.TailApply(tmp6613, tmp6622)
return


}


}, 1)

tmp6625 := Call(__e, PrimFunc(symshen_4_5plus_6), V2710)


tmp6626 := Call(__e, tmp6612, tmp6625)


__e.TailApply(tmp6578, tmp6626)
return


}, 1)

tmp6627 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp6577)


_ = tmp6627

tmp6628 := MakeNative(func(__e *ControlFlow) {
V2727 := __e.Get(1)
_ = V2727
tmp6629 := MakeNative(func(__e *ControlFlow) {
W2728 := __e.Get(1)
_ = W2728
tmp6631 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2728)


if True == tmp6631 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2728)
return
}


}, 1)

tmp6637 := Call(__e, PrimFunc(symshen_4hds_a_2), V2727, MakeNumber(101))


var ifres6632 Obj

if True == tmp6637 {
tmp6633 := MakeNative(func(__e *ControlFlow) {
W2729 := __e.Get(1)
_ = W2729
__e.TailApply(PrimFunc(symshen_4comb), W2729, symshen_4skip)
return
}, 1)

tmp6634 := Call(__e, PrimFunc(symtail), V2727)


tmp6635 := Call(__e, tmp6633, tmp6634)


ifres6632 = tmp6635


} else {
tmp6636 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6632 = tmp6636


}

__e.TailApply(tmp6629, ifres6632)
return


}, 1)

tmp6638 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp6628)


_ = tmp6638

tmp6639 := MakeNative(func(__e *ControlFlow) {
V2730 := __e.Get(1)
_ = V2730
V2731 := __e.Get(2)
_ = V2731
tmp6640 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2731)


__e.Return(PrimNumberMultiply(V2730, tmp6640))
return


}, 2)

tmp6641 := Call(__e, ns2_1set, symshen_4compute_1E, tmp6639)


_ = tmp6641

tmp6642 := MakeNative(func(__e *ControlFlow) {
V2732 := __e.Get(1)
_ = V2732
tmp6643 := MakeNative(func(__e *ControlFlow) {
W2733 := __e.Get(1)
_ = W2733
tmp6655 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2733)


if True == tmp6655 {
tmp6644 := MakeNative(func(__e *ControlFlow) {
W2738 := __e.Get(1)
_ = W2738
tmp6646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2738)


if True == tmp6646 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2738)
return
}


}, 1)

tmp6647 := MakeNative(func(__e *ControlFlow) {
W2739 := __e.Get(1)
_ = W2739
tmp6651 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2739)


if True == tmp6651 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6648 := MakeNative(func(__e *ControlFlow) {
W2740 := __e.Get(1)
_ = W2740
__e.TailApply(PrimFunc(symshen_4comb), W2740, symshen_4skip)
return
}, 1)

tmp6649 := Call(__e, PrimFunc(symshen_4in_1_6), W2739)


__e.TailApply(tmp6648, tmp6649)
return


}


}, 1)

tmp6652 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2732)


tmp6653 := Call(__e, tmp6647, tmp6652)


__e.TailApply(tmp6644, tmp6653)
return


} else {
__e.Return(W2733)
return
}


}, 1)

tmp6656 := MakeNative(func(__e *ControlFlow) {
W2734 := __e.Get(1)
_ = W2734
tmp6666 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2734)


if True == tmp6666 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6657 := MakeNative(func(__e *ControlFlow) {
W2735 := __e.Get(1)
_ = W2735
tmp6658 := MakeNative(func(__e *ControlFlow) {
W2736 := __e.Get(1)
_ = W2736
tmp6662 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2736)


if True == tmp6662 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
tmp6659 := MakeNative(func(__e *ControlFlow) {
W2737 := __e.Get(1)
_ = W2737
__e.TailApply(PrimFunc(symshen_4comb), W2737, symshen_4skip)
return
}, 1)

tmp6660 := Call(__e, PrimFunc(symshen_4in_1_6), W2736)


__e.TailApply(tmp6659, tmp6660)
return


}


}, 1)

tmp6663 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), W2735)


__e.TailApply(tmp6658, tmp6663)
return


}, 1)

tmp6664 := Call(__e, PrimFunc(symshen_4in_1_6), W2734)


__e.TailApply(tmp6657, tmp6664)
return


}


}, 1)

tmp6667 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2732)


tmp6668 := Call(__e, tmp6656, tmp6667)


__e.TailApply(tmp6643, tmp6668)
return


}, 1)

tmp6669 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6642)


_ = tmp6669

tmp6670 := MakeNative(func(__e *ControlFlow) {
V2741 := __e.Get(1)
_ = V2741
tmp6671 := MakeNative(func(__e *ControlFlow) {
W2742 := __e.Get(1)
_ = W2742
tmp6673 := Call(__e, PrimFunc(symshen_4parse_1failure_2), W2742)


if True == tmp6673 {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
} else {
__e.Return(W2742)
return
}


}, 1)

tmp6683 := PrimIsPair(V2741)

var ifres6674 Obj

if True == tmp6683 {
tmp6675 := MakeNative(func(__e *ControlFlow) {
W2743 := __e.Get(1)
_ = W2743
tmp6676 := MakeNative(func(__e *ControlFlow) {
W2744 := __e.Get(1)
_ = W2744
tmp6678 := Call(__e, PrimFunc(symshen_4whitespace_2), W2743)


if True == tmp6678 {
__e.TailApply(PrimFunc(symshen_4comb), W2744, symshen_4skip)
return
} else {
__e.TailApply(PrimFunc(symshen_4parse_1failure))
return
}


}, 1)

tmp6679 := Call(__e, PrimFunc(symtail), V2741)


__e.TailApply(tmp6676, tmp6679)
return


}, 1)

tmp6680 := Call(__e, PrimFunc(symhead), V2741)


tmp6681 := Call(__e, tmp6675, tmp6680)


ifres6674 = tmp6681


} else {
tmp6682 := Call(__e, PrimFunc(symshen_4parse_1failure))


ifres6674 = tmp6682


}

__e.TailApply(tmp6671, ifres6674)
return


}, 1)

tmp6684 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp6670)


_ = tmp6684

tmp6685 := MakeNative(func(__e *ControlFlow) {
V2747 := __e.Get(1)
_ = V2747
tmp6693 := PrimEqual(MakeNumber(32), V2747)

if True == tmp6693 {
__e.Return(True)
return
} else {
tmp6691 := PrimEqual(MakeNumber(13), V2747)

if True == tmp6691 {
__e.Return(True)
return
} else {
tmp6689 := PrimEqual(MakeNumber(10), V2747)

if True == tmp6689 {
__e.Return(True)
return
} else {
tmp6687 := PrimEqual(MakeNumber(9), V2747)

if True == tmp6687 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}, 1)

tmp6694 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp6685)


_ = tmp6694

tmp6695 := MakeNative(func(__e *ControlFlow) {
V2748 := __e.Get(1)
_ = V2748
tmp6718 := PrimEqual(Nil, V2748)

if True == tmp6718 {
__e.Return(Nil)
return
} else {
tmp6716 := PrimIsPair(V2748)

var ifres6712 Obj

if True == tmp6716 {
tmp6714 := PrimHead(V2748)

tmp6715 := Call(__e, PrimFunc(symshen_4packaged_2), tmp6714)


var ifres6713 Obj

if True == tmp6715 {
ifres6713 = True


} else {
ifres6713 = False


}

ifres6712 = ifres6713


} else {
ifres6712 = False


}

if True == ifres6712 {
tmp6696 := PrimHead(V2748)

tmp6697 := Call(__e, PrimFunc(symshen_4unpackage), tmp6696)


tmp6698 := PrimTail(V2748)

tmp6699 := Call(__e, PrimFunc(symappend), tmp6697, tmp6698)


__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6699)
return


} else {
tmp6710 := PrimIsPair(V2748)

if True == tmp6710 {
tmp6700 := MakeNative(func(__e *ControlFlow) {
W2749 := __e.Get(1)
_ = W2749
tmp6706 := Call(__e, PrimFunc(symshen_4packaged_2), W2749)


if True == tmp6706 {
tmp6701 := PrimTail(V2748)

tmp6702 := PrimCons(W2749, tmp6701)

__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6702)
return


} else {
tmp6703 := PrimTail(V2748)

tmp6704 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp6703)


__e.Return(PrimCons(W2749, tmp6704))
return


}


}, 1)

tmp6707 := PrimHead(V2748)

tmp6708 := Call(__e, PrimFunc(symmacroexpand), tmp6707)


__e.TailApply(tmp6700, tmp6708)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpackage&macroexpand")))
return
}


}


}


}, 1)

tmp6719 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp6695)


_ = tmp6719

tmp6720 := MakeNative(func(__e *ControlFlow) {
V2752 := __e.Get(1)
_ = V2752
tmp6735 := PrimIsPair(V2752)

var ifres6722 Obj

if True == tmp6735 {
tmp6733 := PrimHead(V2752)

tmp6734 := PrimEqual(sympackage, tmp6733)

var ifres6724 Obj

if True == tmp6734 {
tmp6731 := PrimTail(V2752)

tmp6732 := PrimIsPair(tmp6731)

var ifres6726 Obj

if True == tmp6732 {
tmp6728 := PrimTail(V2752)

tmp6729 := PrimTail(tmp6728)

tmp6730 := PrimIsPair(tmp6729)

var ifres6727 Obj

if True == tmp6730 {
ifres6727 = True


} else {
ifres6727 = False


}

ifres6726 = ifres6727


} else {
ifres6726 = False


}

var ifres6725 Obj

if True == ifres6726 {
ifres6725 = True


} else {
ifres6725 = False


}

ifres6724 = ifres6725


} else {
ifres6724 = False


}

var ifres6723 Obj

if True == ifres6724 {
ifres6723 = True


} else {
ifres6723 = False


}

ifres6722 = ifres6723


} else {
ifres6722 = False


}

if True == ifres6722 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6736 := Call(__e, ns2_1set, symshen_4packaged_2, tmp6720)


_ = tmp6736

tmp6737 := MakeNative(func(__e *ControlFlow) {
V2755 := __e.Get(1)
_ = V2755
tmp6798 := PrimIsPair(V2755)

var ifres6780 Obj

if True == tmp6798 {
tmp6796 := PrimHead(V2755)

tmp6797 := PrimEqual(sympackage, tmp6796)

var ifres6782 Obj

if True == tmp6797 {
tmp6794 := PrimTail(V2755)

tmp6795 := PrimIsPair(tmp6794)

var ifres6784 Obj

if True == tmp6795 {
tmp6791 := PrimTail(V2755)

tmp6792 := PrimHead(tmp6791)

tmp6793 := PrimEqual(symnull, tmp6792)

var ifres6786 Obj

if True == tmp6793 {
tmp6788 := PrimTail(V2755)

tmp6789 := PrimTail(tmp6788)

tmp6790 := PrimIsPair(tmp6789)

var ifres6787 Obj

if True == tmp6790 {
ifres6787 = True


} else {
ifres6787 = False


}

ifres6786 = ifres6787


} else {
ifres6786 = False


}

var ifres6785 Obj

if True == ifres6786 {
ifres6785 = True


} else {
ifres6785 = False


}

ifres6784 = ifres6785


} else {
ifres6784 = False


}

var ifres6783 Obj

if True == ifres6784 {
ifres6783 = True


} else {
ifres6783 = False


}

ifres6782 = ifres6783


} else {
ifres6782 = False


}

var ifres6781 Obj

if True == ifres6782 {
ifres6781 = True


} else {
ifres6781 = False


}

ifres6780 = ifres6781


} else {
ifres6780 = False


}

if True == ifres6780 {
tmp6738 := PrimTail(V2755)

tmp6739 := PrimTail(tmp6738)

__e.Return(PrimTail(tmp6739))
return


} else {
tmp6778 := PrimIsPair(V2755)

var ifres6765 Obj

if True == tmp6778 {
tmp6776 := PrimHead(V2755)

tmp6777 := PrimEqual(sympackage, tmp6776)

var ifres6767 Obj

if True == tmp6777 {
tmp6774 := PrimTail(V2755)

tmp6775 := PrimIsPair(tmp6774)

var ifres6769 Obj

if True == tmp6775 {
tmp6771 := PrimTail(V2755)

tmp6772 := PrimTail(tmp6771)

tmp6773 := PrimIsPair(tmp6772)

var ifres6770 Obj

if True == tmp6773 {
ifres6770 = True


} else {
ifres6770 = False


}

ifres6769 = ifres6770


} else {
ifres6769 = False


}

var ifres6768 Obj

if True == ifres6769 {
ifres6768 = True


} else {
ifres6768 = False


}

ifres6767 = ifres6768


} else {
ifres6767 = False


}

var ifres6766 Obj

if True == ifres6767 {
ifres6766 = True


} else {
ifres6766 = False


}

ifres6765 = ifres6766


} else {
ifres6765 = False


}

if True == ifres6765 {
tmp6740 := MakeNative(func(__e *ControlFlow) {
W2756 := __e.Get(1)
_ = W2756
tmp6741 := MakeNative(func(__e *ControlFlow) {
W2757 := __e.Get(1)
_ = W2757
tmp6742 := MakeNative(func(__e *ControlFlow) {
W2758 := __e.Get(1)
_ = W2758
tmp6743 := MakeNative(func(__e *ControlFlow) {
W2759 := __e.Get(1)
_ = W2759
__e.Return(W2757)
return
}, 1)

tmp6744 := PrimTail(V2755)

tmp6745 := PrimHead(tmp6744)

tmp6746 := PrimTail(V2755)

tmp6747 := PrimTail(tmp6746)

tmp6748 := PrimTail(tmp6747)

tmp6749 := Call(__e, PrimFunc(symshen_4record_1internal), tmp6745, W2756, tmp6748)


__e.TailApply(tmp6743, tmp6749)
return


}, 1)

tmp6750 := PrimTail(V2755)

tmp6751 := PrimHead(tmp6750)

tmp6752 := Call(__e, PrimFunc(symshen_4record_1external), tmp6751, W2756)


__e.TailApply(tmp6742, tmp6752)
return


}, 1)

tmp6753 := PrimTail(V2755)

tmp6754 := PrimHead(tmp6753)

tmp6755 := PrimStr(tmp6754)

tmp6756 := PrimTail(V2755)

tmp6757 := PrimTail(tmp6756)

tmp6758 := PrimTail(tmp6757)

tmp6759 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp6755, W2756, tmp6758)


__e.TailApply(tmp6741, tmp6759)
return


}, 1)

tmp6760 := PrimTail(V2755)

tmp6761 := PrimTail(tmp6760)

tmp6762 := PrimHead(tmp6761)

tmp6763 := Call(__e, PrimFunc(symeval), tmp6762)


__e.TailApply(tmp6740, tmp6763)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpackage")))
return
}


}


}, 1)

tmp6799 := Call(__e, ns2_1set, symshen_4unpackage, tmp6737)


_ = tmp6799

tmp6800 := MakeNative(func(__e *ControlFlow) {
V2760 := __e.Get(1)
_ = V2760
V2761 := __e.Get(2)
_ = V2761
V2762 := __e.Get(3)
_ = V2762
tmp6801 := MakeNative(func(__e *ControlFlow) {
W2763 := __e.Get(1)
_ = W2763
tmp6802 := MakeNative(func(__e *ControlFlow) {
W2765 := __e.Get(1)
_ = W2765
tmp6803 := Call(__e, PrimFunc(symunion), W2765, W2763)


tmp6804 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2760, symshen_4internal_1symbols, tmp6803, tmp6804)
return


}, 1)

tmp6805 := PrimStr(V2760)

tmp6806 := Call(__e, PrimFunc(symshen_4internal_1symbols), tmp6805, V2761, V2762)


__e.TailApply(tmp6802, tmp6806)
return


}, 1)

tmp6807 := MakeNative(func(__e *ControlFlow) {
tmp6808 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V2760, symshen_4internal_1symbols, tmp6808)
return


}, 0)

tmp6809 := MakeNative(func(__e *ControlFlow) {
Z2764 := __e.Get(1)
_ = Z2764
__e.Return(Nil)
return
}, 1)

tmp6810 := Call(__e, try_1catch, tmp6807, tmp6809)


__e.TailApply(tmp6801, tmp6810)
return


}, 3)

tmp6811 := Call(__e, ns2_1set, symshen_4record_1internal, tmp6800)


_ = tmp6811

tmp6812 := MakeNative(func(__e *ControlFlow) {
V2772 := __e.Get(1)
_ = V2772
V2773 := __e.Get(2)
_ = V2773
V2774 := __e.Get(3)
_ = V2774
tmp6821 := PrimIsPair(V2774)

if True == tmp6821 {
tmp6813 := PrimHead(V2774)

tmp6814 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2772, V2773, tmp6813)


tmp6815 := PrimTail(V2774)

tmp6816 := Call(__e, PrimFunc(symshen_4internal_1symbols), V2772, V2773, tmp6815)


__e.TailApply(PrimFunc(symunion), tmp6814, tmp6816)
return


} else {
tmp6819 := Call(__e, PrimFunc(symshen_4internal_2), V2774, V2772, V2773)


if True == tmp6819 {
tmp6817 := Call(__e, PrimFunc(symshen_4intern_1in_1package), V2772, V2774)


__e.Return(PrimCons(tmp6817, Nil))
return


} else {
__e.Return(Nil)
return
}


}


}, 3)

tmp6822 := Call(__e, ns2_1set, symshen_4internal_1symbols, tmp6812)


_ = tmp6822

tmp6823 := MakeNative(func(__e *ControlFlow) {
V2775 := __e.Get(1)
_ = V2775
V2776 := __e.Get(2)
_ = V2776
tmp6824 := MakeNative(func(__e *ControlFlow) {
W2777 := __e.Get(1)
_ = W2777
tmp6825 := Call(__e, PrimFunc(symunion), V2776, W2777)


tmp6826 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symput), V2775, symshen_4external_1symbols, tmp6825, tmp6826)
return


}, 1)

tmp6827 := MakeNative(func(__e *ControlFlow) {
tmp6828 := PrimValue(sym_dproperty_1vector_d)

__e.TailApply(PrimFunc(symget), V2775, symshen_4external_1symbols, tmp6828)
return


}, 0)

tmp6829 := MakeNative(func(__e *ControlFlow) {
Z2778 := __e.Get(1)
_ = Z2778
__e.Return(Nil)
return
}, 1)

tmp6830 := Call(__e, try_1catch, tmp6827, tmp6829)


__e.TailApply(tmp6824, tmp6830)
return


}, 2)

tmp6831 := Call(__e, ns2_1set, symshen_4record_1external, tmp6823)


_ = tmp6831

tmp6832 := MakeNative(func(__e *ControlFlow) {
V2783 := __e.Get(1)
_ = V2783
V2784 := __e.Get(2)
_ = V2784
V2785 := __e.Get(3)
_ = V2785
tmp6837 := PrimIsPair(V2785)

if True == tmp6837 {
tmp6833 := MakeNative(func(__e *ControlFlow) {
Z2786 := __e.Get(1)
_ = Z2786
__e.TailApply(PrimFunc(symshen_4package_1symbols), V2783, V2784, Z2786)
return
}, 1)

__e.TailApply(PrimFunc(symmap), tmp6833, V2785)
return


} else {
tmp6835 := Call(__e, PrimFunc(symshen_4internal_2), V2785, V2783, V2784)


if True == tmp6835 {
__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V2783, V2785)
return
} else {
__e.Return(V2785)
return
}


}


}, 3)

tmp6838 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp6832)


_ = tmp6838

tmp6839 := MakeNative(func(__e *ControlFlow) {
V2787 := __e.Get(1)
_ = V2787
V2788 := __e.Get(2)
_ = V2788
tmp6840 := PrimStr(V2788)

tmp6841 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp6840)


tmp6842 := Call(__e, PrimFunc(sym_8s), V2787, tmp6841)


__e.Return(PrimIntern(tmp6842))
return


}, 2)

tmp6843 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp6839)


_ = tmp6843

tmp6844 := MakeNative(func(__e *ControlFlow) {
V2789 := __e.Get(1)
_ = V2789
V2790 := __e.Get(2)
_ = V2790
V2791 := __e.Get(3)
_ = V2791
tmp6874 := Call(__e, PrimFunc(symelement_2), V2789, V2791)


tmp6875 := PrimNot(tmp6874)

if True == tmp6875 {
tmp6871 := Call(__e, PrimFunc(symshen_4sng_2), V2789)


tmp6872 := PrimNot(tmp6871)

var ifres6846 Obj

if True == tmp6872 {
tmp6869 := Call(__e, PrimFunc(symshen_4dbl_2), V2789)


tmp6870 := PrimNot(tmp6869)

var ifres6848 Obj

if True == tmp6870 {
tmp6868 := PrimIsSymbol(V2789)

var ifres6850 Obj

if True == tmp6868 {
tmp6866 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2789)


tmp6867 := PrimNot(tmp6866)

var ifres6852 Obj

if True == tmp6867 {
tmp6864 := PrimIsVariable(V2789)

tmp6865 := PrimNot(tmp6864)

var ifres6854 Obj

if True == tmp6865 {
tmp6861 := PrimStr(V2789)

tmp6862 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6861)


tmp6863 := PrimNot(tmp6862)

var ifres6856 Obj

if True == tmp6863 {
tmp6858 := PrimStr(V2789)

tmp6859 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V2790, tmp6858)


tmp6860 := PrimNot(tmp6859)

var ifres6857 Obj

if True == tmp6860 {
ifres6857 = True


} else {
ifres6857 = False


}

ifres6856 = ifres6857


} else {
ifres6856 = False


}

var ifres6855 Obj

if True == ifres6856 {
ifres6855 = True


} else {
ifres6855 = False


}

ifres6854 = ifres6855


} else {
ifres6854 = False


}

var ifres6853 Obj

if True == ifres6854 {
ifres6853 = True


} else {
ifres6853 = False


}

ifres6852 = ifres6853


} else {
ifres6852 = False


}

var ifres6851 Obj

if True == ifres6852 {
ifres6851 = True


} else {
ifres6851 = False


}

ifres6850 = ifres6851


} else {
ifres6850 = False


}

var ifres6849 Obj

if True == ifres6850 {
ifres6849 = True


} else {
ifres6849 = False


}

ifres6848 = ifres6849


} else {
ifres6848 = False


}

var ifres6847 Obj

if True == ifres6848 {
ifres6847 = True


} else {
ifres6847 = False


}

ifres6846 = ifres6847


} else {
ifres6846 = False


}

if True == ifres6846 {
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


}, 3)

tmp6876 := Call(__e, ns2_1set, symshen_4internal_2, tmp6844)


_ = tmp6876

tmp6877 := MakeNative(func(__e *ControlFlow) {
V2796 := __e.Get(1)
_ = V2796
tmp6931 := Call(__e, PrimFunc(symshen_4_7string_2), V2796)


var ifres6879 Obj

if True == tmp6931 {
tmp6929 := Call(__e, PrimFunc(symhdstr), V2796)


tmp6930 := PrimEqual(MakeString("s"), tmp6929)

var ifres6881 Obj

if True == tmp6930 {
tmp6927 := PrimTailString(V2796)

tmp6928 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6927)


var ifres6883 Obj

if True == tmp6928 {
tmp6924 := PrimTailString(V2796)

tmp6925 := Call(__e, PrimFunc(symhdstr), tmp6924)


tmp6926 := PrimEqual(MakeString("h"), tmp6925)

var ifres6885 Obj

if True == tmp6926 {
tmp6921 := PrimTailString(V2796)

tmp6922 := PrimTailString(tmp6921)

tmp6923 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6922)


var ifres6887 Obj

if True == tmp6923 {
tmp6917 := PrimTailString(V2796)

tmp6918 := PrimTailString(tmp6917)

tmp6919 := Call(__e, PrimFunc(symhdstr), tmp6918)


tmp6920 := PrimEqual(MakeString("e"), tmp6919)

var ifres6889 Obj

if True == tmp6920 {
tmp6913 := PrimTailString(V2796)

tmp6914 := PrimTailString(tmp6913)

tmp6915 := PrimTailString(tmp6914)

tmp6916 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6915)


var ifres6891 Obj

if True == tmp6916 {
tmp6908 := PrimTailString(V2796)

tmp6909 := PrimTailString(tmp6908)

tmp6910 := PrimTailString(tmp6909)

tmp6911 := Call(__e, PrimFunc(symhdstr), tmp6910)


tmp6912 := PrimEqual(MakeString("n"), tmp6911)

var ifres6893 Obj

if True == tmp6912 {
tmp6903 := PrimTailString(V2796)

tmp6904 := PrimTailString(tmp6903)

tmp6905 := PrimTailString(tmp6904)

tmp6906 := PrimTailString(tmp6905)

tmp6907 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6906)


var ifres6895 Obj

if True == tmp6907 {
tmp6897 := PrimTailString(V2796)

tmp6898 := PrimTailString(tmp6897)

tmp6899 := PrimTailString(tmp6898)

tmp6900 := PrimTailString(tmp6899)

tmp6901 := Call(__e, PrimFunc(symhdstr), tmp6900)


tmp6902 := PrimEqual(MakeString("."), tmp6901)

var ifres6896 Obj

if True == tmp6902 {
ifres6896 = True


} else {
ifres6896 = False


}

ifres6895 = ifres6896


} else {
ifres6895 = False


}

var ifres6894 Obj

if True == ifres6895 {
ifres6894 = True


} else {
ifres6894 = False


}

ifres6893 = ifres6894


} else {
ifres6893 = False


}

var ifres6892 Obj

if True == ifres6893 {
ifres6892 = True


} else {
ifres6892 = False


}

ifres6891 = ifres6892


} else {
ifres6891 = False


}

var ifres6890 Obj

if True == ifres6891 {
ifres6890 = True


} else {
ifres6890 = False


}

ifres6889 = ifres6890


} else {
ifres6889 = False


}

var ifres6888 Obj

if True == ifres6889 {
ifres6888 = True


} else {
ifres6888 = False


}

ifres6887 = ifres6888


} else {
ifres6887 = False


}

var ifres6886 Obj

if True == ifres6887 {
ifres6886 = True


} else {
ifres6886 = False


}

ifres6885 = ifres6886


} else {
ifres6885 = False


}

var ifres6884 Obj

if True == ifres6885 {
ifres6884 = True


} else {
ifres6884 = False


}

ifres6883 = ifres6884


} else {
ifres6883 = False


}

var ifres6882 Obj

if True == ifres6883 {
ifres6882 = True


} else {
ifres6882 = False


}

ifres6881 = ifres6882


} else {
ifres6881 = False


}

var ifres6880 Obj

if True == ifres6881 {
ifres6880 = True


} else {
ifres6880 = False


}

ifres6879 = ifres6880


} else {
ifres6879 = False


}

if True == ifres6879 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp6932 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp6877)


_ = tmp6932

tmp6933 := MakeNative(func(__e *ControlFlow) {
V2797 := __e.Get(1)
_ = V2797
tmp6934 := PrimValue(sym_dproperty_1vector_d)

tmp6935 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp6934)


__e.TailApply(PrimFunc(symelement_2), V2797, tmp6935)
return


}, 1)

tmp6936 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp6933)


_ = tmp6936

tmp6937 := MakeNative(func(__e *ControlFlow) {
V2805 := __e.Get(1)
_ = V2805
V2806 := __e.Get(2)
_ = V2806
tmp6958 := PrimEqual(MakeString(""), V2805)

var ifres6951 Obj

if True == tmp6958 {
tmp6957 := Call(__e, PrimFunc(symshen_4_7string_2), V2806)


var ifres6953 Obj

if True == tmp6957 {
tmp6955 := Call(__e, PrimFunc(symhdstr), V2806)


tmp6956 := PrimEqual(MakeString("."), tmp6955)

var ifres6954 Obj

if True == tmp6956 {
ifres6954 = True


} else {
ifres6954 = False


}

ifres6953 = ifres6954


} else {
ifres6953 = False


}

var ifres6952 Obj

if True == ifres6953 {
ifres6952 = True


} else {
ifres6952 = False


}

ifres6951 = ifres6952


} else {
ifres6951 = False


}

if True == ifres6951 {
__e.Return(True)
return
} else {
tmp6949 := Call(__e, PrimFunc(symshen_4_7string_2), V2805)


var ifres6941 Obj

if True == tmp6949 {
tmp6948 := Call(__e, PrimFunc(symshen_4_7string_2), V2806)


var ifres6943 Obj

if True == tmp6948 {
tmp6945 := Call(__e, PrimFunc(symhdstr), V2805)


tmp6946 := Call(__e, PrimFunc(symhdstr), V2806)


tmp6947 := PrimEqual(tmp6945, tmp6946)

var ifres6944 Obj

if True == tmp6947 {
ifres6944 = True


} else {
ifres6944 = False


}

ifres6943 = ifres6944


} else {
ifres6943 = False


}

var ifres6942 Obj

if True == ifres6943 {
ifres6942 = True


} else {
ifres6942 = False


}

ifres6941 = ifres6942


} else {
ifres6941 = False


}

if True == ifres6941 {
tmp6938 := PrimTailString(V2805)

tmp6939 := PrimTailString(V2806)

__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp6938, tmp6939)
return


} else {
__e.Return(False)
return
}


}


}, 2)

tmp6959 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp6937)


_ = tmp6959

tmp6960 := MakeNative(func(__e *ControlFlow) {
V2809 := __e.Get(1)
_ = V2809
V2810 := __e.Get(2)
_ = V2810
tmp6973 := Call(__e, PrimFunc(symelement_2), V2809, V2810)


if True == tmp6973 {
__e.Return(V2809)
return
} else {
tmp6971 := PrimIsPair(V2809)

var ifres6967 Obj

if True == tmp6971 {
tmp6969 := PrimHead(V2809)

tmp6970 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp6969)


var ifres6968 Obj

if True == tmp6970 {
ifres6968 = True


} else {
ifres6968 = False


}

ifres6967 = ifres6968


} else {
ifres6967 = False


}

if True == ifres6967 {
tmp6961 := PrimHead(V2809)

__e.TailApply(PrimFunc(symshen_4special_1case), tmp6961, V2809, V2810)
return


} else {
tmp6965 := PrimIsPair(V2809)

if True == tmp6965 {
tmp6962 := MakeNative(func(__e *ControlFlow) {
Z2811 := __e.Get(1)
_ = Z2811
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2811, V2810)
return
}, 1)

tmp6963 := Call(__e, PrimFunc(symmap), tmp6962, V2809)


__e.TailApply(PrimFunc(symshen_4process_1application), tmp6963, V2810)
return


} else {
__e.Return(V2809)
return
}


}


}


}, 2)

tmp6974 := Call(__e, ns2_1set, symshen_4process_1applications, tmp6960)


_ = tmp6974

tmp6975 := MakeNative(func(__e *ControlFlow) {
V2814 := __e.Get(1)
_ = V2814
tmp6985 := PrimEqual(symdefine, V2814)

if True == tmp6985 {
__e.Return(True)
return
} else {
tmp6983 := PrimEqual(symdefun, V2814)

if True == tmp6983 {
__e.Return(True)
return
} else {
tmp6981 := PrimEqual(symsynonyms, V2814)

if True == tmp6981 {
__e.Return(True)
return
} else {
tmp6979 := Call(__e, PrimFunc(symshen_4special_2), V2814)


if True == tmp6979 {
__e.Return(True)
return
} else {
tmp6977 := Call(__e, PrimFunc(symshen_4extraspecial_2), V2814)


if True == tmp6977 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}


}


}


}, 1)

tmp6986 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp6975)


_ = tmp6986

tmp6987 := MakeNative(func(__e *ControlFlow) {
V2819 := __e.Get(1)
_ = V2819
V2820 := __e.Get(2)
_ = V2820
V2821 := __e.Get(3)
_ = V2821
tmp7229 := PrimEqual(symlambda, V2819)

var ifres7207 Obj

if True == tmp7229 {
tmp7228 := PrimIsPair(V2820)

var ifres7209 Obj

if True == tmp7228 {
tmp7226 := PrimHead(V2820)

tmp7227 := PrimEqual(symlambda, tmp7226)

var ifres7211 Obj

if True == tmp7227 {
tmp7224 := PrimTail(V2820)

tmp7225 := PrimIsPair(tmp7224)

var ifres7213 Obj

if True == tmp7225 {
tmp7221 := PrimTail(V2820)

tmp7222 := PrimTail(tmp7221)

tmp7223 := PrimIsPair(tmp7222)

var ifres7215 Obj

if True == tmp7223 {
tmp7217 := PrimTail(V2820)

tmp7218 := PrimTail(tmp7217)

tmp7219 := PrimTail(tmp7218)

tmp7220 := PrimEqual(Nil, tmp7219)

var ifres7216 Obj

if True == tmp7220 {
ifres7216 = True


} else {
ifres7216 = False


}

ifres7215 = ifres7216


} else {
ifres7215 = False


}

var ifres7214 Obj

if True == ifres7215 {
ifres7214 = True


} else {
ifres7214 = False


}

ifres7213 = ifres7214


} else {
ifres7213 = False


}

var ifres7212 Obj

if True == ifres7213 {
ifres7212 = True


} else {
ifres7212 = False


}

ifres7211 = ifres7212


} else {
ifres7211 = False


}

var ifres7210 Obj

if True == ifres7211 {
ifres7210 = True


} else {
ifres7210 = False


}

ifres7209 = ifres7210


} else {
ifres7209 = False


}

var ifres7208 Obj

if True == ifres7209 {
ifres7208 = True


} else {
ifres7208 = False


}

ifres7207 = ifres7208


} else {
ifres7207 = False


}

if True == ifres7207 {
tmp6988 := PrimTail(V2820)

tmp6989 := PrimHead(tmp6988)

tmp6990 := PrimTail(V2820)

tmp6991 := PrimTail(tmp6990)

tmp6992 := PrimHead(tmp6991)

tmp6993 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6992, V2821)


tmp6994 := PrimCons(tmp6993, Nil)

tmp6995 := PrimCons(tmp6989, tmp6994)

__e.Return(PrimCons(symlambda, tmp6995))
return


} else {
tmp7205 := PrimEqual(symlet, V2819)

var ifres7176 Obj

if True == tmp7205 {
tmp7204 := PrimIsPair(V2820)

var ifres7178 Obj

if True == tmp7204 {
tmp7202 := PrimHead(V2820)

tmp7203 := PrimEqual(symlet, tmp7202)

var ifres7180 Obj

if True == tmp7203 {
tmp7200 := PrimTail(V2820)

tmp7201 := PrimIsPair(tmp7200)

var ifres7182 Obj

if True == tmp7201 {
tmp7197 := PrimTail(V2820)

tmp7198 := PrimTail(tmp7197)

tmp7199 := PrimIsPair(tmp7198)

var ifres7184 Obj

if True == tmp7199 {
tmp7193 := PrimTail(V2820)

tmp7194 := PrimTail(tmp7193)

tmp7195 := PrimTail(tmp7194)

tmp7196 := PrimIsPair(tmp7195)

var ifres7186 Obj

if True == tmp7196 {
tmp7188 := PrimTail(V2820)

tmp7189 := PrimTail(tmp7188)

tmp7190 := PrimTail(tmp7189)

tmp7191 := PrimTail(tmp7190)

tmp7192 := PrimEqual(Nil, tmp7191)

var ifres7187 Obj

if True == tmp7192 {
ifres7187 = True


} else {
ifres7187 = False


}

ifres7186 = ifres7187


} else {
ifres7186 = False


}

var ifres7185 Obj

if True == ifres7186 {
ifres7185 = True


} else {
ifres7185 = False


}

ifres7184 = ifres7185


} else {
ifres7184 = False


}

var ifres7183 Obj

if True == ifres7184 {
ifres7183 = True


} else {
ifres7183 = False


}

ifres7182 = ifres7183


} else {
ifres7182 = False


}

var ifres7181 Obj

if True == ifres7182 {
ifres7181 = True


} else {
ifres7181 = False


}

ifres7180 = ifres7181


} else {
ifres7180 = False


}

var ifres7179 Obj

if True == ifres7180 {
ifres7179 = True


} else {
ifres7179 = False


}

ifres7178 = ifres7179


} else {
ifres7178 = False


}

var ifres7177 Obj

if True == ifres7178 {
ifres7177 = True


} else {
ifres7177 = False


}

ifres7176 = ifres7177


} else {
ifres7176 = False


}

if True == ifres7176 {
tmp6996 := PrimTail(V2820)

tmp6997 := PrimHead(tmp6996)

tmp6998 := PrimTail(V2820)

tmp6999 := PrimTail(tmp6998)

tmp7000 := PrimHead(tmp6999)

tmp7001 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7000, V2821)


tmp7002 := PrimTail(V2820)

tmp7003 := PrimTail(tmp7002)

tmp7004 := PrimTail(tmp7003)

tmp7005 := PrimHead(tmp7004)

tmp7006 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7005, V2821)


tmp7007 := PrimCons(tmp7006, Nil)

tmp7008 := PrimCons(tmp7001, tmp7007)

tmp7009 := PrimCons(tmp6997, tmp7008)

__e.Return(PrimCons(symlet, tmp7009))
return


} else {
tmp7174 := PrimEqual(symdefun, V2819)

var ifres7145 Obj

if True == tmp7174 {
tmp7173 := PrimIsPair(V2820)

var ifres7147 Obj

if True == tmp7173 {
tmp7171 := PrimHead(V2820)

tmp7172 := PrimEqual(symdefun, tmp7171)

var ifres7149 Obj

if True == tmp7172 {
tmp7169 := PrimTail(V2820)

tmp7170 := PrimIsPair(tmp7169)

var ifres7151 Obj

if True == tmp7170 {
tmp7166 := PrimTail(V2820)

tmp7167 := PrimTail(tmp7166)

tmp7168 := PrimIsPair(tmp7167)

var ifres7153 Obj

if True == tmp7168 {
tmp7162 := PrimTail(V2820)

tmp7163 := PrimTail(tmp7162)

tmp7164 := PrimTail(tmp7163)

tmp7165 := PrimIsPair(tmp7164)

var ifres7155 Obj

if True == tmp7165 {
tmp7157 := PrimTail(V2820)

tmp7158 := PrimTail(tmp7157)

tmp7159 := PrimTail(tmp7158)

tmp7160 := PrimTail(tmp7159)

tmp7161 := PrimEqual(Nil, tmp7160)

var ifres7156 Obj

if True == tmp7161 {
ifres7156 = True


} else {
ifres7156 = False


}

ifres7155 = ifres7156


} else {
ifres7155 = False


}

var ifres7154 Obj

if True == ifres7155 {
ifres7154 = True


} else {
ifres7154 = False


}

ifres7153 = ifres7154


} else {
ifres7153 = False


}

var ifres7152 Obj

if True == ifres7153 {
ifres7152 = True


} else {
ifres7152 = False


}

ifres7151 = ifres7152


} else {
ifres7151 = False


}

var ifres7150 Obj

if True == ifres7151 {
ifres7150 = True


} else {
ifres7150 = False


}

ifres7149 = ifres7150


} else {
ifres7149 = False


}

var ifres7148 Obj

if True == ifres7149 {
ifres7148 = True


} else {
ifres7148 = False


}

ifres7147 = ifres7148


} else {
ifres7147 = False


}

var ifres7146 Obj

if True == ifres7147 {
ifres7146 = True


} else {
ifres7146 = False


}

ifres7145 = ifres7146


} else {
ifres7145 = False


}

if True == ifres7145 {
__e.Return(V2820)
return
} else {
tmp7143 := PrimEqual(symdefine, V2819)

var ifres7121 Obj

if True == tmp7143 {
tmp7142 := PrimIsPair(V2820)

var ifres7123 Obj

if True == tmp7142 {
tmp7140 := PrimHead(V2820)

tmp7141 := PrimEqual(symdefine, tmp7140)

var ifres7125 Obj

if True == tmp7141 {
tmp7138 := PrimTail(V2820)

tmp7139 := PrimIsPair(tmp7138)

var ifres7127 Obj

if True == tmp7139 {
tmp7135 := PrimTail(V2820)

tmp7136 := PrimTail(tmp7135)

tmp7137 := PrimIsPair(tmp7136)

var ifres7129 Obj

if True == tmp7137 {
tmp7131 := PrimTail(V2820)

tmp7132 := PrimTail(tmp7131)

tmp7133 := PrimHead(tmp7132)

tmp7134 := PrimEqual(sym_i, tmp7133)

var ifres7130 Obj

if True == tmp7134 {
ifres7130 = True


} else {
ifres7130 = False


}

ifres7129 = ifres7130


} else {
ifres7129 = False


}

var ifres7128 Obj

if True == ifres7129 {
ifres7128 = True


} else {
ifres7128 = False


}

ifres7127 = ifres7128


} else {
ifres7127 = False


}

var ifres7126 Obj

if True == ifres7127 {
ifres7126 = True


} else {
ifres7126 = False


}

ifres7125 = ifres7126


} else {
ifres7125 = False


}

var ifres7124 Obj

if True == ifres7125 {
ifres7124 = True


} else {
ifres7124 = False


}

ifres7123 = ifres7124


} else {
ifres7123 = False


}

var ifres7122 Obj

if True == ifres7123 {
ifres7122 = True


} else {
ifres7122 = False


}

ifres7121 = ifres7122


} else {
ifres7121 = False


}

if True == ifres7121 {
tmp7010 := PrimTail(V2820)

tmp7011 := PrimHead(tmp7010)

tmp7012 := PrimTail(V2820)

tmp7013 := PrimHead(tmp7012)

tmp7014 := PrimTail(V2820)

tmp7015 := PrimTail(tmp7014)

tmp7016 := PrimTail(tmp7015)

tmp7017 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp7013, tmp7016, V2821)


tmp7018 := PrimCons(sym_i, tmp7017)

tmp7019 := PrimCons(tmp7011, tmp7018)

__e.Return(PrimCons(symdefine, tmp7019))
return


} else {
tmp7119 := PrimEqual(symdefine, V2819)

var ifres7108 Obj

if True == tmp7119 {
tmp7118 := PrimIsPair(V2820)

var ifres7110 Obj

if True == tmp7118 {
tmp7116 := PrimHead(V2820)

tmp7117 := PrimEqual(symdefine, tmp7116)

var ifres7112 Obj

if True == tmp7117 {
tmp7114 := PrimTail(V2820)

tmp7115 := PrimIsPair(tmp7114)

var ifres7113 Obj

if True == tmp7115 {
ifres7113 = True


} else {
ifres7113 = False


}

ifres7112 = ifres7113


} else {
ifres7112 = False


}

var ifres7111 Obj

if True == ifres7112 {
ifres7111 = True


} else {
ifres7111 = False


}

ifres7110 = ifres7111


} else {
ifres7110 = False


}

var ifres7109 Obj

if True == ifres7110 {
ifres7109 = True


} else {
ifres7109 = False


}

ifres7108 = ifres7109


} else {
ifres7108 = False


}

if True == ifres7108 {
tmp7020 := PrimTail(V2820)

tmp7021 := PrimHead(tmp7020)

tmp7022 := MakeNative(func(__e *ControlFlow) {
Z2822 := __e.Get(1)
_ = Z2822
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2822, V2821)
return
}, 1)

tmp7023 := PrimTail(V2820)

tmp7024 := PrimTail(tmp7023)

tmp7025 := Call(__e, PrimFunc(symmap), tmp7022, tmp7024)


tmp7026 := PrimCons(tmp7021, tmp7025)

__e.Return(PrimCons(symdefine, tmp7026))
return


} else {
tmp7106 := PrimEqual(symsynonyms, V2819)

if True == tmp7106 {
__e.Return(PrimCons(symsynonyms, V2820))
return
} else {
tmp7104 := PrimEqual(symtype, V2819)

var ifres7082 Obj

if True == tmp7104 {
tmp7103 := PrimIsPair(V2820)

var ifres7084 Obj

if True == tmp7103 {
tmp7101 := PrimHead(V2820)

tmp7102 := PrimEqual(symtype, tmp7101)

var ifres7086 Obj

if True == tmp7102 {
tmp7099 := PrimTail(V2820)

tmp7100 := PrimIsPair(tmp7099)

var ifres7088 Obj

if True == tmp7100 {
tmp7096 := PrimTail(V2820)

tmp7097 := PrimTail(tmp7096)

tmp7098 := PrimIsPair(tmp7097)

var ifres7090 Obj

if True == tmp7098 {
tmp7092 := PrimTail(V2820)

tmp7093 := PrimTail(tmp7092)

tmp7094 := PrimTail(tmp7093)

tmp7095 := PrimEqual(Nil, tmp7094)

var ifres7091 Obj

if True == tmp7095 {
ifres7091 = True


} else {
ifres7091 = False


}

ifres7090 = ifres7091


} else {
ifres7090 = False


}

var ifres7089 Obj

if True == ifres7090 {
ifres7089 = True


} else {
ifres7089 = False


}

ifres7088 = ifres7089


} else {
ifres7088 = False


}

var ifres7087 Obj

if True == ifres7088 {
ifres7087 = True


} else {
ifres7087 = False


}

ifres7086 = ifres7087


} else {
ifres7086 = False


}

var ifres7085 Obj

if True == ifres7086 {
ifres7085 = True


} else {
ifres7085 = False


}

ifres7084 = ifres7085


} else {
ifres7084 = False


}

var ifres7083 Obj

if True == ifres7084 {
ifres7083 = True


} else {
ifres7083 = False


}

ifres7082 = ifres7083


} else {
ifres7082 = False


}

if True == ifres7082 {
tmp7027 := PrimTail(V2820)

tmp7028 := PrimHead(tmp7027)

tmp7029 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7028, V2821)


tmp7030 := PrimTail(V2820)

tmp7031 := PrimTail(tmp7030)

tmp7032 := PrimCons(tmp7029, tmp7031)

__e.Return(PrimCons(symtype, tmp7032))
return


} else {
tmp7080 := PrimEqual(syminput_7, V2819)

var ifres7058 Obj

if True == tmp7080 {
tmp7079 := PrimIsPair(V2820)

var ifres7060 Obj

if True == tmp7079 {
tmp7077 := PrimHead(V2820)

tmp7078 := PrimEqual(syminput_7, tmp7077)

var ifres7062 Obj

if True == tmp7078 {
tmp7075 := PrimTail(V2820)

tmp7076 := PrimIsPair(tmp7075)

var ifres7064 Obj

if True == tmp7076 {
tmp7072 := PrimTail(V2820)

tmp7073 := PrimTail(tmp7072)

tmp7074 := PrimIsPair(tmp7073)

var ifres7066 Obj

if True == tmp7074 {
tmp7068 := PrimTail(V2820)

tmp7069 := PrimTail(tmp7068)

tmp7070 := PrimTail(tmp7069)

tmp7071 := PrimEqual(Nil, tmp7070)

var ifres7067 Obj

if True == tmp7071 {
ifres7067 = True


} else {
ifres7067 = False


}

ifres7066 = ifres7067


} else {
ifres7066 = False


}

var ifres7065 Obj

if True == ifres7066 {
ifres7065 = True


} else {
ifres7065 = False


}

ifres7064 = ifres7065


} else {
ifres7064 = False


}

var ifres7063 Obj

if True == ifres7064 {
ifres7063 = True


} else {
ifres7063 = False


}

ifres7062 = ifres7063


} else {
ifres7062 = False


}

var ifres7061 Obj

if True == ifres7062 {
ifres7061 = True


} else {
ifres7061 = False


}

ifres7060 = ifres7061


} else {
ifres7060 = False


}

var ifres7059 Obj

if True == ifres7060 {
ifres7059 = True


} else {
ifres7059 = False


}

ifres7058 = ifres7059


} else {
ifres7058 = False


}

if True == ifres7058 {
tmp7033 := PrimTail(V2820)

tmp7034 := PrimHead(tmp7033)

tmp7035 := PrimTail(V2820)

tmp7036 := PrimTail(tmp7035)

tmp7037 := PrimHead(tmp7036)

tmp7038 := Call(__e, PrimFunc(symshen_4process_1applications), tmp7037, V2821)


tmp7039 := PrimCons(tmp7038, Nil)

tmp7040 := PrimCons(tmp7034, tmp7039)

__e.Return(PrimCons(syminput_7, tmp7040))
return


} else {
tmp7056 := PrimIsPair(V2820)

var ifres7052 Obj

if True == tmp7056 {
tmp7054 := PrimHead(V2820)

tmp7055 := Call(__e, PrimFunc(symshen_4special_2), tmp7054)


var ifres7053 Obj

if True == tmp7055 {
ifres7053 = True


} else {
ifres7053 = False


}

ifres7052 = ifres7053


} else {
ifres7052 = False


}

if True == ifres7052 {
tmp7041 := PrimHead(V2820)

tmp7042 := MakeNative(func(__e *ControlFlow) {
Z2823 := __e.Get(1)
_ = Z2823
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2823, V2821)
return
}, 1)

tmp7043 := PrimTail(V2820)

tmp7044 := Call(__e, PrimFunc(symmap), tmp7042, tmp7043)


__e.Return(PrimCons(tmp7041, tmp7044))
return


} else {
tmp7050 := PrimIsPair(V2820)

var ifres7046 Obj

if True == tmp7050 {
tmp7048 := PrimHead(V2820)

tmp7049 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp7048)


var ifres7047 Obj

if True == tmp7049 {
ifres7047 = True


} else {
ifres7047 = False


}

ifres7046 = ifres7047


} else {
ifres7046 = False


}

if True == ifres7046 {
__e.Return(V2820)
return
} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.special-case")))
return
}


}


}


}


}


}


}


}


}


}


}, 3)

tmp7230 := Call(__e, ns2_1set, symshen_4special_1case, tmp6987)


_ = tmp7230

tmp7231 := MakeNative(func(__e *ControlFlow) {
V2826 := __e.Get(1)
_ = V2826
V2827 := __e.Get(2)
_ = V2827
V2828 := __e.Get(3)
_ = V2828
tmp7247 := PrimIsPair(V2827)

var ifres7243 Obj

if True == tmp7247 {
tmp7245 := PrimHead(V2827)

tmp7246 := PrimEqual(sym_j, tmp7245)

var ifres7244 Obj

if True == tmp7246 {
ifres7244 = True


} else {
ifres7244 = False


}

ifres7243 = ifres7244


} else {
ifres7243 = False


}

if True == ifres7243 {
tmp7232 := MakeNative(func(__e *ControlFlow) {
Z2829 := __e.Get(1)
_ = Z2829
__e.TailApply(PrimFunc(symshen_4process_1applications), Z2829, V2828)
return
}, 1)

tmp7233 := PrimTail(V2827)

tmp7234 := Call(__e, PrimFunc(symmap), tmp7232, tmp7233)


__e.Return(PrimCons(sym_j, tmp7234))
return


} else {
tmp7241 := PrimIsPair(V2827)

if True == tmp7241 {
tmp7235 := PrimHead(V2827)

tmp7236 := PrimTail(V2827)

tmp7237 := Call(__e, PrimFunc(symshen_4process_1after_1type), V2826, tmp7236, V2828)


__e.Return(PrimCons(tmp7235, tmp7237))
return


} else {
tmp7238 := Call(__e, PrimFunc(symshen_4app), V2826, MakeString("\n"), symshen_4a)


tmp7239 := PrimStringConcat(MakeString("missing } in "), tmp7238)

__e.Return(PrimSimpleError(tmp7239))
return


}


}


}, 3)

tmp7248 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp7231)


_ = tmp7248

tmp7249 := MakeNative(func(__e *ControlFlow) {
V2830 := __e.Get(1)
_ = V2830
V2831 := __e.Get(2)
_ = V2831
tmp7294 := PrimIsPair(V2830)

if True == tmp7294 {
tmp7250 := MakeNative(func(__e *ControlFlow) {
W2832 := __e.Get(1)
_ = W2832
tmp7251 := MakeNative(func(__e *ControlFlow) {
W2833 := __e.Get(1)
_ = W2833
tmp7288 := Call(__e, PrimFunc(symelement_2), V2830, V2831)


if True == tmp7288 {
__e.Return(V2830)
return
} else {
tmp7285 := PrimHead(V2830)

tmp7286 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp7285)


if True == tmp7286 {
__e.Return(V2830)
return
} else {
tmp7283 := Call(__e, PrimFunc(symshen_4foreign_2), V2830)


if True == tmp7283 {
__e.TailApply(PrimFunc(symshen_4unpack_1foreign), V2830)
return
} else {
tmp7281 := Call(__e, PrimFunc(symshen_4fn_1call_2), V2830)


if True == tmp7281 {
__e.TailApply(PrimFunc(symshen_4fn_1call), V2830)
return
} else {
tmp7279 := Call(__e, PrimFunc(symshen_4zero_1place_2), V2830)


if True == tmp7279 {
__e.Return(V2830)
return
} else {
tmp7276 := PrimHead(V2830)

tmp7277 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp7276, W2832)


if True == tmp7277 {
tmp7252 := PrimHead(V2830)

tmp7253 := PrimCons(tmp7252, Nil)

tmp7254 := PrimCons(symfn, tmp7253)

tmp7255 := PrimTail(V2830)

tmp7256 := PrimCons(tmp7254, tmp7255)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7256)
return


} else {
tmp7273 := PrimHead(V2830)

tmp7274 := PrimIsVariable(tmp7273)

if True == tmp7274 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2830)
return
} else {
tmp7270 := PrimHead(V2830)

tmp7271 := Call(__e, PrimFunc(symshen_4application_2), tmp7270)


if True == tmp7271 {
__e.TailApply(PrimFunc(symshen_4simple_1curry), V2830)
return
} else {
tmp7267 := PrimHead(V2830)

tmp7268 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp7267, W2832, W2833)


if True == tmp7268 {
tmp7257 := PrimNumberSubtract(W2832, W2833)

__e.TailApply(PrimFunc(symshen_4lambda_1function), V2830, tmp7257)
return


} else {
tmp7264 := PrimHead(V2830)

tmp7265 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp7264, W2832, W2833)


if True == tmp7265 {
tmp7258 := PrimHead(V2830)

tmp7259 := PrimCons(tmp7258, Nil)

tmp7260 := PrimCons(symfn, tmp7259)

tmp7261 := PrimTail(V2830)

tmp7262 := PrimCons(tmp7260, tmp7261)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7262)
return


} else {
__e.Return(V2830)
return
}


}


}


}


}


}


}


}


}


}


}, 1)

tmp7289 := PrimTail(V2830)

tmp7290 := Call(__e, PrimFunc(symlength), tmp7289)


__e.TailApply(tmp7251, tmp7290)
return


}, 1)

tmp7291 := PrimHead(V2830)

tmp7292 := Call(__e, PrimFunc(symarity), tmp7291)


__e.TailApply(tmp7250, tmp7292)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.process-application")))
return
}


}, 2)

tmp7295 := Call(__e, ns2_1set, symshen_4process_1application, tmp7249)


_ = tmp7295

tmp7296 := MakeNative(func(__e *ControlFlow) {
V2834 := __e.Get(1)
_ = V2834
tmp7322 := PrimIsPair(V2834)

var ifres7302 Obj

if True == tmp7322 {
tmp7320 := PrimHead(V2834)

tmp7321 := PrimIsPair(tmp7320)

var ifres7304 Obj

if True == tmp7321 {
tmp7317 := PrimHead(V2834)

tmp7318 := PrimHead(tmp7317)

tmp7319 := PrimEqual(symforeign, tmp7318)

var ifres7306 Obj

if True == tmp7319 {
tmp7314 := PrimHead(V2834)

tmp7315 := PrimTail(tmp7314)

tmp7316 := PrimIsPair(tmp7315)

var ifres7308 Obj

if True == tmp7316 {
tmp7310 := PrimHead(V2834)

tmp7311 := PrimTail(tmp7310)

tmp7312 := PrimTail(tmp7311)

tmp7313 := PrimEqual(Nil, tmp7312)

var ifres7309 Obj

if True == tmp7313 {
ifres7309 = True


} else {
ifres7309 = False


}

ifres7308 = ifres7309


} else {
ifres7308 = False


}

var ifres7307 Obj

if True == ifres7308 {
ifres7307 = True


} else {
ifres7307 = False


}

ifres7306 = ifres7307


} else {
ifres7306 = False


}

var ifres7305 Obj

if True == ifres7306 {
ifres7305 = True


} else {
ifres7305 = False


}

ifres7304 = ifres7305


} else {
ifres7304 = False


}

var ifres7303 Obj

if True == ifres7304 {
ifres7303 = True


} else {
ifres7303 = False


}

ifres7302 = ifres7303


} else {
ifres7302 = False


}

if True == ifres7302 {
tmp7297 := PrimHead(V2834)

tmp7298 := PrimTail(tmp7297)

tmp7299 := PrimHead(tmp7298)

tmp7300 := PrimTail(V2834)

__e.Return(PrimCons(tmp7299, tmp7300))
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.unpack-foreign")))
return
}


}, 1)

tmp7323 := Call(__e, ns2_1set, symshen_4unpack_1foreign, tmp7296)


_ = tmp7323

tmp7324 := MakeNative(func(__e *ControlFlow) {
V2837 := __e.Get(1)
_ = V2837
tmp7346 := PrimIsPair(V2837)

var ifres7326 Obj

if True == tmp7346 {
tmp7344 := PrimHead(V2837)

tmp7345 := PrimIsPair(tmp7344)

var ifres7328 Obj

if True == tmp7345 {
tmp7341 := PrimHead(V2837)

tmp7342 := PrimHead(tmp7341)

tmp7343 := PrimEqual(symforeign, tmp7342)

var ifres7330 Obj

if True == tmp7343 {
tmp7338 := PrimHead(V2837)

tmp7339 := PrimTail(tmp7338)

tmp7340 := PrimIsPair(tmp7339)

var ifres7332 Obj

if True == tmp7340 {
tmp7334 := PrimHead(V2837)

tmp7335 := PrimTail(tmp7334)

tmp7336 := PrimTail(tmp7335)

tmp7337 := PrimEqual(Nil, tmp7336)

var ifres7333 Obj

if True == tmp7337 {
ifres7333 = True


} else {
ifres7333 = False


}

ifres7332 = ifres7333


} else {
ifres7332 = False


}

var ifres7331 Obj

if True == ifres7332 {
ifres7331 = True


} else {
ifres7331 = False


}

ifres7330 = ifres7331


} else {
ifres7330 = False


}

var ifres7329 Obj

if True == ifres7330 {
ifres7329 = True


} else {
ifres7329 = False


}

ifres7328 = ifres7329


} else {
ifres7328 = False


}

var ifres7327 Obj

if True == ifres7328 {
ifres7327 = True


} else {
ifres7327 = False


}

ifres7326 = ifres7327


} else {
ifres7326 = False


}

if True == ifres7326 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7347 := Call(__e, ns2_1set, symshen_4foreign_2, tmp7324)


_ = tmp7347

tmp7348 := MakeNative(func(__e *ControlFlow) {
V2840 := __e.Get(1)
_ = V2840
tmp7354 := PrimIsPair(V2840)

var ifres7350 Obj

if True == tmp7354 {
tmp7352 := PrimTail(V2840)

tmp7353 := PrimEqual(Nil, tmp7352)

var ifres7351 Obj

if True == tmp7353 {
ifres7351 = True


} else {
ifres7351 = False


}

ifres7350 = ifres7351


} else {
ifres7350 = False


}

if True == ifres7350 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}, 1)

tmp7355 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp7348)


_ = tmp7355

tmp7356 := MakeNative(func(__e *ControlFlow) {
V2841 := __e.Get(1)
_ = V2841
tmp7361 := PrimIsSymbol(V2841)

if True == tmp7361 {
tmp7358 := PrimStr(V2841)

tmp7359 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp7358)


if True == tmp7359 {
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


}, 1)

tmp7362 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp7356)


_ = tmp7362

tmp7363 := MakeNative(func(__e *ControlFlow) {
V2846 := __e.Get(1)
_ = V2846
tmp7393 := PrimIsPair(V2846)

var ifres7380 Obj

if True == tmp7393 {
tmp7391 := PrimHead(V2846)

tmp7392 := PrimEqual(symprotect, tmp7391)

var ifres7382 Obj

if True == tmp7392 {
tmp7389 := PrimTail(V2846)

tmp7390 := PrimIsPair(tmp7389)

var ifres7384 Obj

if True == tmp7390 {
tmp7386 := PrimTail(V2846)

tmp7387 := PrimTail(tmp7386)

tmp7388 := PrimEqual(Nil, tmp7387)

var ifres7385 Obj

if True == tmp7388 {
ifres7385 = True


} else {
ifres7385 = False


}

ifres7384 = ifres7385


} else {
ifres7384 = False


}

var ifres7383 Obj

if True == ifres7384 {
ifres7383 = True


} else {
ifres7383 = False


}

ifres7382 = ifres7383


} else {
ifres7382 = False


}

var ifres7381 Obj

if True == ifres7382 {
ifres7381 = True


} else {
ifres7381 = False


}

ifres7380 = ifres7381


} else {
ifres7380 = False


}

if True == ifres7380 {
__e.Return(False)
return
} else {
tmp7378 := PrimIsPair(V2846)

var ifres7365 Obj

if True == tmp7378 {
tmp7376 := PrimHead(V2846)

tmp7377 := PrimEqual(symforeign, tmp7376)

var ifres7367 Obj

if True == tmp7377 {
tmp7374 := PrimTail(V2846)

tmp7375 := PrimIsPair(tmp7374)

var ifres7369 Obj

if True == tmp7375 {
tmp7371 := PrimTail(V2846)

tmp7372 := PrimTail(tmp7371)

tmp7373 := PrimEqual(Nil, tmp7372)

var ifres7370 Obj

if True == tmp7373 {
ifres7370 = True


} else {
ifres7370 = False


}

ifres7369 = ifres7370


} else {
ifres7369 = False


}

var ifres7368 Obj

if True == ifres7369 {
ifres7368 = True


} else {
ifres7368 = False


}

ifres7367 = ifres7368


} else {
ifres7367 = False


}

var ifres7366 Obj

if True == ifres7367 {
ifres7366 = True


} else {
ifres7366 = False


}

ifres7365 = ifres7366


} else {
ifres7365 = False


}

if True == ifres7365 {
__e.Return(False)
return
} else {
__e.Return(PrimIsPair(V2846))
return
}


}


}, 1)

tmp7394 := Call(__e, ns2_1set, symshen_4application_2, tmp7363)


_ = tmp7394

tmp7395 := MakeNative(func(__e *ControlFlow) {
V2851 := __e.Get(1)
_ = V2851
V2852 := __e.Get(2)
_ = V2852
tmp7403 := PrimEqual(MakeNumber(-1), V2852)

if True == tmp7403 {
tmp7401 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V2851)


if True == tmp7401 {
tmp7397 := Call(__e, PrimFunc(symexternal), symshen)


tmp7398 := Call(__e, PrimFunc(symelement_2), V2851, tmp7397)


tmp7399 := PrimNot(tmp7398)

if True == tmp7399 {
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
__e.Return(False)
return
}


}, 2)

tmp7404 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp7395)


_ = tmp7404

tmp7405 := MakeNative(func(__e *ControlFlow) {
V2853 := __e.Get(1)
_ = V2853
tmp7410 := PrimIsSymbol(V2853)

if True == tmp7410 {
tmp7407 := PrimIsVariable(V2853)

tmp7408 := PrimNot(tmp7407)

if True == tmp7408 {
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


}, 1)

tmp7411 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp7405)


_ = tmp7411

tmp7412 := MakeNative(func(__e *ControlFlow) {
V2854 := __e.Get(1)
_ = V2854
tmp7442 := PrimIsPair(V2854)

var ifres7433 Obj

if True == tmp7442 {
tmp7440 := PrimTail(V2854)

tmp7441 := PrimIsPair(tmp7440)

var ifres7435 Obj

if True == tmp7441 {
tmp7437 := PrimTail(V2854)

tmp7438 := PrimTail(tmp7437)

tmp7439 := PrimEqual(Nil, tmp7438)

var ifres7436 Obj

if True == tmp7439 {
ifres7436 = True


} else {
ifres7436 = False


}

ifres7435 = ifres7436


} else {
ifres7435 = False


}

var ifres7434 Obj

if True == ifres7435 {
ifres7434 = True


} else {
ifres7434 = False


}

ifres7433 = ifres7434


} else {
ifres7433 = False


}

if True == ifres7433 {
__e.Return(V2854)
return
} else {
tmp7431 := PrimIsPair(V2854)

var ifres7422 Obj

if True == tmp7431 {
tmp7429 := PrimTail(V2854)

tmp7430 := PrimIsPair(tmp7429)

var ifres7424 Obj

if True == tmp7430 {
tmp7426 := PrimTail(V2854)

tmp7427 := PrimTail(tmp7426)

tmp7428 := PrimIsPair(tmp7427)

var ifres7425 Obj

if True == tmp7428 {
ifres7425 = True


} else {
ifres7425 = False


}

ifres7424 = ifres7425


} else {
ifres7424 = False


}

var ifres7423 Obj

if True == ifres7424 {
ifres7423 = True


} else {
ifres7423 = False


}

ifres7422 = ifres7423


} else {
ifres7422 = False


}

if True == ifres7422 {
tmp7413 := PrimHead(V2854)

tmp7414 := PrimTail(V2854)

tmp7415 := PrimHead(tmp7414)

tmp7416 := PrimCons(tmp7415, Nil)

tmp7417 := PrimCons(tmp7413, tmp7416)

tmp7418 := PrimTail(V2854)

tmp7419 := PrimTail(tmp7418)

tmp7420 := PrimCons(tmp7417, tmp7419)

__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp7420)
return


} else {
__e.Return(V2854)
return
}


}


}, 1)

tmp7443 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp7412)


_ = tmp7443

tmp7444 := MakeNative(func(__e *ControlFlow) {
V2855 := __e.Get(1)
_ = V2855
__e.TailApply(PrimFunc(symfn), V2855)
return
}, 1)

tmp7445 := Call(__e, ns2_1set, symfunction, tmp7444)


_ = tmp7445

tmp7446 := MakeNative(func(__e *ControlFlow) {
V2856 := __e.Get(1)
_ = V2856
tmp7455 := Call(__e, PrimFunc(symarity), V2856)


tmp7456 := PrimEqual(tmp7455, MakeNumber(0))

if True == tmp7456 {
__e.TailApply(V2856)
return
} else {
tmp7447 := MakeNative(func(__e *ControlFlow) {
W2857 := __e.Get(1)
_ = W2857
tmp7451 := Call(__e, PrimFunc(symempty_2), W2857)


if True == tmp7451 {
tmp7448 := Call(__e, PrimFunc(symshen_4app), V2856, MakeString(" is undefined\n"), symshen_4a)


tmp7449 := PrimStringConcat(MakeString("fn: "), tmp7448)

__e.Return(PrimSimpleError(tmp7449))
return


} else {
__e.Return(PrimTail(W2857))
return
}


}, 1)

tmp7452 := PrimValue(symshen_4_dlambdatable_d)

tmp7453 := Call(__e, PrimFunc(symassoc), V2856, tmp7452)


__e.TailApply(tmp7447, tmp7453)
return


}


}, 1)

tmp7457 := Call(__e, ns2_1set, symfn, tmp7446)


_ = tmp7457

tmp7458 := MakeNative(func(__e *ControlFlow) {
V2860 := __e.Get(1)
_ = V2860
tmp7488 := PrimIsPair(V2860)

var ifres7475 Obj

if True == tmp7488 {
tmp7486 := PrimHead(V2860)

tmp7487 := PrimEqual(symfn, tmp7486)

var ifres7477 Obj

if True == tmp7487 {
tmp7484 := PrimTail(V2860)

tmp7485 := PrimIsPair(tmp7484)

var ifres7479 Obj

if True == tmp7485 {
tmp7481 := PrimTail(V2860)

tmp7482 := PrimTail(tmp7481)

tmp7483 := PrimEqual(Nil, tmp7482)

var ifres7480 Obj

if True == tmp7483 {
ifres7480 = True


} else {
ifres7480 = False


}

ifres7479 = ifres7480


} else {
ifres7479 = False


}

var ifres7478 Obj

if True == ifres7479 {
ifres7478 = True


} else {
ifres7478 = False


}

ifres7477 = ifres7478


} else {
ifres7477 = False


}

var ifres7476 Obj

if True == ifres7477 {
ifres7476 = True


} else {
ifres7476 = False


}

ifres7475 = ifres7476


} else {
ifres7475 = False


}

if True == ifres7475 {
__e.Return(True)
return
} else {
tmp7473 := PrimIsPair(V2860)

var ifres7460 Obj

if True == tmp7473 {
tmp7471 := PrimHead(V2860)

tmp7472 := PrimEqual(symfunction, tmp7471)

var ifres7462 Obj

if True == tmp7472 {
tmp7469 := PrimTail(V2860)

tmp7470 := PrimIsPair(tmp7469)

var ifres7464 Obj

if True == tmp7470 {
tmp7466 := PrimTail(V2860)

tmp7467 := PrimTail(tmp7466)

tmp7468 := PrimEqual(Nil, tmp7467)

var ifres7465 Obj

if True == tmp7468 {
ifres7465 = True


} else {
ifres7465 = False


}

ifres7464 = ifres7465


} else {
ifres7464 = False


}

var ifres7463 Obj

if True == ifres7464 {
ifres7463 = True


} else {
ifres7463 = False


}

ifres7462 = ifres7463


} else {
ifres7462 = False


}

var ifres7461 Obj

if True == ifres7462 {
ifres7461 = True


} else {
ifres7461 = False


}

ifres7460 = ifres7461


} else {
ifres7460 = False


}

if True == ifres7460 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

tmp7489 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp7458)


_ = tmp7489

tmp7490 := MakeNative(func(__e *ControlFlow) {
V2861 := __e.Get(1)
_ = V2861
tmp7531 := PrimIsPair(V2861)

var ifres7518 Obj

if True == tmp7531 {
tmp7529 := PrimHead(V2861)

tmp7530 := PrimEqual(symfunction, tmp7529)

var ifres7520 Obj

if True == tmp7530 {
tmp7527 := PrimTail(V2861)

tmp7528 := PrimIsPair(tmp7527)

var ifres7522 Obj

if True == tmp7528 {
tmp7524 := PrimTail(V2861)

tmp7525 := PrimTail(tmp7524)

tmp7526 := PrimEqual(Nil, tmp7525)

var ifres7523 Obj

if True == tmp7526 {
ifres7523 = True


} else {
ifres7523 = False


}

ifres7522 = ifres7523


} else {
ifres7522 = False


}

var ifres7521 Obj

if True == ifres7522 {
ifres7521 = True


} else {
ifres7521 = False


}

ifres7520 = ifres7521


} else {
ifres7520 = False


}

var ifres7519 Obj

if True == ifres7520 {
ifres7519 = True


} else {
ifres7519 = False


}

ifres7518 = ifres7519


} else {
ifres7518 = False


}

if True == ifres7518 {
tmp7491 := PrimTail(V2861)

tmp7492 := PrimCons(symfn, tmp7491)

__e.TailApply(PrimFunc(symshen_4fn_1call), tmp7492)
return


} else {
tmp7516 := PrimIsPair(V2861)

var ifres7503 Obj

if True == tmp7516 {
tmp7514 := PrimHead(V2861)

tmp7515 := PrimEqual(symfn, tmp7514)

var ifres7505 Obj

if True == tmp7515 {
tmp7512 := PrimTail(V2861)

tmp7513 := PrimIsPair(tmp7512)

var ifres7507 Obj

if True == tmp7513 {
tmp7509 := PrimTail(V2861)

tmp7510 := PrimTail(tmp7509)

tmp7511 := PrimEqual(Nil, tmp7510)

var ifres7508 Obj

if True == tmp7511 {
ifres7508 = True


} else {
ifres7508 = False


}

ifres7507 = ifres7508


} else {
ifres7507 = False


}

var ifres7506 Obj

if True == ifres7507 {
ifres7506 = True


} else {
ifres7506 = False


}

ifres7505 = ifres7506


} else {
ifres7505 = False


}

var ifres7504 Obj

if True == ifres7505 {
ifres7504 = True


} else {
ifres7504 = False


}

ifres7503 = ifres7504


} else {
ifres7503 = False


}

if True == ifres7503 {
tmp7493 := MakeNative(func(__e *ControlFlow) {
W2862 := __e.Get(1)
_ = W2862
tmp7498 := PrimEqual(W2862, MakeNumber(-1))

if True == tmp7498 {
__e.Return(V2861)
return
} else {
tmp7496 := PrimEqual(W2862, MakeNumber(0))

if True == tmp7496 {
__e.Return(PrimTail(V2861))
return
} else {
tmp7494 := PrimTail(V2861)

__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp7494, W2862)
return


}


}


}, 1)

tmp7499 := PrimTail(V2861)

tmp7500 := PrimHead(tmp7499)

tmp7501 := Call(__e, PrimFunc(symarity), tmp7500)


__e.TailApply(tmp7493, tmp7501)
return


} else {
__e.Return(PrimSimpleError(MakeString("partial function shen.fn-call")))
return
}


}


}, 1)

tmp7532 := Call(__e, ns2_1set, symshen_4fn_1call, tmp7490)


_ = tmp7532

tmp7533 := MakeNative(func(__e *ControlFlow) {
V2863 := __e.Get(1)
_ = V2863
V2864 := __e.Get(2)
_ = V2864
V2865 := __e.Get(3)
_ = V2865
tmp7534 := MakeNative(func(__e *ControlFlow) {
W2866 := __e.Get(1)
_ = W2866
tmp7535 := MakeNative(func(__e *ControlFlow) {
W2867 := __e.Get(1)
_ = W2867
__e.Return(W2866)
return
}, 1)

var ifres7541 Obj

if True == W2866 {
tmp7549 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7543 Obj

if True == tmp7549 {
tmp7545 := PrimCons(sym_1, Nil)

tmp7546 := PrimCons(sym_7, tmp7545)

tmp7547 := Call(__e, PrimFunc(symelement_2), V2863, tmp7546)


tmp7548 := PrimNot(tmp7547)

var ifres7544 Obj

if True == tmp7548 {
ifres7544 = True


} else {
ifres7544 = False


}

ifres7543 = ifres7544


} else {
ifres7543 = False


}

var ifres7542 Obj

if True == ifres7543 {
ifres7542 = True


} else {
ifres7542 = False


}

ifres7541 = ifres7542


} else {
ifres7541 = False


}

var ifres7536 Obj

if True == ifres7541 {
tmp7537 := Call(__e, PrimFunc(symshen_4app), V2863, MakeString("\n"), symshen_4a)


tmp7538 := PrimStringConcat(MakeString("partial application of "), tmp7537)

tmp7539 := Call(__e, PrimFunc(symstoutput))


tmp7540 := Call(__e, PrimFunc(sympr), tmp7538, tmp7539)


ifres7536 = tmp7540


} else {
ifres7536 = symshen_4skip


}

__e.TailApply(tmp7535, ifres7536)
return


}, 1)

tmp7550 := PrimGreatThan(V2864, V2865)

__e.TailApply(tmp7534, tmp7550)
return


}, 3)

tmp7551 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp7533)


_ = tmp7551

tmp7552 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(symshen_4_dloading_2_d))
return
}, 0)

tmp7553 := Call(__e, ns2_1set, symshen_4loading_2, tmp7552)


_ = tmp7553

tmp7554 := MakeNative(func(__e *ControlFlow) {
V2872 := __e.Get(1)
_ = V2872
V2873 := __e.Get(2)
_ = V2873
V2874 := __e.Get(3)
_ = V2874
tmp7572 := PrimEqual(MakeNumber(-1), V2873)

if True == tmp7572 {
__e.Return(False)
return
} else {
tmp7555 := MakeNative(func(__e *ControlFlow) {
W2875 := __e.Get(1)
_ = W2875
tmp7556 := MakeNative(func(__e *ControlFlow) {
W2876 := __e.Get(1)
_ = W2876
__e.Return(W2875)
return
}, 1)

var ifres7567 Obj

if True == W2875 {
tmp7569 := Call(__e, PrimFunc(symshen_4loading_2))


var ifres7568 Obj

if True == tmp7569 {
ifres7568 = True


} else {
ifres7568 = False


}

ifres7567 = ifres7568


} else {
ifres7567 = False


}

var ifres7557 Obj

if True == ifres7567 {
tmp7559 := PrimEqual(V2874, MakeNumber(1))

var ifres7558 Obj

if True == tmp7559 {
ifres7558 = MakeString("")


} else {
ifres7558 = MakeString("s")


}

tmp7560 := Call(__e, PrimFunc(symshen_4app), ifres7558, MakeString("\n"), symshen_4a)


tmp7561 := PrimStringConcat(MakeString(" argument"), tmp7560)

tmp7562 := Call(__e, PrimFunc(symshen_4app), V2874, tmp7561, symshen_4a)


tmp7563 := PrimStringConcat(MakeString(" might not like "), tmp7562)

tmp7564 := Call(__e, PrimFunc(symshen_4app), V2872, tmp7563, symshen_4a)


tmp7565 := Call(__e, PrimFunc(symstoutput))


tmp7566 := Call(__e, PrimFunc(sympr), tmp7564, tmp7565)


ifres7557 = tmp7566


} else {
ifres7557 = symshen_4skip


}

__e.TailApply(tmp7556, ifres7557)
return


}, 1)

tmp7570 := PrimLessThan(V2873, V2874)

__e.TailApply(tmp7555, tmp7570)
return


}


}, 3)

__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp7554)
return




}, 0)

