package re

import (
	"bytes"
	// "fmt"
	"io"
	"math"
	"strconv"

	"time"
)

// =======================================
//          Object representation
// =======================================

var _ Obj = Integer(0)
var _ Obj = Float64(0.0)
var _ Obj = String("")
var _ Obj = True
var _ Obj = False
var _ Obj = Nil
var _ Obj = &Symbol{}
var _ Obj = &Closure{}
var _ Obj = &Cons{}
var _ Obj = Vector{}
var _ Obj = Stream{}
var _ Obj = Error{}

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
	switch b {
	case True:
		return "true"
	case False:
		return "false"
	default:
		panic("bool type error")
	}
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
	fn  Obj
}

func (s *Symbol) String() string {
	return s.str
}

var symbolMap = make(map[string]*Symbol)

var symIf, symDo, symLambda, symDefun Obj
var symCond, symAnd, symOr, symLet, symType, symFreeze, symTrapError, symSimpleError Obj
var uptime time.Time

// var symQuote Obj

var klMacro map[Obj]func(*Compiler, Obj, *Env, Instr) Instr

func init() {
	uptime = time.Now()
	symIf = MakeSymbol("if")
	symDo = MakeSymbol("do")
	symLambda = MakeSymbol("lambda")
	// symQuote = MakeSymbol("quote")
	symDefun = MakeSymbol("defun")
	symLet = MakeSymbol("let")
	symCond = MakeSymbol("cond")
	symAnd = MakeSymbol("and")
	symOr = MakeSymbol("or")
	symTrapError = MakeSymbol("trap-error")
	symSimpleError = MakeSymbol("simple-error")

	klMacro = map[Obj]func(*Compiler, Obj, *Env, Instr) Instr{
		symDefun:       compileDefunMacro,
		symCond:        compileCondMacro,
		symAnd:         compileAndMacro,
		symOr:          compileOrMacro,
		symTrapError:   compileTrapErrorMacro,
		symSimpleError: compileSimpleErrorMacro,
		symFreeze:      compileFreezeMacro,
		symLet:         compileLetMacro,
		symType:        compileTypeMacro,
	}

	symTry := MakeSymbol("try")
	symTry.fn = &Closure{
		Required: 2,
		Code:     instrFunc(instrForTry),
		Slot:     nil,
	}
	symThrow := MakeSymbol("throw")
	symThrow.fn = &Closure{
		Required: 3,
		Code:     instrFunc(instrForThrow),
		Slot:     nil,
	}
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
	Parent   *Closure
	Slot     map[int]Obj
}

func (c *Closure) String() string {
	return "#0closure"
}

type Curry struct {
	Required int
	Closed   []Obj
	Prim     *Primitive
}

func (c *Curry) String() string {
	return "#0curry"
}

type Stream struct {
	io.ReadWriter
}

func (s Stream) String() string {
	return "#stream"
}

type Error struct {
	error
}

func (err Error) String() string {
	return err.Error()
}

type Vector []Obj

func (v Vector) String() string {
	return "#vector"
}

type Continuation struct {
	Stack stackFrame
	Code  Instr
}

func (c Continuation) String() string {
	return "#cont"
}

type Primitive struct {
	Exec     func(vm *VM)
	Required int
}

func (c *Primitive) String() string {
	return "#primitive"
}

// =======================================
// 	VM and instructions
// =======================================

type VM struct {
	pc    Instr
	stack stackFrame
	val   Obj
}

func New() *VM {
	var initializeStack [2048]Obj
	return &VM{
		stack: stackFrame{
			underlying: &initializeStack,
		},
	}
}

type stackFrame struct {
	underlying *[2048]Obj
	base       int
	pos        int
}

func (sf *stackFrame) Push(o Obj) {
	sf.underlying[sf.pos] = o
	sf.pos++
}

func (sf *stackFrame) Get(idx int) Obj {
	if idx >= 0 {
		return sf.underlying[sf.base+idx]
	}
	return sf.underlying[sf.pos+idx]
}

func (sf *stackFrame) Set(idx int, v Obj) {
	sf.underlying[sf.base+idx] = v
}

func (sf *stackFrame) Pop() Obj {
	if sf.pos <= sf.base {
		panic("pop empty stack")
	}
	sf.pos--
	return sf.underlying[sf.pos]
}

func (sf *stackFrame) Resize(n int) {
	sf.pos = sf.base + n
}

func (vm *VM) push(v Obj) {
	vm.stack.Push(v)
}

func (vm *VM) pop() Obj {
	return vm.stack.Pop()
}

func (vm *VM) Return(x Obj) {
	cc := vm.stack.Get(0)
	cont := cc.(Continuation)
	vm.val = x
	vm.stack = cont.Stack
	vm.pc = cont.Code
}

func (vm *VM) run(code Instr) {
	cc := Continuation{
		Stack: vm.stack,
		Code:  nil,
	}
	vm.stack.base = vm.stack.pos
	vm.push(cc)

	for vm.pc = code; vm.pc != nil; {
		vm.pc.Exec(vm)
	}
}

var _ Instr = InstrConst{}
var _ Instr = InstrIf{}
var _ Instr = InstrNop{}
var _ Instr = InstrPush{}
var _ Instr = InstrLocalRef{}
var _ Instr = InstrClosureRef{}
var _ Instr = InstrGlobalRef{}
var _ Instr = InstrPrimitive{}
var _ Instr = InstrCall{}
var _ Instr = InstrSet{}
var _ Instr = InstrMakeClosure{}
var _ Instr = InstrExit{}

type Instr interface {
	Exec(*VM)
}

type InstrExit struct {}

func (_ InstrExit) Exec(vm *VM) {
	vm.Return(vm.val)
}

type InstrConst struct {
	Val  Obj
	Next Instr
}

func (c InstrConst) Exec(vm *VM) {
	vm.val = c.Val
	vm.pc = c.Next
}

type InstrGlobalRef struct {
	sym  *Symbol
	Next Instr
}

func (i InstrGlobalRef) Exec(vm *VM) {
	if i.sym.fn == nil {
		panic("undefined fn for symbol:" + i.sym.str)
	}
	vm.val = i.sym.fn
	vm.pc = i.Next
}

type InstrSet struct {
	Sym  Obj
	Next Instr
}

func (s InstrSet) Exec(vm *VM) {
	f := vm.val
	s.Sym.(*Symbol).fn = f
	vm.val = s.Sym
	vm.pc = s.Next
}

type InstrNop struct {
	Next Instr
}

func (i InstrNop) Exec(vm *VM) {
	vm.pc = i.Next
}

type InstrPush struct {
	Next Instr
}

func (i InstrPush) Exec(vm *VM) {
	vm.push(vm.val)
	vm.pc = i.Next
}

type InstrLocalRef struct {
	Idx  int
	Next Instr
}

func (i InstrLocalRef) Exec(vm *VM) {
	v := vm.stack.Get(i.Idx + 2)
	vm.val = v
	vm.pc = i.Next
}

type InstrClosureRef struct {
	Up   int
	Idx  int
	Next Instr
}

func (i InstrClosureRef) Exec(vm *VM) {
	tmp := vm.stack.Get(1)
	clo := tmp.(*Closure)
	for v := i.Up; v > 0; v-- {
		clo = clo.Parent
	}

	vm.val = clo.Slot[i.Idx]
	vm.pc = i.Next
}

type InstrIf struct {
	Then Instr
	Else Instr
}

func (i InstrIf) Exec(vm *VM) {
	switch vm.val {
	case True:
		vm.pc = i.Then
	case False:
		vm.pc = i.Else
	default:
		panic("if instr accept only true and false")
	}
}

type InstrMakeClosure struct {
	closed   []int
	code     Instr
	required int
	Next     Instr
}

func (i InstrMakeClosure) Exec(vm *VM) {
	var parent *Closure
	if raw, ok := vm.stack.Get(1).(*Closure); ok {
		parent = raw
	}
	var slot map[int]Obj
	if len(i.closed) > 0 {
		slot = make(map[int]Obj, len(i.closed))
		for _, off := range i.closed {
			slot[off] = vm.stack.Get(off + 2)
		}
	}
	vm.val = &Closure{
		Required: i.required,
		Code:     i.code,
		Slot:     slot,
		Parent:   parent,
	}
	vm.pc = i.Next
}

type InstrPrepareCall struct {
	next Instr
}

func (c InstrPrepareCall) Exec(vm *VM) {
	// This slot is reserved for continuation
	vm.push(Nil)
	vm.pc = c.next
}

type InstrCall struct {
	size int
	Next Instr
}

func (c InstrCall) Exec(vm *VM) {
	fn := vm.stack.Get(-c.size)
	switch raw := fn.(type) {
	case *Closure:
		c.callClosure(vm, raw)
	case *Curry:
		c.callCurry(vm, raw)
	default:
		panic("can't call " + fn.String())
	}
}

func (c InstrCall) callCurry(vm *VM, curry *Curry) {
	tmp := append(curry.Closed, vm.stack.underlying[vm.stack.pos-c.size+1:vm.stack.pos]...)
	vm.stack.pos = vm.stack.pos - c.size
	for _, v := range tmp {
		vm.stack.Push(v)
	}

	if curry.Prim != nil {
		vm.pc = InstrPrimitive{
			size: len(tmp),
			prim: curry.Prim,
			Next: c.Next,
		}
	} else {
		vm.pc = InstrCall{
			size: len(tmp),
			Next: c.Next,
		}
	}
}

func (c InstrCall) callClosure(vm *VM, clo *Closure) {
	argc := c.size - 1
	switch {
	case argc == clo.Required:
		c.callClosureNormal(vm, clo)
	case argc < clo.Required:
		vm.val = &Curry{
			Required: clo.Required - argc,
			Closed:   append([]Obj{}, vm.stack.underlying[vm.stack.pos-c.size:vm.stack.pos]...),
		}
		vm.stack.Resize(vm.stack.pos - c.size - 1)
		if c.Next == nil {
			panic("call closure")
		} else {
			vm.pc = c.Next
		}
	case argc > clo.Required:
		c.callClosurePartial(vm, clo)
	}
}

func (c InstrCall) callClosureNormal(vm *VM, raw *Closure) {
	code := raw.Code
	if c.Next == nil {
		panic("now never here?")
	} else if c.Next == identity { // Jump
		// Reuse the old stack.
		for i := 0; i < c.size; i++ {
			arg := vm.stack.Get(-c.size + i)
			vm.stack.Set(i+1, arg)
		}
		vm.stack.Resize(c.size + 1)
	} else { // Call
		newStackBase := vm.stack.pos - c.size - 1
		// Save the old stack and save the Next instr.
		//  After the call, the function return to here.
		cc := Continuation{
			Stack: stackFrame{
				underlying: vm.stack.underlying,
				base:       vm.stack.base,
				pos:        newStackBase,
			},
			Code: c.Next,
		}

		// Change to the new stack.
		vm.stack.base = newStackBase
		vm.stack.Set(0, cc)
	}
	// change current pc to the callee
	vm.pc = code
}

func (c InstrCall) callClosurePartial(vm *VM, clo *Closure) {
	// prepare a new stack for call
	cc := Continuation{
		Stack: vm.stack,
		Code:  nil,
	}
	newBase := vm.stack.pos
	vm.push(cc)
	for i := 0; i < clo.Required+1; i++ {
		vm.push(vm.stack.Get(i + 2))
	}
	vm.stack.base = newBase

	// Call partial
	for vm.pc = clo.Code; vm.pc != nil; {
		vm.pc.Exec(vm)
	}

	// Handle the remain call
	vm.stack.Set(2, vm.val)
	for i := 0; i < c.size-1-clo.Required; i++ {
		vm.stack.Set(i+3, vm.stack.Get(i+3+clo.Required))
	}
	vm.stack.Resize(c.size - clo.Required + 2)

	c1 := InstrCall{
		size: c.size - clo.Required,
		Next: c.Next,
	}
	vm.pc = c1
}

type InstrPrimitive struct {
	size int
	prim *Primitive
	Next Instr
}

func (c InstrPrimitive) Exec(vm *VM) {
	raw := c.prim
	if c.size == raw.Required {
		// Execute the primitive
		raw.Exec(vm)
	} else if c.size < raw.Required {
		vm.val = &Curry{
			Required: raw.Required - c.size,
			Closed:   append([]Obj{}, vm.stack.underlying[vm.stack.pos-c.size:vm.stack.pos]...),
			Prim:     c.prim,
		}
	} else {
		panic("call primitive argc not match")
	}

	if c.Next == nil { // Jump
		vm.Return(vm.val)
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

type pos struct {
	m int
	n int
}

type Compiler struct {
	closed []pos
}

func (c *Compiler) compile(exp Obj, env *Env, cont Instr) Instr {
	switch raw := exp.(type) {
	case nilObj, booleanObj, Integer, String, Float64:
		return InstrConst{
			Val:  exp,
			Next: cont,
		}
	case *Symbol:
		return c.compileSymbol(raw, env, false, cont)
	}

	raw := exp.(*Cons)
	if _, ok := raw.car.(*Symbol); ok {
		if fn, ok := klMacro[raw.car]; ok {
			return fn(c, cdr(exp), env, cont)
		}
	}

	switch raw.car {
	case symIf:
		thenCont := c.compile(cadr(raw.cdr), env, cont)
		elseCont := c.compile(caddr(raw.cdr), env, cont)
		return c.compile(car(raw.cdr), env, InstrIf{
			Then: thenCont,
			Else: elseCont,
		})
	case symDo:
		return c.compile(cadr(raw), env, InstrNop{
			Next: c.compile(caddr(raw), env, cont),
		})
	case symLambda:
		args := car(raw.cdr)
		if _, ok := args.(*Symbol); ok {
			args = cons(args, Nil)
		}
		body := cadr(raw.cdr)
		newEnv := &Env{
			parent: env,
			args:   args,
		}
		return c.compileLambda(args, body, newEnv, cont)
	}

	return c.compileCall(exp, env, cont)
}

func (c *Compiler) compileSymbol(exp *Symbol, env *Env, isCall bool, cont Instr) Instr {
	m, n := env.findVariable(exp)
	if m == 0 {
		return InstrLocalRef{
			Idx:  n,
			Next: cont,
		}
	}
	if m > 0 {
		// Closure value
		c.closed = append(c.closed, pos{m, n})
		return InstrClosureRef{
			Up:   m - 1,
			Idx:  n,
			Next: cont,
		}
	}
	// Global variable
	if isCall {
		return InstrGlobalRef{
			sym:  exp,
			Next: cont,
		}
	}

	return InstrConst{
		Val:  exp,
		Next: cont,
	}
}

var identity = InstrExit{}

func (c *Compiler) compileLambda(args, body Obj, env *Env, cont Instr) Instr {
	var c1 Compiler
	code := c1.compile(body, env, identity)
	closed := make([]int, 0, len(c1.closed))
	for _, p := range c1.closed {
		if p.m == 1 {
			closed = append(closed, p.n)
		} else {
			c.closed = append(c.closed, pos{p.m - 1, p.n})
		}
	}
	return InstrMakeClosure{
		closed:   closed,
		code:     code,
		required: listLength(args),
		Next:     cont,
	}
}

func (c *Compiler) compileCall(exp Obj, env *Env, cont Instr) Instr {
	if sym, ok := car(exp).(*Symbol); ok {
		remain := cdr(exp)
		switch sym.str {
		case "+":
			return c.compileList(remain, env, InstrPrimitive{
				size: listLength(remain),
				prim: primAdd,
				Next: cont,
			})
		case "-":
			return c.compileList(remain, env, InstrPrimitive{
				size: listLength(remain),
				prim: primSub,
				Next: cont,
			})
		case "=":
			return c.compileList(remain, env, InstrPrimitive{
				size: listLength(remain),
				prim: primEQ,
				Next: cont,
			})
		case "*":
			return c.compileList(remain, env, InstrPrimitive{
				size: listLength(remain),
				prim: primMul,
				Next: cont,
			})
		case "/":
			return c.compileList(remain, env, InstrPrimitive{
				size: listLength(remain),
				prim: primDiv,
				Next: cont,
			})
		case "set":
			return c.compileList(remain, env, InstrPrimitive{
				size: listLength(remain),
				prim: primSet,
				Next: cont,
			})
		case "value":
			return c.compileList(remain, env, InstrPrimitive{
				size: listLength(remain),
				prim: primValue,
				Next: cont,
			})
		default:
			size := listLength(exp)
			return InstrPrepareCall{
				next: c.compileSymbol(sym, env, true, InstrPush{
					Next: c.compileList(remain, env, InstrCall{
						size: size,
						Next: cont,
					}),
				}),
			}
		}
	}
	size := listLength(exp)
	return InstrPrepareCall{
		next: c.compileList(exp, env, InstrCall{
			size: size,
			Next: cont,
		}),
	}
}

func (c *Compiler) compileList(exp Obj, env *Env, cont Instr) Instr {
	if exp == Nil {
		return cont
	}
	return c.compile(car(exp), env, InstrPush{
		Next: c.compileList(cdr(exp), env, cont),
	})
}

func compileDefunMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (defun f (x y) ..) => ((fn 'set) (quote f) (lambda (x y) ...))
	fName := car(exp)
	args := cadr(exp)
	body := caddr(exp)
	newEnv := &Env{
		parent: env,
		args:   args,
	}
	return c.compileLambda(args, body, newEnv, InstrSet{
		Sym:  fName,
		Next: cont,
	})
}

func compileOrMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (or x y) => (if x true (if y true false))
	x := car(exp)
	y := cadr(exp)
	tmp := cons(symIf, cons(y, cons(True, cons(False, Nil))))
	exp = cons(symIf, cons(x, cons(True, cons(tmp, Nil))))
	return c.compile(exp, env, cont)
}

func compileAndMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (and x y) => (if x (if y true false) false)
	x := car(exp)
	y := cadr(exp)
	tmp := cons(symIf, cons(y, cons(True, cons(False, Nil))))
	exp = cons(symIf, cons(x, cons(tmp, cons(False, Nil))))
	return c.compile(exp, env, cont)
}

func compileCondMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	exp = compileCondMacroHelp(c, exp)
	return c.compile(exp, env, cont)
}

func compileCondMacroHelp(c *Compiler, exp Obj) Obj {
	// (cond (x y) ...)
	if exp == Nil {
		return cons(MakeSymbol("simple-error"), cons(String("no cond match"), Nil))
	}
	xy := car(exp)
	x := car(xy)
	y := cadr(xy)
	tmp := compileCondMacroHelp(c, cdr(exp))
	return cons(symIf, cons(x, cons(y, cons(tmp, Nil))))
}

func compileLetMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (let x y z) => ((lambda x z) y)
	x := car(exp)
	y := cadr(exp)
	z := cdr(cdr(exp))
	fn := cons(symLambda, cons(x, z))
	exp = cons(fn, cons(y, Nil))
	return c.compile(exp, env, cont)
}

func compileFreezeMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (freeze body) => (lambda () body)
	return c.compileLambda(Nil, car(exp), env, cont)
}

func compileTypeMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (type exp _) => exp
	return c.compile(car(exp), env, cont)
}

func compileTrapErrorMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (trap-error body handler) => (try (lambda (#cc #handler) body) handler)
	body := car(exp)
	handler := cdr(exp)
	action := cons(symLambda, cons(cons(MakeSymbol("#cc"), cons(MakeSymbol("#handler"), Nil)), cons(body, Nil)))
	exp = cons(MakeSymbol("try"), cons(action, handler))
	return c.compile(exp, env, cont)
}

func compileSimpleErrorMacro(c *Compiler, exp Obj, env *Env, cont Instr) Instr {
	// (simple-error v) => (throw v #cc #handler)
	v := car(exp)
	exp = cons(MakeSymbol("throw"), cons(v, cons(MakeSymbol("#cc"), cons(MakeSymbol("#handler"), Nil))))
	return c.compile(exp, env, cont)
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

func (vm *VM) Eval(exp Obj) Obj {
	return eval(vm, exp)
}

func eval(vm *VM, exp Obj) Obj {
	var c Compiler
	code := c.compile(exp, nil, identity)
	//	fmt.Printf("the code is: %#v\n", code)
	vm.run(code)
	ret := vm.val
	return ret
}

var primSet = &Primitive{
	Exec: func(vm *VM) {
		val := vm.pop()
		key := vm.pop()
		key.(*Symbol).val = val
		vm.val = val
	},
	Required: 2,
}

var primValue = &Primitive{
	Exec: func(vm *VM) {
		key := vm.pop()
		val := key.(*Symbol).val
		vm.val = val
	},
	Required: 1,
}

var primAdd = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop().(Integer)
		y := vm.pop().(Integer)
		vm.val = Integer(x + y)
	},
	Required: 2,
}

var primSub = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop().(Integer)
		y := vm.pop().(Integer)
		vm.val = Integer(y - x)
	},
	Required: 2,
}

var primMul = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop().(Integer)
		y := vm.pop().(Integer)
		vm.val = Integer(x * y)
	},
	Required: 2,
}

var primDiv = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop().(Integer)
		y := vm.pop().(Integer)
		vm.val = Integer(y / x)
	},
	Required: 2,
}

var primEQ = &Primitive{
	Exec: func(vm *VM) {
		x := vm.pop()
		y := vm.pop()
		if x == y {
			vm.val = True
		} else {
			vm.val = False
		}
	},
	Required: 2,
}

var primIdentify = &Primitive{
	Exec: func(vm *VM) {
	},
	Required: 1,
}

/*
(defun try (thunk handler)
	(capture (lambda (cc)
				// cc and handler is available in thunk
				(thunk)))))
*/

type instrFunc func(vm *VM)

func (f instrFunc) Exec(vm *VM) {
	f(vm)
}

func instrForTry(vm *VM) {
	cc := vm.stack.Get(0)
	// self :=  vm.stack.Get(1) the 'try' closure itself.
	thunk := vm.stack.Get(2).(*Closure)
	handler := vm.stack.Get(3)

	// Create a new heap allocated frame.
	var arr [2048]Obj
	newStack := stackFrame{
		underlying: &arr,
	}

	// Prepare the call for the chunk.
	newStack.Push(cc)      // saved return address
	newStack.Push(thunk)   // callee: the chunk itself
	newStack.Push(cc)      // cc as the first argument for thunk
	newStack.Push(handler) // handler as the second parameter for thunk

	// Call the chunk, cc and handler is passed as the parameter.
	vm.stack = newStack
	vm.pc = thunk.Code
}

/*
(defun throw (cc handler v)
	(capture (lambda (cc1)
				(let k (make-special-closure cc1)
					// throw longjump to the old cc
					// and then call  the handler
					(do (long-jump cc)
						(handler k v))))
*/

type savedContAsClosure Continuation

func (s savedContAsClosure) Exec(vm *VM) {
	val := vm.stack.Get(2)
	vm.stack = s.Stack
	vm.pc = s.Code
	vm.val = val
}

func instrForThrow(vm *VM) {
	cc := vm.stack.Get(0).(Continuation)
	v := vm.stack.Get(2)
	oldCC := vm.stack.Get(3).(Continuation)
	handler := vm.stack.Get(4).(*Closure)

	// Wrap the current stack into a closure.
	fn := &Closure{
		Required: 1,
		Code:     savedContAsClosure(cc),
	}

	// Long jump to the call site of 'try'.
	vm.stack = oldCC.Stack

	// New stack.
	vm.stack.base = vm.stack.pos

	// Call the handler function.
	vm.push(oldCC)
	vm.push(handler)
	vm.push(v)
	vm.push(fn)
	vm.pc = handler.Code
}
