package main

import (
	"fmt"

	"github.com/Soln1shko/heap-coursework/internal/benchmark"
	"github.com/Soln1shko/heap-coursework/internal/heap"
	"github.com/Soln1shko/heap-coursework/internal/utils"
)

func main() {
	fmt.Println("Тестирование производительности кучи")
	fmt.Println("===================================")

	fmt.Println("\n1. Тестирование производительности операций с кучей:")
	sizes := []int{1000, 10000, 100000}
	results := benchmark.BenchmarkHeapOperations(sizes)
	benchmark.PrintBenchmarkResults(results)

	fmt.Println("\n2. Сравнение алгоритмов сортировки:")
	size := 10000
	iterations := 5
	benchmark.CompareSortingAlgorithmsAndPrint(size, iterations)

	fmt.Println("\n3. Тестирование масштабируемости:")
	testScalability()

	fmt.Println("\n4. Сравнение производительности минимальной и максимальной кучи:")
	compareHeapTypes()
}

func testScalability() {
	sizes := []int{1000, 10000, 100000, 1000000}

	fmt.Printf("%-10s %-15s %-15s %-15s\n", "Размер", "BuildHeap (мс)", "Insert (мс)", "ExtractRoot (мс)")
	fmt.Println("----------------------------------------------------------")

	for _, size := range sizes {
		arr := utils.GenerateRandomArray(size, 1, 1000000)

		buildHeapTime := utils.MeasureTime(func() {
			heap.Heapify(arr, heap.MinHeap)
		})

		h := heap.NewHeap(heap.MinHeap, size)
		insertTime := utils.MeasureTime(func() {
			for _, val := range arr[:size] {
				h.Insert(val)
			}
		})

		extractTime := utils.MeasureTime(func() {
			for i := 0; i < size; i++ {
				h.ExtractRoot()
			}
		})

		fmt.Printf("%-10d %-15.3f %-15.3f %-15.3f\n",
			size,
			float64(buildHeapTime.Microseconds())/1000.0,
			float64(insertTime.Microseconds())/1000.0,
			float64(extractTime.Microseconds())/1000.0)
	}
}

func compareHeapTypes() {
	size := 100000
	arr := utils.GenerateRandomArray(size, 1, 1000000)

	fmt.Printf("Размер массива: %d\n", size)
	fmt.Println("----------------------------------------------------------")

	minHeapBuildTime := utils.MeasureTime(func() {
		heap.Heapify(utils.CloneArray(arr), heap.MinHeap)
	})

	minHeap := heap.NewHeap(heap.MinHeap, size)
	minHeapInsertTime := utils.MeasureTime(func() {
		for _, val := range arr {
			minHeap.Insert(val)
		}
	})

	minHeapExtractTime := utils.MeasureTime(func() {
		for i := 0; i < size; i++ {
			minHeap.ExtractRoot()
		}
	})

	maxHeapBuildTime := utils.MeasureTime(func() {
		heap.Heapify(utils.CloneArray(arr), heap.MaxHeap)
	})

	maxHeap := heap.NewHeap(heap.MaxHeap, size)
	maxHeapInsertTime := utils.MeasureTime(func() {
		for _, val := range arr {
			maxHeap.Insert(val)
		}
	})

	maxHeapExtractTime := utils.MeasureTime(func() {
		for i := 0; i < size; i++ {
			maxHeap.ExtractRoot()
		}
	})

	fmt.Printf("%-15s %-15s %-15s %-15s\n", "Тип кучи", "BuildHeap (мс)", "Insert (мс)", "ExtractRoot (мс)")
	fmt.Println("----------------------------------------------------------")

	fmt.Printf("%-15s %-15.3f %-15.3f %-15.3f\n",
		"MinHeap",
		float64(minHeapBuildTime.Microseconds())/1000.0,
		float64(minHeapInsertTime.Microseconds())/1000.0,
		float64(minHeapExtractTime.Microseconds())/1000.0)

	fmt.Printf("%-15s %-15.3f %-15.3f %-15.3f\n",
		"MaxHeap",
		float64(maxHeapBuildTime.Microseconds())/1000.0,
		float64(maxHeapInsertTime.Microseconds())/1000.0,
		float64(maxHeapExtractTime.Microseconds())/1000.0)

	fmt.Println("\nСравнение времени сортировки:")

	heapSortTime := utils.MeasureTime(func() {
		arrCopy := utils.CloneArray(arr)
		heap.HeapSort(arrCopy)
	})

	heapSortDescTime := utils.MeasureTime(func() {
		arrCopy := utils.CloneArray(arr)
		heap.HeapSortDesc(arrCopy)
	})

	fmt.Printf("%-20s %-15.3f мс\n", "HeapSort (по возр.)", float64(heapSortTime.Microseconds())/1000.0)
	fmt.Printf("%-20s %-15.3f мс\n", "HeapSort (по убыв.)", float64(heapSortDescTime.Microseconds())/1000.0)
}
