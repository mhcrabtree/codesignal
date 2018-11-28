package models

import (
	"encoding/json"
	"testing"
)

func TestDataIntBool(t *testing.T) {
	d := DataIntBool{Input: 123, Output: true}

	if !d.Assert(&d.Output) {
		j, _ := json.Marshal(d)
		t.Error(string(j))
	}
}

func BenchmarkDataIntBool(b *testing.B) {
	d := DataIntBool{Input: 123, Output: true}

	for i := 0; i < b.N; i++ {
		d.Assert(&d.Output)
	}
}
