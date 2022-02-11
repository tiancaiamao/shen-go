package babashka

import (
	"fmt"

	. "github.com/tiancaiamao/shen-go/cora"
)

func primStrlen(x Obj) Obj {
	str := GetString(x)
	return MakeInteger(len(str))
}

func primDisplay(x Obj) Obj {
	str := ObjString(x)
	fmt.Println(str)
	return Nil
}

func Init() {
	PrimNS1Set(MakeSymbol("strlen"), MakePrimitive("strlen", 1, primStrlen))
	PrimNS1Set(MakeSymbol("display"), MakePrimitive("display", 1, primDisplay))
}
