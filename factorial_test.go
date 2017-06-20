package factorial

import (
	"testing"
)

func TestCountingRecursive(t *testing.T) {
	if res := countingRecursive(5); res != 120 {
		t.Error("Expected 120, but", res)
	}

	if res := counting(6); res != 720 {
		t.Error("Expected 720, but", res)
	}
}
