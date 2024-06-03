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

// 4681            268921 ns/op           24102 B/op        801 allocs/op
func BenchmarkConcurrentCache(b *testing.B) {
	c := cache.NewConcurrentCache()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, 100)
	}
}

// 6073            187311 ns/op           24050 B/op        801 allocs/op
func BenchmarkShardedCache(b *testing.B) {
	c := cache.NewShardedCache()

	for i := 0; i < b.N; i++ {
		emulateLoad(c, 100)
	}
}
