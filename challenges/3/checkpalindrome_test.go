package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataStringBool {
	return []m.DataStringBool{
		m.DataStringBool{Input: "aabaa", Output: true},
		m.DataStringBool{Input: "abac", Output: false},
		m.DataStringBool{Input: "a", Output: true},
		m.DataStringBool{Input: "az", Output: false},
		m.DataStringBool{Input: "z", Output: true},
		m.DataStringBool{Input: "aaabaaaa", Output: false},
		m.DataStringBool{Input: "zzzazzazz", Output: false},
		m.DataStringBool{Input: "", Output: true},
		m.DataStringBool{Input: "hlbeeykoqqqqokyeeblh", Output: true},
		m.DataStringBool{Input: "abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba", Output: true},
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
