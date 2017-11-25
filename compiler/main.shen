(load "primitive.shen")
(load "de-bruijn.shen")
(load "compile.shen")

(define kl->bytecode
        X -> (compile1 true (de-bruijn X)))

(define shen-to-kl
  Code -> (shen.elim-def (shen.walk (function macroexpand) Code)))

(set *bytecode-stream* (open "fifo" out))

(define ep
        E Output -> (let KL (shen-to-kl E)
                         BC (kl->bytecode KL)
                         BC1 (append BC [[iHalt]])
                         Str (make-string "~R" BC1)
                         _ (pr Str Output)
                         BC))

(define repl
        ->  (let Input (read)
                 Output (value *bytecode-stream*)
                 _ (pr (trap-error
                        (ep Input Output)
                        (/. X (error-to-string X))))
                 (repl)))