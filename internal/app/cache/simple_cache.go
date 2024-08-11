package cache

import "time"

var Timeout = time.Millisecond * 100

type SimpleCache struct {
	items map[string]string
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		items: make(map[string]string),
	}
}

func (c *SimpleCache) Get(key string) (string, error) {
	item, ok := c.items[key]

	if !ok {
		return "", ErrorNotFound
	}

	return item, nil
}

func (c *SimpleCache) Set(key string, value string) error {
	time.Sleep(Timeout)
	c.items[key] = value

	return nil
}

func (c *SimpleCache) Delete(key string) error {
	delete(c.items, key)

	return nil
}
