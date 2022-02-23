package klambda

import (
	"math"
	"unsafe"
)

func cadr(o Obj) Obj {
	return car(cdr(o))
}

func caddr(exp Obj) Obj {
	return Car(Cdr(Cdr(exp)))
}

func cdddr(o Obj) Obj {
	return cdr(cdr(cdr(o)))
}

func cadddr(o Obj) Obj {
	return car(cdr(cdr(cdr(o))))
}

// func Reverse(o Obj) Obj {
// 	return reverse(o)
// }

func reverse(o Obj) Obj {
	ret := Nil
	for o != Nil {
		ret = cons(car(o), ret)
		o = cdr(o)
	}
	return ret
}

// func listAppend(l1, l2 Obj) Obj {
// 	res := Nil
// 	for ; l1 != Nil; l1 = cdr(l1) {
// 		res = cons(car(l1), res)
// 	}
// 	for ; l2 != Nil; l2 = cdr(l2) {
// 		res = cons(car(l2), res)
// 	}
// 	return reverse(res)
// }

func equal(x, y Obj) Obj {
	if x == y {
		return True
	}
	if isFixnum(uintptr(unsafe.Pointer(x))) || isFixnum(uintptr(unsafe.Pointer(y))) || *x != *y {
		return False
	}

	switch *x {
	case scmHeadNumber:
		if !isFixnum(uintptr(unsafe.Pointer(x))) && !isFixnum(uintptr(unsafe.Pointer(y))) {
			if mustNumber(x) == mustNumber(y) {
				return True
			}
		}
		// x == y is checked already
		return False
	case scmHeadString:
		if mustString(x) != mustString(y) {
			return False
		}
	case scmHeadBoolean:
		if x != y {
			return False
		}
	case scmHeadSymbol:
		if x != y {
			return False
		}
	case scmHeadNull:
		if *y != scmHeadNull {
			return False
		}
	case scmHeadPair:
		// TODO: maybe cycle exists!
		if x != y {
			if equal(car(x), car(y)) == False {
				return False
			}
			if equal(cdr(x), cdr(y)) == False {
				return False
			}
		}
	case scmHeadStream, scmHeadNative:
		if x != y {
			return False
		}
	case scmHeadVector:
		v1 := mustVector(x)
		v2 := mustVector(y)
		if len(v1) != len(v2) {
			return False
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != nil || v2[i] != nil {
				if equal(v1[i], v2[i]) != True {
					return False
				}
			}
		}
	}

	return True
}

func listLength(l Obj) int {
	count := 0
	for *l == scmHeadPair {
		count++
		l = cdr(l)
	}
	return count
}

// func Cadr(o Obj) Obj {
// 	return cadr(o)
// }

func Car(o Obj) Obj {
	return car(o)
}

func Cdr(o Obj) Obj {
	return cdr(o)
}

func Cons(x, y Obj) Obj {
	return cons(x, y)
}

func IsCons(x Obj) bool {
	if objType(x) == scmHeadPair {
		return true
	}
	return false
}

// isInteger determinate whether a float64 is actually a precise integer.
// Judge is according to IEEE754 standard.
func isPreciseInteger(f float64) bool {
	exp := math.Ilogb(f)
	if exp < 0 && exp != math.MinInt32 {
		return false
	}

	if exp >= 52 {
		return true
	}

	bits := math.Float64bits(f)
	return (bits << uint(12+exp)) == 0
}