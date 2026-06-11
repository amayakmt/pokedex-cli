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
			input:    "PIKACHU   bulbasaur",
			expected: []string{"pikachu", "bulbasaur"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "charmander",
			expected: []string{"charmander"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Incorrect slice length. Actual: %v. Expected: %v", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Incorrect word. Actual: %v. Expected: %v", actual[i], c.expected[i])

			}
		}

	}
}
