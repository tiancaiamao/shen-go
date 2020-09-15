package kl

import (
	"fmt"
	"runtime"
)

func envGet(env Obj, sym Obj) (Obj, bool) {
	for env != Nil {
		pair := car(env)
		if PrimEqual(car(pair), sym) == True {
			return cdr(pair), true
		}
		env = cdr(env)
	}
	return nil, false
}

func envExtend(env Obj, symbols, values []Obj) Obj {
	for i := 0; i < len(symbols); i++ {
		env = cons(cons(symbols[i], values[i]), env)
	}
	return env
}

type ControlFlowKind int

const (
	ControlFlowReturn ControlFlowKind = iota
	ControlFlowEval
	ControlFlowApply
)

const ctlFlowCacheSize = 8

type ControlFlow struct {
	// controlFlowReturn: result = data[0]
	// controlFlowApply: fn, args = data[0], data[1], data[2] ...
	// controlFlowEval: exp, env = data[0], data[1]
	kind  ControlFlowKind

	// len(data) is stack pointer.
	// stack[bp : len(data)] is the arguments to current function.
	stack  []Obj
	bp int
}

func (ctl *ControlFlow) Get(n int) Obj {
	return ctl.stack[ctl.bp + n]
}

// trampoline is introduced for tail call optimization.
func (e *Evaluator) trampoline(exp Obj, env Obj) Obj {
	e.stack = e.stack[:e.bp]
	e.kind = ControlFlowEval
	e.stack = append(e.stack, exp)
	e.stack = append(e.stack, env)
	return e.do(&e.ControlFlow)
}

func (e *Evaluator) Call(f Obj, args ...Obj) Obj {
	// ctl.stack[bp : len(ctl.stack)] is current stack frame
	// 
	// ---------------------------
	//  | bp            |
	//  | arg1 arg2 ... |
	// ---------------------------
	//
	// save old bp, update bp to len(ctl.stack)
	// push arg1, arg2 ...
	// -----------------------------------
	//  | saveBP        | bp           |
	//  | ...           | arg1 arg2 ...|
	// -----------------------------------
	// 
	// Call()
	// recover the old frame: stack <- stack[:bp], bp <-saveBP

	ctl := &e.ControlFlow
	saveBP := ctl.bp

	ctl.bp = len(ctl.stack)
	ctl.stack = append(ctl.stack, f)
	ctl.stack = append(ctl.stack, args...)
	ctl.kind = ControlFlowApply
	ret := e.do(ctl)

	ctl.stack = ctl.stack[:ctl.bp]
	ctl.bp = saveBP
	return ret
}

func (e *Evaluator) do(ctl *ControlFlow) Obj {
	for ctl.kind != ControlFlowReturn {
		switch ctl.kind {
		case ControlFlowEval:
			e.eval(ctl)
		case ControlFlowApply:
			e.apply(ctl)
		}
	}
	ret := ctl.stack[ctl.bp]
	ctl.stack = ctl.stack[:ctl.bp]
	return ret
}

func (ctl *ControlFlow) TailEval(exp Obj, env Obj) {
	ctl.stack = ctl.stack[:ctl.bp]
	ctl.stack = append(ctl.stack, exp)
	ctl.stack = append(ctl.stack, env)
	ctl.kind = ControlFlowEval
}

func (ctl *ControlFlow) TailApply(f Obj, args ...Obj) {
	ctl.stack = ctl.stack[:ctl.bp]
	ctl.stack = append(ctl.stack, f)
	ctl.stack = append(ctl.stack, args...)
	ctl.kind = ControlFlowApply
}

func (ctl *ControlFlow) Return(result Obj) {
	ctl.stack = ctl.stack[:ctl.bp+1]
	ctl.stack[ctl.bp] = result
	ctl.kind = ControlFlowReturn
}

func (e *Evaluator) eval(ctl *ControlFlow) {
	exp := ctl.Get(0)
	env := ctl.Get(1)

	// fmt.Println("eval ===", ObjString(exp), "env ==", ObjString(env), e.stack, e.bp)

	switch *exp { // handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull, scmHeadProcedure, scmHeadPrimitive:
		ctl.Return(exp)
		return
	}

	if ok, _ := isSymbol(exp); ok {
		if val, ok := envGet(env, exp); ok {
			exp = val
		}
		ctl.Return(exp)
		return
	}

	pair := mustPair(exp)
	if ok, sym := isSymbol(pair.car); ok {
		exp = pair.cdr // handle special form
		switch symbolArray[sym.offset].str {
		case "quote":
			// Extension to make vm work.
			// TODO: remove it later
			ctl.Return(car(exp))
			return
		case "defun": // (defun f (x y) z)
			proc := makeProcedure(cadr(exp), caddr(exp), env)
			funName := car(exp)
			e.functionTable[GetSymbol(funName)] = proc
			ctl.Return(funName)
			return
		case "lambda": // (lambda x x)
			ctl.Return(makeProcedure(car(exp), cadr(exp), env))
			return
		case "freeze": // (freeze body)
			ctl.Return(makeProcedure(Nil, car(exp), env))
			return
		case "let": // (let x y z)
			args := e.trampoline(cadr(exp), env)
			newEnv := envExtend(env, []Obj{car(exp)}, []Obj{args})
			ctl.TailEval(caddr(exp), newEnv)
			return
		case "and":
			e.evalAnd(car(exp), cadr(exp), env, ctl)
			return
		case "or":
			e.evalOr(car(exp), cadr(exp), env, ctl)
			return
		case "if": // (if a b c)
			if listLength(pair.cdr) == 3 {
				e.evalIf(car(exp), cadr(exp), caddr(exp), env, ctl)
				return
			} // if may also be a function for partial apply
		case "cond": // (cond (false 1) (true 2))
			e.evalCond(exp, env, ctl)
			return
		case "trap-error": // (trap-error ~body ~handler)
			e.evalTrapError(exp, env, ctl)
			return
		case "do": // (do A A)
			e.trampoline(car(exp), env)
			ctl.TailEval(cadr(exp), env)
			return
		}
	}

	fn := e.evalFunction(pair.car, env)
	args := e.evalArgumentList(pair.cdr, env)
	ctl.TailApply(fn, args...)
	return
}

func (e *Evaluator) evalFunction(fn Obj, env Obj) Obj {
	if ok, _ := isSymbol(fn); ok {
		str := GetSymbol(fn)
		// Native function has higher priority to overload primitive.
		if native, ok := e.nativeFunc[str]; ok {
			return native
		}

		if proc, ok := envGet(env, fn); ok {
			return proc
		}

		if val, ok := e.functionTable[str]; ok {
			return val
		}
	}

	switch *fn {
	case scmHeadPrimitive, scmHeadProcedure, scmHeadNative:
		return fn
	case scmHeadPair:
		return e.trampoline(fn, env)
	}
	panic(fmt.Sprintf("can't apply non function: %#v", (*scmHead)(fn)))
}

func (e *Evaluator) evalIf(a, b, c Obj, env Obj, ctl *ControlFlow) {
	t := e.trampoline(a, env)
	switch t {
	case True:
		ctl.TailEval(b, env)
		return
	case False:
		ctl.TailEval(c, env)
		return
	}
	panic("second argument of if should be boolean")
}

func (e *Evaluator) evalAnd(a, b Obj, env Obj, ctl *ControlFlow) {
	if e.trampoline(a, env) == False {
		ctl.Return(False)
		return
	}
	ctl.TailEval(b, env)
}

func (e *Evaluator) evalOr(a, b Obj, env Obj, ctl *ControlFlow) {
	if e.trampoline(a, env) == True {
		ctl.Return(True)
		return
	}
	ctl.TailEval(b, env)
}

func (e *Evaluator) evalCond(l Obj, env Obj, ctl *ControlFlow) {
	for *l == scmHeadPair {
		curr := car(l)
		if e.trampoline(car(curr), env) == True {
			ctl.TailEval(cadr(curr), env)
			return
		}
		l = cdr(l)
	}
	ctl.Return(Nil)
	return
}

func (e *Evaluator) evalTrapError(exp Obj, env Obj, ctl *ControlFlow) {
	defer func() {
		if err := recover(); err != nil {
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					handle := e.evalFunction(cadr(exp), env)
					ctl.TailApply(handle, val)
					return
				}
			}
			// Unexpected panic?
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Panic:", err)
			fmt.Println("Recovered in trap-error:", ObjString(exp))
			fmt.Println(string(buf[:n]))
		}
	}()
	v := e.trampoline(car(exp), env)
	ctl.Return(v)
	return
}

func (e *Evaluator) apply(ctl *ControlFlow) {
	f := ctl.Get(0)
	args := ctl.stack[ctl.bp+1:]

	if *f == scmHeadPrimitive {
		prim := mustPrimitive(f)
		switch {
		case len(args) < prim.Required:
			ctl.Return(partialApply(prim.Required, args, nil, f))
			return
		case len(args) == prim.Required:
			ret := prim.Function(args...)
			ctl.Return(ret)
			return
		}
	} else if *f == scmHeadProcedure {
		proc := mustProcedure(f)
		switch {
		case len(args) < proc.arity:
			ctl.Return(partialApply(proc.arity, args, proc.env, f))
			return
		case len(args) == proc.arity:
			newEnv := envExtend(proc.env, proc.arg, args)
			ctl.TailEval(proc.body, newEnv)
			return
		case len(args) > proc.arity:
			newEnv := envExtend(proc.env, proc.arg, args[:proc.arity])
			res := e.trampoline(proc.body, newEnv)
			ctl.TailApply(res, args[proc.arity:]...)
			return
		}
	} else if *f == scmHeadNative {
		fn := MustNative(f)
		provided := len(fn.captured) + len(args)
		required := fn.require
		if provided == required {
			var tmp []Obj
			if len(fn.captured) == 0 {
				tmp = args
			} else {
				tmp = make([]Obj, 0, required)
				tmp = append(tmp, fn.captured...)
				tmp = append(tmp, args...)
			}
			fn.fn(e, ctl, tmp...)
			return
		}

		tmp := make([]Obj, 0, fn.require)
		tmp = append(tmp, fn.captured...)
		if provided < required {
			tmp = append(tmp, args...)
			ctl.Return(MakeNative(fn.fn, fn.require, tmp...))
			return
		}

		taken := required - len(tmp)
		tmp = append(tmp, args[:taken]...)
		f = e.Call(f, tmp...)
		args = args[taken:]
		ctl.TailApply(f, args...)
		return
	}
	panic("can't apply object")
}

// partialApply works when Required > providArgs
func partialApply(required int, providArgs []Obj, env Obj, proc Obj) Obj {
	// Partial apply...
	// (f x y z) => (lambda (z) (f x y z)) with x y in env
	symbols := makeTempSymbols(required)
	env1 := envExtend(env, symbols[:len(providArgs)], providArgs)

	args := Nil
	for i, count := len(symbols)-1, required-len(providArgs); count > 0; count-- {
		args = cons(symbols[i], args)
		i--
	}

	body := Nil
	for i := len(symbols) - 1; i >= 0; i-- {
		body = cons(symbols[i], body)
	}
	body = cons(proc, body)

	return makeProcedure(args, body, env1)
}

func (e *Evaluator) evalArgumentList(args Obj, env Obj) []Obj {
	var ret []Obj
	for *args == scmHeadPair {
		v := e.trampoline(car(args), env)
		if *v == scmHeadError {
			return []Obj{v}
		}
		ret = append(ret, v)
		args = cdr(args)
	}
	return ret
}

func makeTempSymbols(n int) []Obj {
	ret := make([]Obj, n)
	for i := 0; i < n; i++ {
		ret[i] = MakeSymbol(fmt.Sprintf("tmp%d", i))
	}
	return ret
}

type tryResult struct {
	e    *Evaluator
	data Obj
}

func (e *Evaluator) Try(f Obj) (res tryResult) {
	defer func() {
		if err := recover(); err != nil {
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					res = tryResult{e: e, data: val}
					return
				}
			}
			// Unexpected panic?
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Panic:", err)
			fmt.Println("Recovered in Try:", ObjString(f))
			fmt.Println(string(buf[:n]))
		}
	}()
	MustNative(f)
	val := e.Call(f)
	res = tryResult{e: e, data: val}
	return
}

func (t tryResult) Catch(f Obj) Obj {
	if IsError(t.data) {
		return t.e.Call(f, t.data)
	}
	return t.data
}
