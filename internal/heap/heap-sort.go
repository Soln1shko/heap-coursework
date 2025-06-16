package heap

func HeapSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	Heapify(arr, MaxHeap)

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		h := &Heap{
			Items:    arr,
			Type:     MaxHeap,
			Size:     i,
			Capacity: n,
		}
		h.SiftDown(0)
	}
}

func HeapSortDesc(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	Heapify(arr, MinHeap)
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		h := &Heap{
			Items:    arr,
			Type:     MinHeap,
			Size:     i,
			Capacity: n,
		}
		h.SiftDown(0)
	}
}

func PartialSort(arr []int, k int) []int {
	n := len(arr)
	if k <= 0 || n <= 0 {
		return []int{}
	}

	if k >= n {
		result := make([]int, n)
		copy(result, arr)
		HeapSort(result)
		return result
	}

	h := NewHeap(MaxHeap, k)

	for i := 0; i < k; i++ {
		h.Insert(arr[i])
	}

	for i := k; i < n; i++ {
		root, _ := h.Peek()
		if arr[i] < root {
			h.ExtractRoot()
			h.Insert(arr[i])
		}
	}

	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i], _ = h.ExtractRoot()
	}

	return result
}
