package models

import (
	"encoding/json"
	"testing"
)

func TestDataIntInt(t *testing.T) {
	d := DataIntInt{Input: 1905, Output: 20}

	if !d.Assert(&d.Output) {
		j, _ := json.Marshal(d)
		t.Error(string(j))
	}
}

func BenchmarkDataIntInt(b *testing.B) {
	d := DataIntInt{Input: 1905, Output: 20}

	for i := 0; i < b.N; i++ {
		d.Assert(&d.Output)
	}
}
