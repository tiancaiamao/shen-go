"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN string->n (S) (LET ((L (COERCE S 'LIST)))
                          (IF (= (LIST-LENGTH L) 1)
                              (CHAR-CODE (CAR L))
                              (ERROR "~S is not a unit string.~%" S))))