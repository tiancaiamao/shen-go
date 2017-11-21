(define de-bruijn
  Exp -> (de-bruijn0 [] Exp))

(define de-bruijn0
  Env [let X Y Z] -> [$app [$abs (de-bruijn0 (cons X Env) Z)] (de-bruijn0 Env Y)]
  Env [lambda X Y] -> [$abs (de-bruijn0 (cons X Env) Y)]
  Env [freeze X] -> [$freeze (de-bruijn0 Env X)]
  Env [if X Y Z] -> [$if (de-bruijn0 Env X) (de-bruijn0 Env Y) (de-bruijn0 Env Z)]
  Env [do X Y] -> [$do (de-bruijn0 Env X) (de-bruijn0 Env Y)]
  Env [defun F X Y] -> [$defun F (de-bruijn0 Env (defun-rewrite X Y))]
  Env [cond | L] -> (de-bruijn0 Env (cond-rewrite L))
  Env [and X Y] -> [$if (de-bruijn0 Env X) (de-bruijn0 Env Y) [$const false]]
  Env [or X Y] -> [$if (de-bruijn0 Env X) [$const true] (de-bruijn0 Env Y)]
  Env [F | X] -> (let F1 (de-bruijn0 Env)
                      F2 (/. X (if (symbol? X) [$const X] X))
                      [$app (F1 F) | (map (compose [F1 F2]) X)])
  Env X -> (de-bruijn-index X Env))

(define de-bruijn-index
  X _ -> [$const X] where (or (boolean? X) (number? X) (string? X))
  X E <- (find-env X E)
  X _ -> X)

(define compose
  [] X -> X
  [F | Fs] X -> (compose Fs (F X)))

(define find-env
  S E -> (find-env0 S 0 E))

(define find-env0
  S I [] -> (fail)
  S I [S | _] -> [$var I]
  S I [_ | E] -> (find-env0 S (+ I 1) E))

(define defun-rewrite
  [] Y -> [freeze Y]
  X Y -> (defun-rewrite0 X Y))

(define defun-rewrite0
  [] Y -> Y
  [X | L] Y -> [lambda X (defun-rewrite0 L Y)])

(define cond-rewrite
  [] -> [simple-error "no match cond"]
  [[X Y] | L] -> [if X Y (cond-rewrite L)])
