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

(define compile1
        [$symbol V] -> [[iConst V]]
        [$const V] -> [[iConst V]]
        [$var I] -> [[iAccess I]]
        [$if X Y Z] -> (append (compile1 X) [[iJF | (compile1 Y)] | [[iJMP | (compile1 Z)]]])
        [$do X Y] -> (append (compile1 X) [[iPop] | (compile1 Y)])
        [$defun F L] -> (append (compile1 L) [[iConst F] [iDefun]])
        [$native | L] -> (append (compile-arg-list L) [[iNativeCall (length L)]])
        [$app F | X] -> [[iMark] | (append (compile-apply F X) [[iApply]])]
        [$abs Body] -> [[iFreeze | [[iGrab] | (compile-tail Body)]]]
        [$freeze Body] -> [[iFreeze | (compile-tail Body)]]
        [$trap X Y] -> (append (compile1 Y) [[iSetJmp | (append (compile1 X) [[iClearJmp]])]]))

(define compile-tail
        [$if X Y Z] -> (append (compile1 X) [[iJF | (compile-tail Y)] | [[iJMP | (compile-tail Z)]]])
        [$do X Y] -> (append (compile1 X) [[iPop] | (compile-tail Y)])
        [$abs Body] -> [[iGrab] | (compile-tail Body)]
        [$app F | X] -> (append (compile-apply F X) [[iTailApply]])
        X -> (append (compile1 X) [[iReturn]]))

(define compile-apply
  F X -> (append (compile-arg-list (reverse X)) (compile-function F)))

(define compile-function
  [$symbol V] -> [[iConst V] [iGetF]]
  F -> (compile1 F))

(define compile-arg-list
  ArgList -> (mapcan (function compile1) ArgList))

(define kl->bytecode
        X -> (append (compile-tail (de-bruijn X)) [[iHalt]]))
