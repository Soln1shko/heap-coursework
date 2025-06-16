package heap

type HeapType int

const (
	MinHeap HeapType = iota
	MaxHeap
)

type Heap struct {
	Items    []int
	Type     HeapType
	Size     int
	Capacity int
}

func NewHeap(heapType HeapType, capacity int) *Heap {
	if capacity <= 0 {
		capacity = 10
	}

	return &Heap{
		Items:    make([]int, capacity),
		Type:     heapType,
		Size:     0,
		Capacity: capacity,
	}
}

func NewHeapFromArray(arr []int, heapType HeapType) *Heap {
	capacity := len(arr)
	heap := &Heap{
		Items:    make([]int, capacity),
		Type:     heapType,
		Size:     capacity,
		Capacity: capacity,
	}

	copy(heap.Items, arr)

	heap.BuildHeap()

	return heap
}
