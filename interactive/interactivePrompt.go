package interactive

import (
	"bufio"
	"d2char/d2"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func ClearScreen() {
	cmd := getClearCommand()
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PromptGameModeCharacters(reader io.Reader) []string {
	var gameMode string

	for {
		fmt.Printf("Select game mode:\n1 - CLASSIC\n2 - LOD EXTENSION\n")
		fmt.Fscanln(reader, &gameMode)
		characters, err := d2.GetGameModeCharacters(gameMode)

		if err != nil {
			fmt.Printf("%s\n\n", err)
			continue
		}

		return characters
	}
}

func PromptNumberOfPlayers(reader io.Reader, minPlayers int, maxPlayers int) int {
	var playersNumber int

	for {
		fmt.Printf("Provide number of players: ")
		fmt.Fscanln(reader, &playersNumber)

		if !d2.IsPlayersNumberValid(playersNumber, minPlayers, maxPlayers) {
			fmt.Printf("Invalid number of players, should be between %d and %d\n\n", minPlayers, maxPlayers)
			continue
		}

		return playersNumber
	}
}

func PromptInputPlayerNames(reader io.Reader, playersNumber int) []string {
	var playerNames []string

	for i := 1; i < playersNumber+1; i++ {
		var playerName string
		fmt.Printf("Player %d name: ", i)
		fmt.Fscanln(reader, &playerName)

		if playerName == "" {
			playerName = fmt.Sprintf("Player %d", i)
		}

		playerNames = append(playerNames, playerName)
	}

	return playerNames
}

func PromptCanCharactersRepeat(reader io.Reader) bool {
	var playerInput string

	for {
		fmt.Printf("Can characters repeat? (yes/no) \n")
		fmt.Fscanln(reader, &playerInput)
		normalizedInput := strings.ToLower(playerInput)
		canCharactersRepeat, err := parseCanCharactersRepeatInput(normalizedInput)

		if err != nil {
			fmt.Printf("%s\n\n", err)
			continue
		}

		return canCharactersRepeat
	}
}

func parseCanCharactersRepeatInput(value string) (bool, error) {
	switch value {
	case
		"yes",
		"y",
		"true",
		"t":
		return true, nil
	case
		"no",
		"n",
		"false",
		"f":
		return false, nil
	}

	return false, errors.New("Invalid option provided, valid options are: yes, y, true, t, no, n, false, f")
}

func getClearCommand() *exec.Cmd {
	if runningOS := runtime.GOOS; runningOS == "windows" {
		return exec.Command("cmd", "/c", "cls")
	}

	return exec.Command("clear")
}

func RunInteractiveSession() {
	reader := bufio.NewReader(os.Stdin)
	ClearScreen()

	for {
		rand.Seed(time.Now().UTC().UnixNano())

		// PROMPT
		characters := PromptGameModeCharacters(reader)
		playersNumber := PromptNumberOfPlayers(reader, d2.MinPlayers, d2.MaxPlayers)
		playerNames := PromptInputPlayerNames(reader, playersNumber)
		canCharactersRepeat := PromptCanCharactersRepeat(reader)

		// ASSIGN CHARACTERS
		assignCharacters := d2.GetCharacterAssignmentFunction(canCharactersRepeat)
		playerCharacters := assignCharacters(playersNumber, characters)

		// PRINT OUT ASSIGNED CHARACTERS
		ClearScreen()

		if !canCharactersRepeat && playersNumber > len(characters) {
			fmt.Printf("Number of players is greater than number of available characters, some characters will repeat\n\n")
		}

		fmt.Println("================== CHARACTERS =====================")
		for idx := range playerCharacters {
			playerName := playerNames[idx]
			playerCharacter := playerCharacters[idx]
			fmt.Printf("%s: %s\n", playerName, playerCharacter)
		}

		break
	}
}
