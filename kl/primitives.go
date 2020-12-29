package kl

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

var klPrimitives = []scmPrimitive{
	// &scmPrimitive{scmHead: scmHeadPrimitive, Name: "load-file", Required: 1},
	// &scmPrimitive{scmHead: scmHeadPrimitive, Name: "type", Required: 2, Function: PrimTypeFunc},
	// &scmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "get-time", Required: 1, Function: PrimGetTime},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "close", Required: 1, Function: PrimCloseStream},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "open", Required: 2, Function: PrimOpenStream},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "read-byte", Required: 1, Function: PrimReadByte},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "write-byte", Required: 2, Function: PrimWriteByte},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector?", Required: 1, Function: PrimIsVector},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "<-address", Required: 2, Function: PrimVectorGet},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "address->", Required: 3, Function: PrimVectorSet},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector", Required: 1, Function: PrimAbsvector},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "str", Required: 1, Function: PrimStr},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "<=", Required: 2, Function: PrimLessEqual},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: ">=", Required: 2, Function: PrimGreatEqual},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "<", Required: 2, Function: PrimLessThan},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: ">", Required: 2, Function: PrimGreatThan},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "error-to-string", Required: 1, Function: PrimErrorToString},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "simple-error", Required: 1, Function: PrimSimpleError},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "=", Required: 2, Function: PrimEqual},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "-", Required: 2, Function: PrimNumberSubtract},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "*", Required: 2, Function: PrimNumberMultiply},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "/", Required: 2, Function: PrimNumberDivide},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "+", Required: 2, Function: PrimNumberAdd},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "string->n", Required: 1, Function: PrimStringToNumber},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "n->string", Required: 1, Function: PrimNumberToString},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "number?", Required: 1, Function: PrimIsNumber},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "string?", Required: 1, Function: PrimIsString},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "pos", Required: 2, Function: PrimPos},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "tlstr", Required: 1, Function: PrimTailString},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "cn", Required: 2, Function: PrimStringConcat},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "intern", Required: 1, Function: PrimIntern},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "hd", Required: 1, Function: PrimHead},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "tl", Required: 1, Function: PrimTail},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "cons", Required: 2, Function: PrimCons},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "cons?", Required: 1, Function: PrimIsPair},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "value", Required: 1, Function: PrimValue},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "set", Required: 2, Function: PrimSet},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "not", Required: 1, Function: PrimNot},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "if", Required: 3, Function: PrimIf},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "symbol?", Required: 1, Function: PrimIsSymbol},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "read-file-as-bytelist", Required: 1, Function: PrimReadFileAsByteList},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "read-file-as-string", Required: 1, Function: PrimReadFileAsString},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "variable?", Required: 1, Function: PrimIsVariable},
	scmPrimitive{scmHead: scmHeadPrimitive, Name: "integer?", Required: 1, Function: PrimIsInteger},
}

func init() {
	for i := 0; i < len(klPrimitives); i++ {
		prim := &klPrimitives[i]
		sym := MakeSymbol(prim.Name)
		CoraSet(sym, Obj(&prim.scmHead))
		BindSymbolFunc(sym, Obj(&prim.scmHead))
	}

	// Overload for primitive set and value.
	tmp := &scmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1, Function: primEvalKL}
	BindSymbolFunc(MakeSymbol("eval-kl"), Obj(&tmp.scmHead))
	tmp = &scmPrimitive{scmHead: scmHeadPrimitive, Name: "load-file", Required: 1, Function: primLoadFile(false)}
	BindSymbolFunc(MakeSymbol("load-file"), Obj(&tmp.scmHead))

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
	tmp = &scmPrimitive{scmHead: scmHeadPrimitive, Name: "cora.", Required: 1, Function: coraValue}
	BindSymbolFunc(MakeSymbol("cora."), Obj(&tmp.scmHead))
	tmp = &scmPrimitive{scmHead: scmHeadPrimitive, Name: "defun", Required: 2, Function: primDefun}
	BindSymbolFunc(MakeSymbol("defun"), Obj(&tmp.scmHead))

	CoraSet(MakeSymbol("car"), MakePrimitive("car", 1, PrimHead))
	CoraSet(MakeSymbol("cdr"), MakePrimitive("car", 1, PrimTail))
	CoraSet(MakeSymbol("gensym"), MakePrimitive("gensym", 1, PrimGenSym))
	CoraSet(MakeSymbol("set"), MakePrimitive("set", 2, coraSet))
	CoraSet(MakeSymbol("value"), MakePrimitive("value", 1, coraValue))
	CoraSet(MakeSymbol("load"), MakePrimitive("load", 1, primLoadFile(true)))
	CoraSet(MakeSymbol("load-so"), MakePrimitive("load-so", 1, primLoadSo))
	CoraSet(MakeSymbol("read-file-as-sexp"), MakePrimitive("read-file-as-sexp", 1, readFileAsSexp))
	CoraSet(MakeSymbol("write-sexp-to-file"), MakePrimitive("write-sexp-to-file", 2, writeSexpToFile))
	CoraSet(MakeSymbol("bc->go"), MakePrimitive("bc->go", 3, codeGen))
	// priv = &scmPrimitive{scmHead: scmHeadPrimitive, Name: "apply", Required: 2, Function: e.primApply}
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
	case scmHeadPrimitive:
		prim := mustPrimitive(e.Get(1))
		e.Return(MakeString(fmt.Sprintf("#<primitive %s>", prim.Name)))
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
		e.Return(MakeString(fmt.Sprintf("#<native %v>", *n)))
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
	e.Return(CoraValue(key))
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
