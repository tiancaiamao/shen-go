package klambda

import (
	"embed"
	"io"
	"os"
	"runtime"

	. "github.com/tiancaiamao/cora/cora_go"
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
	klInit0(e)
	primKLInit(e, test)
}

func klInit0(e *ControlFlow) {
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
}

//go:embed kl.cora
var initFS embed.FS

func primKLInit(e *ControlFlow, test bool) {
	var evaluator Evaluator = ControlFlowAsEvaluator{e}
	if test {
		evaluator = DefaultEvaluator()
	}
	f, err := initFS.Open("kl.cora")
	if err != nil {
		e.Return(MakeError(err.Error()))
		return
	}
	defer f.Close()
	r := NewSexpReader(f, true)
	res := LoadFileFromReader(evaluator, true, r)
	e.Return(res)
}

// func PrimPos(x, y Obj) Obj {
// 	s := []rune(GetString(x))
// 	n := GetInteger(y)
// 	if n >= len(s) {
// 		panic(MakeError(fmt.Sprintf("%d is not valid index for %s", n, string(s))))
// 	}
// 	return MakeString(string([]rune(s)[n]))
// }

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

		res := EvalKL(exp)
		if IsError(res) {
			e.Return(res)
			return
		}
	}
	e.Return(MakeSymbol("loaded"))
}

func primEvalKL(e *ControlFlow) {
	exp := e.Get(1)
	res := EvalKL(exp)
	e.Return(res)
}

// func PrimIsVariable(x Obj) Obj {
// 	if !IsSymbol(x) {
// 		return False
// 	}

// 	sym := GetSymbol(x)
// 	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
// 		return False
// 	}
// 	return True
// }
