package main

import (
	"flag"
	"fmt"
	// "io"
	"net/http"
	_ "net/http/pprof"
	// "os"
	// "path/filepath"
	// "runtime"

	"github.com/tiancaiamao/shen-go/klambda"
)

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

func regist(e *klambda.ControlFlow) {
	for _, init := range []klambda.Obj{
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
		res := klambda.Call(e, init)
		if klambda.IsError(res) {
			fmt.Println("load ...fail")
		}
	}
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e klambda.ControlFlow
	klambda.PrimNS2Set(klambda.MakeSymbol("def"), klambda.MakePrimitive("def", 2, klambda.PrimNS2Set))
	// klInit()

	// klambda.PrimNS1Set(symMacroExpand, klambda.Nil)
	// klambda.Init(&e, false)
	// if klambda.IsError(e.Get(0)) {
	// 	os.Exit(-1)
	// }

	// klambda.KLInit(&e, false)
	// if klambda.IsError(e.Get(0)) {
	// 	os.Exit(-1)
	// }
	// klambda.PrimNS2Set(symEvalKL, klambda.MakeNative(primEvalKL, 1))

	// if err := klambda.Call(&e, klambda.PrimNS1Value(klambda.MakeSymbol("klambda.init"))); klambda.IsError(err) {
	// 	os.Exit(-1)
	// }

	regist(&e)
	e.Eval(klambda.Cons(klambda.MakeSymbol("shen.shen"), klambda.Nil))
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
	symMacroExpand = klambda.MakeSymbol("macroexpand")
	symEvalKL      = klambda.MakeSymbol("eval-kl")
	symKLToKlambda = klambda.MakeSymbol("kl->klambda")
)

// var klPrimitives = []struct {
// 	name  string
// 	arity int
// 	fn    interface{}
// }{
// 	{"get-time", 1, klambda.PrimGetTime},
// 	{"close", 1, klambda.PrimCloseStream},
// 	{"open", 2, klambda.PrimOpenStream},
// 	{"read-byte", 1, klambda.PrimReadByte},
// 	{"write-byte", 2, klambda.PrimWriteByte},
// 	{"absvector?", 1, klambda.PrimIsVector},
// 	{"<-address", 2, klambda.PrimVectorGet},
// 	{"address->", 3, klambda.PrimVectorSet},
// 	{"absvector", 1, klambda.PrimAbsvector},
// 	{"str", 1, klambda.PrimStr},
// 	{"<=", 2, klambda.PrimLessEqual},
// 	{">=", 2, klambda.PrimGreatEqual},
// 	{"<", 2, klambda.PrimLessThan},
// 	{">", 2, klambda.PrimGreatThan},
// 	{"error-to-string", 1, klambda.PrimErrorToString},
// 	{"simple-error", 1, klambda.PrimSimpleError},
// 	{"=", 2, klambda.PrimEqual},
// 	{"-", 2, klambda.PrimNumberSubtract},
// 	{"*", 2, klambda.PrimNumberMultiply},
// 	{"/", 2, klambda.PrimNumberDivide},
// 	{"+", 2, klambda.PrimNumberAdd},
// 	{"string->n", 1, klambda.PrimStringToNumber},
// 	{"n->string", 1, klambda.PrimNumberToString},
// 	{"number?", 1, klambda.PrimIsNumber},
// 	{"string?", 1, klambda.PrimIsString},
// 	{"pos", 2, klambda.PrimPos},
// 	{"tlstr", 1, klambda.PrimTailString},
// 	{"cn", 2, klambda.PrimStringConcat},
// 	{"intern", 1, klambda.PrimIntern},
// 	{"hd", 1, klambda.PrimHead},
// 	{"tl", 1, klambda.PrimTail},
// 	{"cons", 2, klambda.PrimCons},
// 	{"cons?", 1, klambda.PrimIsPair},
// 	{"value", 1, klambda.PrimNS3Value},
// 	{"set", 2, klambda.PrimNS3Set},
// 	{"not", 1, klambda.PrimNot},
// 	{"if", 3, klambda.PrimIf},
// 	{"symbol?", 1, klambda.PrimIsSymbol},
// 	{"read-file-as-bytelist", 1, klambda.PrimReadFileAsByteList},
// 	{"read-file-as-string", 1, klambda.PrimReadFileAsString},
// 	{"variable?", 1, klambda.PrimIsVariable},
// 	{"integer?", 1, klambda.PrimIsInteger},
// }

// func klInit() {
// 	// kl.BindSymbolFunc(kl.MakeSymbol("load-file"), kl.MakeNative(primLoadFile(false), 1))
// 	klambda.PrimNS1Set(klambda.MakeSymbol("ns2-set"), klambda.MakePrimitive("ns2-set", 2, klambda.PrimNS2Set))
// 	klambda.PrimNS1Set(klambda.MakeSymbol("ns2-value"), klambda.MakePrimitive("ns2-value", 1, klambda.PrimNS2Value))
// 	for _, item := range klPrimitives {
// 		sym := klambda.MakeSymbol(item.name)
// 		prim := klambda.MakePrimitive(item.name, item.arity, item.fn)
// 		klambda.PrimNS2Set(sym, prim)
// 	}
// 	klambda.PrimNS2Set(symEvalKL, klambda.MakeNative(primEvalKL, 1))
// 	klambda.PrimNS2Set(klambda.MakeSymbol("load-file"), klambda.MakeNative(primLoad, 1))

// 	// Overload for primitive set and value.
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*stinput*"), klambda.MakeStream(os.Stdin))
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*stoutput*"), klambda.MakeStream(os.Stdout))
// 	dir, _ := os.Getwd()
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*home-directory*"), klambda.MakeString(dir))
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*language*"), klambda.MakeString("Go"))
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*implementation*"), klambda.MakeString("AOT+interpreter"))
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*relase*"), klambda.MakeString(runtime.Version()))
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*os*"), klambda.MakeString(runtime.GOOS))
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*porters*"), klambda.MakeString("Arthur Mao"))
// 	klambda.PrimNS3Set(klambda.MakeSymbol("*port*"), klambda.MakeString("1.0.0-rc1"))
// }

// func primEvalKL(e *klambda.ControlFlow) {
// 	exp1 := e.Get(1)
// 	res := e.Eval(exp1)
// 	e.Return(res)
// }

// func primLoad(e *klambda.ControlFlow) {
// 	file := e.Get(1)
// 	if !klambda.IsString(file) {
// 		e.Return(klambda.MakeError("arg1 must be string"))
// 		return
// 	}
// 	path := klambda.GetString(file)
// 	if _, err := os.Stat(path); err != nil {
// 		e.Return(klambda.MakeError(err.Error()))
// 		return
// 	}

// 	f, err := os.Open(path)
// 	if err != nil {
// 		e.Return(klambda.MakeError(err.Error()))
// 		return
// 	}
// 	defer f.Close()

// 	r := klambda.NewSexpReader(f, false)
// 	for {
// 		exp, err := r.Read()
// 		if err != nil {
// 			if err != io.EOF {
// 				e.Return(klambda.MakeError(err.Error()))
// 				return
// 			}
// 			break
// 		}

// 		res := e.Eval(exp)
// 		if klambda.IsError(res) {
// 			e.Return(res)
// 			return
// 		}
// 	}
// 	e.Return(klambda.MakeSymbol("loaded"))
// }

// func PrimIsVariable(x klambda.Obj) klambda.Obj {
// 	if !klambda.IsSymbol(x) {
// 		return klambda.False
// 	}

// 	sym := klambda.GetSymbol(x)
// 	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
// 		return klambda.False
// 	}
// 	return klambda.True
// }

func primDef(x klambda.Obj) klambda.Obj {
	if !klambda.IsSymbol(x) {
		return klambda.False
	}

	sym := klambda.GetSymbol(x)
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		return klambda.False
	}
	return klambda.True
}
