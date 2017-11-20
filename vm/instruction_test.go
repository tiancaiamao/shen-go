package vm

import (
	// "fmt"
	"strings"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestInstruction(t *testing.T) {
	var a Assember
	a.CONST(kl.Make_integer(4))
	a.CONST(kl.Make_integer(7))
	a.GRAB(2)
	a.PUSHARG()
	a.DEFUN()
	a.POP()
	a.HALT()

	result := `CONST 0
CONST 1
GRAB 2
PUSHARG
DEFUN
POP
HALT`

	expects := strings.Split(result, "\n")
	for i, v := range a.buf {
		if v.String() != expects[i] {
			t.Errorf("expect: %s, get %s\n", v.String(), expects[i])
		}
	}
}

func TestFromSexp(t *testing.T) {
	r := kl.NewSexpReader(strings.NewReader("((iGrab (iGrab (iAccess 1) (iAccess 0) (iPrimCall 2) (iReturn)) (iReturn)) (iConst 1) (iPushArg) (iConst 2) (iPushArg) (iApply) (iHalt))"))
	sexp, err := r.Read()
	if err != nil {
		t.Error(err)
	}

	var a Assember
	err = a.FromSexp(sexp)
	if err != nil {
		t.Error(err)
	}

	code := a.Comiple()
	str := a.Decode(code)
	// fmt.Println(str)
	result := `GRAB 6
GRAB 4
ACCESS 1
ACCESS 0
PRIMCALL
RETURN
RETURN
CONST 0
PUSHARG
CONST 1
PUSHARG
APPLY
HALT
`
	if str != result {
		t.Errorf("get: %s\nexcept:%s\n", str, result)
	}

	// expects := strings.Split(result, "\n")
	// for i, v := range a.buf {
	// 	if v.String() != expects[i] {
	// 		t.Errorf("expect: %s, get %s\n", v.String(), expects[i])
	// 	}
	// }
}
