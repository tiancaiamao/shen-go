package kl

import (
	"fmt"
	"io"
	"os"
	"time"
)

var primitives []*ScmPrimitive = []*ScmPrimitive{
	&ScmPrimitive{scmHead: Primitive, name: "load-file", required: 1, function: primLoadFile},
	&ScmPrimitive{scmHead: Primitive, name: "type", required: 2, function: typeFunc},
	&ScmPrimitive{scmHead: Primitive, name: "get-time", required: 1, function: getTime},
	&ScmPrimitive{scmHead: Primitive, name: "eval-kl", required: 1, function: primEvalKL},
	&ScmPrimitive{scmHead: Primitive, name: "close", required: 1, function: closeStream},
	&ScmPrimitive{scmHead: Primitive, name: "open", required: 2, function: openStream},
	&ScmPrimitive{scmHead: Primitive, name: "read-byte", required: 1, function: readByte},
	&ScmPrimitive{scmHead: Primitive, name: "write-byte", required: 2, function: writeByte},
	&ScmPrimitive{scmHead: Primitive, name: "absvector?", required: 1, function: isVector},
	&ScmPrimitive{scmHead: Primitive, name: "<-address", required: 2, function: vectorGet},
	&ScmPrimitive{scmHead: Primitive, name: "address->", required: 3, function: vectorSet},
	&ScmPrimitive{scmHead: Primitive, name: "absvector", required: 1, function: primAbsvector},
	&ScmPrimitive{scmHead: Primitive, name: "str", required: 1, function: primStr},
	&ScmPrimitive{scmHead: Primitive, name: "<=", required: 2, function: lessEqual},
	&ScmPrimitive{scmHead: Primitive, name: ">=", required: 2, function: greatEqual},
	&ScmPrimitive{scmHead: Primitive, name: "<", required: 2, function: lessThan},
	&ScmPrimitive{scmHead: Primitive, name: ">", required: 2, function: greatThan},
	&ScmPrimitive{scmHead: Primitive, name: "error-to-string", required: 1, function: primErrorToString},
	&ScmPrimitive{scmHead: Primitive, name: "simple-error", required: 1, function: simpleError},
	&ScmPrimitive{scmHead: Primitive, name: "=", required: 2, function: PrimEqual},
	&ScmPrimitive{scmHead: Primitive, name: "-", required: 2, function: primNumberSubtract},
	&ScmPrimitive{scmHead: Primitive, name: "*", required: 2, function: primNumberMultiply},
	&ScmPrimitive{scmHead: Primitive, name: "/", required: 2, function: primNumberDivide},
	&ScmPrimitive{scmHead: Primitive, name: "+", required: 2, function: PrimNumberAdd},
	&ScmPrimitive{scmHead: Primitive, name: "string->n", required: 1, function: primStringToNumber},
	&ScmPrimitive{scmHead: Primitive, name: "n->string", required: 1, function: primNumberToString},
	&ScmPrimitive{scmHead: Primitive, name: "number?", required: 1, function: primIsNumber},
	&ScmPrimitive{scmHead: Primitive, name: "string?", required: 1, function: primIsString},
	&ScmPrimitive{scmHead: Primitive, name: "pos", required: 2, function: pos},
	&ScmPrimitive{scmHead: Primitive, name: "tlstr", required: 1, function: primTailString},
	&ScmPrimitive{scmHead: Primitive, name: "cn", required: 2, function: stringConcat},
	&ScmPrimitive{scmHead: Primitive, name: "intern", required: 1, function: primIntern},
	&ScmPrimitive{scmHead: Primitive, name: "hd", required: 1, function: primHead},
	&ScmPrimitive{scmHead: Primitive, name: "tl", required: 1, function: primTail},
	&ScmPrimitive{scmHead: Primitive, name: "cons", required: 2, function: primCons},
	&ScmPrimitive{scmHead: Primitive, name: "cons?", required: 1, function: primIsPair},
	&ScmPrimitive{scmHead: Primitive, name: "value", required: 1, function: variableGet},
	&ScmPrimitive{scmHead: Primitive, name: "set", required: 2, function: variableSet},
	&ScmPrimitive{scmHead: Primitive, name: "not", required: 1, function: primNot},
	&ScmPrimitive{scmHead: Primitive, name: "if", required: 3, function: primIf},
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
	return Make_symbol(mustString(args[0]))
}

func primHead(args ...Obj) Obj {
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
	default:
		return Make_string("primStr unknown:default value")
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

func variableSet(args ...Obj) Obj {
	sym := mustSymbol(args[0])
	symbolTable[sym.sym] = args[1]
	return args[1]
}

func variableGet(args ...Obj) Obj {
	sym := mustSymbol(args[0])
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

func readByte(args ...Obj) Obj {
	s := mustStream(args[0])
	r, ok := s.raw.(io.Reader)
	if !ok {
		return Make_error("stream is closed of not opened in in mode")
	}
	var buf [1]byte
	_, err := r.Read(buf[:])
	if err != nil {
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

func primEvalKL(args ...Obj) Obj {
	return trampoline(args[0], nil)
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

func primLoadFile(args ...Obj) Obj {
	path := mustString(args[0])
	return LoadFile(path)
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
