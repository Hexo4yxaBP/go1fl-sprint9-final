package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("size zero returns empty slice", func(t *testing.T) {
		got := generateRandomElements(0)
		require.NotNil(t, got)
		assert.Equal(t, 0, len(got))
	})

	t.Run("negative size returns empty slice", func(t *testing.T) {
		got := generateRandomElements(-5)
		require.NotNil(t, got)
		assert.Equal(t, 0, len(got))
	})

	t.Run("size N returns slice of N positive ints", func(t *testing.T) {
		const n = 100
		got := generateRandomElements(n)
		require.Equal(t, n, len(got))
		for i, v := range got {
			assert.Greater(t, v, 0, "index=%d", i)
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slice returns 0", func(t *testing.T) {
		assert.Equal(t, 0, maximum([]int{}))
	})

	t.Run("single element slice returns that element", func(t *testing.T) {
		require.Equal(t, 42, maximum([]int{42}))
	})

	t.Run("multiple elements returns correct maximum", func(t *testing.T) {
		data := []int{1, 3, 2, 8, 5, 7}
		assert.Equal(t, 8, maximum(data))
	})

	t.Run("all equal elements returns that value", func(t *testing.T) {
		data := []int{5, 5, 5, 5}
		assert.Equal(t, 5, maximum(data))
	})

	t.Run("decreasing order returns first element", func(t *testing.T) {
		data := []int{9, 7, 6, 3}
		assert.Equal(t, 9, maximum(data))
	})
}
