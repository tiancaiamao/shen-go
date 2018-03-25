package runtime

import (
	"strings"
	"testing"
)

func TestInstruction(t *testing.T) {
	var a assember
	a.CONST(MakeInteger(4))
	a.CONST(MakeInteger(7))
	a.GRAB()
	a.DEFUN()
	a.POP()
	a.JF(1)
	a.MARK()
	a.JMP(5)
	a.FREEZE(3)
	a.SETJMP(22)
	a.NATIVECALL(4)
	a.HALT()

	result := `CONST 0
CONST 1
GRAB
DEFUN
POP
JF 1
MARK
JMP 5
FREEZE 3
SETJMP 22
NATIVECALL 4
HALT`

	expects := strings.Split(result, "\n")
	for i, v := range a.buf {
		if v.String() != expects[i] {
			t.Errorf("expect: %s, get %s\n", v.String(), expects[i])
		}
	}
}

func TestFromSexp(t *testing.T) {
	r := NewSexpReader(strings.NewReader("((iGrab) (iGrab) (iAccess 1) (iAccess 0) (iPrimCall 2) (iReturn) (iReturn) (iConst 1) (iConst 2) (iApply) (iHalt))"))
	sexp, err := r.Read()
	if err != nil {
		t.Error(err)
	}

	var a assember
	err = a.FromSexp(sexp)
	if err != nil {
		t.Error(err)
	}
}
