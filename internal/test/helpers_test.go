package test

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/alexeysamorodov/cache/internal/app/cache"
	"github.com/stretchr/testify/assert"
)

func emulateLoad(t *testing.T, c cache.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		// Запись в кэш
		wg.Add(1)
		go func(k, v string) {
			err := c.Set(k, v)
			assert.NoError(t, err)
			wg.Done()
		}(key, value)

		// Чтение из кэша
		wg.Add(1)
		go func(k, v string) {
			actualValue, err := c.Get(k)
			if !errors.Is(err, cache.ErrorNotFound) {
				assert.Equal(t, v, actualValue)
			}
			wg.Done()
		}(key, value)

		// Удаление из кэша
		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key)
	}

	wg.Wait()
}
