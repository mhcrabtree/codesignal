package models

import (
	"encoding/json"
	"testing"
)

func TestDataIntArrayBool(t *testing.T) {
	d := DataIntArrayBool{Input: []int{3, 6, -2, -5, 7, 3}, Output: true}

	if !d.Assert(&d.Output) {
		j, _ := json.Marshal(d)
		t.Error(string(j))
	}
}

func BenchmarkDataIntArrayBool(b *testing.B) {
	d := DataIntArrayBool{Input: []int{3, 6, -2, -5, 7, 3}, Output: true}
	for i := 0; i < b.N; i++ {
		d.Assert(&d.Output)
	}
}
