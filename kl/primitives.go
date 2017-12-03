package kl

import (
	"fmt"
	"io"
	"os"
	"time"
)

var allPrimitives []*scmPrimitive = []*scmPrimitive{
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "load-file", Required: 1},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "type", Required: 2, Function: typeFunc},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "get-time", Required: 1, Function: getTime},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "eval-kl", Required: 1},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "close", Required: 1, Function: closeStream},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "open", Required: 2, Function: openStream},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "read-byte", Required: 1, Function: primReadByte},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "write-byte", Required: 2, Function: writeByte},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector?", Required: 1, Function: isVector},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "<-address", Required: 2, Function: primVectorGet},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "address->", Required: 3, Function: primVectorSet},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "absvector", Required: 1, Function: primAbsvector},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "str", Required: 1, Function: primStr},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "<=", Required: 2, Function: lessEqual},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: ">=", Required: 2, Function: greatEqual},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "<", Required: 2, Function: lessThan},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: ">", Required: 2, Function: greatThan},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "error-to-string", Required: 1, Function: primErrorToString},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "simple-error", Required: 1, Function: simpleError},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "=", Required: 2, Function: PrimEqual},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "-", Required: 2, Function: primNumberSubtract},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "*", Required: 2, Function: primNumberMultiply},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "/", Required: 2, Function: primNumberDivide},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "+", Required: 2, Function: PrimNumberAdd},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "string->n", Required: 1, Function: primStringToNumber},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "n->string", Required: 1, Function: primNumberToString},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "number?", Required: 1, Function: primIsNumber},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "string?", Required: 1, Function: primIsString},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "pos", Required: 2, Function: pos},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "tlstr", Required: 1, Function: primTailString},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "cn", Required: 2, Function: stringConcat},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "intern", Required: 1, Function: primIntern},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "hd", Required: 1, Function: PrimHead},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "tl", Required: 1, Function: primTail},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "cons", Required: 2, Function: primCons},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "cons?", Required: 1, Function: primIsPair},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "value", Required: 1},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "set", Required: 2},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "not", Required: 1, Function: primNot},
	&scmPrimitive{scmHead: scmHeadPrimitive, Name: "if", Required: 3, Function: primIf},
}

func GetPrimitiveByID(id int) *scmPrimitive {
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
		sym := mustSymbol(args[0])
		return MakeString(sym.sym)
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

func PrimSet(symbolTable map[string]Obj, key Obj, val Obj) Obj {
	sym := mustSymbol(key)
	symbolTable[sym.sym] = val
	return val
}

func PrimValue(symbolTable map[string]Obj, key Obj) Obj {
	sym := mustSymbol(key)
	if val, ok := symbolTable[sym.sym]; ok {
		return val
	}
	return MakeError(fmt.Sprintf("variable %s not bound", sym.sym))
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
	return vec[off]
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
	mode := mustSymbol(args[1]).sym
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
	kind := mustSymbol(args[0]).sym
	switch kind {
	case "unix":
		return MakeInteger(int(time.Now().Unix()))
	case "run":
		return MakeInteger(int(time.Since(time.Now()).Seconds()))
	}
	return Nil
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
