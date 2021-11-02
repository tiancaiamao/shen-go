"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN address-> (Vector N Value) (SETF (SVREF Vector N) Value) Vector)