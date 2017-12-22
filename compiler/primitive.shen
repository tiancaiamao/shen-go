(define primitive?
        X -> (native primitive? X))

(define primitive-arity
        X -> (native primitive-arity X))

(define primitive-id
        X -> (native primitive-id X))


(defun fail () fail!)
(defun shen.f_error (X) (simple-error (str X)))

(define boolean?
  true -> true
  false -> true
  _ -> false)

(define reverse
  X -> (reverse_help X []))

(define reverse_help
  [] R -> R
  [X | Y] R -> (reverse_help Y [X | R]))

(define map
  F X -> (map-h F X []))

(define map-h
  _ [] X -> (reverse X)
  F [X | Y] Z -> (map-h F Y [(F X) | Z]))

(define append
  [] X -> X
  [X | Y] Z -> [X | (append Y Z)])

(define fold-left
  F Acc [] -> Acc
  F Acc [X | Rest] -> (fold-left F (F Acc X) Rest))

(define mapcan
  _ [] -> []
  F [X | Y] -> (append (F X) (mapcan F Y)))

(defun length (V3293) (shen.length-h V3293 0))

(defun shen.length-h (V3296 V3297)
  (cond ((= () V3296) V3297)
        (true (shen.length-h (tl V3296) (+ V3297 1)))))