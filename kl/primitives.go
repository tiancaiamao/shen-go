package kl

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
	"unsafe"
)

var AllPrimitives = []*ScmPrimitive{
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "load-file", Required: 1, CodeGen: "PrimLoadFile"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "type", Required: 2, Function: PrimTypeFunc, CodeGen: "PrimTypeFunc"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "get-time", Required: 1, Function: PrimGetTime, CodeGen: "PrimGetTime"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1, CodeGen: "PrimEvalKL"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "close", Required: 1, Function: PrimCloseStream, CodeGen: "PrimCloseStream"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "open", Required: 2, Function: PrimOpenStream, CodeGen: "PrimOpenStream"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "read-byte", Required: 1, Function: PrimReadByte, CodeGen: "PrimReadByte"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "write-byte", Required: 2, Function: PrimWriteByte, CodeGen: "PrimWriteByte"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector?", Required: 1, Function: PrimIsVector, CodeGen: "PrimIsVector"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "<-address", Required: 2, Function: PrimVectorGet, CodeGen: "PrimVectorGet"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "address->", Required: 3, Function: PrimVectorSet, CodeGen: "PrimVectorSet"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector", Required: 1, Function: PrimAbsvector, CodeGen: "PrimAbsvector"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "str", Required: 1, Function: PrimStr, CodeGen: "PrimStr"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "<=", Required: 2, Function: PrimLessEqual, CodeGen: "PrimLessEqual"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: ">=", Required: 2, Function: PrimGreatEqual, CodeGen: "PrimGreatEqual"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "<", Required: 2, Function: PrimLessThan, CodeGen: "PrimLessThan"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: ">", Required: 2, Function: PrimGreatThan, CodeGen: "PrimGreatThan"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "error-to-string", Required: 1, Function: PrimErrorToString, CodeGen: "PrimErrorToString"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "simple-error", Required: 1, Function: PrimSimpleError, CodeGen: "PrimSimpleError"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "=", Required: 2, Function: PrimEqual, CodeGen: "PrimEqual"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "-", Required: 2, Function: PrimNumberSubtract, CodeGen: "PrimNumberSubtract"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "*", Required: 2, Function: PrimNumberMultiply, CodeGen: "PrimNumberMultiply"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "/", Required: 2, Function: PrimNumberDivide, CodeGen: "PrimNumberDivide"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "+", Required: 2, Function: PrimNumberAdd, CodeGen: "PrimNumberAdd"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "string->n", Required: 1, Function: PrimStringToNumber, CodeGen: "PrimStringToNumber"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "n->string", Required: 1, Function: PrimNumberToString, CodeGen: "PrimNumberToString"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "number?", Required: 1, Function: PrimIsNumber, CodeGen: "PrimIsNumber"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "string?", Required: 1, Function: PrimIsString, CodeGen: "PrimIsString"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "pos", Required: 2, Function: PrimPos, CodeGen: "PrimPos"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "tlstr", Required: 1, Function: PrimTailString, CodeGen: "PrimTailString"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "cn", Required: 2, Function: PrimStringConcat, CodeGen: "PrimStringConcat"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "intern", Required: 1, Function: PrimIntern, CodeGen: "PrimIntern"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "hd", Required: 1, Function: PrimHead, CodeGen: "PrimHead"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "tl", Required: 1, Function: PrimTail, CodeGen: "PrimTail"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "cons", Required: 2, Function: PrimCons, CodeGen: "PrimCons"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "cons?", Required: 1, Function: PrimIsPair, CodeGen: "PrimIsPair"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "value", Required: 1, CodeGen: "PrimValue"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "set", Required: 2, CodeGen: "PrimSet"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "not", Required: 1, Function: PrimNot, CodeGen: "PrimNot"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "if", Required: 3, Function: PrimIf, CodeGen: "PrimIf"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "symbol?", Required: 1, Function: PrimIsSymbol, CodeGen: "PrimIsSymbol"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "read-file-as-bytelist", Required: 1, Function: PrimReadFileAsByteList, CodeGen: "PrimReadFileAsByteList"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "read-file-as-string", Required: 1, Function: PrimReadFileAsString, CodeGen: "PrimReadFileAsString"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "variable?", Required: 1, Function: PrimIsVariable, CodeGen: "PrimIsVariable"},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "integer?", Required: 1, Function: PrimIsInteger, CodeGen: "PrimIsInteger"},
}

func PrimNumberAdd(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1 + y1)
}

func PrimNumberSubtract(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1 - y1)
}

func PrimNumberMultiply(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1 * y1)
}

func PrimNumberDivide(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1 / y1)
}

func PrimIntern(args ...Obj) Obj {
	str := mustString(args[0])
	switch str {
	case "true":
		return True
	case "false":
		return False
	}
	return MakeSymbol(str)
}

func PrimHead(args ...Obj) Obj {
	return car(args[0])
}

func PrimTail(args ...Obj) Obj {
	return cdr(args[0])
}

func PrimIsNumber(args ...Obj) Obj {
	if *args[0] == scmHeadNumber {
		return True
	}
	return False
}

func PrimStringToNumber(args ...Obj) Obj {
	str := mustString(args[0])
	n := ([]rune(str))[0]
	return MakeInteger(int(n))
}

func PrimNumberToString(args ...Obj) Obj {
	n := mustInteger(args[0])
	return MakeString(string(rune(n)))
}

func PrimStr(args ...Obj) Obj {
	switch *args[0] {
	case scmHeadPair:
		// Pair may contain recursive list.
		panic(MakeError("can't str pair object"))
	case scmHeadNull:
		return MakeString("()")
	case scmHeadSymbol:
		str := GetSymbol(args[0])
		return MakeString(str)
	case scmHeadNumber:
		if isFixnum(args[0]) {
			return MakeString(fmt.Sprintf("%d", mustInteger(args[0])))
		}
		f := mustNumber(args[0])
		if !isPreciseInteger(f) {
			return MakeString(fmt.Sprintf("%f", f))
		}
		return MakeString(fmt.Sprintf("%d", int(f)))
	case scmHeadString:
		return MakeString(fmt.Sprintf(`"%s"`, mustString(args[0])))
	case scmHeadProcedure:
		return MakeString("#<procedure>")
	case scmHeadPrimitive:
		prim := mustPrimitive(args[0])
		return MakeString(fmt.Sprintf("#<primitive %s>", prim.Name))
	case scmHeadBoolean:
		if args[0] == True {
			return MakeString("true")
		} else if args[0] == False {
			return MakeString("false")
		}
	case scmHeadError:
		e := mustError(args[0])
		return MakeString(e.err)
	case scmHeadStream:
		return MakeString("<stream>")
	case scmHeadRaw:
		return MakeString("#<raw>")
	case scmHeadNative:
		return MakeString("#<native>")
	default:
		return MakeString("PrimStr unknown")
	}
	return MakeString("wrong input, the object is not atom ...")
}

func PrimStringConcat(args ...Obj) Obj {
	s1 := mustString(args[0])
	s2 := mustString(args[1])
	return MakeString(s1 + s2)
}

func PrimTailString(args ...Obj) Obj {
	str := mustString(args[0])
	if len(str) == 0 {
		panic(MakeError("empty string"))
	}
	return MakeString(string([]rune(str)[1:]))
}

func PrimPos(args ...Obj) Obj {
	s := []rune(mustString(args[0]))
	n := mustInteger(args[1])
	if n >= len(s) {
		panic(MakeError(fmt.Sprintf("%d is not valid index for %s", n, string(s))))
	}
	return MakeString(string([]rune(s)[n]))
}

func and(args ...Obj) Obj {
	if args[0] == True && args[1] == True {
		return True
	}
	return False
}

func or(args ...Obj) Obj {
	if args[0] == False || args[1] == False {
		return False
	}
	return True
}

func PrimSet(key Obj, val Obj) Obj {
	sym := mustSymbol(key)
	symVal := &symbolArray[sym.offset]
	symVal.value = val
	return val
}

func PrimValue(key Obj) Obj {
	sym := mustSymbol(key)
	symVal := &symbolArray[sym.offset]
	if symVal.value != nil {
		return symVal.value
	}
	panic(MakeError(fmt.Sprintf("variable %s not bound", symVal.str)))
}

func PrimSimpleError(args ...Obj) Obj {
	str := mustString(args[0])
	panic(MakeError(str))
}

func PrimErrorToString(args ...Obj) Obj {
	e := mustError(args[0])
	return MakeString(e.err)
}

func PrimGreatThan(args ...Obj) Obj {
	x := mustInteger(args[0])
	y := mustInteger(args[1])
	if x > y {
		return True
	}
	return False
}

func PrimLessThan(args ...Obj) Obj {
	x := mustInteger(args[0])
	y := mustInteger(args[1])
	if x < y {
		return True
	}
	return False
}

func PrimLessEqual(args ...Obj) Obj {
	x := mustInteger(args[0])
	y := mustInteger(args[1])
	if x <= y {
		return True
	}
	return False
}

func PrimGreatEqual(args ...Obj) Obj {
	x := mustInteger(args[0])
	y := mustInteger(args[1])
	if x >= y {
		return True
	}
	return False
}

func PrimAbsvector(args ...Obj) Obj {
	n := mustInteger(args[0])
	if n < 0 {
		panic(MakeError("absvector wrong argument"))
	}
	return MakeVector(n)
}

func PrimVectorSet(args ...Obj) Obj {
	vec := mustVector(args[0])
	off := mustInteger(args[1])
	val := args[2]
	vec[off] = val
	return args[0]
}

func PrimVectorGet(args ...Obj) Obj {
	vec := mustVector(args[0])
	off := mustInteger(args[1])
	if off >= len(vec) {
		panic(MakeError(fmt.Sprintf("index %d out of range %d", off, len(vec))))
	}
	ret := vec[off]
	if ret == nil {
		return undefined
	}
	return ret
}

func PrimIsVector(args ...Obj) Obj {
	if *args[0] == scmHeadVector {
		return True
	}
	return False
}

func PrimWriteByte(args ...Obj) Obj {
	n := mustInteger(args[0])
	s := mustStream(args[1])
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
	return args[0]
}

func PrimReadByte(args ...Obj) Obj {
	s := mustStream(args[0])
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

func PrimOpenStream(args ...Obj) Obj {
	file := mustString(args[0])
	var flag int
	mode := GetSymbol(args[1])
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

func PrimCloseStream(args ...Obj) Obj {
	s := mustStream(args[0])
	c := s.raw.(io.Closer)
	c.Close()
	return Nil
}

func PrimGetTime(args ...Obj) Obj {
	kind := GetSymbol(args[0])
	switch kind {
	case "unix":
		return MakeNumber(float64(time.Now().Unix()))
	case "run":
		return MakeNumber(time.Since(uptime).Seconds())
	}
	panic(MakeError(fmt.Sprintf("get-time does not understand the parameter %s", kind)))
}

func PrimTypeFunc(args ...Obj) Obj {
	// TODO: seems meanless
	return args[0]
}

func PrimIsString(args ...Obj) Obj {
	if *args[0] == scmHeadString {
		return True
	}
	return False
}

func PrimIsPair(args ...Obj) Obj {
	if *args[0] == scmHeadPair {
		return True
	}
	return False
}

func PrimNot(args ...Obj) Obj {
	if args[0] == False {
		return True
	} else if args[0] == True {
		return False
	}
	panic(MakeError("PrimNot"))

}

func PrimIf(args ...Obj) Obj {
	switch args[0] {
	case True:
		return args[1]
	case False:
		return args[2]
	}
	panic(MakeError("PrimIf"))
}

func PrimEqual(args ...Obj) Obj {
	return equal(args[0], args[1])
}

func PrimCons(args ...Obj) Obj {
	return cons(args[0], args[1])
}

func PrimIsSymbol(args ...Obj) Obj {
	if IsSymbol(args[0]) {
		return True
	}
	return False
}

func objHash(x Obj) int {
	// initialize value is its type, then mixed with value part.
	sum := (int)(*x)
	switch *x {
	case scmHeadNull:
	case scmHeadBoolean:
		if x == True {
			sum = sum<<23 + 17
		} else {
			sum = sum<<23 + 13
		}
	// case scmHeadNumber:
	// 	sum = sum<<23 + *((*int)(unsafe.Pointer(&mustNumber(x).val)))
	case scmHeadString:
		str := mustString(x)
		for _, s := range str {
			sum = ((sum + 13) << 7) + int(s)
		}
	case scmHeadSymbol:
		str := GetSymbol(x)
		for _, s := range str {
			sum = ((sum + 13) << 7) + int(s)
		}
	case scmHeadVector:
		objs := mustVector(x)
		for _, o := range objs {
			sum = ((sum + 17) << 13) + objHash(o)
		}
	case scmHeadError:
		str := mustError(x).err
		for _, s := range str {
			sum = ((sum + 123) << 7) + int(s)
		}
	case scmHeadPair:
		sum = objHash(car(x)) ^ objHash(cdr(x))
	case scmHeadProcedure:
		sum = sum ^ *((*int)(unsafe.Pointer(mustProcedure(x))))
	case scmHeadStream:
		sum = sum ^ *((*int)(unsafe.Pointer(mustStream(x))))
	case scmHeadPrimitive:
		for _, s := range mustPrimitive(x).Name {
			sum = ((sum + 17) << 7) + int(s)
		}
	case scmHeadRaw:
		sum = sum ^ *((*int)(unsafe.Pointer(mustStream(x))))
	}
	return sum
}

func PrimReadFileAsByteList(args ...Obj) Obj {
	fileName := mustString(args[0])
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

func PrimReadFileAsString(args ...Obj) Obj {
	fileName := mustString(args[0])
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(MakeError(err.Error()))
	}

	return MakeString(string(buf))
}

func PrimIsVariable(args ...Obj) Obj {
	if *args[0] != scmHeadSymbol {
		return False
	}

	sym := GetSymbol(args[0])
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		return False
	}
	return True
}

func PrimIsInteger(args ...Obj) Obj {
	if *args[0] != scmHeadNumber {
		return False
	}
	if isFixnum(args[0]) {
		return True
	}
	f := mustNumber(args[0])
	if isPreciseInteger(f) {
		return True
	}
	return False
}

func PrimEvalKL(e *Evaluator, args ...Obj) Obj {
	// fmt.Println("eval-kl:", ObjString(args[0]))
	return e.trampoline(args[0], nil)
}
