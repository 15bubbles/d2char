package interactive

import (
	"fmt"
	"testing"
)

func TestParseCanCharactersRepeatInput(t *testing.T) {
	cases := []struct {
		givenValue       string
		expectedOutput   bool
		shouldThrowError bool
	}{
		{"y", true, false},
		{"yes", true, false},
		{"true", true, false},
		{"t", true, false},
		{"n", false, false},
		{"no", false, false},
		{"f", false, false},
		{"false", false, false},
		{"somethingelse", false, true},
		{"YES", false, true},
		{"YeS", false, true},
		{"Yes", false, true},
	}

	for _, c := range cases {
		testname := fmt.Sprint(c)

		t.Run(testname, func(t *testing.T) {
			actualOutput, err := parseCanCharactersRepeatInput(c.givenValue)

			if c.shouldThrowError && err == nil {
				t.Errorf("expected to throw error")
			}

			if !c.shouldThrowError && err != nil {
				t.Errorf("unexpected error thrown")
			}

			if actualOutput != c.expectedOutput {
				t.Errorf("actual %t, expected %t", actualOutput, c.expectedOutput)
			}
		})
	}
}
