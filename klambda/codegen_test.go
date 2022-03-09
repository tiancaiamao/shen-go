package klambda

import (
	"os"
	"os/exec"
	"io"
	"fmt"
	"testing"
	"strings"
)

const part0 = `
package main

import (
        "fmt"
	. "github.com/tiancaiamao/shen-go/klambda"
)

func main() {
        var ctx ControlFlow
`

func TestCodeGen(t *testing.T) {
	input := "(+ 4 2)"
	// input := `(lambda x
	// 	(if (let a 3
	// 		(> (+ a x) 5))
	// 		4
	// 		7))`

	r := NewSexpReader(strings.NewReader(input), true)
	sexp, err := r.Read()
	if err != nil && err != io.EOF {
		panic(err)
	}

	var cc Compiler
	env := &Env{
		parent: nil,
		args:   Nil,
	}
	insts := cc.compile(sexp, env, true)

	var g GoCodeGenerator
	res := insts.GenGoCode(&g)

	
	f, err := os.CreateTemp("", "code_gen_*.go")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	// fmt.Println(f.Name())

	io.WriteString(f, part0)
	io.Copy(f, &g.Buffer)

	fmt.Fprintf(f, `fmt.Print(ObjString(_reg%d))
}
`, res)

	cmd := exec.Command("go", "run", f.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("go run file.go failed")
		t.Error(err)
	}

	expect := "6"
	if string(output) != expect {
		fmt.Printf("expect: %q\n", expect)
		fmt.Printf("actual: %q\n", string(output))
		t.Fail()
	}
}
