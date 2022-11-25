package main

import . "github.com/tiancaiamao/shen-go/kl"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
	tmp4353 := MakeNative(func(__e *ControlFlow) {
		V2726 := __e.Get(1)
		_ = V2726
		tmp4354 := MakeNative(func(__e *ControlFlow) {
			Bytelist := __e.Get(1)
			_ = Bytelist
			tmp4355 := MakeNative(func(__e *ControlFlow) {
				S_1exprs := __e.Get(1)
				_ = S_1exprs
				tmp4356 := MakeNative(func(__e *ControlFlow) {
					Process := __e.Get(1)
					_ = Process
					__e.Return(Process)
					return
				}, 1)

				tmp4357 := Call(__e, PrimFunc(symshen_4process_1sexprs), S_1exprs)

				__e.TailApply(tmp4356, tmp4357)
				return

			}, 1)

			tmp4358 := MakeNative(func(__e *ControlFlow) {
				tmp4359 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), X)
					return
				}, 1)

				__e.TailApply(PrimFunc(symcompile), tmp4359, Bytelist)
				return

			}, 0)

			tmp4360 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				tmp4361 := PrimValue(symshen_4_dresidue_d)

				__e.TailApply(PrimFunc(symshen_4print_1residue), tmp4361)
				return

			}, 1)

			tmp4362 := Call(__e, try_1catch, tmp4358, tmp4360)

			__e.TailApply(tmp4355, tmp4362)
			return

		}, 1)

		tmp4363 := PrimReadFileAsByteList(V2726)

		__e.TailApply(tmp4354, tmp4363)
		return

	}, 1)

	tmp4364 := Call(__e, ns2_1set, symread_1file, tmp4353)

	_ = tmp4364

	tmp4365 := MakeNative(func(__e *ControlFlow) {
		V2727 := __e.Get(1)
		_ = V2727
		tmp4366 := MakeNative(func(__e *ControlFlow) {
			Err := __e.Get(1)
			_ = Err
			__e.TailApply(PrimFunc(symshen_4nchars), MakeNumber(50), V2727)
			return
		}, 1)

		tmp4367 := Call(__e, PrimFunc(symstoutput))

		tmp4368 := Call(__e, PrimFunc(sympr), MakeString("syntax error here:\n"), tmp4367)

		__e.TailApply(tmp4366, tmp4368)
		return

	}, 1)

	tmp4369 := Call(__e, ns2_1set, symshen_4print_1residue, tmp4365)

	_ = tmp4369

	tmp4370 := MakeNative(func(__e *ControlFlow) {
		V2732 := __e.Get(1)
		_ = V2732
		V2733 := __e.Get(2)
		_ = V2733
		tmp4386 := PrimEqual(MakeNumber(0), V2732)

		if True == tmp4386 {
			tmp4371 := Call(__e, PrimFunc(symstoutput))

			tmp4372 := Call(__e, PrimFunc(sympr), MakeString(" ..."), tmp4371)

			_ = tmp4372

			__e.TailApply(PrimFunc(symabort))
			return

		} else {
			tmp4384 := PrimEqual(Nil, V2733)

			if True == tmp4384 {
				tmp4373 := Call(__e, PrimFunc(symstoutput))

				tmp4374 := Call(__e, PrimFunc(sympr), MakeString(" ..."), tmp4373)

				_ = tmp4374

				__e.TailApply(PrimFunc(symabort))
				return

			} else {
				tmp4382 := PrimIsPair(V2733)

				if True == tmp4382 {
					tmp4375 := PrimHead(V2733)

					tmp4376 := PrimNumberToString(tmp4375)

					tmp4377 := Call(__e, PrimFunc(symstoutput))

					tmp4378 := Call(__e, PrimFunc(sympr), tmp4376, tmp4377)

					_ = tmp4378

					tmp4379 := PrimNumberSubtract(V2732, MakeNumber(1))

					tmp4380 := PrimTail(V2733)

					__e.TailApply(PrimFunc(symshen_4nchars), tmp4379, tmp4380)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4nchars)
					return
				}

			}

		}

	}, 2)

	tmp4387 := Call(__e, ns2_1set, symshen_4nchars, tmp4370)

	_ = tmp4387

	tmp4388 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dit_d))
		return
	}, 0)

	tmp4389 := Call(__e, ns2_1set, symit, tmp4388)

	_ = tmp4389

	tmp4390 := MakeNative(func(__e *ControlFlow) {
		V2734 := __e.Get(1)
		_ = V2734
		tmp4391 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp4392 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp4393 := MakeNative(func(__e *ControlFlow) {
					Bytes := __e.Get(1)
					_ = Bytes
					tmp4394 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						__e.TailApply(PrimFunc(symreverse), Bytes)
						return
					}, 1)

					tmp4395 := PrimCloseStream(Stream)

					__e.TailApply(tmp4394, tmp4395)
					return

				}, 1)

				tmp4396 := Call(__e, PrimFunc(symshen_4read_1file_1as_1bytelist_1help), Stream, Byte, Nil)

				__e.TailApply(tmp4393, tmp4396)
				return

			}, 1)

			tmp4397 := PrimReadByte(Stream)

			__e.TailApply(tmp4392, tmp4397)
			return

		}, 1)

		tmp4398 := PrimOpenStream(V2734, symin)

		__e.TailApply(tmp4391, tmp4398)
		return

	}, 1)

	tmp4399 := Call(__e, ns2_1set, symread_1file_1as_1bytelist, tmp4390)

	_ = tmp4399

	tmp4400 := MakeNative(func(__e *ControlFlow) {
		V2735 := __e.Get(1)
		_ = V2735
		V2736 := __e.Get(2)
		_ = V2736
		V2737 := __e.Get(3)
		_ = V2737
		tmp4404 := PrimEqual(MakeNumber(-1), V2736)

		if True == tmp4404 {
			__e.Return(V2737)
			return
		} else {
			tmp4401 := PrimReadByte(V2735)

			tmp4402 := PrimCons(V2736, V2737)

			__e.TailApply(PrimFunc(symshen_4read_1file_1as_1bytelist_1help), V2735, tmp4401, tmp4402)
			return

		}

	}, 3)

	tmp4405 := Call(__e, ns2_1set, symshen_4read_1file_1as_1bytelist_1help, tmp4400)

	_ = tmp4405

	tmp4406 := MakeNative(func(__e *ControlFlow) {
		V2738 := __e.Get(1)
		_ = V2738
		tmp4407 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp4408 := PrimReadByte(Stream)

			__e.TailApply(PrimFunc(symshen_4rfas_1h), Stream, tmp4408, MakeString(""))
			return

		}, 1)

		tmp4409 := PrimOpenStream(V2738, symin)

		__e.TailApply(tmp4407, tmp4409)
		return

	}, 1)

	tmp4410 := Call(__e, ns2_1set, symread_1file_1as_1string, tmp4406)

	_ = tmp4410

	tmp4411 := MakeNative(func(__e *ControlFlow) {
		V2739 := __e.Get(1)
		_ = V2739
		V2740 := __e.Get(2)
		_ = V2740
		V2741 := __e.Get(3)
		_ = V2741
		tmp4417 := PrimEqual(MakeNumber(-1), V2740)

		if True == tmp4417 {
			tmp4412 := PrimCloseStream(V2739)

			_ = tmp4412

			__e.Return(V2741)
			return

		} else {
			tmp4413 := PrimReadByte(V2739)

			tmp4414 := PrimNumberToString(V2740)

			tmp4415 := PrimStringConcat(V2741, tmp4414)

			__e.TailApply(PrimFunc(symshen_4rfas_1h), V2739, tmp4413, tmp4415)
			return

		}

	}, 3)

	tmp4418 := Call(__e, ns2_1set, symshen_4rfas_1h, tmp4411)

	_ = tmp4418

	tmp4419 := MakeNative(func(__e *ControlFlow) {
		V2742 := __e.Get(1)
		_ = V2742
		tmp4420 := Call(__e, PrimFunc(symread), V2742)

		__e.TailApply(PrimFunc(symeval_1kl), tmp4420)
		return

	}, 1)

	tmp4421 := Call(__e, ns2_1set, syminput, tmp4419)

	_ = tmp4421

	tmp4422 := MakeNative(func(__e *ControlFlow) {
		V2743 := __e.Get(1)
		_ = V2743
		V2744 := __e.Get(2)
		_ = V2744
		tmp4423 := MakeNative(func(__e *ControlFlow) {
			Mono_2 := __e.Get(1)
			_ = Mono_2
			tmp4424 := MakeNative(func(__e *ControlFlow) {
				Input := __e.Get(1)
				_ = Input
				tmp4430 := Call(__e, PrimFunc(symshen_4rectify_1type), V2743)

				tmp4431 := Call(__e, PrimFunc(symshen_4typecheck), Input, tmp4430)

				tmp4432 := PrimEqual(False, tmp4431)

				if True == tmp4432 {
					tmp4425 := Call(__e, PrimFunc(symshen_4app), V2743, MakeString("\n"), symshen_4r)

					tmp4426 := PrimStringConcat(MakeString(" is not of type "), tmp4425)

					tmp4427 := Call(__e, PrimFunc(symshen_4app), Input, tmp4426, symshen_4r)

					tmp4428 := PrimStringConcat(MakeString("type error: "), tmp4427)

					__e.Return(PrimSimpleError(tmp4428))
					return

				} else {
					__e.TailApply(PrimFunc(symeval_1kl), Input)
					return
				}

			}, 1)

			tmp4433 := Call(__e, PrimFunc(symread), V2744)

			__e.TailApply(tmp4424, tmp4433)
			return

		}, 1)

		tmp4434 := Call(__e, PrimFunc(symshen_4monotype), V2743)

		__e.TailApply(tmp4423, tmp4434)
		return

	}, 2)

	tmp4435 := Call(__e, ns2_1set, syminput_7, tmp4422)

	_ = tmp4435

	tmp4436 := MakeNative(func(__e *ControlFlow) {
		V2745 := __e.Get(1)
		_ = V2745
		tmp4443 := PrimIsPair(V2745)

		if True == tmp4443 {
			tmp4437 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				__e.TailApply(PrimFunc(symshen_4monotype), Z)
				return
			}, 1)

			__e.TailApply(PrimFunc(symmap), tmp4437, V2745)
			return

		} else {
			tmp4441 := PrimIsVariable(V2745)

			if True == tmp4441 {
				tmp4438 := Call(__e, PrimFunc(symshen_4app), V2745, MakeString("\n"), symshen_4a)

				tmp4439 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp4438)

				__e.Return(PrimSimpleError(tmp4439))
				return

			} else {
				__e.Return(V2745)
				return
			}

		}

	}, 1)

	tmp4444 := Call(__e, ns2_1set, symshen_4monotype, tmp4436)

	_ = tmp4444

	tmp4445 := MakeNative(func(__e *ControlFlow) {
		V2746 := __e.Get(1)
		_ = V2746
		tmp4446 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2746)

		tmp4447 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimFunc(symshen_4return_2), X)
			return
		}, 1)

		__e.TailApply(PrimFunc(symshen_4read_1loop), V2746, tmp4446, Nil, tmp4447)
		return

	}, 1)

	tmp4448 := Call(__e, ns2_1set, symlineread, tmp4445)

	_ = tmp4448

	tmp4449 := MakeNative(func(__e *ControlFlow) {
		V2747 := __e.Get(1)
		_ = V2747
		tmp4450 := MakeNative(func(__e *ControlFlow) {
			Bytelist := __e.Get(1)
			_ = Bytelist
			tmp4451 := MakeNative(func(__e *ControlFlow) {
				S_1exprs := __e.Get(1)
				_ = S_1exprs
				tmp4452 := MakeNative(func(__e *ControlFlow) {
					Process := __e.Get(1)
					_ = Process
					__e.Return(Process)
					return
				}, 1)

				tmp4453 := Call(__e, PrimFunc(symshen_4process_1sexprs), S_1exprs)

				__e.TailApply(tmp4452, tmp4453)
				return

			}, 1)

			tmp4454 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), X)
				return
			}, 1)

			tmp4455 := Call(__e, PrimFunc(symcompile), tmp4454, Bytelist)

			__e.TailApply(tmp4451, tmp4455)
			return

		}, 1)

		tmp4456 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2747)

		__e.TailApply(tmp4450, tmp4456)
		return

	}, 1)

	tmp4457 := Call(__e, ns2_1set, symread_1from_1string, tmp4449)

	_ = tmp4457

	tmp4458 := MakeNative(func(__e *ControlFlow) {
		V2748 := __e.Get(1)
		_ = V2748
		tmp4459 := MakeNative(func(__e *ControlFlow) {
			Bytelist := __e.Get(1)
			_ = Bytelist
			tmp4460 := MakeNative(func(__e *ControlFlow) {
				S_1exprs := __e.Get(1)
				_ = S_1exprs
				__e.Return(S_1exprs)
				return
			}, 1)

			tmp4461 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), X)
				return
			}, 1)

			tmp4462 := Call(__e, PrimFunc(symcompile), tmp4461, Bytelist)

			__e.TailApply(tmp4460, tmp4462)
			return

		}, 1)

		tmp4463 := Call(__e, PrimFunc(symshen_4str_1_6bytes), V2748)

		__e.TailApply(tmp4459, tmp4463)
		return

	}, 1)

	tmp4464 := Call(__e, ns2_1set, symread_1from_1string_1unprocessed, tmp4458)

	_ = tmp4464

	tmp4465 := MakeNative(func(__e *ControlFlow) {
		V2749 := __e.Get(1)
		_ = V2749
		tmp4473 := PrimEqual(MakeString(""), V2749)

		if True == tmp4473 {
			__e.Return(Nil)
			return
		} else {
			tmp4471 := Call(__e, PrimFunc(symshen_4_7string_2), V2749)

			if True == tmp4471 {
				tmp4466 := Call(__e, PrimFunc(symhdstr), V2749)

				tmp4467 := PrimStringToNumber(tmp4466)

				tmp4468 := PrimTailString(V2749)

				tmp4469 := Call(__e, PrimFunc(symshen_4str_1_6bytes), tmp4468)

				__e.Return(PrimCons(tmp4467, tmp4469))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4str_1_6bytes)
				return
			}

		}

	}, 1)

	tmp4474 := Call(__e, ns2_1set, symshen_4str_1_6bytes, tmp4465)

	_ = tmp4474

	tmp4475 := MakeNative(func(__e *ControlFlow) {
		V2750 := __e.Get(1)
		_ = V2750
		tmp4476 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2750)

		tmp4477 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimFunc(symshen_4whitespace_2), X)
			return
		}, 1)

		tmp4478 := Call(__e, PrimFunc(symshen_4read_1loop), V2750, tmp4476, Nil, tmp4477)

		__e.Return(PrimHead(tmp4478))
		return

	}, 1)

	tmp4479 := Call(__e, ns2_1set, symread, tmp4475)

	_ = tmp4479

	tmp4480 := MakeNative(func(__e *ControlFlow) {
		V2751 := __e.Get(1)
		_ = V2751
		tmp4483 := Call(__e, PrimFunc(symshen_4char_1stinput_2), V2751)

		if True == tmp4483 {
			tmp4481 := Call(__e, PrimFunc(symshen_4read_1unit_1string), V2751)

			__e.Return(PrimStringToNumber(tmp4481))
			return

		} else {
			__e.Return(PrimReadByte(V2751))
			return
		}

	}, 1)

	tmp4484 := Call(__e, ns2_1set, symshen_4my_1read_1byte, tmp4480)

	_ = tmp4484

	tmp4485 := MakeNative(func(__e *ControlFlow) {
		V2756 := __e.Get(1)
		_ = V2756
		V2757 := __e.Get(2)
		_ = V2757
		V2758 := __e.Get(3)
		_ = V2758
		V2759 := __e.Get(4)
		_ = V2759
		tmp4508 := PrimEqual(MakeNumber(94), V2757)

		if True == tmp4508 {
			__e.Return(PrimSimpleError(MakeString("read aborted")))
			return
		} else {
			tmp4506 := PrimEqual(MakeNumber(-1), V2757)

			if True == tmp4506 {
				tmp4488 := Call(__e, PrimFunc(symempty_2), V2758)

				if True == tmp4488 {
					__e.Return(PrimSimpleError(MakeString("error: empty stream")))
					return
				} else {
					tmp4486 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), X)
						return
					}, 1)

					__e.TailApply(PrimFunc(symcompile), tmp4486, V2758)
					return

				}

			} else {
				tmp4504 := PrimEqual(MakeNumber(0), V2757)

				if True == tmp4504 {
					tmp4489 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2756)

					__e.TailApply(PrimFunc(symshen_4read_1loop), V2756, tmp4489, V2758, V2759)
					return

				} else {
					tmp4502 := Call(__e, V2759, V2757)

					if True == tmp4502 {
						tmp4490 := MakeNative(func(__e *ControlFlow) {
							Parse := __e.Get(1)
							_ = Parse
							tmp4496 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), Parse)

							if True == tmp4496 {
								tmp4491 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2756)

								tmp4492 := PrimCons(V2757, Nil)

								tmp4493 := Call(__e, PrimFunc(symappend), V2758, tmp4492)

								__e.TailApply(PrimFunc(symshen_4read_1loop), V2756, tmp4491, tmp4493, V2759)
								return

							} else {
								tmp4494 := Call(__e, PrimFunc(symshen_4record_1it), V2758)

								_ = tmp4494

								__e.Return(Parse)
								return

							}

						}, 1)

						tmp4497 := Call(__e, PrimFunc(symshen_4try_1parse), V2758)

						__e.TailApply(tmp4490, tmp4497)
						return

					} else {
						tmp4498 := Call(__e, PrimFunc(symshen_4my_1read_1byte), V2756)

						tmp4499 := PrimCons(V2757, Nil)

						tmp4500 := Call(__e, PrimFunc(symappend), V2758, tmp4499)

						__e.TailApply(PrimFunc(symshen_4read_1loop), V2756, tmp4498, tmp4500, V2759)
						return

					}

				}

			}

		}

	}, 4)

	tmp4509 := Call(__e, ns2_1set, symshen_4read_1loop, tmp4485)

	_ = tmp4509

	tmp4510 := MakeNative(func(__e *ControlFlow) {
		V2760 := __e.Get(1)
		_ = V2760
		tmp4511 := MakeNative(func(__e *ControlFlow) {
			S_1exprs := __e.Get(1)
			_ = S_1exprs
			tmp4513 := Call(__e, PrimFunc(symshen_4nothing_1doing_2), S_1exprs)

			if True == tmp4513 {
				__e.Return(symshen_4i_1failed_b)
				return
			} else {
				__e.TailApply(PrimFunc(symshen_4process_1sexprs), S_1exprs)
				return
			}

		}, 1)

		tmp4514 := MakeNative(func(__e *ControlFlow) {
			tmp4515 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimFunc(symshen_4_5s_1exprs_6), X)
				return
			}, 1)

			__e.TailApply(PrimFunc(symcompile), tmp4515, V2760)
			return

		}, 0)

		tmp4516 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(symshen_4i_1failed_b)
			return
		}, 1)

		tmp4517 := Call(__e, try_1catch, tmp4514, tmp4516)

		__e.TailApply(tmp4511, tmp4517)
		return

	}, 1)

	tmp4518 := Call(__e, ns2_1set, symshen_4try_1parse, tmp4510)

	_ = tmp4518

	tmp4519 := MakeNative(func(__e *ControlFlow) {
		V2763 := __e.Get(1)
		_ = V2763
		tmp4523 := PrimEqual(symshen_4i_1failed_b, V2763)

		if True == tmp4523 {
			__e.Return(True)
			return
		} else {
			tmp4521 := PrimEqual(Nil, V2763)

			if True == tmp4521 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp4524 := Call(__e, ns2_1set, symshen_4nothing_1doing_2, tmp4519)

	_ = tmp4524

	tmp4525 := MakeNative(func(__e *ControlFlow) {
		V2764 := __e.Get(1)
		_ = V2764
		tmp4526 := Call(__e, PrimFunc(symshen_4bytes_1_6string), V2764)

		__e.Return(PrimSet(symshen_4_dit_d, tmp4526))
		return

	}, 1)

	tmp4527 := Call(__e, ns2_1set, symshen_4record_1it, tmp4525)

	_ = tmp4527

	tmp4528 := MakeNative(func(__e *ControlFlow) {
		V2765 := __e.Get(1)
		_ = V2765
		tmp4536 := PrimEqual(Nil, V2765)

		if True == tmp4536 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp4534 := PrimIsPair(V2765)

			if True == tmp4534 {
				tmp4529 := PrimHead(V2765)

				tmp4530 := PrimNumberToString(tmp4529)

				tmp4531 := PrimTail(V2765)

				tmp4532 := Call(__e, PrimFunc(symshen_4bytes_1_6string), tmp4531)

				__e.Return(PrimStringConcat(tmp4530, tmp4532))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4bytes_1_6string)
				return
			}

		}

	}, 1)

	tmp4537 := Call(__e, ns2_1set, symshen_4bytes_1_6string, tmp4528)

	_ = tmp4537

	tmp4538 := MakeNative(func(__e *ControlFlow) {
		V2766 := __e.Get(1)
		_ = V2766
		tmp4539 := MakeNative(func(__e *ControlFlow) {
			Unpack_eExpand := __e.Get(1)
			_ = Unpack_eExpand
			tmp4540 := MakeNative(func(__e *ControlFlow) {
				FindArities := __e.Get(1)
				_ = FindArities
				tmp4541 := MakeNative(func(__e *ControlFlow) {
					Types := __e.Get(1)
					_ = Types
					tmp4542 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimFunc(symshen_4process_1applications), X, Types)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp4542, Unpack_eExpand)
					return

				}, 1)

				tmp4543 := Call(__e, PrimFunc(symshen_4find_1types), Unpack_eExpand)

				__e.TailApply(tmp4541, tmp4543)
				return

			}, 1)

			tmp4544 := Call(__e, PrimFunc(symshen_4find_1arities), Unpack_eExpand)

			__e.TailApply(tmp4540, tmp4544)
			return

		}, 1)

		tmp4545 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), V2766)

		__e.TailApply(tmp4539, tmp4545)
		return

	}, 1)

	tmp4546 := Call(__e, ns2_1set, symshen_4process_1sexprs, tmp4538)

	_ = tmp4546

	tmp4547 := MakeNative(func(__e *ControlFlow) {
		V2767 := __e.Get(1)
		_ = V2767
		tmp4569 := PrimIsPair(V2767)

		var ifres4560 Obj

		if True == tmp4569 {
			tmp4567 := PrimTail(V2767)

			tmp4568 := PrimIsPair(tmp4567)

			var ifres4562 Obj

			if True == tmp4568 {
				tmp4564 := PrimHead(V2767)

				tmp4565 := PrimIntern(MakeString(":"))

				tmp4566 := PrimEqual(tmp4564, tmp4565)

				var ifres4563 Obj

				if True == tmp4566 {
					ifres4563 = True

				} else {
					ifres4563 = False

				}

				ifres4562 = ifres4563

			} else {
				ifres4562 = False

			}

			var ifres4561 Obj

			if True == ifres4562 {
				ifres4561 = True

			} else {
				ifres4561 = False

			}

			ifres4560 = ifres4561

		} else {
			ifres4560 = False

		}

		if True == ifres4560 {
			tmp4548 := PrimTail(V2767)

			tmp4549 := PrimHead(tmp4548)

			tmp4550 := PrimTail(V2767)

			tmp4551 := PrimTail(tmp4550)

			tmp4552 := Call(__e, PrimFunc(symshen_4find_1types), tmp4551)

			__e.Return(PrimCons(tmp4549, tmp4552))
			return

		} else {
			tmp4558 := PrimIsPair(V2767)

			if True == tmp4558 {
				tmp4553 := PrimHead(V2767)

				tmp4554 := Call(__e, PrimFunc(symshen_4find_1types), tmp4553)

				tmp4555 := PrimTail(V2767)

				tmp4556 := Call(__e, PrimFunc(symshen_4find_1types), tmp4555)

				__e.TailApply(PrimFunc(symappend), tmp4554, tmp4556)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp4570 := Call(__e, ns2_1set, symshen_4find_1types, tmp4547)

	_ = tmp4570

	tmp4571 := MakeNative(func(__e *ControlFlow) {
		V2770 := __e.Get(1)
		_ = V2770
		tmp4620 := PrimIsPair(V2770)

		var ifres4601 Obj

		if True == tmp4620 {
			tmp4618 := PrimHead(V2770)

			tmp4619 := PrimEqual(symdefine, tmp4618)

			var ifres4603 Obj

			if True == tmp4619 {
				tmp4616 := PrimTail(V2770)

				tmp4617 := PrimIsPair(tmp4616)

				var ifres4605 Obj

				if True == tmp4617 {
					tmp4613 := PrimTail(V2770)

					tmp4614 := PrimTail(tmp4613)

					tmp4615 := PrimIsPair(tmp4614)

					var ifres4607 Obj

					if True == tmp4615 {
						tmp4609 := PrimTail(V2770)

						tmp4610 := PrimTail(tmp4609)

						tmp4611 := PrimHead(tmp4610)

						tmp4612 := PrimEqual(sym_i, tmp4611)

						var ifres4608 Obj

						if True == tmp4612 {
							ifres4608 = True

						} else {
							ifres4608 = False

						}

						ifres4607 = ifres4608

					} else {
						ifres4607 = False

					}

					var ifres4606 Obj

					if True == ifres4607 {
						ifres4606 = True

					} else {
						ifres4606 = False

					}

					ifres4605 = ifres4606

				} else {
					ifres4605 = False

				}

				var ifres4604 Obj

				if True == ifres4605 {
					ifres4604 = True

				} else {
					ifres4604 = False

				}

				ifres4603 = ifres4604

			} else {
				ifres4603 = False

			}

			var ifres4602 Obj

			if True == ifres4603 {
				ifres4602 = True

			} else {
				ifres4602 = False

			}

			ifres4601 = ifres4602

		} else {
			ifres4601 = False

		}

		if True == ifres4601 {
			tmp4572 := PrimTail(V2770)

			tmp4573 := PrimHead(tmp4572)

			tmp4574 := PrimTail(V2770)

			tmp4575 := PrimHead(tmp4574)

			tmp4576 := PrimTail(V2770)

			tmp4577 := PrimTail(tmp4576)

			tmp4578 := PrimTail(tmp4577)

			tmp4579 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4575, MakeNumber(1), tmp4578)

			__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4573, tmp4579)
			return

		} else {
			tmp4599 := PrimIsPair(V2770)

			var ifres4591 Obj

			if True == tmp4599 {
				tmp4597 := PrimHead(V2770)

				tmp4598 := PrimEqual(symdefine, tmp4597)

				var ifres4593 Obj

				if True == tmp4598 {
					tmp4595 := PrimTail(V2770)

					tmp4596 := PrimIsPair(tmp4595)

					var ifres4594 Obj

					if True == tmp4596 {
						ifres4594 = True

					} else {
						ifres4594 = False

					}

					ifres4593 = ifres4594

				} else {
					ifres4593 = False

				}

				var ifres4592 Obj

				if True == ifres4593 {
					ifres4592 = True

				} else {
					ifres4592 = False

				}

				ifres4591 = ifres4592

			} else {
				ifres4591 = False

			}

			if True == ifres4591 {
				tmp4580 := PrimTail(V2770)

				tmp4581 := PrimHead(tmp4580)

				tmp4582 := PrimTail(V2770)

				tmp4583 := PrimHead(tmp4582)

				tmp4584 := PrimTail(V2770)

				tmp4585 := PrimTail(tmp4584)

				tmp4586 := Call(__e, PrimFunc(symshen_4find_1arity), tmp4583, MakeNumber(0), tmp4585)

				__e.TailApply(PrimFunc(symshen_4store_1arity), tmp4581, tmp4586)
				return

			} else {
				tmp4589 := PrimIsPair(V2770)

				if True == tmp4589 {
					tmp4587 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(PrimFunc(symshen_4find_1arities), Z)
						return
					}, 1)

					__e.TailApply(PrimFunc(symmap), tmp4587, V2770)
					return

				} else {
					__e.Return(symshen_4skip)
					return
				}

			}

		}

	}, 1)

	tmp4621 := Call(__e, ns2_1set, symshen_4find_1arities, tmp4571)

	_ = tmp4621

	tmp4622 := MakeNative(func(__e *ControlFlow) {
		V2771 := __e.Get(1)
		_ = V2771
		V2772 := __e.Get(2)
		_ = V2772
		tmp4623 := MakeNative(func(__e *ControlFlow) {
			ArityF := __e.Get(1)
			_ = ArityF
			tmp4631 := PrimEqual(ArityF, MakeNumber(-1))

			if True == tmp4631 {
				__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2771, V2772)
				return
			} else {
				tmp4629 := PrimEqual(ArityF, V2772)

				if True == tmp4629 {
					__e.Return(symshen_4skip)
					return
				} else {
					tmp4624 := Call(__e, PrimFunc(symshen_4app), V2771, MakeString(" may cause errors\n"), symshen_4a)

					tmp4625 := PrimStringConcat(MakeString("changing the arity of "), tmp4624)

					tmp4626 := Call(__e, PrimFunc(symstoutput))

					tmp4627 := Call(__e, PrimFunc(sympr), tmp4625, tmp4626)

					_ = tmp4627

					__e.TailApply(PrimFunc(symshen_4execute_1store_1arity), V2771, V2772)
					return

				}

			}

		}, 1)

		tmp4632 := Call(__e, PrimFunc(symarity), V2771)

		__e.TailApply(tmp4623, tmp4632)
		return

	}, 2)

	tmp4633 := Call(__e, ns2_1set, symshen_4store_1arity, tmp4622)

	_ = tmp4633

	tmp4634 := MakeNative(func(__e *ControlFlow) {
		V2773 := __e.Get(1)
		_ = V2773
		V2774 := __e.Get(2)
		_ = V2774
		tmp4639 := PrimEqual(MakeNumber(0), V2774)

		if True == tmp4639 {
			tmp4635 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symput), V2773, symarity, MakeNumber(0), tmp4635)
			return

		} else {
			tmp4636 := PrimValue(sym_dproperty_1vector_d)

			tmp4637 := Call(__e, PrimFunc(symput), V2773, symarity, V2774, tmp4636)

			_ = tmp4637

			__e.TailApply(PrimFunc(symshen_4update_1lambdatable), V2773, V2774)
			return

		}

	}, 2)

	tmp4640 := Call(__e, ns2_1set, symshen_4execute_1store_1arity, tmp4634)

	_ = tmp4640

	tmp4641 := MakeNative(func(__e *ControlFlow) {
		V2775 := __e.Get(1)
		_ = V2775
		V2776 := __e.Get(2)
		_ = V2776
		tmp4642 := MakeNative(func(__e *ControlFlow) {
			LambdaTable := __e.Get(1)
			_ = LambdaTable
			tmp4643 := MakeNative(func(__e *ControlFlow) {
				Lambda := __e.Get(1)
				_ = Lambda
				tmp4644 := MakeNative(func(__e *ControlFlow) {
					Insert := __e.Get(1)
					_ = Insert
					tmp4645 := MakeNative(func(__e *ControlFlow) {
						Reset := __e.Get(1)
						_ = Reset
						__e.Return(Reset)
						return
					}, 1)

					tmp4646 := PrimSet(symshen_4_dlambdatable_d, Insert)

					__e.TailApply(tmp4645, tmp4646)
					return

				}, 1)

				tmp4647 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2775, Lambda, LambdaTable)

				__e.TailApply(tmp4644, tmp4647)
				return

			}, 1)

			tmp4648 := PrimCons(V2775, Nil)

			tmp4649 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp4648, V2776)

			tmp4650 := Call(__e, PrimFunc(symeval_1kl), tmp4649)

			__e.TailApply(tmp4643, tmp4650)
			return

		}, 1)

		tmp4651 := PrimValue(symshen_4_dlambdatable_d)

		__e.TailApply(tmp4642, tmp4651)
		return

	}, 2)

	tmp4652 := Call(__e, ns2_1set, symshen_4update_1lambdatable, tmp4641)

	_ = tmp4652

	tmp4653 := MakeNative(func(__e *ControlFlow) {
		V2779 := __e.Get(1)
		_ = V2779
		V2780 := __e.Get(2)
		_ = V2780
		tmp4673 := PrimEqual(MakeNumber(0), V2780)

		if True == tmp4673 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp4671 := PrimEqual(MakeNumber(1), V2780)

			if True == tmp4671 {
				tmp4654 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp4655 := PrimCons(X, Nil)

					tmp4656 := Call(__e, PrimFunc(symappend), V2779, tmp4655)

					tmp4657 := PrimCons(tmp4656, Nil)

					tmp4658 := PrimCons(X, tmp4657)

					__e.Return(PrimCons(symlambda, tmp4658))
					return

				}, 1)

				tmp4659 := Call(__e, PrimFunc(symgensym), symY)

				tmp4660 := Call(__e, PrimFunc(symprotect), tmp4659)

				__e.TailApply(tmp4654, tmp4660)
				return

			} else {
				tmp4661 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp4662 := PrimCons(X, Nil)

					tmp4663 := Call(__e, PrimFunc(symappend), V2779, tmp4662)

					tmp4664 := PrimNumberSubtract(V2780, MakeNumber(1))

					tmp4665 := Call(__e, PrimFunc(symshen_4lambda_1function), tmp4663, tmp4664)

					tmp4666 := PrimCons(tmp4665, Nil)

					tmp4667 := PrimCons(X, tmp4666)

					__e.Return(PrimCons(symlambda, tmp4667))
					return

				}, 1)

				tmp4668 := Call(__e, PrimFunc(symgensym), symY)

				tmp4669 := Call(__e, PrimFunc(symprotect), tmp4668)

				__e.TailApply(tmp4661, tmp4669)
				return

			}

		}

	}, 2)

	tmp4674 := Call(__e, ns2_1set, symshen_4lambda_1function, tmp4653)

	_ = tmp4674

	tmp4675 := MakeNative(func(__e *ControlFlow) {
		V2790 := __e.Get(1)
		_ = V2790
		V2791 := __e.Get(2)
		_ = V2791
		V2792 := __e.Get(3)
		_ = V2792
		tmp4698 := PrimEqual(Nil, V2792)

		if True == tmp4698 {
			tmp4676 := PrimCons(V2790, V2791)

			__e.Return(PrimCons(tmp4676, Nil))
			return

		} else {
			tmp4696 := PrimIsPair(V2792)

			var ifres4687 Obj

			if True == tmp4696 {
				tmp4694 := PrimHead(V2792)

				tmp4695 := PrimIsPair(tmp4694)

				var ifres4689 Obj

				if True == tmp4695 {
					tmp4691 := PrimHead(V2792)

					tmp4692 := PrimHead(tmp4691)

					tmp4693 := PrimEqual(V2790, tmp4692)

					var ifres4690 Obj

					if True == tmp4693 {
						ifres4690 = True

					} else {
						ifres4690 = False

					}

					ifres4689 = ifres4690

				} else {
					ifres4689 = False

				}

				var ifres4688 Obj

				if True == ifres4689 {
					ifres4688 = True

				} else {
					ifres4688 = False

				}

				ifres4687 = ifres4688

			} else {
				ifres4687 = False

			}

			if True == ifres4687 {
				tmp4677 := PrimHead(V2792)

				tmp4678 := PrimHead(tmp4677)

				tmp4679 := PrimCons(tmp4678, V2791)

				tmp4680 := PrimTail(V2792)

				__e.Return(PrimCons(tmp4679, tmp4680))
				return

			} else {
				tmp4685 := PrimIsPair(V2792)

				if True == tmp4685 {
					tmp4681 := PrimHead(V2792)

					tmp4682 := PrimTail(V2792)

					tmp4683 := Call(__e, PrimFunc(symshen_4assoc_1_6), V2790, V2791, tmp4682)

					__e.Return(PrimCons(tmp4681, tmp4683))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.assoc->")))
					return
				}

			}

		}

	}, 3)

	tmp4699 := Call(__e, ns2_1set, symshen_4assoc_1_6, tmp4675)

	_ = tmp4699

	tmp4700 := MakeNative(func(__e *ControlFlow) {
		V2807 := __e.Get(1)
		_ = V2807
		V2808 := __e.Get(2)
		_ = V2808
		V2809 := __e.Get(3)
		_ = V2809
		tmp4747 := PrimEqual(MakeNumber(0), V2808)

		var ifres4740 Obj

		if True == tmp4747 {
			tmp4746 := PrimIsPair(V2809)

			var ifres4742 Obj

			if True == tmp4746 {
				tmp4744 := PrimHead(V2809)

				tmp4745 := PrimEqual(tmp4744, sym_1_6)

				var ifres4743 Obj

				if True == tmp4745 {
					ifres4743 = True

				} else {
					ifres4743 = False

				}

				ifres4742 = ifres4743

			} else {
				ifres4742 = False

			}

			var ifres4741 Obj

			if True == ifres4742 {
				ifres4741 = True

			} else {
				ifres4741 = False

			}

			ifres4740 = ifres4741

		} else {
			ifres4740 = False

		}

		if True == ifres4740 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp4738 := PrimEqual(MakeNumber(0), V2808)

			var ifres4731 Obj

			if True == tmp4738 {
				tmp4737 := PrimIsPair(V2809)

				var ifres4733 Obj

				if True == tmp4737 {
					tmp4735 := PrimHead(V2809)

					tmp4736 := PrimEqual(tmp4735, sym_5_1)

					var ifres4734 Obj

					if True == tmp4736 {
						ifres4734 = True

					} else {
						ifres4734 = False

					}

					ifres4733 = ifres4734

				} else {
					ifres4733 = False

				}

				var ifres4732 Obj

				if True == ifres4733 {
					ifres4732 = True

				} else {
					ifres4732 = False

				}

				ifres4731 = ifres4732

			} else {
				ifres4731 = False

			}

			if True == ifres4731 {
				__e.Return(MakeNumber(0))
				return
			} else {
				tmp4729 := PrimEqual(MakeNumber(0), V2808)

				var ifres4726 Obj

				if True == tmp4729 {
					tmp4728 := PrimIsPair(V2809)

					var ifres4727 Obj

					if True == tmp4728 {
						ifres4727 = True

					} else {
						ifres4727 = False

					}

					ifres4726 = ifres4727

				} else {
					ifres4726 = False

				}

				if True == ifres4726 {
					tmp4701 := PrimTail(V2809)

					tmp4702 := Call(__e, PrimFunc(symshen_4find_1arity), V2807, MakeNumber(0), tmp4701)

					__e.Return(PrimNumberAdd(MakeNumber(1), tmp4702))
					return

				} else {
					tmp4724 := PrimEqual(MakeNumber(1), V2808)

					var ifres4717 Obj

					if True == tmp4724 {
						tmp4723 := PrimIsPair(V2809)

						var ifres4719 Obj

						if True == tmp4723 {
							tmp4721 := PrimHead(V2809)

							tmp4722 := PrimEqual(sym_j, tmp4721)

							var ifres4720 Obj

							if True == tmp4722 {
								ifres4720 = True

							} else {
								ifres4720 = False

							}

							ifres4719 = ifres4720

						} else {
							ifres4719 = False

						}

						var ifres4718 Obj

						if True == ifres4719 {
							ifres4718 = True

						} else {
							ifres4718 = False

						}

						ifres4717 = ifres4718

					} else {
						ifres4717 = False

					}

					if True == ifres4717 {
						tmp4703 := PrimTail(V2809)

						__e.TailApply(PrimFunc(symshen_4find_1arity), V2807, MakeNumber(0), tmp4703)
						return

					} else {
						tmp4715 := PrimEqual(MakeNumber(1), V2808)

						var ifres4712 Obj

						if True == tmp4715 {
							tmp4714 := PrimIsPair(V2809)

							var ifres4713 Obj

							if True == tmp4714 {
								ifres4713 = True

							} else {
								ifres4713 = False

							}

							ifres4712 = ifres4713

						} else {
							ifres4712 = False

						}

						if True == ifres4712 {
							tmp4704 := PrimTail(V2809)

							__e.TailApply(PrimFunc(symshen_4find_1arity), V2807, MakeNumber(1), tmp4704)
							return

						} else {
							tmp4710 := PrimEqual(MakeNumber(1), V2808)

							if True == tmp4710 {
								tmp4705 := Call(__e, PrimFunc(symshen_4app), V2807, MakeString(" definition: missing }\n"), symshen_4a)

								tmp4706 := PrimStringConcat(MakeString("syntax error in "), tmp4705)

								__e.Return(PrimSimpleError(tmp4706))
								return

							} else {
								tmp4707 := Call(__e, PrimFunc(symshen_4app), V2807, MakeString(" definition: missing -> or <-\n"), symshen_4a)

								tmp4708 := PrimStringConcat(MakeString("syntax error in "), tmp4707)

								__e.Return(PrimSimpleError(tmp4708))
								return

							}

						}

					}

				}

			}

		}

	}, 3)

	tmp4748 := Call(__e, ns2_1set, symshen_4find_1arity, tmp4700)

	_ = tmp4748

	tmp4749 := MakeNative(func(__e *ControlFlow) {
		V2810 := __e.Get(1)
		_ = V2810
		tmp4750 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4942 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp4942 {
				tmp4751 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4919 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp4919 {
						tmp4752 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp4905 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

							if True == tmp4905 {
								tmp4753 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp4891 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

									if True == tmp4891 {
										tmp4754 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp4877 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

											if True == tmp4877 {
												tmp4755 := MakeNative(func(__e *ControlFlow) {
													Result := __e.Get(1)
													_ = Result
													tmp4862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

													if True == tmp4862 {
														tmp4756 := MakeNative(func(__e *ControlFlow) {
															Result := __e.Get(1)
															_ = Result
															tmp4843 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

															if True == tmp4843 {
																tmp4757 := MakeNative(func(__e *ControlFlow) {
																	Result := __e.Get(1)
																	_ = Result
																	tmp4828 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

																	if True == tmp4828 {
																		tmp4758 := MakeNative(func(__e *ControlFlow) {
																			Result := __e.Get(1)
																			_ = Result
																			tmp4813 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

																			if True == tmp4813 {
																				tmp4759 := MakeNative(func(__e *ControlFlow) {
																					Result := __e.Get(1)
																					_ = Result
																					tmp4800 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

																					if True == tmp4800 {
																						tmp4760 := MakeNative(func(__e *ControlFlow) {
																							Result := __e.Get(1)
																							_ = Result
																							tmp4785 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

																							if True == tmp4785 {
																								tmp4761 := MakeNative(func(__e *ControlFlow) {
																									Result := __e.Get(1)
																									_ = Result
																									tmp4772 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

																									if True == tmp4772 {
																										tmp4762 := MakeNative(func(__e *ControlFlow) {
																											Result := __e.Get(1)
																											_ = Result
																											tmp4764 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

																											if True == tmp4764 {
																												__e.TailApply(PrimFunc(symshen_4parse_1failure))
																												return
																											} else {
																												__e.Return(Result)
																												return
																											}

																										}, 1)

																										tmp4765 := MakeNative(func(__e *ControlFlow) {
																											Parse_5e_6 := __e.Get(1)
																											_ = Parse_5e_6
																											tmp4768 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)

																											if True == tmp4768 {
																												__e.TailApply(PrimFunc(symshen_4parse_1failure))
																												return
																											} else {
																												tmp4766 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)

																												__e.TailApply(PrimFunc(symshen_4comb), tmp4766, Nil)
																												return

																											}

																										}, 1)

																										tmp4769 := Call(__e, PrimFunc(sym_5e_6), V2810)

																										tmp4770 := Call(__e, tmp4765, tmp4769)

																										__e.TailApply(tmp4762, tmp4770)
																										return

																									} else {
																										__e.Return(Result)
																										return
																									}

																								}, 1)

																								tmp4773 := MakeNative(func(__e *ControlFlow) {
																									Parseshen_4_5whitespaces_6 := __e.Get(1)
																									_ = Parseshen_4_5whitespaces_6
																									tmp4781 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5whitespaces_6)

																									if True == tmp4781 {
																										__e.TailApply(PrimFunc(symshen_4parse_1failure))
																										return
																									} else {
																										tmp4774 := MakeNative(func(__e *ControlFlow) {
																											Parseshen_4_5s_1exprs_6 := __e.Get(1)
																											_ = Parseshen_4_5s_1exprs_6
																											tmp4778 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																											if True == tmp4778 {
																												__e.TailApply(PrimFunc(symshen_4parse_1failure))
																												return
																											} else {
																												tmp4775 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																												tmp4776 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																												__e.TailApply(PrimFunc(symshen_4comb), tmp4775, tmp4776)
																												return

																											}

																										}, 1)

																										tmp4779 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5whitespaces_6)

																										__e.TailApply(tmp4774, tmp4779)
																										return

																									}

																								}, 1)

																								tmp4782 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), V2810)

																								tmp4783 := Call(__e, tmp4773, tmp4782)

																								__e.TailApply(tmp4761, tmp4783)
																								return

																							} else {
																								__e.Return(Result)
																								return
																							}

																						}, 1)

																						tmp4786 := MakeNative(func(__e *ControlFlow) {
																							Parseshen_4_5atom_6 := __e.Get(1)
																							_ = Parseshen_4_5atom_6
																							tmp4796 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5atom_6)

																							if True == tmp4796 {
																								__e.TailApply(PrimFunc(symshen_4parse_1failure))
																								return
																							} else {
																								tmp4787 := MakeNative(func(__e *ControlFlow) {
																									Parseshen_4_5s_1exprs_6 := __e.Get(1)
																									_ = Parseshen_4_5s_1exprs_6
																									tmp4793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																									if True == tmp4793 {
																										__e.TailApply(PrimFunc(symshen_4parse_1failure))
																										return
																									} else {
																										tmp4788 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																										tmp4789 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5atom_6)

																										tmp4790 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																										tmp4791 := PrimCons(tmp4789, tmp4790)

																										__e.TailApply(PrimFunc(symshen_4comb), tmp4788, tmp4791)
																										return

																									}

																								}, 1)

																								tmp4794 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5atom_6)

																								__e.TailApply(tmp4787, tmp4794)
																								return

																							}

																						}, 1)

																						tmp4797 := Call(__e, PrimFunc(symshen_4_5atom_6), V2810)

																						tmp4798 := Call(__e, tmp4786, tmp4797)

																						__e.TailApply(tmp4760, tmp4798)
																						return

																					} else {
																						__e.Return(Result)
																						return
																					}

																				}, 1)

																				tmp4801 := MakeNative(func(__e *ControlFlow) {
																					Parseshen_4_5comment_6 := __e.Get(1)
																					_ = Parseshen_4_5comment_6
																					tmp4809 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5comment_6)

																					if True == tmp4809 {
																						__e.TailApply(PrimFunc(symshen_4parse_1failure))
																						return
																					} else {
																						tmp4802 := MakeNative(func(__e *ControlFlow) {
																							Parseshen_4_5s_1exprs_6 := __e.Get(1)
																							_ = Parseshen_4_5s_1exprs_6
																							tmp4806 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																							if True == tmp4806 {
																								__e.TailApply(PrimFunc(symshen_4parse_1failure))
																								return
																							} else {
																								tmp4803 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																								tmp4804 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																								__e.TailApply(PrimFunc(symshen_4comb), tmp4803, tmp4804)
																								return

																							}

																						}, 1)

																						tmp4807 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5comment_6)

																						__e.TailApply(tmp4802, tmp4807)
																						return

																					}

																				}, 1)

																				tmp4810 := Call(__e, PrimFunc(symshen_4_5comment_6), V2810)

																				tmp4811 := Call(__e, tmp4801, tmp4810)

																				__e.TailApply(tmp4759, tmp4811)
																				return

																			} else {
																				__e.Return(Result)
																				return
																			}

																		}, 1)

																		tmp4814 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5comma_6 := __e.Get(1)
																			_ = Parseshen_4_5comma_6
																			tmp4824 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5comma_6)

																			if True == tmp4824 {
																				__e.TailApply(PrimFunc(symshen_4parse_1failure))
																				return
																			} else {
																				tmp4815 := MakeNative(func(__e *ControlFlow) {
																					Parseshen_4_5s_1exprs_6 := __e.Get(1)
																					_ = Parseshen_4_5s_1exprs_6
																					tmp4821 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																					if True == tmp4821 {
																						__e.TailApply(PrimFunc(symshen_4parse_1failure))
																						return
																					} else {
																						tmp4816 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																						tmp4817 := PrimIntern(MakeString(","))

																						tmp4818 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																						tmp4819 := PrimCons(tmp4817, tmp4818)

																						__e.TailApply(PrimFunc(symshen_4comb), tmp4816, tmp4819)
																						return

																					}

																				}, 1)

																				tmp4822 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5comma_6)

																				__e.TailApply(tmp4815, tmp4822)
																				return

																			}

																		}, 1)

																		tmp4825 := Call(__e, PrimFunc(symshen_4_5comma_6), V2810)

																		tmp4826 := Call(__e, tmp4814, tmp4825)

																		__e.TailApply(tmp4758, tmp4826)
																		return

																	} else {
																		__e.Return(Result)
																		return
																	}

																}, 1)

																tmp4829 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5colon_6 := __e.Get(1)
																	_ = Parseshen_4_5colon_6
																	tmp4839 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5colon_6)

																	if True == tmp4839 {
																		__e.TailApply(PrimFunc(symshen_4parse_1failure))
																		return
																	} else {
																		tmp4830 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5s_1exprs_6 := __e.Get(1)
																			_ = Parseshen_4_5s_1exprs_6
																			tmp4836 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																			if True == tmp4836 {
																				__e.TailApply(PrimFunc(symshen_4parse_1failure))
																				return
																			} else {
																				tmp4831 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																				tmp4832 := PrimIntern(MakeString(":"))

																				tmp4833 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																				tmp4834 := PrimCons(tmp4832, tmp4833)

																				__e.TailApply(PrimFunc(symshen_4comb), tmp4831, tmp4834)
																				return

																			}

																		}, 1)

																		tmp4837 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5colon_6)

																		__e.TailApply(tmp4830, tmp4837)
																		return

																	}

																}, 1)

																tmp4840 := Call(__e, PrimFunc(symshen_4_5colon_6), V2810)

																tmp4841 := Call(__e, tmp4829, tmp4840)

																__e.TailApply(tmp4757, tmp4841)
																return

															} else {
																__e.Return(Result)
																return
															}

														}, 1)

														tmp4844 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5colon_6 := __e.Get(1)
															_ = Parseshen_4_5colon_6
															tmp4858 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5colon_6)

															if True == tmp4858 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																tmp4845 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5equal_6 := __e.Get(1)
																	_ = Parseshen_4_5equal_6
																	tmp4855 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5equal_6)

																	if True == tmp4855 {
																		__e.TailApply(PrimFunc(symshen_4parse_1failure))
																		return
																	} else {
																		tmp4846 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5s_1exprs_6 := __e.Get(1)
																			_ = Parseshen_4_5s_1exprs_6
																			tmp4852 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																			if True == tmp4852 {
																				__e.TailApply(PrimFunc(symshen_4parse_1failure))
																				return
																			} else {
																				tmp4847 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																				tmp4848 := PrimIntern(MakeString(":="))

																				tmp4849 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																				tmp4850 := PrimCons(tmp4848, tmp4849)

																				__e.TailApply(PrimFunc(symshen_4comb), tmp4847, tmp4850)
																				return

																			}

																		}, 1)

																		tmp4853 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5equal_6)

																		__e.TailApply(tmp4846, tmp4853)
																		return

																	}

																}, 1)

																tmp4856 := Call(__e, PrimFunc(symshen_4_5equal_6), Parseshen_4_5colon_6)

																__e.TailApply(tmp4845, tmp4856)
																return

															}

														}, 1)

														tmp4859 := Call(__e, PrimFunc(symshen_4_5colon_6), V2810)

														tmp4860 := Call(__e, tmp4844, tmp4859)

														__e.TailApply(tmp4756, tmp4860)
														return

													} else {
														__e.Return(Result)
														return
													}

												}, 1)

												tmp4863 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5semicolon_6 := __e.Get(1)
													_ = Parseshen_4_5semicolon_6
													tmp4873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5semicolon_6)

													if True == tmp4873 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp4864 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5s_1exprs_6 := __e.Get(1)
															_ = Parseshen_4_5s_1exprs_6
															tmp4870 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

															if True == tmp4870 {
																__e.TailApply(PrimFunc(symshen_4parse_1failure))
																return
															} else {
																tmp4865 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																tmp4866 := PrimIntern(MakeString(";"))

																tmp4867 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																tmp4868 := PrimCons(tmp4866, tmp4867)

																__e.TailApply(PrimFunc(symshen_4comb), tmp4865, tmp4868)
																return

															}

														}, 1)

														tmp4871 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5semicolon_6)

														__e.TailApply(tmp4864, tmp4871)
														return

													}

												}, 1)

												tmp4874 := Call(__e, PrimFunc(symshen_4_5semicolon_6), V2810)

												tmp4875 := Call(__e, tmp4863, tmp4874)

												__e.TailApply(tmp4755, tmp4875)
												return

											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp4878 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5bar_6 := __e.Get(1)
											_ = Parseshen_4_5bar_6
											tmp4887 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5bar_6)

											if True == tmp4887 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp4879 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5s_1exprs_6 := __e.Get(1)
													_ = Parseshen_4_5s_1exprs_6
													tmp4884 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

													if True == tmp4884 {
														__e.TailApply(PrimFunc(symshen_4parse_1failure))
														return
													} else {
														tmp4880 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

														tmp4881 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

														tmp4882 := PrimCons(symbar_b, tmp4881)

														__e.TailApply(PrimFunc(symshen_4comb), tmp4880, tmp4882)
														return

													}

												}, 1)

												tmp4885 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5bar_6)

												__e.TailApply(tmp4879, tmp4885)
												return

											}

										}, 1)

										tmp4888 := Call(__e, PrimFunc(symshen_4_5bar_6), V2810)

										tmp4889 := Call(__e, tmp4878, tmp4888)

										__e.TailApply(tmp4754, tmp4889)
										return

									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp4892 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5rcurly_6 := __e.Get(1)
									_ = Parseshen_4_5rcurly_6
									tmp4901 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rcurly_6)

									if True == tmp4901 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4893 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5s_1exprs_6 := __e.Get(1)
											_ = Parseshen_4_5s_1exprs_6
											tmp4898 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

											if True == tmp4898 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp4894 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

												tmp4895 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

												tmp4896 := PrimCons(sym_j, tmp4895)

												__e.TailApply(PrimFunc(symshen_4comb), tmp4894, tmp4896)
												return

											}

										}, 1)

										tmp4899 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5rcurly_6)

										__e.TailApply(tmp4893, tmp4899)
										return

									}

								}, 1)

								tmp4902 := Call(__e, PrimFunc(symshen_4_5rcurly_6), V2810)

								tmp4903 := Call(__e, tmp4892, tmp4902)

								__e.TailApply(tmp4753, tmp4903)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp4906 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5lcurly_6 := __e.Get(1)
							_ = Parseshen_4_5lcurly_6
							tmp4915 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5lcurly_6)

							if True == tmp4915 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp4907 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5s_1exprs_6 := __e.Get(1)
									_ = Parseshen_4_5s_1exprs_6
									tmp4912 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

									if True == tmp4912 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4908 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

										tmp4909 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

										tmp4910 := PrimCons(sym_i, tmp4909)

										__e.TailApply(PrimFunc(symshen_4comb), tmp4908, tmp4910)
										return

									}

								}, 1)

								tmp4913 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), Parseshen_4_5lcurly_6)

								__e.TailApply(tmp4907, tmp4913)
								return

							}

						}, 1)

						tmp4916 := Call(__e, PrimFunc(symshen_4_5lcurly_6), V2810)

						tmp4917 := Call(__e, tmp4906, tmp4916)

						__e.TailApply(tmp4752, tmp4917)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4920 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5lrb_6 := __e.Get(1)
					_ = Parseshen_4_5lrb_6
					tmp4938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5lrb_6)

					if True == tmp4938 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4921 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5s_1exprs1_6 := __e.Get(1)
							_ = Parseshen_4_5s_1exprs1_6
							tmp4935 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs1_6)

							if True == tmp4935 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp4922 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5rrb_6 := __e.Get(1)
									_ = Parseshen_4_5rrb_6
									tmp4932 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rrb_6)

									if True == tmp4932 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4923 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5s_1exprs2_6 := __e.Get(1)
											_ = Parseshen_4_5s_1exprs2_6
											tmp4929 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs2_6)

											if True == tmp4929 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp4924 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs2_6)

												tmp4925 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs1_6)

												tmp4926 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs2_6)

												tmp4927 := Call(__e, PrimFunc(symshen_4add_1sexpr), tmp4925, tmp4926)

												__e.TailApply(PrimFunc(symshen_4comb), tmp4924, tmp4927)
												return

											}

										}, 1)

										tmp4930 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), Parseshen_4_5rrb_6)

										__e.TailApply(tmp4923, tmp4930)
										return

									}

								}, 1)

								tmp4933 := Call(__e, PrimFunc(symshen_4_5rrb_6), Parseshen_4_5s_1exprs1_6)

								__e.TailApply(tmp4922, tmp4933)
								return

							}

						}, 1)

						tmp4936 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), Parseshen_4_5lrb_6)

						__e.TailApply(tmp4921, tmp4936)
						return

					}

				}, 1)

				tmp4939 := Call(__e, PrimFunc(symshen_4_5lrb_6), V2810)

				tmp4940 := Call(__e, tmp4920, tmp4939)

				__e.TailApply(tmp4751, tmp4940)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4943 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5lsb_6 := __e.Get(1)
			_ = Parseshen_4_5lsb_6
			tmp4962 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5lsb_6)

			if True == tmp4962 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp4944 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5s_1exprs1_6 := __e.Get(1)
					_ = Parseshen_4_5s_1exprs1_6
					tmp4959 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs1_6)

					if True == tmp4959 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp4945 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5rsb_6 := __e.Get(1)
							_ = Parseshen_4_5rsb_6
							tmp4956 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5rsb_6)

							if True == tmp4956 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp4946 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5s_1exprs2_6 := __e.Get(1)
									_ = Parseshen_4_5s_1exprs2_6
									tmp4953 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs2_6)

									if True == tmp4953 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp4947 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs2_6)

										tmp4948 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs1_6)

										tmp4949 := Call(__e, PrimFunc(symshen_4cons_1form), tmp4948)

										tmp4950 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs2_6)

										tmp4951 := PrimCons(tmp4949, tmp4950)

										__e.TailApply(PrimFunc(symshen_4comb), tmp4947, tmp4951)
										return

									}

								}, 1)

								tmp4954 := Call(__e, PrimFunc(symshen_4_5s_1exprs2_6), Parseshen_4_5rsb_6)

								__e.TailApply(tmp4946, tmp4954)
								return

							}

						}, 1)

						tmp4957 := Call(__e, PrimFunc(symshen_4_5rsb_6), Parseshen_4_5s_1exprs1_6)

						__e.TailApply(tmp4945, tmp4957)
						return

					}

				}, 1)

				tmp4960 := Call(__e, PrimFunc(symshen_4_5s_1exprs1_6), Parseshen_4_5lsb_6)

				__e.TailApply(tmp4944, tmp4960)
				return

			}

		}, 1)

		tmp4963 := Call(__e, PrimFunc(symshen_4_5lsb_6), V2810)

		tmp4964 := Call(__e, tmp4943, tmp4963)

		__e.TailApply(tmp4750, tmp4964)
		return

	}, 1)

	tmp4965 := Call(__e, ns2_1set, symshen_4_5s_1exprs_6, tmp4749)

	_ = tmp4965

	tmp4966 := MakeNative(func(__e *ControlFlow) {
		V2811 := __e.Get(1)
		_ = V2811
		V2812 := __e.Get(2)
		_ = V2812
		tmp4984 := PrimIsPair(V2811)

		var ifres4971 Obj

		if True == tmp4984 {
			tmp4982 := PrimHead(V2811)

			tmp4983 := PrimEqual(sym_3, tmp4982)

			var ifres4973 Obj

			if True == tmp4983 {
				tmp4980 := PrimTail(V2811)

				tmp4981 := PrimIsPair(tmp4980)

				var ifres4975 Obj

				if True == tmp4981 {
					tmp4977 := PrimTail(V2811)

					tmp4978 := PrimTail(tmp4977)

					tmp4979 := PrimEqual(Nil, tmp4978)

					var ifres4976 Obj

					if True == tmp4979 {
						ifres4976 = True

					} else {
						ifres4976 = False

					}

					ifres4975 = ifres4976

				} else {
					ifres4975 = False

				}

				var ifres4974 Obj

				if True == ifres4975 {
					ifres4974 = True

				} else {
					ifres4974 = False

				}

				ifres4973 = ifres4974

			} else {
				ifres4973 = False

			}

			var ifres4972 Obj

			if True == ifres4973 {
				ifres4972 = True

			} else {
				ifres4972 = False

			}

			ifres4971 = ifres4972

		} else {
			ifres4971 = False

		}

		if True == ifres4971 {
			tmp4967 := PrimTail(V2811)

			tmp4968 := PrimHead(tmp4967)

			tmp4969 := Call(__e, PrimFunc(symexplode), tmp4968)

			__e.TailApply(PrimFunc(symappend), tmp4969, V2812)
			return

		} else {
			__e.Return(PrimCons(V2811, V2812))
			return
		}

	}, 2)

	tmp4985 := Call(__e, ns2_1set, symshen_4add_1sexpr, tmp4966)

	_ = tmp4985

	tmp4986 := MakeNative(func(__e *ControlFlow) {
		V2813 := __e.Get(1)
		_ = V2813
		tmp4987 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4989 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp4989 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4996 := Call(__e, PrimFunc(symshen_4_ahd_2), V2813, MakeNumber(91))

		var ifres4990 Obj

		if True == tmp4996 {
			tmp4991 := MakeNative(func(__e *ControlFlow) {
				News2495 := __e.Get(1)
				_ = News2495
				tmp4992 := Call(__e, PrimFunc(symshen_4in_1_6), News2495)

				__e.TailApply(PrimFunc(symshen_4comb), tmp4992, symshen_4skip)
				return

			}, 1)

			tmp4993 := Call(__e, PrimFunc(symshen_4tls), V2813)

			tmp4994 := Call(__e, tmp4991, tmp4993)

			ifres4990 = tmp4994

		} else {
			tmp4995 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres4990 = tmp4995

		}

		__e.TailApply(tmp4987, ifres4990)
		return

	}, 1)

	tmp4997 := Call(__e, ns2_1set, symshen_4_5lsb_6, tmp4986)

	_ = tmp4997

	tmp4998 := MakeNative(func(__e *ControlFlow) {
		V2814 := __e.Get(1)
		_ = V2814
		tmp4999 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5001 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5001 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5008 := Call(__e, PrimFunc(symshen_4_ahd_2), V2814, MakeNumber(93))

		var ifres5002 Obj

		if True == tmp5008 {
			tmp5003 := MakeNative(func(__e *ControlFlow) {
				News2497 := __e.Get(1)
				_ = News2497
				tmp5004 := Call(__e, PrimFunc(symshen_4in_1_6), News2497)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5004, symshen_4skip)
				return

			}, 1)

			tmp5005 := Call(__e, PrimFunc(symshen_4tls), V2814)

			tmp5006 := Call(__e, tmp5003, tmp5005)

			ifres5002 = tmp5006

		} else {
			tmp5007 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5002 = tmp5007

		}

		__e.TailApply(tmp4999, ifres5002)
		return

	}, 1)

	tmp5009 := Call(__e, ns2_1set, symshen_4_5rsb_6, tmp4998)

	_ = tmp5009

	tmp5010 := MakeNative(func(__e *ControlFlow) {
		V2815 := __e.Get(1)
		_ = V2815
		tmp5011 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5013 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5013 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5014 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5s_1exprs_6 := __e.Get(1)
			_ = Parseshen_4_5s_1exprs_6
			tmp5018 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

			if True == tmp5018 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5015 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

				tmp5016 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5015, tmp5016)
				return

			}

		}, 1)

		tmp5019 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2815)

		tmp5020 := Call(__e, tmp5014, tmp5019)

		__e.TailApply(tmp5011, tmp5020)
		return

	}, 1)

	tmp5021 := Call(__e, ns2_1set, symshen_4_5s_1exprs1_6, tmp5010)

	_ = tmp5021

	tmp5022 := MakeNative(func(__e *ControlFlow) {
		V2816 := __e.Get(1)
		_ = V2816
		tmp5023 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5025 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5025 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5026 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5s_1exprs_6 := __e.Get(1)
			_ = Parseshen_4_5s_1exprs_6
			tmp5030 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

			if True == tmp5030 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5027 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

				tmp5028 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5027, tmp5028)
				return

			}

		}, 1)

		tmp5031 := Call(__e, PrimFunc(symshen_4_5s_1exprs_6), V2816)

		tmp5032 := Call(__e, tmp5026, tmp5031)

		__e.TailApply(tmp5023, tmp5032)
		return

	}, 1)

	tmp5033 := Call(__e, ns2_1set, symshen_4_5s_1exprs2_6, tmp5022)

	_ = tmp5033

	tmp5034 := MakeNative(func(__e *ControlFlow) {
		V2818 := __e.Get(1)
		_ = V2818
		tmp5091 := PrimEqual(Nil, V2818)

		if True == tmp5091 {
			__e.Return(Nil)
			return
		} else {
			tmp5089 := PrimIsPair(V2818)

			var ifres5069 Obj

			if True == tmp5089 {
				tmp5087 := PrimTail(V2818)

				tmp5088 := PrimIsPair(tmp5087)

				var ifres5071 Obj

				if True == tmp5088 {
					tmp5084 := PrimTail(V2818)

					tmp5085 := PrimTail(tmp5084)

					tmp5086 := PrimIsPair(tmp5085)

					var ifres5073 Obj

					if True == tmp5086 {
						tmp5080 := PrimTail(V2818)

						tmp5081 := PrimTail(tmp5080)

						tmp5082 := PrimTail(tmp5081)

						tmp5083 := PrimEqual(Nil, tmp5082)

						var ifres5075 Obj

						if True == tmp5083 {
							tmp5077 := PrimTail(V2818)

							tmp5078 := PrimHead(tmp5077)

							tmp5079 := PrimEqual(tmp5078, symbar_b)

							var ifres5076 Obj

							if True == tmp5079 {
								ifres5076 = True

							} else {
								ifres5076 = False

							}

							ifres5075 = ifres5076

						} else {
							ifres5075 = False

						}

						var ifres5074 Obj

						if True == ifres5075 {
							ifres5074 = True

						} else {
							ifres5074 = False

						}

						ifres5073 = ifres5074

					} else {
						ifres5073 = False

					}

					var ifres5072 Obj

					if True == ifres5073 {
						ifres5072 = True

					} else {
						ifres5072 = False

					}

					ifres5071 = ifres5072

				} else {
					ifres5071 = False

				}

				var ifres5070 Obj

				if True == ifres5071 {
					ifres5070 = True

				} else {
					ifres5070 = False

				}

				ifres5069 = ifres5070

			} else {
				ifres5069 = False

			}

			if True == ifres5069 {
				tmp5035 := PrimHead(V2818)

				tmp5036 := PrimTail(V2818)

				tmp5037 := PrimTail(tmp5036)

				tmp5038 := PrimCons(tmp5035, tmp5037)

				__e.Return(PrimCons(symcons, tmp5038))
				return

			} else {
				tmp5067 := PrimIsPair(V2818)

				var ifres5047 Obj

				if True == tmp5067 {
					tmp5065 := PrimTail(V2818)

					tmp5066 := PrimIsPair(tmp5065)

					var ifres5049 Obj

					if True == tmp5066 {
						tmp5062 := PrimTail(V2818)

						tmp5063 := PrimTail(tmp5062)

						tmp5064 := PrimIsPair(tmp5063)

						var ifres5051 Obj

						if True == tmp5064 {
							tmp5058 := PrimTail(V2818)

							tmp5059 := PrimTail(tmp5058)

							tmp5060 := PrimTail(tmp5059)

							tmp5061 := PrimIsPair(tmp5060)

							var ifres5053 Obj

							if True == tmp5061 {
								tmp5055 := PrimTail(V2818)

								tmp5056 := PrimHead(tmp5055)

								tmp5057 := PrimEqual(tmp5056, symbar_b)

								var ifres5054 Obj

								if True == tmp5057 {
									ifres5054 = True

								} else {
									ifres5054 = False

								}

								ifres5053 = ifres5054

							} else {
								ifres5053 = False

							}

							var ifres5052 Obj

							if True == ifres5053 {
								ifres5052 = True

							} else {
								ifres5052 = False

							}

							ifres5051 = ifres5052

						} else {
							ifres5051 = False

						}

						var ifres5050 Obj

						if True == ifres5051 {
							ifres5050 = True

						} else {
							ifres5050 = False

						}

						ifres5049 = ifres5050

					} else {
						ifres5049 = False

					}

					var ifres5048 Obj

					if True == ifres5049 {
						ifres5048 = True

					} else {
						ifres5048 = False

					}

					ifres5047 = ifres5048

				} else {
					ifres5047 = False

				}

				if True == ifres5047 {
					__e.Return(PrimSimpleError(MakeString("misapplication of |\n")))
					return
				} else {
					tmp5045 := PrimIsPair(V2818)

					if True == tmp5045 {
						tmp5039 := PrimHead(V2818)

						tmp5040 := PrimTail(V2818)

						tmp5041 := Call(__e, PrimFunc(symshen_4cons_1form), tmp5040)

						tmp5042 := PrimCons(tmp5041, Nil)

						tmp5043 := PrimCons(tmp5039, tmp5042)

						__e.Return(PrimCons(symcons, tmp5043))
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4cons_1form)
						return
					}

				}

			}

		}

	}, 1)

	tmp5092 := Call(__e, ns2_1set, symshen_4cons_1form, tmp5034)

	_ = tmp5092

	tmp5093 := MakeNative(func(__e *ControlFlow) {
		V2819 := __e.Get(1)
		_ = V2819
		tmp5094 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5096 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5096 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5103 := Call(__e, PrimFunc(symshen_4_ahd_2), V2819, MakeNumber(40))

		var ifres5097 Obj

		if True == tmp5103 {
			tmp5098 := MakeNative(func(__e *ControlFlow) {
				News2501 := __e.Get(1)
				_ = News2501
				tmp5099 := Call(__e, PrimFunc(symshen_4in_1_6), News2501)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5099, symshen_4skip)
				return

			}, 1)

			tmp5100 := Call(__e, PrimFunc(symshen_4tls), V2819)

			tmp5101 := Call(__e, tmp5098, tmp5100)

			ifres5097 = tmp5101

		} else {
			tmp5102 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5097 = tmp5102

		}

		__e.TailApply(tmp5094, ifres5097)
		return

	}, 1)

	tmp5104 := Call(__e, ns2_1set, symshen_4_5lrb_6, tmp5093)

	_ = tmp5104

	tmp5105 := MakeNative(func(__e *ControlFlow) {
		V2820 := __e.Get(1)
		_ = V2820
		tmp5106 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5108 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5108 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5115 := Call(__e, PrimFunc(symshen_4_ahd_2), V2820, MakeNumber(41))

		var ifres5109 Obj

		if True == tmp5115 {
			tmp5110 := MakeNative(func(__e *ControlFlow) {
				News2503 := __e.Get(1)
				_ = News2503
				tmp5111 := Call(__e, PrimFunc(symshen_4in_1_6), News2503)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5111, symshen_4skip)
				return

			}, 1)

			tmp5112 := Call(__e, PrimFunc(symshen_4tls), V2820)

			tmp5113 := Call(__e, tmp5110, tmp5112)

			ifres5109 = tmp5113

		} else {
			tmp5114 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5109 = tmp5114

		}

		__e.TailApply(tmp5106, ifres5109)
		return

	}, 1)

	tmp5116 := Call(__e, ns2_1set, symshen_4_5rrb_6, tmp5105)

	_ = tmp5116

	tmp5117 := MakeNative(func(__e *ControlFlow) {
		V2821 := __e.Get(1)
		_ = V2821
		tmp5118 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5120 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5120 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5127 := Call(__e, PrimFunc(symshen_4_ahd_2), V2821, MakeNumber(123))

		var ifres5121 Obj

		if True == tmp5127 {
			tmp5122 := MakeNative(func(__e *ControlFlow) {
				News2505 := __e.Get(1)
				_ = News2505
				tmp5123 := Call(__e, PrimFunc(symshen_4in_1_6), News2505)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5123, symshen_4skip)
				return

			}, 1)

			tmp5124 := Call(__e, PrimFunc(symshen_4tls), V2821)

			tmp5125 := Call(__e, tmp5122, tmp5124)

			ifres5121 = tmp5125

		} else {
			tmp5126 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5121 = tmp5126

		}

		__e.TailApply(tmp5118, ifres5121)
		return

	}, 1)

	tmp5128 := Call(__e, ns2_1set, symshen_4_5lcurly_6, tmp5117)

	_ = tmp5128

	tmp5129 := MakeNative(func(__e *ControlFlow) {
		V2822 := __e.Get(1)
		_ = V2822
		tmp5130 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5132 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5132 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5139 := Call(__e, PrimFunc(symshen_4_ahd_2), V2822, MakeNumber(125))

		var ifres5133 Obj

		if True == tmp5139 {
			tmp5134 := MakeNative(func(__e *ControlFlow) {
				News2507 := __e.Get(1)
				_ = News2507
				tmp5135 := Call(__e, PrimFunc(symshen_4in_1_6), News2507)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5135, symshen_4skip)
				return

			}, 1)

			tmp5136 := Call(__e, PrimFunc(symshen_4tls), V2822)

			tmp5137 := Call(__e, tmp5134, tmp5136)

			ifres5133 = tmp5137

		} else {
			tmp5138 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5133 = tmp5138

		}

		__e.TailApply(tmp5130, ifres5133)
		return

	}, 1)

	tmp5140 := Call(__e, ns2_1set, symshen_4_5rcurly_6, tmp5129)

	_ = tmp5140

	tmp5141 := MakeNative(func(__e *ControlFlow) {
		V2823 := __e.Get(1)
		_ = V2823
		tmp5142 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5144 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5144 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5151 := Call(__e, PrimFunc(symshen_4_ahd_2), V2823, MakeNumber(124))

		var ifres5145 Obj

		if True == tmp5151 {
			tmp5146 := MakeNative(func(__e *ControlFlow) {
				News2509 := __e.Get(1)
				_ = News2509
				tmp5147 := Call(__e, PrimFunc(symshen_4in_1_6), News2509)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5147, symshen_4skip)
				return

			}, 1)

			tmp5148 := Call(__e, PrimFunc(symshen_4tls), V2823)

			tmp5149 := Call(__e, tmp5146, tmp5148)

			ifres5145 = tmp5149

		} else {
			tmp5150 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5145 = tmp5150

		}

		__e.TailApply(tmp5142, ifres5145)
		return

	}, 1)

	tmp5152 := Call(__e, ns2_1set, symshen_4_5bar_6, tmp5141)

	_ = tmp5152

	tmp5153 := MakeNative(func(__e *ControlFlow) {
		V2824 := __e.Get(1)
		_ = V2824
		tmp5154 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5156 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5156 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5163 := Call(__e, PrimFunc(symshen_4_ahd_2), V2824, MakeNumber(59))

		var ifres5157 Obj

		if True == tmp5163 {
			tmp5158 := MakeNative(func(__e *ControlFlow) {
				News2511 := __e.Get(1)
				_ = News2511
				tmp5159 := Call(__e, PrimFunc(symshen_4in_1_6), News2511)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5159, symshen_4skip)
				return

			}, 1)

			tmp5160 := Call(__e, PrimFunc(symshen_4tls), V2824)

			tmp5161 := Call(__e, tmp5158, tmp5160)

			ifres5157 = tmp5161

		} else {
			tmp5162 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5157 = tmp5162

		}

		__e.TailApply(tmp5154, ifres5157)
		return

	}, 1)

	tmp5164 := Call(__e, ns2_1set, symshen_4_5semicolon_6, tmp5153)

	_ = tmp5164

	tmp5165 := MakeNative(func(__e *ControlFlow) {
		V2825 := __e.Get(1)
		_ = V2825
		tmp5166 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5168 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5168 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5175 := Call(__e, PrimFunc(symshen_4_ahd_2), V2825, MakeNumber(58))

		var ifres5169 Obj

		if True == tmp5175 {
			tmp5170 := MakeNative(func(__e *ControlFlow) {
				News2513 := __e.Get(1)
				_ = News2513
				tmp5171 := Call(__e, PrimFunc(symshen_4in_1_6), News2513)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5171, symshen_4skip)
				return

			}, 1)

			tmp5172 := Call(__e, PrimFunc(symshen_4tls), V2825)

			tmp5173 := Call(__e, tmp5170, tmp5172)

			ifres5169 = tmp5173

		} else {
			tmp5174 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5169 = tmp5174

		}

		__e.TailApply(tmp5166, ifres5169)
		return

	}, 1)

	tmp5176 := Call(__e, ns2_1set, symshen_4_5colon_6, tmp5165)

	_ = tmp5176

	tmp5177 := MakeNative(func(__e *ControlFlow) {
		V2826 := __e.Get(1)
		_ = V2826
		tmp5178 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5180 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5180 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5187 := Call(__e, PrimFunc(symshen_4_ahd_2), V2826, MakeNumber(44))

		var ifres5181 Obj

		if True == tmp5187 {
			tmp5182 := MakeNative(func(__e *ControlFlow) {
				News2515 := __e.Get(1)
				_ = News2515
				tmp5183 := Call(__e, PrimFunc(symshen_4in_1_6), News2515)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5183, symshen_4skip)
				return

			}, 1)

			tmp5184 := Call(__e, PrimFunc(symshen_4tls), V2826)

			tmp5185 := Call(__e, tmp5182, tmp5184)

			ifres5181 = tmp5185

		} else {
			tmp5186 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5181 = tmp5186

		}

		__e.TailApply(tmp5178, ifres5181)
		return

	}, 1)

	tmp5188 := Call(__e, ns2_1set, symshen_4_5comma_6, tmp5177)

	_ = tmp5188

	tmp5189 := MakeNative(func(__e *ControlFlow) {
		V2827 := __e.Get(1)
		_ = V2827
		tmp5190 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5192 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5192 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5199 := Call(__e, PrimFunc(symshen_4_ahd_2), V2827, MakeNumber(61))

		var ifres5193 Obj

		if True == tmp5199 {
			tmp5194 := MakeNative(func(__e *ControlFlow) {
				News2517 := __e.Get(1)
				_ = News2517
				tmp5195 := Call(__e, PrimFunc(symshen_4in_1_6), News2517)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5195, symshen_4skip)
				return

			}, 1)

			tmp5196 := Call(__e, PrimFunc(symshen_4tls), V2827)

			tmp5197 := Call(__e, tmp5194, tmp5196)

			ifres5193 = tmp5197

		} else {
			tmp5198 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5193 = tmp5198

		}

		__e.TailApply(tmp5190, ifres5193)
		return

	}, 1)

	tmp5200 := Call(__e, ns2_1set, symshen_4_5equal_6, tmp5189)

	_ = tmp5200

	tmp5201 := MakeNative(func(__e *ControlFlow) {
		V2828 := __e.Get(1)
		_ = V2828
		tmp5202 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5213 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5213 {
				tmp5203 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5205 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5205 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5206 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5multiline_6 := __e.Get(1)
					_ = Parseshen_4_5multiline_6
					tmp5209 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5multiline_6)

					if True == tmp5209 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5207 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5multiline_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5207, symshen_4skip)
						return

					}

				}, 1)

				tmp5210 := Call(__e, PrimFunc(symshen_4_5multiline_6), V2828)

				tmp5211 := Call(__e, tmp5206, tmp5210)

				__e.TailApply(tmp5203, tmp5211)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5214 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5singleline_6 := __e.Get(1)
			_ = Parseshen_4_5singleline_6
			tmp5217 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5singleline_6)

			if True == tmp5217 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5215 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5singleline_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5215, symshen_4skip)
				return

			}

		}, 1)

		tmp5218 := Call(__e, PrimFunc(symshen_4_5singleline_6), V2828)

		tmp5219 := Call(__e, tmp5214, tmp5218)

		__e.TailApply(tmp5202, tmp5219)
		return

	}, 1)

	tmp5220 := Call(__e, ns2_1set, symshen_4_5comment_6, tmp5201)

	_ = tmp5220

	tmp5221 := MakeNative(func(__e *ControlFlow) {
		V2829 := __e.Get(1)
		_ = V2829
		tmp5222 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5224 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5224 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5225 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5backslash_6 := __e.Get(1)
			_ = Parseshen_4_5backslash_6
			tmp5240 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

			if True == tmp5240 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5226 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5backslash_6 := __e.Get(1)
					_ = Parseshen_4_5backslash_6
					tmp5237 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

					if True == tmp5237 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5227 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5shortnatters_6 := __e.Get(1)
							_ = Parseshen_4_5shortnatters_6
							tmp5234 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5shortnatters_6)

							if True == tmp5234 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5228 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5returns_6 := __e.Get(1)
									_ = Parseshen_4_5returns_6
									tmp5231 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5returns_6)

									if True == tmp5231 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp5229 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5returns_6)

										__e.TailApply(PrimFunc(symshen_4comb), tmp5229, symshen_4skip)
										return

									}

								}, 1)

								tmp5232 := Call(__e, PrimFunc(symshen_4_5returns_6), Parseshen_4_5shortnatters_6)

								__e.TailApply(tmp5228, tmp5232)
								return

							}

						}, 1)

						tmp5235 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), Parseshen_4_5backslash_6)

						__e.TailApply(tmp5227, tmp5235)
						return

					}

				}, 1)

				tmp5238 := Call(__e, PrimFunc(symshen_4_5backslash_6), Parseshen_4_5backslash_6)

				__e.TailApply(tmp5226, tmp5238)
				return

			}

		}, 1)

		tmp5241 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2829)

		tmp5242 := Call(__e, tmp5225, tmp5241)

		__e.TailApply(tmp5222, tmp5242)
		return

	}, 1)

	tmp5243 := Call(__e, ns2_1set, symshen_4_5singleline_6, tmp5221)

	_ = tmp5243

	tmp5244 := MakeNative(func(__e *ControlFlow) {
		V2830 := __e.Get(1)
		_ = V2830
		tmp5245 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5247 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5247 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5254 := Call(__e, PrimFunc(symshen_4_ahd_2), V2830, MakeNumber(92))

		var ifres5248 Obj

		if True == tmp5254 {
			tmp5249 := MakeNative(func(__e *ControlFlow) {
				News2521 := __e.Get(1)
				_ = News2521
				tmp5250 := Call(__e, PrimFunc(symshen_4in_1_6), News2521)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5250, symshen_4skip)
				return

			}, 1)

			tmp5251 := Call(__e, PrimFunc(symshen_4tls), V2830)

			tmp5252 := Call(__e, tmp5249, tmp5251)

			ifres5248 = tmp5252

		} else {
			tmp5253 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5248 = tmp5253

		}

		__e.TailApply(tmp5245, ifres5248)
		return

	}, 1)

	tmp5255 := Call(__e, ns2_1set, symshen_4_5backslash_6, tmp5244)

	_ = tmp5255

	tmp5256 := MakeNative(func(__e *ControlFlow) {
		V2831 := __e.Get(1)
		_ = V2831
		tmp5257 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5268 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5268 {
				tmp5258 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5260 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5260 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5261 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp5264 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp5264 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5262 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5262, symshen_4skip)
						return

					}

				}, 1)

				tmp5265 := Call(__e, PrimFunc(sym_5e_6), V2831)

				tmp5266 := Call(__e, tmp5261, tmp5265)

				__e.TailApply(tmp5258, tmp5266)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5269 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5shortnatter_6 := __e.Get(1)
			_ = Parseshen_4_5shortnatter_6
			tmp5276 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5shortnatter_6)

			if True == tmp5276 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5270 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5shortnatters_6 := __e.Get(1)
					_ = Parseshen_4_5shortnatters_6
					tmp5273 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5shortnatters_6)

					if True == tmp5273 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5271 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5shortnatters_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5271, symshen_4skip)
						return

					}

				}, 1)

				tmp5274 := Call(__e, PrimFunc(symshen_4_5shortnatters_6), Parseshen_4_5shortnatter_6)

				__e.TailApply(tmp5270, tmp5274)
				return

			}

		}, 1)

		tmp5277 := Call(__e, PrimFunc(symshen_4_5shortnatter_6), V2831)

		tmp5278 := Call(__e, tmp5269, tmp5277)

		__e.TailApply(tmp5257, tmp5278)
		return

	}, 1)

	tmp5279 := Call(__e, ns2_1set, symshen_4_5shortnatters_6, tmp5256)

	_ = tmp5279

	tmp5280 := MakeNative(func(__e *ControlFlow) {
		V2832 := __e.Get(1)
		_ = V2832
		tmp5281 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5283 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5283 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5295 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2832)

		var ifres5284 Obj

		if True == tmp5295 {
			tmp5285 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp5286 := MakeNative(func(__e *ControlFlow) {
					News2524 := __e.Get(1)
					_ = News2524
					tmp5289 := Call(__e, PrimFunc(symshen_4return_2), Byte)

					tmp5290 := PrimNot(tmp5289)

					if True == tmp5290 {
						tmp5287 := Call(__e, PrimFunc(symshen_4in_1_6), News2524)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5287, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5291 := Call(__e, PrimFunc(symshen_4tls), V2832)

				__e.TailApply(tmp5286, tmp5291)
				return

			}, 1)

			tmp5292 := Call(__e, PrimFunc(symshen_4hds), V2832)

			tmp5293 := Call(__e, tmp5285, tmp5292)

			ifres5284 = tmp5293

		} else {
			tmp5294 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5284 = tmp5294

		}

		__e.TailApply(tmp5281, ifres5284)
		return

	}, 1)

	tmp5296 := Call(__e, ns2_1set, symshen_4_5shortnatter_6, tmp5280)

	_ = tmp5296

	tmp5297 := MakeNative(func(__e *ControlFlow) {
		V2833 := __e.Get(1)
		_ = V2833
		tmp5298 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5309 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5309 {
				tmp5299 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5301 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5301 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5302 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5return_6 := __e.Get(1)
					_ = Parseshen_4_5return_6
					tmp5305 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5return_6)

					if True == tmp5305 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5303 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5return_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5303, symshen_4skip)
						return

					}

				}, 1)

				tmp5306 := Call(__e, PrimFunc(symshen_4_5return_6), V2833)

				tmp5307 := Call(__e, tmp5302, tmp5306)

				__e.TailApply(tmp5299, tmp5307)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5310 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5return_6 := __e.Get(1)
			_ = Parseshen_4_5return_6
			tmp5317 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5return_6)

			if True == tmp5317 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5311 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5returns_6 := __e.Get(1)
					_ = Parseshen_4_5returns_6
					tmp5314 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5returns_6)

					if True == tmp5314 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5312 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5returns_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5312, symshen_4skip)
						return

					}

				}, 1)

				tmp5315 := Call(__e, PrimFunc(symshen_4_5returns_6), Parseshen_4_5return_6)

				__e.TailApply(tmp5311, tmp5315)
				return

			}

		}, 1)

		tmp5318 := Call(__e, PrimFunc(symshen_4_5return_6), V2833)

		tmp5319 := Call(__e, tmp5310, tmp5318)

		__e.TailApply(tmp5298, tmp5319)
		return

	}, 1)

	tmp5320 := Call(__e, ns2_1set, symshen_4_5returns_6, tmp5297)

	_ = tmp5320

	tmp5321 := MakeNative(func(__e *ControlFlow) {
		V2834 := __e.Get(1)
		_ = V2834
		tmp5322 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5324 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5324 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5335 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2834)

		var ifres5325 Obj

		if True == tmp5335 {
			tmp5326 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp5327 := MakeNative(func(__e *ControlFlow) {
					News2527 := __e.Get(1)
					_ = News2527
					tmp5330 := Call(__e, PrimFunc(symshen_4return_2), Byte)

					if True == tmp5330 {
						tmp5328 := Call(__e, PrimFunc(symshen_4in_1_6), News2527)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5328, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5331 := Call(__e, PrimFunc(symshen_4tls), V2834)

				__e.TailApply(tmp5327, tmp5331)
				return

			}, 1)

			tmp5332 := Call(__e, PrimFunc(symshen_4hds), V2834)

			tmp5333 := Call(__e, tmp5326, tmp5332)

			ifres5325 = tmp5333

		} else {
			tmp5334 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5325 = tmp5334

		}

		__e.TailApply(tmp5322, ifres5325)
		return

	}, 1)

	tmp5336 := Call(__e, ns2_1set, symshen_4_5return_6, tmp5321)

	_ = tmp5336

	tmp5337 := MakeNative(func(__e *ControlFlow) {
		V2835 := __e.Get(1)
		_ = V2835
		tmp5338 := PrimCons(MakeNumber(13), Nil)

		tmp5339 := PrimCons(MakeNumber(10), tmp5338)

		tmp5340 := PrimCons(MakeNumber(9), tmp5339)

		__e.TailApply(PrimFunc(symelement_2), V2835, tmp5340)
		return

	}, 1)

	tmp5341 := Call(__e, ns2_1set, symshen_4return_2, tmp5337)

	_ = tmp5341

	tmp5342 := MakeNative(func(__e *ControlFlow) {
		V2836 := __e.Get(1)
		_ = V2836
		tmp5343 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5345 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5345 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5346 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5backslash_6 := __e.Get(1)
			_ = Parseshen_4_5backslash_6
			tmp5357 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

			if True == tmp5357 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5347 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5times_6 := __e.Get(1)
					_ = Parseshen_4_5times_6
					tmp5354 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5times_6)

					if True == tmp5354 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5348 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5longnatter_6 := __e.Get(1)
							_ = Parseshen_4_5longnatter_6
							tmp5351 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5longnatter_6)

							if True == tmp5351 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5349 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5longnatter_6)

								__e.TailApply(PrimFunc(symshen_4comb), tmp5349, symshen_4skip)
								return

							}

						}, 1)

						tmp5352 := Call(__e, PrimFunc(symshen_4_5longnatter_6), Parseshen_4_5times_6)

						__e.TailApply(tmp5348, tmp5352)
						return

					}

				}, 1)

				tmp5355 := Call(__e, PrimFunc(symshen_4_5times_6), Parseshen_4_5backslash_6)

				__e.TailApply(tmp5347, tmp5355)
				return

			}

		}, 1)

		tmp5358 := Call(__e, PrimFunc(symshen_4_5backslash_6), V2836)

		tmp5359 := Call(__e, tmp5346, tmp5358)

		__e.TailApply(tmp5343, tmp5359)
		return

	}, 1)

	tmp5360 := Call(__e, ns2_1set, symshen_4_5multiline_6, tmp5342)

	_ = tmp5360

	tmp5361 := MakeNative(func(__e *ControlFlow) {
		V2837 := __e.Get(1)
		_ = V2837
		tmp5362 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5364 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5364 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5371 := Call(__e, PrimFunc(symshen_4_ahd_2), V2837, MakeNumber(42))

		var ifres5365 Obj

		if True == tmp5371 {
			tmp5366 := MakeNative(func(__e *ControlFlow) {
				News2530 := __e.Get(1)
				_ = News2530
				tmp5367 := Call(__e, PrimFunc(symshen_4in_1_6), News2530)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5367, symshen_4skip)
				return

			}, 1)

			tmp5368 := Call(__e, PrimFunc(symshen_4tls), V2837)

			tmp5369 := Call(__e, tmp5366, tmp5368)

			ifres5365 = tmp5369

		} else {
			tmp5370 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5365 = tmp5370

		}

		__e.TailApply(tmp5362, ifres5365)
		return

	}, 1)

	tmp5372 := Call(__e, ns2_1set, symshen_4_5times_6, tmp5361)

	_ = tmp5372

	tmp5373 := MakeNative(func(__e *ControlFlow) {
		V2838 := __e.Get(1)
		_ = V2838
		tmp5374 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5403 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5403 {
				tmp5375 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5391 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5391 {
						tmp5376 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp5378 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

							if True == tmp5378 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp5389 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2838)

						var ifres5379 Obj

						if True == tmp5389 {
							tmp5380 := MakeNative(func(__e *ControlFlow) {
								News2532 := __e.Get(1)
								_ = News2532
								tmp5381 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5longnatter_6 := __e.Get(1)
									_ = Parseshen_4_5longnatter_6
									tmp5384 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5longnatter_6)

									if True == tmp5384 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp5382 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5longnatter_6)

										__e.TailApply(PrimFunc(symshen_4comb), tmp5382, symshen_4skip)
										return

									}

								}, 1)

								tmp5385 := Call(__e, PrimFunc(symshen_4_5longnatter_6), News2532)

								__e.TailApply(tmp5381, tmp5385)
								return

							}, 1)

							tmp5386 := Call(__e, PrimFunc(symshen_4tls), V2838)

							tmp5387 := Call(__e, tmp5380, tmp5386)

							ifres5379 = tmp5387

						} else {
							tmp5388 := Call(__e, PrimFunc(symshen_4parse_1failure))

							ifres5379 = tmp5388

						}

						__e.TailApply(tmp5376, ifres5379)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5392 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5times_6 := __e.Get(1)
					_ = Parseshen_4_5times_6
					tmp5399 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5times_6)

					if True == tmp5399 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5393 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5backslash_6 := __e.Get(1)
							_ = Parseshen_4_5backslash_6
							tmp5396 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

							if True == tmp5396 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5394 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5backslash_6)

								__e.TailApply(PrimFunc(symshen_4comb), tmp5394, symshen_4skip)
								return

							}

						}, 1)

						tmp5397 := Call(__e, PrimFunc(symshen_4_5backslash_6), Parseshen_4_5times_6)

						__e.TailApply(tmp5393, tmp5397)
						return

					}

				}, 1)

				tmp5400 := Call(__e, PrimFunc(symshen_4_5times_6), V2838)

				tmp5401 := Call(__e, tmp5392, tmp5400)

				__e.TailApply(tmp5375, tmp5401)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5404 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5comment_6 := __e.Get(1)
			_ = Parseshen_4_5comment_6
			tmp5411 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5comment_6)

			if True == tmp5411 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5405 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5longnatter_6 := __e.Get(1)
					_ = Parseshen_4_5longnatter_6
					tmp5408 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5longnatter_6)

					if True == tmp5408 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5406 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5longnatter_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5406, symshen_4skip)
						return

					}

				}, 1)

				tmp5409 := Call(__e, PrimFunc(symshen_4_5longnatter_6), Parseshen_4_5comment_6)

				__e.TailApply(tmp5405, tmp5409)
				return

			}

		}, 1)

		tmp5412 := Call(__e, PrimFunc(symshen_4_5comment_6), V2838)

		tmp5413 := Call(__e, tmp5404, tmp5412)

		__e.TailApply(tmp5374, tmp5413)
		return

	}, 1)

	tmp5414 := Call(__e, ns2_1set, symshen_4_5longnatter_6, tmp5373)

	_ = tmp5414

	tmp5415 := MakeNative(func(__e *ControlFlow) {
		V2839 := __e.Get(1)
		_ = V2839
		tmp5416 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5444 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5444 {
				tmp5417 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5435 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5435 {
						tmp5418 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp5420 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

							if True == tmp5420 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp5421 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5sym_6 := __e.Get(1)
							_ = Parseshen_4_5sym_6
							tmp5431 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5sym_6)

							if True == tmp5431 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5422 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5sym_6)

								tmp5428 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5sym_6)

								tmp5429 := PrimEqual(tmp5428, MakeString("<>"))

								var ifres5423 Obj

								if True == tmp5429 {
									tmp5424 := PrimCons(MakeNumber(0), Nil)

									tmp5425 := PrimCons(symvector, tmp5424)

									ifres5423 = tmp5425

								} else {
									tmp5426 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5sym_6)

									tmp5427 := PrimIntern(tmp5426)

									ifres5423 = tmp5427

								}

								__e.TailApply(PrimFunc(symshen_4comb), tmp5422, ifres5423)
								return

							}

						}, 1)

						tmp5432 := Call(__e, PrimFunc(symshen_4_5sym_6), V2839)

						tmp5433 := Call(__e, tmp5421, tmp5432)

						__e.TailApply(tmp5418, tmp5433)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5436 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5number_6 := __e.Get(1)
					_ = Parseshen_4_5number_6
					tmp5440 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5number_6)

					if True == tmp5440 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5437 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5number_6)

						tmp5438 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5number_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5437, tmp5438)
						return

					}

				}, 1)

				tmp5441 := Call(__e, PrimFunc(symshen_4_5number_6), V2839)

				tmp5442 := Call(__e, tmp5436, tmp5441)

				__e.TailApply(tmp5417, tmp5442)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5445 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5str_6 := __e.Get(1)
			_ = Parseshen_4_5str_6
			tmp5449 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5str_6)

			if True == tmp5449 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5446 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5str_6)

				tmp5447 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5str_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5446, tmp5447)
				return

			}

		}, 1)

		tmp5450 := Call(__e, PrimFunc(symshen_4_5str_6), V2839)

		tmp5451 := Call(__e, tmp5445, tmp5450)

		__e.TailApply(tmp5416, tmp5451)
		return

	}, 1)

	tmp5452 := Call(__e, ns2_1set, symshen_4_5atom_6, tmp5415)

	_ = tmp5452

	tmp5453 := MakeNative(func(__e *ControlFlow) {
		V2840 := __e.Get(1)
		_ = V2840
		tmp5454 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5456 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5456 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5457 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5alpha_6 := __e.Get(1)
			_ = Parseshen_4_5alpha_6
			tmp5467 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5alpha_6)

			if True == tmp5467 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5458 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5alphanums_6 := __e.Get(1)
					_ = Parseshen_4_5alphanums_6
					tmp5464 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5alphanums_6)

					if True == tmp5464 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5459 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5alphanums_6)

						tmp5460 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5alpha_6)

						tmp5461 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5alphanums_6)

						tmp5462 := PrimStringConcat(tmp5460, tmp5461)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5459, tmp5462)
						return

					}

				}, 1)

				tmp5465 := Call(__e, PrimFunc(symshen_4_5alphanums_6), Parseshen_4_5alpha_6)

				__e.TailApply(tmp5458, tmp5465)
				return

			}

		}, 1)

		tmp5468 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2840)

		tmp5469 := Call(__e, tmp5457, tmp5468)

		__e.TailApply(tmp5454, tmp5469)
		return

	}, 1)

	tmp5470 := Call(__e, ns2_1set, symshen_4_5sym_6, tmp5453)

	_ = tmp5470

	tmp5471 := MakeNative(func(__e *ControlFlow) {
		V2841 := __e.Get(1)
		_ = V2841
		tmp5472 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5474 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5474 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5486 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2841)

		var ifres5475 Obj

		if True == tmp5486 {
			tmp5476 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp5477 := MakeNative(func(__e *ControlFlow) {
					News2536 := __e.Get(1)
					_ = News2536
					tmp5481 := Call(__e, PrimFunc(symshen_4alpha_2), Byte)

					if True == tmp5481 {
						tmp5478 := Call(__e, PrimFunc(symshen_4in_1_6), News2536)

						tmp5479 := PrimNumberToString(Byte)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5478, tmp5479)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5482 := Call(__e, PrimFunc(symshen_4tls), V2841)

				__e.TailApply(tmp5477, tmp5482)
				return

			}, 1)

			tmp5483 := Call(__e, PrimFunc(symshen_4hds), V2841)

			tmp5484 := Call(__e, tmp5476, tmp5483)

			ifres5475 = tmp5484

		} else {
			tmp5485 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5475 = tmp5485

		}

		__e.TailApply(tmp5472, ifres5475)
		return

	}, 1)

	tmp5487 := Call(__e, ns2_1set, symshen_4_5alpha_6, tmp5471)

	_ = tmp5487

	tmp5488 := MakeNative(func(__e *ControlFlow) {
		V2842 := __e.Get(1)
		_ = V2842
		tmp5495 := Call(__e, PrimFunc(symshen_4lowercase_2), V2842)

		if True == tmp5495 {
			__e.Return(True)
			return
		} else {
			tmp5493 := Call(__e, PrimFunc(symshen_4uppercase_2), V2842)

			var ifres5490 Obj

			if True == tmp5493 {
				ifres5490 = True

			} else {
				tmp5492 := Call(__e, PrimFunc(symshen_4misc_2), V2842)

				var ifres5491 Obj

				if True == tmp5492 {
					ifres5491 = True

				} else {
					ifres5491 = False

				}

				ifres5490 = ifres5491

			}

			if True == ifres5490 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp5496 := Call(__e, ns2_1set, symshen_4alpha_2, tmp5488)

	_ = tmp5496

	tmp5497 := MakeNative(func(__e *ControlFlow) {
		V2843 := __e.Get(1)
		_ = V2843
		tmp5501 := PrimGreatEqual(V2843, MakeNumber(97))

		if True == tmp5501 {
			tmp5499 := PrimLessEqual(V2843, MakeNumber(122))

			if True == tmp5499 {
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

	tmp5502 := Call(__e, ns2_1set, symshen_4lowercase_2, tmp5497)

	_ = tmp5502

	tmp5503 := MakeNative(func(__e *ControlFlow) {
		V2844 := __e.Get(1)
		_ = V2844
		tmp5507 := PrimGreatEqual(V2844, MakeNumber(65))

		if True == tmp5507 {
			tmp5505 := PrimLessEqual(V2844, MakeNumber(90))

			if True == tmp5505 {
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

	tmp5508 := Call(__e, ns2_1set, symshen_4uppercase_2, tmp5503)

	_ = tmp5508

	tmp5509 := MakeNative(func(__e *ControlFlow) {
		V2845 := __e.Get(1)
		_ = V2845
		tmp5510 := PrimCons(MakeNumber(96), Nil)

		tmp5511 := PrimCons(MakeNumber(35), tmp5510)

		tmp5512 := PrimCons(MakeNumber(39), tmp5511)

		tmp5513 := PrimCons(MakeNumber(37), tmp5512)

		tmp5514 := PrimCons(MakeNumber(38), tmp5513)

		tmp5515 := PrimCons(MakeNumber(60), tmp5514)

		tmp5516 := PrimCons(MakeNumber(62), tmp5515)

		tmp5517 := PrimCons(MakeNumber(46), tmp5516)

		tmp5518 := PrimCons(MakeNumber(126), tmp5517)

		tmp5519 := PrimCons(MakeNumber(64), tmp5518)

		tmp5520 := PrimCons(MakeNumber(33), tmp5519)

		tmp5521 := PrimCons(MakeNumber(36), tmp5520)

		tmp5522 := PrimCons(MakeNumber(63), tmp5521)

		tmp5523 := PrimCons(MakeNumber(95), tmp5522)

		tmp5524 := PrimCons(MakeNumber(43), tmp5523)

		tmp5525 := PrimCons(MakeNumber(47), tmp5524)

		tmp5526 := PrimCons(MakeNumber(42), tmp5525)

		tmp5527 := PrimCons(MakeNumber(45), tmp5526)

		tmp5528 := PrimCons(MakeNumber(61), tmp5527)

		__e.TailApply(PrimFunc(symelement_2), V2845, tmp5528)
		return

	}, 1)

	tmp5529 := Call(__e, ns2_1set, symshen_4misc_2, tmp5509)

	_ = tmp5529

	tmp5530 := MakeNative(func(__e *ControlFlow) {
		V2846 := __e.Get(1)
		_ = V2846
		tmp5531 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5542 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5542 {
				tmp5532 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5534 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5534 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5535 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp5538 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp5538 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5536 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5536, MakeString(""))
						return

					}

				}, 1)

				tmp5539 := Call(__e, PrimFunc(sym_5e_6), V2846)

				tmp5540 := Call(__e, tmp5535, tmp5539)

				__e.TailApply(tmp5532, tmp5540)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5543 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5alphanum_6 := __e.Get(1)
			_ = Parseshen_4_5alphanum_6
			tmp5553 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5alphanum_6)

			if True == tmp5553 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5544 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5alphanums_6 := __e.Get(1)
					_ = Parseshen_4_5alphanums_6
					tmp5550 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5alphanums_6)

					if True == tmp5550 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5545 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5alphanums_6)

						tmp5546 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5alphanum_6)

						tmp5547 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5alphanums_6)

						tmp5548 := PrimStringConcat(tmp5546, tmp5547)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5545, tmp5548)
						return

					}

				}, 1)

				tmp5551 := Call(__e, PrimFunc(symshen_4_5alphanums_6), Parseshen_4_5alphanum_6)

				__e.TailApply(tmp5544, tmp5551)
				return

			}

		}, 1)

		tmp5554 := Call(__e, PrimFunc(symshen_4_5alphanum_6), V2846)

		tmp5555 := Call(__e, tmp5543, tmp5554)

		__e.TailApply(tmp5531, tmp5555)
		return

	}, 1)

	tmp5556 := Call(__e, ns2_1set, symshen_4_5alphanums_6, tmp5530)

	_ = tmp5556

	tmp5557 := MakeNative(func(__e *ControlFlow) {
		V2847 := __e.Get(1)
		_ = V2847
		tmp5558 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5570 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5570 {
				tmp5559 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5561 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5561 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5562 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5numeral_6 := __e.Get(1)
					_ = Parseshen_4_5numeral_6
					tmp5566 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5numeral_6)

					if True == tmp5566 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5563 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5numeral_6)

						tmp5564 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5numeral_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5563, tmp5564)
						return

					}

				}, 1)

				tmp5567 := Call(__e, PrimFunc(symshen_4_5numeral_6), V2847)

				tmp5568 := Call(__e, tmp5562, tmp5567)

				__e.TailApply(tmp5559, tmp5568)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5571 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5alpha_6 := __e.Get(1)
			_ = Parseshen_4_5alpha_6
			tmp5575 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5alpha_6)

			if True == tmp5575 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5572 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5alpha_6)

				tmp5573 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5alpha_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5572, tmp5573)
				return

			}

		}, 1)

		tmp5576 := Call(__e, PrimFunc(symshen_4_5alpha_6), V2847)

		tmp5577 := Call(__e, tmp5571, tmp5576)

		__e.TailApply(tmp5558, tmp5577)
		return

	}, 1)

	tmp5578 := Call(__e, ns2_1set, symshen_4_5alphanum_6, tmp5557)

	_ = tmp5578

	tmp5579 := MakeNative(func(__e *ControlFlow) {
		V2848 := __e.Get(1)
		_ = V2848
		tmp5580 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5582 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5582 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5594 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2848)

		var ifres5583 Obj

		if True == tmp5594 {
			tmp5584 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp5585 := MakeNative(func(__e *ControlFlow) {
					News2540 := __e.Get(1)
					_ = News2540
					tmp5589 := Call(__e, PrimFunc(symshen_4digit_2), Byte)

					if True == tmp5589 {
						tmp5586 := Call(__e, PrimFunc(symshen_4in_1_6), News2540)

						tmp5587 := PrimNumberToString(Byte)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5586, tmp5587)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5590 := Call(__e, PrimFunc(symshen_4tls), V2848)

				__e.TailApply(tmp5585, tmp5590)
				return

			}, 1)

			tmp5591 := Call(__e, PrimFunc(symshen_4hds), V2848)

			tmp5592 := Call(__e, tmp5584, tmp5591)

			ifres5583 = tmp5592

		} else {
			tmp5593 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5583 = tmp5593

		}

		__e.TailApply(tmp5580, ifres5583)
		return

	}, 1)

	tmp5595 := Call(__e, ns2_1set, symshen_4_5numeral_6, tmp5579)

	_ = tmp5595

	tmp5596 := MakeNative(func(__e *ControlFlow) {
		V2849 := __e.Get(1)
		_ = V2849
		tmp5600 := PrimGreatEqual(V2849, MakeNumber(48))

		if True == tmp5600 {
			tmp5598 := PrimLessEqual(V2849, MakeNumber(57))

			if True == tmp5598 {
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

	tmp5601 := Call(__e, ns2_1set, symshen_4digit_2, tmp5596)

	_ = tmp5601

	tmp5602 := MakeNative(func(__e *ControlFlow) {
		V2850 := __e.Get(1)
		_ = V2850
		tmp5603 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5605 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5605 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5606 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5dbq_6 := __e.Get(1)
			_ = Parseshen_4_5dbq_6
			tmp5618 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5dbq_6)

			if True == tmp5618 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5607 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5strcontents_6 := __e.Get(1)
					_ = Parseshen_4_5strcontents_6
					tmp5615 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5strcontents_6)

					if True == tmp5615 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5608 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5dbq_6 := __e.Get(1)
							_ = Parseshen_4_5dbq_6
							tmp5612 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5dbq_6)

							if True == tmp5612 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5609 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5dbq_6)

								tmp5610 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5strcontents_6)

								__e.TailApply(PrimFunc(symshen_4comb), tmp5609, tmp5610)
								return

							}

						}, 1)

						tmp5613 := Call(__e, PrimFunc(symshen_4_5dbq_6), Parseshen_4_5strcontents_6)

						__e.TailApply(tmp5608, tmp5613)
						return

					}

				}, 1)

				tmp5616 := Call(__e, PrimFunc(symshen_4_5strcontents_6), Parseshen_4_5dbq_6)

				__e.TailApply(tmp5607, tmp5616)
				return

			}

		}, 1)

		tmp5619 := Call(__e, PrimFunc(symshen_4_5dbq_6), V2850)

		tmp5620 := Call(__e, tmp5606, tmp5619)

		__e.TailApply(tmp5603, tmp5620)
		return

	}, 1)

	tmp5621 := Call(__e, ns2_1set, symshen_4_5str_6, tmp5602)

	_ = tmp5621

	tmp5622 := MakeNative(func(__e *ControlFlow) {
		V2851 := __e.Get(1)
		_ = V2851
		tmp5623 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5625 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5625 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5632 := Call(__e, PrimFunc(symshen_4_ahd_2), V2851, MakeNumber(34))

		var ifres5626 Obj

		if True == tmp5632 {
			tmp5627 := MakeNative(func(__e *ControlFlow) {
				News2543 := __e.Get(1)
				_ = News2543
				tmp5628 := Call(__e, PrimFunc(symshen_4in_1_6), News2543)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5628, symshen_4skip)
				return

			}, 1)

			tmp5629 := Call(__e, PrimFunc(symshen_4tls), V2851)

			tmp5630 := Call(__e, tmp5627, tmp5629)

			ifres5626 = tmp5630

		} else {
			tmp5631 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5626 = tmp5631

		}

		__e.TailApply(tmp5623, ifres5626)
		return

	}, 1)

	tmp5633 := Call(__e, ns2_1set, symshen_4_5dbq_6, tmp5622)

	_ = tmp5633

	tmp5634 := MakeNative(func(__e *ControlFlow) {
		V2852 := __e.Get(1)
		_ = V2852
		tmp5635 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5646 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5646 {
				tmp5636 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5638 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5638 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5639 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp5642 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp5642 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5640 := Call(__e, PrimFunc(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5640, MakeString(""))
						return

					}

				}, 1)

				tmp5643 := Call(__e, PrimFunc(sym_5e_6), V2852)

				tmp5644 := Call(__e, tmp5639, tmp5643)

				__e.TailApply(tmp5636, tmp5644)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5647 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5strc_6 := __e.Get(1)
			_ = Parseshen_4_5strc_6
			tmp5657 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5strc_6)

			if True == tmp5657 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5648 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5strcontents_6 := __e.Get(1)
					_ = Parseshen_4_5strcontents_6
					tmp5654 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5strcontents_6)

					if True == tmp5654 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5649 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5strcontents_6)

						tmp5650 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5strc_6)

						tmp5651 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5strcontents_6)

						tmp5652 := PrimStringConcat(tmp5650, tmp5651)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5649, tmp5652)
						return

					}

				}, 1)

				tmp5655 := Call(__e, PrimFunc(symshen_4_5strcontents_6), Parseshen_4_5strc_6)

				__e.TailApply(tmp5648, tmp5655)
				return

			}

		}, 1)

		tmp5658 := Call(__e, PrimFunc(symshen_4_5strc_6), V2852)

		tmp5659 := Call(__e, tmp5647, tmp5658)

		__e.TailApply(tmp5635, tmp5659)
		return

	}, 1)

	tmp5660 := Call(__e, ns2_1set, symshen_4_5strcontents_6, tmp5634)

	_ = tmp5660

	tmp5661 := MakeNative(func(__e *ControlFlow) {
		V2853 := __e.Get(1)
		_ = V2853
		tmp5662 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5674 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5674 {
				tmp5663 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5665 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5665 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5666 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5notdbq_6 := __e.Get(1)
					_ = Parseshen_4_5notdbq_6
					tmp5670 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5notdbq_6)

					if True == tmp5670 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5667 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5notdbq_6)

						tmp5668 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5notdbq_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5667, tmp5668)
						return

					}

				}, 1)

				tmp5671 := Call(__e, PrimFunc(symshen_4_5notdbq_6), V2853)

				tmp5672 := Call(__e, tmp5666, tmp5671)

				__e.TailApply(tmp5663, tmp5672)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5675 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5control_6 := __e.Get(1)
			_ = Parseshen_4_5control_6
			tmp5679 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5control_6)

			if True == tmp5679 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5676 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5control_6)

				tmp5677 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5control_6)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5676, tmp5677)
				return

			}

		}, 1)

		tmp5680 := Call(__e, PrimFunc(symshen_4_5control_6), V2853)

		tmp5681 := Call(__e, tmp5675, tmp5680)

		__e.TailApply(tmp5662, tmp5681)
		return

	}, 1)

	tmp5682 := Call(__e, ns2_1set, symshen_4_5strc_6, tmp5661)

	_ = tmp5682

	tmp5683 := MakeNative(func(__e *ControlFlow) {
		V2854 := __e.Get(1)
		_ = V2854
		tmp5684 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5686 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5686 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5687 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5lowC_6 := __e.Get(1)
			_ = Parseshen_4_5lowC_6
			tmp5704 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5lowC_6)

			if True == tmp5704 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5688 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5hash_6 := __e.Get(1)
					_ = Parseshen_4_5hash_6
					tmp5701 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5hash_6)

					if True == tmp5701 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5689 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5integer_6 := __e.Get(1)
							_ = Parseshen_4_5integer_6
							tmp5698 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

							if True == tmp5698 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5690 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5semicolon_6 := __e.Get(1)
									_ = Parseshen_4_5semicolon_6
									tmp5695 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5semicolon_6)

									if True == tmp5695 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp5691 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5semicolon_6)

										tmp5692 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5integer_6)

										tmp5693 := PrimNumberToString(tmp5692)

										__e.TailApply(PrimFunc(symshen_4comb), tmp5691, tmp5693)
										return

									}

								}, 1)

								tmp5696 := Call(__e, PrimFunc(symshen_4_5semicolon_6), Parseshen_4_5integer_6)

								__e.TailApply(tmp5690, tmp5696)
								return

							}

						}, 1)

						tmp5699 := Call(__e, PrimFunc(symshen_4_5integer_6), Parseshen_4_5hash_6)

						__e.TailApply(tmp5689, tmp5699)
						return

					}

				}, 1)

				tmp5702 := Call(__e, PrimFunc(symshen_4_5hash_6), Parseshen_4_5lowC_6)

				__e.TailApply(tmp5688, tmp5702)
				return

			}

		}, 1)

		tmp5705 := Call(__e, PrimFunc(symshen_4_5lowC_6), V2854)

		tmp5706 := Call(__e, tmp5687, tmp5705)

		__e.TailApply(tmp5684, tmp5706)
		return

	}, 1)

	tmp5707 := Call(__e, ns2_1set, symshen_4_5control_6, tmp5683)

	_ = tmp5707

	tmp5708 := MakeNative(func(__e *ControlFlow) {
		V2855 := __e.Get(1)
		_ = V2855
		tmp5709 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5711 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5711 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5724 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2855)

		var ifres5712 Obj

		if True == tmp5724 {
			tmp5713 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp5714 := MakeNative(func(__e *ControlFlow) {
					News2548 := __e.Get(1)
					_ = News2548
					tmp5718 := PrimEqual(Byte, MakeNumber(34))

					tmp5719 := PrimNot(tmp5718)

					if True == tmp5719 {
						tmp5715 := Call(__e, PrimFunc(symshen_4in_1_6), News2548)

						tmp5716 := PrimNumberToString(Byte)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5715, tmp5716)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5720 := Call(__e, PrimFunc(symshen_4tls), V2855)

				__e.TailApply(tmp5714, tmp5720)
				return

			}, 1)

			tmp5721 := Call(__e, PrimFunc(symshen_4hds), V2855)

			tmp5722 := Call(__e, tmp5713, tmp5721)

			ifres5712 = tmp5722

		} else {
			tmp5723 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5712 = tmp5723

		}

		__e.TailApply(tmp5709, ifres5712)
		return

	}, 1)

	tmp5725 := Call(__e, ns2_1set, symshen_4_5notdbq_6, tmp5708)

	_ = tmp5725

	tmp5726 := MakeNative(func(__e *ControlFlow) {
		V2856 := __e.Get(1)
		_ = V2856
		tmp5727 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5729 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5729 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5736 := Call(__e, PrimFunc(symshen_4_ahd_2), V2856, MakeNumber(99))

		var ifres5730 Obj

		if True == tmp5736 {
			tmp5731 := MakeNative(func(__e *ControlFlow) {
				News2550 := __e.Get(1)
				_ = News2550
				tmp5732 := Call(__e, PrimFunc(symshen_4in_1_6), News2550)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5732, symshen_4skip)
				return

			}, 1)

			tmp5733 := Call(__e, PrimFunc(symshen_4tls), V2856)

			tmp5734 := Call(__e, tmp5731, tmp5733)

			ifres5730 = tmp5734

		} else {
			tmp5735 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5730 = tmp5735

		}

		__e.TailApply(tmp5727, ifres5730)
		return

	}, 1)

	tmp5737 := Call(__e, ns2_1set, symshen_4_5lowC_6, tmp5726)

	_ = tmp5737

	tmp5738 := MakeNative(func(__e *ControlFlow) {
		V2857 := __e.Get(1)
		_ = V2857
		tmp5739 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5741 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5741 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5748 := Call(__e, PrimFunc(symshen_4_ahd_2), V2857, MakeNumber(35))

		var ifres5742 Obj

		if True == tmp5748 {
			tmp5743 := MakeNative(func(__e *ControlFlow) {
				News2552 := __e.Get(1)
				_ = News2552
				tmp5744 := Call(__e, PrimFunc(symshen_4in_1_6), News2552)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5744, symshen_4skip)
				return

			}, 1)

			tmp5745 := Call(__e, PrimFunc(symshen_4tls), V2857)

			tmp5746 := Call(__e, tmp5743, tmp5745)

			ifres5742 = tmp5746

		} else {
			tmp5747 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5742 = tmp5747

		}

		__e.TailApply(tmp5739, ifres5742)
		return

	}, 1)

	tmp5749 := Call(__e, ns2_1set, symshen_4_5hash_6, tmp5738)

	_ = tmp5749

	tmp5750 := MakeNative(func(__e *ControlFlow) {
		V2858 := __e.Get(1)
		_ = V2858
		tmp5751 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5797 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5797 {
				tmp5752 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5784 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5784 {
						tmp5753 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp5775 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

							if True == tmp5775 {
								tmp5754 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp5766 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

									if True == tmp5766 {
										tmp5755 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp5757 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

											if True == tmp5757 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp5758 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5integer_6 := __e.Get(1)
											_ = Parseshen_4_5integer_6
											tmp5762 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

											if True == tmp5762 {
												__e.TailApply(PrimFunc(symshen_4parse_1failure))
												return
											} else {
												tmp5759 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5integer_6)

												tmp5760 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5integer_6)

												__e.TailApply(PrimFunc(symshen_4comb), tmp5759, tmp5760)
												return

											}

										}, 1)

										tmp5763 := Call(__e, PrimFunc(symshen_4_5integer_6), V2858)

										tmp5764 := Call(__e, tmp5758, tmp5763)

										__e.TailApply(tmp5755, tmp5764)
										return

									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp5767 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5float_6 := __e.Get(1)
									_ = Parseshen_4_5float_6
									tmp5771 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5float_6)

									if True == tmp5771 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp5768 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5float_6)

										tmp5769 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5float_6)

										__e.TailApply(PrimFunc(symshen_4comb), tmp5768, tmp5769)
										return

									}

								}, 1)

								tmp5772 := Call(__e, PrimFunc(symshen_4_5float_6), V2858)

								tmp5773 := Call(__e, tmp5767, tmp5772)

								__e.TailApply(tmp5754, tmp5773)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp5776 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5e_1number_6 := __e.Get(1)
							_ = Parseshen_4_5e_1number_6
							tmp5780 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5e_1number_6)

							if True == tmp5780 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5777 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5e_1number_6)

								tmp5778 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5e_1number_6)

								__e.TailApply(PrimFunc(symshen_4comb), tmp5777, tmp5778)
								return

							}

						}, 1)

						tmp5781 := Call(__e, PrimFunc(symshen_4_5e_1number_6), V2858)

						tmp5782 := Call(__e, tmp5776, tmp5781)

						__e.TailApply(tmp5753, tmp5782)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5785 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5plus_6 := __e.Get(1)
					_ = Parseshen_4_5plus_6
					tmp5793 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5plus_6)

					if True == tmp5793 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5786 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5number_6 := __e.Get(1)
							_ = Parseshen_4_5number_6
							tmp5790 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5number_6)

							if True == tmp5790 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5787 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5number_6)

								tmp5788 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5number_6)

								__e.TailApply(PrimFunc(symshen_4comb), tmp5787, tmp5788)
								return

							}

						}, 1)

						tmp5791 := Call(__e, PrimFunc(symshen_4_5number_6), Parseshen_4_5plus_6)

						__e.TailApply(tmp5786, tmp5791)
						return

					}

				}, 1)

				tmp5794 := Call(__e, PrimFunc(symshen_4_5plus_6), V2858)

				tmp5795 := Call(__e, tmp5785, tmp5794)

				__e.TailApply(tmp5752, tmp5795)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5798 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5minus_6 := __e.Get(1)
			_ = Parseshen_4_5minus_6
			tmp5807 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5minus_6)

			if True == tmp5807 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5799 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5number_6 := __e.Get(1)
					_ = Parseshen_4_5number_6
					tmp5804 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5number_6)

					if True == tmp5804 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5800 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5number_6)

						tmp5801 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5number_6)

						tmp5802 := PrimNumberSubtract(MakeNumber(0), tmp5801)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5800, tmp5802)
						return

					}

				}, 1)

				tmp5805 := Call(__e, PrimFunc(symshen_4_5number_6), Parseshen_4_5minus_6)

				__e.TailApply(tmp5799, tmp5805)
				return

			}

		}, 1)

		tmp5808 := Call(__e, PrimFunc(symshen_4_5minus_6), V2858)

		tmp5809 := Call(__e, tmp5798, tmp5808)

		__e.TailApply(tmp5751, tmp5809)
		return

	}, 1)

	tmp5810 := Call(__e, ns2_1set, symshen_4_5number_6, tmp5750)

	_ = tmp5810

	tmp5811 := MakeNative(func(__e *ControlFlow) {
		V2859 := __e.Get(1)
		_ = V2859
		tmp5812 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5814 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5814 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5821 := Call(__e, PrimFunc(symshen_4_ahd_2), V2859, MakeNumber(45))

		var ifres5815 Obj

		if True == tmp5821 {
			tmp5816 := MakeNative(func(__e *ControlFlow) {
				News2555 := __e.Get(1)
				_ = News2555
				tmp5817 := Call(__e, PrimFunc(symshen_4in_1_6), News2555)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5817, symshen_4skip)
				return

			}, 1)

			tmp5818 := Call(__e, PrimFunc(symshen_4tls), V2859)

			tmp5819 := Call(__e, tmp5816, tmp5818)

			ifres5815 = tmp5819

		} else {
			tmp5820 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5815 = tmp5820

		}

		__e.TailApply(tmp5812, ifres5815)
		return

	}, 1)

	tmp5822 := Call(__e, ns2_1set, symshen_4_5minus_6, tmp5811)

	_ = tmp5822

	tmp5823 := MakeNative(func(__e *ControlFlow) {
		V2860 := __e.Get(1)
		_ = V2860
		tmp5824 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5826 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5826 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5833 := Call(__e, PrimFunc(symshen_4_ahd_2), V2860, MakeNumber(43))

		var ifres5827 Obj

		if True == tmp5833 {
			tmp5828 := MakeNative(func(__e *ControlFlow) {
				News2557 := __e.Get(1)
				_ = News2557
				tmp5829 := Call(__e, PrimFunc(symshen_4in_1_6), News2557)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5829, symshen_4skip)
				return

			}, 1)

			tmp5830 := Call(__e, PrimFunc(symshen_4tls), V2860)

			tmp5831 := Call(__e, tmp5828, tmp5830)

			ifres5827 = tmp5831

		} else {
			tmp5832 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5827 = tmp5832

		}

		__e.TailApply(tmp5824, ifres5827)
		return

	}, 1)

	tmp5834 := Call(__e, ns2_1set, symshen_4_5plus_6, tmp5823)

	_ = tmp5834

	tmp5835 := MakeNative(func(__e *ControlFlow) {
		V2861 := __e.Get(1)
		_ = V2861
		tmp5836 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5838 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5838 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5839 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5digits_6 := __e.Get(1)
			_ = Parseshen_4_5digits_6
			tmp5844 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5digits_6)

			if True == tmp5844 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5840 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5digits_6)

				tmp5841 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5digits_6)

				tmp5842 := Call(__e, PrimFunc(symshen_4compute_1integer), tmp5841)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5840, tmp5842)
				return

			}

		}, 1)

		tmp5845 := Call(__e, PrimFunc(symshen_4_5digits_6), V2861)

		tmp5846 := Call(__e, tmp5839, tmp5845)

		__e.TailApply(tmp5836, tmp5846)
		return

	}, 1)

	tmp5847 := Call(__e, ns2_1set, symshen_4_5integer_6, tmp5835)

	_ = tmp5847

	tmp5848 := MakeNative(func(__e *ControlFlow) {
		V2862 := __e.Get(1)
		_ = V2862
		tmp5849 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5862 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5862 {
				tmp5850 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5852 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5852 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5853 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5digit_6 := __e.Get(1)
					_ = Parseshen_4_5digit_6
					tmp5858 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5digit_6)

					if True == tmp5858 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5854 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5digit_6)

						tmp5855 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5digit_6)

						tmp5856 := PrimCons(tmp5855, Nil)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5854, tmp5856)
						return

					}

				}, 1)

				tmp5859 := Call(__e, PrimFunc(symshen_4_5digit_6), V2862)

				tmp5860 := Call(__e, tmp5853, tmp5859)

				__e.TailApply(tmp5850, tmp5860)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5863 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5digit_6 := __e.Get(1)
			_ = Parseshen_4_5digit_6
			tmp5873 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5digit_6)

			if True == tmp5873 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5864 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5digits_6 := __e.Get(1)
					_ = Parseshen_4_5digits_6
					tmp5870 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5digits_6)

					if True == tmp5870 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5865 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5digits_6)

						tmp5866 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5digit_6)

						tmp5867 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5digits_6)

						tmp5868 := PrimCons(tmp5866, tmp5867)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5865, tmp5868)
						return

					}

				}, 1)

				tmp5871 := Call(__e, PrimFunc(symshen_4_5digits_6), Parseshen_4_5digit_6)

				__e.TailApply(tmp5864, tmp5871)
				return

			}

		}, 1)

		tmp5874 := Call(__e, PrimFunc(symshen_4_5digit_6), V2862)

		tmp5875 := Call(__e, tmp5863, tmp5874)

		__e.TailApply(tmp5849, tmp5875)
		return

	}, 1)

	tmp5876 := Call(__e, ns2_1set, symshen_4_5digits_6, tmp5848)

	_ = tmp5876

	tmp5877 := MakeNative(func(__e *ControlFlow) {
		V2863 := __e.Get(1)
		_ = V2863
		tmp5878 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5880 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5880 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5892 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2863)

		var ifres5881 Obj

		if True == tmp5892 {
			tmp5882 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp5883 := MakeNative(func(__e *ControlFlow) {
					News2561 := __e.Get(1)
					_ = News2561
					tmp5887 := Call(__e, PrimFunc(symshen_4digit_2), Byte)

					if True == tmp5887 {
						tmp5884 := Call(__e, PrimFunc(symshen_4in_1_6), News2561)

						tmp5885 := Call(__e, PrimFunc(symshen_4byte_1_6digit), Byte)

						__e.TailApply(PrimFunc(symshen_4comb), tmp5884, tmp5885)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5888 := Call(__e, PrimFunc(symshen_4tls), V2863)

				__e.TailApply(tmp5883, tmp5888)
				return

			}, 1)

			tmp5889 := Call(__e, PrimFunc(symshen_4hds), V2863)

			tmp5890 := Call(__e, tmp5882, tmp5889)

			ifres5881 = tmp5890

		} else {
			tmp5891 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5881 = tmp5891

		}

		__e.TailApply(tmp5878, ifres5881)
		return

	}, 1)

	tmp5893 := Call(__e, ns2_1set, symshen_4_5digit_6, tmp5877)

	_ = tmp5893

	tmp5894 := MakeNative(func(__e *ControlFlow) {
		V2864 := __e.Get(1)
		_ = V2864
		__e.Return(PrimNumberSubtract(V2864, MakeNumber(48)))
		return
	}, 1)

	tmp5895 := Call(__e, ns2_1set, symshen_4byte_1_6digit, tmp5894)

	_ = tmp5895

	tmp5896 := MakeNative(func(__e *ControlFlow) {
		V2865 := __e.Get(1)
		_ = V2865
		tmp5897 := Call(__e, PrimFunc(symreverse), V2865)

		__e.TailApply(PrimFunc(symshen_4compute_1integer_1h), tmp5897, MakeNumber(0))
		return

	}, 1)

	tmp5898 := Call(__e, ns2_1set, symshen_4compute_1integer, tmp5896)

	_ = tmp5898

	tmp5899 := MakeNative(func(__e *ControlFlow) {
		V2868 := __e.Get(1)
		_ = V2868
		V2869 := __e.Get(2)
		_ = V2869
		tmp5909 := PrimEqual(Nil, V2868)

		if True == tmp5909 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp5907 := PrimIsPair(V2868)

			if True == tmp5907 {
				tmp5900 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2869)

				tmp5901 := PrimHead(V2868)

				tmp5902 := PrimNumberMultiply(tmp5900, tmp5901)

				tmp5903 := PrimTail(V2868)

				tmp5904 := PrimNumberAdd(V2869, MakeNumber(1))

				tmp5905 := Call(__e, PrimFunc(symshen_4compute_1integer_1h), tmp5903, tmp5904)

				__e.Return(PrimNumberAdd(tmp5902, tmp5905))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1integer_1h)
				return
			}

		}

	}, 2)

	tmp5910 := Call(__e, ns2_1set, symshen_4compute_1integer_1h, tmp5899)

	_ = tmp5910

	tmp5911 := MakeNative(func(__e *ControlFlow) {
		V2872 := __e.Get(1)
		_ = V2872
		V2873 := __e.Get(2)
		_ = V2873
		tmp5919 := PrimEqual(MakeNumber(0), V2873)

		if True == tmp5919 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp5917 := PrimGreatThan(V2873, MakeNumber(0))

			if True == tmp5917 {
				tmp5912 := PrimNumberSubtract(V2873, MakeNumber(1))

				tmp5913 := Call(__e, PrimFunc(symshen_4expt), V2872, tmp5912)

				__e.Return(PrimNumberMultiply(V2872, tmp5913))
				return

			} else {
				tmp5914 := PrimNumberAdd(V2873, MakeNumber(1))

				tmp5915 := Call(__e, PrimFunc(symshen_4expt), V2872, tmp5914)

				__e.Return(PrimNumberDivide(tmp5915, V2872))
				return

			}

		}

	}, 2)

	tmp5920 := Call(__e, ns2_1set, symshen_4expt, tmp5911)

	_ = tmp5920

	tmp5921 := MakeNative(func(__e *ControlFlow) {
		V2874 := __e.Get(1)
		_ = V2874
		tmp5922 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5938 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5938 {
				tmp5923 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5925 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp5925 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5926 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5stop_6 := __e.Get(1)
					_ = Parseshen_4_5stop_6
					tmp5934 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5stop_6)

					if True == tmp5934 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5927 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5fraction_6 := __e.Get(1)
							_ = Parseshen_4_5fraction_6
							tmp5931 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5fraction_6)

							if True == tmp5931 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5928 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5fraction_6)

								tmp5929 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5fraction_6)

								__e.TailApply(PrimFunc(symshen_4comb), tmp5928, tmp5929)
								return

							}

						}, 1)

						tmp5932 := Call(__e, PrimFunc(symshen_4_5fraction_6), Parseshen_4_5stop_6)

						__e.TailApply(tmp5927, tmp5932)
						return

					}

				}, 1)

				tmp5935 := Call(__e, PrimFunc(symshen_4_5stop_6), V2874)

				tmp5936 := Call(__e, tmp5926, tmp5935)

				__e.TailApply(tmp5923, tmp5936)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5939 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5integer_6 := __e.Get(1)
			_ = Parseshen_4_5integer_6
			tmp5953 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

			if True == tmp5953 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5940 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5stop_6 := __e.Get(1)
					_ = Parseshen_4_5stop_6
					tmp5950 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5stop_6)

					if True == tmp5950 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp5941 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5fraction_6 := __e.Get(1)
							_ = Parseshen_4_5fraction_6
							tmp5947 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5fraction_6)

							if True == tmp5947 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp5942 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5fraction_6)

								tmp5943 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5integer_6)

								tmp5944 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5fraction_6)

								tmp5945 := PrimNumberAdd(tmp5943, tmp5944)

								__e.TailApply(PrimFunc(symshen_4comb), tmp5942, tmp5945)
								return

							}

						}, 1)

						tmp5948 := Call(__e, PrimFunc(symshen_4_5fraction_6), Parseshen_4_5stop_6)

						__e.TailApply(tmp5941, tmp5948)
						return

					}

				}, 1)

				tmp5951 := Call(__e, PrimFunc(symshen_4_5stop_6), Parseshen_4_5integer_6)

				__e.TailApply(tmp5940, tmp5951)
				return

			}

		}, 1)

		tmp5954 := Call(__e, PrimFunc(symshen_4_5integer_6), V2874)

		tmp5955 := Call(__e, tmp5939, tmp5954)

		__e.TailApply(tmp5922, tmp5955)
		return

	}, 1)

	tmp5956 := Call(__e, ns2_1set, symshen_4_5float_6, tmp5921)

	_ = tmp5956

	tmp5957 := MakeNative(func(__e *ControlFlow) {
		V2875 := __e.Get(1)
		_ = V2875
		tmp5958 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5960 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5960 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5967 := Call(__e, PrimFunc(symshen_4_ahd_2), V2875, MakeNumber(46))

		var ifres5961 Obj

		if True == tmp5967 {
			tmp5962 := MakeNative(func(__e *ControlFlow) {
				News2564 := __e.Get(1)
				_ = News2564
				tmp5963 := Call(__e, PrimFunc(symshen_4in_1_6), News2564)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5963, symshen_4skip)
				return

			}, 1)

			tmp5964 := Call(__e, PrimFunc(symshen_4tls), V2875)

			tmp5965 := Call(__e, tmp5962, tmp5964)

			ifres5961 = tmp5965

		} else {
			tmp5966 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres5961 = tmp5966

		}

		__e.TailApply(tmp5958, ifres5961)
		return

	}, 1)

	tmp5968 := Call(__e, ns2_1set, symshen_4_5stop_6, tmp5957)

	_ = tmp5968

	tmp5969 := MakeNative(func(__e *ControlFlow) {
		V2876 := __e.Get(1)
		_ = V2876
		tmp5970 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5972 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp5972 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5973 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5digits_6 := __e.Get(1)
			_ = Parseshen_4_5digits_6
			tmp5978 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5digits_6)

			if True == tmp5978 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp5974 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5digits_6)

				tmp5975 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5digits_6)

				tmp5976 := Call(__e, PrimFunc(symshen_4compute_1fraction), tmp5975)

				__e.TailApply(PrimFunc(symshen_4comb), tmp5974, tmp5976)
				return

			}

		}, 1)

		tmp5979 := Call(__e, PrimFunc(symshen_4_5digits_6), V2876)

		tmp5980 := Call(__e, tmp5973, tmp5979)

		__e.TailApply(tmp5970, tmp5980)
		return

	}, 1)

	tmp5981 := Call(__e, ns2_1set, symshen_4_5fraction_6, tmp5969)

	_ = tmp5981

	tmp5982 := MakeNative(func(__e *ControlFlow) {
		V2877 := __e.Get(1)
		_ = V2877
		__e.TailApply(PrimFunc(symshen_4compute_1fraction_1h), V2877, MakeNumber(-1))
		return
	}, 1)

	tmp5983 := Call(__e, ns2_1set, symshen_4compute_1fraction, tmp5982)

	_ = tmp5983

	tmp5984 := MakeNative(func(__e *ControlFlow) {
		V2880 := __e.Get(1)
		_ = V2880
		V2881 := __e.Get(2)
		_ = V2881
		tmp5994 := PrimEqual(Nil, V2880)

		if True == tmp5994 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp5992 := PrimIsPair(V2880)

			if True == tmp5992 {
				tmp5985 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2881)

				tmp5986 := PrimHead(V2880)

				tmp5987 := PrimNumberMultiply(tmp5985, tmp5986)

				tmp5988 := PrimTail(V2880)

				tmp5989 := PrimNumberSubtract(V2881, MakeNumber(1))

				tmp5990 := Call(__e, PrimFunc(symshen_4compute_1fraction_1h), tmp5988, tmp5989)

				__e.Return(PrimNumberAdd(tmp5987, tmp5990))
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4compute_1fraction_1h)
				return
			}

		}

	}, 2)

	tmp5995 := Call(__e, ns2_1set, symshen_4compute_1fraction_1h, tmp5984)

	_ = tmp5995

	tmp5996 := MakeNative(func(__e *ControlFlow) {
		V2882 := __e.Get(1)
		_ = V2882
		tmp5997 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp6019 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp6019 {
				tmp5998 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp6000 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp6000 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp6001 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5integer_6 := __e.Get(1)
					_ = Parseshen_4_5integer_6
					tmp6015 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

					if True == tmp6015 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp6002 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5lowE_6 := __e.Get(1)
							_ = Parseshen_4_5lowE_6
							tmp6012 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5lowE_6)

							if True == tmp6012 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp6003 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5log10_6 := __e.Get(1)
									_ = Parseshen_4_5log10_6
									tmp6009 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

									if True == tmp6009 {
										__e.TailApply(PrimFunc(symshen_4parse_1failure))
										return
									} else {
										tmp6004 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5log10_6)

										tmp6005 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5integer_6)

										tmp6006 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5log10_6)

										tmp6007 := Call(__e, PrimFunc(symshen_4compute_1E), tmp6005, tmp6006)

										__e.TailApply(PrimFunc(symshen_4comb), tmp6004, tmp6007)
										return

									}

								}, 1)

								tmp6010 := Call(__e, PrimFunc(symshen_4_5log10_6), Parseshen_4_5lowE_6)

								__e.TailApply(tmp6003, tmp6010)
								return

							}

						}, 1)

						tmp6013 := Call(__e, PrimFunc(symshen_4_5lowE_6), Parseshen_4_5integer_6)

						__e.TailApply(tmp6002, tmp6013)
						return

					}

				}, 1)

				tmp6016 := Call(__e, PrimFunc(symshen_4_5integer_6), V2882)

				tmp6017 := Call(__e, tmp6001, tmp6016)

				__e.TailApply(tmp5998, tmp6017)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp6020 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5float_6 := __e.Get(1)
			_ = Parseshen_4_5float_6
			tmp6034 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5float_6)

			if True == tmp6034 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp6021 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5lowE_6 := __e.Get(1)
					_ = Parseshen_4_5lowE_6
					tmp6031 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5lowE_6)

					if True == tmp6031 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp6022 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5log10_6 := __e.Get(1)
							_ = Parseshen_4_5log10_6
							tmp6028 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

							if True == tmp6028 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp6023 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5log10_6)

								tmp6024 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5float_6)

								tmp6025 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5log10_6)

								tmp6026 := Call(__e, PrimFunc(symshen_4compute_1E), tmp6024, tmp6025)

								__e.TailApply(PrimFunc(symshen_4comb), tmp6023, tmp6026)
								return

							}

						}, 1)

						tmp6029 := Call(__e, PrimFunc(symshen_4_5log10_6), Parseshen_4_5lowE_6)

						__e.TailApply(tmp6022, tmp6029)
						return

					}

				}, 1)

				tmp6032 := Call(__e, PrimFunc(symshen_4_5lowE_6), Parseshen_4_5float_6)

				__e.TailApply(tmp6021, tmp6032)
				return

			}

		}, 1)

		tmp6035 := Call(__e, PrimFunc(symshen_4_5float_6), V2882)

		tmp6036 := Call(__e, tmp6020, tmp6035)

		__e.TailApply(tmp5997, tmp6036)
		return

	}, 1)

	tmp6037 := Call(__e, ns2_1set, symshen_4_5e_1number_6, tmp5996)

	_ = tmp6037

	tmp6038 := MakeNative(func(__e *ControlFlow) {
		V2883 := __e.Get(1)
		_ = V2883
		tmp6039 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp6066 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp6066 {
				tmp6040 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp6052 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp6052 {
						tmp6041 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp6043 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

							if True == tmp6043 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp6044 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5integer_6 := __e.Get(1)
							_ = Parseshen_4_5integer_6
							tmp6048 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

							if True == tmp6048 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp6045 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5integer_6)

								tmp6046 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5integer_6)

								__e.TailApply(PrimFunc(symshen_4comb), tmp6045, tmp6046)
								return

							}

						}, 1)

						tmp6049 := Call(__e, PrimFunc(symshen_4_5integer_6), V2883)

						tmp6050 := Call(__e, tmp6044, tmp6049)

						__e.TailApply(tmp6041, tmp6050)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp6053 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5minus_6 := __e.Get(1)
					_ = Parseshen_4_5minus_6
					tmp6062 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5minus_6)

					if True == tmp6062 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp6054 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5log10_6 := __e.Get(1)
							_ = Parseshen_4_5log10_6
							tmp6059 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

							if True == tmp6059 {
								__e.TailApply(PrimFunc(symshen_4parse_1failure))
								return
							} else {
								tmp6055 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5log10_6)

								tmp6056 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5log10_6)

								tmp6057 := PrimNumberSubtract(MakeNumber(0), tmp6056)

								__e.TailApply(PrimFunc(symshen_4comb), tmp6055, tmp6057)
								return

							}

						}, 1)

						tmp6060 := Call(__e, PrimFunc(symshen_4_5log10_6), Parseshen_4_5minus_6)

						__e.TailApply(tmp6054, tmp6060)
						return

					}

				}, 1)

				tmp6063 := Call(__e, PrimFunc(symshen_4_5minus_6), V2883)

				tmp6064 := Call(__e, tmp6053, tmp6063)

				__e.TailApply(tmp6040, tmp6064)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp6067 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5plus_6 := __e.Get(1)
			_ = Parseshen_4_5plus_6
			tmp6075 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5plus_6)

			if True == tmp6075 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp6068 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5log10_6 := __e.Get(1)
					_ = Parseshen_4_5log10_6
					tmp6072 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

					if True == tmp6072 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp6069 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5log10_6)

						tmp6070 := Call(__e, PrimFunc(symshen_4_5_1out), Parseshen_4_5log10_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp6069, tmp6070)
						return

					}

				}, 1)

				tmp6073 := Call(__e, PrimFunc(symshen_4_5log10_6), Parseshen_4_5plus_6)

				__e.TailApply(tmp6068, tmp6073)
				return

			}

		}, 1)

		tmp6076 := Call(__e, PrimFunc(symshen_4_5plus_6), V2883)

		tmp6077 := Call(__e, tmp6067, tmp6076)

		__e.TailApply(tmp6039, tmp6077)
		return

	}, 1)

	tmp6078 := Call(__e, ns2_1set, symshen_4_5log10_6, tmp6038)

	_ = tmp6078

	tmp6079 := MakeNative(func(__e *ControlFlow) {
		V2884 := __e.Get(1)
		_ = V2884
		tmp6080 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp6082 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp6082 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp6089 := Call(__e, PrimFunc(symshen_4_ahd_2), V2884, MakeNumber(101))

		var ifres6083 Obj

		if True == tmp6089 {
			tmp6084 := MakeNative(func(__e *ControlFlow) {
				News2569 := __e.Get(1)
				_ = News2569
				tmp6085 := Call(__e, PrimFunc(symshen_4in_1_6), News2569)

				__e.TailApply(PrimFunc(symshen_4comb), tmp6085, symshen_4skip)
				return

			}, 1)

			tmp6086 := Call(__e, PrimFunc(symshen_4tls), V2884)

			tmp6087 := Call(__e, tmp6084, tmp6086)

			ifres6083 = tmp6087

		} else {
			tmp6088 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres6083 = tmp6088

		}

		__e.TailApply(tmp6080, ifres6083)
		return

	}, 1)

	tmp6090 := Call(__e, ns2_1set, symshen_4_5lowE_6, tmp6079)

	_ = tmp6090

	tmp6091 := MakeNative(func(__e *ControlFlow) {
		V2885 := __e.Get(1)
		_ = V2885
		V2886 := __e.Get(2)
		_ = V2886
		tmp6092 := Call(__e, PrimFunc(symshen_4expt), MakeNumber(10), V2886)

		__e.Return(PrimNumberMultiply(V2885, tmp6092))
		return

	}, 2)

	tmp6093 := Call(__e, ns2_1set, symshen_4compute_1E, tmp6091)

	_ = tmp6093

	tmp6094 := MakeNative(func(__e *ControlFlow) {
		V2887 := __e.Get(1)
		_ = V2887
		tmp6095 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp6106 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp6106 {
				tmp6096 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp6098 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

					if True == tmp6098 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp6099 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5whitespace_6 := __e.Get(1)
					_ = Parseshen_4_5whitespace_6
					tmp6102 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5whitespace_6)

					if True == tmp6102 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp6100 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5whitespace_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp6100, symshen_4skip)
						return

					}

				}, 1)

				tmp6103 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2887)

				tmp6104 := Call(__e, tmp6099, tmp6103)

				__e.TailApply(tmp6096, tmp6104)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp6107 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5whitespace_6 := __e.Get(1)
			_ = Parseshen_4_5whitespace_6
			tmp6114 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5whitespace_6)

			if True == tmp6114 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				tmp6108 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5whitespaces_6 := __e.Get(1)
					_ = Parseshen_4_5whitespaces_6
					tmp6111 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Parseshen_4_5whitespaces_6)

					if True == tmp6111 {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					} else {
						tmp6109 := Call(__e, PrimFunc(symshen_4in_1_6), Parseshen_4_5whitespaces_6)

						__e.TailApply(PrimFunc(symshen_4comb), tmp6109, symshen_4skip)
						return

					}

				}, 1)

				tmp6112 := Call(__e, PrimFunc(symshen_4_5whitespaces_6), Parseshen_4_5whitespace_6)

				__e.TailApply(tmp6108, tmp6112)
				return

			}

		}, 1)

		tmp6115 := Call(__e, PrimFunc(symshen_4_5whitespace_6), V2887)

		tmp6116 := Call(__e, tmp6107, tmp6115)

		__e.TailApply(tmp6095, tmp6116)
		return

	}, 1)

	tmp6117 := Call(__e, ns2_1set, symshen_4_5whitespaces_6, tmp6094)

	_ = tmp6117

	tmp6118 := MakeNative(func(__e *ControlFlow) {
		V2888 := __e.Get(1)
		_ = V2888
		tmp6119 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp6121 := Call(__e, PrimFunc(symshen_4parse_1failure_2), Result)

			if True == tmp6121 {
				__e.TailApply(PrimFunc(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp6132 := Call(__e, PrimFunc(symshen_4non_1empty_1stream_2), V2888)

		var ifres6122 Obj

		if True == tmp6132 {
			tmp6123 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp6124 := MakeNative(func(__e *ControlFlow) {
					News2572 := __e.Get(1)
					_ = News2572
					tmp6127 := Call(__e, PrimFunc(symshen_4whitespace_2), Byte)

					if True == tmp6127 {
						tmp6125 := Call(__e, PrimFunc(symshen_4in_1_6), News2572)

						__e.TailApply(PrimFunc(symshen_4comb), tmp6125, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimFunc(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp6128 := Call(__e, PrimFunc(symshen_4tls), V2888)

				__e.TailApply(tmp6124, tmp6128)
				return

			}, 1)

			tmp6129 := Call(__e, PrimFunc(symshen_4hds), V2888)

			tmp6130 := Call(__e, tmp6123, tmp6129)

			ifres6122 = tmp6130

		} else {
			tmp6131 := Call(__e, PrimFunc(symshen_4parse_1failure))

			ifres6122 = tmp6131

		}

		__e.TailApply(tmp6119, ifres6122)
		return

	}, 1)

	tmp6133 := Call(__e, ns2_1set, symshen_4_5whitespace_6, tmp6118)

	_ = tmp6133

	tmp6134 := MakeNative(func(__e *ControlFlow) {
		V2891 := __e.Get(1)
		_ = V2891
		tmp6142 := PrimEqual(MakeNumber(32), V2891)

		if True == tmp6142 {
			__e.Return(True)
			return
		} else {
			tmp6140 := PrimEqual(MakeNumber(13), V2891)

			if True == tmp6140 {
				__e.Return(True)
				return
			} else {
				tmp6138 := PrimEqual(MakeNumber(10), V2891)

				if True == tmp6138 {
					__e.Return(True)
					return
				} else {
					tmp6136 := PrimEqual(MakeNumber(9), V2891)

					if True == tmp6136 {
						__e.Return(True)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}

		}

	}, 1)

	tmp6143 := Call(__e, ns2_1set, symshen_4whitespace_2, tmp6134)

	_ = tmp6143

	tmp6144 := MakeNative(func(__e *ControlFlow) {
		V2892 := __e.Get(1)
		_ = V2892
		tmp6167 := PrimEqual(Nil, V2892)

		if True == tmp6167 {
			__e.Return(Nil)
			return
		} else {
			tmp6165 := PrimIsPair(V2892)

			var ifres6161 Obj

			if True == tmp6165 {
				tmp6163 := PrimHead(V2892)

				tmp6164 := Call(__e, PrimFunc(symshen_4packaged_2), tmp6163)

				var ifres6162 Obj

				if True == tmp6164 {
					ifres6162 = True

				} else {
					ifres6162 = False

				}

				ifres6161 = ifres6162

			} else {
				ifres6161 = False

			}

			if True == ifres6161 {
				tmp6145 := PrimHead(V2892)

				tmp6146 := Call(__e, PrimFunc(symshen_4unpackage), tmp6145)

				tmp6147 := PrimTail(V2892)

				tmp6148 := Call(__e, PrimFunc(symappend), tmp6146, tmp6147)

				__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6148)
				return

			} else {
				tmp6159 := PrimIsPair(V2892)

				if True == tmp6159 {
					tmp6149 := MakeNative(func(__e *ControlFlow) {
						M := __e.Get(1)
						_ = M
						tmp6155 := Call(__e, PrimFunc(symshen_4packaged_2), M)

						if True == tmp6155 {
							tmp6150 := PrimTail(V2892)

							tmp6151 := PrimCons(M, tmp6150)

							__e.TailApply(PrimFunc(symshen_4unpackage_emacroexpand), tmp6151)
							return

						} else {
							tmp6152 := PrimTail(V2892)

							tmp6153 := Call(__e, PrimFunc(symshen_4unpackage_emacroexpand), tmp6152)

							__e.Return(PrimCons(M, tmp6153))
							return

						}

					}, 1)

					tmp6156 := PrimHead(V2892)

					tmp6157 := Call(__e, PrimFunc(symmacroexpand), tmp6156)

					__e.TailApply(tmp6149, tmp6157)
					return

				} else {
					__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage_emacroexpand)
					return
				}

			}

		}

	}, 1)

	tmp6168 := Call(__e, ns2_1set, symshen_4unpackage_emacroexpand, tmp6144)

	_ = tmp6168

	tmp6169 := MakeNative(func(__e *ControlFlow) {
		V2895 := __e.Get(1)
		_ = V2895
		tmp6184 := PrimIsPair(V2895)

		var ifres6171 Obj

		if True == tmp6184 {
			tmp6182 := PrimHead(V2895)

			tmp6183 := PrimEqual(sympackage, tmp6182)

			var ifres6173 Obj

			if True == tmp6183 {
				tmp6180 := PrimTail(V2895)

				tmp6181 := PrimIsPair(tmp6180)

				var ifres6175 Obj

				if True == tmp6181 {
					tmp6177 := PrimTail(V2895)

					tmp6178 := PrimTail(tmp6177)

					tmp6179 := PrimIsPair(tmp6178)

					var ifres6176 Obj

					if True == tmp6179 {
						ifres6176 = True

					} else {
						ifres6176 = False

					}

					ifres6175 = ifres6176

				} else {
					ifres6175 = False

				}

				var ifres6174 Obj

				if True == ifres6175 {
					ifres6174 = True

				} else {
					ifres6174 = False

				}

				ifres6173 = ifres6174

			} else {
				ifres6173 = False

			}

			var ifres6172 Obj

			if True == ifres6173 {
				ifres6172 = True

			} else {
				ifres6172 = False

			}

			ifres6171 = ifres6172

		} else {
			ifres6171 = False

		}

		if True == ifres6171 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp6185 := Call(__e, ns2_1set, symshen_4packaged_2, tmp6169)

	_ = tmp6185

	tmp6186 := MakeNative(func(__e *ControlFlow) {
		V2898 := __e.Get(1)
		_ = V2898
		tmp6240 := PrimIsPair(V2898)

		var ifres6222 Obj

		if True == tmp6240 {
			tmp6238 := PrimHead(V2898)

			tmp6239 := PrimEqual(sympackage, tmp6238)

			var ifres6224 Obj

			if True == tmp6239 {
				tmp6236 := PrimTail(V2898)

				tmp6237 := PrimIsPair(tmp6236)

				var ifres6226 Obj

				if True == tmp6237 {
					tmp6233 := PrimTail(V2898)

					tmp6234 := PrimHead(tmp6233)

					tmp6235 := PrimEqual(symnull, tmp6234)

					var ifres6228 Obj

					if True == tmp6235 {
						tmp6230 := PrimTail(V2898)

						tmp6231 := PrimTail(tmp6230)

						tmp6232 := PrimIsPair(tmp6231)

						var ifres6229 Obj

						if True == tmp6232 {
							ifres6229 = True

						} else {
							ifres6229 = False

						}

						ifres6228 = ifres6229

					} else {
						ifres6228 = False

					}

					var ifres6227 Obj

					if True == ifres6228 {
						ifres6227 = True

					} else {
						ifres6227 = False

					}

					ifres6226 = ifres6227

				} else {
					ifres6226 = False

				}

				var ifres6225 Obj

				if True == ifres6226 {
					ifres6225 = True

				} else {
					ifres6225 = False

				}

				ifres6224 = ifres6225

			} else {
				ifres6224 = False

			}

			var ifres6223 Obj

			if True == ifres6224 {
				ifres6223 = True

			} else {
				ifres6223 = False

			}

			ifres6222 = ifres6223

		} else {
			ifres6222 = False

		}

		if True == ifres6222 {
			tmp6187 := PrimTail(V2898)

			tmp6188 := PrimTail(tmp6187)

			__e.Return(PrimTail(tmp6188))
			return

		} else {
			tmp6220 := PrimIsPair(V2898)

			var ifres6207 Obj

			if True == tmp6220 {
				tmp6218 := PrimHead(V2898)

				tmp6219 := PrimEqual(sympackage, tmp6218)

				var ifres6209 Obj

				if True == tmp6219 {
					tmp6216 := PrimTail(V2898)

					tmp6217 := PrimIsPair(tmp6216)

					var ifres6211 Obj

					if True == tmp6217 {
						tmp6213 := PrimTail(V2898)

						tmp6214 := PrimTail(tmp6213)

						tmp6215 := PrimIsPair(tmp6214)

						var ifres6212 Obj

						if True == tmp6215 {
							ifres6212 = True

						} else {
							ifres6212 = False

						}

						ifres6211 = ifres6212

					} else {
						ifres6211 = False

					}

					var ifres6210 Obj

					if True == ifres6211 {
						ifres6210 = True

					} else {
						ifres6210 = False

					}

					ifres6209 = ifres6210

				} else {
					ifres6209 = False

				}

				var ifres6208 Obj

				if True == ifres6209 {
					ifres6208 = True

				} else {
					ifres6208 = False

				}

				ifres6207 = ifres6208

			} else {
				ifres6207 = False

			}

			if True == ifres6207 {
				tmp6189 := MakeNative(func(__e *ControlFlow) {
					External_b := __e.Get(1)
					_ = External_b
					tmp6190 := MakeNative(func(__e *ControlFlow) {
						Package := __e.Get(1)
						_ = Package
						tmp6191 := MakeNative(func(__e *ControlFlow) {
							RecordExternal := __e.Get(1)
							_ = RecordExternal
							__e.Return(Package)
							return
						}, 1)

						tmp6192 := PrimTail(V2898)

						tmp6193 := PrimHead(tmp6192)

						tmp6194 := Call(__e, PrimFunc(symshen_4record_1external), tmp6193, External_b)

						__e.TailApply(tmp6191, tmp6194)
						return

					}, 1)

					tmp6195 := PrimTail(V2898)

					tmp6196 := PrimHead(tmp6195)

					tmp6197 := PrimStr(tmp6196)

					tmp6198 := PrimTail(V2898)

					tmp6199 := PrimTail(tmp6198)

					tmp6200 := PrimTail(tmp6199)

					tmp6201 := Call(__e, PrimFunc(symshen_4package_1symbols), tmp6197, External_b, tmp6200)

					__e.TailApply(tmp6190, tmp6201)
					return

				}, 1)

				tmp6202 := PrimTail(V2898)

				tmp6203 := PrimTail(tmp6202)

				tmp6204 := PrimHead(tmp6203)

				tmp6205 := Call(__e, PrimFunc(symeval), tmp6204)

				__e.TailApply(tmp6189, tmp6205)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4unpackage)
				return
			}

		}

	}, 1)

	tmp6241 := Call(__e, ns2_1set, symshen_4unpackage, tmp6186)

	_ = tmp6241

	tmp6242 := MakeNative(func(__e *ControlFlow) {
		V2899 := __e.Get(1)
		_ = V2899
		V2900 := __e.Get(2)
		_ = V2900
		tmp6243 := MakeNative(func(__e *ControlFlow) {
			External := __e.Get(1)
			_ = External
			tmp6244 := Call(__e, PrimFunc(symunion), V2900, External)

			tmp6245 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symput), V2899, symshen_4external_1symbols, tmp6244, tmp6245)
			return

		}, 1)

		tmp6246 := MakeNative(func(__e *ControlFlow) {
			tmp6247 := PrimValue(sym_dproperty_1vector_d)

			__e.TailApply(PrimFunc(symget), V2899, symshen_4external_1symbols, tmp6247)
			return

		}, 0)

		tmp6248 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp6249 := Call(__e, try_1catch, tmp6246, tmp6248)

		__e.TailApply(tmp6243, tmp6249)
		return

	}, 2)

	tmp6250 := Call(__e, ns2_1set, symshen_4record_1external, tmp6242)

	_ = tmp6250

	tmp6251 := MakeNative(func(__e *ControlFlow) {
		V2905 := __e.Get(1)
		_ = V2905
		V2906 := __e.Get(2)
		_ = V2906
		V2907 := __e.Get(3)
		_ = V2907
		tmp6256 := PrimIsPair(V2907)

		if True == tmp6256 {
			tmp6252 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimFunc(symshen_4package_1symbols), V2905, V2906, X)
				return
			}, 1)

			__e.TailApply(PrimFunc(symmap), tmp6252, V2907)
			return

		} else {
			tmp6254 := Call(__e, PrimFunc(symshen_4internal_2), V2907, V2905, V2906)

			if True == tmp6254 {
				__e.TailApply(PrimFunc(symshen_4intern_1in_1package), V2905, V2907)
				return
			} else {
				__e.Return(V2907)
				return
			}

		}

	}, 3)

	tmp6257 := Call(__e, ns2_1set, symshen_4package_1symbols, tmp6251)

	_ = tmp6257

	tmp6258 := MakeNative(func(__e *ControlFlow) {
		V2908 := __e.Get(1)
		_ = V2908
		V2909 := __e.Get(2)
		_ = V2909
		tmp6259 := PrimStr(V2909)

		tmp6260 := Call(__e, PrimFunc(sym_8s), MakeString("."), tmp6259)

		tmp6261 := Call(__e, PrimFunc(sym_8s), V2908, tmp6260)

		__e.Return(PrimIntern(tmp6261))
		return

	}, 2)

	tmp6262 := Call(__e, ns2_1set, symshen_4intern_1in_1package, tmp6258)

	_ = tmp6262

	tmp6263 := MakeNative(func(__e *ControlFlow) {
		V2910 := __e.Get(1)
		_ = V2910
		V2911 := __e.Get(2)
		_ = V2911
		V2912 := __e.Get(3)
		_ = V2912
		tmp6293 := Call(__e, PrimFunc(symelement_2), V2910, V2912)

		tmp6294 := PrimNot(tmp6293)

		if True == tmp6294 {
			tmp6290 := Call(__e, PrimFunc(symshen_4sng_2), V2910)

			tmp6291 := PrimNot(tmp6290)

			var ifres6265 Obj

			if True == tmp6291 {
				tmp6288 := Call(__e, PrimFunc(symshen_4dbl_2), V2910)

				tmp6289 := PrimNot(tmp6288)

				var ifres6267 Obj

				if True == tmp6289 {
					tmp6287 := PrimIsSymbol(V2910)

					var ifres6269 Obj

					if True == tmp6287 {
						tmp6285 := Call(__e, PrimFunc(symshen_4sysfunc_2), V2910)

						tmp6286 := PrimNot(tmp6285)

						var ifres6271 Obj

						if True == tmp6286 {
							tmp6283 := PrimIsVariable(V2910)

							tmp6284 := PrimNot(tmp6283)

							var ifres6273 Obj

							if True == tmp6284 {
								tmp6280 := PrimStr(V2910)

								tmp6281 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6280)

								tmp6282 := PrimNot(tmp6281)

								var ifres6275 Obj

								if True == tmp6282 {
									tmp6277 := PrimStr(V2910)

									tmp6278 := Call(__e, PrimFunc(symshen_4internal_1to_1P_2), V2911, tmp6277)

									tmp6279 := PrimNot(tmp6278)

									var ifres6276 Obj

									if True == tmp6279 {
										ifres6276 = True

									} else {
										ifres6276 = False

									}

									ifres6275 = ifres6276

								} else {
									ifres6275 = False

								}

								var ifres6274 Obj

								if True == ifres6275 {
									ifres6274 = True

								} else {
									ifres6274 = False

								}

								ifres6273 = ifres6274

							} else {
								ifres6273 = False

							}

							var ifres6272 Obj

							if True == ifres6273 {
								ifres6272 = True

							} else {
								ifres6272 = False

							}

							ifres6271 = ifres6272

						} else {
							ifres6271 = False

						}

						var ifres6270 Obj

						if True == ifres6271 {
							ifres6270 = True

						} else {
							ifres6270 = False

						}

						ifres6269 = ifres6270

					} else {
						ifres6269 = False

					}

					var ifres6268 Obj

					if True == ifres6269 {
						ifres6268 = True

					} else {
						ifres6268 = False

					}

					ifres6267 = ifres6268

				} else {
					ifres6267 = False

				}

				var ifres6266 Obj

				if True == ifres6267 {
					ifres6266 = True

				} else {
					ifres6266 = False

				}

				ifres6265 = ifres6266

			} else {
				ifres6265 = False

			}

			if True == ifres6265 {
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

	}, 3)

	tmp6295 := Call(__e, ns2_1set, symshen_4internal_2, tmp6263)

	_ = tmp6295

	tmp6296 := MakeNative(func(__e *ControlFlow) {
		V2917 := __e.Get(1)
		_ = V2917
		tmp6350 := Call(__e, PrimFunc(symshen_4_7string_2), V2917)

		var ifres6298 Obj

		if True == tmp6350 {
			tmp6348 := Call(__e, PrimFunc(symhdstr), V2917)

			tmp6349 := PrimEqual(MakeString("s"), tmp6348)

			var ifres6300 Obj

			if True == tmp6349 {
				tmp6346 := PrimTailString(V2917)

				tmp6347 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6346)

				var ifres6302 Obj

				if True == tmp6347 {
					tmp6343 := PrimTailString(V2917)

					tmp6344 := Call(__e, PrimFunc(symhdstr), tmp6343)

					tmp6345 := PrimEqual(MakeString("h"), tmp6344)

					var ifres6304 Obj

					if True == tmp6345 {
						tmp6340 := PrimTailString(V2917)

						tmp6341 := PrimTailString(tmp6340)

						tmp6342 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6341)

						var ifres6306 Obj

						if True == tmp6342 {
							tmp6336 := PrimTailString(V2917)

							tmp6337 := PrimTailString(tmp6336)

							tmp6338 := Call(__e, PrimFunc(symhdstr), tmp6337)

							tmp6339 := PrimEqual(MakeString("e"), tmp6338)

							var ifres6308 Obj

							if True == tmp6339 {
								tmp6332 := PrimTailString(V2917)

								tmp6333 := PrimTailString(tmp6332)

								tmp6334 := PrimTailString(tmp6333)

								tmp6335 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6334)

								var ifres6310 Obj

								if True == tmp6335 {
									tmp6327 := PrimTailString(V2917)

									tmp6328 := PrimTailString(tmp6327)

									tmp6329 := PrimTailString(tmp6328)

									tmp6330 := Call(__e, PrimFunc(symhdstr), tmp6329)

									tmp6331 := PrimEqual(MakeString("n"), tmp6330)

									var ifres6312 Obj

									if True == tmp6331 {
										tmp6322 := PrimTailString(V2917)

										tmp6323 := PrimTailString(tmp6322)

										tmp6324 := PrimTailString(tmp6323)

										tmp6325 := PrimTailString(tmp6324)

										tmp6326 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6325)

										var ifres6314 Obj

										if True == tmp6326 {
											tmp6316 := PrimTailString(V2917)

											tmp6317 := PrimTailString(tmp6316)

											tmp6318 := PrimTailString(tmp6317)

											tmp6319 := PrimTailString(tmp6318)

											tmp6320 := Call(__e, PrimFunc(symhdstr), tmp6319)

											tmp6321 := PrimEqual(MakeString("."), tmp6320)

											var ifres6315 Obj

											if True == tmp6321 {
												ifres6315 = True

											} else {
												ifres6315 = False

											}

											ifres6314 = ifres6315

										} else {
											ifres6314 = False

										}

										var ifres6313 Obj

										if True == ifres6314 {
											ifres6313 = True

										} else {
											ifres6313 = False

										}

										ifres6312 = ifres6313

									} else {
										ifres6312 = False

									}

									var ifres6311 Obj

									if True == ifres6312 {
										ifres6311 = True

									} else {
										ifres6311 = False

									}

									ifres6310 = ifres6311

								} else {
									ifres6310 = False

								}

								var ifres6309 Obj

								if True == ifres6310 {
									ifres6309 = True

								} else {
									ifres6309 = False

								}

								ifres6308 = ifres6309

							} else {
								ifres6308 = False

							}

							var ifres6307 Obj

							if True == ifres6308 {
								ifres6307 = True

							} else {
								ifres6307 = False

							}

							ifres6306 = ifres6307

						} else {
							ifres6306 = False

						}

						var ifres6305 Obj

						if True == ifres6306 {
							ifres6305 = True

						} else {
							ifres6305 = False

						}

						ifres6304 = ifres6305

					} else {
						ifres6304 = False

					}

					var ifres6303 Obj

					if True == ifres6304 {
						ifres6303 = True

					} else {
						ifres6303 = False

					}

					ifres6302 = ifres6303

				} else {
					ifres6302 = False

				}

				var ifres6301 Obj

				if True == ifres6302 {
					ifres6301 = True

				} else {
					ifres6301 = False

				}

				ifres6300 = ifres6301

			} else {
				ifres6300 = False

			}

			var ifres6299 Obj

			if True == ifres6300 {
				ifres6299 = True

			} else {
				ifres6299 = False

			}

			ifres6298 = ifres6299

		} else {
			ifres6298 = False

		}

		if True == ifres6298 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp6351 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp6296)

	_ = tmp6351

	tmp6352 := MakeNative(func(__e *ControlFlow) {
		V2918 := __e.Get(1)
		_ = V2918
		tmp6353 := PrimValue(sym_dproperty_1vector_d)

		tmp6354 := Call(__e, PrimFunc(symget), symshen, symshen_4external_1symbols, tmp6353)

		__e.TailApply(PrimFunc(symelement_2), V2918, tmp6354)
		return

	}, 1)

	tmp6355 := Call(__e, ns2_1set, symshen_4sysfunc_2, tmp6352)

	_ = tmp6355

	tmp6356 := MakeNative(func(__e *ControlFlow) {
		V2926 := __e.Get(1)
		_ = V2926
		V2927 := __e.Get(2)
		_ = V2927
		tmp6377 := PrimEqual(MakeString(""), V2926)

		var ifres6370 Obj

		if True == tmp6377 {
			tmp6376 := Call(__e, PrimFunc(symshen_4_7string_2), V2927)

			var ifres6372 Obj

			if True == tmp6376 {
				tmp6374 := Call(__e, PrimFunc(symhdstr), V2927)

				tmp6375 := PrimEqual(MakeString("."), tmp6374)

				var ifres6373 Obj

				if True == tmp6375 {
					ifres6373 = True

				} else {
					ifres6373 = False

				}

				ifres6372 = ifres6373

			} else {
				ifres6372 = False

			}

			var ifres6371 Obj

			if True == ifres6372 {
				ifres6371 = True

			} else {
				ifres6371 = False

			}

			ifres6370 = ifres6371

		} else {
			ifres6370 = False

		}

		if True == ifres6370 {
			__e.Return(True)
			return
		} else {
			tmp6368 := Call(__e, PrimFunc(symshen_4_7string_2), V2926)

			var ifres6360 Obj

			if True == tmp6368 {
				tmp6367 := Call(__e, PrimFunc(symshen_4_7string_2), V2927)

				var ifres6362 Obj

				if True == tmp6367 {
					tmp6364 := Call(__e, PrimFunc(symhdstr), V2926)

					tmp6365 := Call(__e, PrimFunc(symhdstr), V2927)

					tmp6366 := PrimEqual(tmp6364, tmp6365)

					var ifres6363 Obj

					if True == tmp6366 {
						ifres6363 = True

					} else {
						ifres6363 = False

					}

					ifres6362 = ifres6363

				} else {
					ifres6362 = False

				}

				var ifres6361 Obj

				if True == ifres6362 {
					ifres6361 = True

				} else {
					ifres6361 = False

				}

				ifres6360 = ifres6361

			} else {
				ifres6360 = False

			}

			if True == ifres6360 {
				tmp6357 := PrimTailString(V2926)

				tmp6358 := PrimTailString(V2927)

				__e.TailApply(PrimFunc(symshen_4internal_1to_1P_2), tmp6357, tmp6358)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp6378 := Call(__e, ns2_1set, symshen_4internal_1to_1P_2, tmp6356)

	_ = tmp6378

	tmp6379 := MakeNative(func(__e *ControlFlow) {
		V2930 := __e.Get(1)
		_ = V2930
		V2931 := __e.Get(2)
		_ = V2931
		tmp6392 := Call(__e, PrimFunc(symelement_2), V2930, V2931)

		if True == tmp6392 {
			__e.Return(V2930)
			return
		} else {
			tmp6390 := PrimIsPair(V2930)

			var ifres6386 Obj

			if True == tmp6390 {
				tmp6388 := PrimHead(V2930)

				tmp6389 := Call(__e, PrimFunc(symshen_4non_1application_2), tmp6388)

				var ifres6387 Obj

				if True == tmp6389 {
					ifres6387 = True

				} else {
					ifres6387 = False

				}

				ifres6386 = ifres6387

			} else {
				ifres6386 = False

			}

			if True == ifres6386 {
				tmp6380 := PrimHead(V2930)

				__e.TailApply(PrimFunc(symshen_4special_1case), tmp6380, V2930, V2931)
				return

			} else {
				tmp6384 := PrimIsPair(V2930)

				if True == tmp6384 {
					tmp6381 := MakeNative(func(__e *ControlFlow) {
						Y := __e.Get(1)
						_ = Y
						__e.TailApply(PrimFunc(symshen_4process_1applications), Y, V2931)
						return
					}, 1)

					tmp6382 := Call(__e, PrimFunc(symmap), tmp6381, V2930)

					__e.TailApply(PrimFunc(symshen_4process_1application), tmp6382, V2931)
					return

				} else {
					__e.Return(V2930)
					return
				}

			}

		}

	}, 2)

	tmp6393 := Call(__e, ns2_1set, symshen_4process_1applications, tmp6379)

	_ = tmp6393

	tmp6394 := MakeNative(func(__e *ControlFlow) {
		V2934 := __e.Get(1)
		_ = V2934
		tmp6404 := PrimEqual(symdefine, V2934)

		if True == tmp6404 {
			__e.Return(True)
			return
		} else {
			tmp6402 := PrimEqual(symdefun, V2934)

			if True == tmp6402 {
				__e.Return(True)
				return
			} else {
				tmp6400 := PrimEqual(symsynonyms, V2934)

				if True == tmp6400 {
					__e.Return(True)
					return
				} else {
					tmp6398 := Call(__e, PrimFunc(symshen_4special_2), V2934)

					if True == tmp6398 {
						__e.Return(True)
						return
					} else {
						tmp6396 := Call(__e, PrimFunc(symshen_4extraspecial_2), V2934)

						if True == tmp6396 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}

					}

				}

			}

		}

	}, 1)

	tmp6405 := Call(__e, ns2_1set, symshen_4non_1application_2, tmp6394)

	_ = tmp6405

	tmp6406 := MakeNative(func(__e *ControlFlow) {
		V2939 := __e.Get(1)
		_ = V2939
		V2940 := __e.Get(2)
		_ = V2940
		V2941 := __e.Get(3)
		_ = V2941
		tmp6648 := PrimEqual(symlambda, V2939)

		var ifres6626 Obj

		if True == tmp6648 {
			tmp6647 := PrimIsPair(V2940)

			var ifres6628 Obj

			if True == tmp6647 {
				tmp6645 := PrimHead(V2940)

				tmp6646 := PrimEqual(symlambda, tmp6645)

				var ifres6630 Obj

				if True == tmp6646 {
					tmp6643 := PrimTail(V2940)

					tmp6644 := PrimIsPair(tmp6643)

					var ifres6632 Obj

					if True == tmp6644 {
						tmp6640 := PrimTail(V2940)

						tmp6641 := PrimTail(tmp6640)

						tmp6642 := PrimIsPair(tmp6641)

						var ifres6634 Obj

						if True == tmp6642 {
							tmp6636 := PrimTail(V2940)

							tmp6637 := PrimTail(tmp6636)

							tmp6638 := PrimTail(tmp6637)

							tmp6639 := PrimEqual(Nil, tmp6638)

							var ifres6635 Obj

							if True == tmp6639 {
								ifres6635 = True

							} else {
								ifres6635 = False

							}

							ifres6634 = ifres6635

						} else {
							ifres6634 = False

						}

						var ifres6633 Obj

						if True == ifres6634 {
							ifres6633 = True

						} else {
							ifres6633 = False

						}

						ifres6632 = ifres6633

					} else {
						ifres6632 = False

					}

					var ifres6631 Obj

					if True == ifres6632 {
						ifres6631 = True

					} else {
						ifres6631 = False

					}

					ifres6630 = ifres6631

				} else {
					ifres6630 = False

				}

				var ifres6629 Obj

				if True == ifres6630 {
					ifres6629 = True

				} else {
					ifres6629 = False

				}

				ifres6628 = ifres6629

			} else {
				ifres6628 = False

			}

			var ifres6627 Obj

			if True == ifres6628 {
				ifres6627 = True

			} else {
				ifres6627 = False

			}

			ifres6626 = ifres6627

		} else {
			ifres6626 = False

		}

		if True == ifres6626 {
			tmp6407 := PrimTail(V2940)

			tmp6408 := PrimHead(tmp6407)

			tmp6409 := PrimTail(V2940)

			tmp6410 := PrimTail(tmp6409)

			tmp6411 := PrimHead(tmp6410)

			tmp6412 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6411, V2941)

			tmp6413 := PrimCons(tmp6412, Nil)

			tmp6414 := PrimCons(tmp6408, tmp6413)

			__e.Return(PrimCons(symlambda, tmp6414))
			return

		} else {
			tmp6624 := PrimEqual(symlet, V2939)

			var ifres6595 Obj

			if True == tmp6624 {
				tmp6623 := PrimIsPair(V2940)

				var ifres6597 Obj

				if True == tmp6623 {
					tmp6621 := PrimHead(V2940)

					tmp6622 := PrimEqual(symlet, tmp6621)

					var ifres6599 Obj

					if True == tmp6622 {
						tmp6619 := PrimTail(V2940)

						tmp6620 := PrimIsPair(tmp6619)

						var ifres6601 Obj

						if True == tmp6620 {
							tmp6616 := PrimTail(V2940)

							tmp6617 := PrimTail(tmp6616)

							tmp6618 := PrimIsPair(tmp6617)

							var ifres6603 Obj

							if True == tmp6618 {
								tmp6612 := PrimTail(V2940)

								tmp6613 := PrimTail(tmp6612)

								tmp6614 := PrimTail(tmp6613)

								tmp6615 := PrimIsPair(tmp6614)

								var ifres6605 Obj

								if True == tmp6615 {
									tmp6607 := PrimTail(V2940)

									tmp6608 := PrimTail(tmp6607)

									tmp6609 := PrimTail(tmp6608)

									tmp6610 := PrimTail(tmp6609)

									tmp6611 := PrimEqual(Nil, tmp6610)

									var ifres6606 Obj

									if True == tmp6611 {
										ifres6606 = True

									} else {
										ifres6606 = False

									}

									ifres6605 = ifres6606

								} else {
									ifres6605 = False

								}

								var ifres6604 Obj

								if True == ifres6605 {
									ifres6604 = True

								} else {
									ifres6604 = False

								}

								ifres6603 = ifres6604

							} else {
								ifres6603 = False

							}

							var ifres6602 Obj

							if True == ifres6603 {
								ifres6602 = True

							} else {
								ifres6602 = False

							}

							ifres6601 = ifres6602

						} else {
							ifres6601 = False

						}

						var ifres6600 Obj

						if True == ifres6601 {
							ifres6600 = True

						} else {
							ifres6600 = False

						}

						ifres6599 = ifres6600

					} else {
						ifres6599 = False

					}

					var ifres6598 Obj

					if True == ifres6599 {
						ifres6598 = True

					} else {
						ifres6598 = False

					}

					ifres6597 = ifres6598

				} else {
					ifres6597 = False

				}

				var ifres6596 Obj

				if True == ifres6597 {
					ifres6596 = True

				} else {
					ifres6596 = False

				}

				ifres6595 = ifres6596

			} else {
				ifres6595 = False

			}

			if True == ifres6595 {
				tmp6415 := PrimTail(V2940)

				tmp6416 := PrimHead(tmp6415)

				tmp6417 := PrimTail(V2940)

				tmp6418 := PrimTail(tmp6417)

				tmp6419 := PrimHead(tmp6418)

				tmp6420 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6419, V2941)

				tmp6421 := PrimTail(V2940)

				tmp6422 := PrimTail(tmp6421)

				tmp6423 := PrimTail(tmp6422)

				tmp6424 := PrimHead(tmp6423)

				tmp6425 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6424, V2941)

				tmp6426 := PrimCons(tmp6425, Nil)

				tmp6427 := PrimCons(tmp6420, tmp6426)

				tmp6428 := PrimCons(tmp6416, tmp6427)

				__e.Return(PrimCons(symlet, tmp6428))
				return

			} else {
				tmp6593 := PrimEqual(symdefun, V2939)

				var ifres6564 Obj

				if True == tmp6593 {
					tmp6592 := PrimIsPair(V2940)

					var ifres6566 Obj

					if True == tmp6592 {
						tmp6590 := PrimHead(V2940)

						tmp6591 := PrimEqual(symdefun, tmp6590)

						var ifres6568 Obj

						if True == tmp6591 {
							tmp6588 := PrimTail(V2940)

							tmp6589 := PrimIsPair(tmp6588)

							var ifres6570 Obj

							if True == tmp6589 {
								tmp6585 := PrimTail(V2940)

								tmp6586 := PrimTail(tmp6585)

								tmp6587 := PrimIsPair(tmp6586)

								var ifres6572 Obj

								if True == tmp6587 {
									tmp6581 := PrimTail(V2940)

									tmp6582 := PrimTail(tmp6581)

									tmp6583 := PrimTail(tmp6582)

									tmp6584 := PrimIsPair(tmp6583)

									var ifres6574 Obj

									if True == tmp6584 {
										tmp6576 := PrimTail(V2940)

										tmp6577 := PrimTail(tmp6576)

										tmp6578 := PrimTail(tmp6577)

										tmp6579 := PrimTail(tmp6578)

										tmp6580 := PrimEqual(Nil, tmp6579)

										var ifres6575 Obj

										if True == tmp6580 {
											ifres6575 = True

										} else {
											ifres6575 = False

										}

										ifres6574 = ifres6575

									} else {
										ifres6574 = False

									}

									var ifres6573 Obj

									if True == ifres6574 {
										ifres6573 = True

									} else {
										ifres6573 = False

									}

									ifres6572 = ifres6573

								} else {
									ifres6572 = False

								}

								var ifres6571 Obj

								if True == ifres6572 {
									ifres6571 = True

								} else {
									ifres6571 = False

								}

								ifres6570 = ifres6571

							} else {
								ifres6570 = False

							}

							var ifres6569 Obj

							if True == ifres6570 {
								ifres6569 = True

							} else {
								ifres6569 = False

							}

							ifres6568 = ifres6569

						} else {
							ifres6568 = False

						}

						var ifres6567 Obj

						if True == ifres6568 {
							ifres6567 = True

						} else {
							ifres6567 = False

						}

						ifres6566 = ifres6567

					} else {
						ifres6566 = False

					}

					var ifres6565 Obj

					if True == ifres6566 {
						ifres6565 = True

					} else {
						ifres6565 = False

					}

					ifres6564 = ifres6565

				} else {
					ifres6564 = False

				}

				if True == ifres6564 {
					__e.Return(V2940)
					return
				} else {
					tmp6562 := PrimEqual(symdefine, V2939)

					var ifres6540 Obj

					if True == tmp6562 {
						tmp6561 := PrimIsPair(V2940)

						var ifres6542 Obj

						if True == tmp6561 {
							tmp6559 := PrimHead(V2940)

							tmp6560 := PrimEqual(symdefine, tmp6559)

							var ifres6544 Obj

							if True == tmp6560 {
								tmp6557 := PrimTail(V2940)

								tmp6558 := PrimIsPair(tmp6557)

								var ifres6546 Obj

								if True == tmp6558 {
									tmp6554 := PrimTail(V2940)

									tmp6555 := PrimTail(tmp6554)

									tmp6556 := PrimIsPair(tmp6555)

									var ifres6548 Obj

									if True == tmp6556 {
										tmp6550 := PrimTail(V2940)

										tmp6551 := PrimTail(tmp6550)

										tmp6552 := PrimHead(tmp6551)

										tmp6553 := PrimEqual(sym_i, tmp6552)

										var ifres6549 Obj

										if True == tmp6553 {
											ifres6549 = True

										} else {
											ifres6549 = False

										}

										ifres6548 = ifres6549

									} else {
										ifres6548 = False

									}

									var ifres6547 Obj

									if True == ifres6548 {
										ifres6547 = True

									} else {
										ifres6547 = False

									}

									ifres6546 = ifres6547

								} else {
									ifres6546 = False

								}

								var ifres6545 Obj

								if True == ifres6546 {
									ifres6545 = True

								} else {
									ifres6545 = False

								}

								ifres6544 = ifres6545

							} else {
								ifres6544 = False

							}

							var ifres6543 Obj

							if True == ifres6544 {
								ifres6543 = True

							} else {
								ifres6543 = False

							}

							ifres6542 = ifres6543

						} else {
							ifres6542 = False

						}

						var ifres6541 Obj

						if True == ifres6542 {
							ifres6541 = True

						} else {
							ifres6541 = False

						}

						ifres6540 = ifres6541

					} else {
						ifres6540 = False

					}

					if True == ifres6540 {
						tmp6429 := PrimTail(V2940)

						tmp6430 := PrimHead(tmp6429)

						tmp6431 := PrimTail(V2940)

						tmp6432 := PrimHead(tmp6431)

						tmp6433 := PrimTail(V2940)

						tmp6434 := PrimTail(tmp6433)

						tmp6435 := PrimTail(tmp6434)

						tmp6436 := Call(__e, PrimFunc(symshen_4process_1after_1type), tmp6432, tmp6435, V2941)

						tmp6437 := PrimCons(sym_i, tmp6436)

						tmp6438 := PrimCons(tmp6430, tmp6437)

						__e.Return(PrimCons(symdefine, tmp6438))
						return

					} else {
						tmp6538 := PrimEqual(symdefine, V2939)

						var ifres6527 Obj

						if True == tmp6538 {
							tmp6537 := PrimIsPair(V2940)

							var ifres6529 Obj

							if True == tmp6537 {
								tmp6535 := PrimHead(V2940)

								tmp6536 := PrimEqual(symdefine, tmp6535)

								var ifres6531 Obj

								if True == tmp6536 {
									tmp6533 := PrimTail(V2940)

									tmp6534 := PrimIsPair(tmp6533)

									var ifres6532 Obj

									if True == tmp6534 {
										ifres6532 = True

									} else {
										ifres6532 = False

									}

									ifres6531 = ifres6532

								} else {
									ifres6531 = False

								}

								var ifres6530 Obj

								if True == ifres6531 {
									ifres6530 = True

								} else {
									ifres6530 = False

								}

								ifres6529 = ifres6530

							} else {
								ifres6529 = False

							}

							var ifres6528 Obj

							if True == ifres6529 {
								ifres6528 = True

							} else {
								ifres6528 = False

							}

							ifres6527 = ifres6528

						} else {
							ifres6527 = False

						}

						if True == ifres6527 {
							tmp6439 := PrimTail(V2940)

							tmp6440 := PrimHead(tmp6439)

							tmp6441 := MakeNative(func(__e *ControlFlow) {
								Y := __e.Get(1)
								_ = Y
								__e.TailApply(PrimFunc(symshen_4process_1applications), Y, V2941)
								return
							}, 1)

							tmp6442 := PrimTail(V2940)

							tmp6443 := PrimTail(tmp6442)

							tmp6444 := Call(__e, PrimFunc(symmap), tmp6441, tmp6443)

							tmp6445 := PrimCons(tmp6440, tmp6444)

							__e.Return(PrimCons(symdefine, tmp6445))
							return

						} else {
							tmp6525 := PrimEqual(symsynonyms, V2939)

							if True == tmp6525 {
								__e.Return(PrimCons(symsynonyms, V2940))
								return
							} else {
								tmp6523 := PrimEqual(symtype, V2939)

								var ifres6501 Obj

								if True == tmp6523 {
									tmp6522 := PrimIsPair(V2940)

									var ifres6503 Obj

									if True == tmp6522 {
										tmp6520 := PrimHead(V2940)

										tmp6521 := PrimEqual(symtype, tmp6520)

										var ifres6505 Obj

										if True == tmp6521 {
											tmp6518 := PrimTail(V2940)

											tmp6519 := PrimIsPair(tmp6518)

											var ifres6507 Obj

											if True == tmp6519 {
												tmp6515 := PrimTail(V2940)

												tmp6516 := PrimTail(tmp6515)

												tmp6517 := PrimIsPair(tmp6516)

												var ifres6509 Obj

												if True == tmp6517 {
													tmp6511 := PrimTail(V2940)

													tmp6512 := PrimTail(tmp6511)

													tmp6513 := PrimTail(tmp6512)

													tmp6514 := PrimEqual(Nil, tmp6513)

													var ifres6510 Obj

													if True == tmp6514 {
														ifres6510 = True

													} else {
														ifres6510 = False

													}

													ifres6509 = ifres6510

												} else {
													ifres6509 = False

												}

												var ifres6508 Obj

												if True == ifres6509 {
													ifres6508 = True

												} else {
													ifres6508 = False

												}

												ifres6507 = ifres6508

											} else {
												ifres6507 = False

											}

											var ifres6506 Obj

											if True == ifres6507 {
												ifres6506 = True

											} else {
												ifres6506 = False

											}

											ifres6505 = ifres6506

										} else {
											ifres6505 = False

										}

										var ifres6504 Obj

										if True == ifres6505 {
											ifres6504 = True

										} else {
											ifres6504 = False

										}

										ifres6503 = ifres6504

									} else {
										ifres6503 = False

									}

									var ifres6502 Obj

									if True == ifres6503 {
										ifres6502 = True

									} else {
										ifres6502 = False

									}

									ifres6501 = ifres6502

								} else {
									ifres6501 = False

								}

								if True == ifres6501 {
									tmp6446 := PrimTail(V2940)

									tmp6447 := PrimHead(tmp6446)

									tmp6448 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6447, V2941)

									tmp6449 := PrimTail(V2940)

									tmp6450 := PrimTail(tmp6449)

									tmp6451 := PrimCons(tmp6448, tmp6450)

									__e.Return(PrimCons(symtype, tmp6451))
									return

								} else {
									tmp6499 := PrimEqual(syminput_7, V2939)

									var ifres6477 Obj

									if True == tmp6499 {
										tmp6498 := PrimIsPair(V2940)

										var ifres6479 Obj

										if True == tmp6498 {
											tmp6496 := PrimHead(V2940)

											tmp6497 := PrimEqual(syminput_7, tmp6496)

											var ifres6481 Obj

											if True == tmp6497 {
												tmp6494 := PrimTail(V2940)

												tmp6495 := PrimIsPair(tmp6494)

												var ifres6483 Obj

												if True == tmp6495 {
													tmp6491 := PrimTail(V2940)

													tmp6492 := PrimTail(tmp6491)

													tmp6493 := PrimIsPair(tmp6492)

													var ifres6485 Obj

													if True == tmp6493 {
														tmp6487 := PrimTail(V2940)

														tmp6488 := PrimTail(tmp6487)

														tmp6489 := PrimTail(tmp6488)

														tmp6490 := PrimEqual(Nil, tmp6489)

														var ifres6486 Obj

														if True == tmp6490 {
															ifres6486 = True

														} else {
															ifres6486 = False

														}

														ifres6485 = ifres6486

													} else {
														ifres6485 = False

													}

													var ifres6484 Obj

													if True == ifres6485 {
														ifres6484 = True

													} else {
														ifres6484 = False

													}

													ifres6483 = ifres6484

												} else {
													ifres6483 = False

												}

												var ifres6482 Obj

												if True == ifres6483 {
													ifres6482 = True

												} else {
													ifres6482 = False

												}

												ifres6481 = ifres6482

											} else {
												ifres6481 = False

											}

											var ifres6480 Obj

											if True == ifres6481 {
												ifres6480 = True

											} else {
												ifres6480 = False

											}

											ifres6479 = ifres6480

										} else {
											ifres6479 = False

										}

										var ifres6478 Obj

										if True == ifres6479 {
											ifres6478 = True

										} else {
											ifres6478 = False

										}

										ifres6477 = ifres6478

									} else {
										ifres6477 = False

									}

									if True == ifres6477 {
										tmp6452 := PrimTail(V2940)

										tmp6453 := PrimHead(tmp6452)

										tmp6454 := PrimTail(V2940)

										tmp6455 := PrimTail(tmp6454)

										tmp6456 := PrimHead(tmp6455)

										tmp6457 := Call(__e, PrimFunc(symshen_4process_1applications), tmp6456, V2941)

										tmp6458 := PrimCons(tmp6457, Nil)

										tmp6459 := PrimCons(tmp6453, tmp6458)

										__e.Return(PrimCons(syminput_7, tmp6459))
										return

									} else {
										tmp6475 := PrimIsPair(V2940)

										var ifres6471 Obj

										if True == tmp6475 {
											tmp6473 := PrimHead(V2940)

											tmp6474 := Call(__e, PrimFunc(symshen_4special_2), tmp6473)

											var ifres6472 Obj

											if True == tmp6474 {
												ifres6472 = True

											} else {
												ifres6472 = False

											}

											ifres6471 = ifres6472

										} else {
											ifres6471 = False

										}

										if True == ifres6471 {
											tmp6460 := PrimHead(V2940)

											tmp6461 := MakeNative(func(__e *ControlFlow) {
												Y := __e.Get(1)
												_ = Y
												__e.TailApply(PrimFunc(symshen_4process_1applications), Y, V2941)
												return
											}, 1)

											tmp6462 := PrimTail(V2940)

											tmp6463 := Call(__e, PrimFunc(symmap), tmp6461, tmp6462)

											__e.Return(PrimCons(tmp6460, tmp6463))
											return

										} else {
											tmp6469 := PrimIsPair(V2940)

											var ifres6465 Obj

											if True == tmp6469 {
												tmp6467 := PrimHead(V2940)

												tmp6468 := Call(__e, PrimFunc(symshen_4extraspecial_2), tmp6467)

												var ifres6466 Obj

												if True == tmp6468 {
													ifres6466 = True

												} else {
													ifres6466 = False

												}

												ifres6465 = ifres6466

											} else {
												ifres6465 = False

											}

											if True == ifres6465 {
												__e.Return(V2940)
												return
											} else {
												__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4special_1case)
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

	}, 3)

	tmp6649 := Call(__e, ns2_1set, symshen_4special_1case, tmp6406)

	_ = tmp6649

	tmp6650 := MakeNative(func(__e *ControlFlow) {
		V2944 := __e.Get(1)
		_ = V2944
		V2945 := __e.Get(2)
		_ = V2945
		V2946 := __e.Get(3)
		_ = V2946
		tmp6666 := PrimIsPair(V2945)

		var ifres6662 Obj

		if True == tmp6666 {
			tmp6664 := PrimHead(V2945)

			tmp6665 := PrimEqual(sym_j, tmp6664)

			var ifres6663 Obj

			if True == tmp6665 {
				ifres6663 = True

			} else {
				ifres6663 = False

			}

			ifres6662 = ifres6663

		} else {
			ifres6662 = False

		}

		if True == ifres6662 {
			tmp6651 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(PrimFunc(symshen_4process_1applications), Y, V2946)
				return
			}, 1)

			tmp6652 := PrimTail(V2945)

			tmp6653 := Call(__e, PrimFunc(symmap), tmp6651, tmp6652)

			__e.Return(PrimCons(sym_j, tmp6653))
			return

		} else {
			tmp6660 := PrimIsPair(V2945)

			if True == tmp6660 {
				tmp6654 := PrimHead(V2945)

				tmp6655 := PrimTail(V2945)

				tmp6656 := Call(__e, PrimFunc(symshen_4process_1after_1type), V2944, tmp6655, V2946)

				__e.Return(PrimCons(tmp6654, tmp6656))
				return

			} else {
				tmp6657 := Call(__e, PrimFunc(symshen_4app), V2944, MakeString("\n"), symshen_4a)

				tmp6658 := PrimStringConcat(MakeString("missing } in "), tmp6657)

				__e.Return(PrimSimpleError(tmp6658))
				return

			}

		}

	}, 3)

	tmp6667 := Call(__e, ns2_1set, symshen_4process_1after_1type, tmp6650)

	_ = tmp6667

	tmp6668 := MakeNative(func(__e *ControlFlow) {
		V2947 := __e.Get(1)
		_ = V2947
		V2948 := __e.Get(2)
		_ = V2948
		tmp6706 := PrimIsPair(V2947)

		if True == tmp6706 {
			tmp6669 := MakeNative(func(__e *ControlFlow) {
				ArityF := __e.Get(1)
				_ = ArityF
				tmp6670 := MakeNative(func(__e *ControlFlow) {
					N := __e.Get(1)
					_ = N
					tmp6700 := Call(__e, PrimFunc(symelement_2), V2947, V2948)

					if True == tmp6700 {
						__e.Return(V2947)
						return
					} else {
						tmp6697 := PrimHead(V2947)

						tmp6698 := Call(__e, PrimFunc(symshen_4shen_1call_2), tmp6697)

						if True == tmp6698 {
							__e.Return(V2947)
							return
						} else {
							tmp6695 := Call(__e, PrimFunc(symshen_4fn_1call_2), V2947)

							if True == tmp6695 {
								__e.TailApply(PrimFunc(symshen_4fn_1call), V2947)
								return
							} else {
								tmp6693 := Call(__e, PrimFunc(symshen_4zero_1place_2), V2947)

								if True == tmp6693 {
									__e.Return(V2947)
									return
								} else {
									tmp6690 := PrimHead(V2947)

									tmp6691 := Call(__e, PrimFunc(symshen_4undefined_1f_2), tmp6690, ArityF)

									if True == tmp6691 {
										tmp6671 := PrimHead(V2947)

										tmp6672 := PrimCons(tmp6671, Nil)

										tmp6673 := PrimCons(symfn, tmp6672)

										tmp6674 := PrimTail(V2947)

										tmp6675 := PrimCons(tmp6673, tmp6674)

										__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6675)
										return

									} else {
										tmp6687 := PrimHead(V2947)

										tmp6688 := PrimIsVariable(tmp6687)

										if True == tmp6688 {
											__e.TailApply(PrimFunc(symshen_4simple_1curry), V2947)
											return
										} else {
											tmp6684 := PrimHead(V2947)

											tmp6685 := Call(__e, PrimFunc(symshen_4application_2), tmp6684)

											if True == tmp6685 {
												__e.TailApply(PrimFunc(symshen_4simple_1curry), V2947)
												return
											} else {
												tmp6681 := PrimHead(V2947)

												tmp6682 := Call(__e, PrimFunc(symshen_4partial_1application_d_2), tmp6681, ArityF, N)

												if True == tmp6682 {
													tmp6676 := PrimNumberSubtract(ArityF, N)

													__e.TailApply(PrimFunc(symshen_4lambda_1function), V2947, tmp6676)
													return

												} else {
													tmp6678 := PrimHead(V2947)

													tmp6679 := Call(__e, PrimFunc(symshen_4overapplication_2), tmp6678, ArityF, N)

													if True == tmp6679 {
														__e.TailApply(PrimFunc(symshen_4simple_1curry), V2947)
														return
													} else {
														__e.Return(V2947)
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

				}, 1)

				tmp6701 := PrimTail(V2947)

				tmp6702 := Call(__e, PrimFunc(symlength), tmp6701)

				__e.TailApply(tmp6670, tmp6702)
				return

			}, 1)

			tmp6703 := PrimHead(V2947)

			tmp6704 := Call(__e, PrimFunc(symarity), tmp6703)

			__e.TailApply(tmp6669, tmp6704)
			return

		} else {
			__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4process_1application)
			return
		}

	}, 2)

	tmp6707 := Call(__e, ns2_1set, symshen_4process_1application, tmp6668)

	_ = tmp6707

	tmp6708 := MakeNative(func(__e *ControlFlow) {
		V2951 := __e.Get(1)
		_ = V2951
		tmp6714 := PrimIsPair(V2951)

		var ifres6710 Obj

		if True == tmp6714 {
			tmp6712 := PrimTail(V2951)

			tmp6713 := PrimEqual(Nil, tmp6712)

			var ifres6711 Obj

			if True == tmp6713 {
				ifres6711 = True

			} else {
				ifres6711 = False

			}

			ifres6710 = ifres6711

		} else {
			ifres6710 = False

		}

		if True == ifres6710 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp6715 := Call(__e, ns2_1set, symshen_4zero_1place_2, tmp6708)

	_ = tmp6715

	tmp6716 := MakeNative(func(__e *ControlFlow) {
		V2952 := __e.Get(1)
		_ = V2952
		tmp6721 := PrimIsSymbol(V2952)

		if True == tmp6721 {
			tmp6718 := PrimStr(V2952)

			tmp6719 := Call(__e, PrimFunc(symshen_4internal_1to_1shen_2), tmp6718)

			if True == tmp6719 {
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

	tmp6722 := Call(__e, ns2_1set, symshen_4shen_1call_2, tmp6716)

	_ = tmp6722

	tmp6723 := MakeNative(func(__e *ControlFlow) {
		V2957 := __e.Get(1)
		_ = V2957
		tmp6777 := Call(__e, PrimFunc(symshen_4_7string_2), V2957)

		var ifres6725 Obj

		if True == tmp6777 {
			tmp6775 := Call(__e, PrimFunc(symhdstr), V2957)

			tmp6776 := PrimEqual(MakeString("s"), tmp6775)

			var ifres6727 Obj

			if True == tmp6776 {
				tmp6773 := PrimTailString(V2957)

				tmp6774 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6773)

				var ifres6729 Obj

				if True == tmp6774 {
					tmp6770 := PrimTailString(V2957)

					tmp6771 := Call(__e, PrimFunc(symhdstr), tmp6770)

					tmp6772 := PrimEqual(MakeString("h"), tmp6771)

					var ifres6731 Obj

					if True == tmp6772 {
						tmp6767 := PrimTailString(V2957)

						tmp6768 := PrimTailString(tmp6767)

						tmp6769 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6768)

						var ifres6733 Obj

						if True == tmp6769 {
							tmp6763 := PrimTailString(V2957)

							tmp6764 := PrimTailString(tmp6763)

							tmp6765 := Call(__e, PrimFunc(symhdstr), tmp6764)

							tmp6766 := PrimEqual(MakeString("e"), tmp6765)

							var ifres6735 Obj

							if True == tmp6766 {
								tmp6759 := PrimTailString(V2957)

								tmp6760 := PrimTailString(tmp6759)

								tmp6761 := PrimTailString(tmp6760)

								tmp6762 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6761)

								var ifres6737 Obj

								if True == tmp6762 {
									tmp6754 := PrimTailString(V2957)

									tmp6755 := PrimTailString(tmp6754)

									tmp6756 := PrimTailString(tmp6755)

									tmp6757 := Call(__e, PrimFunc(symhdstr), tmp6756)

									tmp6758 := PrimEqual(MakeString("n"), tmp6757)

									var ifres6739 Obj

									if True == tmp6758 {
										tmp6749 := PrimTailString(V2957)

										tmp6750 := PrimTailString(tmp6749)

										tmp6751 := PrimTailString(tmp6750)

										tmp6752 := PrimTailString(tmp6751)

										tmp6753 := Call(__e, PrimFunc(symshen_4_7string_2), tmp6752)

										var ifres6741 Obj

										if True == tmp6753 {
											tmp6743 := PrimTailString(V2957)

											tmp6744 := PrimTailString(tmp6743)

											tmp6745 := PrimTailString(tmp6744)

											tmp6746 := PrimTailString(tmp6745)

											tmp6747 := Call(__e, PrimFunc(symhdstr), tmp6746)

											tmp6748 := PrimEqual(MakeString("."), tmp6747)

											var ifres6742 Obj

											if True == tmp6748 {
												ifres6742 = True

											} else {
												ifres6742 = False

											}

											ifres6741 = ifres6742

										} else {
											ifres6741 = False

										}

										var ifres6740 Obj

										if True == ifres6741 {
											ifres6740 = True

										} else {
											ifres6740 = False

										}

										ifres6739 = ifres6740

									} else {
										ifres6739 = False

									}

									var ifres6738 Obj

									if True == ifres6739 {
										ifres6738 = True

									} else {
										ifres6738 = False

									}

									ifres6737 = ifres6738

								} else {
									ifres6737 = False

								}

								var ifres6736 Obj

								if True == ifres6737 {
									ifres6736 = True

								} else {
									ifres6736 = False

								}

								ifres6735 = ifres6736

							} else {
								ifres6735 = False

							}

							var ifres6734 Obj

							if True == ifres6735 {
								ifres6734 = True

							} else {
								ifres6734 = False

							}

							ifres6733 = ifres6734

						} else {
							ifres6733 = False

						}

						var ifres6732 Obj

						if True == ifres6733 {
							ifres6732 = True

						} else {
							ifres6732 = False

						}

						ifres6731 = ifres6732

					} else {
						ifres6731 = False

					}

					var ifres6730 Obj

					if True == ifres6731 {
						ifres6730 = True

					} else {
						ifres6730 = False

					}

					ifres6729 = ifres6730

				} else {
					ifres6729 = False

				}

				var ifres6728 Obj

				if True == ifres6729 {
					ifres6728 = True

				} else {
					ifres6728 = False

				}

				ifres6727 = ifres6728

			} else {
				ifres6727 = False

			}

			var ifres6726 Obj

			if True == ifres6727 {
				ifres6726 = True

			} else {
				ifres6726 = False

			}

			ifres6725 = ifres6726

		} else {
			ifres6725 = False

		}

		if True == ifres6725 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp6778 := Call(__e, ns2_1set, symshen_4internal_1to_1shen_2, tmp6723)

	_ = tmp6778

	tmp6779 := MakeNative(func(__e *ControlFlow) {
		V2958 := __e.Get(1)
		_ = V2958
		__e.Return(PrimIsPair(V2958))
		return
	}, 1)

	tmp6780 := Call(__e, ns2_1set, symshen_4application_2, tmp6779)

	_ = tmp6780

	tmp6781 := MakeNative(func(__e *ControlFlow) {
		V2963 := __e.Get(1)
		_ = V2963
		V2964 := __e.Get(2)
		_ = V2964
		tmp6789 := PrimEqual(MakeNumber(-1), V2964)

		if True == tmp6789 {
			tmp6787 := Call(__e, PrimFunc(symshen_4lowercase_1symbol_2), V2963)

			if True == tmp6787 {
				tmp6783 := Call(__e, PrimFunc(symexternal), symshen)

				tmp6784 := Call(__e, PrimFunc(symelement_2), V2963, tmp6783)

				tmp6785 := PrimNot(tmp6784)

				if True == tmp6785 {
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

		} else {
			__e.Return(False)
			return
		}

	}, 2)

	tmp6790 := Call(__e, ns2_1set, symshen_4undefined_1f_2, tmp6781)

	_ = tmp6790

	tmp6791 := MakeNative(func(__e *ControlFlow) {
		V2965 := __e.Get(1)
		_ = V2965
		tmp6796 := PrimIsSymbol(V2965)

		if True == tmp6796 {
			tmp6793 := PrimIsVariable(V2965)

			tmp6794 := PrimNot(tmp6793)

			if True == tmp6794 {
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

	tmp6797 := Call(__e, ns2_1set, symshen_4lowercase_1symbol_2, tmp6791)

	_ = tmp6797

	tmp6798 := MakeNative(func(__e *ControlFlow) {
		V2966 := __e.Get(1)
		_ = V2966
		tmp6828 := PrimIsPair(V2966)

		var ifres6819 Obj

		if True == tmp6828 {
			tmp6826 := PrimTail(V2966)

			tmp6827 := PrimIsPair(tmp6826)

			var ifres6821 Obj

			if True == tmp6827 {
				tmp6823 := PrimTail(V2966)

				tmp6824 := PrimTail(tmp6823)

				tmp6825 := PrimEqual(Nil, tmp6824)

				var ifres6822 Obj

				if True == tmp6825 {
					ifres6822 = True

				} else {
					ifres6822 = False

				}

				ifres6821 = ifres6822

			} else {
				ifres6821 = False

			}

			var ifres6820 Obj

			if True == ifres6821 {
				ifres6820 = True

			} else {
				ifres6820 = False

			}

			ifres6819 = ifres6820

		} else {
			ifres6819 = False

		}

		if True == ifres6819 {
			__e.Return(V2966)
			return
		} else {
			tmp6817 := PrimIsPair(V2966)

			var ifres6808 Obj

			if True == tmp6817 {
				tmp6815 := PrimTail(V2966)

				tmp6816 := PrimIsPair(tmp6815)

				var ifres6810 Obj

				if True == tmp6816 {
					tmp6812 := PrimTail(V2966)

					tmp6813 := PrimTail(tmp6812)

					tmp6814 := PrimIsPair(tmp6813)

					var ifres6811 Obj

					if True == tmp6814 {
						ifres6811 = True

					} else {
						ifres6811 = False

					}

					ifres6810 = ifres6811

				} else {
					ifres6810 = False

				}

				var ifres6809 Obj

				if True == ifres6810 {
					ifres6809 = True

				} else {
					ifres6809 = False

				}

				ifres6808 = ifres6809

			} else {
				ifres6808 = False

			}

			if True == ifres6808 {
				tmp6799 := PrimHead(V2966)

				tmp6800 := PrimTail(V2966)

				tmp6801 := PrimHead(tmp6800)

				tmp6802 := PrimCons(tmp6801, Nil)

				tmp6803 := PrimCons(tmp6799, tmp6802)

				tmp6804 := PrimTail(V2966)

				tmp6805 := PrimTail(tmp6804)

				tmp6806 := PrimCons(tmp6803, tmp6805)

				__e.TailApply(PrimFunc(symshen_4simple_1curry), tmp6806)
				return

			} else {
				__e.Return(V2966)
				return
			}

		}

	}, 1)

	tmp6829 := Call(__e, ns2_1set, symshen_4simple_1curry, tmp6798)

	_ = tmp6829

	tmp6830 := MakeNative(func(__e *ControlFlow) {
		V2967 := __e.Get(1)
		_ = V2967
		__e.TailApply(PrimFunc(symfn), V2967)
		return
	}, 1)

	tmp6831 := Call(__e, ns2_1set, symfunction, tmp6830)

	_ = tmp6831

	tmp6832 := MakeNative(func(__e *ControlFlow) {
		V2968 := __e.Get(1)
		_ = V2968
		tmp6833 := MakeNative(func(__e *ControlFlow) {
			LookUp := __e.Get(1)
			_ = LookUp
			tmp6837 := Call(__e, PrimFunc(symempty_2), LookUp)

			if True == tmp6837 {
				tmp6834 := Call(__e, PrimFunc(symshen_4app), V2968, MakeString(" is undefined\n"), symshen_4a)

				tmp6835 := PrimStringConcat(MakeString("fn: "), tmp6834)

				__e.Return(PrimSimpleError(tmp6835))
				return

			} else {
				__e.Return(PrimTail(LookUp))
				return
			}

		}, 1)

		tmp6838 := PrimValue(symshen_4_dlambdatable_d)

		tmp6839 := Call(__e, PrimFunc(symassoc), V2968, tmp6838)

		__e.TailApply(tmp6833, tmp6839)
		return

	}, 1)

	tmp6840 := Call(__e, ns2_1set, symfn, tmp6832)

	_ = tmp6840

	tmp6841 := MakeNative(func(__e *ControlFlow) {
		V2971 := __e.Get(1)
		_ = V2971
		tmp6871 := PrimIsPair(V2971)

		var ifres6858 Obj

		if True == tmp6871 {
			tmp6869 := PrimHead(V2971)

			tmp6870 := PrimEqual(symfn, tmp6869)

			var ifres6860 Obj

			if True == tmp6870 {
				tmp6867 := PrimTail(V2971)

				tmp6868 := PrimIsPair(tmp6867)

				var ifres6862 Obj

				if True == tmp6868 {
					tmp6864 := PrimTail(V2971)

					tmp6865 := PrimTail(tmp6864)

					tmp6866 := PrimEqual(Nil, tmp6865)

					var ifres6863 Obj

					if True == tmp6866 {
						ifres6863 = True

					} else {
						ifres6863 = False

					}

					ifres6862 = ifres6863

				} else {
					ifres6862 = False

				}

				var ifres6861 Obj

				if True == ifres6862 {
					ifres6861 = True

				} else {
					ifres6861 = False

				}

				ifres6860 = ifres6861

			} else {
				ifres6860 = False

			}

			var ifres6859 Obj

			if True == ifres6860 {
				ifres6859 = True

			} else {
				ifres6859 = False

			}

			ifres6858 = ifres6859

		} else {
			ifres6858 = False

		}

		if True == ifres6858 {
			__e.Return(True)
			return
		} else {
			tmp6856 := PrimIsPair(V2971)

			var ifres6843 Obj

			if True == tmp6856 {
				tmp6854 := PrimHead(V2971)

				tmp6855 := PrimEqual(symfunction, tmp6854)

				var ifres6845 Obj

				if True == tmp6855 {
					tmp6852 := PrimTail(V2971)

					tmp6853 := PrimIsPair(tmp6852)

					var ifres6847 Obj

					if True == tmp6853 {
						tmp6849 := PrimTail(V2971)

						tmp6850 := PrimTail(tmp6849)

						tmp6851 := PrimEqual(Nil, tmp6850)

						var ifres6848 Obj

						if True == tmp6851 {
							ifres6848 = True

						} else {
							ifres6848 = False

						}

						ifres6847 = ifres6848

					} else {
						ifres6847 = False

					}

					var ifres6846 Obj

					if True == ifres6847 {
						ifres6846 = True

					} else {
						ifres6846 = False

					}

					ifres6845 = ifres6846

				} else {
					ifres6845 = False

				}

				var ifres6844 Obj

				if True == ifres6845 {
					ifres6844 = True

				} else {
					ifres6844 = False

				}

				ifres6843 = ifres6844

			} else {
				ifres6843 = False

			}

			if True == ifres6843 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp6872 := Call(__e, ns2_1set, symshen_4fn_1call_2, tmp6841)

	_ = tmp6872

	tmp6873 := MakeNative(func(__e *ControlFlow) {
		V2972 := __e.Get(1)
		_ = V2972
		tmp6914 := PrimIsPair(V2972)

		var ifres6901 Obj

		if True == tmp6914 {
			tmp6912 := PrimHead(V2972)

			tmp6913 := PrimEqual(symfunction, tmp6912)

			var ifres6903 Obj

			if True == tmp6913 {
				tmp6910 := PrimTail(V2972)

				tmp6911 := PrimIsPair(tmp6910)

				var ifres6905 Obj

				if True == tmp6911 {
					tmp6907 := PrimTail(V2972)

					tmp6908 := PrimTail(tmp6907)

					tmp6909 := PrimEqual(Nil, tmp6908)

					var ifres6906 Obj

					if True == tmp6909 {
						ifres6906 = True

					} else {
						ifres6906 = False

					}

					ifres6905 = ifres6906

				} else {
					ifres6905 = False

				}

				var ifres6904 Obj

				if True == ifres6905 {
					ifres6904 = True

				} else {
					ifres6904 = False

				}

				ifres6903 = ifres6904

			} else {
				ifres6903 = False

			}

			var ifres6902 Obj

			if True == ifres6903 {
				ifres6902 = True

			} else {
				ifres6902 = False

			}

			ifres6901 = ifres6902

		} else {
			ifres6901 = False

		}

		if True == ifres6901 {
			tmp6874 := PrimTail(V2972)

			tmp6875 := PrimCons(symfn, tmp6874)

			__e.TailApply(PrimFunc(symshen_4fn_1call), tmp6875)
			return

		} else {
			tmp6899 := PrimIsPair(V2972)

			var ifres6886 Obj

			if True == tmp6899 {
				tmp6897 := PrimHead(V2972)

				tmp6898 := PrimEqual(symfn, tmp6897)

				var ifres6888 Obj

				if True == tmp6898 {
					tmp6895 := PrimTail(V2972)

					tmp6896 := PrimIsPair(tmp6895)

					var ifres6890 Obj

					if True == tmp6896 {
						tmp6892 := PrimTail(V2972)

						tmp6893 := PrimTail(tmp6892)

						tmp6894 := PrimEqual(Nil, tmp6893)

						var ifres6891 Obj

						if True == tmp6894 {
							ifres6891 = True

						} else {
							ifres6891 = False

						}

						ifres6890 = ifres6891

					} else {
						ifres6890 = False

					}

					var ifres6889 Obj

					if True == ifres6890 {
						ifres6889 = True

					} else {
						ifres6889 = False

					}

					ifres6888 = ifres6889

				} else {
					ifres6888 = False

				}

				var ifres6887 Obj

				if True == ifres6888 {
					ifres6887 = True

				} else {
					ifres6887 = False

				}

				ifres6886 = ifres6887

			} else {
				ifres6886 = False

			}

			if True == ifres6886 {
				tmp6876 := MakeNative(func(__e *ControlFlow) {
					ArityF := __e.Get(1)
					_ = ArityF
					tmp6881 := PrimEqual(ArityF, MakeNumber(-1))

					if True == tmp6881 {
						__e.Return(V2972)
						return
					} else {
						tmp6879 := PrimEqual(ArityF, MakeNumber(0))

						if True == tmp6879 {
							__e.Return(PrimSimpleError(MakeString("fn cannot be applied to a zero place function\n")))
							return
						} else {
							tmp6877 := PrimTail(V2972)

							__e.TailApply(PrimFunc(symshen_4lambda_1function), tmp6877, ArityF)
							return

						}

					}

				}, 1)

				tmp6882 := PrimTail(V2972)

				tmp6883 := PrimHead(tmp6882)

				tmp6884 := Call(__e, PrimFunc(symarity), tmp6883)

				__e.TailApply(tmp6876, tmp6884)
				return

			} else {
				__e.TailApply(PrimFunc(symshen_4f_1error), symshen_4fn_1call)
				return
			}

		}

	}, 1)

	tmp6915 := Call(__e, ns2_1set, symshen_4fn_1call, tmp6873)

	_ = tmp6915

	tmp6916 := MakeNative(func(__e *ControlFlow) {
		V2973 := __e.Get(1)
		_ = V2973
		V2974 := __e.Get(2)
		_ = V2974
		V2975 := __e.Get(3)
		_ = V2975
		tmp6917 := MakeNative(func(__e *ControlFlow) {
			Verdict := __e.Get(1)
			_ = Verdict
			tmp6918 := MakeNative(func(__e *ControlFlow) {
				Message := __e.Get(1)
				_ = Message
				__e.Return(Verdict)
				return
			}, 1)

			var ifres6924 Obj

			if True == Verdict {
				tmp6932 := Call(__e, PrimFunc(symshen_4loading_2))

				var ifres6926 Obj

				if True == tmp6932 {
					tmp6928 := PrimCons(sym_1, Nil)

					tmp6929 := PrimCons(sym_7, tmp6928)

					tmp6930 := Call(__e, PrimFunc(symelement_2), V2973, tmp6929)

					tmp6931 := PrimNot(tmp6930)

					var ifres6927 Obj

					if True == tmp6931 {
						ifres6927 = True

					} else {
						ifres6927 = False

					}

					ifres6926 = ifres6927

				} else {
					ifres6926 = False

				}

				var ifres6925 Obj

				if True == ifres6926 {
					ifres6925 = True

				} else {
					ifres6925 = False

				}

				ifres6924 = ifres6925

			} else {
				ifres6924 = False

			}

			var ifres6919 Obj

			if True == ifres6924 {
				tmp6920 := Call(__e, PrimFunc(symshen_4app), V2973, MakeString("\n"), symshen_4a)

				tmp6921 := PrimStringConcat(MakeString("partial application of "), tmp6920)

				tmp6922 := Call(__e, PrimFunc(symstoutput))

				tmp6923 := Call(__e, PrimFunc(sympr), tmp6921, tmp6922)

				ifres6919 = tmp6923

			} else {
				ifres6919 = symshen_4skip

			}

			__e.TailApply(tmp6918, ifres6919)
			return

		}, 1)

		tmp6933 := PrimGreatThan(V2974, V2975)

		__e.TailApply(tmp6917, tmp6933)
		return

	}, 3)

	tmp6934 := Call(__e, ns2_1set, symshen_4partial_1application_d_2, tmp6916)

	_ = tmp6934

	tmp6935 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(symshen_4_dloading_2_d))
		return
	}, 0)

	tmp6936 := Call(__e, ns2_1set, symshen_4loading_2, tmp6935)

	_ = tmp6936

	tmp6937 := MakeNative(func(__e *ControlFlow) {
		V2976 := __e.Get(1)
		_ = V2976
		V2977 := __e.Get(2)
		_ = V2977
		V2978 := __e.Get(3)
		_ = V2978
		tmp6938 := MakeNative(func(__e *ControlFlow) {
			Verdict := __e.Get(1)
			_ = Verdict
			tmp6939 := MakeNative(func(__e *ControlFlow) {
				Message := __e.Get(1)
				_ = Message
				__e.Return(Verdict)
				return
			}, 1)

			var ifres6950 Obj

			if True == Verdict {
				tmp6952 := Call(__e, PrimFunc(symshen_4loading_2))

				var ifres6951 Obj

				if True == tmp6952 {
					ifres6951 = True

				} else {
					ifres6951 = False

				}

				ifres6950 = ifres6951

			} else {
				ifres6950 = False

			}

			var ifres6940 Obj

			if True == ifres6950 {
				tmp6942 := PrimEqual(V2978, MakeNumber(1))

				var ifres6941 Obj

				if True == tmp6942 {
					ifres6941 = MakeString("")

				} else {
					ifres6941 = MakeString("s")

				}

				tmp6943 := Call(__e, PrimFunc(symshen_4app), ifres6941, MakeString("\n"), symshen_4a)

				tmp6944 := PrimStringConcat(MakeString(" argument"), tmp6943)

				tmp6945 := Call(__e, PrimFunc(symshen_4app), V2978, tmp6944, symshen_4a)

				tmp6946 := PrimStringConcat(MakeString(" might not like "), tmp6945)

				tmp6947 := Call(__e, PrimFunc(symshen_4app), V2976, tmp6946, symshen_4a)

				tmp6948 := Call(__e, PrimFunc(symstoutput))

				tmp6949 := Call(__e, PrimFunc(sympr), tmp6947, tmp6948)

				ifres6940 = tmp6949

			} else {
				ifres6940 = symshen_4skip

			}

			__e.TailApply(tmp6939, ifres6940)
			return

		}, 1)

		tmp6953 := PrimLessThan(V2977, V2978)

		__e.TailApply(tmp6938, tmp6953)
		return

	}, 3)

	__e.TailApply(ns2_1set, symshen_4overapplication_2, tmp6937)
	return

}, 0)
