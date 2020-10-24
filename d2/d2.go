package d2

import (
	"errors"
	"math/rand"
)

const MinPlayers = 1
const MaxPlayers = 8

// can I or should I create some speicifc type for holding characters array?

var ClassicCharacters = [5]string{"Amazon", "Barbarian", "Paladin", "Sorcerer", "Necromancer"}
var ExtensionOnlyCharacters = [2]string{"Druid", "Assasin"}
var ExtensionCharacters = append(ClassicCharacters[:], ExtensionOnlyCharacters[:]...)

func IsPlayersNumberValid(playersNumber int, minPlayers int, maxPlayers int) bool {
	return playersNumber >= minPlayers && playersNumber <= maxPlayers
}

func GetGameModeCharacters(gameMode string) ([]string, error) {
	switch gameMode {
	case "1",
		"classic",
		"c":
		return ClassicCharacters[:], nil
	case "2",
		"extension",
		"lod",
		"e",
		"l":
		return ExtensionCharacters[:], nil
	}

	return nil, errors.New("Invalid game mode provided")
}

func GetCharacterAssignmentFunction(canCharactersRepeat bool) func(int, []string) []string {
	if canCharactersRepeat == true {
		return AssignRepeatableCharacters
	}

	return AssignNonRepeatbleCharacters
}

func AssignRepeatableCharacters(playersNumber int, characters []string) []string {
	playersCharacters := make([]string, playersNumber)
	charactersNumber := len(characters)

	for playerIdx := range playersCharacters {
		randomPick := rand.Intn(charactersNumber)
		randomCharacter := characters[randomPick]
		playersCharacters[playerIdx] = randomCharacter
	}

	return playersCharacters
}

func AssignNonRepeatbleCharacters(playersNumber int, characters []string) []string {
	playersCharacters := make([]string, playersNumber)
	charactersNumber := len(characters)
	characterNumberPermutations := rand.Perm(charactersNumber)

	for idx := range playersCharacters {
		var randomCharacter string

		if idx >= charactersNumber {
			randomPick := rand.Intn(charactersNumber)
			randomCharacter = characters[randomPick]
		} else {
			randomPick := characterNumberPermutations[idx]
			randomCharacter = characters[randomPick]
		}

		playersCharacters[idx] = randomCharacter
	}

	return playersCharacters
}
