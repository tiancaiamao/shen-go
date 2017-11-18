package vm

import (
	"strings"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

func TestInstruction(t *testing.T) {
	var a Assember
	a.CONST(kl.Make_integer(4))
	a.CONST(kl.Make_integer(7))
	a.GRAB()
	a.MARK(5)
	a.PUSHARG(2)

	result := `CONST 0
CONST 1
GRAB
MARK 5
PUSHARG 2`

	expects := strings.Split(result, "\n")
	for i, v := range a.buf {
		if v.String() != expects[i] {
			t.Errorf("expect: %s, get %s\n", v.String(), expects[i])
		}
	}
}
