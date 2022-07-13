package generic_heap

type base[T any] struct {
	data     []T
	lessFunc func(a, b T) bool
}

func (b *base[T]) Len() int {
	return len(b.data)
}

func (b *base[T]) Less(i, j int) bool {
	return b.lessFunc(b.data[i], b.data[j])
}

func (b *base[T]) Swap(i, j int) {
	tmp := b.data[i]
	b.data[i] = b.data[j]
	b.data[j] = tmp
}

func (b *base[T]) Push(x any) {
	b.data = append(b.data, x.(T))
}

func (b *base[T]) Pop() any {
	old := b.data
	oldLen := len(old)
	res := old[oldLen-1]
	b.data = b.data[:oldLen-1]
	return res
}

func (b *base[T]) Peek() any {
	return b.data[0]
}
