package challenges

import (
	"encoding/json"
	"testing"
)

func TestDigitsMod2(t *testing.T) {
	for _, v := range data() {
		output := DigitsMod2(v.Input, 2)
		if !v.Assert(&output) {
			f, _ := json.Marshal(v)
			t.Error(string(f))
		}
	}
}
