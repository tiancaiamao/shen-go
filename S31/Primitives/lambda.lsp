"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"
   
(DEFMACRO lambda (X Y) `(FUNCTION (LAMBDA (,X) ,Y)))