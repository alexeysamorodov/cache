package cache

type ChannelCache struct {
	cache  Cache
	inChan chan Data
}

type Data struct {
	Key   string
	Value string
}

func NewChannelCache() *ChannelCache {
	return &ChannelCache{
		cache:  NewSimpleCache(),
		inChan: make(chan Data),
	}
}

func (c *ChannelCache) GetChannel() chan<- Data {
	return c.inChan
}

func (c *ChannelCache) Run() {
	go func() {
		for {
			x, ok := <-c.inChan
			if !ok {
				return
			}

			c.cache.Set(x.Key, x.Value)
		}
	}()
}

func (c *ChannelCache) Get(key string) (string, error) {
	value, err := c.cache.Get(key)

	if err != nil {
		return "", err
	}

	return value, nil
}
