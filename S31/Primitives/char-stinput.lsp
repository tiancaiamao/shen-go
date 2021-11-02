"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

#+SBCL
(DEFUN shen.char-stinput? (Stream) 'false)

#+CLISP
(DEFUN shen.char-stinput? (Stream) 
  (IF (EQ Stream *stinput*) 'true 'false))
