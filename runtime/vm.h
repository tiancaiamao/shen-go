#ifndef _VM_H
#define _VM_H

struct VM {
  Instr pc;
  Obj val;
  Obj *data;
  int base;
  int pos;
};

struct VM* newVM();

void push(struct VM *vm, Obj v);
Obj pop(struct VM* vm);
void instrCall(struct VM *vm, int argc, Instr next);
Obj vmGet(struct VM* vm, int idx);
void vmSet(struct VM* vm, int idx, Obj v);
void vmResize(struct VM* vm, int sz);
void nextInstr(struct VM *vm, Instr code);
void vmReturn(struct VM *vm, Obj x);
void run(struct VM* vm, Instr code);



#endif
