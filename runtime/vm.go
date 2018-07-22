package runtime

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
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
	stack []Obj
	pc    int // pc register refer to the position in current code
	sp    int // stack top
	fp    int // frame pos

	// Calling convention:
	//
	// [Mark] [Arg3] [Arg2] [Arg1] [Apply]
	//
	// [saved address]
	// [saved fp]  <-- fp
	// [arg3]
	// [arg2]
	// [arg1] <- sp

	env      []Obj // persistent environment
	volatile []Obj // volatile environment
	ebp      int   // volatile[ebp: len(volatile)] is current env

	// jumpBuf is used to implement exception, similar to setjmp/longjmp in C.
	cc []jumpBuf

	code   []instFunc
	status statusType
}

type jumpBuf struct {
	address
	savedStackTop int
	savedFp       int
	savedEnvTop   int
	closure       Obj
}

// address is the information to be saved before apply a closure.
type address struct {
	scmHead
	pc   int
	code Code
	env  []Obj
	ebp  int
}

func makeAddress(pc int, code Code, env []Obj, ebp int) Obj {
	tmp := &address{
		scmHead: scmHeadAddress,
		pc:      pc,
		code:    code,
		env:     env,
		ebp:     ebp,
	}
	return makeObj(&tmp.scmHead)
}

func mustAddress(o Obj) *address {
	if objType(o) != scmHeadAddress {
		panic("mustAddress")
	}
	return (*address)(unsafe.Pointer(o))
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
var nativeFunc = make(map[string]*scmPrimitive)

func NewVM() *VM {
	m := newVM()
	return m
}

func newVM() *VM {
	m := &VM{
		stack:    make([]Obj, initStackSize),
		env:      make([]Obj, 0, 256),
		volatile: make([]Obj, 0, 256),
	}

	return m
}

func RegistNativeCall(name string, arity int, f func(...Obj) Obj) {
	nativeFunc[name] = makePrimitive(name, arity, f)
}

func (m *VM) Run(code Code) Obj {
	m.setup(code)

	// May be the fastest dispatch method in Go.
Dispatch:
	m.status = statusNormal
	for m.status == statusNormal {
		m.code[m.pc](m)
	}

	if m.status == statusException {
		if len(m.cc) == 0 {
			err := m.stack[m.sp-1]
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
	m.stackPush(makeAddress(len(code)-1, code, nil, 0))
	m.stackPush(MakeInteger(m.fp))
	m.fp = m.sp - 1
}

func (m *VM) done() Obj {
	if m.sp != 1 || m.ebp != 0 {
		if enableDebug {
			m.Debug()
		}
		panic("m in wrong status")
	}
	m.sp--
	return m.stack[m.sp]
}

func (m *VM) handleException() {
	// clear jmpBuf
	jmpBuf := m.cc[len(m.cc)-1]
	m.cc = m.cc[:len(m.cc)-1]
	// pop trap-error handler, prepare for call.
	value := m.stack[m.sp-1]
	// recover the snapshot
	m.fp = jmpBuf.savedFp
	m.sp = jmpBuf.savedStackTop
	// make the trap-error handler execute environment
	m.stackPush(nil)
	m.stackPush(MakeInteger(m.fp))
	m.fp = m.sp - 1
	m.stackPush(value)
	// recover savedAddr, for the trap-error handler to return
	m.stack[m.fp-1] = makeObj(&jmpBuf.address.scmHead)
	// longjmp... tail apply
	closure := mustClosure(jmpBuf.closure)
	m.code = closure.code
	m.pc = 0
	m.env = closure.env
	m.ebp = jmpBuf.savedEnvTop
	m.volatile = m.volatile[:m.ebp]
}

func opSetJmp(n int) instFunc {
	return func(m *VM) {
		m.pc++
		if enableDebug {
			debugf("SETJMP\n")
		}

		m.sp--
		cc := jumpBuf{
			address: address{
				scmHead: scmHeadAddress,
				pc:      m.pc + n,
				code:    m.code,
				env:     m.env,
				ebp:     m.ebp,
			},
			savedStackTop: m.sp,
			savedFp:       m.fp,
			closure:       m.stack[m.sp],
			savedEnvTop:   len(m.volatile),
		}
		m.cc = append(m.cc, cc)
	}
}

func opClearJmp(m *VM) {
	m.pc++
	m.cc = m.cc[:len(m.cc)-1]
}

func opConst(o Obj) instFunc {
	return func(m *VM) {
		m.pc++
		m.stackPush(o)
		if enableDebug {
			debugf("CONST %s\n", ObjString(o))
		}
	}
}

func opAccess(n int) instFunc {
	return func(m *VM) {
		m.pc++
		if n+m.ebp < len(m.volatile) {
			// get value from volatile environment
			v := m.volatile[len(m.volatile)-1-n]
			m.stackPush(v)
			if enableDebug {
				debugf("ACCESS %d, get %s\n", n, ObjString(v))
			}
		} else {
			// get value from persistent environment
			xx := n - (len(m.volatile) - m.ebp)
			v := m.env[len(m.env)-1-xx]
			m.stackPush(v)
			if enableDebug {
				debugf("ACCESS %d from env, get %s\n", n, ObjString(v))
			}
		}
	}
}

func opFreeze(n int) instFunc {
	return func(m *VM) {
		m.pc++
		// create closure directly
		// nearly the same with grab, but if need zero arguments.
		env := m.envClose()
		proc := makeClosure(m.code[m.pc:], env)
		if enableDebug {
			debugf("FREEZE len(env)=%d\n", len(env))
		}
		m.stackPush(proc)
		m.pc += n
	}
}

func opMark(m *VM) {
	m.pc++
	m.stackPush(nil) // reserve for return addr
	m.stackPush(MakeInteger(m.fp))
	m.fp = m.sp - 1
	if enableDebug {
		debugln("MARK", m.fp)
	}
}

func opGrab(m *VM) {
	m.pc++
	m.sp--
	v := m.stack[m.sp]
	if m.sp == m.fp {
		// make closure if there are not enough arguments
		oldFp := mustInteger(v)
		m.sp--
		savedAddr := mustAddress(m.stack[m.sp])
		proc := makeClosure(m.code[m.pc-1:], m.envClose())
		m.stackPush(proc)

		// return to saved address
		m.fp = oldFp
		m.code = savedAddr.code
		m.pc = savedAddr.pc
		m.env = savedAddr.env
		m.volatile = m.volatile[:m.ebp]
		m.ebp = savedAddr.ebp
		if enableDebug {
			debugln("GRAB not enough argument")
		}
	} else {
		// grab data from stack to volatile env
		m.volatile = append(m.volatile, v)
		if enableDebug {
			debugf("GRAB %s\n", ObjString(v))
		}
	}
}

func opReturn(m *VM) {
	// stack[sp-1] is the result, so should check sp-2
	if m.sp-2 == m.fp {
		savedAddr := mustAddress(m.stack[m.sp-3])
		m.code = savedAddr.code
		m.pc = savedAddr.pc
		m.env = savedAddr.env
		m.sp--
		m.fp = mustInteger(m.stack[m.sp-1])
		m.stack[m.sp-2] = m.stack[m.sp]
		m.sp--
		m.volatile = m.volatile[:m.ebp]
		m.ebp = savedAddr.ebp
		if enableDebug {
			debugf("RETURN %d %d %s\n", m.sp, savedAddr.ebp, ObjString(m.stack[m.sp-1]))
		}
	} else {
		if enableDebug {
			debugf("RETURN more argument\n")
		}
		// more arguments, continue the beta-reduce.
		// similar to tail apply
		m.sp--
		obj := m.stack[m.sp]
		closure := mustClosure(obj)
		m.code = closure.code
		m.pc = 0
		m.env = closure.env
		m.volatile = m.volatile[:m.ebp]
	}
}

func opTailApply(m *VM) {
	if enableDebug {
		debugln("TAILAPPLY")
	}
	m.sp--
	obj := m.stack[m.sp]
	closure := mustClosure(obj)
	// The only different with Apply is that TailApply doesn't save return address.
	m.code = closure.code
	m.pc = 0
	m.env = closure.env
	m.volatile = m.volatile[:m.ebp]
}

func opApply(m *VM) {
	m.sp--
	m.pc++
	if enableDebug {
		debugf("APPLY %d\n", m.sp)
	}
	obj := m.stack[m.sp]
	closure := mustClosure(obj)
	// save return address
	m.stack[m.fp-1] = makeAddress(m.pc, m.code, m.env, m.ebp)
	// set pc to closure code
	m.code = closure.code
	m.pc = 0
	m.env = closure.env
	m.ebp = len(m.volatile)
}

func opPop(m *VM) {
	m.pc++
	if enableDebug {
		debugln("POP")
	}
	m.sp--
}

func opDefun(m *VM) {
	m.pc++
	symbol := m.stack[m.sp-1]
	function := m.stack[m.sp-2]
	bindSymbolFunc(symbol, function)
	m.sp--
	m.stack[m.sp-1] = m.stack[m.sp]
	if enableDebug {
		debugf("DEFUN %s\n", ObjString(symbol))
	}
}

func opGetF(m *VM) {
	m.pc++
	function := GetSymbolFunc(m.stack[m.sp-1])
	if enableDebug {
		debugf("GETF %s\n", GetSymbol(m.stack[m.sp-1]))
	}
	if function != nil {
		m.stack[m.sp-1] = function
	} else {
		m.stack[m.sp-1] = MakeError("unknown function:" + GetSymbol(m.stack[m.sp-1]))
		m.status = statusException
	}
}

func opJF(n int) instFunc {
	return func(m *VM) {
		m.pc++
		switch m.stack[m.sp-1] {
		case False:
			if enableDebug {
				debugln("JF false")
			}
			m.sp--
			m.pc += n
			return
		case True:
			if enableDebug {
				debugln("JF true")
			}
			m.sp--
		default:
			// TODO: So what?
			m.stack[m.sp-1] = MakeError("test condition need to be boolean")
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

func opNativeCall(arity int) instFunc {
	return func(m *VM) {
		m.pc++
		method := mustString(m.stack[m.sp-arity])
		if enableDebug {
			debugf("NativeCall %s\n", method)
		}
		proc, ok := nativeFunc[method]
		if !ok {
			m.stack[m.sp-1] = MakeError("unknown native function:" + method)
			m.status = statusException
			return
		}
		// Note the invariance arity = len(method + args), so arity-1 = proc.Required
		if arity-1 != proc.Required {
			m.stack[m.sp-1] = MakeError("wrong arity for native " + method)
			m.status = statusException
			return
		}
		args := m.stack[m.sp-proc.Required : m.sp]
		result := proc.Function(args...)

		m.stack[m.sp-arity] = result
		m.sp = m.sp - proc.Required
		if IsError(result) {
			m.status = statusException
		}
	}
}

func (m *VM) envClose() []Obj {
	lenVolatile := len(m.volatile) - m.ebp
	if len(m.env) > 0 || lenVolatile > 0 {
		env := make([]Obj, 0, len(m.env)+lenVolatile)
		env = append(env, m.env...)
		env = append(env, m.volatile[m.ebp:len(m.volatile)]...)
		return env
	}
	return nil
}

func (m *VM) stackPush(o Obj) {
	if m.sp == len(m.stack) {
		stack := make([]Obj, len(m.stack)*2)
		copy(stack, m.stack)
		m.stack = stack
	}
	m.stack[m.sp] = o
	m.sp++
}

func (m *VM) Reset() {
	m.stack = m.stack[:initStackSize]
	m.sp = 0
	m.fp = 0
	m.env = nil
	m.cc = nil
}

func (m *VM) Debug() {
	debugln("pc:", m.pc)
	debugln("sp:", m.sp)
	debugln("fp:", m.fp)
	debugln("ebp:", m.ebp)
	debugln("stack:")
	for i := m.sp - 1; i >= 0; i-- {
		debugln(ObjString(m.stack[i]))
	}
}

var prototype *VM

func klToSexpByteCode(klambda Obj) Obj {
	// TODO: Better way to do it?
	// tailcall (kl->bytecode klambda)
	var a assember
	a.CONST(klambda)
	a.CONST(MakeSymbol("kl->bytecode"))
	a.GetF()
	a.TAILAPPLY()
	a.HALT()

	code := a.Compile()
	return prototype.Run(code)
}

func klToCode(klambda Obj) (Code, error) {
	bc := klToSexpByteCode(klambda)
	if IsError(bc) || bc == Nil {
		return nil, errors.New("klToByteCode return some thing wrong:" + ObjString(bc))
	}
	var a assember
	err := a.FromSexp(bc)
	if err != nil {
		return nil, err
	}
	code := a.Compile()
	return code, nil
}

func (m *VM) Eval(sexp Obj) (res Obj) {
	defer func() {
		if r := recover(); r != nil {
			m.Reset()
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			fmt.Println("Recovered in Eval:", ObjString(sexp))
			fmt.Println(string(buf[:n]))
			res = MakeError("panic")
		}
	}()

	code, err := klToCode(sexp)
	if err != nil {
		return MakeError(err.Error())
	}
	res = m.Run(code)
	return
}

func shenRead(input string) (res Obj) {
	r := strings.NewReader(input)
	stream := MakeStream(r)
	sexp := Cons(MakeSymbol("read"), Cons(stream, Nil))
	return Eval(sexp)
}

func shenEval(obj Obj) (res Obj) {
	sexp := Cons(MakeSymbol("eval"), Cons(obj, Nil))
	return Eval(sexp)
}

func EvalString(input string) Obj {
	obj := shenRead(input)
	if IsError(obj) {
		return obj
	}
	return Eval(obj)
}

func Eval(sexp Obj) Obj {
	vm := auxVM.Get()
	defer auxVM.Put(vm)
	return vm.Eval(sexp)
}

func BootstrapMin() {
	prototype.mustLoadBytecode(MakeString("primitive.bc"))
	prototype.mustLoadBytecode(MakeString("de-bruijn.bc"))
	prototype.mustLoadBytecode(MakeString("compile.bc"))
	prototype.mustLoadBytecode(MakeString("core.bc"))
	prototype.mustLoadBytecode(MakeString("sys.bc"))
	prototype.mustLoadBytecode(MakeString("yacc.bc"))
	prototype.mustLoadBytecode(MakeString("override.bc"))
}

func BootstrapCora() {
	prototype.mustLoadBytecode(MakeString("primitive.bc"))
	prototype.mustLoadBytecode(MakeString("de-bruijn.bc"))
	prototype.mustLoadBytecode(MakeString("compile.bc"))
	prototype.mustLoadBytecode(MakeString("core.bc"))
	prototype.mustLoadBytecode(MakeString("sys.bc"))
	prototype.mustLoadBytecode(MakeString("yacc.bc"))
	prototype.mustLoadBytecode(MakeString("toplevel.bc"))
	prototype.mustLoadBytecode(MakeString("reader.bc"))
	prototype.mustLoadBytecode(MakeString("track.bc"))
	prototype.mustLoadBytecode(MakeString("load.bc"))
	prototype.mustLoadBytecode(MakeString("writer.bc"))
	prototype.mustLoadBytecode(MakeString("macros.bc"))
	prototype.mustLoadBytecode(MakeString("declarations.bc"))
}

func BootstrapShen() {
	BootstrapMin()
	prototype.mustLoadBytecode(MakeString("toplevel.bc"))
	prototype.mustLoadBytecode(MakeString("dict.bc"))
	prototype.mustLoadBytecode(MakeString("sequent.bc"))
	prototype.mustLoadBytecode(MakeString("reader.bc"))
	prototype.mustLoadBytecode(MakeString("prolog.bc"))
	prototype.mustLoadBytecode(MakeString("track.bc"))
	prototype.mustLoadBytecode(MakeString("load.bc"))
	prototype.mustLoadBytecode(MakeString("writer.bc"))
	prototype.mustLoadBytecode(MakeString("macros.bc"))
	prototype.mustLoadBytecode(MakeString("declarations.bc"))
	prototype.mustLoadBytecode(MakeString("t-star.bc"))
	prototype.mustLoadBytecode(MakeString("types.bc"))
}

var Boot string

func (m *VM) mustLoadBytecode(args ...Obj) {
	res := loadBytecode(args...)
	if IsError(res) {
		panic(ObjString(res) + ObjString(args[0]))
	}
}

func loadBytecode(args ...Obj) Obj {
	fileName := GetString(args[0])
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
		return MakeError(err.Error())
	}
	defer f.Close()

	r := NewSexpReader(f)
	obj, err := r.Read()
	for err == nil {
		var a assember
		if err1 := a.FromSexp(obj); err1 != nil {
			return MakeError(err1.Error())
		}
		code := a.Compile()

		tmp := auxVM.Get()
		res := tmp.Run(code)
		auxVM.Put(tmp)

		if IsError(res) {
			return res
		}
		obj, err = r.Read()
	}
	if err != io.EOF {
		return MakeError(err.Error())
	}
	return args[0]
}

var stdDebug io.Writer

func debugf(format string, a ...interface{}) {
	fmt.Fprintf(stdDebug, format, a...)
}

func debugln(a ...interface{}) {
	fmt.Fprintln(stdDebug, a...)
}

func initSymbolTable() {
	dir, _ := os.Getwd()
	primSet(MakeSymbol("*stinput*"), MakeStream(os.Stdin))
	primSet(MakeSymbol("*stoutput*"), MakeStream(os.Stdout))
	primSet(MakeSymbol("*home-directory*"), MakeString(dir))
	primSet(MakeSymbol("*language*"), MakeString("Go"))
	primSet(MakeSymbol("*implementation*"), MakeString("bytecode"))
	primSet(MakeSymbol("*relase*"), MakeString(runtime.Version()))
	primSet(MakeSymbol("*os*"), MakeString(runtime.GOOS))
	primSet(MakeSymbol("*porters*"), MakeString("Arthur Mao"))
	primSet(MakeSymbol("*port*"), MakeString("0.0.1"))
	primSet(MakeSymbol("*package-path*"), MakeString(PackagePath()))
}

func init() {
	uptime = time.Now()
	tmp1 := &scmBoolean{scmHeadBoolean, false}
	False = makeObj(&tmp1.scmHead)

	tmp2 := &scmBoolean{scmHeadBoolean, true}
	True = makeObj(&tmp2.scmHead)

	tmp3 := &scmPair{scmHeadNull, nil, nil}
	Nil = makeObj(&tmp3.scmHead)

	var tmp4 int
	undefined = MakeRaw(&tmp4)

	for i := 0; i < intConstCount; i++ {
		intConst[i] = makeInteger(i)
	}

	symbolArray = make([]symbolArrayObj, 0, 4096)
	trieRoot = &trieNode{}

	if enableDebug {
		stdDebug, _ = os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	}

	initSymbolTable()

	for _, v := range allPrimitives {
		RegistNativeCall(v.Name, v.Required, v.Function)
	}

	prototype = newVM()
}
