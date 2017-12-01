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

	env []kl.Obj // environment stack
	arg queue    // argument queue

	pc        int // pc register refer to the position in current code
	code      *Code
	savedAddr []address // saved return address

	functionTable map[string]*Procedure
	symbolTable   map[string]kl.Obj

	// snapshot is a vm snapshot, used to implement exception.
	snapshot *VM
}

// Code is something executable to VM, it's immutable.
type Code struct {
	bc     []instruction
	consts []kl.Obj
}

type address struct {
	pc   int
	code *Code
	env  []kl.Obj
}

type Procedure struct {
	kl.ScmRaw
	code *Code
	env  []kl.Obj
}

const initStackSize = 128

var auxVM *VM = New()

func New() *VM {
	vm := &VM{
		stack:         make([]kl.Obj, initStackSize),
		env:           make([]kl.Obj, 0, 200),
		functionTable: make(map[string]*Procedure),
		symbolTable:   make(map[string]kl.Obj),
	}
	initSymbolTable(vm.symbolTable)
	vm.arg.init(100)
	return vm
}

func initSymbolTable(symbolTable map[string]kl.Obj) {
	dir, _ := os.Getwd()
	symbolTable["*stinput*"] = kl.Make_stream(os.Stdin)
	symbolTable["*stoutput*"] = kl.Make_stream(os.Stdout)
	symbolTable["*home-directory*"] = kl.Make_string(dir)
	symbolTable["*language*"] = kl.Make_string("Go")
	symbolTable["*implementation*"] = kl.Make_string("bytecode")
	symbolTable["*relase*"] = kl.Make_string(runtime.Version())
	symbolTable["*os*"] = kl.Make_string(runtime.GOOS)
	symbolTable["*porters*"] = kl.Make_string("Arthur Mao")
	symbolTable["*port*"] = kl.Make_string("0.0.1")

	// Extended by shen-go implementation
	gopath := os.Getenv("GOPATH")
	packagePath := path.Join(gopath, "src/github.com/tiancaiamao/shen-go/pkg")
	symbolTable["*package-path*"] = kl.Make_string(packagePath)
}

func (vm *VM) takeSnapshot(pc int) {
	var snapshot VM
	snapshot = *vm
	snapshot.arg.clone(&vm.arg)
	snapshot.stack = make([]kl.Obj, len(vm.stack), cap(vm.stack))
	copy(snapshot.stack, vm.stack)
	snapshot.env = make([]kl.Obj, len(vm.env), cap(vm.env))
	copy(snapshot.env, vm.env)
	snapshot.savedAddr = make([]address, len(vm.savedAddr), cap(vm.savedAddr))
	copy(snapshot.savedAddr, vm.savedAddr)
	snapshot.snapshot = nil
	snapshot.pc = pc
	fmt.Fprintln(StdDebug, "code len:", len(snapshot.code.bc), "pc:", pc)

	vm.snapshot = &snapshot
}

func (vm *VM) Run(code *Code) (kl.Obj, error) {
	vm.code = code
	vm.pc = 0

	vm.savedAddr = append(vm.savedAddr, address{pc: len(code.bc) - 1, code: code})
	halt := false
	for !halt {
		exception := false
		inst := vm.code.bc[vm.pc]
		vm.pc++

		switch instructionCode(inst) {
		case iSetJmp:
			n := instructionOP1(inst)
			fmt.Fprintln(StdBC, "SETJMP", vm.pc+n)
			vm.takeSnapshot(vm.pc + n)
		case iConst:
			n := instructionOP1(inst)
			fmt.Fprintln(StdBC, "CONST ", n, kl.ObjString(vm.code.consts[n]), vm.top)
			vm.stackPush(vm.code.consts[n])
		case iAccess:
			n := instructionOP1(inst)
			// get value from environment
			vm.stackPush(vm.env[len(vm.env)-1-n])
			fmt.Fprintln(StdBC, "ACCESS", n, " get ", kl.ObjString(vm.stack[vm.top-1]))
		case iFreeze:
			// create closure directly
			// nearly the same with grab, but if need zero arguments.
			n := instructionOP1(inst)
			fmt.Fprintln(StdBC, "FREEZE", vm.pc, n)
			tmp := Procedure{
				ScmRaw: kl.Make_raw(),
				code: &Code{
					bc:     vm.code.bc[vm.pc:],
					consts: vm.code.consts,
				},
			}
			if len(vm.env) > 0 {
				tmp.env = make([]kl.Obj, len(vm.env))
				copy(tmp.env, vm.env)
			}
			vm.stackPush(tmp.ScmRaw.Object())
			vm.pc += n
		case iGrab:
			if vm.arg.empty() {
				// make closure if there are not enough arguments
				n := instructionOP1(inst)
				fmt.Fprintln(StdBC, "GRAB, not enough argument, make a closure", vm.pc, n)
				tmp := Procedure{
					ScmRaw: kl.Make_raw(),
					code: &Code{
						bc:     vm.code.bc[vm.pc-1:],
						consts: vm.code.consts,
					},
				}
				if len(vm.env) > 0 {
					tmp.env = make([]kl.Obj, len(vm.env))
					copy(tmp.env, vm.env)
				}
				vm.stackPush(tmp.ScmRaw.Object())
				vm.pc += n
			} else {
				// grab data from stack to env
				v := vm.arg.pop()
				fmt.Fprintln(StdBC, "GRAB, pop a value", kl.ObjString(v))
				vm.env = append(vm.env, v)
			}
		case iReturn:
			if vm.arg.empty() {
				savedAddr := vm.savedAddr[len(vm.savedAddr)-1]
				vm.savedAddr = vm.savedAddr[:len(vm.savedAddr)-1]

				vm.code = savedAddr.code
				vm.pc = savedAddr.pc
				vm.env = savedAddr.env
				fmt.Fprintln(StdBC, "RETURN ", vm.top, vm.arg.count(), len(vm.code.bc), vm.pc)
			} else {
				// more arguments then necessary, continue the beta-reduce.
				// equivalent to tail apply
				fmt.Fprintln(StdBC, "RETURN: TODO, should partial apply")
			}
		case iPop:
			fmt.Fprintln(StdBC, "POP")
			vm.top--
		case iDefun:
			symbol := kl.SymbolString(vm.stack[vm.top-1])
			fmt.Fprintln(StdBC, "DEFUN", symbol)
			function := (*Procedure)(unsafe.Pointer(vm.stack[vm.top-2]))
			vm.functionTable[symbol] = function
			vm.top--
			vm.stack[vm.top-1] = vm.stack[vm.top]
		case iGetF:
			symbol := kl.SymbolString(vm.stack[vm.top-1])
			fmt.Fprintln(StdBC, "GETF", symbol)
			if function, ok := vm.functionTable[symbol]; ok {
				vm.stack[vm.top-1] = function.ScmRaw.Object()
			} else {
				vm.stack[vm.top-1] = kl.Make_error("unknown function:" + symbol)
			}
		case iJF:
			fmt.Fprintln(StdBC, "JF")
			switch vm.stack[vm.top-1] {
			case kl.False:
				n := instructionOP1(inst)
				vm.top--
				vm.pc += n
				fmt.Fprintln(StdDebug, "JF false branch pc=", vm.pc, n)
			case kl.True:
				fmt.Fprintln(StdDebug, "JF true branch pc=", vm.pc)
				vm.top--
			default:
				// TODO: So what?
				vm.stack[vm.top-1] = kl.Make_error("test condition need to be boolean")
			}
		case iJMP:
			n := instructionOP1(inst)
			vm.pc += n
			fmt.Fprintln(StdBC, "JMP", n, vm.pc)
		case iSwap:
			fmt.Fprintln(StdBC, "SWAP")
			vm.stack[vm.top-1], vm.stack[vm.top-2] = vm.stack[vm.top-2], vm.stack[vm.top-1]
		case iHalt:
			fmt.Fprintln(StdBC, "HALT", vm.top, vm.arg.count(), len(vm.savedAddr))
			halt = true
		case iPushArg:
			// pop a value from stack top to argument queue
			fmt.Fprintln(StdBC, "PUSHARG from", vm.top)
			vm.top--
			vm.arg.push(vm.stack[vm.top])
		case iApply:
			vm.top--
			obj := vm.stack[vm.top]
			closure := (*Procedure)(unsafe.Pointer(obj))
			fmt.Fprintln(StdBC, "APPLY", vm.top, len(closure.code.bc), "save pc=", vm.pc)
			// save return address
			// set pc to closure code
			// prepare initialize environment from closure
			vm.savedAddr = append(vm.savedAddr, address{vm.pc, vm.code, vm.env})
			vm.code = closure.code
			vm.pc = 0
			vm.env = closure.env
		case iTailApply:
			vm.top--
			obj := vm.stack[vm.top]
			closure := (*Procedure)(unsafe.Pointer(obj))
			fmt.Fprintln(StdBC, "TAILAPPLY", vm.top, len(closure.code.bc), "save pc=", vm.pc)
			// The only different with Apply is that TailApply doesn't save return address.
			vm.code = closure.code
			vm.pc = 0
			vm.env = vm.env[0:]
			vm.env = append(vm.env, closure.env...)
		case iPrimCall:
			id := instructionOP1(inst)
			prim := kl.Primitives[id]
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

			fmt.Fprintln(StdBC, "PRIMCALL", prim.Name, kl.ObjString(result))
			vm.stack[vm.top-prim.Required] = result
			vm.top = vm.top - prim.Required + 1
			if kl.IsError(result) {
				exception = true
			}
		default:
			panic(fmt.Sprintf("unknown instruction %d", inst))
		}

		if exception {
			fmt.Fprintln(StdDebug, "exception")
			vm.Debug()
			if vm.snapshot == nil {
				vm.Reset()
				break
			}
			saveErr := vm.stack[vm.top-1]
			*vm = *vm.snapshot
			vm.stackPush(saveErr)
		}
	}

	if vm.top != 1 {
		return kl.Nil, errors.New("vm in wrong status")
	}
	vm.top--
	return vm.stack[vm.top], nil
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
	vm.arg.reset()
	vm.stack = vm.stack[:initStackSize]
	vm.top = 0
	vm.env = vm.env[:0]
	vm.savedAddr = vm.savedAddr[:0]
}

func (vm *VM) Debug() {
	fmt.Fprintln(StdDebug, "pc:", vm.pc)
	fmt.Fprintln(StdDebug, "arg:", vm.arg.count())
	fmt.Fprintln(StdDebug, "top:", vm.top)
	fmt.Fprintln(StdDebug, "stack:")
	for i := vm.top - 1; i >= 0; i-- {
		fmt.Fprintln(StdDebug, kl.ObjString(vm.stack[i]))
	}
	fmt.Fprintln(StdDebug, "function:", len(vm.functionTable))
	fmt.Fprintln(StdDebug)
}

func klToSexpByteCode(klambda kl.Obj) kl.Obj {
	fmt.Fprintln(StdDebug, "kl->bytecode:", kl.ObjString(klambda))
	quote := kl.Cons(kl.Make_symbol("quote"), kl.Cons(klambda, kl.Nil))
	f := kl.Cons(kl.Make_symbol("kl->bytecode"), kl.Cons(quote, kl.Nil))
	// Get the bytecode in sexp representation.
	return kl.Eval(f)
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
		return kl.Make_error(err.Error())
	}

	res, err := vm.Run(code)
	if err != nil {
		vm.Reset()
		return kl.Make_error(err.Error())
	}
	return res
}

func BootstrapCompiler() {
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
	filePath := path.Join(kl.PackagePath, file)
	o := kl.LoadFile(filePath)
	if kl.IsError(o) {
		panic(kl.ObjString(o))
	}
}
