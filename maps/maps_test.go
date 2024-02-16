package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "its a test"}

	got := Search(dict, "test")
	want := "its a test"
	if got != want {
		t.Errorf("got diff than want")
	}
}
