# Go Compiler — Functional Operations PoC

A fork of the Go compiler used as a learning sandbox to explore how the
compiler is structured by adding functional programming features from scratch.

## The `|>` pipe operator

This fork adds a `|>` pipe operator to Go. It lets you chain function calls
left-to-right instead of nesting them right-to-left.

```go
// without pipe
result := reduce(fmap(filter(xs, isEven), double), 0, add)

// with pipe
result := xs |> filter(isEven) |> fmap(double) |> reduce(0, add)
```

`xs |> f(args...)` is syntactic sugar for `f(xs, args...)`. The transformation
happens at parse time — the rest of the compiler sees ordinary function calls.

**Rules:**
- Left-associative: `a |> f |> g` → `g(f(a))`
- Lowest precedence: `a + b |> f(c)` → `f(a+b, c)`
- Works with any function, including generics and builtins

## Generic slice operations

The pipe operator pairs naturally with generic helper functions:

```go
func filter[T any](s []T, f func(T) bool) []T
func fmap[T, U any](s []T, f func(T) U) []U
func reduce[T, U any](s []T, init U, f func(U, T) U) U
```

```go
xs := []int{1, 2, 3, 4, 5}

// filter: keep evens
evens := xs |> filter(func(x int) bool { return x%2 == 0 })
// → [2 4]

// fmap: int → string (cross-type)
strs := xs |> fmap(func(x int) string { return fmt.Sprintf("<%d>", x) })
// → [<1> <2> <3> <4> <5>]

// reduce: sum
sum := xs |> reduce(0, func(acc, x int) int { return acc + x })
// → 15

// chained pipeline
result := xs |>
    filter(func(x int) bool { return x%2 != 0 }) |>
    fmap(func(x int) int { return x * x }) |>
    reduce(0, func(acc, x int) int { return acc + x })
// sum of squares of odds: 1 + 9 + 25 = 35
```

## Build

Requires an existing Go toolchain to bootstrap:

```bash
cd src && ./make.bash
```

Run a program with the modified compiler:

```bash
./bin/go run yourfile.go
```

To run the test case above just run:

```bash
./bin/go run src/cmd/compile/internal/syntax/testdata/pipe.go 
```

## Roadmap

| # | Feature | Compiler phases touched | Status |
|---|---------|------------------------|--------|
| 1 | `\|>` pipe operator | scanner, parser | done |
| 2 | `fn` lambda shorthand (`fn(x) x*2`) | scanner, parser | planned |
| 3 | `filter`, `fmap`, `reduce` as builtins | types2, typecheck, walk | planned |
| 4 | Custom IR node (`OPIPELINE` or similar) | ir, walk | planned |

---

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 4.0 Attribution license][cc4-by].*

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://go.dev/dl/.

After downloading a binary release, visit https://go.dev/doc/install
for installation instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://go.dev/doc/install/source
for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines at https://go.dev/doc/contribute.

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://go.dev/wiki/Questions for a list of
places to ask questions about the Go language.

[rf]: https://reneefrench.blogspot.com/
[cc4-by]: https://creativecommons.org/licenses/by/4.0/
