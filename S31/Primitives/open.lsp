"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN open (String Direction)
   (LET ((Path (FORMAT NIL "~A~A" *home-directory* String)))
        (shen.openh Path Direction)))              

(DEFUN shen.openh (Path Direction) 
      (COND ((EQ Direction 'in) 
             (OPEN Path :DIRECTION :INPUT 
                        :ELEMENT-TYPE '(UNSIGNED-BYTE 8))) 
            ((EQ Direction 'out) 
             (OPEN Path :DIRECTION :OUTPUT 
                        :ELEMENT-TYPE '(UNSIGNED-BYTE 8)
                        :IF-EXISTS :SUPERSEDE)) 
            (T (ERROR "invalid direction"))))