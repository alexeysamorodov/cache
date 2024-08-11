package main

import (
	"fmt"
	"strconv"
	"sync"

	cache "github.com/alexeysamorodov/cache/internal/app/cache"
)

func main() {
	cacheImpl := cache.NewSimpleCache()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				recover()
			}()
			defer wg.Done()
			str := strconv.Itoa(i)
			cacheImpl.Set(str, str)
		}(i)
	}
	wg.Wait()
	fmt.Println("here")
}
