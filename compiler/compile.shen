(define compile1
  [$const V] -> [[iConst V]]
  [$var I] -> [[iAccess I]]
  [$if X Y Z] -> (append (compile1 X) [[iJF | (compile1 Y)] | [[iJMP | (compile1 Z)]]])
  [$do X Y] -> (append (compile1 X) [[iPop] | (compile1 Y)])
  [$defun F L] -> (append (compile1 L) [[iConst F] [iDefun]])
  [$app F | X] -> (compile-apply F X)
  [$abs Body] -> [[iGrab | (append (compile1 Body) [[iReturn]])]]
  [$freeze Body] -> [[iFreeze | (append (compile1 Body) [[iReturn]])]]
  X -> X)

(define compile-apply
  F X -> (if (= (length X) (primitive-arity F))
             (append (mapcan (function compile1) X) [[iPrimCall (primitive-id F)]])
             (compile1 (curry-primitive F X))) where (primitive? F)
  F X -> [[iConst F] [iGetF] | (compile-arg-list X)] where (symbol? F)
  F X -> (append (compile1 F) (compile-arg-list X)))

(define compile-arg-list
  ArgList -> (append (mapcan (/. X (append (compile1 X) [[iPushArg]])) ArgList) [[iApply]]))

(define curry-primitive
  F X -> (let Count (- (primitive-arity F) (length X))
              Pad (rrange Count)
              PadList (map (/. X [$var X]) Pad)
              (fold-left (/. X (/. Y [$abs X])) [$app F | (append X PadList)] Pad)))

(define rrange
  N -> (rrange0 N 0 []))

(define rrange0
  N N R -> R
  N I R -> (rrange0 N (+ I 1) [I | R]))
