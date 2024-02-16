package depinjection

import (
	"bytes"
	"testing"
)

func TestDependencyInjection(t *testing.T) {
	buffer := bytes.Buffer{}
	Great(&buffer, "Ale")
	got := buffer.String()

	want := "Hello, Ale"

	if got != want {
		t.Errorf("Want %s diff than got %s", want, got)
	}
}
