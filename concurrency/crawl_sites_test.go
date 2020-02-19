package concurency

import (
	"testing"
	"time"
)

func mockCrawlFn(url URL) bool {
	time.Sleep(5 * time.Millisecond)
	return true
}
func TestCheckWebsites(t *testing.T) {
	websites := []URL{
		"google.com",
		"yahoo.com",
		"cisco.com",
	}

	t.Run("Checking 3 websites simultaneously", func(t *testing.T) {
		healths := crawlSites(mockCrawlFn, websites)
		want := true
		for _, got := range healths {
			if got != want {
				t.Errorf("Got %t expected %t", got, want)
			}
		}
	})
}

func BenchmarkCheckWebsites(benchmark *testing.B) {
	websites := make([]URL, 100)
	for i := 0; i < 100; i++ {
		websites[i] = "random.com"
	}
	for i := 0; i < benchmark.N; i++ {
		crawlSites(mockCrawlFn, websites)
	}

	// Input size: 100 websites
	// Without go routines: 583441898 ns / op
	// With go routines   : 5725240   ns / op  (100 times faster)
}
