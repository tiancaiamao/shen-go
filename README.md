# Shen/Go, a Go port of the Shen language

Shen is a portable functional programming language by [Mark Tarver](http://marktarver.com) that offers

- pattern matching,
- λ calculus consistency,
- macros,
- optional lazy evaluation,
- static type checking,
- an integrated fully functional Prolog,
- and an inbuilt compiler-compiler.

shen-go is a port of the Shen language that runs on top of Go implementations.

## Building

Make sure you have [Go installed](https://golang.org/doc/install).

```
make shen
```

## Running

```
./shen
```

This binary has no dependency, you can move it to any where you want.

## Testing

```
cd ShenOSKernel-21.0/tests
../../shen
(load "README.shen")
(load "tests.shen")
```

## Native Calls

Native call is implemented through Go plugin.

First, write a plugin Go file like:

```Go
package main

import (
    "github.com/tiancaiamao/shen-go/runtime"
)

func hello(args ...runtime.Obj) runtime.Obj {
	return runtime.MakeString("hello " + runtime.ObjString(args[0]))
}

func Main() {
	runtime.RegistNativeCall("hello", 1, hello)
}
```

Compile it to a plugin file:

```
go build -o test.so -buildmode=plugin
```

Run in shen-go repl:

```
(load-plugin "test.so")
(native "hello" "world")
```

The `native` keyword tells the compiler to use the right calling convention, the first argument is the function name, which you regist in the Go plugin file using `runtime.RegistNativeCall`.

Notice that native call is not curry function, If you want to partial apply them, you have to wrap it:

```
(defun hello (Str) (native "hello" Str))
```

## Bootstrap from scratch

Run shen repl, then

```
(load "compiler/bootstrap.shen")
```

This would compile all necessary shen files into bytecode.

Move the `.bc` files to bytecode directory

```
mv ShenOSKernel-21.0/sources/*.bc bytecode/
mv compiler/*.bc bytecode/
```

And try it with

```
./shen -boot=.
```

If everything is ok, you can embedded those bytecode files into `asset.go`

```
make generate
```

And make shen again to get the new binary.

## Learn Shen
* [Official website of Shen](http://shenlanguage.org/)
* [The Shen OS Kernel Manual](http://shenlanguage.org/learn-shen/index.html)
* [The Official Shen Standard](http://www.shenlanguage.org/learn-shen/shendoc.htm)
* [Shen Community Wiki](https://github.com/Shen-Language/wiki/wiki)
* [The Book of Shen: third edition](https://www.amazon.co.uk/Book-Shen-Third-Mark-Tarver/dp/1784562130)

## License

- Shen, Copyright © 2010-2015 Mark Tarver - [License](http://www.shenlanguage.org/license.pdf).
- shen-go, Copyright © 2017-2018 Arthur Mao under [BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause).
