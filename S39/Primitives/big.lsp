"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN absvector (N) (MAKE-ARRAY (LIST N) :INITIAL-ELEMENT 'shen.fail!))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN <-address (Vector N) (SVREF Vector N))

"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN address-> (Vector N Value) (SETF (SVREF Vector N) Value) Vector)"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO and (X Y) `(if ,X (if ,Y 'true 'false) 'false))"Copyright (c) 2010-2021, Mark Tarver

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

(DEFUN number? (N) (IF (NUMBERP N) 'true 'false))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

#+SBCL
(DEFUN shen.char-stinput? (Stream) 'false)

#+CLISP
(DEFUN shen.char-stinput? (Stream) 
  (IF (EQ Stream *stinput*) 'true 'false))
"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN shen.char-stoutput? (Stream) 
  (IF (EQ Stream *stoutput*) 'true 'false))
"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN close (Stream) (CLOSE Stream) NIL)"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN cn (Str1 Str2) (CONCATENATE 'STRING Str1 Str2))

"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN cons (X Y) (CONS X Y))"Copyright (c) 2010-2021, Mark Tarver

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
        (T NIL)))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN error-to-string (E) 
   (IF (TYPEP E 'CONDITION) 
       (FORMAT NIL "~A" E)
       (ERROR "~S is not an exception~%" E)))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN eval-kl (X) 
 (LET ((E (EVAL (cl.kl-to-lisp X))))
      (IF (AND (CONSP X) (EQ (CAR X) 'defun))
          (COMPILE E)
          E)))
                          
"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"
 
(DEFMACRO freeze (X) `(FUNCTION (LAMBDA () ,X)))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN get-time (Time) 
  (COND ((EQ Time 'run)
         (* 1.0 (/ (GET-INTERNAL-RUN-TIME) 
                   INTERNAL-TIME-UNITS-PER-SECOND)))
        ((EQ Time 'unix)
                     (- (GET-UNIVERSAL-TIME) 2208988800))
        (T (ERROR "get-time does not understand the parameter ~A~%" Time))))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(SETQ *language* "Common Lisp"
      *implementation* (LISP-IMPLEMENTATION-TYPE)
      *porters* "Mark Tarver"
      *port* 3.2
      *os* (SOFTWARE-TYPE)
      *stinput* *STANDARD-INPUT*
      *stoutput* *STANDARD-OUTPUT*)

#+SBCL      
(SETQ *release* "2.0.0")      

#+CLISP
(SETQ *release* "2.49")"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN hd (X) (CAR X))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO if (X Y Z)
  `(LET ((*C* ,X))
       (COND ((EQ *C* 'true) ,Y) 
             ((EQ *C* 'false) ,Z)
             (T (ERROR "~S is not a boolean~%" *C*)))))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN intern (String) (INTERN String))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN absvector? (X) (IF (ARRAYP X) 'true 'false))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN cons? (X) (IF (CONSP X) 'true 'false))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN string? (S) (IF (STRINGP S) 'true 'false))
"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"
   
(DEFMACRO lambda (X Y) `(FUNCTION (LAMBDA (,X) ,Y)))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO let (X Y Z) `(LET ((,X ,Y)) ,Z))Copyright (c) 2010-2021, Mark Tarver

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. The name of Mark Tarver may not be used to endorse or promote products
   derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY
EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.c#34;"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN n->string (N) (trap-error (FORMAT NIL "~C" (CODE-CHAR N)) (ERROR "~A is not a natural number~%" N)))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN open (String Direction)
   (LET ((Path (FORMAT NIL "~A~A" *home-directory* String)))
        (cl.openh Path Direction)))              

(DEFUN cl.openh (Path Direction) 
      (COND ((EQ Direction 'in) 
             (OPEN Path :DIRECTION :INPUT 
                        :ELEMENT-TYPE '(UNSIGNED-BYTE 8))) 
            ((EQ Direction 'out) 
             (OPEN Path :DIRECTION :OUTPUT 
                        :ELEMENT-TYPE '(UNSIGNED-BYTE 8)
                        :IF-EXISTS :SUPERSEDE)) 
            (T (ERROR "invalid direction"))))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO or (X Y) `(if ,X 'true (if ,Y 'true 'false)))

"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN pos (X N) (trap-error (COERCE (LIST (CHAR X N)) 'STRING)
                       (LAMBDA (E)
                         (IF (NOT (STRINGP X)) 
                             (ERROR "~A is not a string~%" X)
                             (ERROR "~A is not a natural number less than the length of the string~%" 
                                    N)))))
 "Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN read-byte (S) 
  (READ-BYTE S NIL -1))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"
  
(DEFUN shen.read-unit-string (Stream) (COERCE (LIST (READ-CHAR Stream)) 'STRING))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN set (X Y) (SET X Y))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN simple-error (String) (ERROR "~A" String))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN str (X) 
  (COND ((NULL X) (ERROR "[] is not an atom in cl; str cannot convert it to a string.~%"))
        ((SYMBOLP X) 
         (cl.process-string (SYMBOL-NAME X)))
        ((NUMBERP X) 
         (cl.process-number (FORMAT NIL "~A" X)))
        ((STRINGP X) (FORMAT NIL "~S" X))
        ((STREAMP X) (FORMAT NIL "~A" X))
        ((FUNCTIONP X) (FORMAT NIL "~A" X)) 
        (T (ERROR "~S is not an atom, stream or closure; str cannot convert it to a string.~%" X))))

(DEFUN cl.process-number (S)
  (COND ((STRING-EQUAL S "") "")
        ((STRING-EQUAL (pos S 0) "d")
         (IF (STRING-EQUAL (pos S 1) "0")
             ""
             (cn "e" (tlstr S))))
        (T (cn (pos S 0) (cl.process-number (tlstr S))))))

(DEFUN cl.process-string (X)
   (COND ((STRING-EQUAL X "") X)
         ((AND (> (LENGTH X) 8) 
               (STRING-EQUAL X "_hash1957" :END1 9)) 
          (cn "#" (cl.process-string (SUBSEQ X 9))))
         ((AND (> (LENGTH X) 9) 
               (STRING-EQUAL X "_quote1957" :END1 10)) 
          (cn "'" (cl.process-string (SUBSEQ X 10))))
         ((AND (> (LENGTH X) 13) 
               (STRING-EQUAL X "_backquote1957" :END1 14)) 
          (cn "`" (cl.process-string (SUBSEQ X 14))))
         ((AND (> (LENGTH X) 7) 
               (STRING-EQUAL X "bar!1957" :END1 8))
          (cn "|" (cl.process-string (SUBSEQ X 8))))
         (T (cn (pos X 0) (cl.process-string (tlstr X))))))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN string->n (S) (LET ((L (COERCE S 'LIST)))
                          (IF (= (LIST-LENGTH L) 1)
                              (CHAR-CODE (CAR L))
                              (ERROR "~S is not a unit string.~%" S))))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN tl (X) (CDR X))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"
        
(DEFUN tlstr (X) (trap-error (SUBSEQ X 1) (LAMBDA (E) (ERROR "~S is not a non-empty string~%" X))))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO trap-error (X F) 
`(HANDLER-CASE ,X (ERROR (Condition) (FUNCALL ,F Condition))))
"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN type (X MyType) (DECLARE (IGNORE MyType)) X)

"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN value (X) (SYMBOL-VALUE X))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN write-byte (Byte S) (WRITE-BYTE Byte S))"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFUN shen.write-string (String Stream) 
 (WRITE-STRING String Stream) 
 (FORCE-OUTPUT Stream) 
 String) 

