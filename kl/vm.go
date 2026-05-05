package kl

import (
	"fmt"
	"unsafe"
)

// ---- Opcodes ----

const (
	OP_LOAD_CONST      = uint8(0)  // push consts[A]
	OP_LOAD_LOCAL      = uint8(1)  // push locals[A]
	OP_STORE_LOCAL     = uint8(2)  // locals[A] = pop()
	OP_LOAD_GLOBAL     = uint8(3)  // push mustSymbol(consts[A]).function
	OP_LOAD_UPVAL      = uint8(4)  // push upvals[A]
	OP_CALL            = uint8(5)  // non-tail call: fn at stack[top-A-1], A args above
	OP_TAIL_CALL       = uint8(6)  // tail call: same layout, signal trampoline
	OP_RETURN          = uint8(7)  // ctl.Return(pop())
	OP_JUMP            = uint8(8)  // pc += A (signed)
	OP_JUMP_FALSE      = uint8(9)  // pop(); if == False: pc += A
	OP_MAKE_CLOSURE    = uint8(10) // consts[A]=BytecodeFunc; pop B upvals; push closure
	OP_POP             = uint8(11) // discard top of stack
	OP_SELF_TAIL_CALL  = uint8(12) // A args on stack → update locals[0..A-1], jump pc=0
	// Arithmetic fast-paths (avoid trampoline + float64 boxing for fixnums)
	OP_ADD             = uint8(13) // pop y,x: push x+y
	OP_SUB             = uint8(14) // pop y,x: push x-y
	OP_MUL             = uint8(15) // pop y,x: push x*y
	OP_LT              = uint8(16) // pop y,x: push x<y
	OP_LE              = uint8(17) // pop y,x: push x<=y
	OP_GT              = uint8(18) // pop y,x: push x>y
	OP_GE              = uint8(19) // pop y,x: push x>=y
	OP_EQ              = uint8(20) // pop y,x: push x=y  (numeric equality only)
	OP_NOT             = uint8(21) // pop x: push (not x)
)

// Instr is a single VM instruction.  A and B are signed operands.
type Instr struct {
	Op uint8
	A  int32
	B  int32
}

// BytecodeFunc holds the compiled representation of a KL function.
type BytecodeFunc struct {
	Name    string
	Arity   int
	Nlocals int   // number of local slots (includes arity slots for params)
	Code    []Instr
	Consts  []Obj // constant pool (numbers, strings, symbols, nested BytecodeFunc objs)
}

// scmBytecodeFunc wraps a BytecodeFunc as an Obj so it can live in the Shen heap.
type scmBytecodeFunc struct {
	scmHead
	fn     *BytecodeFunc
	upvals []Obj // captured values for closures
}

const scmHeadBytecodeFunc scmHead = 50

func makeBytecodeObj(fn *BytecodeFunc, upvals []Obj) Obj {
	tmp := &scmBytecodeFunc{
		scmHead: scmHeadBytecodeFunc,
		fn:      fn,
		upvals:  upvals,
	}
	return &tmp.scmHead
}

func isBytecodeFunc(o Obj) bool {
	return *o == scmHeadBytecodeFunc
}

func mustBytecodeFunc(o Obj) *scmBytecodeFunc {
	return (*scmBytecodeFunc)(unsafe.Pointer(o))
}

// ---- Arithmetic helpers (fixnum fast paths) ----

func numAdd(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		return MakeInteger(fixnum(x) + fixnum(y))
	}
	return MakeNumber(mustNumber(x) + mustNumber(y))
}

func numSub(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		return MakeInteger(fixnum(x) - fixnum(y))
	}
	return MakeNumber(mustNumber(x) - mustNumber(y))
}

func numMul(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		return MakeInteger(fixnum(x) * fixnum(y))
	}
	return MakeNumber(mustNumber(x) * mustNumber(y))
}

// numCmp: returns True if sign(x - y) == wantSign (-1 for <, 0 for ==).
func numCmp(x, y Obj, wantSign int) Obj {
	if isFixnum(x) && isFixnum(y) {
		diff := fixnum(x) - fixnum(y)
		if (wantSign < 0 && diff < 0) {
			return True
		}
		return False
	}
	xf := mustNumber(x)
	yf := mustNumber(y)
	if (wantSign < 0 && xf < yf) {
		return True
	}
	return False
}

func numCmpLE(x, y Obj) Obj {
	if isFixnum(x) && isFixnum(y) {
		if fixnum(x) <= fixnum(y) {
			return True
		}
		return False
	}
	if mustNumber(x) <= mustNumber(y) {
		return True
	}
	return False
}

// vmApply is the entry point called from apply() for bytecode functions.
// It handles arity checking and partial application, then calls vmExec.
func vmApply(ctl *ControlFlow, bfObj Obj, args []Obj) {
	bf := mustBytecodeFunc(bfObj)
	fn := bf.fn
	provided := len(args)

	switch {
	case provided < fn.Arity:
		// Partial application: build a closure that waits for the remaining args.
		ctl.Return(vmPartialApply(fn.Arity, args, bfObj, bf.upvals))
	case provided == fn.Arity:
		vmExec(ctl, bf, args)
	case provided > fn.Arity:
		// Over-application: call with the right arity, then apply the result.
		res := Call(ctl, bfObj, args[:fn.Arity]...)
		ctl.TailApply(res, args[fn.Arity:]...)
	}
}

// vmPartialApply creates a closure that captures the supplied args and waits
// for the remaining (required - len(provided)) arguments.
func vmPartialApply(required int, providedArgs []Obj, proc Obj, upvals []Obj) Obj {
	symbols := makeTempSymbols(required)
	env1 := envExtend(Nil, symbols[:len(providedArgs)], providedArgs)

	args := Nil
	for i, count := len(symbols)-1, required-len(providedArgs); count > 0; count-- {
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

// vmExec runs one activation of a bytecode function.
// It either calls ctl.Return with the result or sets up a tail call.
func vmExec(ctl *ControlFlow, bf *scmBytecodeFunc, args []Obj) {
	fn := bf.fn
	upvals := bf.upvals

	locals := make([]Obj, fn.Nlocals)
	copy(locals, args)

	stack := make([]Obj, 0, 16)
	pc := 0
	code := fn.Code
	consts := fn.Consts

	for {
		instr := code[pc]
		pc++
		switch instr.Op {

		case OP_LOAD_CONST:
			stack = append(stack, consts[instr.A])

		case OP_LOAD_LOCAL:
			stack = append(stack, locals[instr.A])

		case OP_STORE_LOCAL:
			locals[instr.A] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

		case OP_LOAD_GLOBAL:
			sym := mustSymbol(consts[instr.A])
			if sym.function == nil {
				panic(MakeError(fmt.Sprintf("function %s not bound", sym.str)))
			}
			stack = append(stack, sym.function)

		case OP_LOAD_UPVAL:
			stack = append(stack, upvals[instr.A])

		case OP_CALL:
			n := int(instr.A)
			base := len(stack) - n - 1
			callee := stack[base]
			callArgs := make([]Obj, n)
			copy(callArgs, stack[base+1:])
			stack = stack[:base]
			result := Call(ctl, callee, callArgs...)
			stack = append(stack, result)

		case OP_TAIL_CALL:
			n := int(instr.A)
			base := len(stack) - n - 1
			callee := stack[base]
			callArgs := make([]Obj, n)
			copy(callArgs, stack[base+1:])
			ctl.TailApply(callee, callArgs...)
			return

		case OP_RETURN:
			ctl.Return(stack[len(stack)-1])
			return

		case OP_JUMP:
			pc += int(instr.A)

		case OP_JUMP_FALSE:
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v == False {
				pc += int(instr.A)
			}

		case OP_MAKE_CLOSURE:
			nUpvals := int(instr.B)
			innerBFObj := consts[instr.A]
			innerBF := mustBytecodeFunc(innerBFObj)
			captured := make([]Obj, nUpvals)
			base := len(stack) - nUpvals
			copy(captured, stack[base:])
			stack = stack[:base]
			closure := makeBytecodeObj(innerBF.fn, captured)
			stack = append(stack, closure)

		case OP_POP:
			stack = stack[:len(stack)-1]

		case OP_SELF_TAIL_CALL:
			n := int(instr.A)
			// All n args are on top of stack; copy into locals[0..n-1] and loop.
			copy(locals[:n], stack[len(stack)-n:])
			stack = stack[:len(stack)-n]
			pc = 0

		case OP_ADD:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, numAdd(x, y))

		case OP_SUB:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, numSub(x, y))

		case OP_MUL:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, numMul(x, y))

		case OP_LT:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, numCmp(x, y, -1))

		case OP_LE:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, numCmpLE(x, y))

		case OP_GT:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, numCmp(y, x, -1)) // x > y ↔ y < x

		case OP_GE:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, numCmpLE(y, x)) // x >= y ↔ y <= x

		case OP_EQ:
			y := stack[len(stack)-1]
			x := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, equal(x, y))

		case OP_NOT:
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if x == True {
				stack = append(stack, False)
			} else if x == False {
				stack = append(stack, True)
			} else {
				panic(MakeError("not: expected boolean"))
			}

		default:
			panic(fmt.Sprintf("vmExec: unknown opcode %d at pc %d", instr.Op, pc-1))
		}
	}
}
