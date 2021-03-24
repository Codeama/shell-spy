//+build integration

package spy_test

import (
	"bytes"
	"spy"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExecuteIntegration(t *testing.T) {
	expectedTerminal := "Hello\n"
	expectedRecorder := "************ START ************\nCOMMAND: echo Hello \n\nOUTPUT:\nHello\n"

	var terminal, recorder bytes.Buffer
	err := spy.Execute(&terminal, &recorder, "echo Hello")

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(expectedTerminal, terminal.String()) {
		t.Errorf("Want: %q, got: %q", expectedRecorder, &terminal)
	}

	if !cmp.Equal(expectedRecorder, recorder.String()) {
		t.Errorf("Want: %q, got: %q", expectedRecorder, &terminal)
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
