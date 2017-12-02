package kl

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	if reverse(Nil) != Nil {
		t.FailNow()
	}

	l := cons(MakeInteger(1), cons(MakeInteger(2), cons(MakeInteger(3), Nil)))
	r := reverse(l)
	if mustInteger(car(r)) != 3 {
		t.FailNow()
	}
	if mustInteger(cadr(r)) != 2 {
		t.FailNow()
	}
	if mustInteger(caddr(r)) != 1 {
		t.FailNow()
	}
	if cdddr(r) != Nil {
		t.FailNow()
	}

	// (1 (1 2 3))
	l1 := cons(MakeInteger(1), cons(l, Nil))
	// ((1 2 3) 1)
	l2 := cons(l, cons(MakeInteger(1), Nil))

	if equal(reverse(l1), l2) != True {
		t.Error("fuck1")
	}
	if equal(reverse(l2), l1) != True {
		t.Error("fuck2")
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		x      Obj
		y      Obj
		expect Obj
	}{
		{True, True, True},
		{True, False, False},
		{False, True, False},
		{True, MakeNumber(10), False},
		{Nil, Nil, True},
		{Nil, False, False},
		{MakeString("asd"), MakeString("abc"), False},
	}

	for _, test := range tests {
		if equal(test.x, test.y) != test.expect {
			t.Error(test.x, test.y)
		}
	}
}

func TestPrimitive(t *testing.T) {
	fmt.Println("(init-primitive-table [")
	for i, prim := range allPrimitives {
		fmt.Println(prim.Name, prim.Required, `"xxx"`, i)
	}
	fmt.Println("])")
}
