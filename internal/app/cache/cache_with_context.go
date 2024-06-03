package cache

import (
	"context"
	"errors"
)

type CacheWithContext struct {
	cache Cache
}

var ErrorTimeout = errors.New("Timeout")

func NewCacheWithContext() *CacheWithContext {
	return &CacheWithContext{
		cache: NewSimpleCache(),
	}
}

func (c *CacheWithContext) Get(ctx context.Context, key string) (string, error) {
	ch := make(chan string)

	go func() {
		defer close(ch)

		value, err := c.cache.Get(key)

		if err == nil {
			ch <- value
		}
	}()

	select {
	case <-ctx.Done():
		return "", ErrorTimeout
	case value, ok := <-ch: // Если здесь сразу попробовать вызвать c.Cache.Get(key) - то функция будет выполнена сразу, независимо от ctx.Done
		if ok {
			return value, nil
		}
		return "", ErrorNotFound
	}
}

func (c *CacheWithContext) Set(ctx context.Context, key, value string) error {
	ch := make(chan error)

	go func() {
		defer close(ch)

		err := c.cache.Set(key, value)

		if err != nil {
			ch <- err
		}
	}()

	select {
	case <-ctx.Done():
		return ErrorTimeout
	case value, ok := <-ch:
		if ok {
			return value
		}
		return nil
	}
}

func (c *CacheWithContext) Delete(ctx context.Context, key string) error {
	ch := make(chan error)

	go func() {
		defer close(ch)

		err := c.cache.Delete(key)

		if err != nil {
			ch <- err
		}
	}()

	select {
	case <-ctx.Done():
		return ErrorTimeout
	case value, ok := <-ch:
		if ok {
			return value
		}
		return nil
	}
}
