package cache

//package main

import (
	"time"
	"fmt"
)

type Node struct {
	key string
	value string
	t time.Time
}

type Cache struct{
	node []Node
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	for i := 0; i < len(c.node) ; i++ {
		if c.node[i].t.Sub(time.Now()) > 0 || c.node[i].t.IsZero() {
			if c.node[i].key == key {
				return c.node[i].value, true
			}
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	k := 0
	for i := 0; i < len(c.node) ; i++ {
		if c.node[i].key == key {
			c.node[i].value = value
			k = 1
		}
	}
	if k == 0 {
		c.node = append(c.node, Node{key:key, value:value})
	}
}

func (c *Cache) Keys() []string {
	k := make([]string, len(c.node))
	for i := 0; i < len(c.node) ; i++ {
		if c.node[i].t.Sub(time.Now()) > 0 ||c.node[i].t.IsZero() {
			k[i] = c.node[i].key
		}
	}
	return k
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	k := 0
	for i := 0; i < len(c.node) ; i++ {
		if c.node[i].key == key {
			c.node[i].value = value
			c.node[i].t = deadline
			k = 1
		}
	}
	if k == 0 {
		c.node = append(c.node, Node{key:key, value:value, t:deadline})
	}
}

/*func main() {
	var c Cache
	c.Put("a","abc")
	c.PutTill("b","1234",time.Now().Add(time.Minute*3))
	c.Put("a","567")
	c.Put("1", "avs")
	c.PutTill("c","rt",time.Now())
	fmt.Println(c)
	fmt.Println(c.Get("1"))
	fmt.Println(c.Get("5"))
	fmt.Println(c.Get("b"))
	fmt.Println(c.Get("c"))
	fmt.Println(c.Keys())
}
*/