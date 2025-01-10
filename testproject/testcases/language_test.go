package main

import (
	"testing"
	"testproject/servicepackage"
)

func TestLanguage(t *testing.T) {

	result := servicepackage.Hello("Hello", "Spanish")
	want := "Hola, Hello"

	if result != want {
		t.Errorf("The result %q and got %q", result, want)
	}
}
