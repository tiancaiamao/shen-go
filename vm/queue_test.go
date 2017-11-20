package vm

import (
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestQueue(t *testing.T) {
	var q queue
	q.init(2)
	q.push(kl.Make_number(1))
	q.push(kl.Make_number(2))
	q.push(kl.Make_number(3))
	q.push(kl.Make_number(4))

	mustEqualNumber(q.pop(), 1, t)
	mustEqualNumber(q.pop(), 2, t)

	q.push(kl.Make_number(5))
	q.push(kl.Make_number(6))

	if q.count() != 4 {
		t.Error("wrong count")
	}

	mustEqualNumber(q.pop(), 3, t)
	mustEqualNumber(q.pop(), 4, t)
}

func mustEqualNumber(o kl.Obj, x float64, t *testing.T) {
	if kl.PrimEqual(o, kl.Make_number(x)) != kl.True {
		t.Fatal("mustEqualNumber:", kl.ObjString(o), x)
	}
}
