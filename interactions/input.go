package interactions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(isSpecialRound bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		}
		if playerChoice == "2" {
			return "HEAL"
		}
		if playerChoice == "3" && isSpecialRound {
			return "SPECIAL ATTACK"
		}
		fmt.Println("Fetching the user input failed")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")
	playerInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Replace(playerInput, "\n", "", -1), nil
}
