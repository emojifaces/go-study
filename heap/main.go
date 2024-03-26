package main

import (
	"container/heap"
	"fmt"
)

// Item 是优先队列中的元素
type Item struct {
	value    string // 元素的值，可以是任意的可比较类型
	priority int    // 元素的优先级
}

// PriorityQueue 实现了 heap.Interface，并持有 Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们希望 Pop 给我们最高的，而不是最低的，优先级所以使用大于号。
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	// 创建一个优先队列，并放入一些元素
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
		}
		i++
	}
	heap.Init(&pq)

	// 插入一个新元素，然后修改其优先级
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	// 取出并打印队列中的所有元素
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
