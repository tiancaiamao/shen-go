package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
		MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")

		gen4214 := MakeNative(func(__e Evaluator, __args ...Obj) {
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*installing-kl*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*history*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*tc*"), False)
			gen3514 := Call(__e, ShenFunc(symshen_4dict), MakeNumber(20000))

			Call(__e, ShenFunc(symset), MakeSymbol("*property-vector*"), gen3514)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*process-counter*"), MakeNumber(0))
			gen3515 := Call(__e, ShenFunc(symvector), MakeNumber(10000))

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*varcounter*"), gen3515)

			gen3516 := Call(__e, ShenFunc(symvector), MakeNumber(10000))

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*prologvectors*"), gen3516)

			gen3517 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.Return(X)
				return
			}, 1)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*demodulation-function*"), gen3517)

			gen3518 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Arg := __args[0]
				_ = Arg
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					OnFail := __args[0]
					_ = OnFail
					__e.TailApply(ShenFunc(symthaw), OnFail)

					return
				}, 1))
				return
			}, 1)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*custom-pattern-compiler*"), gen3518)

			gen3519 := MakeNative(func(__e Evaluator, __args ...Obj) {
				Arg := __args[0]
				_ = Arg
				__e.Return(Arg)
				return
			}, 1)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*custom-pattern-reducer*"), gen3519)

			gen3520 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.function-macro"), Nil)

			gen3521 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.defprolog-macro"), gen3520)

			gen3522 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.@s-macro"), gen3521)

			gen3523 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.nl-macro"), gen3522)

			gen3524 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.synonyms-macro"), gen3523)

			gen3525 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.prolog-macro"), gen3524)

			gen3526 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.error-macro"), gen3525)

			gen3527 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.input-macro"), gen3526)

			gen3528 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.output-macro"), gen3527)

			gen3529 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.make-string-macro"), gen3528)

			gen3530 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.assoc-macro"), gen3529)

			gen3531 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.let-macro"), gen3530)

			gen3532 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.datatype-macro"), gen3531)

			gen3533 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.compile-macro"), gen3532)

			gen3534 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.put/get-macro"), gen3533)

			gen3535 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.abs-macro"), gen3534)

			gen3536 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.cases-macro"), gen3535)

			gen3537 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.timer-macro"), gen3536)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*macroreg*"), gen3537)

			gen3538 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4timer_1macro), X)

				return
			}, 1)
			gen3539 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4cases_1macro), X)

				return
			}, 1)
			gen3540 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4abs_1macro), X)

				return
			}, 1)
			gen3541 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4put_cget_1macro), X)

				return
			}, 1)
			gen3542 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4compile_1macro), X)

				return
			}, 1)
			gen3543 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4datatype_1macro), X)

				return
			}, 1)
			gen3544 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4let_1macro), X)

				return
			}, 1)
			gen3545 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4assoc_1macro), X)

				return
			}, 1)
			gen3546 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4make_1string_1macro), X)

				return
			}, 1)
			gen3547 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4output_1macro), X)

				return
			}, 1)
			gen3548 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4input_1macro), X)

				return
			}, 1)
			gen3549 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4error_1macro), X)

				return
			}, 1)
			gen3550 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4prolog_1macro), X)

				return
			}, 1)
			gen3551 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4synonyms_1macro), X)

				return
			}, 1)
			gen3552 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4nl_1macro), X)

				return
			}, 1)
			gen3553 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4_8s_1macro), X)

				return
			}, 1)
			gen3554 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4defprolog_1macro), X)

				return
			}, 1)
			gen3555 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4function_1macro), X)

				return
			}, 1)
			gen3556 := Call(__e, ShenFunc(symcons), gen3555, Nil)

			gen3557 := Call(__e, ShenFunc(symcons), gen3554, gen3556)

			gen3558 := Call(__e, ShenFunc(symcons), gen3553, gen3557)

			gen3559 := Call(__e, ShenFunc(symcons), gen3552, gen3558)

			gen3560 := Call(__e, ShenFunc(symcons), gen3551, gen3559)

			gen3561 := Call(__e, ShenFunc(symcons), gen3550, gen3560)

			gen3562 := Call(__e, ShenFunc(symcons), gen3549, gen3561)

			gen3563 := Call(__e, ShenFunc(symcons), gen3548, gen3562)

			gen3564 := Call(__e, ShenFunc(symcons), gen3547, gen3563)

			gen3565 := Call(__e, ShenFunc(symcons), gen3546, gen3564)

			gen3566 := Call(__e, ShenFunc(symcons), gen3545, gen3565)

			gen3567 := Call(__e, ShenFunc(symcons), gen3544, gen3566)

			gen3568 := Call(__e, ShenFunc(symcons), gen3543, gen3567)

			gen3569 := Call(__e, ShenFunc(symcons), gen3542, gen3568)

			gen3570 := Call(__e, ShenFunc(symcons), gen3541, gen3569)

			gen3571 := Call(__e, ShenFunc(symcons), gen3540, gen3570)

			gen3572 := Call(__e, ShenFunc(symcons), gen3539, gen3571)

			gen3573 := Call(__e, ShenFunc(symcons), gen3538, gen3572)

			Call(__e, ShenFunc(symset), MakeSymbol("*macros*"), gen3573)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*gensym*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*tracking*"), Nil)
			gen3574 := Call(__e, ShenFunc(symcons), MakeSymbol("Z"), Nil)

			gen3575 := Call(__e, ShenFunc(symcons), MakeSymbol("Y"), gen3574)

			gen3576 := Call(__e, ShenFunc(symcons), MakeSymbol("X"), gen3575)

			gen3577 := Call(__e, ShenFunc(symcons), MakeSymbol("W"), gen3576)

			gen3578 := Call(__e, ShenFunc(symcons), MakeSymbol("V"), gen3577)

			gen3579 := Call(__e, ShenFunc(symcons), MakeSymbol("U"), gen3578)

			gen3580 := Call(__e, ShenFunc(symcons), MakeSymbol("T"), gen3579)

			gen3581 := Call(__e, ShenFunc(symcons), MakeSymbol("S"), gen3580)

			gen3582 := Call(__e, ShenFunc(symcons), MakeSymbol("R"), gen3581)

			gen3583 := Call(__e, ShenFunc(symcons), MakeSymbol("Q"), gen3582)

			gen3584 := Call(__e, ShenFunc(symcons), MakeSymbol("P"), gen3583)

			gen3585 := Call(__e, ShenFunc(symcons), MakeSymbol("O"), gen3584)

			gen3586 := Call(__e, ShenFunc(symcons), MakeSymbol("N"), gen3585)

			gen3587 := Call(__e, ShenFunc(symcons), MakeSymbol("M"), gen3586)

			gen3588 := Call(__e, ShenFunc(symcons), MakeSymbol("L"), gen3587)

			gen3589 := Call(__e, ShenFunc(symcons), MakeSymbol("K"), gen3588)

			gen3590 := Call(__e, ShenFunc(symcons), MakeSymbol("J"), gen3589)

			gen3591 := Call(__e, ShenFunc(symcons), MakeSymbol("I"), gen3590)

			gen3592 := Call(__e, ShenFunc(symcons), MakeSymbol("H"), gen3591)

			gen3593 := Call(__e, ShenFunc(symcons), MakeSymbol("G"), gen3592)

			gen3594 := Call(__e, ShenFunc(symcons), MakeSymbol("F"), gen3593)

			gen3595 := Call(__e, ShenFunc(symcons), MakeSymbol("E"), gen3594)

			gen3596 := Call(__e, ShenFunc(symcons), MakeSymbol("D"), gen3595)

			gen3597 := Call(__e, ShenFunc(symcons), MakeSymbol("C"), gen3596)

			gen3598 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen3597)

			gen3599 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen3598)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*alphabet*"), gen3599)

			gen3600 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), Nil)

			gen3601 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen3600)

			gen3602 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen3601)

			gen3603 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen3602)

			gen3604 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen3603)

			gen3605 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen3604)

			gen3606 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen3605)

			gen3607 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen3606)

			gen3608 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen3607)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*special*"), gen3608)

			gen3609 := Call(__e, ShenFunc(symcons), MakeSymbol("defmacro"), Nil)

			gen3610 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.read+"), gen3609)

			gen3611 := Call(__e, ShenFunc(symcons), MakeSymbol("defcc"), gen3610)

			gen3612 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen3611)

			gen3613 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.process-datatype"), gen3612)

			gen3614 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen3613)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*extraspecial*"), gen3614)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*spy*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*datatypes*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*alldatatypes*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*shen-type-theory-enabled?*"), True)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*synonyms*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*system*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*maxcomplexity*"), MakeNumber(128))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*occurs*"), True)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*maxinferences*"), MakeNumber(1000000))
			Call(__e, ShenFunc(symset), MakeSymbol("*maximum-print-sequence-size*"), MakeNumber(20))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*catch*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*call*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*infs*"), MakeNumber(0))
			Call(__e, ShenFunc(symset), MakeSymbol("*hush*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*optimise*"), False)
			Call(__e, ShenFunc(symset), MakeSymbol("*version*"), MakeString("Shen 22.3"))
			gen3615 := Call(__e, ShenFunc(symbound_2), MakeSymbol("*home-directory*"))

			gen3616 := Call(__e, ShenFunc(symnot), gen3615)

			if True == gen3616 {
				Call(__e, ShenFunc(symset), MakeSymbol("*home-directory*"), MakeString(""))
			} else {
				MakeSymbol("shen.skip")
			}

			gen3618 := Call(__e, ShenFunc(symbound_2), MakeSymbol("*sterror*"))

			gen3619 := Call(__e, ShenFunc(symnot), gen3618)

			if True == gen3619 {
				gen3617 := Call(__e, ShenFunc(symvalue), MakeSymbol("*stoutput*"))

				Call(__e, ShenFunc(symset), MakeSymbol("*sterror*"), gen3617)

			} else {
				MakeSymbol("shen.skip")
			}

			gen3620 := Call(__e, ShenFunc(symcons), MakeNumber(1), Nil)

			gen3621 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen3620)

			gen3622 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3621)

			gen3623 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen3622)

			gen3624 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3623)

			gen3625 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen3624)

			gen3626 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3625)

			gen3627 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen3626)

			gen3628 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3627)

			gen3629 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen3628)

			gen3630 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3629)

			gen3631 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen3630)

			gen3632 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3631)

			gen3633 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen3632)

			gen3634 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3633)

			gen3635 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen3634)

			gen3636 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3635)

			gen3637 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen3636)

			gen3638 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3637)

			gen3639 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen3638)

			gen3640 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3639)

			gen3641 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen3640)

			gen3642 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3641)

			gen3643 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen3642)

			gen3644 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3643)

			gen3645 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen3644)

			gen3646 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3645)

			gen3647 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen3646)

			gen3648 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3647)

			gen3649 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen3648)

			gen3650 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3649)

			gen3651 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen3650)

			gen3652 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3651)

			gen3653 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen3652)

			gen3654 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3653)

			gen3655 := Call(__e, ShenFunc(symcons), MakeSymbol("version"), gen3654)

			gen3656 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3655)

			gen3657 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen3656)

			gen3658 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3657)

			gen3659 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen3658)

			gen3660 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3659)

			gen3661 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen3660)

			gen3662 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3661)

			gen3663 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen3662)

			gen3664 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3663)

			gen3665 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen3664)

			gen3666 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3665)

			gen3667 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen3666)

			gen3668 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3667)

			gen3669 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen3668)

			gen3670 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3669)

			gen3671 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen3670)

			gen3672 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3671)

			gen3673 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen3672)

			gen3674 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen3673)

			gen3675 := Call(__e, ShenFunc(symcons), MakeSymbol("unify!"), gen3674)

			gen3676 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen3675)

			gen3677 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen3676)

			gen3678 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3677)

			gen3679 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen3678)

			gen3680 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3679)

			gen3681 := Call(__e, ShenFunc(symcons), MakeSymbol("unput"), gen3680)

			gen3682 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3681)

			gen3683 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen3682)

			gen3684 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3683)

			gen3685 := Call(__e, ShenFunc(symcons), MakeSymbol("return"), gen3684)

			gen3686 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3685)

			gen3687 := Call(__e, ShenFunc(symcons), MakeSymbol("type"), gen3686)

			gen3688 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3687)

			gen3689 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen3688)

			gen3690 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3689)

			gen3691 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen3690)

			gen3692 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3691)

			gen3693 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen3692)

			gen3694 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3693)

			gen3695 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen3694)

			gen3696 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3695)

			gen3697 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen3696)

			gen3698 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3697)

			gen3699 := Call(__e, ShenFunc(symcons), MakeSymbol("tc?"), gen3698)

			gen3700 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3699)

			gen3701 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen3700)

			gen3702 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3701)

			gen3703 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen3702)

			gen3704 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3703)

			gen3705 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen3704)

			gen3706 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3705)

			gen3707 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen3706)

			gen3708 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3707)

			gen3709 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen3708)

			gen3710 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3709)

			gen3711 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen3710)

			gen3712 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3711)

			gen3713 := Call(__e, ShenFunc(symcons), MakeSymbol("subst"), gen3712)

			gen3714 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3713)

			gen3715 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen3714)

			gen3716 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3715)

			gen3717 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen3716)

			gen3718 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3717)

			gen3719 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen3718)

			gen3720 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3719)

			gen3721 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen3720)

			gen3722 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3721)

			gen3723 := Call(__e, ShenFunc(symcons), MakeSymbol("sterror"), gen3722)

			gen3724 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3723)

			gen3725 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), gen3724)

			gen3726 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3725)

			gen3727 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), gen3726)

			gen3728 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3727)

			gen3729 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen3728)

			gen3730 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3729)

			gen3731 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen3730)

			gen3732 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3731)

			gen3733 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen3732)

			gen3734 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3733)

			gen3735 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen3734)

			gen3736 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3735)

			gen3737 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen3736)

			gen3738 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3737)

			gen3739 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen3738)

			gen3740 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3739)

			gen3741 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen3740)

			gen3742 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3741)

			gen3743 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.require"), gen3742)

			gen3744 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3743)

			gen3745 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen3744)

			gen3746 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3745)

			gen3747 := Call(__e, ShenFunc(symcons), MakeSymbol("release"), gen3746)

			gen3748 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3747)

			gen3749 := Call(__e, ShenFunc(symcons), MakeSymbol("receive"), gen3748)

			gen3750 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3749)

			gen3751 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen3750)

			gen3752 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3751)

			gen3753 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen3752)

			gen3754 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3753)

			gen3755 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen3754)

			gen3756 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3755)

			gen3757 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen3756)

			gen3758 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3757)

			gen3759 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen3758)

			gen3760 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3759)

			gen3761 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen3760)

			gen3762 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3761)

			gen3763 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.reassemble"), gen3762)

			gen3764 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen3763)

			gen3765 := Call(__e, ShenFunc(symcons), MakeSymbol("put"), gen3764)

			gen3766 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3765)

			gen3767 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen3766)

			gen3768 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3767)

			gen3769 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen3768)

			gen3770 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3769)

			gen3771 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen3770)

			gen3772 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3771)

			gen3773 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen3772)

			gen3774 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3773)

			gen3775 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen3774)

			gen3776 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3775)

			gen3777 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen3776)

			gen3778 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3777)

			gen3779 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen3778)

			gen3780 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3779)

			gen3781 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen3780)

			gen3782 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3781)

			gen3783 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen3782)

			gen3784 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3783)

			gen3785 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen3784)

			gen3786 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3785)

			gen3787 := Call(__e, ShenFunc(symcons), MakeSymbol("porters"), gen3786)

			gen3788 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3787)

			gen3789 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen3788)

			gen3790 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3789)

			gen3791 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen3790)

			gen3792 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3791)

			gen3793 := Call(__e, ShenFunc(symcons), MakeSymbol("package"), gen3792)

			gen3794 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3793)

			gen3795 := Call(__e, ShenFunc(symcons), MakeSymbol("os"), gen3794)

			gen3796 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3795)

			gen3797 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen3796)

			gen3798 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3797)

			gen3799 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen3798)

			gen3800 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3799)

			gen3801 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), gen3800)

			gen3802 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3801)

			gen3803 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen3802)

			gen3804 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3803)

			gen3805 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen3804)

			gen3806 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3805)

			gen3807 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen3806)

			gen3808 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3807)

			gen3809 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen3808)

			gen3810 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3809)

			gen3811 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen3810)

			gen3812 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3811)

			gen3813 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen3812)

			gen3814 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3813)

			gen3815 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen3814)

			gen3816 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3815)

			gen3817 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen3816)

			gen3818 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3817)

			gen3819 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen3818)

			gen3820 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3819)

			gen3821 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen3820)

			gen3822 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3821)

			gen3823 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen3822)

			gen3824 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3823)

			gen3825 := Call(__e, ShenFunc(symcons), MakeSymbol("macroexpand"), gen3824)

			gen3826 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3825)

			gen3827 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen3826)

			gen3828 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3827)

			gen3829 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen3828)

			gen3830 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3829)

			gen3831 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen3830)

			gen3832 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3831)

			gen3833 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen3832)

			gen3834 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3833)

			gen3835 := Call(__e, ShenFunc(symcons), MakeSymbol("lineread"), gen3834)

			gen3836 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3835)

			gen3837 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen3836)

			gen3838 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3837)

			gen3839 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen3838)

			gen3840 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3839)

			gen3841 := Call(__e, ShenFunc(symcons), MakeSymbol("language"), gen3840)

			gen3842 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3841)

			gen3843 := Call(__e, ShenFunc(symcons), MakeSymbol("kill"), gen3842)

			gen3844 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3843)

			gen3845 := Call(__e, ShenFunc(symcons), MakeSymbol("it"), gen3844)

			gen3846 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3845)

			gen3847 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen3846)

			gen3848 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3847)

			gen3849 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen3848)

			gen3850 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3849)

			gen3851 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen3850)

			gen3852 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3851)

			gen3853 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen3852)

			gen3854 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3853)

			gen3855 := Call(__e, ShenFunc(symcons), MakeSymbol("input"), gen3854)

			gen3856 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3855)

			gen3857 := Call(__e, ShenFunc(symcons), MakeSymbol("inferences"), gen3856)

			gen3858 := Call(__e, ShenFunc(symcons), MakeNumber(4), gen3857)

			gen3859 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen3858)

			gen3860 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3859)

			gen3861 := Call(__e, ShenFunc(symcons), MakeSymbol("intern"), gen3860)

			gen3862 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3861)

			gen3863 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen3862)

			gen3864 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3863)

			gen3865 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen3864)

			gen3866 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3865)

			gen3867 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen3866)

			gen3868 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3867)

			gen3869 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen3868)

			gen3870 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3869)

			gen3871 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen3870)

			gen3872 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3871)

			gen3873 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen3872)

			gen3874 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3873)

			gen3875 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen3874)

			gen3876 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3875)

			gen3877 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen3876)

			gen3878 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3877)

			gen3879 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen3878)

			gen3880 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3879)

			gen3881 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen3880)

			gen3882 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3881)

			gen3883 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen3882)

			gen3884 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3883)

			gen3885 := Call(__e, ShenFunc(symcons), MakeSymbol("<-address"), gen3884)

			gen3886 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3885)

			gen3887 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen3886)

			gen3888 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3887)

			gen3889 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen3888)

			gen3890 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3889)

			gen3891 := Call(__e, ShenFunc(symcons), MakeSymbol("get"), gen3890)

			gen3892 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3891)

			gen3893 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen3892)

			gen3894 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3893)

			gen3895 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen3894)

			gen3896 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3895)

			gen3897 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen3896)

			gen3898 := Call(__e, ShenFunc(symcons), MakeNumber(5), gen3897)

			gen3899 := Call(__e, ShenFunc(symcons), MakeSymbol("findall"), gen3898)

			gen3900 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3899)

			gen3901 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen3900)

			gen3902 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3901)

			gen3903 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), gen3902)

			gen3904 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3903)

			gen3905 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen3904)

			gen3906 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3905)

			gen3907 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen3906)

			gen3908 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3907)

			gen3909 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen3908)

			gen3910 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3909)

			gen3911 := Call(__e, ShenFunc(symcons), MakeSymbol("eval-kl"), gen3910)

			gen3912 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3911)

			gen3913 := Call(__e, ShenFunc(symcons), MakeSymbol("eval"), gen3912)

			gen3914 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3913)

			gen3915 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.interror"), gen3914)

			gen3916 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3915)

			gen3917 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen3916)

			gen3918 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3917)

			gen3919 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen3918)

			gen3920 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3919)

			gen3921 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen3920)

			gen3922 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3921)

			gen3923 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen3922)

			gen3924 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3923)

			gen3925 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen3924)

			gen3926 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3925)

			gen3927 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen3926)

			gen3928 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3927)

			gen3929 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen3928)

			gen3930 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3929)

			gen3931 := Call(__e, ShenFunc(symcons), MakeSymbol("declare"), gen3930)

			gen3932 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3931)

			gen3933 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen3932)

			gen3934 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3933)

			gen3935 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen3934)

			gen3936 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3935)

			gen3937 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen3936)

			gen3938 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3937)

			gen3939 := Call(__e, ShenFunc(symcons), MakeSymbol("concat"), gen3938)

			gen3940 := Call(__e, ShenFunc(symcons), MakeNumber(3), gen3939)

			gen3941 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen3940)

			gen3942 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3941)

			gen3943 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen3942)

			gen3944 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3943)

			gen3945 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen3944)

			gen3946 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3945)

			gen3947 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen3946)

			gen3948 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3947)

			gen3949 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen3948)

			gen3950 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3949)

			gen3951 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen3950)

			gen3952 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3951)

			gen3953 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen3952)

			gen3954 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3953)

			gen3955 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen3954)

			gen3956 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3955)

			gen3957 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen3956)

			gen3958 := Call(__e, ShenFunc(symcons), MakeNumber(2), gen3957)

			gen3959 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen3958)

			gen3960 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3959)

			gen3961 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector"), gen3960)

			gen3962 := Call(__e, ShenFunc(symcons), MakeNumber(1), gen3961)

			gen3963 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen3962)

			gen3964 := Call(__e, ShenFunc(symcons), MakeNumber(0), gen3963)

			gen3965 := Call(__e, ShenFunc(symcons), MakeSymbol("abort"), gen3964)

			Call(__e, ShenFunc(symshen_4initialise__arity__table), gen3965)

			gen3966 := Call(__e, ShenFunc(symintern), MakeString("shen"))

			gen3967 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector"), Nil)

			gen3968 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen3967)

			gen3969 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen3968)

			gen3970 := Call(__e, ShenFunc(symcons), MakeSymbol("<-address"), gen3969)

			gen3971 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen3970)

			gen3972 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen3971)

			gen3973 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen3972)

			gen3974 := Call(__e, ShenFunc(symcons), MakeSymbol("abort"), gen3973)

			gen3975 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen3974)

			gen3976 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen3975)

			gen3977 := Call(__e, ShenFunc(symcons), MakeSymbol("bar!"), gen3976)

			gen3978 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen3977)

			gen3979 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen3978)

			gen3980 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen3979)

			gen3981 := Call(__e, ShenFunc(symcons), MakeSymbol("bind"), gen3980)

			gen3982 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen3981)

			gen3983 := Call(__e, ShenFunc(symcons), MakeSymbol("call"), gen3982)

			gen3984 := Call(__e, ShenFunc(symcons), MakeSymbol("cases"), gen3983)

			gen3985 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen3984)

			gen3986 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen3985)

			gen3987 := Call(__e, ShenFunc(symcons), MakeSymbol("concat"), gen3986)

			gen3988 := Call(__e, ShenFunc(symcons), MakeSymbol("cond"), gen3987)

			gen3989 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen3988)

			gen3990 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen3989)

			gen3991 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen3990)

			gen3992 := Call(__e, ShenFunc(symcons), MakeSymbol("cut"), gen3991)

			gen3993 := Call(__e, ShenFunc(symcons), MakeSymbol("datatype"), gen3992)

			gen3994 := Call(__e, ShenFunc(symcons), MakeSymbol("declare"), gen3993)

			gen3995 := Call(__e, ShenFunc(symcons), MakeSymbol("defprolog"), gen3994)

			gen3996 := Call(__e, ShenFunc(symcons), MakeSymbol("defcc"), gen3995)

			gen3997 := Call(__e, ShenFunc(symcons), MakeSymbol("defmacro"), gen3996)

			gen3998 := Call(__e, ShenFunc(symcons), MakeSymbol("define"), gen3997)

			gen3999 := Call(__e, ShenFunc(symcons), MakeSymbol("defun"), gen3998)

			gen4000 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen3999)

			gen4001 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen4000)

			gen4002 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen4001)

			gen4003 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen4002)

			gen4004 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen4003)

			gen4005 := Call(__e, ShenFunc(symcons), MakeSymbol("error"), gen4004)

			gen4006 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen4005)

			gen4007 := Call(__e, ShenFunc(symcons), MakeSymbol("eval"), gen4006)

			gen4008 := Call(__e, ShenFunc(symcons), MakeSymbol("eval-kl"), gen4007)

			gen4009 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen4008)

			gen4010 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen4009)

			gen4011 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen4010)

			gen4012 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen4011)

			gen4013 := Call(__e, ShenFunc(symcons), False, gen4012)

			gen4014 := Call(__e, ShenFunc(symcons), MakeSymbol("findall"), gen4013)

			gen4015 := Call(__e, ShenFunc(symcons), MakeSymbol("fwhen"), gen4014)

			gen4016 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen4015)

			gen4017 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), gen4016)

			gen4018 := Call(__e, ShenFunc(symcons), MakeSymbol("file"), gen4017)

			gen4019 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen4018)

			gen4020 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen4019)

			gen4021 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen4020)

			gen4022 := Call(__e, ShenFunc(symcons), MakeSymbol("function"), gen4021)

			gen4023 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen4022)

			gen4024 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen4023)

			gen4025 := Call(__e, ShenFunc(symcons), MakeSymbol("get"), gen4024)

			gen4026 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen4025)

			gen4027 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen4026)

			gen4028 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen4027)

			gen4029 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen4028)

			gen4030 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen4029)

			gen4031 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen4030)

			gen4032 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen4031)

			gen4033 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen4032)

			gen4034 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen4033)

			gen4035 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), gen4034)

			gen4036 := Call(__e, ShenFunc(symcons), MakeSymbol("it"), gen4035)

			gen4037 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen4036)

			gen4038 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen4037)

			gen4039 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen4038)

			gen4040 := Call(__e, ShenFunc(symcons), MakeSymbol("input"), gen4039)

			gen4041 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen4040)

			gen4042 := Call(__e, ShenFunc(symcons), MakeSymbol("intern"), gen4041)

			gen4043 := Call(__e, ShenFunc(symcons), MakeSymbol("inferences"), gen4042)

			gen4044 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen4043)

			gen4045 := Call(__e, ShenFunc(symcons), MakeSymbol("is"), gen4044)

			gen4046 := Call(__e, ShenFunc(symcons), MakeSymbol("kill"), gen4045)

			gen4047 := Call(__e, ShenFunc(symcons), MakeSymbol("language"), gen4046)

			gen4048 := Call(__e, ShenFunc(symcons), MakeSymbol("lambda"), gen4047)

			gen4049 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen4048)

			gen4050 := Call(__e, ShenFunc(symcons), MakeSymbol("let"), gen4049)

			gen4051 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen4050)

			gen4052 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen4051)

			gen4053 := Call(__e, ShenFunc(symcons), MakeSymbol("lineread"), gen4052)

			gen4054 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4053)

			gen4055 := Call(__e, ShenFunc(symcons), MakeSymbol("loaded"), gen4054)

			gen4056 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen4055)

			gen4057 := Call(__e, ShenFunc(symcons), MakeSymbol("make-string"), gen4056)

			gen4058 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen4057)

			gen4059 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen4058)

			gen4060 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen4059)

			gen4061 := Call(__e, ShenFunc(symcons), MakeSymbol("macroexpand"), gen4060)

			gen4062 := Call(__e, ShenFunc(symcons), MakeSymbol("mode"), gen4061)

			gen4063 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen4062)

			gen4064 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen4063)

			gen4065 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen4064)

			gen4066 := Call(__e, ShenFunc(symcons), MakeSymbol("null"), gen4065)

			gen4067 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4066)

			gen4068 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen4067)

			gen4069 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen4068)

			gen4070 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen4069)

			gen4071 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen4070)

			gen4072 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), gen4071)

			gen4073 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen4072)

			gen4074 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen4073)

			gen4075 := Call(__e, ShenFunc(symcons), MakeSymbol("os"), gen4074)

			gen4076 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), gen4075)

			gen4077 := Call(__e, ShenFunc(symcons), MakeSymbol("output"), gen4076)

			gen4078 := Call(__e, ShenFunc(symcons), MakeSymbol("package"), gen4077)

			gen4079 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen4078)

			gen4080 := Call(__e, ShenFunc(symcons), MakeSymbol("porters"), gen4079)

			gen4081 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen4080)

			gen4082 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen4081)

			gen4083 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen4082)

			gen4084 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen4083)

			gen4085 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen4084)

			gen4086 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen4085)

			gen4087 := Call(__e, ShenFunc(symcons), MakeSymbol("prolog?"), gen4086)

			gen4088 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen4087)

			gen4089 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen4088)

			gen4090 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen4089)

			gen4091 := Call(__e, ShenFunc(symcons), MakeSymbol("put"), gen4090)

			gen4092 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen4091)

			gen4093 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen4092)

			gen4094 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen4093)

			gen4095 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen4094)

			gen4096 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen4095)

			gen4097 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen4096)

			gen4098 := Call(__e, ShenFunc(symcons), MakeSymbol("receive"), gen4097)

			gen4099 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen4098)

			gen4100 := Call(__e, ShenFunc(symcons), MakeSymbol("release"), gen4099)

			gen4101 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen4100)

			gen4102 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen4101)

			gen4103 := Call(__e, ShenFunc(symcons), MakeSymbol("run"), gen4102)

			gen4104 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen4103)

			gen4105 := Call(__e, ShenFunc(symcons), MakeSymbol("save"), gen4104)

			gen4106 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen4105)

			gen4107 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen4106)

			gen4108 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen4107)

			gen4109 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen4108)

			gen4110 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen4109)

			gen4111 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen4110)

			gen4112 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), gen4111)

			gen4113 := Call(__e, ShenFunc(symcons), MakeSymbol("sterror"), gen4112)

			gen4114 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), gen4113)

			gen4115 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4114)

			gen4116 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen4115)

			gen4117 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen4116)

			gen4118 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen4117)

			gen4119 := Call(__e, ShenFunc(symcons), MakeSymbol("subst"), gen4118)

			gen4120 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen4119)

			gen4121 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen4120)

			gen4122 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen4121)

			gen4123 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4122)

			gen4124 := Call(__e, ShenFunc(symcons), MakeSymbol("synonyms"), gen4123)

			gen4125 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen4124)

			gen4126 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen4125)

			gen4127 := Call(__e, ShenFunc(symcons), MakeSymbol("tlv"), gen4126)

			gen4128 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen4127)

			gen4129 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen4128)

			gen4130 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen4129)

			gen4131 := Call(__e, ShenFunc(symcons), MakeSymbol("tc?"), gen4130)

			gen4132 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen4131)

			gen4133 := Call(__e, ShenFunc(symcons), MakeSymbol("time"), gen4132)

			gen4134 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen4133)

			gen4135 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen4134)

			gen4136 := Call(__e, ShenFunc(symcons), True, gen4135)

			gen4137 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen4136)

			gen4138 := Call(__e, ShenFunc(symcons), MakeSymbol("type"), gen4137)

			gen4139 := Call(__e, ShenFunc(symcons), MakeSymbol("return"), gen4138)

			gen4140 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen4139)

			gen4141 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen4140)

			gen4142 := Call(__e, ShenFunc(symcons), MakeSymbol("unput"), gen4141)

			gen4143 := Call(__e, ShenFunc(symcons), MakeSymbol("unify!"), gen4142)

			gen4144 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen4143)

			gen4145 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen4144)

			gen4146 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.unix"), gen4145)

			gen4147 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), gen4146)

			gen4148 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen4147)

			gen4149 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen4148)

			gen4150 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen4149)

			gen4151 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4150)

			gen4152 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen4151)

			gen4153 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen4152)

			gen4154 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen4153)

			gen4155 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen4154)

			gen4156 := Call(__e, ShenFunc(symcons), MakeSymbol("verified"), gen4155)

			gen4157 := Call(__e, ShenFunc(symcons), MakeSymbol("version"), gen4156)

			gen4158 := Call(__e, ShenFunc(symcons), MakeSymbol("warn"), gen4157)

			gen4159 := Call(__e, ShenFunc(symcons), MakeSymbol("when"), gen4158)

			gen4160 := Call(__e, ShenFunc(symcons), MakeSymbol("where"), gen4159)

			gen4161 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen4160)

			gen4162 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen4161)

			gen4163 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen4162)

			gen4164 := Call(__e, ShenFunc(symcons), MakeSymbol(">>"), gen4163)

			gen4165 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen4164)

			gen4166 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen4165)

			gen4167 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen4166)

			gen4168 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen4167)

			gen4169 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen4168)

			gen4170 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen4169)

			gen4171 := Call(__e, ShenFunc(symcons), MakeSymbol("$"), gen4170)

			gen4172 := Call(__e, ShenFunc(symcons), MakeSymbol("=!"), gen4171)

			gen4173 := Call(__e, ShenFunc(symcons), MakeSymbol("/."), gen4172)

			gen4174 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen4173)

			gen4175 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen4174)

			gen4176 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen4175)

			gen4177 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen4176)

			gen4178 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen4177)

			gen4179 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen4178)

			gen4180 := Call(__e, ShenFunc(symcons), MakeSymbol("->"), gen4179)

			gen4181 := Call(__e, ShenFunc(symcons), MakeSymbol("<-"), gen4180)

			gen4182 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen4181)

			gen4183 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen4182)

			gen4184 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen4183)

			gen4185 := Call(__e, ShenFunc(symcons), MakeSymbol("*hush*"), gen4184)

			gen4186 := Call(__e, ShenFunc(symcons), MakeSymbol("*porters*"), gen4185)

			gen4187 := Call(__e, ShenFunc(symcons), MakeSymbol("*port*"), gen4186)

			gen4188 := Call(__e, ShenFunc(symcons), MakeSymbol("*property-vector*"), gen4187)

			gen4189 := Call(__e, ShenFunc(symcons), MakeSymbol("*release*"), gen4188)

			gen4190 := Call(__e, ShenFunc(symcons), MakeSymbol("*os*"), gen4189)

			gen4191 := Call(__e, ShenFunc(symcons), MakeSymbol("*macros*"), gen4190)

			gen4192 := Call(__e, ShenFunc(symcons), MakeSymbol("*maximum-print-sequence-size*"), gen4191)

			gen4193 := Call(__e, ShenFunc(symcons), MakeSymbol("*version*"), gen4192)

			gen4194 := Call(__e, ShenFunc(symcons), MakeSymbol("*home-directory*"), gen4193)

			gen4195 := Call(__e, ShenFunc(symcons), MakeSymbol("*sterror*"), gen4194)

			gen4196 := Call(__e, ShenFunc(symcons), MakeSymbol("*stoutput*"), gen4195)

			gen4197 := Call(__e, ShenFunc(symcons), MakeSymbol("*stinput*"), gen4196)

			gen4198 := Call(__e, ShenFunc(symcons), MakeSymbol("*implementation*"), gen4197)

			gen4199 := Call(__e, ShenFunc(symcons), MakeSymbol("*language*"), gen4198)

			gen4200 := Call(__e, ShenFunc(symcons), MakeSymbol(","), gen4199)

			gen4201 := Call(__e, ShenFunc(symcons), MakeSymbol("_"), gen4200)

			gen4202 := Call(__e, ShenFunc(symcons), MakeSymbol(":="), gen4201)

			gen4203 := Call(__e, ShenFunc(symcons), MakeSymbol(":-"), gen4202)

			gen4204 := Call(__e, ShenFunc(symcons), MakeSymbol(";"), gen4203)

			gen4205 := Call(__e, ShenFunc(symcons), MakeSymbol(":"), gen4204)

			gen4206 := Call(__e, ShenFunc(symcons), MakeSymbol("&&"), gen4205)

			gen4207 := Call(__e, ShenFunc(symcons), MakeSymbol("<--"), gen4206)

			gen4208 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4207)

			gen4209 := Call(__e, ShenFunc(symcons), MakeSymbol("{"), gen4208)

			gen4210 := Call(__e, ShenFunc(symcons), MakeSymbol("}"), gen4209)

			gen4211 := Call(__e, ShenFunc(symcons), MakeSymbol("!"), gen4210)

			gen4212 := Call(__e, ShenFunc(symvalue), MakeSymbol("*property-vector*"))

			Call(__e, ShenFunc(symput), gen3966, MakeSymbol("shen.external-symbols"), gen4211, gen4212)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*history*"), Nil)
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*step*"), False)
			gen4213 := Call(__e, ShenFunc(symabsvector), MakeNumber(0))

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*empty-absvector*"), gen4213)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-environment"), gen4214)

		gen5352 := MakeNative(func(__e Evaluator, __args ...Obj) {
			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), Nil)
			gen4215 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4216 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4215)

			gen4217 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4216)

			gen4218 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen4217)

			gen4219 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4220 := Call(__e, ShenFunc(symcons), gen4218, gen4219)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4220)

			gen4221 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4222 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4221)

			gen4223 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4224 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4223)

			gen4225 := Call(__e, ShenFunc(symcons), gen4224, Nil)

			gen4226 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4225)

			gen4227 := Call(__e, ShenFunc(symcons), gen4222, gen4226)

			gen4228 := Call(__e, ShenFunc(symcons), gen4227, Nil)

			gen4229 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4228)

			gen4230 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4229)

			gen4231 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen4230)

			gen4232 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4233 := Call(__e, ShenFunc(symcons), gen4231, gen4232)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4233)

			gen4234 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4235 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4234)

			gen4236 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen4235)

			gen4237 := Call(__e, ShenFunc(symcons), gen4236, Nil)

			gen4238 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4237)

			gen4239 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen4238)

			gen4240 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen4239)

			gen4241 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4242 := Call(__e, ShenFunc(symcons), gen4240, gen4241)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4242)

			gen4243 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4244 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4243)

			gen4245 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4244)

			gen4246 := Call(__e, ShenFunc(symcons), gen4245, Nil)

			gen4247 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4246)

			gen4248 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4247)

			gen4249 := Call(__e, ShenFunc(symcons), gen4248, Nil)

			gen4250 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4249)

			gen4251 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4250)

			gen4252 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.app"), gen4251)

			gen4253 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4254 := Call(__e, ShenFunc(symcons), gen4252, gen4253)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4254)

			gen4255 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4256 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4255)

			gen4257 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4258 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4257)

			gen4259 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4260 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4259)

			gen4261 := Call(__e, ShenFunc(symcons), gen4260, Nil)

			gen4262 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4261)

			gen4263 := Call(__e, ShenFunc(symcons), gen4258, gen4262)

			gen4264 := Call(__e, ShenFunc(symcons), gen4263, Nil)

			gen4265 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4264)

			gen4266 := Call(__e, ShenFunc(symcons), gen4256, gen4265)

			gen4267 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen4266)

			gen4268 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4269 := Call(__e, ShenFunc(symcons), gen4267, gen4268)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4269)

			gen4270 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4271 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4270)

			gen4272 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4271)

			gen4273 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen4272)

			gen4274 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4275 := Call(__e, ShenFunc(symcons), gen4273, gen4274)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4275)

			gen4276 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4277 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4276)

			gen4278 := Call(__e, ShenFunc(symcons), gen4277, Nil)

			gen4279 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4278)

			gen4280 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4281 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4280)

			gen4282 := Call(__e, ShenFunc(symcons), gen4281, Nil)

			gen4283 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4282)

			gen4284 := Call(__e, ShenFunc(symcons), gen4279, gen4283)

			gen4285 := Call(__e, ShenFunc(symcons), gen4284, Nil)

			gen4286 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4285)

			gen4287 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4286)

			gen4288 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen4287)

			gen4289 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4290 := Call(__e, ShenFunc(symcons), gen4288, gen4289)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4290)

			gen4291 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4292 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4291)

			gen4293 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4292)

			gen4294 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen4293)

			gen4295 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4296 := Call(__e, ShenFunc(symcons), gen4294, gen4295)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4296)

			gen4297 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4298 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4297)

			gen4299 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4298)

			gen4300 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen4299)

			gen4301 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4302 := Call(__e, ShenFunc(symcons), gen4300, gen4301)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4302)

			gen4303 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4304 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4303)

			gen4305 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4304)

			gen4306 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen4305)

			gen4307 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4308 := Call(__e, ShenFunc(symcons), gen4306, gen4307)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4308)

			gen4309 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4310 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen4309)

			gen4311 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4312 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4311)

			gen4313 := Call(__e, ShenFunc(symcons), gen4312, Nil)

			gen4314 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4313)

			gen4315 := Call(__e, ShenFunc(symcons), gen4310, gen4314)

			gen4316 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen4315)

			gen4317 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4318 := Call(__e, ShenFunc(symcons), gen4316, gen4317)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4318)

			gen4319 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4320 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4319)

			gen4321 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4320)

			gen4322 := Call(__e, ShenFunc(symcons), gen4321, Nil)

			gen4323 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4322)

			gen4324 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4323)

			gen4325 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen4324)

			gen4326 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4327 := Call(__e, ShenFunc(symcons), gen4325, gen4326)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4327)

			gen4328 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4329 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen4328)

			gen4330 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4329)

			gen4331 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4332 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4331)

			gen4333 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4332)

			gen4334 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4335 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4334)

			gen4336 := Call(__e, ShenFunc(symcons), gen4333, gen4335)

			gen4337 := Call(__e, ShenFunc(symcons), gen4336, Nil)

			gen4338 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4337)

			gen4339 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4338)

			gen4340 := Call(__e, ShenFunc(symcons), gen4339, Nil)

			gen4341 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4340)

			gen4342 := Call(__e, ShenFunc(symcons), gen4330, gen4341)

			gen4343 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen4342)

			gen4344 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4345 := Call(__e, ShenFunc(symcons), gen4343, gen4344)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4345)

			gen4346 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4347 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4346)

			gen4348 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4347)

			gen4349 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen4348)

			gen4350 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4351 := Call(__e, ShenFunc(symcons), gen4349, gen4350)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4351)

			gen4352 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4353 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4352)

			gen4354 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4353)

			gen4355 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4356 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4355)

			gen4357 := Call(__e, ShenFunc(symcons), gen4354, gen4356)

			gen4358 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen4357)

			gen4359 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4360 := Call(__e, ShenFunc(symcons), gen4358, gen4359)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4360)

			gen4361 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4362 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4361)

			gen4363 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4364 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4363)

			gen4365 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4366 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4365)

			gen4367 := Call(__e, ShenFunc(symcons), gen4366, Nil)

			gen4368 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4367)

			gen4369 := Call(__e, ShenFunc(symcons), gen4364, gen4368)

			gen4370 := Call(__e, ShenFunc(symcons), gen4369, Nil)

			gen4371 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4370)

			gen4372 := Call(__e, ShenFunc(symcons), gen4362, gen4371)

			gen4373 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen4372)

			gen4374 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4375 := Call(__e, ShenFunc(symcons), gen4373, gen4374)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4375)

			gen4376 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4377 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4376)

			gen4378 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen4377)

			gen4379 := Call(__e, ShenFunc(symcons), gen4378, Nil)

			gen4380 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4379)

			gen4381 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4380)

			gen4382 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen4381)

			gen4383 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4384 := Call(__e, ShenFunc(symcons), gen4382, gen4383)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4384)

			gen4385 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4386 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4385)

			gen4387 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4388 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4387)

			gen4389 := Call(__e, ShenFunc(symcons), gen4388, Nil)

			gen4390 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen4389)

			gen4391 := Call(__e, ShenFunc(symcons), gen4386, gen4390)

			gen4392 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen4391)

			gen4393 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4394 := Call(__e, ShenFunc(symcons), gen4392, gen4393)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4394)

			gen4395 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4396 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4395)

			gen4397 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4398 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4397)

			gen4399 := Call(__e, ShenFunc(symcons), gen4398, Nil)

			gen4400 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.==>"), gen4399)

			gen4401 := Call(__e, ShenFunc(symcons), gen4396, gen4400)

			gen4402 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen4401)

			gen4403 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4404 := Call(__e, ShenFunc(symcons), gen4402, gen4403)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4404)

			gen4405 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4406 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4405)

			gen4407 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4408 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4407)

			gen4409 := Call(__e, ShenFunc(symcons), gen4406, gen4408)

			gen4410 := Call(__e, ShenFunc(symcons), gen4409, Nil)

			gen4411 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4410)

			gen4412 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4411)

			gen4413 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen4412)

			gen4414 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4415 := Call(__e, ShenFunc(symcons), gen4413, gen4414)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4415)

			gen4416 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4417 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4416)

			gen4418 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4417)

			gen4419 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen4418)

			gen4420 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4421 := Call(__e, ShenFunc(symcons), gen4419, gen4420)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4421)

			gen4422 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4423 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4422)

			gen4424 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4423)

			gen4425 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen4424)

			gen4426 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4427 := Call(__e, ShenFunc(symcons), gen4425, gen4426)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4427)

			gen4428 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4429 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4428)

			gen4430 := Call(__e, ShenFunc(symcons), gen4429, Nil)

			gen4431 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4430)

			gen4432 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4431)

			gen4433 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen4432)

			gen4434 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4435 := Call(__e, ShenFunc(symcons), gen4433, gen4434)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4435)

			gen4436 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4437 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4436)

			gen4438 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen4437)

			gen4439 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen4438)

			gen4440 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4441 := Call(__e, ShenFunc(symcons), gen4439, gen4440)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4441)

			gen4442 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4443 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4442)

			gen4444 := Call(__e, ShenFunc(symcons), gen4443, Nil)

			gen4445 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4444)

			gen4446 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4445)

			gen4447 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen4446)

			gen4448 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4449 := Call(__e, ShenFunc(symcons), gen4447, gen4448)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4449)

			gen4450 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4451 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4450)

			gen4452 := Call(__e, ShenFunc(symcons), MakeSymbol("fail"), gen4451)

			gen4453 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4454 := Call(__e, ShenFunc(symcons), gen4452, gen4453)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4454)

			gen4455 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4456 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4455)

			gen4457 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4456)

			gen4458 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4459 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4458)

			gen4460 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4459)

			gen4461 := Call(__e, ShenFunc(symcons), gen4460, Nil)

			gen4462 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4461)

			gen4463 := Call(__e, ShenFunc(symcons), gen4457, gen4462)

			gen4464 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen4463)

			gen4465 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4466 := Call(__e, ShenFunc(symcons), gen4464, gen4465)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4466)

			gen4467 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4468 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4467)

			gen4469 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4468)

			gen4470 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4471 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4470)

			gen4472 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4471)

			gen4473 := Call(__e, ShenFunc(symcons), gen4472, Nil)

			gen4474 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4473)

			gen4475 := Call(__e, ShenFunc(symcons), gen4469, gen4474)

			gen4476 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen4475)

			gen4477 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4478 := Call(__e, ShenFunc(symcons), gen4476, gen4477)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4478)

			gen4479 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4480 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen4479)

			gen4481 := Call(__e, ShenFunc(symcons), gen4480, Nil)

			gen4482 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4481)

			gen4483 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4482)

			gen4484 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen4483)

			gen4485 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4486 := Call(__e, ShenFunc(symcons), gen4484, gen4485)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4486)

			gen4487 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4488 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen4487)

			gen4489 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4488)

			gen4490 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4491 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4490)

			gen4492 := Call(__e, ShenFunc(symcons), gen4489, gen4491)

			gen4493 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen4492)

			gen4494 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4495 := Call(__e, ShenFunc(symcons), gen4493, gen4494)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4495)

			gen4496 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4497 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4496)

			gen4498 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4497)

			gen4499 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4500 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4499)

			gen4501 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4500)

			gen4502 := Call(__e, ShenFunc(symcons), gen4501, Nil)

			gen4503 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4502)

			gen4504 := Call(__e, ShenFunc(symcons), gen4498, gen4503)

			gen4505 := Call(__e, ShenFunc(symcons), MakeSymbol("function"), gen4504)

			gen4506 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4507 := Call(__e, ShenFunc(symcons), gen4505, gen4506)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4507)

			gen4508 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4509 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4508)

			gen4510 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4509)

			gen4511 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen4510)

			gen4512 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4513 := Call(__e, ShenFunc(symcons), gen4511, gen4512)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4513)

			gen4514 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4515 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4514)

			gen4516 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4517 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4516)

			gen4518 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4517)

			gen4519 := Call(__e, ShenFunc(symcons), gen4518, Nil)

			gen4520 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4519)

			gen4521 := Call(__e, ShenFunc(symcons), gen4515, gen4520)

			gen4522 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen4521)

			gen4523 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4524 := Call(__e, ShenFunc(symcons), gen4522, gen4523)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4524)

			gen4525 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4526 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4525)

			gen4527 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4528 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4527)

			gen4529 := Call(__e, ShenFunc(symcons), gen4528, Nil)

			gen4530 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4529)

			gen4531 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4530)

			gen4532 := Call(__e, ShenFunc(symcons), gen4531, Nil)

			gen4533 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4532)

			gen4534 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4533)

			gen4535 := Call(__e, ShenFunc(symcons), gen4534, Nil)

			gen4536 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4535)

			gen4537 := Call(__e, ShenFunc(symcons), gen4526, gen4536)

			gen4538 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen4537)

			gen4539 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4540 := Call(__e, ShenFunc(symcons), gen4538, gen4539)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4540)

			gen4541 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4542 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4541)

			gen4543 := Call(__e, ShenFunc(symcons), gen4542, Nil)

			gen4544 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4543)

			gen4545 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4544)

			gen4546 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4545)

			gen4547 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4548 := Call(__e, ShenFunc(symcons), gen4546, gen4547)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4548)

			gen4549 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4550 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4549)

			gen4551 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4550)

			gen4552 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen4551)

			gen4553 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4554 := Call(__e, ShenFunc(symcons), gen4552, gen4553)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4554)

			gen4555 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4556 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4555)

			gen4557 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4556)

			gen4558 := Call(__e, ShenFunc(symcons), gen4557, Nil)

			gen4559 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4558)

			gen4560 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4559)

			gen4561 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen4560)

			gen4562 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4563 := Call(__e, ShenFunc(symcons), gen4561, gen4562)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4563)

			gen4564 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4565 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4564)

			gen4566 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4567 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4566)

			gen4568 := Call(__e, ShenFunc(symcons), gen4565, gen4567)

			gen4569 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen4568)

			gen4570 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4571 := Call(__e, ShenFunc(symcons), gen4569, gen4570)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4571)

			gen4572 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4573 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4572)

			gen4574 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4575 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4574)

			gen4576 := Call(__e, ShenFunc(symcons), gen4573, gen4575)

			gen4577 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen4576)

			gen4578 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4579 := Call(__e, ShenFunc(symcons), gen4577, gen4578)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4579)

			gen4580 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4581 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4580)

			gen4582 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4581)

			gen4583 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen4582)

			gen4584 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4585 := Call(__e, ShenFunc(symcons), gen4583, gen4584)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4585)

			gen4586 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4587 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4586)

			gen4588 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4587)

			gen4589 := Call(__e, ShenFunc(symcons), gen4588, Nil)

			gen4590 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4589)

			gen4591 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4590)

			gen4592 := Call(__e, ShenFunc(symcons), gen4591, Nil)

			gen4593 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4592)

			gen4594 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen4593)

			gen4595 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen4594)

			gen4596 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4597 := Call(__e, ShenFunc(symcons), gen4595, gen4596)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4597)

			gen4598 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4599 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4598)

			gen4600 := Call(__e, ShenFunc(symcons), MakeSymbol("it"), gen4599)

			gen4601 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4602 := Call(__e, ShenFunc(symcons), gen4600, gen4601)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4602)

			gen4603 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4604 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4603)

			gen4605 := Call(__e, ShenFunc(symcons), MakeSymbol("implementation"), gen4604)

			gen4606 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4607 := Call(__e, ShenFunc(symcons), gen4605, gen4606)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4607)

			gen4608 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4609 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4608)

			gen4610 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4611 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4610)

			gen4612 := Call(__e, ShenFunc(symcons), gen4611, Nil)

			gen4613 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4612)

			gen4614 := Call(__e, ShenFunc(symcons), gen4609, gen4613)

			gen4615 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen4614)

			gen4616 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4617 := Call(__e, ShenFunc(symcons), gen4615, gen4616)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4617)

			gen4618 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4619 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4618)

			gen4620 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4621 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4620)

			gen4622 := Call(__e, ShenFunc(symcons), gen4621, Nil)

			gen4623 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4622)

			gen4624 := Call(__e, ShenFunc(symcons), gen4619, gen4623)

			gen4625 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen4624)

			gen4626 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4627 := Call(__e, ShenFunc(symcons), gen4625, gen4626)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4627)

			gen4628 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4629 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4628)

			gen4630 := Call(__e, ShenFunc(symcons), MakeSymbol("inferences"), gen4629)

			gen4631 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4632 := Call(__e, ShenFunc(symcons), gen4630, gen4631)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4632)

			gen4633 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4634 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4633)

			gen4635 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4634)

			gen4636 := Call(__e, ShenFunc(symcons), gen4635, Nil)

			gen4637 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4636)

			gen4638 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4637)

			gen4639 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.insert"), gen4638)

			gen4640 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4641 := Call(__e, ShenFunc(symcons), gen4639, gen4640)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4641)

			gen4642 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4643 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4642)

			gen4644 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4643)

			gen4645 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen4644)

			gen4646 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4647 := Call(__e, ShenFunc(symcons), gen4645, gen4646)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4647)

			gen4648 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4649 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4648)

			gen4650 := Call(__e, ShenFunc(symcons), gen4649, Nil)

			gen4651 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4650)

			gen4652 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4651)

			gen4653 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen4652)

			gen4654 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4655 := Call(__e, ShenFunc(symcons), gen4653, gen4654)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4655)

			gen4656 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4657 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4656)

			gen4658 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4659 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4658)

			gen4660 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4661 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4660)

			gen4662 := Call(__e, ShenFunc(symcons), gen4661, Nil)

			gen4663 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4662)

			gen4664 := Call(__e, ShenFunc(symcons), gen4659, gen4663)

			gen4665 := Call(__e, ShenFunc(symcons), gen4664, Nil)

			gen4666 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4665)

			gen4667 := Call(__e, ShenFunc(symcons), gen4657, gen4666)

			gen4668 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen4667)

			gen4669 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4670 := Call(__e, ShenFunc(symcons), gen4668, gen4669)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4670)

			gen4671 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4672 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4671)

			gen4673 := Call(__e, ShenFunc(symcons), MakeSymbol("kill"), gen4672)

			gen4674 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4675 := Call(__e, ShenFunc(symcons), gen4673, gen4674)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4675)

			gen4676 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4677 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4676)

			gen4678 := Call(__e, ShenFunc(symcons), MakeSymbol("language"), gen4677)

			gen4679 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4680 := Call(__e, ShenFunc(symcons), gen4678, gen4679)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4680)

			gen4681 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4682 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4681)

			gen4683 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4684 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4683)

			gen4685 := Call(__e, ShenFunc(symcons), gen4682, gen4684)

			gen4686 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen4685)

			gen4687 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4688 := Call(__e, ShenFunc(symcons), gen4686, gen4687)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4688)

			gen4689 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4690 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen4689)

			gen4691 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4692 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4691)

			gen4693 := Call(__e, ShenFunc(symcons), gen4690, gen4692)

			gen4694 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen4693)

			gen4695 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4696 := Call(__e, ShenFunc(symcons), gen4694, gen4695)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4696)

			gen4697 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4698 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4697)

			gen4699 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4698)

			gen4700 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen4699)

			gen4701 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4702 := Call(__e, ShenFunc(symcons), gen4700, gen4701)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4702)

			gen4703 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4704 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4703)

			gen4705 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4704)

			gen4706 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4707 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4706)

			gen4708 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4709 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4708)

			gen4710 := Call(__e, ShenFunc(symcons), gen4709, Nil)

			gen4711 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4710)

			gen4712 := Call(__e, ShenFunc(symcons), gen4707, gen4711)

			gen4713 := Call(__e, ShenFunc(symcons), gen4712, Nil)

			gen4714 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4713)

			gen4715 := Call(__e, ShenFunc(symcons), gen4705, gen4714)

			gen4716 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen4715)

			gen4717 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4718 := Call(__e, ShenFunc(symcons), gen4716, gen4717)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4718)

			gen4719 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4720 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4719)

			gen4721 := Call(__e, ShenFunc(symcons), gen4720, Nil)

			gen4722 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4721)

			gen4723 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4722)

			gen4724 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4725 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4724)

			gen4726 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4727 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4726)

			gen4728 := Call(__e, ShenFunc(symcons), gen4727, Nil)

			gen4729 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4728)

			gen4730 := Call(__e, ShenFunc(symcons), gen4725, gen4729)

			gen4731 := Call(__e, ShenFunc(symcons), gen4730, Nil)

			gen4732 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4731)

			gen4733 := Call(__e, ShenFunc(symcons), gen4723, gen4732)

			gen4734 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen4733)

			gen4735 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4736 := Call(__e, ShenFunc(symcons), gen4734, gen4735)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4736)

			gen4737 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4738 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4737)

			gen4739 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4738)

			gen4740 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen4739)

			gen4741 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4742 := Call(__e, ShenFunc(symcons), gen4740, gen4741)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4742)

			gen4743 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4744 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4743)

			gen4745 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4744)

			gen4746 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen4745)

			gen4747 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4748 := Call(__e, ShenFunc(symcons), gen4746, gen4747)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4748)

			gen4749 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4750 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4749)

			gen4751 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4750)

			gen4752 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen4751)

			gen4753 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4754 := Call(__e, ShenFunc(symcons), gen4752, gen4753)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4754)

			gen4755 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4756 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4755)

			gen4757 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen4756)

			gen4758 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen4757)

			gen4759 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4760 := Call(__e, ShenFunc(symcons), gen4758, gen4759)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4760)

			gen4761 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4762 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4761)

			gen4763 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4764 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4763)

			gen4765 := Call(__e, ShenFunc(symcons), gen4762, gen4764)

			gen4766 := Call(__e, ShenFunc(symcons), gen4765, Nil)

			gen4767 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4766)

			gen4768 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4767)

			gen4769 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen4768)

			gen4770 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4771 := Call(__e, ShenFunc(symcons), gen4769, gen4770)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4771)

			gen4772 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4773 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4772)

			gen4774 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4773)

			gen4775 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen4774)

			gen4776 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4777 := Call(__e, ShenFunc(symcons), gen4775, gen4776)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4777)

			gen4778 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4779 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4778)

			gen4780 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen4779)

			gen4781 := Call(__e, ShenFunc(symcons), gen4780, Nil)

			gen4782 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4781)

			gen4783 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4782)

			gen4784 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen4783)

			gen4785 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4786 := Call(__e, ShenFunc(symcons), gen4784, gen4785)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4786)

			gen4787 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4788 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4787)

			gen4789 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4788)

			gen4790 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen4789)

			gen4791 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4792 := Call(__e, ShenFunc(symcons), gen4790, gen4791)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4792)

			gen4793 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4794 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4793)

			gen4795 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4794)

			gen4796 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen4795)

			gen4797 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4798 := Call(__e, ShenFunc(symcons), gen4796, gen4797)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4798)

			gen4799 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4800 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4799)

			gen4801 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen4800)

			gen4802 := Call(__e, ShenFunc(symcons), gen4801, Nil)

			gen4803 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4802)

			gen4804 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), gen4803)

			gen4805 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen4804)

			gen4806 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4807 := Call(__e, ShenFunc(symcons), gen4805, gen4806)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4807)

			gen4808 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4809 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4808)

			gen4810 := Call(__e, ShenFunc(symcons), MakeSymbol("os"), gen4809)

			gen4811 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4812 := Call(__e, ShenFunc(symcons), gen4810, gen4811)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4812)

			gen4813 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen4814 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4813)

			gen4815 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4814)

			gen4816 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen4815)

			gen4817 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4818 := Call(__e, ShenFunc(symcons), gen4816, gen4817)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4818)

			gen4819 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4820 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4819)

			gen4821 := Call(__e, ShenFunc(symcons), MakeSymbol("port"), gen4820)

			gen4822 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4823 := Call(__e, ShenFunc(symcons), gen4821, gen4822)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4823)

			gen4824 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4825 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4824)

			gen4826 := Call(__e, ShenFunc(symcons), MakeSymbol("porters"), gen4825)

			gen4827 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4828 := Call(__e, ShenFunc(symcons), gen4826, gen4827)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4828)

			gen4829 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4830 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4829)

			gen4831 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen4830)

			gen4832 := Call(__e, ShenFunc(symcons), gen4831, Nil)

			gen4833 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4832)

			gen4834 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4833)

			gen4835 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen4834)

			gen4836 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4837 := Call(__e, ShenFunc(symcons), gen4835, gen4836)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4837)

			gen4838 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen4839 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen4838)

			gen4840 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4841 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4840)

			gen4842 := Call(__e, ShenFunc(symcons), gen4839, gen4841)

			gen4843 := Call(__e, ShenFunc(symcons), gen4842, Nil)

			gen4844 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4843)

			gen4845 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4844)

			gen4846 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen4845)

			gen4847 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4848 := Call(__e, ShenFunc(symcons), gen4846, gen4847)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4848)

			gen4849 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4850 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4849)

			gen4851 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4850)

			gen4852 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen4851)

			gen4853 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4854 := Call(__e, ShenFunc(symcons), gen4852, gen4853)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4854)

			gen4855 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4856 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4855)

			gen4857 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4856)

			gen4858 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4859 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4858)

			gen4860 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4859)

			gen4861 := Call(__e, ShenFunc(symcons), gen4860, Nil)

			gen4862 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4861)

			gen4863 := Call(__e, ShenFunc(symcons), gen4857, gen4862)

			gen4864 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen4863)

			gen4865 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4866 := Call(__e, ShenFunc(symcons), gen4864, gen4865)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4866)

			gen4867 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4868 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4867)

			gen4869 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4870 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4869)

			gen4871 := Call(__e, ShenFunc(symcons), gen4870, Nil)

			gen4872 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4871)

			gen4873 := Call(__e, ShenFunc(symcons), gen4868, gen4872)

			gen4874 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen4873)

			gen4875 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4876 := Call(__e, ShenFunc(symcons), gen4874, gen4875)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4876)

			gen4877 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4878 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4877)

			gen4879 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4878)

			gen4880 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.proc-nl"), gen4879)

			gen4881 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4882 := Call(__e, ShenFunc(symcons), gen4880, gen4881)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4882)

			gen4883 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4884 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4883)

			gen4885 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4884)

			gen4886 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen4887 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4886)

			gen4888 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4887)

			gen4889 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4890 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen4889)

			gen4891 := Call(__e, ShenFunc(symcons), gen4888, gen4890)

			gen4892 := Call(__e, ShenFunc(symcons), gen4891, Nil)

			gen4893 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4892)

			gen4894 := Call(__e, ShenFunc(symcons), gen4885, gen4893)

			gen4895 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen4894)

			gen4896 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4897 := Call(__e, ShenFunc(symcons), gen4895, gen4896)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4897)

			gen4898 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4899 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4898)

			gen4900 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4899)

			gen4901 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen4900)

			gen4902 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4903 := Call(__e, ShenFunc(symcons), gen4901, gen4902)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4903)

			gen4904 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4905 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4904)

			gen4906 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen4907 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4906)

			gen4908 := Call(__e, ShenFunc(symcons), gen4907, Nil)

			gen4909 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4908)

			gen4910 := Call(__e, ShenFunc(symcons), gen4905, gen4909)

			gen4911 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen4910)

			gen4912 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4913 := Call(__e, ShenFunc(symcons), gen4911, gen4912)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4913)

			gen4914 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen4915 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen4914)

			gen4916 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4917 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4916)

			gen4918 := Call(__e, ShenFunc(symcons), gen4915, gen4917)

			gen4919 := Call(__e, ShenFunc(symcons), gen4918, Nil)

			gen4920 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4919)

			gen4921 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4920)

			gen4922 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.prhush"), gen4921)

			gen4923 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4924 := Call(__e, ShenFunc(symcons), gen4922, gen4923)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4924)

			gen4925 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen4926 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4925)

			gen4927 := Call(__e, ShenFunc(symcons), gen4926, Nil)

			gen4928 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4927)

			gen4929 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen4928)

			gen4930 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen4929)

			gen4931 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4932 := Call(__e, ShenFunc(symcons), gen4930, gen4931)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4932)

			gen4933 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen4934 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen4933)

			gen4935 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen4936 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4935)

			gen4937 := Call(__e, ShenFunc(symcons), gen4934, gen4936)

			gen4938 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen4937)

			gen4939 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4940 := Call(__e, ShenFunc(symcons), gen4938, gen4939)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4940)

			gen4941 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen4942 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen4941)

			gen4943 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4944 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4943)

			gen4945 := Call(__e, ShenFunc(symcons), gen4942, gen4944)

			gen4946 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen4945)

			gen4947 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4948 := Call(__e, ShenFunc(symcons), gen4946, gen4947)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4948)

			gen4949 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen4950 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4949)

			gen4951 := Call(__e, ShenFunc(symcons), gen4950, Nil)

			gen4952 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4951)

			gen4953 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4952)

			gen4954 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen4953)

			gen4955 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4956 := Call(__e, ShenFunc(symcons), gen4954, gen4955)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4956)

			gen4957 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4958 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4957)

			gen4959 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4958)

			gen4960 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen4959)

			gen4961 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4962 := Call(__e, ShenFunc(symcons), gen4960, gen4961)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4962)

			gen4963 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen4964 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4963)

			gen4965 := Call(__e, ShenFunc(symcons), gen4964, Nil)

			gen4966 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4965)

			gen4967 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4966)

			gen4968 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen4967)

			gen4969 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4970 := Call(__e, ShenFunc(symcons), gen4968, gen4969)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4970)

			gen4971 := Call(__e, ShenFunc(symcons), MakeSymbol("unit"), Nil)

			gen4972 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4971)

			gen4973 := Call(__e, ShenFunc(symcons), gen4972, Nil)

			gen4974 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4973)

			gen4975 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen4974)

			gen4976 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen4975)

			gen4977 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4978 := Call(__e, ShenFunc(symcons), gen4976, gen4977)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4978)

			gen4979 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen4980 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4979)

			gen4981 := Call(__e, ShenFunc(symcons), MakeSymbol("release"), gen4980)

			gen4982 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4983 := Call(__e, ShenFunc(symcons), gen4981, gen4982)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4983)

			gen4984 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4985 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4984)

			gen4986 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4987 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4986)

			gen4988 := Call(__e, ShenFunc(symcons), gen4987, Nil)

			gen4989 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4988)

			gen4990 := Call(__e, ShenFunc(symcons), gen4985, gen4989)

			gen4991 := Call(__e, ShenFunc(symcons), gen4990, Nil)

			gen4992 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen4991)

			gen4993 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen4992)

			gen4994 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen4993)

			gen4995 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen4996 := Call(__e, ShenFunc(symcons), gen4994, gen4995)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen4996)

			gen4997 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen4998 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4997)

			gen4999 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5000 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen4999)

			gen5001 := Call(__e, ShenFunc(symcons), gen5000, Nil)

			gen5002 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5001)

			gen5003 := Call(__e, ShenFunc(symcons), gen4998, gen5002)

			gen5004 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen5003)

			gen5005 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5006 := Call(__e, ShenFunc(symcons), gen5004, gen5005)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5006)

			gen5007 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5008 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5007)

			gen5009 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen5008)

			gen5010 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen5009)

			gen5011 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5012 := Call(__e, ShenFunc(symcons), gen5010, gen5011)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5012)

			gen5013 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen5014 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen5013)

			gen5015 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5014)

			gen5016 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen5017 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5016)

			gen5018 := Call(__e, ShenFunc(symcons), gen5015, gen5017)

			gen5019 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen5018)

			gen5020 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5021 := Call(__e, ShenFunc(symcons), gen5019, gen5020)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5021)

			gen5022 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen5023 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5022)

			gen5024 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5023)

			gen5025 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen5024)

			gen5026 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5027 := Call(__e, ShenFunc(symcons), gen5025, gen5026)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5027)

			gen5028 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5029 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5028)

			gen5030 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5029)

			gen5031 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen5030)

			gen5032 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5033 := Call(__e, ShenFunc(symcons), gen5031, gen5032)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5033)

			gen5034 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5035 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5034)

			gen5036 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5035)

			gen5037 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen5036)

			gen5038 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5039 := Call(__e, ShenFunc(symcons), gen5037, gen5038)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5039)

			gen5040 := Call(__e, ShenFunc(symcons), MakeSymbol("in"), Nil)

			gen5041 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen5040)

			gen5042 := Call(__e, ShenFunc(symcons), gen5041, Nil)

			gen5043 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5042)

			gen5044 := Call(__e, ShenFunc(symcons), MakeSymbol("stinput"), gen5043)

			gen5045 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5046 := Call(__e, ShenFunc(symcons), gen5044, gen5045)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5046)

			gen5047 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen5048 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen5047)

			gen5049 := Call(__e, ShenFunc(symcons), gen5048, Nil)

			gen5050 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5049)

			gen5051 := Call(__e, ShenFunc(symcons), MakeSymbol("sterror"), gen5050)

			gen5052 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5053 := Call(__e, ShenFunc(symcons), gen5051, gen5052)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5053)

			gen5054 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen5055 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen5054)

			gen5056 := Call(__e, ShenFunc(symcons), gen5055, Nil)

			gen5057 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5056)

			gen5058 := Call(__e, ShenFunc(symcons), MakeSymbol("stoutput"), gen5057)

			gen5059 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5060 := Call(__e, ShenFunc(symcons), gen5058, gen5059)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5060)

			gen5061 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5062 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5061)

			gen5063 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5062)

			gen5064 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen5063)

			gen5065 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5066 := Call(__e, ShenFunc(symcons), gen5064, gen5065)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5066)

			gen5067 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen5068 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5067)

			gen5069 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5068)

			gen5070 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen5069)

			gen5071 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5072 := Call(__e, ShenFunc(symcons), gen5070, gen5071)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5072)

			gen5073 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5074 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5073)

			gen5075 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen5074)

			gen5076 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen5075)

			gen5077 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5078 := Call(__e, ShenFunc(symcons), gen5076, gen5077)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5078)

			gen5079 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen5080 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5079)

			gen5081 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen5080)

			gen5082 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen5081)

			gen5083 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5084 := Call(__e, ShenFunc(symcons), gen5082, gen5083)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5084)

			gen5085 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5086 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen5085)

			gen5087 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5088 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5087)

			gen5089 := Call(__e, ShenFunc(symcons), gen5086, gen5088)

			gen5090 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen5089)

			gen5091 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5092 := Call(__e, ShenFunc(symcons), gen5090, gen5091)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5092)

			gen5093 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5094 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5093)

			gen5095 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5094)

			gen5096 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen5095)

			gen5097 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5098 := Call(__e, ShenFunc(symcons), gen5096, gen5097)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5098)

			gen5099 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen5100 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5099)

			gen5101 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5100)

			gen5102 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen5101)

			gen5103 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5104 := Call(__e, ShenFunc(symcons), gen5102, gen5103)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5104)

			gen5105 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5106 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen5105)

			gen5107 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5108 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen5107)

			gen5109 := Call(__e, ShenFunc(symcons), gen5108, Nil)

			gen5110 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5109)

			gen5111 := Call(__e, ShenFunc(symcons), gen5106, gen5110)

			gen5112 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen5111)

			gen5113 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5114 := Call(__e, ShenFunc(symcons), gen5112, gen5113)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5114)

			gen5115 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen5116 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5115)

			gen5117 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen5116)

			gen5118 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen5117)

			gen5119 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5120 := Call(__e, ShenFunc(symcons), gen5118, gen5119)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5120)

			gen5121 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5122 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen5121)

			gen5123 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5124 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen5123)

			gen5125 := Call(__e, ShenFunc(symcons), gen5124, Nil)

			gen5126 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5125)

			gen5127 := Call(__e, ShenFunc(symcons), gen5122, gen5126)

			gen5128 := Call(__e, ShenFunc(symcons), MakeSymbol("tlv"), gen5127)

			gen5129 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5130 := Call(__e, ShenFunc(symcons), gen5128, gen5129)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5130)

			gen5131 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5132 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5131)

			gen5133 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5132)

			gen5134 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen5133)

			gen5135 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5136 := Call(__e, ShenFunc(symcons), gen5134, gen5135)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5136)

			gen5137 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5138 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5137)

			gen5139 := Call(__e, ShenFunc(symcons), MakeSymbol("tc?"), gen5138)

			gen5140 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5141 := Call(__e, ShenFunc(symcons), gen5139, gen5140)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5141)

			gen5142 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5143 := Call(__e, ShenFunc(symcons), MakeSymbol("lazy"), gen5142)

			gen5144 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5145 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5144)

			gen5146 := Call(__e, ShenFunc(symcons), gen5143, gen5145)

			gen5147 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen5146)

			gen5148 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5149 := Call(__e, ShenFunc(symcons), gen5147, gen5148)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5149)

			gen5150 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen5151 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5150)

			gen5152 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5151)

			gen5153 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen5152)

			gen5154 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5155 := Call(__e, ShenFunc(symcons), gen5153, gen5154)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5155)

			gen5156 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5157 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5156)

			gen5158 := Call(__e, ShenFunc(symcons), MakeSymbol("exception"), gen5157)

			gen5159 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5160 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5159)

			gen5161 := Call(__e, ShenFunc(symcons), gen5158, gen5160)

			gen5162 := Call(__e, ShenFunc(symcons), gen5161, Nil)

			gen5163 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5162)

			gen5164 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5163)

			gen5165 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen5164)

			gen5166 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5167 := Call(__e, ShenFunc(symcons), gen5165, gen5166)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5167)

			gen5168 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5169 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5168)

			gen5170 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5169)

			gen5171 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen5170)

			gen5172 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5173 := Call(__e, ShenFunc(symcons), gen5171, gen5172)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5173)

			gen5174 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen5175 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5174)

			gen5176 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5175)

			gen5177 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen5176)

			gen5178 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5179 := Call(__e, ShenFunc(symcons), gen5177, gen5178)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5179)

			gen5180 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5181 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen5180)

			gen5182 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5183 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen5182)

			gen5184 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5185 := Call(__e, ShenFunc(symcons), MakeSymbol("list"), gen5184)

			gen5186 := Call(__e, ShenFunc(symcons), gen5185, Nil)

			gen5187 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5186)

			gen5188 := Call(__e, ShenFunc(symcons), gen5183, gen5187)

			gen5189 := Call(__e, ShenFunc(symcons), gen5188, Nil)

			gen5190 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5189)

			gen5191 := Call(__e, ShenFunc(symcons), gen5181, gen5190)

			gen5192 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen5191)

			gen5193 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5194 := Call(__e, ShenFunc(symcons), gen5192, gen5193)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5194)

			gen5195 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen5196 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5195)

			gen5197 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5196)

			gen5198 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), Nil)

			gen5199 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5198)

			gen5200 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5199)

			gen5201 := Call(__e, ShenFunc(symcons), gen5200, Nil)

			gen5202 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5201)

			gen5203 := Call(__e, ShenFunc(symcons), gen5197, gen5202)

			gen5204 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen5203)

			gen5205 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5206 := Call(__e, ShenFunc(symcons), gen5204, gen5205)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5206)

			gen5207 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen5208 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5207)

			gen5209 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5208)

			gen5210 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen5209)

			gen5211 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5212 := Call(__e, ShenFunc(symcons), gen5210, gen5211)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5212)

			gen5213 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), Nil)

			gen5214 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5213)

			gen5215 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol"), gen5214)

			gen5216 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen5215)

			gen5217 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5218 := Call(__e, ShenFunc(symcons), gen5216, gen5217)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5218)

			gen5219 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5220 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5219)

			gen5221 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5220)

			gen5222 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen5221)

			gen5223 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5224 := Call(__e, ShenFunc(symcons), gen5222, gen5223)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5224)

			gen5225 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5226 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5225)

			gen5227 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5226)

			gen5228 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen5227)

			gen5229 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5230 := Call(__e, ShenFunc(symcons), gen5228, gen5229)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5230)

			gen5231 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), Nil)

			gen5232 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5231)

			gen5233 := Call(__e, ShenFunc(symcons), MakeSymbol("version"), gen5232)

			gen5234 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5235 := Call(__e, ShenFunc(symcons), gen5233, gen5234)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5235)

			gen5236 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), Nil)

			gen5237 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5236)

			gen5238 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5237)

			gen5239 := Call(__e, ShenFunc(symcons), gen5238, Nil)

			gen5240 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5239)

			gen5241 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen5240)

			gen5242 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen5241)

			gen5243 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5244 := Call(__e, ShenFunc(symcons), gen5242, gen5243)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5244)

			gen5245 := Call(__e, ShenFunc(symcons), MakeSymbol("out"), Nil)

			gen5246 := Call(__e, ShenFunc(symcons), MakeSymbol("stream"), gen5245)

			gen5247 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5248 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5247)

			gen5249 := Call(__e, ShenFunc(symcons), gen5246, gen5248)

			gen5250 := Call(__e, ShenFunc(symcons), gen5249, Nil)

			gen5251 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5250)

			gen5252 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5251)

			gen5253 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen5252)

			gen5254 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5255 := Call(__e, ShenFunc(symcons), gen5253, gen5254)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5255)

			gen5256 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5257 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5256)

			gen5258 := Call(__e, ShenFunc(symcons), MakeSymbol("string"), gen5257)

			gen5259 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen5258)

			gen5260 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5261 := Call(__e, ShenFunc(symcons), gen5259, gen5260)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5261)

			gen5262 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5263 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5262)

			gen5264 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5263)

			gen5265 := Call(__e, ShenFunc(symcons), gen5264, Nil)

			gen5266 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5265)

			gen5267 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5266)

			gen5268 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen5267)

			gen5269 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5270 := Call(__e, ShenFunc(symcons), gen5268, gen5269)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5270)

			gen5271 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5272 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5271)

			gen5273 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5272)

			gen5274 := Call(__e, ShenFunc(symcons), gen5273, Nil)

			gen5275 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5274)

			gen5276 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5275)

			gen5277 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen5276)

			gen5278 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5279 := Call(__e, ShenFunc(symcons), gen5277, gen5278)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5279)

			gen5280 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5281 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5280)

			gen5282 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5281)

			gen5283 := Call(__e, ShenFunc(symcons), gen5282, Nil)

			gen5284 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5283)

			gen5285 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5284)

			gen5286 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen5285)

			gen5287 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5288 := Call(__e, ShenFunc(symcons), gen5286, gen5287)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5288)

			gen5289 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5290 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5289)

			gen5291 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5290)

			gen5292 := Call(__e, ShenFunc(symcons), gen5291, Nil)

			gen5293 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5292)

			gen5294 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5293)

			gen5295 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen5294)

			gen5296 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5297 := Call(__e, ShenFunc(symcons), gen5295, gen5296)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5297)

			gen5298 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5299 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5298)

			gen5300 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5299)

			gen5301 := Call(__e, ShenFunc(symcons), gen5300, Nil)

			gen5302 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5301)

			gen5303 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5302)

			gen5304 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen5303)

			gen5305 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5306 := Call(__e, ShenFunc(symcons), gen5304, gen5305)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5306)

			gen5307 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5308 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5307)

			gen5309 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5308)

			gen5310 := Call(__e, ShenFunc(symcons), gen5309, Nil)

			gen5311 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5310)

			gen5312 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5311)

			gen5313 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen5312)

			gen5314 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5315 := Call(__e, ShenFunc(symcons), gen5313, gen5314)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5315)

			gen5316 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5317 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5316)

			gen5318 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5317)

			gen5319 := Call(__e, ShenFunc(symcons), gen5318, Nil)

			gen5320 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5319)

			gen5321 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5320)

			gen5322 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen5321)

			gen5323 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5324 := Call(__e, ShenFunc(symcons), gen5322, gen5323)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5324)

			gen5325 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5326 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5325)

			gen5327 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5326)

			gen5328 := Call(__e, ShenFunc(symcons), gen5327, Nil)

			gen5329 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5328)

			gen5330 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5329)

			gen5331 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen5330)

			gen5332 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5333 := Call(__e, ShenFunc(symcons), gen5331, gen5332)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5333)

			gen5334 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), Nil)

			gen5335 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5334)

			gen5336 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5335)

			gen5337 := Call(__e, ShenFunc(symcons), gen5336, Nil)

			gen5338 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5337)

			gen5339 := Call(__e, ShenFunc(symcons), MakeSymbol("number"), gen5338)

			gen5340 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen5339)

			gen5341 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5342 := Call(__e, ShenFunc(symcons), gen5340, gen5341)

			Call(__e, ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5342)

			gen5343 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean"), Nil)

			gen5344 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5343)

			gen5345 := Call(__e, ShenFunc(symcons), MakeSymbol("B"), gen5344)

			gen5346 := Call(__e, ShenFunc(symcons), gen5345, Nil)

			gen5347 := Call(__e, ShenFunc(symcons), MakeSymbol("-->"), gen5346)

			gen5348 := Call(__e, ShenFunc(symcons), MakeSymbol("A"), gen5347)

			gen5349 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen5348)

			gen5350 := Call(__e, ShenFunc(symvalue), MakeSymbol("shen.*signedfuncs*"))

			gen5351 := Call(__e, ShenFunc(symcons), gen5349, gen5350)

			__e.TailApply(ShenFunc(symset), MakeSymbol("shen.*signedfuncs*"), gen5351)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-signedfuncs"), gen5352)

		gen5625 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen5353 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3239 := __args[0]
				_ = V3239
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3240 := __args[0]
					_ = V3240
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3241 := __args[0]
						_ = V3241
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1absvector_2), V3239, V3240, V3241)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5354 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-absvector?"), gen5353)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5354)

			gen5355 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3249 := __args[0]
				_ = V3249
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3250 := __args[0]
					_ = V3250
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3251 := __args[0]
						_ = V3251
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1adjoin), V3249, V3250, V3251)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5356 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-adjoin"), gen5355)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5356)

			gen5357 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3259 := __args[0]
				_ = V3259
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3260 := __args[0]
					_ = V3260
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3261 := __args[0]
						_ = V3261
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1and), V3259, V3260, V3261)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5358 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-and"), gen5357)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5358)

			gen5359 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3269 := __args[0]
				_ = V3269
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3270 := __args[0]
					_ = V3270
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3271 := __args[0]
						_ = V3271
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4app), V3269, V3270, V3271)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5360 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.app"), gen5359)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5360)

			gen5361 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3279 := __args[0]
				_ = V3279
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3280 := __args[0]
					_ = V3280
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3281 := __args[0]
						_ = V3281
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1append), V3279, V3280, V3281)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5362 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-append"), gen5361)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5362)

			gen5363 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3289 := __args[0]
				_ = V3289
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3290 := __args[0]
					_ = V3290
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3291 := __args[0]
						_ = V3291
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1arity), V3289, V3290, V3291)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5364 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-arity"), gen5363)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5364)

			gen5365 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3299 := __args[0]
				_ = V3299
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3300 := __args[0]
					_ = V3300
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3301 := __args[0]
						_ = V3301
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1assoc), V3299, V3300, V3301)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5366 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-assoc"), gen5365)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5366)

			gen5367 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3309 := __args[0]
				_ = V3309
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3310 := __args[0]
					_ = V3310
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3311 := __args[0]
						_ = V3311
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1boolean_2), V3309, V3310, V3311)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5368 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-boolean?"), gen5367)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5368)

			gen5369 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3319 := __args[0]
				_ = V3319
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3320 := __args[0]
					_ = V3320
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3321 := __args[0]
						_ = V3321
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1bound_2), V3319, V3320, V3321)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5370 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-bound?"), gen5369)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5370)

			gen5371 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3329 := __args[0]
				_ = V3329
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3330 := __args[0]
					_ = V3330
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3331 := __args[0]
						_ = V3331
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1cd), V3329, V3330, V3331)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5372 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-cd"), gen5371)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5372)

			gen5373 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3339 := __args[0]
				_ = V3339
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3340 := __args[0]
					_ = V3340
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3341 := __args[0]
						_ = V3341
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1close), V3339, V3340, V3341)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5374 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-close"), gen5373)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5374)

			gen5375 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3349 := __args[0]
				_ = V3349
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3350 := __args[0]
					_ = V3350
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3351 := __args[0]
						_ = V3351
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1cn), V3349, V3350, V3351)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5376 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-cn"), gen5375)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5376)

			gen5377 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3359 := __args[0]
				_ = V3359
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3360 := __args[0]
					_ = V3360
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3361 := __args[0]
						_ = V3361
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1compile), V3359, V3360, V3361)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5378 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-compile"), gen5377)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5378)

			gen5379 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3369 := __args[0]
				_ = V3369
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3370 := __args[0]
					_ = V3370
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3371 := __args[0]
						_ = V3371
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1cons_2), V3369, V3370, V3371)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5380 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-cons?"), gen5379)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5380)

			gen5381 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3379 := __args[0]
				_ = V3379
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3380 := __args[0]
					_ = V3380
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3381 := __args[0]
						_ = V3381
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1destroy), V3379, V3380, V3381)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5382 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-destroy"), gen5381)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5382)

			gen5383 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3389 := __args[0]
				_ = V3389
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3390 := __args[0]
					_ = V3390
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3391 := __args[0]
						_ = V3391
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1difference), V3389, V3390, V3391)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5384 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-difference"), gen5383)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5384)

			gen5385 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3399 := __args[0]
				_ = V3399
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3400 := __args[0]
					_ = V3400
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3401 := __args[0]
						_ = V3401
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1do), V3399, V3400, V3401)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5386 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-do"), gen5385)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5386)

			gen5387 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3409 := __args[0]
				_ = V3409
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3410 := __args[0]
					_ = V3410
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3411 := __args[0]
						_ = V3411
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5e_6), V3409, V3410, V3411)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5388 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<e>"), gen5387)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5388)

			gen5389 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3419 := __args[0]
				_ = V3419
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3420 := __args[0]
					_ = V3420
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3421 := __args[0]
						_ = V3421
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5_b_6), V3419, V3420, V3421)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5390 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<!>"), gen5389)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5390)

			gen5391 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3429 := __args[0]
				_ = V3429
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3430 := __args[0]
					_ = V3430
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3431 := __args[0]
						_ = V3431
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1element_2), V3429, V3430, V3431)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5392 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-element?"), gen5391)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5392)

			gen5393 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3439 := __args[0]
				_ = V3439
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3440 := __args[0]
					_ = V3440
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3441 := __args[0]
						_ = V3441
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1empty_2), V3439, V3440, V3441)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5394 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-empty?"), gen5393)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5394)

			gen5395 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3449 := __args[0]
				_ = V3449
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3450 := __args[0]
					_ = V3450
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3451 := __args[0]
						_ = V3451
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1enable_1type_1theory), V3449, V3450, V3451)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5396 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-enable-type-theory"), gen5395)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5396)

			gen5397 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3459 := __args[0]
				_ = V3459
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3460 := __args[0]
					_ = V3460
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3461 := __args[0]
						_ = V3461
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1external), V3459, V3460, V3461)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5398 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-external"), gen5397)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5398)

			gen5399 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3469 := __args[0]
				_ = V3469
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3470 := __args[0]
					_ = V3470
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3471 := __args[0]
						_ = V3471
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1error_1to_1string), V3469, V3470, V3471)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5400 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-error-to-string"), gen5399)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5400)

			gen5401 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3479 := __args[0]
				_ = V3479
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3480 := __args[0]
					_ = V3480
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3481 := __args[0]
						_ = V3481
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1explode), V3479, V3480, V3481)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5402 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-explode"), gen5401)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5402)

			gen5403 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3489 := __args[0]
				_ = V3489
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3490 := __args[0]
					_ = V3490
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3491 := __args[0]
						_ = V3491
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fail), V3489, V3490, V3491)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5404 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fail"), gen5403)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5404)

			gen5405 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3499 := __args[0]
				_ = V3499
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3500 := __args[0]
					_ = V3500
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3501 := __args[0]
						_ = V3501
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fail_1if), V3499, V3500, V3501)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5406 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fail-if"), gen5405)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5406)

			gen5407 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3509 := __args[0]
				_ = V3509
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3510 := __args[0]
					_ = V3510
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3511 := __args[0]
						_ = V3511
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fix), V3509, V3510, V3511)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5408 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fix"), gen5407)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5408)

			gen5409 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3519 := __args[0]
				_ = V3519
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3520 := __args[0]
					_ = V3520
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3521 := __args[0]
						_ = V3521
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1freeze), V3519, V3520, V3521)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5410 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-freeze"), gen5409)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5410)

			gen5411 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3529 := __args[0]
				_ = V3529
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3530 := __args[0]
					_ = V3530
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3531 := __args[0]
						_ = V3531
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1fst), V3529, V3530, V3531)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5412 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-fst"), gen5411)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5412)

			gen5413 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3539 := __args[0]
				_ = V3539
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3540 := __args[0]
					_ = V3540
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3541 := __args[0]
						_ = V3541
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1function), V3539, V3540, V3541)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5414 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-function"), gen5413)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5414)

			gen5415 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3549 := __args[0]
				_ = V3549
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3550 := __args[0]
					_ = V3550
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3551 := __args[0]
						_ = V3551
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1gensym), V3549, V3550, V3551)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5416 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-gensym"), gen5415)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5416)

			gen5417 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3559 := __args[0]
				_ = V3559
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3560 := __args[0]
					_ = V3560
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3561 := __args[0]
						_ = V3561
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5_1vector), V3559, V3560, V3561)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5418 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<-vector"), gen5417)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5418)

			gen5419 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3569 := __args[0]
				_ = V3569
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3570 := __args[0]
					_ = V3570
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3571 := __args[0]
						_ = V3571
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1vector_1_6), V3569, V3570, V3571)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5420 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-vector->"), gen5419)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5420)

			gen5421 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3579 := __args[0]
				_ = V3579
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3580 := __args[0]
					_ = V3580
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3581 := __args[0]
						_ = V3581
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1vector), V3579, V3580, V3581)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5422 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-vector"), gen5421)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5422)

			gen5423 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3589 := __args[0]
				_ = V3589
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3590 := __args[0]
					_ = V3590
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3591 := __args[0]
						_ = V3591
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1get_1time), V3589, V3590, V3591)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5424 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-get-time"), gen5423)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5424)

			gen5425 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3599 := __args[0]
				_ = V3599
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3600 := __args[0]
					_ = V3600
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3601 := __args[0]
						_ = V3601
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1hash), V3599, V3600, V3601)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5426 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-hash"), gen5425)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5426)

			gen5427 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3609 := __args[0]
				_ = V3609
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3610 := __args[0]
					_ = V3610
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3611 := __args[0]
						_ = V3611
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1head), V3609, V3610, V3611)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5428 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-head"), gen5427)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5428)

			gen5429 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3619 := __args[0]
				_ = V3619
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3620 := __args[0]
					_ = V3620
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3621 := __args[0]
						_ = V3621
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1hdv), V3619, V3620, V3621)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5430 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-hdv"), gen5429)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5430)

			gen5431 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3629 := __args[0]
				_ = V3629
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3630 := __args[0]
					_ = V3630
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3631 := __args[0]
						_ = V3631
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1hdstr), V3629, V3630, V3631)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5432 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-hdstr"), gen5431)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5432)

			gen5433 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3639 := __args[0]
				_ = V3639
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3640 := __args[0]
					_ = V3640
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3641 := __args[0]
						_ = V3641
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1if), V3639, V3640, V3641)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5434 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-if"), gen5433)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5434)

			gen5435 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3649 := __args[0]
				_ = V3649
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3650 := __args[0]
					_ = V3650
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3651 := __args[0]
						_ = V3651
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1it), V3649, V3650, V3651)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5436 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-it"), gen5435)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5436)

			gen5437 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3659 := __args[0]
				_ = V3659
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3660 := __args[0]
					_ = V3660
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3661 := __args[0]
						_ = V3661
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1implementation), V3659, V3660, V3661)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5438 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-implementation"), gen5437)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5438)

			gen5439 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3669 := __args[0]
				_ = V3669
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3670 := __args[0]
					_ = V3670
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3671 := __args[0]
						_ = V3671
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1include), V3669, V3670, V3671)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5440 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-include"), gen5439)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5440)

			gen5441 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3679 := __args[0]
				_ = V3679
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3680 := __args[0]
					_ = V3680
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3681 := __args[0]
						_ = V3681
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1include_1all_1but), V3679, V3680, V3681)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5442 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-include-all-but"), gen5441)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5442)

			gen5443 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3689 := __args[0]
				_ = V3689
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3690 := __args[0]
					_ = V3690
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3691 := __args[0]
						_ = V3691
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1inferences), V3689, V3690, V3691)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5444 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-inferences"), gen5443)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5444)

			gen5445 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3699 := __args[0]
				_ = V3699
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3700 := __args[0]
					_ = V3700
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3701 := __args[0]
						_ = V3701
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4insert), V3699, V3700, V3701)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5446 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.insert"), gen5445)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5446)

			gen5447 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3709 := __args[0]
				_ = V3709
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3710 := __args[0]
					_ = V3710
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3711 := __args[0]
						_ = V3711
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1integer_2), V3709, V3710, V3711)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5448 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-integer?"), gen5447)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5448)

			gen5449 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3719 := __args[0]
				_ = V3719
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3720 := __args[0]
					_ = V3720
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3721 := __args[0]
						_ = V3721
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1internal), V3719, V3720, V3721)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5450 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-internal"), gen5449)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5450)

			gen5451 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3729 := __args[0]
				_ = V3729
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3730 := __args[0]
					_ = V3730
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3731 := __args[0]
						_ = V3731
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1intersection), V3729, V3730, V3731)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5452 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-intersection"), gen5451)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5452)

			gen5453 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3739 := __args[0]
				_ = V3739
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3740 := __args[0]
					_ = V3740
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3741 := __args[0]
						_ = V3741
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1kill), V3739, V3740, V3741)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5454 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-kill"), gen5453)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5454)

			gen5455 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3749 := __args[0]
				_ = V3749
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3750 := __args[0]
					_ = V3750
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3751 := __args[0]
						_ = V3751
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1language), V3749, V3750, V3751)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5456 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-language"), gen5455)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5456)

			gen5457 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3759 := __args[0]
				_ = V3759
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3760 := __args[0]
					_ = V3760
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3761 := __args[0]
						_ = V3761
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1length), V3759, V3760, V3761)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5458 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-length"), gen5457)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5458)

			gen5459 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3769 := __args[0]
				_ = V3769
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3770 := __args[0]
					_ = V3770
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3771 := __args[0]
						_ = V3771
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1limit), V3769, V3770, V3771)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5460 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-limit"), gen5459)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5460)

			gen5461 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3779 := __args[0]
				_ = V3779
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3780 := __args[0]
					_ = V3780
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3781 := __args[0]
						_ = V3781
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1load), V3779, V3780, V3781)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5462 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-load"), gen5461)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5462)

			gen5463 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3789 := __args[0]
				_ = V3789
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3790 := __args[0]
					_ = V3790
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3791 := __args[0]
						_ = V3791
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1map), V3789, V3790, V3791)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5464 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-map"), gen5463)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5464)

			gen5465 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3799 := __args[0]
				_ = V3799
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3800 := __args[0]
					_ = V3800
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3801 := __args[0]
						_ = V3801
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1mapcan), V3799, V3800, V3801)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5466 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-mapcan"), gen5465)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5466)

			gen5467 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3809 := __args[0]
				_ = V3809
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3810 := __args[0]
					_ = V3810
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3811 := __args[0]
						_ = V3811
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1maxinferences), V3809, V3810, V3811)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5468 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-maxinferences"), gen5467)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5468)

			gen5469 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3819 := __args[0]
				_ = V3819
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3820 := __args[0]
					_ = V3820
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3821 := __args[0]
						_ = V3821
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1n_1_6string), V3819, V3820, V3821)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5470 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-n->string"), gen5469)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5470)

			gen5471 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3829 := __args[0]
				_ = V3829
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3830 := __args[0]
					_ = V3830
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3831 := __args[0]
						_ = V3831
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1nl), V3829, V3830, V3831)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5472 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-nl"), gen5471)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5472)

			gen5473 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3839 := __args[0]
				_ = V3839
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3840 := __args[0]
					_ = V3840
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3841 := __args[0]
						_ = V3841
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1not), V3839, V3840, V3841)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5474 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-not"), gen5473)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5474)

			gen5475 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3849 := __args[0]
				_ = V3849
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3850 := __args[0]
					_ = V3850
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3851 := __args[0]
						_ = V3851
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1nth), V3849, V3850, V3851)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5476 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-nth"), gen5475)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5476)

			gen5477 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3859 := __args[0]
				_ = V3859
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3860 := __args[0]
					_ = V3860
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3861 := __args[0]
						_ = V3861
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1number_2), V3859, V3860, V3861)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5478 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-number?"), gen5477)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5478)

			gen5479 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3869 := __args[0]
				_ = V3869
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3870 := __args[0]
					_ = V3870
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3871 := __args[0]
						_ = V3871
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1occurrences), V3869, V3870, V3871)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5480 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-occurrences"), gen5479)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5480)

			gen5481 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3879 := __args[0]
				_ = V3879
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3880 := __args[0]
					_ = V3880
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3881 := __args[0]
						_ = V3881
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1occurs_1check), V3879, V3880, V3881)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5482 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-occurs-check"), gen5481)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5482)

			gen5483 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3889 := __args[0]
				_ = V3889
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3890 := __args[0]
					_ = V3890
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3891 := __args[0]
						_ = V3891
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1optimise), V3889, V3890, V3891)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5484 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-optimise"), gen5483)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5484)

			gen5485 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3899 := __args[0]
				_ = V3899
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3900 := __args[0]
					_ = V3900
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3901 := __args[0]
						_ = V3901
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1or), V3899, V3900, V3901)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5486 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-or"), gen5485)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5486)

			gen5487 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3909 := __args[0]
				_ = V3909
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3910 := __args[0]
					_ = V3910
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3911 := __args[0]
						_ = V3911
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1os), V3909, V3910, V3911)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5488 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-os"), gen5487)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5488)

			gen5489 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3919 := __args[0]
				_ = V3919
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3920 := __args[0]
					_ = V3920
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3921 := __args[0]
						_ = V3921
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1package_2), V3919, V3920, V3921)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5490 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-package?"), gen5489)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5490)

			gen5491 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3929 := __args[0]
				_ = V3929
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3930 := __args[0]
					_ = V3930
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3931 := __args[0]
						_ = V3931
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1port), V3929, V3930, V3931)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5492 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-port"), gen5491)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5492)

			gen5493 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3939 := __args[0]
				_ = V3939
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3940 := __args[0]
					_ = V3940
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3941 := __args[0]
						_ = V3941
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1porters), V3939, V3940, V3941)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5494 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-porters"), gen5493)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5494)

			gen5495 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3949 := __args[0]
				_ = V3949
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3950 := __args[0]
					_ = V3950
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3951 := __args[0]
						_ = V3951
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1pos), V3949, V3950, V3951)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5496 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-pos"), gen5495)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5496)

			gen5497 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3959 := __args[0]
				_ = V3959
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3960 := __args[0]
					_ = V3960
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3961 := __args[0]
						_ = V3961
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1pr), V3959, V3960, V3961)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5498 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-pr"), gen5497)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5498)

			gen5499 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3969 := __args[0]
				_ = V3969
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3970 := __args[0]
					_ = V3970
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3971 := __args[0]
						_ = V3971
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1print), V3969, V3970, V3971)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5500 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-print"), gen5499)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5500)

			gen5501 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3979 := __args[0]
				_ = V3979
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3980 := __args[0]
					_ = V3980
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3981 := __args[0]
						_ = V3981
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1profile), V3979, V3980, V3981)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5502 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-profile"), gen5501)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5502)

			gen5503 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3989 := __args[0]
				_ = V3989
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V3990 := __args[0]
					_ = V3990
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V3991 := __args[0]
						_ = V3991
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1preclude), V3989, V3990, V3991)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5504 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-preclude"), gen5503)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5504)

			gen5505 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V3999 := __args[0]
				_ = V3999
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4000 := __args[0]
					_ = V4000
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4001 := __args[0]
						_ = V4001
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4proc_1nl), V3999, V4000, V4001)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5506 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.proc-nl"), gen5505)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5506)

			gen5507 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4009 := __args[0]
				_ = V4009
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4010 := __args[0]
					_ = V4010
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4011 := __args[0]
						_ = V4011
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1profile_1results), V4009, V4010, V4011)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5508 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-profile-results"), gen5507)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5508)

			gen5509 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4019 := __args[0]
				_ = V4019
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4020 := __args[0]
					_ = V4020
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4021 := __args[0]
						_ = V4021
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1protect), V4019, V4020, V4021)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5510 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-protect"), gen5509)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5510)

			gen5511 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4029 := __args[0]
				_ = V4029
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4030 := __args[0]
					_ = V4030
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4031 := __args[0]
						_ = V4031
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1preclude_1all_1but), V4029, V4030, V4031)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5512 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-preclude-all-but"), gen5511)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5512)

			gen5513 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4039 := __args[0]
				_ = V4039
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4040 := __args[0]
					_ = V4040
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4041 := __args[0]
						_ = V4041
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1shen_4prhush), V4039, V4040, V4041)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5514 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-shen.prhush"), gen5513)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5514)

			gen5515 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4049 := __args[0]
				_ = V4049
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4050 := __args[0]
					_ = V4050
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4051 := __args[0]
						_ = V4051
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1ps), V4049, V4050, V4051)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5516 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-ps"), gen5515)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5516)

			gen5517 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4059 := __args[0]
				_ = V4059
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4060 := __args[0]
					_ = V4060
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4061 := __args[0]
						_ = V4061
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read), V4059, V4060, V4061)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5518 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read"), gen5517)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5518)

			gen5519 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4069 := __args[0]
				_ = V4069
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4070 := __args[0]
					_ = V4070
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4071 := __args[0]
						_ = V4071
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1byte), V4069, V4070, V4071)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5520 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-byte"), gen5519)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5520)

			gen5521 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4079 := __args[0]
				_ = V4079
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4080 := __args[0]
					_ = V4080
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4081 := __args[0]
						_ = V4081
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1file_1as_1bytelist), V4079, V4080, V4081)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5522 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-file-as-bytelist"), gen5521)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5522)

			gen5523 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4089 := __args[0]
				_ = V4089
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4090 := __args[0]
					_ = V4090
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4091 := __args[0]
						_ = V4091
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1file_1as_1string), V4089, V4090, V4091)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5524 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-file-as-string"), gen5523)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5524)

			gen5525 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4099 := __args[0]
				_ = V4099
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4100 := __args[0]
					_ = V4100
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4101 := __args[0]
						_ = V4101
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1file), V4099, V4100, V4101)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5526 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-file"), gen5525)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5526)

			gen5527 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4109 := __args[0]
				_ = V4109
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4110 := __args[0]
					_ = V4110
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4111 := __args[0]
						_ = V4111
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1read_1from_1string), V4109, V4110, V4111)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5528 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-read-from-string"), gen5527)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5528)

			gen5529 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4119 := __args[0]
				_ = V4119
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4120 := __args[0]
					_ = V4120
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4121 := __args[0]
						_ = V4121
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1release), V4119, V4120, V4121)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5530 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-release"), gen5529)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5530)

			gen5531 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4129 := __args[0]
				_ = V4129
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4130 := __args[0]
					_ = V4130
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4131 := __args[0]
						_ = V4131
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1remove), V4129, V4130, V4131)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5532 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-remove"), gen5531)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5532)

			gen5533 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4139 := __args[0]
				_ = V4139
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4140 := __args[0]
					_ = V4140
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4141 := __args[0]
						_ = V4141
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1reverse), V4139, V4140, V4141)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5534 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-reverse"), gen5533)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5534)

			gen5535 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4149 := __args[0]
				_ = V4149
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4150 := __args[0]
					_ = V4150
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4151 := __args[0]
						_ = V4151
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1simple_1error), V4149, V4150, V4151)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5536 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-simple-error"), gen5535)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5536)

			gen5537 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4159 := __args[0]
				_ = V4159
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4160 := __args[0]
					_ = V4160
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4161 := __args[0]
						_ = V4161
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1snd), V4159, V4160, V4161)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5538 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-snd"), gen5537)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5538)

			gen5539 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4169 := __args[0]
				_ = V4169
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4170 := __args[0]
					_ = V4170
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4171 := __args[0]
						_ = V4171
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1specialise), V4169, V4170, V4171)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5540 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-specialise"), gen5539)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5540)

			gen5541 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4179 := __args[0]
				_ = V4179
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4180 := __args[0]
					_ = V4180
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4181 := __args[0]
						_ = V4181
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1spy), V4179, V4180, V4181)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5542 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-spy"), gen5541)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5542)

			gen5543 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4189 := __args[0]
				_ = V4189
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4190 := __args[0]
					_ = V4190
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4191 := __args[0]
						_ = V4191
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1step), V4189, V4190, V4191)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5544 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-step"), gen5543)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5544)

			gen5545 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4199 := __args[0]
				_ = V4199
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4200 := __args[0]
					_ = V4200
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4201 := __args[0]
						_ = V4201
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1stinput), V4199, V4200, V4201)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5546 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-stinput"), gen5545)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5546)

			gen5547 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4209 := __args[0]
				_ = V4209
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4210 := __args[0]
					_ = V4210
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4211 := __args[0]
						_ = V4211
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1sterror), V4209, V4210, V4211)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5548 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-sterror"), gen5547)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5548)

			gen5549 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4219 := __args[0]
				_ = V4219
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4220 := __args[0]
					_ = V4220
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4221 := __args[0]
						_ = V4221
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1stoutput), V4219, V4220, V4221)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5550 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-stoutput"), gen5549)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5550)

			gen5551 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4229 := __args[0]
				_ = V4229
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4230 := __args[0]
					_ = V4230
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4231 := __args[0]
						_ = V4231
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1string_2), V4229, V4230, V4231)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5552 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-string?"), gen5551)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5552)

			gen5553 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4239 := __args[0]
				_ = V4239
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4240 := __args[0]
					_ = V4240
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4241 := __args[0]
						_ = V4241
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1str), V4239, V4240, V4241)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5554 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-str"), gen5553)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5554)

			gen5555 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4249 := __args[0]
				_ = V4249
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4250 := __args[0]
					_ = V4250
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4251 := __args[0]
						_ = V4251
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1string_1_6n), V4249, V4250, V4251)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5556 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-string->n"), gen5555)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5556)

			gen5557 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4259 := __args[0]
				_ = V4259
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4260 := __args[0]
					_ = V4260
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4261 := __args[0]
						_ = V4261
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1string_1_6symbol), V4259, V4260, V4261)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5558 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-string->symbol"), gen5557)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5558)

			gen5559 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4269 := __args[0]
				_ = V4269
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4270 := __args[0]
					_ = V4270
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4271 := __args[0]
						_ = V4271
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1sum), V4269, V4270, V4271)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5560 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-sum"), gen5559)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5560)

			gen5561 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4279 := __args[0]
				_ = V4279
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4280 := __args[0]
					_ = V4280
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4281 := __args[0]
						_ = V4281
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1symbol_2), V4279, V4280, V4281)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5562 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-symbol?"), gen5561)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5562)

			gen5563 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4289 := __args[0]
				_ = V4289
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4290 := __args[0]
					_ = V4290
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4291 := __args[0]
						_ = V4291
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1systemf), V4289, V4290, V4291)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5564 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-systemf"), gen5563)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5564)

			gen5565 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4299 := __args[0]
				_ = V4299
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4300 := __args[0]
					_ = V4300
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4301 := __args[0]
						_ = V4301
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tail), V4299, V4300, V4301)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5566 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tail"), gen5565)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5566)

			gen5567 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4309 := __args[0]
				_ = V4309
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4310 := __args[0]
					_ = V4310
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4311 := __args[0]
						_ = V4311
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tlstr), V4309, V4310, V4311)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5568 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tlstr"), gen5567)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5568)

			gen5569 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4319 := __args[0]
				_ = V4319
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4320 := __args[0]
					_ = V4320
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4321 := __args[0]
						_ = V4321
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tlv), V4319, V4320, V4321)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5570 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tlv"), gen5569)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5570)

			gen5571 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4329 := __args[0]
				_ = V4329
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4330 := __args[0]
					_ = V4330
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4331 := __args[0]
						_ = V4331
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tc), V4329, V4330, V4331)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5572 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tc"), gen5571)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5572)

			gen5573 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4339 := __args[0]
				_ = V4339
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4340 := __args[0]
					_ = V4340
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4341 := __args[0]
						_ = V4341
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tc_2), V4339, V4340, V4341)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5574 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tc?"), gen5573)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5574)

			gen5575 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4349 := __args[0]
				_ = V4349
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4350 := __args[0]
					_ = V4350
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4351 := __args[0]
						_ = V4351
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1thaw), V4349, V4350, V4351)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5576 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-thaw"), gen5575)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5576)

			gen5577 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4359 := __args[0]
				_ = V4359
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4360 := __args[0]
					_ = V4360
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4361 := __args[0]
						_ = V4361
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1track), V4359, V4360, V4361)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5578 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-track"), gen5577)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5578)

			gen5579 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4369 := __args[0]
				_ = V4369
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4370 := __args[0]
					_ = V4370
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4371 := __args[0]
						_ = V4371
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1trap_1error), V4369, V4370, V4371)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5580 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-trap-error"), gen5579)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5580)

			gen5581 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4379 := __args[0]
				_ = V4379
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4380 := __args[0]
					_ = V4380
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4381 := __args[0]
						_ = V4381
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1tuple_2), V4379, V4380, V4381)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5582 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-tuple?"), gen5581)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5582)

			gen5583 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4389 := __args[0]
				_ = V4389
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4390 := __args[0]
					_ = V4390
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4391 := __args[0]
						_ = V4391
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1undefmacro), V4389, V4390, V4391)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5584 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-undefmacro"), gen5583)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5584)

			gen5585 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4399 := __args[0]
				_ = V4399
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4400 := __args[0]
					_ = V4400
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4401 := __args[0]
						_ = V4401
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1union), V4399, V4400, V4401)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5586 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-union"), gen5585)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5586)

			gen5587 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4409 := __args[0]
				_ = V4409
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4410 := __args[0]
					_ = V4410
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4411 := __args[0]
						_ = V4411
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1unprofile), V4409, V4410, V4411)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5588 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-unprofile"), gen5587)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5588)

			gen5589 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4419 := __args[0]
				_ = V4419
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4420 := __args[0]
					_ = V4420
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4421 := __args[0]
						_ = V4421
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1untrack), V4419, V4420, V4421)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5590 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-untrack"), gen5589)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5590)

			gen5591 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4429 := __args[0]
				_ = V4429
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4430 := __args[0]
					_ = V4430
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4431 := __args[0]
						_ = V4431
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1unspecialise), V4429, V4430, V4431)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5592 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-unspecialise"), gen5591)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5592)

			gen5593 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4439 := __args[0]
				_ = V4439
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4440 := __args[0]
					_ = V4440
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4441 := __args[0]
						_ = V4441
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1variable_2), V4439, V4440, V4441)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5594 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-variable?"), gen5593)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5594)

			gen5595 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4449 := __args[0]
				_ = V4449
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4450 := __args[0]
					_ = V4450
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4451 := __args[0]
						_ = V4451
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1vector_2), V4449, V4450, V4451)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5596 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-vector?"), gen5595)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5596)

			gen5597 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4459 := __args[0]
				_ = V4459
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4460 := __args[0]
					_ = V4460
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4461 := __args[0]
						_ = V4461
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1version), V4459, V4460, V4461)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5598 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-version"), gen5597)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5598)

			gen5599 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4469 := __args[0]
				_ = V4469
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4470 := __args[0]
					_ = V4470
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4471 := __args[0]
						_ = V4471
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1write_1to_1file), V4469, V4470, V4471)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5600 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-write-to-file"), gen5599)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5600)

			gen5601 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4479 := __args[0]
				_ = V4479
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4480 := __args[0]
					_ = V4480
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4481 := __args[0]
						_ = V4481
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1write_1byte), V4479, V4480, V4481)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5602 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-write-byte"), gen5601)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5602)

			gen5603 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4489 := __args[0]
				_ = V4489
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4490 := __args[0]
					_ = V4490
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4491 := __args[0]
						_ = V4491
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1y_1or_1n_2), V4489, V4490, V4491)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5604 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-y-or-n?"), gen5603)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5604)

			gen5605 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4499 := __args[0]
				_ = V4499
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4500 := __args[0]
					_ = V4500
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4501 := __args[0]
						_ = V4501
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_6), V4499, V4500, V4501)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5606 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of->"), gen5605)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5606)

			gen5607 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4509 := __args[0]
				_ = V4509
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4510 := __args[0]
					_ = V4510
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4511 := __args[0]
						_ = V4511
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5), V4509, V4510, V4511)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5608 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<"), gen5607)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5608)

			gen5609 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4519 := __args[0]
				_ = V4519
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4520 := __args[0]
					_ = V4520
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4521 := __args[0]
						_ = V4521
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_6_a), V4519, V4520, V4521)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5610 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of->="), gen5609)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5610)

			gen5611 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4529 := __args[0]
				_ = V4529
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4530 := __args[0]
					_ = V4530
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4531 := __args[0]
						_ = V4531
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_5_a), V4529, V4530, V4531)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5612 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-<="), gen5611)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5612)

			gen5613 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4539 := __args[0]
				_ = V4539
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4540 := __args[0]
					_ = V4540
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4541 := __args[0]
						_ = V4541
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_a), V4539, V4540, V4541)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5614 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-="), gen5613)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5614)

			gen5615 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4549 := __args[0]
				_ = V4549
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4550 := __args[0]
					_ = V4550
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4551 := __args[0]
						_ = V4551
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_7), V4549, V4550, V4551)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5616 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-+"), gen5615)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5616)

			gen5617 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4559 := __args[0]
				_ = V4559
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4560 := __args[0]
					_ = V4560
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4561 := __args[0]
						_ = V4561
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_c), V4559, V4560, V4561)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5618 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-/"), gen5617)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5618)

			gen5619 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4569 := __args[0]
				_ = V4569
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4570 := __args[0]
					_ = V4570
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4571 := __args[0]
						_ = V4571
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_1), V4569, V4570, V4571)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5620 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of--"), gen5619)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5620)

			gen5621 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4579 := __args[0]
				_ = V4579
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4580 := __args[0]
					_ = V4580
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4581 := __args[0]
						_ = V4581
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_d), V4579, V4580, V4581)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5622 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-*"), gen5621)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5622)

			gen5623 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V4589 := __args[0]
				_ = V4589
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V4590 := __args[0]
					_ = V4590
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V4591 := __args[0]
						_ = V4591
						__e.TailApply(ShenFunc(symshen_4type_1signature_1of_1_a_a), V4589, V4590, V4591)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5624 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.type-signature-of-=="), gen5623)

			__e.TailApply(ShenFunc(symshen_4set_1lambda_1form_1entry), gen5624)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-signedfunc-lambda-forms"), gen5625)

		gen5927 := MakeNative(func(__e Evaluator, __args ...Obj) {
			gen5626 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4datatype_1error), X)

				return
			}, 1)
			gen5627 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.datatype-error"), gen5626)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5627)

			gen5628 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4tuple), X)

				return
			}, 1)
			gen5629 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.tuple"), gen5628)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5629)

			gen5630 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4pvar), X)

				return
			}, 1)
			gen5631 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.pvar"), gen5630)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5631)

			gen5632 := MakeNative(func(__e Evaluator, __args ...Obj) {
				X := __args[0]
				_ = X
				__e.TailApply(ShenFunc(symshen_4dictionary), X)

				return
			}, 1)
			gen5633 := Call(__e, ShenFunc(symcons), MakeSymbol("shen.dictionary"), gen5632)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5633)

			gen5634 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V476 := __args[0]
				_ = V476
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V477 := __args[0]
					_ = V477
					__e.TailApply(ShenFunc(sym_8v), V476, V477)

					return
				}, 1))
				return
			}, 1)
			gen5635 := Call(__e, ShenFunc(symcons), MakeSymbol("@v"), gen5634)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5635)

			gen5636 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V478 := __args[0]
				_ = V478
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V479 := __args[0]
					_ = V479
					__e.TailApply(ShenFunc(sym_8p), V478, V479)

					return
				}, 1))
				return
			}, 1)
			gen5637 := Call(__e, ShenFunc(symcons), MakeSymbol("@p"), gen5636)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5637)

			gen5638 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V480 := __args[0]
				_ = V480
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V481 := __args[0]
					_ = V481
					__e.TailApply(ShenFunc(sym_8s), V480, V481)

					return
				}, 1))
				return
			}, 1)
			gen5639 := Call(__e, ShenFunc(symcons), MakeSymbol("@s"), gen5638)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5639)

			gen5640 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V482 := __args[0]
				_ = V482
				__e.TailApply(ShenFunc(sym_5e_6), V482)

				return
			}, 1)
			gen5641 := Call(__e, ShenFunc(symcons), MakeSymbol("<e>"), gen5640)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5641)

			gen5642 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V483 := __args[0]
				_ = V483
				__e.TailApply(ShenFunc(sym_5_b_6), V483)

				return
			}, 1)
			gen5643 := Call(__e, ShenFunc(symcons), MakeSymbol("<!>"), gen5642)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5643)

			gen5644 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V484 := __args[0]
				_ = V484
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V485 := __args[0]
					_ = V485
					__e.TailApply(ShenFunc(sym_a_a), V484, V485)

					return
				}, 1))
				return
			}, 1)
			gen5645 := Call(__e, ShenFunc(symcons), MakeSymbol("=="), gen5644)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5645)

			gen5646 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V486 := __args[0]
				_ = V486
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V487 := __args[0]
					_ = V487
					__e.TailApply(ShenFunc(sym_a), V486, V487)

					return
				}, 1))
				return
			}, 1)
			gen5647 := Call(__e, ShenFunc(symcons), MakeSymbol("="), gen5646)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5647)

			gen5648 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V488 := __args[0]
				_ = V488
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V489 := __args[0]
					_ = V489
					__e.TailApply(ShenFunc(sym_6_a), V488, V489)

					return
				}, 1))
				return
			}, 1)
			gen5649 := Call(__e, ShenFunc(symcons), MakeSymbol(">="), gen5648)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5649)

			gen5650 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V490 := __args[0]
				_ = V490
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V491 := __args[0]
					_ = V491
					__e.TailApply(ShenFunc(sym_6), V490, V491)

					return
				}, 1))
				return
			}, 1)
			gen5651 := Call(__e, ShenFunc(symcons), MakeSymbol(">"), gen5650)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5651)

			gen5652 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V492 := __args[0]
				_ = V492
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V493 := __args[0]
					_ = V493
					__e.TailApply(ShenFunc(sym_1), V492, V493)

					return
				}, 1))
				return
			}, 1)
			gen5653 := Call(__e, ShenFunc(symcons), MakeSymbol("-"), gen5652)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5653)

			gen5654 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V494 := __args[0]
				_ = V494
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V495 := __args[0]
					_ = V495
					__e.TailApply(ShenFunc(sym_c), V494, V495)

					return
				}, 1))
				return
			}, 1)
			gen5655 := Call(__e, ShenFunc(symcons), MakeSymbol("/"), gen5654)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5655)

			gen5656 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V496 := __args[0]
				_ = V496
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V497 := __args[0]
					_ = V497
					__e.TailApply(ShenFunc(sym_d), V496, V497)

					return
				}, 1))
				return
			}, 1)
			gen5657 := Call(__e, ShenFunc(symcons), MakeSymbol("*"), gen5656)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5657)

			gen5658 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V498 := __args[0]
				_ = V498
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V499 := __args[0]
					_ = V499
					__e.TailApply(ShenFunc(sym_7), V498, V499)

					return
				}, 1))
				return
			}, 1)
			gen5659 := Call(__e, ShenFunc(symcons), MakeSymbol("+"), gen5658)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5659)

			gen5660 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V500 := __args[0]
				_ = V500
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V501 := __args[0]
					_ = V501
					__e.TailApply(ShenFunc(sym_5_a), V500, V501)

					return
				}, 1))
				return
			}, 1)
			gen5661 := Call(__e, ShenFunc(symcons), MakeSymbol("<="), gen5660)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5661)

			gen5662 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V502 := __args[0]
				_ = V502
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V503 := __args[0]
					_ = V503
					__e.TailApply(ShenFunc(sym_5), V502, V503)

					return
				}, 1))
				return
			}, 1)
			gen5663 := Call(__e, ShenFunc(symcons), MakeSymbol("<"), gen5662)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5663)

			gen5664 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V504 := __args[0]
				_ = V504
				__e.TailApply(ShenFunc(symy_1or_1n_2), V504)

				return
			}, 1)
			gen5665 := Call(__e, ShenFunc(symcons), MakeSymbol("y-or-n?"), gen5664)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5665)

			gen5666 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V505 := __args[0]
				_ = V505
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V506 := __args[0]
					_ = V506
					__e.TailApply(ShenFunc(symwrite_1to_1file), V505, V506)

					return
				}, 1))
				return
			}, 1)
			gen5667 := Call(__e, ShenFunc(symcons), MakeSymbol("write-to-file"), gen5666)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5667)

			gen5668 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V507 := __args[0]
				_ = V507
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V508 := __args[0]
					_ = V508
					__e.TailApply(ShenFunc(symwrite_1byte), V507, V508)

					return
				}, 1))
				return
			}, 1)
			gen5669 := Call(__e, ShenFunc(symcons), MakeSymbol("write-byte"), gen5668)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5669)

			gen5670 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V509 := __args[0]
				_ = V509
				__e.TailApply(ShenFunc(symvariable_2), V509)

				return
			}, 1)
			gen5671 := Call(__e, ShenFunc(symcons), MakeSymbol("variable?"), gen5670)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5671)

			gen5672 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V510 := __args[0]
				_ = V510
				__e.TailApply(ShenFunc(symvalue), V510)

				return
			}, 1)
			gen5673 := Call(__e, ShenFunc(symcons), MakeSymbol("value"), gen5672)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5673)

			gen5674 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V511 := __args[0]
				_ = V511
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V512 := __args[0]
					_ = V512
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V513 := __args[0]
						_ = V513
						__e.TailApply(ShenFunc(symvector_1_6), V511, V512, V513)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5675 := Call(__e, ShenFunc(symcons), MakeSymbol("vector->"), gen5674)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5675)

			gen5676 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V514 := __args[0]
				_ = V514
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V515 := __args[0]
					_ = V515
					__e.TailApply(ShenFunc(sym_5_1vector), V514, V515)

					return
				}, 1))
				return
			}, 1)
			gen5677 := Call(__e, ShenFunc(symcons), MakeSymbol("<-vector"), gen5676)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5677)

			gen5678 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V516 := __args[0]
				_ = V516
				__e.TailApply(ShenFunc(symvector), V516)

				return
			}, 1)
			gen5679 := Call(__e, ShenFunc(symcons), MakeSymbol("vector"), gen5678)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5679)

			gen5680 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V517 := __args[0]
				_ = V517
				__e.TailApply(ShenFunc(symvector_2), V517)

				return
			}, 1)
			gen5681 := Call(__e, ShenFunc(symcons), MakeSymbol("vector?"), gen5680)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5681)

			gen5682 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V518 := __args[0]
				_ = V518
				__e.TailApply(ShenFunc(symunspecialise), V518)

				return
			}, 1)
			gen5683 := Call(__e, ShenFunc(symcons), MakeSymbol("unspecialise"), gen5682)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5683)

			gen5684 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V519 := __args[0]
				_ = V519
				__e.TailApply(ShenFunc(symuntrack), V519)

				return
			}, 1)
			gen5685 := Call(__e, ShenFunc(symcons), MakeSymbol("untrack"), gen5684)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5685)

			gen5686 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V520 := __args[0]
				_ = V520
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V521 := __args[0]
					_ = V521
					__e.TailApply(ShenFunc(symunion), V520, V521)

					return
				}, 1))
				return
			}, 1)
			gen5687 := Call(__e, ShenFunc(symcons), MakeSymbol("union"), gen5686)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5687)

			gen5688 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V522 := __args[0]
				_ = V522
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V523 := __args[0]
					_ = V523
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V524 := __args[0]
						_ = V524
						__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
							V525 := __args[0]
							_ = V525
							__e.TailApply(ShenFunc(symunify), V522, V523, V524, V525)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5689 := Call(__e, ShenFunc(symcons), MakeSymbol("unify"), gen5688)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5689)

			gen5690 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V526 := __args[0]
				_ = V526
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V527 := __args[0]
					_ = V527
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V528 := __args[0]
						_ = V528
						__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
							V529 := __args[0]
							_ = V529
							__e.TailApply(ShenFunc(symunify_b), V526, V527, V528, V529)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5691 := Call(__e, ShenFunc(symcons), MakeSymbol("unify!"), gen5690)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5691)

			gen5692 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V530 := __args[0]
				_ = V530
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V531 := __args[0]
					_ = V531
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V532 := __args[0]
						_ = V532
						__e.TailApply(ShenFunc(symunput), V530, V531, V532)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5693 := Call(__e, ShenFunc(symcons), MakeSymbol("unput"), gen5692)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5693)

			gen5694 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V533 := __args[0]
				_ = V533
				__e.TailApply(ShenFunc(symunprofile), V533)

				return
			}, 1)
			gen5695 := Call(__e, ShenFunc(symcons), MakeSymbol("unprofile"), gen5694)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5695)

			gen5696 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V534 := __args[0]
				_ = V534
				__e.TailApply(ShenFunc(symundefmacro), V534)

				return
			}, 1)
			gen5697 := Call(__e, ShenFunc(symcons), MakeSymbol("undefmacro"), gen5696)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5697)

			gen5698 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V535 := __args[0]
				_ = V535
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V536 := __args[0]
					_ = V536
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V537 := __args[0]
						_ = V537
						__e.TailApply(ShenFunc(symreturn), V535, V536, V537)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5699 := Call(__e, ShenFunc(symcons), MakeSymbol("return"), gen5698)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5699)

			gen5700 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V538 := __args[0]
				_ = V538
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V539 := __args[0]
					_ = V539
					__e.TailApply(ShenFunc(symtype), V538, V539)

					return
				}, 1))
				return
			}, 1)
			gen5701 := Call(__e, ShenFunc(symcons), MakeSymbol("type"), gen5700)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5701)

			gen5702 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V540 := __args[0]
				_ = V540
				__e.TailApply(ShenFunc(symtuple_2), V540)

				return
			}, 1)
			gen5703 := Call(__e, ShenFunc(symcons), MakeSymbol("tuple?"), gen5702)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5703)

			gen5705 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V541 := __args[0]
				_ = V541
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V542 := __args[0]
					_ = V542
					gen5704 := MakeNative(func(__e Evaluator, __args ...Obj) {
						__e.Return(V541)
						return
					}, 0)
					__e.Return(Try(__e, gen5704).Catch(V542))
					return

				}, 1))
				return
			}, 1)
			gen5706 := Call(__e, ShenFunc(symcons), MakeSymbol("trap-error"), gen5705)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5706)

			gen5707 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V543 := __args[0]
				_ = V543
				__e.TailApply(ShenFunc(symtrack), V543)

				return
			}, 1)
			gen5708 := Call(__e, ShenFunc(symcons), MakeSymbol("track"), gen5707)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5708)

			gen5709 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V544 := __args[0]
				_ = V544
				__e.TailApply(ShenFunc(symthaw), V544)

				return
			}, 1)
			gen5710 := Call(__e, ShenFunc(symcons), MakeSymbol("thaw"), gen5709)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5710)

			gen5711 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V545 := __args[0]
				_ = V545
				__e.TailApply(ShenFunc(symtc), V545)

				return
			}, 1)
			gen5712 := Call(__e, ShenFunc(symcons), MakeSymbol("tc"), gen5711)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5712)

			gen5713 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V546 := __args[0]
				_ = V546
				__e.TailApply(ShenFunc(symtl), V546)

				return
			}, 1)
			gen5714 := Call(__e, ShenFunc(symcons), MakeSymbol("tl"), gen5713)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5714)

			gen5715 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V547 := __args[0]
				_ = V547
				__e.TailApply(ShenFunc(symtlstr), V547)

				return
			}, 1)
			gen5716 := Call(__e, ShenFunc(symcons), MakeSymbol("tlstr"), gen5715)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5716)

			gen5717 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V548 := __args[0]
				_ = V548
				__e.TailApply(ShenFunc(symtail), V548)

				return
			}, 1)
			gen5718 := Call(__e, ShenFunc(symcons), MakeSymbol("tail"), gen5717)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5718)

			gen5719 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V549 := __args[0]
				_ = V549
				__e.TailApply(ShenFunc(symsystemf), V549)

				return
			}, 1)
			gen5720 := Call(__e, ShenFunc(symcons), MakeSymbol("systemf"), gen5719)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5720)

			gen5721 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V550 := __args[0]
				_ = V550
				__e.TailApply(ShenFunc(symsymbol_2), V550)

				return
			}, 1)
			gen5722 := Call(__e, ShenFunc(symcons), MakeSymbol("symbol?"), gen5721)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5722)

			gen5723 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V551 := __args[0]
				_ = V551
				__e.TailApply(ShenFunc(symstring_1_6symbol), V551)

				return
			}, 1)
			gen5724 := Call(__e, ShenFunc(symcons), MakeSymbol("string->symbol"), gen5723)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5724)

			gen5725 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V552 := __args[0]
				_ = V552
				__e.TailApply(ShenFunc(symsum), V552)

				return
			}, 1)
			gen5726 := Call(__e, ShenFunc(symcons), MakeSymbol("sum"), gen5725)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5726)

			gen5727 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V553 := __args[0]
				_ = V553
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V554 := __args[0]
					_ = V554
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V555 := __args[0]
						_ = V555
						__e.TailApply(ShenFunc(symsubst), V553, V554, V555)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5728 := Call(__e, ShenFunc(symcons), MakeSymbol("subst"), gen5727)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5728)

			gen5729 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V556 := __args[0]
				_ = V556
				__e.TailApply(ShenFunc(symstring_2), V556)

				return
			}, 1)
			gen5730 := Call(__e, ShenFunc(symcons), MakeSymbol("string?"), gen5729)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5730)

			gen5731 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V557 := __args[0]
				_ = V557
				__e.TailApply(ShenFunc(symstring_1_6n), V557)

				return
			}, 1)
			gen5732 := Call(__e, ShenFunc(symcons), MakeSymbol("string->n"), gen5731)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5732)

			gen5733 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V558 := __args[0]
				_ = V558
				__e.TailApply(ShenFunc(symstep), V558)

				return
			}, 1)
			gen5734 := Call(__e, ShenFunc(symcons), MakeSymbol("step"), gen5733)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5734)

			gen5735 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V559 := __args[0]
				_ = V559
				__e.TailApply(ShenFunc(symspy), V559)

				return
			}, 1)
			gen5736 := Call(__e, ShenFunc(symcons), MakeSymbol("spy"), gen5735)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5736)

			gen5737 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V560 := __args[0]
				_ = V560
				__e.TailApply(ShenFunc(symspecialise), V560)

				return
			}, 1)
			gen5738 := Call(__e, ShenFunc(symcons), MakeSymbol("specialise"), gen5737)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5738)

			gen5739 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V561 := __args[0]
				_ = V561
				__e.TailApply(ShenFunc(symsnd), V561)

				return
			}, 1)
			gen5740 := Call(__e, ShenFunc(symcons), MakeSymbol("snd"), gen5739)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5740)

			gen5741 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V562 := __args[0]
				_ = V562
				__e.TailApply(ShenFunc(symsimple_1error), V562)

				return
			}, 1)
			gen5742 := Call(__e, ShenFunc(symcons), MakeSymbol("simple-error"), gen5741)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5742)

			gen5743 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V563 := __args[0]
				_ = V563
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V564 := __args[0]
					_ = V564
					__e.TailApply(ShenFunc(symset), V563, V564)

					return
				}, 1))
				return
			}, 1)
			gen5744 := Call(__e, ShenFunc(symcons), MakeSymbol("set"), gen5743)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5744)

			gen5745 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V565 := __args[0]
				_ = V565
				__e.TailApply(ShenFunc(symstr), V565)

				return
			}, 1)
			gen5746 := Call(__e, ShenFunc(symcons), MakeSymbol("str"), gen5745)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5746)

			gen5747 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V566 := __args[0]
				_ = V566
				__e.TailApply(ShenFunc(symreverse), V566)

				return
			}, 1)
			gen5748 := Call(__e, ShenFunc(symcons), MakeSymbol("reverse"), gen5747)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5748)

			gen5749 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V567 := __args[0]
				_ = V567
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V568 := __args[0]
					_ = V568
					__e.TailApply(ShenFunc(symremove), V567, V568)

					return
				}, 1))
				return
			}, 1)
			gen5750 := Call(__e, ShenFunc(symcons), MakeSymbol("remove"), gen5749)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5750)

			gen5751 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V569 := __args[0]
				_ = V569
				__e.TailApply(ShenFunc(symread), V569)

				return
			}, 1)
			gen5752 := Call(__e, ShenFunc(symcons), MakeSymbol("read"), gen5751)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5752)

			gen5753 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V570 := __args[0]
				_ = V570
				__e.TailApply(ShenFunc(symread_1file), V570)

				return
			}, 1)
			gen5754 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file"), gen5753)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5754)

			gen5755 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V571 := __args[0]
				_ = V571
				__e.TailApply(ShenFunc(symread_1file_1as_1bytelist), V571)

				return
			}, 1)
			gen5756 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-bytelist"), gen5755)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5756)

			gen5757 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V572 := __args[0]
				_ = V572
				__e.TailApply(ShenFunc(symread_1file_1as_1string), V572)

				return
			}, 1)
			gen5758 := Call(__e, ShenFunc(symcons), MakeSymbol("read-file-as-string"), gen5757)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5758)

			gen5759 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V573 := __args[0]
				_ = V573
				__e.TailApply(ShenFunc(symread_1byte), V573)

				return
			}, 1)
			gen5760 := Call(__e, ShenFunc(symcons), MakeSymbol("read-byte"), gen5759)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5760)

			gen5761 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V574 := __args[0]
				_ = V574
				__e.TailApply(ShenFunc(symread_1from_1string), V574)

				return
			}, 1)
			gen5762 := Call(__e, ShenFunc(symcons), MakeSymbol("read-from-string"), gen5761)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5762)

			gen5763 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V575 := __args[0]
				_ = V575
				__e.TailApply(ShenFunc(sympackage_2), V575)

				return
			}, 1)
			gen5764 := Call(__e, ShenFunc(symcons), MakeSymbol("package?"), gen5763)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5764)

			gen5765 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V576 := __args[0]
				_ = V576
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V577 := __args[0]
					_ = V577
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V578 := __args[0]
						_ = V578
						__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
							V579 := __args[0]
							_ = V579
							__e.TailApply(ShenFunc(symput), V576, V577, V578, V579)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5766 := Call(__e, ShenFunc(symcons), MakeSymbol("put"), gen5765)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5766)

			gen5767 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V580 := __args[0]
				_ = V580
				__e.TailApply(ShenFunc(sympreclude), V580)

				return
			}, 1)
			gen5768 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude"), gen5767)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5768)

			gen5769 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V581 := __args[0]
				_ = V581
				__e.TailApply(ShenFunc(sympreclude_1all_1but), V581)

				return
			}, 1)
			gen5770 := Call(__e, ShenFunc(symcons), MakeSymbol("preclude-all-but"), gen5769)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5770)

			gen5771 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V582 := __args[0]
				_ = V582
				__e.TailApply(ShenFunc(symps), V582)

				return
			}, 1)
			gen5772 := Call(__e, ShenFunc(symcons), MakeSymbol("ps"), gen5771)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5772)

			gen5773 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V583 := __args[0]
				_ = V583
				__e.TailApply(ShenFunc(symprotect), V583)

				return
			}, 1)
			gen5774 := Call(__e, ShenFunc(symcons), MakeSymbol("protect"), gen5773)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5774)

			gen5775 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V584 := __args[0]
				_ = V584
				__e.TailApply(ShenFunc(symprofile_1results), V584)

				return
			}, 1)
			gen5776 := Call(__e, ShenFunc(symcons), MakeSymbol("profile-results"), gen5775)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5776)

			gen5777 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V585 := __args[0]
				_ = V585
				__e.TailApply(ShenFunc(symprofile), V585)

				return
			}, 1)
			gen5778 := Call(__e, ShenFunc(symcons), MakeSymbol("profile"), gen5777)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5778)

			gen5779 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V586 := __args[0]
				_ = V586
				__e.TailApply(ShenFunc(symprint), V586)

				return
			}, 1)
			gen5780 := Call(__e, ShenFunc(symcons), MakeSymbol("print"), gen5779)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5780)

			gen5781 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V587 := __args[0]
				_ = V587
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V588 := __args[0]
					_ = V588
					__e.TailApply(ShenFunc(sympr), V587, V588)

					return
				}, 1))
				return
			}, 1)
			gen5782 := Call(__e, ShenFunc(symcons), MakeSymbol("pr"), gen5781)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5782)

			gen5783 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V589 := __args[0]
				_ = V589
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V590 := __args[0]
					_ = V590
					__e.TailApply(ShenFunc(sympos), V589, V590)

					return
				}, 1))
				return
			}, 1)
			gen5784 := Call(__e, ShenFunc(symcons), MakeSymbol("pos"), gen5783)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5784)

			gen5785 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V591 := __args[0]
				_ = V591
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V592 := __args[0]
					_ = V592
					if True == V591 {
						__e.Return(True)
						return
					} else {
						if True == V592 {
							__e.Return(True)
							return
						} else {
							__e.Return(False)
							return
						}
					}
				}, 1))
				return
			}, 1)
			gen5786 := Call(__e, ShenFunc(symcons), MakeSymbol("or"), gen5785)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5786)

			gen5787 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V593 := __args[0]
				_ = V593
				__e.TailApply(ShenFunc(symoptimise), V593)

				return
			}, 1)
			gen5788 := Call(__e, ShenFunc(symcons), MakeSymbol("optimise"), gen5787)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5788)

			gen5789 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V594 := __args[0]
				_ = V594
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V595 := __args[0]
					_ = V595
					__e.TailApply(ShenFunc(symopen), V594, V595)

					return
				}, 1))
				return
			}, 1)
			gen5790 := Call(__e, ShenFunc(symcons), MakeSymbol("open"), gen5789)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5790)

			gen5791 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V596 := __args[0]
				_ = V596
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V597 := __args[0]
					_ = V597
					__e.TailApply(ShenFunc(symoccurrences), V596, V597)

					return
				}, 1))
				return
			}, 1)
			gen5792 := Call(__e, ShenFunc(symcons), MakeSymbol("occurrences"), gen5791)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5792)

			gen5793 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V598 := __args[0]
				_ = V598
				__e.TailApply(ShenFunc(symoccurs_1check), V598)

				return
			}, 1)
			gen5794 := Call(__e, ShenFunc(symcons), MakeSymbol("occurs-check"), gen5793)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5794)

			gen5795 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V599 := __args[0]
				_ = V599
				__e.TailApply(ShenFunc(symn_1_6string), V599)

				return
			}, 1)
			gen5796 := Call(__e, ShenFunc(symcons), MakeSymbol("n->string"), gen5795)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5796)

			gen5797 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V600 := __args[0]
				_ = V600
				__e.TailApply(ShenFunc(symnumber_2), V600)

				return
			}, 1)
			gen5798 := Call(__e, ShenFunc(symcons), MakeSymbol("number?"), gen5797)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5798)

			gen5799 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V601 := __args[0]
				_ = V601
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V602 := __args[0]
					_ = V602
					__e.TailApply(ShenFunc(symnth), V601, V602)

					return
				}, 1))
				return
			}, 1)
			gen5800 := Call(__e, ShenFunc(symcons), MakeSymbol("nth"), gen5799)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5800)

			gen5801 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V603 := __args[0]
				_ = V603
				__e.TailApply(ShenFunc(symnot), V603)

				return
			}, 1)
			gen5802 := Call(__e, ShenFunc(symcons), MakeSymbol("not"), gen5801)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5802)

			gen5803 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V604 := __args[0]
				_ = V604
				__e.TailApply(ShenFunc(symnl), V604)

				return
			}, 1)
			gen5804 := Call(__e, ShenFunc(symcons), MakeSymbol("nl"), gen5803)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5804)

			gen5805 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V605 := __args[0]
				_ = V605
				__e.TailApply(ShenFunc(symmacroexpand), V605)

				return
			}, 1)
			gen5806 := Call(__e, ShenFunc(symcons), MakeSymbol("macroexpand"), gen5805)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5806)

			gen5807 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V606 := __args[0]
				_ = V606
				__e.TailApply(ShenFunc(symmaxinferences), V606)

				return
			}, 1)
			gen5808 := Call(__e, ShenFunc(symcons), MakeSymbol("maxinferences"), gen5807)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5808)

			gen5809 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V607 := __args[0]
				_ = V607
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V608 := __args[0]
					_ = V608
					__e.TailApply(ShenFunc(symmapcan), V607, V608)

					return
				}, 1))
				return
			}, 1)
			gen5810 := Call(__e, ShenFunc(symcons), MakeSymbol("mapcan"), gen5809)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5810)

			gen5811 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V609 := __args[0]
				_ = V609
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V610 := __args[0]
					_ = V610
					__e.TailApply(ShenFunc(symmap), V609, V610)

					return
				}, 1))
				return
			}, 1)
			gen5812 := Call(__e, ShenFunc(symcons), MakeSymbol("map"), gen5811)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5812)

			gen5813 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V611 := __args[0]
				_ = V611
				__e.TailApply(ShenFunc(symload), V611)

				return
			}, 1)
			gen5814 := Call(__e, ShenFunc(symcons), MakeSymbol("load"), gen5813)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5814)

			gen5815 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V612 := __args[0]
				_ = V612
				__e.TailApply(ShenFunc(symlineread), V612)

				return
			}, 1)
			gen5816 := Call(__e, ShenFunc(symcons), MakeSymbol("lineread"), gen5815)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5816)

			gen5817 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V613 := __args[0]
				_ = V613
				__e.TailApply(ShenFunc(symlimit), V613)

				return
			}, 1)
			gen5818 := Call(__e, ShenFunc(symcons), MakeSymbol("limit"), gen5817)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5818)

			gen5819 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V614 := __args[0]
				_ = V614
				__e.TailApply(ShenFunc(symlength), V614)

				return
			}, 1)
			gen5820 := Call(__e, ShenFunc(symcons), MakeSymbol("length"), gen5819)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5820)

			gen5821 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V615 := __args[0]
				_ = V615
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V616 := __args[0]
					_ = V616
					__e.TailApply(ShenFunc(symintersection), V615, V616)

					return
				}, 1))
				return
			}, 1)
			gen5822 := Call(__e, ShenFunc(symcons), MakeSymbol("intersection"), gen5821)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5822)

			gen5823 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V617 := __args[0]
				_ = V617
				__e.TailApply(ShenFunc(symintern), V617)

				return
			}, 1)
			gen5824 := Call(__e, ShenFunc(symcons), MakeSymbol("intern"), gen5823)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5824)

			gen5825 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V618 := __args[0]
				_ = V618
				__e.TailApply(ShenFunc(syminteger_2), V618)

				return
			}, 1)
			gen5826 := Call(__e, ShenFunc(symcons), MakeSymbol("integer?"), gen5825)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5826)

			gen5827 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V619 := __args[0]
				_ = V619
				__e.TailApply(ShenFunc(syminput), V619)

				return
			}, 1)
			gen5828 := Call(__e, ShenFunc(symcons), MakeSymbol("input"), gen5827)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5828)

			gen5829 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V620 := __args[0]
				_ = V620
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V621 := __args[0]
					_ = V621
					__e.TailApply(ShenFunc(syminput_7), V620, V621)

					return
				}, 1))
				return
			}, 1)
			gen5830 := Call(__e, ShenFunc(symcons), MakeSymbol("input+"), gen5829)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5830)

			gen5831 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V622 := __args[0]
				_ = V622
				__e.TailApply(ShenFunc(syminclude), V622)

				return
			}, 1)
			gen5832 := Call(__e, ShenFunc(symcons), MakeSymbol("include"), gen5831)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5832)

			gen5833 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V623 := __args[0]
				_ = V623
				__e.TailApply(ShenFunc(syminclude_1all_1but), V623)

				return
			}, 1)
			gen5834 := Call(__e, ShenFunc(symcons), MakeSymbol("include-all-but"), gen5833)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5834)

			gen5835 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V624 := __args[0]
				_ = V624
				__e.TailApply(ShenFunc(syminternal), V624)

				return
			}, 1)
			gen5836 := Call(__e, ShenFunc(symcons), MakeSymbol("internal"), gen5835)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5836)

			gen5837 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V625 := __args[0]
				_ = V625
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V626 := __args[0]
					_ = V626
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V627 := __args[0]
						_ = V627
						if True == V625 {
							__e.Return(V626)
							return
						} else {
							__e.Return(V627)
							return
						}
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5838 := Call(__e, ShenFunc(symcons), MakeSymbol("if"), gen5837)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5838)

			gen5839 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V628 := __args[0]
				_ = V628
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V629 := __args[0]
					_ = V629
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V630 := __args[0]
						_ = V630
						__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
							V631 := __args[0]
							_ = V631
							__e.TailApply(ShenFunc(symidentical), V628, V629, V630, V631)

							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5840 := Call(__e, ShenFunc(symcons), MakeSymbol("identical"), gen5839)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5840)

			gen5841 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V632 := __args[0]
				_ = V632
				__e.TailApply(ShenFunc(symhead), V632)

				return
			}, 1)
			gen5842 := Call(__e, ShenFunc(symcons), MakeSymbol("head"), gen5841)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5842)

			gen5843 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V633 := __args[0]
				_ = V633
				__e.TailApply(ShenFunc(symhd), V633)

				return
			}, 1)
			gen5844 := Call(__e, ShenFunc(symcons), MakeSymbol("hd"), gen5843)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5844)

			gen5845 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V634 := __args[0]
				_ = V634
				__e.TailApply(ShenFunc(symhdv), V634)

				return
			}, 1)
			gen5846 := Call(__e, ShenFunc(symcons), MakeSymbol("hdv"), gen5845)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5846)

			gen5847 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V635 := __args[0]
				_ = V635
				__e.TailApply(ShenFunc(symhdstr), V635)

				return
			}, 1)
			gen5848 := Call(__e, ShenFunc(symcons), MakeSymbol("hdstr"), gen5847)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5848)

			gen5849 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V636 := __args[0]
				_ = V636
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V637 := __args[0]
					_ = V637
					__e.TailApply(ShenFunc(symhash), V636, V637)

					return
				}, 1))
				return
			}, 1)
			gen5850 := Call(__e, ShenFunc(symcons), MakeSymbol("hash"), gen5849)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5850)

			gen5851 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V638 := __args[0]
				_ = V638
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V639 := __args[0]
					_ = V639
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V640 := __args[0]
						_ = V640
						__e.TailApply(ShenFunc(symget), V638, V639, V640)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5852 := Call(__e, ShenFunc(symcons), MakeSymbol("get"), gen5851)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5852)

			gen5853 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V641 := __args[0]
				_ = V641
				__e.TailApply(ShenFunc(symget_1time), V641)

				return
			}, 1)
			gen5854 := Call(__e, ShenFunc(symcons), MakeSymbol("get-time"), gen5853)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5854)

			gen5855 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V642 := __args[0]
				_ = V642
				__e.TailApply(ShenFunc(symgensym), V642)

				return
			}, 1)
			gen5856 := Call(__e, ShenFunc(symcons), MakeSymbol("gensym"), gen5855)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5856)

			gen5857 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V643 := __args[0]
				_ = V643
				__e.TailApply(ShenFunc(symfst), V643)

				return
			}, 1)
			gen5858 := Call(__e, ShenFunc(symcons), MakeSymbol("fst"), gen5857)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5858)

			gen5859 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V644 := __args[0]
				_ = V644
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					__e.Return(V644)
					return
				}, 0))
				return
			}, 1)
			gen5860 := Call(__e, ShenFunc(symcons), MakeSymbol("freeze"), gen5859)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5860)

			gen5861 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V645 := __args[0]
				_ = V645
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V646 := __args[0]
					_ = V646
					__e.TailApply(ShenFunc(symfix), V645, V646)

					return
				}, 1))
				return
			}, 1)
			gen5862 := Call(__e, ShenFunc(symcons), MakeSymbol("fix"), gen5861)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5862)

			gen5863 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V647 := __args[0]
				_ = V647
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V648 := __args[0]
					_ = V648
					__e.TailApply(ShenFunc(symfail_1if), V647, V648)

					return
				}, 1))
				return
			}, 1)
			gen5864 := Call(__e, ShenFunc(symcons), MakeSymbol("fail-if"), gen5863)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5864)

			gen5865 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V649 := __args[0]
				_ = V649
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V650 := __args[0]
					_ = V650
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V651 := __args[0]
						_ = V651
						__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
							V652 := __args[0]
							_ = V652
							__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
								V653 := __args[0]
								_ = V653
								__e.TailApply(ShenFunc(symfindall), V649, V650, V651, V652, V653)

								return
							}, 1))
							return
						}, 1))
						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5866 := Call(__e, ShenFunc(symcons), MakeSymbol("findall"), gen5865)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5866)

			gen5867 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V654 := __args[0]
				_ = V654
				__e.TailApply(ShenFunc(symenable_1type_1theory), V654)

				return
			}, 1)
			gen5868 := Call(__e, ShenFunc(symcons), MakeSymbol("enable-type-theory"), gen5867)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5868)

			gen5869 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V655 := __args[0]
				_ = V655
				__e.TailApply(ShenFunc(symexplode), V655)

				return
			}, 1)
			gen5870 := Call(__e, ShenFunc(symcons), MakeSymbol("explode"), gen5869)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5870)

			gen5871 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V656 := __args[0]
				_ = V656
				__e.TailApply(ShenFunc(symexternal), V656)

				return
			}, 1)
			gen5872 := Call(__e, ShenFunc(symcons), MakeSymbol("external"), gen5871)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5872)

			gen5873 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V657 := __args[0]
				_ = V657
				__e.TailApply(ShenFunc(symeval_1kl), V657)

				return
			}, 1)
			gen5874 := Call(__e, ShenFunc(symcons), MakeSymbol("eval-kl"), gen5873)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5874)

			gen5875 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V658 := __args[0]
				_ = V658
				__e.TailApply(ShenFunc(symeval), V658)

				return
			}, 1)
			gen5876 := Call(__e, ShenFunc(symcons), MakeSymbol("eval"), gen5875)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5876)

			gen5877 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V659 := __args[0]
				_ = V659
				__e.TailApply(ShenFunc(symerror_1to_1string), V659)

				return
			}, 1)
			gen5878 := Call(__e, ShenFunc(symcons), MakeSymbol("error-to-string"), gen5877)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5878)

			gen5879 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V660 := __args[0]
				_ = V660
				__e.TailApply(ShenFunc(symempty_2), V660)

				return
			}, 1)
			gen5880 := Call(__e, ShenFunc(symcons), MakeSymbol("empty?"), gen5879)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5880)

			gen5881 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V661 := __args[0]
				_ = V661
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V662 := __args[0]
					_ = V662
					__e.TailApply(ShenFunc(symelement_2), V661, V662)

					return
				}, 1))
				return
			}, 1)
			gen5882 := Call(__e, ShenFunc(symcons), MakeSymbol("element?"), gen5881)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5882)

			gen5883 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V663 := __args[0]
				_ = V663
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V664 := __args[0]
					_ = V664
					// V663
					__e.Return(V664)
					return

				}, 1))
				return
			}, 1)
			gen5884 := Call(__e, ShenFunc(symcons), MakeSymbol("do"), gen5883)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5884)

			gen5885 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V665 := __args[0]
				_ = V665
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V666 := __args[0]
					_ = V666
					__e.TailApply(ShenFunc(symdifference), V665, V666)

					return
				}, 1))
				return
			}, 1)
			gen5886 := Call(__e, ShenFunc(symcons), MakeSymbol("difference"), gen5885)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5886)

			gen5887 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V667 := __args[0]
				_ = V667
				__e.TailApply(ShenFunc(symdestroy), V667)

				return
			}, 1)
			gen5888 := Call(__e, ShenFunc(symcons), MakeSymbol("destroy"), gen5887)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5888)

			gen5889 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V668 := __args[0]
				_ = V668
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V669 := __args[0]
					_ = V669
					__e.TailApply(ShenFunc(symdeclare), V668, V669)

					return
				}, 1))
				return
			}, 1)
			gen5890 := Call(__e, ShenFunc(symcons), MakeSymbol("declare"), gen5889)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5890)

			gen5891 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V670 := __args[0]
				_ = V670
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V671 := __args[0]
					_ = V671
					__e.TailApply(ShenFunc(symcn), V670, V671)

					return
				}, 1))
				return
			}, 1)
			gen5892 := Call(__e, ShenFunc(symcons), MakeSymbol("cn"), gen5891)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5892)

			gen5893 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V672 := __args[0]
				_ = V672
				__e.TailApply(ShenFunc(symcons_2), V672)

				return
			}, 1)
			gen5894 := Call(__e, ShenFunc(symcons), MakeSymbol("cons?"), gen5893)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5894)

			gen5895 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V673 := __args[0]
				_ = V673
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V674 := __args[0]
					_ = V674
					__e.TailApply(ShenFunc(symcons), V673, V674)

					return
				}, 1))
				return
			}, 1)
			gen5896 := Call(__e, ShenFunc(symcons), MakeSymbol("cons"), gen5895)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5896)

			gen5897 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V675 := __args[0]
				_ = V675
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V676 := __args[0]
					_ = V676
					__e.TailApply(ShenFunc(symconcat), V675, V676)

					return
				}, 1))
				return
			}, 1)
			gen5898 := Call(__e, ShenFunc(symcons), MakeSymbol("concat"), gen5897)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5898)

			gen5899 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V677 := __args[0]
				_ = V677
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V678 := __args[0]
					_ = V678
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V679 := __args[0]
						_ = V679
						__e.TailApply(ShenFunc(symcompile), V677, V678, V679)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5900 := Call(__e, ShenFunc(symcons), MakeSymbol("compile"), gen5899)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5900)

			gen5901 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V680 := __args[0]
				_ = V680
				__e.TailApply(ShenFunc(symcd), V680)

				return
			}, 1)
			gen5902 := Call(__e, ShenFunc(symcons), MakeSymbol("cd"), gen5901)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5902)

			gen5903 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V681 := __args[0]
				_ = V681
				__e.TailApply(ShenFunc(symclose), V681)

				return
			}, 1)
			gen5904 := Call(__e, ShenFunc(symcons), MakeSymbol("close"), gen5903)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5904)

			gen5905 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V682 := __args[0]
				_ = V682
				__e.TailApply(ShenFunc(symbound_2), V682)

				return
			}, 1)
			gen5906 := Call(__e, ShenFunc(symcons), MakeSymbol("bound?"), gen5905)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5906)

			gen5907 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V683 := __args[0]
				_ = V683
				__e.TailApply(ShenFunc(symboolean_2), V683)

				return
			}, 1)
			gen5908 := Call(__e, ShenFunc(symcons), MakeSymbol("boolean?"), gen5907)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5908)

			gen5909 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V684 := __args[0]
				_ = V684
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V685 := __args[0]
					_ = V685
					__e.TailApply(ShenFunc(symassoc), V684, V685)

					return
				}, 1))
				return
			}, 1)
			gen5910 := Call(__e, ShenFunc(symcons), MakeSymbol("assoc"), gen5909)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5910)

			gen5911 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V686 := __args[0]
				_ = V686
				__e.TailApply(ShenFunc(symarity), V686)

				return
			}, 1)
			gen5912 := Call(__e, ShenFunc(symcons), MakeSymbol("arity"), gen5911)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5912)

			gen5913 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V687 := __args[0]
				_ = V687
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V688 := __args[0]
					_ = V688
					__e.TailApply(ShenFunc(symappend), V687, V688)

					return
				}, 1))
				return
			}, 1)
			gen5914 := Call(__e, ShenFunc(symcons), MakeSymbol("append"), gen5913)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5914)

			gen5915 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V689 := __args[0]
				_ = V689
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V690 := __args[0]
					_ = V690
					if True == V689 {
						if True == V690 {
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
				}, 1))
				return
			}, 1)
			gen5916 := Call(__e, ShenFunc(symcons), MakeSymbol("and"), gen5915)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5916)

			gen5917 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V691 := __args[0]
				_ = V691
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V692 := __args[0]
					_ = V692
					__e.TailApply(ShenFunc(symadjoin), V691, V692)

					return
				}, 1))
				return
			}, 1)
			gen5918 := Call(__e, ShenFunc(symcons), MakeSymbol("adjoin"), gen5917)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5918)

			gen5919 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V693 := __args[0]
				_ = V693
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V694 := __args[0]
					_ = V694
					__e.TailApply(ShenFunc(sym_5_1address), V693, V694)

					return
				}, 1))
				return
			}, 1)
			gen5920 := Call(__e, ShenFunc(symcons), MakeSymbol("<-address"), gen5919)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5920)

			gen5921 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V695 := __args[0]
				_ = V695
				__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
					V696 := __args[0]
					_ = V696
					__e.Return(MakeNative(func(__e Evaluator, __args ...Obj) {
						V697 := __args[0]
						_ = V697
						__e.TailApply(ShenFunc(symaddress_1_6), V695, V696, V697)

						return
					}, 1))
					return
				}, 1))
				return
			}, 1)
			gen5922 := Call(__e, ShenFunc(symcons), MakeSymbol("address->"), gen5921)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5922)

			gen5923 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V698 := __args[0]
				_ = V698
				__e.TailApply(ShenFunc(symabsvector_2), V698)

				return
			}, 1)
			gen5924 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector?"), gen5923)

			Call(__e, ShenFunc(symshen_4set_1lambda_1form_1entry), gen5924)

			gen5925 := MakeNative(func(__e Evaluator, __args ...Obj) {
				V699 := __args[0]
				_ = V699
				__e.TailApply(ShenFunc(symabsvector), V699)

				return
			}, 1)
			gen5926 := Call(__e, ShenFunc(symcons), MakeSymbol("absvector"), gen5925)

			__e.TailApply(ShenFunc(symshen_4set_1lambda_1form_1entry), gen5926)

			return

		}, 0)
		Call(__e, ShenFunc(symdefun), MakeSymbol("shen.initialise-lambda-forms"), gen5927)

		gen5928 := MakeNative(func(__e Evaluator, __args ...Obj) {
			Call(__e, ShenFunc(symshen_4initialise_1environment))
			Call(__e, ShenFunc(symshen_4initialise_1lambda_1forms))
			Call(__e, ShenFunc(symshen_4initialise_1signedfunc_1lambda_1forms))
			__e.TailApply(ShenFunc(symshen_4initialise_1signedfuncs))

			return

		}, 0)
		__e.TailApply(ShenFunc(symdefun), MakeSymbol("shen.initialise"), gen5928)

		return

	}, 0))
}
