package vm

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/tiancaiamao/shen-go/kl"
)

type instruction uint32

const (
	iAccess = iota
	iGrab
	iFreeze
	iPush
	iPop
	iApply
	iMark
	iTailApply
	iPrimCall
	iConst
	iReturn
	iHalt
	iDefun
	iGetF
	iJF
	iJMP
	iSetJmp
	iClearJmp
)

type instructionInfo struct {
	id   instruction
	name string
}

var instructionTable = []instructionInfo{
	{iAccess, "ACCESS"},
	{iGrab, "GRAB"},
	{iPush, "PUSH"},
	{iPop, "POP"},
	{iApply, "APPLY"},
	{iMark, "MARK"},
	{iPrimCall, "PRIMCALL"},
	{iConst, "CONST"},
	{iReturn, "RETURN"},
	{iHalt, "HALT"},
}

var instructionStrToInfo map[string]*instructionInfo

func init() {
	instructionStrToInfo = make(map[string]*instructionInfo)
	for i := 0; i < len(instructionTable); i++ {
		v := &instructionTable[i]
		instructionStrToInfo[v.name] = v
	}
}

const (
	codeBitShift = 24
	valueBitMask = (1<<codeBitShift - 1)
)

func instructionCode(i instruction) int {
	return int(i >> codeBitShift)
}

func instructionOPN(i instruction) int {
	return int(i) & valueBitMask
}

func instructionOP2(i instruction) int {
	return int((i >> 8) & 0xff)
}

func instructionOP3(i instruction) int {
	return int(i & 0xff)
}

func (i instruction) String() string {
	switch instructionCode(i) {
	case iAccess:
		return fmt.Sprintf("ACCESS %d", instructionOPN(i))
	case iGrab:
		return fmt.Sprintf("GRAB %d", instructionOPN(i))
	case iPush:
		return "PUSH"
	case iPop:
		return "POP"
	case iApply:
		return "APPLY"
	case iTailApply:
		return fmt.Sprintf("TAILAPPLY %d", instructionOPN(i))
	case iPrimCall:
		return "PRIMCALL"
	case iConst:
		return fmt.Sprintf("CONST %d", instructionOPN(i))
	case iReturn:
		return "RETURN"
	case iHalt:
		return "HALT"
	case iDefun:
		return "DEFUN"
	case iGetF:
		return "GETF"
	case iJF:
		return fmt.Sprintf("JF %d", instructionOPN(i))
	case iJMP:
		return fmt.Sprintf("JMP %d", instructionOPN(i))
	case iFreeze:
		return fmt.Sprintf("FREEZE %d", instructionOPN(i))
	case iSetJmp:
		return fmt.Sprintf("SETJMP %d", instructionOPN(i))
	case iClearJmp:
		return "CLEARJMP"
	case iMark:
		return "MARK"
	}
	return "UNKNOWN"
}

type Assember struct {
	buf    []instruction
	consts []kl.Obj
}

func (a *Assember) ACCESS(i int) {
	inst := instruction((iAccess << codeBitShift) | i)
	a.buf = append(a.buf, inst)
}

func (a *Assember) RETURN() {
	a.buf = append(a.buf, instruction(iReturn<<codeBitShift))
}

func (a *Assember) APPLY() {
	a.buf = append(a.buf, instruction(iApply<<codeBitShift))
}

func (a *Assember) TAILAPPLY() {
	a.buf = append(a.buf, instruction(iTailApply<<codeBitShift))
}

func (a *Assember) HALT() {
	a.buf = append(a.buf, instruction(iHalt<<codeBitShift))
}

func (a *Assember) CLEARJMP() {
	a.buf = append(a.buf, instruction(iClearJmp<<codeBitShift))
}

func (a *Assember) DEFUN() {
	a.buf = append(a.buf, instruction(iDefun<<codeBitShift))
}

func (a *Assember) FREEZE(i int) {
	inst := instruction((iFreeze << codeBitShift) | i)
	a.buf = append(a.buf, inst)
}
func (a *Assember) GRAB(i int) {
	inst := instruction((iGrab << codeBitShift) | i)
	a.buf = append(a.buf, inst)
}

func (a *Assember) JF(i int) {
	if i >= (1 << codeBitShift) {
		panic("overflow instruct bits")
	}
	inst := instruction((iJF << codeBitShift) | i)
	a.buf = append(a.buf, inst)
}

func (a *Assember) JMP(i int) {
	if i >= (1 << codeBitShift) {
		panic("overflow instruct bits")
	}
	inst := instruction((iJMP << codeBitShift) | i)
	a.buf = append(a.buf, inst)
}

func (a *Assember) SETJMP(i int) {
	if i >= (1 << codeBitShift) {
		panic("overflow instruct bits")
	}
	inst := instruction((iSetJmp << codeBitShift) | i)
	a.buf = append(a.buf, inst)
}

func (a *Assember) POP() {
	a.buf = append(a.buf, instruction(iPop<<codeBitShift))
}

func (a *Assember) MARK() {
	a.buf = append(a.buf, instruction(iMark<<codeBitShift))
}

func (a *Assember) GetF() {
	a.buf = append(a.buf, instruction(iGetF<<codeBitShift))
}

func (a *Assember) CONST(o kl.Obj) {
	idx := len(a.consts)
	a.consts = append(a.consts, o)
	if idx >= (1 << codeBitShift) {
		panic("overflow instruct bits")
	}
	inst := instruction((iConst << codeBitShift) | idx)
	a.buf = append(a.buf, inst)
}

func (a *Assember) PRIMCALL(id int) {
	if id >= (1 << codeBitShift) {
		panic("overflow instruct bits")
	}
	inst := instruction((iPrimCall << codeBitShift) | id)
	a.buf = append(a.buf, inst)
}

func (a *Assember) Comiple() *Code {
	code := Code{
		bc:     a.buf,
		consts: a.consts,
	}
	a.buf = nil
	a.consts = nil
	return &code
}

func (a *Assember) Encode(str string) (*Code, error) {
	inputs := strings.Split(str, "\n")
	for _, input := range inputs {
		var inst string
		_, err := fmt.Sscanf(input, "%s ", &inst)
		if err != nil {
			return nil, err
		}
		_, ok := instructionStrToInfo[inst]
		if !ok {
			return nil, errors.New("invalid instruct " + inst)
		}
	}
	return nil, errors.New("fuck")
}

func (a *Assember) Decode(code *Code) string {
	var buf bytes.Buffer
	for _, bc := range code.bc {
		fmt.Fprintln(&buf, bc.String())
	}
	return buf.String()
}

func (a *Assember) FromSexp(input kl.Obj) error {
	objs := kl.ListToSlice(input)
	for _, obj := range objs {
		// fmt.Println(kl.ObjString(kl.Car(obj)))
		id := kl.GetSymbol(kl.Car(obj))
		switch id {
		case "iAccess":
			n := kl.GetInteger(kl.Cadr(obj))
			a.ACCESS(n)
		case "iPrimCall":
			n := kl.GetInteger(kl.Cadr(obj))
			a.PRIMCALL(n)
		case "iConst":
			a.CONST(kl.Cadr(obj))
		case "iApply":
			a.APPLY()
		case "iTailApply":
			a.TAILAPPLY()
		case "iMark":
			a.MARK()
		case "iReturn":
			a.RETURN()
		case "iGrab":
			var a1 Assember
			a1.FromSexp(kl.Cdr(obj))
			a.GRAB(len(a1.buf))
			adjustConst(a1.buf, len(a.consts))
			a.buf = append(a.buf, a1.buf...)
			a.consts = append(a.consts, a1.consts...)
		case "iFreeze":
			var a1 Assember
			a1.FromSexp(kl.Cdr(obj))
			a.FREEZE(len(a1.buf))
			adjustConst(a1.buf, len(a.consts))
			a.buf = append(a.buf, a1.buf...)
			a.consts = append(a.consts, a1.consts...)
		case "iHalt":
			a.HALT()
		case "iDefun":
			a.DEFUN()
		case "iClearJmp":
			a.CLEARJMP()
		case "iPop":
			a.POP()
		case "iGetF":
			a.GetF()
		case "iJF":
			var a1 Assember
			a1.FromSexp(kl.Cdr(obj))
			a.JF(len(a1.buf) + 1) // Follow by a JMP
			adjustConst(a1.buf, len(a.consts))
			a.buf = append(a.buf, a1.buf...)
			a.consts = append(a.consts, a1.consts...)
		case "iJMP":
			var a1 Assember
			a1.FromSexp(kl.Cdr(obj))
			a.JMP(len(a1.buf))
			adjustConst(a1.buf, len(a.consts))
			a.buf = append(a.buf, a1.buf...)
			a.consts = append(a.consts, a1.consts...)
		case "iSetJmp":
			var a1 Assember
			a1.FromSexp(kl.Cdr(obj))
			a.SETJMP(len(a1.buf))
			adjustConst(a1.buf, len(a.consts))
			a.buf = append(a.buf, a1.buf...)
			a.consts = append(a.consts, a1.consts...)
		}
	}
	return nil
}

func adjustConst(insts []instruction, ofst int) {
	for i, inst := range insts {
		if instructionCode(inst) == iConst {
			idx := instructionOPN(inst) + ofst
			if i >= (1 << codeBitShift) {
				panic("overflow instruct bits")
			}
			insts[i] = instruction((iConst << codeBitShift) | idx)
		}

	}
}
