package klambda

import (
	"fmt"
	"bytes"
)

type GoCodeGenerator struct {
	bytes.Buffer
	cur int
}

func (g *GoCodeGenerator) newReg() int {
	v := g.cur
	g.cur++
	return v
}

func (i *ifInst) GenGoCode(g *GoCodeGenerator) int {
	// The generated code looks like follows:
	//	var regX Obj
	//	switch regX {
	//	case True:
	//	case False:
	//	default:
	//		panic("")
	//	}
	vTest := i.a.GenGoCode(g)
	vRes := g.newReg()
	fmt.Fprintf(g, "var _reg%d Obj\n", vRes)
	fmt.Fprintf(g, "switch _reg%d {\n", vTest)
	fmt.Fprintf(g, "case True:\n")
	vThen := i.b.GenGoCode(g)
	fmt.Fprintf(g, "_reg%d = _reg%d\n", vRes, vThen)
	fmt.Fprintf(g, "case False:\n")
	vElse := i.c.GenGoCode(g)
	fmt.Fprintf(g, "_reg%d = _reg%d\n", vRes, vElse)
	fmt.Fprintf(g, "default:\n")
	fmt.Fprintf(g, `	panic("if test branch not boolean")
`)
	fmt.Fprintf(g, "}\n")
	return vRes
}

func (i *doInst) GenGoCode(g *GoCodeGenerator) int {
	res := i.op1.GenGoCode(g)
	fmt.Fprintf(g, "_ = _reg%d", res)
	return i.op2.GenGoCode(g)
}

func (idx localRefInst) GenGoCode(g *GoCodeGenerator) int {
	res := g.newReg()
	fmt.Fprintf(g, "_reg%d = ctx.stack[%d]\n", res, idx)
	return res
}

func (i constInst) GenGoCode(g *GoCodeGenerator) int {
	res := g.newReg()
	c := i.val
	switch {
	case IsNumber(c):
		fmt.Fprintf(g, "_reg%d := MakeNumber(%d)\n", res, GetInteger(c))
	case IsString(c):
		str := GetString(c)
		fmt.Fprintf(g, "_reg%d := MakeString(%#v)\n", res, str)
	case IsSymbol(c):
		fmt.Fprintf(g, "_reg%d := sym%s\n", res, symbolAsVar(c))
	case c == Nil:
		fmt.Fprintf(g, "_reg%d := Nil", res)
	case c == True:
		fmt.Fprintf(g, "_reg%d := True", res)
	case c == False:
		fmt.Fprintf(g, "_reg%d := False", res)
	default:
		panic("unknown $const instruct")
	}
	return res
}


func (i closureRefInst) GenGoCode(g *GoCodeGenerator) int {
	fmt.Fprintf(g, "TODO: closureRefInst(%d, %d)\n", i.m, i.n)
	return g.cur
}

func (i globalRefInst)  GenGoCode(g *GoCodeGenerator) int {
	fmt.Fprintf(g, "TODO: globalRef\n")
	return 0
}

func (i klGlobalRefInst)  GenGoCode(g *GoCodeGenerator) int {
	res := g.newReg()
	fmt.Fprintf(g, "_reg%d := PrimNS2Value(MakeSymbol(%q))\n", res, ObjString(i.sym))
	return res
}

func (l letInst) GenGoCode(g *GoCodeGenerator) int {
	res := l.arg.GenGoCode(g)
	fmt.Fprintf(g, "ctx.stack[ebp + %d] = _reg%d\n", l.idx, res)
	return l.body.GenGoCode(g)
}

func (l *makeClosureInst) GenGoCode(g *GoCodeGenerator) int {
	res := g.newReg()
	fmt.Fprintf(g, "_reg%d := MakeNative(func (ctx *ControlFlow, ebp, esp int) {\n", res)
	tmp := l.op.GenGoCode(g)
	fmt.Fprintf(g, "ctx.val = _reg%d\n", tmp)
	fmt.Fprintf(g, "}")
	fmt.Fprintf(g, ", %d", l.nargs)
	if len(l.mark) > 0 {
		fmt.Fprintf(g, ", TODO: args)\n")
	} else {
		fmt.Fprintf(g, ")\n")
	}
	return res
}

func (l klGlobalSetInst) GenGoCode(g *GoCodeGenerator) int {
	res := l.f.GenGoCode(g)
	fmt.Fprintf(g, "PrimNS2Set(%s, _reg%d)\n", symbolAsVar(l.name), res)
	return res
}

func (l callInst) GenGoCode(g *GoCodeGenerator) int {
	regs := make([]int, len(l.insts))
	for i:=0; i<len(l.insts); i++ {
		regs[i] = l.insts[i].GenGoCode(g)
	}
	res := g.newReg()
	fmt.Fprintf(g, "_reg%d := ctx.Call(", res)
	for i, reg := range regs {
		if i == 0 {
			fmt.Fprintf(g, "_reg%d", reg)
		} else {
			fmt.Fprintf(g, ", _reg%d", reg)
		}
	}
	fmt.Fprintf(g, ")\n")
	return res
}

func symbolAsVar(sym Obj) string {
	str := GetSymbol(sym)
	var buf bytes.Buffer
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '_':
			buf.WriteString("__")
		case '-':
			buf.WriteString("_1")
		case '?':
			buf.WriteString("_2")
		case '$':
			buf.WriteString("_3")
		case '.':
			buf.WriteString("_4")
		case '<':
			buf.WriteString("_5")
		case '>':
			buf.WriteString("_6")
		case '+':
			buf.WriteString("_7")
		case '@':
			buf.WriteString("_8")
		case '=':
			buf.WriteString("_a")
		case '!':
			buf.WriteString("_b")
		case '/':
			buf.WriteString("_c")
		case '*':
			buf.WriteString("_d")
		case '&':
			buf.WriteString("_e")
		case '%':
			buf.WriteString("_f")
		case '^':
			buf.WriteString("_g")
		case ':':
			buf.WriteString("_h")
		case '{':
			buf.WriteString("_i")
		case '}':
			buf.WriteString("_j")
		case ';':
			buf.WriteString("_k")
		case ',':
			buf.WriteString("_l")
		default:
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}
