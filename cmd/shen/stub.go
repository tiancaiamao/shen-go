package main

import (
	"github.com/tiancaiamao/shen-go/kl"
)

type defType struct {
	name  string
	value kl.Obj
}

var __initExprs []kl.Obj
var __initDefs []defType

func Regist(e *kl.Evaluator) {
	for _, def := range __initDefs {
		if def.name == "symbol?" {
			continue
		}
		e.RegistNativeCall(def.name, def.value)
	}

	for _, expr := range __initExprs {
		e.Call(expr)
	}
}
