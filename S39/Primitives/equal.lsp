"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN cl.equal? (X Y) 
   (IF (cl.ABSEQUAL X Y) 'true 'false))

(DEFUN cl.ABSEQUAL (X Y)
  (COND ((AND (CONSP X) (CONSP Y) (cl.ABSEQUAL (CAR X) (CAR Y))) 
         (cl.ABSEQUAL (CDR X) (CDR Y)))
        ((AND (STRINGP X) (STRINGP Y)) (STRING= X Y))
        ((AND (NUMBERP X) (NUMBERP Y)) (= X Y))
        ((AND (ARRAYP X) (ARRAYP Y)) (CF-VECTORS X Y (LENGTH X) (LENGTH Y)))
        (T (EQUAL X Y))))

(DEFUN CF-VECTORS (X Y LX LY) 
   (AND (= LX LY) 
        (CF-VECTORS-HELP X Y 0 (1- LX))))

(DEFUN CF-VECTORS-HELP (X Y COUNT MAX)
  (COND ((= COUNT MAX) (cl.ABSEQUAL (AREF X MAX) (AREF Y MAX)))
        ((cl.ABSEQUAL (AREF X COUNT) (AREF Y COUNT)) (CF-VECTORS-HELP X Y (1+ COUNT) MAX))
        (T NIL)))