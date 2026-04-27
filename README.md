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

or for windows

```
make shen-exe
```

## Running

```
./shen
```

This binary has no dependency, you can move it to any where you want.

## Testing

1. Run unit test:

```
make test
```

2. Test the `klambda` implementation:

```
make kl
cd kernel/tests
kl
(load-file "../../cmd/kl/runtests.kl")
(load "runme.shen")
```

3. Test the `shen` binary:


```
make shen
cd kernel/tests
../../shen
(load "runme.shen")
```

## How to bootstrap

You can just do
```
cd compiled
kl
(load-file "script.kl")
cd ..
make shen
``` 
Explanation : 

`kl` implement a simple klambda interpreter in Go, which can be used to bootstrap `shen`

```
(load-file "../kernel/klambda/toplevel.kl")
(load-file "../kernel/klambda/core.kl")
(load-file "../kernel/klambda/sys.kl")
(load-file "../kernel/klambda/sequent.kl")
(load-file "../kernel/klambda/yacc.kl")
(load-file "../kernel/klambda/reader.kl")
(load-file "../kernel/klambda/prolog.kl")
(load-file "../kernel/klambda/track.kl")
(load-file "../kernel/klambda/load.kl")
(load-file "../kernel/klambda/writer.kl")
(load-file "../kernel/klambda/macros.kl")
(load-file "../kernel/klambda/declarations.kl")
(load-file "../kernel/klambda/t-star.kl")
(load-file "../kernel/klambda/types.kl")
(load-file "../kernel/klambda/dict.kl")
(load-file "../kernel/klambda/init.kl")
(shen.initialise)
```

`shen` source files is generated from the `.kl` files. The full transformation path is Shen -> KL -> IR -> Go.

The file `src/compiler.shen` is a transpiler from KL to an intermediate representation(IR), load it:

```
(load "../src/compiler.shen")
```

Compile the klambda to the intermediate representation:

```
(set *maximum-print-sequence-size* 100000)
(compile-file "../kernel/klambda/sys.kl" "sys.tmp")
(compile-file "../kernel/klambda/writer.kl" "writer.tmp")
(compile-file "../kernel/klambda/core.kl" "core.tmp")
(compile-file "../kernel/klambda/reader.kl" "reader.tmp")
(compile-file "../kernel/klambda/declarations.kl" "declarations.tmp")
(compile-file "../kernel/klambda/toplevel.kl" "toplevel.tmp")
(compile-file "../kernel/klambda/macros.kl" "macros.tmp")
(compile-file "../kernel/klambda/load.kl" "load.tmp")
(compile-file "../kernel/klambda/prolog.kl" "prolog.tmp")
(compile-file "../kernel/klambda/sequent.kl" "sequent.tmp")
(compile-file "../kernel/klambda/track.kl" "track.tmp")
(compile-file "../kernel/klambda/t-star.kl" "t-star.tmp")
(compile-file "../kernel/klambda/yacc.kl" "yacc.tmp")
(compile-file "../kernel/klambda/types.kl" "types.tmp")
(compile-file "../kernel/klambda/dict.kl" "dict.tmp")
(compile-file "../kernel/klambda/init.kl" "init.tmp")
```

And generate the Go files from the intermediate representation:

Use `compiled/bctogo.shen` to generate the Go files from the intermediate representation.

Now the shen source files are available, built it:

```
make shen
```


## Learn Shen
* [Official website of Shen](http://shenlanguage.org/)
* [Shen Community Wiki](https://github.com/Shen-Language/wiki/wiki)

## License

- Shen, Copyright © 2010-2015 Mark Tarver - [License](http://www.shenlanguage.org/license.pdf).
- shen-go, Copyright © 2017-2022 Arthur Mao under [BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause).
