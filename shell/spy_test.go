package shell

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseCommand(t *testing.T) {
	tcs := []struct {
		cmd  string
		name string
		args []string
	}{
		{"echo Bukola", "echo", []string{"Bukola"}},
		{"test it works fine", "test", []string{"it", "works", "fine"}},
		{"cat /file-path", "cat", []string{"/file-path"}},
		{"cd ..", "cd", []string{".."}},
		{"echo", "echo", []string{}},
		{"echo'", "echo'", []string{}},
		{" ", " ", []string{}},
		{"   ", "   ", []string{}},
		{"echo 'Hello, 123'", "echo", []string{"Hello, 123"}},
	}

	for _, tc := range tcs {

		name, args := ParseCommand(tc.cmd)

		if tc.name != name {
			t.Errorf("ParseCommand(%q): Expected command name to be %q, got: %q", tc.cmd, tc.name, name)
		}
		if !cmp.Equal(tc.args, args) {
			t.Errorf("ParseCommand(%q): Expected args to be %q, got %q", tc.cmd, tc.args, args)
		}
	}

}
