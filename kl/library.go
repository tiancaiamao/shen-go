package kl

import (
	"os"
	"path"
)

func PackagePath() string {
	gopath := os.Getenv("GOPATH")
	return path.Join(gopath, "src/github.com/tiancaiamao/shen-go")
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

func Cons(x, y Obj) Obj {
	return cons(x, y)
}
