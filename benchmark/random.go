package main

import (
	"math/rand"
	. "github.com/tiancaiamao/shen-go/kl"
)

func Main(__e Evaluator, __args ...Obj) {
	res := CoraSet(MakeSymbol("random"), MakeNative(func(__e Evaluator, __args ...Obj) {
		n := GetInteger(__args[0])
		res := rand.Intn(n)
		__e.Return(MakeInteger(res))
	}, 1))
	__e.Return(res)
	return
}
