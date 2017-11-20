package kl

import (
	"bytes"
	"fmt"
	"unsafe"
)

// The type of all Scheme objects.
type Obj *scmHead

type scmHead int

const (
	Number    scmHead = 0
	Pair              = 1
	Vector            = 2
	Null              = 3
	String            = 4
	Symbol            = 5
	Boolean           = 6
	Procedure         = 14
	Stream            = 17
	Primitive         = 21
	Error             = 22
	Raw               = 42
)

type ScmRaw struct {
	scmHead
}

type ScmNumber struct {
	scmHead
	val float64
}

type ScmChar struct {
	scmHead
	codepoint rune
}

type ScmSymbol struct {
	scmHead
	sym string
}

type ScmPair struct {
	scmHead
	car Obj
	cdr Obj
}

type ScmVector struct {
	scmHead
	vector []Obj
}

type ScmString struct {
	scmHead
	str string
}

type ScmStream struct {
	scmHead
	raw interface{}
}

type ScmBoolean struct {
	scmHead
	bool
}

type ScmProcedure struct {
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
}

type ScmError struct {
	scmHead
	err string
}

func Make_raw() ScmRaw {
	return ScmRaw{scmHead: Raw}
}

func (raw *ScmRaw) Object() Obj {
	return &raw.scmHead
}

func Make_error(err string) Obj {
	tmp := ScmError{Error, err}
	return &tmp.scmHead
}

func mustError(o Obj) *ScmError {
	if *o != Error {
		panic("mustError")
	}
	return (*ScmError)(unsafe.Pointer(o))
}

func isPrimitive(o Obj) (bool, *ScmPrimitive) {
	if *o != Primitive {
		return false, nil
	}
	return true, (*ScmPrimitive)(unsafe.Pointer(o))
}

func mustPrimitive(o Obj) *ScmPrimitive {
	if *o != Primitive {
		panic("mustPrimitive")
	}
	return (*ScmPrimitive)(unsafe.Pointer(o))
}

func mustVector(o Obj) []Obj {
	if (*o) != Vector {
		panic("mustVector")
	}
	tmp := (*ScmVector)(unsafe.Pointer(o))
	return tmp.vector
}

func mustProcedure(o Obj) *ScmProcedure {
	if (*o) != Procedure {
		panic("mustProcedure")
	}
	return (*ScmProcedure)(unsafe.Pointer(o))
}

func mustString(o Obj) string {
	if (*o) != String {
		panic("mustString")
	}
	return (*ScmString)(unsafe.Pointer(o)).str
}

func mustInteger(o Obj) int {
	if (*o) != Number {
		panic("mustNumber")
	}
	f := (*ScmNumber)(unsafe.Pointer(o)).val
	return int(f)
}

func mustNumber(o Obj) *ScmNumber {
	if (*o) != Number {
		panic("mustNumber")
	}
	return (*ScmNumber)(unsafe.Pointer(o))
}

func mustSymbol(o Obj) *ScmSymbol {
	if (*o) != Symbol {
		panic("mustSymbol")
	}
	return (*ScmSymbol)(unsafe.Pointer(o))
}

func isSymbol(o Obj) (bool, *ScmSymbol) {
	if *o == Symbol {
		return true, (*ScmSymbol)(unsafe.Pointer(o))
	}
	return false, nil
}

func mustStream(o Obj) *ScmStream {
	if (*o) != Stream {
		panic("mustStream")
	}
	return (*ScmStream)(unsafe.Pointer(o))
}

func mustPair(o Obj) *ScmPair {
	if (*o) != Pair {
		panic("mustPair")
	}
	return (*ScmPair)(unsafe.Pointer(o))
}

func isPair(o Obj) (bool, *ScmPair) {
	if (*o) == Pair {
		return true, (*ScmPair)(unsafe.Pointer(o))
	}
	return false, nil
}

var True, False, Nil Obj

func init() {
	tmp1 := ScmBoolean{Boolean, false}
	False = Obj(&tmp1.scmHead)

	tmp2 := &ScmBoolean{Boolean, true}
	True = Obj(&tmp2.scmHead)

	tmp3 := &ScmPair{Null, nil, nil}
	Nil = Obj(&tmp3.scmHead)
}

func Make_integer(v int) Obj {
	tmp := ScmNumber{Number, float64(v)}
	return &tmp.scmHead
}

func Make_number(f float64) Obj {
	tmp := ScmNumber{Number, f}
	return &tmp.scmHead
}

func Make_boolean(x bool) Obj {
	if x {
		return True
	}
	return False
}

func Make_stream(raw interface{}) Obj {
	tmp := ScmStream{
		Stream,
		raw,
	}
	return &tmp.scmHead
}

func SymbolString(o Obj) string {
	return mustSymbol(o).sym
}

func cons(x, y Obj) Obj {
	tmp := ScmPair{
		scmHead: Pair,
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

func Make_vector(n int) Obj {
	tmp := ScmVector{
		Vector,
		make([]Obj, n),
	}
	return &tmp.scmHead
}

func Make_string(s string) Obj {
	tmp := ScmString{String, s}
	return &tmp.scmHead
}

func Make_symbol(s string) Obj {
	tmp := ScmSymbol{
		Symbol,
		s,
	}
	return &tmp.scmHead
}

func Make_procedure(arg Obj, body Obj, env *Environment) Obj {
	tmp := ScmProcedure{
		scmHead: Procedure,
		body:    body,
		env:     env,
	}
	if *arg == Symbol {
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

func (o *scmHead) GoString() string {
	switch *o {
	case Number:
		return fmt.Sprintf("Number(%f)", mustNumber(o).val)
	case Pair:
		if o == Nil {
			return "Nil"
		}
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "Pair(")
		fmt.Fprintf(&buf, "%#v", (*scmHead)(car(o)))
		fmt.Fprintf(&buf, "%#v", (*scmHead)(cdr(o)))
		fmt.Fprintf(&buf, ")")
		return buf.String()
	case Vector:
		return fmt.Sprintf("Vector")
	case Null:
		return fmt.Sprintf("Null")
	case String:
		return fmt.Sprintf("String(%s)", mustString(o))
	case Symbol:
		return fmt.Sprintf("Symbol(%s)", mustSymbol(o).sym)
	case Boolean:
		if o == True {
			return fmt.Sprintf("Boolean(True)")
		} else if o == False {
			return fmt.Sprintf("Boolean(False)")
		} else {
			return fmt.Sprintf("Boolean(something wrong)")
		}
	case Error:
		return fmt.Sprintf("Error(%s)", mustError(o).err)
	case Procedure:
		return fmt.Sprintf("Procedure")
	case Stream:
		return fmt.Sprintf("Stream")
	case Primitive:
		prim := mustPrimitive(o)
		return fmt.Sprintf("Primitive(%s)", prim.Name)
	}
	return fmt.Sprintf("unknown type %d", *o)
}
