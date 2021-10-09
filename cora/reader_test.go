package cora

import (
	"io"
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
		r := NewSexpReader(strings.NewReader(test.input), false)
		o, err := r.Read()
		if err != nil && err != io.EOF {
			t.Error("read error", err)
		}
		if equal(o, test.expect) == False {
			t.Errorf("%s fail, expect: %#v, but get: %#v\n", test.input, (*scmHead)(test.expect), (*scmHead)(o))
		}
	}

	r := NewSexpReader(strings.NewReader("(if true 1 false) 2"), false)
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
	a, _ := NewSexpReader(strings.NewReader("'(a b c)"), true).Read()
	b, _ := NewSexpReader(strings.NewReader("(quote (a b c))"), false).Read()
	if equal(a, b) == False {
		t.Error("fail:", ObjString(a), ObjString(b))
	}

	a, err1 := NewSexpReader(strings.NewReader("[a b c]"), true).Read()
	b, err2 := NewSexpReader(strings.NewReader("(list a b c)"), true).Read()
	if err1 != nil || err2 != nil {
		t.Error("fail:", err1, err2)
	}
	if equal(a, b) == False {
		t.Error("fail:", ObjString(a), ObjString(b))
	}

	a, err1 = NewSexpReader(strings.NewReader("[a b . c]"), true).Read()
	b, err2 = NewSexpReader(strings.NewReader("(list-rest a b c)"), true).Read()
	if err1 != nil || err2 != nil {
		t.Error("fail:", err1, err2)
	}
	if equal(a, b) == False {
		t.Error("fail:", ObjString(a), ObjString(b))
	}
}
