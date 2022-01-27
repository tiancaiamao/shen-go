package cora

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"unsafe"
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
	// {"absvector?", 1,  PrimIsVector},
	// {"<-address", 2,  PrimVectorGet},
	// {"address->", 3,  PrimVectorSet},
	// {"absvector", 1,  PrimAbsvector},
	{"str", 1, PrimStr},
	// {"<=", 2,  PrimLessEqual},
	// {">=", 2,  PrimGreatEqual},
	// {"<", 2,  PrimLessThan},
	// {">", 2,  PrimGreatThan},
	{"error-to-string", 1, PrimErrorToString},
	{"simple-error", 1, PrimSimpleError},
	// {"=", 2,  PrimEqual},
	// {"-", 2,  PrimNumberSubtract},
	// {"*", 2,  PrimNumberMultiply},
	// {"/", 2,  PrimNumberDivide},
	// {"+", 2,  PrimNumberAdd},
	{"string->n", 1, PrimStringToNumber},
	{"n->string", 1, PrimNumberToString},
	// {"number?", 1,  PrimIsNumber},
	// {"string?", 1,  PrimIsString},
	{"pos", 2, PrimPos},
	{"tlstr", 1, PrimTailString},
	{"cn", 2, PrimStringConcat},
	// {"intern", 1,  PrimIntern},
	// {"hd", 1,  PrimHead},
	// {"tl", 1,  PrimTail},
	// {"cons", 2,  PrimCons},
	// {"cons?", 1,  PrimIsPair},
	{"value", 1, PrimNS3Value},
	{"set", 2, PrimNS3Set},
	{"def", 2, PrimNS2Set}, // defun in klambda is compiled to def
	// {"not", 1,  PrimNot},
	{"if", 3, PrimIf},
	// {"symbol?", 1,  PrimIsSymbol},
	{"read-file-as-bytelist", 1, PrimReadFileAsByteList},
	{"read-file-as-string", 1, PrimReadFileAsString},
	// {"variable?", 1, PrimIsVariable},
	{"integer?", 1, PrimIsInteger},
	{"shen.char-stoutput?", 1, PrimCharStOutput},
	{"shen.char-stinput?", 1, PrimCharStInput},
}

func KLInit(e *ControlFlow, test bool) {
	// KLambda primitives
	for _, item := range klPrimitives {
		sym := MakeSymbol(item.name)
		prim := MakePrimitive(item.name, item.arity, item.fn)
		PrimNS2Set(sym, prim)
	}
	PrimNS2Set(MakeSymbol("eval-kl"), MakeNative(primEvalKL, 1))
	PrimNS2Set(MakeSymbol("load-file"), MakeNative(primLoad, 1))
	// Overload for primitive set and value.
	PrimNS3Set(MakeSymbol("*stinput*"), MakeStream(os.Stdin))
	PrimNS3Set(MakeSymbol("*stoutput*"), MakeStream(os.Stdout))
	dir, _ := os.Getwd()
	PrimNS3Set(MakeSymbol("*home-directory*"), MakeString(dir))
	PrimNS3Set(MakeSymbol("*release*"), MakeString(runtime.Version()))
	PrimNS3Set(MakeSymbol("*os*"), MakeString(runtime.GOOS))
	// PrimNS1Set(MakeSymbol("kl.init"), MakeNative(primKLInit, 0))

	primKLInit(e, test)
}

func PrimStr(o Obj) Obj {
	if isFixnum(uintptr(unsafe.Pointer(o))) {
		return MakeString(fmt.Sprintf("%d", mustInteger(o)))
	}

	switch *o {
	case scmHeadPair:
		// This is actually a special representation for closure in cora.
		// To make shen happy ...
		if car(o) == symLambda {
			return MakeString("#<closure >")
		}
		// Pair may contain recursive list.
		panic(MakeError("can't str pair object"))
	case scmHeadNull:
		return MakeString("()")
	case scmHeadSymbol:
		str := GetSymbol(o)
		return MakeString(str)
	case scmHeadNumber:
		f := mustNumber(o)
		if !isPreciseInteger(f) {
			return MakeString(fmt.Sprintf("%f", f))
		}
		return MakeString(fmt.Sprintf("%d", int(f)))
	case scmHeadString:
		return MakeString(fmt.Sprintf(`"%s"`, mustString(o)))
	case scmHeadBoolean:
		if o == True {
			return MakeString("true")
		} else if o == False {
			return MakeString("false")
		}
	case scmHeadError:
		err := mustError(o)
		return MakeString("#<err " + err.err + ">")
	case scmHeadStream:
		return MakeString("#<stream >")
	case scmHeadRaw:
		return MakeString("#<raw >")
	case scmHeadNative:
		n := MustNative(o)
		if len(n.name) > 0 {
			return MakeString(fmt.Sprintf("#<primitive %s>", n.name))
		} else {
			return MakeString(fmt.Sprintf("#<native %v>", *n))
		}
	default:
		return MakeString("PrimStr unknown")

	}
	return MakeString("wrong input, the object is not atom ...")
}

func PrimSimpleError(o Obj) Obj {
	str := mustString(o)
	panic(MakeError(str))
}

func PrimErrorToString(o Obj) Obj {
	err := mustError(o)
	return MakeString(err.err)
}

func PrimPos(x, y Obj) Obj {
	s := []rune(mustString(x))
	n := mustInteger(y)
	if n >= len(s) {
		panic(MakeError(fmt.Sprintf("%d is not valid index for %s", n, string(s))))
	}
	return MakeString(string([]rune(s)[n]))
}

func PrimCharStInput(x Obj) Obj {
	return False
}

func PrimCharStOutput(x Obj) Obj {
	return False
}

func primLoad(e *ControlFlow) {
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

		res := evalKL(e, exp)
		if IsError(res) {
			e.Return(res)
			return
		}
	}
	e.Return(MakeSymbol("loaded"))
}

func primEvalKL(e *ControlFlow) {
	exp1 := e.Get(1)
	res := evalKL(e, exp1)
	e.Return(res)
}

func evalKL(e *ControlFlow, exp Obj) Obj {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("evalKL panic:", ObjString(exp))
			panic(r)
		}
	}()
	// exp1 := Call(e, PrimNS1Value(MakeSymbol("kl->cora")), Nil, exp)
	exp1 := NCall(PrimNS1Value(MakeSymbol("kl->cora")), Nil, exp)

	// fmt.Println("evalKL with ===", kl.ObjString(exp1))
	res := Neval(exp1)
	return res
}

func primKLInit(e *ControlFlow, test bool) {
	var evaluator Evaluator = controlFlowAsEvaluator{e}
	if test {
		evaluator = closureEvaluator{}
	}
	f, err := initFS.Open("kl.cora")
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()
	r := NewSexpReader(f, true)
	res := loadFileFromReader(evaluator, true, r)
	e.Return(res)
}
