package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/tiancaiamao/shen-go/kl"
	"github.com/tiancaiamao/shen-go/vm"
)

func BitLeftShift(args ...kl.Obj) kl.Obj {
	if !(kl.IsNumber(args[0]) && kl.IsNumber(args[1])) {
		return kl.MakeError("xxx")
	}
	return kl.MakeInteger(42)
}

type rawTime struct {
	scmHead int
	time.Time
}

func Now(args ...kl.Obj) kl.Obj {
	var tmp rawTime
	ret := kl.MakeRaw(&tmp.scmHead)
	tmp.Time = time.Now()
	return ret
}

func TimeSub(args ...kl.Obj) kl.Obj {
	t1 := (*rawTime)(unsafe.Pointer(args[0]))
	t2 := (*rawTime)(unsafe.Pointer(args[1]))
	duration := t2.Time.Sub(t1.Time)
	return kl.MakeString(duration.String())
}

func Main(vm *vm.VM) {
	fmt.Println("hello world")
	vm.RegistNativeCall("now", 0, Now)
	vm.RegistNativeCall("time-sub", 2, TimeSub)
}
