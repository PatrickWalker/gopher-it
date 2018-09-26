package main

import "testing"

func benchmarkFib(i int, b *testing.B) {
	// b.N is a value that gives
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

//couldn't we use test driven tests for this? we could but as we don't check output we wouldn't save much
func BenchmarkFib1(b *testing.B) { benchmarkFib(1, b) }

//test name is BencharkXxxx not TestXxxxx
func BenchmarkFib2(b *testing.B) { benchmarkFib(2, b) }

//its a different argument as well
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
