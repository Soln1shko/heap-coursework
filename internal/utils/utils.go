package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomArray(size int, min, max int) []int {
	if size <= 0 {
		return []int{}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = r.Intn(max-min+1) + min
	}

	return arr
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func IsSortedDesc(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

func IsMinHeap(arr []int, size int) bool {
	for i := 0; i < size/2; i++ {
		leftChildIdx := 2*i + 1
		rightChildIdx := 2*i + 2

		if leftChildIdx < size && arr[i] > arr[leftChildIdx] {
			return false
		}

		if rightChildIdx < size && arr[i] > arr[rightChildIdx] {
			return false
		}
	}

	return true
}

func IsMaxHeap(arr []int, size int) bool {
	for i := 0; i < size/2; i++ {
		leftChildIdx := 2*i + 1
		rightChildIdx := 2*i + 2

		if leftChildIdx < size && arr[i] < arr[leftChildIdx] {
			return false
		}

		if rightChildIdx < size && arr[i] < arr[rightChildIdx] {
			return false
		}
	}

	return true
}

func PrintArray(arr []int, message string) {
	fmt.Printf("%s: %v\n", message, arr)
}
func CloneArray(arr []int) []int {
	clone := make([]int, len(arr))
	copy(clone, arr)
	return clone
}

func MeasureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
