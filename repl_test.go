package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello    world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this is a test",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "This IS A Test",
			expected: []string{"this", "is", "a", "test"},
		},
	}

	for _, c := range cases {

		actual := CleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		actualLength := len(actual)
		expectedLength := len(c.expected)
		if !reflect.DeepEqual(expectedLength, actualLength) {
			t.Fatalf("\n Testing word slice length \n input: %v \n expected: %v, got: %v", c.input, expectedLength, actualLength)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if !reflect.DeepEqual(expectedWord, word) {
				t.Fatalf("\n Testing words in word slice \n input: %v \n expected: %v, got: %v", word, expectedWord, word)
			}
		}

	}

}
