package LFU

import "container/heap"

//source  https://www.jianshu.com/p/722b29558958

// 最小堆实现的队列
type queue []*lfuElement

// 队列长度
func (q queue) Len() int {
	return len(q)
}

// '<' 是最小堆，'>' 是最大堆
func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

// 交换元素
func (q queue) Swap(i, j int) {
	// 交换元素
	q[i], q[j] = q[j], q[i]
	// 索引不用交换
	q[i].index = i
	q[j].index = j
}

// append ，*q = oldQue[:n-1] 会导致频繁的内存拷贝
// 实际上，如果使用 LFU算法，处于性能考虑，可以将最大内存限制修改为最大记录数限制
// 这样提前分配好 queue 的容量，再使用交换索引和限制索引的方式来实现 Pop 方法，可以免去频繁的内存拷贝，极大提高性能
func (q *queue) Push(v interface{}) {
	n := q.Len()
	en := v.(*lfuElement)
	en.index = n
	*q = append(*q, en) // 这里会重新分配内存，并拷贝数据
}

func (q *queue) Pop() interface{} {
	oldQue := *q
	n := len(oldQue)
	en := oldQue[n-1]
	oldQue[n-1] = nil // 将不再使用的对象置为nil，加快垃圾回收，避免内存泄漏
	*q = oldQue[:n-1] // 这里会重新分配内存，并拷贝数据
	return en
}

// weight更新后，要重新排序，时间复杂度为 O(logN)
func (q *queue) update(en *lfuElement, val interface{}, weight int) {
	en.value = val
	en.weight = weight
	(*q)[en.index] = en
	// 重新排序
	// 分析思路是把 堆(大D) 的树状图画出来，看成一个一个小的堆(小D)，看改变其中一个值，对 大D 有什么影响
	// 可以得出结论，下沉操作和上沉操作分别执行一次能将 queue 排列为堆
	heap.Fix(q, en.index)
}
