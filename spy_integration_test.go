//+build integration

package spy_test

import (
	"bytes"
	"spy"
	"testing"
)

func TestExecuteIntegration(t *testing.T) {
	want := struct{ a, b string }{
		a: "Hello\n",
		b: "************ START ************\nCOMMAND: echo Hello\n\nOUTPUT:\nHello\n",
	}
	var buff, buff2 bytes.Buffer
	err := spy.Execute(&buff, &buff2, "echo Hello")

	if err != nil {
		t.Error(err)
	}
	got := struct{ a, b string }{
		a: buff.String(),
		b: buff2.String(),
	}

	if want.a != got.a {
		t.Errorf("Want: %q, got: %q", want.a, got.a)
	}

	if want.b != got.b {
		t.Errorf("Want: %q, got: %q", want.b, got.b)
	}
}

func TestExecuteIntegrationInvalid(t *testing.T) {
	var buff, buff2 bytes.Buffer
	want := "Error parsing command\n"
	err := spy.Execute(&buff, &buff2, "echo Hello'")

	if err != nil {
		t.Error(err)
	}

	got := buff.String()
	if want != got {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}
