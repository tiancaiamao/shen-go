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
