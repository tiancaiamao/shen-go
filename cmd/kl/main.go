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

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var e kl.ControlFlow
	klInit()

	kl.PrimNS1Set(symMacroExpand, kl.Nil)
	if err := kl.Call(&e, kl.PrimNS1Value(kl.MakeSymbol("cora.init"))); kl.IsError(err) {
		os.Exit(-1)
	}
	if err := kl.Call(&e, kl.PrimNS1Value(kl.MakeSymbol("kl.init"))); kl.IsError(err) {
		os.Exit(-1)
	}
	// gopath := os.Getenv("GOPATH")
	// initFiles := []string{
	// 	filepath.Join(gopath, "src/github.com/tiancaiamao/shen-go/kl", "init.cora"),
	// 	filepath.Join(gopath, "src/github.com/tiancaiamao/shen-go/src", "build.cora"),
	// }
	// for _, file := range initFiles {
	// 	if _, err := os.Stat(file); err == nil {
	// 		res := kl.Call(&e, kl.PrimNS1Value(kl.MakeSymbol("load")), kl.MakeString(file))
	// 		if kl.IsError(res) {
	// 			os.Exit(-1)
	// 		}
	// 	}
	// }
	// kl.Call(&e, kl.PrimNS1Value(kl.MakeSymbol("load")), kl.MakeString("build.cora"))

	r := kl.NewSexpReader(os.Stdin, false)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		res := evalKL(&e, sexp)
		fmt.Println(kl.ObjString(res))
	}
}

var (
	symMacroExpand = kl.MakeSymbol("macroexpand")
	symEvalKL      = kl.MakeSymbol("eval-kl")
	symKLToCora    = kl.MakeSymbol("kl->cora")
)

var klPrimitives = []struct {
	name  string
	arity int
	fn    interface{}
}{
	{"get-time", 1, kl.PrimGetTime},
	{"close", 1, kl.PrimCloseStream},
	{"open", 2, kl.PrimOpenStream},
	{"read-byte", 1, kl.PrimReadByte},
	{"write-byte", 2, kl.PrimWriteByte},
	{"absvector?", 1, kl.PrimIsVector},
	{"<-address", 2, kl.PrimVectorGet},
	{"address->", 3, kl.PrimVectorSet},
	{"absvector", 1, kl.PrimAbsvector},
	{"str", 1, kl.PrimStr},
	{"<=", 2, kl.PrimLessEqual},
	{">=", 2, kl.PrimGreatEqual},
	{"<", 2, kl.PrimLessThan},
	{">", 2, kl.PrimGreatThan},
	{"error-to-string", 1, kl.PrimErrorToString},
	{"simple-error", 1, kl.PrimSimpleError},
	{"=", 2, kl.PrimEqual},
	{"-", 2, kl.PrimNumberSubtract},
	{"*", 2, kl.PrimNumberMultiply},
	{"/", 2, kl.PrimNumberDivide},
	{"+", 2, kl.PrimNumberAdd},
	{"string->n", 1, kl.PrimStringToNumber},
	{"n->string", 1, kl.PrimNumberToString},
	{"number?", 1, kl.PrimIsNumber},
	{"string?", 1, kl.PrimIsString},
	{"pos", 2, kl.PrimPos},
	{"tlstr", 1, kl.PrimTailString},
	{"cn", 2, kl.PrimStringConcat},
	{"intern", 1, kl.PrimIntern},
	{"hd", 1, kl.PrimHead},
	{"tl", 1, kl.PrimTail},
	{"cons", 2, kl.PrimCons},
	{"cons?", 1, kl.PrimIsPair},
	{"value", 1, kl.PrimNS3Value},
	{"set", 2, kl.PrimNS3Set},
	{"not", 1, kl.PrimNot},
	{"if", 3, kl.PrimIf},
	{"symbol?", 1, kl.PrimIsSymbol},
	{"read-file-as-bytelist", 1, kl.PrimReadFileAsByteList},
	{"read-file-as-string", 1, kl.PrimReadFileAsString},
	{"variable?", 1, PrimIsVariable},
	{"integer?", 1, kl.PrimIsInteger},
}

func klInit() {
	// kl.BindSymbolFunc(kl.MakeSymbol("load-file"), kl.MakeNative(primLoadFile(false), 1))
	kl.PrimNS1Set(kl.MakeSymbol("ns2-set"), kl.MakePrimitive("ns2-set", 2, kl.PrimNS2Set))
	kl.PrimNS1Set(kl.MakeSymbol("ns2-value"), kl.MakePrimitive("ns2-value", 1, kl.PrimNS2Value))
	for _, item := range klPrimitives {
		sym := kl.MakeSymbol(item.name)
		prim := kl.MakePrimitive(item.name, item.arity, item.fn)
		kl.PrimNS2Set(sym, prim)
	}
	kl.PrimNS2Set(symEvalKL, kl.MakeNative(primEvalKL, 1))
	kl.PrimNS2Set(kl.MakeSymbol("load-file"), kl.MakeNative(primLoad, 1))

	// Overload for primitive set and value.
	kl.PrimNS3Set(kl.MakeSymbol("*stinput*"), kl.MakeStream(os.Stdin))
	kl.PrimNS3Set(kl.MakeSymbol("*stoutput*"), kl.MakeStream(os.Stdout))
	dir, _ := os.Getwd()
	kl.PrimNS3Set(kl.MakeSymbol("*home-directory*"), kl.MakeString(dir))
	kl.PrimNS3Set(kl.MakeSymbol("*language*"), kl.MakeString("Go"))
	kl.PrimNS3Set(kl.MakeSymbol("*implementation*"), kl.MakeString("AOT+interpreter"))
	kl.PrimNS3Set(kl.MakeSymbol("*relase*"), kl.MakeString(runtime.Version()))
	kl.PrimNS3Set(kl.MakeSymbol("*os*"), kl.MakeString(runtime.GOOS))
	kl.PrimNS3Set(kl.MakeSymbol("*porters*"), kl.MakeString("Arthur Mao"))
	kl.PrimNS3Set(kl.MakeSymbol("*port*"), kl.MakeString("1.0.0-rc1"))
}

func primEvalKL(e *kl.ControlFlow) {
	exp1 := e.Get(1)
	res := evalKL(e, exp1)
	e.Return(res)
}

func evalKL(e *kl.ControlFlow, exp kl.Obj) kl.Obj {
	exp1 := kl.Call(e, kl.PrimNS1Value(symKLToCora), kl.Nil, exp)

	// fmt.Println("evalKL with ===", kl.ObjString(exp1))
	res := kl.Eval(e, exp1)
	return res
}

func primLoad(e *kl.ControlFlow) {
	file := e.Get(1)
	if !kl.IsString(file) {
		e.Return(kl.MakeError("arg1 must be string"))
		return
	}
	path := kl.GetString(file)
	if _, err := os.Stat(path); err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}

	f, err := os.Open(path)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	defer f.Close()

	r := kl.NewSexpReader(f, false)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				e.Return(kl.MakeError(err.Error()))
				return
			}
			break
		}

		res := evalKL(e, exp)
		if kl.IsError(res) {
			e.Return(res)
			return
		}
	}
	e.Return(kl.MakeSymbol("loaded"))
}

func PrimIsVariable(x kl.Obj) kl.Obj {
	if !kl.IsSymbol(x) {
		return kl.False
	}

	sym := kl.GetSymbol(x)
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		return kl.False
	}
	return kl.True
}
