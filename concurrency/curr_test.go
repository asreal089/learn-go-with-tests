package concurrency

import (
	"reflect"
	"testing"
)

func mockWebSiteCheck(url string) bool {
	if url == "waat://futher" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"waat://futher",
	}

	want := map[string]bool{
		"https://google.com":   true,
		"https://facebook.com": true,
		"waat://futher":        false,
	}

	got := CheckWebsites(mockWebSiteCheck, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got diff then want")
	}
}
