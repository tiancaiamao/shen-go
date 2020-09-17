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

	// data[pos : len(data)] is the arguments to current function.
	data  []Obj
	pos int
}

func (e *Evaluator) evalExp(exp Obj, env Obj) Obj {
	e.TailEval(exp, env)
	return e.trampoline()
}

// trampoline is introduced for tail call optimization.
func (e *Evaluator) trampoline() Obj {
	ctl := &e.ControlFlow
	for ctl.kind != ControlFlowReturn {
		switch ctl.kind {
		case ControlFlowEval:
			e.eval()
		case ControlFlowApply:
			e.apply()
		}
	}
	ret := ctl.data[ctl.pos]
	ctl.data = ctl.data[:ctl.pos]
	return ret
}

func (ctl *ControlFlow) TailEval(exp Obj, env Obj) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, exp)
	ctl.data = append(ctl.data, env)
	ctl.kind = ControlFlowEval
}

func (ctl *ControlFlow) TailApply(f Obj, args ...Obj) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, f)
	ctl.data = append(ctl.data, args...)
	ctl.kind = ControlFlowApply
}

func (ctl *ControlFlow) Return(result Obj) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, result)
	ctl.kind = ControlFlowReturn
}

func (e *Evaluator) eval() {
	exp := e.data[e.pos]
	env := e.data[e.pos + 1]

	// fmt.Println("eval ===", ObjString(exp), "env ==", ObjString(env), e.data, e.pos)

	switch *exp { // handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull, scmHeadProcedure, scmHeadPrimitive:
		e.Return(exp)
		return
	}

	if ok, _ := isSymbol(exp); ok {
		if val, ok := envGet(env, exp); ok {
			exp = val
		}
		e.Return(exp)
		return
	}

	pair := mustPair(exp)
	if ok, sym := isSymbol(pair.car); ok {
		exp = pair.cdr // handle special form
		switch symbolArray[sym.offset].str {
		case "quote":
			// Extension to make vm work.
			// TODO: remove it later
			e.Return(car(exp))
			return
		case "defun": // (defun f (x y) z)
			proc := makeProcedure(cadr(exp), caddr(exp), env)
			funName := car(exp)
			e.functionTable[GetSymbol(funName)] = proc
			e.Return(funName)
			return
		case "lambda": // (lambda x x)
			e.Return(makeProcedure(car(exp), cadr(exp), env))
			return
		case "freeze": // (freeze body)
			e.Return(makeProcedure(Nil, car(exp), env))
			return
		case "let": // (let x y z)
			args := e.evalExp(cadr(exp), env)
			newEnv := envExtend(env, []Obj{car(exp)}, []Obj{args})
			e.TailEval(caddr(exp), newEnv)
			return
		case "and":
			e.evalAnd(car(exp), cadr(exp), env)
			return
		case "or":
			e.evalOr(car(exp), cadr(exp), env)
			return
		case "if": // (if a b c)
			if listLength(pair.cdr) == 3 {
				e.evalIf(car(exp), cadr(exp), caddr(exp), env)
				return
			} // if may also be a function for partial apply
		case "cond": // (cond (false 1) (true 2))
			e.evalCond(exp, env)
			return
		case "trap-error": // (trap-error ~body ~handler)
			e.evalTrapError(exp, env)
			return
		case "do": // (do A A)
			e.evalExp(car(exp), env)
			e.TailEval(cadr(exp), env)
			return
		}
	}

	savePOS := e.pos
	fn := e.evalFunction(pair.car, env)
	e.data = e.data[:e.pos]
	e.data = append(e.data, fn)
	e.pos++

	args := pair.cdr
	for *args == scmHeadPair {
		v := e.evalExp(car(args), env)
		if *v == scmHeadError {
			e.pos--
			e.TailApply(fn, v)
			return
		}
		e.data = append(e.data, v)
		e.pos++
		args = cdr(args)
	}
	e.pos = savePOS
	e.kind = ControlFlowApply
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
		return e.evalExp(fn, env)
	}
	panic(fmt.Sprintf("can't apply non function: %#v", (*scmHead)(fn)))
}

func (e *Evaluator) evalIf(a, b, c Obj, env Obj) {
	t := e.evalExp(a, env)
	switch t {
	case True:
		e.TailEval(b, env)
		return
	case False:
		e.TailEval(c, env)
		return
	}
	panic("second argument of if should be boolean")
}

func (e *Evaluator) evalAnd(a, b Obj, env Obj) {
	if e.evalExp(a, env) == False {
		e.Return(False)
		return
	}
	e.TailEval(b, env)
}

func (e *Evaluator) evalOr(a, b Obj, env Obj) {
	if e.evalExp(a, env) == True {
		e.Return(True)
		return
	}
	e.TailEval(b, env)
}

func (e *Evaluator) evalCond(l Obj, env Obj) {
	for *l == scmHeadPair {
		curr := car(l)
		if e.evalExp(car(curr), env) == True {
			e.TailEval(cadr(curr), env)
			return
		}
		l = cdr(l)
	}
	e.Return(Nil)
	return
}

func (e *Evaluator) evalTrapError(exp Obj, env Obj) {
	defer func() {
		if err := recover(); err != nil {
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					handle := e.evalFunction(cadr(exp), env)
					e.TailApply(handle, val)
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
	v := e.evalExp(car(exp), env)
	e.Return(v)
	return
}

func (e *Evaluator) apply() {
	f := e.data[e.pos]
	args := e.data[e.pos+1:]

	if *f == scmHeadPrimitive {
		prim := mustPrimitive(f)
		switch {
		case len(args) < prim.Required:
			e.Return(partialApply(prim.Required, args, Nil, f))
			return
		case len(args) == prim.Required:
			ret := prim.Function(args...)
			e.Return(ret)
			return
		}
	} else if *f == scmHeadProcedure {
		proc := mustProcedure(f)
		switch {
		case len(args) < proc.arity:
			e.Return(partialApply(proc.arity, args, proc.env, f))
			return
		case len(args) == proc.arity:
			newEnv := envExtend(proc.env, proc.arg, args)
			e.TailEval(proc.body, newEnv)
			return
		case len(args) > proc.arity:
			newEnv := envExtend(proc.env, proc.arg, args[:proc.arity])
			res := e.evalExp(proc.body, newEnv)
			e.TailApply(res, args[proc.arity:]...)
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
			fn.fn(e, &e.ControlFlow, tmp...)
			return
		}

		tmp := make([]Obj, 0, fn.require)
		tmp = append(tmp, fn.captured...)
		if provided < required {
			tmp = append(tmp, args...)
			e.Return(MakeNative(fn.fn, fn.require, tmp...))
			return
		}

		taken := required - len(tmp)
		tmp = append(tmp, args[:taken]...)
		f = e.Call(f, tmp...)
		args = args[taken:]
		e.TailApply(f, args...)
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
