package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4initialise_1environment Obj               // shen.initialise-environment
var __defun__shen_4initialise_1signedfuncs Obj               // shen.initialise-signedfuncs
var __defun__shen_4initialise_1signedfunc_1lambda_1forms Obj // shen.initialise-signedfunc-lambda-forms
var __defun__shen_4initialise_1lambda_1forms Obj             // shen.initialise-lambda-forms
var __defun__shen_4initialise Obj                            // shen.initialise

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		reg31663 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__e.Return(reg31663)
		return
	}, 0))
	__defun__shen_4initialise_1environment = MakeNative(func(__e Evaluator, __args ...Obj) {
		reg31664 := MakeSymbol("shen.*installing-kl*")
		reg31665 := False
		reg31666 := PrimSet(reg31664, reg31665)
		_ = reg31666
		reg31667 := MakeSymbol("shen.*history*")
		reg31668 := Nil
		reg31669 := PrimSet(reg31667, reg31668)
		_ = reg31669
		reg31670 := MakeSymbol("shen.*tc*")
		reg31671 := False
		reg31672 := PrimSet(reg31670, reg31671)
		_ = reg31672
		reg31673 := MakeSymbol("*property-vector*")
		reg31674 := MakeNumber(20000)
		reg31675 := Call(__e, __defun__shen_4dict, reg31674)
		reg31676 := PrimSet(reg31673, reg31675)
		_ = reg31676
		reg31677 := MakeSymbol("shen.*process-counter*")
		reg31678 := MakeNumber(0)
		reg31679 := PrimSet(reg31677, reg31678)
		_ = reg31679
		reg31680 := MakeSymbol("shen.*varcounter*")
		reg31681 := MakeNumber(10000)
		reg31682 := Call(__e, __defun__vector, reg31681)
		reg31683 := PrimSet(reg31680, reg31682)
		_ = reg31683
		reg31684 := MakeSymbol("shen.*prologvectors*")
		reg31685 := MakeNumber(10000)
		reg31686 := Call(__e, __defun__vector, reg31685)
		reg31687 := PrimSet(reg31684, reg31686)
		_ = reg31687
		reg31688 := MakeSymbol("shen.*demodulation-function*")
		reg31689 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.Return(X)
			return
		}, 1)
		reg31690 := PrimSet(reg31688, reg31689)
		_ = reg31690
		reg31691 := MakeSymbol("shen.*macroreg*")
		reg31692 := MakeSymbol("shen.timer-macro")
		reg31693 := MakeSymbol("shen.cases-macro")
		reg31694 := MakeSymbol("shen.abs-macro")
		reg31695 := MakeSymbol("shen.put/get-macro")
		reg31696 := MakeSymbol("shen.compile-macro")
		reg31697 := MakeSymbol("shen.datatype-macro")
		reg31698 := MakeSymbol("shen.let-macro")
		reg31699 := MakeSymbol("shen.assoc-macro")
		reg31700 := MakeSymbol("shen.make-string-macro")
		reg31701 := MakeSymbol("shen.output-macro")
		reg31702 := MakeSymbol("shen.input-macro")
		reg31703 := MakeSymbol("shen.error-macro")
		reg31704 := MakeSymbol("shen.prolog-macro")
		reg31705 := MakeSymbol("shen.synonyms-macro")
		reg31706 := MakeSymbol("shen.nl-macro")
		reg31707 := MakeSymbol("shen.@s-macro")
		reg31708 := MakeSymbol("shen.defprolog-macro")
		reg31709 := MakeSymbol("shen.function-macro")
		reg31710 := Nil
		reg31711 := PrimCons(reg31709, reg31710)
		reg31712 := PrimCons(reg31708, reg31711)
		reg31713 := PrimCons(reg31707, reg31712)
		reg31714 := PrimCons(reg31706, reg31713)
		reg31715 := PrimCons(reg31705, reg31714)
		reg31716 := PrimCons(reg31704, reg31715)
		reg31717 := PrimCons(reg31703, reg31716)
		reg31718 := PrimCons(reg31702, reg31717)
		reg31719 := PrimCons(reg31701, reg31718)
		reg31720 := PrimCons(reg31700, reg31719)
		reg31721 := PrimCons(reg31699, reg31720)
		reg31722 := PrimCons(reg31698, reg31721)
		reg31723 := PrimCons(reg31697, reg31722)
		reg31724 := PrimCons(reg31696, reg31723)
		reg31725 := PrimCons(reg31695, reg31724)
		reg31726 := PrimCons(reg31694, reg31725)
		reg31727 := PrimCons(reg31693, reg31726)
		reg31728 := PrimCons(reg31692, reg31727)
		reg31729 := PrimSet(reg31691, reg31728)
		_ = reg31729
		reg31730 := MakeSymbol("*macros*")
		reg31731 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4timer_1macro, X)
			return
		}, 1)
		reg31733 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4cases_1macro, X)
			return
		}, 1)
		reg31735 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4abs_1macro, X)
			return
		}, 1)
		reg31737 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4put_cget_1macro, X)
			return
		}, 1)
		reg31739 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4compile_1macro, X)
			return
		}, 1)
		reg31741 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4datatype_1macro, X)
			return
		}, 1)
		reg31743 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4let_1macro, X)
			return
		}, 1)
		reg31745 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4assoc_1macro, X)
			return
		}, 1)
		reg31747 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4make_1string_1macro, X)
			return
		}, 1)
		reg31749 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4output_1macro, X)
			return
		}, 1)
		reg31751 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4input_1macro, X)
			return
		}, 1)
		reg31753 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4error_1macro, X)
			return
		}, 1)
		reg31755 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4prolog_1macro, X)
			return
		}, 1)
		reg31757 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4synonyms_1macro, X)
			return
		}, 1)
		reg31759 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4nl_1macro, X)
			return
		}, 1)
		reg31761 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4_8s_1macro, X)
			return
		}, 1)
		reg31763 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4defprolog_1macro, X)
			return
		}, 1)
		reg31765 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4function_1macro, X)
			return
		}, 1)
		reg31767 := Nil
		reg31768 := PrimCons(reg31765, reg31767)
		reg31769 := PrimCons(reg31763, reg31768)
		reg31770 := PrimCons(reg31761, reg31769)
		reg31771 := PrimCons(reg31759, reg31770)
		reg31772 := PrimCons(reg31757, reg31771)
		reg31773 := PrimCons(reg31755, reg31772)
		reg31774 := PrimCons(reg31753, reg31773)
		reg31775 := PrimCons(reg31751, reg31774)
		reg31776 := PrimCons(reg31749, reg31775)
		reg31777 := PrimCons(reg31747, reg31776)
		reg31778 := PrimCons(reg31745, reg31777)
		reg31779 := PrimCons(reg31743, reg31778)
		reg31780 := PrimCons(reg31741, reg31779)
		reg31781 := PrimCons(reg31739, reg31780)
		reg31782 := PrimCons(reg31737, reg31781)
		reg31783 := PrimCons(reg31735, reg31782)
		reg31784 := PrimCons(reg31733, reg31783)
		reg31785 := PrimCons(reg31731, reg31784)
		reg31786 := PrimSet(reg31730, reg31785)
		_ = reg31786
		reg31787 := MakeSymbol("shen.*gensym*")
		reg31788 := MakeNumber(0)
		reg31789 := PrimSet(reg31787, reg31788)
		_ = reg31789
		reg31790 := MakeSymbol("shen.*tracking*")
		reg31791 := Nil
		reg31792 := PrimSet(reg31790, reg31791)
		_ = reg31792
		reg31793 := MakeSymbol("shen.*alphabet*")
		reg31794 := MakeSymbol("A")
		reg31795 := MakeSymbol("B")
		reg31796 := MakeSymbol("C")
		reg31797 := MakeSymbol("D")
		reg31798 := MakeSymbol("E")
		reg31799 := MakeSymbol("F")
		reg31800 := MakeSymbol("G")
		reg31801 := MakeSymbol("H")
		reg31802 := MakeSymbol("I")
		reg31803 := MakeSymbol("J")
		reg31804 := MakeSymbol("K")
		reg31805 := MakeSymbol("L")
		reg31806 := MakeSymbol("M")
		reg31807 := MakeSymbol("N")
		reg31808 := MakeSymbol("O")
		reg31809 := MakeSymbol("P")
		reg31810 := MakeSymbol("Q")
		reg31811 := MakeSymbol("R")
		reg31812 := MakeSymbol("S")
		reg31813 := MakeSymbol("T")
		reg31814 := MakeSymbol("U")
		reg31815 := MakeSymbol("V")
		reg31816 := MakeSymbol("W")
		reg31817 := MakeSymbol("X")
		reg31818 := MakeSymbol("Y")
		reg31819 := MakeSymbol("Z")
		reg31820 := Nil
		reg31821 := PrimCons(reg31819, reg31820)
		reg31822 := PrimCons(reg31818, reg31821)
		reg31823 := PrimCons(reg31817, reg31822)
		reg31824 := PrimCons(reg31816, reg31823)
		reg31825 := PrimCons(reg31815, reg31824)
		reg31826 := PrimCons(reg31814, reg31825)
		reg31827 := PrimCons(reg31813, reg31826)
		reg31828 := PrimCons(reg31812, reg31827)
		reg31829 := PrimCons(reg31811, reg31828)
		reg31830 := PrimCons(reg31810, reg31829)
		reg31831 := PrimCons(reg31809, reg31830)
		reg31832 := PrimCons(reg31808, reg31831)
		reg31833 := PrimCons(reg31807, reg31832)
		reg31834 := PrimCons(reg31806, reg31833)
		reg31835 := PrimCons(reg31805, reg31834)
		reg31836 := PrimCons(reg31804, reg31835)
		reg31837 := PrimCons(reg31803, reg31836)
		reg31838 := PrimCons(reg31802, reg31837)
		reg31839 := PrimCons(reg31801, reg31838)
		reg31840 := PrimCons(reg31800, reg31839)
		reg31841 := PrimCons(reg31799, reg31840)
		reg31842 := PrimCons(reg31798, reg31841)
		reg31843 := PrimCons(reg31797, reg31842)
		reg31844 := PrimCons(reg31796, reg31843)
		reg31845 := PrimCons(reg31795, reg31844)
		reg31846 := PrimCons(reg31794, reg31845)
		reg31847 := PrimSet(reg31793, reg31846)
		_ = reg31847
		reg31848 := MakeSymbol("shen.*special*")
		reg31849 := MakeSymbol("@p")
		reg31850 := MakeSymbol("@s")
		reg31851 := MakeSymbol("@v")
		reg31852 := MakeSymbol("cons")
		reg31853 := MakeSymbol("lambda")
		reg31854 := MakeSymbol("let")
		reg31855 := MakeSymbol("where")
		reg31856 := MakeSymbol("set")
		reg31857 := MakeSymbol("open")
		reg31858 := Nil
		reg31859 := PrimCons(reg31857, reg31858)
		reg31860 := PrimCons(reg31856, reg31859)
		reg31861 := PrimCons(reg31855, reg31860)
		reg31862 := PrimCons(reg31854, reg31861)
		reg31863 := PrimCons(reg31853, reg31862)
		reg31864 := PrimCons(reg31852, reg31863)
		reg31865 := PrimCons(reg31851, reg31864)
		reg31866 := PrimCons(reg31850, reg31865)
		reg31867 := PrimCons(reg31849, reg31866)
		reg31868 := PrimSet(reg31848, reg31867)
		_ = reg31868
		reg31869 := MakeSymbol("shen.*extraspecial*")
		reg31870 := MakeSymbol("define")
		reg31871 := MakeSymbol("shen.process-datatype")
		reg31872 := MakeSymbol("input+")
		reg31873 := MakeSymbol("defcc")
		reg31874 := MakeSymbol("shen.read+")
		reg31875 := MakeSymbol("defmacro")
		reg31876 := Nil
		reg31877 := PrimCons(reg31875, reg31876)
		reg31878 := PrimCons(reg31874, reg31877)
		reg31879 := PrimCons(reg31873, reg31878)
		reg31880 := PrimCons(reg31872, reg31879)
		reg31881 := PrimCons(reg31871, reg31880)
		reg31882 := PrimCons(reg31870, reg31881)
		reg31883 := PrimSet(reg31869, reg31882)
		_ = reg31883
		reg31884 := MakeSymbol("shen.*spy*")
		reg31885 := False
		reg31886 := PrimSet(reg31884, reg31885)
		_ = reg31886
		reg31887 := MakeSymbol("shen.*datatypes*")
		reg31888 := Nil
		reg31889 := PrimSet(reg31887, reg31888)
		_ = reg31889
		reg31890 := MakeSymbol("shen.*alldatatypes*")
		reg31891 := Nil
		reg31892 := PrimSet(reg31890, reg31891)
		_ = reg31892
		reg31893 := MakeSymbol("shen.*shen-type-theory-enabled?*")
		reg31894 := True
		reg31895 := PrimSet(reg31893, reg31894)
		_ = reg31895
		reg31896 := MakeSymbol("shen.*synonyms*")
		reg31897 := Nil
		reg31898 := PrimSet(reg31896, reg31897)
		_ = reg31898
		reg31899 := MakeSymbol("shen.*system*")
		reg31900 := Nil
		reg31901 := PrimSet(reg31899, reg31900)
		_ = reg31901
		reg31902 := MakeSymbol("shen.*maxcomplexity*")
		reg31903 := MakeNumber(128)
		reg31904 := PrimSet(reg31902, reg31903)
		_ = reg31904
		reg31905 := MakeSymbol("shen.*occurs*")
		reg31906 := True
		reg31907 := PrimSet(reg31905, reg31906)
		_ = reg31907
		reg31908 := MakeSymbol("shen.*maxinferences*")
		reg31909 := MakeNumber(1000000)
		reg31910 := PrimSet(reg31908, reg31909)
		_ = reg31910
		reg31911 := MakeSymbol("*maximum-print-sequence-size*")
		reg31912 := MakeNumber(20)
		reg31913 := PrimSet(reg31911, reg31912)
		_ = reg31913
		reg31914 := MakeSymbol("shen.*catch*")
		reg31915 := MakeNumber(0)
		reg31916 := PrimSet(reg31914, reg31915)
		_ = reg31916
		reg31917 := MakeSymbol("shen.*call*")
		reg31918 := MakeNumber(0)
		reg31919 := PrimSet(reg31917, reg31918)
		_ = reg31919
		reg31920 := MakeSymbol("shen.*infs*")
		reg31921 := MakeNumber(0)
		reg31922 := PrimSet(reg31920, reg31921)
		_ = reg31922
		reg31923 := MakeSymbol("*hush*")
		reg31924 := False
		reg31925 := PrimSet(reg31923, reg31924)
		_ = reg31925
		reg31926 := MakeSymbol("shen.*optimise*")
		reg31927 := False
		reg31928 := PrimSet(reg31926, reg31927)
		_ = reg31928
		reg31929 := MakeSymbol("*version*")
		reg31930 := MakeString("Shen 22.2")
		reg31931 := PrimSet(reg31929, reg31930)
		_ = reg31931
		reg31932 := MakeSymbol("*home-directory*")
		reg31933 := Call(__e, __defun__bound_2, reg31932)
		reg31934 := PrimNot(reg31933)
		var reg31939 Obj
		if reg31934 == True {
			reg31935 := MakeSymbol("*home-directory*")
			reg31936 := MakeString("")
			reg31937 := PrimSet(reg31935, reg31936)
			reg31939 = reg31937
		} else {
			reg31938 := MakeSymbol("shen.skip")
			reg31939 = reg31938
		}
		_ = reg31939
		reg31940 := MakeSymbol("*sterror*")
		reg31941 := Call(__e, __defun__bound_2, reg31940)
		reg31942 := PrimNot(reg31941)
		var reg31948 Obj
		if reg31942 == True {
			reg31943 := MakeSymbol("*sterror*")
			reg31944 := MakeSymbol("*stoutput*")
			reg31945 := PrimValue(reg31944)
			reg31946 := PrimSet(reg31943, reg31945)
			reg31948 = reg31946
		} else {
			reg31947 := MakeSymbol("shen.skip")
			reg31948 = reg31947
		}
		_ = reg31948
		reg31949 := MakeSymbol("abort")
		reg31950 := MakeNumber(0)
		reg31951 := MakeSymbol("absvector?")
		reg31952 := MakeNumber(1)
		reg31953 := MakeSymbol("absvector")
		reg31954 := MakeNumber(1)
		reg31955 := MakeSymbol("adjoin")
		reg31956 := MakeNumber(2)
		reg31957 := MakeSymbol("and")
		reg31958 := MakeNumber(2)
		reg31959 := MakeSymbol("append")
		reg31960 := MakeNumber(2)
		reg31961 := MakeSymbol("arity")
		reg31962 := MakeNumber(1)
		reg31963 := MakeSymbol("assoc")
		reg31964 := MakeNumber(2)
		reg31965 := MakeSymbol("boolean?")
		reg31966 := MakeNumber(1)
		reg31967 := MakeSymbol("bound?")
		reg31968 := MakeNumber(1)
		reg31969 := MakeSymbol("cd")
		reg31970 := MakeNumber(1)
		reg31971 := MakeSymbol("close")
		reg31972 := MakeNumber(1)
		reg31973 := MakeSymbol("compile")
		reg31974 := MakeNumber(3)
		reg31975 := MakeSymbol("concat")
		reg31976 := MakeNumber(2)
		reg31977 := MakeSymbol("cons")
		reg31978 := MakeNumber(2)
		reg31979 := MakeSymbol("cons?")
		reg31980 := MakeNumber(1)
		reg31981 := MakeSymbol("cn")
		reg31982 := MakeNumber(2)
		reg31983 := MakeSymbol("declare")
		reg31984 := MakeNumber(2)
		reg31985 := MakeSymbol("destroy")
		reg31986 := MakeNumber(1)
		reg31987 := MakeSymbol("difference")
		reg31988 := MakeNumber(2)
		reg31989 := MakeSymbol("do")
		reg31990 := MakeNumber(2)
		reg31991 := MakeSymbol("element?")
		reg31992 := MakeNumber(2)
		reg31993 := MakeSymbol("empty?")
		reg31994 := MakeNumber(1)
		reg31995 := MakeSymbol("enable-type-theory")
		reg31996 := MakeNumber(1)
		reg31997 := MakeSymbol("error-to-string")
		reg31998 := MakeNumber(1)
		reg31999 := MakeSymbol("shen.interror")
		reg32000 := MakeNumber(2)
		reg32001 := MakeSymbol("eval")
		reg32002 := MakeNumber(1)
		reg32003 := MakeSymbol("eval-kl")
		reg32004 := MakeNumber(1)
		reg32005 := MakeSymbol("explode")
		reg32006 := MakeNumber(1)
		reg32007 := MakeSymbol("external")
		reg32008 := MakeNumber(1)
		reg32009 := MakeSymbol("fail-if")
		reg32010 := MakeNumber(2)
		reg32011 := MakeSymbol("fail")
		reg32012 := MakeNumber(0)
		reg32013 := MakeSymbol("fix")
		reg32014 := MakeNumber(2)
		reg32015 := MakeSymbol("findall")
		reg32016 := MakeNumber(5)
		reg32017 := MakeSymbol("freeze")
		reg32018 := MakeNumber(1)
		reg32019 := MakeSymbol("fst")
		reg32020 := MakeNumber(1)
		reg32021 := MakeSymbol("gensym")
		reg32022 := MakeNumber(1)
		reg32023 := MakeSymbol("get")
		reg32024 := MakeNumber(3)
		reg32025 := MakeSymbol("get-time")
		reg32026 := MakeNumber(1)
		reg32027 := MakeSymbol("address->")
		reg32028 := MakeNumber(3)
		reg32029 := MakeSymbol("<-address")
		reg32030 := MakeNumber(2)
		reg32031 := MakeSymbol("<-vector")
		reg32032 := MakeNumber(2)
		reg32033 := MakeSymbol(">")
		reg32034 := MakeNumber(2)
		reg32035 := MakeSymbol(">=")
		reg32036 := MakeNumber(2)
		reg32037 := MakeSymbol("=")
		reg32038 := MakeNumber(2)
		reg32039 := MakeSymbol("hash")
		reg32040 := MakeNumber(2)
		reg32041 := MakeSymbol("hd")
		reg32042 := MakeNumber(1)
		reg32043 := MakeSymbol("hdv")
		reg32044 := MakeNumber(1)
		reg32045 := MakeSymbol("hdstr")
		reg32046 := MakeNumber(1)
		reg32047 := MakeSymbol("head")
		reg32048 := MakeNumber(1)
		reg32049 := MakeSymbol("if")
		reg32050 := MakeNumber(3)
		reg32051 := MakeSymbol("integer?")
		reg32052 := MakeNumber(1)
		reg32053 := MakeSymbol("intern")
		reg32054 := MakeNumber(1)
		reg32055 := MakeSymbol("identical")
		reg32056 := MakeNumber(4)
		reg32057 := MakeSymbol("inferences")
		reg32058 := MakeNumber(0)
		reg32059 := MakeSymbol("input")
		reg32060 := MakeNumber(1)
		reg32061 := MakeSymbol("input+")
		reg32062 := MakeNumber(2)
		reg32063 := MakeSymbol("implementation")
		reg32064 := MakeNumber(0)
		reg32065 := MakeSymbol("intersection")
		reg32066 := MakeNumber(2)
		reg32067 := MakeSymbol("internal")
		reg32068 := MakeNumber(1)
		reg32069 := MakeSymbol("it")
		reg32070 := MakeNumber(0)
		reg32071 := MakeSymbol("kill")
		reg32072 := MakeNumber(0)
		reg32073 := MakeSymbol("language")
		reg32074 := MakeNumber(0)
		reg32075 := MakeSymbol("length")
		reg32076 := MakeNumber(1)
		reg32077 := MakeSymbol("limit")
		reg32078 := MakeNumber(1)
		reg32079 := MakeSymbol("lineread")
		reg32080 := MakeNumber(1)
		reg32081 := MakeSymbol("load")
		reg32082 := MakeNumber(1)
		reg32083 := MakeSymbol("<")
		reg32084 := MakeNumber(2)
		reg32085 := MakeSymbol("<=")
		reg32086 := MakeNumber(2)
		reg32087 := MakeSymbol("vector")
		reg32088 := MakeNumber(1)
		reg32089 := MakeSymbol("macroexpand")
		reg32090 := MakeNumber(1)
		reg32091 := MakeSymbol("map")
		reg32092 := MakeNumber(2)
		reg32093 := MakeSymbol("mapcan")
		reg32094 := MakeNumber(2)
		reg32095 := MakeSymbol("maxinferences")
		reg32096 := MakeNumber(1)
		reg32097 := MakeSymbol("nl")
		reg32098 := MakeNumber(1)
		reg32099 := MakeSymbol("not")
		reg32100 := MakeNumber(1)
		reg32101 := MakeSymbol("nth")
		reg32102 := MakeNumber(2)
		reg32103 := MakeSymbol("n->string")
		reg32104 := MakeNumber(1)
		reg32105 := MakeSymbol("number?")
		reg32106 := MakeNumber(1)
		reg32107 := MakeSymbol("occurs-check")
		reg32108 := MakeNumber(1)
		reg32109 := MakeSymbol("occurrences")
		reg32110 := MakeNumber(2)
		reg32111 := MakeSymbol("occurs-check")
		reg32112 := MakeNumber(1)
		reg32113 := MakeSymbol("open")
		reg32114 := MakeNumber(2)
		reg32115 := MakeSymbol("optimise")
		reg32116 := MakeNumber(1)
		reg32117 := MakeSymbol("or")
		reg32118 := MakeNumber(2)
		reg32119 := MakeSymbol("os")
		reg32120 := MakeNumber(0)
		reg32121 := MakeSymbol("package")
		reg32122 := MakeNumber(3)
		reg32123 := MakeSymbol("package?")
		reg32124 := MakeNumber(1)
		reg32125 := MakeSymbol("port")
		reg32126 := MakeNumber(0)
		reg32127 := MakeSymbol("porters")
		reg32128 := MakeNumber(0)
		reg32129 := MakeSymbol("pos")
		reg32130 := MakeNumber(2)
		reg32131 := MakeSymbol("print")
		reg32132 := MakeNumber(1)
		reg32133 := MakeSymbol("profile")
		reg32134 := MakeNumber(1)
		reg32135 := MakeSymbol("profile-results")
		reg32136 := MakeNumber(1)
		reg32137 := MakeSymbol("pr")
		reg32138 := MakeNumber(2)
		reg32139 := MakeSymbol("ps")
		reg32140 := MakeNumber(1)
		reg32141 := MakeSymbol("preclude")
		reg32142 := MakeNumber(1)
		reg32143 := MakeSymbol("preclude-all-but")
		reg32144 := MakeNumber(1)
		reg32145 := MakeSymbol("protect")
		reg32146 := MakeNumber(1)
		reg32147 := MakeSymbol("address->")
		reg32148 := MakeNumber(3)
		reg32149 := MakeSymbol("put")
		reg32150 := MakeNumber(4)
		reg32151 := MakeSymbol("shen.reassemble")
		reg32152 := MakeNumber(2)
		reg32153 := MakeSymbol("read-file-as-string")
		reg32154 := MakeNumber(1)
		reg32155 := MakeSymbol("read-file")
		reg32156 := MakeNumber(1)
		reg32157 := MakeSymbol("read-file-as-bytelist")
		reg32158 := MakeNumber(1)
		reg32159 := MakeSymbol("read")
		reg32160 := MakeNumber(1)
		reg32161 := MakeSymbol("read-byte")
		reg32162 := MakeNumber(1)
		reg32163 := MakeSymbol("read-from-string")
		reg32164 := MakeNumber(1)
		reg32165 := MakeSymbol("receive")
		reg32166 := MakeNumber(1)
		reg32167 := MakeSymbol("release")
		reg32168 := MakeNumber(0)
		reg32169 := MakeSymbol("remove")
		reg32170 := MakeNumber(2)
		reg32171 := MakeSymbol("shen.require")
		reg32172 := MakeNumber(3)
		reg32173 := MakeSymbol("reverse")
		reg32174 := MakeNumber(1)
		reg32175 := MakeSymbol("set")
		reg32176 := MakeNumber(2)
		reg32177 := MakeSymbol("simple-error")
		reg32178 := MakeNumber(1)
		reg32179 := MakeSymbol("snd")
		reg32180 := MakeNumber(1)
		reg32181 := MakeSymbol("specialise")
		reg32182 := MakeNumber(1)
		reg32183 := MakeSymbol("spy")
		reg32184 := MakeNumber(1)
		reg32185 := MakeSymbol("step")
		reg32186 := MakeNumber(1)
		reg32187 := MakeSymbol("stinput")
		reg32188 := MakeNumber(0)
		reg32189 := MakeSymbol("stoutput")
		reg32190 := MakeNumber(0)
		reg32191 := MakeSymbol("sterror")
		reg32192 := MakeNumber(0)
		reg32193 := MakeSymbol("string->n")
		reg32194 := MakeNumber(1)
		reg32195 := MakeSymbol("string->symbol")
		reg32196 := MakeNumber(1)
		reg32197 := MakeSymbol("string?")
		reg32198 := MakeNumber(1)
		reg32199 := MakeSymbol("str")
		reg32200 := MakeNumber(1)
		reg32201 := MakeSymbol("subst")
		reg32202 := MakeNumber(3)
		reg32203 := MakeSymbol("sum")
		reg32204 := MakeNumber(1)
		reg32205 := MakeSymbol("symbol?")
		reg32206 := MakeNumber(1)
		reg32207 := MakeSymbol("systemf")
		reg32208 := MakeNumber(1)
		reg32209 := MakeSymbol("tail")
		reg32210 := MakeNumber(1)
		reg32211 := MakeSymbol("tl")
		reg32212 := MakeNumber(1)
		reg32213 := MakeSymbol("tc")
		reg32214 := MakeNumber(1)
		reg32215 := MakeSymbol("tc?")
		reg32216 := MakeNumber(0)
		reg32217 := MakeSymbol("thaw")
		reg32218 := MakeNumber(1)
		reg32219 := MakeSymbol("tlstr")
		reg32220 := MakeNumber(1)
		reg32221 := MakeSymbol("track")
		reg32222 := MakeNumber(1)
		reg32223 := MakeSymbol("trap-error")
		reg32224 := MakeNumber(2)
		reg32225 := MakeSymbol("tuple?")
		reg32226 := MakeNumber(1)
		reg32227 := MakeSymbol("type")
		reg32228 := MakeNumber(2)
		reg32229 := MakeSymbol("return")
		reg32230 := MakeNumber(3)
		reg32231 := MakeSymbol("undefmacro")
		reg32232 := MakeNumber(1)
		reg32233 := MakeSymbol("unput")
		reg32234 := MakeNumber(3)
		reg32235 := MakeSymbol("unprofile")
		reg32236 := MakeNumber(1)
		reg32237 := MakeSymbol("unify")
		reg32238 := MakeNumber(4)
		reg32239 := MakeSymbol("unify!")
		reg32240 := MakeNumber(4)
		reg32241 := MakeSymbol("union")
		reg32242 := MakeNumber(2)
		reg32243 := MakeSymbol("untrack")
		reg32244 := MakeNumber(1)
		reg32245 := MakeSymbol("unspecialise")
		reg32246 := MakeNumber(1)
		reg32247 := MakeSymbol("undefmacro")
		reg32248 := MakeNumber(1)
		reg32249 := MakeSymbol("vector")
		reg32250 := MakeNumber(1)
		reg32251 := MakeSymbol("vector?")
		reg32252 := MakeNumber(1)
		reg32253 := MakeSymbol("vector->")
		reg32254 := MakeNumber(3)
		reg32255 := MakeSymbol("value")
		reg32256 := MakeNumber(1)
		reg32257 := MakeSymbol("variable?")
		reg32258 := MakeNumber(1)
		reg32259 := MakeSymbol("version")
		reg32260 := MakeNumber(0)
		reg32261 := MakeSymbol("write-byte")
		reg32262 := MakeNumber(2)
		reg32263 := MakeSymbol("write-to-file")
		reg32264 := MakeNumber(2)
		reg32265 := MakeSymbol("y-or-n?")
		reg32266 := MakeNumber(1)
		reg32267 := MakeSymbol("+")
		reg32268 := MakeNumber(2)
		reg32269 := MakeSymbol("*")
		reg32270 := MakeNumber(2)
		reg32271 := MakeSymbol("/")
		reg32272 := MakeNumber(2)
		reg32273 := MakeSymbol("-")
		reg32274 := MakeNumber(2)
		reg32275 := MakeSymbol("==")
		reg32276 := MakeNumber(2)
		reg32277 := MakeSymbol("<e>")
		reg32278 := MakeNumber(1)
		reg32279 := MakeSymbol("<!>")
		reg32280 := MakeNumber(1)
		reg32281 := MakeSymbol("@p")
		reg32282 := MakeNumber(2)
		reg32283 := MakeSymbol("@v")
		reg32284 := MakeNumber(2)
		reg32285 := MakeSymbol("@s")
		reg32286 := MakeNumber(2)
		reg32287 := MakeSymbol("preclude")
		reg32288 := MakeNumber(1)
		reg32289 := MakeSymbol("include")
		reg32290 := MakeNumber(1)
		reg32291 := MakeSymbol("preclude-all-but")
		reg32292 := MakeNumber(1)
		reg32293 := MakeSymbol("include-all-but")
		reg32294 := MakeNumber(1)
		reg32295 := Nil
		reg32296 := PrimCons(reg32294, reg32295)
		reg32297 := PrimCons(reg32293, reg32296)
		reg32298 := PrimCons(reg32292, reg32297)
		reg32299 := PrimCons(reg32291, reg32298)
		reg32300 := PrimCons(reg32290, reg32299)
		reg32301 := PrimCons(reg32289, reg32300)
		reg32302 := PrimCons(reg32288, reg32301)
		reg32303 := PrimCons(reg32287, reg32302)
		reg32304 := PrimCons(reg32286, reg32303)
		reg32305 := PrimCons(reg32285, reg32304)
		reg32306 := PrimCons(reg32284, reg32305)
		reg32307 := PrimCons(reg32283, reg32306)
		reg32308 := PrimCons(reg32282, reg32307)
		reg32309 := PrimCons(reg32281, reg32308)
		reg32310 := PrimCons(reg32280, reg32309)
		reg32311 := PrimCons(reg32279, reg32310)
		reg32312 := PrimCons(reg32278, reg32311)
		reg32313 := PrimCons(reg32277, reg32312)
		reg32314 := PrimCons(reg32276, reg32313)
		reg32315 := PrimCons(reg32275, reg32314)
		reg32316 := PrimCons(reg32274, reg32315)
		reg32317 := PrimCons(reg32273, reg32316)
		reg32318 := PrimCons(reg32272, reg32317)
		reg32319 := PrimCons(reg32271, reg32318)
		reg32320 := PrimCons(reg32270, reg32319)
		reg32321 := PrimCons(reg32269, reg32320)
		reg32322 := PrimCons(reg32268, reg32321)
		reg32323 := PrimCons(reg32267, reg32322)
		reg32324 := PrimCons(reg32266, reg32323)
		reg32325 := PrimCons(reg32265, reg32324)
		reg32326 := PrimCons(reg32264, reg32325)
		reg32327 := PrimCons(reg32263, reg32326)
		reg32328 := PrimCons(reg32262, reg32327)
		reg32329 := PrimCons(reg32261, reg32328)
		reg32330 := PrimCons(reg32260, reg32329)
		reg32331 := PrimCons(reg32259, reg32330)
		reg32332 := PrimCons(reg32258, reg32331)
		reg32333 := PrimCons(reg32257, reg32332)
		reg32334 := PrimCons(reg32256, reg32333)
		reg32335 := PrimCons(reg32255, reg32334)
		reg32336 := PrimCons(reg32254, reg32335)
		reg32337 := PrimCons(reg32253, reg32336)
		reg32338 := PrimCons(reg32252, reg32337)
		reg32339 := PrimCons(reg32251, reg32338)
		reg32340 := PrimCons(reg32250, reg32339)
		reg32341 := PrimCons(reg32249, reg32340)
		reg32342 := PrimCons(reg32248, reg32341)
		reg32343 := PrimCons(reg32247, reg32342)
		reg32344 := PrimCons(reg32246, reg32343)
		reg32345 := PrimCons(reg32245, reg32344)
		reg32346 := PrimCons(reg32244, reg32345)
		reg32347 := PrimCons(reg32243, reg32346)
		reg32348 := PrimCons(reg32242, reg32347)
		reg32349 := PrimCons(reg32241, reg32348)
		reg32350 := PrimCons(reg32240, reg32349)
		reg32351 := PrimCons(reg32239, reg32350)
		reg32352 := PrimCons(reg32238, reg32351)
		reg32353 := PrimCons(reg32237, reg32352)
		reg32354 := PrimCons(reg32236, reg32353)
		reg32355 := PrimCons(reg32235, reg32354)
		reg32356 := PrimCons(reg32234, reg32355)
		reg32357 := PrimCons(reg32233, reg32356)
		reg32358 := PrimCons(reg32232, reg32357)
		reg32359 := PrimCons(reg32231, reg32358)
		reg32360 := PrimCons(reg32230, reg32359)
		reg32361 := PrimCons(reg32229, reg32360)
		reg32362 := PrimCons(reg32228, reg32361)
		reg32363 := PrimCons(reg32227, reg32362)
		reg32364 := PrimCons(reg32226, reg32363)
		reg32365 := PrimCons(reg32225, reg32364)
		reg32366 := PrimCons(reg32224, reg32365)
		reg32367 := PrimCons(reg32223, reg32366)
		reg32368 := PrimCons(reg32222, reg32367)
		reg32369 := PrimCons(reg32221, reg32368)
		reg32370 := PrimCons(reg32220, reg32369)
		reg32371 := PrimCons(reg32219, reg32370)
		reg32372 := PrimCons(reg32218, reg32371)
		reg32373 := PrimCons(reg32217, reg32372)
		reg32374 := PrimCons(reg32216, reg32373)
		reg32375 := PrimCons(reg32215, reg32374)
		reg32376 := PrimCons(reg32214, reg32375)
		reg32377 := PrimCons(reg32213, reg32376)
		reg32378 := PrimCons(reg32212, reg32377)
		reg32379 := PrimCons(reg32211, reg32378)
		reg32380 := PrimCons(reg32210, reg32379)
		reg32381 := PrimCons(reg32209, reg32380)
		reg32382 := PrimCons(reg32208, reg32381)
		reg32383 := PrimCons(reg32207, reg32382)
		reg32384 := PrimCons(reg32206, reg32383)
		reg32385 := PrimCons(reg32205, reg32384)
		reg32386 := PrimCons(reg32204, reg32385)
		reg32387 := PrimCons(reg32203, reg32386)
		reg32388 := PrimCons(reg32202, reg32387)
		reg32389 := PrimCons(reg32201, reg32388)
		reg32390 := PrimCons(reg32200, reg32389)
		reg32391 := PrimCons(reg32199, reg32390)
		reg32392 := PrimCons(reg32198, reg32391)
		reg32393 := PrimCons(reg32197, reg32392)
		reg32394 := PrimCons(reg32196, reg32393)
		reg32395 := PrimCons(reg32195, reg32394)
		reg32396 := PrimCons(reg32194, reg32395)
		reg32397 := PrimCons(reg32193, reg32396)
		reg32398 := PrimCons(reg32192, reg32397)
		reg32399 := PrimCons(reg32191, reg32398)
		reg32400 := PrimCons(reg32190, reg32399)
		reg32401 := PrimCons(reg32189, reg32400)
		reg32402 := PrimCons(reg32188, reg32401)
		reg32403 := PrimCons(reg32187, reg32402)
		reg32404 := PrimCons(reg32186, reg32403)
		reg32405 := PrimCons(reg32185, reg32404)
		reg32406 := PrimCons(reg32184, reg32405)
		reg32407 := PrimCons(reg32183, reg32406)
		reg32408 := PrimCons(reg32182, reg32407)
		reg32409 := PrimCons(reg32181, reg32408)
		reg32410 := PrimCons(reg32180, reg32409)
		reg32411 := PrimCons(reg32179, reg32410)
		reg32412 := PrimCons(reg32178, reg32411)
		reg32413 := PrimCons(reg32177, reg32412)
		reg32414 := PrimCons(reg32176, reg32413)
		reg32415 := PrimCons(reg32175, reg32414)
		reg32416 := PrimCons(reg32174, reg32415)
		reg32417 := PrimCons(reg32173, reg32416)
		reg32418 := PrimCons(reg32172, reg32417)
		reg32419 := PrimCons(reg32171, reg32418)
		reg32420 := PrimCons(reg32170, reg32419)
		reg32421 := PrimCons(reg32169, reg32420)
		reg32422 := PrimCons(reg32168, reg32421)
		reg32423 := PrimCons(reg32167, reg32422)
		reg32424 := PrimCons(reg32166, reg32423)
		reg32425 := PrimCons(reg32165, reg32424)
		reg32426 := PrimCons(reg32164, reg32425)
		reg32427 := PrimCons(reg32163, reg32426)
		reg32428 := PrimCons(reg32162, reg32427)
		reg32429 := PrimCons(reg32161, reg32428)
		reg32430 := PrimCons(reg32160, reg32429)
		reg32431 := PrimCons(reg32159, reg32430)
		reg32432 := PrimCons(reg32158, reg32431)
		reg32433 := PrimCons(reg32157, reg32432)
		reg32434 := PrimCons(reg32156, reg32433)
		reg32435 := PrimCons(reg32155, reg32434)
		reg32436 := PrimCons(reg32154, reg32435)
		reg32437 := PrimCons(reg32153, reg32436)
		reg32438 := PrimCons(reg32152, reg32437)
		reg32439 := PrimCons(reg32151, reg32438)
		reg32440 := PrimCons(reg32150, reg32439)
		reg32441 := PrimCons(reg32149, reg32440)
		reg32442 := PrimCons(reg32148, reg32441)
		reg32443 := PrimCons(reg32147, reg32442)
		reg32444 := PrimCons(reg32146, reg32443)
		reg32445 := PrimCons(reg32145, reg32444)
		reg32446 := PrimCons(reg32144, reg32445)
		reg32447 := PrimCons(reg32143, reg32446)
		reg32448 := PrimCons(reg32142, reg32447)
		reg32449 := PrimCons(reg32141, reg32448)
		reg32450 := PrimCons(reg32140, reg32449)
		reg32451 := PrimCons(reg32139, reg32450)
		reg32452 := PrimCons(reg32138, reg32451)
		reg32453 := PrimCons(reg32137, reg32452)
		reg32454 := PrimCons(reg32136, reg32453)
		reg32455 := PrimCons(reg32135, reg32454)
		reg32456 := PrimCons(reg32134, reg32455)
		reg32457 := PrimCons(reg32133, reg32456)
		reg32458 := PrimCons(reg32132, reg32457)
		reg32459 := PrimCons(reg32131, reg32458)
		reg32460 := PrimCons(reg32130, reg32459)
		reg32461 := PrimCons(reg32129, reg32460)
		reg32462 := PrimCons(reg32128, reg32461)
		reg32463 := PrimCons(reg32127, reg32462)
		reg32464 := PrimCons(reg32126, reg32463)
		reg32465 := PrimCons(reg32125, reg32464)
		reg32466 := PrimCons(reg32124, reg32465)
		reg32467 := PrimCons(reg32123, reg32466)
		reg32468 := PrimCons(reg32122, reg32467)
		reg32469 := PrimCons(reg32121, reg32468)
		reg32470 := PrimCons(reg32120, reg32469)
		reg32471 := PrimCons(reg32119, reg32470)
		reg32472 := PrimCons(reg32118, reg32471)
		reg32473 := PrimCons(reg32117, reg32472)
		reg32474 := PrimCons(reg32116, reg32473)
		reg32475 := PrimCons(reg32115, reg32474)
		reg32476 := PrimCons(reg32114, reg32475)
		reg32477 := PrimCons(reg32113, reg32476)
		reg32478 := PrimCons(reg32112, reg32477)
		reg32479 := PrimCons(reg32111, reg32478)
		reg32480 := PrimCons(reg32110, reg32479)
		reg32481 := PrimCons(reg32109, reg32480)
		reg32482 := PrimCons(reg32108, reg32481)
		reg32483 := PrimCons(reg32107, reg32482)
		reg32484 := PrimCons(reg32106, reg32483)
		reg32485 := PrimCons(reg32105, reg32484)
		reg32486 := PrimCons(reg32104, reg32485)
		reg32487 := PrimCons(reg32103, reg32486)
		reg32488 := PrimCons(reg32102, reg32487)
		reg32489 := PrimCons(reg32101, reg32488)
		reg32490 := PrimCons(reg32100, reg32489)
		reg32491 := PrimCons(reg32099, reg32490)
		reg32492 := PrimCons(reg32098, reg32491)
		reg32493 := PrimCons(reg32097, reg32492)
		reg32494 := PrimCons(reg32096, reg32493)
		reg32495 := PrimCons(reg32095, reg32494)
		reg32496 := PrimCons(reg32094, reg32495)
		reg32497 := PrimCons(reg32093, reg32496)
		reg32498 := PrimCons(reg32092, reg32497)
		reg32499 := PrimCons(reg32091, reg32498)
		reg32500 := PrimCons(reg32090, reg32499)
		reg32501 := PrimCons(reg32089, reg32500)
		reg32502 := PrimCons(reg32088, reg32501)
		reg32503 := PrimCons(reg32087, reg32502)
		reg32504 := PrimCons(reg32086, reg32503)
		reg32505 := PrimCons(reg32085, reg32504)
		reg32506 := PrimCons(reg32084, reg32505)
		reg32507 := PrimCons(reg32083, reg32506)
		reg32508 := PrimCons(reg32082, reg32507)
		reg32509 := PrimCons(reg32081, reg32508)
		reg32510 := PrimCons(reg32080, reg32509)
		reg32511 := PrimCons(reg32079, reg32510)
		reg32512 := PrimCons(reg32078, reg32511)
		reg32513 := PrimCons(reg32077, reg32512)
		reg32514 := PrimCons(reg32076, reg32513)
		reg32515 := PrimCons(reg32075, reg32514)
		reg32516 := PrimCons(reg32074, reg32515)
		reg32517 := PrimCons(reg32073, reg32516)
		reg32518 := PrimCons(reg32072, reg32517)
		reg32519 := PrimCons(reg32071, reg32518)
		reg32520 := PrimCons(reg32070, reg32519)
		reg32521 := PrimCons(reg32069, reg32520)
		reg32522 := PrimCons(reg32068, reg32521)
		reg32523 := PrimCons(reg32067, reg32522)
		reg32524 := PrimCons(reg32066, reg32523)
		reg32525 := PrimCons(reg32065, reg32524)
		reg32526 := PrimCons(reg32064, reg32525)
		reg32527 := PrimCons(reg32063, reg32526)
		reg32528 := PrimCons(reg32062, reg32527)
		reg32529 := PrimCons(reg32061, reg32528)
		reg32530 := PrimCons(reg32060, reg32529)
		reg32531 := PrimCons(reg32059, reg32530)
		reg32532 := PrimCons(reg32058, reg32531)
		reg32533 := PrimCons(reg32057, reg32532)
		reg32534 := PrimCons(reg32056, reg32533)
		reg32535 := PrimCons(reg32055, reg32534)
		reg32536 := PrimCons(reg32054, reg32535)
		reg32537 := PrimCons(reg32053, reg32536)
		reg32538 := PrimCons(reg32052, reg32537)
		reg32539 := PrimCons(reg32051, reg32538)
		reg32540 := PrimCons(reg32050, reg32539)
		reg32541 := PrimCons(reg32049, reg32540)
		reg32542 := PrimCons(reg32048, reg32541)
		reg32543 := PrimCons(reg32047, reg32542)
		reg32544 := PrimCons(reg32046, reg32543)
		reg32545 := PrimCons(reg32045, reg32544)
		reg32546 := PrimCons(reg32044, reg32545)
		reg32547 := PrimCons(reg32043, reg32546)
		reg32548 := PrimCons(reg32042, reg32547)
		reg32549 := PrimCons(reg32041, reg32548)
		reg32550 := PrimCons(reg32040, reg32549)
		reg32551 := PrimCons(reg32039, reg32550)
		reg32552 := PrimCons(reg32038, reg32551)
		reg32553 := PrimCons(reg32037, reg32552)
		reg32554 := PrimCons(reg32036, reg32553)
		reg32555 := PrimCons(reg32035, reg32554)
		reg32556 := PrimCons(reg32034, reg32555)
		reg32557 := PrimCons(reg32033, reg32556)
		reg32558 := PrimCons(reg32032, reg32557)
		reg32559 := PrimCons(reg32031, reg32558)
		reg32560 := PrimCons(reg32030, reg32559)
		reg32561 := PrimCons(reg32029, reg32560)
		reg32562 := PrimCons(reg32028, reg32561)
		reg32563 := PrimCons(reg32027, reg32562)
		reg32564 := PrimCons(reg32026, reg32563)
		reg32565 := PrimCons(reg32025, reg32564)
		reg32566 := PrimCons(reg32024, reg32565)
		reg32567 := PrimCons(reg32023, reg32566)
		reg32568 := PrimCons(reg32022, reg32567)
		reg32569 := PrimCons(reg32021, reg32568)
		reg32570 := PrimCons(reg32020, reg32569)
		reg32571 := PrimCons(reg32019, reg32570)
		reg32572 := PrimCons(reg32018, reg32571)
		reg32573 := PrimCons(reg32017, reg32572)
		reg32574 := PrimCons(reg32016, reg32573)
		reg32575 := PrimCons(reg32015, reg32574)
		reg32576 := PrimCons(reg32014, reg32575)
		reg32577 := PrimCons(reg32013, reg32576)
		reg32578 := PrimCons(reg32012, reg32577)
		reg32579 := PrimCons(reg32011, reg32578)
		reg32580 := PrimCons(reg32010, reg32579)
		reg32581 := PrimCons(reg32009, reg32580)
		reg32582 := PrimCons(reg32008, reg32581)
		reg32583 := PrimCons(reg32007, reg32582)
		reg32584 := PrimCons(reg32006, reg32583)
		reg32585 := PrimCons(reg32005, reg32584)
		reg32586 := PrimCons(reg32004, reg32585)
		reg32587 := PrimCons(reg32003, reg32586)
		reg32588 := PrimCons(reg32002, reg32587)
		reg32589 := PrimCons(reg32001, reg32588)
		reg32590 := PrimCons(reg32000, reg32589)
		reg32591 := PrimCons(reg31999, reg32590)
		reg32592 := PrimCons(reg31998, reg32591)
		reg32593 := PrimCons(reg31997, reg32592)
		reg32594 := PrimCons(reg31996, reg32593)
		reg32595 := PrimCons(reg31995, reg32594)
		reg32596 := PrimCons(reg31994, reg32595)
		reg32597 := PrimCons(reg31993, reg32596)
		reg32598 := PrimCons(reg31992, reg32597)
		reg32599 := PrimCons(reg31991, reg32598)
		reg32600 := PrimCons(reg31990, reg32599)
		reg32601 := PrimCons(reg31989, reg32600)
		reg32602 := PrimCons(reg31988, reg32601)
		reg32603 := PrimCons(reg31987, reg32602)
		reg32604 := PrimCons(reg31986, reg32603)
		reg32605 := PrimCons(reg31985, reg32604)
		reg32606 := PrimCons(reg31984, reg32605)
		reg32607 := PrimCons(reg31983, reg32606)
		reg32608 := PrimCons(reg31982, reg32607)
		reg32609 := PrimCons(reg31981, reg32608)
		reg32610 := PrimCons(reg31980, reg32609)
		reg32611 := PrimCons(reg31979, reg32610)
		reg32612 := PrimCons(reg31978, reg32611)
		reg32613 := PrimCons(reg31977, reg32612)
		reg32614 := PrimCons(reg31976, reg32613)
		reg32615 := PrimCons(reg31975, reg32614)
		reg32616 := PrimCons(reg31974, reg32615)
		reg32617 := PrimCons(reg31973, reg32616)
		reg32618 := PrimCons(reg31972, reg32617)
		reg32619 := PrimCons(reg31971, reg32618)
		reg32620 := PrimCons(reg31970, reg32619)
		reg32621 := PrimCons(reg31969, reg32620)
		reg32622 := PrimCons(reg31968, reg32621)
		reg32623 := PrimCons(reg31967, reg32622)
		reg32624 := PrimCons(reg31966, reg32623)
		reg32625 := PrimCons(reg31965, reg32624)
		reg32626 := PrimCons(reg31964, reg32625)
		reg32627 := PrimCons(reg31963, reg32626)
		reg32628 := PrimCons(reg31962, reg32627)
		reg32629 := PrimCons(reg31961, reg32628)
		reg32630 := PrimCons(reg31960, reg32629)
		reg32631 := PrimCons(reg31959, reg32630)
		reg32632 := PrimCons(reg31958, reg32631)
		reg32633 := PrimCons(reg31957, reg32632)
		reg32634 := PrimCons(reg31956, reg32633)
		reg32635 := PrimCons(reg31955, reg32634)
		reg32636 := PrimCons(reg31954, reg32635)
		reg32637 := PrimCons(reg31953, reg32636)
		reg32638 := PrimCons(reg31952, reg32637)
		reg32639 := PrimCons(reg31951, reg32638)
		reg32640 := PrimCons(reg31950, reg32639)
		reg32641 := PrimCons(reg31949, reg32640)
		reg32642 := Call(__e, __defun__shen_4initialise__arity__table, reg32641)
		_ = reg32642
		reg32643 := MakeString("shen")
		reg32644 := PrimIntern(reg32643)
		reg32645 := MakeSymbol("shen.external-symbols")
		reg32646 := MakeSymbol("!")
		reg32647 := MakeSymbol("}")
		reg32648 := MakeSymbol("{")
		reg32649 := MakeSymbol("-->")
		reg32650 := MakeSymbol("<--")
		reg32651 := MakeSymbol("&&")
		reg32652 := MakeSymbol(":")
		reg32653 := MakeSymbol(";")
		reg32654 := MakeSymbol(":-")
		reg32655 := MakeSymbol(":=")
		reg32656 := MakeSymbol("_")
		reg32657 := MakeSymbol(",")
		reg32658 := MakeSymbol("*language*")
		reg32659 := MakeSymbol("*implementation*")
		reg32660 := MakeSymbol("*stinput*")
		reg32661 := MakeSymbol("*stoutput*")
		reg32662 := MakeSymbol("*sterror*")
		reg32663 := MakeSymbol("*home-directory*")
		reg32664 := MakeSymbol("*version*")
		reg32665 := MakeSymbol("*maximum-print-sequence-size*")
		reg32666 := MakeSymbol("*macros*")
		reg32667 := MakeSymbol("*os*")
		reg32668 := MakeSymbol("*release*")
		reg32669 := MakeSymbol("*property-vector*")
		reg32670 := MakeSymbol("*port*")
		reg32671 := MakeSymbol("*porters*")
		reg32672 := MakeSymbol("*hush*")
		reg32673 := MakeSymbol("@v")
		reg32674 := MakeSymbol("@p")
		reg32675 := MakeSymbol("@s")
		reg32676 := MakeSymbol("<-")
		reg32677 := MakeSymbol("->")
		reg32678 := MakeSymbol("<e>")
		reg32679 := MakeSymbol("<!>")
		reg32680 := MakeSymbol("==")
		reg32681 := MakeSymbol("=")
		reg32682 := MakeSymbol(">=")
		reg32683 := MakeSymbol(">")
		reg32684 := MakeSymbol("/.")
		reg32685 := MakeSymbol("=!")
		reg32686 := MakeSymbol("$")
		reg32687 := MakeSymbol("-")
		reg32688 := MakeSymbol("/")
		reg32689 := MakeSymbol("*")
		reg32690 := MakeSymbol("+")
		reg32691 := MakeSymbol("<=")
		reg32692 := MakeSymbol("<")
		reg32693 := MakeSymbol(">>")
		reg32694 := MakeSymbol("y-or-n?")
		reg32695 := MakeSymbol("write-to-file")
		reg32696 := MakeSymbol("write-byte")
		reg32697 := MakeSymbol("where")
		reg32698 := MakeSymbol("when")
		reg32699 := MakeSymbol("warn")
		reg32700 := MakeSymbol("version")
		reg32701 := MakeSymbol("verified")
		reg32702 := MakeSymbol("variable?")
		reg32703 := MakeSymbol("value")
		reg32704 := MakeSymbol("vector->")
		reg32705 := MakeSymbol("<-vector")
		reg32706 := MakeSymbol("vector")
		reg32707 := MakeSymbol("vector?")
		reg32708 := MakeSymbol("unspecialise")
		reg32709 := MakeSymbol("untrack")
		reg32710 := MakeSymbol("unit")
		reg32711 := MakeSymbol("shen.unix")
		reg32712 := MakeSymbol("union")
		reg32713 := MakeSymbol("unify")
		reg32714 := MakeSymbol("unify!")
		reg32715 := MakeSymbol("unput")
		reg32716 := MakeSymbol("unprofile")
		reg32717 := MakeSymbol("undefmacro")
		reg32718 := MakeSymbol("return")
		reg32719 := MakeSymbol("type")
		reg32720 := MakeSymbol("tuple?")
		reg32721 := True
		reg32722 := MakeSymbol("trap-error")
		reg32723 := MakeSymbol("track")
		reg32724 := MakeSymbol("time")
		reg32725 := MakeSymbol("thaw")
		reg32726 := MakeSymbol("tc?")
		reg32727 := MakeSymbol("tc")
		reg32728 := MakeSymbol("tl")
		reg32729 := MakeSymbol("tlstr")
		reg32730 := MakeSymbol("tlv")
		reg32731 := MakeSymbol("tail")
		reg32732 := MakeSymbol("systemf")
		reg32733 := MakeSymbol("synonyms")
		reg32734 := MakeSymbol("symbol")
		reg32735 := MakeSymbol("symbol?")
		reg32736 := MakeSymbol("string->symbol")
		reg32737 := MakeSymbol("sum")
		reg32738 := MakeSymbol("subst")
		reg32739 := MakeSymbol("string?")
		reg32740 := MakeSymbol("string->n")
		reg32741 := MakeSymbol("stream")
		reg32742 := MakeSymbol("string")
		reg32743 := MakeSymbol("stinput")
		reg32744 := MakeSymbol("sterror")
		reg32745 := MakeSymbol("stoutput")
		reg32746 := MakeSymbol("step")
		reg32747 := MakeSymbol("spy")
		reg32748 := MakeSymbol("specialise")
		reg32749 := MakeSymbol("snd")
		reg32750 := MakeSymbol("simple-error")
		reg32751 := MakeSymbol("set")
		reg32752 := MakeSymbol("save")
		reg32753 := MakeSymbol("str")
		reg32754 := MakeSymbol("run")
		reg32755 := MakeSymbol("reverse")
		reg32756 := MakeSymbol("remove")
		reg32757 := MakeSymbol("release")
		reg32758 := MakeSymbol("read")
		reg32759 := MakeSymbol("receive")
		reg32760 := MakeSymbol("read-file")
		reg32761 := MakeSymbol("read-file-as-bytelist")
		reg32762 := MakeSymbol("read-file-as-string")
		reg32763 := MakeSymbol("read-byte")
		reg32764 := MakeSymbol("read-from-string")
		reg32765 := MakeSymbol("package?")
		reg32766 := MakeSymbol("put")
		reg32767 := MakeSymbol("preclude")
		reg32768 := MakeSymbol("preclude-all-but")
		reg32769 := MakeSymbol("ps")
		reg32770 := MakeSymbol("prolog?")
		reg32771 := MakeSymbol("protect")
		reg32772 := MakeSymbol("profile-results")
		reg32773 := MakeSymbol("profile")
		reg32774 := MakeSymbol("print")
		reg32775 := MakeSymbol("pr")
		reg32776 := MakeSymbol("pos")
		reg32777 := MakeSymbol("porters")
		reg32778 := MakeSymbol("port")
		reg32779 := MakeSymbol("package")
		reg32780 := MakeSymbol("output")
		reg32781 := MakeSymbol("out")
		reg32782 := MakeSymbol("os")
		reg32783 := MakeSymbol("or")
		reg32784 := MakeSymbol("optimise")
		reg32785 := MakeSymbol("open")
		reg32786 := MakeSymbol("occurrences")
		reg32787 := MakeSymbol("occurs-check")
		reg32788 := MakeSymbol("n->string")
		reg32789 := MakeSymbol("number?")
		reg32790 := MakeSymbol("number")
		reg32791 := MakeSymbol("null")
		reg32792 := MakeSymbol("nth")
		reg32793 := MakeSymbol("not")
		reg32794 := MakeSymbol("nl")
		reg32795 := MakeSymbol("mode")
		reg32796 := MakeSymbol("macroexpand")
		reg32797 := MakeSymbol("maxinferences")
		reg32798 := MakeSymbol("mapcan")
		reg32799 := MakeSymbol("map")
		reg32800 := MakeSymbol("make-string")
		reg32801 := MakeSymbol("load")
		reg32802 := MakeSymbol("loaded")
		reg32803 := MakeSymbol("list")
		reg32804 := MakeSymbol("lineread")
		reg32805 := MakeSymbol("limit")
		reg32806 := MakeSymbol("length")
		reg32807 := MakeSymbol("let")
		reg32808 := MakeSymbol("lazy")
		reg32809 := MakeSymbol("lambda")
		reg32810 := MakeSymbol("language")
		reg32811 := MakeSymbol("kill")
		reg32812 := MakeSymbol("is")
		reg32813 := MakeSymbol("intersection")
		reg32814 := MakeSymbol("inferences")
		reg32815 := MakeSymbol("intern")
		reg32816 := MakeSymbol("integer?")
		reg32817 := MakeSymbol("input")
		reg32818 := MakeSymbol("input+")
		reg32819 := MakeSymbol("include")
		reg32820 := MakeSymbol("include-all-but")
		reg32821 := MakeSymbol("it")
		reg32822 := MakeSymbol("in")
		reg32823 := MakeSymbol("internal")
		reg32824 := MakeSymbol("implementation")
		reg32825 := MakeSymbol("if")
		reg32826 := MakeSymbol("identical")
		reg32827 := MakeSymbol("head")
		reg32828 := MakeSymbol("hd")
		reg32829 := MakeSymbol("hdv")
		reg32830 := MakeSymbol("hdstr")
		reg32831 := MakeSymbol("hash")
		reg32832 := MakeSymbol("get")
		reg32833 := MakeSymbol("get-time")
		reg32834 := MakeSymbol("gensym")
		reg32835 := MakeSymbol("function")
		reg32836 := MakeSymbol("fst")
		reg32837 := MakeSymbol("freeze")
		reg32838 := MakeSymbol("fix")
		reg32839 := MakeSymbol("file")
		reg32840 := MakeSymbol("fail")
		reg32841 := MakeSymbol("fail-if")
		reg32842 := MakeSymbol("fwhen")
		reg32843 := MakeSymbol("findall")
		reg32844 := False
		reg32845 := MakeSymbol("enable-type-theory")
		reg32846 := MakeSymbol("explode")
		reg32847 := MakeSymbol("external")
		reg32848 := MakeSymbol("exception")
		reg32849 := MakeSymbol("eval-kl")
		reg32850 := MakeSymbol("eval")
		reg32851 := MakeSymbol("error-to-string")
		reg32852 := MakeSymbol("error")
		reg32853 := MakeSymbol("empty?")
		reg32854 := MakeSymbol("element?")
		reg32855 := MakeSymbol("do")
		reg32856 := MakeSymbol("difference")
		reg32857 := MakeSymbol("destroy")
		reg32858 := MakeSymbol("defun")
		reg32859 := MakeSymbol("define")
		reg32860 := MakeSymbol("defmacro")
		reg32861 := MakeSymbol("defcc")
		reg32862 := MakeSymbol("defprolog")
		reg32863 := MakeSymbol("declare")
		reg32864 := MakeSymbol("datatype")
		reg32865 := MakeSymbol("cut")
		reg32866 := MakeSymbol("cn")
		reg32867 := MakeSymbol("cons?")
		reg32868 := MakeSymbol("cons")
		reg32869 := MakeSymbol("cond")
		reg32870 := MakeSymbol("concat")
		reg32871 := MakeSymbol("compile")
		reg32872 := MakeSymbol("cd")
		reg32873 := MakeSymbol("cases")
		reg32874 := MakeSymbol("call")
		reg32875 := MakeSymbol("close")
		reg32876 := MakeSymbol("bind")
		reg32877 := MakeSymbol("bound?")
		reg32878 := MakeSymbol("boolean?")
		reg32879 := MakeSymbol("boolean")
		reg32880 := MakeSymbol("bar!")
		reg32881 := MakeSymbol("assoc")
		reg32882 := MakeSymbol("arity")
		reg32883 := MakeSymbol("abort")
		reg32884 := MakeSymbol("append")
		reg32885 := MakeSymbol("and")
		reg32886 := MakeSymbol("adjoin")
		reg32887 := MakeSymbol("<-address")
		reg32888 := MakeSymbol("address->")
		reg32889 := MakeSymbol("absvector?")
		reg32890 := MakeSymbol("absvector")
		reg32891 := Nil
		reg32892 := PrimCons(reg32890, reg32891)
		reg32893 := PrimCons(reg32889, reg32892)
		reg32894 := PrimCons(reg32888, reg32893)
		reg32895 := PrimCons(reg32887, reg32894)
		reg32896 := PrimCons(reg32886, reg32895)
		reg32897 := PrimCons(reg32885, reg32896)
		reg32898 := PrimCons(reg32884, reg32897)
		reg32899 := PrimCons(reg32883, reg32898)
		reg32900 := PrimCons(reg32882, reg32899)
		reg32901 := PrimCons(reg32881, reg32900)
		reg32902 := PrimCons(reg32880, reg32901)
		reg32903 := PrimCons(reg32879, reg32902)
		reg32904 := PrimCons(reg32878, reg32903)
		reg32905 := PrimCons(reg32877, reg32904)
		reg32906 := PrimCons(reg32876, reg32905)
		reg32907 := PrimCons(reg32875, reg32906)
		reg32908 := PrimCons(reg32874, reg32907)
		reg32909 := PrimCons(reg32873, reg32908)
		reg32910 := PrimCons(reg32872, reg32909)
		reg32911 := PrimCons(reg32871, reg32910)
		reg32912 := PrimCons(reg32870, reg32911)
		reg32913 := PrimCons(reg32869, reg32912)
		reg32914 := PrimCons(reg32868, reg32913)
		reg32915 := PrimCons(reg32867, reg32914)
		reg32916 := PrimCons(reg32866, reg32915)
		reg32917 := PrimCons(reg32865, reg32916)
		reg32918 := PrimCons(reg32864, reg32917)
		reg32919 := PrimCons(reg32863, reg32918)
		reg32920 := PrimCons(reg32862, reg32919)
		reg32921 := PrimCons(reg32861, reg32920)
		reg32922 := PrimCons(reg32860, reg32921)
		reg32923 := PrimCons(reg32859, reg32922)
		reg32924 := PrimCons(reg32858, reg32923)
		reg32925 := PrimCons(reg32857, reg32924)
		reg32926 := PrimCons(reg32856, reg32925)
		reg32927 := PrimCons(reg32855, reg32926)
		reg32928 := PrimCons(reg32854, reg32927)
		reg32929 := PrimCons(reg32853, reg32928)
		reg32930 := PrimCons(reg32852, reg32929)
		reg32931 := PrimCons(reg32851, reg32930)
		reg32932 := PrimCons(reg32850, reg32931)
		reg32933 := PrimCons(reg32849, reg32932)
		reg32934 := PrimCons(reg32848, reg32933)
		reg32935 := PrimCons(reg32847, reg32934)
		reg32936 := PrimCons(reg32846, reg32935)
		reg32937 := PrimCons(reg32845, reg32936)
		reg32938 := PrimCons(reg32844, reg32937)
		reg32939 := PrimCons(reg32843, reg32938)
		reg32940 := PrimCons(reg32842, reg32939)
		reg32941 := PrimCons(reg32841, reg32940)
		reg32942 := PrimCons(reg32840, reg32941)
		reg32943 := PrimCons(reg32839, reg32942)
		reg32944 := PrimCons(reg32838, reg32943)
		reg32945 := PrimCons(reg32837, reg32944)
		reg32946 := PrimCons(reg32836, reg32945)
		reg32947 := PrimCons(reg32835, reg32946)
		reg32948 := PrimCons(reg32834, reg32947)
		reg32949 := PrimCons(reg32833, reg32948)
		reg32950 := PrimCons(reg32832, reg32949)
		reg32951 := PrimCons(reg32831, reg32950)
		reg32952 := PrimCons(reg32830, reg32951)
		reg32953 := PrimCons(reg32829, reg32952)
		reg32954 := PrimCons(reg32828, reg32953)
		reg32955 := PrimCons(reg32827, reg32954)
		reg32956 := PrimCons(reg32826, reg32955)
		reg32957 := PrimCons(reg32825, reg32956)
		reg32958 := PrimCons(reg32824, reg32957)
		reg32959 := PrimCons(reg32823, reg32958)
		reg32960 := PrimCons(reg32822, reg32959)
		reg32961 := PrimCons(reg32821, reg32960)
		reg32962 := PrimCons(reg32820, reg32961)
		reg32963 := PrimCons(reg32819, reg32962)
		reg32964 := PrimCons(reg32818, reg32963)
		reg32965 := PrimCons(reg32817, reg32964)
		reg32966 := PrimCons(reg32816, reg32965)
		reg32967 := PrimCons(reg32815, reg32966)
		reg32968 := PrimCons(reg32814, reg32967)
		reg32969 := PrimCons(reg32813, reg32968)
		reg32970 := PrimCons(reg32812, reg32969)
		reg32971 := PrimCons(reg32811, reg32970)
		reg32972 := PrimCons(reg32810, reg32971)
		reg32973 := PrimCons(reg32809, reg32972)
		reg32974 := PrimCons(reg32808, reg32973)
		reg32975 := PrimCons(reg32807, reg32974)
		reg32976 := PrimCons(reg32806, reg32975)
		reg32977 := PrimCons(reg32805, reg32976)
		reg32978 := PrimCons(reg32804, reg32977)
		reg32979 := PrimCons(reg32803, reg32978)
		reg32980 := PrimCons(reg32802, reg32979)
		reg32981 := PrimCons(reg32801, reg32980)
		reg32982 := PrimCons(reg32800, reg32981)
		reg32983 := PrimCons(reg32799, reg32982)
		reg32984 := PrimCons(reg32798, reg32983)
		reg32985 := PrimCons(reg32797, reg32984)
		reg32986 := PrimCons(reg32796, reg32985)
		reg32987 := PrimCons(reg32795, reg32986)
		reg32988 := PrimCons(reg32794, reg32987)
		reg32989 := PrimCons(reg32793, reg32988)
		reg32990 := PrimCons(reg32792, reg32989)
		reg32991 := PrimCons(reg32791, reg32990)
		reg32992 := PrimCons(reg32790, reg32991)
		reg32993 := PrimCons(reg32789, reg32992)
		reg32994 := PrimCons(reg32788, reg32993)
		reg32995 := PrimCons(reg32787, reg32994)
		reg32996 := PrimCons(reg32786, reg32995)
		reg32997 := PrimCons(reg32785, reg32996)
		reg32998 := PrimCons(reg32784, reg32997)
		reg32999 := PrimCons(reg32783, reg32998)
		reg33000 := PrimCons(reg32782, reg32999)
		reg33001 := PrimCons(reg32781, reg33000)
		reg33002 := PrimCons(reg32780, reg33001)
		reg33003 := PrimCons(reg32779, reg33002)
		reg33004 := PrimCons(reg32778, reg33003)
		reg33005 := PrimCons(reg32777, reg33004)
		reg33006 := PrimCons(reg32776, reg33005)
		reg33007 := PrimCons(reg32775, reg33006)
		reg33008 := PrimCons(reg32774, reg33007)
		reg33009 := PrimCons(reg32773, reg33008)
		reg33010 := PrimCons(reg32772, reg33009)
		reg33011 := PrimCons(reg32771, reg33010)
		reg33012 := PrimCons(reg32770, reg33011)
		reg33013 := PrimCons(reg32769, reg33012)
		reg33014 := PrimCons(reg32768, reg33013)
		reg33015 := PrimCons(reg32767, reg33014)
		reg33016 := PrimCons(reg32766, reg33015)
		reg33017 := PrimCons(reg32765, reg33016)
		reg33018 := PrimCons(reg32764, reg33017)
		reg33019 := PrimCons(reg32763, reg33018)
		reg33020 := PrimCons(reg32762, reg33019)
		reg33021 := PrimCons(reg32761, reg33020)
		reg33022 := PrimCons(reg32760, reg33021)
		reg33023 := PrimCons(reg32759, reg33022)
		reg33024 := PrimCons(reg32758, reg33023)
		reg33025 := PrimCons(reg32757, reg33024)
		reg33026 := PrimCons(reg32756, reg33025)
		reg33027 := PrimCons(reg32755, reg33026)
		reg33028 := PrimCons(reg32754, reg33027)
		reg33029 := PrimCons(reg32753, reg33028)
		reg33030 := PrimCons(reg32752, reg33029)
		reg33031 := PrimCons(reg32751, reg33030)
		reg33032 := PrimCons(reg32750, reg33031)
		reg33033 := PrimCons(reg32749, reg33032)
		reg33034 := PrimCons(reg32748, reg33033)
		reg33035 := PrimCons(reg32747, reg33034)
		reg33036 := PrimCons(reg32746, reg33035)
		reg33037 := PrimCons(reg32745, reg33036)
		reg33038 := PrimCons(reg32744, reg33037)
		reg33039 := PrimCons(reg32743, reg33038)
		reg33040 := PrimCons(reg32742, reg33039)
		reg33041 := PrimCons(reg32741, reg33040)
		reg33042 := PrimCons(reg32740, reg33041)
		reg33043 := PrimCons(reg32739, reg33042)
		reg33044 := PrimCons(reg32738, reg33043)
		reg33045 := PrimCons(reg32737, reg33044)
		reg33046 := PrimCons(reg32736, reg33045)
		reg33047 := PrimCons(reg32735, reg33046)
		reg33048 := PrimCons(reg32734, reg33047)
		reg33049 := PrimCons(reg32733, reg33048)
		reg33050 := PrimCons(reg32732, reg33049)
		reg33051 := PrimCons(reg32731, reg33050)
		reg33052 := PrimCons(reg32730, reg33051)
		reg33053 := PrimCons(reg32729, reg33052)
		reg33054 := PrimCons(reg32728, reg33053)
		reg33055 := PrimCons(reg32727, reg33054)
		reg33056 := PrimCons(reg32726, reg33055)
		reg33057 := PrimCons(reg32725, reg33056)
		reg33058 := PrimCons(reg32724, reg33057)
		reg33059 := PrimCons(reg32723, reg33058)
		reg33060 := PrimCons(reg32722, reg33059)
		reg33061 := PrimCons(reg32721, reg33060)
		reg33062 := PrimCons(reg32720, reg33061)
		reg33063 := PrimCons(reg32719, reg33062)
		reg33064 := PrimCons(reg32718, reg33063)
		reg33065 := PrimCons(reg32717, reg33064)
		reg33066 := PrimCons(reg32716, reg33065)
		reg33067 := PrimCons(reg32715, reg33066)
		reg33068 := PrimCons(reg32714, reg33067)
		reg33069 := PrimCons(reg32713, reg33068)
		reg33070 := PrimCons(reg32712, reg33069)
		reg33071 := PrimCons(reg32711, reg33070)
		reg33072 := PrimCons(reg32710, reg33071)
		reg33073 := PrimCons(reg32709, reg33072)
		reg33074 := PrimCons(reg32708, reg33073)
		reg33075 := PrimCons(reg32707, reg33074)
		reg33076 := PrimCons(reg32706, reg33075)
		reg33077 := PrimCons(reg32705, reg33076)
		reg33078 := PrimCons(reg32704, reg33077)
		reg33079 := PrimCons(reg32703, reg33078)
		reg33080 := PrimCons(reg32702, reg33079)
		reg33081 := PrimCons(reg32701, reg33080)
		reg33082 := PrimCons(reg32700, reg33081)
		reg33083 := PrimCons(reg32699, reg33082)
		reg33084 := PrimCons(reg32698, reg33083)
		reg33085 := PrimCons(reg32697, reg33084)
		reg33086 := PrimCons(reg32696, reg33085)
		reg33087 := PrimCons(reg32695, reg33086)
		reg33088 := PrimCons(reg32694, reg33087)
		reg33089 := PrimCons(reg32693, reg33088)
		reg33090 := PrimCons(reg32692, reg33089)
		reg33091 := PrimCons(reg32691, reg33090)
		reg33092 := PrimCons(reg32690, reg33091)
		reg33093 := PrimCons(reg32689, reg33092)
		reg33094 := PrimCons(reg32688, reg33093)
		reg33095 := PrimCons(reg32687, reg33094)
		reg33096 := PrimCons(reg32686, reg33095)
		reg33097 := PrimCons(reg32685, reg33096)
		reg33098 := PrimCons(reg32684, reg33097)
		reg33099 := PrimCons(reg32683, reg33098)
		reg33100 := PrimCons(reg32682, reg33099)
		reg33101 := PrimCons(reg32681, reg33100)
		reg33102 := PrimCons(reg32680, reg33101)
		reg33103 := PrimCons(reg32679, reg33102)
		reg33104 := PrimCons(reg32678, reg33103)
		reg33105 := PrimCons(reg32677, reg33104)
		reg33106 := PrimCons(reg32676, reg33105)
		reg33107 := PrimCons(reg32675, reg33106)
		reg33108 := PrimCons(reg32674, reg33107)
		reg33109 := PrimCons(reg32673, reg33108)
		reg33110 := PrimCons(reg32672, reg33109)
		reg33111 := PrimCons(reg32671, reg33110)
		reg33112 := PrimCons(reg32670, reg33111)
		reg33113 := PrimCons(reg32669, reg33112)
		reg33114 := PrimCons(reg32668, reg33113)
		reg33115 := PrimCons(reg32667, reg33114)
		reg33116 := PrimCons(reg32666, reg33115)
		reg33117 := PrimCons(reg32665, reg33116)
		reg33118 := PrimCons(reg32664, reg33117)
		reg33119 := PrimCons(reg32663, reg33118)
		reg33120 := PrimCons(reg32662, reg33119)
		reg33121 := PrimCons(reg32661, reg33120)
		reg33122 := PrimCons(reg32660, reg33121)
		reg33123 := PrimCons(reg32659, reg33122)
		reg33124 := PrimCons(reg32658, reg33123)
		reg33125 := PrimCons(reg32657, reg33124)
		reg33126 := PrimCons(reg32656, reg33125)
		reg33127 := PrimCons(reg32655, reg33126)
		reg33128 := PrimCons(reg32654, reg33127)
		reg33129 := PrimCons(reg32653, reg33128)
		reg33130 := PrimCons(reg32652, reg33129)
		reg33131 := PrimCons(reg32651, reg33130)
		reg33132 := PrimCons(reg32650, reg33131)
		reg33133 := PrimCons(reg32649, reg33132)
		reg33134 := PrimCons(reg32648, reg33133)
		reg33135 := PrimCons(reg32647, reg33134)
		reg33136 := PrimCons(reg32646, reg33135)
		reg33137 := MakeSymbol("*property-vector*")
		reg33138 := PrimValue(reg33137)
		reg33139 := Call(__e, __defun__put, reg32644, reg32645, reg33136, reg33138)
		_ = reg33139
		reg33140 := MakeSymbol("shen.*history*")
		reg33141 := Nil
		reg33142 := PrimSet(reg33140, reg33141)
		_ = reg33142
		reg33143 := MakeSymbol("shen.*step*")
		reg33144 := False
		reg33145 := PrimSet(reg33143, reg33144)
		_ = reg33145
		reg33146 := MakeSymbol("shen.*empty-absvector*")
		reg33147 := MakeNumber(0)
		reg33148 := PrimAbsvector(reg33147)
		reg33149 := PrimSet(reg33146, reg33148)
		__e.Return(reg33149)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.initialise-environment", value: __defun__shen_4initialise_1environment})

	__defun__shen_4initialise_1signedfuncs = MakeNative(func(__e Evaluator, __args ...Obj) {
		reg33150 := MakeSymbol("shen.*signedfuncs*")
		reg33151 := Nil
		reg33152 := PrimSet(reg33150, reg33151)
		_ = reg33152
		reg33153 := MakeSymbol("shen.*signedfuncs*")
		reg33154 := MakeSymbol("absvector?")
		reg33155 := MakeSymbol("A")
		reg33156 := MakeSymbol("-->")
		reg33157 := MakeSymbol("boolean")
		reg33158 := Nil
		reg33159 := PrimCons(reg33157, reg33158)
		reg33160 := PrimCons(reg33156, reg33159)
		reg33161 := PrimCons(reg33155, reg33160)
		reg33162 := PrimCons(reg33154, reg33161)
		reg33163 := MakeSymbol("shen.*signedfuncs*")
		reg33164 := PrimValue(reg33163)
		reg33165 := PrimCons(reg33162, reg33164)
		reg33166 := PrimSet(reg33153, reg33165)
		_ = reg33166
		reg33167 := MakeSymbol("shen.*signedfuncs*")
		reg33168 := MakeSymbol("adjoin")
		reg33169 := MakeSymbol("A")
		reg33170 := MakeSymbol("-->")
		reg33171 := MakeSymbol("list")
		reg33172 := MakeSymbol("A")
		reg33173 := Nil
		reg33174 := PrimCons(reg33172, reg33173)
		reg33175 := PrimCons(reg33171, reg33174)
		reg33176 := MakeSymbol("-->")
		reg33177 := MakeSymbol("list")
		reg33178 := MakeSymbol("A")
		reg33179 := Nil
		reg33180 := PrimCons(reg33178, reg33179)
		reg33181 := PrimCons(reg33177, reg33180)
		reg33182 := Nil
		reg33183 := PrimCons(reg33181, reg33182)
		reg33184 := PrimCons(reg33176, reg33183)
		reg33185 := PrimCons(reg33175, reg33184)
		reg33186 := Nil
		reg33187 := PrimCons(reg33185, reg33186)
		reg33188 := PrimCons(reg33170, reg33187)
		reg33189 := PrimCons(reg33169, reg33188)
		reg33190 := PrimCons(reg33168, reg33189)
		reg33191 := MakeSymbol("shen.*signedfuncs*")
		reg33192 := PrimValue(reg33191)
		reg33193 := PrimCons(reg33190, reg33192)
		reg33194 := PrimSet(reg33167, reg33193)
		_ = reg33194
		reg33195 := MakeSymbol("shen.*signedfuncs*")
		reg33196 := MakeSymbol("and")
		reg33197 := MakeSymbol("boolean")
		reg33198 := MakeSymbol("-->")
		reg33199 := MakeSymbol("boolean")
		reg33200 := MakeSymbol("-->")
		reg33201 := MakeSymbol("boolean")
		reg33202 := Nil
		reg33203 := PrimCons(reg33201, reg33202)
		reg33204 := PrimCons(reg33200, reg33203)
		reg33205 := PrimCons(reg33199, reg33204)
		reg33206 := Nil
		reg33207 := PrimCons(reg33205, reg33206)
		reg33208 := PrimCons(reg33198, reg33207)
		reg33209 := PrimCons(reg33197, reg33208)
		reg33210 := PrimCons(reg33196, reg33209)
		reg33211 := MakeSymbol("shen.*signedfuncs*")
		reg33212 := PrimValue(reg33211)
		reg33213 := PrimCons(reg33210, reg33212)
		reg33214 := PrimSet(reg33195, reg33213)
		_ = reg33214
		reg33215 := MakeSymbol("shen.*signedfuncs*")
		reg33216 := MakeSymbol("shen.app")
		reg33217 := MakeSymbol("A")
		reg33218 := MakeSymbol("-->")
		reg33219 := MakeSymbol("string")
		reg33220 := MakeSymbol("-->")
		reg33221 := MakeSymbol("symbol")
		reg33222 := MakeSymbol("-->")
		reg33223 := MakeSymbol("string")
		reg33224 := Nil
		reg33225 := PrimCons(reg33223, reg33224)
		reg33226 := PrimCons(reg33222, reg33225)
		reg33227 := PrimCons(reg33221, reg33226)
		reg33228 := Nil
		reg33229 := PrimCons(reg33227, reg33228)
		reg33230 := PrimCons(reg33220, reg33229)
		reg33231 := PrimCons(reg33219, reg33230)
		reg33232 := Nil
		reg33233 := PrimCons(reg33231, reg33232)
		reg33234 := PrimCons(reg33218, reg33233)
		reg33235 := PrimCons(reg33217, reg33234)
		reg33236 := PrimCons(reg33216, reg33235)
		reg33237 := MakeSymbol("shen.*signedfuncs*")
		reg33238 := PrimValue(reg33237)
		reg33239 := PrimCons(reg33236, reg33238)
		reg33240 := PrimSet(reg33215, reg33239)
		_ = reg33240
		reg33241 := MakeSymbol("shen.*signedfuncs*")
		reg33242 := MakeSymbol("append")
		reg33243 := MakeSymbol("list")
		reg33244 := MakeSymbol("A")
		reg33245 := Nil
		reg33246 := PrimCons(reg33244, reg33245)
		reg33247 := PrimCons(reg33243, reg33246)
		reg33248 := MakeSymbol("-->")
		reg33249 := MakeSymbol("list")
		reg33250 := MakeSymbol("A")
		reg33251 := Nil
		reg33252 := PrimCons(reg33250, reg33251)
		reg33253 := PrimCons(reg33249, reg33252)
		reg33254 := MakeSymbol("-->")
		reg33255 := MakeSymbol("list")
		reg33256 := MakeSymbol("A")
		reg33257 := Nil
		reg33258 := PrimCons(reg33256, reg33257)
		reg33259 := PrimCons(reg33255, reg33258)
		reg33260 := Nil
		reg33261 := PrimCons(reg33259, reg33260)
		reg33262 := PrimCons(reg33254, reg33261)
		reg33263 := PrimCons(reg33253, reg33262)
		reg33264 := Nil
		reg33265 := PrimCons(reg33263, reg33264)
		reg33266 := PrimCons(reg33248, reg33265)
		reg33267 := PrimCons(reg33247, reg33266)
		reg33268 := PrimCons(reg33242, reg33267)
		reg33269 := MakeSymbol("shen.*signedfuncs*")
		reg33270 := PrimValue(reg33269)
		reg33271 := PrimCons(reg33268, reg33270)
		reg33272 := PrimSet(reg33241, reg33271)
		_ = reg33272
		reg33273 := MakeSymbol("shen.*signedfuncs*")
		reg33274 := MakeSymbol("arity")
		reg33275 := MakeSymbol("A")
		reg33276 := MakeSymbol("-->")
		reg33277 := MakeSymbol("number")
		reg33278 := Nil
		reg33279 := PrimCons(reg33277, reg33278)
		reg33280 := PrimCons(reg33276, reg33279)
		reg33281 := PrimCons(reg33275, reg33280)
		reg33282 := PrimCons(reg33274, reg33281)
		reg33283 := MakeSymbol("shen.*signedfuncs*")
		reg33284 := PrimValue(reg33283)
		reg33285 := PrimCons(reg33282, reg33284)
		reg33286 := PrimSet(reg33273, reg33285)
		_ = reg33286
		reg33287 := MakeSymbol("shen.*signedfuncs*")
		reg33288 := MakeSymbol("assoc")
		reg33289 := MakeSymbol("A")
		reg33290 := MakeSymbol("-->")
		reg33291 := MakeSymbol("list")
		reg33292 := MakeSymbol("list")
		reg33293 := MakeSymbol("A")
		reg33294 := Nil
		reg33295 := PrimCons(reg33293, reg33294)
		reg33296 := PrimCons(reg33292, reg33295)
		reg33297 := Nil
		reg33298 := PrimCons(reg33296, reg33297)
		reg33299 := PrimCons(reg33291, reg33298)
		reg33300 := MakeSymbol("-->")
		reg33301 := MakeSymbol("list")
		reg33302 := MakeSymbol("A")
		reg33303 := Nil
		reg33304 := PrimCons(reg33302, reg33303)
		reg33305 := PrimCons(reg33301, reg33304)
		reg33306 := Nil
		reg33307 := PrimCons(reg33305, reg33306)
		reg33308 := PrimCons(reg33300, reg33307)
		reg33309 := PrimCons(reg33299, reg33308)
		reg33310 := Nil
		reg33311 := PrimCons(reg33309, reg33310)
		reg33312 := PrimCons(reg33290, reg33311)
		reg33313 := PrimCons(reg33289, reg33312)
		reg33314 := PrimCons(reg33288, reg33313)
		reg33315 := MakeSymbol("shen.*signedfuncs*")
		reg33316 := PrimValue(reg33315)
		reg33317 := PrimCons(reg33314, reg33316)
		reg33318 := PrimSet(reg33287, reg33317)
		_ = reg33318
		reg33319 := MakeSymbol("shen.*signedfuncs*")
		reg33320 := MakeSymbol("boolean?")
		reg33321 := MakeSymbol("A")
		reg33322 := MakeSymbol("-->")
		reg33323 := MakeSymbol("boolean")
		reg33324 := Nil
		reg33325 := PrimCons(reg33323, reg33324)
		reg33326 := PrimCons(reg33322, reg33325)
		reg33327 := PrimCons(reg33321, reg33326)
		reg33328 := PrimCons(reg33320, reg33327)
		reg33329 := MakeSymbol("shen.*signedfuncs*")
		reg33330 := PrimValue(reg33329)
		reg33331 := PrimCons(reg33328, reg33330)
		reg33332 := PrimSet(reg33319, reg33331)
		_ = reg33332
		reg33333 := MakeSymbol("shen.*signedfuncs*")
		reg33334 := MakeSymbol("bound?")
		reg33335 := MakeSymbol("symbol")
		reg33336 := MakeSymbol("-->")
		reg33337 := MakeSymbol("boolean")
		reg33338 := Nil
		reg33339 := PrimCons(reg33337, reg33338)
		reg33340 := PrimCons(reg33336, reg33339)
		reg33341 := PrimCons(reg33335, reg33340)
		reg33342 := PrimCons(reg33334, reg33341)
		reg33343 := MakeSymbol("shen.*signedfuncs*")
		reg33344 := PrimValue(reg33343)
		reg33345 := PrimCons(reg33342, reg33344)
		reg33346 := PrimSet(reg33333, reg33345)
		_ = reg33346
		reg33347 := MakeSymbol("shen.*signedfuncs*")
		reg33348 := MakeSymbol("cd")
		reg33349 := MakeSymbol("string")
		reg33350 := MakeSymbol("-->")
		reg33351 := MakeSymbol("string")
		reg33352 := Nil
		reg33353 := PrimCons(reg33351, reg33352)
		reg33354 := PrimCons(reg33350, reg33353)
		reg33355 := PrimCons(reg33349, reg33354)
		reg33356 := PrimCons(reg33348, reg33355)
		reg33357 := MakeSymbol("shen.*signedfuncs*")
		reg33358 := PrimValue(reg33357)
		reg33359 := PrimCons(reg33356, reg33358)
		reg33360 := PrimSet(reg33347, reg33359)
		_ = reg33360
		reg33361 := MakeSymbol("shen.*signedfuncs*")
		reg33362 := MakeSymbol("close")
		reg33363 := MakeSymbol("stream")
		reg33364 := MakeSymbol("A")
		reg33365 := Nil
		reg33366 := PrimCons(reg33364, reg33365)
		reg33367 := PrimCons(reg33363, reg33366)
		reg33368 := MakeSymbol("-->")
		reg33369 := MakeSymbol("list")
		reg33370 := MakeSymbol("B")
		reg33371 := Nil
		reg33372 := PrimCons(reg33370, reg33371)
		reg33373 := PrimCons(reg33369, reg33372)
		reg33374 := Nil
		reg33375 := PrimCons(reg33373, reg33374)
		reg33376 := PrimCons(reg33368, reg33375)
		reg33377 := PrimCons(reg33367, reg33376)
		reg33378 := PrimCons(reg33362, reg33377)
		reg33379 := MakeSymbol("shen.*signedfuncs*")
		reg33380 := PrimValue(reg33379)
		reg33381 := PrimCons(reg33378, reg33380)
		reg33382 := PrimSet(reg33361, reg33381)
		_ = reg33382
		reg33383 := MakeSymbol("shen.*signedfuncs*")
		reg33384 := MakeSymbol("cn")
		reg33385 := MakeSymbol("string")
		reg33386 := MakeSymbol("-->")
		reg33387 := MakeSymbol("string")
		reg33388 := MakeSymbol("-->")
		reg33389 := MakeSymbol("string")
		reg33390 := Nil
		reg33391 := PrimCons(reg33389, reg33390)
		reg33392 := PrimCons(reg33388, reg33391)
		reg33393 := PrimCons(reg33387, reg33392)
		reg33394 := Nil
		reg33395 := PrimCons(reg33393, reg33394)
		reg33396 := PrimCons(reg33386, reg33395)
		reg33397 := PrimCons(reg33385, reg33396)
		reg33398 := PrimCons(reg33384, reg33397)
		reg33399 := MakeSymbol("shen.*signedfuncs*")
		reg33400 := PrimValue(reg33399)
		reg33401 := PrimCons(reg33398, reg33400)
		reg33402 := PrimSet(reg33383, reg33401)
		_ = reg33402
		reg33403 := MakeSymbol("shen.*signedfuncs*")
		reg33404 := MakeSymbol("compile")
		reg33405 := MakeSymbol("A")
		reg33406 := MakeSymbol("shen.==>")
		reg33407 := MakeSymbol("B")
		reg33408 := Nil
		reg33409 := PrimCons(reg33407, reg33408)
		reg33410 := PrimCons(reg33406, reg33409)
		reg33411 := PrimCons(reg33405, reg33410)
		reg33412 := MakeSymbol("-->")
		reg33413 := MakeSymbol("A")
		reg33414 := MakeSymbol("-->")
		reg33415 := MakeSymbol("A")
		reg33416 := MakeSymbol("-->")
		reg33417 := MakeSymbol("B")
		reg33418 := Nil
		reg33419 := PrimCons(reg33417, reg33418)
		reg33420 := PrimCons(reg33416, reg33419)
		reg33421 := PrimCons(reg33415, reg33420)
		reg33422 := MakeSymbol("-->")
		reg33423 := MakeSymbol("B")
		reg33424 := Nil
		reg33425 := PrimCons(reg33423, reg33424)
		reg33426 := PrimCons(reg33422, reg33425)
		reg33427 := PrimCons(reg33421, reg33426)
		reg33428 := Nil
		reg33429 := PrimCons(reg33427, reg33428)
		reg33430 := PrimCons(reg33414, reg33429)
		reg33431 := PrimCons(reg33413, reg33430)
		reg33432 := Nil
		reg33433 := PrimCons(reg33431, reg33432)
		reg33434 := PrimCons(reg33412, reg33433)
		reg33435 := PrimCons(reg33411, reg33434)
		reg33436 := PrimCons(reg33404, reg33435)
		reg33437 := MakeSymbol("shen.*signedfuncs*")
		reg33438 := PrimValue(reg33437)
		reg33439 := PrimCons(reg33436, reg33438)
		reg33440 := PrimSet(reg33403, reg33439)
		_ = reg33440
		reg33441 := MakeSymbol("shen.*signedfuncs*")
		reg33442 := MakeSymbol("cons?")
		reg33443 := MakeSymbol("A")
		reg33444 := MakeSymbol("-->")
		reg33445 := MakeSymbol("boolean")
		reg33446 := Nil
		reg33447 := PrimCons(reg33445, reg33446)
		reg33448 := PrimCons(reg33444, reg33447)
		reg33449 := PrimCons(reg33443, reg33448)
		reg33450 := PrimCons(reg33442, reg33449)
		reg33451 := MakeSymbol("shen.*signedfuncs*")
		reg33452 := PrimValue(reg33451)
		reg33453 := PrimCons(reg33450, reg33452)
		reg33454 := PrimSet(reg33441, reg33453)
		_ = reg33454
		reg33455 := MakeSymbol("shen.*signedfuncs*")
		reg33456 := MakeSymbol("destroy")
		reg33457 := MakeSymbol("A")
		reg33458 := MakeSymbol("-->")
		reg33459 := MakeSymbol("B")
		reg33460 := Nil
		reg33461 := PrimCons(reg33459, reg33460)
		reg33462 := PrimCons(reg33458, reg33461)
		reg33463 := PrimCons(reg33457, reg33462)
		reg33464 := MakeSymbol("-->")
		reg33465 := MakeSymbol("symbol")
		reg33466 := Nil
		reg33467 := PrimCons(reg33465, reg33466)
		reg33468 := PrimCons(reg33464, reg33467)
		reg33469 := PrimCons(reg33463, reg33468)
		reg33470 := PrimCons(reg33456, reg33469)
		reg33471 := MakeSymbol("shen.*signedfuncs*")
		reg33472 := PrimValue(reg33471)
		reg33473 := PrimCons(reg33470, reg33472)
		reg33474 := PrimSet(reg33455, reg33473)
		_ = reg33474
		reg33475 := MakeSymbol("shen.*signedfuncs*")
		reg33476 := MakeSymbol("difference")
		reg33477 := MakeSymbol("list")
		reg33478 := MakeSymbol("A")
		reg33479 := Nil
		reg33480 := PrimCons(reg33478, reg33479)
		reg33481 := PrimCons(reg33477, reg33480)
		reg33482 := MakeSymbol("-->")
		reg33483 := MakeSymbol("list")
		reg33484 := MakeSymbol("A")
		reg33485 := Nil
		reg33486 := PrimCons(reg33484, reg33485)
		reg33487 := PrimCons(reg33483, reg33486)
		reg33488 := MakeSymbol("-->")
		reg33489 := MakeSymbol("list")
		reg33490 := MakeSymbol("A")
		reg33491 := Nil
		reg33492 := PrimCons(reg33490, reg33491)
		reg33493 := PrimCons(reg33489, reg33492)
		reg33494 := Nil
		reg33495 := PrimCons(reg33493, reg33494)
		reg33496 := PrimCons(reg33488, reg33495)
		reg33497 := PrimCons(reg33487, reg33496)
		reg33498 := Nil
		reg33499 := PrimCons(reg33497, reg33498)
		reg33500 := PrimCons(reg33482, reg33499)
		reg33501 := PrimCons(reg33481, reg33500)
		reg33502 := PrimCons(reg33476, reg33501)
		reg33503 := MakeSymbol("shen.*signedfuncs*")
		reg33504 := PrimValue(reg33503)
		reg33505 := PrimCons(reg33502, reg33504)
		reg33506 := PrimSet(reg33475, reg33505)
		_ = reg33506
		reg33507 := MakeSymbol("shen.*signedfuncs*")
		reg33508 := MakeSymbol("do")
		reg33509 := MakeSymbol("A")
		reg33510 := MakeSymbol("-->")
		reg33511 := MakeSymbol("B")
		reg33512 := MakeSymbol("-->")
		reg33513 := MakeSymbol("B")
		reg33514 := Nil
		reg33515 := PrimCons(reg33513, reg33514)
		reg33516 := PrimCons(reg33512, reg33515)
		reg33517 := PrimCons(reg33511, reg33516)
		reg33518 := Nil
		reg33519 := PrimCons(reg33517, reg33518)
		reg33520 := PrimCons(reg33510, reg33519)
		reg33521 := PrimCons(reg33509, reg33520)
		reg33522 := PrimCons(reg33508, reg33521)
		reg33523 := MakeSymbol("shen.*signedfuncs*")
		reg33524 := PrimValue(reg33523)
		reg33525 := PrimCons(reg33522, reg33524)
		reg33526 := PrimSet(reg33507, reg33525)
		_ = reg33526
		reg33527 := MakeSymbol("shen.*signedfuncs*")
		reg33528 := MakeSymbol("<e>")
		reg33529 := MakeSymbol("list")
		reg33530 := MakeSymbol("A")
		reg33531 := Nil
		reg33532 := PrimCons(reg33530, reg33531)
		reg33533 := PrimCons(reg33529, reg33532)
		reg33534 := MakeSymbol("shen.==>")
		reg33535 := MakeSymbol("list")
		reg33536 := MakeSymbol("B")
		reg33537 := Nil
		reg33538 := PrimCons(reg33536, reg33537)
		reg33539 := PrimCons(reg33535, reg33538)
		reg33540 := Nil
		reg33541 := PrimCons(reg33539, reg33540)
		reg33542 := PrimCons(reg33534, reg33541)
		reg33543 := PrimCons(reg33533, reg33542)
		reg33544 := PrimCons(reg33528, reg33543)
		reg33545 := MakeSymbol("shen.*signedfuncs*")
		reg33546 := PrimValue(reg33545)
		reg33547 := PrimCons(reg33544, reg33546)
		reg33548 := PrimSet(reg33527, reg33547)
		_ = reg33548
		reg33549 := MakeSymbol("shen.*signedfuncs*")
		reg33550 := MakeSymbol("<!>")
		reg33551 := MakeSymbol("list")
		reg33552 := MakeSymbol("A")
		reg33553 := Nil
		reg33554 := PrimCons(reg33552, reg33553)
		reg33555 := PrimCons(reg33551, reg33554)
		reg33556 := MakeSymbol("shen.==>")
		reg33557 := MakeSymbol("list")
		reg33558 := MakeSymbol("A")
		reg33559 := Nil
		reg33560 := PrimCons(reg33558, reg33559)
		reg33561 := PrimCons(reg33557, reg33560)
		reg33562 := Nil
		reg33563 := PrimCons(reg33561, reg33562)
		reg33564 := PrimCons(reg33556, reg33563)
		reg33565 := PrimCons(reg33555, reg33564)
		reg33566 := PrimCons(reg33550, reg33565)
		reg33567 := MakeSymbol("shen.*signedfuncs*")
		reg33568 := PrimValue(reg33567)
		reg33569 := PrimCons(reg33566, reg33568)
		reg33570 := PrimSet(reg33549, reg33569)
		_ = reg33570
		reg33571 := MakeSymbol("shen.*signedfuncs*")
		reg33572 := MakeSymbol("element?")
		reg33573 := MakeSymbol("A")
		reg33574 := MakeSymbol("-->")
		reg33575 := MakeSymbol("list")
		reg33576 := MakeSymbol("A")
		reg33577 := Nil
		reg33578 := PrimCons(reg33576, reg33577)
		reg33579 := PrimCons(reg33575, reg33578)
		reg33580 := MakeSymbol("-->")
		reg33581 := MakeSymbol("boolean")
		reg33582 := Nil
		reg33583 := PrimCons(reg33581, reg33582)
		reg33584 := PrimCons(reg33580, reg33583)
		reg33585 := PrimCons(reg33579, reg33584)
		reg33586 := Nil
		reg33587 := PrimCons(reg33585, reg33586)
		reg33588 := PrimCons(reg33574, reg33587)
		reg33589 := PrimCons(reg33573, reg33588)
		reg33590 := PrimCons(reg33572, reg33589)
		reg33591 := MakeSymbol("shen.*signedfuncs*")
		reg33592 := PrimValue(reg33591)
		reg33593 := PrimCons(reg33590, reg33592)
		reg33594 := PrimSet(reg33571, reg33593)
		_ = reg33594
		reg33595 := MakeSymbol("shen.*signedfuncs*")
		reg33596 := MakeSymbol("empty?")
		reg33597 := MakeSymbol("A")
		reg33598 := MakeSymbol("-->")
		reg33599 := MakeSymbol("boolean")
		reg33600 := Nil
		reg33601 := PrimCons(reg33599, reg33600)
		reg33602 := PrimCons(reg33598, reg33601)
		reg33603 := PrimCons(reg33597, reg33602)
		reg33604 := PrimCons(reg33596, reg33603)
		reg33605 := MakeSymbol("shen.*signedfuncs*")
		reg33606 := PrimValue(reg33605)
		reg33607 := PrimCons(reg33604, reg33606)
		reg33608 := PrimSet(reg33595, reg33607)
		_ = reg33608
		reg33609 := MakeSymbol("shen.*signedfuncs*")
		reg33610 := MakeSymbol("enable-type-theory")
		reg33611 := MakeSymbol("symbol")
		reg33612 := MakeSymbol("-->")
		reg33613 := MakeSymbol("boolean")
		reg33614 := Nil
		reg33615 := PrimCons(reg33613, reg33614)
		reg33616 := PrimCons(reg33612, reg33615)
		reg33617 := PrimCons(reg33611, reg33616)
		reg33618 := PrimCons(reg33610, reg33617)
		reg33619 := MakeSymbol("shen.*signedfuncs*")
		reg33620 := PrimValue(reg33619)
		reg33621 := PrimCons(reg33618, reg33620)
		reg33622 := PrimSet(reg33609, reg33621)
		_ = reg33622
		reg33623 := MakeSymbol("shen.*signedfuncs*")
		reg33624 := MakeSymbol("external")
		reg33625 := MakeSymbol("symbol")
		reg33626 := MakeSymbol("-->")
		reg33627 := MakeSymbol("list")
		reg33628 := MakeSymbol("symbol")
		reg33629 := Nil
		reg33630 := PrimCons(reg33628, reg33629)
		reg33631 := PrimCons(reg33627, reg33630)
		reg33632 := Nil
		reg33633 := PrimCons(reg33631, reg33632)
		reg33634 := PrimCons(reg33626, reg33633)
		reg33635 := PrimCons(reg33625, reg33634)
		reg33636 := PrimCons(reg33624, reg33635)
		reg33637 := MakeSymbol("shen.*signedfuncs*")
		reg33638 := PrimValue(reg33637)
		reg33639 := PrimCons(reg33636, reg33638)
		reg33640 := PrimSet(reg33623, reg33639)
		_ = reg33640
		reg33641 := MakeSymbol("shen.*signedfuncs*")
		reg33642 := MakeSymbol("error-to-string")
		reg33643 := MakeSymbol("exception")
		reg33644 := MakeSymbol("-->")
		reg33645 := MakeSymbol("string")
		reg33646 := Nil
		reg33647 := PrimCons(reg33645, reg33646)
		reg33648 := PrimCons(reg33644, reg33647)
		reg33649 := PrimCons(reg33643, reg33648)
		reg33650 := PrimCons(reg33642, reg33649)
		reg33651 := MakeSymbol("shen.*signedfuncs*")
		reg33652 := PrimValue(reg33651)
		reg33653 := PrimCons(reg33650, reg33652)
		reg33654 := PrimSet(reg33641, reg33653)
		_ = reg33654
		reg33655 := MakeSymbol("shen.*signedfuncs*")
		reg33656 := MakeSymbol("explode")
		reg33657 := MakeSymbol("A")
		reg33658 := MakeSymbol("-->")
		reg33659 := MakeSymbol("list")
		reg33660 := MakeSymbol("string")
		reg33661 := Nil
		reg33662 := PrimCons(reg33660, reg33661)
		reg33663 := PrimCons(reg33659, reg33662)
		reg33664 := Nil
		reg33665 := PrimCons(reg33663, reg33664)
		reg33666 := PrimCons(reg33658, reg33665)
		reg33667 := PrimCons(reg33657, reg33666)
		reg33668 := PrimCons(reg33656, reg33667)
		reg33669 := MakeSymbol("shen.*signedfuncs*")
		reg33670 := PrimValue(reg33669)
		reg33671 := PrimCons(reg33668, reg33670)
		reg33672 := PrimSet(reg33655, reg33671)
		_ = reg33672
		reg33673 := MakeSymbol("shen.*signedfuncs*")
		reg33674 := MakeSymbol("fail")
		reg33675 := MakeSymbol("-->")
		reg33676 := MakeSymbol("symbol")
		reg33677 := Nil
		reg33678 := PrimCons(reg33676, reg33677)
		reg33679 := PrimCons(reg33675, reg33678)
		reg33680 := PrimCons(reg33674, reg33679)
		reg33681 := MakeSymbol("shen.*signedfuncs*")
		reg33682 := PrimValue(reg33681)
		reg33683 := PrimCons(reg33680, reg33682)
		reg33684 := PrimSet(reg33673, reg33683)
		_ = reg33684
		reg33685 := MakeSymbol("shen.*signedfuncs*")
		reg33686 := MakeSymbol("fail-if")
		reg33687 := MakeSymbol("symbol")
		reg33688 := MakeSymbol("-->")
		reg33689 := MakeSymbol("boolean")
		reg33690 := Nil
		reg33691 := PrimCons(reg33689, reg33690)
		reg33692 := PrimCons(reg33688, reg33691)
		reg33693 := PrimCons(reg33687, reg33692)
		reg33694 := MakeSymbol("-->")
		reg33695 := MakeSymbol("symbol")
		reg33696 := MakeSymbol("-->")
		reg33697 := MakeSymbol("symbol")
		reg33698 := Nil
		reg33699 := PrimCons(reg33697, reg33698)
		reg33700 := PrimCons(reg33696, reg33699)
		reg33701 := PrimCons(reg33695, reg33700)
		reg33702 := Nil
		reg33703 := PrimCons(reg33701, reg33702)
		reg33704 := PrimCons(reg33694, reg33703)
		reg33705 := PrimCons(reg33693, reg33704)
		reg33706 := PrimCons(reg33686, reg33705)
		reg33707 := MakeSymbol("shen.*signedfuncs*")
		reg33708 := PrimValue(reg33707)
		reg33709 := PrimCons(reg33706, reg33708)
		reg33710 := PrimSet(reg33685, reg33709)
		_ = reg33710
		reg33711 := MakeSymbol("shen.*signedfuncs*")
		reg33712 := MakeSymbol("fix")
		reg33713 := MakeSymbol("A")
		reg33714 := MakeSymbol("-->")
		reg33715 := MakeSymbol("A")
		reg33716 := Nil
		reg33717 := PrimCons(reg33715, reg33716)
		reg33718 := PrimCons(reg33714, reg33717)
		reg33719 := PrimCons(reg33713, reg33718)
		reg33720 := MakeSymbol("-->")
		reg33721 := MakeSymbol("A")
		reg33722 := MakeSymbol("-->")
		reg33723 := MakeSymbol("A")
		reg33724 := Nil
		reg33725 := PrimCons(reg33723, reg33724)
		reg33726 := PrimCons(reg33722, reg33725)
		reg33727 := PrimCons(reg33721, reg33726)
		reg33728 := Nil
		reg33729 := PrimCons(reg33727, reg33728)
		reg33730 := PrimCons(reg33720, reg33729)
		reg33731 := PrimCons(reg33719, reg33730)
		reg33732 := PrimCons(reg33712, reg33731)
		reg33733 := MakeSymbol("shen.*signedfuncs*")
		reg33734 := PrimValue(reg33733)
		reg33735 := PrimCons(reg33732, reg33734)
		reg33736 := PrimSet(reg33711, reg33735)
		_ = reg33736
		reg33737 := MakeSymbol("shen.*signedfuncs*")
		reg33738 := MakeSymbol("freeze")
		reg33739 := MakeSymbol("A")
		reg33740 := MakeSymbol("-->")
		reg33741 := MakeSymbol("lazy")
		reg33742 := MakeSymbol("A")
		reg33743 := Nil
		reg33744 := PrimCons(reg33742, reg33743)
		reg33745 := PrimCons(reg33741, reg33744)
		reg33746 := Nil
		reg33747 := PrimCons(reg33745, reg33746)
		reg33748 := PrimCons(reg33740, reg33747)
		reg33749 := PrimCons(reg33739, reg33748)
		reg33750 := PrimCons(reg33738, reg33749)
		reg33751 := MakeSymbol("shen.*signedfuncs*")
		reg33752 := PrimValue(reg33751)
		reg33753 := PrimCons(reg33750, reg33752)
		reg33754 := PrimSet(reg33737, reg33753)
		_ = reg33754
		reg33755 := MakeSymbol("shen.*signedfuncs*")
		reg33756 := MakeSymbol("fst")
		reg33757 := MakeSymbol("A")
		reg33758 := MakeSymbol("*")
		reg33759 := MakeSymbol("B")
		reg33760 := Nil
		reg33761 := PrimCons(reg33759, reg33760)
		reg33762 := PrimCons(reg33758, reg33761)
		reg33763 := PrimCons(reg33757, reg33762)
		reg33764 := MakeSymbol("-->")
		reg33765 := MakeSymbol("A")
		reg33766 := Nil
		reg33767 := PrimCons(reg33765, reg33766)
		reg33768 := PrimCons(reg33764, reg33767)
		reg33769 := PrimCons(reg33763, reg33768)
		reg33770 := PrimCons(reg33756, reg33769)
		reg33771 := MakeSymbol("shen.*signedfuncs*")
		reg33772 := PrimValue(reg33771)
		reg33773 := PrimCons(reg33770, reg33772)
		reg33774 := PrimSet(reg33755, reg33773)
		_ = reg33774
		reg33775 := MakeSymbol("shen.*signedfuncs*")
		reg33776 := MakeSymbol("function")
		reg33777 := MakeSymbol("A")
		reg33778 := MakeSymbol("-->")
		reg33779 := MakeSymbol("B")
		reg33780 := Nil
		reg33781 := PrimCons(reg33779, reg33780)
		reg33782 := PrimCons(reg33778, reg33781)
		reg33783 := PrimCons(reg33777, reg33782)
		reg33784 := MakeSymbol("-->")
		reg33785 := MakeSymbol("A")
		reg33786 := MakeSymbol("-->")
		reg33787 := MakeSymbol("B")
		reg33788 := Nil
		reg33789 := PrimCons(reg33787, reg33788)
		reg33790 := PrimCons(reg33786, reg33789)
		reg33791 := PrimCons(reg33785, reg33790)
		reg33792 := Nil
		reg33793 := PrimCons(reg33791, reg33792)
		reg33794 := PrimCons(reg33784, reg33793)
		reg33795 := PrimCons(reg33783, reg33794)
		reg33796 := PrimCons(reg33776, reg33795)
		reg33797 := MakeSymbol("shen.*signedfuncs*")
		reg33798 := PrimValue(reg33797)
		reg33799 := PrimCons(reg33796, reg33798)
		reg33800 := PrimSet(reg33775, reg33799)
		_ = reg33800
		reg33801 := MakeSymbol("shen.*signedfuncs*")
		reg33802 := MakeSymbol("gensym")
		reg33803 := MakeSymbol("symbol")
		reg33804 := MakeSymbol("-->")
		reg33805 := MakeSymbol("symbol")
		reg33806 := Nil
		reg33807 := PrimCons(reg33805, reg33806)
		reg33808 := PrimCons(reg33804, reg33807)
		reg33809 := PrimCons(reg33803, reg33808)
		reg33810 := PrimCons(reg33802, reg33809)
		reg33811 := MakeSymbol("shen.*signedfuncs*")
		reg33812 := PrimValue(reg33811)
		reg33813 := PrimCons(reg33810, reg33812)
		reg33814 := PrimSet(reg33801, reg33813)
		_ = reg33814
		reg33815 := MakeSymbol("shen.*signedfuncs*")
		reg33816 := MakeSymbol("<-vector")
		reg33817 := MakeSymbol("vector")
		reg33818 := MakeSymbol("A")
		reg33819 := Nil
		reg33820 := PrimCons(reg33818, reg33819)
		reg33821 := PrimCons(reg33817, reg33820)
		reg33822 := MakeSymbol("-->")
		reg33823 := MakeSymbol("number")
		reg33824 := MakeSymbol("-->")
		reg33825 := MakeSymbol("A")
		reg33826 := Nil
		reg33827 := PrimCons(reg33825, reg33826)
		reg33828 := PrimCons(reg33824, reg33827)
		reg33829 := PrimCons(reg33823, reg33828)
		reg33830 := Nil
		reg33831 := PrimCons(reg33829, reg33830)
		reg33832 := PrimCons(reg33822, reg33831)
		reg33833 := PrimCons(reg33821, reg33832)
		reg33834 := PrimCons(reg33816, reg33833)
		reg33835 := MakeSymbol("shen.*signedfuncs*")
		reg33836 := PrimValue(reg33835)
		reg33837 := PrimCons(reg33834, reg33836)
		reg33838 := PrimSet(reg33815, reg33837)
		_ = reg33838
		reg33839 := MakeSymbol("shen.*signedfuncs*")
		reg33840 := MakeSymbol("vector->")
		reg33841 := MakeSymbol("vector")
		reg33842 := MakeSymbol("A")
		reg33843 := Nil
		reg33844 := PrimCons(reg33842, reg33843)
		reg33845 := PrimCons(reg33841, reg33844)
		reg33846 := MakeSymbol("-->")
		reg33847 := MakeSymbol("number")
		reg33848 := MakeSymbol("-->")
		reg33849 := MakeSymbol("A")
		reg33850 := MakeSymbol("-->")
		reg33851 := MakeSymbol("vector")
		reg33852 := MakeSymbol("A")
		reg33853 := Nil
		reg33854 := PrimCons(reg33852, reg33853)
		reg33855 := PrimCons(reg33851, reg33854)
		reg33856 := Nil
		reg33857 := PrimCons(reg33855, reg33856)
		reg33858 := PrimCons(reg33850, reg33857)
		reg33859 := PrimCons(reg33849, reg33858)
		reg33860 := Nil
		reg33861 := PrimCons(reg33859, reg33860)
		reg33862 := PrimCons(reg33848, reg33861)
		reg33863 := PrimCons(reg33847, reg33862)
		reg33864 := Nil
		reg33865 := PrimCons(reg33863, reg33864)
		reg33866 := PrimCons(reg33846, reg33865)
		reg33867 := PrimCons(reg33845, reg33866)
		reg33868 := PrimCons(reg33840, reg33867)
		reg33869 := MakeSymbol("shen.*signedfuncs*")
		reg33870 := PrimValue(reg33869)
		reg33871 := PrimCons(reg33868, reg33870)
		reg33872 := PrimSet(reg33839, reg33871)
		_ = reg33872
		reg33873 := MakeSymbol("shen.*signedfuncs*")
		reg33874 := MakeSymbol("vector")
		reg33875 := MakeSymbol("number")
		reg33876 := MakeSymbol("-->")
		reg33877 := MakeSymbol("vector")
		reg33878 := MakeSymbol("A")
		reg33879 := Nil
		reg33880 := PrimCons(reg33878, reg33879)
		reg33881 := PrimCons(reg33877, reg33880)
		reg33882 := Nil
		reg33883 := PrimCons(reg33881, reg33882)
		reg33884 := PrimCons(reg33876, reg33883)
		reg33885 := PrimCons(reg33875, reg33884)
		reg33886 := PrimCons(reg33874, reg33885)
		reg33887 := MakeSymbol("shen.*signedfuncs*")
		reg33888 := PrimValue(reg33887)
		reg33889 := PrimCons(reg33886, reg33888)
		reg33890 := PrimSet(reg33873, reg33889)
		_ = reg33890
		reg33891 := MakeSymbol("shen.*signedfuncs*")
		reg33892 := MakeSymbol("get-time")
		reg33893 := MakeSymbol("symbol")
		reg33894 := MakeSymbol("-->")
		reg33895 := MakeSymbol("number")
		reg33896 := Nil
		reg33897 := PrimCons(reg33895, reg33896)
		reg33898 := PrimCons(reg33894, reg33897)
		reg33899 := PrimCons(reg33893, reg33898)
		reg33900 := PrimCons(reg33892, reg33899)
		reg33901 := MakeSymbol("shen.*signedfuncs*")
		reg33902 := PrimValue(reg33901)
		reg33903 := PrimCons(reg33900, reg33902)
		reg33904 := PrimSet(reg33891, reg33903)
		_ = reg33904
		reg33905 := MakeSymbol("shen.*signedfuncs*")
		reg33906 := MakeSymbol("hash")
		reg33907 := MakeSymbol("A")
		reg33908 := MakeSymbol("-->")
		reg33909 := MakeSymbol("number")
		reg33910 := MakeSymbol("-->")
		reg33911 := MakeSymbol("number")
		reg33912 := Nil
		reg33913 := PrimCons(reg33911, reg33912)
		reg33914 := PrimCons(reg33910, reg33913)
		reg33915 := PrimCons(reg33909, reg33914)
		reg33916 := Nil
		reg33917 := PrimCons(reg33915, reg33916)
		reg33918 := PrimCons(reg33908, reg33917)
		reg33919 := PrimCons(reg33907, reg33918)
		reg33920 := PrimCons(reg33906, reg33919)
		reg33921 := MakeSymbol("shen.*signedfuncs*")
		reg33922 := PrimValue(reg33921)
		reg33923 := PrimCons(reg33920, reg33922)
		reg33924 := PrimSet(reg33905, reg33923)
		_ = reg33924
		reg33925 := MakeSymbol("shen.*signedfuncs*")
		reg33926 := MakeSymbol("head")
		reg33927 := MakeSymbol("list")
		reg33928 := MakeSymbol("A")
		reg33929 := Nil
		reg33930 := PrimCons(reg33928, reg33929)
		reg33931 := PrimCons(reg33927, reg33930)
		reg33932 := MakeSymbol("-->")
		reg33933 := MakeSymbol("A")
		reg33934 := Nil
		reg33935 := PrimCons(reg33933, reg33934)
		reg33936 := PrimCons(reg33932, reg33935)
		reg33937 := PrimCons(reg33931, reg33936)
		reg33938 := PrimCons(reg33926, reg33937)
		reg33939 := MakeSymbol("shen.*signedfuncs*")
		reg33940 := PrimValue(reg33939)
		reg33941 := PrimCons(reg33938, reg33940)
		reg33942 := PrimSet(reg33925, reg33941)
		_ = reg33942
		reg33943 := MakeSymbol("shen.*signedfuncs*")
		reg33944 := MakeSymbol("hdv")
		reg33945 := MakeSymbol("vector")
		reg33946 := MakeSymbol("A")
		reg33947 := Nil
		reg33948 := PrimCons(reg33946, reg33947)
		reg33949 := PrimCons(reg33945, reg33948)
		reg33950 := MakeSymbol("-->")
		reg33951 := MakeSymbol("A")
		reg33952 := Nil
		reg33953 := PrimCons(reg33951, reg33952)
		reg33954 := PrimCons(reg33950, reg33953)
		reg33955 := PrimCons(reg33949, reg33954)
		reg33956 := PrimCons(reg33944, reg33955)
		reg33957 := MakeSymbol("shen.*signedfuncs*")
		reg33958 := PrimValue(reg33957)
		reg33959 := PrimCons(reg33956, reg33958)
		reg33960 := PrimSet(reg33943, reg33959)
		_ = reg33960
		reg33961 := MakeSymbol("shen.*signedfuncs*")
		reg33962 := MakeSymbol("hdstr")
		reg33963 := MakeSymbol("string")
		reg33964 := MakeSymbol("-->")
		reg33965 := MakeSymbol("string")
		reg33966 := Nil
		reg33967 := PrimCons(reg33965, reg33966)
		reg33968 := PrimCons(reg33964, reg33967)
		reg33969 := PrimCons(reg33963, reg33968)
		reg33970 := PrimCons(reg33962, reg33969)
		reg33971 := MakeSymbol("shen.*signedfuncs*")
		reg33972 := PrimValue(reg33971)
		reg33973 := PrimCons(reg33970, reg33972)
		reg33974 := PrimSet(reg33961, reg33973)
		_ = reg33974
		reg33975 := MakeSymbol("shen.*signedfuncs*")
		reg33976 := MakeSymbol("if")
		reg33977 := MakeSymbol("boolean")
		reg33978 := MakeSymbol("-->")
		reg33979 := MakeSymbol("A")
		reg33980 := MakeSymbol("-->")
		reg33981 := MakeSymbol("A")
		reg33982 := MakeSymbol("-->")
		reg33983 := MakeSymbol("A")
		reg33984 := Nil
		reg33985 := PrimCons(reg33983, reg33984)
		reg33986 := PrimCons(reg33982, reg33985)
		reg33987 := PrimCons(reg33981, reg33986)
		reg33988 := Nil
		reg33989 := PrimCons(reg33987, reg33988)
		reg33990 := PrimCons(reg33980, reg33989)
		reg33991 := PrimCons(reg33979, reg33990)
		reg33992 := Nil
		reg33993 := PrimCons(reg33991, reg33992)
		reg33994 := PrimCons(reg33978, reg33993)
		reg33995 := PrimCons(reg33977, reg33994)
		reg33996 := PrimCons(reg33976, reg33995)
		reg33997 := MakeSymbol("shen.*signedfuncs*")
		reg33998 := PrimValue(reg33997)
		reg33999 := PrimCons(reg33996, reg33998)
		reg34000 := PrimSet(reg33975, reg33999)
		_ = reg34000
		reg34001 := MakeSymbol("shen.*signedfuncs*")
		reg34002 := MakeSymbol("it")
		reg34003 := MakeSymbol("-->")
		reg34004 := MakeSymbol("string")
		reg34005 := Nil
		reg34006 := PrimCons(reg34004, reg34005)
		reg34007 := PrimCons(reg34003, reg34006)
		reg34008 := PrimCons(reg34002, reg34007)
		reg34009 := MakeSymbol("shen.*signedfuncs*")
		reg34010 := PrimValue(reg34009)
		reg34011 := PrimCons(reg34008, reg34010)
		reg34012 := PrimSet(reg34001, reg34011)
		_ = reg34012
		reg34013 := MakeSymbol("shen.*signedfuncs*")
		reg34014 := MakeSymbol("implementation")
		reg34015 := MakeSymbol("-->")
		reg34016 := MakeSymbol("string")
		reg34017 := Nil
		reg34018 := PrimCons(reg34016, reg34017)
		reg34019 := PrimCons(reg34015, reg34018)
		reg34020 := PrimCons(reg34014, reg34019)
		reg34021 := MakeSymbol("shen.*signedfuncs*")
		reg34022 := PrimValue(reg34021)
		reg34023 := PrimCons(reg34020, reg34022)
		reg34024 := PrimSet(reg34013, reg34023)
		_ = reg34024
		reg34025 := MakeSymbol("shen.*signedfuncs*")
		reg34026 := MakeSymbol("include")
		reg34027 := MakeSymbol("list")
		reg34028 := MakeSymbol("symbol")
		reg34029 := Nil
		reg34030 := PrimCons(reg34028, reg34029)
		reg34031 := PrimCons(reg34027, reg34030)
		reg34032 := MakeSymbol("-->")
		reg34033 := MakeSymbol("list")
		reg34034 := MakeSymbol("symbol")
		reg34035 := Nil
		reg34036 := PrimCons(reg34034, reg34035)
		reg34037 := PrimCons(reg34033, reg34036)
		reg34038 := Nil
		reg34039 := PrimCons(reg34037, reg34038)
		reg34040 := PrimCons(reg34032, reg34039)
		reg34041 := PrimCons(reg34031, reg34040)
		reg34042 := PrimCons(reg34026, reg34041)
		reg34043 := MakeSymbol("shen.*signedfuncs*")
		reg34044 := PrimValue(reg34043)
		reg34045 := PrimCons(reg34042, reg34044)
		reg34046 := PrimSet(reg34025, reg34045)
		_ = reg34046
		reg34047 := MakeSymbol("shen.*signedfuncs*")
		reg34048 := MakeSymbol("include-all-but")
		reg34049 := MakeSymbol("list")
		reg34050 := MakeSymbol("symbol")
		reg34051 := Nil
		reg34052 := PrimCons(reg34050, reg34051)
		reg34053 := PrimCons(reg34049, reg34052)
		reg34054 := MakeSymbol("-->")
		reg34055 := MakeSymbol("list")
		reg34056 := MakeSymbol("symbol")
		reg34057 := Nil
		reg34058 := PrimCons(reg34056, reg34057)
		reg34059 := PrimCons(reg34055, reg34058)
		reg34060 := Nil
		reg34061 := PrimCons(reg34059, reg34060)
		reg34062 := PrimCons(reg34054, reg34061)
		reg34063 := PrimCons(reg34053, reg34062)
		reg34064 := PrimCons(reg34048, reg34063)
		reg34065 := MakeSymbol("shen.*signedfuncs*")
		reg34066 := PrimValue(reg34065)
		reg34067 := PrimCons(reg34064, reg34066)
		reg34068 := PrimSet(reg34047, reg34067)
		_ = reg34068
		reg34069 := MakeSymbol("shen.*signedfuncs*")
		reg34070 := MakeSymbol("inferences")
		reg34071 := MakeSymbol("-->")
		reg34072 := MakeSymbol("number")
		reg34073 := Nil
		reg34074 := PrimCons(reg34072, reg34073)
		reg34075 := PrimCons(reg34071, reg34074)
		reg34076 := PrimCons(reg34070, reg34075)
		reg34077 := MakeSymbol("shen.*signedfuncs*")
		reg34078 := PrimValue(reg34077)
		reg34079 := PrimCons(reg34076, reg34078)
		reg34080 := PrimSet(reg34069, reg34079)
		_ = reg34080
		reg34081 := MakeSymbol("shen.*signedfuncs*")
		reg34082 := MakeSymbol("shen.insert")
		reg34083 := MakeSymbol("A")
		reg34084 := MakeSymbol("-->")
		reg34085 := MakeSymbol("string")
		reg34086 := MakeSymbol("-->")
		reg34087 := MakeSymbol("string")
		reg34088 := Nil
		reg34089 := PrimCons(reg34087, reg34088)
		reg34090 := PrimCons(reg34086, reg34089)
		reg34091 := PrimCons(reg34085, reg34090)
		reg34092 := Nil
		reg34093 := PrimCons(reg34091, reg34092)
		reg34094 := PrimCons(reg34084, reg34093)
		reg34095 := PrimCons(reg34083, reg34094)
		reg34096 := PrimCons(reg34082, reg34095)
		reg34097 := MakeSymbol("shen.*signedfuncs*")
		reg34098 := PrimValue(reg34097)
		reg34099 := PrimCons(reg34096, reg34098)
		reg34100 := PrimSet(reg34081, reg34099)
		_ = reg34100
		reg34101 := MakeSymbol("shen.*signedfuncs*")
		reg34102 := MakeSymbol("integer?")
		reg34103 := MakeSymbol("A")
		reg34104 := MakeSymbol("-->")
		reg34105 := MakeSymbol("boolean")
		reg34106 := Nil
		reg34107 := PrimCons(reg34105, reg34106)
		reg34108 := PrimCons(reg34104, reg34107)
		reg34109 := PrimCons(reg34103, reg34108)
		reg34110 := PrimCons(reg34102, reg34109)
		reg34111 := MakeSymbol("shen.*signedfuncs*")
		reg34112 := PrimValue(reg34111)
		reg34113 := PrimCons(reg34110, reg34112)
		reg34114 := PrimSet(reg34101, reg34113)
		_ = reg34114
		reg34115 := MakeSymbol("shen.*signedfuncs*")
		reg34116 := MakeSymbol("internal")
		reg34117 := MakeSymbol("symbol")
		reg34118 := MakeSymbol("-->")
		reg34119 := MakeSymbol("list")
		reg34120 := MakeSymbol("symbol")
		reg34121 := Nil
		reg34122 := PrimCons(reg34120, reg34121)
		reg34123 := PrimCons(reg34119, reg34122)
		reg34124 := Nil
		reg34125 := PrimCons(reg34123, reg34124)
		reg34126 := PrimCons(reg34118, reg34125)
		reg34127 := PrimCons(reg34117, reg34126)
		reg34128 := PrimCons(reg34116, reg34127)
		reg34129 := MakeSymbol("shen.*signedfuncs*")
		reg34130 := PrimValue(reg34129)
		reg34131 := PrimCons(reg34128, reg34130)
		reg34132 := PrimSet(reg34115, reg34131)
		_ = reg34132
		reg34133 := MakeSymbol("shen.*signedfuncs*")
		reg34134 := MakeSymbol("intersection")
		reg34135 := MakeSymbol("list")
		reg34136 := MakeSymbol("A")
		reg34137 := Nil
		reg34138 := PrimCons(reg34136, reg34137)
		reg34139 := PrimCons(reg34135, reg34138)
		reg34140 := MakeSymbol("-->")
		reg34141 := MakeSymbol("list")
		reg34142 := MakeSymbol("A")
		reg34143 := Nil
		reg34144 := PrimCons(reg34142, reg34143)
		reg34145 := PrimCons(reg34141, reg34144)
		reg34146 := MakeSymbol("-->")
		reg34147 := MakeSymbol("list")
		reg34148 := MakeSymbol("A")
		reg34149 := Nil
		reg34150 := PrimCons(reg34148, reg34149)
		reg34151 := PrimCons(reg34147, reg34150)
		reg34152 := Nil
		reg34153 := PrimCons(reg34151, reg34152)
		reg34154 := PrimCons(reg34146, reg34153)
		reg34155 := PrimCons(reg34145, reg34154)
		reg34156 := Nil
		reg34157 := PrimCons(reg34155, reg34156)
		reg34158 := PrimCons(reg34140, reg34157)
		reg34159 := PrimCons(reg34139, reg34158)
		reg34160 := PrimCons(reg34134, reg34159)
		reg34161 := MakeSymbol("shen.*signedfuncs*")
		reg34162 := PrimValue(reg34161)
		reg34163 := PrimCons(reg34160, reg34162)
		reg34164 := PrimSet(reg34133, reg34163)
		_ = reg34164
		reg34165 := MakeSymbol("shen.*signedfuncs*")
		reg34166 := MakeSymbol("kill")
		reg34167 := MakeSymbol("-->")
		reg34168 := MakeSymbol("A")
		reg34169 := Nil
		reg34170 := PrimCons(reg34168, reg34169)
		reg34171 := PrimCons(reg34167, reg34170)
		reg34172 := PrimCons(reg34166, reg34171)
		reg34173 := MakeSymbol("shen.*signedfuncs*")
		reg34174 := PrimValue(reg34173)
		reg34175 := PrimCons(reg34172, reg34174)
		reg34176 := PrimSet(reg34165, reg34175)
		_ = reg34176
		reg34177 := MakeSymbol("shen.*signedfuncs*")
		reg34178 := MakeSymbol("language")
		reg34179 := MakeSymbol("-->")
		reg34180 := MakeSymbol("string")
		reg34181 := Nil
		reg34182 := PrimCons(reg34180, reg34181)
		reg34183 := PrimCons(reg34179, reg34182)
		reg34184 := PrimCons(reg34178, reg34183)
		reg34185 := MakeSymbol("shen.*signedfuncs*")
		reg34186 := PrimValue(reg34185)
		reg34187 := PrimCons(reg34184, reg34186)
		reg34188 := PrimSet(reg34177, reg34187)
		_ = reg34188
		reg34189 := MakeSymbol("shen.*signedfuncs*")
		reg34190 := MakeSymbol("length")
		reg34191 := MakeSymbol("list")
		reg34192 := MakeSymbol("A")
		reg34193 := Nil
		reg34194 := PrimCons(reg34192, reg34193)
		reg34195 := PrimCons(reg34191, reg34194)
		reg34196 := MakeSymbol("-->")
		reg34197 := MakeSymbol("number")
		reg34198 := Nil
		reg34199 := PrimCons(reg34197, reg34198)
		reg34200 := PrimCons(reg34196, reg34199)
		reg34201 := PrimCons(reg34195, reg34200)
		reg34202 := PrimCons(reg34190, reg34201)
		reg34203 := MakeSymbol("shen.*signedfuncs*")
		reg34204 := PrimValue(reg34203)
		reg34205 := PrimCons(reg34202, reg34204)
		reg34206 := PrimSet(reg34189, reg34205)
		_ = reg34206
		reg34207 := MakeSymbol("shen.*signedfuncs*")
		reg34208 := MakeSymbol("limit")
		reg34209 := MakeSymbol("vector")
		reg34210 := MakeSymbol("A")
		reg34211 := Nil
		reg34212 := PrimCons(reg34210, reg34211)
		reg34213 := PrimCons(reg34209, reg34212)
		reg34214 := MakeSymbol("-->")
		reg34215 := MakeSymbol("number")
		reg34216 := Nil
		reg34217 := PrimCons(reg34215, reg34216)
		reg34218 := PrimCons(reg34214, reg34217)
		reg34219 := PrimCons(reg34213, reg34218)
		reg34220 := PrimCons(reg34208, reg34219)
		reg34221 := MakeSymbol("shen.*signedfuncs*")
		reg34222 := PrimValue(reg34221)
		reg34223 := PrimCons(reg34220, reg34222)
		reg34224 := PrimSet(reg34207, reg34223)
		_ = reg34224
		reg34225 := MakeSymbol("shen.*signedfuncs*")
		reg34226 := MakeSymbol("load")
		reg34227 := MakeSymbol("string")
		reg34228 := MakeSymbol("-->")
		reg34229 := MakeSymbol("symbol")
		reg34230 := Nil
		reg34231 := PrimCons(reg34229, reg34230)
		reg34232 := PrimCons(reg34228, reg34231)
		reg34233 := PrimCons(reg34227, reg34232)
		reg34234 := PrimCons(reg34226, reg34233)
		reg34235 := MakeSymbol("shen.*signedfuncs*")
		reg34236 := PrimValue(reg34235)
		reg34237 := PrimCons(reg34234, reg34236)
		reg34238 := PrimSet(reg34225, reg34237)
		_ = reg34238
		reg34239 := MakeSymbol("shen.*signedfuncs*")
		reg34240 := MakeSymbol("map")
		reg34241 := MakeSymbol("A")
		reg34242 := MakeSymbol("-->")
		reg34243 := MakeSymbol("B")
		reg34244 := Nil
		reg34245 := PrimCons(reg34243, reg34244)
		reg34246 := PrimCons(reg34242, reg34245)
		reg34247 := PrimCons(reg34241, reg34246)
		reg34248 := MakeSymbol("-->")
		reg34249 := MakeSymbol("list")
		reg34250 := MakeSymbol("A")
		reg34251 := Nil
		reg34252 := PrimCons(reg34250, reg34251)
		reg34253 := PrimCons(reg34249, reg34252)
		reg34254 := MakeSymbol("-->")
		reg34255 := MakeSymbol("list")
		reg34256 := MakeSymbol("B")
		reg34257 := Nil
		reg34258 := PrimCons(reg34256, reg34257)
		reg34259 := PrimCons(reg34255, reg34258)
		reg34260 := Nil
		reg34261 := PrimCons(reg34259, reg34260)
		reg34262 := PrimCons(reg34254, reg34261)
		reg34263 := PrimCons(reg34253, reg34262)
		reg34264 := Nil
		reg34265 := PrimCons(reg34263, reg34264)
		reg34266 := PrimCons(reg34248, reg34265)
		reg34267 := PrimCons(reg34247, reg34266)
		reg34268 := PrimCons(reg34240, reg34267)
		reg34269 := MakeSymbol("shen.*signedfuncs*")
		reg34270 := PrimValue(reg34269)
		reg34271 := PrimCons(reg34268, reg34270)
		reg34272 := PrimSet(reg34239, reg34271)
		_ = reg34272
		reg34273 := MakeSymbol("shen.*signedfuncs*")
		reg34274 := MakeSymbol("mapcan")
		reg34275 := MakeSymbol("A")
		reg34276 := MakeSymbol("-->")
		reg34277 := MakeSymbol("list")
		reg34278 := MakeSymbol("B")
		reg34279 := Nil
		reg34280 := PrimCons(reg34278, reg34279)
		reg34281 := PrimCons(reg34277, reg34280)
		reg34282 := Nil
		reg34283 := PrimCons(reg34281, reg34282)
		reg34284 := PrimCons(reg34276, reg34283)
		reg34285 := PrimCons(reg34275, reg34284)
		reg34286 := MakeSymbol("-->")
		reg34287 := MakeSymbol("list")
		reg34288 := MakeSymbol("A")
		reg34289 := Nil
		reg34290 := PrimCons(reg34288, reg34289)
		reg34291 := PrimCons(reg34287, reg34290)
		reg34292 := MakeSymbol("-->")
		reg34293 := MakeSymbol("list")
		reg34294 := MakeSymbol("B")
		reg34295 := Nil
		reg34296 := PrimCons(reg34294, reg34295)
		reg34297 := PrimCons(reg34293, reg34296)
		reg34298 := Nil
		reg34299 := PrimCons(reg34297, reg34298)
		reg34300 := PrimCons(reg34292, reg34299)
		reg34301 := PrimCons(reg34291, reg34300)
		reg34302 := Nil
		reg34303 := PrimCons(reg34301, reg34302)
		reg34304 := PrimCons(reg34286, reg34303)
		reg34305 := PrimCons(reg34285, reg34304)
		reg34306 := PrimCons(reg34274, reg34305)
		reg34307 := MakeSymbol("shen.*signedfuncs*")
		reg34308 := PrimValue(reg34307)
		reg34309 := PrimCons(reg34306, reg34308)
		reg34310 := PrimSet(reg34273, reg34309)
		_ = reg34310
		reg34311 := MakeSymbol("shen.*signedfuncs*")
		reg34312 := MakeSymbol("maxinferences")
		reg34313 := MakeSymbol("number")
		reg34314 := MakeSymbol("-->")
		reg34315 := MakeSymbol("number")
		reg34316 := Nil
		reg34317 := PrimCons(reg34315, reg34316)
		reg34318 := PrimCons(reg34314, reg34317)
		reg34319 := PrimCons(reg34313, reg34318)
		reg34320 := PrimCons(reg34312, reg34319)
		reg34321 := MakeSymbol("shen.*signedfuncs*")
		reg34322 := PrimValue(reg34321)
		reg34323 := PrimCons(reg34320, reg34322)
		reg34324 := PrimSet(reg34311, reg34323)
		_ = reg34324
		reg34325 := MakeSymbol("shen.*signedfuncs*")
		reg34326 := MakeSymbol("n->string")
		reg34327 := MakeSymbol("number")
		reg34328 := MakeSymbol("-->")
		reg34329 := MakeSymbol("string")
		reg34330 := Nil
		reg34331 := PrimCons(reg34329, reg34330)
		reg34332 := PrimCons(reg34328, reg34331)
		reg34333 := PrimCons(reg34327, reg34332)
		reg34334 := PrimCons(reg34326, reg34333)
		reg34335 := MakeSymbol("shen.*signedfuncs*")
		reg34336 := PrimValue(reg34335)
		reg34337 := PrimCons(reg34334, reg34336)
		reg34338 := PrimSet(reg34325, reg34337)
		_ = reg34338
		reg34339 := MakeSymbol("shen.*signedfuncs*")
		reg34340 := MakeSymbol("nl")
		reg34341 := MakeSymbol("number")
		reg34342 := MakeSymbol("-->")
		reg34343 := MakeSymbol("number")
		reg34344 := Nil
		reg34345 := PrimCons(reg34343, reg34344)
		reg34346 := PrimCons(reg34342, reg34345)
		reg34347 := PrimCons(reg34341, reg34346)
		reg34348 := PrimCons(reg34340, reg34347)
		reg34349 := MakeSymbol("shen.*signedfuncs*")
		reg34350 := PrimValue(reg34349)
		reg34351 := PrimCons(reg34348, reg34350)
		reg34352 := PrimSet(reg34339, reg34351)
		_ = reg34352
		reg34353 := MakeSymbol("shen.*signedfuncs*")
		reg34354 := MakeSymbol("not")
		reg34355 := MakeSymbol("boolean")
		reg34356 := MakeSymbol("-->")
		reg34357 := MakeSymbol("boolean")
		reg34358 := Nil
		reg34359 := PrimCons(reg34357, reg34358)
		reg34360 := PrimCons(reg34356, reg34359)
		reg34361 := PrimCons(reg34355, reg34360)
		reg34362 := PrimCons(reg34354, reg34361)
		reg34363 := MakeSymbol("shen.*signedfuncs*")
		reg34364 := PrimValue(reg34363)
		reg34365 := PrimCons(reg34362, reg34364)
		reg34366 := PrimSet(reg34353, reg34365)
		_ = reg34366
		reg34367 := MakeSymbol("shen.*signedfuncs*")
		reg34368 := MakeSymbol("nth")
		reg34369 := MakeSymbol("number")
		reg34370 := MakeSymbol("-->")
		reg34371 := MakeSymbol("list")
		reg34372 := MakeSymbol("A")
		reg34373 := Nil
		reg34374 := PrimCons(reg34372, reg34373)
		reg34375 := PrimCons(reg34371, reg34374)
		reg34376 := MakeSymbol("-->")
		reg34377 := MakeSymbol("A")
		reg34378 := Nil
		reg34379 := PrimCons(reg34377, reg34378)
		reg34380 := PrimCons(reg34376, reg34379)
		reg34381 := PrimCons(reg34375, reg34380)
		reg34382 := Nil
		reg34383 := PrimCons(reg34381, reg34382)
		reg34384 := PrimCons(reg34370, reg34383)
		reg34385 := PrimCons(reg34369, reg34384)
		reg34386 := PrimCons(reg34368, reg34385)
		reg34387 := MakeSymbol("shen.*signedfuncs*")
		reg34388 := PrimValue(reg34387)
		reg34389 := PrimCons(reg34386, reg34388)
		reg34390 := PrimSet(reg34367, reg34389)
		_ = reg34390
		reg34391 := MakeSymbol("shen.*signedfuncs*")
		reg34392 := MakeSymbol("number?")
		reg34393 := MakeSymbol("A")
		reg34394 := MakeSymbol("-->")
		reg34395 := MakeSymbol("boolean")
		reg34396 := Nil
		reg34397 := PrimCons(reg34395, reg34396)
		reg34398 := PrimCons(reg34394, reg34397)
		reg34399 := PrimCons(reg34393, reg34398)
		reg34400 := PrimCons(reg34392, reg34399)
		reg34401 := MakeSymbol("shen.*signedfuncs*")
		reg34402 := PrimValue(reg34401)
		reg34403 := PrimCons(reg34400, reg34402)
		reg34404 := PrimSet(reg34391, reg34403)
		_ = reg34404
		reg34405 := MakeSymbol("shen.*signedfuncs*")
		reg34406 := MakeSymbol("occurrences")
		reg34407 := MakeSymbol("A")
		reg34408 := MakeSymbol("-->")
		reg34409 := MakeSymbol("B")
		reg34410 := MakeSymbol("-->")
		reg34411 := MakeSymbol("number")
		reg34412 := Nil
		reg34413 := PrimCons(reg34411, reg34412)
		reg34414 := PrimCons(reg34410, reg34413)
		reg34415 := PrimCons(reg34409, reg34414)
		reg34416 := Nil
		reg34417 := PrimCons(reg34415, reg34416)
		reg34418 := PrimCons(reg34408, reg34417)
		reg34419 := PrimCons(reg34407, reg34418)
		reg34420 := PrimCons(reg34406, reg34419)
		reg34421 := MakeSymbol("shen.*signedfuncs*")
		reg34422 := PrimValue(reg34421)
		reg34423 := PrimCons(reg34420, reg34422)
		reg34424 := PrimSet(reg34405, reg34423)
		_ = reg34424
		reg34425 := MakeSymbol("shen.*signedfuncs*")
		reg34426 := MakeSymbol("occurs-check")
		reg34427 := MakeSymbol("symbol")
		reg34428 := MakeSymbol("-->")
		reg34429 := MakeSymbol("boolean")
		reg34430 := Nil
		reg34431 := PrimCons(reg34429, reg34430)
		reg34432 := PrimCons(reg34428, reg34431)
		reg34433 := PrimCons(reg34427, reg34432)
		reg34434 := PrimCons(reg34426, reg34433)
		reg34435 := MakeSymbol("shen.*signedfuncs*")
		reg34436 := PrimValue(reg34435)
		reg34437 := PrimCons(reg34434, reg34436)
		reg34438 := PrimSet(reg34425, reg34437)
		_ = reg34438
		reg34439 := MakeSymbol("shen.*signedfuncs*")
		reg34440 := MakeSymbol("optimise")
		reg34441 := MakeSymbol("symbol")
		reg34442 := MakeSymbol("-->")
		reg34443 := MakeSymbol("boolean")
		reg34444 := Nil
		reg34445 := PrimCons(reg34443, reg34444)
		reg34446 := PrimCons(reg34442, reg34445)
		reg34447 := PrimCons(reg34441, reg34446)
		reg34448 := PrimCons(reg34440, reg34447)
		reg34449 := MakeSymbol("shen.*signedfuncs*")
		reg34450 := PrimValue(reg34449)
		reg34451 := PrimCons(reg34448, reg34450)
		reg34452 := PrimSet(reg34439, reg34451)
		_ = reg34452
		reg34453 := MakeSymbol("shen.*signedfuncs*")
		reg34454 := MakeSymbol("or")
		reg34455 := MakeSymbol("boolean")
		reg34456 := MakeSymbol("-->")
		reg34457 := MakeSymbol("boolean")
		reg34458 := MakeSymbol("-->")
		reg34459 := MakeSymbol("boolean")
		reg34460 := Nil
		reg34461 := PrimCons(reg34459, reg34460)
		reg34462 := PrimCons(reg34458, reg34461)
		reg34463 := PrimCons(reg34457, reg34462)
		reg34464 := Nil
		reg34465 := PrimCons(reg34463, reg34464)
		reg34466 := PrimCons(reg34456, reg34465)
		reg34467 := PrimCons(reg34455, reg34466)
		reg34468 := PrimCons(reg34454, reg34467)
		reg34469 := MakeSymbol("shen.*signedfuncs*")
		reg34470 := PrimValue(reg34469)
		reg34471 := PrimCons(reg34468, reg34470)
		reg34472 := PrimSet(reg34453, reg34471)
		_ = reg34472
		reg34473 := MakeSymbol("shen.*signedfuncs*")
		reg34474 := MakeSymbol("os")
		reg34475 := MakeSymbol("-->")
		reg34476 := MakeSymbol("string")
		reg34477 := Nil
		reg34478 := PrimCons(reg34476, reg34477)
		reg34479 := PrimCons(reg34475, reg34478)
		reg34480 := PrimCons(reg34474, reg34479)
		reg34481 := MakeSymbol("shen.*signedfuncs*")
		reg34482 := PrimValue(reg34481)
		reg34483 := PrimCons(reg34480, reg34482)
		reg34484 := PrimSet(reg34473, reg34483)
		_ = reg34484
		reg34485 := MakeSymbol("shen.*signedfuncs*")
		reg34486 := MakeSymbol("package?")
		reg34487 := MakeSymbol("symbol")
		reg34488 := MakeSymbol("-->")
		reg34489 := MakeSymbol("boolean")
		reg34490 := Nil
		reg34491 := PrimCons(reg34489, reg34490)
		reg34492 := PrimCons(reg34488, reg34491)
		reg34493 := PrimCons(reg34487, reg34492)
		reg34494 := PrimCons(reg34486, reg34493)
		reg34495 := MakeSymbol("shen.*signedfuncs*")
		reg34496 := PrimValue(reg34495)
		reg34497 := PrimCons(reg34494, reg34496)
		reg34498 := PrimSet(reg34485, reg34497)
		_ = reg34498
		reg34499 := MakeSymbol("shen.*signedfuncs*")
		reg34500 := MakeSymbol("port")
		reg34501 := MakeSymbol("-->")
		reg34502 := MakeSymbol("string")
		reg34503 := Nil
		reg34504 := PrimCons(reg34502, reg34503)
		reg34505 := PrimCons(reg34501, reg34504)
		reg34506 := PrimCons(reg34500, reg34505)
		reg34507 := MakeSymbol("shen.*signedfuncs*")
		reg34508 := PrimValue(reg34507)
		reg34509 := PrimCons(reg34506, reg34508)
		reg34510 := PrimSet(reg34499, reg34509)
		_ = reg34510
		reg34511 := MakeSymbol("shen.*signedfuncs*")
		reg34512 := MakeSymbol("porters")
		reg34513 := MakeSymbol("-->")
		reg34514 := MakeSymbol("string")
		reg34515 := Nil
		reg34516 := PrimCons(reg34514, reg34515)
		reg34517 := PrimCons(reg34513, reg34516)
		reg34518 := PrimCons(reg34512, reg34517)
		reg34519 := MakeSymbol("shen.*signedfuncs*")
		reg34520 := PrimValue(reg34519)
		reg34521 := PrimCons(reg34518, reg34520)
		reg34522 := PrimSet(reg34511, reg34521)
		_ = reg34522
		reg34523 := MakeSymbol("shen.*signedfuncs*")
		reg34524 := MakeSymbol("pos")
		reg34525 := MakeSymbol("string")
		reg34526 := MakeSymbol("-->")
		reg34527 := MakeSymbol("number")
		reg34528 := MakeSymbol("-->")
		reg34529 := MakeSymbol("string")
		reg34530 := Nil
		reg34531 := PrimCons(reg34529, reg34530)
		reg34532 := PrimCons(reg34528, reg34531)
		reg34533 := PrimCons(reg34527, reg34532)
		reg34534 := Nil
		reg34535 := PrimCons(reg34533, reg34534)
		reg34536 := PrimCons(reg34526, reg34535)
		reg34537 := PrimCons(reg34525, reg34536)
		reg34538 := PrimCons(reg34524, reg34537)
		reg34539 := MakeSymbol("shen.*signedfuncs*")
		reg34540 := PrimValue(reg34539)
		reg34541 := PrimCons(reg34538, reg34540)
		reg34542 := PrimSet(reg34523, reg34541)
		_ = reg34542
		reg34543 := MakeSymbol("shen.*signedfuncs*")
		reg34544 := MakeSymbol("pr")
		reg34545 := MakeSymbol("string")
		reg34546 := MakeSymbol("-->")
		reg34547 := MakeSymbol("stream")
		reg34548 := MakeSymbol("out")
		reg34549 := Nil
		reg34550 := PrimCons(reg34548, reg34549)
		reg34551 := PrimCons(reg34547, reg34550)
		reg34552 := MakeSymbol("-->")
		reg34553 := MakeSymbol("string")
		reg34554 := Nil
		reg34555 := PrimCons(reg34553, reg34554)
		reg34556 := PrimCons(reg34552, reg34555)
		reg34557 := PrimCons(reg34551, reg34556)
		reg34558 := Nil
		reg34559 := PrimCons(reg34557, reg34558)
		reg34560 := PrimCons(reg34546, reg34559)
		reg34561 := PrimCons(reg34545, reg34560)
		reg34562 := PrimCons(reg34544, reg34561)
		reg34563 := MakeSymbol("shen.*signedfuncs*")
		reg34564 := PrimValue(reg34563)
		reg34565 := PrimCons(reg34562, reg34564)
		reg34566 := PrimSet(reg34543, reg34565)
		_ = reg34566
		reg34567 := MakeSymbol("shen.*signedfuncs*")
		reg34568 := MakeSymbol("print")
		reg34569 := MakeSymbol("A")
		reg34570 := MakeSymbol("-->")
		reg34571 := MakeSymbol("A")
		reg34572 := Nil
		reg34573 := PrimCons(reg34571, reg34572)
		reg34574 := PrimCons(reg34570, reg34573)
		reg34575 := PrimCons(reg34569, reg34574)
		reg34576 := PrimCons(reg34568, reg34575)
		reg34577 := MakeSymbol("shen.*signedfuncs*")
		reg34578 := PrimValue(reg34577)
		reg34579 := PrimCons(reg34576, reg34578)
		reg34580 := PrimSet(reg34567, reg34579)
		_ = reg34580
		reg34581 := MakeSymbol("shen.*signedfuncs*")
		reg34582 := MakeSymbol("profile")
		reg34583 := MakeSymbol("A")
		reg34584 := MakeSymbol("-->")
		reg34585 := MakeSymbol("B")
		reg34586 := Nil
		reg34587 := PrimCons(reg34585, reg34586)
		reg34588 := PrimCons(reg34584, reg34587)
		reg34589 := PrimCons(reg34583, reg34588)
		reg34590 := MakeSymbol("-->")
		reg34591 := MakeSymbol("A")
		reg34592 := MakeSymbol("-->")
		reg34593 := MakeSymbol("B")
		reg34594 := Nil
		reg34595 := PrimCons(reg34593, reg34594)
		reg34596 := PrimCons(reg34592, reg34595)
		reg34597 := PrimCons(reg34591, reg34596)
		reg34598 := Nil
		reg34599 := PrimCons(reg34597, reg34598)
		reg34600 := PrimCons(reg34590, reg34599)
		reg34601 := PrimCons(reg34589, reg34600)
		reg34602 := PrimCons(reg34582, reg34601)
		reg34603 := MakeSymbol("shen.*signedfuncs*")
		reg34604 := PrimValue(reg34603)
		reg34605 := PrimCons(reg34602, reg34604)
		reg34606 := PrimSet(reg34581, reg34605)
		_ = reg34606
		reg34607 := MakeSymbol("shen.*signedfuncs*")
		reg34608 := MakeSymbol("preclude")
		reg34609 := MakeSymbol("list")
		reg34610 := MakeSymbol("symbol")
		reg34611 := Nil
		reg34612 := PrimCons(reg34610, reg34611)
		reg34613 := PrimCons(reg34609, reg34612)
		reg34614 := MakeSymbol("-->")
		reg34615 := MakeSymbol("list")
		reg34616 := MakeSymbol("symbol")
		reg34617 := Nil
		reg34618 := PrimCons(reg34616, reg34617)
		reg34619 := PrimCons(reg34615, reg34618)
		reg34620 := Nil
		reg34621 := PrimCons(reg34619, reg34620)
		reg34622 := PrimCons(reg34614, reg34621)
		reg34623 := PrimCons(reg34613, reg34622)
		reg34624 := PrimCons(reg34608, reg34623)
		reg34625 := MakeSymbol("shen.*signedfuncs*")
		reg34626 := PrimValue(reg34625)
		reg34627 := PrimCons(reg34624, reg34626)
		reg34628 := PrimSet(reg34607, reg34627)
		_ = reg34628
		reg34629 := MakeSymbol("shen.*signedfuncs*")
		reg34630 := MakeSymbol("shen.proc-nl")
		reg34631 := MakeSymbol("string")
		reg34632 := MakeSymbol("-->")
		reg34633 := MakeSymbol("string")
		reg34634 := Nil
		reg34635 := PrimCons(reg34633, reg34634)
		reg34636 := PrimCons(reg34632, reg34635)
		reg34637 := PrimCons(reg34631, reg34636)
		reg34638 := PrimCons(reg34630, reg34637)
		reg34639 := MakeSymbol("shen.*signedfuncs*")
		reg34640 := PrimValue(reg34639)
		reg34641 := PrimCons(reg34638, reg34640)
		reg34642 := PrimSet(reg34629, reg34641)
		_ = reg34642
		reg34643 := MakeSymbol("shen.*signedfuncs*")
		reg34644 := MakeSymbol("profile-results")
		reg34645 := MakeSymbol("A")
		reg34646 := MakeSymbol("-->")
		reg34647 := MakeSymbol("B")
		reg34648 := Nil
		reg34649 := PrimCons(reg34647, reg34648)
		reg34650 := PrimCons(reg34646, reg34649)
		reg34651 := PrimCons(reg34645, reg34650)
		reg34652 := MakeSymbol("-->")
		reg34653 := MakeSymbol("A")
		reg34654 := MakeSymbol("-->")
		reg34655 := MakeSymbol("B")
		reg34656 := Nil
		reg34657 := PrimCons(reg34655, reg34656)
		reg34658 := PrimCons(reg34654, reg34657)
		reg34659 := PrimCons(reg34653, reg34658)
		reg34660 := MakeSymbol("*")
		reg34661 := MakeSymbol("number")
		reg34662 := Nil
		reg34663 := PrimCons(reg34661, reg34662)
		reg34664 := PrimCons(reg34660, reg34663)
		reg34665 := PrimCons(reg34659, reg34664)
		reg34666 := Nil
		reg34667 := PrimCons(reg34665, reg34666)
		reg34668 := PrimCons(reg34652, reg34667)
		reg34669 := PrimCons(reg34651, reg34668)
		reg34670 := PrimCons(reg34644, reg34669)
		reg34671 := MakeSymbol("shen.*signedfuncs*")
		reg34672 := PrimValue(reg34671)
		reg34673 := PrimCons(reg34670, reg34672)
		reg34674 := PrimSet(reg34643, reg34673)
		_ = reg34674
		reg34675 := MakeSymbol("shen.*signedfuncs*")
		reg34676 := MakeSymbol("protect")
		reg34677 := MakeSymbol("symbol")
		reg34678 := MakeSymbol("-->")
		reg34679 := MakeSymbol("symbol")
		reg34680 := Nil
		reg34681 := PrimCons(reg34679, reg34680)
		reg34682 := PrimCons(reg34678, reg34681)
		reg34683 := PrimCons(reg34677, reg34682)
		reg34684 := PrimCons(reg34676, reg34683)
		reg34685 := MakeSymbol("shen.*signedfuncs*")
		reg34686 := PrimValue(reg34685)
		reg34687 := PrimCons(reg34684, reg34686)
		reg34688 := PrimSet(reg34675, reg34687)
		_ = reg34688
		reg34689 := MakeSymbol("shen.*signedfuncs*")
		reg34690 := MakeSymbol("preclude-all-but")
		reg34691 := MakeSymbol("list")
		reg34692 := MakeSymbol("symbol")
		reg34693 := Nil
		reg34694 := PrimCons(reg34692, reg34693)
		reg34695 := PrimCons(reg34691, reg34694)
		reg34696 := MakeSymbol("-->")
		reg34697 := MakeSymbol("list")
		reg34698 := MakeSymbol("symbol")
		reg34699 := Nil
		reg34700 := PrimCons(reg34698, reg34699)
		reg34701 := PrimCons(reg34697, reg34700)
		reg34702 := Nil
		reg34703 := PrimCons(reg34701, reg34702)
		reg34704 := PrimCons(reg34696, reg34703)
		reg34705 := PrimCons(reg34695, reg34704)
		reg34706 := PrimCons(reg34690, reg34705)
		reg34707 := MakeSymbol("shen.*signedfuncs*")
		reg34708 := PrimValue(reg34707)
		reg34709 := PrimCons(reg34706, reg34708)
		reg34710 := PrimSet(reg34689, reg34709)
		_ = reg34710
		reg34711 := MakeSymbol("shen.*signedfuncs*")
		reg34712 := MakeSymbol("shen.prhush")
		reg34713 := MakeSymbol("string")
		reg34714 := MakeSymbol("-->")
		reg34715 := MakeSymbol("stream")
		reg34716 := MakeSymbol("out")
		reg34717 := Nil
		reg34718 := PrimCons(reg34716, reg34717)
		reg34719 := PrimCons(reg34715, reg34718)
		reg34720 := MakeSymbol("-->")
		reg34721 := MakeSymbol("string")
		reg34722 := Nil
		reg34723 := PrimCons(reg34721, reg34722)
		reg34724 := PrimCons(reg34720, reg34723)
		reg34725 := PrimCons(reg34719, reg34724)
		reg34726 := Nil
		reg34727 := PrimCons(reg34725, reg34726)
		reg34728 := PrimCons(reg34714, reg34727)
		reg34729 := PrimCons(reg34713, reg34728)
		reg34730 := PrimCons(reg34712, reg34729)
		reg34731 := MakeSymbol("shen.*signedfuncs*")
		reg34732 := PrimValue(reg34731)
		reg34733 := PrimCons(reg34730, reg34732)
		reg34734 := PrimSet(reg34711, reg34733)
		_ = reg34734
		reg34735 := MakeSymbol("shen.*signedfuncs*")
		reg34736 := MakeSymbol("ps")
		reg34737 := MakeSymbol("symbol")
		reg34738 := MakeSymbol("-->")
		reg34739 := MakeSymbol("list")
		reg34740 := MakeSymbol("unit")
		reg34741 := Nil
		reg34742 := PrimCons(reg34740, reg34741)
		reg34743 := PrimCons(reg34739, reg34742)
		reg34744 := Nil
		reg34745 := PrimCons(reg34743, reg34744)
		reg34746 := PrimCons(reg34738, reg34745)
		reg34747 := PrimCons(reg34737, reg34746)
		reg34748 := PrimCons(reg34736, reg34747)
		reg34749 := MakeSymbol("shen.*signedfuncs*")
		reg34750 := PrimValue(reg34749)
		reg34751 := PrimCons(reg34748, reg34750)
		reg34752 := PrimSet(reg34735, reg34751)
		_ = reg34752
		reg34753 := MakeSymbol("shen.*signedfuncs*")
		reg34754 := MakeSymbol("read")
		reg34755 := MakeSymbol("stream")
		reg34756 := MakeSymbol("in")
		reg34757 := Nil
		reg34758 := PrimCons(reg34756, reg34757)
		reg34759 := PrimCons(reg34755, reg34758)
		reg34760 := MakeSymbol("-->")
		reg34761 := MakeSymbol("unit")
		reg34762 := Nil
		reg34763 := PrimCons(reg34761, reg34762)
		reg34764 := PrimCons(reg34760, reg34763)
		reg34765 := PrimCons(reg34759, reg34764)
		reg34766 := PrimCons(reg34754, reg34765)
		reg34767 := MakeSymbol("shen.*signedfuncs*")
		reg34768 := PrimValue(reg34767)
		reg34769 := PrimCons(reg34766, reg34768)
		reg34770 := PrimSet(reg34753, reg34769)
		_ = reg34770
		reg34771 := MakeSymbol("shen.*signedfuncs*")
		reg34772 := MakeSymbol("read-byte")
		reg34773 := MakeSymbol("stream")
		reg34774 := MakeSymbol("in")
		reg34775 := Nil
		reg34776 := PrimCons(reg34774, reg34775)
		reg34777 := PrimCons(reg34773, reg34776)
		reg34778 := MakeSymbol("-->")
		reg34779 := MakeSymbol("number")
		reg34780 := Nil
		reg34781 := PrimCons(reg34779, reg34780)
		reg34782 := PrimCons(reg34778, reg34781)
		reg34783 := PrimCons(reg34777, reg34782)
		reg34784 := PrimCons(reg34772, reg34783)
		reg34785 := MakeSymbol("shen.*signedfuncs*")
		reg34786 := PrimValue(reg34785)
		reg34787 := PrimCons(reg34784, reg34786)
		reg34788 := PrimSet(reg34771, reg34787)
		_ = reg34788
		reg34789 := MakeSymbol("shen.*signedfuncs*")
		reg34790 := MakeSymbol("read-file-as-bytelist")
		reg34791 := MakeSymbol("string")
		reg34792 := MakeSymbol("-->")
		reg34793 := MakeSymbol("list")
		reg34794 := MakeSymbol("number")
		reg34795 := Nil
		reg34796 := PrimCons(reg34794, reg34795)
		reg34797 := PrimCons(reg34793, reg34796)
		reg34798 := Nil
		reg34799 := PrimCons(reg34797, reg34798)
		reg34800 := PrimCons(reg34792, reg34799)
		reg34801 := PrimCons(reg34791, reg34800)
		reg34802 := PrimCons(reg34790, reg34801)
		reg34803 := MakeSymbol("shen.*signedfuncs*")
		reg34804 := PrimValue(reg34803)
		reg34805 := PrimCons(reg34802, reg34804)
		reg34806 := PrimSet(reg34789, reg34805)
		_ = reg34806
		reg34807 := MakeSymbol("shen.*signedfuncs*")
		reg34808 := MakeSymbol("read-file-as-string")
		reg34809 := MakeSymbol("string")
		reg34810 := MakeSymbol("-->")
		reg34811 := MakeSymbol("string")
		reg34812 := Nil
		reg34813 := PrimCons(reg34811, reg34812)
		reg34814 := PrimCons(reg34810, reg34813)
		reg34815 := PrimCons(reg34809, reg34814)
		reg34816 := PrimCons(reg34808, reg34815)
		reg34817 := MakeSymbol("shen.*signedfuncs*")
		reg34818 := PrimValue(reg34817)
		reg34819 := PrimCons(reg34816, reg34818)
		reg34820 := PrimSet(reg34807, reg34819)
		_ = reg34820
		reg34821 := MakeSymbol("shen.*signedfuncs*")
		reg34822 := MakeSymbol("read-file")
		reg34823 := MakeSymbol("string")
		reg34824 := MakeSymbol("-->")
		reg34825 := MakeSymbol("list")
		reg34826 := MakeSymbol("unit")
		reg34827 := Nil
		reg34828 := PrimCons(reg34826, reg34827)
		reg34829 := PrimCons(reg34825, reg34828)
		reg34830 := Nil
		reg34831 := PrimCons(reg34829, reg34830)
		reg34832 := PrimCons(reg34824, reg34831)
		reg34833 := PrimCons(reg34823, reg34832)
		reg34834 := PrimCons(reg34822, reg34833)
		reg34835 := MakeSymbol("shen.*signedfuncs*")
		reg34836 := PrimValue(reg34835)
		reg34837 := PrimCons(reg34834, reg34836)
		reg34838 := PrimSet(reg34821, reg34837)
		_ = reg34838
		reg34839 := MakeSymbol("shen.*signedfuncs*")
		reg34840 := MakeSymbol("read-from-string")
		reg34841 := MakeSymbol("string")
		reg34842 := MakeSymbol("-->")
		reg34843 := MakeSymbol("list")
		reg34844 := MakeSymbol("unit")
		reg34845 := Nil
		reg34846 := PrimCons(reg34844, reg34845)
		reg34847 := PrimCons(reg34843, reg34846)
		reg34848 := Nil
		reg34849 := PrimCons(reg34847, reg34848)
		reg34850 := PrimCons(reg34842, reg34849)
		reg34851 := PrimCons(reg34841, reg34850)
		reg34852 := PrimCons(reg34840, reg34851)
		reg34853 := MakeSymbol("shen.*signedfuncs*")
		reg34854 := PrimValue(reg34853)
		reg34855 := PrimCons(reg34852, reg34854)
		reg34856 := PrimSet(reg34839, reg34855)
		_ = reg34856
		reg34857 := MakeSymbol("shen.*signedfuncs*")
		reg34858 := MakeSymbol("release")
		reg34859 := MakeSymbol("-->")
		reg34860 := MakeSymbol("string")
		reg34861 := Nil
		reg34862 := PrimCons(reg34860, reg34861)
		reg34863 := PrimCons(reg34859, reg34862)
		reg34864 := PrimCons(reg34858, reg34863)
		reg34865 := MakeSymbol("shen.*signedfuncs*")
		reg34866 := PrimValue(reg34865)
		reg34867 := PrimCons(reg34864, reg34866)
		reg34868 := PrimSet(reg34857, reg34867)
		_ = reg34868
		reg34869 := MakeSymbol("shen.*signedfuncs*")
		reg34870 := MakeSymbol("remove")
		reg34871 := MakeSymbol("A")
		reg34872 := MakeSymbol("-->")
		reg34873 := MakeSymbol("list")
		reg34874 := MakeSymbol("A")
		reg34875 := Nil
		reg34876 := PrimCons(reg34874, reg34875)
		reg34877 := PrimCons(reg34873, reg34876)
		reg34878 := MakeSymbol("-->")
		reg34879 := MakeSymbol("list")
		reg34880 := MakeSymbol("A")
		reg34881 := Nil
		reg34882 := PrimCons(reg34880, reg34881)
		reg34883 := PrimCons(reg34879, reg34882)
		reg34884 := Nil
		reg34885 := PrimCons(reg34883, reg34884)
		reg34886 := PrimCons(reg34878, reg34885)
		reg34887 := PrimCons(reg34877, reg34886)
		reg34888 := Nil
		reg34889 := PrimCons(reg34887, reg34888)
		reg34890 := PrimCons(reg34872, reg34889)
		reg34891 := PrimCons(reg34871, reg34890)
		reg34892 := PrimCons(reg34870, reg34891)
		reg34893 := MakeSymbol("shen.*signedfuncs*")
		reg34894 := PrimValue(reg34893)
		reg34895 := PrimCons(reg34892, reg34894)
		reg34896 := PrimSet(reg34869, reg34895)
		_ = reg34896
		reg34897 := MakeSymbol("shen.*signedfuncs*")
		reg34898 := MakeSymbol("reverse")
		reg34899 := MakeSymbol("list")
		reg34900 := MakeSymbol("A")
		reg34901 := Nil
		reg34902 := PrimCons(reg34900, reg34901)
		reg34903 := PrimCons(reg34899, reg34902)
		reg34904 := MakeSymbol("-->")
		reg34905 := MakeSymbol("list")
		reg34906 := MakeSymbol("A")
		reg34907 := Nil
		reg34908 := PrimCons(reg34906, reg34907)
		reg34909 := PrimCons(reg34905, reg34908)
		reg34910 := Nil
		reg34911 := PrimCons(reg34909, reg34910)
		reg34912 := PrimCons(reg34904, reg34911)
		reg34913 := PrimCons(reg34903, reg34912)
		reg34914 := PrimCons(reg34898, reg34913)
		reg34915 := MakeSymbol("shen.*signedfuncs*")
		reg34916 := PrimValue(reg34915)
		reg34917 := PrimCons(reg34914, reg34916)
		reg34918 := PrimSet(reg34897, reg34917)
		_ = reg34918
		reg34919 := MakeSymbol("shen.*signedfuncs*")
		reg34920 := MakeSymbol("simple-error")
		reg34921 := MakeSymbol("string")
		reg34922 := MakeSymbol("-->")
		reg34923 := MakeSymbol("A")
		reg34924 := Nil
		reg34925 := PrimCons(reg34923, reg34924)
		reg34926 := PrimCons(reg34922, reg34925)
		reg34927 := PrimCons(reg34921, reg34926)
		reg34928 := PrimCons(reg34920, reg34927)
		reg34929 := MakeSymbol("shen.*signedfuncs*")
		reg34930 := PrimValue(reg34929)
		reg34931 := PrimCons(reg34928, reg34930)
		reg34932 := PrimSet(reg34919, reg34931)
		_ = reg34932
		reg34933 := MakeSymbol("shen.*signedfuncs*")
		reg34934 := MakeSymbol("snd")
		reg34935 := MakeSymbol("A")
		reg34936 := MakeSymbol("*")
		reg34937 := MakeSymbol("B")
		reg34938 := Nil
		reg34939 := PrimCons(reg34937, reg34938)
		reg34940 := PrimCons(reg34936, reg34939)
		reg34941 := PrimCons(reg34935, reg34940)
		reg34942 := MakeSymbol("-->")
		reg34943 := MakeSymbol("B")
		reg34944 := Nil
		reg34945 := PrimCons(reg34943, reg34944)
		reg34946 := PrimCons(reg34942, reg34945)
		reg34947 := PrimCons(reg34941, reg34946)
		reg34948 := PrimCons(reg34934, reg34947)
		reg34949 := MakeSymbol("shen.*signedfuncs*")
		reg34950 := PrimValue(reg34949)
		reg34951 := PrimCons(reg34948, reg34950)
		reg34952 := PrimSet(reg34933, reg34951)
		_ = reg34952
		reg34953 := MakeSymbol("shen.*signedfuncs*")
		reg34954 := MakeSymbol("specialise")
		reg34955 := MakeSymbol("symbol")
		reg34956 := MakeSymbol("-->")
		reg34957 := MakeSymbol("symbol")
		reg34958 := Nil
		reg34959 := PrimCons(reg34957, reg34958)
		reg34960 := PrimCons(reg34956, reg34959)
		reg34961 := PrimCons(reg34955, reg34960)
		reg34962 := PrimCons(reg34954, reg34961)
		reg34963 := MakeSymbol("shen.*signedfuncs*")
		reg34964 := PrimValue(reg34963)
		reg34965 := PrimCons(reg34962, reg34964)
		reg34966 := PrimSet(reg34953, reg34965)
		_ = reg34966
		reg34967 := MakeSymbol("shen.*signedfuncs*")
		reg34968 := MakeSymbol("spy")
		reg34969 := MakeSymbol("symbol")
		reg34970 := MakeSymbol("-->")
		reg34971 := MakeSymbol("boolean")
		reg34972 := Nil
		reg34973 := PrimCons(reg34971, reg34972)
		reg34974 := PrimCons(reg34970, reg34973)
		reg34975 := PrimCons(reg34969, reg34974)
		reg34976 := PrimCons(reg34968, reg34975)
		reg34977 := MakeSymbol("shen.*signedfuncs*")
		reg34978 := PrimValue(reg34977)
		reg34979 := PrimCons(reg34976, reg34978)
		reg34980 := PrimSet(reg34967, reg34979)
		_ = reg34980
		reg34981 := MakeSymbol("shen.*signedfuncs*")
		reg34982 := MakeSymbol("step")
		reg34983 := MakeSymbol("symbol")
		reg34984 := MakeSymbol("-->")
		reg34985 := MakeSymbol("boolean")
		reg34986 := Nil
		reg34987 := PrimCons(reg34985, reg34986)
		reg34988 := PrimCons(reg34984, reg34987)
		reg34989 := PrimCons(reg34983, reg34988)
		reg34990 := PrimCons(reg34982, reg34989)
		reg34991 := MakeSymbol("shen.*signedfuncs*")
		reg34992 := PrimValue(reg34991)
		reg34993 := PrimCons(reg34990, reg34992)
		reg34994 := PrimSet(reg34981, reg34993)
		_ = reg34994
		reg34995 := MakeSymbol("shen.*signedfuncs*")
		reg34996 := MakeSymbol("stinput")
		reg34997 := MakeSymbol("-->")
		reg34998 := MakeSymbol("stream")
		reg34999 := MakeSymbol("in")
		reg35000 := Nil
		reg35001 := PrimCons(reg34999, reg35000)
		reg35002 := PrimCons(reg34998, reg35001)
		reg35003 := Nil
		reg35004 := PrimCons(reg35002, reg35003)
		reg35005 := PrimCons(reg34997, reg35004)
		reg35006 := PrimCons(reg34996, reg35005)
		reg35007 := MakeSymbol("shen.*signedfuncs*")
		reg35008 := PrimValue(reg35007)
		reg35009 := PrimCons(reg35006, reg35008)
		reg35010 := PrimSet(reg34995, reg35009)
		_ = reg35010
		reg35011 := MakeSymbol("shen.*signedfuncs*")
		reg35012 := MakeSymbol("sterror")
		reg35013 := MakeSymbol("-->")
		reg35014 := MakeSymbol("stream")
		reg35015 := MakeSymbol("out")
		reg35016 := Nil
		reg35017 := PrimCons(reg35015, reg35016)
		reg35018 := PrimCons(reg35014, reg35017)
		reg35019 := Nil
		reg35020 := PrimCons(reg35018, reg35019)
		reg35021 := PrimCons(reg35013, reg35020)
		reg35022 := PrimCons(reg35012, reg35021)
		reg35023 := MakeSymbol("shen.*signedfuncs*")
		reg35024 := PrimValue(reg35023)
		reg35025 := PrimCons(reg35022, reg35024)
		reg35026 := PrimSet(reg35011, reg35025)
		_ = reg35026
		reg35027 := MakeSymbol("shen.*signedfuncs*")
		reg35028 := MakeSymbol("stoutput")
		reg35029 := MakeSymbol("-->")
		reg35030 := MakeSymbol("stream")
		reg35031 := MakeSymbol("out")
		reg35032 := Nil
		reg35033 := PrimCons(reg35031, reg35032)
		reg35034 := PrimCons(reg35030, reg35033)
		reg35035 := Nil
		reg35036 := PrimCons(reg35034, reg35035)
		reg35037 := PrimCons(reg35029, reg35036)
		reg35038 := PrimCons(reg35028, reg35037)
		reg35039 := MakeSymbol("shen.*signedfuncs*")
		reg35040 := PrimValue(reg35039)
		reg35041 := PrimCons(reg35038, reg35040)
		reg35042 := PrimSet(reg35027, reg35041)
		_ = reg35042
		reg35043 := MakeSymbol("shen.*signedfuncs*")
		reg35044 := MakeSymbol("string?")
		reg35045 := MakeSymbol("A")
		reg35046 := MakeSymbol("-->")
		reg35047 := MakeSymbol("boolean")
		reg35048 := Nil
		reg35049 := PrimCons(reg35047, reg35048)
		reg35050 := PrimCons(reg35046, reg35049)
		reg35051 := PrimCons(reg35045, reg35050)
		reg35052 := PrimCons(reg35044, reg35051)
		reg35053 := MakeSymbol("shen.*signedfuncs*")
		reg35054 := PrimValue(reg35053)
		reg35055 := PrimCons(reg35052, reg35054)
		reg35056 := PrimSet(reg35043, reg35055)
		_ = reg35056
		reg35057 := MakeSymbol("shen.*signedfuncs*")
		reg35058 := MakeSymbol("str")
		reg35059 := MakeSymbol("A")
		reg35060 := MakeSymbol("-->")
		reg35061 := MakeSymbol("string")
		reg35062 := Nil
		reg35063 := PrimCons(reg35061, reg35062)
		reg35064 := PrimCons(reg35060, reg35063)
		reg35065 := PrimCons(reg35059, reg35064)
		reg35066 := PrimCons(reg35058, reg35065)
		reg35067 := MakeSymbol("shen.*signedfuncs*")
		reg35068 := PrimValue(reg35067)
		reg35069 := PrimCons(reg35066, reg35068)
		reg35070 := PrimSet(reg35057, reg35069)
		_ = reg35070
		reg35071 := MakeSymbol("shen.*signedfuncs*")
		reg35072 := MakeSymbol("string->n")
		reg35073 := MakeSymbol("string")
		reg35074 := MakeSymbol("-->")
		reg35075 := MakeSymbol("number")
		reg35076 := Nil
		reg35077 := PrimCons(reg35075, reg35076)
		reg35078 := PrimCons(reg35074, reg35077)
		reg35079 := PrimCons(reg35073, reg35078)
		reg35080 := PrimCons(reg35072, reg35079)
		reg35081 := MakeSymbol("shen.*signedfuncs*")
		reg35082 := PrimValue(reg35081)
		reg35083 := PrimCons(reg35080, reg35082)
		reg35084 := PrimSet(reg35071, reg35083)
		_ = reg35084
		reg35085 := MakeSymbol("shen.*signedfuncs*")
		reg35086 := MakeSymbol("string->symbol")
		reg35087 := MakeSymbol("string")
		reg35088 := MakeSymbol("-->")
		reg35089 := MakeSymbol("symbol")
		reg35090 := Nil
		reg35091 := PrimCons(reg35089, reg35090)
		reg35092 := PrimCons(reg35088, reg35091)
		reg35093 := PrimCons(reg35087, reg35092)
		reg35094 := PrimCons(reg35086, reg35093)
		reg35095 := MakeSymbol("shen.*signedfuncs*")
		reg35096 := PrimValue(reg35095)
		reg35097 := PrimCons(reg35094, reg35096)
		reg35098 := PrimSet(reg35085, reg35097)
		_ = reg35098
		reg35099 := MakeSymbol("shen.*signedfuncs*")
		reg35100 := MakeSymbol("sum")
		reg35101 := MakeSymbol("list")
		reg35102 := MakeSymbol("number")
		reg35103 := Nil
		reg35104 := PrimCons(reg35102, reg35103)
		reg35105 := PrimCons(reg35101, reg35104)
		reg35106 := MakeSymbol("-->")
		reg35107 := MakeSymbol("number")
		reg35108 := Nil
		reg35109 := PrimCons(reg35107, reg35108)
		reg35110 := PrimCons(reg35106, reg35109)
		reg35111 := PrimCons(reg35105, reg35110)
		reg35112 := PrimCons(reg35100, reg35111)
		reg35113 := MakeSymbol("shen.*signedfuncs*")
		reg35114 := PrimValue(reg35113)
		reg35115 := PrimCons(reg35112, reg35114)
		reg35116 := PrimSet(reg35099, reg35115)
		_ = reg35116
		reg35117 := MakeSymbol("shen.*signedfuncs*")
		reg35118 := MakeSymbol("symbol?")
		reg35119 := MakeSymbol("A")
		reg35120 := MakeSymbol("-->")
		reg35121 := MakeSymbol("boolean")
		reg35122 := Nil
		reg35123 := PrimCons(reg35121, reg35122)
		reg35124 := PrimCons(reg35120, reg35123)
		reg35125 := PrimCons(reg35119, reg35124)
		reg35126 := PrimCons(reg35118, reg35125)
		reg35127 := MakeSymbol("shen.*signedfuncs*")
		reg35128 := PrimValue(reg35127)
		reg35129 := PrimCons(reg35126, reg35128)
		reg35130 := PrimSet(reg35117, reg35129)
		_ = reg35130
		reg35131 := MakeSymbol("shen.*signedfuncs*")
		reg35132 := MakeSymbol("systemf")
		reg35133 := MakeSymbol("symbol")
		reg35134 := MakeSymbol("-->")
		reg35135 := MakeSymbol("symbol")
		reg35136 := Nil
		reg35137 := PrimCons(reg35135, reg35136)
		reg35138 := PrimCons(reg35134, reg35137)
		reg35139 := PrimCons(reg35133, reg35138)
		reg35140 := PrimCons(reg35132, reg35139)
		reg35141 := MakeSymbol("shen.*signedfuncs*")
		reg35142 := PrimValue(reg35141)
		reg35143 := PrimCons(reg35140, reg35142)
		reg35144 := PrimSet(reg35131, reg35143)
		_ = reg35144
		reg35145 := MakeSymbol("shen.*signedfuncs*")
		reg35146 := MakeSymbol("tail")
		reg35147 := MakeSymbol("list")
		reg35148 := MakeSymbol("A")
		reg35149 := Nil
		reg35150 := PrimCons(reg35148, reg35149)
		reg35151 := PrimCons(reg35147, reg35150)
		reg35152 := MakeSymbol("-->")
		reg35153 := MakeSymbol("list")
		reg35154 := MakeSymbol("A")
		reg35155 := Nil
		reg35156 := PrimCons(reg35154, reg35155)
		reg35157 := PrimCons(reg35153, reg35156)
		reg35158 := Nil
		reg35159 := PrimCons(reg35157, reg35158)
		reg35160 := PrimCons(reg35152, reg35159)
		reg35161 := PrimCons(reg35151, reg35160)
		reg35162 := PrimCons(reg35146, reg35161)
		reg35163 := MakeSymbol("shen.*signedfuncs*")
		reg35164 := PrimValue(reg35163)
		reg35165 := PrimCons(reg35162, reg35164)
		reg35166 := PrimSet(reg35145, reg35165)
		_ = reg35166
		reg35167 := MakeSymbol("shen.*signedfuncs*")
		reg35168 := MakeSymbol("tlstr")
		reg35169 := MakeSymbol("string")
		reg35170 := MakeSymbol("-->")
		reg35171 := MakeSymbol("string")
		reg35172 := Nil
		reg35173 := PrimCons(reg35171, reg35172)
		reg35174 := PrimCons(reg35170, reg35173)
		reg35175 := PrimCons(reg35169, reg35174)
		reg35176 := PrimCons(reg35168, reg35175)
		reg35177 := MakeSymbol("shen.*signedfuncs*")
		reg35178 := PrimValue(reg35177)
		reg35179 := PrimCons(reg35176, reg35178)
		reg35180 := PrimSet(reg35167, reg35179)
		_ = reg35180
		reg35181 := MakeSymbol("shen.*signedfuncs*")
		reg35182 := MakeSymbol("tlv")
		reg35183 := MakeSymbol("vector")
		reg35184 := MakeSymbol("A")
		reg35185 := Nil
		reg35186 := PrimCons(reg35184, reg35185)
		reg35187 := PrimCons(reg35183, reg35186)
		reg35188 := MakeSymbol("-->")
		reg35189 := MakeSymbol("vector")
		reg35190 := MakeSymbol("A")
		reg35191 := Nil
		reg35192 := PrimCons(reg35190, reg35191)
		reg35193 := PrimCons(reg35189, reg35192)
		reg35194 := Nil
		reg35195 := PrimCons(reg35193, reg35194)
		reg35196 := PrimCons(reg35188, reg35195)
		reg35197 := PrimCons(reg35187, reg35196)
		reg35198 := PrimCons(reg35182, reg35197)
		reg35199 := MakeSymbol("shen.*signedfuncs*")
		reg35200 := PrimValue(reg35199)
		reg35201 := PrimCons(reg35198, reg35200)
		reg35202 := PrimSet(reg35181, reg35201)
		_ = reg35202
		reg35203 := MakeSymbol("shen.*signedfuncs*")
		reg35204 := MakeSymbol("tc")
		reg35205 := MakeSymbol("symbol")
		reg35206 := MakeSymbol("-->")
		reg35207 := MakeSymbol("boolean")
		reg35208 := Nil
		reg35209 := PrimCons(reg35207, reg35208)
		reg35210 := PrimCons(reg35206, reg35209)
		reg35211 := PrimCons(reg35205, reg35210)
		reg35212 := PrimCons(reg35204, reg35211)
		reg35213 := MakeSymbol("shen.*signedfuncs*")
		reg35214 := PrimValue(reg35213)
		reg35215 := PrimCons(reg35212, reg35214)
		reg35216 := PrimSet(reg35203, reg35215)
		_ = reg35216
		reg35217 := MakeSymbol("shen.*signedfuncs*")
		reg35218 := MakeSymbol("tc?")
		reg35219 := MakeSymbol("-->")
		reg35220 := MakeSymbol("boolean")
		reg35221 := Nil
		reg35222 := PrimCons(reg35220, reg35221)
		reg35223 := PrimCons(reg35219, reg35222)
		reg35224 := PrimCons(reg35218, reg35223)
		reg35225 := MakeSymbol("shen.*signedfuncs*")
		reg35226 := PrimValue(reg35225)
		reg35227 := PrimCons(reg35224, reg35226)
		reg35228 := PrimSet(reg35217, reg35227)
		_ = reg35228
		reg35229 := MakeSymbol("shen.*signedfuncs*")
		reg35230 := MakeSymbol("thaw")
		reg35231 := MakeSymbol("lazy")
		reg35232 := MakeSymbol("A")
		reg35233 := Nil
		reg35234 := PrimCons(reg35232, reg35233)
		reg35235 := PrimCons(reg35231, reg35234)
		reg35236 := MakeSymbol("-->")
		reg35237 := MakeSymbol("A")
		reg35238 := Nil
		reg35239 := PrimCons(reg35237, reg35238)
		reg35240 := PrimCons(reg35236, reg35239)
		reg35241 := PrimCons(reg35235, reg35240)
		reg35242 := PrimCons(reg35230, reg35241)
		reg35243 := MakeSymbol("shen.*signedfuncs*")
		reg35244 := PrimValue(reg35243)
		reg35245 := PrimCons(reg35242, reg35244)
		reg35246 := PrimSet(reg35229, reg35245)
		_ = reg35246
		reg35247 := MakeSymbol("shen.*signedfuncs*")
		reg35248 := MakeSymbol("track")
		reg35249 := MakeSymbol("symbol")
		reg35250 := MakeSymbol("-->")
		reg35251 := MakeSymbol("symbol")
		reg35252 := Nil
		reg35253 := PrimCons(reg35251, reg35252)
		reg35254 := PrimCons(reg35250, reg35253)
		reg35255 := PrimCons(reg35249, reg35254)
		reg35256 := PrimCons(reg35248, reg35255)
		reg35257 := MakeSymbol("shen.*signedfuncs*")
		reg35258 := PrimValue(reg35257)
		reg35259 := PrimCons(reg35256, reg35258)
		reg35260 := PrimSet(reg35247, reg35259)
		_ = reg35260
		reg35261 := MakeSymbol("shen.*signedfuncs*")
		reg35262 := MakeSymbol("trap-error")
		reg35263 := MakeSymbol("A")
		reg35264 := MakeSymbol("-->")
		reg35265 := MakeSymbol("exception")
		reg35266 := MakeSymbol("-->")
		reg35267 := MakeSymbol("A")
		reg35268 := Nil
		reg35269 := PrimCons(reg35267, reg35268)
		reg35270 := PrimCons(reg35266, reg35269)
		reg35271 := PrimCons(reg35265, reg35270)
		reg35272 := MakeSymbol("-->")
		reg35273 := MakeSymbol("A")
		reg35274 := Nil
		reg35275 := PrimCons(reg35273, reg35274)
		reg35276 := PrimCons(reg35272, reg35275)
		reg35277 := PrimCons(reg35271, reg35276)
		reg35278 := Nil
		reg35279 := PrimCons(reg35277, reg35278)
		reg35280 := PrimCons(reg35264, reg35279)
		reg35281 := PrimCons(reg35263, reg35280)
		reg35282 := PrimCons(reg35262, reg35281)
		reg35283 := MakeSymbol("shen.*signedfuncs*")
		reg35284 := PrimValue(reg35283)
		reg35285 := PrimCons(reg35282, reg35284)
		reg35286 := PrimSet(reg35261, reg35285)
		_ = reg35286
		reg35287 := MakeSymbol("shen.*signedfuncs*")
		reg35288 := MakeSymbol("tuple?")
		reg35289 := MakeSymbol("A")
		reg35290 := MakeSymbol("-->")
		reg35291 := MakeSymbol("boolean")
		reg35292 := Nil
		reg35293 := PrimCons(reg35291, reg35292)
		reg35294 := PrimCons(reg35290, reg35293)
		reg35295 := PrimCons(reg35289, reg35294)
		reg35296 := PrimCons(reg35288, reg35295)
		reg35297 := MakeSymbol("shen.*signedfuncs*")
		reg35298 := PrimValue(reg35297)
		reg35299 := PrimCons(reg35296, reg35298)
		reg35300 := PrimSet(reg35287, reg35299)
		_ = reg35300
		reg35301 := MakeSymbol("shen.*signedfuncs*")
		reg35302 := MakeSymbol("undefmacro")
		reg35303 := MakeSymbol("symbol")
		reg35304 := MakeSymbol("-->")
		reg35305 := MakeSymbol("symbol")
		reg35306 := Nil
		reg35307 := PrimCons(reg35305, reg35306)
		reg35308 := PrimCons(reg35304, reg35307)
		reg35309 := PrimCons(reg35303, reg35308)
		reg35310 := PrimCons(reg35302, reg35309)
		reg35311 := MakeSymbol("shen.*signedfuncs*")
		reg35312 := PrimValue(reg35311)
		reg35313 := PrimCons(reg35310, reg35312)
		reg35314 := PrimSet(reg35301, reg35313)
		_ = reg35314
		reg35315 := MakeSymbol("shen.*signedfuncs*")
		reg35316 := MakeSymbol("union")
		reg35317 := MakeSymbol("list")
		reg35318 := MakeSymbol("A")
		reg35319 := Nil
		reg35320 := PrimCons(reg35318, reg35319)
		reg35321 := PrimCons(reg35317, reg35320)
		reg35322 := MakeSymbol("-->")
		reg35323 := MakeSymbol("list")
		reg35324 := MakeSymbol("A")
		reg35325 := Nil
		reg35326 := PrimCons(reg35324, reg35325)
		reg35327 := PrimCons(reg35323, reg35326)
		reg35328 := MakeSymbol("-->")
		reg35329 := MakeSymbol("list")
		reg35330 := MakeSymbol("A")
		reg35331 := Nil
		reg35332 := PrimCons(reg35330, reg35331)
		reg35333 := PrimCons(reg35329, reg35332)
		reg35334 := Nil
		reg35335 := PrimCons(reg35333, reg35334)
		reg35336 := PrimCons(reg35328, reg35335)
		reg35337 := PrimCons(reg35327, reg35336)
		reg35338 := Nil
		reg35339 := PrimCons(reg35337, reg35338)
		reg35340 := PrimCons(reg35322, reg35339)
		reg35341 := PrimCons(reg35321, reg35340)
		reg35342 := PrimCons(reg35316, reg35341)
		reg35343 := MakeSymbol("shen.*signedfuncs*")
		reg35344 := PrimValue(reg35343)
		reg35345 := PrimCons(reg35342, reg35344)
		reg35346 := PrimSet(reg35315, reg35345)
		_ = reg35346
		reg35347 := MakeSymbol("shen.*signedfuncs*")
		reg35348 := MakeSymbol("unprofile")
		reg35349 := MakeSymbol("A")
		reg35350 := MakeSymbol("-->")
		reg35351 := MakeSymbol("B")
		reg35352 := Nil
		reg35353 := PrimCons(reg35351, reg35352)
		reg35354 := PrimCons(reg35350, reg35353)
		reg35355 := PrimCons(reg35349, reg35354)
		reg35356 := MakeSymbol("-->")
		reg35357 := MakeSymbol("A")
		reg35358 := MakeSymbol("-->")
		reg35359 := MakeSymbol("B")
		reg35360 := Nil
		reg35361 := PrimCons(reg35359, reg35360)
		reg35362 := PrimCons(reg35358, reg35361)
		reg35363 := PrimCons(reg35357, reg35362)
		reg35364 := Nil
		reg35365 := PrimCons(reg35363, reg35364)
		reg35366 := PrimCons(reg35356, reg35365)
		reg35367 := PrimCons(reg35355, reg35366)
		reg35368 := PrimCons(reg35348, reg35367)
		reg35369 := MakeSymbol("shen.*signedfuncs*")
		reg35370 := PrimValue(reg35369)
		reg35371 := PrimCons(reg35368, reg35370)
		reg35372 := PrimSet(reg35347, reg35371)
		_ = reg35372
		reg35373 := MakeSymbol("shen.*signedfuncs*")
		reg35374 := MakeSymbol("untrack")
		reg35375 := MakeSymbol("symbol")
		reg35376 := MakeSymbol("-->")
		reg35377 := MakeSymbol("symbol")
		reg35378 := Nil
		reg35379 := PrimCons(reg35377, reg35378)
		reg35380 := PrimCons(reg35376, reg35379)
		reg35381 := PrimCons(reg35375, reg35380)
		reg35382 := PrimCons(reg35374, reg35381)
		reg35383 := MakeSymbol("shen.*signedfuncs*")
		reg35384 := PrimValue(reg35383)
		reg35385 := PrimCons(reg35382, reg35384)
		reg35386 := PrimSet(reg35373, reg35385)
		_ = reg35386
		reg35387 := MakeSymbol("shen.*signedfuncs*")
		reg35388 := MakeSymbol("unspecialise")
		reg35389 := MakeSymbol("symbol")
		reg35390 := MakeSymbol("-->")
		reg35391 := MakeSymbol("symbol")
		reg35392 := Nil
		reg35393 := PrimCons(reg35391, reg35392)
		reg35394 := PrimCons(reg35390, reg35393)
		reg35395 := PrimCons(reg35389, reg35394)
		reg35396 := PrimCons(reg35388, reg35395)
		reg35397 := MakeSymbol("shen.*signedfuncs*")
		reg35398 := PrimValue(reg35397)
		reg35399 := PrimCons(reg35396, reg35398)
		reg35400 := PrimSet(reg35387, reg35399)
		_ = reg35400
		reg35401 := MakeSymbol("shen.*signedfuncs*")
		reg35402 := MakeSymbol("variable?")
		reg35403 := MakeSymbol("A")
		reg35404 := MakeSymbol("-->")
		reg35405 := MakeSymbol("boolean")
		reg35406 := Nil
		reg35407 := PrimCons(reg35405, reg35406)
		reg35408 := PrimCons(reg35404, reg35407)
		reg35409 := PrimCons(reg35403, reg35408)
		reg35410 := PrimCons(reg35402, reg35409)
		reg35411 := MakeSymbol("shen.*signedfuncs*")
		reg35412 := PrimValue(reg35411)
		reg35413 := PrimCons(reg35410, reg35412)
		reg35414 := PrimSet(reg35401, reg35413)
		_ = reg35414
		reg35415 := MakeSymbol("shen.*signedfuncs*")
		reg35416 := MakeSymbol("vector?")
		reg35417 := MakeSymbol("A")
		reg35418 := MakeSymbol("-->")
		reg35419 := MakeSymbol("boolean")
		reg35420 := Nil
		reg35421 := PrimCons(reg35419, reg35420)
		reg35422 := PrimCons(reg35418, reg35421)
		reg35423 := PrimCons(reg35417, reg35422)
		reg35424 := PrimCons(reg35416, reg35423)
		reg35425 := MakeSymbol("shen.*signedfuncs*")
		reg35426 := PrimValue(reg35425)
		reg35427 := PrimCons(reg35424, reg35426)
		reg35428 := PrimSet(reg35415, reg35427)
		_ = reg35428
		reg35429 := MakeSymbol("shen.*signedfuncs*")
		reg35430 := MakeSymbol("version")
		reg35431 := MakeSymbol("-->")
		reg35432 := MakeSymbol("string")
		reg35433 := Nil
		reg35434 := PrimCons(reg35432, reg35433)
		reg35435 := PrimCons(reg35431, reg35434)
		reg35436 := PrimCons(reg35430, reg35435)
		reg35437 := MakeSymbol("shen.*signedfuncs*")
		reg35438 := PrimValue(reg35437)
		reg35439 := PrimCons(reg35436, reg35438)
		reg35440 := PrimSet(reg35429, reg35439)
		_ = reg35440
		reg35441 := MakeSymbol("shen.*signedfuncs*")
		reg35442 := MakeSymbol("write-to-file")
		reg35443 := MakeSymbol("string")
		reg35444 := MakeSymbol("-->")
		reg35445 := MakeSymbol("A")
		reg35446 := MakeSymbol("-->")
		reg35447 := MakeSymbol("A")
		reg35448 := Nil
		reg35449 := PrimCons(reg35447, reg35448)
		reg35450 := PrimCons(reg35446, reg35449)
		reg35451 := PrimCons(reg35445, reg35450)
		reg35452 := Nil
		reg35453 := PrimCons(reg35451, reg35452)
		reg35454 := PrimCons(reg35444, reg35453)
		reg35455 := PrimCons(reg35443, reg35454)
		reg35456 := PrimCons(reg35442, reg35455)
		reg35457 := MakeSymbol("shen.*signedfuncs*")
		reg35458 := PrimValue(reg35457)
		reg35459 := PrimCons(reg35456, reg35458)
		reg35460 := PrimSet(reg35441, reg35459)
		_ = reg35460
		reg35461 := MakeSymbol("shen.*signedfuncs*")
		reg35462 := MakeSymbol("write-byte")
		reg35463 := MakeSymbol("number")
		reg35464 := MakeSymbol("-->")
		reg35465 := MakeSymbol("stream")
		reg35466 := MakeSymbol("out")
		reg35467 := Nil
		reg35468 := PrimCons(reg35466, reg35467)
		reg35469 := PrimCons(reg35465, reg35468)
		reg35470 := MakeSymbol("-->")
		reg35471 := MakeSymbol("number")
		reg35472 := Nil
		reg35473 := PrimCons(reg35471, reg35472)
		reg35474 := PrimCons(reg35470, reg35473)
		reg35475 := PrimCons(reg35469, reg35474)
		reg35476 := Nil
		reg35477 := PrimCons(reg35475, reg35476)
		reg35478 := PrimCons(reg35464, reg35477)
		reg35479 := PrimCons(reg35463, reg35478)
		reg35480 := PrimCons(reg35462, reg35479)
		reg35481 := MakeSymbol("shen.*signedfuncs*")
		reg35482 := PrimValue(reg35481)
		reg35483 := PrimCons(reg35480, reg35482)
		reg35484 := PrimSet(reg35461, reg35483)
		_ = reg35484
		reg35485 := MakeSymbol("shen.*signedfuncs*")
		reg35486 := MakeSymbol("y-or-n?")
		reg35487 := MakeSymbol("string")
		reg35488 := MakeSymbol("-->")
		reg35489 := MakeSymbol("boolean")
		reg35490 := Nil
		reg35491 := PrimCons(reg35489, reg35490)
		reg35492 := PrimCons(reg35488, reg35491)
		reg35493 := PrimCons(reg35487, reg35492)
		reg35494 := PrimCons(reg35486, reg35493)
		reg35495 := MakeSymbol("shen.*signedfuncs*")
		reg35496 := PrimValue(reg35495)
		reg35497 := PrimCons(reg35494, reg35496)
		reg35498 := PrimSet(reg35485, reg35497)
		_ = reg35498
		reg35499 := MakeSymbol("shen.*signedfuncs*")
		reg35500 := MakeSymbol(">")
		reg35501 := MakeSymbol("number")
		reg35502 := MakeSymbol("-->")
		reg35503 := MakeSymbol("number")
		reg35504 := MakeSymbol("-->")
		reg35505 := MakeSymbol("boolean")
		reg35506 := Nil
		reg35507 := PrimCons(reg35505, reg35506)
		reg35508 := PrimCons(reg35504, reg35507)
		reg35509 := PrimCons(reg35503, reg35508)
		reg35510 := Nil
		reg35511 := PrimCons(reg35509, reg35510)
		reg35512 := PrimCons(reg35502, reg35511)
		reg35513 := PrimCons(reg35501, reg35512)
		reg35514 := PrimCons(reg35500, reg35513)
		reg35515 := MakeSymbol("shen.*signedfuncs*")
		reg35516 := PrimValue(reg35515)
		reg35517 := PrimCons(reg35514, reg35516)
		reg35518 := PrimSet(reg35499, reg35517)
		_ = reg35518
		reg35519 := MakeSymbol("shen.*signedfuncs*")
		reg35520 := MakeSymbol("<")
		reg35521 := MakeSymbol("number")
		reg35522 := MakeSymbol("-->")
		reg35523 := MakeSymbol("number")
		reg35524 := MakeSymbol("-->")
		reg35525 := MakeSymbol("boolean")
		reg35526 := Nil
		reg35527 := PrimCons(reg35525, reg35526)
		reg35528 := PrimCons(reg35524, reg35527)
		reg35529 := PrimCons(reg35523, reg35528)
		reg35530 := Nil
		reg35531 := PrimCons(reg35529, reg35530)
		reg35532 := PrimCons(reg35522, reg35531)
		reg35533 := PrimCons(reg35521, reg35532)
		reg35534 := PrimCons(reg35520, reg35533)
		reg35535 := MakeSymbol("shen.*signedfuncs*")
		reg35536 := PrimValue(reg35535)
		reg35537 := PrimCons(reg35534, reg35536)
		reg35538 := PrimSet(reg35519, reg35537)
		_ = reg35538
		reg35539 := MakeSymbol("shen.*signedfuncs*")
		reg35540 := MakeSymbol(">=")
		reg35541 := MakeSymbol("number")
		reg35542 := MakeSymbol("-->")
		reg35543 := MakeSymbol("number")
		reg35544 := MakeSymbol("-->")
		reg35545 := MakeSymbol("boolean")
		reg35546 := Nil
		reg35547 := PrimCons(reg35545, reg35546)
		reg35548 := PrimCons(reg35544, reg35547)
		reg35549 := PrimCons(reg35543, reg35548)
		reg35550 := Nil
		reg35551 := PrimCons(reg35549, reg35550)
		reg35552 := PrimCons(reg35542, reg35551)
		reg35553 := PrimCons(reg35541, reg35552)
		reg35554 := PrimCons(reg35540, reg35553)
		reg35555 := MakeSymbol("shen.*signedfuncs*")
		reg35556 := PrimValue(reg35555)
		reg35557 := PrimCons(reg35554, reg35556)
		reg35558 := PrimSet(reg35539, reg35557)
		_ = reg35558
		reg35559 := MakeSymbol("shen.*signedfuncs*")
		reg35560 := MakeSymbol("<=")
		reg35561 := MakeSymbol("number")
		reg35562 := MakeSymbol("-->")
		reg35563 := MakeSymbol("number")
		reg35564 := MakeSymbol("-->")
		reg35565 := MakeSymbol("boolean")
		reg35566 := Nil
		reg35567 := PrimCons(reg35565, reg35566)
		reg35568 := PrimCons(reg35564, reg35567)
		reg35569 := PrimCons(reg35563, reg35568)
		reg35570 := Nil
		reg35571 := PrimCons(reg35569, reg35570)
		reg35572 := PrimCons(reg35562, reg35571)
		reg35573 := PrimCons(reg35561, reg35572)
		reg35574 := PrimCons(reg35560, reg35573)
		reg35575 := MakeSymbol("shen.*signedfuncs*")
		reg35576 := PrimValue(reg35575)
		reg35577 := PrimCons(reg35574, reg35576)
		reg35578 := PrimSet(reg35559, reg35577)
		_ = reg35578
		reg35579 := MakeSymbol("shen.*signedfuncs*")
		reg35580 := MakeSymbol("=")
		reg35581 := MakeSymbol("A")
		reg35582 := MakeSymbol("-->")
		reg35583 := MakeSymbol("A")
		reg35584 := MakeSymbol("-->")
		reg35585 := MakeSymbol("boolean")
		reg35586 := Nil
		reg35587 := PrimCons(reg35585, reg35586)
		reg35588 := PrimCons(reg35584, reg35587)
		reg35589 := PrimCons(reg35583, reg35588)
		reg35590 := Nil
		reg35591 := PrimCons(reg35589, reg35590)
		reg35592 := PrimCons(reg35582, reg35591)
		reg35593 := PrimCons(reg35581, reg35592)
		reg35594 := PrimCons(reg35580, reg35593)
		reg35595 := MakeSymbol("shen.*signedfuncs*")
		reg35596 := PrimValue(reg35595)
		reg35597 := PrimCons(reg35594, reg35596)
		reg35598 := PrimSet(reg35579, reg35597)
		_ = reg35598
		reg35599 := MakeSymbol("shen.*signedfuncs*")
		reg35600 := MakeSymbol("+")
		reg35601 := MakeSymbol("number")
		reg35602 := MakeSymbol("-->")
		reg35603 := MakeSymbol("number")
		reg35604 := MakeSymbol("-->")
		reg35605 := MakeSymbol("number")
		reg35606 := Nil
		reg35607 := PrimCons(reg35605, reg35606)
		reg35608 := PrimCons(reg35604, reg35607)
		reg35609 := PrimCons(reg35603, reg35608)
		reg35610 := Nil
		reg35611 := PrimCons(reg35609, reg35610)
		reg35612 := PrimCons(reg35602, reg35611)
		reg35613 := PrimCons(reg35601, reg35612)
		reg35614 := PrimCons(reg35600, reg35613)
		reg35615 := MakeSymbol("shen.*signedfuncs*")
		reg35616 := PrimValue(reg35615)
		reg35617 := PrimCons(reg35614, reg35616)
		reg35618 := PrimSet(reg35599, reg35617)
		_ = reg35618
		reg35619 := MakeSymbol("shen.*signedfuncs*")
		reg35620 := MakeSymbol("/")
		reg35621 := MakeSymbol("number")
		reg35622 := MakeSymbol("-->")
		reg35623 := MakeSymbol("number")
		reg35624 := MakeSymbol("-->")
		reg35625 := MakeSymbol("number")
		reg35626 := Nil
		reg35627 := PrimCons(reg35625, reg35626)
		reg35628 := PrimCons(reg35624, reg35627)
		reg35629 := PrimCons(reg35623, reg35628)
		reg35630 := Nil
		reg35631 := PrimCons(reg35629, reg35630)
		reg35632 := PrimCons(reg35622, reg35631)
		reg35633 := PrimCons(reg35621, reg35632)
		reg35634 := PrimCons(reg35620, reg35633)
		reg35635 := MakeSymbol("shen.*signedfuncs*")
		reg35636 := PrimValue(reg35635)
		reg35637 := PrimCons(reg35634, reg35636)
		reg35638 := PrimSet(reg35619, reg35637)
		_ = reg35638
		reg35639 := MakeSymbol("shen.*signedfuncs*")
		reg35640 := MakeSymbol("-")
		reg35641 := MakeSymbol("number")
		reg35642 := MakeSymbol("-->")
		reg35643 := MakeSymbol("number")
		reg35644 := MakeSymbol("-->")
		reg35645 := MakeSymbol("number")
		reg35646 := Nil
		reg35647 := PrimCons(reg35645, reg35646)
		reg35648 := PrimCons(reg35644, reg35647)
		reg35649 := PrimCons(reg35643, reg35648)
		reg35650 := Nil
		reg35651 := PrimCons(reg35649, reg35650)
		reg35652 := PrimCons(reg35642, reg35651)
		reg35653 := PrimCons(reg35641, reg35652)
		reg35654 := PrimCons(reg35640, reg35653)
		reg35655 := MakeSymbol("shen.*signedfuncs*")
		reg35656 := PrimValue(reg35655)
		reg35657 := PrimCons(reg35654, reg35656)
		reg35658 := PrimSet(reg35639, reg35657)
		_ = reg35658
		reg35659 := MakeSymbol("shen.*signedfuncs*")
		reg35660 := MakeSymbol("*")
		reg35661 := MakeSymbol("number")
		reg35662 := MakeSymbol("-->")
		reg35663 := MakeSymbol("number")
		reg35664 := MakeSymbol("-->")
		reg35665 := MakeSymbol("number")
		reg35666 := Nil
		reg35667 := PrimCons(reg35665, reg35666)
		reg35668 := PrimCons(reg35664, reg35667)
		reg35669 := PrimCons(reg35663, reg35668)
		reg35670 := Nil
		reg35671 := PrimCons(reg35669, reg35670)
		reg35672 := PrimCons(reg35662, reg35671)
		reg35673 := PrimCons(reg35661, reg35672)
		reg35674 := PrimCons(reg35660, reg35673)
		reg35675 := MakeSymbol("shen.*signedfuncs*")
		reg35676 := PrimValue(reg35675)
		reg35677 := PrimCons(reg35674, reg35676)
		reg35678 := PrimSet(reg35659, reg35677)
		_ = reg35678
		reg35679 := MakeSymbol("shen.*signedfuncs*")
		reg35680 := MakeSymbol("==")
		reg35681 := MakeSymbol("A")
		reg35682 := MakeSymbol("-->")
		reg35683 := MakeSymbol("B")
		reg35684 := MakeSymbol("-->")
		reg35685 := MakeSymbol("boolean")
		reg35686 := Nil
		reg35687 := PrimCons(reg35685, reg35686)
		reg35688 := PrimCons(reg35684, reg35687)
		reg35689 := PrimCons(reg35683, reg35688)
		reg35690 := Nil
		reg35691 := PrimCons(reg35689, reg35690)
		reg35692 := PrimCons(reg35682, reg35691)
		reg35693 := PrimCons(reg35681, reg35692)
		reg35694 := PrimCons(reg35680, reg35693)
		reg35695 := MakeSymbol("shen.*signedfuncs*")
		reg35696 := PrimValue(reg35695)
		reg35697 := PrimCons(reg35694, reg35696)
		reg35698 := PrimSet(reg35679, reg35697)
		__e.Return(reg35698)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.initialise-signedfuncs", value: __defun__shen_4initialise_1signedfuncs})

	__defun__shen_4initialise_1signedfunc_1lambda_1forms = MakeNative(func(__e Evaluator, __args ...Obj) {
		reg35699 := MakeSymbol("shen.type-signature-of-absvector?")
		reg35700 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4559 := __args[0]
			_ = V4559
			reg35701 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4560 := __args[0]
				_ = V4560
				reg35702 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4561 := __args[0]
					_ = V4561
					__e.TailApply(__defun__shen_4type_1signature_1of_1absvector_2, V4559, V4560, V4561)
					return
				}, 1)
				__e.Return(reg35702)
				return
			}, 1)
			__e.Return(reg35701)
			return
		}, 1)
		reg35704 := PrimCons(reg35699, reg35700)
		reg35705 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35704)
		_ = reg35705
		reg35706 := MakeSymbol("shen.type-signature-of-adjoin")
		reg35707 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4549 := __args[0]
			_ = V4549
			reg35708 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4550 := __args[0]
				_ = V4550
				reg35709 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4551 := __args[0]
					_ = V4551
					__e.TailApply(__defun__shen_4type_1signature_1of_1adjoin, V4549, V4550, V4551)
					return
				}, 1)
				__e.Return(reg35709)
				return
			}, 1)
			__e.Return(reg35708)
			return
		}, 1)
		reg35711 := PrimCons(reg35706, reg35707)
		reg35712 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35711)
		_ = reg35712
		reg35713 := MakeSymbol("shen.type-signature-of-and")
		reg35714 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4539 := __args[0]
			_ = V4539
			reg35715 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4540 := __args[0]
				_ = V4540
				reg35716 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4541 := __args[0]
					_ = V4541
					__e.TailApply(__defun__shen_4type_1signature_1of_1and, V4539, V4540, V4541)
					return
				}, 1)
				__e.Return(reg35716)
				return
			}, 1)
			__e.Return(reg35715)
			return
		}, 1)
		reg35718 := PrimCons(reg35713, reg35714)
		reg35719 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35718)
		_ = reg35719
		reg35720 := MakeSymbol("shen.type-signature-of-shen.app")
		reg35721 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4529 := __args[0]
			_ = V4529
			reg35722 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4530 := __args[0]
				_ = V4530
				reg35723 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4531 := __args[0]
					_ = V4531
					__e.TailApply(__defun__shen_4type_1signature_1of_1shen_4app, V4529, V4530, V4531)
					return
				}, 1)
				__e.Return(reg35723)
				return
			}, 1)
			__e.Return(reg35722)
			return
		}, 1)
		reg35725 := PrimCons(reg35720, reg35721)
		reg35726 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35725)
		_ = reg35726
		reg35727 := MakeSymbol("shen.type-signature-of-append")
		reg35728 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4519 := __args[0]
			_ = V4519
			reg35729 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4520 := __args[0]
				_ = V4520
				reg35730 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4521 := __args[0]
					_ = V4521
					__e.TailApply(__defun__shen_4type_1signature_1of_1append, V4519, V4520, V4521)
					return
				}, 1)
				__e.Return(reg35730)
				return
			}, 1)
			__e.Return(reg35729)
			return
		}, 1)
		reg35732 := PrimCons(reg35727, reg35728)
		reg35733 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35732)
		_ = reg35733
		reg35734 := MakeSymbol("shen.type-signature-of-arity")
		reg35735 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4509 := __args[0]
			_ = V4509
			reg35736 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4510 := __args[0]
				_ = V4510
				reg35737 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4511 := __args[0]
					_ = V4511
					__e.TailApply(__defun__shen_4type_1signature_1of_1arity, V4509, V4510, V4511)
					return
				}, 1)
				__e.Return(reg35737)
				return
			}, 1)
			__e.Return(reg35736)
			return
		}, 1)
		reg35739 := PrimCons(reg35734, reg35735)
		reg35740 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35739)
		_ = reg35740
		reg35741 := MakeSymbol("shen.type-signature-of-assoc")
		reg35742 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4499 := __args[0]
			_ = V4499
			reg35743 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4500 := __args[0]
				_ = V4500
				reg35744 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4501 := __args[0]
					_ = V4501
					__e.TailApply(__defun__shen_4type_1signature_1of_1assoc, V4499, V4500, V4501)
					return
				}, 1)
				__e.Return(reg35744)
				return
			}, 1)
			__e.Return(reg35743)
			return
		}, 1)
		reg35746 := PrimCons(reg35741, reg35742)
		reg35747 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35746)
		_ = reg35747
		reg35748 := MakeSymbol("shen.type-signature-of-boolean?")
		reg35749 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4489 := __args[0]
			_ = V4489
			reg35750 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4490 := __args[0]
				_ = V4490
				reg35751 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4491 := __args[0]
					_ = V4491
					__e.TailApply(__defun__shen_4type_1signature_1of_1boolean_2, V4489, V4490, V4491)
					return
				}, 1)
				__e.Return(reg35751)
				return
			}, 1)
			__e.Return(reg35750)
			return
		}, 1)
		reg35753 := PrimCons(reg35748, reg35749)
		reg35754 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35753)
		_ = reg35754
		reg35755 := MakeSymbol("shen.type-signature-of-bound?")
		reg35756 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4479 := __args[0]
			_ = V4479
			reg35757 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4480 := __args[0]
				_ = V4480
				reg35758 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4481 := __args[0]
					_ = V4481
					__e.TailApply(__defun__shen_4type_1signature_1of_1bound_2, V4479, V4480, V4481)
					return
				}, 1)
				__e.Return(reg35758)
				return
			}, 1)
			__e.Return(reg35757)
			return
		}, 1)
		reg35760 := PrimCons(reg35755, reg35756)
		reg35761 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35760)
		_ = reg35761
		reg35762 := MakeSymbol("shen.type-signature-of-cd")
		reg35763 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4469 := __args[0]
			_ = V4469
			reg35764 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4470 := __args[0]
				_ = V4470
				reg35765 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4471 := __args[0]
					_ = V4471
					__e.TailApply(__defun__shen_4type_1signature_1of_1cd, V4469, V4470, V4471)
					return
				}, 1)
				__e.Return(reg35765)
				return
			}, 1)
			__e.Return(reg35764)
			return
		}, 1)
		reg35767 := PrimCons(reg35762, reg35763)
		reg35768 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35767)
		_ = reg35768
		reg35769 := MakeSymbol("shen.type-signature-of-close")
		reg35770 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4459 := __args[0]
			_ = V4459
			reg35771 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4460 := __args[0]
				_ = V4460
				reg35772 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4461 := __args[0]
					_ = V4461
					__e.TailApply(__defun__shen_4type_1signature_1of_1close, V4459, V4460, V4461)
					return
				}, 1)
				__e.Return(reg35772)
				return
			}, 1)
			__e.Return(reg35771)
			return
		}, 1)
		reg35774 := PrimCons(reg35769, reg35770)
		reg35775 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35774)
		_ = reg35775
		reg35776 := MakeSymbol("shen.type-signature-of-cn")
		reg35777 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4449 := __args[0]
			_ = V4449
			reg35778 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4450 := __args[0]
				_ = V4450
				reg35779 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4451 := __args[0]
					_ = V4451
					__e.TailApply(__defun__shen_4type_1signature_1of_1cn, V4449, V4450, V4451)
					return
				}, 1)
				__e.Return(reg35779)
				return
			}, 1)
			__e.Return(reg35778)
			return
		}, 1)
		reg35781 := PrimCons(reg35776, reg35777)
		reg35782 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35781)
		_ = reg35782
		reg35783 := MakeSymbol("shen.type-signature-of-compile")
		reg35784 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4439 := __args[0]
			_ = V4439
			reg35785 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4440 := __args[0]
				_ = V4440
				reg35786 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4441 := __args[0]
					_ = V4441
					__e.TailApply(__defun__shen_4type_1signature_1of_1compile, V4439, V4440, V4441)
					return
				}, 1)
				__e.Return(reg35786)
				return
			}, 1)
			__e.Return(reg35785)
			return
		}, 1)
		reg35788 := PrimCons(reg35783, reg35784)
		reg35789 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35788)
		_ = reg35789
		reg35790 := MakeSymbol("shen.type-signature-of-cons?")
		reg35791 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4429 := __args[0]
			_ = V4429
			reg35792 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4430 := __args[0]
				_ = V4430
				reg35793 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4431 := __args[0]
					_ = V4431
					__e.TailApply(__defun__shen_4type_1signature_1of_1cons_2, V4429, V4430, V4431)
					return
				}, 1)
				__e.Return(reg35793)
				return
			}, 1)
			__e.Return(reg35792)
			return
		}, 1)
		reg35795 := PrimCons(reg35790, reg35791)
		reg35796 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35795)
		_ = reg35796
		reg35797 := MakeSymbol("shen.type-signature-of-destroy")
		reg35798 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4419 := __args[0]
			_ = V4419
			reg35799 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4420 := __args[0]
				_ = V4420
				reg35800 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4421 := __args[0]
					_ = V4421
					__e.TailApply(__defun__shen_4type_1signature_1of_1destroy, V4419, V4420, V4421)
					return
				}, 1)
				__e.Return(reg35800)
				return
			}, 1)
			__e.Return(reg35799)
			return
		}, 1)
		reg35802 := PrimCons(reg35797, reg35798)
		reg35803 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35802)
		_ = reg35803
		reg35804 := MakeSymbol("shen.type-signature-of-difference")
		reg35805 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4409 := __args[0]
			_ = V4409
			reg35806 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4410 := __args[0]
				_ = V4410
				reg35807 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4411 := __args[0]
					_ = V4411
					__e.TailApply(__defun__shen_4type_1signature_1of_1difference, V4409, V4410, V4411)
					return
				}, 1)
				__e.Return(reg35807)
				return
			}, 1)
			__e.Return(reg35806)
			return
		}, 1)
		reg35809 := PrimCons(reg35804, reg35805)
		reg35810 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35809)
		_ = reg35810
		reg35811 := MakeSymbol("shen.type-signature-of-do")
		reg35812 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4399 := __args[0]
			_ = V4399
			reg35813 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4400 := __args[0]
				_ = V4400
				reg35814 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4401 := __args[0]
					_ = V4401
					__e.TailApply(__defun__shen_4type_1signature_1of_1do, V4399, V4400, V4401)
					return
				}, 1)
				__e.Return(reg35814)
				return
			}, 1)
			__e.Return(reg35813)
			return
		}, 1)
		reg35816 := PrimCons(reg35811, reg35812)
		reg35817 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35816)
		_ = reg35817
		reg35818 := MakeSymbol("shen.type-signature-of-<e>")
		reg35819 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4389 := __args[0]
			_ = V4389
			reg35820 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4390 := __args[0]
				_ = V4390
				reg35821 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4391 := __args[0]
					_ = V4391
					__e.TailApply(__defun__shen_4type_1signature_1of_1_5e_6, V4389, V4390, V4391)
					return
				}, 1)
				__e.Return(reg35821)
				return
			}, 1)
			__e.Return(reg35820)
			return
		}, 1)
		reg35823 := PrimCons(reg35818, reg35819)
		reg35824 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35823)
		_ = reg35824
		reg35825 := MakeSymbol("shen.type-signature-of-<!>")
		reg35826 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4379 := __args[0]
			_ = V4379
			reg35827 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4380 := __args[0]
				_ = V4380
				reg35828 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4381 := __args[0]
					_ = V4381
					__e.TailApply(__defun__shen_4type_1signature_1of_1_5_b_6, V4379, V4380, V4381)
					return
				}, 1)
				__e.Return(reg35828)
				return
			}, 1)
			__e.Return(reg35827)
			return
		}, 1)
		reg35830 := PrimCons(reg35825, reg35826)
		reg35831 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35830)
		_ = reg35831
		reg35832 := MakeSymbol("shen.type-signature-of-element?")
		reg35833 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4369 := __args[0]
			_ = V4369
			reg35834 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4370 := __args[0]
				_ = V4370
				reg35835 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4371 := __args[0]
					_ = V4371
					__e.TailApply(__defun__shen_4type_1signature_1of_1element_2, V4369, V4370, V4371)
					return
				}, 1)
				__e.Return(reg35835)
				return
			}, 1)
			__e.Return(reg35834)
			return
		}, 1)
		reg35837 := PrimCons(reg35832, reg35833)
		reg35838 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35837)
		_ = reg35838
		reg35839 := MakeSymbol("shen.type-signature-of-empty?")
		reg35840 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4359 := __args[0]
			_ = V4359
			reg35841 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4360 := __args[0]
				_ = V4360
				reg35842 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4361 := __args[0]
					_ = V4361
					__e.TailApply(__defun__shen_4type_1signature_1of_1empty_2, V4359, V4360, V4361)
					return
				}, 1)
				__e.Return(reg35842)
				return
			}, 1)
			__e.Return(reg35841)
			return
		}, 1)
		reg35844 := PrimCons(reg35839, reg35840)
		reg35845 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35844)
		_ = reg35845
		reg35846 := MakeSymbol("shen.type-signature-of-enable-type-theory")
		reg35847 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4349 := __args[0]
			_ = V4349
			reg35848 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4350 := __args[0]
				_ = V4350
				reg35849 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4351 := __args[0]
					_ = V4351
					__e.TailApply(__defun__shen_4type_1signature_1of_1enable_1type_1theory, V4349, V4350, V4351)
					return
				}, 1)
				__e.Return(reg35849)
				return
			}, 1)
			__e.Return(reg35848)
			return
		}, 1)
		reg35851 := PrimCons(reg35846, reg35847)
		reg35852 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35851)
		_ = reg35852
		reg35853 := MakeSymbol("shen.type-signature-of-external")
		reg35854 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4339 := __args[0]
			_ = V4339
			reg35855 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4340 := __args[0]
				_ = V4340
				reg35856 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4341 := __args[0]
					_ = V4341
					__e.TailApply(__defun__shen_4type_1signature_1of_1external, V4339, V4340, V4341)
					return
				}, 1)
				__e.Return(reg35856)
				return
			}, 1)
			__e.Return(reg35855)
			return
		}, 1)
		reg35858 := PrimCons(reg35853, reg35854)
		reg35859 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35858)
		_ = reg35859
		reg35860 := MakeSymbol("shen.type-signature-of-error-to-string")
		reg35861 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4329 := __args[0]
			_ = V4329
			reg35862 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4330 := __args[0]
				_ = V4330
				reg35863 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4331 := __args[0]
					_ = V4331
					__e.TailApply(__defun__shen_4type_1signature_1of_1error_1to_1string, V4329, V4330, V4331)
					return
				}, 1)
				__e.Return(reg35863)
				return
			}, 1)
			__e.Return(reg35862)
			return
		}, 1)
		reg35865 := PrimCons(reg35860, reg35861)
		reg35866 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35865)
		_ = reg35866
		reg35867 := MakeSymbol("shen.type-signature-of-explode")
		reg35868 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4319 := __args[0]
			_ = V4319
			reg35869 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4320 := __args[0]
				_ = V4320
				reg35870 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4321 := __args[0]
					_ = V4321
					__e.TailApply(__defun__shen_4type_1signature_1of_1explode, V4319, V4320, V4321)
					return
				}, 1)
				__e.Return(reg35870)
				return
			}, 1)
			__e.Return(reg35869)
			return
		}, 1)
		reg35872 := PrimCons(reg35867, reg35868)
		reg35873 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35872)
		_ = reg35873
		reg35874 := MakeSymbol("shen.type-signature-of-fail")
		reg35875 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4309 := __args[0]
			_ = V4309
			reg35876 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4310 := __args[0]
				_ = V4310
				reg35877 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4311 := __args[0]
					_ = V4311
					__e.TailApply(__defun__shen_4type_1signature_1of_1fail, V4309, V4310, V4311)
					return
				}, 1)
				__e.Return(reg35877)
				return
			}, 1)
			__e.Return(reg35876)
			return
		}, 1)
		reg35879 := PrimCons(reg35874, reg35875)
		reg35880 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35879)
		_ = reg35880
		reg35881 := MakeSymbol("shen.type-signature-of-fail-if")
		reg35882 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4299 := __args[0]
			_ = V4299
			reg35883 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4300 := __args[0]
				_ = V4300
				reg35884 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4301 := __args[0]
					_ = V4301
					__e.TailApply(__defun__shen_4type_1signature_1of_1fail_1if, V4299, V4300, V4301)
					return
				}, 1)
				__e.Return(reg35884)
				return
			}, 1)
			__e.Return(reg35883)
			return
		}, 1)
		reg35886 := PrimCons(reg35881, reg35882)
		reg35887 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35886)
		_ = reg35887
		reg35888 := MakeSymbol("shen.type-signature-of-fix")
		reg35889 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4289 := __args[0]
			_ = V4289
			reg35890 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4290 := __args[0]
				_ = V4290
				reg35891 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4291 := __args[0]
					_ = V4291
					__e.TailApply(__defun__shen_4type_1signature_1of_1fix, V4289, V4290, V4291)
					return
				}, 1)
				__e.Return(reg35891)
				return
			}, 1)
			__e.Return(reg35890)
			return
		}, 1)
		reg35893 := PrimCons(reg35888, reg35889)
		reg35894 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35893)
		_ = reg35894
		reg35895 := MakeSymbol("shen.type-signature-of-freeze")
		reg35896 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4279 := __args[0]
			_ = V4279
			reg35897 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4280 := __args[0]
				_ = V4280
				reg35898 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4281 := __args[0]
					_ = V4281
					__e.TailApply(__defun__shen_4type_1signature_1of_1freeze, V4279, V4280, V4281)
					return
				}, 1)
				__e.Return(reg35898)
				return
			}, 1)
			__e.Return(reg35897)
			return
		}, 1)
		reg35900 := PrimCons(reg35895, reg35896)
		reg35901 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35900)
		_ = reg35901
		reg35902 := MakeSymbol("shen.type-signature-of-fst")
		reg35903 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4269 := __args[0]
			_ = V4269
			reg35904 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4270 := __args[0]
				_ = V4270
				reg35905 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4271 := __args[0]
					_ = V4271
					__e.TailApply(__defun__shen_4type_1signature_1of_1fst, V4269, V4270, V4271)
					return
				}, 1)
				__e.Return(reg35905)
				return
			}, 1)
			__e.Return(reg35904)
			return
		}, 1)
		reg35907 := PrimCons(reg35902, reg35903)
		reg35908 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35907)
		_ = reg35908
		reg35909 := MakeSymbol("shen.type-signature-of-function")
		reg35910 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4259 := __args[0]
			_ = V4259
			reg35911 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4260 := __args[0]
				_ = V4260
				reg35912 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4261 := __args[0]
					_ = V4261
					__e.TailApply(__defun__shen_4type_1signature_1of_1function, V4259, V4260, V4261)
					return
				}, 1)
				__e.Return(reg35912)
				return
			}, 1)
			__e.Return(reg35911)
			return
		}, 1)
		reg35914 := PrimCons(reg35909, reg35910)
		reg35915 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35914)
		_ = reg35915
		reg35916 := MakeSymbol("shen.type-signature-of-gensym")
		reg35917 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4249 := __args[0]
			_ = V4249
			reg35918 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4250 := __args[0]
				_ = V4250
				reg35919 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4251 := __args[0]
					_ = V4251
					__e.TailApply(__defun__shen_4type_1signature_1of_1gensym, V4249, V4250, V4251)
					return
				}, 1)
				__e.Return(reg35919)
				return
			}, 1)
			__e.Return(reg35918)
			return
		}, 1)
		reg35921 := PrimCons(reg35916, reg35917)
		reg35922 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35921)
		_ = reg35922
		reg35923 := MakeSymbol("shen.type-signature-of-<-vector")
		reg35924 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4239 := __args[0]
			_ = V4239
			reg35925 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4240 := __args[0]
				_ = V4240
				reg35926 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4241 := __args[0]
					_ = V4241
					__e.TailApply(__defun__shen_4type_1signature_1of_1_5_1vector, V4239, V4240, V4241)
					return
				}, 1)
				__e.Return(reg35926)
				return
			}, 1)
			__e.Return(reg35925)
			return
		}, 1)
		reg35928 := PrimCons(reg35923, reg35924)
		reg35929 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35928)
		_ = reg35929
		reg35930 := MakeSymbol("shen.type-signature-of-vector->")
		reg35931 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4229 := __args[0]
			_ = V4229
			reg35932 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4230 := __args[0]
				_ = V4230
				reg35933 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4231 := __args[0]
					_ = V4231
					__e.TailApply(__defun__shen_4type_1signature_1of_1vector_1_6, V4229, V4230, V4231)
					return
				}, 1)
				__e.Return(reg35933)
				return
			}, 1)
			__e.Return(reg35932)
			return
		}, 1)
		reg35935 := PrimCons(reg35930, reg35931)
		reg35936 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35935)
		_ = reg35936
		reg35937 := MakeSymbol("shen.type-signature-of-vector")
		reg35938 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4219 := __args[0]
			_ = V4219
			reg35939 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4220 := __args[0]
				_ = V4220
				reg35940 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4221 := __args[0]
					_ = V4221
					__e.TailApply(__defun__shen_4type_1signature_1of_1vector, V4219, V4220, V4221)
					return
				}, 1)
				__e.Return(reg35940)
				return
			}, 1)
			__e.Return(reg35939)
			return
		}, 1)
		reg35942 := PrimCons(reg35937, reg35938)
		reg35943 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35942)
		_ = reg35943
		reg35944 := MakeSymbol("shen.type-signature-of-get-time")
		reg35945 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4209 := __args[0]
			_ = V4209
			reg35946 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4210 := __args[0]
				_ = V4210
				reg35947 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4211 := __args[0]
					_ = V4211
					__e.TailApply(__defun__shen_4type_1signature_1of_1get_1time, V4209, V4210, V4211)
					return
				}, 1)
				__e.Return(reg35947)
				return
			}, 1)
			__e.Return(reg35946)
			return
		}, 1)
		reg35949 := PrimCons(reg35944, reg35945)
		reg35950 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35949)
		_ = reg35950
		reg35951 := MakeSymbol("shen.type-signature-of-hash")
		reg35952 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4199 := __args[0]
			_ = V4199
			reg35953 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4200 := __args[0]
				_ = V4200
				reg35954 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4201 := __args[0]
					_ = V4201
					__e.TailApply(__defun__shen_4type_1signature_1of_1hash, V4199, V4200, V4201)
					return
				}, 1)
				__e.Return(reg35954)
				return
			}, 1)
			__e.Return(reg35953)
			return
		}, 1)
		reg35956 := PrimCons(reg35951, reg35952)
		reg35957 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35956)
		_ = reg35957
		reg35958 := MakeSymbol("shen.type-signature-of-head")
		reg35959 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4189 := __args[0]
			_ = V4189
			reg35960 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4190 := __args[0]
				_ = V4190
				reg35961 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4191 := __args[0]
					_ = V4191
					__e.TailApply(__defun__shen_4type_1signature_1of_1head, V4189, V4190, V4191)
					return
				}, 1)
				__e.Return(reg35961)
				return
			}, 1)
			__e.Return(reg35960)
			return
		}, 1)
		reg35963 := PrimCons(reg35958, reg35959)
		reg35964 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35963)
		_ = reg35964
		reg35965 := MakeSymbol("shen.type-signature-of-hdv")
		reg35966 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4179 := __args[0]
			_ = V4179
			reg35967 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4180 := __args[0]
				_ = V4180
				reg35968 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4181 := __args[0]
					_ = V4181
					__e.TailApply(__defun__shen_4type_1signature_1of_1hdv, V4179, V4180, V4181)
					return
				}, 1)
				__e.Return(reg35968)
				return
			}, 1)
			__e.Return(reg35967)
			return
		}, 1)
		reg35970 := PrimCons(reg35965, reg35966)
		reg35971 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35970)
		_ = reg35971
		reg35972 := MakeSymbol("shen.type-signature-of-hdstr")
		reg35973 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4169 := __args[0]
			_ = V4169
			reg35974 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4170 := __args[0]
				_ = V4170
				reg35975 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4171 := __args[0]
					_ = V4171
					__e.TailApply(__defun__shen_4type_1signature_1of_1hdstr, V4169, V4170, V4171)
					return
				}, 1)
				__e.Return(reg35975)
				return
			}, 1)
			__e.Return(reg35974)
			return
		}, 1)
		reg35977 := PrimCons(reg35972, reg35973)
		reg35978 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35977)
		_ = reg35978
		reg35979 := MakeSymbol("shen.type-signature-of-if")
		reg35980 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4159 := __args[0]
			_ = V4159
			reg35981 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4160 := __args[0]
				_ = V4160
				reg35982 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4161 := __args[0]
					_ = V4161
					__e.TailApply(__defun__shen_4type_1signature_1of_1if, V4159, V4160, V4161)
					return
				}, 1)
				__e.Return(reg35982)
				return
			}, 1)
			__e.Return(reg35981)
			return
		}, 1)
		reg35984 := PrimCons(reg35979, reg35980)
		reg35985 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35984)
		_ = reg35985
		reg35986 := MakeSymbol("shen.type-signature-of-it")
		reg35987 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4149 := __args[0]
			_ = V4149
			reg35988 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4150 := __args[0]
				_ = V4150
				reg35989 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4151 := __args[0]
					_ = V4151
					__e.TailApply(__defun__shen_4type_1signature_1of_1it, V4149, V4150, V4151)
					return
				}, 1)
				__e.Return(reg35989)
				return
			}, 1)
			__e.Return(reg35988)
			return
		}, 1)
		reg35991 := PrimCons(reg35986, reg35987)
		reg35992 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35991)
		_ = reg35992
		reg35993 := MakeSymbol("shen.type-signature-of-implementation")
		reg35994 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4139 := __args[0]
			_ = V4139
			reg35995 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4140 := __args[0]
				_ = V4140
				reg35996 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4141 := __args[0]
					_ = V4141
					__e.TailApply(__defun__shen_4type_1signature_1of_1implementation, V4139, V4140, V4141)
					return
				}, 1)
				__e.Return(reg35996)
				return
			}, 1)
			__e.Return(reg35995)
			return
		}, 1)
		reg35998 := PrimCons(reg35993, reg35994)
		reg35999 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg35998)
		_ = reg35999
		reg36000 := MakeSymbol("shen.type-signature-of-include")
		reg36001 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4129 := __args[0]
			_ = V4129
			reg36002 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4130 := __args[0]
				_ = V4130
				reg36003 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4131 := __args[0]
					_ = V4131
					__e.TailApply(__defun__shen_4type_1signature_1of_1include, V4129, V4130, V4131)
					return
				}, 1)
				__e.Return(reg36003)
				return
			}, 1)
			__e.Return(reg36002)
			return
		}, 1)
		reg36005 := PrimCons(reg36000, reg36001)
		reg36006 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36005)
		_ = reg36006
		reg36007 := MakeSymbol("shen.type-signature-of-include-all-but")
		reg36008 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4119 := __args[0]
			_ = V4119
			reg36009 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4120 := __args[0]
				_ = V4120
				reg36010 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4121 := __args[0]
					_ = V4121
					__e.TailApply(__defun__shen_4type_1signature_1of_1include_1all_1but, V4119, V4120, V4121)
					return
				}, 1)
				__e.Return(reg36010)
				return
			}, 1)
			__e.Return(reg36009)
			return
		}, 1)
		reg36012 := PrimCons(reg36007, reg36008)
		reg36013 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36012)
		_ = reg36013
		reg36014 := MakeSymbol("shen.type-signature-of-inferences")
		reg36015 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4109 := __args[0]
			_ = V4109
			reg36016 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4110 := __args[0]
				_ = V4110
				reg36017 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4111 := __args[0]
					_ = V4111
					__e.TailApply(__defun__shen_4type_1signature_1of_1inferences, V4109, V4110, V4111)
					return
				}, 1)
				__e.Return(reg36017)
				return
			}, 1)
			__e.Return(reg36016)
			return
		}, 1)
		reg36019 := PrimCons(reg36014, reg36015)
		reg36020 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36019)
		_ = reg36020
		reg36021 := MakeSymbol("shen.type-signature-of-shen.insert")
		reg36022 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4099 := __args[0]
			_ = V4099
			reg36023 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4100 := __args[0]
				_ = V4100
				reg36024 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4101 := __args[0]
					_ = V4101
					__e.TailApply(__defun__shen_4type_1signature_1of_1shen_4insert, V4099, V4100, V4101)
					return
				}, 1)
				__e.Return(reg36024)
				return
			}, 1)
			__e.Return(reg36023)
			return
		}, 1)
		reg36026 := PrimCons(reg36021, reg36022)
		reg36027 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36026)
		_ = reg36027
		reg36028 := MakeSymbol("shen.type-signature-of-integer?")
		reg36029 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4089 := __args[0]
			_ = V4089
			reg36030 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4090 := __args[0]
				_ = V4090
				reg36031 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4091 := __args[0]
					_ = V4091
					__e.TailApply(__defun__shen_4type_1signature_1of_1integer_2, V4089, V4090, V4091)
					return
				}, 1)
				__e.Return(reg36031)
				return
			}, 1)
			__e.Return(reg36030)
			return
		}, 1)
		reg36033 := PrimCons(reg36028, reg36029)
		reg36034 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36033)
		_ = reg36034
		reg36035 := MakeSymbol("shen.type-signature-of-internal")
		reg36036 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4079 := __args[0]
			_ = V4079
			reg36037 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4080 := __args[0]
				_ = V4080
				reg36038 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4081 := __args[0]
					_ = V4081
					__e.TailApply(__defun__shen_4type_1signature_1of_1internal, V4079, V4080, V4081)
					return
				}, 1)
				__e.Return(reg36038)
				return
			}, 1)
			__e.Return(reg36037)
			return
		}, 1)
		reg36040 := PrimCons(reg36035, reg36036)
		reg36041 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36040)
		_ = reg36041
		reg36042 := MakeSymbol("shen.type-signature-of-intersection")
		reg36043 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4069 := __args[0]
			_ = V4069
			reg36044 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4070 := __args[0]
				_ = V4070
				reg36045 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4071 := __args[0]
					_ = V4071
					__e.TailApply(__defun__shen_4type_1signature_1of_1intersection, V4069, V4070, V4071)
					return
				}, 1)
				__e.Return(reg36045)
				return
			}, 1)
			__e.Return(reg36044)
			return
		}, 1)
		reg36047 := PrimCons(reg36042, reg36043)
		reg36048 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36047)
		_ = reg36048
		reg36049 := MakeSymbol("shen.type-signature-of-kill")
		reg36050 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4059 := __args[0]
			_ = V4059
			reg36051 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4060 := __args[0]
				_ = V4060
				reg36052 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4061 := __args[0]
					_ = V4061
					__e.TailApply(__defun__shen_4type_1signature_1of_1kill, V4059, V4060, V4061)
					return
				}, 1)
				__e.Return(reg36052)
				return
			}, 1)
			__e.Return(reg36051)
			return
		}, 1)
		reg36054 := PrimCons(reg36049, reg36050)
		reg36055 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36054)
		_ = reg36055
		reg36056 := MakeSymbol("shen.type-signature-of-language")
		reg36057 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4049 := __args[0]
			_ = V4049
			reg36058 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4050 := __args[0]
				_ = V4050
				reg36059 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4051 := __args[0]
					_ = V4051
					__e.TailApply(__defun__shen_4type_1signature_1of_1language, V4049, V4050, V4051)
					return
				}, 1)
				__e.Return(reg36059)
				return
			}, 1)
			__e.Return(reg36058)
			return
		}, 1)
		reg36061 := PrimCons(reg36056, reg36057)
		reg36062 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36061)
		_ = reg36062
		reg36063 := MakeSymbol("shen.type-signature-of-length")
		reg36064 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4039 := __args[0]
			_ = V4039
			reg36065 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4040 := __args[0]
				_ = V4040
				reg36066 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4041 := __args[0]
					_ = V4041
					__e.TailApply(__defun__shen_4type_1signature_1of_1length, V4039, V4040, V4041)
					return
				}, 1)
				__e.Return(reg36066)
				return
			}, 1)
			__e.Return(reg36065)
			return
		}, 1)
		reg36068 := PrimCons(reg36063, reg36064)
		reg36069 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36068)
		_ = reg36069
		reg36070 := MakeSymbol("shen.type-signature-of-limit")
		reg36071 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4029 := __args[0]
			_ = V4029
			reg36072 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4030 := __args[0]
				_ = V4030
				reg36073 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4031 := __args[0]
					_ = V4031
					__e.TailApply(__defun__shen_4type_1signature_1of_1limit, V4029, V4030, V4031)
					return
				}, 1)
				__e.Return(reg36073)
				return
			}, 1)
			__e.Return(reg36072)
			return
		}, 1)
		reg36075 := PrimCons(reg36070, reg36071)
		reg36076 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36075)
		_ = reg36076
		reg36077 := MakeSymbol("shen.type-signature-of-load")
		reg36078 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4019 := __args[0]
			_ = V4019
			reg36079 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4020 := __args[0]
				_ = V4020
				reg36080 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4021 := __args[0]
					_ = V4021
					__e.TailApply(__defun__shen_4type_1signature_1of_1load, V4019, V4020, V4021)
					return
				}, 1)
				__e.Return(reg36080)
				return
			}, 1)
			__e.Return(reg36079)
			return
		}, 1)
		reg36082 := PrimCons(reg36077, reg36078)
		reg36083 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36082)
		_ = reg36083
		reg36084 := MakeSymbol("shen.type-signature-of-map")
		reg36085 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V4009 := __args[0]
			_ = V4009
			reg36086 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4010 := __args[0]
				_ = V4010
				reg36087 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4011 := __args[0]
					_ = V4011
					__e.TailApply(__defun__shen_4type_1signature_1of_1map, V4009, V4010, V4011)
					return
				}, 1)
				__e.Return(reg36087)
				return
			}, 1)
			__e.Return(reg36086)
			return
		}, 1)
		reg36089 := PrimCons(reg36084, reg36085)
		reg36090 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36089)
		_ = reg36090
		reg36091 := MakeSymbol("shen.type-signature-of-mapcan")
		reg36092 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3999 := __args[0]
			_ = V3999
			reg36093 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4000 := __args[0]
				_ = V4000
				reg36094 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V4001 := __args[0]
					_ = V4001
					__e.TailApply(__defun__shen_4type_1signature_1of_1mapcan, V3999, V4000, V4001)
					return
				}, 1)
				__e.Return(reg36094)
				return
			}, 1)
			__e.Return(reg36093)
			return
		}, 1)
		reg36096 := PrimCons(reg36091, reg36092)
		reg36097 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36096)
		_ = reg36097
		reg36098 := MakeSymbol("shen.type-signature-of-maxinferences")
		reg36099 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3989 := __args[0]
			_ = V3989
			reg36100 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3990 := __args[0]
				_ = V3990
				reg36101 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3991 := __args[0]
					_ = V3991
					__e.TailApply(__defun__shen_4type_1signature_1of_1maxinferences, V3989, V3990, V3991)
					return
				}, 1)
				__e.Return(reg36101)
				return
			}, 1)
			__e.Return(reg36100)
			return
		}, 1)
		reg36103 := PrimCons(reg36098, reg36099)
		reg36104 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36103)
		_ = reg36104
		reg36105 := MakeSymbol("shen.type-signature-of-n->string")
		reg36106 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3979 := __args[0]
			_ = V3979
			reg36107 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3980 := __args[0]
				_ = V3980
				reg36108 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3981 := __args[0]
					_ = V3981
					__e.TailApply(__defun__shen_4type_1signature_1of_1n_1_6string, V3979, V3980, V3981)
					return
				}, 1)
				__e.Return(reg36108)
				return
			}, 1)
			__e.Return(reg36107)
			return
		}, 1)
		reg36110 := PrimCons(reg36105, reg36106)
		reg36111 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36110)
		_ = reg36111
		reg36112 := MakeSymbol("shen.type-signature-of-nl")
		reg36113 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3969 := __args[0]
			_ = V3969
			reg36114 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3970 := __args[0]
				_ = V3970
				reg36115 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3971 := __args[0]
					_ = V3971
					__e.TailApply(__defun__shen_4type_1signature_1of_1nl, V3969, V3970, V3971)
					return
				}, 1)
				__e.Return(reg36115)
				return
			}, 1)
			__e.Return(reg36114)
			return
		}, 1)
		reg36117 := PrimCons(reg36112, reg36113)
		reg36118 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36117)
		_ = reg36118
		reg36119 := MakeSymbol("shen.type-signature-of-not")
		reg36120 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3959 := __args[0]
			_ = V3959
			reg36121 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3960 := __args[0]
				_ = V3960
				reg36122 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3961 := __args[0]
					_ = V3961
					__e.TailApply(__defun__shen_4type_1signature_1of_1not, V3959, V3960, V3961)
					return
				}, 1)
				__e.Return(reg36122)
				return
			}, 1)
			__e.Return(reg36121)
			return
		}, 1)
		reg36124 := PrimCons(reg36119, reg36120)
		reg36125 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36124)
		_ = reg36125
		reg36126 := MakeSymbol("shen.type-signature-of-nth")
		reg36127 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3949 := __args[0]
			_ = V3949
			reg36128 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3950 := __args[0]
				_ = V3950
				reg36129 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3951 := __args[0]
					_ = V3951
					__e.TailApply(__defun__shen_4type_1signature_1of_1nth, V3949, V3950, V3951)
					return
				}, 1)
				__e.Return(reg36129)
				return
			}, 1)
			__e.Return(reg36128)
			return
		}, 1)
		reg36131 := PrimCons(reg36126, reg36127)
		reg36132 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36131)
		_ = reg36132
		reg36133 := MakeSymbol("shen.type-signature-of-number?")
		reg36134 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3939 := __args[0]
			_ = V3939
			reg36135 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3940 := __args[0]
				_ = V3940
				reg36136 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3941 := __args[0]
					_ = V3941
					__e.TailApply(__defun__shen_4type_1signature_1of_1number_2, V3939, V3940, V3941)
					return
				}, 1)
				__e.Return(reg36136)
				return
			}, 1)
			__e.Return(reg36135)
			return
		}, 1)
		reg36138 := PrimCons(reg36133, reg36134)
		reg36139 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36138)
		_ = reg36139
		reg36140 := MakeSymbol("shen.type-signature-of-occurrences")
		reg36141 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3929 := __args[0]
			_ = V3929
			reg36142 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3930 := __args[0]
				_ = V3930
				reg36143 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3931 := __args[0]
					_ = V3931
					__e.TailApply(__defun__shen_4type_1signature_1of_1occurrences, V3929, V3930, V3931)
					return
				}, 1)
				__e.Return(reg36143)
				return
			}, 1)
			__e.Return(reg36142)
			return
		}, 1)
		reg36145 := PrimCons(reg36140, reg36141)
		reg36146 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36145)
		_ = reg36146
		reg36147 := MakeSymbol("shen.type-signature-of-occurs-check")
		reg36148 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3919 := __args[0]
			_ = V3919
			reg36149 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3920 := __args[0]
				_ = V3920
				reg36150 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3921 := __args[0]
					_ = V3921
					__e.TailApply(__defun__shen_4type_1signature_1of_1occurs_1check, V3919, V3920, V3921)
					return
				}, 1)
				__e.Return(reg36150)
				return
			}, 1)
			__e.Return(reg36149)
			return
		}, 1)
		reg36152 := PrimCons(reg36147, reg36148)
		reg36153 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36152)
		_ = reg36153
		reg36154 := MakeSymbol("shen.type-signature-of-optimise")
		reg36155 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3909 := __args[0]
			_ = V3909
			reg36156 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3910 := __args[0]
				_ = V3910
				reg36157 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3911 := __args[0]
					_ = V3911
					__e.TailApply(__defun__shen_4type_1signature_1of_1optimise, V3909, V3910, V3911)
					return
				}, 1)
				__e.Return(reg36157)
				return
			}, 1)
			__e.Return(reg36156)
			return
		}, 1)
		reg36159 := PrimCons(reg36154, reg36155)
		reg36160 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36159)
		_ = reg36160
		reg36161 := MakeSymbol("shen.type-signature-of-or")
		reg36162 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3899 := __args[0]
			_ = V3899
			reg36163 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3900 := __args[0]
				_ = V3900
				reg36164 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3901 := __args[0]
					_ = V3901
					__e.TailApply(__defun__shen_4type_1signature_1of_1or, V3899, V3900, V3901)
					return
				}, 1)
				__e.Return(reg36164)
				return
			}, 1)
			__e.Return(reg36163)
			return
		}, 1)
		reg36166 := PrimCons(reg36161, reg36162)
		reg36167 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36166)
		_ = reg36167
		reg36168 := MakeSymbol("shen.type-signature-of-os")
		reg36169 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3889 := __args[0]
			_ = V3889
			reg36170 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3890 := __args[0]
				_ = V3890
				reg36171 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3891 := __args[0]
					_ = V3891
					__e.TailApply(__defun__shen_4type_1signature_1of_1os, V3889, V3890, V3891)
					return
				}, 1)
				__e.Return(reg36171)
				return
			}, 1)
			__e.Return(reg36170)
			return
		}, 1)
		reg36173 := PrimCons(reg36168, reg36169)
		reg36174 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36173)
		_ = reg36174
		reg36175 := MakeSymbol("shen.type-signature-of-package?")
		reg36176 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3879 := __args[0]
			_ = V3879
			reg36177 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3880 := __args[0]
				_ = V3880
				reg36178 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3881 := __args[0]
					_ = V3881
					__e.TailApply(__defun__shen_4type_1signature_1of_1package_2, V3879, V3880, V3881)
					return
				}, 1)
				__e.Return(reg36178)
				return
			}, 1)
			__e.Return(reg36177)
			return
		}, 1)
		reg36180 := PrimCons(reg36175, reg36176)
		reg36181 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36180)
		_ = reg36181
		reg36182 := MakeSymbol("shen.type-signature-of-port")
		reg36183 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3869 := __args[0]
			_ = V3869
			reg36184 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3870 := __args[0]
				_ = V3870
				reg36185 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3871 := __args[0]
					_ = V3871
					__e.TailApply(__defun__shen_4type_1signature_1of_1port, V3869, V3870, V3871)
					return
				}, 1)
				__e.Return(reg36185)
				return
			}, 1)
			__e.Return(reg36184)
			return
		}, 1)
		reg36187 := PrimCons(reg36182, reg36183)
		reg36188 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36187)
		_ = reg36188
		reg36189 := MakeSymbol("shen.type-signature-of-porters")
		reg36190 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3859 := __args[0]
			_ = V3859
			reg36191 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3860 := __args[0]
				_ = V3860
				reg36192 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3861 := __args[0]
					_ = V3861
					__e.TailApply(__defun__shen_4type_1signature_1of_1porters, V3859, V3860, V3861)
					return
				}, 1)
				__e.Return(reg36192)
				return
			}, 1)
			__e.Return(reg36191)
			return
		}, 1)
		reg36194 := PrimCons(reg36189, reg36190)
		reg36195 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36194)
		_ = reg36195
		reg36196 := MakeSymbol("shen.type-signature-of-pos")
		reg36197 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3849 := __args[0]
			_ = V3849
			reg36198 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3850 := __args[0]
				_ = V3850
				reg36199 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3851 := __args[0]
					_ = V3851
					__e.TailApply(__defun__shen_4type_1signature_1of_1pos, V3849, V3850, V3851)
					return
				}, 1)
				__e.Return(reg36199)
				return
			}, 1)
			__e.Return(reg36198)
			return
		}, 1)
		reg36201 := PrimCons(reg36196, reg36197)
		reg36202 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36201)
		_ = reg36202
		reg36203 := MakeSymbol("shen.type-signature-of-pr")
		reg36204 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3839 := __args[0]
			_ = V3839
			reg36205 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3840 := __args[0]
				_ = V3840
				reg36206 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3841 := __args[0]
					_ = V3841
					__e.TailApply(__defun__shen_4type_1signature_1of_1pr, V3839, V3840, V3841)
					return
				}, 1)
				__e.Return(reg36206)
				return
			}, 1)
			__e.Return(reg36205)
			return
		}, 1)
		reg36208 := PrimCons(reg36203, reg36204)
		reg36209 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36208)
		_ = reg36209
		reg36210 := MakeSymbol("shen.type-signature-of-print")
		reg36211 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3829 := __args[0]
			_ = V3829
			reg36212 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3830 := __args[0]
				_ = V3830
				reg36213 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3831 := __args[0]
					_ = V3831
					__e.TailApply(__defun__shen_4type_1signature_1of_1print, V3829, V3830, V3831)
					return
				}, 1)
				__e.Return(reg36213)
				return
			}, 1)
			__e.Return(reg36212)
			return
		}, 1)
		reg36215 := PrimCons(reg36210, reg36211)
		reg36216 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36215)
		_ = reg36216
		reg36217 := MakeSymbol("shen.type-signature-of-profile")
		reg36218 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3819 := __args[0]
			_ = V3819
			reg36219 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3820 := __args[0]
				_ = V3820
				reg36220 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3821 := __args[0]
					_ = V3821
					__e.TailApply(__defun__shen_4type_1signature_1of_1profile, V3819, V3820, V3821)
					return
				}, 1)
				__e.Return(reg36220)
				return
			}, 1)
			__e.Return(reg36219)
			return
		}, 1)
		reg36222 := PrimCons(reg36217, reg36218)
		reg36223 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36222)
		_ = reg36223
		reg36224 := MakeSymbol("shen.type-signature-of-preclude")
		reg36225 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3809 := __args[0]
			_ = V3809
			reg36226 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3810 := __args[0]
				_ = V3810
				reg36227 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3811 := __args[0]
					_ = V3811
					__e.TailApply(__defun__shen_4type_1signature_1of_1preclude, V3809, V3810, V3811)
					return
				}, 1)
				__e.Return(reg36227)
				return
			}, 1)
			__e.Return(reg36226)
			return
		}, 1)
		reg36229 := PrimCons(reg36224, reg36225)
		reg36230 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36229)
		_ = reg36230
		reg36231 := MakeSymbol("shen.type-signature-of-shen.proc-nl")
		reg36232 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3799 := __args[0]
			_ = V3799
			reg36233 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3800 := __args[0]
				_ = V3800
				reg36234 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3801 := __args[0]
					_ = V3801
					__e.TailApply(__defun__shen_4type_1signature_1of_1shen_4proc_1nl, V3799, V3800, V3801)
					return
				}, 1)
				__e.Return(reg36234)
				return
			}, 1)
			__e.Return(reg36233)
			return
		}, 1)
		reg36236 := PrimCons(reg36231, reg36232)
		reg36237 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36236)
		_ = reg36237
		reg36238 := MakeSymbol("shen.type-signature-of-profile-results")
		reg36239 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3789 := __args[0]
			_ = V3789
			reg36240 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3790 := __args[0]
				_ = V3790
				reg36241 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3791 := __args[0]
					_ = V3791
					__e.TailApply(__defun__shen_4type_1signature_1of_1profile_1results, V3789, V3790, V3791)
					return
				}, 1)
				__e.Return(reg36241)
				return
			}, 1)
			__e.Return(reg36240)
			return
		}, 1)
		reg36243 := PrimCons(reg36238, reg36239)
		reg36244 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36243)
		_ = reg36244
		reg36245 := MakeSymbol("shen.type-signature-of-protect")
		reg36246 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3779 := __args[0]
			_ = V3779
			reg36247 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3780 := __args[0]
				_ = V3780
				reg36248 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3781 := __args[0]
					_ = V3781
					__e.TailApply(__defun__shen_4type_1signature_1of_1protect, V3779, V3780, V3781)
					return
				}, 1)
				__e.Return(reg36248)
				return
			}, 1)
			__e.Return(reg36247)
			return
		}, 1)
		reg36250 := PrimCons(reg36245, reg36246)
		reg36251 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36250)
		_ = reg36251
		reg36252 := MakeSymbol("shen.type-signature-of-preclude-all-but")
		reg36253 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3769 := __args[0]
			_ = V3769
			reg36254 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3770 := __args[0]
				_ = V3770
				reg36255 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3771 := __args[0]
					_ = V3771
					__e.TailApply(__defun__shen_4type_1signature_1of_1preclude_1all_1but, V3769, V3770, V3771)
					return
				}, 1)
				__e.Return(reg36255)
				return
			}, 1)
			__e.Return(reg36254)
			return
		}, 1)
		reg36257 := PrimCons(reg36252, reg36253)
		reg36258 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36257)
		_ = reg36258
		reg36259 := MakeSymbol("shen.type-signature-of-shen.prhush")
		reg36260 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3759 := __args[0]
			_ = V3759
			reg36261 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3760 := __args[0]
				_ = V3760
				reg36262 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3761 := __args[0]
					_ = V3761
					__e.TailApply(__defun__shen_4type_1signature_1of_1shen_4prhush, V3759, V3760, V3761)
					return
				}, 1)
				__e.Return(reg36262)
				return
			}, 1)
			__e.Return(reg36261)
			return
		}, 1)
		reg36264 := PrimCons(reg36259, reg36260)
		reg36265 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36264)
		_ = reg36265
		reg36266 := MakeSymbol("shen.type-signature-of-ps")
		reg36267 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3749 := __args[0]
			_ = V3749
			reg36268 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3750 := __args[0]
				_ = V3750
				reg36269 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3751 := __args[0]
					_ = V3751
					__e.TailApply(__defun__shen_4type_1signature_1of_1ps, V3749, V3750, V3751)
					return
				}, 1)
				__e.Return(reg36269)
				return
			}, 1)
			__e.Return(reg36268)
			return
		}, 1)
		reg36271 := PrimCons(reg36266, reg36267)
		reg36272 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36271)
		_ = reg36272
		reg36273 := MakeSymbol("shen.type-signature-of-read")
		reg36274 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3739 := __args[0]
			_ = V3739
			reg36275 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3740 := __args[0]
				_ = V3740
				reg36276 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3741 := __args[0]
					_ = V3741
					__e.TailApply(__defun__shen_4type_1signature_1of_1read, V3739, V3740, V3741)
					return
				}, 1)
				__e.Return(reg36276)
				return
			}, 1)
			__e.Return(reg36275)
			return
		}, 1)
		reg36278 := PrimCons(reg36273, reg36274)
		reg36279 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36278)
		_ = reg36279
		reg36280 := MakeSymbol("shen.type-signature-of-read-byte")
		reg36281 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3729 := __args[0]
			_ = V3729
			reg36282 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3730 := __args[0]
				_ = V3730
				reg36283 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3731 := __args[0]
					_ = V3731
					__e.TailApply(__defun__shen_4type_1signature_1of_1read_1byte, V3729, V3730, V3731)
					return
				}, 1)
				__e.Return(reg36283)
				return
			}, 1)
			__e.Return(reg36282)
			return
		}, 1)
		reg36285 := PrimCons(reg36280, reg36281)
		reg36286 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36285)
		_ = reg36286
		reg36287 := MakeSymbol("shen.type-signature-of-read-file-as-bytelist")
		reg36288 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3719 := __args[0]
			_ = V3719
			reg36289 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3720 := __args[0]
				_ = V3720
				reg36290 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3721 := __args[0]
					_ = V3721
					__e.TailApply(__defun__shen_4type_1signature_1of_1read_1file_1as_1bytelist, V3719, V3720, V3721)
					return
				}, 1)
				__e.Return(reg36290)
				return
			}, 1)
			__e.Return(reg36289)
			return
		}, 1)
		reg36292 := PrimCons(reg36287, reg36288)
		reg36293 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36292)
		_ = reg36293
		reg36294 := MakeSymbol("shen.type-signature-of-read-file-as-string")
		reg36295 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3709 := __args[0]
			_ = V3709
			reg36296 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3710 := __args[0]
				_ = V3710
				reg36297 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3711 := __args[0]
					_ = V3711
					__e.TailApply(__defun__shen_4type_1signature_1of_1read_1file_1as_1string, V3709, V3710, V3711)
					return
				}, 1)
				__e.Return(reg36297)
				return
			}, 1)
			__e.Return(reg36296)
			return
		}, 1)
		reg36299 := PrimCons(reg36294, reg36295)
		reg36300 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36299)
		_ = reg36300
		reg36301 := MakeSymbol("shen.type-signature-of-read-file")
		reg36302 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3699 := __args[0]
			_ = V3699
			reg36303 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3700 := __args[0]
				_ = V3700
				reg36304 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3701 := __args[0]
					_ = V3701
					__e.TailApply(__defun__shen_4type_1signature_1of_1read_1file, V3699, V3700, V3701)
					return
				}, 1)
				__e.Return(reg36304)
				return
			}, 1)
			__e.Return(reg36303)
			return
		}, 1)
		reg36306 := PrimCons(reg36301, reg36302)
		reg36307 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36306)
		_ = reg36307
		reg36308 := MakeSymbol("shen.type-signature-of-read-from-string")
		reg36309 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3689 := __args[0]
			_ = V3689
			reg36310 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3690 := __args[0]
				_ = V3690
				reg36311 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3691 := __args[0]
					_ = V3691
					__e.TailApply(__defun__shen_4type_1signature_1of_1read_1from_1string, V3689, V3690, V3691)
					return
				}, 1)
				__e.Return(reg36311)
				return
			}, 1)
			__e.Return(reg36310)
			return
		}, 1)
		reg36313 := PrimCons(reg36308, reg36309)
		reg36314 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36313)
		_ = reg36314
		reg36315 := MakeSymbol("shen.type-signature-of-release")
		reg36316 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3679 := __args[0]
			_ = V3679
			reg36317 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3680 := __args[0]
				_ = V3680
				reg36318 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3681 := __args[0]
					_ = V3681
					__e.TailApply(__defun__shen_4type_1signature_1of_1release, V3679, V3680, V3681)
					return
				}, 1)
				__e.Return(reg36318)
				return
			}, 1)
			__e.Return(reg36317)
			return
		}, 1)
		reg36320 := PrimCons(reg36315, reg36316)
		reg36321 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36320)
		_ = reg36321
		reg36322 := MakeSymbol("shen.type-signature-of-remove")
		reg36323 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3669 := __args[0]
			_ = V3669
			reg36324 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3670 := __args[0]
				_ = V3670
				reg36325 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3671 := __args[0]
					_ = V3671
					__e.TailApply(__defun__shen_4type_1signature_1of_1remove, V3669, V3670, V3671)
					return
				}, 1)
				__e.Return(reg36325)
				return
			}, 1)
			__e.Return(reg36324)
			return
		}, 1)
		reg36327 := PrimCons(reg36322, reg36323)
		reg36328 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36327)
		_ = reg36328
		reg36329 := MakeSymbol("shen.type-signature-of-reverse")
		reg36330 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3659 := __args[0]
			_ = V3659
			reg36331 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3660 := __args[0]
				_ = V3660
				reg36332 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3661 := __args[0]
					_ = V3661
					__e.TailApply(__defun__shen_4type_1signature_1of_1reverse, V3659, V3660, V3661)
					return
				}, 1)
				__e.Return(reg36332)
				return
			}, 1)
			__e.Return(reg36331)
			return
		}, 1)
		reg36334 := PrimCons(reg36329, reg36330)
		reg36335 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36334)
		_ = reg36335
		reg36336 := MakeSymbol("shen.type-signature-of-simple-error")
		reg36337 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3649 := __args[0]
			_ = V3649
			reg36338 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3650 := __args[0]
				_ = V3650
				reg36339 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3651 := __args[0]
					_ = V3651
					__e.TailApply(__defun__shen_4type_1signature_1of_1simple_1error, V3649, V3650, V3651)
					return
				}, 1)
				__e.Return(reg36339)
				return
			}, 1)
			__e.Return(reg36338)
			return
		}, 1)
		reg36341 := PrimCons(reg36336, reg36337)
		reg36342 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36341)
		_ = reg36342
		reg36343 := MakeSymbol("shen.type-signature-of-snd")
		reg36344 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3639 := __args[0]
			_ = V3639
			reg36345 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3640 := __args[0]
				_ = V3640
				reg36346 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3641 := __args[0]
					_ = V3641
					__e.TailApply(__defun__shen_4type_1signature_1of_1snd, V3639, V3640, V3641)
					return
				}, 1)
				__e.Return(reg36346)
				return
			}, 1)
			__e.Return(reg36345)
			return
		}, 1)
		reg36348 := PrimCons(reg36343, reg36344)
		reg36349 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36348)
		_ = reg36349
		reg36350 := MakeSymbol("shen.type-signature-of-specialise")
		reg36351 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3629 := __args[0]
			_ = V3629
			reg36352 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3630 := __args[0]
				_ = V3630
				reg36353 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3631 := __args[0]
					_ = V3631
					__e.TailApply(__defun__shen_4type_1signature_1of_1specialise, V3629, V3630, V3631)
					return
				}, 1)
				__e.Return(reg36353)
				return
			}, 1)
			__e.Return(reg36352)
			return
		}, 1)
		reg36355 := PrimCons(reg36350, reg36351)
		reg36356 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36355)
		_ = reg36356
		reg36357 := MakeSymbol("shen.type-signature-of-spy")
		reg36358 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3619 := __args[0]
			_ = V3619
			reg36359 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3620 := __args[0]
				_ = V3620
				reg36360 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3621 := __args[0]
					_ = V3621
					__e.TailApply(__defun__shen_4type_1signature_1of_1spy, V3619, V3620, V3621)
					return
				}, 1)
				__e.Return(reg36360)
				return
			}, 1)
			__e.Return(reg36359)
			return
		}, 1)
		reg36362 := PrimCons(reg36357, reg36358)
		reg36363 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36362)
		_ = reg36363
		reg36364 := MakeSymbol("shen.type-signature-of-step")
		reg36365 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3609 := __args[0]
			_ = V3609
			reg36366 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3610 := __args[0]
				_ = V3610
				reg36367 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3611 := __args[0]
					_ = V3611
					__e.TailApply(__defun__shen_4type_1signature_1of_1step, V3609, V3610, V3611)
					return
				}, 1)
				__e.Return(reg36367)
				return
			}, 1)
			__e.Return(reg36366)
			return
		}, 1)
		reg36369 := PrimCons(reg36364, reg36365)
		reg36370 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36369)
		_ = reg36370
		reg36371 := MakeSymbol("shen.type-signature-of-stinput")
		reg36372 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3599 := __args[0]
			_ = V3599
			reg36373 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3600 := __args[0]
				_ = V3600
				reg36374 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3601 := __args[0]
					_ = V3601
					__e.TailApply(__defun__shen_4type_1signature_1of_1stinput, V3599, V3600, V3601)
					return
				}, 1)
				__e.Return(reg36374)
				return
			}, 1)
			__e.Return(reg36373)
			return
		}, 1)
		reg36376 := PrimCons(reg36371, reg36372)
		reg36377 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36376)
		_ = reg36377
		reg36378 := MakeSymbol("shen.type-signature-of-sterror")
		reg36379 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3589 := __args[0]
			_ = V3589
			reg36380 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3590 := __args[0]
				_ = V3590
				reg36381 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3591 := __args[0]
					_ = V3591
					__e.TailApply(__defun__shen_4type_1signature_1of_1sterror, V3589, V3590, V3591)
					return
				}, 1)
				__e.Return(reg36381)
				return
			}, 1)
			__e.Return(reg36380)
			return
		}, 1)
		reg36383 := PrimCons(reg36378, reg36379)
		reg36384 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36383)
		_ = reg36384
		reg36385 := MakeSymbol("shen.type-signature-of-stoutput")
		reg36386 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3579 := __args[0]
			_ = V3579
			reg36387 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3580 := __args[0]
				_ = V3580
				reg36388 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3581 := __args[0]
					_ = V3581
					__e.TailApply(__defun__shen_4type_1signature_1of_1stoutput, V3579, V3580, V3581)
					return
				}, 1)
				__e.Return(reg36388)
				return
			}, 1)
			__e.Return(reg36387)
			return
		}, 1)
		reg36390 := PrimCons(reg36385, reg36386)
		reg36391 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36390)
		_ = reg36391
		reg36392 := MakeSymbol("shen.type-signature-of-string?")
		reg36393 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3569 := __args[0]
			_ = V3569
			reg36394 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3570 := __args[0]
				_ = V3570
				reg36395 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3571 := __args[0]
					_ = V3571
					__e.TailApply(__defun__shen_4type_1signature_1of_1string_2, V3569, V3570, V3571)
					return
				}, 1)
				__e.Return(reg36395)
				return
			}, 1)
			__e.Return(reg36394)
			return
		}, 1)
		reg36397 := PrimCons(reg36392, reg36393)
		reg36398 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36397)
		_ = reg36398
		reg36399 := MakeSymbol("shen.type-signature-of-str")
		reg36400 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3559 := __args[0]
			_ = V3559
			reg36401 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3560 := __args[0]
				_ = V3560
				reg36402 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3561 := __args[0]
					_ = V3561
					__e.TailApply(__defun__shen_4type_1signature_1of_1str, V3559, V3560, V3561)
					return
				}, 1)
				__e.Return(reg36402)
				return
			}, 1)
			__e.Return(reg36401)
			return
		}, 1)
		reg36404 := PrimCons(reg36399, reg36400)
		reg36405 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36404)
		_ = reg36405
		reg36406 := MakeSymbol("shen.type-signature-of-string->n")
		reg36407 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3549 := __args[0]
			_ = V3549
			reg36408 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3550 := __args[0]
				_ = V3550
				reg36409 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3551 := __args[0]
					_ = V3551
					__e.TailApply(__defun__shen_4type_1signature_1of_1string_1_6n, V3549, V3550, V3551)
					return
				}, 1)
				__e.Return(reg36409)
				return
			}, 1)
			__e.Return(reg36408)
			return
		}, 1)
		reg36411 := PrimCons(reg36406, reg36407)
		reg36412 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36411)
		_ = reg36412
		reg36413 := MakeSymbol("shen.type-signature-of-string->symbol")
		reg36414 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3539 := __args[0]
			_ = V3539
			reg36415 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3540 := __args[0]
				_ = V3540
				reg36416 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3541 := __args[0]
					_ = V3541
					__e.TailApply(__defun__shen_4type_1signature_1of_1string_1_6symbol, V3539, V3540, V3541)
					return
				}, 1)
				__e.Return(reg36416)
				return
			}, 1)
			__e.Return(reg36415)
			return
		}, 1)
		reg36418 := PrimCons(reg36413, reg36414)
		reg36419 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36418)
		_ = reg36419
		reg36420 := MakeSymbol("shen.type-signature-of-sum")
		reg36421 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3529 := __args[0]
			_ = V3529
			reg36422 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3530 := __args[0]
				_ = V3530
				reg36423 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3531 := __args[0]
					_ = V3531
					__e.TailApply(__defun__shen_4type_1signature_1of_1sum, V3529, V3530, V3531)
					return
				}, 1)
				__e.Return(reg36423)
				return
			}, 1)
			__e.Return(reg36422)
			return
		}, 1)
		reg36425 := PrimCons(reg36420, reg36421)
		reg36426 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36425)
		_ = reg36426
		reg36427 := MakeSymbol("shen.type-signature-of-symbol?")
		reg36428 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3519 := __args[0]
			_ = V3519
			reg36429 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3520 := __args[0]
				_ = V3520
				reg36430 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3521 := __args[0]
					_ = V3521
					__e.TailApply(__defun__shen_4type_1signature_1of_1symbol_2, V3519, V3520, V3521)
					return
				}, 1)
				__e.Return(reg36430)
				return
			}, 1)
			__e.Return(reg36429)
			return
		}, 1)
		reg36432 := PrimCons(reg36427, reg36428)
		reg36433 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36432)
		_ = reg36433
		reg36434 := MakeSymbol("shen.type-signature-of-systemf")
		reg36435 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3509 := __args[0]
			_ = V3509
			reg36436 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3510 := __args[0]
				_ = V3510
				reg36437 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3511 := __args[0]
					_ = V3511
					__e.TailApply(__defun__shen_4type_1signature_1of_1systemf, V3509, V3510, V3511)
					return
				}, 1)
				__e.Return(reg36437)
				return
			}, 1)
			__e.Return(reg36436)
			return
		}, 1)
		reg36439 := PrimCons(reg36434, reg36435)
		reg36440 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36439)
		_ = reg36440
		reg36441 := MakeSymbol("shen.type-signature-of-tail")
		reg36442 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3499 := __args[0]
			_ = V3499
			reg36443 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3500 := __args[0]
				_ = V3500
				reg36444 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3501 := __args[0]
					_ = V3501
					__e.TailApply(__defun__shen_4type_1signature_1of_1tail, V3499, V3500, V3501)
					return
				}, 1)
				__e.Return(reg36444)
				return
			}, 1)
			__e.Return(reg36443)
			return
		}, 1)
		reg36446 := PrimCons(reg36441, reg36442)
		reg36447 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36446)
		_ = reg36447
		reg36448 := MakeSymbol("shen.type-signature-of-tlstr")
		reg36449 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3489 := __args[0]
			_ = V3489
			reg36450 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3490 := __args[0]
				_ = V3490
				reg36451 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3491 := __args[0]
					_ = V3491
					__e.TailApply(__defun__shen_4type_1signature_1of_1tlstr, V3489, V3490, V3491)
					return
				}, 1)
				__e.Return(reg36451)
				return
			}, 1)
			__e.Return(reg36450)
			return
		}, 1)
		reg36453 := PrimCons(reg36448, reg36449)
		reg36454 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36453)
		_ = reg36454
		reg36455 := MakeSymbol("shen.type-signature-of-tlv")
		reg36456 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3479 := __args[0]
			_ = V3479
			reg36457 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3480 := __args[0]
				_ = V3480
				reg36458 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3481 := __args[0]
					_ = V3481
					__e.TailApply(__defun__shen_4type_1signature_1of_1tlv, V3479, V3480, V3481)
					return
				}, 1)
				__e.Return(reg36458)
				return
			}, 1)
			__e.Return(reg36457)
			return
		}, 1)
		reg36460 := PrimCons(reg36455, reg36456)
		reg36461 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36460)
		_ = reg36461
		reg36462 := MakeSymbol("shen.type-signature-of-tc")
		reg36463 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3469 := __args[0]
			_ = V3469
			reg36464 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3470 := __args[0]
				_ = V3470
				reg36465 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3471 := __args[0]
					_ = V3471
					__e.TailApply(__defun__shen_4type_1signature_1of_1tc, V3469, V3470, V3471)
					return
				}, 1)
				__e.Return(reg36465)
				return
			}, 1)
			__e.Return(reg36464)
			return
		}, 1)
		reg36467 := PrimCons(reg36462, reg36463)
		reg36468 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36467)
		_ = reg36468
		reg36469 := MakeSymbol("shen.type-signature-of-tc?")
		reg36470 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3459 := __args[0]
			_ = V3459
			reg36471 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3460 := __args[0]
				_ = V3460
				reg36472 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3461 := __args[0]
					_ = V3461
					__e.TailApply(__defun__shen_4type_1signature_1of_1tc_2, V3459, V3460, V3461)
					return
				}, 1)
				__e.Return(reg36472)
				return
			}, 1)
			__e.Return(reg36471)
			return
		}, 1)
		reg36474 := PrimCons(reg36469, reg36470)
		reg36475 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36474)
		_ = reg36475
		reg36476 := MakeSymbol("shen.type-signature-of-thaw")
		reg36477 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3449 := __args[0]
			_ = V3449
			reg36478 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3450 := __args[0]
				_ = V3450
				reg36479 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3451 := __args[0]
					_ = V3451
					__e.TailApply(__defun__shen_4type_1signature_1of_1thaw, V3449, V3450, V3451)
					return
				}, 1)
				__e.Return(reg36479)
				return
			}, 1)
			__e.Return(reg36478)
			return
		}, 1)
		reg36481 := PrimCons(reg36476, reg36477)
		reg36482 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36481)
		_ = reg36482
		reg36483 := MakeSymbol("shen.type-signature-of-track")
		reg36484 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3439 := __args[0]
			_ = V3439
			reg36485 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3440 := __args[0]
				_ = V3440
				reg36486 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3441 := __args[0]
					_ = V3441
					__e.TailApply(__defun__shen_4type_1signature_1of_1track, V3439, V3440, V3441)
					return
				}, 1)
				__e.Return(reg36486)
				return
			}, 1)
			__e.Return(reg36485)
			return
		}, 1)
		reg36488 := PrimCons(reg36483, reg36484)
		reg36489 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36488)
		_ = reg36489
		reg36490 := MakeSymbol("shen.type-signature-of-trap-error")
		reg36491 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3429 := __args[0]
			_ = V3429
			reg36492 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3430 := __args[0]
				_ = V3430
				reg36493 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3431 := __args[0]
					_ = V3431
					__e.TailApply(__defun__shen_4type_1signature_1of_1trap_1error, V3429, V3430, V3431)
					return
				}, 1)
				__e.Return(reg36493)
				return
			}, 1)
			__e.Return(reg36492)
			return
		}, 1)
		reg36495 := PrimCons(reg36490, reg36491)
		reg36496 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36495)
		_ = reg36496
		reg36497 := MakeSymbol("shen.type-signature-of-tuple?")
		reg36498 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3419 := __args[0]
			_ = V3419
			reg36499 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3420 := __args[0]
				_ = V3420
				reg36500 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3421 := __args[0]
					_ = V3421
					__e.TailApply(__defun__shen_4type_1signature_1of_1tuple_2, V3419, V3420, V3421)
					return
				}, 1)
				__e.Return(reg36500)
				return
			}, 1)
			__e.Return(reg36499)
			return
		}, 1)
		reg36502 := PrimCons(reg36497, reg36498)
		reg36503 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36502)
		_ = reg36503
		reg36504 := MakeSymbol("shen.type-signature-of-undefmacro")
		reg36505 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3409 := __args[0]
			_ = V3409
			reg36506 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3410 := __args[0]
				_ = V3410
				reg36507 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3411 := __args[0]
					_ = V3411
					__e.TailApply(__defun__shen_4type_1signature_1of_1undefmacro, V3409, V3410, V3411)
					return
				}, 1)
				__e.Return(reg36507)
				return
			}, 1)
			__e.Return(reg36506)
			return
		}, 1)
		reg36509 := PrimCons(reg36504, reg36505)
		reg36510 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36509)
		_ = reg36510
		reg36511 := MakeSymbol("shen.type-signature-of-union")
		reg36512 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3399 := __args[0]
			_ = V3399
			reg36513 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3400 := __args[0]
				_ = V3400
				reg36514 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3401 := __args[0]
					_ = V3401
					__e.TailApply(__defun__shen_4type_1signature_1of_1union, V3399, V3400, V3401)
					return
				}, 1)
				__e.Return(reg36514)
				return
			}, 1)
			__e.Return(reg36513)
			return
		}, 1)
		reg36516 := PrimCons(reg36511, reg36512)
		reg36517 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36516)
		_ = reg36517
		reg36518 := MakeSymbol("shen.type-signature-of-unprofile")
		reg36519 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3389 := __args[0]
			_ = V3389
			reg36520 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3390 := __args[0]
				_ = V3390
				reg36521 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3391 := __args[0]
					_ = V3391
					__e.TailApply(__defun__shen_4type_1signature_1of_1unprofile, V3389, V3390, V3391)
					return
				}, 1)
				__e.Return(reg36521)
				return
			}, 1)
			__e.Return(reg36520)
			return
		}, 1)
		reg36523 := PrimCons(reg36518, reg36519)
		reg36524 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36523)
		_ = reg36524
		reg36525 := MakeSymbol("shen.type-signature-of-untrack")
		reg36526 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3379 := __args[0]
			_ = V3379
			reg36527 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3380 := __args[0]
				_ = V3380
				reg36528 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3381 := __args[0]
					_ = V3381
					__e.TailApply(__defun__shen_4type_1signature_1of_1untrack, V3379, V3380, V3381)
					return
				}, 1)
				__e.Return(reg36528)
				return
			}, 1)
			__e.Return(reg36527)
			return
		}, 1)
		reg36530 := PrimCons(reg36525, reg36526)
		reg36531 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36530)
		_ = reg36531
		reg36532 := MakeSymbol("shen.type-signature-of-unspecialise")
		reg36533 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3369 := __args[0]
			_ = V3369
			reg36534 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3370 := __args[0]
				_ = V3370
				reg36535 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3371 := __args[0]
					_ = V3371
					__e.TailApply(__defun__shen_4type_1signature_1of_1unspecialise, V3369, V3370, V3371)
					return
				}, 1)
				__e.Return(reg36535)
				return
			}, 1)
			__e.Return(reg36534)
			return
		}, 1)
		reg36537 := PrimCons(reg36532, reg36533)
		reg36538 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36537)
		_ = reg36538
		reg36539 := MakeSymbol("shen.type-signature-of-variable?")
		reg36540 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3359 := __args[0]
			_ = V3359
			reg36541 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3360 := __args[0]
				_ = V3360
				reg36542 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3361 := __args[0]
					_ = V3361
					__e.TailApply(__defun__shen_4type_1signature_1of_1variable_2, V3359, V3360, V3361)
					return
				}, 1)
				__e.Return(reg36542)
				return
			}, 1)
			__e.Return(reg36541)
			return
		}, 1)
		reg36544 := PrimCons(reg36539, reg36540)
		reg36545 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36544)
		_ = reg36545
		reg36546 := MakeSymbol("shen.type-signature-of-vector?")
		reg36547 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3349 := __args[0]
			_ = V3349
			reg36548 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3350 := __args[0]
				_ = V3350
				reg36549 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3351 := __args[0]
					_ = V3351
					__e.TailApply(__defun__shen_4type_1signature_1of_1vector_2, V3349, V3350, V3351)
					return
				}, 1)
				__e.Return(reg36549)
				return
			}, 1)
			__e.Return(reg36548)
			return
		}, 1)
		reg36551 := PrimCons(reg36546, reg36547)
		reg36552 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36551)
		_ = reg36552
		reg36553 := MakeSymbol("shen.type-signature-of-version")
		reg36554 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3339 := __args[0]
			_ = V3339
			reg36555 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3340 := __args[0]
				_ = V3340
				reg36556 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3341 := __args[0]
					_ = V3341
					__e.TailApply(__defun__shen_4type_1signature_1of_1version, V3339, V3340, V3341)
					return
				}, 1)
				__e.Return(reg36556)
				return
			}, 1)
			__e.Return(reg36555)
			return
		}, 1)
		reg36558 := PrimCons(reg36553, reg36554)
		reg36559 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36558)
		_ = reg36559
		reg36560 := MakeSymbol("shen.type-signature-of-write-to-file")
		reg36561 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3329 := __args[0]
			_ = V3329
			reg36562 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3330 := __args[0]
				_ = V3330
				reg36563 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3331 := __args[0]
					_ = V3331
					__e.TailApply(__defun__shen_4type_1signature_1of_1write_1to_1file, V3329, V3330, V3331)
					return
				}, 1)
				__e.Return(reg36563)
				return
			}, 1)
			__e.Return(reg36562)
			return
		}, 1)
		reg36565 := PrimCons(reg36560, reg36561)
		reg36566 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36565)
		_ = reg36566
		reg36567 := MakeSymbol("shen.type-signature-of-write-byte")
		reg36568 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3319 := __args[0]
			_ = V3319
			reg36569 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3320 := __args[0]
				_ = V3320
				reg36570 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3321 := __args[0]
					_ = V3321
					__e.TailApply(__defun__shen_4type_1signature_1of_1write_1byte, V3319, V3320, V3321)
					return
				}, 1)
				__e.Return(reg36570)
				return
			}, 1)
			__e.Return(reg36569)
			return
		}, 1)
		reg36572 := PrimCons(reg36567, reg36568)
		reg36573 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36572)
		_ = reg36573
		reg36574 := MakeSymbol("shen.type-signature-of-y-or-n?")
		reg36575 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3309 := __args[0]
			_ = V3309
			reg36576 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3310 := __args[0]
				_ = V3310
				reg36577 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3311 := __args[0]
					_ = V3311
					__e.TailApply(__defun__shen_4type_1signature_1of_1y_1or_1n_2, V3309, V3310, V3311)
					return
				}, 1)
				__e.Return(reg36577)
				return
			}, 1)
			__e.Return(reg36576)
			return
		}, 1)
		reg36579 := PrimCons(reg36574, reg36575)
		reg36580 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36579)
		_ = reg36580
		reg36581 := MakeSymbol("shen.type-signature-of->")
		reg36582 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3299 := __args[0]
			_ = V3299
			reg36583 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3300 := __args[0]
				_ = V3300
				reg36584 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3301 := __args[0]
					_ = V3301
					__e.TailApply(__defun__shen_4type_1signature_1of_1_6, V3299, V3300, V3301)
					return
				}, 1)
				__e.Return(reg36584)
				return
			}, 1)
			__e.Return(reg36583)
			return
		}, 1)
		reg36586 := PrimCons(reg36581, reg36582)
		reg36587 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36586)
		_ = reg36587
		reg36588 := MakeSymbol("shen.type-signature-of-<")
		reg36589 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3289 := __args[0]
			_ = V3289
			reg36590 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3290 := __args[0]
				_ = V3290
				reg36591 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3291 := __args[0]
					_ = V3291
					__e.TailApply(__defun__shen_4type_1signature_1of_1_5, V3289, V3290, V3291)
					return
				}, 1)
				__e.Return(reg36591)
				return
			}, 1)
			__e.Return(reg36590)
			return
		}, 1)
		reg36593 := PrimCons(reg36588, reg36589)
		reg36594 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36593)
		_ = reg36594
		reg36595 := MakeSymbol("shen.type-signature-of->=")
		reg36596 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3279 := __args[0]
			_ = V3279
			reg36597 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3280 := __args[0]
				_ = V3280
				reg36598 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3281 := __args[0]
					_ = V3281
					__e.TailApply(__defun__shen_4type_1signature_1of_1_6_a, V3279, V3280, V3281)
					return
				}, 1)
				__e.Return(reg36598)
				return
			}, 1)
			__e.Return(reg36597)
			return
		}, 1)
		reg36600 := PrimCons(reg36595, reg36596)
		reg36601 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36600)
		_ = reg36601
		reg36602 := MakeSymbol("shen.type-signature-of-<=")
		reg36603 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3269 := __args[0]
			_ = V3269
			reg36604 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3270 := __args[0]
				_ = V3270
				reg36605 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3271 := __args[0]
					_ = V3271
					__e.TailApply(__defun__shen_4type_1signature_1of_1_5_a, V3269, V3270, V3271)
					return
				}, 1)
				__e.Return(reg36605)
				return
			}, 1)
			__e.Return(reg36604)
			return
		}, 1)
		reg36607 := PrimCons(reg36602, reg36603)
		reg36608 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36607)
		_ = reg36608
		reg36609 := MakeSymbol("shen.type-signature-of-=")
		reg36610 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3259 := __args[0]
			_ = V3259
			reg36611 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3260 := __args[0]
				_ = V3260
				reg36612 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3261 := __args[0]
					_ = V3261
					__e.TailApply(__defun__shen_4type_1signature_1of_1_a, V3259, V3260, V3261)
					return
				}, 1)
				__e.Return(reg36612)
				return
			}, 1)
			__e.Return(reg36611)
			return
		}, 1)
		reg36614 := PrimCons(reg36609, reg36610)
		reg36615 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36614)
		_ = reg36615
		reg36616 := MakeSymbol("shen.type-signature-of-+")
		reg36617 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3249 := __args[0]
			_ = V3249
			reg36618 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3250 := __args[0]
				_ = V3250
				reg36619 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3251 := __args[0]
					_ = V3251
					__e.TailApply(__defun__shen_4type_1signature_1of_1_7, V3249, V3250, V3251)
					return
				}, 1)
				__e.Return(reg36619)
				return
			}, 1)
			__e.Return(reg36618)
			return
		}, 1)
		reg36621 := PrimCons(reg36616, reg36617)
		reg36622 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36621)
		_ = reg36622
		reg36623 := MakeSymbol("shen.type-signature-of-/")
		reg36624 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3239 := __args[0]
			_ = V3239
			reg36625 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3240 := __args[0]
				_ = V3240
				reg36626 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3241 := __args[0]
					_ = V3241
					__e.TailApply(__defun__shen_4type_1signature_1of_1_c, V3239, V3240, V3241)
					return
				}, 1)
				__e.Return(reg36626)
				return
			}, 1)
			__e.Return(reg36625)
			return
		}, 1)
		reg36628 := PrimCons(reg36623, reg36624)
		reg36629 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36628)
		_ = reg36629
		reg36630 := MakeSymbol("shen.type-signature-of--")
		reg36631 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3229 := __args[0]
			_ = V3229
			reg36632 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3230 := __args[0]
				_ = V3230
				reg36633 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3231 := __args[0]
					_ = V3231
					__e.TailApply(__defun__shen_4type_1signature_1of_1_1, V3229, V3230, V3231)
					return
				}, 1)
				__e.Return(reg36633)
				return
			}, 1)
			__e.Return(reg36632)
			return
		}, 1)
		reg36635 := PrimCons(reg36630, reg36631)
		reg36636 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36635)
		_ = reg36636
		reg36637 := MakeSymbol("shen.type-signature-of-*")
		reg36638 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3219 := __args[0]
			_ = V3219
			reg36639 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3220 := __args[0]
				_ = V3220
				reg36640 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3221 := __args[0]
					_ = V3221
					__e.TailApply(__defun__shen_4type_1signature_1of_1_d, V3219, V3220, V3221)
					return
				}, 1)
				__e.Return(reg36640)
				return
			}, 1)
			__e.Return(reg36639)
			return
		}, 1)
		reg36642 := PrimCons(reg36637, reg36638)
		reg36643 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36642)
		_ = reg36643
		reg36644 := MakeSymbol("shen.type-signature-of-==")
		reg36645 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V3209 := __args[0]
			_ = V3209
			reg36646 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3210 := __args[0]
				_ = V3210
				reg36647 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V3211 := __args[0]
					_ = V3211
					__e.TailApply(__defun__shen_4type_1signature_1of_1_a_a, V3209, V3210, V3211)
					return
				}, 1)
				__e.Return(reg36647)
				return
			}, 1)
			__e.Return(reg36646)
			return
		}, 1)
		reg36649 := PrimCons(reg36644, reg36645)
		__e.TailApply(__defun__shen_4set_1lambda_1form_1entry, reg36649)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.initialise-signedfunc-lambda-forms", value: __defun__shen_4initialise_1signedfunc_1lambda_1forms})

	__defun__shen_4initialise_1lambda_1forms = MakeNative(func(__e Evaluator, __args ...Obj) {
		reg36651 := MakeSymbol("shen.datatype-error")
		reg36652 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4datatype_1error, X)
			return
		}, 1)
		reg36654 := PrimCons(reg36651, reg36652)
		reg36655 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36654)
		_ = reg36655
		reg36656 := MakeSymbol("shen.tuple")
		reg36657 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4tuple, X)
			return
		}, 1)
		reg36659 := PrimCons(reg36656, reg36657)
		reg36660 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36659)
		_ = reg36660
		reg36661 := MakeSymbol("shen.pvar")
		reg36662 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4pvar, X)
			return
		}, 1)
		reg36664 := PrimCons(reg36661, reg36662)
		reg36665 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36664)
		_ = reg36665
		reg36666 := MakeSymbol("shen.dictionary")
		reg36667 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4dictionary, X)
			return
		}, 1)
		reg36669 := PrimCons(reg36666, reg36667)
		reg36670 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36669)
		_ = reg36670
		reg36671 := MakeSymbol("@v")
		reg36672 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V668 := __args[0]
			_ = V668
			reg36673 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V669 := __args[0]
				_ = V669
				__e.TailApply(__defun___8v, V668, V669)
				return
			}, 1)
			__e.Return(reg36673)
			return
		}, 1)
		reg36675 := PrimCons(reg36671, reg36672)
		reg36676 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36675)
		_ = reg36676
		reg36677 := MakeSymbol("@p")
		reg36678 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V666 := __args[0]
			_ = V666
			reg36679 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V667 := __args[0]
				_ = V667
				__e.TailApply(__defun___8p, V666, V667)
				return
			}, 1)
			__e.Return(reg36679)
			return
		}, 1)
		reg36681 := PrimCons(reg36677, reg36678)
		reg36682 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36681)
		_ = reg36682
		reg36683 := MakeSymbol("@s")
		reg36684 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V664 := __args[0]
			_ = V664
			reg36685 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V665 := __args[0]
				_ = V665
				__e.TailApply(__defun___8s, V664, V665)
				return
			}, 1)
			__e.Return(reg36685)
			return
		}, 1)
		reg36687 := PrimCons(reg36683, reg36684)
		reg36688 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36687)
		_ = reg36688
		reg36689 := MakeSymbol("<e>")
		reg36690 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V663 := __args[0]
			_ = V663
			__e.TailApply(__defun___5e_6, V663)
			return
		}, 1)
		reg36692 := PrimCons(reg36689, reg36690)
		reg36693 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36692)
		_ = reg36693
		reg36694 := MakeSymbol("<!>")
		reg36695 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V662 := __args[0]
			_ = V662
			__e.TailApply(__defun___5_b_6, V662)
			return
		}, 1)
		reg36697 := PrimCons(reg36694, reg36695)
		reg36698 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36697)
		_ = reg36698
		reg36699 := MakeSymbol("==")
		reg36700 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V660 := __args[0]
			_ = V660
			reg36701 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V661 := __args[0]
				_ = V661
				__e.TailApply(__defun___a_a, V660, V661)
				return
			}, 1)
			__e.Return(reg36701)
			return
		}, 1)
		reg36703 := PrimCons(reg36699, reg36700)
		reg36704 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36703)
		_ = reg36704
		reg36705 := MakeSymbol("=")
		reg36706 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V658 := __args[0]
			_ = V658
			reg36707 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V659 := __args[0]
				_ = V659
				reg36708 := PrimEqual(V658, V659)
				__e.Return(reg36708)
				return
			}, 1)
			__e.Return(reg36707)
			return
		}, 1)
		reg36709 := PrimCons(reg36705, reg36706)
		reg36710 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36709)
		_ = reg36710
		reg36711 := MakeSymbol(">=")
		reg36712 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V656 := __args[0]
			_ = V656
			reg36713 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V657 := __args[0]
				_ = V657
				reg36714 := PrimGreatEqual(V656, V657)
				__e.Return(reg36714)
				return
			}, 1)
			__e.Return(reg36713)
			return
		}, 1)
		reg36715 := PrimCons(reg36711, reg36712)
		reg36716 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36715)
		_ = reg36716
		reg36717 := MakeSymbol(">")
		reg36718 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V654 := __args[0]
			_ = V654
			reg36719 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V655 := __args[0]
				_ = V655
				reg36720 := PrimGreatThan(V654, V655)
				__e.Return(reg36720)
				return
			}, 1)
			__e.Return(reg36719)
			return
		}, 1)
		reg36721 := PrimCons(reg36717, reg36718)
		reg36722 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36721)
		_ = reg36722
		reg36723 := MakeSymbol("-")
		reg36724 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V652 := __args[0]
			_ = V652
			reg36725 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V653 := __args[0]
				_ = V653
				reg36726 := PrimNumberSubtract(V652, V653)
				__e.Return(reg36726)
				return
			}, 1)
			__e.Return(reg36725)
			return
		}, 1)
		reg36727 := PrimCons(reg36723, reg36724)
		reg36728 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36727)
		_ = reg36728
		reg36729 := MakeSymbol("/")
		reg36730 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V650 := __args[0]
			_ = V650
			reg36731 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V651 := __args[0]
				_ = V651
				reg36732 := PrimNumberDivide(V650, V651)
				__e.Return(reg36732)
				return
			}, 1)
			__e.Return(reg36731)
			return
		}, 1)
		reg36733 := PrimCons(reg36729, reg36730)
		reg36734 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36733)
		_ = reg36734
		reg36735 := MakeSymbol("*")
		reg36736 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V648 := __args[0]
			_ = V648
			reg36737 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V649 := __args[0]
				_ = V649
				reg36738 := PrimNumberMultiply(V648, V649)
				__e.Return(reg36738)
				return
			}, 1)
			__e.Return(reg36737)
			return
		}, 1)
		reg36739 := PrimCons(reg36735, reg36736)
		reg36740 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36739)
		_ = reg36740
		reg36741 := MakeSymbol("+")
		reg36742 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V646 := __args[0]
			_ = V646
			reg36743 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V647 := __args[0]
				_ = V647
				reg36744 := PrimNumberAdd(V646, V647)
				__e.Return(reg36744)
				return
			}, 1)
			__e.Return(reg36743)
			return
		}, 1)
		reg36745 := PrimCons(reg36741, reg36742)
		reg36746 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36745)
		_ = reg36746
		reg36747 := MakeSymbol("<=")
		reg36748 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V644 := __args[0]
			_ = V644
			reg36749 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V645 := __args[0]
				_ = V645
				reg36750 := PrimLessEqual(V644, V645)
				__e.Return(reg36750)
				return
			}, 1)
			__e.Return(reg36749)
			return
		}, 1)
		reg36751 := PrimCons(reg36747, reg36748)
		reg36752 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36751)
		_ = reg36752
		reg36753 := MakeSymbol("<")
		reg36754 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V642 := __args[0]
			_ = V642
			reg36755 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V643 := __args[0]
				_ = V643
				reg36756 := PrimLessThan(V642, V643)
				__e.Return(reg36756)
				return
			}, 1)
			__e.Return(reg36755)
			return
		}, 1)
		reg36757 := PrimCons(reg36753, reg36754)
		reg36758 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36757)
		_ = reg36758
		reg36759 := MakeSymbol("y-or-n?")
		reg36760 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V641 := __args[0]
			_ = V641
			__e.TailApply(__defun__y_1or_1n_2, V641)
			return
		}, 1)
		reg36762 := PrimCons(reg36759, reg36760)
		reg36763 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36762)
		_ = reg36763
		reg36764 := MakeSymbol("write-to-file")
		reg36765 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V639 := __args[0]
			_ = V639
			reg36766 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V640 := __args[0]
				_ = V640
				__e.TailApply(__defun__write_1to_1file, V639, V640)
				return
			}, 1)
			__e.Return(reg36766)
			return
		}, 1)
		reg36768 := PrimCons(reg36764, reg36765)
		reg36769 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36768)
		_ = reg36769
		reg36770 := MakeSymbol("write-byte")
		reg36771 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V637 := __args[0]
			_ = V637
			reg36772 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V638 := __args[0]
				_ = V638
				reg36773 := PrimWriteByte(V637, V638)
				__e.Return(reg36773)
				return
			}, 1)
			__e.Return(reg36772)
			return
		}, 1)
		reg36774 := PrimCons(reg36770, reg36771)
		reg36775 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36774)
		_ = reg36775
		reg36776 := MakeSymbol("variable?")
		reg36777 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V636 := __args[0]
			_ = V636
			reg36778 := PrimIsVariable(V636)
			__e.Return(reg36778)
			return
		}, 1)
		reg36779 := PrimCons(reg36776, reg36777)
		reg36780 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36779)
		_ = reg36780
		reg36781 := MakeSymbol("value")
		reg36782 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V635 := __args[0]
			_ = V635
			reg36783 := PrimValue(V635)
			__e.Return(reg36783)
			return
		}, 1)
		reg36784 := PrimCons(reg36781, reg36782)
		reg36785 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36784)
		_ = reg36785
		reg36786 := MakeSymbol("vector->")
		reg36787 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V632 := __args[0]
			_ = V632
			reg36788 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V633 := __args[0]
				_ = V633
				reg36789 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V634 := __args[0]
					_ = V634
					__e.TailApply(__defun__vector_1_6, V632, V633, V634)
					return
				}, 1)
				__e.Return(reg36789)
				return
			}, 1)
			__e.Return(reg36788)
			return
		}, 1)
		reg36791 := PrimCons(reg36786, reg36787)
		reg36792 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36791)
		_ = reg36792
		reg36793 := MakeSymbol("<-vector")
		reg36794 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V630 := __args[0]
			_ = V630
			reg36795 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V631 := __args[0]
				_ = V631
				__e.TailApply(__defun___5_1vector, V630, V631)
				return
			}, 1)
			__e.Return(reg36795)
			return
		}, 1)
		reg36797 := PrimCons(reg36793, reg36794)
		reg36798 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36797)
		_ = reg36798
		reg36799 := MakeSymbol("vector")
		reg36800 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V629 := __args[0]
			_ = V629
			__e.TailApply(__defun__vector, V629)
			return
		}, 1)
		reg36802 := PrimCons(reg36799, reg36800)
		reg36803 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36802)
		_ = reg36803
		reg36804 := MakeSymbol("vector?")
		reg36805 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V628 := __args[0]
			_ = V628
			__e.TailApply(__defun__vector_2, V628)
			return
		}, 1)
		reg36807 := PrimCons(reg36804, reg36805)
		reg36808 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36807)
		_ = reg36808
		reg36809 := MakeSymbol("unspecialise")
		reg36810 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V627 := __args[0]
			_ = V627
			__e.TailApply(__defun__unspecialise, V627)
			return
		}, 1)
		reg36812 := PrimCons(reg36809, reg36810)
		reg36813 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36812)
		_ = reg36813
		reg36814 := MakeSymbol("untrack")
		reg36815 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V626 := __args[0]
			_ = V626
			__e.TailApply(__defun__untrack, V626)
			return
		}, 1)
		reg36817 := PrimCons(reg36814, reg36815)
		reg36818 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36817)
		_ = reg36818
		reg36819 := MakeSymbol("union")
		reg36820 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V624 := __args[0]
			_ = V624
			reg36821 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V625 := __args[0]
				_ = V625
				__e.TailApply(__defun__union, V624, V625)
				return
			}, 1)
			__e.Return(reg36821)
			return
		}, 1)
		reg36823 := PrimCons(reg36819, reg36820)
		reg36824 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36823)
		_ = reg36824
		reg36825 := MakeSymbol("unify")
		reg36826 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V620 := __args[0]
			_ = V620
			reg36827 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V621 := __args[0]
				_ = V621
				reg36828 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V622 := __args[0]
					_ = V622
					reg36829 := MakeNative(func(__e Evaluator, __args ...Obj) {
						V623 := __args[0]
						_ = V623
						__e.TailApply(__defun__unify, V620, V621, V622, V623)
						return
					}, 1)
					__e.Return(reg36829)
					return
				}, 1)
				__e.Return(reg36828)
				return
			}, 1)
			__e.Return(reg36827)
			return
		}, 1)
		reg36831 := PrimCons(reg36825, reg36826)
		reg36832 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36831)
		_ = reg36832
		reg36833 := MakeSymbol("unify!")
		reg36834 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V616 := __args[0]
			_ = V616
			reg36835 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V617 := __args[0]
				_ = V617
				reg36836 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V618 := __args[0]
					_ = V618
					reg36837 := MakeNative(func(__e Evaluator, __args ...Obj) {
						V619 := __args[0]
						_ = V619
						__e.TailApply(__defun__unify_b, V616, V617, V618, V619)
						return
					}, 1)
					__e.Return(reg36837)
					return
				}, 1)
				__e.Return(reg36836)
				return
			}, 1)
			__e.Return(reg36835)
			return
		}, 1)
		reg36839 := PrimCons(reg36833, reg36834)
		reg36840 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36839)
		_ = reg36840
		reg36841 := MakeSymbol("unput")
		reg36842 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V613 := __args[0]
			_ = V613
			reg36843 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V614 := __args[0]
				_ = V614
				reg36844 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V615 := __args[0]
					_ = V615
					__e.TailApply(__defun__unput, V613, V614, V615)
					return
				}, 1)
				__e.Return(reg36844)
				return
			}, 1)
			__e.Return(reg36843)
			return
		}, 1)
		reg36846 := PrimCons(reg36841, reg36842)
		reg36847 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36846)
		_ = reg36847
		reg36848 := MakeSymbol("unprofile")
		reg36849 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V612 := __args[0]
			_ = V612
			__e.TailApply(__defun__unprofile, V612)
			return
		}, 1)
		reg36851 := PrimCons(reg36848, reg36849)
		reg36852 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36851)
		_ = reg36852
		reg36853 := MakeSymbol("undefmacro")
		reg36854 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V611 := __args[0]
			_ = V611
			__e.TailApply(__defun__undefmacro, V611)
			return
		}, 1)
		reg36856 := PrimCons(reg36853, reg36854)
		reg36857 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36856)
		_ = reg36857
		reg36858 := MakeSymbol("return")
		reg36859 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V608 := __args[0]
			_ = V608
			reg36860 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V609 := __args[0]
				_ = V609
				reg36861 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V610 := __args[0]
					_ = V610
					__e.TailApply(__defun__return, V608, V609, V610)
					return
				}, 1)
				__e.Return(reg36861)
				return
			}, 1)
			__e.Return(reg36860)
			return
		}, 1)
		reg36863 := PrimCons(reg36858, reg36859)
		reg36864 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36863)
		_ = reg36864
		reg36865 := MakeSymbol("type")
		reg36866 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V606 := __args[0]
			_ = V606
			reg36867 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V607 := __args[0]
				_ = V607
				reg36868 := PrimTypeFunc(V606, V607)
				__e.Return(reg36868)
				return
			}, 1)
			__e.Return(reg36867)
			return
		}, 1)
		reg36869 := PrimCons(reg36865, reg36866)
		reg36870 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36869)
		_ = reg36870
		reg36871 := MakeSymbol("tuple?")
		reg36872 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V605 := __args[0]
			_ = V605
			__e.TailApply(__defun__tuple_2, V605)
			return
		}, 1)
		reg36874 := PrimCons(reg36871, reg36872)
		reg36875 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36874)
		_ = reg36875
		reg36876 := MakeSymbol("trap-error")
		reg36877 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V603 := __args[0]
			_ = V603
			reg36878 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V604 := __args[0]
				_ = V604
				reg36879 := MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.Return(V603)
					return
				}, 0)
				reg36880 := Try(__e, reg36879).Catch(V604)
				__e.Return(reg36880)
				return
			}, 1)
			__e.Return(reg36878)
			return
		}, 1)
		reg36881 := PrimCons(reg36876, reg36877)
		reg36882 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36881)
		_ = reg36882
		reg36883 := MakeSymbol("track")
		reg36884 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V602 := __args[0]
			_ = V602
			__e.TailApply(__defun__track, V602)
			return
		}, 1)
		reg36886 := PrimCons(reg36883, reg36884)
		reg36887 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36886)
		_ = reg36887
		reg36888 := MakeSymbol("thaw")
		reg36889 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V601 := __args[0]
			_ = V601
			__e.TailApply(__defun__thaw, V601)
			return
		}, 1)
		reg36891 := PrimCons(reg36888, reg36889)
		reg36892 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36891)
		_ = reg36892
		reg36893 := MakeSymbol("tc")
		reg36894 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V600 := __args[0]
			_ = V600
			__e.TailApply(__defun__tc, V600)
			return
		}, 1)
		reg36896 := PrimCons(reg36893, reg36894)
		reg36897 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36896)
		_ = reg36897
		reg36898 := MakeSymbol("tl")
		reg36899 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V599 := __args[0]
			_ = V599
			reg36900 := PrimTail(V599)
			__e.Return(reg36900)
			return
		}, 1)
		reg36901 := PrimCons(reg36898, reg36899)
		reg36902 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36901)
		_ = reg36902
		reg36903 := MakeSymbol("tlstr")
		reg36904 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V598 := __args[0]
			_ = V598
			reg36905 := PrimTailString(V598)
			__e.Return(reg36905)
			return
		}, 1)
		reg36906 := PrimCons(reg36903, reg36904)
		reg36907 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36906)
		_ = reg36907
		reg36908 := MakeSymbol("tail")
		reg36909 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V597 := __args[0]
			_ = V597
			__e.TailApply(__defun__tail, V597)
			return
		}, 1)
		reg36911 := PrimCons(reg36908, reg36909)
		reg36912 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36911)
		_ = reg36912
		reg36913 := MakeSymbol("systemf")
		reg36914 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V596 := __args[0]
			_ = V596
			__e.TailApply(__defun__systemf, V596)
			return
		}, 1)
		reg36916 := PrimCons(reg36913, reg36914)
		reg36917 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36916)
		_ = reg36917
		reg36918 := MakeSymbol("symbol?")
		reg36919 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V595 := __args[0]
			_ = V595
			reg36920 := PrimIsSymbol(V595)
			__e.Return(reg36920)
			return
		}, 1)
		reg36921 := PrimCons(reg36918, reg36919)
		reg36922 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36921)
		_ = reg36922
		reg36923 := MakeSymbol("string->symbol")
		reg36924 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V594 := __args[0]
			_ = V594
			__e.TailApply(__defun__string_1_6symbol, V594)
			return
		}, 1)
		reg36926 := PrimCons(reg36923, reg36924)
		reg36927 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36926)
		_ = reg36927
		reg36928 := MakeSymbol("sum")
		reg36929 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V593 := __args[0]
			_ = V593
			__e.TailApply(__defun__sum, V593)
			return
		}, 1)
		reg36931 := PrimCons(reg36928, reg36929)
		reg36932 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36931)
		_ = reg36932
		reg36933 := MakeSymbol("subst")
		reg36934 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V590 := __args[0]
			_ = V590
			reg36935 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V591 := __args[0]
				_ = V591
				reg36936 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V592 := __args[0]
					_ = V592
					__e.TailApply(__defun__subst, V590, V591, V592)
					return
				}, 1)
				__e.Return(reg36936)
				return
			}, 1)
			__e.Return(reg36935)
			return
		}, 1)
		reg36938 := PrimCons(reg36933, reg36934)
		reg36939 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36938)
		_ = reg36939
		reg36940 := MakeSymbol("string?")
		reg36941 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V589 := __args[0]
			_ = V589
			reg36942 := PrimIsString(V589)
			__e.Return(reg36942)
			return
		}, 1)
		reg36943 := PrimCons(reg36940, reg36941)
		reg36944 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36943)
		_ = reg36944
		reg36945 := MakeSymbol("string->n")
		reg36946 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V588 := __args[0]
			_ = V588
			reg36947 := PrimStringToNumber(V588)
			__e.Return(reg36947)
			return
		}, 1)
		reg36948 := PrimCons(reg36945, reg36946)
		reg36949 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36948)
		_ = reg36949
		reg36950 := MakeSymbol("step")
		reg36951 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V587 := __args[0]
			_ = V587
			__e.TailApply(__defun__step, V587)
			return
		}, 1)
		reg36953 := PrimCons(reg36950, reg36951)
		reg36954 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36953)
		_ = reg36954
		reg36955 := MakeSymbol("spy")
		reg36956 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V586 := __args[0]
			_ = V586
			__e.TailApply(__defun__spy, V586)
			return
		}, 1)
		reg36958 := PrimCons(reg36955, reg36956)
		reg36959 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36958)
		_ = reg36959
		reg36960 := MakeSymbol("specialise")
		reg36961 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V585 := __args[0]
			_ = V585
			__e.TailApply(__defun__specialise, V585)
			return
		}, 1)
		reg36963 := PrimCons(reg36960, reg36961)
		reg36964 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36963)
		_ = reg36964
		reg36965 := MakeSymbol("snd")
		reg36966 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V584 := __args[0]
			_ = V584
			__e.TailApply(__defun__snd, V584)
			return
		}, 1)
		reg36968 := PrimCons(reg36965, reg36966)
		reg36969 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36968)
		_ = reg36969
		reg36970 := MakeSymbol("simple-error")
		reg36971 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V583 := __args[0]
			_ = V583
			reg36972 := PrimSimpleError(V583)
			__e.Return(reg36972)
			return
		}, 1)
		reg36973 := PrimCons(reg36970, reg36971)
		reg36974 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36973)
		_ = reg36974
		reg36975 := MakeSymbol("set")
		reg36976 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V581 := __args[0]
			_ = V581
			reg36977 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V582 := __args[0]
				_ = V582
				reg36978 := PrimSet(V581, V582)
				__e.Return(reg36978)
				return
			}, 1)
			__e.Return(reg36977)
			return
		}, 1)
		reg36979 := PrimCons(reg36975, reg36976)
		reg36980 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36979)
		_ = reg36980
		reg36981 := MakeSymbol("str")
		reg36982 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V580 := __args[0]
			_ = V580
			reg36983 := PrimStr(V580)
			__e.Return(reg36983)
			return
		}, 1)
		reg36984 := PrimCons(reg36981, reg36982)
		reg36985 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36984)
		_ = reg36985
		reg36986 := MakeSymbol("reverse")
		reg36987 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V579 := __args[0]
			_ = V579
			__e.TailApply(__defun__reverse, V579)
			return
		}, 1)
		reg36989 := PrimCons(reg36986, reg36987)
		reg36990 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36989)
		_ = reg36990
		reg36991 := MakeSymbol("remove")
		reg36992 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V577 := __args[0]
			_ = V577
			reg36993 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V578 := __args[0]
				_ = V578
				__e.TailApply(__defun__remove, V577, V578)
				return
			}, 1)
			__e.Return(reg36993)
			return
		}, 1)
		reg36995 := PrimCons(reg36991, reg36992)
		reg36996 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg36995)
		_ = reg36996
		reg36997 := MakeSymbol("read")
		reg36998 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V576 := __args[0]
			_ = V576
			__e.TailApply(__defun__read, V576)
			return
		}, 1)
		reg37000 := PrimCons(reg36997, reg36998)
		reg37001 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37000)
		_ = reg37001
		reg37002 := MakeSymbol("read-file")
		reg37003 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V575 := __args[0]
			_ = V575
			__e.TailApply(__defun__read_1file, V575)
			return
		}, 1)
		reg37005 := PrimCons(reg37002, reg37003)
		reg37006 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37005)
		_ = reg37006
		reg37007 := MakeSymbol("read-file-as-bytelist")
		reg37008 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V574 := __args[0]
			_ = V574
			reg37009 := PrimReadFileAsByteList(V574)
			__e.Return(reg37009)
			return
		}, 1)
		reg37010 := PrimCons(reg37007, reg37008)
		reg37011 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37010)
		_ = reg37011
		reg37012 := MakeSymbol("read-file-as-string")
		reg37013 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V573 := __args[0]
			_ = V573
			reg37014 := PrimReadFileAsString(V573)
			__e.Return(reg37014)
			return
		}, 1)
		reg37015 := PrimCons(reg37012, reg37013)
		reg37016 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37015)
		_ = reg37016
		reg37017 := MakeSymbol("read-byte")
		reg37018 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V572 := __args[0]
			_ = V572
			reg37019 := PrimReadByte(V572)
			__e.Return(reg37019)
			return
		}, 1)
		reg37020 := PrimCons(reg37017, reg37018)
		reg37021 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37020)
		_ = reg37021
		reg37022 := MakeSymbol("read-from-string")
		reg37023 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V571 := __args[0]
			_ = V571
			__e.TailApply(__defun__read_1from_1string, V571)
			return
		}, 1)
		reg37025 := PrimCons(reg37022, reg37023)
		reg37026 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37025)
		_ = reg37026
		reg37027 := MakeSymbol("package?")
		reg37028 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V570 := __args[0]
			_ = V570
			__e.TailApply(__defun__package_2, V570)
			return
		}, 1)
		reg37030 := PrimCons(reg37027, reg37028)
		reg37031 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37030)
		_ = reg37031
		reg37032 := MakeSymbol("put")
		reg37033 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V566 := __args[0]
			_ = V566
			reg37034 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V567 := __args[0]
				_ = V567
				reg37035 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V568 := __args[0]
					_ = V568
					reg37036 := MakeNative(func(__e Evaluator, __args ...Obj) {
						V569 := __args[0]
						_ = V569
						__e.TailApply(__defun__put, V566, V567, V568, V569)
						return
					}, 1)
					__e.Return(reg37036)
					return
				}, 1)
				__e.Return(reg37035)
				return
			}, 1)
			__e.Return(reg37034)
			return
		}, 1)
		reg37038 := PrimCons(reg37032, reg37033)
		reg37039 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37038)
		_ = reg37039
		reg37040 := MakeSymbol("preclude")
		reg37041 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V565 := __args[0]
			_ = V565
			__e.TailApply(__defun__preclude, V565)
			return
		}, 1)
		reg37043 := PrimCons(reg37040, reg37041)
		reg37044 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37043)
		_ = reg37044
		reg37045 := MakeSymbol("preclude-all-but")
		reg37046 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V564 := __args[0]
			_ = V564
			__e.TailApply(__defun__preclude_1all_1but, V564)
			return
		}, 1)
		reg37048 := PrimCons(reg37045, reg37046)
		reg37049 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37048)
		_ = reg37049
		reg37050 := MakeSymbol("ps")
		reg37051 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V563 := __args[0]
			_ = V563
			__e.TailApply(__defun__ps, V563)
			return
		}, 1)
		reg37053 := PrimCons(reg37050, reg37051)
		reg37054 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37053)
		_ = reg37054
		reg37055 := MakeSymbol("protect")
		reg37056 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V562 := __args[0]
			_ = V562
			__e.TailApply(__defun__protect, V562)
			return
		}, 1)
		reg37058 := PrimCons(reg37055, reg37056)
		reg37059 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37058)
		_ = reg37059
		reg37060 := MakeSymbol("profile-results")
		reg37061 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V561 := __args[0]
			_ = V561
			__e.TailApply(__defun__profile_1results, V561)
			return
		}, 1)
		reg37063 := PrimCons(reg37060, reg37061)
		reg37064 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37063)
		_ = reg37064
		reg37065 := MakeSymbol("profile")
		reg37066 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V560 := __args[0]
			_ = V560
			__e.TailApply(__defun__profile, V560)
			return
		}, 1)
		reg37068 := PrimCons(reg37065, reg37066)
		reg37069 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37068)
		_ = reg37069
		reg37070 := MakeSymbol("print")
		reg37071 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V559 := __args[0]
			_ = V559
			__e.TailApply(__defun__print, V559)
			return
		}, 1)
		reg37073 := PrimCons(reg37070, reg37071)
		reg37074 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37073)
		_ = reg37074
		reg37075 := MakeSymbol("pr")
		reg37076 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V557 := __args[0]
			_ = V557
			reg37077 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V558 := __args[0]
				_ = V558
				__e.TailApply(__defun__pr, V557, V558)
				return
			}, 1)
			__e.Return(reg37077)
			return
		}, 1)
		reg37079 := PrimCons(reg37075, reg37076)
		reg37080 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37079)
		_ = reg37080
		reg37081 := MakeSymbol("pos")
		reg37082 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V555 := __args[0]
			_ = V555
			reg37083 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V556 := __args[0]
				_ = V556
				reg37084 := PrimPos(V555, V556)
				__e.Return(reg37084)
				return
			}, 1)
			__e.Return(reg37083)
			return
		}, 1)
		reg37085 := PrimCons(reg37081, reg37082)
		reg37086 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37085)
		_ = reg37086
		reg37087 := MakeSymbol("or")
		reg37088 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V553 := __args[0]
			_ = V553
			reg37089 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V554 := __args[0]
				_ = V554
				if V553 == True {
					reg37090 := True
					__e.Return(reg37090)
					return
				} else {
					if V554 == True {
						reg37091 := True
						__e.Return(reg37091)
						return
					} else {
						reg37092 := False
						__e.Return(reg37092)
						return
					}
				}
			}, 1)
			__e.Return(reg37089)
			return
		}, 1)
		reg37093 := PrimCons(reg37087, reg37088)
		reg37094 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37093)
		_ = reg37094
		reg37095 := MakeSymbol("optimise")
		reg37096 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V552 := __args[0]
			_ = V552
			__e.TailApply(__defun__optimise, V552)
			return
		}, 1)
		reg37098 := PrimCons(reg37095, reg37096)
		reg37099 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37098)
		_ = reg37099
		reg37100 := MakeSymbol("open")
		reg37101 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V550 := __args[0]
			_ = V550
			reg37102 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V551 := __args[0]
				_ = V551
				reg37103 := PrimOpenStream(V550, V551)
				__e.Return(reg37103)
				return
			}, 1)
			__e.Return(reg37102)
			return
		}, 1)
		reg37104 := PrimCons(reg37100, reg37101)
		reg37105 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37104)
		_ = reg37105
		reg37106 := MakeSymbol("occurrences")
		reg37107 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V548 := __args[0]
			_ = V548
			reg37108 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V549 := __args[0]
				_ = V549
				__e.TailApply(__defun__occurrences, V548, V549)
				return
			}, 1)
			__e.Return(reg37108)
			return
		}, 1)
		reg37110 := PrimCons(reg37106, reg37107)
		reg37111 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37110)
		_ = reg37111
		reg37112 := MakeSymbol("occurs-check")
		reg37113 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V547 := __args[0]
			_ = V547
			__e.TailApply(__defun__occurs_1check, V547)
			return
		}, 1)
		reg37115 := PrimCons(reg37112, reg37113)
		reg37116 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37115)
		_ = reg37116
		reg37117 := MakeSymbol("n->string")
		reg37118 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V546 := __args[0]
			_ = V546
			reg37119 := PrimNumberToString(V546)
			__e.Return(reg37119)
			return
		}, 1)
		reg37120 := PrimCons(reg37117, reg37118)
		reg37121 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37120)
		_ = reg37121
		reg37122 := MakeSymbol("number?")
		reg37123 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V545 := __args[0]
			_ = V545
			reg37124 := PrimIsNumber(V545)
			__e.Return(reg37124)
			return
		}, 1)
		reg37125 := PrimCons(reg37122, reg37123)
		reg37126 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37125)
		_ = reg37126
		reg37127 := MakeSymbol("nth")
		reg37128 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V543 := __args[0]
			_ = V543
			reg37129 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V544 := __args[0]
				_ = V544
				__e.TailApply(__defun__nth, V543, V544)
				return
			}, 1)
			__e.Return(reg37129)
			return
		}, 1)
		reg37131 := PrimCons(reg37127, reg37128)
		reg37132 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37131)
		_ = reg37132
		reg37133 := MakeSymbol("not")
		reg37134 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V542 := __args[0]
			_ = V542
			reg37135 := PrimNot(V542)
			__e.Return(reg37135)
			return
		}, 1)
		reg37136 := PrimCons(reg37133, reg37134)
		reg37137 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37136)
		_ = reg37137
		reg37138 := MakeSymbol("nl")
		reg37139 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V541 := __args[0]
			_ = V541
			__e.TailApply(__defun__nl, V541)
			return
		}, 1)
		reg37141 := PrimCons(reg37138, reg37139)
		reg37142 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37141)
		_ = reg37142
		reg37143 := MakeSymbol("macroexpand")
		reg37144 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V540 := __args[0]
			_ = V540
			__e.TailApply(__defun__macroexpand, V540)
			return
		}, 1)
		reg37146 := PrimCons(reg37143, reg37144)
		reg37147 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37146)
		_ = reg37147
		reg37148 := MakeSymbol("maxinferences")
		reg37149 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V539 := __args[0]
			_ = V539
			__e.TailApply(__defun__maxinferences, V539)
			return
		}, 1)
		reg37151 := PrimCons(reg37148, reg37149)
		reg37152 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37151)
		_ = reg37152
		reg37153 := MakeSymbol("mapcan")
		reg37154 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V537 := __args[0]
			_ = V537
			reg37155 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V538 := __args[0]
				_ = V538
				__e.TailApply(__defun__mapcan, V537, V538)
				return
			}, 1)
			__e.Return(reg37155)
			return
		}, 1)
		reg37157 := PrimCons(reg37153, reg37154)
		reg37158 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37157)
		_ = reg37158
		reg37159 := MakeSymbol("map")
		reg37160 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V535 := __args[0]
			_ = V535
			reg37161 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V536 := __args[0]
				_ = V536
				__e.TailApply(__defun__map, V535, V536)
				return
			}, 1)
			__e.Return(reg37161)
			return
		}, 1)
		reg37163 := PrimCons(reg37159, reg37160)
		reg37164 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37163)
		_ = reg37164
		reg37165 := MakeSymbol("load")
		reg37166 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V534 := __args[0]
			_ = V534
			__e.TailApply(__defun__load, V534)
			return
		}, 1)
		reg37168 := PrimCons(reg37165, reg37166)
		reg37169 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37168)
		_ = reg37169
		reg37170 := MakeSymbol("lineread")
		reg37171 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V533 := __args[0]
			_ = V533
			__e.TailApply(__defun__lineread, V533)
			return
		}, 1)
		reg37173 := PrimCons(reg37170, reg37171)
		reg37174 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37173)
		_ = reg37174
		reg37175 := MakeSymbol("limit")
		reg37176 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V532 := __args[0]
			_ = V532
			__e.TailApply(__defun__limit, V532)
			return
		}, 1)
		reg37178 := PrimCons(reg37175, reg37176)
		reg37179 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37178)
		_ = reg37179
		reg37180 := MakeSymbol("length")
		reg37181 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V531 := __args[0]
			_ = V531
			__e.TailApply(__defun__length, V531)
			return
		}, 1)
		reg37183 := PrimCons(reg37180, reg37181)
		reg37184 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37183)
		_ = reg37184
		reg37185 := MakeSymbol("intersection")
		reg37186 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V529 := __args[0]
			_ = V529
			reg37187 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V530 := __args[0]
				_ = V530
				__e.TailApply(__defun__intersection, V529, V530)
				return
			}, 1)
			__e.Return(reg37187)
			return
		}, 1)
		reg37189 := PrimCons(reg37185, reg37186)
		reg37190 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37189)
		_ = reg37190
		reg37191 := MakeSymbol("intern")
		reg37192 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V528 := __args[0]
			_ = V528
			reg37193 := PrimIntern(V528)
			__e.Return(reg37193)
			return
		}, 1)
		reg37194 := PrimCons(reg37191, reg37192)
		reg37195 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37194)
		_ = reg37195
		reg37196 := MakeSymbol("integer?")
		reg37197 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V527 := __args[0]
			_ = V527
			reg37198 := PrimIsInteger(V527)
			__e.Return(reg37198)
			return
		}, 1)
		reg37199 := PrimCons(reg37196, reg37197)
		reg37200 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37199)
		_ = reg37200
		reg37201 := MakeSymbol("input")
		reg37202 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V526 := __args[0]
			_ = V526
			__e.TailApply(__defun__input, V526)
			return
		}, 1)
		reg37204 := PrimCons(reg37201, reg37202)
		reg37205 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37204)
		_ = reg37205
		reg37206 := MakeSymbol("input+")
		reg37207 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V524 := __args[0]
			_ = V524
			reg37208 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V525 := __args[0]
				_ = V525
				__e.TailApply(__defun__input_7, V524, V525)
				return
			}, 1)
			__e.Return(reg37208)
			return
		}, 1)
		reg37210 := PrimCons(reg37206, reg37207)
		reg37211 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37210)
		_ = reg37211
		reg37212 := MakeSymbol("include")
		reg37213 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V523 := __args[0]
			_ = V523
			__e.TailApply(__defun__include, V523)
			return
		}, 1)
		reg37215 := PrimCons(reg37212, reg37213)
		reg37216 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37215)
		_ = reg37216
		reg37217 := MakeSymbol("include-all-but")
		reg37218 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V522 := __args[0]
			_ = V522
			__e.TailApply(__defun__include_1all_1but, V522)
			return
		}, 1)
		reg37220 := PrimCons(reg37217, reg37218)
		reg37221 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37220)
		_ = reg37221
		reg37222 := MakeSymbol("internal")
		reg37223 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V521 := __args[0]
			_ = V521
			__e.TailApply(__defun__internal, V521)
			return
		}, 1)
		reg37225 := PrimCons(reg37222, reg37223)
		reg37226 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37225)
		_ = reg37226
		reg37227 := MakeSymbol("if")
		reg37228 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V518 := __args[0]
			_ = V518
			reg37229 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V519 := __args[0]
				_ = V519
				reg37230 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V520 := __args[0]
					_ = V520
					if V518 == True {
						__e.Return(V519)
						return
					} else {
						__e.Return(V520)
						return
					}
				}, 1)
				__e.Return(reg37230)
				return
			}, 1)
			__e.Return(reg37229)
			return
		}, 1)
		reg37231 := PrimCons(reg37227, reg37228)
		reg37232 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37231)
		_ = reg37232
		reg37233 := MakeSymbol("identical")
		reg37234 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V514 := __args[0]
			_ = V514
			reg37235 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V515 := __args[0]
				_ = V515
				reg37236 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V516 := __args[0]
					_ = V516
					reg37237 := MakeNative(func(__e Evaluator, __args ...Obj) {
						V517 := __args[0]
						_ = V517
						__e.TailApply(__defun__identical, V514, V515, V516, V517)
						return
					}, 1)
					__e.Return(reg37237)
					return
				}, 1)
				__e.Return(reg37236)
				return
			}, 1)
			__e.Return(reg37235)
			return
		}, 1)
		reg37239 := PrimCons(reg37233, reg37234)
		reg37240 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37239)
		_ = reg37240
		reg37241 := MakeSymbol("head")
		reg37242 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V513 := __args[0]
			_ = V513
			__e.TailApply(__defun__head, V513)
			return
		}, 1)
		reg37244 := PrimCons(reg37241, reg37242)
		reg37245 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37244)
		_ = reg37245
		reg37246 := MakeSymbol("hd")
		reg37247 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V512 := __args[0]
			_ = V512
			reg37248 := PrimHead(V512)
			__e.Return(reg37248)
			return
		}, 1)
		reg37249 := PrimCons(reg37246, reg37247)
		reg37250 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37249)
		_ = reg37250
		reg37251 := MakeSymbol("hdv")
		reg37252 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V511 := __args[0]
			_ = V511
			__e.TailApply(__defun__hdv, V511)
			return
		}, 1)
		reg37254 := PrimCons(reg37251, reg37252)
		reg37255 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37254)
		_ = reg37255
		reg37256 := MakeSymbol("hdstr")
		reg37257 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V510 := __args[0]
			_ = V510
			__e.TailApply(__defun__hdstr, V510)
			return
		}, 1)
		reg37259 := PrimCons(reg37256, reg37257)
		reg37260 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37259)
		_ = reg37260
		reg37261 := MakeSymbol("hash")
		reg37262 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V508 := __args[0]
			_ = V508
			reg37263 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V509 := __args[0]
				_ = V509
				__e.TailApply(__defun__hash, V508, V509)
				return
			}, 1)
			__e.Return(reg37263)
			return
		}, 1)
		reg37265 := PrimCons(reg37261, reg37262)
		reg37266 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37265)
		_ = reg37266
		reg37267 := MakeSymbol("get")
		reg37268 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V505 := __args[0]
			_ = V505
			reg37269 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V506 := __args[0]
				_ = V506
				reg37270 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V507 := __args[0]
					_ = V507
					__e.TailApply(__defun__get, V505, V506, V507)
					return
				}, 1)
				__e.Return(reg37270)
				return
			}, 1)
			__e.Return(reg37269)
			return
		}, 1)
		reg37272 := PrimCons(reg37267, reg37268)
		reg37273 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37272)
		_ = reg37273
		reg37274 := MakeSymbol("get-time")
		reg37275 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V504 := __args[0]
			_ = V504
			reg37276 := PrimGetTime(V504)
			__e.Return(reg37276)
			return
		}, 1)
		reg37277 := PrimCons(reg37274, reg37275)
		reg37278 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37277)
		_ = reg37278
		reg37279 := MakeSymbol("gensym")
		reg37280 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V503 := __args[0]
			_ = V503
			__e.TailApply(__defun__gensym, V503)
			return
		}, 1)
		reg37282 := PrimCons(reg37279, reg37280)
		reg37283 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37282)
		_ = reg37283
		reg37284 := MakeSymbol("fst")
		reg37285 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V502 := __args[0]
			_ = V502
			__e.TailApply(__defun__fst, V502)
			return
		}, 1)
		reg37287 := PrimCons(reg37284, reg37285)
		reg37288 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37287)
		_ = reg37288
		reg37289 := MakeSymbol("freeze")
		reg37290 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V501 := __args[0]
			_ = V501
			reg37291 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__e.Return(V501)
				return
			}, 0)
			__e.Return(reg37291)
			return
		}, 1)
		reg37292 := PrimCons(reg37289, reg37290)
		reg37293 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37292)
		_ = reg37293
		reg37294 := MakeSymbol("fix")
		reg37295 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V499 := __args[0]
			_ = V499
			reg37296 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V500 := __args[0]
				_ = V500
				__e.TailApply(__defun__fix, V499, V500)
				return
			}, 1)
			__e.Return(reg37296)
			return
		}, 1)
		reg37298 := PrimCons(reg37294, reg37295)
		reg37299 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37298)
		_ = reg37299
		reg37300 := MakeSymbol("fail-if")
		reg37301 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V497 := __args[0]
			_ = V497
			reg37302 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V498 := __args[0]
				_ = V498
				__e.TailApply(__defun__fail_1if, V497, V498)
				return
			}, 1)
			__e.Return(reg37302)
			return
		}, 1)
		reg37304 := PrimCons(reg37300, reg37301)
		reg37305 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37304)
		_ = reg37305
		reg37306 := MakeSymbol("findall")
		reg37307 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V492 := __args[0]
			_ = V492
			reg37308 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V493 := __args[0]
				_ = V493
				reg37309 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V494 := __args[0]
					_ = V494
					reg37310 := MakeNative(func(__e Evaluator, __args ...Obj) {
						V495 := __args[0]
						_ = V495
						reg37311 := MakeNative(func(__e Evaluator, __args ...Obj) {
							V496 := __args[0]
							_ = V496
							__e.TailApply(__defun__findall, V492, V493, V494, V495, V496)
							return
						}, 1)
						__e.Return(reg37311)
						return
					}, 1)
					__e.Return(reg37310)
					return
				}, 1)
				__e.Return(reg37309)
				return
			}, 1)
			__e.Return(reg37308)
			return
		}, 1)
		reg37313 := PrimCons(reg37306, reg37307)
		reg37314 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37313)
		_ = reg37314
		reg37315 := MakeSymbol("enable-type-theory")
		reg37316 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V491 := __args[0]
			_ = V491
			__e.TailApply(__defun__enable_1type_1theory, V491)
			return
		}, 1)
		reg37318 := PrimCons(reg37315, reg37316)
		reg37319 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37318)
		_ = reg37319
		reg37320 := MakeSymbol("explode")
		reg37321 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V490 := __args[0]
			_ = V490
			__e.TailApply(__defun__explode, V490)
			return
		}, 1)
		reg37323 := PrimCons(reg37320, reg37321)
		reg37324 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37323)
		_ = reg37324
		reg37325 := MakeSymbol("external")
		reg37326 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V489 := __args[0]
			_ = V489
			__e.TailApply(__defun__external, V489)
			return
		}, 1)
		reg37328 := PrimCons(reg37325, reg37326)
		reg37329 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37328)
		_ = reg37329
		reg37330 := MakeSymbol("eval-kl")
		reg37331 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V488 := __args[0]
			_ = V488
			reg37332 := PrimEvalKL(__e, V488)
			__e.Return(reg37332)
			return
		}, 1)
		reg37333 := PrimCons(reg37330, reg37331)
		reg37334 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37333)
		_ = reg37334
		reg37335 := MakeSymbol("eval")
		reg37336 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V487 := __args[0]
			_ = V487
			__e.TailApply(__defun__eval, V487)
			return
		}, 1)
		reg37338 := PrimCons(reg37335, reg37336)
		reg37339 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37338)
		_ = reg37339
		reg37340 := MakeSymbol("error-to-string")
		reg37341 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V486 := __args[0]
			_ = V486
			reg37342 := PrimErrorToString(V486)
			__e.Return(reg37342)
			return
		}, 1)
		reg37343 := PrimCons(reg37340, reg37341)
		reg37344 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37343)
		_ = reg37344
		reg37345 := MakeSymbol("empty?")
		reg37346 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V485 := __args[0]
			_ = V485
			__e.TailApply(__defun__empty_2, V485)
			return
		}, 1)
		reg37348 := PrimCons(reg37345, reg37346)
		reg37349 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37348)
		_ = reg37349
		reg37350 := MakeSymbol("element?")
		reg37351 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V483 := __args[0]
			_ = V483
			reg37352 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V484 := __args[0]
				_ = V484
				__e.TailApply(__defun__element_2, V483, V484)
				return
			}, 1)
			__e.Return(reg37352)
			return
		}, 1)
		reg37354 := PrimCons(reg37350, reg37351)
		reg37355 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37354)
		_ = reg37355
		reg37356 := MakeSymbol("do")
		reg37357 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V481 := __args[0]
			_ = V481
			reg37358 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V482 := __args[0]
				_ = V482
				_ = V481
				__e.Return(V482)
				return
			}, 1)
			__e.Return(reg37358)
			return
		}, 1)
		reg37359 := PrimCons(reg37356, reg37357)
		reg37360 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37359)
		_ = reg37360
		reg37361 := MakeSymbol("difference")
		reg37362 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V479 := __args[0]
			_ = V479
			reg37363 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V480 := __args[0]
				_ = V480
				__e.TailApply(__defun__difference, V479, V480)
				return
			}, 1)
			__e.Return(reg37363)
			return
		}, 1)
		reg37365 := PrimCons(reg37361, reg37362)
		reg37366 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37365)
		_ = reg37366
		reg37367 := MakeSymbol("destroy")
		reg37368 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V478 := __args[0]
			_ = V478
			__e.TailApply(__defun__destroy, V478)
			return
		}, 1)
		reg37370 := PrimCons(reg37367, reg37368)
		reg37371 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37370)
		_ = reg37371
		reg37372 := MakeSymbol("declare")
		reg37373 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V476 := __args[0]
			_ = V476
			reg37374 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V477 := __args[0]
				_ = V477
				__e.TailApply(__defun__declare, V476, V477)
				return
			}, 1)
			__e.Return(reg37374)
			return
		}, 1)
		reg37376 := PrimCons(reg37372, reg37373)
		reg37377 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37376)
		_ = reg37377
		reg37378 := MakeSymbol("cn")
		reg37379 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V474 := __args[0]
			_ = V474
			reg37380 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V475 := __args[0]
				_ = V475
				reg37381 := PrimStringConcat(V474, V475)
				__e.Return(reg37381)
				return
			}, 1)
			__e.Return(reg37380)
			return
		}, 1)
		reg37382 := PrimCons(reg37378, reg37379)
		reg37383 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37382)
		_ = reg37383
		reg37384 := MakeSymbol("cons?")
		reg37385 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V473 := __args[0]
			_ = V473
			reg37386 := PrimIsPair(V473)
			__e.Return(reg37386)
			return
		}, 1)
		reg37387 := PrimCons(reg37384, reg37385)
		reg37388 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37387)
		_ = reg37388
		reg37389 := MakeSymbol("cons")
		reg37390 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V471 := __args[0]
			_ = V471
			reg37391 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V472 := __args[0]
				_ = V472
				reg37392 := PrimCons(V471, V472)
				__e.Return(reg37392)
				return
			}, 1)
			__e.Return(reg37391)
			return
		}, 1)
		reg37393 := PrimCons(reg37389, reg37390)
		reg37394 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37393)
		_ = reg37394
		reg37395 := MakeSymbol("concat")
		reg37396 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V469 := __args[0]
			_ = V469
			reg37397 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V470 := __args[0]
				_ = V470
				__e.TailApply(__defun__concat, V469, V470)
				return
			}, 1)
			__e.Return(reg37397)
			return
		}, 1)
		reg37399 := PrimCons(reg37395, reg37396)
		reg37400 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37399)
		_ = reg37400
		reg37401 := MakeSymbol("compile")
		reg37402 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V466 := __args[0]
			_ = V466
			reg37403 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V467 := __args[0]
				_ = V467
				reg37404 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V468 := __args[0]
					_ = V468
					__e.TailApply(__defun__compile, V466, V467, V468)
					return
				}, 1)
				__e.Return(reg37404)
				return
			}, 1)
			__e.Return(reg37403)
			return
		}, 1)
		reg37406 := PrimCons(reg37401, reg37402)
		reg37407 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37406)
		_ = reg37407
		reg37408 := MakeSymbol("cd")
		reg37409 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V465 := __args[0]
			_ = V465
			__e.TailApply(__defun__cd, V465)
			return
		}, 1)
		reg37411 := PrimCons(reg37408, reg37409)
		reg37412 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37411)
		_ = reg37412
		reg37413 := MakeSymbol("close")
		reg37414 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V464 := __args[0]
			_ = V464
			reg37415 := PrimCloseStream(V464)
			__e.Return(reg37415)
			return
		}, 1)
		reg37416 := PrimCons(reg37413, reg37414)
		reg37417 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37416)
		_ = reg37417
		reg37418 := MakeSymbol("bound?")
		reg37419 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V463 := __args[0]
			_ = V463
			__e.TailApply(__defun__bound_2, V463)
			return
		}, 1)
		reg37421 := PrimCons(reg37418, reg37419)
		reg37422 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37421)
		_ = reg37422
		reg37423 := MakeSymbol("boolean?")
		reg37424 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V462 := __args[0]
			_ = V462
			__e.TailApply(__defun__boolean_2, V462)
			return
		}, 1)
		reg37426 := PrimCons(reg37423, reg37424)
		reg37427 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37426)
		_ = reg37427
		reg37428 := MakeSymbol("assoc")
		reg37429 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V460 := __args[0]
			_ = V460
			reg37430 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V461 := __args[0]
				_ = V461
				__e.TailApply(__defun__assoc, V460, V461)
				return
			}, 1)
			__e.Return(reg37430)
			return
		}, 1)
		reg37432 := PrimCons(reg37428, reg37429)
		reg37433 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37432)
		_ = reg37433
		reg37434 := MakeSymbol("arity")
		reg37435 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V459 := __args[0]
			_ = V459
			__e.TailApply(__defun__arity, V459)
			return
		}, 1)
		reg37437 := PrimCons(reg37434, reg37435)
		reg37438 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37437)
		_ = reg37438
		reg37439 := MakeSymbol("append")
		reg37440 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V457 := __args[0]
			_ = V457
			reg37441 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V458 := __args[0]
				_ = V458
				__e.TailApply(__defun__append, V457, V458)
				return
			}, 1)
			__e.Return(reg37441)
			return
		}, 1)
		reg37443 := PrimCons(reg37439, reg37440)
		reg37444 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37443)
		_ = reg37444
		reg37445 := MakeSymbol("and")
		reg37446 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V455 := __args[0]
			_ = V455
			reg37447 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V456 := __args[0]
				_ = V456
				if V455 == True {
					if V456 == True {
						reg37448 := True
						__e.Return(reg37448)
						return
					} else {
						reg37449 := False
						__e.Return(reg37449)
						return
					}
				} else {
					reg37450 := False
					__e.Return(reg37450)
					return
				}
			}, 1)
			__e.Return(reg37447)
			return
		}, 1)
		reg37451 := PrimCons(reg37445, reg37446)
		reg37452 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37451)
		_ = reg37452
		reg37453 := MakeSymbol("adjoin")
		reg37454 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V453 := __args[0]
			_ = V453
			reg37455 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V454 := __args[0]
				_ = V454
				__e.TailApply(__defun__adjoin, V453, V454)
				return
			}, 1)
			__e.Return(reg37455)
			return
		}, 1)
		reg37457 := PrimCons(reg37453, reg37454)
		reg37458 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37457)
		_ = reg37458
		reg37459 := MakeSymbol("<-address")
		reg37460 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V451 := __args[0]
			_ = V451
			reg37461 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V452 := __args[0]
				_ = V452
				reg37462 := PrimVectorGet(V451, V452)
				__e.Return(reg37462)
				return
			}, 1)
			__e.Return(reg37461)
			return
		}, 1)
		reg37463 := PrimCons(reg37459, reg37460)
		reg37464 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37463)
		_ = reg37464
		reg37465 := MakeSymbol("address->")
		reg37466 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V448 := __args[0]
			_ = V448
			reg37467 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V449 := __args[0]
				_ = V449
				reg37468 := MakeNative(func(__e Evaluator, __args ...Obj) {
					V450 := __args[0]
					_ = V450
					reg37469 := PrimVectorSet(V448, V449, V450)
					__e.Return(reg37469)
					return
				}, 1)
				__e.Return(reg37468)
				return
			}, 1)
			__e.Return(reg37467)
			return
		}, 1)
		reg37470 := PrimCons(reg37465, reg37466)
		reg37471 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37470)
		_ = reg37471
		reg37472 := MakeSymbol("absvector?")
		reg37473 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V447 := __args[0]
			_ = V447
			reg37474 := PrimIsVector(V447)
			__e.Return(reg37474)
			return
		}, 1)
		reg37475 := PrimCons(reg37472, reg37473)
		reg37476 := Call(__e, __defun__shen_4set_1lambda_1form_1entry, reg37475)
		_ = reg37476
		reg37477 := MakeSymbol("absvector")
		reg37478 := MakeNative(func(__e Evaluator, __args ...Obj) {
			V446 := __args[0]
			_ = V446
			reg37479 := PrimAbsvector(V446)
			__e.Return(reg37479)
			return
		}, 1)
		reg37480 := PrimCons(reg37477, reg37478)
		__e.TailApply(__defun__shen_4set_1lambda_1form_1entry, reg37480)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.initialise-lambda-forms", value: __defun__shen_4initialise_1lambda_1forms})

	__defun__shen_4initialise = MakeNative(func(__e Evaluator, __args ...Obj) {
		reg37482 := Call(__e, __defun__shen_4initialise_1environment)
		_ = reg37482
		reg37483 := Call(__e, __defun__shen_4initialise_1lambda_1forms)
		_ = reg37483
		reg37484 := Call(__e, __defun__shen_4initialise_1signedfunc_1lambda_1forms)
		_ = reg37484
		__e.TailApply(__defun__shen_4initialise_1signedfuncs)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.initialise", value: __defun__shen_4initialise})

}
