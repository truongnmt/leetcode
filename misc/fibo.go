package main

import "fmt"

func main() {
	position := 9
	// result := fiboAtPosition(position)

	m := make(map[int]int)
	result := fiboAtPosition2(position, m)
	fmt.Println(result)
}

func fiboAtPosition(p int) int {
	if p == 0 {
		return 0
	}
	if p == 1 {
		return 1
	}

	return fiboAtPosition(p - 2) + fiboAtPosition(p - 1)
}

// use recursion with memoization
//
// $ go test -run=Fibo -bench=. 
// goos: darwin
// goarch: amd64
// BenchmarkFibo-16          6157921               182 ns/op
// BenchmarkFibo2-16       157318396                7.60 ns/op
func fiboAtPosition2(p int, m map[int]int) int {
	if p == 0 {
		return 0
	}
	if p == 1 {
		return 1
	}

	if m[p] != 0 {
		return m[p]
	}
	m[p] = fiboAtPosition(p - 2) + fiboAtPosition(p - 1)
	return m[p]
}
