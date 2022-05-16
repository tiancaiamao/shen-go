package re

import (
	// "embed"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"time"
)

var klPrimitives = []struct {
	name  string
	arity int
	fn    interface{}
}{
	{"get-time", 1, PrimGetTime},
	{"close", 1, PrimCloseStream},
	{"open", 2, PrimOpenStream},
	{"read-byte", 1, PrimReadByte},
	{"write-byte", 2, PrimWriteByte},
	{"absvector?", 1, PrimIsVector},
	{"<-address", 2, PrimVectorGet},
	{"address->", 3, PrimVectorSet},
	{"absvector", 1, PrimAbsvector},
	{"str", 1, PrimStr},
	{"<=", 2, PrimLessEqual},
	{">=", 2, PrimGreatEqual},
	{"<", 2, PrimLessThan},
	{">", 2, PrimGreatThan},
	{"error-to-string", 1, PrimErrorToString},
	{"simple-error", 1, PrimSimpleError},
	{"=", 2, PrimEqual},
	//	{"-", 2, PrimNumberSubtract},
	//	{"*", 2, PrimNumberMultiply},
	//	{"/", 2, PrimNumberDivide},
	//	{"+", 2, PrimNumberAdd},
	{"string->n", 1, PrimStringToNumber},
	{"n->string", 1, PrimNumberToString},
	{"number?", 1, PrimIsNumber},
	{"string?", 1, PrimIsString},
	{"pos", 2, PrimPos},
	{"tlstr", 1, PrimTailString},
	{"cn", 2, PrimStringConcat},
	{"intern", 1, PrimIntern},
	{"hd", 1, PrimHead},
	{"tl", 1, PrimTail},
	{"cons", 2, PrimCons},
	{"cons?", 1, PrimIsPair},
	//	{"value", 1, PrimNS3Value},
	//	{"set", 2, PrimNS3Set},
	{"not", 1, PrimNot},
	{"if", 3, PrimIf},
	{"symbol?", 1, PrimIsSymbol},
	{"read-file-as-bytelist", 1, PrimReadFileAsByteList},
	{"read-file-as-string", 1, PrimReadFileAsString},
	{"variable?", 1, PrimIsVariable},
	{"integer?", 1, PrimIsInteger},
	{"shen.char-stoutput?", 1, PrimCharStOutput},
	{"shen.char-stinput?", 1, PrimCharStInput},
}

func PrimCharStInput(x Obj) Obj {
	return False
}

func PrimCharStOutput(x Obj) Obj {
	return False
}

/*
func PrimLoad(e *VM) {
	file := e.Get(1)
	if !IsString(file) {
		e.Return(MakeError("arg1 must be string"))
		return
	}
	path := GetString(file)
	if _, err := os.Stat(path); err != nil {
		e.Return(MakeError(err.Error()))
		return
	}

	f, err := os.Open(path)
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()

	r := NewSexpReader(f, false)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				e.Return(MakeError(err.Error()))
				return
			}
			break
		}

		res := e.Eval(exp)
		if IsError(res) {
			e.Return(res)
			return
		}
	}
	e.Return(MakeSymbol("loaded"))
}
*/

/*
func PrimNS3Value(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.value != nil {
		return sym.value
	}
	fmt.Println("variable not bound", sym.str)
	panic(MakeError(fmt.Sprintf("variable %s not bound", sym.str)))
}



*/

func init() {
	//	PrimNS2Set(MakeSymbol("try-catch"), MakeNative(PrimTryCatch(false), 2))
	//	PrimNS2Set(MakeSymbol("load-file"), MakeNative(PrimLoad, 1))

	for _, item := range klPrimitives {
		sym := MakeSymbol(item.name)
		sym.fn = MakePrimitive(item.name, item.arity, item.fn)
	}
	MakeSymbol("eval-kl").fn = &Closure{
		Required: 1,
		Code: instrFunc(func(vm *VM) {
			exp := vm.stack.Get(2)
			res := eval(vm, exp)
			vm.Return(res)
		}),
	}
	// Overload for Primitive set and value.
	MakeSymbol("*stinput*").val = Stream{os.Stdin}
	MakeSymbol("*stoutput*").val = Stream{os.Stdout}
	dir, _ := os.Getwd()
	MakeSymbol("*home-directory*").val = String(dir)
	MakeSymbol("*release*").val = String(runtime.Version())
	MakeSymbol("*os*").val = String(runtime.GOOS)
	MakeSymbol("*language*").val = String("Go")
	MakeSymbol("*implementation*").val = String("AOT+closure")
	MakeSymbol("*porters*").val = String("Author Mao")
	MakeSymbol("*port*").val = String("1.1.0")
}

/*
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
*/

func PrimIntern(o Obj) Obj {
	str := o.(String)
	switch str {
	case "true":
		return True
	case "false":
		return False
	}
	return MakeSymbol(string(str))
}

func PrimHead(o Obj) Obj {
	return car(o)
}

func PrimTail(o Obj) Obj {
	return cdr(o)
}

func PrimIsNumber(o Obj) Obj {
	if _, ok := o.(Integer); ok {
		return True
	}
	return False
}

func PrimStringToNumber(o Obj) Obj {
	str := o.(String)
	n := ([]rune(str))[0]
	return Integer(int(n))
}

func PrimNumberToString(o Obj) Obj {
	n := o.(Integer)
	return String(rune(n))
}

func PrimStringConcat(x, y Obj) Obj {
	s1 := x.(String)
	s2 := y.(String)
	return String(s1 + s2)
}

func PrimTailString(o Obj) Obj {
	str := o.(String)
	if len(str) == 0 {
		panic("empty string")
	}
	return String(string([]rune(str)[1:]))
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
	x1 := x.(Integer)
	y1 := y.(Integer)
	if x1 > y1 {
		return True

	}
	return False
}

func PrimLessThan(a, b Obj) Obj {
	x := a.(Integer)
	y := b.(Integer)
	if x < y {
		return True

	}
	return False
}

func PrimLessEqual(a, b Obj) Obj {
	x := a.(Integer)
	y := b.(Integer)
	if x <= y {
		return True

	}
	return False

}

func PrimGreatEqual(a, b Obj) Obj {
	x := a.(Integer)
	y := b.(Integer)
	if x >= y {
		return True

	}
	return False
}

func PrimAbsvector(o Obj) Obj {
	n := o.(Integer)
	if n < 0 {
		panic("absvector wrong argument")
	}
	return Vector(make([]Obj, n))
}

func PrimVectorSet(x, y, z Obj) Obj {
	vec := x.(Vector)
	off := y.(Integer)
	val := z
	vec[off] = val
	return x
}

func PrimVectorGet(x, y Obj) Obj {
	vec := x.(Vector)
	off := y.(Integer)
	if int(off) >= len(vec) {
		panic(fmt.Sprintf("index %d out of range %d", off, len(vec)))
	}
	ret := vec[off]
	if ret == nil {
		return nil
	}
	return ret
}

func PrimIsVector(x Obj) Obj {
	if _, ok := x.(Vector); ok {
		return True
	}
	return False
}

func PrimWriteByte(x, y Obj) Obj {
	n := x.(Integer)
	w, ok := y.(Stream)
	if !ok {
		panic("stream is not opened in out mode")
	}
	var b [1]byte
	b[0] = byte(n)
	_, err := w.Write(b[:])
	if err != nil {
		panic(err.Error())
	}
	return x
}

func PrimReadByte(x Obj) Obj {
	r, ok := x.(Stream)
	if !ok {
		panic("stream is closed of not opened in in mode")
	}
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		if err == io.EOF {
			return Integer(-1)
		}
		panic(err.Error())
	}
	return Integer(buf[0])
}

func PrimOpenStream(x, y Obj) Obj {
	file := x.(String)
	var flag int
	mode := y.(*Symbol).str
	switch mode {
	case "in":
		flag |= os.O_RDONLY
	case "out":
		flag |= os.O_WRONLY | os.O_CREATE
	default:
		flag = os.O_RDWR | os.O_CREATE
	}

	f, err := os.OpenFile(string(file), flag, 0666)
	if err != nil {
		panic(err.Error())
	}
	return Stream{f}
}

func PrimCloseStream(x Obj) Obj {
	s := x.(Stream)
	if c, ok := s.ReadWriter.(io.Closer); ok {
		c.Close()
	}
	return Nil
}

func PrimGetTime(x Obj) Obj {
	kind := x.(*Symbol).str
	switch kind {
	case "unix":
		return MakeNumber(float64(time.Now().Unix()))
	case "run":
		return MakeNumber(time.Since(uptime).Seconds())

	}
	panic(fmt.Sprintf("get-time does not understand the parameter %s", kind))
}

func PrimIsString(x Obj) Obj {
	if _, ok := x.(String); ok {
		return True

	}
	return False
}

func PrimIsPair(x Obj) Obj {
	if _, ok := x.(*Cons); ok {
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
	panic("PrimNot")
}

func PrimIf(x, y, z Obj) Obj {
	switch x {
	case True:
		return y
	case False:
		return z
	}
	panic("PrimIf")
}

func PrimEqual(x, y Obj) Obj {
	if x == y {
		return True
	}
	return False
}

func PrimCons(x, y Obj) Obj {
	return cons(x, y)
}

func PrimIsSymbol(x Obj) Obj {
	if _, ok := x.(*Symbol); ok {
		return True

	}
	return False
}

func PrimReadFileAsByteList(x Obj) Obj {
	fileName := x.(String)
	buf, err := ioutil.ReadFile(string(fileName))
	if err != nil {
		panic(err.Error())
	}

	ret := cons(Nil, Nil)
	curr := ret.(*Cons)
	for _, b := range buf {
		tmp := cons(Integer(int(b)), Nil)
		curr.cdr = tmp
		curr = tmp.(*Cons)
	}
	return cdr(ret)
}

func PrimReadFileAsString(x Obj) Obj {
	fileName := x.(String)
	buf, err := ioutil.ReadFile(string(fileName))
	if err != nil {
		panic(err.Error())
	}
	return String(buf)
}

func PrimIsInteger(x Obj) Obj {
	if _, ok := x.(Integer); ok {
		return True
	}
	return False
}

var genIdx uint64 = 0

/*
func PrimGenSym(x Obj) Obj {
	sym := mustSymbol(x)
	str := fmt.Sprintf("%s%d", sym.str, genIdx)
	genIdx++
	return MakeSymbol(str)
}


func PrimLoadFile() func(e *VM) {
	return func(e *VM) {
		path := mustString(e.Get(1))
		res := loadFile(true, path)
		e.Return(res)
	}
}

func loadFile(extended bool, file string) Obj {
	if _, err := os.Stat(file); err != nil {
		return MakeError(err.Error())
	}

	f, err := os.Open(file)
	if err != nil {
		return MakeError(err.Error())
	}
	defer f.Close()
	r := NewSexpReader(f, extended)
	res := loadFileFromReader(extended, r)
	if IsError(res) {
		return res
	}
	return MakeString(file)
}


func loadFileFromReader(extended bool, r *SexpReader) Obj {
	var ctx VM
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return Error{err}
			}
			break
		}


		// Macro expand for cora.
		if extended {
			expand := mustSymbol(symMacroExpand).cora
			if expand != Nil {
				exp = ctx.Call(expand, exp)
			}
		}


		res := ctx.Eval(exp)
		if *res == scmHeadError {
			return res
		}
	}
	return Nil
}

func readFileAsSexp(e *VM) {
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


func writeSexpToFile(e *VM) {
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
*/

func wrapf1(f func(Obj) Obj) *Closure {
	return &Closure{
		Required: 1,
		Code: instrFunc(func(e *VM) {
			tmp := e.stack.Get(2)
			res := f(tmp)
			e.Return(res)
		}),
	}
}

func wrapf2(f func(x, y Obj) Obj) *Closure {
	return &Closure{
		Required: 2,
		Code: instrFunc(func(e *VM) {
			tmp1 := e.stack.Get(2)
			tmp2 := e.stack.Get(3)
			res := f(tmp1, tmp2)
			e.Return(res)
		}),
	}
}

func wrapf3(f func(x, y, z Obj) Obj) *Closure {
	return &Closure{
		Required: 3,
		Code: instrFunc(func(e *VM) {
			tmp1 := e.stack.Get(2)
			tmp2 := e.stack.Get(3)
			tmp3 := e.stack.Get(4)
			res := f(tmp1, tmp2, tmp3)
			e.Return(res)
		}),
	}
}

func wrapf4(f func(x, y, z Obj) Obj) *Closure {
	return &Closure{
		Required: 4,
		Code: instrFunc(func(e *VM) {
			tmp1 := e.stack.Get(2)
			tmp2 := e.stack.Get(3)
			tmp3 := e.stack.Get(4)
			res := f(tmp1, tmp2, tmp3)
			e.Return(res)
		}),
	}
}

func MakePrimitive(name string, arity int, f interface{}) *Closure {
	switch arity {
	case 1:
		raw, ok := f.(func(Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		return wrapf1(raw)
	case 2:
		raw, ok := f.(func(x, y Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		return wrapf2(raw)
	case 3:
		raw, ok := f.(func(x, y, z Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		return wrapf3(raw)
	}
	panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
}

func PrimStr(o Obj) Obj {
	return String(o.String())
}

func PrimErrorToString(o Obj) Obj {
	return String(o.String())
}

func PrimSimpleError(o Obj) Obj {
	str := o.(String)
	return Error{errors.New(string(str))}
}

func PrimPos(x, y Obj) Obj {
	s := []rune(x.(String))
	n := y.(Integer)
	if int(n) >= len(s) {
		panic(fmt.Sprintf("%d is not valid index for %s", n, string(s)))
	}
	return String(string([]rune(s)[n]))
}

func PrimIsVariable(x Obj) Obj {
	if _, ok := x.(*Symbol); !ok {
		return False
	}

	sym := x.(*Symbol).str
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		return False
	}
	return True
}
