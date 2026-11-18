package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		// Генерируем числа от 1 до 10000 (положительные)
		elements[i] = rand.Intn(10000) + 1
	}
	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	// Если данных меньше, чем частей, обрабатываем в одном потоке
	if chunkSize == 0 {
		return maximum(data)
	}

	// Слайс для хранения максимумов каждой части
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Последний чанк должен захватить все оставшиеся элементы
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(start, end, index int) {
			defer wg.Done()
			// Находим максимум в текущем срезе и записываем в общий слайс
			maxValues[index] = maximum(data[start:end])
		}(start, end, i)
	}

	wg.Wait()

	// Находим максимум среди максимумов частей
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)}
	