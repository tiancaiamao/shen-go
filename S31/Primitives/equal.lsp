"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN shen.equal? (X Y) 
   (IF (shen.ABSEQUAL X Y) 'true 'false))

(DEFUN shen.ABSEQUAL (X Y)
  (COND ((AND (CONSP X) (CONSP Y) (shen.ABSEQUAL (CAR X) (CAR Y))) 
         (shen.ABSEQUAL (CDR X) (CDR Y)))
        ((AND (STRINGP X) (STRINGP Y)) (STRING= X Y))
        ((AND (NUMBERP X) (NUMBERP Y)) (= X Y))
        ((AND (ARRAYP X) (ARRAYP Y)) (CF-VECTORS X Y (LENGTH X) (LENGTH Y)))
        (T (EQUAL X Y))))

(DEFUN CF-VECTORS (X Y LX LY) 
   (AND (= LX LY) 
        (CF-VECTORS-HELP X Y 0 (1- LX))))

(DEFUN CF-VECTORS-HELP (X Y COUNT MAX)
  (COND ((= COUNT MAX) (shen.ABSEQUAL (AREF X MAX) (AREF Y MAX)))
        ((shen.ABSEQUAL (AREF X COUNT) (AREF Y COUNT)) (CF-VECTORS-HELP X Y (1+ COUNT) MAX))
        (T NIL)))