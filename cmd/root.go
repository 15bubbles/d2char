package cmd

import (
	"d2char/d2"
	"d2char/interactive"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var gameMode string
var players []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "d2char [players]",
	Short: `Randomly selects Diablo 2 characters to players`,
	Long: `
Randomly selects Diablo 2 characters to players given as arguments.

If no players were given or no flags were set it will run interactive prompt asking few questions
to help with selection of characters for players.

For non-interactive way of assigning characters you can either pass player names as arguments or
specify number of players (-p/--players). If both are given, player name arguments take precedence.
`,
	Run: func(cmd *cobra.Command, args []string) {
		totalFlags := cmd.Flags().NFlag()
		totalArgs := cmd.Flags().NArg()

		rand.Seed(time.Now().UTC().UnixNano())

		if totalFlags == 0 && totalArgs == 0 {
			interactive.RunInteractiveSession()
			return
		}

		var playersNumber int
		var playersNames []string
		gameMode, _ := cmd.Flags().GetString("mode")
		canCharactersRepeat, _ := cmd.Flags().GetBool("repeat")

		if totalArgs > 0 {
			playersNumber = totalArgs
			playersNames = args
		} else {
			playersNumber, _ = cmd.Flags().GetInt("players")
			playersNames = make([]string, playersNumber)
			for idx := range playersNames {
				playersNames[idx] = fmt.Sprintf("Player %d", idx+1)
			}
		}

		if !d2.IsPlayersNumberValid(playersNumber, d2.MinPlayers, d2.MaxPlayers) {
			fmt.Printf("Invalid number of players specified, should be between %d and %d", d2.MinPlayers, d2.MaxPlayers)
			return
		}

		if characters, err := d2.GetGameModeCharacters(gameMode); err == nil {
			assignFunction := d2.GetCharacterAssignmentFunction(canCharactersRepeat)
			assignedCharacters := assignFunction(playersNumber, characters)

			if playersNumber > len(characters) {
				fmt.Printf("Number of players is greater than number of available characters, some characters will repeat\n\n")
			}

			for idx := range assignedCharacters {
				playerName := playersNames[idx]
				character := assignedCharacters[idx]

				fmt.Printf("%s: %s\n", playerName, character)
			}
		} else {
			fmt.Println(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("mode", "m", "extension", `game mode that will be played ("classic/c/1", "extension/e/lod/l/2")`)
	rootCmd.Flags().IntP("players", "p", 0, `number of players in a game (between 1 and 8)`)
	rootCmd.Flags().BoolP("repeat", "r", false, `specifies if characters can repeat, by default it will try to assign unique characters`)
}
