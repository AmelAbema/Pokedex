package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input string
		want  []string
	}{
		{
			input: "Hello World",
			want: []string{
				"hello",
				"world",
			},
		},
	}

	for _, s := range cases {
		actual := cleanInput(s.input)
		if len(actual) != len(s.want) {
			t.Errorf("The lengths are not equal: %v != %v", len(actual), len(s.want))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := s.want[i]
			if actualWord != expectedWord {
				t.Errorf("%v != %v", actualWord, expectedWord)
			}
		}
	}

}
