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
