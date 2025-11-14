"Copyright (c) 2010-2021, Mark Tarver

   3 clause BSD; see license"

(DEFMACRO or (X Y) `(if ,X 'true (if ,Y 'true 'false)))

