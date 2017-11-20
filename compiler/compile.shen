(define compile1
  [$const V] -> [[iConst V]]
  [$var I] -> [[iAccess I]]
  [$do X Y] -> (append (compile1 X) (compile1 Y))
  [$app F | X] -> (compile-apply F X)
  [$abs Body] -> [[iGrab | (append (compile1 Body) [[iReturn]])]]
  X -> X)

(define compile-apply
  F X -> (append (mapcan (function compile1) X) [[iPrimCall (primitive-arity F)]]) where (primitive? F)
  F X -> (let ArgList (mapcan (function compile1) X)
              ArgList1 (mapcan (/. X [X [iPushArg]]) ArgList)
              Result (append ArgList1 [[iApply]])
              (append (compile1 F) Result)))
