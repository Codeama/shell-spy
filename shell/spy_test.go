package shell

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	want := fmt.Sprintf("spy.go\nspy_test.go\n")
	got := Execute("ls")
	if got != want {
		t.Errorf("Expected: %s, got: %s", want, got)
	}
}
