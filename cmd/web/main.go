package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Soln1shko/heap-coursework/internal/heap"
	"github.com/Soln1shko/heap-coursework/internal/utils"
)

type HeapData struct {
	Items         []int
	Size          int
	Type          string
	TreeView      string
	ArrayView     string
	Relationships string
}

type HeapOperation struct {
	Operation string      `json:"operation"`
	Value     interface{} `json:"value"`
}

type HeapResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	Items         []int  `json:"items"`
	TreeView      string `json:"treeView"`
	ArrayView     string `json:"arrayView"`
	Relationships string `json:"relationships"`
}

var (
	currentHeap *heap.Heap
)

func main() {
	currentHeap = heap.NewHeap(heap.MinHeap, 100)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/heap", heapAPIHandler)
	http.HandleFunc("/api/sort", sortAPIHandler)

	fs := http.FileServer(http.Dir("cmd/web/templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("cmd/web/templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	heapType := "Минимальная куча"
	if currentHeap.Type == heap.MaxHeap {
		heapType = "Максимальная куча"
	}

	items := currentHeap.GetItems()
	size := currentHeap.GetSize()

	data := HeapData{
		Items:         items,
		Size:          size,
		Type:          heapType,
		TreeView:      utils.ArrayToTreeString(items, size),
		ArrayView:     utils.GetHeapArrayRepresentation(items, size),
		Relationships: utils.GetNodeRelationships(items, size),
	}

	tmpl.Execute(w, data)
}

func heapAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Метод не разрешен"})
		return
	}

	var operation HeapOperation
	err := json.NewDecoder(r.Body).Decode(&operation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Некорректный JSON"})
		return
	}

	var response HeapResponse
	response.Success = true

	switch operation.Operation {
	case "insert":
		var value int
		switch v := operation.Value.(type) {
		case float64:
			value = int(v)
		case int:
			value = v
		default:
			response.Success = false
			response.Message = "Неверный формат данных для вставки"
			break
		}

		err := currentHeap.Insert(value)
		if err != nil {
			response.Success = false
			response.Message = "Ошибка вставки: " + err.Error()
		} else {
			response.Message = fmt.Sprintf("Элемент %d успешно вставлен", value)
		}

	case "extract":
		val, err := currentHeap.ExtractRoot()
		if err != nil {
			response.Success = false
			response.Message = "Ошибка извлечения: " + err.Error()
		} else {
			response.Message = fmt.Sprintf("Извлечен элемент %d", val)
		}

	case "clear":
		currentHeap.Clear()
		response.Message = "Куча очищена"

	case "changeType":
		if currentHeap.Type == heap.MinHeap {
			currentHeap = heap.NewHeapFromArray(currentHeap.GetItems(), heap.MaxHeap)
			response.Message = "Тип кучи изменен на максимальную"
		} else {
			currentHeap = heap.NewHeapFromArray(currentHeap.GetItems(), heap.MinHeap)
			response.Message = "Тип кучи изменен на минимальную"
		}

	case "buildFromArray":
		var numStr string
		switch v := operation.Value.(type) {
		case string:
			numStr = v
		case float64:
			numStr = strconv.FormatFloat(v, 'f', -1, 64)
		case int:
			numStr = strconv.Itoa(v)
		default:
			response.Success = false
			response.Message = "Неверный формат данных для создания кучи"
			break
		}

		parts := strings.Split(numStr, ",")
		nums := make([]int, 0, len(parts))

		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err == nil {
				nums = append(nums, num)
			}
		}

		if len(nums) > 0 {
			currentHeap = heap.NewHeapFromArray(nums, currentHeap.Type)
			response.Message = "Куча создана из массива"
		} else {
			response.Success = false
			response.Message = "Не удалось создать кучу из массива"
		}

	default:
		response.Success = false
		response.Message = "Неизвестная операция"
	}

	items := currentHeap.GetItems()
	size := currentHeap.GetSize()
	response.Items = items
	response.TreeView = utils.ArrayToTreeString(items, size)
	response.ArrayView = utils.GetHeapArrayRepresentation(items, size)
	response.Relationships = utils.GetNodeRelationships(items, size)

	json.NewEncoder(w).Encode(response)
}

func sortAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Метод не разрешен"})
		return
	}

	var request struct {
		Array []int  `json:"array"`
		Order string `json:"order"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Некорректный JSON"})
		return
	}

	arr := make([]int, len(request.Array))
	copy(arr, request.Array)

	if request.Order == "desc" {
		heap.HeapSortDesc(arr)
	} else {
		heap.HeapSort(arr)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":       true,
		"sortedArray":   arr,
		"originalArray": request.Array,
	})
}
