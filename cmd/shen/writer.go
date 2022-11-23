package main

import . "github.com/tiancaiamao/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
	tmp1941 := MakeNative(func(__e *ControlFlow) {
		V4723 := __e.Get(1)
		_ = V4723
		tmp1942 := MakeNative(func(__e *ControlFlow) {
			String := __e.Get(1)
			_ = String
			tmp1943 := MakeNative(func(__e *ControlFlow) {
				Print := __e.Get(1)
				_ = Print
				__e.Return(V4723)
				return
			}, 1)

			tmp1944 := Call(__e, PrimFunc(symstoutput))

			tmp1945 := Call(__e, PrimFunc(sympr), String, tmp1944)

			__e.TailApply(tmp1943, tmp1945)
			return

		}, 1)

		tmp1946 := Call(__e, PrimFunc(symshen_4insert), V4723, MakeString("~S"))

		__e.TailApply(tmp1942, tmp1946)
		return

	}, 1)

	tmp1947 := Call(__e, ns2_1set, symprint, tmp1941)

	_ = tmp1947

	tmp1948 := MakeNative(func(__e *ControlFlow) {
		V4724 := __e.Get(1)
		_ = V4724
		V4725 := __e.Get(2)
		_ = V4725
		tmp1953 := PrimValue(sym_dhush_d)

		if True == tmp1953 {
			__e.Return(V4724)
			return
		} else {
			tmp1951 := Call(__e, PrimFunc(symshen_4char_1stoutput_2), V4725)

			if True == tmp1951 {
				__e.TailApply(PrimFunc(symshen_4write_1string), V4724, V4725)
				return
			} else {
				tmp1949 := Call(__e, PrimFunc(symshen_4string_1_6byte), V4724, MakeNumber(0))

				__e.TailApply(PrimFunc(symshen_4write_1chars), V4724, V4725, tmp1949, MakeNumber(1))
				return

			}

		}

	}, 2)

	tmp1954 := Call(__e, ns2_1set, sympr, tmp1948)

	_ = tmp1954

	tmp1955 := MakeNative(func(__e *ControlFlow) {
		V4726 := __e.Get(1)
		_ = V4726
		V4727 := __e.Get(2)
		_ = V4727
		tmp1956 := MakeNative(func(__e *ControlFlow) {
			tmp1957 := PrimPos(V4726, V4727)

			__e.Return(PrimStringToNumber(tmp1957))
			return

		}, 0)

		tmp1958 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.Return(symshen_4eos)
			return
		}, 1)

		__e.TailApply(try_1catch, tmp1956, tmp1958)
		return

	}, 2)

	tmp1959 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp1955)

	_ = tmp1959

	tmp1960 := MakeNative(func(__e *ControlFlow) {
		V4728 := __e.Get(1)
		_ = V4728
		V4729 := __e.Get(2)
		_ = V4729
		V4730 := __e.Get(3)
		_ = V4730
		V4731 := __e.Get(4)
		_ = V4731
		tmp1965 := PrimEqual(symshen_4eos, V4730)

		if True == tmp1965 {
			__e.Return(V4728)
			return
		} else {
			tmp1961 := PrimWriteByte(V4730, V4729)

			_ = tmp1961

			tmp1962 := Call(__e, PrimFunc(symshen_4string_1_6byte), V4728, V4731)

			tmp1963 := PrimNumberAdd(V4731, MakeNumber(1))

			__e.TailApply(PrimFunc(symshen_4write_1chars), V4728, V4729, tmp1962, tmp1963)
			return

		}

	}, 4)

	tmp1966 := Call(__e, ns2_1set, symshen_4write_1chars, tmp1960)

	_ = tmp1966

	tmp1967 := MakeNative(func(__e *ControlFlow) {
		V4732 := __e.Get(1)
		_ = V4732
		V4733 := __e.Get(2)
		_ = V4733
		tmp1972 := PrimIsString(V4732)

		if True == tmp1972 {
			tmp1968 := Call(__e, PrimFunc(symshen_4proc_1nl), V4732)

			__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp1968, V4733)
			return

		} else {
			tmp1969 := PrimCons(V4732, Nil)

			tmp1970 := PrimCons(symshen_4proc_1nl, tmp1969)

			__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp1970, V4733)
			return

		}

	}, 2)

	tmp1973 := Call(__e, ns2_1set, symshen_4mkstr, tmp1967)

	_ = tmp1973

	tmp1974 := MakeNative(func(__e *ControlFlow) {
		V4738 := __e.Get(1)
		_ = V4738
		V4739 := __e.Get(2)
		_ = V4739
		tmp1981 := PrimEqual(Nil, V4739)

		if True == tmp1981 {
			__e.Return(V4738)
			return
		} else {
			tmp1979 := PrimIsPair(V4739)

			if True == tmp1979 {
				tmp1975 := PrimHead(V4739)

				tmp1976 := Call(__e, PrimFunc(symshen_4insert_1l), tmp1975, V4738)

				tmp1977 := PrimTail(V4739)

				__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp1976, tmp1977)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-l")))
				return
			}

		}

	}, 2)

	tmp1982 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp1974)

	_ = tmp1982

	tmp1983 := MakeNative(func(__e *ControlFlow) {
		V4746 := __e.Get(1)
		_ = V4746
		V4747 := __e.Get(2)
		_ = V4747
		tmp2121 := PrimEqual(MakeString(""), V4747)

		if True == tmp2121 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp2119 := Call(__e, PrimFunc(symshen_4_7string_2), V4747)

			var ifres2106 Obj

			if True == tmp2119 {
				tmp2117 := Call(__e, PrimFunc(symhdstr), V4747)

				tmp2118 := PrimEqual(MakeString("~"), tmp2117)

				var ifres2108 Obj

				if True == tmp2118 {
					tmp2115 := PrimTailString(V4747)

					tmp2116 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2115)

					var ifres2110 Obj

					if True == tmp2116 {
						tmp2112 := PrimTailString(V4747)

						tmp2113 := Call(__e, PrimFunc(symhdstr), tmp2112)

						tmp2114 := PrimEqual(MakeString("A"), tmp2113)

						var ifres2111 Obj

						if True == tmp2114 {
							ifres2111 = True

						} else {
							ifres2111 = False

						}

						ifres2110 = ifres2111

					} else {
						ifres2110 = False

					}

					var ifres2109 Obj

					if True == ifres2110 {
						ifres2109 = True

					} else {
						ifres2109 = False

					}

					ifres2108 = ifres2109

				} else {
					ifres2108 = False

				}

				var ifres2107 Obj

				if True == ifres2108 {
					ifres2107 = True

				} else {
					ifres2107 = False

				}

				ifres2106 = ifres2107

			} else {
				ifres2106 = False

			}

			if True == ifres2106 {
				tmp1984 := PrimTailString(V4747)

				tmp1985 := PrimTailString(tmp1984)

				tmp1986 := PrimCons(symshen_4a, Nil)

				tmp1987 := PrimCons(tmp1985, tmp1986)

				tmp1988 := PrimCons(V4746, tmp1987)

				__e.Return(PrimCons(symshen_4app, tmp1988))
				return

			} else {
				tmp2104 := Call(__e, PrimFunc(symshen_4_7string_2), V4747)

				var ifres2091 Obj

				if True == tmp2104 {
					tmp2102 := Call(__e, PrimFunc(symhdstr), V4747)

					tmp2103 := PrimEqual(MakeString("~"), tmp2102)

					var ifres2093 Obj

					if True == tmp2103 {
						tmp2100 := PrimTailString(V4747)

						tmp2101 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2100)

						var ifres2095 Obj

						if True == tmp2101 {
							tmp2097 := PrimTailString(V4747)

							tmp2098 := Call(__e, PrimFunc(symhdstr), tmp2097)

							tmp2099 := PrimEqual(MakeString("R"), tmp2098)

							var ifres2096 Obj

							if True == tmp2099 {
								ifres2096 = True

							} else {
								ifres2096 = False

							}

							ifres2095 = ifres2096

						} else {
							ifres2095 = False

						}

						var ifres2094 Obj

						if True == ifres2095 {
							ifres2094 = True

						} else {
							ifres2094 = False

						}

						ifres2093 = ifres2094

					} else {
						ifres2093 = False

					}

					var ifres2092 Obj

					if True == ifres2093 {
						ifres2092 = True

					} else {
						ifres2092 = False

					}

					ifres2091 = ifres2092

				} else {
					ifres2091 = False

				}

				if True == ifres2091 {
					tmp1989 := PrimTailString(V4747)

					tmp1990 := PrimTailString(tmp1989)

					tmp1991 := PrimCons(symshen_4r, Nil)

					tmp1992 := PrimCons(tmp1990, tmp1991)

					tmp1993 := PrimCons(V4746, tmp1992)

					__e.Return(PrimCons(symshen_4app, tmp1993))
					return

				} else {
					tmp2089 := Call(__e, PrimFunc(symshen_4_7string_2), V4747)

					var ifres2076 Obj

					if True == tmp2089 {
						tmp2087 := Call(__e, PrimFunc(symhdstr), V4747)

						tmp2088 := PrimEqual(MakeString("~"), tmp2087)

						var ifres2078 Obj

						if True == tmp2088 {
							tmp2085 := PrimTailString(V4747)

							tmp2086 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2085)

							var ifres2080 Obj

							if True == tmp2086 {
								tmp2082 := PrimTailString(V4747)

								tmp2083 := Call(__e, PrimFunc(symhdstr), tmp2082)

								tmp2084 := PrimEqual(MakeString("S"), tmp2083)

								var ifres2081 Obj

								if True == tmp2084 {
									ifres2081 = True

								} else {
									ifres2081 = False

								}

								ifres2080 = ifres2081

							} else {
								ifres2080 = False

							}

							var ifres2079 Obj

							if True == ifres2080 {
								ifres2079 = True

							} else {
								ifres2079 = False

							}

							ifres2078 = ifres2079

						} else {
							ifres2078 = False

						}

						var ifres2077 Obj

						if True == ifres2078 {
							ifres2077 = True

						} else {
							ifres2077 = False

						}

						ifres2076 = ifres2077

					} else {
						ifres2076 = False

					}

					if True == ifres2076 {
						tmp1994 := PrimTailString(V4747)

						tmp1995 := PrimTailString(tmp1994)

						tmp1996 := PrimCons(symshen_4s, Nil)

						tmp1997 := PrimCons(tmp1995, tmp1996)

						tmp1998 := PrimCons(V4746, tmp1997)

						__e.Return(PrimCons(symshen_4app, tmp1998))
						return

					} else {
						tmp2074 := Call(__e, PrimFunc(symshen_4_7string_2), V4747)

						if True == tmp2074 {
							tmp1999 := Call(__e, PrimFunc(symhdstr), V4747)

							tmp2000 := PrimTailString(V4747)

							tmp2001 := Call(__e, PrimFunc(symshen_4insert_1l), V4746, tmp2000)

							tmp2002 := PrimCons(tmp2001, Nil)

							tmp2003 := PrimCons(tmp1999, tmp2002)

							tmp2004 := PrimCons(symcn, tmp2003)

							__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp2004)
							return

						} else {
							tmp2072 := PrimIsPair(V4747)

							var ifres2053 Obj

							if True == tmp2072 {
								tmp2070 := PrimHead(V4747)

								tmp2071 := PrimEqual(symcn, tmp2070)

								var ifres2055 Obj

								if True == tmp2071 {
									tmp2068 := PrimTail(V4747)

									tmp2069 := PrimIsPair(tmp2068)

									var ifres2057 Obj

									if True == tmp2069 {
										tmp2065 := PrimTail(V4747)

										tmp2066 := PrimTail(tmp2065)

										tmp2067 := PrimIsPair(tmp2066)

										var ifres2059 Obj

										if True == tmp2067 {
											tmp2061 := PrimTail(V4747)

											tmp2062 := PrimTail(tmp2061)

											tmp2063 := PrimTail(tmp2062)

											tmp2064 := PrimEqual(Nil, tmp2063)

											var ifres2060 Obj

											if True == tmp2064 {
												ifres2060 = True

											} else {
												ifres2060 = False

											}

											ifres2059 = ifres2060

										} else {
											ifres2059 = False

										}

										var ifres2058 Obj

										if True == ifres2059 {
											ifres2058 = True

										} else {
											ifres2058 = False

										}

										ifres2057 = ifres2058

									} else {
										ifres2057 = False

									}

									var ifres2056 Obj

									if True == ifres2057 {
										ifres2056 = True

									} else {
										ifres2056 = False

									}

									ifres2055 = ifres2056

								} else {
									ifres2055 = False

								}

								var ifres2054 Obj

								if True == ifres2055 {
									ifres2054 = True

								} else {
									ifres2054 = False

								}

								ifres2053 = ifres2054

							} else {
								ifres2053 = False

							}

							if True == ifres2053 {
								tmp2005 := PrimTail(V4747)

								tmp2006 := PrimHead(tmp2005)

								tmp2007 := PrimTail(V4747)

								tmp2008 := PrimTail(tmp2007)

								tmp2009 := PrimHead(tmp2008)

								tmp2010 := Call(__e, PrimFunc(symshen_4insert_1l), V4746, tmp2009)

								tmp2011 := PrimCons(tmp2010, Nil)

								tmp2012 := PrimCons(tmp2006, tmp2011)

								__e.Return(PrimCons(symcn, tmp2012))
								return

							} else {
								tmp2051 := PrimIsPair(V4747)

								var ifres2025 Obj

								if True == tmp2051 {
									tmp2049 := PrimHead(V4747)

									tmp2050 := PrimEqual(symshen_4app, tmp2049)

									var ifres2027 Obj

									if True == tmp2050 {
										tmp2047 := PrimTail(V4747)

										tmp2048 := PrimIsPair(tmp2047)

										var ifres2029 Obj

										if True == tmp2048 {
											tmp2044 := PrimTail(V4747)

											tmp2045 := PrimTail(tmp2044)

											tmp2046 := PrimIsPair(tmp2045)

											var ifres2031 Obj

											if True == tmp2046 {
												tmp2040 := PrimTail(V4747)

												tmp2041 := PrimTail(tmp2040)

												tmp2042 := PrimTail(tmp2041)

												tmp2043 := PrimIsPair(tmp2042)

												var ifres2033 Obj

												if True == tmp2043 {
													tmp2035 := PrimTail(V4747)

													tmp2036 := PrimTail(tmp2035)

													tmp2037 := PrimTail(tmp2036)

													tmp2038 := PrimTail(tmp2037)

													tmp2039 := PrimEqual(Nil, tmp2038)

													var ifres2034 Obj

													if True == tmp2039 {
														ifres2034 = True

													} else {
														ifres2034 = False

													}

													ifres2033 = ifres2034

												} else {
													ifres2033 = False

												}

												var ifres2032 Obj

												if True == ifres2033 {
													ifres2032 = True

												} else {
													ifres2032 = False

												}

												ifres2031 = ifres2032

											} else {
												ifres2031 = False

											}

											var ifres2030 Obj

											if True == ifres2031 {
												ifres2030 = True

											} else {
												ifres2030 = False

											}

											ifres2029 = ifres2030

										} else {
											ifres2029 = False

										}

										var ifres2028 Obj

										if True == ifres2029 {
											ifres2028 = True

										} else {
											ifres2028 = False

										}

										ifres2027 = ifres2028

									} else {
										ifres2027 = False

									}

									var ifres2026 Obj

									if True == ifres2027 {
										ifres2026 = True

									} else {
										ifres2026 = False

									}

									ifres2025 = ifres2026

								} else {
									ifres2025 = False

								}

								if True == ifres2025 {
									tmp2013 := PrimTail(V4747)

									tmp2014 := PrimHead(tmp2013)

									tmp2015 := PrimTail(V4747)

									tmp2016 := PrimTail(tmp2015)

									tmp2017 := PrimHead(tmp2016)

									tmp2018 := Call(__e, PrimFunc(symshen_4insert_1l), V4746, tmp2017)

									tmp2019 := PrimTail(V4747)

									tmp2020 := PrimTail(tmp2019)

									tmp2021 := PrimTail(tmp2020)

									tmp2022 := PrimCons(tmp2018, tmp2021)

									tmp2023 := PrimCons(tmp2014, tmp2022)

									__e.Return(PrimCons(symshen_4app, tmp2023))
									return

								} else {
									__e.Return(PrimSimpleError(MakeString("implementation error in shen.insert-l")))
									return
								}

							}

						}

					}

				}

			}

		}

	}, 2)

	tmp2122 := Call(__e, ns2_1set, symshen_4insert_1l, tmp1983)

	_ = tmp2122

	tmp2123 := MakeNative(func(__e *ControlFlow) {
		V4748 := __e.Get(1)
		_ = V4748
		tmp2208 := PrimIsPair(V4748)

		var ifres2139 Obj

		if True == tmp2208 {
			tmp2206 := PrimHead(V4748)

			tmp2207 := PrimEqual(symcn, tmp2206)

			var ifres2141 Obj

			if True == tmp2207 {
				tmp2204 := PrimTail(V4748)

				tmp2205 := PrimIsPair(tmp2204)

				var ifres2143 Obj

				if True == tmp2205 {
					tmp2201 := PrimTail(V4748)

					tmp2202 := PrimTail(tmp2201)

					tmp2203 := PrimIsPair(tmp2202)

					var ifres2145 Obj

					if True == tmp2203 {
						tmp2197 := PrimTail(V4748)

						tmp2198 := PrimTail(tmp2197)

						tmp2199 := PrimHead(tmp2198)

						tmp2200 := PrimIsPair(tmp2199)

						var ifres2147 Obj

						if True == tmp2200 {
							tmp2192 := PrimTail(V4748)

							tmp2193 := PrimTail(tmp2192)

							tmp2194 := PrimHead(tmp2193)

							tmp2195 := PrimHead(tmp2194)

							tmp2196 := PrimEqual(symcn, tmp2195)

							var ifres2149 Obj

							if True == tmp2196 {
								tmp2187 := PrimTail(V4748)

								tmp2188 := PrimTail(tmp2187)

								tmp2189 := PrimHead(tmp2188)

								tmp2190 := PrimTail(tmp2189)

								tmp2191 := PrimIsPair(tmp2190)

								var ifres2151 Obj

								if True == tmp2191 {
									tmp2181 := PrimTail(V4748)

									tmp2182 := PrimTail(tmp2181)

									tmp2183 := PrimHead(tmp2182)

									tmp2184 := PrimTail(tmp2183)

									tmp2185 := PrimTail(tmp2184)

									tmp2186 := PrimIsPair(tmp2185)

									var ifres2153 Obj

									if True == tmp2186 {
										tmp2174 := PrimTail(V4748)

										tmp2175 := PrimTail(tmp2174)

										tmp2176 := PrimHead(tmp2175)

										tmp2177 := PrimTail(tmp2176)

										tmp2178 := PrimTail(tmp2177)

										tmp2179 := PrimTail(tmp2178)

										tmp2180 := PrimEqual(Nil, tmp2179)

										var ifres2155 Obj

										if True == tmp2180 {
											tmp2170 := PrimTail(V4748)

											tmp2171 := PrimTail(tmp2170)

											tmp2172 := PrimTail(tmp2171)

											tmp2173 := PrimEqual(Nil, tmp2172)

											var ifres2157 Obj

											if True == tmp2173 {
												tmp2167 := PrimTail(V4748)

												tmp2168 := PrimHead(tmp2167)

												tmp2169 := PrimIsString(tmp2168)

												var ifres2159 Obj

												if True == tmp2169 {
													tmp2161 := PrimTail(V4748)

													tmp2162 := PrimTail(tmp2161)

													tmp2163 := PrimHead(tmp2162)

													tmp2164 := PrimTail(tmp2163)

													tmp2165 := PrimHead(tmp2164)

													tmp2166 := PrimIsString(tmp2165)

													var ifres2160 Obj

													if True == tmp2166 {
														ifres2160 = True

													} else {
														ifres2160 = False

													}

													ifres2159 = ifres2160

												} else {
													ifres2159 = False

												}

												var ifres2158 Obj

												if True == ifres2159 {
													ifres2158 = True

												} else {
													ifres2158 = False

												}

												ifres2157 = ifres2158

											} else {
												ifres2157 = False

											}

											var ifres2156 Obj

											if True == ifres2157 {
												ifres2156 = True

											} else {
												ifres2156 = False

											}

											ifres2155 = ifres2156

										} else {
											ifres2155 = False

										}

										var ifres2154 Obj

										if True == ifres2155 {
											ifres2154 = True

										} else {
											ifres2154 = False

										}

										ifres2153 = ifres2154

									} else {
										ifres2153 = False

									}

									var ifres2152 Obj

									if True == ifres2153 {
										ifres2152 = True

									} else {
										ifres2152 = False

									}

									ifres2151 = ifres2152

								} else {
									ifres2151 = False

								}

								var ifres2150 Obj

								if True == ifres2151 {
									ifres2150 = True

								} else {
									ifres2150 = False

								}

								ifres2149 = ifres2150

							} else {
								ifres2149 = False

							}

							var ifres2148 Obj

							if True == ifres2149 {
								ifres2148 = True

							} else {
								ifres2148 = False

							}

							ifres2147 = ifres2148

						} else {
							ifres2147 = False

						}

						var ifres2146 Obj

						if True == ifres2147 {
							ifres2146 = True

						} else {
							ifres2146 = False

						}

						ifres2145 = ifres2146

					} else {
						ifres2145 = False

					}

					var ifres2144 Obj

					if True == ifres2145 {
						ifres2144 = True

					} else {
						ifres2144 = False

					}

					ifres2143 = ifres2144

				} else {
					ifres2143 = False

				}

				var ifres2142 Obj

				if True == ifres2143 {
					ifres2142 = True

				} else {
					ifres2142 = False

				}

				ifres2141 = ifres2142

			} else {
				ifres2141 = False

			}

			var ifres2140 Obj

			if True == ifres2141 {
				ifres2140 = True

			} else {
				ifres2140 = False

			}

			ifres2139 = ifres2140

		} else {
			ifres2139 = False

		}

		if True == ifres2139 {
			tmp2124 := PrimTail(V4748)

			tmp2125 := PrimHead(tmp2124)

			tmp2126 := PrimTail(V4748)

			tmp2127 := PrimTail(tmp2126)

			tmp2128 := PrimHead(tmp2127)

			tmp2129 := PrimTail(tmp2128)

			tmp2130 := PrimHead(tmp2129)

			tmp2131 := PrimStringConcat(tmp2125, tmp2130)

			tmp2132 := PrimTail(V4748)

			tmp2133 := PrimTail(tmp2132)

			tmp2134 := PrimHead(tmp2133)

			tmp2135 := PrimTail(tmp2134)

			tmp2136 := PrimTail(tmp2135)

			tmp2137 := PrimCons(tmp2131, tmp2136)

			__e.Return(PrimCons(symcn, tmp2137))
			return

		} else {
			__e.Return(V4748)
			return
		}

	}, 1)

	tmp2209 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp2123)

	_ = tmp2209

	tmp2210 := MakeNative(func(__e *ControlFlow) {
		V4751 := __e.Get(1)
		_ = V4751
		tmp2236 := PrimEqual(MakeString(""), V4751)

		if True == tmp2236 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp2234 := Call(__e, PrimFunc(symshen_4_7string_2), V4751)

			var ifres2221 Obj

			if True == tmp2234 {
				tmp2232 := Call(__e, PrimFunc(symhdstr), V4751)

				tmp2233 := PrimEqual(MakeString("~"), tmp2232)

				var ifres2223 Obj

				if True == tmp2233 {
					tmp2230 := PrimTailString(V4751)

					tmp2231 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2230)

					var ifres2225 Obj

					if True == tmp2231 {
						tmp2227 := PrimTailString(V4751)

						tmp2228 := Call(__e, PrimFunc(symhdstr), tmp2227)

						tmp2229 := PrimEqual(MakeString("%"), tmp2228)

						var ifres2226 Obj

						if True == tmp2229 {
							ifres2226 = True

						} else {
							ifres2226 = False

						}

						ifres2225 = ifres2226

					} else {
						ifres2225 = False

					}

					var ifres2224 Obj

					if True == ifres2225 {
						ifres2224 = True

					} else {
						ifres2224 = False

					}

					ifres2223 = ifres2224

				} else {
					ifres2223 = False

				}

				var ifres2222 Obj

				if True == ifres2223 {
					ifres2222 = True

				} else {
					ifres2222 = False

				}

				ifres2221 = ifres2222

			} else {
				ifres2221 = False

			}

			if True == ifres2221 {
				tmp2211 := PrimNumberToString(MakeNumber(10))

				tmp2212 := PrimTailString(V4751)

				tmp2213 := PrimTailString(tmp2212)

				tmp2214 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp2213)

				__e.Return(PrimStringConcat(tmp2211, tmp2214))
				return

			} else {
				tmp2219 := Call(__e, PrimFunc(symshen_4_7string_2), V4751)

				if True == tmp2219 {
					tmp2215 := Call(__e, PrimFunc(symhdstr), V4751)

					tmp2216 := PrimTailString(V4751)

					tmp2217 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp2216)

					__e.Return(PrimStringConcat(tmp2215, tmp2217))
					return

				} else {
					__e.Return(PrimSimpleError(MakeString("implementation error in shen.proc-nl")))
					return
				}

			}

		}

	}, 1)

	tmp2237 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp2210)

	_ = tmp2237

	tmp2238 := MakeNative(func(__e *ControlFlow) {
		V4756 := __e.Get(1)
		_ = V4756
		V4757 := __e.Get(2)
		_ = V4757
		tmp2247 := PrimEqual(Nil, V4757)

		if True == tmp2247 {
			__e.Return(V4756)
			return
		} else {
			tmp2245 := PrimIsPair(V4757)

			if True == tmp2245 {
				tmp2239 := PrimHead(V4757)

				tmp2240 := PrimCons(V4756, Nil)

				tmp2241 := PrimCons(tmp2239, tmp2240)

				tmp2242 := PrimCons(symshen_4insert, tmp2241)

				tmp2243 := PrimTail(V4757)

				__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2242, tmp2243)
				return

			} else {
				__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-r")))
				return
			}

		}

	}, 2)

	tmp2248 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp2238)

	_ = tmp2248

	tmp2249 := MakeNative(func(__e *ControlFlow) {
		V4758 := __e.Get(1)
		_ = V4758
		V4759 := __e.Get(2)
		_ = V4759
		__e.TailApply(PrimFunc(symshen_4insert_1h), V4758, V4759, MakeString(""))
		return
	}, 2)

	tmp2250 := Call(__e, ns2_1set, symshen_4insert, tmp2249)

	_ = tmp2250

	tmp2251 := MakeNative(func(__e *ControlFlow) {
		V4768 := __e.Get(1)
		_ = V4768
		V4769 := __e.Get(2)
		_ = V4769
		V4770 := __e.Get(3)
		_ = V4770
		tmp2312 := PrimEqual(MakeString(""), V4769)

		if True == tmp2312 {
			__e.Return(V4770)
			return
		} else {
			tmp2310 := Call(__e, PrimFunc(symshen_4_7string_2), V4769)

			var ifres2297 Obj

			if True == tmp2310 {
				tmp2308 := Call(__e, PrimFunc(symhdstr), V4769)

				tmp2309 := PrimEqual(MakeString("~"), tmp2308)

				var ifres2299 Obj

				if True == tmp2309 {
					tmp2306 := PrimTailString(V4769)

					tmp2307 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2306)

					var ifres2301 Obj

					if True == tmp2307 {
						tmp2303 := PrimTailString(V4769)

						tmp2304 := Call(__e, PrimFunc(symhdstr), tmp2303)

						tmp2305 := PrimEqual(MakeString("A"), tmp2304)

						var ifres2302 Obj

						if True == tmp2305 {
							ifres2302 = True

						} else {
							ifres2302 = False

						}

						ifres2301 = ifres2302

					} else {
						ifres2301 = False

					}

					var ifres2300 Obj

					if True == ifres2301 {
						ifres2300 = True

					} else {
						ifres2300 = False

					}

					ifres2299 = ifres2300

				} else {
					ifres2299 = False

				}

				var ifres2298 Obj

				if True == ifres2299 {
					ifres2298 = True

				} else {
					ifres2298 = False

				}

				ifres2297 = ifres2298

			} else {
				ifres2297 = False

			}

			if True == ifres2297 {
				tmp2252 := PrimTailString(V4769)

				tmp2253 := PrimTailString(tmp2252)

				tmp2254 := Call(__e, PrimFunc(symshen_4app), V4768, tmp2253, symshen_4a)

				__e.Return(PrimStringConcat(V4770, tmp2254))
				return

			} else {
				tmp2295 := Call(__e, PrimFunc(symshen_4_7string_2), V4769)

				var ifres2282 Obj

				if True == tmp2295 {
					tmp2293 := Call(__e, PrimFunc(symhdstr), V4769)

					tmp2294 := PrimEqual(MakeString("~"), tmp2293)

					var ifres2284 Obj

					if True == tmp2294 {
						tmp2291 := PrimTailString(V4769)

						tmp2292 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2291)

						var ifres2286 Obj

						if True == tmp2292 {
							tmp2288 := PrimTailString(V4769)

							tmp2289 := Call(__e, PrimFunc(symhdstr), tmp2288)

							tmp2290 := PrimEqual(MakeString("R"), tmp2289)

							var ifres2287 Obj

							if True == tmp2290 {
								ifres2287 = True

							} else {
								ifres2287 = False

							}

							ifres2286 = ifres2287

						} else {
							ifres2286 = False

						}

						var ifres2285 Obj

						if True == ifres2286 {
							ifres2285 = True

						} else {
							ifres2285 = False

						}

						ifres2284 = ifres2285

					} else {
						ifres2284 = False

					}

					var ifres2283 Obj

					if True == ifres2284 {
						ifres2283 = True

					} else {
						ifres2283 = False

					}

					ifres2282 = ifres2283

				} else {
					ifres2282 = False

				}

				if True == ifres2282 {
					tmp2255 := PrimTailString(V4769)

					tmp2256 := PrimTailString(tmp2255)

					tmp2257 := Call(__e, PrimFunc(symshen_4app), V4768, tmp2256, symshen_4r)

					__e.Return(PrimStringConcat(V4770, tmp2257))
					return

				} else {
					tmp2280 := Call(__e, PrimFunc(symshen_4_7string_2), V4769)

					var ifres2267 Obj

					if True == tmp2280 {
						tmp2278 := Call(__e, PrimFunc(symhdstr), V4769)

						tmp2279 := PrimEqual(MakeString("~"), tmp2278)

						var ifres2269 Obj

						if True == tmp2279 {
							tmp2276 := PrimTailString(V4769)

							tmp2277 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2276)

							var ifres2271 Obj

							if True == tmp2277 {
								tmp2273 := PrimTailString(V4769)

								tmp2274 := Call(__e, PrimFunc(symhdstr), tmp2273)

								tmp2275 := PrimEqual(MakeString("S"), tmp2274)

								var ifres2272 Obj

								if True == tmp2275 {
									ifres2272 = True

								} else {
									ifres2272 = False

								}

								ifres2271 = ifres2272

							} else {
								ifres2271 = False

							}

							var ifres2270 Obj

							if True == ifres2271 {
								ifres2270 = True

							} else {
								ifres2270 = False

							}

							ifres2269 = ifres2270

						} else {
							ifres2269 = False

						}

						var ifres2268 Obj

						if True == ifres2269 {
							ifres2268 = True

						} else {
							ifres2268 = False

						}

						ifres2267 = ifres2268

					} else {
						ifres2267 = False

					}

					if True == ifres2267 {
						tmp2258 := PrimTailString(V4769)

						tmp2259 := PrimTailString(tmp2258)

						tmp2260 := Call(__e, PrimFunc(symshen_4app), V4768, tmp2259, symshen_4s)

						__e.Return(PrimStringConcat(V4770, tmp2260))
						return

					} else {
						tmp2265 := Call(__e, PrimFunc(symshen_4_7string_2), V4769)

						if True == tmp2265 {
							tmp2261 := PrimTailString(V4769)

							tmp2262 := Call(__e, PrimFunc(symhdstr), V4769)

							tmp2263 := PrimStringConcat(V4770, tmp2262)

							__e.TailApply(PrimFunc(symshen_4insert_1h), V4768, tmp2261, tmp2263)
							return

						} else {
							__e.Return(PrimSimpleError(MakeString("implementation error in shen.insert-h")))
							return
						}

					}

				}

			}

		}

	}, 3)

	tmp2313 := Call(__e, ns2_1set, symshen_4insert_1h, tmp2251)

	_ = tmp2313

	tmp2314 := MakeNative(func(__e *ControlFlow) {
		V4771 := __e.Get(1)
		_ = V4771
		V4772 := __e.Get(2)
		_ = V4772
		V4773 := __e.Get(3)
		_ = V4773
		tmp2315 := Call(__e, PrimFunc(symshen_4arg_1_6str), V4771, V4773)

		__e.Return(PrimStringConcat(tmp2315, V4772))
		return

	}, 3)

	tmp2316 := Call(__e, ns2_1set, symshen_4app, tmp2314)

	_ = tmp2316

	tmp2317 := MakeNative(func(__e *ControlFlow) {
		V4777 := __e.Get(1)
		_ = V4777
		V4778 := __e.Get(2)
		_ = V4778
		tmp2325 := Call(__e, PrimFunc(symfail))

		tmp2326 := PrimEqual(V4777, tmp2325)

		if True == tmp2326 {
			__e.Return(MakeString("..."))
			return
		} else {
			tmp2323 := Call(__e, PrimFunc(symshen_4list_2), V4777)

			if True == tmp2323 {
				__e.TailApply(PrimFunc(symshen_4list_1_6str), V4777, V4778)
				return
			} else {
				tmp2321 := PrimIsString(V4777)

				if True == tmp2321 {
					__e.TailApply(PrimFunc(symshen_4str_1_6str), V4777, V4778)
					return
				} else {
					tmp2319 := PrimIsVector(V4777)

					if True == tmp2319 {
						__e.TailApply(PrimFunc(symshen_4vector_1_6str), V4777, V4778)
						return
					} else {
						__e.TailApply(PrimFunc(symshen_4atom_1_6str), V4777)
						return
					}

				}

			}

		}

	}, 2)

	tmp2327 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp2317)

	_ = tmp2327

	tmp2328 := MakeNative(func(__e *ControlFlow) {
		V4779 := __e.Get(1)
		_ = V4779
		V4780 := __e.Get(2)
		_ = V4780
		tmp2336 := PrimEqual(symshen_4r, V4780)

		if True == tmp2336 {
			tmp2329 := Call(__e, PrimFunc(symshen_4maxseq))

			tmp2330 := Call(__e, PrimFunc(symshen_4iter_1list), V4779, symshen_4r, tmp2329)

			tmp2331 := Call(__e, PrimFunc(sym_8s), tmp2330, MakeString(")"))

			__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp2331)
			return

		} else {
			tmp2332 := Call(__e, PrimFunc(symshen_4maxseq))

			tmp2333 := Call(__e, PrimFunc(symshen_4iter_1list), V4779, V4780, tmp2332)

			tmp2334 := Call(__e, PrimFunc(sym_8s), tmp2333, MakeString("]"))

			__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp2334)
			return

		}

	}, 2)

	tmp2337 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp2328)

	_ = tmp2337

	tmp2338 := MakeNative(func(__e *ControlFlow) {
		__e.Return(PrimValue(sym_dmaximum_1print_1sequence_1size_d))
		return
	}, 0)

	tmp2339 := Call(__e, ns2_1set, symshen_4maxseq, tmp2338)

	_ = tmp2339

	tmp2340 := MakeNative(func(__e *ControlFlow) {
		V4791 := __e.Get(1)
		_ = V4791
		V4792 := __e.Get(2)
		_ = V4792
		V4793 := __e.Get(3)
		_ = V4793
		tmp2361 := PrimEqual(Nil, V4791)

		if True == tmp2361 {
			__e.Return(MakeString(""))
			return
		} else {
			tmp2359 := PrimEqual(MakeNumber(0), V4793)

			if True == tmp2359 {
				__e.Return(MakeString("... etc"))
				return
			} else {
				tmp2357 := PrimIsPair(V4791)

				var ifres2353 Obj

				if True == tmp2357 {
					tmp2355 := PrimTail(V4791)

					tmp2356 := PrimEqual(Nil, tmp2355)

					var ifres2354 Obj

					if True == tmp2356 {
						ifres2354 = True

					} else {
						ifres2354 = False

					}

					ifres2353 = ifres2354

				} else {
					ifres2353 = False

				}

				if True == ifres2353 {
					tmp2341 := PrimHead(V4791)

					__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp2341, V4792)
					return

				} else {
					tmp2351 := PrimIsPair(V4791)

					if True == tmp2351 {
						tmp2342 := PrimHead(V4791)

						tmp2343 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2342, V4792)

						tmp2344 := PrimTail(V4791)

						tmp2345 := PrimNumberSubtract(V4793, MakeNumber(1))

						tmp2346 := Call(__e, PrimFunc(symshen_4iter_1list), tmp2344, V4792, tmp2345)

						tmp2347 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2346)

						__e.TailApply(PrimFunc(sym_8s), tmp2343, tmp2347)
						return

					} else {
						tmp2348 := Call(__e, PrimFunc(symshen_4arg_1_6str), V4791, V4792)

						tmp2349 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2348)

						__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp2349)
						return

					}

				}

			}

		}

	}, 3)

	tmp2362 := Call(__e, ns2_1set, symshen_4iter_1list, tmp2340)

	_ = tmp2362

	tmp2363 := MakeNative(func(__e *ControlFlow) {
		V4796 := __e.Get(1)
		_ = V4796
		V4797 := __e.Get(2)
		_ = V4797
		tmp2368 := PrimEqual(symshen_4a, V4797)

		if True == tmp2368 {
			__e.Return(V4796)
			return
		} else {
			tmp2364 := PrimNumberToString(MakeNumber(34))

			tmp2365 := PrimNumberToString(MakeNumber(34))

			tmp2366 := Call(__e, PrimFunc(sym_8s), V4796, tmp2365)

			__e.TailApply(PrimFunc(sym_8s), tmp2364, tmp2366)
			return

		}

	}, 2)

	tmp2369 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp2363)

	_ = tmp2369

	tmp2370 := MakeNative(func(__e *ControlFlow) {
		V4798 := __e.Get(1)
		_ = V4798
		V4799 := __e.Get(2)
		_ = V4799
		tmp2383 := Call(__e, PrimFunc(symshen_4print_1vector_2), V4798)

		if True == tmp2383 {
			tmp2371 := PrimVectorGet(V4798, MakeNumber(0))

			tmp2372 := Call(__e, PrimFunc(symfn), tmp2371)

			__e.TailApply(tmp2372, V4798)
			return

		} else {
			tmp2381 := Call(__e, PrimFunc(symvector_2), V4798)

			if True == tmp2381 {
				tmp2373 := Call(__e, PrimFunc(symshen_4maxseq))

				tmp2374 := Call(__e, PrimFunc(symshen_4iter_1vector), V4798, MakeNumber(1), V4799, tmp2373)

				tmp2375 := Call(__e, PrimFunc(sym_8s), tmp2374, MakeString(">"))

				__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2375)
				return

			} else {
				tmp2376 := Call(__e, PrimFunc(symshen_4maxseq))

				tmp2377 := Call(__e, PrimFunc(symshen_4iter_1vector), V4798, MakeNumber(0), V4799, tmp2376)

				tmp2378 := Call(__e, PrimFunc(sym_8s), tmp2377, MakeString(">>"))

				tmp2379 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp2378)

				__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2379)
				return

			}

		}

	}, 2)

	tmp2384 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp2370)

	_ = tmp2384

	tmp2385 := MakeNative(func(__e *ControlFlow) {
		V4800 := __e.Get(1)
		_ = V4800
		tmp2386 := MakeNative(func(__e *ControlFlow) {
			Zero := __e.Get(1)
			_ = Zero
			tmp2393 := PrimEqual(Zero, symshen_4tuple)

			if True == tmp2393 {
				__e.Return(True)
				return
			} else {
				tmp2391 := PrimEqual(Zero, symshen_4pvar)

				if True == tmp2391 {
					__e.Return(True)
					return
				} else {
					tmp2388 := PrimIsNumber(Zero)

					tmp2389 := PrimNot(tmp2388)

					if True == tmp2389 {
						__e.TailApply(PrimFunc(symshen_4fbound_2), Zero)
						return
					} else {
						__e.Return(False)
						return
					}

				}

			}

		}, 1)

		tmp2394 := PrimVectorGet(V4800, MakeNumber(0))

		__e.TailApply(tmp2386, tmp2394)
		return

	}, 1)

	tmp2395 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp2385)

	_ = tmp2395

	tmp2396 := MakeNative(func(__e *ControlFlow) {
		V4801 := __e.Get(1)
		_ = V4801
		tmp2397 := Call(__e, PrimFunc(symarity), V4801)

		tmp2398 := PrimEqual(tmp2397, MakeNumber(-1))

		__e.Return(PrimNot(tmp2398))
		return

	}, 1)

	tmp2399 := Call(__e, ns2_1set, symshen_4fbound_2, tmp2396)

	_ = tmp2399

	tmp2400 := MakeNative(func(__e *ControlFlow) {
		V4802 := __e.Get(1)
		_ = V4802
		tmp2401 := PrimVectorGet(V4802, MakeNumber(1))

		tmp2402 := PrimVectorGet(V4802, MakeNumber(2))

		tmp2403 := Call(__e, PrimFunc(symshen_4app), tmp2402, MakeString(")"), symshen_4s)

		tmp2404 := PrimStringConcat(MakeString(" "), tmp2403)

		tmp2405 := Call(__e, PrimFunc(symshen_4app), tmp2401, tmp2404, symshen_4s)

		__e.Return(PrimStringConcat(MakeString("(@p "), tmp2405))
		return

	}, 1)

	tmp2406 := Call(__e, ns2_1set, symshen_4tuple, tmp2400)

	_ = tmp2406

	tmp2407 := MakeNative(func(__e *ControlFlow) {
		V4809 := __e.Get(1)
		_ = V4809
		V4810 := __e.Get(2)
		_ = V4810
		V4811 := __e.Get(3)
		_ = V4811
		V4812 := __e.Get(4)
		_ = V4812
		tmp2427 := PrimEqual(MakeNumber(0), V4812)

		if True == tmp2427 {
			__e.Return(MakeString("... etc"))
			return
		} else {
			tmp2408 := MakeNative(func(__e *ControlFlow) {
				Item := __e.Get(1)
				_ = Item
				tmp2409 := MakeNative(func(__e *ControlFlow) {
					Next := __e.Get(1)
					_ = Next
					tmp2418 := PrimEqual(Item, symshen_4out_1of_1bounds)

					if True == tmp2418 {
						__e.Return(MakeString(""))
						return
					} else {
						tmp2416 := PrimEqual(Next, symshen_4out_1of_1bounds)

						if True == tmp2416 {
							__e.TailApply(PrimFunc(symshen_4arg_1_6str), Item, V4811)
							return
						} else {
							tmp2410 := Call(__e, PrimFunc(symshen_4arg_1_6str), Item, V4811)

							tmp2411 := PrimNumberAdd(V4810, MakeNumber(1))

							tmp2412 := PrimNumberSubtract(V4812, MakeNumber(1))

							tmp2413 := Call(__e, PrimFunc(symshen_4iter_1vector), V4809, tmp2411, V4811, tmp2412)

							tmp2414 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2413)

							__e.TailApply(PrimFunc(sym_8s), tmp2410, tmp2414)
							return

						}

					}

				}, 1)

				tmp2419 := MakeNative(func(__e *ControlFlow) {
					tmp2420 := PrimNumberAdd(V4810, MakeNumber(1))

					__e.Return(PrimVectorGet(V4809, tmp2420))
					return

				}, 0)

				tmp2421 := MakeNative(func(__e *ControlFlow) {
					E := __e.Get(1)
					_ = E
					__e.Return(symshen_4out_1of_1bounds)
					return
				}, 1)

				tmp2422 := Call(__e, try_1catch, tmp2419, tmp2421)

				__e.TailApply(tmp2409, tmp2422)
				return

			}, 1)

			tmp2423 := MakeNative(func(__e *ControlFlow) {
				__e.Return(PrimVectorGet(V4809, V4810))
				return
			}, 0)

			tmp2424 := MakeNative(func(__e *ControlFlow) {
				E := __e.Get(1)
				_ = E
				__e.Return(symshen_4out_1of_1bounds)
				return
			}, 1)

			tmp2425 := Call(__e, try_1catch, tmp2423, tmp2424)

			__e.TailApply(tmp2408, tmp2425)
			return

		}

	}, 4)

	tmp2428 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp2407)

	_ = tmp2428

	tmp2429 := MakeNative(func(__e *ControlFlow) {
		V4813 := __e.Get(1)
		_ = V4813
		tmp2430 := MakeNative(func(__e *ControlFlow) {
			__e.Return(PrimStr(V4813))
			return
		}, 0)

		tmp2431 := MakeNative(func(__e *ControlFlow) {
			E := __e.Get(1)
			_ = E
			__e.TailApply(PrimFunc(symshen_4funexstring))
			return
		}, 1)

		__e.TailApply(try_1catch, tmp2430, tmp2431)
		return

	}, 1)

	tmp2432 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp2429)

	_ = tmp2432

	tmp2433 := MakeNative(func(__e *ControlFlow) {
		tmp2434 := PrimIntern(MakeString("x"))

		tmp2435 := Call(__e, PrimFunc(symgensym), tmp2434)

		tmp2436 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2435, symshen_4a)

		tmp2437 := Call(__e, PrimFunc(sym_8s), tmp2436, MakeString("\x11"))

		tmp2438 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp2437)

		tmp2439 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2438)

		tmp2440 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp2439)

		tmp2441 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2440)

		__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp2441)
		return

	}, 0)

	tmp2442 := Call(__e, ns2_1set, symshen_4funexstring, tmp2433)

	_ = tmp2442

	tmp2443 := MakeNative(func(__e *ControlFlow) {
		V4814 := __e.Get(1)
		_ = V4814
		tmp2447 := Call(__e, PrimFunc(symempty_2), V4814)

		if True == tmp2447 {
			__e.Return(True)
			return
		} else {
			tmp2445 := PrimIsPair(V4814)

			if True == tmp2445 {
				__e.Return(True)
				return
			} else {
				__e.Return(False)
				return
			}

		}

	}, 1)

	__e.TailApply(ns2_1set, symshen_4list_2, tmp2443)
	return

}, 0)
