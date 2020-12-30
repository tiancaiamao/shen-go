package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen15409 := MakeNative(func(__e Evaluator) {
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*installing-kl*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*history*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*tc*"), False)
			gen14709 := Call(__e, ShenFunc(symshen_4dict), MakeNumber(20000))

			Call(__e, ShenFunc(symset), MakeSymbol("*property-vector*"), gen14709)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*process-counter*"), MakeNumber(0))
			gen14710 := Call(__e, ShenFunc(symvector), MakeNumber(10000))

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*varcounter*"), gen14710)

			gen14711 := Call(__e, ShenFunc(symvector), MakeNumber(10000))

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*prologvectors*"), gen14711)

			gen14712 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.Return(X)
				return
			}, 1)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*demodulation-function*"), gen14712)

			gen14713 := MakeNative(func(__e Evaluator) {
				Arg := __e.Get(1)
				_ = Arg
				__e.Return(MakeNative(func(__e Evaluator) {
					OnFail := __e.Get(1)
					_ = OnFail
					__e.TailApply(ShenFunc(symthaw), OnFail)

					return
				}, 1))
				return
			}, 1)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*custom-pattern-compiler*"), gen14713)

			gen14714 := MakeNative(func(__e Evaluator) {
				Arg := __e.Get(1)
				_ = Arg
				__e.Return(Arg)
				return
			}, 1)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*custom-pattern-reducer*"), gen14714)

			gen14715 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.function-macro"), Nil)

			gen14716 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.defprolog-macro"), gen14715)

			gen14717 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.@s-macro"), gen14716)

			gen14718 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.nl-macro"), gen14717)

			gen14719 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.synonyms-macro"), gen14718)

			gen14720 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.prolog-macro"), gen14719)

			gen14721 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.error-macro"), gen14720)

			gen14722 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.input-macro"), gen14721)

			gen14723 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.output-macro"), gen14722)

			gen14724 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.make-string-macro"), gen14723)

			gen14725 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.assoc-macro"), gen14724)

			gen14726 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.let-macro"), gen14725)

			gen14727 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.datatype-macro"), gen14726)

			gen14728 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.compile-macro"), gen14727)

			gen14729 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.put/get-macro"), gen14728)

			gen14730 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.abs-macro"), gen14729)

			gen14731 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.cases-macro"), gen14730)

			gen14732 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.timer-macro"), gen14731)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*macroreg*"), gen14732)

			gen14733 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4timer_1macro), X)

				return
			}, 1)
			gen14734 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4cases_1macro), X)

				return
			}, 1)
			gen14735 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4abs_1macro), X)

				return
			}, 1)
			gen14736 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4put_cget_1macro), X)

				return
			}, 1)
			gen14737 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4compile_1macro), X)

				return
			}, 1)
			gen14738 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4datatype_1macro), X)

				return
			}, 1)
			gen14739 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4let_1macro), X)

				return
			}, 1)
			gen14740 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4assoc_1macro), X)

				return
			}, 1)
			gen14741 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4make_1string_1macro), X)

				return
			}, 1)
			gen14742 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4output_1macro), X)

				return
			}, 1)
			gen14743 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4input_1macro), X)

				return
			}, 1)
			gen14744 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4error_1macro), X)

				return
			}, 1)
			gen14745 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4prolog_1macro), X)

				return
			}, 1)
			gen14746 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4synonyms_1macro), X)

				return
			}, 1)
			gen14747 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4nl_1macro), X)

				return
			}, 1)
			gen14748 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4_8s_1macro), X)

				return
			}, 1)
			gen14749 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4defprolog_1macro), X)

				return
			}, 1)
			gen14750 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4function_1macro), X)

				return
			}, 1)
			gen14751 := Call(__e, ShenFunc(symcons), gen14750, Nil)

			gen14752 := Call(__e, ShenFunc(symcons), gen14749, gen14751)

			gen14753 := Call(__e, ShenFunc(symcons), gen14748, gen14752)

			gen14754 := Call(__e, ShenFunc(symcons), gen14747, gen14753)

			gen14755 := Call(__e, ShenFunc(symcons), gen14746, gen14754)

			gen14756 := Call(__e, ShenFunc(symcons), gen14745, gen14755)

			gen14757 := Call(__e, ShenFunc(symcons), gen14744, gen14756)

			gen14758 := Call(__e, ShenFunc(symcons), gen14743, gen14757)

			gen14759 := Call(__e, ShenFunc(symcons), gen14742, gen14758)

			gen14760 := Call(__e, ShenFunc(symcons), gen14741, gen14759)

			gen14761 := Call(__e, ShenFunc(symcons), gen14740, gen14760)

			gen14762 := Call(__e, ShenFunc(symcons), gen14739, gen14761)

			gen14763 := Call(__e, ShenFunc(symcons), gen14738, gen14762)

			gen14764 := Call(__e, ShenFunc(symcons), gen14737, gen14763)

			gen14765 := Call(__e, ShenFunc(symcons), gen14736, gen14764)

			gen14766 := Call(__e, ShenFunc(symcons), gen14735, gen14765)

			gen14767 := Call(__e, ShenFunc(symcons), gen14734, gen14766)

			gen14768 := Call(__e, ShenFunc(symcons), gen14733, gen14767)

			Call(__e, ShenFunc(symset), MakeSymbol("*macros*"), gen14768)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*gensym*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*tracking*"), Nil)
			gen14769 := Call(__e, ShenFunc(symcons), MakeSymbol("Z"), Nil)

			gen14770 := Call(__e, ShenFunc(symcons), MakeSymbol("Y"), gen14769)

			gen14771 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), gen14770)

			gen14772 := Call(__e, ShenFunc(symcons), MakeSymbol("W"), gen14771)

			gen14773 := Call(__e, ShenFunc(symcons), MakeSymbol("V"), gen14772)

			gen14774 := Call(__e, ShenFunc(symcons), MakeSymbol("U"), gen14773)

			gen14775 := Call(__e, ShenFunc(symcons), MakeSymbol("T"), gen14774)

			gen14776 := Call(__e, ShenFunc(symcons), MakeSymbol("S"), gen14775)

			gen14777 := Call(__e, ShenFunc(symcons), MakeSymbol("R"), gen14776)

			gen14778 := Call(__e, ShenFunc(symcons), MakeSymbol("Q"), gen14777)

			gen14779 := Call(__e, ShenFunc(symcons), MakeSymbol("P"), gen14778)

			gen14780 := Call(__e, ShenFunc(symcons), MakeSymbol("O"), gen14779)

			gen14781 := Call(__e, ShenFunc(symcons), MakeSymbol("N"), gen14780)

			gen14782 := Call(__e, ShenFunc(symcons), MakeSymbol("M"), gen14781)

			gen14783 := Call(__e, ShenFunc(symcons), MakeSymbol("L"), gen14782)

			gen14784 := Call(__e, ShenFunc(symcons), MakeSymbol("K"), gen14783)

			gen14785 := Call(__e, ShenFunc(symcons), MakeSymbol("J"), gen14784)

			gen14786 := Call(__e, ShenFunc(symcons), MakeSymbol("I"), gen14785)

			gen14787 := Call(__e, ShenFunc(symcons), MakeSymbol("H"), gen14786)

			gen14788 := Call(__e, ShenFunc(symcons), MakeSymbol("G"), gen14787)

			gen14789 := Call(__e, ShenFunc(symcons), MakeSymbol("F"), gen14788)

			gen14790 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), gen14789)

			gen14791 := Call(__e, ShenFunc(symcons), MakeSymbol("D"), gen14790)

			gen14792 := Call(__e, ShenFunc(symcons), MakeSymbol("C"), gen14791)

			gen14793 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen14792)

			gen14794 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen14793)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*alphabet*"), gen14794)

			gen14795 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), Nil)

			gen14796 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen14795)

			gen14797 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen14796)

			gen14798 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen14797)

			gen14799 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen14798)

			gen14800 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen14799)

			gen14801 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen14800)

			gen14802 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen14801)

			gen14803 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen14802)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*special*"), gen14803)

			gen14804 := Call(__e, ShenFunc(symcons), MakeSymbol("defmacro"), Nil)

			gen14805 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.read+"), gen14804)

			gen14806 := Call(__e, ShenFunc(symcons), MakeSymbol("defcc"), gen14805)

			gen14807 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen14806)

			gen14808 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.process-datatype"), gen14807)

			gen14809 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen14808)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*extraspecial*"), gen14809)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*spy*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*datatypes*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*alldatatypes*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*shen-type-theory-enabled?*"), True)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*synonyms*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*system*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*maxcomplexity*"), MakeNumber(128))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*occurs*"), True)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*maxinferences*"), MakeNumber(1000000))
			Call(__e, ShenFunc(symset), MakeSymbol("*maximum-print-sequence-size*"), MakeNumber(20))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*catch*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*call*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*infs*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("*hush*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*optimise*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("*version*"), MakeString("Shen 22.3"))
			gen14810 := Call(__e, ShenFunc(symbound_2), MakeSymbol("*home-directory*"))

			gen14811 := Call(__e, ShenFunc(symnot), gen14810)

			if True == gen14811 {
				Call(__e, ShenFunc(symset), MakeSymbol("*home-directory*"), MakeString(""))
			} else {
				MakeSymbol("shen.skip")
			}

			gen14813 := Call(__e, ShenFunc(symbound_2), MakeSymbol("*sterror*"))

			gen14814 := Call(__e, ShenFunc(symnot), gen14813)

			if True == gen14814 {
				gen14812 := Call(__e, ShenFunc(symvalue), MakeSymbol("*stoutput*"))

				Call(__e, ShenFunc(symset), MakeSymbol("*sterror*"), gen14812)

			} else {
				MakeSymbol("shen.skip")
			}

			gen14815 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

			gen14816 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen14815)

			gen14817 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14816)

			gen14818 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen14817)

			gen14819 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14818)

			gen14820 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen14819)

			gen14821 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14820)

			gen14822 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen14821)

			gen14823 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14822)

			gen14824 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen14823)

			gen14825 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14824)

			gen14826 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen14825)

			gen14827 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14826)

			gen14828 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen14827)

			gen14829 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14828)

			gen14830 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen14829)

			gen14831 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14830)

			gen14832 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen14831)

			gen14833 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14832)

			gen14834 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen14833)

			gen14835 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14834)

			gen14836 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen14835)

			gen14837 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14836)

			gen14838 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen14837)

			gen14839 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14838)

			gen14840 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen14839)

			gen14841 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14840)

			gen14842 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen14841)

			gen14843 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14842)

			gen14844 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen14843)

			gen14845 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14844)

			gen14846 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen14845)

			gen14847 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14846)

			gen14848 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen14847)

			gen14849 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14848)

			gen14850 := Call(__e, ShenFunc(symcons), MakeSymbol("version"), gen14849)

			gen14851 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14850)

			gen14852 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen14851)

			gen14853 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14852)

			gen14854 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen14853)

			gen14855 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen14854)

			gen14856 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen14855)

			gen14857 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14856)

			gen14858 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen14857)

			gen14859 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14858)

			gen14860 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen14859)

			gen14861 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14860)

			gen14862 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen14861)

			gen14863 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14862)

			gen14864 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen14863)

			gen14865 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14864)

			gen14866 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen14865)

			gen14867 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14866)

			gen14868 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen14867)

			gen14869 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen14868)

			gen14870 := Call(__e, ShenFunc(symcons), MakeSymbol("unify!"), gen14869)

			gen14871 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen14870)

			gen14872 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen14871)

			gen14873 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14872)

			gen14874 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen14873)

			gen14875 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen14874)

			gen14876 := Call(__e, ShenFunc(symcons), MakeSymbol("unput"), gen14875)

			gen14877 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14876)

			gen14878 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen14877)

			gen14879 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen14878)

			gen14880 := Call(__e, ShenFunc(symcons), MakeSymbol("return"), gen14879)

			gen14881 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14880)

			gen14882 := Call(__e, ShenFunc(symcons), MakeSymbol("type"), gen14881)

			gen14883 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14882)

			gen14884 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen14883)

			gen14885 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14884)

			gen14886 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen14885)

			gen14887 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14886)

			gen14888 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen14887)

			gen14889 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14888)

			gen14890 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen14889)

			gen14891 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14890)

			gen14892 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen14891)

			gen14893 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14892)

			gen14894 := Call(__e, ShenFunc(symcons), MakeSymbol("tc?"), gen14893)

			gen14895 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14894)

			gen14896 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen14895)

			gen14897 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14896)

			gen14898 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen14897)

			gen14899 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14898)

			gen14900 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen14899)

			gen14901 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14900)

			gen14902 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen14901)

			gen14903 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14902)

			gen14904 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen14903)

			gen14905 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14904)

			gen14906 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen14905)

			gen14907 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen14906)

			gen14908 := Call(__e, ShenFunc(symcons), MakeSymbol("subst"), gen14907)

			gen14909 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14908)

			gen14910 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen14909)

			gen14911 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14910)

			gen14912 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen14911)

			gen14913 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14912)

			gen14914 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen14913)

			gen14915 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14914)

			gen14916 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen14915)

			gen14917 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14916)

			gen14918 := Call(__e, ShenFunc(symcons), MakeSymbol("sterror"), gen14917)

			gen14919 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14918)

			gen14920 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), gen14919)

			gen14921 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14920)

			gen14922 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), gen14921)

			gen14923 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14922)

			gen14924 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen14923)

			gen14925 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14924)

			gen14926 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen14925)

			gen14927 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14926)

			gen14928 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen14927)

			gen14929 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14928)

			gen14930 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen14929)

			gen14931 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14930)

			gen14932 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen14931)

			gen14933 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14932)

			gen14934 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen14933)

			gen14935 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14934)

			gen14936 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen14935)

			gen14937 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen14936)

			gen14938 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.require"), gen14937)

			gen14939 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14938)

			gen14940 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen14939)

			gen14941 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14940)

			gen14942 := Call(__e, ShenFunc(symcons), MakeSymbol("release"), gen14941)

			gen14943 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14942)

			gen14944 := Call(__e, ShenFunc(symcons), MakeSymbol("receive"), gen14943)

			gen14945 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14944)

			gen14946 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen14945)

			gen14947 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14946)

			gen14948 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen14947)

			gen14949 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14948)

			gen14950 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen14949)

			gen14951 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14950)

			gen14952 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen14951)

			gen14953 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14952)

			gen14954 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen14953)

			gen14955 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14954)

			gen14956 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen14955)

			gen14957 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14956)

			gen14958 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.reassemble"), gen14957)

			gen14959 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen14958)

			gen14960 := Call(__e, ShenFunc(symcons), MakeSymbol("put"), gen14959)

			gen14961 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen14960)

			gen14962 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen14961)

			gen14963 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14962)

			gen14964 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen14963)

			gen14965 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14964)

			gen14966 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen14965)

			gen14967 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14966)

			gen14968 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen14967)

			gen14969 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14968)

			gen14970 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen14969)

			gen14971 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14970)

			gen14972 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen14971)

			gen14973 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14972)

			gen14974 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen14973)

			gen14975 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14974)

			gen14976 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen14975)

			gen14977 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14976)

			gen14978 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen14977)

			gen14979 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14978)

			gen14980 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen14979)

			gen14981 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14980)

			gen14982 := Call(__e, ShenFunc(symcons), MakeSymbol("porters"), gen14981)

			gen14983 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14982)

			gen14984 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen14983)

			gen14985 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14984)

			gen14986 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen14985)

			gen14987 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen14986)

			gen14988 := Call(__e, ShenFunc(symcons), MakeSymbol("package"), gen14987)

			gen14989 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen14988)

			gen14990 := Call(__e, ShenFunc(symcons), MakeSymbol("os"), gen14989)

			gen14991 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14990)

			gen14992 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen14991)

			gen14993 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14992)

			gen14994 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen14993)

			gen14995 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14994)

			gen14996 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), gen14995)

			gen14997 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen14996)

			gen14998 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen14997)

			gen14999 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen14998)

			gen15000 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen14999)

			gen15001 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15000)

			gen15002 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen15001)

			gen15003 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15002)

			gen15004 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen15003)

			gen15005 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15004)

			gen15006 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen15005)

			gen15007 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15006)

			gen15008 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen15007)

			gen15009 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15008)

			gen15010 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen15009)

			gen15011 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15010)

			gen15012 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen15011)

			gen15013 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15012)

			gen15014 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen15013)

			gen15015 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15014)

			gen15016 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen15015)

			gen15017 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15016)

			gen15018 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen15017)

			gen15019 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15018)

			gen15020 := Call(__e, ShenFunc(symcons), MakeSymbol("macroexpand"), gen15019)

			gen15021 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15020)

			gen15022 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15021)

			gen15023 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15022)

			gen15024 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen15023)

			gen15025 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15024)

			gen15026 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen15025)

			gen15027 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15026)

			gen15028 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen15027)

			gen15029 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15028)

			gen15030 := Call(__e, ShenFunc(symcons), MakeSymbol("lineread"), gen15029)

			gen15031 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15030)

			gen15032 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen15031)

			gen15033 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15032)

			gen15034 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen15033)

			gen15035 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15034)

			gen15036 := Call(__e, ShenFunc(symcons), MakeSymbol("language"), gen15035)

			gen15037 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15036)

			gen15038 := Call(__e, ShenFunc(symcons), MakeSymbol("kill"), gen15037)

			gen15039 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15038)

			gen15040 := Call(__e, ShenFunc(symcons), MakeSymbol("it"), gen15039)

			gen15041 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15040)

			gen15042 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen15041)

			gen15043 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15042)

			gen15044 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen15043)

			gen15045 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15044)

			gen15046 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen15045)

			gen15047 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15046)

			gen15048 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen15047)

			gen15049 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15048)

			gen15050 := Call(__e, ShenFunc(symcons), MakeSymbol("input"), gen15049)

			gen15051 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15050)

			gen15052 := Call(__e, ShenFunc(symcons), MakeSymbol("inferences"), gen15051)

			gen15053 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen15052)

			gen15054 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen15053)

			gen15055 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15054)

			gen15056 := Call(__e, ShenFunc(symcons), MakeSymbol("intern"), gen15055)

			gen15057 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15056)

			gen15058 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen15057)

			gen15059 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen15058)

			gen15060 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen15059)

			gen15061 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15060)

			gen15062 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen15061)

			gen15063 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15062)

			gen15064 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen15063)

			gen15065 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15064)

			gen15066 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen15065)

			gen15067 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15066)

			gen15068 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen15067)

			gen15069 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15068)

			gen15070 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen15069)

			gen15071 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15070)

			gen15072 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen15071)

			gen15073 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15072)

			gen15074 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen15073)

			gen15075 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15074)

			gen15076 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen15075)

			gen15077 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15076)

			gen15078 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen15077)

			gen15079 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15078)

			gen15080 := Call(__e, ShenFunc(symcons), MakeSymbol("<-address"), gen15079)

			gen15081 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen15080)

			gen15082 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen15081)

			gen15083 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15082)

			gen15084 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen15083)

			gen15085 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen15084)

			gen15086 := Call(__e, ShenFunc(symcons), MakeSymbol("get"), gen15085)

			gen15087 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15086)

			gen15088 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen15087)

			gen15089 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15088)

			gen15090 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen15089)

			gen15091 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15090)

			gen15092 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen15091)

			gen15093 := Call(__e, ShenFunc(symcons), MakeNumber(5), gen15092)

			gen15094 := Call(__e, ShenFunc(symcons), MakeSymbol("findall"), gen15093)

			gen15095 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15094)

			gen15096 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen15095)

			gen15097 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15096)

			gen15098 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), gen15097)

			gen15099 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15098)

			gen15100 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen15099)

			gen15101 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15100)

			gen15102 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen15101)

			gen15103 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15102)

			gen15104 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen15103)

			gen15105 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15104)

			gen15106 := Call(__e, ShenFunc(symcons), MakeSymbol("eval-kl"), gen15105)

			gen15107 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15106)

			gen15108 := Call(__e, ShenFunc(symcons), MakeSymbol("eval"), gen15107)

			gen15109 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15108)

			gen15110 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.interror"), gen15109)

			gen15111 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15110)

			gen15112 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen15111)

			gen15113 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15112)

			gen15114 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen15113)

			gen15115 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15114)

			gen15116 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen15115)

			gen15117 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15116)

			gen15118 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen15117)

			gen15119 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15118)

			gen15120 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen15119)

			gen15121 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15120)

			gen15122 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen15121)

			gen15123 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15122)

			gen15124 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen15123)

			gen15125 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15124)

			gen15126 := Call(__e, ShenFunc(symcons), MakeSymbol("declare"), gen15125)

			gen15127 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15126)

			gen15128 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen15127)

			gen15129 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15128)

			gen15130 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen15129)

			gen15131 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15130)

			gen15132 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen15131)

			gen15133 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15132)

			gen15134 := Call(__e, ShenFunc(symcons), MakeSymbol("concat"), gen15133)

			gen15135 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen15134)

			gen15136 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen15135)

			gen15137 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15136)

			gen15138 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen15137)

			gen15139 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15138)

			gen15140 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen15139)

			gen15141 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15140)

			gen15142 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen15141)

			gen15143 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15142)

			gen15144 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen15143)

			gen15145 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15144)

			gen15146 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen15145)

			gen15147 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15146)

			gen15148 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen15147)

			gen15149 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15148)

			gen15150 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen15149)

			gen15151 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15150)

			gen15152 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen15151)

			gen15153 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen15152)

			gen15154 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen15153)

			gen15155 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15154)

			gen15156 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector"), gen15155)

			gen15157 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen15156)

			gen15158 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen15157)

			gen15159 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen15158)

			gen15160 := Call(__e, ShenFunc(symcons), MakeSymbol("abort"), gen15159)

			Call(__e, ShenFunc(symshen_4initialise__arity__table), gen15160)

			gen15161 := Call(__e, ShenFunc(symintern), MakeString("shen"))

			gen15162 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector"), Nil)

			gen15163 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen15162)

			gen15164 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen15163)

			gen15165 := Call(__e, ShenFunc(symcons), MakeSymbol("<-address"), gen15164)

			gen15166 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen15165)

			gen15167 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen15166)

			gen15168 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen15167)

			gen15169 := Call(__e, ShenFunc(symcons), MakeSymbol("abort"), gen15168)

			gen15170 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen15169)

			gen15171 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen15170)

			gen15172 := Call(__e, ShenFunc(symcons), MakeSymbol("bar!"), gen15171)

			gen15173 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen15172)

			gen15174 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen15173)

			gen15175 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen15174)

			gen15176 := Call(__e, ShenFunc(symcons), MakeSymbol("bind"), gen15175)

			gen15177 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen15176)

			gen15178 := Call(__e, ShenFunc(symcons), MakeSymbol("call"), gen15177)

			gen15179 := Call(__e, ShenFunc(symcons), MakeSymbol("cases"), gen15178)

			gen15180 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen15179)

			gen15181 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen15180)

			gen15182 := Call(__e, ShenFunc(symcons), MakeSymbol("concat"), gen15181)

			gen15183 := Call(__e, ShenFunc(symcons), MakeSymbol("cond"), gen15182)

			gen15184 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen15183)

			gen15185 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen15184)

			gen15186 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen15185)

			gen15187 := Call(__e, ShenFunc(symcons), MakeSymbol("cut"), gen15186)

			gen15188 := Call(__e, ShenFunc(symcons), MakeSymbol("datatype"), gen15187)

			gen15189 := Call(__e, ShenFunc(symcons), MakeSymbol("declare"), gen15188)

			gen15190 := Call(__e, ShenFunc(symcons), MakeSymbol("defprolog"), gen15189)

			gen15191 := Call(__e, ShenFunc(symcons), MakeSymbol("defcc"), gen15190)

			gen15192 := Call(__e, ShenFunc(symcons), MakeSymbol("defmacro"), gen15191)

			gen15193 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen15192)

			gen15194 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen15193)

			gen15195 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen15194)

			gen15196 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen15195)

			gen15197 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen15196)

			gen15198 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen15197)

			gen15199 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen15198)

			gen15200 := Call(__e, ShenFunc(symcons), MakeSymbol("error"), gen15199)

			gen15201 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen15200)

			gen15202 := Call(__e, ShenFunc(symcons), MakeSymbol("eval"), gen15201)

			gen15203 := Call(__e, ShenFunc(symcons), MakeSymbol("eval-kl"), gen15202)

			gen15204 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen15203)

			gen15205 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen15204)

			gen15206 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen15205)

			gen15207 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen15206)

			gen15208 := Call(__e, ShenFunc(symcons), False, gen15207)

			gen15209 := Call(__e, ShenFunc(symcons), MakeSymbol("findall"), gen15208)

			gen15210 := Call(__e, ShenFunc(symcons), MakeSymbol("fwhen"), gen15209)

			gen15211 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen15210)

			gen15212 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), gen15211)

			gen15213 := Call(__e, ShenFunc(symcons), MakeSymbol("file"), gen15212)

			gen15214 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen15213)

			gen15215 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen15214)

			gen15216 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen15215)

			gen15217 := Call(__e, ShenFunc(symcons), MakeSymbol("function"), gen15216)

			gen15218 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen15217)

			gen15219 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen15218)

			gen15220 := Call(__e, ShenFunc(symcons), MakeSymbol("get"), gen15219)

			gen15221 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen15220)

			gen15222 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen15221)

			gen15223 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen15222)

			gen15224 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen15223)

			gen15225 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen15224)

			gen15226 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen15225)

			gen15227 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen15226)

			gen15228 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen15227)

			gen15229 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen15228)

			gen15230 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen15229)

			gen15231 := Call(__e, ShenFunc(symcons), MakeSymbol("it"), gen15230)

			gen15232 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen15231)

			gen15233 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen15232)

			gen15234 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen15233)

			gen15235 := Call(__e, ShenFunc(symcons), MakeSymbol("input"), gen15234)

			gen15236 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen15235)

			gen15237 := Call(__e, ShenFunc(symcons), MakeSymbol("intern"), gen15236)

			gen15238 := Call(__e, ShenFunc(symcons), MakeSymbol("inferences"), gen15237)

			gen15239 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen15238)

			gen15240 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen15239)

			gen15241 := Call(__e, ShenFunc(symcons), MakeSymbol("kill"), gen15240)

			gen15242 := Call(__e, ShenFunc(symcons), MakeSymbol("language"), gen15241)

			gen15243 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen15242)

			gen15244 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen15243)

			gen15245 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen15244)

			gen15246 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen15245)

			gen15247 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen15246)

			gen15248 := Call(__e, ShenFunc(symcons), MakeSymbol("lineread"), gen15247)

			gen15249 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15248)

			gen15250 := Call(__e, ShenFunc(symcons), MakeSymbol("loaded"), gen15249)

			gen15251 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen15250)

			gen15252 := Call(__e, ShenFunc(symcons), MakeSymbol("make-string"), gen15251)

			gen15253 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen15252)

			gen15254 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen15253)

			gen15255 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen15254)

			gen15256 := Call(__e, ShenFunc(symcons), MakeSymbol("macroexpand"), gen15255)

			gen15257 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen15256)

			gen15258 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen15257)

			gen15259 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen15258)

			gen15260 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen15259)

			gen15261 := Call(__e, ShenFunc(symcons), MakeSymbol("null"), gen15260)

			gen15262 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15261)

			gen15263 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen15262)

			gen15264 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen15263)

			gen15265 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen15264)

			gen15266 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen15265)

			gen15267 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), gen15266)

			gen15268 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen15267)

			gen15269 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen15268)

			gen15270 := Call(__e, ShenFunc(symcons), MakeSymbol("os"), gen15269)

			gen15271 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), gen15270)

			gen15272 := Call(__e, ShenFunc(symcons), MakeSymbol("output"), gen15271)

			gen15273 := Call(__e, ShenFunc(symcons), MakeSymbol("package"), gen15272)

			gen15274 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen15273)

			gen15275 := Call(__e, ShenFunc(symcons), MakeSymbol("porters"), gen15274)

			gen15276 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen15275)

			gen15277 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen15276)

			gen15278 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen15277)

			gen15279 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen15278)

			gen15280 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen15279)

			gen15281 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen15280)

			gen15282 := Call(__e, ShenFunc(symcons), MakeSymbol("prolog?"), gen15281)

			gen15283 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen15282)

			gen15284 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen15283)

			gen15285 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen15284)

			gen15286 := Call(__e, ShenFunc(symcons), MakeSymbol("put"), gen15285)

			gen15287 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen15286)

			gen15288 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen15287)

			gen15289 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen15288)

			gen15290 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen15289)

			gen15291 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen15290)

			gen15292 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen15291)

			gen15293 := Call(__e, ShenFunc(symcons), MakeSymbol("receive"), gen15292)

			gen15294 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen15293)

			gen15295 := Call(__e, ShenFunc(symcons), MakeSymbol("release"), gen15294)

			gen15296 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen15295)

			gen15297 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen15296)

			gen15298 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), gen15297)

			gen15299 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen15298)

			gen15300 := Call(__e, ShenFunc(symcons), MakeSymbol("save"), gen15299)

			gen15301 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen15300)

			gen15302 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen15301)

			gen15303 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen15302)

			gen15304 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen15303)

			gen15305 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen15304)

			gen15306 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen15305)

			gen15307 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), gen15306)

			gen15308 := Call(__e, ShenFunc(symcons), MakeSymbol("sterror"), gen15307)

			gen15309 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), gen15308)

			gen15310 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15309)

			gen15311 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen15310)

			gen15312 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen15311)

			gen15313 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen15312)

			gen15314 := Call(__e, ShenFunc(symcons), MakeSymbol("subst"), gen15313)

			gen15315 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen15314)

			gen15316 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen15315)

			gen15317 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen15316)

			gen15318 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15317)

			gen15319 := Call(__e, ShenFunc(symcons), MakeSymbol("synonyms"), gen15318)

			gen15320 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen15319)

			gen15321 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen15320)

			gen15322 := Call(__e, ShenFunc(symcons), MakeSymbol("tlv"), gen15321)

			gen15323 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen15322)

			gen15324 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen15323)

			gen15325 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen15324)

			gen15326 := Call(__e, ShenFunc(symcons), MakeSymbol("tc?"), gen15325)

			gen15327 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen15326)

			gen15328 := Call(__e, ShenFunc(symcons), MakeSymbol("time"), gen15327)

			gen15329 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen15328)

			gen15330 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen15329)

			gen15331 := Call(__e, ShenFunc(symcons), True, gen15330)

			gen15332 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen15331)

			gen15333 := Call(__e, ShenFunc(symcons), MakeSymbol("type"), gen15332)

			gen15334 := Call(__e, ShenFunc(symcons), MakeSymbol("return"), gen15333)

			gen15335 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen15334)

			gen15336 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen15335)

			gen15337 := Call(__e, ShenFunc(symcons), MakeSymbol("unput"), gen15336)

			gen15338 := Call(__e, ShenFunc(symcons), MakeSymbol("unify!"), gen15337)

			gen15339 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen15338)

			gen15340 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen15339)

			gen15341 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.unix"), gen15340)

			gen15342 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), gen15341)

			gen15343 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen15342)

			gen15344 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen15343)

			gen15345 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen15344)

			gen15346 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15345)

			gen15347 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen15346)

			gen15348 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen15347)

			gen15349 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen15348)

			gen15350 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen15349)

			gen15351 := Call(__e, ShenFunc(symcons), MakeSymbol("verified"), gen15350)

			gen15352 := Call(__e, ShenFunc(symcons), MakeSymbol("version"), gen15351)

			gen15353 := Call(__e, ShenFunc(symcons), MakeSymbol("warn"), gen15352)

			gen15354 := Call(__e, ShenFunc(symcons), MakeSymbol("when"), gen15353)

			gen15355 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen15354)

			gen15356 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen15355)

			gen15357 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen15356)

			gen15358 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen15357)

			gen15359 := Call(__e, ShenFunc(symcons), MakeSymbol(">>"), gen15358)

			gen15360 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen15359)

			gen15361 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen15360)

			gen15362 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen15361)

			gen15363 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen15362)

			gen15364 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen15363)

			gen15365 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen15364)

			gen15366 := Call(__e, ShenFunc(symcons), MakeSymbol("$"), gen15365)

			gen15367 := Call(__e, ShenFunc(symcons), MakeSymbol("=!"), gen15366)

			gen15368 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen15367)

			gen15369 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen15368)

			gen15370 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen15369)

			gen15371 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen15370)

			gen15372 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen15371)

			gen15373 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen15372)

			gen15374 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen15373)

			gen15375 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen15374)

			gen15376 := Call(__e, ShenFunc(symcons), MakeSymbol("<-"), gen15375)

			gen15377 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen15376)

			gen15378 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen15377)

			gen15379 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen15378)

			gen15380 := Call(__e, ShenFunc(symcons), MakeSymbol("*hush*"), gen15379)

			gen15381 := Call(__e, ShenFunc(symcons), MakeSymbol("*porters*"), gen15380)

			gen15382 := Call(__e, ShenFunc(symcons), MakeSymbol("*port*"), gen15381)

			gen15383 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), gen15382)

			gen15384 := Call(__e, ShenFunc(symcons), MakeSymbol("*release*"), gen15383)

			gen15385 := Call(__e, ShenFunc(symcons), MakeSymbol("*os*"), gen15384)

			gen15386 := Call(__e, ShenFunc(symcons), MakeSymbol("*macros*"), gen15385)

			gen15387 := Call(__e, ShenFunc(symcons), MakeSymbol("*maximum-print-sequence-size*"), gen15386)

			gen15388 := Call(__e, ShenFunc(symcons), MakeSymbol("*version*"), gen15387)

			gen15389 := Call(__e, ShenFunc(symcons), MakeSymbol("*home-directory*"), gen15388)

			gen15390 := Call(__e, ShenFunc(symcons), MakeSymbol("*sterror*"), gen15389)

			gen15391 := Call(__e, ShenFunc(symcons), MakeSymbol("*stoutput*"), gen15390)

			gen15392 := Call(__e, ShenFunc(symcons), MakeSymbol("*stinput*"), gen15391)

			gen15393 := Call(__e, ShenFunc(symcons), MakeSymbol("*implementation*"), gen15392)

			gen15394 := Call(__e, ShenFunc(symcons), MakeSymbol("*language*"), gen15393)

			gen15395 := Call(__e, ShenFunc(symcons), MakeSymbol(","), gen15394)

			gen15396 := Call(__e, ShenFunc(symcons), MakeSymbol("_"), gen15395)

			gen15397 := Call(__e, ShenFunc(symcons), MakeSymbol(":="), gen15396)

			gen15398 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen15397)

			gen15399 := Call(__e, ShenFunc(symcons), MakeSymbol(";"), gen15398)

			gen15400 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen15399)

			gen15401 := Call(__e, ShenFunc(symcons), MakeSymbol("&&"), gen15400)

			gen15402 := Call(__e, ShenFunc(symcons), MakeSymbol("<--"), gen15401)

			gen15403 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15402)

			gen15404 := Call(__e, ShenFunc(symcons), MakeSymbol("{"), gen15403)

			gen15405 := Call(__e, ShenFunc(symcons), MakeSymbol("}"), gen15404)

			gen15406 := Call(__e, ShenFunc(symcons), MakeSymbol("!"), gen15405)

			gen15407 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			Call(__e, ShenFunc(symput), gen15161, MakeSymbol("shen.external-symbols"), gen15406, gen15407)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*history*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*step*"), False)
			gen15408 := Call(__e, ShenFunc(symabsvector), MakeNumber(0))

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*empty-absvector*"), gen15408)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-environment"), gen15409)

		gen16547 := MakeNative(func(__e Evaluator) {
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), Nil)
			gen15410 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15411 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15410)

			gen15412 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15411)

			gen15413 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen15412)

			gen15414 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15415 := Call(__e, ShenFunc(symcons), gen15413, gen15414)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15415)

			gen15416 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15417 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15416)

			gen15418 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15419 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15418)

			gen15420 := Call(__e, ShenFunc(symcons), gen15419, Nil)

			gen15421 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15420)

			gen15422 := Call(__e, ShenFunc(symcons), gen15417, gen15421)

			gen15423 := Call(__e, ShenFunc(symcons), gen15422, Nil)

			gen15424 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15423)

			gen15425 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15424)

			gen15426 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen15425)

			gen15427 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15428 := Call(__e, ShenFunc(symcons), gen15426, gen15427)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15428)

			gen15429 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15430 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15429)

			gen15431 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen15430)

			gen15432 := Call(__e, ShenFunc(symcons), gen15431, Nil)

			gen15433 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15432)

			gen15434 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen15433)

			gen15435 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen15434)

			gen15436 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15437 := Call(__e, ShenFunc(symcons), gen15435, gen15436)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15437)

			gen15438 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15439 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15438)

			gen15440 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15439)

			gen15441 := Call(__e, ShenFunc(symcons), gen15440, Nil)

			gen15442 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15441)

			gen15443 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15442)

			gen15444 := Call(__e, ShenFunc(symcons), gen15443, Nil)

			gen15445 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15444)

			gen15446 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15445)

			gen15447 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.app"), gen15446)

			gen15448 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15449 := Call(__e, ShenFunc(symcons), gen15447, gen15448)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15449)

			gen15450 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15451 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15450)

			gen15452 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15453 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15452)

			gen15454 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15455 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15454)

			gen15456 := Call(__e, ShenFunc(symcons), gen15455, Nil)

			gen15457 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15456)

			gen15458 := Call(__e, ShenFunc(symcons), gen15453, gen15457)

			gen15459 := Call(__e, ShenFunc(symcons), gen15458, Nil)

			gen15460 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15459)

			gen15461 := Call(__e, ShenFunc(symcons), gen15451, gen15460)

			gen15462 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen15461)

			gen15463 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15464 := Call(__e, ShenFunc(symcons), gen15462, gen15463)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15464)

			gen15465 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15466 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15465)

			gen15467 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15466)

			gen15468 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen15467)

			gen15469 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15470 := Call(__e, ShenFunc(symcons), gen15468, gen15469)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15470)

			gen15471 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15472 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15471)

			gen15473 := Call(__e, ShenFunc(symcons), gen15472, Nil)

			gen15474 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15473)

			gen15475 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15476 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15475)

			gen15477 := Call(__e, ShenFunc(symcons), gen15476, Nil)

			gen15478 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15477)

			gen15479 := Call(__e, ShenFunc(symcons), gen15474, gen15478)

			gen15480 := Call(__e, ShenFunc(symcons), gen15479, Nil)

			gen15481 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15480)

			gen15482 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15481)

			gen15483 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen15482)

			gen15484 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15485 := Call(__e, ShenFunc(symcons), gen15483, gen15484)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15485)

			gen15486 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15487 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15486)

			gen15488 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15487)

			gen15489 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen15488)

			gen15490 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15491 := Call(__e, ShenFunc(symcons), gen15489, gen15490)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15491)

			gen15492 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15493 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15492)

			gen15494 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15493)

			gen15495 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen15494)

			gen15496 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15497 := Call(__e, ShenFunc(symcons), gen15495, gen15496)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15497)

			gen15498 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15499 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15498)

			gen15500 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15499)

			gen15501 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen15500)

			gen15502 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15503 := Call(__e, ShenFunc(symcons), gen15501, gen15502)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15503)

			gen15504 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15505 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen15504)

			gen15506 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15507 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15506)

			gen15508 := Call(__e, ShenFunc(symcons), gen15507, Nil)

			gen15509 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15508)

			gen15510 := Call(__e, ShenFunc(symcons), gen15505, gen15509)

			gen15511 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen15510)

			gen15512 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15513 := Call(__e, ShenFunc(symcons), gen15511, gen15512)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15513)

			gen15514 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15515 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15514)

			gen15516 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15515)

			gen15517 := Call(__e, ShenFunc(symcons), gen15516, Nil)

			gen15518 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15517)

			gen15519 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15518)

			gen15520 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen15519)

			gen15521 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15522 := Call(__e, ShenFunc(symcons), gen15520, gen15521)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15522)

			gen15523 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15524 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen15523)

			gen15525 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15524)

			gen15526 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15527 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15526)

			gen15528 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15527)

			gen15529 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15530 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15529)

			gen15531 := Call(__e, ShenFunc(symcons), gen15528, gen15530)

			gen15532 := Call(__e, ShenFunc(symcons), gen15531, Nil)

			gen15533 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15532)

			gen15534 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15533)

			gen15535 := Call(__e, ShenFunc(symcons), gen15534, Nil)

			gen15536 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15535)

			gen15537 := Call(__e, ShenFunc(symcons), gen15525, gen15536)

			gen15538 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen15537)

			gen15539 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15540 := Call(__e, ShenFunc(symcons), gen15538, gen15539)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15540)

			gen15541 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15542 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15541)

			gen15543 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15542)

			gen15544 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen15543)

			gen15545 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15546 := Call(__e, ShenFunc(symcons), gen15544, gen15545)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15546)

			gen15547 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15548 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15547)

			gen15549 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15548)

			gen15550 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15551 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15550)

			gen15552 := Call(__e, ShenFunc(symcons), gen15549, gen15551)

			gen15553 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen15552)

			gen15554 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15555 := Call(__e, ShenFunc(symcons), gen15553, gen15554)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15555)

			gen15556 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15557 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15556)

			gen15558 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15559 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15558)

			gen15560 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15561 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15560)

			gen15562 := Call(__e, ShenFunc(symcons), gen15561, Nil)

			gen15563 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15562)

			gen15564 := Call(__e, ShenFunc(symcons), gen15559, gen15563)

			gen15565 := Call(__e, ShenFunc(symcons), gen15564, Nil)

			gen15566 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15565)

			gen15567 := Call(__e, ShenFunc(symcons), gen15557, gen15566)

			gen15568 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen15567)

			gen15569 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15570 := Call(__e, ShenFunc(symcons), gen15568, gen15569)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15570)

			gen15571 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15572 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15571)

			gen15573 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen15572)

			gen15574 := Call(__e, ShenFunc(symcons), gen15573, Nil)

			gen15575 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15574)

			gen15576 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15575)

			gen15577 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen15576)

			gen15578 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15579 := Call(__e, ShenFunc(symcons), gen15577, gen15578)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15579)

			gen15580 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15581 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15580)

			gen15582 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15583 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15582)

			gen15584 := Call(__e, ShenFunc(symcons), gen15583, Nil)

			gen15585 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen15584)

			gen15586 := Call(__e, ShenFunc(symcons), gen15581, gen15585)

			gen15587 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen15586)

			gen15588 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15589 := Call(__e, ShenFunc(symcons), gen15587, gen15588)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15589)

			gen15590 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15591 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15590)

			gen15592 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15593 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15592)

			gen15594 := Call(__e, ShenFunc(symcons), gen15593, Nil)

			gen15595 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen15594)

			gen15596 := Call(__e, ShenFunc(symcons), gen15591, gen15595)

			gen15597 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen15596)

			gen15598 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15599 := Call(__e, ShenFunc(symcons), gen15597, gen15598)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15599)

			gen15600 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15601 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15600)

			gen15602 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15603 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15602)

			gen15604 := Call(__e, ShenFunc(symcons), gen15601, gen15603)

			gen15605 := Call(__e, ShenFunc(symcons), gen15604, Nil)

			gen15606 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15605)

			gen15607 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15606)

			gen15608 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen15607)

			gen15609 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15610 := Call(__e, ShenFunc(symcons), gen15608, gen15609)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15610)

			gen15611 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15612 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15611)

			gen15613 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15612)

			gen15614 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen15613)

			gen15615 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15616 := Call(__e, ShenFunc(symcons), gen15614, gen15615)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15616)

			gen15617 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15618 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15617)

			gen15619 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15618)

			gen15620 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen15619)

			gen15621 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15622 := Call(__e, ShenFunc(symcons), gen15620, gen15621)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15622)

			gen15623 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15624 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15623)

			gen15625 := Call(__e, ShenFunc(symcons), gen15624, Nil)

			gen15626 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15625)

			gen15627 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15626)

			gen15628 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen15627)

			gen15629 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15630 := Call(__e, ShenFunc(symcons), gen15628, gen15629)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15630)

			gen15631 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15632 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15631)

			gen15633 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen15632)

			gen15634 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen15633)

			gen15635 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15636 := Call(__e, ShenFunc(symcons), gen15634, gen15635)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15636)

			gen15637 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15638 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15637)

			gen15639 := Call(__e, ShenFunc(symcons), gen15638, Nil)

			gen15640 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15639)

			gen15641 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15640)

			gen15642 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen15641)

			gen15643 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15644 := Call(__e, ShenFunc(symcons), gen15642, gen15643)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15644)

			gen15645 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15646 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15645)

			gen15647 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), gen15646)

			gen15648 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15649 := Call(__e, ShenFunc(symcons), gen15647, gen15648)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15649)

			gen15650 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15651 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15650)

			gen15652 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15651)

			gen15653 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15654 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15653)

			gen15655 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15654)

			gen15656 := Call(__e, ShenFunc(symcons), gen15655, Nil)

			gen15657 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15656)

			gen15658 := Call(__e, ShenFunc(symcons), gen15652, gen15657)

			gen15659 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen15658)

			gen15660 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15661 := Call(__e, ShenFunc(symcons), gen15659, gen15660)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15661)

			gen15662 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15663 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15662)

			gen15664 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15663)

			gen15665 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15666 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15665)

			gen15667 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15666)

			gen15668 := Call(__e, ShenFunc(symcons), gen15667, Nil)

			gen15669 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15668)

			gen15670 := Call(__e, ShenFunc(symcons), gen15664, gen15669)

			gen15671 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen15670)

			gen15672 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15673 := Call(__e, ShenFunc(symcons), gen15671, gen15672)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15673)

			gen15674 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15675 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen15674)

			gen15676 := Call(__e, ShenFunc(symcons), gen15675, Nil)

			gen15677 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15676)

			gen15678 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15677)

			gen15679 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen15678)

			gen15680 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15681 := Call(__e, ShenFunc(symcons), gen15679, gen15680)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15681)

			gen15682 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15683 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen15682)

			gen15684 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15683)

			gen15685 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15686 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15685)

			gen15687 := Call(__e, ShenFunc(symcons), gen15684, gen15686)

			gen15688 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen15687)

			gen15689 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15690 := Call(__e, ShenFunc(symcons), gen15688, gen15689)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15690)

			gen15691 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15692 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15691)

			gen15693 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15692)

			gen15694 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15695 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15694)

			gen15696 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15695)

			gen15697 := Call(__e, ShenFunc(symcons), gen15696, Nil)

			gen15698 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15697)

			gen15699 := Call(__e, ShenFunc(symcons), gen15693, gen15698)

			gen15700 := Call(__e, ShenFunc(symcons), MakeSymbol("function"), gen15699)

			gen15701 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15702 := Call(__e, ShenFunc(symcons), gen15700, gen15701)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15702)

			gen15703 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15704 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15703)

			gen15705 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15704)

			gen15706 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen15705)

			gen15707 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15708 := Call(__e, ShenFunc(symcons), gen15706, gen15707)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15708)

			gen15709 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15710 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15709)

			gen15711 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15712 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15711)

			gen15713 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15712)

			gen15714 := Call(__e, ShenFunc(symcons), gen15713, Nil)

			gen15715 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15714)

			gen15716 := Call(__e, ShenFunc(symcons), gen15710, gen15715)

			gen15717 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen15716)

			gen15718 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15719 := Call(__e, ShenFunc(symcons), gen15717, gen15718)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15719)

			gen15720 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15721 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15720)

			gen15722 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15723 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15722)

			gen15724 := Call(__e, ShenFunc(symcons), gen15723, Nil)

			gen15725 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15724)

			gen15726 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15725)

			gen15727 := Call(__e, ShenFunc(symcons), gen15726, Nil)

			gen15728 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15727)

			gen15729 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15728)

			gen15730 := Call(__e, ShenFunc(symcons), gen15729, Nil)

			gen15731 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15730)

			gen15732 := Call(__e, ShenFunc(symcons), gen15721, gen15731)

			gen15733 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen15732)

			gen15734 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15735 := Call(__e, ShenFunc(symcons), gen15733, gen15734)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15735)

			gen15736 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15737 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15736)

			gen15738 := Call(__e, ShenFunc(symcons), gen15737, Nil)

			gen15739 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15738)

			gen15740 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15739)

			gen15741 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15740)

			gen15742 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15743 := Call(__e, ShenFunc(symcons), gen15741, gen15742)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15743)

			gen15744 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15745 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15744)

			gen15746 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15745)

			gen15747 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen15746)

			gen15748 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15749 := Call(__e, ShenFunc(symcons), gen15747, gen15748)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15749)

			gen15750 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15751 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15750)

			gen15752 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15751)

			gen15753 := Call(__e, ShenFunc(symcons), gen15752, Nil)

			gen15754 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15753)

			gen15755 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15754)

			gen15756 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen15755)

			gen15757 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15758 := Call(__e, ShenFunc(symcons), gen15756, gen15757)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15758)

			gen15759 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15760 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15759)

			gen15761 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15762 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15761)

			gen15763 := Call(__e, ShenFunc(symcons), gen15760, gen15762)

			gen15764 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen15763)

			gen15765 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15766 := Call(__e, ShenFunc(symcons), gen15764, gen15765)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15766)

			gen15767 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15768 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15767)

			gen15769 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15770 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15769)

			gen15771 := Call(__e, ShenFunc(symcons), gen15768, gen15770)

			gen15772 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen15771)

			gen15773 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15774 := Call(__e, ShenFunc(symcons), gen15772, gen15773)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15774)

			gen15775 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15776 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15775)

			gen15777 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15776)

			gen15778 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen15777)

			gen15779 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15780 := Call(__e, ShenFunc(symcons), gen15778, gen15779)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15780)

			gen15781 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15782 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15781)

			gen15783 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15782)

			gen15784 := Call(__e, ShenFunc(symcons), gen15783, Nil)

			gen15785 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15784)

			gen15786 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15785)

			gen15787 := Call(__e, ShenFunc(symcons), gen15786, Nil)

			gen15788 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15787)

			gen15789 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen15788)

			gen15790 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen15789)

			gen15791 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15792 := Call(__e, ShenFunc(symcons), gen15790, gen15791)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15792)

			gen15793 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15794 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15793)

			gen15795 := Call(__e, ShenFunc(symcons), MakeSymbol("it"), gen15794)

			gen15796 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15797 := Call(__e, ShenFunc(symcons), gen15795, gen15796)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15797)

			gen15798 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15799 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15798)

			gen15800 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen15799)

			gen15801 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15802 := Call(__e, ShenFunc(symcons), gen15800, gen15801)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15802)

			gen15803 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15804 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15803)

			gen15805 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15806 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15805)

			gen15807 := Call(__e, ShenFunc(symcons), gen15806, Nil)

			gen15808 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15807)

			gen15809 := Call(__e, ShenFunc(symcons), gen15804, gen15808)

			gen15810 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen15809)

			gen15811 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15812 := Call(__e, ShenFunc(symcons), gen15810, gen15811)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15812)

			gen15813 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15814 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15813)

			gen15815 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15816 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15815)

			gen15817 := Call(__e, ShenFunc(symcons), gen15816, Nil)

			gen15818 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15817)

			gen15819 := Call(__e, ShenFunc(symcons), gen15814, gen15818)

			gen15820 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen15819)

			gen15821 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15822 := Call(__e, ShenFunc(symcons), gen15820, gen15821)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15822)

			gen15823 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15824 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15823)

			gen15825 := Call(__e, ShenFunc(symcons), MakeSymbol("inferences"), gen15824)

			gen15826 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15827 := Call(__e, ShenFunc(symcons), gen15825, gen15826)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15827)

			gen15828 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15829 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15828)

			gen15830 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15829)

			gen15831 := Call(__e, ShenFunc(symcons), gen15830, Nil)

			gen15832 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15831)

			gen15833 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15832)

			gen15834 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.insert"), gen15833)

			gen15835 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15836 := Call(__e, ShenFunc(symcons), gen15834, gen15835)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15836)

			gen15837 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15838 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15837)

			gen15839 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15838)

			gen15840 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen15839)

			gen15841 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15842 := Call(__e, ShenFunc(symcons), gen15840, gen15841)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15842)

			gen15843 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15844 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15843)

			gen15845 := Call(__e, ShenFunc(symcons), gen15844, Nil)

			gen15846 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15845)

			gen15847 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15846)

			gen15848 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen15847)

			gen15849 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15850 := Call(__e, ShenFunc(symcons), gen15848, gen15849)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15850)

			gen15851 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15852 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15851)

			gen15853 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15854 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15853)

			gen15855 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15856 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15855)

			gen15857 := Call(__e, ShenFunc(symcons), gen15856, Nil)

			gen15858 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15857)

			gen15859 := Call(__e, ShenFunc(symcons), gen15854, gen15858)

			gen15860 := Call(__e, ShenFunc(symcons), gen15859, Nil)

			gen15861 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15860)

			gen15862 := Call(__e, ShenFunc(symcons), gen15852, gen15861)

			gen15863 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen15862)

			gen15864 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15865 := Call(__e, ShenFunc(symcons), gen15863, gen15864)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15865)

			gen15866 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15867 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15866)

			gen15868 := Call(__e, ShenFunc(symcons), MakeSymbol("kill"), gen15867)

			gen15869 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15870 := Call(__e, ShenFunc(symcons), gen15868, gen15869)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15870)

			gen15871 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15872 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15871)

			gen15873 := Call(__e, ShenFunc(symcons), MakeSymbol("language"), gen15872)

			gen15874 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15875 := Call(__e, ShenFunc(symcons), gen15873, gen15874)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15875)

			gen15876 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15877 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15876)

			gen15878 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15879 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15878)

			gen15880 := Call(__e, ShenFunc(symcons), gen15877, gen15879)

			gen15881 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen15880)

			gen15882 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15883 := Call(__e, ShenFunc(symcons), gen15881, gen15882)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15883)

			gen15884 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15885 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen15884)

			gen15886 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15887 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15886)

			gen15888 := Call(__e, ShenFunc(symcons), gen15885, gen15887)

			gen15889 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen15888)

			gen15890 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15891 := Call(__e, ShenFunc(symcons), gen15889, gen15890)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15891)

			gen15892 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen15893 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15892)

			gen15894 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen15893)

			gen15895 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen15894)

			gen15896 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15897 := Call(__e, ShenFunc(symcons), gen15895, gen15896)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15897)

			gen15898 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15899 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15898)

			gen15900 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15899)

			gen15901 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15902 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15901)

			gen15903 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15904 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15903)

			gen15905 := Call(__e, ShenFunc(symcons), gen15904, Nil)

			gen15906 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15905)

			gen15907 := Call(__e, ShenFunc(symcons), gen15902, gen15906)

			gen15908 := Call(__e, ShenFunc(symcons), gen15907, Nil)

			gen15909 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15908)

			gen15910 := Call(__e, ShenFunc(symcons), gen15900, gen15909)

			gen15911 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen15910)

			gen15912 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15913 := Call(__e, ShenFunc(symcons), gen15911, gen15912)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15913)

			gen15914 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15915 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15914)

			gen15916 := Call(__e, ShenFunc(symcons), gen15915, Nil)

			gen15917 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15916)

			gen15918 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15917)

			gen15919 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15920 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15919)

			gen15921 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen15922 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15921)

			gen15923 := Call(__e, ShenFunc(symcons), gen15922, Nil)

			gen15924 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15923)

			gen15925 := Call(__e, ShenFunc(symcons), gen15920, gen15924)

			gen15926 := Call(__e, ShenFunc(symcons), gen15925, Nil)

			gen15927 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15926)

			gen15928 := Call(__e, ShenFunc(symcons), gen15918, gen15927)

			gen15929 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen15928)

			gen15930 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15931 := Call(__e, ShenFunc(symcons), gen15929, gen15930)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15931)

			gen15932 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15933 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15932)

			gen15934 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15933)

			gen15935 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen15934)

			gen15936 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15937 := Call(__e, ShenFunc(symcons), gen15935, gen15936)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15937)

			gen15938 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen15939 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15938)

			gen15940 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15939)

			gen15941 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen15940)

			gen15942 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15943 := Call(__e, ShenFunc(symcons), gen15941, gen15942)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15943)

			gen15944 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15945 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15944)

			gen15946 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15945)

			gen15947 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen15946)

			gen15948 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15949 := Call(__e, ShenFunc(symcons), gen15947, gen15948)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15949)

			gen15950 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15951 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15950)

			gen15952 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen15951)

			gen15953 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen15952)

			gen15954 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15955 := Call(__e, ShenFunc(symcons), gen15953, gen15954)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15955)

			gen15956 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15957 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen15956)

			gen15958 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen15959 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15958)

			gen15960 := Call(__e, ShenFunc(symcons), gen15957, gen15959)

			gen15961 := Call(__e, ShenFunc(symcons), gen15960, Nil)

			gen15962 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15961)

			gen15963 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen15962)

			gen15964 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen15963)

			gen15965 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15966 := Call(__e, ShenFunc(symcons), gen15964, gen15965)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15966)

			gen15967 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15968 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15967)

			gen15969 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15968)

			gen15970 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen15969)

			gen15971 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15972 := Call(__e, ShenFunc(symcons), gen15970, gen15971)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15972)

			gen15973 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen15974 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15973)

			gen15975 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen15974)

			gen15976 := Call(__e, ShenFunc(symcons), gen15975, Nil)

			gen15977 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15976)

			gen15978 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen15977)

			gen15979 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen15978)

			gen15980 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15981 := Call(__e, ShenFunc(symcons), gen15979, gen15980)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15981)

			gen15982 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15983 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15982)

			gen15984 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15983)

			gen15985 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen15984)

			gen15986 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15987 := Call(__e, ShenFunc(symcons), gen15985, gen15986)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15987)

			gen15988 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15989 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15988)

			gen15990 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen15989)

			gen15991 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen15990)

			gen15992 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen15993 := Call(__e, ShenFunc(symcons), gen15991, gen15992)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen15993)

			gen15994 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen15995 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15994)

			gen15996 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen15995)

			gen15997 := Call(__e, ShenFunc(symcons), gen15996, Nil)

			gen15998 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen15997)

			gen15999 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen15998)

			gen16000 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen15999)

			gen16001 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16002 := Call(__e, ShenFunc(symcons), gen16000, gen16001)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16002)

			gen16003 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16004 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16003)

			gen16005 := Call(__e, ShenFunc(symcons), MakeSymbol("os"), gen16004)

			gen16006 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16007 := Call(__e, ShenFunc(symcons), gen16005, gen16006)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16007)

			gen16008 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16009 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16008)

			gen16010 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16009)

			gen16011 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen16010)

			gen16012 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16013 := Call(__e, ShenFunc(symcons), gen16011, gen16012)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16013)

			gen16014 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16015 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16014)

			gen16016 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen16015)

			gen16017 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16018 := Call(__e, ShenFunc(symcons), gen16016, gen16017)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16018)

			gen16019 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16020 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16019)

			gen16021 := Call(__e, ShenFunc(symcons), MakeSymbol("porters"), gen16020)

			gen16022 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16023 := Call(__e, ShenFunc(symcons), gen16021, gen16022)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16023)

			gen16024 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16025 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16024)

			gen16026 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16025)

			gen16027 := Call(__e, ShenFunc(symcons), gen16026, Nil)

			gen16028 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16027)

			gen16029 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16028)

			gen16030 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen16029)

			gen16031 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16032 := Call(__e, ShenFunc(symcons), gen16030, gen16031)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16032)

			gen16033 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen16034 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16033)

			gen16035 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16036 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16035)

			gen16037 := Call(__e, ShenFunc(symcons), gen16034, gen16036)

			gen16038 := Call(__e, ShenFunc(symcons), gen16037, Nil)

			gen16039 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16038)

			gen16040 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16039)

			gen16041 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen16040)

			gen16042 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16043 := Call(__e, ShenFunc(symcons), gen16041, gen16042)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16043)

			gen16044 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16045 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16044)

			gen16046 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16045)

			gen16047 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen16046)

			gen16048 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16049 := Call(__e, ShenFunc(symcons), gen16047, gen16048)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16049)

			gen16050 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16051 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16050)

			gen16052 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16051)

			gen16053 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16054 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16053)

			gen16055 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16054)

			gen16056 := Call(__e, ShenFunc(symcons), gen16055, Nil)

			gen16057 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16056)

			gen16058 := Call(__e, ShenFunc(symcons), gen16052, gen16057)

			gen16059 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen16058)

			gen16060 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16061 := Call(__e, ShenFunc(symcons), gen16059, gen16060)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16061)

			gen16062 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16063 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16062)

			gen16064 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16065 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16064)

			gen16066 := Call(__e, ShenFunc(symcons), gen16065, Nil)

			gen16067 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16066)

			gen16068 := Call(__e, ShenFunc(symcons), gen16063, gen16067)

			gen16069 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen16068)

			gen16070 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16071 := Call(__e, ShenFunc(symcons), gen16069, gen16070)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16071)

			gen16072 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16073 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16072)

			gen16074 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16073)

			gen16075 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.proc-nl"), gen16074)

			gen16076 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16077 := Call(__e, ShenFunc(symcons), gen16075, gen16076)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16077)

			gen16078 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16079 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16078)

			gen16080 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16079)

			gen16081 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16082 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16081)

			gen16083 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16082)

			gen16084 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16085 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen16084)

			gen16086 := Call(__e, ShenFunc(symcons), gen16083, gen16085)

			gen16087 := Call(__e, ShenFunc(symcons), gen16086, Nil)

			gen16088 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16087)

			gen16089 := Call(__e, ShenFunc(symcons), gen16080, gen16088)

			gen16090 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen16089)

			gen16091 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16092 := Call(__e, ShenFunc(symcons), gen16090, gen16091)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16092)

			gen16093 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16094 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16093)

			gen16095 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16094)

			gen16096 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen16095)

			gen16097 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16098 := Call(__e, ShenFunc(symcons), gen16096, gen16097)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16098)

			gen16099 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16100 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16099)

			gen16101 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16102 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16101)

			gen16103 := Call(__e, ShenFunc(symcons), gen16102, Nil)

			gen16104 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16103)

			gen16105 := Call(__e, ShenFunc(symcons), gen16100, gen16104)

			gen16106 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen16105)

			gen16107 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16108 := Call(__e, ShenFunc(symcons), gen16106, gen16107)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16108)

			gen16109 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen16110 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16109)

			gen16111 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16112 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16111)

			gen16113 := Call(__e, ShenFunc(symcons), gen16110, gen16112)

			gen16114 := Call(__e, ShenFunc(symcons), gen16113, Nil)

			gen16115 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16114)

			gen16116 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16115)

			gen16117 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.prhush"), gen16116)

			gen16118 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16119 := Call(__e, ShenFunc(symcons), gen16117, gen16118)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16119)

			gen16120 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen16121 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16120)

			gen16122 := Call(__e, ShenFunc(symcons), gen16121, Nil)

			gen16123 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16122)

			gen16124 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16123)

			gen16125 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen16124)

			gen16126 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16127 := Call(__e, ShenFunc(symcons), gen16125, gen16126)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16127)

			gen16128 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen16129 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16128)

			gen16130 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen16131 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16130)

			gen16132 := Call(__e, ShenFunc(symcons), gen16129, gen16131)

			gen16133 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen16132)

			gen16134 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16135 := Call(__e, ShenFunc(symcons), gen16133, gen16134)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16135)

			gen16136 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen16137 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16136)

			gen16138 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16139 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16138)

			gen16140 := Call(__e, ShenFunc(symcons), gen16137, gen16139)

			gen16141 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen16140)

			gen16142 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16143 := Call(__e, ShenFunc(symcons), gen16141, gen16142)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16143)

			gen16144 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16145 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16144)

			gen16146 := Call(__e, ShenFunc(symcons), gen16145, Nil)

			gen16147 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16146)

			gen16148 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16147)

			gen16149 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen16148)

			gen16150 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16151 := Call(__e, ShenFunc(symcons), gen16149, gen16150)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16151)

			gen16152 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16153 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16152)

			gen16154 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16153)

			gen16155 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen16154)

			gen16156 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16157 := Call(__e, ShenFunc(symcons), gen16155, gen16156)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16157)

			gen16158 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen16159 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16158)

			gen16160 := Call(__e, ShenFunc(symcons), gen16159, Nil)

			gen16161 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16160)

			gen16162 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16161)

			gen16163 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen16162)

			gen16164 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16165 := Call(__e, ShenFunc(symcons), gen16163, gen16164)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16165)

			gen16166 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen16167 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16166)

			gen16168 := Call(__e, ShenFunc(symcons), gen16167, Nil)

			gen16169 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16168)

			gen16170 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16169)

			gen16171 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen16170)

			gen16172 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16173 := Call(__e, ShenFunc(symcons), gen16171, gen16172)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16173)

			gen16174 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16175 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16174)

			gen16176 := Call(__e, ShenFunc(symcons), MakeSymbol("release"), gen16175)

			gen16177 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16178 := Call(__e, ShenFunc(symcons), gen16176, gen16177)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16178)

			gen16179 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16180 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16179)

			gen16181 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16182 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16181)

			gen16183 := Call(__e, ShenFunc(symcons), gen16182, Nil)

			gen16184 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16183)

			gen16185 := Call(__e, ShenFunc(symcons), gen16180, gen16184)

			gen16186 := Call(__e, ShenFunc(symcons), gen16185, Nil)

			gen16187 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16186)

			gen16188 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16187)

			gen16189 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen16188)

			gen16190 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16191 := Call(__e, ShenFunc(symcons), gen16189, gen16190)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16191)

			gen16192 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16193 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16192)

			gen16194 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16195 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16194)

			gen16196 := Call(__e, ShenFunc(symcons), gen16195, Nil)

			gen16197 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16196)

			gen16198 := Call(__e, ShenFunc(symcons), gen16193, gen16197)

			gen16199 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen16198)

			gen16200 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16201 := Call(__e, ShenFunc(symcons), gen16199, gen16200)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16201)

			gen16202 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16203 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16202)

			gen16204 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16203)

			gen16205 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen16204)

			gen16206 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16207 := Call(__e, ShenFunc(symcons), gen16205, gen16206)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16207)

			gen16208 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16209 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen16208)

			gen16210 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16209)

			gen16211 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16212 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16211)

			gen16213 := Call(__e, ShenFunc(symcons), gen16210, gen16212)

			gen16214 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen16213)

			gen16215 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16216 := Call(__e, ShenFunc(symcons), gen16214, gen16215)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16216)

			gen16217 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16218 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16217)

			gen16219 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16218)

			gen16220 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen16219)

			gen16221 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16222 := Call(__e, ShenFunc(symcons), gen16220, gen16221)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16222)

			gen16223 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16224 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16223)

			gen16225 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16224)

			gen16226 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen16225)

			gen16227 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16228 := Call(__e, ShenFunc(symcons), gen16226, gen16227)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16228)

			gen16229 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16230 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16229)

			gen16231 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16230)

			gen16232 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen16231)

			gen16233 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16234 := Call(__e, ShenFunc(symcons), gen16232, gen16233)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16234)

			gen16235 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen16236 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16235)

			gen16237 := Call(__e, ShenFunc(symcons), gen16236, Nil)

			gen16238 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16237)

			gen16239 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), gen16238)

			gen16240 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16241 := Call(__e, ShenFunc(symcons), gen16239, gen16240)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16241)

			gen16242 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen16243 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16242)

			gen16244 := Call(__e, ShenFunc(symcons), gen16243, Nil)

			gen16245 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16244)

			gen16246 := Call(__e, ShenFunc(symcons), MakeSymbol("sterror"), gen16245)

			gen16247 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16248 := Call(__e, ShenFunc(symcons), gen16246, gen16247)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16248)

			gen16249 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen16250 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16249)

			gen16251 := Call(__e, ShenFunc(symcons), gen16250, Nil)

			gen16252 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16251)

			gen16253 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), gen16252)

			gen16254 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16255 := Call(__e, ShenFunc(symcons), gen16253, gen16254)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16255)

			gen16256 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16257 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16256)

			gen16258 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16257)

			gen16259 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen16258)

			gen16260 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16261 := Call(__e, ShenFunc(symcons), gen16259, gen16260)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16261)

			gen16262 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16263 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16262)

			gen16264 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16263)

			gen16265 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen16264)

			gen16266 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16267 := Call(__e, ShenFunc(symcons), gen16265, gen16266)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16267)

			gen16268 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16269 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16268)

			gen16270 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16269)

			gen16271 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen16270)

			gen16272 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16273 := Call(__e, ShenFunc(symcons), gen16271, gen16272)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16273)

			gen16274 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16275 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16274)

			gen16276 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16275)

			gen16277 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen16276)

			gen16278 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16279 := Call(__e, ShenFunc(symcons), gen16277, gen16278)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16279)

			gen16280 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16281 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16280)

			gen16282 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16283 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16282)

			gen16284 := Call(__e, ShenFunc(symcons), gen16281, gen16283)

			gen16285 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen16284)

			gen16286 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16287 := Call(__e, ShenFunc(symcons), gen16285, gen16286)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16287)

			gen16288 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16289 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16288)

			gen16290 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16289)

			gen16291 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen16290)

			gen16292 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16293 := Call(__e, ShenFunc(symcons), gen16291, gen16292)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16293)

			gen16294 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16295 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16294)

			gen16296 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16295)

			gen16297 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen16296)

			gen16298 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16299 := Call(__e, ShenFunc(symcons), gen16297, gen16298)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16299)

			gen16300 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16301 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16300)

			gen16302 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16303 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16302)

			gen16304 := Call(__e, ShenFunc(symcons), gen16303, Nil)

			gen16305 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16304)

			gen16306 := Call(__e, ShenFunc(symcons), gen16301, gen16305)

			gen16307 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen16306)

			gen16308 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16309 := Call(__e, ShenFunc(symcons), gen16307, gen16308)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16309)

			gen16310 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16311 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16310)

			gen16312 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16311)

			gen16313 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen16312)

			gen16314 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16315 := Call(__e, ShenFunc(symcons), gen16313, gen16314)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16315)

			gen16316 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16317 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16316)

			gen16318 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16319 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16318)

			gen16320 := Call(__e, ShenFunc(symcons), gen16319, Nil)

			gen16321 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16320)

			gen16322 := Call(__e, ShenFunc(symcons), gen16317, gen16321)

			gen16323 := Call(__e, ShenFunc(symcons), MakeSymbol("tlv"), gen16322)

			gen16324 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16325 := Call(__e, ShenFunc(symcons), gen16323, gen16324)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16325)

			gen16326 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16327 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16326)

			gen16328 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16327)

			gen16329 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen16328)

			gen16330 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16331 := Call(__e, ShenFunc(symcons), gen16329, gen16330)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16331)

			gen16332 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16333 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16332)

			gen16334 := Call(__e, ShenFunc(symcons), MakeSymbol("tc?"), gen16333)

			gen16335 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16336 := Call(__e, ShenFunc(symcons), gen16334, gen16335)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16336)

			gen16337 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16338 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen16337)

			gen16339 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16340 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16339)

			gen16341 := Call(__e, ShenFunc(symcons), gen16338, gen16340)

			gen16342 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen16341)

			gen16343 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16344 := Call(__e, ShenFunc(symcons), gen16342, gen16343)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16344)

			gen16345 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16346 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16345)

			gen16347 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16346)

			gen16348 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen16347)

			gen16349 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16350 := Call(__e, ShenFunc(symcons), gen16348, gen16349)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16350)

			gen16351 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16352 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16351)

			gen16353 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen16352)

			gen16354 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16355 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16354)

			gen16356 := Call(__e, ShenFunc(symcons), gen16353, gen16355)

			gen16357 := Call(__e, ShenFunc(symcons), gen16356, Nil)

			gen16358 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16357)

			gen16359 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16358)

			gen16360 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen16359)

			gen16361 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16362 := Call(__e, ShenFunc(symcons), gen16360, gen16361)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16362)

			gen16363 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16364 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16363)

			gen16365 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16364)

			gen16366 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen16365)

			gen16367 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16368 := Call(__e, ShenFunc(symcons), gen16366, gen16367)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16368)

			gen16369 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16370 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16369)

			gen16371 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16370)

			gen16372 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen16371)

			gen16373 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16374 := Call(__e, ShenFunc(symcons), gen16372, gen16373)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16374)

			gen16375 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16376 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16375)

			gen16377 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16378 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16377)

			gen16379 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16380 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen16379)

			gen16381 := Call(__e, ShenFunc(symcons), gen16380, Nil)

			gen16382 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16381)

			gen16383 := Call(__e, ShenFunc(symcons), gen16378, gen16382)

			gen16384 := Call(__e, ShenFunc(symcons), gen16383, Nil)

			gen16385 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16384)

			gen16386 := Call(__e, ShenFunc(symcons), gen16376, gen16385)

			gen16387 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen16386)

			gen16388 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16389 := Call(__e, ShenFunc(symcons), gen16387, gen16388)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16389)

			gen16390 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16391 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16390)

			gen16392 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16391)

			gen16393 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen16394 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16393)

			gen16395 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16394)

			gen16396 := Call(__e, ShenFunc(symcons), gen16395, Nil)

			gen16397 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16396)

			gen16398 := Call(__e, ShenFunc(symcons), gen16392, gen16397)

			gen16399 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen16398)

			gen16400 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16401 := Call(__e, ShenFunc(symcons), gen16399, gen16400)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16401)

			gen16402 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16403 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16402)

			gen16404 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16403)

			gen16405 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen16404)

			gen16406 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16407 := Call(__e, ShenFunc(symcons), gen16405, gen16406)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16407)

			gen16408 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen16409 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16408)

			gen16410 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen16409)

			gen16411 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen16410)

			gen16412 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16413 := Call(__e, ShenFunc(symcons), gen16411, gen16412)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16413)

			gen16414 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16415 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16414)

			gen16416 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16415)

			gen16417 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen16416)

			gen16418 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16419 := Call(__e, ShenFunc(symcons), gen16417, gen16418)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16419)

			gen16420 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16421 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16420)

			gen16422 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16421)

			gen16423 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen16422)

			gen16424 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16425 := Call(__e, ShenFunc(symcons), gen16423, gen16424)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16425)

			gen16426 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen16427 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16426)

			gen16428 := Call(__e, ShenFunc(symcons), MakeSymbol("version"), gen16427)

			gen16429 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16430 := Call(__e, ShenFunc(symcons), gen16428, gen16429)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16430)

			gen16431 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen16432 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16431)

			gen16433 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16432)

			gen16434 := Call(__e, ShenFunc(symcons), gen16433, Nil)

			gen16435 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16434)

			gen16436 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16435)

			gen16437 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen16436)

			gen16438 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16439 := Call(__e, ShenFunc(symcons), gen16437, gen16438)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16439)

			gen16440 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen16441 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen16440)

			gen16442 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16443 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16442)

			gen16444 := Call(__e, ShenFunc(symcons), gen16441, gen16443)

			gen16445 := Call(__e, ShenFunc(symcons), gen16444, Nil)

			gen16446 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16445)

			gen16447 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16446)

			gen16448 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen16447)

			gen16449 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16450 := Call(__e, ShenFunc(symcons), gen16448, gen16449)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16450)

			gen16451 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16452 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16451)

			gen16453 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen16452)

			gen16454 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen16453)

			gen16455 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16456 := Call(__e, ShenFunc(symcons), gen16454, gen16455)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16456)

			gen16457 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16458 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16457)

			gen16459 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16458)

			gen16460 := Call(__e, ShenFunc(symcons), gen16459, Nil)

			gen16461 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16460)

			gen16462 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16461)

			gen16463 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen16462)

			gen16464 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16465 := Call(__e, ShenFunc(symcons), gen16463, gen16464)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16465)

			gen16466 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16467 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16466)

			gen16468 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16467)

			gen16469 := Call(__e, ShenFunc(symcons), gen16468, Nil)

			gen16470 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16469)

			gen16471 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16470)

			gen16472 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen16471)

			gen16473 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16474 := Call(__e, ShenFunc(symcons), gen16472, gen16473)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16474)

			gen16475 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16476 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16475)

			gen16477 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16476)

			gen16478 := Call(__e, ShenFunc(symcons), gen16477, Nil)

			gen16479 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16478)

			gen16480 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16479)

			gen16481 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen16480)

			gen16482 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16483 := Call(__e, ShenFunc(symcons), gen16481, gen16482)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16483)

			gen16484 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16485 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16484)

			gen16486 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16485)

			gen16487 := Call(__e, ShenFunc(symcons), gen16486, Nil)

			gen16488 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16487)

			gen16489 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16488)

			gen16490 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen16489)

			gen16491 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16492 := Call(__e, ShenFunc(symcons), gen16490, gen16491)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16492)

			gen16493 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16494 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16493)

			gen16495 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16494)

			gen16496 := Call(__e, ShenFunc(symcons), gen16495, Nil)

			gen16497 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16496)

			gen16498 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16497)

			gen16499 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen16498)

			gen16500 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16501 := Call(__e, ShenFunc(symcons), gen16499, gen16500)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16501)

			gen16502 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16503 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16502)

			gen16504 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16503)

			gen16505 := Call(__e, ShenFunc(symcons), gen16504, Nil)

			gen16506 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16505)

			gen16507 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16506)

			gen16508 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen16507)

			gen16509 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16510 := Call(__e, ShenFunc(symcons), gen16508, gen16509)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16510)

			gen16511 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16512 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16511)

			gen16513 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16512)

			gen16514 := Call(__e, ShenFunc(symcons), gen16513, Nil)

			gen16515 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16514)

			gen16516 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16515)

			gen16517 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen16516)

			gen16518 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16519 := Call(__e, ShenFunc(symcons), gen16517, gen16518)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16519)

			gen16520 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16521 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16520)

			gen16522 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16521)

			gen16523 := Call(__e, ShenFunc(symcons), gen16522, Nil)

			gen16524 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16523)

			gen16525 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16524)

			gen16526 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen16525)

			gen16527 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16528 := Call(__e, ShenFunc(symcons), gen16526, gen16527)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16528)

			gen16529 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen16530 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16529)

			gen16531 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16530)

			gen16532 := Call(__e, ShenFunc(symcons), gen16531, Nil)

			gen16533 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16532)

			gen16534 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen16533)

			gen16535 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen16534)

			gen16536 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16537 := Call(__e, ShenFunc(symcons), gen16535, gen16536)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16537)

			gen16538 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen16539 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16538)

			gen16540 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen16539)

			gen16541 := Call(__e, ShenFunc(symcons), gen16540, Nil)

			gen16542 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen16541)

			gen16543 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen16542)

			gen16544 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen16543)

			gen16545 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen16546 := Call(__e, ShenFunc(symcons), gen16544, gen16545)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen16546)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-signedfuncs"), gen16547)

		gen16820 := MakeNative(func(__e Evaluator) {
			gen16548 := MakeNative(func(__e Evaluator) {
				V3239 := __e.Get(1)
				_ = V3239
				__e.Return(MakeNative(func(__e Evaluator) {
					V3240 := __e.Get(1)
					_ = V3240
					__e.Return(MakeNative(func(__e Evaluator) {
						V3241 := __e.Get(1)
						_ = V3241
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1absvector_2), V3239, V3240, V3241)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16549 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-absvector?"), gen16548)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16549)

			gen16550 := MakeNative(func(__e Evaluator) {
				V3249 := __e.Get(1)
				_ = V3249
				__e.Return(MakeNative(func(__e Evaluator) {
					V3250 := __e.Get(1)
					_ = V3250
					__e.Return(MakeNative(func(__e Evaluator) {
						V3251 := __e.Get(1)
						_ = V3251
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1adjoin), V3249, V3250, V3251)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16551 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-adjoin"), gen16550)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16551)

			gen16552 := MakeNative(func(__e Evaluator) {
				V3259 := __e.Get(1)
				_ = V3259
				__e.Return(MakeNative(func(__e Evaluator) {
					V3260 := __e.Get(1)
					_ = V3260
					__e.Return(MakeNative(func(__e Evaluator) {
						V3261 := __e.Get(1)
						_ = V3261
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1and), V3259, V3260, V3261)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16553 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-and"), gen16552)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16553)

			gen16554 := MakeNative(func(__e Evaluator) {
				V3269 := __e.Get(1)
				_ = V3269
				__e.Return(MakeNative(func(__e Evaluator) {
					V3270 := __e.Get(1)
					_ = V3270
					__e.Return(MakeNative(func(__e Evaluator) {
						V3271 := __e.Get(1)
						_ = V3271
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4app), V3269, V3270, V3271)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16555 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.app"), gen16554)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16555)

			gen16556 := MakeNative(func(__e Evaluator) {
				V3279 := __e.Get(1)
				_ = V3279
				__e.Return(MakeNative(func(__e Evaluator) {
					V3280 := __e.Get(1)
					_ = V3280
					__e.Return(MakeNative(func(__e Evaluator) {
						V3281 := __e.Get(1)
						_ = V3281
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1append), V3279, V3280, V3281)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16557 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-append"), gen16556)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16557)

			gen16558 := MakeNative(func(__e Evaluator) {
				V3289 := __e.Get(1)
				_ = V3289
				__e.Return(MakeNative(func(__e Evaluator) {
					V3290 := __e.Get(1)
					_ = V3290
					__e.Return(MakeNative(func(__e Evaluator) {
						V3291 := __e.Get(1)
						_ = V3291
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1arity), V3289, V3290, V3291)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16559 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-arity"), gen16558)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16559)

			gen16560 := MakeNative(func(__e Evaluator) {
				V3299 := __e.Get(1)
				_ = V3299
				__e.Return(MakeNative(func(__e Evaluator) {
					V3300 := __e.Get(1)
					_ = V3300
					__e.Return(MakeNative(func(__e Evaluator) {
						V3301 := __e.Get(1)
						_ = V3301
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1assoc), V3299, V3300, V3301)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16561 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-assoc"), gen16560)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16561)

			gen16562 := MakeNative(func(__e Evaluator) {
				V3309 := __e.Get(1)
				_ = V3309
				__e.Return(MakeNative(func(__e Evaluator) {
					V3310 := __e.Get(1)
					_ = V3310
					__e.Return(MakeNative(func(__e Evaluator) {
						V3311 := __e.Get(1)
						_ = V3311
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1boolean_2), V3309, V3310, V3311)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16563 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-boolean?"), gen16562)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16563)

			gen16564 := MakeNative(func(__e Evaluator) {
				V3319 := __e.Get(1)
				_ = V3319
				__e.Return(MakeNative(func(__e Evaluator) {
					V3320 := __e.Get(1)
					_ = V3320
					__e.Return(MakeNative(func(__e Evaluator) {
						V3321 := __e.Get(1)
						_ = V3321
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1bound_2), V3319, V3320, V3321)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16565 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-bound?"), gen16564)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16565)

			gen16566 := MakeNative(func(__e Evaluator) {
				V3329 := __e.Get(1)
				_ = V3329
				__e.Return(MakeNative(func(__e Evaluator) {
					V3330 := __e.Get(1)
					_ = V3330
					__e.Return(MakeNative(func(__e Evaluator) {
						V3331 := __e.Get(1)
						_ = V3331
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1cd), V3329, V3330, V3331)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16567 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-cd"), gen16566)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16567)

			gen16568 := MakeNative(func(__e Evaluator) {
				V3339 := __e.Get(1)
				_ = V3339
				__e.Return(MakeNative(func(__e Evaluator) {
					V3340 := __e.Get(1)
					_ = V3340
					__e.Return(MakeNative(func(__e Evaluator) {
						V3341 := __e.Get(1)
						_ = V3341
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1close), V3339, V3340, V3341)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16569 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-close"), gen16568)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16569)

			gen16570 := MakeNative(func(__e Evaluator) {
				V3349 := __e.Get(1)
				_ = V3349
				__e.Return(MakeNative(func(__e Evaluator) {
					V3350 := __e.Get(1)
					_ = V3350
					__e.Return(MakeNative(func(__e Evaluator) {
						V3351 := __e.Get(1)
						_ = V3351
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1cn), V3349, V3350, V3351)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16571 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-cn"), gen16570)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16571)

			gen16572 := MakeNative(func(__e Evaluator) {
				V3359 := __e.Get(1)
				_ = V3359
				__e.Return(MakeNative(func(__e Evaluator) {
					V3360 := __e.Get(1)
					_ = V3360
					__e.Return(MakeNative(func(__e Evaluator) {
						V3361 := __e.Get(1)
						_ = V3361
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1compile), V3359, V3360, V3361)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16573 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-compile"), gen16572)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16573)

			gen16574 := MakeNative(func(__e Evaluator) {
				V3369 := __e.Get(1)
				_ = V3369
				__e.Return(MakeNative(func(__e Evaluator) {
					V3370 := __e.Get(1)
					_ = V3370
					__e.Return(MakeNative(func(__e Evaluator) {
						V3371 := __e.Get(1)
						_ = V3371
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1cons_2), V3369, V3370, V3371)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16575 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-cons?"), gen16574)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16575)

			gen16576 := MakeNative(func(__e Evaluator) {
				V3379 := __e.Get(1)
				_ = V3379
				__e.Return(MakeNative(func(__e Evaluator) {
					V3380 := __e.Get(1)
					_ = V3380
					__e.Return(MakeNative(func(__e Evaluator) {
						V3381 := __e.Get(1)
						_ = V3381
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1destroy), V3379, V3380, V3381)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16577 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-destroy"), gen16576)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16577)

			gen16578 := MakeNative(func(__e Evaluator) {
				V3389 := __e.Get(1)
				_ = V3389
				__e.Return(MakeNative(func(__e Evaluator) {
					V3390 := __e.Get(1)
					_ = V3390
					__e.Return(MakeNative(func(__e Evaluator) {
						V3391 := __e.Get(1)
						_ = V3391
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1difference), V3389, V3390, V3391)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16579 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-difference"), gen16578)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16579)

			gen16580 := MakeNative(func(__e Evaluator) {
				V3399 := __e.Get(1)
				_ = V3399
				__e.Return(MakeNative(func(__e Evaluator) {
					V3400 := __e.Get(1)
					_ = V3400
					__e.Return(MakeNative(func(__e Evaluator) {
						V3401 := __e.Get(1)
						_ = V3401
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1do), V3399, V3400, V3401)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16581 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-do"), gen16580)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16581)

			gen16582 := MakeNative(func(__e Evaluator) {
				V3409 := __e.Get(1)
				_ = V3409
				__e.Return(MakeNative(func(__e Evaluator) {
					V3410 := __e.Get(1)
					_ = V3410
					__e.Return(MakeNative(func(__e Evaluator) {
						V3411 := __e.Get(1)
						_ = V3411
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5e_6), V3409, V3410, V3411)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16583 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<e>"), gen16582)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16583)

			gen16584 := MakeNative(func(__e Evaluator) {
				V3419 := __e.Get(1)
				_ = V3419
				__e.Return(MakeNative(func(__e Evaluator) {
					V3420 := __e.Get(1)
					_ = V3420
					__e.Return(MakeNative(func(__e Evaluator) {
						V3421 := __e.Get(1)
						_ = V3421
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5_b_6), V3419, V3420, V3421)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16585 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<!>"), gen16584)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16585)

			gen16586 := MakeNative(func(__e Evaluator) {
				V3429 := __e.Get(1)
				_ = V3429
				__e.Return(MakeNative(func(__e Evaluator) {
					V3430 := __e.Get(1)
					_ = V3430
					__e.Return(MakeNative(func(__e Evaluator) {
						V3431 := __e.Get(1)
						_ = V3431
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1element_2), V3429, V3430, V3431)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16587 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-element?"), gen16586)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16587)

			gen16588 := MakeNative(func(__e Evaluator) {
				V3439 := __e.Get(1)
				_ = V3439
				__e.Return(MakeNative(func(__e Evaluator) {
					V3440 := __e.Get(1)
					_ = V3440
					__e.Return(MakeNative(func(__e Evaluator) {
						V3441 := __e.Get(1)
						_ = V3441
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1empty_2), V3439, V3440, V3441)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16589 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-empty?"), gen16588)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16589)

			gen16590 := MakeNative(func(__e Evaluator) {
				V3449 := __e.Get(1)
				_ = V3449
				__e.Return(MakeNative(func(__e Evaluator) {
					V3450 := __e.Get(1)
					_ = V3450
					__e.Return(MakeNative(func(__e Evaluator) {
						V3451 := __e.Get(1)
						_ = V3451
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1enable_1type_1theory), V3449, V3450, V3451)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16591 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-enable-type-theory"), gen16590)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16591)

			gen16592 := MakeNative(func(__e Evaluator) {
				V3459 := __e.Get(1)
				_ = V3459
				__e.Return(MakeNative(func(__e Evaluator) {
					V3460 := __e.Get(1)
					_ = V3460
					__e.Return(MakeNative(func(__e Evaluator) {
						V3461 := __e.Get(1)
						_ = V3461
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1external), V3459, V3460, V3461)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16593 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-external"), gen16592)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16593)

			gen16594 := MakeNative(func(__e Evaluator) {
				V3469 := __e.Get(1)
				_ = V3469
				__e.Return(MakeNative(func(__e Evaluator) {
					V3470 := __e.Get(1)
					_ = V3470
					__e.Return(MakeNative(func(__e Evaluator) {
						V3471 := __e.Get(1)
						_ = V3471
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1error_1to_1string), V3469, V3470, V3471)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16595 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-error-to-string"), gen16594)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16595)

			gen16596 := MakeNative(func(__e Evaluator) {
				V3479 := __e.Get(1)
				_ = V3479
				__e.Return(MakeNative(func(__e Evaluator) {
					V3480 := __e.Get(1)
					_ = V3480
					__e.Return(MakeNative(func(__e Evaluator) {
						V3481 := __e.Get(1)
						_ = V3481
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1explode), V3479, V3480, V3481)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16597 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-explode"), gen16596)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16597)

			gen16598 := MakeNative(func(__e Evaluator) {
				V3489 := __e.Get(1)
				_ = V3489
				__e.Return(MakeNative(func(__e Evaluator) {
					V3490 := __e.Get(1)
					_ = V3490
					__e.Return(MakeNative(func(__e Evaluator) {
						V3491 := __e.Get(1)
						_ = V3491
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fail), V3489, V3490, V3491)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16599 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fail"), gen16598)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16599)

			gen16600 := MakeNative(func(__e Evaluator) {
				V3499 := __e.Get(1)
				_ = V3499
				__e.Return(MakeNative(func(__e Evaluator) {
					V3500 := __e.Get(1)
					_ = V3500
					__e.Return(MakeNative(func(__e Evaluator) {
						V3501 := __e.Get(1)
						_ = V3501
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fail_1if), V3499, V3500, V3501)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16601 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fail-if"), gen16600)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16601)

			gen16602 := MakeNative(func(__e Evaluator) {
				V3509 := __e.Get(1)
				_ = V3509
				__e.Return(MakeNative(func(__e Evaluator) {
					V3510 := __e.Get(1)
					_ = V3510
					__e.Return(MakeNative(func(__e Evaluator) {
						V3511 := __e.Get(1)
						_ = V3511
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fix), V3509, V3510, V3511)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16603 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fix"), gen16602)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16603)

			gen16604 := MakeNative(func(__e Evaluator) {
				V3519 := __e.Get(1)
				_ = V3519
				__e.Return(MakeNative(func(__e Evaluator) {
					V3520 := __e.Get(1)
					_ = V3520
					__e.Return(MakeNative(func(__e Evaluator) {
						V3521 := __e.Get(1)
						_ = V3521
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1freeze), V3519, V3520, V3521)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16605 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-freeze"), gen16604)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16605)

			gen16606 := MakeNative(func(__e Evaluator) {
				V3529 := __e.Get(1)
				_ = V3529
				__e.Return(MakeNative(func(__e Evaluator) {
					V3530 := __e.Get(1)
					_ = V3530
					__e.Return(MakeNative(func(__e Evaluator) {
						V3531 := __e.Get(1)
						_ = V3531
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fst), V3529, V3530, V3531)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16607 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fst"), gen16606)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16607)

			gen16608 := MakeNative(func(__e Evaluator) {
				V3539 := __e.Get(1)
				_ = V3539
				__e.Return(MakeNative(func(__e Evaluator) {
					V3540 := __e.Get(1)
					_ = V3540
					__e.Return(MakeNative(func(__e Evaluator) {
						V3541 := __e.Get(1)
						_ = V3541
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1function), V3539, V3540, V3541)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16609 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-function"), gen16608)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16609)

			gen16610 := MakeNative(func(__e Evaluator) {
				V3549 := __e.Get(1)
				_ = V3549
				__e.Return(MakeNative(func(__e Evaluator) {
					V3550 := __e.Get(1)
					_ = V3550
					__e.Return(MakeNative(func(__e Evaluator) {
						V3551 := __e.Get(1)
						_ = V3551
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1gensym), V3549, V3550, V3551)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16611 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-gensym"), gen16610)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16611)

			gen16612 := MakeNative(func(__e Evaluator) {
				V3559 := __e.Get(1)
				_ = V3559
				__e.Return(MakeNative(func(__e Evaluator) {
					V3560 := __e.Get(1)
					_ = V3560
					__e.Return(MakeNative(func(__e Evaluator) {
						V3561 := __e.Get(1)
						_ = V3561
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5_1vector), V3559, V3560, V3561)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16613 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<-vector"), gen16612)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16613)

			gen16614 := MakeNative(func(__e Evaluator) {
				V3569 := __e.Get(1)
				_ = V3569
				__e.Return(MakeNative(func(__e Evaluator) {
					V3570 := __e.Get(1)
					_ = V3570
					__e.Return(MakeNative(func(__e Evaluator) {
						V3571 := __e.Get(1)
						_ = V3571
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1vector_1_6), V3569, V3570, V3571)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16615 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-vector->"), gen16614)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16615)

			gen16616 := MakeNative(func(__e Evaluator) {
				V3579 := __e.Get(1)
				_ = V3579
				__e.Return(MakeNative(func(__e Evaluator) {
					V3580 := __e.Get(1)
					_ = V3580
					__e.Return(MakeNative(func(__e Evaluator) {
						V3581 := __e.Get(1)
						_ = V3581
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1vector), V3579, V3580, V3581)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16617 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-vector"), gen16616)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16617)

			gen16618 := MakeNative(func(__e Evaluator) {
				V3589 := __e.Get(1)
				_ = V3589
				__e.Return(MakeNative(func(__e Evaluator) {
					V3590 := __e.Get(1)
					_ = V3590
					__e.Return(MakeNative(func(__e Evaluator) {
						V3591 := __e.Get(1)
						_ = V3591
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1get_1time), V3589, V3590, V3591)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16619 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-get-time"), gen16618)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16619)

			gen16620 := MakeNative(func(__e Evaluator) {
				V3599 := __e.Get(1)
				_ = V3599
				__e.Return(MakeNative(func(__e Evaluator) {
					V3600 := __e.Get(1)
					_ = V3600
					__e.Return(MakeNative(func(__e Evaluator) {
						V3601 := __e.Get(1)
						_ = V3601
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1hash), V3599, V3600, V3601)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16621 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-hash"), gen16620)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16621)

			gen16622 := MakeNative(func(__e Evaluator) {
				V3609 := __e.Get(1)
				_ = V3609
				__e.Return(MakeNative(func(__e Evaluator) {
					V3610 := __e.Get(1)
					_ = V3610
					__e.Return(MakeNative(func(__e Evaluator) {
						V3611 := __e.Get(1)
						_ = V3611
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1head), V3609, V3610, V3611)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16623 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-head"), gen16622)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16623)

			gen16624 := MakeNative(func(__e Evaluator) {
				V3619 := __e.Get(1)
				_ = V3619
				__e.Return(MakeNative(func(__e Evaluator) {
					V3620 := __e.Get(1)
					_ = V3620
					__e.Return(MakeNative(func(__e Evaluator) {
						V3621 := __e.Get(1)
						_ = V3621
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1hdv), V3619, V3620, V3621)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16625 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-hdv"), gen16624)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16625)

			gen16626 := MakeNative(func(__e Evaluator) {
				V3629 := __e.Get(1)
				_ = V3629
				__e.Return(MakeNative(func(__e Evaluator) {
					V3630 := __e.Get(1)
					_ = V3630
					__e.Return(MakeNative(func(__e Evaluator) {
						V3631 := __e.Get(1)
						_ = V3631
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1hdstr), V3629, V3630, V3631)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16627 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-hdstr"), gen16626)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16627)

			gen16628 := MakeNative(func(__e Evaluator) {
				V3639 := __e.Get(1)
				_ = V3639
				__e.Return(MakeNative(func(__e Evaluator) {
					V3640 := __e.Get(1)
					_ = V3640
					__e.Return(MakeNative(func(__e Evaluator) {
						V3641 := __e.Get(1)
						_ = V3641
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1if), V3639, V3640, V3641)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16629 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-if"), gen16628)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16629)

			gen16630 := MakeNative(func(__e Evaluator) {
				V3649 := __e.Get(1)
				_ = V3649
				__e.Return(MakeNative(func(__e Evaluator) {
					V3650 := __e.Get(1)
					_ = V3650
					__e.Return(MakeNative(func(__e Evaluator) {
						V3651 := __e.Get(1)
						_ = V3651
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1it), V3649, V3650, V3651)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16631 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-it"), gen16630)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16631)

			gen16632 := MakeNative(func(__e Evaluator) {
				V3659 := __e.Get(1)
				_ = V3659
				__e.Return(MakeNative(func(__e Evaluator) {
					V3660 := __e.Get(1)
					_ = V3660
					__e.Return(MakeNative(func(__e Evaluator) {
						V3661 := __e.Get(1)
						_ = V3661
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1implementation), V3659, V3660, V3661)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16633 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-implementation"), gen16632)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16633)

			gen16634 := MakeNative(func(__e Evaluator) {
				V3669 := __e.Get(1)
				_ = V3669
				__e.Return(MakeNative(func(__e Evaluator) {
					V3670 := __e.Get(1)
					_ = V3670
					__e.Return(MakeNative(func(__e Evaluator) {
						V3671 := __e.Get(1)
						_ = V3671
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1include), V3669, V3670, V3671)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16635 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-include"), gen16634)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16635)

			gen16636 := MakeNative(func(__e Evaluator) {
				V3679 := __e.Get(1)
				_ = V3679
				__e.Return(MakeNative(func(__e Evaluator) {
					V3680 := __e.Get(1)
					_ = V3680
					__e.Return(MakeNative(func(__e Evaluator) {
						V3681 := __e.Get(1)
						_ = V3681
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1include_1all_1but), V3679, V3680, V3681)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16637 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-include-all-but"), gen16636)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16637)

			gen16638 := MakeNative(func(__e Evaluator) {
				V3689 := __e.Get(1)
				_ = V3689
				__e.Return(MakeNative(func(__e Evaluator) {
					V3690 := __e.Get(1)
					_ = V3690
					__e.Return(MakeNative(func(__e Evaluator) {
						V3691 := __e.Get(1)
						_ = V3691
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1inferences), V3689, V3690, V3691)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16639 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-inferences"), gen16638)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16639)

			gen16640 := MakeNative(func(__e Evaluator) {
				V3699 := __e.Get(1)
				_ = V3699
				__e.Return(MakeNative(func(__e Evaluator) {
					V3700 := __e.Get(1)
					_ = V3700
					__e.Return(MakeNative(func(__e Evaluator) {
						V3701 := __e.Get(1)
						_ = V3701
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4insert), V3699, V3700, V3701)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16641 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.insert"), gen16640)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16641)

			gen16642 := MakeNative(func(__e Evaluator) {
				V3709 := __e.Get(1)
				_ = V3709
				__e.Return(MakeNative(func(__e Evaluator) {
					V3710 := __e.Get(1)
					_ = V3710
					__e.Return(MakeNative(func(__e Evaluator) {
						V3711 := __e.Get(1)
						_ = V3711
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1integer_2), V3709, V3710, V3711)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16643 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-integer?"), gen16642)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16643)

			gen16644 := MakeNative(func(__e Evaluator) {
				V3719 := __e.Get(1)
				_ = V3719
				__e.Return(MakeNative(func(__e Evaluator) {
					V3720 := __e.Get(1)
					_ = V3720
					__e.Return(MakeNative(func(__e Evaluator) {
						V3721 := __e.Get(1)
						_ = V3721
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1internal), V3719, V3720, V3721)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16645 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-internal"), gen16644)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16645)

			gen16646 := MakeNative(func(__e Evaluator) {
				V3729 := __e.Get(1)
				_ = V3729
				__e.Return(MakeNative(func(__e Evaluator) {
					V3730 := __e.Get(1)
					_ = V3730
					__e.Return(MakeNative(func(__e Evaluator) {
						V3731 := __e.Get(1)
						_ = V3731
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1intersection), V3729, V3730, V3731)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16647 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-intersection"), gen16646)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16647)

			gen16648 := MakeNative(func(__e Evaluator) {
				V3739 := __e.Get(1)
				_ = V3739
				__e.Return(MakeNative(func(__e Evaluator) {
					V3740 := __e.Get(1)
					_ = V3740
					__e.Return(MakeNative(func(__e Evaluator) {
						V3741 := __e.Get(1)
						_ = V3741
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1kill), V3739, V3740, V3741)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16649 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-kill"), gen16648)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16649)

			gen16650 := MakeNative(func(__e Evaluator) {
				V3749 := __e.Get(1)
				_ = V3749
				__e.Return(MakeNative(func(__e Evaluator) {
					V3750 := __e.Get(1)
					_ = V3750
					__e.Return(MakeNative(func(__e Evaluator) {
						V3751 := __e.Get(1)
						_ = V3751
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1language), V3749, V3750, V3751)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16651 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-language"), gen16650)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16651)

			gen16652 := MakeNative(func(__e Evaluator) {
				V3759 := __e.Get(1)
				_ = V3759
				__e.Return(MakeNative(func(__e Evaluator) {
					V3760 := __e.Get(1)
					_ = V3760
					__e.Return(MakeNative(func(__e Evaluator) {
						V3761 := __e.Get(1)
						_ = V3761
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1length), V3759, V3760, V3761)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16653 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-length"), gen16652)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16653)

			gen16654 := MakeNative(func(__e Evaluator) {
				V3769 := __e.Get(1)
				_ = V3769
				__e.Return(MakeNative(func(__e Evaluator) {
					V3770 := __e.Get(1)
					_ = V3770
					__e.Return(MakeNative(func(__e Evaluator) {
						V3771 := __e.Get(1)
						_ = V3771
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1limit), V3769, V3770, V3771)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16655 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-limit"), gen16654)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16655)

			gen16656 := MakeNative(func(__e Evaluator) {
				V3779 := __e.Get(1)
				_ = V3779
				__e.Return(MakeNative(func(__e Evaluator) {
					V3780 := __e.Get(1)
					_ = V3780
					__e.Return(MakeNative(func(__e Evaluator) {
						V3781 := __e.Get(1)
						_ = V3781
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1load), V3779, V3780, V3781)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16657 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-load"), gen16656)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16657)

			gen16658 := MakeNative(func(__e Evaluator) {
				V3789 := __e.Get(1)
				_ = V3789
				__e.Return(MakeNative(func(__e Evaluator) {
					V3790 := __e.Get(1)
					_ = V3790
					__e.Return(MakeNative(func(__e Evaluator) {
						V3791 := __e.Get(1)
						_ = V3791
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1map), V3789, V3790, V3791)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16659 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-map"), gen16658)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16659)

			gen16660 := MakeNative(func(__e Evaluator) {
				V3799 := __e.Get(1)
				_ = V3799
				__e.Return(MakeNative(func(__e Evaluator) {
					V3800 := __e.Get(1)
					_ = V3800
					__e.Return(MakeNative(func(__e Evaluator) {
						V3801 := __e.Get(1)
						_ = V3801
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1mapcan), V3799, V3800, V3801)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16661 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-mapcan"), gen16660)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16661)

			gen16662 := MakeNative(func(__e Evaluator) {
				V3809 := __e.Get(1)
				_ = V3809
				__e.Return(MakeNative(func(__e Evaluator) {
					V3810 := __e.Get(1)
					_ = V3810
					__e.Return(MakeNative(func(__e Evaluator) {
						V3811 := __e.Get(1)
						_ = V3811
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1maxinferences), V3809, V3810, V3811)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16663 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-maxinferences"), gen16662)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16663)

			gen16664 := MakeNative(func(__e Evaluator) {
				V3819 := __e.Get(1)
				_ = V3819
				__e.Return(MakeNative(func(__e Evaluator) {
					V3820 := __e.Get(1)
					_ = V3820
					__e.Return(MakeNative(func(__e Evaluator) {
						V3821 := __e.Get(1)
						_ = V3821
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1n_1_6string), V3819, V3820, V3821)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16665 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-n->string"), gen16664)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16665)

			gen16666 := MakeNative(func(__e Evaluator) {
				V3829 := __e.Get(1)
				_ = V3829
				__e.Return(MakeNative(func(__e Evaluator) {
					V3830 := __e.Get(1)
					_ = V3830
					__e.Return(MakeNative(func(__e Evaluator) {
						V3831 := __e.Get(1)
						_ = V3831
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1nl), V3829, V3830, V3831)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16667 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-nl"), gen16666)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16667)

			gen16668 := MakeNative(func(__e Evaluator) {
				V3839 := __e.Get(1)
				_ = V3839
				__e.Return(MakeNative(func(__e Evaluator) {
					V3840 := __e.Get(1)
					_ = V3840
					__e.Return(MakeNative(func(__e Evaluator) {
						V3841 := __e.Get(1)
						_ = V3841
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1not), V3839, V3840, V3841)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16669 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-not"), gen16668)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16669)

			gen16670 := MakeNative(func(__e Evaluator) {
				V3849 := __e.Get(1)
				_ = V3849
				__e.Return(MakeNative(func(__e Evaluator) {
					V3850 := __e.Get(1)
					_ = V3850
					__e.Return(MakeNative(func(__e Evaluator) {
						V3851 := __e.Get(1)
						_ = V3851
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1nth), V3849, V3850, V3851)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16671 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-nth"), gen16670)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16671)

			gen16672 := MakeNative(func(__e Evaluator) {
				V3859 := __e.Get(1)
				_ = V3859
				__e.Return(MakeNative(func(__e Evaluator) {
					V3860 := __e.Get(1)
					_ = V3860
					__e.Return(MakeNative(func(__e Evaluator) {
						V3861 := __e.Get(1)
						_ = V3861
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1number_2), V3859, V3860, V3861)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16673 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-number?"), gen16672)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16673)

			gen16674 := MakeNative(func(__e Evaluator) {
				V3869 := __e.Get(1)
				_ = V3869
				__e.Return(MakeNative(func(__e Evaluator) {
					V3870 := __e.Get(1)
					_ = V3870
					__e.Return(MakeNative(func(__e Evaluator) {
						V3871 := __e.Get(1)
						_ = V3871
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1occurrences), V3869, V3870, V3871)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16675 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-occurrences"), gen16674)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16675)

			gen16676 := MakeNative(func(__e Evaluator) {
				V3879 := __e.Get(1)
				_ = V3879
				__e.Return(MakeNative(func(__e Evaluator) {
					V3880 := __e.Get(1)
					_ = V3880
					__e.Return(MakeNative(func(__e Evaluator) {
						V3881 := __e.Get(1)
						_ = V3881
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1occurs_1check), V3879, V3880, V3881)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16677 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-occurs-check"), gen16676)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16677)

			gen16678 := MakeNative(func(__e Evaluator) {
				V3889 := __e.Get(1)
				_ = V3889
				__e.Return(MakeNative(func(__e Evaluator) {
					V3890 := __e.Get(1)
					_ = V3890
					__e.Return(MakeNative(func(__e Evaluator) {
						V3891 := __e.Get(1)
						_ = V3891
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1optimise), V3889, V3890, V3891)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16679 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-optimise"), gen16678)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16679)

			gen16680 := MakeNative(func(__e Evaluator) {
				V3899 := __e.Get(1)
				_ = V3899
				__e.Return(MakeNative(func(__e Evaluator) {
					V3900 := __e.Get(1)
					_ = V3900
					__e.Return(MakeNative(func(__e Evaluator) {
						V3901 := __e.Get(1)
						_ = V3901
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1or), V3899, V3900, V3901)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16681 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-or"), gen16680)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16681)

			gen16682 := MakeNative(func(__e Evaluator) {
				V3909 := __e.Get(1)
				_ = V3909
				__e.Return(MakeNative(func(__e Evaluator) {
					V3910 := __e.Get(1)
					_ = V3910
					__e.Return(MakeNative(func(__e Evaluator) {
						V3911 := __e.Get(1)
						_ = V3911
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1os), V3909, V3910, V3911)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16683 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-os"), gen16682)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16683)

			gen16684 := MakeNative(func(__e Evaluator) {
				V3919 := __e.Get(1)
				_ = V3919
				__e.Return(MakeNative(func(__e Evaluator) {
					V3920 := __e.Get(1)
					_ = V3920
					__e.Return(MakeNative(func(__e Evaluator) {
						V3921 := __e.Get(1)
						_ = V3921
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1package_2), V3919, V3920, V3921)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16685 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-package?"), gen16684)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16685)

			gen16686 := MakeNative(func(__e Evaluator) {
				V3929 := __e.Get(1)
				_ = V3929
				__e.Return(MakeNative(func(__e Evaluator) {
					V3930 := __e.Get(1)
					_ = V3930
					__e.Return(MakeNative(func(__e Evaluator) {
						V3931 := __e.Get(1)
						_ = V3931
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1port), V3929, V3930, V3931)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16687 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-port"), gen16686)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16687)

			gen16688 := MakeNative(func(__e Evaluator) {
				V3939 := __e.Get(1)
				_ = V3939
				__e.Return(MakeNative(func(__e Evaluator) {
					V3940 := __e.Get(1)
					_ = V3940
					__e.Return(MakeNative(func(__e Evaluator) {
						V3941 := __e.Get(1)
						_ = V3941
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1porters), V3939, V3940, V3941)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16689 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-porters"), gen16688)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16689)

			gen16690 := MakeNative(func(__e Evaluator) {
				V3949 := __e.Get(1)
				_ = V3949
				__e.Return(MakeNative(func(__e Evaluator) {
					V3950 := __e.Get(1)
					_ = V3950
					__e.Return(MakeNative(func(__e Evaluator) {
						V3951 := __e.Get(1)
						_ = V3951
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1pos), V3949, V3950, V3951)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16691 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-pos"), gen16690)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16691)

			gen16692 := MakeNative(func(__e Evaluator) {
				V3959 := __e.Get(1)
				_ = V3959
				__e.Return(MakeNative(func(__e Evaluator) {
					V3960 := __e.Get(1)
					_ = V3960
					__e.Return(MakeNative(func(__e Evaluator) {
						V3961 := __e.Get(1)
						_ = V3961
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1pr), V3959, V3960, V3961)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16693 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-pr"), gen16692)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16693)

			gen16694 := MakeNative(func(__e Evaluator) {
				V3969 := __e.Get(1)
				_ = V3969
				__e.Return(MakeNative(func(__e Evaluator) {
					V3970 := __e.Get(1)
					_ = V3970
					__e.Return(MakeNative(func(__e Evaluator) {
						V3971 := __e.Get(1)
						_ = V3971
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1print), V3969, V3970, V3971)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16695 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-print"), gen16694)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16695)

			gen16696 := MakeNative(func(__e Evaluator) {
				V3979 := __e.Get(1)
				_ = V3979
				__e.Return(MakeNative(func(__e Evaluator) {
					V3980 := __e.Get(1)
					_ = V3980
					__e.Return(MakeNative(func(__e Evaluator) {
						V3981 := __e.Get(1)
						_ = V3981
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1profile), V3979, V3980, V3981)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16697 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-profile"), gen16696)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16697)

			gen16698 := MakeNative(func(__e Evaluator) {
				V3989 := __e.Get(1)
				_ = V3989
				__e.Return(MakeNative(func(__e Evaluator) {
					V3990 := __e.Get(1)
					_ = V3990
					__e.Return(MakeNative(func(__e Evaluator) {
						V3991 := __e.Get(1)
						_ = V3991
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1preclude), V3989, V3990, V3991)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16699 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-preclude"), gen16698)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16699)

			gen16700 := MakeNative(func(__e Evaluator) {
				V3999 := __e.Get(1)
				_ = V3999
				__e.Return(MakeNative(func(__e Evaluator) {
					V4000 := __e.Get(1)
					_ = V4000
					__e.Return(MakeNative(func(__e Evaluator) {
						V4001 := __e.Get(1)
						_ = V4001
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4proc_1nl), V3999, V4000, V4001)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16701 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.proc-nl"), gen16700)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16701)

			gen16702 := MakeNative(func(__e Evaluator) {
				V4009 := __e.Get(1)
				_ = V4009
				__e.Return(MakeNative(func(__e Evaluator) {
					V4010 := __e.Get(1)
					_ = V4010
					__e.Return(MakeNative(func(__e Evaluator) {
						V4011 := __e.Get(1)
						_ = V4011
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1profile_1results), V4009, V4010, V4011)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16703 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-profile-results"), gen16702)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16703)

			gen16704 := MakeNative(func(__e Evaluator) {
				V4019 := __e.Get(1)
				_ = V4019
				__e.Return(MakeNative(func(__e Evaluator) {
					V4020 := __e.Get(1)
					_ = V4020
					__e.Return(MakeNative(func(__e Evaluator) {
						V4021 := __e.Get(1)
						_ = V4021
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1protect), V4019, V4020, V4021)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16705 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-protect"), gen16704)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16705)

			gen16706 := MakeNative(func(__e Evaluator) {
				V4029 := __e.Get(1)
				_ = V4029
				__e.Return(MakeNative(func(__e Evaluator) {
					V4030 := __e.Get(1)
					_ = V4030
					__e.Return(MakeNative(func(__e Evaluator) {
						V4031 := __e.Get(1)
						_ = V4031
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1preclude_1all_1but), V4029, V4030, V4031)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16707 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-preclude-all-but"), gen16706)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16707)

			gen16708 := MakeNative(func(__e Evaluator) {
				V4039 := __e.Get(1)
				_ = V4039
				__e.Return(MakeNative(func(__e Evaluator) {
					V4040 := __e.Get(1)
					_ = V4040
					__e.Return(MakeNative(func(__e Evaluator) {
						V4041 := __e.Get(1)
						_ = V4041
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4prhush), V4039, V4040, V4041)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16709 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.prhush"), gen16708)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16709)

			gen16710 := MakeNative(func(__e Evaluator) {
				V4049 := __e.Get(1)
				_ = V4049
				__e.Return(MakeNative(func(__e Evaluator) {
					V4050 := __e.Get(1)
					_ = V4050
					__e.Return(MakeNative(func(__e Evaluator) {
						V4051 := __e.Get(1)
						_ = V4051
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1ps), V4049, V4050, V4051)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16711 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-ps"), gen16710)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16711)

			gen16712 := MakeNative(func(__e Evaluator) {
				V4059 := __e.Get(1)
				_ = V4059
				__e.Return(MakeNative(func(__e Evaluator) {
					V4060 := __e.Get(1)
					_ = V4060
					__e.Return(MakeNative(func(__e Evaluator) {
						V4061 := __e.Get(1)
						_ = V4061
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read), V4059, V4060, V4061)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16713 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read"), gen16712)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16713)

			gen16714 := MakeNative(func(__e Evaluator) {
				V4069 := __e.Get(1)
				_ = V4069
				__e.Return(MakeNative(func(__e Evaluator) {
					V4070 := __e.Get(1)
					_ = V4070
					__e.Return(MakeNative(func(__e Evaluator) {
						V4071 := __e.Get(1)
						_ = V4071
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1byte), V4069, V4070, V4071)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16715 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-byte"), gen16714)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16715)

			gen16716 := MakeNative(func(__e Evaluator) {
				V4079 := __e.Get(1)
				_ = V4079
				__e.Return(MakeNative(func(__e Evaluator) {
					V4080 := __e.Get(1)
					_ = V4080
					__e.Return(MakeNative(func(__e Evaluator) {
						V4081 := __e.Get(1)
						_ = V4081
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1file_1as_1bytelist), V4079, V4080, V4081)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16717 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-file-as-bytelist"), gen16716)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16717)

			gen16718 := MakeNative(func(__e Evaluator) {
				V4089 := __e.Get(1)
				_ = V4089
				__e.Return(MakeNative(func(__e Evaluator) {
					V4090 := __e.Get(1)
					_ = V4090
					__e.Return(MakeNative(func(__e Evaluator) {
						V4091 := __e.Get(1)
						_ = V4091
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1file_1as_1string), V4089, V4090, V4091)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16719 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-file-as-string"), gen16718)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16719)

			gen16720 := MakeNative(func(__e Evaluator) {
				V4099 := __e.Get(1)
				_ = V4099
				__e.Return(MakeNative(func(__e Evaluator) {
					V4100 := __e.Get(1)
					_ = V4100
					__e.Return(MakeNative(func(__e Evaluator) {
						V4101 := __e.Get(1)
						_ = V4101
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1file), V4099, V4100, V4101)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16721 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-file"), gen16720)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16721)

			gen16722 := MakeNative(func(__e Evaluator) {
				V4109 := __e.Get(1)
				_ = V4109
				__e.Return(MakeNative(func(__e Evaluator) {
					V4110 := __e.Get(1)
					_ = V4110
					__e.Return(MakeNative(func(__e Evaluator) {
						V4111 := __e.Get(1)
						_ = V4111
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1from_1string), V4109, V4110, V4111)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16723 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-from-string"), gen16722)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16723)

			gen16724 := MakeNative(func(__e Evaluator) {
				V4119 := __e.Get(1)
				_ = V4119
				__e.Return(MakeNative(func(__e Evaluator) {
					V4120 := __e.Get(1)
					_ = V4120
					__e.Return(MakeNative(func(__e Evaluator) {
						V4121 := __e.Get(1)
						_ = V4121
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1release), V4119, V4120, V4121)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16725 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-release"), gen16724)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16725)

			gen16726 := MakeNative(func(__e Evaluator) {
				V4129 := __e.Get(1)
				_ = V4129
				__e.Return(MakeNative(func(__e Evaluator) {
					V4130 := __e.Get(1)
					_ = V4130
					__e.Return(MakeNative(func(__e Evaluator) {
						V4131 := __e.Get(1)
						_ = V4131
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1remove), V4129, V4130, V4131)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16727 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-remove"), gen16726)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16727)

			gen16728 := MakeNative(func(__e Evaluator) {
				V4139 := __e.Get(1)
				_ = V4139
				__e.Return(MakeNative(func(__e Evaluator) {
					V4140 := __e.Get(1)
					_ = V4140
					__e.Return(MakeNative(func(__e Evaluator) {
						V4141 := __e.Get(1)
						_ = V4141
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1reverse), V4139, V4140, V4141)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16729 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-reverse"), gen16728)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16729)

			gen16730 := MakeNative(func(__e Evaluator) {
				V4149 := __e.Get(1)
				_ = V4149
				__e.Return(MakeNative(func(__e Evaluator) {
					V4150 := __e.Get(1)
					_ = V4150
					__e.Return(MakeNative(func(__e Evaluator) {
						V4151 := __e.Get(1)
						_ = V4151
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1simple_1error), V4149, V4150, V4151)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16731 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-simple-error"), gen16730)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16731)

			gen16732 := MakeNative(func(__e Evaluator) {
				V4159 := __e.Get(1)
				_ = V4159
				__e.Return(MakeNative(func(__e Evaluator) {
					V4160 := __e.Get(1)
					_ = V4160
					__e.Return(MakeNative(func(__e Evaluator) {
						V4161 := __e.Get(1)
						_ = V4161
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1snd), V4159, V4160, V4161)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16733 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-snd"), gen16732)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16733)

			gen16734 := MakeNative(func(__e Evaluator) {
				V4169 := __e.Get(1)
				_ = V4169
				__e.Return(MakeNative(func(__e Evaluator) {
					V4170 := __e.Get(1)
					_ = V4170
					__e.Return(MakeNative(func(__e Evaluator) {
						V4171 := __e.Get(1)
						_ = V4171
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1specialise), V4169, V4170, V4171)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16735 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-specialise"), gen16734)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16735)

			gen16736 := MakeNative(func(__e Evaluator) {
				V4179 := __e.Get(1)
				_ = V4179
				__e.Return(MakeNative(func(__e Evaluator) {
					V4180 := __e.Get(1)
					_ = V4180
					__e.Return(MakeNative(func(__e Evaluator) {
						V4181 := __e.Get(1)
						_ = V4181
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1spy), V4179, V4180, V4181)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16737 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-spy"), gen16736)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16737)

			gen16738 := MakeNative(func(__e Evaluator) {
				V4189 := __e.Get(1)
				_ = V4189
				__e.Return(MakeNative(func(__e Evaluator) {
					V4190 := __e.Get(1)
					_ = V4190
					__e.Return(MakeNative(func(__e Evaluator) {
						V4191 := __e.Get(1)
						_ = V4191
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1step), V4189, V4190, V4191)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16739 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-step"), gen16738)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16739)

			gen16740 := MakeNative(func(__e Evaluator) {
				V4199 := __e.Get(1)
				_ = V4199
				__e.Return(MakeNative(func(__e Evaluator) {
					V4200 := __e.Get(1)
					_ = V4200
					__e.Return(MakeNative(func(__e Evaluator) {
						V4201 := __e.Get(1)
						_ = V4201
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1stinput), V4199, V4200, V4201)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16741 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-stinput"), gen16740)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16741)

			gen16742 := MakeNative(func(__e Evaluator) {
				V4209 := __e.Get(1)
				_ = V4209
				__e.Return(MakeNative(func(__e Evaluator) {
					V4210 := __e.Get(1)
					_ = V4210
					__e.Return(MakeNative(func(__e Evaluator) {
						V4211 := __e.Get(1)
						_ = V4211
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1sterror), V4209, V4210, V4211)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16743 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-sterror"), gen16742)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16743)

			gen16744 := MakeNative(func(__e Evaluator) {
				V4219 := __e.Get(1)
				_ = V4219
				__e.Return(MakeNative(func(__e Evaluator) {
					V4220 := __e.Get(1)
					_ = V4220
					__e.Return(MakeNative(func(__e Evaluator) {
						V4221 := __e.Get(1)
						_ = V4221
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1stoutput), V4219, V4220, V4221)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16745 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-stoutput"), gen16744)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16745)

			gen16746 := MakeNative(func(__e Evaluator) {
				V4229 := __e.Get(1)
				_ = V4229
				__e.Return(MakeNative(func(__e Evaluator) {
					V4230 := __e.Get(1)
					_ = V4230
					__e.Return(MakeNative(func(__e Evaluator) {
						V4231 := __e.Get(1)
						_ = V4231
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1string_2), V4229, V4230, V4231)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16747 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-string?"), gen16746)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16747)

			gen16748 := MakeNative(func(__e Evaluator) {
				V4239 := __e.Get(1)
				_ = V4239
				__e.Return(MakeNative(func(__e Evaluator) {
					V4240 := __e.Get(1)
					_ = V4240
					__e.Return(MakeNative(func(__e Evaluator) {
						V4241 := __e.Get(1)
						_ = V4241
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1str), V4239, V4240, V4241)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16749 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-str"), gen16748)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16749)

			gen16750 := MakeNative(func(__e Evaluator) {
				V4249 := __e.Get(1)
				_ = V4249
				__e.Return(MakeNative(func(__e Evaluator) {
					V4250 := __e.Get(1)
					_ = V4250
					__e.Return(MakeNative(func(__e Evaluator) {
						V4251 := __e.Get(1)
						_ = V4251
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1string_1_6n), V4249, V4250, V4251)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16751 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-string->n"), gen16750)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16751)

			gen16752 := MakeNative(func(__e Evaluator) {
				V4259 := __e.Get(1)
				_ = V4259
				__e.Return(MakeNative(func(__e Evaluator) {
					V4260 := __e.Get(1)
					_ = V4260
					__e.Return(MakeNative(func(__e Evaluator) {
						V4261 := __e.Get(1)
						_ = V4261
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1string_1_6symbol), V4259, V4260, V4261)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16753 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-string->symbol"), gen16752)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16753)

			gen16754 := MakeNative(func(__e Evaluator) {
				V4269 := __e.Get(1)
				_ = V4269
				__e.Return(MakeNative(func(__e Evaluator) {
					V4270 := __e.Get(1)
					_ = V4270
					__e.Return(MakeNative(func(__e Evaluator) {
						V4271 := __e.Get(1)
						_ = V4271
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1sum), V4269, V4270, V4271)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16755 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-sum"), gen16754)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16755)

			gen16756 := MakeNative(func(__e Evaluator) {
				V4279 := __e.Get(1)
				_ = V4279
				__e.Return(MakeNative(func(__e Evaluator) {
					V4280 := __e.Get(1)
					_ = V4280
					__e.Return(MakeNative(func(__e Evaluator) {
						V4281 := __e.Get(1)
						_ = V4281
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1symbol_2), V4279, V4280, V4281)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16757 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-symbol?"), gen16756)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16757)

			gen16758 := MakeNative(func(__e Evaluator) {
				V4289 := __e.Get(1)
				_ = V4289
				__e.Return(MakeNative(func(__e Evaluator) {
					V4290 := __e.Get(1)
					_ = V4290
					__e.Return(MakeNative(func(__e Evaluator) {
						V4291 := __e.Get(1)
						_ = V4291
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1systemf), V4289, V4290, V4291)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16759 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-systemf"), gen16758)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16759)

			gen16760 := MakeNative(func(__e Evaluator) {
				V4299 := __e.Get(1)
				_ = V4299
				__e.Return(MakeNative(func(__e Evaluator) {
					V4300 := __e.Get(1)
					_ = V4300
					__e.Return(MakeNative(func(__e Evaluator) {
						V4301 := __e.Get(1)
						_ = V4301
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tail), V4299, V4300, V4301)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16761 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tail"), gen16760)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16761)

			gen16762 := MakeNative(func(__e Evaluator) {
				V4309 := __e.Get(1)
				_ = V4309
				__e.Return(MakeNative(func(__e Evaluator) {
					V4310 := __e.Get(1)
					_ = V4310
					__e.Return(MakeNative(func(__e Evaluator) {
						V4311 := __e.Get(1)
						_ = V4311
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tlstr), V4309, V4310, V4311)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16763 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tlstr"), gen16762)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16763)

			gen16764 := MakeNative(func(__e Evaluator) {
				V4319 := __e.Get(1)
				_ = V4319
				__e.Return(MakeNative(func(__e Evaluator) {
					V4320 := __e.Get(1)
					_ = V4320
					__e.Return(MakeNative(func(__e Evaluator) {
						V4321 := __e.Get(1)
						_ = V4321
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tlv), V4319, V4320, V4321)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16765 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tlv"), gen16764)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16765)

			gen16766 := MakeNative(func(__e Evaluator) {
				V4329 := __e.Get(1)
				_ = V4329
				__e.Return(MakeNative(func(__e Evaluator) {
					V4330 := __e.Get(1)
					_ = V4330
					__e.Return(MakeNative(func(__e Evaluator) {
						V4331 := __e.Get(1)
						_ = V4331
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tc), V4329, V4330, V4331)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16767 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tc"), gen16766)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16767)

			gen16768 := MakeNative(func(__e Evaluator) {
				V4339 := __e.Get(1)
				_ = V4339
				__e.Return(MakeNative(func(__e Evaluator) {
					V4340 := __e.Get(1)
					_ = V4340
					__e.Return(MakeNative(func(__e Evaluator) {
						V4341 := __e.Get(1)
						_ = V4341
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tc_2), V4339, V4340, V4341)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16769 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tc?"), gen16768)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16769)

			gen16770 := MakeNative(func(__e Evaluator) {
				V4349 := __e.Get(1)
				_ = V4349
				__e.Return(MakeNative(func(__e Evaluator) {
					V4350 := __e.Get(1)
					_ = V4350
					__e.Return(MakeNative(func(__e Evaluator) {
						V4351 := __e.Get(1)
						_ = V4351
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1thaw), V4349, V4350, V4351)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16771 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-thaw"), gen16770)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16771)

			gen16772 := MakeNative(func(__e Evaluator) {
				V4359 := __e.Get(1)
				_ = V4359
				__e.Return(MakeNative(func(__e Evaluator) {
					V4360 := __e.Get(1)
					_ = V4360
					__e.Return(MakeNative(func(__e Evaluator) {
						V4361 := __e.Get(1)
						_ = V4361
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1track), V4359, V4360, V4361)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16773 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-track"), gen16772)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16773)

			gen16774 := MakeNative(func(__e Evaluator) {
				V4369 := __e.Get(1)
				_ = V4369
				__e.Return(MakeNative(func(__e Evaluator) {
					V4370 := __e.Get(1)
					_ = V4370
					__e.Return(MakeNative(func(__e Evaluator) {
						V4371 := __e.Get(1)
						_ = V4371
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1trap_1error), V4369, V4370, V4371)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16775 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-trap-error"), gen16774)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16775)

			gen16776 := MakeNative(func(__e Evaluator) {
				V4379 := __e.Get(1)
				_ = V4379
				__e.Return(MakeNative(func(__e Evaluator) {
					V4380 := __e.Get(1)
					_ = V4380
					__e.Return(MakeNative(func(__e Evaluator) {
						V4381 := __e.Get(1)
						_ = V4381
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tuple_2), V4379, V4380, V4381)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16777 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tuple?"), gen16776)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16777)

			gen16778 := MakeNative(func(__e Evaluator) {
				V4389 := __e.Get(1)
				_ = V4389
				__e.Return(MakeNative(func(__e Evaluator) {
					V4390 := __e.Get(1)
					_ = V4390
					__e.Return(MakeNative(func(__e Evaluator) {
						V4391 := __e.Get(1)
						_ = V4391
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1undefmacro), V4389, V4390, V4391)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16779 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-undefmacro"), gen16778)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16779)

			gen16780 := MakeNative(func(__e Evaluator) {
				V4399 := __e.Get(1)
				_ = V4399
				__e.Return(MakeNative(func(__e Evaluator) {
					V4400 := __e.Get(1)
					_ = V4400
					__e.Return(MakeNative(func(__e Evaluator) {
						V4401 := __e.Get(1)
						_ = V4401
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1union), V4399, V4400, V4401)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16781 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-union"), gen16780)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16781)

			gen16782 := MakeNative(func(__e Evaluator) {
				V4409 := __e.Get(1)
				_ = V4409
				__e.Return(MakeNative(func(__e Evaluator) {
					V4410 := __e.Get(1)
					_ = V4410
					__e.Return(MakeNative(func(__e Evaluator) {
						V4411 := __e.Get(1)
						_ = V4411
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1unprofile), V4409, V4410, V4411)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16783 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-unprofile"), gen16782)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16783)

			gen16784 := MakeNative(func(__e Evaluator) {
				V4419 := __e.Get(1)
				_ = V4419
				__e.Return(MakeNative(func(__e Evaluator) {
					V4420 := __e.Get(1)
					_ = V4420
					__e.Return(MakeNative(func(__e Evaluator) {
						V4421 := __e.Get(1)
						_ = V4421
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1untrack), V4419, V4420, V4421)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16785 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-untrack"), gen16784)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16785)

			gen16786 := MakeNative(func(__e Evaluator) {
				V4429 := __e.Get(1)
				_ = V4429
				__e.Return(MakeNative(func(__e Evaluator) {
					V4430 := __e.Get(1)
					_ = V4430
					__e.Return(MakeNative(func(__e Evaluator) {
						V4431 := __e.Get(1)
						_ = V4431
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1unspecialise), V4429, V4430, V4431)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16787 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-unspecialise"), gen16786)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16787)

			gen16788 := MakeNative(func(__e Evaluator) {
				V4439 := __e.Get(1)
				_ = V4439
				__e.Return(MakeNative(func(__e Evaluator) {
					V4440 := __e.Get(1)
					_ = V4440
					__e.Return(MakeNative(func(__e Evaluator) {
						V4441 := __e.Get(1)
						_ = V4441
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1variable_2), V4439, V4440, V4441)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16789 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-variable?"), gen16788)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16789)

			gen16790 := MakeNative(func(__e Evaluator) {
				V4449 := __e.Get(1)
				_ = V4449
				__e.Return(MakeNative(func(__e Evaluator) {
					V4450 := __e.Get(1)
					_ = V4450
					__e.Return(MakeNative(func(__e Evaluator) {
						V4451 := __e.Get(1)
						_ = V4451
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1vector_2), V4449, V4450, V4451)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16791 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-vector?"), gen16790)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16791)

			gen16792 := MakeNative(func(__e Evaluator) {
				V4459 := __e.Get(1)
				_ = V4459
				__e.Return(MakeNative(func(__e Evaluator) {
					V4460 := __e.Get(1)
					_ = V4460
					__e.Return(MakeNative(func(__e Evaluator) {
						V4461 := __e.Get(1)
						_ = V4461
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1version), V4459, V4460, V4461)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16793 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-version"), gen16792)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16793)

			gen16794 := MakeNative(func(__e Evaluator) {
				V4469 := __e.Get(1)
				_ = V4469
				__e.Return(MakeNative(func(__e Evaluator) {
					V4470 := __e.Get(1)
					_ = V4470
					__e.Return(MakeNative(func(__e Evaluator) {
						V4471 := __e.Get(1)
						_ = V4471
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1write_1to_1file), V4469, V4470, V4471)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16795 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-write-to-file"), gen16794)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16795)

			gen16796 := MakeNative(func(__e Evaluator) {
				V4479 := __e.Get(1)
				_ = V4479
				__e.Return(MakeNative(func(__e Evaluator) {
					V4480 := __e.Get(1)
					_ = V4480
					__e.Return(MakeNative(func(__e Evaluator) {
						V4481 := __e.Get(1)
						_ = V4481
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1write_1byte), V4479, V4480, V4481)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16797 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-write-byte"), gen16796)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16797)

			gen16798 := MakeNative(func(__e Evaluator) {
				V4489 := __e.Get(1)
				_ = V4489
				__e.Return(MakeNative(func(__e Evaluator) {
					V4490 := __e.Get(1)
					_ = V4490
					__e.Return(MakeNative(func(__e Evaluator) {
						V4491 := __e.Get(1)
						_ = V4491
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1y_1or_1n_2), V4489, V4490, V4491)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16799 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-y-or-n?"), gen16798)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16799)

			gen16800 := MakeNative(func(__e Evaluator) {
				V4499 := __e.Get(1)
				_ = V4499
				__e.Return(MakeNative(func(__e Evaluator) {
					V4500 := __e.Get(1)
					_ = V4500
					__e.Return(MakeNative(func(__e Evaluator) {
						V4501 := __e.Get(1)
						_ = V4501
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_6), V4499, V4500, V4501)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16801 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of->"), gen16800)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16801)

			gen16802 := MakeNative(func(__e Evaluator) {
				V4509 := __e.Get(1)
				_ = V4509
				__e.Return(MakeNative(func(__e Evaluator) {
					V4510 := __e.Get(1)
					_ = V4510
					__e.Return(MakeNative(func(__e Evaluator) {
						V4511 := __e.Get(1)
						_ = V4511
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5), V4509, V4510, V4511)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16803 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<"), gen16802)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16803)

			gen16804 := MakeNative(func(__e Evaluator) {
				V4519 := __e.Get(1)
				_ = V4519
				__e.Return(MakeNative(func(__e Evaluator) {
					V4520 := __e.Get(1)
					_ = V4520
					__e.Return(MakeNative(func(__e Evaluator) {
						V4521 := __e.Get(1)
						_ = V4521
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_6_a), V4519, V4520, V4521)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16805 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of->="), gen16804)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16805)

			gen16806 := MakeNative(func(__e Evaluator) {
				V4529 := __e.Get(1)
				_ = V4529
				__e.Return(MakeNative(func(__e Evaluator) {
					V4530 := __e.Get(1)
					_ = V4530
					__e.Return(MakeNative(func(__e Evaluator) {
						V4531 := __e.Get(1)
						_ = V4531
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5_a), V4529, V4530, V4531)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16807 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<="), gen16806)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16807)

			gen16808 := MakeNative(func(__e Evaluator) {
				V4539 := __e.Get(1)
				_ = V4539
				__e.Return(MakeNative(func(__e Evaluator) {
					V4540 := __e.Get(1)
					_ = V4540
					__e.Return(MakeNative(func(__e Evaluator) {
						V4541 := __e.Get(1)
						_ = V4541
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_a), V4539, V4540, V4541)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16809 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-="), gen16808)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16809)

			gen16810 := MakeNative(func(__e Evaluator) {
				V4549 := __e.Get(1)
				_ = V4549
				__e.Return(MakeNative(func(__e Evaluator) {
					V4550 := __e.Get(1)
					_ = V4550
					__e.Return(MakeNative(func(__e Evaluator) {
						V4551 := __e.Get(1)
						_ = V4551
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_7), V4549, V4550, V4551)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16811 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-+"), gen16810)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16811)

			gen16812 := MakeNative(func(__e Evaluator) {
				V4559 := __e.Get(1)
				_ = V4559
				__e.Return(MakeNative(func(__e Evaluator) {
					V4560 := __e.Get(1)
					_ = V4560
					__e.Return(MakeNative(func(__e Evaluator) {
						V4561 := __e.Get(1)
						_ = V4561
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_c), V4559, V4560, V4561)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16813 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-/"), gen16812)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16813)

			gen16814 := MakeNative(func(__e Evaluator) {
				V4569 := __e.Get(1)
				_ = V4569
				__e.Return(MakeNative(func(__e Evaluator) {
					V4570 := __e.Get(1)
					_ = V4570
					__e.Return(MakeNative(func(__e Evaluator) {
						V4571 := __e.Get(1)
						_ = V4571
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_1), V4569, V4570, V4571)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16815 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of--"), gen16814)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16815)

			gen16816 := MakeNative(func(__e Evaluator) {
				V4579 := __e.Get(1)
				_ = V4579
				__e.Return(MakeNative(func(__e Evaluator) {
					V4580 := __e.Get(1)
					_ = V4580
					__e.Return(MakeNative(func(__e Evaluator) {
						V4581 := __e.Get(1)
						_ = V4581
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_d), V4579, V4580, V4581)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16817 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-*"), gen16816)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16817)

			gen16818 := MakeNative(func(__e Evaluator) {
				V4589 := __e.Get(1)
				_ = V4589
				__e.Return(MakeNative(func(__e Evaluator) {
					V4590 := __e.Get(1)
					_ = V4590
					__e.Return(MakeNative(func(__e Evaluator) {
						V4591 := __e.Get(1)
						_ = V4591
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_a_a), V4589, V4590, V4591)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16819 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-=="), gen16818)

			__e.TailApply(ShenFunc(symshen_4set_1lambda_1form_1entry), gen16819)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-signedfunc-lambda-forms"), gen16820)

		gen17122 := MakeNative(func(__e Evaluator) {
			gen16821 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4datatype_1error), X)

				return
			}, 1)
			gen16822 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.datatype-error"), gen16821)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16822)

			gen16823 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4tuple), X)

				return
			}, 1)
			gen16824 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tuple"), gen16823)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16824)

			gen16825 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4pvar), X)

				return
			}, 1)
			gen16826 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pvar"), gen16825)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16826)

			gen16827 := MakeNative(func(__e Evaluator) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(ShenFunc(symshen_4dictionary), X)

				return
			}, 1)
			gen16828 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dictionary"), gen16827)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16828)

			gen16829 := MakeNative(func(__e Evaluator) {
				V476 := __e.Get(1)
				_ = V476
				__e.Return(MakeNative(func(__e Evaluator) {
					V477 := __e.Get(1)
					_ = V477
					__e.TailApply(ShenFunc(sym_8v), V476, V477)

					return
				}, 1))
				return
			}, 1)
			gen16830 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen16829)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16830)

			gen16831 := MakeNative(func(__e Evaluator) {
				V478 := __e.Get(1)
				_ = V478
				__e.Return(MakeNative(func(__e Evaluator) {
					V479 := __e.Get(1)
					_ = V479
					__e.TailApply(ShenFunc(sym_8p), V478, V479)

					return
				}, 1))
				return
			}, 1)
			gen16832 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen16831)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16832)

			gen16833 := MakeNative(func(__e Evaluator) {
				V480 := __e.Get(1)
				_ = V480
				__e.Return(MakeNative(func(__e Evaluator) {
					V481 := __e.Get(1)
					_ = V481
					__e.TailApply(ShenFunc(sym_8s), V480, V481)

					return
				}, 1))
				return
			}, 1)
			gen16834 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen16833)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16834)

			gen16835 := MakeNative(func(__e Evaluator) {
				V482 := __e.Get(1)
				_ = V482
				__e.TailApply(ShenFunc(sym_5e_6), V482)

				return
			}, 1)
			gen16836 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen16835)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16836)

			gen16837 := MakeNative(func(__e Evaluator) {
				V483 := __e.Get(1)
				_ = V483
				__e.TailApply(ShenFunc(sym_5_b_6), V483)

				return
			}, 1)
			gen16838 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen16837)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16838)

			gen16839 := MakeNative(func(__e Evaluator) {
				V484 := __e.Get(1)
				_ = V484
				__e.Return(MakeNative(func(__e Evaluator) {
					V485 := __e.Get(1)
					_ = V485
					__e.TailApply(ShenFunc(sym_a_a), V484, V485)

					return
				}, 1))
				return
			}, 1)
			gen16840 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen16839)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16840)

			gen16841 := MakeNative(func(__e Evaluator) {
				V486 := __e.Get(1)
				_ = V486
				__e.Return(MakeNative(func(__e Evaluator) {
					V487 := __e.Get(1)
					_ = V487
					__e.TailApply(ShenFunc(sym_a), V486, V487)

					return
				}, 1))
				return
			}, 1)
			gen16842 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen16841)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16842)

			gen16843 := MakeNative(func(__e Evaluator) {
				V488 := __e.Get(1)
				_ = V488
				__e.Return(MakeNative(func(__e Evaluator) {
					V489 := __e.Get(1)
					_ = V489
					__e.TailApply(ShenFunc(sym_6_a), V488, V489)

					return
				}, 1))
				return
			}, 1)
			gen16844 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen16843)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16844)

			gen16845 := MakeNative(func(__e Evaluator) {
				V490 := __e.Get(1)
				_ = V490
				__e.Return(MakeNative(func(__e Evaluator) {
					V491 := __e.Get(1)
					_ = V491
					__e.TailApply(ShenFunc(sym_6), V490, V491)

					return
				}, 1))
				return
			}, 1)
			gen16846 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen16845)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16846)

			gen16847 := MakeNative(func(__e Evaluator) {
				V492 := __e.Get(1)
				_ = V492
				__e.Return(MakeNative(func(__e Evaluator) {
					V493 := __e.Get(1)
					_ = V493
					__e.TailApply(ShenFunc(sym_1), V492, V493)

					return
				}, 1))
				return
			}, 1)
			gen16848 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen16847)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16848)

			gen16849 := MakeNative(func(__e Evaluator) {
				V494 := __e.Get(1)
				_ = V494
				__e.Return(MakeNative(func(__e Evaluator) {
					V495 := __e.Get(1)
					_ = V495
					__e.TailApply(ShenFunc(sym_c), V494, V495)

					return
				}, 1))
				return
			}, 1)
			gen16850 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen16849)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16850)

			gen16851 := MakeNative(func(__e Evaluator) {
				V496 := __e.Get(1)
				_ = V496
				__e.Return(MakeNative(func(__e Evaluator) {
					V497 := __e.Get(1)
					_ = V497
					__e.TailApply(ShenFunc(sym_d), V496, V497)

					return
				}, 1))
				return
			}, 1)
			gen16852 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen16851)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16852)

			gen16853 := MakeNative(func(__e Evaluator) {
				V498 := __e.Get(1)
				_ = V498
				__e.Return(MakeNative(func(__e Evaluator) {
					V499 := __e.Get(1)
					_ = V499
					__e.TailApply(ShenFunc(sym_7), V498, V499)

					return
				}, 1))
				return
			}, 1)
			gen16854 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen16853)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16854)

			gen16855 := MakeNative(func(__e Evaluator) {
				V500 := __e.Get(1)
				_ = V500
				__e.Return(MakeNative(func(__e Evaluator) {
					V501 := __e.Get(1)
					_ = V501
					__e.TailApply(ShenFunc(sym_5_a), V500, V501)

					return
				}, 1))
				return
			}, 1)
			gen16856 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen16855)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16856)

			gen16857 := MakeNative(func(__e Evaluator) {
				V502 := __e.Get(1)
				_ = V502
				__e.Return(MakeNative(func(__e Evaluator) {
					V503 := __e.Get(1)
					_ = V503
					__e.TailApply(ShenFunc(sym_5), V502, V503)

					return
				}, 1))
				return
			}, 1)
			gen16858 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen16857)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16858)

			gen16859 := MakeNative(func(__e Evaluator) {
				V504 := __e.Get(1)
				_ = V504
				__e.TailApply(ShenFunc(symy_1or_1n_2), V504)

				return
			}, 1)
			gen16860 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen16859)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16860)

			gen16861 := MakeNative(func(__e Evaluator) {
				V505 := __e.Get(1)
				_ = V505
				__e.Return(MakeNative(func(__e Evaluator) {
					V506 := __e.Get(1)
					_ = V506
					__e.TailApply(ShenFunc(symwrite_1to_1file), V505, V506)

					return
				}, 1))
				return
			}, 1)
			gen16862 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen16861)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16862)

			gen16863 := MakeNative(func(__e Evaluator) {
				V507 := __e.Get(1)
				_ = V507
				__e.Return(MakeNative(func(__e Evaluator) {
					V508 := __e.Get(1)
					_ = V508
					__e.TailApply(ShenFunc(symwrite_1byte), V507, V508)

					return
				}, 1))
				return
			}, 1)
			gen16864 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen16863)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16864)

			gen16865 := MakeNative(func(__e Evaluator) {
				V509 := __e.Get(1)
				_ = V509
				__e.TailApply(ShenFunc(symvariable_2), V509)

				return
			}, 1)
			gen16866 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen16865)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16866)

			gen16867 := MakeNative(func(__e Evaluator) {
				V510 := __e.Get(1)
				_ = V510
				__e.TailApply(ShenFunc(symvalue), V510)

				return
			}, 1)
			gen16868 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen16867)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16868)

			gen16869 := MakeNative(func(__e Evaluator) {
				V511 := __e.Get(1)
				_ = V511
				__e.Return(MakeNative(func(__e Evaluator) {
					V512 := __e.Get(1)
					_ = V512
					__e.Return(MakeNative(func(__e Evaluator) {
						V513 := __e.Get(1)
						_ = V513
						__e.TailApply(ShenFunc(symvector_1_6), V511, V512, V513)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16870 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen16869)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16870)

			gen16871 := MakeNative(func(__e Evaluator) {
				V514 := __e.Get(1)
				_ = V514
				__e.Return(MakeNative(func(__e Evaluator) {
					V515 := __e.Get(1)
					_ = V515
					__e.TailApply(ShenFunc(sym_5_1vector), V514, V515)

					return
				}, 1))
				return
			}, 1)
			gen16872 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen16871)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16872)

			gen16873 := MakeNative(func(__e Evaluator) {
				V516 := __e.Get(1)
				_ = V516
				__e.TailApply(ShenFunc(symvector), V516)

				return
			}, 1)
			gen16874 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen16873)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16874)

			gen16875 := MakeNative(func(__e Evaluator) {
				V517 := __e.Get(1)
				_ = V517
				__e.TailApply(ShenFunc(symvector_2), V517)

				return
			}, 1)
			gen16876 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen16875)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16876)

			gen16877 := MakeNative(func(__e Evaluator) {
				V518 := __e.Get(1)
				_ = V518
				__e.TailApply(ShenFunc(symunspecialise), V518)

				return
			}, 1)
			gen16878 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen16877)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16878)

			gen16879 := MakeNative(func(__e Evaluator) {
				V519 := __e.Get(1)
				_ = V519
				__e.TailApply(ShenFunc(symuntrack), V519)

				return
			}, 1)
			gen16880 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen16879)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16880)

			gen16881 := MakeNative(func(__e Evaluator) {
				V520 := __e.Get(1)
				_ = V520
				__e.Return(MakeNative(func(__e Evaluator) {
					V521 := __e.Get(1)
					_ = V521
					__e.TailApply(ShenFunc(symunion), V520, V521)

					return
				}, 1))
				return
			}, 1)
			gen16882 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen16881)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16882)

			gen16883 := MakeNative(func(__e Evaluator) {
				V522 := __e.Get(1)
				_ = V522
				__e.Return(MakeNative(func(__e Evaluator) {
					V523 := __e.Get(1)
					_ = V523
					__e.Return(MakeNative(func(__e Evaluator) {
						V524 := __e.Get(1)
						_ = V524
						__e.Return(MakeNative(func(__e Evaluator) {
							V525 := __e.Get(1)
							_ = V525
							__e.TailApply(ShenFunc(symunify), V522, V523, V524, V525)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16884 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen16883)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16884)

			gen16885 := MakeNative(func(__e Evaluator) {
				V526 := __e.Get(1)
				_ = V526
				__e.Return(MakeNative(func(__e Evaluator) {
					V527 := __e.Get(1)
					_ = V527
					__e.Return(MakeNative(func(__e Evaluator) {
						V528 := __e.Get(1)
						_ = V528
						__e.Return(MakeNative(func(__e Evaluator) {
							V529 := __e.Get(1)
							_ = V529
							__e.TailApply(ShenFunc(symunify_b), V526, V527, V528, V529)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16886 := Call(__e, ShenFunc(symcons), MakeSymbol("unify!"), gen16885)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16886)

			gen16887 := MakeNative(func(__e Evaluator) {
				V530 := __e.Get(1)
				_ = V530
				__e.Return(MakeNative(func(__e Evaluator) {
					V531 := __e.Get(1)
					_ = V531
					__e.Return(MakeNative(func(__e Evaluator) {
						V532 := __e.Get(1)
						_ = V532
						__e.TailApply(ShenFunc(symunput), V530, V531, V532)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16888 := Call(__e, ShenFunc(symcons), MakeSymbol("unput"), gen16887)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16888)

			gen16889 := MakeNative(func(__e Evaluator) {
				V533 := __e.Get(1)
				_ = V533
				__e.TailApply(ShenFunc(symunprofile), V533)

				return
			}, 1)
			gen16890 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen16889)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16890)

			gen16891 := MakeNative(func(__e Evaluator) {
				V534 := __e.Get(1)
				_ = V534
				__e.TailApply(ShenFunc(symundefmacro), V534)

				return
			}, 1)
			gen16892 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen16891)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16892)

			gen16893 := MakeNative(func(__e Evaluator) {
				V535 := __e.Get(1)
				_ = V535
				__e.Return(MakeNative(func(__e Evaluator) {
					V536 := __e.Get(1)
					_ = V536
					__e.Return(MakeNative(func(__e Evaluator) {
						V537 := __e.Get(1)
						_ = V537
						__e.TailApply(ShenFunc(symreturn), V535, V536, V537)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16894 := Call(__e, ShenFunc(symcons), MakeSymbol("return"), gen16893)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16894)

			gen16895 := MakeNative(func(__e Evaluator) {
				V538 := __e.Get(1)
				_ = V538
				__e.Return(MakeNative(func(__e Evaluator) {
					V539 := __e.Get(1)
					_ = V539
					__e.TailApply(ShenFunc(symtype), V538, V539)

					return
				}, 1))
				return
			}, 1)
			gen16896 := Call(__e, ShenFunc(symcons), MakeSymbol("type"), gen16895)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16896)

			gen16897 := MakeNative(func(__e Evaluator) {
				V540 := __e.Get(1)
				_ = V540
				__e.TailApply(ShenFunc(symtuple_2), V540)

				return
			}, 1)
			gen16898 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen16897)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16898)

			gen16900 := MakeNative(func(__e Evaluator) {
				V541 := __e.Get(1)
				_ = V541
				__e.Return(MakeNative(func(__e Evaluator) {
					V542 := __e.Get(1)
					_ = V542
					gen16899 := MakeNative(func(__e Evaluator) {
						__e.Return(V541)
						return
					}, 0)
					__e.Return(Try(__e, gen16899).Catch(V542))
					return

				}, 1))
				return
			}, 1)
			gen16901 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen16900)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16901)

			gen16902 := MakeNative(func(__e Evaluator) {
				V543 := __e.Get(1)
				_ = V543
				__e.TailApply(ShenFunc(symtrack), V543)

				return
			}, 1)
			gen16903 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen16902)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16903)

			gen16904 := MakeNative(func(__e Evaluator) {
				V544 := __e.Get(1)
				_ = V544
				__e.TailApply(ShenFunc(symthaw), V544)

				return
			}, 1)
			gen16905 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen16904)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16905)

			gen16906 := MakeNative(func(__e Evaluator) {
				V545 := __e.Get(1)
				_ = V545
				__e.TailApply(ShenFunc(symtc), V545)

				return
			}, 1)
			gen16907 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen16906)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16907)

			gen16908 := MakeNative(func(__e Evaluator) {
				V546 := __e.Get(1)
				_ = V546
				__e.TailApply(ShenFunc(symtl), V546)

				return
			}, 1)
			gen16909 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen16908)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16909)

			gen16910 := MakeNative(func(__e Evaluator) {
				V547 := __e.Get(1)
				_ = V547
				__e.TailApply(ShenFunc(symtlstr), V547)

				return
			}, 1)
			gen16911 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen16910)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16911)

			gen16912 := MakeNative(func(__e Evaluator) {
				V548 := __e.Get(1)
				_ = V548
				__e.TailApply(ShenFunc(symtail), V548)

				return
			}, 1)
			gen16913 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen16912)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16913)

			gen16914 := MakeNative(func(__e Evaluator) {
				V549 := __e.Get(1)
				_ = V549
				__e.TailApply(ShenFunc(symsystemf), V549)

				return
			}, 1)
			gen16915 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen16914)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16915)

			gen16916 := MakeNative(func(__e Evaluator) {
				V550 := __e.Get(1)
				_ = V550
				__e.TailApply(ShenFunc(symsymbol_2), V550)

				return
			}, 1)
			gen16917 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen16916)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16917)

			gen16918 := MakeNative(func(__e Evaluator) {
				V551 := __e.Get(1)
				_ = V551
				__e.TailApply(ShenFunc(symstring_1_6symbol), V551)

				return
			}, 1)
			gen16919 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen16918)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16919)

			gen16920 := MakeNative(func(__e Evaluator) {
				V552 := __e.Get(1)
				_ = V552
				__e.TailApply(ShenFunc(symsum), V552)

				return
			}, 1)
			gen16921 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen16920)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16921)

			gen16922 := MakeNative(func(__e Evaluator) {
				V553 := __e.Get(1)
				_ = V553
				__e.Return(MakeNative(func(__e Evaluator) {
					V554 := __e.Get(1)
					_ = V554
					__e.Return(MakeNative(func(__e Evaluator) {
						V555 := __e.Get(1)
						_ = V555
						__e.TailApply(ShenFunc(symsubst), V553, V554, V555)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16923 := Call(__e, ShenFunc(symcons), MakeSymbol("subst"), gen16922)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16923)

			gen16924 := MakeNative(func(__e Evaluator) {
				V556 := __e.Get(1)
				_ = V556
				__e.TailApply(ShenFunc(symstring_2), V556)

				return
			}, 1)
			gen16925 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen16924)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16925)

			gen16926 := MakeNative(func(__e Evaluator) {
				V557 := __e.Get(1)
				_ = V557
				__e.TailApply(ShenFunc(symstring_1_6n), V557)

				return
			}, 1)
			gen16927 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen16926)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16927)

			gen16928 := MakeNative(func(__e Evaluator) {
				V558 := __e.Get(1)
				_ = V558
				__e.TailApply(ShenFunc(symstep), V558)

				return
			}, 1)
			gen16929 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen16928)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16929)

			gen16930 := MakeNative(func(__e Evaluator) {
				V559 := __e.Get(1)
				_ = V559
				__e.TailApply(ShenFunc(symspy), V559)

				return
			}, 1)
			gen16931 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen16930)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16931)

			gen16932 := MakeNative(func(__e Evaluator) {
				V560 := __e.Get(1)
				_ = V560
				__e.TailApply(ShenFunc(symspecialise), V560)

				return
			}, 1)
			gen16933 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen16932)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16933)

			gen16934 := MakeNative(func(__e Evaluator) {
				V561 := __e.Get(1)
				_ = V561
				__e.TailApply(ShenFunc(symsnd), V561)

				return
			}, 1)
			gen16935 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen16934)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16935)

			gen16936 := MakeNative(func(__e Evaluator) {
				V562 := __e.Get(1)
				_ = V562
				__e.TailApply(ShenFunc(symsimple_1error), V562)

				return
			}, 1)
			gen16937 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen16936)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16937)

			gen16938 := MakeNative(func(__e Evaluator) {
				V563 := __e.Get(1)
				_ = V563
				__e.Return(MakeNative(func(__e Evaluator) {
					V564 := __e.Get(1)
					_ = V564
					__e.TailApply(ShenFunc(symset), V563, V564)

					return
				}, 1))
				return
			}, 1)
			gen16939 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen16938)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16939)

			gen16940 := MakeNative(func(__e Evaluator) {
				V565 := __e.Get(1)
				_ = V565
				__e.TailApply(ShenFunc(symstr), V565)

				return
			}, 1)
			gen16941 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen16940)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16941)

			gen16942 := MakeNative(func(__e Evaluator) {
				V566 := __e.Get(1)
				_ = V566
				__e.TailApply(ShenFunc(symreverse), V566)

				return
			}, 1)
			gen16943 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen16942)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16943)

			gen16944 := MakeNative(func(__e Evaluator) {
				V567 := __e.Get(1)
				_ = V567
				__e.Return(MakeNative(func(__e Evaluator) {
					V568 := __e.Get(1)
					_ = V568
					__e.TailApply(ShenFunc(symremove), V567, V568)

					return
				}, 1))
				return
			}, 1)
			gen16945 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen16944)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16945)

			gen16946 := MakeNative(func(__e Evaluator) {
				V569 := __e.Get(1)
				_ = V569
				__e.TailApply(ShenFunc(symread), V569)

				return
			}, 1)
			gen16947 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen16946)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16947)

			gen16948 := MakeNative(func(__e Evaluator) {
				V570 := __e.Get(1)
				_ = V570
				__e.TailApply(ShenFunc(symread_1file), V570)

				return
			}, 1)
			gen16949 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen16948)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16949)

			gen16950 := MakeNative(func(__e Evaluator) {
				V571 := __e.Get(1)
				_ = V571
				__e.TailApply(ShenFunc(symread_1file_1as_1bytelist), V571)

				return
			}, 1)
			gen16951 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen16950)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16951)

			gen16952 := MakeNative(func(__e Evaluator) {
				V572 := __e.Get(1)
				_ = V572
				__e.TailApply(ShenFunc(symread_1file_1as_1string), V572)

				return
			}, 1)
			gen16953 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen16952)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16953)

			gen16954 := MakeNative(func(__e Evaluator) {
				V573 := __e.Get(1)
				_ = V573
				__e.TailApply(ShenFunc(symread_1byte), V573)

				return
			}, 1)
			gen16955 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen16954)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16955)

			gen16956 := MakeNative(func(__e Evaluator) {
				V574 := __e.Get(1)
				_ = V574
				__e.TailApply(ShenFunc(symread_1from_1string), V574)

				return
			}, 1)
			gen16957 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen16956)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16957)

			gen16958 := MakeNative(func(__e Evaluator) {
				V575 := __e.Get(1)
				_ = V575
				__e.TailApply(ShenFunc(sympackage_2), V575)

				return
			}, 1)
			gen16959 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen16958)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16959)

			gen16960 := MakeNative(func(__e Evaluator) {
				V576 := __e.Get(1)
				_ = V576
				__e.Return(MakeNative(func(__e Evaluator) {
					V577 := __e.Get(1)
					_ = V577
					__e.Return(MakeNative(func(__e Evaluator) {
						V578 := __e.Get(1)
						_ = V578
						__e.Return(MakeNative(func(__e Evaluator) {
							V579 := __e.Get(1)
							_ = V579
							__e.TailApply(ShenFunc(symput), V576, V577, V578, V579)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen16961 := Call(__e, ShenFunc(symcons), MakeSymbol("put"), gen16960)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16961)

			gen16962 := MakeNative(func(__e Evaluator) {
				V580 := __e.Get(1)
				_ = V580
				__e.TailApply(ShenFunc(sympreclude), V580)

				return
			}, 1)
			gen16963 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen16962)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16963)

			gen16964 := MakeNative(func(__e Evaluator) {
				V581 := __e.Get(1)
				_ = V581
				__e.TailApply(ShenFunc(sympreclude_1all_1but), V581)

				return
			}, 1)
			gen16965 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen16964)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16965)

			gen16966 := MakeNative(func(__e Evaluator) {
				V582 := __e.Get(1)
				_ = V582
				__e.TailApply(ShenFunc(symps), V582)

				return
			}, 1)
			gen16967 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen16966)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16967)

			gen16968 := MakeNative(func(__e Evaluator) {
				V583 := __e.Get(1)
				_ = V583
				__e.TailApply(ShenFunc(symprotect), V583)

				return
			}, 1)
			gen16969 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen16968)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16969)

			gen16970 := MakeNative(func(__e Evaluator) {
				V584 := __e.Get(1)
				_ = V584
				__e.TailApply(ShenFunc(symprofile_1results), V584)

				return
			}, 1)
			gen16971 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen16970)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16971)

			gen16972 := MakeNative(func(__e Evaluator) {
				V585 := __e.Get(1)
				_ = V585
				__e.TailApply(ShenFunc(symprofile), V585)

				return
			}, 1)
			gen16973 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen16972)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16973)

			gen16974 := MakeNative(func(__e Evaluator) {
				V586 := __e.Get(1)
				_ = V586
				__e.TailApply(ShenFunc(symprint), V586)

				return
			}, 1)
			gen16975 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen16974)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16975)

			gen16976 := MakeNative(func(__e Evaluator) {
				V587 := __e.Get(1)
				_ = V587
				__e.Return(MakeNative(func(__e Evaluator) {
					V588 := __e.Get(1)
					_ = V588
					__e.TailApply(ShenFunc(sympr), V587, V588)

					return
				}, 1))
				return
			}, 1)
			gen16977 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen16976)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16977)

			gen16978 := MakeNative(func(__e Evaluator) {
				V589 := __e.Get(1)
				_ = V589
				__e.Return(MakeNative(func(__e Evaluator) {
					V590 := __e.Get(1)
					_ = V590
					__e.TailApply(ShenFunc(sympos), V589, V590)

					return
				}, 1))
				return
			}, 1)
			gen16979 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen16978)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16979)

			gen16980 := MakeNative(func(__e Evaluator) {
				V591 := __e.Get(1)
				_ = V591
				__e.Return(MakeNative(func(__e Evaluator) {
					V592 := __e.Get(1)
					_ = V592
					if True == V591 {
						__e.Return(True)
						return
					} else {
						if True == V592 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}
					}
				}, 1))
				return
			}, 1)
			gen16981 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen16980)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16981)

			gen16982 := MakeNative(func(__e Evaluator) {
				V593 := __e.Get(1)
				_ = V593
				__e.TailApply(ShenFunc(symoptimise), V593)

				return
			}, 1)
			gen16983 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen16982)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16983)

			gen16984 := MakeNative(func(__e Evaluator) {
				V594 := __e.Get(1)
				_ = V594
				__e.Return(MakeNative(func(__e Evaluator) {
					V595 := __e.Get(1)
					_ = V595
					__e.TailApply(ShenFunc(symopen), V594, V595)

					return
				}, 1))
				return
			}, 1)
			gen16985 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), gen16984)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16985)

			gen16986 := MakeNative(func(__e Evaluator) {
				V596 := __e.Get(1)
				_ = V596
				__e.Return(MakeNative(func(__e Evaluator) {
					V597 := __e.Get(1)
					_ = V597
					__e.TailApply(ShenFunc(symoccurrences), V596, V597)

					return
				}, 1))
				return
			}, 1)
			gen16987 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen16986)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16987)

			gen16988 := MakeNative(func(__e Evaluator) {
				V598 := __e.Get(1)
				_ = V598
				__e.TailApply(ShenFunc(symoccurs_1check), V598)

				return
			}, 1)
			gen16989 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen16988)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16989)

			gen16990 := MakeNative(func(__e Evaluator) {
				V599 := __e.Get(1)
				_ = V599
				__e.TailApply(ShenFunc(symn_1_6string), V599)

				return
			}, 1)
			gen16991 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen16990)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16991)

			gen16992 := MakeNative(func(__e Evaluator) {
				V600 := __e.Get(1)
				_ = V600
				__e.TailApply(ShenFunc(symnumber_2), V600)

				return
			}, 1)
			gen16993 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen16992)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16993)

			gen16994 := MakeNative(func(__e Evaluator) {
				V601 := __e.Get(1)
				_ = V601
				__e.Return(MakeNative(func(__e Evaluator) {
					V602 := __e.Get(1)
					_ = V602
					__e.TailApply(ShenFunc(symnth), V601, V602)

					return
				}, 1))
				return
			}, 1)
			gen16995 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen16994)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16995)

			gen16996 := MakeNative(func(__e Evaluator) {
				V603 := __e.Get(1)
				_ = V603
				__e.TailApply(ShenFunc(symnot), V603)

				return
			}, 1)
			gen16997 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen16996)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16997)

			gen16998 := MakeNative(func(__e Evaluator) {
				V604 := __e.Get(1)
				_ = V604
				__e.TailApply(ShenFunc(symnl), V604)

				return
			}, 1)
			gen16999 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen16998)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen16999)

			gen17000 := MakeNative(func(__e Evaluator) {
				V605 := __e.Get(1)
				_ = V605
				__e.TailApply(ShenFunc(symmacroexpand), V605)

				return
			}, 1)
			gen17001 := Call(__e, ShenFunc(symcons), MakeSymbol("macroexpand"), gen17000)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17001)

			gen17002 := MakeNative(func(__e Evaluator) {
				V606 := __e.Get(1)
				_ = V606
				__e.TailApply(ShenFunc(symmaxinferences), V606)

				return
			}, 1)
			gen17003 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen17002)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17003)

			gen17004 := MakeNative(func(__e Evaluator) {
				V607 := __e.Get(1)
				_ = V607
				__e.Return(MakeNative(func(__e Evaluator) {
					V608 := __e.Get(1)
					_ = V608
					__e.TailApply(ShenFunc(symmapcan), V607, V608)

					return
				}, 1))
				return
			}, 1)
			gen17005 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen17004)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17005)

			gen17006 := MakeNative(func(__e Evaluator) {
				V609 := __e.Get(1)
				_ = V609
				__e.Return(MakeNative(func(__e Evaluator) {
					V610 := __e.Get(1)
					_ = V610
					__e.TailApply(ShenFunc(symmap), V609, V610)

					return
				}, 1))
				return
			}, 1)
			gen17007 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen17006)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17007)

			gen17008 := MakeNative(func(__e Evaluator) {
				V611 := __e.Get(1)
				_ = V611
				__e.TailApply(ShenFunc(symload), V611)

				return
			}, 1)
			gen17009 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen17008)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17009)

			gen17010 := MakeNative(func(__e Evaluator) {
				V612 := __e.Get(1)
				_ = V612
				__e.TailApply(ShenFunc(symlineread), V612)

				return
			}, 1)
			gen17011 := Call(__e, ShenFunc(symcons), MakeSymbol("lineread"), gen17010)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17011)

			gen17012 := MakeNative(func(__e Evaluator) {
				V613 := __e.Get(1)
				_ = V613
				__e.TailApply(ShenFunc(symlimit), V613)

				return
			}, 1)
			gen17013 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen17012)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17013)

			gen17014 := MakeNative(func(__e Evaluator) {
				V614 := __e.Get(1)
				_ = V614
				__e.TailApply(ShenFunc(symlength), V614)

				return
			}, 1)
			gen17015 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen17014)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17015)

			gen17016 := MakeNative(func(__e Evaluator) {
				V615 := __e.Get(1)
				_ = V615
				__e.Return(MakeNative(func(__e Evaluator) {
					V616 := __e.Get(1)
					_ = V616
					__e.TailApply(ShenFunc(symintersection), V615, V616)

					return
				}, 1))
				return
			}, 1)
			gen17017 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen17016)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17017)

			gen17018 := MakeNative(func(__e Evaluator) {
				V617 := __e.Get(1)
				_ = V617
				__e.TailApply(ShenFunc(symintern), V617)

				return
			}, 1)
			gen17019 := Call(__e, ShenFunc(symcons), MakeSymbol("intern"), gen17018)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17019)

			gen17020 := MakeNative(func(__e Evaluator) {
				V618 := __e.Get(1)
				_ = V618
				__e.TailApply(ShenFunc(syminteger_2), V618)

				return
			}, 1)
			gen17021 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen17020)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17021)

			gen17022 := MakeNative(func(__e Evaluator) {
				V619 := __e.Get(1)
				_ = V619
				__e.TailApply(ShenFunc(syminput), V619)

				return
			}, 1)
			gen17023 := Call(__e, ShenFunc(symcons), MakeSymbol("input"), gen17022)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17023)

			gen17024 := MakeNative(func(__e Evaluator) {
				V620 := __e.Get(1)
				_ = V620
				__e.Return(MakeNative(func(__e Evaluator) {
					V621 := __e.Get(1)
					_ = V621
					__e.TailApply(ShenFunc(syminput_7), V620, V621)

					return
				}, 1))
				return
			}, 1)
			gen17025 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen17024)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17025)

			gen17026 := MakeNative(func(__e Evaluator) {
				V622 := __e.Get(1)
				_ = V622
				__e.TailApply(ShenFunc(syminclude), V622)

				return
			}, 1)
			gen17027 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen17026)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17027)

			gen17028 := MakeNative(func(__e Evaluator) {
				V623 := __e.Get(1)
				_ = V623
				__e.TailApply(ShenFunc(syminclude_1all_1but), V623)

				return
			}, 1)
			gen17029 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen17028)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17029)

			gen17030 := MakeNative(func(__e Evaluator) {
				V624 := __e.Get(1)
				_ = V624
				__e.TailApply(ShenFunc(syminternal), V624)

				return
			}, 1)
			gen17031 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen17030)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17031)

			gen17032 := MakeNative(func(__e Evaluator) {
				V625 := __e.Get(1)
				_ = V625
				__e.Return(MakeNative(func(__e Evaluator) {
					V626 := __e.Get(1)
					_ = V626
					__e.Return(MakeNative(func(__e Evaluator) {
						V627 := __e.Get(1)
						_ = V627
						if True == V625 {
							__e.Return(V626)
							return
						} else {
							__e.Return(V627)
							return
						}
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen17033 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen17032)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17033)

			gen17034 := MakeNative(func(__e Evaluator) {
				V628 := __e.Get(1)
				_ = V628
				__e.Return(MakeNative(func(__e Evaluator) {
					V629 := __e.Get(1)
					_ = V629
					__e.Return(MakeNative(func(__e Evaluator) {
						V630 := __e.Get(1)
						_ = V630
						__e.Return(MakeNative(func(__e Evaluator) {
							V631 := __e.Get(1)
							_ = V631
							__e.TailApply(ShenFunc(symidentical), V628, V629, V630, V631)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen17035 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen17034)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17035)

			gen17036 := MakeNative(func(__e Evaluator) {
				V632 := __e.Get(1)
				_ = V632
				__e.TailApply(ShenFunc(symhead), V632)

				return
			}, 1)
			gen17037 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen17036)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17037)

			gen17038 := MakeNative(func(__e Evaluator) {
				V633 := __e.Get(1)
				_ = V633
				__e.TailApply(ShenFunc(symhd), V633)

				return
			}, 1)
			gen17039 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen17038)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17039)

			gen17040 := MakeNative(func(__e Evaluator) {
				V634 := __e.Get(1)
				_ = V634
				__e.TailApply(ShenFunc(symhdv), V634)

				return
			}, 1)
			gen17041 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen17040)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17041)

			gen17042 := MakeNative(func(__e Evaluator) {
				V635 := __e.Get(1)
				_ = V635
				__e.TailApply(ShenFunc(symhdstr), V635)

				return
			}, 1)
			gen17043 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen17042)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17043)

			gen17044 := MakeNative(func(__e Evaluator) {
				V636 := __e.Get(1)
				_ = V636
				__e.Return(MakeNative(func(__e Evaluator) {
					V637 := __e.Get(1)
					_ = V637
					__e.TailApply(ShenFunc(symhash), V636, V637)

					return
				}, 1))
				return
			}, 1)
			gen17045 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen17044)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17045)

			gen17046 := MakeNative(func(__e Evaluator) {
				V638 := __e.Get(1)
				_ = V638
				__e.Return(MakeNative(func(__e Evaluator) {
					V639 := __e.Get(1)
					_ = V639
					__e.Return(MakeNative(func(__e Evaluator) {
						V640 := __e.Get(1)
						_ = V640
						__e.TailApply(ShenFunc(symget), V638, V639, V640)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen17047 := Call(__e, ShenFunc(symcons), MakeSymbol("get"), gen17046)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17047)

			gen17048 := MakeNative(func(__e Evaluator) {
				V641 := __e.Get(1)
				_ = V641
				__e.TailApply(ShenFunc(symget_1time), V641)

				return
			}, 1)
			gen17049 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen17048)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17049)

			gen17050 := MakeNative(func(__e Evaluator) {
				V642 := __e.Get(1)
				_ = V642
				__e.TailApply(ShenFunc(symgensym), V642)

				return
			}, 1)
			gen17051 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen17050)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17051)

			gen17052 := MakeNative(func(__e Evaluator) {
				V643 := __e.Get(1)
				_ = V643
				__e.TailApply(ShenFunc(symfst), V643)

				return
			}, 1)
			gen17053 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen17052)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17053)

			gen17054 := MakeNative(func(__e Evaluator) {
				V644 := __e.Get(1)
				_ = V644
				__e.Return(MakeNative(func(__e Evaluator) {
					__e.Return(V644)
					return
				}, 0))
				return
			}, 1)
			gen17055 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen17054)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17055)

			gen17056 := MakeNative(func(__e Evaluator) {
				V645 := __e.Get(1)
				_ = V645
				__e.Return(MakeNative(func(__e Evaluator) {
					V646 := __e.Get(1)
					_ = V646
					__e.TailApply(ShenFunc(symfix), V645, V646)

					return
				}, 1))
				return
			}, 1)
			gen17057 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen17056)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17057)

			gen17058 := MakeNative(func(__e Evaluator) {
				V647 := __e.Get(1)
				_ = V647
				__e.Return(MakeNative(func(__e Evaluator) {
					V648 := __e.Get(1)
					_ = V648
					__e.TailApply(ShenFunc(symfail_1if), V647, V648)

					return
				}, 1))
				return
			}, 1)
			gen17059 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen17058)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17059)

			gen17060 := MakeNative(func(__e Evaluator) {
				V649 := __e.Get(1)
				_ = V649
				__e.Return(MakeNative(func(__e Evaluator) {
					V650 := __e.Get(1)
					_ = V650
					__e.Return(MakeNative(func(__e Evaluator) {
						V651 := __e.Get(1)
						_ = V651
						__e.Return(MakeNative(func(__e Evaluator) {
							V652 := __e.Get(1)
							_ = V652
							__e.Return(MakeNative(func(__e Evaluator) {
								V653 := __e.Get(1)
								_ = V653
								__e.TailApply(ShenFunc(symfindall), V649, V650, V651, V652, V653)

								return
							}, 1))
							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen17061 := Call(__e, ShenFunc(symcons), MakeSymbol("findall"), gen17060)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17061)

			gen17062 := MakeNative(func(__e Evaluator) {
				V654 := __e.Get(1)
				_ = V654
				__e.TailApply(ShenFunc(symenable_1type_1theory), V654)

				return
			}, 1)
			gen17063 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen17062)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17063)

			gen17064 := MakeNative(func(__e Evaluator) {
				V655 := __e.Get(1)
				_ = V655
				__e.TailApply(ShenFunc(symexplode), V655)

				return
			}, 1)
			gen17065 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen17064)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17065)

			gen17066 := MakeNative(func(__e Evaluator) {
				V656 := __e.Get(1)
				_ = V656
				__e.TailApply(ShenFunc(symexternal), V656)

				return
			}, 1)
			gen17067 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen17066)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17067)

			gen17068 := MakeNative(func(__e Evaluator) {
				V657 := __e.Get(1)
				_ = V657
				__e.TailApply(ShenFunc(symeval_1kl), V657)

				return
			}, 1)
			gen17069 := Call(__e, ShenFunc(symcons), MakeSymbol("eval-kl"), gen17068)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17069)

			gen17070 := MakeNative(func(__e Evaluator) {
				V658 := __e.Get(1)
				_ = V658
				__e.TailApply(ShenFunc(symeval), V658)

				return
			}, 1)
			gen17071 := Call(__e, ShenFunc(symcons), MakeSymbol("eval"), gen17070)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17071)

			gen17072 := MakeNative(func(__e Evaluator) {
				V659 := __e.Get(1)
				_ = V659
				__e.TailApply(ShenFunc(symerror_1to_1string), V659)

				return
			}, 1)
			gen17073 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen17072)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17073)

			gen17074 := MakeNative(func(__e Evaluator) {
				V660 := __e.Get(1)
				_ = V660
				__e.TailApply(ShenFunc(symempty_2), V660)

				return
			}, 1)
			gen17075 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen17074)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17075)

			gen17076 := MakeNative(func(__e Evaluator) {
				V661 := __e.Get(1)
				_ = V661
				__e.Return(MakeNative(func(__e Evaluator) {
					V662 := __e.Get(1)
					_ = V662
					__e.TailApply(ShenFunc(symelement_2), V661, V662)

					return
				}, 1))
				return
			}, 1)
			gen17077 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen17076)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17077)

			gen17078 := MakeNative(func(__e Evaluator) {
				V663 := __e.Get(1)
				_ = V663
				__e.Return(MakeNative(func(__e Evaluator) {
					V664 := __e.Get(1)
					_ = V664
					// V663
					__e.Return(V664)
					return

				}, 1))
				return
			}, 1)
			gen17079 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen17078)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17079)

			gen17080 := MakeNative(func(__e Evaluator) {
				V665 := __e.Get(1)
				_ = V665
				__e.Return(MakeNative(func(__e Evaluator) {
					V666 := __e.Get(1)
					_ = V666
					__e.TailApply(ShenFunc(symdifference), V665, V666)

					return
				}, 1))
				return
			}, 1)
			gen17081 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen17080)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17081)

			gen17082 := MakeNative(func(__e Evaluator) {
				V667 := __e.Get(1)
				_ = V667
				__e.TailApply(ShenFunc(symdestroy), V667)

				return
			}, 1)
			gen17083 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen17082)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17083)

			gen17084 := MakeNative(func(__e Evaluator) {
				V668 := __e.Get(1)
				_ = V668
				__e.Return(MakeNative(func(__e Evaluator) {
					V669 := __e.Get(1)
					_ = V669
					__e.TailApply(ShenFunc(symdeclare), V668, V669)

					return
				}, 1))
				return
			}, 1)
			gen17085 := Call(__e, ShenFunc(symcons), MakeSymbol("declare"), gen17084)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17085)

			gen17086 := MakeNative(func(__e Evaluator) {
				V670 := __e.Get(1)
				_ = V670
				__e.Return(MakeNative(func(__e Evaluator) {
					V671 := __e.Get(1)
					_ = V671
					__e.TailApply(ShenFunc(symcn), V670, V671)

					return
				}, 1))
				return
			}, 1)
			gen17087 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen17086)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17087)

			gen17088 := MakeNative(func(__e Evaluator) {
				V672 := __e.Get(1)
				_ = V672
				__e.TailApply(ShenFunc(symcons_2), V672)

				return
			}, 1)
			gen17089 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen17088)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17089)

			gen17090 := MakeNative(func(__e Evaluator) {
				V673 := __e.Get(1)
				_ = V673
				__e.Return(MakeNative(func(__e Evaluator) {
					V674 := __e.Get(1)
					_ = V674
					__e.TailApply(ShenFunc(symcons), V673, V674)

					return
				}, 1))
				return
			}, 1)
			gen17091 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen17090)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17091)

			gen17092 := MakeNative(func(__e Evaluator) {
				V675 := __e.Get(1)
				_ = V675
				__e.Return(MakeNative(func(__e Evaluator) {
					V676 := __e.Get(1)
					_ = V676
					__e.TailApply(ShenFunc(symconcat), V675, V676)

					return
				}, 1))
				return
			}, 1)
			gen17093 := Call(__e, ShenFunc(symcons), MakeSymbol("concat"), gen17092)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17093)

			gen17094 := MakeNative(func(__e Evaluator) {
				V677 := __e.Get(1)
				_ = V677
				__e.Return(MakeNative(func(__e Evaluator) {
					V678 := __e.Get(1)
					_ = V678
					__e.Return(MakeNative(func(__e Evaluator) {
						V679 := __e.Get(1)
						_ = V679
						__e.TailApply(ShenFunc(symcompile), V677, V678, V679)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen17095 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen17094)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17095)

			gen17096 := MakeNative(func(__e Evaluator) {
				V680 := __e.Get(1)
				_ = V680
				__e.TailApply(ShenFunc(symcd), V680)

				return
			}, 1)
			gen17097 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen17096)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17097)

			gen17098 := MakeNative(func(__e Evaluator) {
				V681 := __e.Get(1)
				_ = V681
				__e.TailApply(ShenFunc(symclose), V681)

				return
			}, 1)
			gen17099 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen17098)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17099)

			gen17100 := MakeNative(func(__e Evaluator) {
				V682 := __e.Get(1)
				_ = V682
				__e.TailApply(ShenFunc(symbound_2), V682)

				return
			}, 1)
			gen17101 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen17100)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17101)

			gen17102 := MakeNative(func(__e Evaluator) {
				V683 := __e.Get(1)
				_ = V683
				__e.TailApply(ShenFunc(symboolean_2), V683)

				return
			}, 1)
			gen17103 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen17102)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17103)

			gen17104 := MakeNative(func(__e Evaluator) {
				V684 := __e.Get(1)
				_ = V684
				__e.Return(MakeNative(func(__e Evaluator) {
					V685 := __e.Get(1)
					_ = V685
					__e.TailApply(ShenFunc(symassoc), V684, V685)

					return
				}, 1))
				return
			}, 1)
			gen17105 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen17104)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17105)

			gen17106 := MakeNative(func(__e Evaluator) {
				V686 := __e.Get(1)
				_ = V686
				__e.TailApply(ShenFunc(symarity), V686)

				return
			}, 1)
			gen17107 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen17106)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17107)

			gen17108 := MakeNative(func(__e Evaluator) {
				V687 := __e.Get(1)
				_ = V687
				__e.Return(MakeNative(func(__e Evaluator) {
					V688 := __e.Get(1)
					_ = V688
					__e.TailApply(ShenFunc(symappend), V687, V688)

					return
				}, 1))
				return
			}, 1)
			gen17109 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen17108)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17109)

			gen17110 := MakeNative(func(__e Evaluator) {
				V689 := __e.Get(1)
				_ = V689
				__e.Return(MakeNative(func(__e Evaluator) {
					V690 := __e.Get(1)
					_ = V690
					if True == V689 {
						if True == V690 {
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
				}, 1))
				return
			}, 1)
			gen17111 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen17110)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17111)

			gen17112 := MakeNative(func(__e Evaluator) {
				V691 := __e.Get(1)
				_ = V691
				__e.Return(MakeNative(func(__e Evaluator) {
					V692 := __e.Get(1)
					_ = V692
					__e.TailApply(ShenFunc(symadjoin), V691, V692)

					return
				}, 1))
				return
			}, 1)
			gen17113 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen17112)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17113)

			gen17114 := MakeNative(func(__e Evaluator) {
				V693 := __e.Get(1)
				_ = V693
				__e.Return(MakeNative(func(__e Evaluator) {
					V694 := __e.Get(1)
					_ = V694
					__e.TailApply(ShenFunc(sym_5_1address), V693, V694)

					return
				}, 1))
				return
			}, 1)
			gen17115 := Call(__e, ShenFunc(symcons), MakeSymbol("<-address"), gen17114)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17115)

			gen17116 := MakeNative(func(__e Evaluator) {
				V695 := __e.Get(1)
				_ = V695
				__e.Return(MakeNative(func(__e Evaluator) {
					V696 := __e.Get(1)
					_ = V696
					__e.Return(MakeNative(func(__e Evaluator) {
						V697 := __e.Get(1)
						_ = V697
						__e.TailApply(ShenFunc(symaddress_1_6), V695, V696, V697)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen17117 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen17116)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17117)

			gen17118 := MakeNative(func(__e Evaluator) {
				V698 := __e.Get(1)
				_ = V698
				__e.TailApply(ShenFunc(symabsvector_2), V698)

				return
			}, 1)
			gen17119 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen17118)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen17119)

			gen17120 := MakeNative(func(__e Evaluator) {
				V699 := __e.Get(1)
				_ = V699
				__e.TailApply(ShenFunc(symabsvector), V699)

				return
			}, 1)
			gen17121 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector"), gen17120)

			__e.TailApply(ShenFunc(symshen_4set_1lambda_1form_1entry), gen17121)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-lambda-forms"), gen17122)

		gen17123 := MakeNative(func(__e Evaluator) {
			Call(__e, ShenFunc(symshen_4initialise_1environment))
			Call(__e, ShenFunc(symshen_4initialise_1lambda_1forms))
			Call(__e, ShenFunc(symshen_4initialise_1signedfunc_1lambda_1forms))
			__e.TailApply(ShenFunc(symshen_4initialise_1signedfuncs))

			return

		}, 0)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.initialise"), gen17123)

		return

	}, 0))
}
