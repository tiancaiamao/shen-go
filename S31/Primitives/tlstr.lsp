"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"
        
(DEFUN tlstr (X) (trap-error (SUBSEQ X 1) (LAMBDA (E) (ERROR "~S is not a non-empty string~%" X))))