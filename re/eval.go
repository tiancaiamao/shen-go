package re

import (
	"bytes"
//	"fmt"
	"math"
	"strconv"
)

// =======================================
//          Object representation
// =======================================

type Obj interface {
	String() string
}

type Integer int64

func (i Integer) String() string {
	return strconv.FormatInt(int64(i), 10)
}

type Float64 float64

func (f Float64) String() string {
	return strconv.FormatFloat(float64(f), 'E', -1, 64)
}

// isInteger determinate whether a float64 is actually a precise integer.
// Judge is according to IEEE754 standard.
func isPreciseInteger(f float64) bool {
	exp := math.Ilogb(f)
	if exp < 0 && exp != math.MinInt32 {
		return false
	}

	if exp >= 52 {
		return true
	}

	bits := math.Float64bits(f)
	return (bits << uint(12+exp)) == 0
}

func MakeNumber(f float64) Obj {
	if isPreciseInteger(f) {
		return Integer(int(f))
	}
	return Float64(f)
}

type booleanObj bool

func (b booleanObj) String() string {
	return strconv.FormatBool(bool(b))
}

var True = booleanObj(true)
var False = booleanObj(false)

type nilObj struct{}

func (_ nilObj) String() string {
	return "nil"
}

var Nil Obj = nilObj{}

type Symbol struct {
	str string
	val Obj
}

func (s *Symbol) String() string {
	return s.str
}

var symbolMap = make(map[string]*Symbol)

var symIf, symDo, symLambda Obj
var symQuote Obj

func init() {
	symIf = MakeSymbol("if")
	symDo = MakeSymbol("do")
	symLambda = MakeSymbol("lambda")
	symQuote = MakeSymbol("quote")
}

func MakeSymbol(str string) *Symbol {
	if s, ok := symbolMap[str]; ok {
		return s
	}
	s := &Symbol{
		str: str,
	}
	symbolMap[str] = s
	return s
}

type String string

func (s String) String() string {
	return strconv.Quote(string(s))
}

type Cons struct {
	car Obj
	cdr Obj
}

func (c *Cons) String() string {
	var buf bytes.Buffer
	buf.WriteString("(cons ")
	buf.WriteString(c.car.String())
	buf.WriteString(" ")
	buf.WriteString(c.cdr.String())
	buf.WriteString(")")
	return buf.String()
}

type Closure struct {
	Required int
	Code     Instr
	Data     []Obj // The closure value
}

func (c *Closure) String() string {
	return "#closure"
}

type Continuation struct {
	Stack savedStackFrame
	Code  Instr
}

func (c Continuation) String() string {
	return "#cont"
}

type Primitive struct {
	Exec func(vm *VM)
}

func (c *Primitive) String() string {
	return "#primitive"
}

// =======================================
// 	VM and instructions
// =======================================

type VM struct {
	pc    Instr
	stack Frame
}

func New() *VM {
	var s stackFrame
	s.Push(Continuation{
		Stack: savedStackFrame{
			stackFrame: stackFrame{
				base: s.base,
				pos:  0,
			},
		},
		Code: nil,
	})
	return &VM{
		stack: &s,
	}
}

type Frame interface {
	Get(idx int) Obj
	Set(idx int, v Obj)
	Push(Obj)
	Pop() Obj
	Resize(n int)
}

type stackFrame struct {
	base []Obj
	pos  int
}

type savedStackFrame struct {
	stackFrame
	size int
}

func (s *stackFrame) Get(idx int) Obj {
	if idx >= 0 {
		return s.base[s.pos+idx]
	}
	return s.base[len(s.base)+idx]
}

func (s *stackFrame) Set(idx int, v Obj) {
	if idx >= 0 {
		s.base[s.pos+idx] = v
	} else {
		s.base[len(s.base)+idx] = v
	}
}

func (s *stackFrame) Resize(n int) {
	s.base = s.base[:n]
}

func (s *stackFrame) Push(o Obj) {
	s.base = append(s.base, o)
}

func (s *stackFrame) Pop() Obj {
	tos := len(s.base) - 1
	if tos >= len(s.base) {
		panic("pop() empty stack")
	}
	res := s.base[tos]
	s.base = s.base[:tos]
	return res
}

func (vm *VM) push(v Obj) {
	vm.stack.Push(v)
}

func (vm *VM) pop() Obj {
	return vm.stack.Pop()
}

func (vm *VM) run(code Instr) {
	for vm.pc = code; vm.pc != nil; {
		vm.pc.Exec(vm)
	}
}

var _ Instr = InstrConst{}
var _ Instr = InstrIf{}
var _ Instr = InstrLocalRef{}
var _ Instr = InstrSet{}
var _ Instr = InstrCall{}
var _ Instr = InstrDo{}
var _ Instr = InstrGlobalRef{}

type Instr interface {
	Exec(*VM)
}

type InstrConst struct {
	Val  Obj
	Next Instr
}

func (c InstrConst) Exec(vm *VM) {
	vm.push(c.Val)
	vm.pc = c.Next
}

type InstrGlobalRef struct {
	sym  *Symbol
	Next Instr
}

func (i InstrGlobalRef) Exec(vm *VM) {
	if i.sym.val == nil {
		panic("undefined value for symbol" + i.sym.str)
	}
	vm.push(i.sym.val)
	vm.pc = i.Next
}

type InstrSet struct {
	Sym  Obj
	Next Instr
}

func (s InstrSet) Exec(vm *VM) {
	v := vm.pop()
	s.Sym.(*Symbol).val = v
	vm.pc = s.Next
}

type InstrDo struct {
	Next Instr
}

func (i InstrDo) Exec(vm *VM) {
	vm.pop()
	vm.pc = i.Next
}

type InstrLocalRef struct {
	Idx  int
	Next Instr
}

func (i InstrLocalRef) Exec(vm *VM) {
	v := vm.stack.Get(i.Idx + 2)
	vm.push(v)
	vm.pc = i.Next
}

type InstrIf struct {
	Then Instr
	Else Instr
}

func (i InstrIf) Exec(vm *VM) {
	v := vm.pop()
	switch v {
	case True:
		vm.pc = i.Then
	case False:
		vm.pc = i.Else
	default:
		panic("if instr accept only true and false")
	}
}

type InstrPrepareCall struct {
	next Instr
}

func (c InstrPrepareCall) Exec(vm *VM) {
	// This slot is reserved for continuation
	vm.stack.Push(Nil)
	vm.pc = c.next
}

type InstrCall struct {
	size int
	Next Instr
}

func (c InstrCall) Exec(vm *VM) {
	fn := vm.stack.Get(-c.size)
	raw := fn.(*Closure)
	code := raw.Code
	if c.Next == nil { // Jump
		// Reuse the old stack.
		for i := 0; i < c.size; i++ {
			arg := vm.stack.Get(-c.size + i)
			vm.stack.Set(i+1, arg)
		}
		vm.stack.Resize(c.size + 1)
	} else { // Call
		sf := vm.stack.(*stackFrame)
		newStack := &stackFrame{
			base: sf.base,
			pos:  len(sf.base) - c.size - 1,
		}
		// Save the stack.
		// Save the Next instr, after the call, the function return to here.
		cc := Continuation{
			Stack: savedStackFrame{
				stackFrame: stackFrame{
					base: sf.base,
					pos:  sf.pos,
				},
				size: newStack.pos,
			},
			Code: c.Next,
		}
		newStack.Set(-c.size-1, cc)
		// change current pc to the callee
		vm.stack = newStack
	}
	vm.pc = code
}

type InstrPrimitive struct {
	prim *Primitive
	Next Instr
}

func (c InstrPrimitive) Exec(vm *VM) {
	raw := c.prim
	// Execute the primitive
	raw.Exec(vm)
	if c.Next == nil { // Jump
		val := vm.stack.Get(-1)
		cc := vm.stack.Get(0)
		// Return to the previous saved continuation.
		cont := cc.(Continuation)
		vm.stack = &stackFrame{
			base: cont.Stack.base[:cont.Stack.size],
			pos:  cont.Stack.pos,
		}
		vm.push(val)
		vm.pc = cont.Code
	} else { // Call
		vm.pc = c.Next
	}
}

// =====================================
// 	Compiler utilities
// =====================================

type Env struct {
	parent *Env
	args   Obj
}

func (env *Env) findVariable(s Obj) (m int, n int) {
	e := env
	for e != nil {
		n = 0
		for args := e.args; args != Nil; args = cdr(args) {
			if car(args) == s {
				return
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

func compile(exp Obj, env *Env, cont Instr) Instr {
	switch raw := exp.(type) {
	case nilObj, booleanObj, Integer, String, Float64:
		return InstrConst{
			Val:  exp,
			Next: cont,
		}
	case *Symbol:
		m, n := env.findVariable(exp)
		if m == 0 {
			return InstrLocalRef{
				Idx:  n,
				Next: cont,
			}
		}
		if m > 0 {
			// Closure value
		}
		// Global variable
		return InstrGlobalRef{
			sym:  raw,
			Next: cont,
		}
	}

	raw := exp.(*Cons)
	switch raw.car {
	case symQuote:
		return InstrConst{
			Val:  car(raw.cdr),
			Next: cont,
		}
	case symIf:
		thenCont := compile(cadr(raw.cdr), env, cont)
		elseCont := compile(caddr(raw.cdr), env, cont)
		return compile(car(raw.cdr), env, InstrIf{
			Then: thenCont,
			Else: elseCont,
		})
	case symDo:
		return compile(cadr(raw), env, InstrDo{
			Next: compile(caddr(raw), env, cont),
		})
	case symLambda:
		args := car(raw.cdr)
		body := cadr(raw.cdr)
		newEnv := &Env{
			parent: env,
			args:   args,
		}
		code := compile(body, newEnv, InstrPrimitive{
			prim: primIdentify,
		})
		return InstrConst{
			Val: &Closure{
				Required: listLength(args),
				Code:     code,
			},
			Next: cont,
		}
		// case symLet:
	}

	return compileCall(exp, env, cont)
}

func compileCall(exp Obj, env *Env, cont Instr) Instr {
	if sym, ok := car(exp).(*Symbol); ok {
		switch sym.str {
		case "+":
			return compileList(cdr(exp), env, InstrPrimitive{
				prim: primAdd,
				Next: cont,
			})
		case "-":
			return compileList(cdr(exp), env, InstrPrimitive{
				prim: primSub,
				Next: cont,
			})
		case "=":
			return compileList(cdr(exp), env, InstrPrimitive{
				prim: primEQ,
				Next: cont,
			})
		case "set":
			return compileList(cdr(exp), env, InstrPrimitive{
				prim: primSet,
				Next: cont,
			})
		case "value":
			return compileList(cdr(exp), env, InstrPrimitive{
				prim: primValue,
				Next: cont,
			})
		}
	}
	size := listLength(exp)
	return InstrPrepareCall{
		next: compileList(exp, env, InstrCall{
			size: size,
			Next: cont,
		}),
	}
}

func compileList(exp Obj, env *Env, cont Instr) Instr {
	if exp == Nil {
		return cont
	}
	return compile(car(exp), env, compileList(cdr(exp), env, cont))
}

// ==================================
// 	Library and primitive functions
// ==================================

func cons(a, b Obj) Obj {
	return &Cons{a, b}
}

func car(o Obj) Obj {
	return o.(*Cons).car
}

func cdr(o Obj) Obj {
	return o.(*Cons).cdr
}

func cadr(o Obj) Obj {
	return car(cdr(o))
}

func caddr(o Obj) Obj {
	return car(cdr(cdr(o)))
}

func reverse(o Obj) Obj {
	ret := Nil
	for o != Nil {
		ret = cons(car(o), ret)
		o = cdr(o)
	}
	return ret
}

func listLength(l Obj) int {
	count := 0
	for l != Nil {
		count++
		l = cdr(l)
	}
	return count
}

func eval(vm *VM, exp Obj) Obj {
	var env *Env
	code := compile(exp, env, nil)
	//	fmt.Printf("the code is: %#v\n", code)
	vm.run(code)
	ret := vm.stack.Get(-1)
	//	vm.stack = vm.stack[:0]
	return ret
}

var primSet = &Primitive{
	Exec: func(vm *VM) {
		val := vm.pop()
		key := vm.pop()
		key.(*Symbol).val = val
		vm.push(val)
	},
}

var primValue = &Primitive{
	Exec: func(vm *VM) {
		key := vm.pop()
		val := key.(*Symbol).val
		vm.push(val)
	},
}

var primAdd = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop().(Integer)
		y := vm.pop().(Integer)
		vm.push(Integer(x + y))
	},
}

var primSub = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop().(Integer)
		y := vm.pop().(Integer)
		vm.push(Integer(y - x))
	},
}

var primEQ = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop()
		y := vm.pop()
		if x == y {
			vm.push(True)
		} else {
			vm.push(False)
		}
	},
}

var primIdentify = &Primitive{
	Exec: func(vm *VM) {
	},
}
