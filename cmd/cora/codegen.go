package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"unsafe"

	. "github.com/tiancaiamao/shen-go/kl"
)

var makeCodeGenerator = MakeNative(func(e *ControlFlow) {
	// (make-code-generator 'cora)
	cg := &codeGenerator{
		declare: make(map[Obj]struct{}),
	}
	e.Return(MakeRaw(&cg.scmHead))
	return
}, 0)

var bcToGo = MakeNative(func(e *ControlFlow) {
	// (let cg (make-code-generator)
	//      (bc->go cg "Main" true "xx.bc" "xx.go"))
	cg := (*codeGenerator)(unsafe.Pointer(e.Get(1)))
	exportName := GetString(e.Get(2))
	genSym := e.Get(3)
	inFile := GetString(e.Get(4))
	outFile := GetString(e.Get(5))

	f, err := os.Open(inFile)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()

	out, err := os.Create(outFile)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer out.Close()

	if err := cg.HandleBody(f, exportName, out); err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	if genSym == True {
		cg.HandleSymbol(out)
	}

	e.Return(Nil)
	return
}, 5)

type codeGenerator struct {
	scmHead int
	declare map[Obj]struct{}
}

func (cg *codeGenerator) HandleBody(f io.Reader, export string, out io.Writer) error {
	fmt.Fprintf(out, "package main\n\n")
	fmt.Fprintf(out, "import . \"github.com/tiancaiamao/shen-go/kl\"\n\n")
	fmt.Fprintf(out, `var %s = MakeNative(func(__e *ControlFlow) {
`, export)
	r := NewSexpReader(f, false)
	bc, err := r.Read()
	if err != nil {
		fmt.Println("read bytecode error", err)
		return err
	}
	if err := cg.generateExpr(out, bc); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n\n")
	fmt.Fprintf(out, "}, 0)\n\n")
	return nil
}

func (cg *codeGenerator) HandleSymbol(out io.Writer) {
	for sym, _ := range cg.declare {
		symStr := GetSymbol(sym)
		symVar := "sym" + symbolAsVar(sym)
		fmt.Fprintf(out, "var %s = MakeSymbol(\"%s\")\n", symVar, symStr)
	}
}

func symbolAsVar(sym Obj) string {
	str := GetSymbol(sym)
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
		case ':':
			buf.WriteString("_h")
		case '{':
			buf.WriteString("_i")
		case '}':
			buf.WriteString("_j")
		case ';':
			buf.WriteString("_k")
		case ',':
			buf.WriteString("_l")
		default:
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

func (cg *codeGenerator) generateExpr(w io.Writer, sexp Obj) error {
	// fmt.Printf("handle %s ..\n", ObjString(sexp))
	if IsSymbol(sexp) {
		fmt.Fprintf(w, "%s", symbolAsVar(sexp))
		return nil
	}
	kind := GetSymbol(Car(sexp))
	switch kind {
	case "block":
		for cur := Cdr(sexp); cur != Nil; cur = Cdr(cur) {
			p := Car(cur)
			if err := cg.generateExpr(w, p); err != nil {
				return err
			}
			fmt.Fprintln(w)
		}
		fmt.Fprintln(w)
	case "<-":
		// (<- a b)
		a := Car(Cdr(sexp))
		b := Car(Cdr(Cdr(sexp)))
		fmt.Fprintf(w, "%s = ", symbolAsVar(a))
		cg.generateExpr(w, b)
		fmt.Fprintln(w)
	case "<-:":
		a := Car(Cdr(sexp))
		b := Car(Cdr(Cdr(sexp)))
		fmt.Fprintf(w, "%s := ", symbolAsVar(a))
		cg.generateExpr(w, b)
		fmt.Fprintln(w)
	case "$global":
		sym := Car(Cdr(sexp))
		cg.declare[sym] = struct{}{}
		fmt.Fprintf(w, "PrimNS1Value(sym%s)", symbolAsVar(sym))
	case "$const":
		// (const Number)
		// (const ())
		// (const "xxx")
		c := Cadr(sexp)
		if err := cg.generateConst(w, c); err != nil {
			return err
		}
	case "lambda":
		// (lambda (p1 p2 ...) ...)
		tmp := Car(Cdr(sexp))
		args := ListToSlice(tmp)
		fmt.Fprintf(w, "MakeNative(func(__e *ControlFlow) {\n")
		for i, arg := range args {
			fmt.Fprintf(w, "%s := __e.Get(%d)\n", symbolAsVar(arg), i+1)
			fmt.Fprintf(w, "_ = %s\n", symbolAsVar(arg))
		}
		if err := cg.generateExpr(w, Car(Cdr(Cdr(sexp)))); err != nil {
			return err
		}
		fmt.Fprintf(w, "}, %d)", len(args))
	case "if":
		// (if a b c)
		a := Cadr(sexp)
		b := Car(Cdr(Cdr(sexp)))
		c := Car(Cdr(Cdr(Cdr(sexp))))
		if err := cg.generateIfExpr(w, a, b, c); err != nil {
			return err
		}
	case "ignore": // (ignore xx)
		fmt.Fprintf(w, "_ = ")
		exp := Car(Cdr(sexp))
		cg.generateExpr(w, exp)
		fmt.Fprintln(w)
	case "var":
		// (var xxx)
		sym := Car(Cdr(sexp))
		fmt.Fprintf(w, "var %s Obj\n", symbolAsVar(sym))
	case "return":
		val := Cadr(sexp)
		fmt.Fprintf(w, "__e.Return(")
		cg.generateExpr(w, val)
		fmt.Fprintf(w, ")\nreturn\n")
	case "call":
		// ($call f a b c ...)
		fmt.Fprintf(w, "Call(__e, ")
		args := ListToSlice(Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")\n")
	case "tailapply":
		// ($call f a b c ...)
		fmt.Fprintf(w, "__e.TailApply(")
		args := ListToSlice(Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")\nreturn\n")
	case "$prim":
		// ($prim f a b c ...)
		sym := GetSymbol(Car(Cdr(sexp)))
		prim := Primitives[sym].Name
		fmt.Fprintf(w, "%s(", prim)
		args := ListToSlice(Cdr(Cdr(sexp)))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")")
		fmt.Fprintln(w)
	case "try-catch":
		// (try-catch body handle)
		body := Car(Cdr(sexp))
		handle := Car(Cdr(Cdr(sexp)))
		fmt.Fprintf(w, "Try(__e, %s).Catch(%s)", symbolAsVar(body), symbolAsVar(handle))
	default:
		return fmt.Errorf("unknown instruct: %s\n", kind)
	}
	return nil
}

func (cg *codeGenerator) generateConst(w io.Writer, c Obj) error {
	switch {
	case IsNumber(c):
		fmt.Fprintf(w, "MakeNumber(%d)", GetInteger(c))
	case IsString(c):
		str := GetString(c)
		fmt.Fprintf(w, "MakeString(%#v)", str)
	case IsSymbol(c):
		cg.declare[c] = struct{}{}
		fmt.Fprintf(w, "sym"+symbolAsVar(c))
	case c == Nil:
		fmt.Fprintf(w, "Nil")
	case c == True:
		fmt.Fprintf(w, "True")
	case c == False:
		fmt.Fprintf(w, "False")
	default:
		return errors.New("unknown $const instruct")
	}
	return nil
}

func (cg *codeGenerator) generateIfExpr(w io.Writer, a, b, c Obj) error {
	fmt.Fprintf(w, "if True == ")
	if err := cg.generateExpr(w, a); err != nil {
		return err
	}
	fmt.Fprintf(w, " {\n")
	if err := cg.generateExpr(w, b); err != nil {
		return err
	}
	fmt.Fprintf(w, "} else {\n")
	if err := cg.generateExpr(w, c); err != nil {
		return err
	}
	fmt.Fprintf(w, "}\n")
	return nil
}
