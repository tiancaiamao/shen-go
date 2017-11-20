(define init-primitive-table
  [] -> []
  [Prim Arity Str Id | L] -> (do (put Prim primitive.id Id)
                             (do (put Prim primitive.str Str)
                             (do (put Prim primitive.arity Arity)
                                 (init-primitive-table L)))))

(init-primitive-table [
load-file 1 "xxx" 0
type 2 "xxx" 1
get-time 1 "xxx" 2
eval-kl 1 "xxx" 3
close 1 "xxx" 4
open 2 "xxx" 5
read-byte 1 "xxx" 6
write-byte 2 "xxx" 7
absvector? 1 "xxx" 8
<-address 2 "xxx" 9
address-> 3 "xxx" 10
absvector 1 "xxx" 11
str 1 "xxx" 12
<= 2 "xxx" 13
>= 2 "xxx" 14
< 2 "xxx" 15
> 2 "xxx" 16
error-to-string 1 "xxx" 17
simple-error 1 "xxx" 18
= 2 "xxx" 19
- 2 "xxx" 20
* 2 "xxx" 21
/ 2 "xxx" 22
+ 2 "xxx" 23
string->n 1 "xxx" 24
n->string 1 "xxx" 25
number? 1 "xxx" 26
string? 1 "xxx" 27
pos 2 "xxx" 28
tlstr 1 "xxx" 29
cn 2 "xxx" 30
intern 1 "xxx" 31
hd 1 "xxx" 32
tl 1 "xxx" 33
cons 2 "xxx" 34
cons? 1 "xxx" 35
value 1 "xxx" 36
set 2 "xxx" 37
not 1 "xxx" 38
if 3 "xxx" 39
])

(define primitive?
  X -> (>= (get/or X primitive.id (freeze -1)) 0) where (symbol? X)
  _ -> false)

(define primitive-arity
  X -> (get X primitive.arity))

(define primitive-id
  X -> (get X primitive.id))