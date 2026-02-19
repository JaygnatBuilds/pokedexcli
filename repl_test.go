package main

import (
	"testing"
	"reflect"
)

cases := []struct{
	input	string
	expected []string
}{
	{
		input: "   hello    world    ",
		expected: []string{"hello", "world"},
	},
}


func TestCleanInput(t *testing.T) {

	for _, c := range cases {

		actual := CleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}

	}

}
