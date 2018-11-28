package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntArrayInt {
	return []m.DataIntArrayInt{
		m.DataIntArrayInt{Input: []int{9, 8, 8, 5, 3, 5, 3, 2, 8, 6}, Output: 4},
		m.DataIntArrayInt{Input: []int{4, 4}, Output: 1},
		m.DataIntArrayInt{Input: []int{1}, Output: 1},
		m.DataIntArrayInt{Input: []int{2, 1, 4, 4, 1, 4, 4, 1, 2, 0, 1, 0, 0, 3, 1, 3, 4, 1, 3, 4}, Output: 6},
	}
}

func TestZigZag(t *testing.T) {
	for _, v := range data() {
		output := ZigZag(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}

func BenchmarkZigZag(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		ZigZag(inputStrings[index].Input)
	}
}
