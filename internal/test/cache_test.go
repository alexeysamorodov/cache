package test

import (
	"testing"

	"github.com/alexeysamorodov/cache/internal/app/cache"
	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name  string
		cache cache.Cache
	}{
		{
			name:  "SimpleCache SetGet",
			cache: cache.NewSimpleCache(),
		},
		{
			name:  "ConcurrentCache SetGet",
			cache: cache.NewConcurrentCache(),
		},
		{
			name:  "ShardedCache SetGet",
			cache: cache.NewShardedCache(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			key := "key"
			value := "value"

			// Act
			tt.cache.Set(key, value)
			actual, err := tt.cache.Get(key)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, value, actual)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name  string
		cache cache.Cache
	}{
		{
			name:  "SimpleCache Delete",
			cache: cache.NewSimpleCache(),
		},
		{
			name:  "ConcurrentCache Delete",
			cache: cache.NewConcurrentCache(),
		},
		{
			name:  "ShardedCache SetGet",
			cache: cache.NewShardedCache(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			key := "key"
			value := "value"

			// Act
			tt.cache.Set(key, value)
			deleteErr := tt.cache.Delete(key)

			actualValue, getErr := tt.cache.Get(key)

			// Assert
			assert.NoError(t, deleteErr)
			assert.Error(t, getErr)
			assert.Zero(t, actualValue)
		})
	}
}

func TestParallel(t *testing.T) {
	tests := []struct {
		name           string
		cache          cache.Cache
		parallelFactor int
	}{
		{
			name:           "ConcurrentCache test race",
			cache:          cache.NewConcurrentCache(),
			parallelFactor: 100_000,
		},
		{
			name:           "ShardedCache test race",
			cache:          cache.NewShardedCache(),
			parallelFactor: 100_000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emulateLoad(t, tt.cache, tt.parallelFactor)
		})
	}
}
