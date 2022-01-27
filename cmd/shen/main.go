package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	// "path/filepath"
	"runtime"

	"github.com/tiancaiamao/shen-go/cora"
)

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

func regist(e *cora.ControlFlow) {
	for _, init := range []cora.Obj{
		KLMain,
		SysMain,
		WriterMain,
		CoreMain,
		ReaderMain,
		DeclarationsMain,
		TopLevelMain,
		MacrosMain,
		LoadMain,
		PrologMain,
		SequentMain,
		TrackMain,
		TStarMain,
		YaccMain,
		TypesMain,
	} {
		res := cora.Call(e, init)
		if cora.IsError(res) {
			fmt.Println("load ...fail")
		}
	}
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e cora.ControlFlow
	// klInit()

	cora.PrimNS1Set(symMacroExpand, cora.Nil)
	cora.CoraInit(&e, false)
	if cora.IsError(e.Get(0)) {
		os.Exit(-1)
	}

	cora.KLInit(&e, false)
	if cora.IsError(e.Get(0)) {
		os.Exit(-1)
	}
	cora.PrimNS2Set(symEvalKL, cora.MakeNative(primEvalKL, 1))

	// if err := cora.Call(&e, cora.PrimNS1Value(cora.MakeSymbol("cora.init"))); cora.IsError(err) {
	// 	os.Exit(-1)
	// }

	regist(&e)
	evalKL(&e, cora.Cons(cora.MakeSymbol("shen.shen"), cora.Nil))
	// r := kl.NewSexpReader(os.Stdin, false)
	// for i := 0; ; i++ {
	// 	fmt.Printf("%d #> ", i)
	// 	sexp, err := r.Read()
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println("read error:", err)
	// 		}
	// 		break
	// 	}
	// 	res := evalKL(&e, sexp)
	// 	fmt.Println(kl.ObjString(res))
	// }
}

var (
	symMacroExpand = cora.MakeSymbol("macroexpand")
	symEvalKL      = cora.MakeSymbol("eval-kl")
	symKLToCora    = cora.MakeSymbol("kl->cora")
)

var klPrimitives = []struct {
	name  string
	arity int
	fn    interface{}
}{
	{"get-time", 1, cora.PrimGetTime},
	{"close", 1, cora.PrimCloseStream},
	{"open", 2, cora.PrimOpenStream},
	{"read-byte", 1, cora.PrimReadByte},
	{"write-byte", 2, cora.PrimWriteByte},
	{"absvector?", 1, cora.PrimIsVector},
	{"<-address", 2, cora.PrimVectorGet},
	{"address->", 3, cora.PrimVectorSet},
	{"absvector", 1, cora.PrimAbsvector},
	{"str", 1, cora.PrimStr},
	{"<=", 2, cora.PrimLessEqual},
	{">=", 2, cora.PrimGreatEqual},
	{"<", 2, cora.PrimLessThan},
	{">", 2, cora.PrimGreatThan},
	{"error-to-string", 1, cora.PrimErrorToString},
	{"simple-error", 1, cora.PrimSimpleError},
	{"=", 2, cora.PrimEqual},
	{"-", 2, cora.PrimNumberSubtract},
	{"*", 2, cora.PrimNumberMultiply},
	{"/", 2, cora.PrimNumberDivide},
	{"+", 2, cora.PrimNumberAdd},
	{"string->n", 1, cora.PrimStringToNumber},
	{"n->string", 1, cora.PrimNumberToString},
	{"number?", 1, cora.PrimIsNumber},
	{"string?", 1, cora.PrimIsString},
	{"pos", 2, cora.PrimPos},
	{"tlstr", 1, cora.PrimTailString},
	{"cn", 2, cora.PrimStringConcat},
	{"intern", 1, cora.PrimIntern},
	{"hd", 1, cora.PrimHead},
	{"tl", 1, cora.PrimTail},
	{"cons", 2, cora.PrimCons},
	{"cons?", 1, cora.PrimIsPair},
	{"value", 1, cora.PrimNS3Value},
	{"set", 2, cora.PrimNS3Set},
	{"not", 1, cora.PrimNot},
	{"if", 3, cora.PrimIf},
	{"symbol?", 1, cora.PrimIsSymbol},
	{"read-file-as-bytelist", 1, cora.PrimReadFileAsByteList},
	{"read-file-as-string", 1, cora.PrimReadFileAsString},
	{"variable?", 1, PrimIsVariable},
	{"integer?", 1, cora.PrimIsInteger},
}

func klInit() {
	// kl.BindSymbolFunc(kl.MakeSymbol("load-file"), kl.MakeNative(primLoadFile(false), 1))
	cora.PrimNS1Set(cora.MakeSymbol("ns2-set"), cora.MakePrimitive("ns2-set", 2, cora.PrimNS2Set))
	cora.PrimNS1Set(cora.MakeSymbol("ns2-value"), cora.MakePrimitive("ns2-value", 1, cora.PrimNS2Value))
	for _, item := range klPrimitives {
		sym := cora.MakeSymbol(item.name)
		prim := cora.MakePrimitive(item.name, item.arity, item.fn)
		cora.PrimNS2Set(sym, prim)
	}
	cora.PrimNS2Set(symEvalKL, cora.MakeNative(primEvalKL, 1))
	cora.PrimNS2Set(cora.MakeSymbol("load-file"), cora.MakeNative(primLoad, 1))

	// Overload for primitive set and value.
	cora.PrimNS3Set(cora.MakeSymbol("*stinput*"), cora.MakeStream(os.Stdin))
	cora.PrimNS3Set(cora.MakeSymbol("*stoutput*"), cora.MakeStream(os.Stdout))
	dir, _ := os.Getwd()
	cora.PrimNS3Set(cora.MakeSymbol("*home-directory*"), cora.MakeString(dir))
	cora.PrimNS3Set(cora.MakeSymbol("*language*"), cora.MakeString("Go"))
	cora.PrimNS3Set(cora.MakeSymbol("*implementation*"), cora.MakeString("AOT+interpreter"))
	cora.PrimNS3Set(cora.MakeSymbol("*relase*"), cora.MakeString(runtime.Version()))
	cora.PrimNS3Set(cora.MakeSymbol("*os*"), cora.MakeString(runtime.GOOS))
	cora.PrimNS3Set(cora.MakeSymbol("*porters*"), cora.MakeString("Arthur Mao"))
	cora.PrimNS3Set(cora.MakeSymbol("*port*"), cora.MakeString("1.0.0-rc1"))
}

func primEvalKL(e *cora.ControlFlow) {
	exp1 := e.Get(1)
	res := evalKL(e, exp1)
	e.Return(res)
}

func evalKL(e *cora.ControlFlow, exp cora.Obj) cora.Obj {
	exp1 := cora.Call(e, cora.PrimNS1Value(symKLToCora), cora.Nil, exp)

	// fmt.Println("evalKL with ===", kl.ObjString(exp1))
	res := cora.Eval(e, exp1)
	return res
}

func primLoad(e *cora.ControlFlow) {
	file := e.Get(1)
	if !cora.IsString(file) {
		e.Return(cora.MakeError("arg1 must be string"))
		return
	}
	path := cora.GetString(file)
	if _, err := os.Stat(path); err != nil {
		e.Return(cora.MakeError(err.Error()))
		return
	}

	f, err := os.Open(path)
	if err != nil {
		e.Return(cora.MakeError(err.Error()))
		return
	}
	defer f.Close()

	r := cora.NewSexpReader(f, false)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				e.Return(cora.MakeError(err.Error()))
				return
			}
			break
		}

		res := evalKL(e, exp)
		if cora.IsError(res) {
			e.Return(res)
			return
		}
	}
	e.Return(cora.MakeSymbol("loaded"))
}

func PrimIsVariable(x cora.Obj) cora.Obj {
	if !cora.IsSymbol(x) {
		return cora.False
	}

	sym := cora.GetSymbol(x)
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		return cora.False
	}
	return cora.True
}
