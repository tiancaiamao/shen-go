package main

// KL-native bindings over the shared bytecode-IR -> Go translator (package
// codegen). The translator itself lives in codegen/ so other build tools
// (cmd/yggdrasil-build) can drive it without going through KL.

import (
	"os"
	"unsafe"

	"github.com/tiancaiamao/shen-go/codegen"
	"github.com/tiancaiamao/shen-go/kl"
)

var makeCodeGenerator = kl.MakeNative(func(e *kl.ControlFlow) {
	// (make-code-generator 'cora)
	cg := codegen.New()
	e.Return(kl.MakeRaw(&cg.ScmHead))

}, 0)

var bcToGo = kl.MakeNative(func(e *kl.ControlFlow) {
	// (let cg (make-code-generator)
	//      (bc->go cg "Main" true "xx.bc" "xx.go"))
	cg := (*codegen.CodeGenerator)(unsafe.Pointer(e.Get(1)))
	exportName := kl.GetString(e.Get(2))
	genSym := e.Get(3)
	inFile := kl.GetString(e.Get(4))
	outFile := kl.GetString(e.Get(5))

	f, err := os.Open(inFile)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	defer f.Close()

	out, err := os.Create(outFile)
	if err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	defer out.Close()

	if err := cg.HandleBody(f, exportName, out); err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	if genSym == kl.True {
		cg.HandleSymbol(out)
	}

	e.Return(kl.Nil)

}, 5)
