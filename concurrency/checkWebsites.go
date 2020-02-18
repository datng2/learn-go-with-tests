package concurency

// Task:
// Here's the setup: a colleague has written a function, CheckWebsites,
// that checks the status of a list of URLs.
//
// return `true` if website is still accessible, `false` otherwise.
func checkWebsites(urls []string) {

	for _, url := range urls {
		ok := healthCheck(url)
	}
}
