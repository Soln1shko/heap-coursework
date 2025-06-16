package heap

import (
	"errors"
	"fmt"
)

func (h *Heap) Parent(index int) int {
	return (index - 1) / 2
}
func (h *Heap) LeftChild(index int) int {
	return 2*index + 1
}

func (h *Heap) RightChild(index int) int {
	return 2*index + 2
}

func (h *Heap) Swap(i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
}

func (h *Heap) ShouldSwap(parentIndex, childIndex int) bool {
	if h.Type == MinHeap {
		return h.Items[parentIndex] > h.Items[childIndex]
	}
	return h.Items[parentIndex] < h.Items[childIndex]
}

func (h *Heap) Insert(item int) error {
	if h.Size >= h.Capacity {
		return errors.New("переполнение кучи: достигнута максимальная емкость")
	}
	h.Items[h.Size] = item
	h.Size++
	h.SiftUp(h.Size - 1)

	return nil
}

func (h *Heap) ExtractRoot() (int, error) {
	if h.Size <= 0 {
		return 0, errors.New("куча пуста")
	}
	root := h.Items[0]
	h.Items[0] = h.Items[h.Size-1]
	h.Size--
	h.SiftDown(0)

	return root, nil
}
func (h *Heap) Peek() (int, error) {
	if h.Size <= 0 {
		return 0, errors.New("куча пуста")
	}
	return h.Items[0], nil
}

func (h *Heap) IsEmpty() bool {
	return h.Size == 0
}
func (h *Heap) Clear() {
	h.Size = 0
}

func (h *Heap) GetSize() int {
	return h.Size
}
func (h *Heap) GetCapacity() int {
	return h.Capacity
}

func (h *Heap) GetItems() []int {
	result := make([]int, h.Size)
	copy(result, h.Items[:h.Size])
	return result
}

func (h *Heap) String() string {
	heapType := "MinHeap"
	if h.Type == MaxHeap {
		heapType = "MaxHeap"
	}

	return fmt.Sprintf("%s: %v (размер: %d, емкость: %d)",
		heapType, h.Items[:h.Size], h.Size, h.Capacity)
}

func (h *Heap) Contains(item int) bool {
	for i := 0; i < h.Size; i++ {
		if h.Items[i] == item {
			return true
		}
	}
	return false
}

func (h *Heap) Remove(item int) bool {
	for i := 0; i < h.Size; i++ {
		if h.Items[i] == item {
			h.Items[i] = h.Items[h.Size-1]
			h.Size--

			// Восстанавливаем свойство кучи
			if i > 0 && h.ShouldSwap(h.Parent(i), i) {
				h.SiftUp(i)
			} else {
				h.SiftDown(i)
			}

			return true
		}
	}
	return false
}
