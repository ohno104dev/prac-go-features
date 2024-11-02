package main

import (
	"math/rand"
	rand2 "math/rand/v2"
	"testing"
)

// go test ./rand_test.go -bench=Rand -count=1

const MAX = 1e9

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(MAX)
	}

}

// new version is faster
func BenchmarkRand2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand2.IntN(MAX)
	}
}
