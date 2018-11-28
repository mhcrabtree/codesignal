package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntArrayBool {
	return []m.DataIntArrayBool{
		m.DataIntArrayBool{Input: []int{1, 3, 2, 1}, Output: false},
		m.DataIntArrayBool{Input: []int{1, 3, 2}, Output: true},
		m.DataIntArrayBool{Input: []int{1, 2, 1, 2}, Output: false},
		m.DataIntArrayBool{Input: []int{1, 4, 10, 4, 2}, Output: false},
		m.DataIntArrayBool{Input: []int{10, 1, 2, 3, 4, 5}, Output: true},
		m.DataIntArrayBool{Input: []int{1, 1, 1, 2, 3}, Output: false},
		m.DataIntArrayBool{Input: []int{0, -2, 5, 6}, Output: true},
		m.DataIntArrayBool{Input: []int{1, 2, 3, 4, 5, 3, 5, 6}, Output: false},
		m.DataIntArrayBool{Input: []int{40, 50, 60, 10, 20, 30}, Output: false},
		m.DataIntArrayBool{Input: []int{1, 1}, Output: true},
		m.DataIntArrayBool{Input: []int{1, 2, 5, 3, 5}, Output: true},
		m.DataIntArrayBool{Input: []int{10, 1, 2, 3, 4, 5, 6, 1}, Output: true},
		m.DataIntArrayBool{Input: []int{1, 2, 3, 4, 3, 6}, Output: true},
	}
}

func TestCheckPalindrome(t *testing.T) {
	for _, v := range data() {
		output := CheckPalindrome(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}

func BenchmarkCheckPalendrome(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		CheckPalindrome(inputStrings[index].Input)
	}
}
