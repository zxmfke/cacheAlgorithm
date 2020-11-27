package LRU

import "testing"

func TestLRUCache_Put(t *testing.T) {
	lru := New(5)

	test := []int{0, 1, 2, 3, 0, 4, 5, 6, 7}

	for i := 0; i < len(test); i++ {
		lru.Put(test[i], test[i])
		lru.Print()
	}
}
