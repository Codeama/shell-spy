package shell

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseCommand(t *testing.T) {
	testData := []struct {
		cmd  string
		name string
		args []string
	}{
		{"echo Bukola", "echo", []string{"Bukola"}},
		{"test it works fine", "test", []string{"it", "works", "fine"}},
		{"cat /file-path", "cat", []string{"/file-path"}},
		{"cd ..", "cd", []string{".."}},
		{"echo", "echo", []string{}},
		{" ", "", []string{""}},
		{"   ", "", []string{"", "", ""}},
	}

	for _, td := range testData {

		name, args := ParseCommand(td.cmd)

		if td.name != name {
			t.Errorf("ParseCommand(%q): Expected command name to be %q, got: %q", td.cmd, td.name, name)
		}
		if !cmp.Equal(td.args, args) {
			t.Errorf("ParseCommand(%q): Expected args to be %q, got %q", td.cmd, td.args, args)
		}
	}

}
