package bench

import (
	"errors"
	"fmt"
	"sync"

	"github.com/alexeysamorodov/cache/internal/app/cache"
)

func emulateLoad(c cache.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		// Запись в кэш
		wg.Add(1)
		go func(k, v string) {
			err := c.Set(k, v)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key, value)

		// Чтение из кэша
		wg.Add(1)
		go func(k, v string) {
			_, err := c.Get(k)
			if err != nil && !errors.Is(err, cache.ErrorNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)

		// Удаление из кэша
		wg.Add(1)
		go func(k string) {
			c.Delete(k)
			wg.Done()
		}(key)
	}

	wg.Wait()
}
