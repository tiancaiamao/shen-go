package kl

// klCompiler compiles a KL form to bytecode.
type klCompiler struct {
	fn      *BytecodeFunc
	locals  map[Obj]int    // symbol pointer -> slot index
	upvals  []upvalInfo    // upvalues captured from outer scope
	outer   *klCompiler    // enclosing compiler (for closures)
}

type upvalInfo struct {
	sym      Obj
	outerRef varRef // where to find it in the outer scope
}

type varRef struct {
	kind  varKind
	index int
}

type varKind int

const (
	varLocal  varKind = iota
	varUpval
	varGlobal
)

// CompileFunc compiles a KL top-level defun body to a BytecodeFunc object.
// params is the list of parameter symbols; body is the KL expression.
func CompileFunc(name string, params []Obj, body Obj) Obj {
	c := newCompiler(name, len(params), params, nil)
	c.compileExpr(body, true)
	return makeBytecodeObj(c.fn, nil)
}

func newCompiler(name string, arity int, params []Obj, outer *klCompiler) *klCompiler {
	c := &klCompiler{
		fn: &BytecodeFunc{
			Name:    name,
			Arity:   arity,
			Nlocals: arity,
		},
		locals: make(map[Obj]int),
		outer:  outer,
	}
	for i, p := range params {
		c.locals[p] = i
	}
	return c
}

func (c *klCompiler) emit(op uint8, a, b int32) int {
	idx := len(c.fn.Code)
	c.fn.Code = append(c.fn.Code, Instr{Op: op, A: a, B: b})
	return idx
}

func (c *klCompiler) addConst(v Obj) int32 {
	for i, cv := range c.fn.Consts {
		if cv == v {
			return int32(i)
		}
	}
	c.fn.Consts = append(c.fn.Consts, v)
	return int32(len(c.fn.Consts) - 1)
}

func (c *klCompiler) newLocal(sym Obj) int32 {
	slot := int32(c.fn.Nlocals)
	c.fn.Nlocals++
	c.locals[sym] = int(slot)
	return slot
}

func (c *klCompiler) resolveVar(sym Obj) (varKind, int) {
	if idx, ok := c.locals[sym]; ok {
		return varLocal, idx
	}
	for i, uv := range c.upvals {
		if uv.sym == sym {
			return varUpval, i
		}
	}
	if c.outer != nil {
		outerKind, outerIdx := c.outer.resolveVar(sym)
		if outerKind == varGlobal {
			return varGlobal, 0
		}
		uvIdx := len(c.upvals)
		c.upvals = append(c.upvals, upvalInfo{
			sym:      sym,
			outerRef: varRef{kind: outerKind, index: outerIdx},
		})
		return varUpval, uvIdx
	}
	return varGlobal, 0
}

// patchJump patches a previously emitted jump instruction's A operand.
func (c *klCompiler) patchJump(instrIdx int) {
	target := len(c.fn.Code)
	c.fn.Code[instrIdx].A = int32(target - instrIdx - 1)
}

// compileExpr compiles a KL expression.  tail indicates whether this is
// in tail position (may emit OP_TAIL_CALL / OP_RETURN instead of OP_CALL).
func (c *klCompiler) compileExpr(form Obj, tail bool) {
	switch *form {
	case scmHeadNumber, scmHeadString, scmHeadBoolean, scmHeadNull:
		c.emit(OP_LOAD_CONST, c.addConst(form), 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return

	case scmHeadSymbol:
		kind, idx := c.resolveVar(form)
		switch kind {
		case varLocal:
			c.emit(OP_LOAD_LOCAL, int32(idx), 0)
		case varUpval:
			c.emit(OP_LOAD_UPVAL, int32(idx), 0)
		case varGlobal:
			// In value position, a symbol not in scope evaluates to itself.
			c.emit(OP_LOAD_CONST, c.addConst(form), 0)
		}
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return

	case scmHeadPair:
		c.compileForm(form, tail)
		return

	default:
		// Vector, stream, procedure, native, etc. — treat as constants.
		c.emit(OP_LOAD_CONST, c.addConst(form), 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return
	}
}

// compileForm handles pair forms: special forms and function calls.
func (c *klCompiler) compileForm(form Obj, tail bool) {
	pair := mustPair(form)
	if !IsSymbol(pair.car) {
		// Complex function expression: (expr args...)
		c.compileCall(pair.car, pair.cdr, tail)
		return
	}
	head := pair.car
	rest := pair.cdr

	switch head {
	case symDefun:
		// (defun f (x...) body)
		name := car(rest)
		params := cadr(rest)
		body := caddr(rest)
		c.compileDefun(name, params, body)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}

	case symLambda:
		// (lambda x body)
		param := car(rest)
		body := cadr(rest)
		c.compileLambda([]Obj{param}, body)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}

	case symFreeze:
		// (freeze body)
		body := car(rest)
		c.compileLambda(nil, body)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}

	case symLet:
		// (let x val body)
		x := car(rest)
		val := cadr(rest)
		body := caddr(rest)
		c.compileLet(x, val, body, tail)

	case symIf:
		args := rest
		if listLength(args) == 3 {
			c.compileIf(car(args), cadr(args), caddr(args), tail)
		} else {
			// (if ...) with wrong arg count: fall through as a call
			c.compileCall(head, rest, tail)
		}

	case symAnd:
		// (and a b) => (if a b false)
		a := car(rest)
		b := cadr(rest)
		c.compileIf(a, b, False, tail)

	case symOr:
		// (or a b) => (if a true b)
		a := car(rest)
		b := cadr(rest)
		c.compileIf(a, True, b, tail)

	case symCond:
		c.compileCond(rest, tail)

	case symDo:
		// (do a b)
		a := car(rest)
		b := cadr(rest)
		c.compileExpr(a, false)
		c.emit(OP_POP, 0, 0)
		c.compileExpr(b, tail)

	case symType:
		// (type x _) — just compile x, ignore the type annotation
		c.compileExpr(car(rest), tail)

	case symTrapError:
		// (trap-error body handler)
		// Lower to: (try-catch (freeze body) handler)
		body := car(rest)
		handler := cadr(rest)
		c.compileTrapError(body, handler, tail)

	default:
		// Regular function call with a symbol in head position.
		c.compileCall(head, rest, tail)
	}
}

func (c *klCompiler) compileDefun(name, params, body Obj) {
	paramSlice := ListToSlice(params)
	nameStr := mustSymbol(name).str
	inner := newCompiler(nameStr, len(paramSlice), paramSlice, nil) // top-level: no outer
	inner.compileExpr(body, true)
	bfObj := makeBytecodeObj(inner.fn, nil)
	// Emit: bind and push the name symbol
	// (defun f ...) at runtime: compile to bytecode and bind, then push name.
	// We emit this as a constant + global store via LOAD_CONST + native call,
	// but actually we can just emit a LOAD_CONST of the compiled func and
	// a call to the defun primitive.  For simplicity, use the approach of
	// OP_LOAD_CONST bfObj, OP_LOAD_CONST nameSym, then swap + call defun.
	// Even simpler: use a special compilation-time action.
	//
	// Actually the cleanest approach: emit code that at runtime calls
	// BindSymbolFunc(name, bfObj).  We do this by pushing two consts and
	// calling the built-in "defun" primitive.
	symConst := c.addConst(name)
	fnConst := c.addConst(bfObj)
	defunSym := c.addConst(MakeSymbol("defun"))
	c.emit(OP_LOAD_GLOBAL, defunSym, 0)
	c.emit(OP_LOAD_CONST, symConst, 0)
	c.emit(OP_LOAD_CONST, fnConst, 0)
	c.emit(OP_CALL, 2, 0)
}

func (c *klCompiler) compileLambda(params []Obj, body Obj) {
	inner := newCompiler("lambda", len(params), params, c)
	inner.compileExpr(body, true)

	// The inner compiler may have discovered upvalues; emit loads for them.
	for _, uv := range inner.upvals {
		switch uv.outerRef.kind {
		case varLocal:
			c.emit(OP_LOAD_LOCAL, int32(uv.outerRef.index), 0)
		case varUpval:
			c.emit(OP_LOAD_UPVAL, int32(uv.outerRef.index), 0)
		}
	}

	innerObj := makeBytecodeObj(inner.fn, nil) // upvals populated at runtime
	innerConst := c.addConst(innerObj)
	c.emit(OP_MAKE_CLOSURE, innerConst, int32(len(inner.upvals)))
}

func (c *klCompiler) compileLet(x, val, body Obj, tail bool) {
	// Evaluate val.
	c.compileExpr(val, false)
	// Assign to a new local slot (or reuse if x is already a local).
	// We always allocate a new slot to handle shadowing correctly.
	oldIdx, hadOld := c.locals[x]
	slot := c.newLocal(x)
	c.emit(OP_STORE_LOCAL, slot, 0)
	// Compile body.
	c.compileExpr(body, tail)
	// Restore previous mapping (for shadowing).
	if hadOld {
		c.locals[x] = oldIdx
	} else {
		delete(c.locals, x)
	}
}

func (c *klCompiler) compileIf(cond, thenExpr, elseExpr Obj, tail bool) {
	c.compileExpr(cond, false)
	// Emit JUMP_FALSE with placeholder; patch after then-branch.
	jumpFalseIdx := c.emit(OP_JUMP_FALSE, 0, 0)
	c.compileExpr(thenExpr, tail)
	if !tail {
		// Emit JUMP over else-branch; patch after else.
		jumpOverIdx := c.emit(OP_JUMP, 0, 0)
		c.patchJump(jumpFalseIdx)
		c.compileExpr(elseExpr, tail)
		c.patchJump(jumpOverIdx)
	} else {
		// In tail position both branches already emit RETURN; just patch the false jump.
		c.patchJump(jumpFalseIdx)
		c.compileExpr(elseExpr, tail)
	}
}

func (c *klCompiler) compileCond(clauses Obj, tail bool) {
	if *clauses == scmHeadNull {
		// No matching clause: return Nil.
		c.emit(OP_LOAD_CONST, c.addConst(Nil), 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return
	}
	clause := car(clauses)
	rest := cdr(clauses)
	cond := car(clause)
	action := cadr(clause)

	if cond == True {
		// (cond (true action) ...) — unconditional
		c.compileExpr(action, tail)
		return
	}

	c.compileExpr(cond, false)
	jumpFalseIdx := c.emit(OP_JUMP_FALSE, 0, 0)
	c.compileExpr(action, tail)
	if !tail {
		jumpOverIdx := c.emit(OP_JUMP, 0, 0)
		c.patchJump(jumpFalseIdx)
		c.compileCond(rest, tail)
		c.patchJump(jumpOverIdx)
	} else {
		c.patchJump(jumpFalseIdx)
		c.compileCond(rest, tail)
	}
}

// compileTrapError lowers (trap-error body handler) to (try-catch (freeze body) handler).
// Stack layout for CALL/TAIL_CALL 2: [..., fn, arg1, arg2]
func (c *klCompiler) compileTrapError(body, handler Obj, tail bool) {
	// 1. Push try-catch global function.
	tryCatchSym := c.addConst(MakeSymbol("try-catch"))
	c.emit(OP_LOAD_GLOBAL, tryCatchSym, 0)
	// 2. Push (freeze body) as arg1.
	c.compileLambda(nil, body)
	// 3. Push handler as arg2.
	c.compileExpr(handler, false)
	// 4. Call try-catch with 2 args.
	// trap-error must NOT be a tail call: the panic/recover in try-catch
	// needs a real Go stack frame to catch from.
	c.emit(OP_CALL, 2, 0)
	if tail {
		c.emit(OP_RETURN, 0, 0)
	}
}

// intrinsicOp maps a 2-arg primitive symbol to a fast-path opcode.
// Returns 0 if no fast path exists.
func intrinsicOp2(sym Obj) uint8 {
	switch sym {
	case symAdd:
		return OP_ADD
	case symSub:
		return OP_SUB
	case symMul:
		return OP_MUL
	case symLT:
		return OP_LT
	case symLE:
		return OP_LE
	case symGT:
		return OP_GT
	case symGE:
		return OP_GE
	case symNumEq:
		return OP_EQ
	}
	return 0
}

func (c *klCompiler) compileCall(fn Obj, args Obj, tail bool) {
	nArgs := 0
	argList := make([]Obj, 0, 4)
	for l := args; *l == scmHeadPair; l = cdr(l) {
		argList = append(argList, car(l))
		nArgs++
	}

	// ---- 1-arg intrinsics ----
	if IsSymbol(fn) && nArgs == 1 && fn == symNot {
		c.compileExpr(argList[0], false)
		c.emit(OP_NOT, 0, 0)
		if tail {
			c.emit(OP_RETURN, 0, 0)
		}
		return
	}

	// ---- 2-arg arithmetic intrinsics ----
	if IsSymbol(fn) && nArgs == 2 {
		if op := intrinsicOp2(fn); op != 0 {
			c.compileExpr(argList[0], false)
			c.compileExpr(argList[1], false)
			c.emit(op, 0, 0)
			if tail {
				c.emit(OP_RETURN, 0, 0)
			}
			return
		}
	}

	// ---- Self-tail-call optimization ----
	if tail && IsSymbol(fn) && mustSymbol(fn).str == c.fn.Name {
		kind, _ := c.resolveVar(fn)
		if kind == varGlobal && nArgs == c.fn.Arity {
			// It's a self-call in tail position: emit args then OP_SELF_TAIL_CALL.
			for _, a := range argList {
				c.compileExpr(a, false)
			}
			c.emit(OP_SELF_TAIL_CALL, int32(nArgs), 0)
			return
		}
	}

	// ---- General call ----
	if IsSymbol(fn) {
		kind, idx := c.resolveVar(fn)
		switch kind {
		case varLocal:
			c.emit(OP_LOAD_LOCAL, int32(idx), 0)
		case varUpval:
			c.emit(OP_LOAD_UPVAL, int32(idx), 0)
		case varGlobal:
			c.emit(OP_LOAD_GLOBAL, c.addConst(fn), 0)
		}
	} else {
		c.compileExpr(fn, false)
	}

	for _, a := range argList {
		c.compileExpr(a, false)
	}

	if tail {
		c.emit(OP_TAIL_CALL, int32(nArgs), 0)
	} else {
		c.emit(OP_CALL, int32(nArgs), 0)
	}
}
