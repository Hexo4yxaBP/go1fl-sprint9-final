package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("size zero returns empty slice", func(t *testing.T) {
		got := generateRandomElements(0)
		if len(got) != 0 {
			t.Fatalf("expected empty slice, got len=%d", len(got))
		}
	})

	t.Run("negative size returns empty slice", func(t *testing.T) {
		got := generateRandomElements(-5)
		if len(got) != 0 {
			t.Fatalf("expected empty slice for negative size, got len=%d", len(got))
		}
	})

	t.Run("size N returns slice of N positive ints", func(t *testing.T) {
		const n = 100
		got := generateRandomElements(n)
		if len(got) != n {
			t.Fatalf("expected len=%d, got %d", n, len(got))
		}
		for i, v := range got {
			if v <= 0 {
				t.Fatalf("expected positive integers, got got[%d]=%d", i, v)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slice returns 0", func(t *testing.T) {
		if max := maximum([]int{}); max != 0 {
			t.Fatalf("expected 0 for empty slice, got %d", max)
		}
	})

	t.Run("single element slice returns that element", func(t *testing.T) {
		if max := maximum([]int{42}); max != 42 {
			t.Fatalf("expected 42, got %d", max)
		}
	})

	t.Run("multiple elements returns correct maximum", func(t *testing.T) {
		data := []int{1, 3, 2, 8, 5, 7}
		if max := maximum(data); max != 8 {
			t.Fatalf("expected 8, got %d", max)
		}
	})

	t.Run("all equal elements returns that value", func(t *testing.T) {
		data := []int{5, 5, 5, 5}
		if max := maximum(data); max != 5 {
			t.Fatalf("expected 5, got %d", max)
		}
	})

	t.Run("decreasing order returns first element", func(t *testing.T) {
		data := []int{9, 7, 6, 3}
		if max := maximum(data); max != 9 {
			t.Fatalf("expected 9, got %d", max)
		}
	})
}
