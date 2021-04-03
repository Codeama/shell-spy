//+build integration

package spy_test

import (
	"bytes"
	"fmt"
	"spy"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExecuteIntegration(t *testing.T) {
	commandLine := "echo Hello, World!"
	wantRecorder := fmt.Sprintf("COMMAND: %s \n\nOUTPUT:\nHello, World!\n", commandLine)
	wantTerminal := "Hello, World!\n"

	s, err := spy.NewSession("/tmp/test-file.log")
	if err != nil {
		t.Fatal(err)
	}

	var logFile, terminal bytes.Buffer
	s.Recorder = &logFile
	s.Terminal = &terminal

	err = s.Execute(commandLine)
	if err != nil {
		t.Fatal(err)
	}

	gotRecorder := logFile.String()
	gotTerminal := terminal.String()

	if !cmp.Equal(wantRecorder, gotRecorder) {
		t.Error(cmp.Diff(wantRecorder, gotRecorder))
	}

	if !cmp.Equal(wantTerminal, gotTerminal) {
		t.Error(cmp.Diff(wantTerminal, gotTerminal))
	}
}

func TestExecuteIntegrationInvalid(t *testing.T) {
	var logFile, terminal bytes.Buffer

	want := "Error parsing command\n"
	s, err := spy.NewSession("/tmp/test-file.log")

	if err != nil {
		t.Error(err)
	}
	s.Terminal = &terminal
	s.Recorder = &logFile
	err = s.Execute("echo Hello'")

	if err != nil {
		t.Error(err)
	}

	got := terminal.String()
	if want != got {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}
