package factorial

import (
	"testing"
)

func TestCountingRecursive(t *testing.T) {
	if res := countingRecursive(5); res != 120 {
		t.Error("Expected 120, but", res)
	}
}
