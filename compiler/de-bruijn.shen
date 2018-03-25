\*

Copyright (C) 2018, Arthur Mao <tiancaiamao@gmail.com>

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. The name of Mark Tarver may not be used to endorse or promote products
   derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY Mark Tarver ''AS IS'' AND ANY
EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL Mark Tarver BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

*\

(define de-bruijn
  Exp -> (de-bruijn0 [] Exp))

(define de-bruijn0
  Env [] -> [$const []]
  Env X -> [$const X] where (or (boolean? X) (number? X) (string? X))
  Env [let X Y Z] -> [$app [$abs (de-bruijn0 (cons X Env) Z)] (de-bruijn0 Env Y)]
  Env [lambda X Y] -> [$abs (de-bruijn0 (cons X Env) Y)]
  Env [freeze X] -> [$freeze (de-bruijn0 Env X)]
  Env [if X Y Z] -> [$if (de-bruijn0 Env X) (de-bruijn0 Env Y) (de-bruijn0 Env Z)]
  Env [do X Y] -> [$do (de-bruijn0 Env X) (de-bruijn0 Env Y)]
  Env [defun F X Y] -> [$defun F (de-bruijn0 Env (defun-rewrite X Y))]
  Env [cond | L] -> (de-bruijn0 Env (cond-rewrite L))
  Env [and X Y] -> (de-bruijn0 Env [if X [if Y true false] false])
  Env [or X Y] -> (de-bruijn0 Env [if X true [if Y true false]])
  Env [trap-error X Y] -> [$trap (de-bruijn0 Env X) (de-bruijn0 Env Y)]
  Env [native | L] -> [$native | (map (de-bruijn0 Env) L)]
  Env [F | X] -> [$app (de-bruijn0 Env F) | (map (de-bruijn0 Env) X)]
  Env X -> (de-bruijn-index X Env))

(define de-bruijn-index
  X E <- (find-env X E)
  X _ -> [$symbol X])

(define find-env
  S E -> (find-env0 S 0 E))

(define find-env0
  S I [] -> (fail)
  S I [S | _] -> [$var I]
  S I [_ | E] -> (find-env0 S (+ I 1) E))

(define defun-rewrite
  [] Y -> [freeze Y]
  X Y -> (defun-rewrite0 X Y))

(define defun-rewrite0
  [] Y -> Y
  [X | L] Y -> [lambda X (defun-rewrite0 L Y)])

(define cond-rewrite
  [] -> [simple-error "no match cond"]
  [[X Y] | L] -> [if X Y (cond-rewrite L)])
