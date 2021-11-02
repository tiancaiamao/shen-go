"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN shen.double-precision (X)
  (IF (INTEGERP X) X (COERCE X 'DOUBLE-FLOAT)))

(DECLAIM (INLINE shen.double-precision))

(DEFUN shen.multiply (X Y)
  (IF (OR (ZEROP X) (ZEROP Y))
      0
      (* (shen.double-precision X) (shen.double-precision Y))))

(DEFUN shen.add (X Y) 
  (+ (shen.double-precision X) (shen.double-precision Y)))

(DEFUN shen.subtract (X Y) 
  (- (shen.double-precision X) (shen.double-precision Y)))

(DEFUN shen.divide (X Y) 
  (LET ((Div (/ (shen.double-precision X)
                (shen.double-precision Y))))
                      (IF (INTEGERP Div)
                           Div
                          (* (COERCE 1.0 'DOUBLE-FLOAT) Div))))

(DEFUN shen.greater? (X Y) (IF (> X Y) 'true 'false))

(DEFUN shen.less? (X Y) (IF (< X Y) 'true 'false))

(DEFUN shen.greater-than-or-equal-to? (X Y) 
	(IF (>= X Y) 'true 'false))

(DEFUN shen.less-than-or-equal-to? (X Y) 
	(IF (<= X Y) 'true 'false))

(DEFUN number? (N) (IF (NUMBERP N) 'true 'false))