package main

import (
	"fmt"
)

type LRUCache struct {
	M        map[int]*Element
	li       *List
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		M:        map[int]*Element{},
		li:       New(),
		capacity: capacity,
	}
}

type entry struct {
	key, value int
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.M[key]; !ok {
		return -1
	}
	element := c.M[key]
	c.li.MoveToFront(element)
	return element.Value.(entry).value
}

func (c *LRUCache) Put(key, value int) {
	element, ok := c.M[key]
	e := entry{key: key, value: value}
	if !ok {
		element = c.li.PushFront(e)
		c.M[key] = element
	} else {
		element.Value = e
		c.li.MoveToFront(element)
	}

	if c.li.Len() > c.capacity {
		tail := c.li.Back()
		delete(c.M, tail.Value.(entry).key)
		c.li.Remove(tail)
	}
}

func main() {
	lruCache := NewLRUCache(2)
	res := lruCache.Get(1)
	fmt.Println(res)
	lruCache.Put(1, 100)
	res = lruCache.Get(1)
	fmt.Println(res)
	lruCache.Put(1, 200)
	res = lruCache.Get(1)
	fmt.Println(res)
	lruCache.Put(2, 3)
	lruCache.Put(3, 4)
	res = lruCache.Get(1)
	fmt.Println(res)
}

type Element struct {
	next, prev *Element
	list       *List
	Value      any
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List { return new(List).Init() }

func (l *List) Len() int { return l.len }

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v any, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *List) remove(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}

func (l *List) move(e, at *Element) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

func (l *List) Remove(e *Element) any {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List) PushFront(v any) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}
