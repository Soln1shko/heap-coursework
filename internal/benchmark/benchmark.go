package benchmark

import (
	"fmt"
	"time"

	"github.com/Soln1shko/heap-coursework/internal/heap"
	"github.com/Soln1shko/heap-coursework/internal/utils"
)

type BenchmarkResult struct {
	OperationName string
	ArraySize     int
	Duration      time.Duration
}

func BenchmarkHeapOperations(sizes []int) []BenchmarkResult {
	results := make([]BenchmarkResult, 0)

	for _, size := range sizes {
		arr := utils.GenerateRandomArray(size, 1, 1000)

		buildHeapDuration := utils.MeasureTime(func() {
			heap.Heapify(arr, heap.MinHeap)
		})
		results = append(results, BenchmarkResult{
			OperationName: "BuildHeap",
			ArraySize:     size,
			Duration:      buildHeapDuration,
		})

		heapSortDuration := utils.MeasureTime(func() {
			arrCopy := utils.CloneArray(arr)
			heap.HeapSort(arrCopy)
		})
		results = append(results, BenchmarkResult{
			OperationName: "HeapSort",
			ArraySize:     size,
			Duration:      heapSortDuration,
		})

		h := heap.NewHeap(heap.MinHeap, size)
		insertDuration := utils.MeasureTime(func() {
			for _, val := range arr {
				h.Insert(val)
			}
		})
		results = append(results, BenchmarkResult{
			OperationName: "Insert",
			ArraySize:     size,
			Duration:      insertDuration,
		})

		extractDuration := utils.MeasureTime(func() {
			for i := 0; i < size; i++ {
				h.ExtractRoot()
			}
		})
		results = append(results, BenchmarkResult{
			OperationName: "ExtractRoot",
			ArraySize:     size,
			Duration:      extractDuration,
		})
	}

	return results
}

func PrintBenchmarkResults(results []BenchmarkResult) {
	fmt.Println("Результаты тестирования производительности:")
	fmt.Println("------------------------------------------")
	fmt.Printf("%-15s %-15s %-15s\n", "Операция", "Размер массива", "Время (мс)")
	fmt.Println("------------------------------------------")

	for _, result := range results {
		fmt.Printf("%-15s %-15d %-15.3f\n",
			result.OperationName,
			result.ArraySize,
			float64(result.Duration.Microseconds())/1000.0)
	}
}

func CompareSortingAlgorithms(size int, iterations int) map[string]time.Duration {
	results := make(map[string]time.Duration)
	totalHeapSort := time.Duration(0)
	totalQuickSort := time.Duration(0)
	totalMergeSort := time.Duration(0)

	for i := 0; i < iterations; i++ {
		arr := utils.GenerateRandomArray(size, 1, 1000)

		arrCopy1 := utils.CloneArray(arr)
		heapSortDuration := utils.MeasureTime(func() {
			heap.HeapSort(arrCopy1)
		})
		totalHeapSort += heapSortDuration

		arrCopy2 := utils.CloneArray(arr)
		quickSortDuration := utils.MeasureTime(func() {
			quickSort(arrCopy2, 0, len(arrCopy2)-1)
		})
		totalQuickSort += quickSortDuration

		arrCopy3 := utils.CloneArray(arr)
		mergeSortDuration := utils.MeasureTime(func() {
			mergeSort(arrCopy3)
		})
		totalMergeSort += mergeSortDuration
	}

	results["HeapSort"] = totalHeapSort / time.Duration(iterations)
	results["QuickSort"] = totalQuickSort / time.Duration(iterations)
	results["MergeSort"] = totalMergeSort / time.Duration(iterations)

	return results
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func BenchmarkHeapOperationsWithSizes() {
	sizes := []int{100, 1000, 10000, 100000}
	results := BenchmarkHeapOperations(sizes)
	PrintBenchmarkResults(results)
}

func CompareSortingAlgorithmsAndPrint(size int, iterations int) {
	fmt.Printf("Сравнение алгоритмов сортировки (размер массива: %d, итераций: %d):\n", size, iterations)
	fmt.Println("------------------------------------------")

	results := CompareSortingAlgorithms(size, iterations)

	fmt.Println("Среднее время выполнения:")
	for algo, duration := range results {
		fmt.Printf("%-15s: %.3f мс\n", algo, float64(duration.Microseconds())/1000.0)
	}
}
