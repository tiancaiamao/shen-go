package kl

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
	"unsafe"
)

var allPrimitives []*ScmPrimitive = []*ScmPrimitive{
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "load-file", Required: 1},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "type", Required: 2, Function: typeFunc},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "get-time", Required: 1, Function: getTime},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "close", Required: 1, Function: closeStream},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "open", Required: 2, Function: openStream},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "read-byte", Required: 1, Function: primReadByte},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "write-byte", Required: 2, Function: writeByte},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector?", Required: 1, Function: isVector},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "<-address", Required: 2, Function: primVectorGet},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "address->", Required: 3, Function: primVectorSet},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector", Required: 1, Function: primAbsvector},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "str", Required: 1, Function: primStr},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "<=", Required: 2, Function: lessEqual},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: ">=", Required: 2, Function: greatEqual},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "<", Required: 2, Function: lessThan},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: ">", Required: 2, Function: greatThan},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "error-to-string", Required: 1, Function: primErrorToString},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "simple-error", Required: 1, Function: simpleError},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "=", Required: 2, Function: PrimEqual},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "-", Required: 2, Function: primNumberSubtract},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "*", Required: 2, Function: primNumberMultiply},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "/", Required: 2, Function: primNumberDivide},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "+", Required: 2, Function: PrimNumberAdd},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "string->n", Required: 1, Function: primStringToNumber},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "n->string", Required: 1, Function: primNumberToString},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "number?", Required: 1, Function: primIsNumber},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "string?", Required: 1, Function: primIsString},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "pos", Required: 2, Function: pos},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "tlstr", Required: 1, Function: primTailString},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "cn", Required: 2, Function: stringConcat},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "intern", Required: 1, Function: primIntern},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "hd", Required: 1, Function: PrimHead},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "tl", Required: 1, Function: primTail},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "cons", Required: 2, Function: primCons},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "cons?", Required: 1, Function: primIsPair},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "value", Required: 1},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "set", Required: 2},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "not", Required: 1, Function: primNot},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "if", Required: 3, Function: primIf},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "symbol?", Required: 1, Function: primIsSymbol},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "native", Required: 9999},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "hash", Required: 2, Function: primHash},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "read-file-as-bytelist", Required: 1, Function: primReadFileAsByteList},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "read-file-as-string", Required: 1, Function: primReadFileAsString},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "variable?", Required: 1, Function: primIsVariable},
	&ScmPrimitive{scmHead: scmHeadPrimitive, Name: "integer?", Required: 1, Function: primIsInteger},
}

var primitiveIdx map[string]*ScmPrimitive

func init() {
	primitiveIdx = make(map[string]*ScmPrimitive)
	for i, prim := range allPrimitives {
		prim.id = i
		primitiveIdx[prim.Name] = prim
	}
}

func NativeIsPrimitive(args ...Obj) Obj {
	if *args[0] != scmHeadSymbol {
		return False
	}
	str := GetSymbol(args[0])
	_, ok := primitiveIdx[str]
	if ok {
		return True
	}
	return False
}

func NativePrimitiveArity(args ...Obj) Obj {
	str := GetSymbol(args[0])
	prim, ok := primitiveIdx[str]
	if !ok {
		return MakeInteger(-1)
	}
	return MakeInteger(prim.Required)
}

func NativePrimitiveID(args ...Obj) Obj {
	str := GetSymbol(args[0])
	prim, ok := primitiveIdx[str]
	if !ok {
		return MakeError("not a primitive")
	}
	return MakeInteger(prim.id)
}

func GetPrimitiveByID(id int) *ScmPrimitive {
	return allPrimitives[id]
}

func PrimNumberAdd(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1.val + y1.val)
}

func primNumberSubtract(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1.val - y1.val)
}

func primNumberMultiply(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1.val * y1.val)
}

func primNumberDivide(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return MakeNumber(x1.val / y1.val)
}

func primIntern(args ...Obj) Obj {
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

func primTail(args ...Obj) Obj {
	return cdr(args[0])
}

func primIsNumber(args ...Obj) Obj {
	if *args[0] == scmHeadNumber {
		return True
	}
	return False
}

func primStringToNumber(args ...Obj) Obj {
	str := mustString(args[0])
	n := ([]rune(str))[0]
	return MakeInteger(int(n))
}

func primNumberToString(args ...Obj) Obj {
	n := mustInteger(args[0])
	return MakeString(string(rune(n)))
}

func primStr(args ...Obj) Obj {
	switch *args[0] {
	case scmHeadPair:
		// Pair may contain recursive list.
		return MakeError("can't str pair object")
	case scmHeadNull:
		return MakeString("()")
	case scmHeadSymbol:
		str := GetSymbol(args[0])
		return MakeString(str)
	case scmHeadNumber:
		f := mustNumber(args[0])
		if !isPreciseInteger(f.val) {
			return MakeString(fmt.Sprintf("%f", f.val))
		}
		return MakeString(fmt.Sprintf("%d", int(f.val)))
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
	default:
		return MakeString("primStr unknown")
	}
	return MakeString("wrong input, the object is not atom ...")
}

func stringConcat(args ...Obj) Obj {
	s1 := mustString(args[0])
	s2 := mustString(args[1])
	return MakeString(s1 + s2)
}

func primTailString(args ...Obj) Obj {
	str := mustString(args[0])
	if len(str) == 0 {
		return MakeError("empty string")
	}
	return MakeString(string([]rune(str)[1:]))
}

func pos(args ...Obj) Obj {
	s := []rune(mustString(args[0]))
	n := mustInteger(args[1])
	if n >= len(s) {
		return MakeError(fmt.Sprintf("%d is not valid index for %s", n, string(s)))
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
	return MakeError(fmt.Sprintf("variable %s not bound", symVal.str))
}

func simpleError(args ...Obj) Obj {
	str := mustString(args[0])
	return MakeError(str)
}

func primErrorToString(args ...Obj) Obj {
	e := mustError(args[0])
	return MakeString(e.err)
}

func greatThan(args ...Obj) Obj {
	x := mustNumber(args[0])
	y := mustNumber(args[1])
	if x.val > y.val {
		return True
	}
	return False
}

func lessThan(args ...Obj) Obj {
	x := mustNumber(args[0])
	y := mustNumber(args[1])
	if x.val < y.val {
		return True
	}
	return False
}

func lessEqual(args ...Obj) Obj {
	x := mustNumber(args[0])
	y := mustNumber(args[1])
	if x.val <= y.val {
		return True
	}
	return False
}

func greatEqual(args ...Obj) Obj {
	x := mustNumber(args[0])
	y := mustNumber(args[1])
	if x.val >= y.val {
		return True
	}
	return False
}

func primAbsvector(args ...Obj) Obj {
	n := mustInteger(args[0])
	if n < 0 {
		return MakeError("absvector wrong argument")
	}
	return MakeVector(n)
}

func primVectorSet(args ...Obj) Obj {
	vec := mustVector(args[0])
	off := mustInteger(args[1])
	val := args[2]
	vec[off] = val
	return args[0]
}

func primVectorGet(args ...Obj) Obj {
	vec := mustVector(args[0])
	off := mustInteger(args[1])
	if off >= len(vec) {
		return MakeError(fmt.Sprintf("index %d out of range %d", off, len(vec)))
	}
	ret := vec[off]
	if ret == nil {
		return undefined
	}
	return ret
}

func isVector(args ...Obj) Obj {
	if *args[0] == scmHeadVector {
		return True
	}
	return False
}

func writeByte(args ...Obj) Obj {
	n := mustInteger(args[0])
	s := mustStream(args[1])
	w, ok := s.raw.(io.Writer)
	if !ok {
		return MakeError("stream is not opened in out mode")
	}
	var b [1]byte
	b[0] = byte(n)
	_, err := w.Write(b[:])
	if err != nil {
		return MakeError(err.Error())
	}
	return args[0]
}

func primReadByte(args ...Obj) Obj {
	s := mustStream(args[0])
	r, ok := s.raw.(io.Reader)
	if !ok {
		return MakeError("stream is closed of not opened in in mode")
	}
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		if err == io.EOF {
			return MakeInteger(-1)
		}
		return MakeError(err.Error())
	}
	return MakeInteger(int(buf[0]))
}

func openStream(args ...Obj) Obj {
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
		return MakeError(err.Error())
	}
	return MakeStream(f)
}

func closeStream(args ...Obj) Obj {
	s := mustStream(args[0])
	c := s.raw.(io.Closer)
	c.Close()
	return Nil
}

func getTime(args ...Obj) Obj {
	kind := GetSymbol(args[0])
	switch kind {
	case "unix":
		return MakeNumber(float64(time.Now().Unix()))
	case "run":
		return MakeNumber(time.Since(uptime).Seconds())
	}
	return MakeError(fmt.Sprintf("get-time does not understand the parameter %s", kind))
}

func typeFunc(args ...Obj) Obj {
	// TODO: seems meanless
	return args[0]
}

func primIsString(args ...Obj) Obj {
	if *args[0] == scmHeadString {
		return True
	}
	return False
}

func primIsPair(args ...Obj) Obj {
	if *args[0] == scmHeadPair {
		return True
	}
	return False
}

func primNot(args ...Obj) Obj {
	if args[0] == False {
		return True
	} else if args[0] == True {
		return False
	}
	return MakeError("primNot")

}

func primIf(args ...Obj) Obj {
	switch args[0] {
	case True:
		return args[1]
	case False:
		return args[2]
	}
	return MakeError("primIf")
}

func PrimEqual(args ...Obj) Obj {
	return equal(args[0], args[1])
}

func primCons(args ...Obj) Obj {
	return cons(args[0], args[1])
}

func primIsSymbol(args ...Obj) Obj {
	if IsSymbol(args[0]) {
		return True
	}
	return False
}

// A --> number --> number
func primHash(args ...Obj) Obj {
	mod := mustNumber(args[1]).val
	ret := objHash(args[0]) % int(mod)
	return MakeInteger(ret)
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
	case scmHeadNumber:
		sum = sum<<23 + *((*int)(unsafe.Pointer(&mustNumber(x).val)))
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

func primReadFileAsByteList(args ...Obj) Obj {
	fileName := mustString(args[0])
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return MakeError(err.Error())
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

func primReadFileAsString(args ...Obj) Obj {
	fileName := mustString(args[0])
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return MakeError(err.Error())
	}

	return MakeString(string(buf))
}

func primIsVariable(args ...Obj) Obj {
	if *args[0] != scmHeadSymbol {
		return False
	}

	sym := GetSymbol(args[0])
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		return False
	}
	return True
}

func primIsInteger(args ...Obj) Obj {
	if *args[0] != scmHeadNumber {
		return False
	}
	f := mustNumber(args[0]).val
	if isPreciseInteger(f) {
		return True
	}
	return False
}
