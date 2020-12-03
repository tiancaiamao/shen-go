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

	declare := make(map[kl.Obj]struct{})
	var init bytes.Buffer

	fmt.Fprintf(out, "package main\n\n")
	fmt.Fprintf(out, "import . \"github.com/tiancaiamao/shen-go/kl\"\n\n")
	r := kl.NewSexpReader(f, false)
	bc, err := r.Read()
	if err != nil {
		fmt.Println("read bytecode error", err)
		return
	}

	for bc != kl.Nil {
		curr := kl.Car(bc)
		bc = kl.Cdr(bc)
		if err := generateExpr(declare, &init, curr, bc==kl.Nil); err != nil {
			panic(err)
			return
		}
		fmt.Fprintf(&init, "\n\n")
	}

	for sym, _ := range declare {
		symStr := symbolString(sym)
		symVar := "sym"+symbolAsVar(sym)
		fmt.Fprintf(out, "var %s = MakeSymbol(\"%s\")\n", symVar, symStr)
	}
	// _, err = io.Copy(out, &declare)
	// if err != nil {
	// 	return
	// }
	fmt.Fprintf(out, "\nfunc Main(__e Evaluator, __args ...Obj) {\n")

	_, err = io.Copy(out, &init)
	if err != nil {
		return
	}
	fmt.Fprintf(out, "}")
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
		case ':':
			buf.WriteString("_h")
		default:
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

func generateExpr(declare map[kl.Obj]struct{}, w io.Writer, sexp kl.Obj, tail bool) error {
	// fmt.Printf("handle %s ..\n", kl.ObjString(sexp))
	if kl.IsSymbol(sexp) {
		if tail {
			fmt.Fprintf(w, "__e.Return(%s)\nreturn\n", symbolAsVar(sexp))
		} else {
			fmt.Fprintf(w, "%s", symbolAsVar(sexp))
		}
		return nil
	}
	kind := kl.GetSymbol(kl.Car(sexp))
	switch kind {
	case "block":
		for cur := kl.Cdr(sexp); cur != kl.Nil; cur = kl.Cdr(cur) {
			p := kl.Car(cur)
			tail1 := false
			if kl.Cdr(cur) == kl.Nil {
				tail1 = tail
			}
			if err := generateExpr(declare, w, p, tail1); err != nil {
				return err
			}
		}
		fmt.Fprintln(w)
	case "<-":
		// (<- a b)
		a := kl.Car(kl.Cdr(sexp))
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s = ", symbolAsVar(a))
		generateExpr(declare, w, b, false)
		fmt.Fprintln(w)
	case "<-:":
		a := kl.Car(kl.Cdr(sexp))
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s := ", symbolAsVar(a))
		generateExpr(declare, w, b, false)
		fmt.Fprintln(w)
	case "$global":
		sym := kl.Car(kl.Cdr(sexp))
		declare[sym] = struct{}{}
		fmt.Fprintf(w, "CoraValue(sym%s)", symbolAsVar(sym))
	case "declare":
		// (declare x)
		x := kl.Car(kl.Cdr(sexp))
		fmt.Fprintf(w, "var %s Obj\n", symbolAsVar(x))
	case "lambda":
		// (lambda (p1 p2 ...) ...)
		tmp := kl.Car(kl.Cdr(sexp))
		args := kl.ListToSlice(tmp)
		fmt.Fprintf(w, "MakeNative(func(__e Evaluator, __args ...Obj) {\n")
		for i, arg := range args {
			fmt.Fprintf(w, "%s := __args[%d]\n", symbolAsVar(arg), i)
			fmt.Fprintf(w, "_ = %s\n", symbolAsVar(arg))
		}
		if err := generateExpr(declare, w, kl.Car(kl.Cdr(kl.Cdr(sexp))), true); err != nil {
			return err
		}
		fmt.Fprintf(w, "}, %d)", len(args))
	case "$const":
		// (const Number)
		// (const ())
		// (const "xxx")
		c := kl.Cadr(sexp)
		if tail {
			fmt.Fprintf(w, "__e.Return(")
		}
		if err := generateConst(declare, w, c); err != nil {
			return err
		}
		if tail {
			fmt.Fprintf(w, ")\nreturn\n")
		}
	case "if":
		// (if a b c)
		a := kl.Cadr(sexp)
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		c := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		if err := generateIfExpr(declare, w, a, b, c, tail); err != nil {
			return err
		}
	case "$call":
		// ($call f a b c ...)
		if tail {
			fmt.Fprintf(w, "__e.TailApply(")
		} else {
			fmt.Fprintf(w, "Call(__e, ")
		}
		args := kl.ListToSlice(kl.Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if kl.IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := generateExpr(declare, w, arg, false); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")")
		if tail {
			fmt.Fprintf(w, "\nreturn\n")
		}
	case "try-catch":
		// (try-catch body handle)
		body := kl.Car(kl.Cdr(sexp))
		handle := kl.Car(kl.Cdr(kl.Cdr(sexp)))
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

func generateConst(declare map[kl.Obj]struct{}, w io.Writer, c kl.Obj) error {
	switch {
	case kl.IsNumber(c):
		fmt.Fprintf(w, "MakeNumber(%d)", kl.GetInteger(c))
	case kl.PrimIsString(c) == kl.True:
		str := kl.GetString(c)
		fmt.Fprintf(w, "MakeString(%#v)", str)
	case kl.IsSymbol(c):
		str := kl.GetSymbol(c)
		fmt.Fprintf(w, "MakeSymbol(\"%s\")", str)
	case c == kl.Nil:
		fmt.Fprintf(w, "Nil")
	case c == kl.True:
		fmt.Fprintf(w, "True")
	case c == kl.False:
		fmt.Fprintf(w, "False")
	default:
		return errors.New("unknown $const instruct")
	}
	return nil
}

func generateIfExpr(declare map[kl.Obj]struct{}, w io.Writer, a, b, c kl.Obj, tail bool) error {
	fmt.Fprintf(w, "if True == ")
	if err := generateExpr(declare, w, a, false); err != nil {
		return err
	}
	fmt.Fprintf(w, " {\n")
	if err := generateExpr(declare, w, b, tail); err != nil {
		return err
	}
	fmt.Fprintf(w, "} else {\n")
	if err := generateExpr(declare, w, c, tail); err != nil {
		return err
	}
	fmt.Fprintf(w, "}\n")
	return nil
}

var primMap map[string]*kl.ScmPrimitive

func init() {
	primMap = make(map[string]*kl.ScmPrimitive)
	for _, v := range kl.AllPrimitives {
		primMap[v.Name] = v
	}
}
