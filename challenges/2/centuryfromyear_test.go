package challenges

import (
	"encoding/json"
	"testing"

	m "codesignal/models"
)

func data() []m.DataIntInt {
	return []m.DataIntInt{
		m.DataIntInt{Input: 1905, Output: 20},
		m.DataIntInt{Input: 1700, Output: 17},
		m.DataIntInt{Input: 1988, Output: 20},
		m.DataIntInt{Input: 2000, Output: 20},
		m.DataIntInt{Input: 2001, Output: 21},
		m.DataIntInt{Input: 200, Output: 2},
		m.DataIntInt{Input: 374, Output: 4},
		m.DataIntInt{Input: 45, Output: 1},
		m.DataIntInt{Input: 8, Output: 1},
	}
}

func TestCenturyFromYear(t *testing.T) {
	for _, v := range data() {
		output := CenturyFromYear(v.Input)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}

func BenchmarkCenturyFromYear(b *testing.B) {
	inputStrings := data()
	size := len(inputStrings)
	for i := 0; i < b.N; i++ {
		index := i % size
		CenturyFromYear(inputStrings[index].Input)
	}
}
