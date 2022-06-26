#include <stdarg.h>
#include <stddef.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>
#include <stdio.h>
#include "types.h"

const Obj True = ((1 << (TAG_SHIFT+1)) | TAG_BOOLEAN);
const Obj False = ((2 << (TAG_SHIFT+1)) | TAG_BOOLEAN);
const Obj Nil = ((666 << (TAG_SHIFT+1)) | TAG_IMMEDIATE_CONST);
const Obj Undef = ((42 << TAG_SHIFT) | TAG_IMMEDIATE_CONST);

struct GC;

uintptr_t gcCopy(struct GC *gc, uintptr_t head) { return 42; };

struct scmString {
  scmHead head;
  int sz;
  char data[];
};

static void*
newObj(scmHeadType tp, int sz) {
  scmHead* p = malloc(sz);
  /* scmHead* p = gcAlloc(&gc, sz); */
  assert(((Obj)p & TAG_PTR) == 0);
  p->type = tp;
  return (void*)p;
}

Obj
makeCons(Obj car, Obj cdr) {
  struct scmCons* p = newObj(scmHeadCons, sizeof(struct scmCons));
  p->car = car;
  p->cdr = cdr;
  return ((Obj)(&p->head) | TAG_CONS);
}

static void
consGCFunc(struct GC *gc, void* f, void* t) {
  struct scmCons *from = f;
  struct scmCons *to = t;
  to->car = gcCopy(gc, from->car);
  to->cdr = gcCopy(gc, from->cdr);
}

Obj
consp(Obj x) {
  if (iscons(x)) {
    return True;
  }
  return False;
}

Obj
cadr(Obj x) {
  return car(cdr(x));
}

Obj
caddr(Obj x) {
  return car(cdr(cdr(x)));
}

Obj
cdddr(Obj x) {
  return cdr(cdr(cdr(x)));
}

struct scmClosure {
  scmHead head;
  int required;
  Instr code;
  Obj* slot;
};

Obj
makeClosure(int required, Instr code, Obj* closed) {
  struct scmClosure* clo = newObj(scmHeadClosure, sizeof(struct scmClosure));
  clo->required = required;
  clo->code = code;
  return ((Obj)(&clo->head) | TAG_PTR);
}

Instr
closureCode(Obj o) {
  struct scmClosure* c = ptr(o);
  assert(c->head.type == scmHeadClosure);
  return c->code;
}

struct scmContinuation {
  scmHead head;
  struct stack s;
  Instr code;
};

Obj
makeContinuation(struct stack s, Instr code) {
  struct scmContinuation* cont = newObj(scmHeadContinuation, sizeof(struct scmContinuation));
  cont->s = s;
  cont->code = code;
  return ((Obj)(&cont->head) | TAG_PTR);
}

struct stack
contStack(Obj o) {
  struct scmContinuation* cont = ptr(o);
  assert(cont->head.type == scmHeadContinuation);
  return cont->s;
}

Instr
contCode(Obj o) {
  struct scmContinuation* cont = ptr(o);
  assert(cont->head.type == scmHeadContinuation);
  return cont->code;
}

Obj
makeNumber(int v) {
  if (v < 99999999) {
    // The type of a fixnum is actually intptr_t, although stored as uintptr.
    // Be careful with the sign digit.
    Obj res = (Obj)(((intptr_t)(v)<<1));
    assert((res & 1) == 0);
    return res;
  }
  // TODO
  return (Obj)(99999999);
}

bool
isNumber(Obj o) {
  if (tag(o) == TAG_FIXNUM) {
    return true;
  }
  if (tag(o) == TAG_PTR) {
    if (((scmHead*)ptr(o))->type == scmHeadNumber) {
      return true;
    }
  }
  return false;
}

Obj
makeString(const char *s, int len) {
  // sz is the actural length but we malloc a extra byte to be compatible with C.
  int alloc = len + sizeof(struct scmString) + 1;
  struct scmString* str = newObj(scmHeadString, alloc);
  str->sz = len;
  str->data[len] = '\0';
  if (s != NULL) {
    memcpy(&str->data[0], s, len);
  }
  return ((Obj)(&str->head) | TAG_PTR);
}

char*
stringStr(Obj o) {
  struct scmString* s = ptr(o);
  assert(s->head.type == scmHeadString);
  return s->data;
}

int
stringLen(Obj o) {
  struct scmString* s = ptr(o);
  assert(s->head.type == scmHeadString);
  return s->sz;
}

bool
isstring(Obj o) {
  if (tag(o) == TAG_PTR) {
    if (((scmHead*)ptr(o))->type == scmHeadString) {
      return true;
    }
  }
  return false;
}

static void
stringGCFunc(struct GC *gc, void* f, void* t) {
  // TODO:
}

struct trieNode {
  Obj value;
  char *sym;
  struct trieNode* child[256];
};

struct trieNode root = {};

/* static void */
/* trieNodeGCFunc(struct GC* gc, struct trieNode *node) { */
/*   node->value = gcCopy(gc, node->value); */
/*   for (int i=0; i<256; i++) { */
/*     if (node->child[i] != NULL) { */
/*       trieNodeGCFunc(gc, node->child[i]); */
/*     } */
/*   } */
/* } */

/* void */
/* gcGlobal(struct GC *gc) { */
/*   trieNodeGCFunc(gc, &root); */
/* } */

Obj
makeSymbol(char *s) {
  char *old = s;
  struct trieNode* p = &root;
  for(; *s; s++) {
    int offset = *s;
    if (p->child[offset] == NULL) {
      struct trieNode *n = malloc(sizeof(struct trieNode));
      memset(n, 0, sizeof(struct trieNode));
      p->child[offset] = n;
    }
    p = p->child[offset];
  }
  if (p->sym == NULL) {
    char* tmp = malloc(strlen(old) + 1);
    strcpy(tmp, old);
    p->sym = tmp;
    p->value = Undef;
  }

  return (Obj)(p) | TAG_SYMBOL;
}

Obj
symbolGet(Obj sym) {
  assert(issymbol(sym));
  struct trieNode* s = ptr(sym);
  return s->value;
}

Obj
symbolSet(Obj sym, Obj val) {
  assert(issymbol(sym));
  struct trieNode* s = ptr(sym);
  s->value = val;
  return val;
}

const char*
symbolStr(Obj sym) {
  assert(issymbol(sym));
  struct trieNode *s = ptr(sym);
  return s->sym;
}

struct scmCurry {
  scmHead head;
  int required;
  int captured;
  // The first element is scmNative and the remain is arguments.
  Obj data[];
};

Obj
makeCurry(int required, int captured) {
  int sz = sizeof(struct scmCurry) + captured*sizeof(Obj);
  struct scmCurry* clo = newObj(scmHeadCurry, sz);
  clo->required = required;
  clo->captured = captured;
  assert(captured > 0);
  return ((Obj)(&clo->head) | TAG_PTR);
}

static void
curryGCFunc(struct GC *gc, void* f, void* t) {
  struct scmCurry *from = f;
  struct scmCurry *to = t;
  for (int i=0; i<from->captured; i++) {
    to->data[i] = gcCopy(gc, from->data[i]);
  }
}

int
curryRequired(Obj o) {
  struct scmCurry *curry = ptr(o);
  return curry->required;
}

Obj
curryCaptured(Obj o) {
  struct scmCurry *curry = ptr(o);
  return curry->captured;
}

Obj*
curryData(Obj o) {
  struct scmCurry *curry = ptr(o);
  return curry->data;
}

struct scmNative {
  scmHead head;
  nativeFuncPtr fn;
  // required is the argument number of the nativeFunc.
  int required;
  // captured is the size of the data, it's immutable after makeNative.
  int captured;
  Obj data[];
};

Obj
makeNative(nativeFuncPtr fn, int required, int captured, ...) {
  int sz = sizeof(struct scmNative) + captured*sizeof(Obj);
  struct scmNative* clo = newObj(scmHeadNative, sz);
  clo->fn = fn;
  clo->required = required;
  clo->captured = captured;

  if (captured > 0) {
    va_list ap;
    va_start(ap, captured);
    for (int i=0; i<captured; i++) {
      clo->data[i] = va_arg(ap, Obj);
    }
    va_end(ap);
  }

  return ((Obj)(&clo->head) | TAG_PTR);
}

static void
nativeGCFunc(struct GC *gc, void* f, void* t) {
  struct scmNative *from = f;
  struct scmNative *to = t;
  for (int i=0; i<from->captured; i++) {
    to->data[i] = gcCopy(gc, from->data[i]);
  }
}

nativeFuncPtr
nativeFn(Obj o) {
  struct scmNative* native = ptr(o);
  assert(native->head.type == scmHeadNative);
  return native->fn;
}

int
nativeRequired(Obj o) {
  struct scmNative* native = ptr(o);
  assert(native->head.type == scmHeadNative);
  return native->required;
}

int
nativeCaptured(Obj o) {
  struct scmNative* native = ptr(o);
  assert(native->head.type == scmHeadNative);
  return native->captured;
}

Obj
nativeRef(Obj o, int idx) {
  struct scmNative* native = ptr(o);
  assert(native->head.type == scmHeadNative);
  return native->data[idx];
}

Obj
makeBuiltin(nativeFuncPtr fn, int required) {
  return makeNative(fn, required+1, 0);
}

struct scmVector {
  scmHead head;
  int size;
  Obj data[];
};

Obj
makeVector(int size) {
  struct scmVector* vec = newObj(scmHeadVector, sizeof(struct scmVector)+sizeof(Obj)*size);
  vec->size = size;
  return ((Obj)(&vec->head) | TAG_PTR);
}

Obj
vectorRef(Obj v, int idx) {
  struct scmVector* vec = ptr(v);
  assert(vec->head.type == scmHeadVector);
  assert(idx >= 0 && idx < vec->size);
  return vec->data[idx];
}

Obj
vectorSet(Obj vec, int idx, Obj val) {
  struct scmVector* v = ptr(vec);
  assert(v->head.type == scmHeadVector);
  assert(idx >= 0 && idx < v->size);
  v->data[idx] = val;
  return vec;
}

bool
isvector(Obj o) {
  if (tag(o) == TAG_PTR) {
    if (((scmHead*)ptr(o))->type == scmHeadVector) {
      return true;
    }
  }
  return false;
}

static void
vectorGCFunc(struct GC *gc, void* f, void* t) {
  struct scmVector *from = f;
  struct scmVector *to = t;
  for (int i=0; i<from->size; i++) {
    to->data[i] = gcCopy(gc, from->data[i]);
  }
}

bool
eq(Obj x, Obj y) {
  if (x == y) {
    return true;
  }

  if (iscons(x) && iscons(y)) {
    if (!eq(car(x), car(y))) {
      return false;
    }
    return eq(cdr(x), cdr(y));
  }

  if (isstring(x) && isstring(y)) {
    struct scmString* x1 = ptr(x);
    struct scmString* y1 = ptr(y);
    if (x1->sz != y1->sz) {
      return false;
    }
    for (int i=0; i<x1->sz; i++) {
      if (x1->data[i] != y1->data[i]) {
	return false;
      }
    }
    return true;
  }

  return false;
}

void
typesInit() {
  /* gcRegistForType(scmHeadCons, consGCFunc); */
  /* gcRegistForType(scmHeadNative, nativeGCFunc); */
  /* gcRegistForType(scmHeadCurry, curryGCFunc); */
  /* gcRegistForType(scmHeadString, stringGCFunc); */
  /* gcRegistForType(scmHeadVector, vectorGCFunc); */
}
