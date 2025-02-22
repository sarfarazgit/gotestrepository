package main

import (
	"reflect"
	"testing"
	"testproject/servicepackage"
)

func TestArray(t *testing.T) {
	numbers := []int{10, 20, 30, 40, 50}
	result := servicepackage.SumofArrayElements(numbers)
	wants := 150
	if !reflect.DeepEqual(result, wants) {
		t.Errorf("The result is %q and wants %q", result, wants)
	}
}

func TestSearchElementInMap(t *testing.T) {

	t.Run("Found the key", func(t *testing.T) {
		// Data preparation
		maplements := map[string]string{"Development": "Testing", "Production": "Live"}
		findkey := "Development"
		expectedresult := "Testing"
		// Test with data
		result := servicepackage.SearchElement(maplements, findkey)
		assertStrings(t, result, expectedresult)
	})

	t.Run("Not Found the key", func(t *testing.T) {
		// Data preparation
		maplements := map[string]string{"Developemt": "Testing", "Production": "Testing"}
		expectedkey := "Staging"
		// Test with data
		result := servicepackage.SearchElement(maplements, expectedkey)
		assertStrings(t, result, expectedkey)
	})
}

func assertStrings(t testing.TB, result, expectedkey string) {
	t.Helper()
	if result != expectedkey {
		t.Errorf("Result %q Expected %q", result, expectedkey)
	}

}

func TestRepeat(t *testing.T) {

	result := servicepackage.BenchmarkRepeat()
	expected := "aaaaa"

	if result != expected {
		t.Errorf("The result %q and expected %q", result, expected)
	}
}
