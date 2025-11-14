package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"unsafe"

	"github.com/tiancaiamao/shen-go/kl"
)

var makeCodeGenerator = kl.MakeNative(func(e *kl.ControlFlow) {
	// (make-code-generator 'cora)
	cg := &codeGenerator{
		declare: make(map[kl.Obj]struct{}),
	}
	e.Return(kl.MakeRaw(&cg.scmHead))

}, 0)

var bcToGo = kl.MakeNative(func(e *kl.ControlFlow) {
	// (let cg (make-code-generator)
	//      (bc->go cg "Main" true "xx.bc" "xx.go"))
	cg := (*codeGenerator)(unsafe.Pointer(e.Get(1)))
	exportName := kl.GetString(e.Get(2))
	genSym := e.Get(3)
	inFile := kl.GetString(e.Get(4))
	outFile := kl.GetString(e.Get(5))

	f, err := os.Open(inFile)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	defer f.Close()

	out, err := os.Create(outFile)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	defer out.Close()

	if err := cg.HandleBody(f, exportName, out); err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	if genSym == kl.True {
		cg.HandleSymbol(out)
	}

	e.Return(kl.Nil)

}, 5)

type codeGenerator struct {
	scmHead int
	declare map[kl.Obj]struct{}
}

func (cg *codeGenerator) HandleBody(f io.Reader, export string, out io.Writer) error {
	fmt.Fprintf(out, "package main\n\n")
	fmt.Fprintf(out, "import . \"github.com/tiancaiamao/shen-go/kl\"\n\n")
	fmt.Fprintf(out, `var %s = MakeNative(func(__e *ControlFlow) {
`, export)
	r := kl.NewSexpReader(f, false)
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
	for sym := range cg.declare {
		symStr := kl.GetSymbol(sym)
		symVar := "sym" + symbolAsVar(sym)
		fmt.Fprintf(out, "var %s = MakeSymbol(\"%s\")\n", symVar, symStr)
	}
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

func (cg *codeGenerator) generateExpr(w io.Writer, sexp kl.Obj) error {
	// fmt.Printf("handle %s ..\n", ObjString(sexp))
	if kl.IsSymbol(sexp) {
		fmt.Fprintf(w, "%s", symbolAsVar(sexp))
		return nil
	}
	kind := kl.GetSymbol(kl.Car(sexp))
	switch kind {
	case "block":
		for cur := kl.Cdr(sexp); cur != kl.Nil; cur = kl.Cdr(cur) {
			p := kl.Car(cur)
			if err := cg.generateExpr(w, p); err != nil {
				return err
			}
			fmt.Fprintln(w)
		}
		fmt.Fprintln(w)
	case "<-":
		// (<- a b)
		a := kl.Car(kl.Cdr(sexp))
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s = ", symbolAsVar(a))
		cg.generateExpr(w, b)
		fmt.Fprintln(w)
	case "<=":
		a := kl.Car(kl.Cdr(sexp))
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s := ", symbolAsVar(a))
		cg.generateExpr(w, b)
		fmt.Fprintln(w)
	case "$global":
		sym := kl.Car(kl.Cdr(sexp))
		cg.declare[sym] = struct{}{}
		fmt.Fprintf(w, "PrimFunc(sym%s)", symbolAsVar(sym))
	case "$const":
		// (const Number)
		// (const ())
		// (const "xxx")
		c := kl.Cadr(sexp)
		if err := cg.generateConst(w, c); err != nil {
			return err
		}
	case "lambda":
		// (lambda (p1 p2 ...) ...)
		tmp := kl.Car(kl.Cdr(sexp))
		args := kl.ListToSlice(tmp)
		fmt.Fprintf(w, "MakeNative(func(__e *ControlFlow) {\n")
		for i, arg := range args {
			fmt.Fprintf(w, "%s := __e.Get(%d)\n", symbolAsVar(arg), i+1)
			fmt.Fprintf(w, "_ = %s\n", symbolAsVar(arg))
		}
		if err := cg.generateExpr(w, kl.Car(kl.Cdr(kl.Cdr(sexp)))); err != nil {
			return err
		}
		fmt.Fprintf(w, "}, %d)", len(args))
	case "if":
		// (if a b c)
		a := kl.Cadr(sexp)
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		c := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		if err := cg.generateIfExpr(w, a, b, c); err != nil {
			return err
		}
	case "ignore": // (ignore xx)
		fmt.Fprintf(w, "_ = ")
		exp := kl.Car(kl.Cdr(sexp))
		cg.generateExpr(w, exp)
		fmt.Fprintln(w)
	case "var":
		// (var xxx)
		sym := kl.Car(kl.Cdr(sexp))
		fmt.Fprintf(w, "var %s Obj\n", symbolAsVar(sym))
	case "return":
		val := kl.Cadr(sexp)
		fmt.Fprintf(w, "__e.Return(")
		cg.generateExpr(w, val)
		fmt.Fprintf(w, ")\nreturn\n")
	case "call":
		// (call f a b c ...)
		ok, err := cg.primitiveCallOptimize(w, sexp, false)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}

		fmt.Fprintf(w, "Call(__e, ")
		args := kl.ListToSlice(kl.Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if kl.IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")\n")
	case "tailapply":
		// (tailapply f a b c ...)
		ok, err := cg.primitiveCallOptimize(w, sexp, true)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}

		fmt.Fprintf(w, "__e.TailApply(")
		args := kl.ListToSlice(kl.Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if kl.IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")\nreturn\n")
	default:
		return fmt.Errorf("unknown instruct: %s", kind)
	}
	return nil
}

func (cg *codeGenerator) generateConst(w io.Writer, c kl.Obj) error {
	switch {
	case kl.IsNumber(c):
		fmt.Fprintf(w, "MakeNumber(%d)", kl.GetInteger(c))
	case kl.IsString(c):
		str := kl.GetString(c)
		fmt.Fprintf(w, "MakeString(%#v)", str)
	case kl.IsSymbol(c):
		cg.declare[c] = struct{}{}
		fmt.Fprintf(w, "sym%s", symbolAsVar(c))
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

func (cg *codeGenerator) generateIfExpr(w io.Writer, a, b, c kl.Obj) error {
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

var symGlobal = kl.MakeSymbol("$global")
var symConst = kl.MakeSymbol("$const")

func (cg *codeGenerator) primitiveCallOptimize(w io.Writer, sexp kl.Obj, tail bool) (bool, error) {
	// (call ($global XX) ...)
	// (tailapply ($global XX) ...)
	fn := kl.Cadr(sexp)
	if kl.PrimIsPair(fn) == kl.False || kl.Car(fn) != symGlobal {
		return false, nil
	}
	global := kl.Cadr(fn)
	str := kl.GetSymbol(global)
	args := kl.ListToSlice(kl.Cdr(kl.Cdr(sexp)))
	prim, ok := shenPrimitive[str]
	if !ok || prim.Arity != len(args) {
		return false, nil
	}
	primName := prim.Name

	// ($prim f a b c ...)
	if tail {
		fmt.Fprintf(w, "__e.Return(")
	}

	fmt.Fprintf(w, "%s(", primName)
	for i, arg := range args {
		if i != 0 {
			fmt.Fprintf(w, ", ")
		}
		if kl.IsSymbol(arg) {
			fmt.Fprintf(w, "%s", symbolAsVar(arg))
		} else {
			if err := cg.generateExpr(w, arg); err != nil {
				return true, err
			}
		}
	}
	fmt.Fprintf(w, ")")

	if tail {
		fmt.Fprintf(w, ")\nreturn\n")
	}
	return true, nil
}

var shenPrimitive = map[string]struct {
	Arity int
	Name  string
}{
	"get-time":              {1, "PrimGetTime"},
	"close":                 {1, "PrimCloseStream"},
	"open":                  {2, "PrimOpenStream"},
	"read-byte":             {1, "PrimReadByte"},
	"write-byte":            {2, "PrimWriteByte"},
	"absvector?":            {1, "PrimIsVector"},
	"<-address":             {2, "PrimVectorGet"},
	"address->":             {3, "PrimVectorSet"},
	"absvector":             {1, "PrimAbsvector"},
	"str":                   {1, "PrimStr"},
	"<=":                    {2, "PrimLessEqual"},
	">=":                    {2, "PrimGreatEqual"},
	"<":                     {2, "PrimLessThan"},
	">":                     {2, "PrimGreatThan"},
	"error-to-string":       {1, "PrimErrorToString"},
	"simple-error":          {1, "PrimSimpleError"},
	"=":                     {2, "PrimEqual"},
	"-":                     {2, "PrimNumberSubtract"},
	"*":                     {2, "PrimNumberMultiply"},
	"/":                     {2, "PrimNumberDivide"},
	"+":                     {2, "PrimNumberAdd"},
	"string->n":             {1, "PrimStringToNumber"},
	"n->string":             {1, "PrimNumberToString"},
	"number?":               {1, "PrimIsNumber"},
	"string?":               {1, "PrimIsString"},
	"pos":                   {2, "PrimPos"},
	"tlstr":                 {1, "PrimTailString"},
	"cn":                    {2, "PrimStringConcat"},
	"intern":                {1, "PrimIntern"},
	"hd":                    {1, "PrimHead"},
	"tl":                    {1, "PrimTail"},
	"cons":                  {2, "PrimCons"},
	"cons?":                 {1, "PrimIsPair"},
	"value":                 {1, "PrimValue"},
	"set":                   {2, "PrimSet"},
	"not":                   {1, "PrimNot"},
	"if":                    {3, "PrimIf"},
	"symbol?":               {1, "PrimIsSymbol"},
	"read-file-as-bytelist": {1, "PrimReadFileAsByteList"},
	"read-file-as-string":   {1, "PrimReadFileAsString"},
	"variable?":             {1, "PrimIsVariable"},
	"integer?":              {1, "PrimIsInteger"},
}
