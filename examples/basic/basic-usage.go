package main

import (
	"fmt"

	"github.com/Soln1shko/heap-coursework/internal/heap"
	"github.com/Soln1shko/heap-coursework/internal/utils"
)

func main() {
	fmt.Println("Пример базового использования кучи")
	fmt.Println("==================================")

	fmt.Println("\n1. Создание минимальной кучи:")
	minHeap := heap.NewHeap(heap.MinHeap, 10)
	fmt.Printf("Создана пустая минимальная куча: %v\n", minHeap)

	fmt.Println("\n2. Вставка элементов:")
	elements := []int{5, 3, 8, 1, 9, 4, 7, 2, 6}
	fmt.Printf("Элементы для вставки: %v\n", elements)

	for _, elem := range elements {
		minHeap.Insert(elem)
		fmt.Printf("После вставки %d: %v\n", elem, minHeap.GetItems())
	}

	fmt.Println("\n3. Получение корневого элемента (без удаления):")
	root, _ := minHeap.Peek()
	fmt.Printf("Корневой элемент: %d\n", root)
	fmt.Println("\n4. Извлечение элементов:")
	for !minHeap.IsEmpty() {
		root, _ := minHeap.ExtractRoot()
		fmt.Printf("Извлечен корень: %d, оставшиеся элементы: %v\n", root, minHeap.GetItems())
	}

	fmt.Println("\n5. Создание максимальной кучи:")
	maxHeap := heap.NewHeap(heap.MaxHeap, 10)
	fmt.Printf("Создана пустая максимальная куча: %v\n", maxHeap)
	fmt.Println("\n6. Вставка элементов в максимальную кучу:")
	for _, elem := range elements {
		maxHeap.Insert(elem)
		fmt.Printf("После вставки %d: %v\n", elem, maxHeap.GetItems())
	}

	fmt.Println("\n7. Извлечение элементов из максимальной кучи:")
	for !maxHeap.IsEmpty() {
		root, _ := maxHeap.ExtractRoot()
		fmt.Printf("Извлечен корень: %d, оставшиеся элементы: %v\n", root, maxHeap.GetItems())
	}

	fmt.Println("\n8. Создание кучи из существующего массива:")
	arr := utils.GenerateRandomArray(10, 1, 100)
	fmt.Printf("Исходный массив: %v\n", arr)

	heapFromArray := heap.NewHeapFromArray(arr, heap.MinHeap)
	fmt.Printf("Минимальная куча из массива: %v\n", heapFromArray.GetItems())

	fmt.Println("\n9. Проверка свойств кучи:")
	heapItems := heapFromArray.GetItems()
	isMinHeap := utils.IsMinHeap(heapItems, len(heapItems))
	fmt.Printf("Является ли минимальной кучей: %v\n", isMinHeap)
	fmt.Println("\n10. Представление кучи в виде массива:")
	fmt.Println(utils.GetHeapArrayRepresentation(heapItems, len(heapItems)))
	fmt.Println("\n11. Представление кучи в виде дерева:")
	fmt.Println(utils.ArrayToTreeString(heapItems, len(heapItems)))
	fmt.Println("\n12. Информация о связях между узлами:")
	fmt.Println(utils.GetNodeRelationships(heapItems, len(heapItems)))
}
