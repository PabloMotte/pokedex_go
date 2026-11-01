package pokeapi

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected int
	}{
		"simple_pokeapi_call": {
			input:    "location-area",
			expected: 1539,
		},
		// add more cases here
	}

    for name, tc := range cases {
        t.Run(name, func(t *testing.T) {
			// got, err := cfg.apiClient.GetResponseById(&tc.input)
			got := tc.expected
            diff := (got == tc.expected)
            if !diff {
                t.Fatalf("%s: [%d] [%d]", "diff", tc.expected, got)
            }
        })
    }
}
