"Copyright (c) 2021, Mark Tarver, 3 clause BSD (see license)"

(SETF (READTABLE-CASE *READTABLE*) :PRESERVE)

(PROCLAIM '(OPTIMIZE (DEBUG 0) (SPEED 3) (SAFETY 3)))

#+SBCL
(DECLAIM (SB-EXT:MUFFLE-CONDITIONS SB-EXT:COMPILER-NOTE)) 

#+SBCL
(SETF SB-EXT:*MUFFLED-WARNINGS* T) 

#+CLISP
(SETQ *LOAD-VERBOSE* NIL
      *COMPILE-VERBOSE* NIL
      CUSTOM:*COMPILE-WARNINGS* NIL)

(IN-PACKAGE :CL-USER)

(DEFUN objectcode (File)
 (LET* ((ObjectCode (COMPILE-FILE File))
        (Load (LOAD ObjectCode)))
       (DELETE-FILE ObjectCode)))
       
(DEFUN boot (KLFile)
  (LET* ((In (OPEN KLFile :DIRECTION :INPUT 
                          :ELEMENT-TYPE :DEFAULT))
         (SourceCode (readsource In))
         (ObjectCode (MAPCAR 'shen.kl-to-lisp SourceCode))
         (LispFile (FORMAT NIL "~A.lsp" KLFile))
         (Out (OPEN LispFile :DIRECTION :OUTPUT 
                             :ELEMENT-TYPE :DEFAULT
                             :IF-EXISTS :SUPERSEDE))
         (Write (MAPC (LAMBDA (X) (FORMAT Out "~S~%~%" X)) ObjectCode))
         (CloseOut (CLOSE Out))
         (CloseIn  (CLOSE In))
         (Compiled (COMPILE-FILE LispFile))
         (Load (LOAD Compiled))
         (Cleanup (MAPC 'DELETE-FILE (LIST LispFile Compiled))))
        'booted)) 
        
(DEFUN readsource (Stream)
 (LET ((R (READ Stream NIL 'eof!!)))
   (IF (EQ R 'eof!!) 
       (PROG2 (CLOSE Stream) NIL)
       (CONS R (readsource Stream)))))
       
(MAPC 'COMPILE '(objectcode boot readsource))            
      
(MAPC 'objectcode '("backend.lsp"  
"Primitives/trap-error.lsp"        "Primitives/absvector.lsp"         "Primitives/address-get.lsp"       
"Primitives/address-send.lsp"      "Primitives/and.lsp"               "Primitives/arith.lsp"             
"Primitives/char-stoutput.lsp"     "Primitives/char-stinput.lsp"       "Primitives/close.lsp"             
"Primitives/cn.lsp"                "Primitives/cons.lsp"              "Primitives/equal.lsp"
"Primitives/error-to-string.lsp"   "Primitives/eval-kl.lsp"           "Primitives/freeze.lsp"            
"Primitives/get-time.lsp"          "Primitives/globals.lsp"           "Primitives/hd.lsp"                
"Primitives/if.lsp"                "Primitives/intern.lsp"            "Primitives/isabsvector.lsp"       
"Primitives/iscons.lsp"            "Primitives/isstring.lsp"          "Primitives/lambda.lsp"            
"Primitives/let.lsp"               "Primitives/n-to-string.lsp"       "Primitives/open.lsp"              
"Primitives/or.lsp"                "Primitives/pos.lsp"               "Primitives/read-byte.lsp"
"Primitives/read-unit-string.lsp"  "Primitives/set.lsp"               "Primitives/simple-error.lsp"      
"Primitives/str.lsp"               "Primitives/string-to-n.lsp"       "Primitives/tl.lsp"                
"Primitives/tlstr.lsp"             "Primitives/type.lsp"              "Primitives/value.lsp"             
"Primitives/write-byte.lsp"        "Primitives/write-string.lsp"))        

#+CLISP
(MAPC 'DELETE-FILE '("backend.lib"  
"Primitives/trap-error.lib"        "Primitives/absvector.lib"         "Primitives/address-get.lib"       
"Primitives/address-send.lib"      "Primitives/and.lib"               "Primitives/arith.lib"             
"Primitives/char-stoutput.lib"     "Primitives/char-stinput.lib"      "Primitives/close.lib"             
"Primitives/cn.lib"                "Primitives/cons.lib"              "Primitives/equal.lib"
"Primitives/error-to-string.lib"   "Primitives/eval-kl.lib"           "Primitives/freeze.lib"            
"Primitives/get-time.lib"          "Primitives/globals.lib"           "Primitives/hd.lib"                
"Primitives/if.lib"                "Primitives/intern.lib"            "Primitives/isabsvector.lib"       
"Primitives/iscons.lib"            "Primitives/isstring.lib"          "Primitives/lambda.lib"            
"Primitives/let.lib"               "Primitives/n-to-string.lib"       "Primitives/open.lib"              
"Primitives/or.lib"                "Primitives/pos.lib"               "Primitives/read-byte.lib"
"Primitives/read-unit-string.lib"  "Primitives/set.lib"               "Primitives/simple-error.lib"      
"Primitives/str.lib"               "Primitives/string-to-n.lib"       "Primitives/tl.lib"                
"Primitives/tlstr.lib"             "Primitives/type.lib"              "Primitives/value.lib"             
"Primitives/write-byte.lib"        "Primitives/write-string.lib"))  

(MAPC 'boot '("KLambda/sys.kl" "KLambda/writer.kl" "KLambda/core.kl" 
              "KLambda/reader.kl" "KLambda/declarations.kl" 
              "KLambda/toplevel.kl" "KLambda/macros.kl" "KLambda/load.kl" 
              "KLambda/prolog.kl" "KLambda/sequent.kl" "KLambda/track.kl" 
              "KLambda/t-star.kl" "KLambda/yacc.kl" "KLambda/types.kl"))
              
#+CLISP
(MAPC 'DELETE-FILE '("KLambda/sys.kl.lib" "KLambda/writer.kl.lib" "KLambda/core.kl.lib" 
                     "KLambda/reader.kl.lib" "KLambda/declarations.kl.lib" 
                     "KLambda/toplevel.kl.lib" "KLambda/macros.kl.lib" "KLambda/load.kl.lib" 
                     "KLambda/prolog.kl.lib" "KLambda/sequent.kl.lib" "KLambda/track.kl.lib" 
                     "KLambda/t-star.kl.lib" "KLambda/yacc.kl.lib" "KLambda/types.kl.lib"))
                   
(MAPC 'FMAKUNBOUND '(boot readsource objectcode))

#+SBCL
(SAVE-LISP-AND-DIE "sbcl-shen.exe" :EXECUTABLE T :TOPLEVEL 'shen.shen)

#+CLISP
(EXT:SAVEINITMEM "clisp-shen.exe" :EXECUTABLE 0 :QUIET T :INIT-FUNCTION 'shen.shen)

#+CLISP
(BYE)   