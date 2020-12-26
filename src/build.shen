(define parse
  Env [] -> [$const []]
  Env X -> [$const X] where (or (number? X) (string? X) (boolean? x))
  Env [if true Y Z] -> (parse Env Y)
  Env [if false Y Z] -> (parse Env Z)
  Env [if X Y Z] -> [$if (parse Env X) (parse Env Y) (parse Env Z)]
  Env [do X Y] -> [$do (parse Env X) (parse Env Y)]
  Env [lambda X Y] -> [$lambda [X] (parse [X | Env] Y)]

  Env [defun F X Y] -> [[$global defun] [$const F] [$lambda X (parse X Y)]]
  Env [let X Y Z] -> [$let X (parse Env Y) (parse (cons X Env) Z)]
  Env [trap-error Body Handler] -> [$try-catch [$lambda [] (parse Env Body)] (parse Env Handler)]


  Env [or X Y] -> (parse Env [if X true [if Y true false]])
  Env [and X Y] -> (parse Env [if X [if Y true false] false])
  Env [cond [Case Action]] -> (parse Env [if Case Action [simple-error "no cond match"]])
  Env [cond [Case Action] | More] -> (parse Env [if Case Action [cond | More]])
  Env [freeze Body] -> [$lambda [] (parse Env Body)]

  Env [F | X] -> [(parse-app Env F) | (map (parse Env) X)] where (symbol? F)
  Env [F | X] -> (let GenF (gensym f)
                      [$let GenF (parse Env F)
                            [GenF | (map (parse Env) X)]])
  Env X -> (if (element? X Env) X [$const X]))

(define parse-app
  Env F -> F where (element? F Env)
  _ F -> [$global F])


(define shen-to-kl
  Code -> (shen.elim-def (shen.walk (function macroexpand) Code)))

(define symbol->string
  S -> "|{|" where (= { S)
  S -> "|}|" where (= } S)
  S -> "|;|" where (= ; S)
  S -> (str S))

(define concat-strings
  [] -> ""
  [S | Ss] -> (@s S " " (concat-strings Ss)))

(define sexp->string
  true -> "#t"
  false -> "#f"
  Comma -> "|,|" where (= Comma ,)
  Sym -> (symbol->string Sym) where (symbol? Sym)
  [quote Exp] -> (@s "'" (sexp->string Exp))
  [Sexp | Sexps] -> (@s "(" (concat-strings (map (/. X (sexp->string X))
                                                 [Sexp | Sexps]))
                        ")")
  Sexp -> (make-string "~R" Sexp))

(define codegen
  Src -> (let Ast (parse [] Src)
      	      Norm ((cora. anf) Ast (cora. id))
	      ((cora. expr->stmt) [] Norm)))

(define compile-file
  Fin Fout -> (let Exprs (read-file Fin)
      	      	   BC (map (codegen) Exprs)
		   ((cora. write-sexp-to-file) Fout BC)))


((cora. load-so) "init.so")
((cora. load-so) "compiler.so")