package cache

type SimpleCache struct {
	items map[string]string
}

func NewSimpleCache() Cache {
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
	c.items[key] = value

	return nil
}

func (c *SimpleCache) Delete(key string) error {
	delete(c.items, key)

	return nil
}
