package kl

import (
	"bytes"
	"fmt"
	"io"
	"time"
	"unsafe"
)

// All kinds of Scheme object.
type Obj *scmHead

type scmHead int

const (
	scmHeadNumber    scmHead = 0
	scmHeadPair              = 1
	scmHeadVector            = 2
	scmHeadNull              = 3
	scmHeadString            = 4
	scmHeadSymbol            = 5
	scmHeadBoolean           = 6
	scmHeadProcedure         = 14
	scmHeadStream            = 17
	scmHeadPrimitive         = 21
	scmHeadError             = 22
	scmHeadNative            = 23
	scmHeadRaw               = 42
)

type scmNumber struct {
	scmHead
	val float64
}

type scmSymbol struct {
	scmHead
	offset int
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

type scmProcedure struct {
	scmHead
	name  string
	arg   []Obj
	arity int
	body  Obj
	env   *Environment
}

type ScmPrimitive struct {
	scmHead
	Name     string
	Required int
	Function func(...Obj) Obj
	CodeGen  string
}

type scmNative struct {
	scmHead
	fn       func(*Evaluator, *ControlFlow, ...Obj)
	require  int
	captured []Obj
}

func MakeNative(fn func(*Evaluator, *ControlFlow, ...Obj), require int, captured ...Obj) Obj {
	tmp := scmNative{
		scmHead:  scmHeadNative,
		fn:       fn,
		require:  require,
		captured: captured,
	}
	return &tmp.scmHead
}

func MustNative(o Obj) *scmNative {
	if *o != scmHeadNative {
		panic("mustNative")
	}
	return (*scmNative)(unsafe.Pointer(o))
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

func mustError(o Obj) *scmError {
	if *o != scmHeadError {
		panic(MakeError("mustError"))
	}
	return (*scmError)(unsafe.Pointer(o))
}

func IsError(o Obj) bool {
	return *o == scmHeadError
}

func IsNumber(o Obj) bool {
	return *o == scmHeadNumber
}

func IsSymbol(o Obj) bool {
	return *o == scmHeadSymbol
}

func MakePrimitive(name string, arity int, f func(...Obj) Obj) *ScmPrimitive {
	return &ScmPrimitive{
		scmHead:  scmHeadPrimitive,
		Name:     name,
		Required: arity,
		Function: f,
	}
}

func isPrimitive(o Obj) (bool, *ScmPrimitive) {
	if *o != scmHeadPrimitive {
		return false, nil
	}
	return true, (*ScmPrimitive)(unsafe.Pointer(o))
}

func mustPrimitive(o Obj) *ScmPrimitive {
	if *o != scmHeadPrimitive {
		panic(MakeError("mustPrimitive"))
	}
	return (*ScmPrimitive)(unsafe.Pointer(o))
}

func mustVector(o Obj) []Obj {
	if (*o) != scmHeadVector {
		panic(MakeError("mustVector"))
	}
	tmp := (*scmVector)(unsafe.Pointer(o))
	return tmp.vector
}

func mustProcedure(o Obj) *scmProcedure {
	if (*o) != scmHeadProcedure {
		panic(MakeError("mustProcedure"))
	}
	return (*scmProcedure)(unsafe.Pointer(o))
}

func mustString(o Obj) string {
	if (*o) != scmHeadString {
		panic(MakeError("mustString"))
	}
	return (*scmString)(unsafe.Pointer(o)).str
}

func fixnum(o Obj) int {
	return int(uintptr(unsafe.Pointer(o)) - fixnumBaseAddr)
}

func mustInteger(o Obj) int {
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
	}
	if isFixnum(o) {
		return fixnum(o)
	}

	f := (*scmNumber)(unsafe.Pointer(o)).val
	return int(f)
}

func mustNumber(o Obj) float64 {
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
	}
	if isFixnum(o) {
		return float64(fixnum(o))
	}
	x := (*scmNumber)(unsafe.Pointer(o))
	return x.val
}

func mustSymbol(o Obj) *scmSymbol {
	if (*o) != scmHeadSymbol {
		panic(MakeError("mustSymbol"))
	}
	return (*scmSymbol)(unsafe.Pointer(o))
}

func isSymbol(o Obj) (bool, *scmSymbol) {
	if *o == scmHeadSymbol {
		return true, (*scmSymbol)(unsafe.Pointer(o))
	}
	return false, nil
}

func mustStream(o Obj) *scmStream {
	if (*o) != scmHeadStream {
		panic(MakeError("mustStream"))
	}
	return (*scmStream)(unsafe.Pointer(o))
}

func mustPair(o Obj) *scmPair {
	if (*o) != scmHeadPair {
		fmt.Println(ObjString(o))
		panic(MakeError("mustPair"))
	}
	return (*scmPair)(unsafe.Pointer(o))
}

func isPair(o Obj) (bool, *scmPair) {
	if (*o) == scmHeadPair {
		return true, (*scmPair)(unsafe.Pointer(o))
	}
	return false, nil
}

var True, False, Nil, undefined Obj
var uptime time.Time

var addrForFixnum [1 << 20]byte

const fixnumCount = 8 << 20

var fixnumBaseAddr = uintptr(unsafe.Pointer(&addrForFixnum[0]))
var fixnumEndAddr = fixnumBaseAddr + fixnumCount

type symbolArrayObj struct {
	str      string // The string of this symbol
	value    Obj    // The value bind to this symbol
	function Obj    // The function bind to this symbol
}

var symbolArray []symbolArrayObj

type trieNode struct {
	children [256]*trieNode
	value    int
}

var trieRoot *trieNode

func trieFindOrInsert(str string) *trieNode {
	p := trieRoot
	for i := 0; i < len(str); i++ {
		v := str[i]
		if p.children[v] == nil {
			p.children[v] = &trieNode{value: -1}
		}
		p = p.children[v]
	}
	return p
}

func init() {
	uptime = time.Now()
	tmp1 := &scmBoolean{scmHeadBoolean, false}
	False = Obj(&tmp1.scmHead)

	tmp2 := &scmBoolean{scmHeadBoolean, true}
	True = Obj(&tmp2.scmHead)

	tmp3 := &scmPair{scmHeadNull, nil, nil}
	Nil = Obj(&tmp3.scmHead)

	var tmp4 int
	undefined = MakeRaw(&tmp4)

	symbolArray = make([]symbolArrayObj, 0, 4096)
	trieRoot = &trieNode{}
}

func MakeInteger(v int) Obj {
	if v >= 0 && v < fixnumCount {
		return Obj(unsafe.Pointer(fixnumBaseAddr + uintptr(v)))
	}
	return makeInteger(v)
}

func isFixnum(o Obj) bool {
	v := uintptr(unsafe.Pointer(o))
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

func GetString(o Obj) string {
	return mustString(o)
}

func GetSymbol(o Obj) string {
	return symbolArray[mustSymbol(o).offset].str
}

func BindSymbolFunc(sym Obj, f Obj) {
	o := &symbolArray[mustSymbol(sym).offset]
	o.function = f
}

func GetSymbolFunc(sym Obj) Obj {
	o := &symbolArray[mustSymbol(sym).offset]
	return o.function
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
	var idx int
	p := trieFindOrInsert(s)
	if p.value < 0 {
		idx = len(symbolArray)
		symbolArray = append(symbolArray, symbolArrayObj{str: s})
		p.value = idx
	} else {
		idx = p.value
	}

	tmp := scmSymbol{
		scmHeadSymbol,
		idx,
	}
	return &tmp.scmHead
}

func makeProcedure(arg Obj, body Obj, env *Environment) Obj {
	tmp := scmProcedure{
		scmHead: scmHeadProcedure,
		body:    body,
		env:     env,
	}
	if *arg == scmHeadSymbol {
		tmp.arg = []Obj{arg}
		tmp.arity = 1
	} else {
		tmp.arg = ListToSlice(arg)
		tmp.arity = len(tmp.arg)
	}
	return &tmp.scmHead
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
		return fmt.Sprintf("Error(%s)", mustError(o).err)
	case scmHeadProcedure:
		return "#procedure"
	case scmHeadStream:
		return "#stream"
	case scmHeadPrimitive:
		prim := mustPrimitive(o)
		return fmt.Sprintf("#primitive(%s)", prim.Name)
	case scmHeadRaw:
		return "#raw"
	case scmHeadNative:
		return "#native"
	}
	return fmt.Sprintf("unknown type %d", *o)
}
