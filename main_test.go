package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name           string
		size           int
		wantLen        int
		checkPositive  bool
	}{
		{name: "size zero returns empty slice", size: 0, wantLen: 0, checkPositive: false},
		{name: "negative size returns empty slice", size: -5, wantLen: 0, checkPositive: false},
		{name: "size N returns slice of N positive ints", size: 100, wantLen: 100, checkPositive: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			require.NotNil(t, got)
			require.Equal(t, tt.wantLen, len(got))
			if tt.checkPositive {
				for i, v := range got {
					assert.Greater(t, v, 0, "index=%d", i)
				}
			}
		})
	}
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
