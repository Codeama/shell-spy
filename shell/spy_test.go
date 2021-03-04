package shell

import (
	"reflect"
	"testing"
)

// func TestExecute(t *testing.T) {
// 	var out bytes.Buffer
// 	want := "spy.go\nspy_test.go\n"
// 	Execute(&out, "ls")
// 	got := fmt.Sprint(&out)
// 	if got != want {
// 		t.Errorf("Expected: %s, got: %s", want, got)
// 	}
// }

func TestParseCommand(t *testing.T) {
	testData := struct {
		cmd  string
		name string
		args []string
	}{
		"echo Bukola",
		"echo",
		[]string{"Bukola"},
	}
	name, args := ParseCommand(testData.cmd)

	if testData.name != name {
		t.Errorf("Expected command name: %s, got: %s", testData.name, name)
	}
	if !reflect.DeepEqual(testData.args, args) {
		t.Errorf("Expected args: %s, got %s", testData.args, args)
	}

}
