package main

import (
	"testing"
)

func TestPalindromeCheck(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple Palindrome",
			input:    "madam",
			expected: true,
		},
		{
			name:     "Not a Palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "Palindrome with spaces",
			input:    "race car",
			expected: true,
		},
		{
			name:     "Palindrome with punctuation",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "Mixed Case Palindrome",
			input:    "RaceCar",
			expected: true,
		},
		{
			name:     "Number Palindrome",
			input:    "12321",
			expected: true,
		},
		{
			name:     "Alphanumeric Palindrome",
			input:    "Madam, I'm Adam.",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := palindromeCheck(tc.input)
			if actual != tc.expected {
				t.Errorf("For input '%s', expected %t, got %t", tc.input, tc.expected, actual)
			}
		})
	}
}
