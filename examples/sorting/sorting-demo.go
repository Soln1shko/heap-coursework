package main

import (
	"fmt"
	"time"

	"github.com/Soln1shko/heap-coursework/internal/heap"
	"github.com/Soln1shko/heap-coursework/internal/utils"
)

func main() {
	fmt.Println("Демонстрация сортировки с использованием кучи")
	fmt.Println("============================================")

	size := 15
	arr := utils.GenerateRandomArray(size, 1, 100)
	fmt.Printf("\nИсходный массив (%d элементов): %v\n", size, arr)
	fmt.Println("\n1. Сортировка по возрастанию (HeapSort):")
	arrAsc := utils.CloneArray(arr)

	startTime := time.Now()
	heap.HeapSort(arrAsc)
	duration := time.Since(startTime)

	fmt.Printf("Отсортированный массив: %v\n", arrAsc)
	fmt.Printf("Время сортировки: %v\n", duration)
	fmt.Printf("Проверка сортировки: %v\n", utils.IsSorted(arrAsc))

	fmt.Println("\n2. Сортировка по убыванию (HeapSortDesc):")
	arrDesc := utils.CloneArray(arr)

	startTime = time.Now()
	heap.HeapSortDesc(arrDesc)
	duration = time.Since(startTime)

	fmt.Printf("Отсортированный массив: %v\n", arrDesc)
	fmt.Printf("Время сортировки: %v\n", duration)
	fmt.Printf("Проверка сортировки по убыванию: %v\n", utils.IsSortedDesc(arrDesc))

	fmt.Println("\n3. Пошаговая демонстрация сортировки кучей:")
	demoArr := []int{9, 4, 7, 1, 5, 3, 8, 2, 6}
	fmt.Printf("Исходный массив: %v\n", demoArr)
	fmt.Println("\nШаг 1: Преобразование массива в максимальную кучу")
	h := heap.NewHeapFromArray(demoArr, heap.MaxHeap)
	fmt.Printf("Максимальная куча: %v\n", h.GetItems())
	fmt.Println("Представление в виде дерева:")
	fmt.Println(utils.ArrayToTreeString(h.GetItems(), h.GetSize()))

	fmt.Println("\nШаг 2: Извлечение элементов из кучи")
	sortedArr := make([]int, 0, h.GetSize())

	for !h.IsEmpty() {
		root, _ := h.ExtractRoot()
		sortedArr = append(sortedArr, root)
		fmt.Printf("Извлечен корень: %d, оставшиеся элементы: %v\n", root, h.GetItems())
		if !h.IsEmpty() {
			fmt.Println("Представление в виде дерева:")
			fmt.Println(utils.ArrayToTreeString(h.GetItems(), h.GetSize()))
		}
	}

	fmt.Printf("\nОтсортированный массив: %v\n", sortedArr)

	fmt.Println("\n4. Частичная сортировка (k наименьших элементов):")
	k := 5
	fmt.Printf("Исходный массив: %v\n", arr)
	fmt.Printf("Поиск %d наименьших элементов...\n", k)

	kSmallest := heap.PartialSort(arr, k)
	fmt.Printf("%d наименьших элементов: %v\n", k, kSmallest)
}
