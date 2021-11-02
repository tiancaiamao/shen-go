"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN pos (X N) (trap-error (COERCE (LIST (CHAR X N)) 'STRING)
                       (LAMBDA (E)
                         (IF (NOT (STRINGP X)) 
                             (ERROR "~A is not a string~%" X)
                             (ERROR "~A is not a natural number less than the length of the string~%" 
                                    N)))))
 