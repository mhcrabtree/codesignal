package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntInt {
	return []m.DataIntInt{
		m.DataIntInt{Input: 9, Output: 2},
		m.DataIntInt{Input: 529, Output: 4},
		m.DataIntInt{Input: 20, Output: 1},
		m.DataIntInt{Input: 15, Output: 0},
		m.DataIntInt{Input: 1041, Output: 5},
		m.DataIntInt{Input: 32, Output: 0},
		m.DataIntInt{Input: 0, Output: 0},
		m.DataIntInt{Input: 1, Output: 0},
	}
}

func TestBinaryGap(t *testing.T) {
	for _, v := range data() {
		output := BinaryGap(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}

func BenchmarkBinaryGap(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		BinaryGap(inputStrings[index].Input)
	}
}
