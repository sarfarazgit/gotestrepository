package main

import (
	"testing"
	"testproject/servicepackage"
)

func TestRepeat(t *testing.T) {

	result := servicepackage.BenchmarkRepeat()
	expected := "aaaaa"

	if result != expected {
		t.Errorf("The result %q and expected %q", result, expected)
	}
}
