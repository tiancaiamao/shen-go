(define compile1
  [$const V] -> [[iConst V]]
  [$var I] -> [[iAccess I]]
  [$do X Y] -> (append (compile1 X) (compile1 Y))
  [$defun F L] -> (append (compile1 L) [[iConst F] [iDefun]])
  [$app F | X] -> (compile-apply F X)
  [$abs Body] -> [[iGrab | (append (compile1 Body) [[iReturn]])]]
  X -> X)

(define compile-apply
  F X -> (append (mapcan (function compile1) X) [[iPrimCall (primitive-arity F)]]) where (primitive? F)
  F X -> [[iConst F] [iGetF] | (compile-arg-list X)] where (symbol? F)
  F X -> (append (mapcan (function compile1) X) (compile-arg-list X)))

(define compile-arg-list
  ArgList -> (append (mapcan (/. X [X [iPushArg]]) ArgList) [[iApply]]))
