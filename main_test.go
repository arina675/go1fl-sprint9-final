package main

import (
	"testing"
)

// Тесты для функции generateRandomElements
func TestGenerateRandomElements(t *testing.T) {
	t.Run("Нулевой размер", func(t *testing.T) {
		res := generateRandomElements(0)
		if len(res) != 0 {
			t.Errorf("Ожидался пустой слайс, получено %d элементов", len(res))
		}
	})

	t.Run("Отрицательный размер", func(t *testing.T) {
		res := generateRandomElements(-5)
		if len(res) != 0 {
			t.Errorf("Ожидался пустой слайс, получено %d элементов", len(res))
		}
	})

	t.Run("Корректный размер", func(t *testing.T) {
		size := 100
		res := generateRandomElements(size)
		if len(res) != size {
			t.Errorf("Ожидался размер %d, получено %d", size, len(res))
		}
		for _, v := range res {
			if v <= 0 {
				t.Errorf("Все элементы должны быть положительными, получено %d", v)
			}
		}
	})
}

// Тесты для функции maximum
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Один элемент", []int{42}, 42},
		{"Все элементы одинаковые", []int{5, 5, 5, 5}, 5},
		{"Максимум в начале", []int{100, 2, 3, 4}, 100},
		{"Максимум в конце", []int{1, 2, 3, 100}, 100},
		{"Максимум в середине", []int{1, 2, 100, 3, 4}, 100},
		{"С отрицательными числами", []int{-10, -5, -20}, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

// Тесты для функции maxChunks
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Один элемент", []int{42}, 42},
		{"Меньше 8 элементов", []int{10, 5, 20, 3}, 20},
		{"Ровно 8 элементов", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"Больше 8 элементов", []int{1, 20, 3, 40, 5, 60, 7, 80, 9, 100, 11, 120, 13, 140, 15, 160}, 160},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}