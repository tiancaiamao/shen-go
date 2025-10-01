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
cd 'S31/Test Programs'
kl
(load-file "../../cmd/kl/runtests.kl")
(load "runme.shen")
```

3. Test the `shen` binary:


```
make shen
cd 'S31/Test Programs'
../../shen
(load "runme.shen")
```

## How to bootstrap

`kl` implement a simple klambda interpreter in Go, which can be used to bootstrap `shen`

```
;; mkdir -p compiled
;; cd compiled
;; kl
(load-file "../S31/KLambda/toplevel.kl")
(load-file "../S31/KLambda/core.kl")
(load-file "../S31/KLambda/sys.kl")
(load-file "../S31/KLambda/sequent.kl")
(load-file "../S31/KLambda/yacc.kl")
(load-file "../S31/KLambda/reader.kl")
(load-file "../S31/KLambda/prolog.kl")
(load-file "../S31/KLambda/track.kl")
(load-file "../S31/KLambda/load.kl")
(load-file "../S31/KLambda/writer.kl")
(load-file "../S31/KLambda/macros.kl")
(load-file "../S31/KLambda/declarations.kl")
(load-file "../S31/KLambda/t-star.kl")
(load-file "../S31/KLambda/types.kl")
(shen.shen)
```

`shen` source files is generated from the `.kl` files. The full transformation path is Shen -> KL -> IR -> Go.

The file `src/compiler.shen` is a transpiler from KL to an intermediate representation(IR), load it:

```
(load "../src/compiler.shen")
```

Compile the klambda to the intermediate representation:

```
(set *maximum-print-sequence-size* 100000)
(compile-file "../S31/KLambda/sys.kl" "sys.tmp")
(compile-file "../S31/KLambda/writer.kl" "writer.tmp")
(compile-file "../S31/KLambda/core.kl" "core.tmp")
(compile-file "../S31/KLambda/reader.kl" "reader.tmp")
(compile-file "../S31/KLambda/declarations.kl" "declarations.tmp")
(compile-file "../S31/KLambda/toplevel.kl" "toplevel.tmp")
(compile-file "../S31/KLambda/macros.kl" "macros.tmp")
(compile-file "../S31/KLambda/load.kl"  "load.tmp")
(compile-file "../S31/KLambda/prolog.kl" "prolog.tmp")
(compile-file "../S31/KLambda/sequent.kl" "sequent.tmp")
(compile-file "../S31/KLambda/track.kl" "track.tmp")
(compile-file "../S31/KLambda/t-star.kl" "t-star.tmp")
(compile-file "../S31/KLambda/yacc.kl" "yacc.tmp")
(compile-file "../S31/KLambda/types.kl" "types.tmp")
```

And generate the Go files from the intermediate representation:

```
(put bc->go arity 5)
(let Cg (make-code-generator)
     (do
      (bc->go Cg "SysMain" false "sys.tmp" "../cmd/shen/sys.go")
      (bc->go Cg "WriterMain" false "writer.tmp" "../cmd/shen/writer.go")
      (bc->go Cg "CoreMain" false "core.tmp" "../cmd/shen/core.go")
      (bc->go Cg "ReaderMain" false "reader.tmp" "../cmd/shen/reader.go")
      (bc->go Cg "DeclarationsMain" false "declarations.tmp" "../cmd/shen/declarations.go")
      (bc->go Cg "TopLevelMain" false "toplevel.tmp" "../cmd/shen/toplevel.go")
      (bc->go Cg "MacrosMain" false "macros.tmp" "../cmd/shen/macros.go")
      (bc->go Cg "LoadMain" false "load.tmp" "../cmd/shen/load.go")
      (bc->go Cg "PrologMain" false "prolog.tmp" "../cmd/shen/prolog.go")
      (bc->go Cg "SequentMain" false "sequent.tmp" "../cmd/shen/sequent.go")
      (bc->go Cg "TrackMain" false "track.tmp" "../cmd/shen/track.go")
      (bc->go Cg "TStarMain" false "t-star.tmp" "../cmd/shen/t-star.go")
      (bc->go Cg "YaccMain" false "yacc.tmp" "../cmd/shen/yacc.go")
      (bc->go Cg "TypesMain" true "types.tmp" "../cmd/shen/types.go")))
```

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
