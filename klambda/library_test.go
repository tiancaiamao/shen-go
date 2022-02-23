package klambda

import (
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
		{MakeVector(1), MakeVector(2), False},
		{MakeVector(0), MakeVector(0), True},
		{MakeInteger(5), MakeInteger(fixnumCount + 3), False},
		{MakeInteger(0), MakeInteger(5), False},
		{MakeInteger(5), MakeInteger(5), True},
	}

	for _, test := range tests {
		if equal(test.x, test.y) != test.expect {
			t.Error(test.x, test.y)
		}
	}
}

func TestVectorGet(t *testing.T) {
	var kl ControlFlow
	vec := MakeVector(1)
	kl.Return(PrimVectorSet(vec, MakeInteger(0), MakeNumber(42)))
	func() {
		defer func() {
			err := recover()
			val := err.(Obj)
			if *val != scmHeadError {
				t.Error("should be error out of range")
			}
		}()
		kl.Return(PrimVectorGet(vec, MakeInteger(1)))
	}()
	if equal(PrimVectorGet(vec, MakeInteger(0)), MakeNumber(42)) != True {
		t.Error("vector set or get wrong")
	}
}

func TestIsPreciseInteger(t *testing.T) {
	testPreciseInteger(t, 0.0, true)
	testPreciseInteger(t, 1.0, true)
	testPreciseInteger(t, 3, true)
	testPreciseInteger(t, 7.0, true)
	testPreciseInteger(t, -1.0, true)
	testPreciseInteger(t, 2251799813685248.0, true)

	testPreciseInteger(t, 2.14, false)
	testPreciseInteger(t, 0.14, false)
	testPreciseInteger(t, -0.14, false)
	testPreciseInteger(t, 1.1, false)
	testPreciseInteger(t, 1024.5, false)
	testPreciseInteger(t, 2251799813685248.5, false)
}

func testPreciseInteger(t *testing.T, f float64, expect bool) {
	if isPreciseInteger(f) != expect {
		t.Error("testPreciseInteger wrong:", f, expect)
	}
}
