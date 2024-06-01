package bench

import (
	"testing"

	"github.com/alexeysamorodov/cache/internal/app/cache"
)

const parallelFactor = 10_000

func BenchmarkSimpleCache(b *testing.B) {
	b.Skip("panic in SimpleCache")

	c := cache.NewSimpleCache()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, 1)
	}
}

func BenchmarkConcurrentCache(b *testing.B) {
	c := cache.NewConcurrentCache()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, 1)
	}
}
