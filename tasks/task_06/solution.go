package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	items := make(map[K]*list.Element)
	return &LRUCache[K, V]{capacity: capacity, items: items}
}
func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	elem, ok := c.items[key]
	if !ok {
		return value, false
	}
	c.ll.MoveToFront(elem)
	ent := elem.Value.(*entry[K, V])
	return ent.value, true
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}
	if elem, ok := c.items[key]; ok {
		ent := elem.Value.(*entry[K, V])
		ent.value = value
		return
	}
	if c.ll.Len() >= c.capacity {
		oldest := c.ll.Back()
		ent := oldest.Value.(*entry[K, V])
		c.ll.Remove(oldest)
		delete(c.items, ent.key)
	}

	elem := c.ll.PushFront(&entry[K, V]{key: key, value: value})
	c.items[key] = elem
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
