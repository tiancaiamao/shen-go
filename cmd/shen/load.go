package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__load Obj                       // load
var __defun__shen_4load_1help Obj           // shen.load-help
var __defun__shen_4remove_1synonyms Obj     // shen.remove-synonyms
var __defun__shen_4typecheck_1and_1load Obj // shen.typecheck-and-load
var __defun__shen_4typetable Obj            // shen.typetable
var __defun__shen_4assumetype Obj           // shen.assumetype
var __defun__shen_4unwind_1types Obj        // shen.unwind-types
var __defun__shen_4remtype Obj              // shen.remtype
var __defun__shen_4removetype Obj           // shen.removetype
var __defun__shen_4_5sig_7rest_6 Obj        // shen.<sig+rest>
var __defun__write_1to_1file Obj            // write-to-file

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		reg16463 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__e.Return(reg16463)
		return
	}, 0))
	__defun__load = MakeNative(func(__e Evaluator, __args ...Obj) {
		V671 := __args[0]
		_ = V671
		reg16464 := MakeSymbol("run")
		reg16465 := PrimGetTime(reg16464)
		Start := reg16465
		_ = Start
		reg16466 := MakeSymbol("shen.*tc*")
		reg16467 := PrimValue(reg16466)
		reg16468 := Call(__e, __defun__read_1file, V671)
		reg16469 := Call(__e, __defun__shen_4load_1help, reg16467, reg16468)
		Result := reg16469
		_ = Result
		reg16470 := MakeSymbol("run")
		reg16471 := PrimGetTime(reg16470)
		Finish := reg16471
		_ = Finish
		reg16472 := PrimNumberSubtract(Finish, Start)
		Time := reg16472
		_ = Time
		reg16473 := MakeString("\nrun time: ")
		reg16474 := PrimStr(Time)
		reg16475 := MakeString(" secs\n")
		reg16476 := PrimStringConcat(reg16474, reg16475)
		reg16477 := PrimStringConcat(reg16473, reg16476)
		reg16478 := Call(__e, __defun__stoutput)
		reg16479 := Call(__e, __defun__shen_4prhush, reg16477, reg16478)
		Message := reg16479
		_ = Message
		Load := Result
		_ = Load
		reg16480 := MakeSymbol("shen.*tc*")
		reg16481 := PrimValue(reg16480)
		var reg16491 Obj
		if reg16481 == True {
			reg16482 := MakeString("\ntypechecked in ")
			reg16483 := Call(__e, __defun__inferences)
			reg16484 := MakeString(" inferences\n")
			reg16485 := MakeSymbol("shen.a")
			reg16486 := Call(__e, __defun__shen_4app, reg16483, reg16484, reg16485)
			reg16487 := PrimStringConcat(reg16482, reg16486)
			reg16488 := Call(__e, __defun__stoutput)
			reg16489 := Call(__e, __defun__shen_4prhush, reg16487, reg16488)
			reg16491 = reg16489
		} else {
			reg16490 := MakeSymbol("shen.skip")
			reg16491 = reg16490
		}
		Infs := reg16491
		_ = Infs
		reg16492 := MakeSymbol("loaded")
		__e.Return(reg16492)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "load", value: __defun__load})

	__defun__shen_4load_1help = MakeNative(func(__e Evaluator, __args ...Obj) {
		V678 := __args[0]
		_ = V678
		V679 := __args[1]
		_ = V679
		reg16493 := False
		reg16494 := PrimEqual(reg16493, V678)
		if reg16494 == True {
			reg16495 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				reg16496 := Call(__e, __defun__shen_4eval_1without_1macros, X)
				reg16497 := MakeString("\n")
				reg16498 := MakeSymbol("shen.s")
				reg16499 := Call(__e, __defun__shen_4app, reg16496, reg16497, reg16498)
				reg16500 := Call(__e, __defun__stoutput)
				__e.TailApply(__defun__shen_4prhush, reg16499, reg16500)
				return
			}, 1)
			__e.TailApply(__defun__shen_4for_1each, reg16495, V679)
			return
		} else {
			reg16503 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(__defun__shen_4remove_1synonyms, X)
				return
			}, 1)
			reg16505 := Call(__e, __defun__mapcan, reg16503, V679)
			RemoveSynonyms := reg16505
			_ = RemoveSynonyms
			reg16506 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(__defun__shen_4typetable, X)
				return
			}, 1)
			reg16508 := Call(__e, __defun__mapcan, reg16506, RemoveSynonyms)
			Table := reg16508
			_ = Table
			reg16509 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(__defun__shen_4assumetype, X)
				return
			}, 1)
			reg16511 := Call(__e, __defun__shen_4for_1each, reg16509, Table)
			Assume := reg16511
			_ = Assume
			reg16512 := MakeNative(func(__e Evaluator, __args ...Obj) {
				reg16513 := MakeNative(func(__e Evaluator, __args ...Obj) {
					X := __args[0]
					_ = X
					__e.TailApply(__defun__shen_4typecheck_1and_1load, X)
					return
				}, 1)
				__e.TailApply(__defun__shen_4for_1each, reg16513, RemoveSynonyms)
				return
			}, 0)
			reg16516 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				__e.TailApply(__defun__shen_4unwind_1types, E, Table)
				return
			}, 1)
			reg16518 := Try(__e, reg16512).Catch(reg16516)
			__e.Return(reg16518)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.load-help", value: __defun__shen_4load_1help})

	__defun__shen_4remove_1synonyms = MakeNative(func(__e Evaluator, __args ...Obj) {
		V681 := __args[0]
		_ = V681
		reg16519 := PrimIsPair(V681)
		var reg16527 Obj
		if reg16519 == True {
			reg16520 := MakeSymbol("shen.synonyms-help")
			reg16521 := PrimHead(V681)
			reg16522 := PrimEqual(reg16520, reg16521)
			var reg16525 Obj
			if reg16522 == True {
				reg16523 := True
				reg16525 = reg16523
			} else {
				reg16524 := False
				reg16525 = reg16524
			}
			reg16527 = reg16525
		} else {
			reg16526 := False
			reg16527 = reg16526
		}
		if reg16527 == True {
			reg16528 := Call(__e, __defun__eval, V681)
			_ = reg16528
			reg16529 := Nil
			__e.Return(reg16529)
			return
		} else {
			reg16530 := Nil
			reg16531 := PrimCons(V681, reg16530)
			__e.Return(reg16531)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.remove-synonyms", value: __defun__shen_4remove_1synonyms})

	__defun__shen_4typecheck_1and_1load = MakeNative(func(__e Evaluator, __args ...Obj) {
		V683 := __args[0]
		_ = V683
		reg16532 := MakeNumber(1)
		reg16533 := Call(__e, __defun__nl, reg16532)
		_ = reg16533
		reg16534 := MakeSymbol("A")
		reg16535 := Call(__e, __defun__gensym, reg16534)
		__e.TailApply(__defun__shen_4typecheck_1and_1evaluate, V683, reg16535)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.typecheck-and-load", value: __defun__shen_4typecheck_1and_1load})

	__defun__shen_4typetable = MakeNative(func(__e Evaluator, __args ...Obj) {
		V689 := __args[0]
		_ = V689
		reg16537 := PrimIsPair(V689)
		var reg16552 Obj
		if reg16537 == True {
			reg16538 := MakeSymbol("define")
			reg16539 := PrimHead(V689)
			reg16540 := PrimEqual(reg16538, reg16539)
			var reg16547 Obj
			if reg16540 == True {
				reg16541 := PrimTail(V689)
				reg16542 := PrimIsPair(reg16541)
				var reg16545 Obj
				if reg16542 == True {
					reg16543 := True
					reg16545 = reg16543
				} else {
					reg16544 := False
					reg16545 = reg16544
				}
				reg16547 = reg16545
			} else {
				reg16546 := False
				reg16547 = reg16546
			}
			var reg16550 Obj
			if reg16547 == True {
				reg16548 := True
				reg16550 = reg16548
			} else {
				reg16549 := False
				reg16550 = reg16549
			}
			reg16552 = reg16550
		} else {
			reg16551 := False
			reg16552 = reg16551
		}
		if reg16552 == True {
			reg16553 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Y := __args[0]
				_ = Y
				__e.TailApply(__defun__shen_4_5sig_7rest_6, Y)
				return
			}, 1)
			reg16555 := PrimTail(V689)
			reg16556 := PrimTail(reg16555)
			reg16557 := MakeNative(func(__e Evaluator, __args ...Obj) {
				E := __args[0]
				_ = E
				reg16558 := PrimTail(V689)
				reg16559 := PrimHead(reg16558)
				reg16560 := MakeString(" lacks a proper signature.\n")
				reg16561 := MakeSymbol("shen.a")
				reg16562 := Call(__e, __defun__shen_4app, reg16559, reg16560, reg16561)
				reg16563 := PrimSimpleError(reg16562)
				__e.Return(reg16563)
				return
			}, 1)
			reg16564 := Call(__e, __defun__compile, reg16553, reg16556, reg16557)
			Sig := reg16564
			_ = Sig
			reg16565 := PrimTail(V689)
			reg16566 := PrimHead(reg16565)
			reg16567 := PrimCons(reg16566, Sig)
			reg16568 := Nil
			reg16569 := PrimCons(reg16567, reg16568)
			__e.Return(reg16569)
			return
		} else {
			reg16570 := Nil
			__e.Return(reg16570)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.typetable", value: __defun__shen_4typetable})

	__defun__shen_4assumetype = MakeNative(func(__e Evaluator, __args ...Obj) {
		V691 := __args[0]
		_ = V691
		reg16571 := PrimIsPair(V691)
		if reg16571 == True {
			reg16572 := PrimHead(V691)
			reg16573 := PrimTail(V691)
			__e.TailApply(__defun__declare, reg16572, reg16573)
			return
		} else {
			reg16575 := MakeSymbol("shen.assumetype")
			__e.TailApply(__defun__shen_4f__error, reg16575)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.assumetype", value: __defun__shen_4assumetype})

	__defun__shen_4unwind_1types = MakeNative(func(__e Evaluator, __args ...Obj) {
		V698 := __args[0]
		_ = V698
		V699 := __args[1]
		_ = V699
		reg16577 := Nil
		reg16578 := PrimEqual(reg16577, V699)
		if reg16578 == True {
			reg16579 := PrimErrorToString(V698)
			reg16580 := PrimSimpleError(reg16579)
			__e.Return(reg16580)
			return
		} else {
			reg16581 := PrimIsPair(V699)
			var reg16588 Obj
			if reg16581 == True {
				reg16582 := PrimHead(V699)
				reg16583 := PrimIsPair(reg16582)
				var reg16586 Obj
				if reg16583 == True {
					reg16584 := True
					reg16586 = reg16584
				} else {
					reg16585 := False
					reg16586 = reg16585
				}
				reg16588 = reg16586
			} else {
				reg16587 := False
				reg16588 = reg16587
			}
			if reg16588 == True {
				reg16589 := PrimHead(V699)
				reg16590 := PrimHead(reg16589)
				reg16591 := Call(__e, __defun__shen_4remtype, reg16590)
				_ = reg16591
				reg16592 := PrimTail(V699)
				__e.TailApply(__defun__shen_4unwind_1types, V698, reg16592)
				return
			} else {
				reg16594 := MakeSymbol("shen.unwind-types")
				__e.TailApply(__defun__shen_4f__error, reg16594)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.unwind-types", value: __defun__shen_4unwind_1types})

	__defun__shen_4remtype = MakeNative(func(__e Evaluator, __args ...Obj) {
		V701 := __args[0]
		_ = V701
		reg16596 := MakeSymbol("shen.*signedfuncs*")
		reg16597 := MakeSymbol("shen.*signedfuncs*")
		reg16598 := PrimValue(reg16597)
		reg16599 := Call(__e, __defun__shen_4removetype, V701, reg16598)
		reg16600 := PrimSet(reg16596, reg16599)
		__e.Return(reg16600)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.remtype", value: __defun__shen_4remtype})

	__defun__shen_4removetype = MakeNative(func(__e Evaluator, __args ...Obj) {
		V709 := __args[0]
		_ = V709
		V710 := __args[1]
		_ = V710
		reg16601 := Nil
		reg16602 := PrimEqual(reg16601, V710)
		if reg16602 == True {
			reg16603 := Nil
			__e.Return(reg16603)
			return
		} else {
			reg16604 := PrimIsPair(V710)
			var reg16619 Obj
			if reg16604 == True {
				reg16605 := PrimHead(V710)
				reg16606 := PrimIsPair(reg16605)
				var reg16614 Obj
				if reg16606 == True {
					reg16607 := PrimHead(V710)
					reg16608 := PrimHead(reg16607)
					reg16609 := PrimEqual(reg16608, V709)
					var reg16612 Obj
					if reg16609 == True {
						reg16610 := True
						reg16612 = reg16610
					} else {
						reg16611 := False
						reg16612 = reg16611
					}
					reg16614 = reg16612
				} else {
					reg16613 := False
					reg16614 = reg16613
				}
				var reg16617 Obj
				if reg16614 == True {
					reg16615 := True
					reg16617 = reg16615
				} else {
					reg16616 := False
					reg16617 = reg16616
				}
				reg16619 = reg16617
			} else {
				reg16618 := False
				reg16619 = reg16618
			}
			if reg16619 == True {
				reg16620 := PrimHead(V710)
				reg16621 := PrimHead(reg16620)
				reg16622 := PrimTail(V710)
				__e.TailApply(__defun__shen_4removetype, reg16621, reg16622)
				return
			} else {
				reg16624 := PrimIsPair(V710)
				if reg16624 == True {
					reg16625 := PrimHead(V710)
					reg16626 := PrimTail(V710)
					reg16627 := Call(__e, __defun__shen_4removetype, V709, reg16626)
					reg16628 := PrimCons(reg16625, reg16627)
					__e.Return(reg16628)
					return
				} else {
					reg16629 := MakeSymbol("shen.removetype")
					__e.TailApply(__defun__shen_4f__error, reg16629)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.removetype", value: __defun__shen_4removetype})

	__defun__shen_4_5sig_7rest_6 = MakeNative(func(__e Evaluator, __args ...Obj) {
		V712 := __args[0]
		_ = V712
		reg16631 := Call(__e, __defun__shen_4_5signature_6, V712)
		Parse__shen_4_5signature_6 := reg16631
		_ = Parse__shen_4_5signature_6
		reg16632 := Call(__e, __defun__fail)
		reg16633 := PrimEqual(reg16632, Parse__shen_4_5signature_6)
		reg16634 := PrimNot(reg16633)
		if reg16634 == True {
			reg16635 := Call(__e, __defun___5_b_6, Parse__shen_4_5signature_6)
			Parse___5_b_6 := reg16635
			_ = Parse___5_b_6
			reg16636 := Call(__e, __defun__fail)
			reg16637 := PrimEqual(reg16636, Parse___5_b_6)
			reg16638 := PrimNot(reg16637)
			if reg16638 == True {
				reg16639 := PrimHead(Parse___5_b_6)
				reg16640 := Call(__e, __defun__shen_4hdtl, Parse__shen_4_5signature_6)
				__e.TailApply(__defun__shen_4pair, reg16639, reg16640)
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
	__initDefs = append(__initDefs, defType{name: "shen.<sig+rest>", value: __defun__shen_4_5sig_7rest_6})

	__defun__write_1to_1file = MakeNative(func(__e Evaluator, __args ...Obj) {
		V715 := __args[0]
		_ = V715
		V716 := __args[1]
		_ = V716
		reg16644 := MakeSymbol("out")
		reg16645 := PrimOpenStream(V715, reg16644)
		Stream := reg16645
		_ = Stream
		reg16646 := PrimIsString(V716)
		var reg16653 Obj
		if reg16646 == True {
			reg16647 := MakeString("\n\n")
			reg16648 := MakeSymbol("shen.a")
			reg16649 := Call(__e, __defun__shen_4app, V716, reg16647, reg16648)
			reg16653 = reg16649
		} else {
			reg16650 := MakeString("\n\n")
			reg16651 := MakeSymbol("shen.s")
			reg16652 := Call(__e, __defun__shen_4app, V716, reg16650, reg16651)
			reg16653 = reg16652
		}
		String := reg16653
		_ = String
		reg16654 := Call(__e, __defun__pr, String, Stream)
		Write := reg16654
		_ = Write
		reg16655 := PrimCloseStream(Stream)
		Close := reg16655
		_ = Close
		__e.Return(V716)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "write-to-file", value: __defun__write_1to_1file})

}
