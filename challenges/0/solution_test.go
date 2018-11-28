package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntArrayInt {
	return []m.DataIntArrayInt{
		m.DataIntArrayInt{Input: []int{1, 3, 6, 4, 1, 2}, Output: 5},
		m.DataIntArrayInt{Input: []int{1, 2, 3}, Output: 4},
		m.DataIntArrayInt{Input: []int{-1, -3}, Output: 1},
	}
}

func TestSolution(t *testing.T) {
	for _, v := range data() {
		output := Solution(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}
