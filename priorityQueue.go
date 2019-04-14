package gograph

import (
	"container/heap"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	items []Vertex
	// value to index
	m map[Vertex]int
	// value to priority
	pr map[Vertex]int64
}

func (pq *PriorityQueue) Len() int           { return len(pq.items) }
func (pq *PriorityQueue) Less(i, j int) bool { return pq.pr[pq.items[i]] < pq.pr[pq.items[j]] }
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i]] = i
	pq.m[pq.items[j]] = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(Vertex)
	pq.m[item] = n
	pq.items = append(pq.items, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.m[item] = -1
	pq.items = old[0 : n-1]
	return item
}

// update modifies the priority of an item in the queue.
func (pq *PriorityQueue) Update(item Vertex, priority int64) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}
func (pq *PriorityQueue) AddWithPriority(item Vertex, priority int64) {
	heap.Push(pq, item)
	pq.Update(item, priority)
}
