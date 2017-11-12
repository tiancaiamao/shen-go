package kl

import (
	"io"
	// "fmt"
	"strings"
	"testing"
)

func TestSexpReader(t *testing.T) {
	IF := Make_symbol("if")
	tests := []struct {
		input  string
		expect Obj
	}{
		{"1234", Make_integer(1234)},
		{`"string"`, Make_string("string")},
		{"symbol", Make_symbol("symbol")},
		{"()", Nil},
		{"(1 2)", cons(Make_integer(1), cons(Make_integer(2), Nil))},
		{"true", True},
		{"false", False},
		{"(if true (if false 1 2) 3)", cons(IF, cons(True,
			cons(
				cons(IF, cons(False, cons(Make_integer(1), cons(Make_integer(2), Nil)))),
				cons(Make_integer(3), Nil))))},
		{`"abc
de"`, Make_string("abc\nde")},
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
	if equal(o1, cons(Make_symbol("if"), cons(Make_boolean(true), cons(Make_integer(1), cons(Make_boolean(false), Nil))))) == False {
		t.Errorf("if true... %#v", (*scmHead)(o1))
	}

	o2, err := r.Read()
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	if equal(o2, Make_integer(2)) == False {
		t.Error("read another sexp")
	}
}
