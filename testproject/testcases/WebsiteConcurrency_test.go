package main

import (
	"reflect"
	"testing"
	"testproject/servicepackage"
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
