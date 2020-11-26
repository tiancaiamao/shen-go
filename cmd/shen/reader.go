package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4read_1char_1code Obj            // shen.read-char-code
var __defun__read_1file_1as_1bytelist Obj          // read-file-as-bytelist
var __defun__shen_4read_1file_1as_1charlist Obj    // shen.read-file-as-charlist
var __defun__shen_4read_1file_1as_1Xlist Obj       // shen.read-file-as-Xlist
var __defun__shen_4read_1file_1as_1Xlist_1help Obj // shen.read-file-as-Xlist-help
var __defun__read_1file_1as_1string Obj            // read-file-as-string
var __defun__shen_4rfas_1h Obj                     // shen.rfas-h
var __defun__input Obj                             // input
var __defun__input_7 Obj                           // input+
var __defun__shen_4monotype Obj                    // shen.monotype
var __defun__read Obj                              // read
var __defun__it Obj                                // it
var __defun__shen_4read_1loop Obj                  // shen.read-loop
var __defun__shen_4terminator_2 Obj                // shen.terminator?
var __defun__lineread Obj                          // lineread
var __defun__shen_4lineread_1loop Obj              // shen.lineread-loop
var __defun__shen_4record_1it Obj                  // shen.record-it
var __defun__shen_4trim_1whitespace Obj            // shen.trim-whitespace
var __defun__shen_4record_1it_1h Obj               // shen.record-it-h
var __defun__shen_4cn_1all Obj                     // shen.cn-all
var __defun__read_1file Obj                        // read-file
var __defun__read_1from_1string Obj                // read-from-string
var __defun__shen_4read_1error Obj                 // shen.read-error
var __defun__shen_4compress_150 Obj                // shen.compress-50
var __defun__shen_4_5st__input_6 Obj               // shen.<st_input>
var __defun__shen_4_5lsb_6 Obj                     // shen.<lsb>
var __defun__shen_4_5rsb_6 Obj                     // shen.<rsb>
var __defun__shen_4_5lcurly_6 Obj                  // shen.<lcurly>
var __defun__shen_4_5rcurly_6 Obj                  // shen.<rcurly>
var __defun__shen_4_5bar_6 Obj                     // shen.<bar>
var __defun__shen_4_5semicolon_6 Obj               // shen.<semicolon>
var __defun__shen_4_5colon_6 Obj                   // shen.<colon>
var __defun__shen_4_5comma_6 Obj                   // shen.<comma>
var __defun__shen_4_5equal_6 Obj                   // shen.<equal>
var __defun__shen_4_5minus_6 Obj                   // shen.<minus>
var __defun__shen_4_5lrb_6 Obj                     // shen.<lrb>
var __defun__shen_4_5rrb_6 Obj                     // shen.<rrb>
var __defun__shen_4_5atom_6 Obj                    // shen.<atom>
var __defun__shen_4control_1chars Obj              // shen.control-chars
var __defun__shen_4code_1point Obj                 // shen.code-point
var __defun__shen_4after_1codepoint Obj            // shen.after-codepoint
var __defun__shen_4decimalise Obj                  // shen.decimalise
var __defun__shen_4digits_1_6integers Obj          // shen.digits->integers
var __defun__shen_4_5sym_6 Obj                     // shen.<sym>
var __defun__shen_4_5alphanums_6 Obj               // shen.<alphanums>
var __defun__shen_4_5alphanum_6 Obj                // shen.<alphanum>
var __defun__shen_4_5num_6 Obj                     // shen.<num>
var __defun__shen_4numbyte_2 Obj                   // shen.numbyte?
var __defun__shen_4_5alpha_6 Obj                   // shen.<alpha>
var __defun__shen_4symbol_1code_2 Obj              // shen.symbol-code?
var __defun__shen_4_5str_6 Obj                     // shen.<str>
var __defun__shen_4_5dbq_6 Obj                     // shen.<dbq>
var __defun__shen_4_5strcontents_6 Obj             // shen.<strcontents>
var __defun__shen_4_5byte_6 Obj                    // shen.<byte>
var __defun__shen_4_5strc_6 Obj                    // shen.<strc>
var __defun__shen_4_5number_6 Obj                  // shen.<number>
var __defun__shen_4_5E_6 Obj                       // shen.<E>
var __defun__shen_4_5log10_6 Obj                   // shen.<log10>
var __defun__shen_4_5plus_6 Obj                    // shen.<plus>
var __defun__shen_4_5stop_6 Obj                    // shen.<stop>
var __defun__shen_4_5predigits_6 Obj               // shen.<predigits>
var __defun__shen_4_5postdigits_6 Obj              // shen.<postdigits>
var __defun__shen_4_5digits_6 Obj                  // shen.<digits>
var __defun__shen_4_5digit_6 Obj                   // shen.<digit>
var __defun__shen_4byte_1_6digit Obj               // shen.byte->digit
var __defun__shen_4pre Obj                         // shen.pre
var __defun__shen_4post Obj                        // shen.post
var __defun__shen_4expt Obj                        // shen.expt
var __defun__shen_4_5st__input1_6 Obj              // shen.<st_input1>
var __defun__shen_4_5st__input2_6 Obj              // shen.<st_input2>
var __defun__shen_4_5comment_6 Obj                 // shen.<comment>
var __defun__shen_4_5singleline_6 Obj              // shen.<singleline>
var __defun__shen_4_5backslash_6 Obj               // shen.<backslash>
var __defun__shen_4_5anysingle_6 Obj               // shen.<anysingle>
var __defun__shen_4_5non_1return_6 Obj             // shen.<non-return>
var __defun__shen_4_5return_6 Obj                  // shen.<return>
var __defun__shen_4_5multiline_6 Obj               // shen.<multiline>
var __defun__shen_4_5times_6 Obj                   // shen.<times>
var __defun__shen_4_5anymulti_6 Obj                // shen.<anymulti>
var __defun__shen_4_5whitespaces_6 Obj             // shen.<whitespaces>
var __defun__shen_4_5whitespace_6 Obj              // shen.<whitespace>
var __defun__shen_4cons__form Obj                  // shen.cons_form
var __defun__shen_4package_1macro Obj              // shen.package-macro
var __defun__shen_4record_1exceptions Obj          // shen.record-exceptions
var __defun__shen_4record_1internal Obj            // shen.record-internal
var __defun__shen_4internal_1symbols Obj           // shen.internal-symbols
var __defun__shen_4packageh Obj                    // shen.packageh

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		reg8356 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__e.Return(reg8356)
		return
	}, 0))
	__defun__shen_4read_1char_1code = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1400 := __args[0]
		_ = V1400
		reg8357 := PrimReadByte(V1400)
		__e.Return(reg8357)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.read-char-code", value: __defun__shen_4read_1char_1code})

	__defun__read_1file_1as_1bytelist = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1402 := __args[0]
		_ = V1402
		reg8358 := MakeNative(func(__e Evaluator, __args ...Obj) {
			S := __args[0]
			_ = S
			reg8359 := PrimReadByte(S)
			__e.Return(reg8359)
			return
		}, 1)
		__e.TailApply(__defun__shen_4read_1file_1as_1Xlist, V1402, reg8358)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "read-file-as-bytelist", value: __defun__read_1file_1as_1bytelist})

	__defun__shen_4read_1file_1as_1charlist = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1404 := __args[0]
		_ = V1404
		reg8361 := MakeNative(func(__e Evaluator, __args ...Obj) {
			S := __args[0]
			_ = S
			__e.TailApply(__defun__shen_4read_1char_1code, S)
			return
		}, 1)
		__e.TailApply(__defun__shen_4read_1file_1as_1Xlist, V1404, reg8361)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.read-file-as-charlist", value: __defun__shen_4read_1file_1as_1charlist})

	__defun__shen_4read_1file_1as_1Xlist = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1407 := __args[0]
		_ = V1407
		V1408 := __args[1]
		_ = V1408
		reg8364 := MakeSymbol("in")
		reg8365 := PrimOpenStream(V1407, reg8364)
		Stream := reg8365
		_ = Stream
		reg8366 := Call(__e, V1408, Stream)
		X := reg8366
		_ = X
		reg8367 := Nil
		reg8368 := Call(__e, __defun__shen_4read_1file_1as_1Xlist_1help, Stream, V1408, X, reg8367)
		Xs := reg8368
		_ = Xs
		reg8369 := PrimCloseStream(Stream)
		Close := reg8369
		_ = Close
		__e.TailApply(__defun__reverse, Xs)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.read-file-as-Xlist", value: __defun__shen_4read_1file_1as_1Xlist})

	__defun__shen_4read_1file_1as_1Xlist_1help = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1413 := __args[0]
		_ = V1413
		V1414 := __args[1]
		_ = V1414
		V1415 := __args[2]
		_ = V1415
		V1416 := __args[3]
		_ = V1416
		reg8371 := MakeNumber(-1)
		reg8372 := PrimEqual(reg8371, V1415)
		if reg8372 == True {
			__e.Return(V1416)
			return
		} else {
			reg8373 := Call(__e, V1414, V1413)
			reg8374 := PrimCons(V1415, V1416)
			__e.TailApply(__defun__shen_4read_1file_1as_1Xlist_1help, V1413, V1414, reg8373, reg8374)
			return
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.read-file-as-Xlist-help", value: __defun__shen_4read_1file_1as_1Xlist_1help})

	__defun__read_1file_1as_1string = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1418 := __args[0]
		_ = V1418
		reg8376 := MakeSymbol("in")
		reg8377 := PrimOpenStream(V1418, reg8376)
		Stream := reg8377
		_ = Stream
		reg8378 := Call(__e, __defun__shen_4read_1char_1code, Stream)
		reg8379 := MakeString("")
		__e.TailApply(__defun__shen_4rfas_1h, Stream, reg8378, reg8379)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "read-file-as-string", value: __defun__read_1file_1as_1string})

	__defun__shen_4rfas_1h = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1422 := __args[0]
		_ = V1422
		V1423 := __args[1]
		_ = V1423
		V1424 := __args[2]
		_ = V1424
		reg8381 := MakeNumber(-1)
		reg8382 := PrimEqual(reg8381, V1423)
		if reg8382 == True {
			reg8383 := PrimCloseStream(V1422)
			_ = reg8383
			__e.Return(V1424)
			return
		} else {
			reg8384 := Call(__e, __defun__shen_4read_1char_1code, V1422)
			reg8385 := PrimNumberToString(V1423)
			reg8386 := PrimStringConcat(V1424, reg8385)
			__e.TailApply(__defun__shen_4rfas_1h, V1422, reg8384, reg8386)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.rfas-h", value: __defun__shen_4rfas_1h})

	__defun__input = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1426 := __args[0]
		_ = V1426
		reg8388 := Call(__e, __defun__read, V1426)
		reg8389 := PrimEvalKL(__e, reg8388)
		__e.Return(reg8389)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "input", value: __defun__input})

	__defun__input_7 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1429 := __args[0]
		_ = V1429
		V1430 := __args[1]
		_ = V1430
		reg8390 := Call(__e, __defun__shen_4monotype, V1429)
		Mono_2 := reg8390
		_ = Mono_2
		reg8391 := Call(__e, __defun__read, V1430)
		Input := reg8391
		_ = Input
		reg8392 := False
		reg8393 := Call(__e, __defun__shen_4demodulate, V1429)
		reg8394 := Call(__e, __defun__shen_4typecheck, Input, reg8393)
		reg8395 := PrimEqual(reg8392, reg8394)
		if reg8395 == True {
			reg8396 := MakeString("type error: ")
			reg8397 := MakeString(" is not of type ")
			reg8398 := MakeString("\n")
			reg8399 := MakeSymbol("shen.r")
			reg8400 := Call(__e, __defun__shen_4app, V1429, reg8398, reg8399)
			reg8401 := PrimStringConcat(reg8397, reg8400)
			reg8402 := MakeSymbol("shen.r")
			reg8403 := Call(__e, __defun__shen_4app, Input, reg8401, reg8402)
			reg8404 := PrimStringConcat(reg8396, reg8403)
			reg8405 := PrimSimpleError(reg8404)
			__e.Return(reg8405)
			return
		} else {
			reg8406 := PrimEvalKL(__e, Input)
			__e.Return(reg8406)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "input+", value: __defun__input_7})

	__defun__shen_4monotype = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1432 := __args[0]
		_ = V1432
		reg8407 := PrimIsPair(V1432)
		if reg8407 == True {
			reg8408 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Z := __args[0]
				_ = Z
				__e.TailApply(__defun__shen_4monotype, Z)
				return
			}, 1)
			__e.TailApply(__defun__map, reg8408, V1432)
			return
		} else {
			reg8411 := PrimIsVariable(V1432)
			if reg8411 == True {
				reg8412 := MakeString("input+ expects a monotype: not ")
				reg8413 := MakeString("\n")
				reg8414 := MakeSymbol("shen.a")
				reg8415 := Call(__e, __defun__shen_4app, V1432, reg8413, reg8414)
				reg8416 := PrimStringConcat(reg8412, reg8415)
				reg8417 := PrimSimpleError(reg8416)
				__e.Return(reg8417)
				return
			} else {
				__e.Return(V1432)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.monotype", value: __defun__shen_4monotype})

	__defun__read = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1434 := __args[0]
		_ = V1434
		reg8418 := Call(__e, __defun__shen_4read_1char_1code, V1434)
		reg8419 := Nil
		reg8420 := Call(__e, __defun__shen_4read_1loop, V1434, reg8418, reg8419)
		reg8421 := PrimHead(reg8420)
		__e.Return(reg8421)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "read", value: __defun__read})

	__defun__it = MakeNative(func(__e Evaluator, __args ...Obj) {
		reg8422 := MakeSymbol("shen.*it*")
		reg8423 := PrimValue(reg8422)
		__e.Return(reg8423)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "it", value: __defun__it})

	__defun__shen_4read_1loop = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1442 := __args[0]
		_ = V1442
		V1443 := __args[1]
		_ = V1443
		V1444 := __args[2]
		_ = V1444
		reg8424 := MakeNumber(94)
		reg8425 := PrimEqual(reg8424, V1443)
		if reg8425 == True {
			reg8426 := MakeString("read aborted")
			reg8427 := PrimSimpleError(reg8426)
			__e.Return(reg8427)
			return
		} else {
			reg8428 := MakeNumber(-1)
			reg8429 := PrimEqual(reg8428, V1443)
			if reg8429 == True {
				reg8430 := Call(__e, __defun__empty_2, V1444)
				if reg8430 == True {
					reg8431 := MakeString("error: empty stream")
					reg8432 := PrimSimpleError(reg8431)
					__e.Return(reg8432)
					return
				} else {
					reg8433 := MakeNative(func(__e Evaluator, __args ...Obj) {
						X := __args[0]
						_ = X
						__e.TailApply(__defun__shen_4_5st__input_6, X)
						return
					}, 1)
					reg8435 := MakeNative(func(__e Evaluator, __args ...Obj) {
						E := __args[0]
						_ = E
						__e.Return(E)
						return
					}, 1)
					__e.TailApply(__defun__compile, reg8433, V1444, reg8435)
					return
				}
			} else {
				reg8437 := Call(__e, __defun__shen_4terminator_2, V1443)
				if reg8437 == True {
					reg8438 := Nil
					reg8439 := PrimCons(V1443, reg8438)
					reg8440 := Call(__e, __defun__append, V1444, reg8439)
					AllChars := reg8440
					_ = AllChars
					reg8441 := Call(__e, __defun__shen_4record_1it, AllChars)
					It := reg8441
					_ = It
					reg8442 := MakeNative(func(__e Evaluator, __args ...Obj) {
						X := __args[0]
						_ = X
						__e.TailApply(__defun__shen_4_5st__input_6, X)
						return
					}, 1)
					reg8444 := MakeNative(func(__e Evaluator, __args ...Obj) {
						E := __args[0]
						_ = E
						reg8445 := MakeSymbol("shen.nextbyte")
						__e.Return(reg8445)
						return
					}, 1)
					reg8446 := Call(__e, __defun__compile, reg8442, AllChars, reg8444)
					Read := reg8446
					_ = Read
					reg8447 := MakeSymbol("shen.nextbyte")
					reg8448 := PrimEqual(Read, reg8447)
					var reg8454 Obj
					if reg8448 == True {
						reg8449 := True
						reg8454 = reg8449
					} else {
						reg8450 := Call(__e, __defun__empty_2, Read)
						var reg8453 Obj
						if reg8450 == True {
							reg8451 := True
							reg8453 = reg8451
						} else {
							reg8452 := False
							reg8453 = reg8452
						}
						reg8454 = reg8453
					}
					if reg8454 == True {
						reg8455 := Call(__e, __defun__shen_4read_1char_1code, V1442)
						__e.TailApply(__defun__shen_4read_1loop, V1442, reg8455, AllChars)
						return
					} else {
						__e.Return(Read)
						return
					}
				} else {
					reg8457 := Call(__e, __defun__shen_4read_1char_1code, V1442)
					reg8458 := Nil
					reg8459 := PrimCons(V1443, reg8458)
					reg8460 := Call(__e, __defun__append, V1444, reg8459)
					__e.TailApply(__defun__shen_4read_1loop, V1442, reg8457, reg8460)
					return
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.read-loop", value: __defun__shen_4read_1loop})

	__defun__shen_4terminator_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1446 := __args[0]
		_ = V1446
		reg8462 := MakeNumber(9)
		reg8463 := MakeNumber(10)
		reg8464 := MakeNumber(13)
		reg8465 := MakeNumber(32)
		reg8466 := MakeNumber(34)
		reg8467 := MakeNumber(41)
		reg8468 := MakeNumber(93)
		reg8469 := Nil
		reg8470 := PrimCons(reg8468, reg8469)
		reg8471 := PrimCons(reg8467, reg8470)
		reg8472 := PrimCons(reg8466, reg8471)
		reg8473 := PrimCons(reg8465, reg8472)
		reg8474 := PrimCons(reg8464, reg8473)
		reg8475 := PrimCons(reg8463, reg8474)
		reg8476 := PrimCons(reg8462, reg8475)
		__e.TailApply(__defun__element_2, V1446, reg8476)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.terminator?", value: __defun__shen_4terminator_2})

	__defun__lineread = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1448 := __args[0]
		_ = V1448
		reg8478 := Call(__e, __defun__shen_4read_1char_1code, V1448)
		reg8479 := Nil
		__e.TailApply(__defun__shen_4lineread_1loop, reg8478, reg8479, V1448)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "lineread", value: __defun__lineread})

	__defun__shen_4lineread_1loop = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1453 := __args[0]
		_ = V1453
		V1454 := __args[1]
		_ = V1454
		V1455 := __args[2]
		_ = V1455
		reg8481 := MakeNumber(-1)
		reg8482 := PrimEqual(reg8481, V1453)
		if reg8482 == True {
			reg8483 := Call(__e, __defun__empty_2, V1454)
			if reg8483 == True {
				reg8484 := MakeString("empty stream")
				reg8485 := PrimSimpleError(reg8484)
				__e.Return(reg8485)
				return
			} else {
				reg8486 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(__defun__shen_4_5st__input_6, X)
					return
				}, 1)
				reg8488 := MakeNative(func(__e Evaluator, __args ...Obj) {
					E := __args[0]
					_ = E
					__e.Return(E)
					return
				}, 1)
				__e.TailApply(__defun__compile, reg8486, V1454, reg8488)
				return
			}
		} else {
			reg8490 := Call(__e, __defun__shen_4hat)
			reg8491 := PrimEqual(V1453, reg8490)
			if reg8491 == True {
				reg8492 := MakeString("line read aborted")
				reg8493 := PrimSimpleError(reg8492)
				__e.Return(reg8493)
				return
			} else {
				reg8494 := Call(__e, __defun__shen_4newline)
				reg8495 := Call(__e, __defun__shen_4carriage_1return)
				reg8496 := Nil
				reg8497 := PrimCons(reg8495, reg8496)
				reg8498 := PrimCons(reg8494, reg8497)
				reg8499 := Call(__e, __defun__element_2, V1453, reg8498)
				if reg8499 == True {
					reg8500 := MakeNative(func(__e Evaluator, __args ...Obj) {
						X := __args[0]
						_ = X
						__e.TailApply(__defun__shen_4_5st__input_6, X)
						return
					}, 1)
					reg8502 := MakeNative(func(__e Evaluator, __args ...Obj) {
						E := __args[0]
						_ = E
						reg8503 := MakeSymbol("shen.nextline")
						__e.Return(reg8503)
						return
					}, 1)
					reg8504 := Call(__e, __defun__compile, reg8500, V1454, reg8502)
					Line := reg8504
					_ = Line
					reg8505 := Call(__e, __defun__shen_4record_1it, V1454)
					It := reg8505
					_ = It
					reg8506 := MakeSymbol("shen.nextline")
					reg8507 := PrimEqual(Line, reg8506)
					var reg8513 Obj
					if reg8507 == True {
						reg8508 := True
						reg8513 = reg8508
					} else {
						reg8509 := Call(__e, __defun__empty_2, Line)
						var reg8512 Obj
						if reg8509 == True {
							reg8510 := True
							reg8512 = reg8510
						} else {
							reg8511 := False
							reg8512 = reg8511
						}
						reg8513 = reg8512
					}
					if reg8513 == True {
						reg8514 := Call(__e, __defun__shen_4read_1char_1code, V1455)
						reg8515 := Nil
						reg8516 := PrimCons(V1453, reg8515)
						reg8517 := Call(__e, __defun__append, V1454, reg8516)
						__e.TailApply(__defun__shen_4lineread_1loop, reg8514, reg8517, V1455)
						return
					} else {
						__e.Return(Line)
						return
					}
				} else {
					reg8519 := Call(__e, __defun__shen_4read_1char_1code, V1455)
					reg8520 := Nil
					reg8521 := PrimCons(V1453, reg8520)
					reg8522 := Call(__e, __defun__append, V1454, reg8521)
					__e.TailApply(__defun__shen_4lineread_1loop, reg8519, reg8522, V1455)
					return
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.lineread-loop", value: __defun__shen_4lineread_1loop})

	__defun__shen_4record_1it = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1457 := __args[0]
		_ = V1457
		reg8524 := Call(__e, __defun__shen_4trim_1whitespace, V1457)
		TrimLeft := reg8524
		_ = TrimLeft
		reg8525 := Call(__e, __defun__reverse, TrimLeft)
		reg8526 := Call(__e, __defun__shen_4trim_1whitespace, reg8525)
		TrimRight := reg8526
		_ = TrimRight
		reg8527 := Call(__e, __defun__reverse, TrimRight)
		Trimmed := reg8527
		_ = Trimmed
		__e.TailApply(__defun__shen_4record_1it_1h, Trimmed)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.record-it", value: __defun__shen_4record_1it})

	__defun__shen_4trim_1whitespace = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1459 := __args[0]
		_ = V1459
		reg8529 := PrimIsPair(V1459)
		var reg8545 Obj
		if reg8529 == True {
			reg8530 := PrimHead(V1459)
			reg8531 := MakeNumber(9)
			reg8532 := MakeNumber(10)
			reg8533 := MakeNumber(13)
			reg8534 := MakeNumber(32)
			reg8535 := Nil
			reg8536 := PrimCons(reg8534, reg8535)
			reg8537 := PrimCons(reg8533, reg8536)
			reg8538 := PrimCons(reg8532, reg8537)
			reg8539 := PrimCons(reg8531, reg8538)
			reg8540 := Call(__e, __defun__element_2, reg8530, reg8539)
			var reg8543 Obj
			if reg8540 == True {
				reg8541 := True
				reg8543 = reg8541
			} else {
				reg8542 := False
				reg8543 = reg8542
			}
			reg8545 = reg8543
		} else {
			reg8544 := False
			reg8545 = reg8544
		}
		if reg8545 == True {
			reg8546 := PrimTail(V1459)
			__e.TailApply(__defun__shen_4trim_1whitespace, reg8546)
			return
		} else {
			__e.Return(V1459)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.trim-whitespace", value: __defun__shen_4trim_1whitespace})

	__defun__shen_4record_1it_1h = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1461 := __args[0]
		_ = V1461
		reg8548 := MakeSymbol("shen.*it*")
		reg8549 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			reg8550 := PrimNumberToString(X)
			__e.Return(reg8550)
			return
		}, 1)
		reg8551 := Call(__e, __defun__map, reg8549, V1461)
		reg8552 := Call(__e, __defun__shen_4cn_1all, reg8551)
		reg8553 := PrimSet(reg8548, reg8552)
		_ = reg8553
		__e.Return(V1461)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.record-it-h", value: __defun__shen_4record_1it_1h})

	__defun__shen_4cn_1all = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1463 := __args[0]
		_ = V1463
		reg8554 := Nil
		reg8555 := PrimEqual(reg8554, V1463)
		if reg8555 == True {
			reg8556 := MakeString("")
			__e.Return(reg8556)
			return
		} else {
			reg8557 := PrimIsPair(V1463)
			if reg8557 == True {
				reg8558 := PrimHead(V1463)
				reg8559 := PrimTail(V1463)
				reg8560 := Call(__e, __defun__shen_4cn_1all, reg8559)
				reg8561 := PrimStringConcat(reg8558, reg8560)
				__e.Return(reg8561)
				return
			} else {
				reg8562 := MakeSymbol("shen.cn-all")
				__e.TailApply(__defun__shen_4f__error, reg8562)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.cn-all", value: __defun__shen_4cn_1all})

	__defun__read_1file = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1465 := __args[0]
		_ = V1465
		reg8564 := Call(__e, __defun__shen_4read_1file_1as_1charlist, V1465)
		Charlist := reg8564
		_ = Charlist
		reg8565 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4_5st__input_6, X)
			return
		}, 1)
		reg8567 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4read_1error, X)
			return
		}, 1)
		__e.TailApply(__defun__compile, reg8565, Charlist, reg8567)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "read-file", value: __defun__read_1file})

	__defun__read_1from_1string = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1467 := __args[0]
		_ = V1467
		reg8570 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			reg8571 := PrimStringToNumber(X)
			__e.Return(reg8571)
			return
		}, 1)
		reg8572 := Call(__e, __defun__explode, V1467)
		reg8573 := Call(__e, __defun__map, reg8570, reg8572)
		Ns := reg8573
		_ = Ns
		reg8574 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4_5st__input_6, X)
			return
		}, 1)
		reg8576 := MakeNative(func(__e Evaluator, __args ...Obj) {
			X := __args[0]
			_ = X
			__e.TailApply(__defun__shen_4read_1error, X)
			return
		}, 1)
		__e.TailApply(__defun__compile, reg8574, Ns, reg8576)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "read-from-string", value: __defun__read_1from_1string})

	__defun__shen_4read_1error = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1475 := __args[0]
		_ = V1475
		reg8579 := PrimIsPair(V1475)
		var reg8602 Obj
		if reg8579 == True {
			reg8580 := PrimHead(V1475)
			reg8581 := PrimIsPair(reg8580)
			var reg8597 Obj
			if reg8581 == True {
				reg8582 := PrimTail(V1475)
				reg8583 := PrimIsPair(reg8582)
				var reg8592 Obj
				if reg8583 == True {
					reg8584 := Nil
					reg8585 := PrimTail(V1475)
					reg8586 := PrimTail(reg8585)
					reg8587 := PrimEqual(reg8584, reg8586)
					var reg8590 Obj
					if reg8587 == True {
						reg8588 := True
						reg8590 = reg8588
					} else {
						reg8589 := False
						reg8590 = reg8589
					}
					reg8592 = reg8590
				} else {
					reg8591 := False
					reg8592 = reg8591
				}
				var reg8595 Obj
				if reg8592 == True {
					reg8593 := True
					reg8595 = reg8593
				} else {
					reg8594 := False
					reg8595 = reg8594
				}
				reg8597 = reg8595
			} else {
				reg8596 := False
				reg8597 = reg8596
			}
			var reg8600 Obj
			if reg8597 == True {
				reg8598 := True
				reg8600 = reg8598
			} else {
				reg8599 := False
				reg8600 = reg8599
			}
			reg8602 = reg8600
		} else {
			reg8601 := False
			reg8602 = reg8601
		}
		if reg8602 == True {
			reg8603 := MakeString("read error here:\n\n ")
			reg8604 := MakeNumber(50)
			reg8605 := PrimHead(V1475)
			reg8606 := Call(__e, __defun__shen_4compress_150, reg8604, reg8605)
			reg8607 := MakeString("\n")
			reg8608 := MakeSymbol("shen.a")
			reg8609 := Call(__e, __defun__shen_4app, reg8606, reg8607, reg8608)
			reg8610 := PrimStringConcat(reg8603, reg8609)
			reg8611 := PrimSimpleError(reg8610)
			__e.Return(reg8611)
			return
		} else {
			reg8612 := MakeString("read error\n")
			reg8613 := PrimSimpleError(reg8612)
			__e.Return(reg8613)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.read-error", value: __defun__shen_4read_1error})

	__defun__shen_4compress_150 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1482 := __args[0]
		_ = V1482
		V1483 := __args[1]
		_ = V1483
		reg8614 := Nil
		reg8615 := PrimEqual(reg8614, V1483)
		if reg8615 == True {
			reg8616 := MakeString("")
			__e.Return(reg8616)
			return
		} else {
			reg8617 := MakeNumber(0)
			reg8618 := PrimEqual(reg8617, V1482)
			if reg8618 == True {
				reg8619 := MakeString("")
				__e.Return(reg8619)
				return
			} else {
				reg8620 := PrimIsPair(V1483)
				if reg8620 == True {
					reg8621 := PrimHead(V1483)
					reg8622 := PrimNumberToString(reg8621)
					reg8623 := MakeNumber(1)
					reg8624 := PrimNumberSubtract(V1482, reg8623)
					reg8625 := PrimTail(V1483)
					reg8626 := Call(__e, __defun__shen_4compress_150, reg8624, reg8625)
					reg8627 := PrimStringConcat(reg8622, reg8626)
					__e.Return(reg8627)
					return
				} else {
					reg8628 := MakeSymbol("shen.compress-50")
					__e.TailApply(__defun__shen_4f__error, reg8628)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.compress-50", value: __defun__shen_4compress_150})

	__defun__shen_4_5st__input_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1485 := __args[0]
		_ = V1485
		reg8630 := Call(__e, __defun__shen_4_5lsb_6, V1485)
		Parse__shen_4_5lsb_6 := reg8630
		_ = Parse__shen_4_5lsb_6
		reg8631 := Call(__e, __defun__fail)
		reg8632 := PrimEqual(reg8631, Parse__shen_4_5lsb_6)
		reg8633 := PrimNot(reg8632)
		var reg8660 Obj
		if reg8633 == True {
			reg8634 := Call(__e, __defun__shen_4_5st__input1_6, Parse__shen_4_5lsb_6)
			Parse__shen_4_5st__input1_6 := reg8634
			_ = Parse__shen_4_5st__input1_6
			reg8635 := Call(__e, __defun__fail)
			reg8636 := PrimEqual(reg8635, Parse__shen_4_5st__input1_6)
			reg8637 := PrimNot(reg8636)
			var reg8658 Obj
			if reg8637 == True {
				reg8638 := Call(__e, __defun__shen_4_5rsb_6, Parse__shen_4_5st__input1_6)
				Parse__shen_4_5rsb_6 := reg8638
				_ = Parse__shen_4_5rsb_6
				reg8639 := Call(__e, __defun__fail)
				reg8640 := PrimEqual(reg8639, Parse__shen_4_5rsb_6)
				reg8641 := PrimNot(reg8640)
				var reg8656 Obj
				if reg8641 == True {
					reg8642 := Call(__e, __defun__shen_4_5st__input2_6, Parse__shen_4_5rsb_6)
					Parse__shen_4_5st__input2_6 := reg8642
					_ = Parse__shen_4_5st__input2_6
					reg8643 := Call(__e, __defun__fail)
					reg8644 := PrimEqual(reg8643, Parse__shen_4_5st__input2_6)
					reg8645 := PrimNot(reg8644)
					var reg8654 Obj
					if reg8645 == True {
						reg8646 := PrimHead(Parse__shen_4_5st__input2_6)
						reg8647 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input1_6)
						reg8648 := Call(__e, __defun__shen_4cons__form, reg8647)
						reg8649 := Call(__e, __defun__macroexpand, reg8648)
						reg8650 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input2_6)
						reg8651 := PrimCons(reg8649, reg8650)
						reg8652 := Call(__e, __defun__shen_4pair, reg8646, reg8651)
						reg8654 = reg8652
					} else {
						reg8653 := Call(__e, __defun__fail)
						reg8654 = reg8653
					}
					reg8656 = reg8654
				} else {
					reg8655 := Call(__e, __defun__fail)
					reg8656 = reg8655
				}
				reg8658 = reg8656
			} else {
				reg8657 := Call(__e, __defun__fail)
				reg8658 = reg8657
			}
			reg8660 = reg8658
		} else {
			reg8659 := Call(__e, __defun__fail)
			reg8660 = reg8659
		}
		YaccParse := reg8660
		_ = YaccParse
		reg8661 := Call(__e, __defun__fail)
		reg8662 := PrimEqual(YaccParse, reg8661)
		if reg8662 == True {
			reg8663 := Call(__e, __defun__shen_4_5lrb_6, V1485)
			Parse__shen_4_5lrb_6 := reg8663
			_ = Parse__shen_4_5lrb_6
			reg8664 := Call(__e, __defun__fail)
			reg8665 := PrimEqual(reg8664, Parse__shen_4_5lrb_6)
			reg8666 := PrimNot(reg8665)
			var reg8692 Obj
			if reg8666 == True {
				reg8667 := Call(__e, __defun__shen_4_5st__input1_6, Parse__shen_4_5lrb_6)
				Parse__shen_4_5st__input1_6 := reg8667
				_ = Parse__shen_4_5st__input1_6
				reg8668 := Call(__e, __defun__fail)
				reg8669 := PrimEqual(reg8668, Parse__shen_4_5st__input1_6)
				reg8670 := PrimNot(reg8669)
				var reg8690 Obj
				if reg8670 == True {
					reg8671 := Call(__e, __defun__shen_4_5rrb_6, Parse__shen_4_5st__input1_6)
					Parse__shen_4_5rrb_6 := reg8671
					_ = Parse__shen_4_5rrb_6
					reg8672 := Call(__e, __defun__fail)
					reg8673 := PrimEqual(reg8672, Parse__shen_4_5rrb_6)
					reg8674 := PrimNot(reg8673)
					var reg8688 Obj
					if reg8674 == True {
						reg8675 := Call(__e, __defun__shen_4_5st__input2_6, Parse__shen_4_5rrb_6)
						Parse__shen_4_5st__input2_6 := reg8675
						_ = Parse__shen_4_5st__input2_6
						reg8676 := Call(__e, __defun__fail)
						reg8677 := PrimEqual(reg8676, Parse__shen_4_5st__input2_6)
						reg8678 := PrimNot(reg8677)
						var reg8686 Obj
						if reg8678 == True {
							reg8679 := PrimHead(Parse__shen_4_5st__input2_6)
							reg8680 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input1_6)
							reg8681 := Call(__e, __defun__macroexpand, reg8680)
							reg8682 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input2_6)
							reg8683 := Call(__e, __defun__shen_4package_1macro, reg8681, reg8682)
							reg8684 := Call(__e, __defun__shen_4pair, reg8679, reg8683)
							reg8686 = reg8684
						} else {
							reg8685 := Call(__e, __defun__fail)
							reg8686 = reg8685
						}
						reg8688 = reg8686
					} else {
						reg8687 := Call(__e, __defun__fail)
						reg8688 = reg8687
					}
					reg8690 = reg8688
				} else {
					reg8689 := Call(__e, __defun__fail)
					reg8690 = reg8689
				}
				reg8692 = reg8690
			} else {
				reg8691 := Call(__e, __defun__fail)
				reg8692 = reg8691
			}
			YaccParse := reg8692
			_ = YaccParse
			reg8693 := Call(__e, __defun__fail)
			reg8694 := PrimEqual(YaccParse, reg8693)
			if reg8694 == True {
				reg8695 := Call(__e, __defun__shen_4_5lcurly_6, V1485)
				Parse__shen_4_5lcurly_6 := reg8695
				_ = Parse__shen_4_5lcurly_6
				reg8696 := Call(__e, __defun__fail)
				reg8697 := PrimEqual(reg8696, Parse__shen_4_5lcurly_6)
				reg8698 := PrimNot(reg8697)
				var reg8711 Obj
				if reg8698 == True {
					reg8699 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5lcurly_6)
					Parse__shen_4_5st__input_6 := reg8699
					_ = Parse__shen_4_5st__input_6
					reg8700 := Call(__e, __defun__fail)
					reg8701 := PrimEqual(reg8700, Parse__shen_4_5st__input_6)
					reg8702 := PrimNot(reg8701)
					var reg8709 Obj
					if reg8702 == True {
						reg8703 := PrimHead(Parse__shen_4_5st__input_6)
						reg8704 := MakeSymbol("{")
						reg8705 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
						reg8706 := PrimCons(reg8704, reg8705)
						reg8707 := Call(__e, __defun__shen_4pair, reg8703, reg8706)
						reg8709 = reg8707
					} else {
						reg8708 := Call(__e, __defun__fail)
						reg8709 = reg8708
					}
					reg8711 = reg8709
				} else {
					reg8710 := Call(__e, __defun__fail)
					reg8711 = reg8710
				}
				YaccParse := reg8711
				_ = YaccParse
				reg8712 := Call(__e, __defun__fail)
				reg8713 := PrimEqual(YaccParse, reg8712)
				if reg8713 == True {
					reg8714 := Call(__e, __defun__shen_4_5rcurly_6, V1485)
					Parse__shen_4_5rcurly_6 := reg8714
					_ = Parse__shen_4_5rcurly_6
					reg8715 := Call(__e, __defun__fail)
					reg8716 := PrimEqual(reg8715, Parse__shen_4_5rcurly_6)
					reg8717 := PrimNot(reg8716)
					var reg8730 Obj
					if reg8717 == True {
						reg8718 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5rcurly_6)
						Parse__shen_4_5st__input_6 := reg8718
						_ = Parse__shen_4_5st__input_6
						reg8719 := Call(__e, __defun__fail)
						reg8720 := PrimEqual(reg8719, Parse__shen_4_5st__input_6)
						reg8721 := PrimNot(reg8720)
						var reg8728 Obj
						if reg8721 == True {
							reg8722 := PrimHead(Parse__shen_4_5st__input_6)
							reg8723 := MakeSymbol("}")
							reg8724 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
							reg8725 := PrimCons(reg8723, reg8724)
							reg8726 := Call(__e, __defun__shen_4pair, reg8722, reg8725)
							reg8728 = reg8726
						} else {
							reg8727 := Call(__e, __defun__fail)
							reg8728 = reg8727
						}
						reg8730 = reg8728
					} else {
						reg8729 := Call(__e, __defun__fail)
						reg8730 = reg8729
					}
					YaccParse := reg8730
					_ = YaccParse
					reg8731 := Call(__e, __defun__fail)
					reg8732 := PrimEqual(YaccParse, reg8731)
					if reg8732 == True {
						reg8733 := Call(__e, __defun__shen_4_5bar_6, V1485)
						Parse__shen_4_5bar_6 := reg8733
						_ = Parse__shen_4_5bar_6
						reg8734 := Call(__e, __defun__fail)
						reg8735 := PrimEqual(reg8734, Parse__shen_4_5bar_6)
						reg8736 := PrimNot(reg8735)
						var reg8749 Obj
						if reg8736 == True {
							reg8737 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5bar_6)
							Parse__shen_4_5st__input_6 := reg8737
							_ = Parse__shen_4_5st__input_6
							reg8738 := Call(__e, __defun__fail)
							reg8739 := PrimEqual(reg8738, Parse__shen_4_5st__input_6)
							reg8740 := PrimNot(reg8739)
							var reg8747 Obj
							if reg8740 == True {
								reg8741 := PrimHead(Parse__shen_4_5st__input_6)
								reg8742 := MakeSymbol("bar!")
								reg8743 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
								reg8744 := PrimCons(reg8742, reg8743)
								reg8745 := Call(__e, __defun__shen_4pair, reg8741, reg8744)
								reg8747 = reg8745
							} else {
								reg8746 := Call(__e, __defun__fail)
								reg8747 = reg8746
							}
							reg8749 = reg8747
						} else {
							reg8748 := Call(__e, __defun__fail)
							reg8749 = reg8748
						}
						YaccParse := reg8749
						_ = YaccParse
						reg8750 := Call(__e, __defun__fail)
						reg8751 := PrimEqual(YaccParse, reg8750)
						if reg8751 == True {
							reg8752 := Call(__e, __defun__shen_4_5semicolon_6, V1485)
							Parse__shen_4_5semicolon_6 := reg8752
							_ = Parse__shen_4_5semicolon_6
							reg8753 := Call(__e, __defun__fail)
							reg8754 := PrimEqual(reg8753, Parse__shen_4_5semicolon_6)
							reg8755 := PrimNot(reg8754)
							var reg8768 Obj
							if reg8755 == True {
								reg8756 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5semicolon_6)
								Parse__shen_4_5st__input_6 := reg8756
								_ = Parse__shen_4_5st__input_6
								reg8757 := Call(__e, __defun__fail)
								reg8758 := PrimEqual(reg8757, Parse__shen_4_5st__input_6)
								reg8759 := PrimNot(reg8758)
								var reg8766 Obj
								if reg8759 == True {
									reg8760 := PrimHead(Parse__shen_4_5st__input_6)
									reg8761 := MakeSymbol(";")
									reg8762 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
									reg8763 := PrimCons(reg8761, reg8762)
									reg8764 := Call(__e, __defun__shen_4pair, reg8760, reg8763)
									reg8766 = reg8764
								} else {
									reg8765 := Call(__e, __defun__fail)
									reg8766 = reg8765
								}
								reg8768 = reg8766
							} else {
								reg8767 := Call(__e, __defun__fail)
								reg8768 = reg8767
							}
							YaccParse := reg8768
							_ = YaccParse
							reg8769 := Call(__e, __defun__fail)
							reg8770 := PrimEqual(YaccParse, reg8769)
							if reg8770 == True {
								reg8771 := Call(__e, __defun__shen_4_5colon_6, V1485)
								Parse__shen_4_5colon_6 := reg8771
								_ = Parse__shen_4_5colon_6
								reg8772 := Call(__e, __defun__fail)
								reg8773 := PrimEqual(reg8772, Parse__shen_4_5colon_6)
								reg8774 := PrimNot(reg8773)
								var reg8793 Obj
								if reg8774 == True {
									reg8775 := Call(__e, __defun__shen_4_5equal_6, Parse__shen_4_5colon_6)
									Parse__shen_4_5equal_6 := reg8775
									_ = Parse__shen_4_5equal_6
									reg8776 := Call(__e, __defun__fail)
									reg8777 := PrimEqual(reg8776, Parse__shen_4_5equal_6)
									reg8778 := PrimNot(reg8777)
									var reg8791 Obj
									if reg8778 == True {
										reg8779 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5equal_6)
										Parse__shen_4_5st__input_6 := reg8779
										_ = Parse__shen_4_5st__input_6
										reg8780 := Call(__e, __defun__fail)
										reg8781 := PrimEqual(reg8780, Parse__shen_4_5st__input_6)
										reg8782 := PrimNot(reg8781)
										var reg8789 Obj
										if reg8782 == True {
											reg8783 := PrimHead(Parse__shen_4_5st__input_6)
											reg8784 := MakeSymbol(":=")
											reg8785 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
											reg8786 := PrimCons(reg8784, reg8785)
											reg8787 := Call(__e, __defun__shen_4pair, reg8783, reg8786)
											reg8789 = reg8787
										} else {
											reg8788 := Call(__e, __defun__fail)
											reg8789 = reg8788
										}
										reg8791 = reg8789
									} else {
										reg8790 := Call(__e, __defun__fail)
										reg8791 = reg8790
									}
									reg8793 = reg8791
								} else {
									reg8792 := Call(__e, __defun__fail)
									reg8793 = reg8792
								}
								YaccParse := reg8793
								_ = YaccParse
								reg8794 := Call(__e, __defun__fail)
								reg8795 := PrimEqual(YaccParse, reg8794)
								if reg8795 == True {
									reg8796 := Call(__e, __defun__shen_4_5colon_6, V1485)
									Parse__shen_4_5colon_6 := reg8796
									_ = Parse__shen_4_5colon_6
									reg8797 := Call(__e, __defun__fail)
									reg8798 := PrimEqual(reg8797, Parse__shen_4_5colon_6)
									reg8799 := PrimNot(reg8798)
									var reg8818 Obj
									if reg8799 == True {
										reg8800 := Call(__e, __defun__shen_4_5minus_6, Parse__shen_4_5colon_6)
										Parse__shen_4_5minus_6 := reg8800
										_ = Parse__shen_4_5minus_6
										reg8801 := Call(__e, __defun__fail)
										reg8802 := PrimEqual(reg8801, Parse__shen_4_5minus_6)
										reg8803 := PrimNot(reg8802)
										var reg8816 Obj
										if reg8803 == True {
											reg8804 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5minus_6)
											Parse__shen_4_5st__input_6 := reg8804
											_ = Parse__shen_4_5st__input_6
											reg8805 := Call(__e, __defun__fail)
											reg8806 := PrimEqual(reg8805, Parse__shen_4_5st__input_6)
											reg8807 := PrimNot(reg8806)
											var reg8814 Obj
											if reg8807 == True {
												reg8808 := PrimHead(Parse__shen_4_5st__input_6)
												reg8809 := MakeSymbol(":-")
												reg8810 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
												reg8811 := PrimCons(reg8809, reg8810)
												reg8812 := Call(__e, __defun__shen_4pair, reg8808, reg8811)
												reg8814 = reg8812
											} else {
												reg8813 := Call(__e, __defun__fail)
												reg8814 = reg8813
											}
											reg8816 = reg8814
										} else {
											reg8815 := Call(__e, __defun__fail)
											reg8816 = reg8815
										}
										reg8818 = reg8816
									} else {
										reg8817 := Call(__e, __defun__fail)
										reg8818 = reg8817
									}
									YaccParse := reg8818
									_ = YaccParse
									reg8819 := Call(__e, __defun__fail)
									reg8820 := PrimEqual(YaccParse, reg8819)
									if reg8820 == True {
										reg8821 := Call(__e, __defun__shen_4_5colon_6, V1485)
										Parse__shen_4_5colon_6 := reg8821
										_ = Parse__shen_4_5colon_6
										reg8822 := Call(__e, __defun__fail)
										reg8823 := PrimEqual(reg8822, Parse__shen_4_5colon_6)
										reg8824 := PrimNot(reg8823)
										var reg8837 Obj
										if reg8824 == True {
											reg8825 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5colon_6)
											Parse__shen_4_5st__input_6 := reg8825
											_ = Parse__shen_4_5st__input_6
											reg8826 := Call(__e, __defun__fail)
											reg8827 := PrimEqual(reg8826, Parse__shen_4_5st__input_6)
											reg8828 := PrimNot(reg8827)
											var reg8835 Obj
											if reg8828 == True {
												reg8829 := PrimHead(Parse__shen_4_5st__input_6)
												reg8830 := MakeSymbol(":")
												reg8831 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
												reg8832 := PrimCons(reg8830, reg8831)
												reg8833 := Call(__e, __defun__shen_4pair, reg8829, reg8832)
												reg8835 = reg8833
											} else {
												reg8834 := Call(__e, __defun__fail)
												reg8835 = reg8834
											}
											reg8837 = reg8835
										} else {
											reg8836 := Call(__e, __defun__fail)
											reg8837 = reg8836
										}
										YaccParse := reg8837
										_ = YaccParse
										reg8838 := Call(__e, __defun__fail)
										reg8839 := PrimEqual(YaccParse, reg8838)
										if reg8839 == True {
											reg8840 := Call(__e, __defun__shen_4_5comma_6, V1485)
											Parse__shen_4_5comma_6 := reg8840
											_ = Parse__shen_4_5comma_6
											reg8841 := Call(__e, __defun__fail)
											reg8842 := PrimEqual(reg8841, Parse__shen_4_5comma_6)
											reg8843 := PrimNot(reg8842)
											var reg8857 Obj
											if reg8843 == True {
												reg8844 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5comma_6)
												Parse__shen_4_5st__input_6 := reg8844
												_ = Parse__shen_4_5st__input_6
												reg8845 := Call(__e, __defun__fail)
												reg8846 := PrimEqual(reg8845, Parse__shen_4_5st__input_6)
												reg8847 := PrimNot(reg8846)
												var reg8855 Obj
												if reg8847 == True {
													reg8848 := PrimHead(Parse__shen_4_5st__input_6)
													reg8849 := MakeString(",")
													reg8850 := PrimIntern(reg8849)
													reg8851 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
													reg8852 := PrimCons(reg8850, reg8851)
													reg8853 := Call(__e, __defun__shen_4pair, reg8848, reg8852)
													reg8855 = reg8853
												} else {
													reg8854 := Call(__e, __defun__fail)
													reg8855 = reg8854
												}
												reg8857 = reg8855
											} else {
												reg8856 := Call(__e, __defun__fail)
												reg8857 = reg8856
											}
											YaccParse := reg8857
											_ = YaccParse
											reg8858 := Call(__e, __defun__fail)
											reg8859 := PrimEqual(YaccParse, reg8858)
											if reg8859 == True {
												reg8860 := Call(__e, __defun__shen_4_5comment_6, V1485)
												Parse__shen_4_5comment_6 := reg8860
												_ = Parse__shen_4_5comment_6
												reg8861 := Call(__e, __defun__fail)
												reg8862 := PrimEqual(reg8861, Parse__shen_4_5comment_6)
												reg8863 := PrimNot(reg8862)
												var reg8874 Obj
												if reg8863 == True {
													reg8864 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5comment_6)
													Parse__shen_4_5st__input_6 := reg8864
													_ = Parse__shen_4_5st__input_6
													reg8865 := Call(__e, __defun__fail)
													reg8866 := PrimEqual(reg8865, Parse__shen_4_5st__input_6)
													reg8867 := PrimNot(reg8866)
													var reg8872 Obj
													if reg8867 == True {
														reg8868 := PrimHead(Parse__shen_4_5st__input_6)
														reg8869 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
														reg8870 := Call(__e, __defun__shen_4pair, reg8868, reg8869)
														reg8872 = reg8870
													} else {
														reg8871 := Call(__e, __defun__fail)
														reg8872 = reg8871
													}
													reg8874 = reg8872
												} else {
													reg8873 := Call(__e, __defun__fail)
													reg8874 = reg8873
												}
												YaccParse := reg8874
												_ = YaccParse
												reg8875 := Call(__e, __defun__fail)
												reg8876 := PrimEqual(YaccParse, reg8875)
												if reg8876 == True {
													reg8877 := Call(__e, __defun__shen_4_5atom_6, V1485)
													Parse__shen_4_5atom_6 := reg8877
													_ = Parse__shen_4_5atom_6
													reg8878 := Call(__e, __defun__fail)
													reg8879 := PrimEqual(reg8878, Parse__shen_4_5atom_6)
													reg8880 := PrimNot(reg8879)
													var reg8894 Obj
													if reg8880 == True {
														reg8881 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5atom_6)
														Parse__shen_4_5st__input_6 := reg8881
														_ = Parse__shen_4_5st__input_6
														reg8882 := Call(__e, __defun__fail)
														reg8883 := PrimEqual(reg8882, Parse__shen_4_5st__input_6)
														reg8884 := PrimNot(reg8883)
														var reg8892 Obj
														if reg8884 == True {
															reg8885 := PrimHead(Parse__shen_4_5st__input_6)
															reg8886 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5atom_6)
															reg8887 := Call(__e, __defun__macroexpand, reg8886)
															reg8888 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
															reg8889 := PrimCons(reg8887, reg8888)
															reg8890 := Call(__e, __defun__shen_4pair, reg8885, reg8889)
															reg8892 = reg8890
														} else {
															reg8891 := Call(__e, __defun__fail)
															reg8892 = reg8891
														}
														reg8894 = reg8892
													} else {
														reg8893 := Call(__e, __defun__fail)
														reg8894 = reg8893
													}
													YaccParse := reg8894
													_ = YaccParse
													reg8895 := Call(__e, __defun__fail)
													reg8896 := PrimEqual(YaccParse, reg8895)
													if reg8896 == True {
														reg8897 := Call(__e, __defun__shen_4_5whitespaces_6, V1485)
														Parse__shen_4_5whitespaces_6 := reg8897
														_ = Parse__shen_4_5whitespaces_6
														reg8898 := Call(__e, __defun__fail)
														reg8899 := PrimEqual(reg8898, Parse__shen_4_5whitespaces_6)
														reg8900 := PrimNot(reg8899)
														var reg8911 Obj
														if reg8900 == True {
															reg8901 := Call(__e, __defun__shen_4_5st__input_6, Parse__shen_4_5whitespaces_6)
															Parse__shen_4_5st__input_6 := reg8901
															_ = Parse__shen_4_5st__input_6
															reg8902 := Call(__e, __defun__fail)
															reg8903 := PrimEqual(reg8902, Parse__shen_4_5st__input_6)
															reg8904 := PrimNot(reg8903)
															var reg8909 Obj
															if reg8904 == True {
																reg8905 := PrimHead(Parse__shen_4_5st__input_6)
																reg8906 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
																reg8907 := Call(__e, __defun__shen_4pair, reg8905, reg8906)
																reg8909 = reg8907
															} else {
																reg8908 := Call(__e, __defun__fail)
																reg8909 = reg8908
															}
															reg8911 = reg8909
														} else {
															reg8910 := Call(__e, __defun__fail)
															reg8911 = reg8910
														}
														YaccParse := reg8911
														_ = YaccParse
														reg8912 := Call(__e, __defun__fail)
														reg8913 := PrimEqual(YaccParse, reg8912)
														if reg8913 == True {
															reg8914 := Call(__e, __defun___5e_6, V1485)
															Parse___5e_6 := reg8914
															_ = Parse___5e_6
															reg8915 := Call(__e, __defun__fail)
															reg8916 := PrimEqual(reg8915, Parse___5e_6)
															reg8917 := PrimNot(reg8916)
															if reg8917 == True {
																reg8918 := PrimHead(Parse___5e_6)
																reg8919 := Nil
																__e.TailApply(__defun__shen_4pair, reg8918, reg8919)
																return
															} else {
																__e.TailApply(__defun__fail)
																return
															}
														} else {
															__e.Return(YaccParse)
															return
														}
													} else {
														__e.Return(YaccParse)
														return
													}
												} else {
													__e.Return(YaccParse)
													return
												}
											} else {
												__e.Return(YaccParse)
												return
											}
										} else {
											__e.Return(YaccParse)
											return
										}
									} else {
										__e.Return(YaccParse)
										return
									}
								} else {
									__e.Return(YaccParse)
									return
								}
							} else {
								__e.Return(YaccParse)
								return
							}
						} else {
							__e.Return(YaccParse)
							return
						}
					} else {
						__e.Return(YaccParse)
						return
					}
				} else {
					__e.Return(YaccParse)
					return
				}
			} else {
				__e.Return(YaccParse)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<st_input>", value: __defun__shen_4_5st__input_6})

	__defun__shen_4_5lsb_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1488 := __args[0]
		_ = V1488
		reg8922 := PrimHead(V1488)
		reg8923 := PrimIsPair(reg8922)
		var reg8931 Obj
		if reg8923 == True {
			reg8924 := MakeNumber(91)
			reg8925 := Call(__e, __defun__shen_4hdhd, V1488)
			reg8926 := PrimEqual(reg8924, reg8925)
			var reg8929 Obj
			if reg8926 == True {
				reg8927 := True
				reg8929 = reg8927
			} else {
				reg8928 := False
				reg8929 = reg8928
			}
			reg8931 = reg8929
		} else {
			reg8930 := False
			reg8931 = reg8930
		}
		if reg8931 == True {
			reg8932 := Call(__e, __defun__shen_4tlhd, V1488)
			reg8933 := Call(__e, __defun__shen_4hdtl, V1488)
			reg8934 := Call(__e, __defun__shen_4pair, reg8932, reg8933)
			NewStream1486 := reg8934
			_ = NewStream1486
			reg8935 := PrimHead(NewStream1486)
			reg8936 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg8935, reg8936)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<lsb>", value: __defun__shen_4_5lsb_6})

	__defun__shen_4_5rsb_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1491 := __args[0]
		_ = V1491
		reg8939 := PrimHead(V1491)
		reg8940 := PrimIsPair(reg8939)
		var reg8948 Obj
		if reg8940 == True {
			reg8941 := MakeNumber(93)
			reg8942 := Call(__e, __defun__shen_4hdhd, V1491)
			reg8943 := PrimEqual(reg8941, reg8942)
			var reg8946 Obj
			if reg8943 == True {
				reg8944 := True
				reg8946 = reg8944
			} else {
				reg8945 := False
				reg8946 = reg8945
			}
			reg8948 = reg8946
		} else {
			reg8947 := False
			reg8948 = reg8947
		}
		if reg8948 == True {
			reg8949 := Call(__e, __defun__shen_4tlhd, V1491)
			reg8950 := Call(__e, __defun__shen_4hdtl, V1491)
			reg8951 := Call(__e, __defun__shen_4pair, reg8949, reg8950)
			NewStream1489 := reg8951
			_ = NewStream1489
			reg8952 := PrimHead(NewStream1489)
			reg8953 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg8952, reg8953)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<rsb>", value: __defun__shen_4_5rsb_6})

	__defun__shen_4_5lcurly_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1494 := __args[0]
		_ = V1494
		reg8956 := PrimHead(V1494)
		reg8957 := PrimIsPair(reg8956)
		var reg8965 Obj
		if reg8957 == True {
			reg8958 := MakeNumber(123)
			reg8959 := Call(__e, __defun__shen_4hdhd, V1494)
			reg8960 := PrimEqual(reg8958, reg8959)
			var reg8963 Obj
			if reg8960 == True {
				reg8961 := True
				reg8963 = reg8961
			} else {
				reg8962 := False
				reg8963 = reg8962
			}
			reg8965 = reg8963
		} else {
			reg8964 := False
			reg8965 = reg8964
		}
		if reg8965 == True {
			reg8966 := Call(__e, __defun__shen_4tlhd, V1494)
			reg8967 := Call(__e, __defun__shen_4hdtl, V1494)
			reg8968 := Call(__e, __defun__shen_4pair, reg8966, reg8967)
			NewStream1492 := reg8968
			_ = NewStream1492
			reg8969 := PrimHead(NewStream1492)
			reg8970 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg8969, reg8970)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<lcurly>", value: __defun__shen_4_5lcurly_6})

	__defun__shen_4_5rcurly_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1497 := __args[0]
		_ = V1497
		reg8973 := PrimHead(V1497)
		reg8974 := PrimIsPair(reg8973)
		var reg8982 Obj
		if reg8974 == True {
			reg8975 := MakeNumber(125)
			reg8976 := Call(__e, __defun__shen_4hdhd, V1497)
			reg8977 := PrimEqual(reg8975, reg8976)
			var reg8980 Obj
			if reg8977 == True {
				reg8978 := True
				reg8980 = reg8978
			} else {
				reg8979 := False
				reg8980 = reg8979
			}
			reg8982 = reg8980
		} else {
			reg8981 := False
			reg8982 = reg8981
		}
		if reg8982 == True {
			reg8983 := Call(__e, __defun__shen_4tlhd, V1497)
			reg8984 := Call(__e, __defun__shen_4hdtl, V1497)
			reg8985 := Call(__e, __defun__shen_4pair, reg8983, reg8984)
			NewStream1495 := reg8985
			_ = NewStream1495
			reg8986 := PrimHead(NewStream1495)
			reg8987 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg8986, reg8987)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<rcurly>", value: __defun__shen_4_5rcurly_6})

	__defun__shen_4_5bar_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1500 := __args[0]
		_ = V1500
		reg8990 := PrimHead(V1500)
		reg8991 := PrimIsPair(reg8990)
		var reg8999 Obj
		if reg8991 == True {
			reg8992 := MakeNumber(124)
			reg8993 := Call(__e, __defun__shen_4hdhd, V1500)
			reg8994 := PrimEqual(reg8992, reg8993)
			var reg8997 Obj
			if reg8994 == True {
				reg8995 := True
				reg8997 = reg8995
			} else {
				reg8996 := False
				reg8997 = reg8996
			}
			reg8999 = reg8997
		} else {
			reg8998 := False
			reg8999 = reg8998
		}
		if reg8999 == True {
			reg9000 := Call(__e, __defun__shen_4tlhd, V1500)
			reg9001 := Call(__e, __defun__shen_4hdtl, V1500)
			reg9002 := Call(__e, __defun__shen_4pair, reg9000, reg9001)
			NewStream1498 := reg9002
			_ = NewStream1498
			reg9003 := PrimHead(NewStream1498)
			reg9004 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9003, reg9004)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<bar>", value: __defun__shen_4_5bar_6})

	__defun__shen_4_5semicolon_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1503 := __args[0]
		_ = V1503
		reg9007 := PrimHead(V1503)
		reg9008 := PrimIsPair(reg9007)
		var reg9016 Obj
		if reg9008 == True {
			reg9009 := MakeNumber(59)
			reg9010 := Call(__e, __defun__shen_4hdhd, V1503)
			reg9011 := PrimEqual(reg9009, reg9010)
			var reg9014 Obj
			if reg9011 == True {
				reg9012 := True
				reg9014 = reg9012
			} else {
				reg9013 := False
				reg9014 = reg9013
			}
			reg9016 = reg9014
		} else {
			reg9015 := False
			reg9016 = reg9015
		}
		if reg9016 == True {
			reg9017 := Call(__e, __defun__shen_4tlhd, V1503)
			reg9018 := Call(__e, __defun__shen_4hdtl, V1503)
			reg9019 := Call(__e, __defun__shen_4pair, reg9017, reg9018)
			NewStream1501 := reg9019
			_ = NewStream1501
			reg9020 := PrimHead(NewStream1501)
			reg9021 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9020, reg9021)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<semicolon>", value: __defun__shen_4_5semicolon_6})

	__defun__shen_4_5colon_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1506 := __args[0]
		_ = V1506
		reg9024 := PrimHead(V1506)
		reg9025 := PrimIsPair(reg9024)
		var reg9033 Obj
		if reg9025 == True {
			reg9026 := MakeNumber(58)
			reg9027 := Call(__e, __defun__shen_4hdhd, V1506)
			reg9028 := PrimEqual(reg9026, reg9027)
			var reg9031 Obj
			if reg9028 == True {
				reg9029 := True
				reg9031 = reg9029
			} else {
				reg9030 := False
				reg9031 = reg9030
			}
			reg9033 = reg9031
		} else {
			reg9032 := False
			reg9033 = reg9032
		}
		if reg9033 == True {
			reg9034 := Call(__e, __defun__shen_4tlhd, V1506)
			reg9035 := Call(__e, __defun__shen_4hdtl, V1506)
			reg9036 := Call(__e, __defun__shen_4pair, reg9034, reg9035)
			NewStream1504 := reg9036
			_ = NewStream1504
			reg9037 := PrimHead(NewStream1504)
			reg9038 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9037, reg9038)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<colon>", value: __defun__shen_4_5colon_6})

	__defun__shen_4_5comma_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1509 := __args[0]
		_ = V1509
		reg9041 := PrimHead(V1509)
		reg9042 := PrimIsPair(reg9041)
		var reg9050 Obj
		if reg9042 == True {
			reg9043 := MakeNumber(44)
			reg9044 := Call(__e, __defun__shen_4hdhd, V1509)
			reg9045 := PrimEqual(reg9043, reg9044)
			var reg9048 Obj
			if reg9045 == True {
				reg9046 := True
				reg9048 = reg9046
			} else {
				reg9047 := False
				reg9048 = reg9047
			}
			reg9050 = reg9048
		} else {
			reg9049 := False
			reg9050 = reg9049
		}
		if reg9050 == True {
			reg9051 := Call(__e, __defun__shen_4tlhd, V1509)
			reg9052 := Call(__e, __defun__shen_4hdtl, V1509)
			reg9053 := Call(__e, __defun__shen_4pair, reg9051, reg9052)
			NewStream1507 := reg9053
			_ = NewStream1507
			reg9054 := PrimHead(NewStream1507)
			reg9055 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9054, reg9055)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<comma>", value: __defun__shen_4_5comma_6})

	__defun__shen_4_5equal_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1512 := __args[0]
		_ = V1512
		reg9058 := PrimHead(V1512)
		reg9059 := PrimIsPair(reg9058)
		var reg9067 Obj
		if reg9059 == True {
			reg9060 := MakeNumber(61)
			reg9061 := Call(__e, __defun__shen_4hdhd, V1512)
			reg9062 := PrimEqual(reg9060, reg9061)
			var reg9065 Obj
			if reg9062 == True {
				reg9063 := True
				reg9065 = reg9063
			} else {
				reg9064 := False
				reg9065 = reg9064
			}
			reg9067 = reg9065
		} else {
			reg9066 := False
			reg9067 = reg9066
		}
		if reg9067 == True {
			reg9068 := Call(__e, __defun__shen_4tlhd, V1512)
			reg9069 := Call(__e, __defun__shen_4hdtl, V1512)
			reg9070 := Call(__e, __defun__shen_4pair, reg9068, reg9069)
			NewStream1510 := reg9070
			_ = NewStream1510
			reg9071 := PrimHead(NewStream1510)
			reg9072 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9071, reg9072)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<equal>", value: __defun__shen_4_5equal_6})

	__defun__shen_4_5minus_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1515 := __args[0]
		_ = V1515
		reg9075 := PrimHead(V1515)
		reg9076 := PrimIsPair(reg9075)
		var reg9084 Obj
		if reg9076 == True {
			reg9077 := MakeNumber(45)
			reg9078 := Call(__e, __defun__shen_4hdhd, V1515)
			reg9079 := PrimEqual(reg9077, reg9078)
			var reg9082 Obj
			if reg9079 == True {
				reg9080 := True
				reg9082 = reg9080
			} else {
				reg9081 := False
				reg9082 = reg9081
			}
			reg9084 = reg9082
		} else {
			reg9083 := False
			reg9084 = reg9083
		}
		if reg9084 == True {
			reg9085 := Call(__e, __defun__shen_4tlhd, V1515)
			reg9086 := Call(__e, __defun__shen_4hdtl, V1515)
			reg9087 := Call(__e, __defun__shen_4pair, reg9085, reg9086)
			NewStream1513 := reg9087
			_ = NewStream1513
			reg9088 := PrimHead(NewStream1513)
			reg9089 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9088, reg9089)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<minus>", value: __defun__shen_4_5minus_6})

	__defun__shen_4_5lrb_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1518 := __args[0]
		_ = V1518
		reg9092 := PrimHead(V1518)
		reg9093 := PrimIsPair(reg9092)
		var reg9101 Obj
		if reg9093 == True {
			reg9094 := MakeNumber(40)
			reg9095 := Call(__e, __defun__shen_4hdhd, V1518)
			reg9096 := PrimEqual(reg9094, reg9095)
			var reg9099 Obj
			if reg9096 == True {
				reg9097 := True
				reg9099 = reg9097
			} else {
				reg9098 := False
				reg9099 = reg9098
			}
			reg9101 = reg9099
		} else {
			reg9100 := False
			reg9101 = reg9100
		}
		if reg9101 == True {
			reg9102 := Call(__e, __defun__shen_4tlhd, V1518)
			reg9103 := Call(__e, __defun__shen_4hdtl, V1518)
			reg9104 := Call(__e, __defun__shen_4pair, reg9102, reg9103)
			NewStream1516 := reg9104
			_ = NewStream1516
			reg9105 := PrimHead(NewStream1516)
			reg9106 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9105, reg9106)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<lrb>", value: __defun__shen_4_5lrb_6})

	__defun__shen_4_5rrb_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1521 := __args[0]
		_ = V1521
		reg9109 := PrimHead(V1521)
		reg9110 := PrimIsPair(reg9109)
		var reg9118 Obj
		if reg9110 == True {
			reg9111 := MakeNumber(41)
			reg9112 := Call(__e, __defun__shen_4hdhd, V1521)
			reg9113 := PrimEqual(reg9111, reg9112)
			var reg9116 Obj
			if reg9113 == True {
				reg9114 := True
				reg9116 = reg9114
			} else {
				reg9115 := False
				reg9116 = reg9115
			}
			reg9118 = reg9116
		} else {
			reg9117 := False
			reg9118 = reg9117
		}
		if reg9118 == True {
			reg9119 := Call(__e, __defun__shen_4tlhd, V1521)
			reg9120 := Call(__e, __defun__shen_4hdtl, V1521)
			reg9121 := Call(__e, __defun__shen_4pair, reg9119, reg9120)
			NewStream1519 := reg9121
			_ = NewStream1519
			reg9122 := PrimHead(NewStream1519)
			reg9123 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9122, reg9123)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<rrb>", value: __defun__shen_4_5rrb_6})

	__defun__shen_4_5atom_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1523 := __args[0]
		_ = V1523
		reg9126 := Call(__e, __defun__shen_4_5str_6, V1523)
		Parse__shen_4_5str_6 := reg9126
		_ = Parse__shen_4_5str_6
		reg9127 := Call(__e, __defun__fail)
		reg9128 := PrimEqual(reg9127, Parse__shen_4_5str_6)
		reg9129 := PrimNot(reg9128)
		var reg9135 Obj
		if reg9129 == True {
			reg9130 := PrimHead(Parse__shen_4_5str_6)
			reg9131 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5str_6)
			reg9132 := Call(__e, __defun__shen_4control_1chars, reg9131)
			reg9133 := Call(__e, __defun__shen_4pair, reg9130, reg9132)
			reg9135 = reg9133
		} else {
			reg9134 := Call(__e, __defun__fail)
			reg9135 = reg9134
		}
		YaccParse := reg9135
		_ = YaccParse
		reg9136 := Call(__e, __defun__fail)
		reg9137 := PrimEqual(YaccParse, reg9136)
		if reg9137 == True {
			reg9138 := Call(__e, __defun__shen_4_5number_6, V1523)
			Parse__shen_4_5number_6 := reg9138
			_ = Parse__shen_4_5number_6
			reg9139 := Call(__e, __defun__fail)
			reg9140 := PrimEqual(reg9139, Parse__shen_4_5number_6)
			reg9141 := PrimNot(reg9140)
			var reg9146 Obj
			if reg9141 == True {
				reg9142 := PrimHead(Parse__shen_4_5number_6)
				reg9143 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5number_6)
				reg9144 := Call(__e, __defun__shen_4pair, reg9142, reg9143)
				reg9146 = reg9144
			} else {
				reg9145 := Call(__e, __defun__fail)
				reg9146 = reg9145
			}
			YaccParse := reg9146
			_ = YaccParse
			reg9147 := Call(__e, __defun__fail)
			reg9148 := PrimEqual(YaccParse, reg9147)
			if reg9148 == True {
				reg9149 := Call(__e, __defun__shen_4_5sym_6, V1523)
				Parse__shen_4_5sym_6 := reg9149
				_ = Parse__shen_4_5sym_6
				reg9150 := Call(__e, __defun__fail)
				reg9151 := PrimEqual(reg9150, Parse__shen_4_5sym_6)
				reg9152 := PrimNot(reg9151)
				if reg9152 == True {
					reg9153 := PrimHead(Parse__shen_4_5sym_6)
					reg9154 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5sym_6)
					reg9155 := MakeString("<>")
					reg9156 := PrimEqual(reg9154, reg9155)
					var reg9164 Obj
					if reg9156 == True {
						reg9157 := MakeSymbol("vector")
						reg9158 := MakeNumber(0)
						reg9159 := Nil
						reg9160 := PrimCons(reg9158, reg9159)
						reg9161 := PrimCons(reg9157, reg9160)
						reg9164 = reg9161
					} else {
						reg9162 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5sym_6)
						reg9163 := PrimIntern(reg9162)
						reg9164 = reg9163
					}
					__e.TailApply(__defun__shen_4pair, reg9153, reg9164)
					return
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.Return(YaccParse)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<atom>", value: __defun__shen_4_5atom_6})

	__defun__shen_4control_1chars = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1525 := __args[0]
		_ = V1525
		reg9167 := Nil
		reg9168 := PrimEqual(reg9167, V1525)
		if reg9168 == True {
			reg9169 := MakeString("")
			__e.Return(reg9169)
			return
		} else {
			reg9170 := PrimIsPair(V1525)
			var reg9194 Obj
			if reg9170 == True {
				reg9171 := MakeString("c")
				reg9172 := PrimHead(V1525)
				reg9173 := PrimEqual(reg9171, reg9172)
				var reg9189 Obj
				if reg9173 == True {
					reg9174 := PrimTail(V1525)
					reg9175 := PrimIsPair(reg9174)
					var reg9184 Obj
					if reg9175 == True {
						reg9176 := MakeString("#")
						reg9177 := PrimTail(V1525)
						reg9178 := PrimHead(reg9177)
						reg9179 := PrimEqual(reg9176, reg9178)
						var reg9182 Obj
						if reg9179 == True {
							reg9180 := True
							reg9182 = reg9180
						} else {
							reg9181 := False
							reg9182 = reg9181
						}
						reg9184 = reg9182
					} else {
						reg9183 := False
						reg9184 = reg9183
					}
					var reg9187 Obj
					if reg9184 == True {
						reg9185 := True
						reg9187 = reg9185
					} else {
						reg9186 := False
						reg9187 = reg9186
					}
					reg9189 = reg9187
				} else {
					reg9188 := False
					reg9189 = reg9188
				}
				var reg9192 Obj
				if reg9189 == True {
					reg9190 := True
					reg9192 = reg9190
				} else {
					reg9191 := False
					reg9192 = reg9191
				}
				reg9194 = reg9192
			} else {
				reg9193 := False
				reg9194 = reg9193
			}
			if reg9194 == True {
				reg9195 := PrimTail(V1525)
				reg9196 := PrimTail(reg9195)
				reg9197 := Call(__e, __defun__shen_4code_1point, reg9196)
				CodePoint := reg9197
				_ = CodePoint
				reg9198 := PrimTail(V1525)
				reg9199 := PrimTail(reg9198)
				reg9200 := Call(__e, __defun__shen_4after_1codepoint, reg9199)
				AfterCodePoint := reg9200
				_ = AfterCodePoint
				reg9201 := Call(__e, __defun__shen_4decimalise, CodePoint)
				reg9202 := PrimNumberToString(reg9201)
				reg9203 := Call(__e, __defun__shen_4control_1chars, AfterCodePoint)
				__e.TailApply(__defun___8s, reg9202, reg9203)
				return
			} else {
				reg9205 := PrimIsPair(V1525)
				if reg9205 == True {
					reg9206 := PrimHead(V1525)
					reg9207 := PrimTail(V1525)
					reg9208 := Call(__e, __defun__shen_4control_1chars, reg9207)
					__e.TailApply(__defun___8s, reg9206, reg9208)
					return
				} else {
					reg9210 := MakeSymbol("shen.control-chars")
					__e.TailApply(__defun__shen_4f__error, reg9210)
					return
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.control-chars", value: __defun__shen_4control_1chars})

	__defun__shen_4code_1point = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1529 := __args[0]
		_ = V1529
		reg9212 := PrimIsPair(V1529)
		var reg9220 Obj
		if reg9212 == True {
			reg9213 := MakeString(";")
			reg9214 := PrimHead(V1529)
			reg9215 := PrimEqual(reg9213, reg9214)
			var reg9218 Obj
			if reg9215 == True {
				reg9216 := True
				reg9218 = reg9216
			} else {
				reg9217 := False
				reg9218 = reg9217
			}
			reg9220 = reg9218
		} else {
			reg9219 := False
			reg9220 = reg9219
		}
		if reg9220 == True {
			reg9221 := MakeString("")
			__e.Return(reg9221)
			return
		} else {
			reg9222 := PrimIsPair(V1529)
			var reg9252 Obj
			if reg9222 == True {
				reg9223 := PrimHead(V1529)
				reg9224 := MakeString("0")
				reg9225 := MakeString("1")
				reg9226 := MakeString("2")
				reg9227 := MakeString("3")
				reg9228 := MakeString("4")
				reg9229 := MakeString("5")
				reg9230 := MakeString("6")
				reg9231 := MakeString("7")
				reg9232 := MakeString("8")
				reg9233 := MakeString("9")
				reg9234 := MakeString("0")
				reg9235 := Nil
				reg9236 := PrimCons(reg9234, reg9235)
				reg9237 := PrimCons(reg9233, reg9236)
				reg9238 := PrimCons(reg9232, reg9237)
				reg9239 := PrimCons(reg9231, reg9238)
				reg9240 := PrimCons(reg9230, reg9239)
				reg9241 := PrimCons(reg9229, reg9240)
				reg9242 := PrimCons(reg9228, reg9241)
				reg9243 := PrimCons(reg9227, reg9242)
				reg9244 := PrimCons(reg9226, reg9243)
				reg9245 := PrimCons(reg9225, reg9244)
				reg9246 := PrimCons(reg9224, reg9245)
				reg9247 := Call(__e, __defun__element_2, reg9223, reg9246)
				var reg9250 Obj
				if reg9247 == True {
					reg9248 := True
					reg9250 = reg9248
				} else {
					reg9249 := False
					reg9250 = reg9249
				}
				reg9252 = reg9250
			} else {
				reg9251 := False
				reg9252 = reg9251
			}
			if reg9252 == True {
				reg9253 := PrimHead(V1529)
				reg9254 := PrimTail(V1529)
				reg9255 := Call(__e, __defun__shen_4code_1point, reg9254)
				reg9256 := PrimCons(reg9253, reg9255)
				__e.Return(reg9256)
				return
			} else {
				reg9257 := MakeString("code point parse error ")
				reg9258 := MakeString("\n")
				reg9259 := MakeSymbol("shen.a")
				reg9260 := Call(__e, __defun__shen_4app, V1529, reg9258, reg9259)
				reg9261 := PrimStringConcat(reg9257, reg9260)
				reg9262 := PrimSimpleError(reg9261)
				__e.Return(reg9262)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.code-point", value: __defun__shen_4code_1point})

	__defun__shen_4after_1codepoint = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1535 := __args[0]
		_ = V1535
		reg9263 := Nil
		reg9264 := PrimEqual(reg9263, V1535)
		if reg9264 == True {
			reg9265 := Nil
			__e.Return(reg9265)
			return
		} else {
			reg9266 := PrimIsPair(V1535)
			var reg9274 Obj
			if reg9266 == True {
				reg9267 := MakeString(";")
				reg9268 := PrimHead(V1535)
				reg9269 := PrimEqual(reg9267, reg9268)
				var reg9272 Obj
				if reg9269 == True {
					reg9270 := True
					reg9272 = reg9270
				} else {
					reg9271 := False
					reg9272 = reg9271
				}
				reg9274 = reg9272
			} else {
				reg9273 := False
				reg9274 = reg9273
			}
			if reg9274 == True {
				reg9275 := PrimTail(V1535)
				__e.Return(reg9275)
				return
			} else {
				reg9276 := PrimIsPair(V1535)
				if reg9276 == True {
					reg9277 := PrimTail(V1535)
					__e.TailApply(__defun__shen_4after_1codepoint, reg9277)
					return
				} else {
					reg9279 := MakeSymbol("shen.after-codepoint")
					__e.TailApply(__defun__shen_4f__error, reg9279)
					return
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.after-codepoint", value: __defun__shen_4after_1codepoint})

	__defun__shen_4decimalise = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1537 := __args[0]
		_ = V1537
		reg9281 := Call(__e, __defun__shen_4digits_1_6integers, V1537)
		reg9282 := Call(__e, __defun__reverse, reg9281)
		reg9283 := MakeNumber(0)
		__e.TailApply(__defun__shen_4pre, reg9282, reg9283)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.decimalise", value: __defun__shen_4decimalise})

	__defun__shen_4digits_1_6integers = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1543 := __args[0]
		_ = V1543
		reg9285 := PrimIsPair(V1543)
		var reg9293 Obj
		if reg9285 == True {
			reg9286 := MakeString("0")
			reg9287 := PrimHead(V1543)
			reg9288 := PrimEqual(reg9286, reg9287)
			var reg9291 Obj
			if reg9288 == True {
				reg9289 := True
				reg9291 = reg9289
			} else {
				reg9290 := False
				reg9291 = reg9290
			}
			reg9293 = reg9291
		} else {
			reg9292 := False
			reg9293 = reg9292
		}
		if reg9293 == True {
			reg9294 := MakeNumber(0)
			reg9295 := PrimTail(V1543)
			reg9296 := Call(__e, __defun__shen_4digits_1_6integers, reg9295)
			reg9297 := PrimCons(reg9294, reg9296)
			__e.Return(reg9297)
			return
		} else {
			reg9298 := PrimIsPair(V1543)
			var reg9306 Obj
			if reg9298 == True {
				reg9299 := MakeString("1")
				reg9300 := PrimHead(V1543)
				reg9301 := PrimEqual(reg9299, reg9300)
				var reg9304 Obj
				if reg9301 == True {
					reg9302 := True
					reg9304 = reg9302
				} else {
					reg9303 := False
					reg9304 = reg9303
				}
				reg9306 = reg9304
			} else {
				reg9305 := False
				reg9306 = reg9305
			}
			if reg9306 == True {
				reg9307 := MakeNumber(1)
				reg9308 := PrimTail(V1543)
				reg9309 := Call(__e, __defun__shen_4digits_1_6integers, reg9308)
				reg9310 := PrimCons(reg9307, reg9309)
				__e.Return(reg9310)
				return
			} else {
				reg9311 := PrimIsPair(V1543)
				var reg9319 Obj
				if reg9311 == True {
					reg9312 := MakeString("2")
					reg9313 := PrimHead(V1543)
					reg9314 := PrimEqual(reg9312, reg9313)
					var reg9317 Obj
					if reg9314 == True {
						reg9315 := True
						reg9317 = reg9315
					} else {
						reg9316 := False
						reg9317 = reg9316
					}
					reg9319 = reg9317
				} else {
					reg9318 := False
					reg9319 = reg9318
				}
				if reg9319 == True {
					reg9320 := MakeNumber(2)
					reg9321 := PrimTail(V1543)
					reg9322 := Call(__e, __defun__shen_4digits_1_6integers, reg9321)
					reg9323 := PrimCons(reg9320, reg9322)
					__e.Return(reg9323)
					return
				} else {
					reg9324 := PrimIsPair(V1543)
					var reg9332 Obj
					if reg9324 == True {
						reg9325 := MakeString("3")
						reg9326 := PrimHead(V1543)
						reg9327 := PrimEqual(reg9325, reg9326)
						var reg9330 Obj
						if reg9327 == True {
							reg9328 := True
							reg9330 = reg9328
						} else {
							reg9329 := False
							reg9330 = reg9329
						}
						reg9332 = reg9330
					} else {
						reg9331 := False
						reg9332 = reg9331
					}
					if reg9332 == True {
						reg9333 := MakeNumber(3)
						reg9334 := PrimTail(V1543)
						reg9335 := Call(__e, __defun__shen_4digits_1_6integers, reg9334)
						reg9336 := PrimCons(reg9333, reg9335)
						__e.Return(reg9336)
						return
					} else {
						reg9337 := PrimIsPair(V1543)
						var reg9345 Obj
						if reg9337 == True {
							reg9338 := MakeString("4")
							reg9339 := PrimHead(V1543)
							reg9340 := PrimEqual(reg9338, reg9339)
							var reg9343 Obj
							if reg9340 == True {
								reg9341 := True
								reg9343 = reg9341
							} else {
								reg9342 := False
								reg9343 = reg9342
							}
							reg9345 = reg9343
						} else {
							reg9344 := False
							reg9345 = reg9344
						}
						if reg9345 == True {
							reg9346 := MakeNumber(4)
							reg9347 := PrimTail(V1543)
							reg9348 := Call(__e, __defun__shen_4digits_1_6integers, reg9347)
							reg9349 := PrimCons(reg9346, reg9348)
							__e.Return(reg9349)
							return
						} else {
							reg9350 := PrimIsPair(V1543)
							var reg9358 Obj
							if reg9350 == True {
								reg9351 := MakeString("5")
								reg9352 := PrimHead(V1543)
								reg9353 := PrimEqual(reg9351, reg9352)
								var reg9356 Obj
								if reg9353 == True {
									reg9354 := True
									reg9356 = reg9354
								} else {
									reg9355 := False
									reg9356 = reg9355
								}
								reg9358 = reg9356
							} else {
								reg9357 := False
								reg9358 = reg9357
							}
							if reg9358 == True {
								reg9359 := MakeNumber(5)
								reg9360 := PrimTail(V1543)
								reg9361 := Call(__e, __defun__shen_4digits_1_6integers, reg9360)
								reg9362 := PrimCons(reg9359, reg9361)
								__e.Return(reg9362)
								return
							} else {
								reg9363 := PrimIsPair(V1543)
								var reg9371 Obj
								if reg9363 == True {
									reg9364 := MakeString("6")
									reg9365 := PrimHead(V1543)
									reg9366 := PrimEqual(reg9364, reg9365)
									var reg9369 Obj
									if reg9366 == True {
										reg9367 := True
										reg9369 = reg9367
									} else {
										reg9368 := False
										reg9369 = reg9368
									}
									reg9371 = reg9369
								} else {
									reg9370 := False
									reg9371 = reg9370
								}
								if reg9371 == True {
									reg9372 := MakeNumber(6)
									reg9373 := PrimTail(V1543)
									reg9374 := Call(__e, __defun__shen_4digits_1_6integers, reg9373)
									reg9375 := PrimCons(reg9372, reg9374)
									__e.Return(reg9375)
									return
								} else {
									reg9376 := PrimIsPair(V1543)
									var reg9384 Obj
									if reg9376 == True {
										reg9377 := MakeString("7")
										reg9378 := PrimHead(V1543)
										reg9379 := PrimEqual(reg9377, reg9378)
										var reg9382 Obj
										if reg9379 == True {
											reg9380 := True
											reg9382 = reg9380
										} else {
											reg9381 := False
											reg9382 = reg9381
										}
										reg9384 = reg9382
									} else {
										reg9383 := False
										reg9384 = reg9383
									}
									if reg9384 == True {
										reg9385 := MakeNumber(7)
										reg9386 := PrimTail(V1543)
										reg9387 := Call(__e, __defun__shen_4digits_1_6integers, reg9386)
										reg9388 := PrimCons(reg9385, reg9387)
										__e.Return(reg9388)
										return
									} else {
										reg9389 := PrimIsPair(V1543)
										var reg9397 Obj
										if reg9389 == True {
											reg9390 := MakeString("8")
											reg9391 := PrimHead(V1543)
											reg9392 := PrimEqual(reg9390, reg9391)
											var reg9395 Obj
											if reg9392 == True {
												reg9393 := True
												reg9395 = reg9393
											} else {
												reg9394 := False
												reg9395 = reg9394
											}
											reg9397 = reg9395
										} else {
											reg9396 := False
											reg9397 = reg9396
										}
										if reg9397 == True {
											reg9398 := MakeNumber(8)
											reg9399 := PrimTail(V1543)
											reg9400 := Call(__e, __defun__shen_4digits_1_6integers, reg9399)
											reg9401 := PrimCons(reg9398, reg9400)
											__e.Return(reg9401)
											return
										} else {
											reg9402 := PrimIsPair(V1543)
											var reg9410 Obj
											if reg9402 == True {
												reg9403 := MakeString("9")
												reg9404 := PrimHead(V1543)
												reg9405 := PrimEqual(reg9403, reg9404)
												var reg9408 Obj
												if reg9405 == True {
													reg9406 := True
													reg9408 = reg9406
												} else {
													reg9407 := False
													reg9408 = reg9407
												}
												reg9410 = reg9408
											} else {
												reg9409 := False
												reg9410 = reg9409
											}
											if reg9410 == True {
												reg9411 := MakeNumber(9)
												reg9412 := PrimTail(V1543)
												reg9413 := Call(__e, __defun__shen_4digits_1_6integers, reg9412)
												reg9414 := PrimCons(reg9411, reg9413)
												__e.Return(reg9414)
												return
											} else {
												reg9415 := Nil
												__e.Return(reg9415)
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
	__initDefs = append(__initDefs, defType{name: "shen.digits->integers", value: __defun__shen_4digits_1_6integers})

	__defun__shen_4_5sym_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1545 := __args[0]
		_ = V1545
		reg9416 := Call(__e, __defun__shen_4_5alpha_6, V1545)
		Parse__shen_4_5alpha_6 := reg9416
		_ = Parse__shen_4_5alpha_6
		reg9417 := Call(__e, __defun__fail)
		reg9418 := PrimEqual(reg9417, Parse__shen_4_5alpha_6)
		reg9419 := PrimNot(reg9418)
		if reg9419 == True {
			reg9420 := Call(__e, __defun__shen_4_5alphanums_6, Parse__shen_4_5alpha_6)
			Parse__shen_4_5alphanums_6 := reg9420
			_ = Parse__shen_4_5alphanums_6
			reg9421 := Call(__e, __defun__fail)
			reg9422 := PrimEqual(reg9421, Parse__shen_4_5alphanums_6)
			reg9423 := PrimNot(reg9422)
			if reg9423 == True {
				reg9424 := PrimHead(Parse__shen_4_5alphanums_6)
				reg9425 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5alpha_6)
				reg9426 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5alphanums_6)
				reg9427 := Call(__e, __defun___8s, reg9425, reg9426)
				__e.TailApply(__defun__shen_4pair, reg9424, reg9427)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<sym>", value: __defun__shen_4_5sym_6})

	__defun__shen_4_5alphanums_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1547 := __args[0]
		_ = V1547
		reg9431 := Call(__e, __defun__shen_4_5alphanum_6, V1547)
		Parse__shen_4_5alphanum_6 := reg9431
		_ = Parse__shen_4_5alphanum_6
		reg9432 := Call(__e, __defun__fail)
		reg9433 := PrimEqual(reg9432, Parse__shen_4_5alphanum_6)
		reg9434 := PrimNot(reg9433)
		var reg9447 Obj
		if reg9434 == True {
			reg9435 := Call(__e, __defun__shen_4_5alphanums_6, Parse__shen_4_5alphanum_6)
			Parse__shen_4_5alphanums_6 := reg9435
			_ = Parse__shen_4_5alphanums_6
			reg9436 := Call(__e, __defun__fail)
			reg9437 := PrimEqual(reg9436, Parse__shen_4_5alphanums_6)
			reg9438 := PrimNot(reg9437)
			var reg9445 Obj
			if reg9438 == True {
				reg9439 := PrimHead(Parse__shen_4_5alphanums_6)
				reg9440 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5alphanum_6)
				reg9441 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5alphanums_6)
				reg9442 := Call(__e, __defun___8s, reg9440, reg9441)
				reg9443 := Call(__e, __defun__shen_4pair, reg9439, reg9442)
				reg9445 = reg9443
			} else {
				reg9444 := Call(__e, __defun__fail)
				reg9445 = reg9444
			}
			reg9447 = reg9445
		} else {
			reg9446 := Call(__e, __defun__fail)
			reg9447 = reg9446
		}
		YaccParse := reg9447
		_ = YaccParse
		reg9448 := Call(__e, __defun__fail)
		reg9449 := PrimEqual(YaccParse, reg9448)
		if reg9449 == True {
			reg9450 := Call(__e, __defun___5e_6, V1547)
			Parse___5e_6 := reg9450
			_ = Parse___5e_6
			reg9451 := Call(__e, __defun__fail)
			reg9452 := PrimEqual(reg9451, Parse___5e_6)
			reg9453 := PrimNot(reg9452)
			if reg9453 == True {
				reg9454 := PrimHead(Parse___5e_6)
				reg9455 := MakeString("")
				__e.TailApply(__defun__shen_4pair, reg9454, reg9455)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<alphanums>", value: __defun__shen_4_5alphanums_6})

	__defun__shen_4_5alphanum_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1549 := __args[0]
		_ = V1549
		reg9458 := Call(__e, __defun__shen_4_5alpha_6, V1549)
		Parse__shen_4_5alpha_6 := reg9458
		_ = Parse__shen_4_5alpha_6
		reg9459 := Call(__e, __defun__fail)
		reg9460 := PrimEqual(reg9459, Parse__shen_4_5alpha_6)
		reg9461 := PrimNot(reg9460)
		var reg9466 Obj
		if reg9461 == True {
			reg9462 := PrimHead(Parse__shen_4_5alpha_6)
			reg9463 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5alpha_6)
			reg9464 := Call(__e, __defun__shen_4pair, reg9462, reg9463)
			reg9466 = reg9464
		} else {
			reg9465 := Call(__e, __defun__fail)
			reg9466 = reg9465
		}
		YaccParse := reg9466
		_ = YaccParse
		reg9467 := Call(__e, __defun__fail)
		reg9468 := PrimEqual(YaccParse, reg9467)
		if reg9468 == True {
			reg9469 := Call(__e, __defun__shen_4_5num_6, V1549)
			Parse__shen_4_5num_6 := reg9469
			_ = Parse__shen_4_5num_6
			reg9470 := Call(__e, __defun__fail)
			reg9471 := PrimEqual(reg9470, Parse__shen_4_5num_6)
			reg9472 := PrimNot(reg9471)
			if reg9472 == True {
				reg9473 := PrimHead(Parse__shen_4_5num_6)
				reg9474 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5num_6)
				__e.TailApply(__defun__shen_4pair, reg9473, reg9474)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<alphanum>", value: __defun__shen_4_5alphanum_6})

	__defun__shen_4_5num_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1551 := __args[0]
		_ = V1551
		reg9477 := PrimHead(V1551)
		reg9478 := PrimIsPair(reg9477)
		if reg9478 == True {
			reg9479 := Call(__e, __defun__shen_4hdhd, V1551)
			Parse__Char := reg9479
			_ = Parse__Char
			reg9480 := Call(__e, __defun__shen_4numbyte_2, Parse__Char)
			if reg9480 == True {
				reg9481 := Call(__e, __defun__shen_4tlhd, V1551)
				reg9482 := Call(__e, __defun__shen_4hdtl, V1551)
				reg9483 := Call(__e, __defun__shen_4pair, reg9481, reg9482)
				reg9484 := PrimHead(reg9483)
				reg9485 := PrimNumberToString(Parse__Char)
				__e.TailApply(__defun__shen_4pair, reg9484, reg9485)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<num>", value: __defun__shen_4_5num_6})

	__defun__shen_4numbyte_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1557 := __args[0]
		_ = V1557
		reg9489 := MakeNumber(48)
		reg9490 := PrimEqual(reg9489, V1557)
		if reg9490 == True {
			reg9491 := True
			__e.Return(reg9491)
			return
		} else {
			reg9492 := MakeNumber(49)
			reg9493 := PrimEqual(reg9492, V1557)
			if reg9493 == True {
				reg9494 := True
				__e.Return(reg9494)
				return
			} else {
				reg9495 := MakeNumber(50)
				reg9496 := PrimEqual(reg9495, V1557)
				if reg9496 == True {
					reg9497 := True
					__e.Return(reg9497)
					return
				} else {
					reg9498 := MakeNumber(51)
					reg9499 := PrimEqual(reg9498, V1557)
					if reg9499 == True {
						reg9500 := True
						__e.Return(reg9500)
						return
					} else {
						reg9501 := MakeNumber(52)
						reg9502 := PrimEqual(reg9501, V1557)
						if reg9502 == True {
							reg9503 := True
							__e.Return(reg9503)
							return
						} else {
							reg9504 := MakeNumber(53)
							reg9505 := PrimEqual(reg9504, V1557)
							if reg9505 == True {
								reg9506 := True
								__e.Return(reg9506)
								return
							} else {
								reg9507 := MakeNumber(54)
								reg9508 := PrimEqual(reg9507, V1557)
								if reg9508 == True {
									reg9509 := True
									__e.Return(reg9509)
									return
								} else {
									reg9510 := MakeNumber(55)
									reg9511 := PrimEqual(reg9510, V1557)
									if reg9511 == True {
										reg9512 := True
										__e.Return(reg9512)
										return
									} else {
										reg9513 := MakeNumber(56)
										reg9514 := PrimEqual(reg9513, V1557)
										if reg9514 == True {
											reg9515 := True
											__e.Return(reg9515)
											return
										} else {
											reg9516 := MakeNumber(57)
											reg9517 := PrimEqual(reg9516, V1557)
											if reg9517 == True {
												reg9518 := True
												__e.Return(reg9518)
												return
											} else {
												reg9519 := False
												__e.Return(reg9519)
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
	__initDefs = append(__initDefs, defType{name: "shen.numbyte?", value: __defun__shen_4numbyte_2})

	__defun__shen_4_5alpha_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1559 := __args[0]
		_ = V1559
		reg9520 := PrimHead(V1559)
		reg9521 := PrimIsPair(reg9520)
		if reg9521 == True {
			reg9522 := Call(__e, __defun__shen_4hdhd, V1559)
			Parse__Char := reg9522
			_ = Parse__Char
			reg9523 := Call(__e, __defun__shen_4symbol_1code_2, Parse__Char)
			if reg9523 == True {
				reg9524 := Call(__e, __defun__shen_4tlhd, V1559)
				reg9525 := Call(__e, __defun__shen_4hdtl, V1559)
				reg9526 := Call(__e, __defun__shen_4pair, reg9524, reg9525)
				reg9527 := PrimHead(reg9526)
				reg9528 := PrimNumberToString(Parse__Char)
				__e.TailApply(__defun__shen_4pair, reg9527, reg9528)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<alpha>", value: __defun__shen_4_5alpha_6})

	__defun__shen_4symbol_1code_2 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1561 := __args[0]
		_ = V1561
		reg9532 := MakeNumber(126)
		reg9533 := PrimEqual(V1561, reg9532)
		if reg9533 == True {
			reg9534 := True
			__e.Return(reg9534)
			return
		} else {
			reg9535 := MakeNumber(94)
			reg9536 := PrimGreatThan(V1561, reg9535)
			var reg9543 Obj
			if reg9536 == True {
				reg9537 := MakeNumber(123)
				reg9538 := PrimLessThan(V1561, reg9537)
				var reg9541 Obj
				if reg9538 == True {
					reg9539 := True
					reg9541 = reg9539
				} else {
					reg9540 := False
					reg9541 = reg9540
				}
				reg9543 = reg9541
			} else {
				reg9542 := False
				reg9543 = reg9542
			}
			var reg9600 Obj
			if reg9543 == True {
				reg9544 := True
				reg9600 = reg9544
			} else {
				reg9545 := MakeNumber(59)
				reg9546 := PrimGreatThan(V1561, reg9545)
				var reg9553 Obj
				if reg9546 == True {
					reg9547 := MakeNumber(91)
					reg9548 := PrimLessThan(V1561, reg9547)
					var reg9551 Obj
					if reg9548 == True {
						reg9549 := True
						reg9551 = reg9549
					} else {
						reg9550 := False
						reg9551 = reg9550
					}
					reg9553 = reg9551
				} else {
					reg9552 := False
					reg9553 = reg9552
				}
				var reg9596 Obj
				if reg9553 == True {
					reg9554 := True
					reg9596 = reg9554
				} else {
					reg9555 := MakeNumber(41)
					reg9556 := PrimGreatThan(V1561, reg9555)
					var reg9571 Obj
					if reg9556 == True {
						reg9557 := MakeNumber(58)
						reg9558 := PrimLessThan(V1561, reg9557)
						var reg9566 Obj
						if reg9558 == True {
							reg9559 := MakeNumber(44)
							reg9560 := PrimEqual(V1561, reg9559)
							reg9561 := PrimNot(reg9560)
							var reg9564 Obj
							if reg9561 == True {
								reg9562 := True
								reg9564 = reg9562
							} else {
								reg9563 := False
								reg9564 = reg9563
							}
							reg9566 = reg9564
						} else {
							reg9565 := False
							reg9566 = reg9565
						}
						var reg9569 Obj
						if reg9566 == True {
							reg9567 := True
							reg9569 = reg9567
						} else {
							reg9568 := False
							reg9569 = reg9568
						}
						reg9571 = reg9569
					} else {
						reg9570 := False
						reg9571 = reg9570
					}
					var reg9592 Obj
					if reg9571 == True {
						reg9572 := True
						reg9592 = reg9572
					} else {
						reg9573 := MakeNumber(34)
						reg9574 := PrimGreatThan(V1561, reg9573)
						var reg9581 Obj
						if reg9574 == True {
							reg9575 := MakeNumber(40)
							reg9576 := PrimLessThan(V1561, reg9575)
							var reg9579 Obj
							if reg9576 == True {
								reg9577 := True
								reg9579 = reg9577
							} else {
								reg9578 := False
								reg9579 = reg9578
							}
							reg9581 = reg9579
						} else {
							reg9580 := False
							reg9581 = reg9580
						}
						var reg9588 Obj
						if reg9581 == True {
							reg9582 := True
							reg9588 = reg9582
						} else {
							reg9583 := MakeNumber(33)
							reg9584 := PrimEqual(V1561, reg9583)
							var reg9587 Obj
							if reg9584 == True {
								reg9585 := True
								reg9587 = reg9585
							} else {
								reg9586 := False
								reg9587 = reg9586
							}
							reg9588 = reg9587
						}
						var reg9591 Obj
						if reg9588 == True {
							reg9589 := True
							reg9591 = reg9589
						} else {
							reg9590 := False
							reg9591 = reg9590
						}
						reg9592 = reg9591
					}
					var reg9595 Obj
					if reg9592 == True {
						reg9593 := True
						reg9595 = reg9593
					} else {
						reg9594 := False
						reg9595 = reg9594
					}
					reg9596 = reg9595
				}
				var reg9599 Obj
				if reg9596 == True {
					reg9597 := True
					reg9599 = reg9597
				} else {
					reg9598 := False
					reg9599 = reg9598
				}
				reg9600 = reg9599
			}
			if reg9600 == True {
				reg9601 := True
				__e.Return(reg9601)
				return
			} else {
				reg9602 := False
				__e.Return(reg9602)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.symbol-code?", value: __defun__shen_4symbol_1code_2})

	__defun__shen_4_5str_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1563 := __args[0]
		_ = V1563
		reg9603 := Call(__e, __defun__shen_4_5dbq_6, V1563)
		Parse__shen_4_5dbq_6 := reg9603
		_ = Parse__shen_4_5dbq_6
		reg9604 := Call(__e, __defun__fail)
		reg9605 := PrimEqual(reg9604, Parse__shen_4_5dbq_6)
		reg9606 := PrimNot(reg9605)
		if reg9606 == True {
			reg9607 := Call(__e, __defun__shen_4_5strcontents_6, Parse__shen_4_5dbq_6)
			Parse__shen_4_5strcontents_6 := reg9607
			_ = Parse__shen_4_5strcontents_6
			reg9608 := Call(__e, __defun__fail)
			reg9609 := PrimEqual(reg9608, Parse__shen_4_5strcontents_6)
			reg9610 := PrimNot(reg9609)
			if reg9610 == True {
				reg9611 := Call(__e, __defun__shen_4_5dbq_6, Parse__shen_4_5strcontents_6)
				Parse__shen_4_5dbq_6 := reg9611
				_ = Parse__shen_4_5dbq_6
				reg9612 := Call(__e, __defun__fail)
				reg9613 := PrimEqual(reg9612, Parse__shen_4_5dbq_6)
				reg9614 := PrimNot(reg9613)
				if reg9614 == True {
					reg9615 := PrimHead(Parse__shen_4_5dbq_6)
					reg9616 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5strcontents_6)
					__e.TailApply(__defun__shen_4pair, reg9615, reg9616)
					return
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<str>", value: __defun__shen_4_5str_6})

	__defun__shen_4_5dbq_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1565 := __args[0]
		_ = V1565
		reg9621 := PrimHead(V1565)
		reg9622 := PrimIsPair(reg9621)
		if reg9622 == True {
			reg9623 := Call(__e, __defun__shen_4hdhd, V1565)
			Parse__Char := reg9623
			_ = Parse__Char
			reg9624 := MakeNumber(34)
			reg9625 := PrimEqual(Parse__Char, reg9624)
			if reg9625 == True {
				reg9626 := Call(__e, __defun__shen_4tlhd, V1565)
				reg9627 := Call(__e, __defun__shen_4hdtl, V1565)
				reg9628 := Call(__e, __defun__shen_4pair, reg9626, reg9627)
				reg9629 := PrimHead(reg9628)
				__e.TailApply(__defun__shen_4pair, reg9629, Parse__Char)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<dbq>", value: __defun__shen_4_5dbq_6})

	__defun__shen_4_5strcontents_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1567 := __args[0]
		_ = V1567
		reg9633 := Call(__e, __defun__shen_4_5strc_6, V1567)
		Parse__shen_4_5strc_6 := reg9633
		_ = Parse__shen_4_5strc_6
		reg9634 := Call(__e, __defun__fail)
		reg9635 := PrimEqual(reg9634, Parse__shen_4_5strc_6)
		reg9636 := PrimNot(reg9635)
		var reg9649 Obj
		if reg9636 == True {
			reg9637 := Call(__e, __defun__shen_4_5strcontents_6, Parse__shen_4_5strc_6)
			Parse__shen_4_5strcontents_6 := reg9637
			_ = Parse__shen_4_5strcontents_6
			reg9638 := Call(__e, __defun__fail)
			reg9639 := PrimEqual(reg9638, Parse__shen_4_5strcontents_6)
			reg9640 := PrimNot(reg9639)
			var reg9647 Obj
			if reg9640 == True {
				reg9641 := PrimHead(Parse__shen_4_5strcontents_6)
				reg9642 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5strc_6)
				reg9643 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5strcontents_6)
				reg9644 := PrimCons(reg9642, reg9643)
				reg9645 := Call(__e, __defun__shen_4pair, reg9641, reg9644)
				reg9647 = reg9645
			} else {
				reg9646 := Call(__e, __defun__fail)
				reg9647 = reg9646
			}
			reg9649 = reg9647
		} else {
			reg9648 := Call(__e, __defun__fail)
			reg9649 = reg9648
		}
		YaccParse := reg9649
		_ = YaccParse
		reg9650 := Call(__e, __defun__fail)
		reg9651 := PrimEqual(YaccParse, reg9650)
		if reg9651 == True {
			reg9652 := Call(__e, __defun___5e_6, V1567)
			Parse___5e_6 := reg9652
			_ = Parse___5e_6
			reg9653 := Call(__e, __defun__fail)
			reg9654 := PrimEqual(reg9653, Parse___5e_6)
			reg9655 := PrimNot(reg9654)
			if reg9655 == True {
				reg9656 := PrimHead(Parse___5e_6)
				reg9657 := Nil
				__e.TailApply(__defun__shen_4pair, reg9656, reg9657)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<strcontents>", value: __defun__shen_4_5strcontents_6})

	__defun__shen_4_5byte_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1569 := __args[0]
		_ = V1569
		reg9660 := PrimHead(V1569)
		reg9661 := PrimIsPair(reg9660)
		if reg9661 == True {
			reg9662 := Call(__e, __defun__shen_4hdhd, V1569)
			Parse__Char := reg9662
			_ = Parse__Char
			reg9663 := Call(__e, __defun__shen_4tlhd, V1569)
			reg9664 := Call(__e, __defun__shen_4hdtl, V1569)
			reg9665 := Call(__e, __defun__shen_4pair, reg9663, reg9664)
			reg9666 := PrimHead(reg9665)
			reg9667 := PrimNumberToString(Parse__Char)
			__e.TailApply(__defun__shen_4pair, reg9666, reg9667)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<byte>", value: __defun__shen_4_5byte_6})

	__defun__shen_4_5strc_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1571 := __args[0]
		_ = V1571
		reg9670 := PrimHead(V1571)
		reg9671 := PrimIsPair(reg9670)
		if reg9671 == True {
			reg9672 := Call(__e, __defun__shen_4hdhd, V1571)
			Parse__Char := reg9672
			_ = Parse__Char
			reg9673 := MakeNumber(34)
			reg9674 := PrimEqual(Parse__Char, reg9673)
			reg9675 := PrimNot(reg9674)
			if reg9675 == True {
				reg9676 := Call(__e, __defun__shen_4tlhd, V1571)
				reg9677 := Call(__e, __defun__shen_4hdtl, V1571)
				reg9678 := Call(__e, __defun__shen_4pair, reg9676, reg9677)
				reg9679 := PrimHead(reg9678)
				reg9680 := PrimNumberToString(Parse__Char)
				__e.TailApply(__defun__shen_4pair, reg9679, reg9680)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<strc>", value: __defun__shen_4_5strc_6})

	__defun__shen_4_5number_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1573 := __args[0]
		_ = V1573
		reg9684 := Call(__e, __defun__shen_4_5minus_6, V1573)
		Parse__shen_4_5minus_6 := reg9684
		_ = Parse__shen_4_5minus_6
		reg9685 := Call(__e, __defun__fail)
		reg9686 := PrimEqual(reg9685, Parse__shen_4_5minus_6)
		reg9687 := PrimNot(reg9686)
		var reg9700 Obj
		if reg9687 == True {
			reg9688 := Call(__e, __defun__shen_4_5number_6, Parse__shen_4_5minus_6)
			Parse__shen_4_5number_6 := reg9688
			_ = Parse__shen_4_5number_6
			reg9689 := Call(__e, __defun__fail)
			reg9690 := PrimEqual(reg9689, Parse__shen_4_5number_6)
			reg9691 := PrimNot(reg9690)
			var reg9698 Obj
			if reg9691 == True {
				reg9692 := PrimHead(Parse__shen_4_5number_6)
				reg9693 := MakeNumber(0)
				reg9694 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5number_6)
				reg9695 := PrimNumberSubtract(reg9693, reg9694)
				reg9696 := Call(__e, __defun__shen_4pair, reg9692, reg9695)
				reg9698 = reg9696
			} else {
				reg9697 := Call(__e, __defun__fail)
				reg9698 = reg9697
			}
			reg9700 = reg9698
		} else {
			reg9699 := Call(__e, __defun__fail)
			reg9700 = reg9699
		}
		YaccParse := reg9700
		_ = YaccParse
		reg9701 := Call(__e, __defun__fail)
		reg9702 := PrimEqual(YaccParse, reg9701)
		if reg9702 == True {
			reg9703 := Call(__e, __defun__shen_4_5plus_6, V1573)
			Parse__shen_4_5plus_6 := reg9703
			_ = Parse__shen_4_5plus_6
			reg9704 := Call(__e, __defun__fail)
			reg9705 := PrimEqual(reg9704, Parse__shen_4_5plus_6)
			reg9706 := PrimNot(reg9705)
			var reg9717 Obj
			if reg9706 == True {
				reg9707 := Call(__e, __defun__shen_4_5number_6, Parse__shen_4_5plus_6)
				Parse__shen_4_5number_6 := reg9707
				_ = Parse__shen_4_5number_6
				reg9708 := Call(__e, __defun__fail)
				reg9709 := PrimEqual(reg9708, Parse__shen_4_5number_6)
				reg9710 := PrimNot(reg9709)
				var reg9715 Obj
				if reg9710 == True {
					reg9711 := PrimHead(Parse__shen_4_5number_6)
					reg9712 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5number_6)
					reg9713 := Call(__e, __defun__shen_4pair, reg9711, reg9712)
					reg9715 = reg9713
				} else {
					reg9714 := Call(__e, __defun__fail)
					reg9715 = reg9714
				}
				reg9717 = reg9715
			} else {
				reg9716 := Call(__e, __defun__fail)
				reg9717 = reg9716
			}
			YaccParse := reg9717
			_ = YaccParse
			reg9718 := Call(__e, __defun__fail)
			reg9719 := PrimEqual(YaccParse, reg9718)
			if reg9719 == True {
				reg9720 := Call(__e, __defun__shen_4_5predigits_6, V1573)
				Parse__shen_4_5predigits_6 := reg9720
				_ = Parse__shen_4_5predigits_6
				reg9721 := Call(__e, __defun__fail)
				reg9722 := PrimEqual(reg9721, Parse__shen_4_5predigits_6)
				reg9723 := PrimNot(reg9722)
				var reg9763 Obj
				if reg9723 == True {
					reg9724 := Call(__e, __defun__shen_4_5stop_6, Parse__shen_4_5predigits_6)
					Parse__shen_4_5stop_6 := reg9724
					_ = Parse__shen_4_5stop_6
					reg9725 := Call(__e, __defun__fail)
					reg9726 := PrimEqual(reg9725, Parse__shen_4_5stop_6)
					reg9727 := PrimNot(reg9726)
					var reg9761 Obj
					if reg9727 == True {
						reg9728 := Call(__e, __defun__shen_4_5postdigits_6, Parse__shen_4_5stop_6)
						Parse__shen_4_5postdigits_6 := reg9728
						_ = Parse__shen_4_5postdigits_6
						reg9729 := Call(__e, __defun__fail)
						reg9730 := PrimEqual(reg9729, Parse__shen_4_5postdigits_6)
						reg9731 := PrimNot(reg9730)
						var reg9759 Obj
						if reg9731 == True {
							reg9732 := Call(__e, __defun__shen_4_5E_6, Parse__shen_4_5postdigits_6)
							Parse__shen_4_5E_6 := reg9732
							_ = Parse__shen_4_5E_6
							reg9733 := Call(__e, __defun__fail)
							reg9734 := PrimEqual(reg9733, Parse__shen_4_5E_6)
							reg9735 := PrimNot(reg9734)
							var reg9757 Obj
							if reg9735 == True {
								reg9736 := Call(__e, __defun__shen_4_5log10_6, Parse__shen_4_5E_6)
								Parse__shen_4_5log10_6 := reg9736
								_ = Parse__shen_4_5log10_6
								reg9737 := Call(__e, __defun__fail)
								reg9738 := PrimEqual(reg9737, Parse__shen_4_5log10_6)
								reg9739 := PrimNot(reg9738)
								var reg9755 Obj
								if reg9739 == True {
									reg9740 := PrimHead(Parse__shen_4_5log10_6)
									reg9741 := MakeNumber(10)
									reg9742 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5log10_6)
									reg9743 := Call(__e, __defun__shen_4expt, reg9741, reg9742)
									reg9744 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5predigits_6)
									reg9745 := Call(__e, __defun__reverse, reg9744)
									reg9746 := MakeNumber(0)
									reg9747 := Call(__e, __defun__shen_4pre, reg9745, reg9746)
									reg9748 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5postdigits_6)
									reg9749 := MakeNumber(1)
									reg9750 := Call(__e, __defun__shen_4post, reg9748, reg9749)
									reg9751 := PrimNumberAdd(reg9747, reg9750)
									reg9752 := PrimNumberMultiply(reg9743, reg9751)
									reg9753 := Call(__e, __defun__shen_4pair, reg9740, reg9752)
									reg9755 = reg9753
								} else {
									reg9754 := Call(__e, __defun__fail)
									reg9755 = reg9754
								}
								reg9757 = reg9755
							} else {
								reg9756 := Call(__e, __defun__fail)
								reg9757 = reg9756
							}
							reg9759 = reg9757
						} else {
							reg9758 := Call(__e, __defun__fail)
							reg9759 = reg9758
						}
						reg9761 = reg9759
					} else {
						reg9760 := Call(__e, __defun__fail)
						reg9761 = reg9760
					}
					reg9763 = reg9761
				} else {
					reg9762 := Call(__e, __defun__fail)
					reg9763 = reg9762
				}
				YaccParse := reg9763
				_ = YaccParse
				reg9764 := Call(__e, __defun__fail)
				reg9765 := PrimEqual(YaccParse, reg9764)
				if reg9765 == True {
					reg9766 := Call(__e, __defun__shen_4_5digits_6, V1573)
					Parse__shen_4_5digits_6 := reg9766
					_ = Parse__shen_4_5digits_6
					reg9767 := Call(__e, __defun__fail)
					reg9768 := PrimEqual(reg9767, Parse__shen_4_5digits_6)
					reg9769 := PrimNot(reg9768)
					var reg9793 Obj
					if reg9769 == True {
						reg9770 := Call(__e, __defun__shen_4_5E_6, Parse__shen_4_5digits_6)
						Parse__shen_4_5E_6 := reg9770
						_ = Parse__shen_4_5E_6
						reg9771 := Call(__e, __defun__fail)
						reg9772 := PrimEqual(reg9771, Parse__shen_4_5E_6)
						reg9773 := PrimNot(reg9772)
						var reg9791 Obj
						if reg9773 == True {
							reg9774 := Call(__e, __defun__shen_4_5log10_6, Parse__shen_4_5E_6)
							Parse__shen_4_5log10_6 := reg9774
							_ = Parse__shen_4_5log10_6
							reg9775 := Call(__e, __defun__fail)
							reg9776 := PrimEqual(reg9775, Parse__shen_4_5log10_6)
							reg9777 := PrimNot(reg9776)
							var reg9789 Obj
							if reg9777 == True {
								reg9778 := PrimHead(Parse__shen_4_5log10_6)
								reg9779 := MakeNumber(10)
								reg9780 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5log10_6)
								reg9781 := Call(__e, __defun__shen_4expt, reg9779, reg9780)
								reg9782 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digits_6)
								reg9783 := Call(__e, __defun__reverse, reg9782)
								reg9784 := MakeNumber(0)
								reg9785 := Call(__e, __defun__shen_4pre, reg9783, reg9784)
								reg9786 := PrimNumberMultiply(reg9781, reg9785)
								reg9787 := Call(__e, __defun__shen_4pair, reg9778, reg9786)
								reg9789 = reg9787
							} else {
								reg9788 := Call(__e, __defun__fail)
								reg9789 = reg9788
							}
							reg9791 = reg9789
						} else {
							reg9790 := Call(__e, __defun__fail)
							reg9791 = reg9790
						}
						reg9793 = reg9791
					} else {
						reg9792 := Call(__e, __defun__fail)
						reg9793 = reg9792
					}
					YaccParse := reg9793
					_ = YaccParse
					reg9794 := Call(__e, __defun__fail)
					reg9795 := PrimEqual(YaccParse, reg9794)
					if reg9795 == True {
						reg9796 := Call(__e, __defun__shen_4_5predigits_6, V1573)
						Parse__shen_4_5predigits_6 := reg9796
						_ = Parse__shen_4_5predigits_6
						reg9797 := Call(__e, __defun__fail)
						reg9798 := PrimEqual(reg9797, Parse__shen_4_5predigits_6)
						reg9799 := PrimNot(reg9798)
						var reg9823 Obj
						if reg9799 == True {
							reg9800 := Call(__e, __defun__shen_4_5stop_6, Parse__shen_4_5predigits_6)
							Parse__shen_4_5stop_6 := reg9800
							_ = Parse__shen_4_5stop_6
							reg9801 := Call(__e, __defun__fail)
							reg9802 := PrimEqual(reg9801, Parse__shen_4_5stop_6)
							reg9803 := PrimNot(reg9802)
							var reg9821 Obj
							if reg9803 == True {
								reg9804 := Call(__e, __defun__shen_4_5postdigits_6, Parse__shen_4_5stop_6)
								Parse__shen_4_5postdigits_6 := reg9804
								_ = Parse__shen_4_5postdigits_6
								reg9805 := Call(__e, __defun__fail)
								reg9806 := PrimEqual(reg9805, Parse__shen_4_5postdigits_6)
								reg9807 := PrimNot(reg9806)
								var reg9819 Obj
								if reg9807 == True {
									reg9808 := PrimHead(Parse__shen_4_5postdigits_6)
									reg9809 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5predigits_6)
									reg9810 := Call(__e, __defun__reverse, reg9809)
									reg9811 := MakeNumber(0)
									reg9812 := Call(__e, __defun__shen_4pre, reg9810, reg9811)
									reg9813 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5postdigits_6)
									reg9814 := MakeNumber(1)
									reg9815 := Call(__e, __defun__shen_4post, reg9813, reg9814)
									reg9816 := PrimNumberAdd(reg9812, reg9815)
									reg9817 := Call(__e, __defun__shen_4pair, reg9808, reg9816)
									reg9819 = reg9817
								} else {
									reg9818 := Call(__e, __defun__fail)
									reg9819 = reg9818
								}
								reg9821 = reg9819
							} else {
								reg9820 := Call(__e, __defun__fail)
								reg9821 = reg9820
							}
							reg9823 = reg9821
						} else {
							reg9822 := Call(__e, __defun__fail)
							reg9823 = reg9822
						}
						YaccParse := reg9823
						_ = YaccParse
						reg9824 := Call(__e, __defun__fail)
						reg9825 := PrimEqual(YaccParse, reg9824)
						if reg9825 == True {
							reg9826 := Call(__e, __defun__shen_4_5digits_6, V1573)
							Parse__shen_4_5digits_6 := reg9826
							_ = Parse__shen_4_5digits_6
							reg9827 := Call(__e, __defun__fail)
							reg9828 := PrimEqual(reg9827, Parse__shen_4_5digits_6)
							reg9829 := PrimNot(reg9828)
							if reg9829 == True {
								reg9830 := PrimHead(Parse__shen_4_5digits_6)
								reg9831 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digits_6)
								reg9832 := Call(__e, __defun__reverse, reg9831)
								reg9833 := MakeNumber(0)
								reg9834 := Call(__e, __defun__shen_4pre, reg9832, reg9833)
								__e.TailApply(__defun__shen_4pair, reg9830, reg9834)
								return
							} else {
								__e.TailApply(__defun__fail)
								return
							}
						} else {
							__e.Return(YaccParse)
							return
						}
					} else {
						__e.Return(YaccParse)
						return
					}
				} else {
					__e.Return(YaccParse)
					return
				}
			} else {
				__e.Return(YaccParse)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<number>", value: __defun__shen_4_5number_6})

	__defun__shen_4_5E_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1576 := __args[0]
		_ = V1576
		reg9837 := PrimHead(V1576)
		reg9838 := PrimIsPair(reg9837)
		var reg9846 Obj
		if reg9838 == True {
			reg9839 := MakeNumber(101)
			reg9840 := Call(__e, __defun__shen_4hdhd, V1576)
			reg9841 := PrimEqual(reg9839, reg9840)
			var reg9844 Obj
			if reg9841 == True {
				reg9842 := True
				reg9844 = reg9842
			} else {
				reg9843 := False
				reg9844 = reg9843
			}
			reg9846 = reg9844
		} else {
			reg9845 := False
			reg9846 = reg9845
		}
		if reg9846 == True {
			reg9847 := Call(__e, __defun__shen_4tlhd, V1576)
			reg9848 := Call(__e, __defun__shen_4hdtl, V1576)
			reg9849 := Call(__e, __defun__shen_4pair, reg9847, reg9848)
			NewStream1574 := reg9849
			_ = NewStream1574
			reg9850 := PrimHead(NewStream1574)
			reg9851 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg9850, reg9851)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<E>", value: __defun__shen_4_5E_6})

	__defun__shen_4_5log10_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1578 := __args[0]
		_ = V1578
		reg9854 := Call(__e, __defun__shen_4_5minus_6, V1578)
		Parse__shen_4_5minus_6 := reg9854
		_ = Parse__shen_4_5minus_6
		reg9855 := Call(__e, __defun__fail)
		reg9856 := PrimEqual(reg9855, Parse__shen_4_5minus_6)
		reg9857 := PrimNot(reg9856)
		var reg9873 Obj
		if reg9857 == True {
			reg9858 := Call(__e, __defun__shen_4_5digits_6, Parse__shen_4_5minus_6)
			Parse__shen_4_5digits_6 := reg9858
			_ = Parse__shen_4_5digits_6
			reg9859 := Call(__e, __defun__fail)
			reg9860 := PrimEqual(reg9859, Parse__shen_4_5digits_6)
			reg9861 := PrimNot(reg9860)
			var reg9871 Obj
			if reg9861 == True {
				reg9862 := PrimHead(Parse__shen_4_5digits_6)
				reg9863 := MakeNumber(0)
				reg9864 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digits_6)
				reg9865 := Call(__e, __defun__reverse, reg9864)
				reg9866 := MakeNumber(0)
				reg9867 := Call(__e, __defun__shen_4pre, reg9865, reg9866)
				reg9868 := PrimNumberSubtract(reg9863, reg9867)
				reg9869 := Call(__e, __defun__shen_4pair, reg9862, reg9868)
				reg9871 = reg9869
			} else {
				reg9870 := Call(__e, __defun__fail)
				reg9871 = reg9870
			}
			reg9873 = reg9871
		} else {
			reg9872 := Call(__e, __defun__fail)
			reg9873 = reg9872
		}
		YaccParse := reg9873
		_ = YaccParse
		reg9874 := Call(__e, __defun__fail)
		reg9875 := PrimEqual(YaccParse, reg9874)
		if reg9875 == True {
			reg9876 := Call(__e, __defun__shen_4_5digits_6, V1578)
			Parse__shen_4_5digits_6 := reg9876
			_ = Parse__shen_4_5digits_6
			reg9877 := Call(__e, __defun__fail)
			reg9878 := PrimEqual(reg9877, Parse__shen_4_5digits_6)
			reg9879 := PrimNot(reg9878)
			if reg9879 == True {
				reg9880 := PrimHead(Parse__shen_4_5digits_6)
				reg9881 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digits_6)
				reg9882 := Call(__e, __defun__reverse, reg9881)
				reg9883 := MakeNumber(0)
				reg9884 := Call(__e, __defun__shen_4pre, reg9882, reg9883)
				__e.TailApply(__defun__shen_4pair, reg9880, reg9884)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<log10>", value: __defun__shen_4_5log10_6})

	__defun__shen_4_5plus_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1580 := __args[0]
		_ = V1580
		reg9887 := PrimHead(V1580)
		reg9888 := PrimIsPair(reg9887)
		if reg9888 == True {
			reg9889 := Call(__e, __defun__shen_4hdhd, V1580)
			Parse__Char := reg9889
			_ = Parse__Char
			reg9890 := MakeNumber(43)
			reg9891 := PrimEqual(Parse__Char, reg9890)
			if reg9891 == True {
				reg9892 := Call(__e, __defun__shen_4tlhd, V1580)
				reg9893 := Call(__e, __defun__shen_4hdtl, V1580)
				reg9894 := Call(__e, __defun__shen_4pair, reg9892, reg9893)
				reg9895 := PrimHead(reg9894)
				__e.TailApply(__defun__shen_4pair, reg9895, Parse__Char)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<plus>", value: __defun__shen_4_5plus_6})

	__defun__shen_4_5stop_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1582 := __args[0]
		_ = V1582
		reg9899 := PrimHead(V1582)
		reg9900 := PrimIsPair(reg9899)
		if reg9900 == True {
			reg9901 := Call(__e, __defun__shen_4hdhd, V1582)
			Parse__Char := reg9901
			_ = Parse__Char
			reg9902 := MakeNumber(46)
			reg9903 := PrimEqual(Parse__Char, reg9902)
			if reg9903 == True {
				reg9904 := Call(__e, __defun__shen_4tlhd, V1582)
				reg9905 := Call(__e, __defun__shen_4hdtl, V1582)
				reg9906 := Call(__e, __defun__shen_4pair, reg9904, reg9905)
				reg9907 := PrimHead(reg9906)
				__e.TailApply(__defun__shen_4pair, reg9907, Parse__Char)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<stop>", value: __defun__shen_4_5stop_6})

	__defun__shen_4_5predigits_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1584 := __args[0]
		_ = V1584
		reg9911 := Call(__e, __defun__shen_4_5digits_6, V1584)
		Parse__shen_4_5digits_6 := reg9911
		_ = Parse__shen_4_5digits_6
		reg9912 := Call(__e, __defun__fail)
		reg9913 := PrimEqual(reg9912, Parse__shen_4_5digits_6)
		reg9914 := PrimNot(reg9913)
		var reg9919 Obj
		if reg9914 == True {
			reg9915 := PrimHead(Parse__shen_4_5digits_6)
			reg9916 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digits_6)
			reg9917 := Call(__e, __defun__shen_4pair, reg9915, reg9916)
			reg9919 = reg9917
		} else {
			reg9918 := Call(__e, __defun__fail)
			reg9919 = reg9918
		}
		YaccParse := reg9919
		_ = YaccParse
		reg9920 := Call(__e, __defun__fail)
		reg9921 := PrimEqual(YaccParse, reg9920)
		if reg9921 == True {
			reg9922 := Call(__e, __defun___5e_6, V1584)
			Parse___5e_6 := reg9922
			_ = Parse___5e_6
			reg9923 := Call(__e, __defun__fail)
			reg9924 := PrimEqual(reg9923, Parse___5e_6)
			reg9925 := PrimNot(reg9924)
			if reg9925 == True {
				reg9926 := PrimHead(Parse___5e_6)
				reg9927 := Nil
				__e.TailApply(__defun__shen_4pair, reg9926, reg9927)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<predigits>", value: __defun__shen_4_5predigits_6})

	__defun__shen_4_5postdigits_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1586 := __args[0]
		_ = V1586
		reg9930 := Call(__e, __defun__shen_4_5digits_6, V1586)
		Parse__shen_4_5digits_6 := reg9930
		_ = Parse__shen_4_5digits_6
		reg9931 := Call(__e, __defun__fail)
		reg9932 := PrimEqual(reg9931, Parse__shen_4_5digits_6)
		reg9933 := PrimNot(reg9932)
		if reg9933 == True {
			reg9934 := PrimHead(Parse__shen_4_5digits_6)
			reg9935 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digits_6)
			__e.TailApply(__defun__shen_4pair, reg9934, reg9935)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<postdigits>", value: __defun__shen_4_5postdigits_6})

	__defun__shen_4_5digits_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1588 := __args[0]
		_ = V1588
		reg9938 := Call(__e, __defun__shen_4_5digit_6, V1588)
		Parse__shen_4_5digit_6 := reg9938
		_ = Parse__shen_4_5digit_6
		reg9939 := Call(__e, __defun__fail)
		reg9940 := PrimEqual(reg9939, Parse__shen_4_5digit_6)
		reg9941 := PrimNot(reg9940)
		var reg9954 Obj
		if reg9941 == True {
			reg9942 := Call(__e, __defun__shen_4_5digits_6, Parse__shen_4_5digit_6)
			Parse__shen_4_5digits_6 := reg9942
			_ = Parse__shen_4_5digits_6
			reg9943 := Call(__e, __defun__fail)
			reg9944 := PrimEqual(reg9943, Parse__shen_4_5digits_6)
			reg9945 := PrimNot(reg9944)
			var reg9952 Obj
			if reg9945 == True {
				reg9946 := PrimHead(Parse__shen_4_5digits_6)
				reg9947 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digit_6)
				reg9948 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digits_6)
				reg9949 := PrimCons(reg9947, reg9948)
				reg9950 := Call(__e, __defun__shen_4pair, reg9946, reg9949)
				reg9952 = reg9950
			} else {
				reg9951 := Call(__e, __defun__fail)
				reg9952 = reg9951
			}
			reg9954 = reg9952
		} else {
			reg9953 := Call(__e, __defun__fail)
			reg9954 = reg9953
		}
		YaccParse := reg9954
		_ = YaccParse
		reg9955 := Call(__e, __defun__fail)
		reg9956 := PrimEqual(YaccParse, reg9955)
		if reg9956 == True {
			reg9957 := Call(__e, __defun__shen_4_5digit_6, V1588)
			Parse__shen_4_5digit_6 := reg9957
			_ = Parse__shen_4_5digit_6
			reg9958 := Call(__e, __defun__fail)
			reg9959 := PrimEqual(reg9958, Parse__shen_4_5digit_6)
			reg9960 := PrimNot(reg9959)
			if reg9960 == True {
				reg9961 := PrimHead(Parse__shen_4_5digit_6)
				reg9962 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5digit_6)
				reg9963 := Nil
				reg9964 := PrimCons(reg9962, reg9963)
				__e.TailApply(__defun__shen_4pair, reg9961, reg9964)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<digits>", value: __defun__shen_4_5digits_6})

	__defun__shen_4_5digit_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1590 := __args[0]
		_ = V1590
		reg9967 := PrimHead(V1590)
		reg9968 := PrimIsPair(reg9967)
		if reg9968 == True {
			reg9969 := Call(__e, __defun__shen_4hdhd, V1590)
			Parse__X := reg9969
			_ = Parse__X
			reg9970 := Call(__e, __defun__shen_4numbyte_2, Parse__X)
			if reg9970 == True {
				reg9971 := Call(__e, __defun__shen_4tlhd, V1590)
				reg9972 := Call(__e, __defun__shen_4hdtl, V1590)
				reg9973 := Call(__e, __defun__shen_4pair, reg9971, reg9972)
				reg9974 := PrimHead(reg9973)
				reg9975 := Call(__e, __defun__shen_4byte_1_6digit, Parse__X)
				__e.TailApply(__defun__shen_4pair, reg9974, reg9975)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<digit>", value: __defun__shen_4_5digit_6})

	__defun__shen_4byte_1_6digit = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1592 := __args[0]
		_ = V1592
		reg9979 := MakeNumber(48)
		reg9980 := PrimEqual(reg9979, V1592)
		if reg9980 == True {
			reg9981 := MakeNumber(0)
			__e.Return(reg9981)
			return
		} else {
			reg9982 := MakeNumber(49)
			reg9983 := PrimEqual(reg9982, V1592)
			if reg9983 == True {
				reg9984 := MakeNumber(1)
				__e.Return(reg9984)
				return
			} else {
				reg9985 := MakeNumber(50)
				reg9986 := PrimEqual(reg9985, V1592)
				if reg9986 == True {
					reg9987 := MakeNumber(2)
					__e.Return(reg9987)
					return
				} else {
					reg9988 := MakeNumber(51)
					reg9989 := PrimEqual(reg9988, V1592)
					if reg9989 == True {
						reg9990 := MakeNumber(3)
						__e.Return(reg9990)
						return
					} else {
						reg9991 := MakeNumber(52)
						reg9992 := PrimEqual(reg9991, V1592)
						if reg9992 == True {
							reg9993 := MakeNumber(4)
							__e.Return(reg9993)
							return
						} else {
							reg9994 := MakeNumber(53)
							reg9995 := PrimEqual(reg9994, V1592)
							if reg9995 == True {
								reg9996 := MakeNumber(5)
								__e.Return(reg9996)
								return
							} else {
								reg9997 := MakeNumber(54)
								reg9998 := PrimEqual(reg9997, V1592)
								if reg9998 == True {
									reg9999 := MakeNumber(6)
									__e.Return(reg9999)
									return
								} else {
									reg10000 := MakeNumber(55)
									reg10001 := PrimEqual(reg10000, V1592)
									if reg10001 == True {
										reg10002 := MakeNumber(7)
										__e.Return(reg10002)
										return
									} else {
										reg10003 := MakeNumber(56)
										reg10004 := PrimEqual(reg10003, V1592)
										if reg10004 == True {
											reg10005 := MakeNumber(8)
											__e.Return(reg10005)
											return
										} else {
											reg10006 := MakeNumber(57)
											reg10007 := PrimEqual(reg10006, V1592)
											if reg10007 == True {
												reg10008 := MakeNumber(9)
												__e.Return(reg10008)
												return
											} else {
												reg10009 := MakeSymbol("shen.byte->digit")
												__e.TailApply(__defun__shen_4f__error, reg10009)
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
	__initDefs = append(__initDefs, defType{name: "shen.byte->digit", value: __defun__shen_4byte_1_6digit})

	__defun__shen_4pre = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1597 := __args[0]
		_ = V1597
		V1598 := __args[1]
		_ = V1598
		reg10011 := Nil
		reg10012 := PrimEqual(reg10011, V1597)
		if reg10012 == True {
			reg10013 := MakeNumber(0)
			__e.Return(reg10013)
			return
		} else {
			reg10014 := PrimIsPair(V1597)
			if reg10014 == True {
				reg10015 := MakeNumber(10)
				reg10016 := Call(__e, __defun__shen_4expt, reg10015, V1598)
				reg10017 := PrimHead(V1597)
				reg10018 := PrimNumberMultiply(reg10016, reg10017)
				reg10019 := PrimTail(V1597)
				reg10020 := MakeNumber(1)
				reg10021 := PrimNumberAdd(V1598, reg10020)
				reg10022 := Call(__e, __defun__shen_4pre, reg10019, reg10021)
				reg10023 := PrimNumberAdd(reg10018, reg10022)
				__e.Return(reg10023)
				return
			} else {
				reg10024 := MakeSymbol("shen.pre")
				__e.TailApply(__defun__shen_4f__error, reg10024)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.pre", value: __defun__shen_4pre})

	__defun__shen_4post = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1603 := __args[0]
		_ = V1603
		V1604 := __args[1]
		_ = V1604
		reg10026 := Nil
		reg10027 := PrimEqual(reg10026, V1603)
		if reg10027 == True {
			reg10028 := MakeNumber(0)
			__e.Return(reg10028)
			return
		} else {
			reg10029 := PrimIsPair(V1603)
			if reg10029 == True {
				reg10030 := MakeNumber(10)
				reg10031 := MakeNumber(0)
				reg10032 := PrimNumberSubtract(reg10031, V1604)
				reg10033 := Call(__e, __defun__shen_4expt, reg10030, reg10032)
				reg10034 := PrimHead(V1603)
				reg10035 := PrimNumberMultiply(reg10033, reg10034)
				reg10036 := PrimTail(V1603)
				reg10037 := MakeNumber(1)
				reg10038 := PrimNumberAdd(V1604, reg10037)
				reg10039 := Call(__e, __defun__shen_4post, reg10036, reg10038)
				reg10040 := PrimNumberAdd(reg10035, reg10039)
				__e.Return(reg10040)
				return
			} else {
				reg10041 := MakeSymbol("shen.post")
				__e.TailApply(__defun__shen_4f__error, reg10041)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.post", value: __defun__shen_4post})

	__defun__shen_4expt = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1609 := __args[0]
		_ = V1609
		V1610 := __args[1]
		_ = V1610
		reg10043 := MakeNumber(0)
		reg10044 := PrimEqual(reg10043, V1610)
		if reg10044 == True {
			reg10045 := MakeNumber(1)
			__e.Return(reg10045)
			return
		} else {
			reg10046 := MakeNumber(0)
			reg10047 := PrimGreatThan(V1610, reg10046)
			if reg10047 == True {
				reg10048 := MakeNumber(1)
				reg10049 := PrimNumberSubtract(V1610, reg10048)
				reg10050 := Call(__e, __defun__shen_4expt, V1609, reg10049)
				reg10051 := PrimNumberMultiply(V1609, reg10050)
				__e.Return(reg10051)
				return
			} else {
				reg10052 := MakeNumber(1)
				reg10053 := MakeNumber(1)
				reg10054 := PrimNumberAdd(V1610, reg10053)
				reg10055 := Call(__e, __defun__shen_4expt, V1609, reg10054)
				reg10056 := PrimNumberDivide(reg10055, V1609)
				reg10057 := PrimNumberMultiply(reg10052, reg10056)
				__e.Return(reg10057)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.expt", value: __defun__shen_4expt})

	__defun__shen_4_5st__input1_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1612 := __args[0]
		_ = V1612
		reg10058 := Call(__e, __defun__shen_4_5st__input_6, V1612)
		Parse__shen_4_5st__input_6 := reg10058
		_ = Parse__shen_4_5st__input_6
		reg10059 := Call(__e, __defun__fail)
		reg10060 := PrimEqual(reg10059, Parse__shen_4_5st__input_6)
		reg10061 := PrimNot(reg10060)
		if reg10061 == True {
			reg10062 := PrimHead(Parse__shen_4_5st__input_6)
			reg10063 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
			__e.TailApply(__defun__shen_4pair, reg10062, reg10063)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<st_input1>", value: __defun__shen_4_5st__input1_6})

	__defun__shen_4_5st__input2_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1614 := __args[0]
		_ = V1614
		reg10066 := Call(__e, __defun__shen_4_5st__input_6, V1614)
		Parse__shen_4_5st__input_6 := reg10066
		_ = Parse__shen_4_5st__input_6
		reg10067 := Call(__e, __defun__fail)
		reg10068 := PrimEqual(reg10067, Parse__shen_4_5st__input_6)
		reg10069 := PrimNot(reg10068)
		if reg10069 == True {
			reg10070 := PrimHead(Parse__shen_4_5st__input_6)
			reg10071 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5st__input_6)
			__e.TailApply(__defun__shen_4pair, reg10070, reg10071)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<st_input2>", value: __defun__shen_4_5st__input2_6})

	__defun__shen_4_5comment_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1616 := __args[0]
		_ = V1616
		reg10074 := Call(__e, __defun__shen_4_5singleline_6, V1616)
		Parse__shen_4_5singleline_6 := reg10074
		_ = Parse__shen_4_5singleline_6
		reg10075 := Call(__e, __defun__fail)
		reg10076 := PrimEqual(reg10075, Parse__shen_4_5singleline_6)
		reg10077 := PrimNot(reg10076)
		var reg10082 Obj
		if reg10077 == True {
			reg10078 := PrimHead(Parse__shen_4_5singleline_6)
			reg10079 := MakeSymbol("shen.skip")
			reg10080 := Call(__e, __defun__shen_4pair, reg10078, reg10079)
			reg10082 = reg10080
		} else {
			reg10081 := Call(__e, __defun__fail)
			reg10082 = reg10081
		}
		YaccParse := reg10082
		_ = YaccParse
		reg10083 := Call(__e, __defun__fail)
		reg10084 := PrimEqual(YaccParse, reg10083)
		if reg10084 == True {
			reg10085 := Call(__e, __defun__shen_4_5multiline_6, V1616)
			Parse__shen_4_5multiline_6 := reg10085
			_ = Parse__shen_4_5multiline_6
			reg10086 := Call(__e, __defun__fail)
			reg10087 := PrimEqual(reg10086, Parse__shen_4_5multiline_6)
			reg10088 := PrimNot(reg10087)
			if reg10088 == True {
				reg10089 := PrimHead(Parse__shen_4_5multiline_6)
				reg10090 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg10089, reg10090)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<comment>", value: __defun__shen_4_5comment_6})

	__defun__shen_4_5singleline_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1618 := __args[0]
		_ = V1618
		reg10093 := Call(__e, __defun__shen_4_5backslash_6, V1618)
		Parse__shen_4_5backslash_6 := reg10093
		_ = Parse__shen_4_5backslash_6
		reg10094 := Call(__e, __defun__fail)
		reg10095 := PrimEqual(reg10094, Parse__shen_4_5backslash_6)
		reg10096 := PrimNot(reg10095)
		if reg10096 == True {
			reg10097 := Call(__e, __defun__shen_4_5backslash_6, Parse__shen_4_5backslash_6)
			Parse__shen_4_5backslash_6 := reg10097
			_ = Parse__shen_4_5backslash_6
			reg10098 := Call(__e, __defun__fail)
			reg10099 := PrimEqual(reg10098, Parse__shen_4_5backslash_6)
			reg10100 := PrimNot(reg10099)
			if reg10100 == True {
				reg10101 := Call(__e, __defun__shen_4_5anysingle_6, Parse__shen_4_5backslash_6)
				Parse__shen_4_5anysingle_6 := reg10101
				_ = Parse__shen_4_5anysingle_6
				reg10102 := Call(__e, __defun__fail)
				reg10103 := PrimEqual(reg10102, Parse__shen_4_5anysingle_6)
				reg10104 := PrimNot(reg10103)
				if reg10104 == True {
					reg10105 := Call(__e, __defun__shen_4_5return_6, Parse__shen_4_5anysingle_6)
					Parse__shen_4_5return_6 := reg10105
					_ = Parse__shen_4_5return_6
					reg10106 := Call(__e, __defun__fail)
					reg10107 := PrimEqual(reg10106, Parse__shen_4_5return_6)
					reg10108 := PrimNot(reg10107)
					if reg10108 == True {
						reg10109 := PrimHead(Parse__shen_4_5return_6)
						reg10110 := MakeSymbol("shen.skip")
						__e.TailApply(__defun__shen_4pair, reg10109, reg10110)
						return
					} else {
						__e.TailApply(__defun__fail)
						return
					}
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<singleline>", value: __defun__shen_4_5singleline_6})

	__defun__shen_4_5backslash_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1621 := __args[0]
		_ = V1621
		reg10116 := PrimHead(V1621)
		reg10117 := PrimIsPair(reg10116)
		var reg10125 Obj
		if reg10117 == True {
			reg10118 := MakeNumber(92)
			reg10119 := Call(__e, __defun__shen_4hdhd, V1621)
			reg10120 := PrimEqual(reg10118, reg10119)
			var reg10123 Obj
			if reg10120 == True {
				reg10121 := True
				reg10123 = reg10121
			} else {
				reg10122 := False
				reg10123 = reg10122
			}
			reg10125 = reg10123
		} else {
			reg10124 := False
			reg10125 = reg10124
		}
		if reg10125 == True {
			reg10126 := Call(__e, __defun__shen_4tlhd, V1621)
			reg10127 := Call(__e, __defun__shen_4hdtl, V1621)
			reg10128 := Call(__e, __defun__shen_4pair, reg10126, reg10127)
			NewStream1619 := reg10128
			_ = NewStream1619
			reg10129 := PrimHead(NewStream1619)
			reg10130 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg10129, reg10130)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<backslash>", value: __defun__shen_4_5backslash_6})

	__defun__shen_4_5anysingle_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1623 := __args[0]
		_ = V1623
		reg10133 := Call(__e, __defun__shen_4_5non_1return_6, V1623)
		Parse__shen_4_5non_1return_6 := reg10133
		_ = Parse__shen_4_5non_1return_6
		reg10134 := Call(__e, __defun__fail)
		reg10135 := PrimEqual(reg10134, Parse__shen_4_5non_1return_6)
		reg10136 := PrimNot(reg10135)
		var reg10147 Obj
		if reg10136 == True {
			reg10137 := Call(__e, __defun__shen_4_5anysingle_6, Parse__shen_4_5non_1return_6)
			Parse__shen_4_5anysingle_6 := reg10137
			_ = Parse__shen_4_5anysingle_6
			reg10138 := Call(__e, __defun__fail)
			reg10139 := PrimEqual(reg10138, Parse__shen_4_5anysingle_6)
			reg10140 := PrimNot(reg10139)
			var reg10145 Obj
			if reg10140 == True {
				reg10141 := PrimHead(Parse__shen_4_5anysingle_6)
				reg10142 := MakeSymbol("shen.skip")
				reg10143 := Call(__e, __defun__shen_4pair, reg10141, reg10142)
				reg10145 = reg10143
			} else {
				reg10144 := Call(__e, __defun__fail)
				reg10145 = reg10144
			}
			reg10147 = reg10145
		} else {
			reg10146 := Call(__e, __defun__fail)
			reg10147 = reg10146
		}
		YaccParse := reg10147
		_ = YaccParse
		reg10148 := Call(__e, __defun__fail)
		reg10149 := PrimEqual(YaccParse, reg10148)
		if reg10149 == True {
			reg10150 := Call(__e, __defun___5e_6, V1623)
			Parse___5e_6 := reg10150
			_ = Parse___5e_6
			reg10151 := Call(__e, __defun__fail)
			reg10152 := PrimEqual(reg10151, Parse___5e_6)
			reg10153 := PrimNot(reg10152)
			if reg10153 == True {
				reg10154 := PrimHead(Parse___5e_6)
				reg10155 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg10154, reg10155)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<anysingle>", value: __defun__shen_4_5anysingle_6})

	__defun__shen_4_5non_1return_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1625 := __args[0]
		_ = V1625
		reg10158 := PrimHead(V1625)
		reg10159 := PrimIsPair(reg10158)
		if reg10159 == True {
			reg10160 := Call(__e, __defun__shen_4hdhd, V1625)
			Parse__X := reg10160
			_ = Parse__X
			reg10161 := MakeNumber(10)
			reg10162 := MakeNumber(13)
			reg10163 := Nil
			reg10164 := PrimCons(reg10162, reg10163)
			reg10165 := PrimCons(reg10161, reg10164)
			reg10166 := Call(__e, __defun__element_2, Parse__X, reg10165)
			reg10167 := PrimNot(reg10166)
			if reg10167 == True {
				reg10168 := Call(__e, __defun__shen_4tlhd, V1625)
				reg10169 := Call(__e, __defun__shen_4hdtl, V1625)
				reg10170 := Call(__e, __defun__shen_4pair, reg10168, reg10169)
				reg10171 := PrimHead(reg10170)
				reg10172 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg10171, reg10172)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<non-return>", value: __defun__shen_4_5non_1return_6})

	__defun__shen_4_5return_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1627 := __args[0]
		_ = V1627
		reg10176 := PrimHead(V1627)
		reg10177 := PrimIsPair(reg10176)
		if reg10177 == True {
			reg10178 := Call(__e, __defun__shen_4hdhd, V1627)
			Parse__X := reg10178
			_ = Parse__X
			reg10179 := MakeNumber(10)
			reg10180 := MakeNumber(13)
			reg10181 := Nil
			reg10182 := PrimCons(reg10180, reg10181)
			reg10183 := PrimCons(reg10179, reg10182)
			reg10184 := Call(__e, __defun__element_2, Parse__X, reg10183)
			if reg10184 == True {
				reg10185 := Call(__e, __defun__shen_4tlhd, V1627)
				reg10186 := Call(__e, __defun__shen_4hdtl, V1627)
				reg10187 := Call(__e, __defun__shen_4pair, reg10185, reg10186)
				reg10188 := PrimHead(reg10187)
				reg10189 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg10188, reg10189)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<return>", value: __defun__shen_4_5return_6})

	__defun__shen_4_5multiline_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1629 := __args[0]
		_ = V1629
		reg10193 := Call(__e, __defun__shen_4_5backslash_6, V1629)
		Parse__shen_4_5backslash_6 := reg10193
		_ = Parse__shen_4_5backslash_6
		reg10194 := Call(__e, __defun__fail)
		reg10195 := PrimEqual(reg10194, Parse__shen_4_5backslash_6)
		reg10196 := PrimNot(reg10195)
		if reg10196 == True {
			reg10197 := Call(__e, __defun__shen_4_5times_6, Parse__shen_4_5backslash_6)
			Parse__shen_4_5times_6 := reg10197
			_ = Parse__shen_4_5times_6
			reg10198 := Call(__e, __defun__fail)
			reg10199 := PrimEqual(reg10198, Parse__shen_4_5times_6)
			reg10200 := PrimNot(reg10199)
			if reg10200 == True {
				reg10201 := Call(__e, __defun__shen_4_5anymulti_6, Parse__shen_4_5times_6)
				Parse__shen_4_5anymulti_6 := reg10201
				_ = Parse__shen_4_5anymulti_6
				reg10202 := Call(__e, __defun__fail)
				reg10203 := PrimEqual(reg10202, Parse__shen_4_5anymulti_6)
				reg10204 := PrimNot(reg10203)
				if reg10204 == True {
					reg10205 := PrimHead(Parse__shen_4_5anymulti_6)
					reg10206 := MakeSymbol("shen.skip")
					__e.TailApply(__defun__shen_4pair, reg10205, reg10206)
					return
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<multiline>", value: __defun__shen_4_5multiline_6})

	__defun__shen_4_5times_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1632 := __args[0]
		_ = V1632
		reg10211 := PrimHead(V1632)
		reg10212 := PrimIsPair(reg10211)
		var reg10220 Obj
		if reg10212 == True {
			reg10213 := MakeNumber(42)
			reg10214 := Call(__e, __defun__shen_4hdhd, V1632)
			reg10215 := PrimEqual(reg10213, reg10214)
			var reg10218 Obj
			if reg10215 == True {
				reg10216 := True
				reg10218 = reg10216
			} else {
				reg10217 := False
				reg10218 = reg10217
			}
			reg10220 = reg10218
		} else {
			reg10219 := False
			reg10220 = reg10219
		}
		if reg10220 == True {
			reg10221 := Call(__e, __defun__shen_4tlhd, V1632)
			reg10222 := Call(__e, __defun__shen_4hdtl, V1632)
			reg10223 := Call(__e, __defun__shen_4pair, reg10221, reg10222)
			NewStream1630 := reg10223
			_ = NewStream1630
			reg10224 := PrimHead(NewStream1630)
			reg10225 := MakeSymbol("shen.skip")
			__e.TailApply(__defun__shen_4pair, reg10224, reg10225)
			return
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<times>", value: __defun__shen_4_5times_6})

	__defun__shen_4_5anymulti_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1634 := __args[0]
		_ = V1634
		reg10228 := Call(__e, __defun__shen_4_5comment_6, V1634)
		Parse__shen_4_5comment_6 := reg10228
		_ = Parse__shen_4_5comment_6
		reg10229 := Call(__e, __defun__fail)
		reg10230 := PrimEqual(reg10229, Parse__shen_4_5comment_6)
		reg10231 := PrimNot(reg10230)
		var reg10242 Obj
		if reg10231 == True {
			reg10232 := Call(__e, __defun__shen_4_5anymulti_6, Parse__shen_4_5comment_6)
			Parse__shen_4_5anymulti_6 := reg10232
			_ = Parse__shen_4_5anymulti_6
			reg10233 := Call(__e, __defun__fail)
			reg10234 := PrimEqual(reg10233, Parse__shen_4_5anymulti_6)
			reg10235 := PrimNot(reg10234)
			var reg10240 Obj
			if reg10235 == True {
				reg10236 := PrimHead(Parse__shen_4_5anymulti_6)
				reg10237 := MakeSymbol("shen.skip")
				reg10238 := Call(__e, __defun__shen_4pair, reg10236, reg10237)
				reg10240 = reg10238
			} else {
				reg10239 := Call(__e, __defun__fail)
				reg10240 = reg10239
			}
			reg10242 = reg10240
		} else {
			reg10241 := Call(__e, __defun__fail)
			reg10242 = reg10241
		}
		YaccParse := reg10242
		_ = YaccParse
		reg10243 := Call(__e, __defun__fail)
		reg10244 := PrimEqual(YaccParse, reg10243)
		if reg10244 == True {
			reg10245 := Call(__e, __defun__shen_4_5times_6, V1634)
			Parse__shen_4_5times_6 := reg10245
			_ = Parse__shen_4_5times_6
			reg10246 := Call(__e, __defun__fail)
			reg10247 := PrimEqual(reg10246, Parse__shen_4_5times_6)
			reg10248 := PrimNot(reg10247)
			var reg10259 Obj
			if reg10248 == True {
				reg10249 := Call(__e, __defun__shen_4_5backslash_6, Parse__shen_4_5times_6)
				Parse__shen_4_5backslash_6 := reg10249
				_ = Parse__shen_4_5backslash_6
				reg10250 := Call(__e, __defun__fail)
				reg10251 := PrimEqual(reg10250, Parse__shen_4_5backslash_6)
				reg10252 := PrimNot(reg10251)
				var reg10257 Obj
				if reg10252 == True {
					reg10253 := PrimHead(Parse__shen_4_5backslash_6)
					reg10254 := MakeSymbol("shen.skip")
					reg10255 := Call(__e, __defun__shen_4pair, reg10253, reg10254)
					reg10257 = reg10255
				} else {
					reg10256 := Call(__e, __defun__fail)
					reg10257 = reg10256
				}
				reg10259 = reg10257
			} else {
				reg10258 := Call(__e, __defun__fail)
				reg10259 = reg10258
			}
			YaccParse := reg10259
			_ = YaccParse
			reg10260 := Call(__e, __defun__fail)
			reg10261 := PrimEqual(YaccParse, reg10260)
			if reg10261 == True {
				reg10262 := PrimHead(V1634)
				reg10263 := PrimIsPair(reg10262)
				if reg10263 == True {
					reg10264 := Call(__e, __defun__shen_4hdhd, V1634)
					Parse__X := reg10264
					_ = Parse__X
					reg10265 := Call(__e, __defun__shen_4tlhd, V1634)
					reg10266 := Call(__e, __defun__shen_4hdtl, V1634)
					reg10267 := Call(__e, __defun__shen_4pair, reg10265, reg10266)
					reg10268 := Call(__e, __defun__shen_4_5anymulti_6, reg10267)
					Parse__shen_4_5anymulti_6 := reg10268
					_ = Parse__shen_4_5anymulti_6
					reg10269 := Call(__e, __defun__fail)
					reg10270 := PrimEqual(reg10269, Parse__shen_4_5anymulti_6)
					reg10271 := PrimNot(reg10270)
					if reg10271 == True {
						reg10272 := PrimHead(Parse__shen_4_5anymulti_6)
						reg10273 := MakeSymbol("shen.skip")
						__e.TailApply(__defun__shen_4pair, reg10272, reg10273)
						return
					} else {
						__e.TailApply(__defun__fail)
						return
					}
				} else {
					__e.TailApply(__defun__fail)
					return
				}
			} else {
				__e.Return(YaccParse)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<anymulti>", value: __defun__shen_4_5anymulti_6})

	__defun__shen_4_5whitespaces_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1636 := __args[0]
		_ = V1636
		reg10277 := Call(__e, __defun__shen_4_5whitespace_6, V1636)
		Parse__shen_4_5whitespace_6 := reg10277
		_ = Parse__shen_4_5whitespace_6
		reg10278 := Call(__e, __defun__fail)
		reg10279 := PrimEqual(reg10278, Parse__shen_4_5whitespace_6)
		reg10280 := PrimNot(reg10279)
		var reg10291 Obj
		if reg10280 == True {
			reg10281 := Call(__e, __defun__shen_4_5whitespaces_6, Parse__shen_4_5whitespace_6)
			Parse__shen_4_5whitespaces_6 := reg10281
			_ = Parse__shen_4_5whitespaces_6
			reg10282 := Call(__e, __defun__fail)
			reg10283 := PrimEqual(reg10282, Parse__shen_4_5whitespaces_6)
			reg10284 := PrimNot(reg10283)
			var reg10289 Obj
			if reg10284 == True {
				reg10285 := PrimHead(Parse__shen_4_5whitespaces_6)
				reg10286 := MakeSymbol("shen.skip")
				reg10287 := Call(__e, __defun__shen_4pair, reg10285, reg10286)
				reg10289 = reg10287
			} else {
				reg10288 := Call(__e, __defun__fail)
				reg10289 = reg10288
			}
			reg10291 = reg10289
		} else {
			reg10290 := Call(__e, __defun__fail)
			reg10291 = reg10290
		}
		YaccParse := reg10291
		_ = YaccParse
		reg10292 := Call(__e, __defun__fail)
		reg10293 := PrimEqual(YaccParse, reg10292)
		if reg10293 == True {
			reg10294 := Call(__e, __defun__shen_4_5whitespace_6, V1636)
			Parse__shen_4_5whitespace_6 := reg10294
			_ = Parse__shen_4_5whitespace_6
			reg10295 := Call(__e, __defun__fail)
			reg10296 := PrimEqual(reg10295, Parse__shen_4_5whitespace_6)
			reg10297 := PrimNot(reg10296)
			if reg10297 == True {
				reg10298 := PrimHead(Parse__shen_4_5whitespace_6)
				reg10299 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg10298, reg10299)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.Return(YaccParse)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<whitespaces>", value: __defun__shen_4_5whitespaces_6})

	__defun__shen_4_5whitespace_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1638 := __args[0]
		_ = V1638
		reg10302 := PrimHead(V1638)
		reg10303 := PrimIsPair(reg10302)
		if reg10303 == True {
			reg10304 := Call(__e, __defun__shen_4hdhd, V1638)
			Parse__X := reg10304
			_ = Parse__X
			Parse__Case := Parse__X
			_ = Parse__Case
			reg10305 := MakeNumber(32)
			reg10306 := PrimEqual(Parse__Case, reg10305)
			var reg10327 Obj
			if reg10306 == True {
				reg10307 := True
				reg10327 = reg10307
			} else {
				reg10308 := MakeNumber(13)
				reg10309 := PrimEqual(Parse__Case, reg10308)
				var reg10323 Obj
				if reg10309 == True {
					reg10310 := True
					reg10323 = reg10310
				} else {
					reg10311 := MakeNumber(10)
					reg10312 := PrimEqual(Parse__Case, reg10311)
					var reg10319 Obj
					if reg10312 == True {
						reg10313 := True
						reg10319 = reg10313
					} else {
						reg10314 := MakeNumber(9)
						reg10315 := PrimEqual(Parse__Case, reg10314)
						var reg10318 Obj
						if reg10315 == True {
							reg10316 := True
							reg10318 = reg10316
						} else {
							reg10317 := False
							reg10318 = reg10317
						}
						reg10319 = reg10318
					}
					var reg10322 Obj
					if reg10319 == True {
						reg10320 := True
						reg10322 = reg10320
					} else {
						reg10321 := False
						reg10322 = reg10321
					}
					reg10323 = reg10322
				}
				var reg10326 Obj
				if reg10323 == True {
					reg10324 := True
					reg10326 = reg10324
				} else {
					reg10325 := False
					reg10326 = reg10325
				}
				reg10327 = reg10326
			}
			if reg10327 == True {
				reg10328 := Call(__e, __defun__shen_4tlhd, V1638)
				reg10329 := Call(__e, __defun__shen_4hdtl, V1638)
				reg10330 := Call(__e, __defun__shen_4pair, reg10328, reg10329)
				reg10331 := PrimHead(reg10330)
				reg10332 := MakeSymbol("shen.skip")
				__e.TailApply(__defun__shen_4pair, reg10331, reg10332)
				return
			} else {
				__e.TailApply(__defun__fail)
				return
			}
		} else {
			__e.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.<whitespace>", value: __defun__shen_4_5whitespace_6})

	__defun__shen_4cons__form = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1640 := __args[0]
		_ = V1640
		reg10336 := Nil
		reg10337 := PrimEqual(reg10336, V1640)
		if reg10337 == True {
			reg10338 := Nil
			__e.Return(reg10338)
			return
		} else {
			reg10339 := PrimIsPair(V1640)
			var reg10373 Obj
			if reg10339 == True {
				reg10340 := PrimTail(V1640)
				reg10341 := PrimIsPair(reg10340)
				var reg10368 Obj
				if reg10341 == True {
					reg10342 := PrimTail(V1640)
					reg10343 := PrimTail(reg10342)
					reg10344 := PrimIsPair(reg10343)
					var reg10363 Obj
					if reg10344 == True {
						reg10345 := Nil
						reg10346 := PrimTail(V1640)
						reg10347 := PrimTail(reg10346)
						reg10348 := PrimTail(reg10347)
						reg10349 := PrimEqual(reg10345, reg10348)
						var reg10358 Obj
						if reg10349 == True {
							reg10350 := PrimTail(V1640)
							reg10351 := PrimHead(reg10350)
							reg10352 := MakeSymbol("bar!")
							reg10353 := PrimEqual(reg10351, reg10352)
							var reg10356 Obj
							if reg10353 == True {
								reg10354 := True
								reg10356 = reg10354
							} else {
								reg10355 := False
								reg10356 = reg10355
							}
							reg10358 = reg10356
						} else {
							reg10357 := False
							reg10358 = reg10357
						}
						var reg10361 Obj
						if reg10358 == True {
							reg10359 := True
							reg10361 = reg10359
						} else {
							reg10360 := False
							reg10361 = reg10360
						}
						reg10363 = reg10361
					} else {
						reg10362 := False
						reg10363 = reg10362
					}
					var reg10366 Obj
					if reg10363 == True {
						reg10364 := True
						reg10366 = reg10364
					} else {
						reg10365 := False
						reg10366 = reg10365
					}
					reg10368 = reg10366
				} else {
					reg10367 := False
					reg10368 = reg10367
				}
				var reg10371 Obj
				if reg10368 == True {
					reg10369 := True
					reg10371 = reg10369
				} else {
					reg10370 := False
					reg10371 = reg10370
				}
				reg10373 = reg10371
			} else {
				reg10372 := False
				reg10373 = reg10372
			}
			if reg10373 == True {
				reg10374 := MakeSymbol("cons")
				reg10375 := PrimHead(V1640)
				reg10376 := PrimTail(V1640)
				reg10377 := PrimTail(reg10376)
				reg10378 := PrimCons(reg10375, reg10377)
				reg10379 := PrimCons(reg10374, reg10378)
				__e.Return(reg10379)
				return
			} else {
				reg10380 := PrimIsPair(V1640)
				if reg10380 == True {
					reg10381 := MakeSymbol("cons")
					reg10382 := PrimHead(V1640)
					reg10383 := PrimTail(V1640)
					reg10384 := Call(__e, __defun__shen_4cons__form, reg10383)
					reg10385 := Nil
					reg10386 := PrimCons(reg10384, reg10385)
					reg10387 := PrimCons(reg10382, reg10386)
					reg10388 := PrimCons(reg10381, reg10387)
					__e.Return(reg10388)
					return
				} else {
					reg10389 := MakeSymbol("shen.cons_form")
					__e.TailApply(__defun__shen_4f__error, reg10389)
					return
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.cons_form", value: __defun__shen_4cons__form})

	__defun__shen_4package_1macro = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1645 := __args[0]
		_ = V1645
		V1646 := __args[1]
		_ = V1646
		reg10391 := PrimIsPair(V1645)
		var reg10415 Obj
		if reg10391 == True {
			reg10392 := MakeSymbol("$")
			reg10393 := PrimHead(V1645)
			reg10394 := PrimEqual(reg10392, reg10393)
			var reg10410 Obj
			if reg10394 == True {
				reg10395 := PrimTail(V1645)
				reg10396 := PrimIsPair(reg10395)
				var reg10405 Obj
				if reg10396 == True {
					reg10397 := Nil
					reg10398 := PrimTail(V1645)
					reg10399 := PrimTail(reg10398)
					reg10400 := PrimEqual(reg10397, reg10399)
					var reg10403 Obj
					if reg10400 == True {
						reg10401 := True
						reg10403 = reg10401
					} else {
						reg10402 := False
						reg10403 = reg10402
					}
					reg10405 = reg10403
				} else {
					reg10404 := False
					reg10405 = reg10404
				}
				var reg10408 Obj
				if reg10405 == True {
					reg10406 := True
					reg10408 = reg10406
				} else {
					reg10407 := False
					reg10408 = reg10407
				}
				reg10410 = reg10408
			} else {
				reg10409 := False
				reg10410 = reg10409
			}
			var reg10413 Obj
			if reg10410 == True {
				reg10411 := True
				reg10413 = reg10411
			} else {
				reg10412 := False
				reg10413 = reg10412
			}
			reg10415 = reg10413
		} else {
			reg10414 := False
			reg10415 = reg10414
		}
		if reg10415 == True {
			reg10416 := PrimTail(V1645)
			reg10417 := PrimHead(reg10416)
			reg10418 := Call(__e, __defun__explode, reg10417)
			__e.TailApply(__defun__append, reg10418, V1646)
			return
		} else {
			reg10420 := PrimIsPair(V1645)
			var reg10452 Obj
			if reg10420 == True {
				reg10421 := MakeSymbol("package")
				reg10422 := PrimHead(V1645)
				reg10423 := PrimEqual(reg10421, reg10422)
				var reg10447 Obj
				if reg10423 == True {
					reg10424 := PrimTail(V1645)
					reg10425 := PrimIsPair(reg10424)
					var reg10442 Obj
					if reg10425 == True {
						reg10426 := MakeSymbol("null")
						reg10427 := PrimTail(V1645)
						reg10428 := PrimHead(reg10427)
						reg10429 := PrimEqual(reg10426, reg10428)
						var reg10437 Obj
						if reg10429 == True {
							reg10430 := PrimTail(V1645)
							reg10431 := PrimTail(reg10430)
							reg10432 := PrimIsPair(reg10431)
							var reg10435 Obj
							if reg10432 == True {
								reg10433 := True
								reg10435 = reg10433
							} else {
								reg10434 := False
								reg10435 = reg10434
							}
							reg10437 = reg10435
						} else {
							reg10436 := False
							reg10437 = reg10436
						}
						var reg10440 Obj
						if reg10437 == True {
							reg10438 := True
							reg10440 = reg10438
						} else {
							reg10439 := False
							reg10440 = reg10439
						}
						reg10442 = reg10440
					} else {
						reg10441 := False
						reg10442 = reg10441
					}
					var reg10445 Obj
					if reg10442 == True {
						reg10443 := True
						reg10445 = reg10443
					} else {
						reg10444 := False
						reg10445 = reg10444
					}
					reg10447 = reg10445
				} else {
					reg10446 := False
					reg10447 = reg10446
				}
				var reg10450 Obj
				if reg10447 == True {
					reg10448 := True
					reg10450 = reg10448
				} else {
					reg10449 := False
					reg10450 = reg10449
				}
				reg10452 = reg10450
			} else {
				reg10451 := False
				reg10452 = reg10451
			}
			if reg10452 == True {
				reg10453 := PrimTail(V1645)
				reg10454 := PrimTail(reg10453)
				reg10455 := PrimTail(reg10454)
				__e.TailApply(__defun__append, reg10455, V1646)
				return
			} else {
				reg10457 := PrimIsPair(V1645)
				var reg10480 Obj
				if reg10457 == True {
					reg10458 := MakeSymbol("package")
					reg10459 := PrimHead(V1645)
					reg10460 := PrimEqual(reg10458, reg10459)
					var reg10475 Obj
					if reg10460 == True {
						reg10461 := PrimTail(V1645)
						reg10462 := PrimIsPair(reg10461)
						var reg10470 Obj
						if reg10462 == True {
							reg10463 := PrimTail(V1645)
							reg10464 := PrimTail(reg10463)
							reg10465 := PrimIsPair(reg10464)
							var reg10468 Obj
							if reg10465 == True {
								reg10466 := True
								reg10468 = reg10466
							} else {
								reg10467 := False
								reg10468 = reg10467
							}
							reg10470 = reg10468
						} else {
							reg10469 := False
							reg10470 = reg10469
						}
						var reg10473 Obj
						if reg10470 == True {
							reg10471 := True
							reg10473 = reg10471
						} else {
							reg10472 := False
							reg10473 = reg10472
						}
						reg10475 = reg10473
					} else {
						reg10474 := False
						reg10475 = reg10474
					}
					var reg10478 Obj
					if reg10475 == True {
						reg10476 := True
						reg10478 = reg10476
					} else {
						reg10477 := False
						reg10478 = reg10477
					}
					reg10480 = reg10478
				} else {
					reg10479 := False
					reg10480 = reg10479
				}
				if reg10480 == True {
					reg10481 := PrimTail(V1645)
					reg10482 := PrimTail(reg10481)
					reg10483 := PrimHead(reg10482)
					reg10484 := Call(__e, __defun__shen_4eval_1without_1macros, reg10483)
					ListofExceptions := reg10484
					_ = ListofExceptions
					reg10485 := PrimTail(V1645)
					reg10486 := PrimHead(reg10485)
					reg10487 := Call(__e, __defun__shen_4record_1exceptions, ListofExceptions, reg10486)
					External := reg10487
					_ = External
					reg10488 := PrimTail(V1645)
					reg10489 := PrimHead(reg10488)
					reg10490 := PrimStr(reg10489)
					reg10491 := MakeString(".")
					reg10492 := PrimStringConcat(reg10490, reg10491)
					reg10493 := PrimIntern(reg10492)
					PackageNameDot := reg10493
					_ = PackageNameDot
					reg10494 := Call(__e, __defun__explode, PackageNameDot)
					ExpPackageNameDot := reg10494
					_ = ExpPackageNameDot
					reg10495 := PrimTail(V1645)
					reg10496 := PrimTail(reg10495)
					reg10497 := PrimTail(reg10496)
					reg10498 := Call(__e, __defun__shen_4packageh, PackageNameDot, ListofExceptions, reg10497, ExpPackageNameDot)
					Packaged := reg10498
					_ = Packaged
					reg10499 := PrimTail(V1645)
					reg10500 := PrimHead(reg10499)
					reg10501 := Call(__e, __defun__shen_4internal_1symbols, ExpPackageNameDot, Packaged)
					reg10502 := Call(__e, __defun__shen_4record_1internal, reg10500, reg10501)
					Internal := reg10502
					_ = Internal
					__e.TailApply(__defun__append, Packaged, V1646)
					return
				} else {
					reg10504 := PrimCons(V1645, V1646)
					__e.Return(reg10504)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.package-macro", value: __defun__shen_4package_1macro})

	__defun__shen_4record_1exceptions = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1649 := __args[0]
		_ = V1649
		V1650 := __args[1]
		_ = V1650
		reg10505 := MakeNative(func(__e Evaluator, __args ...Obj) {
			reg10506 := MakeSymbol("shen.external-symbols")
			reg10507 := MakeSymbol("*property-vector*")
			reg10508 := PrimValue(reg10507)
			__e.TailApply(__defun__get, V1650, reg10506, reg10508)
			return
		}, 0)
		reg10510 := MakeNative(func(__e Evaluator, __args ...Obj) {
			E := __args[0]
			_ = E
			reg10511 := Nil
			__e.Return(reg10511)
			return
		}, 1)
		reg10512 := Try(__e, reg10505).Catch(reg10510)
		CurrExceptions := reg10512
		_ = CurrExceptions
		reg10513 := Call(__e, __defun__union, V1649, CurrExceptions)
		AllExceptions := reg10513
		_ = AllExceptions
		reg10514 := MakeSymbol("shen.external-symbols")
		reg10515 := MakeSymbol("*property-vector*")
		reg10516 := PrimValue(reg10515)
		__e.TailApply(__defun__put, V1650, reg10514, AllExceptions, reg10516)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.record-exceptions", value: __defun__shen_4record_1exceptions})

	__defun__shen_4record_1internal = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1653 := __args[0]
		_ = V1653
		V1654 := __args[1]
		_ = V1654
		reg10518 := MakeSymbol("shen.internal-symbols")
		reg10519 := MakeNative(func(__e Evaluator, __args ...Obj) {
			reg10520 := MakeSymbol("shen.internal-symbols")
			reg10521 := MakeSymbol("*property-vector*")
			reg10522 := PrimValue(reg10521)
			__e.TailApply(__defun__get, V1653, reg10520, reg10522)
			return
		}, 0)
		reg10524 := MakeNative(func(__e Evaluator, __args ...Obj) {
			E := __args[0]
			_ = E
			reg10525 := Nil
			__e.Return(reg10525)
			return
		}, 1)
		reg10526 := Try(__e, reg10519).Catch(reg10524)
		reg10527 := Call(__e, __defun__union, V1654, reg10526)
		reg10528 := MakeSymbol("*property-vector*")
		reg10529 := PrimValue(reg10528)
		__e.TailApply(__defun__put, V1653, reg10518, reg10527, reg10529)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.record-internal", value: __defun__shen_4record_1internal})

	__defun__shen_4internal_1symbols = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1665 := __args[0]
		_ = V1665
		V1666 := __args[1]
		_ = V1666
		reg10531 := PrimIsSymbol(V1666)
		var reg10538 Obj
		if reg10531 == True {
			reg10532 := Call(__e, __defun__explode, V1666)
			reg10533 := Call(__e, __defun__shen_4prefix_2, V1665, reg10532)
			var reg10536 Obj
			if reg10533 == True {
				reg10534 := True
				reg10536 = reg10534
			} else {
				reg10535 := False
				reg10536 = reg10535
			}
			reg10538 = reg10536
		} else {
			reg10537 := False
			reg10538 = reg10537
		}
		if reg10538 == True {
			reg10539 := Nil
			reg10540 := PrimCons(V1666, reg10539)
			__e.Return(reg10540)
			return
		} else {
			reg10541 := PrimIsPair(V1666)
			if reg10541 == True {
				reg10542 := PrimHead(V1666)
				reg10543 := Call(__e, __defun__shen_4internal_1symbols, V1665, reg10542)
				reg10544 := PrimTail(V1666)
				reg10545 := Call(__e, __defun__shen_4internal_1symbols, V1665, reg10544)
				__e.TailApply(__defun__union, reg10543, reg10545)
				return
			} else {
				reg10547 := Nil
				__e.Return(reg10547)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.internal-symbols", value: __defun__shen_4internal_1symbols})

	__defun__shen_4packageh = MakeNative(func(__e Evaluator, __args ...Obj) {
		V1683 := __args[0]
		_ = V1683
		V1684 := __args[1]
		_ = V1684
		V1685 := __args[2]
		_ = V1685
		V1686 := __args[3]
		_ = V1686
		reg10548 := PrimIsPair(V1685)
		if reg10548 == True {
			reg10549 := PrimHead(V1685)
			reg10550 := Call(__e, __defun__shen_4packageh, V1683, V1684, reg10549, V1686)
			reg10551 := PrimTail(V1685)
			reg10552 := Call(__e, __defun__shen_4packageh, V1683, V1684, reg10551, V1686)
			reg10553 := PrimCons(reg10550, reg10552)
			__e.Return(reg10553)
			return
		} else {
			reg10554 := Call(__e, __defun__shen_4sysfunc_2, V1685)
			var reg10578 Obj
			if reg10554 == True {
				reg10555 := True
				reg10578 = reg10555
			} else {
				reg10556 := PrimIsVariable(V1685)
				var reg10574 Obj
				if reg10556 == True {
					reg10557 := True
					reg10574 = reg10557
				} else {
					reg10558 := Call(__e, __defun__element_2, V1685, V1684)
					var reg10570 Obj
					if reg10558 == True {
						reg10559 := True
						reg10570 = reg10559
					} else {
						reg10560 := Call(__e, __defun__shen_4doubleunderline_2, V1685)
						var reg10566 Obj
						if reg10560 == True {
							reg10561 := True
							reg10566 = reg10561
						} else {
							reg10562 := Call(__e, __defun__shen_4singleunderline_2, V1685)
							var reg10565 Obj
							if reg10562 == True {
								reg10563 := True
								reg10565 = reg10563
							} else {
								reg10564 := False
								reg10565 = reg10564
							}
							reg10566 = reg10565
						}
						var reg10569 Obj
						if reg10566 == True {
							reg10567 := True
							reg10569 = reg10567
						} else {
							reg10568 := False
							reg10569 = reg10568
						}
						reg10570 = reg10569
					}
					var reg10573 Obj
					if reg10570 == True {
						reg10571 := True
						reg10573 = reg10571
					} else {
						reg10572 := False
						reg10573 = reg10572
					}
					reg10574 = reg10573
				}
				var reg10577 Obj
				if reg10574 == True {
					reg10575 := True
					reg10577 = reg10575
				} else {
					reg10576 := False
					reg10577 = reg10576
				}
				reg10578 = reg10577
			}
			if reg10578 == True {
				__e.Return(V1685)
				return
			} else {
				reg10579 := PrimIsSymbol(V1685)
				var reg10605 Obj
				if reg10579 == True {
					reg10580 := Call(__e, __defun__explode, V1685)
					ExplodeX := reg10580
					_ = ExplodeX
					reg10581 := MakeString("s")
					reg10582 := MakeString("h")
					reg10583 := MakeString("e")
					reg10584 := MakeString("n")
					reg10585 := MakeString(".")
					reg10586 := Nil
					reg10587 := PrimCons(reg10585, reg10586)
					reg10588 := PrimCons(reg10584, reg10587)
					reg10589 := PrimCons(reg10583, reg10588)
					reg10590 := PrimCons(reg10582, reg10589)
					reg10591 := PrimCons(reg10581, reg10590)
					reg10592 := Call(__e, __defun__shen_4prefix_2, reg10591, ExplodeX)
					reg10593 := PrimNot(reg10592)
					var reg10600 Obj
					if reg10593 == True {
						reg10594 := Call(__e, __defun__shen_4prefix_2, V1686, ExplodeX)
						reg10595 := PrimNot(reg10594)
						var reg10598 Obj
						if reg10595 == True {
							reg10596 := True
							reg10598 = reg10596
						} else {
							reg10597 := False
							reg10598 = reg10597
						}
						reg10600 = reg10598
					} else {
						reg10599 := False
						reg10600 = reg10599
					}
					var reg10603 Obj
					if reg10600 == True {
						reg10601 := True
						reg10603 = reg10601
					} else {
						reg10602 := False
						reg10603 = reg10602
					}
					reg10605 = reg10603
				} else {
					reg10604 := False
					reg10605 = reg10604
				}
				if reg10605 == True {
					__e.TailApply(__defun__concat, V1683, V1685)
					return
				} else {
					__e.Return(V1685)
					return
				}
			}
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.packageh", value: __defun__shen_4packageh})

}
