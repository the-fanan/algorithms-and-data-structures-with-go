package main

import (
	"testing"
)
var testVal int = 1000000

func BenchmarkSumNumbers1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumNumbers1(testVal)
	}
}

func BenchmarkSumNumbers2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumNumbers2(testVal)
	}
}