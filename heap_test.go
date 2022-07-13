package heap

import (
	"testing"
)

type node struct {
	val int
}

func TestGenericHeap(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		data := []int{0, 5, 4, 2, -5, 8, 9, -10}
		expect := []int{-10, -5, 0, 2, 4, 5, 8, 9}
		lessFunc := func(i, j int) bool {
			return j-i > 0
		}
		heap := New(lessFunc)
		for _, v := range data {
			heap.Push(v)
		}
		for i := 0; i < len(data); i++ {
			result := heap.Pop()
			if expect[i] != result {
				t.Error("wrong value", result, "expect", expect[i])
			}
		}
	})

	t.Run("node", func(t *testing.T) {
		data := []int{0, 5, 4, 2, -5, 8, 9, -10}
		expect := []int{-10, -5, 0, 2, 4, 5, 8, 9}
		lessFunc := func(i, j node) bool {
			return j.val-i.val > 0
		}
		heap := New(lessFunc)
		for _, v := range data {
			heap.Push(node{val: v})
		}
		for i := 0; i < len(data); i++ {
			result := heap.Pop().val
			if expect[i] != result {
				t.Error("wrong value", result, "expect", expect[i])
			}
		}
	})

	t.Run("ec", func(t *testing.T) {
		lessFunc := func(i, j int) bool {
			return j-i > 0
		}
		heap := New(lessFunc)
		if heap.Pop() != 0 {
			t.Error("err pop empty heap")
		}

		if heap.Peek() != 0 {
			t.Error("err peek empty heap")
		}

		heap.Push(1)
		heap.Push(-1)
		heap.Push(5)
		if heap.Peek() != -1 {
			t.Error("err peek wrong value")
		}
		if heap.Peek() != -1 {
			t.Error("err peek wrong value")
		}
	})
}
