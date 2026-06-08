package main

import . "github.com/tiancaiamao/shen-go/kl"

var DictMain = MakeNative(func(__e *ControlFlow) {
	tmp16822 := MakeNative(func(__e *ControlFlow) {
		V4138 := __e.Get(1)
		_ = V4138
		tmp16838 := PrimLessThan(V4138, MakeNumber(1))

		if True == tmp16838 {
			tmp16823 := Call(__e, PrimFunc(symshen_4app), V4138, MakeString(""), symshen_4s)

			tmp16824 := PrimStringConcat(MakeString("invalid initial dict size: "), tmp16823)

			__e.Return(PrimSimpleError(tmp16824))
			return

		} else {
			tmp16825 := MakeNative(func(__e *ControlFlow) {
				W4139 := __e.Get(1)
				_ = W4139
				tmp16826 := MakeNative(func(__e *ControlFlow) {
					W4140 := __e.Get(1)
					_ = W4140
					tmp16827 := MakeNative(func(__e *ControlFlow) {
						W4141 := __e.Get(1)
						_ = W4141
						tmp16828 := MakeNative(func(__e *ControlFlow) {
							W4142 := __e.Get(1)
							_ = W4142
							tmp16829 := MakeNative(func(__e *ControlFlow) {
								W4143 := __e.Get(1)
								_ = W4143
								__e.Return(W4139)
								return
							}, 1)

							tmp16830 := PrimNumberAdd(MakeNumber(2), V4138)

							tmp16831 := Call(__e, PrimFunc(symshen_4fillvector), W4139, MakeNumber(3), tmp16830, Nil)

							__e.TailApply(tmp16829, tmp16831)
							return

						}, 1)

						tmp16832 := PrimVectorSet(W4139, MakeNumber(2), MakeNumber(0))

						__e.TailApply(tmp16828, tmp16832)
						return

					}, 1)

					tmp16833 := PrimVectorSet(W4139, MakeNumber(1), V4138)

					__e.TailApply(tmp16827, tmp16833)
					return

				}, 1)

				tmp16834 := PrimVectorSet(W4139, MakeNumber(0), symshen_4dictionary)

				__e.TailApply(tmp16826, tmp16834)
				return

			}, 1)

			tmp16835 := PrimNumberAdd(MakeNumber(3), V4138)

			tmp16836 := PrimAbsvector(tmp16835)

			__e.TailApply(tmp16825, tmp16836)
			return

		}

	}, 1)

	tmp16839 := Call(__e, ns2_1set, symshen_4dict, tmp16822)

	_ = tmp16839

	tmp16840 := MakeNative(func(__e *ControlFlow) {
		V4144 := __e.Get(1)
		_ = V4144
		tmp16847 := PrimIsVector(V4144)

		if True == tmp16847 {
			tmp16842 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V4144, MakeNumber(0)))
				return
			}, 0)

			tmp16843 := MakeNative(func(__e *ControlFlow) {
				Z4145 := __e.Get(1)
				_ = Z4145
				__e.Return(symshen_4not_1dictionary)
				return
			}, 1)

			tmp16844 := Call(__e, try_1catch, tmp16842, tmp16843)

			tmp16845 := PrimEqual(tmp16844, symshen_4dictionary)

			if True == tmp16845 {
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

	tmp16848 := Call(__e, ns2_1set, symshen_4dict_2, tmp16840)

	_ = tmp16848

	tmp16849 := MakeNative(func(__e *ControlFlow) {
		V4146 := __e.Get(1)
		_ = V4146
		__e.Return(PrimVectorGet(V4146, MakeNumber(1)))
		return
	}, 1)

	tmp16850 := Call(__e, ns2_1set, symshen_4dict_1capacity, tmp16849)

	_ = tmp16850

	tmp16851 := MakeNative(func(__e *ControlFlow) {
		V4147 := __e.Get(1)
		_ = V4147
		__e.Return(PrimVectorGet(V4147, MakeNumber(2)))
		return
	}, 1)

	tmp16852 := Call(__e, ns2_1set, symshen_4dict_1count, tmp16851)

	_ = tmp16852

	tmp16853 := MakeNative(func(__e *ControlFlow) {
		V4148 := __e.Get(1)
		_ = V4148
		V4149 := __e.Get(2)
		_ = V4149
		__e.Return(PrimVectorSet(V4148, MakeNumber(2), V4149))
		return
	}, 2)

	tmp16854 := Call(__e, ns2_1set, symshen_4dict_1count_1_6, tmp16853)

	_ = tmp16854

	tmp16855 := MakeNative(func(__e *ControlFlow) {
		V4150 := __e.Get(1)
		_ = V4150
		V4151 := __e.Get(2)
		_ = V4151
		tmp16856 := PrimNumberAdd(MakeNumber(3), V4151)

		__e.Return(PrimVectorGet(V4150, tmp16856))
		return

	}, 2)

	tmp16857 := Call(__e, ns2_1set, symshen_4_5_1dict_1bucket, tmp16855)

	_ = tmp16857

	tmp16858 := MakeNative(func(__e *ControlFlow) {
		V4152 := __e.Get(1)
		_ = V4152
		V4153 := __e.Get(2)
		_ = V4153
		V4154 := __e.Get(3)
		_ = V4154
		tmp16859 := PrimNumberAdd(MakeNumber(3), V4153)

		__e.Return(PrimVectorSet(V4152, tmp16859, V4154))
		return

	}, 3)

	tmp16860 := Call(__e, ns2_1set, symshen_4dict_1bucket_1_6, tmp16858)

	_ = tmp16860

	tmp16861 := MakeNative(func(__e *ControlFlow) {
		V4155 := __e.Get(1)
		_ = V4155
		V4156 := __e.Get(2)
		_ = V4156
		V4157 := __e.Get(3)
		_ = V4157
		tmp16862 := MakeNative(func(__e *ControlFlow) {
			W4158 := __e.Get(1)
			_ = W4158
			tmp16863 := Call(__e, PrimFunc(symshen_4dict_1count), V4155)

			tmp16864 := PrimNumberAdd(W4158, tmp16863)

			__e.TailApply(PrimFunc(symshen_4dict_1count_1_6), V4155, tmp16864)
			return

		}, 1)

		tmp16865 := Call(__e, PrimFunc(symlength), V4157)

		tmp16866 := Call(__e, PrimFunc(symlength), V4156)

		tmp16867 := PrimNumberSubtract(tmp16865, tmp16866)

		__e.TailApply(tmp16862, tmp16867)
		return

	}, 3)

	tmp16868 := Call(__e, ns2_1set, symshen_4dict_1update_1count, tmp16861)

	_ = tmp16868

	tmp16869 := MakeNative(func(__e *ControlFlow) {
		V4159 := __e.Get(1)
		_ = V4159
		V4160 := __e.Get(2)
		_ = V4160
		V4161 := __e.Get(3)
		_ = V4161
		tmp16870 := MakeNative(func(__e *ControlFlow) {
			W4162 := __e.Get(1)
			_ = W4162
			tmp16871 := MakeNative(func(__e *ControlFlow) {
				W4163 := __e.Get(1)
				_ = W4163
				tmp16872 := MakeNative(func(__e *ControlFlow) {
					W4164 := __e.Get(1)
					_ = W4164
					tmp16873 := MakeNative(func(__e *ControlFlow) {
						W4165 := __e.Get(1)
						_ = W4165
						tmp16874 := MakeNative(func(__e *ControlFlow) {
							W4166 := __e.Get(1)
							_ = W4166
							__e.Return(V4161)
							return
						}, 1)

						tmp16875 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4159, W4163, W4164)

						__e.TailApply(tmp16874, tmp16875)
						return

					}, 1)

					tmp16876 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4159, W4162, W4164)

					__e.TailApply(tmp16873, tmp16876)
					return

				}, 1)

				tmp16877 := Call(__e, PrimFunc(symshen_4assoc_1set), V4160, V4161, W4163)

				__e.TailApply(tmp16872, tmp16877)
				return

			}, 1)

			tmp16878 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4159, W4162)

			__e.TailApply(tmp16871, tmp16878)
			return

		}, 1)

		tmp16879 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4159)

		tmp16880 := Call(__e, PrimFunc(symhash), V4160, tmp16879)

		__e.TailApply(tmp16870, tmp16880)
		return

	}, 3)

	tmp16881 := Call(__e, ns2_1set, symshen_4dict_1_6, tmp16869)

	_ = tmp16881

	tmp16882 := MakeNative(func(__e *ControlFlow) {
		V4167 := __e.Get(1)
		_ = V4167
		V4168 := __e.Get(2)
		_ = V4168
		tmp16883 := MakeNative(func(__e *ControlFlow) {
			W4169 := __e.Get(1)
			_ = W4169
			tmp16884 := MakeNative(func(__e *ControlFlow) {
				W4170 := __e.Get(1)
				_ = W4170
				tmp16885 := MakeNative(func(__e *ControlFlow) {
					W4171 := __e.Get(1)
					_ = W4171
					tmp16889 := Call(__e, PrimFunc(symempty_2), W4171)

					if True == tmp16889 {
						tmp16886 := Call(__e, PrimFunc(symshen_4app), V4168, MakeString(" not found in dict\n"), symshen_4a)

						tmp16887 := PrimStringConcat(MakeString("value "), tmp16886)

						__e.Return(PrimSimpleError(tmp16887))
						return

					} else {
						__e.Return(PrimTail(W4171))
						return
					}

				}, 1)

				tmp16890 := Call(__e, PrimFunc(symassoc), V4168, W4170)

				__e.TailApply(tmp16885, tmp16890)
				return

			}, 1)

			tmp16891 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4167, W4169)

			__e.TailApply(tmp16884, tmp16891)
			return

		}, 1)

		tmp16892 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4167)

		tmp16893 := Call(__e, PrimFunc(symhash), V4168, tmp16892)

		__e.TailApply(tmp16883, tmp16893)
		return

	}, 2)

	tmp16894 := Call(__e, ns2_1set, symshen_4_5_1dict, tmp16882)

	_ = tmp16894

	tmp16895 := MakeNative(func(__e *ControlFlow) {
		V4172 := __e.Get(1)
		_ = V4172
		V4173 := __e.Get(2)
		_ = V4173
		tmp16896 := MakeNative(func(__e *ControlFlow) {
			W4174 := __e.Get(1)
			_ = W4174
			tmp16897 := MakeNative(func(__e *ControlFlow) {
				W4175 := __e.Get(1)
				_ = W4175
				tmp16898 := MakeNative(func(__e *ControlFlow) {
					W4176 := __e.Get(1)
					_ = W4176
					tmp16899 := MakeNative(func(__e *ControlFlow) {
						W4177 := __e.Get(1)
						_ = W4177
						tmp16900 := MakeNative(func(__e *ControlFlow) {
							W4178 := __e.Get(1)
							_ = W4178
							__e.Return(V4173)
							return
						}, 1)

						tmp16901 := Call(__e, PrimFunc(symshen_4dict_1update_1count), V4172, W4175, W4176)

						__e.TailApply(tmp16900, tmp16901)
						return

					}, 1)

					tmp16902 := Call(__e, PrimFunc(symshen_4dict_1bucket_1_6), V4172, W4174, W4176)

					__e.TailApply(tmp16899, tmp16902)
					return

				}, 1)

				tmp16903 := Call(__e, PrimFunc(symshen_4assoc_1rm), V4173, W4175)

				__e.TailApply(tmp16898, tmp16903)
				return

			}, 1)

			tmp16904 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4172, W4174)

			__e.TailApply(tmp16897, tmp16904)
			return

		}, 1)

		tmp16905 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4172)

		tmp16906 := Call(__e, PrimFunc(symhash), V4173, tmp16905)

		__e.TailApply(tmp16896, tmp16906)
		return

	}, 2)

	tmp16907 := Call(__e, ns2_1set, symshen_4dict_1rm, tmp16895)

	_ = tmp16907

	tmp16908 := MakeNative(func(__e *ControlFlow) {
		V4179 := __e.Get(1)
		_ = V4179
		V4180 := __e.Get(2)
		_ = V4180
		V4181 := __e.Get(3)
		_ = V4181
		tmp16909 := MakeNative(func(__e *ControlFlow) {
			W4182 := __e.Get(1)
			_ = W4182
			__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4179, V4180, V4181, MakeNumber(0), W4182)
			return
		}, 1)

		tmp16910 := Call(__e, PrimFunc(symshen_4dict_1capacity), V4180)

		__e.TailApply(tmp16909, tmp16910)
		return

	}, 3)

	tmp16911 := Call(__e, ns2_1set, symshen_4dict_1fold, tmp16908)

	_ = tmp16911

	tmp16912 := MakeNative(func(__e *ControlFlow) {
		V4184 := __e.Get(1)
		_ = V4184
		V4185 := __e.Get(2)
		_ = V4185
		V4186 := __e.Get(3)
		_ = V4186
		V4187 := __e.Get(4)
		_ = V4187
		V4188 := __e.Get(5)
		_ = V4188
		tmp16919 := PrimEqual(V4187, V4188)

		if True == tmp16919 {
			__e.Return(V4186)
			return
		} else {
			tmp16913 := MakeNative(func(__e *ControlFlow) {
				W4189 := __e.Get(1)
				_ = W4189
				tmp16914 := MakeNative(func(__e *ControlFlow) {
					W4190 := __e.Get(1)
					_ = W4190
					tmp16915 := PrimNumberAdd(MakeNumber(1), V4187)

					__e.TailApply(PrimFunc(symshen_4dict_1fold_1h), V4184, V4185, W4190, tmp16915, V4188)
					return

				}, 1)

				tmp16916 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4184, W4189, V4186)

				__e.TailApply(tmp16914, tmp16916)
				return

			}, 1)

			tmp16917 := Call(__e, PrimFunc(symshen_4_5_1dict_1bucket), V4185, V4187)

			__e.TailApply(tmp16913, tmp16917)
			return

		}

	}, 5)

	tmp16920 := Call(__e, ns2_1set, symshen_4dict_1fold_1h, tmp16912)

	_ = tmp16920

	tmp16921 := MakeNative(func(__e *ControlFlow) {
		V4191 := __e.Get(1)
		_ = V4191
		V4192 := __e.Get(2)
		_ = V4192
		V4193 := __e.Get(3)
		_ = V4193
		tmp16937 := PrimEqual(Nil, V4192)

		if True == tmp16937 {
			__e.Return(V4193)
			return
		} else {
			tmp16935 := PrimIsPair(V4192)

			var ifres16931 Obj

			if True == tmp16935 {
				tmp16933 := PrimHead(V4192)

				tmp16934 := PrimIsPair(tmp16933)

				var ifres16932 Obj

				if True == tmp16934 {
					ifres16932 = True

				} else {
					ifres16932 = False

				}

				ifres16931 = ifres16932

			} else {
				ifres16931 = False

			}

			if True == ifres16931 {
				tmp16922 := PrimHead(V4192)

				tmp16923 := PrimHead(tmp16922)

				tmp16924 := Call(__e, V4191, tmp16923)

				tmp16925 := PrimHead(V4192)

				tmp16926 := PrimTail(tmp16925)

				tmp16927 := Call(__e, tmp16924, tmp16926)

				tmp16928 := PrimTail(V4192)

				tmp16929 := Call(__e, PrimFunc(symshen_4bucket_1fold), V4191, tmp16928, V4193)

				__e.TailApply(tmp16927, tmp16929)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bucket_1fold)
				return
			}

		}

	}, 3)

	tmp16938 := Call(__e, ns2_1set, symshen_4bucket_1fold, tmp16921)

	_ = tmp16938

	tmp16939 := MakeNative(func(__e *ControlFlow) {
		V4194 := __e.Get(1)
		_ = V4194
		tmp16940 := MakeNative(func(__e *ControlFlow) {
			Z4195 := __e.Get(1)
			_ = Z4195
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Z4196 := __e.Get(1)
				_ = Z4196
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Z4197 := __e.Get(1)
					_ = Z4197
					__e.Return(PrimCons(Z4195, Z4197))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp16940, V4194, Nil)
		return

	}, 1)

	tmp16941 := Call(__e, ns2_1set, symshen_4dict_1keys, tmp16939)

	_ = tmp16941

	tmp16942 := MakeNative(func(__e *ControlFlow) {
		V4198 := __e.Get(1)
		_ = V4198
		tmp16943 := MakeNative(func(__e *ControlFlow) {
			Z4199 := __e.Get(1)
			_ = Z4199
			__e.Return(MakeNative(func(__e *ControlFlow) {
				Z4200 := __e.Get(1)
				_ = Z4200
				__e.Return(MakeNative(func(__e *ControlFlow) {
					Z4201 := __e.Get(1)
					_ = Z4201
					__e.Return(PrimCons(Z4200, Z4201))
					return
				}, 1))
				return
			}, 1))
			return
		}, 1)

		__e.TailApply(PrimFunc(symshen_4dict_1fold), tmp16943, V4198, Nil)
		return

	}, 1)

	__e.TailApply(ns2_1set, symshen_4dict_1values, tmp16942)
	return

}, 0)
