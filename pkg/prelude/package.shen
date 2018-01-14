(package prelude [bit-shift now time-sub]

        (native plugin-bind (@s (value *package-path*) "/pkg/prelude/prelude.so")
                      [bit-shift 2 "BitLeftShift"
			time-now 0 "Now"
			time-sub 2 "TimeSub"])

	(define bit-shift
		X Y -> (native bit-shift X Y))

	(define fact
	 0 -> 1
	 N -> (* N (fact (- N 1))))

	(define now
		-> (native time-now))

	(define time-sub
		T1 T2 -> (native time-sub T2 T1))
)
