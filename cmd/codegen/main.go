package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/tiancaiamao/shen-go/kl"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: xxx from.scm to.go")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		panic(err)
		return
	}
	defer out.Close()

	var declare bytes.Buffer
	var init bytes.Buffer

	fmt.Fprintf(out, "package main\n\n")
	fmt.Fprintf(out, "import . \"github.com/tiancaiamao/shen-go/kl\"\n\n")
	r := kl.NewSexpReader(f)
	for {
		bc, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read bytecode error", err)
			return
		}

		if err := generateExpr(&declare, &init, bc); err != nil {
			panic(err)
			return
		}
	}
	_, err = io.Copy(out, &declare)
	if err != nil {
		return
	}
	fmt.Fprintf(out, "\nfunc init() {\n")

	_, err = io.Copy(out, &init)
	if err != nil {
		return
	}
	fmt.Fprintf(out, "}\n")
}

func symbolString(sym kl.Obj) string {
	return kl.GetSymbol(sym)
}

func symbolAsVar(sym kl.Obj) string {
	str := kl.GetSymbol(sym)
	var buf bytes.Buffer
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '_':
			buf.WriteString("__")
		case '-':
			buf.WriteString("_1")
		case '?':
			buf.WriteString("_2")
		case '$':
			buf.WriteString("_3")
		case '.':
			buf.WriteString("_4")
		case '<':
			buf.WriteString("_5")
		case '>':
			buf.WriteString("_6")
		case '+':
			buf.WriteString("_7")
		case '@':
			buf.WriteString("_8")
		case '=':
			buf.WriteString("_a")
		case '!':
			buf.WriteString("_b")
		case '/':
			buf.WriteString("_c")
		case '*':
			buf.WriteString("_d")
		case '&':
			buf.WriteString("_e")
		case '%':
			buf.WriteString("_f")
		case '^':
			buf.WriteString("_g")
		default:
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

func defunSymbolVar(sym kl.Obj) string {
	return "__defun__" + symbolAsVar(sym)
}

func generateExprs(declare, init io.Writer, sexp kl.Obj) error {
	instructs := kl.ListToSlice(sexp)
	for _, inst := range instructs {
		if err := generateExpr(declare, init, inst); err != nil {
			return err
		}
	}
	return nil
}

func generateExpr(declare, w io.Writer, sexp kl.Obj) error {
	fmt.Printf("handle %s ...\n", kl.ObjString(sexp))
	kind := kl.GetSymbol(kl.Car(sexp))
	switch kind {
	case "$defun":
		name := defunSymbolVar(kl.Cadr(sexp))
		fmt.Fprintf(declare, "var %s Obj\n", name)
		fmt.Fprintf(w, "%s = MakeNative(func(__ctx *Trampoline, __args ...Obj) {\n", name)
		args := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		args1 := kl.ListToSlice(args)
		for i, arg := range args1 {
			fmt.Fprintf(w, "%s := __args[%d]\n", symbolAsVar(arg), i)
			fmt.Fprintf(w, "_ = %s\n", symbolAsVar(arg))
		}
		if err := generateExprs(declare, w, kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))); err != nil {
			return err
		}
		fmt.Fprintf(w, "}, %d)\n", len(args1))
	case "$lambda":
		dest := kl.Cadr(sexp)
		args := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s := MakeNative(func(__ctx *Trampoline, __args ...Obj) {\n", symbolString(dest))
		fmt.Fprintf(w, "%s := __args[0]\n", symbolAsVar(args))
		fmt.Fprintf(w, "_ = %s\n", symbolAsVar(args))
		if err := generateExprs(declare, w, kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))); err != nil {
			return err
		}
		fmt.Fprintf(w, "}, 1)\n")
	case "mov":
		// (mov SRC DST)
		src := kl.Car(kl.Cdr(sexp))
		dst := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s = %s\n", symbolString(dst), symbolAsVar(src))
	case "$declare-mov":
		// (mov SRC DST)
		src := kl.Car(kl.Cdr(sexp))
		dst := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s := %s\n", symbolAsVar(dst), symbolAsVar(src))
	case "$const":
		// (const Number DST)
		// (const () DST)
		// (const "xxx" DST)
		dst := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		c := kl.Cadr(sexp)
		switch {
		case kl.IsNumber(c):
			fmt.Fprintf(w, "%s := MakeNumber(%d)\n", symbolString(dst), kl.GetInteger(kl.Cadr(sexp)))
		case kl.PrimIsString(c) == kl.True:
			str := kl.GetString(c)
			fmt.Fprintf(w, "%s := MakeString(%#v)\n", symbolString(dst), str)
		case kl.IsSymbol(c):
			str := kl.GetSymbol(c)
			fmt.Fprintf(w, "%s := MakeSymbol(\"%s\")\n", symbolString(dst), str)
		case c == kl.Nil:
			fmt.Fprintf(w, "%s := Nil;\n", symbolString(dst))
		case c == kl.True:
			fmt.Fprintf(w, "%s := True;\n", symbolString(dst))
		case c == kl.False:
			fmt.Fprintf(w, "%s := False;\n", symbolString(dst))
		default:
			return errors.New("unknown $const instruct")
		}
	case "$declare":
		// ($declare x)
		tmp := kl.Cadr(sexp)
		fmt.Fprintf(w, "var %s Obj\n", symbolString(tmp))
	case "if":
		// (if a b c)
		a := kl.Cadr(sexp)
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		c := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		if err := generateIfExpr(declare, w, a, b, c); err != nil {
			return err
		}
	case "$builtin":
		// (builtin OP (ARG1 ARG2 ...) DST)
		op := kl.Cadr(sexp)
		args := kl.Cadr(kl.Cdr(sexp))
		dst := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		generateBuiltinCall(w, op, args, dst)
	case "$tailcall-def":
		fname := kl.Cadr(sexp)
		args := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "__ctx.TailApply(%s", defunSymbolVar(fname))
		for _, arg := range kl.ListToSlice(args) {
			fmt.Fprintf(w, ", ")
			fmt.Fprintf(w, "%s", symbolAsVar(arg))
		}
		fmt.Fprintf(w, ")\nreturn\n")
	case "$call-def":
		fname := kl.Cadr(sexp)
		args := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		res := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		fmt.Fprintf(w, "%s := Call(%s", symbolAsVar(res), defunSymbolVar(fname))
		for _, arg := range kl.ListToSlice(args) {
			fmt.Fprintf(w, ", ")
			fmt.Fprintf(w, "%s", symbolAsVar(arg))
		}
		fmt.Fprintf(w, ")\n")
	case "$tailcall":
		fname := kl.Cadr(sexp)
		args := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "__ctx.TailApply(%s", symbolAsVar(fname))
		for _, arg := range kl.ListToSlice(args) {
			fmt.Fprintf(w, ", ")
			fmt.Fprintf(w, "%s", symbolAsVar(arg))
		}
		fmt.Fprintf(w, ")\nreturn\n")
	case "$call":
		fname := kl.Cadr(sexp)
		args := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		res := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		fmt.Fprintf(w, "%s := Call(%s", symbolAsVar(res), symbolAsVar(fname))
		for _, arg := range kl.ListToSlice(args) {
			fmt.Fprintf(w, ", ")
			fmt.Fprintf(w, "%s", symbolAsVar(arg))
		}
		fmt.Fprintf(w, ")\n")
	case "$try-catch":
		exp := kl.Cadr(sexp)
		if err := generateExpr(declare, w, exp); err != nil {
			return err
		}
		handle := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		if err := generateExpr(declare, w, handle); err != nil {
			return err
		}
		res := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		exp1 := kl.Cadr(exp)
		handle1 := kl.Cadr(handle)
		fmt.Fprintf(w, "%s := Try(%s).Catch(%s)\n", symbolAsVar(res), symbolAsVar(exp1), symbolAsVar(handle1))
	case "$return":
		fmt.Fprintf(w, "__ctx.Return(%s)\nreturn\n", symbolAsVar(kl.Cadr(sexp)))
	default:
		return fmt.Errorf("unknown instruct: %s\n", kind)
	}
	return nil
}

func generateIfExpr(declare, w io.Writer, a, b, c kl.Obj) error {
	fmt.Fprintf(w, "if %s == True {\n", symbolString(a))
	if err := generateExprs(declare, w, b); err != nil {
		return err
	}
	fmt.Fprintf(w, "} else {\n")
	if err := generateExprs(declare, w, c); err != nil {
		return err
	}
	fmt.Fprintf(w, "}\n")
	return nil
}

func generateBuiltinCall(w io.Writer, op, args, dst kl.Obj) {
	input := kl.ListToSlice(args)
	name := kl.GetSymbol(op)
	prim, ok := primMap[name]
	if !ok {
		fmt.Fprintf(w, "error, unknown builtin %s\n", name)
		return
	}

	fmt.Fprintf(w, "%s := %s(", symbolAsVar(dst), prim.CodeGen)
	for i := 0; i < len(input); i++ {
		if i != 0 {
			fmt.Fprintf(w, ", ")
		}
		fmt.Fprintf(w, "%s", symbolAsVar(input[i]))
	}
	fmt.Fprintf(w, ")\n")
}

var primMap map[string]*kl.ScmPrimitive

func init() {
	primMap = make(map[string]*kl.ScmPrimitive)
	for _, v := range kl.AllPrimitives {
		primMap[v.Name] = v
	}
}
