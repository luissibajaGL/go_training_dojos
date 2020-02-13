package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func TestFibonacci(t *testing.T) {
	data := []struct {
		n    uint64
		want uint64
	}{
		{0, 1}, {1, 2}, {2, 3}, {3, 5}, {5, 8}, {8, 13}, {13, 21},
	}

	for _, d := range data {
		if got := Fibonacci(d.n); got != d.want {
			t.Errorf("Not a Fibonacci sequence for N: %d, got: %d, wanted: %d", d.n, got, d.want)
		}
	}
}
