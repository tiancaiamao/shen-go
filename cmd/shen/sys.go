package main

import . "github.com/tiancaiamao/shen-go/kl"

var __defun__thaw Obj                        // thaw
var __defun__eval Obj                        // eval
var __defun__shen_4eval_1without_1macros Obj // shen.eval-without-macros
var __defun__shen_4proc_1input_7 Obj         // shen.proc-input+
var __defun__shen_4elim_1def Obj             // shen.elim-def
var __defun__shen_4add_1macro Obj            // shen.add-macro
var __defun__shen_4packaged_2 Obj            // shen.packaged?
var __defun__external Obj                    // external
var __defun__internal Obj                    // internal
var __defun__shen_4package_1contents Obj     // shen.package-contents
var __defun__shen_4walk Obj                  // shen.walk
var __defun__compile Obj                     // compile
var __defun__fail_1if Obj                    // fail-if
var __defun___8s Obj                         // @s
var __defun__tc_2 Obj                        // tc?
var __defun__ps Obj                          // ps
var __defun__stinput Obj                     // stinput
var __defun__vector Obj                      // vector
var __defun__shen_4fillvector Obj            // shen.fillvector
var __defun__vector_2 Obj                    // vector?
var __defun__vector_1_6 Obj                  // vector->
var __defun___5_1vector Obj                  // <-vector
var __defun__shen_4posint_2 Obj              // shen.posint?
var __defun__limit Obj                       // limit
var __defun__symbol_2 Obj                    // symbol?
var __defun__shen_4analyse_1symbol_2 Obj     // shen.analyse-symbol?
var __defun__shen_4alpha_2 Obj               // shen.alpha?
var __defun__shen_4alphanums_2 Obj           // shen.alphanums?
var __defun__shen_4alphanum_2 Obj            // shen.alphanum?
var __defun__shen_4digit_2 Obj               // shen.digit?
var __defun__variable_2 Obj                  // variable?
var __defun__shen_4analyse_1variable_2 Obj   // shen.analyse-variable?
var __defun__shen_4uppercase_2 Obj           // shen.uppercase?
var __defun__gensym Obj                      // gensym
var __defun__concat Obj                      // concat
var __defun___8p Obj                         // @p
var __defun__fst Obj                         // fst
var __defun__snd Obj                         // snd
var __defun__tuple_2 Obj                     // tuple?
var __defun__append Obj                      // append
var __defun___8v Obj                         // @v
var __defun__shen_4_8v_1help Obj             // shen.@v-help
var __defun__shen_4copyfromvector Obj        // shen.copyfromvector
var __defun__hdv Obj                         // hdv
var __defun__tlv Obj                         // tlv
var __defun__shen_4tlv_1help Obj             // shen.tlv-help
var __defun__assoc Obj                       // assoc
var __defun__shen_4assoc_1set Obj            // shen.assoc-set
var __defun__shen_4assoc_1rm Obj             // shen.assoc-rm
var __defun__boolean_2 Obj                   // boolean?
var __defun__nl Obj                          // nl
var __defun__difference Obj                  // difference
var __defun__do Obj                          // do
var __defun__element_2 Obj                   // element?
var __defun__empty_2 Obj                     // empty?
var __defun__fix Obj                         // fix
var __defun__shen_4fix_1help Obj             // shen.fix-help
var __defun__put Obj                         // put
var __defun__unput Obj                       // unput
var __defun__get Obj                         // get
var __defun__hash Obj                        // hash
var __defun__shen_4mod Obj                   // shen.mod
var __defun__shen_4multiples Obj             // shen.multiples
var __defun__shen_4modh Obj                  // shen.modh
var __defun__sum Obj                         // sum
var __defun__head Obj                        // head
var __defun__tail Obj                        // tail
var __defun__hdstr Obj                       // hdstr
var __defun__intersection Obj                // intersection
var __defun__reverse Obj                     // reverse
var __defun__shen_4reverse__help Obj         // shen.reverse_help
var __defun__union Obj                       // union
var __defun__y_1or_1n_2 Obj                  // y-or-n?
var __defun__not Obj                         // not
var __defun__subst Obj                       // subst
var __defun__explode Obj                     // explode
var __defun__shen_4explode_1h Obj            // shen.explode-h
var __defun__cd Obj                          // cd
var __defun__shen_4for_1each Obj             // shen.for-each
var __defun__map Obj                         // map
var __defun__length Obj                      // length
var __defun__shen_4length_1h Obj             // shen.length-h
var __defun__occurrences Obj                 // occurrences
var __defun__nth Obj                         // nth
var __defun__integer_2 Obj                   // integer?
var __defun__shen_4abs Obj                   // shen.abs
var __defun__shen_4magless Obj               // shen.magless
var __defun__shen_4integer_1test_2 Obj       // shen.integer-test?
var __defun__mapcan Obj                      // mapcan
var __defun___a_a Obj                        // ==
var __defun__abort Obj                       // abort
var __defun__bound_2 Obj                     // bound?
var __defun__shen_4string_1_6bytes Obj       // shen.string->bytes
var __defun__maxinferences Obj               // maxinferences
var __defun__inferences Obj                  // inferences
var __defun__protect Obj                     // protect
var __defun__stoutput Obj                    // stoutput
var __defun__sterror Obj                     // sterror
var __defun__string_1_6symbol Obj            // string->symbol
var __defun__optimise Obj                    // optimise
var __defun__os Obj                          // os
var __defun__language Obj                    // language
var __defun__version Obj                     // version
var __defun__port Obj                        // port
var __defun__porters Obj                     // porters
var __defun__implementation Obj              // implementation
var __defun__release Obj                     // release
var __defun__package_2 Obj                   // package?
var __defun__function Obj                    // function
var __defun__shen_4lookup_1func Obj          // shen.lookup-func

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg4475 := MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
		__ctx.Return(reg4475)
		return
	}, 0))
	__defun__thaw = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1890 := __args[0]
		_ = V1890
		__ctx.TailApply(V1890)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "thaw", value: __defun__thaw})

	__defun__eval = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1892 := __args[0]
		_ = V1892
		reg4477 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			Y := __args[0]
			_ = Y
			__ctx.TailApply(__defun__macroexpand, Y)
			return
		}, 1)
		reg4479 := __e.Call(__defun__shen_4walk, reg4477, V1892)
		Macroexpand := reg4479
		_ = Macroexpand
		reg4480 := __e.Call(__defun__shen_4packaged_2, Macroexpand)
		if reg4480 == True {
			reg4481 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				Z := __args[0]
				_ = Z
				__ctx.TailApply(__defun__shen_4eval_1without_1macros, Z)
				return
			}, 1)
			reg4483 := __e.Call(__defun__shen_4package_1contents, Macroexpand)
			__ctx.TailApply(__defun__map, reg4481, reg4483)
			return
		} else {
			__ctx.TailApply(__defun__shen_4eval_1without_1macros, Macroexpand)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "eval", value: __defun__eval})

	__defun__shen_4eval_1without_1macros = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1894 := __args[0]
		_ = V1894
		reg4486 := __e.Call(__defun__shen_4proc_1input_7, V1894)
		reg4487 := __e.Call(__defun__shen_4elim_1def, reg4486)
		reg4488 := PrimEvalKL(__e, reg4487)
		__ctx.Return(reg4488)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.eval-without-macros", value: __defun__shen_4eval_1without_1macros})

	__defun__shen_4proc_1input_7 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1896 := __args[0]
		_ = V1896
		reg4489 := PrimIsPair(V1896)
		var reg4522 Obj
		if reg4489 == True {
			reg4490 := MakeSymbol("input+")
			reg4491 := PrimHead(V1896)
			reg4492 := PrimEqual(reg4490, reg4491)
			var reg4517 Obj
			if reg4492 == True {
				reg4493 := PrimTail(V1896)
				reg4494 := PrimIsPair(reg4493)
				var reg4512 Obj
				if reg4494 == True {
					reg4495 := PrimTail(V1896)
					reg4496 := PrimTail(reg4495)
					reg4497 := PrimIsPair(reg4496)
					var reg4507 Obj
					if reg4497 == True {
						reg4498 := Nil
						reg4499 := PrimTail(V1896)
						reg4500 := PrimTail(reg4499)
						reg4501 := PrimTail(reg4500)
						reg4502 := PrimEqual(reg4498, reg4501)
						var reg4505 Obj
						if reg4502 == True {
							reg4503 := True
							reg4505 = reg4503
						} else {
							reg4504 := False
							reg4505 = reg4504
						}
						reg4507 = reg4505
					} else {
						reg4506 := False
						reg4507 = reg4506
					}
					var reg4510 Obj
					if reg4507 == True {
						reg4508 := True
						reg4510 = reg4508
					} else {
						reg4509 := False
						reg4510 = reg4509
					}
					reg4512 = reg4510
				} else {
					reg4511 := False
					reg4512 = reg4511
				}
				var reg4515 Obj
				if reg4512 == True {
					reg4513 := True
					reg4515 = reg4513
				} else {
					reg4514 := False
					reg4515 = reg4514
				}
				reg4517 = reg4515
			} else {
				reg4516 := False
				reg4517 = reg4516
			}
			var reg4520 Obj
			if reg4517 == True {
				reg4518 := True
				reg4520 = reg4518
			} else {
				reg4519 := False
				reg4520 = reg4519
			}
			reg4522 = reg4520
		} else {
			reg4521 := False
			reg4522 = reg4521
		}
		if reg4522 == True {
			reg4523 := MakeSymbol("input+")
			reg4524 := PrimTail(V1896)
			reg4525 := PrimHead(reg4524)
			reg4526 := __e.Call(__defun__shen_4rcons__form, reg4525)
			reg4527 := PrimTail(V1896)
			reg4528 := PrimTail(reg4527)
			reg4529 := PrimCons(reg4526, reg4528)
			reg4530 := PrimCons(reg4523, reg4529)
			__ctx.Return(reg4530)
			return
		} else {
			reg4531 := PrimIsPair(V1896)
			var reg4564 Obj
			if reg4531 == True {
				reg4532 := MakeSymbol("shen.read+")
				reg4533 := PrimHead(V1896)
				reg4534 := PrimEqual(reg4532, reg4533)
				var reg4559 Obj
				if reg4534 == True {
					reg4535 := PrimTail(V1896)
					reg4536 := PrimIsPair(reg4535)
					var reg4554 Obj
					if reg4536 == True {
						reg4537 := PrimTail(V1896)
						reg4538 := PrimTail(reg4537)
						reg4539 := PrimIsPair(reg4538)
						var reg4549 Obj
						if reg4539 == True {
							reg4540 := Nil
							reg4541 := PrimTail(V1896)
							reg4542 := PrimTail(reg4541)
							reg4543 := PrimTail(reg4542)
							reg4544 := PrimEqual(reg4540, reg4543)
							var reg4547 Obj
							if reg4544 == True {
								reg4545 := True
								reg4547 = reg4545
							} else {
								reg4546 := False
								reg4547 = reg4546
							}
							reg4549 = reg4547
						} else {
							reg4548 := False
							reg4549 = reg4548
						}
						var reg4552 Obj
						if reg4549 == True {
							reg4550 := True
							reg4552 = reg4550
						} else {
							reg4551 := False
							reg4552 = reg4551
						}
						reg4554 = reg4552
					} else {
						reg4553 := False
						reg4554 = reg4553
					}
					var reg4557 Obj
					if reg4554 == True {
						reg4555 := True
						reg4557 = reg4555
					} else {
						reg4556 := False
						reg4557 = reg4556
					}
					reg4559 = reg4557
				} else {
					reg4558 := False
					reg4559 = reg4558
				}
				var reg4562 Obj
				if reg4559 == True {
					reg4560 := True
					reg4562 = reg4560
				} else {
					reg4561 := False
					reg4562 = reg4561
				}
				reg4564 = reg4562
			} else {
				reg4563 := False
				reg4564 = reg4563
			}
			if reg4564 == True {
				reg4565 := MakeSymbol("shen.read+")
				reg4566 := PrimTail(V1896)
				reg4567 := PrimHead(reg4566)
				reg4568 := __e.Call(__defun__shen_4rcons__form, reg4567)
				reg4569 := PrimTail(V1896)
				reg4570 := PrimTail(reg4569)
				reg4571 := PrimCons(reg4568, reg4570)
				reg4572 := PrimCons(reg4565, reg4571)
				__ctx.Return(reg4572)
				return
			} else {
				reg4573 := PrimIsPair(V1896)
				if reg4573 == True {
					reg4574 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
						Z := __args[0]
						_ = Z
						__ctx.TailApply(__defun__shen_4proc_1input_7, Z)
						return
					}, 1)
					__ctx.TailApply(__defun__map, reg4574, V1896)
					return
				} else {
					__ctx.Return(V1896)
					return
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.proc-input+", value: __defun__shen_4proc_1input_7})

	__defun__shen_4elim_1def = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1898 := __args[0]
		_ = V1898
		reg4577 := PrimIsPair(V1898)
		var reg4592 Obj
		if reg4577 == True {
			reg4578 := MakeSymbol("define")
			reg4579 := PrimHead(V1898)
			reg4580 := PrimEqual(reg4578, reg4579)
			var reg4587 Obj
			if reg4580 == True {
				reg4581 := PrimTail(V1898)
				reg4582 := PrimIsPair(reg4581)
				var reg4585 Obj
				if reg4582 == True {
					reg4583 := True
					reg4585 = reg4583
				} else {
					reg4584 := False
					reg4585 = reg4584
				}
				reg4587 = reg4585
			} else {
				reg4586 := False
				reg4587 = reg4586
			}
			var reg4590 Obj
			if reg4587 == True {
				reg4588 := True
				reg4590 = reg4588
			} else {
				reg4589 := False
				reg4590 = reg4589
			}
			reg4592 = reg4590
		} else {
			reg4591 := False
			reg4592 = reg4591
		}
		if reg4592 == True {
			reg4593 := PrimTail(V1898)
			reg4594 := PrimHead(reg4593)
			reg4595 := PrimTail(V1898)
			reg4596 := PrimTail(reg4595)
			__ctx.TailApply(__defun__shen_4shen_1_6kl, reg4594, reg4596)
			return
		} else {
			reg4598 := PrimIsPair(V1898)
			var reg4613 Obj
			if reg4598 == True {
				reg4599 := MakeSymbol("defmacro")
				reg4600 := PrimHead(V1898)
				reg4601 := PrimEqual(reg4599, reg4600)
				var reg4608 Obj
				if reg4601 == True {
					reg4602 := PrimTail(V1898)
					reg4603 := PrimIsPair(reg4602)
					var reg4606 Obj
					if reg4603 == True {
						reg4604 := True
						reg4606 = reg4604
					} else {
						reg4605 := False
						reg4606 = reg4605
					}
					reg4608 = reg4606
				} else {
					reg4607 := False
					reg4608 = reg4607
				}
				var reg4611 Obj
				if reg4608 == True {
					reg4609 := True
					reg4611 = reg4609
				} else {
					reg4610 := False
					reg4611 = reg4610
				}
				reg4613 = reg4611
			} else {
				reg4612 := False
				reg4613 = reg4612
			}
			if reg4613 == True {
				reg4614 := MakeSymbol("X")
				reg4615 := MakeSymbol("->")
				reg4616 := MakeSymbol("X")
				reg4617 := Nil
				reg4618 := PrimCons(reg4616, reg4617)
				reg4619 := PrimCons(reg4615, reg4618)
				reg4620 := PrimCons(reg4614, reg4619)
				Default := reg4620
				_ = Default
				reg4621 := MakeSymbol("define")
				reg4622 := PrimTail(V1898)
				reg4623 := PrimHead(reg4622)
				reg4624 := PrimTail(V1898)
				reg4625 := PrimTail(reg4624)
				reg4626 := __e.Call(__defun__append, reg4625, Default)
				reg4627 := PrimCons(reg4623, reg4626)
				reg4628 := PrimCons(reg4621, reg4627)
				reg4629 := __e.Call(__defun__shen_4elim_1def, reg4628)
				Def := reg4629
				_ = Def
				reg4630 := PrimTail(V1898)
				reg4631 := PrimHead(reg4630)
				reg4632 := __e.Call(__defun__shen_4add_1macro, reg4631)
				MacroAdd := reg4632
				_ = MacroAdd
				__ctx.Return(Def)
				return
			} else {
				reg4633 := PrimIsPair(V1898)
				var reg4648 Obj
				if reg4633 == True {
					reg4634 := MakeSymbol("defcc")
					reg4635 := PrimHead(V1898)
					reg4636 := PrimEqual(reg4634, reg4635)
					var reg4643 Obj
					if reg4636 == True {
						reg4637 := PrimTail(V1898)
						reg4638 := PrimIsPair(reg4637)
						var reg4641 Obj
						if reg4638 == True {
							reg4639 := True
							reg4641 = reg4639
						} else {
							reg4640 := False
							reg4641 = reg4640
						}
						reg4643 = reg4641
					} else {
						reg4642 := False
						reg4643 = reg4642
					}
					var reg4646 Obj
					if reg4643 == True {
						reg4644 := True
						reg4646 = reg4644
					} else {
						reg4645 := False
						reg4646 = reg4645
					}
					reg4648 = reg4646
				} else {
					reg4647 := False
					reg4648 = reg4647
				}
				if reg4648 == True {
					reg4649 := __e.Call(__defun__shen_4yacc, V1898)
					__ctx.TailApply(__defun__shen_4elim_1def, reg4649)
					return
				} else {
					reg4651 := PrimIsPair(V1898)
					if reg4651 == True {
						reg4652 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
							Z := __args[0]
							_ = Z
							__ctx.TailApply(__defun__shen_4elim_1def, Z)
							return
						}, 1)
						__ctx.TailApply(__defun__map, reg4652, V1898)
						return
					} else {
						__ctx.Return(V1898)
						return
					}
				}
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.elim-def", value: __defun__shen_4elim_1def})

	__defun__shen_4add_1macro = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1900 := __args[0]
		_ = V1900
		reg4655 := MakeSymbol("shen.*macroreg*")
		reg4656 := PrimValue(reg4655)
		MacroReg := reg4656
		_ = MacroReg
		reg4657 := MakeSymbol("shen.*macroreg*")
		reg4658 := MakeSymbol("shen.*macroreg*")
		reg4659 := PrimValue(reg4658)
		reg4660 := __e.Call(__defun__adjoin, V1900, reg4659)
		reg4661 := PrimSet(reg4657, reg4660)
		NewMacroReg := reg4661
		_ = NewMacroReg
		reg4662 := PrimEqual(MacroReg, NewMacroReg)
		if reg4662 == True {
			reg4663 := MakeSymbol("shen.skip")
			__ctx.Return(reg4663)
			return
		} else {
			reg4664 := MakeSymbol("*macros*")
			reg4665 := __e.Call(__defun__function, V1900)
			reg4666 := MakeSymbol("*macros*")
			reg4667 := PrimValue(reg4666)
			reg4668 := PrimCons(reg4665, reg4667)
			reg4669 := PrimSet(reg4664, reg4668)
			__ctx.Return(reg4669)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.add-macro", value: __defun__shen_4add_1macro})

	__defun__shen_4packaged_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1908 := __args[0]
		_ = V1908
		reg4670 := PrimIsPair(V1908)
		var reg4693 Obj
		if reg4670 == True {
			reg4671 := MakeSymbol("package")
			reg4672 := PrimHead(V1908)
			reg4673 := PrimEqual(reg4671, reg4672)
			var reg4688 Obj
			if reg4673 == True {
				reg4674 := PrimTail(V1908)
				reg4675 := PrimIsPair(reg4674)
				var reg4683 Obj
				if reg4675 == True {
					reg4676 := PrimTail(V1908)
					reg4677 := PrimTail(reg4676)
					reg4678 := PrimIsPair(reg4677)
					var reg4681 Obj
					if reg4678 == True {
						reg4679 := True
						reg4681 = reg4679
					} else {
						reg4680 := False
						reg4681 = reg4680
					}
					reg4683 = reg4681
				} else {
					reg4682 := False
					reg4683 = reg4682
				}
				var reg4686 Obj
				if reg4683 == True {
					reg4684 := True
					reg4686 = reg4684
				} else {
					reg4685 := False
					reg4686 = reg4685
				}
				reg4688 = reg4686
			} else {
				reg4687 := False
				reg4688 = reg4687
			}
			var reg4691 Obj
			if reg4688 == True {
				reg4689 := True
				reg4691 = reg4689
			} else {
				reg4690 := False
				reg4691 = reg4690
			}
			reg4693 = reg4691
		} else {
			reg4692 := False
			reg4693 = reg4692
		}
		if reg4693 == True {
			reg4694 := True
			__ctx.Return(reg4694)
			return
		} else {
			reg4695 := False
			__ctx.Return(reg4695)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.packaged?", value: __defun__shen_4packaged_2})

	__defun__external = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1910 := __args[0]
		_ = V1910
		reg4696 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg4697 := MakeSymbol("shen.external-symbols")
			reg4698 := MakeSymbol("*property-vector*")
			reg4699 := PrimValue(reg4698)
			__ctx.TailApply(__defun__get, V1910, reg4697, reg4699)
			return
		}, 0)
		reg4701 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg4702 := MakeString("package ")
			reg4703 := MakeString(" has not been used.\n")
			reg4704 := MakeSymbol("shen.a")
			reg4705 := __e.Call(__defun__shen_4app, V1910, reg4703, reg4704)
			reg4706 := PrimStringConcat(reg4702, reg4705)
			reg4707 := PrimSimpleError(reg4706)
			__ctx.Return(reg4707)
			return
		}, 1)
		reg4708 := __e.Try(reg4696).Catch(reg4701)
		__ctx.Return(reg4708)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "external", value: __defun__external})

	__defun__internal = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1912 := __args[0]
		_ = V1912
		reg4709 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg4710 := MakeSymbol("shen.internal-symbols")
			reg4711 := MakeSymbol("*property-vector*")
			reg4712 := PrimValue(reg4711)
			__ctx.TailApply(__defun__get, V1912, reg4710, reg4712)
			return
		}, 0)
		reg4714 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg4715 := MakeString("package ")
			reg4716 := MakeString(" has not been used.\n")
			reg4717 := MakeSymbol("shen.a")
			reg4718 := __e.Call(__defun__shen_4app, V1912, reg4716, reg4717)
			reg4719 := PrimStringConcat(reg4715, reg4718)
			reg4720 := PrimSimpleError(reg4719)
			__ctx.Return(reg4720)
			return
		}, 1)
		reg4721 := __e.Try(reg4709).Catch(reg4714)
		__ctx.Return(reg4721)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "internal", value: __defun__internal})

	__defun__shen_4package_1contents = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1916 := __args[0]
		_ = V1916
		reg4722 := PrimIsPair(V1916)
		var reg4754 Obj
		if reg4722 == True {
			reg4723 := MakeSymbol("package")
			reg4724 := PrimHead(V1916)
			reg4725 := PrimEqual(reg4723, reg4724)
			var reg4749 Obj
			if reg4725 == True {
				reg4726 := PrimTail(V1916)
				reg4727 := PrimIsPair(reg4726)
				var reg4744 Obj
				if reg4727 == True {
					reg4728 := MakeSymbol("null")
					reg4729 := PrimTail(V1916)
					reg4730 := PrimHead(reg4729)
					reg4731 := PrimEqual(reg4728, reg4730)
					var reg4739 Obj
					if reg4731 == True {
						reg4732 := PrimTail(V1916)
						reg4733 := PrimTail(reg4732)
						reg4734 := PrimIsPair(reg4733)
						var reg4737 Obj
						if reg4734 == True {
							reg4735 := True
							reg4737 = reg4735
						} else {
							reg4736 := False
							reg4737 = reg4736
						}
						reg4739 = reg4737
					} else {
						reg4738 := False
						reg4739 = reg4738
					}
					var reg4742 Obj
					if reg4739 == True {
						reg4740 := True
						reg4742 = reg4740
					} else {
						reg4741 := False
						reg4742 = reg4741
					}
					reg4744 = reg4742
				} else {
					reg4743 := False
					reg4744 = reg4743
				}
				var reg4747 Obj
				if reg4744 == True {
					reg4745 := True
					reg4747 = reg4745
				} else {
					reg4746 := False
					reg4747 = reg4746
				}
				reg4749 = reg4747
			} else {
				reg4748 := False
				reg4749 = reg4748
			}
			var reg4752 Obj
			if reg4749 == True {
				reg4750 := True
				reg4752 = reg4750
			} else {
				reg4751 := False
				reg4752 = reg4751
			}
			reg4754 = reg4752
		} else {
			reg4753 := False
			reg4754 = reg4753
		}
		if reg4754 == True {
			reg4755 := PrimTail(V1916)
			reg4756 := PrimTail(reg4755)
			reg4757 := PrimTail(reg4756)
			__ctx.Return(reg4757)
			return
		} else {
			reg4758 := PrimIsPair(V1916)
			var reg4781 Obj
			if reg4758 == True {
				reg4759 := MakeSymbol("package")
				reg4760 := PrimHead(V1916)
				reg4761 := PrimEqual(reg4759, reg4760)
				var reg4776 Obj
				if reg4761 == True {
					reg4762 := PrimTail(V1916)
					reg4763 := PrimIsPair(reg4762)
					var reg4771 Obj
					if reg4763 == True {
						reg4764 := PrimTail(V1916)
						reg4765 := PrimTail(reg4764)
						reg4766 := PrimIsPair(reg4765)
						var reg4769 Obj
						if reg4766 == True {
							reg4767 := True
							reg4769 = reg4767
						} else {
							reg4768 := False
							reg4769 = reg4768
						}
						reg4771 = reg4769
					} else {
						reg4770 := False
						reg4771 = reg4770
					}
					var reg4774 Obj
					if reg4771 == True {
						reg4772 := True
						reg4774 = reg4772
					} else {
						reg4773 := False
						reg4774 = reg4773
					}
					reg4776 = reg4774
				} else {
					reg4775 := False
					reg4776 = reg4775
				}
				var reg4779 Obj
				if reg4776 == True {
					reg4777 := True
					reg4779 = reg4777
				} else {
					reg4778 := False
					reg4779 = reg4778
				}
				reg4781 = reg4779
			} else {
				reg4780 := False
				reg4781 = reg4780
			}
			if reg4781 == True {
				reg4782 := PrimTail(V1916)
				reg4783 := PrimHead(reg4782)
				reg4784 := PrimStr(reg4783)
				reg4785 := MakeString(".")
				reg4786 := PrimStringConcat(reg4784, reg4785)
				reg4787 := PrimIntern(reg4786)
				PackageNameDot := reg4787
				_ = PackageNameDot
				reg4788 := __e.Call(__defun__explode, PackageNameDot)
				ExpPackageNameDot := reg4788
				_ = ExpPackageNameDot
				reg4789 := PrimTail(V1916)
				reg4790 := PrimHead(reg4789)
				reg4791 := PrimTail(V1916)
				reg4792 := PrimTail(reg4791)
				reg4793 := PrimHead(reg4792)
				reg4794 := PrimTail(V1916)
				reg4795 := PrimTail(reg4794)
				reg4796 := PrimTail(reg4795)
				__ctx.TailApply(__defun__shen_4packageh, reg4790, reg4793, reg4796, ExpPackageNameDot)
				return
			} else {
				reg4798 := MakeSymbol("shen.package-contents")
				__ctx.TailApply(__defun__shen_4f__error, reg4798)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.package-contents", value: __defun__shen_4package_1contents})

	__defun__shen_4walk = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1919 := __args[0]
		_ = V1919
		V1920 := __args[1]
		_ = V1920
		reg4800 := PrimIsPair(V1920)
		if reg4800 == True {
			reg4801 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				Z := __args[0]
				_ = Z
				__ctx.TailApply(__defun__shen_4walk, V1919, Z)
				return
			}, 1)
			reg4803 := __e.Call(__defun__map, reg4801, V1920)
			__ctx.TailApply(V1919, reg4803)
			return
		} else {
			__ctx.TailApply(V1919, V1920)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.walk", value: __defun__shen_4walk})

	__defun__compile = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1924 := __args[0]
		_ = V1924
		V1925 := __args[1]
		_ = V1925
		V1926 := __args[2]
		_ = V1926
		reg4806 := Nil
		reg4807 := Nil
		reg4808 := PrimCons(reg4806, reg4807)
		reg4809 := PrimCons(V1925, reg4808)
		reg4810 := __e.Call(V1924, reg4809)
		O := reg4810
		_ = O
		reg4811 := __e.Call(__defun__fail)
		reg4812 := PrimEqual(reg4811, O)
		var reg4820 Obj
		if reg4812 == True {
			reg4813 := True
			reg4820 = reg4813
		} else {
			reg4814 := PrimHead(O)
			reg4815 := __e.Call(__defun__empty_2, reg4814)
			reg4816 := PrimNot(reg4815)
			var reg4819 Obj
			if reg4816 == True {
				reg4817 := True
				reg4819 = reg4817
			} else {
				reg4818 := False
				reg4819 = reg4818
			}
			reg4820 = reg4819
		}
		if reg4820 == True {
			__ctx.TailApply(V1926, O)
			return
		} else {
			__ctx.TailApply(__defun__shen_4hdtl, O)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "compile", value: __defun__compile})

	__defun__fail_1if = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1929 := __args[0]
		_ = V1929
		V1930 := __args[1]
		_ = V1930
		reg4823 := __e.Call(V1929, V1930)
		if reg4823 == True {
			__ctx.TailApply(__defun__fail)
			return
		} else {
			__ctx.Return(V1930)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "fail-if", value: __defun__fail_1if})

	__defun___8s = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1933 := __args[0]
		_ = V1933
		V1934 := __args[1]
		_ = V1934
		reg4825 := PrimStringConcat(V1933, V1934)
		__ctx.Return(reg4825)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "@s", value: __defun___8s})

	__defun__tc_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg4826 := MakeSymbol("shen.*tc*")
		reg4827 := PrimValue(reg4826)
		__ctx.Return(reg4827)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "tc?", value: __defun__tc_2})

	__defun__ps = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1936 := __args[0]
		_ = V1936
		reg4828 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg4829 := MakeSymbol("shen.source")
			reg4830 := MakeSymbol("*property-vector*")
			reg4831 := PrimValue(reg4830)
			__ctx.TailApply(__defun__get, V1936, reg4829, reg4831)
			return
		}, 0)
		reg4833 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg4834 := MakeString(" not found.\n")
			reg4835 := MakeSymbol("shen.a")
			reg4836 := __e.Call(__defun__shen_4app, V1936, reg4834, reg4835)
			reg4837 := PrimSimpleError(reg4836)
			__ctx.Return(reg4837)
			return
		}, 1)
		reg4838 := __e.Try(reg4828).Catch(reg4833)
		__ctx.Return(reg4838)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "ps", value: __defun__ps})

	__defun__stinput = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg4839 := MakeSymbol("*stinput*")
		reg4840 := PrimValue(reg4839)
		__ctx.Return(reg4840)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "stinput", value: __defun__stinput})

	__defun__vector = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1938 := __args[0]
		_ = V1938
		reg4841 := MakeNumber(1)
		reg4842 := PrimNumberAdd(V1938, reg4841)
		reg4843 := PrimAbsvector(reg4842)
		Vector := reg4843
		_ = Vector
		reg4844 := MakeNumber(0)
		reg4845 := PrimVectorSet(Vector, reg4844, V1938)
		ZeroStamp := reg4845
		_ = ZeroStamp
		reg4846 := MakeNumber(0)
		reg4847 := PrimEqual(V1938, reg4846)
		var reg4851 Obj
		if reg4847 == True {
			reg4851 = ZeroStamp
		} else {
			reg4848 := MakeNumber(1)
			reg4849 := __e.Call(__defun__fail)
			reg4850 := __e.Call(__defun__shen_4fillvector, ZeroStamp, reg4848, V1938, reg4849)
			reg4851 = reg4850
		}
		Standard := reg4851
		_ = Standard
		__ctx.Return(Standard)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "vector", value: __defun__vector})

	__defun__shen_4fillvector = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1944 := __args[0]
		_ = V1944
		V1945 := __args[1]
		_ = V1945
		V1946 := __args[2]
		_ = V1946
		V1947 := __args[3]
		_ = V1947
		reg4852 := PrimEqual(V1946, V1945)
		if reg4852 == True {
			reg4853 := PrimVectorSet(V1944, V1946, V1947)
			__ctx.Return(reg4853)
			return
		} else {
			reg4854 := PrimVectorSet(V1944, V1945, V1947)
			reg4855 := MakeNumber(1)
			reg4856 := PrimNumberAdd(reg4855, V1945)
			__ctx.TailApply(__defun__shen_4fillvector, reg4854, reg4856, V1946, V1947)
			return
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.fillvector", value: __defun__shen_4fillvector})

	__defun__vector_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1949 := __args[0]
		_ = V1949
		reg4858 := PrimIsVector(V1949)
		if reg4858 == True {
			reg4859 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				reg4860 := MakeNumber(0)
				reg4861 := PrimVectorGet(V1949, reg4860)
				__ctx.Return(reg4861)
				return
			}, 0)
			reg4862 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				E := __args[0]
				_ = E
				reg4863 := MakeNumber(-1)
				__ctx.Return(reg4863)
				return
			}, 1)
			reg4864 := __e.Try(reg4859).Catch(reg4862)
			X := reg4864
			_ = X
			reg4865 := PrimIsNumber(X)
			var reg4872 Obj
			if reg4865 == True {
				reg4866 := MakeNumber(0)
				reg4867 := PrimGreatEqual(X, reg4866)
				var reg4870 Obj
				if reg4867 == True {
					reg4868 := True
					reg4870 = reg4868
				} else {
					reg4869 := False
					reg4870 = reg4869
				}
				reg4872 = reg4870
			} else {
				reg4871 := False
				reg4872 = reg4871
			}
			if reg4872 == True {
				reg4873 := True
				__ctx.Return(reg4873)
				return
			} else {
				reg4874 := False
				__ctx.Return(reg4874)
				return
			}
		} else {
			reg4875 := False
			__ctx.Return(reg4875)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "vector?", value: __defun__vector_2})

	__defun__vector_1_6 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1953 := __args[0]
		_ = V1953
		V1954 := __args[1]
		_ = V1954
		V1955 := __args[2]
		_ = V1955
		reg4876 := MakeNumber(0)
		reg4877 := PrimEqual(V1954, reg4876)
		if reg4877 == True {
			reg4878 := MakeString("cannot access 0th element of a vector\n")
			reg4879 := PrimSimpleError(reg4878)
			__ctx.Return(reg4879)
			return
		} else {
			reg4880 := PrimVectorSet(V1953, V1954, V1955)
			__ctx.Return(reg4880)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "vector->", value: __defun__vector_1_6})

	__defun___5_1vector = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1958 := __args[0]
		_ = V1958
		V1959 := __args[1]
		_ = V1959
		reg4881 := MakeNumber(0)
		reg4882 := PrimEqual(V1959, reg4881)
		if reg4882 == True {
			reg4883 := MakeString("cannot access 0th element of a vector\n")
			reg4884 := PrimSimpleError(reg4883)
			__ctx.Return(reg4884)
			return
		} else {
			reg4885 := PrimVectorGet(V1958, V1959)
			VectorElement := reg4885
			_ = VectorElement
			reg4886 := __e.Call(__defun__fail)
			reg4887 := PrimEqual(VectorElement, reg4886)
			if reg4887 == True {
				reg4888 := MakeString("vector element not found\n")
				reg4889 := PrimSimpleError(reg4888)
				__ctx.Return(reg4889)
				return
			} else {
				__ctx.Return(VectorElement)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "<-vector", value: __defun___5_1vector})

	__defun__shen_4posint_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1961 := __args[0]
		_ = V1961
		reg4890 := PrimIsInteger(V1961)
		if reg4890 == True {
			reg4891 := MakeNumber(0)
			reg4892 := PrimGreatEqual(V1961, reg4891)
			if reg4892 == True {
				reg4893 := True
				__ctx.Return(reg4893)
				return
			} else {
				reg4894 := False
				__ctx.Return(reg4894)
				return
			}
		} else {
			reg4895 := False
			__ctx.Return(reg4895)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.posint?", value: __defun__shen_4posint_2})

	__defun__limit = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1963 := __args[0]
		_ = V1963
		reg4896 := MakeNumber(0)
		reg4897 := PrimVectorGet(V1963, reg4896)
		__ctx.Return(reg4897)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "limit", value: __defun__limit})

	__defun__symbol_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1965 := __args[0]
		_ = V1965
		reg4898 := __e.Call(__defun__boolean_2, V1965)
		var reg4910 Obj
		if reg4898 == True {
			reg4899 := True
			reg4910 = reg4899
		} else {
			reg4900 := PrimIsNumber(V1965)
			var reg4906 Obj
			if reg4900 == True {
				reg4901 := True
				reg4906 = reg4901
			} else {
				reg4902 := PrimIsString(V1965)
				var reg4905 Obj
				if reg4902 == True {
					reg4903 := True
					reg4905 = reg4903
				} else {
					reg4904 := False
					reg4905 = reg4904
				}
				reg4906 = reg4905
			}
			var reg4909 Obj
			if reg4906 == True {
				reg4907 := True
				reg4909 = reg4907
			} else {
				reg4908 := False
				reg4909 = reg4908
			}
			reg4910 = reg4909
		}
		if reg4910 == True {
			reg4911 := False
			__ctx.Return(reg4911)
			return
		} else {
			reg4912 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				reg4913 := PrimStr(V1965)
				String := reg4913
				_ = String
				__ctx.TailApply(__defun__shen_4analyse_1symbol_2, String)
				return
			}, 0)
			reg4915 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				E := __args[0]
				_ = E
				reg4916 := False
				__ctx.Return(reg4916)
				return
			}, 1)
			reg4917 := __e.Try(reg4912).Catch(reg4915)
			__ctx.Return(reg4917)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "symbol?", value: __defun__symbol_2})

	__defun__shen_4analyse_1symbol_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1967 := __args[0]
		_ = V1967
		reg4918 := MakeString("")
		reg4919 := PrimEqual(reg4918, V1967)
		if reg4919 == True {
			reg4920 := False
			__ctx.Return(reg4920)
			return
		} else {
			reg4921 := __e.Call(__defun__shen_4_7string_2, V1967)
			if reg4921 == True {
				reg4922 := MakeNumber(0)
				reg4923 := PrimPos(V1967, reg4922)
				reg4924 := __e.Call(__defun__shen_4alpha_2, reg4923)
				if reg4924 == True {
					reg4925 := PrimTailString(V1967)
					reg4926 := __e.Call(__defun__shen_4alphanums_2, reg4925)
					if reg4926 == True {
						reg4927 := True
						__ctx.Return(reg4927)
						return
					} else {
						reg4928 := False
						__ctx.Return(reg4928)
						return
					}
				} else {
					reg4929 := False
					__ctx.Return(reg4929)
					return
				}
			} else {
				reg4930 := MakeSymbol("shen.analyse-symbol?")
				__ctx.TailApply(__defun__shen_4f__error, reg4930)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.analyse-symbol?", value: __defun__shen_4analyse_1symbol_2})

	__defun__shen_4alpha_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1969 := __args[0]
		_ = V1969
		reg4932 := MakeString("A")
		reg4933 := MakeString("B")
		reg4934 := MakeString("C")
		reg4935 := MakeString("D")
		reg4936 := MakeString("E")
		reg4937 := MakeString("F")
		reg4938 := MakeString("G")
		reg4939 := MakeString("H")
		reg4940 := MakeString("I")
		reg4941 := MakeString("J")
		reg4942 := MakeString("K")
		reg4943 := MakeString("L")
		reg4944 := MakeString("M")
		reg4945 := MakeString("N")
		reg4946 := MakeString("O")
		reg4947 := MakeString("P")
		reg4948 := MakeString("Q")
		reg4949 := MakeString("R")
		reg4950 := MakeString("S")
		reg4951 := MakeString("T")
		reg4952 := MakeString("U")
		reg4953 := MakeString("V")
		reg4954 := MakeString("W")
		reg4955 := MakeString("X")
		reg4956 := MakeString("Y")
		reg4957 := MakeString("Z")
		reg4958 := MakeString("a")
		reg4959 := MakeString("b")
		reg4960 := MakeString("c")
		reg4961 := MakeString("d")
		reg4962 := MakeString("e")
		reg4963 := MakeString("f")
		reg4964 := MakeString("g")
		reg4965 := MakeString("h")
		reg4966 := MakeString("i")
		reg4967 := MakeString("j")
		reg4968 := MakeString("k")
		reg4969 := MakeString("l")
		reg4970 := MakeString("m")
		reg4971 := MakeString("n")
		reg4972 := MakeString("o")
		reg4973 := MakeString("p")
		reg4974 := MakeString("q")
		reg4975 := MakeString("r")
		reg4976 := MakeString("s")
		reg4977 := MakeString("t")
		reg4978 := MakeString("u")
		reg4979 := MakeString("v")
		reg4980 := MakeString("w")
		reg4981 := MakeString("x")
		reg4982 := MakeString("y")
		reg4983 := MakeString("z")
		reg4984 := MakeString("=")
		reg4985 := MakeString("*")
		reg4986 := MakeString("/")
		reg4987 := MakeString("+")
		reg4988 := MakeString("-")
		reg4989 := MakeString("_")
		reg4990 := MakeString("?")
		reg4991 := MakeString("$")
		reg4992 := MakeString("!")
		reg4993 := MakeString("@")
		reg4994 := MakeString("~")
		reg4995 := MakeString(">")
		reg4996 := MakeString("<")
		reg4997 := MakeString("&")
		reg4998 := MakeString("%")
		reg4999 := MakeString("{")
		reg5000 := MakeString("}")
		reg5001 := MakeString(":")
		reg5002 := MakeString(";")
		reg5003 := MakeString("`")
		reg5004 := MakeString("#")
		reg5005 := MakeString("'")
		reg5006 := MakeString(".")
		reg5007 := Nil
		reg5008 := PrimCons(reg5006, reg5007)
		reg5009 := PrimCons(reg5005, reg5008)
		reg5010 := PrimCons(reg5004, reg5009)
		reg5011 := PrimCons(reg5003, reg5010)
		reg5012 := PrimCons(reg5002, reg5011)
		reg5013 := PrimCons(reg5001, reg5012)
		reg5014 := PrimCons(reg5000, reg5013)
		reg5015 := PrimCons(reg4999, reg5014)
		reg5016 := PrimCons(reg4998, reg5015)
		reg5017 := PrimCons(reg4997, reg5016)
		reg5018 := PrimCons(reg4996, reg5017)
		reg5019 := PrimCons(reg4995, reg5018)
		reg5020 := PrimCons(reg4994, reg5019)
		reg5021 := PrimCons(reg4993, reg5020)
		reg5022 := PrimCons(reg4992, reg5021)
		reg5023 := PrimCons(reg4991, reg5022)
		reg5024 := PrimCons(reg4990, reg5023)
		reg5025 := PrimCons(reg4989, reg5024)
		reg5026 := PrimCons(reg4988, reg5025)
		reg5027 := PrimCons(reg4987, reg5026)
		reg5028 := PrimCons(reg4986, reg5027)
		reg5029 := PrimCons(reg4985, reg5028)
		reg5030 := PrimCons(reg4984, reg5029)
		reg5031 := PrimCons(reg4983, reg5030)
		reg5032 := PrimCons(reg4982, reg5031)
		reg5033 := PrimCons(reg4981, reg5032)
		reg5034 := PrimCons(reg4980, reg5033)
		reg5035 := PrimCons(reg4979, reg5034)
		reg5036 := PrimCons(reg4978, reg5035)
		reg5037 := PrimCons(reg4977, reg5036)
		reg5038 := PrimCons(reg4976, reg5037)
		reg5039 := PrimCons(reg4975, reg5038)
		reg5040 := PrimCons(reg4974, reg5039)
		reg5041 := PrimCons(reg4973, reg5040)
		reg5042 := PrimCons(reg4972, reg5041)
		reg5043 := PrimCons(reg4971, reg5042)
		reg5044 := PrimCons(reg4970, reg5043)
		reg5045 := PrimCons(reg4969, reg5044)
		reg5046 := PrimCons(reg4968, reg5045)
		reg5047 := PrimCons(reg4967, reg5046)
		reg5048 := PrimCons(reg4966, reg5047)
		reg5049 := PrimCons(reg4965, reg5048)
		reg5050 := PrimCons(reg4964, reg5049)
		reg5051 := PrimCons(reg4963, reg5050)
		reg5052 := PrimCons(reg4962, reg5051)
		reg5053 := PrimCons(reg4961, reg5052)
		reg5054 := PrimCons(reg4960, reg5053)
		reg5055 := PrimCons(reg4959, reg5054)
		reg5056 := PrimCons(reg4958, reg5055)
		reg5057 := PrimCons(reg4957, reg5056)
		reg5058 := PrimCons(reg4956, reg5057)
		reg5059 := PrimCons(reg4955, reg5058)
		reg5060 := PrimCons(reg4954, reg5059)
		reg5061 := PrimCons(reg4953, reg5060)
		reg5062 := PrimCons(reg4952, reg5061)
		reg5063 := PrimCons(reg4951, reg5062)
		reg5064 := PrimCons(reg4950, reg5063)
		reg5065 := PrimCons(reg4949, reg5064)
		reg5066 := PrimCons(reg4948, reg5065)
		reg5067 := PrimCons(reg4947, reg5066)
		reg5068 := PrimCons(reg4946, reg5067)
		reg5069 := PrimCons(reg4945, reg5068)
		reg5070 := PrimCons(reg4944, reg5069)
		reg5071 := PrimCons(reg4943, reg5070)
		reg5072 := PrimCons(reg4942, reg5071)
		reg5073 := PrimCons(reg4941, reg5072)
		reg5074 := PrimCons(reg4940, reg5073)
		reg5075 := PrimCons(reg4939, reg5074)
		reg5076 := PrimCons(reg4938, reg5075)
		reg5077 := PrimCons(reg4937, reg5076)
		reg5078 := PrimCons(reg4936, reg5077)
		reg5079 := PrimCons(reg4935, reg5078)
		reg5080 := PrimCons(reg4934, reg5079)
		reg5081 := PrimCons(reg4933, reg5080)
		reg5082 := PrimCons(reg4932, reg5081)
		__ctx.TailApply(__defun__element_2, V1969, reg5082)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.alpha?", value: __defun__shen_4alpha_2})

	__defun__shen_4alphanums_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1971 := __args[0]
		_ = V1971
		reg5084 := MakeString("")
		reg5085 := PrimEqual(reg5084, V1971)
		if reg5085 == True {
			reg5086 := True
			__ctx.Return(reg5086)
			return
		} else {
			reg5087 := __e.Call(__defun__shen_4_7string_2, V1971)
			if reg5087 == True {
				reg5088 := MakeNumber(0)
				reg5089 := PrimPos(V1971, reg5088)
				reg5090 := __e.Call(__defun__shen_4alphanum_2, reg5089)
				if reg5090 == True {
					reg5091 := PrimTailString(V1971)
					reg5092 := __e.Call(__defun__shen_4alphanums_2, reg5091)
					if reg5092 == True {
						reg5093 := True
						__ctx.Return(reg5093)
						return
					} else {
						reg5094 := False
						__ctx.Return(reg5094)
						return
					}
				} else {
					reg5095 := False
					__ctx.Return(reg5095)
					return
				}
			} else {
				reg5096 := MakeSymbol("shen.alphanums?")
				__ctx.TailApply(__defun__shen_4f__error, reg5096)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.alphanums?", value: __defun__shen_4alphanums_2})

	__defun__shen_4alphanum_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1973 := __args[0]
		_ = V1973
		reg5098 := __e.Call(__defun__shen_4alpha_2, V1973)
		if reg5098 == True {
			reg5099 := True
			__ctx.Return(reg5099)
			return
		} else {
			reg5100 := __e.Call(__defun__shen_4digit_2, V1973)
			if reg5100 == True {
				reg5101 := True
				__ctx.Return(reg5101)
				return
			} else {
				reg5102 := False
				__ctx.Return(reg5102)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.alphanum?", value: __defun__shen_4alphanum_2})

	__defun__shen_4digit_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1975 := __args[0]
		_ = V1975
		reg5103 := MakeString("1")
		reg5104 := MakeString("2")
		reg5105 := MakeString("3")
		reg5106 := MakeString("4")
		reg5107 := MakeString("5")
		reg5108 := MakeString("6")
		reg5109 := MakeString("7")
		reg5110 := MakeString("8")
		reg5111 := MakeString("9")
		reg5112 := MakeString("0")
		reg5113 := Nil
		reg5114 := PrimCons(reg5112, reg5113)
		reg5115 := PrimCons(reg5111, reg5114)
		reg5116 := PrimCons(reg5110, reg5115)
		reg5117 := PrimCons(reg5109, reg5116)
		reg5118 := PrimCons(reg5108, reg5117)
		reg5119 := PrimCons(reg5107, reg5118)
		reg5120 := PrimCons(reg5106, reg5119)
		reg5121 := PrimCons(reg5105, reg5120)
		reg5122 := PrimCons(reg5104, reg5121)
		reg5123 := PrimCons(reg5103, reg5122)
		__ctx.TailApply(__defun__element_2, V1975, reg5123)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.digit?", value: __defun__shen_4digit_2})

	__defun__variable_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1977 := __args[0]
		_ = V1977
		reg5125 := __e.Call(__defun__boolean_2, V1977)
		var reg5137 Obj
		if reg5125 == True {
			reg5126 := True
			reg5137 = reg5126
		} else {
			reg5127 := PrimIsNumber(V1977)
			var reg5133 Obj
			if reg5127 == True {
				reg5128 := True
				reg5133 = reg5128
			} else {
				reg5129 := PrimIsString(V1977)
				var reg5132 Obj
				if reg5129 == True {
					reg5130 := True
					reg5132 = reg5130
				} else {
					reg5131 := False
					reg5132 = reg5131
				}
				reg5133 = reg5132
			}
			var reg5136 Obj
			if reg5133 == True {
				reg5134 := True
				reg5136 = reg5134
			} else {
				reg5135 := False
				reg5136 = reg5135
			}
			reg5137 = reg5136
		}
		if reg5137 == True {
			reg5138 := False
			__ctx.Return(reg5138)
			return
		} else {
			reg5139 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				reg5140 := PrimStr(V1977)
				String := reg5140
				_ = String
				__ctx.TailApply(__defun__shen_4analyse_1variable_2, String)
				return
			}, 0)
			reg5142 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				E := __args[0]
				_ = E
				reg5143 := False
				__ctx.Return(reg5143)
				return
			}, 1)
			reg5144 := __e.Try(reg5139).Catch(reg5142)
			__ctx.Return(reg5144)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "variable?", value: __defun__variable_2})

	__defun__shen_4analyse_1variable_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1979 := __args[0]
		_ = V1979
		reg5145 := __e.Call(__defun__shen_4_7string_2, V1979)
		if reg5145 == True {
			reg5146 := MakeNumber(0)
			reg5147 := PrimPos(V1979, reg5146)
			reg5148 := __e.Call(__defun__shen_4uppercase_2, reg5147)
			if reg5148 == True {
				reg5149 := PrimTailString(V1979)
				reg5150 := __e.Call(__defun__shen_4alphanums_2, reg5149)
				if reg5150 == True {
					reg5151 := True
					__ctx.Return(reg5151)
					return
				} else {
					reg5152 := False
					__ctx.Return(reg5152)
					return
				}
			} else {
				reg5153 := False
				__ctx.Return(reg5153)
				return
			}
		} else {
			reg5154 := MakeSymbol("shen.analyse-variable?")
			__ctx.TailApply(__defun__shen_4f__error, reg5154)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.analyse-variable?", value: __defun__shen_4analyse_1variable_2})

	__defun__shen_4uppercase_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1981 := __args[0]
		_ = V1981
		reg5156 := MakeString("A")
		reg5157 := MakeString("B")
		reg5158 := MakeString("C")
		reg5159 := MakeString("D")
		reg5160 := MakeString("E")
		reg5161 := MakeString("F")
		reg5162 := MakeString("G")
		reg5163 := MakeString("H")
		reg5164 := MakeString("I")
		reg5165 := MakeString("J")
		reg5166 := MakeString("K")
		reg5167 := MakeString("L")
		reg5168 := MakeString("M")
		reg5169 := MakeString("N")
		reg5170 := MakeString("O")
		reg5171 := MakeString("P")
		reg5172 := MakeString("Q")
		reg5173 := MakeString("R")
		reg5174 := MakeString("S")
		reg5175 := MakeString("T")
		reg5176 := MakeString("U")
		reg5177 := MakeString("V")
		reg5178 := MakeString("W")
		reg5179 := MakeString("X")
		reg5180 := MakeString("Y")
		reg5181 := MakeString("Z")
		reg5182 := Nil
		reg5183 := PrimCons(reg5181, reg5182)
		reg5184 := PrimCons(reg5180, reg5183)
		reg5185 := PrimCons(reg5179, reg5184)
		reg5186 := PrimCons(reg5178, reg5185)
		reg5187 := PrimCons(reg5177, reg5186)
		reg5188 := PrimCons(reg5176, reg5187)
		reg5189 := PrimCons(reg5175, reg5188)
		reg5190 := PrimCons(reg5174, reg5189)
		reg5191 := PrimCons(reg5173, reg5190)
		reg5192 := PrimCons(reg5172, reg5191)
		reg5193 := PrimCons(reg5171, reg5192)
		reg5194 := PrimCons(reg5170, reg5193)
		reg5195 := PrimCons(reg5169, reg5194)
		reg5196 := PrimCons(reg5168, reg5195)
		reg5197 := PrimCons(reg5167, reg5196)
		reg5198 := PrimCons(reg5166, reg5197)
		reg5199 := PrimCons(reg5165, reg5198)
		reg5200 := PrimCons(reg5164, reg5199)
		reg5201 := PrimCons(reg5163, reg5200)
		reg5202 := PrimCons(reg5162, reg5201)
		reg5203 := PrimCons(reg5161, reg5202)
		reg5204 := PrimCons(reg5160, reg5203)
		reg5205 := PrimCons(reg5159, reg5204)
		reg5206 := PrimCons(reg5158, reg5205)
		reg5207 := PrimCons(reg5157, reg5206)
		reg5208 := PrimCons(reg5156, reg5207)
		__ctx.TailApply(__defun__element_2, V1981, reg5208)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.uppercase?", value: __defun__shen_4uppercase_2})

	__defun__gensym = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1983 := __args[0]
		_ = V1983
		reg5210 := MakeSymbol("shen.*gensym*")
		reg5211 := MakeNumber(1)
		reg5212 := MakeSymbol("shen.*gensym*")
		reg5213 := PrimValue(reg5212)
		reg5214 := PrimNumberAdd(reg5211, reg5213)
		reg5215 := PrimSet(reg5210, reg5214)
		__ctx.TailApply(__defun__concat, V1983, reg5215)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "gensym", value: __defun__gensym})

	__defun__concat = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1986 := __args[0]
		_ = V1986
		V1987 := __args[1]
		_ = V1987
		reg5217 := PrimStr(V1986)
		reg5218 := PrimStr(V1987)
		reg5219 := PrimStringConcat(reg5217, reg5218)
		reg5220 := PrimIntern(reg5219)
		__ctx.Return(reg5220)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "concat", value: __defun__concat})

	__defun___8p = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1990 := __args[0]
		_ = V1990
		V1991 := __args[1]
		_ = V1991
		reg5221 := MakeNumber(3)
		reg5222 := PrimAbsvector(reg5221)
		Vector := reg5222
		_ = Vector
		reg5223 := MakeNumber(0)
		reg5224 := MakeSymbol("shen.tuple")
		reg5225 := PrimVectorSet(Vector, reg5223, reg5224)
		Tag := reg5225
		_ = Tag
		reg5226 := MakeNumber(1)
		reg5227 := PrimVectorSet(Vector, reg5226, V1990)
		Fst := reg5227
		_ = Fst
		reg5228 := MakeNumber(2)
		reg5229 := PrimVectorSet(Vector, reg5228, V1991)
		Snd := reg5229
		_ = Snd
		__ctx.Return(Vector)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "@p", value: __defun___8p})

	__defun__fst = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1993 := __args[0]
		_ = V1993
		reg5230 := MakeNumber(1)
		reg5231 := PrimVectorGet(V1993, reg5230)
		__ctx.Return(reg5231)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "fst", value: __defun__fst})

	__defun__snd = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1995 := __args[0]
		_ = V1995
		reg5232 := MakeNumber(2)
		reg5233 := PrimVectorGet(V1995, reg5232)
		__ctx.Return(reg5233)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "snd", value: __defun__snd})

	__defun__tuple_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V1997 := __args[0]
		_ = V1997
		reg5234 := PrimIsVector(V1997)
		if reg5234 == True {
			reg5235 := MakeSymbol("shen.tuple")
			reg5236 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				reg5237 := MakeNumber(0)
				reg5238 := PrimVectorGet(V1997, reg5237)
				__ctx.Return(reg5238)
				return
			}, 0)
			reg5239 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				E := __args[0]
				_ = E
				reg5240 := MakeSymbol("shen.not-tuple")
				__ctx.Return(reg5240)
				return
			}, 1)
			reg5241 := __e.Try(reg5236).Catch(reg5239)
			reg5242 := PrimEqual(reg5235, reg5241)
			if reg5242 == True {
				reg5243 := True
				__ctx.Return(reg5243)
				return
			} else {
				reg5244 := False
				__ctx.Return(reg5244)
				return
			}
		} else {
			reg5245 := False
			__ctx.Return(reg5245)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "tuple?", value: __defun__tuple_2})

	__defun__append = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2000 := __args[0]
		_ = V2000
		V2001 := __args[1]
		_ = V2001
		reg5246 := Nil
		reg5247 := PrimEqual(reg5246, V2000)
		if reg5247 == True {
			__ctx.Return(V2001)
			return
		} else {
			reg5248 := PrimIsPair(V2000)
			if reg5248 == True {
				reg5249 := PrimHead(V2000)
				reg5250 := PrimTail(V2000)
				reg5251 := __e.Call(__defun__append, reg5250, V2001)
				reg5252 := PrimCons(reg5249, reg5251)
				__ctx.Return(reg5252)
				return
			} else {
				reg5253 := MakeSymbol("append")
				__ctx.TailApply(__defun__shen_4f__error, reg5253)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "append", value: __defun__append})

	__defun___8v = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2004 := __args[0]
		_ = V2004
		V2005 := __args[1]
		_ = V2005
		reg5255 := __e.Call(__defun__limit, V2005)
		Limit := reg5255
		_ = Limit
		reg5256 := MakeNumber(1)
		reg5257 := PrimNumberAdd(Limit, reg5256)
		reg5258 := __e.Call(__defun__vector, reg5257)
		NewVector := reg5258
		_ = NewVector
		reg5259 := MakeNumber(1)
		reg5260 := __e.Call(__defun__vector_1_6, NewVector, reg5259, V2004)
		X_7NewVector := reg5260
		_ = X_7NewVector
		reg5261 := MakeNumber(0)
		reg5262 := PrimEqual(Limit, reg5261)
		if reg5262 == True {
			__ctx.Return(X_7NewVector)
			return
		} else {
			reg5263 := MakeNumber(1)
			__ctx.TailApply(__defun__shen_4_8v_1help, V2005, reg5263, Limit, X_7NewVector)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "@v", value: __defun___8v})

	__defun__shen_4_8v_1help = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2011 := __args[0]
		_ = V2011
		V2012 := __args[1]
		_ = V2012
		V2013 := __args[2]
		_ = V2013
		V2014 := __args[3]
		_ = V2014
		reg5265 := PrimEqual(V2013, V2012)
		if reg5265 == True {
			reg5266 := MakeNumber(1)
			reg5267 := PrimNumberAdd(V2013, reg5266)
			__ctx.TailApply(__defun__shen_4copyfromvector, V2011, V2014, V2013, reg5267)
			return
		} else {
			reg5269 := MakeNumber(1)
			reg5270 := PrimNumberAdd(V2012, reg5269)
			reg5271 := MakeNumber(1)
			reg5272 := PrimNumberAdd(V2012, reg5271)
			reg5273 := __e.Call(__defun__shen_4copyfromvector, V2011, V2014, V2012, reg5272)
			__ctx.TailApply(__defun__shen_4_8v_1help, V2011, reg5270, V2013, reg5273)
			return
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.@v-help", value: __defun__shen_4_8v_1help})

	__defun__shen_4copyfromvector = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2019 := __args[0]
		_ = V2019
		V2020 := __args[1]
		_ = V2020
		V2021 := __args[2]
		_ = V2021
		V2022 := __args[3]
		_ = V2022
		reg5275 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg5276 := __e.Call(__defun___5_1vector, V2019, V2021)
			__ctx.TailApply(__defun__vector_1_6, V2020, V2022, reg5276)
			return
		}, 0)
		reg5278 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			__ctx.Return(V2020)
			return
		}, 1)
		reg5279 := __e.Try(reg5275).Catch(reg5278)
		__ctx.Return(reg5279)
		return
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.copyfromvector", value: __defun__shen_4copyfromvector})

	__defun__hdv = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2024 := __args[0]
		_ = V2024
		reg5280 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg5281 := MakeNumber(1)
			__ctx.TailApply(__defun___5_1vector, V2024, reg5281)
			return
		}, 0)
		reg5283 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg5284 := MakeString("hdv needs a non-empty vector as an argument; not ")
			reg5285 := MakeString("\n")
			reg5286 := MakeSymbol("shen.s")
			reg5287 := __e.Call(__defun__shen_4app, V2024, reg5285, reg5286)
			reg5288 := PrimStringConcat(reg5284, reg5287)
			reg5289 := PrimSimpleError(reg5288)
			__ctx.Return(reg5289)
			return
		}, 1)
		reg5290 := __e.Try(reg5280).Catch(reg5283)
		__ctx.Return(reg5290)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "hdv", value: __defun__hdv})

	__defun__tlv = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2026 := __args[0]
		_ = V2026
		reg5291 := __e.Call(__defun__limit, V2026)
		Limit := reg5291
		_ = Limit
		reg5292 := MakeNumber(0)
		reg5293 := PrimEqual(Limit, reg5292)
		if reg5293 == True {
			reg5294 := MakeString("cannot take the tail of the empty vector\n")
			reg5295 := PrimSimpleError(reg5294)
			__ctx.Return(reg5295)
			return
		} else {
			reg5296 := MakeNumber(1)
			reg5297 := PrimEqual(Limit, reg5296)
			if reg5297 == True {
				reg5298 := MakeNumber(0)
				__ctx.TailApply(__defun__vector, reg5298)
				return
			} else {
				reg5300 := MakeNumber(1)
				reg5301 := PrimNumberSubtract(Limit, reg5300)
				reg5302 := __e.Call(__defun__vector, reg5301)
				NewVector := reg5302
				_ = NewVector
				reg5303 := MakeNumber(2)
				reg5304 := MakeNumber(1)
				reg5305 := PrimNumberSubtract(Limit, reg5304)
				reg5306 := __e.Call(__defun__vector, reg5305)
				__ctx.TailApply(__defun__shen_4tlv_1help, V2026, reg5303, Limit, reg5306)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "tlv", value: __defun__tlv})

	__defun__shen_4tlv_1help = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2032 := __args[0]
		_ = V2032
		V2033 := __args[1]
		_ = V2033
		V2034 := __args[2]
		_ = V2034
		V2035 := __args[3]
		_ = V2035
		reg5308 := PrimEqual(V2034, V2033)
		if reg5308 == True {
			reg5309 := MakeNumber(1)
			reg5310 := PrimNumberSubtract(V2034, reg5309)
			__ctx.TailApply(__defun__shen_4copyfromvector, V2032, V2035, V2034, reg5310)
			return
		} else {
			reg5312 := MakeNumber(1)
			reg5313 := PrimNumberAdd(V2033, reg5312)
			reg5314 := MakeNumber(1)
			reg5315 := PrimNumberSubtract(V2033, reg5314)
			reg5316 := __e.Call(__defun__shen_4copyfromvector, V2032, V2035, V2033, reg5315)
			__ctx.TailApply(__defun__shen_4tlv_1help, V2032, reg5313, V2034, reg5316)
			return
		}
	}, 4)
	__initDefs = append(__initDefs, defType{name: "shen.tlv-help", value: __defun__shen_4tlv_1help})

	__defun__assoc = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2047 := __args[0]
		_ = V2047
		V2048 := __args[1]
		_ = V2048
		reg5318 := Nil
		reg5319 := PrimEqual(reg5318, V2048)
		if reg5319 == True {
			reg5320 := Nil
			__ctx.Return(reg5320)
			return
		} else {
			reg5321 := PrimIsPair(V2048)
			var reg5336 Obj
			if reg5321 == True {
				reg5322 := PrimHead(V2048)
				reg5323 := PrimIsPair(reg5322)
				var reg5331 Obj
				if reg5323 == True {
					reg5324 := PrimHead(V2048)
					reg5325 := PrimHead(reg5324)
					reg5326 := PrimEqual(reg5325, V2047)
					var reg5329 Obj
					if reg5326 == True {
						reg5327 := True
						reg5329 = reg5327
					} else {
						reg5328 := False
						reg5329 = reg5328
					}
					reg5331 = reg5329
				} else {
					reg5330 := False
					reg5331 = reg5330
				}
				var reg5334 Obj
				if reg5331 == True {
					reg5332 := True
					reg5334 = reg5332
				} else {
					reg5333 := False
					reg5334 = reg5333
				}
				reg5336 = reg5334
			} else {
				reg5335 := False
				reg5336 = reg5335
			}
			if reg5336 == True {
				reg5337 := PrimHead(V2048)
				__ctx.Return(reg5337)
				return
			} else {
				reg5338 := PrimIsPair(V2048)
				if reg5338 == True {
					reg5339 := PrimTail(V2048)
					__ctx.TailApply(__defun__assoc, V2047, reg5339)
					return
				} else {
					reg5341 := MakeSymbol("assoc")
					__ctx.TailApply(__defun__shen_4f__error, reg5341)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "assoc", value: __defun__assoc})

	__defun__shen_4assoc_1set = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2055 := __args[0]
		_ = V2055
		V2056 := __args[1]
		_ = V2056
		V2057 := __args[2]
		_ = V2057
		reg5343 := Nil
		reg5344 := PrimEqual(reg5343, V2057)
		if reg5344 == True {
			reg5345 := PrimCons(V2055, V2056)
			reg5346 := Nil
			reg5347 := PrimCons(reg5345, reg5346)
			__ctx.Return(reg5347)
			return
		} else {
			reg5348 := PrimIsPair(V2057)
			var reg5363 Obj
			if reg5348 == True {
				reg5349 := PrimHead(V2057)
				reg5350 := PrimIsPair(reg5349)
				var reg5358 Obj
				if reg5350 == True {
					reg5351 := PrimHead(V2057)
					reg5352 := PrimHead(reg5351)
					reg5353 := PrimEqual(reg5352, V2055)
					var reg5356 Obj
					if reg5353 == True {
						reg5354 := True
						reg5356 = reg5354
					} else {
						reg5355 := False
						reg5356 = reg5355
					}
					reg5358 = reg5356
				} else {
					reg5357 := False
					reg5358 = reg5357
				}
				var reg5361 Obj
				if reg5358 == True {
					reg5359 := True
					reg5361 = reg5359
				} else {
					reg5360 := False
					reg5361 = reg5360
				}
				reg5363 = reg5361
			} else {
				reg5362 := False
				reg5363 = reg5362
			}
			if reg5363 == True {
				reg5364 := PrimHead(V2057)
				reg5365 := PrimHead(reg5364)
				reg5366 := PrimCons(reg5365, V2056)
				reg5367 := PrimTail(V2057)
				reg5368 := PrimCons(reg5366, reg5367)
				__ctx.Return(reg5368)
				return
			} else {
				reg5369 := PrimIsPair(V2057)
				if reg5369 == True {
					reg5370 := PrimHead(V2057)
					reg5371 := PrimTail(V2057)
					reg5372 := __e.Call(__defun__shen_4assoc_1set, V2055, V2056, reg5371)
					reg5373 := PrimCons(reg5370, reg5372)
					__ctx.Return(reg5373)
					return
				} else {
					reg5374 := MakeSymbol("shen.assoc-set")
					__ctx.TailApply(__defun__shen_4f__error, reg5374)
					return
				}
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.assoc-set", value: __defun__shen_4assoc_1set})

	__defun__shen_4assoc_1rm = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2063 := __args[0]
		_ = V2063
		V2064 := __args[1]
		_ = V2064
		reg5376 := Nil
		reg5377 := PrimEqual(reg5376, V2064)
		if reg5377 == True {
			reg5378 := Nil
			__ctx.Return(reg5378)
			return
		} else {
			reg5379 := PrimIsPair(V2064)
			var reg5394 Obj
			if reg5379 == True {
				reg5380 := PrimHead(V2064)
				reg5381 := PrimIsPair(reg5380)
				var reg5389 Obj
				if reg5381 == True {
					reg5382 := PrimHead(V2064)
					reg5383 := PrimHead(reg5382)
					reg5384 := PrimEqual(reg5383, V2063)
					var reg5387 Obj
					if reg5384 == True {
						reg5385 := True
						reg5387 = reg5385
					} else {
						reg5386 := False
						reg5387 = reg5386
					}
					reg5389 = reg5387
				} else {
					reg5388 := False
					reg5389 = reg5388
				}
				var reg5392 Obj
				if reg5389 == True {
					reg5390 := True
					reg5392 = reg5390
				} else {
					reg5391 := False
					reg5392 = reg5391
				}
				reg5394 = reg5392
			} else {
				reg5393 := False
				reg5394 = reg5393
			}
			if reg5394 == True {
				reg5395 := PrimTail(V2064)
				__ctx.Return(reg5395)
				return
			} else {
				reg5396 := PrimIsPair(V2064)
				if reg5396 == True {
					reg5397 := PrimHead(V2064)
					reg5398 := PrimTail(V2064)
					reg5399 := __e.Call(__defun__shen_4assoc_1rm, V2063, reg5398)
					reg5400 := PrimCons(reg5397, reg5399)
					__ctx.Return(reg5400)
					return
				} else {
					reg5401 := MakeSymbol("shen.assoc-rm")
					__ctx.TailApply(__defun__shen_4f__error, reg5401)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.assoc-rm", value: __defun__shen_4assoc_1rm})

	__defun__boolean_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2070 := __args[0]
		_ = V2070
		reg5403 := True
		reg5404 := PrimEqual(reg5403, V2070)
		if reg5404 == True {
			reg5405 := True
			__ctx.Return(reg5405)
			return
		} else {
			reg5406 := False
			reg5407 := PrimEqual(reg5406, V2070)
			if reg5407 == True {
				reg5408 := True
				__ctx.Return(reg5408)
				return
			} else {
				reg5409 := False
				__ctx.Return(reg5409)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "boolean?", value: __defun__boolean_2})

	__defun__nl = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2072 := __args[0]
		_ = V2072
		reg5410 := MakeNumber(0)
		reg5411 := PrimEqual(reg5410, V2072)
		if reg5411 == True {
			reg5412 := MakeNumber(0)
			__ctx.Return(reg5412)
			return
		} else {
			reg5413 := MakeString("\n")
			reg5414 := __e.Call(__defun__stoutput)
			reg5415 := __e.Call(__defun__shen_4prhush, reg5413, reg5414)
			_ = reg5415
			reg5416 := MakeNumber(1)
			reg5417 := PrimNumberSubtract(V2072, reg5416)
			__ctx.TailApply(__defun__nl, reg5417)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "nl", value: __defun__nl})

	__defun__difference = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2077 := __args[0]
		_ = V2077
		V2078 := __args[1]
		_ = V2078
		reg5419 := Nil
		reg5420 := PrimEqual(reg5419, V2077)
		if reg5420 == True {
			reg5421 := Nil
			__ctx.Return(reg5421)
			return
		} else {
			reg5422 := PrimIsPair(V2077)
			if reg5422 == True {
				reg5423 := PrimHead(V2077)
				reg5424 := __e.Call(__defun__element_2, reg5423, V2078)
				if reg5424 == True {
					reg5425 := PrimTail(V2077)
					__ctx.TailApply(__defun__difference, reg5425, V2078)
					return
				} else {
					reg5427 := PrimHead(V2077)
					reg5428 := PrimTail(V2077)
					reg5429 := __e.Call(__defun__difference, reg5428, V2078)
					reg5430 := PrimCons(reg5427, reg5429)
					__ctx.Return(reg5430)
					return
				}
			} else {
				reg5431 := MakeSymbol("difference")
				__ctx.TailApply(__defun__shen_4f__error, reg5431)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "difference", value: __defun__difference})

	__defun__do = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2081 := __args[0]
		_ = V2081
		V2082 := __args[1]
		_ = V2082
		__ctx.Return(V2082)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "do", value: __defun__do})

	__defun__element_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2094 := __args[0]
		_ = V2094
		V2095 := __args[1]
		_ = V2095
		reg5433 := Nil
		reg5434 := PrimEqual(reg5433, V2095)
		if reg5434 == True {
			reg5435 := False
			__ctx.Return(reg5435)
			return
		} else {
			reg5436 := PrimIsPair(V2095)
			var reg5443 Obj
			if reg5436 == True {
				reg5437 := PrimHead(V2095)
				reg5438 := PrimEqual(reg5437, V2094)
				var reg5441 Obj
				if reg5438 == True {
					reg5439 := True
					reg5441 = reg5439
				} else {
					reg5440 := False
					reg5441 = reg5440
				}
				reg5443 = reg5441
			} else {
				reg5442 := False
				reg5443 = reg5442
			}
			if reg5443 == True {
				reg5444 := True
				__ctx.Return(reg5444)
				return
			} else {
				reg5445 := PrimIsPair(V2095)
				if reg5445 == True {
					reg5446 := PrimTail(V2095)
					__ctx.TailApply(__defun__element_2, V2094, reg5446)
					return
				} else {
					reg5448 := MakeSymbol("element?")
					__ctx.TailApply(__defun__shen_4f__error, reg5448)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "element?", value: __defun__element_2})

	__defun__empty_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2101 := __args[0]
		_ = V2101
		reg5450 := Nil
		reg5451 := PrimEqual(reg5450, V2101)
		if reg5451 == True {
			reg5452 := True
			__ctx.Return(reg5452)
			return
		} else {
			reg5453 := False
			__ctx.Return(reg5453)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "empty?", value: __defun__empty_2})

	__defun__fix = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2104 := __args[0]
		_ = V2104
		V2105 := __args[1]
		_ = V2105
		reg5454 := __e.Call(V2104, V2105)
		__ctx.TailApply(__defun__shen_4fix_1help, V2104, V2105, reg5454)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "fix", value: __defun__fix})

	__defun__shen_4fix_1help = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2116 := __args[0]
		_ = V2116
		V2117 := __args[1]
		_ = V2117
		V2118 := __args[2]
		_ = V2118
		reg5456 := PrimEqual(V2118, V2117)
		if reg5456 == True {
			__ctx.Return(V2118)
			return
		} else {
			reg5457 := __e.Call(V2116, V2118)
			__ctx.TailApply(__defun__shen_4fix_1help, V2116, V2118, reg5457)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "shen.fix-help", value: __defun__shen_4fix_1help})

	__defun__put = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2123 := __args[0]
		_ = V2123
		V2124 := __args[1]
		_ = V2124
		V2125 := __args[2]
		_ = V2125
		V2126 := __args[3]
		_ = V2126
		reg5459 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			__ctx.TailApply(__defun__shen_4_5_1dict, V2126, V2123)
			return
		}, 0)
		reg5461 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg5462 := Nil
			__ctx.Return(reg5462)
			return
		}, 1)
		reg5463 := __e.Try(reg5459).Catch(reg5461)
		Curr := reg5463
		_ = Curr
		reg5464 := __e.Call(__defun__shen_4assoc_1set, V2124, V2125, Curr)
		Added := reg5464
		_ = Added
		reg5465 := __e.Call(__defun__shen_4dict_1_6, V2126, V2123, Added)
		Update := reg5465
		_ = Update
		__ctx.Return(V2125)
		return
	}, 4)
	__initDefs = append(__initDefs, defType{name: "put", value: __defun__put})

	__defun__unput = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2130 := __args[0]
		_ = V2130
		V2131 := __args[1]
		_ = V2131
		V2132 := __args[2]
		_ = V2132
		reg5466 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			__ctx.TailApply(__defun__shen_4_5_1dict, V2132, V2130)
			return
		}, 0)
		reg5468 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg5469 := Nil
			__ctx.Return(reg5469)
			return
		}, 1)
		reg5470 := __e.Try(reg5466).Catch(reg5468)
		Curr := reg5470
		_ = Curr
		reg5471 := __e.Call(__defun__shen_4assoc_1rm, V2131, Curr)
		Removed := reg5471
		_ = Removed
		reg5472 := __e.Call(__defun__shen_4dict_1_6, V2132, V2130, Removed)
		Update := reg5472
		_ = Update
		__ctx.Return(V2130)
		return
	}, 3)
	__initDefs = append(__initDefs, defType{name: "unput", value: __defun__unput})

	__defun__get = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2136 := __args[0]
		_ = V2136
		V2137 := __args[1]
		_ = V2137
		V2138 := __args[2]
		_ = V2138
		reg5473 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			__ctx.TailApply(__defun__shen_4_5_1dict, V2138, V2136)
			return
		}, 0)
		reg5475 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg5476 := Nil
			__ctx.Return(reg5476)
			return
		}, 1)
		reg5477 := __e.Try(reg5473).Catch(reg5475)
		Entry := reg5477
		_ = Entry
		reg5478 := __e.Call(__defun__assoc, V2137, Entry)
		Result := reg5478
		_ = Result
		reg5479 := __e.Call(__defun__empty_2, Result)
		if reg5479 == True {
			reg5480 := MakeString("value not found\n")
			reg5481 := PrimSimpleError(reg5480)
			__ctx.Return(reg5481)
			return
		} else {
			reg5482 := PrimTail(Result)
			__ctx.Return(reg5482)
			return
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "get", value: __defun__get})

	__defun__hash = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2141 := __args[0]
		_ = V2141
		V2142 := __args[1]
		_ = V2142
		reg5483 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			X := __args[0]
			_ = X
			reg5484 := PrimStringToNumber(X)
			__ctx.Return(reg5484)
			return
		}, 1)
		reg5485 := __e.Call(__defun__explode, V2141)
		reg5486 := __e.Call(__defun__map, reg5483, reg5485)
		reg5487 := __e.Call(__defun__sum, reg5486)
		__ctx.TailApply(__defun__shen_4mod, reg5487, V2142)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "hash", value: __defun__hash})

	__defun__shen_4mod = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2145 := __args[0]
		_ = V2145
		V2146 := __args[1]
		_ = V2146
		reg5489 := Nil
		reg5490 := PrimCons(V2146, reg5489)
		reg5491 := __e.Call(__defun__shen_4multiples, V2145, reg5490)
		__ctx.TailApply(__defun__shen_4modh, V2145, reg5491)
		return
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.mod", value: __defun__shen_4mod})

	__defun__shen_4multiples = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2149 := __args[0]
		_ = V2149
		V2150 := __args[1]
		_ = V2150
		reg5493 := PrimIsPair(V2150)
		var reg5500 Obj
		if reg5493 == True {
			reg5494 := PrimHead(V2150)
			reg5495 := PrimGreatThan(reg5494, V2149)
			var reg5498 Obj
			if reg5495 == True {
				reg5496 := True
				reg5498 = reg5496
			} else {
				reg5497 := False
				reg5498 = reg5497
			}
			reg5500 = reg5498
		} else {
			reg5499 := False
			reg5500 = reg5499
		}
		if reg5500 == True {
			reg5501 := PrimTail(V2150)
			__ctx.Return(reg5501)
			return
		} else {
			reg5502 := PrimIsPair(V2150)
			if reg5502 == True {
				reg5503 := MakeNumber(2)
				reg5504 := PrimHead(V2150)
				reg5505 := PrimNumberMultiply(reg5503, reg5504)
				reg5506 := PrimCons(reg5505, V2150)
				__ctx.TailApply(__defun__shen_4multiples, V2149, reg5506)
				return
			} else {
				reg5508 := MakeSymbol("shen.multiples")
				__ctx.TailApply(__defun__shen_4f__error, reg5508)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.multiples", value: __defun__shen_4multiples})

	__defun__shen_4modh = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2155 := __args[0]
		_ = V2155
		V2156 := __args[1]
		_ = V2156
		reg5510 := MakeNumber(0)
		reg5511 := PrimEqual(reg5510, V2155)
		if reg5511 == True {
			reg5512 := MakeNumber(0)
			__ctx.Return(reg5512)
			return
		} else {
			reg5513 := Nil
			reg5514 := PrimEqual(reg5513, V2156)
			if reg5514 == True {
				__ctx.Return(V2155)
				return
			} else {
				reg5515 := PrimIsPair(V2156)
				var reg5522 Obj
				if reg5515 == True {
					reg5516 := PrimHead(V2156)
					reg5517 := PrimGreatThan(reg5516, V2155)
					var reg5520 Obj
					if reg5517 == True {
						reg5518 := True
						reg5520 = reg5518
					} else {
						reg5519 := False
						reg5520 = reg5519
					}
					reg5522 = reg5520
				} else {
					reg5521 := False
					reg5522 = reg5521
				}
				if reg5522 == True {
					reg5523 := PrimTail(V2156)
					reg5524 := __e.Call(__defun__empty_2, reg5523)
					if reg5524 == True {
						__ctx.Return(V2155)
						return
					} else {
						reg5525 := PrimTail(V2156)
						__ctx.TailApply(__defun__shen_4modh, V2155, reg5525)
						return
					}
				} else {
					reg5527 := PrimIsPair(V2156)
					if reg5527 == True {
						reg5528 := PrimHead(V2156)
						reg5529 := PrimNumberSubtract(V2155, reg5528)
						__ctx.TailApply(__defun__shen_4modh, reg5529, V2156)
						return
					} else {
						reg5531 := MakeSymbol("shen.modh")
						__ctx.TailApply(__defun__shen_4f__error, reg5531)
						return
					}
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.modh", value: __defun__shen_4modh})

	__defun__sum = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2158 := __args[0]
		_ = V2158
		reg5533 := Nil
		reg5534 := PrimEqual(reg5533, V2158)
		if reg5534 == True {
			reg5535 := MakeNumber(0)
			__ctx.Return(reg5535)
			return
		} else {
			reg5536 := PrimIsPair(V2158)
			if reg5536 == True {
				reg5537 := PrimHead(V2158)
				reg5538 := PrimTail(V2158)
				reg5539 := __e.Call(__defun__sum, reg5538)
				reg5540 := PrimNumberAdd(reg5537, reg5539)
				__ctx.Return(reg5540)
				return
			} else {
				reg5541 := MakeSymbol("sum")
				__ctx.TailApply(__defun__shen_4f__error, reg5541)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "sum", value: __defun__sum})

	__defun__head = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2166 := __args[0]
		_ = V2166
		reg5543 := PrimIsPair(V2166)
		if reg5543 == True {
			reg5544 := PrimHead(V2166)
			__ctx.Return(reg5544)
			return
		} else {
			reg5545 := MakeString("head expects a non-empty list")
			reg5546 := PrimSimpleError(reg5545)
			__ctx.Return(reg5546)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "head", value: __defun__head})

	__defun__tail = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2174 := __args[0]
		_ = V2174
		reg5547 := PrimIsPair(V2174)
		if reg5547 == True {
			reg5548 := PrimTail(V2174)
			__ctx.Return(reg5548)
			return
		} else {
			reg5549 := MakeString("tail expects a non-empty list")
			reg5550 := PrimSimpleError(reg5549)
			__ctx.Return(reg5550)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "tail", value: __defun__tail})

	__defun__hdstr = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2176 := __args[0]
		_ = V2176
		reg5551 := MakeNumber(0)
		reg5552 := PrimPos(V2176, reg5551)
		__ctx.Return(reg5552)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "hdstr", value: __defun__hdstr})

	__defun__intersection = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2181 := __args[0]
		_ = V2181
		V2182 := __args[1]
		_ = V2182
		reg5553 := Nil
		reg5554 := PrimEqual(reg5553, V2181)
		if reg5554 == True {
			reg5555 := Nil
			__ctx.Return(reg5555)
			return
		} else {
			reg5556 := PrimIsPair(V2181)
			if reg5556 == True {
				reg5557 := PrimHead(V2181)
				reg5558 := __e.Call(__defun__element_2, reg5557, V2182)
				if reg5558 == True {
					reg5559 := PrimHead(V2181)
					reg5560 := PrimTail(V2181)
					reg5561 := __e.Call(__defun__intersection, reg5560, V2182)
					reg5562 := PrimCons(reg5559, reg5561)
					__ctx.Return(reg5562)
					return
				} else {
					reg5563 := PrimTail(V2181)
					__ctx.TailApply(__defun__intersection, reg5563, V2182)
					return
				}
			} else {
				reg5565 := MakeSymbol("intersection")
				__ctx.TailApply(__defun__shen_4f__error, reg5565)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "intersection", value: __defun__intersection})

	__defun__reverse = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2184 := __args[0]
		_ = V2184
		reg5567 := Nil
		__ctx.TailApply(__defun__shen_4reverse__help, V2184, reg5567)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "reverse", value: __defun__reverse})

	__defun__shen_4reverse__help = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2187 := __args[0]
		_ = V2187
		V2188 := __args[1]
		_ = V2188
		reg5569 := Nil
		reg5570 := PrimEqual(reg5569, V2187)
		if reg5570 == True {
			__ctx.Return(V2188)
			return
		} else {
			reg5571 := PrimIsPair(V2187)
			if reg5571 == True {
				reg5572 := PrimTail(V2187)
				reg5573 := PrimHead(V2187)
				reg5574 := PrimCons(reg5573, V2188)
				__ctx.TailApply(__defun__shen_4reverse__help, reg5572, reg5574)
				return
			} else {
				reg5576 := MakeSymbol("shen.reverse_help")
				__ctx.TailApply(__defun__shen_4f__error, reg5576)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.reverse_help", value: __defun__shen_4reverse__help})

	__defun__union = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2191 := __args[0]
		_ = V2191
		V2192 := __args[1]
		_ = V2192
		reg5578 := Nil
		reg5579 := PrimEqual(reg5578, V2191)
		if reg5579 == True {
			__ctx.Return(V2192)
			return
		} else {
			reg5580 := PrimIsPair(V2191)
			if reg5580 == True {
				reg5581 := PrimHead(V2191)
				reg5582 := __e.Call(__defun__element_2, reg5581, V2192)
				if reg5582 == True {
					reg5583 := PrimTail(V2191)
					__ctx.TailApply(__defun__union, reg5583, V2192)
					return
				} else {
					reg5585 := PrimHead(V2191)
					reg5586 := PrimTail(V2191)
					reg5587 := __e.Call(__defun__union, reg5586, V2192)
					reg5588 := PrimCons(reg5585, reg5587)
					__ctx.Return(reg5588)
					return
				}
			} else {
				reg5589 := MakeSymbol("union")
				__ctx.TailApply(__defun__shen_4f__error, reg5589)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "union", value: __defun__union})

	__defun__y_1or_1n_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2194 := __args[0]
		_ = V2194
		reg5591 := __e.Call(__defun__shen_4proc_1nl, V2194)
		reg5592 := __e.Call(__defun__stoutput)
		reg5593 := __e.Call(__defun__shen_4prhush, reg5591, reg5592)
		Message := reg5593
		_ = Message
		reg5594 := MakeString(" (y/n) ")
		reg5595 := __e.Call(__defun__stoutput)
		reg5596 := __e.Call(__defun__shen_4prhush, reg5594, reg5595)
		Y_1or_1N := reg5596
		_ = Y_1or_1N
		reg5597 := __e.Call(__defun__stinput)
		reg5598 := __e.Call(__defun__read, reg5597)
		reg5599 := MakeString("")
		reg5600 := MakeSymbol("shen.s")
		reg5601 := __e.Call(__defun__shen_4app, reg5598, reg5599, reg5600)
		Input := reg5601
		_ = Input
		reg5602 := MakeString("y")
		reg5603 := PrimEqual(reg5602, Input)
		if reg5603 == True {
			reg5604 := True
			__ctx.Return(reg5604)
			return
		} else {
			reg5605 := MakeString("n")
			reg5606 := PrimEqual(reg5605, Input)
			if reg5606 == True {
				reg5607 := False
				__ctx.Return(reg5607)
				return
			} else {
				reg5608 := MakeString("please answer y or n\n")
				reg5609 := __e.Call(__defun__stoutput)
				reg5610 := __e.Call(__defun__shen_4prhush, reg5608, reg5609)
				_ = reg5610
				__ctx.TailApply(__defun__y_1or_1n_2, V2194)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "y-or-n?", value: __defun__y_1or_1n_2})

	__defun__not = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2196 := __args[0]
		_ = V2196
		if V2196 == True {
			reg5612 := False
			__ctx.Return(reg5612)
			return
		} else {
			reg5613 := True
			__ctx.Return(reg5613)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "not", value: __defun__not})

	__defun__subst = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2209 := __args[0]
		_ = V2209
		V2210 := __args[1]
		_ = V2210
		V2211 := __args[2]
		_ = V2211
		reg5614 := PrimEqual(V2211, V2210)
		if reg5614 == True {
			__ctx.Return(V2209)
			return
		} else {
			reg5615 := PrimIsPair(V2211)
			if reg5615 == True {
				reg5616 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
					W := __args[0]
					_ = W
					__ctx.TailApply(__defun__subst, V2209, V2210, W)
					return
				}, 1)
				__ctx.TailApply(__defun__map, reg5616, V2211)
				return
			} else {
				__ctx.Return(V2211)
				return
			}
		}
	}, 3)
	__initDefs = append(__initDefs, defType{name: "subst", value: __defun__subst})

	__defun__explode = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2213 := __args[0]
		_ = V2213
		reg5619 := MakeString("")
		reg5620 := MakeSymbol("shen.a")
		reg5621 := __e.Call(__defun__shen_4app, V2213, reg5619, reg5620)
		__ctx.TailApply(__defun__shen_4explode_1h, reg5621)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "explode", value: __defun__explode})

	__defun__shen_4explode_1h = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2215 := __args[0]
		_ = V2215
		reg5623 := MakeString("")
		reg5624 := PrimEqual(reg5623, V2215)
		if reg5624 == True {
			reg5625 := Nil
			__ctx.Return(reg5625)
			return
		} else {
			reg5626 := __e.Call(__defun__shen_4_7string_2, V2215)
			if reg5626 == True {
				reg5627 := MakeNumber(0)
				reg5628 := PrimPos(V2215, reg5627)
				reg5629 := PrimTailString(V2215)
				reg5630 := __e.Call(__defun__shen_4explode_1h, reg5629)
				reg5631 := PrimCons(reg5628, reg5630)
				__ctx.Return(reg5631)
				return
			} else {
				reg5632 := MakeSymbol("shen.explode-h")
				__ctx.TailApply(__defun__shen_4f__error, reg5632)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.explode-h", value: __defun__shen_4explode_1h})

	__defun__cd = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2217 := __args[0]
		_ = V2217
		reg5634 := MakeSymbol("*home-directory*")
		reg5635 := MakeString("")
		reg5636 := PrimEqual(V2217, reg5635)
		var reg5641 Obj
		if reg5636 == True {
			reg5637 := MakeString("")
			reg5641 = reg5637
		} else {
			reg5638 := MakeString("/")
			reg5639 := MakeSymbol("shen.a")
			reg5640 := __e.Call(__defun__shen_4app, V2217, reg5638, reg5639)
			reg5641 = reg5640
		}
		reg5642 := PrimSet(reg5634, reg5641)
		__ctx.Return(reg5642)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "cd", value: __defun__cd})

	__defun__shen_4for_1each = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2220 := __args[0]
		_ = V2220
		V2221 := __args[1]
		_ = V2221
		reg5643 := Nil
		reg5644 := PrimEqual(reg5643, V2221)
		if reg5644 == True {
			reg5645 := True
			__ctx.Return(reg5645)
			return
		} else {
			reg5646 := PrimIsPair(V2221)
			if reg5646 == True {
				reg5647 := PrimHead(V2221)
				reg5648 := __e.Call(V2220, reg5647)
				__ := reg5648
				_ = __
				reg5649 := PrimTail(V2221)
				__ctx.TailApply(__defun__shen_4for_1each, V2220, reg5649)
				return
			} else {
				reg5651 := MakeSymbol("shen.for-each")
				__ctx.TailApply(__defun__shen_4f__error, reg5651)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.for-each", value: __defun__shen_4for_1each})

	__defun__map = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2226 := __args[0]
		_ = V2226
		V2227 := __args[1]
		_ = V2227
		reg5653 := Nil
		reg5654 := PrimEqual(reg5653, V2227)
		if reg5654 == True {
			reg5655 := Nil
			__ctx.Return(reg5655)
			return
		} else {
			reg5656 := PrimIsPair(V2227)
			if reg5656 == True {
				reg5657 := PrimHead(V2227)
				reg5658 := __e.Call(V2226, reg5657)
				reg5659 := PrimTail(V2227)
				reg5660 := __e.Call(__defun__map, V2226, reg5659)
				reg5661 := PrimCons(reg5658, reg5660)
				__ctx.Return(reg5661)
				return
			} else {
				__ctx.TailApply(V2226, V2227)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "map", value: __defun__map})

	__defun__length = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2229 := __args[0]
		_ = V2229
		reg5663 := MakeNumber(0)
		__ctx.TailApply(__defun__shen_4length_1h, V2229, reg5663)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "length", value: __defun__length})

	__defun__shen_4length_1h = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2232 := __args[0]
		_ = V2232
		V2233 := __args[1]
		_ = V2233
		reg5665 := Nil
		reg5666 := PrimEqual(reg5665, V2232)
		if reg5666 == True {
			__ctx.Return(V2233)
			return
		} else {
			reg5667 := PrimTail(V2232)
			reg5668 := MakeNumber(1)
			reg5669 := PrimNumberAdd(V2233, reg5668)
			__ctx.TailApply(__defun__shen_4length_1h, reg5667, reg5669)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.length-h", value: __defun__shen_4length_1h})

	__defun__occurrences = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2245 := __args[0]
		_ = V2245
		V2246 := __args[1]
		_ = V2246
		reg5671 := PrimEqual(V2246, V2245)
		if reg5671 == True {
			reg5672 := MakeNumber(1)
			__ctx.Return(reg5672)
			return
		} else {
			reg5673 := PrimIsPair(V2246)
			if reg5673 == True {
				reg5674 := PrimHead(V2246)
				reg5675 := __e.Call(__defun__occurrences, V2245, reg5674)
				reg5676 := PrimTail(V2246)
				reg5677 := __e.Call(__defun__occurrences, V2245, reg5676)
				reg5678 := PrimNumberAdd(reg5675, reg5677)
				__ctx.Return(reg5678)
				return
			} else {
				reg5679 := MakeNumber(0)
				__ctx.Return(reg5679)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "occurrences", value: __defun__occurrences})

	__defun__nth = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2253 := __args[0]
		_ = V2253
		V2254 := __args[1]
		_ = V2254
		reg5680 := MakeNumber(1)
		reg5681 := PrimEqual(reg5680, V2253)
		var reg5687 Obj
		if reg5681 == True {
			reg5682 := PrimIsPair(V2254)
			var reg5685 Obj
			if reg5682 == True {
				reg5683 := True
				reg5685 = reg5683
			} else {
				reg5684 := False
				reg5685 = reg5684
			}
			reg5687 = reg5685
		} else {
			reg5686 := False
			reg5687 = reg5686
		}
		if reg5687 == True {
			reg5688 := PrimHead(V2254)
			__ctx.Return(reg5688)
			return
		} else {
			reg5689 := PrimIsPair(V2254)
			if reg5689 == True {
				reg5690 := MakeNumber(1)
				reg5691 := PrimNumberSubtract(V2253, reg5690)
				reg5692 := PrimTail(V2254)
				__ctx.TailApply(__defun__nth, reg5691, reg5692)
				return
			} else {
				reg5694 := MakeString("nth applied to ")
				reg5695 := MakeString(", ")
				reg5696 := MakeString("\n")
				reg5697 := MakeSymbol("shen.a")
				reg5698 := __e.Call(__defun__shen_4app, V2254, reg5696, reg5697)
				reg5699 := PrimStringConcat(reg5695, reg5698)
				reg5700 := MakeSymbol("shen.a")
				reg5701 := __e.Call(__defun__shen_4app, V2253, reg5699, reg5700)
				reg5702 := PrimStringConcat(reg5694, reg5701)
				reg5703 := PrimSimpleError(reg5702)
				__ctx.Return(reg5703)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "nth", value: __defun__nth})

	__defun__integer_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2256 := __args[0]
		_ = V2256
		reg5704 := PrimIsNumber(V2256)
		if reg5704 == True {
			reg5705 := __e.Call(__defun__shen_4abs, V2256)
			Abs := reg5705
			_ = Abs
			reg5706 := MakeNumber(1)
			reg5707 := __e.Call(__defun__shen_4magless, Abs, reg5706)
			reg5708 := __e.Call(__defun__shen_4integer_1test_2, Abs, reg5707)
			if reg5708 == True {
				reg5709 := True
				__ctx.Return(reg5709)
				return
			} else {
				reg5710 := False
				__ctx.Return(reg5710)
				return
			}
		} else {
			reg5711 := False
			__ctx.Return(reg5711)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "integer?", value: __defun__integer_2})

	__defun__shen_4abs = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2258 := __args[0]
		_ = V2258
		reg5712 := MakeNumber(0)
		reg5713 := PrimGreatThan(V2258, reg5712)
		if reg5713 == True {
			__ctx.Return(V2258)
			return
		} else {
			reg5714 := MakeNumber(0)
			reg5715 := PrimNumberSubtract(reg5714, V2258)
			__ctx.Return(reg5715)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.abs", value: __defun__shen_4abs})

	__defun__shen_4magless = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2261 := __args[0]
		_ = V2261
		V2262 := __args[1]
		_ = V2262
		reg5716 := MakeNumber(2)
		reg5717 := PrimNumberMultiply(V2262, reg5716)
		Nx2 := reg5717
		_ = Nx2
		reg5718 := PrimGreatThan(Nx2, V2261)
		if reg5718 == True {
			__ctx.Return(V2262)
			return
		} else {
			__ctx.TailApply(__defun__shen_4magless, V2261, Nx2)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.magless", value: __defun__shen_4magless})

	__defun__shen_4integer_1test_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2268 := __args[0]
		_ = V2268
		V2269 := __args[1]
		_ = V2269
		reg5720 := MakeNumber(0)
		reg5721 := PrimEqual(reg5720, V2268)
		if reg5721 == True {
			reg5722 := True
			__ctx.Return(reg5722)
			return
		} else {
			reg5723 := MakeNumber(1)
			reg5724 := PrimGreatThan(reg5723, V2268)
			if reg5724 == True {
				reg5725 := False
				__ctx.Return(reg5725)
				return
			} else {
				reg5726 := PrimNumberSubtract(V2268, V2269)
				Abs_1N := reg5726
				_ = Abs_1N
				reg5727 := MakeNumber(0)
				reg5728 := PrimGreatThan(reg5727, Abs_1N)
				if reg5728 == True {
					reg5729 := PrimIsInteger(V2268)
					__ctx.Return(reg5729)
					return
				} else {
					__ctx.TailApply(__defun__shen_4integer_1test_2, Abs_1N, V2269)
					return
				}
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "shen.integer-test?", value: __defun__shen_4integer_1test_2})

	__defun__mapcan = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2274 := __args[0]
		_ = V2274
		V2275 := __args[1]
		_ = V2275
		reg5731 := Nil
		reg5732 := PrimEqual(reg5731, V2275)
		if reg5732 == True {
			reg5733 := Nil
			__ctx.Return(reg5733)
			return
		} else {
			reg5734 := PrimIsPair(V2275)
			if reg5734 == True {
				reg5735 := PrimHead(V2275)
				reg5736 := __e.Call(V2274, reg5735)
				reg5737 := PrimTail(V2275)
				reg5738 := __e.Call(__defun__mapcan, V2274, reg5737)
				__ctx.TailApply(__defun__append, reg5736, reg5738)
				return
			} else {
				reg5740 := MakeSymbol("mapcan")
				__ctx.TailApply(__defun__shen_4f__error, reg5740)
				return
			}
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "mapcan", value: __defun__mapcan})

	__defun___a_a = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2287 := __args[0]
		_ = V2287
		V2288 := __args[1]
		_ = V2288
		reg5742 := PrimEqual(V2288, V2287)
		if reg5742 == True {
			reg5743 := True
			__ctx.Return(reg5743)
			return
		} else {
			reg5744 := False
			__ctx.Return(reg5744)
			return
		}
	}, 2)
	__initDefs = append(__initDefs, defType{name: "==", value: __defun___a_a})

	__defun__abort = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5745 := MakeString("")
		reg5746 := PrimSimpleError(reg5745)
		__ctx.Return(reg5746)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "abort", value: __defun__abort})

	__defun__bound_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2290 := __args[0]
		_ = V2290
		reg5747 := PrimIsSymbol(V2290)
		if reg5747 == True {
			reg5748 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				reg5749 := PrimValue(V2290)
				__ctx.Return(reg5749)
				return
			}, 0)
			reg5750 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
				E := __args[0]
				_ = E
				reg5751 := MakeSymbol("shen.this-symbol-is-unbound")
				__ctx.Return(reg5751)
				return
			}, 1)
			reg5752 := __e.Try(reg5748).Catch(reg5750)
			Val := reg5752
			_ = Val
			reg5753 := MakeSymbol("shen.this-symbol-is-unbound")
			reg5754 := PrimEqual(Val, reg5753)
			var reg5757 Obj
			if reg5754 == True {
				reg5755 := False
				reg5757 = reg5755
			} else {
				reg5756 := True
				reg5757 = reg5756
			}
			if reg5757 == True {
				reg5758 := True
				__ctx.Return(reg5758)
				return
			} else {
				reg5759 := False
				__ctx.Return(reg5759)
				return
			}
		} else {
			reg5760 := False
			__ctx.Return(reg5760)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "bound?", value: __defun__bound_2})

	__defun__shen_4string_1_6bytes = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2292 := __args[0]
		_ = V2292
		reg5761 := MakeString("")
		reg5762 := PrimEqual(reg5761, V2292)
		if reg5762 == True {
			reg5763 := Nil
			__ctx.Return(reg5763)
			return
		} else {
			reg5764 := MakeNumber(0)
			reg5765 := PrimPos(V2292, reg5764)
			reg5766 := PrimStringToNumber(reg5765)
			reg5767 := PrimTailString(V2292)
			reg5768 := __e.Call(__defun__shen_4string_1_6bytes, reg5767)
			reg5769 := PrimCons(reg5766, reg5768)
			__ctx.Return(reg5769)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.string->bytes", value: __defun__shen_4string_1_6bytes})

	__defun__maxinferences = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2294 := __args[0]
		_ = V2294
		reg5770 := MakeSymbol("shen.*maxinferences*")
		reg5771 := PrimSet(reg5770, V2294)
		__ctx.Return(reg5771)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "maxinferences", value: __defun__maxinferences})

	__defun__inferences = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5772 := MakeSymbol("shen.*infs*")
		reg5773 := PrimValue(reg5772)
		__ctx.Return(reg5773)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "inferences", value: __defun__inferences})

	__defun__protect = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2296 := __args[0]
		_ = V2296
		__ctx.Return(V2296)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "protect", value: __defun__protect})

	__defun__stoutput = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5774 := MakeSymbol("*stoutput*")
		reg5775 := PrimValue(reg5774)
		__ctx.Return(reg5775)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "stoutput", value: __defun__stoutput})

	__defun__sterror = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5776 := MakeSymbol("*sterror*")
		reg5777 := PrimValue(reg5776)
		__ctx.Return(reg5777)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "sterror", value: __defun__sterror})

	__defun__string_1_6symbol = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2298 := __args[0]
		_ = V2298
		reg5778 := PrimIntern(V2298)
		Symbol := reg5778
		_ = Symbol
		reg5779 := PrimIsSymbol(Symbol)
		if reg5779 == True {
			__ctx.Return(Symbol)
			return
		} else {
			reg5780 := MakeString("cannot intern ")
			reg5781 := MakeString(" to a symbol")
			reg5782 := MakeSymbol("shen.s")
			reg5783 := __e.Call(__defun__shen_4app, V2298, reg5781, reg5782)
			reg5784 := PrimStringConcat(reg5780, reg5783)
			reg5785 := PrimSimpleError(reg5784)
			__ctx.Return(reg5785)
			return
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "string->symbol", value: __defun__string_1_6symbol})

	__defun__optimise = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2304 := __args[0]
		_ = V2304
		reg5786 := MakeSymbol("+")
		reg5787 := PrimEqual(reg5786, V2304)
		if reg5787 == True {
			reg5788 := MakeSymbol("shen.*optimise*")
			reg5789 := True
			reg5790 := PrimSet(reg5788, reg5789)
			__ctx.Return(reg5790)
			return
		} else {
			reg5791 := MakeSymbol("-")
			reg5792 := PrimEqual(reg5791, V2304)
			if reg5792 == True {
				reg5793 := MakeSymbol("shen.*optimise*")
				reg5794 := False
				reg5795 := PrimSet(reg5793, reg5794)
				__ctx.Return(reg5795)
				return
			} else {
				reg5796 := MakeString("optimise expects a + or a -.\n")
				reg5797 := PrimSimpleError(reg5796)
				__ctx.Return(reg5797)
				return
			}
		}
	}, 1)
	__initDefs = append(__initDefs, defType{name: "optimise", value: __defun__optimise})

	__defun__os = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5798 := MakeSymbol("*os*")
		reg5799 := PrimValue(reg5798)
		__ctx.Return(reg5799)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "os", value: __defun__os})

	__defun__language = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5800 := MakeSymbol("*language*")
		reg5801 := PrimValue(reg5800)
		__ctx.Return(reg5801)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "language", value: __defun__language})

	__defun__version = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5802 := MakeSymbol("*version*")
		reg5803 := PrimValue(reg5802)
		__ctx.Return(reg5803)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "version", value: __defun__version})

	__defun__port = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5804 := MakeSymbol("*port*")
		reg5805 := PrimValue(reg5804)
		__ctx.Return(reg5805)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "port", value: __defun__port})

	__defun__porters = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5806 := MakeSymbol("*porters*")
		reg5807 := PrimValue(reg5806)
		__ctx.Return(reg5807)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "porters", value: __defun__porters})

	__defun__implementation = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5808 := MakeSymbol("*implementation*")
		reg5809 := PrimValue(reg5808)
		__ctx.Return(reg5809)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "implementation", value: __defun__implementation})

	__defun__release = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		reg5810 := MakeSymbol("*release*")
		reg5811 := PrimValue(reg5810)
		__ctx.Return(reg5811)
		return
	}, 0)
	__initDefs = append(__initDefs, defType{name: "release", value: __defun__release})

	__defun__package_2 = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2306 := __args[0]
		_ = V2306
		reg5812 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg5813 := __e.Call(__defun__external, V2306)
			_ = reg5813
			reg5814 := True
			__ctx.Return(reg5814)
			return
		}, 0)
		reg5815 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg5816 := False
			__ctx.Return(reg5816)
			return
		}, 1)
		reg5817 := __e.Try(reg5812).Catch(reg5815)
		__ctx.Return(reg5817)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "package?", value: __defun__package_2})

	__defun__function = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2308 := __args[0]
		_ = V2308
		__ctx.TailApply(__defun__shen_4lookup_1func, V2308)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "function", value: __defun__function})

	__defun__shen_4lookup_1func = MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
		V2310 := __args[0]
		_ = V2310
		reg5819 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			reg5820 := MakeSymbol("shen.lambda-form")
			reg5821 := MakeSymbol("*property-vector*")
			reg5822 := PrimValue(reg5821)
			__ctx.TailApply(__defun__get, V2310, reg5820, reg5822)
			return
		}, 0)
		reg5824 := MakeNative(func(__e Evaluator, __ctx *ControlFlow, __args ...Obj) {
			E := __args[0]
			_ = E
			reg5825 := MakeString(" has no lambda expansion\n")
			reg5826 := MakeSymbol("shen.a")
			reg5827 := __e.Call(__defun__shen_4app, V2310, reg5825, reg5826)
			reg5828 := PrimSimpleError(reg5827)
			__ctx.Return(reg5828)
			return
		}, 1)
		reg5829 := __e.Try(reg5819).Catch(reg5824)
		__ctx.Return(reg5829)
		return
	}, 1)
	__initDefs = append(__initDefs, defType{name: "shen.lookup-func", value: __defun__shen_4lookup_1func})

}
