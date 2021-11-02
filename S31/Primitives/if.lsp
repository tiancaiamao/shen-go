"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO if (X Y Z)
  `(LET ((*C* ,X))
       (COND ((EQ *C* 'true) ,Y) 
             ((EQ *C* 'false) ,Z)
             (T (ERROR "~S is not a boolean~%" *C*)))))