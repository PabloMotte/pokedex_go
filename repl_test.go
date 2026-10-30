package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple_string": {
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		"null_string":{
			input:    "",
			expected: []string{},
		},
		"simple_sentence": {
			input:    "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
		// add more cases here
	}

    for name, tc := range cases {
        t.Run(name, func(t *testing.T) {
            got := cleanInput(tc.input)
			if len(got) != len(tc.expected) {
				t.Errorf("Returned len %d, expecting len %d", len(got), len(tc.expected))
			}
            diff := cmp.Diff(tc.expected, got)
            if diff != "" {
                t.Fatalf("%s", diff)
            }
        })
    }
}
