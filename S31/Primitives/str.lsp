"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN str (X) 
  (COND ((NULL X) (ERROR "[] is not an atom in Shen; str cannot convert it to a string.~%"))
        ((SYMBOLP X) 
         (shen.process-string (SYMBOL-NAME X)))
        ((NUMBERP X) 
         (shen.process-number (FORMAT NIL "~A" X)))
        ((STRINGP X) (FORMAT NIL "~S" X))
        ((STREAMP X) (FORMAT NIL "~A" X))
        ((FUNCTIONP X) (FORMAT NIL "~A" X)) 
        (T (ERROR "~S is not an atom, stream or closure; str cannot convert it to a string.~%" X))))

(DEFUN shen.process-number (S)
  (COND ((STRING-EQUAL S "") "")
        ((STRING-EQUAL (pos S 0) "d")
         (IF (STRING-EQUAL (pos S 1) "0")
             ""
             (cn "e" (tlstr S))))
        (T (cn (pos S 0) (shen.process-number (tlstr S))))))

(DEFUN shen.process-string (X)
   (COND ((STRING-EQUAL X "") X)
         ((AND (> (LENGTH X) 8) 
               (STRING-EQUAL X "_hash1957" :END1 9)) 
          (cn "#" (shen.process-string (SUBSEQ X 9))))
         ((AND (> (LENGTH X) 9) 
               (STRING-EQUAL X "_quote1957" :END1 10)) 
          (cn "'" (shen.process-string (SUBSEQ X 10))))
         ((AND (> (LENGTH X) 13) 
               (STRING-EQUAL X "_backquote1957" :END1 14)) 
          (cn "`" (shen.process-string (SUBSEQ X 14))))
         ((AND (> (LENGTH X) 7) 
               (STRING-EQUAL X "bar!1957" :END1 8))
          (cn "|" (shen.process-string (SUBSEQ X 8))))
         (T (cn (pos X 0) (shen.process-string (tlstr X))))))