#include "types.h"
#include "vm.h"
#include <assert.h>
#include <unistd.h>
#include <stdio.h>

void entry(struct VM *vm);
void fn_xx(struct VM *vm);
void fn_yy(struct VM *vm);


void primAdd(struct VM* vm);
void primID(struct VM* vm);

int main(int argc, char *argv[]) {
  struct VM *vm;
  vm = newVM();
  run(vm, entry);
  Obj ret = pop(vm);
  printf("result == %ld\n", ret);
  return 0;
}

void
entry(struct VM *vm) {
  push(vm, Nil);
  vm->val = makeClosure(1, fn_xx, NULL);
  push(vm, vm->val);
  vm->val = True;
  push(vm, vm->val);
  instrCall(vm, 2, fn_yy);
}


void fn_xx(struct VM *vm) {
  vm->val = vmGet(vm, 2);
  if (vm->val == True) {
    vm->val = makeNumber(4);
    push(vm, vm->val);
    vm->val = makeNumber(7);
    push(vm, vm->val);
    primAdd(vm);
    primID(vm);
    vmReturn(vm, vm->val);
  } else if (vm->val == False) {
    vm->val = makeNumber(42);
    primID(vm);
    vmReturn(vm, vm->val);
  } else {
    perror("if only accept true or false");
  }
}

void fn_yy(struct VM *vm) {
  primID(vm);
  vmReturn(vm, vm->val);
}

void
primAdd(struct VM* vm) {
  Obj x = pop(vm);
  Obj y = pop(vm);
  Obj z = makeNumber(fixnum(x) + fixnum(y));
  vm->val = z;
}

void
primID(struct VM* vm) {
}
