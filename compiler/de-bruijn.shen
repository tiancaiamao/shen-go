(define de-bruijn
  Exp -> (de-bruijn0 [] Exp))

(define de-bruijn0
  Env [lambda X Y] -> [$abs (de-bruijn0 (cons X Env) Y)]
  Env [if X Y Z] -> [$if (de-bruijn0 Env X) (de-bruijn0 Env Y) (de-bruijn0 Env Z)]
  Env [do X Y] -> [$do (de-bruijn0 Env X) (de-bruijn0 Env Y)]
  Env [F | X] -> [$app (de-bruijn0 Env F) | (map (de-bruijn0 Env) X)]
  Env X -> (de-bruijn-index X Env))

(define de-bruijn-index
  X _ -> [$const X] where (or (boolean? X) (number? X) (string? X))
  X E <- (find-env X E)
  X _ -> X)

(define find-env
  S E -> (find-env0 S 0 E))

(define find-env0
  S I [] -> (fail)
  S I [S | _] -> [$var I]
  S I [_ | E] -> (find-env0 S (+ I 1) E))
