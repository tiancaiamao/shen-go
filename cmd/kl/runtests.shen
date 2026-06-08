(set *saved-home* (value *home-directory*))
(set *home-directory* (cn (value *saved-home*) "/kernel/tests"))
\\(cd "../../kernel/tests")
(load "runme.shen")
