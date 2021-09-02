package main

import (
    "testing"
)

func BenchmarkFibo(b *testing.B) {
	position := 9
    for i := 0; i < b.N; i++ {
        fiboAtPosition(position)
    }
}

func BenchmarkFibo2(b *testing.B) {
	position := 9
	m := make(map[int]int)
    for i := 0; i < b.N; i++ {
        fiboAtPosition2(position, m)
    }
}
