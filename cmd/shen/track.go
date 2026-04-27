package main

import . "github.com/tiancaiamao/shen-go/kl"

var TrackMain = MakeNative(func(__e *ControlFlow) {
	tmp11784 := MakeNative(func(__e *ControlFlow) {
		V5766 := __e.Get(1)
		_ = V5766
		tmp11785 := Call(__e, PrimFunc(symshen_4app), V5766, MakeString(";\n"), symshen_4a)

		tmp11786 := PrimStringConcat(MakeString("partial function "), tmp11785)

		tmp11787 := Call(__e, PrimFunc(symstoutput))

		tmp11788 := Call(__e, PrimFunc(sympr), tmp11786, tmp11787)

		_ = tmp11788

		tmp11797 := Call(__e, PrimFunc(symshen_4tracked_2), V5766)

		tmp11798 := PrimNot(tmp11797)

		var ifres11792 Obj

		if True == tmp11798 {
			tmp11794 := Call(__e, PrimFunc(symshen_4app), V5766, MakeString("? "), symshen_4a)

			tmp11795 := PrimStringConcat(MakeString("track "), tmp11794)

			tmp11796 := Call(__e, PrimFunc(symy_1or_1n_2), tmp11795)

			var ifres11793 Obj

			if True == tmp11796 {
				ifres11793 = True

			} else {
				ifres11793 = False

			}

			ifres11792 = ifres11793

		} else {
			ifres11792 = False

		}

		var ifres11789 Obj

		if True == ifres11792 {
			tmp11790 := Call(__e, PrimFunc(symps), V5766)

			tmp11791 := Call(__e, PrimFunc(symshen_4track_1function), tmp11790)

			ifres11789 = tmp11791

		} else {
			ifres11789 = symshen_4ok

		}

		_ = ifres11789

		__e.Return(PrimSimpleError(MakeString("aborted")))
		return

	}, 1)

	tmp11799 := Call(__e, ns2_1set, symshen_4f_1error, tmp11784)

	_ = tmp11799

	tmp11800 := MakeNative(func(__e *ControlFlow) {
		V5767 := __e.Get(1)
		_ = V5767
		tmp11801 := PrimValue(symshen_4_dtracking_d)

		__e.TailApply(PrimFunc(symelement_2), V5767, tmp11801)
		return

	}, 1)

	tmp11802 := Call(__e, ns2_1set, symshen_4tracked_2, tmp11800)

	_ = tmp11802

	tmp11803 := MakeNative(func(__e *ControlFlow) {
		V5768 := __e.Get(1)
		_ = V5768
		tmp11804 := MakeNative(func(__e *ControlFlow) {
			W5769 := __e.Get(1)
			_ = W5769
			__e.TailApply(PrimFunc(symshen_4track_1function), W5769)
			return
		}, 1)

		tmp11805 := Call(__e, PrimFunc(symps), V5768)

		__e.TailApply(tmp11804, tmp11805)
		return

	}, 1)

	tmp11806 := Call(__e, ns2_1set, symtrack, tmp11803)

	_ = tmp11806

	tmp11807 := MakeNative(func(__e *ControlFlow) {
		V5772 := __e.Get(1)
		_ = V5772
		tmp11864 := PrimIsPair(V5772)

		var ifres11838 Obj

		if True == tmp11864 {
			tmp11862 := PrimHead(V5772)

			tmp11863 := PrimEqual(symdefun, tmp11862)

			var ifres11840 Obj

			if True == tmp11863 {
				tmp11860 := PrimTail(V5772)

				tmp11861 := PrimIsPair(tmp11860)

				var ifres11842 Obj

				if True == tmp11861 {
					tmp11857 := PrimTail(V5772)

					tmp11858 := PrimTail(tmp11857)

					tmp11859 := PrimIsPair(tmp11858)

					var ifres11844 Obj

					if True == tmp11859 {
						tmp11853 := PrimTail(V5772)

						tmp11854 := PrimTail(tmp11853)

						tmp11855 := PrimTail(tmp11854)

						tmp11856 := PrimIsPair(tmp11855)

						var ifres11846 Obj

						if True == tmp11856 {
							tmp11848 := PrimTail(V5772)

							tmp11849 := PrimTail(tmp11848)

							tmp11850 := PrimTail(tmp11849)

							tmp11851 := PrimTail(tmp11850)

							tmp11852 := PrimEqual(Nil, tmp11851)

							var ifres11847 Obj

							if True == tmp11852 {
								ifres11847 = True

							} else {
								ifres11847 = False

							}

							ifres11846 = ifres11847

						} else {
							ifres11846 = False

						}

						var ifres11845 Obj

						if True == ifres11846 {
							ifres11845 = True

						} else {
							ifres11845 = False

						}

						ifres11844 = ifres11845

					} else {
						ifres11844 = False

					}

					var ifres11843 Obj

					if True == ifres11844 {
						ifres11843 = True

					} else {
						ifres11843 = False

					}

					ifres11842 = ifres11843

				} else {
					ifres11842 = False

				}

				var ifres11841 Obj

				if True == ifres11842 {
					ifres11841 = True

				} else {
					ifres11841 = False

				}

				ifres11840 = ifres11841

			} else {
				ifres11840 = False

			}

			var ifres11839 Obj

			if True == ifres11840 {
				ifres11839 = True

			} else {
				ifres11839 = False

			}

			ifres11838 = ifres11839

		} else {
			ifres11838 = False

		}

		if True == ifres11838 {
			tmp11808 := MakeNative(func(__e *ControlFlow) {
				W5773 := __e.Get(1)
				_ = W5773
				tmp11809 := MakeNative(func(__e *ControlFlow) {
					W5774 := __e.Get(1)
					_ = W5774
					tmp11810 := MakeNative(func(__e *ControlFlow) {
						W5775 := __e.Get(1)
						_ = W5775
						tmp11811 := PrimTail(V5772)

						__e.Return(PrimHead(tmp11811))
						return

					}, 1)

					tmp11812 := PrimTail(V5772)

					tmp11813 := PrimHead(tmp11812)

					tmp11814 := PrimValue(symshen_4_dtracking_d)

					tmp11815 := Call(__e, PrimFunc(symadjoin), tmp11813, tmp11814)

					tmp11816 := PrimSet(symshen_4_dtracking_d, tmp11815)

					__e.TailApply(tmp11810, tmp11816)
					return

				}, 1)

				tmp11817 := Call(__e, PrimFunc(symeval_1kl), W5773)

				__e.TailApply(tmp11809, tmp11817)
				return

			}, 1)

			tmp11818 := PrimTail(V5772)

			tmp11819 := PrimHead(tmp11818)

			tmp11820 := PrimTail(V5772)

			tmp11821 := PrimTail(tmp11820)

			tmp11822 := PrimHead(tmp11821)

			tmp11823 := PrimTail(V5772)

			tmp11824 := PrimHead(tmp11823)

			tmp11825 := PrimTail(V5772)

			tmp11826 := PrimTail(tmp11825)

			tmp11827 := PrimHead(tmp11826)

			tmp11828 := PrimTail(V5772)

			tmp11829 := PrimTail(tmp11828)

			tmp11830 := PrimTail(tmp11829)

			tmp11831 := PrimHead(tmp11830)

			tmp11832 := Call(__e, PrimFunc(symshen_4insert_1tracking_1code), tmp11824, tmp11827, tmp11831)

			tmp11833 := PrimCons(tmp11832, Nil)

			tmp11834 := PrimCons(tmp11822, tmp11833)

			tmp11835 := PrimCons(tmp11819, tmp11834)

			tmp11836 := PrimCons(symdefun, tmp11835)

			__e.TailApply(tmp11808, tmp11836)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.track-function")))
			return
		}

	}, 1)

	tmp11865 := Call(__e, ns2_1set, symshen_4track_1function, tmp11807)

	_ = tmp11865

	tmp11866 := MakeNative(func(__e *ControlFlow) {
		V5776 := __e.Get(1)
		_ = V5776
		V5777 := __e.Get(2)
		_ = V5777
		V5778 := __e.Get(3)
		_ = V5778
		tmp11867 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11868 := PrimCons(symvalue, tmp11867)

		tmp11869 := PrimCons(MakeNumber(1), Nil)

		tmp11870 := PrimCons(tmp11868, tmp11869)

		tmp11871 := PrimCons(sym_7, tmp11870)

		tmp11872 := PrimCons(tmp11871, Nil)

		tmp11873 := PrimCons(symshen_4_dcall_d, tmp11872)

		tmp11874 := PrimCons(symset, tmp11873)

		tmp11875 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11876 := PrimCons(symvalue, tmp11875)

		tmp11877 := Call(__e, PrimFunc(symshen_4prolog_1track), V5778, V5777)

		tmp11878 := Call(__e, PrimFunc(symshen_4cons_1form), tmp11877)

		tmp11879 := PrimCons(tmp11878, Nil)

		tmp11880 := PrimCons(V5776, tmp11879)

		tmp11881 := PrimCons(tmp11876, tmp11880)

		tmp11882 := PrimCons(symshen_4input_1track, tmp11881)

		tmp11883 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

		tmp11884 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11885 := PrimCons(symvalue, tmp11884)

		tmp11886 := PrimCons(symResult, Nil)

		tmp11887 := PrimCons(V5776, tmp11886)

		tmp11888 := PrimCons(tmp11885, tmp11887)

		tmp11889 := PrimCons(symshen_4output_1track, tmp11888)

		tmp11890 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11891 := PrimCons(symvalue, tmp11890)

		tmp11892 := PrimCons(MakeNumber(1), Nil)

		tmp11893 := PrimCons(tmp11891, tmp11892)

		tmp11894 := PrimCons(sym_1, tmp11893)

		tmp11895 := PrimCons(tmp11894, Nil)

		tmp11896 := PrimCons(symshen_4_dcall_d, tmp11895)

		tmp11897 := PrimCons(symset, tmp11896)

		tmp11898 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

		tmp11899 := PrimCons(symResult, Nil)

		tmp11900 := PrimCons(tmp11898, tmp11899)

		tmp11901 := PrimCons(symdo, tmp11900)

		tmp11902 := PrimCons(tmp11901, Nil)

		tmp11903 := PrimCons(tmp11897, tmp11902)

		tmp11904 := PrimCons(symdo, tmp11903)

		tmp11905 := PrimCons(tmp11904, Nil)

		tmp11906 := PrimCons(tmp11889, tmp11905)

		tmp11907 := PrimCons(symdo, tmp11906)

		tmp11908 := PrimCons(tmp11907, Nil)

		tmp11909 := PrimCons(V5778, tmp11908)

		tmp11910 := PrimCons(symResult, tmp11909)

		tmp11911 := PrimCons(symlet, tmp11910)

		tmp11912 := PrimCons(tmp11911, Nil)

		tmp11913 := PrimCons(tmp11883, tmp11912)

		tmp11914 := PrimCons(symdo, tmp11913)

		tmp11915 := PrimCons(tmp11914, Nil)

		tmp11916 := PrimCons(tmp11882, tmp11915)

		tmp11917 := PrimCons(symdo, tmp11916)

		tmp11918 := PrimCons(tmp11917, Nil)

		tmp11919 := PrimCons(tmp11874, tmp11918)

		__e.Return(PrimCons(symdo, tmp11919))
		return

	}, 3)

	tmp11920 := Call(__e, ns2_1set, symshen_4insert_1tracking_1code, tmp11866)

	_ = tmp11920

	tmp11921 := MakeNative(func(__e *ControlFlow) {
		V5779 := __e.Get(1)
		_ = V5779
		V5780 := __e.Get(2)
		_ = V5780
		tmp11924 := Call(__e, PrimFunc(symoccurrences), symshen_4incinfs, V5779)

		tmp11925 := PrimEqual(tmp11924, MakeNumber(0))

		if True == tmp11925 {
			__e.Return(V5780)
			return
		} else {
			tmp11922 := Call(__e, PrimFunc(symshen_4vector_1parameter), V5780)

			__e.TailApply(PrimFunc(symshen_4vector_1dereference), V5780, tmp11922)
			return

		}

	}, 2)

	tmp11926 := Call(__e, ns2_1set, symshen_4prolog_1track, tmp11921)

	_ = tmp11926

	tmp11927 := MakeNative(func(__e *ControlFlow) {
		V5783 := __e.Get(1)
		_ = V5783
		tmp11956 := PrimEqual(Nil, V5783)

		if True == tmp11956 {
			__e.Return(Nil)
			return
		} else {
			tmp11954 := PrimIsPair(V5783)

			var ifres11932 Obj

			if True == tmp11954 {
				tmp11952 := PrimTail(V5783)

				tmp11953 := PrimIsPair(tmp11952)

				var ifres11934 Obj

				if True == tmp11953 {
					tmp11949 := PrimTail(V5783)

					tmp11950 := PrimTail(tmp11949)

					tmp11951 := PrimIsPair(tmp11950)

					var ifres11936 Obj

					if True == tmp11951 {
						tmp11945 := PrimTail(V5783)

						tmp11946 := PrimTail(tmp11945)

						tmp11947 := PrimTail(tmp11946)

						tmp11948 := PrimIsPair(tmp11947)

						var ifres11938 Obj

						if True == tmp11948 {
							tmp11940 := PrimTail(V5783)

							tmp11941 := PrimTail(tmp11940)

							tmp11942 := PrimTail(tmp11941)

							tmp11943 := PrimTail(tmp11942)

							tmp11944 := PrimEqual(Nil, tmp11943)

							var ifres11939 Obj

							if True == tmp11944 {
								ifres11939 = True

							} else {
								ifres11939 = False

							}

							ifres11938 = ifres11939

						} else {
							ifres11938 = False

						}

						var ifres11937 Obj

						if True == ifres11938 {
							ifres11937 = True

						} else {
							ifres11937 = False

						}

						ifres11936 = ifres11937

					} else {
						ifres11936 = False

					}

					var ifres11935 Obj

					if True == ifres11936 {
						ifres11935 = True

					} else {
						ifres11935 = False

					}

					ifres11934 = ifres11935

				} else {
					ifres11934 = False

				}

				var ifres11933 Obj

				if True == ifres11934 {
					ifres11933 = True

				} else {
					ifres11933 = False

				}

				ifres11932 = ifres11933

			} else {
				ifres11932 = False

			}

			if True == ifres11932 {
				__e.Return(PrimHead(V5783))
				return
			} else {
				tmp11930 := PrimIsPair(V5783)

				if True == tmp11930 {
					tmp11928 := PrimTail(V5783)

					__e.TailApply(PrimFunc(symshen_4vector_1parameter), tmp11928)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1parameter)
					return
				}

			}

		}

	}, 1)

	tmp11957 := Call(__e, ns2_1set, symshen_4vector_1parameter, tmp11927)

	_ = tmp11957

	tmp11958 := MakeNative(func(__e *ControlFlow) {
		V5786 := __e.Get(1)
		_ = V5786
		V5787 := __e.Get(2)
		_ = V5787
		tmp11992 := PrimEqual(Nil, V5787)

		if True == tmp11992 {
			__e.Return(V5786)
			return
		} else {
			tmp11990 := PrimIsPair(V5786)

			var ifres11968 Obj

			if True == tmp11990 {
				tmp11988 := PrimTail(V5786)

				tmp11989 := PrimIsPair(tmp11988)

				var ifres11970 Obj

				if True == tmp11989 {
					tmp11985 := PrimTail(V5786)

					tmp11986 := PrimTail(tmp11985)

					tmp11987 := PrimIsPair(tmp11986)

					var ifres11972 Obj

					if True == tmp11987 {
						tmp11981 := PrimTail(V5786)

						tmp11982 := PrimTail(tmp11981)

						tmp11983 := PrimTail(tmp11982)

						tmp11984 := PrimIsPair(tmp11983)

						var ifres11974 Obj

						if True == tmp11984 {
							tmp11976 := PrimTail(V5786)

							tmp11977 := PrimTail(tmp11976)

							tmp11978 := PrimTail(tmp11977)

							tmp11979 := PrimTail(tmp11978)

							tmp11980 := PrimEqual(Nil, tmp11979)

							var ifres11975 Obj

							if True == tmp11980 {
								ifres11975 = True

							} else {
								ifres11975 = False

							}

							ifres11974 = ifres11975

						} else {
							ifres11974 = False

						}

						var ifres11973 Obj

						if True == ifres11974 {
							ifres11973 = True

						} else {
							ifres11973 = False

						}

						ifres11972 = ifres11973

					} else {
						ifres11972 = False

					}

					var ifres11971 Obj

					if True == ifres11972 {
						ifres11971 = True

					} else {
						ifres11971 = False

					}

					ifres11970 = ifres11971

				} else {
					ifres11970 = False

				}

				var ifres11969 Obj

				if True == ifres11970 {
					ifres11969 = True

				} else {
					ifres11969 = False

				}

				ifres11968 = ifres11969

			} else {
				ifres11968 = False

			}

			if True == ifres11968 {
				__e.Return(V5786)
				return
			} else {
				tmp11966 := PrimIsPair(V5786)

				if True == tmp11966 {
					tmp11959 := PrimHead(V5786)

					tmp11960 := PrimCons(V5787, Nil)

					tmp11961 := PrimCons(tmp11959, tmp11960)

					tmp11962 := PrimCons(symshen_4deref, tmp11961)

					tmp11963 := PrimTail(V5786)

					tmp11964 := Call(__e, PrimFunc(symshen_4vector_1dereference), tmp11963, V5787)

					__e.Return(PrimCons(tmp11962, tmp11964))
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4vector_1dereference)
					return
				}

			}

		}

	}, 2)

	tmp11993 := Call(__e, ns2_1set, symshen_4vector_1dereference, tmp11958)

	_ = tmp11993

	tmp11994 := MakeNative(func(__e *ControlFlow) {
		V5790 := __e.Get(1)
		_ = V5790
		tmp11998 := PrimEqual(sym_7, V5790)

		if True == tmp11998 {
			__e.Return(PrimSet(symshen_4_dstep_d, True))
			return
		} else {
			tmp11996 := PrimEqual(sym_1, V5790)

			if True == tmp11996 {
				__e.Return(PrimSet(symshen_4_dstep_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp11999 := Call(__e, ns2_1set, symstep, tmp11994)

	_ = tmp11999

	tmp12000 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dstep_d))
		return
	}, 0)

	tmp12001 := Call(__e, ns2_1set, symshen_4step_2, tmp12000)

	_ = tmp12001

	tmp12002 := MakeNative(func(__e *ControlFlow) {
		V5793 := __e.Get(1)
		_ = V5793
		tmp12006 := PrimEqual(sym_7, V5793)

		if True == tmp12006 {
			__e.Return(PrimSet(symshen_4_dspy_d, True))
			return
		} else {
			tmp12004 := PrimEqual(sym_1, V5793)

			if True == tmp12004 {
				__e.Return(PrimSet(symshen_4_dspy_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp12007 := Call(__e, ns2_1set, symspy, tmp12002)

	_ = tmp12007

	tmp12008 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dspy_d))
		return
	}, 0)

	tmp12009 := Call(__e, ns2_1set, symshen_4spy_2, tmp12008)

	_ = tmp12009

	tmp12010 := MakeNative(func(__e *ControlFlow) {
		tmp12014 := PrimValue(symshen_4_dstep_d)

		if True == tmp12014 {
			tmp12011 := PrimValue(sym_dstinput_d)

			tmp12012 := PrimReadByte(tmp12011)

			__e.TailApply(PrimFunc(symshen_4check_1byte), tmp12012)
			return

		} else {
			__e.TailApply(PrimFunc(symnl), MakeNumber(1))
			return
		}

	}, 0)

	tmp12015 := Call(__e, ns2_1set, symshen_4terpri_1or_1read_1char, tmp12010)

	_ = tmp12015

	tmp12016 := MakeNative(func(__e *ControlFlow) {
		V5796 := __e.Get(1)
		_ = V5796
		tmp12018 := PrimEqual(MakeNumber(94), V5796)

		if True == tmp12018 {
			__e.Return(PrimSimpleError(MakeString("aborted")))
			return
		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp12019 := Call(__e, ns2_1set, symshen_4check_1byte, tmp12016)

	_ = tmp12019

	tmp12020 := MakeNative(func(__e *ControlFlow) {
		V5797 := __e.Get(1)
		_ = V5797
		V5798 := __e.Get(2)
		_ = V5798
		V5799 := __e.Get(3)
		_ = V5799
		tmp12021 := Call(__e, PrimFunc(symshen_4spaces), V5797)

		tmp12022 := Call(__e, PrimFunc(symshen_4spaces), V5797)

		tmp12023 := Call(__e, PrimFunc(symshen_4app), tmp12022, MakeString(""), symshen_4a)

		tmp12024 := PrimStringConcat(MakeString(" \n"), tmp12023)

		tmp12025 := Call(__e, PrimFunc(symshen_4app), V5798, tmp12024, symshen_4a)

		tmp12026 := PrimStringConcat(MakeString("> Inputs to "), tmp12025)

		tmp12027 := Call(__e, PrimFunc(symshen_4app), V5797, tmp12026, symshen_4a)

		tmp12028 := PrimStringConcat(MakeString("<"), tmp12027)

		tmp12029 := Call(__e, PrimFunc(symshen_4app), tmp12021, tmp12028, symshen_4a)

		tmp12030 := PrimStringConcat(MakeString("\n"), tmp12029)

		tmp12031 := Call(__e, PrimFunc(symstoutput))

		tmp12032 := Call(__e, PrimFunc(sympr), tmp12030, tmp12031)

		_ = tmp12032

		__e.TailApply(PrimFunc(symshen_4recursively_1print), V5799)
		return

	}, 3)

	tmp12033 := Call(__e, ns2_1set, symshen_4input_1track, tmp12020)

	_ = tmp12033

	tmp12034 := MakeNative(func(__e *ControlFlow) {
		V5802 := __e.Get(1)
		_ = V5802
		tmp12044 := PrimEqual(Nil, V5802)

		if True == tmp12044 {
			tmp12035 := Call(__e, PrimFunc(symstoutput))

			__e.TailApply(PrimFunc(sympr), MakeString(" ==>"), tmp12035)
			return

		} else {
			tmp12042 := PrimIsPair(V5802)

			if True == tmp12042 {
				tmp12036 := PrimHead(V5802)

				tmp12037 := Call(__e, PrimFunc(symprint), tmp12036)

				_ = tmp12037

				tmp12038 := Call(__e, PrimFunc(symstoutput))

				tmp12039 := Call(__e, PrimFunc(sympr), MakeString(", "), tmp12038)

				_ = tmp12039

				tmp12040 := PrimTail(V5802)

				__e.TailApply(PrimFunc(symshen_4recursively_1print), tmp12040)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursively-print")))
				return
			}

		}

	}, 1)

	tmp12045 := Call(__e, ns2_1set, symshen_4recursively_1print, tmp12034)

	_ = tmp12045

	tmp12046 := MakeNative(func(__e *ControlFlow) {
		V5803 := __e.Get(1)
		_ = V5803
		tmp12050 := PrimEqual(MakeNumber(0), V5803)

		if True == tmp12050 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp12047 := PrimNumberSubtract(V5803, MakeNumber(1))

			tmp12048 := Call(__e, PrimFunc(symshen_4spaces), tmp12047)

			__e.Return(PrimStringConcat(MakeString(" "), tmp12048))
			return

		}

	}, 1)

	tmp12051 := Call(__e, ns2_1set, symshen_4spaces, tmp12046)

	_ = tmp12051

	tmp12052 := MakeNative(func(__e *ControlFlow) {
		V5804 := __e.Get(1)
		_ = V5804
		V5805 := __e.Get(2)
		_ = V5805
		V5806 := __e.Get(3)
		_ = V5806
		tmp12053 := Call(__e, PrimFunc(symshen_4spaces), V5804)

		tmp12054 := Call(__e, PrimFunc(symshen_4spaces), V5804)

		tmp12055 := Call(__e, PrimFunc(symshen_4app), V5806, MakeString(""), symshen_4s)

		tmp12056 := PrimStringConcat(MakeString("==> "), tmp12055)

		tmp12057 := Call(__e, PrimFunc(symshen_4app), tmp12054, tmp12056, symshen_4a)

		tmp12058 := PrimStringConcat(MakeString(" \n"), tmp12057)

		tmp12059 := Call(__e, PrimFunc(symshen_4app), V5805, tmp12058, symshen_4a)

		tmp12060 := PrimStringConcat(MakeString("> Output from "), tmp12059)

		tmp12061 := Call(__e, PrimFunc(symshen_4app), V5804, tmp12060, symshen_4a)

		tmp12062 := PrimStringConcat(MakeString("<"), tmp12061)

		tmp12063 := Call(__e, PrimFunc(symshen_4app), tmp12053, tmp12062, symshen_4a)

		tmp12064 := PrimStringConcat(MakeString("\n"), tmp12063)

		tmp12065 := Call(__e, PrimFunc(symstoutput))

		__e.TailApply(PrimFunc(sympr), tmp12064, tmp12065)
		return

	}, 3)

	tmp12066 := Call(__e, ns2_1set, symshen_4output_1track, tmp12052)

	_ = tmp12066

	tmp12067 := MakeNative(func(__e *ControlFlow) {
		V5807 := __e.Get(1)
		_ = V5807
		tmp12068 := PrimValue(symshen_4_dtracking_d)

		tmp12069 := Call(__e, PrimFunc(symremove), V5807, tmp12068)

		tmp12070 := PrimSet(symshen_4_dtracking_d, tmp12069)

		_ = tmp12070

		tmp12071 := MakeNative(func(__e *ControlFlow) {
			tmp12072 := Call(__e, PrimFunc(symps), V5807)

			__e.TailApply(PrimFunc(symeval), tmp12072)
			return

		}, 0)

		tmp12073 := MakeNative(func(__e *ControlFlow) {
			Z5808 := __e.Get(1)
			_ = Z5808
			__e.Return(V5807)
			return
		}, 1)

		tmp12074 := Call(__e, try_1catch, tmp12071, tmp12073)

		_ = tmp12074

		__e.Return(V5807)
		return

	}, 1)

	tmp12075 := Call(__e, ns2_1set, symuntrack, tmp12067)

	_ = tmp12075

	tmp12076 := MakeNative(func(__e *ControlFlow) {
		V5809 := __e.Get(1)
		_ = V5809
		V5810 := __e.Get(2)
		_ = V5810
		__e.TailApply(PrimFunc(symshen_4remove_1h), V5809, V5810, Nil)
		return
	}, 2)

	tmp12077 := Call(__e, ns2_1set, symremove, tmp12076)

	_ = tmp12077

	tmp12078 := MakeNative(func(__e *ControlFlow) {
		V5820 := __e.Get(1)
		_ = V5820
		V5821 := __e.Get(2)
		_ = V5821
		V5822 := __e.Get(3)
		_ = V5822
		tmp12093 := PrimEqual(Nil, V5821)

		if True == tmp12093 {
			__e.TailApply(PrimFunc(symreverse), V5822)
			return
		} else {
			tmp12091 := PrimIsPair(V5821)

			var ifres12087 Obj

			if True == tmp12091 {
				tmp12089 := PrimHead(V5821)

				tmp12090 := PrimEqual(V5820, tmp12089)

				var ifres12088 Obj

				if True == tmp12090 {
					ifres12088 = True

				} else {
					ifres12088 = False

				}

				ifres12087 = ifres12088

			} else {
				ifres12087 = False

			}

			if True == ifres12087 {
				tmp12079 := PrimHead(V5821)

				tmp12080 := PrimTail(V5821)

				__e.TailApply(PrimFunc(symshen_4remove_1h), tmp12079, tmp12080, V5822)
				return

			} else {
				tmp12085 := PrimIsPair(V5821)

				if True == tmp12085 {
					tmp12081 := PrimTail(V5821)

					tmp12082 := PrimHead(V5821)

					tmp12083 := PrimCons(tmp12082, V5822)

					__e.TailApply(PrimFunc(symshen_4remove_1h), V5820, tmp12081, tmp12083)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-h")))
					return
				}

			}

		}

	}, 3)

	tmp12094 := Call(__e, ns2_1set, symshen_4remove_1h, tmp12078)

	_ = tmp12094

	tmp12095 := MakeNative(func(__e *ControlFlow) {
		V5823 := __e.Get(1)
		_ = V5823
		tmp12096 := PrimValue(symshen_4_dprofiled_d)

		tmp12097 := PrimCons(V5823, tmp12096)

		tmp12098 := PrimSet(symshen_4_dprofiled_d, tmp12097)

		_ = tmp12098

		tmp12099 := Call(__e, PrimFunc(symps), V5823)

		__e.TailApply(PrimFunc(symshen_4profile_1help), tmp12099)
		return

	}, 1)

	tmp12100 := Call(__e, ns2_1set, symprofile, tmp12095)

	_ = tmp12100

	tmp12101 := MakeNative(func(__e *ControlFlow) {
		V5826 := __e.Get(1)
		_ = V5826
		tmp12171 := PrimIsPair(V5826)

		var ifres12145 Obj

		if True == tmp12171 {
			tmp12169 := PrimHead(V5826)

			tmp12170 := PrimEqual(symdefun, tmp12169)

			var ifres12147 Obj

			if True == tmp12170 {
				tmp12167 := PrimTail(V5826)

				tmp12168 := PrimIsPair(tmp12167)

				var ifres12149 Obj

				if True == tmp12168 {
					tmp12164 := PrimTail(V5826)

					tmp12165 := PrimTail(tmp12164)

					tmp12166 := PrimIsPair(tmp12165)

					var ifres12151 Obj

					if True == tmp12166 {
						tmp12160 := PrimTail(V5826)

						tmp12161 := PrimTail(tmp12160)

						tmp12162 := PrimTail(tmp12161)

						tmp12163 := PrimIsPair(tmp12162)

						var ifres12153 Obj

						if True == tmp12163 {
							tmp12155 := PrimTail(V5826)

							tmp12156 := PrimTail(tmp12155)

							tmp12157 := PrimTail(tmp12156)

							tmp12158 := PrimTail(tmp12157)

							tmp12159 := PrimEqual(Nil, tmp12158)

							var ifres12154 Obj

							if True == tmp12159 {
								ifres12154 = True

							} else {
								ifres12154 = False

							}

							ifres12153 = ifres12154

						} else {
							ifres12153 = False

						}

						var ifres12152 Obj

						if True == ifres12153 {
							ifres12152 = True

						} else {
							ifres12152 = False

						}

						ifres12151 = ifres12152

					} else {
						ifres12151 = False

					}

					var ifres12150 Obj

					if True == ifres12151 {
						ifres12150 = True

					} else {
						ifres12150 = False

					}

					ifres12149 = ifres12150

				} else {
					ifres12149 = False

				}

				var ifres12148 Obj

				if True == ifres12149 {
					ifres12148 = True

				} else {
					ifres12148 = False

				}

				ifres12147 = ifres12148

			} else {
				ifres12147 = False

			}

			var ifres12146 Obj

			if True == ifres12147 {
				ifres12146 = True

			} else {
				ifres12146 = False

			}

			ifres12145 = ifres12146

		} else {
			ifres12145 = False

		}

		if True == ifres12145 {
			tmp12102 := MakeNative(func(__e *ControlFlow) {
				W5827 := __e.Get(1)
				_ = W5827
				tmp12103 := MakeNative(func(__e *ControlFlow) {
					W5828 := __e.Get(1)
					_ = W5828
					tmp12104 := MakeNative(func(__e *ControlFlow) {
						W5829 := __e.Get(1)
						_ = W5829
						tmp12105 := MakeNative(func(__e *ControlFlow) {
							W5830 := __e.Get(1)
							_ = W5830
							tmp12106 := MakeNative(func(__e *ControlFlow) {
								W5831 := __e.Get(1)
								_ = W5831
								tmp12107 := PrimTail(V5826)

								__e.Return(PrimHead(tmp12107))
								return

							}, 1)

							tmp12108 := Call(__e, PrimFunc(symeval_1kl), W5829)

							__e.TailApply(tmp12106, tmp12108)
							return

						}, 1)

						tmp12109 := Call(__e, PrimFunc(symeval_1kl), W5828)

						__e.TailApply(tmp12105, tmp12109)
						return

					}, 1)

					tmp12110 := PrimTail(V5826)

					tmp12111 := PrimTail(tmp12110)

					tmp12112 := PrimHead(tmp12111)

					tmp12113 := PrimTail(V5826)

					tmp12114 := PrimHead(tmp12113)

					tmp12115 := PrimTail(V5826)

					tmp12116 := PrimTail(tmp12115)

					tmp12117 := PrimTail(tmp12116)

					tmp12118 := PrimHead(tmp12117)

					tmp12119 := Call(__e, PrimFunc(symsubst), W5827, tmp12114, tmp12118)

					tmp12120 := PrimCons(tmp12119, Nil)

					tmp12121 := PrimCons(tmp12112, tmp12120)

					tmp12122 := PrimCons(W5827, tmp12121)

					tmp12123 := PrimCons(symdefun, tmp12122)

					__e.TailApply(tmp12104, tmp12123)
					return

				}, 1)

				tmp12124 := PrimTail(V5826)

				tmp12125 := PrimHead(tmp12124)

				tmp12126 := PrimTail(V5826)

				tmp12127 := PrimTail(tmp12126)

				tmp12128 := PrimHead(tmp12127)

				tmp12129 := PrimTail(V5826)

				tmp12130 := PrimHead(tmp12129)

				tmp12131 := PrimTail(V5826)

				tmp12132 := PrimTail(tmp12131)

				tmp12133 := PrimHead(tmp12132)

				tmp12134 := PrimTail(V5826)

				tmp12135 := PrimTail(tmp12134)

				tmp12136 := PrimHead(tmp12135)

				tmp12137 := PrimCons(W5827, tmp12136)

				tmp12138 := Call(__e, PrimFunc(symshen_4profile_1func), tmp12130, tmp12133, tmp12137)

				tmp12139 := PrimCons(tmp12138, Nil)

				tmp12140 := PrimCons(tmp12128, tmp12139)

				tmp12141 := PrimCons(tmp12125, tmp12140)

				tmp12142 := PrimCons(symdefun, tmp12141)

				__e.TailApply(tmp12103, tmp12142)
				return

			}, 1)

			tmp12143 := Call(__e, PrimFunc(symgensym), symshen_4f)

			__e.TailApply(tmp12102, tmp12143)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
			return
		}

	}, 1)

	tmp12172 := Call(__e, ns2_1set, symshen_4profile_1help, tmp12101)

	_ = tmp12172

	tmp12173 := MakeNative(func(__e *ControlFlow) {
		V5832 := __e.Get(1)
		_ = V5832
		tmp12174 := PrimValue(symshen_4_dprofiled_d)

		tmp12175 := Call(__e, PrimFunc(symremove), V5832, tmp12174)

		tmp12176 := PrimSet(symshen_4_dprofiled_d, tmp12175)

		_ = tmp12176

		tmp12177 := MakeNative(func(__e *ControlFlow) {
			tmp12178 := Call(__e, PrimFunc(symps), V5832)

			__e.TailApply(PrimFunc(symeval), tmp12178)
			return

		}, 0)

		tmp12179 := MakeNative(func(__e *ControlFlow) {
			Z5833 := __e.Get(1)
			_ = Z5833
			__e.Return(V5832)
			return
		}, 1)

		__e.TailApply(try_1catch, tmp12177, tmp12179)
		return

	}, 1)

	tmp12180 := Call(__e, ns2_1set, symunprofile, tmp12173)

	_ = tmp12180

	tmp12181 := MakeNative(func(__e *ControlFlow) {
		V5834 := __e.Get(1)
		_ = V5834
		tmp12182 := PrimValue(symshen_4_dprofiled_d)

		__e.TailApply(PrimFunc(symelement_2), V5834, tmp12182)
		return

	}, 1)

	tmp12183 := Call(__e, ns2_1set, symshen_4profiled_2, tmp12181)

	_ = tmp12183

	tmp12184 := MakeNative(func(__e *ControlFlow) {
		V5835 := __e.Get(1)
		_ = V5835
		V5836 := __e.Get(2)
		_ = V5836
		V5837 := __e.Get(3)
		_ = V5837
		tmp12185 := PrimCons(symrun, Nil)

		tmp12186 := PrimCons(symget_1time, tmp12185)

		tmp12187 := PrimCons(symrun, Nil)

		tmp12188 := PrimCons(symget_1time, tmp12187)

		tmp12189 := PrimCons(symStart, Nil)

		tmp12190 := PrimCons(tmp12188, tmp12189)

		tmp12191 := PrimCons(sym_1, tmp12190)

		tmp12192 := PrimCons(V5835, Nil)

		tmp12193 := PrimCons(symshen_4get_1profile, tmp12192)

		tmp12194 := PrimCons(symFinish, Nil)

		tmp12195 := PrimCons(tmp12193, tmp12194)

		tmp12196 := PrimCons(sym_7, tmp12195)

		tmp12197 := PrimCons(tmp12196, Nil)

		tmp12198 := PrimCons(V5835, tmp12197)

		tmp12199 := PrimCons(symshen_4put_1profile, tmp12198)

		tmp12200 := PrimCons(symResult, Nil)

		tmp12201 := PrimCons(tmp12199, tmp12200)

		tmp12202 := PrimCons(symRecord, tmp12201)

		tmp12203 := PrimCons(symlet, tmp12202)

		tmp12204 := PrimCons(tmp12203, Nil)

		tmp12205 := PrimCons(tmp12191, tmp12204)

		tmp12206 := PrimCons(symFinish, tmp12205)

		tmp12207 := PrimCons(symlet, tmp12206)

		tmp12208 := PrimCons(tmp12207, Nil)

		tmp12209 := PrimCons(V5837, tmp12208)

		tmp12210 := PrimCons(symResult, tmp12209)

		tmp12211 := PrimCons(symlet, tmp12210)

		tmp12212 := PrimCons(tmp12211, Nil)

		tmp12213 := PrimCons(tmp12186, tmp12212)

		tmp12214 := PrimCons(symStart, tmp12213)

		__e.Return(PrimCons(symlet, tmp12214))
		return

	}, 3)

	tmp12215 := Call(__e, ns2_1set, symshen_4profile_1func, tmp12184)

	_ = tmp12215

	tmp12216 := MakeNative(func(__e *ControlFlow) {
		V5838 := __e.Get(1)
		_ = V5838
		tmp12217 := MakeNative(func(__e *ControlFlow) {
			W5839 := __e.Get(1)
			_ = W5839
			tmp12218 := MakeNative(func(__e *ControlFlow) {
				W5840 := __e.Get(1)
				_ = W5840
				__e.TailApply(PrimFunc(sym_8p), V5838, W5839)
				return
			}, 1)

			tmp12219 := Call(__e, PrimFunc(symshen_4put_1profile), V5838, MakeNumber(0))

			__e.TailApply(tmp12218, tmp12219)
			return

		}, 1)

		tmp12220 := Call(__e, PrimFunc(symshen_4get_1profile), V5838)

		__e.TailApply(tmp12217, tmp12220)
		return

	}, 1)

	tmp12221 := Call(__e, ns2_1set, symprofile_1results, tmp12216)

	_ = tmp12221

	tmp12222 := MakeNative(func(__e *ControlFlow) {
		V5841 := __e.Get(1)
		_ = V5841
		tmp12223 := MakeNative(func(__e *ControlFlow) {
			tmp12224 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symget), V5841, symprofile, tmp12224)
			return

		}, 0)

		tmp12225 := MakeNative(func(__e *ControlFlow) {
			Z5842 := __e.Get(1)
			_ = Z5842
			__e.Return(MakeNumber(0))
			return
		}, 1)

		__e.TailApply(try_1catch, tmp12223, tmp12225)
		return

	}, 1)

	tmp12226 := Call(__e, ns2_1set, symshen_4get_1profile, tmp12222)

	_ = tmp12226

	tmp12227 := MakeNative(func(__e *ControlFlow) {
		V5843 := __e.Get(1)
		_ = V5843
		V5844 := __e.Get(2)
		_ = V5844
		tmp12228 := PrimValue(sym_dproperty_1vector_d)

		__e.TailApply(PrimFunc(symput), V5843, symprofile, V5844, tmp12228)
		return

	}, 2)

	__e.TailApply(ns2_1set, symshen_4put_1profile, tmp12227)
	return

}, 0)
