(define init-primitive-table
  [] -> []
  [Prim Arity Str Id | L] -> (do (put Prim primitive.id Id)
                             (do (put Prim primitive.str Str)
                             (do (put Prim primitive.arity Arity)
                                 (init-primitive-table L)))))

(init-primitive-table [
                       + 2 "PrimNumberAdd" 0
                       - 2 "PrimNumberSub" 1
                       = 2 "PrimEqual" 2
                       / 2 "PrimNumberDivide" 3
                       ])

(define primitive?
  X -> (>= (get/or X primitive.id (freeze -1)) 0) where (symbol? X)
  _ -> false)

(define primitive-arity
  X -> (get X primitive.arity))
