package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntArrayInt {
	return []m.DataIntArrayInt{
		//		m.DataIntArrayInt{Input: []int{3, 5, 6, 2, 1}, Output: 683},
		m.DataIntArrayInt{Input: []int{3, 5, 6, 2, 1, 9}, Output: 1583},
		//		m.DataIntArrayInt{Input: []int{3, 5, 6, 2}, Output: 115},
		//		m.DataIntArrayInt{Input: []int{5, 6, 2}, Output: 67},
		//		m.DataIntArrayInt{Input: []int{3, 5}, Output: 8},
		//		m.DataIntArrayInt{Input: []int{1, 1, 1, 1, 1, 1}, Output: 222},
		//		m.DataIntArrayInt{Input: []int{1, 1, 1, 1, 1}, Output: 122},
		//		m.DataIntArrayInt{Input: []int{0, 0}, Output: 0},
		//		m.DataIntArrayInt{Input: []int{0}, Output: 0},
		//		m.DataIntArrayInt{Input: []int{3}, Output: 3},
	}
}

func TestDigits(t *testing.T) {
	for _, v := range data() {
		output := Digits(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}
