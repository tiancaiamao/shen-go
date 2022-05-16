package klambda

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const part0 = `package main

import (
        "fmt"
	. "github.com/tiancaiamao/shen-go/klambda"
)

func main() {
        var ctx ControlFlow
        _ = ctx
`

func TestCodeGen(t *testing.T) {
	for _, c := range basicCases {
		t.Run(c.name, func(t *testing.T) {
			r := NewSexpReader(strings.NewReader(c.input), true)
			sexp, err := r.Read()
			if err != nil && err != io.EOF {
				panic(err)
			}

			var cc Compiler
			env := &Env{
				parent: nil,
				args:   Nil,
			}
			insts := cc.compile(sexp, env, false)

			var g GoCodeGenerator
			res := insts.GenGoCode(&g)

			f, err := os.CreateTemp("", "code_gen_*.go")
			if err != nil {
				t.Error(err)
			}
			// defer os.Remove(f.Name())
			// defer f.Close()
			fmt.Println(f.Name())

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

			if string(output) != c.output {
				fmt.Println("input exp:", c.input)
				fmt.Printf("expect: %q\n", c.output)
				fmt.Printf("actual: %q\n", string(output))
				t.Fail()
			}
		})
	}
}
