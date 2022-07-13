package heap

import "container/heap"

// Heap is a binary min heap, generic wrapper of container/heap
type Heap[T any] struct {
	base *base[T]
}

// New return a binary min heap with corresponding type
func New[T any](lessFunc func(T, T) bool) *Heap[T] {
	b := &base[T]{
		lessFunc: lessFunc,
	}
	heap.Init(b)
	return &Heap[T]{base: b}
}

// Push add an element to the heap
func (h *Heap[T]) Push(t T) {
	heap.Push(h.base, t)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Pop() T {
	if h.Len() == 0 {
		var t T
		return t
	}
	return heap.Pop(h.base).(T)
}

// Len return heap size
func (h *Heap[T]) Len() int {
	return h.base.Len()
}

// Peek returns the minimum element (according to Less) from the heap.
// The complexity is O(1).
func (h *Heap[T]) Peek() T {
	if h.Len() == 0 {
		var t T
		return t
	}
	return h.base.Peek().(T)
}
