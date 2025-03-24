package main

import (
	"testing"
)

func TestAveCalculator(t *testing.T) {
	grades := []float64{80.0, 90.0, 75.0}
	subjectNumber := 3
	expectedAverage := (80.0 + 90.0 + 75.0) / 3.0
	actualAverage := ave_calculator(grades, subjectNumber)

	if actualAverage != expectedAverage {
		t.Errorf("Expected average: %f, but got: %f", expectedAverage, actualAverage)
	}
}
