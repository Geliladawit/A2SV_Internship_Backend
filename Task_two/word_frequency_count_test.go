package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "Simple Sentence",
			input:    "This is a simple sentence",
			expected: map[string]int{"this": 1, "is": 1, "a": 1, "simple": 1, "sentence": 1},
		},
		{
			name:     "Repeated Words",
			input:    "The quick brown fox jumps over the lazy dog the",
			expected: map[string]int{"the": 3, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1},
		},
		{
			name:     "Empty String",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "String with only spaces",
			input:    "   ",
			expected: map[string]int{},
		},
		{
			name:     "Mixed Case",
			input:    "The the THE",
			expected: map[string]int{"the": 3},
		},
		{
			name:     "Punctuation", // Test that punctuation doesn't affect the word count when split by spaces
			input:    "Hello, world! Hello world.",
			expected: map[string]int{"hello,": 1, "world!": 1, "hello": 1, "world.": 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := wordFrequency(tc.input)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("For input '%s', expected %v, got %v", tc.input, tc.expected, actual)
			}
		})
	}
}
