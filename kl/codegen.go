package kl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	modeShen = MakeSymbol("shen")
	modeCora = MakeSymbol("cora")
)

func codeGen(e Evaluator) {
	mode := e.Get(1)
	inFile := mustString(e.Get(2))
	outFile := mustString(e.Get(3))

	cg := &codeGenerator{
		declare: make(map[Obj]struct{}),
		mode:    mode,
	}

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

	if err := cg.HandleBody(f, out); err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	cg.HandleSymbol(out)

	e.Return(Nil)
	return
}

// 	fileNames := []string{
// 		"toplevel",
// 		"dict",
// 		"core",
// 		"sys",
// 		"sequent",
// 		"yacc",
// 		"reader",
// 		"prolog",
// 		"track",
// 		"load",
// 		"writer",
// 		"macros",
// 		"declarations",
// 		"t-star",
// 		"types",
// 		"init",
// 		"extension-features",
// 		"extension-launcher",
// 		"extension-factorise-defun",
// 		// "extension-programmable-pattern-matching",
// 	}

type codeGenerator struct {
	mode    Obj
	declare map[Obj]struct{}
}

func (cg *codeGenerator) HandleBody(f io.Reader, out io.Writer) error {
	fmt.Fprintf(out, "package main\n\n")
	fmt.Fprintf(out, "import . \"github.com/tiancaiamao/shen-go/kl\"\n\n")
	fmt.Fprintf(out, `func init() {
    __initExprs = append(__initExprs, MakeNative(func(__e Evaluator, __args ...Obj) {
`)
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
	fmt.Fprintf(out, "}, 0))\n}\n")

	return nil
}

func (cg *codeGenerator) HandleSymbol(out io.Writer) {
	for sym, _ := range cg.declare {
		symStr := symbolString(sym)
		symVar := "sym" + symbolAsVar(sym)
		fmt.Fprintf(out, "var %s = MakeSymbol(\"%s\")\n", symVar, symStr)
	}
}

func symbolString(sym Obj) string {
	return GetSymbol(sym)
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
		switch {
		case cg.mode == modeCora:
			fmt.Fprintf(w, "CoraValue(sym%s)", symbolAsVar(sym))
		case cg.mode == modeShen:
			fmt.Fprintf(w, "ShenFunc(sym%s)", symbolAsVar(sym))
		default:
			return errors.New("wrong mode")
		}
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
		fmt.Fprintf(w, "MakeNative(func(__e Evaluator, __args ...Obj) {\n")
		for i, arg := range args {
			fmt.Fprintf(w, "%s := __args[%d]\n", symbolAsVar(arg), i)
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
		str := GetSymbol(c)
		fmt.Fprintf(w, "MakeSymbol(\"%s\")", str)
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
