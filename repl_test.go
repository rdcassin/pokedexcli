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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " alpha    BETA    gamma  ",
			expected: []string{"alpha", "beta", "gamma"},
		},
		{
			input:    "  Grandpa dAd    mE  ",
			expected: []string{"grandpa", "dad", "me"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput length: %d, actual length: %d\n", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput word: %s, expected word: %s\n", word, expectedWord)
		}
		}

	}
}