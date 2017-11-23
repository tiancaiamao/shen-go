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
	iPushArg
	iPop
	iApply
	iTailApply
	iPrimCall
	iConst
	iReturn
	iAdd
	iHalt
	iDefun
	iGetF
	iJF
	iJMP
	iSetJmp
)

type instructionInfo struct {
	id   instruction
	name string
}

var instructionTable = []instructionInfo{
	{iAccess, "ACCESS"},
	{iGrab, "GRAB"},
	{iPush, "PUSH"},
	{iPushArg, "PUSHARG"},
	{iPop, "POP"},
	{iApply, "APPLY"},
	{iPrimCall, "PRIMCALL"},
	{iConst, "CONST"},
	{iReturn, "RETURN"},
	{iAdd, "ADD"},
	{iAdd, "HALT"},
}

var instructionStrToInfo map[string]*instructionInfo

func init() {
	instructionStrToInfo = make(map[string]*instructionInfo)
	for i := 0; i < len(instructionTable); i++ {
		v := &instructionTable[i]
		instructionStrToInfo[v.name] = v
	}
}

const codeBitShift = 24

func instructionCode(i instruction) int {
	return int(i >> codeBitShift)
}

func instructionOP1(i instruction) int {
	return int((i >> 16) & 0xff)
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
		return fmt.Sprintf("ACCESS %d", instructionOP1(i))
	case iGrab:
		return fmt.Sprintf("GRAB %d", instructionOP1(i))
	case iPush:
		return "PUSH"
	case iPushArg:
		return fmt.Sprintf("PUSHARG")
	case iPop:
		return "POP"
	case iApply:
		return "APPLY"
	case iTailApply:
		return "TAILAPPLY"
	case iPrimCall:
		return "PRIMCALL"
	case iConst:
		return fmt.Sprintf("CONST %d", instructionOP1(i))
	case iReturn:
		return "RETURN"
	case iHalt:
		return "HALT"
	case iDefun:
		return "DEFUN"
	case iGetF:
		return "GETF"
	case iJF:
		return fmt.Sprintf("JF %d", instructionOP1(i))
	case iJMP:
		return fmt.Sprintf("JMP %d", instructionOP1(i))
	case iFreeze:
		return fmt.Sprintf("FREEZE %d", instructionOP1(i))
	case iSetJmp:
		return fmt.Sprintf("SETJMP %d", instructionOP1(i))
	}
	return "UNKNOWN"
}

type Assember struct {
	buf    []instruction
	consts []kl.Obj
}

func (a *Assember) ACCESS(i int) {
	inst := instruction((iAccess << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) RETURN() {
	a.buf = append(a.buf, instruction(iReturn<<codeBitShift))
}

func (a *Assember) PUSHARG() {
	a.buf = append(a.buf, instruction(iPushArg<<codeBitShift))
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

func (a *Assember) DEFUN() {
	a.buf = append(a.buf, instruction(iDefun<<codeBitShift))
}

func (a *Assember) FREEZE(i int) {
	inst := instruction((iFreeze << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}
func (a *Assember) GRAB(i int) {
	inst := instruction((iGrab << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) JF(i int) {
	inst := instruction((iJF << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) JMP(i int) {
	inst := instruction((iJMP << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) SETJMP(i int) {
	inst := instruction((iSetJmp << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) ADD() {
	a.buf = append(a.buf, instruction(iAdd<<codeBitShift))
}

func (a *Assember) POP() {
	a.buf = append(a.buf, instruction(iPop<<codeBitShift))
}

func (a *Assember) GetF() {
	a.buf = append(a.buf, instruction(iGetF<<codeBitShift))
}

func (a *Assember) CONST(o kl.Obj) {
	idx := len(a.consts)
	a.consts = append(a.consts, o)
	inst := instruction((iConst << codeBitShift) | (idx << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) PRIMCALL(id int) {
	inst := instruction((iPrimCall << codeBitShift) | (id << 16))
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
		id := kl.SymbolString(kl.Car(obj))
		switch id {
		case "iAccess":
			n := kl.GetInteger(kl.Cadr(obj))
			a.ACCESS(n)
		case "iPrimCall":
			n := kl.GetInteger(kl.Cadr(obj))
			a.PRIMCALL(n)
		case "iPushArg":
			a.PUSHARG()
		case "iConst":
			a.CONST(kl.Cadr(obj))
		case "iApply":
			a.APPLY()
		case "iTailApply":
			a.TAILAPPLY()
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
			idx := instructionOP1(inst) + ofst
			insts[i] = instruction((iConst << codeBitShift) | (idx << 16))
		}

	}
}
