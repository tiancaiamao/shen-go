package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__macroexpand Obj                       // macroexpand
var __defun__shen_4error_1macro Obj                // shen.error-macro
var __defun__shen_4output_1macro Obj               // shen.output-macro
var __defun__shen_4make_1string_1macro Obj         // shen.make-string-macro
var __defun__shen_4input_1macro Obj                // shen.input-macro
var __defun__shen_4compose Obj                     // shen.compose
var __defun__shen_4compile_1macro Obj              // shen.compile-macro
var __defun__shen_4prolog_1macro Obj               // shen.prolog-macro
var __defun__shen_4externally_1bound Obj           // shen.externally-bound
var __defun__shen_4locally_1bind_1prolog_1vs Obj   // shen.locally-bind-prolog-vs
var __defun__shen_4bld_1prolog_1call Obj           // shen.bld-prolog-call
var __defun__shen_4defprolog_1macro Obj            // shen.defprolog-macro
var __defun__shen_4datatype_1macro Obj             // shen.datatype-macro
var __defun__shen_4intern_1type Obj                // shen.intern-type
var __defun__shen_4_8s_1macro Obj                  // shen.@s-macro
var __defun__shen_4synonyms_1macro Obj             // shen.synonyms-macro
var __defun__shen_4curry_1synonyms Obj             // shen.curry-synonyms
var __defun__shen_4nl_1macro Obj                   // shen.nl-macro
var __defun__shen_4assoc_1macro Obj                // shen.assoc-macro
var __defun__shen_4let_1macro Obj                  // shen.let-macro
var __defun__shen_4abs_1macro Obj                  // shen.abs-macro
var __defun__shen_4cases_1macro Obj                // shen.cases-macro
var __defun__shen_4timer_1macro Obj                // shen.timer-macro
var __defun__shen_4tuple_1up Obj                   // shen.tuple-up
var __defun__shen_4put_cget_1macro Obj             // shen.put/get-macro
var __defun__shen_4function_1macro Obj             // shen.function-macro
var __defun__shen_4function_1abstraction Obj       // shen.function-abstraction
var __defun__shen_4function_1abstraction_1help Obj // shen.function-abstraction-help
var __defun__undefmacro Obj                        // undefmacro
var __defun__shen_4findpos Obj                     // shen.findpos
var __defun__shen_4remove_1nth Obj                 // shen.remove-nth

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg17415 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__ctx.Return(reg17415)
		return
	}, 0))
	__defun__macroexpand = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V718 := __args[0]
		_ = V718
		reg17416 := MakeSymbol("*macros*")
		reg17417 := PrimValue(reg17416)
		reg17418 := __e.Call(__defun__shen_4compose, reg17417, V718)
		Y := reg17418
		_ = Y
		reg17419 := PrimEqual(V718, Y)
		if reg17419 == True {
			__ctx.Return(V718)
			return
		} else {
			reg17420 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				Z := __args[0]
				_ = Z
				__ctx.TailApply(__defun__macroexpand, Z)
				return
			}, 1)
			__ctx.TailApply(__defun__shen_4walk, reg17420, Y)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "macroexpand", value: __defun__macroexpand})

	__defun__shen_4error_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V720 := __args[0]
		_ = V720
		reg17423 := PrimIsPair(V720)
		var reg17438 Obj
		if reg17423 == True {
			reg17424 := MakeSymbol("error")
			reg17425 := PrimHead(V720)
			reg17426 := PrimEqual(reg17424, reg17425)
			var reg17433 Obj
			if reg17426 == True {
				reg17427 := PrimTail(V720)
				reg17428 := PrimIsPair(reg17427)
				var reg17431 Obj
				if reg17428 == True {
					reg17429 := True
					reg17431 = reg17429
				} else {
					reg17430 := False
					reg17431 = reg17430
				}
				reg17433 = reg17431
			} else {
				reg17432 := False
				reg17433 = reg17432
			}
			var reg17436 Obj
			if reg17433 == True {
				reg17434 := True
				reg17436 = reg17434
			} else {
				reg17435 := False
				reg17436 = reg17435
			}
			reg17438 = reg17436
		} else {
			reg17437 := False
			reg17438 = reg17437
		}
		if reg17438 == True {
			reg17439 := MakeSymbol("simple-error")
			reg17440 := PrimTail(V720)
			reg17441 := PrimHead(reg17440)
			reg17442 := PrimTail(V720)
			reg17443 := PrimTail(reg17442)
			reg17444 := __e.Call(__defun__shen_4mkstr, reg17441, reg17443)
			reg17445 := Nil
			reg17446 := PrimCons(reg17444, reg17445)
			reg17447 := PrimCons(reg17439, reg17446)
			__ctx.Return(reg17447)
			return
		} else {
			__ctx.Return(V720)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.error-macro", value: __defun__shen_4error_1macro})

	__defun__shen_4output_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V722 := __args[0]
		_ = V722
		reg17448 := PrimIsPair(V722)
		var reg17463 Obj
		if reg17448 == True {
			reg17449 := MakeSymbol("output")
			reg17450 := PrimHead(V722)
			reg17451 := PrimEqual(reg17449, reg17450)
			var reg17458 Obj
			if reg17451 == True {
				reg17452 := PrimTail(V722)
				reg17453 := PrimIsPair(reg17452)
				var reg17456 Obj
				if reg17453 == True {
					reg17454 := True
					reg17456 = reg17454
				} else {
					reg17455 := False
					reg17456 = reg17455
				}
				reg17458 = reg17456
			} else {
				reg17457 := False
				reg17458 = reg17457
			}
			var reg17461 Obj
			if reg17458 == True {
				reg17459 := True
				reg17461 = reg17459
			} else {
				reg17460 := False
				reg17461 = reg17460
			}
			reg17463 = reg17461
		} else {
			reg17462 := False
			reg17463 = reg17462
		}
		if reg17463 == True {
			reg17464 := MakeSymbol("shen.prhush")
			reg17465 := PrimTail(V722)
			reg17466 := PrimHead(reg17465)
			reg17467 := PrimTail(V722)
			reg17468 := PrimTail(reg17467)
			reg17469 := __e.Call(__defun__shen_4mkstr, reg17466, reg17468)
			reg17470 := MakeSymbol("stoutput")
			reg17471 := Nil
			reg17472 := PrimCons(reg17470, reg17471)
			reg17473 := Nil
			reg17474 := PrimCons(reg17472, reg17473)
			reg17475 := PrimCons(reg17469, reg17474)
			reg17476 := PrimCons(reg17464, reg17475)
			__ctx.Return(reg17476)
			return
		} else {
			reg17477 := PrimIsPair(V722)
			var reg17501 Obj
			if reg17477 == True {
				reg17478 := MakeSymbol("pr")
				reg17479 := PrimHead(V722)
				reg17480 := PrimEqual(reg17478, reg17479)
				var reg17496 Obj
				if reg17480 == True {
					reg17481 := PrimTail(V722)
					reg17482 := PrimIsPair(reg17481)
					var reg17491 Obj
					if reg17482 == True {
						reg17483 := Nil
						reg17484 := PrimTail(V722)
						reg17485 := PrimTail(reg17484)
						reg17486 := PrimEqual(reg17483, reg17485)
						var reg17489 Obj
						if reg17486 == True {
							reg17487 := True
							reg17489 = reg17487
						} else {
							reg17488 := False
							reg17489 = reg17488
						}
						reg17491 = reg17489
					} else {
						reg17490 := False
						reg17491 = reg17490
					}
					var reg17494 Obj
					if reg17491 == True {
						reg17492 := True
						reg17494 = reg17492
					} else {
						reg17493 := False
						reg17494 = reg17493
					}
					reg17496 = reg17494
				} else {
					reg17495 := False
					reg17496 = reg17495
				}
				var reg17499 Obj
				if reg17496 == True {
					reg17497 := True
					reg17499 = reg17497
				} else {
					reg17498 := False
					reg17499 = reg17498
				}
				reg17501 = reg17499
			} else {
				reg17500 := False
				reg17501 = reg17500
			}
			if reg17501 == True {
				reg17502 := MakeSymbol("pr")
				reg17503 := PrimTail(V722)
				reg17504 := PrimHead(reg17503)
				reg17505 := MakeSymbol("stoutput")
				reg17506 := Nil
				reg17507 := PrimCons(reg17505, reg17506)
				reg17508 := Nil
				reg17509 := PrimCons(reg17507, reg17508)
				reg17510 := PrimCons(reg17504, reg17509)
				reg17511 := PrimCons(reg17502, reg17510)
				__ctx.Return(reg17511)
				return
			} else {
				__ctx.Return(V722)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.output-macro", value: __defun__shen_4output_1macro})

	__defun__shen_4make_1string_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V724 := __args[0]
		_ = V724
		reg17512 := PrimIsPair(V724)
		var reg17527 Obj
		if reg17512 == True {
			reg17513 := MakeSymbol("make-string")
			reg17514 := PrimHead(V724)
			reg17515 := PrimEqual(reg17513, reg17514)
			var reg17522 Obj
			if reg17515 == True {
				reg17516 := PrimTail(V724)
				reg17517 := PrimIsPair(reg17516)
				var reg17520 Obj
				if reg17517 == True {
					reg17518 := True
					reg17520 = reg17518
				} else {
					reg17519 := False
					reg17520 = reg17519
				}
				reg17522 = reg17520
			} else {
				reg17521 := False
				reg17522 = reg17521
			}
			var reg17525 Obj
			if reg17522 == True {
				reg17523 := True
				reg17525 = reg17523
			} else {
				reg17524 := False
				reg17525 = reg17524
			}
			reg17527 = reg17525
		} else {
			reg17526 := False
			reg17527 = reg17526
		}
		if reg17527 == True {
			reg17528 := PrimTail(V724)
			reg17529 := PrimHead(reg17528)
			reg17530 := PrimTail(V724)
			reg17531 := PrimTail(reg17530)
			__ctx.TailApply(__defun__shen_4mkstr, reg17529, reg17531)
			return
		} else {
			__ctx.Return(V724)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.make-string-macro", value: __defun__shen_4make_1string_1macro})

	__defun__shen_4input_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V726 := __args[0]
		_ = V726
		reg17533 := PrimIsPair(V726)
		var reg17549 Obj
		if reg17533 == True {
			reg17534 := MakeSymbol("lineread")
			reg17535 := PrimHead(V726)
			reg17536 := PrimEqual(reg17534, reg17535)
			var reg17544 Obj
			if reg17536 == True {
				reg17537 := Nil
				reg17538 := PrimTail(V726)
				reg17539 := PrimEqual(reg17537, reg17538)
				var reg17542 Obj
				if reg17539 == True {
					reg17540 := True
					reg17542 = reg17540
				} else {
					reg17541 := False
					reg17542 = reg17541
				}
				reg17544 = reg17542
			} else {
				reg17543 := False
				reg17544 = reg17543
			}
			var reg17547 Obj
			if reg17544 == True {
				reg17545 := True
				reg17547 = reg17545
			} else {
				reg17546 := False
				reg17547 = reg17546
			}
			reg17549 = reg17547
		} else {
			reg17548 := False
			reg17549 = reg17548
		}
		if reg17549 == True {
			reg17550 := MakeSymbol("lineread")
			reg17551 := MakeSymbol("stinput")
			reg17552 := Nil
			reg17553 := PrimCons(reg17551, reg17552)
			reg17554 := Nil
			reg17555 := PrimCons(reg17553, reg17554)
			reg17556 := PrimCons(reg17550, reg17555)
			__ctx.Return(reg17556)
			return
		} else {
			reg17557 := PrimIsPair(V726)
			var reg17573 Obj
			if reg17557 == True {
				reg17558 := MakeSymbol("input")
				reg17559 := PrimHead(V726)
				reg17560 := PrimEqual(reg17558, reg17559)
				var reg17568 Obj
				if reg17560 == True {
					reg17561 := Nil
					reg17562 := PrimTail(V726)
					reg17563 := PrimEqual(reg17561, reg17562)
					var reg17566 Obj
					if reg17563 == True {
						reg17564 := True
						reg17566 = reg17564
					} else {
						reg17565 := False
						reg17566 = reg17565
					}
					reg17568 = reg17566
				} else {
					reg17567 := False
					reg17568 = reg17567
				}
				var reg17571 Obj
				if reg17568 == True {
					reg17569 := True
					reg17571 = reg17569
				} else {
					reg17570 := False
					reg17571 = reg17570
				}
				reg17573 = reg17571
			} else {
				reg17572 := False
				reg17573 = reg17572
			}
			if reg17573 == True {
				reg17574 := MakeSymbol("input")
				reg17575 := MakeSymbol("stinput")
				reg17576 := Nil
				reg17577 := PrimCons(reg17575, reg17576)
				reg17578 := Nil
				reg17579 := PrimCons(reg17577, reg17578)
				reg17580 := PrimCons(reg17574, reg17579)
				__ctx.Return(reg17580)
				return
			} else {
				reg17581 := PrimIsPair(V726)
				var reg17597 Obj
				if reg17581 == True {
					reg17582 := MakeSymbol("read")
					reg17583 := PrimHead(V726)
					reg17584 := PrimEqual(reg17582, reg17583)
					var reg17592 Obj
					if reg17584 == True {
						reg17585 := Nil
						reg17586 := PrimTail(V726)
						reg17587 := PrimEqual(reg17585, reg17586)
						var reg17590 Obj
						if reg17587 == True {
							reg17588 := True
							reg17590 = reg17588
						} else {
							reg17589 := False
							reg17590 = reg17589
						}
						reg17592 = reg17590
					} else {
						reg17591 := False
						reg17592 = reg17591
					}
					var reg17595 Obj
					if reg17592 == True {
						reg17593 := True
						reg17595 = reg17593
					} else {
						reg17594 := False
						reg17595 = reg17594
					}
					reg17597 = reg17595
				} else {
					reg17596 := False
					reg17597 = reg17596
				}
				if reg17597 == True {
					reg17598 := MakeSymbol("read")
					reg17599 := MakeSymbol("stinput")
					reg17600 := Nil
					reg17601 := PrimCons(reg17599, reg17600)
					reg17602 := Nil
					reg17603 := PrimCons(reg17601, reg17602)
					reg17604 := PrimCons(reg17598, reg17603)
					__ctx.Return(reg17604)
					return
				} else {
					reg17605 := PrimIsPair(V726)
					var reg17629 Obj
					if reg17605 == True {
						reg17606 := MakeSymbol("input+")
						reg17607 := PrimHead(V726)
						reg17608 := PrimEqual(reg17606, reg17607)
						var reg17624 Obj
						if reg17608 == True {
							reg17609 := PrimTail(V726)
							reg17610 := PrimIsPair(reg17609)
							var reg17619 Obj
							if reg17610 == True {
								reg17611 := Nil
								reg17612 := PrimTail(V726)
								reg17613 := PrimTail(reg17612)
								reg17614 := PrimEqual(reg17611, reg17613)
								var reg17617 Obj
								if reg17614 == True {
									reg17615 := True
									reg17617 = reg17615
								} else {
									reg17616 := False
									reg17617 = reg17616
								}
								reg17619 = reg17617
							} else {
								reg17618 := False
								reg17619 = reg17618
							}
							var reg17622 Obj
							if reg17619 == True {
								reg17620 := True
								reg17622 = reg17620
							} else {
								reg17621 := False
								reg17622 = reg17621
							}
							reg17624 = reg17622
						} else {
							reg17623 := False
							reg17624 = reg17623
						}
						var reg17627 Obj
						if reg17624 == True {
							reg17625 := True
							reg17627 = reg17625
						} else {
							reg17626 := False
							reg17627 = reg17626
						}
						reg17629 = reg17627
					} else {
						reg17628 := False
						reg17629 = reg17628
					}
					if reg17629 == True {
						reg17630 := MakeSymbol("input+")
						reg17631 := PrimTail(V726)
						reg17632 := PrimHead(reg17631)
						reg17633 := MakeSymbol("stinput")
						reg17634 := Nil
						reg17635 := PrimCons(reg17633, reg17634)
						reg17636 := Nil
						reg17637 := PrimCons(reg17635, reg17636)
						reg17638 := PrimCons(reg17632, reg17637)
						reg17639 := PrimCons(reg17630, reg17638)
						__ctx.Return(reg17639)
						return
					} else {
						reg17640 := PrimIsPair(V726)
						var reg17656 Obj
						if reg17640 == True {
							reg17641 := MakeSymbol("read-byte")
							reg17642 := PrimHead(V726)
							reg17643 := PrimEqual(reg17641, reg17642)
							var reg17651 Obj
							if reg17643 == True {
								reg17644 := Nil
								reg17645 := PrimTail(V726)
								reg17646 := PrimEqual(reg17644, reg17645)
								var reg17649 Obj
								if reg17646 == True {
									reg17647 := True
									reg17649 = reg17647
								} else {
									reg17648 := False
									reg17649 = reg17648
								}
								reg17651 = reg17649
							} else {
								reg17650 := False
								reg17651 = reg17650
							}
							var reg17654 Obj
							if reg17651 == True {
								reg17652 := True
								reg17654 = reg17652
							} else {
								reg17653 := False
								reg17654 = reg17653
							}
							reg17656 = reg17654
						} else {
							reg17655 := False
							reg17656 = reg17655
						}
						if reg17656 == True {
							reg17657 := MakeSymbol("read-byte")
							reg17658 := MakeSymbol("stinput")
							reg17659 := Nil
							reg17660 := PrimCons(reg17658, reg17659)
							reg17661 := Nil
							reg17662 := PrimCons(reg17660, reg17661)
							reg17663 := PrimCons(reg17657, reg17662)
							__ctx.Return(reg17663)
							return
						} else {
							__ctx.Return(V726)
							return
						}
					}
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.input-macro", value: __defun__shen_4input_1macro})

	__defun__shen_4compose = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V729 := __args[0]
		_ = V729
		V730 := __args[1]
		_ = V730
		reg17665 := Nil
		reg17666 := PrimEqual(reg17665, V729)
		if reg17666 == True {
			__ctx.Return(V730)
			return
		} else {
			reg17667 := PrimIsPair(V729)
			if reg17667 == True {
				reg17668 := PrimTail(V729)
				reg17669 := PrimHead(V729)
				f17664 := reg17669
				_ = f17664
				reg17670 := __e.Call(f17664, V730)
				__ctx.TailApply(__defun__shen_4compose, reg17668, reg17670)
				return
			} else {
				reg17672 := MakeSymbol("shen.compose")
				__ctx.TailApply(__defun__shen_4f__error, reg17672)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.compose", value: __defun__shen_4compose})

	__defun__shen_4compile_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V732 := __args[0]
		_ = V732
		reg17674 := PrimIsPair(V732)
		var reg17707 Obj
		if reg17674 == True {
			reg17675 := MakeSymbol("compile")
			reg17676 := PrimHead(V732)
			reg17677 := PrimEqual(reg17675, reg17676)
			var reg17702 Obj
			if reg17677 == True {
				reg17678 := PrimTail(V732)
				reg17679 := PrimIsPair(reg17678)
				var reg17697 Obj
				if reg17679 == True {
					reg17680 := PrimTail(V732)
					reg17681 := PrimTail(reg17680)
					reg17682 := PrimIsPair(reg17681)
					var reg17692 Obj
					if reg17682 == True {
						reg17683 := Nil
						reg17684 := PrimTail(V732)
						reg17685 := PrimTail(reg17684)
						reg17686 := PrimTail(reg17685)
						reg17687 := PrimEqual(reg17683, reg17686)
						var reg17690 Obj
						if reg17687 == True {
							reg17688 := True
							reg17690 = reg17688
						} else {
							reg17689 := False
							reg17690 = reg17689
						}
						reg17692 = reg17690
					} else {
						reg17691 := False
						reg17692 = reg17691
					}
					var reg17695 Obj
					if reg17692 == True {
						reg17693 := True
						reg17695 = reg17693
					} else {
						reg17694 := False
						reg17695 = reg17694
					}
					reg17697 = reg17695
				} else {
					reg17696 := False
					reg17697 = reg17696
				}
				var reg17700 Obj
				if reg17697 == True {
					reg17698 := True
					reg17700 = reg17698
				} else {
					reg17699 := False
					reg17700 = reg17699
				}
				reg17702 = reg17700
			} else {
				reg17701 := False
				reg17702 = reg17701
			}
			var reg17705 Obj
			if reg17702 == True {
				reg17703 := True
				reg17705 = reg17703
			} else {
				reg17704 := False
				reg17705 = reg17704
			}
			reg17707 = reg17705
		} else {
			reg17706 := False
			reg17707 = reg17706
		}
		if reg17707 == True {
			reg17708 := MakeSymbol("compile")
			reg17709 := PrimTail(V732)
			reg17710 := PrimHead(reg17709)
			reg17711 := PrimTail(V732)
			reg17712 := PrimTail(reg17711)
			reg17713 := PrimHead(reg17712)
			reg17714 := MakeSymbol("lambda")
			reg17715 := MakeSymbol("E")
			reg17716 := MakeSymbol("if")
			reg17717 := MakeSymbol("cons?")
			reg17718 := MakeSymbol("E")
			reg17719 := Nil
			reg17720 := PrimCons(reg17718, reg17719)
			reg17721 := PrimCons(reg17717, reg17720)
			reg17722 := MakeSymbol("error")
			reg17723 := MakeString("parse error here: ~S~%")
			reg17724 := MakeSymbol("E")
			reg17725 := Nil
			reg17726 := PrimCons(reg17724, reg17725)
			reg17727 := PrimCons(reg17723, reg17726)
			reg17728 := PrimCons(reg17722, reg17727)
			reg17729 := MakeSymbol("error")
			reg17730 := MakeString("parse error~%")
			reg17731 := Nil
			reg17732 := PrimCons(reg17730, reg17731)
			reg17733 := PrimCons(reg17729, reg17732)
			reg17734 := Nil
			reg17735 := PrimCons(reg17733, reg17734)
			reg17736 := PrimCons(reg17728, reg17735)
			reg17737 := PrimCons(reg17721, reg17736)
			reg17738 := PrimCons(reg17716, reg17737)
			reg17739 := Nil
			reg17740 := PrimCons(reg17738, reg17739)
			reg17741 := PrimCons(reg17715, reg17740)
			reg17742 := PrimCons(reg17714, reg17741)
			reg17743 := Nil
			reg17744 := PrimCons(reg17742, reg17743)
			reg17745 := PrimCons(reg17713, reg17744)
			reg17746 := PrimCons(reg17710, reg17745)
			reg17747 := PrimCons(reg17708, reg17746)
			__ctx.Return(reg17747)
			return
		} else {
			__ctx.Return(V732)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.compile-macro", value: __defun__shen_4compile_1macro})

	__defun__shen_4prolog_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V734 := __args[0]
		_ = V734
		reg17748 := PrimIsPair(V734)
		var reg17756 Obj
		if reg17748 == True {
			reg17749 := MakeSymbol("prolog?")
			reg17750 := PrimHead(V734)
			reg17751 := PrimEqual(reg17749, reg17750)
			var reg17754 Obj
			if reg17751 == True {
				reg17752 := True
				reg17754 = reg17752
			} else {
				reg17753 := False
				reg17754 = reg17753
			}
			reg17756 = reg17754
		} else {
			reg17755 := False
			reg17756 = reg17755
		}
		if reg17756 == True {
			reg17757 := MakeSymbol("let")
			reg17758 := MakeSymbol("NPP")
			reg17759 := MakeSymbol("shen.start-new-prolog-process")
			reg17760 := Nil
			reg17761 := PrimCons(reg17759, reg17760)
			reg17762 := MakeSymbol("NPP")
			reg17763 := PrimTail(V734)
			reg17764 := __e.Call(__defun__shen_4bld_1prolog_1call, reg17762, reg17763)
			Calls := reg17764
			_ = Calls
			reg17765 := PrimTail(V734)
			reg17766 := __e.Call(__defun__shen_4extract__vars, reg17765)
			Vs := reg17766
			_ = Vs
			reg17767 := PrimTail(V734)
			reg17768 := __e.Call(__defun__shen_4externally_1bound, reg17767)
			External := reg17768
			_ = External
			reg17769 := __e.Call(__defun__difference, Vs, External)
			PrologVs := reg17769
			_ = PrologVs
			reg17770 := MakeSymbol("NPP")
			reg17771 := __e.Call(__defun__shen_4locally_1bind_1prolog_1vs, reg17770, PrologVs, Calls)
			reg17772 := Nil
			reg17773 := PrimCons(reg17771, reg17772)
			reg17774 := PrimCons(reg17761, reg17773)
			reg17775 := PrimCons(reg17758, reg17774)
			reg17776 := PrimCons(reg17757, reg17775)
			__ctx.Return(reg17776)
			return
		} else {
			__ctx.Return(V734)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.prolog-macro", value: __defun__shen_4prolog_1macro})

	__defun__shen_4externally_1bound = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V740 := __args[0]
		_ = V740
		reg17777 := PrimIsPair(V740)
		var reg17801 Obj
		if reg17777 == True {
			reg17778 := MakeSymbol("receive")
			reg17779 := PrimHead(V740)
			reg17780 := PrimEqual(reg17778, reg17779)
			var reg17796 Obj
			if reg17780 == True {
				reg17781 := PrimTail(V740)
				reg17782 := PrimIsPair(reg17781)
				var reg17791 Obj
				if reg17782 == True {
					reg17783 := Nil
					reg17784 := PrimTail(V740)
					reg17785 := PrimTail(reg17784)
					reg17786 := PrimEqual(reg17783, reg17785)
					var reg17789 Obj
					if reg17786 == True {
						reg17787 := True
						reg17789 = reg17787
					} else {
						reg17788 := False
						reg17789 = reg17788
					}
					reg17791 = reg17789
				} else {
					reg17790 := False
					reg17791 = reg17790
				}
				var reg17794 Obj
				if reg17791 == True {
					reg17792 := True
					reg17794 = reg17792
				} else {
					reg17793 := False
					reg17794 = reg17793
				}
				reg17796 = reg17794
			} else {
				reg17795 := False
				reg17796 = reg17795
			}
			var reg17799 Obj
			if reg17796 == True {
				reg17797 := True
				reg17799 = reg17797
			} else {
				reg17798 := False
				reg17799 = reg17798
			}
			reg17801 = reg17799
		} else {
			reg17800 := False
			reg17801 = reg17800
		}
		if reg17801 == True {
			reg17802 := PrimTail(V740)
			__ctx.Return(reg17802)
			return
		} else {
			reg17803 := PrimIsPair(V740)
			if reg17803 == True {
				reg17804 := PrimHead(V740)
				reg17805 := __e.Call(__defun__shen_4externally_1bound, reg17804)
				reg17806 := PrimTail(V740)
				reg17807 := __e.Call(__defun__shen_4externally_1bound, reg17806)
				__ctx.TailApply(__defun__union, reg17805, reg17807)
				return
			} else {
				reg17809 := Nil
				__ctx.Return(reg17809)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.externally-bound", value: __defun__shen_4externally_1bound})

	__defun__shen_4locally_1bind_1prolog_1vs = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V758 := __args[0]
		_ = V758
		V759 := __args[1]
		_ = V759
		V760 := __args[2]
		_ = V760
		reg17810 := Nil
		reg17811 := PrimEqual(reg17810, V759)
		if reg17811 == True {
			__ctx.Return(V760)
			return
		} else {
			reg17812 := PrimIsPair(V759)
			if reg17812 == True {
				reg17813 := MakeSymbol("let")
				reg17814 := PrimHead(V759)
				reg17815 := MakeSymbol("shen.newpv")
				reg17816 := Nil
				reg17817 := PrimCons(V758, reg17816)
				reg17818 := PrimCons(reg17815, reg17817)
				reg17819 := PrimTail(V759)
				reg17820 := __e.Call(__defun__shen_4locally_1bind_1prolog_1vs, V758, reg17819, V760)
				reg17821 := Nil
				reg17822 := PrimCons(reg17820, reg17821)
				reg17823 := PrimCons(reg17818, reg17822)
				reg17824 := PrimCons(reg17814, reg17823)
				reg17825 := PrimCons(reg17813, reg17824)
				__ctx.Return(reg17825)
				return
			} else {
				reg17826 := MakeString("implementation error inp locally-bind-prolog-vs")
				reg17827 := PrimSimpleError(reg17826)
				__ctx.Return(reg17827)
				return
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.locally-bind-prolog-vs", value: __defun__shen_4locally_1bind_1prolog_1vs})

	__defun__shen_4bld_1prolog_1call = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V773 := __args[0]
		_ = V773
		V774 := __args[1]
		_ = V774
		reg17828 := Nil
		reg17829 := PrimEqual(reg17828, V774)
		if reg17829 == True {
			reg17830 := True
			__ctx.Return(reg17830)
			return
		} else {
			reg17831 := PrimIsPair(V774)
			var reg17839 Obj
			if reg17831 == True {
				reg17832 := MakeSymbol("!")
				reg17833 := PrimHead(V774)
				reg17834 := PrimEqual(reg17832, reg17833)
				var reg17837 Obj
				if reg17834 == True {
					reg17835 := True
					reg17837 = reg17835
				} else {
					reg17836 := False
					reg17837 = reg17836
				}
				reg17839 = reg17837
			} else {
				reg17838 := False
				reg17839 = reg17838
			}
			if reg17839 == True {
				reg17840 := MakeSymbol("cut")
				reg17841 := False
				reg17842 := MakeSymbol("freeze")
				reg17843 := PrimTail(V774)
				reg17844 := __e.Call(__defun__shen_4bld_1prolog_1call, V773, reg17843)
				reg17845 := Nil
				reg17846 := PrimCons(reg17844, reg17845)
				reg17847 := PrimCons(reg17842, reg17846)
				reg17848 := Nil
				reg17849 := PrimCons(reg17847, reg17848)
				reg17850 := PrimCons(V773, reg17849)
				reg17851 := PrimCons(reg17841, reg17850)
				reg17852 := PrimCons(reg17840, reg17851)
				__ctx.Return(reg17852)
				return
			} else {
				reg17853 := PrimIsPair(V774)
				var reg17887 Obj
				if reg17853 == True {
					reg17854 := PrimHead(V774)
					reg17855 := PrimIsPair(reg17854)
					var reg17882 Obj
					if reg17855 == True {
						reg17856 := MakeSymbol("when")
						reg17857 := PrimHead(V774)
						reg17858 := PrimHead(reg17857)
						reg17859 := PrimEqual(reg17856, reg17858)
						var reg17877 Obj
						if reg17859 == True {
							reg17860 := PrimHead(V774)
							reg17861 := PrimTail(reg17860)
							reg17862 := PrimIsPair(reg17861)
							var reg17872 Obj
							if reg17862 == True {
								reg17863 := Nil
								reg17864 := PrimHead(V774)
								reg17865 := PrimTail(reg17864)
								reg17866 := PrimTail(reg17865)
								reg17867 := PrimEqual(reg17863, reg17866)
								var reg17870 Obj
								if reg17867 == True {
									reg17868 := True
									reg17870 = reg17868
								} else {
									reg17869 := False
									reg17870 = reg17869
								}
								reg17872 = reg17870
							} else {
								reg17871 := False
								reg17872 = reg17871
							}
							var reg17875 Obj
							if reg17872 == True {
								reg17873 := True
								reg17875 = reg17873
							} else {
								reg17874 := False
								reg17875 = reg17874
							}
							reg17877 = reg17875
						} else {
							reg17876 := False
							reg17877 = reg17876
						}
						var reg17880 Obj
						if reg17877 == True {
							reg17878 := True
							reg17880 = reg17878
						} else {
							reg17879 := False
							reg17880 = reg17879
						}
						reg17882 = reg17880
					} else {
						reg17881 := False
						reg17882 = reg17881
					}
					var reg17885 Obj
					if reg17882 == True {
						reg17883 := True
						reg17885 = reg17883
					} else {
						reg17884 := False
						reg17885 = reg17884
					}
					reg17887 = reg17885
				} else {
					reg17886 := False
					reg17887 = reg17886
				}
				if reg17887 == True {
					reg17888 := MakeSymbol("fwhen")
					reg17889 := PrimHead(V774)
					reg17890 := PrimTail(reg17889)
					reg17891 := PrimHead(reg17890)
					reg17892 := __e.Call(__defun__shen_4insert_1deref, reg17891, V773)
					reg17893 := MakeSymbol("freeze")
					reg17894 := PrimTail(V774)
					reg17895 := __e.Call(__defun__shen_4bld_1prolog_1call, V773, reg17894)
					reg17896 := Nil
					reg17897 := PrimCons(reg17895, reg17896)
					reg17898 := PrimCons(reg17893, reg17897)
					reg17899 := Nil
					reg17900 := PrimCons(reg17898, reg17899)
					reg17901 := PrimCons(V773, reg17900)
					reg17902 := PrimCons(reg17892, reg17901)
					reg17903 := PrimCons(reg17888, reg17902)
					__ctx.Return(reg17903)
					return
				} else {
					reg17904 := PrimIsPair(V774)
					var reg17948 Obj
					if reg17904 == True {
						reg17905 := PrimHead(V774)
						reg17906 := PrimIsPair(reg17905)
						var reg17943 Obj
						if reg17906 == True {
							reg17907 := MakeSymbol("is")
							reg17908 := PrimHead(V774)
							reg17909 := PrimHead(reg17908)
							reg17910 := PrimEqual(reg17907, reg17909)
							var reg17938 Obj
							if reg17910 == True {
								reg17911 := PrimHead(V774)
								reg17912 := PrimTail(reg17911)
								reg17913 := PrimIsPair(reg17912)
								var reg17933 Obj
								if reg17913 == True {
									reg17914 := PrimHead(V774)
									reg17915 := PrimTail(reg17914)
									reg17916 := PrimTail(reg17915)
									reg17917 := PrimIsPair(reg17916)
									var reg17928 Obj
									if reg17917 == True {
										reg17918 := Nil
										reg17919 := PrimHead(V774)
										reg17920 := PrimTail(reg17919)
										reg17921 := PrimTail(reg17920)
										reg17922 := PrimTail(reg17921)
										reg17923 := PrimEqual(reg17918, reg17922)
										var reg17926 Obj
										if reg17923 == True {
											reg17924 := True
											reg17926 = reg17924
										} else {
											reg17925 := False
											reg17926 = reg17925
										}
										reg17928 = reg17926
									} else {
										reg17927 := False
										reg17928 = reg17927
									}
									var reg17931 Obj
									if reg17928 == True {
										reg17929 := True
										reg17931 = reg17929
									} else {
										reg17930 := False
										reg17931 = reg17930
									}
									reg17933 = reg17931
								} else {
									reg17932 := False
									reg17933 = reg17932
								}
								var reg17936 Obj
								if reg17933 == True {
									reg17934 := True
									reg17936 = reg17934
								} else {
									reg17935 := False
									reg17936 = reg17935
								}
								reg17938 = reg17936
							} else {
								reg17937 := False
								reg17938 = reg17937
							}
							var reg17941 Obj
							if reg17938 == True {
								reg17939 := True
								reg17941 = reg17939
							} else {
								reg17940 := False
								reg17941 = reg17940
							}
							reg17943 = reg17941
						} else {
							reg17942 := False
							reg17943 = reg17942
						}
						var reg17946 Obj
						if reg17943 == True {
							reg17944 := True
							reg17946 = reg17944
						} else {
							reg17945 := False
							reg17946 = reg17945
						}
						reg17948 = reg17946
					} else {
						reg17947 := False
						reg17948 = reg17947
					}
					if reg17948 == True {
						reg17949 := MakeSymbol("bind")
						reg17950 := PrimHead(V774)
						reg17951 := PrimTail(reg17950)
						reg17952 := PrimHead(reg17951)
						reg17953 := PrimHead(V774)
						reg17954 := PrimTail(reg17953)
						reg17955 := PrimTail(reg17954)
						reg17956 := PrimHead(reg17955)
						reg17957 := __e.Call(__defun__shen_4insert_1deref, reg17956, V773)
						reg17958 := MakeSymbol("freeze")
						reg17959 := PrimTail(V774)
						reg17960 := __e.Call(__defun__shen_4bld_1prolog_1call, V773, reg17959)
						reg17961 := Nil
						reg17962 := PrimCons(reg17960, reg17961)
						reg17963 := PrimCons(reg17958, reg17962)
						reg17964 := Nil
						reg17965 := PrimCons(reg17963, reg17964)
						reg17966 := PrimCons(V773, reg17965)
						reg17967 := PrimCons(reg17957, reg17966)
						reg17968 := PrimCons(reg17952, reg17967)
						reg17969 := PrimCons(reg17949, reg17968)
						__ctx.Return(reg17969)
						return
					} else {
						reg17970 := PrimIsPair(V774)
						var reg18004 Obj
						if reg17970 == True {
							reg17971 := PrimHead(V774)
							reg17972 := PrimIsPair(reg17971)
							var reg17999 Obj
							if reg17972 == True {
								reg17973 := MakeSymbol("receive")
								reg17974 := PrimHead(V774)
								reg17975 := PrimHead(reg17974)
								reg17976 := PrimEqual(reg17973, reg17975)
								var reg17994 Obj
								if reg17976 == True {
									reg17977 := PrimHead(V774)
									reg17978 := PrimTail(reg17977)
									reg17979 := PrimIsPair(reg17978)
									var reg17989 Obj
									if reg17979 == True {
										reg17980 := Nil
										reg17981 := PrimHead(V774)
										reg17982 := PrimTail(reg17981)
										reg17983 := PrimTail(reg17982)
										reg17984 := PrimEqual(reg17980, reg17983)
										var reg17987 Obj
										if reg17984 == True {
											reg17985 := True
											reg17987 = reg17985
										} else {
											reg17986 := False
											reg17987 = reg17986
										}
										reg17989 = reg17987
									} else {
										reg17988 := False
										reg17989 = reg17988
									}
									var reg17992 Obj
									if reg17989 == True {
										reg17990 := True
										reg17992 = reg17990
									} else {
										reg17991 := False
										reg17992 = reg17991
									}
									reg17994 = reg17992
								} else {
									reg17993 := False
									reg17994 = reg17993
								}
								var reg17997 Obj
								if reg17994 == True {
									reg17995 := True
									reg17997 = reg17995
								} else {
									reg17996 := False
									reg17997 = reg17996
								}
								reg17999 = reg17997
							} else {
								reg17998 := False
								reg17999 = reg17998
							}
							var reg18002 Obj
							if reg17999 == True {
								reg18000 := True
								reg18002 = reg18000
							} else {
								reg18001 := False
								reg18002 = reg18001
							}
							reg18004 = reg18002
						} else {
							reg18003 := False
							reg18004 = reg18003
						}
						if reg18004 == True {
							reg18005 := PrimTail(V774)
							__ctx.TailApply(__defun__shen_4bld_1prolog_1call, V773, reg18005)
							return
						} else {
							reg18007 := PrimIsPair(V774)
							var reg18051 Obj
							if reg18007 == True {
								reg18008 := PrimHead(V774)
								reg18009 := PrimIsPair(reg18008)
								var reg18046 Obj
								if reg18009 == True {
									reg18010 := MakeSymbol("bind")
									reg18011 := PrimHead(V774)
									reg18012 := PrimHead(reg18011)
									reg18013 := PrimEqual(reg18010, reg18012)
									var reg18041 Obj
									if reg18013 == True {
										reg18014 := PrimHead(V774)
										reg18015 := PrimTail(reg18014)
										reg18016 := PrimIsPair(reg18015)
										var reg18036 Obj
										if reg18016 == True {
											reg18017 := PrimHead(V774)
											reg18018 := PrimTail(reg18017)
											reg18019 := PrimTail(reg18018)
											reg18020 := PrimIsPair(reg18019)
											var reg18031 Obj
											if reg18020 == True {
												reg18021 := Nil
												reg18022 := PrimHead(V774)
												reg18023 := PrimTail(reg18022)
												reg18024 := PrimTail(reg18023)
												reg18025 := PrimTail(reg18024)
												reg18026 := PrimEqual(reg18021, reg18025)
												var reg18029 Obj
												if reg18026 == True {
													reg18027 := True
													reg18029 = reg18027
												} else {
													reg18028 := False
													reg18029 = reg18028
												}
												reg18031 = reg18029
											} else {
												reg18030 := False
												reg18031 = reg18030
											}
											var reg18034 Obj
											if reg18031 == True {
												reg18032 := True
												reg18034 = reg18032
											} else {
												reg18033 := False
												reg18034 = reg18033
											}
											reg18036 = reg18034
										} else {
											reg18035 := False
											reg18036 = reg18035
										}
										var reg18039 Obj
										if reg18036 == True {
											reg18037 := True
											reg18039 = reg18037
										} else {
											reg18038 := False
											reg18039 = reg18038
										}
										reg18041 = reg18039
									} else {
										reg18040 := False
										reg18041 = reg18040
									}
									var reg18044 Obj
									if reg18041 == True {
										reg18042 := True
										reg18044 = reg18042
									} else {
										reg18043 := False
										reg18044 = reg18043
									}
									reg18046 = reg18044
								} else {
									reg18045 := False
									reg18046 = reg18045
								}
								var reg18049 Obj
								if reg18046 == True {
									reg18047 := True
									reg18049 = reg18047
								} else {
									reg18048 := False
									reg18049 = reg18048
								}
								reg18051 = reg18049
							} else {
								reg18050 := False
								reg18051 = reg18050
							}
							if reg18051 == True {
								reg18052 := MakeSymbol("bind")
								reg18053 := PrimHead(V774)
								reg18054 := PrimTail(reg18053)
								reg18055 := PrimHead(reg18054)
								reg18056 := PrimHead(V774)
								reg18057 := PrimTail(reg18056)
								reg18058 := PrimTail(reg18057)
								reg18059 := PrimHead(reg18058)
								reg18060 := __e.Call(__defun__shen_4insert_1lazyderef, reg18059, V773)
								reg18061 := MakeSymbol("freeze")
								reg18062 := PrimTail(V774)
								reg18063 := __e.Call(__defun__shen_4bld_1prolog_1call, V773, reg18062)
								reg18064 := Nil
								reg18065 := PrimCons(reg18063, reg18064)
								reg18066 := PrimCons(reg18061, reg18065)
								reg18067 := Nil
								reg18068 := PrimCons(reg18066, reg18067)
								reg18069 := PrimCons(V773, reg18068)
								reg18070 := PrimCons(reg18060, reg18069)
								reg18071 := PrimCons(reg18055, reg18070)
								reg18072 := PrimCons(reg18052, reg18071)
								__ctx.Return(reg18072)
								return
							} else {
								reg18073 := PrimIsPair(V774)
								var reg18107 Obj
								if reg18073 == True {
									reg18074 := PrimHead(V774)
									reg18075 := PrimIsPair(reg18074)
									var reg18102 Obj
									if reg18075 == True {
										reg18076 := MakeSymbol("fwhen")
										reg18077 := PrimHead(V774)
										reg18078 := PrimHead(reg18077)
										reg18079 := PrimEqual(reg18076, reg18078)
										var reg18097 Obj
										if reg18079 == True {
											reg18080 := PrimHead(V774)
											reg18081 := PrimTail(reg18080)
											reg18082 := PrimIsPair(reg18081)
											var reg18092 Obj
											if reg18082 == True {
												reg18083 := Nil
												reg18084 := PrimHead(V774)
												reg18085 := PrimTail(reg18084)
												reg18086 := PrimTail(reg18085)
												reg18087 := PrimEqual(reg18083, reg18086)
												var reg18090 Obj
												if reg18087 == True {
													reg18088 := True
													reg18090 = reg18088
												} else {
													reg18089 := False
													reg18090 = reg18089
												}
												reg18092 = reg18090
											} else {
												reg18091 := False
												reg18092 = reg18091
											}
											var reg18095 Obj
											if reg18092 == True {
												reg18093 := True
												reg18095 = reg18093
											} else {
												reg18094 := False
												reg18095 = reg18094
											}
											reg18097 = reg18095
										} else {
											reg18096 := False
											reg18097 = reg18096
										}
										var reg18100 Obj
										if reg18097 == True {
											reg18098 := True
											reg18100 = reg18098
										} else {
											reg18099 := False
											reg18100 = reg18099
										}
										reg18102 = reg18100
									} else {
										reg18101 := False
										reg18102 = reg18101
									}
									var reg18105 Obj
									if reg18102 == True {
										reg18103 := True
										reg18105 = reg18103
									} else {
										reg18104 := False
										reg18105 = reg18104
									}
									reg18107 = reg18105
								} else {
									reg18106 := False
									reg18107 = reg18106
								}
								if reg18107 == True {
									reg18108 := MakeSymbol("fwhen")
									reg18109 := PrimHead(V774)
									reg18110 := PrimTail(reg18109)
									reg18111 := PrimHead(reg18110)
									reg18112 := __e.Call(__defun__shen_4insert_1lazyderef, reg18111, V773)
									reg18113 := MakeSymbol("freeze")
									reg18114 := PrimTail(V774)
									reg18115 := __e.Call(__defun__shen_4bld_1prolog_1call, V773, reg18114)
									reg18116 := Nil
									reg18117 := PrimCons(reg18115, reg18116)
									reg18118 := PrimCons(reg18113, reg18117)
									reg18119 := Nil
									reg18120 := PrimCons(reg18118, reg18119)
									reg18121 := PrimCons(V773, reg18120)
									reg18122 := PrimCons(reg18112, reg18121)
									reg18123 := PrimCons(reg18108, reg18122)
									__ctx.Return(reg18123)
									return
								} else {
									reg18124 := PrimIsPair(V774)
									if reg18124 == True {
										reg18125 := PrimHead(V774)
										reg18126 := MakeSymbol("freeze")
										reg18127 := PrimTail(V774)
										reg18128 := __e.Call(__defun__shen_4bld_1prolog_1call, V773, reg18127)
										reg18129 := Nil
										reg18130 := PrimCons(reg18128, reg18129)
										reg18131 := PrimCons(reg18126, reg18130)
										reg18132 := Nil
										reg18133 := PrimCons(reg18131, reg18132)
										reg18134 := PrimCons(V773, reg18133)
										__ctx.TailApply(__defun__append, reg18125, reg18134)
										return
									} else {
										reg18136 := MakeString("implementation error in bld-prolog-call")
										reg18137 := PrimSimpleError(reg18136)
										__ctx.Return(reg18137)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.bld-prolog-call", value: __defun__shen_4bld_1prolog_1call})

	__defun__shen_4defprolog_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V776 := __args[0]
		_ = V776
		reg18138 := PrimIsPair(V776)
		var reg18153 Obj
		if reg18138 == True {
			reg18139 := MakeSymbol("defprolog")
			reg18140 := PrimHead(V776)
			reg18141 := PrimEqual(reg18139, reg18140)
			var reg18148 Obj
			if reg18141 == True {
				reg18142 := PrimTail(V776)
				reg18143 := PrimIsPair(reg18142)
				var reg18146 Obj
				if reg18143 == True {
					reg18144 := True
					reg18146 = reg18144
				} else {
					reg18145 := False
					reg18146 = reg18145
				}
				reg18148 = reg18146
			} else {
				reg18147 := False
				reg18148 = reg18147
			}
			var reg18151 Obj
			if reg18148 == True {
				reg18149 := True
				reg18151 = reg18149
			} else {
				reg18150 := False
				reg18151 = reg18150
			}
			reg18153 = reg18151
		} else {
			reg18152 := False
			reg18153 = reg18152
		}
		if reg18153 == True {
			reg18154 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				Y := __args[0]
				_ = Y
				__ctx.TailApply(__defun__shen_4_5defprolog_6, Y)
				return
			}, 1)
			reg18156 := PrimTail(V776)
			reg18157 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				Y := __args[0]
				_ = Y
				reg18158 := PrimTail(V776)
				reg18159 := PrimHead(reg18158)
				__ctx.TailApply(__defun__shen_4prolog_1error, reg18159, Y)
				return
			}, 1)
			__ctx.TailApply(__defun__compile, reg18154, reg18156, reg18157)
			return
		} else {
			__ctx.Return(V776)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.defprolog-macro", value: __defun__shen_4defprolog_1macro})

	__defun__shen_4datatype_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V778 := __args[0]
		_ = V778
		reg18162 := PrimIsPair(V778)
		var reg18177 Obj
		if reg18162 == True {
			reg18163 := MakeSymbol("datatype")
			reg18164 := PrimHead(V778)
			reg18165 := PrimEqual(reg18163, reg18164)
			var reg18172 Obj
			if reg18165 == True {
				reg18166 := PrimTail(V778)
				reg18167 := PrimIsPair(reg18166)
				var reg18170 Obj
				if reg18167 == True {
					reg18168 := True
					reg18170 = reg18168
				} else {
					reg18169 := False
					reg18170 = reg18169
				}
				reg18172 = reg18170
			} else {
				reg18171 := False
				reg18172 = reg18171
			}
			var reg18175 Obj
			if reg18172 == True {
				reg18173 := True
				reg18175 = reg18173
			} else {
				reg18174 := False
				reg18175 = reg18174
			}
			reg18177 = reg18175
		} else {
			reg18176 := False
			reg18177 = reg18176
		}
		if reg18177 == True {
			reg18178 := MakeSymbol("shen.process-datatype")
			reg18179 := PrimTail(V778)
			reg18180 := PrimHead(reg18179)
			reg18181 := __e.Call(__defun__shen_4intern_1type, reg18180)
			reg18182 := MakeSymbol("compile")
			reg18183 := MakeSymbol("lambda")
			reg18184 := MakeSymbol("X")
			reg18185 := MakeSymbol("shen.<datatype-rules>")
			reg18186 := MakeSymbol("X")
			reg18187 := Nil
			reg18188 := PrimCons(reg18186, reg18187)
			reg18189 := PrimCons(reg18185, reg18188)
			reg18190 := Nil
			reg18191 := PrimCons(reg18189, reg18190)
			reg18192 := PrimCons(reg18184, reg18191)
			reg18193 := PrimCons(reg18183, reg18192)
			reg18194 := PrimTail(V778)
			reg18195 := PrimTail(reg18194)
			reg18196 := __e.Call(__defun__shen_4rcons__form, reg18195)
			reg18197 := MakeSymbol("function")
			reg18198 := MakeSymbol("shen.datatype-error")
			reg18199 := Nil
			reg18200 := PrimCons(reg18198, reg18199)
			reg18201 := PrimCons(reg18197, reg18200)
			reg18202 := Nil
			reg18203 := PrimCons(reg18201, reg18202)
			reg18204 := PrimCons(reg18196, reg18203)
			reg18205 := PrimCons(reg18193, reg18204)
			reg18206 := PrimCons(reg18182, reg18205)
			reg18207 := Nil
			reg18208 := PrimCons(reg18206, reg18207)
			reg18209 := PrimCons(reg18181, reg18208)
			reg18210 := PrimCons(reg18178, reg18209)
			__ctx.Return(reg18210)
			return
		} else {
			__ctx.Return(V778)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.datatype-macro", value: __defun__shen_4datatype_1macro})

	__defun__shen_4intern_1type = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V780 := __args[0]
		_ = V780
		reg18211 := PrimStr(V780)
		reg18212 := MakeString("#type")
		reg18213 := PrimStringConcat(reg18211, reg18212)
		reg18214 := PrimIntern(reg18213)
		__ctx.Return(reg18214)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.intern-type", value: __defun__shen_4intern_1type})

	__defun__shen_4_8s_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V782 := __args[0]
		_ = V782
		reg18215 := PrimIsPair(V782)
		var reg18247 Obj
		if reg18215 == True {
			reg18216 := MakeSymbol("@s")
			reg18217 := PrimHead(V782)
			reg18218 := PrimEqual(reg18216, reg18217)
			var reg18242 Obj
			if reg18218 == True {
				reg18219 := PrimTail(V782)
				reg18220 := PrimIsPair(reg18219)
				var reg18237 Obj
				if reg18220 == True {
					reg18221 := PrimTail(V782)
					reg18222 := PrimTail(reg18221)
					reg18223 := PrimIsPair(reg18222)
					var reg18232 Obj
					if reg18223 == True {
						reg18224 := PrimTail(V782)
						reg18225 := PrimTail(reg18224)
						reg18226 := PrimTail(reg18225)
						reg18227 := PrimIsPair(reg18226)
						var reg18230 Obj
						if reg18227 == True {
							reg18228 := True
							reg18230 = reg18228
						} else {
							reg18229 := False
							reg18230 = reg18229
						}
						reg18232 = reg18230
					} else {
						reg18231 := False
						reg18232 = reg18231
					}
					var reg18235 Obj
					if reg18232 == True {
						reg18233 := True
						reg18235 = reg18233
					} else {
						reg18234 := False
						reg18235 = reg18234
					}
					reg18237 = reg18235
				} else {
					reg18236 := False
					reg18237 = reg18236
				}
				var reg18240 Obj
				if reg18237 == True {
					reg18238 := True
					reg18240 = reg18238
				} else {
					reg18239 := False
					reg18240 = reg18239
				}
				reg18242 = reg18240
			} else {
				reg18241 := False
				reg18242 = reg18241
			}
			var reg18245 Obj
			if reg18242 == True {
				reg18243 := True
				reg18245 = reg18243
			} else {
				reg18244 := False
				reg18245 = reg18244
			}
			reg18247 = reg18245
		} else {
			reg18246 := False
			reg18247 = reg18246
		}
		if reg18247 == True {
			reg18248 := MakeSymbol("@s")
			reg18249 := PrimTail(V782)
			reg18250 := PrimHead(reg18249)
			reg18251 := MakeSymbol("@s")
			reg18252 := PrimTail(V782)
			reg18253 := PrimTail(reg18252)
			reg18254 := PrimCons(reg18251, reg18253)
			reg18255 := __e.Call(__defun__shen_4_8s_1macro, reg18254)
			reg18256 := Nil
			reg18257 := PrimCons(reg18255, reg18256)
			reg18258 := PrimCons(reg18250, reg18257)
			reg18259 := PrimCons(reg18248, reg18258)
			__ctx.Return(reg18259)
			return
		} else {
			reg18260 := PrimIsPair(V782)
			var reg18301 Obj
			if reg18260 == True {
				reg18261 := MakeSymbol("@s")
				reg18262 := PrimHead(V782)
				reg18263 := PrimEqual(reg18261, reg18262)
				var reg18296 Obj
				if reg18263 == True {
					reg18264 := PrimTail(V782)
					reg18265 := PrimIsPair(reg18264)
					var reg18291 Obj
					if reg18265 == True {
						reg18266 := PrimTail(V782)
						reg18267 := PrimTail(reg18266)
						reg18268 := PrimIsPair(reg18267)
						var reg18286 Obj
						if reg18268 == True {
							reg18269 := Nil
							reg18270 := PrimTail(V782)
							reg18271 := PrimTail(reg18270)
							reg18272 := PrimTail(reg18271)
							reg18273 := PrimEqual(reg18269, reg18272)
							var reg18281 Obj
							if reg18273 == True {
								reg18274 := PrimTail(V782)
								reg18275 := PrimHead(reg18274)
								reg18276 := PrimIsString(reg18275)
								var reg18279 Obj
								if reg18276 == True {
									reg18277 := True
									reg18279 = reg18277
								} else {
									reg18278 := False
									reg18279 = reg18278
								}
								reg18281 = reg18279
							} else {
								reg18280 := False
								reg18281 = reg18280
							}
							var reg18284 Obj
							if reg18281 == True {
								reg18282 := True
								reg18284 = reg18282
							} else {
								reg18283 := False
								reg18284 = reg18283
							}
							reg18286 = reg18284
						} else {
							reg18285 := False
							reg18286 = reg18285
						}
						var reg18289 Obj
						if reg18286 == True {
							reg18287 := True
							reg18289 = reg18287
						} else {
							reg18288 := False
							reg18289 = reg18288
						}
						reg18291 = reg18289
					} else {
						reg18290 := False
						reg18291 = reg18290
					}
					var reg18294 Obj
					if reg18291 == True {
						reg18292 := True
						reg18294 = reg18292
					} else {
						reg18293 := False
						reg18294 = reg18293
					}
					reg18296 = reg18294
				} else {
					reg18295 := False
					reg18296 = reg18295
				}
				var reg18299 Obj
				if reg18296 == True {
					reg18297 := True
					reg18299 = reg18297
				} else {
					reg18298 := False
					reg18299 = reg18298
				}
				reg18301 = reg18299
			} else {
				reg18300 := False
				reg18301 = reg18300
			}
			if reg18301 == True {
				reg18302 := PrimTail(V782)
				reg18303 := PrimHead(reg18302)
				reg18304 := __e.Call(__defun__explode, reg18303)
				E := reg18304
				_ = E
				reg18305 := __e.Call(__defun__length, E)
				reg18306 := MakeNumber(1)
				reg18307 := PrimGreatThan(reg18305, reg18306)
				if reg18307 == True {
					reg18308 := MakeSymbol("@s")
					reg18309 := PrimTail(V782)
					reg18310 := PrimTail(reg18309)
					reg18311 := __e.Call(__defun__append, E, reg18310)
					reg18312 := PrimCons(reg18308, reg18311)
					__ctx.TailApply(__defun__shen_4_8s_1macro, reg18312)
					return
				} else {
					__ctx.Return(V782)
					return
				}
			} else {
				__ctx.Return(V782)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.@s-macro", value: __defun__shen_4_8s_1macro})

	__defun__shen_4synonyms_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V784 := __args[0]
		_ = V784
		reg18314 := PrimIsPair(V784)
		var reg18322 Obj
		if reg18314 == True {
			reg18315 := MakeSymbol("synonyms")
			reg18316 := PrimHead(V784)
			reg18317 := PrimEqual(reg18315, reg18316)
			var reg18320 Obj
			if reg18317 == True {
				reg18318 := True
				reg18320 = reg18318
			} else {
				reg18319 := False
				reg18320 = reg18319
			}
			reg18322 = reg18320
		} else {
			reg18321 := False
			reg18322 = reg18321
		}
		if reg18322 == True {
			reg18323 := MakeSymbol("shen.synonyms-help")
			reg18324 := PrimTail(V784)
			reg18325 := __e.Call(__defun__shen_4curry_1synonyms, reg18324)
			reg18326 := __e.Call(__defun__shen_4rcons__form, reg18325)
			reg18327 := Nil
			reg18328 := PrimCons(reg18326, reg18327)
			reg18329 := PrimCons(reg18323, reg18328)
			__ctx.Return(reg18329)
			return
		} else {
			__ctx.Return(V784)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.synonyms-macro", value: __defun__shen_4synonyms_1macro})

	__defun__shen_4curry_1synonyms = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V786 := __args[0]
		_ = V786
		reg18330 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			X := __args[0]
			_ = X
			__ctx.TailApply(__defun__shen_4curry_1type, X)
			return
		}, 1)
		__ctx.TailApply(__defun__map, reg18330, V786)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.curry-synonyms", value: __defun__shen_4curry_1synonyms})

	__defun__shen_4nl_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V788 := __args[0]
		_ = V788
		reg18333 := PrimIsPair(V788)
		var reg18349 Obj
		if reg18333 == True {
			reg18334 := MakeSymbol("nl")
			reg18335 := PrimHead(V788)
			reg18336 := PrimEqual(reg18334, reg18335)
			var reg18344 Obj
			if reg18336 == True {
				reg18337 := Nil
				reg18338 := PrimTail(V788)
				reg18339 := PrimEqual(reg18337, reg18338)
				var reg18342 Obj
				if reg18339 == True {
					reg18340 := True
					reg18342 = reg18340
				} else {
					reg18341 := False
					reg18342 = reg18341
				}
				reg18344 = reg18342
			} else {
				reg18343 := False
				reg18344 = reg18343
			}
			var reg18347 Obj
			if reg18344 == True {
				reg18345 := True
				reg18347 = reg18345
			} else {
				reg18346 := False
				reg18347 = reg18346
			}
			reg18349 = reg18347
		} else {
			reg18348 := False
			reg18349 = reg18348
		}
		if reg18349 == True {
			reg18350 := MakeSymbol("nl")
			reg18351 := MakeNumber(1)
			reg18352 := Nil
			reg18353 := PrimCons(reg18351, reg18352)
			reg18354 := PrimCons(reg18350, reg18353)
			__ctx.Return(reg18354)
			return
		} else {
			__ctx.Return(V788)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.nl-macro", value: __defun__shen_4nl_1macro})

	__defun__shen_4assoc_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V790 := __args[0]
		_ = V790
		reg18355 := PrimIsPair(V790)
		var reg18403 Obj
		if reg18355 == True {
			reg18356 := PrimTail(V790)
			reg18357 := PrimIsPair(reg18356)
			var reg18398 Obj
			if reg18357 == True {
				reg18358 := PrimTail(V790)
				reg18359 := PrimTail(reg18358)
				reg18360 := PrimIsPair(reg18359)
				var reg18393 Obj
				if reg18360 == True {
					reg18361 := PrimTail(V790)
					reg18362 := PrimTail(reg18361)
					reg18363 := PrimTail(reg18362)
					reg18364 := PrimIsPair(reg18363)
					var reg18388 Obj
					if reg18364 == True {
						reg18365 := PrimHead(V790)
						reg18366 := MakeSymbol("@p")
						reg18367 := MakeSymbol("@v")
						reg18368 := MakeSymbol("append")
						reg18369 := MakeSymbol("and")
						reg18370 := MakeSymbol("or")
						reg18371 := MakeSymbol("+")
						reg18372 := MakeSymbol("*")
						reg18373 := MakeSymbol("do")
						reg18374 := Nil
						reg18375 := PrimCons(reg18373, reg18374)
						reg18376 := PrimCons(reg18372, reg18375)
						reg18377 := PrimCons(reg18371, reg18376)
						reg18378 := PrimCons(reg18370, reg18377)
						reg18379 := PrimCons(reg18369, reg18378)
						reg18380 := PrimCons(reg18368, reg18379)
						reg18381 := PrimCons(reg18367, reg18380)
						reg18382 := PrimCons(reg18366, reg18381)
						reg18383 := __e.Call(__defun__element_2, reg18365, reg18382)
						var reg18386 Obj
						if reg18383 == True {
							reg18384 := True
							reg18386 = reg18384
						} else {
							reg18385 := False
							reg18386 = reg18385
						}
						reg18388 = reg18386
					} else {
						reg18387 := False
						reg18388 = reg18387
					}
					var reg18391 Obj
					if reg18388 == True {
						reg18389 := True
						reg18391 = reg18389
					} else {
						reg18390 := False
						reg18391 = reg18390
					}
					reg18393 = reg18391
				} else {
					reg18392 := False
					reg18393 = reg18392
				}
				var reg18396 Obj
				if reg18393 == True {
					reg18394 := True
					reg18396 = reg18394
				} else {
					reg18395 := False
					reg18396 = reg18395
				}
				reg18398 = reg18396
			} else {
				reg18397 := False
				reg18398 = reg18397
			}
			var reg18401 Obj
			if reg18398 == True {
				reg18399 := True
				reg18401 = reg18399
			} else {
				reg18400 := False
				reg18401 = reg18400
			}
			reg18403 = reg18401
		} else {
			reg18402 := False
			reg18403 = reg18402
		}
		if reg18403 == True {
			reg18404 := PrimHead(V790)
			reg18405 := PrimTail(V790)
			reg18406 := PrimHead(reg18405)
			reg18407 := PrimHead(V790)
			reg18408 := PrimTail(V790)
			reg18409 := PrimTail(reg18408)
			reg18410 := PrimCons(reg18407, reg18409)
			reg18411 := __e.Call(__defun__shen_4assoc_1macro, reg18410)
			reg18412 := Nil
			reg18413 := PrimCons(reg18411, reg18412)
			reg18414 := PrimCons(reg18406, reg18413)
			reg18415 := PrimCons(reg18404, reg18414)
			__ctx.Return(reg18415)
			return
		} else {
			__ctx.Return(V790)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.assoc-macro", value: __defun__shen_4assoc_1macro})

	__defun__shen_4let_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V792 := __args[0]
		_ = V792
		reg18416 := PrimIsPair(V792)
		var reg18458 Obj
		if reg18416 == True {
			reg18417 := MakeSymbol("let")
			reg18418 := PrimHead(V792)
			reg18419 := PrimEqual(reg18417, reg18418)
			var reg18453 Obj
			if reg18419 == True {
				reg18420 := PrimTail(V792)
				reg18421 := PrimIsPair(reg18420)
				var reg18448 Obj
				if reg18421 == True {
					reg18422 := PrimTail(V792)
					reg18423 := PrimTail(reg18422)
					reg18424 := PrimIsPair(reg18423)
					var reg18443 Obj
					if reg18424 == True {
						reg18425 := PrimTail(V792)
						reg18426 := PrimTail(reg18425)
						reg18427 := PrimTail(reg18426)
						reg18428 := PrimIsPair(reg18427)
						var reg18438 Obj
						if reg18428 == True {
							reg18429 := PrimTail(V792)
							reg18430 := PrimTail(reg18429)
							reg18431 := PrimTail(reg18430)
							reg18432 := PrimTail(reg18431)
							reg18433 := PrimIsPair(reg18432)
							var reg18436 Obj
							if reg18433 == True {
								reg18434 := True
								reg18436 = reg18434
							} else {
								reg18435 := False
								reg18436 = reg18435
							}
							reg18438 = reg18436
						} else {
							reg18437 := False
							reg18438 = reg18437
						}
						var reg18441 Obj
						if reg18438 == True {
							reg18439 := True
							reg18441 = reg18439
						} else {
							reg18440 := False
							reg18441 = reg18440
						}
						reg18443 = reg18441
					} else {
						reg18442 := False
						reg18443 = reg18442
					}
					var reg18446 Obj
					if reg18443 == True {
						reg18444 := True
						reg18446 = reg18444
					} else {
						reg18445 := False
						reg18446 = reg18445
					}
					reg18448 = reg18446
				} else {
					reg18447 := False
					reg18448 = reg18447
				}
				var reg18451 Obj
				if reg18448 == True {
					reg18449 := True
					reg18451 = reg18449
				} else {
					reg18450 := False
					reg18451 = reg18450
				}
				reg18453 = reg18451
			} else {
				reg18452 := False
				reg18453 = reg18452
			}
			var reg18456 Obj
			if reg18453 == True {
				reg18454 := True
				reg18456 = reg18454
			} else {
				reg18455 := False
				reg18456 = reg18455
			}
			reg18458 = reg18456
		} else {
			reg18457 := False
			reg18458 = reg18457
		}
		if reg18458 == True {
			reg18459 := MakeSymbol("let")
			reg18460 := PrimTail(V792)
			reg18461 := PrimHead(reg18460)
			reg18462 := PrimTail(V792)
			reg18463 := PrimTail(reg18462)
			reg18464 := PrimHead(reg18463)
			reg18465 := MakeSymbol("let")
			reg18466 := PrimTail(V792)
			reg18467 := PrimTail(reg18466)
			reg18468 := PrimTail(reg18467)
			reg18469 := PrimCons(reg18465, reg18468)
			reg18470 := __e.Call(__defun__shen_4let_1macro, reg18469)
			reg18471 := Nil
			reg18472 := PrimCons(reg18470, reg18471)
			reg18473 := PrimCons(reg18464, reg18472)
			reg18474 := PrimCons(reg18461, reg18473)
			reg18475 := PrimCons(reg18459, reg18474)
			__ctx.Return(reg18475)
			return
		} else {
			__ctx.Return(V792)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.let-macro", value: __defun__shen_4let_1macro})

	__defun__shen_4abs_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V794 := __args[0]
		_ = V794
		reg18476 := PrimIsPair(V794)
		var reg18508 Obj
		if reg18476 == True {
			reg18477 := MakeSymbol("/.")
			reg18478 := PrimHead(V794)
			reg18479 := PrimEqual(reg18477, reg18478)
			var reg18503 Obj
			if reg18479 == True {
				reg18480 := PrimTail(V794)
				reg18481 := PrimIsPair(reg18480)
				var reg18498 Obj
				if reg18481 == True {
					reg18482 := PrimTail(V794)
					reg18483 := PrimTail(reg18482)
					reg18484 := PrimIsPair(reg18483)
					var reg18493 Obj
					if reg18484 == True {
						reg18485 := PrimTail(V794)
						reg18486 := PrimTail(reg18485)
						reg18487 := PrimTail(reg18486)
						reg18488 := PrimIsPair(reg18487)
						var reg18491 Obj
						if reg18488 == True {
							reg18489 := True
							reg18491 = reg18489
						} else {
							reg18490 := False
							reg18491 = reg18490
						}
						reg18493 = reg18491
					} else {
						reg18492 := False
						reg18493 = reg18492
					}
					var reg18496 Obj
					if reg18493 == True {
						reg18494 := True
						reg18496 = reg18494
					} else {
						reg18495 := False
						reg18496 = reg18495
					}
					reg18498 = reg18496
				} else {
					reg18497 := False
					reg18498 = reg18497
				}
				var reg18501 Obj
				if reg18498 == True {
					reg18499 := True
					reg18501 = reg18499
				} else {
					reg18500 := False
					reg18501 = reg18500
				}
				reg18503 = reg18501
			} else {
				reg18502 := False
				reg18503 = reg18502
			}
			var reg18506 Obj
			if reg18503 == True {
				reg18504 := True
				reg18506 = reg18504
			} else {
				reg18505 := False
				reg18506 = reg18505
			}
			reg18508 = reg18506
		} else {
			reg18507 := False
			reg18508 = reg18507
		}
		if reg18508 == True {
			reg18509 := MakeSymbol("lambda")
			reg18510 := PrimTail(V794)
			reg18511 := PrimHead(reg18510)
			reg18512 := MakeSymbol("/.")
			reg18513 := PrimTail(V794)
			reg18514 := PrimTail(reg18513)
			reg18515 := PrimCons(reg18512, reg18514)
			reg18516 := __e.Call(__defun__shen_4abs_1macro, reg18515)
			reg18517 := Nil
			reg18518 := PrimCons(reg18516, reg18517)
			reg18519 := PrimCons(reg18511, reg18518)
			reg18520 := PrimCons(reg18509, reg18519)
			__ctx.Return(reg18520)
			return
		} else {
			reg18521 := PrimIsPair(V794)
			var reg18554 Obj
			if reg18521 == True {
				reg18522 := MakeSymbol("/.")
				reg18523 := PrimHead(V794)
				reg18524 := PrimEqual(reg18522, reg18523)
				var reg18549 Obj
				if reg18524 == True {
					reg18525 := PrimTail(V794)
					reg18526 := PrimIsPair(reg18525)
					var reg18544 Obj
					if reg18526 == True {
						reg18527 := PrimTail(V794)
						reg18528 := PrimTail(reg18527)
						reg18529 := PrimIsPair(reg18528)
						var reg18539 Obj
						if reg18529 == True {
							reg18530 := Nil
							reg18531 := PrimTail(V794)
							reg18532 := PrimTail(reg18531)
							reg18533 := PrimTail(reg18532)
							reg18534 := PrimEqual(reg18530, reg18533)
							var reg18537 Obj
							if reg18534 == True {
								reg18535 := True
								reg18537 = reg18535
							} else {
								reg18536 := False
								reg18537 = reg18536
							}
							reg18539 = reg18537
						} else {
							reg18538 := False
							reg18539 = reg18538
						}
						var reg18542 Obj
						if reg18539 == True {
							reg18540 := True
							reg18542 = reg18540
						} else {
							reg18541 := False
							reg18542 = reg18541
						}
						reg18544 = reg18542
					} else {
						reg18543 := False
						reg18544 = reg18543
					}
					var reg18547 Obj
					if reg18544 == True {
						reg18545 := True
						reg18547 = reg18545
					} else {
						reg18546 := False
						reg18547 = reg18546
					}
					reg18549 = reg18547
				} else {
					reg18548 := False
					reg18549 = reg18548
				}
				var reg18552 Obj
				if reg18549 == True {
					reg18550 := True
					reg18552 = reg18550
				} else {
					reg18551 := False
					reg18552 = reg18551
				}
				reg18554 = reg18552
			} else {
				reg18553 := False
				reg18554 = reg18553
			}
			if reg18554 == True {
				reg18555 := MakeSymbol("lambda")
				reg18556 := PrimTail(V794)
				reg18557 := PrimCons(reg18555, reg18556)
				__ctx.Return(reg18557)
				return
			} else {
				__ctx.Return(V794)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.abs-macro", value: __defun__shen_4abs_1macro})

	__defun__shen_4cases_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V798 := __args[0]
		_ = V798
		reg18558 := PrimIsPair(V798)
		var reg18590 Obj
		if reg18558 == True {
			reg18559 := MakeSymbol("cases")
			reg18560 := PrimHead(V798)
			reg18561 := PrimEqual(reg18559, reg18560)
			var reg18585 Obj
			if reg18561 == True {
				reg18562 := PrimTail(V798)
				reg18563 := PrimIsPair(reg18562)
				var reg18580 Obj
				if reg18563 == True {
					reg18564 := True
					reg18565 := PrimTail(V798)
					reg18566 := PrimHead(reg18565)
					reg18567 := PrimEqual(reg18564, reg18566)
					var reg18575 Obj
					if reg18567 == True {
						reg18568 := PrimTail(V798)
						reg18569 := PrimTail(reg18568)
						reg18570 := PrimIsPair(reg18569)
						var reg18573 Obj
						if reg18570 == True {
							reg18571 := True
							reg18573 = reg18571
						} else {
							reg18572 := False
							reg18573 = reg18572
						}
						reg18575 = reg18573
					} else {
						reg18574 := False
						reg18575 = reg18574
					}
					var reg18578 Obj
					if reg18575 == True {
						reg18576 := True
						reg18578 = reg18576
					} else {
						reg18577 := False
						reg18578 = reg18577
					}
					reg18580 = reg18578
				} else {
					reg18579 := False
					reg18580 = reg18579
				}
				var reg18583 Obj
				if reg18580 == True {
					reg18581 := True
					reg18583 = reg18581
				} else {
					reg18582 := False
					reg18583 = reg18582
				}
				reg18585 = reg18583
			} else {
				reg18584 := False
				reg18585 = reg18584
			}
			var reg18588 Obj
			if reg18585 == True {
				reg18586 := True
				reg18588 = reg18586
			} else {
				reg18587 := False
				reg18588 = reg18587
			}
			reg18590 = reg18588
		} else {
			reg18589 := False
			reg18590 = reg18589
		}
		if reg18590 == True {
			reg18591 := PrimTail(V798)
			reg18592 := PrimTail(reg18591)
			reg18593 := PrimHead(reg18592)
			__ctx.Return(reg18593)
			return
		} else {
			reg18594 := PrimIsPair(V798)
			var reg18627 Obj
			if reg18594 == True {
				reg18595 := MakeSymbol("cases")
				reg18596 := PrimHead(V798)
				reg18597 := PrimEqual(reg18595, reg18596)
				var reg18622 Obj
				if reg18597 == True {
					reg18598 := PrimTail(V798)
					reg18599 := PrimIsPair(reg18598)
					var reg18617 Obj
					if reg18599 == True {
						reg18600 := PrimTail(V798)
						reg18601 := PrimTail(reg18600)
						reg18602 := PrimIsPair(reg18601)
						var reg18612 Obj
						if reg18602 == True {
							reg18603 := Nil
							reg18604 := PrimTail(V798)
							reg18605 := PrimTail(reg18604)
							reg18606 := PrimTail(reg18605)
							reg18607 := PrimEqual(reg18603, reg18606)
							var reg18610 Obj
							if reg18607 == True {
								reg18608 := True
								reg18610 = reg18608
							} else {
								reg18609 := False
								reg18610 = reg18609
							}
							reg18612 = reg18610
						} else {
							reg18611 := False
							reg18612 = reg18611
						}
						var reg18615 Obj
						if reg18612 == True {
							reg18613 := True
							reg18615 = reg18613
						} else {
							reg18614 := False
							reg18615 = reg18614
						}
						reg18617 = reg18615
					} else {
						reg18616 := False
						reg18617 = reg18616
					}
					var reg18620 Obj
					if reg18617 == True {
						reg18618 := True
						reg18620 = reg18618
					} else {
						reg18619 := False
						reg18620 = reg18619
					}
					reg18622 = reg18620
				} else {
					reg18621 := False
					reg18622 = reg18621
				}
				var reg18625 Obj
				if reg18622 == True {
					reg18623 := True
					reg18625 = reg18623
				} else {
					reg18624 := False
					reg18625 = reg18624
				}
				reg18627 = reg18625
			} else {
				reg18626 := False
				reg18627 = reg18626
			}
			if reg18627 == True {
				reg18628 := MakeSymbol("if")
				reg18629 := PrimTail(V798)
				reg18630 := PrimHead(reg18629)
				reg18631 := PrimTail(V798)
				reg18632 := PrimTail(reg18631)
				reg18633 := PrimHead(reg18632)
				reg18634 := MakeSymbol("simple-error")
				reg18635 := MakeString("error: cases exhausted")
				reg18636 := Nil
				reg18637 := PrimCons(reg18635, reg18636)
				reg18638 := PrimCons(reg18634, reg18637)
				reg18639 := Nil
				reg18640 := PrimCons(reg18638, reg18639)
				reg18641 := PrimCons(reg18633, reg18640)
				reg18642 := PrimCons(reg18630, reg18641)
				reg18643 := PrimCons(reg18628, reg18642)
				__ctx.Return(reg18643)
				return
			} else {
				reg18644 := PrimIsPair(V798)
				var reg18667 Obj
				if reg18644 == True {
					reg18645 := MakeSymbol("cases")
					reg18646 := PrimHead(V798)
					reg18647 := PrimEqual(reg18645, reg18646)
					var reg18662 Obj
					if reg18647 == True {
						reg18648 := PrimTail(V798)
						reg18649 := PrimIsPair(reg18648)
						var reg18657 Obj
						if reg18649 == True {
							reg18650 := PrimTail(V798)
							reg18651 := PrimTail(reg18650)
							reg18652 := PrimIsPair(reg18651)
							var reg18655 Obj
							if reg18652 == True {
								reg18653 := True
								reg18655 = reg18653
							} else {
								reg18654 := False
								reg18655 = reg18654
							}
							reg18657 = reg18655
						} else {
							reg18656 := False
							reg18657 = reg18656
						}
						var reg18660 Obj
						if reg18657 == True {
							reg18658 := True
							reg18660 = reg18658
						} else {
							reg18659 := False
							reg18660 = reg18659
						}
						reg18662 = reg18660
					} else {
						reg18661 := False
						reg18662 = reg18661
					}
					var reg18665 Obj
					if reg18662 == True {
						reg18663 := True
						reg18665 = reg18663
					} else {
						reg18664 := False
						reg18665 = reg18664
					}
					reg18667 = reg18665
				} else {
					reg18666 := False
					reg18667 = reg18666
				}
				if reg18667 == True {
					reg18668 := MakeSymbol("if")
					reg18669 := PrimTail(V798)
					reg18670 := PrimHead(reg18669)
					reg18671 := PrimTail(V798)
					reg18672 := PrimTail(reg18671)
					reg18673 := PrimHead(reg18672)
					reg18674 := MakeSymbol("cases")
					reg18675 := PrimTail(V798)
					reg18676 := PrimTail(reg18675)
					reg18677 := PrimTail(reg18676)
					reg18678 := PrimCons(reg18674, reg18677)
					reg18679 := __e.Call(__defun__shen_4cases_1macro, reg18678)
					reg18680 := Nil
					reg18681 := PrimCons(reg18679, reg18680)
					reg18682 := PrimCons(reg18673, reg18681)
					reg18683 := PrimCons(reg18670, reg18682)
					reg18684 := PrimCons(reg18668, reg18683)
					__ctx.Return(reg18684)
					return
				} else {
					reg18685 := PrimIsPair(V798)
					var reg18709 Obj
					if reg18685 == True {
						reg18686 := MakeSymbol("cases")
						reg18687 := PrimHead(V798)
						reg18688 := PrimEqual(reg18686, reg18687)
						var reg18704 Obj
						if reg18688 == True {
							reg18689 := PrimTail(V798)
							reg18690 := PrimIsPair(reg18689)
							var reg18699 Obj
							if reg18690 == True {
								reg18691 := Nil
								reg18692 := PrimTail(V798)
								reg18693 := PrimTail(reg18692)
								reg18694 := PrimEqual(reg18691, reg18693)
								var reg18697 Obj
								if reg18694 == True {
									reg18695 := True
									reg18697 = reg18695
								} else {
									reg18696 := False
									reg18697 = reg18696
								}
								reg18699 = reg18697
							} else {
								reg18698 := False
								reg18699 = reg18698
							}
							var reg18702 Obj
							if reg18699 == True {
								reg18700 := True
								reg18702 = reg18700
							} else {
								reg18701 := False
								reg18702 = reg18701
							}
							reg18704 = reg18702
						} else {
							reg18703 := False
							reg18704 = reg18703
						}
						var reg18707 Obj
						if reg18704 == True {
							reg18705 := True
							reg18707 = reg18705
						} else {
							reg18706 := False
							reg18707 = reg18706
						}
						reg18709 = reg18707
					} else {
						reg18708 := False
						reg18709 = reg18708
					}
					if reg18709 == True {
						reg18710 := MakeString("error: odd number of case elements\n")
						reg18711 := PrimSimpleError(reg18710)
						__ctx.Return(reg18711)
						return
					} else {
						__ctx.Return(V798)
						return
					}
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.cases-macro", value: __defun__shen_4cases_1macro})

	__defun__shen_4timer_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V800 := __args[0]
		_ = V800
		reg18712 := PrimIsPair(V800)
		var reg18736 Obj
		if reg18712 == True {
			reg18713 := MakeSymbol("time")
			reg18714 := PrimHead(V800)
			reg18715 := PrimEqual(reg18713, reg18714)
			var reg18731 Obj
			if reg18715 == True {
				reg18716 := PrimTail(V800)
				reg18717 := PrimIsPair(reg18716)
				var reg18726 Obj
				if reg18717 == True {
					reg18718 := Nil
					reg18719 := PrimTail(V800)
					reg18720 := PrimTail(reg18719)
					reg18721 := PrimEqual(reg18718, reg18720)
					var reg18724 Obj
					if reg18721 == True {
						reg18722 := True
						reg18724 = reg18722
					} else {
						reg18723 := False
						reg18724 = reg18723
					}
					reg18726 = reg18724
				} else {
					reg18725 := False
					reg18726 = reg18725
				}
				var reg18729 Obj
				if reg18726 == True {
					reg18727 := True
					reg18729 = reg18727
				} else {
					reg18728 := False
					reg18729 = reg18728
				}
				reg18731 = reg18729
			} else {
				reg18730 := False
				reg18731 = reg18730
			}
			var reg18734 Obj
			if reg18731 == True {
				reg18732 := True
				reg18734 = reg18732
			} else {
				reg18733 := False
				reg18734 = reg18733
			}
			reg18736 = reg18734
		} else {
			reg18735 := False
			reg18736 = reg18735
		}
		if reg18736 == True {
			reg18737 := MakeSymbol("let")
			reg18738 := MakeSymbol("Start")
			reg18739 := MakeSymbol("get-time")
			reg18740 := MakeSymbol("run")
			reg18741 := Nil
			reg18742 := PrimCons(reg18740, reg18741)
			reg18743 := PrimCons(reg18739, reg18742)
			reg18744 := MakeSymbol("Result")
			reg18745 := PrimTail(V800)
			reg18746 := PrimHead(reg18745)
			reg18747 := MakeSymbol("Finish")
			reg18748 := MakeSymbol("get-time")
			reg18749 := MakeSymbol("run")
			reg18750 := Nil
			reg18751 := PrimCons(reg18749, reg18750)
			reg18752 := PrimCons(reg18748, reg18751)
			reg18753 := MakeSymbol("Time")
			reg18754 := MakeSymbol("-")
			reg18755 := MakeSymbol("Finish")
			reg18756 := MakeSymbol("Start")
			reg18757 := Nil
			reg18758 := PrimCons(reg18756, reg18757)
			reg18759 := PrimCons(reg18755, reg18758)
			reg18760 := PrimCons(reg18754, reg18759)
			reg18761 := MakeSymbol("Message")
			reg18762 := MakeSymbol("shen.prhush")
			reg18763 := MakeSymbol("cn")
			reg18764 := MakeString("\nrun time: ")
			reg18765 := MakeSymbol("cn")
			reg18766 := MakeSymbol("str")
			reg18767 := MakeSymbol("Time")
			reg18768 := Nil
			reg18769 := PrimCons(reg18767, reg18768)
			reg18770 := PrimCons(reg18766, reg18769)
			reg18771 := MakeString(" secs\n")
			reg18772 := Nil
			reg18773 := PrimCons(reg18771, reg18772)
			reg18774 := PrimCons(reg18770, reg18773)
			reg18775 := PrimCons(reg18765, reg18774)
			reg18776 := Nil
			reg18777 := PrimCons(reg18775, reg18776)
			reg18778 := PrimCons(reg18764, reg18777)
			reg18779 := PrimCons(reg18763, reg18778)
			reg18780 := MakeSymbol("stoutput")
			reg18781 := Nil
			reg18782 := PrimCons(reg18780, reg18781)
			reg18783 := Nil
			reg18784 := PrimCons(reg18782, reg18783)
			reg18785 := PrimCons(reg18779, reg18784)
			reg18786 := PrimCons(reg18762, reg18785)
			reg18787 := MakeSymbol("Result")
			reg18788 := Nil
			reg18789 := PrimCons(reg18787, reg18788)
			reg18790 := PrimCons(reg18786, reg18789)
			reg18791 := PrimCons(reg18761, reg18790)
			reg18792 := PrimCons(reg18760, reg18791)
			reg18793 := PrimCons(reg18753, reg18792)
			reg18794 := PrimCons(reg18752, reg18793)
			reg18795 := PrimCons(reg18747, reg18794)
			reg18796 := PrimCons(reg18746, reg18795)
			reg18797 := PrimCons(reg18744, reg18796)
			reg18798 := PrimCons(reg18743, reg18797)
			reg18799 := PrimCons(reg18738, reg18798)
			reg18800 := PrimCons(reg18737, reg18799)
			__ctx.TailApply(__defun__shen_4let_1macro, reg18800)
			return
		} else {
			__ctx.Return(V800)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.timer-macro", value: __defun__shen_4timer_1macro})

	__defun__shen_4tuple_1up = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V802 := __args[0]
		_ = V802
		reg18802 := PrimIsPair(V802)
		if reg18802 == True {
			reg18803 := MakeSymbol("@p")
			reg18804 := PrimHead(V802)
			reg18805 := PrimTail(V802)
			reg18806 := __e.Call(__defun__shen_4tuple_1up, reg18805)
			reg18807 := Nil
			reg18808 := PrimCons(reg18806, reg18807)
			reg18809 := PrimCons(reg18804, reg18808)
			reg18810 := PrimCons(reg18803, reg18809)
			__ctx.Return(reg18810)
			return
		} else {
			__ctx.Return(V802)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.tuple-up", value: __defun__shen_4tuple_1up})

	__defun__shen_4put_cget_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V804 := __args[0]
		_ = V804
		reg18811 := PrimIsPair(V804)
		var reg18854 Obj
		if reg18811 == True {
			reg18812 := MakeSymbol("put")
			reg18813 := PrimHead(V804)
			reg18814 := PrimEqual(reg18812, reg18813)
			var reg18849 Obj
			if reg18814 == True {
				reg18815 := PrimTail(V804)
				reg18816 := PrimIsPair(reg18815)
				var reg18844 Obj
				if reg18816 == True {
					reg18817 := PrimTail(V804)
					reg18818 := PrimTail(reg18817)
					reg18819 := PrimIsPair(reg18818)
					var reg18839 Obj
					if reg18819 == True {
						reg18820 := PrimTail(V804)
						reg18821 := PrimTail(reg18820)
						reg18822 := PrimTail(reg18821)
						reg18823 := PrimIsPair(reg18822)
						var reg18834 Obj
						if reg18823 == True {
							reg18824 := Nil
							reg18825 := PrimTail(V804)
							reg18826 := PrimTail(reg18825)
							reg18827 := PrimTail(reg18826)
							reg18828 := PrimTail(reg18827)
							reg18829 := PrimEqual(reg18824, reg18828)
							var reg18832 Obj
							if reg18829 == True {
								reg18830 := True
								reg18832 = reg18830
							} else {
								reg18831 := False
								reg18832 = reg18831
							}
							reg18834 = reg18832
						} else {
							reg18833 := False
							reg18834 = reg18833
						}
						var reg18837 Obj
						if reg18834 == True {
							reg18835 := True
							reg18837 = reg18835
						} else {
							reg18836 := False
							reg18837 = reg18836
						}
						reg18839 = reg18837
					} else {
						reg18838 := False
						reg18839 = reg18838
					}
					var reg18842 Obj
					if reg18839 == True {
						reg18840 := True
						reg18842 = reg18840
					} else {
						reg18841 := False
						reg18842 = reg18841
					}
					reg18844 = reg18842
				} else {
					reg18843 := False
					reg18844 = reg18843
				}
				var reg18847 Obj
				if reg18844 == True {
					reg18845 := True
					reg18847 = reg18845
				} else {
					reg18846 := False
					reg18847 = reg18846
				}
				reg18849 = reg18847
			} else {
				reg18848 := False
				reg18849 = reg18848
			}
			var reg18852 Obj
			if reg18849 == True {
				reg18850 := True
				reg18852 = reg18850
			} else {
				reg18851 := False
				reg18852 = reg18851
			}
			reg18854 = reg18852
		} else {
			reg18853 := False
			reg18854 = reg18853
		}
		if reg18854 == True {
			reg18855 := MakeSymbol("put")
			reg18856 := PrimTail(V804)
			reg18857 := PrimHead(reg18856)
			reg18858 := PrimTail(V804)
			reg18859 := PrimTail(reg18858)
			reg18860 := PrimHead(reg18859)
			reg18861 := PrimTail(V804)
			reg18862 := PrimTail(reg18861)
			reg18863 := PrimTail(reg18862)
			reg18864 := PrimHead(reg18863)
			reg18865 := MakeSymbol("value")
			reg18866 := MakeSymbol("*property-vector*")
			reg18867 := Nil
			reg18868 := PrimCons(reg18866, reg18867)
			reg18869 := PrimCons(reg18865, reg18868)
			reg18870 := Nil
			reg18871 := PrimCons(reg18869, reg18870)
			reg18872 := PrimCons(reg18864, reg18871)
			reg18873 := PrimCons(reg18860, reg18872)
			reg18874 := PrimCons(reg18857, reg18873)
			reg18875 := PrimCons(reg18855, reg18874)
			__ctx.Return(reg18875)
			return
		} else {
			reg18876 := PrimIsPair(V804)
			var reg18909 Obj
			if reg18876 == True {
				reg18877 := MakeSymbol("get")
				reg18878 := PrimHead(V804)
				reg18879 := PrimEqual(reg18877, reg18878)
				var reg18904 Obj
				if reg18879 == True {
					reg18880 := PrimTail(V804)
					reg18881 := PrimIsPair(reg18880)
					var reg18899 Obj
					if reg18881 == True {
						reg18882 := PrimTail(V804)
						reg18883 := PrimTail(reg18882)
						reg18884 := PrimIsPair(reg18883)
						var reg18894 Obj
						if reg18884 == True {
							reg18885 := Nil
							reg18886 := PrimTail(V804)
							reg18887 := PrimTail(reg18886)
							reg18888 := PrimTail(reg18887)
							reg18889 := PrimEqual(reg18885, reg18888)
							var reg18892 Obj
							if reg18889 == True {
								reg18890 := True
								reg18892 = reg18890
							} else {
								reg18891 := False
								reg18892 = reg18891
							}
							reg18894 = reg18892
						} else {
							reg18893 := False
							reg18894 = reg18893
						}
						var reg18897 Obj
						if reg18894 == True {
							reg18895 := True
							reg18897 = reg18895
						} else {
							reg18896 := False
							reg18897 = reg18896
						}
						reg18899 = reg18897
					} else {
						reg18898 := False
						reg18899 = reg18898
					}
					var reg18902 Obj
					if reg18899 == True {
						reg18900 := True
						reg18902 = reg18900
					} else {
						reg18901 := False
						reg18902 = reg18901
					}
					reg18904 = reg18902
				} else {
					reg18903 := False
					reg18904 = reg18903
				}
				var reg18907 Obj
				if reg18904 == True {
					reg18905 := True
					reg18907 = reg18905
				} else {
					reg18906 := False
					reg18907 = reg18906
				}
				reg18909 = reg18907
			} else {
				reg18908 := False
				reg18909 = reg18908
			}
			if reg18909 == True {
				reg18910 := MakeSymbol("get")
				reg18911 := PrimTail(V804)
				reg18912 := PrimHead(reg18911)
				reg18913 := PrimTail(V804)
				reg18914 := PrimTail(reg18913)
				reg18915 := PrimHead(reg18914)
				reg18916 := MakeSymbol("value")
				reg18917 := MakeSymbol("*property-vector*")
				reg18918 := Nil
				reg18919 := PrimCons(reg18917, reg18918)
				reg18920 := PrimCons(reg18916, reg18919)
				reg18921 := Nil
				reg18922 := PrimCons(reg18920, reg18921)
				reg18923 := PrimCons(reg18915, reg18922)
				reg18924 := PrimCons(reg18912, reg18923)
				reg18925 := PrimCons(reg18910, reg18924)
				__ctx.Return(reg18925)
				return
			} else {
				reg18926 := PrimIsPair(V804)
				var reg18959 Obj
				if reg18926 == True {
					reg18927 := MakeSymbol("unput")
					reg18928 := PrimHead(V804)
					reg18929 := PrimEqual(reg18927, reg18928)
					var reg18954 Obj
					if reg18929 == True {
						reg18930 := PrimTail(V804)
						reg18931 := PrimIsPair(reg18930)
						var reg18949 Obj
						if reg18931 == True {
							reg18932 := PrimTail(V804)
							reg18933 := PrimTail(reg18932)
							reg18934 := PrimIsPair(reg18933)
							var reg18944 Obj
							if reg18934 == True {
								reg18935 := Nil
								reg18936 := PrimTail(V804)
								reg18937 := PrimTail(reg18936)
								reg18938 := PrimTail(reg18937)
								reg18939 := PrimEqual(reg18935, reg18938)
								var reg18942 Obj
								if reg18939 == True {
									reg18940 := True
									reg18942 = reg18940
								} else {
									reg18941 := False
									reg18942 = reg18941
								}
								reg18944 = reg18942
							} else {
								reg18943 := False
								reg18944 = reg18943
							}
							var reg18947 Obj
							if reg18944 == True {
								reg18945 := True
								reg18947 = reg18945
							} else {
								reg18946 := False
								reg18947 = reg18946
							}
							reg18949 = reg18947
						} else {
							reg18948 := False
							reg18949 = reg18948
						}
						var reg18952 Obj
						if reg18949 == True {
							reg18950 := True
							reg18952 = reg18950
						} else {
							reg18951 := False
							reg18952 = reg18951
						}
						reg18954 = reg18952
					} else {
						reg18953 := False
						reg18954 = reg18953
					}
					var reg18957 Obj
					if reg18954 == True {
						reg18955 := True
						reg18957 = reg18955
					} else {
						reg18956 := False
						reg18957 = reg18956
					}
					reg18959 = reg18957
				} else {
					reg18958 := False
					reg18959 = reg18958
				}
				if reg18959 == True {
					reg18960 := MakeSymbol("unput")
					reg18961 := PrimTail(V804)
					reg18962 := PrimHead(reg18961)
					reg18963 := PrimTail(V804)
					reg18964 := PrimTail(reg18963)
					reg18965 := PrimHead(reg18964)
					reg18966 := MakeSymbol("value")
					reg18967 := MakeSymbol("*property-vector*")
					reg18968 := Nil
					reg18969 := PrimCons(reg18967, reg18968)
					reg18970 := PrimCons(reg18966, reg18969)
					reg18971 := Nil
					reg18972 := PrimCons(reg18970, reg18971)
					reg18973 := PrimCons(reg18965, reg18972)
					reg18974 := PrimCons(reg18962, reg18973)
					reg18975 := PrimCons(reg18960, reg18974)
					__ctx.Return(reg18975)
					return
				} else {
					__ctx.Return(V804)
					return
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.put/get-macro", value: __defun__shen_4put_cget_1macro})

	__defun__shen_4function_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V806 := __args[0]
		_ = V806
		reg18976 := PrimIsPair(V806)
		var reg19000 Obj
		if reg18976 == True {
			reg18977 := MakeSymbol("function")
			reg18978 := PrimHead(V806)
			reg18979 := PrimEqual(reg18977, reg18978)
			var reg18995 Obj
			if reg18979 == True {
				reg18980 := PrimTail(V806)
				reg18981 := PrimIsPair(reg18980)
				var reg18990 Obj
				if reg18981 == True {
					reg18982 := Nil
					reg18983 := PrimTail(V806)
					reg18984 := PrimTail(reg18983)
					reg18985 := PrimEqual(reg18982, reg18984)
					var reg18988 Obj
					if reg18985 == True {
						reg18986 := True
						reg18988 = reg18986
					} else {
						reg18987 := False
						reg18988 = reg18987
					}
					reg18990 = reg18988
				} else {
					reg18989 := False
					reg18990 = reg18989
				}
				var reg18993 Obj
				if reg18990 == True {
					reg18991 := True
					reg18993 = reg18991
				} else {
					reg18992 := False
					reg18993 = reg18992
				}
				reg18995 = reg18993
			} else {
				reg18994 := False
				reg18995 = reg18994
			}
			var reg18998 Obj
			if reg18995 == True {
				reg18996 := True
				reg18998 = reg18996
			} else {
				reg18997 := False
				reg18998 = reg18997
			}
			reg19000 = reg18998
		} else {
			reg18999 := False
			reg19000 = reg18999
		}
		if reg19000 == True {
			reg19001 := PrimTail(V806)
			reg19002 := PrimHead(reg19001)
			reg19003 := PrimTail(V806)
			reg19004 := PrimHead(reg19003)
			reg19005 := __e.Call(__defun__arity, reg19004)
			__ctx.TailApply(__defun__shen_4function_1abstraction, reg19002, reg19005)
			return
		} else {
			__ctx.Return(V806)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.function-macro", value: __defun__shen_4function_1macro})

	__defun__shen_4function_1abstraction = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V809 := __args[0]
		_ = V809
		V810 := __args[1]
		_ = V810
		reg19007 := MakeNumber(0)
		reg19008 := PrimEqual(reg19007, V810)
		if reg19008 == True {
			reg19009 := MakeString(" has no lambda form\n")
			reg19010 := MakeSymbol("shen.a")
			reg19011 := __e.Call(__defun__shen_4app, V809, reg19009, reg19010)
			reg19012 := PrimSimpleError(reg19011)
			__ctx.Return(reg19012)
			return
		} else {
			reg19013 := MakeNumber(-1)
			reg19014 := PrimEqual(reg19013, V810)
			if reg19014 == True {
				reg19015 := MakeSymbol("function")
				reg19016 := Nil
				reg19017 := PrimCons(V809, reg19016)
				reg19018 := PrimCons(reg19015, reg19017)
				__ctx.Return(reg19018)
				return
			} else {
				reg19019 := Nil
				__ctx.TailApply(__defun__shen_4function_1abstraction_1help, V809, V810, reg19019)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.function-abstraction", value: __defun__shen_4function_1abstraction})

	__defun__shen_4function_1abstraction_1help = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V814 := __args[0]
		_ = V814
		V815 := __args[1]
		_ = V815
		V816 := __args[2]
		_ = V816
		reg19021 := MakeNumber(0)
		reg19022 := PrimEqual(reg19021, V815)
		if reg19022 == True {
			reg19023 := PrimCons(V814, V816)
			__ctx.Return(reg19023)
			return
		} else {
			reg19024 := MakeSymbol("V")
			reg19025 := __e.Call(__defun__gensym, reg19024)
			X := reg19025
			_ = X
			reg19026 := MakeSymbol("/.")
			reg19027 := MakeNumber(1)
			reg19028 := PrimNumberSubtract(V815, reg19027)
			reg19029 := Nil
			reg19030 := PrimCons(X, reg19029)
			reg19031 := __e.Call(__defun__append, V816, reg19030)
			reg19032 := __e.Call(__defun__shen_4function_1abstraction_1help, V814, reg19028, reg19031)
			reg19033 := Nil
			reg19034 := PrimCons(reg19032, reg19033)
			reg19035 := PrimCons(X, reg19034)
			reg19036 := PrimCons(reg19026, reg19035)
			__ctx.Return(reg19036)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.function-abstraction-help", value: __defun__shen_4function_1abstraction_1help})

	__defun__undefmacro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V818 := __args[0]
		_ = V818
		reg19037 := MakeSymbol("shen.*macroreg*")
		reg19038 := PrimValue(reg19037)
		MacroReg := reg19038
		_ = MacroReg
		reg19039 := __e.Call(__defun__shen_4findpos, V818, MacroReg)
		Pos := reg19039
		_ = Pos
		reg19040 := MakeSymbol("shen.*macroreg*")
		reg19041 := __e.Call(__defun__remove, V818, MacroReg)
		reg19042 := PrimSet(reg19040, reg19041)
		Remove1 := reg19042
		_ = Remove1
		reg19043 := MakeSymbol("*macros*")
		reg19044 := MakeSymbol("*macros*")
		reg19045 := PrimValue(reg19044)
		reg19046 := __e.Call(__defun__shen_4remove_1nth, Pos, reg19045)
		reg19047 := PrimSet(reg19043, reg19046)
		Remove2 := reg19047
		_ = Remove2
		__ctx.Return(V818)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "undefmacro", value: __defun__undefmacro})

	__defun__shen_4findpos = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V828 := __args[0]
		_ = V828
		V829 := __args[1]
		_ = V829
		reg19048 := Nil
		reg19049 := PrimEqual(reg19048, V829)
		if reg19049 == True {
			reg19050 := MakeString(" is not a macro\n")
			reg19051 := MakeSymbol("shen.a")
			reg19052 := __e.Call(__defun__shen_4app, V828, reg19050, reg19051)
			reg19053 := PrimSimpleError(reg19052)
			__ctx.Return(reg19053)
			return
		} else {
			reg19054 := PrimIsPair(V829)
			var reg19061 Obj
			if reg19054 == True {
				reg19055 := PrimHead(V829)
				reg19056 := PrimEqual(reg19055, V828)
				var reg19059 Obj
				if reg19056 == True {
					reg19057 := True
					reg19059 = reg19057
				} else {
					reg19058 := False
					reg19059 = reg19058
				}
				reg19061 = reg19059
			} else {
				reg19060 := False
				reg19061 = reg19060
			}
			if reg19061 == True {
				reg19062 := MakeNumber(1)
				__ctx.Return(reg19062)
				return
			} else {
				reg19063 := PrimIsPair(V829)
				if reg19063 == True {
					reg19064 := MakeNumber(1)
					reg19065 := PrimTail(V829)
					reg19066 := __e.Call(__defun__shen_4findpos, V828, reg19065)
					reg19067 := PrimNumberAdd(reg19064, reg19066)
					__ctx.Return(reg19067)
					return
				} else {
					reg19068 := MakeSymbol("shen.findpos")
					__ctx.TailApply(__defun__shen_4f__error, reg19068)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.findpos", value: __defun__shen_4findpos})

	__defun__shen_4remove_1nth = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V834 := __args[0]
		_ = V834
		V835 := __args[1]
		_ = V835
		reg19070 := MakeNumber(1)
		reg19071 := PrimEqual(reg19070, V834)
		var reg19077 Obj
		if reg19071 == True {
			reg19072 := PrimIsPair(V835)
			var reg19075 Obj
			if reg19072 == True {
				reg19073 := True
				reg19075 = reg19073
			} else {
				reg19074 := False
				reg19075 = reg19074
			}
			reg19077 = reg19075
		} else {
			reg19076 := False
			reg19077 = reg19076
		}
		if reg19077 == True {
			reg19078 := PrimTail(V835)
			__ctx.Return(reg19078)
			return
		} else {
			reg19079 := PrimIsPair(V835)
			if reg19079 == True {
				reg19080 := PrimHead(V835)
				reg19081 := MakeNumber(1)
				reg19082 := PrimNumberSubtract(V834, reg19081)
				reg19083 := PrimTail(V835)
				reg19084 := __e.Call(__defun__shen_4remove_1nth, reg19082, reg19083)
				reg19085 := PrimCons(reg19080, reg19084)
				__ctx.Return(reg19085)
				return
			} else {
				reg19086 := MakeSymbol("shen.remove-nth")
				__ctx.TailApply(__defun__shen_4f__error, reg19086)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.remove-nth", value: __defun__shen_4remove_1nth})

}
