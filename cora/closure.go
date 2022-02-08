package cora

import (
	"fmt"
)

func Neval(exp Obj) Obj {
	c := compile(exp, nil, true)
	run(c)
	stack = stack[:sp]
	return val
}

type compileEnv struct {
	parent *compileEnv
	args   Obj
	// mark is the refered variable, i.e closure value
	mark map[int]struct{}
}

func (env *compileEnv) findVariable(s Obj) (m int, n int, e *compileEnv) {
	e = env
	for e != nil {
		n = 0
		for args := e.args; args != Nil; args = cdr(args) {
			if car(args) == s {
				return // local variable
			}
			n++
		}
		m = m + 1
		e = e.parent
	}
	if e == nil {
		m = -1 // global variable
	}
	return // closure variable
}

func compile(exp Obj, env *compileEnv, tail bool) func(env Env) {
	switch *exp {
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull:
		return genConstInst(exp)
	case scmHeadSymbol:
		m, n, e1 := env.findVariable(exp)
		if m == 0 {
			// local[0] is the closure object, the real args begins from i+1
			return genLocalRefInst(n+1)
		}
		if m > 0 {
			e1.mark[n] = struct{}{}
			return genClosureRefInst(m, n)
		}

		return genGlobalRefInst(exp)
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		tl := pair.cdr // handle special form
		switch pair.car {
		case symQuote:
			return genConstInst(car(tl))
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				x := compile(car(tl), env, false)
				y := compile(cadr(tl), env, tail)
				z := compile(caddr(tl), env, tail)
				return genIfInst(x, y, z)
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			x := compile(car(tl), env, false)
			y := compile(cadr(tl), env, tail)
			return genDoInst(x, y, tail)
		case symLambda: // (lambda (x y) x)
			args := car(tl)
			body := cadr(tl)
			env1 := &compileEnv{
				parent: env,
				args:   args,
				mark:   make(map[int]struct{}),
			}
			op := compile(body, env1, true)
			nargs := listLength(args)
			return genMakeClosureInst(op, nargs, env1.mark)
		}
	}

	var cached [5]func(Env)
	insts := cached[:0]
	debugStr := ObjString(exp)
	for *exp == scmHeadPair {
		inst := compile(car(exp), env, false)
		insts = append(insts, inst)
		exp = cdr(exp)
	}
	return genCallInst(insts, debugStr, tail)
}

type Env []Obj

func genConstInst(v Obj) func(env Env) {
	return func(env Env) {
		val = v
	}
}

func genIfInst(a, b, c func(Env)) func(Env) {
	return func(env Env) {
		a(env)
		switch val {
		case True:
			b(env)
		case False:
			c(env)
		default:
			panic("test branch for if must be a boolean")
		}
	}
}

func genDoInst(op1, op2 func(env Env), tail bool) func(env Env) {
	return func(env Env) {
		op1(env)
		if tail {
			globalEnv = env
			pc = op2
			return
		}
		op2(env)
	}
}

func genLocalRefInst(idx int) func(Env) {
	return func(env Env) {
		val = env[idx]
	}
}

func genClosureRefInst(m, n int) func(Env) {
	return func(env Env) {
		clo := mustClosure(env[0])
		for i:=0; i<m; i++{
			clo = clo.parent
		}
		val = clo.freeVars[n]
	}
}

func genGlobalRefInst(sym Obj) func(Env) {
	return func(env Env) {
		s := mustSymbol(sym)
		if s.cora == nil {
			fmt.Println("NOT DEFINED")
			panic("undefined symbol:" + s.str)
		}
		val = s.cora
	}
}

func genMakeClosureInst(op func(Env), nargs int, mark map[int]struct{}) func(Env) {
	return func(env Env) {
		var parent *scmClosure
		if len(env) > 0 {
			if *env[0] == scmHeadClosure {
				parent = mustClosure(env[0])
			}
		}
		// Save some variables into the closure
		// Because they are free variables.
		freeVars := make(map[int]Obj)
		for k, _ := range mark {
			freeVars[k] = env[k+1]
			fmt.Println("... free var ==", k, ObjString(freeVars[k]))
		}
		val = MakeClosure(op, env, nargs, parent, freeVars)
	}
}

func getRequired(o Obj) int {
	switch *o {
	case scmHeadClosure:
		return mustClosure(o).params
	case scmHeadCurry:
		return mustCurry(o).params
	case scmHeadNative:
		n := MustNative(o)
		return n.require - len(n.captured)
	}
	panic("not implemented" + ObjString(o))
}

func prepareCall(env Env, insts []func(Env)) int {
	// Save SP
	saveSP := sp
	sp = len(stack)
	for _, inst := range insts {
		inst(env)
		stack = append(stack, val)
	}
	return saveSP
}

func recoverCall(saveSP int) {
	stack = stack[:sp]
	sp = saveSP
}

func NCall(f Obj, args ...Obj) Obj {
	saveSP := sp
	sp = len(stack)
	stack = append(stack, f)
	for _, arg := range args {
		stack = append(stack, arg)
	}

	myCall(nil)
	recoverCall(saveSP)
	return val
}

func genCallInst(insts []func(env Env), debugStr string, tail bool) func(Env) {
	return func(env Env) {
		// printTabs()
		saveSP := prepareCall(env, insts)
		fmt.Println("call:", debugStr, "sp==", sp, "len(stack)==", len(stack))
		printStack()

	again:

		f := stack[sp]
		required := getRequired(f)
		provided := len(stack[sp+1:])

		if required > provided {
			upvalues := make([]Obj, provided)
			copy(upvalues, stack[sp+1:])
			val = MakeCurry(f, upvalues, required-provided)
			recoverCall(saveSP)
			return
		}

		if required < provided {
			// Prepare a new stack for calling...
			saveSP1 := sp
			sp = len(stack)
			for i := saveSP1; i < saveSP1+1+required; i++ {
				stack = append(stack, stack[i])
			}
			// Make the call
			myCall(env)
			// Recover the stack
			stack[saveSP1] = val
			to := saveSP1 + 1
			from := saveSP1 + 1 + required
			for from < saveSP1+1+provided {
				stack[to] = stack[from]
				from++
				to++
			}
			stack = stack[:to]
			sp = saveSP1

			goto again
		}

		if tail {
			f := stack[sp]
			if *f == scmHeadClosure {
				clo := mustClosure(f)
				nargs := len(stack[sp:])
				copy(stack[saveSP:saveSP+nargs], stack[sp:])
				stack = stack[:saveSP+nargs]
				sp = saveSP
				env = stack[sp:]

				// Some of the data is used by the body (marked),
				// They need persistent, should be heap referenced!
				// if len(clo.mark) > 0 {
				// 	for k, addr := range clo.mark {
				// 		*addr = env[k]
				// 	}
				// }

				pc = clo.code
				globalEnv = env
				return
			}
		}

		myCall(env)
		recoverCall(saveSP)
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

var globalEnv Env
var pc func(Env)

func myCall(env Env) {
retry:
	f := stack[sp]
	switch *f {

	case scmHeadNative:
		fn := MustNative(f)
		var ctl ControlFlow
		ctl.kind = ControlFlowApply
		ctl.data = stack
		ctl.pos = sp
		fn.fn(&ctl)
		val = ctl.Get(0)

	case scmHeadClosure:
		clo := mustClosure(f)
		env = stack[sp:]

		// Some of the data is used by the body (marked),
		// They need persistent, should be heap referenced!
		// if len(clo.mark) > 0 {
		// 	for k, v := range clo.mark {
		// 		v[0] = env[k]
		// 	}
		// }

		globalEnv = env
		run(clo.code)

	case scmHeadCurry:
		curry := mustCurry(f)
		saved := append([]Obj{}, stack[sp+1:]...)
		stack = stack[:sp]
		stack = append(stack, curry.origin)
		stack = append(stack, curry.upvalue...)
		stack = append(stack, saved...)
		f = stack[sp]
		goto retry
	default:
		panic("invalid call operation on:" + ObjString(f))
	}
}

func printStack() {
	for i := 0; i < len(stack); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		if i == sp {
			fmt.Print(" | ")
		}
		fmt.Printf("%s", ObjString(stack[i]))
	}
	fmt.Println()
}

var val Obj
var sp int
var stack []Obj
