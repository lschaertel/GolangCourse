package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestIteration(t *testing.T) {
	for i := 0; i < 50; i++ {
		if i == 50 {
			t.Errorf("i is 50 or bigger")
		}
	}

}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(math.MaxInt64)
	}
}