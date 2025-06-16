package heap

func (h *Heap) SiftUp(index int) {
	for index > 0 {
		parentIndex := h.Parent(index)

		if h.ShouldSwap(parentIndex, index) {
			h.Swap(parentIndex, index)
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *Heap) SiftDown(index int) {
	largest := index
	smallest := index

	for {
		leftChildIndex := h.LeftChild(index)
		rightChildIndex := h.RightChild(index)

		if h.Type == MaxHeap {
			if leftChildIndex < h.Size && h.Items[leftChildIndex] > h.Items[largest] {
				largest = leftChildIndex
			}

			if rightChildIndex < h.Size && h.Items[rightChildIndex] > h.Items[largest] {
				largest = rightChildIndex
			}

			if largest != index {
				h.Swap(index, largest)
				index = largest
			} else {
				break
			}
		} else {
			if leftChildIndex < h.Size && h.Items[leftChildIndex] < h.Items[smallest] {
				smallest = leftChildIndex
			}

			if rightChildIndex < h.Size && h.Items[rightChildIndex] < h.Items[smallest] {
				smallest = rightChildIndex
			}

			if smallest != index {
				h.Swap(index, smallest)
				index = smallest
			} else {
				break
			}
		}
	}
}

func (h *Heap) BuildHeap() {
	for i := h.Size/2 - 1; i >= 0; i-- {
		h.SiftDown(i)
	}
}

func Heapify(arr []int, heapType HeapType) {
	n := len(arr)
	h := &Heap{
		Items:    arr,
		Type:     heapType,
		Size:     n,
		Capacity: n,
	}

	h.BuildHeap()
}

func (h *Heap) HeapifyAt(index int) {
	h.SiftDown(index)
}
