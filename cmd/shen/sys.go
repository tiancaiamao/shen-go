package main

import . "github.com/tiancaiamao/shen-go/kl"

var SysMain = MakeNative(func(__e *ControlFlow) {
	MakeString("Copyright (c) 2010-2015, Mark Tarver\n\nAll rights reserved.\n\nRedistribution and use in source and binary forms, with or without\nmodification, are permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS ''AS IS'' AND\nANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED\nWARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE\nDISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE\nFOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL\nDAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR\nSERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER\nCAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,\nOR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE\nOF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n")
	gen5604 := MakeNative(func(__e *ControlFlow) {
		V1920 := __e.Get(1)
		_ = V1920
		__e.TailApply(V1920)

		return
	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symthaw, gen5604)

	gen5618 := MakeNative(func(__e *ControlFlow) {
		V1922 := __e.Get(1)
		_ = V1922
		gen5613 := MakeNative(func(__e *ControlFlow) {
			Macroexpand := __e.Get(1)
			_ = Macroexpand
			gen5611 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packaged_2)

			gen5612 := Call(__e, gen5611, Macroexpand)

			if True == gen5612 {
				gen5606 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				gen5608 := MakeNative(func(__e *ControlFlow) {
					Z := __e.Get(1)
					_ = Z
					gen5607 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

					__e.TailApply(gen5607, Z)

					return

				}, 1)

				gen5609 := Call(__e, PrimNS1Value(symns2_1value), symshen_4package_1contents)

				gen5610 := Call(__e, gen5609, Macroexpand)

				__e.TailApply(gen5606, gen5608, gen5610)

				return

			} else {
				gen5605 := Call(__e, PrimNS1Value(symns2_1value), symshen_4eval_1without_1macros)

				__e.TailApply(gen5605, Macroexpand)

				return

			}

		}, 1)

		gen5614 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

		gen5616 := MakeNative(func(__e *ControlFlow) {
			Y := __e.Get(1)
			_ = Y
			gen5615 := Call(__e, PrimNS1Value(symns2_1value), symmacroexpand)

			__e.TailApply(gen5615, Y)

			return

		}, 1)

		gen5617 := Call(__e, gen5614, gen5616, V1922)

		__e.TailApply(gen5613, gen5617)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symeval, gen5618)

	gen5624 := MakeNative(func(__e *ControlFlow) {
		V1924 := __e.Get(1)
		_ = V1924
		gen5619 := Call(__e, PrimNS1Value(symns2_1value), symeval_1kl)

		gen5620 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

		gen5621 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1input_7)

		gen5622 := Call(__e, gen5621, V1924)

		gen5623 := Call(__e, gen5620, gen5622)

		__e.TailApply(gen5619, gen5623)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4eval_1without_1macros, gen5624)

	gen5713 := MakeNative(func(__e *ControlFlow) {
		V1926 := __e.Get(1)
		_ = V1926
		gen5710 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen5711 := Call(__e, gen5710, V1926)

		var gen5712 Obj

		if True == gen5711 {
			gen5705 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen5706 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen5707 := Call(__e, gen5706, V1926)

			gen5708 := Call(__e, gen5705, syminput_7, gen5707)

			var gen5709 Obj

			if True == gen5708 {
				gen5700 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen5701 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5702 := Call(__e, gen5701, V1926)

				gen5703 := Call(__e, gen5700, gen5702)

				var gen5704 Obj

				if True == gen5703 {
					gen5693 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen5694 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5695 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5696 := Call(__e, gen5695, V1926)

					gen5697 := Call(__e, gen5694, gen5696)

					gen5698 := Call(__e, gen5693, gen5697)

					var gen5699 Obj

					if True == gen5698 {
						gen5685 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen5686 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5687 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5688 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5689 := Call(__e, gen5688, V1926)

						gen5690 := Call(__e, gen5687, gen5689)

						gen5691 := Call(__e, gen5686, gen5690)

						gen5692 := Call(__e, gen5685, Nil, gen5691)

						if True == gen5692 {
							gen5699 = True
						} else {
							gen5699 = False
						}

					} else {
						gen5699 = False
					}

					if True == gen5699 {
						gen5704 = True
					} else {
						gen5704 = False
					}

				} else {
					gen5704 = False
				}

				if True == gen5704 {
					gen5709 = True
				} else {
					gen5709 = False
				}

			} else {
				gen5709 = False
			}

			if True == gen5709 {
				gen5712 = True
			} else {
				gen5712 = False
			}

		} else {
			gen5712 = False
		}

		if True == gen5712 {
			gen5672 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen5673 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen5674 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

			gen5675 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen5676 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5677 := Call(__e, gen5676, V1926)

			gen5678 := Call(__e, gen5675, gen5677)

			gen5679 := Call(__e, gen5674, gen5678)

			gen5680 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5681 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5682 := Call(__e, gen5681, V1926)

			gen5683 := Call(__e, gen5680, gen5682)

			gen5684 := Call(__e, gen5673, gen5679, gen5683)

			__e.TailApply(gen5672, syminput_7, gen5684)

			return

		} else {
			gen5669 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen5670 := Call(__e, gen5669, V1926)

			var gen5671 Obj

			if True == gen5670 {
				gen5664 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen5665 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen5666 := Call(__e, gen5665, V1926)

				gen5667 := Call(__e, gen5664, symshen_4read_7, gen5666)

				var gen5668 Obj

				if True == gen5667 {
					gen5659 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen5660 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5661 := Call(__e, gen5660, V1926)

					gen5662 := Call(__e, gen5659, gen5661)

					var gen5663 Obj

					if True == gen5662 {
						gen5652 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen5653 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5654 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5655 := Call(__e, gen5654, V1926)

						gen5656 := Call(__e, gen5653, gen5655)

						gen5657 := Call(__e, gen5652, gen5656)

						var gen5658 Obj

						if True == gen5657 {
							gen5644 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

							gen5645 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen5646 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen5647 := Call(__e, PrimNS1Value(symns2_1value), symtl)

							gen5648 := Call(__e, gen5647, V1926)

							gen5649 := Call(__e, gen5646, gen5648)

							gen5650 := Call(__e, gen5645, gen5649)

							gen5651 := Call(__e, gen5644, Nil, gen5650)

							if True == gen5651 {
								gen5658 = True
							} else {
								gen5658 = False
							}

						} else {
							gen5658 = False
						}

						if True == gen5658 {
							gen5663 = True
						} else {
							gen5663 = False
						}

					} else {
						gen5663 = False
					}

					if True == gen5663 {
						gen5668 = True
					} else {
						gen5668 = False
					}

				} else {
					gen5668 = False
				}

				if True == gen5668 {
					gen5671 = True
				} else {
					gen5671 = False
				}

			} else {
				gen5671 = False
			}

			if True == gen5671 {
				gen5631 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen5632 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen5633 := Call(__e, PrimNS1Value(symns2_1value), symshen_4rcons__form)

				gen5634 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen5635 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5636 := Call(__e, gen5635, V1926)

				gen5637 := Call(__e, gen5634, gen5636)

				gen5638 := Call(__e, gen5633, gen5637)

				gen5639 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5640 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5641 := Call(__e, gen5640, V1926)

				gen5642 := Call(__e, gen5639, gen5641)

				gen5643 := Call(__e, gen5632, gen5638, gen5642)

				__e.TailApply(gen5631, symshen_4read_7, gen5643)

				return

			} else {
				gen5629 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen5630 := Call(__e, gen5629, V1926)

				if True == gen5630 {
					gen5626 := Call(__e, PrimNS1Value(symns2_1value), symmap)

					gen5628 := MakeNative(func(__e *ControlFlow) {
						Z := __e.Get(1)
						_ = Z
						gen5627 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1input_7)

						__e.TailApply(gen5627, Z)

						return

					}, 1)

					__e.TailApply(gen5626, gen5628, V1926)

					return

				} else {
					if True == True {
						__e.Return(V1926)
						return
					} else {
						gen5625 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen5625, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4proc_1input_7, gen5713)

	gen5799 := MakeNative(func(__e *ControlFlow) {
		V1928 := __e.Get(1)
		_ = V1928
		gen5796 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen5797 := Call(__e, gen5796, V1928)

		var gen5798 Obj

		if True == gen5797 {
			gen5791 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen5792 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen5793 := Call(__e, gen5792, V1928)

			gen5794 := Call(__e, gen5791, symdefine, gen5793)

			var gen5795 Obj

			if True == gen5794 {
				gen5787 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen5788 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5789 := Call(__e, gen5788, V1928)

				gen5790 := Call(__e, gen5787, gen5789)

				if True == gen5790 {
					gen5795 = True
				} else {
					gen5795 = False
				}

			} else {
				gen5795 = False
			}

			if True == gen5795 {
				gen5798 = True
			} else {
				gen5798 = False
			}

		} else {
			gen5798 = False
		}

		if True == gen5798 {
			gen5778 := Call(__e, PrimNS1Value(symns2_1value), symshen_4shen_1_6kl)

			gen5779 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen5780 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5781 := Call(__e, gen5780, V1928)

			gen5782 := Call(__e, gen5779, gen5781)

			gen5783 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5784 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5785 := Call(__e, gen5784, V1928)

			gen5786 := Call(__e, gen5783, gen5785)

			__e.TailApply(gen5778, gen5782, gen5786)

			return

		} else {
			gen5775 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen5776 := Call(__e, gen5775, V1928)

			var gen5777 Obj

			if True == gen5776 {
				gen5770 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen5771 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen5772 := Call(__e, gen5771, V1928)

				gen5773 := Call(__e, gen5770, symdefmacro, gen5772)

				var gen5774 Obj

				if True == gen5773 {
					gen5766 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen5767 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5768 := Call(__e, gen5767, V1928)

					gen5769 := Call(__e, gen5766, gen5768)

					if True == gen5769 {
						gen5774 = True
					} else {
						gen5774 = False
					}

				} else {
					gen5774 = False
				}

				if True == gen5774 {
					gen5777 = True
				} else {
					gen5777 = False
				}

			} else {
				gen5777 = False
			}

			if True == gen5777 {
				gen5759 := MakeNative(func(__e *ControlFlow) {
					Default := __e.Get(1)
					_ = Default
					gen5742 := MakeNative(func(__e *ControlFlow) {
						Def := __e.Get(1)
						_ = Def
						gen5735 := MakeNative(func(__e *ControlFlow) {
							MacroAdd := __e.Get(1)
							_ = MacroAdd
							__e.Return(Def)
							return
						}, 1)

						gen5736 := Call(__e, PrimNS1Value(symns2_1value), symshen_4add_1macro)

						gen5737 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen5738 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5739 := Call(__e, gen5738, V1928)

						gen5740 := Call(__e, gen5737, gen5739)

						gen5741 := Call(__e, gen5736, gen5740)

						__e.TailApply(gen5735, gen5741)

						return

					}, 1)

					gen5743 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

					gen5744 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen5745 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen5746 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen5747 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5748 := Call(__e, gen5747, V1928)

					gen5749 := Call(__e, gen5746, gen5748)

					gen5750 := Call(__e, PrimNS1Value(symns2_1value), symappend)

					gen5751 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5752 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5753 := Call(__e, gen5752, V1928)

					gen5754 := Call(__e, gen5751, gen5753)

					gen5755 := Call(__e, gen5750, gen5754, Default)

					gen5756 := Call(__e, gen5745, gen5749, gen5755)

					gen5757 := Call(__e, gen5744, symdefine, gen5756)

					gen5758 := Call(__e, gen5743, gen5757)

					__e.TailApply(gen5742, gen5758)

					return

				}, 1)

				gen5760 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen5761 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen5762 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen5763 := Call(__e, gen5762, symX, Nil)

				gen5764 := Call(__e, gen5761, sym_1_6, gen5763)

				gen5765 := Call(__e, gen5760, symX, gen5764)

				__e.TailApply(gen5759, gen5765)

				return

			} else {
				gen5732 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen5733 := Call(__e, gen5732, V1928)

				var gen5734 Obj

				if True == gen5733 {
					gen5727 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen5728 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen5729 := Call(__e, gen5728, V1928)

					gen5730 := Call(__e, gen5727, symdefcc, gen5729)

					var gen5731 Obj

					if True == gen5730 {
						gen5723 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen5724 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5725 := Call(__e, gen5724, V1928)

						gen5726 := Call(__e, gen5723, gen5725)

						if True == gen5726 {
							gen5731 = True
						} else {
							gen5731 = False
						}

					} else {
						gen5731 = False
					}

					if True == gen5731 {
						gen5734 = True
					} else {
						gen5734 = False
					}

				} else {
					gen5734 = False
				}

				if True == gen5734 {
					gen5720 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

					gen5721 := Call(__e, PrimNS1Value(symns2_1value), symshen_4yacc)

					gen5722 := Call(__e, gen5721, V1928)

					__e.TailApply(gen5720, gen5722)

					return

				} else {
					gen5718 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen5719 := Call(__e, gen5718, V1928)

					if True == gen5719 {
						gen5715 := Call(__e, PrimNS1Value(symns2_1value), symmap)

						gen5717 := MakeNative(func(__e *ControlFlow) {
							Z := __e.Get(1)
							_ = Z
							gen5716 := Call(__e, PrimNS1Value(symns2_1value), symshen_4elim_1def)

							__e.TailApply(gen5716, Z)

							return

						}, 1)

						__e.TailApply(gen5715, gen5717, V1928)

						return

					} else {
						if True == True {
							__e.Return(V1928)
							return
						} else {
							gen5714 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen5714, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4elim_1def, gen5799)

	gen5819 := MakeNative(func(__e *ControlFlow) {
		V1930 := __e.Get(1)
		_ = V1930
		gen5816 := MakeNative(func(__e *ControlFlow) {
			MacroReg := __e.Get(1)
			_ = MacroReg
			gen5809 := MakeNative(func(__e *ControlFlow) {
				NewMacroReg := __e.Get(1)
				_ = NewMacroReg
				gen5807 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen5808 := Call(__e, gen5807, MacroReg, NewMacroReg)

				if True == gen5808 {
					__e.Return(symshen_4skip)
					return
				} else {
					gen5800 := Call(__e, PrimNS1Value(symns2_1value), symset)

					gen5801 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen5802 := Call(__e, PrimNS1Value(symns2_1value), symfunction)

					gen5803 := Call(__e, gen5802, V1930)

					gen5804 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

					gen5805 := Call(__e, gen5804, sym_dmacros_d)

					gen5806 := Call(__e, gen5801, gen5803, gen5805)

					__e.TailApply(gen5800, sym_dmacros_d, gen5806)

					return

				}

			}, 1)

			gen5810 := Call(__e, PrimNS1Value(symns2_1value), symset)

			gen5811 := Call(__e, PrimNS1Value(symns2_1value), symadjoin)

			gen5812 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen5813 := Call(__e, gen5812, symshen_4_dmacroreg_d)

			gen5814 := Call(__e, gen5811, V1930, gen5813)

			gen5815 := Call(__e, gen5810, symshen_4_dmacroreg_d, gen5814)

			__e.TailApply(gen5809, gen5815)

			return

		}, 1)

		gen5817 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen5818 := Call(__e, gen5817, symshen_4_dmacroreg_d)

		__e.TailApply(gen5816, gen5818)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4add_1macro, gen5819)

	gen5840 := MakeNative(func(__e *ControlFlow) {
		V1938 := __e.Get(1)
		_ = V1938
		gen5837 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen5838 := Call(__e, gen5837, V1938)

		var gen5839 Obj

		if True == gen5838 {
			gen5832 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen5833 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen5834 := Call(__e, gen5833, V1938)

			gen5835 := Call(__e, gen5832, sympackage, gen5834)

			var gen5836 Obj

			if True == gen5835 {
				gen5827 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen5828 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5829 := Call(__e, gen5828, V1938)

				gen5830 := Call(__e, gen5827, gen5829)

				var gen5831 Obj

				if True == gen5830 {
					gen5821 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen5822 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5823 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5824 := Call(__e, gen5823, V1938)

					gen5825 := Call(__e, gen5822, gen5824)

					gen5826 := Call(__e, gen5821, gen5825)

					if True == gen5826 {
						gen5831 = True
					} else {
						gen5831 = False
					}

				} else {
					gen5831 = False
				}

				if True == gen5831 {
					gen5836 = True
				} else {
					gen5836 = False
				}

			} else {
				gen5836 = False
			}

			if True == gen5836 {
				gen5839 = True
			} else {
				gen5839 = False
			}

		} else {
			gen5839 = False
		}

		if True == gen5839 {
			__e.Return(True)
			return
		} else {
			if True == True {
				__e.Return(False)
				return
			} else {
				gen5820 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen5820, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4packaged_2, gen5840)

	gen5851 := MakeNative(func(__e *ControlFlow) {
		V1940 := __e.Get(1)
		_ = V1940
		gen5844 := MakeNative(func(__e *ControlFlow) {
			gen5841 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen5842 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen5843 := Call(__e, gen5842, sym_dproperty_1vector_d)

			__e.TailApply(gen5841, V1940, symshen_4external_1symbols, gen5843)

			return

		}, 0)

		gen5850 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen5845 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen5846 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen5847 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen5848 := Call(__e, gen5847, V1940, MakeString(" has not been used.\n"), symshen_4a)

			gen5849 := Call(__e, gen5846, MakeString("package "), gen5848)

			__e.TailApply(gen5845, gen5849)

			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen5844, gen5850)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symexternal, gen5851)

	gen5862 := MakeNative(func(__e *ControlFlow) {
		V1942 := __e.Get(1)
		_ = V1942
		gen5855 := MakeNative(func(__e *ControlFlow) {
			gen5852 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen5853 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen5854 := Call(__e, gen5853, sym_dproperty_1vector_d)

			__e.TailApply(gen5852, V1942, symshen_4internal_1symbols, gen5854)

			return

		}, 0)

		gen5861 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen5856 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen5857 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen5858 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen5859 := Call(__e, gen5858, V1942, MakeString(" has not been used.\n"), symshen_4a)

			gen5860 := Call(__e, gen5857, MakeString("package "), gen5859)

			__e.TailApply(gen5856, gen5860)

			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen5855, gen5861)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), syminternal, gen5862)

	gen5946 := MakeNative(func(__e *ControlFlow) {
		V1946 := __e.Get(1)
		_ = V1946
		gen5943 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen5944 := Call(__e, gen5943, V1946)

		var gen5945 Obj

		if True == gen5944 {
			gen5938 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen5939 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen5940 := Call(__e, gen5939, V1946)

			gen5941 := Call(__e, gen5938, sympackage, gen5940)

			var gen5942 Obj

			if True == gen5941 {
				gen5933 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen5934 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5935 := Call(__e, gen5934, V1946)

				gen5936 := Call(__e, gen5933, gen5935)

				var gen5937 Obj

				if True == gen5936 {
					gen5926 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen5927 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen5928 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5929 := Call(__e, gen5928, V1946)

					gen5930 := Call(__e, gen5927, gen5929)

					gen5931 := Call(__e, gen5926, symnull, gen5930)

					var gen5932 Obj

					if True == gen5931 {
						gen5920 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen5921 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5922 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5923 := Call(__e, gen5922, V1946)

						gen5924 := Call(__e, gen5921, gen5923)

						gen5925 := Call(__e, gen5920, gen5924)

						if True == gen5925 {
							gen5932 = True
						} else {
							gen5932 = False
						}

					} else {
						gen5932 = False
					}

					if True == gen5932 {
						gen5937 = True
					} else {
						gen5937 = False
					}

				} else {
					gen5937 = False
				}

				if True == gen5937 {
					gen5942 = True
				} else {
					gen5942 = False
				}

			} else {
				gen5942 = False
			}

			if True == gen5942 {
				gen5945 = True
			} else {
				gen5945 = False
			}

		} else {
			gen5945 = False
		}

		if True == gen5945 {
			gen5915 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5916 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5917 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			gen5918 := Call(__e, gen5917, V1946)

			gen5919 := Call(__e, gen5916, gen5918)

			__e.TailApply(gen5915, gen5919)

			return

		} else {
			gen5912 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen5913 := Call(__e, gen5912, V1946)

			var gen5914 Obj

			if True == gen5913 {
				gen5907 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen5908 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen5909 := Call(__e, gen5908, V1946)

				gen5910 := Call(__e, gen5907, sympackage, gen5909)

				var gen5911 Obj

				if True == gen5910 {
					gen5902 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen5903 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen5904 := Call(__e, gen5903, V1946)

					gen5905 := Call(__e, gen5902, gen5904)

					var gen5906 Obj

					if True == gen5905 {
						gen5896 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

						gen5897 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5898 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5899 := Call(__e, gen5898, V1946)

						gen5900 := Call(__e, gen5897, gen5899)

						gen5901 := Call(__e, gen5896, gen5900)

						if True == gen5901 {
							gen5906 = True
						} else {
							gen5906 = False
						}

					} else {
						gen5906 = False
					}

					if True == gen5906 {
						gen5911 = True
					} else {
						gen5911 = False
					}

				} else {
					gen5911 = False
				}

				if True == gen5911 {
					gen5914 = True
				} else {
					gen5914 = False
				}

			} else {
				gen5914 = False
			}

			if True == gen5914 {
				gen5885 := MakeNative(func(__e *ControlFlow) {
					PackageNameDot := __e.Get(1)
					_ = PackageNameDot
					gen5882 := MakeNative(func(__e *ControlFlow) {
						ExpPackageNameDot := __e.Get(1)
						_ = ExpPackageNameDot
						gen5865 := Call(__e, PrimNS1Value(symns2_1value), symshen_4packageh)

						gen5866 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen5867 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5868 := Call(__e, gen5867, V1946)

						gen5869 := Call(__e, gen5866, gen5868)

						gen5870 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen5871 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5872 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5873 := Call(__e, gen5872, V1946)

						gen5874 := Call(__e, gen5871, gen5873)

						gen5875 := Call(__e, gen5870, gen5874)

						gen5876 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5877 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5878 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen5879 := Call(__e, gen5878, V1946)

						gen5880 := Call(__e, gen5877, gen5879)

						gen5881 := Call(__e, gen5876, gen5880)

						__e.TailApply(gen5865, gen5869, gen5875, gen5881, ExpPackageNameDot)

						return

					}, 1)

					gen5883 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

					gen5884 := Call(__e, gen5883, PackageNameDot)

					__e.TailApply(gen5882, gen5884)

					return

				}, 1)

				gen5886 := Call(__e, PrimNS1Value(symns2_1value), symintern)

				gen5887 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen5888 := Call(__e, PrimNS1Value(symns2_1value), symstr)

				gen5889 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen5890 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen5891 := Call(__e, gen5890, V1946)

				gen5892 := Call(__e, gen5889, gen5891)

				gen5893 := Call(__e, gen5888, gen5892)

				gen5894 := Call(__e, gen5887, gen5893, MakeString("."))

				gen5895 := Call(__e, gen5886, gen5894)

				__e.TailApply(gen5885, gen5895)

				return

			} else {
				if True == True {
					gen5864 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen5864, symshen_4package_1contents)

					return

				} else {
					gen5863 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen5863, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4package_1contents, gen5946)

	gen5954 := MakeNative(func(__e *ControlFlow) {
		V1949 := __e.Get(1)
		_ = V1949
		V1950 := __e.Get(2)
		_ = V1950
		gen5952 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen5953 := Call(__e, gen5952, V1950)

		if True == gen5953 {
			gen5948 := Call(__e, PrimNS1Value(symns2_1value), symmap)

			gen5950 := MakeNative(func(__e *ControlFlow) {
				Z := __e.Get(1)
				_ = Z
				gen5949 := Call(__e, PrimNS1Value(symns2_1value), symshen_4walk)

				__e.TailApply(gen5949, V1949, Z)

				return

			}, 1)

			gen5951 := Call(__e, gen5948, gen5950, V1950)

			__e.TailApply(V1949, gen5951)

			return

		} else {
			if True == True {
				__e.TailApply(V1949, V1950)

				return
			} else {
				gen5947 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen5947, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4walk, gen5954)

	gen5973 := MakeNative(func(__e *ControlFlow) {
		V1954 := __e.Get(1)
		_ = V1954
		V1955 := __e.Get(2)
		_ = V1955
		V1956 := __e.Get(3)
		_ = V1956
		gen5967 := MakeNative(func(__e *ControlFlow) {
			O := __e.Get(1)
			_ = O
			gen5962 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen5963 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			gen5964 := Call(__e, gen5963)

			gen5965 := Call(__e, gen5962, gen5964, O)

			var gen5966 Obj

			if True == gen5965 {
				gen5966 = True
			} else {
				gen5956 := Call(__e, PrimNS1Value(symns2_1value), symnot)

				gen5957 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

				gen5958 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen5959 := Call(__e, gen5958, O)

				gen5960 := Call(__e, gen5957, gen5959)

				gen5961 := Call(__e, gen5956, gen5960)

				if True == gen5961 {
					gen5966 = True
				} else {
					gen5966 = False
				}

			}

			if True == gen5966 {
				__e.TailApply(V1956, O)

				return
			} else {
				gen5955 := Call(__e, PrimNS1Value(symns2_1value), symshen_4hdtl)

				__e.TailApply(gen5955, O)

				return

			}

		}, 1)

		gen5968 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen5969 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen5970 := Call(__e, gen5969, Nil, Nil)

		gen5971 := Call(__e, gen5968, V1955, gen5970)

		gen5972 := Call(__e, V1954, gen5971)

		__e.TailApply(gen5967, gen5972)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symcompile, gen5973)

	gen5976 := MakeNative(func(__e *ControlFlow) {
		V1959 := __e.Get(1)
		_ = V1959
		V1960 := __e.Get(2)
		_ = V1960
		gen5975 := Call(__e, V1959, V1960)

		if True == gen5975 {
			gen5974 := Call(__e, PrimNS1Value(symns2_1value), symfail)

			__e.TailApply(gen5974)

			return

		} else {
			__e.Return(V1960)
			return
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symfail_1if, gen5976)

	gen5978 := MakeNative(func(__e *ControlFlow) {
		V1963 := __e.Get(1)
		_ = V1963
		V1964 := __e.Get(2)
		_ = V1964
		gen5977 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		__e.TailApply(gen5977, V1963, V1964)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), sym_8s, gen5978)

	gen5980 := MakeNative(func(__e *ControlFlow) {
		gen5979 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen5979, symshen_4_dtc_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symtc_2, gen5980)

	gen5989 := MakeNative(func(__e *ControlFlow) {
		V1966 := __e.Get(1)
		_ = V1966
		gen5984 := MakeNative(func(__e *ControlFlow) {
			gen5981 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen5982 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen5983 := Call(__e, gen5982, sym_dproperty_1vector_d)

			__e.TailApply(gen5981, V1966, symshen_4source, gen5983)

			return

		}, 0)

		gen5988 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen5985 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen5986 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen5987 := Call(__e, gen5986, V1966, MakeString(" not found.\n"), symshen_4a)

			__e.TailApply(gen5985, gen5987)

			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen5984, gen5988)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symps, gen5989)

	gen5991 := MakeNative(func(__e *ControlFlow) {
		gen5990 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen5990, sym_dstinput_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symstinput, gen5991)

	gen6007 := MakeNative(func(__e *ControlFlow) {
		V1968 := __e.Get(1)
		_ = V1968
		gen6002 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			gen5999 := MakeNative(func(__e *ControlFlow) {
				ZeroStamp := __e.Get(1)
				_ = ZeroStamp
				gen5992 := MakeNative(func(__e *ControlFlow) {
					Standard := __e.Get(1)
					_ = Standard
					__e.Return(Standard)
					return
				}, 1)

				gen5996 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen5997 := Call(__e, gen5996, V1968, MakeNumber(0))

				var gen5998 Obj

				if True == gen5997 {
					gen5998 = ZeroStamp
				} else {
					gen5993 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fillvector)

					gen5994 := Call(__e, PrimNS1Value(symns2_1value), symfail)

					gen5995 := Call(__e, gen5994)

					gen5998 = Call(__e, gen5993, ZeroStamp, MakeNumber(1), V1968, gen5995)

				}

				__e.TailApply(gen5992, gen5998)

				return

			}, 1)

			gen6000 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			gen6001 := Call(__e, gen6000, Vector, MakeNumber(0), V1968)

			__e.TailApply(gen5999, gen6001)

			return

		}, 1)

		gen6003 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

		gen6004 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		gen6005 := Call(__e, gen6004, V1968, MakeNumber(1))

		gen6006 := Call(__e, gen6003, gen6005)

		__e.TailApply(gen6002, gen6006)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symvector, gen6007)

	gen6017 := MakeNative(func(__e *ControlFlow) {
		V1974 := __e.Get(1)
		_ = V1974
		V1975 := __e.Get(2)
		_ = V1975
		V1976 := __e.Get(3)
		_ = V1976
		V1977 := __e.Get(4)
		_ = V1977
		gen6015 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6016 := Call(__e, gen6015, V1976, V1975)

		if True == gen6016 {
			gen6014 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			__e.TailApply(gen6014, V1974, V1976, V1977)

			return

		} else {
			if True == True {
				gen6009 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fillvector)

				gen6010 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

				gen6011 := Call(__e, gen6010, V1974, V1975, V1977)

				gen6012 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen6013 := Call(__e, gen6012, MakeNumber(1), V1975)

				__e.TailApply(gen6009, gen6011, gen6013, V1976, V1977)

				return

			} else {
				gen6008 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6008, MakeString("no cond match"))

				return

			}
		}

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4fillvector, gen6017)

	gen6030 := MakeNative(func(__e *ControlFlow) {
		V1979 := __e.Get(1)
		_ = V1979
		gen6028 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		gen6029 := Call(__e, gen6028, V1979)

		if True == gen6029 {
			gen6022 := MakeNative(func(__e *ControlFlow) {
				X := __e.Get(1)
				_ = X
				gen6020 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

				gen6021 := Call(__e, gen6020, X)

				if True == gen6021 {
					gen6018 := Call(__e, PrimNS1Value(symns2_1value), sym_6_a)

					gen6019 := Call(__e, gen6018, X, MakeNumber(0))

					if True == gen6019 {
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

			gen6024 := MakeNative(func(__e *ControlFlow) {
				gen6023 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(gen6023, V1979, MakeNumber(0))

				return

			}, 0)

			gen6025 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(MakeNumber(-1))
				return
			}, 1)

			gen6026 := Call(__e, PrimNS1Value(symtry_1catch), gen6024, gen6025)

			gen6027 := Call(__e, gen6022, gen6026)

			if True == gen6027 {
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

	Call(__e, PrimNS1Value(symns2_1set), symvector_2, gen6030)

	gen6035 := MakeNative(func(__e *ControlFlow) {
		V1983 := __e.Get(1)
		_ = V1983
		V1984 := __e.Get(2)
		_ = V1984
		V1985 := __e.Get(3)
		_ = V1985
		gen6033 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6034 := Call(__e, gen6033, V1984, MakeNumber(0))

		if True == gen6034 {
			gen6032 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(gen6032, MakeString("cannot access 0th element of a vector\n"))

			return

		} else {
			gen6031 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			__e.TailApply(gen6031, V1983, V1984, V1985)

			return

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symvector_1_6, gen6035)

	gen6047 := MakeNative(func(__e *ControlFlow) {
		V1988 := __e.Get(1)
		_ = V1988
		V1989 := __e.Get(2)
		_ = V1989
		gen6045 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6046 := Call(__e, gen6045, V1989, MakeNumber(0))

		if True == gen6046 {
			gen6044 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			__e.TailApply(gen6044, MakeString("cannot access 0th element of a vector\n"))

			return

		} else {
			gen6041 := MakeNative(func(__e *ControlFlow) {
				VectorElement := __e.Get(1)
				_ = VectorElement
				gen6037 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen6038 := Call(__e, PrimNS1Value(symns2_1value), symfail)

				gen6039 := Call(__e, gen6038)

				gen6040 := Call(__e, gen6037, VectorElement, gen6039)

				if True == gen6040 {
					gen6036 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6036, MakeString("vector element not found\n"))

					return

				} else {
					__e.Return(VectorElement)
					return
				}

			}, 1)

			gen6042 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

			gen6043 := Call(__e, gen6042, V1988, V1989)

			__e.TailApply(gen6041, gen6043)

			return

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), sym_5_1vector, gen6047)

	gen6052 := MakeNative(func(__e *ControlFlow) {
		V1991 := __e.Get(1)
		_ = V1991
		gen6050 := Call(__e, PrimNS1Value(symns2_1value), syminteger_2)

		gen6051 := Call(__e, gen6050, V1991)

		if True == gen6051 {
			gen6048 := Call(__e, PrimNS1Value(symns2_1value), sym_6_a)

			gen6049 := Call(__e, gen6048, V1991, MakeNumber(0))

			if True == gen6049 {
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

	Call(__e, PrimNS1Value(symns2_1set), symshen_4posint_2, gen6052)

	gen6054 := MakeNative(func(__e *ControlFlow) {
		V1993 := __e.Get(1)
		_ = V1993
		gen6053 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(gen6053, V1993, MakeNumber(0))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symlimit, gen6054)

	gen6070 := MakeNative(func(__e *ControlFlow) {
		V1995 := __e.Get(1)
		_ = V1995
		gen6067 := Call(__e, PrimNS1Value(symns2_1value), symboolean_2)

		gen6068 := Call(__e, gen6067, V1995)

		var gen6069 Obj

		if True == gen6068 {
			gen6069 = True
		} else {
			gen6064 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

			gen6065 := Call(__e, gen6064, V1995)

			var gen6066 Obj

			if True == gen6065 {
				gen6066 = True
			} else {
				gen6062 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

				gen6063 := Call(__e, gen6062, V1995)

				if True == gen6063 {
					gen6066 = True
				} else {
					gen6066 = False
				}

			}

			if True == gen6066 {
				gen6069 = True
			} else {
				gen6069 = False
			}

		}

		if True == gen6069 {
			__e.Return(False)
			return
		} else {
			if True == True {
				gen6060 := MakeNative(func(__e *ControlFlow) {
					gen6057 := MakeNative(func(__e *ControlFlow) {
						String := __e.Get(1)
						_ = String
						gen6056 := Call(__e, PrimNS1Value(symns2_1value), symshen_4analyse_1symbol_2)

						__e.TailApply(gen6056, String)

						return

					}, 1)

					gen6058 := Call(__e, PrimNS1Value(symns2_1value), symstr)

					gen6059 := Call(__e, gen6058, V1995)

					__e.TailApply(gen6057, gen6059)

					return

				}, 0)

				gen6061 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(False)
					return
				}, 1)

				__e.TailApply(PrimNS1Value(symtry_1catch), gen6060, gen6061)

				return

			} else {
				gen6055 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6055, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symsymbol_2, gen6070)

	gen6085 := MakeNative(func(__e *ControlFlow) {
		V1997 := __e.Get(1)
		_ = V1997
		gen6083 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6084 := Call(__e, gen6083, MakeString(""), V1997)

		if True == gen6084 {
			__e.Return(False)
			return
		} else {
			gen6081 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			gen6082 := Call(__e, gen6081, V1997)

			if True == gen6082 {
				gen6077 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alpha_2)

				gen6078 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen6079 := Call(__e, gen6078, V1997, MakeNumber(0))

				gen6080 := Call(__e, gen6077, gen6079)

				if True == gen6080 {
					gen6073 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanums_2)

					gen6074 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen6075 := Call(__e, gen6074, V1997)

					gen6076 := Call(__e, gen6073, gen6075)

					if True == gen6076 {
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
				if True == True {
					gen6072 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6072, symshen_4analyse_1symbol_2)

					return

				} else {
					gen6071 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6071, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1symbol_2, gen6085)

	gen6237 := MakeNative(func(__e *ControlFlow) {
		V1999 := __e.Get(1)
		_ = V1999
		gen6086 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen6087 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6088 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6089 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6090 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6091 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6092 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6093 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6094 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6095 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6096 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6097 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6098 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6099 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6100 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6101 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6102 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6103 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6104 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6105 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6106 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6107 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6108 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6109 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6110 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6111 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6112 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6113 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6114 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6115 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6116 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6117 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6118 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6119 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6120 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6121 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6122 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6123 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6124 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6125 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6126 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6127 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6128 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6129 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6130 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6131 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6132 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6133 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6134 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6135 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6136 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6137 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6138 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6139 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6140 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6141 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6142 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6143 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6144 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6145 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6146 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6147 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6148 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6149 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6150 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6151 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6152 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6153 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6154 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6155 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6156 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6157 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6158 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6159 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6160 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6161 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6162 := Call(__e, gen6161, MakeString("."), Nil)

		gen6163 := Call(__e, gen6160, MakeString("'"), gen6162)

		gen6164 := Call(__e, gen6159, MakeString("#"), gen6163)

		gen6165 := Call(__e, gen6158, MakeString("`"), gen6164)

		gen6166 := Call(__e, gen6157, MakeString(";"), gen6165)

		gen6167 := Call(__e, gen6156, MakeString(":"), gen6166)

		gen6168 := Call(__e, gen6155, MakeString("}"), gen6167)

		gen6169 := Call(__e, gen6154, MakeString("{"), gen6168)

		gen6170 := Call(__e, gen6153, MakeString("%"), gen6169)

		gen6171 := Call(__e, gen6152, MakeString("&"), gen6170)

		gen6172 := Call(__e, gen6151, MakeString("<"), gen6171)

		gen6173 := Call(__e, gen6150, MakeString(">"), gen6172)

		gen6174 := Call(__e, gen6149, MakeString("~"), gen6173)

		gen6175 := Call(__e, gen6148, MakeString("@"), gen6174)

		gen6176 := Call(__e, gen6147, MakeString("!"), gen6175)

		gen6177 := Call(__e, gen6146, MakeString("$"), gen6176)

		gen6178 := Call(__e, gen6145, MakeString("?"), gen6177)

		gen6179 := Call(__e, gen6144, MakeString("_"), gen6178)

		gen6180 := Call(__e, gen6143, MakeString("-"), gen6179)

		gen6181 := Call(__e, gen6142, MakeString("+"), gen6180)

		gen6182 := Call(__e, gen6141, MakeString("/"), gen6181)

		gen6183 := Call(__e, gen6140, MakeString("*"), gen6182)

		gen6184 := Call(__e, gen6139, MakeString("="), gen6183)

		gen6185 := Call(__e, gen6138, MakeString("z"), gen6184)

		gen6186 := Call(__e, gen6137, MakeString("y"), gen6185)

		gen6187 := Call(__e, gen6136, MakeString("x"), gen6186)

		gen6188 := Call(__e, gen6135, MakeString("w"), gen6187)

		gen6189 := Call(__e, gen6134, MakeString("v"), gen6188)

		gen6190 := Call(__e, gen6133, MakeString("u"), gen6189)

		gen6191 := Call(__e, gen6132, MakeString("t"), gen6190)

		gen6192 := Call(__e, gen6131, MakeString("s"), gen6191)

		gen6193 := Call(__e, gen6130, MakeString("r"), gen6192)

		gen6194 := Call(__e, gen6129, MakeString("q"), gen6193)

		gen6195 := Call(__e, gen6128, MakeString("p"), gen6194)

		gen6196 := Call(__e, gen6127, MakeString("o"), gen6195)

		gen6197 := Call(__e, gen6126, MakeString("n"), gen6196)

		gen6198 := Call(__e, gen6125, MakeString("m"), gen6197)

		gen6199 := Call(__e, gen6124, MakeString("l"), gen6198)

		gen6200 := Call(__e, gen6123, MakeString("k"), gen6199)

		gen6201 := Call(__e, gen6122, MakeString("j"), gen6200)

		gen6202 := Call(__e, gen6121, MakeString("i"), gen6201)

		gen6203 := Call(__e, gen6120, MakeString("h"), gen6202)

		gen6204 := Call(__e, gen6119, MakeString("g"), gen6203)

		gen6205 := Call(__e, gen6118, MakeString("f"), gen6204)

		gen6206 := Call(__e, gen6117, MakeString("e"), gen6205)

		gen6207 := Call(__e, gen6116, MakeString("d"), gen6206)

		gen6208 := Call(__e, gen6115, MakeString("c"), gen6207)

		gen6209 := Call(__e, gen6114, MakeString("b"), gen6208)

		gen6210 := Call(__e, gen6113, MakeString("a"), gen6209)

		gen6211 := Call(__e, gen6112, MakeString("Z"), gen6210)

		gen6212 := Call(__e, gen6111, MakeString("Y"), gen6211)

		gen6213 := Call(__e, gen6110, MakeString("X"), gen6212)

		gen6214 := Call(__e, gen6109, MakeString("W"), gen6213)

		gen6215 := Call(__e, gen6108, MakeString("V"), gen6214)

		gen6216 := Call(__e, gen6107, MakeString("U"), gen6215)

		gen6217 := Call(__e, gen6106, MakeString("T"), gen6216)

		gen6218 := Call(__e, gen6105, MakeString("S"), gen6217)

		gen6219 := Call(__e, gen6104, MakeString("R"), gen6218)

		gen6220 := Call(__e, gen6103, MakeString("Q"), gen6219)

		gen6221 := Call(__e, gen6102, MakeString("P"), gen6220)

		gen6222 := Call(__e, gen6101, MakeString("O"), gen6221)

		gen6223 := Call(__e, gen6100, MakeString("N"), gen6222)

		gen6224 := Call(__e, gen6099, MakeString("M"), gen6223)

		gen6225 := Call(__e, gen6098, MakeString("L"), gen6224)

		gen6226 := Call(__e, gen6097, MakeString("K"), gen6225)

		gen6227 := Call(__e, gen6096, MakeString("J"), gen6226)

		gen6228 := Call(__e, gen6095, MakeString("I"), gen6227)

		gen6229 := Call(__e, gen6094, MakeString("H"), gen6228)

		gen6230 := Call(__e, gen6093, MakeString("G"), gen6229)

		gen6231 := Call(__e, gen6092, MakeString("F"), gen6230)

		gen6232 := Call(__e, gen6091, MakeString("E"), gen6231)

		gen6233 := Call(__e, gen6090, MakeString("D"), gen6232)

		gen6234 := Call(__e, gen6089, MakeString("C"), gen6233)

		gen6235 := Call(__e, gen6088, MakeString("B"), gen6234)

		gen6236 := Call(__e, gen6087, MakeString("A"), gen6235)

		__e.TailApply(gen6086, V1999, gen6236)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4alpha_2, gen6237)

	gen6252 := MakeNative(func(__e *ControlFlow) {
		V2001 := __e.Get(1)
		_ = V2001
		gen6250 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6251 := Call(__e, gen6250, MakeString(""), V2001)

		if True == gen6251 {
			__e.Return(True)
			return
		} else {
			gen6248 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			gen6249 := Call(__e, gen6248, V2001)

			if True == gen6249 {
				gen6244 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanum_2)

				gen6245 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen6246 := Call(__e, gen6245, V2001, MakeNumber(0))

				gen6247 := Call(__e, gen6244, gen6246)

				if True == gen6247 {
					gen6240 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanums_2)

					gen6241 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

					gen6242 := Call(__e, gen6241, V2001)

					gen6243 := Call(__e, gen6240, gen6242)

					if True == gen6243 {
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
				if True == True {
					gen6239 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6239, symshen_4alphanums_2)

					return

				} else {
					gen6238 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6238, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4alphanums_2, gen6252)

	gen6257 := MakeNative(func(__e *ControlFlow) {
		V2003 := __e.Get(1)
		_ = V2003
		gen6255 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alpha_2)

		gen6256 := Call(__e, gen6255, V2003)

		if True == gen6256 {
			__e.Return(True)
			return
		} else {
			gen6253 := Call(__e, PrimNS1Value(symns2_1value), symshen_4digit_2)

			gen6254 := Call(__e, gen6253, V2003)

			if True == gen6254 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4alphanum_2, gen6257)

	gen6279 := MakeNative(func(__e *ControlFlow) {
		V2005 := __e.Get(1)
		_ = V2005
		gen6258 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen6259 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6260 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6261 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6262 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6263 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6264 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6265 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6266 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6267 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6268 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6269 := Call(__e, gen6268, MakeString("0"), Nil)

		gen6270 := Call(__e, gen6267, MakeString("9"), gen6269)

		gen6271 := Call(__e, gen6266, MakeString("8"), gen6270)

		gen6272 := Call(__e, gen6265, MakeString("7"), gen6271)

		gen6273 := Call(__e, gen6264, MakeString("6"), gen6272)

		gen6274 := Call(__e, gen6263, MakeString("5"), gen6273)

		gen6275 := Call(__e, gen6262, MakeString("4"), gen6274)

		gen6276 := Call(__e, gen6261, MakeString("3"), gen6275)

		gen6277 := Call(__e, gen6260, MakeString("2"), gen6276)

		gen6278 := Call(__e, gen6259, MakeString("1"), gen6277)

		__e.TailApply(gen6258, V2005, gen6278)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4digit_2, gen6279)

	gen6295 := MakeNative(func(__e *ControlFlow) {
		V2007 := __e.Get(1)
		_ = V2007
		gen6292 := Call(__e, PrimNS1Value(symns2_1value), symboolean_2)

		gen6293 := Call(__e, gen6292, V2007)

		var gen6294 Obj

		if True == gen6293 {
			gen6294 = True
		} else {
			gen6289 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

			gen6290 := Call(__e, gen6289, V2007)

			var gen6291 Obj

			if True == gen6290 {
				gen6291 = True
			} else {
				gen6287 := Call(__e, PrimNS1Value(symns2_1value), symstring_2)

				gen6288 := Call(__e, gen6287, V2007)

				if True == gen6288 {
					gen6291 = True
				} else {
					gen6291 = False
				}

			}

			if True == gen6291 {
				gen6294 = True
			} else {
				gen6294 = False
			}

		}

		if True == gen6294 {
			__e.Return(False)
			return
		} else {
			if True == True {
				gen6285 := MakeNative(func(__e *ControlFlow) {
					gen6282 := MakeNative(func(__e *ControlFlow) {
						String := __e.Get(1)
						_ = String
						gen6281 := Call(__e, PrimNS1Value(symns2_1value), symshen_4analyse_1variable_2)

						__e.TailApply(gen6281, String)

						return

					}, 1)

					gen6283 := Call(__e, PrimNS1Value(symns2_1value), symstr)

					gen6284 := Call(__e, gen6283, V2007)

					__e.TailApply(gen6282, gen6284)

					return

				}, 0)

				gen6286 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(False)
					return
				}, 1)

				__e.TailApply(PrimNS1Value(symtry_1catch), gen6285, gen6286)

				return

			} else {
				gen6280 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6280, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symvariable_2, gen6295)

	gen6308 := MakeNative(func(__e *ControlFlow) {
		V2009 := __e.Get(1)
		_ = V2009
		gen6306 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

		gen6307 := Call(__e, gen6306, V2009)

		if True == gen6307 {
			gen6302 := Call(__e, PrimNS1Value(symns2_1value), symshen_4uppercase_2)

			gen6303 := Call(__e, PrimNS1Value(symns2_1value), sympos)

			gen6304 := Call(__e, gen6303, V2009, MakeNumber(0))

			gen6305 := Call(__e, gen6302, gen6304)

			if True == gen6305 {
				gen6298 := Call(__e, PrimNS1Value(symns2_1value), symshen_4alphanums_2)

				gen6299 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen6300 := Call(__e, gen6299, V2009)

				gen6301 := Call(__e, gen6298, gen6300)

				if True == gen6301 {
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
			if True == True {
				gen6297 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

				__e.TailApply(gen6297, symshen_4analyse_1variable_2)

				return

			} else {
				gen6296 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6296, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4analyse_1variable_2, gen6308)

	gen6362 := MakeNative(func(__e *ControlFlow) {
		V2011 := __e.Get(1)
		_ = V2011
		gen6309 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

		gen6310 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6311 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6312 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6313 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6314 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6315 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6316 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6317 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6318 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6319 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6320 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6321 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6322 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6323 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6324 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6325 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6326 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6327 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6328 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6329 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6330 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6331 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6332 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6333 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6334 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6335 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6336 := Call(__e, gen6335, MakeString("Z"), Nil)

		gen6337 := Call(__e, gen6334, MakeString("Y"), gen6336)

		gen6338 := Call(__e, gen6333, MakeString("X"), gen6337)

		gen6339 := Call(__e, gen6332, MakeString("W"), gen6338)

		gen6340 := Call(__e, gen6331, MakeString("V"), gen6339)

		gen6341 := Call(__e, gen6330, MakeString("U"), gen6340)

		gen6342 := Call(__e, gen6329, MakeString("T"), gen6341)

		gen6343 := Call(__e, gen6328, MakeString("S"), gen6342)

		gen6344 := Call(__e, gen6327, MakeString("R"), gen6343)

		gen6345 := Call(__e, gen6326, MakeString("Q"), gen6344)

		gen6346 := Call(__e, gen6325, MakeString("P"), gen6345)

		gen6347 := Call(__e, gen6324, MakeString("O"), gen6346)

		gen6348 := Call(__e, gen6323, MakeString("N"), gen6347)

		gen6349 := Call(__e, gen6322, MakeString("M"), gen6348)

		gen6350 := Call(__e, gen6321, MakeString("L"), gen6349)

		gen6351 := Call(__e, gen6320, MakeString("K"), gen6350)

		gen6352 := Call(__e, gen6319, MakeString("J"), gen6351)

		gen6353 := Call(__e, gen6318, MakeString("I"), gen6352)

		gen6354 := Call(__e, gen6317, MakeString("H"), gen6353)

		gen6355 := Call(__e, gen6316, MakeString("G"), gen6354)

		gen6356 := Call(__e, gen6315, MakeString("F"), gen6355)

		gen6357 := Call(__e, gen6314, MakeString("E"), gen6356)

		gen6358 := Call(__e, gen6313, MakeString("D"), gen6357)

		gen6359 := Call(__e, gen6312, MakeString("C"), gen6358)

		gen6360 := Call(__e, gen6311, MakeString("B"), gen6359)

		gen6361 := Call(__e, gen6310, MakeString("A"), gen6360)

		__e.TailApply(gen6309, V2011, gen6361)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4uppercase_2, gen6362)

	gen6370 := MakeNative(func(__e *ControlFlow) {
		V2013 := __e.Get(1)
		_ = V2013
		gen6363 := Call(__e, PrimNS1Value(symns2_1value), symconcat)

		gen6364 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen6365 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

		gen6366 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		gen6367 := Call(__e, gen6366, symshen_4_dgensym_d)

		gen6368 := Call(__e, gen6365, MakeNumber(1), gen6367)

		gen6369 := Call(__e, gen6364, symshen_4_dgensym_d, gen6368)

		__e.TailApply(gen6363, V2013, gen6369)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symgensym, gen6370)

	gen6378 := MakeNative(func(__e *ControlFlow) {
		V2016 := __e.Get(1)
		_ = V2016
		V2017 := __e.Get(2)
		_ = V2017
		gen6371 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		gen6372 := Call(__e, PrimNS1Value(symns2_1value), symcn)

		gen6373 := Call(__e, PrimNS1Value(symns2_1value), symstr)

		gen6374 := Call(__e, gen6373, V2016)

		gen6375 := Call(__e, PrimNS1Value(symns2_1value), symstr)

		gen6376 := Call(__e, gen6375, V2017)

		gen6377 := Call(__e, gen6372, gen6374, gen6376)

		__e.TailApply(gen6371, gen6377)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symconcat, gen6378)

	gen6391 := MakeNative(func(__e *ControlFlow) {
		V2020 := __e.Get(1)
		_ = V2020
		V2021 := __e.Get(2)
		_ = V2021
		gen6388 := MakeNative(func(__e *ControlFlow) {
			Vector := __e.Get(1)
			_ = Vector
			gen6385 := MakeNative(func(__e *ControlFlow) {
				Tag := __e.Get(1)
				_ = Tag
				gen6382 := MakeNative(func(__e *ControlFlow) {
					Fst := __e.Get(1)
					_ = Fst
					gen6379 := MakeNative(func(__e *ControlFlow) {
						Snd := __e.Get(1)
						_ = Snd
						__e.Return(Vector)
						return
					}, 1)

					gen6380 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

					gen6381 := Call(__e, gen6380, Vector, MakeNumber(2), V2021)

					__e.TailApply(gen6379, gen6381)

					return

				}, 1)

				gen6383 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

				gen6384 := Call(__e, gen6383, Vector, MakeNumber(1), V2020)

				__e.TailApply(gen6382, gen6384)

				return

			}, 1)

			gen6386 := Call(__e, PrimNS1Value(symns2_1value), symaddress_1_6)

			gen6387 := Call(__e, gen6386, Vector, MakeNumber(0), symshen_4tuple)

			__e.TailApply(gen6385, gen6387)

			return

		}, 1)

		gen6389 := Call(__e, PrimNS1Value(symns2_1value), symabsvector)

		gen6390 := Call(__e, gen6389, MakeNumber(3))

		__e.TailApply(gen6388, gen6390)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), sym_8p, gen6391)

	gen6393 := MakeNative(func(__e *ControlFlow) {
		V2023 := __e.Get(1)
		_ = V2023
		gen6392 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(gen6392, V2023, MakeNumber(1))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symfst, gen6393)

	gen6395 := MakeNative(func(__e *ControlFlow) {
		V2025 := __e.Get(1)
		_ = V2025
		gen6394 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

		__e.TailApply(gen6394, V2025, MakeNumber(2))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symsnd, gen6395)

	gen6404 := MakeNative(func(__e *ControlFlow) {
		V2027 := __e.Get(1)
		_ = V2027
		gen6402 := Call(__e, PrimNS1Value(symns2_1value), symabsvector_2)

		gen6403 := Call(__e, gen6402, V2027)

		if True == gen6403 {
			gen6396 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen6398 := MakeNative(func(__e *ControlFlow) {
				gen6397 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1address)

				__e.TailApply(gen6397, V2027, MakeNumber(0))

				return

			}, 0)

			gen6399 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4not_1tuple)
				return
			}, 1)

			gen6400 := Call(__e, PrimNS1Value(symtry_1catch), gen6398, gen6399)

			gen6401 := Call(__e, gen6396, symshen_4tuple, gen6400)

			if True == gen6401 {
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

	Call(__e, PrimNS1Value(symns2_1set), symtuple_2, gen6404)

	gen6418 := MakeNative(func(__e *ControlFlow) {
		V2030 := __e.Get(1)
		_ = V2030
		V2031 := __e.Get(2)
		_ = V2031
		gen6416 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6417 := Call(__e, gen6416, Nil, V2030)

		if True == gen6417 {
			__e.Return(V2031)
			return
		} else {
			gen6414 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6415 := Call(__e, gen6414, V2030)

			if True == gen6415 {
				gen6407 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen6408 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6409 := Call(__e, gen6408, V2030)

				gen6410 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				gen6411 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6412 := Call(__e, gen6411, V2030)

				gen6413 := Call(__e, gen6410, gen6412, V2031)

				__e.TailApply(gen6407, gen6409, gen6413)

				return

			} else {
				if True == True {
					gen6406 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6406, symappend)

					return

				} else {
					gen6405 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6405, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symappend, gen6418)

	gen6433 := MakeNative(func(__e *ControlFlow) {
		V2034 := __e.Get(1)
		_ = V2034
		V2035 := __e.Get(2)
		_ = V2035
		gen6430 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			gen6425 := MakeNative(func(__e *ControlFlow) {
				NewVector := __e.Get(1)
				_ = NewVector
				gen6422 := MakeNative(func(__e *ControlFlow) {
					X_7NewVector := __e.Get(1)
					_ = X_7NewVector
					gen6420 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen6421 := Call(__e, gen6420, Limit, MakeNumber(0))

					if True == gen6421 {
						__e.Return(X_7NewVector)
						return
					} else {
						gen6419 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8v_1help)

						__e.TailApply(gen6419, V2035, MakeNumber(1), Limit, X_7NewVector)

						return

					}

				}, 1)

				gen6423 := Call(__e, PrimNS1Value(symns2_1value), symvector_1_6)

				gen6424 := Call(__e, gen6423, NewVector, MakeNumber(1), V2034)

				__e.TailApply(gen6422, gen6424)

				return

			}, 1)

			gen6426 := Call(__e, PrimNS1Value(symns2_1value), symvector)

			gen6427 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			gen6428 := Call(__e, gen6427, Limit, MakeNumber(1))

			gen6429 := Call(__e, gen6426, gen6428)

			__e.TailApply(gen6425, gen6429)

			return

		}, 1)

		gen6431 := Call(__e, PrimNS1Value(symns2_1value), symlimit)

		gen6432 := Call(__e, gen6431, V2035)

		__e.TailApply(gen6430, gen6432)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), sym_8v, gen6433)

	gen6447 := MakeNative(func(__e *ControlFlow) {
		V2041 := __e.Get(1)
		_ = V2041
		V2042 := __e.Get(2)
		_ = V2042
		V2043 := __e.Get(3)
		_ = V2043
		V2044 := __e.Get(4)
		_ = V2044
		gen6445 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6446 := Call(__e, gen6445, V2043, V2042)

		if True == gen6446 {
			gen6442 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

			gen6443 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

			gen6444 := Call(__e, gen6443, V2043, MakeNumber(1))

			__e.TailApply(gen6442, V2041, V2044, V2043, gen6444)

			return

		} else {
			if True == True {
				gen6435 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_8v_1help)

				gen6436 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen6437 := Call(__e, gen6436, V2042, MakeNumber(1))

				gen6438 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

				gen6439 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen6440 := Call(__e, gen6439, V2042, MakeNumber(1))

				gen6441 := Call(__e, gen6438, V2041, V2044, V2042, gen6440)

				__e.TailApply(gen6435, V2041, gen6437, V2043, gen6441)

				return

			} else {
				gen6434 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6434, MakeString("no cond match"))

				return

			}
		}

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4_8v_1help, gen6447)

	gen6453 := MakeNative(func(__e *ControlFlow) {
		V2049 := __e.Get(1)
		_ = V2049
		V2050 := __e.Get(2)
		_ = V2050
		V2051 := __e.Get(3)
		_ = V2051
		V2052 := __e.Get(4)
		_ = V2052
		gen6451 := MakeNative(func(__e *ControlFlow) {
			gen6448 := Call(__e, PrimNS1Value(symns2_1value), symvector_1_6)

			gen6449 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1vector)

			gen6450 := Call(__e, gen6449, V2049, V2051)

			__e.TailApply(gen6448, V2050, V2052, gen6450)

			return

		}, 0)

		gen6452 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(V2050)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen6451, gen6452)

		return

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4copyfromvector, gen6453)

	gen6462 := MakeNative(func(__e *ControlFlow) {
		V2054 := __e.Get(1)
		_ = V2054
		gen6455 := MakeNative(func(__e *ControlFlow) {
			gen6454 := Call(__e, PrimNS1Value(symns2_1value), sym_5_1vector)

			__e.TailApply(gen6454, V2054, MakeNumber(1))

			return

		}, 0)

		gen6461 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen6456 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen6457 := Call(__e, PrimNS1Value(symns2_1value), symcn)

			gen6458 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen6459 := Call(__e, gen6458, V2054, MakeString("\n"), symshen_4s)

			gen6460 := Call(__e, gen6457, MakeString("hdv needs a non-empty vector as an argument; not "), gen6459)

			__e.TailApply(gen6456, gen6460)

			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen6455, gen6461)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symhdv, gen6462)

	gen6482 := MakeNative(func(__e *ControlFlow) {
		V2056 := __e.Get(1)
		_ = V2056
		gen6479 := MakeNative(func(__e *ControlFlow) {
			Limit := __e.Get(1)
			_ = Limit
			gen6477 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen6478 := Call(__e, gen6477, Limit, MakeNumber(0))

			if True == gen6478 {
				gen6476 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6476, MakeString("cannot take the tail of the empty vector\n"))

				return

			} else {
				gen6474 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen6475 := Call(__e, gen6474, Limit, MakeNumber(1))

				if True == gen6475 {
					gen6473 := Call(__e, PrimNS1Value(symns2_1value), symvector)

					__e.TailApply(gen6473, MakeNumber(0))

					return

				} else {
					gen6468 := MakeNative(func(__e *ControlFlow) {
						NewVector := __e.Get(1)
						_ = NewVector
						gen6463 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlv_1help)

						gen6464 := Call(__e, PrimNS1Value(symns2_1value), symvector)

						gen6465 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						gen6466 := Call(__e, gen6465, Limit, MakeNumber(1))

						gen6467 := Call(__e, gen6464, gen6466)

						__e.TailApply(gen6463, V2056, MakeNumber(2), Limit, gen6467)

						return

					}, 1)

					gen6469 := Call(__e, PrimNS1Value(symns2_1value), symvector)

					gen6470 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					gen6471 := Call(__e, gen6470, Limit, MakeNumber(1))

					gen6472 := Call(__e, gen6469, gen6471)

					__e.TailApply(gen6468, gen6472)

					return

				}

			}

		}, 1)

		gen6480 := Call(__e, PrimNS1Value(symns2_1value), symlimit)

		gen6481 := Call(__e, gen6480, V2056)

		__e.TailApply(gen6479, gen6481)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symtlv, gen6482)

	gen6496 := MakeNative(func(__e *ControlFlow) {
		V2062 := __e.Get(1)
		_ = V2062
		V2063 := __e.Get(2)
		_ = V2063
		V2064 := __e.Get(3)
		_ = V2064
		V2065 := __e.Get(4)
		_ = V2065
		gen6494 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6495 := Call(__e, gen6494, V2064, V2063)

		if True == gen6495 {
			gen6491 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

			gen6492 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			gen6493 := Call(__e, gen6492, V2064, MakeNumber(1))

			__e.TailApply(gen6491, V2062, V2065, V2064, gen6493)

			return

		} else {
			if True == True {
				gen6484 := Call(__e, PrimNS1Value(symns2_1value), symshen_4tlv_1help)

				gen6485 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen6486 := Call(__e, gen6485, V2063, MakeNumber(1))

				gen6487 := Call(__e, PrimNS1Value(symns2_1value), symshen_4copyfromvector)

				gen6488 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				gen6489 := Call(__e, gen6488, V2063, MakeNumber(1))

				gen6490 := Call(__e, gen6487, V2062, V2065, V2063, gen6489)

				__e.TailApply(gen6484, V2062, gen6486, V2064, gen6490)

				return

			} else {
				gen6483 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6483, MakeString("no cond match"))

				return

			}
		}

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4tlv_1help, gen6496)

	gen6521 := MakeNative(func(__e *ControlFlow) {
		V2077 := __e.Get(1)
		_ = V2077
		V2078 := __e.Get(2)
		_ = V2078
		gen6519 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6520 := Call(__e, gen6519, Nil, V2078)

		if True == gen6520 {
			__e.Return(Nil)
			return
		} else {
			gen6516 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6517 := Call(__e, gen6516, V2078)

			var gen6518 Obj

			if True == gen6517 {
				gen6511 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6512 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6513 := Call(__e, gen6512, V2078)

				gen6514 := Call(__e, gen6511, gen6513)

				var gen6515 Obj

				if True == gen6514 {
					gen6505 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen6506 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6507 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6508 := Call(__e, gen6507, V2078)

					gen6509 := Call(__e, gen6506, gen6508)

					gen6510 := Call(__e, gen6505, gen6509, V2077)

					if True == gen6510 {
						gen6515 = True
					} else {
						gen6515 = False
					}

				} else {
					gen6515 = False
				}

				if True == gen6515 {
					gen6518 = True
				} else {
					gen6518 = False
				}

			} else {
				gen6518 = False
			}

			if True == gen6518 {
				gen6504 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				__e.TailApply(gen6504, V2078)

				return

			} else {
				gen6502 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6503 := Call(__e, gen6502, V2078)

				if True == gen6503 {
					gen6499 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

					gen6500 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6501 := Call(__e, gen6500, V2078)

					__e.TailApply(gen6499, V2077, gen6501)

					return

				} else {
					if True == True {
						gen6498 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen6498, symassoc)

						return

					} else {
						gen6497 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen6497, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symassoc, gen6521)

	gen6561 := MakeNative(func(__e *ControlFlow) {
		V2085 := __e.Get(1)
		_ = V2085
		V2086 := __e.Get(2)
		_ = V2086
		V2087 := __e.Get(3)
		_ = V2087
		gen6559 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6560 := Call(__e, gen6559, Nil, V2087)

		if True == gen6560 {
			gen6556 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen6557 := Call(__e, PrimNS1Value(symns2_1value), symcons)

			gen6558 := Call(__e, gen6557, V2085, V2086)

			__e.TailApply(gen6556, gen6558, Nil)

			return

		} else {
			gen6553 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6554 := Call(__e, gen6553, V2087)

			var gen6555 Obj

			if True == gen6554 {
				gen6548 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6549 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6550 := Call(__e, gen6549, V2087)

				gen6551 := Call(__e, gen6548, gen6550)

				var gen6552 Obj

				if True == gen6551 {
					gen6542 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen6543 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6544 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6545 := Call(__e, gen6544, V2087)

					gen6546 := Call(__e, gen6543, gen6545)

					gen6547 := Call(__e, gen6542, gen6546, V2085)

					if True == gen6547 {
						gen6552 = True
					} else {
						gen6552 = False
					}

				} else {
					gen6552 = False
				}

				if True == gen6552 {
					gen6555 = True
				} else {
					gen6555 = False
				}

			} else {
				gen6555 = False
			}

			if True == gen6555 {
				gen6533 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen6534 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen6535 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6536 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6537 := Call(__e, gen6536, V2087)

				gen6538 := Call(__e, gen6535, gen6537)

				gen6539 := Call(__e, gen6534, gen6538, V2086)

				gen6540 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6541 := Call(__e, gen6540, V2087)

				__e.TailApply(gen6533, gen6539, gen6541)

				return

			} else {
				gen6531 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6532 := Call(__e, gen6531, V2087)

				if True == gen6532 {
					gen6524 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen6525 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6526 := Call(__e, gen6525, V2087)

					gen6527 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1set)

					gen6528 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6529 := Call(__e, gen6528, V2087)

					gen6530 := Call(__e, gen6527, V2085, V2086, gen6529)

					__e.TailApply(gen6524, gen6526, gen6530)

					return

				} else {
					if True == True {
						gen6523 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen6523, symshen_4assoc_1set)

						return

					} else {
						gen6522 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen6522, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1set, gen6561)

	gen6590 := MakeNative(func(__e *ControlFlow) {
		V2093 := __e.Get(1)
		_ = V2093
		V2094 := __e.Get(2)
		_ = V2094
		gen6588 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6589 := Call(__e, gen6588, Nil, V2094)

		if True == gen6589 {
			__e.Return(Nil)
			return
		} else {
			gen6585 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6586 := Call(__e, gen6585, V2094)

			var gen6587 Obj

			if True == gen6586 {
				gen6580 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6581 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6582 := Call(__e, gen6581, V2094)

				gen6583 := Call(__e, gen6580, gen6582)

				var gen6584 Obj

				if True == gen6583 {
					gen6574 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen6575 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6576 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6577 := Call(__e, gen6576, V2094)

					gen6578 := Call(__e, gen6575, gen6577)

					gen6579 := Call(__e, gen6574, gen6578, V2093)

					if True == gen6579 {
						gen6584 = True
					} else {
						gen6584 = False
					}

				} else {
					gen6584 = False
				}

				if True == gen6584 {
					gen6587 = True
				} else {
					gen6587 = False
				}

			} else {
				gen6587 = False
			}

			if True == gen6587 {
				gen6573 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				__e.TailApply(gen6573, V2094)

				return

			} else {
				gen6571 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6572 := Call(__e, gen6571, V2094)

				if True == gen6572 {
					gen6564 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen6565 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6566 := Call(__e, gen6565, V2094)

					gen6567 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1rm)

					gen6568 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6569 := Call(__e, gen6568, V2094)

					gen6570 := Call(__e, gen6567, V2093, gen6569)

					__e.TailApply(gen6564, gen6566, gen6570)

					return

				} else {
					if True == True {
						gen6563 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen6563, symshen_4assoc_1rm)

						return

					} else {
						gen6562 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen6562, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4assoc_1rm, gen6590)

	gen6596 := MakeNative(func(__e *ControlFlow) {
		V2100 := __e.Get(1)
		_ = V2100
		gen6594 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6595 := Call(__e, gen6594, True, V2100)

		if True == gen6595 {
			__e.Return(True)
			return
		} else {
			gen6592 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen6593 := Call(__e, gen6592, False, V2100)

			if True == gen6593 {
				__e.Return(True)
				return
			} else {
				if True == True {
					__e.Return(False)
					return
				} else {
					gen6591 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6591, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symboolean_2, gen6596)

	gen6606 := MakeNative(func(__e *ControlFlow) {
		V2102 := __e.Get(1)
		_ = V2102
		gen6604 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6605 := Call(__e, gen6604, MakeNumber(0), V2102)

		if True == gen6605 {
			__e.Return(MakeNumber(0))
			return
		} else {
			if True == True {
				gen6598 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

				gen6599 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

				gen6600 := Call(__e, gen6599)

				Call(__e, gen6598, MakeString("\n"), gen6600)

				gen6601 := Call(__e, PrimNS1Value(symns2_1value), symnl)

				gen6602 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				gen6603 := Call(__e, gen6602, V2102, MakeNumber(1))

				__e.TailApply(gen6601, gen6603)

				return

			} else {
				gen6597 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6597, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symnl, gen6606)

	gen6627 := MakeNative(func(__e *ControlFlow) {
		V2107 := __e.Get(1)
		_ = V2107
		V2108 := __e.Get(2)
		_ = V2108
		gen6625 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6626 := Call(__e, gen6625, Nil, V2107)

		if True == gen6626 {
			__e.Return(Nil)
			return
		} else {
			gen6623 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6624 := Call(__e, gen6623, V2107)

			if True == gen6624 {
				gen6619 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen6620 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6621 := Call(__e, gen6620, V2107)

				gen6622 := Call(__e, gen6619, gen6621, V2108)

				if True == gen6622 {
					gen6616 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

					gen6617 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6618 := Call(__e, gen6617, V2107)

					__e.TailApply(gen6616, gen6618, V2108)

					return

				} else {
					gen6609 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen6610 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6611 := Call(__e, gen6610, V2107)

					gen6612 := Call(__e, PrimNS1Value(symns2_1value), symdifference)

					gen6613 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6614 := Call(__e, gen6613, V2107)

					gen6615 := Call(__e, gen6612, gen6614, V2108)

					__e.TailApply(gen6609, gen6611, gen6615)

					return

				}

			} else {
				if True == True {
					gen6608 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6608, symdifference)

					return

				} else {
					gen6607 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6607, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symdifference, gen6627)

	gen6628 := MakeNative(func(__e *ControlFlow) {
		V2111 := __e.Get(1)
		_ = V2111
		V2112 := __e.Get(2)
		_ = V2112
		__e.Return(V2112)
		return
	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symdo, gen6628)

	gen6645 := MakeNative(func(__e *ControlFlow) {
		V2124 := __e.Get(1)
		_ = V2124
		V2125 := __e.Get(2)
		_ = V2125
		gen6643 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6644 := Call(__e, gen6643, Nil, V2125)

		if True == gen6644 {
			__e.Return(False)
			return
		} else {
			gen6640 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6641 := Call(__e, gen6640, V2125)

			var gen6642 Obj

			if True == gen6641 {
				gen6636 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen6637 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6638 := Call(__e, gen6637, V2125)

				gen6639 := Call(__e, gen6636, gen6638, V2124)

				if True == gen6639 {
					gen6642 = True
				} else {
					gen6642 = False
				}

			} else {
				gen6642 = False
			}

			if True == gen6642 {
				__e.Return(True)
				return
			} else {
				gen6634 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6635 := Call(__e, gen6634, V2125)

				if True == gen6635 {
					gen6631 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

					gen6632 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6633 := Call(__e, gen6632, V2125)

					__e.TailApply(gen6631, V2124, gen6633)

					return

				} else {
					if True == True {
						gen6630 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

						__e.TailApply(gen6630, symelement_2)

						return

					} else {
						gen6629 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

						__e.TailApply(gen6629, MakeString("no cond match"))

						return

					}
				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symelement_2, gen6645)

	gen6649 := MakeNative(func(__e *ControlFlow) {
		V2131 := __e.Get(1)
		_ = V2131
		gen6647 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6648 := Call(__e, gen6647, Nil, V2131)

		if True == gen6648 {
			__e.Return(True)
			return
		} else {
			if True == True {
				__e.Return(False)
				return
			} else {
				gen6646 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6646, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symempty_2, gen6649)

	gen6652 := MakeNative(func(__e *ControlFlow) {
		V2134 := __e.Get(1)
		_ = V2134
		V2135 := __e.Get(2)
		_ = V2135
		gen6650 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fix_1help)

		gen6651 := Call(__e, V2134, V2135)

		__e.TailApply(gen6650, V2134, V2135, gen6651)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symfix, gen6652)

	gen6658 := MakeNative(func(__e *ControlFlow) {
		V2146 := __e.Get(1)
		_ = V2146
		V2147 := __e.Get(2)
		_ = V2147
		V2148 := __e.Get(3)
		_ = V2148
		gen6656 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6657 := Call(__e, gen6656, V2148, V2147)

		if True == gen6657 {
			__e.Return(V2148)
			return
		} else {
			if True == True {
				gen6654 := Call(__e, PrimNS1Value(symns2_1value), symshen_4fix_1help)

				gen6655 := Call(__e, V2146, V2148)

				__e.TailApply(gen6654, V2146, V2148, gen6655)

				return

			} else {
				gen6653 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6653, MakeString("no cond match"))

				return

			}
		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4fix_1help, gen6658)

	gen6670 := MakeNative(func(__e *ControlFlow) {
		V2153 := __e.Get(1)
		_ = V2153
		V2154 := __e.Get(2)
		_ = V2154
		V2155 := __e.Get(3)
		_ = V2155
		V2156 := __e.Get(4)
		_ = V2156
		gen6665 := MakeNative(func(__e *ControlFlow) {
			Curr := __e.Get(1)
			_ = Curr
			gen6662 := MakeNative(func(__e *ControlFlow) {
				Added := __e.Get(1)
				_ = Added
				gen6659 := MakeNative(func(__e *ControlFlow) {
					Update := __e.Get(1)
					_ = Update
					__e.Return(V2155)
					return
				}, 1)

				gen6660 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1_6)

				gen6661 := Call(__e, gen6660, V2156, V2153, Added)

				__e.TailApply(gen6659, gen6661)

				return

			}, 1)

			gen6663 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1set)

			gen6664 := Call(__e, gen6663, V2154, V2155, Curr)

			__e.TailApply(gen6662, gen6664)

			return

		}, 1)

		gen6667 := MakeNative(func(__e *ControlFlow) {
			gen6666 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict)

			__e.TailApply(gen6666, V2156, V2153)

			return

		}, 0)

		gen6668 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		gen6669 := Call(__e, PrimNS1Value(symtry_1catch), gen6667, gen6668)

		__e.TailApply(gen6665, gen6669)

		return

	}, 4)

	Call(__e, PrimNS1Value(symns2_1set), symput, gen6670)

	gen6682 := MakeNative(func(__e *ControlFlow) {
		V2160 := __e.Get(1)
		_ = V2160
		V2161 := __e.Get(2)
		_ = V2161
		V2162 := __e.Get(3)
		_ = V2162
		gen6677 := MakeNative(func(__e *ControlFlow) {
			Curr := __e.Get(1)
			_ = Curr
			gen6674 := MakeNative(func(__e *ControlFlow) {
				Removed := __e.Get(1)
				_ = Removed
				gen6671 := MakeNative(func(__e *ControlFlow) {
					Update := __e.Get(1)
					_ = Update
					__e.Return(V2160)
					return
				}, 1)

				gen6672 := Call(__e, PrimNS1Value(symns2_1value), symshen_4dict_1_6)

				gen6673 := Call(__e, gen6672, V2162, V2160, Removed)

				__e.TailApply(gen6671, gen6673)

				return

			}, 1)

			gen6675 := Call(__e, PrimNS1Value(symns2_1value), symshen_4assoc_1rm)

			gen6676 := Call(__e, gen6675, V2161, Curr)

			__e.TailApply(gen6674, gen6676)

			return

		}, 1)

		gen6679 := MakeNative(func(__e *ControlFlow) {
			gen6678 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict)

			__e.TailApply(gen6678, V2162, V2160)

			return

		}, 0)

		gen6680 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		gen6681 := Call(__e, PrimNS1Value(symtry_1catch), gen6679, gen6680)

		__e.TailApply(gen6677, gen6681)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symunput, gen6682)

	gen6695 := MakeNative(func(__e *ControlFlow) {
		V2166 := __e.Get(1)
		_ = V2166
		V2167 := __e.Get(2)
		_ = V2167
		V2168 := __e.Get(3)
		_ = V2168
		gen6690 := MakeNative(func(__e *ControlFlow) {
			Entry := __e.Get(1)
			_ = Entry
			gen6687 := MakeNative(func(__e *ControlFlow) {
				Result := __e.Get(1)
				_ = Result
				gen6685 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

				gen6686 := Call(__e, gen6685, Result)

				if True == gen6686 {
					gen6684 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6684, MakeString("value not found\n"))

					return

				} else {
					gen6683 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					__e.TailApply(gen6683, Result)

					return

				}

			}, 1)

			gen6688 := Call(__e, PrimNS1Value(symns2_1value), symassoc)

			gen6689 := Call(__e, gen6688, V2167, Entry)

			__e.TailApply(gen6687, gen6689)

			return

		}, 1)

		gen6692 := MakeNative(func(__e *ControlFlow) {
			gen6691 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_5_1dict)

			__e.TailApply(gen6691, V2168, V2166)

			return

		}, 0)

		gen6693 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(Nil)
			return
		}, 1)

		gen6694 := Call(__e, PrimNS1Value(symtry_1catch), gen6692, gen6693)

		__e.TailApply(gen6690, gen6694)

		return

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symget, gen6695)

	gen6705 := MakeNative(func(__e *ControlFlow) {
		V2171 := __e.Get(1)
		_ = V2171
		V2172 := __e.Get(2)
		_ = V2172
		gen6696 := Call(__e, PrimNS1Value(symns2_1value), symshen_4mod)

		gen6697 := Call(__e, PrimNS1Value(symns2_1value), symsum)

		gen6698 := Call(__e, PrimNS1Value(symns2_1value), symmap)

		gen6700 := MakeNative(func(__e *ControlFlow) {
			X := __e.Get(1)
			_ = X
			gen6699 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

			__e.TailApply(gen6699, X)

			return

		}, 1)

		gen6701 := Call(__e, PrimNS1Value(symns2_1value), symexplode)

		gen6702 := Call(__e, gen6701, V2171)

		gen6703 := Call(__e, gen6698, gen6700, gen6702)

		gen6704 := Call(__e, gen6697, gen6703)

		__e.TailApply(gen6696, gen6704, V2172)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symhash, gen6705)

	gen6711 := MakeNative(func(__e *ControlFlow) {
		V2175 := __e.Get(1)
		_ = V2175
		V2176 := __e.Get(2)
		_ = V2176
		gen6706 := Call(__e, PrimNS1Value(symns2_1value), symshen_4modh)

		gen6707 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiples)

		gen6708 := Call(__e, PrimNS1Value(symns2_1value), symcons)

		gen6709 := Call(__e, gen6708, V2176, Nil)

		gen6710 := Call(__e, gen6707, V2175, gen6709)

		__e.TailApply(gen6706, V2175, gen6710)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4mod, gen6711)

	gen6731 := MakeNative(func(__e *ControlFlow) {
		V2179 := __e.Get(1)
		_ = V2179
		V2180 := __e.Get(2)
		_ = V2180
		gen6728 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen6729 := Call(__e, gen6728, V2180)

		var gen6730 Obj

		if True == gen6729 {
			gen6724 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			gen6725 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			gen6726 := Call(__e, gen6725, V2180)

			gen6727 := Call(__e, gen6724, gen6726, V2179)

			if True == gen6727 {
				gen6730 = True
			} else {
				gen6730 = False
			}

		} else {
			gen6730 = False
		}

		if True == gen6730 {
			gen6723 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(gen6723, V2180)

			return

		} else {
			gen6721 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6722 := Call(__e, gen6721, V2180)

			if True == gen6722 {
				gen6714 := Call(__e, PrimNS1Value(symns2_1value), symshen_4multiples)

				gen6715 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen6716 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

				gen6717 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6718 := Call(__e, gen6717, V2180)

				gen6719 := Call(__e, gen6716, MakeNumber(2), gen6718)

				gen6720 := Call(__e, gen6715, gen6719, V2180)

				__e.TailApply(gen6714, V2179, gen6720)

				return

			} else {
				if True == True {
					gen6713 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6713, symshen_4multiples)

					return

				} else {
					gen6712 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6712, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4multiples, gen6731)

	gen6759 := MakeNative(func(__e *ControlFlow) {
		V2185 := __e.Get(1)
		_ = V2185
		V2186 := __e.Get(2)
		_ = V2186
		gen6757 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6758 := Call(__e, gen6757, MakeNumber(0), V2185)

		if True == gen6758 {
			__e.Return(MakeNumber(0))
			return
		} else {
			gen6755 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen6756 := Call(__e, gen6755, Nil, V2186)

			if True == gen6756 {
				__e.Return(V2185)
				return
			} else {
				gen6752 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

				gen6753 := Call(__e, gen6752, V2186)

				var gen6754 Obj

				if True == gen6753 {
					gen6748 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

					gen6749 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6750 := Call(__e, gen6749, V2186)

					gen6751 := Call(__e, gen6748, gen6750, V2185)

					if True == gen6751 {
						gen6754 = True
					} else {
						gen6754 = False
					}

				} else {
					gen6754 = False
				}

				if True == gen6754 {
					gen6744 := Call(__e, PrimNS1Value(symns2_1value), symempty_2)

					gen6745 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6746 := Call(__e, gen6745, V2186)

					gen6747 := Call(__e, gen6744, gen6746)

					if True == gen6747 {
						__e.Return(V2185)
						return
					} else {
						gen6741 := Call(__e, PrimNS1Value(symns2_1value), symshen_4modh)

						gen6742 := Call(__e, PrimNS1Value(symns2_1value), symtl)

						gen6743 := Call(__e, gen6742, V2186)

						__e.TailApply(gen6741, V2185, gen6743)

						return

					}

				} else {
					gen6739 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

					gen6740 := Call(__e, gen6739, V2186)

					if True == gen6740 {
						gen6734 := Call(__e, PrimNS1Value(symns2_1value), symshen_4modh)

						gen6735 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

						gen6736 := Call(__e, PrimNS1Value(symns2_1value), symhd)

						gen6737 := Call(__e, gen6736, V2186)

						gen6738 := Call(__e, gen6735, V2185, gen6737)

						__e.TailApply(gen6734, gen6738, V2186)

						return

					} else {
						if True == True {
							gen6733 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

							__e.TailApply(gen6733, symshen_4modh)

							return

						} else {
							gen6732 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

							__e.TailApply(gen6732, MakeString("no cond match"))

							return

						}
					}

				}

			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4modh, gen6759)

	gen6773 := MakeNative(func(__e *ControlFlow) {
		V2188 := __e.Get(1)
		_ = V2188
		gen6771 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6772 := Call(__e, gen6771, Nil, V2188)

		if True == gen6772 {
			__e.Return(MakeNumber(0))
			return
		} else {
			gen6769 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6770 := Call(__e, gen6769, V2188)

			if True == gen6770 {
				gen6762 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen6763 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6764 := Call(__e, gen6763, V2188)

				gen6765 := Call(__e, PrimNS1Value(symns2_1value), symsum)

				gen6766 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6767 := Call(__e, gen6766, V2188)

				gen6768 := Call(__e, gen6765, gen6767)

				__e.TailApply(gen6762, gen6764, gen6768)

				return

			} else {
				if True == True {
					gen6761 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6761, symsum)

					return

				} else {
					gen6760 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6760, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symsum, gen6773)

	gen6779 := MakeNative(func(__e *ControlFlow) {
		V2196 := __e.Get(1)
		_ = V2196
		gen6777 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen6778 := Call(__e, gen6777, V2196)

		if True == gen6778 {
			gen6776 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(gen6776, V2196)

			return

		} else {
			if True == True {
				gen6775 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6775, MakeString("head expects a non-empty list"))

				return

			} else {
				gen6774 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6774, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symhead, gen6779)

	gen6785 := MakeNative(func(__e *ControlFlow) {
		V2204 := __e.Get(1)
		_ = V2204
		gen6783 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

		gen6784 := Call(__e, gen6783, V2204)

		if True == gen6784 {
			gen6782 := Call(__e, PrimNS1Value(symns2_1value), symtl)

			__e.TailApply(gen6782, V2204)

			return

		} else {
			if True == True {
				gen6781 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6781, MakeString("tail expects a non-empty list"))

				return

			} else {
				gen6780 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6780, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symtail, gen6785)

	gen6787 := MakeNative(func(__e *ControlFlow) {
		V2206 := __e.Get(1)
		_ = V2206
		gen6786 := Call(__e, PrimNS1Value(symns2_1value), sympos)

		__e.TailApply(gen6786, V2206, MakeNumber(0))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symhdstr, gen6787)

	gen6808 := MakeNative(func(__e *ControlFlow) {
		V2211 := __e.Get(1)
		_ = V2211
		V2212 := __e.Get(2)
		_ = V2212
		gen6806 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6807 := Call(__e, gen6806, Nil, V2211)

		if True == gen6807 {
			__e.Return(Nil)
			return
		} else {
			gen6804 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6805 := Call(__e, gen6804, V2211)

			if True == gen6805 {
				gen6800 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen6801 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6802 := Call(__e, gen6801, V2211)

				gen6803 := Call(__e, gen6800, gen6802, V2212)

				if True == gen6803 {
					gen6793 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen6794 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6795 := Call(__e, gen6794, V2211)

					gen6796 := Call(__e, PrimNS1Value(symns2_1value), symintersection)

					gen6797 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6798 := Call(__e, gen6797, V2211)

					gen6799 := Call(__e, gen6796, gen6798, V2212)

					__e.TailApply(gen6793, gen6795, gen6799)

					return

				} else {
					gen6790 := Call(__e, PrimNS1Value(symns2_1value), symintersection)

					gen6791 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6792 := Call(__e, gen6791, V2211)

					__e.TailApply(gen6790, gen6792, V2212)

					return

				}

			} else {
				if True == True {
					gen6789 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6789, symintersection)

					return

				} else {
					gen6788 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6788, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symintersection, gen6808)

	gen6810 := MakeNative(func(__e *ControlFlow) {
		V2214 := __e.Get(1)
		_ = V2214
		gen6809 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reverse__help)

		__e.TailApply(gen6809, V2214, Nil)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symreverse, gen6810)

	gen6824 := MakeNative(func(__e *ControlFlow) {
		V2217 := __e.Get(1)
		_ = V2217
		V2218 := __e.Get(2)
		_ = V2218
		gen6822 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6823 := Call(__e, gen6822, Nil, V2217)

		if True == gen6823 {
			__e.Return(V2218)
			return
		} else {
			gen6820 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6821 := Call(__e, gen6820, V2217)

			if True == gen6821 {
				gen6813 := Call(__e, PrimNS1Value(symns2_1value), symshen_4reverse__help)

				gen6814 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6815 := Call(__e, gen6814, V2217)

				gen6816 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen6817 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6818 := Call(__e, gen6817, V2217)

				gen6819 := Call(__e, gen6816, gen6818, V2218)

				__e.TailApply(gen6813, gen6815, gen6819)

				return

			} else {
				if True == True {
					gen6812 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6812, symshen_4reverse__help)

					return

				} else {
					gen6811 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6811, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4reverse__help, gen6824)

	gen6845 := MakeNative(func(__e *ControlFlow) {
		V2221 := __e.Get(1)
		_ = V2221
		V2222 := __e.Get(2)
		_ = V2222
		gen6843 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6844 := Call(__e, gen6843, Nil, V2221)

		if True == gen6844 {
			__e.Return(V2222)
			return
		} else {
			gen6841 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6842 := Call(__e, gen6841, V2221)

			if True == gen6842 {
				gen6837 := Call(__e, PrimNS1Value(symns2_1value), symelement_2)

				gen6838 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6839 := Call(__e, gen6838, V2221)

				gen6840 := Call(__e, gen6837, gen6839, V2222)

				if True == gen6840 {
					gen6834 := Call(__e, PrimNS1Value(symns2_1value), symunion)

					gen6835 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6836 := Call(__e, gen6835, V2221)

					__e.TailApply(gen6834, gen6836, V2222)

					return

				} else {
					gen6827 := Call(__e, PrimNS1Value(symns2_1value), symcons)

					gen6828 := Call(__e, PrimNS1Value(symns2_1value), symhd)

					gen6829 := Call(__e, gen6828, V2221)

					gen6830 := Call(__e, PrimNS1Value(symns2_1value), symunion)

					gen6831 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6832 := Call(__e, gen6831, V2221)

					gen6833 := Call(__e, gen6830, gen6832, V2222)

					__e.TailApply(gen6827, gen6829, gen6833)

					return

				}

			} else {
				if True == True {
					gen6826 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6826, symunion)

					return

				} else {
					gen6825 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6825, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symunion, gen6845)

	gen6873 := MakeNative(func(__e *ControlFlow) {
		V2224 := __e.Get(1)
		_ = V2224
		gen6866 := MakeNative(func(__e *ControlFlow) {
			Message := __e.Get(1)
			_ = Message
			gen6861 := MakeNative(func(__e *ControlFlow) {
				Y_1or_1N := __e.Get(1)
				_ = Y_1or_1N
				gen6854 := MakeNative(func(__e *ControlFlow) {
					Input := __e.Get(1)
					_ = Input
					gen6852 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

					gen6853 := Call(__e, gen6852, MakeString("y"), Input)

					if True == gen6853 {
						__e.Return(True)
						return
					} else {
						gen6850 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

						gen6851 := Call(__e, gen6850, MakeString("n"), Input)

						if True == gen6851 {
							__e.Return(False)
							return
						} else {
							gen6846 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

							gen6847 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

							gen6848 := Call(__e, gen6847)

							Call(__e, gen6846, MakeString("please answer y or n\n"), gen6848)

							gen6849 := Call(__e, PrimNS1Value(symns2_1value), symy_1or_1n_2)

							__e.TailApply(gen6849, V2224)

							return

						}

					}

				}, 1)

				gen6855 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen6856 := Call(__e, PrimNS1Value(symns2_1value), symread)

				gen6857 := Call(__e, PrimNS1Value(symns2_1value), symstinput)

				gen6858 := Call(__e, gen6857)

				gen6859 := Call(__e, gen6856, gen6858)

				gen6860 := Call(__e, gen6855, gen6859, MakeString(""), symshen_4s)

				__e.TailApply(gen6854, gen6860)

				return

			}, 1)

			gen6862 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

			gen6863 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

			gen6864 := Call(__e, gen6863)

			gen6865 := Call(__e, gen6862, MakeString(" (y/n) "), gen6864)

			__e.TailApply(gen6861, gen6865)

			return

		}, 1)

		gen6867 := Call(__e, PrimNS1Value(symns2_1value), symshen_4prhush)

		gen6868 := Call(__e, PrimNS1Value(symns2_1value), symshen_4proc_1nl)

		gen6869 := Call(__e, gen6868, V2224)

		gen6870 := Call(__e, PrimNS1Value(symns2_1value), symstoutput)

		gen6871 := Call(__e, gen6870)

		gen6872 := Call(__e, gen6867, gen6869, gen6871)

		__e.TailApply(gen6866, gen6872)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symy_1or_1n_2, gen6873)

	gen6874 := MakeNative(func(__e *ControlFlow) {
		V2226 := __e.Get(1)
		_ = V2226
		if True == V2226 {
			__e.Return(False)
			return
		} else {
			__e.Return(True)
			return
		}
	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symnot, gen6874)

	gen6883 := MakeNative(func(__e *ControlFlow) {
		V2239 := __e.Get(1)
		_ = V2239
		V2240 := __e.Get(2)
		_ = V2240
		V2241 := __e.Get(3)
		_ = V2241
		gen6881 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6882 := Call(__e, gen6881, V2241, V2240)

		if True == gen6882 {
			__e.Return(V2239)
			return
		} else {
			gen6879 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6880 := Call(__e, gen6879, V2241)

			if True == gen6880 {
				gen6876 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				gen6878 := MakeNative(func(__e *ControlFlow) {
					W := __e.Get(1)
					_ = W
					gen6877 := Call(__e, PrimNS1Value(symns2_1value), symsubst)

					__e.TailApply(gen6877, V2239, V2240, W)

					return

				}, 1)

				__e.TailApply(gen6876, gen6878, V2241)

				return

			} else {
				if True == True {
					__e.Return(V2241)
					return
				} else {
					gen6875 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6875, MakeString("no cond match"))

					return

				}
			}

		}

	}, 3)

	Call(__e, PrimNS1Value(symns2_1set), symsubst, gen6883)

	gen6887 := MakeNative(func(__e *ControlFlow) {
		V2243 := __e.Get(1)
		_ = V2243
		gen6884 := Call(__e, PrimNS1Value(symns2_1value), symshen_4explode_1h)

		gen6885 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

		gen6886 := Call(__e, gen6885, V2243, MakeString(""), symshen_4a)

		__e.TailApply(gen6884, gen6886)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symexplode, gen6887)

	gen6901 := MakeNative(func(__e *ControlFlow) {
		V2245 := __e.Get(1)
		_ = V2245
		gen6899 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6900 := Call(__e, gen6899, MakeString(""), V2245)

		if True == gen6900 {
			__e.Return(Nil)
			return
		} else {
			gen6897 := Call(__e, PrimNS1Value(symns2_1value), symshen_4_7string_2)

			gen6898 := Call(__e, gen6897, V2245)

			if True == gen6898 {
				gen6890 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen6891 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen6892 := Call(__e, gen6891, V2245, MakeNumber(0))

				gen6893 := Call(__e, PrimNS1Value(symns2_1value), symshen_4explode_1h)

				gen6894 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen6895 := Call(__e, gen6894, V2245)

				gen6896 := Call(__e, gen6893, gen6895)

				__e.TailApply(gen6890, gen6892, gen6896)

				return

			} else {
				if True == True {
					gen6889 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6889, symshen_4explode_1h)

					return

				} else {
					gen6888 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6888, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4explode_1h, gen6901)

	gen6907 := MakeNative(func(__e *ControlFlow) {
		V2247 := __e.Get(1)
		_ = V2247
		gen6902 := Call(__e, PrimNS1Value(symns2_1value), symset)

		gen6904 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6905 := Call(__e, gen6904, V2247, MakeString(""))

		var gen6906 Obj

		if True == gen6905 {
			gen6906 = MakeString("")
		} else {
			gen6903 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen6906 = Call(__e, gen6903, V2247, MakeString("/"), symshen_4a)

		}

		__e.TailApply(gen6902, sym_dhome_1directory_d, gen6906)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symcd, gen6907)

	gen6921 := MakeNative(func(__e *ControlFlow) {
		V2250 := __e.Get(1)
		_ = V2250
		V2251 := __e.Get(2)
		_ = V2251
		gen6919 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6920 := Call(__e, gen6919, Nil, V2251)

		if True == gen6920 {
			__e.Return(True)
			return
		} else {
			gen6917 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6918 := Call(__e, gen6917, V2251)

			if True == gen6918 {
				gen6913 := MakeNative(func(__e *ControlFlow) {
					__ := __e.Get(1)
					_ = __
					gen6910 := Call(__e, PrimNS1Value(symns2_1value), symshen_4for_1each)

					gen6911 := Call(__e, PrimNS1Value(symns2_1value), symtl)

					gen6912 := Call(__e, gen6911, V2251)

					__e.TailApply(gen6910, V2250, gen6912)

					return

				}, 1)

				gen6914 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6915 := Call(__e, gen6914, V2251)

				gen6916 := Call(__e, V2250, gen6915)

				__e.TailApply(gen6913, gen6916)

				return

			} else {
				if True == True {
					gen6909 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen6909, symshen_4for_1each)

					return

				} else {
					gen6908 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6908, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4for_1each, gen6921)

	gen6935 := MakeNative(func(__e *ControlFlow) {
		V2256 := __e.Get(1)
		_ = V2256
		V2257 := __e.Get(2)
		_ = V2257
		gen6933 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6934 := Call(__e, gen6933, Nil, V2257)

		if True == gen6934 {
			__e.Return(Nil)
			return
		} else {
			gen6931 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6932 := Call(__e, gen6931, V2257)

			if True == gen6932 {
				gen6923 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen6924 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6925 := Call(__e, gen6924, V2257)

				gen6926 := Call(__e, V2256, gen6925)

				gen6927 := Call(__e, PrimNS1Value(symns2_1value), symmap)

				gen6928 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6929 := Call(__e, gen6928, V2257)

				gen6930 := Call(__e, gen6927, V2256, gen6929)

				__e.TailApply(gen6923, gen6926, gen6930)

				return

			} else {
				if True == True {
					__e.TailApply(V2256, V2257)

					return
				} else {
					gen6922 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6922, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symmap, gen6935)

	gen6937 := MakeNative(func(__e *ControlFlow) {
		V2259 := __e.Get(1)
		_ = V2259
		gen6936 := Call(__e, PrimNS1Value(symns2_1value), symshen_4length_1h)

		__e.TailApply(gen6936, V2259, MakeNumber(0))

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symlength, gen6937)

	gen6946 := MakeNative(func(__e *ControlFlow) {
		V2262 := __e.Get(1)
		_ = V2262
		V2263 := __e.Get(2)
		_ = V2263
		gen6944 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6945 := Call(__e, gen6944, Nil, V2262)

		if True == gen6945 {
			__e.Return(V2263)
			return
		} else {
			if True == True {
				gen6939 := Call(__e, PrimNS1Value(symns2_1value), symshen_4length_1h)

				gen6940 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6941 := Call(__e, gen6940, V2262)

				gen6942 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen6943 := Call(__e, gen6942, V2263, MakeNumber(1))

				__e.TailApply(gen6939, gen6941, gen6943)

				return

			} else {
				gen6938 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen6938, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4length_1h, gen6946)

	gen6961 := MakeNative(func(__e *ControlFlow) {
		V2275 := __e.Get(1)
		_ = V2275
		V2276 := __e.Get(2)
		_ = V2276
		gen6959 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6960 := Call(__e, gen6959, V2276, V2275)

		if True == gen6960 {
			__e.Return(MakeNumber(1))
			return
		} else {
			gen6957 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6958 := Call(__e, gen6957, V2276)

			if True == gen6958 {
				gen6948 := Call(__e, PrimNS1Value(symns2_1value), sym_7)

				gen6949 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

				gen6950 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen6951 := Call(__e, gen6950, V2276)

				gen6952 := Call(__e, gen6949, V2275, gen6951)

				gen6953 := Call(__e, PrimNS1Value(symns2_1value), symoccurrences)

				gen6954 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6955 := Call(__e, gen6954, V2276)

				gen6956 := Call(__e, gen6953, V2275, gen6955)

				__e.TailApply(gen6948, gen6952, gen6956)

				return

			} else {
				if True == True {
					__e.Return(MakeNumber(0))
					return
				} else {
					gen6947 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6947, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symoccurrences, gen6961)

	gen6985 := MakeNative(func(__e *ControlFlow) {
		V2283 := __e.Get(1)
		_ = V2283
		V2284 := __e.Get(2)
		_ = V2284
		gen6982 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen6983 := Call(__e, gen6982, MakeNumber(1), V2283)

		var gen6984 Obj

		if True == gen6983 {
			gen6980 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6981 := Call(__e, gen6980, V2284)

			if True == gen6981 {
				gen6984 = True
			} else {
				gen6984 = False
			}

		} else {
			gen6984 = False
		}

		if True == gen6984 {
			gen6979 := Call(__e, PrimNS1Value(symns2_1value), symhd)

			__e.TailApply(gen6979, V2284)

			return

		} else {
			gen6977 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen6978 := Call(__e, gen6977, V2284)

			if True == gen6978 {
				gen6972 := Call(__e, PrimNS1Value(symns2_1value), symnth)

				gen6973 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

				gen6974 := Call(__e, gen6973, V2283, MakeNumber(1))

				gen6975 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen6976 := Call(__e, gen6975, V2284)

				__e.TailApply(gen6972, gen6974, gen6976)

				return

			} else {
				if True == True {
					gen6963 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					gen6964 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen6965 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen6966 := Call(__e, PrimNS1Value(symns2_1value), symcn)

					gen6967 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

					gen6968 := Call(__e, gen6967, V2284, MakeString("\n"), symshen_4a)

					gen6969 := Call(__e, gen6966, MakeString(", "), gen6968)

					gen6970 := Call(__e, gen6965, V2283, gen6969, symshen_4a)

					gen6971 := Call(__e, gen6964, MakeString("nth applied to "), gen6970)

					__e.TailApply(gen6963, gen6971)

					return

				} else {
					gen6962 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen6962, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symnth, gen6985)

	gen6995 := MakeNative(func(__e *ControlFlow) {
		V2286 := __e.Get(1)
		_ = V2286
		gen6993 := Call(__e, PrimNS1Value(symns2_1value), symnumber_2)

		gen6994 := Call(__e, gen6993, V2286)

		if True == gen6994 {
			gen6989 := MakeNative(func(__e *ControlFlow) {
				Abs := __e.Get(1)
				_ = Abs
				gen6986 := Call(__e, PrimNS1Value(symns2_1value), symshen_4integer_1test_2)

				gen6987 := Call(__e, PrimNS1Value(symns2_1value), symshen_4magless)

				gen6988 := Call(__e, gen6987, Abs, MakeNumber(1))

				__e.TailApply(gen6986, Abs, gen6988)

				return

			}, 1)

			gen6990 := Call(__e, PrimNS1Value(symns2_1value), symshen_4abs)

			gen6991 := Call(__e, gen6990, V2286)

			gen6992 := Call(__e, gen6989, gen6991)

			if True == gen6992 {
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

	Call(__e, PrimNS1Value(symns2_1set), syminteger_2, gen6995)

	gen6999 := MakeNative(func(__e *ControlFlow) {
		V2288 := __e.Get(1)
		_ = V2288
		gen6997 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

		gen6998 := Call(__e, gen6997, V2288, MakeNumber(0))

		if True == gen6998 {
			__e.Return(V2288)
			return
		} else {
			gen6996 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

			__e.TailApply(gen6996, MakeNumber(0), V2288)

			return

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4abs, gen6999)

	gen7006 := MakeNative(func(__e *ControlFlow) {
		V2291 := __e.Get(1)
		_ = V2291
		V2292 := __e.Get(2)
		_ = V2292
		gen7003 := MakeNative(func(__e *ControlFlow) {
			Nx2 := __e.Get(1)
			_ = Nx2
			gen7001 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			gen7002 := Call(__e, gen7001, Nx2, V2291)

			if True == gen7002 {
				__e.Return(V2292)
				return
			} else {
				gen7000 := Call(__e, PrimNS1Value(symns2_1value), symshen_4magless)

				__e.TailApply(gen7000, V2291, Nx2)

				return

			}

		}, 1)

		gen7004 := Call(__e, PrimNS1Value(symns2_1value), sym_d)

		gen7005 := Call(__e, gen7004, V2292, MakeNumber(2))

		__e.TailApply(gen7003, gen7005)

		return

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4magless, gen7006)

	gen7019 := MakeNative(func(__e *ControlFlow) {
		V2298 := __e.Get(1)
		_ = V2298
		V2299 := __e.Get(2)
		_ = V2299
		gen7017 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen7018 := Call(__e, gen7017, MakeNumber(0), V2298)

		if True == gen7018 {
			__e.Return(True)
			return
		} else {
			gen7015 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

			gen7016 := Call(__e, gen7015, MakeNumber(1), V2298)

			if True == gen7016 {
				__e.Return(False)
				return
			} else {
				if True == True {
					gen7012 := MakeNative(func(__e *ControlFlow) {
						Abs_1N := __e.Get(1)
						_ = Abs_1N
						gen7010 := Call(__e, PrimNS1Value(symns2_1value), sym_6)

						gen7011 := Call(__e, gen7010, MakeNumber(0), Abs_1N)

						if True == gen7011 {
							gen7009 := Call(__e, PrimNS1Value(symns2_1value), syminteger_2)

							__e.TailApply(gen7009, V2298)

							return

						} else {
							gen7008 := Call(__e, PrimNS1Value(symns2_1value), symshen_4integer_1test_2)

							__e.TailApply(gen7008, Abs_1N, V2299)

							return

						}

					}, 1)

					gen7013 := Call(__e, PrimNS1Value(symns2_1value), sym_1)

					gen7014 := Call(__e, gen7013, V2298, V2299)

					__e.TailApply(gen7012, gen7014)

					return

				} else {
					gen7007 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen7007, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4integer_1test_2, gen7019)

	gen7034 := MakeNative(func(__e *ControlFlow) {
		V2304 := __e.Get(1)
		_ = V2304
		V2305 := __e.Get(2)
		_ = V2305
		gen7032 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen7033 := Call(__e, gen7032, Nil, V2305)

		if True == gen7033 {
			__e.Return(Nil)
			return
		} else {
			gen7030 := Call(__e, PrimNS1Value(symns2_1value), symcons_2)

			gen7031 := Call(__e, gen7030, V2305)

			if True == gen7031 {
				gen7022 := Call(__e, PrimNS1Value(symns2_1value), symappend)

				gen7023 := Call(__e, PrimNS1Value(symns2_1value), symhd)

				gen7024 := Call(__e, gen7023, V2305)

				gen7025 := Call(__e, V2304, gen7024)

				gen7026 := Call(__e, PrimNS1Value(symns2_1value), symmapcan)

				gen7027 := Call(__e, PrimNS1Value(symns2_1value), symtl)

				gen7028 := Call(__e, gen7027, V2305)

				gen7029 := Call(__e, gen7026, V2304, gen7028)

				__e.TailApply(gen7022, gen7025, gen7029)

				return

			} else {
				if True == True {
					gen7021 := Call(__e, PrimNS1Value(symns2_1value), symshen_4f__error)

					__e.TailApply(gen7021, symmapcan)

					return

				} else {
					gen7020 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen7020, MakeString("no cond match"))

					return

				}
			}

		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), symmapcan, gen7034)

	gen7038 := MakeNative(func(__e *ControlFlow) {
		V2317 := __e.Get(1)
		_ = V2317
		V2318 := __e.Get(2)
		_ = V2318
		gen7036 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen7037 := Call(__e, gen7036, V2318, V2317)

		if True == gen7037 {
			__e.Return(True)
			return
		} else {
			if True == True {
				__e.Return(False)
				return
			} else {
				gen7035 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen7035, MakeString("no cond match"))

				return

			}
		}

	}, 2)

	Call(__e, PrimNS1Value(symns2_1set), sym_a_a, gen7038)

	gen7040 := MakeNative(func(__e *ControlFlow) {
		gen7039 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

		__e.TailApply(gen7039, MakeString(""))

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symabort, gen7040)

	gen7051 := MakeNative(func(__e *ControlFlow) {
		V2320 := __e.Get(1)
		_ = V2320
		gen7049 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

		gen7050 := Call(__e, gen7049, V2320)

		if True == gen7050 {
			gen7043 := MakeNative(func(__e *ControlFlow) {
				Val := __e.Get(1)
				_ = Val
				gen7041 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

				gen7042 := Call(__e, gen7041, Val, symshen_4this_1symbol_1is_1unbound)

				if True == gen7042 {
					__e.Return(False)
					return
				} else {
					__e.Return(True)
					return
				}

			}, 1)

			gen7045 := MakeNative(func(__e *ControlFlow) {
				gen7044 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

				__e.TailApply(gen7044, V2320)

				return

			}, 0)

			gen7046 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4this_1symbol_1is_1unbound)
				return
			}, 1)

			gen7047 := Call(__e, PrimNS1Value(symtry_1catch), gen7045, gen7046)

			gen7048 := Call(__e, gen7043, gen7047)

			if True == gen7048 {
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

	Call(__e, PrimNS1Value(symns2_1set), symbound_2, gen7051)

	gen7064 := MakeNative(func(__e *ControlFlow) {
		V2322 := __e.Get(1)
		_ = V2322
		gen7062 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen7063 := Call(__e, gen7062, MakeString(""), V2322)

		if True == gen7063 {
			__e.Return(Nil)
			return
		} else {
			if True == True {
				gen7053 := Call(__e, PrimNS1Value(symns2_1value), symcons)

				gen7054 := Call(__e, PrimNS1Value(symns2_1value), symstring_1_6n)

				gen7055 := Call(__e, PrimNS1Value(symns2_1value), sympos)

				gen7056 := Call(__e, gen7055, V2322, MakeNumber(0))

				gen7057 := Call(__e, gen7054, gen7056)

				gen7058 := Call(__e, PrimNS1Value(symns2_1value), symshen_4string_1_6bytes)

				gen7059 := Call(__e, PrimNS1Value(symns2_1value), symtlstr)

				gen7060 := Call(__e, gen7059, V2322)

				gen7061 := Call(__e, gen7058, gen7060)

				__e.TailApply(gen7053, gen7057, gen7061)

				return

			} else {
				gen7052 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				__e.TailApply(gen7052, MakeString("no cond match"))

				return

			}
		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symshen_4string_1_6bytes, gen7064)

	gen7066 := MakeNative(func(__e *ControlFlow) {
		V2324 := __e.Get(1)
		_ = V2324
		gen7065 := Call(__e, PrimNS1Value(symns2_1value), symset)

		__e.TailApply(gen7065, symshen_4_dmaxinferences_d, V2324)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symmaxinferences, gen7066)

	gen7068 := MakeNative(func(__e *ControlFlow) {
		gen7067 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7067, symshen_4_dinfs_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), syminferences, gen7068)

	gen7069 := MakeNative(func(__e *ControlFlow) {
		V2326 := __e.Get(1)
		_ = V2326
		__e.Return(V2326)
		return
	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symprotect, gen7069)

	gen7071 := MakeNative(func(__e *ControlFlow) {
		gen7070 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7070, sym_dstoutput_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symstoutput, gen7071)

	gen7073 := MakeNative(func(__e *ControlFlow) {
		gen7072 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7072, sym_dsterror_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symsterror, gen7073)

	gen7084 := MakeNative(func(__e *ControlFlow) {
		V2328 := __e.Get(1)
		_ = V2328
		gen7081 := MakeNative(func(__e *ControlFlow) {
			Symbol := __e.Get(1)
			_ = Symbol
			gen7079 := Call(__e, PrimNS1Value(symns2_1value), symsymbol_2)

			gen7080 := Call(__e, gen7079, Symbol)

			if True == gen7080 {
				__e.Return(Symbol)
				return
			} else {
				gen7074 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

				gen7075 := Call(__e, PrimNS1Value(symns2_1value), symcn)

				gen7076 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

				gen7077 := Call(__e, gen7076, V2328, MakeString(" to a symbol"), symshen_4s)

				gen7078 := Call(__e, gen7075, MakeString("cannot intern "), gen7077)

				__e.TailApply(gen7074, gen7078)

				return

			}

		}, 1)

		gen7082 := Call(__e, PrimNS1Value(symns2_1value), symintern)

		gen7083 := Call(__e, gen7082, V2328)

		__e.TailApply(gen7081, gen7083)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symstring_1_6symbol, gen7084)

	gen7093 := MakeNative(func(__e *ControlFlow) {
		V2334 := __e.Get(1)
		_ = V2334
		gen7091 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

		gen7092 := Call(__e, gen7091, sym_7, V2334)

		if True == gen7092 {
			gen7090 := Call(__e, PrimNS1Value(symns2_1value), symset)

			__e.TailApply(gen7090, symshen_4_doptimise_d, True)

			return

		} else {
			gen7088 := Call(__e, PrimNS1Value(symns2_1value), sym_a)

			gen7089 := Call(__e, gen7088, sym_1, V2334)

			if True == gen7089 {
				gen7087 := Call(__e, PrimNS1Value(symns2_1value), symset)

				__e.TailApply(gen7087, symshen_4_doptimise_d, False)

				return

			} else {
				if True == True {
					gen7086 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen7086, MakeString("optimise expects a + or a -.\n"))

					return

				} else {
					gen7085 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

					__e.TailApply(gen7085, MakeString("no cond match"))

					return

				}
			}

		}

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symoptimise, gen7093)

	gen7095 := MakeNative(func(__e *ControlFlow) {
		gen7094 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7094, sym_dos_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symos, gen7095)

	gen7097 := MakeNative(func(__e *ControlFlow) {
		gen7096 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7096, sym_dlanguage_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symlanguage, gen7097)

	gen7099 := MakeNative(func(__e *ControlFlow) {
		gen7098 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7098, sym_dversion_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symversion, gen7099)

	gen7101 := MakeNative(func(__e *ControlFlow) {
		gen7100 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7100, sym_dport_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symport, gen7101)

	gen7103 := MakeNative(func(__e *ControlFlow) {
		gen7102 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7102, sym_dporters_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symporters, gen7103)

	gen7105 := MakeNative(func(__e *ControlFlow) {
		gen7104 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7104, sym_dimplementation_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symimplementation, gen7105)

	gen7107 := MakeNative(func(__e *ControlFlow) {
		gen7106 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

		__e.TailApply(gen7106, sym_drelease_d)

		return

	}, 0)

	Call(__e, PrimNS1Value(symns2_1set), symrelease, gen7107)

	gen7111 := MakeNative(func(__e *ControlFlow) {
		V2336 := __e.Get(1)
		_ = V2336
		gen7109 := MakeNative(func(__e *ControlFlow) {
			gen7108 := Call(__e, PrimNS1Value(symns2_1value), symexternal)

			Call(__e, gen7108, V2336)

			__e.Return(True)
			return

		}, 0)

		gen7110 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(False)
			return
		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen7109, gen7110)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), sympackage_2, gen7111)

	gen7113 := MakeNative(func(__e *ControlFlow) {
		V2338 := __e.Get(1)
		_ = V2338
		gen7112 := Call(__e, PrimNS1Value(symns2_1value), symshen_4lookup_1func)

		__e.TailApply(gen7112, V2338)

		return

	}, 1)

	Call(__e, PrimNS1Value(symns2_1set), symfunction, gen7113)

	gen7122 := MakeNative(func(__e *ControlFlow) {
		V2340 := __e.Get(1)
		_ = V2340
		gen7117 := MakeNative(func(__e *ControlFlow) {
			gen7114 := Call(__e, PrimNS1Value(symns2_1value), symget)

			gen7115 := Call(__e, PrimNS1Value(symns2_1value), symvalue)

			gen7116 := Call(__e, gen7115, sym_dproperty_1vector_d)

			__e.TailApply(gen7114, V2340, symshen_4lambda_1form, gen7116)

			return

		}, 0)

		gen7121 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			gen7118 := Call(__e, PrimNS1Value(symns2_1value), symsimple_1error)

			gen7119 := Call(__e, PrimNS1Value(symns2_1value), symshen_4app)

			gen7120 := Call(__e, gen7119, V2340, MakeString(" has no lambda expansion\n"), symshen_4a)

			__e.TailApply(gen7118, gen7120)

			return

		}, 1)

		__e.TailApply(PrimNS1Value(symtry_1catch), gen7117, gen7121)

		return

	}, 1)

	__e.TailApply(PrimNS1Value(symns2_1set), symshen_4lookup_1func, gen7122)

	return

}, 0)
