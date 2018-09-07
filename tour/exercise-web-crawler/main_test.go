// Run the benchmark and dig into the performance issue
// go test -bench . -cpuprofile [fileName]
// go tool pprof [fileName]
package exercise

import (
	"testing"
)

func BenchmarkCrawler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Crawl("https://golang.org/", 4, fetcher)
	}
}
