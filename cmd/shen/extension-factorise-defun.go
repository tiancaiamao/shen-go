package main

import . "github.com/tiancaiamao/shen-go/kl"

func init() {
	__initExprs = append(__initExprs, MakeNative(func(__e Evaluator) {
		MakeString("Copyright (c) 2012-2019 Bruno Deferrari.  All rights reserved.\nBSD 3-Clause License: http://opensource.org/licenses/BSD-3-Clause")

		gen17712 := MakeNative(func(__e Evaluator) {
			V4931 := __e.Get(1)
			_ = V4931
			gen17710 := Call(__e, __e.Global(symcons_2), V4931)

			var gen17711 Obj
			if True == gen17710 {
				gen17707 := Call(__e, __e.Global(symhd), V4931)

				gen17708 := Call(__e, __e.Global(sym_a), MakeSymbol("defun"), gen17707)

				var gen17709 Obj
				if True == gen17708 {
					gen17704 := Call(__e, __e.Global(symtl), V4931)

					gen17705 := Call(__e, __e.Global(symcons_2), gen17704)

					var gen17706 Obj
					if True == gen17705 {
						gen17700 := Call(__e, __e.Global(symtl), V4931)

						gen17701 := Call(__e, __e.Global(symtl), gen17700)

						gen17702 := Call(__e, __e.Global(symcons_2), gen17701)

						var gen17703 Obj
						if True == gen17702 {
							gen17695 := Call(__e, __e.Global(symtl), V4931)

							gen17696 := Call(__e, __e.Global(symtl), gen17695)

							gen17697 := Call(__e, __e.Global(symtl), gen17696)

							gen17698 := Call(__e, __e.Global(symcons_2), gen17697)

							var gen17699 Obj
							if True == gen17698 {
								gen17689 := Call(__e, __e.Global(symtl), V4931)

								gen17690 := Call(__e, __e.Global(symtl), gen17689)

								gen17691 := Call(__e, __e.Global(symtl), gen17690)

								gen17692 := Call(__e, __e.Global(symhd), gen17691)

								gen17693 := Call(__e, __e.Global(symcons_2), gen17692)

								var gen17694 Obj
								if True == gen17693 {
									gen17682 := Call(__e, __e.Global(symtl), V4931)

									gen17683 := Call(__e, __e.Global(symtl), gen17682)

									gen17684 := Call(__e, __e.Global(symtl), gen17683)

									gen17685 := Call(__e, __e.Global(symhd), gen17684)

									gen17686 := Call(__e, __e.Global(symhd), gen17685)

									gen17687 := Call(__e, __e.Global(sym_a), MakeSymbol("cond"), gen17686)

									var gen17688 Obj
									if True == gen17687 {
										gen17677 := Call(__e, __e.Global(symtl), V4931)

										gen17678 := Call(__e, __e.Global(symtl), gen17677)

										gen17679 := Call(__e, __e.Global(symtl), gen17678)

										gen17680 := Call(__e, __e.Global(symtl), gen17679)

										gen17681 := Call(__e, __e.Global(sym_a), Nil, gen17680)

										if True == gen17681 {
											gen17688 = True
										} else {
											gen17688 = False
										}

									} else {
										gen17688 = False
									}
									if True == gen17688 {
										gen17694 = True
									} else {
										gen17694 = False
									}

								} else {
									gen17694 = False
								}
								if True == gen17694 {
									gen17699 = True
								} else {
									gen17699 = False
								}

							} else {
								gen17699 = False
							}
							if True == gen17699 {
								gen17703 = True
							} else {
								gen17703 = False
							}

						} else {
							gen17703 = False
						}
						if True == gen17703 {
							gen17706 = True
						} else {
							gen17706 = False
						}

					} else {
						gen17706 = False
					}
					if True == gen17706 {
						gen17709 = True
					} else {
						gen17709 = False
					}

				} else {
					gen17709 = False
				}
				if True == gen17709 {
					gen17711 = True
				} else {
					gen17711 = False
				}

			} else {
				gen17711 = False
			}
			if True == gen17711 {
				gen17657 := Call(__e, __e.Global(symtl), V4931)

				gen17658 := Call(__e, __e.Global(symhd), gen17657)

				gen17659 := Call(__e, __e.Global(symtl), V4931)

				gen17660 := Call(__e, __e.Global(symtl), gen17659)

				gen17661 := Call(__e, __e.Global(symhd), gen17660)

				gen17662 := Call(__e, __e.Global(symtl), V4931)

				gen17663 := Call(__e, __e.Global(symtl), gen17662)

				gen17664 := Call(__e, __e.Global(symtl), gen17663)

				gen17665 := Call(__e, __e.Global(symhd), gen17664)

				gen17666 := Call(__e, __e.Global(symtl), V4931)

				gen17667 := Call(__e, __e.Global(symhd), gen17666)

				gen17668 := Call(__e, __e.Global(symcons), gen17667, Nil)

				gen17669 := Call(__e, __e.Global(symcons), MakeSymbol("shen.f_error"), gen17668)

				gen17670 := Call(__e, __e.Global(symtl), V4931)

				gen17671 := Call(__e, __e.Global(symtl), gen17670)

				gen17672 := Call(__e, __e.Global(symhd), gen17671)

				gen17673 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4factorise_1cond), gen17665, gen17669, gen17672)

				gen17674 := Call(__e, __e.Global(symcons), gen17673, Nil)

				gen17675 := Call(__e, __e.Global(symcons), gen17661, gen17674)

				gen17676 := Call(__e, __e.Global(symcons), gen17658, gen17675)

				__e.TailApply(__e.Global(symcons), MakeSymbol("defun"), gen17676)

				return

			} else {
				__e.Return(V4931)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.factorise-defun"), gen17712)

		gen17720 := MakeNative(func(__e Evaluator) {
			V4943 := __e.Get(1)
			_ = V4943
			V4944 := __e.Get(2)
			_ = V4944
			V4945 := __e.Get(3)
			_ = V4945
			gen17718 := Call(__e, __e.Global(symcons_2), V4943)

			var gen17719 Obj
			if True == gen17718 {
				gen17716 := Call(__e, __e.Global(symhd), V4943)

				gen17717 := Call(__e, __e.Global(sym_a), MakeSymbol("cond"), gen17716)

				if True == gen17717 {
					gen17719 = True
				} else {
					gen17719 = False
				}

			} else {
				gen17719 = False
			}
			if True == gen17719 {
				gen17713 := Call(__e, __e.Global(symtl), V4943)

				gen17714 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4add_1returns), gen17713)

				gen17715 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4rebranch), gen17714, V4944)

				__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen17715, V4945)

				return

			} else {
				__e.Return(V4943)
				return
			}

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.factorise-cond"), gen17720)

		gen17744 := MakeNative(func(__e Evaluator) {
			V4947 := __e.Get(1)
			_ = V4947
			gen17743 := Call(__e, __e.Global(sym_a), Nil, V4947)

			if True == gen17743 {
				__e.Return(Nil)
				return
			} else {
				gen17741 := Call(__e, __e.Global(symcons_2), V4947)

				var gen17742 Obj
				if True == gen17741 {
					gen17738 := Call(__e, __e.Global(symhd), V4947)

					gen17739 := Call(__e, __e.Global(symcons_2), gen17738)

					var gen17740 Obj
					if True == gen17739 {
						gen17734 := Call(__e, __e.Global(symhd), V4947)

						gen17735 := Call(__e, __e.Global(symtl), gen17734)

						gen17736 := Call(__e, __e.Global(symcons_2), gen17735)

						var gen17737 Obj
						if True == gen17736 {
							gen17730 := Call(__e, __e.Global(symhd), V4947)

							gen17731 := Call(__e, __e.Global(symtl), gen17730)

							gen17732 := Call(__e, __e.Global(symtl), gen17731)

							gen17733 := Call(__e, __e.Global(sym_a), Nil, gen17732)

							if True == gen17733 {
								gen17737 = True
							} else {
								gen17737 = False
							}

						} else {
							gen17737 = False
						}
						if True == gen17737 {
							gen17740 = True
						} else {
							gen17740 = False
						}

					} else {
						gen17740 = False
					}
					if True == gen17740 {
						gen17742 = True
					} else {
						gen17742 = False
					}

				} else {
					gen17742 = False
				}
				if True == gen17742 {
					gen17721 := Call(__e, __e.Global(symhd), V4947)

					gen17722 := Call(__e, __e.Global(symhd), gen17721)

					gen17723 := Call(__e, __e.Global(symhd), V4947)

					gen17724 := Call(__e, __e.Global(symtl), gen17723)

					gen17725 := Call(__e, __e.Global(symcons), MakeSymbol("%%return"), gen17724)

					gen17726 := Call(__e, __e.Global(symcons), gen17725, Nil)

					gen17727 := Call(__e, __e.Global(symcons), gen17722, gen17726)

					gen17728 := Call(__e, __e.Global(symtl), V4947)

					gen17729 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4add_1returns), gen17728)

					__e.TailApply(__e.Global(symcons), gen17727, gen17729)

					return

				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.add-returns"))

					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.add-returns"), gen17744)

		gen17745 := MakeNative(func(__e Evaluator) {
			__e.TailApply(__e.Global(symgensym), MakeSymbol("%%label"))

			return
		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.generate-label"), gen17745)

		gen17747 := MakeNative(func(__e Evaluator) {
			V4950 := __e.Get(1)
			_ = V4950
			V4951 := __e.Get(2)
			_ = V4951
			gen17746 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4free_1variables_1h), V4950, V4951, Nil)

			__e.TailApply(__e.Global(symreverse), gen17746)

			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.free-variables"), gen17747)

		gen17808 := MakeNative(func(__e Evaluator) {
			V4963 := __e.Get(1)
			_ = V4963
			V4964 := __e.Get(2)
			_ = V4964
			V4965 := __e.Get(3)
			_ = V4965
			gen17806 := Call(__e, __e.Global(symcons_2), V4963)

			var gen17807 Obj
			if True == gen17806 {
				gen17803 := Call(__e, __e.Global(symhd), V4963)

				gen17804 := Call(__e, __e.Global(sym_a), MakeSymbol("let"), gen17803)

				var gen17805 Obj
				if True == gen17804 {
					gen17800 := Call(__e, __e.Global(symtl), V4963)

					gen17801 := Call(__e, __e.Global(symcons_2), gen17800)

					var gen17802 Obj
					if True == gen17801 {
						gen17796 := Call(__e, __e.Global(symtl), V4963)

						gen17797 := Call(__e, __e.Global(symtl), gen17796)

						gen17798 := Call(__e, __e.Global(symcons_2), gen17797)

						var gen17799 Obj
						if True == gen17798 {
							gen17791 := Call(__e, __e.Global(symtl), V4963)

							gen17792 := Call(__e, __e.Global(symtl), gen17791)

							gen17793 := Call(__e, __e.Global(symtl), gen17792)

							gen17794 := Call(__e, __e.Global(symcons_2), gen17793)

							var gen17795 Obj
							if True == gen17794 {
								gen17786 := Call(__e, __e.Global(symtl), V4963)

								gen17787 := Call(__e, __e.Global(symtl), gen17786)

								gen17788 := Call(__e, __e.Global(symtl), gen17787)

								gen17789 := Call(__e, __e.Global(symtl), gen17788)

								gen17790 := Call(__e, __e.Global(sym_a), Nil, gen17789)

								if True == gen17790 {
									gen17795 = True
								} else {
									gen17795 = False
								}

							} else {
								gen17795 = False
							}
							if True == gen17795 {
								gen17799 = True
							} else {
								gen17799 = False
							}

						} else {
							gen17799 = False
						}
						if True == gen17799 {
							gen17802 = True
						} else {
							gen17802 = False
						}

					} else {
						gen17802 = False
					}
					if True == gen17802 {
						gen17805 = True
					} else {
						gen17805 = False
					}

				} else {
					gen17805 = False
				}
				if True == gen17805 {
					gen17807 = True
				} else {
					gen17807 = False
				}

			} else {
				gen17807 = False
			}
			if True == gen17807 {
				gen17775 := Call(__e, __e.Global(symtl), V4963)

				gen17776 := Call(__e, __e.Global(symtl), gen17775)

				gen17777 := Call(__e, __e.Global(symtl), gen17776)

				gen17778 := Call(__e, __e.Global(symhd), gen17777)

				gen17779 := Call(__e, __e.Global(symtl), V4963)

				gen17780 := Call(__e, __e.Global(symhd), gen17779)

				gen17781 := Call(__e, __e.Global(symremove), gen17780, V4964)

				gen17782 := Call(__e, __e.Global(symtl), V4963)

				gen17783 := Call(__e, __e.Global(symtl), gen17782)

				gen17784 := Call(__e, __e.Global(symhd), gen17783)

				gen17785 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4free_1variables_1h), gen17784, V4964, V4965)

				__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4free_1variables_1h), gen17778, gen17781, gen17785)

				return

			} else {
				gen17773 := Call(__e, __e.Global(symcons_2), V4963)

				var gen17774 Obj
				if True == gen17773 {
					gen17770 := Call(__e, __e.Global(symhd), V4963)

					gen17771 := Call(__e, __e.Global(sym_a), MakeSymbol("lambda"), gen17770)

					var gen17772 Obj
					if True == gen17771 {
						gen17767 := Call(__e, __e.Global(symtl), V4963)

						gen17768 := Call(__e, __e.Global(symcons_2), gen17767)

						var gen17769 Obj
						if True == gen17768 {
							gen17763 := Call(__e, __e.Global(symtl), V4963)

							gen17764 := Call(__e, __e.Global(symtl), gen17763)

							gen17765 := Call(__e, __e.Global(symcons_2), gen17764)

							var gen17766 Obj
							if True == gen17765 {
								gen17759 := Call(__e, __e.Global(symtl), V4963)

								gen17760 := Call(__e, __e.Global(symtl), gen17759)

								gen17761 := Call(__e, __e.Global(symtl), gen17760)

								gen17762 := Call(__e, __e.Global(sym_a), Nil, gen17761)

								if True == gen17762 {
									gen17766 = True
								} else {
									gen17766 = False
								}

							} else {
								gen17766 = False
							}
							if True == gen17766 {
								gen17769 = True
							} else {
								gen17769 = False
							}

						} else {
							gen17769 = False
						}
						if True == gen17769 {
							gen17772 = True
						} else {
							gen17772 = False
						}

					} else {
						gen17772 = False
					}
					if True == gen17772 {
						gen17774 = True
					} else {
						gen17774 = False
					}

				} else {
					gen17774 = False
				}
				if True == gen17774 {
					gen17753 := Call(__e, __e.Global(symtl), V4963)

					gen17754 := Call(__e, __e.Global(symtl), gen17753)

					gen17755 := Call(__e, __e.Global(symhd), gen17754)

					gen17756 := Call(__e, __e.Global(symtl), V4963)

					gen17757 := Call(__e, __e.Global(symhd), gen17756)

					gen17758 := Call(__e, __e.Global(symremove), gen17757, V4964)

					__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4free_1variables_1h), gen17755, gen17758, V4965)

					return

				} else {
					gen17752 := Call(__e, __e.Global(symcons_2), V4963)

					if True == gen17752 {
						gen17749 := Call(__e, __e.Global(symtl), V4963)

						gen17750 := Call(__e, __e.Global(symhd), V4963)

						gen17751 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4free_1variables_1h), gen17750, V4964, V4965)

						__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4free_1variables_1h), gen17749, V4964, gen17751)

						return

					} else {
						gen17748 := Call(__e, __e.Global(symelement_2), V4963, V4964)

						if True == gen17748 {
							__e.TailApply(__e.Global(symadjoin), V4963, V4965)

							return
						} else {
							__e.Return(V4965)
							return
						}

					}

				}

			}

		}, 3)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.free-variables-h"), gen17808)

		gen17862 := MakeNative(func(__e Evaluator) {
			V4968 := __e.Get(1)
			_ = V4968
			V4969 := __e.Get(2)
			_ = V4969
			gen17860 := Call(__e, __e.Global(symcons_2), V4968)

			var gen17861 Obj
			if True == gen17860 {
				gen17857 := Call(__e, __e.Global(symhd), V4968)

				gen17858 := Call(__e, __e.Global(sym_a), MakeSymbol("%%let-label"), gen17857)

				var gen17859 Obj
				if True == gen17858 {
					gen17854 := Call(__e, __e.Global(symtl), V4968)

					gen17855 := Call(__e, __e.Global(symcons_2), gen17854)

					var gen17856 Obj
					if True == gen17855 {
						gen17850 := Call(__e, __e.Global(symtl), V4968)

						gen17851 := Call(__e, __e.Global(symtl), gen17850)

						gen17852 := Call(__e, __e.Global(symcons_2), gen17851)

						var gen17853 Obj
						if True == gen17852 {
							gen17845 := Call(__e, __e.Global(symtl), V4968)

							gen17846 := Call(__e, __e.Global(symtl), gen17845)

							gen17847 := Call(__e, __e.Global(symtl), gen17846)

							gen17848 := Call(__e, __e.Global(symcons_2), gen17847)

							var gen17849 Obj
							if True == gen17848 {
								gen17840 := Call(__e, __e.Global(symtl), V4968)

								gen17841 := Call(__e, __e.Global(symtl), gen17840)

								gen17842 := Call(__e, __e.Global(symtl), gen17841)

								gen17843 := Call(__e, __e.Global(symtl), gen17842)

								gen17844 := Call(__e, __e.Global(sym_a), Nil, gen17843)

								if True == gen17844 {
									gen17849 = True
								} else {
									gen17849 = False
								}

							} else {
								gen17849 = False
							}
							if True == gen17849 {
								gen17853 = True
							} else {
								gen17853 = False
							}

						} else {
							gen17853 = False
						}
						if True == gen17853 {
							gen17856 = True
						} else {
							gen17856 = False
						}

					} else {
						gen17856 = False
					}
					if True == gen17856 {
						gen17859 = True
					} else {
						gen17859 = False
					}

				} else {
					gen17859 = False
				}
				if True == gen17859 {
					gen17861 = True
				} else {
					gen17861 = False
				}

			} else {
				gen17861 = False
			}
			if True == gen17861 {
				gen17809 := Call(__e, __e.Global(symtl), V4968)

				gen17810 := Call(__e, __e.Global(symtl), gen17809)

				gen17811 := Call(__e, __e.Global(symhd), gen17810)

				gen17812 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4free_1variables), gen17811, V4969)

				FreeVars := gen17812
				gen17828 := Call(__e, __e.Global(sym_a), Nil, FreeVars)

				var gen17829 Obj
				if True == gen17828 {
					gen17825 := Call(__e, __e.Global(symtl), V4968)

					gen17826 := Call(__e, __e.Global(symtl), gen17825)

					gen17827 := Call(__e, __e.Global(symtl), gen17826)

					gen17829 = Call(__e, __e.Global(symhd), gen17827)

				} else {
					gen17813 := Call(__e, __e.Global(symtl), V4968)

					gen17814 := Call(__e, __e.Global(symhd), gen17813)

					gen17815 := Call(__e, __e.Global(symcons), gen17814, FreeVars)

					gen17816 := Call(__e, __e.Global(symcons), MakeSymbol("%%goto-label"), gen17815)

					gen17817 := Call(__e, __e.Global(symtl), V4968)

					gen17818 := Call(__e, __e.Global(symhd), gen17817)

					gen17819 := Call(__e, __e.Global(symcons), gen17818, Nil)

					gen17820 := Call(__e, __e.Global(symcons), MakeSymbol("%%goto-label"), gen17819)

					gen17821 := Call(__e, __e.Global(symtl), V4968)

					gen17822 := Call(__e, __e.Global(symtl), gen17821)

					gen17823 := Call(__e, __e.Global(symtl), gen17822)

					gen17824 := Call(__e, __e.Global(symhd), gen17823)

					gen17829 = Call(__e, __e.Global(symsubst), gen17816, gen17820, gen17824)

				}
				NewBody := gen17829
				gen17830 := Call(__e, __e.Global(symtl), V4968)

				gen17831 := Call(__e, __e.Global(symhd), gen17830)

				gen17832 := Call(__e, __e.Global(symcons), gen17831, FreeVars)

				gen17833 := Call(__e, __e.Global(symtl), V4968)

				gen17834 := Call(__e, __e.Global(symtl), gen17833)

				gen17835 := Call(__e, __e.Global(symhd), gen17834)

				gen17836 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), NewBody, V4969)

				gen17837 := Call(__e, __e.Global(symcons), gen17836, Nil)

				gen17838 := Call(__e, __e.Global(symcons), gen17835, gen17837)

				gen17839 := Call(__e, __e.Global(symcons), gen17832, gen17838)

				__e.TailApply(__e.Global(symcons), MakeSymbol("%%let-label"), gen17839)

				return

			} else {
				__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.attach-free-variables"))

				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.attach-free-variables"), gen17862)

		gen18017 := MakeNative(func(__e Evaluator) {
			V4976 := __e.Get(1)
			_ = V4976
			V4977 := __e.Get(2)
			_ = V4977
			gen18015 := Call(__e, __e.Global(symcons_2), V4976)

			var gen18016 Obj
			if True == gen18015 {
				gen18012 := Call(__e, __e.Global(symhd), V4976)

				gen18013 := Call(__e, __e.Global(sym_a), MakeSymbol("%%let-label"), gen18012)

				var gen18014 Obj
				if True == gen18013 {
					gen18009 := Call(__e, __e.Global(symtl), V4976)

					gen18010 := Call(__e, __e.Global(symcons_2), gen18009)

					var gen18011 Obj
					if True == gen18010 {
						gen18005 := Call(__e, __e.Global(symtl), V4976)

						gen18006 := Call(__e, __e.Global(symtl), gen18005)

						gen18007 := Call(__e, __e.Global(symcons_2), gen18006)

						var gen18008 Obj
						if True == gen18007 {
							gen18000 := Call(__e, __e.Global(symtl), V4976)

							gen18001 := Call(__e, __e.Global(symtl), gen18000)

							gen18002 := Call(__e, __e.Global(symtl), gen18001)

							gen18003 := Call(__e, __e.Global(symcons_2), gen18002)

							var gen18004 Obj
							if True == gen18003 {
								gen17994 := Call(__e, __e.Global(symtl), V4976)

								gen17995 := Call(__e, __e.Global(symtl), gen17994)

								gen17996 := Call(__e, __e.Global(symtl), gen17995)

								gen17997 := Call(__e, __e.Global(symtl), gen17996)

								gen17998 := Call(__e, __e.Global(sym_a), Nil, gen17997)

								var gen17999 Obj
								if True == gen17998 {
									gen17984 := Call(__e, __e.Global(symtl), V4976)

									gen17985 := Call(__e, __e.Global(symhd), gen17984)

									gen17986 := Call(__e, __e.Global(symcons), gen17985, Nil)

									gen17987 := Call(__e, __e.Global(symcons), MakeSymbol("%%goto-label"), gen17986)

									gen17988 := Call(__e, __e.Global(symtl), V4976)

									gen17989 := Call(__e, __e.Global(symtl), gen17988)

									gen17990 := Call(__e, __e.Global(symtl), gen17989)

									gen17991 := Call(__e, __e.Global(symhd), gen17990)

									gen17992 := Call(__e, __e.Global(symoccurrences), gen17987, gen17991)

									gen17993 := Call(__e, __e.Global(sym_6), gen17992, MakeNumber(1))

									if True == gen17993 {
										gen17999 = True
									} else {
										gen17999 = False
									}

								} else {
									gen17999 = False
								}
								if True == gen17999 {
									gen18004 = True
								} else {
									gen18004 = False
								}

							} else {
								gen18004 = False
							}
							if True == gen18004 {
								gen18008 = True
							} else {
								gen18008 = False
							}

						} else {
							gen18008 = False
						}
						if True == gen18008 {
							gen18011 = True
						} else {
							gen18011 = False
						}

					} else {
						gen18011 = False
					}
					if True == gen18011 {
						gen18014 = True
					} else {
						gen18014 = False
					}

				} else {
					gen18014 = False
				}
				if True == gen18014 {
					gen18016 = True
				} else {
					gen18016 = False
				}

			} else {
				gen18016 = False
			}
			if True == gen18016 {
				gen17972 := Call(__e, __e.Global(symtl), V4976)

				gen17973 := Call(__e, __e.Global(symhd), gen17972)

				gen17974 := Call(__e, __e.Global(symtl), V4976)

				gen17975 := Call(__e, __e.Global(symtl), gen17974)

				gen17976 := Call(__e, __e.Global(symhd), gen17975)

				gen17977 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen17976, V4977)

				gen17978 := Call(__e, __e.Global(symtl), V4976)

				gen17979 := Call(__e, __e.Global(symtl), gen17978)

				gen17980 := Call(__e, __e.Global(symtl), gen17979)

				gen17981 := Call(__e, __e.Global(symcons), gen17977, gen17980)

				gen17982 := Call(__e, __e.Global(symcons), gen17973, gen17981)

				gen17983 := Call(__e, __e.Global(symcons), MakeSymbol("%%let-label"), gen17982)

				__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4attach_1free_1variables), gen17983, V4977)

				return

			} else {
				gen17970 := Call(__e, __e.Global(symcons_2), V4976)

				var gen17971 Obj
				if True == gen17970 {
					gen17967 := Call(__e, __e.Global(symhd), V4976)

					gen17968 := Call(__e, __e.Global(sym_a), MakeSymbol("%%let-label"), gen17967)

					var gen17969 Obj
					if True == gen17968 {
						gen17964 := Call(__e, __e.Global(symtl), V4976)

						gen17965 := Call(__e, __e.Global(symcons_2), gen17964)

						var gen17966 Obj
						if True == gen17965 {
							gen17960 := Call(__e, __e.Global(symtl), V4976)

							gen17961 := Call(__e, __e.Global(symtl), gen17960)

							gen17962 := Call(__e, __e.Global(symcons_2), gen17961)

							var gen17963 Obj
							if True == gen17962 {
								gen17955 := Call(__e, __e.Global(symtl), V4976)

								gen17956 := Call(__e, __e.Global(symtl), gen17955)

								gen17957 := Call(__e, __e.Global(symtl), gen17956)

								gen17958 := Call(__e, __e.Global(symcons_2), gen17957)

								var gen17959 Obj
								if True == gen17958 {
									gen17950 := Call(__e, __e.Global(symtl), V4976)

									gen17951 := Call(__e, __e.Global(symtl), gen17950)

									gen17952 := Call(__e, __e.Global(symtl), gen17951)

									gen17953 := Call(__e, __e.Global(symtl), gen17952)

									gen17954 := Call(__e, __e.Global(sym_a), Nil, gen17953)

									if True == gen17954 {
										gen17959 = True
									} else {
										gen17959 = False
									}

								} else {
									gen17959 = False
								}
								if True == gen17959 {
									gen17963 = True
								} else {
									gen17963 = False
								}

							} else {
								gen17963 = False
							}
							if True == gen17963 {
								gen17966 = True
							} else {
								gen17966 = False
							}

						} else {
							gen17966 = False
						}
						if True == gen17966 {
							gen17969 = True
						} else {
							gen17969 = False
						}

					} else {
						gen17969 = False
					}
					if True == gen17969 {
						gen17971 = True
					} else {
						gen17971 = False
					}

				} else {
					gen17971 = False
				}
				if True == gen17971 {
					gen17937 := Call(__e, __e.Global(symtl), V4976)

					gen17938 := Call(__e, __e.Global(symtl), gen17937)

					gen17939 := Call(__e, __e.Global(symhd), gen17938)

					gen17940 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen17939, V4977)

					gen17941 := Call(__e, __e.Global(symtl), V4976)

					gen17942 := Call(__e, __e.Global(symhd), gen17941)

					gen17943 := Call(__e, __e.Global(symcons), gen17942, Nil)

					gen17944 := Call(__e, __e.Global(symcons), MakeSymbol("%%goto-label"), gen17943)

					gen17945 := Call(__e, __e.Global(symtl), V4976)

					gen17946 := Call(__e, __e.Global(symtl), gen17945)

					gen17947 := Call(__e, __e.Global(symtl), gen17946)

					gen17948 := Call(__e, __e.Global(symhd), gen17947)

					gen17949 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen17948, V4977)

					__e.TailApply(__e.Global(symsubst), gen17940, gen17944, gen17949)

					return

				} else {
					gen17935 := Call(__e, __e.Global(symcons_2), V4976)

					var gen17936 Obj
					if True == gen17935 {
						gen17932 := Call(__e, __e.Global(symhd), V4976)

						gen17933 := Call(__e, __e.Global(sym_a), MakeSymbol("if"), gen17932)

						var gen17934 Obj
						if True == gen17933 {
							gen17929 := Call(__e, __e.Global(symtl), V4976)

							gen17930 := Call(__e, __e.Global(symcons_2), gen17929)

							var gen17931 Obj
							if True == gen17930 {
								gen17925 := Call(__e, __e.Global(symtl), V4976)

								gen17926 := Call(__e, __e.Global(symtl), gen17925)

								gen17927 := Call(__e, __e.Global(symcons_2), gen17926)

								var gen17928 Obj
								if True == gen17927 {
									gen17920 := Call(__e, __e.Global(symtl), V4976)

									gen17921 := Call(__e, __e.Global(symtl), gen17920)

									gen17922 := Call(__e, __e.Global(symtl), gen17921)

									gen17923 := Call(__e, __e.Global(symcons_2), gen17922)

									var gen17924 Obj
									if True == gen17923 {
										gen17915 := Call(__e, __e.Global(symtl), V4976)

										gen17916 := Call(__e, __e.Global(symtl), gen17915)

										gen17917 := Call(__e, __e.Global(symtl), gen17916)

										gen17918 := Call(__e, __e.Global(symtl), gen17917)

										gen17919 := Call(__e, __e.Global(sym_a), Nil, gen17918)

										if True == gen17919 {
											gen17924 = True
										} else {
											gen17924 = False
										}

									} else {
										gen17924 = False
									}
									if True == gen17924 {
										gen17928 = True
									} else {
										gen17928 = False
									}

								} else {
									gen17928 = False
								}
								if True == gen17928 {
									gen17931 = True
								} else {
									gen17931 = False
								}

							} else {
								gen17931 = False
							}
							if True == gen17931 {
								gen17934 = True
							} else {
								gen17934 = False
							}

						} else {
							gen17934 = False
						}
						if True == gen17934 {
							gen17936 = True
						} else {
							gen17936 = False
						}

					} else {
						gen17936 = False
					}
					if True == gen17936 {
						gen17901 := Call(__e, __e.Global(symtl), V4976)

						gen17902 := Call(__e, __e.Global(symhd), gen17901)

						gen17903 := Call(__e, __e.Global(symtl), V4976)

						gen17904 := Call(__e, __e.Global(symtl), gen17903)

						gen17905 := Call(__e, __e.Global(symhd), gen17904)

						gen17906 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen17905, V4977)

						gen17907 := Call(__e, __e.Global(symtl), V4976)

						gen17908 := Call(__e, __e.Global(symtl), gen17907)

						gen17909 := Call(__e, __e.Global(symtl), gen17908)

						gen17910 := Call(__e, __e.Global(symhd), gen17909)

						gen17911 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen17910, V4977)

						gen17912 := Call(__e, __e.Global(symcons), gen17911, Nil)

						gen17913 := Call(__e, __e.Global(symcons), gen17906, gen17912)

						gen17914 := Call(__e, __e.Global(symcons), gen17902, gen17913)

						__e.TailApply(__e.Global(symcons), MakeSymbol("if"), gen17914)

						return

					} else {
						gen17899 := Call(__e, __e.Global(symcons_2), V4976)

						var gen17900 Obj
						if True == gen17899 {
							gen17896 := Call(__e, __e.Global(symhd), V4976)

							gen17897 := Call(__e, __e.Global(sym_a), MakeSymbol("let"), gen17896)

							var gen17898 Obj
							if True == gen17897 {
								gen17893 := Call(__e, __e.Global(symtl), V4976)

								gen17894 := Call(__e, __e.Global(symcons_2), gen17893)

								var gen17895 Obj
								if True == gen17894 {
									gen17889 := Call(__e, __e.Global(symtl), V4976)

									gen17890 := Call(__e, __e.Global(symtl), gen17889)

									gen17891 := Call(__e, __e.Global(symcons_2), gen17890)

									var gen17892 Obj
									if True == gen17891 {
										gen17884 := Call(__e, __e.Global(symtl), V4976)

										gen17885 := Call(__e, __e.Global(symtl), gen17884)

										gen17886 := Call(__e, __e.Global(symtl), gen17885)

										gen17887 := Call(__e, __e.Global(symcons_2), gen17886)

										var gen17888 Obj
										if True == gen17887 {
											gen17879 := Call(__e, __e.Global(symtl), V4976)

											gen17880 := Call(__e, __e.Global(symtl), gen17879)

											gen17881 := Call(__e, __e.Global(symtl), gen17880)

											gen17882 := Call(__e, __e.Global(symtl), gen17881)

											gen17883 := Call(__e, __e.Global(sym_a), Nil, gen17882)

											if True == gen17883 {
												gen17888 = True
											} else {
												gen17888 = False
											}

										} else {
											gen17888 = False
										}
										if True == gen17888 {
											gen17892 = True
										} else {
											gen17892 = False
										}

									} else {
										gen17892 = False
									}
									if True == gen17892 {
										gen17895 = True
									} else {
										gen17895 = False
									}

								} else {
									gen17895 = False
								}
								if True == gen17895 {
									gen17898 = True
								} else {
									gen17898 = False
								}

							} else {
								gen17898 = False
							}
							if True == gen17898 {
								gen17900 = True
							} else {
								gen17900 = False
							}

						} else {
							gen17900 = False
						}
						if True == gen17900 {
							gen17863 := Call(__e, __e.Global(symtl), V4976)

							gen17864 := Call(__e, __e.Global(symhd), gen17863)

							gen17865 := Call(__e, __e.Global(symtl), V4976)

							gen17866 := Call(__e, __e.Global(symtl), gen17865)

							gen17867 := Call(__e, __e.Global(symhd), gen17866)

							gen17868 := Call(__e, __e.Global(symtl), V4976)

							gen17869 := Call(__e, __e.Global(symtl), gen17868)

							gen17870 := Call(__e, __e.Global(symtl), gen17869)

							gen17871 := Call(__e, __e.Global(symhd), gen17870)

							gen17872 := Call(__e, __e.Global(symtl), V4976)

							gen17873 := Call(__e, __e.Global(symhd), gen17872)

							gen17874 := Call(__e, __e.Global(symcons), gen17873, V4977)

							gen17875 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4inline_1mono_1labels), gen17871, gen17874)

							gen17876 := Call(__e, __e.Global(symcons), gen17875, Nil)

							gen17877 := Call(__e, __e.Global(symcons), gen17867, gen17876)

							gen17878 := Call(__e, __e.Global(symcons), gen17864, gen17877)

							__e.TailApply(__e.Global(symcons), MakeSymbol("let"), gen17878)

							return

						} else {
							__e.Return(V4976)
							return
						}

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.inline-mono-labels"), gen18017)

		gen18113 := MakeNative(func(__e Evaluator) {
			V4984 := __e.Get(1)
			_ = V4984
			V4985 := __e.Get(2)
			_ = V4985
			gen18112 := Call(__e, __e.Global(sym_a), Nil, V4984)

			if True == gen18112 {
				__e.Return(V4985)
				return
			} else {
				gen18110 := Call(__e, __e.Global(symcons_2), V4984)

				var gen18111 Obj
				if True == gen18110 {
					gen18107 := Call(__e, __e.Global(symhd), V4984)

					gen18108 := Call(__e, __e.Global(symcons_2), gen18107)

					var gen18109 Obj
					if True == gen18108 {
						gen18103 := Call(__e, __e.Global(symhd), V4984)

						gen18104 := Call(__e, __e.Global(symhd), gen18103)

						gen18105 := Call(__e, __e.Global(sym_a), True, gen18104)

						var gen18106 Obj
						if True == gen18105 {
							gen18099 := Call(__e, __e.Global(symhd), V4984)

							gen18100 := Call(__e, __e.Global(symtl), gen18099)

							gen18101 := Call(__e, __e.Global(symcons_2), gen18100)

							var gen18102 Obj
							if True == gen18101 {
								gen18095 := Call(__e, __e.Global(symhd), V4984)

								gen18096 := Call(__e, __e.Global(symtl), gen18095)

								gen18097 := Call(__e, __e.Global(symtl), gen18096)

								gen18098 := Call(__e, __e.Global(sym_a), Nil, gen18097)

								if True == gen18098 {
									gen18102 = True
								} else {
									gen18102 = False
								}

							} else {
								gen18102 = False
							}
							if True == gen18102 {
								gen18106 = True
							} else {
								gen18106 = False
							}

						} else {
							gen18106 = False
						}
						if True == gen18106 {
							gen18109 = True
						} else {
							gen18109 = False
						}

					} else {
						gen18109 = False
					}
					if True == gen18109 {
						gen18111 = True
					} else {
						gen18111 = False
					}

				} else {
					gen18111 = False
				}
				if True == gen18111 {
					gen18093 := Call(__e, __e.Global(symhd), V4984)

					gen18094 := Call(__e, __e.Global(symtl), gen18093)

					__e.TailApply(__e.Global(symhd), gen18094)

					return

				} else {
					gen18091 := Call(__e, __e.Global(symcons_2), V4984)

					var gen18092 Obj
					if True == gen18091 {
						gen18088 := Call(__e, __e.Global(symhd), V4984)

						gen18089 := Call(__e, __e.Global(symcons_2), gen18088)

						var gen18090 Obj
						if True == gen18089 {
							gen18084 := Call(__e, __e.Global(symhd), V4984)

							gen18085 := Call(__e, __e.Global(symhd), gen18084)

							gen18086 := Call(__e, __e.Global(symcons_2), gen18085)

							var gen18087 Obj
							if True == gen18086 {
								gen18079 := Call(__e, __e.Global(symhd), V4984)

								gen18080 := Call(__e, __e.Global(symhd), gen18079)

								gen18081 := Call(__e, __e.Global(symhd), gen18080)

								gen18082 := Call(__e, __e.Global(sym_a), MakeSymbol("and"), gen18081)

								var gen18083 Obj
								if True == gen18082 {
									gen18074 := Call(__e, __e.Global(symhd), V4984)

									gen18075 := Call(__e, __e.Global(symhd), gen18074)

									gen18076 := Call(__e, __e.Global(symtl), gen18075)

									gen18077 := Call(__e, __e.Global(symcons_2), gen18076)

									var gen18078 Obj
									if True == gen18077 {
										gen18068 := Call(__e, __e.Global(symhd), V4984)

										gen18069 := Call(__e, __e.Global(symhd), gen18068)

										gen18070 := Call(__e, __e.Global(symtl), gen18069)

										gen18071 := Call(__e, __e.Global(symtl), gen18070)

										gen18072 := Call(__e, __e.Global(symcons_2), gen18071)

										var gen18073 Obj
										if True == gen18072 {
											gen18061 := Call(__e, __e.Global(symhd), V4984)

											gen18062 := Call(__e, __e.Global(symhd), gen18061)

											gen18063 := Call(__e, __e.Global(symtl), gen18062)

											gen18064 := Call(__e, __e.Global(symtl), gen18063)

											gen18065 := Call(__e, __e.Global(symtl), gen18064)

											gen18066 := Call(__e, __e.Global(sym_a), Nil, gen18065)

											var gen18067 Obj
											if True == gen18066 {
												gen18057 := Call(__e, __e.Global(symhd), V4984)

												gen18058 := Call(__e, __e.Global(symtl), gen18057)

												gen18059 := Call(__e, __e.Global(symcons_2), gen18058)

												var gen18060 Obj
												if True == gen18059 {
													gen18053 := Call(__e, __e.Global(symhd), V4984)

													gen18054 := Call(__e, __e.Global(symtl), gen18053)

													gen18055 := Call(__e, __e.Global(symtl), gen18054)

													gen18056 := Call(__e, __e.Global(sym_a), Nil, gen18055)

													if True == gen18056 {
														gen18060 = True
													} else {
														gen18060 = False
													}

												} else {
													gen18060 = False
												}
												if True == gen18060 {
													gen18067 = True
												} else {
													gen18067 = False
												}

											} else {
												gen18067 = False
											}
											if True == gen18067 {
												gen18073 = True
											} else {
												gen18073 = False
											}

										} else {
											gen18073 = False
										}
										if True == gen18073 {
											gen18078 = True
										} else {
											gen18078 = False
										}

									} else {
										gen18078 = False
									}
									if True == gen18078 {
										gen18083 = True
									} else {
										gen18083 = False
									}

								} else {
									gen18083 = False
								}
								if True == gen18083 {
									gen18087 = True
								} else {
									gen18087 = False
								}

							} else {
								gen18087 = False
							}
							if True == gen18087 {
								gen18090 = True
							} else {
								gen18090 = False
							}

						} else {
							gen18090 = False
						}
						if True == gen18090 {
							gen18092 = True
						} else {
							gen18092 = False
						}

					} else {
						gen18092 = False
					}
					if True == gen18092 {
						gen18039 := Call(__e, __e.Global(symhd), V4984)

						gen18040 := Call(__e, __e.Global(symhd), gen18039)

						gen18041 := Call(__e, __e.Global(symtl), gen18040)

						gen18042 := Call(__e, __e.Global(symhd), gen18041)

						gen18043 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4true_1branch), gen18042, V4984)

						TrueBranch := gen18043
						gen18044 := Call(__e, __e.Global(symhd), V4984)

						gen18045 := Call(__e, __e.Global(symhd), gen18044)

						gen18046 := Call(__e, __e.Global(symtl), gen18045)

						gen18047 := Call(__e, __e.Global(symhd), gen18046)

						gen18048 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4false_1branch), gen18047, V4984)

						FalseBranch := gen18048
						gen18049 := Call(__e, __e.Global(symhd), V4984)

						gen18050 := Call(__e, __e.Global(symhd), gen18049)

						gen18051 := Call(__e, __e.Global(symtl), gen18050)

						gen18052 := Call(__e, __e.Global(symhd), gen18051)

						__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4rebranch_1h), gen18052, TrueBranch, FalseBranch, V4985)

						return

					} else {
						gen18037 := Call(__e, __e.Global(symcons_2), V4984)

						var gen18038 Obj
						if True == gen18037 {
							gen18034 := Call(__e, __e.Global(symhd), V4984)

							gen18035 := Call(__e, __e.Global(symcons_2), gen18034)

							var gen18036 Obj
							if True == gen18035 {
								gen18030 := Call(__e, __e.Global(symhd), V4984)

								gen18031 := Call(__e, __e.Global(symtl), gen18030)

								gen18032 := Call(__e, __e.Global(symcons_2), gen18031)

								var gen18033 Obj
								if True == gen18032 {
									gen18026 := Call(__e, __e.Global(symhd), V4984)

									gen18027 := Call(__e, __e.Global(symtl), gen18026)

									gen18028 := Call(__e, __e.Global(symtl), gen18027)

									gen18029 := Call(__e, __e.Global(sym_a), Nil, gen18028)

									if True == gen18029 {
										gen18033 = True
									} else {
										gen18033 = False
									}

								} else {
									gen18033 = False
								}
								if True == gen18033 {
									gen18036 = True
								} else {
									gen18036 = False
								}

							} else {
								gen18036 = False
							}
							if True == gen18036 {
								gen18038 = True
							} else {
								gen18038 = False
							}

						} else {
							gen18038 = False
						}
						if True == gen18038 {
							gen18018 := Call(__e, __e.Global(symhd), V4984)

							gen18019 := Call(__e, __e.Global(symhd), gen18018)

							gen18020 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4true_1branch), gen18019, V4984)

							TrueBranch := gen18020
							gen18021 := Call(__e, __e.Global(symhd), V4984)

							gen18022 := Call(__e, __e.Global(symhd), gen18021)

							gen18023 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4false_1branch), gen18022, V4984)

							FalseBranch := gen18023
							gen18024 := Call(__e, __e.Global(symhd), V4984)

							gen18025 := Call(__e, __e.Global(symhd), gen18024)

							__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4rebranch_1h), gen18025, TrueBranch, FalseBranch, V4985)

							return

						} else {
							__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.rebranch"))

							return
						}

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.rebranch"), gen18113)

		gen18122 := MakeNative(func(__e Evaluator) {
			V4990 := __e.Get(1)
			_ = V4990
			V4991 := __e.Get(2)
			_ = V4991
			V4992 := __e.Get(3)
			_ = V4992
			V4993 := __e.Get(4)
			_ = V4993
			gen18114 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4rebranch), V4992, V4993)

			NewElse := gen18114
			gen18121 := MakeNative(func(__e Evaluator) {
				GotoElse := __e.Get(1)
				_ = GotoElse
				gen18115 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4rebranch), V4991, GotoElse)

				gen18116 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4optimize_1selectors), V4990, gen18115)

				gen18117 := Call(__e, __e.Global(symcons), GotoElse, Nil)

				gen18118 := Call(__e, __e.Global(symcons), gen18116, gen18117)

				gen18119 := Call(__e, __e.Global(symcons), V4990, gen18118)

				gen18120 := Call(__e, __e.Global(symcons), MakeSymbol("if"), gen18119)

				__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs), gen18120)

				return

			}, 1)
			__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4with_1labelled_1else), NewElse, gen18121)

			return

		}, 4)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.rebranch-h"), gen18122)

		gen18203 := MakeNative(func(__e Evaluator) {
			V5006 := __e.Get(1)
			_ = V5006
			V5007 := __e.Get(2)
			_ = V5007
			gen18201 := Call(__e, __e.Global(symcons_2), V5007)

			var gen18202 Obj
			if True == gen18201 {
				gen18198 := Call(__e, __e.Global(symhd), V5007)

				gen18199 := Call(__e, __e.Global(symcons_2), gen18198)

				var gen18200 Obj
				if True == gen18199 {
					gen18194 := Call(__e, __e.Global(symhd), V5007)

					gen18195 := Call(__e, __e.Global(symhd), gen18194)

					gen18196 := Call(__e, __e.Global(symcons_2), gen18195)

					var gen18197 Obj
					if True == gen18196 {
						gen18189 := Call(__e, __e.Global(symhd), V5007)

						gen18190 := Call(__e, __e.Global(symhd), gen18189)

						gen18191 := Call(__e, __e.Global(symhd), gen18190)

						gen18192 := Call(__e, __e.Global(sym_a), MakeSymbol("and"), gen18191)

						var gen18193 Obj
						if True == gen18192 {
							gen18184 := Call(__e, __e.Global(symhd), V5007)

							gen18185 := Call(__e, __e.Global(symhd), gen18184)

							gen18186 := Call(__e, __e.Global(symtl), gen18185)

							gen18187 := Call(__e, __e.Global(symcons_2), gen18186)

							var gen18188 Obj
							if True == gen18187 {
								gen18178 := Call(__e, __e.Global(symhd), V5007)

								gen18179 := Call(__e, __e.Global(symhd), gen18178)

								gen18180 := Call(__e, __e.Global(symtl), gen18179)

								gen18181 := Call(__e, __e.Global(symtl), gen18180)

								gen18182 := Call(__e, __e.Global(symcons_2), gen18181)

								var gen18183 Obj
								if True == gen18182 {
									gen18171 := Call(__e, __e.Global(symhd), V5007)

									gen18172 := Call(__e, __e.Global(symhd), gen18171)

									gen18173 := Call(__e, __e.Global(symtl), gen18172)

									gen18174 := Call(__e, __e.Global(symtl), gen18173)

									gen18175 := Call(__e, __e.Global(symtl), gen18174)

									gen18176 := Call(__e, __e.Global(sym_a), Nil, gen18175)

									var gen18177 Obj
									if True == gen18176 {
										gen18167 := Call(__e, __e.Global(symhd), V5007)

										gen18168 := Call(__e, __e.Global(symtl), gen18167)

										gen18169 := Call(__e, __e.Global(symcons_2), gen18168)

										var gen18170 Obj
										if True == gen18169 {
											gen18162 := Call(__e, __e.Global(symhd), V5007)

											gen18163 := Call(__e, __e.Global(symtl), gen18162)

											gen18164 := Call(__e, __e.Global(symtl), gen18163)

											gen18165 := Call(__e, __e.Global(sym_a), Nil, gen18164)

											var gen18166 Obj
											if True == gen18165 {
												gen18157 := Call(__e, __e.Global(symhd), V5007)

												gen18158 := Call(__e, __e.Global(symhd), gen18157)

												gen18159 := Call(__e, __e.Global(symtl), gen18158)

												gen18160 := Call(__e, __e.Global(symhd), gen18159)

												gen18161 := Call(__e, __e.Global(sym_a), gen18160, V5006)

												if True == gen18161 {
													gen18166 = True
												} else {
													gen18166 = False
												}

											} else {
												gen18166 = False
											}
											if True == gen18166 {
												gen18170 = True
											} else {
												gen18170 = False
											}

										} else {
											gen18170 = False
										}
										if True == gen18170 {
											gen18177 = True
										} else {
											gen18177 = False
										}

									} else {
										gen18177 = False
									}
									if True == gen18177 {
										gen18183 = True
									} else {
										gen18183 = False
									}

								} else {
									gen18183 = False
								}
								if True == gen18183 {
									gen18188 = True
								} else {
									gen18188 = False
								}

							} else {
								gen18188 = False
							}
							if True == gen18188 {
								gen18193 = True
							} else {
								gen18193 = False
							}

						} else {
							gen18193 = False
						}
						if True == gen18193 {
							gen18197 = True
						} else {
							gen18197 = False
						}

					} else {
						gen18197 = False
					}
					if True == gen18197 {
						gen18200 = True
					} else {
						gen18200 = False
					}

				} else {
					gen18200 = False
				}
				if True == gen18200 {
					gen18202 = True
				} else {
					gen18202 = False
				}

			} else {
				gen18202 = False
			}
			if True == gen18202 {
				gen18143 := Call(__e, __e.Global(symhd), V5007)

				gen18144 := Call(__e, __e.Global(symhd), gen18143)

				gen18145 := Call(__e, __e.Global(symtl), gen18144)

				gen18146 := Call(__e, __e.Global(symtl), gen18145)

				gen18147 := Call(__e, __e.Global(symhd), gen18146)

				gen18148 := Call(__e, __e.Global(symhd), V5007)

				gen18149 := Call(__e, __e.Global(symtl), gen18148)

				gen18150 := Call(__e, __e.Global(symcons), gen18147, gen18149)

				gen18151 := Call(__e, __e.Global(symhd), V5007)

				gen18152 := Call(__e, __e.Global(symhd), gen18151)

				gen18153 := Call(__e, __e.Global(symtl), gen18152)

				gen18154 := Call(__e, __e.Global(symhd), gen18153)

				gen18155 := Call(__e, __e.Global(symtl), V5007)

				gen18156 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4true_1branch), gen18154, gen18155)

				__e.TailApply(__e.Global(symcons), gen18150, gen18156)

				return

			} else {
				gen18141 := Call(__e, __e.Global(symcons_2), V5007)

				var gen18142 Obj
				if True == gen18141 {
					gen18138 := Call(__e, __e.Global(symhd), V5007)

					gen18139 := Call(__e, __e.Global(symcons_2), gen18138)

					var gen18140 Obj
					if True == gen18139 {
						gen18134 := Call(__e, __e.Global(symhd), V5007)

						gen18135 := Call(__e, __e.Global(symtl), gen18134)

						gen18136 := Call(__e, __e.Global(symcons_2), gen18135)

						var gen18137 Obj
						if True == gen18136 {
							gen18129 := Call(__e, __e.Global(symhd), V5007)

							gen18130 := Call(__e, __e.Global(symtl), gen18129)

							gen18131 := Call(__e, __e.Global(symtl), gen18130)

							gen18132 := Call(__e, __e.Global(sym_a), Nil, gen18131)

							var gen18133 Obj
							if True == gen18132 {
								gen18126 := Call(__e, __e.Global(symhd), V5007)

								gen18127 := Call(__e, __e.Global(symhd), gen18126)

								gen18128 := Call(__e, __e.Global(sym_a), gen18127, V5006)

								if True == gen18128 {
									gen18133 = True
								} else {
									gen18133 = False
								}

							} else {
								gen18133 = False
							}
							if True == gen18133 {
								gen18137 = True
							} else {
								gen18137 = False
							}

						} else {
							gen18137 = False
						}
						if True == gen18137 {
							gen18140 = True
						} else {
							gen18140 = False
						}

					} else {
						gen18140 = False
					}
					if True == gen18140 {
						gen18142 = True
					} else {
						gen18142 = False
					}

				} else {
					gen18142 = False
				}
				if True == gen18142 {
					gen18123 := Call(__e, __e.Global(symhd), V5007)

					gen18124 := Call(__e, __e.Global(symtl), gen18123)

					gen18125 := Call(__e, __e.Global(symcons), True, gen18124)

					__e.TailApply(__e.Global(symcons), gen18125, Nil)

					return

				} else {
					__e.Return(Nil)
					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.true-branch"), gen18203)

		gen18275 := MakeNative(func(__e Evaluator) {
			V5016 := __e.Get(1)
			_ = V5016
			V5017 := __e.Get(2)
			_ = V5017
			gen18273 := Call(__e, __e.Global(symcons_2), V5017)

			var gen18274 Obj
			if True == gen18273 {
				gen18270 := Call(__e, __e.Global(symhd), V5017)

				gen18271 := Call(__e, __e.Global(symcons_2), gen18270)

				var gen18272 Obj
				if True == gen18271 {
					gen18266 := Call(__e, __e.Global(symhd), V5017)

					gen18267 := Call(__e, __e.Global(symhd), gen18266)

					gen18268 := Call(__e, __e.Global(symcons_2), gen18267)

					var gen18269 Obj
					if True == gen18268 {
						gen18261 := Call(__e, __e.Global(symhd), V5017)

						gen18262 := Call(__e, __e.Global(symhd), gen18261)

						gen18263 := Call(__e, __e.Global(symhd), gen18262)

						gen18264 := Call(__e, __e.Global(sym_a), MakeSymbol("and"), gen18263)

						var gen18265 Obj
						if True == gen18264 {
							gen18256 := Call(__e, __e.Global(symhd), V5017)

							gen18257 := Call(__e, __e.Global(symhd), gen18256)

							gen18258 := Call(__e, __e.Global(symtl), gen18257)

							gen18259 := Call(__e, __e.Global(symcons_2), gen18258)

							var gen18260 Obj
							if True == gen18259 {
								gen18250 := Call(__e, __e.Global(symhd), V5017)

								gen18251 := Call(__e, __e.Global(symhd), gen18250)

								gen18252 := Call(__e, __e.Global(symtl), gen18251)

								gen18253 := Call(__e, __e.Global(symtl), gen18252)

								gen18254 := Call(__e, __e.Global(symcons_2), gen18253)

								var gen18255 Obj
								if True == gen18254 {
									gen18243 := Call(__e, __e.Global(symhd), V5017)

									gen18244 := Call(__e, __e.Global(symhd), gen18243)

									gen18245 := Call(__e, __e.Global(symtl), gen18244)

									gen18246 := Call(__e, __e.Global(symtl), gen18245)

									gen18247 := Call(__e, __e.Global(symtl), gen18246)

									gen18248 := Call(__e, __e.Global(sym_a), Nil, gen18247)

									var gen18249 Obj
									if True == gen18248 {
										gen18239 := Call(__e, __e.Global(symhd), V5017)

										gen18240 := Call(__e, __e.Global(symtl), gen18239)

										gen18241 := Call(__e, __e.Global(symcons_2), gen18240)

										var gen18242 Obj
										if True == gen18241 {
											gen18234 := Call(__e, __e.Global(symhd), V5017)

											gen18235 := Call(__e, __e.Global(symtl), gen18234)

											gen18236 := Call(__e, __e.Global(symtl), gen18235)

											gen18237 := Call(__e, __e.Global(sym_a), Nil, gen18236)

											var gen18238 Obj
											if True == gen18237 {
												gen18229 := Call(__e, __e.Global(symhd), V5017)

												gen18230 := Call(__e, __e.Global(symhd), gen18229)

												gen18231 := Call(__e, __e.Global(symtl), gen18230)

												gen18232 := Call(__e, __e.Global(symhd), gen18231)

												gen18233 := Call(__e, __e.Global(sym_a), gen18232, V5016)

												if True == gen18233 {
													gen18238 = True
												} else {
													gen18238 = False
												}

											} else {
												gen18238 = False
											}
											if True == gen18238 {
												gen18242 = True
											} else {
												gen18242 = False
											}

										} else {
											gen18242 = False
										}
										if True == gen18242 {
											gen18249 = True
										} else {
											gen18249 = False
										}

									} else {
										gen18249 = False
									}
									if True == gen18249 {
										gen18255 = True
									} else {
										gen18255 = False
									}

								} else {
									gen18255 = False
								}
								if True == gen18255 {
									gen18260 = True
								} else {
									gen18260 = False
								}

							} else {
								gen18260 = False
							}
							if True == gen18260 {
								gen18265 = True
							} else {
								gen18265 = False
							}

						} else {
							gen18265 = False
						}
						if True == gen18265 {
							gen18269 = True
						} else {
							gen18269 = False
						}

					} else {
						gen18269 = False
					}
					if True == gen18269 {
						gen18272 = True
					} else {
						gen18272 = False
					}

				} else {
					gen18272 = False
				}
				if True == gen18272 {
					gen18274 = True
				} else {
					gen18274 = False
				}

			} else {
				gen18274 = False
			}
			if True == gen18274 {
				gen18224 := Call(__e, __e.Global(symhd), V5017)

				gen18225 := Call(__e, __e.Global(symhd), gen18224)

				gen18226 := Call(__e, __e.Global(symtl), gen18225)

				gen18227 := Call(__e, __e.Global(symhd), gen18226)

				gen18228 := Call(__e, __e.Global(symtl), V5017)

				__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4false_1branch), gen18227, gen18228)

				return

			} else {
				gen18222 := Call(__e, __e.Global(symcons_2), V5017)

				var gen18223 Obj
				if True == gen18222 {
					gen18219 := Call(__e, __e.Global(symhd), V5017)

					gen18220 := Call(__e, __e.Global(symcons_2), gen18219)

					var gen18221 Obj
					if True == gen18220 {
						gen18215 := Call(__e, __e.Global(symhd), V5017)

						gen18216 := Call(__e, __e.Global(symtl), gen18215)

						gen18217 := Call(__e, __e.Global(symcons_2), gen18216)

						var gen18218 Obj
						if True == gen18217 {
							gen18210 := Call(__e, __e.Global(symhd), V5017)

							gen18211 := Call(__e, __e.Global(symtl), gen18210)

							gen18212 := Call(__e, __e.Global(symtl), gen18211)

							gen18213 := Call(__e, __e.Global(sym_a), Nil, gen18212)

							var gen18214 Obj
							if True == gen18213 {
								gen18207 := Call(__e, __e.Global(symhd), V5017)

								gen18208 := Call(__e, __e.Global(symhd), gen18207)

								gen18209 := Call(__e, __e.Global(sym_a), gen18208, V5016)

								if True == gen18209 {
									gen18214 = True
								} else {
									gen18214 = False
								}

							} else {
								gen18214 = False
							}
							if True == gen18214 {
								gen18218 = True
							} else {
								gen18218 = False
							}

						} else {
							gen18218 = False
						}
						if True == gen18218 {
							gen18221 = True
						} else {
							gen18221 = False
						}

					} else {
						gen18221 = False
					}
					if True == gen18221 {
						gen18223 = True
					} else {
						gen18223 = False
					}

				} else {
					gen18223 = False
				}
				if True == gen18223 {
					gen18204 := Call(__e, __e.Global(symhd), V5017)

					gen18205 := Call(__e, __e.Global(symhd), gen18204)

					gen18206 := Call(__e, __e.Global(symtl), V5017)

					__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4false_1branch), gen18205, gen18206)

					return

				} else {
					__e.Return(V5017)
					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.false-branch"), gen18275)

		gen18317 := MakeNative(func(__e Evaluator) {
			V5020 := __e.Get(1)
			_ = V5020
			V5021 := __e.Get(2)
			_ = V5021
			gen18315 := Call(__e, __e.Global(symcons_2), V5020)

			var gen18316 Obj
			if True == gen18315 {
				gen18312 := Call(__e, __e.Global(symhd), V5020)

				gen18313 := Call(__e, __e.Global(sym_a), MakeSymbol("%%return"), gen18312)

				var gen18314 Obj
				if True == gen18313 {
					gen18309 := Call(__e, __e.Global(symtl), V5020)

					gen18310 := Call(__e, __e.Global(symcons_2), gen18309)

					var gen18311 Obj
					if True == gen18310 {
						gen18305 := Call(__e, __e.Global(symtl), V5020)

						gen18306 := Call(__e, __e.Global(symtl), gen18305)

						gen18307 := Call(__e, __e.Global(sym_a), Nil, gen18306)

						var gen18308 Obj
						if True == gen18307 {
							gen18301 := Call(__e, __e.Global(symtl), V5020)

							gen18302 := Call(__e, __e.Global(symhd), gen18301)

							gen18303 := Call(__e, __e.Global(symcons_2), gen18302)

							gen18304 := Call(__e, __e.Global(symnot), gen18303)

							if True == gen18304 {
								gen18308 = True
							} else {
								gen18308 = False
							}

						} else {
							gen18308 = False
						}
						if True == gen18308 {
							gen18311 = True
						} else {
							gen18311 = False
						}

					} else {
						gen18311 = False
					}
					if True == gen18311 {
						gen18314 = True
					} else {
						gen18314 = False
					}

				} else {
					gen18314 = False
				}
				if True == gen18314 {
					gen18316 = True
				} else {
					gen18316 = False
				}

			} else {
				gen18316 = False
			}
			if True == gen18316 {
				__e.TailApply(V5021, V5020)

				return
			} else {
				gen18299 := Call(__e, __e.Global(symcons_2), V5020)

				var gen18300 Obj
				if True == gen18299 {
					gen18296 := Call(__e, __e.Global(symhd), V5020)

					gen18297 := Call(__e, __e.Global(sym_a), MakeSymbol("fail"), gen18296)

					var gen18298 Obj
					if True == gen18297 {
						gen18294 := Call(__e, __e.Global(symtl), V5020)

						gen18295 := Call(__e, __e.Global(sym_a), Nil, gen18294)

						if True == gen18295 {
							gen18298 = True
						} else {
							gen18298 = False
						}

					} else {
						gen18298 = False
					}
					if True == gen18298 {
						gen18300 = True
					} else {
						gen18300 = False
					}

				} else {
					gen18300 = False
				}
				if True == gen18300 {
					__e.TailApply(V5021, V5020)

					return
				} else {
					gen18292 := Call(__e, __e.Global(symcons_2), V5020)

					var gen18293 Obj
					if True == gen18292 {
						gen18289 := Call(__e, __e.Global(symhd), V5020)

						gen18290 := Call(__e, __e.Global(sym_a), MakeSymbol("%%goto-label"), gen18289)

						var gen18291 Obj
						if True == gen18290 {
							gen18286 := Call(__e, __e.Global(symtl), V5020)

							gen18287 := Call(__e, __e.Global(symcons_2), gen18286)

							var gen18288 Obj
							if True == gen18287 {
								gen18283 := Call(__e, __e.Global(symtl), V5020)

								gen18284 := Call(__e, __e.Global(symtl), gen18283)

								gen18285 := Call(__e, __e.Global(sym_a), Nil, gen18284)

								if True == gen18285 {
									gen18288 = True
								} else {
									gen18288 = False
								}

							} else {
								gen18288 = False
							}
							if True == gen18288 {
								gen18291 = True
							} else {
								gen18291 = False
							}

						} else {
							gen18291 = False
						}
						if True == gen18291 {
							gen18293 = True
						} else {
							gen18293 = False
						}

					} else {
						gen18293 = False
					}
					if True == gen18293 {
						__e.TailApply(V5021, V5020)

						return
					} else {
						gen18276 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4generate_1label))

						Label := gen18276
						gen18277 := Call(__e, __e.Global(symcons), Label, Nil)

						gen18278 := Call(__e, __e.Global(symcons), MakeSymbol("%%goto-label"), gen18277)

						gen18279 := Call(__e, V5021, gen18278)

						gen18280 := Call(__e, __e.Global(symcons), gen18279, Nil)

						gen18281 := Call(__e, __e.Global(symcons), V5020, gen18280)

						gen18282 := Call(__e, __e.Global(symcons), Label, gen18281)

						__e.TailApply(__e.Global(symcons), MakeSymbol("%%let-label"), gen18282)

						return

					}

				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.with-labelled-else"), gen18317)

		gen18415 := MakeNative(func(__e Evaluator) {
			V5024 := __e.Get(1)
			_ = V5024
			gen18413 := Call(__e, __e.Global(symcons_2), V5024)

			var gen18414 Obj
			if True == gen18413 {
				gen18410 := Call(__e, __e.Global(symhd), V5024)

				gen18411 := Call(__e, __e.Global(sym_a), MakeSymbol("if"), gen18410)

				var gen18412 Obj
				if True == gen18411 {
					gen18407 := Call(__e, __e.Global(symtl), V5024)

					gen18408 := Call(__e, __e.Global(symcons_2), gen18407)

					var gen18409 Obj
					if True == gen18408 {
						gen18403 := Call(__e, __e.Global(symtl), V5024)

						gen18404 := Call(__e, __e.Global(symtl), gen18403)

						gen18405 := Call(__e, __e.Global(symcons_2), gen18404)

						var gen18406 Obj
						if True == gen18405 {
							gen18398 := Call(__e, __e.Global(symtl), V5024)

							gen18399 := Call(__e, __e.Global(symtl), gen18398)

							gen18400 := Call(__e, __e.Global(symhd), gen18399)

							gen18401 := Call(__e, __e.Global(symcons_2), gen18400)

							var gen18402 Obj
							if True == gen18401 {
								gen18392 := Call(__e, __e.Global(symtl), V5024)

								gen18393 := Call(__e, __e.Global(symtl), gen18392)

								gen18394 := Call(__e, __e.Global(symhd), gen18393)

								gen18395 := Call(__e, __e.Global(symhd), gen18394)

								gen18396 := Call(__e, __e.Global(sym_a), MakeSymbol("if"), gen18395)

								var gen18397 Obj
								if True == gen18396 {
									gen18386 := Call(__e, __e.Global(symtl), V5024)

									gen18387 := Call(__e, __e.Global(symtl), gen18386)

									gen18388 := Call(__e, __e.Global(symhd), gen18387)

									gen18389 := Call(__e, __e.Global(symtl), gen18388)

									gen18390 := Call(__e, __e.Global(symcons_2), gen18389)

									var gen18391 Obj
									if True == gen18390 {
										gen18379 := Call(__e, __e.Global(symtl), V5024)

										gen18380 := Call(__e, __e.Global(symtl), gen18379)

										gen18381 := Call(__e, __e.Global(symhd), gen18380)

										gen18382 := Call(__e, __e.Global(symtl), gen18381)

										gen18383 := Call(__e, __e.Global(symtl), gen18382)

										gen18384 := Call(__e, __e.Global(symcons_2), gen18383)

										var gen18385 Obj
										if True == gen18384 {
											gen18371 := Call(__e, __e.Global(symtl), V5024)

											gen18372 := Call(__e, __e.Global(symtl), gen18371)

											gen18373 := Call(__e, __e.Global(symhd), gen18372)

											gen18374 := Call(__e, __e.Global(symtl), gen18373)

											gen18375 := Call(__e, __e.Global(symtl), gen18374)

											gen18376 := Call(__e, __e.Global(symtl), gen18375)

											gen18377 := Call(__e, __e.Global(symcons_2), gen18376)

											var gen18378 Obj
											if True == gen18377 {
												gen18362 := Call(__e, __e.Global(symtl), V5024)

												gen18363 := Call(__e, __e.Global(symtl), gen18362)

												gen18364 := Call(__e, __e.Global(symhd), gen18363)

												gen18365 := Call(__e, __e.Global(symtl), gen18364)

												gen18366 := Call(__e, __e.Global(symtl), gen18365)

												gen18367 := Call(__e, __e.Global(symtl), gen18366)

												gen18368 := Call(__e, __e.Global(symtl), gen18367)

												gen18369 := Call(__e, __e.Global(sym_a), Nil, gen18368)

												var gen18370 Obj
												if True == gen18369 {
													gen18357 := Call(__e, __e.Global(symtl), V5024)

													gen18358 := Call(__e, __e.Global(symtl), gen18357)

													gen18359 := Call(__e, __e.Global(symtl), gen18358)

													gen18360 := Call(__e, __e.Global(symcons_2), gen18359)

													var gen18361 Obj
													if True == gen18360 {
														gen18351 := Call(__e, __e.Global(symtl), V5024)

														gen18352 := Call(__e, __e.Global(symtl), gen18351)

														gen18353 := Call(__e, __e.Global(symtl), gen18352)

														gen18354 := Call(__e, __e.Global(symtl), gen18353)

														gen18355 := Call(__e, __e.Global(sym_a), Nil, gen18354)

														var gen18356 Obj
														if True == gen18355 {
															gen18339 := Call(__e, __e.Global(symtl), V5024)

															gen18340 := Call(__e, __e.Global(symtl), gen18339)

															gen18341 := Call(__e, __e.Global(symtl), gen18340)

															gen18342 := Call(__e, __e.Global(symhd), gen18341)

															gen18343 := Call(__e, __e.Global(symtl), V5024)

															gen18344 := Call(__e, __e.Global(symtl), gen18343)

															gen18345 := Call(__e, __e.Global(symhd), gen18344)

															gen18346 := Call(__e, __e.Global(symtl), gen18345)

															gen18347 := Call(__e, __e.Global(symtl), gen18346)

															gen18348 := Call(__e, __e.Global(symtl), gen18347)

															gen18349 := Call(__e, __e.Global(symhd), gen18348)

															gen18350 := Call(__e, __e.Global(sym_a), gen18342, gen18349)

															if True == gen18350 {
																gen18356 = True
															} else {
																gen18356 = False
															}

														} else {
															gen18356 = False
														}
														if True == gen18356 {
															gen18361 = True
														} else {
															gen18361 = False
														}

													} else {
														gen18361 = False
													}
													if True == gen18361 {
														gen18370 = True
													} else {
														gen18370 = False
													}

												} else {
													gen18370 = False
												}
												if True == gen18370 {
													gen18378 = True
												} else {
													gen18378 = False
												}

											} else {
												gen18378 = False
											}
											if True == gen18378 {
												gen18385 = True
											} else {
												gen18385 = False
											}

										} else {
											gen18385 = False
										}
										if True == gen18385 {
											gen18391 = True
										} else {
											gen18391 = False
										}

									} else {
										gen18391 = False
									}
									if True == gen18391 {
										gen18397 = True
									} else {
										gen18397 = False
									}

								} else {
									gen18397 = False
								}
								if True == gen18397 {
									gen18402 = True
								} else {
									gen18402 = False
								}

							} else {
								gen18402 = False
							}
							if True == gen18402 {
								gen18406 = True
							} else {
								gen18406 = False
							}

						} else {
							gen18406 = False
						}
						if True == gen18406 {
							gen18409 = True
						} else {
							gen18409 = False
						}

					} else {
						gen18409 = False
					}
					if True == gen18409 {
						gen18412 = True
					} else {
						gen18412 = False
					}

				} else {
					gen18412 = False
				}
				if True == gen18412 {
					gen18414 = True
				} else {
					gen18414 = False
				}

			} else {
				gen18414 = False
			}
			if True == gen18414 {
				gen18318 := Call(__e, __e.Global(symtl), V5024)

				gen18319 := Call(__e, __e.Global(symhd), gen18318)

				gen18320 := Call(__e, __e.Global(symtl), V5024)

				gen18321 := Call(__e, __e.Global(symtl), gen18320)

				gen18322 := Call(__e, __e.Global(symhd), gen18321)

				gen18323 := Call(__e, __e.Global(symtl), gen18322)

				gen18324 := Call(__e, __e.Global(symhd), gen18323)

				gen18325 := Call(__e, __e.Global(symcons), gen18324, Nil)

				gen18326 := Call(__e, __e.Global(symcons), gen18319, gen18325)

				gen18327 := Call(__e, __e.Global(symcons), MakeSymbol("and"), gen18326)

				gen18328 := Call(__e, __e.Global(symtl), V5024)

				gen18329 := Call(__e, __e.Global(symtl), gen18328)

				gen18330 := Call(__e, __e.Global(symhd), gen18329)

				gen18331 := Call(__e, __e.Global(symtl), gen18330)

				gen18332 := Call(__e, __e.Global(symtl), gen18331)

				gen18333 := Call(__e, __e.Global(symhd), gen18332)

				gen18334 := Call(__e, __e.Global(symtl), V5024)

				gen18335 := Call(__e, __e.Global(symtl), gen18334)

				gen18336 := Call(__e, __e.Global(symtl), gen18335)

				gen18337 := Call(__e, __e.Global(symcons), gen18333, gen18336)

				gen18338 := Call(__e, __e.Global(symcons), gen18327, gen18337)

				__e.TailApply(__e.Global(symcons), MakeSymbol("if"), gen18338)

				return

			} else {
				__e.Return(V5024)
				return
			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.merge-same-else-ifs"), gen18415)

		gen18417 := MakeNative(func(__e Evaluator) {
			V5027 := __e.Get(1)
			_ = V5027
			V5028 := __e.Get(2)
			_ = V5028
			gen18416 := Call(__e, __e.Global(symconcat), MakeSymbol("/"), V5028)

			__e.TailApply(__e.Global(symconcat), V5027, gen18416)

			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.concat/"), gen18417)

		gen18435 := MakeNative(func(__e Evaluator) {
			V5032 := __e.Get(1)
			_ = V5032
			gen18433 := Call(__e, __e.Global(symcons_2), V5032)

			var gen18434 Obj
			if True == gen18433 {
				gen18430 := Call(__e, __e.Global(symtl), V5032)

				gen18431 := Call(__e, __e.Global(symcons_2), gen18430)

				var gen18432 Obj
				if True == gen18431 {
					gen18426 := Call(__e, __e.Global(symtl), V5032)

					gen18427 := Call(__e, __e.Global(symtl), gen18426)

					gen18428 := Call(__e, __e.Global(sym_a), Nil, gen18427)

					var gen18429 Obj
					if True == gen18428 {
						gen18424 := Call(__e, __e.Global(symhd), V5032)

						gen18425 := Call(__e, __e.Global(symsymbol_2), gen18424)

						if True == gen18425 {
							gen18429 = True
						} else {
							gen18429 = False
						}

					} else {
						gen18429 = False
					}
					if True == gen18429 {
						gen18432 = True
					} else {
						gen18432 = False
					}

				} else {
					gen18432 = False
				}
				if True == gen18432 {
					gen18434 = True
				} else {
					gen18434 = False
				}

			} else {
				gen18434 = False
			}
			if True == gen18434 {
				gen18420 := Call(__e, __e.Global(symtl), V5032)

				gen18421 := Call(__e, __e.Global(symhd), gen18420)

				gen18422 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4exp_1var), gen18421)

				gen18423 := Call(__e, __e.Global(symhd), V5032)

				__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4concat_c), gen18422, gen18423)

				return

			} else {
				gen18419 := Call(__e, __e.Global(symcons_2), V5032)

				if True == gen18419 {
					gen18418 := Call(__e, __e.Global(symhd), V5032)

					__e.TailApply(__e.Global(symgensym), gen18418)

					return

				} else {
					__e.Return(V5032)
					return
				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.exp-var"), gen18435)

		gen18437 := MakeNative(func(__e Evaluator) {
			V5035 := __e.Get(1)
			_ = V5035
			V5036 := __e.Get(2)
			_ = V5036
			gen18436 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4test_1_6selectors), V5035)

			__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4bind_1repeating_1selectors), gen18436, V5036)

			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.optimize-selectors"), gen18437)

		gen18506 := MakeNative(func(__e Evaluator) {
			V5042 := __e.Get(1)
			_ = V5042
			gen18504 := Call(__e, __e.Global(symcons_2), V5042)

			var gen18505 Obj
			if True == gen18504 {
				gen18501 := Call(__e, __e.Global(symhd), V5042)

				gen18502 := Call(__e, __e.Global(sym_a), MakeSymbol("cons?"), gen18501)

				var gen18503 Obj
				if True == gen18502 {
					gen18498 := Call(__e, __e.Global(symtl), V5042)

					gen18499 := Call(__e, __e.Global(symcons_2), gen18498)

					var gen18500 Obj
					if True == gen18499 {
						gen18495 := Call(__e, __e.Global(symtl), V5042)

						gen18496 := Call(__e, __e.Global(symtl), gen18495)

						gen18497 := Call(__e, __e.Global(sym_a), Nil, gen18496)

						if True == gen18497 {
							gen18500 = True
						} else {
							gen18500 = False
						}

					} else {
						gen18500 = False
					}
					if True == gen18500 {
						gen18503 = True
					} else {
						gen18503 = False
					}

				} else {
					gen18503 = False
				}
				if True == gen18503 {
					gen18505 = True
				} else {
					gen18505 = False
				}

			} else {
				gen18505 = False
			}
			if True == gen18505 {
				gen18490 := Call(__e, __e.Global(symtl), V5042)

				gen18491 := Call(__e, __e.Global(symcons), MakeSymbol("hd"), gen18490)

				gen18492 := Call(__e, __e.Global(symtl), V5042)

				gen18493 := Call(__e, __e.Global(symcons), MakeSymbol("tl"), gen18492)

				gen18494 := Call(__e, __e.Global(symcons), gen18493, Nil)

				__e.TailApply(__e.Global(symcons), gen18491, gen18494)

				return

			} else {
				gen18488 := Call(__e, __e.Global(symcons_2), V5042)

				var gen18489 Obj
				if True == gen18488 {
					gen18485 := Call(__e, __e.Global(symhd), V5042)

					gen18486 := Call(__e, __e.Global(sym_a), MakeSymbol("tuple?"), gen18485)

					var gen18487 Obj
					if True == gen18486 {
						gen18482 := Call(__e, __e.Global(symtl), V5042)

						gen18483 := Call(__e, __e.Global(symcons_2), gen18482)

						var gen18484 Obj
						if True == gen18483 {
							gen18479 := Call(__e, __e.Global(symtl), V5042)

							gen18480 := Call(__e, __e.Global(symtl), gen18479)

							gen18481 := Call(__e, __e.Global(sym_a), Nil, gen18480)

							if True == gen18481 {
								gen18484 = True
							} else {
								gen18484 = False
							}

						} else {
							gen18484 = False
						}
						if True == gen18484 {
							gen18487 = True
						} else {
							gen18487 = False
						}

					} else {
						gen18487 = False
					}
					if True == gen18487 {
						gen18489 = True
					} else {
						gen18489 = False
					}

				} else {
					gen18489 = False
				}
				if True == gen18489 {
					gen18474 := Call(__e, __e.Global(symtl), V5042)

					gen18475 := Call(__e, __e.Global(symcons), MakeSymbol("fst"), gen18474)

					gen18476 := Call(__e, __e.Global(symtl), V5042)

					gen18477 := Call(__e, __e.Global(symcons), MakeSymbol("snd"), gen18476)

					gen18478 := Call(__e, __e.Global(symcons), gen18477, Nil)

					__e.TailApply(__e.Global(symcons), gen18475, gen18478)

					return

				} else {
					gen18472 := Call(__e, __e.Global(symcons_2), V5042)

					var gen18473 Obj
					if True == gen18472 {
						gen18469 := Call(__e, __e.Global(symhd), V5042)

						gen18470 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.+string?"), gen18469)

						var gen18471 Obj
						if True == gen18470 {
							gen18466 := Call(__e, __e.Global(symtl), V5042)

							gen18467 := Call(__e, __e.Global(symcons_2), gen18466)

							var gen18468 Obj
							if True == gen18467 {
								gen18463 := Call(__e, __e.Global(symtl), V5042)

								gen18464 := Call(__e, __e.Global(symtl), gen18463)

								gen18465 := Call(__e, __e.Global(sym_a), Nil, gen18464)

								if True == gen18465 {
									gen18468 = True
								} else {
									gen18468 = False
								}

							} else {
								gen18468 = False
							}
							if True == gen18468 {
								gen18471 = True
							} else {
								gen18471 = False
							}

						} else {
							gen18471 = False
						}
						if True == gen18471 {
							gen18473 = True
						} else {
							gen18473 = False
						}

					} else {
						gen18473 = False
					}
					if True == gen18473 {
						gen18458 := Call(__e, __e.Global(symtl), V5042)

						gen18459 := Call(__e, __e.Global(symcons), MakeSymbol("hdstr"), gen18458)

						gen18460 := Call(__e, __e.Global(symtl), V5042)

						gen18461 := Call(__e, __e.Global(symcons), MakeSymbol("tlstr"), gen18460)

						gen18462 := Call(__e, __e.Global(symcons), gen18461, Nil)

						__e.TailApply(__e.Global(symcons), gen18459, gen18462)

						return

					} else {
						gen18456 := Call(__e, __e.Global(symcons_2), V5042)

						var gen18457 Obj
						if True == gen18456 {
							gen18453 := Call(__e, __e.Global(symhd), V5042)

							gen18454 := Call(__e, __e.Global(sym_a), MakeSymbol("shen.+vector?"), gen18453)

							var gen18455 Obj
							if True == gen18454 {
								gen18450 := Call(__e, __e.Global(symtl), V5042)

								gen18451 := Call(__e, __e.Global(symcons_2), gen18450)

								var gen18452 Obj
								if True == gen18451 {
									gen18447 := Call(__e, __e.Global(symtl), V5042)

									gen18448 := Call(__e, __e.Global(symtl), gen18447)

									gen18449 := Call(__e, __e.Global(sym_a), Nil, gen18448)

									if True == gen18449 {
										gen18452 = True
									} else {
										gen18452 = False
									}

								} else {
									gen18452 = False
								}
								if True == gen18452 {
									gen18455 = True
								} else {
									gen18455 = False
								}

							} else {
								gen18455 = False
							}
							if True == gen18455 {
								gen18457 = True
							} else {
								gen18457 = False
							}

						} else {
							gen18457 = False
						}
						if True == gen18457 {
							gen18442 := Call(__e, __e.Global(symtl), V5042)

							gen18443 := Call(__e, __e.Global(symcons), MakeSymbol("hdv"), gen18442)

							gen18444 := Call(__e, __e.Global(symtl), V5042)

							gen18445 := Call(__e, __e.Global(symcons), MakeSymbol("tlv"), gen18444)

							gen18446 := Call(__e, __e.Global(symcons), gen18445, Nil)

							__e.TailApply(__e.Global(symcons), gen18443, gen18446)

							return

						} else {
							gen18438 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

							gen18439 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4apply_1selector_1handlers), gen18438, V5042)

							Result := gen18439
							gen18440 := Call(__e, __e.Global(symfail))

							gen18441 := Call(__e, __e.Global(sym_a), Result, gen18440)

							if True == gen18441 {
								__e.Return(Nil)
								return
							} else {
								__e.Return(Result)
								return
							}

						}

					}

				}

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.test->selectors"), gen18506)

		gen18512 := MakeNative(func(__e Evaluator) {
			V5045 := __e.Get(1)
			_ = V5045
			V5046 := __e.Get(2)
			_ = V5046
			gen18511 := Call(__e, __e.Global(symcons_2), V5045)

			if True == gen18511 {
				gen18508 := Call(__e, __e.Global(symhd), V5045)

				gen18509 := Call(__e, __e.Global(symtl), V5045)

				gen18510 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4bind_1repeating_1selectors), gen18509, V5046)

				__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4bind_1selector), gen18508, gen18510)

				return

			} else {
				gen18507 := Call(__e, __e.Global(sym_a), Nil, V5045)

				if True == gen18507 {
					__e.Return(V5046)
					return
				} else {
					__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors"))

					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors"), gen18512)

		gen18520 := MakeNative(func(__e Evaluator) {
			V5053 := __e.Get(1)
			_ = V5053
			V5054 := __e.Get(2)
			_ = V5054
			gen18518 := Call(__e, __e.Global(symoccurrences), V5053, V5054)

			gen18519 := Call(__e, __e.Global(sym_6), gen18518, MakeNumber(1))

			if True == gen18519 {
				gen18513 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4exp_1var), V5053)

				Var := gen18513
				gen18514 := Call(__e, __e.Global(symsubst), Var, V5053, V5054)

				gen18515 := Call(__e, __e.Global(symcons), gen18514, Nil)

				gen18516 := Call(__e, __e.Global(symcons), V5053, gen18515)

				gen18517 := Call(__e, __e.Global(symcons), Var, gen18516)

				__e.TailApply(__e.Global(symcons), MakeSymbol("let"), gen18517)

				return

			} else {
				__e.Return(V5054)
				return
			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.bind-selector"), gen18520)

		gen18530 := MakeNative(func(__e Evaluator) {
			V5067 := __e.Get(1)
			_ = V5067
			V5068 := __e.Get(2)
			_ = V5068
			gen18529 := Call(__e, __e.Global(sym_a), Nil, V5067)

			if True == gen18529 {
				__e.TailApply(__e.Global(symfail))

				return
			} else {
				gen18523 := MakeNative(func(__e Evaluator) {
					gen18522 := Call(__e, __e.Global(symcons_2), V5067)

					if True == gen18522 {
						gen18521 := Call(__e, __e.Global(symtl), V5067)

						__e.TailApply(__e.Global(symshen_4x_4factorise_1defun_4apply_1selector_1handlers), gen18521, V5068)

						return

					} else {
						__e.TailApply(__e.Global(symshen_4f__error), MakeSymbol("shen.x.factorise-defun.apply-selector-handlers"))

						return
					}

				}, 0)
				Freeze := gen18523
				gen18528 := Call(__e, __e.Global(symcons_2), V5067)

				if True == gen18528 {
					gen18524 := Call(__e, __e.Global(symhd), V5067)

					f34 := gen18524
					gen18525 := Call(__e, f34, V5068)

					Result := gen18525
					gen18526 := Call(__e, __e.Global(symfail))

					gen18527 := Call(__e, __e.Global(sym_a), Result, gen18526)

					if True == gen18527 {
						__e.TailApply(__e.Global(symthaw), Freeze)

						return
					} else {
						__e.Return(Result)
						return
					}

				} else {
					__e.TailApply(__e.Global(symthaw), Freeze)

					return
				}

			}

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.apply-selector-handlers"), gen18530)

		gen18531 := MakeNative(func(__e Evaluator) {
			Call(__e, __e.Global(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"), Nil)
			Call(__e, __e.Global(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"), Nil)
			__e.Return(MakeSymbol("shen.x.factorise-defun.done"))
			return

		}, 0)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.initialise"), gen18531)

		gen18539 := MakeNative(func(__e Evaluator) {
			V5070 := __e.Get(1)
			_ = V5070
			gen18537 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

			gen18538 := Call(__e, __e.Global(symelement_2), V5070, gen18537)

			if True == gen18538 {
				__e.Return(V5070)
				return
			} else {
				gen18532 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

				gen18533 := Call(__e, __e.Global(symcons), V5070, gen18532)

				Call(__e, __e.Global(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"), gen18533)

				gen18534 := Call(__e, __e.Global(symfunction), V5070)

				gen18535 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

				gen18536 := Call(__e, __e.Global(symcons), gen18534, gen18535)

				Call(__e, __e.Global(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"), gen18536)

				__e.Return(V5070)
				return

			}

		}, 1)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.register-selector-handler"), gen18539)

		gen18543 := MakeNative(func(__e Evaluator) {
			V5073 := __e.Get(1)
			_ = V5073
			V5074 := __e.Get(2)
			_ = V5074
			gen18541 := MakeNative(func(__e Evaluator) {
				__ := __e.Get(1)
				_ = __
				gen18540 := Call(__e, __e.Global(symshen_4app), V5073, MakeString(" is not a selector handler\n"), MakeSymbol("shen.a"))

				__e.TailApply(__e.Global(symsimple_1error), gen18540)

				return

			}, 1)
			gen18542 := MakeNative(func(__e Evaluator) {
				__e.TailApply(__e.Global(symshen_4findpos), V5073, V5074)

				return
			}, 0)
			__e.Return(Try(__e, gen18542).Catch(gen18541))
			return

		}, 2)
		Call(__e, __e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.findpos"), gen18543)

		gen18551 := MakeNative(func(__e Evaluator) {
			V5076 := __e.Get(1)
			_ = V5076
			gen18544 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"))

			Reg := gen18544
			gen18545 := Call(__e, __e.Global(symshen_4x_4factorise_1defun_4findpos), V5076, Reg)

			Pos := gen18545
			gen18546 := Call(__e, __e.Global(symremove), V5076, Reg)

			gen18547 := Call(__e, __e.Global(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers-reg*"), gen18546)

			RemoveReg := gen18547
			_ = RemoveReg
			gen18548 := Call(__e, __e.Global(symvalue), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"))

			gen18549 := Call(__e, __e.Global(symshen_4remove_1nth), Pos, gen18548)

			gen18550 := Call(__e, __e.Global(symset), MakeSymbol("shen.x.factorise-defun.*selector-handlers*"), gen18549)

			RemoveFun := gen18550
			_ = RemoveFun
			__e.Return(V5076)
			return

		}, 1)
		__e.TailApply(__e.Global(symdefun), MakeSymbol("shen.x.factorise-defun.unregister-selector-handler"), gen18551)

		return

	}, 0))
}

var symshen_4ue = MakeSymbol("shen.ue")
var symshen_4compile__to__machine__code = MakeSymbol("shen.compile_to_machine_code")
var symadjoin = MakeSymbol("adjoin")
var sympreclude_1all_1but = MakeSymbol("preclude-all-but")
var symshen_4deref = MakeSymbol("shen.deref")
var symshen_4type_1signature_1of_1empty_2 = MakeSymbol("shen.type-signature-of-empty?")
var sym_a_a = MakeSymbol("==")
var symshen_4insert_1runon = MakeSymbol("shen.insert-runon")
var symshen_4th_d = MakeSymbol("shen.th*")
var symshen_4type_1signature_1of_1fst = MakeSymbol("shen.type-signature-of-fst")
var symshen_4dh_2 = MakeSymbol("shen.dh?")
var symsubst = MakeSymbol("subst")
var symshen_4string_1_6bytes = MakeSymbol("shen.string->bytes")
var symstring_1_6symbol = MakeSymbol("string->symbol")
var symshen_4aah = MakeSymbol("shen.aah")
var symshen_4datatype_1error = MakeSymbol("shen.datatype-error")
var symshen_4custom_1pattern_1compiler = MakeSymbol("shen.custom-pattern-compiler")
var symshen_4cond_1form = MakeSymbol("shen.cond-form")
var symshen_4extract__vars = MakeSymbol("shen.extract_vars")
var symshen_4active_1cons = MakeSymbol("shen.active-cons")
var symshen_4resize_1vector = MakeSymbol("shen.resize-vector")
var symshen_4reduce__help = MakeSymbol("shen.reduce_help")
var symshen_4_5premises_6 = MakeSymbol("shen.<premises>")
var symshen_4resizeprocessvector = MakeSymbol("shen.resizeprocessvector")
var symshen_4copy_1vector_1stage_11 = MakeSymbol("shen.copy-vector-stage-1")
var symshen_4atom_1type = MakeSymbol("shen.atom-type")
var symshen_4_5defprolog_6 = MakeSymbol("shen.<defprolog>")
var symvariable_2 = MakeSymbol("variable?")
var symshen_4singleunderline_2 = MakeSymbol("shen.singleunderline?")
var symshen_4type_1signature_1of_1and = MakeSymbol("shen.type-signature-of-and")
var symshen_4compile__to__lambda_7 = MakeSymbol("shen.compile_to_lambda+")
var symshen_4free__variable__check = MakeSymbol("shen.free_variable_check")
var symshen_4_5side_1conditions_6 = MakeSymbol("shen.<side-conditions>")
var symshen_4decimalise = MakeSymbol("shen.decimalise")
var sym_6 = MakeSymbol(">")
var symshen_4shen_1_6kl = MakeSymbol("shen.shen->kl")
var symshen_4type_1signature_1of_1difference = MakeSymbol("shen.type-signature-of-difference")
var symtail = MakeSymbol("tail")
var symshen_4_5rcurly_6 = MakeSymbol("shen.<rcurly>")
var symshen_4_5sig_7rest_6 = MakeSymbol("shen.<sig+rest>")
var symshen_4add_1macro = MakeSymbol("shen.add-macro")
var symshen_4explode_1h = MakeSymbol("shen.explode-h")
var symshen_4ephemeral__variable_2 = MakeSymbol("shen.ephemeral_variable?")
var symshen_4type_1signature_1of_1include = MakeSymbol("shen.type-signature-of-include")
var symshen_4_5_1dict_1bucket = MakeSymbol("shen.<-dict-bucket")
var symshen_4tlhd = MakeSymbol("shen.tlhd")
var symshen_4prolog__constant_2 = MakeSymbol("shen.prolog_constant?")
var symshen_4curry = MakeSymbol("shen.curry")
var symshen_4default__semantics = MakeSymbol("shen.default_semantics")
var symshen_4mk_1pvar = MakeSymbol("shen.mk-pvar")
var symshen_4errormaxinfs = MakeSymbol("shen.errormaxinfs")
var symshen_4type_1signature_1of_1systemf = MakeSymbol("shen.type-signature-of-systemf")
var symshen_4_5pattern1_6 = MakeSymbol("shen.<pattern1>")
var symshen_4t_d = MakeSymbol("shen.t*")
var symshen_4assumetype = MakeSymbol("shen.assumetype")
var symshen_4line = MakeSymbol("shen.line")
var symshen_4type_1signature_1of_1specialise = MakeSymbol("shen.type-signature-of-specialise")
var symshen_4pvar_2 = MakeSymbol("shen.pvar?")
var symnumber_2 = MakeSymbol("number?")
var symshen_4make__mu__application = MakeSymbol("shen.make_mu_application")
var syminferences = MakeSymbol("inferences")
var symshen_4show = MakeSymbol("shen.show")
var symarity = MakeSymbol("arity")
var symshen_4add__test = MakeSymbol("shen.add_test")
var symwrite_1to_1file = MakeSymbol("write-to-file")
var symshen_4construct_1base_1search_1clause = MakeSymbol("shen.construct-base-search-clause")
var symshen_4read_1error = MakeSymbol("shen.read-error")
var symshen_4function_1macro = MakeSymbol("shen.function-macro")
var symshen_4type_1signature_1of_1freeze = MakeSymbol("shen.type-signature-of-freeze")
var symshen_4type_1signature_1of_1read_1from_1string = MakeSymbol("shen.type-signature-of-read-from-string")
var symoccurrences = MakeSymbol("occurrences")
var symshen_4mode_1ify = MakeSymbol("shen.mode-ify")
var symclose = MakeSymbol("close")
var symshen_4newcontinuation = MakeSymbol("shen.newcontinuation")
var symshen_4memo = MakeSymbol("shen.memo")
var sym_5_1address = MakeSymbol("<-address")
var symshen_4dict_1fold = MakeSymbol("shen.dict-fold")
var symshen_4type_1signature_1of_1enable_1type_1theory = MakeSymbol("shen.type-signature-of-enable-type-theory")
var symnot = MakeSymbol("not")
var symshen_4assoc_1rm = MakeSymbol("shen.assoc-rm")
var sym_7 = MakeSymbol("+")
var symshen_4intprolog_1help = MakeSymbol("shen.intprolog-help")
var symshen_4x_4factorise_1defun_4free_1variables_1h = MakeSymbol("shen.x.factorise-defun.free-variables-h")
var sym_1 = MakeSymbol("-")
var symshen_4synonyms_1help = MakeSymbol("shen.synonyms-help")
var symshen_4uppercase_2 = MakeSymbol("shen.uppercase?")
var symshen_4rule_1_6horn_1clause_1head = MakeSymbol("shen.rule->horn-clause-head")
var symshen_4_5alphanum_6 = MakeSymbol("shen.<alphanum>")
var symshen_4cutpoint = MakeSymbol("shen.cutpoint")
var symprotect = MakeSymbol("protect")
var symshen_4left_1round = MakeSymbol("shen.left-round")
var symshen_4yacc = MakeSymbol("shen.yacc")
var symshen_4dict = MakeSymbol("shen.dict")
var symshen_4type_1signature_1of_1step = MakeSymbol("shen.type-signature-of-step")
var symenable_1type_1theory = MakeSymbol("enable-type-theory")
var symshen_4abs = MakeSymbol("shen.abs")
var symshen_4write_1char_1and_1inc = MakeSymbol("shen.write-char-and-inc")
var symshen_4_5datatype_1rules_6 = MakeSymbol("shen.<datatype-rules>")
var symshen_4cf__help = MakeSymbol("shen.cf_help")
var symshen_4x_4features_4cond_1expand_1macro = MakeSymbol("shen.x.features.cond-expand-macro")
var symelement_2 = MakeSymbol("element?")
var symshen_4pair = MakeSymbol("shen.pair")
var symshen_4type_1signature_1of_1string_2 = MakeSymbol("shen.type-signature-of-string?")
var symshen_4hdhd = MakeSymbol("shen.hdhd")
var symshen_4continuation__call = MakeSymbol("shen.continuation_call")
var symshen_4_5minus_6 = MakeSymbol("shen.<minus>")
var symshen_4credits = MakeSymbol("shen.credits")
var symshen_4_5datatype_1rule_6 = MakeSymbol("shen.<datatype-rule>")
var symhdv = MakeSymbol("hdv")
var symshen_4extract_1pvars = MakeSymbol("shen.extract-pvars")
var symexternal = MakeSymbol("external")
var symshen_4bld_1prolog_1call = MakeSymbol("shen.bld-prolog-call")
var symshen_4typextable = MakeSymbol("shen.typextable")
var symshen_4check__stream = MakeSymbol("shen.check_stream")
var symshen_4_5rules_6 = MakeSymbol("shen.<rules>")
var symshen_4typedf_2 = MakeSymbol("shen.typedf?")
var symshen_4type_1signature_1of_1_5_b_6 = MakeSymbol("shen.type-signature-of-<!>")
var symshen_4type_1signature_1of_1gensym = MakeSymbol("shen.type-signature-of-gensym")
var symshen_4x_4launcher_4execute_1all = MakeSymbol("shen.x.launcher.execute-all")
var symnth = MakeSymbol("nth")
var sym_5 = MakeSymbol("<")
var symoccurs_1check = MakeSymbol("occurs-check")
var symshen_4alpha_2 = MakeSymbol("shen.alpha?")
var symshen_4_5side_1condition_6 = MakeSymbol("shen.<side-condition>")
var symshen_4prbytes = MakeSymbol("shen.prbytes")
var symshen_4type_1signature_1of_1hdv = MakeSymbol("shen.type-signature-of-hdv")
var symshen_4x_4launcher_4default_1handle_1result = MakeSymbol("shen.x.launcher.default-handle-result")
var symshen_4factor_1cn = MakeSymbol("shen.factor-cn")
var symshen_4type_1signature_1of_1profile = MakeSymbol("shen.type-signature-of-profile")
var symshen_4removetype = MakeSymbol("shen.removetype")
var symshen_4x_4factorise_1defun_4with_1labelled_1else = MakeSymbol("shen.x.factorise-defun.with-labelled-else")
var symshen_4empty_1absvector_2 = MakeSymbol("shen.empty-absvector?")
var symshen_4type_1signature_1of_1track = MakeSymbol("shen.type-signature-of-track")
var symshen_4x_4launcher_4eval_1string = MakeSymbol("shen.x.launcher.eval-string")
var symshen_4case_1form = MakeSymbol("shen.case-form")
var symshen_4_5anymulti_6 = MakeSymbol("shen.<anymulti>")
var symshen_4curry_1type_1h = MakeSymbol("shen.curry-type-h")
var symshen_4application__build = MakeSymbol("shen.application_build")
var symshen_4type_1signature_1of_1stinput = MakeSymbol("shen.type-signature-of-stinput")
var symshen_4toplevel__evaluate = MakeSymbol("shen.toplevel_evaluate")
var symshen_4type_1signature_1of_1hdstr = MakeSymbol("shen.type-signature-of-hdstr")
var symput = MakeSymbol("put")
var symundefmacro = MakeSymbol("undefmacro")
var symvalue = MakeSymbol("value")
var symshen_4typetable = MakeSymbol("shen.typetable")
var symshen_4type_1signature_1of_1release = MakeSymbol("shen.type-signature-of-release")
var symshen_4fillvector = MakeSymbol("shen.fillvector")
var symshen_4valvector = MakeSymbol("shen.valvector")
var symshen_4_7string_2 = MakeSymbol("shen.+string?")
var symshen_4_5clauses_d_6 = MakeSymbol("shen.<clauses*>")
var symshen_4_5body_d_6 = MakeSymbol("shen.<body*>")
var symlanguage = MakeSymbol("language")
var symshen_4_5simple__pattern_6 = MakeSymbol("shen.<simple_pattern>")
var symshen_4custom_1pattern_1reducer = MakeSymbol("shen.custom-pattern-reducer")
var symfwhen = MakeSymbol("fwhen")
var symps = MakeSymbol("ps")
var symshen_4variancy_1test = MakeSymbol("shen.variancy-test")
var symshen_4type_1signature_1of_1thaw = MakeSymbol("shen.type-signature-of-thaw")
var symreverse = MakeSymbol("reverse")
var symshen_4initialise_1prolog = MakeSymbol("shen.initialise-prolog")
var symshen_4iter_1vector = MakeSymbol("shen.iter-vector")
var symshen_4type_1signature_1of_1port = MakeSymbol("shen.type-signature-of-port")
var symshen_4compose = MakeSymbol("shen.compose")
var symsimple_1error = MakeSymbol("simple-error")
var symshen_4integer_1test_2 = MakeSymbol("shen.integer-test?")
var symshen_4_5anysingle_6 = MakeSymbol("shen.<anysingle>")
var symshen_4maxseq = MakeSymbol("shen.maxseq")
var sym_d = MakeSymbol("*")
var symshen_4_5formula_6 = MakeSymbol("shen.<formula>")
var symshen_4type_1signature_1of_1function = MakeSymbol("shen.type-signature-of-function")
var symsystemf = MakeSymbol("systemf")
var symn_1_6string = MakeSymbol("n->string")
var sympos = MakeSymbol("pos")
var symshen_4type_1signature_1of_1ps = MakeSymbol("shen.type-signature-of-ps")
var sympr = MakeSymbol("pr")
var symshen_4t_d_1rule = MakeSymbol("shen.t*-rule")
var symshen_4_5dbq_6 = MakeSymbol("shen.<dbq>")
var symunprofile = MakeSymbol("unprofile")
var symcons_2 = MakeSymbol("cons?")
var symshen_4variable_1match = MakeSymbol("shen.variable-match")
var symshen_4cc__help = MakeSymbol("shen.cc_help")
var symshen_4type_1signature_1of_1if = MakeSymbol("shen.type-signature-of-if")
var symshen_4rfas_1h = MakeSymbol("shen.rfas-h")
var symshen_4_5non_1return_6 = MakeSymbol("shen.<non-return>")
var symshen_4type_1signature_1of_1do = MakeSymbol("shen.type-signature-of-do")
var symshen_4type_1signature_1of_1inferences = MakeSymbol("shen.type-signature-of-inferences")
var symshen_4type_1signature_1of_1trap_1error = MakeSymbol("shen.type-signature-of-trap-error")
var sym_a = MakeSymbol("=")
var symshen_4record_1it = MakeSymbol("shen.record-it")
var symshen_4prolog_1error = MakeSymbol("shen.prolog-error")
var symshen_4_5str_6 = MakeSymbol("shen.<str>")
var symshen_4_5digits_6 = MakeSymbol("shen.<digits>")
var symshen_4type_1signature_1of_1shen_4prhush = MakeSymbol("shen.type-signature-of-shen.prhush")
var symshen_4type_1signature_1of_1tlstr = MakeSymbol("shen.type-signature-of-tlstr")
var symshen_4flatten = MakeSymbol("shen.flatten")
var symshen_4type_1theory_1enabled_2 = MakeSymbol("shen.type-theory-enabled?")
var symshen_4dict_1count = MakeSymbol("shen.dict-count")
var symshen_4expt = MakeSymbol("shen.expt")
var symshen_4_5comma_1symbol_6 = MakeSymbol("shen.<comma-symbol>")
var symshen_4type_1signature_1of_1cn = MakeSymbol("shen.type-signature-of-cn")
var symshen_4repl = MakeSymbol("shen.repl")
var symshen_4x_4factorise_1defun_4bind_1selector = MakeSymbol("shen.x.factorise-defun.bind-selector")
var symshen_4dict_1fold_1h = MakeSymbol("shen.dict-fold-h")
var symshen_4curry_1type = MakeSymbol("shen.curry-type")
var symshen_4add_1end = MakeSymbol("shen.add-end")
var symshen_4t_d_1defh = MakeSymbol("shen.t*-defh")
var symvector = MakeSymbol("vector")
var symshen_4same__predicate_2 = MakeSymbol("shen.same_predicate?")
var symshen_4type_1signature_1of_1arity = MakeSymbol("shen.type-signature-of-arity")
var symfail_1if = MakeSymbol("fail-if")
var symshen_4linearise = MakeSymbol("shen.linearise")
var symshen_4numbyte_2 = MakeSymbol("shen.numbyte?")
var symshen_4_5stop_6 = MakeSymbol("shen.<stop>")
var symshen_4_5backslash_6 = MakeSymbol("shen.<backslash>")
var symshen_4nl_1macro = MakeSymbol("shen.nl-macro")
var symoptimise = MakeSymbol("optimise")
var symstoutput = MakeSymbol("stoutput")
var symeval_1kl = MakeSymbol("eval-kl")
var sym_5_a = MakeSymbol("<=")
var symshen_4demod_1rule = MakeSymbol("shen.demod-rule")
var symshen_4type_1signature_1of_1y_1or_1n_2 = MakeSymbol("shen.type-signature-of-y-or-n?")
var symshen_4type_1signature_1of_1bound_2 = MakeSymbol("shen.type-signature-of-bound?")
var symdestroy = MakeSymbol("destroy")
var symshen_4type_1signature_1of_1symbol_2 = MakeSymbol("shen.type-signature-of-symbol?")
var symshen_4dict_1_6 = MakeSymbol("shen.dict->")
var symshen_4compile_1macro = MakeSymbol("shen.compile-macro")
var symassoc = MakeSymbol("assoc")
var symshen_4get_1profile = MakeSymbol("shen.get-profile")
var symshen_4type_1signature_1of_1cons_2 = MakeSymbol("shen.type-signature-of-cons?")
var symshen_4modh = MakeSymbol("shen.modh")
var symshen_4placeholders = MakeSymbol("shen.placeholders")
var sym_5e_6 = MakeSymbol("<e>")
var symshen_4package_1contents = MakeSymbol("shen.package-contents")
var symshen_4terminator_2 = MakeSymbol("shen.terminator?")
var symshen_4defprolog_1macro = MakeSymbol("shen.defprolog-macro")
var symshen_4type_1signature_1of_1unprofile = MakeSymbol("shen.type-signature-of-unprofile")
var symshen_4x_4launcher_4quiet_1load = MakeSymbol("shen.x.launcher.quiet-load")
var symshen_4f__error = MakeSymbol("shen.f_error")
var symempty_2 = MakeSymbol("empty?")
var symshen_4type_1signature_1of_1nth = MakeSymbol("shen.type-signature-of-nth")
var symshen_4error_1macro = MakeSymbol("shen.error-macro")
var symshen_4type_1signature_1of_1not = MakeSymbol("shen.type-signature-of-not")
var symshen_4assign_1types = MakeSymbol("shen.assign-types")
var symshen_4recursive__cons__form = MakeSymbol("shen.recursive_cons_form")
var symshen_4bindv = MakeSymbol("shen.bindv")
var symshen_4track_1function = MakeSymbol("shen.track-function")
var symshen_4remove_1synonyms = MakeSymbol("shen.remove-synonyms")
var symshen_4tuple_1up = MakeSymbol("shen.tuple-up")
var symshen_4t_d_1hyps = MakeSymbol("shen.t*-hyps")
var symshen_4t_d_1defhh = MakeSymbol("shen.t*-defhh")
var symshen_4findallhelp = MakeSymbol("shen.findallhelp")
var symshen_4x_4factorise_1defun_4false_1branch = MakeSymbol("shen.x.factorise-defun.false-branch")
var symshen_4sigf = MakeSymbol("shen.sigf")
var symshen_4type_1signature_1of_1load = MakeSymbol("shen.type-signature-of-load")
var symport = MakeSymbol("port")
var symshen_4x_4launcher_4eval_1flag_1map = MakeSymbol("shen.x.launcher.eval-flag-map")
var symshen_4x_4factorise_1defun_4rebranch = MakeSymbol("shen.x.factorise-defun.rebranch")
var symshen_4doubleunderline_2 = MakeSymbol("shen.doubleunderline?")
var symshen_4funexstring = MakeSymbol("shen.funexstring")
var symshen_4_5alphanums_6 = MakeSymbol("shen.<alphanums>")
var symshen_4type_1signature_1of_1mapcan = MakeSymbol("shen.type-signature-of-mapcan")
var symset = MakeSymbol("set")
var symshen_4_5colon_6 = MakeSymbol("shen.<colon>")
var symshen_4proc_1input_7 = MakeSymbol("shen.proc-input+")
var symshen_4tlv_1help = MakeSymbol("shen.tlv-help")
var symshen_4start_1new_1prolog_1process = MakeSymbol("shen.start-new-prolog-process")
var symshen_4type_1signature_1of_1occurrences = MakeSymbol("shen.type-signature-of-occurrences")
var symshen_4compile__to__kl = MakeSymbol("shen.compile_to_kl")
var symshen_4x_4factorise_1defun_4test_1_6selectors = MakeSymbol("shen.x.factorise-defun.test->selectors")
var symshen_4x_4launcher_4launch_1shen = MakeSymbol("shen.x.launcher.launch-shen")
var symshen_4tuple = MakeSymbol("shen.tuple")
var symshen_4t_d_1action = MakeSymbol("shen.t*-action")
var symshen_4type_1signature_1of_1fail_1if = MakeSymbol("shen.type-signature-of-fail-if")
var symshen_4x_4factorise_1defun_4findpos = MakeSymbol("shen.x.factorise-defun.findpos")
var symshen_4initialise__environment = MakeSymbol("shen.initialise_environment")
var symshen_4_5st__input_6 = MakeSymbol("shen.<st_input>")
var symshen_4spaces = MakeSymbol("shen.spaces")
var symshen_4x_4factorise_1defun_4merge_1same_1else_1ifs = MakeSymbol("shen.x.factorise-defun.merge-same-else-ifs")
var symshen_4find = MakeSymbol("shen.find")
var symthaw = MakeSymbol("thaw")
var symshen_4type_1signature_1of_1tc_2 = MakeSymbol("shen.type-signature-of-tc?")
var symshen_4type_1signature_1of_1unspecialise = MakeSymbol("shen.type-signature-of-unspecialise")
var symshen_4_5doubleunderline_6 = MakeSymbol("shen.<doubleunderline>")
var symshen_4lookup_1func = MakeSymbol("shen.lookup-func")
var symshen_4list_1stream = MakeSymbol("shen.list-stream")
var symshen_4lazyderef = MakeSymbol("shen.lazyderef")
var symshen_4initialise_1lambda_1forms = MakeSymbol("shen.initialise-lambda-forms")
var symshen_4loop = MakeSymbol("shen.loop")
var symshen_4free__variable__warnings = MakeSymbol("shen.free_variable_warnings")
var symshen_4_5number_6 = MakeSymbol("shen.<number>")
var symshen_4complexity__head = MakeSymbol("shen.complexity_head")
var symshen_4type_1signature_1of_1adjoin = MakeSymbol("shen.type-signature-of-adjoin")
var symshen_4app = MakeSymbol("shen.app")
var symshen_4sysfunc_2 = MakeSymbol("shen.sysfunc?")
var symshen_4type_1signature_1of_1stoutput = MakeSymbol("shen.type-signature-of-stoutput")
var symnl = MakeSymbol("nl")
var symshen_4parameters = MakeSymbol("shen.parameters")
var symshen_4x_4factorise_1defun_4true_1branch = MakeSymbol("shen.x.factorise-defun.true-branch")
var symshen_4type_1signature_1of_1_a = MakeSymbol("shen.type-signature-of-=")
var symshen_4type_1signature_1of_1_d = MakeSymbol("shen.type-signature-of-*")
var symshen_4type_1signature_1of_1cd = MakeSymbol("shen.type-signature-of-cd")
var symshen_4update_1demodulation_1function = MakeSymbol("shen.update-demodulation-function")
var symtrack = MakeSymbol("track")
var symshen_4carriage_1return = MakeSymbol("shen.carriage-return")
var symshen_4x_4factorise_1defun_4optimize_1selectors = MakeSymbol("shen.x.factorise-defun.optimize-selectors")
var symcn = MakeSymbol("cn")
var symtc = MakeSymbol("tc")
var symshen_4eval_1cons = MakeSymbol("shen.eval-cons")
var symshen_4put_cget_1macro = MakeSymbol("shen.put/get-macro")
var symshen_4type_1signature_1of_1append = MakeSymbol("shen.type-signature-of-append")
var sympreclude = MakeSymbol("preclude")
var symshen_4type_1signature_1of_1include_1all_1but = MakeSymbol("shen.type-signature-of-include-all-but")
var symshen_4type_1signature_1of_1os = MakeSymbol("shen.type-signature-of-os")
var symshen_4type_1signature_1of_1string_1_6n = MakeSymbol("shen.type-signature-of-string->n")
var symshen_4typecheck = MakeSymbol("shen.typecheck")
var symread_1file_1as_1string = MakeSymbol("read-file-as-string")
var symshen_4_5digit_6 = MakeSymbol("shen.<digit>")
var symshen_4construct_1search_1clause = MakeSymbol("shen.construct-search-clause")
var symshen_4load_1help = MakeSymbol("shen.load-help")
var symshen_4ue_1sig = MakeSymbol("shen.ue-sig")
var symgensym = MakeSymbol("gensym")
var symshen_4split__cc__rules = MakeSymbol("shen.split_cc_rules")
var symshen_4intprolog_1help_1help = MakeSymbol("shen.intprolog-help-help")
var symshen_4type_1signature_1of_1element_2 = MakeSymbol("shen.type-signature-of-element?")
var symshen_4hdtl = MakeSymbol("shen.hdtl")
var symshen_4packaged_2 = MakeSymbol("shen.packaged?")
var symshen_4aum = MakeSymbol("shen.aum")
var symshen_4x_4factorise_1defun_4generate_1label = MakeSymbol("shen.x.factorise-defun.generate-label")
var symshen_4make_1key = MakeSymbol("shen.make-key")
var symshen_4embed_1and = MakeSymbol("shen.embed-and")
var symread_1from_1string = MakeSymbol("read-from-string")
var symshen_4_5postdigits_6 = MakeSymbol("shen.<postdigits>")
var symeval = MakeSymbol("eval")
var symshen_4nest_1disjunct = MakeSymbol("shen.nest-disjunct")
var symshen_4incinfs = MakeSymbol("shen.incinfs")
var symshen_4record_1source = MakeSymbol("shen.record-source")
var symshen_4safe_1product = MakeSymbol("shen.safe-product")
var symshen_4remove_1nth = MakeSymbol("shen.remove-nth")
var symshen_4_5sig_7rules_6 = MakeSymbol("shen.<sig+rules>")
var symshen_4type_1signature_1of_1_5 = MakeSymbol("shen.type-signature-of-<")
var symshen_4digit_2 = MakeSymbol("shen.digit?")
var symshen_4preclude_1h = MakeSymbol("shen.preclude-h")
var symshen_4x_4factorise_1defun_4exp_1var = MakeSymbol("shen.x.factorise-defun.exp-var")
var symappend = MakeSymbol("append")
var symshen_4type_1signature_1of_1profile_1results = MakeSymbol("shen.type-signature-of-profile-results")
var symshen_4_5whitespace_6 = MakeSymbol("shen.<whitespace>")
var symshen_4post = MakeSymbol("shen.post")
var symshen_4type_1signature_1of_1string_1_6symbol = MakeSymbol("shen.type-signature-of-string->symbol")
var symshen_4type_1signature_1of_1_c = MakeSymbol("shen.type-signature-of-/")
var symshen_4occurs_2 = MakeSymbol("shen.occurs?")
var symshen_4pretty_1type = MakeSymbol("shen.pretty-type")
var symshen_4package_1macro = MakeSymbol("shen.package-macro")
var symshen_4type_1signature_1of_1shen_4proc_1nl = MakeSymbol("shen.type-signature-of-shen.proc-nl")
var symtuple_2 = MakeSymbol("tuple?")
var sym_8p = MakeSymbol("@p")
var symmacroexpand = MakeSymbol("macroexpand")
var symshen_4profile_1help = MakeSymbol("shen.profile-help")
var symshen_4assoc_1macro = MakeSymbol("shen.assoc-macro")
var symshen_4cases_1macro = MakeSymbol("shen.cases-macro")
var symlength = MakeSymbol("length")
var symshen_4x_4factorise_1defun_4free_1variables = MakeSymbol("shen.x.factorise-defun.free-variables")
var symshen_4type_1signature_1of_1_5_1vector = MakeSymbol("shen.type-signature-of-<-vector")
var symshen_4_5expr_6 = MakeSymbol("shen.<expr>")
var symshen_4record_1exceptions = MakeSymbol("shen.record-exceptions")
var symshen_4pause_1for_1user = MakeSymbol("shen.pause-for-user")
var symshen_4type_1signature_1of_1pr = MakeSymbol("shen.type-signature-of-pr")
var symshen_4hat = MakeSymbol("shen.hat")
var symshen_4x_4launcher_4eval_1command = MakeSymbol("shen.x.launcher.eval-command")
var symshen_4_5st__input1_6 = MakeSymbol("shen.<st_input1>")
var symshen_4right_1_6left = MakeSymbol("shen.right->left")
var symshen_4ebr = MakeSymbol("shen.ebr")
var symshen_4_5action_6 = MakeSymbol("shen.<action>")
var symboolean_2 = MakeSymbol("boolean?")
var symshen_4list_2 = MakeSymbol("shen.list?")
var symshen_4type_1signature_1of_1_a_a = MakeSymbol("shen.type-signature-of-==")
var symshen_4sequent = MakeSymbol("shen.sequent")
var symbind = MakeSymbol("bind")
var symshen_4type_1signature_1of_1external = MakeSymbol("shen.type-signature-of-external")
var symshen_4type_1signature_1of_1map = MakeSymbol("shen.type-signature-of-map")
var symshen_4cons__form = MakeSymbol("shen.cons_form")
var symshen_4prh = MakeSymbol("shen.prh")
var symshen_4clause__form = MakeSymbol("shen.clause_form")
var symshen_4type_1signature_1of_1spy = MakeSymbol("shen.type-signature-of-spy")
var symshen_4type_1signature_1of_1sum = MakeSymbol("shen.type-signature-of-sum")
var symmap = MakeSymbol("map")
var symshen_4s_1prolog = MakeSymbol("shen.s-prolog")
var symshen_4cc__body = MakeSymbol("shen.cc_body")
var symshen_4type_1signature_1of_1preclude = MakeSymbol("shen.type-signature-of-preclude")
var symshen_4type_1signature_1of_1tc = MakeSymbol("shen.type-signature-of-tc")
var symshen_4read_1file_1as_1charlist = MakeSymbol("shen.read-file-as-charlist")
var symshen_4_5strc_6 = MakeSymbol("shen.<strc>")
var symload = MakeSymbol("load")
var syminternal = MakeSymbol("internal")
var symsnd = MakeSymbol("snd")
var symtlstr = MakeSymbol("tlstr")
var symshen_4mkstr_1r = MakeSymbol("shen.mkstr-r")
var symshen_4x_4launcher_4eval_1command_1h = MakeSymbol("shen.x.launcher.eval-command-h")
var symshen_4syntax = MakeSymbol("shen.syntax")
var symshen_4mkstr_1l = MakeSymbol("shen.mkstr-l")
var symshen_4type_1signature_1of_1implementation = MakeSymbol("shen.type-signature-of-implementation")
var symdeclare = MakeSymbol("declare")
var symshen_4length_1h = MakeSymbol("shen.length-h")
var symshen_4locally_1bind_1prolog_1vs = MakeSymbol("shen.locally-bind-prolog-vs")
var symshen_4type_1signature_1of_1or = MakeSymbol("shen.type-signature-of-or")
var symintersection = MakeSymbol("intersection")
var symshen_4_5predigits_6 = MakeSymbol("shen.<predigits>")
var symshen_4double_1_6singles = MakeSymbol("shen.double->singles")
var symshen_4_5E_6 = MakeSymbol("shen.<E>")
var symshen_4lambda_1form = MakeSymbol("shen.lambda-form")
var symshen_4multiples = MakeSymbol("shen.multiples")
var symshen_4sh_2 = MakeSymbol("shen.sh?")
var symshen_4pre = MakeSymbol("shen.pre")
var symshen_4function_1abstraction_1help = MakeSymbol("shen.function-abstraction-help")
var symshen_4type_1signature_1of_1_1 = MakeSymbol("shen.type-signature-of--")
var symshen_4type_1signature_1of_1language = MakeSymbol("shen.type-signature-of-language")
var symshen_4space = MakeSymbol("shen.space")
var symshen_4find_1past_1inputs = MakeSymbol("shen.find-past-inputs")
var symstinput = MakeSymbol("stinput")
var symshen_4next_150 = MakeSymbol("shen.next-50")
var symshen_4_8v_1help = MakeSymbol("shen.@v-help")
var symshen_4_5singleunderline_6 = MakeSymbol("shen.<singleunderline>")
var symshen_4toplevel_1display_1exception = MakeSymbol("shen.toplevel-display-exception")
var symtl = MakeSymbol("tl")
var symshen_4type_1signature_1of_1vector_2 = MakeSymbol("shen.type-signature-of-vector?")
var symspy = MakeSymbol("spy")
var symshen_4rules_1_6horn_1clauses = MakeSymbol("shen.rules->horn-clauses")
var symshen_4_5semicolon_1symbol_6 = MakeSymbol("shen.<semicolon-symbol>")
var symshen_4read_1file_1as_1Xlist = MakeSymbol("shen.read-file-as-Xlist")
var symshen_4_5multiline_6 = MakeSymbol("shen.<multiline>")
var symshen_4type_1signature_1of_1limit = MakeSymbol("shen.type-signature-of-limit")
var symshen_4type_1signature_1of_1simple_1error = MakeSymbol("shen.type-signature-of-simple-error")
var symshen_4type_1signature_1of_1write_1to_1file = MakeSymbol("shen.type-signature-of-write-to-file")
var symshen_4reduce = MakeSymbol("shen.reduce")
var symshen_4proc_1nl = MakeSymbol("shen.proc-nl")
var symshen_4prolog_1aritycheck = MakeSymbol("shen.prolog-aritycheck")
var symshen_4call__the__continuation = MakeSymbol("shen.call_the_continuation")
var symshen_4type_1signature_1of_1shen_4insert = MakeSymbol("shen.type-signature-of-shen.insert")
var symshen_4_5equal_6 = MakeSymbol("shen.<equal>")
var symshen_4left_1rule = MakeSymbol("shen.left-rule")
var symshen_4profile_1func = MakeSymbol("shen.profile-func")
var symbound_2 = MakeSymbol("bound?")
var symshen_4type_1signature_1of_1destroy = MakeSymbol("shen.type-signature-of-destroy")
var symunspecialise = MakeSymbol("unspecialise")
var symshen_4for_1each = MakeSymbol("shen.for-each")
var symshen_4rule_1_6horn_1clause = MakeSymbol("shen.rule->horn-clause")
var symshen_4insert_1prolog_1variables_1help = MakeSymbol("shen.insert-prolog-variables-help")
var symshen_4copy_1vector = MakeSymbol("shen.copy-vector")
var symshen_4fbound_2 = MakeSymbol("shen.fbound?")
var symshen_4symbol_1code_2 = MakeSymbol("shen.symbol-code?")
var symshen_4_5semicolon_6 = MakeSymbol("shen.<semicolon>")
var symshen_4insert_1predicate = MakeSymbol("shen.insert-predicate")
var symremove = MakeSymbol("remove")
var symshen_4call_1help = MakeSymbol("shen.call-help")
var symshen_4udefs_d = MakeSymbol("shen.udefs*")
var syminclude = MakeSymbol("include")
var symshen_4cond_1expression = MakeSymbol("shen.cond-expression")
var symconcat = MakeSymbol("concat")
var symshen_4x_4factorise_1defun_4concat_c = MakeSymbol("shen.x.factorise-defun.concat/")
var sym_5_b_6 = MakeSymbol("<!>")
var symshen_4type_1signature_1of_1integer_2 = MakeSymbol("shen.type-signature-of-integer?")
var symshen_4remove_1bar = MakeSymbol("shen.remove-bar")
var symshen_4type_1signature_1of_1get_1time = MakeSymbol("shen.type-signature-of-get-time")
var symshen_4function_1abstraction = MakeSymbol("shen.function-abstraction")
var symshen_4_5head_d_6 = MakeSymbol("shen.<head*>")
var symshen_4s_1prolog__literal = MakeSymbol("shen.s-prolog_literal")
var symget_1time = MakeSymbol("get-time")
var symshen_4prolog_1macro = MakeSymbol("shen.prolog-macro")
var symshen_4type_1signature_1of_1union = MakeSymbol("shen.type-signature-of-union")
var symshen_4x_4factorise_1defun_4bind_1repeating_1selectors = MakeSymbol("shen.x.factorise-defun.bind-repeating-selectors")
var symshen_4list__variables = MakeSymbol("shen.list_variables")
var symshen_4construct_1context = MakeSymbol("shen.construct-context")
var symshen_4type_1signature_1of_1number_2 = MakeSymbol("shen.type-signature-of-number?")
var symmaxinferences = MakeSymbol("maxinferences")
var symshen_4_5plus_6 = MakeSymbol("shen.<plus>")
var symshen_4chwild = MakeSymbol("shen.chwild")
var symshen_4base = MakeSymbol("shen.base")
var symshen_4type_1signature_1of_1assoc = MakeSymbol("shen.type-signature-of-assoc")
var symshen_4type_1signature_1of_1intersection = MakeSymbol("shen.type-signature-of-intersection")
var symshen_4alphanum_2 = MakeSymbol("shen.alphanum?")
var symshen_4remove__modes = MakeSymbol("shen.remove_modes")
var symshen_4initialise_1signedfuncs = MakeSymbol("shen.initialise-signedfuncs")
var symshen_4extract__free__vars = MakeSymbol("shen.extract_free_vars")
var symshen_4jump__stream = MakeSymbol("shen.jump_stream")
var symshen_4yacc_1_6shen = MakeSymbol("shen.yacc->shen")
var symshen_4recursive__descent = MakeSymbol("shen.recursive_descent")
var symread_1file = MakeSymbol("read-file")
var symshen_4pvar = MakeSymbol("shen.pvar")
var symfail = MakeSymbol("fail")
var symshen_4aritycheck_1action = MakeSymbol("shen.aritycheck-action")
var symshen_4type_1signature_1of_1kill = MakeSymbol("shen.type-signature-of-kill")
var symshen_4type_1signature_1of_1occurs_1check = MakeSymbol("shen.type-signature-of-occurs-check")
var symshen_4tab = MakeSymbol("shen.tab")
var symsum = MakeSymbol("sum")
var symshen_4type_1signature_1of_1version = MakeSymbol("shen.type-signature-of-version")
var symimplementation = MakeSymbol("implementation")
var symshen_4rcons__form = MakeSymbol("shen.rcons_form")
var symshen_4print_1vector_2 = MakeSymbol("shen.print-vector?")
var symshen_4safe_1multiply = MakeSymbol("shen.safe-multiply")
var symshen_4typecheck_1and_1load = MakeSymbol("shen.typecheck-and-load")
var syminclude_1all_1but = MakeSymbol("include-all-but")
var symshen_4internal_1symbols = MakeSymbol("shen.internal-symbols")
var symshen_4s_1prolog__clause = MakeSymbol("shen.s-prolog_clause")
var symshen_4type_1signature_1of_1remove = MakeSymbol("shen.type-signature-of-remove")
var symhead = MakeSymbol("head")
var symshen_4insert_1deref = MakeSymbol("shen.insert-deref")
var symshen_4print_1past_1inputs = MakeSymbol("shen.print-past-inputs")
var symwrite_1byte = MakeSymbol("write-byte")
var symshen_4_5non_1ll_1rules_6 = MakeSymbol("shen.<non-ll-rules>")
var symshen_4variant_2 = MakeSymbol("shen.variant?")
var symshen_4type_1signature_1of_1str = MakeSymbol("shen.type-signature-of-str")
var symshen_4analyse_1symbol_2 = MakeSymbol("shen.analyse-symbol?")
var symshen_4_5rrb_6 = MakeSymbol("shen.<rrb>")
var syminteger_2 = MakeSymbol("integer?")
var symshen_4err_1condition = MakeSymbol("shen.err-condition")
var symshen_4special_2 = MakeSymbol("shen.special?")
var symshen_4prolog_1failure = MakeSymbol("shen.prolog-failure")
var symshen_4abstraction__build = MakeSymbol("shen.abstraction_build")
var symshen_4analyse_1variable_2 = MakeSymbol("shen.analyse-variable?")
var symshen_4type_1signature_1of_1length = MakeSymbol("shen.type-signature-of-length")
var symshen_4alphanums_2 = MakeSymbol("shen.alphanums?")
var symshen_4construct_1recursive_1search_1clause = MakeSymbol("shen.construct-recursive-search-clause")
var symshen_4call_1rest = MakeSymbol("shen.call-rest")
var symshen_4by__hypothesis = MakeSymbol("shen.by_hypothesis")
var symshen_4newhyps = MakeSymbol("shen.newhyps")
var symshen_4demodulate = MakeSymbol("shen.demodulate")
var symshen_4store_1arity = MakeSymbol("shen.store-arity")
var symshen_4get_1type = MakeSymbol("shen.get-type")
var sym_5_1vector = MakeSymbol("<-vector")
var symshen_4initialise__arity__table = MakeSymbol("shen.initialise_arity_table")
var symshen_4_8s_1macro = MakeSymbol("shen.@s-macro")
var symcut = MakeSymbol("cut")
var symshen_4type_1signature_1of_1variable_2 = MakeSymbol("shen.type-signature-of-variable?")
var symdefun = MakeSymbol("defun")
var symshen_4_5rule_6 = MakeSymbol("shen.<rule>")
var symshen_4procedure__name = MakeSymbol("shen.procedure_name")
var symshen_4mult__subst = MakeSymbol("shen.mult_subst")
var symshen_4x_4factorise_1defun_4attach_1free_1variables = MakeSymbol("shen.x.factorise-defun.attach-free-variables")
var symshen_4set_1lambda_1form_1entry = MakeSymbol("shen.set-lambda-form-entry")
var symshen_4type_1signature_1of_1write_1byte = MakeSymbol("shen.type-signature-of-write-byte")
var symshen_4decons = MakeSymbol("shen.decons")
var symshen_4_5sym_6 = MakeSymbol("shen.<sym>")
var symrelease = MakeSymbol("release")
var symshen_4fix_1help = MakeSymbol("shen.fix-help")
var symshen_4terminal_2 = MakeSymbol("shen.terminal?")
var symshen_4read_1loop = MakeSymbol("shen.read-loop")
var symshen_4check_1byte = MakeSymbol("shen.check-byte")
var symshen_4byte_1_6digit = MakeSymbol("shen.byte->digit")
var symshen_4_5return_6 = MakeSymbol("shen.<return>")
var symshen_4type_1signature_1of_1print = MakeSymbol("shen.type-signature-of-print")
var symshen_4bucket_1fold = MakeSymbol("shen.bucket-fold")
var symshen_4_5bar_6 = MakeSymbol("shen.<bar>")
var symshen_4magless = MakeSymbol("shen.magless")
var symshen_4jump__stream_2 = MakeSymbol("shen.jump_stream?")
var symshen_4_5rsb_6 = MakeSymbol("shen.<rsb>")
var symshen_4_5predicate_d_6 = MakeSymbol("shen.<predicate*>")
var symshen_4head__abstraction = MakeSymbol("shen.head_abstraction")
var symversion = MakeSymbol("version")
var symshen_4update__history = MakeSymbol("shen.update_history")
var symshen_4linearise__X = MakeSymbol("shen.linearise_X")
var symshen_4output_1macro = MakeSymbol("shen.output-macro")
var symshen_4type_1signature_1of_1_5e_6 = MakeSymbol("shen.type-signature-of-<e>")
var symshen_4abstract__rule = MakeSymbol("shen.abstract_rule")
var symshen_4encode_1choices = MakeSymbol("shen.encode-choices")
var symshen_4copy_1vector_1stage_12 = MakeSymbol("shen.copy-vector-stage-2")
var symshen_4type_1signature_1of_1_6_a = MakeSymbol("shen.type-signature-of->=")
var symshen_4_5premise_6 = MakeSymbol("shen.<premise>")
var symshen_4type_1signature_1of_1tail = MakeSymbol("shen.type-signature-of-tail")
var symshen_4linearise__help = MakeSymbol("shen.linearise_help")
var symshen_4_5times_6 = MakeSymbol("shen.<times>")
var symshen_4datatype_1macro = MakeSymbol("shen.datatype-macro")
var symshen_4default_1rule = MakeSymbol("shen.default-rule")
var symshen_4_5whitespaces_6 = MakeSymbol("shen.<whitespaces>")
var symshen_4type_1signature_1of_1nl = MakeSymbol("shen.type-signature-of-nl")
var symshen_4type_1signature_1of_1_5_a = MakeSymbol("shen.type-signature-of-<=")
var symspecialise = MakeSymbol("specialise")
var symshen_4packageh = MakeSymbol("shen.packageh")
var symshen_4_5lsb_6 = MakeSymbol("shen.<lsb>")
var symshen_4extraspecial_2 = MakeSymbol("shen.extraspecial?")
var symshen_4ue_2 = MakeSymbol("shen.ue?")
var symshen_4type_1signature_1of_1error_1to_1string = MakeSymbol("shen.type-signature-of-error-to-string")
var symshen_4_5signature_6 = MakeSymbol("shen.<signature>")
var symmapcan = MakeSymbol("mapcan")
var symshen_4atom_1_6str = MakeSymbol("shen.atom->str")
var symprint = MakeSymbol("print")
var symshen_4lzy_a_b = MakeSymbol("shen.lzy=!")
var symshen_4curry_1synonyms = MakeSymbol("shen.curry-synonyms")
var symshen_4type_1signature_1of_1explode = MakeSymbol("shen.type-signature-of-explode")
var symshen_4type_1signature_1of_1sterror = MakeSymbol("shen.type-signature-of-sterror")
var symshen_4prompt = MakeSymbol("shen.prompt")
var symshen_4aritycheck = MakeSymbol("shen.aritycheck")
var symshen_4type_1signature_1of_1vector = MakeSymbol("shen.type-signature-of-vector")
var symshen_4type_1signature_1of_1_7 = MakeSymbol("shen.type-signature-of-+")
var symshen_4x_4factorise_1defun_4factorise_1cond = MakeSymbol("shen.x.factorise-defun.factorise-cond")
var symshen_4_5formulae_6 = MakeSymbol("shen.<formulae>")
var symshen_4unbindv = MakeSymbol("shen.unbindv")
var symshen_4patthyps = MakeSymbol("shen.patthyps")
var symshen_4read_1evaluate_1print = MakeSymbol("shen.read-evaluate-print")
var symshen_4_5st__input2_6 = MakeSymbol("shen.<st_input2>")
var symvector_2 = MakeSymbol("vector?")
var symshen_4semantic_1completion_1warning = MakeSymbol("shen.semantic-completion-warning")
var symshen_4digits_1_6integers = MakeSymbol("shen.digits->integers")
var symshen_4prefix_2 = MakeSymbol("shen.prefix?")
var symunion = MakeSymbol("union")
var symshen_4findpos = MakeSymbol("shen.findpos")
var symshen_4maxinfexceeded_2 = MakeSymbol("shen.maxinfexceeded?")
var symidentical = MakeSymbol("identical")
var symshen_4t_d_1patterns = MakeSymbol("shen.t*-patterns")
var symshen_4remember = MakeSymbol("shen.remember")
var symshen_4clash_2 = MakeSymbol("shen.clash?")
var symshen_4_5variable_2_6 = MakeSymbol("shen.<variable?>")
var symshen_4_5literal_d_6 = MakeSymbol("shen.<literal*>")
var symabsvector = MakeSymbol("absvector")
var sym_6_a = MakeSymbol(">=")
var symshen_4_5term_d_6 = MakeSymbol("shen.<term*>")
var symshen_4type_1signature_1of_1porters = MakeSymbol("shen.type-signature-of-porters")
var symshen_4type_1signature_1of_1snd = MakeSymbol("shen.type-signature-of-snd")
var symshen_4toplevel = MakeSymbol("shen.toplevel")
var symshen_4percent = MakeSymbol("shen.percent")
var symshen_4assoc_1set = MakeSymbol("shen.assoc-set")
var symstring_2 = MakeSymbol("string?")
var symshen_4multiple_1set = MakeSymbol("shen.multiple-set")
var symcompile = MakeSymbol("compile")
var symshen_4initialise_1environment = MakeSymbol("shen.initialise-environment")
var symshen_4toplineread__loop = MakeSymbol("shen.toplineread_loop")
var symshen_4_5singleline_6 = MakeSymbol("shen.<singleline>")
var symshen_4type_1signature_1of_1read_1byte = MakeSymbol("shen.type-signature-of-read-byte")
var symshen_4kill_1code = MakeSymbol("shen.kill-code")
var symshen_4mu__reduction = MakeSymbol("shen.mu_reduction")
var symshen_4em__help = MakeSymbol("shen.em_help")
var symshen_4type_1signature_1of_1internal = MakeSymbol("shen.type-signature-of-internal")
var symlineread = MakeSymbol("lineread")
var symerror_1to_1string = MakeSymbol("error-to-string")
var symshen_4copyfromvector = MakeSymbol("shen.copyfromvector")
var symshen_4_5pattern_6 = MakeSymbol("shen.<pattern>")
var symshen_4_5lrb_6 = MakeSymbol("shen.<lrb>")
var symcons = MakeSymbol("cons")
var symshen_4x_4factorise_1defun_4rebranch_1h = MakeSymbol("shen.x.factorise-defun.rebranch-h")
var symshen_4show_1assumptions = MakeSymbol("shen.show-assumptions")
var symshen_4timer_1macro = MakeSymbol("shen.timer-macro")
var symshen_4type_1signature_1of_1tuple_2 = MakeSymbol("shen.type-signature-of-tuple?")
var symshen_4dict_1update_1count = MakeSymbol("shen.dict-update-count")
var symstr = MakeSymbol("str")
var symshen_4toplineread = MakeSymbol("shen.toplineread")
var symshen_4x_4launcher_4help_1text = MakeSymbol("shen.x.launcher.help-text")
var symshen_4insert_1h = MakeSymbol("shen.insert-h")
var symshen_4mkstr = MakeSymbol("shen.mkstr")
var symshen_4t_d_1rules = MakeSymbol("shen.t*-rules")
var symshen_4x_4factorise_1defun_4apply_1selector_1handlers = MakeSymbol("shen.x.factorise-defun.apply-selector-handlers")
var symget = MakeSymbol("get")
var symshen_4semantics = MakeSymbol("shen.semantics")
var symshen_4pushnew = MakeSymbol("shen.pushnew")
var symshen_4_5comma_6 = MakeSymbol("shen.<comma>")
var symshen_4t_d_1def = MakeSymbol("shen.t*-def")
var symshen_4type_1signature_1of_1tlv = MakeSymbol("shen.type-signature-of-tlv")
var symshen_4walk = MakeSymbol("shen.walk")
var symstring_1_6n = MakeSymbol("string->n")
var symshen_4record_1internal = MakeSymbol("shen.record-internal")
var symshen_4type_1signature_1of_1it = MakeSymbol("shen.type-signature-of-it")
var symintern = MakeSymbol("intern")
var symshen_4split__cc__rule = MakeSymbol("shen.split_cc_rule")
var symshen_4_5atom_6 = MakeSymbol("shen.<atom>")
var symshen_4iter_1list = MakeSymbol("shen.iter-list")
var symshen_4_5clause_d_6 = MakeSymbol("shen.<clause*>")
var symshen_4type_1signature_1of_1_6 = MakeSymbol("shen.type-signature-of->")
var symdifference = MakeSymbol("difference")
var symy_1or_1n_2 = MakeSymbol("y-or-n?")
var symshen_4type_1signature_1of_1reverse = MakeSymbol("shen.type-signature-of-reverse")
var symprofile = MakeSymbol("profile")
var symread = MakeSymbol("read")
var symshen_4cn_1all = MakeSymbol("shen.cn-all")
var symshen_4_5define_6 = MakeSymbol("shen.<define>")
var symshen_4lambda_1of_1defun = MakeSymbol("shen.lambda-of-defun")
var symshen_4legitimate_1term_2 = MakeSymbol("shen.legitimate-term?")
var symshen_4make_1string_1macro = MakeSymbol("shen.make-string-macro")
var symshen_4dict_1count_1_6 = MakeSymbol("shen.dict-count->")
var symshen_4dict_1bucket_1_6 = MakeSymbol("shen.dict-bucket->")
var symshen_4insert = MakeSymbol("shen.insert")
var symshen_4type_1signature_1of_1fail = MakeSymbol("shen.type-signature-of-fail")
var symshen_4update_1symbol_1table = MakeSymbol("shen.update-symbol-table")
var symshen_4_5lcurly_6 = MakeSymbol("shen.<lcurly>")
var sym_c = MakeSymbol("/")
var symshen_4list_1_6str = MakeSymbol("shen.list->str")
var symabort = MakeSymbol("abort")
var symshen_4shen_1syntax_1error = MakeSymbol("shen.shen-syntax-error")
var symshen_4str_1_6str = MakeSymbol("shen.str->str")
var symshen_4construct_1search_1literals = MakeSymbol("shen.construct-search-literals")
var symshen_4_5num_6 = MakeSymbol("shen.<num>")
var symshen_4trim_1whitespace = MakeSymbol("shen.trim-whitespace")
var symshen_4type_1signature_1of_1read_1file_1as_1string = MakeSymbol("shen.type-signature-of-read-file-as-string")
var symshen_4_5alpha_6 = MakeSymbol("shen.<alpha>")
var symshen_4collect = MakeSymbol("shen.collect")
var syminput = MakeSymbol("input")
var symshen_4type_1signature_1of_1close = MakeSymbol("shen.type-signature-of-close")
var symshen_4eval_1without_1macros = MakeSymbol("shen.eval-without-macros")
var symshen_4strip_1protect = MakeSymbol("shen.strip-protect")
var symshen_4after_1codepoint = MakeSymbol("shen.after-codepoint")
var symshen_4_5strcontents_6 = MakeSymbol("shen.<strcontents>")
var symshen_4read_1char_1code = MakeSymbol("shen.read-char-code")
var symshen_4type_1signature_1of_1untrack = MakeSymbol("shen.type-signature-of-untrack")
var symread_1file_1as_1bytelist = MakeSymbol("read-file-as-bytelist")
var symshen_4dictionary = MakeSymbol("shen.dictionary")
var symshen_4type_1signature_1of_1maxinferences = MakeSymbol("shen.type-signature-of-maxinferences")
var symunput = MakeSymbol("unput")
var symshen_4lineread_1loop = MakeSymbol("shen.lineread-loop")
var symshen_4remtype = MakeSymbol("shen.remtype")
var symshen_4_5guard_6 = MakeSymbol("shen.<guard>")
var symshen_4type_1signature_1of_1optimise = MakeSymbol("shen.type-signature-of-optimise")
var symshen_4tracked_2 = MakeSymbol("shen.tracked?")
var symshen_4vector_1_6str = MakeSymbol("shen.vector->str")
var symfst = MakeSymbol("fst")
var symshen_4yacc__cases = MakeSymbol("shen.yacc_cases")
var symshen_4input_1macro = MakeSymbol("shen.input-macro")
var symshen_4type_1signature_1of_1read = MakeSymbol("shen.type-signature-of-read")
var symshen_4_5comment_6 = MakeSymbol("shen.<comment>")
var symshen_4record_1it_1h = MakeSymbol("shen.record-it-h")
var symshen_4ues = MakeSymbol("shen.ues")
var symshen_4mod = MakeSymbol("shen.mod")
var symshen_4x_4launcher_4script_1command = MakeSymbol("shen.x.launcher.script-command")
var symshen_4explicit__modes = MakeSymbol("shen.explicit_modes")
var symshen_4type_1signature_1of_1hash = MakeSymbol("shen.type-signature-of-hash")
var symsymbol_2 = MakeSymbol("symbol?")
var symshen_4construct_1search_1clauses = MakeSymbol("shen.construct-search-clauses")
var symshen_4elim_1def = MakeSymbol("shen.elim-def")
var symshen_4include_1h = MakeSymbol("shen.include-h")
var symshen_4clauses_1to_1shen = MakeSymbol("shen.clauses-to-shen")
var symshen_4abs_1macro = MakeSymbol("shen.abs-macro")
var symshen_4type_1signature_1of_1absvector_2 = MakeSymbol("shen.type-signature-of-absvector?")
var symuntrack = MakeSymbol("untrack")
var symshen_4type_1signature_1of_1vector_1_6 = MakeSymbol("shen.type-signature-of-vector->")
var symshen_4insert__modes = MakeSymbol("shen.insert_modes")
var symunify = MakeSymbol("unify")
var symshen_4insert_1tracking_1code = MakeSymbol("shen.insert-tracking-code")
var symshen_4_5conclusion_6 = MakeSymbol("shen.<conclusion>")
var symshen_4intern_1type = MakeSymbol("shen.intern-type")
var symopen = MakeSymbol("open")
var symshen_4type_1signature_1of_1head = MakeSymbol("shen.type-signature-of-head")
var sym_8v = MakeSymbol("@v")
var symshen_4remember_1datatype = MakeSymbol("shen.remember-datatype")
var symshen_4complexity = MakeSymbol("shen.complexity")
var symshen_4externally_1bound = MakeSymbol("shen.externally-bound")
var symshen_4type_1signature_1of_1preclude_1all_1but = MakeSymbol("shen.type-signature-of-preclude-all-but")
var symshen_4prhush = MakeSymbol("shen.prhush")
var sym_8s = MakeSymbol("@s")
var symshen_4_5end_d_6 = MakeSymbol("shen.<end*>")
var symshen_4_5patterns_6 = MakeSymbol("shen.<patterns>")
var symshen_4_5_1dict = MakeSymbol("shen.<-dict")
var symshen_4type_1signature_1of_1undefmacro = MakeSymbol("shen.type-signature-of-undefmacro")
var symstep = MakeSymbol("step")
var symshen_4type_1signature_1of_1pos = MakeSymbol("shen.type-signature-of-pos")
var symshen_4type_1signature_1of_1read_1file = MakeSymbol("shen.type-signature-of-read-file")
var symhash = MakeSymbol("hash")
var symshen_4insert_1lazyderef = MakeSymbol("shen.insert-lazyderef")
var symshen_4right_1rule = MakeSymbol("shen.right-rule")
var symshen_4insert_1l = MakeSymbol("shen.insert-l")
var symshen_4type_1signature_1of_1protect = MakeSymbol("shen.type-signature-of-protect")
var sympackage_2 = MakeSymbol("package?")
var symshen_4_5signature_1help_6 = MakeSymbol("shen.<signature-help>")
var symshen_4x_4factorise_1defun_4add_1returns = MakeSymbol("shen.x.factorise-defun.add-returns")
var symshen_4construct_1side_1literals = MakeSymbol("shen.construct-side-literals")
var symcall = MakeSymbol("call")
var symshen_4type_1signature_1of_1package_2 = MakeSymbol("shen.type-signature-of-package?")
var symshen_4exclamation = MakeSymbol("shen.exclamation")
var symshen_4type_1signature_1of_1boolean_2 = MakeSymbol("shen.type-signature-of-boolean?")
var symshen_4compile__prolog__procedure = MakeSymbol("shen.compile_prolog_procedure")
var symshen_4lisp_1or = MakeSymbol("shen.lisp-or")
var symshen_4placeholder = MakeSymbol("shen.placeholder")
var symshen_4typecheck_1and_1evaluate = MakeSymbol("shen.typecheck-and-evaluate")
var symexplode = MakeSymbol("explode")
var symshen_4prolog_1_6shen = MakeSymbol("shen.prolog->shen")
var symshen_4decons_1string = MakeSymbol("shen.decons-string")
var symshen_4_5name_6 = MakeSymbol("shen.<name>")
var symshen_4newline = MakeSymbol("shen.newline")
var symshen_4_5type_6 = MakeSymbol("shen.<type>")
var symshen_4type_1signature_1of_1shen_4app = MakeSymbol("shen.type-signature-of-shen.app")
var symshen_4lzy_a_a = MakeSymbol("shen.lzy==")
var symshen_4result_1type = MakeSymbol("shen.result-type")
var symshen_4recursively_1print = MakeSymbol("shen.recursively-print")
var symcd = MakeSymbol("cd")
var symshen_4retrieve_1from_1history_1if_1needed = MakeSymbol("shen.retrieve-from-history-if-needed")
var symshen_4grammar__symbol_2 = MakeSymbol("shen.grammar_symbol?")
var symshen_4code_1point = MakeSymbol("shen.code-point")
var symunify_b = MakeSymbol("unify!")
var symshen_4type_1signature_1of_1compile = MakeSymbol("shen.type-signature-of-compile")
var symshen_4initialise_1signedfunc_1lambda_1forms = MakeSymbol("shen.initialise-signedfunc-lambda-forms")
var symshen_4lzy_a = MakeSymbol("shen.lzy=")
var symlimit = MakeSymbol("limit")
var symshen_4monotype = MakeSymbol("shen.monotype")
var symshen_4trim_1gubbins = MakeSymbol("shen.trim-gubbins")
var symshen_4reverse__help = MakeSymbol("shen.reverse_help")
var symshen_4control_1chars = MakeSymbol("shen.control-chars")
var symshen_4constructor_1error = MakeSymbol("shen.constructor-error")
var symshen_4linearise_1clause = MakeSymbol("shen.linearise-clause")
var symreturn = MakeSymbol("return")
var symshen_4aritycheck_1name = MakeSymbol("shen.aritycheck-name")
var symshen_4type_1signature_1of_1read_1file_1as_1bytelist = MakeSymbol("shen.type-signature-of-read-file-as-bytelist")
var symshen_4insert_1prolog_1variables = MakeSymbol("shen.insert-prolog-variables")
var symshen_4catchpoint = MakeSymbol("shen.catchpoint")
var symvector_1_6 = MakeSymbol("vector->")
var symprofile_1results = MakeSymbol("profile-results")
var symshen_4show_1p = MakeSymbol("shen.show-p")
var symhdstr = MakeSymbol("hdstr")
var symshen_4rule_1_6horn_1clause_1body = MakeSymbol("shen.rule->horn-clause-body")
var symshen_4read_1file_1as_1Xlist_1help = MakeSymbol("shen.read-file-as-Xlist-help")
var symshen_4put_1profile = MakeSymbol("shen.put-profile")
var symshen_4synonyms_1macro = MakeSymbol("shen.synonyms-macro")
var symhd = MakeSymbol("hd")
var symshen_4csl_1help = MakeSymbol("shen.csl-help")
var symread_1byte = MakeSymbol("read-byte")
var symshen_4compress_150 = MakeSymbol("shen.compress-50")
var symshen_4construct_1premiss_1literal = MakeSymbol("shen.construct-premiss-literal")
var symshen_4group__clauses = MakeSymbol("shen.group_clauses")
var symshen_4newpv = MakeSymbol("shen.newpv")
var symfindall = MakeSymbol("findall")
var symshen_4dict_1capacity = MakeSymbol("shen.dict-capacity")
var symshen_4x_4factorise_1defun_4inline_1mono_1labels = MakeSymbol("shen.x.factorise-defun.inline-mono-labels")
var symshen_4aum__to__shen = MakeSymbol("shen.aum_to_shen")
var symtype = MakeSymbol("type")
var syminput_7 = MakeSymbol("input+")
var symabsvector_2 = MakeSymbol("absvector?")
var symshen_4x_4features_4current = MakeSymbol("shen.x.features.current")
var symshen_4_5log10_6 = MakeSymbol("shen.<log10>")
var symshen_4type_1signature_1of_1fix = MakeSymbol("shen.type-signature-of-fix")
var symshen_4x_4launcher_4version_1string = MakeSymbol("shen.x.launcher.version-string")
var symshen_4_5pattern2_6 = MakeSymbol("shen.<pattern2>")
var symfunction = MakeSymbol("function")
var symshen_4catch_1cut = MakeSymbol("shen.catch-cut")
var symshen_4let_1macro = MakeSymbol("shen.let-macro")
var symshen_4ue_1h_2 = MakeSymbol("shen.ue-h?")
var symshen_4strip_1pathname = MakeSymbol("shen.strip-pathname")
var symshen_4type_1signature_1of_1n_1_6string = MakeSymbol("shen.type-signature-of-n->string")
var symshen_4unwind_1types = MakeSymbol("shen.unwind-types")
var symshen_4arg_1_6str = MakeSymbol("shen.arg->str")
var symaddress_1_6 = MakeSymbol("address->")
var symshen_4remove_1h = MakeSymbol("shen.remove-h")
var symfix = MakeSymbol("fix")
