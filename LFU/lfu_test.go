package LFU

import "testing"

func TestLRUCache_Put(t *testing.T) {
	lfu := New(5)

	test := []int{0, 1, 2, 3, 0, 3, 7, 9}

	for i := 0; i < len(test); i++ {
		lfu.Set(test[i], test[i])
	}
	lfu.Print()
}
