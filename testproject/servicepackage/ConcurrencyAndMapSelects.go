package servicepackage

import (
	"net/http"
	"time"
)

// This struct will use for channel
type Result struct {
	bool
}

func SearchElement(elements map[string]string, serachkey string) string {
	return elements[serachkey]
}

func CheckWebsites(urls []string) map[string]bool {
	results := make(map[string]bool)
	resultchan := make(chan Result)

	for i := 0; i < len(urls); i++ {
		// Using anonymous functions when we want to start a goroutine
		go func() {
			resultchan <- Result{checkUrl(urls[i])}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultchan
		results[urls[i]] = r.bool
	}
	time.Sleep(2 * time.Second)
	return results
}

func checkUrl(url string) bool {
	return url == "https://www.geeksforgeeks.org/"
}

func Racer(slowURL, fastURL string) (winner string) {
	slowDuration := measureResponseTime(slowURL)
	fastDuration := measureResponseTime(fastURL)

	if slowDuration < fastDuration {
		return slowURL
	}
	return fastURL
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// synchronise processes really easily and clearly
func RacerSelect(slowURL, fastURL string) (winner string) {
	select {
	case <-ping(slowURL):
		return slowURL
	case <-ping(fastURL):
		return fastURL
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
