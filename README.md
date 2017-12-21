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
cd ShenOSKernel-20.1/tests
../../shen
(load "README.shen")
(load "tests.shen")
```

## Bootstrap from scratch

Run shen repl, then

```
(load "compiler/primitive.shen")
```

This would compile all necessary shen files into bytecode.

Move the `.bc` files to bytecode directory

```
mv ShenOSKernel-20.1/sources/*.bc bytecode/
mv compiler/*.bc bytecode/
```

And try it with

```
./shen -boot=true
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
- shen-go, Copyright © 2017-2017 Arthur Mao under [BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause).
