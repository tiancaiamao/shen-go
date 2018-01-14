(systemf native)
(systemf plugin-bind)
(systemf *package-path*)

(define import
	Pkg -> (load (@s (value *package-path*) "/pkg/" Pkg "/package.shen")))
