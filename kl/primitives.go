package kl

import (
	"fmt"
	"io"
	"os"
	"time"
)

var Primitives []*ScmPrimitive = []*ScmPrimitive{
	&ScmPrimitive{scmHead: Primitive, Name: "load-file", Required: 1},
	&ScmPrimitive{scmHead: Primitive, Name: "type", Required: 2, Function: typeFunc},
	&ScmPrimitive{scmHead: Primitive, Name: "get-time", Required: 1, Function: getTime},
	&ScmPrimitive{scmHead: Primitive, Name: "eval-kl", Required: 1},
	&ScmPrimitive{scmHead: Primitive, Name: "close", Required: 1, Function: closeStream},
	&ScmPrimitive{scmHead: Primitive, Name: "open", Required: 2, Function: openStream},
	&ScmPrimitive{scmHead: Primitive, Name: "read-byte", Required: 1, Function: primReadByte},
	&ScmPrimitive{scmHead: Primitive, Name: "write-byte", Required: 2, Function: writeByte},
	&ScmPrimitive{scmHead: Primitive, Name: "absvector?", Required: 1, Function: isVector},
	&ScmPrimitive{scmHead: Primitive, Name: "<-address", Required: 2, Function: vectorGet},
	&ScmPrimitive{scmHead: Primitive, Name: "address->", Required: 3, Function: vectorSet},
	&ScmPrimitive{scmHead: Primitive, Name: "absvector", Required: 1, Function: primAbsvector},
	&ScmPrimitive{scmHead: Primitive, Name: "str", Required: 1, Function: primStr},
	&ScmPrimitive{scmHead: Primitive, Name: "<=", Required: 2, Function: lessEqual},
	&ScmPrimitive{scmHead: Primitive, Name: ">=", Required: 2, Function: greatEqual},
	&ScmPrimitive{scmHead: Primitive, Name: "<", Required: 2, Function: lessThan},
	&ScmPrimitive{scmHead: Primitive, Name: ">", Required: 2, Function: greatThan},
	&ScmPrimitive{scmHead: Primitive, Name: "error-to-string", Required: 1, Function: primErrorToString},
	&ScmPrimitive{scmHead: Primitive, Name: "simple-error", Required: 1, Function: simpleError},
	&ScmPrimitive{scmHead: Primitive, Name: "=", Required: 2, Function: PrimEqual},
	&ScmPrimitive{scmHead: Primitive, Name: "-", Required: 2, Function: primNumberSubtract},
	&ScmPrimitive{scmHead: Primitive, Name: "*", Required: 2, Function: primNumberMultiply},
	&ScmPrimitive{scmHead: Primitive, Name: "/", Required: 2, Function: primNumberDivide},
	&ScmPrimitive{scmHead: Primitive, Name: "+", Required: 2, Function: PrimNumberAdd},
	&ScmPrimitive{scmHead: Primitive, Name: "string->n", Required: 1, Function: primStringToNumber},
	&ScmPrimitive{scmHead: Primitive, Name: "n->string", Required: 1, Function: primNumberToString},
	&ScmPrimitive{scmHead: Primitive, Name: "number?", Required: 1, Function: primIsNumber},
	&ScmPrimitive{scmHead: Primitive, Name: "string?", Required: 1, Function: primIsString},
	&ScmPrimitive{scmHead: Primitive, Name: "pos", Required: 2, Function: pos},
	&ScmPrimitive{scmHead: Primitive, Name: "tlstr", Required: 1, Function: primTailString},
	&ScmPrimitive{scmHead: Primitive, Name: "cn", Required: 2, Function: stringConcat},
	&ScmPrimitive{scmHead: Primitive, Name: "intern", Required: 1, Function: primIntern},
	&ScmPrimitive{scmHead: Primitive, Name: "hd", Required: 1, Function: PrimHead},
	&ScmPrimitive{scmHead: Primitive, Name: "tl", Required: 1, Function: primTail},
	&ScmPrimitive{scmHead: Primitive, Name: "cons", Required: 2, Function: primCons},
	&ScmPrimitive{scmHead: Primitive, Name: "cons?", Required: 1, Function: primIsPair},
	&ScmPrimitive{scmHead: Primitive, Name: "value", Required: 1},
	&ScmPrimitive{scmHead: Primitive, Name: "set", Required: 2},
	&ScmPrimitive{scmHead: Primitive, Name: "not", Required: 1, Function: primNot},
	&ScmPrimitive{scmHead: Primitive, Name: "if", Required: 3, Function: primIf},
}

func PrimNumberAdd(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return Make_number(x1.val + y1.val)
}

func primNumberSubtract(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return Make_number(x1.val - y1.val)
}

func primNumberMultiply(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return Make_number(x1.val * y1.val)
}

func primNumberDivide(args ...Obj) Obj {
	x1 := mustNumber(args[0])
	y1 := mustNumber(args[1])
	return Make_number(x1.val / y1.val)
}

func primIntern(args ...Obj) Obj {
	str := mustString(args[0])
	switch str {
	case "true":
		return True
	case "false":
		return False
	}
	return Make_symbol(str)
}

func PrimHead(args ...Obj) Obj {
	return car(args[0])
}

func primTail(args ...Obj) Obj {
	return cdr(args[0])
}

func primIsNumber(args ...Obj) Obj {
	if *args[0] == Number {
		return True
	}
	return False
}

func primStringToNumber(args ...Obj) Obj {
	str := mustString(args[0])
	n := ([]rune(str))[0]
	return Make_integer(int(n))
}

func primNumberToString(args ...Obj) Obj {
	n := mustInteger(args[0])
	return Make_string(string(rune(n)))
}

func primStr(args ...Obj) Obj {
	switch *args[0] {
	// TODO: Pair
	case Symbol:
		sym := mustSymbol(args[0])
		return Make_string(sym.sym)
	case Number:
		f := mustNumber(args[0])
		return Make_string(fmt.Sprintf("%f", f.val))
	case String:
		return Make_string(fmt.Sprintf(`"%s"`, mustString(args[0])))
	case Procedure:
		return Make_string("#<procedure>")
	case Primitive:
		return Make_string("#<primitive>")
	case Boolean:
		if args[0] == True {
			return Make_string("true")
		} else if args[0] == False {
			return Make_string("false")
		}
	case Error:
		e := mustError(args[0])
		return Make_string(e.err)
	case Stream:
		return Make_string("<stream>")
	default:
		return Make_string("primStr unknown")
	}
	return Make_string("wrong input, the object is not atom ...")
}

func stringConcat(args ...Obj) Obj {
	s1 := mustString(args[0])
	s2 := mustString(args[1])
	return Make_string(s1 + s2)
}

func primTailString(args ...Obj) Obj {
	str := mustString(args[0])
	if len(str) == 0 {
		return Make_error("empty string")
	}
	return Make_string(string([]rune(str)[1:]))
}

func pos(args ...Obj) Obj {
	s := []rune(mustString(args[0]))
	n := mustInteger(args[1])
	if n >= len(s) {
		return Make_error(fmt.Sprintf("%d is not valid index for %s", n, string(s)))
	}
	return Make_string(string([]rune(s)[n]))
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
	return Make_error(fmt.Sprintf("variable %s not bound", sym.sym))
}

func simpleError(args ...Obj) Obj {
	str := mustString(args[0])
	return Make_error(str)
}

func primErrorToString(args ...Obj) Obj {
	e := mustError(args[0])
	return Make_string(e.err)
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
		return Make_error("absvector wrong argument")
	}
	return Make_vector(n)
}

func vectorSet(args ...Obj) Obj {
	vec := mustVector(args[0])
	off := mustInteger(args[1])
	val := args[2]
	vec[off] = val
	return args[0]
}

func vectorGet(args ...Obj) Obj {
	vec := mustVector(args[0])
	off := mustInteger(args[1])
	return vec[off]
}

func isVector(args ...Obj) Obj {
	if *args[0] == Vector {
		return True
	}
	return False
}

func writeByte(args ...Obj) Obj {
	n := mustInteger(args[0])
	s := mustStream(args[1])
	w, ok := s.raw.(io.Writer)
	if !ok {
		return Make_error("stream is not opened in out mode")
	}
	var b [1]byte
	b[0] = byte(n)
	_, err := w.Write(b[:])
	if err != nil {
		return Make_error(err.Error())
	}
	return args[0]
}

func primReadByte(args ...Obj) Obj {
	s := mustStream(args[0])
	r, ok := s.raw.(io.Reader)
	if !ok {
		return Make_error("stream is closed of not opened in in mode")
	}
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
		if err == io.EOF {
			return Make_integer(-1)
		}
		return Make_error(err.Error())
	}
	return Make_integer(int(buf[0]))
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
		return Make_error(err.Error())
	}
	return Make_stream(f)
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
		return Make_integer(int(time.Now().Unix()))
	case "run":
		return Make_integer(int(time.Since(time.Now()).Seconds()))
	}
	return Nil
}

func typeFunc(args ...Obj) Obj {
	// TODO: seems meanless
	return args[0]
}

func primIsString(args ...Obj) Obj {
	if *args[0] == String {
		return True
	}
	return False
}

func primIsPair(args ...Obj) Obj {
	if *args[0] == Pair {
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
	return Make_error("primNot")

}

func primIf(args ...Obj) Obj {
	switch args[0] {
	case True:
		return args[1]
	case False:
		return args[2]
	}
	return Make_error("primIf")
}

func PrimEqual(args ...Obj) Obj {
	return equal(args[0], args[1])
}

func primCons(args ...Obj) Obj {
	return cons(args[0], args[1])
}
