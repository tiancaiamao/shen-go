package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__shen_4yacc Obj                          // shen.yacc
var __defun__shen_4yacc_1_6shen Obj                  // shen.yacc->shen
var __defun__shen_4kill_1code Obj                    // shen.kill-code
var __defun__kill Obj                                // kill
var __defun__shen_4analyse_1kill Obj                 // shen.analyse-kill
var __defun__shen_4split__cc__rules Obj              // shen.split_cc_rules
var __defun__shen_4split__cc__rule Obj               // shen.split_cc_rule
var __defun__shen_4semantic_1completion_1warning Obj // shen.semantic-completion-warning
var __defun__shen_4default__semantics Obj            // shen.default_semantics
var __defun__shen_4grammar__symbol_2 Obj             // shen.grammar_symbol?
var __defun__shen_4yacc__cases Obj                   // shen.yacc_cases
var __defun__shen_4cc__body Obj                      // shen.cc_body
var __defun__shen_4syntax Obj                        // shen.syntax
var __defun__shen_4list_1stream Obj                  // shen.list-stream
var __defun__shen_4decons Obj                        // shen.decons
var __defun__shen_4insert_1runon Obj                 // shen.insert-runon
var __defun__shen_4strip_1pathname Obj               // shen.strip-pathname
var __defun__shen_4recursive__descent Obj            // shen.recursive_descent
var __defun__shen_4variable_1match Obj               // shen.variable-match
var __defun__shen_4terminal_2 Obj                    // shen.terminal?
var __defun__shen_4jump__stream_2 Obj                // shen.jump_stream?
var __defun__shen_4check__stream Obj                 // shen.check_stream
var __defun__shen_4jump__stream Obj                  // shen.jump_stream
var __defun__shen_4semantics Obj                     // shen.semantics
var __defun__shen_4pair Obj                          // shen.pair
var __defun__shen_4hdtl Obj                          // shen.hdtl
var __defun__shen_4hdhd Obj                          // shen.hdhd
var __defun__shen_4tlhd Obj                          // shen.tlhd
var __defun__shen_4snd_1or_1fail Obj                 // shen.snd-or-fail
var __defun__fail Obj                                // fail
var __defun___5_b_6 Obj                              // <!>
var __defun___5e_6 Obj                               // <e>

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg7394 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__ctx.Return(reg7394)
		return
	}, 0))
	__defun__shen_4yacc = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4694 := __args[0]
		_ = V4694
		reg7395 := PrimIsPair(V4694)
		var reg7410 Obj
		if reg7395 == True {
			reg7396 := MakeSymbol("defcc")
			reg7397 := PrimHead(V4694)
			reg7398 := PrimEqual(reg7396, reg7397)
			var reg7405 Obj
			if reg7398 == True {
				reg7399 := PrimTail(V4694)
				reg7400 := PrimIsPair(reg7399)
				var reg7403 Obj
				if reg7400 == True {
					reg7401 := True
					reg7403 = reg7401
				} else {
					reg7402 := False
					reg7403 = reg7402
				}
				reg7405 = reg7403
			} else {
				reg7404 := False
				reg7405 = reg7404
			}
			var reg7408 Obj
			if reg7405 == True {
				reg7406 := True
				reg7408 = reg7406
			} else {
				reg7407 := False
				reg7408 = reg7407
			}
			reg7410 = reg7408
		} else {
			reg7409 := False
			reg7410 = reg7409
		}
		if reg7410 == True {
			reg7411 := PrimTail(V4694)
			reg7412 := PrimHead(reg7411)
			reg7413 := PrimTail(V4694)
			reg7414 := PrimTail(reg7413)
			__ctx.TailApply(__defun__shen_4yacc_1_6shen, reg7412, reg7414)
			return
		} else {
			reg7416 := MakeSymbol("shen.yacc")
			__ctx.TailApply(__defun__shen_4f__error, reg7416)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.yacc", value: __defun__shen_4yacc})

	__defun__shen_4yacc_1_6shen = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4697 := __args[0]
		_ = V4697
		V4698 := __args[1]
		_ = V4698
		reg7418 := True
		reg7419 := Nil
		reg7420 := __e.Call(__defun__shen_4split__cc__rules, reg7418, V4698, reg7419)
		CCRules := reg7420
		_ = CCRules
		reg7421 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			X := __args[0]
			_ = X
			__ctx.TailApply(__defun__shen_4cc__body, X)
			return
		}, 1)
		reg7423 := __e.Call(__defun__map, reg7421, CCRules)
		CCBody := reg7423
		_ = CCBody
		reg7424 := __e.Call(__defun__shen_4yacc__cases, CCBody)
		YaccCases := reg7424
		_ = YaccCases
		reg7425 := MakeSymbol("define")
		reg7426 := MakeSymbol("Stream")
		reg7427 := MakeSymbol("->")
		reg7428 := __e.Call(__defun__shen_4kill_1code, YaccCases)
		reg7429 := Nil
		reg7430 := PrimCons(reg7428, reg7429)
		reg7431 := PrimCons(reg7427, reg7430)
		reg7432 := PrimCons(reg7426, reg7431)
		reg7433 := PrimCons(V4697, reg7432)
		reg7434 := PrimCons(reg7425, reg7433)
		__ctx.Return(reg7434)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.yacc->shen", value: __defun__shen_4yacc_1_6shen})

	__defun__shen_4kill_1code = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4700 := __args[0]
		_ = V4700
		reg7435 := MakeSymbol("kill")
		reg7436 := __e.Call(__defun__occurrences, reg7435, V4700)
		reg7437 := MakeNumber(0)
		reg7438 := PrimGreatThan(reg7436, reg7437)
		if reg7438 == True {
			reg7439 := MakeSymbol("trap-error")
			reg7440 := MakeSymbol("lambda")
			reg7441 := MakeSymbol("E")
			reg7442 := MakeSymbol("shen.analyse-kill")
			reg7443 := MakeSymbol("E")
			reg7444 := Nil
			reg7445 := PrimCons(reg7443, reg7444)
			reg7446 := PrimCons(reg7442, reg7445)
			reg7447 := Nil
			reg7448 := PrimCons(reg7446, reg7447)
			reg7449 := PrimCons(reg7441, reg7448)
			reg7450 := PrimCons(reg7440, reg7449)
			reg7451 := Nil
			reg7452 := PrimCons(reg7450, reg7451)
			reg7453 := PrimCons(V4700, reg7452)
			reg7454 := PrimCons(reg7439, reg7453)
			__ctx.Return(reg7454)
			return
		} else {
			__ctx.Return(V4700)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.kill-code", value: __defun__shen_4kill_1code})

	__defun__kill = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg7455 := MakeString("yacc kill")
		reg7456 := PrimSimpleError(reg7455)
		__ctx.Return(reg7456)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "kill", value: __defun__kill})

	__defun__shen_4analyse_1kill = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4702 := __args[0]
		_ = V4702
		reg7457 := PrimErrorToString(V4702)
		String := reg7457
		_ = String
		reg7458 := MakeString("yacc kill")
		reg7459 := PrimEqual(String, reg7458)
		if reg7459 == True {
			__ctx.TailApply(__defun__fail)
			return
		} else {
			__ctx.Return(V4702)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.analyse-kill", value: __defun__shen_4analyse_1kill})

	__defun__shen_4split__cc__rules = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4708 := __args[0]
		_ = V4708
		V4709 := __args[1]
		_ = V4709
		V4710 := __args[2]
		_ = V4710
		reg7461 := Nil
		reg7462 := PrimEqual(reg7461, V4709)
		var reg7469 Obj
		if reg7462 == True {
			reg7463 := Nil
			reg7464 := PrimEqual(reg7463, V4710)
			var reg7467 Obj
			if reg7464 == True {
				reg7465 := True
				reg7467 = reg7465
			} else {
				reg7466 := False
				reg7467 = reg7466
			}
			reg7469 = reg7467
		} else {
			reg7468 := False
			reg7469 = reg7468
		}
		if reg7469 == True {
			reg7470 := Nil
			__ctx.Return(reg7470)
			return
		} else {
			reg7471 := Nil
			reg7472 := PrimEqual(reg7471, V4709)
			if reg7472 == True {
				reg7473 := __e.Call(__defun__reverse, V4710)
				reg7474 := Nil
				reg7475 := __e.Call(__defun__shen_4split__cc__rule, V4708, reg7473, reg7474)
				reg7476 := Nil
				reg7477 := PrimCons(reg7475, reg7476)
				__ctx.Return(reg7477)
				return
			} else {
				reg7478 := PrimIsPair(V4709)
				var reg7486 Obj
				if reg7478 == True {
					reg7479 := MakeSymbol(";")
					reg7480 := PrimHead(V4709)
					reg7481 := PrimEqual(reg7479, reg7480)
					var reg7484 Obj
					if reg7481 == True {
						reg7482 := True
						reg7484 = reg7482
					} else {
						reg7483 := False
						reg7484 = reg7483
					}
					reg7486 = reg7484
				} else {
					reg7485 := False
					reg7486 = reg7485
				}
				if reg7486 == True {
					reg7487 := __e.Call(__defun__reverse, V4710)
					reg7488 := Nil
					reg7489 := __e.Call(__defun__shen_4split__cc__rule, V4708, reg7487, reg7488)
					reg7490 := PrimTail(V4709)
					reg7491 := Nil
					reg7492 := __e.Call(__defun__shen_4split__cc__rules, V4708, reg7490, reg7491)
					reg7493 := PrimCons(reg7489, reg7492)
					__ctx.Return(reg7493)
					return
				} else {
					reg7494 := PrimIsPair(V4709)
					if reg7494 == True {
						reg7495 := PrimTail(V4709)
						reg7496 := PrimHead(V4709)
						reg7497 := PrimCons(reg7496, V4710)
						__ctx.TailApply(__defun__shen_4split__cc__rules, V4708, reg7495, reg7497)
						return
					} else {
						reg7499 := MakeSymbol("shen.split_cc_rules")
						__ctx.TailApply(__defun__shen_4f__error, reg7499)
						return
					}
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.split_cc_rules", value: __defun__shen_4split__cc__rules})

	__defun__shen_4split__cc__rule = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4718 := __args[0]
		_ = V4718
		V4719 := __args[1]
		_ = V4719
		V4720 := __args[2]
		_ = V4720
		reg7501 := PrimIsPair(V4719)
		var reg7525 Obj
		if reg7501 == True {
			reg7502 := MakeSymbol(":=")
			reg7503 := PrimHead(V4719)
			reg7504 := PrimEqual(reg7502, reg7503)
			var reg7520 Obj
			if reg7504 == True {
				reg7505 := PrimTail(V4719)
				reg7506 := PrimIsPair(reg7505)
				var reg7515 Obj
				if reg7506 == True {
					reg7507 := Nil
					reg7508 := PrimTail(V4719)
					reg7509 := PrimTail(reg7508)
					reg7510 := PrimEqual(reg7507, reg7509)
					var reg7513 Obj
					if reg7510 == True {
						reg7511 := True
						reg7513 = reg7511
					} else {
						reg7512 := False
						reg7513 = reg7512
					}
					reg7515 = reg7513
				} else {
					reg7514 := False
					reg7515 = reg7514
				}
				var reg7518 Obj
				if reg7515 == True {
					reg7516 := True
					reg7518 = reg7516
				} else {
					reg7517 := False
					reg7518 = reg7517
				}
				reg7520 = reg7518
			} else {
				reg7519 := False
				reg7520 = reg7519
			}
			var reg7523 Obj
			if reg7520 == True {
				reg7521 := True
				reg7523 = reg7521
			} else {
				reg7522 := False
				reg7523 = reg7522
			}
			reg7525 = reg7523
		} else {
			reg7524 := False
			reg7525 = reg7524
		}
		if reg7525 == True {
			reg7526 := __e.Call(__defun__reverse, V4720)
			reg7527 := PrimTail(V4719)
			reg7528 := PrimCons(reg7526, reg7527)
			__ctx.Return(reg7528)
			return
		} else {
			reg7529 := PrimIsPair(V4719)
			var reg7582 Obj
			if reg7529 == True {
				reg7530 := MakeSymbol(":=")
				reg7531 := PrimHead(V4719)
				reg7532 := PrimEqual(reg7530, reg7531)
				var reg7577 Obj
				if reg7532 == True {
					reg7533 := PrimTail(V4719)
					reg7534 := PrimIsPair(reg7533)
					var reg7572 Obj
					if reg7534 == True {
						reg7535 := PrimTail(V4719)
						reg7536 := PrimTail(reg7535)
						reg7537 := PrimIsPair(reg7536)
						var reg7567 Obj
						if reg7537 == True {
							reg7538 := MakeSymbol("where")
							reg7539 := PrimTail(V4719)
							reg7540 := PrimTail(reg7539)
							reg7541 := PrimHead(reg7540)
							reg7542 := PrimEqual(reg7538, reg7541)
							var reg7562 Obj
							if reg7542 == True {
								reg7543 := PrimTail(V4719)
								reg7544 := PrimTail(reg7543)
								reg7545 := PrimTail(reg7544)
								reg7546 := PrimIsPair(reg7545)
								var reg7557 Obj
								if reg7546 == True {
									reg7547 := Nil
									reg7548 := PrimTail(V4719)
									reg7549 := PrimTail(reg7548)
									reg7550 := PrimTail(reg7549)
									reg7551 := PrimTail(reg7550)
									reg7552 := PrimEqual(reg7547, reg7551)
									var reg7555 Obj
									if reg7552 == True {
										reg7553 := True
										reg7555 = reg7553
									} else {
										reg7554 := False
										reg7555 = reg7554
									}
									reg7557 = reg7555
								} else {
									reg7556 := False
									reg7557 = reg7556
								}
								var reg7560 Obj
								if reg7557 == True {
									reg7558 := True
									reg7560 = reg7558
								} else {
									reg7559 := False
									reg7560 = reg7559
								}
								reg7562 = reg7560
							} else {
								reg7561 := False
								reg7562 = reg7561
							}
							var reg7565 Obj
							if reg7562 == True {
								reg7563 := True
								reg7565 = reg7563
							} else {
								reg7564 := False
								reg7565 = reg7564
							}
							reg7567 = reg7565
						} else {
							reg7566 := False
							reg7567 = reg7566
						}
						var reg7570 Obj
						if reg7567 == True {
							reg7568 := True
							reg7570 = reg7568
						} else {
							reg7569 := False
							reg7570 = reg7569
						}
						reg7572 = reg7570
					} else {
						reg7571 := False
						reg7572 = reg7571
					}
					var reg7575 Obj
					if reg7572 == True {
						reg7573 := True
						reg7575 = reg7573
					} else {
						reg7574 := False
						reg7575 = reg7574
					}
					reg7577 = reg7575
				} else {
					reg7576 := False
					reg7577 = reg7576
				}
				var reg7580 Obj
				if reg7577 == True {
					reg7578 := True
					reg7580 = reg7578
				} else {
					reg7579 := False
					reg7580 = reg7579
				}
				reg7582 = reg7580
			} else {
				reg7581 := False
				reg7582 = reg7581
			}
			if reg7582 == True {
				reg7583 := __e.Call(__defun__reverse, V4720)
				reg7584 := MakeSymbol("where")
				reg7585 := PrimTail(V4719)
				reg7586 := PrimTail(reg7585)
				reg7587 := PrimTail(reg7586)
				reg7588 := PrimHead(reg7587)
				reg7589 := PrimTail(V4719)
				reg7590 := PrimHead(reg7589)
				reg7591 := Nil
				reg7592 := PrimCons(reg7590, reg7591)
				reg7593 := PrimCons(reg7588, reg7592)
				reg7594 := PrimCons(reg7584, reg7593)
				reg7595 := Nil
				reg7596 := PrimCons(reg7594, reg7595)
				reg7597 := PrimCons(reg7583, reg7596)
				__ctx.Return(reg7597)
				return
			} else {
				reg7598 := Nil
				reg7599 := PrimEqual(reg7598, V4719)
				if reg7599 == True {
					reg7600 := __e.Call(__defun__shen_4semantic_1completion_1warning, V4718, V4720)
					_ = reg7600
					reg7601 := MakeSymbol(":=")
					reg7602 := __e.Call(__defun__reverse, V4720)
					reg7603 := __e.Call(__defun__shen_4default__semantics, reg7602)
					reg7604 := Nil
					reg7605 := PrimCons(reg7603, reg7604)
					reg7606 := PrimCons(reg7601, reg7605)
					__ctx.TailApply(__defun__shen_4split__cc__rule, V4718, reg7606, V4720)
					return
				} else {
					reg7608 := PrimIsPair(V4719)
					if reg7608 == True {
						reg7609 := PrimTail(V4719)
						reg7610 := PrimHead(V4719)
						reg7611 := PrimCons(reg7610, V4720)
						__ctx.TailApply(__defun__shen_4split__cc__rule, V4718, reg7609, reg7611)
						return
					} else {
						reg7613 := MakeSymbol("shen.split_cc_rule")
						__ctx.TailApply(__defun__shen_4f__error, reg7613)
						return
					}
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.split_cc_rule", value: __defun__shen_4split__cc__rule})

	__defun__shen_4semantic_1completion_1warning = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4731 := __args[0]
		_ = V4731
		V4732 := __args[1]
		_ = V4732
		reg7615 := True
		reg7616 := PrimEqual(reg7615, V4731)
		if reg7616 == True {
			reg7617 := MakeString("warning: ")
			reg7618 := __e.Call(__defun__stoutput)
			reg7619 := __e.Call(__defun__shen_4prhush, reg7617, reg7618)
			_ = reg7619
			reg7620 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				X := __args[0]
				_ = X
				reg7621 := MakeString(" ")
				reg7622 := MakeSymbol("shen.a")
				reg7623 := __e.Call(__defun__shen_4app, X, reg7621, reg7622)
				reg7624 := __e.Call(__defun__stoutput)
				__ctx.TailApply(__defun__shen_4prhush, reg7623, reg7624)
				return
			}, 1)
			reg7626 := __e.Call(__defun__reverse, V4732)
			reg7627 := __e.Call(__defun__shen_4for_1each, reg7620, reg7626)
			_ = reg7627
			reg7628 := MakeString("has no semantics.\n")
			reg7629 := __e.Call(__defun__stoutput)
			__ctx.TailApply(__defun__shen_4prhush, reg7628, reg7629)
			return
		} else {
			reg7631 := MakeSymbol("shen.skip")
			__ctx.Return(reg7631)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.semantic-completion-warning", value: __defun__shen_4semantic_1completion_1warning})

	__defun__shen_4default__semantics = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4734 := __args[0]
		_ = V4734
		reg7632 := Nil
		reg7633 := PrimEqual(reg7632, V4734)
		if reg7633 == True {
			reg7634 := Nil
			__ctx.Return(reg7634)
			return
		} else {
			reg7635 := PrimIsPair(V4734)
			var reg7650 Obj
			if reg7635 == True {
				reg7636 := Nil
				reg7637 := PrimTail(V4734)
				reg7638 := PrimEqual(reg7636, reg7637)
				var reg7645 Obj
				if reg7638 == True {
					reg7639 := PrimHead(V4734)
					reg7640 := __e.Call(__defun__shen_4grammar__symbol_2, reg7639)
					var reg7643 Obj
					if reg7640 == True {
						reg7641 := True
						reg7643 = reg7641
					} else {
						reg7642 := False
						reg7643 = reg7642
					}
					reg7645 = reg7643
				} else {
					reg7644 := False
					reg7645 = reg7644
				}
				var reg7648 Obj
				if reg7645 == True {
					reg7646 := True
					reg7648 = reg7646
				} else {
					reg7647 := False
					reg7648 = reg7647
				}
				reg7650 = reg7648
			} else {
				reg7649 := False
				reg7650 = reg7649
			}
			if reg7650 == True {
				reg7651 := PrimHead(V4734)
				__ctx.Return(reg7651)
				return
			} else {
				reg7652 := PrimIsPair(V4734)
				var reg7659 Obj
				if reg7652 == True {
					reg7653 := PrimHead(V4734)
					reg7654 := __e.Call(__defun__shen_4grammar__symbol_2, reg7653)
					var reg7657 Obj
					if reg7654 == True {
						reg7655 := True
						reg7657 = reg7655
					} else {
						reg7656 := False
						reg7657 = reg7656
					}
					reg7659 = reg7657
				} else {
					reg7658 := False
					reg7659 = reg7658
				}
				if reg7659 == True {
					reg7660 := MakeSymbol("append")
					reg7661 := PrimHead(V4734)
					reg7662 := PrimTail(V4734)
					reg7663 := __e.Call(__defun__shen_4default__semantics, reg7662)
					reg7664 := Nil
					reg7665 := PrimCons(reg7663, reg7664)
					reg7666 := PrimCons(reg7661, reg7665)
					reg7667 := PrimCons(reg7660, reg7666)
					__ctx.Return(reg7667)
					return
				} else {
					reg7668 := PrimIsPair(V4734)
					if reg7668 == True {
						reg7669 := MakeSymbol("cons")
						reg7670 := PrimHead(V4734)
						reg7671 := PrimTail(V4734)
						reg7672 := __e.Call(__defun__shen_4default__semantics, reg7671)
						reg7673 := Nil
						reg7674 := PrimCons(reg7672, reg7673)
						reg7675 := PrimCons(reg7670, reg7674)
						reg7676 := PrimCons(reg7669, reg7675)
						__ctx.Return(reg7676)
						return
					} else {
						reg7677 := MakeSymbol("shen.default_semantics")
						__ctx.TailApply(__defun__shen_4f__error, reg7677)
						return
					}
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.default_semantics", value: __defun__shen_4default__semantics})

	__defun__shen_4grammar__symbol_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4736 := __args[0]
		_ = V4736
		reg7679 := PrimIsSymbol(V4736)
		if reg7679 == True {
			reg7680 := __e.Call(__defun__explode, V4736)
			reg7681 := __e.Call(__defun__shen_4strip_1pathname, reg7680)
			Cs := reg7681
			_ = Cs
			reg7682 := PrimHead(Cs)
			reg7683 := MakeString("<")
			reg7684 := PrimEqual(reg7682, reg7683)
			var reg7693 Obj
			if reg7684 == True {
				reg7685 := __e.Call(__defun__reverse, Cs)
				reg7686 := PrimHead(reg7685)
				reg7687 := MakeString(">")
				reg7688 := PrimEqual(reg7686, reg7687)
				var reg7691 Obj
				if reg7688 == True {
					reg7689 := True
					reg7691 = reg7689
				} else {
					reg7690 := False
					reg7691 = reg7690
				}
				reg7693 = reg7691
			} else {
				reg7692 := False
				reg7693 = reg7692
			}
			if reg7693 == True {
				reg7694 := True
				__ctx.Return(reg7694)
				return
			} else {
				reg7695 := False
				__ctx.Return(reg7695)
				return
			}
		} else {
			reg7696 := False
			__ctx.Return(reg7696)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.grammar_symbol?", value: __defun__shen_4grammar__symbol_2})

	__defun__shen_4yacc__cases = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4738 := __args[0]
		_ = V4738
		reg7697 := PrimIsPair(V4738)
		var reg7705 Obj
		if reg7697 == True {
			reg7698 := Nil
			reg7699 := PrimTail(V4738)
			reg7700 := PrimEqual(reg7698, reg7699)
			var reg7703 Obj
			if reg7700 == True {
				reg7701 := True
				reg7703 = reg7701
			} else {
				reg7702 := False
				reg7703 = reg7702
			}
			reg7705 = reg7703
		} else {
			reg7704 := False
			reg7705 = reg7704
		}
		if reg7705 == True {
			reg7706 := PrimHead(V4738)
			__ctx.Return(reg7706)
			return
		} else {
			reg7707 := PrimIsPair(V4738)
			if reg7707 == True {
				reg7708 := MakeSymbol("YaccParse")
				P := reg7708
				_ = P
				reg7709 := MakeSymbol("let")
				reg7710 := PrimHead(V4738)
				reg7711 := MakeSymbol("if")
				reg7712 := MakeSymbol("=")
				reg7713 := MakeSymbol("fail")
				reg7714 := Nil
				reg7715 := PrimCons(reg7713, reg7714)
				reg7716 := Nil
				reg7717 := PrimCons(reg7715, reg7716)
				reg7718 := PrimCons(P, reg7717)
				reg7719 := PrimCons(reg7712, reg7718)
				reg7720 := PrimTail(V4738)
				reg7721 := __e.Call(__defun__shen_4yacc__cases, reg7720)
				reg7722 := Nil
				reg7723 := PrimCons(P, reg7722)
				reg7724 := PrimCons(reg7721, reg7723)
				reg7725 := PrimCons(reg7719, reg7724)
				reg7726 := PrimCons(reg7711, reg7725)
				reg7727 := Nil
				reg7728 := PrimCons(reg7726, reg7727)
				reg7729 := PrimCons(reg7710, reg7728)
				reg7730 := PrimCons(P, reg7729)
				reg7731 := PrimCons(reg7709, reg7730)
				__ctx.Return(reg7731)
				return
			} else {
				reg7732 := MakeSymbol("shen.yacc_cases")
				__ctx.TailApply(__defun__shen_4f__error, reg7732)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.yacc_cases", value: __defun__shen_4yacc__cases})

	__defun__shen_4cc__body = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4740 := __args[0]
		_ = V4740
		reg7734 := PrimIsPair(V4740)
		var reg7750 Obj
		if reg7734 == True {
			reg7735 := PrimTail(V4740)
			reg7736 := PrimIsPair(reg7735)
			var reg7745 Obj
			if reg7736 == True {
				reg7737 := Nil
				reg7738 := PrimTail(V4740)
				reg7739 := PrimTail(reg7738)
				reg7740 := PrimEqual(reg7737, reg7739)
				var reg7743 Obj
				if reg7740 == True {
					reg7741 := True
					reg7743 = reg7741
				} else {
					reg7742 := False
					reg7743 = reg7742
				}
				reg7745 = reg7743
			} else {
				reg7744 := False
				reg7745 = reg7744
			}
			var reg7748 Obj
			if reg7745 == True {
				reg7746 := True
				reg7748 = reg7746
			} else {
				reg7747 := False
				reg7748 = reg7747
			}
			reg7750 = reg7748
		} else {
			reg7749 := False
			reg7750 = reg7749
		}
		if reg7750 == True {
			reg7751 := PrimHead(V4740)
			reg7752 := MakeSymbol("Stream")
			reg7753 := PrimTail(V4740)
			reg7754 := PrimHead(reg7753)
			__ctx.TailApply(__defun__shen_4syntax, reg7751, reg7752, reg7754)
			return
		} else {
			reg7756 := MakeSymbol("shen.cc_body")
			__ctx.TailApply(__defun__shen_4f__error, reg7756)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.cc_body", value: __defun__shen_4cc__body})

	__defun__shen_4syntax = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4744 := __args[0]
		_ = V4744
		V4745 := __args[1]
		_ = V4745
		V4746 := __args[2]
		_ = V4746
		reg7758 := Nil
		reg7759 := PrimEqual(reg7758, V4744)
		var reg7798 Obj
		if reg7759 == True {
			reg7760 := PrimIsPair(V4746)
			var reg7793 Obj
			if reg7760 == True {
				reg7761 := MakeSymbol("where")
				reg7762 := PrimHead(V4746)
				reg7763 := PrimEqual(reg7761, reg7762)
				var reg7788 Obj
				if reg7763 == True {
					reg7764 := PrimTail(V4746)
					reg7765 := PrimIsPair(reg7764)
					var reg7783 Obj
					if reg7765 == True {
						reg7766 := PrimTail(V4746)
						reg7767 := PrimTail(reg7766)
						reg7768 := PrimIsPair(reg7767)
						var reg7778 Obj
						if reg7768 == True {
							reg7769 := Nil
							reg7770 := PrimTail(V4746)
							reg7771 := PrimTail(reg7770)
							reg7772 := PrimTail(reg7771)
							reg7773 := PrimEqual(reg7769, reg7772)
							var reg7776 Obj
							if reg7773 == True {
								reg7774 := True
								reg7776 = reg7774
							} else {
								reg7775 := False
								reg7776 = reg7775
							}
							reg7778 = reg7776
						} else {
							reg7777 := False
							reg7778 = reg7777
						}
						var reg7781 Obj
						if reg7778 == True {
							reg7779 := True
							reg7781 = reg7779
						} else {
							reg7780 := False
							reg7781 = reg7780
						}
						reg7783 = reg7781
					} else {
						reg7782 := False
						reg7783 = reg7782
					}
					var reg7786 Obj
					if reg7783 == True {
						reg7784 := True
						reg7786 = reg7784
					} else {
						reg7785 := False
						reg7786 = reg7785
					}
					reg7788 = reg7786
				} else {
					reg7787 := False
					reg7788 = reg7787
				}
				var reg7791 Obj
				if reg7788 == True {
					reg7789 := True
					reg7791 = reg7789
				} else {
					reg7790 := False
					reg7791 = reg7790
				}
				reg7793 = reg7791
			} else {
				reg7792 := False
				reg7793 = reg7792
			}
			var reg7796 Obj
			if reg7793 == True {
				reg7794 := True
				reg7796 = reg7794
			} else {
				reg7795 := False
				reg7796 = reg7795
			}
			reg7798 = reg7796
		} else {
			reg7797 := False
			reg7798 = reg7797
		}
		if reg7798 == True {
			reg7799 := MakeSymbol("if")
			reg7800 := PrimTail(V4746)
			reg7801 := PrimHead(reg7800)
			reg7802 := __e.Call(__defun__shen_4semantics, reg7801)
			reg7803 := MakeSymbol("shen.pair")
			reg7804 := MakeSymbol("hd")
			reg7805 := Nil
			reg7806 := PrimCons(V4745, reg7805)
			reg7807 := PrimCons(reg7804, reg7806)
			reg7808 := PrimTail(V4746)
			reg7809 := PrimTail(reg7808)
			reg7810 := PrimHead(reg7809)
			reg7811 := __e.Call(__defun__shen_4semantics, reg7810)
			reg7812 := Nil
			reg7813 := PrimCons(reg7811, reg7812)
			reg7814 := PrimCons(reg7807, reg7813)
			reg7815 := PrimCons(reg7803, reg7814)
			reg7816 := MakeSymbol("fail")
			reg7817 := Nil
			reg7818 := PrimCons(reg7816, reg7817)
			reg7819 := Nil
			reg7820 := PrimCons(reg7818, reg7819)
			reg7821 := PrimCons(reg7815, reg7820)
			reg7822 := PrimCons(reg7802, reg7821)
			reg7823 := PrimCons(reg7799, reg7822)
			__ctx.Return(reg7823)
			return
		} else {
			reg7824 := Nil
			reg7825 := PrimEqual(reg7824, V4744)
			if reg7825 == True {
				reg7826 := MakeSymbol("shen.pair")
				reg7827 := MakeSymbol("hd")
				reg7828 := Nil
				reg7829 := PrimCons(V4745, reg7828)
				reg7830 := PrimCons(reg7827, reg7829)
				reg7831 := __e.Call(__defun__shen_4semantics, V4746)
				reg7832 := Nil
				reg7833 := PrimCons(reg7831, reg7832)
				reg7834 := PrimCons(reg7830, reg7833)
				reg7835 := PrimCons(reg7826, reg7834)
				__ctx.Return(reg7835)
				return
			} else {
				reg7836 := PrimIsPair(V4744)
				if reg7836 == True {
					reg7837 := PrimHead(V4744)
					reg7838 := __e.Call(__defun__shen_4grammar__symbol_2, reg7837)
					if reg7838 == True {
						__ctx.TailApply(__defun__shen_4recursive__descent, V4744, V4745, V4746)
						return
					} else {
						reg7840 := PrimHead(V4744)
						reg7841 := PrimIsVariable(reg7840)
						if reg7841 == True {
							__ctx.TailApply(__defun__shen_4variable_1match, V4744, V4745, V4746)
							return
						} else {
							reg7843 := PrimHead(V4744)
							reg7844 := __e.Call(__defun__shen_4jump__stream_2, reg7843)
							if reg7844 == True {
								__ctx.TailApply(__defun__shen_4jump__stream, V4744, V4745, V4746)
								return
							} else {
								reg7846 := PrimHead(V4744)
								reg7847 := __e.Call(__defun__shen_4terminal_2, reg7846)
								if reg7847 == True {
									__ctx.TailApply(__defun__shen_4check__stream, V4744, V4745, V4746)
									return
								} else {
									reg7849 := PrimHead(V4744)
									reg7850 := PrimIsPair(reg7849)
									if reg7850 == True {
										reg7851 := PrimHead(V4744)
										reg7852 := __e.Call(__defun__shen_4decons, reg7851)
										reg7853 := PrimTail(V4744)
										__ctx.TailApply(__defun__shen_4list_1stream, reg7852, reg7853, V4745, V4746)
										return
									} else {
										reg7855 := PrimHead(V4744)
										reg7856 := MakeString(" is not legal syntax\n")
										reg7857 := MakeSymbol("shen.a")
										reg7858 := __e.Call(__defun__shen_4app, reg7855, reg7856, reg7857)
										reg7859 := PrimSimpleError(reg7858)
										__ctx.Return(reg7859)
										return
									}
								}
							}
						}
					}
				} else {
					reg7860 := MakeSymbol("shen.syntax")
					__ctx.TailApply(__defun__shen_4f__error, reg7860)
					return
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.syntax", value: __defun__shen_4syntax})

	__defun__shen_4list_1stream = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4751 := __args[0]
		_ = V4751
		V4752 := __args[1]
		_ = V4752
		V4753 := __args[2]
		_ = V4753
		V4754 := __args[3]
		_ = V4754
		reg7862 := MakeSymbol("and")
		reg7863 := MakeSymbol("cons?")
		reg7864 := MakeSymbol("hd")
		reg7865 := Nil
		reg7866 := PrimCons(V4753, reg7865)
		reg7867 := PrimCons(reg7864, reg7866)
		reg7868 := Nil
		reg7869 := PrimCons(reg7867, reg7868)
		reg7870 := PrimCons(reg7863, reg7869)
		reg7871 := MakeSymbol("cons?")
		reg7872 := MakeSymbol("shen.hdhd")
		reg7873 := Nil
		reg7874 := PrimCons(V4753, reg7873)
		reg7875 := PrimCons(reg7872, reg7874)
		reg7876 := Nil
		reg7877 := PrimCons(reg7875, reg7876)
		reg7878 := PrimCons(reg7871, reg7877)
		reg7879 := Nil
		reg7880 := PrimCons(reg7878, reg7879)
		reg7881 := PrimCons(reg7870, reg7880)
		reg7882 := PrimCons(reg7862, reg7881)
		Test := reg7882
		_ = Test
		reg7883 := MakeSymbol("shen.place")
		reg7884 := __e.Call(__defun__gensym, reg7883)
		Placeholder := reg7884
		_ = Placeholder
		reg7885 := MakeSymbol("shen.pair")
		reg7886 := MakeSymbol("shen.tlhd")
		reg7887 := Nil
		reg7888 := PrimCons(V4753, reg7887)
		reg7889 := PrimCons(reg7886, reg7888)
		reg7890 := MakeSymbol("shen.hdtl")
		reg7891 := Nil
		reg7892 := PrimCons(V4753, reg7891)
		reg7893 := PrimCons(reg7890, reg7892)
		reg7894 := Nil
		reg7895 := PrimCons(reg7893, reg7894)
		reg7896 := PrimCons(reg7889, reg7895)
		reg7897 := PrimCons(reg7885, reg7896)
		reg7898 := __e.Call(__defun__shen_4syntax, V4752, reg7897, V4754)
		RunOn := reg7898
		_ = RunOn
		reg7899 := MakeSymbol("shen.pair")
		reg7900 := MakeSymbol("shen.hdhd")
		reg7901 := Nil
		reg7902 := PrimCons(V4753, reg7901)
		reg7903 := PrimCons(reg7900, reg7902)
		reg7904 := MakeSymbol("shen.hdtl")
		reg7905 := Nil
		reg7906 := PrimCons(V4753, reg7905)
		reg7907 := PrimCons(reg7904, reg7906)
		reg7908 := Nil
		reg7909 := PrimCons(reg7907, reg7908)
		reg7910 := PrimCons(reg7903, reg7909)
		reg7911 := PrimCons(reg7899, reg7910)
		reg7912 := __e.Call(__defun__shen_4syntax, V4751, reg7911, Placeholder)
		reg7913 := __e.Call(__defun__shen_4insert_1runon, RunOn, Placeholder, reg7912)
		Action := reg7913
		_ = Action
		reg7914 := MakeSymbol("if")
		reg7915 := MakeSymbol("fail")
		reg7916 := Nil
		reg7917 := PrimCons(reg7915, reg7916)
		reg7918 := Nil
		reg7919 := PrimCons(reg7917, reg7918)
		reg7920 := PrimCons(Action, reg7919)
		reg7921 := PrimCons(Test, reg7920)
		reg7922 := PrimCons(reg7914, reg7921)
		__ctx.Return(reg7922)
		return
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.list-stream", value: __defun__shen_4list_1stream})

	__defun__shen_4decons = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4756 := __args[0]
		_ = V4756
		reg7923 := PrimIsPair(V4756)
		var reg7966 Obj
		if reg7923 == True {
			reg7924 := MakeSymbol("cons")
			reg7925 := PrimHead(V4756)
			reg7926 := PrimEqual(reg7924, reg7925)
			var reg7961 Obj
			if reg7926 == True {
				reg7927 := PrimTail(V4756)
				reg7928 := PrimIsPair(reg7927)
				var reg7956 Obj
				if reg7928 == True {
					reg7929 := PrimTail(V4756)
					reg7930 := PrimTail(reg7929)
					reg7931 := PrimIsPair(reg7930)
					var reg7951 Obj
					if reg7931 == True {
						reg7932 := Nil
						reg7933 := PrimTail(V4756)
						reg7934 := PrimTail(reg7933)
						reg7935 := PrimHead(reg7934)
						reg7936 := PrimEqual(reg7932, reg7935)
						var reg7946 Obj
						if reg7936 == True {
							reg7937 := Nil
							reg7938 := PrimTail(V4756)
							reg7939 := PrimTail(reg7938)
							reg7940 := PrimTail(reg7939)
							reg7941 := PrimEqual(reg7937, reg7940)
							var reg7944 Obj
							if reg7941 == True {
								reg7942 := True
								reg7944 = reg7942
							} else {
								reg7943 := False
								reg7944 = reg7943
							}
							reg7946 = reg7944
						} else {
							reg7945 := False
							reg7946 = reg7945
						}
						var reg7949 Obj
						if reg7946 == True {
							reg7947 := True
							reg7949 = reg7947
						} else {
							reg7948 := False
							reg7949 = reg7948
						}
						reg7951 = reg7949
					} else {
						reg7950 := False
						reg7951 = reg7950
					}
					var reg7954 Obj
					if reg7951 == True {
						reg7952 := True
						reg7954 = reg7952
					} else {
						reg7953 := False
						reg7954 = reg7953
					}
					reg7956 = reg7954
				} else {
					reg7955 := False
					reg7956 = reg7955
				}
				var reg7959 Obj
				if reg7956 == True {
					reg7957 := True
					reg7959 = reg7957
				} else {
					reg7958 := False
					reg7959 = reg7958
				}
				reg7961 = reg7959
			} else {
				reg7960 := False
				reg7961 = reg7960
			}
			var reg7964 Obj
			if reg7961 == True {
				reg7962 := True
				reg7964 = reg7962
			} else {
				reg7963 := False
				reg7964 = reg7963
			}
			reg7966 = reg7964
		} else {
			reg7965 := False
			reg7966 = reg7965
		}
		if reg7966 == True {
			reg7967 := PrimTail(V4756)
			reg7968 := PrimHead(reg7967)
			reg7969 := Nil
			reg7970 := PrimCons(reg7968, reg7969)
			__ctx.Return(reg7970)
			return
		} else {
			reg7971 := PrimIsPair(V4756)
			var reg8004 Obj
			if reg7971 == True {
				reg7972 := MakeSymbol("cons")
				reg7973 := PrimHead(V4756)
				reg7974 := PrimEqual(reg7972, reg7973)
				var reg7999 Obj
				if reg7974 == True {
					reg7975 := PrimTail(V4756)
					reg7976 := PrimIsPair(reg7975)
					var reg7994 Obj
					if reg7976 == True {
						reg7977 := PrimTail(V4756)
						reg7978 := PrimTail(reg7977)
						reg7979 := PrimIsPair(reg7978)
						var reg7989 Obj
						if reg7979 == True {
							reg7980 := Nil
							reg7981 := PrimTail(V4756)
							reg7982 := PrimTail(reg7981)
							reg7983 := PrimTail(reg7982)
							reg7984 := PrimEqual(reg7980, reg7983)
							var reg7987 Obj
							if reg7984 == True {
								reg7985 := True
								reg7987 = reg7985
							} else {
								reg7986 := False
								reg7987 = reg7986
							}
							reg7989 = reg7987
						} else {
							reg7988 := False
							reg7989 = reg7988
						}
						var reg7992 Obj
						if reg7989 == True {
							reg7990 := True
							reg7992 = reg7990
						} else {
							reg7991 := False
							reg7992 = reg7991
						}
						reg7994 = reg7992
					} else {
						reg7993 := False
						reg7994 = reg7993
					}
					var reg7997 Obj
					if reg7994 == True {
						reg7995 := True
						reg7997 = reg7995
					} else {
						reg7996 := False
						reg7997 = reg7996
					}
					reg7999 = reg7997
				} else {
					reg7998 := False
					reg7999 = reg7998
				}
				var reg8002 Obj
				if reg7999 == True {
					reg8000 := True
					reg8002 = reg8000
				} else {
					reg8001 := False
					reg8002 = reg8001
				}
				reg8004 = reg8002
			} else {
				reg8003 := False
				reg8004 = reg8003
			}
			if reg8004 == True {
				reg8005 := PrimTail(V4756)
				reg8006 := PrimHead(reg8005)
				reg8007 := PrimTail(V4756)
				reg8008 := PrimTail(reg8007)
				reg8009 := PrimHead(reg8008)
				reg8010 := __e.Call(__defun__shen_4decons, reg8009)
				reg8011 := PrimCons(reg8006, reg8010)
				__ctx.Return(reg8011)
				return
			} else {
				__ctx.Return(V4756)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.decons", value: __defun__shen_4decons})

	__defun__shen_4insert_1runon = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4771 := __args[0]
		_ = V4771
		V4772 := __args[1]
		_ = V4772
		V4773 := __args[2]
		_ = V4773
		reg8012 := PrimIsPair(V4773)
		var reg8054 Obj
		if reg8012 == True {
			reg8013 := MakeSymbol("shen.pair")
			reg8014 := PrimHead(V4773)
			reg8015 := PrimEqual(reg8013, reg8014)
			var reg8049 Obj
			if reg8015 == True {
				reg8016 := PrimTail(V4773)
				reg8017 := PrimIsPair(reg8016)
				var reg8044 Obj
				if reg8017 == True {
					reg8018 := PrimTail(V4773)
					reg8019 := PrimTail(reg8018)
					reg8020 := PrimIsPair(reg8019)
					var reg8039 Obj
					if reg8020 == True {
						reg8021 := Nil
						reg8022 := PrimTail(V4773)
						reg8023 := PrimTail(reg8022)
						reg8024 := PrimTail(reg8023)
						reg8025 := PrimEqual(reg8021, reg8024)
						var reg8034 Obj
						if reg8025 == True {
							reg8026 := PrimTail(V4773)
							reg8027 := PrimTail(reg8026)
							reg8028 := PrimHead(reg8027)
							reg8029 := PrimEqual(reg8028, V4772)
							var reg8032 Obj
							if reg8029 == True {
								reg8030 := True
								reg8032 = reg8030
							} else {
								reg8031 := False
								reg8032 = reg8031
							}
							reg8034 = reg8032
						} else {
							reg8033 := False
							reg8034 = reg8033
						}
						var reg8037 Obj
						if reg8034 == True {
							reg8035 := True
							reg8037 = reg8035
						} else {
							reg8036 := False
							reg8037 = reg8036
						}
						reg8039 = reg8037
					} else {
						reg8038 := False
						reg8039 = reg8038
					}
					var reg8042 Obj
					if reg8039 == True {
						reg8040 := True
						reg8042 = reg8040
					} else {
						reg8041 := False
						reg8042 = reg8041
					}
					reg8044 = reg8042
				} else {
					reg8043 := False
					reg8044 = reg8043
				}
				var reg8047 Obj
				if reg8044 == True {
					reg8045 := True
					reg8047 = reg8045
				} else {
					reg8046 := False
					reg8047 = reg8046
				}
				reg8049 = reg8047
			} else {
				reg8048 := False
				reg8049 = reg8048
			}
			var reg8052 Obj
			if reg8049 == True {
				reg8050 := True
				reg8052 = reg8050
			} else {
				reg8051 := False
				reg8052 = reg8051
			}
			reg8054 = reg8052
		} else {
			reg8053 := False
			reg8054 = reg8053
		}
		if reg8054 == True {
			__ctx.Return(V4771)
			return
		} else {
			reg8055 := PrimIsPair(V4773)
			if reg8055 == True {
				reg8056 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
					Z := __args[0]
					_ = Z
					__ctx.TailApply(__defun__shen_4insert_1runon, V4771, V4772, Z)
					return
				}, 1)
				__ctx.TailApply(__defun__map, reg8056, V4773)
				return
			} else {
				__ctx.Return(V4773)
				return
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.insert-runon", value: __defun__shen_4insert_1runon})

	__defun__shen_4strip_1pathname = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4779 := __args[0]
		_ = V4779
		reg8059 := MakeString(".")
		reg8060 := __e.Call(__defun__element_2, reg8059, V4779)
		reg8061 := PrimNot(reg8060)
		if reg8061 == True {
			__ctx.Return(V4779)
			return
		} else {
			reg8062 := PrimIsPair(V4779)
			if reg8062 == True {
				reg8063 := PrimTail(V4779)
				__ctx.TailApply(__defun__shen_4strip_1pathname, reg8063)
				return
			} else {
				reg8065 := MakeSymbol("shen.strip-pathname")
				__ctx.TailApply(__defun__shen_4f__error, reg8065)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.strip-pathname", value: __defun__shen_4strip_1pathname})

	__defun__shen_4recursive__descent = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4783 := __args[0]
		_ = V4783
		V4784 := __args[1]
		_ = V4784
		V4785 := __args[2]
		_ = V4785
		reg8067 := PrimIsPair(V4783)
		if reg8067 == True {
			reg8068 := PrimHead(V4783)
			reg8069 := Nil
			reg8070 := PrimCons(V4784, reg8069)
			reg8071 := PrimCons(reg8068, reg8070)
			Test := reg8071
			_ = Test
			reg8072 := PrimTail(V4783)
			reg8073 := MakeSymbol("Parse_")
			reg8074 := PrimHead(V4783)
			reg8075 := __e.Call(__defun__concat, reg8073, reg8074)
			reg8076 := __e.Call(__defun__shen_4syntax, reg8072, reg8075, V4785)
			Action := reg8076
			_ = Action
			reg8077 := MakeSymbol("fail")
			reg8078 := Nil
			reg8079 := PrimCons(reg8077, reg8078)
			Else := reg8079
			_ = Else
			reg8080 := MakeSymbol("let")
			reg8081 := MakeSymbol("Parse_")
			reg8082 := PrimHead(V4783)
			reg8083 := __e.Call(__defun__concat, reg8081, reg8082)
			reg8084 := MakeSymbol("if")
			reg8085 := MakeSymbol("not")
			reg8086 := MakeSymbol("=")
			reg8087 := MakeSymbol("fail")
			reg8088 := Nil
			reg8089 := PrimCons(reg8087, reg8088)
			reg8090 := MakeSymbol("Parse_")
			reg8091 := PrimHead(V4783)
			reg8092 := __e.Call(__defun__concat, reg8090, reg8091)
			reg8093 := Nil
			reg8094 := PrimCons(reg8092, reg8093)
			reg8095 := PrimCons(reg8089, reg8094)
			reg8096 := PrimCons(reg8086, reg8095)
			reg8097 := Nil
			reg8098 := PrimCons(reg8096, reg8097)
			reg8099 := PrimCons(reg8085, reg8098)
			reg8100 := Nil
			reg8101 := PrimCons(Else, reg8100)
			reg8102 := PrimCons(Action, reg8101)
			reg8103 := PrimCons(reg8099, reg8102)
			reg8104 := PrimCons(reg8084, reg8103)
			reg8105 := Nil
			reg8106 := PrimCons(reg8104, reg8105)
			reg8107 := PrimCons(Test, reg8106)
			reg8108 := PrimCons(reg8083, reg8107)
			reg8109 := PrimCons(reg8080, reg8108)
			__ctx.Return(reg8109)
			return
		} else {
			reg8110 := MakeSymbol("shen.recursive_descent")
			__ctx.TailApply(__defun__shen_4f__error, reg8110)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.recursive_descent", value: __defun__shen_4recursive__descent})

	__defun__shen_4variable_1match = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4789 := __args[0]
		_ = V4789
		V4790 := __args[1]
		_ = V4790
		V4791 := __args[2]
		_ = V4791
		reg8112 := PrimIsPair(V4789)
		if reg8112 == True {
			reg8113 := MakeSymbol("cons?")
			reg8114 := MakeSymbol("hd")
			reg8115 := Nil
			reg8116 := PrimCons(V4790, reg8115)
			reg8117 := PrimCons(reg8114, reg8116)
			reg8118 := Nil
			reg8119 := PrimCons(reg8117, reg8118)
			reg8120 := PrimCons(reg8113, reg8119)
			Test := reg8120
			_ = Test
			reg8121 := MakeSymbol("let")
			reg8122 := MakeSymbol("Parse_")
			reg8123 := PrimHead(V4789)
			reg8124 := __e.Call(__defun__concat, reg8122, reg8123)
			reg8125 := MakeSymbol("shen.hdhd")
			reg8126 := Nil
			reg8127 := PrimCons(V4790, reg8126)
			reg8128 := PrimCons(reg8125, reg8127)
			reg8129 := PrimTail(V4789)
			reg8130 := MakeSymbol("shen.pair")
			reg8131 := MakeSymbol("shen.tlhd")
			reg8132 := Nil
			reg8133 := PrimCons(V4790, reg8132)
			reg8134 := PrimCons(reg8131, reg8133)
			reg8135 := MakeSymbol("shen.hdtl")
			reg8136 := Nil
			reg8137 := PrimCons(V4790, reg8136)
			reg8138 := PrimCons(reg8135, reg8137)
			reg8139 := Nil
			reg8140 := PrimCons(reg8138, reg8139)
			reg8141 := PrimCons(reg8134, reg8140)
			reg8142 := PrimCons(reg8130, reg8141)
			reg8143 := __e.Call(__defun__shen_4syntax, reg8129, reg8142, V4791)
			reg8144 := Nil
			reg8145 := PrimCons(reg8143, reg8144)
			reg8146 := PrimCons(reg8128, reg8145)
			reg8147 := PrimCons(reg8124, reg8146)
			reg8148 := PrimCons(reg8121, reg8147)
			Action := reg8148
			_ = Action
			reg8149 := MakeSymbol("fail")
			reg8150 := Nil
			reg8151 := PrimCons(reg8149, reg8150)
			Else := reg8151
			_ = Else
			reg8152 := MakeSymbol("if")
			reg8153 := Nil
			reg8154 := PrimCons(Else, reg8153)
			reg8155 := PrimCons(Action, reg8154)
			reg8156 := PrimCons(Test, reg8155)
			reg8157 := PrimCons(reg8152, reg8156)
			__ctx.Return(reg8157)
			return
		} else {
			reg8158 := MakeSymbol("shen.variable-match")
			__ctx.TailApply(__defun__shen_4f__error, reg8158)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.variable-match", value: __defun__shen_4variable_1match})

	__defun__shen_4terminal_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4801 := __args[0]
		_ = V4801
		reg8160 := PrimIsPair(V4801)
		if reg8160 == True {
			reg8161 := False
			__ctx.Return(reg8161)
			return
		} else {
			reg8162 := PrimIsVariable(V4801)
			if reg8162 == True {
				reg8163 := False
				__ctx.Return(reg8163)
				return
			} else {
				reg8164 := True
				__ctx.Return(reg8164)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.terminal?", value: __defun__shen_4terminal_2})

	__defun__shen_4jump__stream_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4807 := __args[0]
		_ = V4807
		reg8165 := MakeSymbol("_")
		reg8166 := PrimEqual(V4807, reg8165)
		if reg8166 == True {
			reg8167 := True
			__ctx.Return(reg8167)
			return
		} else {
			reg8168 := False
			__ctx.Return(reg8168)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.jump_stream?", value: __defun__shen_4jump__stream_2})

	__defun__shen_4check__stream = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4811 := __args[0]
		_ = V4811
		V4812 := __args[1]
		_ = V4812
		V4813 := __args[2]
		_ = V4813
		reg8169 := PrimIsPair(V4811)
		if reg8169 == True {
			reg8170 := MakeSymbol("and")
			reg8171 := MakeSymbol("cons?")
			reg8172 := MakeSymbol("hd")
			reg8173 := Nil
			reg8174 := PrimCons(V4812, reg8173)
			reg8175 := PrimCons(reg8172, reg8174)
			reg8176 := Nil
			reg8177 := PrimCons(reg8175, reg8176)
			reg8178 := PrimCons(reg8171, reg8177)
			reg8179 := MakeSymbol("=")
			reg8180 := PrimHead(V4811)
			reg8181 := MakeSymbol("shen.hdhd")
			reg8182 := Nil
			reg8183 := PrimCons(V4812, reg8182)
			reg8184 := PrimCons(reg8181, reg8183)
			reg8185 := Nil
			reg8186 := PrimCons(reg8184, reg8185)
			reg8187 := PrimCons(reg8180, reg8186)
			reg8188 := PrimCons(reg8179, reg8187)
			reg8189 := Nil
			reg8190 := PrimCons(reg8188, reg8189)
			reg8191 := PrimCons(reg8178, reg8190)
			reg8192 := PrimCons(reg8170, reg8191)
			Test := reg8192
			_ = Test
			reg8193 := MakeSymbol("NewStream")
			reg8194 := __e.Call(__defun__gensym, reg8193)
			NewStr := reg8194
			_ = NewStr
			reg8195 := MakeSymbol("let")
			reg8196 := MakeSymbol("shen.pair")
			reg8197 := MakeSymbol("shen.tlhd")
			reg8198 := Nil
			reg8199 := PrimCons(V4812, reg8198)
			reg8200 := PrimCons(reg8197, reg8199)
			reg8201 := MakeSymbol("shen.hdtl")
			reg8202 := Nil
			reg8203 := PrimCons(V4812, reg8202)
			reg8204 := PrimCons(reg8201, reg8203)
			reg8205 := Nil
			reg8206 := PrimCons(reg8204, reg8205)
			reg8207 := PrimCons(reg8200, reg8206)
			reg8208 := PrimCons(reg8196, reg8207)
			reg8209 := PrimTail(V4811)
			reg8210 := __e.Call(__defun__shen_4syntax, reg8209, NewStr, V4813)
			reg8211 := Nil
			reg8212 := PrimCons(reg8210, reg8211)
			reg8213 := PrimCons(reg8208, reg8212)
			reg8214 := PrimCons(NewStr, reg8213)
			reg8215 := PrimCons(reg8195, reg8214)
			Action := reg8215
			_ = Action
			reg8216 := MakeSymbol("fail")
			reg8217 := Nil
			reg8218 := PrimCons(reg8216, reg8217)
			Else := reg8218
			_ = Else
			reg8219 := MakeSymbol("if")
			reg8220 := Nil
			reg8221 := PrimCons(Else, reg8220)
			reg8222 := PrimCons(Action, reg8221)
			reg8223 := PrimCons(Test, reg8222)
			reg8224 := PrimCons(reg8219, reg8223)
			__ctx.Return(reg8224)
			return
		} else {
			reg8225 := MakeSymbol("shen.check_stream")
			__ctx.TailApply(__defun__shen_4f__error, reg8225)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.check_stream", value: __defun__shen_4check__stream})

	__defun__shen_4jump__stream = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4817 := __args[0]
		_ = V4817
		V4818 := __args[1]
		_ = V4818
		V4819 := __args[2]
		_ = V4819
		reg8227 := PrimIsPair(V4817)
		if reg8227 == True {
			reg8228 := MakeSymbol("cons?")
			reg8229 := MakeSymbol("hd")
			reg8230 := Nil
			reg8231 := PrimCons(V4818, reg8230)
			reg8232 := PrimCons(reg8229, reg8231)
			reg8233 := Nil
			reg8234 := PrimCons(reg8232, reg8233)
			reg8235 := PrimCons(reg8228, reg8234)
			Test := reg8235
			_ = Test
			reg8236 := PrimTail(V4817)
			reg8237 := MakeSymbol("shen.pair")
			reg8238 := MakeSymbol("shen.tlhd")
			reg8239 := Nil
			reg8240 := PrimCons(V4818, reg8239)
			reg8241 := PrimCons(reg8238, reg8240)
			reg8242 := MakeSymbol("shen.hdtl")
			reg8243 := Nil
			reg8244 := PrimCons(V4818, reg8243)
			reg8245 := PrimCons(reg8242, reg8244)
			reg8246 := Nil
			reg8247 := PrimCons(reg8245, reg8246)
			reg8248 := PrimCons(reg8241, reg8247)
			reg8249 := PrimCons(reg8237, reg8248)
			reg8250 := __e.Call(__defun__shen_4syntax, reg8236, reg8249, V4819)
			Action := reg8250
			_ = Action
			reg8251 := MakeSymbol("fail")
			reg8252 := Nil
			reg8253 := PrimCons(reg8251, reg8252)
			Else := reg8253
			_ = Else
			reg8254 := MakeSymbol("if")
			reg8255 := Nil
			reg8256 := PrimCons(Else, reg8255)
			reg8257 := PrimCons(Action, reg8256)
			reg8258 := PrimCons(Test, reg8257)
			reg8259 := PrimCons(reg8254, reg8258)
			__ctx.Return(reg8259)
			return
		} else {
			reg8260 := MakeSymbol("shen.jump_stream")
			__ctx.TailApply(__defun__shen_4f__error, reg8260)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.jump_stream", value: __defun__shen_4jump__stream})

	__defun__shen_4semantics = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4821 := __args[0]
		_ = V4821
		reg8262 := Nil
		reg8263 := PrimEqual(reg8262, V4821)
		if reg8263 == True {
			reg8264 := Nil
			__ctx.Return(reg8264)
			return
		} else {
			reg8265 := __e.Call(__defun__shen_4grammar__symbol_2, V4821)
			if reg8265 == True {
				reg8266 := MakeSymbol("shen.hdtl")
				reg8267 := MakeSymbol("Parse_")
				reg8268 := __e.Call(__defun__concat, reg8267, V4821)
				reg8269 := Nil
				reg8270 := PrimCons(reg8268, reg8269)
				reg8271 := PrimCons(reg8266, reg8270)
				__ctx.Return(reg8271)
				return
			} else {
				reg8272 := PrimIsVariable(V4821)
				if reg8272 == True {
					reg8273 := MakeSymbol("Parse_")
					__ctx.TailApply(__defun__concat, reg8273, V4821)
					return
				} else {
					reg8275 := PrimIsPair(V4821)
					if reg8275 == True {
						reg8276 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
							Z := __args[0]
							_ = Z
							__ctx.TailApply(__defun__shen_4semantics, Z)
							return
						}, 1)
						__ctx.TailApply(__defun__map, reg8276, V4821)
						return
					} else {
						__ctx.Return(V4821)
						return
					}
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.semantics", value: __defun__shen_4semantics})

	__defun__shen_4pair = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4824 := __args[0]
		_ = V4824
		V4825 := __args[1]
		_ = V4825
		reg8279 := Nil
		reg8280 := PrimCons(V4825, reg8279)
		reg8281 := PrimCons(V4824, reg8280)
		__ctx.Return(reg8281)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.pair", value: __defun__shen_4pair})

	__defun__shen_4hdtl = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4827 := __args[0]
		_ = V4827
		reg8282 := PrimTail(V4827)
		reg8283 := PrimHead(reg8282)
		__ctx.Return(reg8283)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.hdtl", value: __defun__shen_4hdtl})

	__defun__shen_4hdhd = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4829 := __args[0]
		_ = V4829
		reg8284 := PrimHead(V4829)
		reg8285 := PrimHead(reg8284)
		__ctx.Return(reg8285)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.hdhd", value: __defun__shen_4hdhd})

	__defun__shen_4tlhd = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4831 := __args[0]
		_ = V4831
		reg8286 := PrimHead(V4831)
		reg8287 := PrimTail(reg8286)
		__ctx.Return(reg8287)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.tlhd", value: __defun__shen_4tlhd})

	__defun__shen_4snd_1or_1fail = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4839 := __args[0]
		_ = V4839
		reg8288 := PrimIsPair(V4839)
		var reg8304 Obj
		if reg8288 == True {
			reg8289 := PrimTail(V4839)
			reg8290 := PrimIsPair(reg8289)
			var reg8299 Obj
			if reg8290 == True {
				reg8291 := Nil
				reg8292 := PrimTail(V4839)
				reg8293 := PrimTail(reg8292)
				reg8294 := PrimEqual(reg8291, reg8293)
				var reg8297 Obj
				if reg8294 == True {
					reg8295 := True
					reg8297 = reg8295
				} else {
					reg8296 := False
					reg8297 = reg8296
				}
				reg8299 = reg8297
			} else {
				reg8298 := False
				reg8299 = reg8298
			}
			var reg8302 Obj
			if reg8299 == True {
				reg8300 := True
				reg8302 = reg8300
			} else {
				reg8301 := False
				reg8302 = reg8301
			}
			reg8304 = reg8302
		} else {
			reg8303 := False
			reg8304 = reg8303
		}
		if reg8304 == True {
			reg8305 := PrimTail(V4839)
			reg8306 := PrimHead(reg8305)
			__ctx.Return(reg8306)
			return
		} else {
			__ctx.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.snd-or-fail", value: __defun__shen_4snd_1or_1fail})

	__defun__fail = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg8308 := MakeSymbol("...")
		__ctx.Return(reg8308)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "fail", value: __defun__fail})

	__defun___5_b_6 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4847 := __args[0]
		_ = V4847
		reg8309 := PrimIsPair(V4847)
		var reg8325 Obj
		if reg8309 == True {
			reg8310 := PrimTail(V4847)
			reg8311 := PrimIsPair(reg8310)
			var reg8320 Obj
			if reg8311 == True {
				reg8312 := Nil
				reg8313 := PrimTail(V4847)
				reg8314 := PrimTail(reg8313)
				reg8315 := PrimEqual(reg8312, reg8314)
				var reg8318 Obj
				if reg8315 == True {
					reg8316 := True
					reg8318 = reg8316
				} else {
					reg8317 := False
					reg8318 = reg8317
				}
				reg8320 = reg8318
			} else {
				reg8319 := False
				reg8320 = reg8319
			}
			var reg8323 Obj
			if reg8320 == True {
				reg8321 := True
				reg8323 = reg8321
			} else {
				reg8322 := False
				reg8323 = reg8322
			}
			reg8325 = reg8323
		} else {
			reg8324 := False
			reg8325 = reg8324
		}
		if reg8325 == True {
			reg8326 := Nil
			reg8327 := PrimHead(V4847)
			reg8328 := Nil
			reg8329 := PrimCons(reg8327, reg8328)
			reg8330 := PrimCons(reg8326, reg8329)
			__ctx.Return(reg8330)
			return
		} else {
			__ctx.TailApply(__defun__fail)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "<!>", value: __defun___5_b_6})

	__defun___5e_6 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V4853 := __args[0]
		_ = V4853
		reg8332 := PrimIsPair(V4853)
		var reg8348 Obj
		if reg8332 == True {
			reg8333 := PrimTail(V4853)
			reg8334 := PrimIsPair(reg8333)
			var reg8343 Obj
			if reg8334 == True {
				reg8335 := Nil
				reg8336 := PrimTail(V4853)
				reg8337 := PrimTail(reg8336)
				reg8338 := PrimEqual(reg8335, reg8337)
				var reg8341 Obj
				if reg8338 == True {
					reg8339 := True
					reg8341 = reg8339
				} else {
					reg8340 := False
					reg8341 = reg8340
				}
				reg8343 = reg8341
			} else {
				reg8342 := False
				reg8343 = reg8342
			}
			var reg8346 Obj
			if reg8343 == True {
				reg8344 := True
				reg8346 = reg8344
			} else {
				reg8345 := False
				reg8346 = reg8345
			}
			reg8348 = reg8346
		} else {
			reg8347 := False
			reg8348 = reg8347
		}
		if reg8348 == True {
			reg8349 := PrimHead(V4853)
			reg8350 := Nil
			reg8351 := Nil
			reg8352 := PrimCons(reg8350, reg8351)
			reg8353 := PrimCons(reg8349, reg8352)
			__ctx.Return(reg8353)
			return
		} else {
			reg8354 := MakeSymbol("<e>")
			__ctx.TailApply(__defun__shen_4f__error, reg8354)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "<e>", value: __defun___5e_6})

}
