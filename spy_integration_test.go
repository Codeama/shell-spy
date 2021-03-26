//+build integration

package spy_test

import (
	"bytes"
	"spy"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExecuteIntegration(t *testing.T) {
	wantTerminal := "Hello\n"
	wantRecorder := "************ START ************\nCOMMAND: echo Hello \n\nOUTPUT:\nHello\n"

	var terminal, recorder bytes.Buffer
	err := spy.Execute(&terminal, &recorder, "echo Hello")

	if err != nil {
		t.Error(err)
	}
	gotTerminal := terminal.String()
	if !cmp.Equal(wantTerminal, gotTerminal) {
		t.Error(cmp.Diff(wantTerminal, gotTerminal))
	}
	gotRecorder := recorder.String()
	if !cmp.Equal(wantRecorder, gotRecorder) {
		t.Error(cmp.Diff(wantRecorder, gotRecorder))
	}
}

func TestExecuteIntegrationInvalid(t *testing.T) {
	var terminal, recorder bytes.Buffer
	want := "Error parsing command\n"
	err := spy.Execute(&terminal, &recorder, "echo Hello'")

	if err != nil {
		t.Error(err)
	}

	got := terminal.String()
	if want != got {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}
