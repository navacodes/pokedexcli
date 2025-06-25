package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	//... implement test

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmader Bulbasaur PIKACHU",
			expected: []string{"charmader", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch. \nExpected: %v\nGot: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test

			if word != expectedWord {
				t.Errorf("Mismatch at index %d. \nExpected: %v\nGot: %v", i, c.expected, actual)
			}
		}
	}

}
