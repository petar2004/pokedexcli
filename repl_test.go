package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " nope nope nope ",
			expected: []string{"nope", "nope", "nope"},
		},
		{
			input:    " Wowoww woww ",
			expected: []string{"wowoww", "woww"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length is not the same")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word(%q) == %q, want %q", word, word, expectedWord)
			}
		}
	}
}
