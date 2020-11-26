package kl

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

type Evaluator interface {
	base() *ControlFlow
	eval()

	// All the following methods are provided by ControlFlow
	TailEval(exp Obj, env Obj)
	TailApply(f Obj, args ...Obj)
	Return(result Obj)
	Get(n int) Obj
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

func (ctl *ControlFlow) base() *ControlFlow {
	return ctl
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
	return ctl.data[n]
}

// trampoline is introduced for tail call optimization.
func trampoline(e Evaluator) Obj {
	ctl := e.base()
	for ctl.kind != ControlFlowReturn {
		switch ctl.kind {
		case ControlFlowEval:
			e.eval()
		case ControlFlowApply:
			apply(e)
		}
	}
	ret := ctl.data[ctl.pos]
	ctl.data = ctl.data[:ctl.pos]
	return ret
}

func evalExp(e Evaluator, exp Obj, env Obj) Obj {
	e.TailEval(exp, env)
	return trampoline(e)
}

func evalIf(e Evaluator, a, b, c Obj, env Obj) {
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

func evalAnd(e Evaluator, a, b Obj, env Obj) {
	if evalExp(e, a, env) == False {
		e.Return(False)
		return
	}
	e.TailEval(b, env)
}

func evalOr(e Evaluator, a, b Obj, env Obj) {
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

func Call(e Evaluator, f Obj, args ...Obj) Obj {
	e.TailApply(f, args...)
	return trampoline(e)
}

func Eval(e Evaluator, exp Obj) (res Obj) {
	defer func() {
		if r := recover(); r != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Panic:", r)
			fmt.Println("Recovered in Eval:", ObjString(exp))
			fmt.Println(string(buf[:n]))
			res = Nil
		}
	}()
	res = evalExp(e, exp, Nil)
	return
}

type tryResult struct {
	e    Evaluator
	data Obj
}

func Try(e Evaluator, f Obj) (res tryResult) {
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

func loadFile(e Evaluator, extended bool, file string) Obj {
	var filePath string
	if _, err := os.Stat(file); err == nil {
		filePath = file
	} else {
		filePath = path.Join(PackagePath(), file)
		if _, err := os.Stat(filePath); err != nil {
			return MakeError(err.Error())
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return MakeError(err.Error())
	}
	defer f.Close()

	r := NewSexpReader(f, extended)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return MakeError(err.Error())
			}
			break
		}

		// Macro expand for cora.
		if extended {
			expand := CoraValue(symMacroExpand)
			if expand != Nil {
				exp = Call(e, expand, exp)
			}
		}

		res := evalExp(e, exp, Nil)
		if *res == scmHeadError {
			return res
		}
	}
	return MakeString(file)
}

func apply(e Evaluator) {
	ctl := e.base()
	f := ctl.data[ctl.pos]
	args := ctl.data[ctl.pos+1:]

	if *f == scmHeadPrimitive {
		prim := mustPrimitive(f)
		switch {
		case len(args) < prim.Required:
			ctl.Return(partialApply(prim.Required, args, Nil, f))
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
			res := evalExp(e, proc.body, newEnv)
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
			fn.fn(e, tmp...)
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
		f = Call(e, f, tmp...)
		args = args[taken:]
		ctl.TailApply(f, args...)
		return
	}
	panic("can't apply object")
}
