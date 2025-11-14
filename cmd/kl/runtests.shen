(set *saved-home* (value *home-directory*))
(set *home-directory* (cn (value *saved-home*) "/S31/Test Programs"))
\\(cd "../../S31/Test Programs")
(load "runme.shen")
