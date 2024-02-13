package iteration

import "testing"

func TestLoop(t *testing.T) {
	repeat := Repeat("L")
	expected := "LLLLLLL"

	if repeat != expected {
		t.Errorf("Expect was %q, got %q", expected, repeat)
	}
}
