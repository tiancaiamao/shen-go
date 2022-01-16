package cora

import (
	"fmt"
)

func ceval(exp Obj) Obj {
	c := compile(exp, nil, true)
	exec(c)
	return val
}

type compileEnv struct {
	parent *compileEnv
	args Obj
	// mark is the refered variable, i.e closure value
	mark map[int]struct{}
}

func (env *compileEnv) findVariable(s Obj) (m int, n int, e *compileEnv) {
	e = env
	for e != nil {
		n  = 0
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

func compile(exp Obj, env *compileEnv, tail bool) func() {
	switch *exp {
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull:
		return genConstInst(exp)
	case scmHeadSymbol:
		m, n, e1 := env.findVariable(exp)
		if m == 0 {
			// 需要把闭包变量收集起来
			return genLocalRefInst(n)
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
				x := compile(car(tl), env, false);
				y := compile(cadr(tl), env, tail);
				z := compile(caddr(tl), env, tail);
				return genIfInst(x, y, z);
			} // if may also be a function for partial apply
		case symDo: // (do A A)
			x := compile(car(tl), env, false)
			y := compile(cadr(tl), env, tail)
			return genDoInst(x, y)
		case symLambda: // (lambda (x y) x)
			args := car(tl)
			body := cadr(tl)
			// compile body 的时候，把 body 里面的闭包变量全部收集起来
			// 然后用于构造 closure 的闭包变量
			env1 := &compileEnv{
				parent: env,
				args: args,
				mark: make(map[int]struct{}),
			}
			op := compile(body, env1, true)
			nargs := listLength(args)
			return genClosureInst(op, nargs, env1.mark)
		}
	}

	var cached [5]func()
	insts := cached[:0]
	debugStr := ObjString(exp)
	for ;*exp == scmHeadPair; {
		// fmt.Println("compiling...", ObjString(car(exp)))
		inst := compile(car(exp), env, false)
		insts = append(insts, inst)
		exp = cdr(exp)
	}
	return genCallInst(insts, debugStr, tail)
}

type Env struct {
	parent *Env
	data []Obj
	heap map[int]Obj
}

func genConstInst(v Obj) func() {
	return func() {
		val = v
	}
}

func genIfInst(a, b, c func()) func() {
	return func() {
		exec(a)
		switch val {
		case True:
			pc = b
		case False:
			pc = c
		default:
			panic("test branch for if must be a boolean")
		}
	}
}

func genDoInst(op1, op2 func()) func() {
	return func() {
		op1()
		pc = op2
	}
}

func genLocalRefInst(idx int) func() {
	return func() {
		val = env.data[idx]
	}
}

func genClosureRefInst(m, n int) func() {
	return func() {
		e := env
		for i:=0; i<m; i++ {
			e = e.parent
		}
		val = e.heap[n]
	}
}

func genGlobalRefInst(sym Obj) func() {
	return func() {
		// fmt.Println("run here global ref")
		s := mustSymbol(sym)
		val = s.cora
	}
}

func genClosureInst(op func(), nargs int, mark map[int]struct{}) func() {
	return func() {
		val = MakeClosure(op, env, nargs, mark)
	}
}

func getRequired(o Obj) int {
	switch (*o) {
	case scmHeadClosure:
		return mustClosure(o).params
	case scmHeadCurry:
		return mustCurry(o).params
	case scmHeadPrimitive:
		return mustPrimitive(o).params
	}
	panic("not implemented"+ObjString(o))
}

func genCallInst(insts []func(), debugStr string, tail bool) func() {
	return func() {
		// printTabs()
		// fmt.Println("call:", debugStr, "sp==", sp, "len(stack)==", len(stack))
		// printStack()

		saveSP := sp
		sp = len(stack)
		for _, inst := range insts {
			exec(inst)
			stack = append(stack, val)
		}

		f := stack[sp]
		required := getRequired(f)
		provided := len(stack[sp+1:])

		if required > provided {
			upvalues := make([]Obj, provided)
			copy(upvalues, stack[sp+1:])
			val = MakeCurry(f, upvalues, required - provided)
		} else if required < provided {
			panic("required < provided?")
		} else {

retry:
			switch *f {
			case scmHeadPrimitive:
				priv := mustPrimitive(f)
				val = priv.fn(stack[sp+1:]...)

			case scmHeadClosure:
				clo := mustClosure(f)
				saveEnv := env
				env = &Env{
					parent: clo.env,
					data: stack[sp+1:],
				}

				// Some of the data is used by the body (marked),
				// They need persistent, should be heap referenced! 
				if len(clo.mark) > 0 {
					env.heap = make(map[int]Obj, len(clo.mark))
					for idx := range clo.mark {
						env.heap[idx] = env.data[idx]
					}
				}

				if tail {
					pc = func() {
						exec(clo.code)
						// clo.code()
						env = saveEnv
						sp = saveSP
						stack = stack[:saveSP]
						return
					}
					return
				} else {
					exec(clo.code)
					env = saveEnv
				}

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

		if tail {
			stack = stack[:saveSP]
		} else {
			stack = stack[:sp]
		}
		sp = saveSP
	}
}

func printStack() {
	for i:=0; i<len(stack); i++ {
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

func exec(f func()) {
	pc = f
	for pc != nil {
		tmp := pc
		pc = nil
		tmp()
	}
}


var env *Env
var val Obj
var sp int
var stack []Obj
var pc func()
