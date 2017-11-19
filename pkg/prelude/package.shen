(package prelude []

         (plugin-bind "prelude"
                      [bit-shift 2 "BitLeftShift"])

         (define fact
           0 -> 1
           N -> (* N (fact (- N 1))))

         )
