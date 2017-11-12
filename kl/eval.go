package kl

import (
	"fmt"
)

type Environment struct {
	parent *Environment
	bind   map[string]Obj
}

func (env *Environment) Get(sym string) (Obj, bool) {
	for env != nil {
		if v, ok := env.bind[sym]; ok {
			return v, true
		}
		env = env.parent
	}
	return Nil, false
}

func (env *Environment) Extend(symbols, values Obj) *Environment {
	if *symbols == Symbol {
		symbols = cons(symbols, Nil)
	}

	bind := make(map[string]Obj)
	for symbols != Nil && values != Nil {
		name := mustSymbol(car(symbols)).sym
		bind[name] = car(values)

		symbols = cdr(symbols)
		values = cdr(values)
	}
	return &Environment{
		parent: env,
		bind:   bind,
	}
}

func Eval(exp Obj) Obj {
	return trampoline(exp, nil)
}

type controlFlowKind int

const (
	controlFlowReturn controlFlowKind = iota
	controlFlowEval
	controlFlowApply
)

type controlFlow struct {
	kind controlFlowKind
	// arguments for eval
	exp Obj
	env *Environment
	// arguments for apply
	f    Obj
	args Obj
	// return result
	result Obj
}

// trampoline is introduced for tail call optimization.
func trampoline(exp Obj, env *Environment) Obj {
	var ctl = controlFlow{
		kind: controlFlowEval,
		exp:  exp,
		env:  env,
	}
	for ctl.kind != controlFlowReturn {
		switch ctl.kind {
		case controlFlowEval:
			eval(&ctl)
		case controlFlowApply:
			apply(&ctl)
		}
	}
	return ctl.result
}

func (ctl *controlFlow) TailEval(exp Obj, env *Environment) {
	ctl.exp = exp
	ctl.env = env
	ctl.kind = controlFlowEval
}

func (ctl *controlFlow) TailApply(f, args Obj) {
	ctl.f = f
	ctl.args = args
	ctl.kind = controlFlowApply
}

func (ctl *controlFlow) Return(result Obj) {
	ctl.result = result
	ctl.kind = controlFlowReturn
}

// Exception can be replaced by Return totally, just defined as a name alias.
func (ctl *controlFlow) Exception(err Obj) {
	mustError(err)
	ctl.result = err
	ctl.kind = controlFlowReturn
}

func eval(ctl *controlFlow) {
	exp := ctl.exp
	env := ctl.env

	switch *exp {
	case Number, String, Vector, Boolean, Procedure, Null:
		ctl.Return(exp)
		return
	}

	if *exp == Symbol {
		sym := mustSymbol(exp)
		if val, ok := env.Get(sym.sym); ok {
			ctl.Return(val)
			return
		}
		ctl.Return(exp)
		return
	}

	pair := mustPair(exp)
	if *pair.car == Symbol {
		sym := mustSymbol(pair.car)
		exp = pair.cdr
		switch sym.sym {
		case "defun": // (defun f (x y) z)
			proc := Make_procedure(cadr(exp), caddr(exp), env)
			funName := car(exp)
			functionTable[mustSymbol(funName).sym] = proc
			ctl.Return(funName)
			return
		case "lambda": // (lambda x x)
			ctl.Return(Make_procedure(car(exp), cadr(exp), env))
			return
		case "freeze": // (freeze body)
			ctl.Return(Make_procedure(Nil, car(exp), env))
			return
		case "let": // (let x y z)
			args := trampoline(cadr(exp), env)
			if *args == Error {
				ctl.Exception(args)
			}
			newEnv := env.Extend(car(exp), cons(args, Nil))
			ctl.TailEval(caddr(exp), newEnv)
			return
		case "and":
			if trampoline(car(exp), env) == False {
				ctl.Return(False)
				return
			}
			ctl.TailEval(cadr(exp), env)
			return
		case "or":
			if trampoline(car(exp), env) == True {
				ctl.Return(True)
				return
			}
			ctl.TailEval(cadr(exp), env)
			return
		case "if": // (if a b c)
			if listLength(pair.cdr) == 3 {
				evalIf(car(exp), cadr(exp), caddr(exp), env, ctl)
				return
			} // if may also be a function for partial apply
		case "cond": // (cond (false 1) (true 2))
			evalCond(exp, env, ctl)
			return
		case "trap-error": // (trap-error ~body ~handler)
			evalTrapError(exp, env, ctl)
			return
		case "do": // (do A A)
			if tmp := trampoline(car(exp), env); *tmp == Error {
				ctl.Exception(tmp)
				return
			}
			ctl.TailEval(cadr(exp), env)
			return
		}
	}

	fn := functionCall(pair.car, env)
	if *fn == Error {
		ctl.Exception(fn)
		return
	}
	args := evalList(pair.cdr, env)
	ctl.TailApply(fn, args)
	return
}

func functionCall(fn Obj, env *Environment) Obj {
	switch *fn {
	case Symbol:
		sym := mustSymbol(fn)
		if proc, ok := env.Get(sym.sym); ok {
			return proc
		}

		if val, ok := functionTable[sym.sym]; ok {
			return val
		}
	case Primitive, Procedure:
		return fn
	case Pair:
		return trampoline(fn, env)
	}
	return Make_error(fmt.Sprintf("can't apply non function: %#v", (*scmHead)(fn)))
}

func evalIf(a, b, c Obj, env *Environment, ctl *controlFlow) {
	t := trampoline(a, env)
	switch t {
	case True:
		ctl.TailEval(b, env)
		return
	case False:
		ctl.TailEval(c, env)
		return
	}
	ctl.Exception(Make_error("second argument of if should be boolean"))
}

func evalCond(l Obj, env *Environment, ctl *controlFlow) {
	for l != Nil {
		curr := car(l)
		if trampoline(car(curr), env) == True {
			ctl.TailEval(cadr(curr), env)
			return
		}
		l = cdr(l)
	}
	ctl.Return(Nil)
	return
}

func evalTrapError(exp Obj, env *Environment, ctl *controlFlow) {
	e := trampoline(car(exp), env)
	if *e == Error {
		handler := functionCall(cadr(exp), env)
		ctl.TailApply(handler, cons(e, Nil))
		return
	}
	ctl.Return(e)
	return
}

func apply(ctl *controlFlow) {
	f := ctl.f
	args := ctl.args
	if *f == Primitive {
		prim := mustPrimitive(f)
		args1 := listToSlice(args)
		switch {
		case len(args1) < prim.required:
			ctl.Return(partialApply(prim.required, args1, nil, f))
		case len(args1) == prim.required:
			ctl.Return(prim.function(args1...))
		default:
			msg := fmt.Sprintf("too many arguments for %s, required: %d, given: %d", prim.name, prim.required, len(args1))
			ctl.Exception(Make_error(msg))
		}
		return
	}

	if *f == Procedure {
		proc := mustProcedure(f)
		if listLength(args) < proc.arity {
			ctl.Return(partialApply(proc.arity, listToSlice(args), proc.env, f))
			return
		}
		newEnv := proc.env.Extend(proc.arg, args)
		ctl.TailEval(proc.body, newEnv)
		return
	}
	panic("not implement")
}

// partialApply works when required > providArgs
func partialApply(required int, providArgs []Obj, env *Environment, proc Obj) Obj {
	// Partial apply...
	// (f x y z) => (lambda (z) (f x y z)) with x y in env
	symbols := makeTempSymbols(required)
	bind := make(map[string]Obj)
	for i := 0; i < len(providArgs); i++ {
		symbol := symbols[i]
		bind[symbol.sym] = providArgs[i]
	}
	env1 := &Environment{
		parent: env,
		bind:   bind,
	}

	args := Nil
	for i, count := len(symbols)-1, required-len(providArgs); count > 0; count-- {
		args = cons(&symbols[i].scmHead, args)
		i--
	}

	body := Nil
	for i := len(symbols) - 1; i >= 0; i-- {
		body = cons(&symbols[i].scmHead, body)
	}
	body = cons(proc, body)

	return Make_procedure(args, body, env1)
}

func evalList(args Obj, env *Environment) Obj {
	res := Nil
	for args != Nil {
		v := trampoline(car(args), env)
		res = cons(v, res)
		args = cdr(args)
	}
	return reverse(res)
}

func makeTempSymbols(n int) []*ScmSymbol {
	ret := make([]*ScmSymbol, n)
	for i := 0; i < n; i++ {
		ret[i] = mustSymbol(Make_symbol(fmt.Sprintf("tmp%d", i)))
	}
	return ret
}

func listToSlice(l Obj) []Obj {
	var ret []Obj
	for l != Nil {
		ret = append(ret, car(l))
		l = cdr(l)
	}
	return ret
}
