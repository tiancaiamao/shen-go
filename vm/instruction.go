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
	iPush
	iPushArg
	iPop
	iMark
	iApply
	iPrimCall
	iConst
	iReturn
	iAdd
	iHalt
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
	{iMark, "MARK"},
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
		return "GRAB"
	case iPush:
		return "PUSH"
	case iPushArg:
		return fmt.Sprintf("PUSHARG %d", instructionOP1(i))
	case iPop:
		return "POP"
	case iMark:
		return fmt.Sprintf("MARK %d", instructionOP1(i))
	case iApply:
		return "APPLY"
	case iPrimCall:
		return "PRIMCALL"
	case iConst:
		return fmt.Sprintf("CONST %d", instructionOP1(i))
	case iReturn:
		return "RETURN"
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

func (a *Assember) PUSHARG(i int) {
	inst := instruction((iPushArg << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) MARK(i int) {
	inst := instruction((iMark << codeBitShift) | (i << 16))
	a.buf = append(a.buf, inst)
}

func (a *Assember) APPLY() {
	a.buf = append(a.buf, instruction(iApply<<codeBitShift))
}

func (a *Assember) HALT() {
	a.buf = append(a.buf, instruction(iHalt<<codeBitShift))
}

func (a *Assember) GRAB() {
	a.buf = append(a.buf, instruction(iGrab<<codeBitShift))
}

func (a *Assember) ADD() {
	a.buf = append(a.buf, instruction(iAdd<<codeBitShift))
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
