package vm

import (
	"fmt"
	"unsafe"

	"github.com/tiancaiamao/shen-go/kl"
)

type VM struct {
	stack []kl.Obj
	top   int // stack top

	mark      int // stack[mark:top] are reserved for arguments
	savedMark []int

	env []kl.Obj // environment stack

	pc        int // pc register refer to the position in current code
	code      *Code
	savedAddr []address // saved return address
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
	return &VM{
		stack: make([]kl.Obj, 200),
		env:   make([]kl.Obj, 0, 200),
	}
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
			fmt.Println("GRAB", vm.top, vm.mark)
			if vm.top > vm.mark {
				// grab data from stack to env
				vm.top--
				vm.env = append(vm.env, vm.stack[vm.top])
			} else {
				// make closure if there are not enough arguments
			}
		case iReturn:
			if vm.top-1 == vm.mark {
				savedMark := vm.savedMark[len(vm.savedMark)-1]
				vm.savedMark = vm.savedMark[:len(vm.savedMark)-1]

				savedAddr := vm.savedAddr[len(vm.savedAddr)-1]
				vm.savedAddr = vm.savedAddr[:len(vm.savedAddr)-1]

				vm.mark = savedMark
				vm.code = savedAddr.code
				vm.pc = savedAddr.pc
				fmt.Println("RETURN ", vm.top, vm.mark, len(vm.code.bc), vm.pc)
			} else {
				// more arguments then necessary, continue the beta-reduce.
			}
		case iHalt:
			fmt.Println("HALT", vm.top, vm.mark, len(vm.savedMark), len(vm.savedAddr))
			halt = true
		case iPush:
		case iPop:
		case iPushArg:
			n := instructionOP1(inst)
			// pop the stack top value, assign to stack[mark + N]
			vm.top--
			fmt.Println("PUSHARG from", vm.top, "=>", vm.mark+n)
			vm.stack[vm.mark+n] = vm.stack[vm.top]
		case iMark:
			n := instructionOP1(inst)
			fmt.Println("MARK ", n, vm.top, vm.mark)
			// save previous mark position
			// set stack top as current mark position
			// push stack top to reserve space for arguments
			vm.savedMark = append(vm.savedMark, vm.mark)
			vm.mark = vm.top
			vm.top += n
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
