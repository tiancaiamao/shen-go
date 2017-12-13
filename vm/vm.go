package vm

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"unsafe"

	"github.com/tiancaiamao/shen-go/kl"
)

type VM struct {
	stack []kl.Obj
	top   int // stack top

	env []kl.Obj // environment

	pc        int       // pc register refer to the position in current code
	savedAddr []address // saved return address

	functionTable map[string]*Procedure
	symbolTable   map[string]kl.Obj

	// jumpBuf is used to implement exception, similar to setjmp/longjmp in C.
	cc *jumpBuf
}

type jumpBuf struct {
	address
	savedAddrPos int
	closure      kl.Obj
}

// Code is something executable to VM, it's immutable.
type Code struct {
	bc     []instruction
	consts []kl.Obj
}

// address is the information to be saved before apply a closure.
type address struct {
	pc   int
	code *Code
	env  []kl.Obj
}

type Procedure struct {
	scmHead int
	code    *Code
	env     []kl.Obj
}

const initStackSize = 128

var auxVM *VM = New()
var stackMarkDummyValue int
var stackMark = kl.MakeRaw(&stackMarkDummyValue)

func New() *VM {
	vm := &VM{
		stack:         make([]kl.Obj, initStackSize),
		env:           make([]kl.Obj, 0, 200),
		functionTable: make(map[string]*Procedure),
		symbolTable:   make(map[string]kl.Obj),
	}
	initSymbolTable(vm.symbolTable)
	return vm
}

func initSymbolTable(symbolTable map[string]kl.Obj) {
	dir, _ := os.Getwd()
	symbolTable["*stinput*"] = kl.MakeStream(os.Stdin)
	symbolTable["*stoutput*"] = kl.MakeStream(os.Stdout)
	symbolTable["*home-directory*"] = kl.MakeString(dir)
	symbolTable["*language*"] = kl.MakeString("Go")
	symbolTable["*implementation*"] = kl.MakeString("bytecode")
	symbolTable["*relase*"] = kl.MakeString(runtime.Version())
	symbolTable["*os*"] = kl.MakeString(runtime.GOOS)
	symbolTable["*porters*"] = kl.MakeString("Arthur Mao")
	symbolTable["*port*"] = kl.MakeString("0.0.1")

	// Extended by shen-go implementation
	symbolTable["*package-path*"] = kl.MakeString(kl.PackagePath())
}

func (vm *VM) Run(code *Code) kl.Obj {
	vm.pc = 0

	// From the example:
	// case one: 1 	[[iConst 1] [iHalt]]
	// case two: ((lambda X X) 1) [[iConst 1] [iGrab [iAccess 0] [iReturn]] [iTailApply 1] [iHalt]]
	//
	// We see that case one doesn't need the initial savedAddr, while case does.
	// If savedAddr begins with empty, case two would panic. If savedAddr begins with 1,
	// run case two would accumulate more and more savedAddrs.
	// So reset savedAddr to length 1 and both case one and case two feel happy.
	vm.savedAddr = vm.savedAddr[:0]
	vm.savedAddr = append(vm.savedAddr, address{pc: len(code.bc) - 1, code: code})
	vm.stack[0] = stackMark
	vm.top = 1

	halt := false
	for !halt {
		exception := false
		inst := code.bc[vm.pc]
		vm.pc++

		switch instructionCode(inst) {
		case iSetJmp:
			n := instructionOPN(inst)
			vm.top--
			vm.cc = &jumpBuf{
				address: address{
					pc:   vm.pc + n,
					code: code,
					env:  vm.env,
				},
				savedAddrPos: len(vm.savedAddr),
				closure:      vm.stack[vm.top],
			}
			fmt.Fprintln(StdBC, "SETJMP", vm.cc.savedAddrPos, vm.cc.address.pc)
		case iClearJmp:
			fmt.Fprintln(StdBC, "CLEARJMP")
			vm.cc = nil
		case iConst:
			n := instructionOPN(inst)
			fmt.Fprintln(StdBC, "CONST ", n, kl.ObjString(code.consts[n]), vm.top)
			vm.stackPush(code.consts[n])
		case iAccess:
			n := instructionOPN(inst)
			// get value from environment
			vm.stackPush(vm.env[len(vm.env)-1-n])
			fmt.Fprintln(StdBC, "ACCESS", n, " get ", kl.ObjString(vm.stack[vm.top-1]))
		case iFreeze:
			// create closure directly
			// nearly the same with grab, but if need zero arguments.
			n := instructionOPN(inst)
			fmt.Fprintln(StdBC, "FREEZE", vm.pc, n)
			tmp := &Procedure{
				code: &Code{
					bc:     code.bc[vm.pc:],
					consts: code.consts,
				},
			}
			raw := kl.MakeRaw(&tmp.scmHead)
			if len(vm.env) > 0 {
				tmp.env = make([]kl.Obj, len(vm.env))
				copy(tmp.env, vm.env)
			}
			vm.stackPush(raw)
			vm.pc += n
		case iMark:
			fmt.Fprintln(StdBC, "MARK")
			vm.stackPush(stackMark)
		case iGrab:
			vm.top--
			if v := vm.stack[vm.top]; v == stackMark {
				// make closure if there are not enough arguments
				n := instructionOPN(inst)
				fmt.Fprintln(StdBC, "GRAB, not enough argument, make a closure", vm.pc, n)
				tmp := Procedure{
					code: &Code{
						bc:     code.bc[vm.pc-1:],
						consts: code.consts,
					},
				}
				raw := kl.MakeRaw(&tmp.scmHead)
				tmp.env = vm.env
				vm.stackPush(raw)

				// return to saved address
				savedAddr := vm.savedAddr[len(vm.savedAddr)-1]
				vm.savedAddr = vm.savedAddr[:len(vm.savedAddr)-1]
				code = savedAddr.code
				vm.pc = savedAddr.pc
				vm.env = savedAddr.env
			} else {
				// grab data from stack to env
				fmt.Fprintln(StdBC, "GRAB, pop a value", kl.ObjString(v), len(vm.savedAddr))
				vm.env = append(vm.env, v)
			}
		case iReturn:
			// stack[top-1] is the result, so should check top-2
			if vm.stack[vm.top-2] == stackMark {
				savedAddr := vm.savedAddr[len(vm.savedAddr)-1]
				vm.savedAddr = vm.savedAddr[:len(vm.savedAddr)-1]

				code = savedAddr.code
				vm.pc = savedAddr.pc
				vm.env = savedAddr.env
				vm.top--
				vm.stack[vm.top-1] = vm.stack[vm.top]
				fmt.Fprintln(StdBC, "RETURN ", len(vm.savedAddr), vm.top, vm.pc)
			} else {
				// more arguments, continue the beta-reduce.
				// similar to tail apply
				fmt.Fprintln(StdBC, "RETURN more arguments to be consumed!", len(vm.savedAddr))
				vm.top--
				obj := vm.stack[vm.top]
				// TODO: panic if obj is not a closure
				closure := (*Procedure)(unsafe.Pointer(obj))
				code = closure.code
				vm.pc = 0
				vm.env = closure.env
			}
		case iTailApply:
			vm.top--
			obj := vm.stack[vm.top]
			closure := (*Procedure)(unsafe.Pointer(obj))
			// The only different with Apply is that TailApply doesn't save return address.
			code = closure.code
			fmt.Fprintln(StdBC, "TAILAPPLY", vm.top, len(code.bc), closure)
			vm.pc = 0
			vm.env = closure.env
		case iApply:
			vm.top--
			obj := vm.stack[vm.top]
			closure := (*Procedure)(unsafe.Pointer(obj))
			// save return address
			vm.savedAddr = append(vm.savedAddr, address{vm.pc, code, vm.env})
			fmt.Fprintln(StdBC, "APPLY", vm.top, len(vm.savedAddr), len(code.bc))
			// set pc to closure code
			code = closure.code
			vm.pc = 0
			vm.env = closure.env
		case iPop:
			fmt.Fprintln(StdBC, "POP")
			vm.top--
		case iDefun:
			symbol := kl.GetSymbol(vm.stack[vm.top-1])
			function := (*Procedure)(unsafe.Pointer(vm.stack[vm.top-2]))
			fmt.Fprintln(StdBC, "DEFUN", symbol, len(function.code.bc), function)
			vm.functionTable[symbol] = function
			vm.top--
			vm.stack[vm.top-1] = vm.stack[vm.top]
		case iGetF:
			symbol := kl.GetSymbol(vm.stack[vm.top-1])
			fmt.Fprintln(StdBC, "GETF", symbol)
			if function, ok := vm.functionTable[symbol]; ok {
				vm.stack[vm.top-1] = kl.MakeRaw(&function.scmHead)
			} else {
				vm.stack[vm.top-1] = kl.MakeError("unknown function:" + symbol)
			}
		case iJF:
			fmt.Fprintln(StdBC, "JF")
			switch vm.stack[vm.top-1] {
			case kl.False:
				n := instructionOPN(inst)
				vm.top--
				vm.pc += n
				fmt.Fprintln(StdDebug, "JF false branch pc=", vm.pc, n)
			case kl.True:
				fmt.Fprintln(StdDebug, "JF true branch pc=", vm.pc)
				vm.top--
			default:
				// TODO: So what?
				vm.stack[vm.top-1] = kl.MakeError("test condition need to be boolean")
			}
		case iJMP:
			n := instructionOPN(inst)
			vm.pc += n
			fmt.Fprintln(StdBC, "JMP", n, vm.pc)
		case iHalt:
			fmt.Fprintln(StdBC, "HALT", vm.top, len(vm.savedAddr))
			halt = true
		case iPrimCall:
			id := instructionOPN(inst)
			prim := kl.GetPrimitiveByID(id)
			args := vm.stack[vm.top-prim.Required : vm.top]

			var result kl.Obj
			// Ugly hack: set function should not be global.
			switch prim.Name {
			case "set":
				result = kl.PrimSet(vm.symbolTable, args[0], args[1])
			case "value":
				result = kl.PrimValue(vm.symbolTable, args[0])
			case "eval-kl":
				auxVM.symbolTable = vm.symbolTable
				auxVM.functionTable = vm.functionTable
				result = auxVM.Eval(args[0])
			default:
				result = prim.Function(args...)
			}

			fmt.Fprintln(StdBC, "PRIMCALL", prim.Name)
			vm.stack[vm.top-prim.Required] = result
			vm.top = vm.top - prim.Required + 1
			if kl.IsError(result) {
				exception = true
			}
		default:
			panic(fmt.Sprintf("unknown instruction %d", inst))
		}

		if exception {
			fmt.Fprintln(StdBC, "exception")
			vm.Debug()
			if vm.cc == nil {
				err := vm.stack[vm.top-1]
				vm.Reset()
				return err
			}

			// clear jmpBuf
			jmpBuf := vm.cc
			vm.cc = nil
			// pop trap-error handler, prepare for call.
			vm.stackPush(vm.stack[vm.top-1])
			vm.stack[vm.top-2] = stackMark
			// recover savedAddr
			vm.savedAddr = vm.savedAddr[:jmpBuf.savedAddrPos]
			vm.savedAddr = append(vm.savedAddr, jmpBuf.address)
			fmt.Fprintln(StdDebug, "exception", len(vm.savedAddr))
			// longjmp... tail apply
			closure := (*Procedure)(unsafe.Pointer(jmpBuf.closure))
			code = closure.code
			fmt.Fprintf(StdDebug, "len code = %d\n", len(code.bc))
			fmt.Fprintf(StdDebug, "address = %#v\n", jmpBuf.address)
			vm.pc = 0
			vm.env = closure.env
		}
	}

	if vm.top != 1 {
		vm.Debug()
		panic("vm in wrong status")
	}
	vm.top--
	return vm.stack[vm.top]
}

func (vm *VM) stackPush(o kl.Obj) {
	if vm.top == len(vm.stack) {
		stack := make([]kl.Obj, len(vm.stack)*2)
		copy(stack, vm.stack)
		vm.stack = stack
	}
	vm.stack[vm.top] = o
	vm.top++
}

func (vm *VM) Reset() {
	vm.stack = vm.stack[:initStackSize]
	vm.top = 0
	vm.env = vm.env[:0]
	vm.savedAddr = vm.savedAddr[:0]
	vm.cc = nil
}

func (vm *VM) Debug() {
	fmt.Fprintln(StdDebug, "pc:", vm.pc)
	fmt.Fprintln(StdDebug, "top:", vm.top)
	fmt.Fprintln(StdDebug, "stack:")
	for i := vm.top - 1; i >= 0; i-- {
		if vm.stack[i] == stackMark {
			fmt.Fprintln(StdDebug, "MARK")
		} else {
			fmt.Fprintln(StdDebug, kl.ObjString(vm.stack[i]))
		}
	}
	fmt.Fprintln(StdDebug, "function:", len(vm.functionTable))
	fmt.Fprintln(StdDebug)
}

var evaluator = kl.NewEvaluator()

func klToSexpByteCode(klambda kl.Obj) kl.Obj {
	fmt.Fprintln(StdDebug, "kl->bytecode:", kl.ObjString(klambda))
	quote := kl.Cons(kl.MakeSymbol("quote"), kl.Cons(klambda, kl.Nil))
	f := kl.Cons(kl.MakeSymbol("kl->bytecode"), kl.Cons(quote, kl.Nil))
	// Get the bytecode in sexp representation.
	return evaluator.Eval(f)
}

func klToByteCode(klambda kl.Obj) (*Code, error) {
	bc := klToSexpByteCode(klambda)
	if bc == kl.Nil {
		return nil, errors.New("klToByteCode return some thing wrong")
	}
	fmt.Fprintln(StdDebug, "bytecode in sexp:", kl.ObjString(bc))
	var a Assember
	err := a.FromSexp(bc)
	if err != nil {
		return nil, err
	}
	code := a.Comiple()
	fmt.Fprintln(StdDebug, a.Decode(code))
	return code, nil
}

func (vm *VM) Eval(sexp kl.Obj) kl.Obj {
	code, err := klToByteCode(sexp)
	if err != nil {
		return kl.MakeError(err.Error())
	}

	return vm.Run(code)
}

func BootstrapCompiler() {
	evaluator.Silence = true
	mustLoadKLFile("ShenOSKernel-20.1/klambda/toplevel.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/core.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/sys.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/sequent.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/yacc.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/reader.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/prolog.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/track.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/load.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/writer.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/macros.kl")
	mustLoadKLFile("ShenOSKernel-20.1/klambda/declarations.kl")
	mustLoadKLFile("cmd/shen/primitive.kl")
	mustLoadKLFile("cmd/shen/de-bruijn.kl")
	mustLoadKLFile("cmd/shen/compile.kl")
}

func mustLoadKLFile(file string) {
	filePath := path.Join(kl.PackagePath(), file)
	o := evaluator.LoadFile(filePath)
	if kl.IsError(o) {
		panic(kl.ObjString(o))
	}
}
