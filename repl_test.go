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
			input:    "   hello, world  ",
			expected: []string{"hello,", "world"},
		},
		{
			input:    "   Dream a little Dream !",
			expected: []string{"dream", "a", "little", "dream", "!"},
		},
		{
			input:    " He was more   Than surprised    ",
			expected: []string{"he", "was", "more", "than", "surprised"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slice length: %v is not equal to expeted length %v",
				len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expetedWord := c.expected[i]
			if word != expetedWord {
				t.Errorf("Word: %v doesn't match expected word %v", word, expetedWord)
			}
		}
	}
}
