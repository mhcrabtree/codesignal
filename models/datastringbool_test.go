package models

import (
	"encoding/json"
	"testing"
)

func TestDataStringBool(t *testing.T) {
	d := DataStringBool{Input: "aabbcc", Output: true}

	if !d.Assert(&d.Output) {
		j, _ := json.Marshal(d)
		t.Error(string(j))
	}
}

func BenchmarkDataStringBool(b *testing.B) {
	d := DataStringBool{Input: "aabbcc", Output: true}

	for i := 0; i < b.N; i++ {
		d.Assert(&d.Output)
	}
}
