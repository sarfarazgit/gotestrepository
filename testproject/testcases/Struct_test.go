package main

import (
	"reflect"
	"testing"
	"testproject/servicepackage"
)

func TestCircle(t *testing.T) {

	circle := servicepackage.Circle{Radius: 2.5}
	result := circle.AreaofCircle()
	wants := 10.0

	if !reflect.DeepEqual(result, wants) {
		t.Errorf("Result %.2f and wants  %.2f", result, wants)
	}
}
