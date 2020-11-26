package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4dict Obj                // shen.dict
var __defun__shen_4dict_2 Obj              // shen.dict?
var __defun__shen_4dict_1capacity Obj      // shen.dict-capacity
var __defun__shen_4dict_1count Obj         // shen.dict-count
var __defun__shen_4dict_1count_1_6 Obj     // shen.dict-count->
var __defun__shen_4_5_1dict_1bucket Obj    // shen.<-dict-bucket
var __defun__shen_4dict_1bucket_1_6 Obj    // shen.dict-bucket->
var __defun__shen_4dict_1update_1count Obj // shen.dict-update-count
var __defun__shen_4dict_1_6 Obj            // shen.dict->
var __defun__shen_4_5_1dict Obj            // shen.<-dict
var __defun__shen_4dict_1rm Obj            // shen.dict-rm
var __defun__shen_4dict_1fold Obj          // shen.dict-fold
var __defun__shen_4dict_1fold_1h Obj       // shen.dict-fold-h
var __defun__shen_4bucket_1fold Obj        // shen.bucket-fold
var __defun__shen_4dict_1keys Obj          // shen.dict-keys
var __defun__shen_4dict_1values Obj        // shen.dict-values

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		reg5830 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__e.Return(reg5830)
		return
	}, 0))
	__defun__shen_4dict = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2312 := __args[0]
		_ = V2312
		reg5831 := MakeNumber(1)
		reg5832 := PrimLessThan(V2312, reg5831)
		if reg5832 == True {
			reg5833 := MakeString("invalid initial dict size: ")
			reg5834 := MakeString("")
			reg5835 := MakeSymbol("shen.s")
			reg5836 := Call(__e, __defun__shen_4app, V2312, reg5834, reg5835)
			reg5837 := PrimStringConcat(reg5833, reg5836)
			reg5838 := PrimSimpleError(reg5837)
			__e.Return(reg5838)
			return
		} else {
			reg5839 := MakeNumber(3)
			reg5840 := PrimNumberAdd(reg5839, V2312)
			reg5841 := PrimAbsvector(reg5840)
			D := reg5841
			_ = D
			reg5842 := MakeNumber(0)
			reg5843 := MakeSymbol("shen.dictionary")
			reg5844 := PrimVectorSet(D, reg5842, reg5843)
			Tag := reg5844
			_ = Tag
			reg5845 := MakeNumber(1)
			reg5846 := PrimVectorSet(D, reg5845, V2312)
			Capacity := reg5846
			_ = Capacity
			reg5847 := MakeNumber(2)
			reg5848 := MakeNumber(0)
			reg5849 := PrimVectorSet(D, reg5847, reg5848)
			Count := reg5849
			_ = Count
			reg5850 := MakeNumber(3)
			reg5851 := MakeNumber(2)
			reg5852 := PrimNumberAdd(reg5851, V2312)
			reg5853 := Nil
			reg5854 := Call(__e, __defun__shen_4fillvector, D, reg5850, reg5852, reg5853)
			Fill := reg5854
			_ = Fill
			__e.Return(D)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dict", value: __defun__shen_4dict})

	__defun__shen_4dict_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2314 := __args[0]
		_ = V2314
		reg5855 := PrimIsVector(V2314)
		if reg5855 == True {
			reg5856 := MakeNative(func(__e Evaluator, __args ...Obj) {
				reg5857 := MakeNumber(0)
				reg5858 := PrimVectorGet(V2314, reg5857)
				__e.Return(reg5858)
				return
			}, 0)
			reg5859 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				reg5860 := MakeSymbol("shen.not-dictionary")
				__e.Return(reg5860)
				return
			}, 1)
			reg5861 := Try(__e, reg5856).Catch(reg5859)
			reg5862 := MakeSymbol("shen.dictionary")
			reg5863 := PrimEqual(reg5861, reg5862)
			if reg5863 == True {
				reg5864 := True
				__e.Return(reg5864)
				return
			} else {
				reg5865 := False
				__e.Return(reg5865)
				return
			}
		} else {
			reg5866 := False
			__e.Return(reg5866)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dict?", value: __defun__shen_4dict_2})

	__defun__shen_4dict_1capacity = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2316 := __args[0]
		_ = V2316
		reg5867 := MakeNumber(1)
		reg5868 := PrimVectorGet(V2316, reg5867)
		__e.Return(reg5868)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dict-capacity", value: __defun__shen_4dict_1capacity})

	__defun__shen_4dict_1count = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2318 := __args[0]
		_ = V2318
		reg5869 := MakeNumber(2)
		reg5870 := PrimVectorGet(V2318, reg5869)
		__e.Return(reg5870)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dict-count", value: __defun__shen_4dict_1count})

	__defun__shen_4dict_1count_1_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2321 := __args[0]
		_ = V2321
		V2322 := __args[1]
		_ = V2322
		reg5871 := MakeNumber(2)
		reg5872 := PrimVectorSet(V2321, reg5871, V2322)
		__e.Return(reg5872)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.dict-count->", value: __defun__shen_4dict_1count_1_6})

	__defun__shen_4_5_1dict_1bucket = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2325 := __args[0]
		_ = V2325
		V2326 := __args[1]
		_ = V2326
		reg5873 := MakeNumber(3)
		reg5874 := PrimNumberAdd(reg5873, V2326)
		reg5875 := PrimVectorGet(V2325, reg5874)
		__e.Return(reg5875)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.<-dict-bucket", value: __defun__shen_4_5_1dict_1bucket})

	__defun__shen_4dict_1bucket_1_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2330 := __args[0]
		_ = V2330
		V2331 := __args[1]
		_ = V2331
		V2332 := __args[2]
		_ = V2332
		reg5876 := MakeNumber(3)
		reg5877 := PrimNumberAdd(reg5876, V2331)
		reg5878 := PrimVectorSet(V2330, reg5877, V2332)
		__e.Return(reg5878)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.dict-bucket->", value: __defun__shen_4dict_1bucket_1_6})

	__defun__shen_4dict_1update_1count = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2336 := __args[0]
		_ = V2336
		V2337 := __args[1]
		_ = V2337
		V2338 := __args[2]
		_ = V2338
		reg5879 := Call(__e, __defun__length, V2338)
		reg5880 := Call(__e, __defun__length, V2337)
		reg5881 := PrimNumberSubtract(reg5879, reg5880)
		Diff := reg5881
		_ = Diff
		reg5882 := Call(__e, __defun__shen_4dict_1count, V2336)
		reg5883 := PrimNumberAdd(Diff, reg5882)
		__e.TailApply(__defun__shen_4dict_1count_1_6, V2336, reg5883)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.dict-update-count", value: __defun__shen_4dict_1update_1count})

	__defun__shen_4dict_1_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2342 := __args[0]
		_ = V2342
		V2343 := __args[1]
		_ = V2343
		V2344 := __args[2]
		_ = V2344
		reg5885 := Call(__e, __defun__shen_4dict_1capacity, V2342)
		reg5886 := Call(__e, __defun__hash, V2343, reg5885)
		N := reg5886
		_ = N
		reg5887 := Call(__e, __defun__shen_4_5_1dict_1bucket, V2342, N)
		Bucket := reg5887
		_ = Bucket
		reg5888 := Call(__e, __defun__shen_4assoc_1set, V2343, V2344, Bucket)
		NewBucket := reg5888
		_ = NewBucket
		reg5889 := Call(__e, __defun__shen_4dict_1bucket_1_6, V2342, N, NewBucket)
		Change := reg5889
		_ = Change
		reg5890 := Call(__e, __defun__shen_4dict_1update_1count, V2342, Bucket, NewBucket)
		Count := reg5890
		_ = Count
		__e.Return(V2344)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.dict->", value: __defun__shen_4dict_1_6})

	__defun__shen_4_5_1dict = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2347 := __args[0]
		_ = V2347
		V2348 := __args[1]
		_ = V2348
		reg5891 := Call(__e, __defun__shen_4dict_1capacity, V2347)
		reg5892 := Call(__e, __defun__hash, V2348, reg5891)
		N := reg5892
		_ = N
		reg5893 := Call(__e, __defun__shen_4_5_1dict_1bucket, V2347, N)
		Bucket := reg5893
		_ = Bucket
		reg5894 := Call(__e, __defun__assoc, V2348, Bucket)
		Result := reg5894
		_ = Result
		reg5895 := Call(__e, __defun__empty_2, Result)
		if reg5895 == True {
			reg5896 := MakeString("value ")
			reg5897 := MakeString(" not found in dict\n")
			reg5898 := MakeSymbol("shen.a")
			reg5899 := Call(__e, __defun__shen_4app, V2348, reg5897, reg5898)
			reg5900 := PrimStringConcat(reg5896, reg5899)
			reg5901 := PrimSimpleError(reg5900)
			__e.Return(reg5901)
			return
		} else {
			reg5902 := PrimTail(Result)
			__e.Return(reg5902)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.<-dict", value: __defun__shen_4_5_1dict})

	__defun__shen_4dict_1rm = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2351 := __args[0]
		_ = V2351
		V2352 := __args[1]
		_ = V2352
		reg5903 := Call(__e, __defun__shen_4dict_1capacity, V2351)
		reg5904 := Call(__e, __defun__hash, V2352, reg5903)
		N := reg5904
		_ = N
		reg5905 := Call(__e, __defun__shen_4_5_1dict_1bucket, V2351, N)
		Bucket := reg5905
		_ = Bucket
		reg5906 := Call(__e, __defun__shen_4assoc_1rm, V2352, Bucket)
		NewBucket := reg5906
		_ = NewBucket
		reg5907 := Call(__e, __defun__shen_4dict_1bucket_1_6, V2351, N, NewBucket)
		Change := reg5907
		_ = Change
		reg5908 := Call(__e, __defun__shen_4dict_1update_1count, V2351, Bucket, NewBucket)
		Count := reg5908
		_ = Count
		__e.Return(V2352)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.dict-rm", value: __defun__shen_4dict_1rm})

	__defun__shen_4dict_1fold = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2356 := __args[0]
		_ = V2356
		V2357 := __args[1]
		_ = V2357
		V2358 := __args[2]
		_ = V2358
		reg5909 := Call(__e, __defun__shen_4dict_1capacity, V2357)
		Limit := reg5909
		_ = Limit
		reg5910 := MakeNumber(0)
		__e.TailApply(__defun__shen_4dict_1fold_1h, V2356, V2357, V2358, reg5910, Limit)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.dict-fold", value: __defun__shen_4dict_1fold})

	__defun__shen_4dict_1fold_1h = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2365 := __args[0]
		_ = V2365
		V2366 := __args[1]
		_ = V2366
		V2367 := __args[2]
		_ = V2367
		V2368 := __args[3]
		_ = V2368
		V2369 := __args[4]
		_ = V2369
		reg5912 := PrimEqual(V2369, V2368)
		if reg5912 == True {
			__e.Return(V2367)
			return
		} else {
			reg5913 := Call(__e, __defun__shen_4_5_1dict_1bucket, V2366, V2368)
			B := reg5913
			_ = B
			reg5914 := Call(__e, __defun__shen_4bucket_1fold, V2365, B, V2367)
			Acc := reg5914
			_ = Acc
			reg5915 := MakeNumber(1)
			reg5916 := PrimNumberAdd(reg5915, V2368)
			__e.TailApply(__defun__shen_4dict_1fold_1h, V2365, V2366, Acc, reg5916, V2369)
			return
		}
	}, 5)
	__initDefs = append(__initDefs, defType{name: "shen.dict-fold-h", value: __defun__shen_4dict_1fold_1h})

	__defun__shen_4bucket_1fold = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2373 := __args[0]
		_ = V2373
		V2374 := __args[1]
		_ = V2374
		V2375 := __args[2]
		_ = V2375
		reg5918 := Nil
		reg5919 := PrimEqual(reg5918, V2374)
		if reg5919 == True {
			__e.Return(V2375)
			return
		} else {
			reg5920 := PrimIsPair(V2374)
			var reg5927 Obj
			if reg5920 == True {
				reg5921 := PrimHead(V2374)
				reg5922 := PrimIsPair(reg5921)
				var reg5925 Obj
				if reg5922 == True {
					reg5923 := True
					reg5925 = reg5923
				} else {
					reg5924 := False
					reg5925 = reg5924
				}
				reg5927 = reg5925
			} else {
				reg5926 := False
				reg5927 = reg5926
			}
			if reg5927 == True {
				reg5928 := PrimHead(V2374)
				reg5929 := PrimHead(reg5928)
				reg5930 := PrimHead(V2374)
				reg5931 := PrimTail(reg5930)
				reg5932 := PrimTail(V2374)
				reg5933 := Call(__e, __defun__shen_4bucket_1fold, V2373, reg5932, V2375)
				__e.TailApply(V2373, reg5929, reg5931, reg5933)
				return
			} else {
				reg5935 := MakeSymbol("shen.bucket-fold")
				__e.TailApply(__defun__shen_4f__error, reg5935)
				return
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.bucket-fold", value: __defun__shen_4bucket_1fold})

	__defun__shen_4dict_1keys = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2377 := __args[0]
		_ = V2377
		reg5937 := MakeNative(func(__e Evaluator, __args ...Obj) {
			K := __args[0]
			_ = K
			reg5938 := MakeNative(func(__e Evaluator, __args ...Obj) {
				__ := __args[0]
				_ = __
				reg5939 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Acc := __args[0]
					_ = Acc
					reg5940 := PrimCons(K, Acc)
					__e.Return(reg5940)
					return
				}, 1)
				__e.Return(reg5939)
				return
			}, 1)
			__e.Return(reg5938)
			return
		}, 1)
		reg5941 := Nil
		__e.TailApply(__defun__shen_4dict_1fold, reg5937, V2377, reg5941)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dict-keys", value: __defun__shen_4dict_1keys})

	__defun__shen_4dict_1values = MakeNative(func(__e Evaluator, __args ...Obj) {
		V2379 := __args[0]
		_ = V2379
		reg5943 := MakeNative(func(__e Evaluator, __args ...Obj) {
			__ := __args[0]
			_ = __
			reg5944 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V := __args[0]
				_ = V
				reg5945 := MakeNative(func(__e Evaluator, __args ...Obj) {
					Acc := __args[0]
					_ = Acc
					reg5946 := PrimCons(V, Acc)
					__e.Return(reg5946)
					return
				}, 1)
				__e.Return(reg5945)
				return
			}, 1)
			__e.Return(reg5944)
			return
		}, 1)
		reg5947 := Nil
		__e.TailApply(__defun__shen_4dict_1fold, reg5943, V2379, reg5947)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.dict-values", value: __defun__shen_4dict_1values})

}
