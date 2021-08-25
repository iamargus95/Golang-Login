package main

import (
	"testing"
)

func TestJosephusSimulation(t *testing.T) {
	if josephusSimulation(41) != 19 {
		t.Errorf("Test failed, Expected result is 19.")
	}
}

func TestJosephusSimulationArray(t *testing.T) {
	var testArray = []struct {
		input    int
		expected int
	}{
		{4, 1},
		{22, 13},
		{53, 43},
		{999, 975},
	}

	for _, test := range testArray {
		if output := josephusSimulation(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestValidateInput(t *testing.T) {
	var validArray = []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"13", 13},
		{"43", 43},
		{"975", 975},
	}

	for _, test := range validArray {
		if output, err := validateInput(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.input, test.expected, err)
		}
	}
}
