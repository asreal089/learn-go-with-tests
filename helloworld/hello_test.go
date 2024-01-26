package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello()

	want := "Hello there!!"

	if got != want {
		t.Errorf("Want %q different from got %q", want, got)
	}
}
