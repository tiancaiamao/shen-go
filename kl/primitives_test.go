package kl

import (
	"bytes"
	"testing"
)

func TestReadByte(t *testing.T) {
	buf := bytes.NewBufferString("a")
	stream := MakeStream(buf)
	b := primReadByte(stream)
	if mustInteger(b) != 97 {
		t.Error("should be 97")
	}

	b = primReadByte(stream)
	if mustInteger(b) != -1 {
		t.Error("read EOF should return -1")
	}
}

func TestIntern(t *testing.T) {
	if primIntern(MakeString("true")) != True {
		t.Error("intern(true) should be boolean")
	}
	if primIntern(MakeString("false")) != False {
		t.Error("intern(false) should be boolean")
	}
	if equal(primIntern(MakeString("asdf")), MakeSymbol("asdf")) != True {
		t.FailNow()
	}
}

func TestHash(t *testing.T) {
	objs := []Obj{
		MakeInteger(3),
		MakeNumber(3.1),
		MakeSymbol("sdsd"),
		MakeString("sdsd"),
		MakeError("sdsd"),
		Obj(&allPrimitives[3].scmHead),
		True,
		False,
		Nil,
	}

	m := make(map[int]Obj)
	for _, o := range objs {
		hash := objHash(o)
		t.Log(hash, ObjString(o))
		if m[hash] != nil {
			t.FailNow()
		}
		m[hash] = o
	}
}
