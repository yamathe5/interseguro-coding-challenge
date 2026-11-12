package matrix

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	for _, raw := range []string{"null", "[]", "[[]]", "[[1],[2,3]]", "[[null]]", "[[\"1\"]]", "[[1e309]]", "[[" + strings.Repeat("0,", 100) + "0]]", "[" + strings.Repeat("[0],", 100) + "[0]]"} {
		if _, err := Decode(json.RawMessage(raw)); err == nil {
			t.Errorf("accepted %s", raw)
		}
	}
	if value, err := Decode(json.RawMessage("[[1,-2,0.25]]")); err != nil || value[0][2] != 0.25 {
		t.Fatalf("%v %v", value, err)
	}
}
