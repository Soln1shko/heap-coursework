# Реализация и визуализация кучи на Go

Курсовая работа по теме "Кучи. Представления кучи в компьютере, алгоритм окучивания массива"

## Описание проекта

Данный проект представляет собой полную реализацию структуры данных "куча" (heap) на языке Go с визуализацией, тестами и примерами использования. Реализованы как минимальная (MinHeap), так и максимальная (MaxHeap) кучи.

## Структура проекта

```
heap-coursework/
├── cmd/
│   └── web/
│       ├── main.go              # HTTP-сервер и маршруты
│       └── templates/
│           └── index.html       # HTML-шаблон интерфейса
├── docs/                        # Теория
│   ├── algorithms.md
│   └── complexity-analysis.md
├── examples/
│   ├── basic/basic-usage.go
│   ├── sorting/sorting-demo.go
│   └── performance/performance-test.go
├── internal/
│   ├── heap/                    # Реализация кучи
│   │   ├── heap.go
│   │   ├── heapify.go
│   │   ├── heap-sort.go
│   │   └── heap-types.go
│   ├── utils/                   # Вспомогательные функции
│   │   ├── utils.go
│   │   └── array-representation.go
│   ├── visualization/           
│   │   └── visualization.go
│   └── benchmark/               
│       └── benchmark.go
├── pkg/tests/                   # Юнит-тесты
│   └── tests.go
├── go.mod
├── go.sum
└── work.pdf                     # Курсач
```

## Установка и запуск

### Требования

- Go 1.16 или выше
- Для визуализации: библиотека `github.com/fogleman/gg`

### Установка зависимостей

```bash
go mod download
```

### Запуск примеров

```bash
# Базовое использование кучи
go run examples/basic/basic-usage.go

# Демонстрация сортировки кучей
go run examples/sorting/sorting-demo.go

# Тесты производительности
go run examples/performance/performance-test.go
```

### Запуск веб-интерфейса

```bash
go run cmd/web/main.go
```

После запуска веб-интерфейс будет доступен по адресу: http://localhost:8080

## Документация

Сгенерированная с помощью ИИ документация доступна в директории `docs/` ^-^:
- `algorithms.md` - описание алгоритмов работы с кучей
- `complexity-analysis.md` - анализ временной и пространственной сложности алгоритмов

## Курсовая работа

📄 [Текст работы](./work.pdf)

