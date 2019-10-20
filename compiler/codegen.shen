(set *builtin* [load-file type get-time eval-kl close open read-byte write-byte absvector? <-address address-> absvector str <= >= < > error-to-string simple-error = - * / + string->n n->string number? string? pos tlstr cn intern hd tl cons cons? value set not if symbol? read-file-as-bytelist read-file-as-string variable? integer?])

(define parse
  Env [] -> [$const []]
  Env X -> [$const X] where (or (number? X) (string? X) (boolean? x))
  Env [let X Y Z] -> [$let X (parse Env Y) (parse (cons X Env) Z)]
  Env [if true Y Z] -> (parse Env Y)
  Env [if false Y Z] -> (parse Env Z)
  Env [if X Y Z] -> [$if (parse Env X) (parse Env Y) (parse Env Z)]
  Env [lambda X Y] -> [$abs [X] (parse [X | Env] Y)]
  Env [freeze Body] -> [$abs [] (parse Env Body)]
  Env [trap-error Body Handler] -> [$try-catch [$abs [] (parse Env Body)] (parse Env Handler)]
  Env [defun F X Y] -> [$defun F X (parse X Y)]
  Env [do X Y] -> [$do (parse Env X) (parse Env Y)]
  Env [or X Y] -> (parse Env [if X true [if Y true false]])
  Env [and X Y] -> (parse Env [if X [if Y true false] false])
  Env [cond [Case Action]] -> (parse Env [if Case Action [simple-error "no cond match"]])
  Env [cond [Case Action] | More] -> (parse Env [if Case Action [cond | More]])
  Env [F | X] -> [$app (parse-app Env F) | (map (parse Env) X)] where (symbol? F)
  Env [F | X] -> (let GenF (gensym f)
                      [$let GenF (parse Env F)
                            [$app [$var GenF] | (map (parse Env) X)]])
  Env X -> (if (element? X Env) [$var X] [$const X]))

(define parse-app
  Env F -> [$var F] where (element? F Env)
  _ F -> [$builtin F] where (element? F (value *builtin*))
  _ F -> [$global F])

(define handle-return
  false Return BC Reg -> (Return BC Reg)
  true Return BC Reg -> (Return (cons [$return Reg] BC) _))

(define code-gen
  BC Tail [$const X] Return -> (let Reg (gensym reg)
                                    (handle-return Tail Return (cons [$const X Reg] BC) Reg))
  BC Tail [$var X] Return -> (handle-return Tail Return BC X)
  BC Tail [$if X Y Z] Return ->
  (code-gen BC false X
            (/. BC1 (/. R1
                        (code-gen [] Tail Y
                                  (/. BCy (/. R2
                                              (code-gen [] Tail Z
                                                        (/. BCz (/. R3
                                                                    (if Tail
                                                                        (Return (cons [if R1 (reverse BCy) (reverse BCz)] BC1) _)
                                                                        (let Reg (gensym reg)
                                                                             BC2 (cons [$declare Reg] BC1)
                                                                             (Return (cons [if R1
                                                                                           (reverse (cons [mov R2 Reg] BCy))
                                                                                           (reverse (cons [mov R3 Reg] BCz))]
                                                                                           BC2) Reg))))))))))))
  BC Tail [$do X Y] Return -> (code-gen BC false X
                                        (/. BC1 (/. X1
                                            (let BC2 (cons [mov X1 _] BC1)
                                                 (code-gen BC2 Tail Y Return)))))
  BC Tail [$abs Args Body] Return -> (let Reg (gensym reg)
                                          Body2 (code-gen [] true Body
                                                          (/. Body1 (/. Reg1
                                                                        (reverse Body1))))
                                         (handle-return Tail Return
                                                        (cons [$lambda Reg Args Body2] BC)
                                                        Reg))
  BC Tail [$let X Y Z] Return -> (code-gen BC false Y
                                           (/. BC1 (/. Y1
                                                       (code-gen (cons [mov X _] (cons [$declare-mov Y1 X] BC1)) Tail Z Return))))
  BC Tail [$try-catch Exp Handle] Return -> (code-gen BC false Exp (/. BC1 (/. Exp1
                                                      (code-gen BC1 false Handle (/. BC2 (/. Handle1
                                                                (let Reg (gensym reg)
                                                                     (handle-return Tail Return
                                                                                    (cons [$try-catch Exp1 Handle1 Reg] BC2) Reg))))))))
  BC Tail [$defun F Params Body] Return -> [$defun F Params (code-gen-body Body)]
  BC Tail [$app F | Args] Return ->
      (code-gen-list [] BC Args
                      (/.  BC1 (/. RegList
                                   (let Reg (gensym reg)
                                        Self (code-gen-call BC1 Tail F RegList Reg)
                                        (Return Self (if Tail _ Reg)))))))

(define code-gen-body
  Code -> (code-gen [] true Code
                    (/. BC (/. Reg
                               (reverse (if (= _ Reg) BC (cons [$return Reg] BC)))))))

(define code-gen-list
  Regs BC [] Return -> (Return BC (reverse Regs))
  Regs BC [X | Y] Return -> (code-gen BC false X
                                          (/. BC1 (/. R1
                                            (code-gen-list [R1 | Regs] BC1
                                            Y Return)))))

(define code-gen-call
  BC Tail [$builtin F] Args Reg -> (let X (cons [$builtin F Args Reg] BC)
                                        (if Tail (cons [$return Reg] X) X))
  BC false [$global F] Args Reg -> (cons [$call-def F Args Reg] BC)
  BC true [$global F] Args Reg -> (cons [$tailcall-def F Args] BC)
  BC false [$var F] Args Reg -> (cons [$call F Args Reg] BC)
  BC true [$var F] Args Reg -> (cons [$tailcall F Args] BC))

(define compile-defun
  X -> (let Input (parse [] X)
            (code-gen [] true Input (/. BC (/. Reg BC)))))

(define compile-file-content
  Content -> (let Fn (/. X (and (cons? X) (= (hd X) defun)))
                  Defuns (filter Fn Content)
                  (map compile-defun Defuns)))

(set *maximum-print-sequence-size* 2000)

(define list->string
  [] -> ""
  \* Needed for 19.2, doesn't prefix `shen` correctly *\
  [[defun shen | Args] | Body] -> (list->string [[defun shen.shen | Args] | Body])
  \* shen.fail! prints as "...", needs to be handled separately *\
  [[defun fail | _] | Y] -> (@s "(defun fail () shen.fail!)" (list->string Y))
  [X | Y] -> (@s (make-string "~R~%~%" X) (list->string Y)))

(define compile-file
  File TO -> (let Content (read-file File)
               Data (compile-file-content Content)
               (write-to-file TO (list->string Data))))

(define bcfile
  ".kl" -> ".bc"
  (@s S Ss) -> (@s S (bcfile Ss)))