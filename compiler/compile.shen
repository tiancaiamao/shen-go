(define compile1
  Tail [$symbol V] -> [[iConst V] | (if Tail [[iReturn]] [])]
  Tail [$const V] -> [[iConst V] | (if Tail [[iReturn]] [])]
  Tail [$var I] -> [[iAccess I] | (if Tail [[iReturn]] [])]
  Tail [$if X Y Z] -> (append (compile1 false X) [[iJF | (compile1 Tail Y)] | [[iJMP | (compile1 Tail Z)]]])
  Tail [$do X Y] -> (append (compile1 false X) [[iPop] | (compile1 Tail Y)])
  Tail [$defun F L] -> (append (compile1 Tail L) [[iConst F] [iDefun] | (if Tail [[iReturn]] [])])
  Tail [$app F | X] -> (compile-apply Tail F X)
  Tail [$abs Body] -> [[iGrab | (append (compile1 Tail Body) [[iReturn]])]]
  Tail [$freeze Body] -> [[iFreeze | (append (compile1 Tail Body) [[iReturn]])] | (if Tail [[iReturn]] [])]
  Tail [$trap X Y] -> [[iSetJmp | (compile1 Tail Y)] | (compile1 Tail X)]
  Tail X -> X)

(define compile-apply
  Tail [$symbol F] X -> (if (= (length X) (primitive-arity F))
                  (append (mapcan (compile1 false) X) [[iPrimCall (primitive-id F)] | (if Tail [[iReturn]] [])])
                  (compile1 Tail (curry-primitive F X)))
              where (primitive? F)
  Tail [$symbol F] X -> [[iConst F] [iGetF] | (append (compile-arg-list X) (if Tail [[iTailApply]] [[iApply]]))]
  Tail F X -> (append (compile1 false F) (append (compile-arg-list X) (if Tail [[iTailApply]] [[iApply]]))))

(define compile-arg-list
  ArgList -> (mapcan (/. X (append (compile1 false X) [[iPushArg]])) ArgList))

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
