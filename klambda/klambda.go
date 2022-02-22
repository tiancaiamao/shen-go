package klambda

import (
	. "github.com/tiancaiamao/cora/cora_go"
)

func EvalKL(exp Obj) Obj {
	c := compileKL(exp, nil, true, nil)
	return Run(c)
}

var symType = MakeSymbol("type")
var symDefun = MakeSymbol("defun")
var symFreeze = MakeSymbol("freeze")
var symLet = MakeSymbol("let")
var symCond = MakeSymbol("cond")
var symAnd = MakeSymbol("and")
var symOr = MakeSymbol("or")
var symIf = MakeSymbol("if")
var symLambda = MakeSymbol("lambda")
var symTrapError = MakeSymbol("trap-error")
var symDo = MakeSymbol("do")


var globalEnv Env
var pc func(Env)
var ctl ControlFlow

var klMacro map[Obj]func(obj Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env)

func init() {
	klMacro = map[Obj]func(obj Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env){
		symType:      compileTypeMacro,
		symDefun:     compileDefunMacro,
		symFreeze:    compileFreezeMacro,
		symLet:       compileLetMacro,
		symCond:      compileCondMacro,
		symAnd:       compileAndMacro,
		symOr:        compileOrMacro,
		symTrapError: compileTrapErrorMacro,
	}
}

func run(code func(Env)) {
	pc = code
	for pc != nil {
		tmp := pc
		pc = nil
		tmp(globalEnv)
	}
}

func isBoolean(exp Obj) bool {
	return exp == True || exp == False
}

func isNil(exp Obj) bool {
	return exp == Nil
}

func compileKL(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	switch {
	case IsNumber(exp), IsString(exp), IsVector(exp), isBoolean(exp), isNil(exp):
		return GenConstInst(exp)
	case IsSymbol(exp):
		m, n := env.FindVariable(exp)
		if m == 0 {
			// local[0] is the closure object, the real args begins from i+1
			return GenLocalRefInst(n + 1)
		}
		if m > 0 {
			freeVars[exp] = PosRef{Up: m, Offset: n}
			return GenClosureRefInst(m, n)
		}

		return GenConstInst(exp)
	}

	if IsSymbol(Car(exp)) {
		tl := Cdr(exp) // handle special form
		switch Car(exp) {
		case symIf: // (if a b c)
			if ListLength(Cdr(exp)) == 3 {
				x := compileKL(Car(tl), env, false, freeVars)
				y := compileKL(Cadr(tl), env, tail, freeVars)
				z := compileKL(caddr(tl), env, tail, freeVars)
				return GenIfInst(x, y, z)
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			x := compileKL(Car(tl), env, false, freeVars)
			y := compileKL(Cadr(tl), env, tail, freeVars)
			return GenDoInst(x, y, tail)
		case symLambda: // (lambda x body)
			args := Cons(Car(tl), Nil)
			body := Cadr(tl)
			return CompileLambda(args, body, env, freeVars, compileKL)
		}
	}

	var cached [5]func(Env)
	insts := cached[:0]
	if IsSymbol(Car(exp)) {
		if fn, ok := klMacro[Car(exp)]; ok {
			return fn(Cdr(exp), env, tail, freeVars)
		}
		inst := compileSymbolInCall(Car(exp), env, false, freeVars)
		insts = append(insts, inst)
		exp = Cdr(exp)
	}

	for IsCons(exp) {
		inst := compileKL(Car(exp), env, false, freeVars)
		insts = append(insts, inst)
		exp = Cdr(exp)
	}
	return GenCallInst(insts, exp, tail)
}

func compileOrMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	// (or x y) => (if x true (if y true false))
	x := Car(exp)
	y := Cadr(exp)
	tmp := Cons(symIf, Cons(y, Cons(True, Cons(False, Nil))))
	exp = Cons(symIf, Cons(x, Cons(True, Cons(tmp, Nil))))
	return compileKL(exp, env, tail, freeVars)
}

func compileAndMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	// (and x y) => (if x (if y true false) false)
	x := Car(exp)
	y := Cadr(exp)
	tmp := Cons(symIf, Cons(y, Cons(True, Cons(False, Nil))))
	exp = Cons(symIf, Cons(x, Cons(tmp, Cons(False, Nil))))
	return compileKL(exp, env, tail, freeVars)
}

func compileCondMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	exp = compileCondMacroHelp(exp)
	return compileKL(exp, env, tail, freeVars)
}

func compileCondMacroHelp(exp Obj) Obj {
	// (cond (x y) ...)
	if exp == Nil {
		return Cons(MakeSymbol("simple-error"), Cons(MakeString("no cond match"), Nil))
	}
	xy := Car(exp)
	x := Car(xy)
	y := Cadr(xy)
	tmp := compileCondMacroHelp(Cdr(exp))
	return Cons(symIf, Cons(x, Cons(y, Cons(tmp, Nil))))
}

func compileDefunMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	// (defun f (x y) ..) => ((fn 'set) (quote f) (lambda (x y) ...))
	fName := Car(exp)
	args := Cadr(exp)
	body := caddr(exp)
	inst := CompileLambda(args, body, env, freeVars, compileKL)
	return genKLGlobalSetInst(fName, inst)
}

func compileSymbolInCall(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	m, n := env.FindVariable(exp)
	if m == 0 {
		// local[0] is the closure object, the real args begins from i+1
		return GenLocalRefInst(n + 1)
	}
	if m > 0 {
		freeVars[exp] = PosRef{Up: m, Offset: n}
		return GenClosureRefInst(m, n)
	}
	return genKLGlobalRefInst(exp)
}

func genKLGlobalRefInst(sym Obj) func(Env) {
	return func(env Env) {
		// s := mustSymbol(sym)
		// if s.function == nil {
		// 	fmt.Println("NOT DEFINED")
		// 	panic("undefined symbol:" + s.str)
		// }
		// val = s.function
		val := PrimNS2Value(sym)
		SetVal(val)
	}
}

func genKLGlobalSetInst(fName Obj, inst func(Env)) func(Env) {
	return func(env Env) {
		inst(env)
		val := PrimNS2Set(fName, GetVal())
		SetVal(val)
	}
}

func compileLetMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	// (let x y z) => ((lambda x z) y)
	x := Car(exp)
	y := Cadr(exp)
	z := Cdr(Cdr(exp))
	fn := Cons(symLambda, Cons(x, z))
	exp = Cons(fn, Cons(y, Nil))
	return compileKL(exp, env, tail, freeVars)
}

func compileFreezeMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	// (freeze body) => (lambda () body)
	return CompileLambda(Nil, Car(exp), env, freeVars, compileKL)
}

func compileTrapErrorMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	// (trap-error body handler) => (try-catch (lambda () body) handler)
	body := Car(exp)
	handler := Cdr(exp)
	action := Cons(symLambda, Cons(Nil, Cons(body, Nil)))
	exp = Cons(MakeSymbol("try-catch"), Cons(action, handler))
	return compileKL(exp, env, tail, freeVars)
}

func compileTypeMacro(exp Obj, env *CompileEnv, tail bool, freeVars map[Obj]PosRef) func(env Env) {
	// (type exp _) => exp
	return compileKL(Car(exp), env, tail, freeVars)
}

func caddr(exp Obj) Obj {
	return Car(Cdr(Cdr(exp)))
}
