(systemf internal)
(systemf receive)
(systemf <!>)
(systemf dict)
(systemf dict?)
(systemf dict-count)
(systemf dict->)
(systemf <-dict/or)
(systemf <-dict)
(systemf dict-rm)
(systemf dict-fold)
(systemf dict-keys)
(systemf dict-values)
(systemf <-vector/or)
(systemf <-address/or)
(systemf value/or)
(systemf get/or)
(systemf for-each)
(systemf fold-left)
(systemf fold-right)
(systemf exit)
(systemf filter)
(systemf sterror)
(systemf read-char-code)
(systemf read-file-as-charlist)
(systemf *sterror*)
(systemf *argv*)
(systemf command-line)

(define make-file
  ShenFile -> (let Message (output "compiling ~A~%" ShenFile)
                   Shen (read-file ShenFile)
                   KL (map (function make-kl-code) Shen)
                   BC (map (function make-bc-code) KL)
                   StringBC (list->string BC)
                   BCFile (bcfile ShenFile)
                   Write (write-to-file BCFile StringBC)
                 BCFile))


\* Required to avoid errors when processing functions with system names *\
(defcc shen.<name>
  X := (if (symbol? X)
           X
           (error "~A is not a legitimate function name.~%" X)))

\* Required so that older versions process get/or correctly *\
(define shen.put/get-macro
  [put X Pointer Y] -> [put X Pointer Y [value *property-vector*]]
  [get X Pointer] -> [get X Pointer [value *property-vector*]]
  [get/or X Pointer Or] -> [get/or X Pointer Or [value *property-vector*]]
  [unput X Pointer] -> [unput X Pointer [value *property-vector*]]
  X -> X)

(define make-kl-code
  [define F | Rules] -> (shen.elim-def [define F | Rules])
  [defcc F | Rules] -> (shen.elim-def [defcc F | Rules])
  Code -> Code)

(define make-bc-code
        X -> (kl->bytecode X))

(define bcfile
  ".shen" -> ".bc"
  (@s S Ss) -> (@s S (bcfile Ss)))

(define list->string
  [] -> ""
  \* Needed for 19.2, doesn't prefix `shen` correctly *\
  [[defun shen | Args] | Body] -> (list->string [[defun shen.shen | Args] | Body])
  \* shen.fail! prints as "...", needs to be handled separately *\
  [[defun fail | _] | Y] -> (@s "(defun fail () shen.fail!)" (list->string Y))
  [X | Y] -> (@s (make-string "~R~%~%" X) (list->string Y)))

(load "compiler/primitive.shen")
(load "compiler/de-bruijn.shen")
(load "compiler/compile.shen")
(set *maximum-print-sequence-size* 2000)

(map (function make-file)
     ["ShenOSKernel-20.1/sources/toplevel.shen"
      "ShenOSKernel-20.1/sources/core.shen"
      "ShenOSKernel-20.1/sources/sys.shen"
      "ShenOSKernel-20.1/sources/sequent.shen"
      "ShenOSKernel-20.1/sources/yacc.shen"
      "ShenOSKernel-20.1/sources/reader.shen"
      "ShenOSKernel-20.1/sources/prolog.shen"
      "ShenOSKernel-20.1/sources/track.shen"
      "ShenOSKernel-20.1/sources/load.shen"
      "ShenOSKernel-20.1/sources/writer.shen"
      "ShenOSKernel-20.1/sources/macros.shen"
      "ShenOSKernel-20.1/sources/declarations.shen"
      "ShenOSKernel-20.1/sources/t-star.shen"
      "ShenOSKernel-20.1/sources/types.shen"
      "compiler/primitive.shen"
      "compiler/de-bruijn.shen"
      "compiler/compile.shen"])