package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntArrayInt {
	return []m.DataIntArrayInt{
		m.DataIntArrayInt{Input: []int{3, 5, 6, 2, 1}, Output: 683},
		m.DataIntArrayInt{Input: []int{3, 5, 6, 2, 1, 9}, Output: 1583},
		m.DataIntArrayInt{Input: []int{3, 5, 6, 2}, Output: 115},
		m.DataIntArrayInt{Input: []int{5, 6, 2}, Output: 67},
		m.DataIntArrayInt{Input: []int{3, 5}, Output: 8},
		m.DataIntArrayInt{Input: []int{1, 1, 1, 1, 1, 1}, Output: 222},
		m.DataIntArrayInt{Input: []int{1, 1, 1, 1, 1}, Output: 122},
		m.DataIntArrayInt{Input: []int{0, 0}, Output: 0},
		m.DataIntArrayInt{Input: []int{0}, Output: 0},
		m.DataIntArrayInt{Input: []int{3}, Output: 3},
		m.DataIntArrayInt{Input: []int{9, 4, 2, 7, 9, 0}, Output: 1912},
		m.DataIntArrayInt{Input: []int{3, 2, 1, 3, 2, 1, 3, 2, 1}, Output: 36432},
		m.DataIntArrayInt{Input: []int{9, 4, 2, 7, 9, 0, 1}, Output: 10661},
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

func BenchmarkDigits(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		Digits(inputStrings[index].Input)
	}
}
