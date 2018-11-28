package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntArrayInt {
	return []m.DataIntArrayInt{
		m.DataIntArrayInt{Input: []int{6, 2, 3, 8}, Output: 3},
		m.DataIntArrayInt{Input: []int{0, 3}, Output: 2},
		m.DataIntArrayInt{Input: []int{5, 4, 6}, Output: 0},
		m.DataIntArrayInt{Input: []int{6, 3}, Output: 2},
		m.DataIntArrayInt{Input: []int{1}, Output: 0},
	}
}

func TestMakeArrayConsecutive(t *testing.T) {
	for _, v := range data() {
		output := MakeArrayConsecutive(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}

func BenchmarkMakeArrayConsecutive(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		MakeArrayConsecutive(inputStrings[index].Input)
	}
}
