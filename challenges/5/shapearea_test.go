package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntInt {
	return []m.DataIntInt{
		m.DataIntInt{Input: 2, Output: 5},
		m.DataIntInt{Input: 3, Output: 13},
		m.DataIntInt{Input: 1, Output: 1},
		m.DataIntInt{Input: 5, Output: 41},
	}
}

func TestShapeArea(t *testing.T) {
	for _, v := range data() {
		output := ShapeArea(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}

func BenchmarkShapeArea(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		ShapeArea(inputStrings[index].Input)
	}
}
