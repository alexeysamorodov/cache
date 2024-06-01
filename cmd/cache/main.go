package main

import (
	"fmt"

	cache "github.com/alexeysamorodov/cache/internal/app/cache"
)

func main() {
	cacheImpl := cache.NewSimpleCache()

	fmt.Println("Set several values to the cache")
	cacheImpl.Set("Russia", "Moscow")
	cacheImpl.Set("England", "London")
	fmt.Println()

	fmt.Println("Read existing value from the cache")
	key := "Russia"
	item, _ := cacheImpl.Get(key)
	fmt.Printf("Read data: Key:%s, Value:%s\n", key, item)
	key = "England"
	item, _ = cacheImpl.Get(key)
	fmt.Printf("Read data: Key:%s, Value:%s\n", key, item)
	fmt.Println()

	fmt.Println("Delete data from the cache")
	cacheImpl.Delete("Russia")
	fmt.Println()

	fmt.Println("Read not existing value from the cache")
	item, err := cacheImpl.Get("Russia")
	if err != nil {
		fmt.Println(err)
	}
}
