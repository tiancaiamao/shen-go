package kl

import (
	"fmt"
	"unsafe"
)

// ---- Opcodes ----

const (
	OP_LOAD_CONST   = uint8(0)  // push consts[A]
	OP_LOAD_LOCAL   = uint8(1)  // push locals[A]
	OP_STORE_LOCAL  = uint8(2)  // locals[A] = pop(); does NOT pop from value stack
	OP_LOAD_GLOBAL  = uint8(3)  // push mustSymbol(consts[A]).function
	OP_LOAD_UPVAL   = uint8(4)  // push upvals[A]
	OP_CALL         = uint8(5)  // non-tail call: fn at stack[top-A-1], A args above
	OP_TAIL_CALL    = uint8(6)  // tail call: same layout, signal trampoline
	OP_RETURN       = uint8(7)  // ctl.Return(pop())
	OP_JUMP         = uint8(8)  // pc += A (signed, A is int32)
	OP_JUMP_FALSE   = uint8(9)  // pop(); if == False: pc += A
	OP_MAKE_CLOSURE = uint8(10) // consts[A] is *BytecodeFunc; pop B upvals; push closure
	OP_POP          = uint8(11) // discard top of stack
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

		default:
			panic(fmt.Sprintf("vmExec: unknown opcode %d at pc %d", instr.Op, pc-1))
		}
	}
}
