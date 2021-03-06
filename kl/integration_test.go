package kl

import (
	"testing"
)

func TestIntegration(t *testing.T) {
	var e ControlFlow
	PrimNS1Set(MakeSymbol("check"), MakePrimitive("check", 2, func(res, expect Obj) Obj {
		if PrimEqual(res, expect) != True {
			t.Fail()
		}
		return True
	}))
	PrimNS1Set(symMacroExpand, Nil)

	Call(&e, PrimNS1Value(MakeSymbol("cora.init")))
	Call(&e, PrimNS1Value(MakeSymbol("load")), MakeString("issue_25.cora"))
}
