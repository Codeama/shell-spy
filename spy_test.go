package spy

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestParseCommand(t *testing.T) {
	tcs := []struct {
		cmd         string
		name        string
		args        []string
		errExpected bool
	}{
		{"echo Bukola", "echo", []string{"Bukola"}, false},
		{"test it works fine", "test", []string{"it", "works", "fine"}, false},
		{"cat /file-path", "cat", []string{"/file-path"}, false},
		{"cd ..", "cd", []string{".."}, false},
		{"echo", "echo", []string{}, false},
		{"echo'", "echo'", []string{}, true},
		{" ", " ", []string{}, false},
		{"   ", "   ", []string{}, false},
		{"echo 'Hello, 123'", "echo", []string{"Hello, 123"}, false},
		{"sh -c 'ls", "BOGUS", []string{}, true},
	}

	for _, tc := range tcs {

		name, args, err := ParseCommand(tc.cmd)

		errorReturned := (err != nil)

		if errorReturned != tc.errExpected {
			t.Fatalf("ParseCommand(%q): Unexpected error status: %v", tc.cmd, err)
		}

		if !errorReturned && tc.name != name {
			t.Errorf("ParseCommand(%q): Expected command name to be %q, got: %q", tc.cmd, tc.name, name)
		}
		if !errorReturned && !cmp.Equal(tc.args, args) {
			t.Errorf("ParseCommand(%q): Expected args to be %q, got %q", tc.cmd, tc.args, args)
		}
	}

}

func TestTimestamp(t *testing.T) {
	wantTimestamp := "Log time: 2020-04-04T20:15:02Z\n\n"
	currentTime := time.Date(2020, time.April, 4, 20, 15, 02, 00, time.UTC)

	fakeRecorder := bytes.Buffer{}

	s, _ := NewSession("")
	s.Recorder = &fakeRecorder
	s.RecordTime(currentTime)
	got := fakeRecorder.String()
	if !cmp.Equal(wantTimestamp, got) {
		t.Error(cmp.Diff(wantTimestamp, got))
	}
}
