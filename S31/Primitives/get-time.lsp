"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN get-time (Time) 
  (COND ((EQ Time 'run)
         (* 1.0 (/ (GET-INTERNAL-RUN-TIME) 
                   INTERNAL-TIME-UNITS-PER-SECOND)))
        ((EQ Time 'unix)
                     (- (GET-UNIVERSAL-TIME) 2208988800))
        (T (ERROR "get-time does not understand the parameter ~A~%" Time))))