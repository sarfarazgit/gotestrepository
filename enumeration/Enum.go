package main

import "fmt"

type Status int
type Tracker string

const (
	CANCELLED Tracker = "CANCELLED"
	DELEVERED Tracker = "DELEVERED"
)

const (
	SUCCESS Status = iota
	FAILED
	PENDING
)

func showstatus(sval Status) {
	fmt.Println("Status value is::", sval)
}

func showstracker(tval Tracker) {
	fmt.Println("Tracker value is::", tval)
}

func main() {

	// Data preparing for status
	showstatus(FAILED)
	showstatus(SUCCESS)
	showstatus(PENDING)

	// Data preparation for Tracker

	showstracker(DELEVERED)
	showstracker(CANCELLED)

	// Test data for fake tracker
	fmt.Println("******* Unit Testing Result *********")
	fakeTracker(DELEVERED)
	fakeTracker(CANCELLED)
}

// Fack Tracker data for unit testing

func fakeTracker(tval Tracker) {

	fmt.Println("Fake Tracker value is::", tval)
}
