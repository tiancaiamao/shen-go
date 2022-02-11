package cora

import (
	"embed"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"plugin"
	"runtime"
	"time"
	"unsafe"
)

var Primitives = map[string]struct {
	Arity int
	fn    interface{}
	Name  string
}{
	"set":         {2, PrimNS1Set, "PrimNS1Set"},
	"value":       {1, PrimNS1Value, "PrimNS1Value"},
	"car":         {1, PrimHead, "PrimHead"},
	"cdr":         {1, PrimTail, "PrimTail"},
	"vector?":     {1, PrimIsVector, "PrimIsVector"},
	"vector-ref":  {2, PrimVectorGet, "PrimVectorGet"},
	"vector-set!": {3, PrimVectorSet, "PrimVectorSet"},
	"vector":      {1, PrimAbsvector, "PrimAbsvector"},
	"<=":          {2, PrimLessEqual, "PrimLessEqual"},
	">=":          {2, PrimGreatEqual, "PrimGreatEqual"},
	"<":           {2, PrimLessThan, "PrimLessThan"},
	">":           {2, PrimGreatThan, "PrimGreatThan"},
	"=":           {2, PrimEqual, "PrimEqual"},
	"-":           {2, PrimNumberSubtract, "PrimNumberSubtract"},
	"*":           {2, PrimNumberMultiply, "PrimNumberMultiply"},
	"/":           {2, PrimNumberDivide, "PrimNumberDivide"},
	"+":           {2, PrimNumberAdd, "PrimNumberAdd"},
	"number?":     {1, PrimIsNumber, "PrimIsNumber"},
	"string?":     {1, PrimIsString, "PrimIsString"},
	"pos":         {2, PrimPos, "PrimPos"},
	"cn":          {2, PrimStringConcat, "PrimStringConcat"},
	"intern":      {1, PrimIntern, "PrimIntern"},
	"cons":        {2, PrimCons, "PrimCons"},
	"cons?":       {1, PrimIsPair, "PrimIsPair"},
	"not":         {1, PrimNot, "PrimNot"},
	"symbol?":     {1, PrimIsSymbol, "PrimIsSymbol"},
	"gensym":      {1, PrimGenSym, "PrimGenSym"},
	"integer?":    {1, PrimIsInteger, "PrimIsInteger"},

	// The only special primitive required to support KLambda
	"fn": {1, PrimNS2Value, "PrimNS2Value"},
}

func init() {
	// Cora primitives
	for name, item := range Primitives {
		sym := MakeSymbol(name)
		prim := MakePrimitive(name, item.Arity, item.fn)
		PrimNS1Set(sym, prim)
	}
	PrimNS1Set(MakeSymbol("load"), MakeNative(PrimLoadFile(false), 1))
	PrimNS1Set(MakeSymbol("load-so"), MakeNative(primLoadSo, 1))
	PrimNS1Set(MakeSymbol("read-file-as-sexp"), MakeNative(readFileAsSexp, 2))
	PrimNS1Set(MakeSymbol("write-sexp-to-file"), MakeNative(writeSexpToFile, 2))
	PrimNS1Set(MakeSymbol("try-catch"), MakeNative(PrimTryCatch(false), 2))
}

func PrimNumberAdd(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	return MakeNumber(x1 + y1)
}

func PrimNumberSubtract(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	return MakeNumber(x1 - y1)
}

func PrimNumberMultiply(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	return MakeNumber(x1 * y1)
}

func PrimNumberDivide(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	return MakeNumber(x1 / y1)
}

func PrimIntern(o Obj) Obj {
	str := mustString(o)
	switch str {
	case "true":
		return True
	case "false":
		return False
	}
	return MakeSymbol(str)
}

func PrimHead(o Obj) Obj {
	return car(o)
}

func PrimTail(o Obj) Obj {
	return cdr(o)
}

func PrimIsNumber(o Obj) Obj {
	if objType(o) == scmHeadNumber {
		return True
	}
	return False
}

func PrimStringToNumber(o Obj) Obj {
	str := mustString(o)
	n := ([]rune(str))[0]
	return MakeInteger(int(n))
}

func PrimNumberToString(o Obj) Obj {
	n := mustInteger(o)
	return MakeString(string(rune(n)))
}

func PrimStringConcat(x, y Obj) Obj {
	s1 := mustString(x)
	s2 := mustString(y)
	return MakeString(s1 + s2)
}

func PrimTailString(o Obj) Obj {
	str := mustString(o)
	if len(str) == 0 {
		panic(MakeError("empty string"))
	}
	return MakeString(string([]rune(str)[1:]))
}

func and(x, y Obj) Obj {
	if x == True && y == True {
		return True
	}
	return False
}

func or(x, y Obj) Obj {
	if x == False || y == False {
		return False
	}
	return True
}

func PrimGreatThan(x, y Obj) Obj {
	x1 := mustInteger(x)
	y1 := mustInteger(y)
	if x1 > y1 {
		return True

	}
	return False
}

func PrimLessThan(a, b Obj) Obj {
	x := mustInteger(a)
	y := mustInteger(b)
	if x < y {
		return True

	}
	return False
}

func PrimLessEqual(a, b Obj) Obj {
	x := mustInteger(a)
	y := mustInteger(b)
	if x <= y {
		return True

	}
	return False

}

func PrimGreatEqual(a, b Obj) Obj {
	x := mustInteger(a)
	y := mustInteger(b)
	if x >= y {
		return True

	}
	return False
}

func PrimAbsvector(o Obj) Obj {
	n := mustInteger(o)
	if n < 0 {
		panic(MakeError("absvector wrong argument"))
	}
	return MakeVector(n)
}

func PrimVectorSet(x, y, z Obj) Obj {
	vec := mustVector(x)
	off := mustInteger(y)
	val := z
	vec[off] = val
	return x
}

func PrimVectorGet(x, y Obj) Obj {
	vec := mustVector(x)
	off := mustInteger(y)
	if off >= len(vec) {
		panic(MakeError(fmt.Sprintf("index %d out of range %d", off, len(vec))))
	}
	ret := vec[off]
	if ret == nil {
		return undefined
	}
	return ret
}

func PrimIsVector(x Obj) Obj {
	if objType(x) == scmHeadVector {
		return True

	}
	return False
}

func PrimWriteByte(x, y Obj) Obj {
	n := mustInteger(x)
	s := mustStream(y)
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
	return x
}

func PrimReadByte(x Obj) Obj {
	s := mustStream(x)
	r, ok := s.raw.(io.Reader)
	if !ok {
		panic(MakeError("stream is closed of not opened in in mode"))
	}
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		if err == io.EOF {
			return MakeInteger(-1)
		}
		panic(MakeError(err.Error()))
	}
	return MakeInteger(int(buf[0]))
}

func PrimOpenStream(x, y Obj) Obj {
	file := mustString(x)
	var flag int
	mode := GetSymbol(y)
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
	return MakeStream(f)
}

func PrimCloseStream(x Obj) Obj {
	s := mustStream(x)
	c := s.raw.(io.Closer)
	c.Close()
	return Nil
}

func PrimGetTime(x Obj) Obj {
	kind := GetSymbol(x)
	switch kind {
	case "unix":
		return MakeNumber(float64(time.Now().Unix()))
	case "run":
		return MakeNumber(time.Since(uptime).Seconds())

	}
	panic(MakeError(fmt.Sprintf("get-time does not understand the parameter %s", kind)))
}

func PrimIsString(x Obj) Obj {
	if objType(x) == scmHeadString {
		return True

	}
	return False
}

func PrimIsPair(x Obj) Obj {
	if objType(x) == scmHeadPair {
		return True
	}
	return False
}

func PrimNot(x Obj) Obj {
	if x == False {
		return True
	} else if x == True {
		return False
	}
	panic(MakeError("PrimNot"))
}

func PrimIf(x, y, z Obj) Obj {
	switch x {
	case True:
		return y
	case False:
		return z
	}
	panic(MakeError("PrimIf"))
}

func PrimEqual(x, y Obj) Obj {
	return equal(x, y)
}

func PrimCons(x, y Obj) Obj {
	return cons(x, y)
}

func PrimIsSymbol(x Obj) Obj {
	if IsSymbol(x) {
		return True

	}
	return False
}

func PrimReadFileAsByteList(x Obj) Obj {
	fileName := mustString(x)
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
	return cdr(ret)
}

func PrimReadFileAsString(x Obj) Obj {
	fileName := mustString(x)
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(MakeError(err.Error()))
	}
	return MakeString(string(buf))
}

func PrimIsInteger(x Obj) Obj {
	if isFixnum(uintptr(unsafe.Pointer(x))) {
		return True

	}
	if *x != scmHeadNumber {
		return False

	}
	f := mustNumber(x)
	if isPreciseInteger(f) {
		return True

	}
	return False
}

var genIdx uint64 = 0

func PrimGenSym(x Obj) Obj {
	sym := mustSymbol(x)
	str := fmt.Sprintf("%s%d", sym.str, genIdx)
	genIdx++
	return MakeSymbol(str)
}

func PrimLoadFile(test bool) func(e *ControlFlow) {
	return func(e *ControlFlow) {
		var evaluator Evaluator = controlFlowAsEvaluator{e}
		if test {
			evaluator = closureEvaluator{}
		}
		path := mustString(e.Get(1))
		res := loadFile(evaluator, true, path)
		e.Return(res)
	}
}

func loadFile(e Evaluator, extended bool, file string) Obj {
	if _, err := os.Stat(file); err != nil {
		return MakeError(err.Error())
	}

	f, err := os.Open(file)
	if err != nil {
		return MakeError(err.Error())
	}
	defer f.Close()
	r := NewSexpReader(f, extended)
	res := loadFileFromReader(e, extended, r)
	if IsError(res) {
		return res
	}
	return MakeString(file)
}

type Evaluator interface {
	Eval(exp Obj) Obj
	Call(f Obj, args ...Obj) Obj
}

var _ Evaluator = controlFlowAsEvaluator{}

type controlFlowAsEvaluator struct {
	*ControlFlow
}

func (c controlFlowAsEvaluator) Eval(exp Obj) Obj {
	return Eval(c.ControlFlow, exp)
}

func (c controlFlowAsEvaluator) Call(f Obj, args ...Obj) Obj {
	return Call(c.ControlFlow, f, args...)
}

var _ Evaluator = closureEvaluator{}

type closureEvaluator struct{}

func (closureEvaluator) Eval(exp Obj) Obj {
	return Neval(exp)
}

func (closureEvaluator) Call(f Obj, args ...Obj) Obj {
	return NCall(f, args...)
}

func loadFileFromReader(e Evaluator, extended bool, r *SexpReader) Obj {
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
			expand := mustSymbol(symMacroExpand).cora
			if expand != Nil {
				exp = e.Call(expand, exp)
			}
		}

		res := e.Eval(exp)
		if *res == scmHeadError {
			return res
		}
	}
	return Nil
}

func primLoadSo(e *ControlFlow) {
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

	f, ok := entry.(*Obj)
	if !ok {
		e.Return(MakeError("plugin Main must be kl.Obj"))
		return
	}
	if !IsNative(*f) {
		e.Return(MakeError("plugin Main must be a scmNative object"))
		return
	}
	n := MustNative(*f)
	if n.require != 0 {
		e.Return(MakeError("plugin Main arity is incorrect"))
		return
	}

	res := Call(e, *f)
	e.Return(res)
}

func readFileAsSexp(e *ControlFlow) {
	filePath := mustString(e.Get(1))
	extend := (e.Get(2) == True)
	f, err := os.Open(filePath)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()

	ret := Nil
	r := NewSexpReader(f, extend)
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

func writeSexpToFile(e *ControlFlow) {
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

func wrapf1(f func(Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp := e.Get(1)
		res := f(tmp)
		e.Return(res)
	}
}

func wrapf2(f func(x, y Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp1 := e.Get(1)
		tmp2 := e.Get(2)
		res := f(tmp1, tmp2)
		e.Return(res)
	}
}

func wrapf3(f func(x, y, z Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp1 := e.Get(1)
		tmp2 := e.Get(2)
		tmp3 := e.Get(3)
		res := f(tmp1, tmp2, tmp3)
		e.Return(res)
	}
}

func wrapf4(f func(x, y, z Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp1 := e.Get(1)
		tmp2 := e.Get(2)
		tmp3 := e.Get(3)
		res := f(tmp1, tmp2, tmp3)
		e.Return(res)
	}
}

func MakePrimitive(name string, arity int, f interface{}) Obj {
	var fn func(*ControlFlow)
	switch arity {
	case 1:
		raw, ok := f.(func(Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		fn = wrapf1(raw)
	case 2:
		raw, ok := f.(func(x, y Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		fn = wrapf2(raw)
	case 3:
		raw, ok := f.(func(x, y, z Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		fn = wrapf3(raw)
	default:
		panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
	}
	tmp := &scmNative{
		scmHead: scmHeadNative,
		name:    name,
		require: arity,
		fn:      fn,
	}
	return Obj(&tmp.scmHead)
}

type tryResult struct {
	e    *ControlFlow
	data Obj
}

func try(e *ControlFlow, f Obj) (res tryResult) {
	savePos := e.pos
	defer func() {
		if err := recover(); err != nil {
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					// Don't forget to recover the calling stack!
					e.pos = savePos
					e.data = e.data[:e.pos]
					res = tryResult{e: e, data: val}
					return
				}
			}
			// Unexpected panic?
			var buf [8192]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Unexpected Panic:", err)
			fmt.Println("Recovered in Try:", ObjString(f))
			fmt.Println(string(buf[:n]))
			// throw the panic again...
			panic(err)
		}
	}()
	val := Call(e, f)
	res = tryResult{e: e, data: val}
	return
}

func try1(e *ControlFlow, f Obj) (res tryResult) {
	saveSP := sp
	// savePos := e.pos
	defer func() {
		if err := recover(); err != nil {
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					// Don't forget to recover the calling stack!
					// e.pos = savePos
					// e.data = e.data[:e.pos]
					sp = saveSP
					stack = stack[:sp]
					res = tryResult{e: e, data: val}
					return
				}
			}
			// Unexpected panic?
			var buf [8192]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Unexpected Panic:", err)
			fmt.Println("Recovered in Try:", ObjString(f))
			fmt.Println(string(buf[:n]))
			// throw the panic again...
			panic(err)
		}
	}()
	// val := Call(e, f)
	val := NCall(f)
	res = tryResult{e: e, data: val}
	return
}

func (t tryResult) Catch(f Obj) Obj {

	if IsError(t.data) {
		return Call(t.e, f, t.data)
	}
	return t.data
}

func (t tryResult) Catch1(f Obj) Obj {
	if IsError(t.data) {
		return NCall(f, t.data)
	}
	return t.data
}

func PrimTryCatch(test bool) func(e *ControlFlow) {
	return func(e *ControlFlow) {
		exp := e.Get(1)
		act := e.Get(2)
		var res Obj
		if test {
			res = try1(e, exp).Catch1(act)
		} else {
			res = try(e, exp).Catch(act)
		}
		e.Return(res)
	}
}

func PrimNS1Set(key, val Obj) Obj {
	sym := mustSymbol(key)
	sym.cora = val
	return val
}

func PrimNS2Set(key, val Obj) Obj {
	sym := mustSymbol(key)
	sym.function = val
	return key
}

func PrimNS3Set(key, val Obj) Obj {
	sym := mustSymbol(key)
	sym.value = val
	return val
}

func PrimNS1Value(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.cora != nil {
		return sym.cora
	}
	panic(fmt.Sprintf("variable %s not bound", sym.str))
}

func PrimNS2Value(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.function != nil {
		return sym.function
	}
	panic(MakeError(fmt.Sprintf("variable %s not bound", sym.str)))
}

func PrimNS3Value(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.value != nil {
		return sym.value
	}
	panic(MakeError(fmt.Sprintf("variable %s not bound", sym.str)))
}

//go:embed init.cora kl.cora
var initFS embed.FS

func CoraInit(e *ControlFlow, test bool) {
	var evaluator Evaluator = controlFlowAsEvaluator{e}
	if test {
		evaluator = closureEvaluator{}
	}
	f, err := initFS.Open("init.cora")
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()
	r := NewSexpReader(f, true)
	res := loadFileFromReader(evaluator, true, r)
	e.Return(res)
}
