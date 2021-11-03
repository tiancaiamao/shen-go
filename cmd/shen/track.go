package main

import . "github.com/tiancaiamao/shen-go/cora"

var TrackMain = MakeNative(func(__e *ControlFlow) {
	tmp11770 := MakeNative(func(__e *ControlFlow) {
		V4559 := __e.Get(1)
		_ = V4559
		tmp11771 := Call(__e, PrimNS2Value(symshen_4app), V4559, MakeString(";\n"), symshen_4a)

		tmp11772 := PrimStringConcat(MakeString("partial function "), tmp11771)

		tmp11773 := Call(__e, PrimNS2Value(symstoutput))

		tmp11774 := Call(__e, PrimNS2Value(sympr), tmp11772, tmp11773)

		_ = tmp11774

		tmp11783 := Call(__e, PrimNS2Value(symshen_4tracked_2), V4559)

		tmp11784 := PrimNot(tmp11783)

		var ifres11778 Obj

		if True == tmp11784 {
			tmp11780 := Call(__e, PrimNS2Value(symshen_4app), V4559, MakeString("? "), symshen_4a)

			tmp11781 := PrimStringConcat(MakeString("track "), tmp11780)

			tmp11782 := Call(__e, PrimNS2Value(symy_1or_1n_2), tmp11781)

			var ifres11779 Obj

			if True == tmp11782 {
				ifres11779 = True

			} else {
				ifres11779 = False

			}

			ifres11778 = ifres11779

		} else {
			ifres11778 = False

		}

		var ifres11775 Obj

		if True == ifres11778 {
			tmp11776 := Call(__e, PrimNS2Value(symps), V4559)

			tmp11777 := Call(__e, PrimNS2Value(symshen_4track_1function), tmp11776)

			ifres11775 = tmp11777

		} else {
			ifres11775 = symshen_4ok

		}

		_ = ifres11775

		__e.Return(PrimSimpleError(MakeString("aborted")))
		return

	}, 1)

	tmp11785 := Call(__e, PrimNS2Value(symdef), symshen_4f_1error, tmp11770)

	_ = tmp11785

	tmp11786 := MakeNative(func(__e *ControlFlow) {
		V4560 := __e.Get(1)
		_ = V4560
		tmp11787 := PrimNS3Value(symshen_4_dtracking_d)

		__e.TailApply(PrimNS2Value(symelement_2), V4560, tmp11787)
		return

	}, 1)

	tmp11788 := Call(__e, PrimNS2Value(symdef), symshen_4tracked_2, tmp11786)

	_ = tmp11788

	tmp11789 := MakeNative(func(__e *ControlFlow) {
		V4561 := __e.Get(1)
		_ = V4561
		tmp11790 := MakeNative(func(__e *ControlFlow) {
			Source := __e.Get(1)
			_ = Source
			__e.TailApply(PrimNS2Value(symshen_4track_1function), Source)
			return
		}, 1)

		tmp11791 := Call(__e, PrimNS2Value(symps), V4561)

		__e.TailApply(tmp11790, tmp11791)
		return

	}, 1)

	tmp11792 := Call(__e, PrimNS2Value(symdef), symtrack, tmp11789)

	_ = tmp11792

	tmp11793 := MakeNative(func(__e *ControlFlow) {
		V4564 := __e.Get(1)
		_ = V4564
		tmp11847 := PrimIsPair(V4564)

		var ifres11821 Obj

		if True == tmp11847 {
			tmp11845 := PrimHead(V4564)

			tmp11846 := PrimEqual(symdefun, tmp11845)

			var ifres11823 Obj

			if True == tmp11846 {
				tmp11843 := PrimTail(V4564)

				tmp11844 := PrimIsPair(tmp11843)

				var ifres11825 Obj

				if True == tmp11844 {
					tmp11840 := PrimTail(V4564)

					tmp11841 := PrimTail(tmp11840)

					tmp11842 := PrimIsPair(tmp11841)

					var ifres11827 Obj

					if True == tmp11842 {
						tmp11836 := PrimTail(V4564)

						tmp11837 := PrimTail(tmp11836)

						tmp11838 := PrimTail(tmp11837)

						tmp11839 := PrimIsPair(tmp11838)

						var ifres11829 Obj

						if True == tmp11839 {
							tmp11831 := PrimTail(V4564)

							tmp11832 := PrimTail(tmp11831)

							tmp11833 := PrimTail(tmp11832)

							tmp11834 := PrimTail(tmp11833)

							tmp11835 := PrimEqual(Nil, tmp11834)

							var ifres11830 Obj

							if True == tmp11835 {
								ifres11830 = True

							} else {
								ifres11830 = False

							}

							ifres11829 = ifres11830

						} else {
							ifres11829 = False

						}

						var ifres11828 Obj

						if True == ifres11829 {
							ifres11828 = True

						} else {
							ifres11828 = False

						}

						ifres11827 = ifres11828

					} else {
						ifres11827 = False

					}

					var ifres11826 Obj

					if True == ifres11827 {
						ifres11826 = True

					} else {
						ifres11826 = False

					}

					ifres11825 = ifres11826

				} else {
					ifres11825 = False

				}

				var ifres11824 Obj

				if True == ifres11825 {
					ifres11824 = True

				} else {
					ifres11824 = False

				}

				ifres11823 = ifres11824

			} else {
				ifres11823 = False

			}

			var ifres11822 Obj

			if True == ifres11823 {
				ifres11822 = True

			} else {
				ifres11822 = False

			}

			ifres11821 = ifres11822

		} else {
			ifres11821 = False

		}

		if True == ifres11821 {
			tmp11795 := MakeNative(func(__e *ControlFlow) {
				KL := __e.Get(1)
				_ = KL
				tmp11796 := MakeNative(func(__e *ControlFlow) {
					Ob := __e.Get(1)
					_ = Ob
					tmp11797 := MakeNative(func(__e *ControlFlow) {
						Tr := __e.Get(1)
						_ = Tr
						__e.Return(Ob)
						return
					}, 1)

					tmp11798 := PrimNS3Value(symshen_4_dtracking_d)

					tmp11799 := PrimCons(Ob, tmp11798)

					tmp11800 := PrimNS3Set(symshen_4_dtracking_d, tmp11799)

					__e.TailApply(tmp11797, tmp11800)
					return

				}, 1)

				tmp11801 := Call(__e, PrimNS2Value(symeval_1kl), KL)

				__e.TailApply(tmp11796, tmp11801)
				return

			}, 1)

			tmp11802 := PrimTail(V4564)

			tmp11803 := PrimHead(tmp11802)

			tmp11804 := PrimTail(V4564)

			tmp11805 := PrimTail(tmp11804)

			tmp11806 := PrimHead(tmp11805)

			tmp11807 := PrimTail(V4564)

			tmp11808 := PrimHead(tmp11807)

			tmp11809 := PrimTail(V4564)

			tmp11810 := PrimTail(tmp11809)

			tmp11811 := PrimHead(tmp11810)

			tmp11812 := PrimTail(V4564)

			tmp11813 := PrimTail(tmp11812)

			tmp11814 := PrimTail(tmp11813)

			tmp11815 := PrimHead(tmp11814)

			tmp11816 := Call(__e, PrimNS2Value(symshen_4insert_1tracking_1code), tmp11808, tmp11811, tmp11815)

			tmp11817 := PrimCons(tmp11816, Nil)

			tmp11818 := PrimCons(tmp11806, tmp11817)

			tmp11819 := PrimCons(tmp11803, tmp11818)

			tmp11820 := PrimCons(symdefun, tmp11819)

			__e.TailApply(tmp11795, tmp11820)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("implementation error in shen.track-function")))
			return
		}

	}, 1)

	tmp11848 := Call(__e, PrimNS2Value(symdef), symshen_4track_1function, tmp11793)

	_ = tmp11848

	tmp11849 := MakeNative(func(__e *ControlFlow) {
		V4565 := __e.Get(1)
		_ = V4565
		V4566 := __e.Get(2)
		_ = V4566
		V4567 := __e.Get(3)
		_ = V4567
		tmp11850 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11851 := PrimCons(symvalue, tmp11850)

		tmp11852 := PrimCons(MakeNumber(1), Nil)

		tmp11853 := PrimCons(tmp11851, tmp11852)

		tmp11854 := PrimCons(sym_7, tmp11853)

		tmp11855 := PrimCons(tmp11854, Nil)

		tmp11856 := PrimCons(symshen_4_dcall_d, tmp11855)

		tmp11857 := PrimCons(symset, tmp11856)

		tmp11858 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11859 := PrimCons(symvalue, tmp11858)

		tmp11860 := Call(__e, PrimNS2Value(symshen_4prolog_1track), V4567, V4566)

		tmp11861 := Call(__e, PrimNS2Value(symshen_4cons_1form), tmp11860)

		tmp11862 := PrimCons(tmp11861, Nil)

		tmp11863 := PrimCons(V4565, tmp11862)

		tmp11864 := PrimCons(tmp11859, tmp11863)

		tmp11865 := PrimCons(symshen_4input_1track, tmp11864)

		tmp11866 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

		tmp11867 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp11868 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11869 := PrimCons(symvalue, tmp11868)

		tmp11870 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp11871 := PrimCons(tmp11870, Nil)

		tmp11872 := PrimCons(V4565, tmp11871)

		tmp11873 := PrimCons(tmp11869, tmp11872)

		tmp11874 := PrimCons(symshen_4output_1track, tmp11873)

		tmp11875 := PrimCons(symshen_4_dcall_d, Nil)

		tmp11876 := PrimCons(symvalue, tmp11875)

		tmp11877 := PrimCons(MakeNumber(1), Nil)

		tmp11878 := PrimCons(tmp11876, tmp11877)

		tmp11879 := PrimCons(sym_1, tmp11878)

		tmp11880 := PrimCons(tmp11879, Nil)

		tmp11881 := PrimCons(symshen_4_dcall_d, tmp11880)

		tmp11882 := PrimCons(symset, tmp11881)

		tmp11883 := PrimCons(symshen_4terpri_1or_1read_1char, Nil)

		tmp11884 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp11885 := PrimCons(tmp11884, Nil)

		tmp11886 := PrimCons(tmp11883, tmp11885)

		tmp11887 := PrimCons(symdo, tmp11886)

		tmp11888 := PrimCons(tmp11887, Nil)

		tmp11889 := PrimCons(tmp11882, tmp11888)

		tmp11890 := PrimCons(symdo, tmp11889)

		tmp11891 := PrimCons(tmp11890, Nil)

		tmp11892 := PrimCons(tmp11874, tmp11891)

		tmp11893 := PrimCons(symdo, tmp11892)

		tmp11894 := PrimCons(tmp11893, Nil)

		tmp11895 := PrimCons(V4567, tmp11894)

		tmp11896 := PrimCons(tmp11867, tmp11895)

		tmp11897 := PrimCons(symlet, tmp11896)

		tmp11898 := PrimCons(tmp11897, Nil)

		tmp11899 := PrimCons(tmp11866, tmp11898)

		tmp11900 := PrimCons(symdo, tmp11899)

		tmp11901 := PrimCons(tmp11900, Nil)

		tmp11902 := PrimCons(tmp11865, tmp11901)

		tmp11903 := PrimCons(symdo, tmp11902)

		tmp11904 := PrimCons(tmp11903, Nil)

		tmp11905 := PrimCons(tmp11857, tmp11904)

		__e.Return(PrimCons(symdo, tmp11905))
		return

	}, 3)

	tmp11906 := Call(__e, PrimNS2Value(symdef), symshen_4insert_1tracking_1code, tmp11849)

	_ = tmp11906

	tmp11907 := MakeNative(func(__e *ControlFlow) {
		V4568 := __e.Get(1)
		_ = V4568
		V4569 := __e.Get(2)
		_ = V4569
		tmp11910 := Call(__e, PrimNS2Value(symoccurrences), symshen_4incinfs, V4568)

		tmp11911 := PrimEqual(tmp11910, MakeNumber(0))

		if True == tmp11911 {
			__e.Return(V4569)
			return
		} else {
			tmp11909 := Call(__e, PrimNS2Value(symshen_4vector_1parameter), V4569)

			__e.TailApply(PrimNS2Value(symshen_4vector_1dereference), V4569, tmp11909)
			return

		}

	}, 2)

	tmp11912 := Call(__e, PrimNS2Value(symdef), symshen_4prolog_1track, tmp11907)

	_ = tmp11912

	tmp11913 := MakeNative(func(__e *ControlFlow) {
		V4572 := __e.Get(1)
		_ = V4572
		tmp11942 := PrimEqual(Nil, V4572)

		if True == tmp11942 {
			__e.Return(Nil)
			return
		} else {
			tmp11941 := PrimIsPair(V4572)

			var ifres11919 Obj

			if True == tmp11941 {
				tmp11939 := PrimTail(V4572)

				tmp11940 := PrimIsPair(tmp11939)

				var ifres11921 Obj

				if True == tmp11940 {
					tmp11936 := PrimTail(V4572)

					tmp11937 := PrimTail(tmp11936)

					tmp11938 := PrimIsPair(tmp11937)

					var ifres11923 Obj

					if True == tmp11938 {
						tmp11932 := PrimTail(V4572)

						tmp11933 := PrimTail(tmp11932)

						tmp11934 := PrimTail(tmp11933)

						tmp11935 := PrimIsPair(tmp11934)

						var ifres11925 Obj

						if True == tmp11935 {
							tmp11927 := PrimTail(V4572)

							tmp11928 := PrimTail(tmp11927)

							tmp11929 := PrimTail(tmp11928)

							tmp11930 := PrimTail(tmp11929)

							tmp11931 := PrimEqual(Nil, tmp11930)

							var ifres11926 Obj

							if True == tmp11931 {
								ifres11926 = True

							} else {
								ifres11926 = False

							}

							ifres11925 = ifres11926

						} else {
							ifres11925 = False

						}

						var ifres11924 Obj

						if True == ifres11925 {
							ifres11924 = True

						} else {
							ifres11924 = False

						}

						ifres11923 = ifres11924

					} else {
						ifres11923 = False

					}

					var ifres11922 Obj

					if True == ifres11923 {
						ifres11922 = True

					} else {
						ifres11922 = False

					}

					ifres11921 = ifres11922

				} else {
					ifres11921 = False

				}

				var ifres11920 Obj

				if True == ifres11921 {
					ifres11920 = True

				} else {
					ifres11920 = False

				}

				ifres11919 = ifres11920

			} else {
				ifres11919 = False

			}

			if True == ifres11919 {
				__e.Return(PrimHead(V4572))
				return
			} else {
				tmp11918 := PrimIsPair(V4572)

				if True == tmp11918 {
					tmp11917 := PrimTail(V4572)

					__e.TailApply(PrimNS2Value(symshen_4vector_1parameter), tmp11917)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4vector_1parameter)
					return
				}

			}

		}

	}, 1)

	tmp11943 := Call(__e, PrimNS2Value(symdef), symshen_4vector_1parameter, tmp11913)

	_ = tmp11943

	tmp11944 := MakeNative(func(__e *ControlFlow) {
		V4575 := __e.Get(1)
		_ = V4575
		V4576 := __e.Get(2)
		_ = V4576
		tmp11978 := PrimEqual(Nil, V4576)

		if True == tmp11978 {
			__e.Return(V4575)
			return
		} else {
			tmp11977 := PrimIsPair(V4575)

			var ifres11955 Obj

			if True == tmp11977 {
				tmp11975 := PrimTail(V4575)

				tmp11976 := PrimIsPair(tmp11975)

				var ifres11957 Obj

				if True == tmp11976 {
					tmp11972 := PrimTail(V4575)

					tmp11973 := PrimTail(tmp11972)

					tmp11974 := PrimIsPair(tmp11973)

					var ifres11959 Obj

					if True == tmp11974 {
						tmp11968 := PrimTail(V4575)

						tmp11969 := PrimTail(tmp11968)

						tmp11970 := PrimTail(tmp11969)

						tmp11971 := PrimIsPair(tmp11970)

						var ifres11961 Obj

						if True == tmp11971 {
							tmp11963 := PrimTail(V4575)

							tmp11964 := PrimTail(tmp11963)

							tmp11965 := PrimTail(tmp11964)

							tmp11966 := PrimTail(tmp11965)

							tmp11967 := PrimEqual(Nil, tmp11966)

							var ifres11962 Obj

							if True == tmp11967 {
								ifres11962 = True

							} else {
								ifres11962 = False

							}

							ifres11961 = ifres11962

						} else {
							ifres11961 = False

						}

						var ifres11960 Obj

						if True == ifres11961 {
							ifres11960 = True

						} else {
							ifres11960 = False

						}

						ifres11959 = ifres11960

					} else {
						ifres11959 = False

					}

					var ifres11958 Obj

					if True == ifres11959 {
						ifres11958 = True

					} else {
						ifres11958 = False

					}

					ifres11957 = ifres11958

				} else {
					ifres11957 = False

				}

				var ifres11956 Obj

				if True == ifres11957 {
					ifres11956 = True

				} else {
					ifres11956 = False

				}

				ifres11955 = ifres11956

			} else {
				ifres11955 = False

			}

			if True == ifres11955 {
				__e.Return(V4575)
				return
			} else {
				tmp11954 := PrimIsPair(V4575)

				if True == tmp11954 {
					tmp11948 := PrimHead(V4575)

					tmp11949 := PrimCons(V4576, Nil)

					tmp11950 := PrimCons(tmp11948, tmp11949)

					tmp11951 := PrimCons(symshen_4deref, tmp11950)

					tmp11952 := PrimTail(V4575)

					tmp11953 := Call(__e, PrimNS2Value(symshen_4vector_1dereference), tmp11952, V4576)

					__e.Return(PrimCons(tmp11951, tmp11953))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4vector_1dereference)
					return
				}

			}

		}

	}, 2)

	tmp11979 := Call(__e, PrimNS2Value(symdef), symshen_4vector_1dereference, tmp11944)

	_ = tmp11979

	tmp11980 := MakeNative(func(__e *ControlFlow) {
		V4579 := __e.Get(1)
		_ = V4579
		tmp11984 := PrimEqual(sym_7, V4579)

		if True == tmp11984 {
			__e.Return(PrimNS3Set(symshen_4_dstep_d, True))
			return
		} else {
			tmp11983 := PrimEqual(sym_1, V4579)

			if True == tmp11983 {
				__e.Return(PrimNS3Set(symshen_4_dstep_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("step expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp11985 := Call(__e, PrimNS2Value(symdef), symstep, tmp11980)

	_ = tmp11985

	tmp11986 := MakeNative(func(__e *ControlFlow) {
		V4582 := __e.Get(1)
		_ = V4582
		tmp11990 := PrimEqual(sym_7, V4582)

		if True == tmp11990 {
			__e.Return(PrimNS3Set(symshen_4_dspy_d, True))
			return
		} else {
			tmp11989 := PrimEqual(sym_1, V4582)

			if True == tmp11989 {
				__e.Return(PrimNS3Set(symshen_4_dspy_d, False))
				return
			} else {
				__e.Return(PrimSimpleError(MakeString("spy expects a + or a -.\n")))
				return
			}

		}

	}, 1)

	tmp11991 := Call(__e, PrimNS2Value(symdef), symspy, tmp11986)

	_ = tmp11991

	tmp11992 := MakeNative(func(__e *ControlFlow) {
		tmp11996 := PrimNS3Value(symshen_4_dstep_d)

		if True == tmp11996 {
			tmp11994 := PrimNS3Value(sym_dstinput_d)

			tmp11995 := PrimReadByte(tmp11994)

			__e.TailApply(PrimNS2Value(symshen_4check_1byte), tmp11995)
			return

		} else {
			__e.TailApply(PrimNS2Value(symnl), MakeNumber(1))
			return
		}

	}, 0)

	tmp11997 := Call(__e, PrimNS2Value(symdef), symshen_4terpri_1or_1read_1char, tmp11992)

	_ = tmp11997

	tmp11998 := MakeNative(func(__e *ControlFlow) {
		V4585 := __e.Get(1)
		_ = V4585
		tmp12000 := PrimEqual(MakeNumber(94), V4585)

		if True == tmp12000 {
			__e.Return(PrimSimpleError(MakeString("aborted")))
			return
		} else {
			__e.Return(True)
			return
		}

	}, 1)

	tmp12001 := Call(__e, PrimNS2Value(symdef), symshen_4check_1byte, tmp11998)

	_ = tmp12001

	tmp12002 := MakeNative(func(__e *ControlFlow) {
		V4586 := __e.Get(1)
		_ = V4586
		V4587 := __e.Get(2)
		_ = V4587
		V4588 := __e.Get(3)
		_ = V4588
		tmp12003 := Call(__e, PrimNS2Value(symshen_4spaces), V4586)

		tmp12004 := Call(__e, PrimNS2Value(symshen_4spaces), V4586)

		tmp12005 := Call(__e, PrimNS2Value(symshen_4app), tmp12004, MakeString(""), symshen_4a)

		tmp12006 := PrimStringConcat(MakeString(" \n"), tmp12005)

		tmp12007 := Call(__e, PrimNS2Value(symshen_4app), V4587, tmp12006, symshen_4a)

		tmp12008 := PrimStringConcat(MakeString("> Inputs to "), tmp12007)

		tmp12009 := Call(__e, PrimNS2Value(symshen_4app), V4586, tmp12008, symshen_4a)

		tmp12010 := PrimStringConcat(MakeString("<"), tmp12009)

		tmp12011 := Call(__e, PrimNS2Value(symshen_4app), tmp12003, tmp12010, symshen_4a)

		tmp12012 := PrimStringConcat(MakeString("\n"), tmp12011)

		tmp12013 := Call(__e, PrimNS2Value(symstoutput))

		tmp12014 := Call(__e, PrimNS2Value(sympr), tmp12012, tmp12013)

		_ = tmp12014

		__e.TailApply(PrimNS2Value(symshen_4recursively_1print), V4588)
		return

	}, 3)

	tmp12015 := Call(__e, PrimNS2Value(symdef), symshen_4input_1track, tmp12002)

	_ = tmp12015

	tmp12016 := MakeNative(func(__e *ControlFlow) {
		V4591 := __e.Get(1)
		_ = V4591
		tmp12026 := PrimEqual(Nil, V4591)

		if True == tmp12026 {
			tmp12025 := Call(__e, PrimNS2Value(symstoutput))

			__e.TailApply(PrimNS2Value(sympr), MakeString(" ==>"), tmp12025)
			return

		} else {
			tmp12024 := PrimIsPair(V4591)

			if True == tmp12024 {
				tmp12019 := PrimHead(V4591)

				tmp12020 := Call(__e, PrimNS2Value(symprint), tmp12019)

				_ = tmp12020

				tmp12021 := Call(__e, PrimNS2Value(symstoutput))

				tmp12022 := Call(__e, PrimNS2Value(sympr), MakeString(", "), tmp12021)

				_ = tmp12022

				tmp12023 := PrimTail(V4591)

				__e.TailApply(PrimNS2Value(symshen_4recursively_1print), tmp12023)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.recursively-print")))
				return
			}

		}

	}, 1)

	tmp12027 := Call(__e, PrimNS2Value(symdef), symshen_4recursively_1print, tmp12016)

	_ = tmp12027

	tmp12028 := MakeNative(func(__e *ControlFlow) {
		V4592 := __e.Get(1)
		_ = V4592
		tmp12032 := PrimEqual(MakeNumber(0), V4592)

		if True == tmp12032 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp12030 := PrimNumberSubtract(V4592, MakeNumber(1))

			tmp12031 := Call(__e, PrimNS2Value(symshen_4spaces), tmp12030)

			__e.Return(PrimStringConcat(MakeString(" "), tmp12031))
			return

		}

	}, 1)

	tmp12033 := Call(__e, PrimNS2Value(symdef), symshen_4spaces, tmp12028)

	_ = tmp12033

	tmp12034 := MakeNative(func(__e *ControlFlow) {
		V4593 := __e.Get(1)
		_ = V4593
		V4594 := __e.Get(2)
		_ = V4594
		V4595 := __e.Get(3)
		_ = V4595
		tmp12035 := Call(__e, PrimNS2Value(symshen_4spaces), V4593)

		tmp12036 := Call(__e, PrimNS2Value(symshen_4spaces), V4593)

		tmp12037 := Call(__e, PrimNS2Value(symshen_4app), V4595, MakeString(""), symshen_4s)

		tmp12038 := PrimStringConcat(MakeString("==> "), tmp12037)

		tmp12039 := Call(__e, PrimNS2Value(symshen_4app), tmp12036, tmp12038, symshen_4a)

		tmp12040 := PrimStringConcat(MakeString(" \n"), tmp12039)

		tmp12041 := Call(__e, PrimNS2Value(symshen_4app), V4594, tmp12040, symshen_4a)

		tmp12042 := PrimStringConcat(MakeString("> Output from "), tmp12041)

		tmp12043 := Call(__e, PrimNS2Value(symshen_4app), V4593, tmp12042, symshen_4a)

		tmp12044 := PrimStringConcat(MakeString("<"), tmp12043)

		tmp12045 := Call(__e, PrimNS2Value(symshen_4app), tmp12035, tmp12044, symshen_4a)

		tmp12046 := PrimStringConcat(MakeString("\n"), tmp12045)

		tmp12047 := Call(__e, PrimNS2Value(symstoutput))

		__e.TailApply(PrimNS2Value(sympr), tmp12046, tmp12047)
		return

	}, 3)

	tmp12048 := Call(__e, PrimNS2Value(symdef), symshen_4output_1track, tmp12034)

	_ = tmp12048

	tmp12049 := MakeNative(func(__e *ControlFlow) {
		V4596 := __e.Get(1)
		_ = V4596
		tmp12050 := PrimNS3Value(symshen_4_dtracking_d)

		tmp12051 := Call(__e, PrimNS2Value(symremove), V4596, tmp12050)

		tmp12052 := PrimNS3Set(symshen_4_dtracking_d, tmp12051)

		_ = tmp12052

		tmp12053 := MakeNative(func(__e *ControlFlow) {
			tmp12054 := Call(__e, PrimNS2Value(symps), V4596)

			__e.TailApply(PrimNS2Value(symeval), tmp12054)
			return

		}, 0)

		tmp12055 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V4596)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp12053, tmp12055)
		return

	}, 1)

	tmp12056 := Call(__e, PrimNS2Value(symdef), symuntrack, tmp12049)

	_ = tmp12056

	tmp12057 := MakeNative(func(__e *ControlFlow) {
		V4597 := __e.Get(1)
		_ = V4597
		V4598 := __e.Get(2)
		_ = V4598
		__e.TailApply(PrimNS2Value(symshen_4remove_1h), V4597, V4598, Nil)
		return
	}, 2)

	tmp12058 := Call(__e, PrimNS2Value(symdef), symremove, tmp12057)

	_ = tmp12058

	tmp12059 := MakeNative(func(__e *ControlFlow) {
		V4608 := __e.Get(1)
		_ = V4608
		V4609 := __e.Get(2)
		_ = V4609
		V4610 := __e.Get(3)
		_ = V4610
		tmp12074 := PrimEqual(Nil, V4609)

		if True == tmp12074 {
			__e.TailApply(PrimNS2Value(symreverse), V4610)
			return
		} else {
			tmp12073 := PrimIsPair(V4609)

			var ifres12069 Obj

			if True == tmp12073 {
				tmp12071 := PrimHead(V4609)

				tmp12072 := PrimEqual(V4608, tmp12071)

				var ifres12070 Obj

				if True == tmp12072 {
					ifres12070 = True

				} else {
					ifres12070 = False

				}

				ifres12069 = ifres12070

			} else {
				ifres12069 = False

			}

			if True == ifres12069 {
				tmp12067 := PrimHead(V4609)

				tmp12068 := PrimTail(V4609)

				__e.TailApply(PrimNS2Value(symshen_4remove_1h), tmp12067, tmp12068, V4610)
				return

			} else {
				tmp12066 := PrimIsPair(V4609)

				if True == tmp12066 {
					tmp12063 := PrimTail(V4609)

					tmp12064 := PrimHead(V4609)

					tmp12065 := PrimCons(tmp12064, V4610)

					__e.TailApply(PrimNS2Value(symshen_4remove_1h), V4608, tmp12063, tmp12065)
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.remove-h")))
					return
				}

			}

		}

	}, 3)

	tmp12075 := Call(__e, PrimNS2Value(symdef), symshen_4remove_1h, tmp12059)

	_ = tmp12075

	tmp12076 := MakeNative(func(__e *ControlFlow) {
		V4611 := __e.Get(1)
		_ = V4611
		tmp12077 := PrimNS3Value(symshen_4_dprofiled_d)

		tmp12078 := PrimCons(V4611, tmp12077)

		tmp12079 := PrimNS3Set(symshen_4_dprofiled_d, tmp12078)

		_ = tmp12079

		tmp12080 := Call(__e, PrimNS2Value(symps), V4611)

		__e.TailApply(PrimNS2Value(symshen_4profile_1help), tmp12080)
		return

	}, 1)

	tmp12081 := Call(__e, PrimNS2Value(symdef), symprofile, tmp12076)

	_ = tmp12081

	tmp12082 := MakeNative(func(__e *ControlFlow) {
		V4614 := __e.Get(1)
		_ = V4614
		tmp12152 := PrimIsPair(V4614)

		var ifres12126 Obj

		if True == tmp12152 {
			tmp12150 := PrimHead(V4614)

			tmp12151 := PrimEqual(symdefun, tmp12150)

			var ifres12128 Obj

			if True == tmp12151 {
				tmp12148 := PrimTail(V4614)

				tmp12149 := PrimIsPair(tmp12148)

				var ifres12130 Obj

				if True == tmp12149 {
					tmp12145 := PrimTail(V4614)

					tmp12146 := PrimTail(tmp12145)

					tmp12147 := PrimIsPair(tmp12146)

					var ifres12132 Obj

					if True == tmp12147 {
						tmp12141 := PrimTail(V4614)

						tmp12142 := PrimTail(tmp12141)

						tmp12143 := PrimTail(tmp12142)

						tmp12144 := PrimIsPair(tmp12143)

						var ifres12134 Obj

						if True == tmp12144 {
							tmp12136 := PrimTail(V4614)

							tmp12137 := PrimTail(tmp12136)

							tmp12138 := PrimTail(tmp12137)

							tmp12139 := PrimTail(tmp12138)

							tmp12140 := PrimEqual(Nil, tmp12139)

							var ifres12135 Obj

							if True == tmp12140 {
								ifres12135 = True

							} else {
								ifres12135 = False

							}

							ifres12134 = ifres12135

						} else {
							ifres12134 = False

						}

						var ifres12133 Obj

						if True == ifres12134 {
							ifres12133 = True

						} else {
							ifres12133 = False

						}

						ifres12132 = ifres12133

					} else {
						ifres12132 = False

					}

					var ifres12131 Obj

					if True == ifres12132 {
						ifres12131 = True

					} else {
						ifres12131 = False

					}

					ifres12130 = ifres12131

				} else {
					ifres12130 = False

				}

				var ifres12129 Obj

				if True == ifres12130 {
					ifres12129 = True

				} else {
					ifres12129 = False

				}

				ifres12128 = ifres12129

			} else {
				ifres12128 = False

			}

			var ifres12127 Obj

			if True == ifres12128 {
				ifres12127 = True

			} else {
				ifres12127 = False

			}

			ifres12126 = ifres12127

		} else {
			ifres12126 = False

		}

		if True == ifres12126 {
			tmp12084 := MakeNative(func(__e *ControlFlow) {
				G := __e.Get(1)
				_ = G
				tmp12085 := MakeNative(func(__e *ControlFlow) {
					Profile := __e.Get(1)
					_ = Profile
					tmp12086 := MakeNative(func(__e *ControlFlow) {
						Def := __e.Get(1)
						_ = Def
						tmp12087 := MakeNative(func(__e *ControlFlow) {
							CompileProfile := __e.Get(1)
							_ = CompileProfile
							tmp12088 := MakeNative(func(__e *ControlFlow) {
								CompileG := __e.Get(1)
								_ = CompileG
								tmp12089 := PrimTail(V4614)

								__e.Return(PrimHead(tmp12089))
								return

							}, 1)

							tmp12090 := Call(__e, PrimNS2Value(symeval_1kl), Def)

							__e.TailApply(tmp12088, tmp12090)
							return

						}, 1)

						tmp12091 := Call(__e, PrimNS2Value(symeval_1kl), Profile)

						__e.TailApply(tmp12087, tmp12091)
						return

					}, 1)

					tmp12092 := PrimTail(V4614)

					tmp12093 := PrimTail(tmp12092)

					tmp12094 := PrimHead(tmp12093)

					tmp12095 := PrimTail(V4614)

					tmp12096 := PrimHead(tmp12095)

					tmp12097 := PrimTail(V4614)

					tmp12098 := PrimTail(tmp12097)

					tmp12099 := PrimTail(tmp12098)

					tmp12100 := PrimHead(tmp12099)

					tmp12101 := Call(__e, PrimNS2Value(symsubst), G, tmp12096, tmp12100)

					tmp12102 := PrimCons(tmp12101, Nil)

					tmp12103 := PrimCons(tmp12094, tmp12102)

					tmp12104 := PrimCons(G, tmp12103)

					tmp12105 := PrimCons(symdefun, tmp12104)

					__e.TailApply(tmp12086, tmp12105)
					return

				}, 1)

				tmp12106 := PrimTail(V4614)

				tmp12107 := PrimHead(tmp12106)

				tmp12108 := PrimTail(V4614)

				tmp12109 := PrimTail(tmp12108)

				tmp12110 := PrimHead(tmp12109)

				tmp12111 := PrimTail(V4614)

				tmp12112 := PrimHead(tmp12111)

				tmp12113 := PrimTail(V4614)

				tmp12114 := PrimTail(tmp12113)

				tmp12115 := PrimHead(tmp12114)

				tmp12116 := PrimTail(V4614)

				tmp12117 := PrimTail(tmp12116)

				tmp12118 := PrimHead(tmp12117)

				tmp12119 := PrimCons(G, tmp12118)

				tmp12120 := Call(__e, PrimNS2Value(symshen_4profile_1func), tmp12112, tmp12115, tmp12119)

				tmp12121 := PrimCons(tmp12120, Nil)

				tmp12122 := PrimCons(tmp12110, tmp12121)

				tmp12123 := PrimCons(tmp12107, tmp12122)

				tmp12124 := PrimCons(symdefun, tmp12123)

				__e.TailApply(tmp12085, tmp12124)
				return

			}, 1)

			tmp12125 := Call(__e, PrimNS2Value(symgensym), symshen_4f)

			__e.TailApply(tmp12084, tmp12125)
			return

		} else {
			__e.Return(PrimSimpleError(MakeString("Cannot profile.\n")))
			return
		}

	}, 1)

	tmp12153 := Call(__e, PrimNS2Value(symdef), symshen_4profile_1help, tmp12082)

	_ = tmp12153

	tmp12154 := MakeNative(func(__e *ControlFlow) {
		V4615 := __e.Get(1)
		_ = V4615
		tmp12155 := PrimNS3Value(symshen_4_dprofiled_d)

		tmp12156 := Call(__e, PrimNS2Value(symremove), V4615, tmp12155)

		tmp12157 := PrimNS3Set(symshen_4_dprofiled_d, tmp12156)

		_ = tmp12157

		tmp12158 := MakeNative(func(__e *ControlFlow) {
			tmp12159 := Call(__e, PrimNS2Value(symps), V4615)

			__e.TailApply(PrimNS2Value(symeval), tmp12159)
			return

		}, 0)

		tmp12160 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V4615)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp12158, tmp12160)
		return

	}, 1)

	tmp12161 := Call(__e, PrimNS2Value(symdef), symunprofile, tmp12154)

	_ = tmp12161

	tmp12162 := MakeNative(func(__e *ControlFlow) {
		V4616 := __e.Get(1)
		_ = V4616
		tmp12163 := PrimNS3Value(symshen_4_dprofiled_d)

		__e.TailApply(PrimNS2Value(symelement_2), V4616, tmp12163)
		return

	}, 1)

	tmp12164 := Call(__e, PrimNS2Value(symdef), symshen_4profiled_2, tmp12162)

	_ = tmp12164

	tmp12165 := MakeNative(func(__e *ControlFlow) {
		V4617 := __e.Get(1)
		_ = V4617
		V4618 := __e.Get(2)
		_ = V4618
		V4619 := __e.Get(3)
		_ = V4619
		tmp12166 := Call(__e, PrimNS2Value(symprotect), symStart)

		tmp12167 := PrimCons(symrun, Nil)

		tmp12168 := PrimCons(symget_1time, tmp12167)

		tmp12169 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp12170 := Call(__e, PrimNS2Value(symprotect), symFinish)

		tmp12171 := PrimCons(symrun, Nil)

		tmp12172 := PrimCons(symget_1time, tmp12171)

		tmp12173 := Call(__e, PrimNS2Value(symprotect), symStart)

		tmp12174 := PrimCons(tmp12173, Nil)

		tmp12175 := PrimCons(tmp12172, tmp12174)

		tmp12176 := PrimCons(sym_1, tmp12175)

		tmp12177 := Call(__e, PrimNS2Value(symprotect), symRecord)

		tmp12178 := PrimCons(V4617, Nil)

		tmp12179 := PrimCons(symshen_4get_1profile, tmp12178)

		tmp12180 := Call(__e, PrimNS2Value(symprotect), symFinish)

		tmp12181 := PrimCons(tmp12180, Nil)

		tmp12182 := PrimCons(tmp12179, tmp12181)

		tmp12183 := PrimCons(sym_7, tmp12182)

		tmp12184 := PrimCons(tmp12183, Nil)

		tmp12185 := PrimCons(V4617, tmp12184)

		tmp12186 := PrimCons(symshen_4put_1profile, tmp12185)

		tmp12187 := Call(__e, PrimNS2Value(symprotect), symResult)

		tmp12188 := PrimCons(tmp12187, Nil)

		tmp12189 := PrimCons(tmp12186, tmp12188)

		tmp12190 := PrimCons(tmp12177, tmp12189)

		tmp12191 := PrimCons(symlet, tmp12190)

		tmp12192 := PrimCons(tmp12191, Nil)

		tmp12193 := PrimCons(tmp12176, tmp12192)

		tmp12194 := PrimCons(tmp12170, tmp12193)

		tmp12195 := PrimCons(symlet, tmp12194)

		tmp12196 := PrimCons(tmp12195, Nil)

		tmp12197 := PrimCons(V4619, tmp12196)

		tmp12198 := PrimCons(tmp12169, tmp12197)

		tmp12199 := PrimCons(symlet, tmp12198)

		tmp12200 := PrimCons(tmp12199, Nil)

		tmp12201 := PrimCons(tmp12168, tmp12200)

		tmp12202 := PrimCons(tmp12166, tmp12201)

		__e.Return(PrimCons(symlet, tmp12202))
		return

	}, 3)

	tmp12203 := Call(__e, PrimNS2Value(symdef), symshen_4profile_1func, tmp12165)

	_ = tmp12203

	tmp12204 := MakeNative(func(__e *ControlFlow) {
		V4620 := __e.Get(1)
		_ = V4620
		tmp12205 := MakeNative(func(__e *ControlFlow) {
			Results := __e.Get(1)
			_ = Results
			tmp12206 := MakeNative(func(__e *ControlFlow) {
				Initialise := __e.Get(1)
				_ = Initialise
				__e.TailApply(PrimNS2Value(sym_8p), V4620, Results)
				return
			}, 1)

			tmp12207 := Call(__e, PrimNS2Value(symshen_4put_1profile), V4620, MakeNumber(0))

			__e.TailApply(tmp12206, tmp12207)
			return

		}, 1)

		tmp12208 := Call(__e, PrimNS2Value(symshen_4get_1profile), V4620)

		__e.TailApply(tmp12205, tmp12208)
		return

	}, 1)

	tmp12209 := Call(__e, PrimNS2Value(symdef), symprofile_1results, tmp12204)

	_ = tmp12209

	tmp12210 := MakeNative(func(__e *ControlFlow) {
		V4621 := __e.Get(1)
		_ = V4621
		tmp12211 := MakeNative(func(__e *ControlFlow) {
			tmp12212 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V4621, symprofile, tmp12212)
			return

		}, 0)

		tmp12213 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(MakeNumber(0))
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), tmp12211, tmp12213)
		return

	}, 1)

	tmp12214 := Call(__e, PrimNS2Value(symdef), symshen_4get_1profile, tmp12210)

	_ = tmp12214

	tmp12215 := MakeNative(func(__e *ControlFlow) {
		V4622 := __e.Get(1)
		_ = V4622
		V4623 := __e.Get(2)
		_ = V4623
		tmp12216 := PrimNS3Value(sym_dproperty_1vector_d)

		__e.TailApply(PrimNS2Value(symput), V4622, symprofile, V4623, tmp12216)
		return

	}, 2)

	__e.TailApply(PrimNS2Value(symdef), symshen_4put_1profile, tmp12215)
	return

}, 0)
