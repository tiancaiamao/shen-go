#include "types.h"
#include "vm.h"
#include <stdlib.h>
#include <stdio.h>

struct VM*
newVM() {
  struct VM* vm = (struct VM*)malloc(sizeof(struct VM));
  vm->pc = NULL;
  vm->data = (Obj*)malloc(sizeof(Obj)*2048);
  vm->base = 0;
  vm->pos = 0;
  printf("addr of vm is %p\n", vm);
  return vm;
}

void
run(struct VM* vm, Instr code) {
  struct stack save;
  save.data = vm->data;
  save.base = vm->base;
  save.pos = vm->pos;
  Obj cont = makeContinuation(save, NULL);
  vm->base = vm->pos;
  push(vm, cont);

  vm->pc = code;
  while(vm->pc != NULL) {
    vm->pc(vm);
  }
}

void
nextInstr(struct VM *vm, Instr code) {
  vm->pc = code;
}

void
vmReturn(struct VM *vm, Obj x) {
  Obj cc = vmGet(vm, 0);
  struct stack s = contStack(cc);
  vm->data = s.data;
  vm->base = s.base;
  vm->pos = s.pos;

  push(vm, x);
  nextInstr(vm, contCode(cc));
}

void
push(struct VM *vm, Obj v) {
  vm->data[vm->pos++] = v;
}

Obj
pop(struct VM* vm) {
  vm->pos--;
  return vm->data[vm->pos];
}

Obj
vmGet(struct VM* vm, int idx) {
  if (idx >= 0) {
    return vm->data[vm->base + idx];
  } else {
    return vm->data[vm->pos + idx];
  }
}

void
vmSet(struct VM* vm, int idx, Obj v) {
  vm->data[vm->base + idx] = v;
}

void
vmResize(struct VM* vm, int sz) {
  vm->pos = vm->base + sz;
}

void
instrCall(struct VM *vm, int argc, Instr next) {
  Obj fn = vmGet(vm, -argc);
  Instr code = closureCode(fn);

  int newBase = vm->pos - argc - 1;
  struct stack save;
  save.data = vm->data;
  save.base = vm->base;
  save.pos = newBase;
  
  // Save the old continuation.
  Obj cont = makeContinuation(save, next);
  vm->base = newBase;
  vmSet(vm, 0, cont);

  // Change the PC register.
  nextInstr(vm, code);
}
