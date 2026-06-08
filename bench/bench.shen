\ Benchmark suite for shen-go compiler work.
\ Measures tak (arithmetic/recursion), fib (recursion), nrev (list), map-double (HOF).
\ Run: printf '(load "bench/bench.shen")\n' | ./shen

(define bench-tak
  X Y Z N ->
    (if (= N 0)
        done
        (let _ (tak X Y Z)
          (bench-tak X Y Z (- N 1)))))

(define tak
  X Y Z ->
    (if (not (< Y X))
        Z
        (tak (tak (- X 1) Y Z)
             (tak (- Y 1) Z X)
             (tak (- Z 1) X Y))))

(define fib
  0 -> 0
  1 -> 1
  N -> (+ (fib (- N 1)) (fib (- N 2))))

(define iota-help
  N Acc -> Acc where (= N 0)
  N Acc -> (iota-help (- N 1) (cons N Acc)))

(define iota
  N -> (iota-help N []))

(define nrev
  [] -> []
  [H | T] -> (append (nrev T) [H]))

(define map-double
  [] -> []
  [H | T] -> (cons (* H 2) (map-double T)))

(define run-bench
  Name Thunk ->
    (let T0 (get-time run)
         _  (Thunk)
         T1 (get-time run)
      (do (output "~A: ~A s~%" Name (- T1 T0))
          (- T1 T0))))

(output "~%--- shen-go benchmarks ---~%")

(run-bench "tak(18,12,6)x200" (freeze (bench-tak 18 12 6 200)))
(run-bench "fib(25)" (freeze (fib 25)))
(let L (iota 1000)
  (run-bench "nrev(1000)" (freeze (nrev L))))
(let L (iota 10000)
  (run-bench "map-double(10000)" (freeze (map-double L))))

(output "~%--- done ---~%")
