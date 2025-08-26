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
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}

	result := make([]int, size)
	for i := range result {
		result[i] = rand.Intn(1_000_000) + 1 // positive integers
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}
	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	chunkSize := 0
	if CHUNKS > 0 {
		chunkSize = len(data) / CHUNKS
	}

	maxima := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data) // include remainder in the last chunk
		}
		if start > len(data) {
			start = len(data)
		}
		if end > len(data) {
			end = len(data)
		}

		idx := i
		segment := data[start:end]
		go func() {
			defer wg.Done()
			maxima[idx] = maximum(segment)
		}()
	}
	wg.Wait()

	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	rand.Seed(time.Now().UnixNano())
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
