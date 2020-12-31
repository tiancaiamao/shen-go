package kl

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"plugin"
	"runtime"
	"time"
)

var klPrimitives = []scmNative{
	// &scmNative{scmHead: scmHeadNative, name: "load-file", require: 1},
	// &scmNative{scmHead: scmHeadNative, name: "type", require: 2, fn: PrimTypeFunc},
	// &scmNative{scmHead: scmHeadNative, name: "eval-kl", require: 1},
	scmNative{scmHead: scmHeadNative, name: "get-time", require: 1, fn: PrimGetTime},
	scmNative{scmHead: scmHeadNative, name: "close", require: 1, fn: PrimCloseStream},
	scmNative{scmHead: scmHeadNative, name: "open", require: 2, fn: PrimOpenStream},
	scmNative{scmHead: scmHeadNative, name: "read-byte", require: 1, fn: PrimReadByte},
	scmNative{scmHead: scmHeadNative, name: "write-byte", require: 2, fn: PrimWriteByte},
	scmNative{scmHead: scmHeadNative, name: "absvector?", require: 1, fn: PrimIsVector},
	scmNative{scmHead: scmHeadNative, name: "<-address", require: 2, fn: PrimVectorGet},
	scmNative{scmHead: scmHeadNative, name: "address->", require: 3, fn: PrimVectorSet},
	scmNative{scmHead: scmHeadNative, name: "absvector", require: 1, fn: PrimAbsvector},
	scmNative{scmHead: scmHeadNative, name: "str", require: 1, fn: PrimStr},
	scmNative{scmHead: scmHeadNative, name: "<=", require: 2, fn: PrimLessEqual},
	scmNative{scmHead: scmHeadNative, name: ">=", require: 2, fn: PrimGreatEqual},
	scmNative{scmHead: scmHeadNative, name: "<", require: 2, fn: PrimLessThan},
	scmNative{scmHead: scmHeadNative, name: ">", require: 2, fn: PrimGreatThan},
	scmNative{scmHead: scmHeadNative, name: "error-to-string", require: 1, fn: PrimErrorToString},
	scmNative{scmHead: scmHeadNative, name: "simple-error", require: 1, fn: PrimSimpleError},
	scmNative{scmHead: scmHeadNative, name: "=", require: 2, fn: PrimEqual},
	scmNative{scmHead: scmHeadNative, name: "-", require: 2, fn: PrimNumberSubtract},
	scmNative{scmHead: scmHeadNative, name: "*", require: 2, fn: PrimNumberMultiply},
	scmNative{scmHead: scmHeadNative, name: "/", require: 2, fn: PrimNumberDivide},
	scmNative{scmHead: scmHeadNative, name: "+", require: 2, fn: PrimNumberAdd},
	scmNative{scmHead: scmHeadNative, name: "string->n", require: 1, fn: PrimStringToNumber},
	scmNative{scmHead: scmHeadNative, name: "n->string", require: 1, fn: PrimNumberToString},
	scmNative{scmHead: scmHeadNative, name: "number?", require: 1, fn: PrimIsNumber},
	scmNative{scmHead: scmHeadNative, name: "string?", require: 1, fn: PrimIsString},
	scmNative{scmHead: scmHeadNative, name: "pos", require: 2, fn: PrimPos},
	scmNative{scmHead: scmHeadNative, name: "tlstr", require: 1, fn: PrimTailString},
	scmNative{scmHead: scmHeadNative, name: "cn", require: 2, fn: PrimStringConcat},
	scmNative{scmHead: scmHeadNative, name: "intern", require: 1, fn: PrimIntern},
	scmNative{scmHead: scmHeadNative, name: "hd", require: 1, fn: PrimHead},
	scmNative{scmHead: scmHeadNative, name: "tl", require: 1, fn: PrimTail},
	scmNative{scmHead: scmHeadNative, name: "cons", require: 2, fn: PrimCons},
	scmNative{scmHead: scmHeadNative, name: "cons?", require: 1, fn: PrimIsPair},
	scmNative{scmHead: scmHeadNative, name: "value", require: 1, fn: PrimValue},
	scmNative{scmHead: scmHeadNative, name: "set", require: 2, fn: PrimSet},
	scmNative{scmHead: scmHeadNative, name: "not", require: 1, fn: PrimNot},
	scmNative{scmHead: scmHeadNative, name: "if", require: 3, fn: PrimIf},
	scmNative{scmHead: scmHeadNative, name: "symbol?", require: 1, fn: PrimIsSymbol},
	scmNative{scmHead: scmHeadNative, name: "read-file-as-bytelist", require: 1, fn: PrimReadFileAsByteList},
	scmNative{scmHead: scmHeadNative, name: "read-file-as-string", require: 1, fn: PrimReadFileAsString},
	scmNative{scmHead: scmHeadNative, name: "variable?", require: 1, fn: PrimIsVariable},
	scmNative{scmHead: scmHeadNative, name: "integer?", require: 1, fn: PrimIsInteger},
}

func init() {
	for i := 0; i < len(klPrimitives); i++ {
		prim := &klPrimitives[i]
		sym := MakeSymbol(prim.name)
		CoraSet(sym, Obj(&prim.scmHead))
		BindSymbolFunc(sym, Obj(&prim.scmHead))
	}

	// Overload for primitive set and value.
	BindSymbolFunc(MakeSymbol("eval-kl"), MakePrimitive("eval-kl", 1, primEvalKL))
	BindSymbolFunc(MakeSymbol("load-file"), MakePrimitive("load-file", 1, primLoadFile(false)))

	primSet(MakeSymbol("*stinput*"), MakeStream(os.Stdin))
	primSet(MakeSymbol("*stoutput*"), MakeStream(os.Stdout))
	dir, _ := os.Getwd()
	primSet(MakeSymbol("*home-directory*"), MakeString(dir))
	primSet(MakeSymbol("*language*"), MakeString("Go"))
	primSet(MakeSymbol("*implementation*"), MakeString("AOT+interpreter"))
	primSet(MakeSymbol("*relase*"), MakeString(runtime.Version()))
	primSet(MakeSymbol("*os*"), MakeString(runtime.GOOS))
	primSet(MakeSymbol("*porters*"), MakeString("Arthur Mao"))
	primSet(MakeSymbol("*port*"), MakeString("1.0.0-rc1"))

	// Extended by shen-go implementation
	primSet(MakeSymbol("*package-path*"), MakeString(PackagePath()))

	BindSymbolFunc(MakeSymbol("cora."), MakePrimitive("cora.", 1, coraValue))
	BindSymbolFunc(MakeSymbol("defun"), MakePrimitive("defun", 2, primDefun))
	CoraSet(MakeSymbol("set"), MakePrimitive("set", 2, coraSet))
	CoraSet(MakeSymbol("value"), MakePrimitive("value", 1, coraValue))

	CoraSet(MakeSymbol("car"), MakePrimitive("car", 1, PrimHead))
	CoraSet(MakeSymbol("cdr"), MakePrimitive("car", 1, PrimTail))
	CoraSet(MakeSymbol("gensym"), MakePrimitive("gensym", 1, PrimGenSym))
	CoraSet(MakeSymbol("load"), MakePrimitive("load", 1, primLoadFile(true)))
	CoraSet(MakeSymbol("load-so"), MakePrimitive("load-so", 1, primLoadSo))
	CoraSet(MakeSymbol("read-file-as-sexp"), MakePrimitive("read-file-as-sexp", 1, readFileAsSexp))
	CoraSet(MakeSymbol("write-sexp-to-file"), MakePrimitive("write-sexp-to-file", 2, writeSexpToFile))
	CoraSet(MakeSymbol("make-code-generator"), makeCodeGenerator)
	CoraSet(MakeSymbol("cg:bc->go"), bcToGo)
	// priv = &scmNative{scmHead: scmHeadNative, Name: "apply", require: 2, fn: e.primApply}
	// coraSet(MakeSymbol("apply"), Obj(&priv.scmHead))
}

func PrimNumberAdd(e Evaluator) {
	x1 := mustNumber(e.Get(1))
	y1 := mustNumber(e.Get(2))
	e.Return(MakeNumber(x1 + y1))
	return
}

func PrimNumberSubtract(e Evaluator) {
	x1 := mustNumber(e.Get(1))
	y1 := mustNumber(e.Get(2))
	e.Return(MakeNumber(x1 - y1))
	return
}

func PrimNumberMultiply(e Evaluator) {
	x1 := mustNumber(e.Get(1))
	y1 := mustNumber(e.Get(2))
	e.Return(MakeNumber(x1 * y1))
	return
}

func PrimNumberDivide(e Evaluator) {
	x1 := mustNumber(e.Get(1))
	y1 := mustNumber(e.Get(2))
	e.Return(MakeNumber(x1 / y1))
	return
}

func PrimIntern(e Evaluator) {
	str := mustString(e.Get(1))
	switch str {
	case "true":
		e.Return(True)
		return
	case "false":
		e.Return(False)
		return
	}
	e.Return(MakeSymbol(str))
	return
}

func PrimHead(e Evaluator) {
	e.Return(car(e.Get(1)))
	return
}

func PrimTail(e Evaluator) {
	e.Return(cdr(e.Get(1)))
	return
}

func PrimIsNumber(e Evaluator) {
	if *e.Get(1) == scmHeadNumber {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimStringToNumber(e Evaluator) {
	str := mustString(e.Get(1))
	n := ([]rune(str))[0]
	e.Return(MakeInteger(int(n)))
	return
}

func PrimNumberToString(e Evaluator) {
	n := mustInteger(e.Get(1))
	e.Return(MakeString(string(rune(n))))
	return
}

func PrimStr(e Evaluator) {
	switch *e.Get(1) {
	case scmHeadPair:
		// Pair may contain recursive list.
		panic(MakeError("can't str pair object"))
	case scmHeadNull:
		e.Return(MakeString("()"))
		return
	case scmHeadSymbol:
		str := GetSymbol(e.Get(1))
		e.Return(MakeString(str))
		return
	case scmHeadNumber:
		if isFixnum(e.Get(1)) {
			e.Return(MakeString(fmt.Sprintf("%d", mustInteger(e.Get(1)))))
			return
		}
		f := mustNumber(e.Get(1))
		if !isPreciseInteger(f) {
			e.Return(MakeString(fmt.Sprintf("%f", f)))
			return
		}
		e.Return(MakeString(fmt.Sprintf("%d", int(f))))
		return
	case scmHeadString:
		e.Return(MakeString(fmt.Sprintf(`"%s"`, mustString(e.Get(1)))))
		return
	case scmHeadProcedure:
		e.Return(MakeString("#<procedure >"))
		return
	case scmHeadBoolean:
		if e.Get(1) == True {
			e.Return(MakeString("true"))
			return
		} else if e.Get(1) == False {
			e.Return(MakeString("false"))
			return
		}
	case scmHeadError:
		err := mustError(e.Get(1))
		e.Return(MakeString("#<err " + err.err + ">"))
		return
	case scmHeadStream:
		e.Return(MakeString("#<stream >"))
		return
	case scmHeadRaw:
		e.Return(MakeString("#<raw >"))
		return
	case scmHeadNative:
		n := MustNative(e.Get(1))
		if len(n.name) > 0 {
			e.Return(MakeString(fmt.Sprintf("#<primitive %s>", n.name)))
		} else {
			e.Return(MakeString(fmt.Sprintf("#<native %v>", *n)))
		}
		return
	default:
		e.Return(MakeString("PrimStr unknown"))
		return
	}
	e.Return(MakeString("wrong input, the object is not atom ..."))
	return
}

func PrimStringConcat(e Evaluator) {
	s1 := mustString(e.Get(1))
	s2 := mustString(e.Get(2))
	e.Return(MakeString(s1 + s2))
	return
}

func PrimTailString(e Evaluator) {
	str := mustString(e.Get(1))
	if len(str) == 0 {
		panic(MakeError("empty string"))
	}
	e.Return(MakeString(string([]rune(str)[1:])))
	return
}

func PrimPos(e Evaluator) {
	s := []rune(mustString(e.Get(1)))
	n := mustInteger(e.Get(2))
	if n >= len(s) {
		panic(MakeError(fmt.Sprintf("%d is not valid index for %s", n, string(s))))
	}
	e.Return(MakeString(string([]rune(s)[n])))
	return
}

func and(e Evaluator) {
	if e.Get(1) == True && e.Get(2) == True {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func or(e Evaluator) {
	if e.Get(1) == False || e.Get(2) == False {
		e.Return(False)
		return
	}
	e.Return(True)
	return
}

func primSet(key, val Obj) Obj {
	sym := mustSymbol(key)
	sym.value = val
	return val
}

func PrimSet(e Evaluator) {
	key, val := e.Get(1), e.Get(2)
	sym := mustSymbol(key)
	sym.value = val
	e.Return(val)
	return
}

func PrimValue(e Evaluator) {
	key := e.Get(1)
	sym := mustSymbol(key)
	if sym.value != nil {
		e.Return(sym.value)
		return
	}
	panic(MakeError(fmt.Sprintf("variable %s not bound", sym.str)))
}

func coraSet(e Evaluator) {
	key, val := e.Get(1), e.Get(2)
	CoraSet(key, val)
	e.Return(val)
}

func coraValue(e Evaluator) {
	key := e.Get(1)
	e.Return(e.Global(key))
	return
}

func primDefun(e Evaluator) {
	key, val := e.Get(1), e.Get(2)
	sym := mustSymbol(key)
	sym.function = val
	e.Return(val)
	return
}

func PrimSimpleError(e Evaluator) {
	str := mustString(e.Get(1))
	panic(MakeError(str))
}

func PrimErrorToString(e Evaluator) {
	err := mustError(e.Get(1))
	e.Return(MakeString(err.err))
	return
}

func PrimGreatThan(e Evaluator) {
	x := mustInteger(e.Get(1))
	y := mustInteger(e.Get(2))
	if x > y {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimLessThan(e Evaluator) {
	x := mustInteger(e.Get(1))
	y := mustInteger(e.Get(2))
	if x < y {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimLessEqual(e Evaluator) {
	x := mustInteger(e.Get(1))
	y := mustInteger(e.Get(2))
	if x <= y {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimGreatEqual(e Evaluator) {
	x := mustInteger(e.Get(1))
	y := mustInteger(e.Get(2))
	if x >= y {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimAbsvector(e Evaluator) {
	n := mustInteger(e.Get(1))
	if n < 0 {
		panic(MakeError("absvector wrong argument"))
	}
	e.Return(MakeVector(n))
	return
}

func PrimVectorSet(e Evaluator) {
	vec := mustVector(e.Get(1))
	off := mustInteger(e.Get(2))
	val := e.Get(3)
	vec[off] = val
	e.Return(e.Get(1))
	return
}

func PrimVectorGet(e Evaluator) {
	vec := mustVector(e.Get(1))
	off := mustInteger(e.Get(2))
	if off >= len(vec) {
		panic(MakeError(fmt.Sprintf("index %d out of range %d", off, len(vec))))
	}
	ret := vec[off]
	if ret == nil {
		e.Return(undefined)
		return
	}
	e.Return(ret)
	return
}

func PrimIsVector(e Evaluator) {
	if *e.Get(1) == scmHeadVector {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimWriteByte(e Evaluator) {
	n := mustInteger(e.Get(1))
	s := mustStream(e.Get(2))
	w, ok := s.raw.(io.Writer)
	if !ok {
		panic(MakeError("stream is not opened in out mode"))
	}
	var b [1]byte
	b[0] = byte(n)
	_, err := w.Write(b[:])
	if err != nil {
		panic(MakeError(err.Error()))
	}
	e.Return(e.Get(1))
	return
}

func PrimReadByte(e Evaluator) {
	s := mustStream(e.Get(1))
	r, ok := s.raw.(io.Reader)
	if !ok {
		panic(MakeError("stream is closed of not opened in in mode"))
	}
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		if err == io.EOF {
			e.Return(MakeInteger(-1))
			return
		}
		panic(MakeError(err.Error()))
	}
	e.Return(MakeInteger(int(buf[0])))
	return
}

func PrimOpenStream(e Evaluator) {
	file := mustString(e.Get(1))
	var flag int
	mode := GetSymbol(e.Get(2))
	switch mode {
	case "in":
		flag |= os.O_RDONLY
	case "out":
		flag |= os.O_WRONLY | os.O_CREATE
	default:
		flag = os.O_RDWR | os.O_CREATE
	}

	f, err := os.OpenFile(file, flag, 0666)
	if err != nil {
		panic(MakeError(err.Error()))
	}
	e.Return(MakeStream(f))
	return
}

func PrimCloseStream(e Evaluator) {
	s := mustStream(e.Get(1))
	c := s.raw.(io.Closer)
	c.Close()
	e.Return(Nil)
	return
}

func PrimGetTime(e Evaluator) {
	kind := GetSymbol(e.Get(1))
	switch kind {
	case "unix":
		e.Return(MakeNumber(float64(time.Now().Unix())))
		return
	case "run":
		e.Return(MakeNumber(time.Since(uptime).Seconds()))
		return
	}
	panic(MakeError(fmt.Sprintf("get-time does not understand the parameter %s", kind)))
}

func PrimIsString(e Evaluator) {
	if *e.Get(1) == scmHeadString {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimIsPair(e Evaluator) {
	if *e.Get(1) == scmHeadPair {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimNot(e Evaluator) {
	if e.Get(1) == False {
		e.Return(True)
		return
	} else if e.Get(1) == True {
		e.Return(False)
		return
	}
	panic(MakeError("PrimNot"))

}

func PrimIf(e Evaluator) {
	switch e.Get(1) {
	case True:
		e.Return(e.Get(2))
		return
	case False:
		e.Return(e.Get(3))
		return
	}
	panic(MakeError("PrimIf"))
}

func PrimEqual(e Evaluator) {
	e.Return(equal(e.Get(1), e.Get(2)))
	return
}

func PrimCons(e Evaluator) {
	e.Return(cons(e.Get(1), e.Get(2)))
	return
}

func PrimIsSymbol(e Evaluator) {
	if IsSymbol(e.Get(1)) {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

func PrimReadFileAsByteList(e Evaluator) {
	fileName := mustString(e.Get(1))
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(MakeError(err.Error()))
	}

	ret := cons(Nil, Nil)
	curr := mustPair(ret)
	for _, b := range buf {
		tmp := cons(MakeInteger(int(b)), Nil)
		curr.cdr = tmp
		curr = mustPair(tmp)
	}
	e.Return(cdr(ret))
	return
}

func PrimReadFileAsString(e Evaluator) {
	fileName := mustString(e.Get(1))
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(MakeError(err.Error()))
	}

	e.Return(MakeString(string(buf)))
	return
}

func PrimIsVariable(e Evaluator) {
	if *e.Get(1) != scmHeadSymbol {
		e.Return(False)
		return
	}

	sym := GetSymbol(e.Get(1))
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		e.Return(False)
		return
	}
	e.Return(True)
	return
}

func PrimIsInteger(e Evaluator) {
	if *e.Get(1) != scmHeadNumber {
		e.Return(False)
		return
	}
	if isFixnum(e.Get(1)) {
		e.Return(True)
		return
	}
	f := mustNumber(e.Get(1))
	if isPreciseInteger(f) {
		e.Return(True)
		return
	}
	e.Return(False)
	return
}

// func PrimEvalKL(e Evaluator, args ...Obj) Obj {
// 	e.Return(evalExp(e, e.Get(1), Nil))
// 	return
// }

var genIdx uint64 = 0

func PrimGenSym(e Evaluator) {
	sym := mustSymbol(e.Get(1))
	str := fmt.Sprintf("%s%d", sym.str, genIdx)
	genIdx++
	e.Return(MakeSymbol(str))
	return
}

func primEvalKL(e Evaluator) {
	e.Return(evalExp(e, e.Get(1), Nil))
	return
}

func primLoadFile(isCora bool) func(e Evaluator) {
	return func(e Evaluator) {
		path := mustString(e.Get(1))
		res := loadFile(e, isCora, path)
		e.Return(res)
		return
	}
}

func loadFile(e Evaluator, extended bool, file string) Obj {
	var filePath string
	if _, err := os.Stat(file); err == nil {
		filePath = file
	} else {
		filePath = path.Join(PackagePath(), file)
		if _, err := os.Stat(filePath); err != nil {
			return MakeError(err.Error())
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return MakeError(err.Error())
	}
	defer f.Close()

	r := NewSexpReader(f, extended)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return MakeError(err.Error())
			}
			break
		}

		// Macro expand for cora.
		if extended {
			expand := e.Global(symMacroExpand)
			if expand != Nil {
				exp = Call(e, expand, exp)
			}
		}

		res := evalExp(e, exp, Nil)
		if *res == scmHeadError {
			return res
		}
	}
	return MakeString(file)
}

func primLoadSo(e Evaluator) {
	pluginPath := GetString(e.Get(1))
	p, err := plugin.Open(pluginPath)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}

	entry, err := p.Lookup("Main")
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}

	f, ok := entry.(func(__e Evaluator))
	if !ok {
		e.Return(MakeError("plugin Main should be func(__e Evaluator)"))
		return
	}

	res := Call(e, MakeNative(f, 0))
	e.Return(res)
	return
}

func readFileAsSexp(e Evaluator) {
	filePath := mustString(e.Get(1))
	f, err := os.Open(filePath)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()

	ret := Nil
	r := NewSexpReader(f, true)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				e.Return(MakeError(err.Error()))
				return
			}
			break
		}
		ret = cons(exp, ret)
	}
	e.Return(reverse(ret))
	return
}

func writeSexpToFile(e Evaluator) {
	filePath := mustString(e.Get(1))
	str := ObjString(e.Get(2))
	err := ioutil.WriteFile(filePath, []byte(str), 0644)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	e.Return(Nil)
	return
}
