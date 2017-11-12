package kl

import (
	"os"
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
	functionTable = make(map[string]Obj, len(primitives))
	for _, prim := range primitives {
		functionTable[prim.name] = Obj(&prim.scmHead)
	}

	dir, _ := os.Getwd()
	symbolTable["*stinput*"] = Make_stream(os.Stdin)
	symbolTable["*stoutput*"] = Make_stream(os.Stdout)
	symbolTable["*home-directory*"] = Make_string(dir)
	symbolTable["*language*"] = Make_string("Go")
	symbolTable["*implementation*"] = Make_string(runtime.Compiler)
	symbolTable["*relase*"] = Make_string(runtime.Version())
	symbolTable["*os*"] = Make_string(runtime.GOOS)
	symbolTable["*porters*"] = Make_string("Arthur Mao")
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

func listToSlice(l Obj) []Obj {
	var ret []Obj
	for *l == Pair {
		ret = append(ret, car(l))
		l = cdr(l)
	}
	return ret
}
