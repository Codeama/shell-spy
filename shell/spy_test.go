package shell

import (
	"reflect"
	"testing"
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
	}

	for _, td := range testData {

		name, args := ParseCommand(td.cmd)

		if td.name != name {
			t.Errorf("Expected command name: %s, got: %s", td.name, name)
		}
		if !reflect.DeepEqual(td.args, args) {
			t.Errorf("Expected args: %s, got %s", td.args, args)
		}
	}

}
