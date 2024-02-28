package concurrency

import (
	"testing"
	"time"
)

func slowWebSiteBenchMark(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func Benchmark(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebSiteBenchMark, urls)
	}
}
