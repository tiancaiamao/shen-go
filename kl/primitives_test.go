package kl

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReadByte(t *testing.T) {
	buf := bytes.NewBufferString("a")
	stream := MakeStream(buf)
	b := PrimReadByte(stream)
	if mustInteger(b) != 97 {
		t.Error("should be 97")
	}

	b = PrimReadByte(stream)
	if mustInteger(b) != -1 {
		t.Error("read EOF should return -1")
	}
}

func TestIntern(t *testing.T) {
	if PrimIntern(MakeString("true")) != True {
		t.Error("intern(true) should be boolean")
	}
	if PrimIntern(MakeString("false")) != False {
		t.Error("intern(false) should be boolean")
	}
	if equal(PrimIntern(MakeString("asdf")), MakeSymbol("asdf")) != True {
		t.FailNow()
	}
}

func TestListAllPrimName(t *testing.T) {
	for _, prim := range AllPrimitives {
		fmt.Printf(" %s", prim.Name)
	}
}
