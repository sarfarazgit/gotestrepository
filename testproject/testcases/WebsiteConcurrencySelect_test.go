package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"testproject/servicepackage"
	"time"
)

func TestWebsites(t *testing.T) {

	websites := []string{"https://www.geeksforgeeks.org/",
		"http://www.geeksforgeeks.org/", "http://www.geeksforgeeks.com/"}

	actualresult := servicepackage.CheckWebsites(websites)

	expectedresult := map[string]bool{"https://www.geeksforgeeks.org/": true,
		"http://www.geeksforgeeks.org/": false, "http://www.geeksforgeeks.com/": false}

	if reflect.DeepEqual(actualresult, expectedresult) {
		t.Fatalf("wanted %v, got %v", expectedresult, actualresult)
	}

}

func TestRaceCondiition(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expected := fastURL
	//result := servicepackage.Racer(slowURL, fastURL)
	result := servicepackage.RacerSelect(slowURL, fastURL)

	if expected != result {
		t.Errorf("result %q, expected %q", result, expected)
	}
}

// Mock Http server
func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
