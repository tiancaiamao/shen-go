package kl

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"
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

func (env *Environment) Extend(symbols, values []Obj) *Environment {
	if len(symbols) == 0 {
		return env
	}

	bind := make(map[string]Obj)
	for i := 0; i < len(symbols); i++ {
		name := mustSymbol(symbols[i]).sym
		bind[name] = values[i]
	}
	return &Environment{
		parent: env,
		bind:   bind,
	}
}

type Evaluator struct {
	upTime        time.Time
	symbolTable   map[string]Obj
	functionTable map[string]Obj
}

func NewEvaluator() *Evaluator {
	var e Evaluator
	e.upTime = time.Now()
	e.symbolTable = make(map[string]Obj)

	e.functionTable = make(map[string]Obj, len(Primitives))
	for _, prim := range Primitives {
		e.functionTable[prim.Name] = Obj(&prim.scmHead)
	}
	// Overload for primitive set and value.
	tmp := &ScmPrimitive{scmHead: Primitive, Name: "set", Required: 2, Function: e.primSet}
	e.functionTable["set"] = Obj(&tmp.scmHead)
	tmp = &ScmPrimitive{scmHead: Primitive, Name: "value", Required: 1, Function: e.primValue}
	e.functionTable["value"] = Obj(&tmp.scmHead)
	tmp = &ScmPrimitive{scmHead: Primitive, Name: "eval-kl", Required: 1, Function: e.primEvalKL}
	e.functionTable["eval-kl"] = Obj(&tmp.scmHead)
	tmp = &ScmPrimitive{scmHead: Primitive, Name: "load-file", Required: 1, Function: e.primLoadFile}
	e.functionTable["load-file"] = Obj(&tmp.scmHead)

	e.symbolTable["*stinput*"] = Make_stream(os.Stdin)
	e.symbolTable["*stoutput*"] = Make_stream(os.Stdout)
	dir, _ := os.Getwd()
	e.symbolTable["*home-directory*"] = Make_string(dir)
	e.symbolTable["*language*"] = Make_string("Go")
	e.symbolTable["*implementation*"] = Make_string("interpreter")
	e.symbolTable["*relase*"] = Make_string(runtime.Version())
	e.symbolTable["*os*"] = Make_string(runtime.GOOS)
	e.symbolTable["*porters*"] = Make_string("Arthur Mao")
	e.symbolTable["*port*"] = Make_string("0.0.1")

	// Extended by shen-go implementation
	e.symbolTable["*package-path*"] = Make_string(PackagePath())
	return &e
}

func (e *Evaluator) primSet(args ...Obj) Obj {
	return PrimSet(e.symbolTable, args[0], args[1])
}

func (e *Evaluator) primValue(args ...Obj) Obj {
	return PrimValue(e.symbolTable, args[0])
}

func (e *Evaluator) primEvalKL(args ...Obj) Obj {
	return e.trampoline(args[0], nil)
}

func (e *Evaluator) primLoadFile(args ...Obj) Obj {
	path := mustString(args[0])
	return e.LoadFile(path)
}

func (e *Evaluator) LoadFile(file string) Obj {
	var filePath string
	if _, err := os.Stat(file); err == nil {
		filePath = file
	} else {
		filePath = path.Join(PackagePath(), file)
		if _, err := os.Stat(filePath); err != nil {
			return Make_error(err.Error())
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return Make_error(err.Error())
	}
	defer f.Close()

	r := NewSexpReader(f)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return Make_error(err.Error())
			}
			break
		}

		res := e.trampoline(exp, nil)
		if *res == Error {
			return res
		}
		fmt.Printf("%#v\n", (*scmHead)(res))
	}
	return Nil
}

func (e *Evaluator) Eval(exp Obj) (res Obj) {
	defer func() {
		if r := recover(); r != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Recovered in Eval:", r)
			fmt.Println(string(buf[:n]))
			res = Nil
		}
	}()
	res = e.trampoline(exp, nil)
	return
}

type controlFlowKind int

const (
	controlFlowReturn controlFlowKind = iota
	controlFlowEval
	controlFlowApply
)

type controlFlow struct {
	kind        controlFlowKind
	inException bool
	// arguments for eval
	exp Obj
	env *Environment
	// arguments for apply
	f    Obj
	args []Obj
	// return result
	result Obj
}

// trampoline is introduced for tail call optimization.
func (e *Evaluator) trampoline(exp Obj, env *Environment) Obj {
	var ctl = controlFlow{
		kind: controlFlowEval,
		exp:  exp,
		env:  env,
	}
	for ctl.kind != controlFlowReturn {
		switch ctl.kind {
		case controlFlowEval:
			e.eval(&ctl)
		case controlFlowApply:
			e.apply(&ctl)
		}
	}
	return ctl.result
}

func (ctl *controlFlow) TailEval(exp Obj, env *Environment) {
	ctl.exp = exp
	ctl.env = env
	ctl.kind = controlFlowEval
}

func (ctl *controlFlow) TailApply(f Obj, args []Obj) {
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

func (e *Evaluator) eval(ctl *controlFlow) {
	exp := ctl.exp
	env := ctl.env

	switch *exp { // handle constant
	case Number, String, Vector, Boolean, Null, Procedure, Primitive:
		ctl.Return(exp)
		return
	}

	if ok, sym := isSymbol(exp); ok {
		if val, ok := env.Get(sym.sym); ok {
			exp = val
		}
		ctl.Return(exp)
		return
	}

	pair := mustPair(exp)
	if ok, sym := isSymbol(pair.car); ok {
		exp = pair.cdr // handle special form
		switch sym.sym {
		case "quote":
			// Extension to make vm work.
			// TODO: remove it later
			ctl.Return(car(exp))
			return
		case "defun": // (defun f (x y) z)
			proc := Make_procedure(cadr(exp), caddr(exp), env)
			funName := car(exp)
			e.functionTable[mustSymbol(funName).sym] = proc
			ctl.Return(funName)
			return
		case "lambda": // (lambda x x)
			ctl.Return(Make_procedure(car(exp), cadr(exp), env))
			return
		case "freeze": // (freeze body)
			ctl.Return(Make_procedure(Nil, car(exp), env))
			return
		case "let": // (let x y z)
			args := e.trampoline(cadr(exp), env)
			if *args == Error {
				ctl.Exception(args)
				return
			}
			newEnv := env.Extend([]Obj{car(exp)}, []Obj{args})
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
			if tmp := e.trampoline(car(exp), env); *tmp == Error {
				ctl.Exception(tmp)
				return
			}
			ctl.TailEval(cadr(exp), env)
			return
		}
	}

	fn := e.evalFunction(pair.car, env)
	if *fn == Error {
		ctl.Exception(fn)
		return
	}
	args := e.evalArgumentList(pair.cdr, env)
	if !ctl.inException && len(args) == 1 && *args[0] == Error {
		ctl.Exception(args[0])
		return
	}
	ctl.TailApply(fn, args)
	return
}

func (e *Evaluator) evalFunction(fn Obj, env *Environment) Obj {
	if ok, sym := isSymbol(fn); ok {
		if proc, ok := env.Get(sym.sym); ok {
			return proc
		}
		if val, ok := e.functionTable[sym.sym]; ok {
			return val
		}
	}

	switch *fn {
	case Primitive, Procedure:
		return fn
	case Pair:
		return e.trampoline(fn, env)
	}
	panic(fmt.Sprintf("can't apply non function: %#v", (*scmHead)(fn)))
}

func (e *Evaluator) evalIf(a, b, c Obj, env *Environment, ctl *controlFlow) {
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

func (e *Evaluator) evalAnd(a, b Obj, env *Environment, ctl *controlFlow) {
	if e.trampoline(a, env) == False {
		ctl.Return(False)
		return
	}
	ctl.TailEval(b, env)
}

func (e *Evaluator) evalOr(a, b Obj, env *Environment, ctl *controlFlow) {
	if e.trampoline(a, env) == True {
		ctl.Return(True)
		return
	}
	ctl.TailEval(b, env)
}

func (e *Evaluator) evalCond(l Obj, env *Environment, ctl *controlFlow) {
	for *l == Pair {
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

func (e *Evaluator) evalTrapError(exp Obj, env *Environment, ctl *controlFlow) {
	v := e.trampoline(car(exp), env)
	if *v == Error {
		ctl.inException = true
		handler := e.evalFunction(cadr(exp), env)
		ctl.TailApply(handler, []Obj{v})
		return
	}
	ctl.Return(v)
	return
}

func (e *Evaluator) apply(ctl *controlFlow) {
	f := ctl.f
	args := ctl.args

	if *f == Primitive {
		prim := mustPrimitive(f)
		switch {
		case len(args) < prim.Required:
			ctl.Return(partialApply(prim.Required, args, nil, f))
			return
		case len(args) == prim.Required:
			ctl.Return(prim.Function(args...))
			return
		}
	} else if *f == Procedure {
		proc := mustProcedure(f)
		switch {
		case len(args) < proc.arity:
			ctl.Return(partialApply(proc.arity, args, proc.env, f))
			return
		case len(args) == proc.arity:
			newEnv := proc.env.Extend(proc.arg, args)
			ctl.TailEval(proc.body, newEnv)
			return
		case len(args) > proc.arity:
			newEnv := proc.env.Extend(proc.arg, args[:proc.arity])
			res := e.trampoline(proc.body, newEnv)
			ctl.TailApply(res, args[proc.arity:])
			return
		}
	}
	panic("can't apply object")
}

// partialApply works when Required > providArgs
func partialApply(required int, providArgs []Obj, env *Environment, proc Obj) Obj {
	// Partial apply...
	// (f x y z) => (lambda (z) (f x y z)) with x y in env
	symbols := makeTempSymbols(required)
	env1 := env.Extend(symbols[:len(providArgs)], providArgs)

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

	return Make_procedure(args, body, env1)
}

func (e *Evaluator) evalArgumentList(args Obj, env *Environment) []Obj {
	var ret []Obj
	for *args == Pair {
		v := e.trampoline(car(args), env)
		if *v == Error {
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
		ret[i] = Make_symbol(fmt.Sprintf("tmp%d", i))
	}
	return ret
}
