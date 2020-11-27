package LRU

/**
 * @author zxm
 * @description lru : Least Recently Used
 * @description lru cache implementation by list
 * @date 2:54 PM 11/27/2020
 **/

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[interface{}]*list.Element
	list     *list.List

	//expireTime time.Duration
}

type cacheElement struct {
	key   interface{}
	value interface{}
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[interface{}]*list.Element),
		list:     list.New(),
		//expireTime: 0,
	}
}

//func (l *LRUCache) WithExpireTime(expireTime time.Duration) {
//	l.expireTime = expireTime
//}

func (l *LRUCache) len() int {
	return l.list.Len()
}

func (l *LRUCache) moveToFront(e *list.Element) {
	l.list.MoveToFront(e)
}

func (l *LRUCache) removeBack() cacheElement {
	removeElement := l.list.Remove(l.list.Back())
	return removeElement.(cacheElement)
}

func (l *LRUCache) Print() {
	if l.len() == 0 {
		fmt.Println("LRU Cache is empty!")
		return
	}

	front := l.list.Front()
	fmt.Printf("[lru cache] : %v", front.Value.(cacheElement).value)
	for ; front.Next() != nil; {
		front = front.Next()
		fmt.Printf("=>%v", front.Value.(cacheElement).value)
	}
	fmt.Printf("\n")
}

func (l *LRUCache) HasKey(key interface{}) bool {
	_, ok := l.cache[key]
	return ok
}

func (l *LRUCache) Put(key interface{}, value interface{}) {
	e := cacheElement{
		key:   key,
		value: value,
	}
	if l.HasKey(key) {
		fmt.Println("[W] cache move to front")
		element := l.cache[key]
		element.Value = e
		l.moveToFront(element)
		return
	}

	if l.len() == l.capacity {
		fmt.Println("[W] lru run out of !")
		removeElement := l.removeBack()
		delete(l.cache, removeElement.key)
	}

	l.cache[key] = l.list.PushFront(e)
}

func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	if !l.HasKey(key) {
		return nil, false
	}

	element := l.cache[key]
	l.moveToFront(element)

	return element.Value.(cacheElement).value, true
}
