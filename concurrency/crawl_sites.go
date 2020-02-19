package concurency

type URL string

type CrawlHandler func(URL) bool

type CrawlResult struct {
	website URL
	data    bool
}

// Task:
// Here's the setup: a colleague has written a function, CheckWebsites,
// that checks the status of a list of URLs.
//
// return `true` if website is still accessible, `false` otherwise.
func crawlSites(crawler CrawlHandler, websites []URL) map[URL]bool {
	results := map[URL]bool{}
	channel := make(chan CrawlResult)

	for _, url := range websites {
		// results[url] = crawler(url)
		go func(website URL) {
			channel <- CrawlResult{website, crawler(website)}
		}(url)
	}

	for i := 0; i < len(websites); i++ {
		res := <-channel
		results[res.website] = res.data
	}
	return results
}

// 23406615 ns / op
// 2249228637 ns / op

// func crawlFn(url URL) bool {
// 	time.Sleep(5 * time.Millisecond)
// 	return true
// }
