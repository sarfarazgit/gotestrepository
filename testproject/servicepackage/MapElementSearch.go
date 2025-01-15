package servicepackage

import "time"

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
