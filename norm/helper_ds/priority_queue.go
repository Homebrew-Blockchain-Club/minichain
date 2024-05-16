// helperds提供了默克尔树所需要的辅助数据结构
package helperds

import (
	"container/heap"
)

type PriorityQueue[T any] struct {
	elements []T
	less     func(a, b T) bool // 用于确定优先级的函数
}

// 以下方法用于满足 heap.Interface 接口
func (pq PriorityQueue[T]) Len() int {
	return len(pq.elements)
}

func (pq PriorityQueue[T]) Empty() bool {
	return pq.Len() == 0
}
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.elements[i], pq.elements[j])
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.elements[i], pq.elements[j] = pq.elements[j], pq.elements[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	pq.elements = append(pq.elements, x.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	n := len(pq.elements)
	item := pq.elements[n-1]
	pq.elements = pq.elements[:n-1]
	return item
}

// 创建一个优先队列
func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		elements: make([]T, 0),
		less:     less,
	}
	heap.Init(pq)
	return pq
}
