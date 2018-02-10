package vm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"plugin"
	"runtime"
	"sync"
	"unsafe"

	"github.com/tiancaiamao/shen-go/kl"
)

// Go doesn't provide MACRO like C, but the compiler optimization can eliminate dead code "if false XX".
const enableDebug = false

type statusType int

const (
	statusNormal statusType = iota
	statusException
	statusHalt
)

type instFunc func(*VM)
type Code []instFunc

type VM struct {
	stack []kl.Obj
	top   int // stack top

	env      []kl.Obj // persistent environment
	volatile []kl.Obj // volatile environment
	envMark  int      // volatile[envMark: len(volatile)] is current env

	pc        int       // pc register refer to the position in current code
	savedAddr []address // saved return address

	nativeFunc map[string]*kl.ScmPrimitive

	// jumpBuf is used to implement exception, similar to setjmp/longjmp in C.
	cc []jumpBuf

	code   []instFunc
	status statusType
}

type jumpBuf struct {
	address
	savedAddrPos  int
	savedStackTop int
	savedEnvTop   int
	closure       kl.Obj
}

// address is the information to be saved before apply a closure.
type address struct {
	pc      int
	code    Code
	env     []kl.Obj
	envMark int
}

type Procedure struct {
	scmHead int
	code    Code
	env     []kl.Obj
}

const initStackSize = 128

type pool struct {
	sync.Mutex
	data []*VM
}

func (p *pool) Get() *VM {
	p.Lock()
	defer p.Unlock()

	if len(p.data) == 0 {
		return newVM()
	}
	ret := p.data[0]
	p.data = p.data[1:]
	return ret
}

func (p *pool) Put(v *VM) {
	p.Lock()
	p.data = append(p.data, v)
	p.Unlock()
}

var auxVM pool
var stackMarkDummyValue int
var stackMark = kl.MakeRaw(&stackMarkDummyValue)

func New() *VM {
	m := newVM()
	m.RegistNativeCall("load-bytecode", 1, m.loadBytecode)
	m.RegistNativeCall("load-file", 1, m.loadFile)
	m.RegistNativeCall("load-plugin", 1, m.loadPlugin)
	m.RegistNativeCall("plugin-bind", 2, m.pluginBind)
	m.RegistNativeCall("primitive?", 1, kl.NativeIsPrimitive)
	m.RegistNativeCall("primitive-arity", 1, kl.NativePrimitiveArity)
	m.RegistNativeCall("primitive-id", 1, kl.NativePrimitiveID)
	return m
}

func newVM() *VM {
	m := &VM{
		stack:      make([]kl.Obj, initStackSize),
		env:        make([]kl.Obj, 0, 256),
		volatile:   make([]kl.Obj, 0, 256),
		nativeFunc: make(map[string]*kl.ScmPrimitive),
	}
	initSymbolTable()
	return m
}

func (m *VM) RegistNativeCall(name string, arity int, f func(...kl.Obj) kl.Obj) {
	m.nativeFunc[name] = kl.MakePrimitive(name, arity, f)
}

func initSymbolTable() {
	dir, _ := os.Getwd()
	kl.PrimSet(kl.MakeSymbol("*stinput*"), kl.MakeStream(os.Stdin))
	kl.PrimSet(kl.MakeSymbol("*stoutput*"), kl.MakeStream(os.Stdout))
	kl.PrimSet(kl.MakeSymbol("*home-directory*"), kl.MakeString(dir))
	kl.PrimSet(kl.MakeSymbol("*language*"), kl.MakeString("Go"))
	kl.PrimSet(kl.MakeSymbol("*implementation*"), kl.MakeString("bytecode"))
	kl.PrimSet(kl.MakeSymbol("*relase*"), kl.MakeString(runtime.Version()))
	kl.PrimSet(kl.MakeSymbol("*os*"), kl.MakeString(runtime.GOOS))
	kl.PrimSet(kl.MakeSymbol("*porters*"), kl.MakeString("Arthur Mao"))
	kl.PrimSet(kl.MakeSymbol("*port*"), kl.MakeString("0.0.1"))

	// Extended by shen-go implementation
	kl.PrimSet(kl.MakeSymbol("*package-path*"), kl.MakeString(kl.PackagePath()))
}

func (m *VM) Run(code Code) kl.Obj {
	m.setup(code)

	// May be the fastest dispatch method in Go.
Dispatch:
	m.status = statusNormal
	for m.status == statusNormal {
		m.code[m.pc](m)
	}

	if m.status == statusException {
		if len(m.cc) == 0 {
			err := m.stack[m.top-1]
			m.Reset()
			return err
		}
		m.handleException()
		goto Dispatch
	}
	return m.done()
}

func (m *VM) setup(code Code) {
	m.code = code
	m.pc = 0
	m.stackPush(stackMark)
	m.savedAddr = append(m.savedAddr, address{pc: len(code) - 1, code: code})
}

func (m *VM) done() kl.Obj {
	if m.top != 1 || len(m.savedAddr) != 0 || m.envMark != 0 {
		if enableDebug {
			m.Debug()
		}
		panic("m in wrong status")
	}
	m.top--
	return m.stack[m.top]
}

func (m *VM) handleException() {
	// clear jmpBuf
	jmpBuf := m.cc[len(m.cc)-1]
	m.cc = m.cc[:len(m.cc)-1]
	// pop trap-error handler, prepare for call.
	value := m.stack[m.top-1]
	m.top = jmpBuf.savedStackTop
	m.stackPush(stackMark)
	m.stackPush(value)
	// recover savedAddr
	m.savedAddr = m.savedAddr[:jmpBuf.savedAddrPos]
	m.savedAddr = append(m.savedAddr, jmpBuf.address)
	// longjmp... tail apply
	closure := (*Procedure)(unsafe.Pointer(jmpBuf.closure))
	m.code = closure.code
	m.pc = 0
	m.env = closure.env
	m.envMark = jmpBuf.savedEnvTop
	m.volatile = m.volatile[:m.envMark]
}

func opSetJmp(n int) instFunc {
	return func(m *VM) {
		m.pc++
		if enableDebug {
			debugf("SETJMP\n")
		}

		m.top--
		cc := jumpBuf{
			address: address{
				pc:      m.pc + n,
				code:    m.code,
				env:     m.env,
				envMark: m.envMark,
			},
			savedAddrPos:  len(m.savedAddr),
			savedStackTop: m.top,
			closure:       m.stack[m.top],
			savedEnvTop:   len(m.volatile),
		}
		m.cc = append(m.cc, cc)
	}
}

func opClearJmp(m *VM) {
	m.pc++
	m.cc = m.cc[:len(m.cc)-1]
}

func opConst(o kl.Obj) instFunc {
	return func(m *VM) {
		m.pc++
		m.stackPush(o)
		if enableDebug {
			debugf("CONST %s\n", kl.ObjString(o))
		}
	}
}

func opAccess(n int) instFunc {
	return func(m *VM) {
		m.pc++
		if n+m.envMark < len(m.volatile) {
			// get value from volatile environment
			v := m.volatile[len(m.volatile)-1-n]
			m.stackPush(v)
			if enableDebug {
				debugf("ACCESS %d, get %s\n", n, kl.ObjString(v))
			}
		} else {
			// get value from persistent environment
			xx := n - (len(m.volatile) - m.envMark)
			v := m.env[len(m.env)-1-xx]
			m.stackPush(v)
			if enableDebug {
				debugf("ACCESS %d from env, get %s\n", n, kl.ObjString(v))
			}
		}
	}
}

func opFreeze(n int) instFunc {
	return func(m *VM) {
		m.pc++
		// create closure directly
		// nearly the same with grab, but if need zero arguments.
		tmp := &Procedure{
			code: m.code[m.pc:],
			env:  m.envClose(),
		}
		raw := kl.MakeRaw(&tmp.scmHead)
		if enableDebug {
			debugf("FREEZE len(env)=%d\n", len(tmp.env))
		}
		m.stackPush(raw)
		m.pc += n
	}
}

func opMark(m *VM) {
	m.pc++
	m.stackPush(stackMark)
	if enableDebug {
		debugln("MARK")
	}
}

func opGrab(m *VM) {
	m.pc++
	m.top--
	if v := m.stack[m.top]; v == stackMark {
		// make closure if there are not enough arguments
		tmp := Procedure{
			code: m.code[m.pc-1:],
			env:  m.envClose(),
		}
		raw := kl.MakeRaw(&tmp.scmHead)
		m.stackPush(raw)

		// return to saved address
		savedAddr := m.savedAddr[len(m.savedAddr)-1]
		m.savedAddr = m.savedAddr[:len(m.savedAddr)-1]
		m.code = savedAddr.code
		m.pc = savedAddr.pc
		m.env = savedAddr.env
		m.volatile = m.volatile[:m.envMark]
		m.envMark = savedAddr.envMark
		if enableDebug {
			debugln("GRAB not enough argument")
		}
	} else {
		// grab data from stack to volatile env
		m.volatile = append(m.volatile, v)
		if enableDebug {
			debugf("GRAB %s\n", kl.ObjString(v))
		}
	}
}

func opReturn(m *VM) {
	// stack[top-1] is the result, so should check top-2
	if m.stack[m.top-2] == stackMark {
		savedAddr := m.savedAddr[len(m.savedAddr)-1]
		m.savedAddr = m.savedAddr[:len(m.savedAddr)-1]

		m.code = savedAddr.code
		m.pc = savedAddr.pc
		m.env = savedAddr.env
		m.top--
		m.stack[m.top-1] = m.stack[m.top]
		m.volatile = m.volatile[:m.envMark]
		m.envMark = savedAddr.envMark
		if enableDebug {
			debugf("RETURN %d %d %s\n", m.top, savedAddr.envMark, kl.ObjString(m.stack[m.top-1]))
		}
	} else {
		if enableDebug {
			debugf("RETURN more argument\n")
		}
		// more arguments, continue the beta-reduce.
		// similar to tail apply
		m.top--
		obj := m.stack[m.top]
		// TODO: panic if obj is not a closure
		closure := (*Procedure)(unsafe.Pointer(obj))
		m.code = closure.code
		m.pc = 0
		m.env = closure.env
		m.volatile = m.volatile[:m.envMark]
	}
}

func opTailApply(m *VM) {
	if enableDebug {
		debugln("TAILAPPLY")
	}
	m.top--
	obj := m.stack[m.top]
	// TODO: panic if obj is not a closure
	closure := (*Procedure)(unsafe.Pointer(obj))
	// The only different with Apply is that TailApply doesn't save return address.
	m.code = closure.code
	m.pc = 0
	m.env = closure.env
	m.volatile = m.volatile[:m.envMark]
}

func opApply(m *VM) {
	m.top--
	m.pc++
	if enableDebug {
		debugf("APPLY %d\n", m.top)
	}
	obj := m.stack[m.top]
	// TODO: panic if obj is not a closure
	closure := (*Procedure)(unsafe.Pointer(obj))
	// save return address
	m.savedAddr = append(m.savedAddr, address{m.pc, m.code, m.env, m.envMark})
	// set pc to closure code
	m.code = closure.code
	m.pc = 0
	m.env = closure.env
	m.envMark = len(m.volatile)
}

func opPop(m *VM) {
	m.pc++
	if enableDebug {
		debugln("POP")
	}
	m.top--
}

func opDefun(m *VM) {
	m.pc++
	symbol := m.stack[m.top-1]
	function := m.stack[m.top-2]
	kl.BindSymbolFunc(symbol, function)
	m.top--
	m.stack[m.top-1] = m.stack[m.top]
	if enableDebug {
		debugf("DEFUN %s\n", kl.ObjString(symbol))
	}
}

func opGetF(m *VM) {
	m.pc++
	function := kl.GetSymbolFunc(m.stack[m.top-1])
	if enableDebug {
		debugf("GETF %s\n", kl.GetSymbol(m.stack[m.top-1]))
	}
	if function != nil {
		m.stack[m.top-1] = function
	} else {
		m.stack[m.top-1] = kl.MakeError("unknown function:" + kl.GetSymbol(m.stack[m.top-1]))
		m.status = statusException
	}
}

func opJF(n int) instFunc {
	return func(m *VM) {
		m.pc++
		switch m.stack[m.top-1] {
		case kl.False:
			if enableDebug {
				debugln("JF false")
			}
			m.top--
			m.pc += n
			return
		case kl.True:
			if enableDebug {
				debugln("JF true")
			}
			m.top--
		default:
			// TODO: So what?
			m.stack[m.top-1] = kl.MakeError("test condition need to be boolean")
			m.status = statusException
		}
	}
}

func opJMP(n int) instFunc {
	return func(m *VM) {
		m.pc++
		if enableDebug {
			debugln("JMP")
		}
		m.pc += n
	}
}

func opHalt(m *VM) {
	if enableDebug {
		debugln("HALT")
	}
	m.status = statusHalt
}

func opPrimCall(id int) instFunc {
	return func(m *VM) {
		m.pc++
		prim := kl.GetPrimitiveByID(id)
		args := m.stack[m.top-prim.Required : m.top]

		if enableDebug {
			debugf("PRIMCALL %s\n", prim.Name)
		}

		var result kl.Obj
		// Ugly hack: set function should not be global.
		switch prim.Name {
		case "set":
			result = kl.PrimSet(args[0], args[1])
		case "value":
			result = kl.PrimValue(args[0])
		case "eval-kl":
			tmp := auxVM.Get()
			tmp.nativeFunc = m.nativeFunc
			result = tmp.Eval(args[0])
			auxVM.Put(tmp)
		default:
			result = prim.Function(args...)
		}

		m.stack[m.top-prim.Required] = result
		m.top = m.top - prim.Required + 1
		if kl.IsError(result) {
			m.status = statusException
		}
	}
}

func opNativeCall(arity int) instFunc {
	return func(m *VM) {
		m.pc++
		method := kl.GetSymbol(m.stack[m.top-arity])
		if enableDebug {
			debugf("NativeCall %s\n", method)
		}
		proc, ok := m.nativeFunc[method]
		if !ok {
			m.stack[m.top-1] = kl.MakeError("unknown native function:" + method)
			m.status = statusException
			return
		}
		// Note the invariance arity = len(method + args), so arity-1 = proc.Required
		if arity-1 != proc.Required {
			m.stack[m.top-1] = kl.MakeError("wrong arity for native " + method)
			m.status = statusException
			return
		}
		args := m.stack[m.top-proc.Required : m.top]
		result := proc.Function(args...)

		m.stack[m.top-arity] = result
		m.top = m.top - proc.Required
		if kl.IsError(result) {
			m.status = statusException
		}
	}
}

func (m *VM) envClose() []kl.Obj {
	lenVolatile := len(m.volatile) - m.envMark
	if len(m.env) > 0 || lenVolatile > 0 {
		env := make([]kl.Obj, 0, len(m.env)+lenVolatile)
		env = append(env, m.env...)
		env = append(env, m.volatile[m.envMark:len(m.volatile)]...)
		return env
	}
	return nil
}

func (m *VM) stackPush(o kl.Obj) {
	if m.top == len(m.stack) {
		stack := make([]kl.Obj, len(m.stack)*2)
		copy(stack, m.stack)
		m.stack = stack
	}
	m.stack[m.top] = o
	m.top++
}

func (m *VM) Reset() {
	m.stack = m.stack[:initStackSize]
	m.top = 0
	m.env = m.env[:0]
	m.savedAddr = m.savedAddr[:0]
	m.cc = nil
}

func (m *VM) Debug() {
	debugln("pc:", m.pc)
	debugln("top:", m.top)
	debugln("envMark:", m.envMark)
	debugln("stack:")
	for i := m.top - 1; i >= 0; i-- {
		if m.stack[i] == stackMark {
			debugln("MARK")
		} else {
			debugln(kl.ObjString(m.stack[i]))
		}
	}
}

type Compiler interface {
	KLToSexpByteCode(kl.Obj) kl.Obj
}

var prototype = newVM()
var compiler Compiler

type compileWithEvaluator struct {
	*kl.Evaluator
}

func (cc *compileWithEvaluator) KLToSexpByteCode(klambda kl.Obj) kl.Obj {
	input := kl.Cons(kl.MakeSymbol("kl->bytecode"), kl.Cons(kl.RconsForm(klambda), kl.Nil))
	return cc.Eval(input)
}

func (m *VM) KLToSexpByteCode(klambda kl.Obj) kl.Obj {
	// TODO: Better way to do it?
	// tailcall (kl->bytecode klambda)
	var a Assember
	a.CONST(klambda)
	a.CONST(kl.MakeSymbol("kl->bytecode"))
	a.GetF()
	a.TAILAPPLY()
	a.HALT()

	code := a.Compile()
	return m.Run(code)
}

func klToByteCode(klambda kl.Obj, cc Compiler) (Code, error) {
	bc := cc.KLToSexpByteCode(klambda)
	if kl.IsError(bc) || bc == kl.Nil {
		return nil, errors.New("klToByteCode return some thing wrong:" + kl.ObjString(bc))
	}
	var a Assember
	err := a.FromSexp(bc)
	if err != nil {
		return nil, err
	}
	code := a.Compile()
	return code, nil
}

func (m *VM) Eval(sexp kl.Obj) (res kl.Obj) {
	defer func() {
		if r := recover(); r != nil {
			m.Reset()
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Recovered in Eval:", kl.ObjString(sexp))
			fmt.Println(string(buf[:n]))
			res = kl.MakeError("panic")
		}
	}()

	code, err := klToByteCode(sexp, compiler)
	if err != nil {
		return kl.MakeError(err.Error())
	}
	res = m.Run(code)
	return
}

func BootstrapKL() {
	evaluator := kl.NewEvaluator()
	evaluator.Silence = true
	evaluator.LoadFile("compiler/primitive.kl")
	evaluator.LoadFile("compiler/de-bruijn.kl")
	evaluator.LoadFile("compiler/compile.kl")
	compiler = &compileWithEvaluator{evaluator}
}

func BootstrapMin() {
	prototype.RegistNativeCall("primitive?", 1, kl.NativeIsPrimitive)
	prototype.RegistNativeCall("primitive-arity", 1, kl.NativePrimitiveArity)
	prototype.RegistNativeCall("primitive-id", 1, kl.NativePrimitiveID)

	prototype.mustLoadBytecode(kl.MakeString("primitive.bc"))
	prototype.mustLoadBytecode(kl.MakeString("de-bruijn.bc"))
	prototype.mustLoadBytecode(kl.MakeString("compile.bc"))
	compiler = prototype
}

func BootstrapCora() {
	BootstrapMin()
	prototype.mustLoadBytecode(kl.MakeString("toplevel.bc"))
	prototype.mustLoadBytecode(kl.MakeString("core.bc"))
	prototype.mustLoadBytecode(kl.MakeString("sys.bc"))
	prototype.mustLoadBytecode(kl.MakeString("yacc.bc"))
	prototype.mustLoadBytecode(kl.MakeString("reader.bc"))
	prototype.mustLoadBytecode(kl.MakeString("track.bc"))
	prototype.mustLoadBytecode(kl.MakeString("load.bc"))
	prototype.mustLoadBytecode(kl.MakeString("writer.bc"))
	prototype.mustLoadBytecode(kl.MakeString("macros.bc"))
	prototype.mustLoadBytecode(kl.MakeString("declarations.bc"))
}

func BootstrapShen() {
	BootstrapCora()
	prototype.mustLoadBytecode(kl.MakeString("sequent.bc"))
	prototype.mustLoadBytecode(kl.MakeString("prolog.bc"))
	prototype.mustLoadBytecode(kl.MakeString("t-star.bc"))
	prototype.mustLoadBytecode(kl.MakeString("types.bc"))
}

var Boot string

func (m *VM) mustLoadBytecode(args ...kl.Obj) {
	res := m.loadBytecode(args...)
	if kl.IsError(res) {
		panic(kl.ObjString(res))
	}
}

func (m *VM) loadBytecode(args ...kl.Obj) kl.Obj {
	fileName := kl.GetString(args[0])
	var f io.ReadCloser
	var err error
	if Boot != "" {
		filePath := path.Join(Boot, "bytecode", fileName)
		f, err = os.Open(filePath)
	} else {
		filePath := path.Join("/bytecode", fileName)
		f, err = FS(false).Open(filePath)
	}
	if err != nil {
		return kl.MakeError(err.Error())
	}
	defer f.Close()

	r := kl.NewSexpReader(f)
	obj, err := r.Read()
	for err == nil {
		var a Assember
		if err1 := a.FromSexp(obj); err1 != nil {
			return kl.MakeError(err1.Error())
		}
		code := a.Compile()

		tmp := auxVM.Get()
		tmp.nativeFunc = m.nativeFunc
		res := tmp.Run(code)
		auxVM.Put(tmp)

		if kl.IsError(res) {
			return res
		}
		obj, err = r.Read()
	}
	if err != io.EOF {
		return kl.MakeError(err.Error())
	}
	return args[0]
}

func (m *VM) loadFile(args ...kl.Obj) kl.Obj {
	file := kl.GetString(args[0])
	var filePath string
	if _, err := os.Stat(file); err == nil {
		filePath = file
	} else {
		filePath = path.Join(kl.PackagePath(), file)
		if _, err := os.Stat(filePath); err != nil {
			return kl.MakeError(err.Error())
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return kl.MakeError(err.Error())
	}
	defer f.Close()

	r := kl.NewSexpReader(f)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return kl.MakeError(err.Error())
			}
			break
		}

		tmp := auxVM.Get()
		tmp.nativeFunc = m.nativeFunc
		res := tmp.Eval(exp)
		auxVM.Put(tmp)

		if kl.IsError(res) {
			return res
		}
	}
	return args[0]
}

func (m *VM) loadPlugin(args ...kl.Obj) kl.Obj {
	pluginPath := kl.GetString(args[0])
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return kl.MakeError(err.Error())
	}

	entry, err := p.Lookup("Main")
	if err != nil {
		return kl.MakeError(err.Error())
	}

	f, ok := entry.(func(*VM))
	if !ok {
		return kl.MakeError("plugin Main should be func(*vm.VM)")
	}

	f(m)
	return args[0]
}

// (native plugin-bind "/path/to/plugin.so" [bit-shift 2 "BitLeftShift"])
func (m *VM) pluginBind(args ...kl.Obj) kl.Obj {
	pluginPath := kl.GetString(args[0])
	bindInfo := getBindInfo(args[1])
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return kl.MakeError(err.Error())
	}

	for _, info := range bindInfo {
		f, err := p.Lookup(info.PluginFunc)
		if err != nil {
			return kl.MakeError(err.Error())
		}
		if funcAddr, ok := f.(func(...kl.Obj) kl.Obj); !ok {
			return kl.MakeError(fmt.Sprintf("func %s signature is illeagel", info.PluginFunc))
		} else {
			m.RegistNativeCall(info.Name, info.Arity, funcAddr)
		}
	}
	return args[0]
}

type bindInfo struct {
	Name       string
	Arity      int
	PluginFunc string
}

func getBindInfo(l kl.Obj) []bindInfo {
	ret := make([]bindInfo, 0, 1)
	for l != kl.Nil {
		name := kl.GetSymbol(kl.Car(l))
		l = kl.Cdr(l)
		arity := kl.GetInteger(kl.Car(l))
		l = kl.Cdr(l)
		pluginName := kl.GetString(kl.Car(l))
		l = kl.Cdr(l)
		ret = append(ret, bindInfo{name, arity, pluginName})
	}
	return ret
}

var stdDebug io.Writer

func init() {
	if enableDebug {
		stdDebug, _ = os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	}
}

func debugf(format string, a ...interface{}) {
	fmt.Fprintf(stdDebug, format, a...)
}

func debugln(a ...interface{}) {
	fmt.Fprintln(stdDebug, a...)
}
