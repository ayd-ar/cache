package cache

type Cache struct {
	item map[string]interface{}
}

func New() *Cache {
	return &Cache{
		item: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.item[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.item[key]
}

func (c *Cache) Delete(key string) {
	delete(c.item, key)
}
