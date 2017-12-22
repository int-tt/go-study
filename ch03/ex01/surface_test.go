package main

import (
	"math"
	"testing"
)

func TestF(t *testing.T) {
	x := float64(1000)
	y := float64(2000)
	z := f(x, y)
	if math.IsNaN(z) {
		t.Errorf("Outside the float")
	}
}

func TestInfinite(t *testing.T) {
	x := math.MaxFloat64 + 100
	y := math.MaxFloat64 + 100
	z := f(x, y)
	if !math.IsNaN(z) {
		t.Errorf("Outside the float")
	}
}
