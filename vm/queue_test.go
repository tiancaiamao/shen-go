package vm

import (
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestQueue(t *testing.T) {
	var q queue
	q.init(2)
	q.push(kl.MakeNumber(1))
	q.push(kl.MakeNumber(2))
	q.push(kl.MakeNumber(3))
	q.push(kl.MakeNumber(4))

	mustEqualNumber(q.pop(), 1, t)
	mustEqualNumber(q.pop(), 2, t)

	q.push(kl.MakeNumber(5))
	q.push(kl.MakeNumber(6))

	if q.count() != 4 {
		t.Error("wrong count")
	}

	mustEqualNumber(q.pop(), 3, t)
	mustEqualNumber(q.pop(), 4, t)
}

func mustEqualNumber(o kl.Obj, x float64, t *testing.T) {
	if kl.PrimEqual(o, kl.MakeNumber(x)) != kl.True {
		t.Fatal("mustEqualNumber:", kl.ObjString(o), x)
	}
}
