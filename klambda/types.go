package klambda

import (
	"bytes"
	"fmt"
	"io"
	"syscall"
	"time"
	"unsafe"
)

// All kinds of Scheme object.
type Obj *scmHead

type scmHead int

const (
	scmHeadNumber  scmHead = 0
	scmHeadPair            = 1
	scmHeadVector          = 2
	scmHeadNull            = 3
	scmHeadString          = 4
	scmHeadSymbol          = 5
	scmHeadBoolean         = 6
	scmHeadStream          = 17
	scmHeadError           = 22
	scmHeadNative          = 23
	scmHeadClosure         = 24
	scmHeadRaw             = 42
)

type scmClosure struct {
	scmHead
	code   inst
	params int

	parent   *scmClosure
	freeVars map[int]Obj
	nlocal   int
}

type scmNumber struct {
	scmHead
	val float64
}

type scmSymbol struct {
	scmHead
	// The string of this symbol.
	str      string
	value    Obj
	function Obj
	cora     Obj
}

type scmPair struct {
	scmHead
	car Obj
	cdr Obj
}

type scmVector struct {
	scmHead
	vector []Obj
}

type scmString struct {
	scmHead
	str string
}

type scmStream struct {
	scmHead
	raw interface{}
}

type scmBoolean struct {
	scmHead
	bool
}

type scmNative struct {
	scmHead
	name     string
	fn       func(*ControlFlow)
	require  int
	captured []Obj
}

func objType(o Obj) scmHead {
	if isFixnum(uintptr(unsafe.Pointer(o))) {
		return scmHeadNumber
	}
	return *o
}

func MakeNative(fn func(*ControlFlow), require int, captured ...Obj) Obj {
	tmp := scmNative{
		scmHead:  scmHeadNative,
		fn:       fn,
		require:  require,
		captured: captured,
	}
	return &tmp.scmHead
}

func MustNative(o Obj) *scmNative {
	if objType(o) != scmHeadNative {
		panic("mustNative")
	}
	return (*scmNative)(unsafe.Pointer(o))
}

func IsNative(o Obj) bool {
	return objType(o) == scmHeadNative
}

type scmError struct {
	scmHead
	err string
}

// MakeRaw makes a struct into a raw object.
// Usage:
// type T struct {
//    scmHead int
//    ... // xxx
// }
// tmp := &T{}
// raw := MakeRaw(&tmp.scmHead)
func MakeRaw(scmHead *int) Obj {
	*scmHead = scmHeadRaw
	return Obj(unsafe.Pointer(scmHead))
}

func MakeError(err string) Obj {
	tmp := scmError{scmHeadError, err}
	return &tmp.scmHead
}

func MustError(o Obj) *scmError {
	if objType(o) != scmHeadError {
		panic(MakeError("mustError"))
	}
	return (*scmError)(unsafe.Pointer(o))
}

func IsError(o Obj) bool {
	return objType(o) == scmHeadError
}

func IsNumber(o Obj) bool {
	return objType(o) == scmHeadNumber
}

func IsSymbol(o Obj) bool {
	return objType(o) == scmHeadSymbol
}

func IsVector(o Obj) bool {
	return objType(o) == scmHeadVector
}

func mustVector(o Obj) []Obj {
	if (*o) != scmHeadVector {
		panic(MakeError("mustVector"))
	}
	tmp := (*scmVector)(unsafe.Pointer(o))
	return tmp.vector
}

func mustString(o Obj) string {
	if objType(o) != scmHeadString {
		panic(MakeError("mustString"))
	}
	return (*scmString)(unsafe.Pointer(o)).str
}

func fixnum(o Obj) int {
	return int(uintptr(unsafe.Pointer(o)) - fixnumBaseAddr)
}

func MustInteger(o Obj) int {
	return mustInteger(o)
}

func mustInteger(o Obj) int {
	if isFixnum(uintptr(unsafe.Pointer(o))) {
		return fixnum(o)
	}
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
	}

	f := (*scmNumber)(unsafe.Pointer(o)).val
	return int(f)
}

func GetInteger(o Obj) int {
	if isFixnum(uintptr(unsafe.Pointer(o))) {
		return fixnum(o)
	}
	f := (*scmNumber)(unsafe.Pointer(o)).val
	return int(f)
}

func mustNumber(o Obj) float64 {
	if isFixnum(uintptr(unsafe.Pointer(o))) {
		return float64(fixnum(o))
	}
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
	}
	x := (*scmNumber)(unsafe.Pointer(o))
	return x.val
}

func mustSymbol(o Obj) *scmSymbol {
	if objType(o) != scmHeadSymbol {
		panic(MakeError("mustSymbol"))
	}
	return (*scmSymbol)(unsafe.Pointer(o))
}

func isSymbol(o Obj) (bool, *scmSymbol) {
	if objType(o) == scmHeadSymbol {
		return true, (*scmSymbol)(unsafe.Pointer(o))
	}
	return false, nil
}

func mustStream(o Obj) *scmStream {
	if objType(o) != scmHeadStream {
		panic(MakeError("mustStream"))
	}
	return (*scmStream)(unsafe.Pointer(o))
}

func mustClosure(o Obj) *scmClosure {
	if objType(o) != scmHeadClosure {
		panic(MakeError("mustClosure"))
	}
	return (*scmClosure)(unsafe.Pointer(o))
}

func mustPair(o Obj) *scmPair {
	/*
		if objType(o) != scmHeadPair {
			fmt.Println(ObjString(o))
			panic(MakeError("mustPair"))
		}
	*/
	return (*scmPair)(unsafe.Pointer(o))
}

func isPair(o Obj) (bool, *scmPair) {
	if objType(o) == scmHeadPair {
		return true, (*scmPair)(unsafe.Pointer(o))
	}
	return false, nil
}

var True, False, Nil, undefined Obj
var uptime time.Time
var symMacroExpand Obj
var symType Obj
var symDefun Obj
var symFreeze Obj
var symLet Obj
var symCond Obj
var symAnd Obj
var symOr Obj
var symIf Obj
var symLambda Obj
var symTrapError Obj
var symDo Obj
var symQuote Obj

var fixnumBaseAddr uintptr
var fixnumEndAddr uintptr

const fixnumCount = 1 << 30

type trieNode struct {
	children [256]*trieNode
	value    scmSymbol
}

var symbolRoot trieNode

func trieFindOrInsert(str string) *trieNode {
	p := &symbolRoot
	for i := 0; i < len(str); i++ {
		v := str[i]
		if p.children[v] == nil {
			p.children[v] = &trieNode{
				value: scmSymbol{
					scmHead: scmHeadSymbol,
					str:     str[:i+1],
				},
			}
		}
		p = p.children[v]
	}
	return p
}

var klMacro map[Obj]func(cc *Compiler, obj Obj, env *Env, tail bool) inst

func init() {
	// Create a mmap area for the memory address of fixnums.
	data, err := syscall.Mmap(-1, /*fd int*/
		0,                 /*offset int64*/
		fixnumCount,       /*length int*/
		syscall.PROT_READ, /*prot int*/
		syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE /*flags int*/)
	if err != nil {
		panic(err)
	}
	fixnumBaseAddr = uintptr(unsafe.Pointer(&data[0]))
	fixnumEndAddr = fixnumBaseAddr + fixnumCount

	uptime = time.Now()
	tmp1 := &scmBoolean{scmHeadBoolean, false}
	False = Obj(&tmp1.scmHead)

	tmp2 := &scmBoolean{scmHeadBoolean, true}
	True = Obj(&tmp2.scmHead)

	tmp3 := &scmPair{scmHeadNull, nil, nil}
	Nil = Obj(&tmp3.scmHead)

	var tmp4 int
	undefined = MakeRaw(&tmp4)

	symQuote = MakeSymbol("quote")
	symDefun = MakeSymbol("defun")
	symLambda = MakeSymbol("lambda")
	symLet = MakeSymbol("let")
	symIf = MakeSymbol("if")
	symDo = MakeSymbol("do")
	symMacroExpand = MakeSymbol("macroexpand")

	symType = MakeSymbol("type")
	symDefun = MakeSymbol("defun")
	symFreeze = MakeSymbol("freeze")
	symLet = MakeSymbol("let")
	symCond = MakeSymbol("cond")
	symAnd = MakeSymbol("and")
	symOr = MakeSymbol("or")
	symIf = MakeSymbol("if")
	symLambda = MakeSymbol("lambda")
	symTrapError = MakeSymbol("trap-error")
	symDo = MakeSymbol("do")

	klMacro = map[Obj]func(cc *Compiler, obj Obj, env *Env, tail bool) inst{
		symDefun:     compileDefunMacro,
		symCond:      compileCondMacro,
		symAnd:       compileAndMacro,
		symOr:        compileOrMacro,
		symTrapError: compileTrapErrorMacro,
		symFreeze:    compileFreezeMacro,
		// symLet: compileLetMacro,
		symType: compileTypeMacro,
	}
}

func MakeInteger(v int) Obj {
	if v >= 0 && v < fixnumCount {
		return Obj(unsafe.Pointer(fixnumBaseAddr + uintptr(v)))
	}
	return makeInteger(v)
}

func IsFixnum(v uintptr) bool {
	return isFixnum(v)
}

func isFixnum(v uintptr) bool {
	if v >= fixnumBaseAddr && v < fixnumEndAddr {
		return true
	}
	return false
}

func makeInteger(v int) Obj {
	tmp := scmNumber{scmHeadNumber, float64(v)}
	return &tmp.scmHead
}

func MakeNumber(f float64) Obj {
	if isPreciseInteger(f) {
		return MakeInteger(int(f))
	}

	tmp := scmNumber{scmHeadNumber, f}
	return &tmp.scmHead
}

func MakeStream(raw interface{}) Obj {
	tmp := scmStream{
		scmHeadStream,
		raw,
	}
	return &tmp.scmHead
}

func IsString(o Obj) bool {
	return objType(o) == scmHeadString
}

func GetString(o Obj) string {
	return mustString(o)
}

func GetSymbol(o Obj) string {
	return mustSymbol(o).str
}

func cons(x, y Obj) Obj {
	tmp := scmPair{
		scmHead: scmHeadPair,
		car:     x,
		cdr:     y,
	}
	return &tmp.scmHead
}

func car(x Obj) Obj {
	return mustPair(x).car
}

func cdr(x Obj) Obj {
	return mustPair(x).cdr
}

func MakeVector(n int) Obj {
	tmp := scmVector{
		scmHeadVector,
		make([]Obj, n),
	}
	return &tmp.scmHead
}

func MakeString(s string) Obj {
	tmp := scmString{scmHeadString, s}
	return &tmp.scmHead
}

func MakeSymbol(s string) Obj {
	p := trieFindOrInsert(s)
	return &p.value.scmHead
}

func MakeClosure(code inst, ebp, esp int, nargs int, parent *scmClosure, freeVars map[int]Obj, nlocal int) Obj {
	tmp := scmClosure{
		scmHead:  scmHeadClosure,
		code:     code,
		params:   nargs,
		parent:   parent,
		freeVars: freeVars,
		nlocal:   nlocal,
	}
	return &tmp.scmHead
}

func makeClosure(params, body, env Obj) Obj {
	return cons(symLambda, cons(params, cons(body, env)))
}

func isClosure(o Obj) bool {
	return objType(o) == scmHeadPair && car(o) == symLambda
}

func ObjString(o Obj) string {
	return (*scmHead)(o).GoString()
}

func (o *scmPair) fmt(buf io.Writer, start bool) {
	if start {
		fmt.Fprintf(buf, "(%s", ObjString(o.car))
	} else {
		fmt.Fprintf(buf, " %s", ObjString(o.car))
	}
	switch *o.cdr {
	case scmHeadNull:
		fmt.Fprintf(buf, ")")
	case scmHeadPair:
		mustPair(o.cdr).fmt(buf, false)
	default:
		fmt.Fprintf(buf, " . %s)", ObjString(o.cdr))
	}
}

func (o *scmHead) GoString() string {
	if isFixnum(uintptr(unsafe.Pointer(o))) {
		return fmt.Sprintf("%d", fixnum(o))
	}
	switch *o {
	case scmHeadNumber:
		f := mustNumber(o)
		if !isPreciseInteger(f) {
			return fmt.Sprintf("%f", f)
		}
		return fmt.Sprintf("%d", int(f))
	case scmHeadPair:
		var buf bytes.Buffer
		mustPair(o).fmt(&buf, true)
		return buf.String()
	case scmHeadVector:
		return fmt.Sprintf("#vector")
	case scmHeadNull:
		return fmt.Sprintf("()")
	case scmHeadString:
		return fmt.Sprintf(`"%s"`, mustString(o))
	case scmHeadSymbol:
		return fmt.Sprintf("%s", GetSymbol(o))
	case scmHeadBoolean:
		if o == True {
			return fmt.Sprintf("true")
		} else if o == False {
			return fmt.Sprintf("false")
		} else {
			return fmt.Sprintf("Boolean(something wrong)")
		}
	case scmHeadError:
		return fmt.Sprintf("Error(%s)", MustError(o).err)
	case scmHeadStream:
		return "#stream"
	case scmHeadRaw:
		return "#raw"
	case scmHeadNative:
		prim := MustNative(o)
		if len(prim.name) > 0 {
			return fmt.Sprintf("#primitive(%s)", prim.name)
		}
		return "#native"
	case scmHeadClosure:
		return "#closure"
	}
	return fmt.Sprintf("unknown type %d", *o)
}
