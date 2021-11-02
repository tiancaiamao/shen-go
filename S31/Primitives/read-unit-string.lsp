"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"
  
(DEFUN shen.read-unit-string (Stream) (COERCE (LIST (READ-CHAR Stream)) 'STRING))