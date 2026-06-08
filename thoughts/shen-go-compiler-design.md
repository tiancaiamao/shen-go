# shen-go Compiler Design

**Date**: 2026-05-05  
**Branch**: `claude/shen-go-compiler-Stnpj`  
**Goal**: Close the 60× interpreter gap to within 3–5× of shen-cl on bench/bench.shen

---

## Problem Statement

The current shen-go evaluates user-defined Shen functions entirely through a
KL tree-walking interpreter.  The profiler shows roughly 85% of CPU time in
interpreter plumbing (`eval`, `trampoline`, `apply`) and ~25% in Go GC.
Arithmetic and other "useful" work is sub-1%.

The three concrete bottlenecks are:

1. **`envGet` (O(depth) alist scan)** – every variable lookup walks a
   linked list of `(sym . val)` cons cells.
2. **`envExtend` (heap allocation)** – every function call allocates
   `2 × arity` cons cells to extend the environment.
3. **`eval` tree dispatch** – every node requires a tag switch and a
   9-way special-form dispatch, all resolvable at compile time.

A secondary correctness bug: `<`, `<=`, `>`, `>=` called `mustInteger`
(which truncates floats) instead of `mustNumber`.  Fixed in Phase 0.

---

## Architecture Choice: Bytecode VM (Option A)

**Chosen**: stack-based bytecode VM with per-call flat local-variable frames.

**Rejected alternatives**:
- *AOT-to-Go plugin*: Go plugin ABI is fragile; per-defun compile latency
  is hundreds of ms; no Windows; no clean redefinition.
- *CPS/direct-threaded*: harder to debug, limited advantage over a switch VM
  on Go (no tail-call guarantee in Go).
- *LLVM/native*: too heavy; no JIT in the Go toolchain.

**Why the bytecode VM**:
- Each `defun` is compiled once at define time to a flat `[]Instr` sequence.
- Variables become integer-indexed slots in a per-call `[]Obj` frame.
  `envGet` (O(depth)) → array read O(1).  `envExtend` cons allocation →
  eliminated entirely for the common case.
- Special-form dispatch happens at compile time; the VM loop is a single
  `switch` over `uint8` opcodes.
- Closures (lambda/freeze) carry `upvals []Obj` — a snapshot of captured
  values at creation time.  No alist chain.
- TCO: tail calls set `ControlFlowApply` and return to the existing
  trampoline.  Self-tail calls are trampolined without growing the Go stack.
- `trap-error` is lowered to `(try-catch (freeze body) handler)` at compile
  time, reusing the existing panic/recover primitive.

---

## Instruction Set

```
OP_LOAD_CONST  A   — push consts[A]
OP_LOAD_LOCAL  A   — push locals[A]
OP_STORE_LOCAL A   — locals[A] = pop()
OP_LOAD_GLOBAL A   — push mustSymbol(consts[A]).function  (runtime lookup)
OP_LOAD_UPVAL  A   — push upvals[A]
OP_CALL        A   — non-tail call: pop fn + A args, push result
OP_TAIL_CALL   A   — tail call: set ControlFlowApply, return to trampoline
OP_RETURN          — ctl.Return(pop())
OP_JUMP        A   — pc += A  (signed)
OP_JUMP_FALSE  A   — if pop()==False: pc += A
OP_MAKE_CLOSURE A B — consts[A] is a *BytecodeFunc; pop B upvals from stack;
                       push closure Obj
OP_POP             — discard top of stack
```

`LOAD_GLOBAL` resolves the symbol's `.function` at call time so that
redefined functions are always picked up without re-compilation.

---

## Compilation Rules

| KL form | Compiled to |
|---|---|
| number/string/bool/nil | `OP_LOAD_CONST` |
| local symbol | `OP_LOAD_LOCAL idx` |
| upvalue symbol | `OP_LOAD_UPVAL idx` |
| global symbol (value pos) | `OP_LOAD_CONST sym` (returns symbol itself) |
| `(defun f (x…) body)` | compile body; bind `scmBytecodeFunc` to `f.function` |
| `(lambda x body)` | inner compile; `OP_MAKE_CLOSURE` with captured upvals |
| `(freeze body)` | same as lambda with 0 params |
| `(let x val body)` | compile val, `OP_STORE_LOCAL`, compile body |
| `(if a b c)` | compile a, `OP_JUMP_FALSE`, compile b, `OP_JUMP`, compile c |
| `(and a b)` | → `(if a b false)` |
| `(or a b)` | → `(if a true b)` |
| `(cond …)` | → nested ifs |
| `(do a b)` | compile a (non-tail), `OP_POP`, compile b (inherits tail) |
| `(type x _)` | compile x |
| `(trap-error b h)` | → `(try-catch (freeze b) h)` |
| `(f arg…)` | compile f + args, `OP_CALL`/`OP_TAIL_CALL` |

---

## Integration with Existing Code

- `eval()` special-form handler for `defun` calls `CompileFunc` and stores
  the resulting `scmBytecodeFunc` in the symbol's `.function` slot.
- `apply()` gains a `scmHeadBytecodeFunc` case that calls `vmExec`.
- The existing `scmProcedure` / tree-walker path remains as a fallback for
  any form that fails compilation (should not happen in practice).
- `primDefun` (called by Shen-level compiler) also compiles its argument if
  it receives a `scmProcedure`.
- `ObjString`, `equal`, and the `apply` partial-application logic all handle
  `scmBytecodeFunc` alongside the existing types.

---

## Upvalue Capture

Upvalues are discovered during compilation of a nested lambda/freeze.
When the inner compiler cannot resolve a symbol as a local or already-known
upvalue, it queries the outer compiler's `resolveVar`.  The outer compiler
may itself recurse upward for multi-level closures.

At closure creation (in the outer function's bytecode), the outer compiler
emits `OP_LOAD_LOCAL` or `OP_LOAD_UPVAL` for each captured variable before
`OP_MAKE_CLOSURE`.  The inner function's bytecode accesses them via
`OP_LOAD_UPVAL`.

---

## Phased Plan (remaining milestones)

| Phase | Description | Expected gain |
|---|---|---|
| 0 | Float fix; design doc; bench baseline | correctness |
| 2 | Bytecode VM (params → slots; compile defun/lambda/let) | 4–8× |
| 3 | Numeric specialization (fixnum fast paths) | 2–3× on tak/fib |
| 4 | Pattern-match compilation to decision trees | 1.5–2× on nrev/map |
| 5 | Inline cache for global lookup; self-tail jump | 1.2–1.5× |

---

## Things Tried and Reverted

*(fill in as work proceeds)*

- Phase 1 (annotated-AST indexed slots) was folded directly into Phase 2.
  Keeping the tree-walker with slot annotations would have required threading
  a `frame []Obj` through all eval/apply paths — essentially a partial VM —
  so it was cleaner to build the full VM directly.
