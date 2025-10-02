(put shen shen.external-symbols (difference (get shen shen.external-symbols)
                                            [warn included datatypes tracked userdefs tc? spy? step? occurs? factorise? hush? optimise? system-S?] ))
(declare p.cores [--> number])

(declare thread [[lazy A] --> thread])
(declare tc? [--> boolean])
(declare spy? [--> boolean])
(declare step? [--> boolean])
(declare occurs? [--> boolean])
(declare factorise? [--> boolean])
(declare hush? [--> boolean])
(declare optimise? [--> boolean])
(declare system-S? [--> boolean])
(declare userdefs [--> [list symbol]])
(declare tracked  [--> [list symbol]])
(declare datatypes  [--> [list symbol]])
(declare included [--> [list symbol]])
(declare step? [--> boolean])

(define datatypes
  -> (map (/. X (shen.typename X)) (value shen.*alldatatypes*)))

(define included
  -> (map (/. X (shen.typename X)) (value shen.*datatypes*)))
  
(define userdefs
  -> (value shen.*userdefs*))  
  
(define optimise?
  -> (value shen.*optimise*)) 
  
(define hush?
  -> (value *hush*))   
  
(define system-S?
  -> (value shen.*shen-type-theory-enabled?*))   

(define tc?
  -> (value shen.*tc*))
  
(define occurs? 
  -> (value shen.*occurs*))
  
(define factorise?  
  -> (value shen.*factorise?*))
  
(define tracked
  -> (value shen.*tracking*))
  
(define spy?
   -> (value shen.*spy*))
   
(define step?
    -> (value shen.*step*))   
   
(map (fn systemf) [included datatypes tracked userdefs tc? spy? step? occurs? factorise? hush? optimise? system-S? step?])   