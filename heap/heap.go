package heap

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	Data []T
	Size int
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Size == 0 && len(h.Data) == 0
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{
		Data: []T{},
		Size: 0,
	}
}

func (h *Heap[T]) Add(i T) {
	h.Size++
	h.Data = append(h.Data, i)
	h.HeapfyUp(h.Size - 1)
}

func (h *Heap[T]) HeapfyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if h.Data[parent] < h.Data[idx] {
			break
		}
		h.Data[parent], h.Data[idx] = h.Data[idx], h.Data[parent]
		idx = parent
	}
}

func (h *Heap[T]) HeapfyDown(idx int) {
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		min := idx

		if left < h.Size && h.Data[left] < h.Data[min] {
			min = left
		}
		if right < h.Size && h.Data[right] < h.Data[min] {
			min = right
		}
		if min == idx {
			break
		}
		h.Data[idx], h.Data[min] = h.Data[min], h.Data[idx]
		idx = min
	}
}
