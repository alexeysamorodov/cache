package cache

import (
	"context"
	"testing"
	"time"
)

func TestCacheWithContext_SetWithSmallTimeout(t *testing.T) {
	cache := NewCacheWithContext()

	timeout := time.Millisecond

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, timeout)

	err := cache.Set(ctx, "k", "v")
	if err != ErrorTimeout {
		t.Error("Expected timeout")
	}
}

func TestCacheWithContext_SetWithBigTimeout(t *testing.T) {
	cache := NewCacheWithContext()

	timeout := time.Millisecond * 200

	ctxBase := context.Background()
	ctx, _ := context.WithTimeout(ctxBase, timeout)

	err := cache.Set(ctx, "k", "v")
	if err != nil {
		t.Errorf("Expected error, %v", err)
	}
}
