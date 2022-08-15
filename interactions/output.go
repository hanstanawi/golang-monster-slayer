package interactions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundStats struct {
	ActionType       string
	PlayerAttackVal  int
	PlayerHealVal    int
	MonsterAttackVal int
	PlayerHealth     int
	MonsterHealth    int
}

func (roundStat *RoundStats) PrintStats() {
	switch roundStat.ActionType {
	case "ATTACK":
		fmt.Printf("Player attacked monster for %v damage.\n", roundStat.PlayerAttackVal)
	case "HEAL":
		fmt.Printf("Player healed for %v\n", roundStat.PlayerHealVal)
	case "SPECIAL ATTACK":
		fmt.Printf("Player perfomed a strong attack against monster for %v damage.\n", roundStat.PlayerAttackVal)
	default:
		break
	}
	fmt.Printf("Monster attacked player for %v damage.\n", roundStat.MonsterAttackVal)
	fmt.Printf("Player Health: %v\n", roundStat.PlayerHealth)
	fmt.Printf("Monster Health: %v\n\n", roundStat.MonsterHealth)
}

func PrintGreeting() {
	monsterSlayerFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	monsterSlayerFigure.Print()
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func PrintAvailableActions(isSpecialRound bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")
	if isSpecialRound {
		fmt.Println("(3) Special Attack")
	}
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	gameOverFigure := figure.NewColorFigure("GAME OVER", "", "red", true)
	gameOverFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundStats) {
	exPath, err := os.Executable()

	if err != nil {
		fmt.Println("Writing log file failed")
		return
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt")
	// file, err := os.Create("gamelog.txt") // For "go run"

	if err != nil {
		fmt.Println("Writing log file failed. Exitting...")
		return
	}

	for index, round := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                round.ActionType,
			"Player Attack Damage":  fmt.Sprint(round.PlayerAttackVal),
			"Player Heal Value":     fmt.Sprint(round.PlayerHealVal),
			"Monster Attack Damage": fmt.Sprint(round.MonsterAttackVal),
			"Player Health":         fmt.Sprint(round.PlayerHealth),
			"Monster Health":        fmt.Sprint(round.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing log data failed! Exiting...")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote log data to gamelog.txt file")
}
