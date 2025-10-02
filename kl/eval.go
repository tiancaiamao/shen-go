package kl

import (
	"fmt"
	"runtime"
)

type ControlFlowKind int

const (
	ControlFlowReturn ControlFlowKind = iota
	ControlFlowEval
	ControlFlowApply
)

type ControlFlow struct {
	// controlFlowReturn: result = data[0]
	// controlFlowApply: fn, args = data[0], data[1], data[2] ...
	// controlFlowEval: exp, env = data[0], data[1]
	kind ControlFlowKind

	// data[pos : len(data)] is the arguments to current function.
	// Why it needs to be a stack?
	// Most of the time, an array is sufficient. But when eval expression like
	// (f (g1 h) (g2 n) ...), the stack is used to store the temporary arguments
	data []Obj
	pos  int
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
	// TODO: Change the function signature.
	// Using args ... is not good for performance, there is a copy here.
	ctl.data = append(ctl.data, args...)
	ctl.kind = ControlFlowApply
}

func (ctl *ControlFlow) Return(result Obj) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, result)
	ctl.kind = ControlFlowReturn
}

func (ctl *ControlFlow) Get(n int) Obj {
	return ctl.data[ctl.pos+n]
}

// trampoline is introduced for tail call optimization.
func trampoline(ctl *ControlFlow) Obj {
	for ctl.kind != ControlFlowReturn {
		switch ctl.kind {
		case ControlFlowEval:
			eval(ctl)
		case ControlFlowApply:
			apply(ctl)
		}
	}
	ret := ctl.data[ctl.pos]
	ctl.data = ctl.data[:ctl.pos]
	return ret
}

func evalExp(e *ControlFlow, exp Obj, env Obj) Obj {
	e.TailEval(exp, env)
	return trampoline(e)
}

func evalIf(e *ControlFlow, a, b, c Obj, env Obj) {
	t := evalExp(e, a, env)
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

func evalAnd(e *ControlFlow, a, b Obj, env Obj) {
	if evalExp(e, a, env) == False {
		e.Return(False)
		return
	}
	e.TailEval(b, env)
}

func evalOr(e *ControlFlow, a, b Obj, env Obj) {
	if evalExp(e, a, env) == True {
		e.Return(True)
		return
	}
	e.TailEval(b, env)
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

func Call(e *ControlFlow, f Obj, args ...Obj) Obj {
	e.TailApply(f, args...)
	return trampoline(e)
}

func Eval(e *ControlFlow, exp Obj) (res Obj) {
	defer func() {
		if r := recover(); r != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			if x, ok := r.(Obj); ok && IsError(x) {
				fmt.Println("Panic:", mustError(x))
			} else {
				fmt.Println("Panic:", r)
			}
			fmt.Println("Recovered in Eval:", ObjString(exp))
			str := string(buf[:n])
			// fmt.Println(str)
			res = MakeError(str)
		}
	}()
	res = evalExp(e, exp, Nil)
	return
}

type tryResult struct {
	e    *ControlFlow
	data Obj
}

func Try(e *ControlFlow, f Obj) (res tryResult) {
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
	val := Call(e, f)
	res = tryResult{e: e, data: val}
	return
}

func (t tryResult) Catch(f Obj) Obj {
	if IsError(t.data) {
		return Call(t.e, f, t.data)
	}
	return t.data
}

func apply(ctl *ControlFlow) {
	f := ctl.data[ctl.pos]
	args := ctl.data[ctl.pos+1:]
	switch *f {
	case scmHeadProcedure:
		args := ctl.data[ctl.pos+1:]
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
			res := evalExp(ctl, proc.body, newEnv)
			ctl.TailApply(res, args[proc.arity:]...)
			return
		}
	case scmHeadNative:
		fn := MustNative(f)
		provided := len(fn.captured) + len(args)
		required := fn.require
		if provided != required {
			panic("partial apply unsupported")
		}
		if len(fn.captured) > 0 {
			// Captured data should come fisrt, then the arguments.
			ctl.data = append(ctl.data, make([]Obj, len(fn.captured))...)
			dst := ctl.data[len(ctl.data)-len(args):]
			copy(dst, args)
			dst = ctl.data[ctl.pos+1 : ctl.pos+1+len(fn.captured)]
			copy(dst, fn.captured)
		}
		fn.fn(ctl)
		return
	}
	panic("can't apply object")
}

func envGet(env Obj, sym Obj) (Obj, bool) {
	for env != Nil {
		pair := car(env)
		if car(pair) == sym {
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

func eval(e *ControlFlow) {
	exp := e.data[e.pos]
	env := e.data[e.pos+1]

	switch *exp {
	// handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull, scmHeadProcedure /* , scmHeadPrimitive */ :
		e.Return(exp)
		return
	case scmHeadSymbol:
		if val, ok := envGet(env, exp); ok {
			exp = val
		}
		e.Return(exp)
		return
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		exp = pair.cdr // handle special form
		switch pair.car {
		case symDefun: // (defun f (x y) z)
			proc := makeProcedure(cadr(exp), caddr(exp), env)
			funName := car(exp)
			BindSymbolFunc(funName, proc)
			e.Return(funName)
			return
		case symLambda: // (lambda x x)
			e.Return(makeProcedure(car(exp), cadr(exp), env))
			return
		case symFreeze: // (freeze body)
			e.Return(makeProcedure(Nil, car(exp), env))
			return
		case symLet: // (let x y z)
			args := evalExp(e, cadr(exp), env)
			newEnv := envExtend(env, []Obj{car(exp)}, []Obj{args})
			e.TailEval(caddr(exp), newEnv)
			return
		case symAnd:
			evalAnd(e, car(exp), cadr(exp), env)
			return
		case symOr:
			evalOr(e, car(exp), cadr(exp), env)
			return
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				evalIf(e, car(exp), cadr(exp), caddr(exp), env)
				return
			} // if may also be a function for partial apply
		case symCond: // (cond (false 1) (true 2))
			evalCond(e, exp, env)
			return
		case symType:
			// (type XXX (list Symbol)) => XXX, just ignore the second one
			e.TailEval(car(exp), env)
			return
		case symTrapError: // (trap-error ~body ~handler)
			evalTrapError(e, exp, env)
			return
		case symDo: // (do A A)
			evalExp(e, car(exp), env)
			e.TailEval(cadr(exp), env)
			return
		}
	}

	savePOS := e.pos
	fn := evalFunction(e, pair.car, env)
	e.data = e.data[:e.pos]
	e.data = append(e.data, fn)
	e.pos++

	args := pair.cdr
	for *args == scmHeadPair {
		v := evalExp(e, car(args), env)
		if *v == scmHeadError {
			e.pos = savePOS
			e.TailApply(fn, v)
			return
		}
		e.data = append(e.data, v)
		e.pos++
		args = cdr(args)
	}
	e.pos = savePOS
	e.kind = ControlFlowApply
}

func evalCond(e *ControlFlow, l Obj, env Obj) {
	for *l == scmHeadPair {
		curr := car(l)
		if evalExp(e, car(curr), env) == True {
			e.TailEval(cadr(curr), env)
			return
		}
		l = cdr(l)
	}
	e.Return(Nil)
}

func evalTrapError(e *ControlFlow, exp Obj, env Obj) {
	savePOS := e.pos
	defer func() {
		if err := recover(); err != nil {
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					e.pos = savePOS
					e.data = e.data[:e.pos]
					handle := evalFunction(e, cadr(exp), env)
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
			e.Return(MakeError("trap-error result is not Obj"))
			return
		}
	}()
	v := evalExp(e, car(exp), env)
	e.Return(v)
}

func evalFunction(e *ControlFlow, fn Obj, env Obj) Obj {
	if ok, sym := isSymbol(fn); ok {
		if sym.function != nil {
			return sym.function
		}

		if proc, ok := envGet(env, fn); ok {
			return proc
		}
	}

	switch *fn {
	case /* scmHeadPrimitive, */ scmHeadProcedure, scmHeadNative:
		return fn
	case scmHeadPair:
		return evalExp(e, fn, env)
	}
	panic(fmt.Sprintf("can't apply non function: %#v", (*scmHead)(fn)))
}

func (e *ControlFlow) Global(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.function != nil {
		return sym.function
	}
	errMsg := fmt.Sprintf("variable %s not bound", sym.str)
	panic(MakeError(errMsg))
}
