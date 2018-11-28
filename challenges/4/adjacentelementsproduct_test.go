package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntArrayInt {
	return []m.DataIntArrayInt{
		m.DataIntArrayInt{Input: []int{3, 6, -2, -5, 7, 3}, Output: 21},
		m.DataIntArrayInt{Input: []int{-1, -2}, Output: 2},
		m.DataIntArrayInt{Input: []int{5, 1, 2, 3, 1, 4}, Output: 6},
		m.DataIntArrayInt{Input: []int{1, 2, 3, 0}, Output: 6},
		m.DataIntArrayInt{Input: []int{9, 5, 10, 2, 24, -1, -48}, Output: 50},
		m.DataIntArrayInt{Input: []int{5, 6, -4, 2, 3, 2, -23}, Output: 30},
		m.DataIntArrayInt{Input: []int{4, 1, 2, 3, 1, 5}, Output: 6},
		m.DataIntArrayInt{Input: []int{-23, 4, -3, 8, -12}, Output: -12},
		m.DataIntArrayInt{Input: []int{1, 0, 1, 0, 1000}, Output: 0},
	}
}

func TestAdjacentElementsProduct(t *testing.T) {
	for _, v := range data() {
		output := AdjacentElementsProduct(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}

func BenchmarkAdjacentElementsProduct(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		AdjacentElementsProduct(inputStrings[index].Input)
	}
}
