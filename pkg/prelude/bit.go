package main

import (
	"github.com/tiancaiamao/shen-go/kl"
)

func BitLeftShift(a, b kl.Obj) kl.Obj {
	if *a != kl.Number {
		return kl.Make_error("xxx")
	}
	return kl.Make_integer(42)
}
