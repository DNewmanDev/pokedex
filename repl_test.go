package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "          ayyeee wats good        ",
			expected: []string{"ayyeee", "wats", "good"},
		},
		{
			input:    "Three MCdOUbles PLz",
			expected: []string{"three", "mcdoubles", "plz"},
		},
		{
			input:    "          Jeffrey    lookin        ",
			expected: []string{"jeffrey", "lookin"},
		},
		{
			input:    "",
			expected: nil,
		},
		{
			input:    "       ",
			expected: nil,
		},
	}

	for _, testcase := range cases {
		actual := cleanInput(testcase.input)

		if len(actual) != len(testcase.expected) {
			t.Errorf("Input: '%s' \n Expected length '%d', got '%d' ", testcase.input, len(actual), len(testcase.expected))
		}

		for i := range actual {

			if actual[i] != testcase.expected[i] {
				t.Errorf("Incorrect match. Expected '%s' , got '%s", actual[i], testcase.expected[i])
			}
		}
	}
}
