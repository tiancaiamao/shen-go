package main

import . "github.com/tiancaiamao/shen-go/klambda"

var ReaderMain = MakeNative(func(__e *ControlFlow) {
	tmp3375 := MakeNative(func(__e *ControlFlow) {
		V2726 := __e.Get(1)
		_ = V2726
		tmp3376 := MakeNative(func(__e *ControlFlow) {
			Bytelist := __e.Get(1)
			_ = Bytelist
			tmp3377 := MakeNative(func(__e *ControlFlow) {
				S_1exprs := __e.Get(1)
				_ = S_1exprs
				tmp3378 := MakeNative(func(__e *ControlFlow) {
					Process := __e.Get(1)
					_ = Process
					__e.Return(Process)
					return
				}, 1)

				tmp3379 := Call(__e, PrimNS2Value(symshen_4process_1sexprs), S_1exprs)

				__e.TailApply(tmp3378, tmp3379)
				return

			}, 1)

			tmp3380 := MakeNative(func(__e *ControlFlow) {
				tmp3381 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					__e.TailApply(PrimNS2Value(symshen_4_5s_1exprs_6), X)
					return
				}, 1)

				__e.TailApply(PrimNS2Value(symcompile), tmp3381, Bytelist)
				return

			}, 0)

			tmp3382 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				tmp3383 := PrimNS3Value(symshen_4_dresidue_d)

				__e.TailApply(PrimNS2Value(symshen_4print_1residue), tmp3383)
				return

			}, 1)

			tmp3384 := Call(__e, PrimNS1Value(symtry_1catch), tmp3380, tmp3382)

			__e.TailApply(tmp3377, tmp3384)
			return

		}, 1)

		tmp3385 := PrimReadFileAsByteList(V2726)

		__e.TailApply(tmp3376, tmp3385)
		return

	}, 1)

	tmp3386 := Call(__e, PrimNS2Value(symdef), symread_1file, tmp3375)

	_ = tmp3386

	tmp3387 := MakeNative(func(__e *ControlFlow) {
		V2727 := __e.Get(1)
		_ = V2727
		tmp3388 := MakeNative(func(__e *ControlFlow) {
			Err := __e.Get(1)
			_ = Err
			__e.TailApply(PrimNS2Value(symshen_4nchars), MakeNumber(50), V2727)
			return
		}, 1)

		tmp3389 := Call(__e, PrimNS2Value(symstoutput))

		tmp3390 := Call(__e, PrimNS2Value(sympr), MakeString("syntax error here:\n"), tmp3389)

		__e.TailApply(tmp3388, tmp3390)
		return

	}, 1)

	tmp3391 := Call(__e, PrimNS2Value(symdef), symshen_4print_1residue, tmp3387)

	_ = tmp3391

	tmp3392 := MakeNative(func(__e *ControlFlow) {
		V2732 := __e.Get(1)
		_ = V2732
		V2733 := __e.Get(2)
		_ = V2733
		tmp3408 := PrimEqual(MakeNumber(0), V2732)

		if True == tmp3408 {
			tmp3406 := Call(__e, PrimNS2Value(symstoutput))

			tmp3407 := Call(__e, PrimNS2Value(sympr), MakeString(" ..."), tmp3406)

			_ = tmp3407

			__e.TailApply(PrimNS2Value(symabort))
			return

		} else {
			tmp3405 := PrimEqual(Nil, V2733)

			if True == tmp3405 {
				tmp3403 := Call(__e, PrimNS2Value(symstoutput))

				tmp3404 := Call(__e, PrimNS2Value(sympr), MakeString(" ..."), tmp3403)

				_ = tmp3404

				__e.TailApply(PrimNS2Value(symabort))
				return

			} else {
				tmp3402 := PrimIsPair(V2733)

				if True == tmp3402 {
					tmp3396 := PrimHead(V2733)

					tmp3397 := PrimNumberToString(tmp3396)

					tmp3398 := Call(__e, PrimNS2Value(symstoutput))

					tmp3399 := Call(__e, PrimNS2Value(sympr), tmp3397, tmp3398)

					_ = tmp3399

					tmp3400 := PrimNumberSubtract(V2732, MakeNumber(1))

					tmp3401 := PrimTail(V2733)

					__e.TailApply(PrimNS2Value(symshen_4nchars), tmp3400, tmp3401)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4nchars)
					return
				}

			}

		}

	}, 2)

	tmp3409 := Call(__e, PrimNS2Value(symdef), symshen_4nchars, tmp3392)

	_ = tmp3409

	tmp3410 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dit_d))
		return
	}, 0)

	tmp3411 := Call(__e, PrimNS2Value(symdef), symit, tmp3410)

	_ = tmp3411

	tmp3412 := MakeNative(func(__e *ControlFlow) {
		V2734 := __e.Get(1)
		_ = V2734
		tmp3413 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp3414 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp3415 := MakeNative(func(__e *ControlFlow) {
					Bytes := __e.Get(1)
					_ = Bytes
					tmp3416 := MakeNative(func(__e *ControlFlow) {
						Close := __e.Get(1)
						_ = Close
						__e.TailApply(PrimNS2Value(symreverse), Bytes)
						return
					}, 1)

					tmp3417 := PrimCloseStream(Stream)

					__e.TailApply(tmp3416, tmp3417)
					return

				}, 1)

				tmp3418 := Call(__e, PrimNS2Value(symshen_4read_1file_1as_1bytelist_1help), Stream, Byte, Nil)

				__e.TailApply(tmp3415, tmp3418)
				return

			}, 1)

			tmp3419 := PrimReadByte(Stream)

			__e.TailApply(tmp3414, tmp3419)
			return

		}, 1)

		tmp3420 := PrimOpenStream(V2734, symin)

		__e.TailApply(tmp3413, tmp3420)
		return

	}, 1)

	tmp3421 := Call(__e, PrimNS2Value(symdef), symread_1file_1as_1bytelist, tmp3412)

	_ = tmp3421

	tmp3422 := MakeNative(func(__e *ControlFlow) {
		V2735 := __e.Get(1)
		_ = V2735
		V2736 := __e.Get(2)
		_ = V2736
		V2737 := __e.Get(3)
		_ = V2737
		tmp3426 := PrimEqual(MakeNumber(-1), V2736)

		if True == tmp3426 {
			__e.Return(V2737)
			return
		} else {
			tmp3424 := PrimReadByte(V2735)

			tmp3425 := PrimCons(V2736, V2737)

			__e.TailApply(PrimNS2Value(symshen_4read_1file_1as_1bytelist_1help), V2735, tmp3424, tmp3425)
			return

		}

	}, 3)

	tmp3427 := Call(__e, PrimNS2Value(symdef), symshen_4read_1file_1as_1bytelist_1help, tmp3422)

	_ = tmp3427

	tmp3428 := MakeNative(func(__e *ControlFlow) {
		V2738 := __e.Get(1)
		_ = V2738
		tmp3429 := MakeNative(func(__e *ControlFlow) {
			Stream := __e.Get(1)
			_ = Stream
			tmp3430 := PrimReadByte(Stream)

			__e.TailApply(PrimNS2Value(symshen_4rfas_1h), Stream, tmp3430, MakeString(""))
			return

		}, 1)

		tmp3431 := PrimOpenStream(V2738, symin)

		__e.TailApply(tmp3429, tmp3431)
		return

	}, 1)

	tmp3432 := Call(__e, PrimNS2Value(symdef), symread_1file_1as_1string, tmp3428)

	_ = tmp3432

	tmp3433 := MakeNative(func(__e *ControlFlow) {
		V2739 := __e.Get(1)
		_ = V2739
		V2740 := __e.Get(2)
		_ = V2740
		V2741 := __e.Get(3)
		_ = V2741
		tmp3439 := PrimEqual(MakeNumber(-1), V2740)

		if True == tmp3439 {
			tmp3438 := PrimCloseStream(V2739)

			_ = tmp3438

			__e.Return(V2741)
			return

		} else {
			tmp3435 := PrimReadByte(V2739)

			tmp3436 := PrimNumberToString(V2740)

			tmp3437 := PrimStringConcat(V2741, tmp3436)

			__e.TailApply(PrimNS2Value(symshen_4rfas_1h), V2739, tmp3435, tmp3437)
			return

		}

	}, 3)

	tmp3440 := Call(__e, PrimNS2Value(symdef), symshen_4rfas_1h, tmp3433)

	_ = tmp3440

	tmp3441 := MakeNative(func(__e *ControlFlow) {
		V2742 := __e.Get(1)
		_ = V2742
		tmp3442 := Call(__e, PrimNS2Value(symread), V2742)

		__e.TailApply(PrimNS2Value(symeval_1kl), tmp3442)
		return

	}, 1)

	tmp3443 := Call(__e, PrimNS2Value(symdef), syminput, tmp3441)

	_ = tmp3443

	tmp3444 := MakeNative(func(__e *ControlFlow) {
		V2743 := __e.Get(1)
		_ = V2743
		V2744 := __e.Get(2)
		_ = V2744
		tmp3445 := MakeNative(func(__e *ControlFlow) {
			Mono_2 := __e.Get(1)
			_ = Mono_2
			tmp3446 := MakeNative(func(__e *ControlFlow) {
				Input := __e.Get(1)
				_ = Input
				tmp3452 := Call(__e, PrimNS2Value(symshen_4rectify_1type), V2743)

				tmp3453 := Call(__e, PrimNS2Value(symshen_4typecheck), Input, tmp3452)

				tmp3454 := PrimEqual(False, tmp3453)

				if True == tmp3454 {
					tmp3448 := Call(__e, PrimNS2Value(symshen_4app), V2743, MakeString("\n"), symshen_4r)

					tmp3449 := PrimStringConcat(MakeString(" is not of type "), tmp3448)

					tmp3450 := Call(__e, PrimNS2Value(symshen_4app), Input, tmp3449, symshen_4r)

					tmp3451 := PrimStringConcat(MakeString("type error: "), tmp3450)

					__e.Return(PrimSimpleError(tmp3451))
					return

				} else {
					__e.TailApply(PrimNS2Value(symeval_1kl), Input)
					return
				}

			}, 1)

			tmp3455 := Call(__e, PrimNS2Value(symread), V2744)

			__e.TailApply(tmp3446, tmp3455)
			return

		}, 1)

		tmp3456 := Call(__e, PrimNS2Value(symshen_4monotype), V2743)

		__e.TailApply(tmp3445, tmp3456)
		return

	}, 2)

	tmp3457 := Call(__e, PrimNS2Value(symdef), syminput_7, tmp3444)

	_ = tmp3457

	tmp3458 := MakeNative(func(__e *ControlFlow) {
		V2745 := __e.Get(1)
		_ = V2745
		tmp3465 := PrimIsPair(V2745)

		if True == tmp3465 {
			tmp3464 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				__e.TailApply(PrimNS2Value(symshen_4monotype), Z)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symmap), tmp3464, V2745)
			return

		} else {
			tmp3463 := PrimIsVariable(V2745)

			if True == tmp3463 {
				tmp3461 := Call(__e, PrimNS2Value(symshen_4app), V2745, MakeString("\n"), symshen_4a)

				tmp3462 := PrimStringConcat(MakeString("input+ expects a monotype: not "), tmp3461)

				__e.Return(PrimSimpleError(tmp3462))
				return

			} else {
				__e.Return(V2745)
				return
			}

		}

	}, 1)

	tmp3466 := Call(__e, PrimNS2Value(symdef), symshen_4monotype, tmp3458)

	_ = tmp3466

	tmp3467 := MakeNative(func(__e *ControlFlow) {
		V2746 := __e.Get(1)
		_ = V2746
		tmp3468 := Call(__e, PrimNS2Value(symshen_4my_1read_1byte), V2746)

		tmp3469 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4return_2), X)
			return
		}, 1)

		__e.TailApply(PrimNS2Value(symshen_4read_1loop), V2746, tmp3468, Nil, tmp3469)
		return

	}, 1)

	tmp3470 := Call(__e, PrimNS2Value(symdef), symlineread, tmp3467)

	_ = tmp3470

	tmp3471 := MakeNative(func(__e *ControlFlow) {
		V2747 := __e.Get(1)
		_ = V2747
		tmp3472 := MakeNative(func(__e *ControlFlow) {
			Bytelist := __e.Get(1)
			_ = Bytelist
			tmp3473 := MakeNative(func(__e *ControlFlow) {
				S_1exprs := __e.Get(1)
				_ = S_1exprs
				tmp3474 := MakeNative(func(__e *ControlFlow) {
					Process := __e.Get(1)
					_ = Process
					__e.Return(Process)
					return
				}, 1)

				tmp3475 := Call(__e, PrimNS2Value(symshen_4process_1sexprs), S_1exprs)

				__e.TailApply(tmp3474, tmp3475)
				return

			}, 1)

			tmp3476 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4_5s_1exprs_6), X)
				return
			}, 1)

			tmp3477 := Call(__e, PrimNS2Value(symcompile), tmp3476, Bytelist)

			__e.TailApply(tmp3473, tmp3477)
			return

		}, 1)

		tmp3478 := Call(__e, PrimNS2Value(symshen_4str_1_6bytes), V2747)

		__e.TailApply(tmp3472, tmp3478)
		return

	}, 1)

	tmp3479 := Call(__e, PrimNS2Value(symdef), symread_1from_1string, tmp3471)

	_ = tmp3479

	tmp3480 := MakeNative(func(__e *ControlFlow) {
		V2748 := __e.Get(1)
		_ = V2748
		tmp3481 := MakeNative(func(__e *ControlFlow) {
			Bytelist := __e.Get(1)
			_ = Bytelist
			tmp3482 := MakeNative(func(__e *ControlFlow) {
				S_1exprs := __e.Get(1)
				_ = S_1exprs
				__e.Return(S_1exprs)
				return
			}, 1)

			tmp3483 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4_5s_1exprs_6), X)
				return
			}, 1)

			tmp3484 := Call(__e, PrimNS2Value(symcompile), tmp3483, Bytelist)

			__e.TailApply(tmp3482, tmp3484)
			return

		}, 1)

		tmp3485 := Call(__e, PrimNS2Value(symshen_4str_1_6bytes), V2748)

		__e.TailApply(tmp3481, tmp3485)
		return

	}, 1)

	tmp3486 := Call(__e, PrimNS2Value(symdef), symread_1from_1string_1unprocessed, tmp3480)

	_ = tmp3486

	tmp3487 := MakeNative(func(__e *ControlFlow) {
		V2749 := __e.Get(1)
		_ = V2749
		tmp3495 := PrimEqual(MakeString(""), V2749)

		if True == tmp3495 {
			__e.Return(Nil)
			return
		} else {
			tmp3494 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2749)

			if True == tmp3494 {
				tmp3490 := Call(__e, PrimNS2Value(symhdstr), V2749)

				tmp3491 := PrimStringToNumber(tmp3490)

				tmp3492 := PrimTailString(V2749)

				tmp3493 := Call(__e, PrimNS2Value(symshen_4str_1_6bytes), tmp3492)

				__e.Return(PrimCons(tmp3491, tmp3493))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4str_1_6bytes)
				return
			}

		}

	}, 1)

	tmp3496 := Call(__e, PrimNS2Value(symdef), symshen_4str_1_6bytes, tmp3487)

	_ = tmp3496

	tmp3497 := MakeNative(func(__e *ControlFlow) {
		V2750 := __e.Get(1)
		_ = V2750
		tmp3498 := Call(__e, PrimNS2Value(symshen_4my_1read_1byte), V2750)

		tmp3499 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			__e.TailApply(PrimNS2Value(symshen_4whitespace_2), X)
			return
		}, 1)

		tmp3500 := Call(__e, PrimNS2Value(symshen_4read_1loop), V2750, tmp3498, Nil, tmp3499)

		__e.Return(PrimHead(tmp3500))
		return

	}, 1)

	tmp3501 := Call(__e, PrimNS2Value(symdef), symread, tmp3497)

	_ = tmp3501

	tmp3502 := MakeNative(func(__e *ControlFlow) {
		V2751 := __e.Get(1)
		_ = V2751
		tmp3505 := Call(__e, PrimNS2Value(symshen_4char_1stinput_2), V2751)

		if True == tmp3505 {
			tmp3504 := Call(__e, PrimNS2Value(symshen_4read_1unit_1string), V2751)

			__e.Return(PrimStringToNumber(tmp3504))
			return

		} else {
			__e.Return(PrimReadByte(V2751))
			return
		}

	}, 1)

	tmp3506 := Call(__e, PrimNS2Value(symdef), symshen_4my_1read_1byte, tmp3502)

	_ = tmp3506

	tmp3507 := MakeNative(func(__e *ControlFlow) {
		V2756 := __e.Get(1)
		_ = V2756
		V2757 := __e.Get(2)
		_ = V2757
		V2758 := __e.Get(3)
		_ = V2758
		V2759 := __e.Get(4)
		_ = V2759
		tmp3530 := PrimEqual(MakeNumber(94), V2757)

		if True == tmp3530 {
			__e.Return(PrimSimpleError(MakeString("read aborted")))
			return
		} else {
			tmp3529 := PrimEqual(MakeNumber(-1), V2757)

			if True == tmp3529 {
				tmp3528 := Call(__e, PrimNS2Value(symempty_2), V2758)

				if True == tmp3528 {
					__e.Return(PrimSimpleError(MakeString("error: empty stream")))
					return
				} else {
					tmp3527 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4_5s_1exprs_6), X)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symcompile), tmp3527, V2758)
					return

				}

			} else {
				tmp3525 := PrimEqual(MakeNumber(0), V2757)

				if True == tmp3525 {
					tmp3524 := Call(__e, PrimNS2Value(symshen_4my_1read_1byte), V2756)

					__e.TailApply(PrimNS2Value(symshen_4read_1loop), V2756, tmp3524, V2758, V2759)
					return

				} else {
					tmp3523 := Call(__e, V2759, V2757)

					if True == tmp3523 {
						tmp3515 := MakeNative(func(__e *ControlFlow) {
							Parse := __e.Get(1)
							_ = Parse
							tmp3521 := Call(__e, PrimNS2Value(symshen_4nothing_1doing_2), Parse)

							if True == tmp3521 {
								tmp3518 := Call(__e, PrimNS2Value(symshen_4my_1read_1byte), V2756)

								tmp3519 := PrimCons(V2757, Nil)

								tmp3520 := Call(__e, PrimNS2Value(symappend), V2758, tmp3519)

								__e.TailApply(PrimNS2Value(symshen_4read_1loop), V2756, tmp3518, tmp3520, V2759)
								return

							} else {
								tmp3517 := Call(__e, PrimNS2Value(symshen_4record_1it), V2758)

								_ = tmp3517

								__e.Return(Parse)
								return

							}

						}, 1)

						tmp3522 := Call(__e, PrimNS2Value(symshen_4try_1parse), V2758)

						__e.TailApply(tmp3515, tmp3522)
						return

					} else {
						tmp3512 := Call(__e, PrimNS2Value(symshen_4my_1read_1byte), V2756)

						tmp3513 := PrimCons(V2757, Nil)

						tmp3514 := Call(__e, PrimNS2Value(symappend), V2758, tmp3513)

						__e.TailApply(PrimNS2Value(symshen_4read_1loop), V2756, tmp3512, tmp3514, V2759)
						return

					}

				}

			}

		}

	}, 4)

	tmp3531 := Call(__e, PrimNS2Value(symdef), symshen_4read_1loop, tmp3507)

	_ = tmp3531

	tmp3532 := MakeNative(func(__e *ControlFlow) {
		V2760 := __e.Get(1)
		_ = V2760
		tmp3533 := MakeNative(func(__e *ControlFlow) {
			S_1exprs := __e.Get(1)
			_ = S_1exprs
			tmp3535 := Call(__e, PrimNS2Value(symshen_4nothing_1doing_2), S_1exprs)

			if True == tmp3535 {
				__e.Return(symshen_4i_1failed_b)
				return
			} else {
				__e.TailApply(PrimNS2Value(symshen_4process_1sexprs), S_1exprs)
				return
			}

		}, 1)

		tmp3536 := MakeNative(func(__e *ControlFlow) {
			tmp3537 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4_5s_1exprs_6), X)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symcompile), tmp3537, V2760)
			return

		}, 0)

		tmp3538 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(symshen_4i_1failed_b)
			return
		}, 1)

		tmp3539 := Call(__e, PrimNS1Value(symtry_1catch), tmp3536, tmp3538)

		__e.TailApply(tmp3533, tmp3539)
		return

	}, 1)

	tmp3540 := Call(__e, PrimNS2Value(symdef), symshen_4try_1parse, tmp3532)

	_ = tmp3540

	tmp3541 := MakeNative(func(__e *ControlFlow) {
		V2763 := __e.Get(1)
		_ = V2763
		tmp3545 := PrimEqual(symshen_4i_1failed_b, V2763)

		if True == tmp3545 {
			__e.Return(True)
			return
		} else {
			tmp3544 := PrimEqual(Nil, V2763)

			if True == tmp3544 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp3546 := Call(__e, PrimNS2Value(symdef), symshen_4nothing_1doing_2, tmp3541)

	_ = tmp3546

	tmp3547 := MakeNative(func(__e *ControlFlow) {
		V2764 := __e.Get(1)
		_ = V2764
		tmp3548 := Call(__e, PrimNS2Value(symshen_4bytes_1_6string), V2764)

		__e.Return(PrimNS3Set(symshen_4_dit_d, tmp3548))
		return

	}, 1)

	tmp3549 := Call(__e, PrimNS2Value(symdef), symshen_4record_1it, tmp3547)

	_ = tmp3549

	tmp3550 := MakeNative(func(__e *ControlFlow) {
		V2765 := __e.Get(1)
		_ = V2765
		tmp3558 := PrimEqual(Nil, V2765)

		if True == tmp3558 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp3557 := PrimIsPair(V2765)

			if True == tmp3557 {
				tmp3553 := PrimHead(V2765)

				tmp3554 := PrimNumberToString(tmp3553)

				tmp3555 := PrimTail(V2765)

				tmp3556 := Call(__e, PrimNS2Value(symshen_4bytes_1_6string), tmp3555)

				__e.Return(PrimStringConcat(tmp3554, tmp3556))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4bytes_1_6string)
				return
			}

		}

	}, 1)

	tmp3559 := Call(__e, PrimNS2Value(symdef), symshen_4bytes_1_6string, tmp3550)

	_ = tmp3559

	tmp3560 := MakeNative(func(__e *ControlFlow) {
		V2766 := __e.Get(1)
		_ = V2766
		tmp3561 := MakeNative(func(__e *ControlFlow) {
			Unpack_eExpand := __e.Get(1)
			_ = Unpack_eExpand
			tmp3562 := MakeNative(func(__e *ControlFlow) {
				FindArities := __e.Get(1)
				_ = FindArities
				tmp3563 := MakeNative(func(__e *ControlFlow) {
					Types := __e.Get(1)
					_ = Types
					tmp3564 := MakeNative(func(__e *ControlFlow) {
						X := __e.Get(1)
						_ = X
						__e.TailApply(PrimNS2Value(symshen_4process_1applications), X, Types)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symmap), tmp3564, Unpack_eExpand)
					return

				}, 1)

				tmp3565 := Call(__e, PrimNS2Value(symshen_4find_1types), Unpack_eExpand)

				__e.TailApply(tmp3563, tmp3565)
				return

			}, 1)

			tmp3566 := Call(__e, PrimNS2Value(symshen_4find_1arities), Unpack_eExpand)

			__e.TailApply(tmp3562, tmp3566)
			return

		}, 1)

		tmp3567 := Call(__e, PrimNS2Value(symshen_4unpackage_emacroexpand), V2766)

		__e.TailApply(tmp3561, tmp3567)
		return

	}, 1)

	tmp3568 := Call(__e, PrimNS2Value(symdef), symshen_4process_1sexprs, tmp3560)

	_ = tmp3568

	tmp3569 := MakeNative(func(__e *ControlFlow) {
		V2767 := __e.Get(1)
		_ = V2767
		tmp3591 := PrimIsPair(V2767)

		var ifres3582 Obj

		if True == tmp3591 {
			tmp3589 := PrimTail(V2767)

			tmp3590 := PrimIsPair(tmp3589)

			var ifres3584 Obj

			if True == tmp3590 {
				tmp3586 := PrimHead(V2767)

				tmp3587 := PrimIntern(MakeString(":"))

				tmp3588 := PrimEqual(tmp3586, tmp3587)

				var ifres3585 Obj

				if True == tmp3588 {
					ifres3585 = True

				} else {
					ifres3585 = False

				}

				ifres3584 = ifres3585

			} else {
				ifres3584 = False

			}

			var ifres3583 Obj

			if True == ifres3584 {
				ifres3583 = True

			} else {
				ifres3583 = False

			}

			ifres3582 = ifres3583

		} else {
			ifres3582 = False

		}

		if True == ifres3582 {
			tmp3577 := PrimTail(V2767)

			tmp3578 := PrimHead(tmp3577)

			tmp3579 := PrimTail(V2767)

			tmp3580 := PrimTail(tmp3579)

			tmp3581 := Call(__e, PrimNS2Value(symshen_4find_1types), tmp3580)

			__e.Return(PrimCons(tmp3578, tmp3581))
			return

		} else {
			tmp3576 := PrimIsPair(V2767)

			if True == tmp3576 {
				tmp3572 := PrimHead(V2767)

				tmp3573 := Call(__e, PrimNS2Value(symshen_4find_1types), tmp3572)

				tmp3574 := PrimTail(V2767)

				tmp3575 := Call(__e, PrimNS2Value(symshen_4find_1types), tmp3574)

				__e.TailApply(PrimNS2Value(symappend), tmp3573, tmp3575)
				return

			} else {
				__e.Return(Nil)
				return
			}

		}

	}, 1)

	tmp3592 := Call(__e, PrimNS2Value(symdef), symshen_4find_1types, tmp3569)

	_ = tmp3592

	tmp3593 := MakeNative(func(__e *ControlFlow) {
		V2770 := __e.Get(1)
		_ = V2770
		tmp3642 := PrimIsPair(V2770)

		var ifres3623 Obj

		if True == tmp3642 {
			tmp3640 := PrimHead(V2770)

			tmp3641 := PrimEqual(symdefine, tmp3640)

			var ifres3625 Obj

			if True == tmp3641 {
				tmp3638 := PrimTail(V2770)

				tmp3639 := PrimIsPair(tmp3638)

				var ifres3627 Obj

				if True == tmp3639 {
					tmp3635 := PrimTail(V2770)

					tmp3636 := PrimTail(tmp3635)

					tmp3637 := PrimIsPair(tmp3636)

					var ifres3629 Obj

					if True == tmp3637 {
						tmp3631 := PrimTail(V2770)

						tmp3632 := PrimTail(tmp3631)

						tmp3633 := PrimHead(tmp3632)

						tmp3634 := PrimEqual(sym_i, tmp3633)

						var ifres3630 Obj

						if True == tmp3634 {
							ifres3630 = True

						} else {
							ifres3630 = False

						}

						ifres3629 = ifres3630

					} else {
						ifres3629 = False

					}

					var ifres3628 Obj

					if True == ifres3629 {
						ifres3628 = True

					} else {
						ifres3628 = False

					}

					ifres3627 = ifres3628

				} else {
					ifres3627 = False

				}

				var ifres3626 Obj

				if True == ifres3627 {
					ifres3626 = True

				} else {
					ifres3626 = False

				}

				ifres3625 = ifres3626

			} else {
				ifres3625 = False

			}

			var ifres3624 Obj

			if True == ifres3625 {
				ifres3624 = True

			} else {
				ifres3624 = False

			}

			ifres3623 = ifres3624

		} else {
			ifres3623 = False

		}

		if True == ifres3623 {
			tmp3615 := PrimTail(V2770)

			tmp3616 := PrimHead(tmp3615)

			tmp3617 := PrimTail(V2770)

			tmp3618 := PrimHead(tmp3617)

			tmp3619 := PrimTail(V2770)

			tmp3620 := PrimTail(tmp3619)

			tmp3621 := PrimTail(tmp3620)

			tmp3622 := Call(__e, PrimNS2Value(symshen_4find_1arity), tmp3618, MakeNumber(1), tmp3621)

			__e.TailApply(PrimNS2Value(symshen_4store_1arity), tmp3616, tmp3622)
			return

		} else {
			tmp3614 := PrimIsPair(V2770)

			var ifres3606 Obj

			if True == tmp3614 {
				tmp3612 := PrimHead(V2770)

				tmp3613 := PrimEqual(symdefine, tmp3612)

				var ifres3608 Obj

				if True == tmp3613 {
					tmp3610 := PrimTail(V2770)

					tmp3611 := PrimIsPair(tmp3610)

					var ifres3609 Obj

					if True == tmp3611 {
						ifres3609 = True

					} else {
						ifres3609 = False

					}

					ifres3608 = ifres3609

				} else {
					ifres3608 = False

				}

				var ifres3607 Obj

				if True == ifres3608 {
					ifres3607 = True

				} else {
					ifres3607 = False

				}

				ifres3606 = ifres3607

			} else {
				ifres3606 = False

			}

			if True == ifres3606 {
				tmp3599 := PrimTail(V2770)

				tmp3600 := PrimHead(tmp3599)

				tmp3601 := PrimTail(V2770)

				tmp3602 := PrimHead(tmp3601)

				tmp3603 := PrimTail(V2770)

				tmp3604 := PrimTail(tmp3603)

				tmp3605 := Call(__e, PrimNS2Value(symshen_4find_1arity), tmp3602, MakeNumber(0), tmp3604)

				__e.TailApply(PrimNS2Value(symshen_4store_1arity), tmp3600, tmp3605)
				return

			} else {
				tmp3598 := PrimIsPair(V2770)

				if True == tmp3598 {
					tmp3597 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						__e.TailApply(PrimNS2Value(symshen_4find_1arities), Z)
						return
					}, 1)

					__e.TailApply(PrimNS2Value(symmap), tmp3597, V2770)
					return

				} else {
					__e.Return(symshen_4skip)
					return
				}

			}

		}

	}, 1)

	tmp3643 := Call(__e, PrimNS2Value(symdef), symshen_4find_1arities, tmp3593)

	_ = tmp3643

	tmp3644 := MakeNative(func(__e *ControlFlow) {
		V2771 := __e.Get(1)
		_ = V2771
		V2772 := __e.Get(2)
		_ = V2772
		tmp3645 := MakeNative(func(__e *ControlFlow) {
			ArityF := __e.Get(1)
			_ = ArityF
			tmp3653 := PrimEqual(ArityF, MakeNumber(-1))

			if True == tmp3653 {
				__e.TailApply(PrimNS2Value(symshen_4execute_1store_1arity), V2771, V2772)
				return
			} else {
				tmp3652 := PrimEqual(ArityF, V2772)

				if True == tmp3652 {
					__e.Return(symshen_4skip)
					return
				} else {
					tmp3648 := Call(__e, PrimNS2Value(symshen_4app), V2771, MakeString(" may cause errors\n"), symshen_4a)

					tmp3649 := PrimStringConcat(MakeString("changing the arity of "), tmp3648)

					tmp3650 := Call(__e, PrimNS2Value(symstoutput))

					tmp3651 := Call(__e, PrimNS2Value(sympr), tmp3649, tmp3650)

					_ = tmp3651

					__e.TailApply(PrimNS2Value(symshen_4execute_1store_1arity), V2771, V2772)
					return

				}

			}

		}, 1)

		tmp3654 := Call(__e, PrimNS2Value(symarity), V2771)

		__e.TailApply(tmp3645, tmp3654)
		return

	}, 2)

	tmp3655 := Call(__e, PrimNS2Value(symdef), symshen_4store_1arity, tmp3644)

	_ = tmp3655

	tmp3656 := MakeNative(func(__e *ControlFlow) {
		V2773 := __e.Get(1)
		_ = V2773
		V2774 := __e.Get(2)
		_ = V2774
		tmp3661 := PrimEqual(MakeNumber(0), V2774)

		if True == tmp3661 {
			tmp3660 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symput), V2773, symarity, MakeNumber(0), tmp3660)
			return

		} else {
			tmp3658 := PrimNS3Value(sym_dproperty_1vector_d)

			tmp3659 := Call(__e, PrimNS2Value(symput), V2773, symarity, V2774, tmp3658)

			_ = tmp3659

			__e.TailApply(PrimNS2Value(symshen_4update_1lambdatable), V2773, V2774)
			return

		}

	}, 2)

	tmp3662 := Call(__e, PrimNS2Value(symdef), symshen_4execute_1store_1arity, tmp3656)

	_ = tmp3662

	tmp3663 := MakeNative(func(__e *ControlFlow) {
		V2775 := __e.Get(1)
		_ = V2775
		V2776 := __e.Get(2)
		_ = V2776
		tmp3664 := MakeNative(func(__e *ControlFlow) {
			LambdaTable := __e.Get(1)
			_ = LambdaTable
			tmp3665 := MakeNative(func(__e *ControlFlow) {
				Lambda := __e.Get(1)
				_ = Lambda
				tmp3666 := MakeNative(func(__e *ControlFlow) {
					Insert := __e.Get(1)
					_ = Insert
					tmp3667 := MakeNative(func(__e *ControlFlow) {
						Reset := __e.Get(1)
						_ = Reset
						__e.Return(Reset)
						return
					}, 1)

					tmp3668 := PrimNS3Set(symshen_4_dlambdatable_d, Insert)

					__e.TailApply(tmp3667, tmp3668)
					return

				}, 1)

				tmp3669 := Call(__e, PrimNS2Value(symshen_4assoc_1_6), V2775, Lambda, LambdaTable)

				__e.TailApply(tmp3666, tmp3669)
				return

			}, 1)

			tmp3670 := PrimCons(V2775, Nil)

			tmp3671 := Call(__e, PrimNS2Value(symshen_4lambda_1function), tmp3670, V2776)

			tmp3672 := Call(__e, PrimNS2Value(symeval_1kl), tmp3671)

			__e.TailApply(tmp3665, tmp3672)
			return

		}, 1)

		tmp3673 := PrimNS3Value(symshen_4_dlambdatable_d)

		__e.TailApply(tmp3664, tmp3673)
		return

	}, 2)

	tmp3674 := Call(__e, PrimNS2Value(symdef), symshen_4update_1lambdatable, tmp3663)

	_ = tmp3674

	tmp3675 := MakeNative(func(__e *ControlFlow) {
		V2779 := __e.Get(1)
		_ = V2779
		V2780 := __e.Get(2)
		_ = V2780
		tmp3695 := PrimEqual(MakeNumber(0), V2780)

		if True == tmp3695 {
			__e.Return(symshen_4skip)
			return
		} else {
			tmp3694 := PrimEqual(MakeNumber(1), V2780)

			if True == tmp3694 {
				tmp3687 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp3688 := PrimCons(X, Nil)

					tmp3689 := Call(__e, PrimNS2Value(symappend), V2779, tmp3688)

					tmp3690 := PrimCons(tmp3689, Nil)

					tmp3691 := PrimCons(X, tmp3690)

					__e.Return(PrimCons(symlambda, tmp3691))
					return

				}, 1)

				tmp3692 := Call(__e, PrimNS2Value(symgensym), symY)

				tmp3693 := Call(__e, PrimNS2Value(symprotect), tmp3692)

				__e.TailApply(tmp3687, tmp3693)
				return

			} else {
				tmp3678 := MakeNative(func(__e *ControlFlow) {
					X := __e.Get(1)
					_ = X
					tmp3679 := PrimCons(X, Nil)

					tmp3680 := Call(__e, PrimNS2Value(symappend), V2779, tmp3679)

					tmp3681 := PrimNumberSubtract(V2780, MakeNumber(1))

					tmp3682 := Call(__e, PrimNS2Value(symshen_4lambda_1function), tmp3680, tmp3681)

					tmp3683 := PrimCons(tmp3682, Nil)

					tmp3684 := PrimCons(X, tmp3683)

					__e.Return(PrimCons(symlambda, tmp3684))
					return

				}, 1)

				tmp3685 := Call(__e, PrimNS2Value(symgensym), symY)

				tmp3686 := Call(__e, PrimNS2Value(symprotect), tmp3685)

				__e.TailApply(tmp3678, tmp3686)
				return

			}

		}

	}, 2)

	tmp3696 := Call(__e, PrimNS2Value(symdef), symshen_4lambda_1function, tmp3675)

	_ = tmp3696

	tmp3697 := MakeNative(func(__e *ControlFlow) {
		V2790 := __e.Get(1)
		_ = V2790
		V2791 := __e.Get(2)
		_ = V2791
		V2792 := __e.Get(3)
		_ = V2792
		tmp3720 := PrimEqual(Nil, V2792)

		if True == tmp3720 {
			tmp3719 := PrimCons(V2790, V2791)

			__e.Return(PrimCons(tmp3719, Nil))
			return

		} else {
			tmp3718 := PrimIsPair(V2792)

			var ifres3709 Obj

			if True == tmp3718 {
				tmp3716 := PrimHead(V2792)

				tmp3717 := PrimIsPair(tmp3716)

				var ifres3711 Obj

				if True == tmp3717 {
					tmp3713 := PrimHead(V2792)

					tmp3714 := PrimHead(tmp3713)

					tmp3715 := PrimEqual(V2790, tmp3714)

					var ifres3712 Obj

					if True == tmp3715 {
						ifres3712 = True

					} else {
						ifres3712 = False

					}

					ifres3711 = ifres3712

				} else {
					ifres3711 = False

				}

				var ifres3710 Obj

				if True == ifres3711 {
					ifres3710 = True

				} else {
					ifres3710 = False

				}

				ifres3709 = ifres3710

			} else {
				ifres3709 = False

			}

			if True == ifres3709 {
				tmp3705 := PrimHead(V2792)

				tmp3706 := PrimHead(tmp3705)

				tmp3707 := PrimCons(tmp3706, V2791)

				tmp3708 := PrimTail(V2792)

				__e.Return(PrimCons(tmp3707, tmp3708))
				return

			} else {
				tmp3704 := PrimIsPair(V2792)

				if True == tmp3704 {
					tmp3701 := PrimHead(V2792)

					tmp3702 := PrimTail(V2792)

					tmp3703 := Call(__e, PrimNS2Value(symshen_4assoc_1_6), V2790, V2791, tmp3702)

					__e.Return(PrimCons(tmp3701, tmp3703))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.assoc->")))
					return
				}

			}

		}

	}, 3)

	tmp3721 := Call(__e, PrimNS2Value(symdef), symshen_4assoc_1_6, tmp3697)

	_ = tmp3721

	tmp3722 := MakeNative(func(__e *ControlFlow) {
		V2807 := __e.Get(1)
		_ = V2807
		V2808 := __e.Get(2)
		_ = V2808
		V2809 := __e.Get(3)
		_ = V2809
		tmp3769 := PrimEqual(MakeNumber(0), V2808)

		var ifres3762 Obj

		if True == tmp3769 {
			tmp3768 := PrimIsPair(V2809)

			var ifres3764 Obj

			if True == tmp3768 {
				tmp3766 := PrimHead(V2809)

				tmp3767 := PrimEqual(tmp3766, sym_1_6)

				var ifres3765 Obj

				if True == tmp3767 {
					ifres3765 = True

				} else {
					ifres3765 = False

				}

				ifres3764 = ifres3765

			} else {
				ifres3764 = False

			}

			var ifres3763 Obj

			if True == ifres3764 {
				ifres3763 = True

			} else {
				ifres3763 = False

			}

			ifres3762 = ifres3763

		} else {
			ifres3762 = False

		}

		if True == ifres3762 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp3761 := PrimEqual(MakeNumber(0), V2808)

			var ifres3754 Obj

			if True == tmp3761 {
				tmp3760 := PrimIsPair(V2809)

				var ifres3756 Obj

				if True == tmp3760 {
					tmp3758 := PrimHead(V2809)

					tmp3759 := PrimEqual(tmp3758, sym_5_1)

					var ifres3757 Obj

					if True == tmp3759 {
						ifres3757 = True

					} else {
						ifres3757 = False

					}

					ifres3756 = ifres3757

				} else {
					ifres3756 = False

				}

				var ifres3755 Obj

				if True == ifres3756 {
					ifres3755 = True

				} else {
					ifres3755 = False

				}

				ifres3754 = ifres3755

			} else {
				ifres3754 = False

			}

			if True == ifres3754 {
				__e.Return(MakeNumber(0))
				return
			} else {
				tmp3753 := PrimEqual(MakeNumber(0), V2808)

				var ifres3750 Obj

				if True == tmp3753 {
					tmp3752 := PrimIsPair(V2809)

					var ifres3751 Obj

					if True == tmp3752 {
						ifres3751 = True

					} else {
						ifres3751 = False

					}

					ifres3750 = ifres3751

				} else {
					ifres3750 = False

				}

				if True == ifres3750 {
					tmp3748 := PrimTail(V2809)

					tmp3749 := Call(__e, PrimNS2Value(symshen_4find_1arity), V2807, MakeNumber(0), tmp3748)

					__e.Return(PrimNumberAdd(MakeNumber(1), tmp3749))
					return

				} else {
					tmp3747 := PrimEqual(MakeNumber(1), V2808)

					var ifres3740 Obj

					if True == tmp3747 {
						tmp3746 := PrimIsPair(V2809)

						var ifres3742 Obj

						if True == tmp3746 {
							tmp3744 := PrimHead(V2809)

							tmp3745 := PrimEqual(sym_j, tmp3744)

							var ifres3743 Obj

							if True == tmp3745 {
								ifres3743 = True

							} else {
								ifres3743 = False

							}

							ifres3742 = ifres3743

						} else {
							ifres3742 = False

						}

						var ifres3741 Obj

						if True == ifres3742 {
							ifres3741 = True

						} else {
							ifres3741 = False

						}

						ifres3740 = ifres3741

					} else {
						ifres3740 = False

					}

					if True == ifres3740 {
						tmp3739 := PrimTail(V2809)

						__e.TailApply(PrimNS2Value(symshen_4find_1arity), V2807, MakeNumber(0), tmp3739)
						return

					} else {
						tmp3738 := PrimEqual(MakeNumber(1), V2808)

						var ifres3735 Obj

						if True == tmp3738 {
							tmp3737 := PrimIsPair(V2809)

							var ifres3736 Obj

							if True == tmp3737 {
								ifres3736 = True

							} else {
								ifres3736 = False

							}

							ifres3735 = ifres3736

						} else {
							ifres3735 = False

						}

						if True == ifres3735 {
							tmp3734 := PrimTail(V2809)

							__e.TailApply(PrimNS2Value(symshen_4find_1arity), V2807, MakeNumber(1), tmp3734)
							return

						} else {
							tmp3733 := PrimEqual(MakeNumber(1), V2808)

							if True == tmp3733 {
								tmp3731 := Call(__e, PrimNS2Value(symshen_4app), V2807, MakeString(" definition: missing }\n"), symshen_4a)

								tmp3732 := PrimStringConcat(MakeString("syntax error in "), tmp3731)

								__e.Return(PrimSimpleError(tmp3732))
								return

							} else {
								tmp3729 := Call(__e, PrimNS2Value(symshen_4app), V2807, MakeString(" definition: missing -> or <-\n"), symshen_4a)

								tmp3730 := PrimStringConcat(MakeString("syntax error in "), tmp3729)

								__e.Return(PrimSimpleError(tmp3730))
								return

							}

						}

					}

				}

			}

		}

	}, 3)

	tmp3770 := Call(__e, PrimNS2Value(symdef), symshen_4find_1arity, tmp3722)

	_ = tmp3770

	tmp3771 := MakeNative(func(__e *ControlFlow) {
		V2810 := __e.Get(1)
		_ = V2810
		tmp3772 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp3964 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp3964 {
				tmp3774 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp3942 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp3942 {
						tmp3776 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp3929 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp3929 {
								tmp3778 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp3916 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

									if True == tmp3916 {
										tmp3780 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp3903 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

											if True == tmp3903 {
												tmp3782 := MakeNative(func(__e *ControlFlow) {
													Result := __e.Get(1)
													_ = Result
													tmp3889 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

													if True == tmp3889 {
														tmp3784 := MakeNative(func(__e *ControlFlow) {
															Result := __e.Get(1)
															_ = Result
															tmp3871 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

															if True == tmp3871 {
																tmp3786 := MakeNative(func(__e *ControlFlow) {
																	Result := __e.Get(1)
																	_ = Result
																	tmp3857 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

																	if True == tmp3857 {
																		tmp3788 := MakeNative(func(__e *ControlFlow) {
																			Result := __e.Get(1)
																			_ = Result
																			tmp3843 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

																			if True == tmp3843 {
																				tmp3790 := MakeNative(func(__e *ControlFlow) {
																					Result := __e.Get(1)
																					_ = Result
																					tmp3831 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

																					if True == tmp3831 {
																						tmp3792 := MakeNative(func(__e *ControlFlow) {
																							Result := __e.Get(1)
																							_ = Result
																							tmp3817 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

																							if True == tmp3817 {
																								tmp3794 := MakeNative(func(__e *ControlFlow) {
																									Result := __e.Get(1)
																									_ = Result
																									tmp3805 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

																									if True == tmp3805 {
																										tmp3796 := MakeNative(func(__e *ControlFlow) {
																											Result := __e.Get(1)
																											_ = Result
																											tmp3798 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

																											if True == tmp3798 {
																												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																												return
																											} else {
																												__e.Return(Result)
																												return
																											}

																										}, 1)

																										tmp3799 := MakeNative(func(__e *ControlFlow) {
																											Parse_5e_6 := __e.Get(1)
																											_ = Parse_5e_6
																											tmp3802 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

																											if True == tmp3802 {
																												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																												return
																											} else {
																												tmp3801 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

																												__e.TailApply(PrimNS2Value(symshen_4comb), tmp3801, Nil)
																												return

																											}

																										}, 1)

																										tmp3803 := Call(__e, PrimNS2Value(sym_5e_6), V2810)

																										tmp3804 := Call(__e, tmp3799, tmp3803)

																										__e.TailApply(tmp3796, tmp3804)
																										return

																									} else {
																										__e.Return(Result)
																										return
																									}

																								}, 1)

																								tmp3806 := MakeNative(func(__e *ControlFlow) {
																									Parseshen_4_5whitespaces_6 := __e.Get(1)
																									_ = Parseshen_4_5whitespaces_6
																									tmp3814 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5whitespaces_6)

																									if True == tmp3814 {
																										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																										return
																									} else {
																										tmp3808 := MakeNative(func(__e *ControlFlow) {
																											Parseshen_4_5s_1exprs_6 := __e.Get(1)
																											_ = Parseshen_4_5s_1exprs_6
																											tmp3812 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																											if True == tmp3812 {
																												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																												return
																											} else {
																												tmp3810 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																												tmp3811 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																												__e.TailApply(PrimNS2Value(symshen_4comb), tmp3810, tmp3811)
																												return

																											}

																										}, 1)

																										tmp3813 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5whitespaces_6)

																										__e.TailApply(tmp3808, tmp3813)
																										return

																									}

																								}, 1)

																								tmp3815 := Call(__e, PrimNS2Value(symshen_4_5whitespaces_6), V2810)

																								tmp3816 := Call(__e, tmp3806, tmp3815)

																								__e.TailApply(tmp3794, tmp3816)
																								return

																							} else {
																								__e.Return(Result)
																								return
																							}

																						}, 1)

																						tmp3818 := MakeNative(func(__e *ControlFlow) {
																							Parseshen_4_5atom_6 := __e.Get(1)
																							_ = Parseshen_4_5atom_6
																							tmp3828 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5atom_6)

																							if True == tmp3828 {
																								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																								return
																							} else {
																								tmp3820 := MakeNative(func(__e *ControlFlow) {
																									Parseshen_4_5s_1exprs_6 := __e.Get(1)
																									_ = Parseshen_4_5s_1exprs_6
																									tmp3826 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																									if True == tmp3826 {
																										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																										return
																									} else {
																										tmp3822 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																										tmp3823 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5atom_6)

																										tmp3824 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																										tmp3825 := PrimCons(tmp3823, tmp3824)

																										__e.TailApply(PrimNS2Value(symshen_4comb), tmp3822, tmp3825)
																										return

																									}

																								}, 1)

																								tmp3827 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5atom_6)

																								__e.TailApply(tmp3820, tmp3827)
																								return

																							}

																						}, 1)

																						tmp3829 := Call(__e, PrimNS2Value(symshen_4_5atom_6), V2810)

																						tmp3830 := Call(__e, tmp3818, tmp3829)

																						__e.TailApply(tmp3792, tmp3830)
																						return

																					} else {
																						__e.Return(Result)
																						return
																					}

																				}, 1)

																				tmp3832 := MakeNative(func(__e *ControlFlow) {
																					Parseshen_4_5comment_6 := __e.Get(1)
																					_ = Parseshen_4_5comment_6
																					tmp3840 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5comment_6)

																					if True == tmp3840 {
																						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																						return
																					} else {
																						tmp3834 := MakeNative(func(__e *ControlFlow) {
																							Parseshen_4_5s_1exprs_6 := __e.Get(1)
																							_ = Parseshen_4_5s_1exprs_6
																							tmp3838 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																							if True == tmp3838 {
																								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																								return
																							} else {
																								tmp3836 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																								tmp3837 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																								__e.TailApply(PrimNS2Value(symshen_4comb), tmp3836, tmp3837)
																								return

																							}

																						}, 1)

																						tmp3839 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5comment_6)

																						__e.TailApply(tmp3834, tmp3839)
																						return

																					}

																				}, 1)

																				tmp3841 := Call(__e, PrimNS2Value(symshen_4_5comment_6), V2810)

																				tmp3842 := Call(__e, tmp3832, tmp3841)

																				__e.TailApply(tmp3790, tmp3842)
																				return

																			} else {
																				__e.Return(Result)
																				return
																			}

																		}, 1)

																		tmp3844 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5comma_6 := __e.Get(1)
																			_ = Parseshen_4_5comma_6
																			tmp3854 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5comma_6)

																			if True == tmp3854 {
																				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																				return
																			} else {
																				tmp3846 := MakeNative(func(__e *ControlFlow) {
																					Parseshen_4_5s_1exprs_6 := __e.Get(1)
																					_ = Parseshen_4_5s_1exprs_6
																					tmp3852 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																					if True == tmp3852 {
																						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																						return
																					} else {
																						tmp3848 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																						tmp3849 := PrimIntern(MakeString(","))

																						tmp3850 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																						tmp3851 := PrimCons(tmp3849, tmp3850)

																						__e.TailApply(PrimNS2Value(symshen_4comb), tmp3848, tmp3851)
																						return

																					}

																				}, 1)

																				tmp3853 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5comma_6)

																				__e.TailApply(tmp3846, tmp3853)
																				return

																			}

																		}, 1)

																		tmp3855 := Call(__e, PrimNS2Value(symshen_4_5comma_6), V2810)

																		tmp3856 := Call(__e, tmp3844, tmp3855)

																		__e.TailApply(tmp3788, tmp3856)
																		return

																	} else {
																		__e.Return(Result)
																		return
																	}

																}, 1)

																tmp3858 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5colon_6 := __e.Get(1)
																	_ = Parseshen_4_5colon_6
																	tmp3868 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5colon_6)

																	if True == tmp3868 {
																		__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																		return
																	} else {
																		tmp3860 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5s_1exprs_6 := __e.Get(1)
																			_ = Parseshen_4_5s_1exprs_6
																			tmp3866 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																			if True == tmp3866 {
																				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																				return
																			} else {
																				tmp3862 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																				tmp3863 := PrimIntern(MakeString(":"))

																				tmp3864 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																				tmp3865 := PrimCons(tmp3863, tmp3864)

																				__e.TailApply(PrimNS2Value(symshen_4comb), tmp3862, tmp3865)
																				return

																			}

																		}, 1)

																		tmp3867 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5colon_6)

																		__e.TailApply(tmp3860, tmp3867)
																		return

																	}

																}, 1)

																tmp3869 := Call(__e, PrimNS2Value(symshen_4_5colon_6), V2810)

																tmp3870 := Call(__e, tmp3858, tmp3869)

																__e.TailApply(tmp3786, tmp3870)
																return

															} else {
																__e.Return(Result)
																return
															}

														}, 1)

														tmp3872 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5colon_6 := __e.Get(1)
															_ = Parseshen_4_5colon_6
															tmp3886 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5colon_6)

															if True == tmp3886 {
																__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																return
															} else {
																tmp3874 := MakeNative(func(__e *ControlFlow) {
																	Parseshen_4_5equal_6 := __e.Get(1)
																	_ = Parseshen_4_5equal_6
																	tmp3884 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5equal_6)

																	if True == tmp3884 {
																		__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																		return
																	} else {
																		tmp3876 := MakeNative(func(__e *ControlFlow) {
																			Parseshen_4_5s_1exprs_6 := __e.Get(1)
																			_ = Parseshen_4_5s_1exprs_6
																			tmp3882 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

																			if True == tmp3882 {
																				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																				return
																			} else {
																				tmp3878 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																				tmp3879 := PrimIntern(MakeString(":="))

																				tmp3880 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																				tmp3881 := PrimCons(tmp3879, tmp3880)

																				__e.TailApply(PrimNS2Value(symshen_4comb), tmp3878, tmp3881)
																				return

																			}

																		}, 1)

																		tmp3883 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5equal_6)

																		__e.TailApply(tmp3876, tmp3883)
																		return

																	}

																}, 1)

																tmp3885 := Call(__e, PrimNS2Value(symshen_4_5equal_6), Parseshen_4_5colon_6)

																__e.TailApply(tmp3874, tmp3885)
																return

															}

														}, 1)

														tmp3887 := Call(__e, PrimNS2Value(symshen_4_5colon_6), V2810)

														tmp3888 := Call(__e, tmp3872, tmp3887)

														__e.TailApply(tmp3784, tmp3888)
														return

													} else {
														__e.Return(Result)
														return
													}

												}, 1)

												tmp3890 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5semicolon_6 := __e.Get(1)
													_ = Parseshen_4_5semicolon_6
													tmp3900 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5semicolon_6)

													if True == tmp3900 {
														__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
														return
													} else {
														tmp3892 := MakeNative(func(__e *ControlFlow) {
															Parseshen_4_5s_1exprs_6 := __e.Get(1)
															_ = Parseshen_4_5s_1exprs_6
															tmp3898 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

															if True == tmp3898 {
																__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
																return
															} else {
																tmp3894 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

																tmp3895 := PrimIntern(MakeString(";"))

																tmp3896 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

																tmp3897 := PrimCons(tmp3895, tmp3896)

																__e.TailApply(PrimNS2Value(symshen_4comb), tmp3894, tmp3897)
																return

															}

														}, 1)

														tmp3899 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5semicolon_6)

														__e.TailApply(tmp3892, tmp3899)
														return

													}

												}, 1)

												tmp3901 := Call(__e, PrimNS2Value(symshen_4_5semicolon_6), V2810)

												tmp3902 := Call(__e, tmp3890, tmp3901)

												__e.TailApply(tmp3782, tmp3902)
												return

											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp3904 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5bar_6 := __e.Get(1)
											_ = Parseshen_4_5bar_6
											tmp3913 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5bar_6)

											if True == tmp3913 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp3906 := MakeNative(func(__e *ControlFlow) {
													Parseshen_4_5s_1exprs_6 := __e.Get(1)
													_ = Parseshen_4_5s_1exprs_6
													tmp3911 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

													if True == tmp3911 {
														__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
														return
													} else {
														tmp3908 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

														tmp3909 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

														tmp3910 := PrimCons(symbar_b, tmp3909)

														__e.TailApply(PrimNS2Value(symshen_4comb), tmp3908, tmp3910)
														return

													}

												}, 1)

												tmp3912 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5bar_6)

												__e.TailApply(tmp3906, tmp3912)
												return

											}

										}, 1)

										tmp3914 := Call(__e, PrimNS2Value(symshen_4_5bar_6), V2810)

										tmp3915 := Call(__e, tmp3904, tmp3914)

										__e.TailApply(tmp3780, tmp3915)
										return

									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp3917 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5rcurly_6 := __e.Get(1)
									_ = Parseshen_4_5rcurly_6
									tmp3926 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rcurly_6)

									if True == tmp3926 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp3919 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5s_1exprs_6 := __e.Get(1)
											_ = Parseshen_4_5s_1exprs_6
											tmp3924 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

											if True == tmp3924 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp3921 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

												tmp3922 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

												tmp3923 := PrimCons(sym_j, tmp3922)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp3921, tmp3923)
												return

											}

										}, 1)

										tmp3925 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5rcurly_6)

										__e.TailApply(tmp3919, tmp3925)
										return

									}

								}, 1)

								tmp3927 := Call(__e, PrimNS2Value(symshen_4_5rcurly_6), V2810)

								tmp3928 := Call(__e, tmp3917, tmp3927)

								__e.TailApply(tmp3778, tmp3928)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp3930 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5lcurly_6 := __e.Get(1)
							_ = Parseshen_4_5lcurly_6
							tmp3939 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5lcurly_6)

							if True == tmp3939 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp3932 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5s_1exprs_6 := __e.Get(1)
									_ = Parseshen_4_5s_1exprs_6
									tmp3937 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

									if True == tmp3937 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp3934 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

										tmp3935 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

										tmp3936 := PrimCons(sym_i, tmp3935)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp3934, tmp3936)
										return

									}

								}, 1)

								tmp3938 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), Parseshen_4_5lcurly_6)

								__e.TailApply(tmp3932, tmp3938)
								return

							}

						}, 1)

						tmp3940 := Call(__e, PrimNS2Value(symshen_4_5lcurly_6), V2810)

						tmp3941 := Call(__e, tmp3930, tmp3940)

						__e.TailApply(tmp3776, tmp3941)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp3943 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5lrb_6 := __e.Get(1)
					_ = Parseshen_4_5lrb_6
					tmp3961 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5lrb_6)

					if True == tmp3961 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp3945 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5s_1exprs1_6 := __e.Get(1)
							_ = Parseshen_4_5s_1exprs1_6
							tmp3959 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs1_6)

							if True == tmp3959 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp3947 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5rrb_6 := __e.Get(1)
									_ = Parseshen_4_5rrb_6
									tmp3957 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rrb_6)

									if True == tmp3957 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp3949 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5s_1exprs2_6 := __e.Get(1)
											_ = Parseshen_4_5s_1exprs2_6
											tmp3955 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs2_6)

											if True == tmp3955 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp3951 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs2_6)

												tmp3952 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs1_6)

												tmp3953 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs2_6)

												tmp3954 := Call(__e, PrimNS2Value(symshen_4add_1sexpr), tmp3952, tmp3953)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp3951, tmp3954)
												return

											}

										}, 1)

										tmp3956 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs2_6), Parseshen_4_5rrb_6)

										__e.TailApply(tmp3949, tmp3956)
										return

									}

								}, 1)

								tmp3958 := Call(__e, PrimNS2Value(symshen_4_5rrb_6), Parseshen_4_5s_1exprs1_6)

								__e.TailApply(tmp3947, tmp3958)
								return

							}

						}, 1)

						tmp3960 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs1_6), Parseshen_4_5lrb_6)

						__e.TailApply(tmp3945, tmp3960)
						return

					}

				}, 1)

				tmp3962 := Call(__e, PrimNS2Value(symshen_4_5lrb_6), V2810)

				tmp3963 := Call(__e, tmp3943, tmp3962)

				__e.TailApply(tmp3774, tmp3963)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp3965 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5lsb_6 := __e.Get(1)
			_ = Parseshen_4_5lsb_6
			tmp3984 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5lsb_6)

			if True == tmp3984 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp3967 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5s_1exprs1_6 := __e.Get(1)
					_ = Parseshen_4_5s_1exprs1_6
					tmp3982 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs1_6)

					if True == tmp3982 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp3969 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5rsb_6 := __e.Get(1)
							_ = Parseshen_4_5rsb_6
							tmp3980 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5rsb_6)

							if True == tmp3980 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp3971 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5s_1exprs2_6 := __e.Get(1)
									_ = Parseshen_4_5s_1exprs2_6
									tmp3978 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs2_6)

									if True == tmp3978 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp3973 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs2_6)

										tmp3974 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs1_6)

										tmp3975 := Call(__e, PrimNS2Value(symshen_4cons_1form), tmp3974)

										tmp3976 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs2_6)

										tmp3977 := PrimCons(tmp3975, tmp3976)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp3973, tmp3977)
										return

									}

								}, 1)

								tmp3979 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs2_6), Parseshen_4_5rsb_6)

								__e.TailApply(tmp3971, tmp3979)
								return

							}

						}, 1)

						tmp3981 := Call(__e, PrimNS2Value(symshen_4_5rsb_6), Parseshen_4_5s_1exprs1_6)

						__e.TailApply(tmp3969, tmp3981)
						return

					}

				}, 1)

				tmp3983 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs1_6), Parseshen_4_5lsb_6)

				__e.TailApply(tmp3967, tmp3983)
				return

			}

		}, 1)

		tmp3985 := Call(__e, PrimNS2Value(symshen_4_5lsb_6), V2810)

		tmp3986 := Call(__e, tmp3965, tmp3985)

		__e.TailApply(tmp3772, tmp3986)
		return

	}, 1)

	tmp3987 := Call(__e, PrimNS2Value(symdef), symshen_4_5s_1exprs_6, tmp3771)

	_ = tmp3987

	tmp3988 := MakeNative(func(__e *ControlFlow) {
		V2811 := __e.Get(1)
		_ = V2811
		V2812 := __e.Get(2)
		_ = V2812
		tmp4006 := PrimIsPair(V2811)

		var ifres3993 Obj

		if True == tmp4006 {
			tmp4004 := PrimHead(V2811)

			tmp4005 := PrimEqual(sym_3, tmp4004)

			var ifres3995 Obj

			if True == tmp4005 {
				tmp4002 := PrimTail(V2811)

				tmp4003 := PrimIsPair(tmp4002)

				var ifres3997 Obj

				if True == tmp4003 {
					tmp3999 := PrimTail(V2811)

					tmp4000 := PrimTail(tmp3999)

					tmp4001 := PrimEqual(Nil, tmp4000)

					var ifres3998 Obj

					if True == tmp4001 {
						ifres3998 = True

					} else {
						ifres3998 = False

					}

					ifres3997 = ifres3998

				} else {
					ifres3997 = False

				}

				var ifres3996 Obj

				if True == ifres3997 {
					ifres3996 = True

				} else {
					ifres3996 = False

				}

				ifres3995 = ifres3996

			} else {
				ifres3995 = False

			}

			var ifres3994 Obj

			if True == ifres3995 {
				ifres3994 = True

			} else {
				ifres3994 = False

			}

			ifres3993 = ifres3994

		} else {
			ifres3993 = False

		}

		if True == ifres3993 {
			tmp3990 := PrimTail(V2811)

			tmp3991 := PrimHead(tmp3990)

			tmp3992 := Call(__e, PrimNS2Value(symexplode), tmp3991)

			__e.TailApply(PrimNS2Value(symappend), tmp3992, V2812)
			return

		} else {
			__e.Return(PrimCons(V2811, V2812))
			return
		}

	}, 2)

	tmp4007 := Call(__e, PrimNS2Value(symdef), symshen_4add_1sexpr, tmp3988)

	_ = tmp4007

	tmp4008 := MakeNative(func(__e *ControlFlow) {
		V2813 := __e.Get(1)
		_ = V2813
		tmp4009 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4011 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4011 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4018 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2813, MakeNumber(91))

		var ifres4012 Obj

		if True == tmp4018 {
			tmp4014 := MakeNative(func(__e *ControlFlow) {
				News2495 := __e.Get(1)
				_ = News2495
				tmp4015 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2495)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4015, symshen_4skip)
				return

			}, 1)

			tmp4016 := Call(__e, PrimNS2Value(symshen_4tls), V2813)

			tmp4017 := Call(__e, tmp4014, tmp4016)

			ifres4012 = tmp4017

		} else {
			tmp4013 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4012 = tmp4013

		}

		__e.TailApply(tmp4009, ifres4012)
		return

	}, 1)

	tmp4019 := Call(__e, PrimNS2Value(symdef), symshen_4_5lsb_6, tmp4008)

	_ = tmp4019

	tmp4020 := MakeNative(func(__e *ControlFlow) {
		V2814 := __e.Get(1)
		_ = V2814
		tmp4021 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4023 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4023 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4030 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2814, MakeNumber(93))

		var ifres4024 Obj

		if True == tmp4030 {
			tmp4026 := MakeNative(func(__e *ControlFlow) {
				News2497 := __e.Get(1)
				_ = News2497
				tmp4027 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2497)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4027, symshen_4skip)
				return

			}, 1)

			tmp4028 := Call(__e, PrimNS2Value(symshen_4tls), V2814)

			tmp4029 := Call(__e, tmp4026, tmp4028)

			ifres4024 = tmp4029

		} else {
			tmp4025 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4024 = tmp4025

		}

		__e.TailApply(tmp4021, ifres4024)
		return

	}, 1)

	tmp4031 := Call(__e, PrimNS2Value(symdef), symshen_4_5rsb_6, tmp4020)

	_ = tmp4031

	tmp4032 := MakeNative(func(__e *ControlFlow) {
		V2815 := __e.Get(1)
		_ = V2815
		tmp4033 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4035 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4035 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4036 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5s_1exprs_6 := __e.Get(1)
			_ = Parseshen_4_5s_1exprs_6
			tmp4040 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

			if True == tmp4040 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4038 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

				tmp4039 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4038, tmp4039)
				return

			}

		}, 1)

		tmp4041 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), V2815)

		tmp4042 := Call(__e, tmp4036, tmp4041)

		__e.TailApply(tmp4033, tmp4042)
		return

	}, 1)

	tmp4043 := Call(__e, PrimNS2Value(symdef), symshen_4_5s_1exprs1_6, tmp4032)

	_ = tmp4043

	tmp4044 := MakeNative(func(__e *ControlFlow) {
		V2816 := __e.Get(1)
		_ = V2816
		tmp4045 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4047 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4047 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4048 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5s_1exprs_6 := __e.Get(1)
			_ = Parseshen_4_5s_1exprs_6
			tmp4052 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5s_1exprs_6)

			if True == tmp4052 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4050 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5s_1exprs_6)

				tmp4051 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5s_1exprs_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4050, tmp4051)
				return

			}

		}, 1)

		tmp4053 := Call(__e, PrimNS2Value(symshen_4_5s_1exprs_6), V2816)

		tmp4054 := Call(__e, tmp4048, tmp4053)

		__e.TailApply(tmp4045, tmp4054)
		return

	}, 1)

	tmp4055 := Call(__e, PrimNS2Value(symdef), symshen_4_5s_1exprs2_6, tmp4044)

	_ = tmp4055

	tmp4056 := MakeNative(func(__e *ControlFlow) {
		V2818 := __e.Get(1)
		_ = V2818
		tmp4113 := PrimEqual(Nil, V2818)

		if True == tmp4113 {
			__e.Return(Nil)
			return
		} else {
			tmp4112 := PrimIsPair(V2818)

			var ifres4092 Obj

			if True == tmp4112 {
				tmp4110 := PrimTail(V2818)

				tmp4111 := PrimIsPair(tmp4110)

				var ifres4094 Obj

				if True == tmp4111 {
					tmp4107 := PrimTail(V2818)

					tmp4108 := PrimTail(tmp4107)

					tmp4109 := PrimIsPair(tmp4108)

					var ifres4096 Obj

					if True == tmp4109 {
						tmp4103 := PrimTail(V2818)

						tmp4104 := PrimTail(tmp4103)

						tmp4105 := PrimTail(tmp4104)

						tmp4106 := PrimEqual(Nil, tmp4105)

						var ifres4098 Obj

						if True == tmp4106 {
							tmp4100 := PrimTail(V2818)

							tmp4101 := PrimHead(tmp4100)

							tmp4102 := PrimEqual(tmp4101, symbar_b)

							var ifres4099 Obj

							if True == tmp4102 {
								ifres4099 = True

							} else {
								ifres4099 = False

							}

							ifres4098 = ifres4099

						} else {
							ifres4098 = False

						}

						var ifres4097 Obj

						if True == ifres4098 {
							ifres4097 = True

						} else {
							ifres4097 = False

						}

						ifres4096 = ifres4097

					} else {
						ifres4096 = False

					}

					var ifres4095 Obj

					if True == ifres4096 {
						ifres4095 = True

					} else {
						ifres4095 = False

					}

					ifres4094 = ifres4095

				} else {
					ifres4094 = False

				}

				var ifres4093 Obj

				if True == ifres4094 {
					ifres4093 = True

				} else {
					ifres4093 = False

				}

				ifres4092 = ifres4093

			} else {
				ifres4092 = False

			}

			if True == ifres4092 {
				tmp4088 := PrimHead(V2818)

				tmp4089 := PrimTail(V2818)

				tmp4090 := PrimTail(tmp4089)

				tmp4091 := PrimCons(tmp4088, tmp4090)

				__e.Return(PrimCons(symcons, tmp4091))
				return

			} else {
				tmp4087 := PrimIsPair(V2818)

				var ifres4067 Obj

				if True == tmp4087 {
					tmp4085 := PrimTail(V2818)

					tmp4086 := PrimIsPair(tmp4085)

					var ifres4069 Obj

					if True == tmp4086 {
						tmp4082 := PrimTail(V2818)

						tmp4083 := PrimTail(tmp4082)

						tmp4084 := PrimIsPair(tmp4083)

						var ifres4071 Obj

						if True == tmp4084 {
							tmp4078 := PrimTail(V2818)

							tmp4079 := PrimTail(tmp4078)

							tmp4080 := PrimTail(tmp4079)

							tmp4081 := PrimIsPair(tmp4080)

							var ifres4073 Obj

							if True == tmp4081 {
								tmp4075 := PrimTail(V2818)

								tmp4076 := PrimHead(tmp4075)

								tmp4077 := PrimEqual(tmp4076, symbar_b)

								var ifres4074 Obj

								if True == tmp4077 {
									ifres4074 = True

								} else {
									ifres4074 = False

								}

								ifres4073 = ifres4074

							} else {
								ifres4073 = False

							}

							var ifres4072 Obj

							if True == ifres4073 {
								ifres4072 = True

							} else {
								ifres4072 = False

							}

							ifres4071 = ifres4072

						} else {
							ifres4071 = False

						}

						var ifres4070 Obj

						if True == ifres4071 {
							ifres4070 = True

						} else {
							ifres4070 = False

						}

						ifres4069 = ifres4070

					} else {
						ifres4069 = False

					}

					var ifres4068 Obj

					if True == ifres4069 {
						ifres4068 = True

					} else {
						ifres4068 = False

					}

					ifres4067 = ifres4068

				} else {
					ifres4067 = False

				}

				if True == ifres4067 {
					__e.Return(PrimSimpleError(MakeString("misapplication of |\n")))
					return
				} else {
					tmp4066 := PrimIsPair(V2818)

					if True == tmp4066 {
						tmp4061 := PrimHead(V2818)

						tmp4062 := PrimTail(V2818)

						tmp4063 := Call(__e, PrimNS2Value(symshen_4cons_1form), tmp4062)

						tmp4064 := PrimCons(tmp4063, Nil)

						tmp4065 := PrimCons(tmp4061, tmp4064)

						__e.Return(PrimCons(symcons, tmp4065))
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4cons_1form)
						return
					}

				}

			}

		}

	}, 1)

	tmp4114 := Call(__e, PrimNS2Value(symdef), symshen_4cons_1form, tmp4056)

	_ = tmp4114

	tmp4115 := MakeNative(func(__e *ControlFlow) {
		V2819 := __e.Get(1)
		_ = V2819
		tmp4116 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4118 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4118 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4125 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2819, MakeNumber(40))

		var ifres4119 Obj

		if True == tmp4125 {
			tmp4121 := MakeNative(func(__e *ControlFlow) {
				News2501 := __e.Get(1)
				_ = News2501
				tmp4122 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2501)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4122, symshen_4skip)
				return

			}, 1)

			tmp4123 := Call(__e, PrimNS2Value(symshen_4tls), V2819)

			tmp4124 := Call(__e, tmp4121, tmp4123)

			ifres4119 = tmp4124

		} else {
			tmp4120 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4119 = tmp4120

		}

		__e.TailApply(tmp4116, ifres4119)
		return

	}, 1)

	tmp4126 := Call(__e, PrimNS2Value(symdef), symshen_4_5lrb_6, tmp4115)

	_ = tmp4126

	tmp4127 := MakeNative(func(__e *ControlFlow) {
		V2820 := __e.Get(1)
		_ = V2820
		tmp4128 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4130 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4130 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4137 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2820, MakeNumber(41))

		var ifres4131 Obj

		if True == tmp4137 {
			tmp4133 := MakeNative(func(__e *ControlFlow) {
				News2503 := __e.Get(1)
				_ = News2503
				tmp4134 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2503)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4134, symshen_4skip)
				return

			}, 1)

			tmp4135 := Call(__e, PrimNS2Value(symshen_4tls), V2820)

			tmp4136 := Call(__e, tmp4133, tmp4135)

			ifres4131 = tmp4136

		} else {
			tmp4132 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4131 = tmp4132

		}

		__e.TailApply(tmp4128, ifres4131)
		return

	}, 1)

	tmp4138 := Call(__e, PrimNS2Value(symdef), symshen_4_5rrb_6, tmp4127)

	_ = tmp4138

	tmp4139 := MakeNative(func(__e *ControlFlow) {
		V2821 := __e.Get(1)
		_ = V2821
		tmp4140 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4142 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4142 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4149 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2821, MakeNumber(123))

		var ifres4143 Obj

		if True == tmp4149 {
			tmp4145 := MakeNative(func(__e *ControlFlow) {
				News2505 := __e.Get(1)
				_ = News2505
				tmp4146 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2505)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4146, symshen_4skip)
				return

			}, 1)

			tmp4147 := Call(__e, PrimNS2Value(symshen_4tls), V2821)

			tmp4148 := Call(__e, tmp4145, tmp4147)

			ifres4143 = tmp4148

		} else {
			tmp4144 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4143 = tmp4144

		}

		__e.TailApply(tmp4140, ifres4143)
		return

	}, 1)

	tmp4150 := Call(__e, PrimNS2Value(symdef), symshen_4_5lcurly_6, tmp4139)

	_ = tmp4150

	tmp4151 := MakeNative(func(__e *ControlFlow) {
		V2822 := __e.Get(1)
		_ = V2822
		tmp4152 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4154 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4154 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4161 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2822, MakeNumber(125))

		var ifres4155 Obj

		if True == tmp4161 {
			tmp4157 := MakeNative(func(__e *ControlFlow) {
				News2507 := __e.Get(1)
				_ = News2507
				tmp4158 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2507)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4158, symshen_4skip)
				return

			}, 1)

			tmp4159 := Call(__e, PrimNS2Value(symshen_4tls), V2822)

			tmp4160 := Call(__e, tmp4157, tmp4159)

			ifres4155 = tmp4160

		} else {
			tmp4156 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4155 = tmp4156

		}

		__e.TailApply(tmp4152, ifres4155)
		return

	}, 1)

	tmp4162 := Call(__e, PrimNS2Value(symdef), symshen_4_5rcurly_6, tmp4151)

	_ = tmp4162

	tmp4163 := MakeNative(func(__e *ControlFlow) {
		V2823 := __e.Get(1)
		_ = V2823
		tmp4164 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4166 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4166 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4173 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2823, MakeNumber(124))

		var ifres4167 Obj

		if True == tmp4173 {
			tmp4169 := MakeNative(func(__e *ControlFlow) {
				News2509 := __e.Get(1)
				_ = News2509
				tmp4170 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2509)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4170, symshen_4skip)
				return

			}, 1)

			tmp4171 := Call(__e, PrimNS2Value(symshen_4tls), V2823)

			tmp4172 := Call(__e, tmp4169, tmp4171)

			ifres4167 = tmp4172

		} else {
			tmp4168 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4167 = tmp4168

		}

		__e.TailApply(tmp4164, ifres4167)
		return

	}, 1)

	tmp4174 := Call(__e, PrimNS2Value(symdef), symshen_4_5bar_6, tmp4163)

	_ = tmp4174

	tmp4175 := MakeNative(func(__e *ControlFlow) {
		V2824 := __e.Get(1)
		_ = V2824
		tmp4176 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4178 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4178 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4185 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2824, MakeNumber(59))

		var ifres4179 Obj

		if True == tmp4185 {
			tmp4181 := MakeNative(func(__e *ControlFlow) {
				News2511 := __e.Get(1)
				_ = News2511
				tmp4182 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2511)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4182, symshen_4skip)
				return

			}, 1)

			tmp4183 := Call(__e, PrimNS2Value(symshen_4tls), V2824)

			tmp4184 := Call(__e, tmp4181, tmp4183)

			ifres4179 = tmp4184

		} else {
			tmp4180 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4179 = tmp4180

		}

		__e.TailApply(tmp4176, ifres4179)
		return

	}, 1)

	tmp4186 := Call(__e, PrimNS2Value(symdef), symshen_4_5semicolon_6, tmp4175)

	_ = tmp4186

	tmp4187 := MakeNative(func(__e *ControlFlow) {
		V2825 := __e.Get(1)
		_ = V2825
		tmp4188 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4190 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4190 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4197 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2825, MakeNumber(58))

		var ifres4191 Obj

		if True == tmp4197 {
			tmp4193 := MakeNative(func(__e *ControlFlow) {
				News2513 := __e.Get(1)
				_ = News2513
				tmp4194 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2513)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4194, symshen_4skip)
				return

			}, 1)

			tmp4195 := Call(__e, PrimNS2Value(symshen_4tls), V2825)

			tmp4196 := Call(__e, tmp4193, tmp4195)

			ifres4191 = tmp4196

		} else {
			tmp4192 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4191 = tmp4192

		}

		__e.TailApply(tmp4188, ifres4191)
		return

	}, 1)

	tmp4198 := Call(__e, PrimNS2Value(symdef), symshen_4_5colon_6, tmp4187)

	_ = tmp4198

	tmp4199 := MakeNative(func(__e *ControlFlow) {
		V2826 := __e.Get(1)
		_ = V2826
		tmp4200 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4202 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4202 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4209 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2826, MakeNumber(44))

		var ifres4203 Obj

		if True == tmp4209 {
			tmp4205 := MakeNative(func(__e *ControlFlow) {
				News2515 := __e.Get(1)
				_ = News2515
				tmp4206 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2515)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4206, symshen_4skip)
				return

			}, 1)

			tmp4207 := Call(__e, PrimNS2Value(symshen_4tls), V2826)

			tmp4208 := Call(__e, tmp4205, tmp4207)

			ifres4203 = tmp4208

		} else {
			tmp4204 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4203 = tmp4204

		}

		__e.TailApply(tmp4200, ifres4203)
		return

	}, 1)

	tmp4210 := Call(__e, PrimNS2Value(symdef), symshen_4_5comma_6, tmp4199)

	_ = tmp4210

	tmp4211 := MakeNative(func(__e *ControlFlow) {
		V2827 := __e.Get(1)
		_ = V2827
		tmp4212 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4214 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4214 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4221 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2827, MakeNumber(61))

		var ifres4215 Obj

		if True == tmp4221 {
			tmp4217 := MakeNative(func(__e *ControlFlow) {
				News2517 := __e.Get(1)
				_ = News2517
				tmp4218 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2517)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4218, symshen_4skip)
				return

			}, 1)

			tmp4219 := Call(__e, PrimNS2Value(symshen_4tls), V2827)

			tmp4220 := Call(__e, tmp4217, tmp4219)

			ifres4215 = tmp4220

		} else {
			tmp4216 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4215 = tmp4216

		}

		__e.TailApply(tmp4212, ifres4215)
		return

	}, 1)

	tmp4222 := Call(__e, PrimNS2Value(symdef), symshen_4_5equal_6, tmp4211)

	_ = tmp4222

	tmp4223 := MakeNative(func(__e *ControlFlow) {
		V2828 := __e.Get(1)
		_ = V2828
		tmp4224 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4235 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4235 {
				tmp4226 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4228 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4228 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4229 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5multiline_6 := __e.Get(1)
					_ = Parseshen_4_5multiline_6
					tmp4232 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5multiline_6)

					if True == tmp4232 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4231 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5multiline_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4231, symshen_4skip)
						return

					}

				}, 1)

				tmp4233 := Call(__e, PrimNS2Value(symshen_4_5multiline_6), V2828)

				tmp4234 := Call(__e, tmp4229, tmp4233)

				__e.TailApply(tmp4226, tmp4234)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4236 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5singleline_6 := __e.Get(1)
			_ = Parseshen_4_5singleline_6
			tmp4239 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5singleline_6)

			if True == tmp4239 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4238 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5singleline_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4238, symshen_4skip)
				return

			}

		}, 1)

		tmp4240 := Call(__e, PrimNS2Value(symshen_4_5singleline_6), V2828)

		tmp4241 := Call(__e, tmp4236, tmp4240)

		__e.TailApply(tmp4224, tmp4241)
		return

	}, 1)

	tmp4242 := Call(__e, PrimNS2Value(symdef), symshen_4_5comment_6, tmp4223)

	_ = tmp4242

	tmp4243 := MakeNative(func(__e *ControlFlow) {
		V2829 := __e.Get(1)
		_ = V2829
		tmp4244 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4246 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4246 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4247 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5backslash_6 := __e.Get(1)
			_ = Parseshen_4_5backslash_6
			tmp4262 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

			if True == tmp4262 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4249 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5backslash_6 := __e.Get(1)
					_ = Parseshen_4_5backslash_6
					tmp4260 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

					if True == tmp4260 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4251 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5shortnatters_6 := __e.Get(1)
							_ = Parseshen_4_5shortnatters_6
							tmp4258 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5shortnatters_6)

							if True == tmp4258 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4253 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5returns_6 := __e.Get(1)
									_ = Parseshen_4_5returns_6
									tmp4256 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5returns_6)

									if True == tmp4256 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp4255 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5returns_6)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp4255, symshen_4skip)
										return

									}

								}, 1)

								tmp4257 := Call(__e, PrimNS2Value(symshen_4_5returns_6), Parseshen_4_5shortnatters_6)

								__e.TailApply(tmp4253, tmp4257)
								return

							}

						}, 1)

						tmp4259 := Call(__e, PrimNS2Value(symshen_4_5shortnatters_6), Parseshen_4_5backslash_6)

						__e.TailApply(tmp4251, tmp4259)
						return

					}

				}, 1)

				tmp4261 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), Parseshen_4_5backslash_6)

				__e.TailApply(tmp4249, tmp4261)
				return

			}

		}, 1)

		tmp4263 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), V2829)

		tmp4264 := Call(__e, tmp4247, tmp4263)

		__e.TailApply(tmp4244, tmp4264)
		return

	}, 1)

	tmp4265 := Call(__e, PrimNS2Value(symdef), symshen_4_5singleline_6, tmp4243)

	_ = tmp4265

	tmp4266 := MakeNative(func(__e *ControlFlow) {
		V2830 := __e.Get(1)
		_ = V2830
		tmp4267 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4269 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4269 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4276 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2830, MakeNumber(92))

		var ifres4270 Obj

		if True == tmp4276 {
			tmp4272 := MakeNative(func(__e *ControlFlow) {
				News2521 := __e.Get(1)
				_ = News2521
				tmp4273 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2521)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4273, symshen_4skip)
				return

			}, 1)

			tmp4274 := Call(__e, PrimNS2Value(symshen_4tls), V2830)

			tmp4275 := Call(__e, tmp4272, tmp4274)

			ifres4270 = tmp4275

		} else {
			tmp4271 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4270 = tmp4271

		}

		__e.TailApply(tmp4267, ifres4270)
		return

	}, 1)

	tmp4277 := Call(__e, PrimNS2Value(symdef), symshen_4_5backslash_6, tmp4266)

	_ = tmp4277

	tmp4278 := MakeNative(func(__e *ControlFlow) {
		V2831 := __e.Get(1)
		_ = V2831
		tmp4279 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4290 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4290 {
				tmp4281 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4283 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4283 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4284 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp4287 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp4287 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4286 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4286, symshen_4skip)
						return

					}

				}, 1)

				tmp4288 := Call(__e, PrimNS2Value(sym_5e_6), V2831)

				tmp4289 := Call(__e, tmp4284, tmp4288)

				__e.TailApply(tmp4281, tmp4289)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4291 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5shortnatter_6 := __e.Get(1)
			_ = Parseshen_4_5shortnatter_6
			tmp4298 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5shortnatter_6)

			if True == tmp4298 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4293 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5shortnatters_6 := __e.Get(1)
					_ = Parseshen_4_5shortnatters_6
					tmp4296 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5shortnatters_6)

					if True == tmp4296 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4295 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5shortnatters_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4295, symshen_4skip)
						return

					}

				}, 1)

				tmp4297 := Call(__e, PrimNS2Value(symshen_4_5shortnatters_6), Parseshen_4_5shortnatter_6)

				__e.TailApply(tmp4293, tmp4297)
				return

			}

		}, 1)

		tmp4299 := Call(__e, PrimNS2Value(symshen_4_5shortnatter_6), V2831)

		tmp4300 := Call(__e, tmp4291, tmp4299)

		__e.TailApply(tmp4279, tmp4300)
		return

	}, 1)

	tmp4301 := Call(__e, PrimNS2Value(symdef), symshen_4_5shortnatters_6, tmp4278)

	_ = tmp4301

	tmp4302 := MakeNative(func(__e *ControlFlow) {
		V2832 := __e.Get(1)
		_ = V2832
		tmp4303 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4305 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4305 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4317 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2832)

		var ifres4306 Obj

		if True == tmp4317 {
			tmp4308 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp4309 := MakeNative(func(__e *ControlFlow) {
					News2524 := __e.Get(1)
					_ = News2524
					tmp4312 := Call(__e, PrimNS2Value(symshen_4return_2), Byte)

					tmp4313 := PrimNot(tmp4312)

					if True == tmp4313 {
						tmp4311 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2524)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4311, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4314 := Call(__e, PrimNS2Value(symshen_4tls), V2832)

				__e.TailApply(tmp4309, tmp4314)
				return

			}, 1)

			tmp4315 := Call(__e, PrimNS2Value(symshen_4hds), V2832)

			tmp4316 := Call(__e, tmp4308, tmp4315)

			ifres4306 = tmp4316

		} else {
			tmp4307 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4306 = tmp4307

		}

		__e.TailApply(tmp4303, ifres4306)
		return

	}, 1)

	tmp4318 := Call(__e, PrimNS2Value(symdef), symshen_4_5shortnatter_6, tmp4302)

	_ = tmp4318

	tmp4319 := MakeNative(func(__e *ControlFlow) {
		V2833 := __e.Get(1)
		_ = V2833
		tmp4320 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4331 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4331 {
				tmp4322 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4324 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4324 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4325 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5return_6 := __e.Get(1)
					_ = Parseshen_4_5return_6
					tmp4328 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5return_6)

					if True == tmp4328 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4327 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5return_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4327, symshen_4skip)
						return

					}

				}, 1)

				tmp4329 := Call(__e, PrimNS2Value(symshen_4_5return_6), V2833)

				tmp4330 := Call(__e, tmp4325, tmp4329)

				__e.TailApply(tmp4322, tmp4330)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4332 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5return_6 := __e.Get(1)
			_ = Parseshen_4_5return_6
			tmp4339 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5return_6)

			if True == tmp4339 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4334 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5returns_6 := __e.Get(1)
					_ = Parseshen_4_5returns_6
					tmp4337 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5returns_6)

					if True == tmp4337 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4336 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5returns_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4336, symshen_4skip)
						return

					}

				}, 1)

				tmp4338 := Call(__e, PrimNS2Value(symshen_4_5returns_6), Parseshen_4_5return_6)

				__e.TailApply(tmp4334, tmp4338)
				return

			}

		}, 1)

		tmp4340 := Call(__e, PrimNS2Value(symshen_4_5return_6), V2833)

		tmp4341 := Call(__e, tmp4332, tmp4340)

		__e.TailApply(tmp4320, tmp4341)
		return

	}, 1)

	tmp4342 := Call(__e, PrimNS2Value(symdef), symshen_4_5returns_6, tmp4319)

	_ = tmp4342

	tmp4343 := MakeNative(func(__e *ControlFlow) {
		V2834 := __e.Get(1)
		_ = V2834
		tmp4344 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4346 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4346 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4357 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2834)

		var ifres4347 Obj

		if True == tmp4357 {
			tmp4349 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp4350 := MakeNative(func(__e *ControlFlow) {
					News2527 := __e.Get(1)
					_ = News2527
					tmp4353 := Call(__e, PrimNS2Value(symshen_4return_2), Byte)

					if True == tmp4353 {
						tmp4352 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2527)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4352, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4354 := Call(__e, PrimNS2Value(symshen_4tls), V2834)

				__e.TailApply(tmp4350, tmp4354)
				return

			}, 1)

			tmp4355 := Call(__e, PrimNS2Value(symshen_4hds), V2834)

			tmp4356 := Call(__e, tmp4349, tmp4355)

			ifres4347 = tmp4356

		} else {
			tmp4348 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4347 = tmp4348

		}

		__e.TailApply(tmp4344, ifres4347)
		return

	}, 1)

	tmp4358 := Call(__e, PrimNS2Value(symdef), symshen_4_5return_6, tmp4343)

	_ = tmp4358

	tmp4359 := MakeNative(func(__e *ControlFlow) {
		V2835 := __e.Get(1)
		_ = V2835
		tmp4360 := PrimCons(MakeNumber(13), Nil)

		tmp4361 := PrimCons(MakeNumber(10), tmp4360)

		tmp4362 := PrimCons(MakeNumber(9), tmp4361)

		__e.TailApply(PrimNS2Value(symelement_2), V2835, tmp4362)
		return

	}, 1)

	tmp4363 := Call(__e, PrimNS2Value(symdef), symshen_4return_2, tmp4359)

	_ = tmp4363

	tmp4364 := MakeNative(func(__e *ControlFlow) {
		V2836 := __e.Get(1)
		_ = V2836
		tmp4365 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4367 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4367 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4368 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5backslash_6 := __e.Get(1)
			_ = Parseshen_4_5backslash_6
			tmp4379 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

			if True == tmp4379 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4370 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5times_6 := __e.Get(1)
					_ = Parseshen_4_5times_6
					tmp4377 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5times_6)

					if True == tmp4377 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4372 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5longnatter_6 := __e.Get(1)
							_ = Parseshen_4_5longnatter_6
							tmp4375 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5longnatter_6)

							if True == tmp4375 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4374 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5longnatter_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4374, symshen_4skip)
								return

							}

						}, 1)

						tmp4376 := Call(__e, PrimNS2Value(symshen_4_5longnatter_6), Parseshen_4_5times_6)

						__e.TailApply(tmp4372, tmp4376)
						return

					}

				}, 1)

				tmp4378 := Call(__e, PrimNS2Value(symshen_4_5times_6), Parseshen_4_5backslash_6)

				__e.TailApply(tmp4370, tmp4378)
				return

			}

		}, 1)

		tmp4380 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), V2836)

		tmp4381 := Call(__e, tmp4368, tmp4380)

		__e.TailApply(tmp4365, tmp4381)
		return

	}, 1)

	tmp4382 := Call(__e, PrimNS2Value(symdef), symshen_4_5multiline_6, tmp4364)

	_ = tmp4382

	tmp4383 := MakeNative(func(__e *ControlFlow) {
		V2837 := __e.Get(1)
		_ = V2837
		tmp4384 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4386 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4386 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4393 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2837, MakeNumber(42))

		var ifres4387 Obj

		if True == tmp4393 {
			tmp4389 := MakeNative(func(__e *ControlFlow) {
				News2530 := __e.Get(1)
				_ = News2530
				tmp4390 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2530)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4390, symshen_4skip)
				return

			}, 1)

			tmp4391 := Call(__e, PrimNS2Value(symshen_4tls), V2837)

			tmp4392 := Call(__e, tmp4389, tmp4391)

			ifres4387 = tmp4392

		} else {
			tmp4388 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4387 = tmp4388

		}

		__e.TailApply(tmp4384, ifres4387)
		return

	}, 1)

	tmp4394 := Call(__e, PrimNS2Value(symdef), symshen_4_5times_6, tmp4383)

	_ = tmp4394

	tmp4395 := MakeNative(func(__e *ControlFlow) {
		V2838 := __e.Get(1)
		_ = V2838
		tmp4396 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4425 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4425 {
				tmp4398 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4414 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4414 {
						tmp4400 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp4402 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp4402 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp4413 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2838)

						var ifres4403 Obj

						if True == tmp4413 {
							tmp4405 := MakeNative(func(__e *ControlFlow) {
								News2532 := __e.Get(1)
								_ = News2532
								tmp4406 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5longnatter_6 := __e.Get(1)
									_ = Parseshen_4_5longnatter_6
									tmp4409 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5longnatter_6)

									if True == tmp4409 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp4408 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5longnatter_6)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp4408, symshen_4skip)
										return

									}

								}, 1)

								tmp4410 := Call(__e, PrimNS2Value(symshen_4_5longnatter_6), News2532)

								__e.TailApply(tmp4406, tmp4410)
								return

							}, 1)

							tmp4411 := Call(__e, PrimNS2Value(symshen_4tls), V2838)

							tmp4412 := Call(__e, tmp4405, tmp4411)

							ifres4403 = tmp4412

						} else {
							tmp4404 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

							ifres4403 = tmp4404

						}

						__e.TailApply(tmp4400, ifres4403)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4415 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5times_6 := __e.Get(1)
					_ = Parseshen_4_5times_6
					tmp4422 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5times_6)

					if True == tmp4422 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4417 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5backslash_6 := __e.Get(1)
							_ = Parseshen_4_5backslash_6
							tmp4420 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5backslash_6)

							if True == tmp4420 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4419 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5backslash_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4419, symshen_4skip)
								return

							}

						}, 1)

						tmp4421 := Call(__e, PrimNS2Value(symshen_4_5backslash_6), Parseshen_4_5times_6)

						__e.TailApply(tmp4417, tmp4421)
						return

					}

				}, 1)

				tmp4423 := Call(__e, PrimNS2Value(symshen_4_5times_6), V2838)

				tmp4424 := Call(__e, tmp4415, tmp4423)

				__e.TailApply(tmp4398, tmp4424)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4426 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5comment_6 := __e.Get(1)
			_ = Parseshen_4_5comment_6
			tmp4433 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5comment_6)

			if True == tmp4433 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4428 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5longnatter_6 := __e.Get(1)
					_ = Parseshen_4_5longnatter_6
					tmp4431 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5longnatter_6)

					if True == tmp4431 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4430 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5longnatter_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4430, symshen_4skip)
						return

					}

				}, 1)

				tmp4432 := Call(__e, PrimNS2Value(symshen_4_5longnatter_6), Parseshen_4_5comment_6)

				__e.TailApply(tmp4428, tmp4432)
				return

			}

		}, 1)

		tmp4434 := Call(__e, PrimNS2Value(symshen_4_5comment_6), V2838)

		tmp4435 := Call(__e, tmp4426, tmp4434)

		__e.TailApply(tmp4396, tmp4435)
		return

	}, 1)

	tmp4436 := Call(__e, PrimNS2Value(symdef), symshen_4_5longnatter_6, tmp4395)

	_ = tmp4436

	tmp4437 := MakeNative(func(__e *ControlFlow) {
		V2839 := __e.Get(1)
		_ = V2839
		tmp4438 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4466 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4466 {
				tmp4440 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4458 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4458 {
						tmp4442 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp4444 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp4444 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp4445 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5sym_6 := __e.Get(1)
							_ = Parseshen_4_5sym_6
							tmp4455 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5sym_6)

							if True == tmp4455 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4447 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5sym_6)

								tmp4453 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5sym_6)

								tmp4454 := PrimEqual(tmp4453, MakeString("<>"))

								var ifres4448 Obj

								if True == tmp4454 {
									tmp4451 := PrimCons(MakeNumber(0), Nil)

									tmp4452 := PrimCons(symvector, tmp4451)

									ifres4448 = tmp4452

								} else {
									tmp4449 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5sym_6)

									tmp4450 := PrimIntern(tmp4449)

									ifres4448 = tmp4450

								}

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4447, ifres4448)
								return

							}

						}, 1)

						tmp4456 := Call(__e, PrimNS2Value(symshen_4_5sym_6), V2839)

						tmp4457 := Call(__e, tmp4445, tmp4456)

						__e.TailApply(tmp4442, tmp4457)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4459 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5number_6 := __e.Get(1)
					_ = Parseshen_4_5number_6
					tmp4463 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5number_6)

					if True == tmp4463 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4461 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5number_6)

						tmp4462 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5number_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4461, tmp4462)
						return

					}

				}, 1)

				tmp4464 := Call(__e, PrimNS2Value(symshen_4_5number_6), V2839)

				tmp4465 := Call(__e, tmp4459, tmp4464)

				__e.TailApply(tmp4440, tmp4465)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4467 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5str_6 := __e.Get(1)
			_ = Parseshen_4_5str_6
			tmp4471 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5str_6)

			if True == tmp4471 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4469 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5str_6)

				tmp4470 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5str_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4469, tmp4470)
				return

			}

		}, 1)

		tmp4472 := Call(__e, PrimNS2Value(symshen_4_5str_6), V2839)

		tmp4473 := Call(__e, tmp4467, tmp4472)

		__e.TailApply(tmp4438, tmp4473)
		return

	}, 1)

	tmp4474 := Call(__e, PrimNS2Value(symdef), symshen_4_5atom_6, tmp4437)

	_ = tmp4474

	tmp4475 := MakeNative(func(__e *ControlFlow) {
		V2840 := __e.Get(1)
		_ = V2840
		tmp4476 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4478 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4478 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4479 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5alpha_6 := __e.Get(1)
			_ = Parseshen_4_5alpha_6
			tmp4489 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5alpha_6)

			if True == tmp4489 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4481 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5alphanums_6 := __e.Get(1)
					_ = Parseshen_4_5alphanums_6
					tmp4487 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5alphanums_6)

					if True == tmp4487 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4483 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5alphanums_6)

						tmp4484 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5alpha_6)

						tmp4485 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5alphanums_6)

						tmp4486 := PrimStringConcat(tmp4484, tmp4485)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4483, tmp4486)
						return

					}

				}, 1)

				tmp4488 := Call(__e, PrimNS2Value(symshen_4_5alphanums_6), Parseshen_4_5alpha_6)

				__e.TailApply(tmp4481, tmp4488)
				return

			}

		}, 1)

		tmp4490 := Call(__e, PrimNS2Value(symshen_4_5alpha_6), V2840)

		tmp4491 := Call(__e, tmp4479, tmp4490)

		__e.TailApply(tmp4476, tmp4491)
		return

	}, 1)

	tmp4492 := Call(__e, PrimNS2Value(symdef), symshen_4_5sym_6, tmp4475)

	_ = tmp4492

	tmp4493 := MakeNative(func(__e *ControlFlow) {
		V2841 := __e.Get(1)
		_ = V2841
		tmp4494 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4496 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4496 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4508 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2841)

		var ifres4497 Obj

		if True == tmp4508 {
			tmp4499 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp4500 := MakeNative(func(__e *ControlFlow) {
					News2536 := __e.Get(1)
					_ = News2536
					tmp4504 := Call(__e, PrimNS2Value(symshen_4alpha_2), Byte)

					if True == tmp4504 {
						tmp4502 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2536)

						tmp4503 := PrimNumberToString(Byte)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4502, tmp4503)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4505 := Call(__e, PrimNS2Value(symshen_4tls), V2841)

				__e.TailApply(tmp4500, tmp4505)
				return

			}, 1)

			tmp4506 := Call(__e, PrimNS2Value(symshen_4hds), V2841)

			tmp4507 := Call(__e, tmp4499, tmp4506)

			ifres4497 = tmp4507

		} else {
			tmp4498 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4497 = tmp4498

		}

		__e.TailApply(tmp4494, ifres4497)
		return

	}, 1)

	tmp4509 := Call(__e, PrimNS2Value(symdef), symshen_4_5alpha_6, tmp4493)

	_ = tmp4509

	tmp4510 := MakeNative(func(__e *ControlFlow) {
		V2842 := __e.Get(1)
		_ = V2842
		tmp4517 := Call(__e, PrimNS2Value(symshen_4lowercase_2), V2842)

		if True == tmp4517 {
			__e.Return(True)
			return
		} else {
			tmp4516 := Call(__e, PrimNS2Value(symshen_4uppercase_2), V2842)

			var ifres4513 Obj

			if True == tmp4516 {
				ifres4513 = True

			} else {
				tmp4515 := Call(__e, PrimNS2Value(symshen_4misc_2), V2842)

				var ifres4514 Obj

				if True == tmp4515 {
					ifres4514 = True

				} else {
					ifres4514 = False

				}

				ifres4513 = ifres4514

			}

			if True == ifres4513 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp4518 := Call(__e, PrimNS2Value(symdef), symshen_4alpha_2, tmp4510)

	_ = tmp4518

	tmp4519 := MakeNative(func(__e *ControlFlow) {
		V2843 := __e.Get(1)
		_ = V2843
		tmp4523 := PrimGreatEqual(V2843, MakeNumber(97))

		if True == tmp4523 {
			tmp4522 := PrimLessEqual(V2843, MakeNumber(122))

			if True == tmp4522 {
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

	tmp4524 := Call(__e, PrimNS2Value(symdef), symshen_4lowercase_2, tmp4519)

	_ = tmp4524

	tmp4525 := MakeNative(func(__e *ControlFlow) {
		V2844 := __e.Get(1)
		_ = V2844
		tmp4529 := PrimGreatEqual(V2844, MakeNumber(65))

		if True == tmp4529 {
			tmp4528 := PrimLessEqual(V2844, MakeNumber(90))

			if True == tmp4528 {
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

	tmp4530 := Call(__e, PrimNS2Value(symdef), symshen_4uppercase_2, tmp4525)

	_ = tmp4530

	tmp4531 := MakeNative(func(__e *ControlFlow) {
		V2845 := __e.Get(1)
		_ = V2845
		tmp4532 := PrimCons(MakeNumber(96), Nil)

		tmp4533 := PrimCons(MakeNumber(35), tmp4532)

		tmp4534 := PrimCons(MakeNumber(39), tmp4533)

		tmp4535 := PrimCons(MakeNumber(37), tmp4534)

		tmp4536 := PrimCons(MakeNumber(38), tmp4535)

		tmp4537 := PrimCons(MakeNumber(60), tmp4536)

		tmp4538 := PrimCons(MakeNumber(62), tmp4537)

		tmp4539 := PrimCons(MakeNumber(46), tmp4538)

		tmp4540 := PrimCons(MakeNumber(126), tmp4539)

		tmp4541 := PrimCons(MakeNumber(64), tmp4540)

		tmp4542 := PrimCons(MakeNumber(33), tmp4541)

		tmp4543 := PrimCons(MakeNumber(36), tmp4542)

		tmp4544 := PrimCons(MakeNumber(63), tmp4543)

		tmp4545 := PrimCons(MakeNumber(95), tmp4544)

		tmp4546 := PrimCons(MakeNumber(43), tmp4545)

		tmp4547 := PrimCons(MakeNumber(47), tmp4546)

		tmp4548 := PrimCons(MakeNumber(42), tmp4547)

		tmp4549 := PrimCons(MakeNumber(45), tmp4548)

		tmp4550 := PrimCons(MakeNumber(61), tmp4549)

		__e.TailApply(PrimNS2Value(symelement_2), V2845, tmp4550)
		return

	}, 1)

	tmp4551 := Call(__e, PrimNS2Value(symdef), symshen_4misc_2, tmp4531)

	_ = tmp4551

	tmp4552 := MakeNative(func(__e *ControlFlow) {
		V2846 := __e.Get(1)
		_ = V2846
		tmp4553 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4564 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4564 {
				tmp4555 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4557 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4557 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4558 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp4561 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp4561 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4560 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4560, MakeString(""))
						return

					}

				}, 1)

				tmp4562 := Call(__e, PrimNS2Value(sym_5e_6), V2846)

				tmp4563 := Call(__e, tmp4558, tmp4562)

				__e.TailApply(tmp4555, tmp4563)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4565 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5alphanum_6 := __e.Get(1)
			_ = Parseshen_4_5alphanum_6
			tmp4575 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5alphanum_6)

			if True == tmp4575 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4567 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5alphanums_6 := __e.Get(1)
					_ = Parseshen_4_5alphanums_6
					tmp4573 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5alphanums_6)

					if True == tmp4573 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4569 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5alphanums_6)

						tmp4570 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5alphanum_6)

						tmp4571 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5alphanums_6)

						tmp4572 := PrimStringConcat(tmp4570, tmp4571)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4569, tmp4572)
						return

					}

				}, 1)

				tmp4574 := Call(__e, PrimNS2Value(symshen_4_5alphanums_6), Parseshen_4_5alphanum_6)

				__e.TailApply(tmp4567, tmp4574)
				return

			}

		}, 1)

		tmp4576 := Call(__e, PrimNS2Value(symshen_4_5alphanum_6), V2846)

		tmp4577 := Call(__e, tmp4565, tmp4576)

		__e.TailApply(tmp4553, tmp4577)
		return

	}, 1)

	tmp4578 := Call(__e, PrimNS2Value(symdef), symshen_4_5alphanums_6, tmp4552)

	_ = tmp4578

	tmp4579 := MakeNative(func(__e *ControlFlow) {
		V2847 := __e.Get(1)
		_ = V2847
		tmp4580 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4592 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4592 {
				tmp4582 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4584 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4584 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4585 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5numeral_6 := __e.Get(1)
					_ = Parseshen_4_5numeral_6
					tmp4589 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5numeral_6)

					if True == tmp4589 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4587 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5numeral_6)

						tmp4588 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5numeral_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4587, tmp4588)
						return

					}

				}, 1)

				tmp4590 := Call(__e, PrimNS2Value(symshen_4_5numeral_6), V2847)

				tmp4591 := Call(__e, tmp4585, tmp4590)

				__e.TailApply(tmp4582, tmp4591)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4593 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5alpha_6 := __e.Get(1)
			_ = Parseshen_4_5alpha_6
			tmp4597 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5alpha_6)

			if True == tmp4597 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4595 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5alpha_6)

				tmp4596 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5alpha_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4595, tmp4596)
				return

			}

		}, 1)

		tmp4598 := Call(__e, PrimNS2Value(symshen_4_5alpha_6), V2847)

		tmp4599 := Call(__e, tmp4593, tmp4598)

		__e.TailApply(tmp4580, tmp4599)
		return

	}, 1)

	tmp4600 := Call(__e, PrimNS2Value(symdef), symshen_4_5alphanum_6, tmp4579)

	_ = tmp4600

	tmp4601 := MakeNative(func(__e *ControlFlow) {
		V2848 := __e.Get(1)
		_ = V2848
		tmp4602 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4604 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4604 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4616 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2848)

		var ifres4605 Obj

		if True == tmp4616 {
			tmp4607 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp4608 := MakeNative(func(__e *ControlFlow) {
					News2540 := __e.Get(1)
					_ = News2540
					tmp4612 := Call(__e, PrimNS2Value(symshen_4digit_2), Byte)

					if True == tmp4612 {
						tmp4610 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2540)

						tmp4611 := PrimNumberToString(Byte)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4610, tmp4611)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4613 := Call(__e, PrimNS2Value(symshen_4tls), V2848)

				__e.TailApply(tmp4608, tmp4613)
				return

			}, 1)

			tmp4614 := Call(__e, PrimNS2Value(symshen_4hds), V2848)

			tmp4615 := Call(__e, tmp4607, tmp4614)

			ifres4605 = tmp4615

		} else {
			tmp4606 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4605 = tmp4606

		}

		__e.TailApply(tmp4602, ifres4605)
		return

	}, 1)

	tmp4617 := Call(__e, PrimNS2Value(symdef), symshen_4_5numeral_6, tmp4601)

	_ = tmp4617

	tmp4618 := MakeNative(func(__e *ControlFlow) {
		V2849 := __e.Get(1)
		_ = V2849
		tmp4622 := PrimGreatEqual(V2849, MakeNumber(48))

		if True == tmp4622 {
			tmp4621 := PrimLessEqual(V2849, MakeNumber(57))

			if True == tmp4621 {
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

	tmp4623 := Call(__e, PrimNS2Value(symdef), symshen_4digit_2, tmp4618)

	_ = tmp4623

	tmp4624 := MakeNative(func(__e *ControlFlow) {
		V2850 := __e.Get(1)
		_ = V2850
		tmp4625 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4627 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4627 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4628 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5dbq_6 := __e.Get(1)
			_ = Parseshen_4_5dbq_6
			tmp4640 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5dbq_6)

			if True == tmp4640 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4630 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5strcontents_6 := __e.Get(1)
					_ = Parseshen_4_5strcontents_6
					tmp4638 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5strcontents_6)

					if True == tmp4638 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4632 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5dbq_6 := __e.Get(1)
							_ = Parseshen_4_5dbq_6
							tmp4636 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5dbq_6)

							if True == tmp4636 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4634 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5dbq_6)

								tmp4635 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5strcontents_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4634, tmp4635)
								return

							}

						}, 1)

						tmp4637 := Call(__e, PrimNS2Value(symshen_4_5dbq_6), Parseshen_4_5strcontents_6)

						__e.TailApply(tmp4632, tmp4637)
						return

					}

				}, 1)

				tmp4639 := Call(__e, PrimNS2Value(symshen_4_5strcontents_6), Parseshen_4_5dbq_6)

				__e.TailApply(tmp4630, tmp4639)
				return

			}

		}, 1)

		tmp4641 := Call(__e, PrimNS2Value(symshen_4_5dbq_6), V2850)

		tmp4642 := Call(__e, tmp4628, tmp4641)

		__e.TailApply(tmp4625, tmp4642)
		return

	}, 1)

	tmp4643 := Call(__e, PrimNS2Value(symdef), symshen_4_5str_6, tmp4624)

	_ = tmp4643

	tmp4644 := MakeNative(func(__e *ControlFlow) {
		V2851 := __e.Get(1)
		_ = V2851
		tmp4645 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4647 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4647 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4654 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2851, MakeNumber(34))

		var ifres4648 Obj

		if True == tmp4654 {
			tmp4650 := MakeNative(func(__e *ControlFlow) {
				News2543 := __e.Get(1)
				_ = News2543
				tmp4651 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2543)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4651, symshen_4skip)
				return

			}, 1)

			tmp4652 := Call(__e, PrimNS2Value(symshen_4tls), V2851)

			tmp4653 := Call(__e, tmp4650, tmp4652)

			ifres4648 = tmp4653

		} else {
			tmp4649 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4648 = tmp4649

		}

		__e.TailApply(tmp4645, ifres4648)
		return

	}, 1)

	tmp4655 := Call(__e, PrimNS2Value(symdef), symshen_4_5dbq_6, tmp4644)

	_ = tmp4655

	tmp4656 := MakeNative(func(__e *ControlFlow) {
		V2852 := __e.Get(1)
		_ = V2852
		tmp4657 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4668 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4668 {
				tmp4659 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4661 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4661 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4662 := MakeNative(func(__e *ControlFlow) {
					Parse_5e_6 := __e.Get(1)
					_ = Parse_5e_6
					tmp4665 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parse_5e_6)

					if True == tmp4665 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4664 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parse_5e_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4664, MakeString(""))
						return

					}

				}, 1)

				tmp4666 := Call(__e, PrimNS2Value(sym_5e_6), V2852)

				tmp4667 := Call(__e, tmp4662, tmp4666)

				__e.TailApply(tmp4659, tmp4667)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4669 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5strc_6 := __e.Get(1)
			_ = Parseshen_4_5strc_6
			tmp4679 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5strc_6)

			if True == tmp4679 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4671 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5strcontents_6 := __e.Get(1)
					_ = Parseshen_4_5strcontents_6
					tmp4677 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5strcontents_6)

					if True == tmp4677 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4673 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5strcontents_6)

						tmp4674 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5strc_6)

						tmp4675 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5strcontents_6)

						tmp4676 := PrimStringConcat(tmp4674, tmp4675)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4673, tmp4676)
						return

					}

				}, 1)

				tmp4678 := Call(__e, PrimNS2Value(symshen_4_5strcontents_6), Parseshen_4_5strc_6)

				__e.TailApply(tmp4671, tmp4678)
				return

			}

		}, 1)

		tmp4680 := Call(__e, PrimNS2Value(symshen_4_5strc_6), V2852)

		tmp4681 := Call(__e, tmp4669, tmp4680)

		__e.TailApply(tmp4657, tmp4681)
		return

	}, 1)

	tmp4682 := Call(__e, PrimNS2Value(symdef), symshen_4_5strcontents_6, tmp4656)

	_ = tmp4682

	tmp4683 := MakeNative(func(__e *ControlFlow) {
		V2853 := __e.Get(1)
		_ = V2853
		tmp4684 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4696 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4696 {
				tmp4686 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4688 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4688 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4689 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5notdbq_6 := __e.Get(1)
					_ = Parseshen_4_5notdbq_6
					tmp4693 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5notdbq_6)

					if True == tmp4693 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4691 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5notdbq_6)

						tmp4692 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5notdbq_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4691, tmp4692)
						return

					}

				}, 1)

				tmp4694 := Call(__e, PrimNS2Value(symshen_4_5notdbq_6), V2853)

				tmp4695 := Call(__e, tmp4689, tmp4694)

				__e.TailApply(tmp4686, tmp4695)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4697 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5control_6 := __e.Get(1)
			_ = Parseshen_4_5control_6
			tmp4701 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5control_6)

			if True == tmp4701 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4699 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5control_6)

				tmp4700 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5control_6)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4699, tmp4700)
				return

			}

		}, 1)

		tmp4702 := Call(__e, PrimNS2Value(symshen_4_5control_6), V2853)

		tmp4703 := Call(__e, tmp4697, tmp4702)

		__e.TailApply(tmp4684, tmp4703)
		return

	}, 1)

	tmp4704 := Call(__e, PrimNS2Value(symdef), symshen_4_5strc_6, tmp4683)

	_ = tmp4704

	tmp4705 := MakeNative(func(__e *ControlFlow) {
		V2854 := __e.Get(1)
		_ = V2854
		tmp4706 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4708 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4708 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4709 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5lowC_6 := __e.Get(1)
			_ = Parseshen_4_5lowC_6
			tmp4726 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5lowC_6)

			if True == tmp4726 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4711 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5hash_6 := __e.Get(1)
					_ = Parseshen_4_5hash_6
					tmp4724 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5hash_6)

					if True == tmp4724 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4713 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5integer_6 := __e.Get(1)
							_ = Parseshen_4_5integer_6
							tmp4722 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

							if True == tmp4722 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4715 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5semicolon_6 := __e.Get(1)
									_ = Parseshen_4_5semicolon_6
									tmp4720 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5semicolon_6)

									if True == tmp4720 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp4717 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5semicolon_6)

										tmp4718 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5integer_6)

										tmp4719 := PrimNumberToString(tmp4718)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp4717, tmp4719)
										return

									}

								}, 1)

								tmp4721 := Call(__e, PrimNS2Value(symshen_4_5semicolon_6), Parseshen_4_5integer_6)

								__e.TailApply(tmp4715, tmp4721)
								return

							}

						}, 1)

						tmp4723 := Call(__e, PrimNS2Value(symshen_4_5integer_6), Parseshen_4_5hash_6)

						__e.TailApply(tmp4713, tmp4723)
						return

					}

				}, 1)

				tmp4725 := Call(__e, PrimNS2Value(symshen_4_5hash_6), Parseshen_4_5lowC_6)

				__e.TailApply(tmp4711, tmp4725)
				return

			}

		}, 1)

		tmp4727 := Call(__e, PrimNS2Value(symshen_4_5lowC_6), V2854)

		tmp4728 := Call(__e, tmp4709, tmp4727)

		__e.TailApply(tmp4706, tmp4728)
		return

	}, 1)

	tmp4729 := Call(__e, PrimNS2Value(symdef), symshen_4_5control_6, tmp4705)

	_ = tmp4729

	tmp4730 := MakeNative(func(__e *ControlFlow) {
		V2855 := __e.Get(1)
		_ = V2855
		tmp4731 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4733 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4733 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4746 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2855)

		var ifres4734 Obj

		if True == tmp4746 {
			tmp4736 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp4737 := MakeNative(func(__e *ControlFlow) {
					News2548 := __e.Get(1)
					_ = News2548
					tmp4741 := PrimEqual(Byte, MakeNumber(34))

					tmp4742 := PrimNot(tmp4741)

					if True == tmp4742 {
						tmp4739 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2548)

						tmp4740 := PrimNumberToString(Byte)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4739, tmp4740)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4743 := Call(__e, PrimNS2Value(symshen_4tls), V2855)

				__e.TailApply(tmp4737, tmp4743)
				return

			}, 1)

			tmp4744 := Call(__e, PrimNS2Value(symshen_4hds), V2855)

			tmp4745 := Call(__e, tmp4736, tmp4744)

			ifres4734 = tmp4745

		} else {
			tmp4735 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4734 = tmp4735

		}

		__e.TailApply(tmp4731, ifres4734)
		return

	}, 1)

	tmp4747 := Call(__e, PrimNS2Value(symdef), symshen_4_5notdbq_6, tmp4730)

	_ = tmp4747

	tmp4748 := MakeNative(func(__e *ControlFlow) {
		V2856 := __e.Get(1)
		_ = V2856
		tmp4749 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4751 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4751 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4758 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2856, MakeNumber(99))

		var ifres4752 Obj

		if True == tmp4758 {
			tmp4754 := MakeNative(func(__e *ControlFlow) {
				News2550 := __e.Get(1)
				_ = News2550
				tmp4755 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2550)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4755, symshen_4skip)
				return

			}, 1)

			tmp4756 := Call(__e, PrimNS2Value(symshen_4tls), V2856)

			tmp4757 := Call(__e, tmp4754, tmp4756)

			ifres4752 = tmp4757

		} else {
			tmp4753 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4752 = tmp4753

		}

		__e.TailApply(tmp4749, ifres4752)
		return

	}, 1)

	tmp4759 := Call(__e, PrimNS2Value(symdef), symshen_4_5lowC_6, tmp4748)

	_ = tmp4759

	tmp4760 := MakeNative(func(__e *ControlFlow) {
		V2857 := __e.Get(1)
		_ = V2857
		tmp4761 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4763 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4763 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4770 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2857, MakeNumber(35))

		var ifres4764 Obj

		if True == tmp4770 {
			tmp4766 := MakeNative(func(__e *ControlFlow) {
				News2552 := __e.Get(1)
				_ = News2552
				tmp4767 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2552)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4767, symshen_4skip)
				return

			}, 1)

			tmp4768 := Call(__e, PrimNS2Value(symshen_4tls), V2857)

			tmp4769 := Call(__e, tmp4766, tmp4768)

			ifres4764 = tmp4769

		} else {
			tmp4765 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4764 = tmp4765

		}

		__e.TailApply(tmp4761, ifres4764)
		return

	}, 1)

	tmp4771 := Call(__e, PrimNS2Value(symdef), symshen_4_5hash_6, tmp4760)

	_ = tmp4771

	tmp4772 := MakeNative(func(__e *ControlFlow) {
		V2858 := __e.Get(1)
		_ = V2858
		tmp4773 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4819 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4819 {
				tmp4775 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4807 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4807 {
						tmp4777 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp4799 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp4799 {
								tmp4779 := MakeNative(func(__e *ControlFlow) {
									Result := __e.Get(1)
									_ = Result
									tmp4791 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

									if True == tmp4791 {
										tmp4781 := MakeNative(func(__e *ControlFlow) {
											Result := __e.Get(1)
											_ = Result
											tmp4783 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

											if True == tmp4783 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												__e.Return(Result)
												return
											}

										}, 1)

										tmp4784 := MakeNative(func(__e *ControlFlow) {
											Parseshen_4_5integer_6 := __e.Get(1)
											_ = Parseshen_4_5integer_6
											tmp4788 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

											if True == tmp4788 {
												__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
												return
											} else {
												tmp4786 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5integer_6)

												tmp4787 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5integer_6)

												__e.TailApply(PrimNS2Value(symshen_4comb), tmp4786, tmp4787)
												return

											}

										}, 1)

										tmp4789 := Call(__e, PrimNS2Value(symshen_4_5integer_6), V2858)

										tmp4790 := Call(__e, tmp4784, tmp4789)

										__e.TailApply(tmp4781, tmp4790)
										return

									} else {
										__e.Return(Result)
										return
									}

								}, 1)

								tmp4792 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5float_6 := __e.Get(1)
									_ = Parseshen_4_5float_6
									tmp4796 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5float_6)

									if True == tmp4796 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp4794 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5float_6)

										tmp4795 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5float_6)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp4794, tmp4795)
										return

									}

								}, 1)

								tmp4797 := Call(__e, PrimNS2Value(symshen_4_5float_6), V2858)

								tmp4798 := Call(__e, tmp4792, tmp4797)

								__e.TailApply(tmp4779, tmp4798)
								return

							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp4800 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5e_1number_6 := __e.Get(1)
							_ = Parseshen_4_5e_1number_6
							tmp4804 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5e_1number_6)

							if True == tmp4804 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4802 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5e_1number_6)

								tmp4803 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5e_1number_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4802, tmp4803)
								return

							}

						}, 1)

						tmp4805 := Call(__e, PrimNS2Value(symshen_4_5e_1number_6), V2858)

						tmp4806 := Call(__e, tmp4800, tmp4805)

						__e.TailApply(tmp4777, tmp4806)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4808 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5plus_6 := __e.Get(1)
					_ = Parseshen_4_5plus_6
					tmp4816 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5plus_6)

					if True == tmp4816 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4810 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5number_6 := __e.Get(1)
							_ = Parseshen_4_5number_6
							tmp4814 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5number_6)

							if True == tmp4814 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4812 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5number_6)

								tmp4813 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5number_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4812, tmp4813)
								return

							}

						}, 1)

						tmp4815 := Call(__e, PrimNS2Value(symshen_4_5number_6), Parseshen_4_5plus_6)

						__e.TailApply(tmp4810, tmp4815)
						return

					}

				}, 1)

				tmp4817 := Call(__e, PrimNS2Value(symshen_4_5plus_6), V2858)

				tmp4818 := Call(__e, tmp4808, tmp4817)

				__e.TailApply(tmp4775, tmp4818)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4820 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5minus_6 := __e.Get(1)
			_ = Parseshen_4_5minus_6
			tmp4829 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5minus_6)

			if True == tmp4829 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4822 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5number_6 := __e.Get(1)
					_ = Parseshen_4_5number_6
					tmp4827 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5number_6)

					if True == tmp4827 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4824 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5number_6)

						tmp4825 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5number_6)

						tmp4826 := PrimNumberSubtract(MakeNumber(0), tmp4825)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4824, tmp4826)
						return

					}

				}, 1)

				tmp4828 := Call(__e, PrimNS2Value(symshen_4_5number_6), Parseshen_4_5minus_6)

				__e.TailApply(tmp4822, tmp4828)
				return

			}

		}, 1)

		tmp4830 := Call(__e, PrimNS2Value(symshen_4_5minus_6), V2858)

		tmp4831 := Call(__e, tmp4820, tmp4830)

		__e.TailApply(tmp4773, tmp4831)
		return

	}, 1)

	tmp4832 := Call(__e, PrimNS2Value(symdef), symshen_4_5number_6, tmp4772)

	_ = tmp4832

	tmp4833 := MakeNative(func(__e *ControlFlow) {
		V2859 := __e.Get(1)
		_ = V2859
		tmp4834 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4836 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4836 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4843 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2859, MakeNumber(45))

		var ifres4837 Obj

		if True == tmp4843 {
			tmp4839 := MakeNative(func(__e *ControlFlow) {
				News2555 := __e.Get(1)
				_ = News2555
				tmp4840 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2555)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4840, symshen_4skip)
				return

			}, 1)

			tmp4841 := Call(__e, PrimNS2Value(symshen_4tls), V2859)

			tmp4842 := Call(__e, tmp4839, tmp4841)

			ifres4837 = tmp4842

		} else {
			tmp4838 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4837 = tmp4838

		}

		__e.TailApply(tmp4834, ifres4837)
		return

	}, 1)

	tmp4844 := Call(__e, PrimNS2Value(symdef), symshen_4_5minus_6, tmp4833)

	_ = tmp4844

	tmp4845 := MakeNative(func(__e *ControlFlow) {
		V2860 := __e.Get(1)
		_ = V2860
		tmp4846 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4848 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4848 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4855 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2860, MakeNumber(43))

		var ifres4849 Obj

		if True == tmp4855 {
			tmp4851 := MakeNative(func(__e *ControlFlow) {
				News2557 := __e.Get(1)
				_ = News2557
				tmp4852 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2557)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4852, symshen_4skip)
				return

			}, 1)

			tmp4853 := Call(__e, PrimNS2Value(symshen_4tls), V2860)

			tmp4854 := Call(__e, tmp4851, tmp4853)

			ifres4849 = tmp4854

		} else {
			tmp4850 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4849 = tmp4850

		}

		__e.TailApply(tmp4846, ifres4849)
		return

	}, 1)

	tmp4856 := Call(__e, PrimNS2Value(symdef), symshen_4_5plus_6, tmp4845)

	_ = tmp4856

	tmp4857 := MakeNative(func(__e *ControlFlow) {
		V2861 := __e.Get(1)
		_ = V2861
		tmp4858 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4860 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4860 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4861 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5digits_6 := __e.Get(1)
			_ = Parseshen_4_5digits_6
			tmp4866 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5digits_6)

			if True == tmp4866 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4863 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5digits_6)

				tmp4864 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5digits_6)

				tmp4865 := Call(__e, PrimNS2Value(symshen_4compute_1integer), tmp4864)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4863, tmp4865)
				return

			}

		}, 1)

		tmp4867 := Call(__e, PrimNS2Value(symshen_4_5digits_6), V2861)

		tmp4868 := Call(__e, tmp4861, tmp4867)

		__e.TailApply(tmp4858, tmp4868)
		return

	}, 1)

	tmp4869 := Call(__e, PrimNS2Value(symdef), symshen_4_5integer_6, tmp4857)

	_ = tmp4869

	tmp4870 := MakeNative(func(__e *ControlFlow) {
		V2862 := __e.Get(1)
		_ = V2862
		tmp4871 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4884 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4884 {
				tmp4873 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4875 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4875 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4876 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5digit_6 := __e.Get(1)
					_ = Parseshen_4_5digit_6
					tmp4881 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5digit_6)

					if True == tmp4881 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4878 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5digit_6)

						tmp4879 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5digit_6)

						tmp4880 := PrimCons(tmp4879, Nil)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4878, tmp4880)
						return

					}

				}, 1)

				tmp4882 := Call(__e, PrimNS2Value(symshen_4_5digit_6), V2862)

				tmp4883 := Call(__e, tmp4876, tmp4882)

				__e.TailApply(tmp4873, tmp4883)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4885 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5digit_6 := __e.Get(1)
			_ = Parseshen_4_5digit_6
			tmp4895 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5digit_6)

			if True == tmp4895 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4887 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5digits_6 := __e.Get(1)
					_ = Parseshen_4_5digits_6
					tmp4893 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5digits_6)

					if True == tmp4893 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4889 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5digits_6)

						tmp4890 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5digit_6)

						tmp4891 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5digits_6)

						tmp4892 := PrimCons(tmp4890, tmp4891)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4889, tmp4892)
						return

					}

				}, 1)

				tmp4894 := Call(__e, PrimNS2Value(symshen_4_5digits_6), Parseshen_4_5digit_6)

				__e.TailApply(tmp4887, tmp4894)
				return

			}

		}, 1)

		tmp4896 := Call(__e, PrimNS2Value(symshen_4_5digit_6), V2862)

		tmp4897 := Call(__e, tmp4885, tmp4896)

		__e.TailApply(tmp4871, tmp4897)
		return

	}, 1)

	tmp4898 := Call(__e, PrimNS2Value(symdef), symshen_4_5digits_6, tmp4870)

	_ = tmp4898

	tmp4899 := MakeNative(func(__e *ControlFlow) {
		V2863 := __e.Get(1)
		_ = V2863
		tmp4900 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4902 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4902 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4914 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2863)

		var ifres4903 Obj

		if True == tmp4914 {
			tmp4905 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp4906 := MakeNative(func(__e *ControlFlow) {
					News2561 := __e.Get(1)
					_ = News2561
					tmp4910 := Call(__e, PrimNS2Value(symshen_4digit_2), Byte)

					if True == tmp4910 {
						tmp4908 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2561)

						tmp4909 := Call(__e, PrimNS2Value(symshen_4byte_1_6digit), Byte)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp4908, tmp4909)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp4911 := Call(__e, PrimNS2Value(symshen_4tls), V2863)

				__e.TailApply(tmp4906, tmp4911)
				return

			}, 1)

			tmp4912 := Call(__e, PrimNS2Value(symshen_4hds), V2863)

			tmp4913 := Call(__e, tmp4905, tmp4912)

			ifres4903 = tmp4913

		} else {
			tmp4904 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4903 = tmp4904

		}

		__e.TailApply(tmp4900, ifres4903)
		return

	}, 1)

	tmp4915 := Call(__e, PrimNS2Value(symdef), symshen_4_5digit_6, tmp4899)

	_ = tmp4915

	tmp4916 := MakeNative(func(__e *ControlFlow) {
		V2864 := __e.Get(1)
		_ = V2864
		__e.Return(PrimNumberSubtract(V2864, MakeNumber(48)))
		return
	}, 1)

	tmp4917 := Call(__e, PrimNS2Value(symdef), symshen_4byte_1_6digit, tmp4916)

	_ = tmp4917

	tmp4918 := MakeNative(func(__e *ControlFlow) {
		V2865 := __e.Get(1)
		_ = V2865
		tmp4919 := Call(__e, PrimNS2Value(symreverse), V2865)

		__e.TailApply(PrimNS2Value(symshen_4compute_1integer_1h), tmp4919, MakeNumber(0))
		return

	}, 1)

	tmp4920 := Call(__e, PrimNS2Value(symdef), symshen_4compute_1integer, tmp4918)

	_ = tmp4920

	tmp4921 := MakeNative(func(__e *ControlFlow) {
		V2868 := __e.Get(1)
		_ = V2868
		V2869 := __e.Get(2)
		_ = V2869
		tmp4931 := PrimEqual(Nil, V2868)

		if True == tmp4931 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp4930 := PrimIsPair(V2868)

			if True == tmp4930 {
				tmp4924 := Call(__e, PrimNS2Value(symshen_4expt), MakeNumber(10), V2869)

				tmp4925 := PrimHead(V2868)

				tmp4926 := PrimNumberMultiply(tmp4924, tmp4925)

				tmp4927 := PrimTail(V2868)

				tmp4928 := PrimNumberAdd(V2869, MakeNumber(1))

				tmp4929 := Call(__e, PrimNS2Value(symshen_4compute_1integer_1h), tmp4927, tmp4928)

				__e.Return(PrimNumberAdd(tmp4926, tmp4929))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4compute_1integer_1h)
				return
			}

		}

	}, 2)

	tmp4932 := Call(__e, PrimNS2Value(symdef), symshen_4compute_1integer_1h, tmp4921)

	_ = tmp4932

	tmp4933 := MakeNative(func(__e *ControlFlow) {
		V2872 := __e.Get(1)
		_ = V2872
		V2873 := __e.Get(2)
		_ = V2873
		tmp4941 := PrimEqual(MakeNumber(0), V2873)

		if True == tmp4941 {
			__e.Return(MakeNumber(1))
			return
		} else {
			tmp4940 := PrimGreatThan(V2873, MakeNumber(0))

			if True == tmp4940 {
				tmp4938 := PrimNumberSubtract(V2873, MakeNumber(1))

				tmp4939 := Call(__e, PrimNS2Value(symshen_4expt), V2872, tmp4938)

				__e.Return(PrimNumberMultiply(V2872, tmp4939))
				return

			} else {
				tmp4936 := PrimNumberAdd(V2873, MakeNumber(1))

				tmp4937 := Call(__e, PrimNS2Value(symshen_4expt), V2872, tmp4936)

				__e.Return(PrimNumberDivide(tmp4937, V2872))
				return

			}

		}

	}, 2)

	tmp4942 := Call(__e, PrimNS2Value(symdef), symshen_4expt, tmp4933)

	_ = tmp4942

	tmp4943 := MakeNative(func(__e *ControlFlow) {
		V2874 := __e.Get(1)
		_ = V2874
		tmp4944 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4960 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4960 {
				tmp4946 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp4948 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp4948 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp4949 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5stop_6 := __e.Get(1)
					_ = Parseshen_4_5stop_6
					tmp4957 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5stop_6)

					if True == tmp4957 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4951 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5fraction_6 := __e.Get(1)
							_ = Parseshen_4_5fraction_6
							tmp4955 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5fraction_6)

							if True == tmp4955 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4953 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5fraction_6)

								tmp4954 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5fraction_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4953, tmp4954)
								return

							}

						}, 1)

						tmp4956 := Call(__e, PrimNS2Value(symshen_4_5fraction_6), Parseshen_4_5stop_6)

						__e.TailApply(tmp4951, tmp4956)
						return

					}

				}, 1)

				tmp4958 := Call(__e, PrimNS2Value(symshen_4_5stop_6), V2874)

				tmp4959 := Call(__e, tmp4949, tmp4958)

				__e.TailApply(tmp4946, tmp4959)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4961 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5integer_6 := __e.Get(1)
			_ = Parseshen_4_5integer_6
			tmp4975 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

			if True == tmp4975 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4963 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5stop_6 := __e.Get(1)
					_ = Parseshen_4_5stop_6
					tmp4973 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5stop_6)

					if True == tmp4973 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp4965 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5fraction_6 := __e.Get(1)
							_ = Parseshen_4_5fraction_6
							tmp4971 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5fraction_6)

							if True == tmp4971 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp4967 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5fraction_6)

								tmp4968 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5integer_6)

								tmp4969 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5fraction_6)

								tmp4970 := PrimNumberAdd(tmp4968, tmp4969)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp4967, tmp4970)
								return

							}

						}, 1)

						tmp4972 := Call(__e, PrimNS2Value(symshen_4_5fraction_6), Parseshen_4_5stop_6)

						__e.TailApply(tmp4965, tmp4972)
						return

					}

				}, 1)

				tmp4974 := Call(__e, PrimNS2Value(symshen_4_5stop_6), Parseshen_4_5integer_6)

				__e.TailApply(tmp4963, tmp4974)
				return

			}

		}, 1)

		tmp4976 := Call(__e, PrimNS2Value(symshen_4_5integer_6), V2874)

		tmp4977 := Call(__e, tmp4961, tmp4976)

		__e.TailApply(tmp4944, tmp4977)
		return

	}, 1)

	tmp4978 := Call(__e, PrimNS2Value(symdef), symshen_4_5float_6, tmp4943)

	_ = tmp4978

	tmp4979 := MakeNative(func(__e *ControlFlow) {
		V2875 := __e.Get(1)
		_ = V2875
		tmp4980 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4982 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4982 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4989 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2875, MakeNumber(46))

		var ifres4983 Obj

		if True == tmp4989 {
			tmp4985 := MakeNative(func(__e *ControlFlow) {
				News2564 := __e.Get(1)
				_ = News2564
				tmp4986 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2564)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4986, symshen_4skip)
				return

			}, 1)

			tmp4987 := Call(__e, PrimNS2Value(symshen_4tls), V2875)

			tmp4988 := Call(__e, tmp4985, tmp4987)

			ifres4983 = tmp4988

		} else {
			tmp4984 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres4983 = tmp4984

		}

		__e.TailApply(tmp4980, ifres4983)
		return

	}, 1)

	tmp4990 := Call(__e, PrimNS2Value(symdef), symshen_4_5stop_6, tmp4979)

	_ = tmp4990

	tmp4991 := MakeNative(func(__e *ControlFlow) {
		V2876 := __e.Get(1)
		_ = V2876
		tmp4992 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp4994 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp4994 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp4995 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5digits_6 := __e.Get(1)
			_ = Parseshen_4_5digits_6
			tmp5000 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5digits_6)

			if True == tmp5000 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp4997 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5digits_6)

				tmp4998 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5digits_6)

				tmp4999 := Call(__e, PrimNS2Value(symshen_4compute_1fraction), tmp4998)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp4997, tmp4999)
				return

			}

		}, 1)

		tmp5001 := Call(__e, PrimNS2Value(symshen_4_5digits_6), V2876)

		tmp5002 := Call(__e, tmp4995, tmp5001)

		__e.TailApply(tmp4992, tmp5002)
		return

	}, 1)

	tmp5003 := Call(__e, PrimNS2Value(symdef), symshen_4_5fraction_6, tmp4991)

	_ = tmp5003

	tmp5004 := MakeNative(func(__e *ControlFlow) {
		V2877 := __e.Get(1)
		_ = V2877
		__e.TailApply(PrimNS2Value(symshen_4compute_1fraction_1h), V2877, MakeNumber(-1))
		return
	}, 1)

	tmp5005 := Call(__e, PrimNS2Value(symdef), symshen_4compute_1fraction, tmp5004)

	_ = tmp5005

	tmp5006 := MakeNative(func(__e *ControlFlow) {
		V2880 := __e.Get(1)
		_ = V2880
		V2881 := __e.Get(2)
		_ = V2881
		tmp5016 := PrimEqual(Nil, V2880)

		if True == tmp5016 {
			__e.Return(MakeNumber(0))
			return
		} else {
			tmp5015 := PrimIsPair(V2880)

			if True == tmp5015 {
				tmp5009 := Call(__e, PrimNS2Value(symshen_4expt), MakeNumber(10), V2881)

				tmp5010 := PrimHead(V2880)

				tmp5011 := PrimNumberMultiply(tmp5009, tmp5010)

				tmp5012 := PrimTail(V2880)

				tmp5013 := PrimNumberSubtract(V2881, MakeNumber(1))

				tmp5014 := Call(__e, PrimNS2Value(symshen_4compute_1fraction_1h), tmp5012, tmp5013)

				__e.Return(PrimNumberAdd(tmp5011, tmp5014))
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4compute_1fraction_1h)
				return
			}

		}

	}, 2)

	tmp5017 := Call(__e, PrimNS2Value(symdef), symshen_4compute_1fraction_1h, tmp5006)

	_ = tmp5017

	tmp5018 := MakeNative(func(__e *ControlFlow) {
		V2882 := __e.Get(1)
		_ = V2882
		tmp5019 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5041 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp5041 {
				tmp5021 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5023 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp5023 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5024 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5integer_6 := __e.Get(1)
					_ = Parseshen_4_5integer_6
					tmp5038 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

					if True == tmp5038 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp5026 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5lowE_6 := __e.Get(1)
							_ = Parseshen_4_5lowE_6
							tmp5036 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5lowE_6)

							if True == tmp5036 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp5028 := MakeNative(func(__e *ControlFlow) {
									Parseshen_4_5log10_6 := __e.Get(1)
									_ = Parseshen_4_5log10_6
									tmp5034 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

									if True == tmp5034 {
										__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
										return
									} else {
										tmp5030 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5log10_6)

										tmp5031 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5integer_6)

										tmp5032 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5log10_6)

										tmp5033 := Call(__e, PrimNS2Value(symshen_4compute_1E), tmp5031, tmp5032)

										__e.TailApply(PrimNS2Value(symshen_4comb), tmp5030, tmp5033)
										return

									}

								}, 1)

								tmp5035 := Call(__e, PrimNS2Value(symshen_4_5log10_6), Parseshen_4_5lowE_6)

								__e.TailApply(tmp5028, tmp5035)
								return

							}

						}, 1)

						tmp5037 := Call(__e, PrimNS2Value(symshen_4_5lowE_6), Parseshen_4_5integer_6)

						__e.TailApply(tmp5026, tmp5037)
						return

					}

				}, 1)

				tmp5039 := Call(__e, PrimNS2Value(symshen_4_5integer_6), V2882)

				tmp5040 := Call(__e, tmp5024, tmp5039)

				__e.TailApply(tmp5021, tmp5040)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5042 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5float_6 := __e.Get(1)
			_ = Parseshen_4_5float_6
			tmp5056 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5float_6)

			if True == tmp5056 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp5044 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5lowE_6 := __e.Get(1)
					_ = Parseshen_4_5lowE_6
					tmp5054 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5lowE_6)

					if True == tmp5054 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp5046 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5log10_6 := __e.Get(1)
							_ = Parseshen_4_5log10_6
							tmp5052 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

							if True == tmp5052 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp5048 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5log10_6)

								tmp5049 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5float_6)

								tmp5050 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5log10_6)

								tmp5051 := Call(__e, PrimNS2Value(symshen_4compute_1E), tmp5049, tmp5050)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp5048, tmp5051)
								return

							}

						}, 1)

						tmp5053 := Call(__e, PrimNS2Value(symshen_4_5log10_6), Parseshen_4_5lowE_6)

						__e.TailApply(tmp5046, tmp5053)
						return

					}

				}, 1)

				tmp5055 := Call(__e, PrimNS2Value(symshen_4_5lowE_6), Parseshen_4_5float_6)

				__e.TailApply(tmp5044, tmp5055)
				return

			}

		}, 1)

		tmp5057 := Call(__e, PrimNS2Value(symshen_4_5float_6), V2882)

		tmp5058 := Call(__e, tmp5042, tmp5057)

		__e.TailApply(tmp5019, tmp5058)
		return

	}, 1)

	tmp5059 := Call(__e, PrimNS2Value(symdef), symshen_4_5e_1number_6, tmp5018)

	_ = tmp5059

	tmp5060 := MakeNative(func(__e *ControlFlow) {
		V2883 := __e.Get(1)
		_ = V2883
		tmp5061 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5088 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp5088 {
				tmp5063 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5075 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp5075 {
						tmp5065 := MakeNative(func(__e *ControlFlow) {
							Result := __e.Get(1)
							_ = Result
							tmp5067 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

							if True == tmp5067 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								__e.Return(Result)
								return
							}

						}, 1)

						tmp5068 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5integer_6 := __e.Get(1)
							_ = Parseshen_4_5integer_6
							tmp5072 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5integer_6)

							if True == tmp5072 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp5070 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5integer_6)

								tmp5071 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5integer_6)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp5070, tmp5071)
								return

							}

						}, 1)

						tmp5073 := Call(__e, PrimNS2Value(symshen_4_5integer_6), V2883)

						tmp5074 := Call(__e, tmp5068, tmp5073)

						__e.TailApply(tmp5065, tmp5074)
						return

					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5076 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5minus_6 := __e.Get(1)
					_ = Parseshen_4_5minus_6
					tmp5085 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5minus_6)

					if True == tmp5085 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp5078 := MakeNative(func(__e *ControlFlow) {
							Parseshen_4_5log10_6 := __e.Get(1)
							_ = Parseshen_4_5log10_6
							tmp5083 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

							if True == tmp5083 {
								__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
								return
							} else {
								tmp5080 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5log10_6)

								tmp5081 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5log10_6)

								tmp5082 := PrimNumberSubtract(MakeNumber(0), tmp5081)

								__e.TailApply(PrimNS2Value(symshen_4comb), tmp5080, tmp5082)
								return

							}

						}, 1)

						tmp5084 := Call(__e, PrimNS2Value(symshen_4_5log10_6), Parseshen_4_5minus_6)

						__e.TailApply(tmp5078, tmp5084)
						return

					}

				}, 1)

				tmp5086 := Call(__e, PrimNS2Value(symshen_4_5minus_6), V2883)

				tmp5087 := Call(__e, tmp5076, tmp5086)

				__e.TailApply(tmp5063, tmp5087)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5089 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5plus_6 := __e.Get(1)
			_ = Parseshen_4_5plus_6
			tmp5097 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5plus_6)

			if True == tmp5097 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp5091 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5log10_6 := __e.Get(1)
					_ = Parseshen_4_5log10_6
					tmp5095 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5log10_6)

					if True == tmp5095 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp5093 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5log10_6)

						tmp5094 := Call(__e, PrimNS2Value(symshen_4_5_1out), Parseshen_4_5log10_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp5093, tmp5094)
						return

					}

				}, 1)

				tmp5096 := Call(__e, PrimNS2Value(symshen_4_5log10_6), Parseshen_4_5plus_6)

				__e.TailApply(tmp5091, tmp5096)
				return

			}

		}, 1)

		tmp5098 := Call(__e, PrimNS2Value(symshen_4_5plus_6), V2883)

		tmp5099 := Call(__e, tmp5089, tmp5098)

		__e.TailApply(tmp5061, tmp5099)
		return

	}, 1)

	tmp5100 := Call(__e, PrimNS2Value(symdef), symshen_4_5log10_6, tmp5060)

	_ = tmp5100

	tmp5101 := MakeNative(func(__e *ControlFlow) {
		V2884 := __e.Get(1)
		_ = V2884
		tmp5102 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5104 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp5104 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5111 := Call(__e, PrimNS2Value(symshen_4_ahd_2), V2884, MakeNumber(101))

		var ifres5105 Obj

		if True == tmp5111 {
			tmp5107 := MakeNative(func(__e *ControlFlow) {
				News2569 := __e.Get(1)
				_ = News2569
				tmp5108 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2569)

				__e.TailApply(PrimNS2Value(symshen_4comb), tmp5108, symshen_4skip)
				return

			}, 1)

			tmp5109 := Call(__e, PrimNS2Value(symshen_4tls), V2884)

			tmp5110 := Call(__e, tmp5107, tmp5109)

			ifres5105 = tmp5110

		} else {
			tmp5106 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres5105 = tmp5106

		}

		__e.TailApply(tmp5102, ifres5105)
		return

	}, 1)

	tmp5112 := Call(__e, PrimNS2Value(symdef), symshen_4_5lowE_6, tmp5101)

	_ = tmp5112

	tmp5113 := MakeNative(func(__e *ControlFlow) {
		V2885 := __e.Get(1)
		_ = V2885
		V2886 := __e.Get(2)
		_ = V2886
		tmp5114 := Call(__e, PrimNS2Value(symshen_4expt), MakeNumber(10), V2886)

		__e.Return(PrimNumberMultiply(V2885, tmp5114))
		return

	}, 2)

	tmp5115 := Call(__e, PrimNS2Value(symdef), symshen_4compute_1E, tmp5113)

	_ = tmp5115

	tmp5116 := MakeNative(func(__e *ControlFlow) {
		V2887 := __e.Get(1)
		_ = V2887
		tmp5117 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5128 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp5128 {
				tmp5119 := MakeNative(func(__e *ControlFlow) {
					Result := __e.Get(1)
					_ = Result
					tmp5121 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

					if True == tmp5121 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						__e.Return(Result)
						return
					}

				}, 1)

				tmp5122 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5whitespace_6 := __e.Get(1)
					_ = Parseshen_4_5whitespace_6
					tmp5125 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5whitespace_6)

					if True == tmp5125 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp5124 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5whitespace_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp5124, symshen_4skip)
						return

					}

				}, 1)

				tmp5126 := Call(__e, PrimNS2Value(symshen_4_5whitespace_6), V2887)

				tmp5127 := Call(__e, tmp5122, tmp5126)

				__e.TailApply(tmp5119, tmp5127)
				return

			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5129 := MakeNative(func(__e *ControlFlow) {
			Parseshen_4_5whitespace_6 := __e.Get(1)
			_ = Parseshen_4_5whitespace_6
			tmp5136 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5whitespace_6)

			if True == tmp5136 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				tmp5131 := MakeNative(func(__e *ControlFlow) {
					Parseshen_4_5whitespaces_6 := __e.Get(1)
					_ = Parseshen_4_5whitespaces_6
					tmp5134 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Parseshen_4_5whitespaces_6)

					if True == tmp5134 {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					} else {
						tmp5133 := Call(__e, PrimNS2Value(symshen_4in_1_6), Parseshen_4_5whitespaces_6)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp5133, symshen_4skip)
						return

					}

				}, 1)

				tmp5135 := Call(__e, PrimNS2Value(symshen_4_5whitespaces_6), Parseshen_4_5whitespace_6)

				__e.TailApply(tmp5131, tmp5135)
				return

			}

		}, 1)

		tmp5137 := Call(__e, PrimNS2Value(symshen_4_5whitespace_6), V2887)

		tmp5138 := Call(__e, tmp5129, tmp5137)

		__e.TailApply(tmp5117, tmp5138)
		return

	}, 1)

	tmp5139 := Call(__e, PrimNS2Value(symdef), symshen_4_5whitespaces_6, tmp5116)

	_ = tmp5139

	tmp5140 := MakeNative(func(__e *ControlFlow) {
		V2888 := __e.Get(1)
		_ = V2888
		tmp5141 := MakeNative(func(__e *ControlFlow) {
			Result := __e.Get(1)
			_ = Result
			tmp5143 := Call(__e, PrimNS2Value(symshen_4parse_1failure_2), Result)

			if True == tmp5143 {
				__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
				return
			} else {
				__e.Return(Result)
				return
			}

		}, 1)

		tmp5154 := Call(__e, PrimNS2Value(symshen_4non_1empty_1stream_2), V2888)

		var ifres5144 Obj

		if True == tmp5154 {
			tmp5146 := MakeNative(func(__e *ControlFlow) {
				Byte := __e.Get(1)
				_ = Byte
				tmp5147 := MakeNative(func(__e *ControlFlow) {
					News2572 := __e.Get(1)
					_ = News2572
					tmp5150 := Call(__e, PrimNS2Value(symshen_4whitespace_2), Byte)

					if True == tmp5150 {
						tmp5149 := Call(__e, PrimNS2Value(symshen_4in_1_6), News2572)

						__e.TailApply(PrimNS2Value(symshen_4comb), tmp5149, symshen_4skip)
						return

					} else {
						__e.TailApply(PrimNS2Value(symshen_4parse_1failure))
						return
					}

				}, 1)

				tmp5151 := Call(__e, PrimNS2Value(symshen_4tls), V2888)

				__e.TailApply(tmp5147, tmp5151)
				return

			}, 1)

			tmp5152 := Call(__e, PrimNS2Value(symshen_4hds), V2888)

			tmp5153 := Call(__e, tmp5146, tmp5152)

			ifres5144 = tmp5153

		} else {
			tmp5145 := Call(__e, PrimNS2Value(symshen_4parse_1failure))

			ifres5144 = tmp5145

		}

		__e.TailApply(tmp5141, ifres5144)
		return

	}, 1)

	tmp5155 := Call(__e, PrimNS2Value(symdef), symshen_4_5whitespace_6, tmp5140)

	_ = tmp5155

	tmp5156 := MakeNative(func(__e *ControlFlow) {
		V2891 := __e.Get(1)
		_ = V2891
		tmp5164 := PrimEqual(MakeNumber(32), V2891)

		if True == tmp5164 {
			__e.Return(True)
			return
		} else {
			tmp5163 := PrimEqual(MakeNumber(13), V2891)

			if True == tmp5163 {
				__e.Return(True)
				return
			} else {
				tmp5162 := PrimEqual(MakeNumber(10), V2891)

				if True == tmp5162 {
					__e.Return(True)
					return
				} else {
					tmp5161 := PrimEqual(MakeNumber(9), V2891)

					if True == tmp5161 {
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

	tmp5165 := Call(__e, PrimNS2Value(symdef), symshen_4whitespace_2, tmp5156)

	_ = tmp5165

	tmp5166 := MakeNative(func(__e *ControlFlow) {
		V2892 := __e.Get(1)
		_ = V2892
		tmp5189 := PrimEqual(Nil, V2892)

		if True == tmp5189 {
			__e.Return(Nil)
			return
		} else {
			tmp5188 := PrimIsPair(V2892)

			var ifres5184 Obj

			if True == tmp5188 {
				tmp5186 := PrimHead(V2892)

				tmp5187 := Call(__e, PrimNS2Value(symshen_4packaged_2), tmp5186)

				var ifres5185 Obj

				if True == tmp5187 {
					ifres5185 = True

				} else {
					ifres5185 = False

				}

				ifres5184 = ifres5185

			} else {
				ifres5184 = False

			}

			if True == ifres5184 {
				tmp5180 := PrimHead(V2892)

				tmp5181 := Call(__e, PrimNS2Value(symshen_4unpackage), tmp5180)

				tmp5182 := PrimTail(V2892)

				tmp5183 := Call(__e, PrimNS2Value(symappend), tmp5181, tmp5182)

				__e.TailApply(PrimNS2Value(symshen_4unpackage_emacroexpand), tmp5183)
				return

			} else {
				tmp5179 := PrimIsPair(V2892)

				if True == tmp5179 {
					tmp5170 := MakeNative(func(__e *ControlFlow) {
						M := __e.Get(1)
						_ = M
						tmp5176 := Call(__e, PrimNS2Value(symshen_4packaged_2), M)

						if True == tmp5176 {
							tmp5174 := PrimTail(V2892)

							tmp5175 := PrimCons(M, tmp5174)

							__e.TailApply(PrimNS2Value(symshen_4unpackage_emacroexpand), tmp5175)
							return

						} else {
							tmp5172 := PrimTail(V2892)

							tmp5173 := Call(__e, PrimNS2Value(symshen_4unpackage_emacroexpand), tmp5172)

							__e.Return(PrimCons(M, tmp5173))
							return

						}

					}, 1)

					tmp5177 := PrimHead(V2892)

					tmp5178 := Call(__e, PrimNS2Value(symmacroexpand), tmp5177)

					__e.TailApply(tmp5170, tmp5178)
					return

				} else {
					__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4unpackage_emacroexpand)
					return
				}

			}

		}

	}, 1)

	tmp5190 := Call(__e, PrimNS2Value(symdef), symshen_4unpackage_emacroexpand, tmp5166)

	_ = tmp5190

	tmp5191 := MakeNative(func(__e *ControlFlow) {
		V2895 := __e.Get(1)
		_ = V2895
		tmp5206 := PrimIsPair(V2895)

		var ifres5193 Obj

		if True == tmp5206 {
			tmp5204 := PrimHead(V2895)

			tmp5205 := PrimEqual(sympackage, tmp5204)

			var ifres5195 Obj

			if True == tmp5205 {
				tmp5202 := PrimTail(V2895)

				tmp5203 := PrimIsPair(tmp5202)

				var ifres5197 Obj

				if True == tmp5203 {
					tmp5199 := PrimTail(V2895)

					tmp5200 := PrimTail(tmp5199)

					tmp5201 := PrimIsPair(tmp5200)

					var ifres5198 Obj

					if True == tmp5201 {
						ifres5198 = True

					} else {
						ifres5198 = False

					}

					ifres5197 = ifres5198

				} else {
					ifres5197 = False

				}

				var ifres5196 Obj

				if True == ifres5197 {
					ifres5196 = True

				} else {
					ifres5196 = False

				}

				ifres5195 = ifres5196

			} else {
				ifres5195 = False

			}

			var ifres5194 Obj

			if True == ifres5195 {
				ifres5194 = True

			} else {
				ifres5194 = False

			}

			ifres5193 = ifres5194

		} else {
			ifres5193 = False

		}

		if True == ifres5193 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp5207 := Call(__e, PrimNS2Value(symdef), symshen_4packaged_2, tmp5191)

	_ = tmp5207

	tmp5208 := MakeNative(func(__e *ControlFlow) {
		V2898 := __e.Get(1)
		_ = V2898
		tmp5262 := PrimIsPair(V2898)

		var ifres5244 Obj

		if True == tmp5262 {
			tmp5260 := PrimHead(V2898)

			tmp5261 := PrimEqual(sympackage, tmp5260)

			var ifres5246 Obj

			if True == tmp5261 {
				tmp5258 := PrimTail(V2898)

				tmp5259 := PrimIsPair(tmp5258)

				var ifres5248 Obj

				if True == tmp5259 {
					tmp5255 := PrimTail(V2898)

					tmp5256 := PrimHead(tmp5255)

					tmp5257 := PrimEqual(symnull, tmp5256)

					var ifres5250 Obj

					if True == tmp5257 {
						tmp5252 := PrimTail(V2898)

						tmp5253 := PrimTail(tmp5252)

						tmp5254 := PrimIsPair(tmp5253)

						var ifres5251 Obj

						if True == tmp5254 {
							ifres5251 = True

						} else {
							ifres5251 = False

						}

						ifres5250 = ifres5251

					} else {
						ifres5250 = False

					}

					var ifres5249 Obj

					if True == ifres5250 {
						ifres5249 = True

					} else {
						ifres5249 = False

					}

					ifres5248 = ifres5249

				} else {
					ifres5248 = False

				}

				var ifres5247 Obj

				if True == ifres5248 {
					ifres5247 = True

				} else {
					ifres5247 = False

				}

				ifres5246 = ifres5247

			} else {
				ifres5246 = False

			}

			var ifres5245 Obj

			if True == ifres5246 {
				ifres5245 = True

			} else {
				ifres5245 = False

			}

			ifres5244 = ifres5245

		} else {
			ifres5244 = False

		}

		if True == ifres5244 {
			tmp5242 := PrimTail(V2898)

			tmp5243 := PrimTail(tmp5242)

			__e.Return(PrimTail(tmp5243))
			return

		} else {
			tmp5241 := PrimIsPair(V2898)

			var ifres5228 Obj

			if True == tmp5241 {
				tmp5239 := PrimHead(V2898)

				tmp5240 := PrimEqual(sympackage, tmp5239)

				var ifres5230 Obj

				if True == tmp5240 {
					tmp5237 := PrimTail(V2898)

					tmp5238 := PrimIsPair(tmp5237)

					var ifres5232 Obj

					if True == tmp5238 {
						tmp5234 := PrimTail(V2898)

						tmp5235 := PrimTail(tmp5234)

						tmp5236 := PrimIsPair(tmp5235)

						var ifres5233 Obj

						if True == tmp5236 {
							ifres5233 = True

						} else {
							ifres5233 = False

						}

						ifres5232 = ifres5233

					} else {
						ifres5232 = False

					}

					var ifres5231 Obj

					if True == ifres5232 {
						ifres5231 = True

					} else {
						ifres5231 = False

					}

					ifres5230 = ifres5231

				} else {
					ifres5230 = False

				}

				var ifres5229 Obj

				if True == ifres5230 {
					ifres5229 = True

				} else {
					ifres5229 = False

				}

				ifres5228 = ifres5229

			} else {
				ifres5228 = False

			}

			if True == ifres5228 {
				tmp5211 := MakeNative(func(__e *ControlFlow) {
					External_b := __e.Get(1)
					_ = External_b
					tmp5212 := MakeNative(func(__e *ControlFlow) {
						Package := __e.Get(1)
						_ = Package
						tmp5213 := MakeNative(func(__e *ControlFlow) {
							RecordExternal := __e.Get(1)
							_ = RecordExternal
							__e.Return(Package)
							return
						}, 1)

						tmp5214 := PrimTail(V2898)

						tmp5215 := PrimHead(tmp5214)

						tmp5216 := Call(__e, PrimNS2Value(symshen_4record_1external), tmp5215, External_b)

						__e.TailApply(tmp5213, tmp5216)
						return

					}, 1)

					tmp5217 := PrimTail(V2898)

					tmp5218 := PrimHead(tmp5217)

					tmp5219 := PrimStr(tmp5218)

					tmp5220 := PrimTail(V2898)

					tmp5221 := PrimTail(tmp5220)

					tmp5222 := PrimTail(tmp5221)

					tmp5223 := Call(__e, PrimNS2Value(symshen_4package_1symbols), tmp5219, External_b, tmp5222)

					__e.TailApply(tmp5212, tmp5223)
					return

				}, 1)

				tmp5224 := PrimTail(V2898)

				tmp5225 := PrimTail(tmp5224)

				tmp5226 := PrimHead(tmp5225)

				tmp5227 := Call(__e, PrimNS2Value(symeval), tmp5226)

				__e.TailApply(tmp5211, tmp5227)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4unpackage)
				return
			}

		}

	}, 1)

	tmp5263 := Call(__e, PrimNS2Value(symdef), symshen_4unpackage, tmp5208)

	_ = tmp5263

	tmp5264 := MakeNative(func(__e *ControlFlow) {
		V2899 := __e.Get(1)
		_ = V2899
		V2900 := __e.Get(2)
		_ = V2900
		tmp5265 := MakeNative(func(__e *ControlFlow) {
			External := __e.Get(1)
			_ = External
			tmp5266 := Call(__e, PrimNS2Value(symunion), V2900, External)

			tmp5267 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symput), V2899, symshen_4external_1symbols, tmp5266, tmp5267)
			return

		}, 1)

		tmp5268 := MakeNative(func(__e *ControlFlow) {
			tmp5269 := PrimNS3Value(sym_dproperty_1vector_d)

			__e.TailApply(PrimNS2Value(symget), V2899, symshen_4external_1symbols, tmp5269)
			return

		}, 0)

		tmp5270 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		tmp5271 := Call(__e, PrimNS1Value(symtry_1catch), tmp5268, tmp5270)

		__e.TailApply(tmp5265, tmp5271)
		return

	}, 2)

	tmp5272 := Call(__e, PrimNS2Value(symdef), symshen_4record_1external, tmp5264)

	_ = tmp5272

	tmp5273 := MakeNative(func(__e *ControlFlow) {
		V2905 := __e.Get(1)
		_ = V2905
		V2906 := __e.Get(2)
		_ = V2906
		V2907 := __e.Get(3)
		_ = V2907
		tmp5278 := PrimIsPair(V2907)

		if True == tmp5278 {
			tmp5277 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				__e.TailApply(PrimNS2Value(symshen_4package_1symbols), V2905, V2906, X)
				return
			}, 1)

			__e.TailApply(PrimNS2Value(symmap), tmp5277, V2907)
			return

		} else {
			tmp5276 := Call(__e, PrimNS2Value(symshen_4internal_2), V2907, V2905, V2906)

			if True == tmp5276 {
				__e.TailApply(PrimNS2Value(symshen_4intern_1in_1package), V2905, V2907)
				return
			} else {
				__e.Return(V2907)
				return
			}

		}

	}, 3)

	tmp5279 := Call(__e, PrimNS2Value(symdef), symshen_4package_1symbols, tmp5273)

	_ = tmp5279

	tmp5280 := MakeNative(func(__e *ControlFlow) {
		V2908 := __e.Get(1)
		_ = V2908
		V2909 := __e.Get(2)
		_ = V2909
		tmp5281 := PrimStr(V2909)

		tmp5282 := Call(__e, PrimNS2Value(sym_8s), MakeString("."), tmp5281)

		tmp5283 := Call(__e, PrimNS2Value(sym_8s), V2908, tmp5282)

		__e.Return(PrimIntern(tmp5283))
		return

	}, 2)

	tmp5284 := Call(__e, PrimNS2Value(symdef), symshen_4intern_1in_1package, tmp5280)

	_ = tmp5284

	tmp5285 := MakeNative(func(__e *ControlFlow) {
		V2910 := __e.Get(1)
		_ = V2910
		V2911 := __e.Get(2)
		_ = V2911
		V2912 := __e.Get(3)
		_ = V2912
		tmp5315 := Call(__e, PrimNS2Value(symelement_2), V2910, V2912)

		tmp5316 := PrimNot(tmp5315)

		if True == tmp5316 {
			tmp5313 := Call(__e, PrimNS2Value(symshen_4sng_2), V2910)

			tmp5314 := PrimNot(tmp5313)

			var ifres5288 Obj

			if True == tmp5314 {
				tmp5311 := Call(__e, PrimNS2Value(symshen_4dbl_2), V2910)

				tmp5312 := PrimNot(tmp5311)

				var ifres5290 Obj

				if True == tmp5312 {
					tmp5310 := PrimIsSymbol(V2910)

					var ifres5292 Obj

					if True == tmp5310 {
						tmp5308 := Call(__e, PrimNS2Value(symshen_4sysfunc_2), V2910)

						tmp5309 := PrimNot(tmp5308)

						var ifres5294 Obj

						if True == tmp5309 {
							tmp5306 := PrimIsVariable(V2910)

							tmp5307 := PrimNot(tmp5306)

							var ifres5296 Obj

							if True == tmp5307 {
								tmp5303 := PrimStr(V2910)

								tmp5304 := Call(__e, PrimNS2Value(symshen_4internal_1to_1shen_2), tmp5303)

								tmp5305 := PrimNot(tmp5304)

								var ifres5298 Obj

								if True == tmp5305 {
									tmp5300 := PrimStr(V2910)

									tmp5301 := Call(__e, PrimNS2Value(symshen_4internal_1to_1P_2), V2911, tmp5300)

									tmp5302 := PrimNot(tmp5301)

									var ifres5299 Obj

									if True == tmp5302 {
										ifres5299 = True

									} else {
										ifres5299 = False

									}

									ifres5298 = ifres5299

								} else {
									ifres5298 = False

								}

								var ifres5297 Obj

								if True == ifres5298 {
									ifres5297 = True

								} else {
									ifres5297 = False

								}

								ifres5296 = ifres5297

							} else {
								ifres5296 = False

							}

							var ifres5295 Obj

							if True == ifres5296 {
								ifres5295 = True

							} else {
								ifres5295 = False

							}

							ifres5294 = ifres5295

						} else {
							ifres5294 = False

						}

						var ifres5293 Obj

						if True == ifres5294 {
							ifres5293 = True

						} else {
							ifres5293 = False

						}

						ifres5292 = ifres5293

					} else {
						ifres5292 = False

					}

					var ifres5291 Obj

					if True == ifres5292 {
						ifres5291 = True

					} else {
						ifres5291 = False

					}

					ifres5290 = ifres5291

				} else {
					ifres5290 = False

				}

				var ifres5289 Obj

				if True == ifres5290 {
					ifres5289 = True

				} else {
					ifres5289 = False

				}

				ifres5288 = ifres5289

			} else {
				ifres5288 = False

			}

			if True == ifres5288 {
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

	tmp5317 := Call(__e, PrimNS2Value(symdef), symshen_4internal_2, tmp5285)

	_ = tmp5317

	tmp5318 := MakeNative(func(__e *ControlFlow) {
		V2917 := __e.Get(1)
		_ = V2917
		tmp5372 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2917)

		var ifres5320 Obj

		if True == tmp5372 {
			tmp5370 := Call(__e, PrimNS2Value(symhdstr), V2917)

			tmp5371 := PrimEqual(MakeString("s"), tmp5370)

			var ifres5322 Obj

			if True == tmp5371 {
				tmp5368 := PrimTailString(V2917)

				tmp5369 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5368)

				var ifres5324 Obj

				if True == tmp5369 {
					tmp5365 := PrimTailString(V2917)

					tmp5366 := Call(__e, PrimNS2Value(symhdstr), tmp5365)

					tmp5367 := PrimEqual(MakeString("h"), tmp5366)

					var ifres5326 Obj

					if True == tmp5367 {
						tmp5362 := PrimTailString(V2917)

						tmp5363 := PrimTailString(tmp5362)

						tmp5364 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5363)

						var ifres5328 Obj

						if True == tmp5364 {
							tmp5358 := PrimTailString(V2917)

							tmp5359 := PrimTailString(tmp5358)

							tmp5360 := Call(__e, PrimNS2Value(symhdstr), tmp5359)

							tmp5361 := PrimEqual(MakeString("e"), tmp5360)

							var ifres5330 Obj

							if True == tmp5361 {
								tmp5354 := PrimTailString(V2917)

								tmp5355 := PrimTailString(tmp5354)

								tmp5356 := PrimTailString(tmp5355)

								tmp5357 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5356)

								var ifres5332 Obj

								if True == tmp5357 {
									tmp5349 := PrimTailString(V2917)

									tmp5350 := PrimTailString(tmp5349)

									tmp5351 := PrimTailString(tmp5350)

									tmp5352 := Call(__e, PrimNS2Value(symhdstr), tmp5351)

									tmp5353 := PrimEqual(MakeString("n"), tmp5352)

									var ifres5334 Obj

									if True == tmp5353 {
										tmp5344 := PrimTailString(V2917)

										tmp5345 := PrimTailString(tmp5344)

										tmp5346 := PrimTailString(tmp5345)

										tmp5347 := PrimTailString(tmp5346)

										tmp5348 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5347)

										var ifres5336 Obj

										if True == tmp5348 {
											tmp5338 := PrimTailString(V2917)

											tmp5339 := PrimTailString(tmp5338)

											tmp5340 := PrimTailString(tmp5339)

											tmp5341 := PrimTailString(tmp5340)

											tmp5342 := Call(__e, PrimNS2Value(symhdstr), tmp5341)

											tmp5343 := PrimEqual(MakeString("."), tmp5342)

											var ifres5337 Obj

											if True == tmp5343 {
												ifres5337 = True

											} else {
												ifres5337 = False

											}

											ifres5336 = ifres5337

										} else {
											ifres5336 = False

										}

										var ifres5335 Obj

										if True == ifres5336 {
											ifres5335 = True

										} else {
											ifres5335 = False

										}

										ifres5334 = ifres5335

									} else {
										ifres5334 = False

									}

									var ifres5333 Obj

									if True == ifres5334 {
										ifres5333 = True

									} else {
										ifres5333 = False

									}

									ifres5332 = ifres5333

								} else {
									ifres5332 = False

								}

								var ifres5331 Obj

								if True == ifres5332 {
									ifres5331 = True

								} else {
									ifres5331 = False

								}

								ifres5330 = ifres5331

							} else {
								ifres5330 = False

							}

							var ifres5329 Obj

							if True == ifres5330 {
								ifres5329 = True

							} else {
								ifres5329 = False

							}

							ifres5328 = ifres5329

						} else {
							ifres5328 = False

						}

						var ifres5327 Obj

						if True == ifres5328 {
							ifres5327 = True

						} else {
							ifres5327 = False

						}

						ifres5326 = ifres5327

					} else {
						ifres5326 = False

					}

					var ifres5325 Obj

					if True == ifres5326 {
						ifres5325 = True

					} else {
						ifres5325 = False

					}

					ifres5324 = ifres5325

				} else {
					ifres5324 = False

				}

				var ifres5323 Obj

				if True == ifres5324 {
					ifres5323 = True

				} else {
					ifres5323 = False

				}

				ifres5322 = ifres5323

			} else {
				ifres5322 = False

			}

			var ifres5321 Obj

			if True == ifres5322 {
				ifres5321 = True

			} else {
				ifres5321 = False

			}

			ifres5320 = ifres5321

		} else {
			ifres5320 = False

		}

		if True == ifres5320 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp5373 := Call(__e, PrimNS2Value(symdef), symshen_4internal_1to_1shen_2, tmp5318)

	_ = tmp5373

	tmp5374 := MakeNative(func(__e *ControlFlow) {
		V2918 := __e.Get(1)
		_ = V2918
		tmp5375 := PrimNS3Value(sym_dproperty_1vector_d)

		tmp5376 := Call(__e, PrimNS2Value(symget), symshen, symshen_4external_1symbols, tmp5375)

		__e.TailApply(PrimNS2Value(symelement_2), V2918, tmp5376)
		return

	}, 1)

	tmp5377 := Call(__e, PrimNS2Value(symdef), symshen_4sysfunc_2, tmp5374)

	_ = tmp5377

	tmp5378 := MakeNative(func(__e *ControlFlow) {
		V2926 := __e.Get(1)
		_ = V2926
		V2927 := __e.Get(2)
		_ = V2927
		tmp5399 := PrimEqual(MakeString(""), V2926)

		var ifres5392 Obj

		if True == tmp5399 {
			tmp5398 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2927)

			var ifres5394 Obj

			if True == tmp5398 {
				tmp5396 := Call(__e, PrimNS2Value(symhdstr), V2927)

				tmp5397 := PrimEqual(MakeString("."), tmp5396)

				var ifres5395 Obj

				if True == tmp5397 {
					ifres5395 = True

				} else {
					ifres5395 = False

				}

				ifres5394 = ifres5395

			} else {
				ifres5394 = False

			}

			var ifres5393 Obj

			if True == ifres5394 {
				ifres5393 = True

			} else {
				ifres5393 = False

			}

			ifres5392 = ifres5393

		} else {
			ifres5392 = False

		}

		if True == ifres5392 {
			__e.Return(True)
			return
		} else {
			tmp5391 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2926)

			var ifres5383 Obj

			if True == tmp5391 {
				tmp5390 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2927)

				var ifres5385 Obj

				if True == tmp5390 {
					tmp5387 := Call(__e, PrimNS2Value(symhdstr), V2926)

					tmp5388 := Call(__e, PrimNS2Value(symhdstr), V2927)

					tmp5389 := PrimEqual(tmp5387, tmp5388)

					var ifres5386 Obj

					if True == tmp5389 {
						ifres5386 = True

					} else {
						ifres5386 = False

					}

					ifres5385 = ifres5386

				} else {
					ifres5385 = False

				}

				var ifres5384 Obj

				if True == ifres5385 {
					ifres5384 = True

				} else {
					ifres5384 = False

				}

				ifres5383 = ifres5384

			} else {
				ifres5383 = False

			}

			if True == ifres5383 {
				tmp5381 := PrimTailString(V2926)

				tmp5382 := PrimTailString(V2927)

				__e.TailApply(PrimNS2Value(symshen_4internal_1to_1P_2), tmp5381, tmp5382)
				return

			} else {
				__e.Return(False)
				return
			}

		}

	}, 2)

	tmp5400 := Call(__e, PrimNS2Value(symdef), symshen_4internal_1to_1P_2, tmp5378)

	_ = tmp5400

	tmp5401 := MakeNative(func(__e *ControlFlow) {
		V2930 := __e.Get(1)
		_ = V2930
		V2931 := __e.Get(2)
		_ = V2931
		tmp5414 := Call(__e, PrimNS2Value(symelement_2), V2930, V2931)

		if True == tmp5414 {
			__e.Return(V2930)
			return
		} else {
			tmp5413 := PrimIsPair(V2930)

			var ifres5409 Obj

			if True == tmp5413 {
				tmp5411 := PrimHead(V2930)

				tmp5412 := Call(__e, PrimNS2Value(symshen_4non_1application_2), tmp5411)

				var ifres5410 Obj

				if True == tmp5412 {
					ifres5410 = True

				} else {
					ifres5410 = False

				}

				ifres5409 = ifres5410

			} else {
				ifres5409 = False

			}

			if True == ifres5409 {
				tmp5408 := PrimHead(V2930)

				__e.TailApply(PrimNS2Value(symshen_4special_1case), tmp5408, V2930, V2931)
				return

			} else {
				tmp5407 := PrimIsPair(V2930)

				if True == tmp5407 {
					tmp5405 := MakeNative(func(__e *ControlFlow) {
						Y := __e.Get(1)
						_ = Y
						__e.TailApply(PrimNS2Value(symshen_4process_1applications), Y, V2931)
						return
					}, 1)

					tmp5406 := Call(__e, PrimNS2Value(symmap), tmp5405, V2930)

					__e.TailApply(PrimNS2Value(symshen_4process_1application), tmp5406, V2931)
					return

				} else {
					__e.Return(V2930)
					return
				}

			}

		}

	}, 2)

	tmp5415 := Call(__e, PrimNS2Value(symdef), symshen_4process_1applications, tmp5401)

	_ = tmp5415

	tmp5416 := MakeNative(func(__e *ControlFlow) {
		V2934 := __e.Get(1)
		_ = V2934
		tmp5426 := PrimEqual(symdefine, V2934)

		if True == tmp5426 {
			__e.Return(True)
			return
		} else {
			tmp5425 := PrimEqual(symdefun, V2934)

			if True == tmp5425 {
				__e.Return(True)
				return
			} else {
				tmp5424 := PrimEqual(symsynonyms, V2934)

				if True == tmp5424 {
					__e.Return(True)
					return
				} else {
					tmp5423 := Call(__e, PrimNS2Value(symshen_4special_2), V2934)

					if True == tmp5423 {
						__e.Return(True)
						return
					} else {
						tmp5422 := Call(__e, PrimNS2Value(symshen_4extraspecial_2), V2934)

						if True == tmp5422 {
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

	tmp5427 := Call(__e, PrimNS2Value(symdef), symshen_4non_1application_2, tmp5416)

	_ = tmp5427

	tmp5428 := MakeNative(func(__e *ControlFlow) {
		V2939 := __e.Get(1)
		_ = V2939
		V2940 := __e.Get(2)
		_ = V2940
		V2941 := __e.Get(3)
		_ = V2941
		tmp5670 := PrimEqual(symlambda, V2939)

		var ifres5648 Obj

		if True == tmp5670 {
			tmp5669 := PrimIsPair(V2940)

			var ifres5650 Obj

			if True == tmp5669 {
				tmp5667 := PrimHead(V2940)

				tmp5668 := PrimEqual(symlambda, tmp5667)

				var ifres5652 Obj

				if True == tmp5668 {
					tmp5665 := PrimTail(V2940)

					tmp5666 := PrimIsPair(tmp5665)

					var ifres5654 Obj

					if True == tmp5666 {
						tmp5662 := PrimTail(V2940)

						tmp5663 := PrimTail(tmp5662)

						tmp5664 := PrimIsPair(tmp5663)

						var ifres5656 Obj

						if True == tmp5664 {
							tmp5658 := PrimTail(V2940)

							tmp5659 := PrimTail(tmp5658)

							tmp5660 := PrimTail(tmp5659)

							tmp5661 := PrimEqual(Nil, tmp5660)

							var ifres5657 Obj

							if True == tmp5661 {
								ifres5657 = True

							} else {
								ifres5657 = False

							}

							ifres5656 = ifres5657

						} else {
							ifres5656 = False

						}

						var ifres5655 Obj

						if True == ifres5656 {
							ifres5655 = True

						} else {
							ifres5655 = False

						}

						ifres5654 = ifres5655

					} else {
						ifres5654 = False

					}

					var ifres5653 Obj

					if True == ifres5654 {
						ifres5653 = True

					} else {
						ifres5653 = False

					}

					ifres5652 = ifres5653

				} else {
					ifres5652 = False

				}

				var ifres5651 Obj

				if True == ifres5652 {
					ifres5651 = True

				} else {
					ifres5651 = False

				}

				ifres5650 = ifres5651

			} else {
				ifres5650 = False

			}

			var ifres5649 Obj

			if True == ifres5650 {
				ifres5649 = True

			} else {
				ifres5649 = False

			}

			ifres5648 = ifres5649

		} else {
			ifres5648 = False

		}

		if True == ifres5648 {
			tmp5640 := PrimTail(V2940)

			tmp5641 := PrimHead(tmp5640)

			tmp5642 := PrimTail(V2940)

			tmp5643 := PrimTail(tmp5642)

			tmp5644 := PrimHead(tmp5643)

			tmp5645 := Call(__e, PrimNS2Value(symshen_4process_1applications), tmp5644, V2941)

			tmp5646 := PrimCons(tmp5645, Nil)

			tmp5647 := PrimCons(tmp5641, tmp5646)

			__e.Return(PrimCons(symlambda, tmp5647))
			return

		} else {
			tmp5639 := PrimEqual(symlet, V2939)

			var ifres5610 Obj

			if True == tmp5639 {
				tmp5638 := PrimIsPair(V2940)

				var ifres5612 Obj

				if True == tmp5638 {
					tmp5636 := PrimHead(V2940)

					tmp5637 := PrimEqual(symlet, tmp5636)

					var ifres5614 Obj

					if True == tmp5637 {
						tmp5634 := PrimTail(V2940)

						tmp5635 := PrimIsPair(tmp5634)

						var ifres5616 Obj

						if True == tmp5635 {
							tmp5631 := PrimTail(V2940)

							tmp5632 := PrimTail(tmp5631)

							tmp5633 := PrimIsPair(tmp5632)

							var ifres5618 Obj

							if True == tmp5633 {
								tmp5627 := PrimTail(V2940)

								tmp5628 := PrimTail(tmp5627)

								tmp5629 := PrimTail(tmp5628)

								tmp5630 := PrimIsPair(tmp5629)

								var ifres5620 Obj

								if True == tmp5630 {
									tmp5622 := PrimTail(V2940)

									tmp5623 := PrimTail(tmp5622)

									tmp5624 := PrimTail(tmp5623)

									tmp5625 := PrimTail(tmp5624)

									tmp5626 := PrimEqual(Nil, tmp5625)

									var ifres5621 Obj

									if True == tmp5626 {
										ifres5621 = True

									} else {
										ifres5621 = False

									}

									ifres5620 = ifres5621

								} else {
									ifres5620 = False

								}

								var ifres5619 Obj

								if True == ifres5620 {
									ifres5619 = True

								} else {
									ifres5619 = False

								}

								ifres5618 = ifres5619

							} else {
								ifres5618 = False

							}

							var ifres5617 Obj

							if True == ifres5618 {
								ifres5617 = True

							} else {
								ifres5617 = False

							}

							ifres5616 = ifres5617

						} else {
							ifres5616 = False

						}

						var ifres5615 Obj

						if True == ifres5616 {
							ifres5615 = True

						} else {
							ifres5615 = False

						}

						ifres5614 = ifres5615

					} else {
						ifres5614 = False

					}

					var ifres5613 Obj

					if True == ifres5614 {
						ifres5613 = True

					} else {
						ifres5613 = False

					}

					ifres5612 = ifres5613

				} else {
					ifres5612 = False

				}

				var ifres5611 Obj

				if True == ifres5612 {
					ifres5611 = True

				} else {
					ifres5611 = False

				}

				ifres5610 = ifres5611

			} else {
				ifres5610 = False

			}

			if True == ifres5610 {
				tmp5596 := PrimTail(V2940)

				tmp5597 := PrimHead(tmp5596)

				tmp5598 := PrimTail(V2940)

				tmp5599 := PrimTail(tmp5598)

				tmp5600 := PrimHead(tmp5599)

				tmp5601 := Call(__e, PrimNS2Value(symshen_4process_1applications), tmp5600, V2941)

				tmp5602 := PrimTail(V2940)

				tmp5603 := PrimTail(tmp5602)

				tmp5604 := PrimTail(tmp5603)

				tmp5605 := PrimHead(tmp5604)

				tmp5606 := Call(__e, PrimNS2Value(symshen_4process_1applications), tmp5605, V2941)

				tmp5607 := PrimCons(tmp5606, Nil)

				tmp5608 := PrimCons(tmp5601, tmp5607)

				tmp5609 := PrimCons(tmp5597, tmp5608)

				__e.Return(PrimCons(symlet, tmp5609))
				return

			} else {
				tmp5595 := PrimEqual(symdefun, V2939)

				var ifres5566 Obj

				if True == tmp5595 {
					tmp5594 := PrimIsPair(V2940)

					var ifres5568 Obj

					if True == tmp5594 {
						tmp5592 := PrimHead(V2940)

						tmp5593 := PrimEqual(symdefun, tmp5592)

						var ifres5570 Obj

						if True == tmp5593 {
							tmp5590 := PrimTail(V2940)

							tmp5591 := PrimIsPair(tmp5590)

							var ifres5572 Obj

							if True == tmp5591 {
								tmp5587 := PrimTail(V2940)

								tmp5588 := PrimTail(tmp5587)

								tmp5589 := PrimIsPair(tmp5588)

								var ifres5574 Obj

								if True == tmp5589 {
									tmp5583 := PrimTail(V2940)

									tmp5584 := PrimTail(tmp5583)

									tmp5585 := PrimTail(tmp5584)

									tmp5586 := PrimIsPair(tmp5585)

									var ifres5576 Obj

									if True == tmp5586 {
										tmp5578 := PrimTail(V2940)

										tmp5579 := PrimTail(tmp5578)

										tmp5580 := PrimTail(tmp5579)

										tmp5581 := PrimTail(tmp5580)

										tmp5582 := PrimEqual(Nil, tmp5581)

										var ifres5577 Obj

										if True == tmp5582 {
											ifres5577 = True

										} else {
											ifres5577 = False

										}

										ifres5576 = ifres5577

									} else {
										ifres5576 = False

									}

									var ifres5575 Obj

									if True == ifres5576 {
										ifres5575 = True

									} else {
										ifres5575 = False

									}

									ifres5574 = ifres5575

								} else {
									ifres5574 = False

								}

								var ifres5573 Obj

								if True == ifres5574 {
									ifres5573 = True

								} else {
									ifres5573 = False

								}

								ifres5572 = ifres5573

							} else {
								ifres5572 = False

							}

							var ifres5571 Obj

							if True == ifres5572 {
								ifres5571 = True

							} else {
								ifres5571 = False

							}

							ifres5570 = ifres5571

						} else {
							ifres5570 = False

						}

						var ifres5569 Obj

						if True == ifres5570 {
							ifres5569 = True

						} else {
							ifres5569 = False

						}

						ifres5568 = ifres5569

					} else {
						ifres5568 = False

					}

					var ifres5567 Obj

					if True == ifres5568 {
						ifres5567 = True

					} else {
						ifres5567 = False

					}

					ifres5566 = ifres5567

				} else {
					ifres5566 = False

				}

				if True == ifres5566 {
					__e.Return(V2940)
					return
				} else {
					tmp5565 := PrimEqual(symdefine, V2939)

					var ifres5543 Obj

					if True == tmp5565 {
						tmp5564 := PrimIsPair(V2940)

						var ifres5545 Obj

						if True == tmp5564 {
							tmp5562 := PrimHead(V2940)

							tmp5563 := PrimEqual(symdefine, tmp5562)

							var ifres5547 Obj

							if True == tmp5563 {
								tmp5560 := PrimTail(V2940)

								tmp5561 := PrimIsPair(tmp5560)

								var ifres5549 Obj

								if True == tmp5561 {
									tmp5557 := PrimTail(V2940)

									tmp5558 := PrimTail(tmp5557)

									tmp5559 := PrimIsPair(tmp5558)

									var ifres5551 Obj

									if True == tmp5559 {
										tmp5553 := PrimTail(V2940)

										tmp5554 := PrimTail(tmp5553)

										tmp5555 := PrimHead(tmp5554)

										tmp5556 := PrimEqual(sym_i, tmp5555)

										var ifres5552 Obj

										if True == tmp5556 {
											ifres5552 = True

										} else {
											ifres5552 = False

										}

										ifres5551 = ifres5552

									} else {
										ifres5551 = False

									}

									var ifres5550 Obj

									if True == ifres5551 {
										ifres5550 = True

									} else {
										ifres5550 = False

									}

									ifres5549 = ifres5550

								} else {
									ifres5549 = False

								}

								var ifres5548 Obj

								if True == ifres5549 {
									ifres5548 = True

								} else {
									ifres5548 = False

								}

								ifres5547 = ifres5548

							} else {
								ifres5547 = False

							}

							var ifres5546 Obj

							if True == ifres5547 {
								ifres5546 = True

							} else {
								ifres5546 = False

							}

							ifres5545 = ifres5546

						} else {
							ifres5545 = False

						}

						var ifres5544 Obj

						if True == ifres5545 {
							ifres5544 = True

						} else {
							ifres5544 = False

						}

						ifres5543 = ifres5544

					} else {
						ifres5543 = False

					}

					if True == ifres5543 {
						tmp5533 := PrimTail(V2940)

						tmp5534 := PrimHead(tmp5533)

						tmp5535 := PrimTail(V2940)

						tmp5536 := PrimHead(tmp5535)

						tmp5537 := PrimTail(V2940)

						tmp5538 := PrimTail(tmp5537)

						tmp5539 := PrimTail(tmp5538)

						tmp5540 := Call(__e, PrimNS2Value(symshen_4process_1after_1type), tmp5536, tmp5539, V2941)

						tmp5541 := PrimCons(sym_i, tmp5540)

						tmp5542 := PrimCons(tmp5534, tmp5541)

						__e.Return(PrimCons(symdefine, tmp5542))
						return

					} else {
						tmp5532 := PrimEqual(symdefine, V2939)

						var ifres5521 Obj

						if True == tmp5532 {
							tmp5531 := PrimIsPair(V2940)

							var ifres5523 Obj

							if True == tmp5531 {
								tmp5529 := PrimHead(V2940)

								tmp5530 := PrimEqual(symdefine, tmp5529)

								var ifres5525 Obj

								if True == tmp5530 {
									tmp5527 := PrimTail(V2940)

									tmp5528 := PrimIsPair(tmp5527)

									var ifres5526 Obj

									if True == tmp5528 {
										ifres5526 = True

									} else {
										ifres5526 = False

									}

									ifres5525 = ifres5526

								} else {
									ifres5525 = False

								}

								var ifres5524 Obj

								if True == ifres5525 {
									ifres5524 = True

								} else {
									ifres5524 = False

								}

								ifres5523 = ifres5524

							} else {
								ifres5523 = False

							}

							var ifres5522 Obj

							if True == ifres5523 {
								ifres5522 = True

							} else {
								ifres5522 = False

							}

							ifres5521 = ifres5522

						} else {
							ifres5521 = False

						}

						if True == ifres5521 {
							tmp5514 := PrimTail(V2940)

							tmp5515 := PrimHead(tmp5514)

							tmp5516 := MakeNative(func(__e *ControlFlow) {
								Y := __e.Get(1)
								_ = Y
								__e.TailApply(PrimNS2Value(symshen_4process_1applications), Y, V2941)
								return
							}, 1)

							tmp5517 := PrimTail(V2940)

							tmp5518 := PrimTail(tmp5517)

							tmp5519 := Call(__e, PrimNS2Value(symmap), tmp5516, tmp5518)

							tmp5520 := PrimCons(tmp5515, tmp5519)

							__e.Return(PrimCons(symdefine, tmp5520))
							return

						} else {
							tmp5513 := PrimEqual(symsynonyms, V2939)

							if True == tmp5513 {
								__e.Return(PrimCons(symsynonyms, V2940))
								return
							} else {
								tmp5512 := PrimEqual(symtype, V2939)

								var ifres5490 Obj

								if True == tmp5512 {
									tmp5511 := PrimIsPair(V2940)

									var ifres5492 Obj

									if True == tmp5511 {
										tmp5509 := PrimHead(V2940)

										tmp5510 := PrimEqual(symtype, tmp5509)

										var ifres5494 Obj

										if True == tmp5510 {
											tmp5507 := PrimTail(V2940)

											tmp5508 := PrimIsPair(tmp5507)

											var ifres5496 Obj

											if True == tmp5508 {
												tmp5504 := PrimTail(V2940)

												tmp5505 := PrimTail(tmp5504)

												tmp5506 := PrimIsPair(tmp5505)

												var ifres5498 Obj

												if True == tmp5506 {
													tmp5500 := PrimTail(V2940)

													tmp5501 := PrimTail(tmp5500)

													tmp5502 := PrimTail(tmp5501)

													tmp5503 := PrimEqual(Nil, tmp5502)

													var ifres5499 Obj

													if True == tmp5503 {
														ifres5499 = True

													} else {
														ifres5499 = False

													}

													ifres5498 = ifres5499

												} else {
													ifres5498 = False

												}

												var ifres5497 Obj

												if True == ifres5498 {
													ifres5497 = True

												} else {
													ifres5497 = False

												}

												ifres5496 = ifres5497

											} else {
												ifres5496 = False

											}

											var ifres5495 Obj

											if True == ifres5496 {
												ifres5495 = True

											} else {
												ifres5495 = False

											}

											ifres5494 = ifres5495

										} else {
											ifres5494 = False

										}

										var ifres5493 Obj

										if True == ifres5494 {
											ifres5493 = True

										} else {
											ifres5493 = False

										}

										ifres5492 = ifres5493

									} else {
										ifres5492 = False

									}

									var ifres5491 Obj

									if True == ifres5492 {
										ifres5491 = True

									} else {
										ifres5491 = False

									}

									ifres5490 = ifres5491

								} else {
									ifres5490 = False

								}

								if True == ifres5490 {
									tmp5484 := PrimTail(V2940)

									tmp5485 := PrimHead(tmp5484)

									tmp5486 := Call(__e, PrimNS2Value(symshen_4process_1applications), tmp5485, V2941)

									tmp5487 := PrimTail(V2940)

									tmp5488 := PrimTail(tmp5487)

									tmp5489 := PrimCons(tmp5486, tmp5488)

									__e.Return(PrimCons(symtype, tmp5489))
									return

								} else {
									tmp5483 := PrimEqual(syminput_7, V2939)

									var ifres5461 Obj

									if True == tmp5483 {
										tmp5482 := PrimIsPair(V2940)

										var ifres5463 Obj

										if True == tmp5482 {
											tmp5480 := PrimHead(V2940)

											tmp5481 := PrimEqual(syminput_7, tmp5480)

											var ifres5465 Obj

											if True == tmp5481 {
												tmp5478 := PrimTail(V2940)

												tmp5479 := PrimIsPair(tmp5478)

												var ifres5467 Obj

												if True == tmp5479 {
													tmp5475 := PrimTail(V2940)

													tmp5476 := PrimTail(tmp5475)

													tmp5477 := PrimIsPair(tmp5476)

													var ifres5469 Obj

													if True == tmp5477 {
														tmp5471 := PrimTail(V2940)

														tmp5472 := PrimTail(tmp5471)

														tmp5473 := PrimTail(tmp5472)

														tmp5474 := PrimEqual(Nil, tmp5473)

														var ifres5470 Obj

														if True == tmp5474 {
															ifres5470 = True

														} else {
															ifres5470 = False

														}

														ifres5469 = ifres5470

													} else {
														ifres5469 = False

													}

													var ifres5468 Obj

													if True == ifres5469 {
														ifres5468 = True

													} else {
														ifres5468 = False

													}

													ifres5467 = ifres5468

												} else {
													ifres5467 = False

												}

												var ifres5466 Obj

												if True == ifres5467 {
													ifres5466 = True

												} else {
													ifres5466 = False

												}

												ifres5465 = ifres5466

											} else {
												ifres5465 = False

											}

											var ifres5464 Obj

											if True == ifres5465 {
												ifres5464 = True

											} else {
												ifres5464 = False

											}

											ifres5463 = ifres5464

										} else {
											ifres5463 = False

										}

										var ifres5462 Obj

										if True == ifres5463 {
											ifres5462 = True

										} else {
											ifres5462 = False

										}

										ifres5461 = ifres5462

									} else {
										ifres5461 = False

									}

									if True == ifres5461 {
										tmp5453 := PrimTail(V2940)

										tmp5454 := PrimHead(tmp5453)

										tmp5455 := PrimTail(V2940)

										tmp5456 := PrimTail(tmp5455)

										tmp5457 := PrimHead(tmp5456)

										tmp5458 := Call(__e, PrimNS2Value(symshen_4process_1applications), tmp5457, V2941)

										tmp5459 := PrimCons(tmp5458, Nil)

										tmp5460 := PrimCons(tmp5454, tmp5459)

										__e.Return(PrimCons(syminput_7, tmp5460))
										return

									} else {
										tmp5452 := PrimIsPair(V2940)

										var ifres5448 Obj

										if True == tmp5452 {
											tmp5450 := PrimHead(V2940)

											tmp5451 := Call(__e, PrimNS2Value(symshen_4special_2), tmp5450)

											var ifres5449 Obj

											if True == tmp5451 {
												ifres5449 = True

											} else {
												ifres5449 = False

											}

											ifres5448 = ifres5449

										} else {
											ifres5448 = False

										}

										if True == ifres5448 {
											tmp5444 := PrimHead(V2940)

											tmp5445 := MakeNative(func(__e *ControlFlow) {
												Y := __e.Get(1)
												_ = Y
												__e.TailApply(PrimNS2Value(symshen_4process_1applications), Y, V2941)
												return
											}, 1)

											tmp5446 := PrimTail(V2940)

											tmp5447 := Call(__e, PrimNS2Value(symmap), tmp5445, tmp5446)

											__e.Return(PrimCons(tmp5444, tmp5447))
											return

										} else {
											tmp5443 := PrimIsPair(V2940)

											var ifres5439 Obj

											if True == tmp5443 {
												tmp5441 := PrimHead(V2940)

												tmp5442 := Call(__e, PrimNS2Value(symshen_4extraspecial_2), tmp5441)

												var ifres5440 Obj

												if True == tmp5442 {
													ifres5440 = True

												} else {
													ifres5440 = False

												}

												ifres5439 = ifres5440

											} else {
												ifres5439 = False

											}

											if True == ifres5439 {
												__e.Return(V2940)
												return
											} else {
												__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4special_1case)
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

	tmp5671 := Call(__e, PrimNS2Value(symdef), symshen_4special_1case, tmp5428)

	_ = tmp5671

	tmp5672 := MakeNative(func(__e *ControlFlow) {
		V2944 := __e.Get(1)
		_ = V2944
		V2945 := __e.Get(2)
		_ = V2945
		V2946 := __e.Get(3)
		_ = V2946
		tmp5688 := PrimIsPair(V2945)

		var ifres5684 Obj

		if True == tmp5688 {
			tmp5686 := PrimHead(V2945)

			tmp5687 := PrimEqual(sym_j, tmp5686)

			var ifres5685 Obj

			if True == tmp5687 {
				ifres5685 = True

			} else {
				ifres5685 = False

			}

			ifres5684 = ifres5685

		} else {
			ifres5684 = False

		}

		if True == ifres5684 {
			tmp5681 := MakeNative(func(__e *ControlFlow) {
				Y := __e.Get(1)
				_ = Y
				__e.TailApply(PrimNS2Value(symshen_4process_1applications), Y, V2946)
				return
			}, 1)

			tmp5682 := PrimTail(V2945)

			tmp5683 := Call(__e, PrimNS2Value(symmap), tmp5681, tmp5682)

			__e.Return(PrimCons(sym_j, tmp5683))
			return

		} else {
			tmp5680 := PrimIsPair(V2945)

			if True == tmp5680 {
				tmp5677 := PrimHead(V2945)

				tmp5678 := PrimTail(V2945)

				tmp5679 := Call(__e, PrimNS2Value(symshen_4process_1after_1type), V2944, tmp5678, V2946)

				__e.Return(PrimCons(tmp5677, tmp5679))
				return

			} else {
				tmp5675 := Call(__e, PrimNS2Value(symshen_4app), V2944, MakeString("\n"), symshen_4a)

				tmp5676 := PrimStringConcat(MakeString("missing } in "), tmp5675)

				__e.Return(PrimSimpleError(tmp5676))
				return

			}

		}

	}, 3)

	tmp5689 := Call(__e, PrimNS2Value(symdef), symshen_4process_1after_1type, tmp5672)

	_ = tmp5689

	tmp5690 := MakeNative(func(__e *ControlFlow) {
		V2947 := __e.Get(1)
		_ = V2947
		V2948 := __e.Get(2)
		_ = V2948
		tmp5728 := PrimIsPair(V2947)

		if True == tmp5728 {
			tmp5692 := MakeNative(func(__e *ControlFlow) {
				ArityF := __e.Get(1)
				_ = ArityF
				tmp5693 := MakeNative(func(__e *ControlFlow) {
					N := __e.Get(1)
					_ = N
					tmp5723 := Call(__e, PrimNS2Value(symelement_2), V2947, V2948)

					if True == tmp5723 {
						__e.Return(V2947)
						return
					} else {
						tmp5721 := PrimHead(V2947)

						tmp5722 := Call(__e, PrimNS2Value(symshen_4shen_1call_2), tmp5721)

						if True == tmp5722 {
							__e.Return(V2947)
							return
						} else {
							tmp5720 := Call(__e, PrimNS2Value(symshen_4fn_1call_2), V2947)

							if True == tmp5720 {
								__e.TailApply(PrimNS2Value(symshen_4fn_1call), V2947)
								return
							} else {
								tmp5719 := Call(__e, PrimNS2Value(symshen_4zero_1place_2), V2947)

								if True == tmp5719 {
									__e.Return(V2947)
									return
								} else {
									tmp5717 := PrimHead(V2947)

									tmp5718 := Call(__e, PrimNS2Value(symshen_4undefined_1f_2), tmp5717, ArityF)

									if True == tmp5718 {
										tmp5712 := PrimHead(V2947)

										tmp5713 := PrimCons(tmp5712, Nil)

										tmp5714 := PrimCons(symfn, tmp5713)

										tmp5715 := PrimTail(V2947)

										tmp5716 := PrimCons(tmp5714, tmp5715)

										__e.TailApply(PrimNS2Value(symshen_4simple_1curry), tmp5716)
										return

									} else {
										tmp5710 := PrimHead(V2947)

										tmp5711 := PrimIsVariable(tmp5710)

										if True == tmp5711 {
											__e.TailApply(PrimNS2Value(symshen_4simple_1curry), V2947)
											return
										} else {
											tmp5708 := PrimHead(V2947)

											tmp5709 := Call(__e, PrimNS2Value(symshen_4application_2), tmp5708)

											if True == tmp5709 {
												__e.TailApply(PrimNS2Value(symshen_4simple_1curry), V2947)
												return
											} else {
												tmp5706 := PrimHead(V2947)

												tmp5707 := Call(__e, PrimNS2Value(symshen_4partial_1application_d_2), tmp5706, ArityF, N)

												if True == tmp5707 {
													tmp5705 := PrimNumberSubtract(ArityF, N)

													__e.TailApply(PrimNS2Value(symshen_4lambda_1function), V2947, tmp5705)
													return

												} else {
													tmp5703 := PrimHead(V2947)

													tmp5704 := Call(__e, PrimNS2Value(symshen_4overapplication_2), tmp5703, ArityF, N)

													if True == tmp5704 {
														__e.TailApply(PrimNS2Value(symshen_4simple_1curry), V2947)
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

				tmp5724 := PrimTail(V2947)

				tmp5725 := Call(__e, PrimNS2Value(symlength), tmp5724)

				__e.TailApply(tmp5693, tmp5725)
				return

			}, 1)

			tmp5726 := PrimHead(V2947)

			tmp5727 := Call(__e, PrimNS2Value(symarity), tmp5726)

			__e.TailApply(tmp5692, tmp5727)
			return

		} else {
			__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4process_1application)
			return
		}

	}, 2)

	tmp5729 := Call(__e, PrimNS2Value(symdef), symshen_4process_1application, tmp5690)

	_ = tmp5729

	tmp5730 := MakeNative(func(__e *ControlFlow) {
		V2951 := __e.Get(1)
		_ = V2951
		tmp5736 := PrimIsPair(V2951)

		var ifres5732 Obj

		if True == tmp5736 {
			tmp5734 := PrimTail(V2951)

			tmp5735 := PrimEqual(Nil, tmp5734)

			var ifres5733 Obj

			if True == tmp5735 {
				ifres5733 = True

			} else {
				ifres5733 = False

			}

			ifres5732 = ifres5733

		} else {
			ifres5732 = False

		}

		if True == ifres5732 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp5737 := Call(__e, PrimNS2Value(symdef), symshen_4zero_1place_2, tmp5730)

	_ = tmp5737

	tmp5738 := MakeNative(func(__e *ControlFlow) {
		V2952 := __e.Get(1)
		_ = V2952
		tmp5743 := PrimIsSymbol(V2952)

		if True == tmp5743 {
			tmp5741 := PrimStr(V2952)

			tmp5742 := Call(__e, PrimNS2Value(symshen_4internal_1to_1shen_2), tmp5741)

			if True == tmp5742 {
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

	tmp5744 := Call(__e, PrimNS2Value(symdef), symshen_4shen_1call_2, tmp5738)

	_ = tmp5744

	tmp5745 := MakeNative(func(__e *ControlFlow) {
		V2957 := __e.Get(1)
		_ = V2957
		tmp5799 := Call(__e, PrimNS2Value(symshen_4_7string_2), V2957)

		var ifres5747 Obj

		if True == tmp5799 {
			tmp5797 := Call(__e, PrimNS2Value(symhdstr), V2957)

			tmp5798 := PrimEqual(MakeString("s"), tmp5797)

			var ifres5749 Obj

			if True == tmp5798 {
				tmp5795 := PrimTailString(V2957)

				tmp5796 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5795)

				var ifres5751 Obj

				if True == tmp5796 {
					tmp5792 := PrimTailString(V2957)

					tmp5793 := Call(__e, PrimNS2Value(symhdstr), tmp5792)

					tmp5794 := PrimEqual(MakeString("h"), tmp5793)

					var ifres5753 Obj

					if True == tmp5794 {
						tmp5789 := PrimTailString(V2957)

						tmp5790 := PrimTailString(tmp5789)

						tmp5791 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5790)

						var ifres5755 Obj

						if True == tmp5791 {
							tmp5785 := PrimTailString(V2957)

							tmp5786 := PrimTailString(tmp5785)

							tmp5787 := Call(__e, PrimNS2Value(symhdstr), tmp5786)

							tmp5788 := PrimEqual(MakeString("e"), tmp5787)

							var ifres5757 Obj

							if True == tmp5788 {
								tmp5781 := PrimTailString(V2957)

								tmp5782 := PrimTailString(tmp5781)

								tmp5783 := PrimTailString(tmp5782)

								tmp5784 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5783)

								var ifres5759 Obj

								if True == tmp5784 {
									tmp5776 := PrimTailString(V2957)

									tmp5777 := PrimTailString(tmp5776)

									tmp5778 := PrimTailString(tmp5777)

									tmp5779 := Call(__e, PrimNS2Value(symhdstr), tmp5778)

									tmp5780 := PrimEqual(MakeString("n"), tmp5779)

									var ifres5761 Obj

									if True == tmp5780 {
										tmp5771 := PrimTailString(V2957)

										tmp5772 := PrimTailString(tmp5771)

										tmp5773 := PrimTailString(tmp5772)

										tmp5774 := PrimTailString(tmp5773)

										tmp5775 := Call(__e, PrimNS2Value(symshen_4_7string_2), tmp5774)

										var ifres5763 Obj

										if True == tmp5775 {
											tmp5765 := PrimTailString(V2957)

											tmp5766 := PrimTailString(tmp5765)

											tmp5767 := PrimTailString(tmp5766)

											tmp5768 := PrimTailString(tmp5767)

											tmp5769 := Call(__e, PrimNS2Value(symhdstr), tmp5768)

											tmp5770 := PrimEqual(MakeString("."), tmp5769)

											var ifres5764 Obj

											if True == tmp5770 {
												ifres5764 = True

											} else {
												ifres5764 = False

											}

											ifres5763 = ifres5764

										} else {
											ifres5763 = False

										}

										var ifres5762 Obj

										if True == ifres5763 {
											ifres5762 = True

										} else {
											ifres5762 = False

										}

										ifres5761 = ifres5762

									} else {
										ifres5761 = False

									}

									var ifres5760 Obj

									if True == ifres5761 {
										ifres5760 = True

									} else {
										ifres5760 = False

									}

									ifres5759 = ifres5760

								} else {
									ifres5759 = False

								}

								var ifres5758 Obj

								if True == ifres5759 {
									ifres5758 = True

								} else {
									ifres5758 = False

								}

								ifres5757 = ifres5758

							} else {
								ifres5757 = False

							}

							var ifres5756 Obj

							if True == ifres5757 {
								ifres5756 = True

							} else {
								ifres5756 = False

							}

							ifres5755 = ifres5756

						} else {
							ifres5755 = False

						}

						var ifres5754 Obj

						if True == ifres5755 {
							ifres5754 = True

						} else {
							ifres5754 = False

						}

						ifres5753 = ifres5754

					} else {
						ifres5753 = False

					}

					var ifres5752 Obj

					if True == ifres5753 {
						ifres5752 = True

					} else {
						ifres5752 = False

					}

					ifres5751 = ifres5752

				} else {
					ifres5751 = False

				}

				var ifres5750 Obj

				if True == ifres5751 {
					ifres5750 = True

				} else {
					ifres5750 = False

				}

				ifres5749 = ifres5750

			} else {
				ifres5749 = False

			}

			var ifres5748 Obj

			if True == ifres5749 {
				ifres5748 = True

			} else {
				ifres5748 = False

			}

			ifres5747 = ifres5748

		} else {
			ifres5747 = False

		}

		if True == ifres5747 {
			__e.Return(True)
			return
		} else {
			__e.Return(False)
			return
		}

	}, 1)

	tmp5800 := Call(__e, PrimNS2Value(symdef), symshen_4internal_1to_1shen_2, tmp5745)

	_ = tmp5800

	tmp5801 := MakeNative(func(__e *ControlFlow) {
		V2958 := __e.Get(1)
		_ = V2958
		__e.Return(PrimIsPair(V2958))
		return
	}, 1)

	tmp5802 := Call(__e, PrimNS2Value(symdef), symshen_4application_2, tmp5801)

	_ = tmp5802

	tmp5803 := MakeNative(func(__e *ControlFlow) {
		V2963 := __e.Get(1)
		_ = V2963
		V2964 := __e.Get(2)
		_ = V2964
		tmp5811 := PrimEqual(MakeNumber(-1), V2964)

		if True == tmp5811 {
			tmp5810 := Call(__e, PrimNS2Value(symshen_4lowercase_1symbol_2), V2963)

			if True == tmp5810 {
				tmp5807 := Call(__e, PrimNS2Value(symexternal), symshen)

				tmp5808 := Call(__e, PrimNS2Value(symelement_2), V2963, tmp5807)

				tmp5809 := PrimNot(tmp5808)

				if True == tmp5809 {
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

	tmp5812 := Call(__e, PrimNS2Value(symdef), symshen_4undefined_1f_2, tmp5803)

	_ = tmp5812

	tmp5813 := MakeNative(func(__e *ControlFlow) {
		V2965 := __e.Get(1)
		_ = V2965
		tmp5818 := PrimIsSymbol(V2965)

		if True == tmp5818 {
			tmp5816 := PrimIsVariable(V2965)

			tmp5817 := PrimNot(tmp5816)

			if True == tmp5817 {
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

	tmp5819 := Call(__e, PrimNS2Value(symdef), symshen_4lowercase_1symbol_2, tmp5813)

	_ = tmp5819

	tmp5820 := MakeNative(func(__e *ControlFlow) {
		V2966 := __e.Get(1)
		_ = V2966
		tmp5850 := PrimIsPair(V2966)

		var ifres5841 Obj

		if True == tmp5850 {
			tmp5848 := PrimTail(V2966)

			tmp5849 := PrimIsPair(tmp5848)

			var ifres5843 Obj

			if True == tmp5849 {
				tmp5845 := PrimTail(V2966)

				tmp5846 := PrimTail(tmp5845)

				tmp5847 := PrimEqual(Nil, tmp5846)

				var ifres5844 Obj

				if True == tmp5847 {
					ifres5844 = True

				} else {
					ifres5844 = False

				}

				ifres5843 = ifres5844

			} else {
				ifres5843 = False

			}

			var ifres5842 Obj

			if True == ifres5843 {
				ifres5842 = True

			} else {
				ifres5842 = False

			}

			ifres5841 = ifres5842

		} else {
			ifres5841 = False

		}

		if True == ifres5841 {
			__e.Return(V2966)
			return
		} else {
			tmp5840 := PrimIsPair(V2966)

			var ifres5831 Obj

			if True == tmp5840 {
				tmp5838 := PrimTail(V2966)

				tmp5839 := PrimIsPair(tmp5838)

				var ifres5833 Obj

				if True == tmp5839 {
					tmp5835 := PrimTail(V2966)

					tmp5836 := PrimTail(tmp5835)

					tmp5837 := PrimIsPair(tmp5836)

					var ifres5834 Obj

					if True == tmp5837 {
						ifres5834 = True

					} else {
						ifres5834 = False

					}

					ifres5833 = ifres5834

				} else {
					ifres5833 = False

				}

				var ifres5832 Obj

				if True == ifres5833 {
					ifres5832 = True

				} else {
					ifres5832 = False

				}

				ifres5831 = ifres5832

			} else {
				ifres5831 = False

			}

			if True == ifres5831 {
				tmp5823 := PrimHead(V2966)

				tmp5824 := PrimTail(V2966)

				tmp5825 := PrimHead(tmp5824)

				tmp5826 := PrimCons(tmp5825, Nil)

				tmp5827 := PrimCons(tmp5823, tmp5826)

				tmp5828 := PrimTail(V2966)

				tmp5829 := PrimTail(tmp5828)

				tmp5830 := PrimCons(tmp5827, tmp5829)

				__e.TailApply(PrimNS2Value(symshen_4simple_1curry), tmp5830)
				return

			} else {
				__e.Return(V2966)
				return
			}

		}

	}, 1)

	tmp5851 := Call(__e, PrimNS2Value(symdef), symshen_4simple_1curry, tmp5820)

	_ = tmp5851

	tmp5852 := MakeNative(func(__e *ControlFlow) {
		V2967 := __e.Get(1)
		_ = V2967
		__e.TailApply(PrimNS2Value(symfn), V2967)
		return
	}, 1)

	tmp5853 := Call(__e, PrimNS2Value(symdef), symfunction, tmp5852)

	_ = tmp5853

	tmp5854 := MakeNative(func(__e *ControlFlow) {
		V2968 := __e.Get(1)
		_ = V2968
		tmp5855 := MakeNative(func(__e *ControlFlow) {
			LookUp := __e.Get(1)
			_ = LookUp
			tmp5859 := Call(__e, PrimNS2Value(symempty_2), LookUp)

			if True == tmp5859 {
				tmp5857 := Call(__e, PrimNS2Value(symshen_4app), V2968, MakeString(" is undefined\n"), symshen_4a)

				tmp5858 := PrimStringConcat(MakeString("fn: "), tmp5857)

				__e.Return(PrimSimpleError(tmp5858))
				return

			} else {
				__e.Return(PrimTail(LookUp))
				return
			}

		}, 1)

		tmp5860 := PrimNS3Value(symshen_4_dlambdatable_d)

		tmp5861 := Call(__e, PrimNS2Value(symassoc), V2968, tmp5860)

		__e.TailApply(tmp5855, tmp5861)
		return

	}, 1)

	tmp5862 := Call(__e, PrimNS2Value(symdef), symfn, tmp5854)

	_ = tmp5862

	tmp5863 := MakeNative(func(__e *ControlFlow) {
		V2971 := __e.Get(1)
		_ = V2971
		tmp5893 := PrimIsPair(V2971)

		var ifres5880 Obj

		if True == tmp5893 {
			tmp5891 := PrimHead(V2971)

			tmp5892 := PrimEqual(symfn, tmp5891)

			var ifres5882 Obj

			if True == tmp5892 {
				tmp5889 := PrimTail(V2971)

				tmp5890 := PrimIsPair(tmp5889)

				var ifres5884 Obj

				if True == tmp5890 {
					tmp5886 := PrimTail(V2971)

					tmp5887 := PrimTail(tmp5886)

					tmp5888 := PrimEqual(Nil, tmp5887)

					var ifres5885 Obj

					if True == tmp5888 {
						ifres5885 = True

					} else {
						ifres5885 = False

					}

					ifres5884 = ifres5885

				} else {
					ifres5884 = False

				}

				var ifres5883 Obj

				if True == ifres5884 {
					ifres5883 = True

				} else {
					ifres5883 = False

				}

				ifres5882 = ifres5883

			} else {
				ifres5882 = False

			}

			var ifres5881 Obj

			if True == ifres5882 {
				ifres5881 = True

			} else {
				ifres5881 = False

			}

			ifres5880 = ifres5881

		} else {
			ifres5880 = False

		}

		if True == ifres5880 {
			__e.Return(True)
			return
		} else {
			tmp5879 := PrimIsPair(V2971)

			var ifres5866 Obj

			if True == tmp5879 {
				tmp5877 := PrimHead(V2971)

				tmp5878 := PrimEqual(symfunction, tmp5877)

				var ifres5868 Obj

				if True == tmp5878 {
					tmp5875 := PrimTail(V2971)

					tmp5876 := PrimIsPair(tmp5875)

					var ifres5870 Obj

					if True == tmp5876 {
						tmp5872 := PrimTail(V2971)

						tmp5873 := PrimTail(tmp5872)

						tmp5874 := PrimEqual(Nil, tmp5873)

						var ifres5871 Obj

						if True == tmp5874 {
							ifres5871 = True

						} else {
							ifres5871 = False

						}

						ifres5870 = ifres5871

					} else {
						ifres5870 = False

					}

					var ifres5869 Obj

					if True == ifres5870 {
						ifres5869 = True

					} else {
						ifres5869 = False

					}

					ifres5868 = ifres5869

				} else {
					ifres5868 = False

				}

				var ifres5867 Obj

				if True == ifres5868 {
					ifres5867 = True

				} else {
					ifres5867 = False

				}

				ifres5866 = ifres5867

			} else {
				ifres5866 = False

			}

			if True == ifres5866 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	tmp5894 := Call(__e, PrimNS2Value(symdef), symshen_4fn_1call_2, tmp5863)

	_ = tmp5894

	tmp5895 := MakeNative(func(__e *ControlFlow) {
		V2972 := __e.Get(1)
		_ = V2972
		tmp5936 := PrimIsPair(V2972)

		var ifres5923 Obj

		if True == tmp5936 {
			tmp5934 := PrimHead(V2972)

			tmp5935 := PrimEqual(symfunction, tmp5934)

			var ifres5925 Obj

			if True == tmp5935 {
				tmp5932 := PrimTail(V2972)

				tmp5933 := PrimIsPair(tmp5932)

				var ifres5927 Obj

				if True == tmp5933 {
					tmp5929 := PrimTail(V2972)

					tmp5930 := PrimTail(tmp5929)

					tmp5931 := PrimEqual(Nil, tmp5930)

					var ifres5928 Obj

					if True == tmp5931 {
						ifres5928 = True

					} else {
						ifres5928 = False

					}

					ifres5927 = ifres5928

				} else {
					ifres5927 = False

				}

				var ifres5926 Obj

				if True == ifres5927 {
					ifres5926 = True

				} else {
					ifres5926 = False

				}

				ifres5925 = ifres5926

			} else {
				ifres5925 = False

			}

			var ifres5924 Obj

			if True == ifres5925 {
				ifres5924 = True

			} else {
				ifres5924 = False

			}

			ifres5923 = ifres5924

		} else {
			ifres5923 = False

		}

		if True == ifres5923 {
			tmp5921 := PrimTail(V2972)

			tmp5922 := PrimCons(symfn, tmp5921)

			__e.TailApply(PrimNS2Value(symshen_4fn_1call), tmp5922)
			return

		} else {
			tmp5920 := PrimIsPair(V2972)

			var ifres5907 Obj

			if True == tmp5920 {
				tmp5918 := PrimHead(V2972)

				tmp5919 := PrimEqual(symfn, tmp5918)

				var ifres5909 Obj

				if True == tmp5919 {
					tmp5916 := PrimTail(V2972)

					tmp5917 := PrimIsPair(tmp5916)

					var ifres5911 Obj

					if True == tmp5917 {
						tmp5913 := PrimTail(V2972)

						tmp5914 := PrimTail(tmp5913)

						tmp5915 := PrimEqual(Nil, tmp5914)

						var ifres5912 Obj

						if True == tmp5915 {
							ifres5912 = True

						} else {
							ifres5912 = False

						}

						ifres5911 = ifres5912

					} else {
						ifres5911 = False

					}

					var ifres5910 Obj

					if True == ifres5911 {
						ifres5910 = True

					} else {
						ifres5910 = False

					}

					ifres5909 = ifres5910

				} else {
					ifres5909 = False

				}

				var ifres5908 Obj

				if True == ifres5909 {
					ifres5908 = True

				} else {
					ifres5908 = False

				}

				ifres5907 = ifres5908

			} else {
				ifres5907 = False

			}

			if True == ifres5907 {
				tmp5898 := MakeNative(func(__e *ControlFlow) {
					ArityF := __e.Get(1)
					_ = ArityF
					tmp5903 := PrimEqual(ArityF, MakeNumber(-1))

					if True == tmp5903 {
						__e.Return(V2972)
						return
					} else {
						tmp5902 := PrimEqual(ArityF, MakeNumber(0))

						if True == tmp5902 {
							__e.Return(PrimSimpleError(MakeString("fn cannot be applied to a zero place function\n")))
							return
						} else {
							tmp5901 := PrimTail(V2972)

							__e.TailApply(PrimNS2Value(symshen_4lambda_1function), tmp5901, ArityF)
							return

						}

					}

				}, 1)

				tmp5904 := PrimTail(V2972)

				tmp5905 := PrimHead(tmp5904)

				tmp5906 := Call(__e, PrimNS2Value(symarity), tmp5905)

				__e.TailApply(tmp5898, tmp5906)
				return

			} else {
				__e.TailApply(PrimNS2Value(symshen_4f_1error), symshen_4fn_1call)
				return
			}

		}

	}, 1)

	tmp5937 := Call(__e, PrimNS2Value(symdef), symshen_4fn_1call, tmp5895)

	_ = tmp5937

	tmp5938 := MakeNative(func(__e *ControlFlow) {
		V2973 := __e.Get(1)
		_ = V2973
		V2974 := __e.Get(2)
		_ = V2974
		V2975 := __e.Get(3)
		_ = V2975
		tmp5939 := MakeNative(func(__e *ControlFlow) {
			Verdict := __e.Get(1)
			_ = Verdict
			tmp5940 := MakeNative(func(__e *ControlFlow) {
				Message := __e.Get(1)
				_ = Message
				__e.Return(Verdict)
				return
			}, 1)

			var ifres5946 Obj

			if True == Verdict {
				tmp5954 := Call(__e, PrimNS2Value(symshen_4loading_2))

				var ifres5948 Obj

				if True == tmp5954 {
					tmp5950 := PrimCons(sym_1, Nil)

					tmp5951 := PrimCons(sym_7, tmp5950)

					tmp5952 := Call(__e, PrimNS2Value(symelement_2), V2973, tmp5951)

					tmp5953 := PrimNot(tmp5952)

					var ifres5949 Obj

					if True == tmp5953 {
						ifres5949 = True

					} else {
						ifres5949 = False

					}

					ifres5948 = ifres5949

				} else {
					ifres5948 = False

				}

				var ifres5947 Obj

				if True == ifres5948 {
					ifres5947 = True

				} else {
					ifres5947 = False

				}

				ifres5946 = ifres5947

			} else {
				ifres5946 = False

			}

			var ifres5941 Obj

			if True == ifres5946 {
				tmp5942 := Call(__e, PrimNS2Value(symshen_4app), V2973, MakeString("\n"), symshen_4a)

				tmp5943 := PrimStringConcat(MakeString("partial application of "), tmp5942)

				tmp5944 := Call(__e, PrimNS2Value(symstoutput))

				tmp5945 := Call(__e, PrimNS2Value(sympr), tmp5943, tmp5944)

				ifres5941 = tmp5945

			} else {
				ifres5941 = symshen_4skip

			}

			__e.TailApply(tmp5940, ifres5941)
			return

		}, 1)

		tmp5955 := PrimGreatThan(V2974, V2975)

		__e.TailApply(tmp5939, tmp5955)
		return

	}, 3)

	tmp5956 := Call(__e, PrimNS2Value(symdef), symshen_4partial_1application_d_2, tmp5938)

	_ = tmp5956

	tmp5957 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimNS3Value(symshen_4_dloading_2_d))
		return
	}, 0)

	tmp5958 := Call(__e, PrimNS2Value(symdef), symshen_4loading_2, tmp5957)

	_ = tmp5958

	tmp5959 := MakeNative(func(__e *ControlFlow) {
		V2976 := __e.Get(1)
		_ = V2976
		V2977 := __e.Get(2)
		_ = V2977
		V2978 := __e.Get(3)
		_ = V2978
		tmp5960 := MakeNative(func(__e *ControlFlow) {
			Verdict := __e.Get(1)
			_ = Verdict
			tmp5961 := MakeNative(func(__e *ControlFlow) {
				Message := __e.Get(1)
				_ = Message
				__e.Return(Verdict)
				return
			}, 1)

			var ifres5972 Obj

			if True == Verdict {
				tmp5974 := Call(__e, PrimNS2Value(symshen_4loading_2))

				var ifres5973 Obj

				if True == tmp5974 {
					ifres5973 = True

				} else {
					ifres5973 = False

				}

				ifres5972 = ifres5973

			} else {
				ifres5972 = False

			}

			var ifres5962 Obj

			if True == ifres5972 {
				tmp5964 := PrimEqual(V2978, MakeNumber(1))

				var ifres5963 Obj

				if True == tmp5964 {
					ifres5963 = MakeString("")

				} else {
					ifres5963 = MakeString("s")

				}

				tmp5965 := Call(__e, PrimNS2Value(symshen_4app), ifres5963, MakeString("\n"), symshen_4a)

				tmp5966 := PrimStringConcat(MakeString(" argument"), tmp5965)

				tmp5967 := Call(__e, PrimNS2Value(symshen_4app), V2978, tmp5966, symshen_4a)

				tmp5968 := PrimStringConcat(MakeString(" might not like "), tmp5967)

				tmp5969 := Call(__e, PrimNS2Value(symshen_4app), V2976, tmp5968, symshen_4a)

				tmp5970 := Call(__e, PrimNS2Value(symstoutput))

				tmp5971 := Call(__e, PrimNS2Value(sympr), tmp5969, tmp5970)

				ifres5962 = tmp5971

			} else {
				ifres5962 = symshen_4skip

			}

			__e.TailApply(tmp5961, ifres5962)
			return

		}, 1)

		tmp5975 := PrimLessThan(V2977, V2978)

		__e.TailApply(tmp5960, tmp5975)
		return

	}, 3)

	__e.TailApply(PrimNS2Value(symdef), symshen_4overapplication_2, tmp5959)
	return

}, 0)
