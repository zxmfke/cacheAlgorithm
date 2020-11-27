package LFU

import (
	"container/heap"
	"fmt"
)

type LFUCache struct {
	capacity int

	// 最小堆实现的队列
	queue *queue

	// map的key是字符串，value是entry
	cache map[interface{}]*lfuElement

	//expireTime time.Duration
}

// 定义元素
type lfuElement struct {
	key    interface{}
	value  interface{}
	weight int // 访问次数
	index  int // queue索引
}

func New(capacity int) *LFUCache {
	q := make(queue, 0)
	return &LFUCache{
		capacity: capacity,
		cache:    make(map[interface{}]*lfuElement),
		queue:    &q,
		//expireTime: 0,
	}
}

//func (l *LRUCache) WithExpireTime(expireTime time.Duration) {
//	l.expireTime = expireTime
//}

func (l *LFUCache) len() int {
	return l.queue.Len()
}

func (l *LFUCache) Print() {
	if l.len() == 0 {
		fmt.Println("LRU Cache is empty!")
		return
	}

	element := l.queue.Pop()
	fmt.Printf("[lfu cache] : {%v,%d}", element.(*lfuElement).value, element.(*lfuElement).weight)
	for ; l.len() != 0; {
		element = l.queue.Pop()
		fmt.Printf("=>{%v,%d}", element.(*lfuElement).value, element.(*lfuElement).weight)
	}
	fmt.Printf("\n")
}

// 通过 Set 方法往 Cache 头部增加一个元素，如果存在则更新值
func (l *LFUCache) Set(key interface{}, value interface{}) {
	if en, ok := l.cache[key]; ok {
		l.queue.update(en, value, en.weight+1)
	} else {
		en := &lfuElement{
			key:    key,
			value:  value,
			weight: 1,
		}

		heap.Push(l.queue, en) // 插入queue 并重新排序为堆
		l.cache[key] = en      // 插入 map

		// 如果超出内存长度，则删除最 '无用' 的元素，0表示无内存限制
		if l.len() > l.capacity {
			l.DelOldest()
		}
	}
}

func (l *LFUCache) Get(key interface{}) interface{} {
	if en, ok := l.cache[key]; ok {
		l.queue.update(en, en.value, en.weight+1)
		return en.value
	}
	return nil
}

func (l *LFUCache) Del(key interface{}) {
	if en, ok := l.cache[key]; ok {
		heap.Remove(l.queue, en.index)
		l.removeElement(en)
	}
}

func (l *LFUCache) DelOldest() {
	if l.len() == 0 {
		return
	}
	val := heap.Pop(l.queue)
	l.removeElement(val)
}

func (l *LFUCache) removeElement(v interface{}) {
	if v == nil {
		return
	}

	en := v.(*lfuElement)

	delete(l.cache, en.key)
}
