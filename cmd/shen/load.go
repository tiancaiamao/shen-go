package main

import . "github.com/tiancaiamao/shen-go/cora"

var LoadMain = MakeNative(func(__e *ControlFlow) {
	_ = MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

	tmp7851 := MakeNative(func(__e *ControlFlow) {
		V701 := __e.Get(1)
		_ = V701
		tmp7852 := MakeNative(func(__e *ControlFlow) {
			Load := __e.Get(1)
			_ = Load
			tmp7853 := MakeNative(func(__e *ControlFlow) {
				Infs := __e.Get(1)
				_ = Infs
				__e.Return(symloaded)
				return
			}, 1)

			tmp7860 := PrimNS3Value(symshen_4_dtc_d)

			var ifres7854 Obj

			if True == tmp7860 {
				tmp7855 := Call(__e, PrimNS2Value(syminferences))

				tmp7856 := Call(__e, PrimNS2Value(symshen_4app), tmp7855, MakeString(" inferences\n"), symshen_4a)

				tmp7857 := PrimStringConcat(MakeString("\ntypechecked in "), tmp7856)

				tmp7858 := Call(__e, PrimNS2Value(symstoutput))

				tmp7859 := Call(__e, PrimNS2Value(symshen_4prhush), tmp7857, tmp7858)

				ifres7854 = tmp7859

			} else {
				ifres7854 = symshen_4skip

			}

			__e.TailApply(tmp7853, ifres7854)
			return

		}, 1)

		tmp7861 := MakeNative(func(__e *ControlFlow) {
			Start := __e.Get(1)
			_ = Start
			tmp7862 := MakeNative(func(__e *ControlFlow) {
				Result := __e.Get(1)
				_ = Result
				tmp7863 := MakeNative(func(__e *ControlFlow) {
					Finish := __e.Get(1)
					_ = Finish
					tmp7864 := MakeNative(func(__e *ControlFlow) {
						Time := __e.Get(1)
						_ = Time
						tmp7865 := MakeNative(func(__e *ControlFlow) {
							Message := __e.Get(1)
							_ = Message
							__e.Return(Result)
							return
						}, 1)

						tmp7866 := PrimStr(Time)

						tmp7867 := PrimStringConcat(tmp7866, MakeString(" secs\n"))

						tmp7868 := PrimStringConcat(MakeString("\nrun time: "), tmp7867)

						tmp7869 := Call(__e, PrimNS2Value(symstoutput))

						tmp7870 := Call(__e, PrimNS2Value(symshen_4prhush), tmp7868, tmp7869)

						__e.TailApply(tmp7865, tmp7870)
						return

					}, 1)

					tmp7871 := PrimNumberSubtract(Finish, Start)

					__e.TailApply(tmp7864, tmp7871)
					return

				}, 1)

				tmp7872 := PrimGetTime(symrun)

				__e.TailApply(tmp7863, tmp7872)
				return

			}, 1)

			tmp7873 := PrimNS3Value(symshen_4_dtc_d)

			tmp7874 := Call(__e, PrimNS2Value(symread_1file), V701)

			tmp7875 := Call(__e, PrimNS2Value(symshen_4load_1help), tmp7873, tmp7874)

			__e.TailApply(tmp7862, tmp7875)
			return

		}, 1)

		tmp7876 := PrimGetTime(symrun)

		tmp7877 := Call(__e, tmp7861, tmp7876)

		__e.TailApply(tmp7852, tmp7877)
		return

	}, 1)

	tmp7878 := Call(__e, PrimNS1Value(symns2_1set), symload, tmp7851)

	_ = tmp7878

	tmp7879 := MakeNative(func(__e *ControlFlow) {
		V708 := __e.Get(1)
		_ = V708
		V709 := __e.Get(2)
		_ = V709
		tmp7897 := PrimEqual(False, V708)

		if True == tmp7897 {
			tmp7893 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				tmp7894 := Call(__e, PrimNS2Value(symshen_4eval_1without_1macros), X)

				tmp7895 := Call(__e, PrimNS2Value(symshen_4app), tmp7894, MakeString("\n"), symshen_4s)

				tmp7896 := Call(__e, PrimNS2Value(symstoutput))

				__e.TailApply(PrimNS2Value(symshen_4prhush), tmp7895, tmp7896)
				return

			}, 1)

			__e.TailApply(PrimNS2Value(symshen_4for_1each), tmp7893, V709)
			return

		} else {
			tmp7881 := MakeNative(func(__e *ControlFlow) {
				RemoveSynonyms := __e.Get(1)
				_ = RemoveSynonyms
				tmp7882 := MakeNative(func(__e *ControlFlow) {
					Table := __e.Get(1)
					_ = Table
					tmp7883 := MakeNative(func(__e *ControlFlow) {
						Assume := __e.Get(1)
						_ = Assume
						tmp7884 := MakeNative(func(__e *ControlFlow) {
							tmp7885 := MakeNative(func(__e *ControlFlow) {
								X := __e.Get(1)
								_ = X
								__e.TailApply(PrimNS2Value(symshen_4typecheck_1and_1load), X)
								return
							}, 1)

							__e.TailApply(PrimNS2Value(symshen_4for_1each), tmp7885, RemoveSynonyms)
							return

						}, 0)

						tmp7886 := MakeNative(func(__e *ControlFlow) {
							E := __e.Get(1)
							_ = E
							__e.TailApply(PrimNS2Value(symshen_4unwind_1types), E, Table)
							return
						}, 1)

						__e.TailApply(PrimNS1Value(symtry_1catch), tmp7884, tmp7886)
						return

					}, 1)

					tmp7887 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4assumetype), X)
						return
					}, 1)

					tmp7888 := Call(__e, PrimNS2Value(symshen_4for_1each), tmp7887, Table)

					__e.TailApply(tmp7883, tmp7888)
					return

				}, 1)

				tmp7889 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimNS2Value(symshen_4typetable), X)
					return
				}, 1)

				tmp7890 := Call(__e, PrimNS2Value(symmapcan), tmp7889, RemoveSynonyms)

				__e.TailApply(tmp7882, tmp7890)
				return

			}, 1)

			tmp7891 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4remove_1synonyms), X)
				return
			}, 1)

			tmp7892 := Call(__e, PrimNS2Value(symmapcan), tmp7891, V709)

			__e.TailApply(tmp7881, tmp7892)
			return

		}

	}, 2)

	tmp7898 := Call(__e, PrimNS1Value(symns2_1set), symshen_4load_1help, tmp7879)

	_ = tmp7898

	tmp7899 := MakeNative(func(__e *ControlFlow) {
		V711 := __e.Get(1)
		_ = V711
		tmp7906 := PrimIsPair(V711)

		var ifres7902 Obj

		if True == tmp7906 {
			tmp7904 := PrimHead(V711)

			tmp7905 := PrimEqual(symshen_4synonyms_1help, tmp7904)

			var ifres7903 Obj

			if True == tmp7905 {
				ifres7903 = True

			} else {
				ifres7903 = False

			}

			ifres7902 = ifres7903

		} else {
			ifres7902 = False

		}

		if True == ifres7902 {
			tmp7901 := Call(__e, PrimNS2Value(symeval), V711)

			_ = tmp7901

			__e.Return(Nil)
			return

		} else {
			__e.Return(PrimCons(V711, Nil))
			return
		}

	}, 1)

	tmp7907 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remove_1synonyms, tmp7899)

	_ = tmp7907

	tmp7908 := MakeNative(func(__e *ControlFlow) {
		V713 := __e.Get(1)
		_ = V713
		tmp7909 := Call(__e, PrimNS2Value(symnl), MakeNumber(1))

		_ = tmp7909

		tmp7910 := Call(__e, PrimNS2Value(symgensym), symA)

		__e.TailApply(PrimNS2Value(symshen_4typecheck_1and_1evaluate), V713, tmp7910)
		return

	}, 1)

	tmp7911 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typecheck_1and_1load, tmp7908)

	_ = tmp7911

	tmp7912 := MakeNative(func(__e *ControlFlow) {
		V719 := __e.Get(1)
		_ = V719
		tmp7934 := PrimIsPair(V719)

		var ifres7926 Obj

		if True == tmp7934 {
			tmp7932 := PrimHead(V719)

			tmp7933 := PrimEqual(symdefine, tmp7932)

			var ifres7928 Obj

			if True == tmp7933 {
				tmp7930 := PrimTail(V719)

				tmp7931 := PrimIsPair(tmp7930)

				var ifres7929 Obj

				if True == tmp7931 {
					ifres7929 = True

				} else {
					ifres7929 = False

				}

				ifres7928 = ifres7929

			} else {
				ifres7928 = False

			}

			var ifres7927 Obj

			if True == ifres7928 {
				ifres7927 = True

			} else {
				ifres7927 = False

			}

			ifres7926 = ifres7927

		} else {
			ifres7926 = False

		}

		if True == ifres7926 {
			tmp7914 := MakeNative(func(__e *ControlFlow) {
				Sig := __e.Get(1)
				_ = Sig
				tmp7915 := PrimTail(V719)

				tmp7916 := PrimHead(tmp7915)

				tmp7917 := PrimCons(tmp7916, Sig)

				__e.Return(PrimCons(tmp7917, Nil))
				return

			}, 1)

			tmp7918 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(PrimNS2Value(symshen_4_5sig_7rest_6), Y)
				return
			}, 1)

			tmp7919 := PrimTail(V719)

			tmp7920 := PrimTail(tmp7919)

			tmp7921 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				tmp7922 := PrimTail(V719)

				tmp7923 := PrimHead(tmp7922)

				tmp7924 := Call(__e, PrimNS2Value(symshen_4app), tmp7923, MakeString(" lacks a proper signature.\n"), symshen_4a)

				__e.Return(PrimSimpleError(tmp7924))
				return

			}, 1)

			tmp7925 := Call(__e, PrimNS2Value(symcompile), tmp7918, tmp7920, tmp7921)

			__e.TailApply(tmp7914, tmp7925)
			return

		} else {
			__e.Return(Nil)
			return
		}

	}, 1)

	tmp7935 := Call(__e, PrimNS1Value(symns2_1set), symshen_4typetable, tmp7912)

	_ = tmp7935

	tmp7936 := MakeNative(func(__e *ControlFlow) {
		V721 := __e.Get(1)
		_ = V721
		tmp7940 := PrimIsPair(V721)

		if True == tmp7940 {
			tmp7938 := PrimHead(V721)

			tmp7939 := PrimTail(V721)

			__e.TailApply(PrimNS2Value(symdeclare), tmp7938, tmp7939)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4assumetype)
			return
		}

	}, 1)

	tmp7941 := Call(__e, PrimNS1Value(symns2_1set), symshen_4assumetype, tmp7936)

	_ = tmp7941

	tmp7942 := MakeNative(func(__e *ControlFlow) {
		V728 := __e.Get(1)
		_ = V728
		V729 := __e.Get(2)
		_ = V729
		tmp7955 := PrimEqual(Nil, V729)

		if True == tmp7955 {
			tmp7954 := PrimErrorToString(V728)

			__e.Return(PrimSimpleError(tmp7954))
			return

		} else {
			tmp7953 := PrimIsPair(V729)

			var ifres7949 Obj

			if True == tmp7953 {
				tmp7951 := PrimHead(V729)

				tmp7952 := PrimIsPair(tmp7951)

				var ifres7950 Obj

				if True == tmp7952 {
					ifres7950 = True

				} else {
					ifres7950 = False

				}

				ifres7949 = ifres7950

			} else {
				ifres7949 = False

			}

			if True == ifres7949 {
				tmp7945 := PrimHead(V729)

				tmp7946 := PrimHead(tmp7945)

				tmp7947 := Call(__e, PrimNS2Value(symshen_4remtype), tmp7946)

				_ = tmp7947

				tmp7948 := PrimTail(V729)

				__e.TailApply(PrimNS2Value(symshen_4unwind_1types), V728, tmp7948)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4unwind_1types)
				return
			}

		}

	}, 2)

	tmp7956 := Call(__e, PrimNS1Value(symns2_1set), symshen_4unwind_1types, tmp7942)

	_ = tmp7956

	tmp7957 := MakeNative(func(__e *ControlFlow) {
		V731 := __e.Get(1)
		_ = V731
		tmp7958 := PrimNS3Value(symshen_4_dsignedfuncs_d)

		tmp7959 := Call(__e, PrimNS2Value(symshen_4removetype), V731, tmp7958)

		__e.Return(PrimNS3Set(symshen_4_dsignedfuncs_d, tmp7959))
		return

	}, 1)

	tmp7960 := Call(__e, PrimNS1Value(symns2_1set), symshen_4remtype, tmp7957)

	_ = tmp7960

	tmp7961 := MakeNative(func(__e *ControlFlow) {
		V739 := __e.Get(1)
		_ = V739
		V740 := __e.Get(2)
		_ = V740
		tmp7982 := PrimEqual(Nil, V740)

		if True == tmp7982 {
			__e.Return(Nil)
			return
		} else {
			tmp7981 := PrimIsPair(V740)

			var ifres7972 Obj

			if True == tmp7981 {
				tmp7979 := PrimHead(V740)

				tmp7980 := PrimIsPair(tmp7979)

				var ifres7974 Obj

				if True == tmp7980 {
					tmp7976 := PrimHead(V740)

					tmp7977 := PrimHead(tmp7976)

					tmp7978 := PrimEqual(tmp7977, V739)

					var ifres7975 Obj

					if True == tmp7978 {
						ifres7975 = True

					} else {
						ifres7975 = False

					}

					ifres7974 = ifres7975

				} else {
					ifres7974 = False

				}

				var ifres7973 Obj

				if True == ifres7974 {
					ifres7973 = True

				} else {
					ifres7973 = False

				}

				ifres7972 = ifres7973

			} else {
				ifres7972 = False

			}

			if True == ifres7972 {
				tmp7969 := PrimHead(V740)

				tmp7970 := PrimHead(tmp7969)

				tmp7971 := PrimTail(V740)

				__e.TailApply(PrimNS2Value(symshen_4removetype), tmp7970, tmp7971)
				return

			} else {
				tmp7968 := PrimIsPair(V740)

				if True == tmp7968 {
					tmp7965 := PrimHead(V740)

					tmp7966 := PrimTail(V740)

					tmp7967 := Call(__e, PrimNS2Value(symshen_4removetype), V739, tmp7966)

					__e.Return(PrimCons(tmp7965, tmp7967))
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f__error), symshen_4removetype)
					return
				}

			}

		}

	}, 2)

	tmp7983 := Call(__e, PrimNS1Value(symns2_1set), symshen_4removetype, tmp7961)

	_ = tmp7983

	tmp7984 := MakeNative(func(__e *ControlFlow) {
		V742 := __e.Get(1)
		_ = V742
		tmp7985 := MakeNative(func(__e *ControlFlow) {
			Parse__shen_4_5signature_6 := __e.Get(1)
			_ = Parse__shen_4_5signature_6
			tmp7995 := Call(__e, PrimNS2Value(symfail))

			tmp7996 := PrimEqual(tmp7995, Parse__shen_4_5signature_6)

			tmp7997 := PrimNot(tmp7996)

			if True == tmp7997 {
				tmp7987 := MakeNative(func(__e *ControlFlow) {
					Parse___5_b_6 := __e.Get(1)
					_ = Parse___5_b_6
					tmp7991 := Call(__e, PrimNS2Value(symfail))

					tmp7992 := PrimEqual(tmp7991, Parse___5_b_6)

					tmp7993 := PrimNot(tmp7992)

					if True == tmp7993 {
						tmp7989 := PrimHead(Parse___5_b_6)

						tmp7990 := Call(__e, PrimNS2Value(symshen_4hdtl), Parse__shen_4_5signature_6)

						__e.TailApply(PrimNS2Value(symshen_4pair), tmp7989, tmp7990)
						return

					} else {
						__e.TailApply(PrimNS2Value(symfail))
						return
					}

				}, 1)

				tmp7994 := Call(__e, PrimNS2Value(sym_5_b_6), Parse__shen_4_5signature_6)

				__e.TailApply(tmp7987, tmp7994)
				return

			} else {
				__e.TailApply(PrimNS2Value(symfail))
				return
			}

		}, 1)

		tmp7998 := Call(__e, PrimNS2Value(symshen_4_5signature_6), V742)

		__e.TailApply(tmp7985, tmp7998)
		return

	}, 1)

	tmp7999 := Call(__e, PrimNS1Value(symns2_1set), symshen_4_5sig_7rest_6, tmp7984)

	_ = tmp7999

	tmp8000 := MakeNative(func(__e *ControlFlow) {
		V745 := __e.Get(1)
		_ = V745
		V746 := __e.Get(2)
		_ = V746
		tmp8001 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp8002 := MakeNative(func(__e *ControlFlow) {
				String := __e.Get(1)
				_ = String
				tmp8003 := MakeNative(func(__e *ControlFlow) {
					Write := __e.Get(1)
					_ = Write
					tmp8004 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						__e.Return(V746)
						return
					}, 1)

					tmp8005 := PrimCloseStream(Stream)

					__e.TailApply(tmp8004, tmp8005)
					return

				}, 1)

				tmp8006 := Call(__e, PrimNS2Value(sympr), String, Stream)

				__e.TailApply(tmp8003, tmp8006)
				return

			}, 1)

			tmp8010 := PrimIsString(V746)

			var ifres8007 Obj

			if True == tmp8010 {
				tmp8009 := Call(__e, PrimNS2Value(symshen_4app), V746, MakeString("\n\n"), symshen_4a)

				ifres8007 = tmp8009

			} else {
				tmp8008 := Call(__e, PrimNS2Value(symshen_4app), V746, MakeString("\n\n"), symshen_4s)

				ifres8007 = tmp8008

			}

			__e.TailApply(tmp8002, ifres8007)
			return

		}, 1)

		tmp8011 := PrimOpenStream(V745, symout)

		__e.TailApply(tmp8001, tmp8011)
		return

	}, 2)

	__e.TailApply(PrimNS1Value(symns2_1set), symwrite_1to_1file, tmp8000)
	return

}, 0)
