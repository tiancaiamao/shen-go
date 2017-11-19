(define compile1
  [$const V] -> [iConst V]
  [$var I] -> [iAccess I]
  [$do X Y] -> [(compile1 X) (compile1 Y)]
  [$app F | X] -> (compile-apply F X)
  [$abs Body] -> [iGrab (compile1 Body) iReturn]
  X -> X)

(define compile-apply
  F X -> [(map (function compile1) X) [iPrimCall (primitive-arity F)]] where (primitive? F)
  F X -> (let EvalList (map (function compile1) (cons F X))
              (compile-apply-procedure EvalList)))

(define compile-apply-procedure
  L -> (let Len (length L)
            Arg (compile-apply-procedure0 Len 1 L)
            (cons [iMark Len] (append Arg [[iApply]]))))

(define compile-apply-procedure0
  Len Pos [] -> []
  Len Pos [N | Ns] -> [N [iPushArg (- Len Pos)] | (compile-apply-procedure0 Len (+ Pos 1) Ns)])
