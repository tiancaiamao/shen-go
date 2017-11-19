(load "primitive.shen")
(load "de-bruijn.shen")
(load "compile.shen")

(define kl->bytecode
        X -> (compile1 (de-bruijn X)))