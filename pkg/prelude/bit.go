package main

import (
	"github.com/tiancaiamao/shen-go/kl"
)

func BitLeftShift(a, b kl.Obj) kl.Obj {
	if !(kl.IsNumber(a) && kl.IsNumber(b)) {
		return kl.MakeError("xxx")
	}
	return kl.MakeInteger(42)
}
