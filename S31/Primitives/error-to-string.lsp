"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN error-to-string (E) 
   (IF (TYPEP E 'CONDITION) 
       (FORMAT NIL "~A" E)
       (ERROR "~S is not an exception~%" E)))