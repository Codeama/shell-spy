//+build integration

package spy_test

import (
	"bytes"
	"spy"
	"testing"
)

func TestExecuteIntegration(t *testing.T) {
	var buff bytes.Buffer
	want := "echo Hello\nHello\n"
	err := spy.Execute(&buff, "echo Hello")

	if err != nil {
		t.Error(err)
	}

	got := buff.String()
	if want != got {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}

func TestExecuteIntegrationInvalid(t *testing.T) {
	var buff bytes.Buffer
	want := "echo Hello'\nError parsing command\n"
	err := spy.Execute(&buff, "echo Hello'")

	if err != nil {
		t.Error(err)
	}

	got := buff.String()
	if want != got {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}
