"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN cl.double-precision (X)
  (IF (INTEGERP X) X (COERCE X 'DOUBLE-FLOAT)))

(DECLAIM (INLINE cl.double-precision))

(DEFUN cl.multiply (X Y)
  (IF (OR (ZEROP X) (ZEROP Y))
      0
      (* (cl.double-precision X) (cl.double-precision Y))))

(DEFUN cl.add (X Y) 
  (+ (cl.double-precision X) (cl.double-precision Y)))

(DEFUN cl.subtract (X Y) 
  (- (cl.double-precision X) (cl.double-precision Y)))

(DEFUN cl.divide (X Y) 
  (LET ((Div (/ (cl.double-precision X)
                (cl.double-precision Y))))
                      (IF (INTEGERP Div)
                           Div
                          (* (COERCE 1.0 'DOUBLE-FLOAT) Div))))

(DEFUN cl.greater? (X Y) (IF (> X Y) 'true 'false))

(DEFUN cl.less? (X Y) (IF (< X Y) 'true 'false))

(DEFUN cl.greater-than-or-equal-to? (X Y) 
	(IF (>= X Y) 'true 'false))

(DEFUN cl.less-than-or-equal-to? (X Y) 
	(IF (<= X Y) 'true 'false))

(DEFUN number? (N) (IF (NUMBERP N) 'true 'false))