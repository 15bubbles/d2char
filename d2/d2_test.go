package d2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetGameModeCharacters(t *testing.T) {
	cases := []struct {
		givenGameMode    string
		expectedOutput   []string
		shouldThrowError bool
	}{
		{"1", ClassicCharacters[:], false},
		{"classic", ClassicCharacters[:], false},
		{"c", ClassicCharacters[:], false},
		{"2", ExtensionCharacters, false},
		{"extension", ExtensionCharacters, false},
		{"lod", ExtensionCharacters, false},
		{"e", ExtensionCharacters, false},
		{"l", ExtensionCharacters, false},
	}

	for _, c := range cases {
		testname := fmt.Sprint(c)

		t.Run(testname, func(t *testing.T) {
			actualOutput, _ := GetGameModeCharacters(c.givenGameMode)
			areEqual := reflect.DeepEqual(actualOutput, c.expectedOutput)

			if !areEqual {
				t.Errorf("actual %s, expected %s", actualOutput, c.expectedOutput)
			}
		})
	}
}

func TestIsPlayersNumberValid(t *testing.T) {
	cases := []struct {
		givenPlayersNumber, givenMinPlayers, givenMaxPlayers int
		expectedOutput                                       bool
	}{
		{1, 1, 1, true},
		{2, 1, 3, true},
		{1, 2, 3, false},
		{4, 2, 3, false},
	}

	for _, c := range cases {
		testname := fmt.Sprint(c)

		t.Run(testname, func(t *testing.T) {
			actualOutput := IsPlayersNumberValid(c.givenPlayersNumber, c.givenMinPlayers, c.givenMaxPlayers)
			if actualOutput != c.expectedOutput {
				t.Errorf("actual %t, expected %t", actualOutput, c.expectedOutput)
			}
		})
	}
}
