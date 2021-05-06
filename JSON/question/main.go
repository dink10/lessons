package main

import "sync"

type Link struct {
	URL string
}

type Cache struct {
	items  map[string]interface{}
	status bool    // false
	len    int     // 0
	link   *Link   // nil correct
	link2  *string // nil wrong
	mux    sync.Mutex
}

func NewCache() *Cache {
	url := "google.com"
	return &Cache{
		//items: map[string]interface{}{},
		items: make(map[string]interface{}),
		link2: &url,
	}
}

func (c *Cache) AddItem(key string, val interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.items[key] = val
}

func (c *Cache) GetItem(key string) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.items[key]
}

func main() {
	cache := NewCache()

	for i := 0; i < 10; i++ {
		go func(i int) {
			cache.AddItem("val1", i)
		}(i)
	}

	_ = cache.GetItem("val2")
}
