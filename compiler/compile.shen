(define compile1
        [$symbol V] -> [[iConst V]]
        [$const V] -> [[iConst V]]
        [$var I] -> [[iAccess I]]
        [$if X Y Z] -> (append (compile1 X) [[iJF | (compile1 Y)] | [[iJMP | (compile1 Z)]]])
        [$do X Y] -> (append (compile1 X) [[iPop] | (compile1 Y)])
        [$defun F L] -> (append (compile1 L) [[iConst F] [iDefun]])
        [$prim F | X] -> (compile-primitive-call F X)
        [$app F | X] -> [[iMark] | (append (compile-apply F X) [[iApply]])]
        [$abs Body] -> [[iFreeze | [[iGrab] | (compile-tail Body)]]]
        [$freeze Body] -> [[iFreeze | (compile-tail Body)]]
        [$trap X Y] -> (append (compile1 Y) [[iSetJmp | (append (compile1 X) [[iClearJmp]])]]))

(define compile-tail
        [$if X Y Z] -> (append (compile1 X) [[iJF | (compile-tail Y)] | [[iJMP | (compile-tail Z)]]])
        [$do X Y] -> (append (compile1 X) [[iPop] | (compile-tail Y)])
        [$abs Body] -> [[iGrab] | (compile-tail Body)]
        [$app F | X] -> (append (compile-apply F X) [[iTailApply]])
        X -> (append (compile1 X) [[iReturn]]))

(define compile-primitive-call
  F X -> (append (compile-arg-list X) [[iPrimCall (primitive-id F)]]) where (= (length X) (primitive-arity F))
  F X -> (compile1 (curry-primitive F X)))

(define compile-apply
  F X -> (append (compile-arg-list (reverse X)) (compile-function F)))

(define compile-function
  [$symbol V] -> [[iConst V] [iGetF]]
  F -> (compile1 F))

(define compile-arg-list
  ArgList -> (mapcan (function compile1) ArgList))

(define curry-primitive
  F X -> (let Count (- (primitive-arity F) (length X))
              Pad (rrange Count)
              PadList (map (/. X [$var X]) Pad)
              (fold-left (/. X (/. Y [$abs X])) [$prim F | (append X PadList)] Pad)))

(define rrange
  N -> (rrange0 N 0 []))

(define rrange0
  N N R -> R
  N I R -> (rrange0 N (+ I 1) [I | R]))

(define kl->bytecode
        X -> (append (compile-tail (de-bruijn X)) [[iHalt]]))
