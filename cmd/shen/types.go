package main

import . "github.com/tiancaiamao/shen-go/kl"

var TypesMain = MakeNative(func(__e *ControlFlow) {
	tmp16705 := MakeNative(func(__e *ControlFlow) {
		V5883 := __e.Get(1)
		_ = V5883
		V5884 := __e.Get(2)
		_ = V5884
		tmp16706 := MakeNative(func(__e *ControlFlow) {
			W5885 := __e.Get(1)
			_ = W5885
			tmp16707 := MakeNative(func(__e *ControlFlow) {
				W5886 := __e.Get(1)
				_ = W5886
				tmp16708 := MakeNative(func(__e *ControlFlow) {
					W5891 := __e.Get(1)
					_ = W5891
					tmp16709 := MakeNative(func(__e *ControlFlow) {
						W5892 := __e.Get(1)
						_ = W5892
						__e.Return(V5883)
						return
					}, 1)

					tmp16710 := PrimValue(symshen_4_dsigf_d)

					tmp16711 := Call(__e, PrimFunc(symshen_4assoc_1_6), V5883, W5891, tmp16710)

					tmp16712 := PrimSet(symshen_4_dsigf_d, tmp16711)

					__e.TailApply(tmp16709, tmp16712)
					return

				}, 1)

				tmp16713 := Call(__e, PrimFunc(symshen_4prolog_1abstraction), V5884)

				tmp16714 := Call(__e, PrimFunc(symeval_1kl), tmp16713)

				__e.TailApply(tmp16708, tmp16714)
				return

			}, 1)

			tmp16715 := MakeNative(func(__e *ControlFlow) {
				Z5887 := __e.Get(1)
				_ = Z5887
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Z5888 := __e.Get(1)
					_ = Z5888
					__e.Return(MakeNative(func(__e *ControlFlow) {
						Z5889 := __e.Get(1)
						_ = Z5889
						__e.Return(MakeNative(func(__e *ControlFlow) {
							Z5890 := __e.Get(1)
							_ = Z5890
							tmp16716 := Call(__e, PrimFunc(symshen_4incinfs))

							_ = tmp16716

							tmp16717 := Call(__e, PrimFunc(symshen_4deref), V5883, Z5887)

							tmp16718 := Call(__e, PrimFunc(symreceive), tmp16717)

							tmp16719 := Call(__e, PrimFunc(symshen_4deref), W5885, Z5887)

							tmp16720 := Call(__e, PrimFunc(symreceive), tmp16719)

							__e.TailApply(PrimFunc(symshen_4variancy), tmp16718, tmp16720, Z5887, Z5888, Z5889, Z5890)
							return

						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)

			tmp16721 := Call(__e, PrimFunc(symshen_4prolog_1vector))

			tmp16722 := Call(__e, tmp16715, tmp16721)

			tmp16723 := Call(__e, PrimFunc(symvector), MakeNumber(0))

			tmp16724 := Call(__e, PrimFunc(sym_8v), MakeNumber(0), tmp16723)

			tmp16725 := Call(__e, PrimFunc(sym_8v), True, tmp16724)

			tmp16726 := Call(__e, tmp16722, tmp16725)

			tmp16727 := Call(__e, tmp16726, MakeNumber(0))

			tmp16728 := MakeNative(func(__e *ControlFlow) {
				__e.Return(True)
				return
			}, 0)

			tmp16729 := Call(__e, tmp16727, tmp16728)

			__e.TailApply(tmp16707, tmp16729)
			return

		}, 1)

		tmp16730 := Call(__e, PrimFunc(symshen_4rectify_1type), V5884)

		__e.TailApply(tmp16706, tmp16730)
		return

	}, 2)

	tmp16731 := Call(__e, ns2_1set, symdeclare, tmp16705)

	_ = tmp16731

	tmp16732 := MakeNative(func(__e *ControlFlow) {
		V5893 := __e.Get(1)
		_ = V5893
		V5894 := __e.Get(2)
		_ = V5894
		V5895 := __e.Get(3)
		_ = V5895
		V5896 := __e.Get(4)
		_ = V5896
		V5897 := __e.Get(5)
		_ = V5897
		V5898 := __e.Get(6)
		_ = V5898
		tmp16739 := Call(__e, PrimFunc(symshen_4unlocked_2), V5896)

		if True == tmp16739 {
			tmp16733 := MakeNative(func(__e *ControlFlow) {
				W5899 := __e.Get(1)
				_ = W5899
				tmp16734 := Call(__e, PrimFunc(symshen_4incinfs))

				_ = tmp16734

				tmp16735 := MakeNative(func(__e *ControlFlow) {
					__e.TailApply(PrimFunc(symshen_4variants_2), V5893, W5899, V5894, V5895, V5896, V5897, V5898)
					return
				}, 0)

				tmp16736 := Call(__e, PrimFunc(symshen_4system_1S_1h), V5893, W5899, Nil, V5895, V5896, V5897, tmp16735)

				__e.TailApply(PrimFunc(symshen_4gc), V5895, tmp16736)
				return

			}, 1)

			tmp16737 := Call(__e, PrimFunc(symshen_4newpv), V5895)

			__e.TailApply(tmp16733, tmp16737)
			return

		} else {
			__e.Return(False)
			return
		}

	}, 6)

	tmp16740 := Call(__e, ns2_1set, symshen_4variancy, tmp16732)

	_ = tmp16740

	tmp16741 := MakeNative(func(__e *ControlFlow) {
		V5900 := __e.Get(1)
		_ = V5900
		V5901 := __e.Get(2)
		_ = V5901
		V5902 := __e.Get(3)
		_ = V5902
		V5903 := __e.Get(4)
		_ = V5903
		V5904 := __e.Get(5)
		_ = V5904
		V5905 := __e.Get(6)
		_ = V5905
		V5906 := __e.Get(7)
		_ = V5906
		tmp16742 := MakeNative(func(__e *ControlFlow) {
			W5907 := __e.Get(1)
			_ = W5907
			tmp16743 := MakeNative(func(__e *ControlFlow) {
				W5908 := __e.Get(1)
				_ = W5908
				tmp16767 := PrimEqual(W5908, False)

				if True == tmp16767 {
					tmp16744 := MakeNative(func(__e *ControlFlow) {
						W5911 := __e.Get(1)
						_ = W5911
						tmp16761 := PrimEqual(W5911, False)

						if True == tmp16761 {
							tmp16745 := MakeNative(func(__e *ControlFlow) {
								W5912 := __e.Get(1)
								_ = W5912
								tmp16747 := PrimEqual(W5912, False)

								if True == tmp16747 {
									__e.TailApply(PrimFunc(symshen_4unlock), V5904, W5907)
									return
								} else {
									__e.Return(W5912)
									return
								}

							}, 1)

							tmp16759 := Call(__e, PrimFunc(symshen_4unlocked_2), V5904)

							var ifres16748 Obj

							if True == tmp16759 {
								tmp16749 := MakeNative(func(__e *ControlFlow) {
									W5913 := __e.Get(1)
									_ = W5913
									tmp16750 := Call(__e, PrimFunc(symshen_4incinfs))

									_ = tmp16750

									tmp16751 := Call(__e, PrimFunc(symshen_4deref), V5900, V5903)

									tmp16752 := Call(__e, PrimFunc(symshen_4app), tmp16751, MakeString(" may create errors\n"), symshen_4a)

									tmp16753 := PrimStringConcat(MakeString("warning: changing the type of "), tmp16752)

									tmp16754 := Call(__e, PrimFunc(symstoutput))

									tmp16755 := Call(__e, PrimFunc(sympr), tmp16753, tmp16754)

									tmp16756 := Call(__e, PrimFunc(symis), W5913, tmp16755, V5903, V5904, W5907, V5906)

									__e.TailApply(PrimFunc(symshen_4gc), V5903, tmp16756)
									return

								}, 1)

								tmp16757 := Call(__e, PrimFunc(symshen_4newpv), V5903)

								tmp16758 := Call(__e, tmp16749, tmp16757)

								ifres16748 = tmp16758

							} else {
								ifres16748 = False

							}

							__e.TailApply(tmp16745, ifres16748)
							return

						} else {
							__e.Return(W5911)
							return
						}

					}, 1)

					tmp16765 := Call(__e, PrimFunc(symshen_4unlocked_2), V5904)

					var ifres16762 Obj

					if True == tmp16765 {
						tmp16763 := Call(__e, PrimFunc(symshen_4incinfs))

						_ = tmp16763

						tmp16764 := Call(__e, PrimFunc(symis_b), V5901, V5902, V5903, V5904, W5907, V5906)

						ifres16762 = tmp16764

					} else {
						ifres16762 = False

					}

					__e.TailApply(tmp16744, ifres16762)
					return

				} else {
					__e.Return(W5908)
					return
				}

			}, 1)

			tmp16779 := Call(__e, PrimFunc(symshen_4unlocked_2), V5904)

			var ifres16768 Obj

			if True == tmp16779 {
				tmp16769 := MakeNative(func(__e *ControlFlow) {
					W5909 := __e.Get(1)
					_ = W5909
					tmp16770 := MakeNative(func(__e *ControlFlow) {
						W5910 := __e.Get(1)
						_ = W5910
						tmp16774 := PrimEqual(W5909, symsymbol)

						if True == tmp16774 {
							__e.TailApply(PrimFunc(symthaw), W5910)
							return
						} else {
							tmp16772 := Call(__e, PrimFunc(symshen_4pvar_2), W5909)

							if True == tmp16772 {
								__e.TailApply(PrimFunc(symshen_4bind_b), W5909, symsymbol, V5903, W5910)
								return
							} else {
								__e.Return(False)
								return
							}

						}

					}, 1)

					tmp16775 := MakeNative(func(__e *ControlFlow) {
						tmp16776 := Call(__e, PrimFunc(symshen_4incinfs))

						_ = tmp16776

						__e.TailApply(PrimFunc(symshen_4cut), V5903, V5904, W5907, V5906)
						return

					}, 0)

					__e.TailApply(tmp16770, tmp16775)
					return

				}, 1)

				tmp16777 := Call(__e, PrimFunc(symshen_4lazyderef), V5901, V5903)

				tmp16778 := Call(__e, tmp16769, tmp16777)

				ifres16768 = tmp16778

			} else {
				ifres16768 = False

			}

			__e.TailApply(tmp16743, ifres16768)
			return

		}, 1)

		tmp16780 := PrimNumberAdd(V5905, MakeNumber(1))

		__e.TailApply(tmp16742, tmp16780)
		return

	}, 7)

	tmp16781 := Call(__e, ns2_1set, symshen_4variants_2, tmp16741)

	_ = tmp16781

	tmp16782 := MakeNative(func(__e *ControlFlow) {
		V5914 := __e.Get(1)
		_ = V5914
		tmp16783 := MakeNative(func(__e *ControlFlow) {
			W5915 := __e.Get(1)
			_ = W5915
			tmp16784 := MakeNative(func(__e *ControlFlow) {
				W5916 := __e.Get(1)
				_ = W5916
				tmp16785 := MakeNative(func(__e *ControlFlow) {
					W5917 := __e.Get(1)
					_ = W5917
					tmp16786 := MakeNative(func(__e *ControlFlow) {
						W5918 := __e.Get(1)
						_ = W5918
						tmp16787 := MakeNative(func(__e *ControlFlow) {
							W5919 := __e.Get(1)
							_ = W5919
							tmp16788 := MakeNative(func(__e *ControlFlow) {
								W5920 := __e.Get(1)
								_ = W5920
								tmp16789 := Call(__e, PrimFunc(symshen_4rcons__form), V5914)

								tmp16790 := PrimCons(W5918, Nil)

								tmp16791 := PrimCons(W5917, tmp16790)

								tmp16792 := PrimCons(W5916, tmp16791)

								tmp16793 := PrimCons(W5915, tmp16792)

								tmp16794 := PrimCons(tmp16789, tmp16793)

								tmp16795 := PrimCons(W5919, tmp16794)

								tmp16796 := PrimCons(symis_b, tmp16795)

								tmp16797 := Call(__e, PrimFunc(symshen_4stpart), W5920, tmp16796, W5915)

								tmp16798 := PrimCons(tmp16797, Nil)

								tmp16799 := PrimCons(W5918, tmp16798)

								tmp16800 := PrimCons(symlambda, tmp16799)

								tmp16801 := PrimCons(tmp16800, Nil)

								tmp16802 := PrimCons(W5917, tmp16801)

								tmp16803 := PrimCons(symlambda, tmp16802)

								tmp16804 := PrimCons(tmp16803, Nil)

								tmp16805 := PrimCons(W5916, tmp16804)

								tmp16806 := PrimCons(symlambda, tmp16805)

								tmp16807 := PrimCons(tmp16806, Nil)

								tmp16808 := PrimCons(W5915, tmp16807)

								tmp16809 := PrimCons(symlambda, tmp16808)

								tmp16810 := PrimCons(tmp16809, Nil)

								tmp16811 := PrimCons(W5919, tmp16810)

								__e.Return(PrimCons(symlambda, tmp16811))
								return

							}, 1)

							tmp16812 := Call(__e, PrimFunc(symshen_4extract_1vars), V5914)

							__e.TailApply(tmp16788, tmp16812)
							return

						}, 1)

						tmp16813 := Call(__e, PrimFunc(symgensym), symV)

						__e.TailApply(tmp16787, tmp16813)
						return

					}, 1)

					tmp16814 := Call(__e, PrimFunc(symgensym), symC)

					__e.TailApply(tmp16786, tmp16814)
					return

				}, 1)

				tmp16815 := Call(__e, PrimFunc(symgensym), symKey)

				__e.TailApply(tmp16785, tmp16815)
				return

			}, 1)

			tmp16816 := Call(__e, PrimFunc(symgensym), symL)

			__e.TailApply(tmp16784, tmp16816)
			return

		}, 1)

		tmp16817 := Call(__e, PrimFunc(symgensym), symB)

		__e.TailApply(tmp16783, tmp16817)
		return

	}, 1)

	tmp16818 := Call(__e, ns2_1set, symshen_4prolog_1abstraction, tmp16782)

	_ = tmp16818

	tmp16819 := MakeNative(func(__e *ControlFlow) {
		V5921 := __e.Get(1)
		_ = V5921
		tmp16820 := MakeNative(func(__e *ControlFlow) {
			W5922 := __e.Get(1)
			_ = W5922
			__e.TailApply(W5922, V5921)
			return
		}, 1)

		tmp16821 := PrimValue(symshen_4_ddemodulation_1function_d)

		__e.TailApply(tmp16820, tmp16821)
		return

	}, 1)

	__e.TailApply(ns2_1set, symshen_4demod, tmp16819)
	return

}, 0)
