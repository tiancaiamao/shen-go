"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(SETQ *language* "Common Lisp"
      *implementation* (LISP-IMPLEMENTATION-TYPE)
      *porters* "Mark Tarver"
      *port* 3.1
      *os* (SOFTWARE-TYPE)
      *stinput* *STANDARD-INPUT*
      *stoutput* *STANDARD-OUTPUT*)

#+SBCL      
(SETQ *release* "1.3.10")      

#+CLISP
(SETQ *release* "2.49")