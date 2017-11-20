package vm

import (
	"fmt"
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
}

// Code is something executable to VM, it's immutable.
type Code struct {
	bc     []instruction
	consts []kl.Obj
}

type address struct {
	pc   int
	code *Code
}

type Procedure struct {
	kl.ScmRaw
	code *Code
	env  []kl.Obj
}

func NewVM() *VM {
	vm := &VM{
		stack:         make([]kl.Obj, 200),
		env:           make([]kl.Obj, 0, 200),
		functionTable: make(map[string]*Procedure),
	}
	vm.arg.init(100)
	return vm
}

func (vm *VM) Run(code *Code) {
	vm.code = code
	vm.pc = 0

	halt := false
	for !halt {
		inst := vm.code.bc[vm.pc]
		vm.pc++

		switch instructionCode(inst) {
		case iConst:
			n := instructionOP1(inst)
			fmt.Println("CONST ", n, kl.ObjString(vm.code.consts[n]), vm.top)
			vm.stack[vm.top] = vm.code.consts[n]
			vm.top++
		case iAccess:
			n := instructionOP1(inst)
			// get value from environment
			vm.stack[vm.top] = vm.env[len(vm.env)-1-n]
			fmt.Println("ACCESS", n, " get ", kl.ObjString(vm.stack[vm.top]))
			vm.top++
		case iGrab:
			if vm.arg.empty() {
				// make closure if there are not enough arguments
				n := instructionOP1(inst)
				fmt.Println("GRAB, not enough argument, make a closure", vm.pc, n)
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
				vm.stack[vm.top] = tmp.ScmRaw.Object()
				vm.top++
				vm.pc += n
			} else {
				// grab data from stack to env
				v := vm.arg.pop()
				fmt.Println("GRAB, pop a value", kl.ObjString(v))
				vm.env = append(vm.env, v)
			}
		case iReturn:
			if vm.arg.empty() {
				savedAddr := vm.savedAddr[len(vm.savedAddr)-1]
				vm.savedAddr = vm.savedAddr[:len(vm.savedAddr)-1]

				vm.code = savedAddr.code
				vm.pc = savedAddr.pc
				fmt.Println("RETURN ", vm.top, vm.arg.count(), len(vm.code.bc), vm.pc)
			} else {
				// more arguments then necessary, continue the beta-reduce.
				fmt.Println("RETURN: TODO, should partial apply")
			}
		case iPop:
			fmt.Println("POP")
			vm.top--
		case iDefun:
			symbol := kl.SymbolString(vm.stack[vm.top-1])
			fmt.Println("DEFUN", symbol)
			function := (*Procedure)(unsafe.Pointer(vm.stack[vm.top-2]))
			vm.functionTable[symbol] = function
			vm.top--
			vm.stack[vm.top-1] = vm.stack[vm.top]
		case iGetF:
			symbol := kl.SymbolString(vm.stack[vm.top-1])
			if function, ok := vm.functionTable[symbol]; ok {
				vm.stack[vm.top-1] = function.ScmRaw.Object()
			} else {
				vm.stack[vm.top-1] = kl.Make_error("unknown function:" + symbol)
			}
		case iHalt:
			fmt.Println("HALT", vm.top, vm.arg.count(), len(vm.savedAddr))
			halt = true
		case iPushArg:
			// pop a value from stack top to argument queue
			fmt.Println("PUSHARG from", vm.top)
			vm.top--
			vm.arg.push(vm.stack[vm.top])
		case iApply:
			vm.top--
			obj := vm.stack[vm.top]
			closure := (*Procedure)(unsafe.Pointer(obj))
			fmt.Println("APPLY", vm.top, len(closure.code.bc), "save pc=", vm.pc)
			// save return address
			// set pc to closure code
			// prepare initialize environment from closure
			vm.savedAddr = append(vm.savedAddr, address{vm.pc, vm.code})
			vm.code = closure.code
			vm.pc = 0
			vm.env = vm.env[0:]
			vm.env = append(vm.env, closure.env...)
		case iPrimCall:
			fallthrough
		case iAdd:
			vm.top--
			o1 := vm.stack[vm.top]
			o2 := vm.stack[vm.top-1]
			o := kl.PrimNumberAdd(o1, o2)
			vm.stack[vm.top-1] = o
			fmt.Println("ADD result = ", kl.ObjString(o))
		default:
			panic(fmt.Sprintf("unknown instruction %d", inst))
		}
	}
}

func (vm *VM) debug() {
	fmt.Println("top:", vm.top)
	fmt.Println("arg:", vm.arg.count())
	fmt.Println("result:", kl.ObjString(vm.stack[vm.top-1]))
}
