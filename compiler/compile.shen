(define compile1
  Tail [$symbol V] -> [[iConst V]]
  Tail [$const V] -> [[iConst V]]
  Tail [$var I] -> [[iAccess I]]
  Tail [$if X Y Z] -> (append (compile1 false X) [[iJF | (compile1 Tail Y)] | [[iJMP | (compile1 Tail Z)]]])
  Tail [$do X Y] -> (append (compile1 false X) [[iPop] | (compile1 Tail Y)])
  Tail [$defun F L] -> (append (compile1 Tail L) [[iConst F] [iDefun]])
  Tail [$app F | X] -> (compile-apply Tail F X)
  Tail [$abs Body] -> [[iGrab | (append (compile1 Tail Body) [[iReturn]])]]
  Tail [$freeze Body] -> [[iFreeze | (append (compile1 Tail Body) [[iReturn]])]]
  Tail [$trap X Y] -> [[iFreeze | (compile1 Tail Y)] | [[iSetJmp | (append (compile1 false X) [[iClearJmp]])]]]
  Tail X -> X)

(define compile-apply
  Tail [$symbol F] X -> (if (= (length X) (primitive-arity F))
                  (append (compile-arg-list X) [[iPrimCall (primitive-id F)]])
                  (compile1 Tail (curry-primitive F X)))
              where (primitive? F)
  Tail F [] -> (apply-function Tail F)
  false F X -> [[iMark] | (append (compile-arg-list (reverse X)) (apply-function false F))]
  true F X -> (append (compile-arg-list (reverse X)) (apply-function true F)))

(define apply-function
        Tail [$symbol F] -> [[iConst F] [iGetF] (if Tail [iTailApply] [iApply])]
        Tail F -> (compile1 false F))

(define compile-arg-list
  ArgList -> (mapcan (compile1 false) ArgList))

(define curry-primitive
  F X -> (let Count (- (primitive-arity F) (length X))
              Pad (rrange Count)
              PadList (map (/. X [$var X]) Pad)
              (fold-left (/. X (/. Y [$abs X])) [$app [$symbol F] | (append X PadList)] Pad)))

(define rrange
  N -> (rrange0 N 0 []))

(define rrange0
  N N R -> R
  N I R -> (rrange0 N (+ I 1) [I | R]))


(define kl->bytecode
        X -> (append (compile1 true (de-bruijn X)) [[iHalt]]))
