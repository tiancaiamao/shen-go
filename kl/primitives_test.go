package kl

import (
	"bytes"
	"testing"
)

func TestReadByte(t *testing.T) {
	buf := bytes.NewBufferString("a")
	stream := Make_stream(buf)
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
	if primIntern(Make_string("true")) != True {
		t.Error("intern(true) should be boolean")
	}
	if primIntern(Make_string("false")) != False {
		t.Error("intern(false) should be boolean")
	}
	if equal(primIntern(Make_string("asdf")), Make_symbol("asdf")) != True {
		t.FailNow()
	}
}
