package main

import . "github.com/tiancaiamao/shen-go/kl"

var WriterMain = MakeNative(func(__e *ControlFlow) {
tmp2078 := MakeNative(func(__e *ControlFlow) {
V5619 := __e.Get(1)
_ = V5619
tmp2079 := MakeNative(func(__e *ControlFlow) {
W5620 := __e.Get(1)
_ = W5620
tmp2080 := MakeNative(func(__e *ControlFlow) {
W5621 := __e.Get(1)
_ = W5621
__e.Return(V5619)
return
}, 1)

tmp2081 := Call(__e, PrimFunc(symstoutput))


tmp2082 := Call(__e, PrimFunc(sympr), W5620, tmp2081)


__e.TailApply(tmp2080, tmp2082)
return


}, 1)

tmp2083 := Call(__e, PrimFunc(symshen_4insert), V5619, MakeString("~S"))


__e.TailApply(tmp2079, tmp2083)
return


}, 1)

tmp2084 := Call(__e, ns2_1set, symprint, tmp2078)


_ = tmp2084

tmp2085 := MakeNative(func(__e *ControlFlow) {
V5622 := __e.Get(1)
_ = V5622
V5623 := __e.Get(2)
_ = V5623
tmp2090 := PrimValue(sym_dhush_d)

if True == tmp2090 {
__e.Return(V5622)
return
} else {
tmp2088 := Call(__e, PrimFunc(symshen_4char_1stoutput_2), V5623)


if True == tmp2088 {
__e.TailApply(PrimFunc(symshen_4write_1string), V5622, V5623)
return
} else {
tmp2086 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5622, MakeNumber(0))


__e.TailApply(PrimFunc(symshen_4write_1chars), V5622, V5623, tmp2086, MakeNumber(1))
return


}


}


}, 2)

tmp2091 := Call(__e, ns2_1set, sympr, tmp2085)


_ = tmp2091

tmp2092 := MakeNative(func(__e *ControlFlow) {
V5624 := __e.Get(1)
_ = V5624
V5625 := __e.Get(2)
_ = V5625
tmp2093 := MakeNative(func(__e *ControlFlow) {
tmp2094 := PrimPos(V5624, V5625)

__e.Return(PrimStringToNumber(tmp2094))
return


}, 0)

tmp2095 := MakeNative(func(__e *ControlFlow) {
Z5626 := __e.Get(1)
_ = Z5626
__e.Return(symshen_4eos)
return
}, 1)

__e.TailApply(try_1catch, tmp2093, tmp2095)
return


}, 2)

tmp2096 := Call(__e, ns2_1set, symshen_4string_1_6byte, tmp2092)


_ = tmp2096

tmp2097 := MakeNative(func(__e *ControlFlow) {
V5627 := __e.Get(1)
_ = V5627
V5628 := __e.Get(2)
_ = V5628
V5629 := __e.Get(3)
_ = V5629
V5630 := __e.Get(4)
_ = V5630
tmp2102 := PrimEqual(symshen_4eos, V5629)

if True == tmp2102 {
__e.Return(V5627)
return
} else {
tmp2098 := PrimWriteByte(V5629, V5628)

_ = tmp2098

tmp2099 := Call(__e, PrimFunc(symshen_4string_1_6byte), V5627, V5630)


tmp2100 := PrimNumberAdd(V5630, MakeNumber(1))

__e.TailApply(PrimFunc(symshen_4write_1chars), V5627, V5628, tmp2099, tmp2100)
return


}


}, 4)

tmp2103 := Call(__e, ns2_1set, symshen_4write_1chars, tmp2097)


_ = tmp2103

tmp2104 := MakeNative(func(__e *ControlFlow) {
V5631 := __e.Get(1)
_ = V5631
V5632 := __e.Get(2)
_ = V5632
tmp2109 := PrimIsString(V5631)

if True == tmp2109 {
tmp2105 := Call(__e, PrimFunc(symshen_4proc_1nl), V5631)


__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2105, V5632)
return


} else {
tmp2106 := PrimCons(V5631, Nil)

tmp2107 := PrimCons(symshen_4proc_1nl, tmp2106)

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2107, V5632)
return


}


}, 2)

tmp2110 := Call(__e, ns2_1set, symshen_4mkstr, tmp2104)


_ = tmp2110

tmp2111 := MakeNative(func(__e *ControlFlow) {
V5637 := __e.Get(1)
_ = V5637
V5638 := __e.Get(2)
_ = V5638
tmp2118 := PrimEqual(Nil, V5638)

if True == tmp2118 {
__e.Return(V5637)
return
} else {
tmp2116 := PrimIsPair(V5638)

if True == tmp2116 {
tmp2112 := PrimHead(V5638)

tmp2113 := Call(__e, PrimFunc(symshen_4insert_1l), tmp2112, V5637)


tmp2114 := PrimTail(V5638)

__e.TailApply(PrimFunc(symshen_4mkstr_1l), tmp2113, tmp2114)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-l")))
return
}


}


}, 2)

tmp2119 := Call(__e, ns2_1set, symshen_4mkstr_1l, tmp2111)


_ = tmp2119

tmp2120 := MakeNative(func(__e *ControlFlow) {
V5645 := __e.Get(1)
_ = V5645
V5646 := __e.Get(2)
_ = V5646
tmp2258 := PrimEqual(MakeString(""), V5646)

if True == tmp2258 {
__e.Return(MakeString(""))
return
} else {
tmp2256 := Call(__e, PrimFunc(symshen_4_7string_2), V5646)


var ifres2243 Obj

if True == tmp2256 {
tmp2254 := Call(__e, PrimFunc(symhdstr), V5646)


tmp2255 := PrimEqual(MakeString("~"), tmp2254)

var ifres2245 Obj

if True == tmp2255 {
tmp2252 := PrimTailString(V5646)

tmp2253 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2252)


var ifres2247 Obj

if True == tmp2253 {
tmp2249 := PrimTailString(V5646)

tmp2250 := Call(__e, PrimFunc(symhdstr), tmp2249)


tmp2251 := PrimEqual(MakeString("A"), tmp2250)

var ifres2248 Obj

if True == tmp2251 {
ifres2248 = True


} else {
ifres2248 = False


}

ifres2247 = ifres2248


} else {
ifres2247 = False


}

var ifres2246 Obj

if True == ifres2247 {
ifres2246 = True


} else {
ifres2246 = False


}

ifres2245 = ifres2246


} else {
ifres2245 = False


}

var ifres2244 Obj

if True == ifres2245 {
ifres2244 = True


} else {
ifres2244 = False


}

ifres2243 = ifres2244


} else {
ifres2243 = False


}

if True == ifres2243 {
tmp2121 := PrimTailString(V5646)

tmp2122 := PrimTailString(tmp2121)

tmp2123 := PrimCons(symshen_4a, Nil)

tmp2124 := PrimCons(tmp2122, tmp2123)

tmp2125 := PrimCons(V5645, tmp2124)

__e.Return(PrimCons(symshen_4app, tmp2125))
return


} else {
tmp2241 := Call(__e, PrimFunc(symshen_4_7string_2), V5646)


var ifres2228 Obj

if True == tmp2241 {
tmp2239 := Call(__e, PrimFunc(symhdstr), V5646)


tmp2240 := PrimEqual(MakeString("~"), tmp2239)

var ifres2230 Obj

if True == tmp2240 {
tmp2237 := PrimTailString(V5646)

tmp2238 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2237)


var ifres2232 Obj

if True == tmp2238 {
tmp2234 := PrimTailString(V5646)

tmp2235 := Call(__e, PrimFunc(symhdstr), tmp2234)


tmp2236 := PrimEqual(MakeString("R"), tmp2235)

var ifres2233 Obj

if True == tmp2236 {
ifres2233 = True


} else {
ifres2233 = False


}

ifres2232 = ifres2233


} else {
ifres2232 = False


}

var ifres2231 Obj

if True == ifres2232 {
ifres2231 = True


} else {
ifres2231 = False


}

ifres2230 = ifres2231


} else {
ifres2230 = False


}

var ifres2229 Obj

if True == ifres2230 {
ifres2229 = True


} else {
ifres2229 = False


}

ifres2228 = ifres2229


} else {
ifres2228 = False


}

if True == ifres2228 {
tmp2126 := PrimTailString(V5646)

tmp2127 := PrimTailString(tmp2126)

tmp2128 := PrimCons(symshen_4r, Nil)

tmp2129 := PrimCons(tmp2127, tmp2128)

tmp2130 := PrimCons(V5645, tmp2129)

__e.Return(PrimCons(symshen_4app, tmp2130))
return


} else {
tmp2226 := Call(__e, PrimFunc(symshen_4_7string_2), V5646)


var ifres2213 Obj

if True == tmp2226 {
tmp2224 := Call(__e, PrimFunc(symhdstr), V5646)


tmp2225 := PrimEqual(MakeString("~"), tmp2224)

var ifres2215 Obj

if True == tmp2225 {
tmp2222 := PrimTailString(V5646)

tmp2223 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2222)


var ifres2217 Obj

if True == tmp2223 {
tmp2219 := PrimTailString(V5646)

tmp2220 := Call(__e, PrimFunc(symhdstr), tmp2219)


tmp2221 := PrimEqual(MakeString("S"), tmp2220)

var ifres2218 Obj

if True == tmp2221 {
ifres2218 = True


} else {
ifres2218 = False


}

ifres2217 = ifres2218


} else {
ifres2217 = False


}

var ifres2216 Obj

if True == ifres2217 {
ifres2216 = True


} else {
ifres2216 = False


}

ifres2215 = ifres2216


} else {
ifres2215 = False


}

var ifres2214 Obj

if True == ifres2215 {
ifres2214 = True


} else {
ifres2214 = False


}

ifres2213 = ifres2214


} else {
ifres2213 = False


}

if True == ifres2213 {
tmp2131 := PrimTailString(V5646)

tmp2132 := PrimTailString(tmp2131)

tmp2133 := PrimCons(symshen_4s, Nil)

tmp2134 := PrimCons(tmp2132, tmp2133)

tmp2135 := PrimCons(V5645, tmp2134)

__e.Return(PrimCons(symshen_4app, tmp2135))
return


} else {
tmp2211 := Call(__e, PrimFunc(symshen_4_7string_2), V5646)


if True == tmp2211 {
tmp2136 := Call(__e, PrimFunc(symhdstr), V5646)


tmp2137 := PrimTailString(V5646)

tmp2138 := Call(__e, PrimFunc(symshen_4insert_1l), V5645, tmp2137)


tmp2139 := PrimCons(tmp2138, Nil)

tmp2140 := PrimCons(tmp2136, tmp2139)

tmp2141 := PrimCons(symcn, tmp2140)

__e.TailApply(PrimFunc(symshen_4factor_1cn), tmp2141)
return


} else {
tmp2209 := PrimIsPair(V5646)

var ifres2190 Obj

if True == tmp2209 {
tmp2207 := PrimHead(V5646)

tmp2208 := PrimEqual(symcn, tmp2207)

var ifres2192 Obj

if True == tmp2208 {
tmp2205 := PrimTail(V5646)

tmp2206 := PrimIsPair(tmp2205)

var ifres2194 Obj

if True == tmp2206 {
tmp2202 := PrimTail(V5646)

tmp2203 := PrimTail(tmp2202)

tmp2204 := PrimIsPair(tmp2203)

var ifres2196 Obj

if True == tmp2204 {
tmp2198 := PrimTail(V5646)

tmp2199 := PrimTail(tmp2198)

tmp2200 := PrimTail(tmp2199)

tmp2201 := PrimEqual(Nil, tmp2200)

var ifres2197 Obj

if True == tmp2201 {
ifres2197 = True


} else {
ifres2197 = False


}

ifres2196 = ifres2197


} else {
ifres2196 = False


}

var ifres2195 Obj

if True == ifres2196 {
ifres2195 = True


} else {
ifres2195 = False


}

ifres2194 = ifres2195


} else {
ifres2194 = False


}

var ifres2193 Obj

if True == ifres2194 {
ifres2193 = True


} else {
ifres2193 = False


}

ifres2192 = ifres2193


} else {
ifres2192 = False


}

var ifres2191 Obj

if True == ifres2192 {
ifres2191 = True


} else {
ifres2191 = False


}

ifres2190 = ifres2191


} else {
ifres2190 = False


}

if True == ifres2190 {
tmp2142 := PrimTail(V5646)

tmp2143 := PrimHead(tmp2142)

tmp2144 := PrimTail(V5646)

tmp2145 := PrimTail(tmp2144)

tmp2146 := PrimHead(tmp2145)

tmp2147 := Call(__e, PrimFunc(symshen_4insert_1l), V5645, tmp2146)


tmp2148 := PrimCons(tmp2147, Nil)

tmp2149 := PrimCons(tmp2143, tmp2148)

__e.Return(PrimCons(symcn, tmp2149))
return


} else {
tmp2188 := PrimIsPair(V5646)

var ifres2162 Obj

if True == tmp2188 {
tmp2186 := PrimHead(V5646)

tmp2187 := PrimEqual(symshen_4app, tmp2186)

var ifres2164 Obj

if True == tmp2187 {
tmp2184 := PrimTail(V5646)

tmp2185 := PrimIsPair(tmp2184)

var ifres2166 Obj

if True == tmp2185 {
tmp2181 := PrimTail(V5646)

tmp2182 := PrimTail(tmp2181)

tmp2183 := PrimIsPair(tmp2182)

var ifres2168 Obj

if True == tmp2183 {
tmp2177 := PrimTail(V5646)

tmp2178 := PrimTail(tmp2177)

tmp2179 := PrimTail(tmp2178)

tmp2180 := PrimIsPair(tmp2179)

var ifres2170 Obj

if True == tmp2180 {
tmp2172 := PrimTail(V5646)

tmp2173 := PrimTail(tmp2172)

tmp2174 := PrimTail(tmp2173)

tmp2175 := PrimTail(tmp2174)

tmp2176 := PrimEqual(Nil, tmp2175)

var ifres2171 Obj

if True == tmp2176 {
ifres2171 = True


} else {
ifres2171 = False


}

ifres2170 = ifres2171


} else {
ifres2170 = False


}

var ifres2169 Obj

if True == ifres2170 {
ifres2169 = True


} else {
ifres2169 = False


}

ifres2168 = ifres2169


} else {
ifres2168 = False


}

var ifres2167 Obj

if True == ifres2168 {
ifres2167 = True


} else {
ifres2167 = False


}

ifres2166 = ifres2167


} else {
ifres2166 = False


}

var ifres2165 Obj

if True == ifres2166 {
ifres2165 = True


} else {
ifres2165 = False


}

ifres2164 = ifres2165


} else {
ifres2164 = False


}

var ifres2163 Obj

if True == ifres2164 {
ifres2163 = True


} else {
ifres2163 = False


}

ifres2162 = ifres2163


} else {
ifres2162 = False


}

if True == ifres2162 {
tmp2150 := PrimTail(V5646)

tmp2151 := PrimHead(tmp2150)

tmp2152 := PrimTail(V5646)

tmp2153 := PrimTail(tmp2152)

tmp2154 := PrimHead(tmp2153)

tmp2155 := Call(__e, PrimFunc(symshen_4insert_1l), V5645, tmp2154)


tmp2156 := PrimTail(V5646)

tmp2157 := PrimTail(tmp2156)

tmp2158 := PrimTail(tmp2157)

tmp2159 := PrimCons(tmp2155, tmp2158)

tmp2160 := PrimCons(tmp2151, tmp2159)

__e.Return(PrimCons(symshen_4app, tmp2160))
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

tmp2259 := Call(__e, ns2_1set, symshen_4insert_1l, tmp2120)


_ = tmp2259

tmp2260 := MakeNative(func(__e *ControlFlow) {
V5647 := __e.Get(1)
_ = V5647
tmp2345 := PrimIsPair(V5647)

var ifres2276 Obj

if True == tmp2345 {
tmp2343 := PrimHead(V5647)

tmp2344 := PrimEqual(symcn, tmp2343)

var ifres2278 Obj

if True == tmp2344 {
tmp2341 := PrimTail(V5647)

tmp2342 := PrimIsPair(tmp2341)

var ifres2280 Obj

if True == tmp2342 {
tmp2338 := PrimTail(V5647)

tmp2339 := PrimTail(tmp2338)

tmp2340 := PrimIsPair(tmp2339)

var ifres2282 Obj

if True == tmp2340 {
tmp2334 := PrimTail(V5647)

tmp2335 := PrimTail(tmp2334)

tmp2336 := PrimHead(tmp2335)

tmp2337 := PrimIsPair(tmp2336)

var ifres2284 Obj

if True == tmp2337 {
tmp2329 := PrimTail(V5647)

tmp2330 := PrimTail(tmp2329)

tmp2331 := PrimHead(tmp2330)

tmp2332 := PrimHead(tmp2331)

tmp2333 := PrimEqual(symcn, tmp2332)

var ifres2286 Obj

if True == tmp2333 {
tmp2324 := PrimTail(V5647)

tmp2325 := PrimTail(tmp2324)

tmp2326 := PrimHead(tmp2325)

tmp2327 := PrimTail(tmp2326)

tmp2328 := PrimIsPair(tmp2327)

var ifres2288 Obj

if True == tmp2328 {
tmp2318 := PrimTail(V5647)

tmp2319 := PrimTail(tmp2318)

tmp2320 := PrimHead(tmp2319)

tmp2321 := PrimTail(tmp2320)

tmp2322 := PrimTail(tmp2321)

tmp2323 := PrimIsPair(tmp2322)

var ifres2290 Obj

if True == tmp2323 {
tmp2311 := PrimTail(V5647)

tmp2312 := PrimTail(tmp2311)

tmp2313 := PrimHead(tmp2312)

tmp2314 := PrimTail(tmp2313)

tmp2315 := PrimTail(tmp2314)

tmp2316 := PrimTail(tmp2315)

tmp2317 := PrimEqual(Nil, tmp2316)

var ifres2292 Obj

if True == tmp2317 {
tmp2307 := PrimTail(V5647)

tmp2308 := PrimTail(tmp2307)

tmp2309 := PrimTail(tmp2308)

tmp2310 := PrimEqual(Nil, tmp2309)

var ifres2294 Obj

if True == tmp2310 {
tmp2304 := PrimTail(V5647)

tmp2305 := PrimHead(tmp2304)

tmp2306 := PrimIsString(tmp2305)

var ifres2296 Obj

if True == tmp2306 {
tmp2298 := PrimTail(V5647)

tmp2299 := PrimTail(tmp2298)

tmp2300 := PrimHead(tmp2299)

tmp2301 := PrimTail(tmp2300)

tmp2302 := PrimHead(tmp2301)

tmp2303 := PrimIsString(tmp2302)

var ifres2297 Obj

if True == tmp2303 {
ifres2297 = True


} else {
ifres2297 = False


}

ifres2296 = ifres2297


} else {
ifres2296 = False


}

var ifres2295 Obj

if True == ifres2296 {
ifres2295 = True


} else {
ifres2295 = False


}

ifres2294 = ifres2295


} else {
ifres2294 = False


}

var ifres2293 Obj

if True == ifres2294 {
ifres2293 = True


} else {
ifres2293 = False


}

ifres2292 = ifres2293


} else {
ifres2292 = False


}

var ifres2291 Obj

if True == ifres2292 {
ifres2291 = True


} else {
ifres2291 = False


}

ifres2290 = ifres2291


} else {
ifres2290 = False


}

var ifres2289 Obj

if True == ifres2290 {
ifres2289 = True


} else {
ifres2289 = False


}

ifres2288 = ifres2289


} else {
ifres2288 = False


}

var ifres2287 Obj

if True == ifres2288 {
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

var ifres2281 Obj

if True == ifres2282 {
ifres2281 = True


} else {
ifres2281 = False


}

ifres2280 = ifres2281


} else {
ifres2280 = False


}

var ifres2279 Obj

if True == ifres2280 {
ifres2279 = True


} else {
ifres2279 = False


}

ifres2278 = ifres2279


} else {
ifres2278 = False


}

var ifres2277 Obj

if True == ifres2278 {
ifres2277 = True


} else {
ifres2277 = False


}

ifres2276 = ifres2277


} else {
ifres2276 = False


}

if True == ifres2276 {
tmp2261 := PrimTail(V5647)

tmp2262 := PrimHead(tmp2261)

tmp2263 := PrimTail(V5647)

tmp2264 := PrimTail(tmp2263)

tmp2265 := PrimHead(tmp2264)

tmp2266 := PrimTail(tmp2265)

tmp2267 := PrimHead(tmp2266)

tmp2268 := PrimStringConcat(tmp2262, tmp2267)

tmp2269 := PrimTail(V5647)

tmp2270 := PrimTail(tmp2269)

tmp2271 := PrimHead(tmp2270)

tmp2272 := PrimTail(tmp2271)

tmp2273 := PrimTail(tmp2272)

tmp2274 := PrimCons(tmp2268, tmp2273)

__e.Return(PrimCons(symcn, tmp2274))
return


} else {
__e.Return(V5647)
return
}


}, 1)

tmp2346 := Call(__e, ns2_1set, symshen_4factor_1cn, tmp2260)


_ = tmp2346

tmp2347 := MakeNative(func(__e *ControlFlow) {
V5650 := __e.Get(1)
_ = V5650
tmp2373 := PrimEqual(MakeString(""), V5650)

if True == tmp2373 {
__e.Return(MakeString(""))
return
} else {
tmp2371 := Call(__e, PrimFunc(symshen_4_7string_2), V5650)


var ifres2358 Obj

if True == tmp2371 {
tmp2369 := Call(__e, PrimFunc(symhdstr), V5650)


tmp2370 := PrimEqual(MakeString("~"), tmp2369)

var ifres2360 Obj

if True == tmp2370 {
tmp2367 := PrimTailString(V5650)

tmp2368 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2367)


var ifres2362 Obj

if True == tmp2368 {
tmp2364 := PrimTailString(V5650)

tmp2365 := Call(__e, PrimFunc(symhdstr), tmp2364)


tmp2366 := PrimEqual(MakeString("%"), tmp2365)

var ifres2363 Obj

if True == tmp2366 {
ifres2363 = True


} else {
ifres2363 = False


}

ifres2362 = ifres2363


} else {
ifres2362 = False


}

var ifres2361 Obj

if True == ifres2362 {
ifres2361 = True


} else {
ifres2361 = False


}

ifres2360 = ifres2361


} else {
ifres2360 = False


}

var ifres2359 Obj

if True == ifres2360 {
ifres2359 = True


} else {
ifres2359 = False


}

ifres2358 = ifres2359


} else {
ifres2358 = False


}

if True == ifres2358 {
tmp2348 := PrimNumberToString(MakeNumber(10))

tmp2349 := PrimTailString(V5650)

tmp2350 := PrimTailString(tmp2349)

tmp2351 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp2350)


__e.Return(PrimStringConcat(tmp2348, tmp2351))
return


} else {
tmp2356 := Call(__e, PrimFunc(symshen_4_7string_2), V5650)


if True == tmp2356 {
tmp2352 := Call(__e, PrimFunc(symhdstr), V5650)


tmp2353 := PrimTailString(V5650)

tmp2354 := Call(__e, PrimFunc(symshen_4proc_1nl), tmp2353)


__e.Return(PrimStringConcat(tmp2352, tmp2354))
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.proc-nl")))
return
}


}


}


}, 1)

tmp2374 := Call(__e, ns2_1set, symshen_4proc_1nl, tmp2347)


_ = tmp2374

tmp2375 := MakeNative(func(__e *ControlFlow) {
V5655 := __e.Get(1)
_ = V5655
V5656 := __e.Get(2)
_ = V5656
tmp2384 := PrimEqual(Nil, V5656)

if True == tmp2384 {
__e.Return(V5655)
return
} else {
tmp2382 := PrimIsPair(V5656)

if True == tmp2382 {
tmp2376 := PrimHead(V5656)

tmp2377 := PrimCons(V5655, Nil)

tmp2378 := PrimCons(tmp2376, tmp2377)

tmp2379 := PrimCons(symshen_4insert, tmp2378)

tmp2380 := PrimTail(V5656)

__e.TailApply(PrimFunc(symshen_4mkstr_1r), tmp2379, tmp2380)
return


} else {
__e.Return(PrimSimpleError(MakeString("implementation error in shen.mkstr-r")))
return
}


}


}, 2)

tmp2385 := Call(__e, ns2_1set, symshen_4mkstr_1r, tmp2375)


_ = tmp2385

tmp2386 := MakeNative(func(__e *ControlFlow) {
V5657 := __e.Get(1)
_ = V5657
V5658 := __e.Get(2)
_ = V5658
__e.TailApply(PrimFunc(symshen_4insert_1h), V5657, V5658, MakeString(""))
return
}, 2)

tmp2387 := Call(__e, ns2_1set, symshen_4insert, tmp2386)


_ = tmp2387

tmp2388 := MakeNative(func(__e *ControlFlow) {
V5667 := __e.Get(1)
_ = V5667
V5668 := __e.Get(2)
_ = V5668
V5669 := __e.Get(3)
_ = V5669
tmp2449 := PrimEqual(MakeString(""), V5668)

if True == tmp2449 {
__e.Return(V5669)
return
} else {
tmp2447 := Call(__e, PrimFunc(symshen_4_7string_2), V5668)


var ifres2434 Obj

if True == tmp2447 {
tmp2445 := Call(__e, PrimFunc(symhdstr), V5668)


tmp2446 := PrimEqual(MakeString("~"), tmp2445)

var ifres2436 Obj

if True == tmp2446 {
tmp2443 := PrimTailString(V5668)

tmp2444 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2443)


var ifres2438 Obj

if True == tmp2444 {
tmp2440 := PrimTailString(V5668)

tmp2441 := Call(__e, PrimFunc(symhdstr), tmp2440)


tmp2442 := PrimEqual(MakeString("A"), tmp2441)

var ifres2439 Obj

if True == tmp2442 {
ifres2439 = True


} else {
ifres2439 = False


}

ifres2438 = ifres2439


} else {
ifres2438 = False


}

var ifres2437 Obj

if True == ifres2438 {
ifres2437 = True


} else {
ifres2437 = False


}

ifres2436 = ifres2437


} else {
ifres2436 = False


}

var ifres2435 Obj

if True == ifres2436 {
ifres2435 = True


} else {
ifres2435 = False


}

ifres2434 = ifres2435


} else {
ifres2434 = False


}

if True == ifres2434 {
tmp2389 := PrimTailString(V5668)

tmp2390 := PrimTailString(tmp2389)

tmp2391 := Call(__e, PrimFunc(symshen_4app), V5667, tmp2390, symshen_4a)


__e.Return(PrimStringConcat(V5669, tmp2391))
return


} else {
tmp2432 := Call(__e, PrimFunc(symshen_4_7string_2), V5668)


var ifres2419 Obj

if True == tmp2432 {
tmp2430 := Call(__e, PrimFunc(symhdstr), V5668)


tmp2431 := PrimEqual(MakeString("~"), tmp2430)

var ifres2421 Obj

if True == tmp2431 {
tmp2428 := PrimTailString(V5668)

tmp2429 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2428)


var ifres2423 Obj

if True == tmp2429 {
tmp2425 := PrimTailString(V5668)

tmp2426 := Call(__e, PrimFunc(symhdstr), tmp2425)


tmp2427 := PrimEqual(MakeString("R"), tmp2426)

var ifres2424 Obj

if True == tmp2427 {
ifres2424 = True


} else {
ifres2424 = False


}

ifres2423 = ifres2424


} else {
ifres2423 = False


}

var ifres2422 Obj

if True == ifres2423 {
ifres2422 = True


} else {
ifres2422 = False


}

ifres2421 = ifres2422


} else {
ifres2421 = False


}

var ifres2420 Obj

if True == ifres2421 {
ifres2420 = True


} else {
ifres2420 = False


}

ifres2419 = ifres2420


} else {
ifres2419 = False


}

if True == ifres2419 {
tmp2392 := PrimTailString(V5668)

tmp2393 := PrimTailString(tmp2392)

tmp2394 := Call(__e, PrimFunc(symshen_4app), V5667, tmp2393, symshen_4r)


__e.Return(PrimStringConcat(V5669, tmp2394))
return


} else {
tmp2417 := Call(__e, PrimFunc(symshen_4_7string_2), V5668)


var ifres2404 Obj

if True == tmp2417 {
tmp2415 := Call(__e, PrimFunc(symhdstr), V5668)


tmp2416 := PrimEqual(MakeString("~"), tmp2415)

var ifres2406 Obj

if True == tmp2416 {
tmp2413 := PrimTailString(V5668)

tmp2414 := Call(__e, PrimFunc(symshen_4_7string_2), tmp2413)


var ifres2408 Obj

if True == tmp2414 {
tmp2410 := PrimTailString(V5668)

tmp2411 := Call(__e, PrimFunc(symhdstr), tmp2410)


tmp2412 := PrimEqual(MakeString("S"), tmp2411)

var ifres2409 Obj

if True == tmp2412 {
ifres2409 = True


} else {
ifres2409 = False


}

ifres2408 = ifres2409


} else {
ifres2408 = False


}

var ifres2407 Obj

if True == ifres2408 {
ifres2407 = True


} else {
ifres2407 = False


}

ifres2406 = ifres2407


} else {
ifres2406 = False


}

var ifres2405 Obj

if True == ifres2406 {
ifres2405 = True


} else {
ifres2405 = False


}

ifres2404 = ifres2405


} else {
ifres2404 = False


}

if True == ifres2404 {
tmp2395 := PrimTailString(V5668)

tmp2396 := PrimTailString(tmp2395)

tmp2397 := Call(__e, PrimFunc(symshen_4app), V5667, tmp2396, symshen_4s)


__e.Return(PrimStringConcat(V5669, tmp2397))
return


} else {
tmp2402 := Call(__e, PrimFunc(symshen_4_7string_2), V5668)


if True == tmp2402 {
tmp2398 := PrimTailString(V5668)

tmp2399 := Call(__e, PrimFunc(symhdstr), V5668)


tmp2400 := PrimStringConcat(V5669, tmp2399)

__e.TailApply(PrimFunc(symshen_4insert_1h), V5667, tmp2398, tmp2400)
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

tmp2450 := Call(__e, ns2_1set, symshen_4insert_1h, tmp2388)


_ = tmp2450

tmp2451 := MakeNative(func(__e *ControlFlow) {
V5670 := __e.Get(1)
_ = V5670
V5671 := __e.Get(2)
_ = V5671
V5672 := __e.Get(3)
_ = V5672
tmp2452 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5670, V5672)


__e.Return(PrimStringConcat(tmp2452, V5671))
return


}, 3)

tmp2453 := Call(__e, ns2_1set, symshen_4app, tmp2451)


_ = tmp2453

tmp2454 := MakeNative(func(__e *ControlFlow) {
V5676 := __e.Get(1)
_ = V5676
V5677 := __e.Get(2)
_ = V5677
tmp2462 := Call(__e, PrimFunc(symfail))


tmp2463 := PrimEqual(V5676, tmp2462)

if True == tmp2463 {
__e.Return(MakeString("..."))
return
} else {
tmp2460 := Call(__e, PrimFunc(symshen_4list_2), V5676)


if True == tmp2460 {
__e.TailApply(PrimFunc(symshen_4list_1_6str), V5676, V5677)
return
} else {
tmp2458 := PrimIsString(V5676)

if True == tmp2458 {
__e.TailApply(PrimFunc(symshen_4str_1_6str), V5676, V5677)
return
} else {
tmp2456 := PrimIsVector(V5676)

if True == tmp2456 {
__e.TailApply(PrimFunc(symshen_4vector_1_6str), V5676, V5677)
return
} else {
__e.TailApply(PrimFunc(symshen_4atom_1_6str), V5676)
return
}


}


}


}


}, 2)

tmp2464 := Call(__e, ns2_1set, symshen_4arg_1_6str, tmp2454)


_ = tmp2464

tmp2465 := MakeNative(func(__e *ControlFlow) {
V5678 := __e.Get(1)
_ = V5678
V5679 := __e.Get(2)
_ = V5679
tmp2473 := PrimEqual(symshen_4r, V5679)

if True == tmp2473 {
tmp2466 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2467 := Call(__e, PrimFunc(symshen_4iter_1list), V5678, symshen_4r, tmp2466)


tmp2468 := Call(__e, PrimFunc(sym_8s), tmp2467, MakeString(")"))


__e.TailApply(PrimFunc(sym_8s), MakeString("("), tmp2468)
return


} else {
tmp2469 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2470 := Call(__e, PrimFunc(symshen_4iter_1list), V5678, V5679, tmp2469)


tmp2471 := Call(__e, PrimFunc(sym_8s), tmp2470, MakeString("]"))


__e.TailApply(PrimFunc(sym_8s), MakeString("["), tmp2471)
return


}


}, 2)

tmp2474 := Call(__e, ns2_1set, symshen_4list_1_6str, tmp2465)


_ = tmp2474

tmp2475 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimValue(sym_dmaximum_1print_1sequence_1size_d))
return
}, 0)

tmp2476 := Call(__e, ns2_1set, symshen_4maxseq, tmp2475)


_ = tmp2476

tmp2477 := MakeNative(func(__e *ControlFlow) {
V5690 := __e.Get(1)
_ = V5690
V5691 := __e.Get(2)
_ = V5691
V5692 := __e.Get(3)
_ = V5692
tmp2498 := PrimEqual(Nil, V5690)

if True == tmp2498 {
__e.Return(MakeString(""))
return
} else {
tmp2496 := PrimEqual(MakeNumber(0), V5692)

if True == tmp2496 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2494 := PrimIsPair(V5690)

var ifres2490 Obj

if True == tmp2494 {
tmp2492 := PrimTail(V5690)

tmp2493 := PrimEqual(Nil, tmp2492)

var ifres2491 Obj

if True == tmp2493 {
ifres2491 = True


} else {
ifres2491 = False


}

ifres2490 = ifres2491


} else {
ifres2490 = False


}

if True == ifres2490 {
tmp2478 := PrimHead(V5690)

__e.TailApply(PrimFunc(symshen_4arg_1_6str), tmp2478, V5691)
return


} else {
tmp2488 := PrimIsPair(V5690)

if True == tmp2488 {
tmp2479 := PrimHead(V5690)

tmp2480 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2479, V5691)


tmp2481 := PrimTail(V5690)

tmp2482 := PrimNumberSubtract(V5692, MakeNumber(1))

tmp2483 := Call(__e, PrimFunc(symshen_4iter_1list), tmp2481, V5691, tmp2482)


tmp2484 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2483)


__e.TailApply(PrimFunc(sym_8s), tmp2480, tmp2484)
return


} else {
tmp2485 := Call(__e, PrimFunc(symshen_4arg_1_6str), V5690, V5691)


tmp2486 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2485)


__e.TailApply(PrimFunc(sym_8s), MakeString("|"), tmp2486)
return


}


}


}


}


}, 3)

tmp2499 := Call(__e, ns2_1set, symshen_4iter_1list, tmp2477)


_ = tmp2499

tmp2500 := MakeNative(func(__e *ControlFlow) {
V5695 := __e.Get(1)
_ = V5695
V5696 := __e.Get(2)
_ = V5696
tmp2505 := PrimEqual(symshen_4a, V5696)

if True == tmp2505 {
__e.Return(V5695)
return
} else {
tmp2501 := PrimNumberToString(MakeNumber(34))

tmp2502 := PrimNumberToString(MakeNumber(34))

tmp2503 := Call(__e, PrimFunc(sym_8s), V5695, tmp2502)


__e.TailApply(PrimFunc(sym_8s), tmp2501, tmp2503)
return


}


}, 2)

tmp2506 := Call(__e, ns2_1set, symshen_4str_1_6str, tmp2500)


_ = tmp2506

tmp2507 := MakeNative(func(__e *ControlFlow) {
V5697 := __e.Get(1)
_ = V5697
V5698 := __e.Get(2)
_ = V5698
tmp2520 := Call(__e, PrimFunc(symshen_4print_1vector_2), V5697)


if True == tmp2520 {
tmp2508 := PrimVectorGet(V5697, MakeNumber(0))

tmp2509 := Call(__e, PrimFunc(symfn), tmp2508)


__e.TailApply(tmp2509, V5697)
return


} else {
tmp2518 := Call(__e, PrimFunc(symvector_2), V5697)


if True == tmp2518 {
tmp2510 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2511 := Call(__e, PrimFunc(symshen_4iter_1vector), V5697, MakeNumber(1), V5698, tmp2510)


tmp2512 := Call(__e, PrimFunc(sym_8s), tmp2511, MakeString(">"))


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2512)
return


} else {
tmp2513 := Call(__e, PrimFunc(symshen_4maxseq))


tmp2514 := Call(__e, PrimFunc(symshen_4iter_1vector), V5697, MakeNumber(0), V5698, tmp2513)


tmp2515 := Call(__e, PrimFunc(sym_8s), tmp2514, MakeString(">>"))


tmp2516 := Call(__e, PrimFunc(sym_8s), MakeString("<"), tmp2515)


__e.TailApply(PrimFunc(sym_8s), MakeString("<"), tmp2516)
return


}


}


}, 2)

tmp2521 := Call(__e, ns2_1set, symshen_4vector_1_6str, tmp2507)


_ = tmp2521

tmp2522 := MakeNative(func(__e *ControlFlow) {
V5699 := __e.Get(1)
_ = V5699
tmp2523 := MakeNative(func(__e *ControlFlow) {
W5700 := __e.Get(1)
_ = W5700
tmp2530 := PrimEqual(W5700, symshen_4tuple)

if True == tmp2530 {
__e.Return(True)
return
} else {
tmp2528 := PrimEqual(W5700, symshen_4pvar)

if True == tmp2528 {
__e.Return(True)
return
} else {
tmp2525 := PrimIsNumber(W5700)

tmp2526 := PrimNot(tmp2525)

if True == tmp2526 {
__e.TailApply(PrimFunc(symshen_4fbound_2), W5700)
return
} else {
__e.Return(False)
return
}


}


}


}, 1)

tmp2531 := PrimVectorGet(V5699, MakeNumber(0))

__e.TailApply(tmp2523, tmp2531)
return


}, 1)

tmp2532 := Call(__e, ns2_1set, symshen_4print_1vector_2, tmp2522)


_ = tmp2532

tmp2533 := MakeNative(func(__e *ControlFlow) {
V5701 := __e.Get(1)
_ = V5701
tmp2534 := Call(__e, PrimFunc(symarity), V5701)


tmp2535 := PrimEqual(tmp2534, MakeNumber(-1))

__e.Return(PrimNot(tmp2535))
return


}, 1)

tmp2536 := Call(__e, ns2_1set, symshen_4fbound_2, tmp2533)


_ = tmp2536

tmp2537 := MakeNative(func(__e *ControlFlow) {
V5702 := __e.Get(1)
_ = V5702
tmp2538 := PrimVectorGet(V5702, MakeNumber(1))

tmp2539 := PrimVectorGet(V5702, MakeNumber(2))

tmp2540 := Call(__e, PrimFunc(symshen_4app), tmp2539, MakeString(")"), symshen_4s)


tmp2541 := PrimStringConcat(MakeString(" "), tmp2540)

tmp2542 := Call(__e, PrimFunc(symshen_4app), tmp2538, tmp2541, symshen_4s)


__e.Return(PrimStringConcat(MakeString("(@p "), tmp2542))
return


}, 1)

tmp2543 := Call(__e, ns2_1set, symshen_4tuple, tmp2537)


_ = tmp2543

tmp2544 := MakeNative(func(__e *ControlFlow) {
V5709 := __e.Get(1)
_ = V5709
V5710 := __e.Get(2)
_ = V5710
V5711 := __e.Get(3)
_ = V5711
V5712 := __e.Get(4)
_ = V5712
tmp2564 := PrimEqual(MakeNumber(0), V5712)

if True == tmp2564 {
__e.Return(MakeString("... etc"))
return
} else {
tmp2545 := MakeNative(func(__e *ControlFlow) {
W5713 := __e.Get(1)
_ = W5713
tmp2546 := MakeNative(func(__e *ControlFlow) {
W5715 := __e.Get(1)
_ = W5715
tmp2555 := PrimEqual(W5713, symshen_4out_1of_1bounds)

if True == tmp2555 {
__e.Return(MakeString(""))
return
} else {
tmp2553 := PrimEqual(W5715, symshen_4out_1of_1bounds)

if True == tmp2553 {
__e.TailApply(PrimFunc(symshen_4arg_1_6str), W5713, V5711)
return
} else {
tmp2547 := Call(__e, PrimFunc(symshen_4arg_1_6str), W5713, V5711)


tmp2548 := PrimNumberAdd(V5710, MakeNumber(1))

tmp2549 := PrimNumberSubtract(V5712, MakeNumber(1))

tmp2550 := Call(__e, PrimFunc(symshen_4iter_1vector), V5709, tmp2548, V5711, tmp2549)


tmp2551 := Call(__e, PrimFunc(sym_8s), MakeString(" "), tmp2550)


__e.TailApply(PrimFunc(sym_8s), tmp2547, tmp2551)
return


}


}


}, 1)

tmp2556 := MakeNative(func(__e *ControlFlow) {
tmp2557 := PrimNumberAdd(V5710, MakeNumber(1))

__e.Return(PrimVectorGet(V5709, tmp2557))
return


}, 0)

tmp2558 := MakeNative(func(__e *ControlFlow) {
Z5716 := __e.Get(1)
_ = Z5716
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2559 := Call(__e, try_1catch, tmp2556, tmp2558)


__e.TailApply(tmp2546, tmp2559)
return


}, 1)

tmp2560 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimVectorGet(V5709, V5710))
return
}, 0)

tmp2561 := MakeNative(func(__e *ControlFlow) {
Z5714 := __e.Get(1)
_ = Z5714
__e.Return(symshen_4out_1of_1bounds)
return
}, 1)

tmp2562 := Call(__e, try_1catch, tmp2560, tmp2561)


__e.TailApply(tmp2545, tmp2562)
return


}


}, 4)

tmp2565 := Call(__e, ns2_1set, symshen_4iter_1vector, tmp2544)


_ = tmp2565

tmp2566 := MakeNative(func(__e *ControlFlow) {
V5717 := __e.Get(1)
_ = V5717
tmp2567 := MakeNative(func(__e *ControlFlow) {
__e.Return(PrimStr(V5717))
return
}, 0)

tmp2568 := MakeNative(func(__e *ControlFlow) {
Z5718 := __e.Get(1)
_ = Z5718
__e.TailApply(PrimFunc(symshen_4funexstring))
return
}, 1)

__e.TailApply(try_1catch, tmp2567, tmp2568)
return


}, 1)

tmp2569 := Call(__e, ns2_1set, symshen_4atom_1_6str, tmp2566)


_ = tmp2569

tmp2570 := MakeNative(func(__e *ControlFlow) {
tmp2571 := PrimIntern(MakeString("x"))

tmp2572 := Call(__e, PrimFunc(symgensym), tmp2571)


tmp2573 := Call(__e, PrimFunc(symshen_4arg_1_6str), tmp2572, symshen_4a)


tmp2574 := Call(__e, PrimFunc(sym_8s), tmp2573, MakeString("\x11"))


tmp2575 := Call(__e, PrimFunc(sym_8s), MakeString("e"), tmp2574)


tmp2576 := Call(__e, PrimFunc(sym_8s), MakeString("n"), tmp2575)


tmp2577 := Call(__e, PrimFunc(sym_8s), MakeString("u"), tmp2576)


tmp2578 := Call(__e, PrimFunc(sym_8s), MakeString("f"), tmp2577)


__e.TailApply(PrimFunc(sym_8s), MakeString("\x10"), tmp2578)
return


}, 0)

tmp2579 := Call(__e, ns2_1set, symshen_4funexstring, tmp2570)


_ = tmp2579

tmp2580 := MakeNative(func(__e *ControlFlow) {
V5719 := __e.Get(1)
_ = V5719
tmp2584 := Call(__e, PrimFunc(symempty_2), V5719)


if True == tmp2584 {
__e.Return(True)
return
} else {
tmp2582 := PrimIsPair(V5719)

if True == tmp2582 {
__e.Return(True)
return
} else {
__e.Return(False)
return
}


}


}, 1)

__e.TailApply(ns2_1set, symshen_4list_2, tmp2580)
return




}, 0)

