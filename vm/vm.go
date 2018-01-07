package vm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"unsafe"

	"github.com/tiancaiamao/shen-go/kl"
)

// Go doesn't provide MACRO like C, but the compiler optimization can eliminate dead code "if false XX".
const enableDebug = false

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
}

type jumpBuf struct {
	address
	savedAddrPos  int
	savedStackTop int
	savedEnvTop   int
	closure       kl.Obj
}

// Code is something executable to VM, it's immutable.
type Code struct {
	bc     []instruction
	consts []kl.Obj
}

// address is the information to be saved before apply a closure.
type address struct {
	pc      int
	code    *Code
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
	vm := newVM()
	vm.RegistNativeCall("load-bytecode", 1, vm.loadBytecode)
	vm.RegistNativeCall("load-file", 1, vm.loadFile)
	vm.RegistNativeCall("primitive?", 1, kl.NativeIsPrimitive)
	vm.RegistNativeCall("primitive-arity", 1, kl.NativePrimitiveArity)
	vm.RegistNativeCall("primitive-id", 1, kl.NativePrimitiveID)
	return vm
}

func newVM() *VM {
	vm := &VM{
		stack:      make([]kl.Obj, initStackSize),
		env:        make([]kl.Obj, 0, 256),
		volatile:   make([]kl.Obj, 0, 256),
		nativeFunc: make(map[string]*kl.ScmPrimitive),
	}
	initSymbolTable()
	return vm
}

func (vm *VM) RegistNativeCall(name string, arity int, f func(...kl.Obj) kl.Obj) {
	vm.nativeFunc[name] = kl.MakePrimitive(name, arity, f)
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

func (vm *VM) Run(code *Code) kl.Obj {
	vm.pc = 0
	vm.stackPush(stackMark)
	vm.savedAddr = append(vm.savedAddr, address{pc: len(code.bc) - 1, code: code})

	halt := false
	for !halt {
		inst := code.bc[vm.pc]
		vm.pc++

		exception := false
		switch instructionCode(inst) {
		case iSetJmp:
			if enableDebug {
				debugf("SETJMP\n")
			}
			n := instructionOPN(inst)
			vm.top--
			cc := jumpBuf{
				address: address{
					pc:      vm.pc + n,
					code:    code,
					env:     vm.env,
					envMark: vm.envMark,
				},
				savedAddrPos:  len(vm.savedAddr),
				savedStackTop: vm.top,
				closure:       vm.stack[vm.top],
				savedEnvTop:   len(vm.volatile),
			}
			vm.cc = append(vm.cc, cc)
		case iClearJmp:
			vm.cc = vm.cc[:len(vm.cc)-1]
		case iConst:
			n := instructionOPN(inst)
			vm.stackPush(code.consts[n])
			if enableDebug {
				debugf("CONST %s\n", kl.ObjString(code.consts[n]))
			}
		case iAccess:
			n := instructionOPN(inst)
			if n+vm.envMark < len(vm.volatile) {
				// get value from volatile environment
				v := vm.volatile[len(vm.volatile)-1-n]
				vm.stackPush(v)
				if enableDebug {
					debugf("ACCESS %d, get %s\n", n, kl.ObjString(v))
				}
			} else {
				// get value from persistent environment
				n -= (len(vm.volatile) - vm.envMark)
				v := vm.env[len(vm.env)-1-n]
				vm.stackPush(v)
				if enableDebug {
					debugf("ACCESS %d from env, get %s\n", n, kl.ObjString(v))
				}
			}
		case iFreeze:
			// create closure directly
			// nearly the same with grab, but if need zero arguments.
			n := instructionOPN(inst)
			tmp := &Procedure{
				code: Code{
					bc:     code.bc[vm.pc:],
					consts: code.consts,
				},
				env: vm.envClose(),
			}
			raw := kl.MakeRaw(&tmp.scmHead)
			if enableDebug {
				debugf("FREEZE len(env)=%d\n", len(tmp.env))
			}
			vm.stackPush(raw)
			vm.pc += n
		case iMark:
			if enableDebug {
				debugln("MARK")
			}
			vm.stackPush(stackMark)
		case iGrab:
			vm.top--
			if v := vm.stack[vm.top]; v == stackMark {
				// make closure if there are not enough arguments
				tmp := Procedure{
					code: Code{
						bc:     code.bc[vm.pc-1:],
						consts: code.consts,
					},
					env: vm.envClose(),
				}
				raw := kl.MakeRaw(&tmp.scmHead)
				vm.stackPush(raw)

				// return to saved address
				savedAddr := vm.savedAddr[len(vm.savedAddr)-1]
				vm.savedAddr = vm.savedAddr[:len(vm.savedAddr)-1]
				code = savedAddr.code
				vm.pc = savedAddr.pc
				vm.env = savedAddr.env
				vm.volatile = vm.volatile[:vm.envMark]
				vm.envMark = savedAddr.envMark
				if enableDebug {
					debugln("GRAB not enough argument")
				}
			} else {
				// grab data from stack to volatile env
				vm.volatile = append(vm.volatile, v)
				if enableDebug {
					debugf("GRAB %s\n", kl.ObjString(v))
				}
			}
		case iReturn:
			// stack[top-1] is the result, so should check top-2
			if vm.stack[vm.top-2] == stackMark {
				savedAddr := vm.savedAddr[len(vm.savedAddr)-1]
				vm.savedAddr = vm.savedAddr[:len(vm.savedAddr)-1]

				code = savedAddr.code
				vm.pc = savedAddr.pc
				vm.env = savedAddr.env
				vm.top--
				vm.stack[vm.top-1] = vm.stack[vm.top]
				vm.volatile = vm.volatile[:vm.envMark]
				vm.envMark = savedAddr.envMark
				if enableDebug {
					debugf("RETURN %d %d\n", vm.top, savedAddr.envMark)
				}
			} else {
				if enableDebug {
					debugf("RETURN more argument\n")
				}
				// more arguments, continue the beta-reduce.
				// similar to tail apply
				vm.top--
				obj := vm.stack[vm.top]
				// TODO: panic if obj is not a closure
				closure := (*Procedure)(unsafe.Pointer(obj))
				code = &closure.code
				vm.pc = 0
				vm.env = closure.env
				vm.volatile = vm.volatile[:vm.envMark]
			}
		case iTailApply:
			if enableDebug {
				debugln("TAILAPPLY")
			}
			vm.top--
			obj := vm.stack[vm.top]
			// TODO: panic if obj is not a closure
			closure := (*Procedure)(unsafe.Pointer(obj))
			// The only different with Apply is that TailApply doesn't save return address.
			code = &closure.code
			vm.pc = 0
			vm.env = closure.env
			vm.volatile = vm.volatile[:vm.envMark]
		case iApply:
			vm.top--
			if enableDebug {
				debugf("APPLY %d\n", vm.top)
			}
			obj := vm.stack[vm.top]
			// TODO: panic if obj is not a closure
			closure := (*Procedure)(unsafe.Pointer(obj))
			// save return address
			vm.savedAddr = append(vm.savedAddr, address{vm.pc, code, vm.env, vm.envMark})
			// set pc to closure code
			code = &closure.code
			vm.pc = 0
			vm.env = closure.env
			vm.envMark = len(vm.volatile)
		case iPop:
			if enableDebug {
				debugln("POP")
			}
			vm.top--
		case iDefun:
			symbol := vm.stack[vm.top-1]
			function := vm.stack[vm.top-2]
			kl.BindSymbolFunc(symbol, function)
			vm.top--
			vm.stack[vm.top-1] = vm.stack[vm.top]
			if enableDebug {
				debugf("DEFUN %s\n", symbol)
			}
		case iGetF:
			function := kl.GetSymbolFunc(vm.stack[vm.top-1])
			if function != nil {
				vm.stack[vm.top-1] = function
			} else {
				vm.stack[vm.top-1] = kl.MakeError("unknown function:" + kl.GetSymbol(vm.stack[vm.top-1]))
				exception = true
			}
			if enableDebug {
				debugf("GETF %s\n", kl.GetSymbol(vm.stack[vm.top-1]))
			}
		case iJF:
			switch vm.stack[vm.top-1] {
			case kl.False:
				if enableDebug {
					debugln("JF false")
				}
				n := instructionOPN(inst)
				vm.top--
				vm.pc += n
			case kl.True:
				if enableDebug {
					debugln("JF true")
				}
				vm.top--
			default:
				// TODO: So what?
				vm.stack[vm.top-1] = kl.MakeError("test condition need to be boolean")
				exception = true
			}
		case iJMP:
			if enableDebug {
				debugln("JMP")
			}
			n := instructionOPN(inst)
			vm.pc += n
		case iHalt:
			if enableDebug {
				debugln("HALT")
			}
			halt = true
		case iPrimCall:
			id := instructionOPN(inst)
			prim := kl.GetPrimitiveByID(id)
			args := vm.stack[vm.top-prim.Required : vm.top]

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
				tmp.nativeFunc = vm.nativeFunc
				result = tmp.Eval(args[0])
				auxVM.Put(tmp)
			default:
				result = prim.Function(args...)
			}

			vm.stack[vm.top-prim.Required] = result
			vm.top = vm.top - prim.Required + 1
			if kl.IsError(result) {
				exception = true
			}
		case iNativeCall:
			arity := instructionOPN(inst)
			method := kl.GetSymbol(vm.stack[vm.top-arity])
			if enableDebug {
				debugf("NativeCall %s\n", method)
			}
			proc, ok := vm.nativeFunc[method]
			if !ok {
				vm.stack[vm.top-1] = kl.MakeError("unknown native function:" + method)
				exception = true
				break
			}
			// Note the invariance arity = len(method + args), so arity-1 = proc.Required
			if arity-1 != proc.Required {
				vm.stack[vm.top-1] = kl.MakeError("wrong arity for native " + method)
				exception = true
				break
			}
			args := vm.stack[vm.top-proc.Required : vm.top]
			result := proc.Function(args...)

			vm.stack[vm.top-arity] = result
			vm.top = vm.top - proc.Required
			if kl.IsError(result) {
				exception = true
			}
		default:
			panic(fmt.Sprintf("unknown instruction %d", inst))
		}

		if exception {
			if len(vm.cc) == 0 {
				err := vm.stack[vm.top-1]
				vm.Reset()
				return err
			}

			// clear jmpBuf
			jmpBuf := vm.cc[len(vm.cc)-1]
			vm.cc = vm.cc[:len(vm.cc)-1]
			// pop trap-error handler, prepare for call.
			value := vm.stack[vm.top-1]
			vm.top = jmpBuf.savedStackTop
			vm.stackPush(stackMark)
			vm.stackPush(value)
			// recover savedAddr
			vm.savedAddr = vm.savedAddr[:jmpBuf.savedAddrPos]
			vm.savedAddr = append(vm.savedAddr, jmpBuf.address)
			// longjmp... tail apply
			closure := (*Procedure)(unsafe.Pointer(jmpBuf.closure))
			code = &closure.code
			vm.pc = 0
			vm.env = closure.env
			vm.envMark = jmpBuf.savedEnvTop
			vm.volatile = vm.volatile[:vm.envMark]
		}
	}

	if vm.top != 1 || len(vm.savedAddr) != 0 || vm.envMark != 0 {
		if enableDebug {
			vm.Debug()
		}
		panic("vm in wrong status")
	}
	vm.top--
	return vm.stack[vm.top]
}

func (vm *VM) envClose() []kl.Obj {
	lenVolatile := len(vm.volatile) - vm.envMark
	if len(vm.env) > 0 || lenVolatile > 0 {
		env := make([]kl.Obj, 0, len(vm.env)+lenVolatile)
		env = append(env, vm.env...)
		env = append(env, vm.volatile[vm.envMark:len(vm.volatile)]...)
		return env
	}
	return nil
}

func (vm *VM) stackPush(o kl.Obj) {
	if vm.top == len(vm.stack) {
		stack := make([]kl.Obj, len(vm.stack)*2)
		copy(stack, vm.stack)
		vm.stack = stack
	}
	vm.stack[vm.top] = o
	vm.top++
}

func (vm *VM) Reset() {
	vm.stack = vm.stack[:initStackSize]
	vm.top = 0
	vm.env = vm.env[:0]
	vm.savedAddr = vm.savedAddr[:0]
	vm.cc = nil
}

func (vm *VM) Debug() {
	debugln("pc:", vm.pc)
	debugln("top:", vm.top)
	debugln("envMark:", vm.envMark)
	debugln("stack:")
	for i := vm.top - 1; i >= 0; i-- {
		if vm.stack[i] == stackMark {
			debugln("MARK")
		} else {
			debugln(kl.ObjString(vm.stack[i]))
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

func (vm *VM) KLToSexpByteCode(klambda kl.Obj) kl.Obj {
	// TODO: Better way to do it?
	// tailcall (kl->bytecode klambda)
	var a Assember
	a.CONST(klambda)
	a.CONST(kl.MakeSymbol("kl->bytecode"))
	a.GetF()
	a.TAILAPPLY()
	a.HALT()

	code := a.Comiple()
	return vm.Run(code)
}

func klToByteCode(klambda kl.Obj, cc Compiler) (*Code, error) {
	bc := cc.KLToSexpByteCode(klambda)
	if kl.IsError(bc) || bc == kl.Nil {
		return nil, errors.New("klToByteCode return some thing wrong:" + kl.ObjString(bc))
	}
	var a Assember
	err := a.FromSexp(bc)
	if err != nil {
		return nil, err
	}
	code := a.Comiple()
	return code, nil
}

func (vm *VM) Eval(sexp kl.Obj) (res kl.Obj) {
	defer func() {
		if r := recover(); r != nil {
			vm.Reset()
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
	res = vm.Run(code)
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

func BootstrapShen() {
	BootstrapMin()
	prototype.mustLoadBytecode(kl.MakeString("toplevel.bc"))
	prototype.mustLoadBytecode(kl.MakeString("core.bc"))
	prototype.mustLoadBytecode(kl.MakeString("sys.bc"))
	prototype.mustLoadBytecode(kl.MakeString("sequent.bc"))
	prototype.mustLoadBytecode(kl.MakeString("yacc.bc"))
	prototype.mustLoadBytecode(kl.MakeString("reader.bc"))
	prototype.mustLoadBytecode(kl.MakeString("reader.bc"))
	prototype.mustLoadBytecode(kl.MakeString("prolog.bc"))
	prototype.mustLoadBytecode(kl.MakeString("track.bc"))
	prototype.mustLoadBytecode(kl.MakeString("load.bc"))
	prototype.mustLoadBytecode(kl.MakeString("writer.bc"))
	prototype.mustLoadBytecode(kl.MakeString("macros.bc"))
	prototype.mustLoadBytecode(kl.MakeString("declarations.bc"))
	prototype.mustLoadBytecode(kl.MakeString("t-star.bc"))
	prototype.mustLoadBytecode(kl.MakeString("types.bc"))
}

var Debug bool

func (vm *VM) mustLoadBytecode(args ...kl.Obj) {
	res := vm.loadBytecode(args...)
	if kl.IsError(res) {
		panic(kl.ObjString(res))
	}
}

func (vm *VM) loadBytecode(args ...kl.Obj) kl.Obj {
	fileName := kl.GetString(args[0])
	var f io.ReadCloser
	var err error
	if Debug {
		filePath := path.Join(kl.PackagePath(), "bytecode", fileName)
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
		code := a.Comiple()

		tmp := auxVM.Get()
		tmp.nativeFunc = vm.nativeFunc
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

func (vm *VM) loadFile(args ...kl.Obj) kl.Obj {
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
		tmp.nativeFunc = vm.nativeFunc
		res := tmp.Eval(exp)
		auxVM.Put(tmp)

		if kl.IsError(res) {
			return res
		}
	}
	return args[0]
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
