package re

import (
	"os"
	"bytes"
	"io"
	"fmt"
)

type CodeGen struct {
	globals []io.ReadWriter
	label int
}

func (cg *CodeGen) getLabel() int {
	ret := cg.label
	cg.label++
	return ret
}

func (cg *CodeGen) GenC(i Instr) {
	var to bytes.Buffer
	instrGenC(i, cg, &to)

	for _, out := range cg.globals {
		io.Copy(os.Stdout, out)
		fmt.Println()
	}

	fmt.Printf("void entry(struct VM *vm) {\n")
	io.Copy(os.Stdout, &to)
	fmt.Printf("}\n")
}

func instrGenC(i Instr, cg *CodeGen, to io.Writer) {
	gen, ok := i.(interface{ GenC(*CodeGen, io.Writer) })
	if !ok {
		fmt.Println("unknown instruct")
		panic(i)
	}
	gen.GenC(cg, to)
}

func (c InstrPush) GenC(cg *CodeGen, to io.Writer) {
	fmt.Fprintf(to, "push(vm, vm->val);\n")
	instrGenC(c.Next, cg, to)
}

func (c InstrConst) GenC(cg *CodeGen, to io.Writer) {
	// fmt.Println("// InstrConst")
	switch raw := c.Val.(type) {
	case nilObj:
		fmt.Fprintf(to, "vm->val = Nil;\n")
	case booleanObj:
		if c.Val == True {
			fmt.Fprintf(to, "vm->val = True;\n")
		} else if c.Val == False {
			fmt.Fprintf(to, "vm->val = False;\n")
		}
	case Integer:
		fmt.Fprintf(to, "vm->val = makeNumber(%d);\n", raw)
	case Float64:
		fmt.Fprintf(to, "TODO InstrConst\n")
	case String:
		fmt.Fprintf(to, "vm->val = makeString(\"%s\"));\n", raw)
	case *Symbol:
		fmt.Fprintf(to, "vm->val = makeSymbol(\"%s\"));\n", raw)
	default:
		fmt.Fprintf(to, "InstrConst unsupport type\n")
	}
	instrGenC(c.Next, cg, to)
}

func (i InstrNop) GenC(cg *CodeGen, to io.Writer) {
	// nop
	fmt.Fprintf(to, "vm->val = NULL; // nop\n")
	instrGenC(i.Next, cg, to)
}

func (i InstrLocalRef) GenC(cg *CodeGen, to io.Writer) {
	// fmt.Println("// InstrLocalRef")
	fmt.Fprintf(to, "vm->val = vmGet(vm, %d);\n", i.Idx+2)
	instrGenC(i.Next, cg, to)
}

func (i InstrGlobalRef) GenC(cg *CodeGen, to io.Writer) {
	// fmt.Println("// InstrGlobalRef")
	fmt.Fprintf(to, "vm->val = symbolFn(%s);\n", i.sym)
	instrGenC(i.Next, cg, to)
}

func (i InstrIf) GenC(cg *CodeGen, to io.Writer) {
	// fmt.Println("// InstrIf")
	fmt.Fprintf(to, "if (vm->val == True) {\n")

	instrGenC(i.Then, cg, to)

	fmt.Fprintf(to, "} else if (vm->val == False) {\n")

	instrGenC(i.Else, cg, to)

	fmt.Fprintf(to, "} else {\n")
	fmt.Fprintf(to, "perror(\"if only accept true or false\");\n")
	fmt.Fprintf(to, "}\n")
}

func (i InstrPrepareCall) GenC(cg *CodeGen, to io.Writer) {
	// fmt.Println("// InstrPrepareCall")
	fmt.Fprintf(to, "push(vm, Nil);\n")
	instrGenC(i.next, cg, to)
}

func (c InstrCall) GenC(cg *CodeGen, to io.Writer) {
	if c.Next == nil { // Jump
		for i := 0; i < c.size; i++ {
			fmt.Fprintf(to, "Obj arg = get(vm, %d);\n", -c.size+i)
			fmt.Fprintf(to, "vmSet(vm, %d, arg);\n", i+1)
		}
		fmt.Fprintf(to, "vmResize(vm, %d);\n", c.size+1)
		fmt.Fprintf(to, "vm->pc = code;\n")
	} else { // Call
		label := cg.getLabel()
		var global bytes.Buffer
		fmt.Fprintf(&global, "void fn_label_%d(struct VM *vm) {\n", label)
		instrGenC(c.Next, cg, &global)
		fmt.Fprintf(&global, "}\n")
		cg.globals = append(cg.globals, &global)

		fmt.Fprintf(to, "instrCall(vm, %d, fn_label_%d);\n", c.size, label)
	}
}

func (i InstrPrimitive) GenC(cg *CodeGen, to io.Writer) {
	// fmt.Println("// InstrPrimitive")
	fmt.Fprintf(to, "prim%s(vm);\n", i.prim.Name)
	if i.Next == nil { // Jump
		fmt.Fprintf(to, "vmReturn(vm, vm->val);\n")
	} else { // Call
		instrGenC(i.Next, cg, to)
	}
}

func (i InstrMakeClosure) GenC(cg *CodeGen, to io.Writer) {
	var global bytes.Buffer
	label := cg.getLabel()
	fmt.Fprintf(&global, "void fn_label_%d(struct VM *vm) {\n", label)
	instrGenC(i.code, cg, &global)
	fmt.Fprintf(&global, "}\n")
	cg.globals = append(cg.globals, &global)

	fmt.Fprintf(to, "vm->val = makeClosure(%d, fn_label_%d, NULL);\n", i.required, label)
	instrGenC(i.Next, cg, to)
}

func (i InstrExit) GenC(cg *CodeGen, to io.Writer) {
	fmt.Fprintf(to, "vmReturn(vm, vm->val);\n")
}

func (i InstrClosureRef) GenC(cg *CodeGen, to io.Writer) {
	panic("TODO: instr closure ref")
}

func (i InstrSet) GenC(cg *CodeGen, to io.Writer) {
	panic("TODO: instr closure ref")
}
