package generic_heap

import "container/heap"

type Heap[T any] struct {
	base *base[T]
}

func New[T any](lessFunc func(T, T) bool) *Heap[T] {
	b := &base[T]{
		lessFunc: lessFunc,
	}
	heap.Init(b)
	return &Heap[T]{base: b}
}

func (h *Heap[T]) Push(t T) {
	heap.Push(h.base, t)
}

func (h *Heap[T]) Pop() T {
	if h.Len() == 0 {
		var t T
		return t
	}
	return heap.Pop(h.base).(T)
}

func (h *Heap[T]) Len() int {
	return h.base.Len()
}

func (h *Heap[T]) Peek() T {
	if h.Len() == 0 {
		var t T
		return t
	}
	return h.base.Peek().(T)
}
