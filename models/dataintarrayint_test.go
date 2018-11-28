package models

import (
	"encoding/json"
	"testing"
)

func TestDataIntArrayInt(t *testing.T) {
	d := DataIntArrayInt{Input: []int{3, 6, -2, -5, 7, 3}, Output: 21}

	if !d.Assert(&d.Output) {
		j, _ := json.Marshal(d)
		t.Error(string(j))
	}
}

func BenchmarkDataIntArrayInt(b *testing.B) {
	d := DataIntArrayInt{Input: []int{3, 6, -2, -5, 7, 3}, Output: 21}
	for i := 0; i < b.N; i++ {
		d.Assert(&d.Output)
	}
}
