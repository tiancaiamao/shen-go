package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4f__error Obj               // shen.f_error
var __defun__shen_4tracked_2 Obj              // shen.tracked?
var __defun__track Obj                        // track
var __defun__shen_4track_1function Obj        // shen.track-function
var __defun__shen_4insert_1tracking_1code Obj // shen.insert-tracking-code
var __defun__step Obj                         // step
var __defun__spy Obj                          // spy
var __defun__shen_4terpri_1or_1read_1char Obj // shen.terpri-or-read-char
var __defun__shen_4check_1byte Obj            // shen.check-byte
var __defun__shen_4input_1track Obj           // shen.input-track
var __defun__shen_4recursively_1print Obj     // shen.recursively-print
var __defun__shen_4spaces Obj                 // shen.spaces
var __defun__shen_4output_1track Obj          // shen.output-track
var __defun__untrack Obj                      // untrack
var __defun__profile Obj                      // profile
var __defun__shen_4profile_1help Obj          // shen.profile-help
var __defun__unprofile Obj                    // unprofile
var __defun__shen_4profile_1func Obj          // shen.profile-func
var __defun__profile_1results Obj             // profile-results
var __defun__shen_4get_1profile Obj           // shen.get-profile
var __defun__shen_4put_1profile Obj           // shen.put-profile

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg307342 := MakeString("Copyright (c) 2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n1. Redistributions of source code must retain the above copyright\n   notice, this list of conditions and the following disclaimer.\n2. Redistributions in binary form must reproduce the above copyright\n   notice, this list of conditions and the following disclaimer in the\n   documentation and/or other materials provided with the distribution.\n3. The name of Mark Tarver may not be used to endorse or promote products\n   derived from this software without specific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY\nEXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY\nDIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\n(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\nLOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\nON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\n(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\nSOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.")
		__ctx.Return(reg307342)
		return
	}, 0))
	__defun__shen_4f__error = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4048 := __args[0]
		_ = V4048
		reg307343 := MakeString("partial function ")
		reg307344 := MakeString(";\n")
		reg307345 := MakeSymbol("shen.a")
		reg307346 := __e.Call(__defun__shen_4app, V4048, reg307344, reg307345)
		reg307347 := PrimStringConcat(reg307343, reg307346)
		reg307348 := __e.Call(__defun__stoutput)
		reg307349 := __e.Call(__defun__shen_4prhush, reg307347, reg307348)
		_ = reg307349
		reg307350 := __e.Call(__defun__shen_4tracked_2, V4048)
		reg307351 := PrimNot(reg307350)
		var reg307362 Obj
		if reg307351 == True {
			reg307352 := MakeString("track ")
			reg307353 := MakeString("? ")
			reg307354 := MakeSymbol("shen.a")
			reg307355 := __e.Call(__defun__shen_4app, V4048, reg307353, reg307354)
			reg307356 := PrimStringConcat(reg307352, reg307355)
			reg307357 := __e.Call(__defun__y_1or_1n_2, reg307356)
			var reg307360 Obj
			if reg307357 == True {
				reg307358 := True
				reg307360 = reg307358
			} else {
				reg307359 := False
				reg307360 = reg307359
			}
			reg307362 = reg307360
		} else {
			reg307361 := False
			reg307362 = reg307361
		}
		var reg307366 Obj
		if reg307362 == True {
			reg307363 := __e.Call(__defun__ps, V4048)
			reg307364 := __e.Call(__defun__shen_4track_1function, reg307363)
			reg307366 = reg307364
		} else {
			reg307365 := MakeSymbol("shen.ok")
			reg307366 = reg307365
		}
		_ = reg307366
		reg307367 := MakeString("aborted")
		reg307368 := PrimSimpleError(reg307367)
		__ctx.Return(reg307368)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.f_error", value: __defun__shen_4f__error})

	__defun__shen_4tracked_2 = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4050 := __args[0]
		_ = V4050
		reg307369 := MakeSymbol("shen.*tracking*")
		reg307370 := PrimValue(reg307369)
		__ctx.TailApply(__defun__element_2, V4050, reg307370)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.tracked?", value: __defun__shen_4tracked_2})

	__defun__track = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4052 := __args[0]
		_ = V4052
		reg307372 := __e.Call(__defun__ps, V4052)
		Source := reg307372
		_ = Source
		__ctx.TailApply(__defun__shen_4track_1function, Source)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "track", value: __defun__track})

	__defun__shen_4track_1function = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4054 := __args[0]
		_ = V4054
		reg307374 := PrimIsPair(V4054)
		var reg307417 Obj
		if reg307374 == True {
			reg307375 := MakeSymbol("defun")
			reg307376 := PrimHead(V4054)
			reg307377 := PrimEqual(reg307375, reg307376)
			var reg307412 Obj
			if reg307377 == True {
				reg307378 := PrimTail(V4054)
				reg307379 := PrimIsPair(reg307378)
				var reg307407 Obj
				if reg307379 == True {
					reg307380 := PrimTail(V4054)
					reg307381 := PrimTail(reg307380)
					reg307382 := PrimIsPair(reg307381)
					var reg307402 Obj
					if reg307382 == True {
						reg307383 := PrimTail(V4054)
						reg307384 := PrimTail(reg307383)
						reg307385 := PrimTail(reg307384)
						reg307386 := PrimIsPair(reg307385)
						var reg307397 Obj
						if reg307386 == True {
							reg307387 := Nil
							reg307388 := PrimTail(V4054)
							reg307389 := PrimTail(reg307388)
							reg307390 := PrimTail(reg307389)
							reg307391 := PrimTail(reg307390)
							reg307392 := PrimEqual(reg307387, reg307391)
							var reg307395 Obj
							if reg307392 == True {
								reg307393 := True
								reg307395 = reg307393
							} else {
								reg307394 := False
								reg307395 = reg307394
							}
							reg307397 = reg307395
						} else {
							reg307396 := False
							reg307397 = reg307396
						}
						var reg307400 Obj
						if reg307397 == True {
							reg307398 := True
							reg307400 = reg307398
						} else {
							reg307399 := False
							reg307400 = reg307399
						}
						reg307402 = reg307400
					} else {
						reg307401 := False
						reg307402 = reg307401
					}
					var reg307405 Obj
					if reg307402 == True {
						reg307403 := True
						reg307405 = reg307403
					} else {
						reg307404 := False
						reg307405 = reg307404
					}
					reg307407 = reg307405
				} else {
					reg307406 := False
					reg307407 = reg307406
				}
				var reg307410 Obj
				if reg307407 == True {
					reg307408 := True
					reg307410 = reg307408
				} else {
					reg307409 := False
					reg307410 = reg307409
				}
				reg307412 = reg307410
			} else {
				reg307411 := False
				reg307412 = reg307411
			}
			var reg307415 Obj
			if reg307412 == True {
				reg307413 := True
				reg307415 = reg307413
			} else {
				reg307414 := False
				reg307415 = reg307414
			}
			reg307417 = reg307415
		} else {
			reg307416 := False
			reg307417 = reg307416
		}
		if reg307417 == True {
			reg307418 := MakeSymbol("defun")
			reg307419 := PrimTail(V4054)
			reg307420 := PrimHead(reg307419)
			reg307421 := PrimTail(V4054)
			reg307422 := PrimTail(reg307421)
			reg307423 := PrimHead(reg307422)
			reg307424 := PrimTail(V4054)
			reg307425 := PrimHead(reg307424)
			reg307426 := PrimTail(V4054)
			reg307427 := PrimTail(reg307426)
			reg307428 := PrimHead(reg307427)
			reg307429 := PrimTail(V4054)
			reg307430 := PrimTail(reg307429)
			reg307431 := PrimTail(reg307430)
			reg307432 := PrimHead(reg307431)
			reg307433 := __e.Call(__defun__shen_4insert_1tracking_1code, reg307425, reg307428, reg307432)
			reg307434 := Nil
			reg307435 := PrimCons(reg307433, reg307434)
			reg307436 := PrimCons(reg307423, reg307435)
			reg307437 := PrimCons(reg307420, reg307436)
			reg307438 := PrimCons(reg307418, reg307437)
			KL := reg307438
			_ = KL
			reg307439 := PrimEvalKL(__e, KL)
			Ob := reg307439
			_ = Ob
			reg307440 := MakeSymbol("shen.*tracking*")
			reg307441 := MakeSymbol("shen.*tracking*")
			reg307442 := PrimValue(reg307441)
			reg307443 := PrimCons(Ob, reg307442)
			reg307444 := PrimSet(reg307440, reg307443)
			Tr := reg307444
			_ = Tr
			__ctx.Return(Ob)
			return
		} else {
			reg307445 := MakeSymbol("shen.track-function")
			__ctx.TailApply(__defun__shen_4f__error, reg307445)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.track-function", value: __defun__shen_4track_1function})

	__defun__shen_4insert_1tracking_1code = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4058 := __args[0]
		_ = V4058
		V4059 := __args[1]
		_ = V4059
		V4060 := __args[2]
		_ = V4060
		reg307447 := MakeSymbol("do")
		reg307448 := MakeSymbol("set")
		reg307449 := MakeSymbol("shen.*call*")
		reg307450 := MakeSymbol("+")
		reg307451 := MakeSymbol("value")
		reg307452 := MakeSymbol("shen.*call*")
		reg307453 := Nil
		reg307454 := PrimCons(reg307452, reg307453)
		reg307455 := PrimCons(reg307451, reg307454)
		reg307456 := MakeNumber(1)
		reg307457 := Nil
		reg307458 := PrimCons(reg307456, reg307457)
		reg307459 := PrimCons(reg307455, reg307458)
		reg307460 := PrimCons(reg307450, reg307459)
		reg307461 := Nil
		reg307462 := PrimCons(reg307460, reg307461)
		reg307463 := PrimCons(reg307449, reg307462)
		reg307464 := PrimCons(reg307448, reg307463)
		reg307465 := MakeSymbol("do")
		reg307466 := MakeSymbol("shen.input-track")
		reg307467 := MakeSymbol("value")
		reg307468 := MakeSymbol("shen.*call*")
		reg307469 := Nil
		reg307470 := PrimCons(reg307468, reg307469)
		reg307471 := PrimCons(reg307467, reg307470)
		reg307472 := __e.Call(__defun__shen_4cons__form, V4059)
		reg307473 := Nil
		reg307474 := PrimCons(reg307472, reg307473)
		reg307475 := PrimCons(V4058, reg307474)
		reg307476 := PrimCons(reg307471, reg307475)
		reg307477 := PrimCons(reg307466, reg307476)
		reg307478 := MakeSymbol("do")
		reg307479 := MakeSymbol("shen.terpri-or-read-char")
		reg307480 := Nil
		reg307481 := PrimCons(reg307479, reg307480)
		reg307482 := MakeSymbol("let")
		reg307483 := MakeSymbol("Result")
		reg307484 := MakeSymbol("do")
		reg307485 := MakeSymbol("shen.output-track")
		reg307486 := MakeSymbol("value")
		reg307487 := MakeSymbol("shen.*call*")
		reg307488 := Nil
		reg307489 := PrimCons(reg307487, reg307488)
		reg307490 := PrimCons(reg307486, reg307489)
		reg307491 := MakeSymbol("Result")
		reg307492 := Nil
		reg307493 := PrimCons(reg307491, reg307492)
		reg307494 := PrimCons(V4058, reg307493)
		reg307495 := PrimCons(reg307490, reg307494)
		reg307496 := PrimCons(reg307485, reg307495)
		reg307497 := MakeSymbol("do")
		reg307498 := MakeSymbol("set")
		reg307499 := MakeSymbol("shen.*call*")
		reg307500 := MakeSymbol("-")
		reg307501 := MakeSymbol("value")
		reg307502 := MakeSymbol("shen.*call*")
		reg307503 := Nil
		reg307504 := PrimCons(reg307502, reg307503)
		reg307505 := PrimCons(reg307501, reg307504)
		reg307506 := MakeNumber(1)
		reg307507 := Nil
		reg307508 := PrimCons(reg307506, reg307507)
		reg307509 := PrimCons(reg307505, reg307508)
		reg307510 := PrimCons(reg307500, reg307509)
		reg307511 := Nil
		reg307512 := PrimCons(reg307510, reg307511)
		reg307513 := PrimCons(reg307499, reg307512)
		reg307514 := PrimCons(reg307498, reg307513)
		reg307515 := MakeSymbol("do")
		reg307516 := MakeSymbol("shen.terpri-or-read-char")
		reg307517 := Nil
		reg307518 := PrimCons(reg307516, reg307517)
		reg307519 := MakeSymbol("Result")
		reg307520 := Nil
		reg307521 := PrimCons(reg307519, reg307520)
		reg307522 := PrimCons(reg307518, reg307521)
		reg307523 := PrimCons(reg307515, reg307522)
		reg307524 := Nil
		reg307525 := PrimCons(reg307523, reg307524)
		reg307526 := PrimCons(reg307514, reg307525)
		reg307527 := PrimCons(reg307497, reg307526)
		reg307528 := Nil
		reg307529 := PrimCons(reg307527, reg307528)
		reg307530 := PrimCons(reg307496, reg307529)
		reg307531 := PrimCons(reg307484, reg307530)
		reg307532 := Nil
		reg307533 := PrimCons(reg307531, reg307532)
		reg307534 := PrimCons(V4060, reg307533)
		reg307535 := PrimCons(reg307483, reg307534)
		reg307536 := PrimCons(reg307482, reg307535)
		reg307537 := Nil
		reg307538 := PrimCons(reg307536, reg307537)
		reg307539 := PrimCons(reg307481, reg307538)
		reg307540 := PrimCons(reg307478, reg307539)
		reg307541 := Nil
		reg307542 := PrimCons(reg307540, reg307541)
		reg307543 := PrimCons(reg307477, reg307542)
		reg307544 := PrimCons(reg307465, reg307543)
		reg307545 := Nil
		reg307546 := PrimCons(reg307544, reg307545)
		reg307547 := PrimCons(reg307464, reg307546)
		reg307548 := PrimCons(reg307447, reg307547)
		__ctx.Return(reg307548)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.insert-tracking-code", value: __defun__shen_4insert_1tracking_1code})

	__initExprs = append(__initExprs, MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg307549 := MakeSymbol("shen.*step*")
		reg307550 := False
		reg307551 := PrimSet(reg307549, reg307550)
		__ctx.Return(reg307551)
		return
	}, 0))
	__defun__step = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4066 := __args[0]
		_ = V4066
		reg307552 := MakeSymbol("+")
		reg307553 := PrimEqual(reg307552, V4066)
		if reg307553 == True {
			reg307554 := MakeSymbol("shen.*step*")
			reg307555 := True
			reg307556 := PrimSet(reg307554, reg307555)
			__ctx.Return(reg307556)
			return
		} else {
			reg307557 := MakeSymbol("-")
			reg307558 := PrimEqual(reg307557, V4066)
			if reg307558 == True {
				reg307559 := MakeSymbol("shen.*step*")
				reg307560 := False
				reg307561 := PrimSet(reg307559, reg307560)
				__ctx.Return(reg307561)
				return
			} else {
				reg307562 := MakeString("step expects a + or a -.\n")
				reg307563 := PrimSimpleError(reg307562)
				__ctx.Return(reg307563)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "step", value: __defun__step})

	__defun__spy = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4072 := __args[0]
		_ = V4072
		reg307564 := MakeSymbol("+")
		reg307565 := PrimEqual(reg307564, V4072)
		if reg307565 == True {
			reg307566 := MakeSymbol("shen.*spy*")
			reg307567 := True
			reg307568 := PrimSet(reg307566, reg307567)
			__ctx.Return(reg307568)
			return
		} else {
			reg307569 := MakeSymbol("-")
			reg307570 := PrimEqual(reg307569, V4072)
			if reg307570 == True {
				reg307571 := MakeSymbol("shen.*spy*")
				reg307572 := False
				reg307573 := PrimSet(reg307571, reg307572)
				__ctx.Return(reg307573)
				return
			} else {
				reg307574 := MakeString("spy expects a + or a -.\n")
				reg307575 := PrimSimpleError(reg307574)
				__ctx.Return(reg307575)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "spy", value: __defun__spy})

	__defun__shen_4terpri_1or_1read_1char = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg307576 := MakeSymbol("shen.*step*")
		reg307577 := PrimValue(reg307576)
		if reg307577 == True {
			reg307578 := MakeSymbol("*stinput*")
			reg307579 := PrimValue(reg307578)
			reg307580 := PrimReadByte(reg307579)
			__ctx.TailApply(__defun__shen_4check_1byte, reg307580)
			return
		} else {
			reg307582 := MakeNumber(1)
			__ctx.TailApply(__defun__nl, reg307582)
			return
		}
	}, 0)
	__initDefs = append(__initDefs, defType{name: "shen.terpri-or-read-char", value: __defun__shen_4terpri_1or_1read_1char})

	__defun__shen_4check_1byte = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4078 := __args[0]
		_ = V4078
		reg307584 := __e.Call(__defun__shen_4hat)
		reg307585 := PrimEqual(V4078, reg307584)
		if reg307585 == True {
			reg307586 := MakeString("aborted")
			reg307587 := PrimSimpleError(reg307586)
			__ctx.Return(reg307587)
			return
		} else {
			reg307588 := True
			__ctx.Return(reg307588)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.check-byte", value: __defun__shen_4check_1byte})

	__defun__shen_4input_1track = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4082 := __args[0]
		_ = V4082
		V4083 := __args[1]
		_ = V4083
		V4084 := __args[2]
		_ = V4084
		reg307589 := MakeString("\n")
		reg307590 := __e.Call(__defun__shen_4spaces, V4082)
		reg307591 := MakeString("<")
		reg307592 := MakeString("> Inputs to ")
		reg307593 := MakeString(" \n")
		reg307594 := __e.Call(__defun__shen_4spaces, V4082)
		reg307595 := MakeString("")
		reg307596 := MakeSymbol("shen.a")
		reg307597 := __e.Call(__defun__shen_4app, reg307594, reg307595, reg307596)
		reg307598 := PrimStringConcat(reg307593, reg307597)
		reg307599 := MakeSymbol("shen.a")
		reg307600 := __e.Call(__defun__shen_4app, V4083, reg307598, reg307599)
		reg307601 := PrimStringConcat(reg307592, reg307600)
		reg307602 := MakeSymbol("shen.a")
		reg307603 := __e.Call(__defun__shen_4app, V4082, reg307601, reg307602)
		reg307604 := PrimStringConcat(reg307591, reg307603)
		reg307605 := MakeSymbol("shen.a")
		reg307606 := __e.Call(__defun__shen_4app, reg307590, reg307604, reg307605)
		reg307607 := PrimStringConcat(reg307589, reg307606)
		reg307608 := __e.Call(__defun__stoutput)
		reg307609 := __e.Call(__defun__shen_4prhush, reg307607, reg307608)
		_ = reg307609
		__ctx.TailApply(__defun__shen_4recursively_1print, V4084)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.input-track", value: __defun__shen_4input_1track})

	__defun__shen_4recursively_1print = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4086 := __args[0]
		_ = V4086
		reg307611 := Nil
		reg307612 := PrimEqual(reg307611, V4086)
		if reg307612 == True {
			reg307613 := MakeString(" ==>")
			reg307614 := __e.Call(__defun__stoutput)
			__ctx.TailApply(__defun__shen_4prhush, reg307613, reg307614)
			return
		} else {
			reg307616 := PrimIsPair(V4086)
			if reg307616 == True {
				reg307617 := PrimHead(V4086)
				reg307618 := __e.Call(__defun__print, reg307617)
				_ = reg307618
				reg307619 := MakeString(", ")
				reg307620 := __e.Call(__defun__stoutput)
				reg307621 := __e.Call(__defun__shen_4prhush, reg307619, reg307620)
				_ = reg307621
				reg307622 := PrimTail(V4086)
				__ctx.TailApply(__defun__shen_4recursively_1print, reg307622)
				return
			} else {
				reg307624 := MakeSymbol("shen.recursively-print")
				__ctx.TailApply(__defun__shen_4f__error, reg307624)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.recursively-print", value: __defun__shen_4recursively_1print})

	__defun__shen_4spaces = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4088 := __args[0]
		_ = V4088
		reg307626 := MakeNumber(0)
		reg307627 := PrimEqual(reg307626, V4088)
		if reg307627 == True {
			reg307628 := MakeString("")
			__ctx.Return(reg307628)
			return
		} else {
			reg307629 := MakeString(" ")
			reg307630 := MakeNumber(1)
			reg307631 := PrimNumberSubtract(V4088, reg307630)
			reg307632 := __e.Call(__defun__shen_4spaces, reg307631)
			reg307633 := PrimStringConcat(reg307629, reg307632)
			__ctx.Return(reg307633)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.spaces", value: __defun__shen_4spaces})

	__defun__shen_4output_1track = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4092 := __args[0]
		_ = V4092
		V4093 := __args[1]
		_ = V4093
		V4094 := __args[2]
		_ = V4094
		reg307634 := MakeString("\n")
		reg307635 := __e.Call(__defun__shen_4spaces, V4092)
		reg307636 := MakeString("<")
		reg307637 := MakeString("> Output from ")
		reg307638 := MakeString(" \n")
		reg307639 := __e.Call(__defun__shen_4spaces, V4092)
		reg307640 := MakeString("==> ")
		reg307641 := MakeString("")
		reg307642 := MakeSymbol("shen.s")
		reg307643 := __e.Call(__defun__shen_4app, V4094, reg307641, reg307642)
		reg307644 := PrimStringConcat(reg307640, reg307643)
		reg307645 := MakeSymbol("shen.a")
		reg307646 := __e.Call(__defun__shen_4app, reg307639, reg307644, reg307645)
		reg307647 := PrimStringConcat(reg307638, reg307646)
		reg307648 := MakeSymbol("shen.a")
		reg307649 := __e.Call(__defun__shen_4app, V4093, reg307647, reg307648)
		reg307650 := PrimStringConcat(reg307637, reg307649)
		reg307651 := MakeSymbol("shen.a")
		reg307652 := __e.Call(__defun__shen_4app, V4092, reg307650, reg307651)
		reg307653 := PrimStringConcat(reg307636, reg307652)
		reg307654 := MakeSymbol("shen.a")
		reg307655 := __e.Call(__defun__shen_4app, reg307635, reg307653, reg307654)
		reg307656 := PrimStringConcat(reg307634, reg307655)
		reg307657 := __e.Call(__defun__stoutput)
		__ctx.TailApply(__defun__shen_4prhush, reg307656, reg307657)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.output-track", value: __defun__shen_4output_1track})

	__defun__untrack = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4096 := __args[0]
		_ = V4096
		reg307659 := MakeSymbol("shen.*tracking*")
		reg307660 := PrimValue(reg307659)
		Tracking := reg307660
		_ = Tracking
		reg307661 := MakeSymbol("shen.*tracking*")
		reg307662 := __e.Call(__defun__remove, V4096, Tracking)
		reg307663 := PrimSet(reg307661, reg307662)
		Tracking = reg307663
		_ = Tracking
		reg307664 := __e.Call(__defun__ps, V4096)
		__ctx.TailApply(__defun__eval, reg307664)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "untrack", value: __defun__untrack})

	__defun__profile = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4098 := __args[0]
		_ = V4098
		reg307666 := __e.Call(__defun__ps, V4098)
		__ctx.TailApply(__defun__shen_4profile_1help, reg307666)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "profile", value: __defun__profile})

	__defun__shen_4profile_1help = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4104 := __args[0]
		_ = V4104
		reg307668 := PrimIsPair(V4104)
		var reg307711 Obj
		if reg307668 == True {
			reg307669 := MakeSymbol("defun")
			reg307670 := PrimHead(V4104)
			reg307671 := PrimEqual(reg307669, reg307670)
			var reg307706 Obj
			if reg307671 == True {
				reg307672 := PrimTail(V4104)
				reg307673 := PrimIsPair(reg307672)
				var reg307701 Obj
				if reg307673 == True {
					reg307674 := PrimTail(V4104)
					reg307675 := PrimTail(reg307674)
					reg307676 := PrimIsPair(reg307675)
					var reg307696 Obj
					if reg307676 == True {
						reg307677 := PrimTail(V4104)
						reg307678 := PrimTail(reg307677)
						reg307679 := PrimTail(reg307678)
						reg307680 := PrimIsPair(reg307679)
						var reg307691 Obj
						if reg307680 == True {
							reg307681 := Nil
							reg307682 := PrimTail(V4104)
							reg307683 := PrimTail(reg307682)
							reg307684 := PrimTail(reg307683)
							reg307685 := PrimTail(reg307684)
							reg307686 := PrimEqual(reg307681, reg307685)
							var reg307689 Obj
							if reg307686 == True {
								reg307687 := True
								reg307689 = reg307687
							} else {
								reg307688 := False
								reg307689 = reg307688
							}
							reg307691 = reg307689
						} else {
							reg307690 := False
							reg307691 = reg307690
						}
						var reg307694 Obj
						if reg307691 == True {
							reg307692 := True
							reg307694 = reg307692
						} else {
							reg307693 := False
							reg307694 = reg307693
						}
						reg307696 = reg307694
					} else {
						reg307695 := False
						reg307696 = reg307695
					}
					var reg307699 Obj
					if reg307696 == True {
						reg307697 := True
						reg307699 = reg307697
					} else {
						reg307698 := False
						reg307699 = reg307698
					}
					reg307701 = reg307699
				} else {
					reg307700 := False
					reg307701 = reg307700
				}
				var reg307704 Obj
				if reg307701 == True {
					reg307702 := True
					reg307704 = reg307702
				} else {
					reg307703 := False
					reg307704 = reg307703
				}
				reg307706 = reg307704
			} else {
				reg307705 := False
				reg307706 = reg307705
			}
			var reg307709 Obj
			if reg307706 == True {
				reg307707 := True
				reg307709 = reg307707
			} else {
				reg307708 := False
				reg307709 = reg307708
			}
			reg307711 = reg307709
		} else {
			reg307710 := False
			reg307711 = reg307710
		}
		if reg307711 == True {
			reg307712 := MakeSymbol("shen.f")
			reg307713 := __e.Call(__defun__gensym, reg307712)
			G := reg307713
			_ = G
			reg307714 := MakeSymbol("defun")
			reg307715 := PrimTail(V4104)
			reg307716 := PrimHead(reg307715)
			reg307717 := PrimTail(V4104)
			reg307718 := PrimTail(reg307717)
			reg307719 := PrimHead(reg307718)
			reg307720 := PrimTail(V4104)
			reg307721 := PrimHead(reg307720)
			reg307722 := PrimTail(V4104)
			reg307723 := PrimTail(reg307722)
			reg307724 := PrimHead(reg307723)
			reg307725 := PrimTail(V4104)
			reg307726 := PrimTail(reg307725)
			reg307727 := PrimHead(reg307726)
			reg307728 := PrimCons(G, reg307727)
			reg307729 := __e.Call(__defun__shen_4profile_1func, reg307721, reg307724, reg307728)
			reg307730 := Nil
			reg307731 := PrimCons(reg307729, reg307730)
			reg307732 := PrimCons(reg307719, reg307731)
			reg307733 := PrimCons(reg307716, reg307732)
			reg307734 := PrimCons(reg307714, reg307733)
			Profile := reg307734
			_ = Profile
			reg307735 := MakeSymbol("defun")
			reg307736 := PrimTail(V4104)
			reg307737 := PrimTail(reg307736)
			reg307738 := PrimHead(reg307737)
			reg307739 := PrimTail(V4104)
			reg307740 := PrimHead(reg307739)
			reg307741 := PrimTail(V4104)
			reg307742 := PrimTail(reg307741)
			reg307743 := PrimTail(reg307742)
			reg307744 := PrimHead(reg307743)
			reg307745 := __e.Call(__defun__subst, G, reg307740, reg307744)
			reg307746 := Nil
			reg307747 := PrimCons(reg307745, reg307746)
			reg307748 := PrimCons(reg307738, reg307747)
			reg307749 := PrimCons(G, reg307748)
			reg307750 := PrimCons(reg307735, reg307749)
			Def := reg307750
			_ = Def
			reg307751 := __e.Call(__defun__shen_4eval_1without_1macros, Profile)
			CompileProfile := reg307751
			_ = CompileProfile
			reg307752 := __e.Call(__defun__shen_4eval_1without_1macros, Def)
			CompileG := reg307752
			_ = CompileG
			reg307753 := PrimTail(V4104)
			reg307754 := PrimHead(reg307753)
			__ctx.Return(reg307754)
			return
		} else {
			reg307755 := MakeString("Cannot profile.\n")
			reg307756 := PrimSimpleError(reg307755)
			__ctx.Return(reg307756)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.profile-help", value: __defun__shen_4profile_1help})

	__defun__unprofile = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4106 := __args[0]
		_ = V4106
		__ctx.TailApply(__defun__untrack, V4106)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "unprofile", value: __defun__unprofile})

	__defun__shen_4profile_1func = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4110 := __args[0]
		_ = V4110
		V4111 := __args[1]
		_ = V4111
		V4112 := __args[2]
		_ = V4112
		reg307758 := MakeSymbol("let")
		reg307759 := MakeSymbol("Start")
		reg307760 := MakeSymbol("get-time")
		reg307761 := MakeSymbol("run")
		reg307762 := Nil
		reg307763 := PrimCons(reg307761, reg307762)
		reg307764 := PrimCons(reg307760, reg307763)
		reg307765 := MakeSymbol("let")
		reg307766 := MakeSymbol("Result")
		reg307767 := MakeSymbol("let")
		reg307768 := MakeSymbol("Finish")
		reg307769 := MakeSymbol("-")
		reg307770 := MakeSymbol("get-time")
		reg307771 := MakeSymbol("run")
		reg307772 := Nil
		reg307773 := PrimCons(reg307771, reg307772)
		reg307774 := PrimCons(reg307770, reg307773)
		reg307775 := MakeSymbol("Start")
		reg307776 := Nil
		reg307777 := PrimCons(reg307775, reg307776)
		reg307778 := PrimCons(reg307774, reg307777)
		reg307779 := PrimCons(reg307769, reg307778)
		reg307780 := MakeSymbol("let")
		reg307781 := MakeSymbol("Record")
		reg307782 := MakeSymbol("shen.put-profile")
		reg307783 := MakeSymbol("+")
		reg307784 := MakeSymbol("shen.get-profile")
		reg307785 := Nil
		reg307786 := PrimCons(V4110, reg307785)
		reg307787 := PrimCons(reg307784, reg307786)
		reg307788 := MakeSymbol("Finish")
		reg307789 := Nil
		reg307790 := PrimCons(reg307788, reg307789)
		reg307791 := PrimCons(reg307787, reg307790)
		reg307792 := PrimCons(reg307783, reg307791)
		reg307793 := Nil
		reg307794 := PrimCons(reg307792, reg307793)
		reg307795 := PrimCons(V4110, reg307794)
		reg307796 := PrimCons(reg307782, reg307795)
		reg307797 := MakeSymbol("Result")
		reg307798 := Nil
		reg307799 := PrimCons(reg307797, reg307798)
		reg307800 := PrimCons(reg307796, reg307799)
		reg307801 := PrimCons(reg307781, reg307800)
		reg307802 := PrimCons(reg307780, reg307801)
		reg307803 := Nil
		reg307804 := PrimCons(reg307802, reg307803)
		reg307805 := PrimCons(reg307779, reg307804)
		reg307806 := PrimCons(reg307768, reg307805)
		reg307807 := PrimCons(reg307767, reg307806)
		reg307808 := Nil
		reg307809 := PrimCons(reg307807, reg307808)
		reg307810 := PrimCons(V4112, reg307809)
		reg307811 := PrimCons(reg307766, reg307810)
		reg307812 := PrimCons(reg307765, reg307811)
		reg307813 := Nil
		reg307814 := PrimCons(reg307812, reg307813)
		reg307815 := PrimCons(reg307764, reg307814)
		reg307816 := PrimCons(reg307759, reg307815)
		reg307817 := PrimCons(reg307758, reg307816)
		__ctx.Return(reg307817)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.profile-func", value: __defun__shen_4profile_1func})

	__defun__profile_1results = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4114 := __args[0]
		_ = V4114
		reg307818 := __e.Call(__defun__shen_4get_1profile, V4114)
		Results := reg307818
		_ = Results
		reg307819 := MakeNumber(0)
		reg307820 := __e.Call(__defun__shen_4put_1profile, V4114, reg307819)
		Initialise := reg307820
		_ = Initialise
		__ctx.TailApply(__defun___8p, V4114, Results)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "profile-results", value: __defun__profile_1results})

	__defun__shen_4get_1profile = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4116 := __args[0]
		_ = V4116
		reg307822 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg307823 := MakeSymbol("profile")
			reg307824 := MakeSymbol("*property-vector*")
			reg307825 := PrimValue(reg307824)
			__ctx.TailApply(__defun__get, V4116, reg307823, reg307825)
			return
		}, 0)
		reg307827 := MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg307828 := MakeNumber(0)
			__ctx.Return(reg307828)
			return
		}, 1)
		reg307829 := __e.Try(reg307822).Catch(reg307827)
		__ctx.Return(reg307829)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.get-profile", value: __defun__shen_4get_1profile})

	__defun__shen_4put_1profile = MakeNative(func(__e *Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4119 := __args[0]
		_ = V4119
		V4120 := __args[1]
		_ = V4120
		reg307830 := MakeSymbol("profile")
		reg307831 := MakeSymbol("*property-vector*")
		reg307832 := PrimValue(reg307831)
		__ctx.TailApply(__defun__put, V4119, reg307830, V4120, reg307832)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.put-profile", value: __defun__shen_4put_1profile})

}
