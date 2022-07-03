package re

import (
	"fmt"
)

func instrGenC(i Instr) {
	gen, ok := i.(interface{ GenC() })
	if !ok {
		fmt.Println("unknown instruct")
		panic(i)
	}
	gen.GenC()
}

func (c InstrPush) GenC() {
	fmt.Printf("push(vm, vm.val);\n")
	instrGenC(c.Next)
}

func (c InstrConst) GenC() {
	// fmt.Println("// InstrConst")
	switch raw := c.Val.(type) {
	case nilObj:
		fmt.Printf("vm.val = Nil;\n")
	case booleanObj:
		if c.Val == True {
			fmt.Printf("vm.val = True;\n")
		} else if c.Val == False {
			fmt.Printf("vm.val = False;\n")
		}
	case Integer:
		fmt.Printf("vm.val = makeNumber(%d);\n", raw)
	case Float64:
		fmt.Printf("TODO InstrConst\n")
	case String:
		fmt.Printf("vm.val = makeString(\"%s\"));\n", raw)
	case *Symbol:
		fmt.Printf("vm.val = makeSymbol(\"%s\"));\n", raw)
	default:
		fmt.Printf("InstrConst unsupport type\n")
	}
	instrGenC(c.Next)
}

func (i InstrNop) GenC() {
	// nop
	fmt.Printf("vm.val = NULL; // nop\n")
	instrGenC(i.Next)
}

func (i InstrLocalRef) GenC() {
	// fmt.Println("// InstrLocalRef")
	fmt.Printf("vm.val = vmGet(vm, %d);\n", i.Idx+2)
	instrGenC(i.Next)
}

func (i InstrGlobalRef) GenC() {
	// fmt.Println("// InstrGlobalRef")
	fmt.Printf("vm.val = symbolFn(%s);\n", i.sym)
	instrGenC(i.Next)
}

func (i InstrIf) GenC() {
	// fmt.Println("// InstrIf")
	fmt.Printf("if (vm.val == True) {\n")

	instrGenC(i.Then)

	fmt.Printf("} else if (vm.val == False) {\n")

	instrGenC(i.Else)

	fmt.Printf("} else {\n")
	fmt.Printf("perror(\"if only accept true or false\");\n")
	fmt.Printf("}\n")
}

func (i InstrPrepareCall) GenC() {
	// fmt.Println("// InstrPrepareCall")
	fmt.Printf("push(vm, Nil);\n")
	instrGenC(i.next)
}

func (c InstrCall) GenC() {
	if c.Next == nil { // Jump
		for i := 0; i < c.size; i++ {
			fmt.Printf("Obj arg = get(vm, %d);\n", -c.size+i)
			fmt.Printf("vmSet(vm, %d, arg);\n", i+1)
		}
		fmt.Printf("vmResize(vm, %d);\n", c.size+1)
		fmt.Printf("vm->pc = code;\n")
	} else { // Call
		fmt.Printf("void fn_yy(struct VM *vm) {\n")
		instrGenC(c.Next)
		fmt.Printf("}\n")
		fmt.Printf("instrCall(vm, %d, fn_yy);\n", c.size)
	}
}

func (i InstrPrimitive) GenC() {
	// fmt.Println("// InstrPrimitive")
	fmt.Printf("primitiveXX();\n")
	if i.Next == nil { // Jump
		fmt.Printf("vmReturn(vm, vm.val);\n")
	} else { // Call
		instrGenC(i.Next)
	}
}

func (i InstrMakeClosure) GenC() {
	fmt.Printf("// move this to global definition section\n")
	fmt.Printf("void fn_xx(struct VM *vm) {\n")
	instrGenC(i.code)
	fmt.Printf("}\n")
	fmt.Printf("vm.val = makeClosure(%d, fn_xx, NULL);\n", i.required)
	instrGenC(i.Next)
}

func (i InstrExit) GenC() {
	fmt.Printf("vmReturn(vm, vm.val);\n")
}
