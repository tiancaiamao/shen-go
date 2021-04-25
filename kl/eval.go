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
			fmt.Println("Panic:", r)
			fmt.Println("Recovered in Eval:", ObjString(exp))
			str := string(buf[:n])
			// fmt.Println(str)
			res = MakeError(str)
		}
	}()
	res = evalExp(e, exp, Nil)
	return
}

func apply(ctl *ControlFlow) {
	f := ctl.data[ctl.pos]
	args := ctl.data[ctl.pos+1:]
	if isClosure(f) {
		// (lambda (params) body . env)
		params := cadr(f)
		body := caddr(f)
		env := cdddr(f)

		env, params, args = envExtend(env, params, args)
		if params != Nil {
			// More params than args, auto curry.
			closure := makeClosure(params, body, env)
			ctl.Return(closure)
			return
		}
		if len(args) > 0 {
			savePos := ctl.pos
			ctl.pos = len(ctl.data)
			f := evalExp(ctl, body, env)
			ctl.pos = savePos
			ctl.TailApply(f, args...)
			return
		}
		ctl.TailEval(body, env)
		return
	}

	if *f == scmHeadNative {
		fn := MustNative(f)
		provided := len(fn.captured) + len(args)
		required := fn.require
		if provided == required {
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

		if provided < required {
			tmp := make([]Obj, 0, fn.require)
			tmp = append(tmp, fn.captured...)
			tmp = append(tmp, args...)
			ctl.Return(MakeNative(fn.fn, fn.require, tmp...))
			return
		}

		// Trick: instead of adjust the arguments, we use a new call stack here.
		savePos := ctl.pos
		ctl.pos = len(ctl.data)
		ctl.data = append(ctl.data, f) // Actually, it's not f any more.
		ctl.data = append(ctl.data, fn.captured...)
		taken := required - len(fn.captured)
		ctl.data = append(ctl.data, ctl.data[savePos+1:savePos+1+taken]...)
		fn.fn(ctl)
		f := ctl.data[ctl.pos]
		ctl.pos = savePos
		args = args[taken:]
		ctl.TailApply(f, args...)
		return
	}
	panic("can't apply object")
}

var symDebug = MakeSymbol("*debug-eval*")

func eval(e *ControlFlow) {
	exp := e.data[e.pos]
	env := e.data[e.pos+1]

	enableDebug := mustSymbol(symDebug).cora
	if enableDebug != nil && enableDebug != Nil {
		fmt.Println("eval === exp:", ObjString(exp))
	}

	switch *exp {
	// handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull /*, scmHeadProcedure , scmHeadPrimitive */ :
		e.Return(exp)
		return
	case scmHeadSymbol:
		if found := envGet(env, exp); found != Nil {
			e.Return(cdr(found))
			return
		}
		sym := mustSymbol(exp)
		if sym.cora == nil {
			// TODO: xxx
			fmt.Println("NOT DEFINED")
			panic("undefined symbol:" + sym.str)
		}
		e.Return(sym.cora)
		return
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		tl := pair.cdr // handle special form
		switch pair.car {
		case symQuote:
			e.Return(car(tl))
			return
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				evalIf(e, car(tl), cadr(tl), caddr(tl), env)
				return
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			evalExp(e, car(tl), env)
			e.TailEval(cadr(tl), env)
			return
		case symLambda: // (lambda (x y) x)
			e.Return(makeClosure(car(tl), cadr(tl), env))
			return
		}
	}

	savePOS := e.pos
	args := exp
	for *args == scmHeadPair {
		v := evalExp(e, car(args), env)
		e.data = append(e.data, v)
		e.pos++
		args = cdr(args)
	}
	e.pos = savePOS
	e.kind = ControlFlowApply
	return
}

func envGet(env, sym Obj) Obj {
	for ; env != Nil; env = cdr(env) {
		pair := car(env)
		if car(pair) == sym {
			return pair
		}
	}
	return Nil
}

// envExtend return the new (env, params, args)
func envExtend(env Obj, params Obj, args []Obj) (Obj, Obj, []Obj) {
	for params != Nil && len(args) > 0 {
		pair := cons(car(params), args[0])
		env = cons(pair, env)
		params = cdr(params)
		args = args[1:]
	}
	return env, params, args
}
