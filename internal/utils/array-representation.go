package utils

import (
	"fmt"
	"strings"
)

func ArrayToTreeString(arr []int, size int) string {
	if size == 0 {
		return "Пустое дерево"
	}

	var result strings.Builder
	maxLevel := 0

	for i := 0; (1<<i)-1 < size; i++ {
		maxLevel = i
	}

	for level := 0; level <= maxLevel; level++ {
		indent := strings.Repeat(" ", (1<<(maxLevel-level))-1)

		spacing := strings.Repeat(" ", (1<<(maxLevel-level+1))-1)

		result.WriteString(indent)

		startIdx := (1 << level) - 1
		endIdx := (1 << (level + 1)) - 1

		for i := startIdx; i < endIdx && i < size; i++ {
			result.WriteString(fmt.Sprintf("%d", arr[i]))

			if i < endIdx-1 && i+1 < size {
				result.WriteString(spacing)
			}
		}

		result.WriteString("\n")
	}

	return result.String()
}

func GetNodeRelationships(arr []int, size int) string {
	if size == 0 {
		return "Пустая куча"
	}

	var result strings.Builder
	result.WriteString("Связи между узлами в куче:\n")

	for i := 0; i < size; i++ {
		result.WriteString(fmt.Sprintf("Узел[%d] = %d\n", i, arr[i]))

		parentIdx := (i - 1) / 2
		leftChildIdx := 2*i + 1
		rightChildIdx := 2*i + 2

		if i > 0 {
			result.WriteString(fmt.Sprintf("  Родитель[%d] = %d\n", parentIdx, arr[parentIdx]))
		} else {
			result.WriteString("  Корневой узел\n")
		}

		if leftChildIdx < size {
			result.WriteString(fmt.Sprintf("  Левый потомок[%d] = %d\n", leftChildIdx, arr[leftChildIdx]))
		} else {
			result.WriteString("  Нет левого потомка\n")
		}

		if rightChildIdx < size {
			result.WriteString(fmt.Sprintf("  Правый потомок[%d] = %d\n", rightChildIdx, arr[rightChildIdx]))
		} else {
			result.WriteString("  Нет правого потомка\n")
		}

		result.WriteString("\n")
	}

	return result.String()
}

func GetHeapArrayRepresentation(arr []int, size int) string {
	if size == 0 {
		return "Пустая куча"
	}

	var result strings.Builder
	result.WriteString("Представление кучи в виде массива:\n")

	result.WriteString("Индекс: ")
	for i := 0; i < size; i++ {
		result.WriteString(fmt.Sprintf("%3d ", i))
	}
	result.WriteString("\n")

	result.WriteString("Значение: ")
	for i := 0; i < size; i++ {
		result.WriteString(fmt.Sprintf("%3d ", arr[i]))
	}
	result.WriteString("\n")

	return result.String()
}
