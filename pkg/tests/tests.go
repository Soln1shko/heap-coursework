package tests

import (
	"fmt"
	"testing"

	"github.com/Soln1shko/heap-coursework/internal/heap"
	"github.com/Soln1shko/heap-coursework/internal/utils"
)

func TestHeapOperations(t *testing.T) {
	minHeap := heap.NewHeap(heap.MinHeap, 10)
	elements := []int{5, 3, 8, 1, 9, 4, 7, 2, 6}
	for _, elem := range elements {
		err := minHeap.Insert(elem)
		if err != nil {
			t.Errorf("Ошибка при вставке элемента %d: %v", elem, err)
		}
	}

	root, err := minHeap.Peek()
	if err != nil {
		t.Errorf("Ошибка при получении корня: %v", err)
	}
	if root != 1 {
		t.Errorf("Ожидался корень 1, получен %d", root)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, exp := range expected {
		val, err := minHeap.ExtractRoot()
		if err != nil {
			t.Errorf("Ошибка при извлечении корня #%d: %v", i, err)
		}
		if val != exp {
			t.Errorf("Ожидался элемент %d, получен %d", exp, val)
		}
	}

	if !minHeap.IsEmpty() {
		t.Error("Куча должна быть пуста после извлечения всех элементов")
	}
}

func TestMaxHeap(t *testing.T) {
	maxHeap := heap.NewHeap(heap.MaxHeap, 10)
	elements := []int{5, 3, 8, 1, 9, 4, 7, 2, 6}
	for _, elem := range elements {
		maxHeap.Insert(elem)
	}

	root, _ := maxHeap.Peek()
	if root != 9 {
		t.Errorf("Ожидался корень 9, получен %d", root)
	}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i, exp := range expected {
		val, _ := maxHeap.ExtractRoot()
		if val != exp {
			t.Errorf("Ожидался элемент %d, получен %d на позиции %d", exp, val, i)
		}
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{9, 4, 7, 1, 5, 3, 8, 2, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	heap.HeapSort(arr)

	for i, val := range arr {
		if val != expected[i] {
			t.Errorf("Ошибка сортировки: ожидался %d на позиции %d, получен %d", expected[i], i, val)
		}
	}
}

func TestHeapSortDesc(t *testing.T) {
	arr := []int{9, 4, 7, 1, 5, 3, 8, 2, 6}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	heap.HeapSortDesc(arr)

	for i, val := range arr {
		if val != expected[i] {
			t.Errorf("Ошибка сортировки по убыванию: ожидался %d на позиции %d, получен %d", expected[i], i, val)
		}
	}
}

func TestHeapFromArray(t *testing.T) {
	arr := []int{9, 4, 7, 1, 5, 3, 8, 2, 6}
	minHeap := heap.NewHeapFromArray(arr, heap.MinHeap)

	root, _ := minHeap.Peek()
	if root != 1 {
		t.Errorf("Ожидался корень 1, получен %d", root)
	}

	maxHeap := heap.NewHeapFromArray(arr, heap.MaxHeap)
	root, _ = maxHeap.Peek()
	if root != 9 {
		t.Errorf("Ожидался корень 9, получен %d", root)
	}
}

func RunAllTests() {
	fmt.Println("Запуск тестов кучи...")

	fmt.Println("\nТест минимальной кучи:")
	minHeap := heap.NewHeap(heap.MinHeap, 10)
	elements := []int{5, 3, 8, 1, 9, 4, 7, 2, 6}

	fmt.Println("Вставка элементов:", elements)
	for _, elem := range elements {
		minHeap.Insert(elem)
		fmt.Printf("После вставки %d: %v\n", elem, minHeap.GetItems())
	}

	fmt.Println("\nИзвлечение элементов:")
	for !minHeap.IsEmpty() {
		root, _ := minHeap.ExtractRoot()
		fmt.Printf("Извлечен корень: %d, оставшиеся элементы: %v\n", root, minHeap.GetItems())
	}

	fmt.Println("\nТест максимальной кучи:")
	maxHeap := heap.NewHeap(heap.MaxHeap, 10)

	fmt.Println("Вставка элементов:", elements)
	for _, elem := range elements {
		maxHeap.Insert(elem)
		fmt.Printf("После вставки %d: %v\n", elem, maxHeap.GetItems())
	}

	fmt.Println("\nИзвлечение элементов:")
	for !maxHeap.IsEmpty() {
		root, _ := maxHeap.ExtractRoot()
		fmt.Printf("Извлечен корень: %d, оставшиеся элементы: %v\n", root, maxHeap.GetItems())
	}

	fmt.Println("\nТест сортировки кучей:")
	arr := utils.GenerateRandomArray(10, 1, 100)
	fmt.Println("Исходный массив:", arr)

	arrCopy := utils.CloneArray(arr)
	heap.HeapSort(arrCopy)
	fmt.Println("После сортировки по возрастанию:", arrCopy)

	arrCopy = utils.CloneArray(arr)
	heap.HeapSortDesc(arrCopy)
	fmt.Println("После сортировки по убыванию:", arrCopy)

	fmt.Println("\nТесты завершены.")
}
