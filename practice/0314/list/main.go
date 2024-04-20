package main

import "container/list"

type LRUCache struct {
	capacity int
	li       *list.List
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	li := list.New()
	return LRUCache{
		capacity: capacity,
		li:       li,
		cache:    map[int]*list.Element{},
	}
}

type Node struct {
	key, val int
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.li.MoveToFront(e)
	return (e.Value.(Node)).val
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.cache[key]
	newE := Node{key, value}
	if ok {
		e.Value = newE
		this.li.MoveToFront(e)
		this.cache[key] = e
		return
	}
	e = &list.Element{Value: newE}
	this.li.PushFront(e)
	this.cache[key] = e
	if this.li.Len() > this.capacity {
		old := this.li.Front()
		this.li.Remove(old)
		delete(this.cache, (old.Value.(Node)).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
