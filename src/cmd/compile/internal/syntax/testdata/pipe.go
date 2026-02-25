// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Syntax test for the |> pipe operator with generic slice operations.
// xs |> f(args...) desugars at parse time to f(xs, args...).

package main

import "fmt"

func filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, x := range s {
		if f(x) {
			r = append(r, x)
		}
	}
	return r
}

func fmap[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, x := range s {
		r[i] = f(x)
	}
	return r
}

func reduce[T, U any](s []T, init U, f func(U, T) U) U {
	acc := init
	for _, x := range s {
		acc = f(acc, x)
	}
	return acc
}

func main() {
	xs := []int{1, 2, 3, 4, 5}

	// filter: keep evens
	evens := xs |> filter(func(x int) bool { return x%2 == 0 })
	fmt.Println(evens) // → [2 4]

	// fmap: int → string (cross-type)
	strs := xs |> fmap(func(x int) string { return fmt.Sprintf("<%d>", x) })
	fmt.Println(strs) // → [<1> <2> <3> <4> <5>]

	// reduce: sum
	sum := xs |> reduce(0, func(acc, x int) int { return acc + x })
	fmt.Println(sum) // → 15

	// chained pipeline
	result := xs |>
		filter(func(x int) bool { return x%2 != 0 }) |>
		fmap(func(x int) int { return x * x }) |>
		reduce(0, func(acc, x int) int { return acc + x })
	fmt.Println(result) // sum of squares of odds: 1 + 9 + 25 = 35
}
