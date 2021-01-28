package kl

import (
	"bytes"
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

// func TestListAllPrimName(t *testing.T) {
// 	for _, prim := range AllPrimitives {
// 		fmt.Printf(" %s", prim.Name)
// 	}
// }

func TestStr(t *testing.T) {
	// str primitive prints the viewable format of a object.
	// shen symbol? defun rely on this function contain non-alpha chars.
	var alphaTable [256]bool
	for _, c := range []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'=', '*', '/', '+', '-', '_', '?', '$', '!', '@', '~', '>', '<',
		'&', '%', '{', '}', ':', ';', '`', '#', '\'', '.'} {
		alphaTable[c] = true
	}
	allAlpha := func(str string) bool {
		for i := 0; i < len(str); i++ {
			c := str[i]
			if alphaTable[c] != true {
				return false
			}
		}
		return true
	}

	str := PrimStr(makeClosure(MakeSymbol("x"), MakeSymbol("x"), Nil))
	if allAlpha(mustString(str)) {
		t.Error("str of procedure should not be all alpha")
	}
}

func TestFixnum(t *testing.T) {
	n := MakeNumber(3000001)
	if PrimIsPair(n) != False {
		t.Error("3000000 should not be a cons")
	}
}
