"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO let (X Y Z) `(LET ((,X ,Y)) ,Z))