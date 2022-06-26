#ifndef _TYPES_H_
#define _TYPES_H_

#include <stdint.h>
#include <stdbool.h>
/* #include "gc.h" */


typedef uint8_t scmHeadType;

typedef struct {
  scmHeadType type;
  uint16_t size;
  uintptr_t forwarding;
} scmHead;

#define TAG_SHIFT 3
#define TAG_MASK 0x7
#define TAG_PTR 0x7

#define ptr(x) ((void*)((x)&~TAG_PTR))
#define tag(x) ((x) & TAG_MASK)

typedef uintptr_t Obj;

// 000 fixnum
// 001 non-fixnum
// 001 symbol
// 011 cons
// 101 immediate const (boolean, null, undef...)
// 111 general pointer (string, vector, number, error...)
#define TAG_FIXNUM 0x0
#define TAG_SYMBOL 0x1
#define TAG_CONS 0x3
#define TAG_IMMEDIATE_CONST 0x5
// 1-101 boolean
// 0-101 other constant
#define TAG_BOOLEAN 0xd

#define iscons(x) (tag(x) == TAG_CONS)
#define issymbol(x) (tag(x) == TAG_SYMBOL)
#define isfixnum(x) (((x) & 1) == 0)
#define isboolean(x) (((x) & 0xf) == TAG_BOOLEAN)

extern const Obj True;
extern const Obj False;
extern const Obj Nil;
extern const Obj Undef;

struct controlFlow;

enum {
      // Instant values.
      scmHeadBoolean,
      scmHeadNull,
      // Symbol is a special pointer, but it's basically instant value.
      // Number may be or may not be pointer.
      scmHeadNumber,
      // The followings are all pointer types.
      scmHeadCons,
      scmHeadNative,
      scmHeadCurry,
      scmHeadString,
      scmHeadVector,
      scmHeadClosure,
      scmHeadContinuation,
};

void typesInit();

struct VM;
typedef void (*Instr)(struct VM *vm);
typedef void (*nativeFuncPtr) (struct controlFlow *ctx);

struct scmCons {
  scmHead head;
  Obj car;
  Obj cdr;
};

#define fixnum(x) ((intptr_t)(x)>>1)
bool eq(Obj x, Obj y);

#define intern(x) makeSymbol(x)
Obj makeSymbol(char *s);
Obj symbolGet(Obj sym);
Obj symbolSet(Obj sym, Obj val);
const char* symbolStr(Obj sym);

#define cons(x, y) makeCons(x, y)
Obj makeCons(Obj car, Obj cdr);
Obj consp(Obj v);
Obj cadr(Obj v);
Obj caddr(Obj v);
Obj cdddr(Obj v);
#define car(v) (((struct scmCons*)(ptr(v)))->car)
#define cdr(v) (((struct scmCons*)(ptr(v)))->cdr)

Obj makeBuiltin(nativeFuncPtr fn, int required);
Obj makeNative(nativeFuncPtr fn, int required, int captured, ...);
Obj nativeRef(Obj o, int idx);
int nativeRequired(Obj o);
int nativeCaptured(Obj o);
nativeFuncPtr nativeFn(Obj o);

Obj makeCurry(int required, int captured);
int curryRequired(Obj curry);
Obj curryCaptured(Obj curry);
Obj* curryData(Obj curry);

Obj makeClosure(int required, Instr code, Obj* closed);
Instr closureCode(Obj);

Obj makeString(const char *s, int len);
char* stringStr(Obj o);
int stringLen(Obj o);
Obj makeNumber(int v);
bool isstring(Obj o);
bool isNumber(Obj o);

struct stack {
  Obj *data;
  int base;
  int pos;
};

Obj makeContinuation(struct stack s, Instr code);
Instr contCode(Obj o);

Obj symQuote, symIf, symLambda, symDo, symMacroExpand, symDebugEval;
/* void gcSymbols(struct GC *gc); */

Obj makeVector(int c);
Obj vectorRef(Obj vec, int idx);
Obj vectorSet(Obj vec, int idx, Obj val);
bool isvector(Obj o);

struct stack contStack(Obj o);


#endif
