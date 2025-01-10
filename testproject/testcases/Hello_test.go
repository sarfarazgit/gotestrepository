package main

import (
	"testing"
	"testproject/servicepackage"
)

func TestHello(t *testing.T) {

	result := servicepackage.SayHello("Ali")
	expected := "HelloAli"

	if result != expected {
		t.Errorf("Result %q Expected %q", result, expected)
	}
}

func TestEmptyString(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := servicepackage.SayHello("Chris")
		want := "HelloChris"
		assertCorrectMessage(t, got, want)

		// if got != want {
		// 	t.Errorf("We got %q but want %q", got, want)
		// }
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := servicepackage.SayHello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
		// if got != want {
		// 	t.Errorf("We got %q but want %q", got, want)
		// }
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
