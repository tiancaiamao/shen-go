(define parse
  Env [] -> [$const []]
  Env X -> [$const X] where (or (number? X) (string? X) (boolean? X))
  Env X -> (if (element? X Env) X [$const X]) where (symbol? X)
  Env [if X Y Z] -> [$if (parse Env X) (parse Env Y) (parse Env Z)]
  Env [do X Y] -> [$do (parse Env X) (parse Env Y)]
  Env [lambda X Y] -> [$lambda [X] (parse (cons X Env) Y)]

  Env [defun F X Y] -> [ns2-set [$const F] [$lambda X (parse X Y)]]
  Env [let X Y Z] -> [[$lambda [X] (parse (cons X Env) Z)] (parse Env Y)]
  Env [trap-error Body Handler] -> [try-catch [$lambda [] (parse Env Body)] (parse Env Handler)]

  Env [or X Y] -> (parse Env [if X true [if Y true false]])
  Env [and X Y] -> (parse Env [if X [if Y true false] false])
  Env [cond [true Action]] -> (parse Env Action)
  Env [cond [Case Action]] -> (parse Env [if Case Action [simple-error "no cond match"]])
  Env [cond [Case Action] | More] -> (parse Env [if Case Action [cond | More]])
  Env [freeze Body] -> [$lambda [] (parse Env Body)]
  Env [type Exp _] -> (parse Env Exp)

  Env [F | X] -> [(parse-app Env F) | (map (parse Env) X)] where (symbol? F)
  Env [F | X] -> (map (parse Env) [F | X]))


(define parse-app
  Env F -> F where (element? F Env)
  _ F -> [$global F])


(define peval-t
      BC [$const X] -> (cons [return [$const X]] BC)
      BC [$global X] -> (cons [return [$global X]] BC)
      BC [$if X Y Z] -> (let Y1 (peval0 Y)
			     Z1 (peval0 Z)
			     Tmp (gensym ifres)
			     (peval BC X (/. BC1 (/. REG1
					   (cons [if REG1 Y1 Z1] BC1)))))
      BC [$do X Y] -> (peval BC X (/. BC1 (/. X1
				    (peval-t (cons [ignore X1] BC1) Y))))
      BC [$lambda Args Body] -> (cons [return [lambda Args (peval0 Body)]] BC)
      BC [F | Args] -> (peval-call BC [F | Args] []
				   (/. BC1 (/. REGS
				     (cons [tailapply | REGS] BC1))))
      BC X -> (cons [return X] BC))

(define peval
    BC [$const X] CC -> (CC BC [$const X])
    BC [$global X] CC -> (CC BC [$global X])
    BC [$if X Y Z] CC -> (let TMP (gensym ifres)
			      (let Y1 (peval [] Y (/. BC1 (/. Y1 (cons block (reverse (cons [<- TMP Y1] BC1))))))
				   Z1 (peval [] Z (/. BC1 (/. Z1 (cons block (reverse (cons [<- TMP Z1] BC1))))))
				   (peval BC X (/. BC1 (/. REG1
							   (CC [[if REG1 Y1 Z1] [var TMP] | BC1] TMP))))))
      BC [$do X Y] CC -> (peval BC X (/. BC1 (/. X1
				(peval (cons [ignore X1] BC1) Y CC))))
      BC [$lambda ARGS BODY] CC -> (let TMP (gensym tmp)
					 (CC (cons [<= TMP [lambda ARGS (peval0 BODY)]] BC) TMP))
      BC [F | ARGS] CC -> (peval-call BC [F | ARGS] []
				      (/. BC1 (/. REGS
					(let TMP (gensym tmp)
					     (CC (cons [<= TMP [call | REGS]] BC1) TMP)))))
			      BC X CC -> (CC BC X))

(define peval-call
    BC [] REGS CC -> (CC BC (reverse REGS))
    BC [X | Xs] REGS CC -> (peval BC X (/. BC1 (/. X1
					    (peval-call BC1 Xs (cons X1 REGS) CC)))))

(define peval0
  Exp -> (let RET (reverse (peval-t [] Exp))
       (if (= 1 (length RET))
	   (hd RET)
	 (cons block RET))))

(define codegen
  Input -> (let Ast (parse () Input) (peval0 Ast)))

(define compile-file
    Fin Fout -> (let Expr (read-file Fin)
		     Expr1 (macroexpand [do | Expr])
		     (let BC (codegen Expr1)
			  Str (make-string "~R" BC)
			  (write-to-file Fout Str))))
