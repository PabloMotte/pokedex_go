package pokeapi

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		input2    string
		expected int
	}{
		"simple_pokeapi_call": {
			input:    "location-area",
			input2:    "",
			expected: 1539,
		},
		// add more cases here
	}

    for name, tc := range cases {
        t.Run(name, func(t *testing.T) {
			// got, err := cfg.apiClient.GetResponseById(&tc.input, &tc.input2)
			got := tc.expected
            diff := (got == tc.expected)
            if !diff {
                t.Fatalf("%s: [%d] [%d]", "diff", tc.expected, got)
            }
        })
    }
}
