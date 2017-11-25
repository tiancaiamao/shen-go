package kl

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

var (
	initTime      time.Time
	symbolTable   map[string]Obj
	functionTable map[string]Obj
)

func init() {
	initTime = time.Now()
	symbolTable = make(map[string]Obj)
	functionTable = make(map[string]Obj, len(Primitives))
	for _, prim := range Primitives {
		functionTable[prim.Name] = Obj(&prim.scmHead)
	}

	dir, _ := os.Getwd()
	symbolTable["*stinput*"] = Make_stream(os.Stdin)
	symbolTable["*stoutput*"] = Make_stream(os.Stdout)
	symbolTable["*home-directory*"] = Make_string(dir)
	symbolTable["*language*"] = Make_string("Go")
	symbolTable["*implementation*"] = Make_string("interpreter")
	symbolTable["*relase*"] = Make_string(runtime.Version())
	symbolTable["*os*"] = Make_string(runtime.GOOS)
	symbolTable["*porters*"] = Make_string("Arthur Mao")
	symbolTable["*port*"] = Make_string("0.0.1")

	// Extended by shen-go implementation
	gopath := os.Getenv("GOPATH")
	packagePath := path.Join(gopath, "src/github.com/tiancaiamao/shen-go/pkg")
	symbolTable["*package-path*"] = Make_string(packagePath)
}

func cadr(o Obj) Obj {
	return car(cdr(o))
}

func caddr(o Obj) Obj {
	return car(cdr(cdr(o)))
}

func cdddr(o Obj) Obj {
	return cdr(cdr(cdr(o)))
}

func cadddr(o Obj) Obj {
	return car(cdr(cdr(cdr(o))))
}

func reverse(o Obj) Obj {
	ret := Nil
	for o != Nil {
		ret = cons(car(o), ret)
		o = cdr(o)
	}
	return ret
}

func equal(x, y Obj) Obj {
	if *x != *y {
		return False
	}

	switch *x {
	case Number:
		if mustNumber(x).val != mustNumber(y).val {
			return False
		}
	case String:
		if mustString(x) != mustString(y) {
			return False
		}
	case Boolean:
		if x != y {
			return False
		}
	case Symbol:
		if mustSymbol(x).sym != mustSymbol(y).sym {
			return False
		}
	case Null:
		if *y != Null {
			return False
		}
	case Pair:
		// TODO: maybe cycle exists!
		if x != y {
			if equal(car(x), car(y)) == False {
				return False
			}
			if equal(cdr(x), cdr(y)) == False {
				return False
			}
		}
	case Stream, Procedure, Primitive:
		if x != y {
			return False
		}
	}

	return True
}

func listLength(l Obj) int {
	count := 0
	for *l == Pair {
		count++
		l = cdr(l)
	}
	return count
}

func ListToSlice(l Obj) []Obj {
	var ret []Obj
	for *l == Pair {
		ret = append(ret, car(l))
		l = cdr(l)
	}
	return ret
}

func LoadFile(path string) Obj {
	f, err := os.Open(path)
	if err != nil {
		return Make_error(err.Error())
	}
	defer f.Close()

	r := NewSexpReader(f)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return Make_error(err.Error())
			}
			break
		}

		// fmt.Printf("BEGIN: %#v\n", (*scmHead)(exp))
		res := trampoline(exp, nil)
		if *res == Error {
			return res
		}
		fmt.Printf("%#v\n", (*scmHead)(res))
		// fmt.Printf("END: %#v\n", (*scmHead)(res))
	}
	return Nil
}

func GetNumber(o Obj) float64 {
	return mustNumber(o).val
}

func GetInteger(o Obj) int {
	return int(mustNumber(o).val)
}

func Cadr(o Obj) Obj {
	return cadr(o)
}

func Car(o Obj) Obj {
	return car(o)
}

func Cdr(o Obj) Obj {
	return cdr(o)
}
