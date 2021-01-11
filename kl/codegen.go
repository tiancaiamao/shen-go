package kl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"unsafe"
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
	// (let cg (make-code-generator 'shen)
	//      (bc->go cg "Main" true "xx.bc" "xx.go"))
	cg := (*codeGenerator)(unsafe.Pointer(e.Get(1)))
	exportName := mustString(e.Get(2))
	genSym := e.Get(3)
	inFile := mustString(e.Get(4))
	outFile := mustString(e.Get(5))

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
	for bc != Nil {
		curr := Car(bc)
		bc = Cdr(bc)
		if err := cg.generateExpr(out, curr, bc == Nil); err != nil {
			return err
		}
		fmt.Fprintf(out, "\n\n")
	}
	fmt.Fprintf(out, "}, 0)\n")
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

func (cg *codeGenerator) generateExpr(w io.Writer, sexp Obj, tail bool) error {
	// fmt.Printf("handle %s ..\n", ObjString(sexp))
	if IsSymbol(sexp) {
		if tail {
			fmt.Fprintf(w, "__e.Return(%s)\nreturn\n", symbolAsVar(sexp))
		} else {
			fmt.Fprintf(w, "%s", symbolAsVar(sexp))
		}
		return nil
	}
	kind := GetSymbol(Car(sexp))
	switch kind {
	case "block":
		for cur := Cdr(sexp); cur != Nil; cur = Cdr(cur) {
			p := Car(cur)
			tail1 := false
			if Cdr(cur) == Nil {
				tail1 = tail
			}
			if err := cg.generateExpr(w, p, tail1); err != nil {
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
		cg.generateExpr(w, b, false)
		fmt.Fprintln(w)
	case "<-:":
		a := Car(Cdr(sexp))
		b := Car(Cdr(Cdr(sexp)))
		fmt.Fprintf(w, "%s := ", symbolAsVar(a))
		cg.generateExpr(w, b, false)
		fmt.Fprintln(w)
	case "$global":
		sym := Car(Cdr(sexp))
		cg.declare[sym] = struct{}{}
		fmt.Fprintf(w, "PrimNS1Value(sym%s)", symbolAsVar(sym))
	case "declare":
		// (declare x)
		x := Car(Cdr(sexp))
		fmt.Fprintf(w, "var %s Obj\n", symbolAsVar(x))
	case "lambda":
		// (lambda (p1 p2 ...) ...)
		tmp := Car(Cdr(sexp))
		args := ListToSlice(tmp)
		if tail {
			fmt.Fprintf(w, "__e.Return(")
		}
		fmt.Fprintf(w, "MakeNative(func(__e *ControlFlow) {\n")
		for i, arg := range args {
			fmt.Fprintf(w, "%s := __e.Get(%d)\n", symbolAsVar(arg), i+1)
			fmt.Fprintf(w, "_ = %s\n", symbolAsVar(arg))
		}
		if err := cg.generateExpr(w, Car(Cdr(Cdr(sexp))), true); err != nil {
			return err
		}
		fmt.Fprintf(w, "}, %d)", len(args))
		if tail {
			fmt.Fprintf(w, ")\nreturn\n")
		}
	case "$const":
		// (const Number)
		// (const ())
		// (const "xxx")
		c := Cadr(sexp)
		if tail {
			fmt.Fprintf(w, "__e.Return(")
		}
		if err := cg.generateConst(w, c); err != nil {
			return err
		}
		if tail {
			fmt.Fprintf(w, ")\nreturn\n")
		}
	case "if":
		// (if a b c)
		a := Cadr(sexp)
		b := Car(Cdr(Cdr(sexp)))
		c := Car(Cdr(Cdr(Cdr(sexp))))
		if err := cg.generateIfExpr(w, a, b, c, tail); err != nil {
			return err
		}
	case "$call":
		// ($call f a b c ...)
		if tail {
			fmt.Fprintf(w, "__e.TailApply(")
		} else {
			fmt.Fprintf(w, "Call(__e, ")
		}
		args := ListToSlice(Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg, false); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")\n")
		if tail {
			fmt.Fprintf(w, "\nreturn\n")
		}
	case "try-catch":
		// (try-catch body handle)
		body := Car(Cdr(sexp))
		handle := Car(Cdr(Cdr(sexp)))
		if tail {
			fmt.Fprintf(w, "__e.Return(")
		}
		fmt.Fprintf(w, "Try(__e, %s).Catch(%s)", symbolAsVar(body), symbolAsVar(handle))
		if tail {
			fmt.Fprintf(w, ")\nreturn\n")
		}
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

func (cg *codeGenerator) generateIfExpr(w io.Writer, a, b, c Obj, tail bool) error {
	fmt.Fprintf(w, "if True == ")
	if err := cg.generateExpr(w, a, false); err != nil {
		return err
	}
	fmt.Fprintf(w, " {\n")
	if err := cg.generateExpr(w, b, tail); err != nil {
		return err
	}
	fmt.Fprintf(w, "} else {\n")
	if err := cg.generateExpr(w, c, tail); err != nil {
		return err
	}
	fmt.Fprintf(w, "}\n")
	return nil
}
