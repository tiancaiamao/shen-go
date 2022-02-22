package cora

import (
	"fmt"
)

func EvalKL(exp Obj) Obj {
	c := compileKL(exp, nil, true, nil)
	run(c)
	stack = stack[:sp]
	return val
}

var klMacro map[Obj]func(obj Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env)

func compileKL(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	switch *exp {
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull:
		return genConstInst(exp)
	case scmHeadSymbol:
		m, n := env.findVariable(exp)
		if m == 0 {
			// local[0] is the closure object, the real args begins from i+1
			return genLocalRefInst(n + 1)
		}
		if m > 0 {
			freeVars[exp] = posRef{up: m, offset: n}
			return genClosureRefInst(m, n)
		}

		return genConstInst(exp)
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		tl := pair.cdr // handle special form
		switch pair.car {
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				x := compileKL(car(tl), env, false, freeVars)
				y := compileKL(cadr(tl), env, tail, freeVars)
				z := compileKL(caddr(tl), env, tail, freeVars)
				return genIfInst(x, y, z)
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			x := compileKL(car(tl), env, false, freeVars)
			y := compileKL(cadr(tl), env, tail, freeVars)
			return genDoInst(x, y, tail)
		case symLambda: // (lambda x body)
			args := cons(car(tl), Nil)
			body := cadr(tl)
			return compileLambda(args, body, env, freeVars)
		}
	}

	var cached [5]func(Env)
	insts := cached[:0]
	if IsSymbol(pair.car) {
		if fn, ok := klMacro[pair.car]; ok {
			return fn(pair.cdr, env, tail, freeVars)
		}
		inst := compileSymbolInCall(pair.car, env, false, freeVars)
		insts = append(insts, inst)
		exp = pair.cdr
	}

	for *exp == scmHeadPair {
		inst := compileKL(car(exp), env, false, freeVars)
		insts = append(insts, inst)
		exp = cdr(exp)
	}
	return genCallInst(insts, exp, tail)
}

func compileLambda(args, body Obj, env *compileEnv, freeVars map[Obj]posRef) func(env Env) {
	env1 := &compileEnv{
		parent: env,
		args:   args,
	}

	collectFVs := make(map[Obj]posRef)
	op := compileKL(body, env1, true, collectFVs)
	nargs := listLength(args)

	for v, p := range collectFVs {
		if p.up != 1 {
			// Collect to freeVars and delete from this one.
			freeVars[v] = posRef{up: p.up - 1, offset: p.offset}
			delete(collectFVs, v)
		}
	}
	return genMakeClosureInst(op, nargs, collectFVs)
}

func compileOrMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	// (or x y) => (if x true (if y true false))
	x := car(exp)
	y := cadr(exp)
	tmp := cons(symIf, cons(y, cons(True, cons(False, Nil))))
	exp = cons(symIf, cons(x, cons(True, cons(tmp, Nil))))
	return compileKL(exp, env, tail, freeVars)
}

func compileAndMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	// (and x y) => (if x (if y true false) false)
	x := car(exp)
	y := cadr(exp)
	tmp := cons(symIf, cons(y, cons(True, cons(False, Nil))))
	exp = cons(symIf, cons(x, cons(tmp, cons(False, Nil))))
	return compileKL(exp, env, tail, freeVars)
}

func compileCondMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	exp = compileCondMacroHelp(exp)
	return compileKL(exp, env, tail, freeVars)
}

func compileCondMacroHelp(exp Obj) Obj {
	// (cond (x y) ...)
	if exp == Nil {
		return cons(MakeSymbol("simple-error"), cons(MakeString("no cond match"), Nil))
	}
	xy := car(exp)
	x := car(xy)
	y := cadr(xy)
	tmp := compileCondMacroHelp(cdr(exp))
	return cons(symIf, cons(x, cons(y, cons(tmp, Nil))))
}

func compileDefunMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	// (defun f (x y) ..) => ((fn 'set) (quote f) (lambda (x y) ...))
	fName := car(exp)
	args := cadr(exp)
	body := caddr(exp)
	inst := compileLambda(args, body, env, freeVars)
	return genKLGlobalSetInst(fName, inst)
}

func compileSymbolInCall(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	m, n := env.findVariable(exp)
	if m == 0 {
		// local[0] is the closure object, the real args begins from i+1
		return genLocalRefInst(n + 1)
	}
	if m > 0 {
		freeVars[exp] = posRef{up: m, offset: n}
		return genClosureRefInst(m, n)
	}
	return genKLGlobalRefInst(exp)
}

func genKLGlobalRefInst(sym Obj) func(Env) {
	return func(env Env) {
		s := mustSymbol(sym)
		if s.function == nil {
			fmt.Println("NOT DEFINED")
			panic("undefined symbol:" + s.str)
		}
		val = s.function
	}
}

func genKLGlobalSetInst(fName Obj, inst func(Env)) func(Env) {
	return func(env Env) {
		inst(env)
		val = PrimNS2Set(fName, val)
	}
}

func compileLetMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	// (let x y z) => ((lambda x z) y)
	x := car(exp)
	y := cadr(exp)
	z := cdr(cdr(exp))
	fn := cons(symLambda, cons(x, z))
	exp = cons(fn, cons(y, Nil))
	return compileKL(exp, env, tail, freeVars)
}

func compileFreezeMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	// (freeze body) => (lambda () body)
	return compileLambda(Nil, car(exp), env, freeVars)
}

func compileTrapErrorMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	// (trap-error body handler) => (try-catch (lambda () body) handler)
	body := car(exp)
	handler := cdr(exp)
	action := cons(symLambda, cons(Nil, cons(body, Nil)))
	exp = cons(MakeSymbol("try-catch"), cons(action, handler))
	return compileKL(exp, env, tail, freeVars)
}

func compileTypeMacro(exp Obj, env *compileEnv, tail bool, freeVars map[Obj]posRef) func(env Env) {
	// (type exp _) => exp
	return compileKL(car(exp), env, tail, freeVars)
}
