"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN n->string (N) (trap-error (FORMAT NIL "~C" (CODE-CHAR N)) (ERROR "~A is not a natural number~%" N)))