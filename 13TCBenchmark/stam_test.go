package main

import "testing"

// name the file whatece_test.go
// import testing
// write the benchmarks
// benchmark function name must start with Benchmark and after it capital letter anything
// parameter is of type *testing.B
// go test -bench='.'
// go test -bench='.*01'
func BenchmarkA01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("kuku")
	}
}

func BenchmarkA02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoSum(i, i)
	}
}
