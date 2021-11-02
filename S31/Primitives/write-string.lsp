"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN shen.write-string (String Stream) 
 (WRITE-STRING String Stream) 
 (FORCE-OUTPUT Stream) 
 String) 

