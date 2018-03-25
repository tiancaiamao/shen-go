package runtime

import (
	"io"
	// "fmt"
	"strings"
	"testing"
)

func TestSexpReader(t *testing.T) {
	IF := MakeSymbol("if")
	tests := []struct {
		input  string
		expect Obj
	}{
		{"1234", MakeInteger(1234)},
		{`"string"`, MakeString("string")},
		{"symbol", MakeSymbol("symbol")},
		{"()", Nil},
		{"(1 2)", cons(MakeInteger(1), cons(MakeInteger(2), Nil))},
		{"true", True},
		{"false", False},
		{"(if true (if false 1 2) 3)", cons(IF, cons(True,
			cons(
				cons(IF, cons(False, cons(MakeInteger(1), cons(MakeInteger(2), Nil)))),
				cons(MakeInteger(3), Nil))))},
		{`"abc
de"`, MakeString("abc\nde")},
	}
	for _, test := range tests {
		r := NewSexpReader(strings.NewReader(test.input))
		o, err := r.Read()
		if err != nil && err != io.EOF {
			t.Error("read error", err)
		}
		if equal(o, test.expect) == False {
			t.Errorf("%s fail, expect: %#v, but get: %#v\n", test.input, (*scmHead)(test.expect), (*scmHead)(o))
		}
	}

	r := NewSexpReader(strings.NewReader("(if true 1 false) 2"))
	o1, err := r.Read()
	if err != nil && err != io.EOF {
		t.Fatal("read error", err)
	}
	if equal(o1, cons(MakeSymbol("if"), cons(True, cons(MakeInteger(1), cons(False, Nil))))) == False {
		t.Errorf("if true... %#v", (*scmHead)(o1))
	}

	o2, err := r.Read()
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	if equal(o2, MakeInteger(2)) == False {
		t.Error("read another sexp")
	}
}

func TestReaderMacro(t *testing.T) {
	a, _ := NewSexpReader(strings.NewReader("^'(a b c)")).Read()
	b, _ := NewSexpReader(strings.NewReader("(cons a (cons b (cons c ())))")).Read()
	if equal(a, b) == False {
		t.Error("fail:", ObjString(a), ObjString(b))
	}
}
