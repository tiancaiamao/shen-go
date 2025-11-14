"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN eval-kl (X) 
 (LET ((E (EVAL (cl.kl-to-lisp X))))
      (IF (AND (CONSP X) (EQ (CAR X) 'defun))
          (COMPILE E)
          E)))
                          
